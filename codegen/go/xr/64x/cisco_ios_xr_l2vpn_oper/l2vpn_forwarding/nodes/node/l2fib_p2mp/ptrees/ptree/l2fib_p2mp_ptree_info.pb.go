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
// source: l2fib_p2mp_ptree_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_p2mp_ptrees_ptree

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

type L2FibP2MpPtreeInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PtreeType            string   `protobuf:"bytes,2,opt,name=ptree_type,json=ptreeType,proto3" json:"ptree_type,omitempty"`
	LsmId                uint32   `protobuf:"varint,3,opt,name=lsm_id,json=lsmId,proto3" json:"lsm_id,omitempty"`
	TunnelId             uint32   `protobuf:"varint,4,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	P2MpId               uint32   `protobuf:"varint,5,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	ExtendedTunnelId     uint32   `protobuf:"varint,6,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibP2MpPtreeInfo_KEYS) Reset()         { *m = L2FibP2MpPtreeInfo_KEYS{} }
func (m *L2FibP2MpPtreeInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibP2MpPtreeInfo_KEYS) ProtoMessage()    {}
func (*L2FibP2MpPtreeInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2389d832b2b54910, []int{0}
}

func (m *L2FibP2MpPtreeInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibP2MpPtreeInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibP2MpPtreeInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS.Merge(m, src)
}
func (m *L2FibP2MpPtreeInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS.Size(m)
}
func (m *L2FibP2MpPtreeInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibP2MpPtreeInfo_KEYS proto.InternalMessageInfo

func (m *L2FibP2MpPtreeInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibP2MpPtreeInfo_KEYS) GetPtreeType() string {
	if m != nil {
		return m.PtreeType
	}
	return ""
}

func (m *L2FibP2MpPtreeInfo_KEYS) GetLsmId() uint32 {
	if m != nil {
		return m.LsmId
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo_KEYS) GetExtendedTunnelId() uint32 {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return 0
}

type L2FibP2MpPtreeInfo struct {
	LsmIdXr              uint32   `protobuf:"varint,50,opt,name=lsm_id_xr,json=lsmIdXr,proto3" json:"lsm_id_xr,omitempty"`
	TunnelIdXr           uint32   `protobuf:"varint,51,opt,name=tunnel_id_xr,json=tunnelIdXr,proto3" json:"tunnel_id_xr,omitempty"`
	P2MpIdXr             uint32   `protobuf:"varint,52,opt,name=p2mp_id_xr,json=p2mpIdXr,proto3" json:"p2mp_id_xr,omitempty"`
	ExtendedTunnelIdXr   string   `protobuf:"bytes,53,opt,name=extended_tunnel_id_xr,json=extendedTunnelIdXr,proto3" json:"extended_tunnel_id_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibP2MpPtreeInfo) Reset()         { *m = L2FibP2MpPtreeInfo{} }
func (m *L2FibP2MpPtreeInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibP2MpPtreeInfo) ProtoMessage()    {}
func (*L2FibP2MpPtreeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2389d832b2b54910, []int{1}
}

func (m *L2FibP2MpPtreeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibP2MpPtreeInfo.Unmarshal(m, b)
}
func (m *L2FibP2MpPtreeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibP2MpPtreeInfo.Marshal(b, m, deterministic)
}
func (m *L2FibP2MpPtreeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibP2MpPtreeInfo.Merge(m, src)
}
func (m *L2FibP2MpPtreeInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibP2MpPtreeInfo.Size(m)
}
func (m *L2FibP2MpPtreeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibP2MpPtreeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibP2MpPtreeInfo proto.InternalMessageInfo

func (m *L2FibP2MpPtreeInfo) GetLsmIdXr() uint32 {
	if m != nil {
		return m.LsmIdXr
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo) GetTunnelIdXr() uint32 {
	if m != nil {
		return m.TunnelIdXr
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo) GetP2MpIdXr() uint32 {
	if m != nil {
		return m.P2MpIdXr
	}
	return 0
}

func (m *L2FibP2MpPtreeInfo) GetExtendedTunnelIdXr() string {
	if m != nil {
		return m.ExtendedTunnelIdXr
	}
	return ""
}

func init() {
	proto.RegisterType((*L2FibP2MpPtreeInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_p2mp.ptrees.ptree.l2fib_p2mp_ptree_info_KEYS")
	proto.RegisterType((*L2FibP2MpPtreeInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_p2mp.ptrees.ptree.l2fib_p2mp_ptree_info")
}

func init() { proto.RegisterFile("l2fib_p2mp_ptree_info.proto", fileDescriptor_2389d832b2b54910) }

var fileDescriptor_2389d832b2b54910 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xa9, 0xba, 0xce, 0x1e, 0x14, 0xe4, 0xc0, 0xb0, 0x6c, 0x0a, 0x65, 0x57, 0xbb, 0x90,
	0x82, 0x9d, 0x3e, 0x82, 0x17, 0x65, 0x77, 0x73, 0x17, 0xf5, 0x2a, 0xb8, 0x25, 0x95, 0x40, 0x9b,
	0x84, 0xa4, 0x6a, 0xf7, 0x42, 0xbe, 0x8e, 0xaf, 0x24, 0x39, 0xd9, 0x1c, 0xe8, 0x6e, 0xda, 0x26,
	0xdf, 0xf9, 0xff, 0x7c, 0x6d, 0x61, 0xd2, 0x14, 0xb5, 0x5c, 0x33, 0x53, 0xb4, 0x86, 0x99, 0xce,
	0x0a, 0xc1, 0xa4, 0xaa, 0x75, 0x6e, 0xac, 0xee, 0x34, 0x2e, 0x36, 0xd2, 0x6d, 0x34, 0x93, 0xda,
	0xb1, 0xde, 0xb2, 0xa6, 0xf8, 0x30, 0x8a, 0x69, 0x23, 0x6c, 0x1e, 0x1e, 0x6b, 0x6d, 0x3f, 0x5f,
	0x2d, 0x97, 0xea, 0x2d, 0x57, 0x9a, 0x0b, 0x47, 0xd7, 0xfc, 0x50, 0x98, 0x53, 0xa1, 0x0b, 0xb7,
	0xe9, 0x77, 0x04, 0xe3, 0xa3, 0x87, 0xb1, 0xc5, 0xd3, 0xcb, 0x33, 0x5e, 0xc3, 0xd0, 0xc7, 0x99,
	0xe4, 0x69, 0x94, 0x45, 0xb3, 0x64, 0x19, 0xfb, 0x65, 0xc9, 0xf1, 0x16, 0x20, 0xcc, 0x76, 0x5b,
	0x23, 0xd2, 0x13, 0x62, 0x09, 0xed, 0xac, 0xb6, 0x46, 0xe0, 0x08, 0xe2, 0xc6, 0xb5, 0x3e, 0x76,
	0x9a, 0x45, 0xb3, 0xcb, 0xe5, 0xa0, 0x71, 0x6d, 0xc9, 0x71, 0x02, 0x49, 0xf7, 0xae, 0x94, 0x68,
	0x3c, 0x39, 0x23, 0x72, 0x1e, 0x36, 0x4a, 0xee, 0xcf, 0x22, 0x07, 0xc9, 0xd3, 0x01, 0xa1, 0xd8,
	0x2f, 0x4b, 0x8e, 0x77, 0x80, 0xa2, 0xef, 0x84, 0xe2, 0x82, 0xb3, 0x43, 0x3c, 0xa6, 0x99, 0xab,
	0x3d, 0x59, 0xed, 0x6a, 0xa6, 0x5f, 0x11, 0x8c, 0x8e, 0xbe, 0x11, 0x8e, 0x21, 0x09, 0x52, 0xac,
	0xb7, 0x69, 0x41, 0xf1, 0x21, 0x79, 0x55, 0x16, 0x33, 0xb8, 0xf8, 0xad, 0xf6, 0x78, 0x4e, 0x18,
	0xf6, 0x72, 0x95, 0xc5, 0x1b, 0x80, 0x9d, 0x9e, 0xe7, 0x0f, 0x41, 0x3e, 0x18, 0x56, 0x16, 0xef,
	0x61, 0xf4, 0xdf, 0xd1, 0x0f, 0x3e, 0xd2, 0xa7, 0xc1, 0xbf, 0x9a, 0x95, 0x5d, 0xc7, 0xf4, 0x3b,
	0xe7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0xc5, 0x0f, 0x89, 0xed, 0x01, 0x00, 0x00,
}
