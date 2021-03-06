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
	Encapsulation        uint32   `protobuf:"varint,3,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	NeighborIp           string   `protobuf:"bytes,4,opt,name=neighbor_ip,json=neighborIp,proto3" json:"neighbor_ip,omitempty"`
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

func (m *L2VpnEvpnNeighbor_KEYS) GetEncapsulation() uint32 {
	if m != nil {
		return m.Encapsulation
	}
	return 0
}

func (m *L2VpnEvpnNeighbor_KEYS) GetNeighborIp() string {
	if m != nil {
		return m.NeighborIp
	}
	return ""
}

type L2VpnEvpnEviSummary struct {
	EthernetVpnId        uint32   `protobuf:"varint,1,opt,name=ethernet_vpn_id,json=ethernetVpnId,proto3" json:"ethernet_vpn_id,omitempty"`
	EncapsulationXr      uint32   `protobuf:"varint,2,opt,name=encapsulation_xr,json=encapsulationXr,proto3" json:"encapsulation_xr,omitempty"`
	BdName               string   `protobuf:"bytes,3,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEviSummary) Reset()         { *m = L2VpnEvpnEviSummary{} }
func (m *L2VpnEvpnEviSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviSummary) ProtoMessage()    {}
func (*L2VpnEvpnEviSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f930afb239d68c, []int{1}
}

func (m *L2VpnEvpnEviSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviSummary.Merge(m, src)
}
func (m *L2VpnEvpnEviSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Size(m)
}
func (m *L2VpnEvpnEviSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviSummary proto.InternalMessageInfo

func (m *L2VpnEvpnEviSummary) GetEthernetVpnId() uint32 {
	if m != nil {
		return m.EthernetVpnId
	}
	return 0
}

func (m *L2VpnEvpnEviSummary) GetEncapsulationXr() uint32 {
	if m != nil {
		return m.EncapsulationXr
	}
	return 0
}

func (m *L2VpnEvpnEviSummary) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

func (m *L2VpnEvpnEviSummary) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type L2VpnEvpnNeighbor struct {
	EvpnInstance         *L2VpnEvpnEviSummary `protobuf:"bytes,50,opt,name=evpn_instance,json=evpnInstance,proto3" json:"evpn_instance,omitempty"`
	Neighbor             string               `protobuf:"bytes,51,opt,name=neighbor,proto3" json:"neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *L2VpnEvpnNeighbor) Reset()         { *m = L2VpnEvpnNeighbor{} }
func (m *L2VpnEvpnNeighbor) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnNeighbor) ProtoMessage()    {}
func (*L2VpnEvpnNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_78f930afb239d68c, []int{2}
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

func (m *L2VpnEvpnNeighbor) GetEvpnInstance() *L2VpnEvpnEviSummary {
	if m != nil {
		return m.EvpnInstance
	}
	return nil
}

func (m *L2VpnEvpnNeighbor) GetNeighbor() string {
	if m != nil {
		return m.Neighbor
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnEvpnNeighbor_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnNeighbor)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor")
}

func init() { proto.RegisterFile("l2vpn_evpn_neighbor.proto", fileDescriptor_78f930afb239d68c) }

var fileDescriptor_78f930afb239d68c = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0x26, 0x7f, 0x4b, 0xff, 0xbf, 0xd3, 0xbf, 0xb4, 0x44, 0xd0, 0xd5, 0x8b, 0xa5, 0x88, 0xd4,
	0x4b, 0x0e, 0xed, 0x33, 0x78, 0x58, 0x04, 0xc1, 0x15, 0x44, 0x4f, 0x21, 0xbb, 0x19, 0x6c, 0x60,
	0x37, 0x09, 0x49, 0x5a, 0xda, 0x77, 0x10, 0xcf, 0x3e, 0x8f, 0x4f, 0x26, 0x89, 0xdb, 0x62, 0xa1,
	0x47, 0x2f, 0xc3, 0x37, 0xdf, 0x24, 0xf3, 0x7d, 0x33, 0x09, 0x9c, 0xd7, 0xf3, 0xb5, 0xd5, 0x1c,
	0x63, 0xd0, 0xa8, 0x5e, 0x97, 0xa5, 0x71, 0xcc, 0x3a, 0x13, 0x0c, 0x7d, 0xa8, 0x94, 0xaf, 0x0c,
	0x57, 0xc6, 0xf3, 0x8d, 0xfb, 0x3e, 0x61, 0x2c, 0x3a, 0x16, 0x11, 0xd3, 0x46, 0xa2, 0x4f, 0x91,
	0xe1, 0x5a, 0x71, 0x89, 0x41, 0xa8, 0x3a, 0xc1, 0x6a, 0xa9, 0x6a, 0xe9, 0x50, 0xb3, 0x5d, 0x3b,
	0xbf, 0x47, 0xd3, 0x37, 0x02, 0xd9, 0x11, 0x41, 0x7e, 0x77, 0xfb, 0xf2, 0x48, 0xcf, 0xe0, 0x6f,
	0xec, 0xc7, 0x95, 0xcc, 0xc8, 0x84, 0xcc, 0xfa, 0x45, 0x2f, 0xa6, 0xb9, 0xa4, 0x63, 0xe8, 0xe0,
	0x5a, 0x65, 0x7f, 0x26, 0x64, 0x36, 0x2c, 0x22, 0xa4, 0x57, 0x30, 0x44, 0x5d, 0x09, 0xeb, 0x57,
	0xb5, 0x08, 0xca, 0xe8, 0xac, 0x93, 0x6a, 0x87, 0x24, 0xbd, 0x84, 0xc1, 0x5e, 0x41, 0xd9, 0xac,
	0x9b, 0x9a, 0xc2, 0x8e, 0xca, 0xed, 0xf4, 0x83, 0xc0, 0xe9, 0x0f, 0x3b, 0xd1, 0xbd, 0x5f, 0x35,
	0x8d, 0x70, 0x5b, 0x7a, 0x0d, 0x23, 0x0c, 0x4b, 0x74, 0x1a, 0x03, 0x8f, 0xb5, 0xd6, 0x54, 0xd4,
	0x68, 0xe9, 0x27, 0xab, 0x73, 0x49, 0x6f, 0x60, 0x7c, 0x20, 0xca, 0x37, 0xae, 0x35, 0x3a, 0x3a,
	0xe0, 0x9f, 0x5d, 0x9c, 0xaf, 0x94, 0x5c, 0x8b, 0x06, 0x93, 0xdd, 0x7e, 0xd1, 0x2b, 0xe5, 0xbd,
	0x68, 0x90, 0x52, 0xe8, 0x86, 0xad, 0xc5, 0xd6, 0x60, 0xc2, 0xd3, 0x4f, 0x02, 0x27, 0x47, 0x36,
	0x45, 0xdf, 0x09, 0x0c, 0x13, 0xa3, 0xb4, 0x0f, 0x42, 0x57, 0x98, 0xcd, 0x27, 0x64, 0x36, 0x98,
	0x2b, 0xf6, 0xeb, 0xaf, 0xc5, 0x8e, 0xaf, 0xa6, 0xf8, 0x1f, 0x99, 0xbc, 0x95, 0xa7, 0x17, 0xf0,
	0x6f, 0x77, 0x21, 0x5b, 0xa4, 0x01, 0xf6, 0x79, 0xd9, 0x4b, 0x1f, 0x69, 0xf1, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0x19, 0x6d, 0x78, 0x65, 0x02, 0x00, 0x00,
}
