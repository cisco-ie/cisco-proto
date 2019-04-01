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
	Shelltype            string   `protobuf:"bytes,52,opt,name=shelltype,proto3" json:"shelltype,omitempty"`
	Directory            string   `protobuf:"bytes,53,opt,name=directory,proto3" json:"directory,omitempty"`
	Authenmethod         int32    `protobuf:"zigzag32,54,opt,name=authenmethod,proto3" json:"authenmethod,omitempty"`
	Taskmap              []string `protobuf:"bytes,55,rep,name=taskmap,proto3" json:"taskmap,omitempty"`
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

func (m *RemoteusrDet) GetShelltype() string {
	if m != nil {
		return m.Shelltype
	}
	return ""
}

func (m *RemoteusrDet) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
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
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbf, 0x4e, 0xc4, 0x30,
	0x0c, 0x87, 0x55, 0x81, 0x40, 0x8d, 0x60, 0x20, 0x30, 0x64, 0x60, 0xa8, 0x3a, 0x75, 0xea, 0xc0,
	0xf1, 0xe7, 0x09, 0x98, 0xd8, 0xca, 0xc4, 0x64, 0x99, 0xd6, 0xa2, 0x15, 0xcd, 0x39, 0x72, 0x1c,
	0x89, 0x7b, 0x36, 0x5e, 0x0e, 0x5d, 0x2a, 0xae, 0xea, 0x66, 0x7f, 0x9f, 0x6c, 0xff, 0x64, 0x73,
	0x2b, 0xe4, 0x59, 0x29, 0x45, 0x81, 0x81, 0xb4, 0x0d, 0xc2, 0xca, 0xb6, 0xed, 0xa7, 0xd8, 0x33,
	0x4c, 0x1c, 0xe1, 0x47, 0x00, 0x11, 0x61, 0xe6, 0x1e, 0xe7, 0x01, 0x38, 0x90, 0xb4, 0x88, 0xd8,
	0x62, 0xd2, 0x91, 0xf6, 0xe0, 0x49, 0x47, 0x1e, 0xea, 0x3b, 0x63, 0x37, 0x6b, 0xe0, 0xed, 0xf5,
	0xe3, 0xbd, 0xfe, 0x2d, 0xcc, 0xf5, 0x06, 0x5b, 0x6b, 0xce, 0xf7, 0xe8, 0xc9, 0x3d, 0x54, 0x45,
	0x53, 0x76, 0xb9, 0xb6, 0xf7, 0xa6, 0x4c, 0x91, 0xe4, 0x4b, 0x38, 0x05, 0xb7, 0xab, 0xce, 0x9a,
	0xb2, 0x5b, 0xc1, 0xd1, 0xc6, 0x91, 0xe6, 0x59, 0x0f, 0x81, 0xdc, 0x63, 0x1e, 0x5b, 0xc1, 0xd1,
	0x0e, 0x93, 0x50, 0xaf, 0x2c, 0x07, 0xf7, 0xb4, 0xd8, 0x13, 0xb0, 0xb5, 0xb9, 0x5a, 0x62, 0x2e,
	0x29, 0xdd, 0x73, 0x55, 0x34, 0x37, 0xdd, 0x86, 0x59, 0x67, 0x2e, 0x15, 0xe3, 0xb7, 0xc7, 0xe0,
	0x5e, 0xf2, 0xed, 0xff, 0xf6, 0xf3, 0x22, 0xbf, 0x62, 0xf7, 0x17, 0x00, 0x00, 0xff, 0xff, 0xdc,
	0xb8, 0xb7, 0xcf, 0x21, 0x01, 0x00, 0x00,
}
