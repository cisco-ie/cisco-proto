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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_vrfs_vrf_ifrs_summary

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
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

func (m *PimIdbSummBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimIdbSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.ifrs_summary.pim_idb_summ_bag_KEYS")
	proto.RegisterType((*PimIdbSummBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.ifrs_summary.pim_idb_summ_bag")
}

func init() { proto.RegisterFile("pim_idb_summ_bag.proto", fileDescriptor_2e335ab0c2fccc75) }

var fileDescriptor_2e335ab0c2fccc75 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0xc6, 0x8f, 0x80, 0x1f, 0x44, 0x94, 0xb3, 0x3b, 0xae, 0xf1, 0xaa, 0x08, 0x77,
	0x62, 0x2f, 0x62, 0x25, 0x58, 0x9c, 0x95, 0xd5, 0x90, 0xe4, 0x92, 0x63, 0xc0, 0x64, 0xc2, 0x24,
	0x1b, 0xdc, 0x7f, 0x2f, 0x09, 0x22, 0xdc, 0x36, 0x03, 0xf3, 0xce, 0xf3, 0x3e, 0x30, 0xe2, 0x2e,
	0x61, 0x00, 0xdc, 0x1b, 0xc8, 0x63, 0x08, 0x60, 0xf4, 0x41, 0x25, 0xa6, 0x42, 0xf2, 0xc5, 0x62,
	0xb6, 0x04, 0x48, 0x19, 0x7e, 0x18, 0x30, 0xd5, 0x27, 0x68, 0x24, 0x25, 0xc7, 0x0a, 0x53, 0x7d,
	0x6e, 0x9b, 0xca, 0x45, 0xc7, 0xbd, 0x99, 0x54, 0x65, 0x9f, 0xdb, 0x50, 0xe8, 0x39, 0x77, 0x93,
	0xe6, 0x69, 0xb5, 0x11, 0xb7, 0x73, 0x39, 0xbc, 0xbf, 0x7d, 0x7d, 0xca, 0x7b, 0x71, 0x56, 0xd9,
	0x43, 0xd4, 0xc1, 0x2d, 0x86, 0xe5, 0xb0, 0x3e, 0xdf, 0x9d, 0x56, 0xf6, 0x1f, 0x3a, 0xb8, 0xd5,
	0xb7, 0xb8, 0x9e, 0x77, 0xe4, 0x83, 0xb8, 0xc2, 0x58, 0x1c, 0x7b, 0x6d, 0x1d, 0x58, 0x1a, 0x63,
	0x59, 0x6c, 0x96, 0xc3, 0xfa, 0x62, 0x77, 0xf9, 0x1f, 0xbf, 0xb6, 0x54, 0x3e, 0x8a, 0x1b, 0x4b,
	0xd1, 0xe3, 0x61, 0x64, 0x5d, 0x90, 0xe2, 0x1f, 0xbc, 0xed, 0xb0, 0x3c, 0x3a, 0xf5, 0x82, 0x39,
	0xe9, 0xbf, 0x6e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x49, 0x99, 0xfc, 0xa8, 0x05, 0x01, 0x00,
	0x00,
}
