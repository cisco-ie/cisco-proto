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
// source: ip_udp.proto

package cisco_ios_xr_ip_udp_oper_udp_nodes_node_statistics_ipv6_traffic

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

type IpUdp_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpUdp_KEYS) Reset()         { *m = IpUdp_KEYS{} }
func (m *IpUdp_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpUdp_KEYS) ProtoMessage()    {}
func (*IpUdp_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_df82d98c5823780f, []int{0}
}

func (m *IpUdp_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpUdp_KEYS.Unmarshal(m, b)
}
func (m *IpUdp_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpUdp_KEYS.Marshal(b, m, deterministic)
}
func (m *IpUdp_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpUdp_KEYS.Merge(m, src)
}
func (m *IpUdp_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpUdp_KEYS.Size(m)
}
func (m *IpUdp_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpUdp_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpUdp_KEYS proto.InternalMessageInfo

func (m *IpUdp_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type IpUdp struct {
	UdpInputPackets         uint32   `protobuf:"varint,50,opt,name=udp_input_packets,json=udpInputPackets,proto3" json:"udp_input_packets,omitempty"`
	UdpChecksumErrorPackets uint32   `protobuf:"varint,51,opt,name=udp_checksum_error_packets,json=udpChecksumErrorPackets,proto3" json:"udp_checksum_error_packets,omitempty"`
	UdpNoPortPackets        uint32   `protobuf:"varint,52,opt,name=udp_no_port_packets,json=udpNoPortPackets,proto3" json:"udp_no_port_packets,omitempty"`
	UdpBadLengthPackets     uint32   `protobuf:"varint,53,opt,name=udp_bad_length_packets,json=udpBadLengthPackets,proto3" json:"udp_bad_length_packets,omitempty"`
	UdpOutputPackets        uint32   `protobuf:"varint,54,opt,name=udp_output_packets,json=udpOutputPackets,proto3" json:"udp_output_packets,omitempty"`
	UdpDroppedPackets       uint32   `protobuf:"varint,55,opt,name=udp_dropped_packets,json=udpDroppedPackets,proto3" json:"udp_dropped_packets,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *IpUdp) Reset()         { *m = IpUdp{} }
func (m *IpUdp) String() string { return proto.CompactTextString(m) }
func (*IpUdp) ProtoMessage()    {}
func (*IpUdp) Descriptor() ([]byte, []int) {
	return fileDescriptor_df82d98c5823780f, []int{1}
}

func (m *IpUdp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpUdp.Unmarshal(m, b)
}
func (m *IpUdp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpUdp.Marshal(b, m, deterministic)
}
func (m *IpUdp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpUdp.Merge(m, src)
}
func (m *IpUdp) XXX_Size() int {
	return xxx_messageInfo_IpUdp.Size(m)
}
func (m *IpUdp) XXX_DiscardUnknown() {
	xxx_messageInfo_IpUdp.DiscardUnknown(m)
}

var xxx_messageInfo_IpUdp proto.InternalMessageInfo

func (m *IpUdp) GetUdpInputPackets() uint32 {
	if m != nil {
		return m.UdpInputPackets
	}
	return 0
}

func (m *IpUdp) GetUdpChecksumErrorPackets() uint32 {
	if m != nil {
		return m.UdpChecksumErrorPackets
	}
	return 0
}

func (m *IpUdp) GetUdpNoPortPackets() uint32 {
	if m != nil {
		return m.UdpNoPortPackets
	}
	return 0
}

func (m *IpUdp) GetUdpBadLengthPackets() uint32 {
	if m != nil {
		return m.UdpBadLengthPackets
	}
	return 0
}

func (m *IpUdp) GetUdpOutputPackets() uint32 {
	if m != nil {
		return m.UdpOutputPackets
	}
	return 0
}

func (m *IpUdp) GetUdpDroppedPackets() uint32 {
	if m != nil {
		return m.UdpDroppedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*IpUdp_KEYS)(nil), "cisco_ios_xr_ip_udp_oper.udp.nodes.node.statistics.ipv6_traffic.ip_udp_KEYS")
	proto.RegisterType((*IpUdp)(nil), "cisco_ios_xr_ip_udp_oper.udp.nodes.node.statistics.ipv6_traffic.ip_udp")
}

func init() { proto.RegisterFile("ip_udp.proto", fileDescriptor_df82d98c5823780f) }

var fileDescriptor_df82d98c5823780f = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x3d, 0x4f, 0x02, 0x41,
	0x10, 0x86, 0x83, 0x05, 0x91, 0x55, 0xa3, 0x9e, 0x89, 0x12, 0x6d, 0x08, 0x15, 0x21, 0x7a, 0x85,
	0x28, 0x16, 0x16, 0x26, 0x2a, 0x85, 0xd1, 0x20, 0xc1, 0xca, 0x6a, 0xb2, 0xec, 0x2e, 0xb2, 0x41,
	0x6e, 0x26, 0xfb, 0x61, 0xfc, 0x5f, 0xfe, 0x41, 0x33, 0x7b, 0x1f, 0xda, 0x5c, 0x31, 0xef, 0xf3,
	0xdc, 0xbb, 0x3b, 0x2b, 0x76, 0x2d, 0x41, 0xd4, 0x94, 0x93, 0xc3, 0x80, 0xd9, 0x9d, 0xb2, 0x5e,
	0x21, 0x58, 0xf4, 0xf0, 0xed, 0xa0, 0x8c, 0x00, 0xc9, 0xb8, 0x9c, 0x99, 0x02, 0xb5, 0xf1, 0xe9,
	0x9b, 0xfb, 0x20, 0x83, 0xf5, 0xc1, 0x2a, 0x9f, 0x5b, 0xfa, 0x1a, 0x43, 0x70, 0x72, 0xb9, 0xb4,
	0xaa, 0x3f, 0x14, 0x3b, 0x95, 0xf5, 0x3c, 0x79, 0x7f, 0xcb, 0xce, 0x44, 0x87, 0x71, 0x28, 0xe4,
	0xc6, 0x74, 0x5b, 0xbd, 0xd6, 0xa0, 0x33, 0xdf, 0xe6, 0xc1, 0x54, 0x6e, 0x4c, 0xff, 0x67, 0x4b,
	0xb4, 0x4b, 0x38, 0x1b, 0x8a, 0x43, 0x76, 0x6c, 0x41, 0x31, 0x00, 0x49, 0xb5, 0x36, 0xc1, 0x77,
	0x2f, 0x7b, 0xad, 0xc1, 0xde, 0x7c, 0x3f, 0x6a, 0x7a, 0xe2, 0xf9, 0xac, 0x1c, 0x67, 0xb7, 0xe2,
	0x94, 0x59, 0xb5, 0x32, 0x6a, 0xed, 0xe3, 0x06, 0x8c, 0x73, 0xe8, 0x1a, 0x69, 0x94, 0xa4, 0x93,
	0xa8, 0xe9, 0xa1, 0x02, 0x26, 0x9c, 0xd7, 0xf2, 0x85, 0x38, 0x62, 0xb9, 0x40, 0x20, 0x74, 0x7f,
	0x55, 0x57, 0xc9, 0x3a, 0x88, 0x9a, 0xa6, 0x38, 0x43, 0xd7, 0x74, 0x8d, 0xc4, 0x31, 0xe3, 0x0b,
	0xa9, 0xe1, 0xd3, 0x14, 0x1f, 0x61, 0xd5, 0x18, 0xd7, 0xc9, 0xe0, 0x9f, 0xdd, 0x4b, 0xfd, 0x92,
	0xb2, 0x5a, 0x3a, 0x17, 0x59, 0x5a, 0x5b, 0x0c, 0xff, 0x6f, 0x33, 0x6e, 0x2a, 0x5e, 0x53, 0x50,
	0xd3, 0x79, 0x79, 0x22, 0xed, 0x90, 0xc8, 0xe8, 0x06, 0xbf, 0x49, 0x38, 0x6f, 0xe5, 0xb1, 0x4c,
	0x2a, 0x7e, 0xd1, 0x4e, 0x2f, 0x35, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x8b, 0x2d, 0xc0,
	0xb9, 0x01, 0x00, 0x00,
}
