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
// source: mpls_te_soft_preemption_global_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_soft_preemption_global_info

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

type MplsTeSoftPreemptionGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) Reset()         { *m = MplsTeSoftPreemptionGlobalInfo_KEYS{} }
func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeSoftPreemptionGlobalInfo_KEYS) ProtoMessage()    {}
func (*MplsTeSoftPreemptionGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db9b368a82987f, []int{0}
}

func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS.Merge(m, src)
}
func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS.Size(m)
}
func (m *MplsTeSoftPreemptionGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo_KEYS proto.InternalMessageInfo

type MplsTeSoftPreemptionGlobalInfo struct {
	IsConfigured                bool     `protobuf:"varint,50,opt,name=is_configured,json=isConfigured,proto3" json:"is_configured,omitempty"`
	IsTimeoutIntervalConfigured bool     `protobuf:"varint,51,opt,name=is_timeout_interval_configured,json=isTimeoutIntervalConfigured,proto3" json:"is_timeout_interval_configured,omitempty"`
	TimeoutInterval             uint32   `protobuf:"varint,52,opt,name=timeout_interval,json=timeoutInterval,proto3" json:"timeout_interval,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MplsTeSoftPreemptionGlobalInfo) Reset()         { *m = MplsTeSoftPreemptionGlobalInfo{} }
func (m *MplsTeSoftPreemptionGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeSoftPreemptionGlobalInfo) ProtoMessage()    {}
func (*MplsTeSoftPreemptionGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9db9b368a82987f, []int{1}
}

func (m *MplsTeSoftPreemptionGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo.Unmarshal(m, b)
}
func (m *MplsTeSoftPreemptionGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeSoftPreemptionGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo.Merge(m, src)
}
func (m *MplsTeSoftPreemptionGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo.Size(m)
}
func (m *MplsTeSoftPreemptionGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeSoftPreemptionGlobalInfo proto.InternalMessageInfo

func (m *MplsTeSoftPreemptionGlobalInfo) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *MplsTeSoftPreemptionGlobalInfo) GetIsTimeoutIntervalConfigured() bool {
	if m != nil {
		return m.IsTimeoutIntervalConfigured
	}
	return false
}

func (m *MplsTeSoftPreemptionGlobalInfo) GetTimeoutInterval() uint32 {
	if m != nil {
		return m.TimeoutInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeSoftPreemptionGlobalInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.soft_preemption_global_info.mpls_te_soft_preemption_global_info_KEYS")
	proto.RegisterType((*MplsTeSoftPreemptionGlobalInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.soft_preemption_global_info.mpls_te_soft_preemption_global_info")
}

func init() {
	proto.RegisterFile("mpls_te_soft_preemption_global_info.proto", fileDescriptor_e9db9b368a82987f)
}

var fileDescriptor_e9db9b368a82987f = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0x85, 0x40,
	0x10, 0x80, 0xf1, 0x12, 0xb1, 0x24, 0x85, 0x27, 0x21, 0x08, 0xd1, 0x8b, 0x76, 0xf0, 0x90, 0xdd,
	0x3b, 0x48, 0x87, 0xe8, 0x66, 0x5d, 0x3a, 0x0d, 0xba, 0xad, 0x32, 0xb0, 0xee, 0x2c, 0x3b, 0x63,
	0xf4, 0xb7, 0xf5, 0xd7, 0x05, 0xf6, 0x03, 0xdf, 0x3b, 0x3c, 0x3c, 0xce, 0xf0, 0x7d, 0xdf, 0xc0,
	0xa8, 0x6a, 0xf6, 0x96, 0x41, 0x0c, 0x30, 0x8d, 0x02, 0x3e, 0x18, 0x33, 0x7b, 0x41, 0x72, 0x30,
	0x59, 0x1a, 0x7a, 0x0b, 0xe8, 0x46, 0xaa, 0x7d, 0x20, 0xa1, 0xe4, 0x41, 0x23, 0x6b, 0x02, 0x24,
	0x86, 0xcf, 0x00, 0x7f, 0x1e, 0x79, 0x13, 0xea, 0x75, 0xb0, 0xba, 0xd7, 0xf5, 0x89, 0x4c, 0x7e,
	0xab, 0xca, 0x1d, 0xd7, 0xe0, 0xf9, 0xf1, 0xed, 0x25, 0xff, 0x8a, 0x54, 0xb1, 0x03, 0x4e, 0x0a,
	0x15, 0x23, 0x83, 0x26, 0x37, 0xe2, 0xb4, 0x04, 0xf3, 0x9e, 0xde, 0x65, 0x51, 0x79, 0xde, 0x5d,
	0x20, 0xb7, 0xff, 0xbb, 0xa4, 0x55, 0x37, 0xc8, 0x20, 0x38, 0x1b, 0x5a, 0x04, 0xd0, 0x89, 0x09,
	0x1f, 0xbd, 0xdd, 0x5a, 0xcd, 0x6a, 0x5d, 0x23, 0xbf, 0xfe, 0x40, 0x4f, 0xbf, 0xcc, 0x26, 0x52,
	0xa9, 0xab, 0xe3, 0x42, 0x7a, 0x9f, 0x45, 0x65, 0xdc, 0x5d, 0xca, 0xa1, 0x34, 0x9c, 0xad, 0x0f,
	0x6b, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x3a, 0xdb, 0x8e, 0x5d, 0x01, 0x00, 0x00,
}
