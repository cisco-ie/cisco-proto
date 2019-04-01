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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_topology_interface_flag_route_counts_topology_interface_flag_route_count

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
	proto.RegisterType((*PimTtSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag_KEYS")
	proto.RegisterType((*PimTtSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.topology_interface_flag_route_counts.topology_interface_flag_route_count.pim_tt_summ_bag")
}

func init() { proto.RegisterFile("pim_tt_summ_bag.proto", fileDescriptor_dc46104ddf9a9167) }

var fileDescriptor_dc46104ddf9a9167 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0x3d, 0x4f, 0xf3, 0x30,
	0x14, 0x05, 0x60, 0xa5, 0xc3, 0xfb, 0x16, 0xd3, 0x0f, 0x61, 0x40, 0x0a, 0x5b, 0xa9, 0x00, 0x55,
	0x42, 0xca, 0x40, 0x0b, 0xfc, 0x00, 0x04, 0x0c, 0x40, 0x87, 0xb0, 0xc0, 0x74, 0xe5, 0xa6, 0x8e,
	0x65, 0x29, 0xce, 0xb5, 0x7c, 0x9d, 0xb4, 0x5d, 0xf9, 0xe5, 0xa8, 0xe6, 0xa3, 0x4d, 0x27, 0x16,
	0x0f, 0xc7, 0x8f, 0x8e, 0x7c, 0x64, 0x76, 0x6c, 0xb5, 0x01, 0xef, 0x81, 0x2a, 0x63, 0x60, 0x26,
	0x54, 0x62, 0x1d, 0x7a, 0xe4, 0xcb, 0x4c, 0x53, 0x86, 0xa0, 0x91, 0x60, 0xe9, 0x40, 0xdb, 0x7a,
	0x02, 0x6b, 0x88, 0x56, 0xba, 0xc4, 0x6a, 0x93, 0x90, 0x17, 0xe5, 0x7c, 0xb6, 0x4a, 0x6a, 0x97,
	0xd3, 0xfa, 0x48, 0x3c, 0x5a, 0x2c, 0x50, 0xad, 0x40, 0x97, 0x5e, 0xba, 0x5c, 0x64, 0x12, 0xf2,
	0x42, 0x28, 0x70, 0x58, 0x79, 0x09, 0x19, 0x56, 0xa5, 0xa7, 0xbf, 0xa0, 0xe1, 0x1b, 0x3b, 0xda,
	0x79, 0x12, 0x3c, 0xdd, 0xbf, 0xbf, 0xf2, 0x13, 0xd6, 0xae, 0x5d, 0x0e, 0xa5, 0x30, 0x32, 0x8e,
	0x06, 0xd1, 0x68, 0x2f, 0xfd, 0x5f, 0xbb, 0x7c, 0x2a, 0x8c, 0xe4, 0xe7, 0xac, 0xd7, 0x2c, 0x8c,
	0x5b, 0x01, 0x74, 0x7f, 0xd3, 0x87, 0x42, 0xa8, 0xe1, 0x47, 0x8b, 0xf5, 0x77, 0xaa, 0xf9, 0x29,
	0xeb, 0x28, 0x87, 0x95, 0x05, 0x27, 0x4a, 0x25, 0x29, 0xbe, 0x1a, 0x44, 0xa3, 0x6e, 0xba, 0x1f,
	0xb2, 0x34, 0x44, 0x3c, 0x61, 0x87, 0x22, 0xf3, 0xba, 0x96, 0xd0, 0x90, 0xe3, 0x20, 0x0f, 0xbe,
	0xae, 0x1e, 0xb7, 0xfc, 0x77, 0xe5, 0xcf, 0xa0, 0x78, 0xb2, 0xa9, 0xf4, 0xf2, 0x6e, 0x1d, 0xf1,
	0x33, 0xd6, 0xa3, 0xc6, 0xea, 0xf8, 0x3a, 0xa0, 0x0e, 0xa9, 0x74, 0xa3, 0x2e, 0x58, 0x9f, 0x94,
	0x6b, 0xb0, 0x9b, 0xc0, 0xba, 0xa4, 0xdc, 0x96, 0xbb, 0x64, 0x5c, 0x13, 0x94, 0x38, 0x97, 0x50,
	0xe0, 0x02, 0x8c, 0x34, 0xe8, 0x56, 0xf1, 0xed, 0x20, 0x1a, 0xb5, 0xd3, 0xbe, 0xa6, 0x29, 0xce,
	0xe5, 0x33, 0x2e, 0x5e, 0x42, 0x3c, 0xfb, 0x17, 0xfe, 0x77, 0xfc, 0x19, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0xa2, 0xc6, 0x81, 0xf8, 0x01, 0x00, 0x00,
}