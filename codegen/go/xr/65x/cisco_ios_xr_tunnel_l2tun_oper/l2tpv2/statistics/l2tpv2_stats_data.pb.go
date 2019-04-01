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
// source: l2tpv2_stats_data.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tpv2_statistics

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

type L2Tpv2StatsData_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2Tpv2StatsData_KEYS) Reset()         { *m = L2Tpv2StatsData_KEYS{} }
func (m *L2Tpv2StatsData_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2Tpv2StatsData_KEYS) ProtoMessage()    {}
func (*L2Tpv2StatsData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_803612e42cbd3326, []int{0}
}

func (m *L2Tpv2StatsData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2Tpv2StatsData_KEYS.Unmarshal(m, b)
}
func (m *L2Tpv2StatsData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2Tpv2StatsData_KEYS.Marshal(b, m, deterministic)
}
func (m *L2Tpv2StatsData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2Tpv2StatsData_KEYS.Merge(m, src)
}
func (m *L2Tpv2StatsData_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2Tpv2StatsData_KEYS.Size(m)
}
func (m *L2Tpv2StatsData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2Tpv2StatsData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2Tpv2StatsData_KEYS proto.InternalMessageInfo

type L2Tpv2StatsData struct {
	Tunnels                     uint32   `protobuf:"varint,50,opt,name=tunnels,proto3" json:"tunnels,omitempty"`
	Sessions                    uint32   `protobuf:"varint,51,opt,name=sessions,proto3" json:"sessions,omitempty"`
	SentPackets                 uint32   `protobuf:"varint,52,opt,name=sent_packets,json=sentPackets,proto3" json:"sent_packets,omitempty"`
	ReceivedPackets             uint32   `protobuf:"varint,53,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	AveragePacketProcessingTime uint32   `protobuf:"varint,54,opt,name=average_packet_processing_time,json=averagePacketProcessingTime,proto3" json:"average_packet_processing_time,omitempty"`
	ReceivedOutOfOrderPackets   uint32   `protobuf:"varint,55,opt,name=received_out_of_order_packets,json=receivedOutOfOrderPackets,proto3" json:"received_out_of_order_packets,omitempty"`
	ReorderPackets              uint32   `protobuf:"varint,56,opt,name=reorder_packets,json=reorderPackets,proto3" json:"reorder_packets,omitempty"`
	ReorderDeviationPackets     uint32   `protobuf:"varint,57,opt,name=reorder_deviation_packets,json=reorderDeviationPackets,proto3" json:"reorder_deviation_packets,omitempty"`
	IncomingDroppedPackets      uint32   `protobuf:"varint,58,opt,name=incoming_dropped_packets,json=incomingDroppedPackets,proto3" json:"incoming_dropped_packets,omitempty"`
	BufferedPackets             uint32   `protobuf:"varint,59,opt,name=buffered_packets,json=bufferedPackets,proto3" json:"buffered_packets,omitempty"`
	NetioPackets                uint32   `protobuf:"varint,60,opt,name=netio_packets,json=netioPackets,proto3" json:"netio_packets,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *L2Tpv2StatsData) Reset()         { *m = L2Tpv2StatsData{} }
func (m *L2Tpv2StatsData) String() string { return proto.CompactTextString(m) }
func (*L2Tpv2StatsData) ProtoMessage()    {}
func (*L2Tpv2StatsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_803612e42cbd3326, []int{1}
}

func (m *L2Tpv2StatsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2Tpv2StatsData.Unmarshal(m, b)
}
func (m *L2Tpv2StatsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2Tpv2StatsData.Marshal(b, m, deterministic)
}
func (m *L2Tpv2StatsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2Tpv2StatsData.Merge(m, src)
}
func (m *L2Tpv2StatsData) XXX_Size() int {
	return xxx_messageInfo_L2Tpv2StatsData.Size(m)
}
func (m *L2Tpv2StatsData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2Tpv2StatsData.DiscardUnknown(m)
}

var xxx_messageInfo_L2Tpv2StatsData proto.InternalMessageInfo

func (m *L2Tpv2StatsData) GetTunnels() uint32 {
	if m != nil {
		return m.Tunnels
	}
	return 0
}

func (m *L2Tpv2StatsData) GetSessions() uint32 {
	if m != nil {
		return m.Sessions
	}
	return 0
}

func (m *L2Tpv2StatsData) GetSentPackets() uint32 {
	if m != nil {
		return m.SentPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetAveragePacketProcessingTime() uint32 {
	if m != nil {
		return m.AveragePacketProcessingTime
	}
	return 0
}

func (m *L2Tpv2StatsData) GetReceivedOutOfOrderPackets() uint32 {
	if m != nil {
		return m.ReceivedOutOfOrderPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetReorderPackets() uint32 {
	if m != nil {
		return m.ReorderPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetReorderDeviationPackets() uint32 {
	if m != nil {
		return m.ReorderDeviationPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetIncomingDroppedPackets() uint32 {
	if m != nil {
		return m.IncomingDroppedPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetBufferedPackets() uint32 {
	if m != nil {
		return m.BufferedPackets
	}
	return 0
}

func (m *L2Tpv2StatsData) GetNetioPackets() uint32 {
	if m != nil {
		return m.NetioPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*L2Tpv2StatsData_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.statistics.l2tpv2_stats_data_KEYS")
	proto.RegisterType((*L2Tpv2StatsData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.statistics.l2tpv2_stats_data")
}

func init() { proto.RegisterFile("l2tpv2_stats_data.proto", fileDescriptor_803612e42cbd3326) }

var fileDescriptor_803612e42cbd3326 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x86, 0x63, 0x62, 0xd4, 0xac, 0x20, 0xda, 0x03, 0x14, 0x8d, 0x46, 0xf1, 0xa0, 0x5e, 0x88,
	0x29, 0x7e, 0x20, 0x7a, 0x30, 0x11, 0x4f, 0x1e, 0x20, 0xea, 0xc5, 0xd3, 0xa6, 0xb4, 0x53, 0xb2,
	0x11, 0x76, 0x36, 0xbb, 0xd3, 0xc6, 0x3f, 0xe1, 0x7f, 0x36, 0xdd, 0x76, 0x37, 0x18, 0x8e, 0xfb,
	0xbe, 0xcf, 0x33, 0x03, 0x93, 0xb2, 0xce, 0x22, 0x22, 0x55, 0x44, 0xdc, 0x50, 0x4c, 0x86, 0xa7,
	0x31, 0xc5, 0x7d, 0xa5, 0x91, 0x30, 0xb8, 0x4e, 0x84, 0x49, 0x90, 0x0b, 0x34, 0xfc, 0x47, 0x73,
	0xca, 0xa5, 0x84, 0x05, 0x5f, 0x44, 0x94, 0x4b, 0x8e, 0x0a, 0x74, 0xbf, 0xf2, 0xfa, 0xa5, 0x27,
	0x0c, 0x89, 0xc4, 0xf4, 0x42, 0xd6, 0x5e, 0x1b, 0xc6, 0xdf, 0x5e, 0xbf, 0x3e, 0x7a, 0xbf, 0x9b,
	0xec, 0x60, 0xad, 0x0a, 0x42, 0xb6, 0x5d, 0x8d, 0x35, 0x61, 0x74, 0xba, 0x71, 0xd9, 0x7c, 0x77,
	0xcf, 0xe0, 0x90, 0xed, 0x18, 0x30, 0x46, 0xa0, 0x34, 0xe1, 0xc0, 0x56, 0xfe, 0x1d, 0x9c, 0xb1,
	0x86, 0x01, 0x49, 0x5c, 0xc5, 0xc9, 0x37, 0x90, 0x09, 0x6f, 0x6c, 0xbf, 0x5b, 0x66, 0xd3, 0x2a,
	0x0a, 0xae, 0xd8, 0xbe, 0x86, 0x04, 0x44, 0x01, 0xa9, 0xc7, 0x6e, 0x2d, 0xd6, 0x72, 0xb9, 0x43,
	0x5f, 0xd8, 0x49, 0x5c, 0x80, 0x8e, 0xe7, 0x50, 0x93, 0x5c, 0x69, 0x4c, 0xca, 0x5d, 0x72, 0xce,
	0x49, 0x2c, 0x21, 0xbc, 0xb3, 0xe2, 0x51, 0x4d, 0x55, 0xde, 0xd4, 0x33, 0x9f, 0x62, 0x09, 0xc1,
	0x33, 0x3b, 0xf6, 0xfb, 0x30, 0x27, 0x8e, 0x19, 0x47, 0x9d, 0x82, 0xf6, 0xcb, 0xef, 0xed, 0x8c,
	0xae, 0x83, 0x26, 0x39, 0x4d, 0xb2, 0x49, 0x49, 0xb8, 0x9f, 0x71, 0xc1, 0x5a, 0x1a, 0xfe, 0x3b,
	0x43, 0xeb, 0xec, 0xd5, 0xb1, 0x03, 0x47, 0xac, 0xeb, 0xc0, 0x14, 0x0a, 0x11, 0x93, 0x40, 0xe9,
	0x95, 0x07, 0xab, 0x74, 0x6a, 0x60, 0xec, 0x7a, 0xe7, 0x0e, 0x59, 0x28, 0x64, 0x82, 0xcb, 0xf2,
	0xaf, 0xa5, 0x1a, 0x95, 0x5a, 0x39, 0xcf, 0xc8, 0xaa, 0x6d, 0xd7, 0x8f, 0xab, 0x7a, 0xe5, 0xa0,
	0xb3, 0x3c, 0xcb, 0x40, 0xaf, 0x18, 0x8f, 0xd5, 0x41, 0x5d, 0xee, 0xd0, 0x73, 0xd6, 0x94, 0x40,
	0x02, 0x3d, 0xf7, 0x64, 0xb9, 0x86, 0x0d, 0x6b, 0x68, 0xb6, 0x65, 0x3f, 0xb1, 0xc1, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x99, 0x83, 0x80, 0x12, 0x7d, 0x02, 0x00, 0x00,
}
