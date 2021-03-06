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

package cisco_ios_xr_dwdm_ui_oper_dwdm_ports_port_prbs_twenty_four_hours_bucket_twenty_four_hours_statistics

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
	proto.RegisterType((*DwdmEdmPrbsInfo_KEYS)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.twenty_four_hours_bucket.twenty_four_hours_statistics.dwdm_edm_prbs_info_KEYS")
	proto.RegisterType((*PrbsInterval)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.twenty_four_hours_bucket.twenty_four_hours_statistics.prbs_interval")
	proto.RegisterType((*DwdmEdmPrbsInfo)(nil), "cisco_ios_xr_dwdm_ui_oper.dwdm.ports.port.prbs.twenty_four_hours_bucket.twenty_four_hours_statistics.dwdm_edm_prbs_info")
}

func init() { proto.RegisterFile("dwdm_edm_prbs_info.proto", fileDescriptor_4f688e45422f29b9) }

var fileDescriptor_4f688e45422f29b9 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcb, 0x6e, 0xdb, 0x3c,
	0x10, 0x85, 0xa1, 0xc4, 0xbf, 0x2f, 0x13, 0x38, 0xc9, 0xcf, 0x4d, 0x98, 0x45, 0x51, 0xc3, 0x40,
	0x0b, 0x75, 0x11, 0x2d, 0x92, 0xbe, 0x80, 0x11, 0x78, 0x51, 0x14, 0x05, 0x02, 0x77, 0xd5, 0xd5,
	0x80, 0x12, 0xe9, 0x96, 0xa8, 0xcd, 0x11, 0xc8, 0x51, 0x2e, 0xaf, 0xd1, 0x27, 0xec, 0x7b, 0x74,
	0x53, 0x70, 0x24, 0xa3, 0x2d, 0xd2, 0x75, 0x37, 0x02, 0xcf, 0x39, 0xdf, 0x48, 0x33, 0x1a, 0x82,
	0xb6, 0x0f, 0x76, 0x8f, 0xce, 0xee, 0xb1, 0x8d, 0x75, 0x42, 0x1f, 0xb6, 0x54, 0xb5, 0x91, 0x98,
	0x94, 0x6d, 0x7c, 0x6a, 0x08, 0x3d, 0x25, 0x7c, 0x8c, 0x28, 0x58, 0xe7, 0x91, 0x5a, 0x17, 0xab,
	0x2c, 0xaa, 0x96, 0x22, 0x27, 0x79, 0x56, 0xb9, 0xb4, 0xe2, 0x07, 0x17, 0xf8, 0x09, 0xb7, 0xd4,
	0x45, 0xfc, 0x42, 0x5d, 0x4c, 0x58, 0x77, 0xcd, 0x57, 0xc7, 0x7f, 0x09, 0x12, 0x1b, 0xf6, 0x89,
	0x7d, 0x93, 0x96, 0x57, 0x70, 0xf1, 0xbc, 0x03, 0x7c, 0xbf, 0xfe, 0xf4, 0x51, 0x29, 0x18, 0x05,
	0xb3, 0x77, 0xba, 0x58, 0x14, 0xe5, 0x6c, 0x23, 0xe7, 0xe5, 0xf7, 0x23, 0x98, 0x0f, 0x18, 0xbb,
	0x78, 0x6f, 0x76, 0xea, 0x15, 0x9c, 0x1e, 0xce, 0xe8, 0x83, 0x75, 0x8f, 0x03, 0x3f, 0x3f, 0xb8,
	0xef, 0xb2, 0xa9, 0xae, 0x40, 0x35, 0x14, 0xb6, 0xfe, 0x73, 0x17, 0x9d, 0xc5, 0xd6, 0x30, 0xbb,
	0x18, 0xf4, 0x91, 0xa0, 0xff, 0xff, 0x4a, 0xee, 0xfa, 0x40, 0x5d, 0xc2, 0x34, 0xb1, 0x89, 0x8c,
	0x86, 0xf5, 0xb1, 0x40, 0x13, 0xd1, 0x2b, 0x56, 0x17, 0x30, 0x49, 0x4c, 0x6d, 0x4e, 0x46, 0x92,
	0x8c, 0xb3, 0x5c, 0xb1, 0x7a, 0x03, 0xe7, 0xd1, 0x35, 0xce, 0xdf, 0xff, 0xf6, 0x81, 0xff, 0x84,
	0x38, 0x3b, 0xf8, 0x87, 0xd7, 0xbf, 0x86, 0xb3, 0xda, 0x33, 0xba, 0x18, 0x29, 0x62, 0x43, 0x5d,
	0x60, 0x3d, 0x5e, 0x14, 0xe5, 0x68, 0x33, 0xaf, 0x3d, 0xaf, 0xb3, 0x7b, 0x9b, 0x4d, 0xf5, 0x12,
	0x4e, 0xb6, 0xd4, 0x05, 0x3b, 0x30, 0x13, 0x61, 0x40, 0xac, 0x1e, 0x78, 0x01, 0xb0, 0xa3, 0xc4,
	0x43, 0x3e, 0x95, 0x7c, 0x96, 0x9d, 0x3e, 0xbe, 0x84, 0x69, 0x5f, 0x6f, 0x58, 0xcf, 0xfa, 0x31,
	0x44, 0xf7, 0x63, 0x48, 0xa5, 0x61, 0x0d, 0xfd, 0x18, 0x59, 0xae, 0x78, 0xf9, 0xa3, 0x00, 0xf5,
	0x7c, 0x25, 0xea, 0x5b, 0x01, 0x20, 0xca, 0x05, 0x8e, 0x4f, 0xfa, 0x7a, 0x71, 0x5c, 0x9e, 0x5c,
	0xa7, 0xea, 0x5f, 0x5c, 0x92, 0xea, 0x8f, 0x8d, 0x6f, 0x66, 0x59, 0xae, 0x73, 0x17, 0xf9, 0x3f,
	0xfa, 0x84, 0x43, 0x5b, 0xa6, 0xde, 0x39, 0xab, 0x6f, 0x16, 0x45, 0x39, 0xdd, 0xcc, 0x7d, 0xba,
	0x13, 0x4a, 0x4c, 0x55, 0xc2, 0xb9, 0x40, 0xfd, 0xa2, 0x71, 0x4f, 0xd6, 0xe9, 0xb7, 0x32, 0xf5,
	0x69, 0xf6, 0x6f, 0xc5, 0xfe, 0x40, 0xd6, 0xd5, 0x63, 0xb9, 0xfc, 0x37, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x78, 0x51, 0x89, 0x10, 0x18, 0x03, 0x00, 0x00,
}
