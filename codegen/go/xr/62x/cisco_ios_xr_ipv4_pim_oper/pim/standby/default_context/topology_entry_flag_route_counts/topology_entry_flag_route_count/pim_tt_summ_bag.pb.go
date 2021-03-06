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
// source: pim_tt_summ_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_topology_entry_flag_route_counts_topology_entry_flag_route_count

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

type PimTtSummBag_KEYS struct {
	EntryFlag            string   `protobuf:"bytes,1,opt,name=entry_flag,json=entryFlag,proto3" json:"entry_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimTtSummBag_KEYS) Reset()         { *m = PimTtSummBag_KEYS{} }
func (m *PimTtSummBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimTtSummBag_KEYS) ProtoMessage()    {}
func (*PimTtSummBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc46104ddf9a9167, []int{0}
}

func (m *PimTtSummBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTtSummBag_KEYS.Unmarshal(m, b)
}
func (m *PimTtSummBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTtSummBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimTtSummBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTtSummBag_KEYS.Merge(m, src)
}
func (m *PimTtSummBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimTtSummBag_KEYS.Size(m)
}
func (m *PimTtSummBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTtSummBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimTtSummBag_KEYS proto.InternalMessageInfo

func (m *PimTtSummBag_KEYS) GetEntryFlag() string {
	if m != nil {
		return m.EntryFlag
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
	return fileDescriptor_dc46104ddf9a9167, []int{1}
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

type PimTtSummBag struct {
	GroupRanges          uint32       `protobuf:"varint,50,opt,name=group_ranges,json=groupRanges,proto3" json:"group_ranges,omitempty"`
	ActiveGroupRanges    uint32       `protobuf:"varint,51,opt,name=active_group_ranges,json=activeGroupRanges,proto3" json:"active_group_ranges,omitempty"`
	GroupAddress         *PimAddrtype `protobuf:"bytes,52,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	GrouteCount          uint32       `protobuf:"varint,53,opt,name=groute_count,json=grouteCount,proto3" json:"groute_count,omitempty"`
	SgRouteCount         uint32       `protobuf:"varint,54,opt,name=sg_route_count,json=sgRouteCount,proto3" json:"sg_route_count,omitempty"`
	SgrRouteCount        uint32       `protobuf:"varint,55,opt,name=sgr_route_count,json=sgrRouteCount,proto3" json:"sgr_route_count,omitempty"`
	IsNodeLowMemory      bool         `protobuf:"varint,56,opt,name=is_node_low_memory,json=isNodeLowMemory,proto3" json:"is_node_low_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimTtSummBag) Reset()         { *m = PimTtSummBag{} }
func (m *PimTtSummBag) String() string { return proto.CompactTextString(m) }
func (*PimTtSummBag) ProtoMessage()    {}
func (*PimTtSummBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc46104ddf9a9167, []int{2}
}

func (m *PimTtSummBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTtSummBag.Unmarshal(m, b)
}
func (m *PimTtSummBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTtSummBag.Marshal(b, m, deterministic)
}
func (m *PimTtSummBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTtSummBag.Merge(m, src)
}
func (m *PimTtSummBag) XXX_Size() int {
	return xxx_messageInfo_PimTtSummBag.Size(m)
}
func (m *PimTtSummBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTtSummBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimTtSummBag proto.InternalMessageInfo

func (m *PimTtSummBag) GetGroupRanges() uint32 {
	if m != nil {
		return m.GroupRanges
	}
	return 0
}

func (m *PimTtSummBag) GetActiveGroupRanges() uint32 {
	if m != nil {
		return m.ActiveGroupRanges
	}
	return 0
}

func (m *PimTtSummBag) GetGroupAddress() *PimAddrtype {
	if m != nil {
		return m.GroupAddress
	}
	return nil
}

func (m *PimTtSummBag) GetGrouteCount() uint32 {
	if m != nil {
		return m.GrouteCount
	}
	return 0
}

func (m *PimTtSummBag) GetSgRouteCount() uint32 {
	if m != nil {
		return m.SgRouteCount
	}
	return 0
}

func (m *PimTtSummBag) GetSgrRouteCount() uint32 {
	if m != nil {
		return m.SgrRouteCount
	}
	return 0
}

func (m *PimTtSummBag) GetIsNodeLowMemory() bool {
	if m != nil {
		return m.IsNodeLowMemory
	}
	return false
}

func init() {
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_addrtype")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0xd2, 0x3d, 0xcf, 0xd3, 0x30,
	0x10, 0x07, 0x70, 0x85, 0x4a, 0x85, 0xba, 0x2d, 0x15, 0x01, 0x44, 0x16, 0xa4, 0x50, 0x21, 0x14,
	0x09, 0xc9, 0x43, 0xdf, 0x60, 0x45, 0x08, 0x18, 0x80, 0x0e, 0x61, 0x62, 0x3a, 0xb9, 0x89, 0x63,
	0x59, 0x8a, 0x73, 0x96, 0xed, 0xb4, 0xcd, 0x07, 0xe0, 0xd3, 0xb0, 0xf0, 0x11, 0x51, 0x9c, 0x3e,
	0x6d, 0xda, 0xe5, 0x19, 0x9f, 0xf5, 0x7f, 0xbf, 0x8b, 0xee, 0x2e, 0x26, 0x2f, 0xb5, 0x54, 0xe0,
	0x1c, 0xd8, 0x5a, 0x29, 0xd8, 0x31, 0x41, 0xb5, 0x41, 0x87, 0xe1, 0x21, 0x93, 0x36, 0x43, 0x90,
	0x68, 0xe1, 0x68, 0x40, 0xea, 0xfd, 0x0a, 0x5a, 0x88, 0x9a, 0x1b, 0xaa, 0xa5, 0xa2, 0xd6, 0xb1,
	0x2a, 0xdf, 0x35, 0x34, 0xe7, 0x05, 0xab, 0x4b, 0x07, 0x19, 0x56, 0x8e, 0x1f, 0x1d, 0x75, 0xa8,
	0xb1, 0x44, 0xd1, 0x00, 0xaf, 0x9c, 0x69, 0xa0, 0x28, 0x99, 0x00, 0x83, 0xb5, 0xe3, 0x90, 0x61,
	0x5d, 0x39, 0x7b, 0x1f, 0x98, 0xaf, 0xc9, 0x8b, 0x9b, 0x89, 0xe0, 0xfb, 0x97, 0xdf, 0xbf, 0xc2,
	0xd7, 0x84, 0x5c, 0x3a, 0xa2, 0x20, 0x0e, 0x92, 0x51, 0x3a, 0xf2, 0xc9, 0xd7, 0x92, 0x89, 0xb9,
	0x22, 0x93, 0xb6, 0x8d, 0xe5, 0xb9, 0x71, 0x8d, 0xe6, 0xe1, 0x2b, 0xf2, 0x98, 0x15, 0x50, 0x31,
	0xc5, 0x4f, 0x76, 0xc8, 0x8a, 0x2d, 0x53, 0x3c, 0x7c, 0x43, 0x26, 0x7e, 0x9b, 0x56, 0x72, 0x6b,
	0xa3, 0x47, 0xbe, 0x3a, 0x6e, 0xb3, 0x4f, 0x5d, 0x74, 0x22, 0x9b, 0x33, 0x19, 0x9c, 0xc9, 0xe6,
	0x44, 0xe6, 0xff, 0x06, 0x64, 0x76, 0x33, 0x66, 0xdb, 0x26, 0x0c, 0xd6, 0x1a, 0x0c, 0xab, 0x04,
	0xb7, 0xd1, 0x22, 0x0e, 0x92, 0x69, 0x3a, 0xf6, 0x59, 0xea, 0xa3, 0x90, 0x92, 0xe7, 0x2c, 0x73,
	0x72, 0xcf, 0xe1, 0x4a, 0x2e, 0xbd, 0x7c, 0xd6, 0x95, 0xbe, 0xf5, 0xfc, 0xdf, 0x80, 0x4c, 0x3b,
	0x79, 0x37, 0xcb, 0x2a, 0x0e, 0x92, 0xf1, 0xe2, 0x4f, 0x40, 0x1f, 0xe8, 0xff, 0xd0, 0xfe, 0x95,
	0xd3, 0x6e, 0xe1, 0xde, 0xdd, 0x44, 0x8f, 0x46, 0xeb, 0xcb, 0x01, 0x1c, 0xff, 0xdc, 0x46, 0xe1,
	0x5b, 0xf2, 0xd4, 0x5e, 0x7d, 0x2f, 0xda, 0x78, 0x34, 0xb1, 0x22, 0xbd, 0xa8, 0x77, 0x64, 0x66,
	0x85, 0xb9, 0x62, 0x1f, 0x3c, 0x9b, 0x5a, 0x61, 0x7a, 0xee, 0x3d, 0x09, 0xa5, 0x85, 0x0a, 0x73,
	0x0e, 0x25, 0x1e, 0x40, 0x71, 0x85, 0xa6, 0x89, 0x3e, 0xc6, 0x41, 0xf2, 0x24, 0x9d, 0x49, 0xbb,
	0xc5, 0x9c, 0xff, 0xc0, 0xc3, 0x4f, 0x1f, 0xef, 0x86, 0xfe, 0x61, 0x2f, 0xff, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x00, 0x65, 0xc7, 0xf1, 0x02, 0x00, 0x00,
}
