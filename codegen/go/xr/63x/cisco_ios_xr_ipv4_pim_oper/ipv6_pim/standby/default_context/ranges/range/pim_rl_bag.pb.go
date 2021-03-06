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
// source: pim_rl_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_ranges_range

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

type PimRlBag_KEYS struct {
	RpAddress            string   `protobuf:"bytes,1,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRlBag_KEYS) Reset()         { *m = PimRlBag_KEYS{} }
func (m *PimRlBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRlBag_KEYS) ProtoMessage()    {}
func (*PimRlBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7cdc6ebd16448, []int{0}
}

func (m *PimRlBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRlBag_KEYS.Unmarshal(m, b)
}
func (m *PimRlBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRlBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRlBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRlBag_KEYS.Merge(m, src)
}
func (m *PimRlBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRlBag_KEYS.Size(m)
}
func (m *PimRlBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRlBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRlBag_KEYS proto.InternalMessageInfo

func (m *PimRlBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

func (m *PimRlBag_KEYS) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7cdc6ebd16448, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimRlRpRangeBag struct {
	Prefix               *PimAddrtype `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         int32        `protobuf:"zigzag32,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Uptime               uint64       `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Expires              uint64       `protobuf:"varint,4,opt,name=expires,proto3" json:"expires,omitempty"`
	SourceOfInformation  *PimAddrtype `protobuf:"bytes,5,opt,name=source_of_information,json=sourceOfInformation,proto3" json:"source_of_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimRlRpRangeBag) Reset()         { *m = PimRlRpRangeBag{} }
func (m *PimRlRpRangeBag) String() string { return proto.CompactTextString(m) }
func (*PimRlRpRangeBag) ProtoMessage()    {}
func (*PimRlRpRangeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7cdc6ebd16448, []int{2}
}

func (m *PimRlRpRangeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRlRpRangeBag.Unmarshal(m, b)
}
func (m *PimRlRpRangeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRlRpRangeBag.Marshal(b, m, deterministic)
}
func (m *PimRlRpRangeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRlRpRangeBag.Merge(m, src)
}
func (m *PimRlRpRangeBag) XXX_Size() int {
	return xxx_messageInfo_PimRlRpRangeBag.Size(m)
}
func (m *PimRlRpRangeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRlRpRangeBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRlRpRangeBag proto.InternalMessageInfo

func (m *PimRlRpRangeBag) GetPrefix() *PimAddrtype {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PimRlRpRangeBag) GetPrefixLength() int32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *PimRlRpRangeBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *PimRlRpRangeBag) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *PimRlRpRangeBag) GetSourceOfInformation() *PimAddrtype {
	if m != nil {
		return m.SourceOfInformation
	}
	return nil
}

type PimRlBag struct {
	RpAddressXr          *PimAddrtype       `protobuf:"bytes,50,opt,name=rp_address_xr,json=rpAddressXr,proto3" json:"rp_address_xr,omitempty"`
	Protocol             string             `protobuf:"bytes,51,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ClientXr             string             `protobuf:"bytes,52,opt,name=client_xr,json=clientXr,proto3" json:"client_xr,omitempty"`
	SourceOfInformation  *PimAddrtype       `protobuf:"bytes,53,opt,name=source_of_information,json=sourceOfInformation,proto3" json:"source_of_information,omitempty"`
	Expires              uint64             `protobuf:"varint,54,opt,name=expires,proto3" json:"expires,omitempty"`
	GroupRange           []*PimRlRpRangeBag `protobuf:"bytes,55,rep,name=group_range,json=groupRange,proto3" json:"group_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PimRlBag) Reset()         { *m = PimRlBag{} }
func (m *PimRlBag) String() string { return proto.CompactTextString(m) }
func (*PimRlBag) ProtoMessage()    {}
func (*PimRlBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6a7cdc6ebd16448, []int{3}
}

func (m *PimRlBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRlBag.Unmarshal(m, b)
}
func (m *PimRlBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRlBag.Marshal(b, m, deterministic)
}
func (m *PimRlBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRlBag.Merge(m, src)
}
func (m *PimRlBag) XXX_Size() int {
	return xxx_messageInfo_PimRlBag.Size(m)
}
func (m *PimRlBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRlBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRlBag proto.InternalMessageInfo

func (m *PimRlBag) GetRpAddressXr() *PimAddrtype {
	if m != nil {
		return m.RpAddressXr
	}
	return nil
}

func (m *PimRlBag) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PimRlBag) GetClientXr() string {
	if m != nil {
		return m.ClientXr
	}
	return ""
}

func (m *PimRlBag) GetSourceOfInformation() *PimAddrtype {
	if m != nil {
		return m.SourceOfInformation
	}
	return nil
}

func (m *PimRlBag) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *PimRlBag) GetGroupRange() []*PimRlRpRangeBag {
	if m != nil {
		return m.GroupRange
	}
	return nil
}

func init() {
	proto.RegisterType((*PimRlBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.ranges.range.pim_rl_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.ranges.range.pim_addrtype")
	proto.RegisterType((*PimRlRpRangeBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.ranges.range.pim_rl_rp_range_bag")
	proto.RegisterType((*PimRlBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.ranges.range.pim_rl_bag")
}

func init() { proto.RegisterFile("pim_rl_bag.proto", fileDescriptor_b6a7cdc6ebd16448) }

var fileDescriptor_b6a7cdc6ebd16448 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xd5, 0x76, 0xcb, 0x96, 0x9d, 0x6c, 0x05, 0xb8, 0x02, 0x2c, 0x10, 0x52, 0x59, 0x2e, 0x3d,
	0xe5, 0xb0, 0x2d, 0xcb, 0x99, 0x03, 0x52, 0x11, 0x08, 0xa4, 0x20, 0xa1, 0x72, 0x40, 0x96, 0x37,
	0x71, 0x82, 0xa5, 0xc4, 0xb6, 0xc6, 0x4e, 0x95, 0x72, 0xe0, 0x03, 0xf8, 0x52, 0xf8, 0x0b, 0x64,
	0x3b, 0x9b, 0x14, 0x0e, 0x9c, 0xba, 0x3d, 0x25, 0x6f, 0x66, 0x3c, 0x6f, 0xde, 0xf8, 0x19, 0xee,
	0x1b, 0xd9, 0x30, 0xac, 0xd9, 0x86, 0x57, 0xa9, 0x41, 0xed, 0x34, 0x39, 0xcf, 0xa5, 0xcd, 0x35,
	0x93, 0xda, 0xb2, 0x0e, 0x99, 0x34, 0x97, 0x67, 0xcc, 0xd7, 0x68, 0x23, 0x30, 0x95, 0xe6, 0x72,
	0xed, 0x51, 0x6a, 0x1d, 0x57, 0xc5, 0xe6, 0x2a, 0x2d, 0x44, 0xc9, 0xdb, 0xda, 0xb1, 0x5c, 0x2b,
	0x27, 0x3a, 0x97, 0x22, 0x57, 0x95, 0xb0, 0xf1, 0xb3, 0x3c, 0x87, 0x7b, 0x63, 0x77, 0xf6, 0xee,
	0xcd, 0x97, 0x4f, 0xe4, 0x19, 0x00, 0x1a, 0xc6, 0x8b, 0x02, 0x85, 0xb5, 0x74, 0x72, 0x3c, 0x39,
	0x99, 0x67, 0x73, 0x34, 0xaf, 0x63, 0x80, 0x3c, 0x82, 0x59, 0x5e, 0x4b, 0xa1, 0x1c, 0xdd, 0x0b,
	0xa9, 0x1e, 0x2d, 0x1b, 0x58, 0xf8, 0x4e, 0xfe, 0x9c, 0xbb, 0x32, 0x82, 0x3c, 0x86, 0x03, 0x5e,
	0x32, 0xc5, 0x1b, 0xd1, 0xf7, 0x98, 0xf1, 0xf2, 0x03, 0x6f, 0x04, 0x79, 0x0e, 0x8b, 0x30, 0xf1,
	0x96, 0x21, 0xb6, 0x49, 0x7c, 0x6c, 0xcb, 0x11, 0x4b, 0xd6, 0x43, 0xc9, 0x74, 0x28, 0x59, 0xf7,
	0x25, 0xcb, 0xdf, 0x7b, 0x70, 0xd4, 0x4f, 0x8e, 0x86, 0x05, 0x31, 0x5e, 0x02, 0x51, 0x30, 0x33,
	0x28, 0x4a, 0xd9, 0x05, 0xd6, 0x64, 0xf5, 0x39, 0xbd, 0xa9, 0x5d, 0xa5, 0xd7, 0xe5, 0x65, 0x3d,
	0x0b, 0x79, 0x01, 0x87, 0xf1, 0x8f, 0xd5, 0x42, 0x55, 0xee, 0x5b, 0x90, 0xf3, 0x20, 0x5b, 0xc4,
	0xe0, 0xfb, 0x10, 0xf3, 0x3b, 0x6b, 0x8d, 0x93, 0x8d, 0x08, 0x4a, 0xf6, 0xb3, 0x1e, 0x11, 0x0a,
	0x07, 0xa2, 0x33, 0x12, 0x85, 0xa5, 0xfb, 0x21, 0xb1, 0x85, 0xe4, 0xe7, 0x04, 0x1e, 0x5a, 0xdd,
	0x62, 0x2e, 0x98, 0x2e, 0x99, 0x54, 0xa5, 0xc6, 0x86, 0x3b, 0xa9, 0x15, 0xbd, 0xb3, 0x53, 0x59,
	0x47, 0x91, 0xf4, 0x63, 0xf9, 0x76, 0xa4, 0x5c, 0xfe, 0x9a, 0x02, 0x8c, 0x2e, 0x21, 0xdf, 0xe1,
	0x70, 0x34, 0x08, 0xeb, 0x90, 0xae, 0x76, 0x3a, 0x52, 0x32, 0x78, 0xef, 0x02, 0xc9, 0x13, 0xb8,
	0x1b, 0x9e, 0x40, 0xae, 0x6b, 0x7a, 0x1a, 0x5c, 0x31, 0x60, 0xf2, 0x14, 0xe6, 0xd1, 0x8b, 0x7e,
	0xa6, 0xb3, 0x98, 0x8c, 0x81, 0x0b, 0xfc, 0xcf, 0x42, 0x5f, 0xde, 0xfa, 0x42, 0xaf, 0xdf, 0xfb,
	0xfa, 0xef, 0x7b, 0xff, 0x01, 0x49, 0x85, 0xba, 0xed, 0x1d, 0x4d, 0x5f, 0x1d, 0x4f, 0x4f, 0x92,
	0xd5, 0xd7, 0x9b, 0x9d, 0xed, 0x9f, 0x27, 0x93, 0x41, 0x60, 0xcc, 0x3c, 0xde, 0xcc, 0xc2, 0x36,
	0x4f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x75, 0x26, 0x1f, 0x74, 0x04, 0x00, 0x00,
}
