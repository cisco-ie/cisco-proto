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
// source: mpls_te_affinity_map_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_affinity_map

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

type MplsTeAffinityMapBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAffinityMapBag_KEYS) Reset()         { *m = MplsTeAffinityMapBag_KEYS{} }
func (m *MplsTeAffinityMapBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeAffinityMapBag_KEYS) ProtoMessage()    {}
func (*MplsTeAffinityMapBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_16bfef2ee65f8ac1, []int{0}
}

func (m *MplsTeAffinityMapBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAffinityMapBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeAffinityMapBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAffinityMapBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeAffinityMapBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAffinityMapBag_KEYS.Merge(m, src)
}
func (m *MplsTeAffinityMapBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeAffinityMapBag_KEYS.Size(m)
}
func (m *MplsTeAffinityMapBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAffinityMapBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAffinityMapBag_KEYS proto.InternalMessageInfo

type MplsTeAffinityMapInfo struct {
	AffinityName          string   `protobuf:"bytes,1,opt,name=affinity_name,json=affinityName,proto3" json:"affinity_name,omitempty"`
	AffinityValue         uint32   `protobuf:"varint,2,opt,name=affinity_value,json=affinityValue,proto3" json:"affinity_value,omitempty"`
	AffinityExtendedValue []uint32 `protobuf:"varint,3,rep,packed,name=affinity_extended_value,json=affinityExtendedValue,proto3" json:"affinity_extended_value,omitempty"`
	AffinityBitPosition   uint32   `protobuf:"varint,4,opt,name=affinity_bit_position,json=affinityBitPosition,proto3" json:"affinity_bit_position,omitempty"`
	AffinityTableId       string   `protobuf:"bytes,5,opt,name=affinity_table_id,json=affinityTableId,proto3" json:"affinity_table_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *MplsTeAffinityMapInfo) Reset()         { *m = MplsTeAffinityMapInfo{} }
func (m *MplsTeAffinityMapInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAffinityMapInfo) ProtoMessage()    {}
func (*MplsTeAffinityMapInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16bfef2ee65f8ac1, []int{1}
}

func (m *MplsTeAffinityMapInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAffinityMapInfo.Unmarshal(m, b)
}
func (m *MplsTeAffinityMapInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAffinityMapInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeAffinityMapInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAffinityMapInfo.Merge(m, src)
}
func (m *MplsTeAffinityMapInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAffinityMapInfo.Size(m)
}
func (m *MplsTeAffinityMapInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAffinityMapInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAffinityMapInfo proto.InternalMessageInfo

func (m *MplsTeAffinityMapInfo) GetAffinityName() string {
	if m != nil {
		return m.AffinityName
	}
	return ""
}

func (m *MplsTeAffinityMapInfo) GetAffinityValue() uint32 {
	if m != nil {
		return m.AffinityValue
	}
	return 0
}

func (m *MplsTeAffinityMapInfo) GetAffinityExtendedValue() []uint32 {
	if m != nil {
		return m.AffinityExtendedValue
	}
	return nil
}

func (m *MplsTeAffinityMapInfo) GetAffinityBitPosition() uint32 {
	if m != nil {
		return m.AffinityBitPosition
	}
	return 0
}

func (m *MplsTeAffinityMapInfo) GetAffinityTableId() string {
	if m != nil {
		return m.AffinityTableId
	}
	return ""
}

type MplsTeAffinityMapBag struct {
	AffinityMapArray     []*MplsTeAffinityMapInfo `protobuf:"bytes,50,rep,name=affinity_map_array,json=affinityMapArray,proto3" json:"affinity_map_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsTeAffinityMapBag) Reset()         { *m = MplsTeAffinityMapBag{} }
func (m *MplsTeAffinityMapBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeAffinityMapBag) ProtoMessage()    {}
func (*MplsTeAffinityMapBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_16bfef2ee65f8ac1, []int{2}
}

func (m *MplsTeAffinityMapBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAffinityMapBag.Unmarshal(m, b)
}
func (m *MplsTeAffinityMapBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAffinityMapBag.Marshal(b, m, deterministic)
}
func (m *MplsTeAffinityMapBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAffinityMapBag.Merge(m, src)
}
func (m *MplsTeAffinityMapBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeAffinityMapBag.Size(m)
}
func (m *MplsTeAffinityMapBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAffinityMapBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAffinityMapBag proto.InternalMessageInfo

func (m *MplsTeAffinityMapBag) GetAffinityMapArray() []*MplsTeAffinityMapInfo {
	if m != nil {
		return m.AffinityMapArray
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeAffinityMapBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.affinity_map.mpls_te_affinity_map_bag_KEYS")
	proto.RegisterType((*MplsTeAffinityMapInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.affinity_map.mpls_te_affinity_map_info")
	proto.RegisterType((*MplsTeAffinityMapBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.affinity_map.mpls_te_affinity_map_bag")
}

func init() { proto.RegisterFile("mpls_te_affinity_map_bag.proto", fileDescriptor_16bfef2ee65f8ac1) }

var fileDescriptor_16bfef2ee65f8ac1 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x55, 0xc1, 0x68, 0xfd, 0x88, 0x88, 0xf1, 0xa0, 0x96, 0x8a, 0x50, 0x3c, 0xf4,
	0x50, 0xc1, 0xbb, 0xc2, 0x1e, 0x16, 0x51, 0xa4, 0x8a, 0xe0, 0x69, 0x48, 0xb7, 0xa9, 0x0c, 0xb4,
	0x4d, 0x48, 0xa3, 0xee, 0xfe, 0x0e, 0x7f, 0xb0, 0xd2, 0xda, 0x04, 0x05, 0x7b, 0xd8, 0x63, 0x9f,
	0xe7, 0x7d, 0x3b, 0x61, 0x86, 0x9c, 0xd4, 0xaa, 0x6a, 0xc1, 0x08, 0xe0, 0x65, 0x89, 0x0d, 0x9a,
	0x05, 0xd4, 0x5c, 0x41, 0xce, 0x5f, 0x13, 0xa5, 0xa5, 0x91, 0x34, 0x99, 0x61, 0x3b, 0x93, 0x80,
	0xb2, 0x85, 0xb9, 0x06, 0x1b, 0x96, 0x4a, 0xe8, 0x64, 0xf8, 0x48, 0x7e, 0x37, 0xa3, 0x53, 0x72,
	0x3c, 0xf6, 0x47, 0xb8, 0x9d, 0xbc, 0x3c, 0x46, 0x5f, 0x1e, 0x39, 0xfa, 0x37, 0x81, 0x4d, 0x29,
	0xe9, 0x19, 0x09, 0x1c, 0x6c, 0x78, 0x2d, 0x98, 0x17, 0x7a, 0xf1, 0x46, 0xb6, 0x65, 0xe1, 0x3d,
	0xaf, 0x05, 0x3d, 0x27, 0xdb, 0x2e, 0xf4, 0xce, 0xab, 0x37, 0xc1, 0x56, 0x42, 0x2f, 0x0e, 0x32,
	0x57, 0x7d, 0xee, 0x20, 0xbd, 0x22, 0x87, 0x2e, 0x26, 0xe6, 0x46, 0x34, 0x85, 0x28, 0x86, 0xbc,
	0x1f, 0xfa, 0x71, 0x90, 0x1d, 0x58, 0x3d, 0x19, 0xec, 0x4f, 0x2f, 0x25, 0x4e, 0x40, 0x8e, 0x06,
	0x94, 0x6c, 0xd1, 0xa0, 0x6c, 0xd8, 0x6a, 0x3f, 0x65, 0xdf, 0xca, 0x1b, 0x34, 0x0f, 0x83, 0xa2,
	0x17, 0x64, 0xcf, 0x75, 0x0c, 0xcf, 0x2b, 0x01, 0x58, 0xb0, 0xb5, 0xfe, 0xed, 0x3b, 0x56, 0x3c,
	0x75, 0x7c, 0x5a, 0x44, 0x9f, 0x1e, 0x61, 0x63, 0x3b, 0xa2, 0x1f, 0x84, 0xfe, 0x61, 0x5c, 0x6b,
	0xbe, 0x60, 0x69, 0xe8, 0xc7, 0x9b, 0xe9, 0x74, 0xc9, 0x63, 0x24, 0xa3, 0x7b, 0xce, 0x76, 0x2d,
	0xba, 0xe3, 0xea, 0xba, 0x1b, 0x91, 0xaf, 0xf7, 0xf7, 0xbe, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xd6, 0x84, 0xb8, 0x9c, 0x11, 0x02, 0x00, 0x00,
}
