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
// source: igmp_edm_ssm_map_detail_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_vrfs_vrf_ssm_map_details_ssm_map_detail

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

type IgmpEdmSsmMapDetailBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	SsmMapType           string   `protobuf:"bytes,2,opt,name=ssm_map_type,json=ssmMapType,proto3" json:"ssm_map_type,omitempty"`
	GroupAddress         string   `protobuf:"bytes,3,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmSsmMapDetailBag_KEYS) Reset()         { *m = IgmpEdmSsmMapDetailBag_KEYS{} }
func (m *IgmpEdmSsmMapDetailBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSsmMapDetailBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmSsmMapDetailBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d4f289c0843057, []int{0}
}

func (m *IgmpEdmSsmMapDetailBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmSsmMapDetailBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSsmMapDetailBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmSsmMapDetailBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS.Size(m)
}
func (m *IgmpEdmSsmMapDetailBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSsmMapDetailBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmSsmMapDetailBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IgmpEdmSsmMapDetailBag_KEYS) GetSsmMapType() string {
	if m != nil {
		return m.SsmMapType
	}
	return ""
}

func (m *IgmpEdmSsmMapDetailBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

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
	return fileDescriptor_c1d4f289c0843057, []int{1}
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

type IgmpEdmSsmMapBag struct {
	GroupAddressXr       *IgmpAddrtype `protobuf:"bytes,1,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	MapType              uint32        `protobuf:"varint,2,opt,name=map_type,json=mapType,proto3" json:"map_type,omitempty"`
	SourceCounts         uint32        `protobuf:"varint,3,opt,name=source_counts,json=sourceCounts,proto3" json:"source_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpEdmSsmMapBag) Reset()         { *m = IgmpEdmSsmMapBag{} }
func (m *IgmpEdmSsmMapBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSsmMapBag) ProtoMessage()    {}
func (*IgmpEdmSsmMapBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d4f289c0843057, []int{2}
}

func (m *IgmpEdmSsmMapBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSsmMapBag.Unmarshal(m, b)
}
func (m *IgmpEdmSsmMapBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSsmMapBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSsmMapBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSsmMapBag.Merge(m, src)
}
func (m *IgmpEdmSsmMapBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSsmMapBag.Size(m)
}
func (m *IgmpEdmSsmMapBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSsmMapBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSsmMapBag proto.InternalMessageInfo

func (m *IgmpEdmSsmMapBag) GetGroupAddressXr() *IgmpAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *IgmpEdmSsmMapBag) GetMapType() uint32 {
	if m != nil {
		return m.MapType
	}
	return 0
}

func (m *IgmpEdmSsmMapBag) GetSourceCounts() uint32 {
	if m != nil {
		return m.SourceCounts
	}
	return 0
}

type IgmpEdmSsmMapDetailBag struct {
	MapInfo              *IgmpEdmSsmMapBag `protobuf:"bytes,50,opt,name=map_info,json=mapInfo,proto3" json:"map_info,omitempty"`
	Sources              []*IgmpAddrtype   `protobuf:"bytes,51,rep,name=sources,proto3" json:"sources,omitempty"`
	ExpirationTime       uint32            `protobuf:"varint,52,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	ResponsePending      bool              `protobuf:"varint,53,opt,name=response_pending,json=responsePending,proto3" json:"response_pending,omitempty"`
	QueryInterval        uint32            `protobuf:"varint,54,opt,name=query_interval,json=queryInterval,proto3" json:"query_interval,omitempty"`
	ElapsedTime          uint64            `protobuf:"varint,55,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IgmpEdmSsmMapDetailBag) Reset()         { *m = IgmpEdmSsmMapDetailBag{} }
func (m *IgmpEdmSsmMapDetailBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSsmMapDetailBag) ProtoMessage()    {}
func (*IgmpEdmSsmMapDetailBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1d4f289c0843057, []int{3}
}

func (m *IgmpEdmSsmMapDetailBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag.Unmarshal(m, b)
}
func (m *IgmpEdmSsmMapDetailBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSsmMapDetailBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSsmMapDetailBag.Merge(m, src)
}
func (m *IgmpEdmSsmMapDetailBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSsmMapDetailBag.Size(m)
}
func (m *IgmpEdmSsmMapDetailBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSsmMapDetailBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSsmMapDetailBag proto.InternalMessageInfo

func (m *IgmpEdmSsmMapDetailBag) GetMapInfo() *IgmpEdmSsmMapBag {
	if m != nil {
		return m.MapInfo
	}
	return nil
}

func (m *IgmpEdmSsmMapDetailBag) GetSources() []*IgmpAddrtype {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IgmpEdmSsmMapDetailBag) GetExpirationTime() uint32 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *IgmpEdmSsmMapDetailBag) GetResponsePending() bool {
	if m != nil {
		return m.ResponsePending
	}
	return false
}

func (m *IgmpEdmSsmMapDetailBag) GetQueryInterval() uint32 {
	if m != nil {
		return m.QueryInterval
	}
	return 0
}

func (m *IgmpEdmSsmMapDetailBag) GetElapsedTime() uint64 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmSsmMapDetailBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.ssm_map_details.ssm_map_detail.igmp_edm_ssm_map_detail_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.ssm_map_details.ssm_map_detail.igmp_addrtype")
	proto.RegisterType((*IgmpEdmSsmMapBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.ssm_map_details.ssm_map_detail.igmp_edm_ssm_map_bag")
	proto.RegisterType((*IgmpEdmSsmMapDetailBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.vrfs.vrf.ssm_map_details.ssm_map_detail.igmp_edm_ssm_map_detail_bag")
}

func init() { proto.RegisterFile("igmp_edm_ssm_map_detail_bag.proto", fileDescriptor_c1d4f289c0843057) }

var fileDescriptor_c1d4f289c0843057 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xbd, 0x8e, 0xd3, 0x40,
	0x14, 0x85, 0x65, 0x82, 0x36, 0xcb, 0x24, 0xce, 0xae, 0x46, 0x48, 0x18, 0xd1, 0x78, 0x83, 0x10,
	0xa1, 0x71, 0xb1, 0xbb, 0x84, 0x1a, 0x21, 0x8a, 0x15, 0xe2, 0x47, 0x66, 0x0b, 0x68, 0xb8, 0x9a,
	0xd8, 0xd7, 0xd6, 0x48, 0x99, 0x1f, 0x66, 0x1c, 0x2b, 0xe9, 0x28, 0x28, 0x78, 0x24, 0x5e, 0x87,
	0x37, 0x41, 0x33, 0x63, 0x6b, 0x71, 0x90, 0x52, 0xc1, 0x36, 0x51, 0xfc, 0xf9, 0xce, 0x39, 0xf7,
	0x1c, 0xdb, 0xe4, 0x8c, 0xd7, 0x42, 0x03, 0x96, 0x02, 0xac, 0x15, 0x20, 0x98, 0x86, 0x12, 0x1b,
	0xc6, 0xd7, 0xb0, 0x62, 0x75, 0xa6, 0x8d, 0x6a, 0x14, 0x7d, 0x5f, 0x70, 0x5b, 0x28, 0xe0, 0xca,
	0xc2, 0xd6, 0x00, 0xd7, 0xed, 0x25, 0xf8, 0x43, 0x4a, 0xa3, 0xc9, 0xdc, 0xbf, 0x8c, 0x15, 0x0d,
	0x6f, 0x31, 0x6b, 0x4d, 0x65, 0xdd, 0x4f, 0x36, 0x94, 0xb2, 0x7b, 0xd7, 0xf3, 0xef, 0x11, 0x49,
	0x0f, 0xd8, 0xc2, 0x9b, 0xd7, 0x9f, 0x3f, 0xd2, 0x87, 0xe4, 0xb8, 0x35, 0x15, 0x48, 0x26, 0x30,
	0x89, 0xd2, 0x68, 0x71, 0x2f, 0x1f, 0xb7, 0xa6, 0x7a, 0xc7, 0x04, 0xd2, 0x94, 0x4c, 0xfb, 0x53,
	0xcd, 0x4e, 0x63, 0x72, 0xc7, 0xdf, 0x26, 0xd6, 0x8a, 0xb7, 0x4c, 0x5f, 0xef, 0x34, 0xd2, 0xc7,
	0x24, 0xae, 0x8d, 0xda, 0x68, 0x60, 0x65, 0x69, 0xd0, 0xda, 0x64, 0xe4, 0x47, 0xa6, 0x1e, 0xbe,
	0x0c, 0x6c, 0x2e, 0x49, 0xec, 0xb7, 0x70, 0x33, 0x4e, 0x87, 0x3e, 0x20, 0x63, 0x36, 0x70, 0x3c,
	0x62, 0xc1, 0xf0, 0x8c, 0x4c, 0x7d, 0xec, 0x5e, 0x2d, 0x18, 0x4e, 0x1c, 0xeb, 0xc4, 0xba, 0x91,
	0xe5, 0x9e, 0xa1, 0x1b, 0x59, 0xf6, 0x7e, 0xbf, 0x22, 0x72, 0xff, 0xaf, 0xd8, 0x2b, 0x56, 0xd3,
	0x1f, 0x11, 0x39, 0x1d, 0xac, 0x0b, 0x5b, 0xe3, 0x37, 0x98, 0x9c, 0x7f, 0xc9, 0xfe, 0x71, 0xf9,
	0xd9, 0x20, 0x72, 0x3e, 0xfb, 0xb3, 0x91, 0x4f, 0xc6, 0xb5, 0x3e, 0xa8, 0x35, 0xce, 0xc7, 0xe2,
	0xa6, 0x53, 0xab, 0x36, 0xa6, 0x40, 0x28, 0xd4, 0x46, 0x36, 0x21, 0x62, 0x9c, 0x4f, 0x03, 0x7c,
	0xe5, 0xd9, 0xfc, 0xe7, 0x88, 0x3c, 0x3a, 0xf0, 0x68, 0xe9, 0xb7, 0x28, 0x18, 0x70, 0x59, 0xa9,
	0xe4, 0xdc, 0x47, 0xc4, 0xff, 0x13, 0x71, 0xaf, 0x64, 0x9f, 0xe3, 0x4a, 0x56, 0x8a, 0x6e, 0xc9,
	0x38, 0xac, 0x6c, 0x93, 0x8b, 0x74, 0x74, 0x0b, 0x1d, 0xf7, 0x76, 0xf4, 0x29, 0x39, 0xc1, 0xad,
	0xe6, 0x86, 0x35, 0x5c, 0x49, 0x68, 0xb8, 0xc0, 0xe4, 0xd2, 0x77, 0x38, 0xbb, 0xc1, 0xd7, 0x5c,
	0x20, 0x7d, 0x46, 0x4e, 0x0d, 0x5a, 0xad, 0xa4, 0x45, 0xd0, 0x28, 0x4b, 0x2e, 0xeb, 0xe4, 0x79,
	0x1a, 0x2d, 0x8e, 0xf3, 0x93, 0x9e, 0x7f, 0x08, 0x98, 0x3e, 0x21, 0xb3, 0xaf, 0x1b, 0x34, 0x3b,
	0xe0, 0xb2, 0x41, 0xd3, 0xb2, 0x75, 0xb2, 0xf4, 0x92, 0xb1, 0xa7, 0x57, 0x1d, 0x74, 0xaf, 0x27,
	0xae, 0x99, 0xb6, 0x58, 0x06, 0xdf, 0x17, 0x69, 0xb4, 0xb8, 0x9b, 0x4f, 0x3a, 0xe6, 0x4c, 0x57,
	0x47, 0xfe, 0x6b, 0xbf, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xcc, 0xd6, 0x65, 0x12, 0x04,
	0x00, 0x00,
}