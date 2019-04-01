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
// source: ospf_sh_ncd_stats.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_statistics_nsr_pl_stats

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

type OspfShNcdStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNcdStats_KEYS) Reset()         { *m = OspfShNcdStats_KEYS{} }
func (m *OspfShNcdStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNcdStats_KEYS) ProtoMessage()    {}
func (*OspfShNcdStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69aa6118418380e, []int{0}
}

func (m *OspfShNcdStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNcdStats_KEYS.Unmarshal(m, b)
}
func (m *OspfShNcdStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNcdStats_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShNcdStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNcdStats_KEYS.Merge(m, src)
}
func (m *OspfShNcdStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNcdStats_KEYS.Size(m)
}
func (m *OspfShNcdStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNcdStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNcdStats_KEYS proto.InternalMessageInfo

func (m *OspfShNcdStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type NsrPlClientStatsType struct {
	NumSent              []uint64 `protobuf:"varint,1,rep,packed,name=num_sent,json=numSent,proto3" json:"num_sent,omitempty"`
	NumRecv              []uint64 `protobuf:"varint,2,rep,packed,name=num_recv,json=numRecv,proto3" json:"num_recv,omitempty"`
	NumSentDrop          []string `protobuf:"bytes,3,rep,name=num_sent_drop,json=numSentDrop,proto3" json:"num_sent_drop,omitempty"`
	NumRecvDrop          []string `protobuf:"bytes,4,rep,name=num_recv_drop,json=numRecvDrop,proto3" json:"num_recv_drop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NsrPlClientStatsType) Reset()         { *m = NsrPlClientStatsType{} }
func (m *NsrPlClientStatsType) String() string { return proto.CompactTextString(m) }
func (*NsrPlClientStatsType) ProtoMessage()    {}
func (*NsrPlClientStatsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69aa6118418380e, []int{1}
}

func (m *NsrPlClientStatsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NsrPlClientStatsType.Unmarshal(m, b)
}
func (m *NsrPlClientStatsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NsrPlClientStatsType.Marshal(b, m, deterministic)
}
func (m *NsrPlClientStatsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NsrPlClientStatsType.Merge(m, src)
}
func (m *NsrPlClientStatsType) XXX_Size() int {
	return xxx_messageInfo_NsrPlClientStatsType.Size(m)
}
func (m *NsrPlClientStatsType) XXX_DiscardUnknown() {
	xxx_messageInfo_NsrPlClientStatsType.DiscardUnknown(m)
}

var xxx_messageInfo_NsrPlClientStatsType proto.InternalMessageInfo

func (m *NsrPlClientStatsType) GetNumSent() []uint64 {
	if m != nil {
		return m.NumSent
	}
	return nil
}

func (m *NsrPlClientStatsType) GetNumRecv() []uint64 {
	if m != nil {
		return m.NumRecv
	}
	return nil
}

func (m *NsrPlClientStatsType) GetNumSentDrop() []string {
	if m != nil {
		return m.NumSentDrop
	}
	return nil
}

func (m *NsrPlClientStatsType) GetNumRecvDrop() []string {
	if m != nil {
		return m.NumRecvDrop
	}
	return nil
}

type OspfShNcdStats struct {
	NcdPri               []*NsrPlClientStatsType `protobuf:"bytes,50,rep,name=ncd_pri,json=ncdPri,proto3" json:"ncd_pri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OspfShNcdStats) Reset()         { *m = OspfShNcdStats{} }
func (m *OspfShNcdStats) String() string { return proto.CompactTextString(m) }
func (*OspfShNcdStats) ProtoMessage()    {}
func (*OspfShNcdStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c69aa6118418380e, []int{2}
}

func (m *OspfShNcdStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNcdStats.Unmarshal(m, b)
}
func (m *OspfShNcdStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNcdStats.Marshal(b, m, deterministic)
}
func (m *OspfShNcdStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNcdStats.Merge(m, src)
}
func (m *OspfShNcdStats) XXX_Size() int {
	return xxx_messageInfo_OspfShNcdStats.Size(m)
}
func (m *OspfShNcdStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNcdStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNcdStats proto.InternalMessageInfo

func (m *OspfShNcdStats) GetNcdPri() []*NsrPlClientStatsType {
	if m != nil {
		return m.NcdPri
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShNcdStats_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.statistics.nsr_pl_stats.ospf_sh_ncd_stats_KEYS")
	proto.RegisterType((*NsrPlClientStatsType)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.statistics.nsr_pl_stats.nsr_pl_client_stats_type")
	proto.RegisterType((*OspfShNcdStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.statistics.nsr_pl_stats.ospf_sh_ncd_stats")
}

func init() { proto.RegisterFile("ospf_sh_ncd_stats.proto", fileDescriptor_c69aa6118418380e) }

var fileDescriptor_c69aa6118418380e = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x65, 0x52, 0xb5, 0xe0, 0xc2, 0x40, 0x06, 0x30, 0x5b, 0xc8, 0xd4, 0xc9, 0x43, 0x61,
	0x63, 0x85, 0x09, 0x09, 0xa1, 0x54, 0x42, 0x62, 0x3a, 0x05, 0xe7, 0x22, 0x2c, 0x25, 0xb6, 0xe5,
	0x73, 0x22, 0x78, 0x04, 0xde, 0x80, 0x8d, 0x57, 0x45, 0x6e, 0x9d, 0x2e, 0x15, 0x23, 0xdb, 0xd9,
	0xf7, 0xfd, 0xdf, 0xc9, 0x67, 0x7e, 0x69, 0xc9, 0xb5, 0x40, 0xef, 0x60, 0x54, 0x03, 0x14, 0xea,
	0x40, 0xd2, 0x79, 0x1b, 0x6c, 0xfe, 0xa2, 0x34, 0x29, 0x0b, 0xda, 0x12, 0x7c, 0x78, 0xd0, 0x6e,
	0xbc, 0x85, 0x2d, 0x6a, 0x1d, 0x7a, 0x19, 0xab, 0xc8, 0x29, 0x24, 0x42, 0x9a, 0x2a, 0xd9, 0x60,
	0x5b, 0x0f, 0x5d, 0x80, 0xd1, 0xb7, 0x32, 0xba, 0x34, 0x05, 0xad, 0x48, 0x1a, 0xf2, 0xe0, 0xba,
	0x9d, 0xbd, 0xbc, 0xe3, 0x17, 0x07, 0x23, 0xe1, 0xf1, 0xe1, 0x75, 0x93, 0x5f, 0xf3, 0xd3, 0x24,
	0x02, 0x53, 0xf7, 0x28, 0x58, 0xc1, 0x56, 0x27, 0xd5, 0x32, 0xdd, 0x3d, 0xd5, 0x3d, 0x96, 0xdf,
	0x8c, 0x8b, 0x64, 0x53, 0x9d, 0x46, 0x13, 0x52, 0x3e, 0x7c, 0x3a, 0xcc, 0xaf, 0xf8, 0xb1, 0x19,
	0x7a, 0x20, 0x34, 0x41, 0xb0, 0x22, 0x5b, 0xcd, 0xaa, 0x85, 0x19, 0xfa, 0x0d, 0x9a, 0x30, 0xb5,
	0x3c, 0xaa, 0x51, 0x1c, 0xed, 0x5b, 0x15, 0xaa, 0x31, 0x2f, 0xf9, 0xd9, 0x94, 0x82, 0xc6, 0x5b,
	0x27, 0xb2, 0x22, 0x8b, 0x63, 0x53, 0xf4, 0xde, 0x5b, 0x37, 0x31, 0x31, 0xbe, 0x63, 0x66, 0x7b,
	0x26, 0x3a, 0x22, 0x53, 0xfe, 0x30, 0x7e, 0x7e, 0xf0, 0xb0, 0xfc, 0x8b, 0xf1, 0x45, 0x3c, 0x39,
	0xaf, 0xc5, 0xba, 0xc8, 0x56, 0xcb, 0xb5, 0x93, 0xff, 0xb3, 0x58, 0xf9, 0xd7, 0x5e, 0xaa, 0xb9,
	0x51, 0xcd, 0xb3, 0xd7, 0x6f, 0xf3, 0xed, 0xc7, 0xde, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xeb,
	0xf0, 0xcf, 0xcb, 0xf3, 0x01, 0x00, 0x00,
}
