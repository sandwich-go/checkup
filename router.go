package internalcmd

import (
	"bitbucket.org/funplus/sandwich-plugins/internalcmd/handlers"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func init() {
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&netutils.CmdPing{})), handlers.HandleCmdPing)
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&internal_cmd.CmdStream{})), handlers.HandleCmdStream)
	GetManager().Register(string(protoreflect.MessageDescriptor.FullName(&netutils.CmdCheckup{})), handlers.HandleCmdCheckUp)
}
