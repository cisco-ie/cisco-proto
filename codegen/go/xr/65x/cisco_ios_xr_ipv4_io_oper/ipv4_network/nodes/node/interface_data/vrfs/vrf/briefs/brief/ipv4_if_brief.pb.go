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
// source: ipv4_if_brief.proto

package cisco_ios_xr_ipv4_io_oper_ipv4_network_nodes_node_interface_data_vrfs_vrf_briefs_brief

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

type Ipv4IfBrief_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief_KEYS) Reset()         { *m = Ipv4IfBrief_KEYS{} }
func (m *Ipv4IfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief_KEYS) ProtoMessage()    {}
func (*Ipv4IfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_306830dfda42ed3e, []int{0}
}

func (m *Ipv4IfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ipv4IfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4IfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief_KEYS.Merge(m, src)
}
func (m *Ipv4IfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Size(m)
}
func (m *Ipv4IfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief_KEYS proto.InternalMessageInfo

func (m *Ipv4IfBrief_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv4IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv4IfBrief struct {
	PrimaryAddress       string   `protobuf:"bytes,50,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	VrfId                uint32   `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	VrfName              string   `protobuf:"bytes,52,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LineState            string   `protobuf:"bytes,53,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief) Reset()         { *m = Ipv4IfBrief{} }
func (m *Ipv4IfBrief) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief) ProtoMessage()    {}
func (*Ipv4IfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_306830dfda42ed3e, []int{1}
}

func (m *Ipv4IfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief.Unmarshal(m, b)
}
func (m *Ipv4IfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief.Marshal(b, m, deterministic)
}
func (m *Ipv4IfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief.Merge(m, src)
}
func (m *Ipv4IfBrief) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief.Size(m)
}
func (m *Ipv4IfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief proto.InternalMessageInfo

func (m *Ipv4IfBrief) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *Ipv4IfBrief) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Ipv4IfBrief) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv4IfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv4_if_brief_KEYS")
	proto.RegisterType((*Ipv4IfBrief)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.interface_data.vrfs.vrf.briefs.brief.ipv4_if_brief")
}

func init() { proto.RegisterFile("ipv4_if_brief.proto", fileDescriptor_306830dfda42ed3e) }

var fileDescriptor_306830dfda42ed3e = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0xff, 0x1f, 0x6b, 0x77, 0x60, 0x2b, 0x44, 0x84, 0x15, 0x11, 0x4a, 0x41, 0xec,
	0x29, 0x07, 0x5b, 0x3f, 0x80, 0x07, 0x0f, 0x22, 0x78, 0x68, 0x41, 0xf0, 0x34, 0xa4, 0x9b, 0x09,
	0x04, 0xdd, 0xcd, 0x32, 0x09, 0xab, 0x7e, 0x04, 0xbf, 0xb5, 0x64, 0x56, 0x94, 0xbd, 0x4c, 0xc8,
	0xef, 0xbd, 0xcc, 0x7b, 0x04, 0x4e, 0x7d, 0x3f, 0x6c, 0xd1, 0x3b, 0x3c, 0xb0, 0x27, 0xa7, 0x7b,
	0x0e, 0x29, 0xa8, 0xe7, 0xc6, 0xc7, 0x26, 0xa0, 0x0f, 0x11, 0x3f, 0x18, 0x47, 0x47, 0xc0, 0xd0,
	0x13, 0x6b, 0xb9, 0x74, 0x94, 0xde, 0x03, 0xbf, 0xea, 0x2e, 0x58, 0x8a, 0x32, 0xb5, 0xef, 0x12,
	0xb1, 0x33, 0x0d, 0xa1, 0x35, 0xc9, 0xe8, 0x81, 0x5d, 0xcc, 0x43, 0xcb, 0xda, 0x38, 0x1e, 0xab,
	0x08, 0x6a, 0x12, 0x87, 0x8f, 0xf7, 0x2f, 0x7b, 0x75, 0x01, 0x65, 0x5e, 0x81, 0x9d, 0x69, 0xa9,
	0x2e, 0x96, 0xc5, 0xba, 0xdc, 0xcd, 0x33, 0x78, 0x32, 0x2d, 0xa9, 0x73, 0x98, 0x0f, 0xec, 0x46,
	0xed, 0x9f, 0x68, 0xc7, 0x03, 0x3b, 0x91, 0xae, 0x60, 0xf1, 0x97, 0x2a, 0x86, 0xff, 0x62, 0xa8,
	0x7e, 0x69, 0xb6, 0xad, 0xbe, 0x0a, 0xa8, 0x26, 0xa9, 0xea, 0x1a, 0x4e, 0x7a, 0xf6, 0xad, 0xe1,
	0x4f, 0x34, 0xd6, 0x32, 0xc5, 0x58, 0xdf, 0xc8, 0xcb, 0xc5, 0x0f, 0xbe, 0x1b, 0xa9, 0x3a, 0x83,
	0x59, 0x0e, 0xf7, 0xb6, 0xde, 0x2c, 0x8b, 0x75, 0xb5, 0x3b, 0x1a, 0xd8, 0x3d, 0xd8, 0x49, 0xa7,
	0xed, 0xb4, 0xd3, 0x25, 0xc0, 0x9b, 0xef, 0x08, 0x63, 0x32, 0x89, 0xea, 0x5b, 0x11, 0xcb, 0x4c,
	0xf6, 0x19, 0x1c, 0x66, 0xf2, 0xbf, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x71, 0x9c,
	0x99, 0x76, 0x01, 0x00, 0x00,
}
