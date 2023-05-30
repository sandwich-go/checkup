package checkup

import (
	"context"
	"fmt"
	"github.com/sandwich-go/boost/singleflight"
	"github.com/sandwich-go/boost/z"
	"github.com/sandwich-go/checkup/protocol/gen/golang/common"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	"google.golang.org/protobuf/proto"
	"time"
)

var PacketCheckup = &Packet{Uri: "internal_command.CmdCheckup"}

func init() { register(checkupPacketProcessor{}) }

type checkupPacketProcessor struct{}

func (h checkupPacketProcessor) packet() *Packet { return PacketCheckup }
func (h checkupPacketProcessor) process(ctx context.Context, fight *singleflight.Group, cc *Options) ([]byte, error) {
	var resp *internal_command.CmdCheckup
	tsStart := time.Now()
	rr, err := fight.Do(h.packet().Uri, func() (interface{}, error) {
		return cc.GetDevopsCheckup()(ctx), nil
	})
	if err != nil {
		resp = &internal_command.CmdCheckup{Code: common.ErrorCode_Unknown.NumberInt32(), Message: err.Error()}
	} else {
		resp = rr.(*internal_command.CmdCheckup)
		if len(resp.CustomMeasurements) == 0 {
			resp.CustomMeasurements = z.StringToBytes(fmt.Sprintf("%s_%s", tsStart, time.Now().Sub(tsStart)))
		}
	}
	return proto.Marshal(resp)
}
func (h checkupPacketProcessor) decode(in []byte) (proto.Message, error) {
	var resp = &internal_command.CmdCheckup{}
	err := proto.Unmarshal(in, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
