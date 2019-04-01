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
// source: eem_dir_usr.proto

package cisco_ios_xr_ha_eem_policy_oper_eem_dir_user_library

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

type EemDirUsr_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EemDirUsr_KEYS) Reset()         { *m = EemDirUsr_KEYS{} }
func (m *EemDirUsr_KEYS) String() string { return proto.CompactTextString(m) }
func (*EemDirUsr_KEYS) ProtoMessage()    {}
func (*EemDirUsr_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ea73252327ad00c, []int{0}
}

func (m *EemDirUsr_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EemDirUsr_KEYS.Unmarshal(m, b)
}
func (m *EemDirUsr_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EemDirUsr_KEYS.Marshal(b, m, deterministic)
}
func (m *EemDirUsr_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EemDirUsr_KEYS.Merge(m, src)
}
func (m *EemDirUsr_KEYS) XXX_Size() int {
	return xxx_messageInfo_EemDirUsr_KEYS.Size(m)
}
func (m *EemDirUsr_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EemDirUsr_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EemDirUsr_KEYS proto.InternalMessageInfo

type EemDirUsr struct {
	Policy               string   `protobuf:"bytes,50,opt,name=policy,proto3" json:"policy,omitempty"`
	Library              string   `protobuf:"bytes,51,opt,name=library,proto3" json:"library,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EemDirUsr) Reset()         { *m = EemDirUsr{} }
func (m *EemDirUsr) String() string { return proto.CompactTextString(m) }
func (*EemDirUsr) ProtoMessage()    {}
func (*EemDirUsr) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ea73252327ad00c, []int{1}
}

func (m *EemDirUsr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EemDirUsr.Unmarshal(m, b)
}
func (m *EemDirUsr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EemDirUsr.Marshal(b, m, deterministic)
}
func (m *EemDirUsr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EemDirUsr.Merge(m, src)
}
func (m *EemDirUsr) XXX_Size() int {
	return xxx_messageInfo_EemDirUsr.Size(m)
}
func (m *EemDirUsr) XXX_DiscardUnknown() {
	xxx_messageInfo_EemDirUsr.DiscardUnknown(m)
}

var xxx_messageInfo_EemDirUsr proto.InternalMessageInfo

func (m *EemDirUsr) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func (m *EemDirUsr) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func init() {
	proto.RegisterType((*EemDirUsr_KEYS)(nil), "cisco_ios_xr_ha_eem_policy_oper.eem.dir_user.library.eem_dir_usr_KEYS")
	proto.RegisterType((*EemDirUsr)(nil), "cisco_ios_xr_ha_eem_policy_oper.eem.dir_user.library.eem_dir_usr")
}

func init() { proto.RegisterFile("eem_dir_usr.proto", fileDescriptor_4ea73252327ad00c) }

var fileDescriptor_4ea73252327ad00c = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4d, 0xcd, 0x8d,
	0x4f, 0xc9, 0x2c, 0x8a, 0x2f, 0x2d, 0x2e, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x49,
	0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0x48, 0x8c, 0x07,
	0x29, 0x29, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x8c, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x4b, 0x4d, 0xcd,
	0xd5, 0x83, 0x68, 0x49, 0x2d, 0xd2, 0xcb, 0xc9, 0x4c, 0x2a, 0x4a, 0x2c, 0xaa, 0x54, 0x12, 0xe2,
	0x12, 0x40, 0x32, 0x2a, 0xde, 0xdb, 0x35, 0x32, 0x58, 0xc9, 0x9e, 0x8b, 0x1b, 0x49, 0x4c, 0x48,
	0x8c, 0x8b, 0x0d, 0x62, 0x8c, 0x84, 0x91, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x24, 0xc1,
	0xc5, 0x0e, 0x35, 0x45, 0xc2, 0x18, 0x2c, 0x01, 0xe3, 0x26, 0xb1, 0x81, 0x5d, 0x64, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x25, 0xe7, 0x59, 0x9a, 0xa6, 0x00, 0x00, 0x00,
}