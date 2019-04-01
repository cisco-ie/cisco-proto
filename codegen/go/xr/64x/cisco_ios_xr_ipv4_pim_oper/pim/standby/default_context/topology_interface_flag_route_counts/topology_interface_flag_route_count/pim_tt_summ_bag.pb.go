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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_topology_interface_flag_route_counts_topology_interface_flag_route_count

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
	InterfaceFlag        string   `protobuf:"bytes,1,opt,name=interface_flag,json=interfaceFlag,proto3" json:"interface_flag,omitempty"`
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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xcb, 0x4b, 0xfb, 0x40,
	0x10, 0xc0, 0x71, 0xf2, 0x3b, 0xfc, 0xd0, 0xb5, 0x0f, 0x8c, 0x0a, 0x39, 0xd6, 0xa2, 0x52, 0x10,
	0x72, 0xb0, 0x55, 0x4f, 0x9e, 0x44, 0x3d, 0xf8, 0x38, 0xc4, 0x93, 0xa7, 0x61, 0x9b, 0x6c, 0x97,
	0x85, 0x64, 0x67, 0xd9, 0x99, 0xb4, 0xcd, 0x49, 0xf0, 0x2f, 0x97, 0xae, 0x8f, 0x36, 0x3d, 0x79,
	0xfd, 0xe6, 0x33, 0xc3, 0x84, 0x15, 0x47, 0xce, 0x54, 0xc0, 0x0c, 0x54, 0x57, 0x15, 0x4c, 0xa5,
	0x4e, 0x9d, 0x47, 0xc6, 0xf8, 0x3d, 0x37, 0x94, 0x23, 0x18, 0x24, 0x58, 0x7a, 0x30, 0x6e, 0x3e,
	0x81, 0x15, 0x44, 0xa7, 0x7c, 0xea, 0x4c, 0x95, 0x12, 0x4b, 0x5b, 0x4c, 0x9b, 0xb4, 0x50, 0x33,
	0x59, 0x97, 0x0c, 0x39, 0x5a, 0x56, 0x4b, 0x4e, 0x19, 0x1d, 0x96, 0xa8, 0x1b, 0x30, 0x96, 0x95,
	0x9f, 0xc9, 0x5c, 0xc1, 0xac, 0x94, 0x1a, 0x3c, 0xd6, 0xac, 0x20, 0xc7, 0xda, 0x32, 0xfd, 0x05,
	0x0d, 0x6f, 0xc4, 0xe1, 0xd6, 0x65, 0xf0, 0x78, 0xf7, 0xf6, 0x1a, 0x9f, 0x8a, 0x5e, 0x7b, 0x2a,
	0x89, 0x06, 0xd1, 0x68, 0x37, 0xeb, 0xfe, 0xd6, 0xfb, 0x52, 0xea, 0xe1, 0xc7, 0x3f, 0xd1, 0xdf,
	0x9a, 0x8f, 0x8f, 0x45, 0x47, 0x7b, 0xac, 0x1d, 0x78, 0x69, 0xb5, 0xa2, 0xe4, 0x62, 0x10, 0x8d,
	0xba, 0xd9, 0x5e, 0x68, 0x59, 0x48, 0x71, 0x2a, 0x0e, 0x64, 0xce, 0x66, 0xae, 0xa0, 0x25, 0xc7,
	0x41, 0xee, 0x7f, 0x7d, 0x7a, 0xd8, 0xf0, 0xdf, 0x2b, 0x7f, 0xae, 0x4e, 0x26, 0xeb, 0x95, 0xac,
	0x6e, 0x57, 0x29, 0x3e, 0x11, 0x3d, 0x6a, 0xfd, 0x5a, 0x72, 0x19, 0x50, 0x87, 0x74, 0xb6, 0x56,
	0x67, 0xa2, 0x4f, 0xda, 0xb7, 0xd8, 0x55, 0x60, 0x5d, 0xd2, 0x7e, 0xc3, 0x9d, 0x8b, 0xd8, 0x10,
	0x58, 0x2c, 0x14, 0x94, 0xb8, 0x80, 0x4a, 0x55, 0xe8, 0x9b, 0xe4, 0x7a, 0x10, 0x8d, 0x76, 0xb2,
	0xbe, 0xa1, 0x17, 0x2c, 0xd4, 0x13, 0x2e, 0x9e, 0x43, 0x9e, 0xfe, 0x0f, 0x6f, 0x39, 0xfe, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x64, 0x69, 0xb9, 0x60, 0xe4, 0x01, 0x00, 0x00,
}