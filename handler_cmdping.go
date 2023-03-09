package internalcmd

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"time"
)

func handleCmdPing(ctx context.Context, req interface{}, opts ...interface{}) (proto.Message, error) {
	if _, ok := req.(*internal_command.CmdPing); !ok {
		return nil, ErrReqType
	}
	return &internal_command.CmdPingAck{Timestamp: time.Now().Unix()}, nil
}
