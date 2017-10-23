// Code generated by protoc-gen-go. DO NOT EDIT.
// source: version.proto

/*
Package cmd is a generated protocol buffer package.

It is generated from these files:
	version.proto

It has these top-level messages:
	FlagInfo
	VersionTest
*/
package cmd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FlagInfo struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Shorthand        *string `protobuf:"bytes,2,opt,name=shorthand" json:"shorthand,omitempty"`
	Usage            *string `protobuf:"bytes,3,opt,name=usage" json:"usage,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *FlagInfo) Reset()                    { *m = FlagInfo{} }
func (m *FlagInfo) String() string            { return proto.CompactTextString(m) }
func (*FlagInfo) ProtoMessage()               {}
func (*FlagInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FlagInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *FlagInfo) GetShorthand() string {
	if m != nil && m.Shorthand != nil {
		return *m.Shorthand
	}
	return ""
}

func (m *FlagInfo) GetUsage() string {
	if m != nil && m.Usage != nil {
		return *m.Usage
	}
	return ""
}

type VersionTest struct {
	Client           *bool  `protobuf:"varint,2,opt,name=client" json:"client,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *VersionTest) Reset()                    { *m = VersionTest{} }
func (m *VersionTest) String() string            { return proto.CompactTextString(m) }
func (*VersionTest) ProtoMessage()               {}
func (*VersionTest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VersionTest) GetClient() bool {
	if m != nil && m.Client != nil {
		return *m.Client
	}
	return false
}

var E_Info = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FlagInfo)(nil),
	Field:         1234,
	Name:          "cmd.info",
	Tag:           "bytes,1234,opt,name=info",
	//Filename:      "version.proto",
}

func init() {
	proto.RegisterType((*FlagInfo)(nil), "cmd.FlagInfo")
	proto.RegisterType((*VersionTest)(nil), "cmd.VersionTest")
	proto.RegisterExtension(E_Info)
}

func init() { proto.RegisterFile("version.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xa9, 0xd2, 0x6c, 0xe9, 0x65, 0xf1, 0x10, 0x8a, 0xc2, 0xd2, 0x53, 0x7a, 0xd9,
	0x82, 0x27, 0xd1, 0x93, 0x0a, 0x05, 0x0f, 0x22, 0x44, 0xf1, 0xbe, 0xc9, 0x4e, 0xd3, 0x85, 0xcd,
	0x4e, 0x9c, 0xdd, 0x14, 0xf4, 0x67, 0xf8, 0xd3, 0xfc, 0x45, 0xe2, 0x26, 0xa1, 0xb7, 0x79, 0xef,
	0xcd, 0x83, 0xf7, 0xb1, 0xe5, 0x11, 0xc8, 0x1b, 0x74, 0xb2, 0x23, 0x0c, 0xc8, 0xd3, 0xba, 0xd5,
	0x2b, 0xd1, 0x20, 0x36, 0x16, 0xb6, 0xd1, 0xaa, 0xfa, 0xfd, 0x56, 0x83, 0xaf, 0xc9, 0x74, 0x01,
	0x69, 0x78, 0x5b, 0x97, 0x6c, 0xbe, 0xb3, 0xaa, 0x79, 0x76, 0x7b, 0xe4, 0x9c, 0xcd, 0x9c, 0x6a,
	0x21, 0x4f, 0x44, 0x52, 0x64, 0x65, 0xbc, 0xf9, 0x15, 0xcb, 0xfc, 0x01, 0x29, 0x1c, 0x94, 0xd3,
	0xf9, 0x59, 0x0c, 0x4e, 0x06, 0xbf, 0x64, 0xe7, 0xbd, 0x57, 0x0d, 0xe4, 0x69, 0x4c, 0x06, 0xb1,
	0xae, 0xd8, 0xe2, 0x63, 0xd8, 0xf2, 0x0e, 0x3e, 0xf0, 0x37, 0x76, 0x51, 0x5b, 0x03, 0x2e, 0xc4,
	0xfe, 0xfc, 0xf1, 0xfe, 0xe7, 0xe5, 0x76, 0xb2, 0x78, 0x5a, 0xa9, 0xef, 0xd5, 0xe6, 0x29, 0x0a,
	0x31, 0x12, 0x08, 0x74, 0xf6, 0x4b, 0x14, 0x0e, 0x85, 0x07, 0x3a, 0x02, 0x09, 0x82, 0xcf, 0xde,
	0x10, 0xe8, 0x8d, 0x2c, 0xc7, 0xde, 0xdd, 0x03, 0x9b, 0x99, 0xff, 0xcd, 0xd7, 0x72, 0x40, 0x94,
	0x13, 0xa2, 0xdc, 0x19, 0xb0, 0xfa, 0xb5, 0x0b, 0x06, 0x9d, 0xcf, 0x7f, 0x33, 0x91, 0x14, 0x8b,
	0x9b, 0xa5, 0xac, 0x5b, 0x2d, 0x27, 0xd0, 0x32, 0x56, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba,
	0x5b, 0x09, 0x9a, 0x31, 0x01, 0x00, 0x00,
}
