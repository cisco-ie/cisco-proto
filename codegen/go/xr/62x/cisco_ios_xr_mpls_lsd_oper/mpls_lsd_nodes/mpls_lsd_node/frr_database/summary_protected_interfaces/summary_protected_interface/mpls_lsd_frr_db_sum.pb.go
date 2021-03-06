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
// source: mpls_lsd_frr_db_sum.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_frr_database_summary_protected_interfaces_summary_protected_interface

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

type MplsLsdFrrDbSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbSum_KEYS) Reset()         { *m = MplsLsdFrrDbSum_KEYS{} }
func (m *MplsLsdFrrDbSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbSum_KEYS) ProtoMessage()    {}
func (*MplsLsdFrrDbSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c23a653836dc035, []int{0}
}

func (m *MplsLsdFrrDbSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Merge(m, src)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbSum_KEYS.Size(m)
}
func (m *MplsLsdFrrDbSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbSum_KEYS proto.InternalMessageInfo

func (m *MplsLsdFrrDbSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsLsdFrrDbSum_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLsdFrrDbSum struct {
	Active               uint32   `protobuf:"varint,50,opt,name=active,proto3" json:"active,omitempty"`
	Ready                uint32   `protobuf:"varint,51,opt,name=ready,proto3" json:"ready,omitempty"`
	Partial              uint32   `protobuf:"varint,52,opt,name=partial,proto3" json:"partial,omitempty"`
	Igp                  uint32   `protobuf:"varint,53,opt,name=igp,proto3" json:"igp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdFrrDbSum) Reset()         { *m = MplsLsdFrrDbSum{} }
func (m *MplsLsdFrrDbSum) String() string { return proto.CompactTextString(m) }
func (*MplsLsdFrrDbSum) ProtoMessage()    {}
func (*MplsLsdFrrDbSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c23a653836dc035, []int{1}
}

func (m *MplsLsdFrrDbSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdFrrDbSum.Unmarshal(m, b)
}
func (m *MplsLsdFrrDbSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdFrrDbSum.Marshal(b, m, deterministic)
}
func (m *MplsLsdFrrDbSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdFrrDbSum.Merge(m, src)
}
func (m *MplsLsdFrrDbSum) XXX_Size() int {
	return xxx_messageInfo_MplsLsdFrrDbSum.Size(m)
}
func (m *MplsLsdFrrDbSum) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdFrrDbSum.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdFrrDbSum proto.InternalMessageInfo

func (m *MplsLsdFrrDbSum) GetActive() uint32 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetReady() uint32 {
	if m != nil {
		return m.Ready
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetPartial() uint32 {
	if m != nil {
		return m.Partial
	}
	return 0
}

func (m *MplsLsdFrrDbSum) GetIgp() uint32 {
	if m != nil {
		return m.Igp
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLsdFrrDbSum_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.summary_protected_interfaces.summary_protected_interface.mpls_lsd_frr_db_sum_KEYS")
	proto.RegisterType((*MplsLsdFrrDbSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.frr_database.summary_protected_interfaces.summary_protected_interface.mpls_lsd_frr_db_sum")
}

func init() { proto.RegisterFile("mpls_lsd_frr_db_sum.proto", fileDescriptor_0c23a653836dc035) }

var fileDescriptor_0c23a653836dc035 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0x86, 0xd9, 0xef, 0xc3, 0x6a, 0x03, 0x15, 0x89, 0x22, 0x11, 0x2f, 0xa5, 0x20, 0xf4, 0xb4,
	0x07, 0xab, 0x3f, 0xc1, 0x93, 0xe0, 0xa1, 0x9e, 0xbc, 0x38, 0xcc, 0x26, 0x53, 0x09, 0x6c, 0x36,
	0x61, 0x26, 0x15, 0x7b, 0xf0, 0xbf, 0x4b, 0x52, 0x5d, 0x10, 0x8a, 0xb7, 0x3c, 0x4f, 0x5e, 0xde,
	0x19, 0x46, 0x5d, 0x85, 0xd4, 0x0b, 0xf4, 0xe2, 0x60, 0xc3, 0x0c, 0xae, 0x03, 0xd9, 0x86, 0x36,
	0x71, 0xcc, 0x51, 0x7f, 0x5a, 0x2f, 0x36, 0x82, 0x8f, 0x02, 0x1f, 0x0c, 0x63, 0x2e, 0x26, 0xe2,
	0x76, 0xa4, 0x21, 0x3a, 0x92, 0xdf, 0xd8, 0xd6, 0x26, 0xcc, 0xd8, 0xa1, 0x50, 0x2b, 0xdb, 0x10,
	0x90, 0x77, 0x50, 0x3a, 0xc9, 0x66, 0x72, 0xe0, 0x87, 0x4c, 0xbc, 0x41, 0x4b, 0xf2, 0xd7, 0xe7,
	0xe2, 0x55, 0x99, 0x03, 0xbb, 0xc1, 0xe3, 0xc3, 0xcb, 0xb3, 0xbe, 0x56, 0xd3, 0x32, 0x09, 0x06,
	0x0c, 0x64, 0x9a, 0x79, 0xb3, 0x9c, 0xae, 0x4f, 0x8a, 0x78, 0xc2, 0x40, 0xfa, 0x46, 0x9d, 0x8e,
	0x2d, 0xfb, 0xc4, 0xbf, 0x9a, 0x98, 0x8d, 0xb6, 0xc4, 0x16, 0x51, 0x9d, 0x1f, 0xe8, 0xd7, 0x97,
	0x6a, 0x82, 0x36, 0xfb, 0x77, 0x32, 0xb7, 0xf3, 0x66, 0x39, 0x5b, 0x7f, 0x93, 0xbe, 0x50, 0x47,
	0x4c, 0xe8, 0x76, 0x66, 0x55, 0xf5, 0x1e, 0xb4, 0x51, 0xc7, 0x09, 0x39, 0x7b, 0xec, 0xcd, 0x5d,
	0xf5, 0x3f, 0xa8, 0xcf, 0xd4, 0x7f, 0xff, 0x96, 0xcc, 0x7d, 0xb5, 0xe5, 0xd9, 0x4d, 0xea, 0x59,
	0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x37, 0xd7, 0x3d, 0x33, 0x73, 0x01, 0x00, 0x00,
}
