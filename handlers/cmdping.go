package handlers

import (
	"context"
	"github.com/golang/protobuf/proto"
	"time"
)

func HandleCmdPing(ctx context.Context, msg proto.Message) (proto.Message, error) {
	return netutils.CmdPingAck{Timestamp: time.Now().Unix()}, nil
}
