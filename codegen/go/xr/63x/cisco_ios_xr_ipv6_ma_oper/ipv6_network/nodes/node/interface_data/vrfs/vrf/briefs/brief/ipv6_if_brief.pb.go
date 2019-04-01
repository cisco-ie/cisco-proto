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
// source: ipv6_if_brief.proto

package cisco_ios_xr_ipv6_ma_oper_ipv6_network_nodes_node_interface_data_vrfs_vrf_briefs_brief

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

type Ipv6IfBrief_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6IfBrief_KEYS) Reset()         { *m = Ipv6IfBrief_KEYS{} }
func (m *Ipv6IfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfBrief_KEYS) ProtoMessage()    {}
func (*Ipv6IfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_12de899f81f054b5, []int{0}
}

func (m *Ipv6IfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ipv6IfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6IfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfBrief_KEYS.Merge(m, src)
}
func (m *Ipv6IfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfBrief_KEYS.Size(m)
}
func (m *Ipv6IfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfBrief_KEYS proto.InternalMessageInfo

func (m *Ipv6IfBrief_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6AddrNode struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	AddressState         string   `protobuf:"bytes,3,opt,name=address_state,json=addressState,proto3" json:"address_state,omitempty"`
	IsAnycast            bool     `protobuf:"varint,4,opt,name=is_anycast,json=isAnycast,proto3" json:"is_anycast,omitempty"`
	RouteTag             uint32   `protobuf:"varint,5,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6AddrNode) Reset()         { *m = Ipv6AddrNode{} }
func (m *Ipv6AddrNode) String() string { return proto.CompactTextString(m) }
func (*Ipv6AddrNode) ProtoMessage()    {}
func (*Ipv6AddrNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_12de899f81f054b5, []int{1}
}

func (m *Ipv6AddrNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6AddrNode.Unmarshal(m, b)
}
func (m *Ipv6AddrNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6AddrNode.Marshal(b, m, deterministic)
}
func (m *Ipv6AddrNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6AddrNode.Merge(m, src)
}
func (m *Ipv6AddrNode) XXX_Size() int {
	return xxx_messageInfo_Ipv6AddrNode.Size(m)
}
func (m *Ipv6AddrNode) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6AddrNode.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6AddrNode proto.InternalMessageInfo

func (m *Ipv6AddrNode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Ipv6AddrNode) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6AddrNode) GetAddressState() string {
	if m != nil {
		return m.AddressState
	}
	return ""
}

func (m *Ipv6AddrNode) GetIsAnycast() bool {
	if m != nil {
		return m.IsAnycast
	}
	return false
}

func (m *Ipv6AddrNode) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

type Ipv6IfBrief struct {
	LineState            string          `protobuf:"bytes,50,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	VrfName              string          `protobuf:"bytes,51,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Address              []*Ipv6AddrNode `protobuf:"bytes,52,rep,name=address,proto3" json:"address,omitempty"`
	LinkLocalAddress     *Ipv6AddrNode   `protobuf:"bytes,53,opt,name=link_local_address,json=linkLocalAddress,proto3" json:"link_local_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Ipv6IfBrief) Reset()         { *m = Ipv6IfBrief{} }
func (m *Ipv6IfBrief) String() string { return proto.CompactTextString(m) }
func (*Ipv6IfBrief) ProtoMessage()    {}
func (*Ipv6IfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_12de899f81f054b5, []int{2}
}

func (m *Ipv6IfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6IfBrief.Unmarshal(m, b)
}
func (m *Ipv6IfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6IfBrief.Marshal(b, m, deterministic)
}
func (m *Ipv6IfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6IfBrief.Merge(m, src)
}
func (m *Ipv6IfBrief) XXX_Size() int {
	return xxx_messageInfo_Ipv6IfBrief.Size(m)
}
func (m *Ipv6IfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6IfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6IfBrief proto.InternalMessageInfo

func (m *Ipv6IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *Ipv6IfBrief) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6IfBrief) GetAddress() []*Ipv6AddrNode {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Ipv6IfBrief) GetLinkLocalAddress() *Ipv6AddrNode {
	if m != nil {
		return m.LinkLocalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6IfBrief_KEYS)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_if_brief_KEYS")
	proto.RegisterType((*Ipv6AddrNode)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_addr_node")
	proto.RegisterType((*Ipv6IfBrief)(nil), "cisco_ios_xr_ipv6_ma_oper.ipv6_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv6_if_brief")
}

func init() { proto.RegisterFile("ipv6_if_brief.proto", fileDescriptor_12de899f81f054b5) }

var fileDescriptor_12de899f81f054b5 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0x95, 0xf6, 0xff, 0x69, 0x62, 0x9a, 0x0a, 0x99, 0x25, 0x08, 0x21, 0x45, 0x45, 0x48,
	0x9d, 0x3c, 0xb4, 0xc0, 0xde, 0x81, 0x89, 0x8a, 0xa1, 0x45, 0x48, 0x4c, 0x57, 0x6e, 0x62, 0x17,
	0xab, 0xa9, 0x1d, 0xd9, 0x26, 0x94, 0x8d, 0x87, 0xe0, 0x41, 0x78, 0x27, 0x5e, 0x04, 0xd9, 0x49,
	0x29, 0x79, 0x00, 0x58, 0x1c, 0xdd, 0x73, 0x8f, 0xee, 0x77, 0x7d, 0x62, 0x74, 0x2c, 0xca, 0xea,
	0x1a, 0x04, 0x87, 0xa5, 0x16, 0x8c, 0x93, 0x52, 0x2b, 0xab, 0xf0, 0x43, 0x26, 0x4c, 0xa6, 0x40,
	0x28, 0x03, 0x5b, 0x0d, 0xde, 0xb1, 0xa1, 0xa0, 0x4a, 0xa6, 0x89, 0x2f, 0x24, 0xb3, 0x2f, 0x4a,
	0xaf, 0x89, 0x54, 0x39, 0x33, 0xfe, 0x24, 0x42, 0x5a, 0xa6, 0x39, 0xcd, 0x18, 0xe4, 0xd4, 0x52,
	0x52, 0x69, 0x6e, 0xdc, 0x41, 0xfc, 0x58, 0x53, 0x7f, 0x86, 0x06, 0xe1, 0x16, 0x0e, 0x6e, 0x6f,
	0x1e, 0x17, 0xf8, 0x14, 0x45, 0x6e, 0x04, 0x48, 0xba, 0x61, 0x49, 0x90, 0x06, 0xa3, 0x68, 0x1e,
	0x3a, 0xe1, 0x8e, 0x6e, 0x18, 0x3e, 0x41, 0x61, 0xa5, 0x79, 0xdd, 0xeb, 0xf8, 0x5e, 0xaf, 0xd2,
	0xdc, 0xb7, 0x2e, 0xd0, 0x60, 0x4f, 0xf5, 0x86, 0xae, 0x37, 0xc4, 0xdf, 0xaa, 0xb3, 0x0d, 0x3f,
	0x02, 0x34, 0xf0, 0x54, 0x9a, 0xe7, 0x1a, 0xdc, 0x60, 0x9c, 0xa0, 0x9e, 0x2b, 0x98, 0x31, 0x0d,
	0x6f, 0x57, 0xe2, 0x73, 0x14, 0x97, 0x9a, 0x71, 0xb1, 0x85, 0x82, 0xc9, 0x95, 0x7d, 0xf2, 0xcc,
	0x78, 0xde, 0xaf, 0xc5, 0x99, 0xd7, 0x9c, 0xa9, 0xf1, 0x83, 0xb1, 0xd4, 0xee, 0xb8, 0xfd, 0x46,
	0x5c, 0x38, 0x0d, 0x9f, 0x21, 0x24, 0x0c, 0x50, 0xf9, 0x9a, 0x51, 0x63, 0x93, 0x7f, 0x69, 0x30,
	0x0a, 0xe7, 0x91, 0x30, 0xd3, 0x5a, 0x70, 0x97, 0xd6, 0xea, 0xd9, 0x32, 0xb0, 0x74, 0x95, 0xfc,
	0xf7, 0x90, 0xd0, 0x0b, 0xf7, 0x74, 0x35, 0xfc, 0xec, 0xa0, 0xb8, 0x15, 0x94, 0x9b, 0x56, 0x08,
	0xc9, 0x1a, 0xde, 0xd8, 0xf3, 0x22, 0xa7, 0xd4, 0xb0, 0x9f, 0x29, 0x4d, 0xda, 0x29, 0xbd, 0x05,
	0xfb, 0xcb, 0x5e, 0xa6, 0xdd, 0xd1, 0xe1, 0x98, 0x93, 0xdf, 0xf9, 0xbd, 0xa4, 0x9d, 0xf2, 0x3e,
	0xd4, 0xf7, 0x00, 0xe1, 0x42, 0xc8, 0x35, 0x14, 0x2a, 0xa3, 0x05, 0xec, 0xb6, 0xb9, 0x4a, 0x83,
	0x3f, 0xdc, 0xe6, 0xc8, 0x6d, 0x30, 0x73, 0x0b, 0x4c, 0x6b, 0xfe, 0xf2, 0xc0, 0x3f, 0xf6, 0xc9,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0x4a, 0x1d, 0xa1, 0x03, 0x03, 0x00, 0x00,
}