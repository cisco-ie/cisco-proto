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
// source: te_stats_darkbw_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_bandwidth_accounting_counters

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

type TeStatsDarkbwInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsDarkbwInfo_KEYS) Reset()         { *m = TeStatsDarkbwInfo_KEYS{} }
func (m *TeStatsDarkbwInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsDarkbwInfo_KEYS) ProtoMessage()    {}
func (*TeStatsDarkbwInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4dbc8bb0eeff2c3, []int{0}
}

func (m *TeStatsDarkbwInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsDarkbwInfo_KEYS.Unmarshal(m, b)
}
func (m *TeStatsDarkbwInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsDarkbwInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *TeStatsDarkbwInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsDarkbwInfo_KEYS.Merge(m, src)
}
func (m *TeStatsDarkbwInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsDarkbwInfo_KEYS.Size(m)
}
func (m *TeStatsDarkbwInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsDarkbwInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsDarkbwInfo_KEYS proto.InternalMessageInfo

type TeStatsDarkbwInfo struct {
	LastClearedTimestamp      uint32   `protobuf:"varint,50,opt,name=last_cleared_timestamp,json=lastClearedTimestamp,proto3" json:"last_cleared_timestamp,omitempty"`
	BandwidthApplicationCount uint32   `protobuf:"varint,51,opt,name=bandwidth_application_count,json=bandwidthApplicationCount,proto3" json:"bandwidth_application_count,omitempty"`
	BandwidthSampleCount      uint32   `protobuf:"varint,52,opt,name=bandwidth_sample_count,json=bandwidthSampleCount,proto3" json:"bandwidth_sample_count,omitempty"`
	InvalidSampleCount        uint32   `protobuf:"varint,53,opt,name=invalid_sample_count,json=invalidSampleCount,proto3" json:"invalid_sample_count,omitempty"`
	SkippedApplicationCount   uint32   `protobuf:"varint,54,opt,name=skipped_application_count,json=skippedApplicationCount,proto3" json:"skipped_application_count,omitempty"`
	CappedSampleCount         uint32   `protobuf:"varint,55,opt,name=capped_sample_count,json=cappedSampleCount,proto3" json:"capped_sample_count,omitempty"`
	HighLatencyCount          uint32   `protobuf:"varint,56,opt,name=high_latency_count,json=highLatencyCount,proto3" json:"high_latency_count,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *TeStatsDarkbwInfo) Reset()         { *m = TeStatsDarkbwInfo{} }
func (m *TeStatsDarkbwInfo) String() string { return proto.CompactTextString(m) }
func (*TeStatsDarkbwInfo) ProtoMessage()    {}
func (*TeStatsDarkbwInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4dbc8bb0eeff2c3, []int{1}
}

func (m *TeStatsDarkbwInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsDarkbwInfo.Unmarshal(m, b)
}
func (m *TeStatsDarkbwInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsDarkbwInfo.Marshal(b, m, deterministic)
}
func (m *TeStatsDarkbwInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsDarkbwInfo.Merge(m, src)
}
func (m *TeStatsDarkbwInfo) XXX_Size() int {
	return xxx_messageInfo_TeStatsDarkbwInfo.Size(m)
}
func (m *TeStatsDarkbwInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsDarkbwInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsDarkbwInfo proto.InternalMessageInfo

func (m *TeStatsDarkbwInfo) GetLastClearedTimestamp() uint32 {
	if m != nil {
		return m.LastClearedTimestamp
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetBandwidthApplicationCount() uint32 {
	if m != nil {
		return m.BandwidthApplicationCount
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetBandwidthSampleCount() uint32 {
	if m != nil {
		return m.BandwidthSampleCount
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetInvalidSampleCount() uint32 {
	if m != nil {
		return m.InvalidSampleCount
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetSkippedApplicationCount() uint32 {
	if m != nil {
		return m.SkippedApplicationCount
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetCappedSampleCount() uint32 {
	if m != nil {
		return m.CappedSampleCount
	}
	return 0
}

func (m *TeStatsDarkbwInfo) GetHighLatencyCount() uint32 {
	if m != nil {
		return m.HighLatencyCount
	}
	return 0
}

func init() {
	proto.RegisterType((*TeStatsDarkbwInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bandwidth_accounting_counters.te_stats_darkbw_info_KEYS")
	proto.RegisterType((*TeStatsDarkbwInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bandwidth_accounting_counters.te_stats_darkbw_info")
}

func init() { proto.RegisterFile("te_stats_darkbw_info.proto", fileDescriptor_a4dbc8bb0eeff2c3) }

var fileDescriptor_a4dbc8bb0eeff2c3 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3d, 0x4b, 0x43, 0x31,
	0x14, 0x86, 0x11, 0xc1, 0x21, 0x20, 0xe8, 0xb5, 0x68, 0x6b, 0x17, 0xe9, 0xe4, 0x20, 0x17, 0xb1,
	0xf5, 0x03, 0x07, 0x45, 0x8a, 0x93, 0x4e, 0xd6, 0xc5, 0xe9, 0x70, 0x9a, 0xc4, 0xf6, 0xd0, 0xdc,
	0x24, 0x24, 0x47, 0xab, 0x7f, 0xc2, 0xdf, 0x2c, 0x37, 0xb9, 0xf6, 0x03, 0xbb, 0x25, 0x3c, 0xef,
	0x93, 0xf7, 0x85, 0x88, 0x63, 0xd6, 0x10, 0x19, 0x39, 0x82, 0xc2, 0x30, 0x1b, 0xcf, 0x81, 0xec,
	0xbb, 0x2b, 0x7d, 0x70, 0xec, 0x8a, 0x7b, 0x49, 0x51, 0x3a, 0x20, 0x17, 0xe1, 0x2b, 0x40, 0xe5,
	0x4d, 0x04, 0xd6, 0xe0, 0xbc, 0x0e, 0x65, 0x73, 0x29, 0xc7, 0x68, 0xd5, 0x9c, 0x14, 0x4f, 0x01,
	0xa5, 0x74, 0x1f, 0x96, 0xc9, 0x4e, 0x20, 0x1d, 0x74, 0x88, 0xbd, 0xae, 0xe8, 0x6c, 0x7a, 0x1e,
	0x9e, 0x1e, 0xdf, 0x46, 0xbd, 0x9f, 0x6d, 0xd1, 0xda, 0x44, 0x8b, 0x81, 0x38, 0x34, 0x18, 0x19,
	0xa4, 0xd1, 0x18, 0xb4, 0x02, 0xa6, 0x4a, 0x47, 0xc6, 0xca, 0xb7, 0x2f, 0x4e, 0xb6, 0x4e, 0x77,
	0x5f, 0x5a, 0x35, 0x1d, 0x66, 0xf8, 0xfa, 0xc7, 0x8a, 0x3b, 0xd1, 0x5d, 0x19, 0xe3, 0xbd, 0x21,
	0x89, 0x4c, 0xce, 0xe6, 0x35, 0xed, 0x7e, 0x52, 0x3b, 0x8b, 0xc8, 0xc3, 0x32, 0x31, 0xac, 0x03,
	0x75, 0xeb, 0xd2, 0x8f, 0x58, 0x79, 0xa3, 0x1b, 0x75, 0x90, 0x5b, 0x17, 0x74, 0x94, 0x60, 0xb6,
	0xce, 0x45, 0x8b, 0xec, 0x27, 0x1a, 0x52, 0xeb, 0xce, 0x65, 0x72, 0x8a, 0x86, 0xad, 0x1a, 0xb7,
	0xa2, 0x13, 0x67, 0xe4, 0xbd, 0x56, 0x1b, 0x56, 0x5e, 0x25, 0xed, 0xa8, 0x09, 0xfc, 0xdb, 0x58,
	0x8a, 0x03, 0x89, 0x49, 0x5d, 0x2b, 0xbb, 0x4e, 0xd6, 0x7e, 0x46, 0xab, 0x5d, 0x67, 0xa2, 0x98,
	0xd2, 0x64, 0x0a, 0x06, 0x59, 0x5b, 0xf9, 0xdd, 0xc4, 0x6f, 0x52, 0x7c, 0xaf, 0x26, 0xcf, 0x19,
	0xa4, 0xf4, 0x78, 0x27, 0xfd, 0x7a, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfc, 0xde, 0x42,
	0x13, 0x02, 0x00, 0x00,
}