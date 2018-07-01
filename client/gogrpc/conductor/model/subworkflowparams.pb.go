// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/subworkflowparams.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubWorkflowParams struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version              *_struct.Value `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SubWorkflowParams) Reset()         { *m = SubWorkflowParams{} }
func (m *SubWorkflowParams) String() string { return proto.CompactTextString(m) }
func (*SubWorkflowParams) ProtoMessage()    {}
func (*SubWorkflowParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_subworkflowparams_c15f17d5a5a4cfdd, []int{0}
}
func (m *SubWorkflowParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubWorkflowParams.Unmarshal(m, b)
}
func (m *SubWorkflowParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubWorkflowParams.Marshal(b, m, deterministic)
}
func (dst *SubWorkflowParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubWorkflowParams.Merge(dst, src)
}
func (m *SubWorkflowParams) XXX_Size() int {
	return xxx_messageInfo_SubWorkflowParams.Size(m)
}
func (m *SubWorkflowParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SubWorkflowParams.DiscardUnknown(m)
}

var xxx_messageInfo_SubWorkflowParams proto.InternalMessageInfo

func (m *SubWorkflowParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubWorkflowParams) GetVersion() *_struct.Value {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*SubWorkflowParams)(nil), "com.netflix.conductor.proto.SubWorkflowParams")
}

func init() {
	proto.RegisterFile("model/subworkflowparams.proto", fileDescriptor_subworkflowparams_c15f17d5a5a4cfdd)
}

var fileDescriptor_subworkflowparams_c15f17d5a5a4cfdd = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x88, 0x62, 0x3c, 0x19, 0x41, 0x8a, 0x1f, 0x50, 0x3c, 0xf5, 0x94, 0x88, 0xde,
	0x3c, 0xf6, 0x17, 0xd4, 0x0a, 0x8a, 0xde, 0x92, 0x34, 0xcd, 0x86, 0x4d, 0x32, 0x25, 0x1f, 0xdb,
	0xfd, 0xf9, 0x0b, 0x69, 0xbb, 0x2c, 0xec, 0x6d, 0x66, 0xde, 0xe1, 0x99, 0x87, 0x41, 0x2f, 0x16,
	0x7a, 0x69, 0x68, 0x48, 0x7c, 0x02, 0xbf, 0x1d, 0x0c, 0x4c, 0x23, 0xf3, 0xcc, 0x06, 0x32, 0x7a,
	0x88, 0x80, 0x9f, 0x04, 0x58, 0xe2, 0x64, 0x1c, 0x8c, 0xde, 0x13, 0x01, 0xae, 0x4f, 0x22, 0x82,
	0x9f, 0xc3, 0xc7, 0x67, 0x05, 0xa0, 0x8c, 0xa4, 0xb9, 0xe3, 0x69, 0xa0, 0x21, 0xfa, 0x24, 0xe2,
	0x9c, 0xbe, 0xfe, 0xa1, 0xbb, 0xef, 0xc4, 0x7f, 0x17, 0x6a, 0x9b, 0xa9, 0x18, 0xa3, 0x4b, 0xc7,
	0xac, 0x2c, 0x8b, 0xaa, 0xa8, 0x6f, 0xba, 0x5c, 0xe3, 0x37, 0x74, 0xbd, 0x93, 0x3e, 0x68, 0x70,
	0xe5, 0x45, 0x55, 0xd4, 0xb7, 0xef, 0x0f, 0x64, 0x06, 0x93, 0x15, 0x4c, 0x7e, 0x98, 0x49, 0xb2,
	0x5b, 0xd7, 0x9a, 0xaf, 0xe6, 0xfe, 0x0c, 0xdd, 0xf2, 0xff, 0x4f, 0xa5, 0xe3, 0x26, 0x71, 0x22,
	0xc0, 0xd2, 0xc5, 0x99, 0x1e, 0x9d, 0xa9, 0x30, 0x5a, 0xba, 0x48, 0x15, 0x28, 0x3f, 0x8a, 0x93,
	0x79, 0x7e, 0x00, 0xbf, 0xca, 0xb7, 0x3e, 0x0e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x5e, 0xa1,
	0x60, 0x10, 0x01, 0x00, 0x00,
}
