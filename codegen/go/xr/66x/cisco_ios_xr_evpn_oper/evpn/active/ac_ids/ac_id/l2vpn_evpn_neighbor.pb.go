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

package cisco_ios_xr_evpn_oper_evpn_active_ac_ids_ac_id

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
	AcId                 uint32   `protobuf:"varint,2,opt,name=ac_id,json=acId,proto3" json:"ac_id,omitempty"`
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
	proto.RegisterType((*L2VpnEvpnNeighbor_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ac_ids.ac_id.l2vpn_evpn_neighbor_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ac_ids.ac_id.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnNeighbor)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ac_ids.ac_id.l2vpn_evpn_neighbor")
}

func init() { proto.RegisterFile("l2vpn_evpn_neighbor.proto", fileDescriptor_78f930afb239d68c) }

var fileDescriptor_78f930afb239d68c = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0xbf, 0xfd, 0x6b, 0x7b, 0x75, 0x68, 0x49, 0x41, 0xa3, 0x6e, 0x86, 0x2e, 0x64,
	0xdc, 0x8c, 0xd0, 0xfa, 0x02, 0x2e, 0x8a, 0x14, 0xc1, 0xc5, 0x08, 0xa2, 0xab, 0x90, 0xc9, 0x5c,
	0x6d, 0x60, 0x26, 0x19, 0x92, 0x74, 0x68, 0xdf, 0xc4, 0x87, 0xf1, 0xe1, 0x24, 0x69, 0x2b, 0x0a,
	0xdd, 0xb8, 0x49, 0x4e, 0xbe, 0xe4, 0x9e, 0x9b, 0x93, 0xc0, 0x79, 0x35, 0x6d, 0x1b, 0xc5, 0xd0,
	0x0f, 0x0a, 0xe5, 0xfb, 0xb2, 0xd0, 0x26, 0x6b, 0x8c, 0x76, 0x9a, 0xdc, 0x08, 0x69, 0x85, 0x66,
	0x52, 0x5b, 0xb6, 0x36, 0xdb, 0x13, 0xba, 0x41, 0x93, 0x79, 0x95, 0x71, 0xe1, 0x64, 0x8b, 0x19,
	0x17, 0x4c, 0x96, 0x76, 0x3b, 0x4d, 0xee, 0x80, 0x1e, 0x70, 0x63, 0x0f, 0xf3, 0xd7, 0x27, 0x32,
	0x82, 0x0e, 0xb6, 0x92, 0x46, 0x49, 0x94, 0xc6, 0xb9, 0x97, 0x64, 0x0c, 0xff, 0x43, 0x19, 0xfd,
	0x17, 0x58, 0x97, 0x8b, 0x45, 0x39, 0xf9, 0x88, 0xe0, 0xf4, 0x87, 0x07, 0xb6, 0x92, 0xd9, 0x55,
	0x5d, 0x73, 0xb3, 0x21, 0x57, 0x30, 0x44, 0xb7, 0x44, 0xa3, 0xd0, 0x31, 0xbf, 0x27, 0xcb, 0x9d,
	0x5b, 0xbc, 0xc7, 0xcf, 0x8d, 0x5a, 0x94, 0xe4, 0x1a, 0x46, 0xa8, 0x04, 0x6f, 0xec, 0xaa, 0xe2,
	0x4e, 0x6a, 0xc5, 0xd6, 0x66, 0xd7, 0x62, 0xf8, 0x8b, 0xbf, 0x18, 0x72, 0x06, 0x47, 0x45, 0xc9,
	0x14, 0xaf, 0x91, 0x76, 0x92, 0x28, 0x1d, 0xe4, 0xbd, 0xa2, 0x7c, 0xe4, 0x35, 0x12, 0x02, 0x5d,
	0xb7, 0x69, 0x90, 0x76, 0x03, 0x0d, 0x7a, 0xf2, 0x19, 0xc1, 0xf8, 0x40, 0x3c, 0x52, 0x41, 0x1c,
	0x80, 0x54, 0xd6, 0x71, 0x25, 0x90, 0x4e, 0x93, 0x28, 0x3d, 0x9e, 0xde, 0x67, 0x7f, 0x7c, 0xbe,
	0xec, 0x70, 0xee, 0xfc, 0xc4, 0x93, 0xc5, 0xce, 0x9c, 0x5c, 0x40, 0x7f, 0xdf, 0x99, 0xce, 0xc2,
	0xed, 0xbe, 0xd7, 0xe4, 0x12, 0x06, 0x15, 0xf2, 0x37, 0x5f, 0x6d, 0xe9, 0x6d, 0xd2, 0x49, 0xe3,
	0xbc, 0xef, 0xc1, 0xbc, 0x95, 0xb6, 0xe8, 0x85, 0x4f, 0x9d, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x11, 0xfa, 0x41, 0xb1, 0xf1, 0x01, 0x00, 0x00,
}
