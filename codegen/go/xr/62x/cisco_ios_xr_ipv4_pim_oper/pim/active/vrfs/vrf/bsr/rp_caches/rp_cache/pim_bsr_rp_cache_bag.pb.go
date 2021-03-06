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

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_bsr_rp_caches_rp_cache

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	GroupPrefix          string   `protobuf:"bytes,2,opt,name=group_prefix,json=groupPrefix,proto3" json:"group_prefix,omitempty"`
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

func (m *PimBsrRpCacheBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	PimBsrCrpBag         []*PimBsrCrpBagItem `protobuf:"bytes,7,rep,name=pim_bsr_crp_bag,json=pimBsrCrpBag,proto3" json:"pim_bsr_crp_bag,omitempty"`
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
	proto.RegisterType((*PimBsrRpCacheBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.bsr.rp_caches.rp_cache.pim_bsr_rp_cache_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.bsr.rp_caches.rp_cache.pim_addrtype")
	proto.RegisterType((*PimBsrCrpBagItem)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.bsr.rp_caches.rp_cache.pim_bsr_crp_bag_item")
	proto.RegisterType((*PimBsrCrpBagEntry)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.bsr.rp_caches.rp_cache.pim_bsr_crp_bag_entry")
	proto.RegisterType((*PimBsrRpCacheBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.bsr.rp_caches.rp_cache.pim_bsr_rp_cache_bag")
}

func init() { proto.RegisterFile("pim_bsr_rp_cache_bag.proto", fileDescriptor_2d28aea04a03df8b) }

var fileDescriptor_2d28aea04a03df8b = []byte{
	// 491 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6a, 0x1b, 0x31,
	0x10, 0x65, 0xeb, 0xd6, 0x4e, 0xe5, 0x04, 0x13, 0x25, 0xa6, 0x4a, 0x4e, 0xe9, 0x9e, 0x72, 0xda,
	0x16, 0x27, 0x4d, 0xcf, 0x4d, 0x30, 0x2d, 0x34, 0x94, 0xb0, 0x69, 0xa1, 0xa1, 0x05, 0x21, 0x6b,
	0x65, 0x5b, 0xe0, 0x5d, 0x89, 0x91, 0xbc, 0xd8, 0xf4, 0x56, 0x08, 0xed, 0x47, 0xf4, 0x63, 0x8b,
	0x64, 0xef, 0xa2, 0xdd, 0xf6, 0x98, 0x5c, 0x8c, 0x3d, 0x6f, 0xde, 0xcc, 0x3c, 0xbd, 0x67, 0x74,
	0xac, 0x65, 0x4e, 0x27, 0x06, 0x28, 0x68, 0xca, 0x19, 0x9f, 0x0b, 0x3a, 0x61, 0xb3, 0x44, 0x83,
	0xb2, 0x0a, 0x8f, 0xb9, 0x34, 0x5c, 0x51, 0xa9, 0x0c, 0x5d, 0x01, 0x95, 0xba, 0x3c, 0xa7, 0xae,
	0x5b, 0x69, 0x01, 0x89, 0x96, 0x79, 0xc2, 0xb8, 0x95, 0xa5, 0x48, 0x4a, 0x98, 0x1a, 0xf7, 0x91,
	0x4c, 0x0c, 0x24, 0xd5, 0x18, 0x53, 0x7f, 0x8b, 0xef, 0xd0, 0xd1, 0xff, 0x96, 0xd0, 0x8f, 0xe3,
	0xbb, 0x5b, 0x7c, 0x84, 0x76, 0x4a, 0x98, 0xd2, 0x82, 0xe5, 0x82, 0x44, 0x27, 0xd1, 0xe9, 0xf3,
	0xb4, 0x57, 0xc2, 0xf4, 0x13, 0xcb, 0x05, 0x7e, 0x89, 0x76, 0x67, 0xa0, 0x96, 0x9a, 0x6a, 0x10,
	0x53, 0xb9, 0x22, 0x4f, 0x3c, 0xdc, 0xf7, 0xb5, 0x1b, 0x5f, 0x8a, 0x73, 0xb4, 0xeb, 0x46, 0xb3,
	0x2c, 0x03, 0xbb, 0xd6, 0x02, 0xbf, 0x40, 0x3d, 0xd6, 0x18, 0xd6, 0x65, 0xf5, 0x2c, 0x7f, 0xbf,
	0xeb, 0x14, 0xc6, 0x54, 0xb3, 0x5c, 0xed, 0xdd, 0xa6, 0xb4, 0x6d, 0xb9, 0xa8, 0x5b, 0x3a, 0x75,
	0xcb, 0xc5, 0xb6, 0x25, 0xbe, 0xef, 0xa0, 0xc3, 0x4a, 0x0a, 0x07, 0xed, 0x55, 0x48, 0x2b, 0x72,
	0x7c, 0x1f, 0xa1, 0x43, 0xce, 0x8a, 0x4c, 0x66, 0xcc, 0x0a, 0xa7, 0xb2, 0x1a, 0xe2, 0xae, 0xe8,
	0x8f, 0x6e, 0x93, 0x07, 0x79, 0xc9, 0x24, 0xd4, 0x9a, 0xe2, 0x7a, 0x61, 0xaa, 0x2b, 0x0d, 0x23,
	0x34, 0x6c, 0x9c, 0x31, 0x57, 0x8b, 0xcc, 0xca, 0x5c, 0x78, 0xbd, 0x7b, 0xe9, 0x41, 0x40, 0xf9,
	0xb0, 0x85, 0xfe, 0xe1, 0x68, 0x90, 0x0a, 0xa4, 0x5d, 0xfb, 0x07, 0x68, 0x72, 0x6e, 0xb6, 0x10,
	0x7e, 0xd5, 0x92, 0xbb, 0xd4, 0xd4, 0xaf, 0x79, 0xea, 0x29, 0xfb, 0x01, 0xe5, 0x8b, 0xfe, 0xec,
	0x96, 0xbc, 0x6e, 0x11, 0xc4, 0x4a, 0x4b, 0x10, 0x86, 0x3c, 0xf3, 0x84, 0x50, 0xca, 0x78, 0x83,
	0xe0, 0x63, 0xb4, 0xe3, 0x53, 0xc8, 0xd5, 0x82, 0x74, 0xbd, 0x15, 0xf5, 0xef, 0xf8, 0x4f, 0x84,
	0x86, 0x6d, 0x1f, 0x44, 0x61, 0x61, 0x8d, 0x7f, 0x46, 0x68, 0xd0, 0x42, 0x48, 0xef, 0xa4, 0x73,
	0xda, 0x1f, 0x7d, 0x7b, 0x40, 0x0f, 0xda, 0xfe, 0xa7, 0x2e, 0x85, 0x97, 0x06, 0xae, 0x40, 0x5f,
	0xb2, 0x59, 0xfc, 0x2b, 0x88, 0x49, 0x98, 0x78, 0xfc, 0x03, 0x0d, 0xc2, 0x44, 0xd3, 0x15, 0x90,
	0xd1, 0xe3, 0x05, 0x64, 0x2f, 0xf8, 0xa7, 0x7c, 0x05, 0x9c, 0xa0, 0x83, 0xc6, 0xf2, 0x85, 0x28,
	0x66, 0x76, 0x4e, 0xce, 0x36, 0x96, 0x05, 0xbd, 0xd7, 0x1e, 0xc0, 0x6f, 0x11, 0x69, 0x58, 0xb6,
	0x21, 0x73, 0xb5, 0x2c, 0x2c, 0x39, 0xf7, 0xa4, 0x61, 0x60, 0xdb, 0x7b, 0x87, 0x5e, 0x39, 0x10,
	0xff, 0x8e, 0xd0, 0x7e, 0x83, 0xb9, 0x90, 0xc6, 0x92, 0x37, 0x5e, 0xe8, 0xf7, 0x47, 0x72, 0xc1,
	0xbb, 0x9f, 0x0e, 0x82, 0x83, 0xae, 0xa5, 0xb1, 0x93, 0xae, 0x8f, 0xcc, 0xd9, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x75, 0x6b, 0x12, 0xa4, 0xe6, 0x04, 0x00, 0x00,
}
