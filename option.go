package internalcmd

import (
	"context"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/common"
	"github.com/sandwich-go/internalcmd/protocol/gen/golang/internal_command"
)

//go:generate optionGen  --option_return_previous=false
func OptionsOptionDeclareWithDefault() interface{} {
	return map[string]interface{}{
		"DevopsCheckup": func(ctx context.Context) *internal_command.CmdCheckup {
			return &internal_command.CmdCheckup{Code: common.ErrorCode_OK.NumberInt32(), Message: "default ok"}
		},
	}
}
