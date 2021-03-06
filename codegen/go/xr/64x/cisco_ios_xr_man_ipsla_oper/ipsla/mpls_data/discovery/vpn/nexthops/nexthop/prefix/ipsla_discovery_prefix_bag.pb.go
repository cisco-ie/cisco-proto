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
// source: ipsla_discovery_prefix_bag.proto

package cisco_ios_xr_man_ipsla_oper_ipsla_mpls_data_discovery_vpn_nexthops_nexthop_prefix

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

type IpslaDiscoveryPrefixBag_KEYS struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaDiscoveryPrefixBag_KEYS) Reset()         { *m = IpslaDiscoveryPrefixBag_KEYS{} }
func (m *IpslaDiscoveryPrefixBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpslaDiscoveryPrefixBag_KEYS) ProtoMessage()    {}
func (*IpslaDiscoveryPrefixBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d205a7526b48c057, []int{0}
}

func (m *IpslaDiscoveryPrefixBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS.Unmarshal(m, b)
}
func (m *IpslaDiscoveryPrefixBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IpslaDiscoveryPrefixBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS.Merge(m, src)
}
func (m *IpslaDiscoveryPrefixBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS.Size(m)
}
func (m *IpslaDiscoveryPrefixBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaDiscoveryPrefixBag_KEYS proto.InternalMessageInfo

func (m *IpslaDiscoveryPrefixBag_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type IpslaDiscoveryPrefixBag struct {
	TargetAddress        string   `protobuf:"bytes,50,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	TargetMask           uint32   `protobuf:"varint,51,opt,name=target_mask,json=targetMask,proto3" json:"target_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaDiscoveryPrefixBag) Reset()         { *m = IpslaDiscoveryPrefixBag{} }
func (m *IpslaDiscoveryPrefixBag) String() string { return proto.CompactTextString(m) }
func (*IpslaDiscoveryPrefixBag) ProtoMessage()    {}
func (*IpslaDiscoveryPrefixBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d205a7526b48c057, []int{1}
}

func (m *IpslaDiscoveryPrefixBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag.Unmarshal(m, b)
}
func (m *IpslaDiscoveryPrefixBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag.Marshal(b, m, deterministic)
}
func (m *IpslaDiscoveryPrefixBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaDiscoveryPrefixBag.Merge(m, src)
}
func (m *IpslaDiscoveryPrefixBag) XXX_Size() int {
	return xxx_messageInfo_IpslaDiscoveryPrefixBag.Size(m)
}
func (m *IpslaDiscoveryPrefixBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaDiscoveryPrefixBag.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaDiscoveryPrefixBag proto.InternalMessageInfo

func (m *IpslaDiscoveryPrefixBag) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *IpslaDiscoveryPrefixBag) GetTargetMask() uint32 {
	if m != nil {
		return m.TargetMask
	}
	return 0
}

func init() {
	proto.RegisterType((*IpslaDiscoveryPrefixBag_KEYS)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.mpls_data.discovery.vpn.nexthops.nexthop.prefix.ipsla_discovery_prefix_bag_KEYS")
	proto.RegisterType((*IpslaDiscoveryPrefixBag)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.mpls_data.discovery.vpn.nexthops.nexthop.prefix.ipsla_discovery_prefix_bag")
}

func init() { proto.RegisterFile("ipsla_discovery_prefix_bag.proto", fileDescriptor_d205a7526b48c057) }

var fileDescriptor_d205a7526b48c057 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xbf, 0x6f, 0x83, 0x30,
	0x10, 0x85, 0xc5, 0xd2, 0xaa, 0xae, 0xe8, 0xe0, 0x09, 0x75, 0x01, 0x21, 0x55, 0x62, 0xf2, 0x50,
	0xc6, 0x4e, 0x1d, 0x32, 0x45, 0x19, 0x42, 0xa6, 0x4c, 0xa7, 0x03, 0x3b, 0xc4, 0xe2, 0x87, 0x2d,
	0x9f, 0x85, 0xc8, 0x7f, 0x1f, 0x05, 0x87, 0x6c, 0x6c, 0xcf, 0x9f, 0xef, 0x7b, 0xa7, 0x63, 0x99,
	0xb6, 0xd4, 0x23, 0x48, 0x4d, 0x8d, 0x99, 0x94, 0xbb, 0x81, 0x75, 0xea, 0xa2, 0x67, 0xa8, 0xb1,
	0x15, 0xd6, 0x19, 0x6f, 0xf8, 0xb1, 0x79, 0xfc, 0x81, 0x36, 0x04, 0xb3, 0x83, 0x01, 0x47, 0x08,
	0x8a, 0xb1, 0xca, 0x89, 0x25, 0x8a, 0xc1, 0xf6, 0x04, 0x12, 0x3d, 0x8a, 0x57, 0x8f, 0x98, 0xec,
	0x28, 0x46, 0x35, 0xfb, 0xab, 0xb1, 0xb4, 0x06, 0x11, 0xca, 0xf3, 0x3f, 0x96, 0x6e, 0xaf, 0x85,
	0xfd, 0xee, 0x7c, 0xe2, 0x09, 0x7b, 0x47, 0x29, 0x9d, 0x22, 0x4a, 0xa2, 0x2c, 0x2a, 0x3e, 0xaa,
	0xf5, 0x99, 0x4b, 0xf6, 0xbd, 0x2d, 0xf3, 0x1f, 0xf6, 0xe5, 0xd1, 0xb5, 0xca, 0xc3, 0xaa, 0xff,
	0x2e, 0x7a, 0x1c, 0xe8, 0x7f, 0x80, 0x3c, 0x65, 0x9f, 0xcf, 0xb1, 0x01, 0xa9, 0x4b, 0xca, 0x2c,
	0x2a, 0xe2, 0x8a, 0x05, 0x74, 0x40, 0xea, 0xea, 0xb7, 0xe5, 0xf8, 0xf2, 0x1e, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0x0e, 0x14, 0x8a, 0x20, 0x01, 0x00, 0x00,
}
