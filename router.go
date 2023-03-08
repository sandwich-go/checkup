package internalcmd

import (
	"github.com/sandwich-go/internalcmd/handlers"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	//"github.com/sandwich-go/internalcmd/protocol/netutils"
	"github.com/golang/protobuf/proto"
)

func init() {
	GetManager().Register(proto.MessageName(&internal_command.CmdPing{}), handlers.HandleCmdPing)
	GetManager().Register(proto.MessageName(&internal_command.CmdStream{}), handlers.HandleCmdStream)
	GetManager().Register(proto.MessageName(&internal_command.CmdCheckup{}), handlers.HandleCmdCheckUp)
}
