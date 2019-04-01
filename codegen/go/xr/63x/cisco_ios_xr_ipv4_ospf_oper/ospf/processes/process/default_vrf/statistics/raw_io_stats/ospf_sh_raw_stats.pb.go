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
// source: ospf_sh_raw_stats.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_statistics_raw_io_stats

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

type OspfShRawStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRawStats_KEYS) Reset()         { *m = OspfShRawStats_KEYS{} }
func (m *OspfShRawStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRawStats_KEYS) ProtoMessage()    {}
func (*OspfShRawStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_071b990fcaaffdfa, []int{0}
}

func (m *OspfShRawStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRawStats_KEYS.Unmarshal(m, b)
}
func (m *OspfShRawStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRawStats_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRawStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRawStats_KEYS.Merge(m, src)
}
func (m *OspfShRawStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRawStats_KEYS.Size(m)
}
func (m *OspfShRawStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRawStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRawStats_KEYS proto.InternalMessageInfo

func (m *OspfShRawStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShRawStats struct {
	OutSent                uint32   `protobuf:"varint,50,opt,name=out_sent,json=outSent,proto3" json:"out_sent,omitempty"`
	OutBytesSent           uint32   `protobuf:"varint,51,opt,name=out_bytes_sent,json=outBytesSent,proto3" json:"out_bytes_sent,omitempty"`
	OutNoConnDiscarded     uint32   `protobuf:"varint,52,opt,name=out_no_conn_discarded,json=outNoConnDiscarded,proto3" json:"out_no_conn_discarded,omitempty"`
	OutNullSrcDiscarded    uint32   `protobuf:"varint,53,opt,name=out_null_src_discarded,json=outNullSrcDiscarded,proto3" json:"out_null_src_discarded,omitempty"`
	OutNoPakDiscarded      uint32   `protobuf:"varint,54,opt,name=out_no_pak_discarded,json=outNoPakDiscarded,proto3" json:"out_no_pak_discarded,omitempty"`
	OutIpv4HdrErrDiscarded uint32   `protobuf:"varint,55,opt,name=out_ipv4_hdr_err_discarded,json=outIpv4HdrErrDiscarded,proto3" json:"out_ipv4_hdr_err_discarded,omitempty"`
	OutSendPakErrDiscarded uint32   `protobuf:"varint,56,opt,name=out_send_pak_err_discarded,json=outSendPakErrDiscarded,proto3" json:"out_send_pak_err_discarded,omitempty"`
	InRcv                  uint32   `protobuf:"varint,57,opt,name=in_rcv,json=inRcv,proto3" json:"in_rcv,omitempty"`
	InBytesRcv             uint32   `protobuf:"varint,58,opt,name=in_bytes_rcv,json=inBytesRcv,proto3" json:"in_bytes_rcv,omitempty"`
	InShortMsgDiscarded    uint32   `protobuf:"varint,59,opt,name=in_short_msg_discarded,json=inShortMsgDiscarded,proto3" json:"in_short_msg_discarded,omitempty"`
	InNoMemDiscarded       uint32   `protobuf:"varint,60,opt,name=in_no_mem_discarded,json=inNoMemDiscarded,proto3" json:"in_no_mem_discarded,omitempty"`
	InRawEvent             uint32   `protobuf:"varint,61,opt,name=in_raw_event,json=inRawEvent,proto3" json:"in_raw_event,omitempty"`
	Disconnects            uint32   `protobuf:"varint,62,opt,name=disconnects,proto3" json:"disconnects,omitempty"`
	InStandbyDiscarded     uint32   `protobuf:"varint,63,opt,name=in_standby_discarded,json=inStandbyDiscarded,proto3" json:"in_standby_discarded,omitempty"`
	NsrNotReadyDiscarded   uint32   `protobuf:"varint,64,opt,name=nsr_not_ready_discarded,json=nsrNotReadyDiscarded,proto3" json:"nsr_not_ready_discarded,omitempty"`
	RawConnected           bool     `protobuf:"varint,65,opt,name=raw_connected,json=rawConnected,proto3" json:"raw_connected,omitempty"`
	SlRawConnected         bool     `protobuf:"varint,66,opt,name=sl_raw_connected,json=slRawConnected,proto3" json:"sl_raw_connected,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *OspfShRawStats) Reset()         { *m = OspfShRawStats{} }
func (m *OspfShRawStats) String() string { return proto.CompactTextString(m) }
func (*OspfShRawStats) ProtoMessage()    {}
func (*OspfShRawStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_071b990fcaaffdfa, []int{1}
}

func (m *OspfShRawStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRawStats.Unmarshal(m, b)
}
func (m *OspfShRawStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRawStats.Marshal(b, m, deterministic)
}
func (m *OspfShRawStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRawStats.Merge(m, src)
}
func (m *OspfShRawStats) XXX_Size() int {
	return xxx_messageInfo_OspfShRawStats.Size(m)
}
func (m *OspfShRawStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRawStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRawStats proto.InternalMessageInfo

func (m *OspfShRawStats) GetOutSent() uint32 {
	if m != nil {
		return m.OutSent
	}
	return 0
}

func (m *OspfShRawStats) GetOutBytesSent() uint32 {
	if m != nil {
		return m.OutBytesSent
	}
	return 0
}

func (m *OspfShRawStats) GetOutNoConnDiscarded() uint32 {
	if m != nil {
		return m.OutNoConnDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetOutNullSrcDiscarded() uint32 {
	if m != nil {
		return m.OutNullSrcDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetOutNoPakDiscarded() uint32 {
	if m != nil {
		return m.OutNoPakDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetOutIpv4HdrErrDiscarded() uint32 {
	if m != nil {
		return m.OutIpv4HdrErrDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetOutSendPakErrDiscarded() uint32 {
	if m != nil {
		return m.OutSendPakErrDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetInRcv() uint32 {
	if m != nil {
		return m.InRcv
	}
	return 0
}

func (m *OspfShRawStats) GetInBytesRcv() uint32 {
	if m != nil {
		return m.InBytesRcv
	}
	return 0
}

func (m *OspfShRawStats) GetInShortMsgDiscarded() uint32 {
	if m != nil {
		return m.InShortMsgDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetInNoMemDiscarded() uint32 {
	if m != nil {
		return m.InNoMemDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetInRawEvent() uint32 {
	if m != nil {
		return m.InRawEvent
	}
	return 0
}

func (m *OspfShRawStats) GetDisconnects() uint32 {
	if m != nil {
		return m.Disconnects
	}
	return 0
}

func (m *OspfShRawStats) GetInStandbyDiscarded() uint32 {
	if m != nil {
		return m.InStandbyDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetNsrNotReadyDiscarded() uint32 {
	if m != nil {
		return m.NsrNotReadyDiscarded
	}
	return 0
}

func (m *OspfShRawStats) GetRawConnected() bool {
	if m != nil {
		return m.RawConnected
	}
	return false
}

func (m *OspfShRawStats) GetSlRawConnected() bool {
	if m != nil {
		return m.SlRawConnected
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShRawStats_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.statistics.raw_io_stats.ospf_sh_raw_stats_KEYS")
	proto.RegisterType((*OspfShRawStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.statistics.raw_io_stats.ospf_sh_raw_stats")
}

func init() { proto.RegisterFile("ospf_sh_raw_stats.proto", fileDescriptor_071b990fcaaffdfa) }

var fileDescriptor_071b990fcaaffdfa = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcb, 0x4f, 0xdb, 0x40,
	0x10, 0x87, 0x95, 0x03, 0x94, 0x2e, 0x01, 0x81, 0x81, 0xe0, 0xf6, 0x94, 0xd2, 0x1e, 0x72, 0x69,
	0xfa, 0x48, 0xe8, 0x03, 0xfa, 0x84, 0x46, 0x6a, 0x55, 0x11, 0x21, 0x5b, 0xaa, 0xd4, 0xd3, 0x6a,
	0xe3, 0xdd, 0x90, 0x55, 0xec, 0x59, 0x6b, 0x67, 0xed, 0x94, 0x7f, 0xbe, 0xaa, 0x66, 0x9d, 0xd4,
	0x8e, 0x72, 0xb3, 0x66, 0xbe, 0x6f, 0x7e, 0xa3, 0x59, 0x99, 0x9d, 0x1a, 0xcc, 0xa7, 0x1c, 0x67,
	0xdc, 0x8a, 0x05, 0x47, 0x27, 0x1c, 0xf6, 0x73, 0x6b, 0x9c, 0x09, 0x7e, 0x25, 0x1a, 0x13, 0xc3,
	0xb5, 0x41, 0xfe, 0xc7, 0x72, 0x9d, 0x97, 0x43, 0xee, 0x51, 0x93, 0x2b, 0xdb, 0xa7, 0x2f, 0xe2,
	0x12, 0x85, 0xa8, 0x70, 0xf5, 0xd5, 0x97, 0x6a, 0x2a, 0x8a, 0xd4, 0xf1, 0xd2, 0x4e, 0xfb, 0x34,
	0x4b, 0xa3, 0xd3, 0x09, 0xf6, 0x69, 0xb4, 0x36, 0xd5, 0xf4, 0xb3, 0x4b, 0xd6, 0xd9, 0x88, 0xe4,
	0x3f, 0x47, 0xbf, 0xe3, 0xe0, 0x09, 0x6b, 0x2f, 0x07, 0x71, 0x10, 0x99, 0x0a, 0x5b, 0xdd, 0x56,
	0xef, 0x61, 0xb4, 0xbb, 0xac, 0x8d, 0x45, 0xa6, 0xce, 0xfe, 0x6e, 0xb1, 0xc3, 0x0d, 0x3b, 0x78,
	0xc4, 0x76, 0x4c, 0xe1, 0x38, 0x2a, 0x70, 0xe1, 0xeb, 0x6e, 0xab, 0xb7, 0x17, 0x3d, 0x30, 0x85,
	0x8b, 0x15, 0xb8, 0xe0, 0x19, 0xdb, 0xa7, 0xd6, 0xe4, 0xde, 0x29, 0xac, 0x80, 0x81, 0x07, 0xda,
	0xa6, 0x70, 0x57, 0x54, 0xf4, 0xd4, 0x2b, 0x76, 0x42, 0x14, 0x18, 0x9e, 0x18, 0x00, 0x2e, 0x35,
	0x26, 0xc2, 0x4a, 0x25, 0xc3, 0xa1, 0x87, 0x03, 0x53, 0xb8, 0xb1, 0xb9, 0x36, 0x00, 0xdf, 0x56,
	0x9d, 0x60, 0xc0, 0x3a, 0x5e, 0x29, 0xd2, 0x94, 0xa3, 0x4d, 0x1a, 0xce, 0xb9, 0x77, 0x8e, 0xc8,
	0x29, 0xd2, 0x34, 0xb6, 0x49, 0x2d, 0xbd, 0x60, 0xc7, 0xcb, 0x9c, 0x5c, 0xcc, 0x1b, 0xca, 0x1b,
	0xaf, 0x1c, 0xfa, 0x98, 0x5b, 0x31, 0xaf, 0x85, 0x0b, 0xf6, 0x98, 0x04, 0x7f, 0xfd, 0x99, 0xb4,
	0x5c, 0x59, 0xdb, 0xd0, 0xde, 0x7a, 0x8d, 0xf6, 0xf8, 0x91, 0x97, 0xc3, 0xef, 0xd2, 0x8e, 0xac,
	0xdd, 0x70, 0x51, 0x81, 0xf4, 0x71, 0xeb, 0xee, 0xbb, 0xff, 0x6e, 0xac, 0x40, 0xde, 0x8a, 0xf9,
	0x9a, 0x7b, 0xc2, 0xb6, 0x35, 0x70, 0x9b, 0x94, 0xe1, 0x7b, 0xcf, 0x6d, 0x69, 0x88, 0x92, 0x32,
	0xe8, 0xb2, 0xb6, 0x86, 0xe5, 0x31, 0xa9, 0x79, 0xe1, 0x9b, 0x4c, 0x83, 0x3f, 0x25, 0x11, 0x03,
	0xd6, 0xd1, 0xc0, 0x71, 0x66, 0xac, 0xe3, 0x19, 0xde, 0x35, 0x02, 0x2f, 0xab, 0xb3, 0x68, 0x88,
	0xa9, 0x79, 0x83, 0x77, 0x75, 0xda, 0x73, 0x76, 0xa4, 0x81, 0xae, 0x92, 0xa9, 0xac, 0x61, 0x7c,
	0xf0, 0xc6, 0x81, 0x86, 0xb1, 0xb9, 0x51, 0x59, 0x8d, 0x57, 0x5b, 0xd0, 0xf3, 0xab, 0x92, 0x5e,
	0xf4, 0xe3, 0x6a, 0x8b, 0x48, 0x2c, 0x46, 0x54, 0x09, 0xba, 0x6c, 0x97, 0xc6, 0x18, 0x00, 0x95,
	0x38, 0x0c, 0x3f, 0x79, 0xa0, 0x59, 0x0a, 0x5e, 0xb2, 0x63, 0xda, 0xd3, 0x09, 0x90, 0x93, 0xfb,
	0x46, 0xe6, 0xe7, 0xea, 0xc1, 0x35, 0xc4, 0x55, 0xab, 0x4e, 0x3d, 0x67, 0xa7, 0x80, 0x96, 0x83,
	0x71, 0xdc, 0x2a, 0x21, 0x9b, 0xd2, 0x17, 0x2f, 0x1d, 0x03, 0xda, 0xb1, 0x71, 0x11, 0x35, 0x6b,
	0xed, 0x29, 0xdb, 0xa3, 0x4d, 0x97, 0xc1, 0x4a, 0x86, 0x5f, 0xbb, 0xad, 0xde, 0x4e, 0xd4, 0xb6,
	0x62, 0x71, 0xbd, 0xaa, 0x05, 0x3d, 0x76, 0x80, 0x29, 0x5f, 0xe7, 0xae, 0x3c, 0xb7, 0x8f, 0x69,
	0xd4, 0x20, 0x27, 0xdb, 0xfe, 0xe7, 0x1c, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x50, 0x4e, 0xe6,
	0xcf, 0xb7, 0x03, 0x00, 0x00,
}