package internalcmd

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"sync"
)

var internalCmdMgr *InternalCmdManager
var mgrOnce sync.Once

func GetManager() *InternalCmdManager {
	mgrOnce.Do(func() {
		internalCmdMgr = &InternalCmdManager{m: make(map[string]handler)}
	})
	return internalCmdMgr
}

type handler func(context.Context, proto.Message) (proto.Message, error)
type InternalCmdManager struct {
	m map[string]handler
}

func (i *InternalCmdManager) Register(uri string, h handler) {
	i.m[uri] = h
}

func (i *InternalCmdManager) IsInternalCmd(msg proto.Message) bool {
	uri := proto.MessageName(msg)
	if _, h := i.m[uri]; !h {
		return false
	}
	return true
}

func (i *InternalCmdManager) HandleInternalCmd(ctx context.Context, msg proto.Message) (proto.Message, error) {
	uri := proto.MessageName(msg)
	hdl, h := i.m[uri]
	if !h {
		return nil, fmt.Errorf("msg %s is not internalCmd", uri)
	}
	return hdl(ctx, msg)
}
