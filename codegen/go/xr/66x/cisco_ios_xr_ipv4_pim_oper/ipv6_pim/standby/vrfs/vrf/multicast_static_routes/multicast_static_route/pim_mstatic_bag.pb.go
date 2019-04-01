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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_vrfs_vrf_multicast_static_routes_multicast_static_route

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
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

func (m *PimMstaticBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimMstaticBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.multicast_static_routes.multicast_static_route.pim_mstatic_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.multicast_static_routes.multicast_static_route.pim_addrtype")
	proto.RegisterType((*PimMstaticBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.multicast_static_routes.multicast_static_route.pim_mstatic_bag")
}

func init() { proto.RegisterFile("pim_mstatic_bag.proto", fileDescriptor_cc7fc2c56e20ecff) }

var fileDescriptor_cc7fc2c56e20ecff = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x16, 0x9a, 0x38, 0xfd, 0x50, 0x16, 0xc5, 0x55, 0x3c, 0xc4, 0x8a, 0x90, 0x53, 0x0e,
	0x6d, 0x8d, 0x67, 0x0f, 0x9e, 0x2c, 0x1e, 0x22, 0x88, 0x9e, 0x96, 0x6d, 0xb2, 0x69, 0x17, 0xba,
	0xc9, 0xba, 0xbb, 0x0d, 0xe9, 0x4f, 0x10, 0xff, 0x8f, 0xbf, 0x4f, 0xf2, 0x61, 0x68, 0x0b, 0x5e,
	0xf5, 0x12, 0x98, 0x97, 0xb7, 0xfb, 0xde, 0xcc, 0x9b, 0x85, 0x53, 0xc9, 0x05, 0x11, 0xda, 0x50,
	0xc3, 0x23, 0x32, 0xa7, 0x0b, 0x5f, 0xaa, 0xcc, 0x64, 0x28, 0x8a, 0xb8, 0x8e, 0x32, 0xc2, 0x33,
	0x4d, 0x0a, 0x45, 0xb8, 0xcc, 0xa7, 0xa4, 0x24, 0x66, 0x92, 0x29, 0x9f, 0xcb, 0x3c, 0x28, 0x2b,
	0x5f, 0x1b, 0x9a, 0xc6, 0xf3, 0x8d, 0x9f, 0xab, 0x44, 0x97, 0x1f, 0x5f, 0xac, 0x57, 0x86, 0x47,
	0x54, 0x1b, 0xd2, 0x5c, 0xa8, 0xb2, 0xb5, 0x61, 0xfa, 0x17, 0x7c, 0x24, 0xe1, 0x64, 0x4f, 0x9d,
	0x3c, 0x3e, 0xbc, 0x3d, 0xa3, 0x73, 0x70, 0x72, 0x95, 0x90, 0x94, 0x0a, 0x86, 0x2d, 0xd7, 0xf2,
	0x0e, 0x43, 0x3b, 0x57, 0xc9, 0x13, 0x15, 0x0c, 0x61, 0xb0, 0x69, 0x1c, 0x2b, 0xa6, 0x35, 0x3e,
	0xa8, 0xff, 0x34, 0x25, 0xba, 0x86, 0x81, 0x54, 0x2c, 0xe1, 0x05, 0x59, 0xb1, 0x74, 0x61, 0x96,
	0xb8, 0xe3, 0x5a, 0xde, 0x20, 0xec, 0xd7, 0xe0, 0xac, 0xc2, 0x46, 0x02, 0xfa, 0xa5, 0x62, 0x79,
	0xc6, 0x6c, 0x24, 0x43, 0x67, 0x60, 0xd3, 0x1d, 0xa1, 0x2e, 0xad, 0x75, 0xae, 0xa0, 0x5f, 0x35,
	0xbd, 0x2b, 0xd6, 0x2b, 0xb1, 0xfb, 0x46, 0xb0, 0xa6, 0x04, 0x2d, 0xa5, 0xd3, 0x52, 0x82, 0x86,
	0x32, 0xfa, 0xea, 0xc0, 0xd1, 0x5e, 0x87, 0xe8, 0xc3, 0x82, 0x6e, 0xed, 0x09, 0x8f, 0x5d, 0xcb,
	0xeb, 0x8d, 0xdf, 0xfd, 0x3f, 0x98, 0xb5, 0xbf, 0xdd, 0x76, 0xd8, 0x18, 0x40, 0x9f, 0x16, 0xd8,
	0x29, 0x2b, 0xcc, 0x32, 0x93, 0x78, 0xf2, 0x5f, 0x66, 0x7e, 0x1c, 0xa0, 0x1b, 0x18, 0xf2, 0xd4,
	0x30, 0x95, 0xd0, 0x88, 0xd5, 0x99, 0x4c, 0xab, 0x91, 0x0e, 0x5a, 0xb4, 0x8a, 0xe6, 0x02, 0x9c,
	0x98, 0x97, 0x0e, 0x22, 0x86, 0x6f, 0xab, 0x8c, 0xdb, 0x1a, 0x79, 0x70, 0xbc, 0xb3, 0x04, 0xa4,
	0x50, 0x38, 0xa8, 0x38, 0xc3, 0xed, 0x3d, 0x78, 0x55, 0xe8, 0x12, 0x80, 0x6b, 0x92, 0x73, 0x4a,
	0x56, 0x5a, 0xe0, 0x3b, 0xd7, 0xf2, 0x9c, 0xd0, 0xe1, 0xfa, 0x85, 0xd3, 0x99, 0x16, 0xf3, 0x6e,
	0xf5, 0x0a, 0x26, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0xa1, 0x5b, 0xc4, 0x1e, 0x03, 0x00,
	0x00,
}