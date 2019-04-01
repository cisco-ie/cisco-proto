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
// source: rsvp_sync_count_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_counters_issu

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

type RsvpSyncCountInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpSyncCountInfo_KEYS) Reset()         { *m = RsvpSyncCountInfo_KEYS{} }
func (m *RsvpSyncCountInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpSyncCountInfo_KEYS) ProtoMessage()    {}
func (*RsvpSyncCountInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a7ee7d371292513, []int{0}
}

func (m *RsvpSyncCountInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSyncCountInfo_KEYS.Merge(m, src)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpSyncCountInfo_KEYS.Size(m)
}
func (m *RsvpSyncCountInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSyncCountInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSyncCountInfo_KEYS proto.InternalMessageInfo

type RsvpSyncCountInfo struct {
	LastClearedTimestamp uint32   `protobuf:"varint,50,opt,name=last_cleared_timestamp,json=lastClearedTimestamp,proto3" json:"last_cleared_timestamp,omitempty"`
	RsvpProcessRole      string   `protobuf:"bytes,51,opt,name=rsvp_process_role,json=rsvpProcessRole,proto3" json:"rsvp_process_role,omitempty"`
	LastIdtStates        uint32   `protobuf:"varint,52,opt,name=last_idt_states,json=lastIdtStates,proto3" json:"last_idt_states,omitempty"`
	TotalStates          uint32   `protobuf:"varint,53,opt,name=total_states,json=totalStates,proto3" json:"total_states,omitempty"`
	TotalDeletions       uint32   `protobuf:"varint,54,opt,name=total_deletions,json=totalDeletions,proto3" json:"total_deletions,omitempty"`
	TotalNacks           uint64   `protobuf:"varint,55,opt,name=total_nacks,json=totalNacks,proto3" json:"total_nacks,omitempty"`
	TotalIdTs            uint32   `protobuf:"varint,56,opt,name=total_id_ts,json=totalIdTs,proto3" json:"total_id_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpSyncCountInfo) Reset()         { *m = RsvpSyncCountInfo{} }
func (m *RsvpSyncCountInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpSyncCountInfo) ProtoMessage()    {}
func (*RsvpSyncCountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a7ee7d371292513, []int{1}
}

func (m *RsvpSyncCountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSyncCountInfo.Unmarshal(m, b)
}
func (m *RsvpSyncCountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSyncCountInfo.Marshal(b, m, deterministic)
}
func (m *RsvpSyncCountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSyncCountInfo.Merge(m, src)
}
func (m *RsvpSyncCountInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpSyncCountInfo.Size(m)
}
func (m *RsvpSyncCountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSyncCountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSyncCountInfo proto.InternalMessageInfo

func (m *RsvpSyncCountInfo) GetLastClearedTimestamp() uint32 {
	if m != nil {
		return m.LastClearedTimestamp
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetRsvpProcessRole() string {
	if m != nil {
		return m.RsvpProcessRole
	}
	return ""
}

func (m *RsvpSyncCountInfo) GetLastIdtStates() uint32 {
	if m != nil {
		return m.LastIdtStates
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalStates() uint32 {
	if m != nil {
		return m.TotalStates
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalDeletions() uint32 {
	if m != nil {
		return m.TotalDeletions
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalNacks() uint64 {
	if m != nil {
		return m.TotalNacks
	}
	return 0
}

func (m *RsvpSyncCountInfo) GetTotalIdTs() uint32 {
	if m != nil {
		return m.TotalIdTs
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpSyncCountInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.issu.rsvp_sync_count_info_KEYS")
	proto.RegisterType((*RsvpSyncCountInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.issu.rsvp_sync_count_info")
}

func init() { proto.RegisterFile("rsvp_sync_count_info.proto", fileDescriptor_8a7ee7d371292513) }

var fileDescriptor_8a7ee7d371292513 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0xe9, 0xcb, 0x8b, 0xd0, 0xd5, 0x5a, 0x5c, 0x8a, 0x44, 0x05, 0xad, 0x3d, 0x68, 0x10,
	0xc9, 0xc1, 0xd6, 0x3f, 0x77, 0xf5, 0x50, 0x04, 0x91, 0xb4, 0x17, 0x4f, 0x43, 0xdc, 0x5d, 0x61,
	0x71, 0x9b, 0x59, 0x76, 0xa6, 0xa2, 0x1f, 0xca, 0xef, 0x28, 0x99, 0x98, 0x5b, 0x6f, 0xe1, 0x79,
	0x9e, 0xf9, 0xe5, 0xb0, 0xea, 0x30, 0xd1, 0x67, 0x04, 0xfa, 0xae, 0x0d, 0x18, 0x5c, 0xd7, 0x0c,
	0xbe, 0x7e, 0xc7, 0x22, 0x26, 0x64, 0xd4, 0x97, 0xc6, 0x93, 0x41, 0xf0, 0x48, 0xf0, 0x95, 0xc0,
	0x47, 0x90, 0x16, 0xa3, 0x4b, 0x45, 0xf3, 0x55, 0xc8, 0x81, 0x4b, 0x54, 0x78, 0xa2, 0xf5, 0xe4,
	0x48, 0x1d, 0x6c, 0xda, 0x82, 0xa7, 0xc7, 0xd7, 0xc5, 0xe4, 0xe7, 0x9f, 0x1a, 0x6d, 0xb2, 0x7a,
	0xa6, 0xf6, 0x43, 0x45, 0x0c, 0x26, 0xb8, 0x2a, 0x39, 0x0b, 0xec, 0x57, 0x8e, 0xb8, 0x5a, 0xc5,
	0xec, 0x6a, 0xdc, 0xcb, 0x07, 0xe5, 0xa8, 0xb1, 0xf7, 0xad, 0x5c, 0x76, 0x4e, 0x5f, 0xa8, 0x3d,
	0x59, 0x8b, 0x09, 0x8d, 0x23, 0x82, 0x84, 0xc1, 0x65, 0xd3, 0x71, 0x2f, 0xef, 0x97, 0xc3, 0x46,
	0xbc, 0xb4, 0xbc, 0xc4, 0xe0, 0xf4, 0x99, 0x1a, 0xca, 0x1f, 0xbc, 0x65, 0x20, 0xae, 0xd8, 0x51,
	0x36, 0x93, 0xe9, 0x41, 0x83, 0xe7, 0x96, 0x17, 0x02, 0xf5, 0xa9, 0xda, 0x61, 0xe4, 0x2a, 0x74,
	0xd1, 0xb5, 0x44, 0xdb, 0xc2, 0xfe, 0x92, 0x73, 0x35, 0x6c, 0x13, 0xeb, 0x82, 0x63, 0x8f, 0x35,
	0x65, 0x37, 0x52, 0xed, 0x0a, 0x7e, 0xe8, 0xa8, 0x3e, 0x51, 0xed, 0x1d, 0xd4, 0x95, 0xf9, 0xa0,
	0xec, 0x76, 0xdc, 0xcb, 0xff, 0x97, 0x4a, 0xd0, 0x73, 0x43, 0xf4, 0x71, 0x17, 0x78, 0x0b, 0x4c,
	0xd9, 0x9d, 0xac, 0xf4, 0x05, 0xcd, 0xed, 0x92, 0xde, 0xb6, 0xe4, 0x05, 0xa6, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd7, 0x46, 0xd3, 0x8c, 0x9f, 0x01, 0x00, 0x00,
}