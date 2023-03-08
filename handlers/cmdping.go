package handlers

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/protocol/netutils"
	"time"
)

func HandleCmdPing(ctx context.Context, msg proto.Message) (proto.Message, error) {
	return netutils.CmdPingAck{Timestamp: time.Now().Unix()}, nil
}
