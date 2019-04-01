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
// source: lacp_counters.proto

package cisco_ios_xr_bundlemgr_oper_lacp_bundle_members_members_member_counters

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

type LacpCounters_KEYS struct {
	MemberInterface      string   `protobuf:"bytes,1,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpCounters_KEYS) Reset()         { *m = LacpCounters_KEYS{} }
func (m *LacpCounters_KEYS) String() string { return proto.CompactTextString(m) }
func (*LacpCounters_KEYS) ProtoMessage()    {}
func (*LacpCounters_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e05b38dbf41ae75, []int{0}
}

func (m *LacpCounters_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpCounters_KEYS.Unmarshal(m, b)
}
func (m *LacpCounters_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpCounters_KEYS.Marshal(b, m, deterministic)
}
func (m *LacpCounters_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpCounters_KEYS.Merge(m, src)
}
func (m *LacpCounters_KEYS) XXX_Size() int {
	return xxx_messageInfo_LacpCounters_KEYS.Size(m)
}
func (m *LacpCounters_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpCounters_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LacpCounters_KEYS proto.InternalMessageInfo

func (m *LacpCounters_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
}

type LacpCounters struct {
	LacpdUsReceived             uint32   `protobuf:"varint,50,opt,name=lacpd_us_received,json=lacpdUsReceived,proto3" json:"lacpd_us_received,omitempty"`
	LacpdUsTransmitted          uint32   `protobuf:"varint,51,opt,name=lacpd_us_transmitted,json=lacpdUsTransmitted,proto3" json:"lacpd_us_transmitted,omitempty"`
	MarkerPacketsReceived       uint32   `protobuf:"varint,52,opt,name=marker_packets_received,json=markerPacketsReceived,proto3" json:"marker_packets_received,omitempty"`
	MarkerResponsesTransmitted  uint32   `protobuf:"varint,53,opt,name=marker_responses_transmitted,json=markerResponsesTransmitted,proto3" json:"marker_responses_transmitted,omitempty"`
	IllegalPacketsReceived      uint32   `protobuf:"varint,54,opt,name=illegal_packets_received,json=illegalPacketsReceived,proto3" json:"illegal_packets_received,omitempty"`
	ExcessLacpdUsReceived       uint32   `protobuf:"varint,55,opt,name=excess_lacpd_us_received,json=excessLacpdUsReceived,proto3" json:"excess_lacpd_us_received,omitempty"`
	ExcessMarkerPacketsReceived uint32   `protobuf:"varint,56,opt,name=excess_marker_packets_received,json=excessMarkerPacketsReceived,proto3" json:"excess_marker_packets_received,omitempty"`
	Defaulted                   uint32   `protobuf:"varint,57,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
	Expired                     uint32   `protobuf:"varint,58,opt,name=expired,proto3" json:"expired,omitempty"`
	TimeSinceLastLacpduReceived uint64   `protobuf:"varint,59,opt,name=time_since_last_lacpdu_received,json=timeSinceLastLacpduReceived,proto3" json:"time_since_last_lacpdu_received,omitempty"`
	TimeSinceCleared            uint64   `protobuf:"varint,60,opt,name=time_since_cleared,json=timeSinceCleared,proto3" json:"time_since_cleared,omitempty"`
	TimeSinceUnexpectedEvent    uint64   `protobuf:"varint,61,opt,name=time_since_unexpected_event,json=timeSinceUnexpectedEvent,proto3" json:"time_since_unexpected_event,omitempty"`
	LastClearedSec              uint32   `protobuf:"varint,62,opt,name=last_cleared_sec,json=lastClearedSec,proto3" json:"last_cleared_sec,omitempty"`
	LastClearedNsec             uint32   `protobuf:"varint,63,opt,name=last_cleared_nsec,json=lastClearedNsec,proto3" json:"last_cleared_nsec,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *LacpCounters) Reset()         { *m = LacpCounters{} }
func (m *LacpCounters) String() string { return proto.CompactTextString(m) }
func (*LacpCounters) ProtoMessage()    {}
func (*LacpCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e05b38dbf41ae75, []int{1}
}

func (m *LacpCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpCounters.Unmarshal(m, b)
}
func (m *LacpCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpCounters.Marshal(b, m, deterministic)
}
func (m *LacpCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpCounters.Merge(m, src)
}
func (m *LacpCounters) XXX_Size() int {
	return xxx_messageInfo_LacpCounters.Size(m)
}
func (m *LacpCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpCounters.DiscardUnknown(m)
}

var xxx_messageInfo_LacpCounters proto.InternalMessageInfo

func (m *LacpCounters) GetLacpdUsReceived() uint32 {
	if m != nil {
		return m.LacpdUsReceived
	}
	return 0
}

func (m *LacpCounters) GetLacpdUsTransmitted() uint32 {
	if m != nil {
		return m.LacpdUsTransmitted
	}
	return 0
}

func (m *LacpCounters) GetMarkerPacketsReceived() uint32 {
	if m != nil {
		return m.MarkerPacketsReceived
	}
	return 0
}

func (m *LacpCounters) GetMarkerResponsesTransmitted() uint32 {
	if m != nil {
		return m.MarkerResponsesTransmitted
	}
	return 0
}

func (m *LacpCounters) GetIllegalPacketsReceived() uint32 {
	if m != nil {
		return m.IllegalPacketsReceived
	}
	return 0
}

func (m *LacpCounters) GetExcessLacpdUsReceived() uint32 {
	if m != nil {
		return m.ExcessLacpdUsReceived
	}
	return 0
}

func (m *LacpCounters) GetExcessMarkerPacketsReceived() uint32 {
	if m != nil {
		return m.ExcessMarkerPacketsReceived
	}
	return 0
}

func (m *LacpCounters) GetDefaulted() uint32 {
	if m != nil {
		return m.Defaulted
	}
	return 0
}

func (m *LacpCounters) GetExpired() uint32 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *LacpCounters) GetTimeSinceLastLacpduReceived() uint64 {
	if m != nil {
		return m.TimeSinceLastLacpduReceived
	}
	return 0
}

func (m *LacpCounters) GetTimeSinceCleared() uint64 {
	if m != nil {
		return m.TimeSinceCleared
	}
	return 0
}

func (m *LacpCounters) GetTimeSinceUnexpectedEvent() uint64 {
	if m != nil {
		return m.TimeSinceUnexpectedEvent
	}
	return 0
}

func (m *LacpCounters) GetLastClearedSec() uint32 {
	if m != nil {
		return m.LastClearedSec
	}
	return 0
}

func (m *LacpCounters) GetLastClearedNsec() uint32 {
	if m != nil {
		return m.LastClearedNsec
	}
	return 0
}

func init() {
	proto.RegisterType((*LacpCounters_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundle_members.members.member.counters.lacp_counters_KEYS")
	proto.RegisterType((*LacpCounters)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundle_members.members.member.counters.lacp_counters")
}

func init() { proto.RegisterFile("lacp_counters.proto", fileDescriptor_2e05b38dbf41ae75) }

var fileDescriptor_2e05b38dbf41ae75 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdd, 0x6b, 0x14, 0x31,
	0x14, 0xc5, 0x59, 0xf0, 0x83, 0x5e, 0xa8, 0x5d, 0xe3, 0x57, 0x70, 0x8b, 0x2e, 0x7d, 0x5a, 0x45,
	0x16, 0xb1, 0xda, 0xd6, 0x8f, 0x5a, 0xa1, 0x16, 0x11, 0xab, 0xc8, 0xac, 0x7d, 0xf0, 0xe9, 0x92,
	0xcd, 0xdc, 0x96, 0xd0, 0x4c, 0x66, 0x48, 0x32, 0x65, 0xff, 0x5c, 0xff, 0x14, 0x49, 0x32, 0x33,
	0x3b, 0xdb, 0xd5, 0xa7, 0x65, 0xcf, 0xf9, 0x9d, 0x7b, 0x72, 0x43, 0x06, 0xee, 0x69, 0x21, 0x2b,
	0x94, 0x65, 0x6d, 0x3c, 0x59, 0x37, 0xad, 0x6c, 0xe9, 0x4b, 0xf6, 0x45, 0x2a, 0x27, 0x4b, 0x54,
	0xa5, 0xc3, 0x85, 0xc5, 0x79, 0x6d, 0x72, 0x4d, 0xc5, 0x85, 0xc5, 0xb2, 0x22, 0x3b, 0x8d, 0x81,
	0xa4, 0x61, 0x41, 0xc5, 0x3c, 0xc4, 0x56, 0x7f, 0xa7, 0xed, 0xb8, 0x9d, 0x23, 0x60, 0x2b, 0xf3,
	0xf1, 0xdb, 0xc9, 0xef, 0x19, 0x7b, 0x06, 0xc3, 0x04, 0xa2, 0x0a, 0xea, 0xb9, 0x90, 0xc4, 0x07,
	0xe3, 0xc1, 0x64, 0x23, 0xdb, 0x4a, 0xfa, 0xd7, 0x56, 0xde, 0xf9, 0x73, 0x13, 0x36, 0x57, 0x26,
	0xb0, 0xe7, 0x70, 0x37, 0x08, 0x39, 0xd6, 0x0e, 0x2d, 0x49, 0x52, 0x57, 0x94, 0xf3, 0x57, 0xe3,
	0xc1, 0x64, 0x33, 0xdb, 0x8a, 0xc6, 0x99, 0xcb, 0x1a, 0x99, 0xbd, 0x84, 0xfb, 0x1d, 0xeb, 0xad,
	0x30, 0xae, 0x50, 0xde, 0x53, 0xce, 0x77, 0x23, 0xce, 0x1a, 0xfc, 0xd7, 0xd2, 0x61, 0x7b, 0xf0,
	0xa8, 0x10, 0xf6, 0x92, 0x2c, 0x56, 0x42, 0x5e, 0x92, 0xef, 0x75, 0xbc, 0x8e, 0xa1, 0x07, 0xc9,
	0xfe, 0x99, 0xdc, 0xae, 0xe9, 0x13, 0x6c, 0x37, 0x39, 0x4b, 0xae, 0x2a, 0x8d, 0xa3, 0xd5, 0xc6,
	0x37, 0x31, 0xfc, 0x38, 0x31, 0x59, 0x8b, 0xf4, 0x9b, 0x0f, 0x80, 0x2b, 0xad, 0xe9, 0x42, 0xe8,
	0xf5, 0xea, 0xbd, 0x98, 0x7e, 0xd8, 0xf8, 0xd7, 0xbb, 0xf7, 0x81, 0xd3, 0x42, 0x92, 0x73, 0xb8,
	0x7e, 0x31, 0xfb, 0xe9, 0xd0, 0xc9, 0x3f, 0xbd, 0x76, 0x3d, 0xc7, 0xf0, 0xa4, 0x09, 0xfe, 0x6f,
	0xe7, 0x83, 0x18, 0x1f, 0x25, 0xea, 0xfb, 0x3f, 0x37, 0xdf, 0x86, 0x8d, 0x9c, 0xce, 0x45, 0xad,
	0xc3, 0x9a, 0x6f, 0x23, 0xbf, 0x14, 0x18, 0x87, 0xdb, 0xb4, 0xa8, 0x94, 0xa5, 0x9c, 0xbf, 0x8b,
	0x5e, 0xfb, 0x97, 0x7d, 0x86, 0xa7, 0x5e, 0x15, 0x84, 0x4e, 0x19, 0x49, 0xa8, 0x85, 0xf3, 0xe9,
	0xf8, 0xf5, 0xb2, 0xfd, 0xfd, 0x78, 0x30, 0xb9, 0x91, 0x8d, 0x02, 0x36, 0x0b, 0xd4, 0xa9, 0x70,
	0x3e, 0xee, 0x50, 0x77, 0xed, 0x2f, 0x80, 0xf5, 0xa6, 0x48, 0x4d, 0x22, 0x54, 0x7d, 0x88, 0xc1,
	0x61, 0x17, 0x3c, 0x4e, 0x3a, 0x3b, 0x84, 0x51, 0x8f, 0xae, 0x0d, 0x2d, 0x2a, 0x92, 0x9e, 0x72,
	0xa4, 0x2b, 0x32, 0x9e, 0x1f, 0xc6, 0x18, 0xef, 0x62, 0x67, 0x1d, 0x70, 0x12, 0x7c, 0x36, 0x81,
	0x61, 0x3c, 0x67, 0x53, 0x83, 0x8e, 0x24, 0xff, 0x18, 0xb7, 0xba, 0x13, 0xf4, 0xa6, 0x65, 0x46,
	0x32, 0x3d, 0xd2, 0x1e, 0x69, 0x02, 0x7a, 0xd4, 0x3e, 0xd2, 0x0e, 0xfd, 0xe1, 0x48, 0xce, 0x6f,
	0xc5, 0x6f, 0x6e, 0xf7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x63, 0x53, 0x62, 0x8a, 0x03,
	0x00, 0x00,
}
