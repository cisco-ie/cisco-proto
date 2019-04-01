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

package cisco_ios_xr_evpn_oper_evpn_standby_evi_detail_evi_children_inclusive_multicasts_inclusive_multicast

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
	proto.RegisterType((*L2VpnEvpnImcast_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnEvpnImcast)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast")
}

func init() { proto.RegisterFile("l2vpn_evpn_imcast.proto", fileDescriptor_eb00aa10138a6cf7) }

var fileDescriptor_eb00aa10138a6cf7 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x15, 0x56, 0x6d, 0xab, 0xfb, 0x77, 0x3e, 0xb0, 0x70, 0x40, 0x94, 0x6a, 0xa0, 0x82,
	0x50, 0x0e, 0x1d, 0xff, 0xe1, 0x5a, 0x89, 0x8a, 0x0a, 0x4d, 0xa5, 0x42, 0xe5, 0x64, 0xb9, 0xf1,
	0xbb, 0xd4, 0x92, 0x63, 0x5b, 0xb6, 0x53, 0x35, 0x12, 0x1f, 0x04, 0x71, 0xe3, 0x03, 0xf0, 0x1d,
	0x91, 0x9d, 0x14, 0x11, 0x6d, 0xe7, 0x5d, 0xa2, 0xd7, 0xbf, 0xf7, 0x89, 0xfd, 0xe4, 0x79, 0x1d,
	0x74, 0x2e, 0xa6, 0x3b, 0x2d, 0x09, 0xf8, 0x07, 0xcf, 0x53, 0x6a, 0x5d, 0xa2, 0x8d, 0x72, 0x0a,
	0xb3, 0x94, 0xdb, 0x54, 0x11, 0xae, 0x2c, 0xd9, 0x9b, 0xaa, 0xaf, 0x34, 0x98, 0xc4, 0x57, 0x89,
	0x75, 0x54, 0xb2, 0x4d, 0x99, 0xc0, 0x8e, 0x13, 0x06, 0x8e, 0x72, 0x11, 0xca, 0x74, 0xcb, 0x05,
	0x33, 0x20, 0x13, 0x2e, 0x53, 0x51, 0x58, 0xbe, 0x03, 0x92, 0x17, 0xc2, 0x71, 0xbf, 0xad, 0xbd,
	0x0d, 0x8e, 0x7f, 0x45, 0xe8, 0xfe, 0x0d, 0x07, 0xe4, 0xf3, 0xec, 0xfb, 0x57, 0x3c, 0x44, 0x47,
	0xb0, 0xe3, 0x71, 0x34, 0x8a, 0x26, 0xbd, 0xa5, 0x2f, 0xf1, 0x05, 0xea, 0x81, 0x4c, 0xa9, 0xb6,
	0x85, 0xa0, 0x8e, 0x2b, 0x19, 0xdf, 0x0b, 0xbd, 0x26, 0xc4, 0x8f, 0x51, 0x17, 0xdc, 0x16, 0x8c,
	0x04, 0x47, 0x1c, 0xcd, 0xe2, 0xa3, 0x20, 0xea, 0x1c, 0xd8, 0x8a, 0x66, 0xf8, 0x09, 0xea, 0x2b,
	0xc3, 0x33, 0x2e, 0xa9, 0xe3, 0x32, 0x23, 0x5c, 0xc7, 0xad, 0x51, 0x34, 0x69, 0x2f, 0x7b, 0xff,
	0xd1, 0xb9, 0x1e, 0xff, 0x6c, 0x9a, 0xf3, 0x5f, 0x68, 0x8b, 0x3c, 0xa7, 0xa6, 0xc4, 0x4f, 0xd1,
	0xe0, 0xdf, 0x21, 0xc1, 0x38, 0xab, 0x8d, 0xf6, 0x0e, 0xf8, 0x9b, 0x96, 0x73, 0x86, 0x9f, 0xa1,
	0x61, 0xc3, 0x1d, 0xd9, 0x9b, 0xda, 0xf5, 0xa0, 0xc1, 0xd7, 0x06, 0x9f, 0xa3, 0x93, 0x0d, 0x23,
	0x92, 0xe6, 0x10, 0x2c, 0xb7, 0x97, 0xc7, 0x1b, 0xf6, 0x85, 0xe6, 0x80, 0x31, 0x6a, 0xb9, 0x52,
	0x43, 0xed, 0x31, 0xd4, 0xe3, 0x3f, 0x2d, 0x74, 0x76, 0x23, 0x37, 0xfc, 0x3b, 0x42, 0xbd, 0x6a,
	0x2d, 0xfd, 0x84, 0x52, 0x88, 0xa7, 0xa3, 0x68, 0xd2, 0x99, 0xfe, 0x48, 0xee, 0x62, 0x98, 0xc9,
	0xed, 0x59, 0x2d, 0xbb, 0x9e, 0xcc, 0x6b, 0x47, 0x8d, 0xe4, 0x1c, 0xcd, 0x7c, 0x20, 0x97, 0xcd,
	0xe4, 0x56, 0x34, 0x5b, 0x1b, 0xfc, 0x1c, 0x9d, 0x35, 0x67, 0xe4, 0x95, 0x2f, 0x43, 0x04, 0x83,
	0xc6, 0x98, 0xd6, 0x06, 0xbf, 0x40, 0xd8, 0x15, 0x52, 0x82, 0x20, 0x20, 0x99, 0x56, 0x5c, 0x3a,
	0x3f, 0x90, 0x57, 0x61, 0xdb, 0x61, 0xd5, 0x99, 0xd5, 0x8d, 0x39, 0xc3, 0x13, 0x34, 0xd4, 0xb9,
	0xe5, 0xa4, 0x7e, 0x25, 0x64, 0xfb, 0x3a, 0x68, 0xfb, 0x9e, 0xaf, 0x02, 0x5e, 0x95, 0x1a, 0xf0,
	0x03, 0x74, 0x2a, 0x61, 0xef, 0xc8, 0x56, 0xe9, 0xf8, 0x4d, 0x38, 0xfa, 0xc4, 0xaf, 0x3f, 0x29,
	0xed, 0x6f, 0x99, 0x2a, 0x9c, 0x2e, 0x1c, 0x11, 0x74, 0x03, 0x22, 0x7e, 0x5b, 0xdd, 0xb2, 0x8a,
	0x2d, 0x3c, 0xc2, 0x17, 0xa8, 0xcf, 0x2d, 0x11, 0x2a, 0xa5, 0xde, 0x97, 0x33, 0x65, 0xfc, 0x6e,
	0x14, 0x4d, 0x4e, 0x97, 0x5d, 0x6e, 0x17, 0x1e, 0xce, 0x3c, 0xab, 0x55, 0xda, 0xa8, 0x7d, 0x59,
	0xab, 0xde, 0x1f, 0x54, 0x57, 0x1e, 0x56, 0xaa, 0x47, 0xa8, 0x63, 0x8d, 0x03, 0xa2, 0x95, 0xe0,
	0x69, 0x19, 0x7f, 0x08, 0x66, 0x90, 0x47, 0x57, 0x81, 0xe0, 0x87, 0x08, 0x81, 0x33, 0x00, 0x44,
	0x00, 0xbd, 0x8e, 0x3f, 0x86, 0x2d, 0xda, 0x81, 0x2c, 0x80, 0x5e, 0x6f, 0x8e, 0xc3, 0x4f, 0x7d,
	0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x27, 0xb7, 0x2e, 0xef, 0x03, 0x00, 0x00,
}