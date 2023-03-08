package internalcmd

import (
	"github.com/sandwich-go/internalcmd/handlers"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"github.com/sandwich-go/internalcmd/protocol/netutils"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&netutils.CmdPing{})), handlers.HandleCmdPing)
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&internal_command.CmdStream{})), handlers.HandleCmdStream)
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&netutils.CmdCheckup{})), handlers.HandleCmdCheckUp)
}
