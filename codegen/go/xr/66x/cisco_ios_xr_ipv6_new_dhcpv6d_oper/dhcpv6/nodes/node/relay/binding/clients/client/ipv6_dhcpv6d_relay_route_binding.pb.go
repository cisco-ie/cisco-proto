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
// source: ipv6_dhcpv6d_relay_route_binding.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_relay_binding_clients_client

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

type Ipv6Dhcpv6DRelayRouteBinding_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) Reset()         { *m = Ipv6Dhcpv6DRelayRouteBinding_KEYS{} }
func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayRouteBinding_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DRelayRouteBinding_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fbcd329c99578d, []int{0}
}

func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding_KEYS) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type Ipv6Dhcpv6DRelayRouteBinding struct {
	Duid                 string   `protobuf:"bytes,50,opt,name=duid,proto3" json:"duid,omitempty"`
	ClientIdXr           uint32   `protobuf:"varint,51,opt,name=client_id_xr,json=clientIdXr,proto3" json:"client_id_xr,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,52,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Prefix               string   `protobuf:"bytes,53,opt,name=prefix,proto3" json:"prefix,omitempty"`
	VrfName              string   `protobuf:"bytes,54,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Lifetime             uint32   `protobuf:"varint,55,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	RemLifeTime          uint32   `protobuf:"varint,56,opt,name=rem_life_time,json=remLifeTime,proto3" json:"rem_life_time,omitempty"`
	InterfaceName        string   `protobuf:"bytes,57,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NextHopAddr          string   `protobuf:"bytes,58,opt,name=next_hop_addr,json=nextHopAddr,proto3" json:"next_hop_addr,omitempty"`
	IaId                 uint32   `protobuf:"varint,59,opt,name=ia_id,json=iaId,proto3" json:"ia_id,omitempty"`
	RelayProfileName     string   `protobuf:"bytes,60,opt,name=relay_profile_name,json=relayProfileName,proto3" json:"relay_profile_name,omitempty"`
	L2IntfAcName         string   `protobuf:"bytes,61,opt,name=l2_intf_ac_name,json=l2IntfAcName,proto3" json:"l2_intf_ac_name,omitempty"`
	Ifname               []string `protobuf:"bytes,62,rep,name=ifname,proto3" json:"ifname,omitempty"`
	SergState            uint32   `protobuf:"varint,63,opt,name=serg_state,json=sergState,proto3" json:"serg_state,omitempty"`
	SergIntfRole         uint32   `protobuf:"varint,64,opt,name=serg_intf_role,json=sergIntfRole,proto3" json:"serg_intf_role,omitempty"`
	RelayBindingChaddr   []uint32 `protobuf:"varint,65,rep,packed,name=relay_binding_chaddr,json=relayBindingChaddr,proto3" json:"relay_binding_chaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) Reset()         { *m = Ipv6Dhcpv6DRelayRouteBinding{} }
func (m *Ipv6Dhcpv6DRelayRouteBinding) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DRelayRouteBinding) ProtoMessage()    {}
func (*Ipv6Dhcpv6DRelayRouteBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_d9fbcd329c99578d, []int{1}
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding.Merge(m, src)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding.Size(m)
}
func (m *Ipv6Dhcpv6DRelayRouteBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DRelayRouteBinding proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetClientIdXr() uint32 {
	if m != nil {
		return m.ClientIdXr
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetLifetime() uint32 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetRemLifeTime() uint32 {
	if m != nil {
		return m.RemLifeTime
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetNextHopAddr() string {
	if m != nil {
		return m.NextHopAddr
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetIaId() uint32 {
	if m != nil {
		return m.IaId
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetRelayProfileName() string {
	if m != nil {
		return m.RelayProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetL2IntfAcName() string {
	if m != nil {
		return m.L2IntfAcName
	}
	return ""
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetIfname() []string {
	if m != nil {
		return m.Ifname
	}
	return nil
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetSergState() uint32 {
	if m != nil {
		return m.SergState
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetSergIntfRole() uint32 {
	if m != nil {
		return m.SergIntfRole
	}
	return 0
}

func (m *Ipv6Dhcpv6DRelayRouteBinding) GetRelayBindingChaddr() []uint32 {
	if m != nil {
		return m.RelayBindingChaddr
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DRelayRouteBinding_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.clients.client.ipv6_dhcpv6d_relay_route_binding_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DRelayRouteBinding)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.relay.binding.clients.client.ipv6_dhcpv6d_relay_route_binding")
}

func init() {
	proto.RegisterFile("ipv6_dhcpv6d_relay_route_binding.proto", fileDescriptor_d9fbcd329c99578d)
}

var fileDescriptor_d9fbcd329c99578d = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x5d, 0x8b, 0x13, 0x31,
	0x14, 0xa5, 0x6e, 0xad, 0xed, 0xdd, 0xce, 0x2a, 0x51, 0x96, 0xa8, 0x08, 0x43, 0xb5, 0xd2, 0x07,
	0x29, 0xd2, 0xd5, 0xfa, 0xfd, 0x51, 0x45, 0xb0, 0xb8, 0x88, 0x76, 0x7d, 0xd0, 0xa7, 0x90, 0x9d,
	0xdc, 0xb4, 0x81, 0x99, 0x64, 0xc8, 0x64, 0x6b, 0xfd, 0x5f, 0xfe, 0xc0, 0x25, 0x37, 0x6d, 0x5f,
	0xf7, 0x65, 0x66, 0xee, 0x39, 0xe7, 0x9e, 0x93, 0x7b, 0x27, 0xf0, 0xd8, 0xd4, 0xeb, 0xa9, 0x50,
	0xab, 0xa2, 0x5e, 0x4f, 0x95, 0xf0, 0x58, 0xca, 0x7f, 0xc2, 0xbb, 0x8b, 0x80, 0xe2, 0xdc, 0x58,
	0x65, 0xec, 0x72, 0x5c, 0x7b, 0x17, 0x1c, 0xfb, 0x59, 0x98, 0xa6, 0x70, 0xc2, 0xb8, 0x46, 0x6c,
	0xbc, 0xa0, 0x26, 0x8b, 0x7f, 0xf7, 0x8d, 0xae, 0x46, 0x3f, 0x4e, 0xc5, 0xd8, 0x3a, 0x85, 0x0d,
	0x3d, 0xc7, 0xe4, 0x37, 0xde, 0x39, 0x15, 0xa5, 0x41, 0x1b, 0x9a, 0xed, 0x7b, 0x20, 0x61, 0x78,
	0x55, 0xb8, 0xf8, 0xf6, 0xe5, 0xcf, 0x19, 0xbb, 0x0f, 0xbd, 0xe8, 0x26, 0xac, 0xac, 0x90, 0xb7,
	0xf2, 0xd6, 0xa8, 0xb7, 0xe8, 0x46, 0xe0, 0xbb, 0xac, 0x30, 0x92, 0xc9, 0x4f, 0x18, 0xc5, 0xaf,
	0x25, 0x32, 0x01, 0x73, 0x35, 0xf8, 0xdf, 0x86, 0xfc, 0xaa, 0x0c, 0xc6, 0xa0, 0xad, 0x2e, 0x8c,
	0xe2, 0x13, 0x6a, 0xa6, 0x6f, 0x96, 0x43, 0x7f, 0xef, 0x2a, 0x36, 0x9e, 0x9f, 0xe4, 0xad, 0x51,
	0xb6, 0x80, 0x9d, 0xf1, 0x6f, 0xcf, 0x1e, 0x42, 0x56, 0x7b, 0xd4, 0x66, 0x23, 0x4a, 0xb4, 0xcb,
	0xb0, 0xe2, 0xcf, 0x48, 0xd2, 0x4f, 0xe0, 0x29, 0x61, 0xec, 0x18, 0x3a, 0xa9, 0xe6, 0xcf, 0xc9,
	0x7c, 0x5b, 0xb1, 0xbb, 0xd0, 0x5d, 0x7b, 0x9d, 0x06, 0x9a, 0x12, 0x73, 0x63, 0xed, 0x35, 0xcd,
	0x73, 0x0f, 0xba, 0xa5, 0xd1, 0x18, 0x4c, 0x85, 0xfc, 0x05, 0x59, 0xee, 0x6b, 0x36, 0x80, 0xcc,
	0x63, 0x25, 0x62, 0x2d, 0x48, 0xf0, 0x92, 0x04, 0x87, 0x1e, 0xab, 0x53, 0xa3, 0xf1, 0x57, 0xd4,
	0x0c, 0xe1, 0xc8, 0xd8, 0x80, 0x5e, 0xcb, 0x62, 0xbb, 0xb1, 0x57, 0x14, 0x90, 0xed, 0x51, 0x8a,
	0x19, 0x40, 0x66, 0x71, 0x13, 0xc4, 0xca, 0xd5, 0x42, 0x2a, 0xe5, 0xf9, 0x6b, 0x52, 0x1d, 0x46,
	0xf0, 0xab, 0xab, 0x67, 0x4a, 0x79, 0x76, 0x1b, 0xae, 0x1b, 0x19, 0xd7, 0xfa, 0x86, 0x62, 0xda,
	0x46, 0xce, 0x15, 0x7b, 0x02, 0x2c, 0x2d, 0xb1, 0xf6, 0x4e, 0x9b, 0x72, 0x9b, 0xf1, 0x96, 0xba,
	0x6f, 0x11, 0xf3, 0x23, 0x11, 0x14, 0x33, 0x84, 0x9b, 0xe5, 0x44, 0x18, 0x1b, 0xb4, 0x90, 0x45,
	0x92, 0xbe, 0x23, 0x69, 0xbf, 0x9c, 0xcc, 0x6d, 0xd0, 0xb3, 0x82, 0x64, 0xc7, 0xd0, 0x31, 0x9a,
	0xd8, 0xf7, 0xf9, 0x41, 0xdc, 0x53, 0xaa, 0xd8, 0x03, 0x80, 0x06, 0xfd, 0x52, 0x34, 0x41, 0x06,
	0xe4, 0x1f, 0xe8, 0x18, 0xbd, 0x88, 0x9c, 0x45, 0x80, 0x3d, 0x82, 0x23, 0xa2, 0xc9, 0xdf, 0xbb,
	0x12, 0xf9, 0xc7, 0xf4, 0x13, 0x22, 0x1a, 0xed, 0x17, 0xae, 0x44, 0xf6, 0x14, 0xee, 0xa4, 0x13,
	0xef, 0x2e, 0x55, 0xb1, 0xa2, 0x89, 0x67, 0xf9, 0xc1, 0x28, 0x5b, 0xa4, 0x69, 0x3e, 0x25, 0xea,
	0x33, 0x31, 0xe7, 0x1d, 0xba, 0xf3, 0x27, 0x97, 0x01, 0x00, 0x00, 0xff, 0xff, 0xed, 0x7c, 0x54,
	0x81, 0x1d, 0x03, 0x00, 0x00,
}
