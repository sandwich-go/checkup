package internalcmd

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/boost/z"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"github.com/sandwich-go/internalcmd/protocol/netutils"
	"time"
)

func handleCmdCheckUp(ctx context.Context, r interface{}, opts ...interface{}) (proto.Message, error) {
	_, ok := r.(*internal_command.CmdCheckup)
	if !ok {
		return &internal_command.CmdCheckup{Code: netutils.ErrorCode_Unknown.NumberInt32(), Message: ErrReqType.Error()}, ErrReqType
	}

	tsStart := time.Now()
	rr, err := GetCheckUpFlight().Do("checkup", func() (interface{}, error) {
		return GetOptions().GetDevopsCheckup()(ctx), nil
	})
	if err != nil {
		return &internal_command.CmdCheckup{Code: netutils.ErrorCode_Unknown.NumberInt32(), Message: err.Error()}, err
	}
	cresp := rr.(*internal_command.CmdCheckup)
	if cresp.CustomMeasurements == nil || len(cresp.CustomMeasurements) == 0 {
		cresp.CustomMeasurements = z.StringToBytes(fmt.Sprintf("%s_%s", tsStart, time.Now().Sub(tsStart)))
	}
	return cresp, nil
}
