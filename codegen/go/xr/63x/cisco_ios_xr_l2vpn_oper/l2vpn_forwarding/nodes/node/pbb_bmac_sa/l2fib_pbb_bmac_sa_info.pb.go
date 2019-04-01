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
// source: l2fib_pbb_bmac_sa_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_pbb_bmac_sa

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

type L2FibPbbBmacSaInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPbbBmacSaInfo_KEYS) Reset()         { *m = L2FibPbbBmacSaInfo_KEYS{} }
func (m *L2FibPbbBmacSaInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibPbbBmacSaInfo_KEYS) ProtoMessage()    {}
func (*L2FibPbbBmacSaInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_034a85b69e3ea2a3, []int{0}
}

func (m *L2FibPbbBmacSaInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibPbbBmacSaInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibPbbBmacSaInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS.Merge(m, src)
}
func (m *L2FibPbbBmacSaInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS.Size(m)
}
func (m *L2FibPbbBmacSaInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPbbBmacSaInfo_KEYS proto.InternalMessageInfo

func (m *L2FibPbbBmacSaInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type L2FibPbbBmacSaInfo struct {
	Configured           bool     `protobuf:"varint,50,opt,name=configured,proto3" json:"configured,omitempty"`
	Mac                  string   `protobuf:"bytes,51,opt,name=mac,proto3" json:"mac,omitempty"`
	ChassisMac           string   `protobuf:"bytes,52,opt,name=chassis_mac,json=chassisMac,proto3" json:"chassis_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPbbBmacSaInfo) Reset()         { *m = L2FibPbbBmacSaInfo{} }
func (m *L2FibPbbBmacSaInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPbbBmacSaInfo) ProtoMessage()    {}
func (*L2FibPbbBmacSaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_034a85b69e3ea2a3, []int{1}
}

func (m *L2FibPbbBmacSaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPbbBmacSaInfo.Unmarshal(m, b)
}
func (m *L2FibPbbBmacSaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPbbBmacSaInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPbbBmacSaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPbbBmacSaInfo.Merge(m, src)
}
func (m *L2FibPbbBmacSaInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPbbBmacSaInfo.Size(m)
}
func (m *L2FibPbbBmacSaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPbbBmacSaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPbbBmacSaInfo proto.InternalMessageInfo

func (m *L2FibPbbBmacSaInfo) GetConfigured() bool {
	if m != nil {
		return m.Configured
	}
	return false
}

func (m *L2FibPbbBmacSaInfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *L2FibPbbBmacSaInfo) GetChassisMac() string {
	if m != nil {
		return m.ChassisMac
	}
	return ""
}

func init() {
	proto.RegisterType((*L2FibPbbBmacSaInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.pbb_bmac_sa.l2fib_pbb_bmac_sa_info_KEYS")
	proto.RegisterType((*L2FibPbbBmacSaInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.pbb_bmac_sa.l2fib_pbb_bmac_sa_info")
}

func init() { proto.RegisterFile("l2fib_pbb_bmac_sa_info.proto", fileDescriptor_034a85b69e3ea2a3) }

var fileDescriptor_034a85b69e3ea2a3 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xbd, 0x4b, 0xc5, 0x40,
	0x10, 0xc4, 0x09, 0xc2, 0x53, 0xd7, 0x46, 0xae, 0xd0, 0x80, 0xa2, 0x8f, 0x54, 0xa9, 0xae, 0x48,
	0xc4, 0xd6, 0xca, 0x42, 0xc4, 0x26, 0x56, 0x56, 0xcb, 0x7d, 0xe4, 0xe2, 0x62, 0x72, 0x7b, 0xdc,
	0xf9, 0xf5, 0xe7, 0xcb, 0x9d, 0xaf, 0x48, 0x91, 0x66, 0x99, 0x9d, 0x61, 0x7e, 0xcb, 0xc2, 0xf5,
	0xdc, 0x39, 0xd2, 0x18, 0xb4, 0x46, 0xbd, 0x28, 0x83, 0x49, 0x21, 0x79, 0xc7, 0x32, 0x44, 0xfe,
	0x64, 0xf1, 0x60, 0x28, 0x19, 0x46, 0xe2, 0x84, 0xbf, 0x11, 0xe7, 0xee, 0x3b, 0x78, 0xe4, 0x30,
	0x46, 0xf9, 0x2f, 0x1d, 0xc7, 0x1f, 0x15, 0x2d, 0xf9, 0x49, 0x7a, 0xb6, 0x63, 0x2a, 0x53, 0xae,
	0x58, 0xcd, 0x3d, 0x5c, 0x6d, 0x1f, 0xc0, 0xe7, 0xc7, 0xb7, 0x57, 0x71, 0x09, 0xc7, 0xb9, 0x82,
	0x64, 0xeb, 0x6a, 0x5f, 0xb5, 0xa7, 0xc3, 0x2e, 0xaf, 0x4f, 0xb6, 0xf9, 0x80, 0x8b, 0xed, 0x9e,
	0xb8, 0x01, 0x30, 0xec, 0x1d, 0x4d, 0x5f, 0x71, 0xb4, 0x75, 0xb7, 0xaf, 0xda, 0x93, 0x61, 0xe5,
	0x88, 0x73, 0x38, 0x5a, 0x94, 0xa9, 0xfb, 0x82, 0xcb, 0x52, 0xdc, 0xc2, 0x99, 0x79, 0x57, 0x29,
	0x51, 0xc2, 0x9c, 0xdc, 0x95, 0x04, 0x0e, 0xd6, 0x8b, 0x32, 0x7a, 0x57, 0x9e, 0xed, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x82, 0x3a, 0x31, 0x48, 0x0c, 0x01, 0x00, 0x00,
}
