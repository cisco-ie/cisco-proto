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

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_default_context_ranges_range

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
	proto.RegisterType((*IgmpGroupRange_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.ranges.range.igmp_group_range_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.ranges.range.igmp_addrtype")
	proto.RegisterType((*IgmpGroupRange)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.ranges.range.igmp_group_range")
}

func init() { proto.RegisterFile("igmp_group_range.proto", fileDescriptor_41fc09d7bbd7746f) }

var fileDescriptor_41fc09d7bbd7746f = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x3d, 0x4f, 0xf3, 0x30,
	0x14, 0x85, 0x95, 0xbe, 0x52, 0x3f, 0xdc, 0xf6, 0x55, 0x65, 0x09, 0x08, 0x48, 0x48, 0xa1, 0x2c,
	0x99, 0x3c, 0xb4, 0xa5, 0x3b, 0x43, 0x27, 0x3e, 0x86, 0x94, 0x01, 0xc4, 0x70, 0xe5, 0x26, 0x4e,
	0xb0, 0x9a, 0xc4, 0x96, 0xed, 0xa2, 0x74, 0xe1, 0x27, 0xf3, 0x1b, 0x50, 0x6e, 0xda, 0x8a, 0x76,
	0x66, 0xb2, 0xef, 0xf1, 0xa3, 0x73, 0x75, 0x8e, 0xc9, 0xb9, 0xcc, 0x0a, 0x0d, 0x99, 0x51, 0x1b,
	0x0d, 0x86, 0x97, 0x99, 0x60, 0xda, 0x28, 0xa7, 0xe8, 0x22, 0x96, 0x36, 0x56, 0x20, 0x95, 0x85,
	0xca, 0x80, 0xd4, 0x9f, 0x33, 0x40, 0x52, 0x69, 0x61, 0x58, 0x7d, 0x63, 0xd6, 0xf1, 0x32, 0x59,
	0x6d, 0x59, 0x22, 0x52, 0xbe, 0xc9, 0x1d, 0xc4, 0xaa, 0x74, 0xa2, 0x72, 0x0c, 0x5d, 0x6c, 0x73,
	0x8c, 0xdf, 0xc9, 0xd9, 0xe9, 0x02, 0x78, 0x58, 0xbc, 0x2d, 0xe9, 0x2d, 0x19, 0x36, 0x1a, 0x4f,
	0x12, 0x23, 0xac, 0xf5, 0xbd, 0xc0, 0x0b, 0x7b, 0xd1, 0x00, 0xc5, 0xfb, 0x46, 0xa3, 0xd7, 0x84,
	0x34, 0x50, 0xc1, 0xed, 0xda, 0x6f, 0x05, 0x5e, 0x38, 0x8c, 0x7a, 0xa8, 0x3c, 0x71, 0xbb, 0x1e,
	0x97, 0x64, 0x88, 0xe6, 0xb5, 0x85, 0xdb, 0x6a, 0x41, 0x2f, 0x48, 0x87, 0xa7, 0x50, 0xf2, 0x42,
	0xec, 0xec, 0xda, 0x3c, 0x7d, 0xe6, 0x85, 0xa0, 0x37, 0x64, 0x80, 0x11, 0xf6, 0xcb, 0x5a, 0xf8,
	0xda, 0xaf, 0xb5, 0xfd, 0xae, 0x06, 0x99, 0x1f, 0x90, 0x7f, 0x07, 0x64, 0xbe, 0x43, 0xc6, 0xdf,
	0x1e, 0x19, 0x9d, 0xa6, 0xa1, 0x5f, 0x64, 0x74, 0x14, 0x04, 0x2a, 0xe3, 0x4f, 0x02, 0x2f, 0xec,
	0x4f, 0x5e, 0xd8, 0x9f, 0x74, 0xc8, 0x8e, 0x32, 0x46, 0xff, 0x7f, 0x37, 0xf4, 0x6a, 0xea, 0x22,
	0xb5, 0x11, 0xa9, 0xac, 0x20, 0x17, 0x65, 0xe6, 0x3e, 0xfc, 0x29, 0xd6, 0x34, 0x68, 0xc4, 0x47,
	0xd4, 0xe8, 0x15, 0xe9, 0xe2, 0xb7, 0xc6, 0x2a, 0xf7, 0x67, 0x18, 0xec, 0x30, 0xd3, 0x4b, 0xd2,
	0x95, 0x16, 0xac, 0xe3, 0xb9, 0xf0, 0xef, 0x02, 0x2f, 0xec, 0x46, 0x1d, 0x69, 0x97, 0xf5, 0xb8,
	0x6a, 0x23, 0x34, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x85, 0x11, 0xc8, 0x2c, 0x25, 0x02, 0x00,
	0x00,
}
