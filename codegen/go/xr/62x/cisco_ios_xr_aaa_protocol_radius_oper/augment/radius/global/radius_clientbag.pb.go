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
// source: radius_clientbag.proto

package cisco_ios_xr_aaa_protocol_radius_oper_augment_radius_global

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

type RadiusClientbag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusClientbag_KEYS) Reset()         { *m = RadiusClientbag_KEYS{} }
func (m *RadiusClientbag_KEYS) String() string { return proto.CompactTextString(m) }
func (*RadiusClientbag_KEYS) ProtoMessage()    {}
func (*RadiusClientbag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c6f2ed9afa7647, []int{0}
}

func (m *RadiusClientbag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusClientbag_KEYS.Unmarshal(m, b)
}
func (m *RadiusClientbag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusClientbag_KEYS.Marshal(b, m, deterministic)
}
func (m *RadiusClientbag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusClientbag_KEYS.Merge(m, src)
}
func (m *RadiusClientbag_KEYS) XXX_Size() int {
	return xxx_messageInfo_RadiusClientbag_KEYS.Size(m)
}
func (m *RadiusClientbag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusClientbag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusClientbag_KEYS proto.InternalMessageInfo

type RadiusClientbag struct {
	UnknownAuthenticationResponse uint32   `protobuf:"varint,50,opt,name=unknown_authentication_response,json=unknownAuthenticationResponse,proto3" json:"unknown_authentication_response,omitempty"`
	AuthenticationNasId           string   `protobuf:"bytes,51,opt,name=authentication_nas_id,json=authenticationNasId,proto3" json:"authentication_nas_id,omitempty"`
	UnknownAccountingResponse     uint32   `protobuf:"varint,52,opt,name=unknown_accounting_response,json=unknownAccountingResponse,proto3" json:"unknown_accounting_response,omitempty"`
	AccountingNasId               string   `protobuf:"bytes,53,opt,name=accounting_nas_id,json=accountingNasId,proto3" json:"accounting_nas_id,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *RadiusClientbag) Reset()         { *m = RadiusClientbag{} }
func (m *RadiusClientbag) String() string { return proto.CompactTextString(m) }
func (*RadiusClientbag) ProtoMessage()    {}
func (*RadiusClientbag) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c6f2ed9afa7647, []int{1}
}

func (m *RadiusClientbag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusClientbag.Unmarshal(m, b)
}
func (m *RadiusClientbag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusClientbag.Marshal(b, m, deterministic)
}
func (m *RadiusClientbag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusClientbag.Merge(m, src)
}
func (m *RadiusClientbag) XXX_Size() int {
	return xxx_messageInfo_RadiusClientbag.Size(m)
}
func (m *RadiusClientbag) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusClientbag.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusClientbag proto.InternalMessageInfo

func (m *RadiusClientbag) GetUnknownAuthenticationResponse() uint32 {
	if m != nil {
		return m.UnknownAuthenticationResponse
	}
	return 0
}

func (m *RadiusClientbag) GetAuthenticationNasId() string {
	if m != nil {
		return m.AuthenticationNasId
	}
	return ""
}

func (m *RadiusClientbag) GetUnknownAccountingResponse() uint32 {
	if m != nil {
		return m.UnknownAccountingResponse
	}
	return 0
}

func (m *RadiusClientbag) GetAccountingNasId() string {
	if m != nil {
		return m.AccountingNasId
	}
	return ""
}

func init() {
	proto.RegisterType((*RadiusClientbag_KEYS)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.augment.radius.global.radius_clientbag_KEYS")
	proto.RegisterType((*RadiusClientbag)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.augment.radius.global.radius_clientbag")
}

func init() { proto.RegisterFile("radius_clientbag.proto", fileDescriptor_09c6f2ed9afa7647) }

var fileDescriptor_09c6f2ed9afa7647 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0x3f, 0x4b, 0xc4, 0x40,
	0x10, 0x05, 0x70, 0xd2, 0x08, 0x2e, 0x88, 0x1a, 0x39, 0x3d, 0x11, 0xf1, 0xb8, 0xea, 0xb0, 0x48,
	0x71, 0xa7, 0x95, 0x20, 0x58, 0x28, 0x88, 0x60, 0x11, 0x2b, 0xab, 0x61, 0xb2, 0x59, 0xe2, 0xe0,
	0x3a, 0x13, 0xf6, 0x0f, 0xfa, 0xd9, 0xad, 0x84, 0x75, 0xcd, 0x99, 0xb4, 0xf3, 0xde, 0xdb, 0x1f,
	0xac, 0x3a, 0x76, 0xd8, 0x52, 0xf4, 0xa0, 0x2d, 0x19, 0x0e, 0x0d, 0x76, 0x55, 0xef, 0x24, 0x48,
	0x79, 0xa3, 0xc9, 0x6b, 0x01, 0x12, 0x0f, 0x5f, 0x0e, 0x10, 0x11, 0xd2, 0x5d, 0x8b, 0x85, 0xbc,
	0x90, 0xde, 0xb8, 0x0a, 0x63, 0xf7, 0x61, 0x38, 0x54, 0xbf, 0xb7, 0xaa, 0xb3, 0xd2, 0xa0, 0x5d,
	0x9e, 0xa8, 0xd9, 0xf4, 0x59, 0x78, 0xba, 0x7f, 0x7d, 0x59, 0x7e, 0x17, 0xea, 0x60, 0x9a, 0x94,
	0x0f, 0xea, 0x22, 0xf2, 0x3b, 0xcb, 0x27, 0x03, 0xc6, 0xf0, 0x66, 0x38, 0x90, 0xc6, 0x40, 0xc2,
	0xe0, 0x8c, 0xef, 0x85, 0xbd, 0x99, 0xaf, 0x17, 0xc5, 0x6a, 0xaf, 0x3e, 0xcf, 0xb5, 0xbb, 0x51,
	0xab, 0xce, 0xa5, 0x72, 0xad, 0x66, 0x93, 0x3d, 0xa3, 0x07, 0x6a, 0xe7, 0x9b, 0x45, 0xb1, 0xda,
	0xad, 0x8f, 0xc6, 0xe1, 0x33, 0xfa, 0xc7, 0xb6, 0xbc, 0x55, 0x67, 0x83, 0xad, 0xb5, 0x44, 0x0e,
	0xc4, 0xdd, 0xd6, 0xbd, 0x4a, 0xee, 0xe9, 0x9f, 0x3b, 0x34, 0x06, 0xf3, 0x52, 0x1d, 0xfe, 0xdb,
	0x65, 0xef, 0x3a, 0x79, 0xfb, 0xdb, 0x20, 0x59, 0xcd, 0x4e, 0xfa, 0xc1, 0xcd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x19, 0x45, 0x48, 0xdd, 0x73, 0x01, 0x00, 0x00,
}
