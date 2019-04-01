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
// source: srms_l2_adj_t_b.proto

package cisco_ios_xr_segment_routing_ms_oper_srms_adjacency_sid_l2_adjacency_interfaces_interface_address_family_ipv4_sid_record

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

type SrmsL2AdjTB_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SidType              string   `protobuf:"bytes,2,opt,name=sid_type,json=sidType,proto3" json:"sid_type,omitempty"`
	SidValue             uint32   `protobuf:"varint,3,opt,name=sid_value,json=sidValue,proto3" json:"sid_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsL2AdjTB_KEYS) Reset()         { *m = SrmsL2AdjTB_KEYS{} }
func (m *SrmsL2AdjTB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SrmsL2AdjTB_KEYS) ProtoMessage()    {}
func (*SrmsL2AdjTB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_179677fb0b32cf11, []int{0}
}

func (m *SrmsL2AdjTB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsL2AdjTB_KEYS.Unmarshal(m, b)
}
func (m *SrmsL2AdjTB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsL2AdjTB_KEYS.Marshal(b, m, deterministic)
}
func (m *SrmsL2AdjTB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsL2AdjTB_KEYS.Merge(m, src)
}
func (m *SrmsL2AdjTB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SrmsL2AdjTB_KEYS.Size(m)
}
func (m *SrmsL2AdjTB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsL2AdjTB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsL2AdjTB_KEYS proto.InternalMessageInfo

func (m *SrmsL2AdjTB_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *SrmsL2AdjTB_KEYS) GetSidType() string {
	if m != nil {
		return m.SidType
	}
	return ""
}

func (m *SrmsL2AdjTB_KEYS) GetSidValue() uint32 {
	if m != nil {
		return m.SidValue
	}
	return 0
}

type AddrT struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddrT) Reset()         { *m = AddrT{} }
func (m *AddrT) String() string { return proto.CompactTextString(m) }
func (*AddrT) ProtoMessage()    {}
func (*AddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_179677fb0b32cf11, []int{1}
}

func (m *AddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddrT.Unmarshal(m, b)
}
func (m *AddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddrT.Marshal(b, m, deterministic)
}
func (m *AddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddrT.Merge(m, src)
}
func (m *AddrT) XXX_Size() int {
	return xxx_messageInfo_AddrT.Size(m)
}
func (m *AddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_AddrT.DiscardUnknown(m)
}

var xxx_messageInfo_AddrT proto.InternalMessageInfo

func (m *AddrT) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *AddrT) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *AddrT) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type SrmsL2AdjTB struct {
	InterfaceName        string   `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SidValueXr           uint32   `protobuf:"varint,51,opt,name=sid_value_xr,json=sidValueXr,proto3" json:"sid_value_xr,omitempty"`
	SidTypeXr            uint32   `protobuf:"varint,52,opt,name=sid_type_xr,json=sidTypeXr,proto3" json:"sid_type_xr,omitempty"`
	AddressFamily        uint32   `protobuf:"varint,53,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	HasNexthop           bool     `protobuf:"varint,54,opt,name=has_nexthop,json=hasNexthop,proto3" json:"has_nexthop,omitempty"`
	InterfaceCount       int32    `protobuf:"zigzag32,55,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	InterfaceDeleteCount int32    `protobuf:"zigzag32,56,opt,name=interface_delete_count,json=interfaceDeleteCount,proto3" json:"interface_delete_count,omitempty"`
	NexthopAddress       *AddrT   `protobuf:"bytes,57,opt,name=nexthop_address,json=nexthopAddress,proto3" json:"nexthop_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsL2AdjTB) Reset()         { *m = SrmsL2AdjTB{} }
func (m *SrmsL2AdjTB) String() string { return proto.CompactTextString(m) }
func (*SrmsL2AdjTB) ProtoMessage()    {}
func (*SrmsL2AdjTB) Descriptor() ([]byte, []int) {
	return fileDescriptor_179677fb0b32cf11, []int{2}
}

func (m *SrmsL2AdjTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsL2AdjTB.Unmarshal(m, b)
}
func (m *SrmsL2AdjTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsL2AdjTB.Marshal(b, m, deterministic)
}
func (m *SrmsL2AdjTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsL2AdjTB.Merge(m, src)
}
func (m *SrmsL2AdjTB) XXX_Size() int {
	return xxx_messageInfo_SrmsL2AdjTB.Size(m)
}
func (m *SrmsL2AdjTB) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsL2AdjTB.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsL2AdjTB proto.InternalMessageInfo

func (m *SrmsL2AdjTB) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *SrmsL2AdjTB) GetSidValueXr() uint32 {
	if m != nil {
		return m.SidValueXr
	}
	return 0
}

func (m *SrmsL2AdjTB) GetSidTypeXr() uint32 {
	if m != nil {
		return m.SidTypeXr
	}
	return 0
}

func (m *SrmsL2AdjTB) GetAddressFamily() uint32 {
	if m != nil {
		return m.AddressFamily
	}
	return 0
}

func (m *SrmsL2AdjTB) GetHasNexthop() bool {
	if m != nil {
		return m.HasNexthop
	}
	return false
}

func (m *SrmsL2AdjTB) GetInterfaceCount() int32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *SrmsL2AdjTB) GetInterfaceDeleteCount() int32 {
	if m != nil {
		return m.InterfaceDeleteCount
	}
	return 0
}

func (m *SrmsL2AdjTB) GetNexthopAddress() *AddrT {
	if m != nil {
		return m.NexthopAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*SrmsL2AdjTB_KEYS)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.adjacency_sid.l2_adjacency.interfaces.interface.address_family.ipv4.sid_record.srms_l2_adj_t_b_KEYS")
	proto.RegisterType((*AddrT)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.adjacency_sid.l2_adjacency.interfaces.interface.address_family.ipv4.sid_record.addr_t")
	proto.RegisterType((*SrmsL2AdjTB)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.adjacency_sid.l2_adjacency.interfaces.interface.address_family.ipv4.sid_record.srms_l2_adj_t_b")
}

func init() { proto.RegisterFile("srms_l2_adj_t_b.proto", fileDescriptor_179677fb0b32cf11) }

var fileDescriptor_179677fb0b32cf11 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x29, 0x2a, 0xcd, 0x84, 0x24, 0xc2, 0x2a, 0xc8, 0x08, 0x09, 0x56, 0x91, 0x10,
	0x7b, 0xf2, 0x21, 0x0d, 0x01, 0x6e, 0x20, 0xfe, 0x5c, 0x90, 0x7a, 0x58, 0x10, 0x2a, 0xa7, 0x91,
	0xbb, 0x9e, 0x34, 0xae, 0xb2, 0xeb, 0x95, 0xed, 0x54, 0xc9, 0x8d, 0xe7, 0xe1, 0x61, 0x78, 0x26,
	0x64, 0xaf, 0x93, 0x88, 0xaa, 0xf7, 0xde, 0xbc, 0xdf, 0x7c, 0xf3, 0xcd, 0x6f, 0x76, 0xd7, 0xf0,
	0xc4, 0xd9, 0xda, 0xe1, 0x6a, 0x8a, 0x52, 0x5d, 0xa3, 0xc7, 0x4b, 0xd1, 0x5a, 0xe3, 0x0d, 0xdb,
	0x54, 0xda, 0x55, 0x06, 0xb5, 0x71, 0xb8, 0xb1, 0xe8, 0xe8, 0xaa, 0xa6, 0xc6, 0xa3, 0x35, 0x6b,
	0xaf, 0x9b, 0x2b, 0xac, 0x1d, 0x9a, 0x96, 0xac, 0x08, 0xbd, 0x42, 0xaa, 0x6b, 0x59, 0x51, 0x53,
	0x6d, 0xd1, 0x69, 0x25, 0xba, 0xa4, 0x4e, 0x10, 0xba, 0xf1, 0x64, 0x17, 0xb2, 0x22, 0x77, 0x38,
	0x0a, 0xa9, 0x94, 0x25, 0xe7, 0x70, 0x21, 0x6b, 0xbd, 0xda, 0x0a, 0xdd, 0xde, 0xcc, 0x84, 0xd3,
	0x0a, 0x2d, 0x55, 0xc6, 0xaa, 0xc9, 0x1a, 0x4e, 0x6f, 0x21, 0xe1, 0xb7, 0x2f, 0xbf, 0xbe, 0xb3,
	0x57, 0x30, 0xda, 0x67, 0x60, 0x23, 0x6b, 0xe2, 0x59, 0x9e, 0x15, 0xfd, 0x72, 0xb8, 0x57, 0xcf,
	0x65, 0x4d, 0xec, 0x19, 0x9c, 0x84, 0x30, 0xbf, 0x6d, 0x89, 0xf7, 0xa2, 0xe1, 0xa1, 0xd3, 0xea,
	0xc7, 0xb6, 0x25, 0xf6, 0x1c, 0xfa, 0xa1, 0x74, 0x23, 0x57, 0x6b, 0xe2, 0x47, 0x79, 0x56, 0x0c,
	0xcb, 0xe0, 0xfd, 0x19, 0x9e, 0x27, 0x1f, 0xe0, 0x38, 0x80, 0xa1, 0x67, 0x23, 0xe8, 0xc9, 0x45,
	0x0a, 0xef, 0xc9, 0x05, 0x63, 0xf0, 0x20, 0x30, 0xa6, 0xb4, 0x78, 0x4e, 0xda, 0x3c, 0xa6, 0x74,
	0xda, 0x7c, 0xf2, 0xf7, 0x08, 0xc6, 0xb7, 0xc8, 0xef, 0x80, 0x9e, 0xde, 0x05, 0x9d, 0xc3, 0xa3,
	0x3d, 0x19, 0x6e, 0x2c, 0x3f, 0x8b, 0x70, 0xb0, 0x83, 0xbb, 0xb0, 0xec, 0x05, 0x0c, 0x76, 0x6b,
	0x05, 0xc3, 0x2c, 0x1a, 0xfa, 0x69, 0xb3, 0x0b, 0x1b, 0x06, 0xfd, 0xff, 0x5e, 0xf9, 0x9b, 0x68,
	0x19, 0x26, 0xf5, 0x6b, 0x14, 0xd9, 0x4b, 0x18, 0x2c, 0xa5, 0xc3, 0x86, 0x36, 0x7e, 0x69, 0x5a,
	0x3e, 0xcf, 0xb3, 0xe2, 0xa4, 0x84, 0xa5, 0x74, 0xe7, 0x9d, 0xc2, 0x5e, 0xc3, 0xf8, 0x00, 0x5c,
	0x99, 0x75, 0xe3, 0xf9, 0xdb, 0x3c, 0x2b, 0x1e, 0x97, 0x87, 0x3d, 0x3e, 0x05, 0x95, 0xcd, 0xe0,
	0xe9, 0xc1, 0xa8, 0x68, 0x45, 0x7e, 0xe7, 0x7f, 0x17, 0xfd, 0xa7, 0xfb, 0xea, 0xe7, 0x58, 0xec,
	0xba, 0xfe, 0x64, 0x30, 0x4e, 0xc3, 0x31, 0x91, 0xf1, 0xf7, 0x79, 0x56, 0x0c, 0xa6, 0xbf, 0x33,
	0x71, 0x5f, 0xbf, 0x9c, 0xe8, 0x3e, 0x7c, 0x39, 0x4a, 0x64, 0x1f, 0x3b, 0xf7, 0xe5, 0x71, 0xbc,
	0x0a, 0x67, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0x17, 0x8f, 0xec, 0x23, 0x03, 0x00, 0x00,
}
