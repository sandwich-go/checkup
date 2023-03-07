package internalcmd

import (
	"reflect"
	"testing"
)

func TestInternalCmd_Marshal(t *testing.T) {
	type fields struct {
		Uri         string
		Raw         []byte
		PassThrough string
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"correct", {"netutils.Ping"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternalCmd{
				Uri:         tt.fields.Uri,
				Raw:         tt.fields.Raw,
				PassThrough: tt.fields.PassThrough,
			}
			got, err := i.Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternalCmd_Unmarshal(t *testing.T) {
	type fields struct {
		Uri         string
		Raw         []byte
		PassThrough string
	}
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternalCmd{
				Uri:         tt.fields.Uri,
				Raw:         tt.fields.Raw,
				PassThrough: tt.fields.PassThrough,
			}
			if err := i.Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readBytes(t *testing.T) {
	type args struct {
		data   []byte
		pos    uint32
		tlvLen uint32
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readBytes(tt.args.data, tt.args.pos, tt.args.tlvLen)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readBytes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readBytes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_readString(t *testing.T) {
	type args struct {
		data   []byte
		pos    uint32
		tlvLen uint32
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := readString(tt.args.data, tt.args.pos, tt.args.tlvLen)
			if got != tt.want {
				t.Errorf("readString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("readString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_writeBytes(t *testing.T) {
	type args struct {
		data    []byte
		pos     uint32
		tlvLen  uint32
		toWrite []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeBytes(tt.args.data, tt.args.pos, tt.args.tlvLen, tt.args.toWrite); got != tt.want {
				t.Errorf("writeBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeString(t *testing.T) {
	type args struct {
		data    []byte
		pos     uint32
		tlvLen  uint32
		toWrite string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeString(tt.args.data, tt.args.pos, tt.args.tlvLen, tt.args.toWrite); got != tt.want {
				t.Errorf("writeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
