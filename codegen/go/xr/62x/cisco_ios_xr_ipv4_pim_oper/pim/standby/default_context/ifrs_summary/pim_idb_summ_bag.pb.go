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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_ifrs_summary

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
	proto.RegisterType((*PimIdbSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.ifrs_summary.pim_idb_summ_bag_KEYS")
	proto.RegisterType((*PimIdbSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.ifrs_summary.pim_idb_summ_bag")
}

func init() { proto.RegisterFile("pim_idb_summ_bag.proto", fileDescriptor_2e335ab0c2fccc75) }

var fileDescriptor_2e335ab0c2fccc75 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xbd, 0x6a, 0xc4, 0x30,
	0x10, 0x06, 0x49, 0x93, 0x42, 0x90, 0x1f, 0x14, 0xf2, 0x53, 0x06, 0x37, 0x49, 0xa5, 0x40, 0x9c,
	0x37, 0x30, 0xa9, 0xd2, 0x25, 0x55, 0xaa, 0x45, 0x92, 0x25, 0xb3, 0x60, 0x69, 0x85, 0xb4, 0x3a,
	0xec, 0xb7, 0x3f, 0x2c, 0x8e, 0x83, 0x73, 0xfb, 0xcd, 0xec, 0xc0, 0x8a, 0xa7, 0x84, 0x01, 0x70,
	0x34, 0x50, 0x6a, 0x08, 0x60, 0xf4, 0xa4, 0x52, 0x26, 0x26, 0x39, 0x58, 0x2c, 0x96, 0x00, 0xa9,
	0xc0, 0x92, 0x01, 0xd3, 0xe1, 0x0b, 0x36, 0x93, 0x92, 0xcb, 0x2a, 0x61, 0x50, 0x85, 0x75, 0x1c,
	0xcd, 0xaa, 0x46, 0xe7, 0x75, 0x9d, 0x19, 0x2c, 0x45, 0x76, 0x0b, 0x2b, 0xf4, 0xb9, 0xb4, 0x96,
	0xce, 0x6b, 0xf7, 0x2c, 0x1e, 0xf7, 0x79, 0xf8, 0xf9, 0xfe, 0xff, 0xeb, 0x66, 0x71, 0xbf, 0x07,
	0xf2, 0x4d, 0xdc, 0x61, 0x64, 0x97, 0xbd, 0xb6, 0x0e, 0x2c, 0xd5, 0xc8, 0x2f, 0x9f, 0xaf, 0x57,
	0xef, 0x37, 0xbf, 0xb7, 0xe7, 0x79, 0xd8, 0x56, 0xf9, 0x21, 0x1e, 0x2c, 0x45, 0x8f, 0x53, 0xcd,
	0x9a, 0x91, 0xe2, 0x49, 0xee, 0x9b, 0x2c, 0x2f, 0x50, 0x3b, 0x30, 0xd7, 0xed, 0xa5, 0xfe, 0x18,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x39, 0xcc, 0xc5, 0xec, 0x00, 0x00, 0x00,
}