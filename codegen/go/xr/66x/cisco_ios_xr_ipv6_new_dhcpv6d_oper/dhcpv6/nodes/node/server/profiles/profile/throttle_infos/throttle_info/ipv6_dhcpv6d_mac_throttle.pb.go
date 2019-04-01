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
// source: ipv6_dhcpv6d_mac_throttle.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_profiles_profile_throttle_infos_throttle_info

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

type Ipv6Dhcpv6DMacThrottle_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProfileName          string   `protobuf:"bytes,2,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	MacAddress           string   `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DMacThrottle_KEYS) Reset()         { *m = Ipv6Dhcpv6DMacThrottle_KEYS{} }
func (m *Ipv6Dhcpv6DMacThrottle_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DMacThrottle_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DMacThrottle_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9c12f637bc36f3, []int{0}
}

func (m *Ipv6Dhcpv6DMacThrottle_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DMacThrottle_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DMacThrottle_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DMacThrottle_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DMacThrottle_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DMacThrottle_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DMacThrottle_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DMacThrottle_KEYS) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DMacThrottle_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

type Ipv6Dhcpv6DMacThrottle struct {
	BindingChaddr        []uint32 `protobuf:"varint,50,rep,packed,name=binding_chaddr,json=bindingChaddr,proto3" json:"binding_chaddr,omitempty"`
	Ifname               []string `protobuf:"bytes,51,rep,name=ifname,proto3" json:"ifname,omitempty"`
	State                uint32   `protobuf:"varint,52,opt,name=state,proto3" json:"state,omitempty"`
	TimeLeft             uint32   `protobuf:"varint,53,opt,name=time_left,json=timeLeft,proto3" json:"time_left,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DMacThrottle) Reset()         { *m = Ipv6Dhcpv6DMacThrottle{} }
func (m *Ipv6Dhcpv6DMacThrottle) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DMacThrottle) ProtoMessage()    {}
func (*Ipv6Dhcpv6DMacThrottle) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9c12f637bc36f3, []int{1}
}

func (m *Ipv6Dhcpv6DMacThrottle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DMacThrottle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DMacThrottle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DMacThrottle.Merge(m, src)
}
func (m *Ipv6Dhcpv6DMacThrottle) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DMacThrottle.Size(m)
}
func (m *Ipv6Dhcpv6DMacThrottle) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DMacThrottle.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DMacThrottle proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DMacThrottle) GetBindingChaddr() []uint32 {
	if m != nil {
		return m.BindingChaddr
	}
	return nil
}

func (m *Ipv6Dhcpv6DMacThrottle) GetIfname() []string {
	if m != nil {
		return m.Ifname
	}
	return nil
}

func (m *Ipv6Dhcpv6DMacThrottle) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Ipv6Dhcpv6DMacThrottle) GetTimeLeft() uint32 {
	if m != nil {
		return m.TimeLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DMacThrottle_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.throttle_infos.throttle_info.ipv6_dhcpv6d_mac_throttle_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DMacThrottle)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.throttle_infos.throttle_info.ipv6_dhcpv6d_mac_throttle")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_mac_throttle.proto", fileDescriptor_9b9c12f637bc36f3) }

var fileDescriptor_9b9c12f637bc36f3 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4b, 0xc4, 0x30,
	0x18, 0xc5, 0xa9, 0x87, 0xc7, 0x35, 0x67, 0x1d, 0x82, 0x48, 0x45, 0xf0, 0xea, 0x81, 0xd0, 0xa9,
	0x83, 0xa7, 0xb7, 0x8b, 0x38, 0x29, 0x0e, 0x75, 0x72, 0xfa, 0xc8, 0x25, 0x5f, 0x6d, 0xa0, 0x49,
	0x4a, 0x12, 0xaa, 0xab, 0x7f, 0x80, 0xff, 0xb3, 0x24, 0xed, 0x09, 0x0e, 0xb7, 0x84, 0xbc, 0xdf,
	0xf7, 0xf2, 0xf2, 0x42, 0xc8, 0x4a, 0xf6, 0xc3, 0x16, 0x44, 0xcb, 0xfb, 0x61, 0x2b, 0x40, 0x31,
	0x0e, 0xbe, 0xb5, 0xc6, 0xfb, 0x0e, 0xab, 0xde, 0x1a, 0x6f, 0xa8, 0xe4, 0xd2, 0x71, 0x03, 0xd2,
	0x38, 0xf8, 0xb2, 0x10, 0xdd, 0x1a, 0x3f, 0xff, 0x4e, 0x98, 0x1e, 0x6d, 0x35, 0x8a, 0x4a, 0x1b,
	0x81, 0x2e, 0xae, 0x95, 0x43, 0x3b, 0xa0, 0x0d, 0x11, 0x8d, 0xec, 0xd0, 0xed, 0x37, 0xd5, 0x3e,
	0x1c, 0xa4, 0x6e, 0x8c, 0xfb, 0x2f, 0xd7, 0xdf, 0x09, 0xb9, 0x3a, 0x58, 0x07, 0x9e, 0x9f, 0xde,
	0xdf, 0xe8, 0x25, 0x49, 0x43, 0x3e, 0x68, 0xa6, 0x30, 0x4f, 0x8a, 0xa4, 0x4c, 0xeb, 0x45, 0x00,
	0xaf, 0x4c, 0x21, 0xbd, 0x26, 0x27, 0xd3, 0x3d, 0xe3, 0xfc, 0x28, 0xce, 0x97, 0x13, 0x8b, 0x96,
	0x15, 0x59, 0x86, 0x50, 0x26, 0x84, 0x45, 0xe7, 0xf2, 0x59, 0x74, 0x10, 0xc5, 0xf8, 0xc3, 0x48,
	0xd6, 0x3f, 0x09, 0xb9, 0x38, 0xd8, 0x81, 0xde, 0x90, 0xd3, 0x9d, 0xd4, 0x42, 0xea, 0x0f, 0xe0,
	0x6d, 0x08, 0xc9, 0x6f, 0x8b, 0x59, 0x99, 0xd5, 0xd9, 0x44, 0x1f, 0x23, 0xa4, 0xe7, 0x64, 0x2e,
	0x9b, 0x58, 0x61, 0x53, 0xcc, 0xca, 0xb4, 0x9e, 0x14, 0x3d, 0x23, 0xc7, 0xce, 0x33, 0x8f, 0xf9,
	0x5d, 0x91, 0x94, 0x59, 0x3d, 0x8a, 0xf0, 0x26, 0x2f, 0x15, 0x42, 0x87, 0x8d, 0xcf, 0xef, 0xe3,
	0x64, 0x11, 0xc0, 0x0b, 0x36, 0x7e, 0x37, 0x8f, 0xbf, 0xb0, 0xf9, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x31, 0x7f, 0xb9, 0xa8, 0x01, 0x00, 0x00,
}
