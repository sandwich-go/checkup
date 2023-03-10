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
	/*if data[0] != magicNumber {
		return ErrUnmarshalNotInternalCmd
	}*/
	err := json.Unmarshal(data, v)
	return err
}
