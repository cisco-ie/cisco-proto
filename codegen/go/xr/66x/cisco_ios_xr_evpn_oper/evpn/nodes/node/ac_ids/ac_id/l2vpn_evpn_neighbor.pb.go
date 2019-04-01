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

package cisco_ios_xr_evpn_oper_evpn_nodes_node_ac_ids_ac_id

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
	AcId                 uint32   `protobuf:"varint,3,opt,name=ac_id,json=acId,proto3" json:"ac_id,omitempty"`
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

func (m *L2VpnEvpnNeighbor_KEYS) GetAcId() uint32 {
	if m != nil {
		return m.AcId
	}
	return 0
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
	LeafEvis             []uint32             `protobuf:"varint,52,rep,packed,name=leaf_evis,json=leafEvis,proto3" json:"leaf_evis,omitempty"`
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

func (m *L2VpnEvpnNeighbor) GetLeafEvis() []uint32 {
	if m != nil {
		return m.LeafEvis
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnEvpnNeighbor_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ac_ids.ac_id.l2vpn_evpn_neighbor_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ac_ids.ac_id.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnNeighbor)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ac_ids.ac_id.l2vpn_evpn_neighbor")
}

func init() { proto.RegisterFile("l2vpn_evpn_neighbor.proto", fileDescriptor_78f930afb239d68c) }

var fileDescriptor_78f930afb239d68c = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x5a, 0x4a, 0x7b, 0x10, 0xb5, 0x72, 0x25, 0x30, 0xb0, 0x44, 0x1d, 0x50, 0x58,
	0x32, 0xb4, 0xbc, 0x42, 0x87, 0xa8, 0x12, 0x43, 0x90, 0x10, 0xb0, 0x58, 0x4e, 0x7c, 0x50, 0x4b,
	0x8d, 0x6d, 0xd9, 0x69, 0xd4, 0xbe, 0x09, 0x8f, 0xc3, 0xa3, 0x21, 0xbb, 0x2d, 0x02, 0xa9, 0x13,
	0xcb, 0xe5, 0xee, 0x8b, 0xef, 0xfe, 0xff, 0x6c, 0xb8, 0x5e, 0x4d, 0x5b, 0xa3, 0x18, 0xfa, 0xa0,
	0x50, 0x7e, 0x2c, 0x4b, 0x6d, 0x33, 0x63, 0x75, 0xa3, 0xc9, 0xac, 0x92, 0xae, 0xd2, 0x4c, 0x6a,
	0xc7, 0x36, 0x76, 0x77, 0x42, 0x1b, 0xb4, 0x99, 0xcf, 0x32, 0xa5, 0x05, 0xba, 0x10, 0x33, 0x5e,
	0x31, 0x29, 0xdc, 0xee, 0x33, 0x79, 0x03, 0x7a, 0x64, 0x22, 0x5b, 0xcc, 0x5f, 0x9f, 0xc8, 0x15,
	0x9c, 0xf9, 0x06, 0x26, 0x05, 0x8d, 0x92, 0x28, 0x1d, 0x14, 0x3d, 0x5f, 0xe6, 0x82, 0x8c, 0xa0,
	0x83, 0xad, 0xa4, 0x27, 0x49, 0x94, 0xc6, 0x85, 0x4f, 0xc9, 0x18, 0x4e, 0xc3, 0x3c, 0xda, 0x09,
	0xac, 0xcb, 0xab, 0x5c, 0x4c, 0x3e, 0x23, 0xb8, 0xfc, 0x35, 0x1c, 0x5b, 0xc9, 0xdc, 0xba, 0xae,
	0xb9, 0xdd, 0x92, 0x3b, 0x18, 0x62, 0xb3, 0x44, 0xab, 0xb0, 0x61, 0xfe, 0xdf, 0x5e, 0x22, 0x2e,
	0xe2, 0x03, 0x7e, 0x36, 0x2a, 0x17, 0xe4, 0x1e, 0x46, 0xa8, 0x2a, 0x6e, 0xdc, 0x7a, 0xc5, 0x1b,
	0xa9, 0x15, 0xdb, 0xd8, 0xbd, 0xec, 0xf0, 0x0f, 0x7f, 0xb1, 0xde, 0x6d, 0x29, 0x98, 0xe2, 0x35,
	0x06, 0x13, 0x83, 0xa2, 0x57, 0x8a, 0x47, 0x5e, 0x23, 0x21, 0xd0, 0x6d, 0xb6, 0x06, 0x69, 0x37,
	0xd0, 0x90, 0x4f, 0xbe, 0x22, 0x18, 0x1f, 0xd9, 0x9b, 0x18, 0x88, 0x03, 0x90, 0xca, 0x35, 0x5c,
	0x55, 0x48, 0xa7, 0x49, 0x94, 0x9e, 0x4f, 0x17, 0xd9, 0x3f, 0xee, 0x36, 0x3b, 0xbe, 0x7b, 0x71,
	0xe1, 0x49, 0xbe, 0x17, 0x20, 0x37, 0xd0, 0x3f, 0xa8, 0xd3, 0x59, 0x70, 0xf8, 0x53, 0x93, 0x5b,
	0x18, 0xac, 0x90, 0xbf, 0xfb, 0x6e, 0x47, 0x1f, 0x92, 0x4e, 0x1a, 0x17, 0x7d, 0x0f, 0xe6, 0xad,
	0x74, 0x65, 0x2f, 0xbc, 0xfa, 0xec, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x00, 0xef, 0xdc, 0x12,
	0x02, 0x00, 0x00,
}