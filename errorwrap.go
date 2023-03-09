package internalcmd

import "fmt"

type internalCmdError struct {
	Uri string
	Err error
}

func (e *internalCmdError) Error() string {
	return fmt.Sprintf("internalcmd uir:%s, err:%v", e.Uri, e.Err)
}

func warpError(uri string, err error) error {
	if err == nil {
		return nil
	}
	return &internalCmdError{Uri: uri, Err: err}
}
