package handlers

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
)

func HandleCmdCheckUp(ctx context.Context, cmd *internalcmd.InternalCmd, req interface{}) (proto.Message, error) {
	if _, ok := req.(*internal_command.CmdCheckup); !ok {
		return nil, internalcmd.ErrReqType
	}
	return msg, nil
}
