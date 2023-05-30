package checkup

import (
	"context"
	"encoding/json"
	"github.com/sandwich-go/checkup/protocol/gen/golang/common"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
)

type LogErrorFunc func(err error)

//go:generate optionGen  --option_return_previous=false
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"DevopsCheckup": func(ctx context.Context) *internal_command.CmdCheckup {
			return &internal_command.CmdCheckup{Code: common.ErrorCode_OK.NumberInt32(), Message: "default ok"}
		},
		"Codec": (Codec)(jsonCodec{}),
	}
}

type jsonCodec struct{}

func (c jsonCodec) Marshal(v any) ([]byte, error)      { return json.Marshal(v) }
func (c jsonCodec) Unmarshal(data []byte, v any) error { return json.Unmarshal(data, v) }
