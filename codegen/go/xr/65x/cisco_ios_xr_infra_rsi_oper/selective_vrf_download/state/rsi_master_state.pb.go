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
// source: rsi_master_state.proto

package cisco_ios_xr_infra_rsi_oper_selective_vrf_download_state

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

type RsiMasterState_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiMasterState_KEYS) Reset()         { *m = RsiMasterState_KEYS{} }
func (m *RsiMasterState_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsiMasterState_KEYS) ProtoMessage()    {}
func (*RsiMasterState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f411c231d8bb09ca, []int{0}
}

func (m *RsiMasterState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiMasterState_KEYS.Unmarshal(m, b)
}
func (m *RsiMasterState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiMasterState_KEYS.Marshal(b, m, deterministic)
}
func (m *RsiMasterState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiMasterState_KEYS.Merge(m, src)
}
func (m *RsiMasterState_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsiMasterState_KEYS.Size(m)
}
func (m *RsiMasterState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiMasterState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsiMasterState_KEYS proto.InternalMessageInfo

type RsiMasterState struct {
	IsSvdEnabled         bool     `protobuf:"varint,50,opt,name=is_svd_enabled,json=isSvdEnabled,proto3" json:"is_svd_enabled,omitempty"`
	IsSvdEnabledCfg      bool     `protobuf:"varint,51,opt,name=is_svd_enabled_cfg,json=isSvdEnabledCfg,proto3" json:"is_svd_enabled_cfg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiMasterState) Reset()         { *m = RsiMasterState{} }
func (m *RsiMasterState) String() string { return proto.CompactTextString(m) }
func (*RsiMasterState) ProtoMessage()    {}
func (*RsiMasterState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f411c231d8bb09ca, []int{1}
}

func (m *RsiMasterState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiMasterState.Unmarshal(m, b)
}
func (m *RsiMasterState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiMasterState.Marshal(b, m, deterministic)
}
func (m *RsiMasterState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiMasterState.Merge(m, src)
}
func (m *RsiMasterState) XXX_Size() int {
	return xxx_messageInfo_RsiMasterState.Size(m)
}
func (m *RsiMasterState) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiMasterState.DiscardUnknown(m)
}

var xxx_messageInfo_RsiMasterState proto.InternalMessageInfo

func (m *RsiMasterState) GetIsSvdEnabled() bool {
	if m != nil {
		return m.IsSvdEnabled
	}
	return false
}

func (m *RsiMasterState) GetIsSvdEnabledCfg() bool {
	if m != nil {
		return m.IsSvdEnabledCfg
	}
	return false
}

func init() {
	proto.RegisterType((*RsiMasterState_KEYS)(nil), "cisco_ios_xr_infra_rsi_oper.selective_vrf_download.state.rsi_master_state_KEYS")
	proto.RegisterType((*RsiMasterState)(nil), "cisco_ios_xr_infra_rsi_oper.selective_vrf_download.state.rsi_master_state")
}

func init() { proto.RegisterFile("rsi_master_state.proto", fileDescriptor_f411c231d8bb09ca) }

var fileDescriptor_f411c231d8bb09ca = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xcf, 0xb1, 0xab, 0xc2, 0x30,
	0x10, 0x06, 0x70, 0xde, 0xf2, 0x78, 0x84, 0x87, 0x4a, 0x40, 0xed, 0x28, 0xc5, 0x41, 0x10, 0x3a,
	0xd8, 0xc5, 0x5d, 0x3a, 0xb9, 0xd9, 0xc9, 0xe9, 0x48, 0x93, 0x4b, 0x09, 0xd4, 0x5e, 0xb9, 0x0b,
	0xd1, 0x3f, 0x5f, 0xac, 0x8b, 0x76, 0xbd, 0xfb, 0x7d, 0x1f, 0x7c, 0x6a, 0xc5, 0x12, 0xe0, 0x66,
	0x24, 0x22, 0x83, 0x44, 0x13, 0xb1, 0x18, 0x98, 0x22, 0xe9, 0xa3, 0x0d, 0x62, 0x09, 0x02, 0x09,
	0x3c, 0x18, 0x42, 0xef, 0xd9, 0xc0, 0x8b, 0xd2, 0x80, 0x5c, 0x08, 0x76, 0x68, 0x63, 0x48, 0x08,
	0x89, 0x3d, 0x38, 0xba, 0xf7, 0x1d, 0x19, 0x57, 0x8c, 0xf9, 0x7c, 0xad, 0x96, 0xd3, 0x4e, 0x38,
	0x57, 0xd7, 0x3a, 0x47, 0xb5, 0x98, 0x3e, 0xf4, 0x56, 0xcd, 0x82, 0x80, 0x24, 0x07, 0xd8, 0x9b,
	0xa6, 0x43, 0x97, 0x1d, 0x36, 0x3f, 0xbb, 0xbf, 0xcb, 0x7f, 0x90, 0x3a, 0xb9, 0xea, 0x7d, 0xd3,
	0x7b, 0xa5, 0xbf, 0x15, 0x58, 0xdf, 0x66, 0xe5, 0x28, 0xe7, 0x9f, 0xf2, 0xe4, 0xdb, 0xe6, 0x77,
	0x1c, 0x50, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x06, 0x28, 0xc7, 0xda, 0x00, 0x00, 0x00,
}