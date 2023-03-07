package handlers

import (
	"bitbucket.org/funplus/sandwich/client"
	"context"
	"github.com/golang/protobuf/proto"
	"k8s.io/apimachinery/pkg/util/rand"
)

func HandleCmdStream(ctx context.Context, msg proto.Message) (proto.Message, error) {
	req.Addr = w.streamAddr
	req.Token = rand.String(streamTokenLen)
	w.cachedTokens.Set(req.Token, &client.StreamArgs{
		ID:    req.Meta[client.MetaKeyStreamID],
		Addr:  w.streamAddr,
		Token: req.Token,
		Meta:  req.Meta,
	}, time.Minute*5)
	return req, nil
	return s.StreamHandler(ctx, req)
}
