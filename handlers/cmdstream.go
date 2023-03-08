package handlers

import (
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
	//"bitbucket.org/funplus/sandwich/client"
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/internalcmd"
	//"bitbucket.org/funplus/sandwich/internal/rand"
)

const streamTokenLen = 16

func HandleCmdStream(ctx context.Context, cmd *internalcmd.InternalCmd, req interface{}) (proto.Message, error) {
	if _, ok := req.(*internal_command.CmdStream); !ok {
		return nil, internalcmd.ErrReqType
	}
	/*req.Addr = w.streamAddr
	req.Token = rand.String(streamTokenLen)
	w.cachedTokens.Set(req.Token, &client.StreamArgs{
		ID:    req.Meta[client.MetaKeyStreamID],
		Addr:  w.streamAddr,
		Token: req.Token,
		Meta:  req.Meta,
	}, time.Minute*5)
	return req, nil
	return s.StreamHandler(ctx, req)*/
	return msg, nil
}
