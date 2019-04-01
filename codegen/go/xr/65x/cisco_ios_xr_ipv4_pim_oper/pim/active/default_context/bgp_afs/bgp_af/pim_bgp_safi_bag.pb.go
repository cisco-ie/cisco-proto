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
// source: pim_bgp_safi_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_bgp_afs_bgp_af

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

type PimBgpSafiBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         string   `protobuf:"bytes,2,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBgpSafiBag_KEYS) Reset()         { *m = PimBgpSafiBag_KEYS{} }
func (m *PimBgpSafiBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBgpSafiBag_KEYS) ProtoMessage()    {}
func (*PimBgpSafiBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02a26094c211637, []int{0}
}

func (m *PimBgpSafiBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBgpSafiBag_KEYS.Unmarshal(m, b)
}
func (m *PimBgpSafiBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBgpSafiBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBgpSafiBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBgpSafiBag_KEYS.Merge(m, src)
}
func (m *PimBgpSafiBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBgpSafiBag_KEYS.Size(m)
}
func (m *PimBgpSafiBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBgpSafiBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBgpSafiBag_KEYS proto.InternalMessageInfo

func (m *PimBgpSafiBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimBgpSafiBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
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
	return fileDescriptor_a02a26094c211637, []int{1}
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

type PimBgpSafiBag struct {
	RouteDistinguisher   string       `protobuf:"bytes,50,opt,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	Source               *PimAddrtype `protobuf:"bytes,51,opt,name=source,proto3" json:"source,omitempty"`
	Group                *PimAddrtype `protobuf:"bytes,52,opt,name=group,proto3" json:"group,omitempty"`
	NextHop              *PimAddrtype `protobuf:"bytes,53,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	ExtranetPathCount    uint32       `protobuf:"varint,54,opt,name=extranet_path_count,json=extranetPathCount,proto3" json:"extranet_path_count,omitempty"`
	IsBgpAdded           bool         `protobuf:"varint,55,opt,name=is_bgp_added,json=isBgpAdded,proto3" json:"is_bgp_added,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimBgpSafiBag) Reset()         { *m = PimBgpSafiBag{} }
func (m *PimBgpSafiBag) String() string { return proto.CompactTextString(m) }
func (*PimBgpSafiBag) ProtoMessage()    {}
func (*PimBgpSafiBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a02a26094c211637, []int{2}
}

func (m *PimBgpSafiBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBgpSafiBag.Unmarshal(m, b)
}
func (m *PimBgpSafiBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBgpSafiBag.Marshal(b, m, deterministic)
}
func (m *PimBgpSafiBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBgpSafiBag.Merge(m, src)
}
func (m *PimBgpSafiBag) XXX_Size() int {
	return xxx_messageInfo_PimBgpSafiBag.Size(m)
}
func (m *PimBgpSafiBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBgpSafiBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBgpSafiBag proto.InternalMessageInfo

func (m *PimBgpSafiBag) GetRouteDistinguisher() string {
	if m != nil {
		return m.RouteDistinguisher
	}
	return ""
}

func (m *PimBgpSafiBag) GetSource() *PimAddrtype {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PimBgpSafiBag) GetGroup() *PimAddrtype {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *PimBgpSafiBag) GetNextHop() *PimAddrtype {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *PimBgpSafiBag) GetExtranetPathCount() uint32 {
	if m != nil {
		return m.ExtranetPathCount
	}
	return 0
}

func (m *PimBgpSafiBag) GetIsBgpAdded() bool {
	if m != nil {
		return m.IsBgpAdded
	}
	return false
}

func init() {
	proto.RegisterType((*PimBgpSafiBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bgp_afs.bgp_af.pim_bgp_safi_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bgp_afs.bgp_af.pim_addrtype")
	proto.RegisterType((*PimBgpSafiBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bgp_afs.bgp_af.pim_bgp_safi_bag")
}

func init() { proto.RegisterFile("pim_bgp_safi_bag.proto", fileDescriptor_a02a26094c211637) }

var fileDescriptor_a02a26094c211637 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xf1, 0xc2, 0x92, 0x4c, 0x71, 0xc6, 0xa6, 0xb0, 0xcd, 0x47, 0x2f, 0x63, 0xe0, 0x93,
	0x06, 0x49, 0x96, 0x9d, 0xd7, 0xa6, 0x50, 0x28, 0x94, 0xe2, 0x9e, 0x7a, 0x12, 0x8a, 0xf5, 0x6c,
	0xab, 0xd4, 0x96, 0x90, 0xe4, 0xe0, 0x7e, 0xd5, 0x7e, 0x9a, 0x22, 0xd9, 0x09, 0x69, 0xce, 0xcd,
	0xc9, 0xf0, 0xde, 0xcf, 0xff, 0x1f, 0xef, 0xe9, 0xa1, 0xef, 0x4a, 0x54, 0x74, 0x5b, 0x28, 0x6a,
	0x58, 0x2e, 0xe8, 0x96, 0x15, 0x44, 0x69, 0x69, 0x25, 0xde, 0x64, 0xc2, 0x64, 0x92, 0x0a, 0x69,
	0x68, 0xab, 0xa9, 0x50, 0xbb, 0x15, 0x75, 0xa4, 0x54, 0xa0, 0x89, 0x12, 0x15, 0x61, 0x99, 0x15,
	0x3b, 0x20, 0x1c, 0x72, 0xd6, 0x3c, 0x59, 0x9a, 0xc9, 0xda, 0x42, 0x6b, 0x89, 0x4b, 0x62, 0xb9,
	0xe9, 0xbf, 0xf3, 0x0c, 0x7d, 0x3b, 0xcd, 0xa7, 0x37, 0x57, 0x0f, 0xf7, 0xf8, 0x37, 0xfa, 0x6c,
	0x64, 0xa3, 0x33, 0xa0, 0x8c, 0x73, 0x0d, 0xc6, 0x44, 0x41, 0x1c, 0x24, 0x9f, 0xd2, 0x69, 0x57,
	0xfd, 0xdf, 0x15, 0xf1, 0x2f, 0x34, 0x2d, 0xb4, 0x6c, 0xd4, 0x81, 0xfa, 0xe0, 0xa9, 0xd0, 0x17,
	0x7b, 0x68, 0x5e, 0xa1, 0xd0, 0x49, 0x1c, 0x62, 0x9f, 0x15, 0xe0, 0x1f, 0x68, 0xc4, 0x72, 0x5a,
	0xb3, 0x0a, 0xfa, 0xd0, 0x21, 0xcb, 0x6f, 0x59, 0x05, 0xf8, 0x27, 0x0a, 0xfd, 0x20, 0x6f, 0xc3,
	0x26, 0xae, 0xb6, 0x17, 0x76, 0xc8, 0xfa, 0x80, 0x0c, 0x0e, 0xc8, 0x7a, 0xaf, 0x7b, 0x19, 0xa0,
	0x2f, 0xa7, 0x43, 0xe1, 0x3f, 0x68, 0xa6, 0x65, 0x63, 0x81, 0x72, 0x61, 0xac, 0xa8, 0x8b, 0x46,
	0x98, 0x12, 0x74, 0xb4, 0xf0, 0xbf, 0x63, 0xdf, 0xda, 0x1c, 0x77, 0xf0, 0x23, 0x1a, 0x76, 0xa3,
	0x46, 0xcb, 0x38, 0x48, 0x26, 0x8b, 0x94, 0xbc, 0xc7, 0xc2, 0xc9, 0xf1, 0x22, 0xd2, 0xde, 0x80,
	0x4b, 0xf4, 0xd1, 0x2f, 0x2c, 0x5a, 0x9d, 0x4d, 0xd5, 0x09, 0x70, 0x85, 0xc6, 0x35, 0xb4, 0x96,
	0x96, 0x52, 0x45, 0x7f, 0xcf, 0x26, 0x1b, 0x39, 0xc7, 0xb5, 0x54, 0x98, 0xa0, 0x19, 0xb4, 0x56,
	0xb3, 0x1a, 0x2c, 0x55, 0xcc, 0x96, 0x34, 0x93, 0x4d, 0x6d, 0xa3, 0x75, 0x1c, 0x24, 0xd3, 0xf4,
	0xeb, 0xbe, 0x75, 0xc7, 0x6c, 0x79, 0xe9, 0x1a, 0x38, 0x46, 0xa1, 0x30, 0xfe, 0xe1, 0x18, 0xe7,
	0xc0, 0xa3, 0x7f, 0x71, 0x90, 0x8c, 0x53, 0x24, 0xcc, 0x45, 0xe1, 0xae, 0x09, 0xf8, 0x76, 0xe8,
	0xaf, 0x7f, 0xf9, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xa5, 0xdc, 0xb1, 0x17, 0x03, 0x00, 0x00,
}