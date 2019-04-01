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
// source: ipv6_dhcpv6d_server_binding_options.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_options_mac_bind_options_mac_bind_option

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

type Ipv6Dhcpv6DServerBindingOptions_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	MacAddress           string   `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerBindingOptions_KEYS{} }
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingOptions_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0286bae149069f3, []int{0}
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

type Ipv6Dhcpv6DServerBindingOptions struct {
	MacAddressXr         string   `protobuf:"bytes,50,opt,name=mac_address_xr,json=macAddressXr,proto3" json:"mac_address_xr,omitempty"`
	DuidXr               string   `protobuf:"bytes,51,opt,name=duid_xr,json=duidXr,proto3" json:"duid_xr,omitempty"`
	DnsCount             uint32   `protobuf:"varint,52,opt,name=dns_count,json=dnsCount,proto3" json:"dns_count,omitempty"`
	DnsAddress           []string `protobuf:"bytes,53,rep,name=dns_address,json=dnsAddress,proto3" json:"dns_address,omitempty"`
	Opt17                string   `protobuf:"bytes,54,opt,name=opt17,proto3" json:"opt17,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBindingOptions) Reset()         { *m = Ipv6Dhcpv6DServerBindingOptions{} }
func (m *Ipv6Dhcpv6DServerBindingOptions) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBindingOptions) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBindingOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0286bae149069f3, []int{1}
}

func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.Size(m)
}
func (m *Ipv6Dhcpv6DServerBindingOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBindingOptions proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBindingOptions) GetMacAddressXr() string {
	if m != nil {
		return m.MacAddressXr
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDuidXr() string {
	if m != nil {
		return m.DuidXr
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDnsCount() uint32 {
	if m != nil {
		return m.DnsCount
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetDnsAddress() []string {
	if m != nil {
		return m.DnsAddress
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerBindingOptions) GetOpt17() string {
	if m != nil {
		return m.Opt17
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingOptions_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DServerBindingOptions)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding_options.mac_bind_options.mac_bind_option.ipv6_dhcpv6d_server_binding_options")
}

func init() {
	proto.RegisterFile("ipv6_dhcpv6d_server_binding_options.proto", fileDescriptor_c0286bae149069f3)
}

var fileDescriptor_c0286bae149069f3 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0xc5, 0xda, 0xc6, 0x3f, 0x87, 0x20, 0xb8, 0xe0, 0xc1, 0xa5, 0x7a, 0x58, 0x2f,
	0x0b, 0x5a, 0x5d, 0xcf, 0x22, 0x9e, 0x04, 0x0f, 0xf5, 0x52, 0x4f, 0x43, 0x9a, 0x09, 0x36, 0xe0,
	0x26, 0x21, 0xb3, 0x5d, 0xfb, 0xd1, 0xfc, 0x78, 0x32, 0x59, 0xff, 0x21, 0x14, 0x7a, 0x09, 0xbc,
	0xf7, 0x98, 0xdf, 0x7b, 0x10, 0x71, 0x61, 0x43, 0x57, 0x03, 0x2e, 0x74, 0xe8, 0x6a, 0x04, 0x32,
	0xb1, 0x33, 0x11, 0xe6, 0xd6, 0xa1, 0x75, 0xaf, 0xe0, 0x43, 0x6b, 0xbd, 0xa3, 0x2a, 0x44, 0xdf,
	0x7a, 0xf9, 0xa6, 0x2d, 0x69, 0x0f, 0xd6, 0x13, 0xac, 0x22, 0xa4, 0x3b, 0x67, 0xde, 0x7f, 0x6e,
	0x7d, 0x30, 0xb1, 0xea, 0x45, 0xe5, 0x3c, 0x1a, 0x4a, 0x6f, 0xd5, 0x23, 0xab, 0xff, 0xc8, 0x46,
	0xe9, 0x54, 0xb3, 0xce, 0x18, 0x2f, 0x44, 0xb9, 0xc1, 0x34, 0x78, 0x7c, 0x78, 0x79, 0x96, 0x27,
	0x62, 0xc4, 0x5d, 0xe0, 0x54, 0x63, 0xf2, 0xac, 0xc8, 0xca, 0xd1, 0x74, 0xc8, 0xc6, 0x93, 0x6a,
	0x8c, 0x3c, 0x15, 0x7b, 0xcc, 0x56, 0x88, 0xd1, 0x10, 0xe5, 0x5b, 0x29, 0x16, 0x8d, 0xd2, 0x77,
	0xbd, 0x33, 0xfe, 0xc8, 0xc4, 0xd9, 0x06, 0x55, 0xf2, 0x5c, 0x1c, 0xfe, 0x01, 0xc1, 0x2a, 0xe6,
	0x57, 0x89, 0xb5, 0xff, 0xcb, 0x9a, 0x45, 0x79, 0x2c, 0x76, 0x71, 0x69, 0x91, 0xe3, 0x49, 0x8a,
	0x07, 0x2c, 0x67, 0x91, 0x47, 0xa2, 0x23, 0xd0, 0x7e, 0xe9, 0xda, 0xfc, 0xba, 0xc8, 0xca, 0x83,
	0xe9, 0x10, 0x1d, 0xdd, 0xb3, 0xe6, 0x91, 0x1c, 0x7e, 0x8f, 0xbc, 0x29, 0xb6, 0x79, 0x24, 0x3a,
	0xfa, 0x02, 0xcb, 0x23, 0xb1, 0xe3, 0x43, 0x7b, 0x79, 0x9b, 0xd7, 0x09, 0xda, 0x8b, 0xf9, 0x20,
	0xfd, 0xcc, 0xe4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x28, 0xfb, 0xaf, 0xc2, 0xc6, 0x01, 0x00, 0x00,
}
