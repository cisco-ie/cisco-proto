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
// source: mpls_lm_bw_alloc_common_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_bandwidth_allocation_global

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

type MplsLmBwAllocCommonInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBwAllocCommonInfo_KEYS) Reset()         { *m = MplsLmBwAllocCommonInfo_KEYS{} }
func (m *MplsLmBwAllocCommonInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmBwAllocCommonInfo_KEYS) ProtoMessage()    {}
func (*MplsLmBwAllocCommonInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3976b8434f6c40, []int{0}
}

func (m *MplsLmBwAllocCommonInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmBwAllocCommonInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmBwAllocCommonInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS.Merge(m, src)
}
func (m *MplsLmBwAllocCommonInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS.Size(m)
}
func (m *MplsLmBwAllocCommonInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBwAllocCommonInfo_KEYS proto.InternalMessageInfo

type MplsLmBwAllocCommonInfo struct {
	IsRoleStandby        bool     `protobuf:"varint,50,opt,name=is_role_standby,json=isRoleStandby,proto3" json:"is_role_standby,omitempty"`
	Links                uint32   `protobuf:"varint,51,opt,name=links,proto3" json:"links,omitempty"`
	BandwidthHoldTime    uint32   `protobuf:"varint,52,opt,name=bandwidth_hold_time,json=bandwidthHoldTime,proto3" json:"bandwidth_hold_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBwAllocCommonInfo) Reset()         { *m = MplsLmBwAllocCommonInfo{} }
func (m *MplsLmBwAllocCommonInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBwAllocCommonInfo) ProtoMessage()    {}
func (*MplsLmBwAllocCommonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e3976b8434f6c40, []int{1}
}

func (m *MplsLmBwAllocCommonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo.Unmarshal(m, b)
}
func (m *MplsLmBwAllocCommonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmBwAllocCommonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBwAllocCommonInfo.Merge(m, src)
}
func (m *MplsLmBwAllocCommonInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBwAllocCommonInfo.Size(m)
}
func (m *MplsLmBwAllocCommonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBwAllocCommonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBwAllocCommonInfo proto.InternalMessageInfo

func (m *MplsLmBwAllocCommonInfo) GetIsRoleStandby() bool {
	if m != nil {
		return m.IsRoleStandby
	}
	return false
}

func (m *MplsLmBwAllocCommonInfo) GetLinks() uint32 {
	if m != nil {
		return m.Links
	}
	return 0
}

func (m *MplsLmBwAllocCommonInfo) GetBandwidthHoldTime() uint32 {
	if m != nil {
		return m.BandwidthHoldTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLmBwAllocCommonInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.bandwidth_allocation.global.mpls_lm_bw_alloc_common_info_KEYS")
	proto.RegisterType((*MplsLmBwAllocCommonInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.bandwidth_allocation.global.mpls_lm_bw_alloc_common_info")
}

func init() { proto.RegisterFile("mpls_lm_bw_alloc_common_info.proto", fileDescriptor_5e3976b8434f6c40) }

var fileDescriptor_5e3976b8434f6c40 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0x4b, 0xc6, 0x30,
	0x10, 0x87, 0xe9, 0xa0, 0x48, 0xe0, 0x45, 0x8c, 0x0e, 0x1d, 0x1c, 0x5e, 0x2b, 0x48, 0xa7, 0x0c,
	0xd6, 0xdd, 0x49, 0x10, 0xdc, 0x5a, 0x17, 0xa7, 0x23, 0xff, 0xb4, 0x87, 0x97, 0x5c, 0x49, 0x02,
	0xd5, 0xef, 0xe0, 0x87, 0x16, 0x1a, 0xd0, 0xad, 0xe3, 0xdd, 0x3d, 0xcf, 0xdd, 0xfd, 0x44, 0x17,
	0x16, 0xca, 0x40, 0x01, 0xcc, 0x0a, 0x9a, 0x88, 0x2d, 0x58, 0x0e, 0x81, 0x23, 0x60, 0x7c, 0x67,
	0xb5, 0x24, 0x2e, 0x2c, 0x1f, 0x2d, 0x66, 0xcb, 0x80, 0x9c, 0xe1, 0x2b, 0xc1, 0x26, 0x14, 0x0f,
	0xbc, 0xf8, 0xa4, 0xaa, 0x6d, 0xb5, 0x55, 0x46, 0x47, 0xb7, 0xa2, 0x2b, 0x73, 0x5d, 0xa3, 0x0b,
	0x72, 0x54, 0x1f, 0xc4, 0x46, 0x53, 0x77, 0x2b, 0x6e, 0xf6, 0xce, 0xc0, 0xcb, 0xd3, 0xdb, 0xd4,
	0xfd, 0x34, 0xe2, 0x7a, 0x8f, 0x92, 0x77, 0xe2, 0x1c, 0x33, 0x24, 0x26, 0x0f, 0xb9, 0xe8, 0xe8,
	0xcc, 0x77, 0x7b, 0x7f, 0x6c, 0xfa, 0xb3, 0xf1, 0x80, 0x79, 0x64, 0xf2, 0x53, 0x6d, 0xca, 0x2b,
	0x71, 0x42, 0x18, 0x3f, 0x73, 0x3b, 0x1c, 0x9b, 0xfe, 0x30, 0xd6, 0x42, 0x2a, 0x71, 0xf9, 0xff,
	0xe2, 0xcc, 0xe4, 0xa0, 0x60, 0xf0, 0xed, 0xc3, 0xc6, 0x5c, 0xfc, 0x8d, 0x9e, 0x99, 0xdc, 0x2b,
	0x06, 0x6f, 0x4e, 0xb7, 0xec, 0xc3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xcf, 0x90, 0xac,
	0x21, 0x01, 0x00, 0x00,
}
