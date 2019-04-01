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
// source: pim_gre_hash_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_gre_gre_hashes_gre_hash

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

type PimGreHashBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	SourceAddress        string   `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,3,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Ifname               string   `protobuf:"bytes,4,opt,name=ifname,proto3" json:"ifname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimGreHashBag_KEYS) Reset()         { *m = PimGreHashBag_KEYS{} }
func (m *PimGreHashBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimGreHashBag_KEYS) ProtoMessage()    {}
func (*PimGreHashBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7705297c4213ce4e, []int{0}
}

func (m *PimGreHashBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGreHashBag_KEYS.Unmarshal(m, b)
}
func (m *PimGreHashBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGreHashBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimGreHashBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGreHashBag_KEYS.Merge(m, src)
}
func (m *PimGreHashBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimGreHashBag_KEYS.Size(m)
}
func (m *PimGreHashBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGreHashBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimGreHashBag_KEYS proto.InternalMessageInfo

func (m *PimGreHashBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimGreHashBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimGreHashBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *PimGreHashBag_KEYS) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

type PimGreHashBag struct {
	NextHopInterface     string   `protobuf:"bytes,50,opt,name=next_hop_interface,json=nextHopInterface,proto3" json:"next_hop_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimGreHashBag) Reset()         { *m = PimGreHashBag{} }
func (m *PimGreHashBag) String() string { return proto.CompactTextString(m) }
func (*PimGreHashBag) ProtoMessage()    {}
func (*PimGreHashBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7705297c4213ce4e, []int{1}
}

func (m *PimGreHashBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimGreHashBag.Unmarshal(m, b)
}
func (m *PimGreHashBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimGreHashBag.Marshal(b, m, deterministic)
}
func (m *PimGreHashBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimGreHashBag.Merge(m, src)
}
func (m *PimGreHashBag) XXX_Size() int {
	return xxx_messageInfo_PimGreHashBag.Size(m)
}
func (m *PimGreHashBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimGreHashBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimGreHashBag proto.InternalMessageInfo

func (m *PimGreHashBag) GetNextHopInterface() string {
	if m != nil {
		return m.NextHopInterface
	}
	return ""
}

func init() {
	proto.RegisterType((*PimGreHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.gre.gre_hashes.gre_hash.pim_gre_hash_bag_KEYS")
	proto.RegisterType((*PimGreHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.gre.gre_hashes.gre_hash.pim_gre_hash_bag")
}

func init() { proto.RegisterFile("pim_gre_hash_bag.proto", fileDescriptor_7705297c4213ce4e) }

var fileDescriptor_7705297c4213ce4e = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0x95, 0x53, 0x03, 0xca, 0x11, 0xf1, 0x58, 0x3b, 0x39, 0x10, 0x2c, 0x24, 0x82,
	0x8a, 0xb5, 0x16, 0x82, 0x72, 0x60, 0xa1, 0x95, 0xd5, 0x90, 0xcb, 0x4e, 0x76, 0xa7, 0xd8, 0x24,
	0x4c, 0x62, 0xb8, 0xdf, 0xe3, 0x2f, 0x95, 0xcd, 0xb9, 0x87, 0x6c, 0x13, 0x32, 0xef, 0x7d, 0x2f,
	0x2f, 0x8c, 0x58, 0x04, 0xea, 0xa1, 0x65, 0x84, 0x4e, 0xc7, 0x0e, 0xd6, 0xba, 0x55, 0x81, 0x7d,
	0xf2, 0x72, 0x65, 0x28, 0x1a, 0x0f, 0xe4, 0x23, 0x6c, 0x18, 0x28, 0xe4, 0x07, 0x18, 0x48, 0x1f,
	0x90, 0x15, 0x85, 0xfc, 0x38, 0x4c, 0x4a, 0x9b, 0x44, 0x19, 0x55, 0x66, 0x1b, 0x87, 0x43, 0xb5,
	0x8c, 0x6a, 0x7c, 0x0b, 0xe3, 0xee, 0xba, 0xfc, 0xa9, 0xc4, 0xf9, 0xb4, 0x07, 0x56, 0x2f, 0x5f,
	0x9f, 0xf2, 0x42, 0x1c, 0x65, 0xb6, 0xe0, 0x74, 0x8f, 0x75, 0x75, 0x59, 0x5d, 0x1f, 0x7f, 0x1c,
	0x66, 0xb6, 0xef, 0xba, 0x47, 0x79, 0x25, 0x4e, 0xa3, 0xff, 0x66, 0x83, 0xa0, 0x9b, 0x86, 0x31,
	0xc6, 0x7a, 0xaf, 0x00, 0x27, 0x5b, 0xf5, 0x79, 0x2b, 0xca, 0x5b, 0x71, 0xd6, 0x60, 0x4c, 0xe4,
	0x74, 0x22, 0xef, 0x76, 0xec, 0x7e, 0x61, 0xe5, 0x3f, 0x6b, 0x0c, 0x2c, 0xc4, 0x8c, 0x6c, 0x29,
	0x3c, 0x28, 0xcc, 0xdf, 0xb4, 0x7c, 0x12, 0xf3, 0xe9, 0x1f, 0xe5, 0x8d, 0x90, 0x0e, 0x37, 0x09,
	0x3a, 0x1f, 0x80, 0x5c, 0x42, 0xb6, 0xda, 0x60, 0x7d, 0x57, 0x72, 0xf3, 0xc1, 0x79, 0xf5, 0xe1,
	0x6d, 0xd4, 0xd7, 0xb3, 0xb2, 0xba, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x53, 0xd8,
	0xd5, 0x54, 0x01, 0x00, 0x00,
}
