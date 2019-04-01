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
// source: igmp_edm_ssm_map_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_vrfs_vrf_ssm_maps_ssm_map

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

type IgmpEdmSsmMapBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	SsmMapType           string   `protobuf:"bytes,2,opt,name=ssm_map_type,json=ssmMapType,proto3" json:"ssm_map_type,omitempty"`
	GroupAddress         string   `protobuf:"bytes,3,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmSsmMapBag_KEYS) Reset()         { *m = IgmpEdmSsmMapBag_KEYS{} }
func (m *IgmpEdmSsmMapBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSsmMapBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmSsmMapBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_51be9443925d0d07, []int{0}
}

func (m *IgmpEdmSsmMapBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmSsmMapBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmSsmMapBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmSsmMapBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmSsmMapBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmSsmMapBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmSsmMapBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmSsmMapBag_KEYS.Size(m)
}
func (m *IgmpEdmSsmMapBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmSsmMapBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmSsmMapBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmSsmMapBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IgmpEdmSsmMapBag_KEYS) GetSsmMapType() string {
	if m != nil {
		return m.SsmMapType
	}
	return ""
}

func (m *IgmpEdmSsmMapBag_KEYS) GetGroupAddress() string {
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
	return fileDescriptor_51be9443925d0d07, []int{1}
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
	GroupAddressXr       *IgmpAddrtype `protobuf:"bytes,50,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	MapType              uint32        `protobuf:"varint,51,opt,name=map_type,json=mapType,proto3" json:"map_type,omitempty"`
	SourceCounts         uint32        `protobuf:"varint,52,opt,name=source_counts,json=sourceCounts,proto3" json:"source_counts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpEdmSsmMapBag) Reset()         { *m = IgmpEdmSsmMapBag{} }
func (m *IgmpEdmSsmMapBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmSsmMapBag) ProtoMessage()    {}
func (*IgmpEdmSsmMapBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_51be9443925d0d07, []int{2}
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

func init() {
	proto.RegisterType((*IgmpEdmSsmMapBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.ssm_maps.ssm_map.igmp_edm_ssm_map_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.ssm_maps.ssm_map.igmp_addrtype")
	proto.RegisterType((*IgmpEdmSsmMapBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.vrfs.vrf.ssm_maps.ssm_map.igmp_edm_ssm_map_bag")
}

func init() { proto.RegisterFile("igmp_edm_ssm_map_bag.proto", fileDescriptor_51be9443925d0d07) }

var fileDescriptor_51be9443925d0d07 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xa9, 0xc2, 0xa6, 0x59, 0x27, 0x12, 0x04, 0x3b, 0x4f, 0x75, 0x5e, 0x76, 0xca, 0x61,
	0x1b, 0x3b, 0x2b, 0xe2, 0x49, 0x54, 0xa8, 0x1e, 0xf4, 0xf4, 0xc8, 0xda, 0xb4, 0x14, 0x4c, 0x13,
	0xf2, 0xda, 0xb2, 0x1e, 0xfc, 0x23, 0xfd, 0x8f, 0x24, 0x69, 0xab, 0x56, 0x76, 0xf3, 0x12, 0xc2,
	0xf7, 0x3e, 0xbe, 0x5f, 0xde, 0x47, 0xc8, 0x45, 0x9e, 0x49, 0x0d, 0x22, 0x91, 0x80, 0x28, 0x41,
	0x72, 0x0d, 0x5b, 0x9e, 0x31, 0x6d, 0x54, 0xa9, 0xe8, 0x75, 0x9c, 0x63, 0xac, 0x20, 0x57, 0x08,
	0x3b, 0x03, 0xb9, 0xae, 0xd7, 0xe0, 0xdc, 0x4a, 0x0b, 0xc3, 0xe4, 0x7b, 0xc2, 0x78, 0x5c, 0xe6,
	0xb5, 0x60, 0xb5, 0x49, 0xd1, 0x1e, 0xac, 0x8b, 0xc0, 0xfe, 0x32, 0xff, 0x20, 0xb3, 0x7d, 0xf9,
	0x70, 0x7f, 0xf7, 0xf6, 0x4c, 0x67, 0xe4, 0xa8, 0x36, 0x29, 0x14, 0x5c, 0x8a, 0xc0, 0x0b, 0xbd,
	0xc5, 0x71, 0x34, 0xae, 0x4d, 0xfa, 0xc8, 0xa5, 0xa0, 0x21, 0xf1, 0x7b, 0x7b, 0xd9, 0x68, 0x11,
	0x1c, 0xb8, 0x31, 0x41, 0x94, 0x0f, 0x5c, 0xbf, 0x34, 0x5a, 0xd0, 0x2b, 0x32, 0xcd, 0x8c, 0xaa,
	0x34, 0xf0, 0x24, 0x31, 0x02, 0x31, 0x38, 0x74, 0x16, 0xdf, 0x89, 0x37, 0xad, 0x36, 0x2f, 0xc8,
	0xd4, 0xe1, 0xad, 0xc7, 0xe6, 0xd0, 0x73, 0x32, 0xe6, 0x03, 0xe2, 0x88, 0xb7, 0xc0, 0x4b, 0xe2,
	0xbb, 0xfd, 0xfa, 0xb4, 0x16, 0x38, 0xb1, 0x5a, 0x17, 0xd6, 0x59, 0x36, 0x7f, 0x80, 0xd6, 0xb2,
	0xe9, 0x79, 0x9f, 0x1e, 0x39, 0xdb, 0xb7, 0x2f, 0x6d, 0xc8, 0xe9, 0xe0, 0xb5, 0xb0, 0x33, 0xc1,
	0x32, 0xf4, 0x16, 0x93, 0xe5, 0x13, 0xfb, 0x6f, 0xc9, 0x6c, 0xb0, 0x62, 0x74, 0xf2, 0xbb, 0x81,
	0x57, 0x63, 0x5b, 0xfe, 0xae, 0x71, 0x15, 0x7a, 0x8b, 0x69, 0x34, 0x96, 0x3f, 0x1d, 0xa2, 0xaa,
	0x4c, 0x2c, 0x20, 0x56, 0x55, 0x51, 0x62, 0xb0, 0x76, 0x73, 0xbf, 0x15, 0x6f, 0x9d, 0xb6, 0x1d,
	0xb9, 0xbf, 0xb0, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xb0, 0x98, 0xda, 0x29, 0x02, 0x00,
	0x00,
}