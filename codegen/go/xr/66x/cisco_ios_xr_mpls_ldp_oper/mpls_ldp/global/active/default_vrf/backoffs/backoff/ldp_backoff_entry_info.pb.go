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
// source: ldp_backoff_entry_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_backoffs_backoff

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

type LdpBackoffEntryInfo_KEYS struct {
	LsrId                string   `protobuf:"bytes,1,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,2,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffEntryInfo_KEYS) Reset()         { *m = LdpBackoffEntryInfo_KEYS{} }
func (m *LdpBackoffEntryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffEntryInfo_KEYS) ProtoMessage()    {}
func (*LdpBackoffEntryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b1b2b369a41567, []int{0}
}

func (m *LdpBackoffEntryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Merge(m, src)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Size(m)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffEntryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffEntryInfo_KEYS proto.InternalMessageInfo

func (m *LdpBackoffEntryInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpBackoffEntryInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

type LdpBackoffEntryInfo struct {
	BackoffSeconds       uint32   `protobuf:"varint,50,opt,name=backoff_seconds,json=backoffSeconds,proto3" json:"backoff_seconds,omitempty"`
	WaitingSeconds       uint32   `protobuf:"varint,51,opt,name=waiting_seconds,json=waitingSeconds,proto3" json:"waiting_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffEntryInfo) Reset()         { *m = LdpBackoffEntryInfo{} }
func (m *LdpBackoffEntryInfo) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffEntryInfo) ProtoMessage()    {}
func (*LdpBackoffEntryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b1b2b369a41567, []int{1}
}

func (m *LdpBackoffEntryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffEntryInfo.Unmarshal(m, b)
}
func (m *LdpBackoffEntryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffEntryInfo.Marshal(b, m, deterministic)
}
func (m *LdpBackoffEntryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffEntryInfo.Merge(m, src)
}
func (m *LdpBackoffEntryInfo) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffEntryInfo.Size(m)
}
func (m *LdpBackoffEntryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffEntryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffEntryInfo proto.InternalMessageInfo

func (m *LdpBackoffEntryInfo) GetBackoffSeconds() uint32 {
	if m != nil {
		return m.BackoffSeconds
	}
	return 0
}

func (m *LdpBackoffEntryInfo) GetWaitingSeconds() uint32 {
	if m != nil {
		return m.WaitingSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpBackoffEntryInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.backoffs.backoff.ldp_backoff_entry_info_KEYS")
	proto.RegisterType((*LdpBackoffEntryInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.backoffs.backoff.ldp_backoff_entry_info")
}

func init() { proto.RegisterFile("ldp_backoff_entry_info.proto", fileDescriptor_91b1b2b369a41567) }

var fileDescriptor_91b1b2b369a41567 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x40, 0x59, 0xc1, 0x82, 0x41, 0x2b, 0x2c, 0x28, 0x05, 0x3d, 0x94, 0x22, 0xd8, 0x53, 0x0e,
	0xf6, 0x37, 0x78, 0x28, 0x82, 0x87, 0xee, 0x49, 0x2f, 0x43, 0x36, 0x1f, 0x25, 0x3a, 0x66, 0x42,
	0x26, 0x56, 0xfd, 0xf7, 0x92, 0x6d, 0xb6, 0xa7, 0xbd, 0x25, 0x6f, 0x1e, 0x6f, 0x60, 0xc4, 0x3d,
	0x9a, 0x08, 0xbd, 0xd2, 0x9f, 0xe4, 0x1c, 0xd8, 0x90, 0xd3, 0x1f, 0xf8, 0xe0, 0x48, 0xc6, 0x44,
	0x99, 0xda, 0x57, 0xed, 0x59, 0x13, 0x78, 0x62, 0xf8, 0x4d, 0xf0, 0x15, 0x91, 0xa1, 0xf8, 0x14,
	0x6d, 0x92, 0xe3, 0x4f, 0xee, 0x91, 0x7a, 0x85, 0x52, 0xe9, 0xec, 0x0f, 0x56, 0x1a, 0xeb, 0xd4,
	0x37, 0x66, 0x38, 0x24, 0x27, 0x6b, 0x97, 0xc7, 0xc7, 0xea, 0x5d, 0xdc, 0x4d, 0xef, 0x83, 0x97,
	0xe7, 0xb7, 0xae, 0xbd, 0x11, 0x33, 0xe4, 0x04, 0xde, 0x2c, 0x9a, 0x65, 0xb3, 0xbe, 0xd8, 0x9d,
	0x23, 0xa7, 0xad, 0x69, 0x1f, 0xc4, 0x1c, 0x55, 0x6f, 0x11, 0x38, 0x2a, 0x6d, 0xcb, 0xf8, 0x6c,
	0xd9, 0xac, 0xaf, 0x76, 0x97, 0x03, 0xed, 0x0a, 0xdc, 0x9a, 0xd5, 0x87, 0xb8, 0x9d, 0x6e, 0xb7,
	0x8f, 0xe2, 0x7a, 0xa4, 0x6c, 0x35, 0x05, 0xc3, 0x8b, 0xa7, 0x21, 0x30, 0xaf, 0xb8, 0x3b, 0xd2,
	0x22, 0xfe, 0x28, 0x9f, 0x7d, 0xd8, 0x9f, 0xc4, 0xcd, 0x51, 0xac, 0xb8, 0x8a, 0xfd, 0x6c, 0x38,
	0xcf, 0xe6, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xff, 0x42, 0x3a, 0x3e, 0x01, 0x00, 0x00,
}
