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
// source: tty_general_stats.proto

package cisco_ios_xr_tty_server_oper_tty_vty_lines_vty_line_vty_statistics_general_statistics

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

type TtyGeneralStats_KEYS struct {
	LineNumber           uint32   `protobuf:"varint,1,opt,name=line_number,json=lineNumber,proto3" json:"line_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TtyGeneralStats_KEYS) Reset()         { *m = TtyGeneralStats_KEYS{} }
func (m *TtyGeneralStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*TtyGeneralStats_KEYS) ProtoMessage()    {}
func (*TtyGeneralStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ad9e5d4e858fd, []int{0}
}

func (m *TtyGeneralStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TtyGeneralStats_KEYS.Unmarshal(m, b)
}
func (m *TtyGeneralStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TtyGeneralStats_KEYS.Marshal(b, m, deterministic)
}
func (m *TtyGeneralStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TtyGeneralStats_KEYS.Merge(m, src)
}
func (m *TtyGeneralStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_TtyGeneralStats_KEYS.Size(m)
}
func (m *TtyGeneralStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TtyGeneralStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TtyGeneralStats_KEYS proto.InternalMessageInfo

func (m *TtyGeneralStats_KEYS) GetLineNumber() uint32 {
	if m != nil {
		return m.LineNumber
	}
	return 0
}

type TtyGeneralStats struct {
	TerminalLength            uint32   `protobuf:"varint,50,opt,name=terminal_length,json=terminalLength,proto3" json:"terminal_length,omitempty"`
	TerminalWidth             uint32   `protobuf:"varint,51,opt,name=terminal_width,json=terminalWidth,proto3" json:"terminal_width,omitempty"`
	AsyncInterface            bool     `protobuf:"varint,52,opt,name=async_interface,json=asyncInterface,proto3" json:"async_interface,omitempty"`
	FlowControlStartCharacter string   `protobuf:"bytes,53,opt,name=flow_control_start_character,json=flowControlStartCharacter,proto3" json:"flow_control_start_character,omitempty"`
	FlowControlStopCharacter  string   `protobuf:"bytes,54,opt,name=flow_control_stop_character,json=flowControlStopCharacter,proto3" json:"flow_control_stop_character,omitempty"`
	DomainLookupEnabled       bool     `protobuf:"varint,55,opt,name=domain_lookup_enabled,json=domainLookupEnabled,proto3" json:"domain_lookup_enabled,omitempty"`
	MotdBannerEnabled         bool     `protobuf:"varint,56,opt,name=motd_banner_enabled,json=motdBannerEnabled,proto3" json:"motd_banner_enabled,omitempty"`
	PrivateFlag               bool     `protobuf:"varint,57,opt,name=private_flag,json=privateFlag,proto3" json:"private_flag,omitempty"`
	TerminalType              string   `protobuf:"bytes,58,opt,name=terminal_type,json=terminalType,proto3" json:"terminal_type,omitempty"`
	AbsoluteTimeout           uint32   `protobuf:"varint,59,opt,name=absolute_timeout,json=absoluteTimeout,proto3" json:"absolute_timeout,omitempty"`
	IdleTime                  uint32   `protobuf:"varint,60,opt,name=idle_time,json=idleTime,proto3" json:"idle_time,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *TtyGeneralStats) Reset()         { *m = TtyGeneralStats{} }
func (m *TtyGeneralStats) String() string { return proto.CompactTextString(m) }
func (*TtyGeneralStats) ProtoMessage()    {}
func (*TtyGeneralStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_f34ad9e5d4e858fd, []int{1}
}

func (m *TtyGeneralStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TtyGeneralStats.Unmarshal(m, b)
}
func (m *TtyGeneralStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TtyGeneralStats.Marshal(b, m, deterministic)
}
func (m *TtyGeneralStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TtyGeneralStats.Merge(m, src)
}
func (m *TtyGeneralStats) XXX_Size() int {
	return xxx_messageInfo_TtyGeneralStats.Size(m)
}
func (m *TtyGeneralStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TtyGeneralStats.DiscardUnknown(m)
}

var xxx_messageInfo_TtyGeneralStats proto.InternalMessageInfo

func (m *TtyGeneralStats) GetTerminalLength() uint32 {
	if m != nil {
		return m.TerminalLength
	}
	return 0
}

func (m *TtyGeneralStats) GetTerminalWidth() uint32 {
	if m != nil {
		return m.TerminalWidth
	}
	return 0
}

func (m *TtyGeneralStats) GetAsyncInterface() bool {
	if m != nil {
		return m.AsyncInterface
	}
	return false
}

func (m *TtyGeneralStats) GetFlowControlStartCharacter() string {
	if m != nil {
		return m.FlowControlStartCharacter
	}
	return ""
}

func (m *TtyGeneralStats) GetFlowControlStopCharacter() string {
	if m != nil {
		return m.FlowControlStopCharacter
	}
	return ""
}

func (m *TtyGeneralStats) GetDomainLookupEnabled() bool {
	if m != nil {
		return m.DomainLookupEnabled
	}
	return false
}

func (m *TtyGeneralStats) GetMotdBannerEnabled() bool {
	if m != nil {
		return m.MotdBannerEnabled
	}
	return false
}

func (m *TtyGeneralStats) GetPrivateFlag() bool {
	if m != nil {
		return m.PrivateFlag
	}
	return false
}

func (m *TtyGeneralStats) GetTerminalType() string {
	if m != nil {
		return m.TerminalType
	}
	return ""
}

func (m *TtyGeneralStats) GetAbsoluteTimeout() uint32 {
	if m != nil {
		return m.AbsoluteTimeout
	}
	return 0
}

func (m *TtyGeneralStats) GetIdleTime() uint32 {
	if m != nil {
		return m.IdleTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TtyGeneralStats_KEYS)(nil), "cisco_ios_xr_tty_server_oper.tty.vty_lines.vty_line.vty_statistics.general_statistics.tty_general_stats_KEYS")
	proto.RegisterType((*TtyGeneralStats)(nil), "cisco_ios_xr_tty_server_oper.tty.vty_lines.vty_line.vty_statistics.general_statistics.tty_general_stats")
}

func init() { proto.RegisterFile("tty_general_stats.proto", fileDescriptor_f34ad9e5d4e858fd) }

var fileDescriptor_f34ad9e5d4e858fd = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdb, 0x6b, 0x14, 0x31,
	0x14, 0xc6, 0x59, 0x10, 0x69, 0xd3, 0x9b, 0x4d, 0x51, 0x23, 0x15, 0x5c, 0x2b, 0xe2, 0xfa, 0x32,
	0x0f, 0xad, 0xb7, 0x7a, 0x41, 0xb0, 0x54, 0x10, 0x8b, 0x0f, 0xdb, 0x8a, 0xf8, 0x74, 0xc8, 0xcc,
	0x9c, 0xdd, 0x0d, 0x66, 0x92, 0x90, 0x9c, 0xdd, 0x3a, 0x6f, 0xfe, 0xe9, 0x92, 0x33, 0xce, 0x58,
	0xed, 0x5b, 0xf8, 0x7d, 0xbf, 0xef, 0x9b, 0x10, 0x46, 0xdc, 0x25, 0x6a, 0x61, 0x8e, 0x0e, 0xa3,
	0xb6, 0x90, 0x48, 0x53, 0x2a, 0x42, 0xf4, 0xe4, 0xe5, 0xd7, 0xca, 0xa4, 0xca, 0x83, 0xf1, 0x09,
	0x7e, 0x46, 0xc8, 0x56, 0xc2, 0xb8, 0xc2, 0x08, 0x3e, 0x60, 0x2c, 0x88, 0xda, 0x62, 0x45, 0x2d,
	0x58, 0xe3, 0x30, 0x0d, 0x27, 0x3e, 0xe4, 0x11, 0x93, 0xc8, 0x54, 0xa9, 0xb8, 0xba, 0xdb, 0xa1,
	0x83, 0x63, 0x71, 0xe7, 0xda, 0x17, 0xe1, 0xf3, 0xe9, 0xf7, 0x73, 0xf9, 0x40, 0x6c, 0xe4, 0x0d,
	0x70, 0xcb, 0xa6, 0xc4, 0xa8, 0x46, 0xe3, 0xd1, 0x64, 0x6b, 0x2a, 0x32, 0xfa, 0xc2, 0xe4, 0xe0,
	0xd7, 0x0d, 0xb1, 0x7b, 0xad, 0x2b, 0x9f, 0x88, 0x1d, 0xc2, 0xd8, 0x18, 0xa7, 0x2d, 0x58, 0x74,
	0x73, 0x5a, 0xa8, 0x43, 0xae, 0x6e, 0xf7, 0xf8, 0x8c, 0xa9, 0x7c, 0x2c, 0x06, 0x02, 0x97, 0xa6,
	0xa6, 0x85, 0x3a, 0x62, 0x6f, 0xab, 0xa7, 0xdf, 0x32, 0xcc, 0x7b, 0x3a, 0xb5, 0xae, 0x02, 0xe3,
	0x08, 0xe3, 0x4c, 0x57, 0xa8, 0x9e, 0x8d, 0x47, 0x93, 0xb5, 0xe9, 0x36, 0xe3, 0x4f, 0x3d, 0x95,
	0xef, 0xc5, 0xfd, 0x99, 0xf5, 0x97, 0x50, 0x79, 0x47, 0xd1, 0xf3, 0x75, 0x22, 0x41, 0xb5, 0xd0,
	0x51, 0x57, 0x84, 0x51, 0x3d, 0x1f, 0x8f, 0x26, 0xeb, 0xd3, 0x7b, 0xd9, 0x39, 0xe9, 0x94, 0xf3,
	0x6c, 0x9c, 0xf4, 0x82, 0x7c, 0x27, 0xf6, 0xff, 0x1b, 0xf0, 0xe1, 0x4a, 0xff, 0x05, 0xf7, 0xd5,
	0x3f, 0x7d, 0x1f, 0xfe, 0xd6, 0x0f, 0xc5, 0xed, 0xda, 0x37, 0xda, 0x38, 0xb0, 0xde, 0xff, 0x58,
	0x06, 0x40, 0xa7, 0x4b, 0x8b, 0xb5, 0x7a, 0xc9, 0xd7, 0xdd, 0xeb, 0xc2, 0x33, 0xce, 0x4e, 0xbb,
	0x48, 0x16, 0x62, 0xaf, 0xf1, 0x54, 0x43, 0xa9, 0x9d, 0xc3, 0x38, 0x34, 0x5e, 0x71, 0x63, 0x37,
	0x47, 0x1f, 0x38, 0xe9, 0xfd, 0x87, 0x62, 0x33, 0x44, 0xb3, 0xd2, 0x84, 0x30, 0xb3, 0x7a, 0xae,
	0x8e, 0x59, 0xdc, 0xf8, 0xc3, 0x3e, 0x5a, 0x3d, 0x97, 0x8f, 0xc4, 0xf0, 0x80, 0x40, 0x6d, 0x40,
	0xf5, 0x9a, 0xef, 0xbd, 0xd9, 0xc3, 0x8b, 0x36, 0xa0, 0x7c, 0x2a, 0x6e, 0xe9, 0x32, 0x79, 0xbb,
	0x24, 0x04, 0x32, 0x0d, 0xfa, 0x25, 0xa9, 0x37, 0xfc, 0xfa, 0x3b, 0x3d, 0xbf, 0xe8, 0xb0, 0xdc,
	0x17, 0xeb, 0xa6, 0xb6, 0x9d, 0xa6, 0xde, 0xb2, 0xb3, 0x96, 0x41, 0xce, 0xcb, 0x9b, 0xfc, 0x6f,
	0x1e, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xae, 0xfe, 0x55, 0xb6, 0x02, 0x00, 0x00,
}
