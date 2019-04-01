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
// source: l2fib_g8032_ring_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_g8032_l2fib_g8032_rings_l2fib_g8032_ring_l2fib_g8032_ring_summary

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

type L2FibG8032RingSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	RingName             string   `protobuf:"bytes,2,opt,name=ring_name,json=ringName,proto3" json:"ring_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibG8032RingSummaryInfo_KEYS) Reset()         { *m = L2FibG8032RingSummaryInfo_KEYS{} }
func (m *L2FibG8032RingSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibG8032RingSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2FibG8032RingSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca57dc90e34e9db, []int{0}
}

func (m *L2FibG8032RingSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibG8032RingSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibG8032RingSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibG8032RingSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS.Size(m)
}
func (m *L2FibG8032RingSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibG8032RingSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibG8032RingSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibG8032RingSummaryInfo_KEYS) GetRingName() string {
	if m != nil {
		return m.RingName
	}
	return ""
}

type L2FibG8032RingSummaryInfo struct {
	RingName             string   `protobuf:"bytes,50,opt,name=ring_name,json=ringName,proto3" json:"ring_name,omitempty"`
	Port0                string   `protobuf:"bytes,51,opt,name=port0,proto3" json:"port0,omitempty"`
	Port1                string   `protobuf:"bytes,52,opt,name=port1,proto3" json:"port1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibG8032RingSummaryInfo) Reset()         { *m = L2FibG8032RingSummaryInfo{} }
func (m *L2FibG8032RingSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibG8032RingSummaryInfo) ProtoMessage()    {}
func (*L2FibG8032RingSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ca57dc90e34e9db, []int{1}
}

func (m *L2FibG8032RingSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibG8032RingSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibG8032RingSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibG8032RingSummaryInfo.Merge(m, src)
}
func (m *L2FibG8032RingSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibG8032RingSummaryInfo.Size(m)
}
func (m *L2FibG8032RingSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibG8032RingSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibG8032RingSummaryInfo proto.InternalMessageInfo

func (m *L2FibG8032RingSummaryInfo) GetRingName() string {
	if m != nil {
		return m.RingName
	}
	return ""
}

func (m *L2FibG8032RingSummaryInfo) GetPort0() string {
	if m != nil {
		return m.Port0
	}
	return ""
}

func (m *L2FibG8032RingSummaryInfo) GetPort1() string {
	if m != nil {
		return m.Port1
	}
	return ""
}

func init() {
	proto.RegisterType((*L2FibG8032RingSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_g8032.l2fib_g8032_rings.l2fib_g8032_ring.l2fib_g8032_ring_summary.l2fib_g8032_ring_summary_info_KEYS")
	proto.RegisterType((*L2FibG8032RingSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_g8032.l2fib_g8032_rings.l2fib_g8032_ring.l2fib_g8032_ring_summary.l2fib_g8032_ring_summary_info")
}

func init() {
	proto.RegisterFile("l2fib_g8032_ring_summary_info.proto", fileDescriptor_9ca57dc90e34e9db)
}

var fileDescriptor_9ca57dc90e34e9db = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0x31, 0x4a, 0xcb,
	0x4c, 0x8a, 0x4f, 0xb7, 0x30, 0x30, 0x36, 0x8a, 0x2f, 0xca, 0xcc, 0x4b, 0x8f, 0x2f, 0x2e, 0xcd,
	0xcd, 0x4d, 0x2c, 0xaa, 0x8c, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xaa, 0x4e, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0x31,
	0x2a, 0x2b, 0xc8, 0x8b, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x83, 0x30, 0xd3, 0xf2, 0x8b, 0xca, 0x13,
	0x8b, 0x52, 0x32, 0xf3, 0xd2, 0xf5, 0xf2, 0xf2, 0x53, 0x52, 0x8b, 0xc1, 0xa4, 0x1e, 0x92, 0xc1,
	0x7a, 0xe8, 0x96, 0x14, 0x63, 0x88, 0xe8, 0xe1, 0x72, 0x87, 0x52, 0x14, 0x97, 0x12, 0x5e, 0x37,
	0xc6, 0x7b, 0xbb, 0x46, 0x06, 0x0b, 0x89, 0x73, 0xb1, 0x83, 0x6c, 0x8d, 0xcf, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0xa4, 0xb9, 0x38, 0xc1, 0x5a,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0x52, 0x1c, 0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0xa5,
	0x0c, 0x2e, 0x59, 0xbc, 0x66, 0xa3, 0xea, 0x36, 0x42, 0xd5, 0x2d, 0x24, 0xc2, 0xc5, 0x5a, 0x90,
	0x5f, 0x54, 0x62, 0x20, 0x61, 0x0c, 0x96, 0x80, 0x70, 0x60, 0xa2, 0x86, 0x12, 0x26, 0x08, 0x51,
	0xc3, 0x24, 0x36, 0x70, 0x48, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xdc, 0xb7, 0xb8,
	0x70, 0x01, 0x00, 0x00,
}