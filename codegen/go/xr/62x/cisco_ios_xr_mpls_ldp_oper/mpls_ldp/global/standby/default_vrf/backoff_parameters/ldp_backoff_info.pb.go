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
// source: ldp_backoff_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_backoff_parameters

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

type LdpBackoffInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffInfo_KEYS) Reset()         { *m = LdpBackoffInfo_KEYS{} }
func (m *LdpBackoffInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffInfo_KEYS) ProtoMessage()    {}
func (*LdpBackoffInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_453f70b0799265f5, []int{0}
}

func (m *LdpBackoffInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpBackoffInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpBackoffInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffInfo_KEYS.Merge(m, src)
}
func (m *LdpBackoffInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffInfo_KEYS.Size(m)
}
func (m *LdpBackoffInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffInfo_KEYS proto.InternalMessageInfo

type LdpBackoffInfo struct {
	InitialSeconds       uint32   `protobuf:"varint,50,opt,name=initial_seconds,json=initialSeconds,proto3" json:"initial_seconds,omitempty"`
	MaximumSeconds       uint32   `protobuf:"varint,51,opt,name=maximum_seconds,json=maximumSeconds,proto3" json:"maximum_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffInfo) Reset()         { *m = LdpBackoffInfo{} }
func (m *LdpBackoffInfo) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffInfo) ProtoMessage()    {}
func (*LdpBackoffInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_453f70b0799265f5, []int{1}
}

func (m *LdpBackoffInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffInfo.Unmarshal(m, b)
}
func (m *LdpBackoffInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffInfo.Marshal(b, m, deterministic)
}
func (m *LdpBackoffInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffInfo.Merge(m, src)
}
func (m *LdpBackoffInfo) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffInfo.Size(m)
}
func (m *LdpBackoffInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffInfo proto.InternalMessageInfo

func (m *LdpBackoffInfo) GetInitialSeconds() uint32 {
	if m != nil {
		return m.InitialSeconds
	}
	return 0
}

func (m *LdpBackoffInfo) GetMaximumSeconds() uint32 {
	if m != nil {
		return m.MaximumSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpBackoffInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.backoff_parameters.ldp_backoff_info_KEYS")
	proto.RegisterType((*LdpBackoffInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.backoff_parameters.ldp_backoff_info")
}

func init() { proto.RegisterFile("ldp_backoff_info.proto", fileDescriptor_453f70b0799265f5) }

var fileDescriptor_453f70b0799265f5 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0xbd, 0x8e, 0xc2, 0x30,
	0x0c, 0xc0, 0x71, 0xdd, 0x72, 0x43, 0xa4, 0xfb, 0x50, 0xa5, 0x3b, 0x18, 0x51, 0x17, 0x98, 0x32,
	0xd0, 0x67, 0x60, 0x62, 0x82, 0x4e, 0x4c, 0x96, 0xf3, 0x85, 0x22, 0x92, 0x38, 0x4a, 0x52, 0x54,
	0xde, 0x1e, 0x95, 0xb6, 0x0c, 0x1d, 0xf3, 0xcf, 0x4f, 0xb6, 0xcc, 0xfe, 0x9d, 0x8a, 0x20, 0x50,
	0xde, 0xc8, 0x18, 0xb0, 0xc1, 0x10, 0x8f, 0x89, 0x0a, 0x55, 0x27, 0x69, 0xb3, 0x24, 0xb0, 0x94,
	0xa1, 0x4f, 0xe0, 0xa3, 0xcb, 0x30, 0x48, 0x8a, 0x3a, 0xf1, 0xf9, 0xc5, 0xaf, 0x8e, 0x04, 0x3a,
	0x9e, 0x0b, 0x06, 0x25, 0x1e, 0x5c, 0x69, 0x83, 0x9d, 0x2b, 0x70, 0x4f, 0x86, 0xcf, 0x23, 0x23,
	0x26, 0xf4, 0xba, 0xe8, 0x94, 0xeb, 0x15, 0xfb, 0x5b, 0x2e, 0x83, 0xe3, 0xe1, 0xd2, 0xd6, 0x8a,
	0xfd, 0x2e, 0x3f, 0xaa, 0x2d, 0xfb, 0xb1, 0xc1, 0x16, 0x8b, 0x0e, 0xb2, 0x96, 0x14, 0x54, 0x5e,
	0xef, 0x37, 0x1f, 0xbb, 0xaf, 0xf3, 0xf7, 0x94, 0xdb, 0xb1, 0x0e, 0xd0, 0x63, 0x6f, 0x7d, 0xe7,
	0xdf, 0xb0, 0x19, 0xe1, 0x94, 0x27, 0x28, 0x3e, 0x5f, 0x87, 0x35, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x04, 0xdb, 0x1d, 0xf7, 0xf2, 0x00, 0x00, 0x00,
}
