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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_topology_entry_flag_route_counts_topology_entry_flag_route_count

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	EntryFlag            string   `protobuf:"bytes,2,opt,name=entry_flag,json=entryFlag,proto3" json:"entry_flag,omitempty"`
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

func (m *PimTtSummBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0x06, 0x70, 0xb6, 0x87, 0xf7, 0x6d, 0x63, 0x6b, 0x31, 0x2a, 0xac, 0x07, 0xa1, 0x16, 0x91,
	0x82, 0xb0, 0x07, 0x5b, 0xeb, 0x07, 0x10, 0xf5, 0xa0, 0x16, 0x59, 0x4f, 0x9e, 0x86, 0x74, 0x9b,
	0x0d, 0x81, 0xcd, 0x4e, 0x98, 0x64, 0xb7, 0xf6, 0xea, 0x27, 0x97, 0x4d, 0x95, 0xfe, 0xb9, 0x78,
	0x09, 0xe4, 0x99, 0x1f, 0x0f, 0x19, 0xc2, 0x4e, 0xad, 0x36, 0xe0, 0x3d, 0xb8, 0xca, 0x18, 0x98,
	0x0b, 0x95, 0x58, 0x42, 0x8f, 0xdc, 0x67, 0xda, 0x65, 0x08, 0x1a, 0x1d, 0x7c, 0x12, 0x68, 0x5b,
	0x4f, 0xa0, 0x81, 0x68, 0x25, 0x25, 0xda, 0xd6, 0xd3, 0xe6, 0x96, 0x88, 0xcc, 0xeb, 0x5a, 0x26,
	0x35, 0xe5, 0xae, 0x39, 0x12, 0x8f, 0x16, 0x0b, 0x54, 0x2b, 0x90, 0xa5, 0xa7, 0x15, 0xe4, 0x85,
	0x50, 0x40, 0x58, 0x79, 0x09, 0x19, 0x56, 0xa5, 0x77, 0x7f, 0x81, 0xe1, 0x1b, 0x3b, 0xd9, 0x7b,
	0x0e, 0x3c, 0x3f, 0x7c, 0xbc, 0xf3, 0x33, 0xd6, 0xae, 0x29, 0x87, 0x52, 0x18, 0x19, 0x47, 0x83,
	0x68, 0xd4, 0x49, 0xff, 0xd7, 0x94, 0xcf, 0x84, 0x91, 0xfc, 0x9c, 0xb1, 0x4d, 0x59, 0xdc, 0x0a,
	0xc3, 0x4e, 0x48, 0x1e, 0x0b, 0xa1, 0x86, 0x5f, 0x2d, 0xd6, 0xdf, 0xab, 0xe4, 0x17, 0xac, 0xab,
	0x08, 0x2b, 0x0b, 0x24, 0x4a, 0x25, 0x5d, 0x7c, 0x33, 0x88, 0x46, 0xbd, 0xf4, 0x20, 0x64, 0x69,
	0x88, 0x78, 0xc2, 0x8e, 0xd7, 0xcb, 0xc1, 0x8e, 0x1c, 0x07, 0x79, 0xb4, 0x1e, 0x3d, 0x6d, 0xf9,
	0x9f, 0xca, 0xdf, 0x45, 0xe2, 0xc9, 0xa6, 0xd2, 0xcb, 0xfb, 0x26, 0xe2, 0x97, 0xec, 0xd0, 0xed,
	0x6c, 0x1b, 0xdf, 0x06, 0xd4, 0x75, 0x2a, 0xdd, 0xa8, 0x2b, 0xd6, 0x77, 0x8a, 0x76, 0xd8, 0x34,
	0xb0, 0x9e, 0x53, 0xb4, 0xe5, 0xae, 0x19, 0xd7, 0x0e, 0x4a, 0x5c, 0x48, 0x28, 0x70, 0x09, 0x46,
	0x1a, 0xa4, 0x55, 0x7c, 0x37, 0x88, 0x46, 0xed, 0xb4, 0xaf, 0xdd, 0x0c, 0x17, 0xf2, 0x05, 0x97,
	0xaf, 0x21, 0x9e, 0xff, 0x0b, 0x7f, 0x3a, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x86, 0x5d, 0xba,
	0xb1, 0xec, 0x01, 0x00, 0x00,
}
