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
// source: igmp_edm_not_active_allgroups_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_default_context_non_active_groups

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

type IgmpEdmNotActiveAllgroupsBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) Reset()         { *m = IgmpEdmNotActiveAllgroupsBag_KEYS{} }
func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNotActiveAllgroupsBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmNotActiveAllgroupsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeadb3d1969200b1, []int{0}
}

func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS.Size(m)
}
func (m *IgmpEdmNotActiveAllgroupsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag_KEYS proto.InternalMessageInfo

type IgmpAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpAddrtype) Reset()         { *m = IgmpAddrtype{} }
func (m *IgmpAddrtype) String() string { return proto.CompactTextString(m) }
func (*IgmpAddrtype) ProtoMessage()    {}
func (*IgmpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeadb3d1969200b1, []int{1}
}

func (m *IgmpAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpAddrtype.Unmarshal(m, b)
}
func (m *IgmpAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpAddrtype.Marshal(b, m, deterministic)
}
func (m *IgmpAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpAddrtype.Merge(m, src)
}
func (m *IgmpAddrtype) XXX_Size() int {
	return xxx_messageInfo_IgmpAddrtype.Size(m)
}
func (m *IgmpAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpAddrtype proto.InternalMessageInfo

func (m *IgmpAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type IgmpEdmNotActiveGroupBag struct {
	Interface            string        `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	ReasonForNonActivity string        `protobuf:"bytes,2,opt,name=reason_for_non_activity,json=reasonForNonActivity,proto3" json:"reason_for_non_activity,omitempty"`
	GroupAddress         *IgmpAddrtype `protobuf:"bytes,3,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	SourceAddress        *IgmpAddrtype `protobuf:"bytes,4,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpEdmNotActiveGroupBag) Reset()         { *m = IgmpEdmNotActiveGroupBag{} }
func (m *IgmpEdmNotActiveGroupBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNotActiveGroupBag) ProtoMessage()    {}
func (*IgmpEdmNotActiveGroupBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeadb3d1969200b1, []int{2}
}

func (m *IgmpEdmNotActiveGroupBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNotActiveGroupBag.Unmarshal(m, b)
}
func (m *IgmpEdmNotActiveGroupBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNotActiveGroupBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNotActiveGroupBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNotActiveGroupBag.Merge(m, src)
}
func (m *IgmpEdmNotActiveGroupBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNotActiveGroupBag.Size(m)
}
func (m *IgmpEdmNotActiveGroupBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNotActiveGroupBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNotActiveGroupBag proto.InternalMessageInfo

func (m *IgmpEdmNotActiveGroupBag) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *IgmpEdmNotActiveGroupBag) GetReasonForNonActivity() string {
	if m != nil {
		return m.ReasonForNonActivity
	}
	return ""
}

func (m *IgmpEdmNotActiveGroupBag) GetGroupAddress() *IgmpAddrtype {
	if m != nil {
		return m.GroupAddress
	}
	return nil
}

func (m *IgmpEdmNotActiveGroupBag) GetSourceAddress() *IgmpAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

type IgmpEdmNotActiveAllgroupsBag struct {
	NonActiveGroups      []*IgmpEdmNotActiveGroupBag `protobuf:"bytes,50,rep,name=non_active_groups,json=nonActiveGroups,proto3" json:"non_active_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *IgmpEdmNotActiveAllgroupsBag) Reset()         { *m = IgmpEdmNotActiveAllgroupsBag{} }
func (m *IgmpEdmNotActiveAllgroupsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmNotActiveAllgroupsBag) ProtoMessage()    {}
func (*IgmpEdmNotActiveAllgroupsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeadb3d1969200b1, []int{3}
}

func (m *IgmpEdmNotActiveAllgroupsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag.Unmarshal(m, b)
}
func (m *IgmpEdmNotActiveAllgroupsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmNotActiveAllgroupsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag.Merge(m, src)
}
func (m *IgmpEdmNotActiveAllgroupsBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag.Size(m)
}
func (m *IgmpEdmNotActiveAllgroupsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmNotActiveAllgroupsBag proto.InternalMessageInfo

func (m *IgmpEdmNotActiveAllgroupsBag) GetNonActiveGroups() []*IgmpEdmNotActiveGroupBag {
	if m != nil {
		return m.NonActiveGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*IgmpEdmNotActiveAllgroupsBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.non_active_groups.igmp_edm_not_active_allgroups_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.non_active_groups.igmp_addrtype")
	proto.RegisterType((*IgmpEdmNotActiveGroupBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.non_active_groups.igmp_edm_not_active_group_bag")
	proto.RegisterType((*IgmpEdmNotActiveAllgroupsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.non_active_groups.igmp_edm_not_active_allgroups_bag")
}

func init() {
	proto.RegisterFile("igmp_edm_not_active_allgroups_bag.proto", fileDescriptor_eeadb3d1969200b1)
}

var fileDescriptor_eeadb3d1969200b1 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0x56, 0x2a, 0x4d, 0x5b, 0xc5, 0x20, 0xb4, 0x0b, 0x85, 0xb6, 0x0b, 0xed, 0x2a,
	0x8b, 0xaa, 0xdd, 0x77, 0xe1, 0x0f, 0x08, 0x5d, 0xd4, 0x85, 0xb8, 0xba, 0xa4, 0x33, 0x99, 0x21,
	0x30, 0x93, 0x3b, 0x24, 0x69, 0x69, 0x41, 0x7c, 0x09, 0x9f, 0xc6, 0xd7, 0xf2, 0x09, 0x64, 0x32,
	0x3f, 0x2a, 0x2a, 0xdd, 0xd4, 0xed, 0xb9, 0x27, 0xf7, 0x7c, 0x73, 0x32, 0x21, 0xe7, 0x32, 0x4a,
	0x52, 0x10, 0x41, 0x02, 0x0a, 0x2d, 0x70, 0xdf, 0xca, 0x95, 0x00, 0x1e, 0xc7, 0x91, 0xc6, 0x65,
	0x6a, 0x60, 0xc1, 0x23, 0x96, 0x6a, 0xb4, 0x48, 0xef, 0x7c, 0x69, 0x7c, 0x04, 0x89, 0x06, 0xd6,
	0x1a, 0x64, 0xba, 0xba, 0x04, 0x77, 0x14, 0x53, 0xa1, 0x59, 0x12, 0x07, 0x2c, 0x3f, 0xcb, 0x02,
	0x11, 0xf2, 0x65, 0x6c, 0xc1, 0x47, 0x65, 0xc5, 0xda, 0x32, 0x85, 0xaa, 0x5c, 0x9b, 0xef, 0x1c,
	0x8e, 0xc8, 0xd9, 0xd6, 0x50, 0xb8, 0xbf, 0x7e, 0x7a, 0x18, 0x2a, 0xd2, 0x71, 0x4e, 0x1e, 0x04,
	0xda, 0x6e, 0x52, 0x41, 0xbb, 0x64, 0x9f, 0x87, 0xa0, 0x78, 0x22, 0x7a, 0x5e, 0xdf, 0x1b, 0x35,
	0xe7, 0x0d, 0x1e, 0xce, 0x78, 0x22, 0xe8, 0x80, 0xb4, 0x1d, 0x52, 0xe6, 0x14, 0xc6, 0xf4, 0x6a,
	0x6e, 0xda, 0xca, 0xb4, 0x69, 0x2e, 0x15, 0x96, 0x49, 0x65, 0xa9, 0x57, 0x96, 0x49, 0x61, 0x19,
	0xbe, 0xd7, 0xc8, 0xe9, 0x6f, 0x68, 0x8e, 0x2b, 0xc3, 0xa2, 0x27, 0xa4, 0x29, 0x95, 0x15, 0x3a,
	0xe4, 0x7e, 0x89, 0xf0, 0x29, 0xd0, 0x2b, 0xd2, 0xd5, 0x82, 0x1b, 0x54, 0x10, 0xa2, 0x86, 0xea,
	0xcb, 0xa5, 0xdd, 0x14, 0x40, 0xc7, 0xf9, 0xf8, 0x06, 0xf5, 0x0c, 0xd5, 0xb4, 0x98, 0xd1, 0x67,
	0xd2, 0xc9, 0x13, 0xbe, 0xa2, 0xb5, 0xc6, 0x8f, 0x6c, 0x57, 0x95, 0xb3, 0x6f, 0x2d, 0xce, 0xdb,
	0x4e, 0x2d, 0x7b, 0x79, 0x21, 0x07, 0x06, 0x97, 0xda, 0x17, 0x55, 0xfc, 0xde, 0xff, 0xc6, 0x77,
	0xf2, 0xb8, 0xb2, 0xf4, 0x37, 0x8f, 0x0c, 0xb6, 0xfe, 0x0f, 0xf4, 0xd5, 0x23, 0x47, 0x3f, 0x16,
	0xf7, 0xc6, 0xfd, 0xfa, 0xa8, 0x35, 0x8e, 0x76, 0x4c, 0xfa, 0xd7, 0xed, 0xcf, 0x0f, 0x55, 0x71,
	0x6b, 0xe2, 0xd6, 0xd9, 0x17, 0x0d, 0xf7, 0x36, 0x2e, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xea,
	0x2e, 0x13, 0xf8, 0x46, 0x03, 0x00, 0x00,
}