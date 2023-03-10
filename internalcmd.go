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

func MarshalInternalCmd(b []byte) []byte {
	data := make([]byte, len(b)+1)
	data[0] = magicNumber
	copy(data[1:], b)
	return data
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

func HandleInternalCmd(ctx context.Context, b []byte, opts ...interface{}) (bo []byte, err error) {
	icmd := &InternalCmd{}
	err = Unmarshal(b, icmd)
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
