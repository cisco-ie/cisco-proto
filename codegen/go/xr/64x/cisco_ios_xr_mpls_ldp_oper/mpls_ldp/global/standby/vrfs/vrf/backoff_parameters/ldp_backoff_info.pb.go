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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_backoff_parameters

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
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

func (m *LdpBackoffInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*LdpBackoffInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.backoff_parameters.ldp_backoff_info_KEYS")
	proto.RegisterType((*LdpBackoffInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.backoff_parameters.ldp_backoff_info")
}

func init() { proto.RegisterFile("ldp_backoff_info.proto", fileDescriptor_453f70b0799265f5) }

var fileDescriptor_453f70b0799265f5 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0x3b, 0x6b, 0x03, 0x31,
	0x0c, 0x07, 0x70, 0x6e, 0xe9, 0xc3, 0xd0, 0x07, 0x07, 0x2d, 0xe9, 0x16, 0xb2, 0x34, 0x93, 0x87,
	0xe4, 0x33, 0x74, 0x2a, 0x64, 0x48, 0xa6, 0x4e, 0x42, 0x7e, 0x15, 0x53, 0xdb, 0x32, 0x92, 0x7b,
	0xa4, 0xdf, 0xbe, 0x5c, 0x73, 0xd7, 0xe1, 0x16, 0x81, 0xfe, 0xfc, 0x24, 0x21, 0xf5, 0x9c, 0x5c,
	0x05, 0x83, 0xf6, 0x8b, 0x42, 0x80, 0x58, 0x02, 0xe9, 0xca, 0xd4, 0xa8, 0x3f, 0xd8, 0x28, 0x96,
	0x20, 0x92, 0xc0, 0x99, 0x21, 0xd7, 0x24, 0x30, 0x4a, 0xaa, 0x9e, 0xf5, 0xdc, 0xe9, 0xcf, 0x44,
	0x06, 0x93, 0x96, 0x86, 0xc5, 0x99, 0x1f, 0x3d, 0x70, 0x90, 0xb1, 0xe8, 0x79, 0x5f, 0x45, 0xc6,
	0xec, 0x9b, 0x67, 0xd9, 0xec, 0xd4, 0xd3, 0xf2, 0x12, 0xbc, 0xbf, 0x7d, 0x9c, 0xfa, 0x17, 0x75,
	0x33, 0x70, 0x80, 0x82, 0xd9, 0xaf, 0xba, 0x75, 0xb7, 0xbd, 0x3d, 0x5e, 0x0f, 0x1c, 0x0e, 0x98,
	0xfd, 0xc6, 0xa9, 0xc7, 0xe5, 0x4c, 0xff, 0xaa, 0x1e, 0x62, 0x89, 0x2d, 0x62, 0x02, 0xf1, 0x96,
	0x8a, 0x93, 0xd5, 0x6e, 0xdd, 0x6d, 0xef, 0x8e, 0xf7, 0x53, 0x7c, 0xba, 0xa4, 0x23, 0xcc, 0x78,
	0x8e, 0xf9, 0x3b, 0xff, 0xc3, 0xfd, 0x05, 0x4e, 0xf1, 0x04, 0xcd, 0xd5, 0xdf, 0xc3, 0xfb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xc4, 0xd5, 0x3b, 0x0a, 0x01, 0x00, 0x00,
}
