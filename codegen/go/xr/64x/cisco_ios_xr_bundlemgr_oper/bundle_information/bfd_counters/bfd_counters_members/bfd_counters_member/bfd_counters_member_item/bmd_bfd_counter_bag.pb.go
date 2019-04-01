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
// source: bmd_bfd_counter_bag.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_bfd_counters_bfd_counters_members_bfd_counters_member_bfd_counters_member_item

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

type BmdBfdCounterBag_KEYS struct {
	MemberInterface      string   `protobuf:"bytes,1,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBfdCounterBag_KEYS) Reset()         { *m = BmdBfdCounterBag_KEYS{} }
func (m *BmdBfdCounterBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBfdCounterBag_KEYS) ProtoMessage()    {}
func (*BmdBfdCounterBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfea9523118c2d3a, []int{0}
}

func (m *BmdBfdCounterBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBfdCounterBag_KEYS.Unmarshal(m, b)
}
func (m *BmdBfdCounterBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBfdCounterBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBfdCounterBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBfdCounterBag_KEYS.Merge(m, src)
}
func (m *BmdBfdCounterBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBfdCounterBag_KEYS.Size(m)
}
func (m *BmdBfdCounterBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBfdCounterBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBfdCounterBag_KEYS proto.InternalMessageInfo

func (m *BmdBfdCounterBag_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
}

type BmdBfdCounterBag struct {
	MemberName                   string   `protobuf:"bytes,50,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	LastTimeCleared              uint64   `protobuf:"varint,51,opt,name=last_time_cleared,json=lastTimeCleared,proto3" json:"last_time_cleared,omitempty"`
	Starting                     uint32   `protobuf:"varint,52,opt,name=starting,proto3" json:"starting,omitempty"`
	Up                           uint32   `protobuf:"varint,53,opt,name=up,proto3" json:"up,omitempty"`
	Down                         uint32   `protobuf:"varint,54,opt,name=down,proto3" json:"down,omitempty"`
	NeighborUnconfigured         uint32   `protobuf:"varint,55,opt,name=neighbor_unconfigured,json=neighborUnconfigured,proto3" json:"neighbor_unconfigured,omitempty"`
	StartTimeouts                uint32   `protobuf:"varint,56,opt,name=start_timeouts,json=startTimeouts,proto3" json:"start_timeouts,omitempty"`
	NeighborUnconfiguredTimeouts uint32   `protobuf:"varint,57,opt,name=neighbor_unconfigured_timeouts,json=neighborUnconfiguredTimeouts,proto3" json:"neighbor_unconfigured_timeouts,omitempty"`
	TimeSinceCleared             uint64   `protobuf:"varint,58,opt,name=time_since_cleared,json=timeSinceCleared,proto3" json:"time_since_cleared,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *BmdBfdCounterBag) Reset()         { *m = BmdBfdCounterBag{} }
func (m *BmdBfdCounterBag) String() string { return proto.CompactTextString(m) }
func (*BmdBfdCounterBag) ProtoMessage()    {}
func (*BmdBfdCounterBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfea9523118c2d3a, []int{1}
}

func (m *BmdBfdCounterBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBfdCounterBag.Unmarshal(m, b)
}
func (m *BmdBfdCounterBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBfdCounterBag.Marshal(b, m, deterministic)
}
func (m *BmdBfdCounterBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBfdCounterBag.Merge(m, src)
}
func (m *BmdBfdCounterBag) XXX_Size() int {
	return xxx_messageInfo_BmdBfdCounterBag.Size(m)
}
func (m *BmdBfdCounterBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBfdCounterBag.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBfdCounterBag proto.InternalMessageInfo

func (m *BmdBfdCounterBag) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

func (m *BmdBfdCounterBag) GetLastTimeCleared() uint64 {
	if m != nil {
		return m.LastTimeCleared
	}
	return 0
}

func (m *BmdBfdCounterBag) GetStarting() uint32 {
	if m != nil {
		return m.Starting
	}
	return 0
}

func (m *BmdBfdCounterBag) GetUp() uint32 {
	if m != nil {
		return m.Up
	}
	return 0
}

func (m *BmdBfdCounterBag) GetDown() uint32 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *BmdBfdCounterBag) GetNeighborUnconfigured() uint32 {
	if m != nil {
		return m.NeighborUnconfigured
	}
	return 0
}

func (m *BmdBfdCounterBag) GetStartTimeouts() uint32 {
	if m != nil {
		return m.StartTimeouts
	}
	return 0
}

func (m *BmdBfdCounterBag) GetNeighborUnconfiguredTimeouts() uint32 {
	if m != nil {
		return m.NeighborUnconfiguredTimeouts
	}
	return 0
}

func (m *BmdBfdCounterBag) GetTimeSinceCleared() uint64 {
	if m != nil {
		return m.TimeSinceCleared
	}
	return 0
}

func init() {
	proto.RegisterType((*BmdBfdCounterBag_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_members.bfd_counters_member.bfd_counters_member_item.bmd_bfd_counter_bag_KEYS")
	proto.RegisterType((*BmdBfdCounterBag)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bfd_counters.bfd_counters_members.bfd_counters_member.bfd_counters_member_item.bmd_bfd_counter_bag")
}

func init() { proto.RegisterFile("bmd_bfd_counter_bag.proto", fileDescriptor_cfea9523118c2d3a) }

var fileDescriptor_cfea9523118c2d3a = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4b, 0x4b, 0x2b, 0x41,
	0x10, 0x85, 0x99, 0x10, 0x2e, 0xf7, 0xd6, 0x25, 0x0f, 0x5b, 0x85, 0x56, 0x44, 0x43, 0x40, 0x88,
	0x22, 0x59, 0x18, 0xdf, 0x5b, 0xcd, 0x42, 0x04, 0x17, 0x49, 0x5c, 0xb8, 0x2a, 0x7a, 0x66, 0x7a,
	0xc6, 0x86, 0x74, 0xf7, 0xd0, 0x0f, 0x74, 0xe3, 0x9f, 0xf6, 0x17, 0xc8, 0xf4, 0xe4, 0x25, 0x8c,
	0xbb, 0xae, 0xef, 0x54, 0x9d, 0x2a, 0x0e, 0x0d, 0x7b, 0xb1, 0x4c, 0x31, 0xce, 0x52, 0x4c, 0xb4,
	0x57, 0x8e, 0x1b, 0x8c, 0x59, 0x3e, 0x2c, 0x8c, 0x76, 0x9a, 0x7c, 0x26, 0xc2, 0x26, 0x1a, 0x85,
	0xb6, 0xf8, 0x61, 0x30, 0xf6, 0x2a, 0x9d, 0x73, 0x99, 0x1b, 0xd4, 0x05, 0x37, 0xc3, 0xaa, 0x44,
	0xa1, 0x32, 0x6d, 0x24, 0x73, 0x42, 0xab, 0xe1, 0x86, 0x8b, 0xfd, 0x51, 0xa0, 0xe4, 0x32, 0xfe,
	0x05, 0xd6, 0x31, 0x14, 0x8e, 0xcb, 0xfe, 0x18, 0x68, 0xcd, 0x6d, 0xf8, 0x34, 0x7e, 0x9d, 0x92,
	0x13, 0xe8, 0x2e, 0x5b, 0x4b, 0x21, 0x63, 0x09, 0xa7, 0x51, 0x2f, 0x1a, 0xfc, 0x9b, 0x74, 0x2a,
	0xfe, 0xb8, 0xc4, 0xfd, 0xaf, 0x06, 0x6c, 0xd7, 0xf8, 0x90, 0x23, 0xf8, 0xbf, 0xb0, 0x50, 0x4c,
	0x72, 0x7a, 0x1e, 0xa6, 0xa1, 0x42, 0xcf, 0x4c, 0x72, 0x72, 0x0a, 0x5b, 0x73, 0x66, 0x1d, 0x3a,
	0x21, 0x39, 0x26, 0x73, 0xce, 0x0c, 0x4f, 0xe9, 0xa8, 0x17, 0x0d, 0x9a, 0x93, 0x4e, 0x29, 0xcc,
	0x84, 0xe4, 0xf7, 0x15, 0x26, 0xfb, 0xf0, 0xd7, 0x3a, 0x66, 0x9c, 0x50, 0x39, 0xbd, 0xe8, 0x45,
	0x83, 0xd6, 0x64, 0x55, 0x93, 0x36, 0x34, 0x7c, 0x41, 0x2f, 0x03, 0x6d, 0xf8, 0x82, 0x10, 0x68,
	0xa6, 0xfa, 0x5d, 0xd1, 0xab, 0x40, 0xc2, 0x9b, 0x8c, 0x60, 0x57, 0x71, 0x91, 0xbf, 0xc5, 0xda,
	0xa0, 0x57, 0x89, 0x56, 0x99, 0xc8, 0x7d, 0xb9, 0xef, 0x3a, 0x34, 0xed, 0x2c, 0xc5, 0x97, 0x0d,
	0x8d, 0x1c, 0x43, 0x3b, 0x2c, 0x09, 0x17, 0x6a, 0xef, 0x2c, 0xbd, 0x09, 0xdd, 0xad, 0x40, 0x67,
	0x0b, 0x48, 0x1e, 0xe0, 0xb0, 0xd6, 0x7b, 0x3d, 0x76, 0x1b, 0xc6, 0x0e, 0xea, 0x96, 0xac, 0x5c,
	0xce, 0x80, 0x84, 0x20, 0xac, 0x50, 0xc9, 0x3a, 0x8e, 0xbb, 0x10, 0x47, 0xb7, 0x54, 0xa6, 0xa5,
	0xb0, 0xc8, 0x23, 0xfe, 0x13, 0x7e, 0xd0, 0xe8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x57, 0x34, 0xe7,
	0x63, 0x5e, 0x02, 0x00, 0x00,
}
