package checkup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	"sync"
)

var internalCmdMgr *Router
var mgrOnce sync.Once

var (
	ErrReqType     = errors.New("request type invalid")
	ErrMarshalType = errors.New("marshal InternalCmd failed, v is not an InternalCmd")
	ErrUriNil      = errors.New("internal cmd uri is nil")
)

func init() {
	GetRouter().RegisterHandler(proto.MessageName(&internal_command.CmdCheckup{}), handleCmdCheckUp)
}

func GetRouter() *Router {
	mgrOnce.Do(func() {
		internalCmdMgr = &Router{m: make(map[string]handler)}
	})
	return internalCmdMgr
}

type handler func(context.Context, interface{}, ...interface{}) (proto.Message, error)
type Router struct {
	m map[string]handler
}

func (i *Router) RegisterHandler(uri string, h handler) {
	i.m[uri] = h
}

func (i *Router) Handle(ctx context.Context, cmd *InternalCmd) (proto.Message, error) {
	o, ok := globalRegistry.NewObject(cmd.Uri)
	if !ok {
		return nil, fmt.Errorf("get handler failed for internal cmd uri:%s", cmd.Uri)
	}

	if err := json.Unmarshal(cmd.Raw, o); err != nil {
		return nil, err
	}

	return i.m[cmd.Uri](ctx, o)
}
