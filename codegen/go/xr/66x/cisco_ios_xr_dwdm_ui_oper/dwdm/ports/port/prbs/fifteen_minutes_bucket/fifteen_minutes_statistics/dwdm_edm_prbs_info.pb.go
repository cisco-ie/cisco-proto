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
// source: dwdm_edm_prbs_info.proto

package cisco_ios_xr_dwdm_ui_oper_dwdm_ports_port_prbs_fifteen_minutes_bucket_fifteen_minutes_statistics

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

type DwdmEdmPrbsInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DwdmEdmPrbsInfo_KEYS) Reset()         { *m = DwdmEdmPrbsInfo_KEYS{} }
func (m *DwdmEdmPrbsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*DwdmEdmPrbsInfo_KEYS) ProtoMessage()    {}
func (*DwdmEdmPrbsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f688e45422f29b9, []int{0}
}

func (m *DwdmEdmPrbsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DwdmEdmPrbsInfo_KEYS.Unmarshal(m, b)
}
func (m *DwdmEdmPrbsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DwdmEdmPrbsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *DwdmEdmPrbsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DwdmEdmPrbsInfo_KEYS.Merge(m, src)
}
func (m *DwdmEdmPrbsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_DwdmEdmPrbsInfo_KEYS.Size(m)
}
func (m *DwdmEdmPrbsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DwdmEdmPrbsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DwdmEdmPrbsInfo_KEYS proto.InternalMessageInfo

func (m *DwdmEdmPrbsInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PrbsInterval struct {
	IntervalIndex        string   `protobuf:"bytes,1,opt,name=interval_index,json=intervalIndex,proto3" json:"interval_index,omitempty"`
	ConfiguredPattern    string   `protobuf:"bytes,2,opt,name=configured_pattern,json=configuredPattern,proto3" json:"configured_pattern,omitempty"`
	StartAt              string   `protobuf:"bytes,3,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	StopAt               string   `protobuf:"bytes,4,opt,name=stop_at,json=stopAt,proto3" json:"stop_at,omitempty"`
	ReceivedPattern      string   `protobuf:"bytes,5,opt,name=received_pattern,json=receivedPattern,proto3" json:"received_pattern,omitempty"`
	BitErrorCount        uint64   `protobuf:"varint,6,opt,name=bit_error_count,json=bitErrorCount,proto3" json:"bit_error_count,omitempty"`
	FoundCount           uint64   `protobuf:"varint,7,opt,name=found_count,json=foundCount,proto3" json:"found_count,omitempty"`
	LostCount            uint64   `protobuf:"varint,8,opt,name=lost_count,json=lostCount,proto3" json:"lost_count,omitempty"`
	FoundAt              string   `protobuf:"bytes,9,opt,name=found_at,json=foundAt,proto3" json:"found_at,omitempty"`
	LostAt               string   `protobuf:"bytes,10,opt,name=lost_at,json=lostAt,proto3" json:"lost_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrbsInterval) Reset()         { *m = PrbsInterval{} }
func (m *PrbsInterval) String() string { return proto.CompactTextString(m) }
func (*PrbsInterval) ProtoMessage()    {}
func (*PrbsInterval) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f688e45422f29b9, []int{1}
}

func (m *PrbsInterval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrbsInterval.Unmarshal(m, b)
}
func (m *PrbsInterval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrbsInterval.Marshal(b, m, deterministic)
}
func (m *PrbsInterval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrbsInterval.Merge(m, src)
}
func (m *PrbsInterval) XXX_Size() int {
	return xxx_messageInfo_PrbsInterval.Size(m)
}
func (m *PrbsInterval) XXX_DiscardUnknown() {
	xxx_messageInfo_PrbsInterval.DiscardUnknown(m)
}

var xxx_messageInfo_PrbsInterval proto.InternalMessageInfo

func (m *PrbsInterval) GetIntervalIndex() string {
	if m != nil {
		return m.IntervalIndex
	}
	return ""
}

func (m *PrbsInterval) GetConfiguredPattern() string {
	if m != nil {
		return m.ConfiguredPattern
	}
	return ""
}

func (m *PrbsInterval) GetStartAt() string {
	if m != nil {
		return m.StartAt
	}
	return ""
}

func (m *PrbsInterval) GetStopAt() string {
	if m != nil {
		return m.StopAt
	}
	return ""
}

func (m *PrbsInterval) GetReceivedPattern() string {
	if m != nil {
		return m.ReceivedPattern
	}
	return ""
}

func (m *PrbsInterval) GetBitErrorCount() uint64 {
	if m != nil {
		return m.BitErrorCount
	}
	return 0
}

func (m *PrbsInterval) GetFoundCount() uint64 {
	if m != nil {
		return m.FoundCount
	}
	return 0
}

func (m *PrbsInterval) GetLostCount() uint64 {
	if m != nil {
		return m.LostCount
	}
	return 0
}

func (m *PrbsInterval) GetFoundAt() string {
	if m != nil {
		return m.FoundAt
	}
	return ""
}

func (m *PrbsInterval) GetLostAt() string {
	if m != nil {
		return m.LostAt
	}
	return ""
}

type DwdmEdmPrbsInfo struct {
	PrbsEntry            []*PrbsInterval `protobuf:"bytes,50,rep,name=prbs_entry,json=prbsEntry,proto3" json:"prbs_entry,omitempty"`
	IsPrbsEnabled        bool            `protobuf:"varint,51,opt,name=is_prbs_enabled,json=isPrbsEnabled,proto3" json:"is_prbs_enabled,omitempty"`
	PrbsConfigMode       string          `protobuf:"bytes,52,opt,name=prbs_config_mode,json=prbsConfigMode,proto3" json:"prbs_config_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DwdmEdmPrbsInfo) Reset()         { *m = DwdmEdmPrbsInfo{} }
func (m *DwdmEdmPrbsInfo) String() string { return proto.CompactTextString(m) }
func (*DwdmEdmPrbsInfo) ProtoMessage()    {}
func (*DwdmEdmPrbsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f688e45422f29b9, []int{2}
}

func (m *DwdmEdmPrbsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DwdmEdmPrbsInfo.Unmarshal(m, b)
}
func (m *DwdmEdmPrbsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DwdmEdmPrbsInfo.Marshal(b, m, deterministic)
}
func (m *DwdmEdmPrbsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DwdmEdmPrbsInfo.Merge(m, src)
}
func (m *DwdmEdmPrbsInfo) XXX_Size() int {
	return xxx_messageInfo_DwdmEdmPrbsInfo.Size(m)
}
func (m *DwdmEdmPrbsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DwdmEdmPrbsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DwdmEdmPrbsInfo proto.InternalMessageInfo

func (m *DwdmEdmPrbsInfo) GetPrbsEntry() []*PrbsInterval {
	if m != nil {
		return m.PrbsEntry
	}
	return nil
}

func (m *DwdmEdmPrbsInfo) GetIsPrbsEnabled() bool {
	if m != nil {
		return m.IsPrbsEnabled
	}
	return false
}

func (m *DwdmEdmPrbsInfo) GetPrbsConfigMode() string {
	if m != nil {
		return m.PrbsConfigMode
	}
	return ""
}

func init() {
	proto.RegisterType((*DwdmEdmPrbsInfo_KEYS)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.fifteen_minutes_bucket.fifteen_minutes_statistics.dwdm_edm_prbs_info_KEYS")
	proto.RegisterType((*PrbsInterval)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.fifteen_minutes_bucket.fifteen_minutes_statistics.prbs_interval")
	proto.RegisterType((*DwdmEdmPrbsInfo)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.fifteen_minutes_bucket.fifteen_minutes_statistics.dwdm_edm_prbs_info")
}

func init() { proto.RegisterFile("dwdm_edm_prbs_info.proto", fileDescriptor_4f688e45422f29b9) }

var fileDescriptor_4f688e45422f29b9 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0x76, 0xd9, 0x3f, 0x53, 0x6d, 0x5b, 0x7c, 0xa9, 0x7b, 0x40, 0xac, 0x56, 0x02,
	0x85, 0x43, 0x73, 0x68, 0x79, 0x81, 0x55, 0xb5, 0x07, 0x84, 0x90, 0xaa, 0xe5, 0xc4, 0x69, 0x48,
	0xe2, 0x09, 0xb2, 0xd8, 0xd8, 0x91, 0x3d, 0x29, 0xe5, 0x25, 0x78, 0x46, 0xde, 0x82, 0x2b, 0xf2,
	0x24, 0x2b, 0x40, 0x7b, 0xe6, 0x12, 0xf9, 0xfb, 0xbe, 0xdf, 0xc4, 0x1e, 0x8f, 0x41, 0x9b, 0x6f,
	0xa6, 0x45, 0x32, 0x2d, 0x76, 0xa1, 0x8a, 0x68, 0x5d, 0xe3, 0x8b, 0x2e, 0x78, 0xf6, 0xea, 0x73,
	0x6d, 0x63, 0xed, 0xd1, 0xfa, 0x88, 0x4f, 0x01, 0x05, 0xeb, 0x2d, 0xfa, 0x8e, 0x42, 0x91, 0x44,
	0xd1, 0xf9, 0xc0, 0x51, 0xbe, 0x45, 0x2a, 0x2d, 0x1a, 0xdb, 0x30, 0x91, 0xc3, 0xd6, 0xba, 0x9e,
	0x29, 0x62, 0xd5, 0xd7, 0x5f, 0x89, 0x8f, 0xec, 0xc8, 0x25, 0xdb, 0xc8, 0xb6, 0x8e, 0xeb, 0x1b,
	0xb8, 0x3a, 0xde, 0x1d, 0xdf, 0x6f, 0x3f, 0x7d, 0x54, 0x0a, 0x26, 0xae, 0x6c, 0x49, 0x67, 0xab,
	0x2c, 0x5f, 0xec, 0x64, 0xbd, 0xfe, 0x79, 0x02, 0xcb, 0x11, 0x63, 0x0a, 0x8f, 0xe5, 0x5e, 0xbd,
	0x82, 0xf3, 0xc3, 0x1a, 0xad, 0x33, 0xf4, 0x34, 0xf2, 0xcb, 0x83, 0xfb, 0x2e, 0x99, 0xea, 0x06,
	0x54, 0xed, 0x5d, 0x63, 0xbf, 0xf4, 0x81, 0x0c, 0x76, 0x25, 0x33, 0x05, 0xa7, 0x4f, 0x04, 0x7d,
	0xfe, 0x27, 0x79, 0x18, 0x02, 0x75, 0x0d, 0xf3, 0xc8, 0x65, 0x60, 0x2c, 0x59, 0x9f, 0x0a, 0x34,
	0x13, 0xbd, 0x61, 0x75, 0x05, 0xb3, 0xc8, 0xbe, 0x4b, 0xc9, 0x44, 0x92, 0x69, 0x92, 0x1b, 0x56,
	0x6f, 0xe0, 0x32, 0x50, 0x4d, 0xf6, 0xf1, 0xaf, 0x0d, 0x9e, 0x09, 0x71, 0x71, 0xf0, 0x0f, 0xbf,
	0x7f, 0x0d, 0x17, 0x95, 0x65, 0xa4, 0x10, 0x7c, 0xc0, 0xda, 0xf7, 0x8e, 0xf5, 0x74, 0x95, 0xe5,
	0x93, 0xdd, 0xb2, 0xb2, 0xbc, 0x4d, 0xee, 0x7d, 0x32, 0xd5, 0x4b, 0x38, 0x6b, 0x7c, 0xef, 0xcc,
	0xc8, 0xcc, 0x84, 0x01, 0xb1, 0x06, 0xe0, 0x05, 0xc0, 0xde, 0x47, 0x1e, 0xf3, 0xb9, 0xe4, 0x8b,
	0xe4, 0x0c, 0xf1, 0x35, 0xcc, 0x87, 0xfa, 0x92, 0xf5, 0x62, 0x68, 0x43, 0xf4, 0xd0, 0x86, 0x54,
	0x96, 0xac, 0x61, 0x68, 0x23, 0xc9, 0x0d, 0xaf, 0x7f, 0x65, 0xa0, 0x8e, 0x47, 0xa2, 0x7e, 0x64,
	0x00, 0xa2, 0xc8, 0x71, 0xf8, 0xae, 0x6f, 0x57, 0xa7, 0xf9, 0xd9, 0xad, 0x2f, 0xfe, 0xf7, 0x03,
	0x29, 0xfe, 0x99, 0xf6, 0x6e, 0x91, 0xe4, 0x36, 0x9d, 0x20, 0xdd, 0xa1, 0x8d, 0x38, 0x1e, 0xa9,
	0xac, 0xf6, 0x64, 0xf4, 0xdd, 0x2a, 0xcb, 0xe7, 0xbb, 0xa5, 0x8d, 0x0f, 0x42, 0x89, 0xa9, 0x72,
	0xb8, 0x14, 0x68, 0x18, 0x32, 0xb6, 0xde, 0x90, 0x7e, 0x2b, 0x1d, 0x9f, 0x27, 0xff, 0x5e, 0xec,
	0x0f, 0xde, 0x50, 0x35, 0x95, 0x47, 0x7f, 0xf7, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x65, 0xe2, 0x02,
	0x2a, 0x10, 0x03, 0x00, 0x00,
}
