package internalcmd

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	"time"
)

func handleCmdStream(ctx context.Context, r interface{}, opts ...interface{}) (proto.Message, error) {
	req, ok := r.(*internal_command.CmdStream)
	if !ok {
		return nil, ErrReqType
	}
	req.Addr = GetOptions().GetIStream().GetStreamAddr()
	req.Token = GetOptions().GetIStream().NewStreamToken()
	GetOptions().GetIStream().GetCache().Set(req.Token, &StreamArgs{
		ID:    req.Meta[MetaKeyStreamID],
		Addr:  GetOptions().GetIStream().GetStreamAddr(),
		Token: req.Token,
		Meta:  req.Meta,
	}, time.Minute*5)
	return req, nil
}
