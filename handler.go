package checkup

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/sandwich-go/boost/singleflight"
	"github.com/sandwich-go/boost/xpanic"
	"github.com/sandwich-go/boost/z"
	"google.golang.org/protobuf/proto"

	"github.com/sandwich-go/checkup/protocol/gen/golang/common"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
)

const uri = "internal_command.CmdCheckup"

var uriSuffix string

func init() {
	ss := strings.Split(uri, ".")
	uriSuffix = ss[len(ss)-1]
}

type handler struct {
	bytes, errBytes []byte
	cc              *Options
	fight           *singleflight.Group
}

func New(opts ...Option) Handler {
	cc := NewOptions(opts...)
	h := &handler{cc: cc, fight: &singleflight.Group{}}
	h.bytes = z.StringToBytes(uri)
	var err error
	h.errBytes, err = cc.Codec.Marshal(&Packet{Uri: uri, Raw: h.bytes})
	xpanic.WhenError(err)
	return h
}

func (h handler) unmarshal(in []byte) (*internal_command.CmdCheckup, error) {
	if bytes.Compare(in, h.errBytes) == 0 {
		return nil, ErrHandleRequest
	}
	var p = &Packet{}
	err := h.cc.Codec.Unmarshal(in, p)
	if err != nil {
		return nil, err
	}
	if p.Uri != uri {
		return nil, ErrUnknownPacket
	}
	var resp = &internal_command.CmdCheckup{}
	if len(p.Raw) > 0 {
		var raw []byte
		raw, err = base64.StdEncoding.DecodeString(z.BytesToString(p.Raw))
		if err != nil {
			return nil, err
		}
		if err = proto.Unmarshal(raw, resp); err != nil {
			return nil, err
		}
	}
	return resp, nil
}
func (h handler) onMarshalError(err error) []byte {
	if f := h.cc.GetOnError(); f != nil {
		f(err)
	}
	return h.errBytes
}
func (h handler) marshal(in *internal_command.CmdCheckup) []byte {
	raw, err := proto.Marshal(in)
	if err != nil {
		return h.onMarshalError(err)
	}
	if len(raw) > 0 {
		raw = z.StringToBytes(base64.StdEncoding.EncodeToString(raw))
	}
	var out []byte
	out, err = h.cc.Codec.Marshal(&Packet{Uri: uri, Raw: raw})
	if err != nil {
		return h.onMarshalError(err)
	}
	return out
}
func (h handler) filter(rr interface{}, ts ...time.Time) *internal_command.CmdCheckup {
	var resp *internal_command.CmdCheckup
	switch v := rr.(type) {
	case error:
		resp = &internal_command.CmdCheckup{Code: common.ErrorCode_Unknown.NumberInt32(), Message: v.Error()}
	case *internal_command.CmdCheckup:
		resp = v
	}
	if resp == nil {
		resp = &internal_command.CmdCheckup{}
	}
	if len(ts) > 0 && len(resp.CustomMeasurements) == 0 {
		resp.CustomMeasurements = z.StringToBytes(fmt.Sprintf("%s_%s", ts[0], time.Now().Sub(ts[0])))
	}
	return resp
}
func (h handler) RequestBytes() []byte      { return h.bytes }
func (h handler) IsRequestPath(s string) Is { return strings.HasSuffix(s, uriSuffix) }
func (h handler) HandleIfRequestBytes(ctx context.Context, in []byte) ResultInterface {
	if bytes.Compare(in, h.bytes) != 0 {
		return &ProcessResult{
			Resp:      nil,
			IsCheckup: false,
		}
	}
	var tsStart = time.Now()
	if rr, err := h.fight.Do(uri, func() (interface{}, error) {
		if f := h.cc.GetDevopsCheckup(); f != nil {
			return f(ctx), nil
		}
		return nil, nil
	}); err != nil {
		return &ProcessResult{
			Resp:      h.marshal(h.filter(err)),
			IsCheckup: true,
			Result:    common.ErrorCode_Unknown,
		}
	} else {
		rsp := h.filter(rr, tsStart)
		return &ProcessResult{
			Resp:      h.marshal(rsp),
			IsCheckup: true,
			Result: func() common.ErrorCode {
				if rsp.Code != common.ErrorCode_OK.NumberInt32() {
					return common.ErrorCode_Unknown
				}
				return common.ErrorCode_OK
			}(),
		}
	}
}
func (h handler) HandleResponseBytes(in []byte) (*internal_command.CmdCheckup, error) {
	return h.unmarshal(in)
}
