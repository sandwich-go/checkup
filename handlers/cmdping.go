package handlers

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"time"
)

func HandleCmdPing(ctx context.Context, msg proto.Message) (proto.Message, error) {
	return &internal_command.CmdPingAck{Timestamp: time.Now().Unix()}, nil
}
