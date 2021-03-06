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
// source: ipv6_nd_vr_entry.proto

package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_nd_virtual_routers_nd_virtual_router

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

type Ipv6NdVrEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdVrEntry_KEYS) Reset()         { *m = Ipv6NdVrEntry_KEYS{} }
func (m *Ipv6NdVrEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdVrEntry_KEYS) ProtoMessage()    {}
func (*Ipv6NdVrEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a4a3a96b19770d7, []int{0}
}

func (m *Ipv6NdVrEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdVrEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv6NdVrEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdVrEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6NdVrEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdVrEntry_KEYS.Merge(m, src)
}
func (m *Ipv6NdVrEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdVrEntry_KEYS.Size(m)
}
func (m *Ipv6NdVrEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdVrEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdVrEntry_KEYS proto.InternalMessageInfo

func (m *Ipv6NdVrEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdVrEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6NdAddr struct {
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	ValidLifetime        uint32   `protobuf:"varint,2,opt,name=valid_lifetime,json=validLifetime,proto3" json:"valid_lifetime,omitempty"`
	PrefLifetime         uint32   `protobuf:"varint,3,opt,name=pref_lifetime,json=prefLifetime,proto3" json:"pref_lifetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdAddr) Reset()         { *m = Ipv6NdAddr{} }
func (m *Ipv6NdAddr) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdAddr) ProtoMessage()    {}
func (*Ipv6NdAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a4a3a96b19770d7, []int{1}
}

func (m *Ipv6NdAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdAddr.Unmarshal(m, b)
}
func (m *Ipv6NdAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdAddr.Marshal(b, m, deterministic)
}
func (m *Ipv6NdAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdAddr.Merge(m, src)
}
func (m *Ipv6NdAddr) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdAddr.Size(m)
}
func (m *Ipv6NdAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdAddr.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdAddr proto.InternalMessageInfo

func (m *Ipv6NdAddr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *Ipv6NdAddr) GetValidLifetime() uint32 {
	if m != nil {
		return m.ValidLifetime
	}
	return 0
}

func (m *Ipv6NdAddr) GetPrefLifetime() uint32 {
	if m != nil {
		return m.PrefLifetime
	}
	return 0
}

type Ipv6NdVrEntry struct {
	LinkLayerAddress     string        `protobuf:"bytes,50,opt,name=link_layer_address,json=linkLayerAddress,proto3" json:"link_layer_address,omitempty"`
	LocalAddress         *Ipv6NdAddr   `protobuf:"bytes,51,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	Context              uint32        `protobuf:"varint,52,opt,name=context,proto3" json:"context,omitempty"`
	State                string        `protobuf:"bytes,53,opt,name=state,proto3" json:"state,omitempty"`
	Flags                string        `protobuf:"bytes,54,opt,name=flags,proto3" json:"flags,omitempty"`
	VrGlAddrCt           uint32        `protobuf:"varint,55,opt,name=vr_gl_addr_ct,json=vrGlAddrCt,proto3" json:"vr_gl_addr_ct,omitempty"`
	VrGlobalAddress      []*Ipv6NdAddr `protobuf:"bytes,56,rep,name=vr_global_address,json=vrGlobalAddress,proto3" json:"vr_global_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ipv6NdVrEntry) Reset()         { *m = Ipv6NdVrEntry{} }
func (m *Ipv6NdVrEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdVrEntry) ProtoMessage()    {}
func (*Ipv6NdVrEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a4a3a96b19770d7, []int{2}
}

func (m *Ipv6NdVrEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdVrEntry.Unmarshal(m, b)
}
func (m *Ipv6NdVrEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdVrEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6NdVrEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdVrEntry.Merge(m, src)
}
func (m *Ipv6NdVrEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdVrEntry.Size(m)
}
func (m *Ipv6NdVrEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdVrEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdVrEntry proto.InternalMessageInfo

func (m *Ipv6NdVrEntry) GetLinkLayerAddress() string {
	if m != nil {
		return m.LinkLayerAddress
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetLocalAddress() *Ipv6NdAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Ipv6NdVrEntry) GetContext() uint32 {
	if m != nil {
		return m.Context
	}
	return 0
}

func (m *Ipv6NdVrEntry) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetFlags() string {
	if m != nil {
		return m.Flags
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetVrGlAddrCt() uint32 {
	if m != nil {
		return m.VrGlAddrCt
	}
	return 0
}

func (m *Ipv6NdVrEntry) GetVrGlobalAddress() []*Ipv6NdAddr {
	if m != nil {
		return m.VrGlobalAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6NdVrEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_vr_entry_KEYS")
	proto.RegisterType((*Ipv6NdAddr)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_addr")
	proto.RegisterType((*Ipv6NdVrEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_vr_entry")
}

func init() { proto.RegisterFile("ipv6_nd_vr_entry.proto", fileDescriptor_2a4a3a96b19770d7) }

var fileDescriptor_2a4a3a96b19770d7 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x4a, 0xeb, 0x40,
	0x18, 0x85, 0xc9, 0x2d, 0xf7, 0x5e, 0x3b, 0x6d, 0xb4, 0x0e, 0x2a, 0x01, 0x37, 0x6d, 0x45, 0xe8,
	0x42, 0xb2, 0x68, 0xb5, 0xba, 0x15, 0x11, 0x17, 0x16, 0x17, 0x75, 0x25, 0x22, 0xc3, 0x34, 0x99,
	0x94, 0xa1, 0xd3, 0x4c, 0xf8, 0x67, 0x1a, 0xda, 0x85, 0xef, 0xe0, 0xce, 0x37, 0xf3, 0x79, 0x64,
	0xfe, 0x24, 0xb5, 0xd4, 0xb5, 0x6e, 0x0a, 0xe7, 0xcc, 0xc7, 0x99, 0x73, 0xa6, 0x84, 0x1c, 0xc9,
	0x2c, 0x1f, 0xb2, 0x34, 0x66, 0x39, 0x30, 0x91, 0x5a, 0x58, 0x85, 0x19, 0x68, 0xab, 0xe9, 0x4b,
	0x24, 0x4d, 0xa4, 0x99, 0xd4, 0x86, 0x2d, 0x81, 0x55, 0x90, 0xce, 0x04, 0x84, 0x85, 0xd0, 0xb1,
	0x60, 0xb1, 0x63, 0x72, 0x01, 0xab, 0xd0, 0x49, 0x83, 0xbf, 0xa1, 0xcb, 0x92, 0x60, 0x17, 0x5c,
	0x31, 0xd0, 0x0b, 0x2b, 0xc0, 0x7c, 0xb7, 0xba, 0xcf, 0xe4, 0x70, 0xfb, 0x62, 0x76, 0x7f, 0xfb,
	0xf4, 0x48, 0x8f, 0x49, 0x1d, 0xa3, 0x53, 0x3e, 0x17, 0x81, 0xd7, 0xf6, 0x7a, 0xf5, 0xf1, 0x8e,
	0x33, 0x1e, 0xf8, 0x5c, 0xd0, 0x53, 0xb2, 0x2b, 0x53, 0x2b, 0x20, 0xe1, 0x51, 0x49, 0xfc, 0x41,
	0xc2, 0x5f, 0xbb, 0x0e, 0xeb, 0xbe, 0x92, 0x66, 0x15, 0xce, 0xe3, 0x18, 0x68, 0xa7, 0xd4, 0x4e,
	0x08, 0x63, 0xca, 0xd8, 0x86, 0xf3, 0xae, 0x0b, 0xcb, 0x25, 0xe7, 0x5c, 0xc9, 0x98, 0x29, 0x99,
	0x08, 0x2b, 0xcb, 0x64, 0x7f, 0xec, 0xa3, 0x3b, 0x2a, 0x4d, 0x7a, 0x42, 0xfc, 0x0c, 0x44, 0xf2,
	0x45, 0xd5, 0x90, 0x6a, 0x3a, 0xb3, 0x82, 0xba, 0x1f, 0x35, 0xd2, 0xda, 0x1e, 0x47, 0xcf, 0x08,
	0x55, 0x32, 0x9d, 0x31, 0xc5, 0x57, 0x02, 0xd6, 0x4d, 0xfa, 0xd8, 0xa4, 0xe5, 0x4e, 0x46, 0xee,
	0xa0, 0xaa, 0xf3, 0xe6, 0x11, 0x5f, 0xe9, 0x88, 0xab, 0x35, 0x39, 0x68, 0x7b, 0xbd, 0x46, 0x7f,
	0x16, 0xfe, 0xe8, 0xdf, 0x12, 0x6e, 0x3e, 0xdb, 0xb8, 0x89, 0x0d, 0xaa, 0x4a, 0x01, 0xf9, 0x1f,
	0xe9, 0xd4, 0x8a, 0xa5, 0x0d, 0xce, 0x71, 0x74, 0x25, 0xe9, 0x01, 0xf9, 0x6b, 0x2c, 0xb7, 0x22,
	0xb8, 0xc0, 0x35, 0x85, 0x70, 0x6e, 0xa2, 0xf8, 0xd4, 0x04, 0xc3, 0xc2, 0x45, 0x41, 0x3b, 0xc4,
	0xcf, 0x81, 0x4d, 0x8b, 0x5d, 0x2c, 0xb2, 0xc1, 0x25, 0x66, 0x91, 0x1c, 0xee, 0xf0, 0xa6, 0x1b,
	0x4b, 0xdf, 0x3d, 0xb2, 0x8f, 0x8c, 0x9e, 0x6c, 0xec, 0xbf, 0x6a, 0xd7, 0x7e, 0x7b, 0xff, 0x9e,
	0x2b, 0xe5, 0x4a, 0x94, 0x4f, 0x30, 0xf9, 0x87, 0x9f, 0xc6, 0xe0, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x97, 0xdc, 0x32, 0xb7, 0x34, 0x03, 0x00, 0x00,
}
