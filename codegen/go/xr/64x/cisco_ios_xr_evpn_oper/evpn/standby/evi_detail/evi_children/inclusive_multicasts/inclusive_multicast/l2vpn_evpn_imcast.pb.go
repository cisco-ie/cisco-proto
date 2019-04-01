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
	EthernetTag          uint32   `protobuf:"varint,2,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	OriginatingIp        string   `protobuf:"bytes,3,opt,name=originating_ip,json=originatingIp,proto3" json:"originating_ip,omitempty"`
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

type L2VpnEvpnImcast struct {
	EviXr                uint32   `protobuf:"varint,50,opt,name=evi_xr,json=eviXr,proto3" json:"evi_xr,omitempty"`
	EthernetTagXr        uint32   `protobuf:"varint,51,opt,name=ethernet_tag_xr,json=ethernetTagXr,proto3" json:"ethernet_tag_xr,omitempty"`
	OriginatingIpXr      string   `protobuf:"bytes,52,opt,name=originating_ip_xr,json=originatingIpXr,proto3" json:"originating_ip_xr,omitempty"`
	NextHop              string   `protobuf:"bytes,53,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	OutputLabel          uint32   `protobuf:"varint,54,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	IsLocalEntry         bool     `protobuf:"varint,55,opt,name=is_local_entry,json=isLocalEntry,proto3" json:"is_local_entry,omitempty"`
	IsProxyEntry         bool     `protobuf:"varint,56,opt,name=is_proxy_entry,json=isProxyEntry,proto3" json:"is_proxy_entry,omitempty"`
	EncapType            uint32   `protobuf:"varint,57,opt,name=encap_type,json=encapType,proto3" json:"encap_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnImcast) Reset()         { *m = L2VpnEvpnImcast{} }
func (m *L2VpnEvpnImcast) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnImcast) ProtoMessage()    {}
func (*L2VpnEvpnImcast) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb00aa10138a6cf7, []int{1}
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

func (m *L2VpnEvpnImcast) GetEviXr() uint32 {
	if m != nil {
		return m.EviXr
	}
	return 0
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

func (m *L2VpnEvpnImcast) GetEncapType() uint32 {
	if m != nil {
		return m.EncapType
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnEvpnImcast_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast_KEYS")
	proto.RegisterType((*L2VpnEvpnImcast)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.evi_detail.evi_children.inclusive_multicasts.inclusive_multicast.l2vpn_evpn_imcast")
}

func init() { proto.RegisterFile("l2vpn_evpn_imcast.proto", fileDescriptor_eb00aa10138a6cf7) }

var fileDescriptor_eb00aa10138a6cf7 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5d, 0xcb, 0xda, 0x30,
	0x14, 0xc7, 0xa9, 0x32, 0xa7, 0x99, 0x2f, 0x33, 0xb0, 0x2d, 0xbb, 0x18, 0x38, 0xd9, 0x86, 0xec,
	0xa2, 0x17, 0xba, 0xd7, 0x0f, 0x20, 0x6c, 0xcc, 0x8b, 0xe1, 0xbc, 0x70, 0x57, 0x21, 0xb6, 0x87,
	0x7a, 0x20, 0x26, 0x21, 0x49, 0x4b, 0xfb, 0x89, 0x9e, 0xaf, 0xf9, 0x90, 0x54, 0x41, 0xf1, 0xb9,
	0x29, 0xa7, 0xbf, 0xff, 0x9f, 0xf3, 0xeb, 0x81, 0x92, 0x37, 0x72, 0x59, 0x19, 0xc5, 0x21, 0x3c,
	0xf0, 0x94, 0x09, 0xe7, 0x53, 0x63, 0xb5, 0xd7, 0x34, 0xcf, 0xd0, 0x65, 0x9a, 0xa3, 0x76, 0xbc,
	0xb6, 0x6d, 0xae, 0x0d, 0xd8, 0x34, 0x4c, 0xa9, 0xf3, 0x42, 0xe5, 0x87, 0x26, 0x85, 0x0a, 0x79,
	0x0e, 0x5e, 0xa0, 0x8c, 0x63, 0x76, 0x44, 0x99, 0x5b, 0x50, 0x29, 0xaa, 0x4c, 0x96, 0x0e, 0x2b,
	0xe0, 0xa7, 0x52, 0x7a, 0x0c, 0x6b, 0xdd, 0x53, 0x70, 0xee, 0xc9, 0xeb, 0xbb, 0x0f, 0xe0, 0x7f,
	0xd6, 0xff, 0xff, 0xd1, 0x97, 0xa4, 0x0b, 0x15, 0xb2, 0x64, 0x96, 0x2c, 0x46, 0xdb, 0x30, 0xd2,
	0xf7, 0x64, 0x08, 0xfe, 0x08, 0x56, 0x81, 0xe7, 0x5e, 0x14, 0xac, 0x13, 0xa3, 0x17, 0x17, 0xb6,
	0x13, 0x05, 0xfd, 0x48, 0xc6, 0xda, 0x62, 0x81, 0x4a, 0x78, 0x54, 0x05, 0x47, 0xc3, 0xba, 0xb3,
	0x64, 0x31, 0xd8, 0x8e, 0xae, 0xe8, 0x6f, 0x33, 0x7f, 0xe8, 0x90, 0xe9, 0x9d, 0x96, 0xbe, 0x22,
	0xbd, 0x70, 0x44, 0x6d, 0xd9, 0x32, 0x6e, 0x7e, 0x06, 0x15, 0xee, 0x2d, 0xfd, 0x44, 0x26, 0xd7,
	0xda, 0x90, 0xaf, 0x62, 0x3e, 0xba, 0x32, 0xef, 0x2d, 0xfd, 0x4c, 0xa6, 0xb7, 0xee, 0xd0, 0xfc,
	0x12, 0xf5, 0x93, 0x1b, 0xfd, 0xde, 0xd2, 0xb7, 0xa4, 0xaf, 0xa0, 0xf6, 0xfc, 0xa8, 0x0d, 0xfb,
	0x1a, 0x2b, 0xcf, 0xc3, 0xfb, 0x2f, 0x6d, 0xc2, 0x95, 0xba, 0xf4, 0xa6, 0xf4, 0x5c, 0x8a, 0x03,
	0x48, 0xf6, 0xad, 0xbd, 0xb2, 0x65, 0x9b, 0x80, 0xe8, 0x07, 0x32, 0x46, 0xc7, 0xa5, 0xce, 0x84,
	0xe4, 0xa0, 0xbc, 0x6d, 0xd8, 0xf7, 0x59, 0xb2, 0xe8, 0x6f, 0x87, 0xe8, 0x36, 0x01, 0xae, 0x03,
	0x3b, 0xb7, 0x8c, 0xd5, 0x75, 0x73, 0x6e, 0xfd, 0xb8, 0xb4, 0xfe, 0x06, 0xd8, 0xb6, 0xde, 0x11,
	0x02, 0x2a, 0x13, 0x86, 0xfb, 0xc6, 0x00, 0xfb, 0x19, 0x65, 0x83, 0x48, 0x76, 0x8d, 0x81, 0x43,
	0x2f, 0xfe, 0x0c, 0xab, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0xb5, 0xe5, 0xfc, 0x27, 0x02,
	0x00, 0x00,
}