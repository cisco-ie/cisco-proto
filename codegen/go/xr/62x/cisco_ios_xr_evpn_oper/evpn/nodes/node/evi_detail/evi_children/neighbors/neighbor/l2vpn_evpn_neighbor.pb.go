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
// source: l2vpn_evpn_neighbor.proto

package cisco_ios_xr_evpn_oper_evpn_nodes_node_evi_detail_evi_children_neighbors_neighbor

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

type L2VpnEvpnNeighbor_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Evi                  uint32   `protobuf:"varint,2,opt,name=evi,proto3" json:"evi,omitempty"`
	NeighborIp           string   `protobuf:"bytes,3,opt,name=neighbor_ip,json=neighborIp,proto3" json:"neighbor_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnNeighbor_KEYS) Reset()         { *m = L2VpnEvpnNeighbor_KEYS{} }
func (m *L2VpnEvpnNeighbor_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnNeighbor_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnNeighbor_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f930afb239d68c, []int{0}
}

func (m *L2VpnEvpnNeighbor_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnNeighbor_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnNeighbor_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnNeighbor_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnNeighbor_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnNeighbor_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnNeighbor_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnNeighbor_KEYS.Size(m)
}
func (m *L2VpnEvpnNeighbor_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnNeighbor_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnNeighbor_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnNeighbor_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnEvpnNeighbor_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnNeighbor_KEYS) GetNeighborIp() string {
	if m != nil {
		return m.NeighborIp
	}
	return ""
}

type L2VpnEvpnNeighbor struct {
	EviXr                uint32   `protobuf:"varint,50,opt,name=evi_xr,json=eviXr,proto3" json:"evi_xr,omitempty"`
	Neighbor             string   `protobuf:"bytes,51,opt,name=neighbor,proto3" json:"neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnNeighbor) Reset()         { *m = L2VpnEvpnNeighbor{} }
func (m *L2VpnEvpnNeighbor) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnNeighbor) ProtoMessage()    {}
func (*L2VpnEvpnNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f930afb239d68c, []int{1}
}

func (m *L2VpnEvpnNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnNeighbor.Unmarshal(m, b)
}
func (m *L2VpnEvpnNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnNeighbor.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnNeighbor.Merge(m, src)
}
func (m *L2VpnEvpnNeighbor) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnNeighbor.Size(m)
}
func (m *L2VpnEvpnNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnNeighbor proto.InternalMessageInfo

func (m *L2VpnEvpnNeighbor) GetEviXr() uint32 {
	if m != nil {
		return m.EviXr
	}
	return 0
}

func (m *L2VpnEvpnNeighbor) GetNeighbor() string {
	if m != nil {
		return m.Neighbor
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnEvpnNeighbor_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor_KEYS")
	proto.RegisterType((*L2VpnEvpnNeighbor)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor")
}

func init() { proto.RegisterFile("l2vpn_evpn_neighbor.proto", fileDescriptor_78f930afb239d68c) }

var fileDescriptor_78f930afb239d68c = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x86, 0xa9, 0xc5, 0xaa, 0x27, 0x82, 0x44, 0xc4, 0xe8, 0x62, 0xe9, 0xd4, 0x29, 0x43, 0xfb,
	0x1b, 0x04, 0x8b, 0x93, 0x75, 0xd1, 0x29, 0xd8, 0xe6, 0xb4, 0x07, 0x25, 0x09, 0x49, 0x09, 0xfd,
	0xf9, 0x1f, 0xc9, 0x47, 0x3b, 0x75, 0x39, 0x9e, 0x0b, 0x6f, 0x9e, 0x37, 0x04, 0x9e, 0xe7, 0x26,
	0x58, 0x2d, 0x31, 0x0e, 0x8d, 0xf4, 0x3f, 0x0d, 0xc6, 0x09, 0xeb, 0xcc, 0x62, 0xd8, 0xe7, 0x48,
	0x7e, 0x34, 0x92, 0x8c, 0x97, 0xab, 0x3b, 0x27, 0x8c, 0x45, 0x27, 0x22, 0x09, 0x6d, 0x14, 0xfa,
	0x34, 0x05, 0x06, 0x92, 0x0a, 0x97, 0x5f, 0x9a, 0x13, 0x8e, 0x13, 0xcd, 0xca, 0xa1, 0x16, 0x9b,
	0xce, 0xef, 0x54, 0xfd, 0x01, 0x3f, 0xe8, 0x93, 0x1f, 0x6f, 0x3f, 0x5f, 0xec, 0x09, 0xae, 0xa2,
	0x4e, 0x92, 0xe2, 0x59, 0x99, 0xd5, 0x37, 0x7d, 0x11, 0xd7, 0x4e, 0xb1, 0x7b, 0xc8, 0x31, 0x10,
	0xbf, 0x28, 0xb3, 0xfa, 0xae, 0x8f, 0xc8, 0x5e, 0xe1, 0x76, 0xbf, 0x4b, 0x96, 0xe7, 0x29, 0x0e,
	0xdb, 0x51, 0x67, 0xab, 0x77, 0x78, 0x38, 0xe8, 0x61, 0x8f, 0x50, 0xc4, 0x17, 0xae, 0x8e, 0x37,
	0x49, 0x76, 0x89, 0x81, 0xbe, 0x1d, 0x7b, 0x81, 0xeb, 0x2d, 0xc2, 0xdb, 0xe4, 0xda, 0xf7, 0xa1,
	0x48, 0x7f, 0xd1, 0x9e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x84, 0xeb, 0xc7, 0xa3, 0x28, 0x01, 0x00,
	0x00,
}
