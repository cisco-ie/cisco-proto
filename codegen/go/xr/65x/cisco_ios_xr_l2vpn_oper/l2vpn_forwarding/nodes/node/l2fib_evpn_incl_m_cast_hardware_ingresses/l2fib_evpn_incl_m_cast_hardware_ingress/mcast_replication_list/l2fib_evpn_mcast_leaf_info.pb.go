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
// source: l2fib_evpn_mcast_leaf_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_evpn_incl_m_cast_hardware_ingresses_l2fib_evpn_incl_m_cast_hardware_ingress_mcast_replication_list

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

type L2FibEvpnMcastLeafInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	BdName               string   `protobuf:"bytes,2,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibEvpnMcastLeafInfo_KEYS) Reset()         { *m = L2FibEvpnMcastLeafInfo_KEYS{} }
func (m *L2FibEvpnMcastLeafInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnMcastLeafInfo_KEYS) ProtoMessage()    {}
func (*L2FibEvpnMcastLeafInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3801fdb27ef638, []int{0}
}

func (m *L2FibEvpnMcastLeafInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibEvpnMcastLeafInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnMcastLeafInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS.Merge(m, src)
}
func (m *L2FibEvpnMcastLeafInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS.Size(m)
}
func (m *L2FibEvpnMcastLeafInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnMcastLeafInfo_KEYS proto.InternalMessageInfo

func (m *L2FibEvpnMcastLeafInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibEvpnMcastLeafInfo_KEYS) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

type L2FibEvpnMcastLeafInfo struct {
	BridgeDomainName     string   `protobuf:"bytes,50,opt,name=bridge_domain_name,json=bridgeDomainName,proto3" json:"bridge_domain_name,omitempty"`
	BridgeDomainId       uint32   `protobuf:"varint,51,opt,name=bridge_domain_id,json=bridgeDomainId,proto3" json:"bridge_domain_id,omitempty"`
	Xcid                 uint32   `protobuf:"varint,52,opt,name=xcid,proto3" json:"xcid,omitempty"`
	IsBound              bool     `protobuf:"varint,53,opt,name=is_bound,json=isBound,proto3" json:"is_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibEvpnMcastLeafInfo) Reset()         { *m = L2FibEvpnMcastLeafInfo{} }
func (m *L2FibEvpnMcastLeafInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnMcastLeafInfo) ProtoMessage()    {}
func (*L2FibEvpnMcastLeafInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3801fdb27ef638, []int{1}
}

func (m *L2FibEvpnMcastLeafInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo.Unmarshal(m, b)
}
func (m *L2FibEvpnMcastLeafInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnMcastLeafInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnMcastLeafInfo.Merge(m, src)
}
func (m *L2FibEvpnMcastLeafInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnMcastLeafInfo.Size(m)
}
func (m *L2FibEvpnMcastLeafInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnMcastLeafInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnMcastLeafInfo proto.InternalMessageInfo

func (m *L2FibEvpnMcastLeafInfo) GetBridgeDomainName() string {
	if m != nil {
		return m.BridgeDomainName
	}
	return ""
}

func (m *L2FibEvpnMcastLeafInfo) GetBridgeDomainId() uint32 {
	if m != nil {
		return m.BridgeDomainId
	}
	return 0
}

func (m *L2FibEvpnMcastLeafInfo) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

func (m *L2FibEvpnMcastLeafInfo) GetIsBound() bool {
	if m != nil {
		return m.IsBound
	}
	return false
}

func init() {
	proto.RegisterType((*L2FibEvpnMcastLeafInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_ingresses.l2fib_evpn_incl_m_cast_hardware_ingress.mcast_replication_list.l2fib_evpn_mcast_leaf_info_KEYS")
	proto.RegisterType((*L2FibEvpnMcastLeafInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_evpn_incl_m_cast_hardware_ingresses.l2fib_evpn_incl_m_cast_hardware_ingress.mcast_replication_list.l2fib_evpn_mcast_leaf_info")
}

func init() { proto.RegisterFile("l2fib_evpn_mcast_leaf_info.proto", fileDescriptor_9b3801fdb27ef638) }

var fileDescriptor_9b3801fdb27ef638 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x03, 0x31,
	0x14, 0xc7, 0x89, 0x48, 0x5b, 0x03, 0x4a, 0xc9, 0xe2, 0xe9, 0xe2, 0xd1, 0xe9, 0x06, 0xb9, 0xa1,
	0xd5, 0x2f, 0x20, 0x3a, 0x14, 0xc1, 0xe1, 0x3a, 0x39, 0x3d, 0x92, 0xbc, 0xdc, 0xf9, 0x20, 0x97,
	0x1c, 0xc9, 0x69, 0xfb, 0x61, 0xdc, 0xfc, 0xa2, 0x92, 0xdc, 0x52, 0x87, 0x82, 0x4b, 0x78, 0x8f,
	0xdf, 0x9f, 0x1f, 0xff, 0xf0, 0x78, 0x69, 0xd7, 0x2d, 0x29, 0x30, 0x5f, 0x83, 0x83, 0x5e, 0xcb,
	0x38, 0x82, 0x35, 0xb2, 0x05, 0x72, 0xad, 0xaf, 0x87, 0xe0, 0x47, 0x2f, 0xbe, 0x99, 0xa6, 0xa8,
	0x3d, 0x90, 0x8f, 0x70, 0x08, 0x60, 0xd7, 0x29, 0xea, 0x07, 0x13, 0xea, 0x69, 0x6c, 0x7d, 0xd8,
	0xcb, 0x80, 0xe4, 0xba, 0xda, 0x79, 0x34, 0x31, 0xbf, 0xf5, 0x91, 0x96, 0x9c, 0xb6, 0xd0, 0x43,
	0x96, 0x7f, 0xc8, 0x80, 0x7b, 0x19, 0x0c, 0x90, 0xeb, 0x82, 0x89, 0xd1, 0xc4, 0xff, 0x26, 0xeb,
	0xa9, 0x5d, 0x30, 0x83, 0x25, 0x2d, 0x47, 0xf2, 0x0e, 0x2c, 0xc5, 0x71, 0xb5, 0xe3, 0x77, 0xa7,
	0xbf, 0x00, 0xaf, 0x2f, 0xef, 0x3b, 0x71, 0xcd, 0xe7, 0xa9, 0x13, 0x10, 0x16, 0xac, 0x64, 0xd5,
	0x45, 0x33, 0x4b, 0xeb, 0x16, 0x13, 0x50, 0x08, 0x4e, 0xf6, 0xa6, 0x38, 0x9b, 0x80, 0xc2, 0x37,
	0xd9, 0x9b, 0xd5, 0x0f, 0xe3, 0xb7, 0xa7, 0xad, 0xe2, 0x9e, 0x0b, 0x15, 0x08, 0x3b, 0x03, 0xe8,
	0x7b, 0x49, 0x6e, 0x52, 0xac, 0xb3, 0x62, 0x39, 0x91, 0xe7, 0x0c, 0x92, 0x4c, 0x54, 0x7c, 0xf9,
	0x37, 0x4d, 0x58, 0x6c, 0x4a, 0x56, 0x5d, 0x36, 0x57, 0xc7, 0xd9, 0x2d, 0x0a, 0xc1, 0xcf, 0x0f,
	0x9a, 0xb0, 0x78, 0xc8, 0x34, 0xcf, 0xe2, 0x86, 0x2f, 0x28, 0x82, 0xf2, 0x9f, 0x0e, 0x8b, 0xc7,
	0x92, 0x55, 0x8b, 0x66, 0x4e, 0xf1, 0x29, 0xad, 0x6a, 0x96, 0x0f, 0xb4, 0xf9, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0x20, 0x4b, 0x0c, 0xc4, 0x01, 0x00, 0x00,
}
