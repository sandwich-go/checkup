package checkup

import (
	"context"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"
	"testing"
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
		opts = append(opts, WithDevopsCheckup(func(ctx context.Context) *internal_command.CmdCheckup {
			cmd.Code++
			return cmd
		}))

		d := New(opts...)
		bs = d.RequestBytes()
		So(len(bs), ShouldNotBeZeroValue)

		var valid Is
		So(d.IsRequestPath("aaaaa"), ShouldBeFalse)
		So(d.IsRequestPath(packetCheckup.Uri), ShouldBeTrue)
		_, valid = d.HandleIfRequestBytes(context.Background(), []byte{'1', '2'})
		So(err, ShouldBeNil)
		So(valid, ShouldBeFalse)

		bs, valid = d.HandleIfRequestBytes(context.Background(), d.RequestBytes())
		So(err, ShouldBeNil)
		So(valid, ShouldBeTrue)
		So(len(bs), ShouldNotBeZeroValue)

		var msg proto.Message
		msg, err = d.HandleResponseBytes(bs)
		So(err, ShouldBeNil)
		So(msg, ShouldNotBeNil)
		cmdRes, ok := msg.(*internal_command.CmdCheckup)
		So(ok, ShouldBeTrue)
		So(cmdRes.Code, ShouldEqual, 1)
		So(cmdRes.Message, ShouldEqual, cmd.Message)
		So(cmdRes.CustomMeasurements, ShouldNotBeEmpty)
	})
}
