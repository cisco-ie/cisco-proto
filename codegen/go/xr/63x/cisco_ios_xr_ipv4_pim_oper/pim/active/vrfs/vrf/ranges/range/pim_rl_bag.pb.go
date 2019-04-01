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

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_ranges_range

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RpAddress            string   `protobuf:"bytes,2,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	Client               string   `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
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

func (m *PimRlBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimRlBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.ranges.range.pim_rl_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.ranges.range.pim_addrtype")
	proto.RegisterType((*PimRlRpRangeBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.ranges.range.pim_rl_rp_range_bag")
	proto.RegisterType((*PimRlBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.ranges.range.pim_rl_bag")
}

func init() { proto.RegisterFile("pim_rl_bag.proto", fileDescriptor_b6a7cdc6ebd16448) }

var fileDescriptor_b6a7cdc6ebd16448 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0x9b, 0x92, 0xed, 0x4e, 0xb6, 0x02, 0x5c, 0x01, 0x06, 0x84, 0x54, 0xc2, 0xa5, 0xa7,
	0x1c, 0xb6, 0x65, 0x39, 0x70, 0xe2, 0xc0, 0xa1, 0x02, 0x01, 0x0a, 0x97, 0x72, 0xb2, 0xdc, 0xd4,
	0x09, 0x96, 0xe2, 0xd8, 0x8c, 0xb3, 0x51, 0x38, 0xf0, 0xdf, 0xf8, 0x05, 0xfc, 0x26, 0x64, 0x3b,
	0x4d, 0xba, 0x48, 0x9c, 0x76, 0xd5, 0x4b, 0x92, 0xf9, 0xc8, 0xbc, 0x79, 0x6f, 0x1e, 0x3c, 0x30,
	0x52, 0x31, 0xac, 0xd9, 0x15, 0xaf, 0x32, 0x83, 0xba, 0xd5, 0xe4, 0x6d, 0x21, 0x6d, 0xa1, 0x99,
	0xd4, 0x96, 0xf5, 0xc8, 0xa4, 0xe9, 0xce, 0x99, 0xeb, 0xd1, 0x46, 0x60, 0x66, 0xa4, 0xca, 0x78,
	0xd1, 0xca, 0x4e, 0x64, 0x1d, 0x96, 0xd6, 0x3d, 0x32, 0xe4, 0x4d, 0x25, 0x6c, 0x78, 0xa5, 0x05,
	0xdc, 0x9f, 0x06, 0xb2, 0x0f, 0xef, 0xbf, 0x7d, 0x25, 0x4f, 0xe1, 0xb0, 0xc3, 0x92, 0x35, 0x5c,
	0x09, 0xba, 0x77, 0xb2, 0x77, 0x3a, 0xcf, 0x67, 0x1d, 0x96, 0x9f, 0xb8, 0x12, 0xe4, 0x05, 0x00,
	0x1a, 0xc6, 0xaf, 0xaf, 0x51, 0x58, 0x4b, 0xf7, 0x7d, 0x71, 0x8e, 0xe6, 0x5d, 0x48, 0x90, 0xc7,
	0x10, 0x17, 0xb5, 0x14, 0x4d, 0x4b, 0x23, 0x5f, 0x1a, 0xa2, 0x54, 0xc1, 0xc2, 0x81, 0xb8, 0xff,
	0xda, 0x9f, 0x46, 0x90, 0x27, 0x30, 0xe3, 0x1b, 0x00, 0x31, 0x0f, 0xf3, 0x5f, 0xc2, 0xc2, 0xef,
	0xbf, 0x89, 0x90, 0xb8, 0xdc, 0x0d, 0x46, 0x68, 0x59, 0x8d, 0x2d, 0xd1, 0xd8, 0xb2, 0x1a, 0x5a,
	0xd2, 0x3f, 0xfb, 0x70, 0x3c, 0x90, 0x42, 0xc3, 0x3c, 0x4f, 0xc7, 0x8e, 0x70, 0x88, 0x0d, 0x8a,
	0x52, 0xf6, 0x1e, 0x35, 0x59, 0x5e, 0x64, 0x5b, 0x28, 0x97, 0xdd, 0x66, 0x94, 0x0f, 0x83, 0xc9,
	0x2b, 0x38, 0x0a, 0x5f, 0xac, 0x16, 0x4d, 0xd5, 0x7e, 0xf7, 0x0c, 0x1e, 0xe6, 0x8b, 0x90, 0xfc,
	0xe8, 0x73, 0x4e, 0xa6, 0xb5, 0x69, 0xa5, 0x12, 0x7e, 0xf9, 0x83, 0x7c, 0x88, 0x08, 0x85, 0x99,
	0xe8, 0x8d, 0x44, 0x61, 0xe9, 0x81, 0x2f, 0xdc, 0x84, 0xe4, 0x17, 0x3c, 0xb2, 0x7a, 0x8d, 0x85,
	0x60, 0xba, 0x64, 0xb2, 0x29, 0x35, 0x2a, 0xde, 0x4a, 0xdd, 0xd0, 0x7b, 0xbb, 0x26, 0x72, 0x1c,
	0x70, 0x3e, 0x97, 0x17, 0x13, 0x4a, 0xfa, 0x3b, 0x02, 0x98, 0x5c, 0x42, 0x14, 0x1c, 0x4d, 0x2e,
	0x60, 0x3d, 0xd2, 0xe5, 0xae, 0xb7, 0x48, 0x46, 0x4f, 0x5d, 0x22, 0x79, 0x06, 0x87, 0xde, 0xe8,
	0x85, 0xae, 0xe9, 0x99, 0xbf, 0xf6, 0x18, 0x93, 0xe7, 0x30, 0x0f, 0x1e, 0x73, 0x6b, 0x9c, 0x87,
	0x62, 0x48, 0x5c, 0xe2, 0xff, 0x55, 0x7b, 0x7d, 0x17, 0xaa, 0xdd, 0x3e, 0xe7, 0x6a, 0xf3, 0x9c,
	0x3f, 0x20, 0xa9, 0x50, 0xaf, 0x07, 0x6f, 0xd2, 0x37, 0x27, 0xd1, 0x69, 0xb2, 0xfc, 0xb2, 0xf5,
	0x3a, 0xff, 0xf8, 0x3d, 0x07, 0x0f, 0x92, 0xbb, 0xf8, 0x2a, 0xf6, 0x92, 0x9d, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x7a, 0xa9, 0xc7, 0x74, 0x3f, 0x04, 0x00, 0x00,
}
