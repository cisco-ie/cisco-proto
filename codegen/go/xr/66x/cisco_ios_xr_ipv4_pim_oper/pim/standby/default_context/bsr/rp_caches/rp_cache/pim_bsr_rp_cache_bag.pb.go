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
// source: pim_bsr_rp_cache_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_bsr_rp_caches_rp_cache

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

type PimBsrRpCacheBag_KEYS struct {
	GroupPrefix          string   `protobuf:"bytes,1,opt,name=group_prefix,json=groupPrefix,proto3" json:"group_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBsrRpCacheBag_KEYS) Reset()         { *m = PimBsrRpCacheBag_KEYS{} }
func (m *PimBsrRpCacheBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBsrRpCacheBag_KEYS) ProtoMessage()    {}
func (*PimBsrRpCacheBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d28aea04a03df8b, []int{0}
}

func (m *PimBsrRpCacheBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrRpCacheBag_KEYS.Unmarshal(m, b)
}
func (m *PimBsrRpCacheBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrRpCacheBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBsrRpCacheBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrRpCacheBag_KEYS.Merge(m, src)
}
func (m *PimBsrRpCacheBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBsrRpCacheBag_KEYS.Size(m)
}
func (m *PimBsrRpCacheBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrRpCacheBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrRpCacheBag_KEYS proto.InternalMessageInfo

func (m *PimBsrRpCacheBag_KEYS) GetGroupPrefix() string {
	if m != nil {
		return m.GroupPrefix
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
	return fileDescriptor_2d28aea04a03df8b, []int{1}
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

type PimBsrCrpBagItem struct {
	CandidateRpAddress   *PimAddrtype `protobuf:"bytes,1,opt,name=candidate_rp_address,json=candidateRpAddress,proto3" json:"candidate_rp_address,omitempty"`
	CandidateRpHoldtime  uint32       `protobuf:"varint,2,opt,name=candidate_rp_holdtime,json=candidateRpHoldtime,proto3" json:"candidate_rp_holdtime,omitempty"`
	CandidateRpPriority  uint32       `protobuf:"varint,3,opt,name=candidate_rp_priority,json=candidateRpPriority,proto3" json:"candidate_rp_priority,omitempty"`
	CandidateRpUpTime    uint32       `protobuf:"varint,4,opt,name=candidate_rp_up_time,json=candidateRpUpTime,proto3" json:"candidate_rp_up_time,omitempty"`
	CandidateRpExpires   uint32       `protobuf:"varint,5,opt,name=candidate_rp_expires,json=candidateRpExpires,proto3" json:"candidate_rp_expires,omitempty"`
	Protocol             string       `protobuf:"bytes,6,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimBsrCrpBagItem) Reset()         { *m = PimBsrCrpBagItem{} }
func (m *PimBsrCrpBagItem) String() string { return proto.CompactTextString(m) }
func (*PimBsrCrpBagItem) ProtoMessage()    {}
func (*PimBsrCrpBagItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d28aea04a03df8b, []int{2}
}

func (m *PimBsrCrpBagItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrCrpBagItem.Unmarshal(m, b)
}
func (m *PimBsrCrpBagItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrCrpBagItem.Marshal(b, m, deterministic)
}
func (m *PimBsrCrpBagItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrCrpBagItem.Merge(m, src)
}
func (m *PimBsrCrpBagItem) XXX_Size() int {
	return xxx_messageInfo_PimBsrCrpBagItem.Size(m)
}
func (m *PimBsrCrpBagItem) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrCrpBagItem.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrCrpBagItem proto.InternalMessageInfo

func (m *PimBsrCrpBagItem) GetCandidateRpAddress() *PimAddrtype {
	if m != nil {
		return m.CandidateRpAddress
	}
	return nil
}

func (m *PimBsrCrpBagItem) GetCandidateRpHoldtime() uint32 {
	if m != nil {
		return m.CandidateRpHoldtime
	}
	return 0
}

func (m *PimBsrCrpBagItem) GetCandidateRpPriority() uint32 {
	if m != nil {
		return m.CandidateRpPriority
	}
	return 0
}

func (m *PimBsrCrpBagItem) GetCandidateRpUpTime() uint32 {
	if m != nil {
		return m.CandidateRpUpTime
	}
	return 0
}

func (m *PimBsrCrpBagItem) GetCandidateRpExpires() uint32 {
	if m != nil {
		return m.CandidateRpExpires
	}
	return 0
}

func (m *PimBsrCrpBagItem) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type PimBsrCrpBagEntry struct {
	PimBsrCrpBag         []*PimBsrCrpBagItem `protobuf:"bytes,1,rep,name=pim_bsr_crp_bag,json=pimBsrCrpBag,proto3" json:"pim_bsr_crp_bag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PimBsrCrpBagEntry) Reset()         { *m = PimBsrCrpBagEntry{} }
func (m *PimBsrCrpBagEntry) String() string { return proto.CompactTextString(m) }
func (*PimBsrCrpBagEntry) ProtoMessage()    {}
func (*PimBsrCrpBagEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d28aea04a03df8b, []int{3}
}

func (m *PimBsrCrpBagEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrCrpBagEntry.Unmarshal(m, b)
}
func (m *PimBsrCrpBagEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrCrpBagEntry.Marshal(b, m, deterministic)
}
func (m *PimBsrCrpBagEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrCrpBagEntry.Merge(m, src)
}
func (m *PimBsrCrpBagEntry) XXX_Size() int {
	return xxx_messageInfo_PimBsrCrpBagEntry.Size(m)
}
func (m *PimBsrCrpBagEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrCrpBagEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrCrpBagEntry proto.InternalMessageInfo

func (m *PimBsrCrpBagEntry) GetPimBsrCrpBag() []*PimBsrCrpBagItem {
	if m != nil {
		return m.PimBsrCrpBag
	}
	return nil
}

type PimBsrRpCacheBag struct {
	GroupPrefixXr         *PimAddrtype       `protobuf:"bytes,50,opt,name=group_prefix_xr,json=groupPrefixXr,proto3" json:"group_prefix_xr,omitempty"`
	GroupPrefixLength     uint32             `protobuf:"varint,51,opt,name=group_prefix_length,json=groupPrefixLength,proto3" json:"group_prefix_length,omitempty"`
	CandidateRpGroupCount uint32             `protobuf:"varint,52,opt,name=candidate_rp_group_count,json=candidateRpGroupCount,proto3" json:"candidate_rp_group_count,omitempty"`
	CandidateRpList       *PimBsrCrpBagEntry `protobuf:"bytes,53,opt,name=candidate_rp_list,json=candidateRpList,proto3" json:"candidate_rp_list,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *PimBsrRpCacheBag) Reset()         { *m = PimBsrRpCacheBag{} }
func (m *PimBsrRpCacheBag) String() string { return proto.CompactTextString(m) }
func (*PimBsrRpCacheBag) ProtoMessage()    {}
func (*PimBsrRpCacheBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d28aea04a03df8b, []int{4}
}

func (m *PimBsrRpCacheBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrRpCacheBag.Unmarshal(m, b)
}
func (m *PimBsrRpCacheBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrRpCacheBag.Marshal(b, m, deterministic)
}
func (m *PimBsrRpCacheBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrRpCacheBag.Merge(m, src)
}
func (m *PimBsrRpCacheBag) XXX_Size() int {
	return xxx_messageInfo_PimBsrRpCacheBag.Size(m)
}
func (m *PimBsrRpCacheBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrRpCacheBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrRpCacheBag proto.InternalMessageInfo

func (m *PimBsrRpCacheBag) GetGroupPrefixXr() *PimAddrtype {
	if m != nil {
		return m.GroupPrefixXr
	}
	return nil
}

func (m *PimBsrRpCacheBag) GetGroupPrefixLength() uint32 {
	if m != nil {
		return m.GroupPrefixLength
	}
	return 0
}

func (m *PimBsrRpCacheBag) GetCandidateRpGroupCount() uint32 {
	if m != nil {
		return m.CandidateRpGroupCount
	}
	return 0
}

func (m *PimBsrRpCacheBag) GetCandidateRpList() *PimBsrCrpBagEntry {
	if m != nil {
		return m.CandidateRpList
	}
	return nil
}

func init() {
	proto.RegisterType((*PimBsrRpCacheBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.rp_caches.rp_cache.pim_bsr_rp_cache_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.rp_caches.rp_cache.pim_addrtype")
	proto.RegisterType((*PimBsrCrpBagItem)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.rp_caches.rp_cache.pim_bsr_crp_bag_item")
	proto.RegisterType((*PimBsrCrpBagEntry)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.rp_caches.rp_cache.pim_bsr_crp_bag_entry")
	proto.RegisterType((*PimBsrRpCacheBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.rp_caches.rp_cache.pim_bsr_rp_cache_bag")
}

func init() { proto.RegisterFile("pim_bsr_rp_cache_bag.proto", fileDescriptor_2d28aea04a03df8b) }

var fileDescriptor_2d28aea04a03df8b = []byte{
	// 490 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x18, 0x14, 0x70, 0x37, 0x55, 0xf3, 0x56, 0x61, 0x76, 0x1a, 0x39, 0xed, 0x14, 0x50,
	0x37, 0xc6, 0x0d, 0x89, 0x4d, 0x13, 0x48, 0x0c, 0x34, 0x05, 0x90, 0x40, 0x1c, 0x2c, 0xc7, 0x71,
	0x5b, 0x4b, 0x49, 0x6c, 0x3d, 0xbb, 0x28, 0x3d, 0x73, 0x43, 0x9a, 0xc4, 0x91, 0x1f, 0xc0, 0x0f,
	0x45, 0x76, 0xd2, 0xc8, 0x09, 0x1c, 0xe9, 0xad, 0x7d, 0xef, 0xfb, 0xde, 0x7b, 0x9f, 0xbf, 0x2f,
	0xe8, 0x48, 0xcb, 0x92, 0x66, 0x06, 0x28, 0x68, 0xca, 0x19, 0x5f, 0x0a, 0x9a, 0xb1, 0x45, 0xa2,
	0x41, 0x59, 0x85, 0xdf, 0x71, 0x69, 0xb8, 0xa2, 0x52, 0x19, 0x5a, 0x03, 0x95, 0xfa, 0xdb, 0x19,
	0x75, 0x68, 0xa5, 0x05, 0x24, 0x5a, 0x96, 0x89, 0xb1, 0xac, 0xca, 0xb3, 0x75, 0x92, 0x8b, 0x39,
	0x5b, 0x15, 0x96, 0x72, 0x55, 0x59, 0x51, 0xdb, 0x24, 0x33, 0x90, 0x6c, 0xc6, 0x99, 0xee, 0x57,
	0xfc, 0x12, 0x3d, 0xfe, 0xd7, 0x32, 0xfa, 0xf6, 0xea, 0xcb, 0x07, 0xfc, 0x04, 0xed, 0x2e, 0x40,
	0xad, 0x34, 0xd5, 0x20, 0xe6, 0xb2, 0x26, 0xd1, 0x71, 0x74, 0xf2, 0x30, 0x1d, 0xfb, 0xda, 0x8d,
	0x2f, 0xc5, 0x25, 0xda, 0x75, 0x7c, 0x96, 0xe7, 0x60, 0xd7, 0x5a, 0xe0, 0x47, 0xe8, 0x3e, 0x9b,
	0xd3, 0x8a, 0x95, 0xa2, 0x45, 0x8f, 0xd8, 0xfc, 0x3d, 0x2b, 0x85, 0x9b, 0xe5, 0x8f, 0x75, 0x48,
	0x61, 0x0c, 0xb9, 0xd3, 0xcc, 0x72, 0xb5, 0x57, 0x4d, 0xa9, 0x85, 0x9c, 0x77, 0x90, 0x9d, 0x0e,
	0x72, 0xde, 0x42, 0xe2, 0xdb, 0x1d, 0x74, 0xb8, 0xb9, 0x97, 0x83, 0xf6, 0xa7, 0x4a, 0x2b, 0x4a,
	0x7c, 0x1b, 0xa1, 0x43, 0xce, 0xaa, 0x5c, 0xe6, 0xcc, 0x0a, 0x27, 0x65, 0x33, 0xc4, 0x5d, 0x31,
	0x9e, 0x7d, 0x4d, 0xfe, 0xeb, 0xb3, 0x25, 0xa1, 0xe6, 0x14, 0x77, 0x8b, 0x53, 0xbd, 0xd1, 0x32,
	0x43, 0xd3, 0xde, 0x39, 0x4b, 0x55, 0xe4, 0x56, 0x96, 0xc2, 0xeb, 0xde, 0x4b, 0x0f, 0x02, 0xca,
	0x9b, 0xb6, 0xf5, 0x17, 0x47, 0x83, 0x54, 0x20, 0xed, 0xda, 0x3f, 0x44, 0x9f, 0x73, 0xd3, 0xb6,
	0xf0, 0xd3, 0x81, 0xec, 0x95, 0xa6, 0x7e, 0xcd, 0x5d, 0x4f, 0xd9, 0x0f, 0x28, 0x9f, 0xf4, 0x47,
	0xb7, 0xe4, 0xd9, 0x80, 0x20, 0x6a, 0x2d, 0x41, 0x18, 0x72, 0xcf, 0x13, 0x42, 0x29, 0x57, 0x4d,
	0x07, 0x1f, 0xa1, 0x07, 0x3e, 0x7a, 0x5c, 0x15, 0x64, 0xe4, 0x2d, 0xe9, 0xfe, 0xc7, 0xbf, 0x23,
	0x34, 0x1d, 0xfa, 0x21, 0x2a, 0x0b, 0x6b, 0xfc, 0x23, 0x42, 0x93, 0x41, 0x87, 0x44, 0xc7, 0x3b,
	0x27, 0xe3, 0x19, 0xdf, 0x82, 0x17, 0xc3, 0x3c, 0xa4, 0x2e, 0x95, 0x17, 0x06, 0x2e, 0x41, 0x5f,
	0xb0, 0x45, 0xfc, 0x2b, 0x88, 0x4d, 0x18, 0x73, 0xfc, 0x3d, 0x42, 0x93, 0x30, 0xe2, 0xb4, 0x06,
	0x32, 0xdb, 0x7e, 0x62, 0xf6, 0x82, 0x4f, 0xe8, 0x33, 0xe0, 0x04, 0x1d, 0xf4, 0x8e, 0x28, 0x44,
	0xb5, 0xb0, 0x4b, 0x72, 0xda, 0x78, 0x18, 0x60, 0xaf, 0x7d, 0x03, 0xbf, 0x40, 0xa4, 0xe7, 0x61,
	0x43, 0xe6, 0x6a, 0x55, 0x59, 0x72, 0xe6, 0x49, 0xd3, 0xc0, 0xc7, 0xd7, 0xae, 0x7b, 0xe9, 0x9a,
	0xf8, 0x67, 0x84, 0xf6, 0x7b, 0xcc, 0x42, 0x1a, 0x4b, 0x9e, 0x7b, 0xc1, 0xf9, 0x96, 0x6d, 0xf1,
	0xb1, 0x48, 0x27, 0xc1, 0x61, 0xd7, 0xd2, 0xd8, 0x6c, 0xe4, 0xb3, 0x74, 0xfa, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x46, 0x20, 0xc8, 0xf4, 0x04, 0x00, 0x00,
}
