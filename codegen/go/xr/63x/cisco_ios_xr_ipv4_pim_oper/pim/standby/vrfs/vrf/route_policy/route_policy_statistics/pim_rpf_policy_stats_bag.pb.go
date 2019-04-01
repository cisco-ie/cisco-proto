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
// source: pim_rpf_policy_stats_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_route_policy_route_policy_statistics

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

type PimRpfPolicyStatsBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfPolicyStatsBag_KEYS) Reset()         { *m = PimRpfPolicyStatsBag_KEYS{} }
func (m *PimRpfPolicyStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfPolicyStatsBag_KEYS) ProtoMessage()    {}
func (*PimRpfPolicyStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e3c28953e18359f, []int{0}
}

func (m *PimRpfPolicyStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfPolicyStatsBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfPolicyStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfPolicyStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfPolicyStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfPolicyStatsBag_KEYS.Merge(m, src)
}
func (m *PimRpfPolicyStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfPolicyStatsBag_KEYS.Size(m)
}
func (m *PimRpfPolicyStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfPolicyStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfPolicyStatsBag_KEYS proto.InternalMessageInfo

func (m *PimRpfPolicyStatsBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type PimRpfPolicyStatsBag struct {
	PolicyName           string   `protobuf:"bytes,50,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Requests             uint32   `protobuf:"varint,51,opt,name=requests,proto3" json:"requests,omitempty"`
	Pass                 uint32   `protobuf:"varint,52,opt,name=pass,proto3" json:"pass,omitempty"`
	Drop                 uint32   `protobuf:"varint,53,opt,name=drop,proto3" json:"drop,omitempty"`
	DefaultTable         uint32   `protobuf:"varint,54,opt,name=default_table,json=defaultTable,proto3" json:"default_table,omitempty"`
	AnyTable             uint32   `protobuf:"varint,55,opt,name=any_table,json=anyTable,proto3" json:"any_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfPolicyStatsBag) Reset()         { *m = PimRpfPolicyStatsBag{} }
func (m *PimRpfPolicyStatsBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfPolicyStatsBag) ProtoMessage()    {}
func (*PimRpfPolicyStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e3c28953e18359f, []int{1}
}

func (m *PimRpfPolicyStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfPolicyStatsBag.Unmarshal(m, b)
}
func (m *PimRpfPolicyStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfPolicyStatsBag.Marshal(b, m, deterministic)
}
func (m *PimRpfPolicyStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfPolicyStatsBag.Merge(m, src)
}
func (m *PimRpfPolicyStatsBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfPolicyStatsBag.Size(m)
}
func (m *PimRpfPolicyStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfPolicyStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfPolicyStatsBag proto.InternalMessageInfo

func (m *PimRpfPolicyStatsBag) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PimRpfPolicyStatsBag) GetRequests() uint32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *PimRpfPolicyStatsBag) GetPass() uint32 {
	if m != nil {
		return m.Pass
	}
	return 0
}

func (m *PimRpfPolicyStatsBag) GetDrop() uint32 {
	if m != nil {
		return m.Drop
	}
	return 0
}

func (m *PimRpfPolicyStatsBag) GetDefaultTable() uint32 {
	if m != nil {
		return m.DefaultTable
	}
	return 0
}

func (m *PimRpfPolicyStatsBag) GetAnyTable() uint32 {
	if m != nil {
		return m.AnyTable
	}
	return 0
}

func init() {
	proto.RegisterType((*PimRpfPolicyStatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.route_policy.route_policy_statistics.pim_rpf_policy_stats_bag_KEYS")
	proto.RegisterType((*PimRpfPolicyStatsBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.route_policy.route_policy_statistics.pim_rpf_policy_stats_bag")
}

func init() { proto.RegisterFile("pim_rpf_policy_stats_bag.proto", fileDescriptor_6e3c28953e18359f) }

var fileDescriptor_6e3c28953e18359f = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x09, 0x41, 0x6b, 0xe8, 0xe2, 0xc9, 0x80, 0x80, 0xaa, 0x2c, 0x9d, 0x32, 0xd0,
	0x02, 0x12, 0x3b, 0x13, 0x12, 0x43, 0xe9, 0xc2, 0x74, 0xba, 0x24, 0x36, 0xb2, 0x94, 0xc4, 0xe6,
	0xce, 0x89, 0xc8, 0xe3, 0xf1, 0x66, 0xc8, 0x4e, 0x84, 0xd4, 0xa1, 0x8b, 0x75, 0xf7, 0x9d, 0xbf,
	0xfb, 0xa5, 0x13, 0xb7, 0xde, 0x36, 0x40, 0xde, 0x80, 0x77, 0xb5, 0x2d, 0x07, 0xe0, 0x80, 0x81,
	0xa1, 0xc0, 0xaf, 0xdc, 0x93, 0x0b, 0x4e, 0xee, 0x4b, 0xcb, 0xa5, 0x03, 0xeb, 0x18, 0x7e, 0x08,
	0xac, 0xef, 0xb7, 0x10, 0x0d, 0xe7, 0x35, 0xe5, 0xde, 0x36, 0x39, 0x07, 0x6c, 0xab, 0x62, 0xc8,
	0x7b, 0x32, 0x1c, 0x9f, 0x9c, 0x5c, 0x17, 0xf4, 0xb4, 0xed, 0xa0, 0x49, 0xab, 0x2d, 0x07, 0x5b,
	0xf2, 0xea, 0x45, 0xdc, 0x1c, 0xcb, 0x85, 0xb7, 0xd7, 0xcf, 0x0f, 0x79, 0x29, 0x66, 0x3d, 0x19,
	0x68, 0xb1, 0xd1, 0x2a, 0x5b, 0x66, 0xeb, 0xf9, 0xee, 0xac, 0x27, 0xf3, 0x8e, 0x8d, 0x5e, 0xfd,
	0x66, 0x42, 0x1d, 0x93, 0xe5, 0x9d, 0x38, 0x9f, 0x58, 0x52, 0x1f, 0x92, 0x2a, 0x46, 0x14, 0x6d,
	0x79, 0x25, 0x66, 0xa4, 0xbf, 0x3b, 0xcd, 0x81, 0xd5, 0x66, 0x99, 0xad, 0x17, 0xbb, 0xff, 0x5e,
	0x4a, 0x71, 0xe2, 0x91, 0x59, 0x6d, 0x13, 0x4f, 0x75, 0x64, 0x15, 0x39, 0xaf, 0x1e, 0x47, 0x16,
	0x6b, 0x79, 0x2f, 0x16, 0x95, 0x36, 0xd8, 0xd5, 0x01, 0x02, 0x16, 0xb5, 0x56, 0x4f, 0x69, 0x78,
	0x31, 0xc1, 0x7d, 0x64, 0xf2, 0x5a, 0xcc, 0xb1, 0x1d, 0xa6, 0x0f, 0xcf, 0x63, 0x12, 0xb6, 0x43,
	0x1a, 0x16, 0xa7, 0xe9, 0xb8, 0x9b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x49, 0x36, 0x51,
	0x7e, 0x01, 0x00, 0x00,
}
