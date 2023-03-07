package netutils

import (
	"fmt"
	"github.com/sandwich-go/boost/xerror"
)

func (x *ErrorResponse) Error() string {
	if x != nil {
		return fmt.Sprintf("code :%d msg:%s", x.Code, x.Message)
	}
	return ""
}

func (x *ErrorResponse) IsLogicException() bool {
	return x.LogicException
}

func (x *ErrorResponse) SetError(err error) {
	x.err = err
}

func (x *ErrorResponse) Logic() bool {
	return x.LogicException
}

func (x *ErrorResponse) Cause() error {
	if x.err != nil {
		return x.err
	}
	return x
}
func (x *ErrorResponse) Caller(skip int) (file, funcName string, line int) {
	if x.err != nil {
		return xerror.Caller(x.err, skip)
	}
	return
}
func (x *ErrorResponse) Stack() string {
	if x.err != nil {
		return xerror.Stack(x.err)
	}
	return ""
}
