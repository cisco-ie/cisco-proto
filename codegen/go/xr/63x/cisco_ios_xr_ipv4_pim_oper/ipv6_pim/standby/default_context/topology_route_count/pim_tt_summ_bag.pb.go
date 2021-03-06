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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_topology_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.topology_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.topology_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd0, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x07, 0x70, 0xfa, 0x1c, 0x1e, 0x64, 0x6d, 0x0d, 0xc6, 0x17, 0x72, 0xac, 0x45, 0xa4, 0x20,
	0xec, 0xc1, 0xd6, 0xfa, 0x01, 0x44, 0x3c, 0xf8, 0x82, 0xc4, 0x93, 0xa7, 0x21, 0x2f, 0xeb, 0xb2,
	0x90, 0x64, 0x96, 0x9d, 0xd9, 0xb4, 0xb9, 0xfa, 0xc9, 0x25, 0xab, 0xd2, 0xb4, 0xc7, 0xfd, 0xcf,
	0x8f, 0xd9, 0x3f, 0x23, 0xce, 0xac, 0xa9, 0x81, 0x19, 0xc8, 0xd7, 0x35, 0xe4, 0x99, 0x96, 0xd6,
	0x21, 0x63, 0xfc, 0x56, 0x18, 0x2a, 0x10, 0x0c, 0x12, 0x6c, 0x1c, 0x18, 0xdb, 0x2e, 0xa1, 0x87,
	0x68, 0x95, 0x93, 0xc6, 0xb6, 0xab, 0xfe, 0x25, 0x89, 0xb3, 0xa6, 0xcc, 0x3b, 0x59, 0xaa, 0xcf,
	0xcc, 0x57, 0x0c, 0x05, 0x36, 0xac, 0x36, 0x2c, 0x19, 0x2d, 0x56, 0xa8, 0x3b, 0x70, 0xe8, 0x59,
	0x41, 0x81, 0xbe, 0xe1, 0xd9, 0xb9, 0x38, 0xdd, 0xfb, 0x0a, 0x9e, 0x1e, 0x3e, 0xde, 0x67, 0x5f,
	0xff, 0x44, 0xb4, 0x37, 0x88, 0x2f, 0xc4, 0x58, 0x3b, 0xf4, 0x16, 0x5c, 0xd6, 0x68, 0x45, 0xc9,
	0xcd, 0x74, 0x34, 0x9f, 0xa4, 0x87, 0x21, 0x4b, 0x43, 0x14, 0x4b, 0x71, 0x92, 0x15, 0x6c, 0x5a,
	0x05, 0x3b, 0x72, 0x11, 0xe4, 0xf1, 0xcf, 0xe8, 0x71, 0xe0, 0x7f, 0x57, 0xfe, 0xd5, 0x49, 0x96,
	0xdb, 0x95, 0xac, 0xee, 0xfb, 0x28, 0xbe, 0x14, 0x47, 0xa4, 0x87, 0x9d, 0x93, 0xdb, 0x80, 0xc6,
	0xa4, 0xd3, 0xad, 0xba, 0x12, 0x11, 0x69, 0xb7, 0xc3, 0x56, 0x81, 0x4d, 0x48, 0xbb, 0x81, 0xbb,
	0x16, 0xb1, 0x21, 0x68, 0xb0, 0x54, 0x50, 0xe1, 0x1a, 0x6a, 0x55, 0xa3, 0xeb, 0x92, 0xbb, 0xe9,
	0x68, 0x7e, 0x90, 0x46, 0x86, 0x5e, 0xb1, 0x54, 0xcf, 0xb8, 0x7e, 0x09, 0x71, 0xfe, 0x3f, 0x5c,
	0x7d, 0xf1, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x98, 0xaa, 0xfa, 0x8e, 0x01, 0x00, 0x00,
}
