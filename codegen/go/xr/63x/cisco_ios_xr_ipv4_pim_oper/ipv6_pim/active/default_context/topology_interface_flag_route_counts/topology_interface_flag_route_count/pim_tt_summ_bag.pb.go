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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_topology_interface_flag_route_counts_topology_interface_flag_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xbb, 0x4f, 0xfb, 0x30,
	0x10, 0x07, 0x70, 0xe5, 0x37, 0xfc, 0x04, 0xa6, 0x0f, 0x11, 0x40, 0xca, 0x58, 0x2a, 0x40, 0x95,
	0x90, 0x32, 0xd0, 0x52, 0x26, 0x26, 0x04, 0x0c, 0x3c, 0x86, 0x30, 0x31, 0x9d, 0x5c, 0xf7, 0x6a,
	0x59, 0x4a, 0x72, 0x96, 0x7d, 0xe9, 0x63, 0x85, 0x7f, 0x1c, 0xc5, 0x05, 0xda, 0x74, 0x62, 0xf4,
	0xd7, 0x9f, 0x3b, 0xdd, 0xe9, 0xc4, 0x89, 0x35, 0x05, 0x30, 0x83, 0xaf, 0x8a, 0x02, 0x26, 0x52,
	0xa7, 0xd6, 0x11, 0x53, 0xfc, 0x19, 0x29, 0xe3, 0x15, 0x81, 0x21, 0x0f, 0x4b, 0x07, 0xc6, 0xce,
	0x47, 0x50, 0x4b, 0xb2, 0xe8, 0x52, 0x63, 0xe7, 0xe3, 0xfa, 0x95, 0x4a, 0xc5, 0x66, 0x8e, 0xe9,
	0x14, 0x67, 0xb2, 0xca, 0x19, 0x14, 0x95, 0x8c, 0x4b, 0x4e, 0x99, 0x2c, 0xe5, 0xa4, 0x57, 0x60,
	0x4a, 0x46, 0x37, 0x93, 0x0a, 0x61, 0x96, 0x4b, 0x0d, 0x8e, 0x2a, 0x46, 0x50, 0x54, 0x95, 0xec,
	0xff, 0x82, 0xfa, 0xb7, 0xe2, 0x78, 0x67, 0x3c, 0x78, 0xba, 0x7f, 0x7f, 0x8b, 0xcf, 0x45, 0xa7,
	0x59, 0x95, 0x44, 0xbd, 0x68, 0xb0, 0x9f, 0xb5, 0x7f, 0xd3, 0x87, 0x5c, 0xea, 0xfe, 0xc7, 0x3f,
	0xd1, 0xdd, 0xa9, 0x8f, 0x4f, 0x45, 0x4b, 0x3b, 0xaa, 0x2c, 0x38, 0x59, 0x6a, 0xf4, 0xc9, 0x55,
	0x2f, 0x1a, 0xb4, 0xb3, 0x83, 0x90, 0x65, 0x21, 0x8a, 0x53, 0x71, 0xb4, 0x5e, 0x0c, 0x1a, 0x72,
	0x18, 0xe4, 0xe1, 0xfa, 0xeb, 0x71, 0xcb, 0x7f, 0xb7, 0xfc, 0x99, 0x3a, 0x19, 0x6d, 0x5a, 0x32,
	0xde, 0xd5, 0x51, 0x7c, 0x26, 0x3a, 0xbe, 0xb1, 0x5a, 0x72, 0x1d, 0x50, 0xcb, 0xeb, 0x6c, 0xa3,
	0x2e, 0x44, 0xd7, 0x6b, 0xd7, 0x60, 0xe3, 0xc0, 0xda, 0x5e, 0xbb, 0x2d, 0x77, 0x29, 0x62, 0xe3,
	0xa1, 0xa4, 0x29, 0x42, 0x4e, 0x0b, 0x28, 0xb0, 0x20, 0xb7, 0x4a, 0x6e, 0x7a, 0xd1, 0x60, 0x2f,
	0xeb, 0x1a, 0xff, 0x4a, 0x53, 0x7c, 0xa6, 0xc5, 0x4b, 0x88, 0x27, 0xff, 0xc3, 0x41, 0x87, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x56, 0x23, 0xb5, 0xe9, 0x01, 0x00, 0x00,
}
