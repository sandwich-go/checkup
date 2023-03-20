package checkup

import (
	"encoding/json"
)

const magicNumber byte = 0x7B

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
	return b, nil
}

func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}
