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

type PimTtSummBag struct {
	GroupRanges          uint32   `protobuf:"varint,50,opt,name=group_ranges,json=groupRanges,proto3" json:"group_ranges,omitempty"`
	ActiveGroupRanges    uint32   `protobuf:"varint,51,opt,name=active_group_ranges,json=activeGroupRanges,proto3" json:"active_group_ranges,omitempty"`
	GrouteCount          uint32   `protobuf:"varint,52,opt,name=groute_count,json=grouteCount,proto3" json:"groute_count,omitempty"`
	SgRouteCount         uint32   `protobuf:"varint,53,opt,name=sg_route_count,json=sgRouteCount,proto3" json:"sg_route_count,omitempty"`
	SgrRouteCount        uint32   `protobuf:"varint,54,opt,name=sgr_route_count,json=sgrRouteCount,proto3" json:"sgr_route_count,omitempty"`
	IsNodeLowMemory      bool     `protobuf:"varint,55,opt,name=is_node_low_memory,json=isNodeLowMemory,proto3" json:"is_node_low_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimTtSummBag) Reset()         { *m = PimTtSummBag{} }
func (m *PimTtSummBag) String() string { return proto.CompactTextString(m) }
func (*PimTtSummBag) ProtoMessage()    {}
func (*PimTtSummBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc46104ddf9a9167, []int{1}
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
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0xc0, 0x71, 0xf2, 0x1c, 0x1e, 0xec, 0xda, 0x5a, 0x8c, 0x0a, 0xb9, 0x08, 0xb5, 0x88, 0x14,
	0x84, 0x1c, 0x6c, 0xab, 0x1f, 0x40, 0xd4, 0x83, 0x2f, 0x87, 0x78, 0xf2, 0x34, 0x6c, 0x93, 0xed,
	0xb2, 0x90, 0xdd, 0x59, 0x76, 0x26, 0x6d, 0x73, 0xf5, 0x93, 0x4b, 0x57, 0xa5, 0x2f, 0x17, 0xaf,
	0xff, 0xfc, 0x66, 0x98, 0xb0, 0xe2, 0xcc, 0x1b, 0x0b, 0xcc, 0x40, 0x8d, 0xb5, 0x30, 0x93, 0x3a,
	0xf7, 0x01, 0x19, 0xd3, 0x65, 0x69, 0xa8, 0x44, 0x30, 0x48, 0xb0, 0x0a, 0x60, 0xfc, 0x62, 0x02,
	0x6b, 0x88, 0x5e, 0x85, 0xdc, 0x1b, 0x9b, 0x13, 0x4b, 0x57, 0xcd, 0xda, 0xbc, 0x52, 0x73, 0xd9,
	0xd4, 0x0c, 0x25, 0x3a, 0x56, 0x2b, 0xce, 0x19, 0x3d, 0xd6, 0xa8, 0x5b, 0x50, 0x8e, 0x43, 0x0b,
	0xf3, 0x5a, 0x6a, 0x08, 0xd8, 0xb0, 0x82, 0x12, 0x1b, 0xc7, 0xf4, 0x17, 0x18, 0x4e, 0xc5, 0xe9,
	0xde, 0x45, 0xf0, 0xfc, 0xf0, 0xf1, 0x9e, 0x9e, 0x0b, 0xb1, 0x99, 0xc8, 0x92, 0x41, 0x32, 0xea,
	0x14, 0x9d, 0x58, 0x1e, 0x6b, 0xa9, 0x87, 0x9f, 0xff, 0x44, 0x7f, 0x6f, 0x2e, 0xbd, 0x10, 0x5d,
	0x1d, 0xb0, 0xf1, 0x10, 0xa4, 0xd3, 0x8a, 0xb2, 0x9b, 0x41, 0x32, 0xea, 0x15, 0x87, 0xb1, 0x15,
	0x31, 0xa5, 0xb9, 0x38, 0x91, 0x25, 0x9b, 0x85, 0x82, 0x1d, 0x39, 0x8e, 0xf2, 0xf8, 0xfb, 0xd3,
	0xd3, 0x96, 0xff, 0x59, 0xf9, 0x7b, 0x6d, 0x36, 0xd9, 0xac, 0x64, 0x75, 0xbf, 0x4e, 0xe9, 0xa5,
	0x38, 0xa2, 0x9d, 0x5f, 0xca, 0xa6, 0x11, 0x75, 0x49, 0x17, 0x1b, 0x75, 0x25, 0xfa, 0xa4, 0xc3,
	0x0e, 0xbb, 0x8d, 0xac, 0x47, 0x3a, 0x6c, 0xb9, 0x6b, 0x91, 0x1a, 0x02, 0x87, 0x95, 0x82, 0x1a,
	0x97, 0x60, 0x95, 0xc5, 0xd0, 0x66, 0x77, 0x83, 0x64, 0x74, 0x50, 0xf4, 0x0d, 0xbd, 0x61, 0xa5,
	0x5e, 0x70, 0xf9, 0x1a, 0xf3, 0xec, 0x7f, 0x7c, 0xbb, 0xf1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x91, 0xa3, 0xee, 0x44, 0xd4, 0x01, 0x00, 0x00,
}