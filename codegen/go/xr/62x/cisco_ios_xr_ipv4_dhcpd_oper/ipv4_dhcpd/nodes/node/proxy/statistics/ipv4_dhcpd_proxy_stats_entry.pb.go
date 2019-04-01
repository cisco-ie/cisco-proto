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

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_proxy_statistics

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
	proto.RegisterType((*Ipv4DhcpdProxyStatsEntry_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.statistics.ipv4_dhcpd_proxy_stats_entry_KEYS")
	proto.RegisterType((*Ipv4DhcpdProxyFilteredStats)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.statistics.ipv4_dhcpd_proxy_filtered_stats")
	proto.RegisterType((*Ipv4DhcpdProxyStatsItem)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.statistics.ipv4_dhcpd_proxy_stats_item")
	proto.RegisterType((*Ipv4DhcpdProxyStatsEntry)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.statistics.ipv4_dhcpd_proxy_stats_entry")
}

func init() { proto.RegisterFile("ipv4_dhcpd_proxy_stats_entry.proto", fileDescriptor_499d3f02475931b4) }

var fileDescriptor_499d3f02475931b4 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcf, 0x4a, 0x33, 0x31,
	0x10, 0x27, 0x5f, 0x3f, 0xaa, 0x4e, 0xc1, 0x4a, 0x0a, 0xa5, 0xa2, 0x60, 0xdd, 0x8b, 0xf5, 0x12,
	0xa1, 0x7a, 0xf3, 0xa8, 0x9e, 0x04, 0x29, 0xdb, 0x93, 0xa7, 0xb8, 0x26, 0x53, 0x0c, 0xba, 0x9b,
	0x90, 0x84, 0xa5, 0xbd, 0xfb, 0x0a, 0xbe, 0x82, 0x8f, 0xe0, 0xd5, 0x57, 0x93, 0xa4, 0x5b, 0x76,
	0x41, 0xad, 0x17, 0xbd, 0x2c, 0xcc, 0xfc, 0xfe, 0xcc, 0xcc, 0x6f, 0x03, 0x89, 0x32, 0xe5, 0x19,
	0x97, 0x0f, 0xc2, 0x48, 0x6e, 0xac, 0x9e, 0x2f, 0xb8, 0xf3, 0x99, 0x77, 0x1c, 0x0b, 0x6f, 0x17,
	0xcc, 0x58, 0xed, 0x35, 0xbd, 0x10, 0xca, 0x09, 0xcd, 0x95, 0x76, 0x7c, 0x6e, 0x79, 0x43, 0xa0,
	0x0d, 0x5a, 0x56, 0xd7, 0xac, 0xd0, 0x12, 0x5d, 0xfc, 0xb2, 0xe8, 0xc5, 0x82, 0x97, 0x72, 0x5e,
	0x09, 0x97, 0x9c, 0xc3, 0xe1, 0xba, 0x51, 0xfc, 0xfa, 0xea, 0x76, 0x4a, 0xfb, 0xd0, 0x0e, 0x6a,
	0x25, 0x07, 0x64, 0x48, 0x46, 0x5b, 0x69, 0x55, 0x25, 0xaf, 0x04, 0x0e, 0x3e, 0xa9, 0x67, 0xea,
	0xc9, 0xa3, 0x45, 0xb9, 0xb4, 0xa1, 0xc7, 0xb0, 0x63, 0x51, 0xa0, 0x2a, 0x51, 0x72, 0x93, 0x89,
	0x47, 0xf4, 0x2e, 0xba, 0xfc, 0x4f, 0xbb, 0xab, 0xfe, 0x64, 0xd9, 0xa6, 0x27, 0xd0, 0xf3, 0x36,
	0x2b, 0x5c, 0xae, 0xbc, 0x6f, 0xb0, 0xff, 0x45, 0x36, 0x6d, 0x40, 0x2b, 0xc1, 0x11, 0x74, 0xa5,
	0xd5, 0xc6, 0x34, 0xc8, 0xad, 0x48, 0xde, 0xae, 0xda, 0x15, 0x31, 0x79, 0x27, 0xb0, 0xf7, 0xcd,
	0x99, 0xca, 0x63, 0x4e, 0x77, 0x61, 0xb3, 0xb4, 0x33, 0x5e, 0x64, 0x39, 0x56, 0x27, 0x6e, 0x94,
	0x76, 0x76, 0x93, 0xe5, 0x48, 0x9f, 0x09, 0x40, 0x9d, 0x57, 0x5c, 0xa6, 0x33, 0x96, 0xec, 0x17,
	0xb2, 0x67, 0x3f, 0x44, 0x97, 0x36, 0xe6, 0x26, 0x6f, 0x04, 0xf6, 0xd7, 0xfd, 0x28, 0xfa, 0x42,
	0xa0, 0xff, 0x35, 0x61, 0x30, 0x1e, 0xb6, 0x46, 0x9d, 0xf1, 0xdd, 0xdf, 0xec, 0x5c, 0xa7, 0x98,
	0xf6, 0x02, 0x78, 0x19, 0xb0, 0x49, 0x80, 0xa6, 0x01, 0xb9, 0x6f, 0xc7, 0xc7, 0x7a, 0xfa, 0x11,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xae, 0x37, 0x98, 0xd2, 0x02, 0x00, 0x00,
}
