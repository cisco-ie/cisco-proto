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
// source: vrf_info.proto

package cisco_ios_xr_lib_mpp_oper_management_plane_protection_outband_vrf

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

type VrfInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VrfInfo_KEYS) Reset()         { *m = VrfInfo_KEYS{} }
func (m *VrfInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*VrfInfo_KEYS) ProtoMessage()    {}
func (*VrfInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d21d3a79e81a1c5, []int{0}
}

func (m *VrfInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrfInfo_KEYS.Unmarshal(m, b)
}
func (m *VrfInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrfInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *VrfInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfInfo_KEYS.Merge(m, src)
}
func (m *VrfInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_VrfInfo_KEYS.Size(m)
}
func (m *VrfInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VrfInfo_KEYS proto.InternalMessageInfo

type VrfInfo struct {
	VrfName              string   `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VrfInfo) Reset()         { *m = VrfInfo{} }
func (m *VrfInfo) String() string { return proto.CompactTextString(m) }
func (*VrfInfo) ProtoMessage()    {}
func (*VrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d21d3a79e81a1c5, []int{1}
}

func (m *VrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VrfInfo.Unmarshal(m, b)
}
func (m *VrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VrfInfo.Marshal(b, m, deterministic)
}
func (m *VrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VrfInfo.Merge(m, src)
}
func (m *VrfInfo) XXX_Size() int {
	return xxx_messageInfo_VrfInfo.Size(m)
}
func (m *VrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VrfInfo proto.InternalMessageInfo

func (m *VrfInfo) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func init() {
	proto.RegisterType((*VrfInfo_KEYS)(nil), "cisco_ios_xr_lib_mpp_oper.management_plane_protection.outband.vrf.vrf_info_KEYS")
	proto.RegisterType((*VrfInfo)(nil), "cisco_ios_xr_lib_mpp_oper.management_plane_protection.outband.vrf.vrf_info")
}

func init() { proto.RegisterFile("vrf_info.proto", fileDescriptor_0d21d3a79e81a1c5) }

var fileDescriptor_0d21d3a79e81a1c5 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2b, 0x4a, 0x8b,
	0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x4c, 0xce, 0x2c, 0x4e,
	0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0xc9, 0x4c, 0x8a, 0xcf, 0x2d, 0x28,
	0x88, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0xcb, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0xcd, 0x4d, 0xcd, 0x2b,
	0x89, 0x2f, 0xc8, 0x49, 0xcc, 0x4b, 0x8d, 0x07, 0xe9, 0x48, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3,
	0xcb, 0x2f, 0x2d, 0x49, 0x4a, 0xcc, 0x4b, 0xd1, 0x2b, 0x2b, 0x4a, 0x53, 0xe2, 0xe7, 0xe2, 0x85,
	0x19, 0x1a, 0xef, 0xed, 0x1a, 0x19, 0xac, 0xa4, 0xca, 0xc5, 0x01, 0x13, 0x10, 0x92, 0x84, 0xb0,
	0xf3, 0x12, 0x73, 0x53, 0x25, 0x8c, 0x14, 0x18, 0x35, 0x38, 0x83, 0xd8, 0xcb, 0x8a, 0xd2, 0xfc,
	0x12, 0x73, 0x53, 0x93, 0xd8, 0xc0, 0x2e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xc0,
	0xc8, 0x9c, 0x93, 0x00, 0x00, 0x00,
}
