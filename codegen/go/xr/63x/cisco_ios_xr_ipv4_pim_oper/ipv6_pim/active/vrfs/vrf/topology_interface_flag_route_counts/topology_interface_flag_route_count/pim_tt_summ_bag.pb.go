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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_topology_interface_flag_route_counts_topology_interface_flag_route_count

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
	InterfaceFlag        string   `protobuf:"bytes,2,opt,name=interface_flag,json=interfaceFlag,proto3" json:"interface_flag,omitempty"`
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

func (m *PimTtSummBag_KEYS) GetInterfaceFlag() string {
	if m != nil {
		return m.InterfaceFlag
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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0x06, 0x70, 0xba, 0xc3, 0xf7, 0x3b, 0xe3, 0x7e, 0x60, 0x55, 0xa8, 0xb7, 0x39, 0x54, 0x06,
	0x42, 0x0f, 0x6e, 0xce, 0x3f, 0x40, 0xd4, 0x83, 0xba, 0x43, 0xbd, 0xe8, 0xe9, 0x25, 0xeb, 0xd2,
	0x10, 0x68, 0xfa, 0x86, 0x37, 0x69, 0xe7, 0xc0, 0x93, 0x7f, 0xb9, 0x34, 0x53, 0xb7, 0xee, 0xe4,
	0xa5, 0xd0, 0x27, 0x1f, 0x1e, 0xf2, 0x10, 0x76, 0x6c, 0x94, 0x06, 0xe7, 0xc0, 0x96, 0x5a, 0xc3,
	0x9c, 0xcb, 0xd8, 0x10, 0x3a, 0x0c, 0x3f, 0x52, 0x65, 0x53, 0x04, 0x85, 0x16, 0xde, 0x09, 0x94,
	0xa9, 0x26, 0x50, 0x43, 0x34, 0x82, 0x62, 0x65, 0xaa, 0x69, 0xfd, 0x17, 0xf3, 0xd4, 0xa9, 0x4a,
	0xc4, 0x15, 0x65, 0xb6, 0xfe, 0xc4, 0x0e, 0x0d, 0xe6, 0x28, 0x57, 0xa0, 0x0a, 0x27, 0x28, 0xe3,
	0xa9, 0x80, 0x2c, 0xe7, 0x12, 0x08, 0x4b, 0x27, 0x20, 0xc5, 0xb2, 0x70, 0xf6, 0x2f, 0x68, 0xf8,
	0xca, 0x8e, 0x76, 0xae, 0x05, 0x8f, 0x77, 0x6f, 0x2f, 0xe1, 0x09, 0x6b, 0x57, 0x94, 0x41, 0xc1,
	0xb5, 0x88, 0x82, 0x41, 0x30, 0xda, 0x4b, 0xfe, 0x57, 0x94, 0xcd, 0xb8, 0x16, 0xe1, 0x39, 0xeb,
	0x35, 0x0b, 0xa3, 0x96, 0x07, 0xdd, 0xdf, 0xf4, 0x3e, 0xe7, 0x72, 0xf8, 0xd9, 0x62, 0xfd, 0x9d,
	0xea, 0xf0, 0x94, 0x75, 0x24, 0x61, 0x69, 0x80, 0x78, 0x21, 0x85, 0x8d, 0xae, 0x06, 0xc1, 0xa8,
	0x9b, 0xec, 0xfb, 0x2c, 0xf1, 0x51, 0x18, 0xb3, 0xc3, 0xf5, 0x58, 0x68, 0xc8, 0xb1, 0x97, 0x07,
	0xeb, 0xa3, 0x87, 0x2d, 0xff, 0x5d, 0xf9, 0x33, 0x28, 0x9a, 0x6c, 0x2a, 0x9d, 0xb8, 0xad, 0xa3,
	0xf0, 0x8c, 0xf5, 0x6c, 0x63, 0x75, 0x74, 0xed, 0x51, 0xc7, 0xca, 0x64, 0xa3, 0x2e, 0x58, 0xdf,
	0x4a, 0x6a, 0xb0, 0xa9, 0x67, 0x5d, 0x2b, 0x69, 0xcb, 0x5d, 0xb2, 0x50, 0x59, 0x28, 0x70, 0x21,
	0x20, 0xc7, 0x25, 0x68, 0xa1, 0x91, 0x56, 0xd1, 0xcd, 0x20, 0x18, 0xb5, 0x93, 0xbe, 0xb2, 0x33,
	0x5c, 0x88, 0x27, 0x5c, 0x3e, 0xfb, 0x78, 0xfe, 0xcf, 0xbf, 0xf1, 0xf8, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x56, 0xaa, 0x89, 0x7e, 0xfc, 0x01, 0x00, 0x00,
}
