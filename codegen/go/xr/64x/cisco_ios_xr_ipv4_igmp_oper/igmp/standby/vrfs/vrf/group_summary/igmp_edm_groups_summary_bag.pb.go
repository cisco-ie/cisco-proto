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
// source: igmp_edm_groups_summary_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_vrfs_vrf_group_summary

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

type IgmpEdmGroupsSummaryBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmGroupsSummaryBag_KEYS) Reset()         { *m = IgmpEdmGroupsSummaryBag_KEYS{} }
func (m *IgmpEdmGroupsSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsSummaryBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmGroupsSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f606fa9135171ea, []int{0}
}

func (m *IgmpEdmGroupsSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmGroupsSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS.Size(m)
}
func (m *IgmpEdmGroupsSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsSummaryBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmGroupsSummaryBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IgmpEdmGroupsSummaryBag struct {
	Groutes              uint32   `protobuf:"varint,50,opt,name=groutes,proto3" json:"groutes,omitempty"`
	SgRoutes             uint32   `protobuf:"varint,51,opt,name=sg_routes,json=sgRoutes,proto3" json:"sg_routes,omitempty"`
	GroupCount           uint32   `protobuf:"varint,52,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	IsLowMemory          bool     `protobuf:"varint,53,opt,name=is_low_memory,json=isLowMemory,proto3" json:"is_low_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmGroupsSummaryBag) Reset()         { *m = IgmpEdmGroupsSummaryBag{} }
func (m *IgmpEdmGroupsSummaryBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsSummaryBag) ProtoMessage()    {}
func (*IgmpEdmGroupsSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f606fa9135171ea, []int{1}
}

func (m *IgmpEdmGroupsSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsSummaryBag.Merge(m, src)
}
func (m *IgmpEdmGroupsSummaryBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsSummaryBag.Size(m)
}
func (m *IgmpEdmGroupsSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsSummaryBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsSummaryBag) GetGroutes() uint32 {
	if m != nil {
		return m.Groutes
	}
	return 0
}

func (m *IgmpEdmGroupsSummaryBag) GetSgRoutes() uint32 {
	if m != nil {
		return m.SgRoutes
	}
	return 0
}

func (m *IgmpEdmGroupsSummaryBag) GetGroupCount() uint32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *IgmpEdmGroupsSummaryBag) GetIsLowMemory() bool {
	if m != nil {
		return m.IsLowMemory
	}
	return false
}

func init() {
	proto.RegisterType((*IgmpEdmGroupsSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.vrfs.vrf.group_summary.igmp_edm_groups_summary_bag_KEYS")
	proto.RegisterType((*IgmpEdmGroupsSummaryBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.vrfs.vrf.group_summary.igmp_edm_groups_summary_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_summary_bag.proto", fileDescriptor_4f606fa9135171ea) }

var fileDescriptor_4f606fa9135171ea = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x85, 0xa6, 0xae, 0xba, 0x78, 0x32, 0xea, 0x40, 0xc8, 0x94, 0xc9, 0x03, 0x2d,
	0x23, 0x62, 0x40, 0x4c, 0x7c, 0x0c, 0x61, 0x62, 0x3a, 0x39, 0xa9, 0x63, 0x59, 0xc2, 0x39, 0xcb,
	0x97, 0xa4, 0xe4, 0xa7, 0xf0, 0x6f, 0x51, 0xdc, 0x76, 0xcd, 0x62, 0xd9, 0xcf, 0xa3, 0xf7, 0xf5,
	0xe9, 0xf8, 0xbd, 0xb3, 0x3e, 0x80, 0x39, 0x7a, 0xb0, 0x11, 0x87, 0x40, 0x40, 0x83, 0xf7, 0x3a,
	0x4e, 0x50, 0x6b, 0xab, 0x42, 0xc4, 0x1e, 0xc5, 0x73, 0xe3, 0xa8, 0x41, 0x70, 0x48, 0xf0, 0x1b,
	0xc1, 0x85, 0xf1, 0x00, 0x29, 0x84, 0xc1, 0x44, 0x35, 0xdf, 0x14, 0xf5, 0xba, 0x3b, 0xd6, 0x93,
	0x1a, 0x63, 0x4b, 0xf3, 0xa1, 0x52, 0xd7, 0xb5, 0xaa, 0x78, 0xe2, 0xf9, 0xc2, 0x2f, 0xf0, 0xf6,
	0xfa, 0xfd, 0x25, 0x6e, 0x79, 0x36, 0xc6, 0x16, 0x3a, 0xed, 0x8d, 0x64, 0x39, 0x2b, 0xd7, 0xd5,
	0x6a, 0x8c, 0xed, 0xa7, 0xf6, 0xa6, 0xf8, 0x63, 0x7c, 0xb7, 0x90, 0x17, 0x92, 0xaf, 0x66, 0xda,
	0x1b, 0x92, 0x0f, 0x39, 0x2b, 0xb7, 0xd5, 0xf5, 0x29, 0x76, 0x7c, 0x4d, 0x16, 0x2e, 0x6e, 0x9f,
	0x5c, 0x46, 0xb6, 0x3a, 0xcb, 0x3b, 0xbe, 0x39, 0x8f, 0xd9, 0xe0, 0xd0, 0xf5, 0xf2, 0x90, 0x34,
	0x4f, 0xe8, 0x65, 0x26, 0xa2, 0xe0, 0x5b, 0x47, 0xf0, 0x83, 0x27, 0xf0, 0xc6, 0x63, 0x9c, 0xe4,
	0x63, 0xce, 0xca, 0xac, 0xda, 0x38, 0x7a, 0xc7, 0xd3, 0x47, 0x42, 0xf5, 0x4d, 0x5a, 0xd1, 0xfe,
	0x3f, 0x00, 0x00, 0xff, 0xff, 0xae, 0x15, 0x81, 0x99, 0x47, 0x01, 0x00, 0x00,
}
