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

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_affinity_map

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
	proto.RegisterType((*MplsTeAffinityMapBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.affinity_map.mpls_te_affinity_map_bag_KEYS")
	proto.RegisterType((*MplsTeAffinityMapInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.affinity_map.mpls_te_affinity_map_info")
	proto.RegisterType((*MplsTeAffinityMapBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.affinity_map.mpls_te_affinity_map_bag")
}

func init() { proto.RegisterFile("mpls_te_affinity_map_bag.proto", fileDescriptor_16bfef2ee65f8ac1) }

var fileDescriptor_16bfef2ee65f8ac1 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0xa9, 0xab, 0x82, 0xd1, 0xfa, 0x27, 0x22, 0xc6, 0x83, 0x5a, 0x2a, 0x42, 0xf1, 0xd0,
	0x43, 0x85, 0xbd, 0x2b, 0xec, 0x41, 0x44, 0xd1, 0x2a, 0x82, 0xa7, 0x61, 0xba, 0x4d, 0x25, 0xd0,
	0x26, 0xa1, 0x8d, 0xb2, 0x3d, 0xf9, 0x20, 0x3e, 0xac, 0xd2, 0xda, 0x04, 0x05, 0x7b, 0xf1, 0x98,
	0xef, 0xfb, 0xcd, 0x4c, 0x98, 0x21, 0x47, 0x95, 0x2e, 0x1b, 0x30, 0x1c, 0xb0, 0x28, 0x84, 0x14,
	0xa6, 0x85, 0x0a, 0x35, 0x64, 0xf8, 0x12, 0xeb, 0x5a, 0x19, 0x45, 0xa7, 0x73, 0xd1, 0xcc, 0x15,
	0x08, 0xd5, 0xc0, 0xa2, 0x06, 0x1b, 0x56, 0x9a, 0xd7, 0xb1, 0x7d, 0x34, 0x06, 0x65, 0x9e, 0xb5,
	0xf1, 0xcf, 0x0e, 0xe1, 0x31, 0x39, 0x1c, 0xeb, 0x0c, 0xd7, 0xb3, 0xe7, 0x87, 0xf0, 0xd3, 0x23,
	0x07, 0x7f, 0x26, 0x84, 0x2c, 0x14, 0x3d, 0x21, 0xbe, 0x83, 0x12, 0x2b, 0xce, 0xbc, 0xc0, 0x8b,
	0xd6, 0xd2, 0x0d, 0x0b, 0x6f, 0xb1, 0xe2, 0xf4, 0x94, 0x6c, 0xba, 0xd0, 0x1b, 0x96, 0xaf, 0x9c,
	0x2d, 0x05, 0x5e, 0xe4, 0xa7, 0xae, 0xf4, 0xa9, 0x83, 0x74, 0x4a, 0xf6, 0x5d, 0x8c, 0x2f, 0x0c,
	0x97, 0x39, 0xcf, 0x87, 0xfc, 0x24, 0x98, 0x44, 0x7e, 0xba, 0x67, 0xf5, 0x6c, 0xb0, 0xdf, 0x75,
	0x09, 0x71, 0x02, 0x32, 0x61, 0x40, 0xab, 0x46, 0x18, 0xa1, 0x24, 0x5b, 0xee, 0xa7, 0xec, 0x5a,
	0x79, 0x29, 0xcc, 0xdd, 0xa0, 0xe8, 0x19, 0xd9, 0x71, 0x35, 0x06, 0xb3, 0x92, 0x83, 0xc8, 0xd9,
	0x4a, 0xff, 0xf7, 0x2d, 0x2b, 0x1e, 0x3b, 0x7e, 0x95, 0x87, 0x1f, 0x1e, 0x61, 0x63, 0x3b, 0xa2,
	0xef, 0x84, 0xfe, 0x62, 0x58, 0xd7, 0xd8, 0xb2, 0x24, 0x98, 0x44, 0xeb, 0xc9, 0x7d, 0xfc, 0xbf,
	0xa3, 0xc4, 0xa3, 0xfb, 0x4e, 0xb7, 0x2d, 0xba, 0x41, 0x7d, 0xd1, 0x8d, 0xca, 0x56, 0xfb, 0xfb,
	0x9f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x61, 0x1f, 0x01, 0x16, 0x21, 0x02, 0x00, 0x00,
}
