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
// source: pim_mstatic_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_multicast_static_routes_multicast_static_route

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

type PimMstaticBag_KEYS struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimMstaticBag_KEYS) Reset()         { *m = PimMstaticBag_KEYS{} }
func (m *PimMstaticBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimMstaticBag_KEYS) ProtoMessage()    {}
func (*PimMstaticBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7fc2c56e20ecff, []int{0}
}

func (m *PimMstaticBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMstaticBag_KEYS.Unmarshal(m, b)
}
func (m *PimMstaticBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMstaticBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimMstaticBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMstaticBag_KEYS.Merge(m, src)
}
func (m *PimMstaticBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimMstaticBag_KEYS.Size(m)
}
func (m *PimMstaticBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMstaticBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimMstaticBag_KEYS proto.InternalMessageInfo

func (m *PimMstaticBag_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PimMstaticBag_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
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
	return fileDescriptor_cc7fc2c56e20ecff, []int{1}
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

type PimMstaticBag struct {
	Prefix               *PimAddrtype `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Nexthop              *PimAddrtype `protobuf:"bytes,51,opt,name=nexthop,proto3" json:"nexthop,omitempty"`
	InterfaceName        string       `protobuf:"bytes,52,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Distance             uint32       `protobuf:"varint,53,opt,name=distance,proto3" json:"distance,omitempty"`
	PrefixLengthXr       uint32       `protobuf:"varint,54,opt,name=prefix_length_xr,json=prefixLengthXr,proto3" json:"prefix_length_xr,omitempty"`
	IsViaLsm             bool         `protobuf:"varint,55,opt,name=is_via_lsm,json=isViaLsm,proto3" json:"is_via_lsm,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimMstaticBag) Reset()         { *m = PimMstaticBag{} }
func (m *PimMstaticBag) String() string { return proto.CompactTextString(m) }
func (*PimMstaticBag) ProtoMessage()    {}
func (*PimMstaticBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc7fc2c56e20ecff, []int{2}
}

func (m *PimMstaticBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMstaticBag.Unmarshal(m, b)
}
func (m *PimMstaticBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMstaticBag.Marshal(b, m, deterministic)
}
func (m *PimMstaticBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMstaticBag.Merge(m, src)
}
func (m *PimMstaticBag) XXX_Size() int {
	return xxx_messageInfo_PimMstaticBag.Size(m)
}
func (m *PimMstaticBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMstaticBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimMstaticBag proto.InternalMessageInfo

func (m *PimMstaticBag) GetPrefix() *PimAddrtype {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PimMstaticBag) GetNexthop() *PimAddrtype {
	if m != nil {
		return m.Nexthop
	}
	return nil
}

func (m *PimMstaticBag) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PimMstaticBag) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *PimMstaticBag) GetPrefixLengthXr() uint32 {
	if m != nil {
		return m.PrefixLengthXr
	}
	return 0
}

func (m *PimMstaticBag) GetIsViaLsm() bool {
	if m != nil {
		return m.IsViaLsm
	}
	return false
}

func init() {
	proto.RegisterType((*PimMstaticBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.multicast_static_routes.multicast_static_route.pim_mstatic_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.multicast_static_routes.multicast_static_route.pim_addrtype")
	proto.RegisterType((*PimMstaticBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.multicast_static_routes.multicast_static_route.pim_mstatic_bag")
}

func init() { proto.RegisterFile("pim_mstatic_bag.proto", fileDescriptor_cc7fc2c56e20ecff) }

var fileDescriptor_cc7fc2c56e20ecff = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcd, 0x6a, 0xe3, 0x30,
	0x10, 0xc7, 0x71, 0x02, 0x49, 0x56, 0xf9, 0xd8, 0x45, 0xec, 0xb2, 0x62, 0xd9, 0x83, 0x37, 0x4b,
	0xc1, 0x27, 0x1f, 0x92, 0x34, 0x3d, 0xf7, 0xd0, 0x53, 0x43, 0x0f, 0x2e, 0x2d, 0xed, 0x49, 0x28,
	0xb2, 0x9c, 0x08, 0x2c, 0x4b, 0x48, 0x93, 0xe0, 0xbe, 0x42, 0xe9, 0x03, 0xf5, 0xf1, 0x8a, 0x65,
	0xc7, 0xa4, 0x81, 0x5e, 0xdb, 0xe3, 0xfc, 0xf5, 0xd3, 0xcc, 0x7f, 0x3e, 0xd0, 0x2f, 0x23, 0x15,
	0x55, 0x0e, 0x18, 0x48, 0x4e, 0xd7, 0x6c, 0x13, 0x1b, 0xab, 0x41, 0xe3, 0x94, 0x4b, 0xc7, 0x35,
	0x95, 0xda, 0xd1, 0xd2, 0x52, 0x69, 0xf6, 0x0b, 0x5a, 0x81, 0xda, 0x08, 0x1b, 0x1b, 0xa9, 0x62,
	0xc6, 0x41, 0xee, 0x45, 0x9c, 0x8a, 0x8c, 0xed, 0x72, 0xa0, 0x5c, 0x17, 0x20, 0x4a, 0x88, 0xd5,
	0x2e, 0x07, 0xc9, 0x99, 0x03, 0xda, 0x64, 0xb4, 0x7a, 0x07, 0xc2, 0x7d, 0xa0, 0x4f, 0xef, 0xd0,
	0xcf, 0x93, 0xf2, 0xf4, 0xfa, 0xea, 0xf1, 0x16, 0x13, 0xd4, 0x67, 0x69, 0x6a, 0x85, 0x73, 0x24,
	0x08, 0x83, 0xe8, 0x5b, 0x72, 0x08, 0xf1, 0x7f, 0x34, 0x36, 0x56, 0x64, 0xb2, 0xa4, 0xb9, 0x28,
	0x36, 0xb0, 0x25, 0x9d, 0x30, 0x88, 0xc6, 0xc9, 0xa8, 0x16, 0x57, 0x5e, 0x9b, 0x2a, 0x34, 0xaa,
	0xd2, 0x56, 0x7f, 0xe0, 0xc9, 0x08, 0xfc, 0x1b, 0xf5, 0x59, 0x46, 0x0b, 0xa6, 0x44, 0x93, 0xae,
	0xc7, 0xb2, 0x1b, 0xa6, 0x04, 0xfe, 0x87, 0x46, 0xbe, 0xb5, 0x43, 0xb1, 0x8e, 0x7f, 0x1d, 0x56,
	0xda, 0x65, 0x53, 0xb0, 0x46, 0x96, 0x2d, 0xd2, 0x6d, 0x91, 0x65, 0x83, 0x4c, 0x5f, 0xbb, 0xe8,
	0xfb, 0x49, 0x1b, 0xf8, 0x39, 0x40, 0xbd, 0xda, 0x13, 0x99, 0x85, 0x41, 0x34, 0x9c, 0xd9, 0xf8,
	0x33, 0x26, 0x1a, 0x1f, 0xf7, 0x9d, 0x34, 0x0e, 0xf0, 0x4b, 0x80, 0xfa, 0x85, 0x28, 0x61, 0xab,
	0x0d, 0x99, 0x7f, 0x99, 0x9b, 0x83, 0x05, 0x7c, 0x86, 0x26, 0xb2, 0x00, 0x61, 0x33, 0xc6, 0x45,
	0xbd, 0x95, 0x85, 0x1f, 0xea, 0xb8, 0x55, 0xfd, 0x72, 0xfe, 0xa0, 0x41, 0x2a, 0x1d, 0xb0, 0x82,
	0x0b, 0x72, 0xee, 0xb7, 0xdc, 0xc6, 0x38, 0x42, 0x3f, 0xde, 0x9d, 0x01, 0x2d, 0x2d, 0x59, 0x7a,
	0x66, 0x72, 0x7c, 0x09, 0x0f, 0x16, 0xff, 0x45, 0x48, 0x3a, 0xba, 0x97, 0x8c, 0xe6, 0x4e, 0x91,
	0x8b, 0x30, 0x88, 0x06, 0xc9, 0x40, 0xba, 0x7b, 0xc9, 0x56, 0x4e, 0xad, 0x7b, 0xfe, 0xda, 0xe7,
	0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xab, 0x3d, 0x86, 0xe0, 0x06, 0x03, 0x00, 0x00,
}
