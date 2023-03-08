package internalcmd

import (
	"encoding/json"
	"fmt"
)

const magicNumber byte = 0xCD

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// InternalCmd 和 Raw 都以json来序列化
// 新增Internal Cmd时需要新增定义到 protocol/protos/internal_command/internal.proto 中
// 发送者只要能保证发送的格式符合 InternalCmd 和 Raw 以json格式解析即可
type InternalCmd struct {
	Uri         string `json:"uri"`
	Raw         []byte `json:"raw"`
	PassThrough string `json:"passThrough"`
}

func Marshal(v interface{}) ([]byte, error) {
	if _, ok := v.(*InternalCmd); !ok {
		return nil, ErrMarshalType
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	data := make([]byte, len(b)+1)
	data[0] = magicNumber
	copy(data[1:], b)
	fmt.Println(data)
	return data, nil
}

func Unmarshal(data []byte, v interface{}) error {
	if data[0] != magicNumber {
		return ErrUnmarshalNotIM
	}
	err := json.Unmarshal(data[1:], v)
	return err
}
