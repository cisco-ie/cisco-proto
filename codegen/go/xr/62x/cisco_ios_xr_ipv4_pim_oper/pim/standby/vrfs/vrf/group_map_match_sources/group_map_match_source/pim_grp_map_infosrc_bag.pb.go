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
// source: pim_grp_map_infosrc_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_group_map_match_sources_group_map_match_source

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

type PimGrpMapInfosrcBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	GroupAddress         string   `protobuf:"bytes,2,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimGrpMapInfosrcBag_KEYS) Reset()         { *m = PimGrpMapInfosrcBag_KEYS{} }
func (m *PimGrpMapInfosrcBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimGrpMapInfosrcBag_KEYS) ProtoMessage()    {}
func (*PimGrpMapInfosrcBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a55d304f0512423d, []int{0}
}

func (m *PimGrpMapInfosrcBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrpMapInfosrcBag_KEYS.Unmarshal(m, b)
}
func (m *PimGrpMapInfosrcBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrpMapInfosrcBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimGrpMapInfosrcBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrpMapInfosrcBag_KEYS.Merge(m, src)
}
func (m *PimGrpMapInfosrcBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimGrpMapInfosrcBag_KEYS.Size(m)
}
func (m *PimGrpMapInfosrcBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrpMapInfosrcBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrpMapInfosrcBag_KEYS proto.InternalMessageInfo

func (m *PimGrpMapInfosrcBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimGrpMapInfosrcBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
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
	return fileDescriptor_a55d304f0512423d, []int{1}
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

type PimGrpMapBag struct {
	Prefix               *PimAddrtype `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         int32        `protobuf:"zigzag32,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Client               string       `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Protocol             string       `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RpAddress            *PimAddrtype `protobuf:"bytes,5,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	GroupCount           uint32       `protobuf:"varint,6,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	IsUsed               bool         `protobuf:"varint,7,opt,name=is_used,json=isUsed,proto3" json:"is_used,omitempty"`
	MribActive           bool         `protobuf:"varint,8,opt,name=mrib_active,json=mribActive,proto3" json:"mrib_active,omitempty"`
	IsOverride           bool         `protobuf:"varint,9,opt,name=is_override,json=isOverride,proto3" json:"is_override,omitempty"`
	Priority             uint32       `protobuf:"varint,10,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimGrpMapBag) Reset()         { *m = PimGrpMapBag{} }
func (m *PimGrpMapBag) String() string { return proto.CompactTextString(m) }
func (*PimGrpMapBag) ProtoMessage()    {}
func (*PimGrpMapBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a55d304f0512423d, []int{2}
}

func (m *PimGrpMapBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrpMapBag.Unmarshal(m, b)
}
func (m *PimGrpMapBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrpMapBag.Marshal(b, m, deterministic)
}
func (m *PimGrpMapBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrpMapBag.Merge(m, src)
}
func (m *PimGrpMapBag) XXX_Size() int {
	return xxx_messageInfo_PimGrpMapBag.Size(m)
}
func (m *PimGrpMapBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrpMapBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrpMapBag proto.InternalMessageInfo

func (m *PimGrpMapBag) GetPrefix() *PimAddrtype {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PimGrpMapBag) GetPrefixLength() int32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *PimGrpMapBag) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *PimGrpMapBag) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PimGrpMapBag) GetRpAddress() *PimAddrtype {
	if m != nil {
		return m.RpAddress
	}
	return nil
}

func (m *PimGrpMapBag) GetGroupCount() uint32 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *PimGrpMapBag) GetIsUsed() bool {
	if m != nil {
		return m.IsUsed
	}
	return false
}

func (m *PimGrpMapBag) GetMribActive() bool {
	if m != nil {
		return m.MribActive
	}
	return false
}

func (m *PimGrpMapBag) GetIsOverride() bool {
	if m != nil {
		return m.IsOverride
	}
	return false
}

func (m *PimGrpMapBag) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

type PimGrpMapInfosrcBag struct {
	SourceOfInformation  *PimAddrtype  `protobuf:"bytes,50,opt,name=source_of_information,json=sourceOfInformation,proto3" json:"source_of_information,omitempty"`
	Holdtime             int32         `protobuf:"zigzag32,51,opt,name=holdtime,proto3" json:"holdtime,omitempty"`
	Expires              uint64        `protobuf:"varint,52,opt,name=expires,proto3" json:"expires,omitempty"`
	Uptime               uint64        `protobuf:"varint,53,opt,name=uptime,proto3" json:"uptime,omitempty"`
	GroupMapInformation  *PimGrpMapBag `protobuf:"bytes,54,opt,name=group_map_information,json=groupMapInformation,proto3" json:"group_map_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PimGrpMapInfosrcBag) Reset()         { *m = PimGrpMapInfosrcBag{} }
func (m *PimGrpMapInfosrcBag) String() string { return proto.CompactTextString(m) }
func (*PimGrpMapInfosrcBag) ProtoMessage()    {}
func (*PimGrpMapInfosrcBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a55d304f0512423d, []int{3}
}

func (m *PimGrpMapInfosrcBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrpMapInfosrcBag.Unmarshal(m, b)
}
func (m *PimGrpMapInfosrcBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrpMapInfosrcBag.Marshal(b, m, deterministic)
}
func (m *PimGrpMapInfosrcBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrpMapInfosrcBag.Merge(m, src)
}
func (m *PimGrpMapInfosrcBag) XXX_Size() int {
	return xxx_messageInfo_PimGrpMapInfosrcBag.Size(m)
}
func (m *PimGrpMapInfosrcBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrpMapInfosrcBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrpMapInfosrcBag proto.InternalMessageInfo

func (m *PimGrpMapInfosrcBag) GetSourceOfInformation() *PimAddrtype {
	if m != nil {
		return m.SourceOfInformation
	}
	return nil
}

func (m *PimGrpMapInfosrcBag) GetHoldtime() int32 {
	if m != nil {
		return m.Holdtime
	}
	return 0
}

func (m *PimGrpMapInfosrcBag) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *PimGrpMapInfosrcBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *PimGrpMapInfosrcBag) GetGroupMapInformation() *PimGrpMapBag {
	if m != nil {
		return m.GroupMapInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*PimGrpMapInfosrcBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.group_map_match_sources.group_map_match_source.pim_grp_map_infosrc_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.group_map_match_sources.group_map_match_source.pim_addrtype")
	proto.RegisterType((*PimGrpMapBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.group_map_match_sources.group_map_match_source.pim_grp_map_bag")
	proto.RegisterType((*PimGrpMapInfosrcBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.group_map_match_sources.group_map_match_source.pim_grp_map_infosrc_bag")
}

func init() { proto.RegisterFile("pim_grp_map_infosrc_bag.proto", fileDescriptor_a55d304f0512423d) }

var fileDescriptor_a55d304f0512423d = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0xd2, 0xb2, 0x49, 0x26, 0xa9, 0x10, 0xae, 0x4a, 0x0d, 0x02, 0x11, 0xc2, 0x25, 0xa7,
	0x3d, 0xb4, 0x25, 0xf7, 0x0a, 0x71, 0x40, 0x7c, 0x54, 0x5a, 0xc4, 0x81, 0x4b, 0x2d, 0x67, 0xd7,
	0x9b, 0x58, 0x8a, 0xd7, 0xd6, 0xd8, 0x59, 0x25, 0x77, 0x6e, 0xfc, 0x00, 0xae, 0x48, 0xfc, 0x01,
	0x7e, 0x22, 0xb2, 0xbd, 0xf9, 0xe8, 0x21, 0x47, 0xe8, 0x25, 0xca, 0xbc, 0x79, 0xf3, 0x66, 0xde,
	0x78, 0x16, 0x5e, 0x18, 0xa9, 0xd8, 0x0c, 0x0d, 0x53, 0xdc, 0x30, 0x59, 0x57, 0xda, 0x62, 0xc1,
	0xa6, 0x7c, 0x96, 0x19, 0xd4, 0x4e, 0x93, 0xdb, 0x42, 0xda, 0x42, 0x33, 0xa9, 0x2d, 0x5b, 0x21,
	0x93, 0xa6, 0xb9, 0x62, 0xbe, 0x40, 0x1b, 0x81, 0x99, 0x91, 0x2a, 0xb3, 0x8e, 0xd7, 0xe5, 0x74,
	0x9d, 0x35, 0x58, 0x59, 0xff, 0x93, 0xcd, 0x50, 0x2f, 0xa3, 0x98, 0xe2, 0xae, 0x98, 0x33, 0xab,
	0x97, 0x58, 0x08, 0x7b, 0x00, 0x1f, 0xdd, 0xc2, 0xf3, 0x03, 0x03, 0xb0, 0x0f, 0xef, 0xbe, 0x7d,
	0x21, 0x4f, 0xa1, 0xdb, 0x60, 0xc5, 0x6a, 0xae, 0x04, 0x4d, 0x86, 0xc9, 0xb8, 0x97, 0x77, 0x1a,
	0xac, 0x3e, 0x73, 0x25, 0xc8, 0x6b, 0x38, 0x89, 0xa2, 0xbc, 0x2c, 0x51, 0x58, 0x4b, 0x1f, 0x84,
	0xfc, 0x20, 0x80, 0xd7, 0x11, 0x1b, 0x29, 0x18, 0x78, 0x7d, 0x4f, 0x71, 0x6b, 0x23, 0xc8, 0x39,
	0x74, 0xf8, 0x1d, 0xb9, 0x94, 0x47, 0xb5, 0x57, 0x30, 0x08, 0xee, 0xee, 0x8a, 0xf5, 0x3d, 0xd6,
	0x6a, 0xb5, 0x94, 0xc9, 0x96, 0x72, 0xb4, 0xa5, 0x4c, 0x36, 0xed, 0x7e, 0x1e, 0xc3, 0xa3, 0x7d,
	0x3f, 0x53, 0x3e, 0x23, 0xdf, 0x13, 0x48, 0x0d, 0x8a, 0x4a, 0xae, 0x42, 0xcb, 0xfe, 0xc5, 0x22,
	0xfb, 0xb7, 0x4b, 0xcd, 0xf6, 0x1d, 0xe7, 0x6d, 0x6f, 0xbf, 0xae, 0xf8, 0x8f, 0x2d, 0x44, 0x3d,
	0x73, 0xf3, 0xe0, 0xf0, 0x71, 0x3e, 0x88, 0xe0, 0xc7, 0x80, 0x91, 0x27, 0x90, 0x16, 0x0b, 0x29,
	0x6a, 0xd7, 0x9a, 0x6b, 0x23, 0xf2, 0x0c, 0xba, 0xe1, 0x1e, 0x0a, 0xbd, 0xa0, 0xc7, 0x21, 0xb3,
	0x8d, 0xc9, 0x8f, 0x04, 0x00, 0x77, 0xaf, 0xf0, 0xf0, 0x1e, 0x3c, 0xf6, 0x70, 0xf3, 0xe0, 0xe4,
	0x25, 0xf4, 0x63, 0x45, 0xa1, 0x97, 0xb5, 0xa3, 0xe9, 0x30, 0x19, 0x9f, 0xe4, 0x10, 0xa0, 0xb7,
	0x1e, 0xf1, 0x17, 0x20, 0x2d, 0x5b, 0x5a, 0x51, 0xd2, 0xce, 0x30, 0x19, 0x77, 0xf3, 0x54, 0xda,
	0xaf, 0x56, 0x94, 0xbe, 0x52, 0xa1, 0x9c, 0x32, 0x5e, 0x38, 0xd9, 0x08, 0xda, 0x0d, 0x49, 0xf0,
	0xd0, 0x75, 0x40, 0x3c, 0x41, 0x5a, 0xa6, 0x1b, 0x81, 0x28, 0x4b, 0x41, 0x7b, 0x91, 0x20, 0xed,
	0x4d, 0x8b, 0xc4, 0x2d, 0x49, 0x8d, 0xd2, 0xad, 0x29, 0x84, 0xc6, 0xdb, 0x78, 0xf4, 0xe7, 0x08,
	0xce, 0x0f, 0x5c, 0x3a, 0xf9, 0x95, 0xc0, 0x59, 0xb4, 0xc5, 0x74, 0x15, 0x32, 0xa8, 0xb8, 0x93,
	0xba, 0xa6, 0x17, 0xf7, 0xb0, 0xcc, 0xd3, 0x08, 0xde, 0x54, 0xef, 0x77, 0x83, 0x78, 0x6b, 0x73,
	0xbd, 0x28, 0x9d, 0x54, 0x82, 0x5e, 0x86, 0xc3, 0xd9, 0xc6, 0x84, 0x42, 0x47, 0xac, 0x8c, 0x44,
	0x61, 0xe9, 0xd5, 0x30, 0x19, 0x1f, 0xe7, 0x9b, 0xd0, 0x9f, 0xd3, 0xd2, 0x84, 0x9a, 0x37, 0x21,
	0xd1, 0x46, 0xe4, 0x77, 0x02, 0x67, 0xbb, 0x51, 0xf6, 0x0d, 0x4f, 0x82, 0x61, 0xfd, 0x3f, 0x0c,
	0xef, 0x7d, 0xa3, 0xf9, 0x69, 0xe0, 0x7d, 0xe2, 0x66, 0xcf, 0xf3, 0x34, 0x0d, 0x27, 0x7e, 0xf9,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x76, 0x86, 0x57, 0x8d, 0x23, 0x05, 0x00, 0x00,
}
