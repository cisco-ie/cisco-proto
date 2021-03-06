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
// source: srms_mi_t_b.proto

package cisco_ios_xr_segment_routing_ms_oper_srms_mapping_mapping_ipv6_mapping_mi

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

type SrmsMiTB_KEYS struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Prefix               uint32   `protobuf:"varint,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB_KEYS) Reset()         { *m = SrmsMiTB_KEYS{} }
func (m *SrmsMiTB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB_KEYS) ProtoMessage()    {}
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{0}
}

func (m *SrmsMiTB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB_KEYS.Unmarshal(m, b)
}
func (m *SrmsMiTB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB_KEYS.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB_KEYS.Merge(m, src)
}
func (m *SrmsMiTB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB_KEYS.Size(m)
}
func (m *SrmsMiTB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB_KEYS proto.InternalMessageInfo

func (m *SrmsMiTB_KEYS) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *SrmsMiTB_KEYS) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

type Addr struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{1}
}

func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *Addr) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Addr) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type SrmsMiTB struct {
	Src                  string   `protobuf:"bytes,50,opt,name=src,proto3" json:"src,omitempty"`
	Router               string   `protobuf:"bytes,51,opt,name=router,proto3" json:"router,omitempty"`
	Area                 string   `protobuf:"bytes,52,opt,name=area,proto3" json:"area,omitempty"`
	Addr                 *Addr    `protobuf:"bytes,53,opt,name=addr,proto3" json:"addr,omitempty"`
	PrefixXr             uint32   `protobuf:"varint,54,opt,name=prefix_xr,json=prefixXr,proto3" json:"prefix_xr,omitempty"`
	SidStart             uint32   `protobuf:"varint,55,opt,name=sid_start,json=sidStart,proto3" json:"sid_start,omitempty"`
	SidCount             uint32   `protobuf:"varint,56,opt,name=sid_count,json=sidCount,proto3" json:"sid_count,omitempty"`
	LastPrefix           string   `protobuf:"bytes,57,opt,name=last_prefix,json=lastPrefix,proto3" json:"last_prefix,omitempty"`
	LastSidIndex         uint32   `protobuf:"varint,58,opt,name=last_sid_index,json=lastSidIndex,proto3" json:"last_sid_index,omitempty"`
	FlagAttached         string   `protobuf:"bytes,59,opt,name=flag_attached,json=flagAttached,proto3" json:"flag_attached,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB) Reset()         { *m = SrmsMiTB{} }
func (m *SrmsMiTB) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB) ProtoMessage()    {}
func (*SrmsMiTB) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{2}
}

func (m *SrmsMiTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB.Unmarshal(m, b)
}
func (m *SrmsMiTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB.Merge(m, src)
}
func (m *SrmsMiTB) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB.Size(m)
}
func (m *SrmsMiTB) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB proto.InternalMessageInfo

func (m *SrmsMiTB) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *SrmsMiTB) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *SrmsMiTB) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *SrmsMiTB) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *SrmsMiTB) GetPrefixXr() uint32 {
	if m != nil {
		return m.PrefixXr
	}
	return 0
}

func (m *SrmsMiTB) GetSidStart() uint32 {
	if m != nil {
		return m.SidStart
	}
	return 0
}

func (m *SrmsMiTB) GetSidCount() uint32 {
	if m != nil {
		return m.SidCount
	}
	return 0
}

func (m *SrmsMiTB) GetLastPrefix() string {
	if m != nil {
		return m.LastPrefix
	}
	return ""
}

func (m *SrmsMiTB) GetLastSidIndex() uint32 {
	if m != nil {
		return m.LastSidIndex
	}
	return 0
}

func (m *SrmsMiTB) GetFlagAttached() string {
	if m != nil {
		return m.FlagAttached
	}
	return ""
}

func init() {
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.mapping.mapping_ipv6.mapping_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.mapping.mapping_ipv6.mapping_mi.addr")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_segment_routing_ms_oper.srms.mapping.mapping_ipv6.mapping_mi.srms_mi_t_b")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor_bb21d79d36617d23) }

var fileDescriptor_bb21d79d36617d23 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x6c, 0x4c, 0xbd, 0xfe, 0x83, 0xbb, 0x87, 0xb2, 0xd0, 0x43, 0x8d, 0xdb, 0x83,
	0x4f, 0x3a, 0xd8, 0xae, 0xda, 0xba, 0x50, 0x28, 0x21, 0x07, 0x93, 0x43, 0x82, 0x7c, 0x49, 0x4e,
	0xc3, 0x5a, 0x5a, 0x29, 0x03, 0x96, 0xb4, 0xec, 0xae, 0x8d, 0xde, 0x3b, 0x2f, 0x10, 0x66, 0x25,
	0x1b, 0x3f, 0x40, 0x4e, 0x9a, 0xef, 0x37, 0x9f, 0xbe, 0xd1, 0x0c, 0x62, 0x9f, 0xad, 0x29, 0x2d,
	0x94, 0x08, 0x0e, 0x0e, 0x91, 0x36, 0xb5, 0xab, 0xf9, 0x2e, 0x45, 0x9b, 0xd6, 0x80, 0xb5, 0x85,
	0xc6, 0x80, 0x55, 0x45, 0xa9, 0x2a, 0x07, 0xa6, 0x3e, 0x39, 0xac, 0x0a, 0x28, 0x2d, 0xd4, 0x5a,
	0x99, 0x88, 0xde, 0x8b, 0x4a, 0xa9, 0x35, 0x56, 0xc5, 0xe5, 0x09, 0xa8, 0xcf, 0xf1, 0x55, 0x94,
	0xb8, 0xd8, 0xb2, 0xd9, 0x4d, 0x3e, 0x3c, 0xdc, 0xbf, 0xec, 0xf9, 0x94, 0x85, 0xa8, 0x45, 0x30,
	0x0f, 0x96, 0xc3, 0x24, 0x44, 0xcd, 0xbf, 0xb0, 0x81, 0x36, 0x2a, 0xc7, 0x46, 0x84, 0xf3, 0x60,
	0x39, 0x49, 0x3a, 0xb5, 0xf8, 0xc7, 0xfa, 0x32, 0xcb, 0x0c, 0xf9, 0x65, 0x7e, 0xf1, 0xcb, 0x9c,
	0x73, 0xd6, 0x47, 0x7d, 0xde, 0x78, 0xf7, 0x30, 0xf1, 0x75, 0xc7, 0x62, 0xd1, 0xbb, 0xb2, 0x78,
	0xf1, 0x16, 0xb2, 0xd1, 0xcd, 0x70, 0x3e, 0x63, 0x3d, 0x6b, 0x52, 0xb1, 0xf2, 0x16, 0x2a, 0x69,
	0x32, 0x6d, 0xa5, 0x8c, 0x58, 0x7b, 0xd8, 0x29, 0x4a, 0x93, 0x46, 0x49, 0xb1, 0x69, 0xd3, 0xa8,
	0xe6, 0x69, 0xfb, 0x35, 0xe2, 0xe7, 0x3c, 0x58, 0x8e, 0x56, 0x8f, 0xd1, 0x87, 0xdd, 0x28, 0xa2,
	0xd8, 0xa4, 0x5d, 0xf5, 0x2b, 0x1b, 0xb6, 0xcb, 0x43, 0x63, 0x44, 0xec, 0xaf, 0xf1, 0xa9, 0x05,
	0xcf, 0xbe, 0x69, 0x31, 0x03, 0xeb, 0xa4, 0x71, 0xe2, 0x57, 0xdb, 0xb4, 0x98, 0xed, 0x49, 0x5f,
	0x9a, 0x69, 0x7d, 0xaa, 0x9c, 0xf8, 0x7d, 0x6d, 0xde, 0x91, 0xe6, 0xdf, 0xd8, 0xe8, 0x28, 0xad,
	0x83, 0xee, 0xcc, 0x7f, 0xfc, 0x5a, 0x8c, 0xd0, 0x93, 0x27, 0xfc, 0x07, 0x9b, 0x7a, 0x03, 0x45,
	0x60, 0x95, 0xa9, 0x46, 0x6c, 0x7d, 0xc4, 0x98, 0xe8, 0x1e, 0xb3, 0x1d, 0x31, 0xfe, 0x9d, 0x4d,
	0xf2, 0xa3, 0x2c, 0x40, 0x3a, 0x27, 0xd3, 0x57, 0x95, 0x89, 0xbf, 0x3e, 0x68, 0x4c, 0xf0, 0x7f,
	0xc7, 0x0e, 0x03, 0xff, 0x0f, 0xad, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x54, 0x6a, 0x98,
	0x58, 0x02, 0x00, 0x00,
}
