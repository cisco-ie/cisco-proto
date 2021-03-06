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

package cisco_ios_xr_evpn_oper_evpn_active_evi_detail_evi_children_neighbors_neighbor

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
	Evi                  uint32   `protobuf:"varint,1,opt,name=evi,proto3" json:"evi,omitempty"`
	Encapsulation        uint32   `protobuf:"varint,2,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
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
	proto.RegisterType((*L2VpnEvpnNeighbor_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnNeighbor)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.neighbors.neighbor.l2vpn_evpn_neighbor")
}

func init() { proto.RegisterFile("l2vpn_evpn_neighbor.proto", fileDescriptor_78f930afb239d68c) }

var fileDescriptor_78f930afb239d68c = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x89, 0x2d, 0x55, 0x5f, 0x1d, 0x5a, 0x22, 0x68, 0x74, 0x63, 0x29, 0x22, 0x75, 0x93,
	0x45, 0x7b, 0x06, 0x17, 0x45, 0x74, 0x51, 0x41, 0x74, 0x15, 0x32, 0x33, 0x0f, 0x1b, 0x98, 0x49,
	0x42, 0x92, 0x0e, 0xed, 0x15, 0x3c, 0x81, 0xa7, 0xf1, 0x6c, 0x32, 0x71, 0xa6, 0x58, 0x98, 0xa5,
	0x9b, 0xf0, 0xe7, 0x7b, 0xc9, 0xfb, 0x7f, 0x5e, 0x02, 0x57, 0xc5, 0xbc, 0xb2, 0x5a, 0x60, 0xbd,
	0x68, 0x54, 0x1f, 0xeb, 0xd4, 0x38, 0x6e, 0x9d, 0x09, 0x86, 0x3e, 0x65, 0xca, 0x67, 0x46, 0x28,
	0xe3, 0xc5, 0xd6, 0xfd, 0x9e, 0x30, 0x16, 0x1d, 0xaf, 0x15, 0x97, 0x59, 0x50, 0x15, 0x72, 0xac,
	0x94, 0xc8, 0x31, 0x48, 0x55, 0x44, 0x99, 0xad, 0x55, 0x91, 0x3b, 0xd4, 0xbc, 0x6d, 0xe5, 0xf7,
	0x6a, 0xea, 0x81, 0x75, 0x78, 0x89, 0xc7, 0x87, 0xf7, 0x17, 0x3a, 0x86, 0x1e, 0x56, 0x8a, 0x91,
	0x09, 0x99, 0x25, 0xab, 0x5a, 0xd2, 0x5b, 0x48, 0x50, 0x67, 0xd2, 0xfa, 0x4d, 0x21, 0x83, 0x32,
	0x9a, 0x1d, 0xc5, 0xda, 0x21, 0xa4, 0x37, 0x30, 0xdc, 0x37, 0x52, 0x96, 0xf5, 0x26, 0x64, 0x76,
	0xba, 0x82, 0x16, 0x2d, 0xed, 0xf4, 0x8b, 0xc0, 0xc5, 0x1f, 0xd7, 0x3a, 0xa3, 0xdf, 0x94, 0xa5,
	0x74, 0x3b, 0x7a, 0x07, 0x23, 0x0c, 0x6b, 0x74, 0x1a, 0x83, 0xa8, 0x6b, 0x2a, 0x6f, 0xfc, 0x93,
	0x16, 0xbf, 0x5a, 0xbd, 0xcc, 0xe9, 0x3d, 0x8c, 0x0f, 0x4c, 0xc5, 0xd6, 0x35, 0x61, 0x46, 0x07,
	0xfc, 0xcd, 0xd1, 0x4b, 0x38, 0x4e, 0x73, 0xa1, 0x65, 0x89, 0x4d, 0x94, 0x41, 0x9a, 0x3f, 0xcb,
	0x12, 0x29, 0x85, 0x7e, 0xd8, 0x59, 0x64, 0xfd, 0x48, 0xa3, 0x9e, 0x7e, 0x13, 0x38, 0xef, 0x18,
	0x08, 0xfd, 0x24, 0x90, 0x44, 0xa2, 0xb4, 0x0f, 0x52, 0x67, 0xc8, 0xe6, 0x13, 0x32, 0x1b, 0xce,
	0x91, 0xff, 0xeb, 0x7b, 0xf0, 0xee, 0xb1, 0xac, 0xce, 0x6a, 0xb2, 0x6c, 0xac, 0xe9, 0x35, 0x9c,
	0xb4, 0x17, 0xd8, 0x22, 0x86, 0xdf, 0xef, 0xd3, 0x41, 0xfc, 0x26, 0x8b, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf8, 0x3c, 0xd5, 0x46, 0x43, 0x02, 0x00, 0x00,
}
