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
// source: l2vpn_evpn_imcast.proto

package cisco_ios_xr_evpn_oper_evpn_active_evi_detail_evi_children_inclusive_multicasts_inclusive_multicast

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

type L2VpnEvpnImcast_KEYS struct {
	Evi                  uint32   `protobuf:"varint,1,opt,name=evi,proto3" json:"evi,omitempty"`
	Encapsulation        uint32   `protobuf:"varint,2,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	EthernetTag          uint32   `protobuf:"varint,3,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	OriginatingIp        string   `protobuf:"bytes,4,opt,name=originating_ip,json=originatingIp,proto3" json:"originating_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnImcast_KEYS) Reset()         { *m = L2VpnEvpnImcast_KEYS{} }
func (m *L2VpnEvpnImcast_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnImcast_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnImcast_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb00aa10138a6cf7, []int{0}
}

func (m *L2VpnEvpnImcast_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnImcast_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnImcast_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnImcast_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnImcast_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnImcast_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnImcast_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnImcast_KEYS.Size(m)
}
func (m *L2VpnEvpnImcast_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnImcast_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnImcast_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnImcast_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnImcast_KEYS) GetEncapsulation() uint32 {
	if m != nil {
		return m.Encapsulation
	}
	return 0
}

func (m *L2VpnEvpnImcast_KEYS) GetEthernetTag() uint32 {
	if m != nil {
		return m.EthernetTag
	}
	return 0
}

func (m *L2VpnEvpnImcast_KEYS) GetOriginatingIp() string {
	if m != nil {
		return m.OriginatingIp
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
	return fileDescriptor_eb00aa10138a6cf7, []int{1}
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

type L2VpnEvpnImcast struct {
	EvpnInstance         *L2VpnEvpnEviSummary `protobuf:"bytes,50,opt,name=evpn_instance,json=evpnInstance,proto3" json:"evpn_instance,omitempty"`
	EthernetTagXr        uint32               `protobuf:"varint,51,opt,name=ethernet_tag_xr,json=ethernetTagXr,proto3" json:"ethernet_tag_xr,omitempty"`
	OriginatingIpXr      string               `protobuf:"bytes,52,opt,name=originating_ip_xr,json=originatingIpXr,proto3" json:"originating_ip_xr,omitempty"`
	TunnelEndpointId     uint32               `protobuf:"varint,53,opt,name=tunnel_endpoint_id,json=tunnelEndpointId,proto3" json:"tunnel_endpoint_id,omitempty"`
	PmsiTunnelType       uint32               `protobuf:"varint,54,opt,name=pmsi_tunnel_type,json=pmsiTunnelType,proto3" json:"pmsi_tunnel_type,omitempty"`
	NextHop              string               `protobuf:"bytes,55,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	OutputLabel          uint32               `protobuf:"varint,56,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	IsLocalEntry         bool                 `protobuf:"varint,57,opt,name=is_local_entry,json=isLocalEntry,proto3" json:"is_local_entry,omitempty"`
	IsProxyEntry         bool                 `protobuf:"varint,58,opt,name=is_proxy_entry,json=isProxyEntry,proto3" json:"is_proxy_entry,omitempty"`
	SrtePolicy           string               `protobuf:"bytes,59,opt,name=srte_policy,json=srtePolicy,proto3" json:"srte_policy,omitempty"`
	EtreeLeaf            bool                 `protobuf:"varint,60,opt,name=etree_leaf,json=etreeLeaf,proto3" json:"etree_leaf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *L2VpnEvpnImcast) Reset()         { *m = L2VpnEvpnImcast{} }
func (m *L2VpnEvpnImcast) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnImcast) ProtoMessage()    {}
func (*L2VpnEvpnImcast) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb00aa10138a6cf7, []int{2}
}

func (m *L2VpnEvpnImcast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnImcast.Unmarshal(m, b)
}
func (m *L2VpnEvpnImcast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnImcast.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnImcast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnImcast.Merge(m, src)
}
func (m *L2VpnEvpnImcast) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnImcast.Size(m)
}
func (m *L2VpnEvpnImcast) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnImcast.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnImcast proto.InternalMessageInfo

func (m *L2VpnEvpnImcast) GetEvpnInstance() *L2VpnEvpnEviSummary {
	if m != nil {
		return m.EvpnInstance
	}
	return nil
}

func (m *L2VpnEvpnImcast) GetEthernetTagXr() uint32 {
	if m != nil {
		return m.EthernetTagXr
	}
	return 0
}

func (m *L2VpnEvpnImcast) GetOriginatingIpXr() string {
	if m != nil {
		return m.OriginatingIpXr
	}
	return ""
}

func (m *L2VpnEvpnImcast) GetTunnelEndpointId() uint32 {
	if m != nil {
		return m.TunnelEndpointId
	}
	return 0
}

func (m *L2VpnEvpnImcast) GetPmsiTunnelType() uint32 {
	if m != nil {
		return m.PmsiTunnelType
	}
	return 0
}

func (m *L2VpnEvpnImcast) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnEvpnImcast) GetOutputLabel() uint32 {
	if m != nil {
		return m.OutputLabel
	}
	return 0
}

func (m *L2VpnEvpnImcast) GetIsLocalEntry() bool {
	if m != nil {
		return m.IsLocalEntry
	}
	return false
}

func (m *L2VpnEvpnImcast) GetIsProxyEntry() bool {
	if m != nil {
		return m.IsProxyEntry
	}
	return false
}

func (m *L2VpnEvpnImcast) GetSrtePolicy() string {
	if m != nil {
		return m.SrtePolicy
	}
	return ""
}

func (m *L2VpnEvpnImcast) GetEtreeLeaf() bool {
	if m != nil {
		return m.EtreeLeaf
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnEvpnImcast_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnImcast)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast")
}

func init() { proto.RegisterFile("l2vpn_evpn_imcast.proto", fileDescriptor_eb00aa10138a6cf7) }

var fileDescriptor_eb00aa10138a6cf7 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x15, 0x56, 0x6d, 0xab, 0xfb, 0x77, 0x3e, 0xb0, 0x70, 0x40, 0x94, 0x6a, 0xa0, 0x82,
	0x50, 0x0e, 0x1d, 0xff, 0xe1, 0x5a, 0x89, 0x8a, 0x0a, 0x4d, 0xa5, 0x42, 0xe5, 0x64, 0xb9, 0xce,
	0xbb, 0xd4, 0x92, 0x63, 0x5b, 0xb6, 0x53, 0x35, 0xe2, 0x8b, 0x20, 0x4e, 0x7c, 0x00, 0x3e, 0x24,
	0xb2, 0x93, 0x22, 0xa2, 0xed, 0xbc, 0x4b, 0xf4, 0xfa, 0xf7, 0x3e, 0xb1, 0x9f, 0x3c, 0xaf, 0x83,
	0xce, 0xc5, 0x74, 0xa7, 0x25, 0x01, 0xff, 0xe0, 0x39, 0xa3, 0xd6, 0x25, 0xda, 0x28, 0xa7, 0x30,
	0x63, 0xdc, 0x32, 0x45, 0xb8, 0xb2, 0x64, 0x6f, 0xaa, 0xbe, 0xd2, 0x60, 0x12, 0x5f, 0x25, 0x94,
	0x39, 0xbe, 0x83, 0x04, 0x76, 0x9c, 0xa4, 0xe0, 0x28, 0x17, 0xa1, 0x64, 0x5b, 0x2e, 0x52, 0x03,
	0x32, 0xe1, 0x92, 0x89, 0xc2, 0xf2, 0x1d, 0x90, 0xbc, 0x10, 0x8e, 0xfb, 0x5d, 0xed, 0x6d, 0x70,
	0xfc, 0x2b, 0x42, 0xf7, 0x6f, 0x18, 0x20, 0x9f, 0x67, 0xdf, 0xbf, 0xe2, 0x21, 0x3a, 0x82, 0x1d,
	0x8f, 0xa3, 0x51, 0x34, 0xe9, 0x2d, 0x7d, 0x89, 0x2f, 0x50, 0x0f, 0x24, 0xa3, 0xda, 0x16, 0x82,
	0x3a, 0xae, 0x64, 0x7c, 0x2f, 0xf4, 0x9a, 0x10, 0x3f, 0x46, 0x5d, 0x70, 0x5b, 0x30, 0x12, 0x1c,
	0x71, 0x34, 0x8b, 0x8f, 0x82, 0xa8, 0x73, 0x60, 0x2b, 0x9a, 0xe1, 0x27, 0xa8, 0xaf, 0x0c, 0xcf,
	0xb8, 0xa4, 0x8e, 0xcb, 0x8c, 0x70, 0x1d, 0xb7, 0x46, 0xd1, 0xa4, 0xbd, 0xec, 0xfd, 0x47, 0xe7,
	0x7a, 0xfc, 0xb3, 0x69, 0xce, 0x7f, 0xa1, 0x2d, 0xf2, 0x9c, 0x9a, 0x12, 0x3f, 0x45, 0x83, 0x7f,
	0x87, 0x04, 0xe3, 0x69, 0x6d, 0xb4, 0x77, 0xc0, 0xdf, 0xb4, 0x9c, 0xa7, 0xf8, 0x19, 0x1a, 0x36,
	0xdc, 0x91, 0xbd, 0xa9, 0x5d, 0x0f, 0x1a, 0x7c, 0x6d, 0xf0, 0x39, 0x3a, 0xd9, 0xa4, 0x44, 0xd2,
	0x1c, 0x82, 0xe5, 0xf6, 0xf2, 0x78, 0x93, 0x7e, 0xa1, 0x39, 0x60, 0x8c, 0x5a, 0xae, 0xd4, 0x50,
	0x7b, 0x0c, 0xf5, 0xf8, 0x4f, 0x0b, 0x9d, 0xdd, 0xc8, 0x0d, 0xff, 0x8e, 0x50, 0xaf, 0x5a, 0x4b,
	0xeb, 0xa8, 0x64, 0x10, 0x4f, 0x47, 0xd1, 0xa4, 0x33, 0xfd, 0x91, 0xdc, 0xc1, 0x2c, 0x93, 0xdb,
	0xa3, 0x5a, 0x76, 0x3d, 0x99, 0xd7, 0x86, 0x1a, 0xc1, 0x39, 0x9a, 0xf9, 0x3c, 0x2e, 0x9b, 0xc1,
	0xad, 0x68, 0xb6, 0x36, 0xf8, 0x39, 0x3a, 0x6b, 0x8e, 0xc8, 0x2b, 0x5f, 0x86, 0x04, 0x06, 0x8d,
	0x29, 0xad, 0x0d, 0x7e, 0x81, 0xb0, 0x2b, 0xa4, 0x04, 0x41, 0x40, 0xa6, 0x5a, 0x71, 0xe9, 0xfc,
	0x3c, 0x5e, 0x85, 0x6d, 0x87, 0x55, 0x67, 0x56, 0x37, 0xe6, 0x29, 0x9e, 0xa0, 0xa1, 0xce, 0x2d,
	0x27, 0xf5, 0x2b, 0x21, 0xda, 0xd7, 0x41, 0xdb, 0xf7, 0x7c, 0x15, 0xf0, 0xaa, 0xd4, 0x80, 0x1f,
	0xa0, 0x53, 0x09, 0x7b, 0x47, 0xb6, 0x4a, 0xc7, 0x6f, 0xc2, 0xd1, 0x27, 0x7e, 0xfd, 0x49, 0x69,
	0x7f, 0xc9, 0x54, 0xe1, 0x74, 0xe1, 0x88, 0xa0, 0x1b, 0x10, 0xf1, 0xdb, 0xea, 0x92, 0x55, 0x6c,
	0xe1, 0x11, 0xbe, 0x40, 0x7d, 0x6e, 0x89, 0x50, 0x8c, 0x7a, 0x5f, 0xce, 0x94, 0xf1, 0xbb, 0x51,
	0x34, 0x39, 0x5d, 0x76, 0xb9, 0x5d, 0x78, 0x38, 0xf3, 0xac, 0x56, 0x69, 0xa3, 0xf6, 0x65, 0xad,
	0x7a, 0x7f, 0x50, 0x5d, 0x79, 0x58, 0xa9, 0x1e, 0xa1, 0x8e, 0x35, 0x0e, 0x88, 0x56, 0x82, 0xb3,
	0x32, 0xfe, 0x10, 0xcc, 0x20, 0x8f, 0xae, 0x02, 0xc1, 0x0f, 0x11, 0x02, 0x67, 0x00, 0x88, 0x00,
	0x7a, 0x1d, 0x7f, 0x0c, 0x5b, 0xb4, 0x03, 0x59, 0x00, 0xbd, 0xde, 0x1c, 0x87, 0x5f, 0xfa, 0xf2,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xb1, 0x1a, 0x2c, 0xed, 0x03, 0x00, 0x00,
}
