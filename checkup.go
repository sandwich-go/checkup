package checkup

import (
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
)

// Packet 传输的包体
type Packet struct {
	Uri URI    `json:"uri"`
	Raw []byte `json:"raw"`
}

// ErrUnknownPacket Decode 或 Handle 的时候，若不合法，则抛出该错误
var ErrUnknownPacket = errors.New("unknown packet")

// Valid 是否为有效的消息
type Valid = bool

// URI uri
type URI = string

// Codec 序列化/反序列化
type Codec interface {
	// Marshal 序列化
	Marshal(v any) ([]byte, error)
	// Unmarshal 反序列化
	Unmarshal(data []byte, v any) error
}

// Handler 处理器
type Handler interface {
	// Bytes URI 对应的字节数组
	Bytes(uri URI) []byte
	// Check 检查是否有效
	Check(in []byte) Valid
	// Handle 处理
	Handle(ctx context.Context, in []byte) ([]byte, Valid, error)
	// Decode 解析 Packet 对应的字节数组
	Decode(in []byte) (URI, proto.Message, error)
}
