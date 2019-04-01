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
// source: snmp_trap_stats.proto

package cisco_ios_xr_snmp_agent_oper_snmp_trap_servers_trap_server

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

type SnmpTrapStats_KEYS struct {
	TrapHost             string   `protobuf:"bytes,1,opt,name=trap_host,json=trapHost,proto3" json:"trap_host,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpTrapStats_KEYS) Reset()         { *m = SnmpTrapStats_KEYS{} }
func (m *SnmpTrapStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapStats_KEYS) ProtoMessage()    {}
func (*SnmpTrapStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bbd136ce10fbc16, []int{0}
}

func (m *SnmpTrapStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapStats_KEYS.Unmarshal(m, b)
}
func (m *SnmpTrapStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapStats_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpTrapStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapStats_KEYS.Merge(m, src)
}
func (m *SnmpTrapStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapStats_KEYS.Size(m)
}
func (m *SnmpTrapStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapStats_KEYS proto.InternalMessageInfo

func (m *SnmpTrapStats_KEYS) GetTrapHost() string {
	if m != nil {
		return m.TrapHost
	}
	return ""
}

func (m *SnmpTrapStats_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type SnmpTrapStats struct {
	NumberOfPktsInTrapQ  uint32   `protobuf:"varint,50,opt,name=number_of_pkts_in_trap_q,json=numberOfPktsInTrapQ,proto3" json:"number_of_pkts_in_trap_q,omitempty"`
	MaxQLengthOfTrapQ    uint32   `protobuf:"varint,51,opt,name=max_q_length_of_trap_q,json=maxQLengthOfTrapQ,proto3" json:"max_q_length_of_trap_q,omitempty"`
	NumberOfPktsSent     uint32   `protobuf:"varint,52,opt,name=number_of_pkts_sent,json=numberOfPktsSent,proto3" json:"number_of_pkts_sent,omitempty"`
	NumberOfPktsDropped  uint32   `protobuf:"varint,53,opt,name=number_of_pkts_dropped,json=numberOfPktsDropped,proto3" json:"number_of_pkts_dropped,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpTrapStats) Reset()         { *m = SnmpTrapStats{} }
func (m *SnmpTrapStats) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapStats) ProtoMessage()    {}
func (*SnmpTrapStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bbd136ce10fbc16, []int{1}
}

func (m *SnmpTrapStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapStats.Unmarshal(m, b)
}
func (m *SnmpTrapStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapStats.Marshal(b, m, deterministic)
}
func (m *SnmpTrapStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapStats.Merge(m, src)
}
func (m *SnmpTrapStats) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapStats.Size(m)
}
func (m *SnmpTrapStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapStats.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapStats proto.InternalMessageInfo

func (m *SnmpTrapStats) GetNumberOfPktsInTrapQ() uint32 {
	if m != nil {
		return m.NumberOfPktsInTrapQ
	}
	return 0
}

func (m *SnmpTrapStats) GetMaxQLengthOfTrapQ() uint32 {
	if m != nil {
		return m.MaxQLengthOfTrapQ
	}
	return 0
}

func (m *SnmpTrapStats) GetNumberOfPktsSent() uint32 {
	if m != nil {
		return m.NumberOfPktsSent
	}
	return 0
}

func (m *SnmpTrapStats) GetNumberOfPktsDropped() uint32 {
	if m != nil {
		return m.NumberOfPktsDropped
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpTrapStats_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.trap_servers.trap_server.snmp_trap_stats_KEYS")
	proto.RegisterType((*SnmpTrapStats)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.trap_servers.trap_server.snmp_trap_stats")
}

func init() { proto.RegisterFile("snmp_trap_stats.proto", fileDescriptor_7bbd136ce10fbc16) }

var fileDescriptor_7bbd136ce10fbc16 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0x88, 0xd8, 0x05, 0x51, 0xb7, 0x5a, 0x16, 0xbc, 0x94, 0x9e, 0x7a, 0x31, 0xa0,
	0xb1, 0x17, 0xcf, 0x8a, 0x8a, 0x42, 0x6d, 0xeb, 0xc5, 0xd3, 0xb0, 0x6d, 0x37, 0x6d, 0xa8, 0xd9,
	0xd9, 0xce, 0x8c, 0x92, 0xb7, 0xf5, 0x55, 0x24, 0x1b, 0x0f, 0x25, 0xbd, 0xed, 0xcc, 0xf7, 0x7f,
	0x33, 0xcb, 0xa8, 0x4b, 0xf6, 0x65, 0x00, 0x21, 0x1b, 0x80, 0xc5, 0x0a, 0xa7, 0x81, 0x50, 0x50,
	0xdf, 0x2f, 0x0a, 0x5e, 0x20, 0x14, 0xc8, 0x50, 0x11, 0xc4, 0x8c, 0x5d, 0x39, 0x2f, 0x80, 0xc1,
	0x51, 0x5a, 0xd7, 0x69, 0xe3, 0x38, 0xfa, 0x71, 0xc4, 0xbb, 0xc5, 0xe0, 0x49, 0x5d, 0xb4, 0x86,
	0xc2, 0xeb, 0xe3, 0xe7, 0x4c, 0x5f, 0xa9, 0x4e, 0x6c, 0xad, 0x91, 0xc5, 0x24, 0xfd, 0x64, 0xd8,
	0x99, 0x1e, 0xd7, 0x8d, 0x67, 0x64, 0xd1, 0x5a, 0x1d, 0x06, 0x24, 0x31, 0x07, 0xfd, 0x64, 0x78,
	0x32, 0x8d, 0xef, 0xc1, 0x6f, 0xa2, 0x4e, 0x5b, 0x93, 0xf4, 0x48, 0x19, 0xff, 0x5d, 0xce, 0x1d,
	0x01, 0xe6, 0x10, 0x36, 0xc2, 0x50, 0xf8, 0x86, 0x6f, 0xcd, 0x6d, 0x74, 0xbb, 0x0d, 0x1f, 0xe7,
	0xef, 0x1b, 0xe1, 0x17, 0xff, 0x41, 0x36, 0x4c, 0xf4, 0x8d, 0xea, 0x95, 0xb6, 0x82, 0x2d, 0x7c,
	0x39, 0xbf, 0x92, 0x75, 0x2d, 0xff, 0x4b, 0x59, 0x94, 0xce, 0x4b, 0x5b, 0x4d, 0xde, 0x22, 0x1b,
	0xe7, 0x8d, 0x72, 0xad, 0xba, 0xad, 0x4d, 0xec, 0xbc, 0x98, 0xbb, 0x98, 0x3f, 0xdb, 0x5d, 0x32,
	0x73, 0x5e, 0x74, 0xa6, 0x7a, 0xad, 0xf8, 0x92, 0x30, 0x04, 0xb7, 0x34, 0xa3, 0xfd, 0x6f, 0x3d,
	0x34, 0x68, 0x7e, 0x14, 0xaf, 0x9d, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x37, 0x34, 0xe9,
	0x86, 0x01, 0x00, 0x00,
}
