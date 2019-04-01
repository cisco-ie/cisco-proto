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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_group_map_match_sources_group_map_match_source

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
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
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
	proto.RegisterType((*PimGrpMapInfosrcBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.group_map_match_sources.group_map_match_source.pim_grp_map_infosrc_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.group_map_match_sources.group_map_match_source.pim_addrtype")
	proto.RegisterType((*PimGrpMapBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.group_map_match_sources.group_map_match_source.pim_grp_map_bag")
	proto.RegisterType((*PimGrpMapInfosrcBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.group_map_match_sources.group_map_match_source.pim_grp_map_infosrc_bag")
}

func init() { proto.RegisterFile("pim_grp_map_infosrc_bag.proto", fileDescriptor_a55d304f0512423d) }

var fileDescriptor_a55d304f0512423d = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0x69, 0x71, 0x92, 0x49, 0x2a, 0xc4, 0x56, 0xa5, 0x2b, 0x04, 0x22, 0x84, 0x4b, 0x4e,
	0x3e, 0xb4, 0x25, 0xf7, 0xaa, 0xe2, 0x80, 0xf8, 0xa9, 0x64, 0xc4, 0x81, 0xd3, 0x6a, 0x63, 0x8f,
	0x93, 0x91, 0x62, 0xef, 0x6a, 0x77, 0x1d, 0x25, 0xcf, 0xc0, 0x8d, 0x77, 0x40, 0x48, 0x3c, 0x04,
	0xcf, 0x86, 0x76, 0xed, 0x24, 0xbe, 0xe4, 0xda, 0xde, 0x3c, 0xdf, 0x7c, 0xf3, 0xf3, 0xcd, 0x7e,
	0x86, 0xd7, 0x9a, 0x4a, 0xb1, 0x30, 0x5a, 0x94, 0x52, 0x0b, 0xaa, 0x0a, 0x65, 0x4d, 0x26, 0xe6,
	0x72, 0x91, 0x68, 0xa3, 0x9c, 0x62, 0x98, 0x91, 0xcd, 0x94, 0x20, 0x65, 0xc5, 0xc6, 0x08, 0xd2,
	0xeb, 0x1b, 0xe1, 0x0b, 0x94, 0x46, 0x93, 0x68, 0x2a, 0x13, 0xeb, 0x64, 0x95, 0xcf, 0xb7, 0x49,
	0x8e, 0x85, 0xac, 0x57, 0x4e, 0x64, 0xaa, 0x72, 0xb8, 0x71, 0xc9, 0xc2, 0xa8, 0xba, 0xe9, 0x59,
	0x4a, 0x97, 0x2d, 0x85, 0x55, 0xb5, 0xc9, 0xd0, 0x1e, 0xc1, 0x27, 0x77, 0xf0, 0xea, 0xc8, 0x1e,
	0xe2, 0xd3, 0x87, 0x1f, 0xdf, 0xd8, 0x3b, 0x38, 0x6b, 0x2a, 0x65, 0x9e, 0x1b, 0xb4, 0x96, 0x47,
	0xe3, 0x68, 0x3a, 0x48, 0x47, 0x01, 0xbc, 0x6d, 0xb0, 0x49, 0x09, 0x23, 0xdf, 0xc4, 0x53, 0xdc,
	0x56, 0x23, 0xbb, 0x84, 0x9e, 0x2c, 0x44, 0x25, 0x4b, 0x6c, 0xe9, 0xb1, 0x2c, 0xbe, 0xca, 0x12,
	0xd9, 0x5b, 0x18, 0x05, 0x25, 0xbb, 0x66, 0x4f, 0x42, 0x76, 0xe8, 0xb1, 0xb6, 0x57, 0x4b, 0x99,
	0xed, 0x29, 0x27, 0x7b, 0xca, 0x6c, 0x37, 0xee, 0xf7, 0x29, 0x3c, 0xeb, 0x2e, 0x3d, 0x97, 0x0b,
	0xf6, 0x33, 0x82, 0x58, 0x1b, 0x2c, 0x68, 0x13, 0x46, 0x0e, 0xaf, 0x6c, 0xf2, 0x20, 0x07, 0x4c,
	0xba, 0xc2, 0xd3, 0x76, 0x05, 0x7f, 0xb5, 0xe6, 0x4b, 0xac, 0xb0, 0x5a, 0xb8, 0x65, 0x10, 0xfa,
	0x3c, 0x1d, 0x35, 0xe0, 0xe7, 0x80, 0xb1, 0x17, 0x10, 0x67, 0x2b, 0xc2, 0xca, 0xb5, 0x1a, 0xdb,
	0x88, 0xbd, 0x84, 0x7e, 0xb0, 0x40, 0xa6, 0x56, 0xfc, 0x34, 0x64, 0xf6, 0x31, 0xfb, 0x15, 0x01,
	0x98, 0xc3, 0x63, 0x3c, 0x7d, 0x3c, 0xa9, 0x03, 0xb3, 0x7b, 0x7e, 0xf6, 0x06, 0x86, 0x4d, 0x45,
	0xa6, 0xea, 0xca, 0xf1, 0x78, 0x1c, 0x4d, 0xcf, 0x52, 0x08, 0xd0, 0x9d, 0x47, 0xbc, 0x1f, 0xc8,
	0x8a, 0xda, 0x62, 0xce, 0x7b, 0xe3, 0x68, 0xda, 0x4f, 0x63, 0xb2, 0xdf, 0x2d, 0xe6, 0xbe, 0xb2,
	0x34, 0x34, 0x17, 0x32, 0x73, 0xb4, 0x46, 0xde, 0x0f, 0x49, 0xf0, 0xd0, 0x6d, 0x40, 0x3c, 0x81,
	0xac, 0x50, 0x6b, 0x34, 0x86, 0x72, 0xe4, 0x83, 0x86, 0x40, 0xf6, 0xbe, 0x45, 0x9a, 0x63, 0x91,
	0x32, 0xe4, 0xb6, 0x1c, 0xc2, 0xe0, 0x7d, 0x3c, 0xf9, 0x77, 0x02, 0x97, 0x47, 0xcc, 0xcd, 0xfe,
	0x44, 0x70, 0xd1, 0xc8, 0x12, 0xaa, 0x08, 0x19, 0x53, 0x4a, 0x47, 0xaa, 0xe2, 0x57, 0x8f, 0x77,
	0xd3, 0xf3, 0x06, 0xbc, 0x2f, 0x3e, 0x1e, 0xf6, 0xf1, 0x0a, 0x97, 0x6a, 0x95, 0x3b, 0x2a, 0x91,
	0x5f, 0x07, 0x1b, 0xed, 0x63, 0xc6, 0xa1, 0x87, 0x1b, 0x4d, 0x06, 0x2d, 0xbf, 0x19, 0x47, 0xd3,
	0xd3, 0x74, 0x17, 0x7a, 0x73, 0xd5, 0x3a, 0xd4, 0xbc, 0x0f, 0x89, 0x36, 0x62, 0x7f, 0x23, 0xb8,
	0x38, 0xac, 0xd2, 0xd5, 0x3d, 0x0b, 0xba, 0xd7, 0x0f, 0xa8, 0xbb, 0xf3, 0xff, 0xa6, 0xe7, 0x81,
	0xf7, 0x45, 0xea, 0x8e, 0xf4, 0x79, 0x1c, 0x7c, 0x7f, 0xfd, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0xe3, 0x02, 0x12, 0x2b, 0x05, 0x00, 0x00,
}