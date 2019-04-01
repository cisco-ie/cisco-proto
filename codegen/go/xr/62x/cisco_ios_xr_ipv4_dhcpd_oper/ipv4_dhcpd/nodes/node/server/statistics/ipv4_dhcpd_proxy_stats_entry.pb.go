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
// source: ipv4_dhcpd_proxy_stats_entry.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_server_statistics

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

type Ipv4DhcpdProxyStatsEntry_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyStatsEntry_KEYS) Reset()         { *m = Ipv4DhcpdProxyStatsEntry_KEYS{} }
func (m *Ipv4DhcpdProxyStatsEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyStatsEntry_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdProxyStatsEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_499d3f02475931b4, []int{0}
}

func (m *Ipv4DhcpdProxyStatsEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyStatsEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyStatsEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdProxyStatsEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS.Size(m)
}
func (m *Ipv4DhcpdProxyStatsEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyStatsEntry_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyStatsEntry_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

type Ipv4DhcpdProxyFilteredStats struct {
	ReceivedPackets      uint64   `protobuf:"varint,1,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	TransmittedPackets   uint64   `protobuf:"varint,2,opt,name=transmitted_packets,json=transmittedPackets,proto3" json:"transmitted_packets,omitempty"`
	DroppedPackets       uint64   `protobuf:"varint,3,opt,name=dropped_packets,json=droppedPackets,proto3" json:"dropped_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyFilteredStats) Reset()         { *m = Ipv4DhcpdProxyFilteredStats{} }
func (m *Ipv4DhcpdProxyFilteredStats) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyFilteredStats) ProtoMessage()    {}
func (*Ipv4DhcpdProxyFilteredStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_499d3f02475931b4, []int{1}
}

func (m *Ipv4DhcpdProxyFilteredStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyFilteredStats.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyFilteredStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyFilteredStats.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyFilteredStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyFilteredStats.Merge(m, src)
}
func (m *Ipv4DhcpdProxyFilteredStats) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyFilteredStats.Size(m)
}
func (m *Ipv4DhcpdProxyFilteredStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyFilteredStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyFilteredStats proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyFilteredStats) GetReceivedPackets() uint64 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv4DhcpdProxyFilteredStats) GetTransmittedPackets() uint64 {
	if m != nil {
		return m.TransmittedPackets
	}
	return 0
}

func (m *Ipv4DhcpdProxyFilteredStats) GetDroppedPackets() uint64 {
	if m != nil {
		return m.DroppedPackets
	}
	return 0
}

type Ipv4DhcpdProxyStatsItem struct {
	VrfName              string                       `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Statistics           *Ipv4DhcpdProxyFilteredStats `protobuf:"bytes,2,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Ipv4DhcpdProxyStatsItem) Reset()         { *m = Ipv4DhcpdProxyStatsItem{} }
func (m *Ipv4DhcpdProxyStatsItem) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyStatsItem) ProtoMessage()    {}
func (*Ipv4DhcpdProxyStatsItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_499d3f02475931b4, []int{2}
}

func (m *Ipv4DhcpdProxyStatsItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsItem.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyStatsItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsItem.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyStatsItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyStatsItem.Merge(m, src)
}
func (m *Ipv4DhcpdProxyStatsItem) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsItem.Size(m)
}
func (m *Ipv4DhcpdProxyStatsItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyStatsItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyStatsItem proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyStatsItem) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4DhcpdProxyStatsItem) GetStatistics() *Ipv4DhcpdProxyFilteredStats {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type Ipv4DhcpdProxyStatsEntry struct {
	Ipv4DhcpdProxyStats  []*Ipv4DhcpdProxyStatsItem `protobuf:"bytes,50,rep,name=ipv4_dhcpd_proxy_stats,json=ipv4DhcpdProxyStats,proto3" json:"ipv4_dhcpd_proxy_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Ipv4DhcpdProxyStatsEntry) Reset()         { *m = Ipv4DhcpdProxyStatsEntry{} }
func (m *Ipv4DhcpdProxyStatsEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyStatsEntry) ProtoMessage()    {}
func (*Ipv4DhcpdProxyStatsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_499d3f02475931b4, []int{3}
}

func (m *Ipv4DhcpdProxyStatsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyStatsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyStatsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyStatsEntry.Merge(m, src)
}
func (m *Ipv4DhcpdProxyStatsEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyStatsEntry.Size(m)
}
func (m *Ipv4DhcpdProxyStatsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyStatsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyStatsEntry proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyStatsEntry) GetIpv4DhcpdProxyStats() []*Ipv4DhcpdProxyStatsItem {
	if m != nil {
		return m.Ipv4DhcpdProxyStats
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4DhcpdProxyStatsEntry_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.statistics.ipv4_dhcpd_proxy_stats_entry_KEYS")
	proto.RegisterType((*Ipv4DhcpdProxyFilteredStats)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.statistics.ipv4_dhcpd_proxy_filtered_stats")
	proto.RegisterType((*Ipv4DhcpdProxyStatsItem)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.statistics.ipv4_dhcpd_proxy_stats_item")
	proto.RegisterType((*Ipv4DhcpdProxyStatsEntry)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.statistics.ipv4_dhcpd_proxy_stats_entry")
}

func init() { proto.RegisterFile("ipv4_dhcpd_proxy_stats_entry.proto", fileDescriptor_499d3f02475931b4) }

var fileDescriptor_499d3f02475931b4 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x66, 0xad, 0x54, 0x9d, 0x82, 0x95, 0x2d, 0x94, 0x8a, 0x82, 0x35, 0x17, 0xeb, 0x65, 0x85,
	0xea, 0xcd, 0x6b, 0x3d, 0x09, 0x52, 0xd2, 0x93, 0xa7, 0x25, 0x66, 0xa7, 0xb8, 0x68, 0xb2, 0xcb,
	0xec, 0x12, 0xda, 0x07, 0xf0, 0x15, 0x7c, 0x04, 0x5f, 0xc1, 0xab, 0x8f, 0x26, 0xd9, 0xa6, 0x24,
	0xa0, 0xd6, 0x8b, 0x5e, 0x02, 0x33, 0xdf, 0xcf, 0xcc, 0x7c, 0x59, 0x88, 0xb4, 0x2d, 0xae, 0xa4,
	0x7a, 0x4c, 0xad, 0x92, 0x96, 0xcc, 0x62, 0x29, 0x9d, 0x4f, 0xbc, 0x93, 0x98, 0x7b, 0x5a, 0x0a,
	0x4b, 0xc6, 0x1b, 0x3e, 0x49, 0xb5, 0x4b, 0x8d, 0xd4, 0xc6, 0xc9, 0x05, 0xc9, 0x86, 0xc0, 0x58,
	0x24, 0x51, 0xd7, 0x22, 0x37, 0x0a, 0x5d, 0xf8, 0x0a, 0x87, 0x54, 0x20, 0x89, 0xd2, 0x4c, 0x3b,
	0xaf, 0x53, 0x17, 0x5d, 0xc3, 0xe9, 0xa6, 0x59, 0xf2, 0xf6, 0xe6, 0x7e, 0xc6, 0xfb, 0xd0, 0x2e,
	0xe5, 0x5a, 0x0d, 0xd8, 0x90, 0x8d, 0xf6, 0xe2, 0xaa, 0x8a, 0xde, 0x18, 0x9c, 0x7c, 0x51, 0xcf,
	0xf5, 0xb3, 0x47, 0x42, 0xb5, 0xb2, 0xe1, 0xe7, 0x70, 0x40, 0x98, 0xa2, 0x2e, 0x50, 0x49, 0x9b,
	0xa4, 0x4f, 0xe8, 0x5d, 0x70, 0xd9, 0x8e, 0xbb, 0xeb, 0xfe, 0x74, 0xd5, 0xe6, 0x17, 0xd0, 0xf3,
	0x94, 0xe4, 0x2e, 0xd3, 0xde, 0x37, 0xd8, 0x5b, 0x81, 0xcd, 0x1b, 0xd0, 0x5a, 0x70, 0x06, 0x5d,
	0x45, 0xc6, 0xda, 0x06, 0xb9, 0x15, 0xc8, 0xfb, 0x55, 0xbb, 0x22, 0x46, 0x1f, 0x0c, 0x8e, 0x7e,
	0x38, 0x53, 0x7b, 0xcc, 0xf8, 0x21, 0xec, 0x16, 0x34, 0x97, 0x79, 0x92, 0x61, 0x75, 0xe2, 0x4e,
	0x41, 0xf3, 0xbb, 0x24, 0x43, 0xfe, 0xc2, 0x00, 0xea, 0xbc, 0xc2, 0x32, 0x9d, 0x31, 0x8a, 0xbf,
	0x08, 0x5f, 0xfc, 0x92, 0x5d, 0xdc, 0x18, 0x1c, 0xbd, 0x33, 0x38, 0xde, 0xf4, 0xa7, 0xf8, 0x2b,
	0x83, 0xfe, 0xf7, 0x84, 0xc1, 0x78, 0xd8, 0x1a, 0x75, 0xc6, 0xc9, 0x3f, 0x2d, 0x5d, 0xe7, 0x18,
	0xf7, 0x4a, 0x70, 0x52, 0x62, 0xd3, 0x12, 0x9a, 0x95, 0xc8, 0x43, 0x3b, 0xbc, 0xd7, 0xcb, 0xcf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xbf, 0x9c, 0x03, 0xd5, 0x02, 0x00, 0x00,
}
