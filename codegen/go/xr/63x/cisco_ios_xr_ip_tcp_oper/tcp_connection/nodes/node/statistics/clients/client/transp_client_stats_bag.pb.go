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
// source: transp_client_stats_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_connection_nodes_node_statistics_clients_client

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

type TranspClientStatsBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId             uint32   `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranspClientStatsBag_KEYS) Reset()         { *m = TranspClientStatsBag_KEYS{} }
func (m *TranspClientStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TranspClientStatsBag_KEYS) ProtoMessage()    {}
func (*TranspClientStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1c4979b5ee7f929, []int{0}
}

func (m *TranspClientStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranspClientStatsBag_KEYS.Unmarshal(m, b)
}
func (m *TranspClientStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranspClientStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TranspClientStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranspClientStatsBag_KEYS.Merge(m, src)
}
func (m *TranspClientStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TranspClientStatsBag_KEYS.Size(m)
}
func (m *TranspClientStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TranspClientStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TranspClientStatsBag_KEYS proto.InternalMessageInfo

func (m *TranspClientStatsBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TranspClientStatsBag_KEYS) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type TranspClientStatsBag struct {
	ClientJid            int32    `protobuf:"zigzag32,50,opt,name=client_jid,json=clientJid,proto3" json:"client_jid,omitempty"`
	ClientName           string   `protobuf:"bytes,51,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	Ipv4ReceivedPackets  uint32   `protobuf:"varint,52,opt,name=ipv4_received_packets,json=ipv4ReceivedPackets,proto3" json:"ipv4_received_packets,omitempty"`
	Ipv4SentPackets      uint32   `protobuf:"varint,53,opt,name=ipv4_sent_packets,json=ipv4SentPackets,proto3" json:"ipv4_sent_packets,omitempty"`
	Ipv6ReceivedPackets  uint32   `protobuf:"varint,54,opt,name=ipv6_received_packets,json=ipv6ReceivedPackets,proto3" json:"ipv6_received_packets,omitempty"`
	Ipv6SentPackets      uint32   `protobuf:"varint,55,opt,name=ipv6_sent_packets,json=ipv6SentPackets,proto3" json:"ipv6_sent_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranspClientStatsBag) Reset()         { *m = TranspClientStatsBag{} }
func (m *TranspClientStatsBag) String() string { return proto.CompactTextString(m) }
func (*TranspClientStatsBag) ProtoMessage()    {}
func (*TranspClientStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1c4979b5ee7f929, []int{1}
}

func (m *TranspClientStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranspClientStatsBag.Unmarshal(m, b)
}
func (m *TranspClientStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranspClientStatsBag.Marshal(b, m, deterministic)
}
func (m *TranspClientStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranspClientStatsBag.Merge(m, src)
}
func (m *TranspClientStatsBag) XXX_Size() int {
	return xxx_messageInfo_TranspClientStatsBag.Size(m)
}
func (m *TranspClientStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TranspClientStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TranspClientStatsBag proto.InternalMessageInfo

func (m *TranspClientStatsBag) GetClientJid() int32 {
	if m != nil {
		return m.ClientJid
	}
	return 0
}

func (m *TranspClientStatsBag) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *TranspClientStatsBag) GetIpv4ReceivedPackets() uint32 {
	if m != nil {
		return m.Ipv4ReceivedPackets
	}
	return 0
}

func (m *TranspClientStatsBag) GetIpv4SentPackets() uint32 {
	if m != nil {
		return m.Ipv4SentPackets
	}
	return 0
}

func (m *TranspClientStatsBag) GetIpv6ReceivedPackets() uint32 {
	if m != nil {
		return m.Ipv6ReceivedPackets
	}
	return 0
}

func (m *TranspClientStatsBag) GetIpv6SentPackets() uint32 {
	if m != nil {
		return m.Ipv6SentPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*TranspClientStatsBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.clients.client.transp_client_stats_bag_KEYS")
	proto.RegisterType((*TranspClientStatsBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.clients.client.transp_client_stats_bag")
}

func init() { proto.RegisterFile("transp_client_stats_bag.proto", fileDescriptor_c1c4979b5ee7f929) }

var fileDescriptor_c1c4979b5ee7f929 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd1, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0x07, 0x70, 0xda, 0xc3, 0x8f, 0xdf, 0x22, 0x2a, 0x8b, 0x88, 0x05, 0x1d, 0x96, 0x9d, 0x8a,
	0x87, 0x1e, 0xb6, 0x59, 0xff, 0x02, 0x0f, 0x3a, 0x11, 0xe9, 0x4e, 0x9e, 0x1e, 0x59, 0xf2, 0x90,
	0xa7, 0x2e, 0x09, 0x4d, 0x18, 0xfe, 0x2b, 0xfe, 0xb7, 0x92, 0x34, 0x1d, 0x8a, 0xee, 0xd2, 0x96,
	0x7c, 0xbf, 0xef, 0xf3, 0x0a, 0x61, 0x13, 0xdf, 0x09, 0xed, 0x2c, 0xc8, 0x77, 0x42, 0xed, 0xc1,
	0x79, 0xe1, 0x1d, 0xac, 0xc5, 0x4b, 0x6d, 0x3b, 0xe3, 0x0d, 0x7f, 0x90, 0xe4, 0xa4, 0x01, 0x32,
	0x0e, 0x3e, 0x3a, 0x20, 0x0b, 0x5e, 0x5a, 0x30, 0x16, 0xbb, 0x3a, 0x7c, 0x48, 0xa3, 0x35, 0x4a,
	0x4f, 0x46, 0xd7, 0xda, 0x28, 0x74, 0xf1, 0x59, 0x07, 0x83, 0x9c, 0x27, 0xe9, 0xea, 0x5e, 0x1d,
	0xde, 0xd3, 0x25, 0xbb, 0xd8, 0xb3, 0x0e, 0x96, 0xb7, 0xcf, 0x2b, 0x7e, 0xc4, 0x72, 0x52, 0x45,
	0x56, 0x66, 0xd5, 0xa8, 0xcd, 0x49, 0xf1, 0x73, 0x36, 0x4a, 0x45, 0x52, 0x45, 0x5e, 0x66, 0xd5,
	0x61, 0xfb, 0xbf, 0x3f, 0xb8, 0x53, 0xd3, 0xcf, 0x9c, 0x9d, 0xed, 0xd1, 0xf8, 0x84, 0xb1, 0x74,
	0xf6, 0x4a, 0xaa, 0x98, 0x95, 0x59, 0x35, 0x6e, 0x13, 0x75, 0x4f, 0x8a, 0x5f, 0xb2, 0x83, 0x14,
	0x6b, 0xb1, 0xc1, 0x62, 0x1e, 0x17, 0xa6, 0x89, 0x47, 0xb1, 0x41, 0x3e, 0x63, 0xa7, 0x64, 0xb7,
	0x0b, 0xe8, 0x50, 0x22, 0x6d, 0x51, 0x81, 0x15, 0xf2, 0x0d, 0xbd, 0x2b, 0x16, 0xf1, 0x27, 0x4e,
	0x42, 0xd8, 0xa6, 0xec, 0xa9, 0x8f, 0xf8, 0x15, 0x1b, 0xc7, 0x19, 0x17, 0xdc, 0xa1, 0x7f, 0x1d,
	0xfb, 0xc7, 0x21, 0x58, 0xa1, 0xf6, 0x43, 0xb7, 0xf7, 0x9b, 0xdf, 0x7e, 0xb3, 0xf3, 0x9b, 0xbf,
	0xfd, 0xe6, 0xa7, 0x7f, 0xb3, 0xf3, 0x9b, 0x6f, 0xfe, 0xfa, 0x5f, 0xbc, 0xbd, 0xf9, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x50, 0xd3, 0x13, 0x96, 0xde, 0x01, 0x00, 0x00,
}