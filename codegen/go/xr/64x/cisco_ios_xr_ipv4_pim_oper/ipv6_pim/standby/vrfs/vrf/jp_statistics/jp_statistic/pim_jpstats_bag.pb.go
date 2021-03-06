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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_vrfs_vrf_jp_statistics_jp_statistic

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *PimJpstatsBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimJpstatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.jp_statistics.jp_statistic.pim_jpstats_bag_KEYS")
	proto.RegisterType((*PimJpstatsBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.jp_statistics.jp_statistic.pim_jpstats_bag")
}

func init() { proto.RegisterFile("pim_jpstats_bag.proto", fileDescriptor_db4f2f6deefad4d3) }

var fileDescriptor_db4f2f6deefad4d3 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0x5d, 0x4b, 0xf3, 0x30,
	0x14, 0x07, 0x70, 0xf6, 0x0c, 0x1e, 0xe7, 0xd1, 0x6d, 0x1a, 0x14, 0xa2, 0x37, 0xce, 0x81, 0x3a,
	0xbc, 0x28, 0x99, 0x5b, 0xe7, 0xfb, 0xa5, 0x57, 0x82, 0xc2, 0xbc, 0x99, 0x57, 0x21, 0xeb, 0x52,
	0xc9, 0x42, 0xdb, 0x90, 0xc4, 0x32, 0x3f, 0x99, 0x5f, 0x4f, 0x92, 0xe1, 0xd6, 0x76, 0xde, 0x94,
	0x9c, 0x7f, 0x7f, 0x39, 0x3d, 0x07, 0x0a, 0x87, 0x4a, 0x24, 0x74, 0xae, 0x8c, 0x65, 0xd6, 0xd0,
	0x29, 0xfb, 0x08, 0x94, 0xce, 0x6c, 0x86, 0x5e, 0x23, 0x61, 0xa2, 0x8c, 0x8a, 0xcc, 0xd0, 0x85,
	0xa6, 0x42, 0xe5, 0x43, 0xea, 0x60, 0xa6, 0xb8, 0x0e, 0x84, 0xca, 0x47, 0xae, 0x0a, 0x8c, 0x65,
	0xe9, 0x6c, 0xfa, 0x15, 0xe4, 0x3a, 0x36, 0xee, 0x11, 0xcc, 0x15, 0x75, 0x8d, 0x84, 0xb1, 0x22,
	0x32, 0xa5, 0xaa, 0x3b, 0x81, 0x83, 0xca, 0x97, 0xe8, 0xf3, 0xd3, 0xfb, 0x1b, 0x3a, 0x82, 0x46,
	0xae, 0x63, 0x9a, 0xb2, 0x84, 0xe3, 0x5a, 0xa7, 0xd6, 0xdb, 0x1e, 0x6f, 0xe5, 0x3a, 0x7e, 0x61,
	0x09, 0x47, 0x67, 0xd0, 0x12, 0xa9, 0xe5, 0x3a, 0x66, 0x11, 0x5f, 0x82, 0x7f, 0x1e, 0x34, 0x57,
	0xa9, 0x63, 0xdd, 0xef, 0x3a, 0xb4, 0x2b, 0xad, 0xd1, 0x25, 0xec, 0x97, 0xaf, 0xd2, 0x85, 0xc6,
	0x57, 0xfe, 0x76, 0xbb, 0x74, 0x7b, 0xa2, 0xd1, 0x1e, 0xd4, 0x13, 0xfb, 0x89, 0x07, 0x9d, 0x5a,
	0xaf, 0x39, 0x76, 0x47, 0xd4, 0x81, 0x1d, 0xab, 0x59, 0x6a, 0x12, 0x61, 0x2d, 0x9f, 0xe1, 0xa1,
	0x7f, 0x53, 0x8c, 0xd0, 0x39, 0xb4, 0x0a, 0x65, 0x9f, 0x10, 0x1c, 0x7a, 0x54, 0x49, 0xdd, 0x0a,
	0x85, 0x84, 0xf6, 0x25, 0x1e, 0x79, 0xd7, 0x2c, 0x3a, 0x89, 0x2e, 0xa0, 0x5d, 0x62, 0x44, 0xe2,
	0xeb, 0x3f, 0xfa, 0x6d, 0xc0, 0x90, 0x48, 0x7c, 0xb3, 0x01, 0x43, 0x22, 0xd1, 0x31, 0x34, 0x34,
	0x8f, 0xb8, 0xc8, 0xf9, 0x0c, 0xdf, 0x7a, 0xb1, 0xaa, 0xdd, 0x7a, 0xbf, 0x67, 0x37, 0xf9, 0xdd,
	0x72, 0xbd, 0x42, 0x84, 0x4e, 0xd6, 0xc2, 0xcd, 0x7c, 0xef, 0x05, 0xac, 0x84, 0x44, 0xa7, 0xb0,
	0xbb, 0x06, 0x44, 0xe2, 0x87, 0x6a, 0x8f, 0x32, 0x71, 0x73, 0x3e, 0x96, 0x49, 0x48, 0xe4, 0xf4,
	0xbf, 0xff, 0xd7, 0x06, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x4c, 0xf9, 0x47, 0x84, 0x02,
	0x00, 0x00,
}
