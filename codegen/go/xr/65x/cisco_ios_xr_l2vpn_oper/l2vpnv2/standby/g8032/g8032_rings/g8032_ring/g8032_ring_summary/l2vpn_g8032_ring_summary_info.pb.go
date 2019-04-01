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
// source: l2vpn_g8032_ring_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_g8032_g8032_rings_g8032_ring_g8032_ring_summary

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

type L2VpnG8032RingSummaryInfo_KEYS struct {
	RingName             string   `protobuf:"bytes,1,opt,name=ring_name,json=ringName,proto3" json:"ring_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnG8032RingSummaryInfo_KEYS) Reset()         { *m = L2VpnG8032RingSummaryInfo_KEYS{} }
func (m *L2VpnG8032RingSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnG8032RingSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2VpnG8032RingSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3c965c504e3424, []int{0}
}

func (m *L2VpnG8032RingSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2VpnG8032RingSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnG8032RingSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2VpnG8032RingSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS.Size(m)
}
func (m *L2VpnG8032RingSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnG8032RingSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2VpnG8032RingSummaryInfo_KEYS) GetRingName() string {
	if m != nil {
		return m.RingName
	}
	return ""
}

type L2VpnG8032RingSummaryInfo struct {
	RingName             string   `protobuf:"bytes,50,opt,name=ring_name,json=ringName,proto3" json:"ring_name,omitempty"`
	Port0                string   `protobuf:"bytes,51,opt,name=port0,proto3" json:"port0,omitempty"`
	Port1                string   `protobuf:"bytes,52,opt,name=port1,proto3" json:"port1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnG8032RingSummaryInfo) Reset()         { *m = L2VpnG8032RingSummaryInfo{} }
func (m *L2VpnG8032RingSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnG8032RingSummaryInfo) ProtoMessage()    {}
func (*L2VpnG8032RingSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a3c965c504e3424, []int{1}
}

func (m *L2VpnG8032RingSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo.Unmarshal(m, b)
}
func (m *L2VpnG8032RingSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnG8032RingSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnG8032RingSummaryInfo.Merge(m, src)
}
func (m *L2VpnG8032RingSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnG8032RingSummaryInfo.Size(m)
}
func (m *L2VpnG8032RingSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnG8032RingSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnG8032RingSummaryInfo proto.InternalMessageInfo

func (m *L2VpnG8032RingSummaryInfo) GetRingName() string {
	if m != nil {
		return m.RingName
	}
	return ""
}

func (m *L2VpnG8032RingSummaryInfo) GetPort0() string {
	if m != nil {
		return m.Port0
	}
	return ""
}

func (m *L2VpnG8032RingSummaryInfo) GetPort1() string {
	if m != nil {
		return m.Port1
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnG8032RingSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.g8032.g8032_rings.g8032_ring.g8032_ring_summary.l2vpn_g8032_ring_summary_info_KEYS")
	proto.RegisterType((*L2VpnG8032RingSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.g8032.g8032_rings.g8032_ring.g8032_ring_summary.l2vpn_g8032_ring_summary_info")
}

func init() {
	proto.RegisterFile("l2vpn_g8032_ring_summary_info.proto", fileDescriptor_9a3c965c504e3424)
}

var fileDescriptor_9a3c965c504e3424 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x31, 0x2a, 0x2b,
	0xc8, 0x8b, 0x4f, 0xb7, 0x30, 0x30, 0x36, 0x8a, 0x2f, 0xca, 0xcc, 0x4b, 0x8f, 0x2f, 0x2e, 0xcd,
	0xcd, 0x4d, 0x2c, 0xaa, 0x8c, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x0a, 0x4f, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x87, 0xe8,
	0xc8, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0x33, 0xcb, 0x8c, 0xf4, 0x8a, 0x4b, 0x12, 0xf3, 0x52, 0x92,
	0x2a, 0xf5, 0xc0, 0xc6, 0xe8, 0x21, 0x0c, 0x2b, 0x46, 0x62, 0xeb, 0x61, 0xda, 0xa1, 0xe4, 0xc8,
	0xa5, 0x84, 0xd7, 0xfe, 0x78, 0x6f, 0xd7, 0xc8, 0x60, 0x21, 0x69, 0x2e, 0x4e, 0xb0, 0x4c, 0x5e,
	0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x07, 0x48, 0xc0, 0x2f, 0x31, 0x37,
	0x55, 0x29, 0x83, 0x4b, 0x16, 0xaf, 0x11, 0xa8, 0xba, 0x8d, 0x50, 0x75, 0x0b, 0x89, 0x70, 0xb1,
	0x16, 0xe4, 0x17, 0x95, 0x18, 0x48, 0x18, 0x83, 0x25, 0x20, 0x1c, 0x98, 0xa8, 0xa1, 0x84, 0x09,
	0x42, 0xd4, 0x30, 0x89, 0x0d, 0x1c, 0x18, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x63,
	0x1d, 0x70, 0x33, 0x01, 0x00, 0x00,
}