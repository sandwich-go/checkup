// Code generated by protokitgo. DO NOT EDIT.
package internal_command

import (
	proto "google.golang.org/protobuf/proto"
	"sync"
)

var _ proto.MarshalOptions
var _ sync.Pool

var cmdCheckupPool = sync.Pool{New: func() interface{} { return new(CmdCheckup) }}

func NewCmdCheckup() *CmdCheckup {
	return cmdCheckupPool.Get().(*CmdCheckup)
}

func (x *CmdCheckup) Release() {
	if x != nil {
		x.Reset()
		cmdCheckupPool.Put(x)
	}
}