package checkup

import (
	"bytes"
	"context"
	"encoding/base64"
	"github.com/sandwich-go/boost/singleflight"
	"github.com/sandwich-go/boost/xpanic"
	"github.com/sandwich-go/boost/z"
	"google.golang.org/protobuf/proto"
)

type packetProcessor interface {
	packet() *Packet
	decode(in []byte) (proto.Message, error)
	process(ctx context.Context, fight *singleflight.Group, cc *Options) ([]byte, error)
}

var processors = make(map[URI]packetProcessor)

func register(processor packetProcessor) {
	xpanic.WhenTrue(processor == nil, "register processor is nil")
	uri := processor.packet().Uri
	if _, dup := processors[uri]; dup {
		xpanic.WhenTrue(dup, "register called twice for processor %s", uri)
	}
	processors[uri] = processor
}

type handler struct {
	bytes map[URI][]byte
	cc    *Options
	fight *singleflight.Group
}

func New(opts ...Option) Handler {
	cc := NewOptions(opts...)
	h := &handler{cc: cc, fight: &singleflight.Group{}, bytes: make(map[URI][]byte)}
	for k, v := range processors {
		bs, err := cc.Codec.Marshal(v.packet())
		xpanic.WhenError(err)
		h.bytes[k] = bs
	}
	return h
}

func (h handler) Bytes(uri URI) []byte { return h.bytes[uri] }
func (h handler) Check(in []byte) Valid {
	for _, bs := range h.bytes {
		if bytes.Compare(in, bs) == 0 {
			return true
		}
	}
	return false
}
func (h handler) Decode(in []byte) (URI, proto.Message, error) {
	var p = &Packet{}
	err := h.cc.Codec.Unmarshal(in, p)
	if err != nil {
		return "", nil, err
	}
	if len(p.Uri) == 0 || len(h.bytes[p.Uri]) == 0 {
		return "", nil, ErrUnknownPacket
	}
	var msg proto.Message
	if len(p.Raw) > 0 {
		var raw []byte
		raw, err = base64.StdEncoding.DecodeString(z.BytesToString(p.Raw))
		if err == nil {
			msg, err = processors[p.Uri].decode(raw)
		}
		if err != nil {
			return "", nil, err
		}
	}
	return p.Uri, msg, nil
}
func (h handler) Handle(ctx context.Context, in []byte) ([]byte, Valid, error) {
	if !h.Check(in) {
		return nil, false, nil
	}
	uri, _, err := h.Decode(in)
	if err != nil {
		return nil, true, err
	}
	var raw []byte
	raw, err = processors[uri].process(ctx, h.fight, h.cc)
	if err != nil {
		return nil, true, err
	}
	if len(raw) > 0 {
		raw = z.StringToBytes(base64.StdEncoding.EncodeToString(raw))
	}
	var out []byte
	out, err = h.cc.Codec.Marshal(&Packet{Uri: uri, Raw: raw})
	return out, true, err
}
