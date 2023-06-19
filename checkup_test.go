package checkup

import (
	"context"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"

	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
)

func TestCheckup(t *testing.T) {
	Convey("checkup", t, func() {
		packetCheckup := &Packet{Uri: uri}
		codec := jsonCodec{}
		bs, err := codec.Marshal(packetCheckup)
		So(err, ShouldBeNil)
		var p = &Packet{}
		err = codec.Unmarshal(bs, p)
		So(err, ShouldBeNil)
		So(p.Uri, ShouldNotBeZeroValue)
		So(packetCheckup.Uri, ShouldEqual, p.Uri)

		var cmd = &internal_command.CmdCheckup{Code: 0, Message: "this is checkup message"}
		var opts []Option
		exed := false
		opts = append(opts, WithDevopsCheckup(func(ctx context.Context) *internal_command.CmdCheckup {
			exed = true
			return cmd
		}))

		d := New(opts...)
		bs = d.RequestBytes()
		So(len(bs), ShouldNotBeZeroValue)

		So(d.IsRequestPath("aaaaa"), ShouldBeFalse)
		So(d.IsRequestPath(packetCheckup.Uri), ShouldBeTrue)
		result := d.HandleIfRequestBytes(context.Background(), []byte{'1', '2'})
		So(result.RespBytes(), ShouldBeNil)
		So(result.Is(), ShouldBeFalse)
		So(result.Success(), ShouldBeTrue)
		So(exed, ShouldBeFalse)

		exed = false
		result = d.HandleIfRequestBytes(context.Background(), d.RequestBytes())
		So(result.RespBytes(), ShouldNotBeNil)
		So(result.Is(), ShouldBeTrue)
		So(result.Success(), ShouldBeTrue)
		So(len(result.RespBytes()), ShouldNotBeZeroValue)

		var msg proto.Message
		bs = result.RespBytes()
		msg, err = d.HandleResponseBytes(bs)
		So(err, ShouldBeNil)
		So(msg, ShouldNotBeNil)
		cmdRes, ok := msg.(*internal_command.CmdCheckup)
		So(ok, ShouldBeTrue)
		So(exed, ShouldBeTrue)
		So(cmdRes.Code, ShouldEqual, 0)
		So(cmdRes.Message, ShouldEqual, cmd.Message)
		So(cmdRes.CustomMeasurements, ShouldNotBeEmpty)
	})
}
