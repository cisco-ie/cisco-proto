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
// source: pim_nbr_summ_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_vrfs_vrf_neighbor_summaries_neighbor_summary

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

type PimNbrSummBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimNbrSummBag_KEYS) Reset()         { *m = PimNbrSummBag_KEYS{} }
func (m *PimNbrSummBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimNbrSummBag_KEYS) ProtoMessage()    {}
func (*PimNbrSummBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a7ffa4c9198eb10, []int{0}
}

func (m *PimNbrSummBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNbrSummBag_KEYS.Unmarshal(m, b)
}
func (m *PimNbrSummBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNbrSummBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimNbrSummBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNbrSummBag_KEYS.Merge(m, src)
}
func (m *PimNbrSummBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimNbrSummBag_KEYS.Size(m)
}
func (m *PimNbrSummBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNbrSummBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimNbrSummBag_KEYS proto.InternalMessageInfo

func (m *PimNbrSummBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimNbrSummBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PimNbrSummBag struct {
	NumberOfNeighbors         int32    `protobuf:"zigzag32,50,opt,name=number_of_neighbors,json=numberOfNeighbors,proto3" json:"number_of_neighbors,omitempty"`
	NumberOfExternalNeighbors int32    `protobuf:"zigzag32,51,opt,name=number_of_external_neighbors,json=numberOfExternalNeighbors,proto3" json:"number_of_external_neighbors,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *PimNbrSummBag) Reset()         { *m = PimNbrSummBag{} }
func (m *PimNbrSummBag) String() string { return proto.CompactTextString(m) }
func (*PimNbrSummBag) ProtoMessage()    {}
func (*PimNbrSummBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a7ffa4c9198eb10, []int{1}
}

func (m *PimNbrSummBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNbrSummBag.Unmarshal(m, b)
}
func (m *PimNbrSummBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNbrSummBag.Marshal(b, m, deterministic)
}
func (m *PimNbrSummBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNbrSummBag.Merge(m, src)
}
func (m *PimNbrSummBag) XXX_Size() int {
	return xxx_messageInfo_PimNbrSummBag.Size(m)
}
func (m *PimNbrSummBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNbrSummBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimNbrSummBag proto.InternalMessageInfo

func (m *PimNbrSummBag) GetNumberOfNeighbors() int32 {
	if m != nil {
		return m.NumberOfNeighbors
	}
	return 0
}

func (m *PimNbrSummBag) GetNumberOfExternalNeighbors() int32 {
	if m != nil {
		return m.NumberOfExternalNeighbors
	}
	return 0
}

func init() {
	proto.RegisterType((*PimNbrSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.neighbor_summaries.neighbor_summary.pim_nbr_summ_bag_KEYS")
	proto.RegisterType((*PimNbrSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.vrfs.vrf.neighbor_summaries.neighbor_summary.pim_nbr_summ_bag")
}

func init() { proto.RegisterFile("pim_nbr_summ_bag.proto", fileDescriptor_7a7ffa4c9198eb10) }

var fileDescriptor_7a7ffa4c9198eb10 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0x39, 0x0b, 0x3f, 0x16, 0x14, 0xb3, 0xa2, 0x24, 0x60, 0x11, 0x02, 0x42, 0xaa, 0x2d,
	0x8c, 0xbd, 0x55, 0x2a, 0x21, 0xc2, 0x59, 0xa5, 0x1a, 0x76, 0x8f, 0xd9, 0x38, 0xe0, 0x7e, 0x30,
	0xbb, 0x59, 0x62, 0xed, 0x3f, 0x2e, 0x77, 0xe7, 0xc5, 0x70, 0xcd, 0x14, 0xef, 0xfd, 0xde, 0x9b,
	0x61, 0xc4, 0x43, 0x24, 0x07, 0xde, 0x30, 0xa4, 0xbd, 0x73, 0x60, 0xf4, 0x4e, 0x45, 0x0e, 0x39,
	0xc8, 0xba, 0xa1, 0xd4, 0x04, 0xa0, 0x90, 0xe0, 0xc0, 0x40, 0xb1, 0xbc, 0x40, 0x4b, 0x86, 0x88,
	0xac, 0x22, 0x39, 0xa5, 0x9b, 0x4c, 0x05, 0x55, 0x61, 0x9b, 0xda, 0xa1, 0x3c, 0xd2, 0xee, 0xd3,
	0x84, 0xbe, 0x47, 0x33, 0x61, 0x1a, 0x4b, 0xdf, 0x8b, 0xad, 0xb8, 0x1f, 0x6f, 0x83, 0xb7, 0xf5,
	0xf6, 0x43, 0xce, 0xc4, 0x65, 0x61, 0x0b, 0x5e, 0x3b, 0x9c, 0x56, 0xf3, 0x6a, 0x79, 0x55, 0x5f,
	0x14, 0xb6, 0x1b, 0xed, 0x50, 0x3e, 0x89, 0x1b, 0xf2, 0x19, 0xd9, 0xea, 0x06, 0x7b, 0xe0, 0xac,
	0x03, 0xae, 0x8f, 0x6a, 0x8b, 0x2d, 0x7e, 0x2a, 0x71, 0x3b, 0xee, 0x96, 0x4a, 0xdc, 0xf9, 0xbd,
	0x33, 0xc8, 0x10, 0x2c, 0x0c, 0xd7, 0xa4, 0xe9, 0xf3, 0xbc, 0x5a, 0x4e, 0xea, 0x49, 0x6f, 0xbd,
	0xdb, 0xcd, 0x60, 0xc8, 0x57, 0xf1, 0xf8, 0xcf, 0xe3, 0x21, 0x23, 0x7b, 0xfd, 0x75, 0x12, 0x5c,
	0x75, 0xc1, 0xd9, 0x10, 0x5c, 0xff, 0x11, 0xc7, 0x02, 0x73, 0xde, 0xfd, 0x6e, 0xf5, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x2b, 0xbf, 0x72, 0x1b, 0x55, 0x01, 0x00, 0x00,
}