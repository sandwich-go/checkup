package checkup

import (
	"encoding/json"
)

const magicNumber byte = 0x7B

// Codec is the interface that wraps the basic Marshal and Unmarshal methods.
type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// Marshal ErrMarshalType is returned when the type of the value passed to Marshal is not supported.
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

// Unmarshal ErrUnmarshalNotInternalCmd is returned when the magicNumber verify failed.
func Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}
