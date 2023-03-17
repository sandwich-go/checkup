package internalcmd

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

// InternalCmd 和 Raw 都以json来序列化
// 新增Internal Cmd时需要新增定义到 protocol/protos/internal_command/internal.proto 中
// 发送者只要能保证发送的格式符合 InternalCmd 和 Raw 以json格式解析即可
type InternalCmd struct {
	Uri         string `json:"uri"`
	Raw         []byte `json:"raw"`
	PassThrough string `json:"passThrough,omitempty"`
}

func IsInternalCmd(b []byte) bool {
	if len(b) < 1 {
		return false
	}
	if b[0] != magicNumber {
		return false
	}
	return true
}

func NewInternalCmd(msg proto.Message, pt string) []byte {
	b, _ := json.Marshal(&InternalCmd{
		Uri: proto.MessageName(msg),
		Raw: func() []byte {
			b, err := json.Marshal(msg)
			if err != nil {
				return nil
			}
			return b
		}(),
		PassThrough: pt,
	})

	return b
}

func HandleInternalCmd(ctx context.Context, bytesIn []byte, opts ...interface{}) (bo []byte, err error) {
	ok := IsInternalCmd(bytesIn)
	if !ok {
		return nil, nil
	}

	icmd := &InternalCmd{}
	err = Unmarshal(bytesIn, icmd)
	if err != nil {
		return nil, err
	}
	if icmd.Uri == "" {
		return nil, ErrUriNil
	}

	ret, err := GetRouter().Handle(ctx, icmd)
	err = warpError(icmd.Uri, err)

	bo, _ = Marshal(&InternalCmd{
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

	return
}
