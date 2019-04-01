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

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_vrfs_vrf_group_summary

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
	proto.RegisterType((*IgmpEdmGroupsSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.group_summary.igmp_edm_groups_summary_bag_KEYS")
	proto.RegisterType((*IgmpEdmGroupsSummaryBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.group_summary.igmp_edm_groups_summary_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_summary_bag.proto", fileDescriptor_4f606fa9135171ea) }

var fileDescriptor_4f606fa9135171ea = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x85, 0xb6, 0xae, 0xba, 0x78, 0x32, 0xea, 0x40, 0xc8, 0x94, 0xc9, 0x03, 0x2d,
	0x23, 0x2c, 0x88, 0x89, 0x8f, 0x21, 0x4c, 0x4c, 0x27, 0x37, 0x38, 0x96, 0x25, 0xdc, 0xb3, 0xee,
	0x12, 0x97, 0xfe, 0x14, 0xfe, 0x2d, 0x8a, 0x4b, 0xd7, 0x2e, 0x96, 0xfd, 0x3c, 0x7a, 0x5f, 0x9f,
	0x4e, 0xde, 0x06, 0x1f, 0x13, 0xb8, 0xaf, 0x08, 0x9e, 0x70, 0x4c, 0x0c, 0x3c, 0xc6, 0x68, 0xe9,
	0x08, 0x3b, 0xeb, 0x4d, 0x22, 0x1c, 0x50, 0x3d, 0x76, 0x81, 0x3b, 0x84, 0x80, 0x0c, 0x3f, 0x04,
	0x21, 0xe5, 0x2d, 0x94, 0x10, 0x26, 0x47, 0x66, 0xba, 0x19, 0xdb, 0x0d, 0x21, 0x3b, 0x93, 0xa9,
	0xe7, 0xe9, 0x30, 0xa5, 0xea, 0xdc, 0x54, 0x3f, 0xc8, 0xea, 0xc2, 0x27, 0xf0, 0xf2, 0xfc, 0xf9,
	0xa1, 0xae, 0xe5, 0x3c, 0x53, 0x0f, 0x7b, 0x1b, 0x9d, 0x16, 0x95, 0x68, 0x16, 0xed, 0x2c, 0x53,
	0xff, 0x6e, 0xa3, 0xab, 0x7f, 0x85, 0x5c, 0x5f, 0xc8, 0x2b, 0x2d, 0x67, 0x13, 0x1d, 0x1c, 0xeb,
	0xbb, 0x4a, 0x34, 0xab, 0xf6, 0xfc, 0x54, 0x6b, 0xb9, 0x60, 0x0f, 0xff, 0x6e, 0x53, 0xdc, 0x9c,
	0x7d, 0x7b, 0x92, 0x37, 0x72, 0x79, 0x1a, 0xb3, 0xc3, 0x71, 0x3f, 0xe8, 0x6d, 0xd1, 0xb2, 0xa0,
	0xa7, 0x89, 0xa8, 0x5a, 0xae, 0x02, 0xc3, 0x37, 0x1e, 0x20, 0xba, 0x88, 0x74, 0xd4, 0xf7, 0x95,
	0x68, 0xe6, 0xed, 0x32, 0xf0, 0x2b, 0x1e, 0xde, 0x0a, 0xda, 0x5d, 0x95, 0x0d, 0x6d, 0xfe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xdf, 0x23, 0xc2, 0xde, 0x46, 0x01, 0x00, 0x00,
}
