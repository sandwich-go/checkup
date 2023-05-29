package checkup

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sandwich-go/boost/singleflight"
)

const magicNumber byte = 0x7B

// InternalCmd 和 Raw 都以json来序列化
// 新增Internal Cmd类型的消息时, 需要新增定义到 protocol/protos/internal_command/internal.proto 中
// 发送者只要能保证发送的格式符合 InternalCmd 和 Raw 以json格式解析即可
type InternalCmd struct {
	Uri         string `json:"uri"`
	Raw         []byte `json:"raw"`
	PassThrough string `json:"passThrough,omitempty"`
}

func CheckMagicNumber(b []byte) bool {
	if len(b) < 1 {
		return false
	}
	if b[0] != magicNumber {
		return false
	}
	return true
}

func NewFrame(msg proto.Message) []byte {
	b, _ := json.Marshal(&InternalCmd{
		Uri: proto.MessageName(msg),
		Raw: func() []byte {
			b, err := json.Marshal(msg)
			if err != nil {
				return nil
			}
			return b
		}(),
	})

	return b
}

type Manager struct {
	Cc    *Options
	Fight *singleflight.Group
}

func NewManager(opts ...Option) *Manager {
	cfg := NewOptions(opts...)
	return &Manager{Cc: cfg, Fight: &singleflight.Group{}}
}

func (m *Manager) HandleInternalCmd(ctx context.Context, bytesIn []byte) (bo []byte, b bool) {
	b = false
	ok := CheckMagicNumber(bytesIn)
	if !ok {
		return nil, false
	}

	icmd := &InternalCmd{}
	err := json.Unmarshal(bytesIn, icmd)
	if err != nil || icmd.Uri == "" {
		m.LogError(err)
		return nil, false
	}

	ret := GetRouter().Handle(ctx, icmd, m)

	bo, err1 := json.Marshal(&InternalCmd{
		Uri: proto.MessageName(ret),
		Raw: func() []byte {
			b, err2 := json.Marshal(ret)
			if err2 != nil {
				return nil
			}
			return b
		}(),
		PassThrough: "",
	})
	if err1 != nil {
		m.LogError(fmt.Errorf("unmarshal internal cmd failed. Err %s", err1.Error()))
	}

	return bo, true
}

func (m *Manager) LogError(err error) {
	m.Cc.GetLogErrorFunc()(err)
}
