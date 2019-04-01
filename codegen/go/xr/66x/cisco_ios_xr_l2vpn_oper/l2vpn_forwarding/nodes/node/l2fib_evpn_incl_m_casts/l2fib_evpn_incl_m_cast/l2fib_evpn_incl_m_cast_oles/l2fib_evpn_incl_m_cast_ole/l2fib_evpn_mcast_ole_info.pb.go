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

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_evpn_incl_m_casts_l2fib_evpn_incl_m_cast_l2fib_evpn_incl_m_cast_oles_l2fib_evpn_incl_m_cast_ole

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
	TunnelEndpointId     uint32   `protobuf:"varint,1,opt,name=tunnel_endpoint_id,json=tunnelEndpointId,proto3" json:"tunnel_endpoint_id,omitempty"`
	NextHopIpv6Addr      string   `protobuf:"bytes,2,opt,name=next_hop_ipv6_addr,json=nextHopIpv6Addr,proto3" json:"next_hop_ipv6_addr,omitempty"`
	McastLabel           uint32   `protobuf:"varint,3,opt,name=mcast_label,json=mcastLabel,proto3" json:"mcast_label,omitempty"`
	McastEncapsulation   uint32   `protobuf:"varint,4,opt,name=mcast_encapsulation,json=mcastEncapsulation,proto3" json:"mcast_encapsulation,omitempty"`
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

func (m *L2VpnEvpnMoiInfo) GetTunnelEndpointId() uint32 {
	if m != nil {
		return m.TunnelEndpointId
	}
	return 0
}

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

func (m *L2VpnEvpnMoiInfo) GetMcastEncapsulation() uint32 {
	if m != nil {
		return m.McastEncapsulation
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
	Base                  *L2FibBaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	NextHopAddress        string         `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	NextHopAddressV6      string         `protobuf:"bytes,3,opt,name=next_hop_address_v6,json=nextHopAddressV6,proto3" json:"next_hop_address_v6,omitempty"`
	NextHopInternalLabel  uint32         `protobuf:"varint,4,opt,name=next_hop_internal_label,json=nextHopInternalLabel,proto3" json:"next_hop_internal_label,omitempty"`
	EcdPlaformtDataValid  bool           `protobuf:"varint,5,opt,name=ecd_plaformt_data_valid,json=ecdPlaformtDataValid,proto3" json:"ecd_plaformt_data_valid,omitempty"`
	EcdPlatformDataLength uint32         `protobuf:"varint,6,opt,name=ecd_platform_data_length,json=ecdPlatformDataLength,proto3" json:"ecd_platform_data_length,omitempty"`
	ChildrenCount         uint32         `protobuf:"varint,7,opt,name=children_count,json=childrenCount,proto3" json:"children_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
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

func (m *L2FibNhopInfo) GetNextHopAddressV6() string {
	if m != nil {
		return m.NextHopAddressV6
	}
	return ""
}

func (m *L2FibNhopInfo) GetNextHopInternalLabel() uint32 {
	if m != nil {
		return m.NextHopInternalLabel
	}
	return 0
}

func (m *L2FibNhopInfo) GetEcdPlaformtDataValid() bool {
	if m != nil {
		return m.EcdPlaformtDataValid
	}
	return false
}

func (m *L2FibNhopInfo) GetEcdPlatformDataLength() uint32 {
	if m != nil {
		return m.EcdPlatformDataLength
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
	IsEtreeLeaf          bool              `protobuf:"varint,52,opt,name=is_etree_leaf,json=isEtreeLeaf,proto3" json:"is_etree_leaf,omitempty"`
	NextHopValid         bool              `protobuf:"varint,53,opt,name=next_hop_valid,json=nextHopValid,proto3" json:"next_hop_valid,omitempty"`
	NextHop              *L2FibNhopInfo    `protobuf:"bytes,54,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
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

func (m *L2FibEvpnMcastOleInfo) GetIsEtreeLeaf() bool {
	if m != nil {
		return m.IsEtreeLeaf
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
	proto.RegisterType((*L2FibEvpnMcastOleInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_casts.l2fib_evpn_incl_m_cast.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_evpn_mcast_ole_info_KEYS")
	proto.RegisterType((*L2VpnEvpnMoiInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_casts.l2fib_evpn_incl_m_cast.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2vpn_evpn_moi_info")
	proto.RegisterType((*L2FibBaseInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_casts.l2fib_evpn_incl_m_cast.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_base_info")
	proto.RegisterType((*L2FibNhopInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_casts.l2fib_evpn_incl_m_cast.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_nhop_info")
	proto.RegisterType((*L2FibEvpnMcastOleInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_casts.l2fib_evpn_incl_m_cast.l2fib_evpn_incl_m_cast_oles.l2fib_evpn_incl_m_cast_ole.l2fib_evpn_mcast_ole_info")
}

func init() { proto.RegisterFile("l2fib_evpn_mcast_ole_info.proto", fileDescriptor_8f4eb1aecc50742a) }

var fileDescriptor_8f4eb1aecc50742a = []byte{
	// 595 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0xb4, 0xfd, 0x9a, 0xd4, 0xfd, 0xd2, 0x06, 0xa7, 0x28, 0xee, 0x86, 0x56, 0x23, 0x90,
	0x22, 0x01, 0x83, 0x94, 0x92, 0xb0, 0xe6, 0x27, 0x12, 0x11, 0x11, 0xa0, 0x41, 0xaa, 0xc4, 0xca,
	0xf2, 0x8c, 0xef, 0x34, 0x96, 0x1c, 0x7b, 0x34, 0x76, 0x42, 0x17, 0x3c, 0x08, 0xbc, 0x02, 0x0b,
	0x5e, 0x01, 0x89, 0x05, 0x2b, 0x9e, 0x88, 0x15, 0xb2, 0x3d, 0x09, 0xa5, 0x34, 0xac, 0xcb, 0x26,
	0xca, 0x9c, 0x73, 0x3c, 0xe7, 0xde, 0x7b, 0x7c, 0x35, 0xe8, 0x48, 0xf6, 0x0b, 0x91, 0x51, 0x58,
	0x94, 0x8a, 0xce, 0x72, 0x66, 0x2c, 0xd5, 0x12, 0xa8, 0x50, 0x85, 0x4e, 0xca, 0x4a, 0x5b, 0x8d,
	0x3f, 0x46, 0xb9, 0x30, 0xb9, 0xa6, 0x42, 0x1b, 0x7a, 0x5e, 0x51, 0xd9, 0x77, 0x4a, 0x5d, 0x42,
	0x95, 0x84, 0xbf, 0x85, 0xae, 0xde, 0xb1, 0x8a, 0x0b, 0x75, 0x96, 0x28, 0xcd, 0xc1, 0xf8, 0xdf,
	0xe4, 0xc2, 0x5b, 0x85, 0xca, 0x25, 0x9d, 0x51, 0xf7, 0x6e, 0xb3, 0x06, 0x5f, 0x03, 0xbb, 0x52,
	0xcc, 0x5f, 0xb8, 0xf8, 0x3d, 0xba, 0xb5, 0xb6, 0x7c, 0xfa, 0x62, 0xf4, 0xf6, 0x0d, 0xee, 0xa2,
	0x86, 0xab, 0x87, 0x0a, 0x4e, 0xa2, 0xe3, 0xa8, 0xb7, 0x93, 0x6e, 0xbb, 0xc7, 0x31, 0x77, 0x44,
	0xc6, 0xa9, 0x62, 0x33, 0x20, 0x1b, 0x81, 0xc8, 0xf8, 0x4b, 0x36, 0x03, 0xdc, 0x43, 0x6d, 0x05,
	0xe7, 0x96, 0x4e, 0x75, 0x49, 0x19, 0xe7, 0x15, 0x18, 0x43, 0x36, 0xbd, 0x62, 0xcf, 0xe1, 0xcf,
	0x75, 0xf9, 0x38, 0xa0, 0xf1, 0xd7, 0x08, 0x75, 0xc2, 0x0c, 0x82, 0xbd, 0x16, 0xde, 0x18, 0xdf,
	0x43, 0xd8, 0xce, 0x95, 0x02, 0x49, 0x41, 0xf1, 0x52, 0x0b, 0x65, 0x97, 0xf6, 0xad, 0xb4, 0x1d,
	0x98, 0x51, 0x4d, 0x8c, 0x39, 0xbe, 0x8b, 0xf0, 0xca, 0x4f, 0x94, 0x8b, 0xa1, 0x37, 0xad, 0x6b,
	0xda, 0xaf, 0x1d, 0xc7, 0xe5, 0x62, 0xe8, 0x5c, 0xf1, 0x11, 0xda, 0x0d, 0x5d, 0x4a, 0x96, 0x81,
	0xf4, 0x75, 0xb5, 0x52, 0xe4, 0xa1, 0x89, 0x43, 0xf0, 0x03, 0xd4, 0x09, 0x02, 0x50, 0x39, 0x2b,
	0xcd, 0x5c, 0x32, 0x2b, 0xb4, 0x22, 0x5b, 0x5e, 0x88, 0x3d, 0x35, 0xba, 0xc8, 0xc4, 0x37, 0xd0,
	0x7e, 0x18, 0x61, 0xc6, 0x4c, 0x18, 0x5c, 0xfc, 0x63, 0x73, 0x89, 0x29, 0x5f, 0x94, 0xeb, 0xe9,
	0x4b, 0x84, 0xb6, 0x9c, 0xc2, 0xb7, 0xb1, 0xdb, 0xff, 0x14, 0x25, 0xd7, 0xf6, 0x56, 0x24, 0x97,
	0xfa, 0x49, 0x7d, 0xe1, 0x57, 0xe6, 0xba, 0x71, 0x55, 0xae, 0xf8, 0x3e, 0xea, 0x5c, 0x56, 0xd2,
	0xc5, 0xb0, 0xbe, 0x04, 0xed, 0xdf, 0xc5, 0xa7, 0x43, 0x3c, 0x40, 0xdd, 0x5f, 0x01, 0x2a, 0x0b,
	0x95, 0x62, 0xb2, 0xce, 0x27, 0x8c, 0xfd, 0x60, 0x99, 0x62, 0x4d, 0x86, 0xa4, 0x06, 0xa8, 0x0b,
	0x39, 0xa7, 0xa5, 0x64, 0x85, 0xae, 0x66, 0x96, 0x72, 0x66, 0x19, 0x5d, 0x30, 0x29, 0x38, 0xf9,
	0xef, 0x38, 0xea, 0x35, 0xd3, 0x03, 0xc8, 0xf9, 0xeb, 0x9a, 0x7d, 0xc6, 0x2c, 0x3b, 0x75, 0x1c,
	0x7e, 0x84, 0x48, 0x7d, 0xcc, 0x3a, 0x26, 0x1c, 0x93, 0xa0, 0xce, 0xec, 0x94, 0x6c, 0x7b, 0xbb,
	0x9b, 0xe1, 0x9c, 0xa7, 0xdd, 0xb9, 0x89, 0x27, 0xf1, 0x1d, 0xb4, 0x97, 0x4f, 0x85, 0xe4, 0x15,
	0x28, 0x9a, 0xeb, 0xb9, 0xb2, 0xa4, 0xe1, 0xe5, 0xad, 0x25, 0xfa, 0xd4, 0x81, 0xf1, 0x87, 0x2d,
	0x74, 0xb8, 0x76, 0xa7, 0xf0, 0xf7, 0x08, 0xed, 0xac, 0x20, 0xd2, 0xf7, 0x77, 0xe1, 0xf3, 0xf5,
	0xbe, 0x0b, 0x7f, 0xec, 0x67, 0xda, 0xf4, 0x1d, 0xbc, 0x92, 0x80, 0x0f, 0x51, 0x53, 0x18, 0x9a,
	0xe9, 0xb9, 0xe2, 0xe4, 0xc4, 0x0f, 0xbd, 0x21, 0xcc, 0x13, 0xf7, 0x88, 0x63, 0xd4, 0x12, 0x86,
	0x82, 0xad, 0x00, 0xa8, 0x04, 0x56, 0x90, 0x87, 0x9e, 0xdf, 0x15, 0x66, 0xe4, 0xb0, 0x09, 0xb0,
	0x02, 0xdf, 0x46, 0x7b, 0xab, 0xe4, 0x43, 0x72, 0x03, 0x2f, 0xfa, 0xbf, 0x0e, 0x3c, 0x24, 0xf6,
	0x2d, 0x42, 0xcd, 0xa5, 0x8c, 0x0c, 0xff, 0x95, 0xf5, 0x59, 0xad, 0x7e, 0xda, 0xa8, 0xbb, 0xc9,
	0xb6, 0xfd, 0x07, 0xe1, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x5e, 0x1d, 0xfb, 0x33,
	0x06, 0x00, 0x00,
}
