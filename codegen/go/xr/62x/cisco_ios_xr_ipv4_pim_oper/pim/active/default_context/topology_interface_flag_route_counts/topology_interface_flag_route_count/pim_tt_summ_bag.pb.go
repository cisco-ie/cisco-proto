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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_topology_interface_flag_route_counts_topology_interface_flag_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_addrtype")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0xd2, 0xbf, 0x8e, 0xd4, 0x30,
	0x10, 0x06, 0x70, 0x85, 0x93, 0x0e, 0xf0, 0x6e, 0x6e, 0x45, 0x00, 0x91, 0x32, 0xac, 0x00, 0x45,
	0x42, 0x72, 0x71, 0x77, 0x2c, 0x34, 0x14, 0x08, 0x01, 0x05, 0x70, 0x45, 0xa8, 0xa8, 0x46, 0xbe,
	0x64, 0x62, 0x59, 0x8a, 0x33, 0x96, 0xed, 0xec, 0x9f, 0x86, 0x27, 0xe0, 0x69, 0x68, 0x78, 0x3d,
	0x14, 0x67, 0xd9, 0x4d, 0xb6, 0xa2, 0xa3, 0xfd, 0xfc, 0x9b, 0x68, 0xfc, 0xc5, 0xec, 0xb1, 0x51,
	0x1a, 0xbc, 0x07, 0xd7, 0x69, 0x0d, 0xb7, 0x42, 0x72, 0x63, 0xc9, 0x53, 0xf2, 0xa3, 0x54, 0xae,
	0x24, 0x50, 0xe4, 0x60, 0x6b, 0x41, 0x99, 0xf5, 0x35, 0xf4, 0x90, 0x0c, 0x5a, 0x6e, 0x94, 0xe6,
	0xa2, 0xf4, 0x6a, 0x8d, 0xbc, 0xc2, 0x5a, 0x74, 0x8d, 0x87, 0x92, 0x5a, 0x8f, 0x5b, 0xcf, 0x3d,
	0x19, 0x6a, 0x48, 0xee, 0x40, 0xb5, 0x1e, 0x6d, 0x2d, 0x4a, 0x84, 0xba, 0x11, 0x12, 0x2c, 0x75,
	0x1e, 0xa1, 0xa4, 0xae, 0xf5, 0xee, 0x5f, 0xd0, 0xf2, 0x2d, 0x7b, 0x74, 0xb2, 0x18, 0x7c, 0xfe,
	0xf0, 0xfd, 0x5b, 0xf2, 0x9c, 0x5d, 0x4c, 0xa7, 0xd2, 0x28, 0x8b, 0xf2, 0xfb, 0x45, 0x7c, 0x48,
	0x3f, 0x36, 0x42, 0x2e, 0x35, 0x9b, 0xf7, 0xe3, 0xa2, 0xaa, 0xac, 0xdf, 0x19, 0x4c, 0x9e, 0xb0,
	0xbb, 0xa2, 0x86, 0x56, 0x68, 0xdc, 0xfb, 0x73, 0x51, 0xdf, 0x08, 0x8d, 0xc9, 0x53, 0x36, 0x0f,
	0x97, 0xeb, 0x25, 0x3a, 0x97, 0xde, 0x09, 0xa7, 0xb3, 0x3e, 0x7b, 0x37, 0x44, 0x7b, 0xb2, 0x3a,
	0x90, 0xb3, 0x03, 0x59, 0xed, 0xc9, 0xf2, 0xf7, 0x19, 0x5b, 0x9c, 0xac, 0xdb, 0x8f, 0x49, 0x4b,
	0x9d, 0x01, 0x2b, 0x5a, 0x89, 0x2e, 0xbd, 0xcc, 0xa2, 0x3c, 0x2e, 0x66, 0x21, 0x2b, 0x42, 0x94,
	0x70, 0xf6, 0x70, 0xe8, 0x11, 0x26, 0xf2, 0x2a, 0xc8, 0x07, 0xc3, 0xd1, 0xa7, 0x91, 0xff, 0x15,
	0xb1, 0x78, 0x90, 0x7f, 0x77, 0xb9, 0xce, 0xa2, 0x7c, 0x76, 0xf9, 0x33, 0xe2, 0xff, 0xf7, 0x77,
	0xf1, 0x71, 0xd9, 0xc5, 0x70, 0xef, 0x51, 0x7d, 0x72, 0x44, 0xd3, 0x57, 0xc7, 0x1e, 0x3c, 0xbe,
	0xef, 0xa3, 0xe4, 0x19, 0xbb, 0x70, 0x93, 0xef, 0xa5, 0xab, 0x80, 0xe6, 0x4e, 0x16, 0x47, 0xf5,
	0x82, 0x2d, 0x9c, 0xb4, 0x13, 0xf6, 0x3a, 0xb0, 0xd8, 0x49, 0x3b, 0x72, 0x2f, 0x59, 0xa2, 0x1c,
	0xb4, 0x54, 0x21, 0x34, 0xb4, 0x01, 0x8d, 0x9a, 0xec, 0x2e, 0x7d, 0x93, 0x45, 0xf9, 0xbd, 0x62,
	0xa1, 0xdc, 0x0d, 0x55, 0xf8, 0x85, 0x36, 0x5f, 0x43, 0x7c, 0x7b, 0x1e, 0x9e, 0xfb, 0xd5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x3b, 0xfe, 0x9a, 0x07, 0x03, 0x00, 0x00,
}
