package checkup

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/boost/z"
	"github.com/sandwich-go/checkup/protocol/gen/golang/common"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	"time"
)

func handleCmdCheckUp(ctx context.Context, r interface{}, m *Manager) proto.Message {
	_, ok := r.(*internal_command.CmdCheckup)
	if !ok {
		return &internal_command.CmdCheckup{Code: common.ErrorCode_Unknown.NumberInt32(), Message: ErrReqType.Error()}
	}

	tsStart := time.Now()
	rr, err := m.Fight.Do("checkup", func() (interface{}, error) {
		return m.Cc.GetDevopsCheckup()(ctx), nil
	})
	if err != nil {
		return &internal_command.CmdCheckup{Code: common.ErrorCode_Unknown.NumberInt32(), Message: err.Error()}
	}
	resp := rr.(*internal_command.CmdCheckup)
	if resp.CustomMeasurements == nil || len(resp.CustomMeasurements) == 0 {
		resp.CustomMeasurements = z.StringToBytes(fmt.Sprintf("%s_%s", tsStart, time.Now().Sub(tsStart)))
	}
	return resp
}
