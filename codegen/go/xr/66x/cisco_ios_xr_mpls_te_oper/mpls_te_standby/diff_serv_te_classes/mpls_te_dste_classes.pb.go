/*
Copyright 2019 Cisco Systems

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mpls_te_dste_classes.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_diff_serv_te_classes

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

type MplsTeDsteClasses_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeDsteClasses_KEYS) Reset()         { *m = MplsTeDsteClasses_KEYS{} }
func (m *MplsTeDsteClasses_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeDsteClasses_KEYS) ProtoMessage()    {}
func (*MplsTeDsteClasses_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_54b768e82f778f9f, []int{0}
}

func (m *MplsTeDsteClasses_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeDsteClasses_KEYS.Unmarshal(m, b)
}
func (m *MplsTeDsteClasses_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeDsteClasses_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeDsteClasses_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeDsteClasses_KEYS.Merge(m, src)
}
func (m *MplsTeDsteClasses_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeDsteClasses_KEYS.Size(m)
}
func (m *MplsTeDsteClasses_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeDsteClasses_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeDsteClasses_KEYS proto.InternalMessageInfo

type MplsTeDsteClass struct {
	ClassNumber          uint32   `protobuf:"varint,1,opt,name=class_number,json=classNumber,proto3" json:"class_number,omitempty"`
	ClassType            uint32   `protobuf:"varint,2,opt,name=class_type,json=classType,proto3" json:"class_type,omitempty"`
	Priority             uint32   `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	ClassStatus          string   `protobuf:"bytes,4,opt,name=class_status,json=classStatus,proto3" json:"class_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeDsteClass) Reset()         { *m = MplsTeDsteClass{} }
func (m *MplsTeDsteClass) String() string { return proto.CompactTextString(m) }
func (*MplsTeDsteClass) ProtoMessage()    {}
func (*MplsTeDsteClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_54b768e82f778f9f, []int{1}
}

func (m *MplsTeDsteClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeDsteClass.Unmarshal(m, b)
}
func (m *MplsTeDsteClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeDsteClass.Marshal(b, m, deterministic)
}
func (m *MplsTeDsteClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeDsteClass.Merge(m, src)
}
func (m *MplsTeDsteClass) XXX_Size() int {
	return xxx_messageInfo_MplsTeDsteClass.Size(m)
}
func (m *MplsTeDsteClass) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeDsteClass.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeDsteClass proto.InternalMessageInfo

func (m *MplsTeDsteClass) GetClassNumber() uint32 {
	if m != nil {
		return m.ClassNumber
	}
	return 0
}

func (m *MplsTeDsteClass) GetClassType() uint32 {
	if m != nil {
		return m.ClassType
	}
	return 0
}

func (m *MplsTeDsteClass) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *MplsTeDsteClass) GetClassStatus() string {
	if m != nil {
		return m.ClassStatus
	}
	return ""
}

type MplsTeDsteClasses struct {
	TeClass              []*MplsTeDsteClass `protobuf:"bytes,50,rep,name=te_class,json=teClass,proto3" json:"te_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MplsTeDsteClasses) Reset()         { *m = MplsTeDsteClasses{} }
func (m *MplsTeDsteClasses) String() string { return proto.CompactTextString(m) }
func (*MplsTeDsteClasses) ProtoMessage()    {}
func (*MplsTeDsteClasses) Descriptor() ([]byte, []int) {
	return fileDescriptor_54b768e82f778f9f, []int{2}
}

func (m *MplsTeDsteClasses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeDsteClasses.Unmarshal(m, b)
}
func (m *MplsTeDsteClasses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeDsteClasses.Marshal(b, m, deterministic)
}
func (m *MplsTeDsteClasses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeDsteClasses.Merge(m, src)
}
func (m *MplsTeDsteClasses) XXX_Size() int {
	return xxx_messageInfo_MplsTeDsteClasses.Size(m)
}
func (m *MplsTeDsteClasses) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeDsteClasses.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeDsteClasses proto.InternalMessageInfo

func (m *MplsTeDsteClasses) GetTeClass() []*MplsTeDsteClass {
	if m != nil {
		return m.TeClass
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeDsteClasses_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.diff_serv_te_classes.mpls_te_dste_classes_KEYS")
	proto.RegisterType((*MplsTeDsteClass)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.diff_serv_te_classes.mpls_te_dste_class")
	proto.RegisterType((*MplsTeDsteClasses)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.diff_serv_te_classes.mpls_te_dste_classes")
}

func init() { proto.RegisterFile("mpls_te_dste_classes.proto", fileDescriptor_54b768e82f778f9f) }

var fileDescriptor_54b768e82f778f9f = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x65, 0x8a, 0xa0, 0xbd, 0xc2, 0x62, 0x31, 0x98, 0x22, 0xa4, 0x90, 0x29, 0x93, 0x87,
	0xb2, 0xb3, 0x20, 0x26, 0x24, 0x86, 0x94, 0x85, 0xe9, 0x94, 0x3f, 0x57, 0xc9, 0x52, 0x53, 0x5b,
	0xbe, 0x2b, 0x22, 0x0f, 0xc0, 0x23, 0xf0, 0xbe, 0x08, 0x43, 0xe8, 0xd0, 0x6c, 0xdd, 0xee, 0xfb,
	0xdd, 0xa7, 0x9f, 0x2d, 0x1b, 0x16, 0x5d, 0xd8, 0x30, 0x0a, 0x61, 0xcb, 0x42, 0xd8, 0x6c, 0x2a,
	0x66, 0x62, 0x1b, 0xa2, 0x17, 0xaf, 0x1f, 0x1a, 0xc7, 0x8d, 0x47, 0xe7, 0x19, 0x3f, 0x22, 0x0e,
	0x45, 0x1f, 0x28, 0xda, 0x21, 0xb0, 0x54, 0xdb, 0xb6, 0xee, 0x6d, 0xeb, 0xd6, 0x6b, 0x64, 0x8a,
	0xef, 0xb8, 0xb7, 0xe4, 0x37, 0x70, 0x3d, 0x66, 0xc7, 0xe7, 0xa7, 0xb7, 0x55, 0xfe, 0xa5, 0x40,
	0x1f, 0x6e, 0xf5, 0x1d, 0x5c, 0xa4, 0x01, 0xb7, 0xbb, 0xae, 0xa6, 0x68, 0x54, 0xa6, 0x8a, 0xcb,
	0x72, 0x9e, 0xd8, 0x4b, 0x42, 0xfa, 0x16, 0xe0, 0xb7, 0x22, 0x7d, 0x20, 0x73, 0x92, 0x0a, 0xb3,
	0x44, 0x5e, 0xfb, 0x40, 0x7a, 0x01, 0xd3, 0x10, 0x9d, 0x8f, 0x4e, 0x7a, 0x33, 0x49, 0xcb, 0xff,
	0xbc, 0xb7, 0xb3, 0x54, 0xb2, 0x63, 0x73, 0x9a, 0xa9, 0x62, 0xf6, 0x67, 0x5f, 0x25, 0x94, 0x7f,
	0x2a, 0xb8, 0x1a, 0xbb, 0xb5, 0xee, 0x60, 0x3a, 0x24, 0xb3, 0xcc, 0x26, 0xc5, 0x7c, 0x59, 0xda,
	0xe3, 0x1e, 0xc8, 0x1e, 0x9e, 0x53, 0x9e, 0x0b, 0x3d, 0xfe, 0x0c, 0xf5, 0x59, 0xfa, 0x83, 0xfb,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x10, 0xb9, 0xab, 0xa1, 0x01, 0x00, 0x00,
}
