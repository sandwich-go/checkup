// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/internal_command/internal.proto

package internal_command

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//CSHARP_NAMESPACE
type CmdStream struct {
	Addr                 string            `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Token                string            `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,3,rep,name=Meta,proto3" json:"Meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CmdStream) Reset()         { *m = CmdStream{} }
func (m *CmdStream) String() string { return proto.CompactTextString(m) }
func (*CmdStream) ProtoMessage()    {}
func (*CmdStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa374f2c537ab1b, []int{0}
}

func (m *CmdStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdStream.Unmarshal(m, b)
}
func (m *CmdStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdStream.Marshal(b, m, deterministic)
}
func (m *CmdStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdStream.Merge(m, src)
}
func (m *CmdStream) XXX_Size() int {
	return xxx_messageInfo_CmdStream.Size(m)
}
func (m *CmdStream) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdStream.DiscardUnknown(m)
}

var xxx_messageInfo_CmdStream proto.InternalMessageInfo

func (m *CmdStream) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *CmdStream) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CmdStream) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

// for internal use
type CmdPing struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdPing) Reset()         { *m = CmdPing{} }
func (m *CmdPing) String() string { return proto.CompactTextString(m) }
func (*CmdPing) ProtoMessage()    {}
func (*CmdPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa374f2c537ab1b, []int{1}
}

func (m *CmdPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdPing.Unmarshal(m, b)
}
func (m *CmdPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdPing.Marshal(b, m, deterministic)
}
func (m *CmdPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdPing.Merge(m, src)
}
func (m *CmdPing) XXX_Size() int {
	return xxx_messageInfo_CmdPing.Size(m)
}
func (m *CmdPing) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdPing.DiscardUnknown(m)
}

var xxx_messageInfo_CmdPing proto.InternalMessageInfo

func (m *CmdPing) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CmdPingAck struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdPingAck) Reset()         { *m = CmdPingAck{} }
func (m *CmdPingAck) String() string { return proto.CompactTextString(m) }
func (*CmdPingAck) ProtoMessage()    {}
func (*CmdPingAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa374f2c537ab1b, []int{2}
}

func (m *CmdPingAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdPingAck.Unmarshal(m, b)
}
func (m *CmdPingAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdPingAck.Marshal(b, m, deterministic)
}
func (m *CmdPingAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdPingAck.Merge(m, src)
}
func (m *CmdPingAck) XXX_Size() int {
	return xxx_messageInfo_CmdPingAck.Size(m)
}
func (m *CmdPingAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdPingAck.DiscardUnknown(m)
}

var xxx_messageInfo_CmdPingAck proto.InternalMessageInfo

func (m *CmdPingAck) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// for devops checkup
type CmdCheckup struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	CustomMeasurements   []byte   `protobuf:"bytes,3,opt,name=CustomMeasurements,proto3" json:"CustomMeasurements,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdCheckup) Reset()         { *m = CmdCheckup{} }
func (m *CmdCheckup) String() string { return proto.CompactTextString(m) }
func (*CmdCheckup) ProtoMessage()    {}
func (*CmdCheckup) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa374f2c537ab1b, []int{3}
}

func (m *CmdCheckup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdCheckup.Unmarshal(m, b)
}
func (m *CmdCheckup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdCheckup.Marshal(b, m, deterministic)
}
func (m *CmdCheckup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdCheckup.Merge(m, src)
}
func (m *CmdCheckup) XXX_Size() int {
	return xxx_messageInfo_CmdCheckup.Size(m)
}
func (m *CmdCheckup) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdCheckup.DiscardUnknown(m)
}

var xxx_messageInfo_CmdCheckup proto.InternalMessageInfo

func (m *CmdCheckup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CmdCheckup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CmdCheckup) GetCustomMeasurements() []byte {
	if m != nil {
		return m.CustomMeasurements
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdStream)(nil), "internal_command.cmdStream")
	proto.RegisterMapType((map[string]string)(nil), "internal_command.cmdStream.MetaEntry")
	proto.RegisterType((*CmdPing)(nil), "internal_command.CmdPing")
	proto.RegisterType((*CmdPingAck)(nil), "internal_command.CmdPingAck")
	proto.RegisterType((*CmdCheckup)(nil), "internal_command.CmdCheckup")
}

func init() {
	proto.RegisterFile("protos/internal_command/internal.proto", fileDescriptor_0fa374f2c537ab1b)
}

var fileDescriptor_0fa374f2c537ab1b = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xd7, 0x75, 0x7b, 0x47, 0xb3, 0x17, 0x19, 0xc1, 0x8b, 0x5a, 0x77, 0x31, 0x0a, 0xea,
	0x10, 0xec, 0x60, 0x5e, 0xf8, 0x07, 0x11, 0xb6, 0xea, 0x85, 0xe2, 0x40, 0xaa, 0x57, 0xde, 0x48,
	0x96, 0x84, 0xae, 0xb6, 0x49, 0x46, 0x93, 0x2a, 0xfb, 0x48, 0x7e, 0x38, 0xbf, 0x83, 0x34, 0xb6,
	0x1d, 0xce, 0x39, 0xef, 0xce, 0x79, 0x9e, 0x5f, 0x72, 0x9a, 0xd3, 0x07, 0xec, 0xcf, 0x53, 0xa1,
	0x84, 0x1c, 0x44, 0x5c, 0xd1, 0x94, 0xa3, 0xe4, 0x19, 0x0b, 0xc6, 0x10, 0x27, 0x95, 0xe0, 0x69,
	0x00, 0x76, 0x56, 0x01, 0xf7, 0xdd, 0x00, 0x96, 0xcf, 0xc8, 0x83, 0x4a, 0x29, 0x62, 0x10, 0x82,
	0xc6, 0x88, 0x90, 0xd4, 0x36, 0x7a, 0x46, 0xdf, 0x0a, 0x74, 0x0d, 0xb7, 0x41, 0xf3, 0x51, 0xc4,
	0x94, 0xdb, 0x75, 0x2d, 0x7e, 0x35, 0xf0, 0x0c, 0x34, 0x26, 0x54, 0x21, 0xdb, 0xec, 0x99, 0xfd,
	0xf6, 0x70, 0xcf, 0x5b, 0xbd, 0xd8, 0xab, 0x2e, 0xf5, 0x72, 0xee, 0x9a, 0xab, 0x74, 0x11, 0xe8,
	0x23, 0xce, 0x09, 0xb0, 0x2a, 0x09, 0x76, 0x80, 0x19, 0xd3, 0x45, 0x31, 0x30, 0x2f, 0xf3, 0x79,
	0xaf, 0x28, 0xc9, 0x68, 0x39, 0x4f, 0x37, 0xe7, 0xf5, 0x53, 0xc3, 0x3d, 0x00, 0x2d, 0x9f, 0x91,
	0xfb, 0x88, 0x87, 0xb0, 0x0b, 0x2c, 0x15, 0x31, 0x2a, 0x15, 0x62, 0x73, 0x7d, 0xd8, 0x0c, 0x96,
	0x82, 0x7b, 0x08, 0x40, 0x01, 0x8e, 0x70, 0xfc, 0x07, 0xfb, 0xa2, 0x59, 0x7f, 0x46, 0x71, 0x9c,
	0xcd, 0xf3, 0x05, 0x60, 0x41, 0xa8, 0xc6, 0x9a, 0x81, 0xae, 0xa1, 0x0d, 0x5a, 0x8c, 0x4a, 0x89,
	0xc2, 0xf2, 0x93, 0xca, 0x16, 0x7a, 0x00, 0xfa, 0x99, 0x54, 0x82, 0x4d, 0x28, 0x92, 0x59, 0x4a,
	0x19, 0xe5, 0x4a, 0xda, 0x66, 0xcf, 0xe8, 0xff, 0x0f, 0xd6, 0x38, 0xc3, 0x0f, 0x03, 0xb4, 0x6f,
	0x8a, 0x45, 0xf9, 0x8c, 0xc0, 0xab, 0xe5, 0x83, 0x76, 0xd6, 0x6e, 0x30, 0xb7, 0x9c, 0xee, 0xaf,
	0xd6, 0x08, 0xc7, 0x6e, 0x0d, 0xde, 0x7e, 0x7b, 0xc1, 0x7a, 0xba, 0x70, 0x9d, 0x8d, 0xae, 0x5b,
	0x83, 0x77, 0x60, 0xcb, 0x67, 0x64, 0x9c, 0x45, 0x49, 0x19, 0x89, 0xdd, 0x0d, 0xbf, 0xd6, 0xd9,
	0x64, 0xba, 0xb5, 0xf1, 0xe5, 0xd3, 0x45, 0x18, 0xa9, 0x59, 0x36, 0xf5, 0xb0, 0x60, 0x03, 0x89,
	0x38, 0x79, 0x8b, 0xf0, 0xec, 0x28, 0x14, 0x55, 0x2e, 0x31, 0x23, 0x03, 0x1d, 0x4d, 0x2c, 0x92,
	0x1f, 0xe9, 0x9d, 0xfe, 0xd3, 0xd6, 0xf1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x1e, 0x40,
	0xdb, 0xdf, 0x02, 0x00, 0x00,
}
