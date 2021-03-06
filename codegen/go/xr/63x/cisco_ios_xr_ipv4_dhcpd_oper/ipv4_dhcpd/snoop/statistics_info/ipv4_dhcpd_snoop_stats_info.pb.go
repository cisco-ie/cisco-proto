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
// source: ipv4_dhcpd_snoop_stats_info.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_snoop_statistics_info

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

type Ipv4DhcpdSnoopStatsInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) Reset()         { *m = Ipv4DhcpdSnoopStatsInfo_KEYS{} }
func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdSnoopStatsInfo_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdSnoopStatsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_60738a6073eacb82, []int{0}
}

func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS.Size(m)
}
func (m *Ipv4DhcpdSnoopStatsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo_KEYS proto.InternalMessageInfo

type Ipv4DhcpdSnoopStatsInfo struct {
	SnoopStatsTimestamp  uint32   `protobuf:"varint,50,opt,name=snoop_stats_timestamp,json=snoopStatsTimestamp,proto3" json:"snoop_stats_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdSnoopStatsInfo) Reset()         { *m = Ipv4DhcpdSnoopStatsInfo{} }
func (m *Ipv4DhcpdSnoopStatsInfo) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdSnoopStatsInfo) ProtoMessage()    {}
func (*Ipv4DhcpdSnoopStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_60738a6073eacb82, []int{1}
}

func (m *Ipv4DhcpdSnoopStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo.Unmarshal(m, b)
}
func (m *Ipv4DhcpdSnoopStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdSnoopStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo.Merge(m, src)
}
func (m *Ipv4DhcpdSnoopStatsInfo) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo.Size(m)
}
func (m *Ipv4DhcpdSnoopStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdSnoopStatsInfo proto.InternalMessageInfo

func (m *Ipv4DhcpdSnoopStatsInfo) GetSnoopStatsTimestamp() uint32 {
	if m != nil {
		return m.SnoopStatsTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4DhcpdSnoopStatsInfo_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.snoop.statistics_info.ipv4_dhcpd_snoop_stats_info_KEYS")
	proto.RegisterType((*Ipv4DhcpdSnoopStatsInfo)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.snoop.statistics_info.ipv4_dhcpd_snoop_stats_info")
}

func init() { proto.RegisterFile("ipv4_dhcpd_snoop_stats_info.proto", fileDescriptor_60738a6073eacb82) }

var fileDescriptor_60738a6073eacb82 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x2c, 0x28, 0x33,
	0x89, 0x4f, 0xc9, 0x48, 0x2e, 0x48, 0x89, 0x2f, 0xce, 0xcb, 0xcf, 0x2f, 0x88, 0x2f, 0x2e, 0x49,
	0x2c, 0x29, 0x8e, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4d,
	0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x47, 0x52, 0x9f, 0x5f,
	0x90, 0x5a, 0xa4, 0x87, 0xe0, 0xeb, 0x81, 0xf5, 0xeb, 0x81, 0xf4, 0x67, 0x16, 0x97, 0x64, 0x26,
	0x43, 0x0c, 0x51, 0x52, 0xe2, 0x52, 0xc0, 0x63, 0x47, 0xbc, 0xb7, 0x6b, 0x64, 0xb0, 0x52, 0x20,
	0x97, 0x34, 0x1e, 0x35, 0x42, 0x46, 0x5c, 0xa2, 0xc8, 0x62, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25,
	0x89, 0xb9, 0x05, 0x12, 0x46, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0xc2, 0x60, 0xc9, 0x60, 0x90, 0x5c,
	0x08, 0x4c, 0x2a, 0x89, 0x0d, 0xec, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0xba,
	0xdc, 0x8a, 0xe1, 0x00, 0x00, 0x00,
}
