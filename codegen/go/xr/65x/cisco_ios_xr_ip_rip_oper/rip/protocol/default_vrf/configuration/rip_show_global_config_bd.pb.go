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
// source: rip_show_global_config_bd.proto

package cisco_ios_xr_ip_rip_oper_rip_protocol_default_vrf_configuration

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

type RipShowGlobalConfigBd_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RipShowGlobalConfigBd_KEYS) Reset()         { *m = RipShowGlobalConfigBd_KEYS{} }
func (m *RipShowGlobalConfigBd_KEYS) String() string { return proto.CompactTextString(m) }
func (*RipShowGlobalConfigBd_KEYS) ProtoMessage()    {}
func (*RipShowGlobalConfigBd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa8f8c81e47d6e2c, []int{0}
}

func (m *RipShowGlobalConfigBd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowGlobalConfigBd_KEYS.Unmarshal(m, b)
}
func (m *RipShowGlobalConfigBd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowGlobalConfigBd_KEYS.Marshal(b, m, deterministic)
}
func (m *RipShowGlobalConfigBd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowGlobalConfigBd_KEYS.Merge(m, src)
}
func (m *RipShowGlobalConfigBd_KEYS) XXX_Size() int {
	return xxx_messageInfo_RipShowGlobalConfigBd_KEYS.Size(m)
}
func (m *RipShowGlobalConfigBd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowGlobalConfigBd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowGlobalConfigBd_KEYS proto.InternalMessageInfo

type RipShowGlobalConfigBd struct {
	Active               bool     `protobuf:"varint,50,opt,name=active,proto3" json:"active,omitempty"`
	VrFisedSocket        bool     `protobuf:"varint,51,opt,name=vr_fised_socket,json=vrFisedSocket,proto3" json:"vr_fised_socket,omitempty"`
	RipVersion           int32    `protobuf:"zigzag32,52,opt,name=rip_version,json=ripVersion,proto3" json:"rip_version,omitempty"`
	DefaultMetric        uint32   `protobuf:"varint,53,opt,name=default_metric,json=defaultMetric,proto3" json:"default_metric,omitempty"`
	MaximumPaths         uint32   `protobuf:"varint,54,opt,name=maximum_paths,json=maximumPaths,proto3" json:"maximum_paths,omitempty"`
	AutoSummarize        bool     `protobuf:"varint,55,opt,name=auto_summarize,json=autoSummarize,proto3" json:"auto_summarize,omitempty"`
	MulticastAddress     bool     `protobuf:"varint,56,opt,name=multicast_address,json=multicastAddress,proto3" json:"multicast_address,omitempty"`
	FlashThreshold       uint32   `protobuf:"varint,57,opt,name=flash_threshold,json=flashThreshold,proto3" json:"flash_threshold,omitempty"`
	InputQLength         uint32   `protobuf:"varint,58,opt,name=input_q_length,json=inputQLength,proto3" json:"input_q_length,omitempty"`
	TriggeredRip         bool     `protobuf:"varint,59,opt,name=triggered_rip,json=triggeredRip,proto3" json:"triggered_rip,omitempty"`
	ValidationIndicator  bool     `protobuf:"varint,60,opt,name=validation_indicator,json=validationIndicator,proto3" json:"validation_indicator,omitempty"`
	UpdateTimer          uint32   `protobuf:"varint,61,opt,name=update_timer,json=updateTimer,proto3" json:"update_timer,omitempty"`
	NextUpdateTime       uint32   `protobuf:"varint,62,opt,name=next_update_time,json=nextUpdateTime,proto3" json:"next_update_time,omitempty"`
	InvalidTimer         uint32   `protobuf:"varint,63,opt,name=invalid_timer,json=invalidTimer,proto3" json:"invalid_timer,omitempty"`
	HoldDownTimer        uint32   `protobuf:"varint,64,opt,name=hold_down_timer,json=holdDownTimer,proto3" json:"hold_down_timer,omitempty"`
	FlushTimer           uint32   `protobuf:"varint,65,opt,name=flush_timer,json=flushTimer,proto3" json:"flush_timer,omitempty"`
	OomFlags             uint32   `protobuf:"varint,66,opt,name=oom_flags,json=oomFlags,proto3" json:"oom_flags,omitempty"`
	NsfStatus            bool     `protobuf:"varint,67,opt,name=nsf_status,json=nsfStatus,proto3" json:"nsf_status,omitempty"`
	NsfLifeTime          uint32   `protobuf:"varint,68,opt,name=nsf_life_time,json=nsfLifeTime,proto3" json:"nsf_life_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RipShowGlobalConfigBd) Reset()         { *m = RipShowGlobalConfigBd{} }
func (m *RipShowGlobalConfigBd) String() string { return proto.CompactTextString(m) }
func (*RipShowGlobalConfigBd) ProtoMessage()    {}
func (*RipShowGlobalConfigBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa8f8c81e47d6e2c, []int{1}
}

func (m *RipShowGlobalConfigBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowGlobalConfigBd.Unmarshal(m, b)
}
func (m *RipShowGlobalConfigBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowGlobalConfigBd.Marshal(b, m, deterministic)
}
func (m *RipShowGlobalConfigBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowGlobalConfigBd.Merge(m, src)
}
func (m *RipShowGlobalConfigBd) XXX_Size() int {
	return xxx_messageInfo_RipShowGlobalConfigBd.Size(m)
}
func (m *RipShowGlobalConfigBd) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowGlobalConfigBd.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowGlobalConfigBd proto.InternalMessageInfo

func (m *RipShowGlobalConfigBd) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetVrFisedSocket() bool {
	if m != nil {
		return m.VrFisedSocket
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetRipVersion() int32 {
	if m != nil {
		return m.RipVersion
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetDefaultMetric() uint32 {
	if m != nil {
		return m.DefaultMetric
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetMaximumPaths() uint32 {
	if m != nil {
		return m.MaximumPaths
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetAutoSummarize() bool {
	if m != nil {
		return m.AutoSummarize
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetMulticastAddress() bool {
	if m != nil {
		return m.MulticastAddress
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetFlashThreshold() uint32 {
	if m != nil {
		return m.FlashThreshold
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetInputQLength() uint32 {
	if m != nil {
		return m.InputQLength
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetTriggeredRip() bool {
	if m != nil {
		return m.TriggeredRip
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetValidationIndicator() bool {
	if m != nil {
		return m.ValidationIndicator
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetUpdateTimer() uint32 {
	if m != nil {
		return m.UpdateTimer
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetNextUpdateTime() uint32 {
	if m != nil {
		return m.NextUpdateTime
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetInvalidTimer() uint32 {
	if m != nil {
		return m.InvalidTimer
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetHoldDownTimer() uint32 {
	if m != nil {
		return m.HoldDownTimer
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetFlushTimer() uint32 {
	if m != nil {
		return m.FlushTimer
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetOomFlags() uint32 {
	if m != nil {
		return m.OomFlags
	}
	return 0
}

func (m *RipShowGlobalConfigBd) GetNsfStatus() bool {
	if m != nil {
		return m.NsfStatus
	}
	return false
}

func (m *RipShowGlobalConfigBd) GetNsfLifeTime() uint32 {
	if m != nil {
		return m.NsfLifeTime
	}
	return 0
}

func init() {
	proto.RegisterType((*RipShowGlobalConfigBd_KEYS)(nil), "cisco_ios_xr_ip_rip_oper.rip.protocol.default_vrf.configuration.rip_show_global_config_bd_KEYS")
	proto.RegisterType((*RipShowGlobalConfigBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.protocol.default_vrf.configuration.rip_show_global_config_bd")
}

func init() { proto.RegisterFile("rip_show_global_config_bd.proto", fileDescriptor_aa8f8c81e47d6e2c) }

var fileDescriptor_aa8f8c81e47d6e2c = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5b, 0x6b, 0x14, 0x31,
	0x14, 0xc7, 0xe9, 0x83, 0xa5, 0x4d, 0xbb, 0xbd, 0x44, 0x91, 0x88, 0xe8, 0xae, 0xab, 0xd6, 0x05,
	0x61, 0x41, 0xeb, 0xfd, 0x56, 0xab, 0xb5, 0x20, 0x56, 0xd0, 0xdd, 0x2a, 0xf8, 0x74, 0xc8, 0xce,
	0x24, 0x33, 0x07, 0x33, 0xc9, 0x98, 0x64, 0x66, 0x17, 0x3f, 0xac, 0x9f, 0x45, 0x92, 0xcc, 0xae,
	0xbe, 0xf4, 0x31, 0xbf, 0xf3, 0x9b, 0x73, 0xce, 0x3f, 0x61, 0x48, 0xdf, 0x62, 0x0d, 0xae, 0x34,
	0x73, 0x28, 0x94, 0x99, 0x71, 0x05, 0x99, 0xd1, 0x12, 0x0b, 0x98, 0xe5, 0xe3, 0xda, 0x1a, 0x6f,
	0xe8, 0x51, 0x86, 0x2e, 0x33, 0x80, 0xc6, 0xc1, 0xc2, 0x02, 0xd6, 0x10, 0x3e, 0x30, 0xb5, 0xb0,
	0x63, 0x8b, 0x75, 0x72, 0x32, 0xa3, 0xc6, 0xb9, 0x90, 0xbc, 0x51, 0x1e, 0x5a, 0x2b, 0xc7, 0xa9,
	0x45, 0x63, 0xb9, 0x47, 0xa3, 0x87, 0x03, 0x72, 0xf3, 0xc2, 0x19, 0xf0, 0xe9, 0xc3, 0x8f, 0xe9,
	0xf0, 0xcf, 0x25, 0x72, 0xed, 0x42, 0x85, 0x5e, 0x25, 0xeb, 0x3c, 0xf3, 0xd8, 0x0a, 0xf6, 0x70,
	0xb0, 0x36, 0xda, 0x98, 0x74, 0x27, 0x7a, 0x40, 0x76, 0x5b, 0x0b, 0x12, 0x9d, 0xc8, 0xc1, 0x99,
	0xec, 0xa7, 0xf0, 0xec, 0x30, 0x0a, 0xbd, 0xd6, 0x9e, 0x06, 0x3a, 0x8d, 0x90, 0xf6, 0xc9, 0x56,
	0x68, 0xde, 0x0a, 0xeb, 0xd0, 0x68, 0xf6, 0x68, 0xb0, 0x36, 0xda, 0x9f, 0x10, 0x8b, 0xf5, 0xf7,
	0x44, 0xe8, 0x5d, 0xb2, 0xb3, 0xdc, 0xbe, 0x12, 0xde, 0x62, 0xc6, 0x1e, 0x0f, 0xd6, 0x46, 0xbd,
	0x49, 0xaf, 0xa3, 0x9f, 0x23, 0xa4, 0xb7, 0x49, 0xaf, 0xe2, 0x0b, 0xac, 0x9a, 0x0a, 0x6a, 0xee,
	0x4b, 0xc7, 0x9e, 0x44, 0x6b, 0xbb, 0x83, 0x5f, 0x02, 0x0b, 0xbd, 0x78, 0xe3, 0x0d, 0xb8, 0xa6,
	0xaa, 0xb8, 0xc5, 0xdf, 0x82, 0x3d, 0x4d, 0x3b, 0x05, 0x3a, 0x5d, 0x42, 0x7a, 0x9f, 0xec, 0x57,
	0x8d, 0xf2, 0x98, 0x71, 0xe7, 0x81, 0xe7, 0xb9, 0x15, 0xce, 0xb1, 0x67, 0xd1, 0xdc, 0x5b, 0x15,
	0x8e, 0x13, 0xa7, 0xf7, 0xc8, 0xae, 0x54, 0xdc, 0x95, 0xe0, 0x4b, 0x2b, 0x5c, 0x69, 0x54, 0xce,
	0x9e, 0xc7, 0xd1, 0x3b, 0x11, 0x9f, 0x2f, 0x29, 0xbd, 0x43, 0x76, 0x50, 0xd7, 0x8d, 0x87, 0x5f,
	0xa0, 0x84, 0x2e, 0x7c, 0xc9, 0x5e, 0xa4, 0x15, 0x23, 0xfd, 0x7a, 0x16, 0x59, 0xc8, 0xe1, 0x2d,
	0x16, 0x85, 0xb0, 0x22, 0x0f, 0x8f, 0xc9, 0x5e, 0xc6, 0xb9, 0xdb, 0x2b, 0x38, 0xc1, 0x9a, 0x3e,
	0x20, 0x57, 0x5a, 0xae, 0x30, 0x8f, 0x4f, 0x08, 0xa8, 0x73, 0xcc, 0xb8, 0x37, 0x96, 0xbd, 0x8a,
	0xee, 0xe5, 0x7f, 0xb5, 0x8f, 0xcb, 0x12, 0xbd, 0x45, 0xb6, 0x9b, 0x3a, 0xe7, 0x5e, 0x80, 0xc7,
	0x4a, 0x58, 0xf6, 0x3a, 0xce, 0xde, 0x4a, 0xec, 0x3c, 0x20, 0x3a, 0x22, 0x7b, 0x5a, 0x2c, 0x3c,
	0xfc, 0xe7, 0xb1, 0x37, 0x29, 0x4a, 0xe0, 0xdf, 0x56, 0x6a, 0x58, 0x12, 0x75, 0x9c, 0xd2, 0x75,
	0x3b, 0x5a, 0x26, 0x89, 0x30, 0xb5, 0x3b, 0x20, 0xbb, 0x21, 0x37, 0xe4, 0x66, 0xae, 0x3b, 0xed,
	0x6d, 0x7a, 0xb9, 0x80, 0x4f, 0xcc, 0x5c, 0x27, 0xaf, 0x4f, 0xb6, 0xa4, 0x6a, 0xc2, 0x05, 0x46,
	0xe7, 0x38, 0x3a, 0x24, 0xa2, 0x24, 0x5c, 0x27, 0x9b, 0xc6, 0x54, 0x20, 0x15, 0x2f, 0x1c, 0x7b,
	0x17, 0xcb, 0x1b, 0xc6, 0x54, 0xa7, 0xe1, 0x4c, 0x6f, 0x10, 0xa2, 0x9d, 0x04, 0xe7, 0xb9, 0x6f,
	0x1c, 0x7b, 0x1f, 0x2f, 0x60, 0x53, 0x3b, 0x39, 0x8d, 0x80, 0x0e, 0x49, 0x2f, 0x94, 0x15, 0xca,
	0x2e, 0xd0, 0x49, 0xca, 0xad, 0x9d, 0x3c, 0x43, 0x19, 0xd3, 0xcc, 0xd6, 0xe3, 0x6f, 0x72, 0xf8,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xf0, 0xb0, 0x72, 0x6d, 0x03, 0x00, 0x00,
}
