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
// source: pim_grp_map_inforpf_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_group_map_rpfs_group_map_rpf

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

type PimGrpMapInforpfBag_KEYS struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Client               string   `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Protocol             string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RpAddress            string   `protobuf:"bytes,5,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	RpPriority           uint32   `protobuf:"varint,6,opt,name=rp_priority,json=rpPriority,proto3" json:"rp_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimGrpMapInforpfBag_KEYS) Reset()         { *m = PimGrpMapInforpfBag_KEYS{} }
func (m *PimGrpMapInforpfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimGrpMapInforpfBag_KEYS) ProtoMessage()    {}
func (*PimGrpMapInforpfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0199b284b053a4, []int{0}
}

func (m *PimGrpMapInforpfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrpMapInforpfBag_KEYS.Unmarshal(m, b)
}
func (m *PimGrpMapInforpfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrpMapInforpfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimGrpMapInforpfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrpMapInforpfBag_KEYS.Merge(m, src)
}
func (m *PimGrpMapInforpfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimGrpMapInforpfBag_KEYS.Size(m)
}
func (m *PimGrpMapInforpfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrpMapInforpfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrpMapInforpfBag_KEYS proto.InternalMessageInfo

func (m *PimGrpMapInforpfBag_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *PimGrpMapInforpfBag_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *PimGrpMapInforpfBag_KEYS) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *PimGrpMapInforpfBag_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PimGrpMapInforpfBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

func (m *PimGrpMapInforpfBag_KEYS) GetRpPriority() uint32 {
	if m != nil {
		return m.RpPriority
	}
	return 0
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
	return fileDescriptor_fc0199b284b053a4, []int{1}
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
	return fileDescriptor_fc0199b284b053a4, []int{2}
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

type PimGrpMapInforpfBag struct {
	AreWeRp              bool          `protobuf:"varint,50,opt,name=are_we_rp,json=areWeRp,proto3" json:"are_we_rp,omitempty"`
	RpfInterfaceName     string        `protobuf:"bytes,51,opt,name=rpf_interface_name,json=rpfInterfaceName,proto3" json:"rpf_interface_name,omitempty"`
	RpfNeighbor          *PimAddrtype  `protobuf:"bytes,52,opt,name=rpf_neighbor,json=rpfNeighbor,proto3" json:"rpf_neighbor,omitempty"`
	RpfVrfName           string        `protobuf:"bytes,53,opt,name=rpf_vrf_name,json=rpfVrfName,proto3" json:"rpf_vrf_name,omitempty"`
	GroupMapInformation  *PimGrpMapBag `protobuf:"bytes,54,opt,name=group_map_information,json=groupMapInformation,proto3" json:"group_map_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PimGrpMapInforpfBag) Reset()         { *m = PimGrpMapInforpfBag{} }
func (m *PimGrpMapInforpfBag) String() string { return proto.CompactTextString(m) }
func (*PimGrpMapInforpfBag) ProtoMessage()    {}
func (*PimGrpMapInforpfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc0199b284b053a4, []int{3}
}

func (m *PimGrpMapInforpfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGrpMapInforpfBag.Unmarshal(m, b)
}
func (m *PimGrpMapInforpfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGrpMapInforpfBag.Marshal(b, m, deterministic)
}
func (m *PimGrpMapInforpfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGrpMapInforpfBag.Merge(m, src)
}
func (m *PimGrpMapInforpfBag) XXX_Size() int {
	return xxx_messageInfo_PimGrpMapInforpfBag.Size(m)
}
func (m *PimGrpMapInforpfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGrpMapInforpfBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGrpMapInforpfBag proto.InternalMessageInfo

func (m *PimGrpMapInforpfBag) GetAreWeRp() bool {
	if m != nil {
		return m.AreWeRp
	}
	return false
}

func (m *PimGrpMapInforpfBag) GetRpfInterfaceName() string {
	if m != nil {
		return m.RpfInterfaceName
	}
	return ""
}

func (m *PimGrpMapInforpfBag) GetRpfNeighbor() *PimAddrtype {
	if m != nil {
		return m.RpfNeighbor
	}
	return nil
}

func (m *PimGrpMapInforpfBag) GetRpfVrfName() string {
	if m != nil {
		return m.RpfVrfName
	}
	return ""
}

func (m *PimGrpMapInforpfBag) GetGroupMapInformation() *PimGrpMapBag {
	if m != nil {
		return m.GroupMapInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*PimGrpMapInforpfBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.group_map_rpfs.group_map_rpf.pim_grp_map_inforpf_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.group_map_rpfs.group_map_rpf.pim_addrtype")
	proto.RegisterType((*PimGrpMapBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.group_map_rpfs.group_map_rpf.pim_grp_map_bag")
	proto.RegisterType((*PimGrpMapInforpfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.group_map_rpfs.group_map_rpf.pim_grp_map_inforpf_bag")
}

func init() { proto.RegisterFile("pim_grp_map_inforpf_bag.proto", fileDescriptor_fc0199b284b053a4) }

var fileDescriptor_fc0199b284b053a4 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x36, 0x90, 0x1f, 0x27, 0x15, 0x60, 0x04, 0xb5, 0x2a, 0x2a, 0x42, 0xb8, 0xe4, 0x80,
	0xf6, 0xd0, 0x96, 0xdc, 0x2b, 0xc4, 0xa1, 0x02, 0x0a, 0x5a, 0x04, 0x88, 0x93, 0x71, 0x76, 0x67,
	0xb7, 0x96, 0xb2, 0xeb, 0xd1, 0xac, 0x13, 0xd2, 0x13, 0x07, 0x1e, 0x82, 0x27, 0xe0, 0x79, 0x78,
	0x22, 0x24, 0x64, 0x7b, 0x93, 0x26, 0x42, 0xe5, 0x54, 0x7a, 0xf3, 0x7c, 0xf3, 0x79, 0xfe, 0xbe,
	0x19, 0x76, 0x80, 0xba, 0x94, 0x05, 0xa1, 0x2c, 0x15, 0x4a, 0x5d, 0xe5, 0x86, 0x30, 0x97, 0x53,
	0x55, 0xc4, 0x48, 0xc6, 0x1a, 0x9e, 0xa4, 0xba, 0x4e, 0x8d, 0xd4, 0xa6, 0x96, 0x4b, 0x92, 0x1a,
	0x17, 0xc7, 0xd2, 0x7d, 0x30, 0x08, 0x14, 0xa3, 0x2e, 0x63, 0x95, 0x5a, 0xbd, 0x80, 0x38, 0x83,
	0x5c, 0xcd, 0x67, 0x56, 0xa6, 0xa6, 0xb2, 0xb0, 0xb4, 0x71, 0x41, 0x66, 0x1e, 0x42, 0x12, 0xe6,
	0xf5, 0xb6, 0x39, 0xfa, 0x15, 0xb1, 0x47, 0x57, 0x64, 0x95, 0xaf, 0x5e, 0x7e, 0x7e, 0xcf, 0x1f,
	0xb2, 0x36, 0x12, 0xe4, 0x7a, 0x29, 0xa2, 0x61, 0x34, 0xee, 0x25, 0x8d, 0xc5, 0x9f, 0xb2, 0xdd,
	0xf0, 0x92, 0x33, 0xa8, 0x0a, 0x7b, 0x2e, 0x76, 0x86, 0xd1, 0x78, 0x37, 0x19, 0x04, 0xf0, 0xb5,
	0xc7, 0xdc, 0xe7, 0x74, 0xa6, 0xa1, 0xb2, 0xa2, 0x15, 0x3e, 0x07, 0x8b, 0xef, 0xb3, 0xae, 0x6f,
	0x29, 0x35, 0x33, 0x71, 0xcb, 0x7b, 0xd6, 0x36, 0x3f, 0x60, 0x8c, 0x50, 0xaa, 0x2c, 0x23, 0xa8,
	0x6b, 0x71, 0xdb, 0x7b, 0x7b, 0x84, 0x27, 0x01, 0xe0, 0x8f, 0x59, 0x9f, 0x50, 0x22, 0x69, 0x43,
	0xda, 0x5e, 0x88, 0xb6, 0xcf, 0xca, 0x08, 0xdf, 0x35, 0xc8, 0xa8, 0x64, 0x03, 0xd7, 0x90, 0x0b,
	0x60, 0x2f, 0x10, 0xf8, 0x1e, 0xeb, 0xa8, 0x5c, 0x56, 0xaa, 0x84, 0x55, 0x07, 0x2a, 0x3f, 0x53,
	0x25, 0xf0, 0x27, 0x6c, 0xe0, 0x67, 0xb8, 0x4a, 0xb5, 0xe3, 0xbd, 0x7d, 0x87, 0xad, 0x92, 0x05,
	0xca, 0x64, 0x4d, 0x69, 0xad, 0x29, 0x93, 0x86, 0x32, 0xfa, 0xdd, 0x62, 0x77, 0x36, 0x07, 0x38,
	0x55, 0x05, 0x5f, 0x6e, 0xcd, 0xac, 0x7f, 0xf8, 0x25, 0xbe, 0x7e, 0xe5, 0xe2, 0xcd, 0x26, 0xff,
	0xad, 0xca, 0xbd, 0x6b, 0x50, 0xe5, 0xdb, 0x5f, 0xaa, 0xdc, 0x44, 0x5b, 0xdb, 0xba, 0x07, 0x62,
	0x6a, 0xe6, 0x95, 0x5d, 0xe9, 0xee, 0xa1, 0x17, 0x0e, 0x71, 0x3a, 0xeb, 0x5a, 0xce, 0x6b, 0xc8,
	0x44, 0x67, 0x18, 0x8d, 0xbb, 0x49, 0x5b, 0xd7, 0x1f, 0x6a, 0xc8, 0xdc, 0xcf, 0x92, 0xf4, 0x54,
	0x86, 0x6a, 0x44, 0xd7, 0x3b, 0x99, 0x83, 0x4e, 0x3c, 0xe2, 0x08, 0xba, 0x96, 0x66, 0x01, 0x44,
	0x3a, 0x03, 0xd1, 0x0b, 0x04, 0x5d, 0xbf, 0x6d, 0x90, 0x30, 0x98, 0x66, 0xe1, 0x98, 0x4f, 0xbc,
	0xb6, 0x47, 0x3f, 0x5b, 0x6c, 0xef, 0x8a, 0x03, 0xe2, 0xfb, 0xac, 0xa7, 0x08, 0xe4, 0x57, 0x90,
	0x84, 0xe2, 0xd0, 0x87, 0xed, 0x28, 0x82, 0x4f, 0x90, 0x20, 0x7f, 0xc6, 0xb8, 0xa3, 0xe9, 0xca,
	0x02, 0xe5, 0x2a, 0x85, 0xb0, 0xa1, 0x47, 0x7e, 0xec, 0x77, 0x09, 0xf3, 0xd3, 0x95, 0xc3, 0xef,
	0xea, 0xf7, 0x88, 0x0d, 0x1c, 0xbd, 0x02, 0x5d, 0x9c, 0x4f, 0x0d, 0x89, 0xe3, 0x1b, 0x52, 0xa0,
	0x4f, 0x98, 0x9f, 0x35, 0x49, 0xf9, 0x30, 0x14, 0xb1, 0xa0, 0xe6, 0x9e, 0x9e, 0xfb, 0x6a, 0x19,
	0x61, 0xfe, 0x91, 0xc2, 0x4d, 0xfd, 0x88, 0xd8, 0x83, 0xcb, 0x68, 0x7e, 0x16, 0xa5, 0xb2, 0xda,
	0x54, 0x62, 0xe2, 0x0b, 0x4e, 0xff, 0x57, 0xc1, 0x1b, 0xe7, 0x97, 0xdc, 0xf7, 0xee, 0x37, 0x0a,
	0x4f, 0x2f, 0xf3, 0x4f, 0xdb, 0x7e, 0x95, 0x8f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x0c,
	0xca, 0x1b, 0x64, 0x05, 0x00, 0x00,
}
