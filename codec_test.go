package checkup

import (
	"encoding/json"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "correct",
			args: args{&InternalCmd{
				Uri: "abc",
				Raw: func() []byte {
					b, err := json.Marshal(&internal_command.CmdPing{Timestamp: 1})
					if err != nil {
						return nil
					}
					return b
				}(),
				PassThrough: "123",
			}},
			want:    []byte{205, 123, 34, 117, 114, 105, 34, 58, 34, 97, 98, 99, 34, 44, 34, 114, 97, 119, 34, 58, 34, 101, 121, 74, 48, 97, 87, 49, 108, 99, 51, 82, 104, 98, 88, 65, 105, 79, 106, 70, 57, 34, 44, 34, 112, 97, 115, 115, 84, 104, 114, 111, 117, 103, 104, 34, 58, 34, 49, 50, 51, 34, 125},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
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

func TestUnmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []byte{205, 123, 34, 117, 114, 105, 34, 58, 34, 97, 98, 99, 34, 44, 34, 114, 97, 119, 34, 58, 34, 101, 121, 74, 48, 97, 87, 49, 108, 99, 51, 82, 104, 98, 88, 65, 105, 79, 106, 70, 57, 34, 44, 34, 112, 97, 115, 115, 84, 104, 114, 111, 117, 103, 104, 34, 58, 34, 49, 50, 51, 34, 125},
				v:    &InternalCmd{},
			},
			want: &InternalCmd{
				Uri: "abc",
				Raw: func() []byte {
					b, err := json.Marshal(&internal_command.CmdPing{Timestamp: 1})
					if err != nil {
						return nil
					}
					return b
				}(),
				PassThrough: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.args.v, tt.want) {
				t.Errorf("Unmarshal() got = %v, want %v", tt.args.v, tt.want)
			}
		})
	}
}
