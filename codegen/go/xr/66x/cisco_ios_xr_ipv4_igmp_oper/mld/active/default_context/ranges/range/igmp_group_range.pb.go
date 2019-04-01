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
// source: igmp_group_range.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_default_context_ranges_range

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

type IgmpGroupRange_KEYS struct {
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	GroupMask            uint32   `protobuf:"varint,2,opt,name=group_mask,json=groupMask,proto3" json:"group_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpGroupRange_KEYS) Reset()         { *m = IgmpGroupRange_KEYS{} }
func (m *IgmpGroupRange_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpGroupRange_KEYS) ProtoMessage()    {}
func (*IgmpGroupRange_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_41fc09d7bbd7746f, []int{0}
}

func (m *IgmpGroupRange_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpGroupRange_KEYS.Unmarshal(m, b)
}
func (m *IgmpGroupRange_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpGroupRange_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpGroupRange_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpGroupRange_KEYS.Merge(m, src)
}
func (m *IgmpGroupRange_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpGroupRange_KEYS.Size(m)
}
func (m *IgmpGroupRange_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpGroupRange_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpGroupRange_KEYS proto.InternalMessageInfo

func (m *IgmpGroupRange_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *IgmpGroupRange_KEYS) GetGroupMask() uint32 {
	if m != nil {
		return m.GroupMask
	}
	return 0
}

type IgmpAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpAddrtype) Reset()         { *m = IgmpAddrtype{} }
func (m *IgmpAddrtype) String() string { return proto.CompactTextString(m) }
func (*IgmpAddrtype) ProtoMessage()    {}
func (*IgmpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_41fc09d7bbd7746f, []int{1}
}

func (m *IgmpAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpAddrtype.Unmarshal(m, b)
}
func (m *IgmpAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpAddrtype.Marshal(b, m, deterministic)
}
func (m *IgmpAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpAddrtype.Merge(m, src)
}
func (m *IgmpAddrtype) XXX_Size() int {
	return xxx_messageInfo_IgmpAddrtype.Size(m)
}
func (m *IgmpAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpAddrtype proto.InternalMessageInfo

func (m *IgmpAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type IgmpGroupRange struct {
	GroupAddressXr       *IgmpAddrtype `protobuf:"bytes,50,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	PrefixLength         uint32        `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Protocol             string        `protobuf:"bytes,52,opt,name=protocol,proto3" json:"protocol,omitempty"`
	IsStale              bool          `protobuf:"varint,53,opt,name=is_stale,json=isStale,proto3" json:"is_stale,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpGroupRange) Reset()         { *m = IgmpGroupRange{} }
func (m *IgmpGroupRange) String() string { return proto.CompactTextString(m) }
func (*IgmpGroupRange) ProtoMessage()    {}
func (*IgmpGroupRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_41fc09d7bbd7746f, []int{2}
}

func (m *IgmpGroupRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpGroupRange.Unmarshal(m, b)
}
func (m *IgmpGroupRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpGroupRange.Marshal(b, m, deterministic)
}
func (m *IgmpGroupRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpGroupRange.Merge(m, src)
}
func (m *IgmpGroupRange) XXX_Size() int {
	return xxx_messageInfo_IgmpGroupRange.Size(m)
}
func (m *IgmpGroupRange) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpGroupRange.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpGroupRange proto.InternalMessageInfo

func (m *IgmpGroupRange) GetGroupAddressXr() *IgmpAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *IgmpGroupRange) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *IgmpGroupRange) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *IgmpGroupRange) GetIsStale() bool {
	if m != nil {
		return m.IsStale
	}
	return false
}

func init() {
	proto.RegisterType((*IgmpGroupRange_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ranges.range.igmp_group_range_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ranges.range.igmp_addrtype")
	proto.RegisterType((*IgmpGroupRange)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.ranges.range.igmp_group_range")
}

func init() { proto.RegisterFile("igmp_group_range.proto", fileDescriptor_41fc09d7bbd7746f) }

var fileDescriptor_41fc09d7bbd7746f = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x4b, 0xfb, 0x30,
	0x1c, 0xc6, 0xe9, 0x7e, 0xb0, 0x3f, 0xd9, 0xf6, 0x63, 0x04, 0xd4, 0x2a, 0x08, 0x75, 0x5e, 0x7a,
	0xca, 0x61, 0x9b, 0xbb, 0x8b, 0x78, 0xf2, 0xcf, 0xa1, 0xbb, 0x28, 0x1e, 0x42, 0x6c, 0xbf, 0xad,
	0x61, 0x6d, 0x13, 0x92, 0x6c, 0x54, 0xf0, 0x25, 0xfb, 0x22, 0xa4, 0xdf, 0x6e, 0xc5, 0xed, 0xec,
	0xa9, 0xe4, 0xd3, 0x0f, 0x4f, 0x78, 0x9e, 0x90, 0x53, 0x99, 0x15, 0x9a, 0x67, 0x46, 0x6d, 0x34,
	0x37, 0xa2, 0xcc, 0x80, 0x69, 0xa3, 0x9c, 0xa2, 0x77, 0xb1, 0xb4, 0xb1, 0xe2, 0x52, 0x59, 0x5e,
	0x19, 0x2e, 0xf5, 0x76, 0xc1, 0xd1, 0x54, 0x1a, 0x0c, 0x2b, 0xf2, 0x84, 0x89, 0xd8, 0xc9, 0x2d,
	0xb0, 0x04, 0x52, 0xb1, 0xc9, 0x1d, 0x8f, 0x55, 0xe9, 0xa0, 0x72, 0x0c, 0x33, 0x6c, 0xf3, 0x99,
	0xbe, 0x91, 0x93, 0xe3, 0x78, 0xfe, 0x70, 0xff, 0xba, 0xa2, 0xd7, 0x64, 0xdc, 0x30, 0x91, 0x24,
	0x06, 0xac, 0xf5, 0xbd, 0xc0, 0x0b, 0x07, 0xd1, 0x08, 0xe1, 0x6d, 0xc3, 0xe8, 0x25, 0x21, 0x8d,
	0x54, 0x08, 0xbb, 0xf6, 0x3b, 0x81, 0x17, 0x8e, 0xa3, 0x01, 0x92, 0x27, 0x61, 0xd7, 0xd3, 0x92,
	0x8c, 0x31, 0xbc, 0x8e, 0x70, 0x9f, 0x1a, 0xe8, 0x19, 0xe9, 0x89, 0x94, 0x97, 0xa2, 0x80, 0x5d,
	0x5c, 0x57, 0xa4, 0xcf, 0xa2, 0x00, 0x7a, 0x45, 0x46, 0x58, 0x60, 0x7f, 0x59, 0x07, 0xff, 0x0e,
	0x6b, 0xb6, 0xbf, 0xab, 0x51, 0x96, 0xad, 0xf2, 0xaf, 0x55, 0x96, 0x3b, 0x65, 0xfa, 0xed, 0x91,
	0xc9, 0x71, 0x1b, 0xfa, 0x45, 0x26, 0x07, 0x45, 0x78, 0x65, 0xfc, 0x59, 0xe0, 0x85, 0xc3, 0x59,
	0xc4, 0xfe, 0x60, 0x41, 0x76, 0xd0, 0x30, 0xfa, 0xff, 0x7b, 0x9f, 0x17, 0x53, 0xcf, 0xa8, 0x0d,
	0xa4, 0xb2, 0xe2, 0x39, 0x94, 0x99, 0xfb, 0xf0, 0xe7, 0x38, 0xd2, 0xa8, 0x81, 0x8f, 0xc8, 0xe8,
	0x05, 0xe9, 0xe3, 0x93, 0xc6, 0x2a, 0xf7, 0x17, 0x58, 0xab, 0x3d, 0xd3, 0x73, 0xd2, 0x97, 0x96,
	0x5b, 0x27, 0x72, 0xf0, 0x6f, 0x02, 0x2f, 0xec, 0x47, 0x3d, 0x69, 0x57, 0xf5, 0xf1, 0xbd, 0x8b,
	0xd2, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xb0, 0x22, 0x31, 0x21, 0x02, 0x00, 0x00,
}