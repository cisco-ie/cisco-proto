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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_gre_gre_hashes_gre_hash

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
	proto.RegisterType((*PimGreHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.gre.gre_hashes.gre_hash.pim_gre_hash_bag_KEYS")
	proto.RegisterType((*PimGreHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.gre.gre_hashes.gre_hash.pim_gre_hash_bag")
}

func init() { proto.RegisterFile("pim_gre_hash_bag.proto", fileDescriptor_7705297c4213ce4e) }

var fileDescriptor_7705297c4213ce4e = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0x03, 0x41,
	0x0c, 0x86, 0xa9, 0x42, 0xc1, 0x01, 0xa5, 0x8c, 0x58, 0xf6, 0x28, 0x05, 0xc1, 0x83, 0x8c, 0xa0,
	0x3e, 0x80, 0x1e, 0x04, 0x45, 0xf0, 0xa0, 0x27, 0x4f, 0x61, 0x76, 0x26, 0xbb, 0x1b, 0x70, 0x27,
	0xc3, 0x24, 0x95, 0xf5, 0x09, 0x7c, 0x6d, 0xe9, 0xd6, 0x2d, 0xd2, 0x5b, 0xf2, 0xff, 0xdf, 0x9f,
	0x84, 0x98, 0x65, 0xa6, 0x1e, 0xda, 0x82, 0xd0, 0x79, 0xe9, 0xa0, 0xf6, 0xad, 0xcb, 0x85, 0x95,
	0xed, 0x6b, 0x20, 0x09, 0x0c, 0xc4, 0x02, 0x43, 0x01, 0xca, 0x5f, 0x77, 0xb0, 0x21, 0x39, 0x63,
	0x71, 0x99, 0x7a, 0x27, 0xea, 0x53, 0xac, 0xbf, 0x5d, 0xc4, 0xc6, 0xaf, 0x3f, 0x15, 0x02, 0x27,
	0xc5, 0x41, 0x5d, 0x5b, 0xd0, 0x4d, 0xe3, 0x50, 0x76, 0xe5, 0xea, 0x67, 0x66, 0xce, 0xf6, 0x57,
	0xc1, 0xcb, 0xe3, 0xc7, 0xbb, 0xbd, 0x30, 0x27, 0xc2, 0xeb, 0x12, 0x10, 0x7c, 0x8c, 0x05, 0x45,
	0xaa, 0xd9, 0xf9, 0xec, 0xf2, 0xe8, 0xed, 0x78, 0xab, 0x3e, 0x6c, 0x45, 0x7b, 0x6d, 0x4e, 0x23,
	0x8a, 0x52, 0xf2, 0x4a, 0x9c, 0x76, 0xec, 0xc1, 0xc8, 0xda, 0x7f, 0xd6, 0x14, 0x58, 0x9a, 0x39,
	0x35, 0xc9, 0xf7, 0x58, 0x1d, 0x8e, 0xcc, 0x5f, 0xb7, 0xba, 0x37, 0x8b, 0xfd, 0x43, 0xec, 0x95,
	0xb1, 0x09, 0x07, 0x85, 0x8e, 0x33, 0x50, 0x52, 0x2c, 0x8d, 0x0f, 0x58, 0xdd, 0x8c, 0xb9, 0xc5,
	0xc6, 0x79, 0xe2, 0xfc, 0x3c, 0xe9, 0xf5, 0x7c, 0x7c, 0xd1, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0x08, 0xe4, 0x2e, 0x3c, 0x01, 0x00, 0x00,
}
