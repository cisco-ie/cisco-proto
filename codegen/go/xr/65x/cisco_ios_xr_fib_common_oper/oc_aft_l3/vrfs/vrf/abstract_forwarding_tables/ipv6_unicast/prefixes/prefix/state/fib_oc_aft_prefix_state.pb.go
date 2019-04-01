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
// source: fib_oc_aft_prefix_state.proto

package cisco_ios_xr_fib_common_oper_oc_aft_l3_vrfs_vrf_abstract_forwarding_tables_ipv6_unicast_prefixes_prefix_state

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

type FibOcAftPrefixState_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcAftPrefixState_KEYS) Reset()         { *m = FibOcAftPrefixState_KEYS{} }
func (m *FibOcAftPrefixState_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPrefixState_KEYS) ProtoMessage()    {}
func (*FibOcAftPrefixState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a171655de3f15d5, []int{0}
}

func (m *FibOcAftPrefixState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPrefixState_KEYS.Unmarshal(m, b)
}
func (m *FibOcAftPrefixState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPrefixState_KEYS.Marshal(b, m, deterministic)
}
func (m *FibOcAftPrefixState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPrefixState_KEYS.Merge(m, src)
}
func (m *FibOcAftPrefixState_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPrefixState_KEYS.Size(m)
}
func (m *FibOcAftPrefixState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPrefixState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPrefixState_KEYS proto.InternalMessageInfo

func (m *FibOcAftPrefixState_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibOcAftPrefixState_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type FibOcAftPrefixState struct {
	PrefixIndex          string   `protobuf:"bytes,50,opt,name=prefix_index,json=prefixIndex,proto3" json:"prefix_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcAftPrefixState) Reset()         { *m = FibOcAftPrefixState{} }
func (m *FibOcAftPrefixState) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPrefixState) ProtoMessage()    {}
func (*FibOcAftPrefixState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a171655de3f15d5, []int{1}
}

func (m *FibOcAftPrefixState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPrefixState.Unmarshal(m, b)
}
func (m *FibOcAftPrefixState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPrefixState.Marshal(b, m, deterministic)
}
func (m *FibOcAftPrefixState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPrefixState.Merge(m, src)
}
func (m *FibOcAftPrefixState) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPrefixState.Size(m)
}
func (m *FibOcAftPrefixState) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPrefixState.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPrefixState proto.InternalMessageInfo

func (m *FibOcAftPrefixState) GetPrefixIndex() string {
	if m != nil {
		return m.PrefixIndex
	}
	return ""
}

func init() {
	proto.RegisterType((*FibOcAftPrefixState_KEYS)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv6_unicast.prefixes.prefix.state.fib_oc_aft_prefix_state_KEYS")
	proto.RegisterType((*FibOcAftPrefixState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv6_unicast.prefixes.prefix.state.fib_oc_aft_prefix_state")
}

func init() { proto.RegisterFile("fib_oc_aft_prefix_state.proto", fileDescriptor_5a171655de3f15d5) }

var fileDescriptor_5a171655de3f15d5 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x50, 0xb1, 0x4a, 0x04, 0x41,
	0x0c, 0xe5, 0x2c, 0x4e, 0x1d, 0xad, 0xb6, 0xd0, 0x15, 0x14, 0xf4, 0x2a, 0xab, 0x29, 0x3c, 0xb0,
	0xb2, 0xb5, 0x10, 0x41, 0x50, 0x2b, 0xab, 0x90, 0x9d, 0x9b, 0x91, 0xc0, 0xcd, 0x64, 0x49, 0xe2,
	0xba, 0x9f, 0x7f, 0xec, 0xee, 0xb4, 0xd7, 0x84, 0xbc, 0x97, 0xf7, 0xf2, 0x42, 0xdc, 0x5d, 0xa2,
	0x0e, 0x38, 0x00, 0x26, 0x83, 0x5e, 0x62, 0xa2, 0x11, 0xd4, 0xd0, 0xa2, 0xef, 0x85, 0x8d, 0x9b,
	0x1c, 0x48, 0x03, 0x03, 0xb1, 0xc2, 0x28, 0x30, 0x69, 0x03, 0xe7, 0xcc, 0x05, 0xb8, 0x8f, 0xe2,
	0xab, 0x6f, 0xbf, 0xf5, 0x83, 0x24, 0x9d, 0x8a, 0xc7, 0x4e, 0x4d, 0x30, 0x18, 0x24, 0x96, 0x7f,
	0x94, 0x1d, 0x95, 0x5f, 0x30, 0xec, 0xf6, 0x51, 0x3d, 0xf5, 0xc3, 0x33, 0xfc, 0x15, 0x0a, 0xa8,
	0xe6, 0x97, 0xac, 0xa8, 0xb5, 0xf1, 0x73, 0xe8, 0xe6, 0xd3, 0xdd, 0x1e, 0xb9, 0x07, 0xde, 0x5f,
	0x7f, 0xbe, 0x9b, 0x1b, 0x77, 0x36, 0x48, 0x82, 0x82, 0x39, 0xb6, 0xab, 0xfb, 0xd5, 0xe3, 0xf9,
	0xd7, 0xe9, 0x20, 0xe9, 0x03, 0x73, 0x6c, 0xae, 0xdc, 0x7a, 0xd1, 0xb7, 0x27, 0xf3, 0xa0, 0xa2,
	0xcd, 0x8b, 0xbb, 0x3e, 0xb2, 0xb2, 0x79, 0x70, 0x97, 0x15, 0x53, 0xd9, 0xc5, 0xb1, 0x7d, 0x9a,
	0x8d, 0x17, 0x0b, 0xf7, 0x36, 0x51, 0xdd, 0x7a, 0x7e, 0xc3, 0xf6, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0x72, 0xfa, 0x28, 0x27, 0x01, 0x00, 0x00,
}
