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

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_topology_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.topology_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.topology_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xd1, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0xea, 0x41, 0x67, 0xdc, 0x1c, 0x46, 0x85, 0x7a, 0xab, 0x43, 0xa4, 0x20, 0x04,
	0x74, 0x53, 0x3f, 0xc0, 0x10, 0x0f, 0xea, 0x0e, 0xf5, 0xe4, 0xe9, 0xd1, 0x75, 0x69, 0x08, 0x2c,
	0x7d, 0xe1, 0x25, 0xed, 0xdc, 0xd5, 0x4f, 0x2e, 0xcd, 0x94, 0x75, 0xbb, 0xe4, 0xf0, 0xcf, 0x2f,
	0x7f, 0x5e, 0x78, 0xec, 0xd2, 0x6a, 0x03, 0xde, 0x83, 0xab, 0x8d, 0x81, 0x79, 0xae, 0x84, 0x25,
	0xf4, 0xc8, 0xa7, 0x85, 0x76, 0x05, 0x82, 0x46, 0x07, 0xdf, 0x04, 0xda, 0x36, 0x13, 0x68, 0x21,
	0x5a, 0x49, 0xc2, 0x6a, 0x23, 0xf2, 0xc2, 0xeb, 0x46, 0x8a, 0x86, 0x4a, 0xd7, 0x1e, 0xc2, 0xa3,
	0xc5, 0x25, 0xaa, 0x35, 0x10, 0xd6, 0x5e, 0x42, 0x81, 0x75, 0xe5, 0x47, 0xf7, 0xec, 0x62, 0xaf,
	0x1d, 0xde, 0x5e, 0xbe, 0x3e, 0xf9, 0x15, 0xeb, 0x35, 0x54, 0x42, 0x95, 0x1b, 0x19, 0x47, 0x49,
	0x94, 0x1e, 0x67, 0x47, 0x0d, 0x95, 0xb3, 0xdc, 0xc8, 0xd1, 0xcf, 0x01, 0x1b, 0xee, 0xbd, 0xe1,
	0xd7, 0xac, 0xaf, 0x08, 0x6b, 0x0b, 0x94, 0x57, 0x4a, 0xba, 0xf8, 0x21, 0x89, 0xd2, 0x41, 0x76,
	0x12, 0xb2, 0x2c, 0x44, 0x5c, 0xb0, 0xf3, 0xcd, 0x44, 0xb0, 0x23, 0xc7, 0x41, 0x9e, 0x6d, 0xae,
	0x5e, 0x3b, 0xfe, 0xaf, 0xf2, 0x7f, 0xd2, 0x78, 0xb2, 0xad, 0xf4, 0x72, 0xda, 0x46, 0xfc, 0x86,
	0x9d, 0x3a, 0xd5, 0xfd, 0x4e, 0xfc, 0x18, 0x50, 0xdf, 0xa9, 0x6c, 0xab, 0x6e, 0xd9, 0xd0, 0x29,
	0xda, 0x61, 0x4f, 0x81, 0x0d, 0x9c, 0xa2, 0x8e, 0xbb, 0x63, 0x5c, 0x3b, 0xa8, 0x70, 0x21, 0x61,
	0x89, 0x2b, 0x30, 0xd2, 0x20, 0xad, 0xe3, 0xe7, 0x24, 0x4a, 0x7b, 0xd9, 0x50, 0xbb, 0x19, 0x2e,
	0xe4, 0x3b, 0xae, 0x3e, 0x42, 0x3c, 0x3f, 0x0c, 0x3b, 0x18, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x12, 0xf2, 0x00, 0xc3, 0x9c, 0x01, 0x00, 0x00,
}