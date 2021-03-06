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
// source: bfd_mgmt_packet_count.proto

package cisco_ios_xr_ip_bfd_oper_bfd_ipv4_single_hop_counters_ipv4_single_hop_packet_counters_ipv4_single_hop_packet_counter

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

type BfdMgmtPacketCount_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtPacketCount_KEYS) Reset()         { *m = BfdMgmtPacketCount_KEYS{} }
func (m *BfdMgmtPacketCount_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtPacketCount_KEYS) ProtoMessage()    {}
func (*BfdMgmtPacketCount_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_448f2d88deed3ba8, []int{0}
}

func (m *BfdMgmtPacketCount_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtPacketCount_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtPacketCount_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtPacketCount_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtPacketCount_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtPacketCount_KEYS.Merge(m, src)
}
func (m *BfdMgmtPacketCount_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtPacketCount_KEYS.Size(m)
}
func (m *BfdMgmtPacketCount_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtPacketCount_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtPacketCount_KEYS proto.InternalMessageInfo

func (m *BfdMgmtPacketCount_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type BfdMgmtPacketCount struct {
	HelloTransmitCount   uint32   `protobuf:"varint,50,opt,name=hello_transmit_count,json=helloTransmitCount,proto3" json:"hello_transmit_count,omitempty"`
	HelloReceiveCount    uint32   `protobuf:"varint,51,opt,name=hello_receive_count,json=helloReceiveCount,proto3" json:"hello_receive_count,omitempty"`
	EchoTransmitCount    uint32   `protobuf:"varint,52,opt,name=echo_transmit_count,json=echoTransmitCount,proto3" json:"echo_transmit_count,omitempty"`
	EchoReceiveCount     uint32   `protobuf:"varint,53,opt,name=echo_receive_count,json=echoReceiveCount,proto3" json:"echo_receive_count,omitempty"`
	DisplayType          string   `protobuf:"bytes,54,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtPacketCount) Reset()         { *m = BfdMgmtPacketCount{} }
func (m *BfdMgmtPacketCount) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtPacketCount) ProtoMessage()    {}
func (*BfdMgmtPacketCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_448f2d88deed3ba8, []int{1}
}

func (m *BfdMgmtPacketCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtPacketCount.Unmarshal(m, b)
}
func (m *BfdMgmtPacketCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtPacketCount.Marshal(b, m, deterministic)
}
func (m *BfdMgmtPacketCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtPacketCount.Merge(m, src)
}
func (m *BfdMgmtPacketCount) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtPacketCount.Size(m)
}
func (m *BfdMgmtPacketCount) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtPacketCount.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtPacketCount proto.InternalMessageInfo

func (m *BfdMgmtPacketCount) GetHelloTransmitCount() uint32 {
	if m != nil {
		return m.HelloTransmitCount
	}
	return 0
}

func (m *BfdMgmtPacketCount) GetHelloReceiveCount() uint32 {
	if m != nil {
		return m.HelloReceiveCount
	}
	return 0
}

func (m *BfdMgmtPacketCount) GetEchoTransmitCount() uint32 {
	if m != nil {
		return m.EchoTransmitCount
	}
	return 0
}

func (m *BfdMgmtPacketCount) GetEchoReceiveCount() uint32 {
	if m != nil {
		return m.EchoReceiveCount
	}
	return 0
}

func (m *BfdMgmtPacketCount) GetDisplayType() string {
	if m != nil {
		return m.DisplayType
	}
	return ""
}

func init() {
	proto.RegisterType((*BfdMgmtPacketCount_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_single_hop_counters.ipv4_single_hop_packet_counters.ipv4_single_hop_packet_counter.bfd_mgmt_packet_count_KEYS")
	proto.RegisterType((*BfdMgmtPacketCount)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_single_hop_counters.ipv4_single_hop_packet_counters.ipv4_single_hop_packet_counter.bfd_mgmt_packet_count")
}

func init() { proto.RegisterFile("bfd_mgmt_packet_count.proto", fileDescriptor_448f2d88deed3ba8) }

var fileDescriptor_448f2d88deed3ba8 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0x45, 0x30, 0x3a, 0xd1, 0xa8, 0x50, 0xf4, 0x32, 0x07, 0xc2, 0x0e, 0x52, 0xc4,
	0x4d, 0xff, 0xc0, 0xf0, 0x24, 0x78, 0xa8, 0xbb, 0x78, 0xfa, 0x48, 0xd3, 0xaf, 0x6b, 0xb0, 0x69,
	0x3e, 0x92, 0x38, 0xec, 0x5f, 0xf7, 0x24, 0x4d, 0xcb, 0xb0, 0xac, 0xb0, 0xeb, 0xfb, 0x3c, 0xef,
	0x9b, 0x84, 0xb0, 0xdb, 0xac, 0xc8, 0x41, 0x6f, 0xb4, 0x07, 0x12, 0xf2, 0x0b, 0x3d, 0x48, 0xf3,
	0x5d, 0xfb, 0x84, 0xac, 0xf1, 0x86, 0x7b, 0xa9, 0x9c, 0x34, 0xa0, 0x8c, 0x83, 0x1f, 0x0b, 0x8a,
	0xa0, 0x95, 0x0d, 0xa1, 0x4d, 0xb2, 0x22, 0x4f, 0x14, 0x6d, 0x97, 0xe0, 0x54, 0xbd, 0xa9, 0x10,
	0x4a, 0x43, 0x5d, 0x13, 0xad, 0xdb, 0x03, 0xff, 0x97, 0x0f, 0xf3, 0xd9, 0x8a, 0xdd, 0x8c, 0x5e,
	0x0a, 0xde, 0x5e, 0x3f, 0x3f, 0xf8, 0x3d, 0x3b, 0x53, 0xad, 0x56, 0x08, 0x89, 0x50, 0x0b, 0x8d,
	0x71, 0x34, 0x8d, 0xe6, 0xc7, 0xe9, 0x64, 0x97, 0xbe, 0x0b, 0x8d, 0xb3, 0xdf, 0x88, 0x5d, 0x8f,
	0xae, 0xf0, 0x47, 0x76, 0x55, 0x62, 0x55, 0x19, 0xf0, 0x56, 0xd4, 0x4e, 0xab, 0x3e, 0x8f, 0x9f,
	0xa6, 0xd1, 0x7c, 0x92, 0xf2, 0xc0, 0xd6, 0x3d, 0x5a, 0x85, 0x46, 0xc2, 0x2e, 0xbb, 0x86, 0x45,
	0x89, 0x6a, 0x8b, 0x7d, 0x61, 0x11, 0x0a, 0x17, 0x01, 0xa5, 0x1d, 0xd9, 0xf9, 0x28, 0xcb, 0xbd,
	0x03, 0x96, 0x9d, 0xdf, 0xa2, 0xe1, 0xfe, 0x03, 0xe3, 0xc1, 0x1f, 0xce, 0x3f, 0x07, 0xfd, 0xbc,
	0x25, 0x83, 0xf5, 0x3b, 0x76, 0x9a, 0x2b, 0x47, 0x95, 0x68, 0xc0, 0x37, 0x84, 0xf1, 0x4b, 0x78,
	0xfe, 0x49, 0x9f, 0xad, 0x1b, 0xc2, 0xec, 0x28, 0x7c, 0xdf, 0xe2, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x94, 0x76, 0x65, 0xbc, 0xdd, 0x01, 0x00, 0x00,
}
