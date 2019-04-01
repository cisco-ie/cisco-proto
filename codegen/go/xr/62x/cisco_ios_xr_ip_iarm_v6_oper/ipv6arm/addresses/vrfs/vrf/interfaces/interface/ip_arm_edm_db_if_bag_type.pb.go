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
// source: ip_arm_edm_db_if_bag_type.proto

package cisco_ios_xr_ip_iarm_v6_oper_ipv6arm_addresses_vrfs_vrf_interfaces_interface

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

type IpArmEdmDbIfBagType_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArmEdmDbIfBagType_KEYS) Reset()         { *m = IpArmEdmDbIfBagType_KEYS{} }
func (m *IpArmEdmDbIfBagType_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpArmEdmDbIfBagType_KEYS) ProtoMessage()    {}
func (*IpArmEdmDbIfBagType_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cabb67f0dd4132, []int{0}
}

func (m *IpArmEdmDbIfBagType_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArmEdmDbIfBagType_KEYS.Unmarshal(m, b)
}
func (m *IpArmEdmDbIfBagType_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArmEdmDbIfBagType_KEYS.Marshal(b, m, deterministic)
}
func (m *IpArmEdmDbIfBagType_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArmEdmDbIfBagType_KEYS.Merge(m, src)
}
func (m *IpArmEdmDbIfBagType_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpArmEdmDbIfBagType_KEYS.Size(m)
}
func (m *IpArmEdmDbIfBagType_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArmEdmDbIfBagType_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpArmEdmDbIfBagType_KEYS proto.InternalMessageInfo

func (m *IpArmEdmDbIfBagType_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IpArmEdmDbIfBagType_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

type ArmAddrtype struct {
	Afi                  int32    `protobuf:"zigzag32,1,opt,name=afi,proto3" json:"afi,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArmAddrtype) Reset()         { *m = ArmAddrtype{} }
func (m *ArmAddrtype) String() string { return proto.CompactTextString(m) }
func (*ArmAddrtype) ProtoMessage()    {}
func (*ArmAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cabb67f0dd4132, []int{1}
}

func (m *ArmAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArmAddrtype.Unmarshal(m, b)
}
func (m *ArmAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArmAddrtype.Marshal(b, m, deterministic)
}
func (m *ArmAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArmAddrtype.Merge(m, src)
}
func (m *ArmAddrtype) XXX_Size() int {
	return xxx_messageInfo_ArmAddrtype.Size(m)
}
func (m *ArmAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_ArmAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_ArmAddrtype proto.InternalMessageInfo

func (m *ArmAddrtype) GetAfi() int32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

func (m *ArmAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *ArmAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type IpArmEdmAddr struct {
	Address              *ArmAddrtype `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32       `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	RouteTag             uint32       `protobuf:"varint,3,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	IsPrimary            bool         `protobuf:"varint,4,opt,name=is_primary,json=isPrimary,proto3" json:"is_primary,omitempty"`
	IsTentative          bool         `protobuf:"varint,5,opt,name=is_tentative,json=isTentative,proto3" json:"is_tentative,omitempty"`
	IsPrefixSid          bool         `protobuf:"varint,6,opt,name=is_prefix_sid,json=isPrefixSid,proto3" json:"is_prefix_sid,omitempty"`
	Producer             string       `protobuf:"bytes,7,opt,name=producer,proto3" json:"producer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IpArmEdmAddr) Reset()         { *m = IpArmEdmAddr{} }
func (m *IpArmEdmAddr) String() string { return proto.CompactTextString(m) }
func (*IpArmEdmAddr) ProtoMessage()    {}
func (*IpArmEdmAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cabb67f0dd4132, []int{2}
}

func (m *IpArmEdmAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArmEdmAddr.Unmarshal(m, b)
}
func (m *IpArmEdmAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArmEdmAddr.Marshal(b, m, deterministic)
}
func (m *IpArmEdmAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArmEdmAddr.Merge(m, src)
}
func (m *IpArmEdmAddr) XXX_Size() int {
	return xxx_messageInfo_IpArmEdmAddr.Size(m)
}
func (m *IpArmEdmAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArmEdmAddr.DiscardUnknown(m)
}

var xxx_messageInfo_IpArmEdmAddr proto.InternalMessageInfo

func (m *IpArmEdmAddr) GetAddress() *ArmAddrtype {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IpArmEdmAddr) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *IpArmEdmAddr) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

func (m *IpArmEdmAddr) GetIsPrimary() bool {
	if m != nil {
		return m.IsPrimary
	}
	return false
}

func (m *IpArmEdmAddr) GetIsTentative() bool {
	if m != nil {
		return m.IsTentative
	}
	return false
}

func (m *IpArmEdmAddr) GetIsPrefixSid() bool {
	if m != nil {
		return m.IsPrefixSid
	}
	return false
}

func (m *IpArmEdmAddr) GetProducer() string {
	if m != nil {
		return m.Producer
	}
	return ""
}

type ReferencedInterfaceItem struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReferencedInterfaceItem) Reset()         { *m = ReferencedInterfaceItem{} }
func (m *ReferencedInterfaceItem) String() string { return proto.CompactTextString(m) }
func (*ReferencedInterfaceItem) ProtoMessage()    {}
func (*ReferencedInterfaceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cabb67f0dd4132, []int{3}
}

func (m *ReferencedInterfaceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReferencedInterfaceItem.Unmarshal(m, b)
}
func (m *ReferencedInterfaceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReferencedInterfaceItem.Marshal(b, m, deterministic)
}
func (m *ReferencedInterfaceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferencedInterfaceItem.Merge(m, src)
}
func (m *ReferencedInterfaceItem) XXX_Size() int {
	return xxx_messageInfo_ReferencedInterfaceItem.Size(m)
}
func (m *ReferencedInterfaceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferencedInterfaceItem.DiscardUnknown(m)
}

var xxx_messageInfo_ReferencedInterfaceItem proto.InternalMessageInfo

func (m *ReferencedInterfaceItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IpArmEdmDbIfBagType struct {
	Address              []*IpArmEdmAddr          `protobuf:"bytes,50,rep,name=address,proto3" json:"address,omitempty"`
	ReferencedInterface  *ReferencedInterfaceItem `protobuf:"bytes,51,opt,name=referenced_interface,json=referencedInterface,proto3" json:"referenced_interface,omitempty"`
	VrfName              string                   `protobuf:"bytes,52,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IpArmEdmDbIfBagType) Reset()         { *m = IpArmEdmDbIfBagType{} }
func (m *IpArmEdmDbIfBagType) String() string { return proto.CompactTextString(m) }
func (*IpArmEdmDbIfBagType) ProtoMessage()    {}
func (*IpArmEdmDbIfBagType) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9cabb67f0dd4132, []int{4}
}

func (m *IpArmEdmDbIfBagType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArmEdmDbIfBagType.Unmarshal(m, b)
}
func (m *IpArmEdmDbIfBagType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArmEdmDbIfBagType.Marshal(b, m, deterministic)
}
func (m *IpArmEdmDbIfBagType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArmEdmDbIfBagType.Merge(m, src)
}
func (m *IpArmEdmDbIfBagType) XXX_Size() int {
	return xxx_messageInfo_IpArmEdmDbIfBagType.Size(m)
}
func (m *IpArmEdmDbIfBagType) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArmEdmDbIfBagType.DiscardUnknown(m)
}

var xxx_messageInfo_IpArmEdmDbIfBagType proto.InternalMessageInfo

func (m *IpArmEdmDbIfBagType) GetAddress() []*IpArmEdmAddr {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IpArmEdmDbIfBagType) GetReferencedInterface() *ReferencedInterfaceItem {
	if m != nil {
		return m.ReferencedInterface
	}
	return nil
}

func (m *IpArmEdmDbIfBagType) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func init() {
	proto.RegisterType((*IpArmEdmDbIfBagType_KEYS)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.addresses.vrfs.vrf.interfaces.interface.ip_arm_edm_db_if_bag_type_KEYS")
	proto.RegisterType((*ArmAddrtype)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.addresses.vrfs.vrf.interfaces.interface.arm_addrtype")
	proto.RegisterType((*IpArmEdmAddr)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.addresses.vrfs.vrf.interfaces.interface.ip_arm_edm_addr")
	proto.RegisterType((*ReferencedInterfaceItem)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.addresses.vrfs.vrf.interfaces.interface.referenced_interface_item")
	proto.RegisterType((*IpArmEdmDbIfBagType)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.addresses.vrfs.vrf.interfaces.interface.ip_arm_edm_db_if_bag_type")
}

func init() { proto.RegisterFile("ip_arm_edm_db_if_bag_type.proto", fileDescriptor_b9cabb67f0dd4132) }

var fileDescriptor_b9cabb67f0dd4132 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x55, 0x12, 0xda, 0x24, 0x93, 0x44, 0xc0, 0xd2, 0x83, 0xc3, 0x67, 0x30, 0x97, 0x9c, 0x2c,
	0x91, 0x56, 0xb9, 0x73, 0xe0, 0x80, 0xa8, 0x50, 0xe5, 0xf6, 0x52, 0x24, 0x34, 0xda, 0xc4, 0xb3,
	0x66, 0xa4, 0xda, 0x5e, 0xed, 0x6e, 0x4c, 0xfb, 0x57, 0xf8, 0x4b, 0xf0, 0xa3, 0x90, 0x37, 0xb1,
	0xdd, 0xa2, 0xe6, 0x96, 0x8b, 0xb5, 0xfb, 0xf6, 0xf9, 0xcd, 0xbc, 0xf9, 0x80, 0x77, 0xac, 0x51,
	0x9a, 0x0c, 0x29, 0xc9, 0x30, 0x59, 0x21, 0x2b, 0x5c, 0xc9, 0x14, 0xdd, 0x9d, 0xa6, 0x48, 0x9b,
	0xc2, 0x15, 0xe2, 0x7c, 0xcd, 0x76, 0x5d, 0x20, 0x17, 0x16, 0x6f, 0x0d, 0xb2, 0x46, 0xae, 0xe8,
	0xe5, 0x12, 0x0b, 0x4d, 0x26, 0x62, 0x5d, 0x2e, 0xa5, 0xc9, 0x22, 0x99, 0x24, 0x86, 0xac, 0x25,
	0x1b, 0x95, 0x46, 0xf9, 0x4f, 0xc4, 0xb9, 0x23, 0xa3, 0xe4, 0x9a, 0x6c, 0x7b, 0x0c, 0xaf, 0xe1,
	0xed, 0xde, 0x80, 0xf8, 0xf5, 0xf3, 0xf5, 0xa5, 0x98, 0xc2, 0xa0, 0x34, 0x0a, 0x73, 0x99, 0x51,
	0xd0, 0x99, 0x75, 0xe6, 0xc3, 0xb8, 0x5f, 0x1a, 0xf5, 0x4d, 0x66, 0x24, 0x5e, 0xc3, 0xb0, 0x51,
	0x0a, 0xba, 0xfe, 0xad, 0x05, 0x42, 0x05, 0xe3, 0x4a, 0xb7, 0xca, 0xa4, 0x52, 0x13, 0xcf, 0xa0,
	0x27, 0x15, 0x7b, 0x8d, 0xe7, 0x71, 0x75, 0x14, 0xef, 0x61, 0xcc, 0xba, 0x3c, 0xc3, 0x5d, 0xb2,
	0x3b, 0x89, 0x51, 0x85, 0x7d, 0xda, 0x42, 0x3b, 0xca, 0xb2, 0xa1, 0xf4, 0x1a, 0xca, 0x72, 0x47,
	0x09, 0xff, 0x74, 0xe1, 0xe9, 0x3d, 0x0f, 0x15, 0x53, 0x38, 0xe8, 0xd7, 0x7f, 0x54, 0xf1, 0x46,
	0x8b, 0xef, 0xd1, 0x21, 0xcb, 0x16, 0xdd, 0x37, 0x16, 0xd7, 0xa1, 0xc4, 0x07, 0x98, 0x68, 0x43,
	0x8a, 0x6f, 0xf1, 0x86, 0xf2, 0xd4, 0xfd, 0xf4, 0x86, 0x26, 0xf1, 0x78, 0x0b, 0x9e, 0x7b, 0x4c,
	0xbc, 0x82, 0xa1, 0x29, 0x36, 0x8e, 0xd0, 0xc9, 0xd4, 0xdb, 0x99, 0xc4, 0x03, 0x0f, 0x5c, 0xc9,
	0x54, 0xbc, 0x01, 0x60, 0x8b, 0xda, 0x70, 0x26, 0xcd, 0x5d, 0xf0, 0x64, 0xd6, 0x99, 0x0f, 0xe2,
	0x21, 0xdb, 0x8b, 0x2d, 0xe0, 0xab, 0x61, 0xd1, 0x51, 0xee, 0xa4, 0xe3, 0x92, 0x82, 0x23, 0x4f,
	0x18, 0xb1, 0xbd, 0xaa, 0x21, 0x11, 0xc2, 0xc4, 0x2b, 0xf8, 0x34, 0x2c, 0x27, 0xc1, 0x71, 0xcd,
	0xb9, 0xf0, 0xd8, 0x25, 0x27, 0xe2, 0x25, 0x0c, 0xb4, 0x29, 0x92, 0xcd, 0x9a, 0x4c, 0xd0, 0xf7,
	0x05, 0x6d, 0xee, 0xe1, 0x47, 0x98, 0x1a, 0x52, 0x64, 0x28, 0x5f, 0x53, 0x82, 0x8d, 0x63, 0x64,
	0x47, 0x99, 0x38, 0x81, 0xa3, 0x52, 0xde, 0x6c, 0xea, 0x41, 0xd8, 0x5e, 0xc2, 0xbf, 0x5d, 0x98,
	0xee, 0x1d, 0x22, 0xf1, 0xab, 0x6d, 0xc5, 0x62, 0xd6, 0x9b, 0x8f, 0x16, 0x3f, 0x0e, 0xdb, 0x8a,
	0xff, 0x5a, 0xdf, 0x76, 0xe3, 0x77, 0x07, 0x4e, 0x1e, 0xb3, 0x12, 0x9c, 0xfa, 0x89, 0x48, 0x0f,
	0x9b, 0xc6, 0xde, 0xa2, 0xc5, 0x2f, 0xda, 0xa7, 0x2f, 0xf5, 0xcb, 0x83, 0xad, 0x3a, 0x7b, 0xb0,
	0x55, 0xab, 0x63, 0xbf, 0xe7, 0xa7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x50, 0xcb, 0x96,
	0x0a, 0x04, 0x00, 0x00,
}
