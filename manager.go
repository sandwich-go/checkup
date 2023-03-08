package internalcmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/handlers"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"sync"
)

var internalCmdMgr *Router
var mgrOnce sync.Once

var (
	ErrReqType        = errors.New("Request type invalid")
	ErrMarshalType    = errors.New("Marshal InternalCmd failed, v is not an InternalCmd")
	ErrUnmarshalNotIM = errors.New("Unmarshal InternalCmd err, magicNumber verify failed")
)

func init() {
	GetRouter().RegisterHandler(proto.MessageName(&internal_command.CmdPing{}), handlers.HandleCmdPing)
	GetRouter().RegisterHandler(proto.MessageName(&internal_command.CmdStream{}), handlers.HandleCmdStream)
	GetRouter().RegisterHandler(proto.MessageName(&internal_command.CmdCheckup{}), handlers.HandleCmdCheckUp)
}

func GetRouter() *Router {
	mgrOnce.Do(func() {
		internalCmdMgr = &Router{m: make(map[string]handler)}
	})
	return internalCmdMgr
}

type handler func(context.Context, *InternalCmd, interface{}) (proto.Message, error)
type Router struct {
	m map[string]handler
}

func (i *Router) RegisterHandler(uri string, h handler) {
	i.m[uri] = h
}

func (i *Router) Handle(ctx context.Context, cmd *InternalCmd) (proto.Message, error) {
	o, ok := globalRegistry.NewObject(cmd.Uri)
	if !ok {
		return nil, fmt.Errorf("Get handler failed for internal cmd uir:%s", cmd.Uri)
	}

	if err := json.Unmarshal(cmd.Raw, o); err != nil {
		return nil, err
	}

	return i.m[cmd.Uri](ctx, cmd, o)
}

func (i *Router) IsInternalCmd(b []byte) bool {
	if len(b) < 1 {
		return false
	}
	if b[0] != magicNumber {
		return false
	}
	return true
}

func (i *Router) HandleInternalCmd(ctx context.Context, b []byte) ([]byte, error) {
	if !i.IsInternalCmd(b) {
		return nil, nil
	}

	icmd := &InternalCmd{}
	err := Unmarshal(b, icmd)
	if err != nil {
		return nil, err
	}
	if icmd.Uri == "" {
		return nil, errors.New("internal cmd uri is nil")
	}

	ret, err := i.Handle(ctx, icmd)
	if err != nil {
		return nil, err
	}

	bo, err := Marshal(&InternalCmd{
		Uri: proto.MessageName(ret),
		Raw: func() []byte {
			b, err := json.Marshal(ret)
			if err != nil {
				return nil
			}
			return b
		}(),
		PassThrough: "",
	})
	if err != nil {
		return nil, err
	}

	return bo, nil
}
