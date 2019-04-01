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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_vrfs_vrf_topology_entry_flag_route_counts_topology_entry_flag_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.topology_entry_flag_route_counts.topology_entry_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0xcf, 0x4a, 0x2b, 0x31,
	0x14, 0x06, 0x70, 0xa6, 0x8b, 0x7b, 0xdb, 0xdc, 0xf6, 0x16, 0xa3, 0xc2, 0xb8, 0x10, 0x6a, 0x11,
	0x29, 0x08, 0xb3, 0xb0, 0xb5, 0x3e, 0x80, 0xa8, 0x0b, 0xb5, 0xc8, 0xb8, 0x72, 0x75, 0x48, 0xa7,
	0x99, 0x10, 0x98, 0xcc, 0x09, 0x39, 0x99, 0xa9, 0xb3, 0xf5, 0xc9, 0x65, 0xa2, 0xd2, 0x3f, 0x1b,
	0x37, 0x81, 0x7c, 0xe7, 0xc7, 0x47, 0x0e, 0x61, 0xc7, 0x56, 0x1b, 0xf0, 0x1e, 0xa8, 0x32, 0x06,
	0x96, 0x42, 0x25, 0xd6, 0xa1, 0x47, 0x5e, 0x65, 0x9a, 0x32, 0x04, 0x8d, 0x04, 0xef, 0x0e, 0xb4,
	0xad, 0x67, 0xd0, 0x42, 0xb4, 0xd2, 0x25, 0xda, 0xd6, 0xf3, 0xf6, 0x96, 0x90, 0x17, 0xe5, 0x6a,
	0xd9, 0x24, 0xb5, 0xcb, 0xa9, 0x3d, 0x12, 0x8f, 0x16, 0x0b, 0x54, 0x0d, 0xc8, 0xd2, 0xbb, 0x06,
	0xf2, 0x42, 0x28, 0x70, 0x58, 0x79, 0x09, 0x19, 0x56, 0xa5, 0xa7, 0xdf, 0xc0, 0xf8, 0x85, 0x1d,
	0xed, 0xbd, 0x07, 0x1e, 0xef, 0xde, 0x5e, 0xf9, 0x09, 0xeb, 0xd6, 0x2e, 0x87, 0x52, 0x18, 0x19,
	0x47, 0xa3, 0x68, 0xd2, 0x4b, 0xff, 0xd6, 0x2e, 0x5f, 0x08, 0x23, 0xf9, 0x29, 0x63, 0x9b, 0xb2,
	0xb8, 0x13, 0x86, 0xbd, 0x90, 0xdc, 0x17, 0x42, 0x8d, 0x3f, 0x3a, 0x6c, 0xb8, 0x57, 0xc9, 0xcf,
	0x58, 0x5f, 0x39, 0xac, 0x2c, 0x38, 0x51, 0x2a, 0x49, 0xf1, 0xd5, 0x28, 0x9a, 0x0c, 0xd2, 0x7f,
	0x21, 0x4b, 0x43, 0xc4, 0x13, 0x76, 0x28, 0x32, 0xaf, 0x6b, 0x09, 0x3b, 0x72, 0x1a, 0xe4, 0xc1,
	0xd7, 0xe8, 0x61, 0xcb, 0x7f, 0x57, 0xfe, 0x2c, 0x12, 0xcf, 0x36, 0x95, 0x5e, 0xde, 0xb6, 0x11,
	0x3f, 0x67, 0xff, 0x69, 0x67, 0xdb, 0xf8, 0x3a, 0xa0, 0x3e, 0xa9, 0x74, 0xa3, 0x2e, 0xd8, 0x90,
	0x94, 0xdb, 0x61, 0xf3, 0xc0, 0x06, 0xa4, 0xdc, 0x96, 0xbb, 0x64, 0x5c, 0x13, 0x94, 0xb8, 0x92,
	0x50, 0xe0, 0x1a, 0x8c, 0x34, 0xe8, 0x9a, 0xf8, 0x66, 0x14, 0x4d, 0xba, 0xe9, 0x50, 0xd3, 0x02,
	0x57, 0xf2, 0x09, 0xd7, 0xcf, 0x21, 0x5e, 0xfe, 0x09, 0x9f, 0x3a, 0xfd, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0xb3, 0xff, 0x2f, 0xed, 0x01, 0x00, 0x00,
}