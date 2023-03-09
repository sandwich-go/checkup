package internalcmd

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/boost/z"
	"github.com/sandwich-go/internalcmd/protocol/netutils"
	"time"
)

func handleCmdCheckUp(ctx context.Context, r interface{}, opts ...interface{}) (proto.Message, error) {
	_, ok := r.(*netutils.CmdCheckup)
	if !ok {
		return &netutils.CmdCheckup{Code: netutils.ErrorCode_Unknown.NumberInt32(), Message: ErrReqType.Error()}, ErrReqType
	}

	tsStart := time.Now()
	rr, err := GetCheckUpFlight().Do("checkup", func() (interface{}, error) {
		return GetOptions().GetDevopsCheckup()(ctx), nil
	})
	if err != nil {
		return &netutils.CmdCheckup{Code: netutils.ErrorCode_Unknown.NumberInt32(), Message: err.Error()}, err
	}
	cresp := rr.(*netutils.CmdCheckup)
	if cresp.CustomMeasurements == nil || len(cresp.CustomMeasurements) == 0 {
		cresp.CustomMeasurements = z.StringToBytes(fmt.Sprintf("%s_%s", tsStart, time.Now().Sub(tsStart)))
	}
	return cresp, nil
}
