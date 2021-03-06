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
// source: udp_sh_stats_bag.proto

package cisco_ios_xr_ip_udp_oper_udp_connection_nodes_node_statistics_summary

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

type UdpShStatsBag_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UdpShStatsBag_KEYS) Reset()         { *m = UdpShStatsBag_KEYS{} }
func (m *UdpShStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*UdpShStatsBag_KEYS) ProtoMessage()    {}
func (*UdpShStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_adcdbbd98a7cc431, []int{0}
}

func (m *UdpShStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShStatsBag_KEYS.Unmarshal(m, b)
}
func (m *UdpShStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *UdpShStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShStatsBag_KEYS.Merge(m, src)
}
func (m *UdpShStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_UdpShStatsBag_KEYS.Size(m)
}
func (m *UdpShStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShStatsBag_KEYS proto.InternalMessageInfo

func (m *UdpShStatsBag_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type UdpShStatsBag struct {
	ReceivedTotalPackets       uint32   `protobuf:"varint,50,opt,name=received_total_packets,json=receivedTotalPackets,proto3" json:"received_total_packets,omitempty"`
	ReceivedNoPortPackets      uint32   `protobuf:"varint,51,opt,name=received_no_port_packets,json=receivedNoPortPackets,proto3" json:"received_no_port_packets,omitempty"`
	ReceivedBadChecksumPackets uint32   `protobuf:"varint,52,opt,name=received_bad_checksum_packets,json=receivedBadChecksumPackets,proto3" json:"received_bad_checksum_packets,omitempty"`
	ReceivedTooShortPackets    uint32   `protobuf:"varint,53,opt,name=received_too_short_packets,json=receivedTooShortPackets,proto3" json:"received_too_short_packets,omitempty"`
	ReceivedDropPackets        uint32   `protobuf:"varint,54,opt,name=received_drop_packets,json=receivedDropPackets,proto3" json:"received_drop_packets,omitempty"`
	SentTotalPackets           uint32   `protobuf:"varint,55,opt,name=sent_total_packets,json=sentTotalPackets,proto3" json:"sent_total_packets,omitempty"`
	SentErrorPackets           uint32   `protobuf:"varint,56,opt,name=sent_error_packets,json=sentErrorPackets,proto3" json:"sent_error_packets,omitempty"`
	ForwardBroadcastPackets    uint32   `protobuf:"varint,57,opt,name=forward_broadcast_packets,json=forwardBroadcastPackets,proto3" json:"forward_broadcast_packets,omitempty"`
	ClonedPackets              uint32   `protobuf:"varint,58,opt,name=cloned_packets,json=clonedPackets,proto3" json:"cloned_packets,omitempty"`
	FailedClonePackets         uint32   `protobuf:"varint,59,opt,name=failed_clone_packets,json=failedClonePackets,proto3" json:"failed_clone_packets,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *UdpShStatsBag) Reset()         { *m = UdpShStatsBag{} }
func (m *UdpShStatsBag) String() string { return proto.CompactTextString(m) }
func (*UdpShStatsBag) ProtoMessage()    {}
func (*UdpShStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_adcdbbd98a7cc431, []int{1}
}

func (m *UdpShStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShStatsBag.Unmarshal(m, b)
}
func (m *UdpShStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShStatsBag.Marshal(b, m, deterministic)
}
func (m *UdpShStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShStatsBag.Merge(m, src)
}
func (m *UdpShStatsBag) XXX_Size() int {
	return xxx_messageInfo_UdpShStatsBag.Size(m)
}
func (m *UdpShStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShStatsBag proto.InternalMessageInfo

func (m *UdpShStatsBag) GetReceivedTotalPackets() uint32 {
	if m != nil {
		return m.ReceivedTotalPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetReceivedNoPortPackets() uint32 {
	if m != nil {
		return m.ReceivedNoPortPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetReceivedBadChecksumPackets() uint32 {
	if m != nil {
		return m.ReceivedBadChecksumPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetReceivedTooShortPackets() uint32 {
	if m != nil {
		return m.ReceivedTooShortPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetReceivedDropPackets() uint32 {
	if m != nil {
		return m.ReceivedDropPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetSentTotalPackets() uint32 {
	if m != nil {
		return m.SentTotalPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetSentErrorPackets() uint32 {
	if m != nil {
		return m.SentErrorPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetForwardBroadcastPackets() uint32 {
	if m != nil {
		return m.ForwardBroadcastPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetClonedPackets() uint32 {
	if m != nil {
		return m.ClonedPackets
	}
	return 0
}

func (m *UdpShStatsBag) GetFailedClonePackets() uint32 {
	if m != nil {
		return m.FailedClonePackets
	}
	return 0
}

func init() {
	proto.RegisterType((*UdpShStatsBag_KEYS)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.statistics.summary.udp_sh_stats_bag_KEYS")
	proto.RegisterType((*UdpShStatsBag)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.statistics.summary.udp_sh_stats_bag")
}

func init() { proto.RegisterFile("udp_sh_stats_bag.proto", fileDescriptor_adcdbbd98a7cc431) }

var fileDescriptor_adcdbbd98a7cc431 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0xcb, 0x1a, 0x31,
	0x10, 0x87, 0x11, 0x4a, 0xe9, 0x1b, 0xb0, 0xc8, 0x56, 0xad, 0xb5, 0x14, 0x44, 0x28, 0x78, 0x28,
	0x4b, 0x51, 0x5b, 0x5b, 0x3d, 0x55, 0xeb, 0xa9, 0x20, 0xa2, 0xbd, 0xf4, 0x34, 0x64, 0x93, 0xd8,
	0x0d, 0xba, 0x99, 0x90, 0xc4, 0xfe, 0xf9, 0x42, 0xfd, 0x9c, 0x25, 0x59, 0x13, 0x7d, 0xbd, 0x2c,
	0xcb, 0xfc, 0x9e, 0x27, 0x33, 0x19, 0x42, 0xba, 0x67, 0xae, 0xc1, 0x96, 0x60, 0x1d, 0x75, 0x16,
	0x0a, 0xfa, 0x33, 0xd7, 0x06, 0x1d, 0x66, 0x6b, 0x26, 0x2d, 0x43, 0x90, 0x68, 0xe1, 0x8f, 0x01,
	0xa9, 0xc1, 0x73, 0xa8, 0x85, 0xc9, 0xfd, 0x0f, 0x43, 0xa5, 0x04, 0x73, 0x12, 0x55, 0xae, 0x90,
	0x0b, 0x1b, 0xbe, 0xb9, 0x3f, 0x43, 0x5a, 0x27, 0x99, 0xcd, 0xed, 0xb9, 0xaa, 0xa8, 0xf9, 0x3b,
	0x9c, 0x92, 0xce, 0x7d, 0x03, 0xf8, 0xb6, 0xfe, 0xb1, 0xcf, 0x5e, 0x93, 0x07, 0xef, 0x80, 0xa2,
	0x95, 0xe8, 0x35, 0x06, 0x8d, 0xd1, 0xc3, 0xee, 0x99, 0x2f, 0x6c, 0x68, 0x25, 0x86, 0xff, 0x9e,
	0x90, 0xd6, 0xbd, 0x96, 0x4d, 0x49, 0xd7, 0x08, 0x26, 0xe4, 0x2f, 0xc1, 0xc1, 0xa1, 0xa3, 0x27,
	0xd0, 0x94, 0x1d, 0x85, 0xb3, 0xbd, 0xf1, 0xa0, 0x31, 0x6a, 0xee, 0xda, 0x31, 0xfd, 0xee, 0xc3,
	0x6d, 0x9d, 0x65, 0x33, 0xd2, 0x4b, 0x96, 0x42, 0xd0, 0x68, 0x5c, 0xf2, 0x26, 0xc1, 0xeb, 0xc4,
	0x7c, 0x83, 0x5b, 0x34, 0x2e, 0x8a, 0x5f, 0xc8, 0x9b, 0x24, 0x16, 0x94, 0x03, 0x2b, 0x05, 0x3b,
	0xda, 0x73, 0x95, 0xec, 0x69, 0xb0, 0xfb, 0x11, 0x5a, 0x52, 0xbe, 0xba, 0x20, 0xf1, 0x88, 0x05,
	0xe9, 0xdf, 0x4c, 0x8c, 0x60, 0xcb, 0xdb, 0xee, 0x1f, 0x82, 0xff, 0xf2, 0x3a, 0x35, 0xee, 0xcb,
	0x9b, 0xfe, 0x63, 0x92, 0x06, 0x03, 0x6e, 0x50, 0x27, 0xef, 0x63, 0xf0, 0x5e, 0xc4, 0xf0, 0xab,
	0x41, 0x1d, 0x9d, 0x77, 0x24, 0xb3, 0x42, 0xb9, 0xbb, 0xf5, 0xcc, 0x82, 0xd0, 0xf2, 0xc9, 0xa3,
	0xd5, 0x44, 0x5a, 0x18, 0x83, 0x26, 0xd1, 0x9f, 0xae, 0xf4, 0xda, 0x07, 0x91, 0x9e, 0x93, 0x57,
	0x07, 0x34, 0xbf, 0xa9, 0xe1, 0x50, 0x18, 0xa4, 0x9c, 0x51, 0x7b, 0xbd, 0xcb, 0xe7, 0xfa, 0x2e,
	0x17, 0x60, 0x19, 0xf3, 0xe8, 0xbe, 0x25, 0xcf, 0xd9, 0x09, 0x95, 0xe0, 0x49, 0x98, 0x07, 0xa1,
	0x59, 0x57, 0x23, 0xf6, 0x9e, 0xb4, 0x0f, 0x54, 0x9e, 0x04, 0x87, 0x50, 0x4f, 0xf0, 0x22, 0xc0,
	0x59, 0x9d, 0xad, 0x7c, 0x74, 0x31, 0x8a, 0xa7, 0xe1, 0xb1, 0x4e, 0xfe, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x9f, 0x63, 0x89, 0xc6, 0x02, 0x00, 0x00,
}
