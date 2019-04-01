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
// source: l2fib_pwhe_mp_show_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_pwhe_main_ports_l2fib_pwhe_main_port

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

type L2FibPwheMpShowInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPwheMpShowInfo_KEYS) Reset()         { *m = L2FibPwheMpShowInfo_KEYS{} }
func (m *L2FibPwheMpShowInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibPwheMpShowInfo_KEYS) ProtoMessage()    {}
func (*L2FibPwheMpShowInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06bf72e7b39bc8, []int{0}
}

func (m *L2FibPwheMpShowInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwheMpShowInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibPwheMpShowInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwheMpShowInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibPwheMpShowInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwheMpShowInfo_KEYS.Merge(m, src)
}
func (m *L2FibPwheMpShowInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibPwheMpShowInfo_KEYS.Size(m)
}
func (m *L2FibPwheMpShowInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwheMpShowInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwheMpShowInfo_KEYS proto.InternalMessageInfo

func (m *L2FibPwheMpShowInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibPwheMpShowInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type L2FibPwheMpShowInfo struct {
	NextHopValid           bool     `protobuf:"varint,50,opt,name=next_hop_valid,json=nextHopValid,proto3" json:"next_hop_valid,omitempty"`
	NextHopAddress         string   `protobuf:"bytes,51,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	PseudoWireType         uint32   `protobuf:"varint,52,opt,name=pseudo_wire_type,json=pseudoWireType,proto3" json:"pseudo_wire_type,omitempty"`
	GenericInterfaceListId uint32   `protobuf:"varint,53,opt,name=generic_interface_list_id,json=genericInterfaceListId,proto3" json:"generic_interface_list_id,omitempty"`
	InternalLabel          uint32   `protobuf:"varint,54,opt,name=internal_label,json=internalLabel,proto3" json:"internal_label,omitempty"`
	RemoteLabel            uint32   `protobuf:"varint,55,opt,name=remote_label,json=remoteLabel,proto3" json:"remote_label,omitempty"`
	ControlWordEnabled     bool     `protobuf:"varint,56,opt,name=control_word_enabled,json=controlWordEnabled,proto3" json:"control_word_enabled,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *L2FibPwheMpShowInfo) Reset()         { *m = L2FibPwheMpShowInfo{} }
func (m *L2FibPwheMpShowInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPwheMpShowInfo) ProtoMessage()    {}
func (*L2FibPwheMpShowInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e06bf72e7b39bc8, []int{1}
}

func (m *L2FibPwheMpShowInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwheMpShowInfo.Unmarshal(m, b)
}
func (m *L2FibPwheMpShowInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwheMpShowInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPwheMpShowInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwheMpShowInfo.Merge(m, src)
}
func (m *L2FibPwheMpShowInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPwheMpShowInfo.Size(m)
}
func (m *L2FibPwheMpShowInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwheMpShowInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwheMpShowInfo proto.InternalMessageInfo

func (m *L2FibPwheMpShowInfo) GetNextHopValid() bool {
	if m != nil {
		return m.NextHopValid
	}
	return false
}

func (m *L2FibPwheMpShowInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *L2FibPwheMpShowInfo) GetPseudoWireType() uint32 {
	if m != nil {
		return m.PseudoWireType
	}
	return 0
}

func (m *L2FibPwheMpShowInfo) GetGenericInterfaceListId() uint32 {
	if m != nil {
		return m.GenericInterfaceListId
	}
	return 0
}

func (m *L2FibPwheMpShowInfo) GetInternalLabel() uint32 {
	if m != nil {
		return m.InternalLabel
	}
	return 0
}

func (m *L2FibPwheMpShowInfo) GetRemoteLabel() uint32 {
	if m != nil {
		return m.RemoteLabel
	}
	return 0
}

func (m *L2FibPwheMpShowInfo) GetControlWordEnabled() bool {
	if m != nil {
		return m.ControlWordEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*L2FibPwheMpShowInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_pwhe_main_ports.l2fib_pwhe_main_port.l2fib_pwhe_mp_show_info_KEYS")
	proto.RegisterType((*L2FibPwheMpShowInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_pwhe_main_ports.l2fib_pwhe_main_port.l2fib_pwhe_mp_show_info")
}

func init() { proto.RegisterFile("l2fib_pwhe_mp_show_info.proto", fileDescriptor_3e06bf72e7b39bc8) }

var fileDescriptor_3e06bf72e7b39bc8 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0xab, 0xd3, 0x30,
	0x14, 0xc7, 0xe9, 0x7d, 0xb8, 0x6a, 0xbc, 0x2b, 0x12, 0xc4, 0x5b, 0x41, 0x61, 0x5e, 0x14, 0xfa,
	0x54, 0x64, 0xf3, 0xe7, 0xa3, 0x0f, 0x03, 0x87, 0xc3, 0x87, 0x29, 0x0e, 0x5f, 0x76, 0x48, 0x9b,
	0xd3, 0x2d, 0x90, 0xe6, 0x84, 0x24, 0xae, 0xdb, 0xbf, 0xe8, 0x5f, 0x25, 0x4d, 0xbb, 0xb9, 0x87,
	0xbb, 0x97, 0xd2, 0x7c, 0xbe, 0x9f, 0x9c, 0x93, 0x13, 0xc2, 0x5e, 0xea, 0x49, 0xad, 0x4a, 0xb0,
	0xed, 0x16, 0xa1, 0xb1, 0xe0, 0xb7, 0xd4, 0x82, 0x32, 0x35, 0x15, 0xd6, 0x51, 0x20, 0xbe, 0xae,
	0x94, 0xaf, 0x08, 0x14, 0x79, 0xd8, 0x3b, 0xd0, 0x93, 0x9d, 0x35, 0x40, 0x16, 0x5d, 0xd1, 0xff,
	0xd6, 0xe4, 0x5a, 0xe1, 0xa4, 0x32, 0x9b, 0xc2, 0x90, 0x44, 0x1f, 0xbf, 0xc5, 0x79, 0x49, 0xa1,
	0x0c, 0x58, 0x72, 0xc1, 0xdf, 0x4b, 0xef, 0xd6, 0xec, 0xc5, 0x85, 0x03, 0xc0, 0xb7, 0xd9, 0xef,
	0x1f, 0xfc, 0x96, 0x3d, 0xe8, 0x4a, 0x82, 0x92, 0x59, 0x32, 0x4e, 0xf2, 0x47, 0xcb, 0xeb, 0x6e,
	0x39, 0x97, 0xfc, 0x0d, 0x4b, 0x95, 0x09, 0xe8, 0x6a, 0x51, 0x21, 0x18, 0xd1, 0x60, 0x76, 0x15,
	0xf3, 0xd1, 0x89, 0x7e, 0x17, 0x0d, 0xde, 0xfd, 0xbd, 0x62, 0xb7, 0x17, 0x1a, 0xf0, 0xd7, 0x2c,
	0x35, 0xb8, 0x0f, 0xb0, 0x25, 0x0b, 0x3b, 0xa1, 0x95, 0xcc, 0x26, 0xe3, 0x24, 0x7f, 0xb8, 0xbc,
	0xe9, 0xe8, 0x57, 0xb2, 0xbf, 0x3a, 0xc6, 0x73, 0xf6, 0xe4, 0x64, 0x09, 0x29, 0x1d, 0x7a, 0x9f,
	0x4d, 0x63, 0xab, 0x74, 0xf0, 0xbe, 0xf4, 0xb4, 0x33, 0xad, 0xc7, 0x3f, 0x92, 0xa0, 0x55, 0x0e,
	0x21, 0x1c, 0x2c, 0x66, 0xef, 0xc6, 0x49, 0x3e, 0x5a, 0xa6, 0x3d, 0x5f, 0x29, 0x87, 0x3f, 0x0f,
	0x16, 0xf9, 0x67, 0xf6, 0x7c, 0x83, 0x06, 0x9d, 0xaa, 0xe0, 0xff, 0x10, 0x5a, 0xf9, 0xd0, 0xcd,
	0xf9, 0x3e, 0x6e, 0x79, 0x36, 0x08, 0xf3, 0x63, 0xbe, 0x50, 0x3e, 0x9c, 0xcd, 0x6d, 0x84, 0x06,
	0x2d, 0x4a, 0xd4, 0xd9, 0x87, 0xe8, 0x8f, 0x8e, 0x74, 0xd1, 0x41, 0xfe, 0x8a, 0xdd, 0x38, 0x6c,
	0x28, 0xe0, 0x20, 0x7d, 0x8c, 0xd2, 0xe3, 0x9e, 0xf5, 0xca, 0x5b, 0xf6, 0xb4, 0x22, 0x13, 0x1c,
	0x69, 0x68, 0xc9, 0x49, 0x40, 0x23, 0x4a, 0x8d, 0x32, 0xfb, 0x14, 0x2f, 0x81, 0x0f, 0xd9, 0x8a,
	0x9c, 0x9c, 0xf5, 0x49, 0x79, 0x1d, 0xdf, 0xc4, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0x6b, 0x1b, 0xd3, 0x34, 0x02, 0x00, 0x00,
}
