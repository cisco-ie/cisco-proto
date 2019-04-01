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
// source: pce_topology_global_info_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_pce_topology_global

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

type PceTopologyGlobalInfoBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceTopologyGlobalInfoBag_KEYS) Reset()         { *m = PceTopologyGlobalInfoBag_KEYS{} }
func (m *PceTopologyGlobalInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceTopologyGlobalInfoBag_KEYS) ProtoMessage()    {}
func (*PceTopologyGlobalInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e3fe020bf2e7f23, []int{0}
}

func (m *PceTopologyGlobalInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS.Unmarshal(m, b)
}
func (m *PceTopologyGlobalInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceTopologyGlobalInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS.Merge(m, src)
}
func (m *PceTopologyGlobalInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS.Size(m)
}
func (m *PceTopologyGlobalInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceTopologyGlobalInfoBag_KEYS proto.InternalMessageInfo

type PceNodeBag struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	IgpAreaId            uint32   `protobuf:"varint,2,opt,name=igp_area_id,json=igpAreaId,proto3" json:"igp_area_id,omitempty"`
	IgpAreaFormat        string   `protobuf:"bytes,3,opt,name=igp_area_format,json=igpAreaFormat,proto3" json:"igp_area_format,omitempty"`
	PceCost              uint32   `protobuf:"varint,4,opt,name=pce_cost,json=pceCost,proto3" json:"pce_cost,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	IsPceBorderNode      bool     `protobuf:"varint,6,opt,name=is_pce_border_node,json=isPceBorderNode,proto3" json:"is_pce_border_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceNodeBag) Reset()         { *m = PceNodeBag{} }
func (m *PceNodeBag) String() string { return proto.CompactTextString(m) }
func (*PceNodeBag) ProtoMessage()    {}
func (*PceNodeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e3fe020bf2e7f23, []int{1}
}

func (m *PceNodeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNodeBag.Unmarshal(m, b)
}
func (m *PceNodeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNodeBag.Marshal(b, m, deterministic)
}
func (m *PceNodeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNodeBag.Merge(m, src)
}
func (m *PceNodeBag) XXX_Size() int {
	return xxx_messageInfo_PceNodeBag.Size(m)
}
func (m *PceNodeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNodeBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceNodeBag proto.InternalMessageInfo

func (m *PceNodeBag) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *PceNodeBag) GetIgpAreaId() uint32 {
	if m != nil {
		return m.IgpAreaId
	}
	return 0
}

func (m *PceNodeBag) GetIgpAreaFormat() string {
	if m != nil {
		return m.IgpAreaFormat
	}
	return ""
}

func (m *PceNodeBag) GetPceCost() uint32 {
	if m != nil {
		return m.PceCost
	}
	return 0
}

func (m *PceNodeBag) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PceNodeBag) GetIsPceBorderNode() bool {
	if m != nil {
		return m.IsPceBorderNode
	}
	return false
}

type PceTopologyGlobalInfoBag struct {
	Nodes                []*PceNodeBag `protobuf:"bytes,50,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PceTopologyGlobalInfoBag) Reset()         { *m = PceTopologyGlobalInfoBag{} }
func (m *PceTopologyGlobalInfoBag) String() string { return proto.CompactTextString(m) }
func (*PceTopologyGlobalInfoBag) ProtoMessage()    {}
func (*PceTopologyGlobalInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e3fe020bf2e7f23, []int{2}
}

func (m *PceTopologyGlobalInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTopologyGlobalInfoBag.Unmarshal(m, b)
}
func (m *PceTopologyGlobalInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTopologyGlobalInfoBag.Marshal(b, m, deterministic)
}
func (m *PceTopologyGlobalInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTopologyGlobalInfoBag.Merge(m, src)
}
func (m *PceTopologyGlobalInfoBag) XXX_Size() int {
	return xxx_messageInfo_PceTopologyGlobalInfoBag.Size(m)
}
func (m *PceTopologyGlobalInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTopologyGlobalInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceTopologyGlobalInfoBag proto.InternalMessageInfo

func (m *PceTopologyGlobalInfoBag) GetNodes() []*PceNodeBag {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func init() {
	proto.RegisterType((*PceTopologyGlobalInfoBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.topology.global.pce_topology_global_info_bag_KEYS")
	proto.RegisterType((*PceNodeBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.topology.global.pce_node_bag")
	proto.RegisterType((*PceTopologyGlobalInfoBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.topology.global.pce_topology_global_info_bag")
}

func init() { proto.RegisterFile("pce_topology_global_info_bag.proto", fileDescriptor_3e3fe020bf2e7f23) }

var fileDescriptor_3e3fe020bf2e7f23 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x80, 0x59, 0x6b, 0xff, 0x52, 0x4b, 0x21, 0x17, 0x57, 0x14, 0x59, 0x57, 0x90, 0x05, 0x21,
	0x87, 0xfa, 0x02, 0xfe, 0xa0, 0x50, 0x04, 0x91, 0x15, 0x04, 0x4f, 0x43, 0x36, 0x99, 0x2e, 0x81,
	0x6d, 0x27, 0x24, 0x51, 0xf4, 0x3d, 0x7d, 0x20, 0x49, 0xaa, 0xe2, 0xa9, 0xe0, 0x71, 0xbe, 0x99,
	0xf9, 0x32, 0x33, 0x61, 0xa5, 0x55, 0x08, 0x81, 0x2c, 0x75, 0xd4, 0x7e, 0x40, 0xdb, 0x51, 0x23,
	0x3b, 0x30, 0xeb, 0x25, 0x41, 0x23, 0x5b, 0x61, 0x1d, 0x05, 0xe2, 0x73, 0x65, 0xbc, 0x22, 0x30,
	0xe4, 0xe1, 0xdd, 0xc1, 0xca, 0x76, 0x1e, 0x02, 0x02, 0x59, 0x74, 0x22, 0x05, 0x56, 0xa1, 0xf8,
	0x51, 0x88, 0x8d, 0xa2, 0x3c, 0x65, 0x27, 0xdb, 0xcc, 0x70, 0x7f, 0xfb, 0xf2, 0x54, 0x7e, 0x66,
	0x6c, 0x2f, 0x56, 0xad, 0x49, 0x63, 0xa4, 0xfc, 0x90, 0x8d, 0x1d, 0xbd, 0x06, 0x74, 0x60, 0x74,
	0x9e, 0x15, 0x59, 0x35, 0xae, 0x47, 0x1b, 0xb0, 0xd0, 0xfc, 0x98, 0x4d, 0x4c, 0x6b, 0x41, 0x3a,
	0x94, 0x31, 0xbd, 0x53, 0x64, 0xd5, 0xb4, 0x1e, 0x9b, 0xd6, 0x5e, 0x39, 0x94, 0x0b, 0xcd, 0xcf,
	0xd8, 0xec, 0x37, 0xbf, 0x24, 0xb7, 0x92, 0x21, 0xef, 0x25, 0xc5, 0xf4, 0xbb, 0xe6, 0x2e, 0x41,
	0x7e, 0xc0, 0x46, 0xf1, 0x51, 0x45, 0x3e, 0xe4, 0xbb, 0x49, 0x32, 0xb4, 0x0a, 0x6f, 0xc8, 0x07,
	0xbe, 0xcf, 0x86, 0x69, 0x16, 0xa3, 0xf3, 0x7e, 0x6a, 0x1d, 0xc4, 0x70, 0xa1, 0xf9, 0x39, 0xe3,
	0x26, 0x2d, 0x0a, 0x0d, 0x39, 0x8d, 0x2e, 0x8d, 0x9c, 0x0f, 0x8a, 0xac, 0x1a, 0xd5, 0x33, 0xe3,
	0x1f, 0x15, 0x5e, 0x27, 0xfe, 0x40, 0x1a, 0xcb, 0x37, 0x76, 0xb4, 0x6d, 0x77, 0xfe, 0xcc, 0xfa,
	0xb1, 0xdd, 0xe7, 0xf3, 0xa2, 0x57, 0x4d, 0xe6, 0x97, 0xe2, 0xff, 0xf7, 0x15, 0x7f, 0xcf, 0x56,
	0x6f, 0x74, 0xcd, 0x20, 0x7d, 0xd7, 0xc5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x48, 0x2a,
	0x2a, 0xd4, 0x01, 0x00, 0x00,
}