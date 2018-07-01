// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflowdef.proto

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

type WorkflowDef struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description          string                    `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Version              int32                     `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Tasks                []*WorkflowTask           `protobuf:"bytes,4,rep,name=tasks" json:"tasks,omitempty"`
	InputParameters      []string                  `protobuf:"bytes,5,rep,name=input_parameters,json=inputParameters" json:"input_parameters,omitempty"`
	OutputParameters     map[string]*_struct.Value `protobuf:"bytes,6,rep,name=output_parameters,json=outputParameters" json:"output_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	FailureWorkflow      string                    `protobuf:"bytes,7,opt,name=failure_workflow,json=failureWorkflow" json:"failure_workflow,omitempty"`
	SchemaVersion        int32                     `protobuf:"varint,8,opt,name=schema_version,json=schemaVersion" json:"schema_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *WorkflowDef) Reset()         { *m = WorkflowDef{} }
func (m *WorkflowDef) String() string { return proto.CompactTextString(m) }
func (*WorkflowDef) ProtoMessage()    {}
func (*WorkflowDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflowdef_bf5b9fbf7e32cdd5, []int{0}
}
func (m *WorkflowDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowDef.Unmarshal(m, b)
}
func (m *WorkflowDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowDef.Marshal(b, m, deterministic)
}
func (dst *WorkflowDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowDef.Merge(dst, src)
}
func (m *WorkflowDef) XXX_Size() int {
	return xxx_messageInfo_WorkflowDef.Size(m)
}
func (m *WorkflowDef) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowDef.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowDef proto.InternalMessageInfo

func (m *WorkflowDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WorkflowDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *WorkflowDef) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *WorkflowDef) GetTasks() []*WorkflowTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *WorkflowDef) GetInputParameters() []string {
	if m != nil {
		return m.InputParameters
	}
	return nil
}

func (m *WorkflowDef) GetOutputParameters() map[string]*_struct.Value {
	if m != nil {
		return m.OutputParameters
	}
	return nil
}

func (m *WorkflowDef) GetFailureWorkflow() string {
	if m != nil {
		return m.FailureWorkflow
	}
	return ""
}

func (m *WorkflowDef) GetSchemaVersion() int32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*WorkflowDef)(nil), "com.netflix.conductor.proto.WorkflowDef")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "com.netflix.conductor.proto.WorkflowDef.OutputParametersEntry")
}

func init() {
	proto.RegisterFile("model/workflowdef.proto", fileDescriptor_workflowdef_bf5b9fbf7e32cdd5)
}

var fileDescriptor_workflowdef_bf5b9fbf7e32cdd5 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4b, 0xab, 0xd3, 0x40,
	0x14, 0xc7, 0xc9, 0xcd, 0xcd, 0xad, 0x9d, 0x50, 0x5b, 0x07, 0xd4, 0x50, 0x5d, 0x04, 0x41, 0x48,
	0x41, 0x26, 0x50, 0x37, 0xd2, 0x85, 0x42, 0xa9, 0x2b, 0x17, 0x96, 0x20, 0x15, 0x74, 0x51, 0x26,
	0x93, 0x93, 0x34, 0xe4, 0x31, 0x61, 0x1e, 0xad, 0xfd, 0xc0, 0x7e, 0x0f, 0xc9, 0x24, 0xa9, 0xb5,
	0x88, 0xb8, 0x9b, 0xf9, 0xfd, 0xe7, 0xbc, 0xfe, 0x67, 0xd0, 0xf3, 0x8a, 0x27, 0x50, 0x86, 0x27,
	0x2e, 0x8a, 0xb4, 0xe4, 0xa7, 0x04, 0x52, 0xd2, 0x08, 0xae, 0x38, 0x7e, 0xc1, 0x78, 0x45, 0x6a,
	0x50, 0x69, 0x99, 0xff, 0x20, 0x8c, 0xd7, 0x89, 0x66, 0x8a, 0x8b, 0x4e, 0x9c, 0x7b, 0x7f, 0x46,
	0x29, 0x2a, 0x8b, 0x5e, 0x79, 0x99, 0x71, 0x9e, 0x95, 0x10, 0x9a, 0x5b, 0xac, 0xd3, 0x50, 0x2a,
	0xa1, 0x99, 0xea, 0xd4, 0x57, 0x3f, 0x6d, 0xe4, 0x7e, 0xed, 0x83, 0x36, 0x90, 0x62, 0x8c, 0xee,
	0x6b, 0x5a, 0x81, 0x67, 0xf9, 0x56, 0x30, 0x8e, 0xcc, 0x19, 0xfb, 0xc8, 0x4d, 0x40, 0x32, 0x91,
	0x37, 0x2a, 0xe7, 0xb5, 0x77, 0x67, 0xa4, 0x6b, 0x84, 0x3d, 0x34, 0x3a, 0x82, 0x90, 0xad, 0x6a,
	0xfb, 0x56, 0xe0, 0x44, 0xc3, 0x15, 0x7f, 0x40, 0x4e, 0xdb, 0x8b, 0xf4, 0xee, 0x7d, 0x3b, 0x70,
	0x97, 0x0b, 0xf2, 0x8f, 0x21, 0xc8, 0xd0, 0xc8, 0x17, 0x2a, 0x8b, 0xa8, 0x8b, 0xc3, 0x0b, 0x34,
	0xcb, 0xeb, 0x46, 0xab, 0x7d, 0x43, 0x05, 0xad, 0x40, 0x81, 0x90, 0x9e, 0xe3, 0xdb, 0xc1, 0x38,
	0x9a, 0x1a, 0xbe, 0xbd, 0x60, 0x5c, 0xa0, 0x27, 0x5c, 0xab, 0x9b, 0xb7, 0x0f, 0xa6, 0xee, 0xfb,
	0xff, 0xaa, 0xbb, 0x81, 0x94, 0x7c, 0x36, 0x19, 0x7e, 0x67, 0xfd, 0x58, 0x2b, 0x71, 0x8e, 0x66,
	0xfc, 0x06, 0xb7, 0x7d, 0xa5, 0x34, 0x2f, 0xb5, 0x80, 0xfd, 0x60, 0xba, 0x37, 0x32, 0xce, 0x4c,
	0x7b, 0x3e, 0x64, 0xc5, 0xaf, 0xd1, 0x63, 0xc9, 0x0e, 0x50, 0xd1, 0xfd, 0x60, 0xd2, 0x23, 0x63,
	0xd2, 0xa4, 0xa3, 0xbb, 0x0e, 0xce, 0xbf, 0xa3, 0xa7, 0x7f, 0x2d, 0x8e, 0x67, 0xc8, 0x2e, 0xe0,
	0xdc, 0xaf, 0xa4, 0x3d, 0xe2, 0x37, 0xc8, 0x39, 0xd2, 0x52, 0x83, 0xd9, 0x85, 0xbb, 0x7c, 0x46,
	0xba, 0x1d, 0x93, 0x61, 0xc7, 0x64, 0xd7, 0xaa, 0x51, 0xf7, 0x68, 0x75, 0xf7, 0xce, 0x5a, 0x7f,
	0x5a, 0x4f, 0xae, 0xa6, 0xdc, 0xc6, 0xdf, 0x56, 0x59, 0xae, 0x0e, 0x3a, 0x6e, 0x7d, 0x09, 0x7b,
	0x5f, 0xc2, 0x8b, 0x2f, 0x21, 0x2b, 0x73, 0xa8, 0x55, 0x98, 0xf1, 0x4c, 0x34, 0xec, 0x8a, 0x9b,
	0x5f, 0x16, 0x3f, 0x98, 0x3a, 0x6f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x9d, 0x18, 0xb5,
	0xab, 0x02, 0x00, 0x00,
}
