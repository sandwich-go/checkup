// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: internal_command/internal_cmd.proto

package internal_command

import (
	netutils "bitbucket.org/funplus/sandwich/protocol/netutils"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_internal_command_internal_cmd_proto protoreflect.FileDescriptor

var file_internal_command_internal_cmd_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6d, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x1a,
	0x13, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x63, 0x6d, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd, 0x01, 0x0a,
	0x0b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6d, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x43, 0x6d, 0x64, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69,
	0x6c, 0x73, 0x2e, 0x43, 0x6d, 0x64, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x2e, 0x6e, 0x65, 0x74,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6d, 0x64, 0x50, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x75, 0x70,
	0x12, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6d, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x75, 0x70, 0x1a, 0x14, 0x2e, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c,
	0x73, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x75, 0x70, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0e, 0x43, 0x6d, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x43, 0x6d, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x6e, 0x65, 0x74, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2e, 0x43, 0x6d, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38,
	0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x75,
	0x6e, 0x70, 0x6c, 0x75, 0x73, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x63, 0x68, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_command_internal_cmd_proto_goTypes = []interface{}{
	(*netutils.CmdPing)(nil),    // 0: netutils.CmdPing
	(*netutils.CmdCheckup)(nil), // 1: netutils.CmdCheckup
	(*CmdStream)(nil),           // 2: netutils.CmdStream
	(*netutils.CmdPingAck)(nil), // 3: netutils.CmdPingAck
}
var file_internal_command_internal_cmd_proto_depIdxs = []int32{
	0, // 0: netutils.InternalCmd.CmdPing:input_type -> netutils.CmdPing
	1, // 1: netutils.InternalCmd.CmdCheckup:input_type -> netutils.CmdCheckup
	2, // 2: netutils.InternalCmd.CmdBuildStream:input_type -> netutils.CmdStream
	3, // 3: netutils.InternalCmd.CmdPing:output_type -> netutils.CmdPingAck
	1, // 4: netutils.InternalCmd.CmdCheckup:output_type -> netutils.CmdCheckup
	2, // 5: netutils.InternalCmd.CmdBuildStream:output_type -> netutils.CmdStream
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_command_internal_cmd_proto_init() }
func file_internal_command_internal_cmd_proto_init() {
	if File_internal_command_internal_cmd_proto != nil {
		return
	}
	file_internal_command_internal_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_command_internal_cmd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_command_internal_cmd_proto_goTypes,
		DependencyIndexes: file_internal_command_internal_cmd_proto_depIdxs,
	}.Build()
	File_internal_command_internal_cmd_proto = out.File
	file_internal_command_internal_cmd_proto_rawDesc = nil
	file_internal_command_internal_cmd_proto_goTypes = nil
	file_internal_command_internal_cmd_proto_depIdxs = nil
}
