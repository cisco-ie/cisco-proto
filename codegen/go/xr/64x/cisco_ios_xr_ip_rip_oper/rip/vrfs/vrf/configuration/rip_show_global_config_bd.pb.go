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

package cisco_ios_xr_ip_rip_oper_rip_vrfs_vrf_configuration

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
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

func (m *RipShowGlobalConfigBd_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*RipShowGlobalConfigBd_KEYS)(nil), "cisco_ios_xr_ip_rip_oper.rip.vrfs.vrf.configuration.rip_show_global_config_bd_KEYS")
	proto.RegisterType((*RipShowGlobalConfigBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.vrfs.vrf.configuration.rip_show_global_config_bd")
}

func init() { proto.RegisterFile("rip_show_global_config_bd.proto", fileDescriptor_aa8f8c81e47d6e2c) }

var fileDescriptor_aa8f8c81e47d6e2c = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x5d, 0x6b, 0x14, 0x3d,
	0x14, 0xc7, 0xd9, 0x8b, 0xa7, 0x4f, 0x9b, 0x76, 0xfb, 0x12, 0x45, 0x52, 0x44, 0xbb, 0x56, 0xad,
	0x0b, 0xc2, 0x82, 0xd6, 0xf7, 0xfa, 0x56, 0xad, 0x05, 0xb1, 0x8a, 0xee, 0x56, 0xc1, 0xab, 0x43,
	0x76, 0x26, 0x99, 0x39, 0x38, 0x49, 0xc6, 0x24, 0x33, 0xbb, 0xf8, 0x61, 0xfd, 0x2c, 0x92, 0x64,
	0x76, 0xf5, 0xa6, 0x37, 0x03, 0xf9, 0x9d, 0x5f, 0xce, 0xc9, 0x3f, 0x61, 0xc8, 0x9e, 0xc5, 0x1a,
	0x5c, 0x69, 0x66, 0x50, 0x54, 0x66, 0xca, 0x2b, 0xc8, 0x8c, 0x96, 0x58, 0xc0, 0x34, 0x1f, 0xd5,
	0xd6, 0x78, 0x43, 0x0f, 0x33, 0x74, 0x99, 0x01, 0x34, 0x0e, 0xe6, 0x16, 0xb0, 0x86, 0xb0, 0xc1,
	0xd4, 0xc2, 0x8e, 0x2c, 0xd6, 0xa3, 0xd6, 0x4a, 0x17, 0x3e, 0xa3, 0xb4, 0xad, 0xb1, 0xdc, 0xa3,
	0xd1, 0xfb, 0x47, 0xe4, 0xfa, 0x85, 0x7d, 0xe1, 0xc3, 0xbb, 0xef, 0x13, 0xba, 0x4b, 0x56, 0x5b,
	0x2b, 0x41, 0x73, 0x25, 0x58, 0x6f, 0xd0, 0x1b, 0xae, 0x8d, 0xff, 0x6f, 0xad, 0xfc, 0xc4, 0x95,
	0xd8, 0xff, 0xfd, 0x1f, 0xd9, 0xbd, 0x70, 0x37, 0xbd, 0x42, 0x56, 0x78, 0xe6, 0xb1, 0x15, 0xec,
	0xfe, 0xa0, 0x37, 0x5c, 0x1d, 0x77, 0x2b, 0x7a, 0x40, 0xb6, 0x5a, 0x0b, 0x12, 0x9d, 0xc8, 0xc1,
	0x99, 0xec, 0x87, 0xf0, 0xec, 0x30, 0x0a, 0xfd, 0xd6, 0x9e, 0x06, 0x3a, 0x89, 0x90, 0xee, 0x91,
	0xf5, 0xd0, 0xbc, 0x15, 0xd6, 0xa1, 0xd1, 0xec, 0xc1, 0xa0, 0x37, 0xdc, 0x19, 0x13, 0x8b, 0xf5,
	0xb7, 0x44, 0xe8, 0x6d, 0xb2, 0x99, 0x0b, 0xc9, 0x9b, 0xca, 0x83, 0x12, 0xde, 0x62, 0xc6, 0x1e,
	0x0e, 0x7a, 0xc3, 0xfe, 0xb8, 0xdf, 0xd1, 0x8f, 0x11, 0xd2, 0x9b, 0xa4, 0xaf, 0xf8, 0x1c, 0x55,
	0xa3, 0xa0, 0xe6, 0xbe, 0x74, 0xec, 0x51, 0xb4, 0x36, 0x3a, 0xf8, 0x39, 0xb0, 0xd0, 0x8b, 0x37,
	0xde, 0x80, 0x6b, 0x94, 0xe2, 0x16, 0x7f, 0x09, 0xf6, 0x38, 0x9d, 0x29, 0xd0, 0xc9, 0x02, 0xd2,
	0xbb, 0x64, 0x47, 0x35, 0x95, 0xc7, 0x8c, 0x3b, 0x0f, 0x3c, 0xcf, 0xad, 0x70, 0x8e, 0x3d, 0x89,
	0xe6, 0xf6, 0xb2, 0x70, 0x9c, 0x38, 0xbd, 0x43, 0xb6, 0x64, 0xc5, 0x5d, 0x09, 0xbe, 0xb4, 0xc2,
	0x95, 0xa6, 0xca, 0xd9, 0xd3, 0x38, 0x7a, 0x33, 0xe2, 0xf3, 0x05, 0xa5, 0xb7, 0xc8, 0x26, 0xea,
	0xba, 0xf1, 0xf0, 0x13, 0x2a, 0xa1, 0x0b, 0x5f, 0xb2, 0x67, 0xe9, 0x88, 0x91, 0x7e, 0x39, 0x8b,
	0x2c, 0xe4, 0xf0, 0x16, 0x8b, 0x42, 0x58, 0x91, 0x87, 0xb7, 0x65, 0x47, 0x71, 0xee, 0xc6, 0x12,
	0x8e, 0xb1, 0xa6, 0xf7, 0xc8, 0xe5, 0x96, 0x57, 0x98, 0xc7, 0xd7, 0x05, 0xd4, 0x39, 0x66, 0xdc,
	0x1b, 0xcb, 0x9e, 0x47, 0xf7, 0xd2, 0xdf, 0xda, 0xfb, 0x45, 0x89, 0xde, 0x20, 0x1b, 0x4d, 0x9d,
	0x73, 0x2f, 0xc0, 0xa3, 0x12, 0x96, 0xbd, 0x88, 0xb3, 0xd7, 0x13, 0x3b, 0x0f, 0x88, 0x0e, 0xc9,
	0xb6, 0x16, 0x73, 0x0f, 0xff, 0x78, 0xec, 0x65, 0x8a, 0x12, 0xf8, 0xd7, 0xa5, 0x1a, 0x0e, 0x89,
	0x3a, 0x4e, 0xe9, 0xba, 0xbd, 0x5a, 0x24, 0x89, 0x30, 0xb5, 0x3b, 0x20, 0x5b, 0x21, 0x37, 0xe4,
	0x66, 0xa6, 0x3b, 0xed, 0x75, 0x7a, 0xb9, 0x80, 0x4f, 0xcc, 0x4c, 0x27, 0x6f, 0x8f, 0xac, 0xcb,
	0xaa, 0x09, 0x17, 0x18, 0x9d, 0xe3, 0xe8, 0x90, 0x88, 0x92, 0x70, 0x95, 0xac, 0x19, 0xa3, 0x40,
	0x56, 0xbc, 0x70, 0xec, 0x4d, 0x2c, 0xaf, 0x1a, 0xa3, 0x4e, 0xc3, 0x9a, 0x5e, 0x23, 0x44, 0x3b,
	0x09, 0xce, 0x73, 0xdf, 0x38, 0xf6, 0x36, 0x5e, 0xc0, 0x9a, 0x76, 0x72, 0x12, 0x01, 0xdd, 0x27,
	0xfd, 0x50, 0xae, 0x50, 0x76, 0x81, 0x4e, 0x52, 0x6e, 0xed, 0xe4, 0x19, 0xca, 0x98, 0x66, 0xba,
	0x12, 0xff, 0xac, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x5c, 0x51, 0x0c, 0x7c, 0x03,
	0x00, 0x00,
}