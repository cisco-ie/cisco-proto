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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_group_map_match_rpfs_group_map_match_rpf

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
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
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

func (m *PimGrpMapInforpfBag_KEYS) GetGroupAddress() string {
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
	proto.RegisterType((*PimGrpMapInforpfBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.group_map_match_rpfs.group_map_match_rpf.pim_grp_map_inforpf_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.group_map_match_rpfs.group_map_match_rpf.pim_addrtype")
	proto.RegisterType((*PimGrpMapBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.group_map_match_rpfs.group_map_match_rpf.pim_grp_map_bag")
	proto.RegisterType((*PimGrpMapInforpfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.group_map_match_rpfs.group_map_match_rpf.pim_grp_map_inforpf_bag")
}

func init() { proto.RegisterFile("pim_grp_map_inforpf_bag.proto", fileDescriptor_fc0199b284b053a4) }

var fileDescriptor_fc0199b284b053a4 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xd2, 0xb2, 0x49, 0x26, 0xa9, 0x00, 0x23, 0xe8, 0xaa, 0x02, 0x11, 0xc2, 0x25, 0x07,
	0xb4, 0x87, 0xb6, 0xe4, 0x5e, 0x55, 0x1c, 0x2a, 0xa0, 0x48, 0x8b, 0x00, 0x71, 0xb2, 0x9c, 0xdd,
	0xd9, 0x8d, 0xa5, 0xec, 0xda, 0x8c, 0x9d, 0x90, 0x3e, 0x02, 0x37, 0x0e, 0xbc, 0x00, 0x37, 0xde,
	0x84, 0xd7, 0x42, 0xb6, 0x37, 0x3f, 0x87, 0xf6, 0xda, 0xde, 0x3c, 0xdf, 0x7c, 0x1e, 0x7f, 0xdf,
	0xcc, 0x18, 0x9e, 0x6b, 0x59, 0xf3, 0x8a, 0x34, 0xaf, 0x85, 0xe6, 0xb2, 0x29, 0x15, 0xe9, 0x92,
	0x4f, 0x45, 0x95, 0x6a, 0x52, 0x56, 0xb1, 0x3c, 0x97, 0x26, 0x57, 0x5c, 0x2a, 0xc3, 0x57, 0xc4,
	0xa5, 0x5e, 0x9e, 0x72, 0x77, 0x41, 0x69, 0xa4, 0x54, 0xea, 0xe5, 0xc4, 0x45, 0xa9, 0xc8, 0xad,
	0x5c, 0x62, 0x5a, 0x60, 0x29, 0x16, 0x73, 0xcb, 0x73, 0xd5, 0x58, 0x5c, 0xd9, 0xb4, 0x22, 0xb5,
	0x08, 0x75, 0x6b, 0x61, 0xf3, 0x19, 0x27, 0x5d, 0x9a, 0xeb, 0xc0, 0xd1, 0x39, 0x3c, 0xbb, 0x41,
	0x05, 0x7f, 0xf7, 0xf6, 0xdb, 0x27, 0xf6, 0x0a, 0x0e, 0xc2, 0x35, 0x51, 0x14, 0x84, 0xc6, 0x24,
	0xd1, 0x30, 0x1a, 0xf7, 0xb2, 0x81, 0x07, 0xcf, 0x02, 0x36, 0xaa, 0x61, 0xe0, 0x8a, 0x38, 0x8a,
	0xbd, 0xd2, 0xc8, 0x0e, 0xa1, 0x23, 0x4a, 0xde, 0x88, 0x1a, 0x5b, 0x7a, 0x2c, 0xca, 0x4b, 0x51,
	0x23, 0x7b, 0x09, 0x03, 0xef, 0x63, 0x5d, 0xec, 0x9e, 0xcf, 0xf6, 0x1d, 0xd6, 0xd6, 0x6a, 0x29,
	0x93, 0x0d, 0x65, 0x6f, 0x43, 0x99, 0xac, 0x9f, 0xfb, 0xb3, 0x0f, 0x0f, 0x76, 0x45, 0x4f, 0x45,
	0xc5, 0x7e, 0x46, 0x10, 0x6b, 0xc2, 0x52, 0xae, 0xfc, 0x93, 0xfd, 0xe3, 0xef, 0xe9, 0x2d, 0xb4,
	0x2f, 0xdd, 0xb5, 0x9d, 0xb5, 0x02, 0x5c, 0xcf, 0xc2, 0x89, 0xcf, 0xb1, 0xa9, 0xec, 0xcc, 0xdb,
	0x7c, 0x94, 0x0d, 0x02, 0xf8, 0xde, 0x63, 0xec, 0x29, 0xc4, 0xf9, 0x5c, 0x62, 0x63, 0x5b, 0x87,
	0x6d, 0xc4, 0x8e, 0xa0, 0xeb, 0xc7, 0x9f, 0xab, 0x79, 0xb2, 0xef, 0x33, 0x9b, 0x98, 0xfd, 0x8a,
	0x00, 0x68, 0x3b, 0x8a, 0xfb, 0x77, 0x65, 0xb4, 0x47, 0xeb, 0xd1, 0xb3, 0x17, 0xd0, 0x0f, 0xf4,
	0x5c, 0x2d, 0x1a, 0x9b, 0xc4, 0xc3, 0x68, 0x7c, 0x90, 0x81, 0x87, 0xce, 0x1d, 0xe2, 0x76, 0x41,
	0x1a, 0xbe, 0x30, 0x58, 0x24, 0x9d, 0x61, 0x34, 0xee, 0x66, 0xb1, 0x34, 0x9f, 0x0d, 0x16, 0xee,
	0x66, 0x4d, 0x72, 0xca, 0x83, 0xb2, 0xa4, 0xeb, 0x93, 0xe0, 0xa0, 0x33, 0x8f, 0x38, 0x82, 0x34,
	0x5c, 0x2d, 0x91, 0x48, 0x16, 0x98, 0xf4, 0x02, 0x41, 0x9a, 0x8f, 0x2d, 0x12, 0x5a, 0x25, 0x15,
	0x49, 0x7b, 0x95, 0x80, 0x7f, 0x78, 0x13, 0x8f, 0xfe, 0xed, 0xc1, 0xe1, 0x0d, 0x8b, 0xcd, 0x8e,
	0xa0, 0x27, 0x08, 0xf9, 0x0f, 0xe4, 0xa4, 0x93, 0x63, 0x5f, 0xb6, 0x23, 0x08, 0xbf, 0x62, 0xa6,
	0xd9, 0x6b, 0x60, 0x8e, 0x26, 0x1b, 0x8b, 0x54, 0x8a, 0x1c, 0xc3, 0x16, 0x9f, 0xf8, 0x41, 0x3c,
	0x24, 0x5d, 0x5e, 0xac, 0x13, 0x7e, 0x9f, 0x7f, 0x47, 0x30, 0x70, 0xf4, 0x06, 0x65, 0x35, 0x9b,
	0x2a, 0x4a, 0x4e, 0xef, 0x6a, 0x24, 0x7d, 0xd2, 0xe5, 0x65, 0xab, 0x82, 0x0d, 0x83, 0xaa, 0x25,
	0xb5, 0x9f, 0xf0, 0x8d, 0x97, 0x0f, 0xa4, 0xcb, 0x2f, 0x14, 0x3e, 0xe2, 0xdf, 0x08, 0x9e, 0x6c,
	0x6b, 0xfa, 0xe6, 0xd4, 0xc2, 0x4a, 0xd5, 0x24, 0x13, 0xef, 0xc0, 0xde, 0x9a, 0x83, 0x9d, 0x4f,
	0x9c, 0x3d, 0xf6, 0xa4, 0x0f, 0x42, 0x5f, 0x6c, 0x05, 0x4d, 0x63, 0xbf, 0xfe, 0x27, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x26, 0xc8, 0xfb, 0x2e, 0x05, 0x00, 0x00,
}
