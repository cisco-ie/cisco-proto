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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_gre_gre_hashes_gre_hash

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
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Ifname               string   `protobuf:"bytes,3,opt,name=ifname,proto3" json:"ifname,omitempty"`
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
	proto.RegisterType((*PimGreHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.gre.gre_hashes.gre_hash.pim_gre_hash_bag_KEYS")
	proto.RegisterType((*PimGreHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.gre.gre_hashes.gre_hash.pim_gre_hash_bag")
}

func init() { proto.RegisterFile("pim_gre_hash_bag.proto", fileDescriptor_7705297c4213ce4e) }

var fileDescriptor_7705297c4213ce4e = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0x42, 0xc1, 0x05, 0xa5, 0xac, 0x58, 0x72, 0x94, 0x82, 0xe0, 0x41, 0x56, 0x50,
	0xf1, 0xac, 0x07, 0x41, 0xf1, 0x56, 0x4f, 0x9e, 0x86, 0xed, 0x66, 0x92, 0x0c, 0xd8, 0x9d, 0x65,
	0x67, 0x1a, 0xf2, 0x0f, 0xfc, 0xdb, 0x92, 0xd4, 0x94, 0xd2, 0xdb, 0xbe, 0xf7, 0xbe, 0xb7, 0x33,
	0x8c, 0x59, 0x24, 0xda, 0x40, 0x93, 0x11, 0x5a, 0x2f, 0x2d, 0xac, 0x7d, 0xe3, 0x52, 0x66, 0x65,
	0xbb, 0x0a, 0x24, 0x81, 0x81, 0x58, 0xa0, 0xcf, 0x40, 0xa9, 0x7b, 0x82, 0x81, 0xe4, 0x84, 0xd9,
	0x51, 0xea, 0x9e, 0x07, 0xe5, 0x7c, 0x50, 0xea, 0xd0, 0x55, 0x58, 0xfb, 0xed, 0x8f, 0x42, 0xe0,
	0xa8, 0xd8, 0xab, 0x6b, 0x32, 0xba, 0xe9, 0x4b, 0x94, 0xfd, 0x73, 0xf9, 0x5b, 0x98, 0xab, 0xe3,
	0x71, 0xf0, 0xf9, 0xf6, 0xfd, 0x65, 0x6f, 0xcc, 0x85, 0xf0, 0x36, 0x07, 0x04, 0x5f, 0x55, 0x19,
	0x45, 0xca, 0xe2, 0xba, 0xb8, 0x3d, 0x5b, 0x9d, 0xef, 0xdc, 0xd7, 0x9d, 0x69, 0xef, 0xcd, 0x65,
	0x85, 0xa2, 0x14, 0xbd, 0x12, 0xc7, 0x3d, 0x7b, 0x32, 0xb2, 0xf6, 0x20, 0x9a, 0x0a, 0x0b, 0x33,
	0xa3, 0x3a, 0xfa, 0x0d, 0x96, 0xa7, 0x23, 0xf3, 0xaf, 0x96, 0x2f, 0x66, 0x7e, 0xbc, 0x88, 0xbd,
	0x33, 0x36, 0x62, 0xaf, 0xd0, 0x72, 0x02, 0x8a, 0x8a, 0xb9, 0xf6, 0x01, 0xcb, 0x87, 0xb1, 0x37,
	0x1f, 0x92, 0x77, 0x4e, 0x1f, 0x93, 0xbf, 0x9e, 0x8d, 0x67, 0x7a, 0xfc, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xe9, 0x19, 0xac, 0x70, 0x40, 0x01, 0x00, 0x00,
}
