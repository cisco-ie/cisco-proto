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
// source: bgp_rpki_summary_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_rpki_summary

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

type BgpRpkiSummaryBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRpkiSummaryBag_KEYS) Reset()         { *m = BgpRpkiSummaryBag_KEYS{} }
func (m *BgpRpkiSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiSummaryBag_KEYS) ProtoMessage()    {}
func (*BgpRpkiSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_392a39b2fb54b382, []int{0}
}

func (m *BgpRpkiSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRpkiSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRpkiSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiSummaryBag_KEYS.Merge(m, src)
}
func (m *BgpRpkiSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiSummaryBag_KEYS.Size(m)
}
func (m *BgpRpkiSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiSummaryBag_KEYS proto.InternalMessageInfo

func (m *BgpRpkiSummaryBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpRpkiSummaryBag struct {
	Servers              uint32   `protobuf:"varint,50,opt,name=servers,proto3" json:"servers,omitempty"`
	Ipv4RoaNets          uint32   `protobuf:"varint,51,opt,name=ipv4roa_nets,json=ipv4roaNets,proto3" json:"ipv4roa_nets,omitempty"`
	Ipv4RoaPaths         uint32   `protobuf:"varint,52,opt,name=ipv4roa_paths,json=ipv4roaPaths,proto3" json:"ipv4roa_paths,omitempty"`
	Ipv6RoaNets          uint32   `protobuf:"varint,53,opt,name=ipv6roa_nets,json=ipv6roaNets,proto3" json:"ipv6roa_nets,omitempty"`
	Ipv6RoaPaths         uint32   `protobuf:"varint,54,opt,name=ipv6roa_paths,json=ipv6roaPaths,proto3" json:"ipv6roa_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRpkiSummaryBag) Reset()         { *m = BgpRpkiSummaryBag{} }
func (m *BgpRpkiSummaryBag) String() string { return proto.CompactTextString(m) }
func (*BgpRpkiSummaryBag) ProtoMessage()    {}
func (*BgpRpkiSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_392a39b2fb54b382, []int{1}
}

func (m *BgpRpkiSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpkiSummaryBag.Unmarshal(m, b)
}
func (m *BgpRpkiSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpkiSummaryBag.Marshal(b, m, deterministic)
}
func (m *BgpRpkiSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpkiSummaryBag.Merge(m, src)
}
func (m *BgpRpkiSummaryBag) XXX_Size() int {
	return xxx_messageInfo_BgpRpkiSummaryBag.Size(m)
}
func (m *BgpRpkiSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpkiSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpkiSummaryBag proto.InternalMessageInfo

func (m *BgpRpkiSummaryBag) GetServers() uint32 {
	if m != nil {
		return m.Servers
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv4RoaNets() uint32 {
	if m != nil {
		return m.Ipv4RoaNets
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv4RoaPaths() uint32 {
	if m != nil {
		return m.Ipv4RoaPaths
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv6RoaNets() uint32 {
	if m != nil {
		return m.Ipv6RoaNets
	}
	return 0
}

func (m *BgpRpkiSummaryBag) GetIpv6RoaPaths() uint32 {
	if m != nil {
		return m.Ipv6RoaPaths
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpRpkiSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_summary.bgp_rpki_summary_bag_KEYS")
	proto.RegisterType((*BgpRpkiSummaryBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.rpki_summary.bgp_rpki_summary_bag")
}

func init() { proto.RegisterFile("bgp_rpki_summary_bag.proto", fileDescriptor_392a39b2fb54b382) }

var fileDescriptor_392a39b2fb54b382 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x05, 0x84, 0x69, 0x97, 0x88, 0xc1, 0x30, 0x95, 0xb2, 0x74, 0xca, 0x40, 0x8b,
	0x67, 0x16, 0x26, 0xa4, 0x82, 0xca, 0xc4, 0x74, 0xb2, 0xc3, 0x29, 0x58, 0x28, 0xb6, 0x75, 0x67,
	0x2a, 0xfa, 0xef, 0xf8, 0x69, 0xc8, 0x56, 0x6c, 0x40, 0xea, 0xe4, 0xf3, 0xf3, 0x7b, 0xdf, 0xb3,
	0x4e, 0x5c, 0x99, 0x21, 0x00, 0x85, 0x0f, 0x0b, 0xfc, 0x39, 0x8e, 0x9a, 0x0e, 0x60, 0xf4, 0xd0,
	0x05, 0xf2, 0xd1, 0xb7, 0x4f, 0xbd, 0xe5, 0xde, 0x83, 0xf5, 0x0c, 0x5f, 0x04, 0x36, 0xec, 0x37,
	0x90, 0xdc, 0x3e, 0x20, 0x75, 0x66, 0x08, 0x9d, 0x75, 0x1c, 0xb5, 0xeb, 0x91, 0xeb, 0x54, 0x07,
	0x48, 0xc7, 0x9b, 0x39, 0x74, 0x7f, 0xd1, 0xcb, 0x7b, 0x71, 0x79, 0xac, 0x0e, 0x1e, 0x1f, 0x5e,
	0x5f, 0xda, 0x1b, 0x31, 0xaf, 0x69, 0xa7, 0x47, 0x94, 0xcd, 0xa2, 0x59, 0x9d, 0xed, 0x66, 0x45,
	0xdc, 0xea, 0x11, 0x97, 0xdf, 0x8d, 0xb8, 0x38, 0x86, 0x68, 0xa5, 0x38, 0x65, 0xa4, 0x3d, 0x12,
	0xcb, 0xdb, 0x45, 0xb3, 0x9a, 0xef, 0xca, 0xb5, 0xbd, 0x16, 0xb3, 0xf4, 0x75, 0xf2, 0x1a, 0x1c,
	0x46, 0x96, 0xeb, 0xfc, 0x7c, 0x3e, 0x69, 0x5b, 0x8c, 0x9c, 0xab, 0x27, 0x4b, 0xd0, 0xf1, 0x9d,
	0xe5, 0x26, 0x7b, 0x4a, 0xee, 0x39, 0x69, 0x13, 0x47, 0x55, 0xce, 0x5d, 0xe5, 0xa8, 0xff, 0x1c,
	0xf5, 0xcb, 0x51, 0x95, 0xa3, 0x0a, 0xc7, 0x9c, 0xe4, 0xe5, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xac, 0x9a, 0x8a, 0xab, 0x7a, 0x01, 0x00, 0x00,
}
