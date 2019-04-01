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
// source: pim_idb_summ_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_ifrs_summary

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

type PimIdbSummBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimIdbSummBag_KEYS) Reset()         { *m = PimIdbSummBag_KEYS{} }
func (m *PimIdbSummBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimIdbSummBag_KEYS) ProtoMessage()    {}
func (*PimIdbSummBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e335ab0c2fccc75, []int{0}
}

func (m *PimIdbSummBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIdbSummBag_KEYS.Unmarshal(m, b)
}
func (m *PimIdbSummBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIdbSummBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimIdbSummBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIdbSummBag_KEYS.Merge(m, src)
}
func (m *PimIdbSummBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimIdbSummBag_KEYS.Size(m)
}
func (m *PimIdbSummBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIdbSummBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimIdbSummBag_KEYS proto.InternalMessageInfo

type PimIdbSummBag struct {
	InterfaceCount       uint32   `protobuf:"varint,50,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	ConfigurationCount   uint32   `protobuf:"varint,51,opt,name=configuration_count,json=configurationCount,proto3" json:"configuration_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimIdbSummBag) Reset()         { *m = PimIdbSummBag{} }
func (m *PimIdbSummBag) String() string { return proto.CompactTextString(m) }
func (*PimIdbSummBag) ProtoMessage()    {}
func (*PimIdbSummBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e335ab0c2fccc75, []int{1}
}

func (m *PimIdbSummBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIdbSummBag.Unmarshal(m, b)
}
func (m *PimIdbSummBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIdbSummBag.Marshal(b, m, deterministic)
}
func (m *PimIdbSummBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIdbSummBag.Merge(m, src)
}
func (m *PimIdbSummBag) XXX_Size() int {
	return xxx_messageInfo_PimIdbSummBag.Size(m)
}
func (m *PimIdbSummBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIdbSummBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimIdbSummBag proto.InternalMessageInfo

func (m *PimIdbSummBag) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *PimIdbSummBag) GetConfigurationCount() uint32 {
	if m != nil {
		return m.ConfigurationCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PimIdbSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.ifrs_summary.pim_idb_summ_bag_KEYS")
	proto.RegisterType((*PimIdbSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.ifrs_summary.pim_idb_summ_bag")
}

func init() { proto.RegisterFile("pim_idb_summ_bag.proto", fileDescriptor_2e335ab0c2fccc75) }

var fileDescriptor_2e335ab0c2fccc75 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3d, 0x6b, 0xc3, 0x30,
	0x10, 0x86, 0xe9, 0xd2, 0x41, 0xd0, 0x0f, 0x54, 0xfa, 0x31, 0x16, 0x2f, 0xed, 0xa4, 0x42, 0x1d,
	0xf2, 0x07, 0x42, 0xc8, 0x90, 0x2d, 0x99, 0x32, 0x1d, 0xb2, 0x2c, 0x99, 0x03, 0x5b, 0x27, 0xa4,
	0x93, 0x71, 0xfe, 0x7d, 0xb0, 0x08, 0x81, 0x78, 0xbc, 0xf7, 0x79, 0xee, 0x8e, 0x57, 0x7c, 0x04,
	0x1c, 0x00, 0xdb, 0x06, 0x52, 0x1e, 0x06, 0x68, 0x74, 0xa7, 0x42, 0x24, 0x26, 0xb9, 0x33, 0x98,
	0x0c, 0x01, 0x52, 0x82, 0x29, 0x02, 0x86, 0x71, 0x05, 0xb3, 0x49, 0xc1, 0x46, 0x85, 0x61, 0x5c,
	0xcf, 0x93, 0xd2, 0x86, 0x71, 0xb4, 0xaa, 0xb5, 0x4e, 0xe7, 0x9e, 0xc1, 0x90, 0x67, 0x3b, 0xb1,
	0x42, 0x17, 0x53, 0xb9, 0xa7, 0xe3, 0xb9, 0xfa, 0x14, 0xef, 0xcb, 0x17, 0xb0, 0xdf, 0x9e, 0x8e,
	0x55, 0x2f, 0x5e, 0x97, 0x40, 0xfe, 0x88, 0x17, 0xf4, 0x6c, 0xa3, 0xd3, 0xc6, 0x82, 0xa1, 0xec,
	0xf9, 0xeb, 0xff, 0xfb, 0xe1, 0xf7, 0xe9, 0xf0, 0x7c, 0x8b, 0x37, 0x73, 0x2a, 0xff, 0xc4, 0x9b,
	0x21, 0xef, 0xb0, 0xcb, 0x51, 0x33, 0x92, 0xbf, 0xca, 0x75, 0x91, 0xe5, 0x1d, 0x2a, 0x0b, 0xcd,
	0x63, 0xa9, 0x55, 0x5f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x03, 0x09, 0xad, 0xf0, 0x00, 0x00,
	0x00,
}
