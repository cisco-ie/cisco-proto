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
// source: remoteusr_det.proto

package cisco_ios_xr_aaa_locald_oper_aaa_authen_method

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

type RemoteusrDet_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteusrDet_KEYS) Reset()         { *m = RemoteusrDet_KEYS{} }
func (m *RemoteusrDet_KEYS) String() string { return proto.CompactTextString(m) }
func (*RemoteusrDet_KEYS) ProtoMessage()    {}
func (*RemoteusrDet_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cf4956e430e79ab, []int{0}
}

func (m *RemoteusrDet_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteusrDet_KEYS.Unmarshal(m, b)
}
func (m *RemoteusrDet_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteusrDet_KEYS.Marshal(b, m, deterministic)
}
func (m *RemoteusrDet_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteusrDet_KEYS.Merge(m, src)
}
func (m *RemoteusrDet_KEYS) XXX_Size() int {
	return xxx_messageInfo_RemoteusrDet_KEYS.Size(m)
}
func (m *RemoteusrDet_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteusrDet_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteusrDet_KEYS proto.InternalMessageInfo

type RemoteusrDet struct {
	Name                 string   `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Usergroup            []string `protobuf:"bytes,51,rep,name=usergroup,proto3" json:"usergroup,omitempty"`
	Authenmethod         int32    `protobuf:"zigzag32,52,opt,name=authenmethod,proto3" json:"authenmethod,omitempty"`
	Taskmap              []string `protobuf:"bytes,53,rep,name=taskmap,proto3" json:"taskmap,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteusrDet) Reset()         { *m = RemoteusrDet{} }
func (m *RemoteusrDet) String() string { return proto.CompactTextString(m) }
func (*RemoteusrDet) ProtoMessage()    {}
func (*RemoteusrDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cf4956e430e79ab, []int{1}
}

func (m *RemoteusrDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteusrDet.Unmarshal(m, b)
}
func (m *RemoteusrDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteusrDet.Marshal(b, m, deterministic)
}
func (m *RemoteusrDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteusrDet.Merge(m, src)
}
func (m *RemoteusrDet) XXX_Size() int {
	return xxx_messageInfo_RemoteusrDet.Size(m)
}
func (m *RemoteusrDet) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteusrDet.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteusrDet proto.InternalMessageInfo

func (m *RemoteusrDet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoteusrDet) GetUsergroup() []string {
	if m != nil {
		return m.Usergroup
	}
	return nil
}

func (m *RemoteusrDet) GetAuthenmethod() int32 {
	if m != nil {
		return m.Authenmethod
	}
	return 0
}

func (m *RemoteusrDet) GetTaskmap() []string {
	if m != nil {
		return m.Taskmap
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteusrDet_KEYS)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.authen_method.remoteusr_det_KEYS")
	proto.RegisterType((*RemoteusrDet)(nil), "cisco_ios_xr_aaa_locald_oper.aaa.authen_method.remoteusr_det")
}

func init() { proto.RegisterFile("remoteusr_det.proto", fileDescriptor_3cf4956e430e79ab) }

var fileDescriptor_3cf4956e430e79ab = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xbd, 0xaa, 0xc2, 0x40,
	0x10, 0x85, 0x09, 0xf7, 0xa2, 0x64, 0xd1, 0xc2, 0xd5, 0x62, 0x0b, 0x8b, 0x90, 0x2a, 0xd5, 0x16,
	0x46, 0x1f, 0xc1, 0xca, 0x2e, 0x56, 0x56, 0xc3, 0x98, 0x0c, 0x26, 0x98, 0xcd, 0x2c, 0xfb, 0x03,
	0x76, 0xbe, 0xba, 0x90, 0x20, 0x92, 0x6e, 0xe6, 0x83, 0x73, 0x3e, 0x8e, 0xd8, 0x3a, 0x32, 0x1c,
	0x28, 0x7a, 0x07, 0x0d, 0x05, 0x6d, 0x1d, 0x07, 0x96, 0xba, 0xee, 0x7c, 0xcd, 0xd0, 0xb1, 0x87,
	0x97, 0x03, 0x44, 0x84, 0x9e, 0x6b, 0xec, 0x1b, 0x60, 0x4b, 0x4e, 0x23, 0xa2, 0xc6, 0x18, 0x5a,
	0x1a, 0xc0, 0x50, 0x68, 0xb9, 0xc9, 0x77, 0x42, 0xce, 0x6a, 0xe0, 0x72, 0xbe, 0x5d, 0xf3, 0xb7,
	0x58, 0xcf, 0xa8, 0x94, 0xe2, 0x7f, 0x40, 0x43, 0xea, 0x90, 0x25, 0x45, 0x5a, 0x8d, 0xb7, 0xdc,
	0x8b, 0x34, 0x7a, 0x72, 0x0f, 0xc7, 0xd1, 0xaa, 0x32, 0xfb, 0x2b, 0xd2, 0xea, 0x07, 0x64, 0x2e,
	0x56, 0x93, 0x69, 0x12, 0xa9, 0x63, 0x96, 0x14, 0x9b, 0x6a, 0xc6, 0xa4, 0x12, 0xcb, 0x80, 0xfe,
	0x69, 0xd0, 0xaa, 0xd3, 0x98, 0xff, 0xbe, 0xf7, 0xc5, 0xb8, 0xa6, 0xfc, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x8b, 0xf7, 0x89, 0xcd, 0xe4, 0x00, 0x00, 0x00,
}
