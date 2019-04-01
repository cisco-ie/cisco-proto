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
// source: ptp_leap_seconds_info.proto

package cisco_ios_xr_ptp_oper_ptp_utc_offset_info

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

type PtpLeapSecondsInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpLeapSecondsInfo_KEYS) Reset()         { *m = PtpLeapSecondsInfo_KEYS{} }
func (m *PtpLeapSecondsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpLeapSecondsInfo_KEYS) ProtoMessage()    {}
func (*PtpLeapSecondsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d25d359c0960f84, []int{0}
}

func (m *PtpLeapSecondsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpLeapSecondsInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpLeapSecondsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpLeapSecondsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpLeapSecondsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpLeapSecondsInfo_KEYS.Merge(m, src)
}
func (m *PtpLeapSecondsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpLeapSecondsInfo_KEYS.Size(m)
}
func (m *PtpLeapSecondsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpLeapSecondsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpLeapSecondsInfo_KEYS proto.InternalMessageInfo

type PtpBagUtcOffsetInfo struct {
	Offset               string   `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Flag                 uint32   `protobuf:"varint,3,opt,name=flag,proto3" json:"flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagUtcOffsetInfo) Reset()         { *m = PtpBagUtcOffsetInfo{} }
func (m *PtpBagUtcOffsetInfo) String() string { return proto.CompactTextString(m) }
func (*PtpBagUtcOffsetInfo) ProtoMessage()    {}
func (*PtpBagUtcOffsetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d25d359c0960f84, []int{1}
}

func (m *PtpBagUtcOffsetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagUtcOffsetInfo.Unmarshal(m, b)
}
func (m *PtpBagUtcOffsetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagUtcOffsetInfo.Marshal(b, m, deterministic)
}
func (m *PtpBagUtcOffsetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagUtcOffsetInfo.Merge(m, src)
}
func (m *PtpBagUtcOffsetInfo) XXX_Size() int {
	return xxx_messageInfo_PtpBagUtcOffsetInfo.Size(m)
}
func (m *PtpBagUtcOffsetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagUtcOffsetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagUtcOffsetInfo proto.InternalMessageInfo

func (m *PtpBagUtcOffsetInfo) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *PtpBagUtcOffsetInfo) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PtpBagUtcOffsetInfo) GetFlag() uint32 {
	if m != nil {
		return m.Flag
	}
	return 0
}

type PtpBagLeapSecondsEntry struct {
	Offset               string   `protobuf:"bytes,1,opt,name=offset,proto3" json:"offset,omitempty"`
	OffsetStartDate      uint64   `protobuf:"varint,2,opt,name=offset_start_date,json=offsetStartDate,proto3" json:"offset_start_date,omitempty"`
	OffsetChange         string   `protobuf:"bytes,3,opt,name=offset_change,json=offsetChange,proto3" json:"offset_change,omitempty"`
	OffsetApplied        bool     `protobuf:"varint,4,opt,name=offset_applied,json=offsetApplied,proto3" json:"offset_applied,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagLeapSecondsEntry) Reset()         { *m = PtpBagLeapSecondsEntry{} }
func (m *PtpBagLeapSecondsEntry) String() string { return proto.CompactTextString(m) }
func (*PtpBagLeapSecondsEntry) ProtoMessage()    {}
func (*PtpBagLeapSecondsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d25d359c0960f84, []int{2}
}

func (m *PtpBagLeapSecondsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagLeapSecondsEntry.Unmarshal(m, b)
}
func (m *PtpBagLeapSecondsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagLeapSecondsEntry.Marshal(b, m, deterministic)
}
func (m *PtpBagLeapSecondsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagLeapSecondsEntry.Merge(m, src)
}
func (m *PtpBagLeapSecondsEntry) XXX_Size() int {
	return xxx_messageInfo_PtpBagLeapSecondsEntry.Size(m)
}
func (m *PtpBagLeapSecondsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagLeapSecondsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagLeapSecondsEntry proto.InternalMessageInfo

func (m *PtpBagLeapSecondsEntry) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *PtpBagLeapSecondsEntry) GetOffsetStartDate() uint64 {
	if m != nil {
		return m.OffsetStartDate
	}
	return 0
}

func (m *PtpBagLeapSecondsEntry) GetOffsetChange() string {
	if m != nil {
		return m.OffsetChange
	}
	return ""
}

func (m *PtpBagLeapSecondsEntry) GetOffsetApplied() bool {
	if m != nil {
		return m.OffsetApplied
	}
	return false
}

type PtpLeapSecondsInfo struct {
	CurrentOffsetInfo    *PtpBagUtcOffsetInfo      `protobuf:"bytes,50,opt,name=current_offset_info,json=currentOffsetInfo,proto3" json:"current_offset_info,omitempty"`
	CurrentGmOffsetInfo  *PtpBagUtcOffsetInfo      `protobuf:"bytes,51,opt,name=current_gm_offset_info,json=currentGmOffsetInfo,proto3" json:"current_gm_offset_info,omitempty"`
	ConfiguredOffsetInfo *PtpBagUtcOffsetInfo      `protobuf:"bytes,52,opt,name=configured_offset_info,json=configuredOffsetInfo,proto3" json:"configured_offset_info,omitempty"`
	PreviousGmOffsetInfo *PtpBagUtcOffsetInfo      `protobuf:"bytes,53,opt,name=previous_gm_offset_info,json=previousGmOffsetInfo,proto3" json:"previous_gm_offset_info,omitempty"`
	HardwareOffsetInfo   *PtpBagUtcOffsetInfo      `protobuf:"bytes,54,opt,name=hardware_offset_info,json=hardwareOffsetInfo,proto3" json:"hardware_offset_info,omitempty"`
	SourceType           uint32                    `protobuf:"varint,55,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	SourceFile           string                    `protobuf:"bytes,56,opt,name=source_file,json=sourceFile,proto3" json:"source_file,omitempty"`
	SourceExpiryDate     uint32                    `protobuf:"varint,57,opt,name=source_expiry_date,json=sourceExpiryDate,proto3" json:"source_expiry_date,omitempty"`
	PollingFrequency     uint32                    `protobuf:"varint,58,opt,name=polling_frequency,json=pollingFrequency,proto3" json:"polling_frequency,omitempty"`
	ConfiguredLeapSecond []*PtpBagLeapSecondsEntry `protobuf:"bytes,59,rep,name=configured_leap_second,json=configuredLeapSecond,proto3" json:"configured_leap_second,omitempty"`
	GmLeapSecond         *PtpBagLeapSecondsEntry   `protobuf:"bytes,60,opt,name=gm_leap_second,json=gmLeapSecond,proto3" json:"gm_leap_second,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PtpLeapSecondsInfo) Reset()         { *m = PtpLeapSecondsInfo{} }
func (m *PtpLeapSecondsInfo) String() string { return proto.CompactTextString(m) }
func (*PtpLeapSecondsInfo) ProtoMessage()    {}
func (*PtpLeapSecondsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d25d359c0960f84, []int{3}
}

func (m *PtpLeapSecondsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpLeapSecondsInfo.Unmarshal(m, b)
}
func (m *PtpLeapSecondsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpLeapSecondsInfo.Marshal(b, m, deterministic)
}
func (m *PtpLeapSecondsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpLeapSecondsInfo.Merge(m, src)
}
func (m *PtpLeapSecondsInfo) XXX_Size() int {
	return xxx_messageInfo_PtpLeapSecondsInfo.Size(m)
}
func (m *PtpLeapSecondsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpLeapSecondsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpLeapSecondsInfo proto.InternalMessageInfo

func (m *PtpLeapSecondsInfo) GetCurrentOffsetInfo() *PtpBagUtcOffsetInfo {
	if m != nil {
		return m.CurrentOffsetInfo
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetCurrentGmOffsetInfo() *PtpBagUtcOffsetInfo {
	if m != nil {
		return m.CurrentGmOffsetInfo
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetConfiguredOffsetInfo() *PtpBagUtcOffsetInfo {
	if m != nil {
		return m.ConfiguredOffsetInfo
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetPreviousGmOffsetInfo() *PtpBagUtcOffsetInfo {
	if m != nil {
		return m.PreviousGmOffsetInfo
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetHardwareOffsetInfo() *PtpBagUtcOffsetInfo {
	if m != nil {
		return m.HardwareOffsetInfo
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetSourceType() uint32 {
	if m != nil {
		return m.SourceType
	}
	return 0
}

func (m *PtpLeapSecondsInfo) GetSourceFile() string {
	if m != nil {
		return m.SourceFile
	}
	return ""
}

func (m *PtpLeapSecondsInfo) GetSourceExpiryDate() uint32 {
	if m != nil {
		return m.SourceExpiryDate
	}
	return 0
}

func (m *PtpLeapSecondsInfo) GetPollingFrequency() uint32 {
	if m != nil {
		return m.PollingFrequency
	}
	return 0
}

func (m *PtpLeapSecondsInfo) GetConfiguredLeapSecond() []*PtpBagLeapSecondsEntry {
	if m != nil {
		return m.ConfiguredLeapSecond
	}
	return nil
}

func (m *PtpLeapSecondsInfo) GetGmLeapSecond() *PtpBagLeapSecondsEntry {
	if m != nil {
		return m.GmLeapSecond
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpLeapSecondsInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.utc_offset_info.ptp_leap_seconds_info_KEYS")
	proto.RegisterType((*PtpBagUtcOffsetInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.utc_offset_info.ptp_bag_utc_offset_info")
	proto.RegisterType((*PtpBagLeapSecondsEntry)(nil), "cisco_ios_xr_ptp_oper.ptp.utc_offset_info.ptp_bag_leap_seconds_entry")
	proto.RegisterType((*PtpLeapSecondsInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.utc_offset_info.ptp_leap_seconds_info")
}

func init() { proto.RegisterFile("ptp_leap_seconds_info.proto", fileDescriptor_0d25d359c0960f84) }

var fileDescriptor_0d25d359c0960f84 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4b, 0x6f, 0xd4, 0x30,
	0x10, 0xc7, 0x15, 0x5a, 0x0a, 0x9d, 0x6d, 0x0b, 0x6b, 0x96, 0x6d, 0x04, 0x48, 0xac, 0x16, 0x21,
	0x85, 0x87, 0x72, 0xd8, 0xf2, 0x86, 0x0b, 0x8f, 0x2d, 0x42, 0x20, 0x21, 0x65, 0xb9, 0x20, 0x0e,
	0x96, 0x9b, 0x4c, 0x52, 0x8b, 0x6c, 0x6c, 0x1c, 0xa7, 0xdd, 0x88, 0xaf, 0xc4, 0xe7, 0xe3, 0x8c,
	0x62, 0x27, 0x6a, 0x52, 0x5a, 0x09, 0xa1, 0xdc, 0xe2, 0xff, 0xfc, 0x3d, 0xbf, 0x78, 0x3c, 0x63,
	0xb8, 0x29, 0xb5, 0xa4, 0x29, 0x32, 0x49, 0x73, 0x0c, 0x45, 0x16, 0xe5, 0x94, 0x67, 0xb1, 0xf0,
	0xa5, 0x12, 0x5a, 0x90, 0x7b, 0x21, 0xcf, 0x43, 0x41, 0xb9, 0xc8, 0xe9, 0x4a, 0xd1, 0xca, 0x29,
	0x24, 0x2a, 0x5f, 0x6a, 0xe9, 0x17, 0x3a, 0xa4, 0x22, 0x8e, 0x73, 0xd4, 0x66, 0xc3, 0xf4, 0x16,
	0xdc, 0x38, 0x33, 0x13, 0xfd, 0x38, 0xff, 0xba, 0x98, 0x7e, 0x83, 0xdd, 0x2a, 0x7a, 0xc0, 0x12,
	0x7a, 0x6a, 0x23, 0x19, 0xc3, 0x86, 0x5d, 0xba, 0xce, 0xc4, 0xf1, 0x36, 0x83, 0x7a, 0x45, 0x46,
	0x70, 0xf1, 0x88, 0xa5, 0x3c, 0x72, 0x2f, 0x4c, 0x1c, 0xef, 0x72, 0x60, 0x17, 0x84, 0xc0, 0x7a,
	0x9c, 0xb2, 0xc4, 0x5d, 0x9b, 0x38, 0xde, 0x76, 0x60, 0xbe, 0xa7, 0xbf, 0x1c, 0xcb, 0xae, 0xb2,
	0x77, 0xf8, 0x98, 0x69, 0x55, 0x9e, 0x0b, 0xb8, 0x0f, 0xc3, 0xfa, 0x3f, 0x72, 0xcd, 0x94, 0xa6,
	0x11, 0xd3, 0x68, 0x60, 0xeb, 0xc1, 0x15, 0x1b, 0x58, 0x54, 0xfa, 0x3b, 0xa6, 0x91, 0xdc, 0x81,
	0xed, 0xda, 0x1b, 0x1e, 0xb2, 0x2c, 0x41, 0xc3, 0xdf, 0x0c, 0xb6, 0xac, 0xf8, 0xd6, 0x68, 0xe4,
	0x2e, 0xec, 0xd4, 0x26, 0x26, 0x65, 0xca, 0x31, 0x72, 0xd7, 0xcd, 0xaf, 0xd7, 0x5b, 0x5f, 0x5b,
	0x71, 0xfa, 0xfb, 0x12, 0x5c, 0x3f, 0xb3, 0x54, 0x44, 0xc1, 0xb5, 0xb0, 0x50, 0x0a, 0x33, 0xdd,
	0xae, 0x90, 0x3b, 0x9b, 0x38, 0xde, 0x60, 0xf6, 0xc6, 0xff, 0xe7, 0xcb, 0xf0, 0xcf, 0xa9, 0x75,
	0x30, 0xac, 0xd3, 0x7f, 0x36, 0xda, 0x87, 0x8a, 0x79, 0x0c, 0xe3, 0x86, 0x99, 0x2c, 0x3b, 0xd8,
	0xbd, 0xde, 0xb0, 0xcd, 0xa9, 0xde, 0x2f, 0x5b, 0xe0, 0x15, 0x8c, 0x43, 0x91, 0xc5, 0x3c, 0x29,
	0x14, 0x46, 0x1d, 0xf0, 0xa3, 0xde, 0xc0, 0xa3, 0x13, 0x42, 0x8b, 0x5c, 0xc2, 0xae, 0x54, 0x78,
	0xc4, 0x45, 0x91, 0x9f, 0x3e, 0xf3, 0xe3, 0xfe, 0xd0, 0x0d, 0xa2, 0x73, 0x68, 0x0d, 0xa3, 0x43,
	0xa6, 0xa2, 0x63, 0xa6, 0xb0, 0xc3, 0x7d, 0xd2, 0x1b, 0x97, 0x34, 0xf9, 0x5b, 0xd4, 0xdb, 0x30,
	0xc8, 0x45, 0xa1, 0x42, 0xa4, 0xba, 0x94, 0xe8, 0x3e, 0x35, 0xb3, 0x03, 0x56, 0xfa, 0x52, 0x4a,
	0x6c, 0x19, 0x62, 0x9e, 0xa2, 0xfb, 0xcc, 0x34, 0x77, 0x6d, 0xd8, 0xe7, 0x29, 0x92, 0x87, 0x40,
	0x6a, 0x03, 0xae, 0x24, 0x57, 0xa5, 0x1d, 0x96, 0xe7, 0x26, 0xd1, 0x55, 0x1b, 0x99, 0x9b, 0x80,
	0x99, 0x96, 0x07, 0x30, 0x94, 0x22, 0x4d, 0x79, 0x96, 0xd0, 0x58, 0xe1, 0x8f, 0x02, 0xb3, 0xb0,
	0x74, 0x5f, 0x58, 0x73, 0x1d, 0xd8, 0x6f, 0x74, 0xf2, 0xb3, 0xd3, 0x07, 0xad, 0xa1, 0x70, 0x5f,
	0x4e, 0xd6, 0xbc, 0xc1, 0x6c, 0xfe, 0x1f, 0x45, 0xf9, 0xfb, 0x15, 0x68, 0xb7, 0xc2, 0x27, 0x64,
	0x72, 0x61, 0x82, 0xe4, 0x3b, 0xec, 0x24, 0xcb, 0x0e, 0xf4, 0x95, 0xb9, 0x89, 0x9e, 0xa0, 0x5b,
	0xc9, 0xf2, 0x04, 0x76, 0xb0, 0x61, 0x1e, 0xd5, 0xbd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98,
	0x69, 0x8e, 0xdf, 0x73, 0x05, 0x00, 0x00,
}
