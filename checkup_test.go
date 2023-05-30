package checkup

import (
	"context"
	"github.com/sandwich-go/boost/singleflight"
	"github.com/sandwich-go/checkup/protocol/gen/golang/internal_command"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestCheckup(t *testing.T) {
	Convey("checkup", t, func() {
		codec := jsonCodec{}
		bs, err := codec.Marshal(PacketCheckup)
		So(err, ShouldBeNil)
		var p = &Packet{}
		err = codec.Unmarshal(bs, p)
		So(err, ShouldBeNil)
		So(p.Uri, ShouldNotBeZeroValue)
		So(PacketCheckup.Uri, ShouldEqual, p.Uri)

		processor := &checkupPacketProcessor{}
		So(processor.packet(), ShouldNotBeNil)
		So(PacketCheckup.Uri, ShouldEqual, processor.packet().Uri)
		So(processor.packet().Raw, ShouldBeNil)

		var cmd = &internal_command.CmdCheckup{Code: 0, Message: "this is checkup message"}
		var opts []Option
		opts = append(opts, WithDevopsCheckup(func(ctx context.Context) *internal_command.CmdCheckup {
			cmd.Code++
			return cmd
		}))
		cc := NewOptions(opts...)

		bs, err = processor.process(context.Background(), &singleflight.Group{}, cc)
		So(err, ShouldBeNil)

		bs, err = processor.process(context.Background(), &singleflight.Group{}, cc)
		So(err, ShouldBeNil)

		d := New(opts...)
		bs = d.Bytes(PacketCheckup.Uri)
		So(len(bs), ShouldNotBeZeroValue)
		bs = d.Bytes("invalid")
		So(len(bs), ShouldBeZeroValue)

		So(d.Check(d.Bytes(PacketCheckup.Uri)), ShouldBeTrue)

		var valid Valid
		_, valid, err = d.Handle(context.Background(), []byte{'1', '2'})
		So(err, ShouldBeNil)
		So(valid, ShouldBeFalse)

		bs, valid, err = d.Handle(context.Background(), d.Bytes(PacketCheckup.Uri))
		So(err, ShouldBeNil)
		So(valid, ShouldBeTrue)
		So(len(bs), ShouldNotBeZeroValue)

		var uri URI
		var msg proto.Message
		uri, msg, err = d.Decode(bs)
		So(err, ShouldBeNil)
		So(uri, ShouldNotBeZeroValue)
		So(PacketCheckup.Uri, ShouldEqual, uri)
		So(msg, ShouldNotBeNil)
		cmdRes, ok := msg.(*internal_command.CmdCheckup)
		So(ok, ShouldBeTrue)
		So(cmdRes.Code, ShouldEqual, 3)
		So(cmdRes.Message, ShouldEqual, cmd.Message)
		So(cmdRes.CustomMeasurements, ShouldNotBeEmpty)
	})
}
