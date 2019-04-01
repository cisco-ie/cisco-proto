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
// source: mpls_lsd_frr_db_sum.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_frr_database_tunnel_midpoint_summary

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

type MplsLsdFrrDbSum_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbSum_KEYS) Reset()         { *m = MplsLsdFrrDbSum_KEYS{} }
func (m *MplsLsdFrrDbSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbSum_KEYS) ProtoMessage()    {}
func (*MplsLsdFrrDbSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c23a653836dc035, []int{0}
}

func (m *MplsLsdFrrDbSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Merge(m, src)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Size(m)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbSum_KEYS proto.InternalMessageInfo

type MplsLsdFrrDbSum struct {
	Active               uint32   `protobuf:"varint,50,opt,name=active,proto3" json:"active,omitempty"`
	Ready                uint32   `protobuf:"varint,51,opt,name=ready,proto3" json:"ready,omitempty"`
	Partial              uint32   `protobuf:"varint,52,opt,name=partial,proto3" json:"partial,omitempty"`
	Igp                  uint32   `protobuf:"varint,53,opt,name=igp,proto3" json:"igp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbSum) Reset()         { *m = MplsLsdFrrDbSum{} }
func (m *MplsLsdFrrDbSum) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbSum) ProtoMessage()    {}
func (*MplsLsdFrrDbSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c23a653836dc035, []int{1}
}

func (m *MplsLsdFrrDbSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbSum.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbSum.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbSum.Merge(m, src)
}
func (m *MplsLsdFrrDbSum) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbSum.Size(m)
}
func (m *MplsLsdFrrDbSum) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbSum.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbSum proto.InternalMessageInfo

func (m *MplsLsdFrrDbSum) GetActive() uint32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetReady() uint32 {
	if m != nil {
		return m.Ready
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetPartial() uint32 {
	if m != nil {
		return m.Partial
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetIgp() uint32 {
	if m != nil {
		return m.Igp
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLsdFrrDbSum_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd.frr_database.tunnel_midpoint_summary.mpls_lsd_frr_db_sum_KEYS")
	proto.RegisterType((*MplsLsdFrrDbSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd.frr_database.tunnel_midpoint_summary.mpls_lsd_frr_db_sum")
}

func init() { proto.RegisterFile("mpls_lsd_frr_db_sum.proto", fileDescriptor_0c23a653836dc035) }

var fileDescriptor_0c23a653836dc035 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcf, 0x8a, 0xc2, 0x30,
	0x10, 0x87, 0x59, 0x96, 0xed, 0x42, 0x60, 0x61, 0xc9, 0x2e, 0x12, 0x3d, 0x49, 0x4f, 0x9e, 0x72,
	0xb0, 0xfa, 0x08, 0x82, 0xe0, 0x4d, 0x4f, 0x9e, 0x86, 0x69, 0x13, 0x25, 0x90, 0x34, 0x61, 0x92,
	0x8a, 0x7d, 0x7b, 0x69, 0xb4, 0x9e, 0x7a, 0x9b, 0xef, 0x9b, 0x7f, 0xfc, 0xd8, 0xdc, 0x05, 0x1b,
	0xc1, 0x46, 0x05, 0x17, 0x22, 0x50, 0x35, 0xc4, 0xce, 0xc9, 0x40, 0x3e, 0x79, 0xbe, 0x6f, 0x4c,
	0x6c, 0x3c, 0x18, 0x1f, 0xe1, 0x4e, 0xf0, 0x9e, 0xf3, 0x41, 0x93, 0x1c, 0x49, 0xe6, 0x2d, 0x4c,
	0x58, 0x63, 0xd4, 0x32, 0x75, 0x6d, 0xab, 0x2d, 0x38, 0xa3, 0x82, 0x37, 0x6d, 0x1a, 0x6e, 0x39,
	0xa4, 0xbe, 0x5c, 0x30, 0x31, 0xf1, 0x06, 0x0e, 0xbb, 0xf3, 0xa9, 0xf4, 0xec, 0x6f, 0xa2, 0xc7,
	0x67, 0xac, 0xc0, 0x26, 0x99, 0x9b, 0x16, 0xeb, 0xe5, 0xc7, 0xea, 0xe7, 0xf8, 0x22, 0xfe, 0xcf,
	0xbe, 0x48, 0xa3, 0xea, 0x45, 0x95, 0xf5, 0x13, 0xb8, 0x60, 0xdf, 0x01, 0x29, 0x19, 0xb4, 0x62,
	0x93, 0xfd, 0x88, 0xfc, 0x97, 0x7d, 0x9a, 0x6b, 0x10, 0xdb, 0x6c, 0x87, 0xb2, 0x2e, 0x72, 0xba,
	0xea, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x10, 0xce, 0xf7, 0xfa, 0x00, 0x00, 0x00,
}