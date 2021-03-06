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
// source: pim_jpstats_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_jp_statistics_jp_statistic

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

type PimJpstatsBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimJpstatsBag_KEYS) Reset()         { *m = PimJpstatsBag_KEYS{} }
func (m *PimJpstatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimJpstatsBag_KEYS) ProtoMessage()    {}
func (*PimJpstatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4f2f6deefad4d3, []int{0}
}

func (m *PimJpstatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimJpstatsBag_KEYS.Unmarshal(m, b)
}
func (m *PimJpstatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimJpstatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimJpstatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimJpstatsBag_KEYS.Merge(m, src)
}
func (m *PimJpstatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimJpstatsBag_KEYS.Size(m)
}
func (m *PimJpstatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimJpstatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimJpstatsBag_KEYS proto.InternalMessageInfo

func (m *PimJpstatsBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PimJpstatsBag struct {
	InterfaceNameXr      string   `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	Mtu                  uint32   `protobuf:"varint,51,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Transmitted          uint32   `protobuf:"varint,52,opt,name=transmitted,proto3" json:"transmitted,omitempty"`
	Transmitted100       uint32   `protobuf:"varint,53,opt,name=transmitted100,proto3" json:"transmitted100,omitempty"`
	Transmitted_1K       uint32   `protobuf:"varint,54,opt,name=transmitted_1k,json=transmitted1k,proto3" json:"transmitted_1k,omitempty"`
	Transmitted_10K      uint32   `protobuf:"varint,55,opt,name=transmitted_10k,json=transmitted10k,proto3" json:"transmitted_10k,omitempty"`
	Transmitted_50K      uint32   `protobuf:"varint,56,opt,name=transmitted_50k,json=transmitted50k,proto3" json:"transmitted_50k,omitempty"`
	Received             uint32   `protobuf:"varint,57,opt,name=received,proto3" json:"received,omitempty"`
	Received100          uint32   `protobuf:"varint,58,opt,name=received100,proto3" json:"received100,omitempty"`
	Received_1K          uint32   `protobuf:"varint,59,opt,name=received_1k,json=received1k,proto3" json:"received_1k,omitempty"`
	Received_10K         uint32   `protobuf:"varint,60,opt,name=received_10k,json=received10k,proto3" json:"received_10k,omitempty"`
	Received_50K         uint32   `protobuf:"varint,61,opt,name=received_50k,json=received50k,proto3" json:"received_50k,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimJpstatsBag) Reset()         { *m = PimJpstatsBag{} }
func (m *PimJpstatsBag) String() string { return proto.CompactTextString(m) }
func (*PimJpstatsBag) ProtoMessage()    {}
func (*PimJpstatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4f2f6deefad4d3, []int{1}
}

func (m *PimJpstatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimJpstatsBag.Unmarshal(m, b)
}
func (m *PimJpstatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimJpstatsBag.Marshal(b, m, deterministic)
}
func (m *PimJpstatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimJpstatsBag.Merge(m, src)
}
func (m *PimJpstatsBag) XXX_Size() int {
	return xxx_messageInfo_PimJpstatsBag.Size(m)
}
func (m *PimJpstatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimJpstatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimJpstatsBag proto.InternalMessageInfo

func (m *PimJpstatsBag) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *PimJpstatsBag) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *PimJpstatsBag) GetTransmitted() uint32 {
	if m != nil {
		return m.Transmitted
	}
	return 0
}

func (m *PimJpstatsBag) GetTransmitted100() uint32 {
	if m != nil {
		return m.Transmitted100
	}
	return 0
}

func (m *PimJpstatsBag) GetTransmitted_1K() uint32 {
	if m != nil {
		return m.Transmitted_1K
	}
	return 0
}

func (m *PimJpstatsBag) GetTransmitted_10K() uint32 {
	if m != nil {
		return m.Transmitted_10K
	}
	return 0
}

func (m *PimJpstatsBag) GetTransmitted_50K() uint32 {
	if m != nil {
		return m.Transmitted_50K
	}
	return 0
}

func (m *PimJpstatsBag) GetReceived() uint32 {
	if m != nil {
		return m.Received
	}
	return 0
}

func (m *PimJpstatsBag) GetReceived100() uint32 {
	if m != nil {
		return m.Received100
	}
	return 0
}

func (m *PimJpstatsBag) GetReceived_1K() uint32 {
	if m != nil {
		return m.Received_1K
	}
	return 0
}

func (m *PimJpstatsBag) GetReceived_10K() uint32 {
	if m != nil {
		return m.Received_10K
	}
	return 0
}

func (m *PimJpstatsBag) GetReceived_50K() uint32 {
	if m != nil {
		return m.Received_50K
	}
	return 0
}

func init() {
	proto.RegisterType((*PimJpstatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.jp_statistics.jp_statistic.pim_jpstats_bag_KEYS")
	proto.RegisterType((*PimJpstatsBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.jp_statistics.jp_statistic.pim_jpstats_bag")
}

func init() { proto.RegisterFile("pim_jpstats_bag.proto", fileDescriptor_db4f2f6deefad4d3) }

var fileDescriptor_db4f2f6deefad4d3 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0x29, 0x85, 0x8f, 0xcf, 0x68, 0x5b, 0x0d, 0x0a, 0xc1, 0x8d, 0xb5, 0xa0, 0x16, 0x17,
	0x43, 0x6a, 0x3b, 0xfe, 0xef, 0xd2, 0x95, 0x20, 0x58, 0x37, 0xba, 0x0a, 0xe9, 0x4c, 0x2a, 0x69,
	0x9c, 0x99, 0x90, 0xdc, 0x96, 0xfa, 0x64, 0xbe, 0x9e, 0x24, 0xe2, 0x74, 0x32, 0x75, 0x77, 0xef,
	0x99, 0xdf, 0xb9, 0x73, 0x0e, 0x04, 0x1d, 0x68, 0x99, 0xb1, 0xb9, 0xb6, 0xc0, 0xc1, 0xb2, 0x29,
	0x7f, 0x8f, 0xb4, 0x29, 0xa0, 0xc0, 0xcf, 0x89, 0xb4, 0x49, 0xc1, 0x64, 0x61, 0xd9, 0xca, 0x30,
	0xa9, 0x97, 0x23, 0xe6, 0xc0, 0x42, 0x0b, 0x13, 0x69, 0x99, 0x45, 0x16, 0x78, 0x9e, 0x4e, 0x3f,
	0xa3, 0x54, 0xcc, 0xf8, 0xe2, 0x03, 0x58, 0x52, 0xe4, 0x20, 0x56, 0x10, 0xcd, 0x35, 0x73, 0xa7,
	0xa4, 0x05, 0x99, 0xd8, 0x60, 0xeb, 0x8d, 0xd1, 0x7e, 0xed, 0x5f, 0xec, 0xf1, 0xe1, 0xed, 0x05,
	0x9f, 0xa0, 0xb6, 0xcc, 0x41, 0x98, 0x19, 0x4f, 0x04, 0xcb, 0x79, 0x26, 0x48, 0xa3, 0xdb, 0xe8,
	0x6f, 0x4d, 0x5a, 0xa5, 0xfa, 0xc4, 0x33, 0xd1, 0xfb, 0x6a, 0xa2, 0x4e, 0xcd, 0x8f, 0xcf, 0xd1,
	0x5e, 0x68, 0x65, 0x2b, 0x43, 0x2e, 0xbc, 0xbb, 0x13, 0xb8, 0x5f, 0x0d, 0xde, 0x45, 0xcd, 0x0c,
	0x16, 0x64, 0xd8, 0x6d, 0xf4, 0x5b, 0x13, 0x37, 0xe2, 0x2e, 0xda, 0x06, 0xc3, 0x73, 0x9b, 0x49,
	0x00, 0x91, 0x92, 0x91, 0xff, 0x52, 0x95, 0xf0, 0x29, 0x6a, 0x57, 0xd6, 0x01, 0xa5, 0x24, 0xf6,
	0x50, 0x4d, 0x75, 0x15, 0x2a, 0x0a, 0x1b, 0x28, 0x72, 0xe9, 0xb9, 0x56, 0x95, 0x53, 0xf8, 0x0c,
	0x75, 0x02, 0x8c, 0x2a, 0x72, 0xf5, 0xc7, 0xbd, 0x0d, 0x30, 0xa6, 0x8a, 0x5c, 0x6f, 0x80, 0x31,
	0x55, 0xf8, 0x10, 0xfd, 0x37, 0x22, 0x11, 0x72, 0x29, 0x52, 0x72, 0xe3, 0x89, 0x72, 0x77, 0xf5,
	0x7e, 0x67, 0x97, 0xfc, 0xf6, 0xa7, 0x5e, 0x45, 0xc2, 0x47, 0x6b, 0xc2, 0x65, 0xbe, 0xf3, 0x04,
	0x2a, 0x09, 0x85, 0x8f, 0xd1, 0xce, 0x1a, 0xa0, 0x8a, 0xdc, 0xd7, 0x6f, 0x84, 0x88, 0xcb, 0x39,
	0x0e, 0x91, 0x98, 0xaa, 0xe9, 0x3f, 0xff, 0xa4, 0x86, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x31, 0xaf, 0x0c, 0x6b, 0x02, 0x00, 0x00,
}
