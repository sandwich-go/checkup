package internalcmd

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/sandwich-go/boost/xencoding"
	"github.com/sandwich-go/boost/z"
)

const magicNumber byte = 0xCD

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type InternalCmd struct {
	Uri         string `json:"uri"`
	Raw         []byte `json:"raw"`
	PassThrough string `json:"passThrough"`
}

func (i InternalCmd) Marshal(v interface{}) ([]byte, error) {
	if _, ok := v.(*InternalCmd); !ok {
		return nil, errors.New("Marshal InternalCmd failed, v is not an InternalCmd")
	}
	b, err := xencoding.GetCodec("json").Marshal(v)
	if err != nil {
		return nil, err
	}
	data := make([]byte, len(b)+1)
	data[0] = magicNumber
	copy(data[1:], b)
	return data, nil
}

func (i InternalCmd) Unmarshal(data []byte, v interface{}) error {
	if data[0] != magicNumber {
		return errors.New("Unmarshal InternalCmd err, magicNumber verify failed")
	}
	return xencoding.GetCodec("json").Unmarshal(data[1:], v)
}

func readString(data []byte, pos uint32, tlvLen uint32) (string, uint32) {
	bb, pos := readBytes(data, pos, tlvLen)
	if bb == nil {
		return "", pos
	}
	return z.BytesToString(bb), pos
}

func readBytes(data []byte, pos uint32, tlvLen uint32) ([]byte, uint32) {
	toReadLen := uint32(0)
	if tlvLen == 1 {
		toReadLen = uint32(data[pos])
	} else if tlvLen == 2 {
		toReadLen = uint32(binary.LittleEndian.Uint16(data[pos : pos+2]))
	} else if tlvLen == 4 {
		toReadLen = binary.LittleEndian.Uint32(data[pos : pos+4])
	} else {
		panic(fmt.Sprintf("not support for tlvLen:%d", tlvLen))
	}
	pos += tlvLen
	return data[pos : pos+toReadLen], pos + toReadLen
}

func writeString(data []byte, pos uint32, tlvLen uint32, toWrite string) uint32 {
	return writeBytes(data, pos, tlvLen, z.StringToBytes(toWrite))
}
func writeBytes(data []byte, pos uint32, tlvLen uint32, toWrite []byte) uint32 {
	sLen := uint32(len(toWrite))
	if tlvLen == 1 {
		data[pos] = byte(sLen)
	} else if tlvLen == 2 {
		binary.LittleEndian.PutUint16(data[pos:pos+tlvLen], uint16(sLen))
	} else if tlvLen == 4 {
		binary.LittleEndian.PutUint32(data[pos:pos+tlvLen], uint32(sLen))
	} else {
		panic(fmt.Sprintf("not support for tlvLen:%d", tlvLen))
	}
	pos += tlvLen
	if sLen > 0 {
		copy(data[pos:pos+sLen], toWrite)
	}
	pos += sLen
	return pos
}
