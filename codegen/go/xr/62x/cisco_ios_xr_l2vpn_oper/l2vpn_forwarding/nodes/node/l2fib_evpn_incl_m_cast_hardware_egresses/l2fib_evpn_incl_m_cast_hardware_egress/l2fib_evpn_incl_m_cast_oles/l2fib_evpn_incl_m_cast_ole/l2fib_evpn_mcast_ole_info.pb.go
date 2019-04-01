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
// source: l2fib_evpn_mcast_ole_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_evpn_incl_m_cast_hardware_egresses_l2fib_evpn_incl_m_cast_hardware_egress_l2fib_evpn_incl_m_cast_oles_l2fib_evpn_incl_m_cast_ole

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

type L2FibEvpnMcastOleInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	BdName               string   `protobuf:"bytes,2,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,3,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibEvpnMcastOleInfo_KEYS) Reset()         { *m = L2FibEvpnMcastOleInfo_KEYS{} }
func (m *L2FibEvpnMcastOleInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnMcastOleInfo_KEYS) ProtoMessage()    {}
func (*L2FibEvpnMcastOleInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4eb1aecc50742a, []int{0}
}

func (m *L2FibEvpnMcastOleInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibEvpnMcastOleInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnMcastOleInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS.Merge(m, src)
}
func (m *L2FibEvpnMcastOleInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS.Size(m)
}
func (m *L2FibEvpnMcastOleInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnMcastOleInfo_KEYS proto.InternalMessageInfo

func (m *L2FibEvpnMcastOleInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibEvpnMcastOleInfo_KEYS) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

func (m *L2FibEvpnMcastOleInfo_KEYS) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

type L2VpnEvpnMoiInfo struct {
	NextHopIpv6Addr      string   `protobuf:"bytes,1,opt,name=next_hop_ipv6_addr,json=nextHopIpv6Addr,proto3" json:"next_hop_ipv6_addr,omitempty"`
	McastLabel           uint32   `protobuf:"varint,2,opt,name=mcast_label,json=mcastLabel,proto3" json:"mcast_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnMoiInfo) Reset()         { *m = L2VpnEvpnMoiInfo{} }
func (m *L2VpnEvpnMoiInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnMoiInfo) ProtoMessage()    {}
func (*L2VpnEvpnMoiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4eb1aecc50742a, []int{1}
}

func (m *L2VpnEvpnMoiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnMoiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnMoiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnMoiInfo.Merge(m, src)
}
func (m *L2VpnEvpnMoiInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Size(m)
}
func (m *L2VpnEvpnMoiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnMoiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnMoiInfo proto.InternalMessageInfo

func (m *L2VpnEvpnMoiInfo) GetNextHopIpv6Addr() string {
	if m != nil {
		return m.NextHopIpv6Addr
	}
	return ""
}

func (m *L2VpnEvpnMoiInfo) GetMcastLabel() uint32 {
	if m != nil {
		return m.McastLabel
	}
	return 0
}

type L2FibBaseInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBaseInfo) Reset()         { *m = L2FibBaseInfo{} }
func (m *L2FibBaseInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBaseInfo) ProtoMessage()    {}
func (*L2FibBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4eb1aecc50742a, []int{2}
}

func (m *L2FibBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBaseInfo.Unmarshal(m, b)
}
func (m *L2FibBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBaseInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBaseInfo.Merge(m, src)
}
func (m *L2FibBaseInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBaseInfo.Size(m)
}
func (m *L2FibBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBaseInfo proto.InternalMessageInfo

type L2FibNhopInfo struct {
	Base                 *L2FibBaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	NextHopAddress       string         `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	NextHopInternalLabel uint32         `protobuf:"varint,3,opt,name=next_hop_internal_label,json=nextHopInternalLabel,proto3" json:"next_hop_internal_label,omitempty"`
	PlaformtDataValid    bool           `protobuf:"varint,4,opt,name=plaformt_data_valid,json=plaformtDataValid,proto3" json:"plaformt_data_valid,omitempty"`
	PlatformDataLength   uint32         `protobuf:"varint,5,opt,name=platform_data_length,json=platformDataLength,proto3" json:"platform_data_length,omitempty"`
	ChildrenCount        uint32         `protobuf:"varint,6,opt,name=children_count,json=childrenCount,proto3" json:"children_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *L2FibNhopInfo) Reset()         { *m = L2FibNhopInfo{} }
func (m *L2FibNhopInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibNhopInfo) ProtoMessage()    {}
func (*L2FibNhopInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4eb1aecc50742a, []int{3}
}

func (m *L2FibNhopInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibNhopInfo.Unmarshal(m, b)
}
func (m *L2FibNhopInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibNhopInfo.Marshal(b, m, deterministic)
}
func (m *L2FibNhopInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibNhopInfo.Merge(m, src)
}
func (m *L2FibNhopInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibNhopInfo.Size(m)
}
func (m *L2FibNhopInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibNhopInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibNhopInfo proto.InternalMessageInfo

func (m *L2FibNhopInfo) GetBase() *L2FibBaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *L2FibNhopInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *L2FibNhopInfo) GetNextHopInternalLabel() uint32 {
	if m != nil {
		return m.NextHopInternalLabel
	}
	return 0
}

func (m *L2FibNhopInfo) GetPlaformtDataValid() bool {
	if m != nil {
		return m.PlaformtDataValid
	}
	return false
}

func (m *L2FibNhopInfo) GetPlatformDataLength() uint32 {
	if m != nil {
		return m.PlatformDataLength
	}
	return 0
}

func (m *L2FibNhopInfo) GetChildrenCount() uint32 {
	if m != nil {
		return m.ChildrenCount
	}
	return 0
}

type L2FibEvpnMcastOleInfo struct {
	McastOle             *L2VpnEvpnMoiInfo `protobuf:"bytes,50,opt,name=mcast_ole,json=mcastOle,proto3" json:"mcast_ole,omitempty"`
	IsBound              bool              `protobuf:"varint,51,opt,name=is_bound,json=isBound,proto3" json:"is_bound,omitempty"`
	NextHopValid         bool              `protobuf:"varint,52,opt,name=next_hop_valid,json=nextHopValid,proto3" json:"next_hop_valid,omitempty"`
	NextHop              *L2FibNhopInfo    `protobuf:"bytes,53,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2FibEvpnMcastOleInfo) Reset()         { *m = L2FibEvpnMcastOleInfo{} }
func (m *L2FibEvpnMcastOleInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnMcastOleInfo) ProtoMessage()    {}
func (*L2FibEvpnMcastOleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f4eb1aecc50742a, []int{4}
}

func (m *L2FibEvpnMcastOleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo.Unmarshal(m, b)
}
func (m *L2FibEvpnMcastOleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnMcastOleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnMcastOleInfo.Merge(m, src)
}
func (m *L2FibEvpnMcastOleInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnMcastOleInfo.Size(m)
}
func (m *L2FibEvpnMcastOleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnMcastOleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnMcastOleInfo proto.InternalMessageInfo

func (m *L2FibEvpnMcastOleInfo) GetMcastOle() *L2VpnEvpnMoiInfo {
	if m != nil {
		return m.McastOle
	}
	return nil
}

func (m *L2FibEvpnMcastOleInfo) GetIsBound() bool {
	if m != nil {
		return m.IsBound
	}
	return false
}

func (m *L2FibEvpnMcastOleInfo) GetNextHopValid() bool {
	if m != nil {
		return m.NextHopValid
	}
	return false
}

func (m *L2FibEvpnMcastOleInfo) GetNextHop() *L2FibNhopInfo {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func init() {
	proto.RegisterType((*L2FibEvpnMcastOleInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_egresses.l2fib_evpn_incl_m_cast_hardware_egress.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_evpn_mcast_ole_info_KEYS")
	proto.RegisterType((*L2VpnEvpnMoiInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_egresses.l2fib_evpn_incl_m_cast_hardware_egress.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2vpn_evpn_moi_info")
	proto.RegisterType((*L2FibBaseInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_egresses.l2fib_evpn_incl_m_cast_hardware_egress.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_base_info")
	proto.RegisterType((*L2FibNhopInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_egresses.l2fib_evpn_incl_m_cast_hardware_egress.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_nhop_info")
	proto.RegisterType((*L2FibEvpnMcastOleInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_egresses.l2fib_evpn_incl_m_cast_hardware_egress.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_evpn_mcast_ole_info")
}

func init() { proto.RegisterFile("l2fib_evpn_mcast_ole_info.proto", fileDescriptor_8f4eb1aecc50742a) }

var fileDescriptor_8f4eb1aecc50742a = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0xcf, 0x6b, 0x14, 0x31,
	0x14, 0x66, 0xda, 0x75, 0x77, 0x9b, 0xda, 0xad, 0x4d, 0x0b, 0x9d, 0x5e, 0x6c, 0x19, 0x14, 0x16,
	0x84, 0x41, 0xb6, 0xb6, 0x77, 0x7f, 0x81, 0xc5, 0xa2, 0x30, 0x82, 0xe0, 0xe9, 0x91, 0x99, 0x64,
	0x77, 0x03, 0x99, 0x64, 0x48, 0xd2, 0x6d, 0x0f, 0xfe, 0x5d, 0x1e, 0xfc, 0x17, 0x3c, 0xf8, 0x1f,
	0x88, 0x37, 0x4f, 0xfe, 0x0f, 0x92, 0x97, 0x99, 0x55, 0xb4, 0x2b, 0x5e, 0xeb, 0x65, 0xd9, 0x7c,
	0xdf, 0x37, 0xef, 0xbd, 0xef, 0x7d, 0x93, 0x21, 0x87, 0x6a, 0x32, 0x95, 0x25, 0x88, 0x45, 0xa3,
	0xa1, 0xae, 0x98, 0xf3, 0x60, 0x94, 0x00, 0xa9, 0xa7, 0x26, 0x6f, 0xac, 0xf1, 0x86, 0x7e, 0x4c,
	0x2a, 0xe9, 0x2a, 0x03, 0xd2, 0x38, 0xb8, 0xb2, 0xa0, 0x26, 0x41, 0x69, 0x1a, 0x61, 0xf3, 0xf8,
	0x77, 0x6a, 0xec, 0x25, 0xb3, 0x5c, 0xea, 0x59, 0xae, 0x0d, 0x17, 0x0e, 0x7f, 0xf3, 0x5f, 0xaa,
	0x4a, 0x5d, 0x29, 0xa8, 0x01, 0x6b, 0xcf, 0x99, 0xe5, 0x97, 0xcc, 0x0a, 0x10, 0x33, 0x2b, 0x9c,
	0x13, 0xee, 0x1f, 0x85, 0xab, 0x64, 0x46, 0x89, 0xbf, 0x71, 0xd9, 0x7b, 0x72, 0x77, 0xa5, 0x3f,
	0x78, 0xf9, 0xfc, 0xdd, 0x1b, 0xba, 0x4f, 0x06, 0x61, 0x60, 0x90, 0x3c, 0x4d, 0x8e, 0x92, 0xf1,
	0x46, 0xd1, 0x0f, 0xc7, 0x33, 0x1e, 0x88, 0x92, 0x83, 0x66, 0xb5, 0x48, 0xd7, 0x22, 0x51, 0xf2,
	0x57, 0xac, 0x16, 0x74, 0x4c, 0xee, 0x68, 0x71, 0xe5, 0x61, 0x6e, 0x1a, 0x60, 0x9c, 0x87, 0xf1,
	0xd2, 0x75, 0x54, 0x8c, 0x02, 0xfe, 0xc2, 0x34, 0x8f, 0x23, 0x9a, 0x55, 0x64, 0x37, 0xee, 0x28,
	0x76, 0x37, 0x12, 0xfb, 0xd2, 0x07, 0x84, 0x2e, 0x0b, 0xc8, 0x66, 0x71, 0x8a, 0x55, 0xda, 0xee,
	0xdb, 0x6d, 0x89, 0xb3, 0x66, 0x71, 0x1a, 0xca, 0xd0, 0x43, 0xb2, 0x19, 0xc7, 0x56, 0xac, 0x14,
	0x0a, 0x47, 0xd9, 0x2a, 0x08, 0x42, 0xe7, 0x01, 0xc9, 0x76, 0xc8, 0x76, 0xb4, 0x58, 0x32, 0x17,
	0x8d, 0x65, 0x5f, 0xd6, 0x3b, 0x4c, 0x63, 0x8f, 0xd0, 0xf4, 0x6b, 0x42, 0x7a, 0x41, 0x81, 0x7d,
	0x36, 0x27, 0x9f, 0x92, 0xfc, 0xe6, 0xc6, 0x9a, 0xff, 0x66, 0xb8, 0x40, 0x67, 0xd7, 0x06, 0xb3,
	0x76, 0x5d, 0x30, 0xf4, 0x84, 0xec, 0xff, 0x4c, 0x40, 0x7b, 0x61, 0x35, 0x53, 0xed, 0x82, 0xd7,
	0x71, 0xc1, 0x7b, 0x5d, 0x0c, 0x2d, 0x89, 0xab, 0xa6, 0x39, 0xd9, 0x6d, 0x14, 0x9b, 0x1a, 0x5b,
	0x7b, 0xe0, 0xcc, 0x33, 0x58, 0x30, 0x25, 0x79, 0xda, 0x3b, 0x4a, 0xc6, 0xc3, 0x62, 0xa7, 0xa3,
	0x9e, 0x31, 0xcf, 0xde, 0x06, 0x82, 0x3e, 0x24, 0x7b, 0x8d, 0x62, 0x3e, 0xa0, 0x51, 0xaf, 0x84,
	0x9e, 0xf9, 0x79, 0x7a, 0x0b, 0x7b, 0xd0, 0x8e, 0x0b, 0x0f, 0x9c, 0x23, 0x43, 0xef, 0x93, 0x51,
	0x35, 0x97, 0x8a, 0x5b, 0xa1, 0xa1, 0x32, 0x17, 0xda, 0xa7, 0x7d, 0xd4, 0x6e, 0x75, 0xe8, 0xd3,
	0x00, 0x66, 0x1f, 0x7a, 0xe4, 0x60, 0xe5, 0x7b, 0x4d, 0xbf, 0x27, 0x64, 0x63, 0x09, 0xa5, 0x13,
	0xcc, 0xfb, 0xf3, 0x0d, 0xcf, 0xfb, 0x8f, 0x5b, 0x54, 0x0c, 0xd1, 0xe2, 0x6b, 0x25, 0xe8, 0x01,
	0x19, 0x4a, 0x07, 0xa5, 0xb9, 0xd0, 0x3c, 0x3d, 0xc6, 0x2c, 0x06, 0xd2, 0x3d, 0x09, 0x47, 0x7a,
	0x8f, 0x8c, 0x96, 0x41, 0xc7, 0xb0, 0x1e, 0xa1, 0xe0, 0x76, 0x9b, 0x6f, 0xcc, 0xe9, 0x5b, 0x42,
	0x86, 0x9d, 0x2c, 0x3d, 0xf9, 0x6f, 0xee, 0xc7, 0xf2, 0xf2, 0x17, 0x83, 0xd6, 0x6e, 0xd9, 0xc7,
	0x6f, 0xfa, 0xf1, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x6f, 0xf5, 0xfd, 0xf6, 0x05, 0x00,
	0x00,
}
