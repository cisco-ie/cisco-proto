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
// source: ipv6_dhcpv6d_server_interface.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_interfaces_interface

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

type Ipv6Dhcpv6DServerInterface_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerInterface_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerInterface_KEYS{} }
func (m *Ipv6Dhcpv6DServerInterface_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerInterface_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerInterface_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6126e5918d440c02, []int{0}
}

func (m *Ipv6Dhcpv6DServerInterface_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerInterface_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerInterface_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerInterface_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerInterface_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerInterface_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerInterface_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6Dhcpv6DServerInterface struct {
	ServerVrfName                 string   `protobuf:"bytes,50,opt,name=server_vrf_name,json=serverVrfName,proto3" json:"server_vrf_name,omitempty"`
	ServerInterfaceMode           string   `protobuf:"bytes,51,opt,name=server_interface_mode,json=serverInterfaceMode,proto3" json:"server_interface_mode,omitempty"`
	IsServerInterfaceAmbiguous    uint32   `protobuf:"varint,52,opt,name=is_server_interface_ambiguous,json=isServerInterfaceAmbiguous,proto3" json:"is_server_interface_ambiguous,omitempty"`
	ServerInterfaceProfileName    string   `protobuf:"bytes,53,opt,name=server_interface_profile_name,json=serverInterfaceProfileName,proto3" json:"server_interface_profile_name,omitempty"`
	ServerInterfaceLeaseLimitType string   `protobuf:"bytes,54,opt,name=server_interface_lease_limit_type,json=serverInterfaceLeaseLimitType,proto3" json:"server_interface_lease_limit_type,omitempty"`
	ServerInterfaceLeaseLimits    uint32   `protobuf:"varint,55,opt,name=server_interface_lease_limits,json=serverInterfaceLeaseLimits,proto3" json:"server_interface_lease_limits,omitempty"`
	SrgRole                       string   `protobuf:"bytes,56,opt,name=srg_role,json=srgRole,proto3" json:"srg_role,omitempty"`
	SergRole                      string   `protobuf:"bytes,57,opt,name=serg_role,json=sergRole,proto3" json:"serg_role,omitempty"`
	MacThrottle                   bool     `protobuf:"varint,58,opt,name=mac_throttle,json=macThrottle,proto3" json:"mac_throttle,omitempty"`
	SrgVrfName                    string   `protobuf:"bytes,59,opt,name=srg_vrf_name,json=srgVrfName,proto3" json:"srg_vrf_name,omitempty"`
	Srgp2P                        bool     `protobuf:"varint,60,opt,name=srgp2p,proto3" json:"srgp2p,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerInterface) Reset()         { *m = Ipv6Dhcpv6DServerInterface{} }
func (m *Ipv6Dhcpv6DServerInterface) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerInterface) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_6126e5918d440c02, []int{1}
}

func (m *Ipv6Dhcpv6DServerInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterface.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerInterface) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterface.Size(m)
}
func (m *Ipv6Dhcpv6DServerInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterface.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerInterface proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerInterface) GetServerVrfName() string {
	if m != nil {
		return m.ServerVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetServerInterfaceMode() string {
	if m != nil {
		return m.ServerInterfaceMode
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetIsServerInterfaceAmbiguous() uint32 {
	if m != nil {
		return m.IsServerInterfaceAmbiguous
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerInterface) GetServerInterfaceProfileName() string {
	if m != nil {
		return m.ServerInterfaceProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetServerInterfaceLeaseLimitType() string {
	if m != nil {
		return m.ServerInterfaceLeaseLimitType
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetServerInterfaceLeaseLimits() uint32 {
	if m != nil {
		return m.ServerInterfaceLeaseLimits
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerInterface) GetSrgRole() string {
	if m != nil {
		return m.SrgRole
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetSergRole() string {
	if m != nil {
		return m.SergRole
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetMacThrottle() bool {
	if m != nil {
		return m.MacThrottle
	}
	return false
}

func (m *Ipv6Dhcpv6DServerInterface) GetSrgVrfName() string {
	if m != nil {
		return m.SrgVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerInterface) GetSrgp2P() bool {
	if m != nil {
		return m.Srgp2P
	}
	return false
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerInterface_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.interfaces.interface.ipv6_dhcpv6d_server_interface_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DServerInterface)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.interfaces.interface.ipv6_dhcpv6d_server_interface")
}

func init() {
	proto.RegisterFile("ipv6_dhcpv6d_server_interface.proto", fileDescriptor_6126e5918d440c02)
}

var fileDescriptor_6126e5918d440c02 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xeb, 0xd3, 0x40,
	0x10, 0xc5, 0x89, 0xc8, 0xd7, 0x76, 0x6d, 0x15, 0x22, 0xca, 0x5a, 0x29, 0xa4, 0x15, 0xa5, 0xa7,
	0x1c, 0x5a, 0xad, 0x3f, 0x2f, 0x3d, 0x08, 0x8a, 0x55, 0x4a, 0x5a, 0x04, 0x4f, 0x4b, 0x9a, 0x4c,
	0xd2, 0x85, 0x24, 0xbb, 0xec, 0xa4, 0xd1, 0xfe, 0x3d, 0xfe, 0xa3, 0xb2, 0xbb, 0x49, 0xaa, 0x8d,
	0xf6, 0x52, 0xba, 0x6f, 0xde, 0x7c, 0xde, 0x0c, 0x13, 0xf2, 0x94, 0xcb, 0x6a, 0xc9, 0xe2, 0x43,
	0x24, 0xab, 0x65, 0xcc, 0x10, 0x54, 0x05, 0x8a, 0xf1, 0xa2, 0x04, 0x95, 0x84, 0x11, 0xf8, 0x52,
	0x89, 0x52, 0xb8, 0x9b, 0x88, 0x63, 0x24, 0x18, 0x17, 0xc8, 0x7e, 0x2a, 0x66, 0x3a, 0x0a, 0xf8,
	0xd1, 0x76, 0x09, 0x09, 0xca, 0xb7, 0x0f, 0xbf, 0x10, 0x31, 0xa0, 0xf9, 0xf5, 0x2d, 0xcc, 0x6f,
	0x61, 0x78, 0xfe, 0x3b, 0x3d, 0x90, 0xe9, 0xd5, 0x60, 0xf6, 0xf9, 0xc3, 0xf7, 0xad, 0xfb, 0x84,
	0xf4, 0x35, 0x89, 0x15, 0x61, 0x0e, 0xd4, 0xf1, 0x9c, 0x59, 0x3f, 0xe8, 0x69, 0xe1, 0x6b, 0x98,
	0x83, 0xfb, 0x8c, 0xdc, 0x3b, 0xdb, 0x8d, 0xe3, 0x96, 0x71, 0x0c, 0x5b, 0x55, 0xdb, 0xa6, 0xbf,
	0x6e, 0x93, 0xf1, 0xd5, 0x28, 0xf7, 0x39, 0xb9, 0x5f, 0x6b, 0x95, 0x4a, 0x2c, 0x69, 0x6e, 0x49,
	0x56, 0xfe, 0xa6, 0x12, 0x13, 0x38, 0x27, 0x0f, 0x3b, 0x63, 0xe6, 0x22, 0x06, 0xba, 0x30, 0xee,
	0x07, 0xb6, 0xf8, 0xa9, 0xa9, 0x7d, 0x11, 0x31, 0xb8, 0x2b, 0x32, 0xe6, 0xd8, 0xdd, 0x2e, 0xcc,
	0xf7, 0x3c, 0x3d, 0x8a, 0x23, 0xd2, 0x17, 0x9e, 0x33, 0x1b, 0x06, 0x23, 0x8e, 0xdb, 0xbf, 0xbb,
	0x57, 0x8d, 0x43, 0x23, 0x3a, 0xfd, 0x52, 0x89, 0x84, 0x67, 0xf5, 0xda, 0x2f, 0x4d, 0xfc, 0xe8,
	0x22, 0x7e, 0x63, 0x2d, 0x66, 0xf2, 0x8f, 0x64, 0xd2, 0x41, 0x64, 0x10, 0x22, 0xb0, 0x8c, 0xe7,
	0xbc, 0x64, 0xe5, 0x49, 0x02, 0x5d, 0x1a, 0xcc, 0xf8, 0x02, 0xb3, 0xd6, 0xb6, 0xb5, 0x76, 0xed,
	0x4e, 0x12, 0xfe, 0x39, 0xcc, 0x1f, 0x24, 0xa4, 0xaf, 0xec, 0x3e, 0xff, 0xa5, 0xa0, 0xfb, 0x98,
	0xf4, 0x50, 0xa5, 0x4c, 0x89, 0x0c, 0xe8, 0x6b, 0x93, 0x79, 0x07, 0x55, 0x1a, 0x88, 0x0c, 0xf4,
	0xbd, 0x11, 0x9a, 0xda, 0x1b, 0x7b, 0x6f, 0x2d, 0x98, 0xe2, 0x84, 0x0c, 0xf2, 0x30, 0x62, 0xe5,
	0x41, 0x89, 0xb2, 0xcc, 0x80, 0xbe, 0xf5, 0x9c, 0x59, 0x2f, 0xb8, 0x9b, 0x87, 0xd1, 0xae, 0x96,
	0x5c, 0x8f, 0x0c, 0x34, 0xba, 0x3d, 0xe3, 0x3b, 0x83, 0x20, 0xa8, 0xd2, 0xe6, 0x86, 0x8f, 0xc8,
	0x0d, 0xaa, 0x54, 0xce, 0x25, 0x7d, 0x6f, 0xda, 0xeb, 0xd7, 0xfe, 0xc6, 0x7c, 0xe8, 0x8b, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0x02, 0xb2, 0x18, 0x0f, 0x03, 0x00, 0x00,
}
