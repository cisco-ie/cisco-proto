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
// source: lspv_interface_info_brief.proto

package cisco_ios_xr_mpls_oam_oper_mpls_oam_interface_briefs_brief

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

type LspvInterfaceInfoBrief_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LspvInterfaceInfoBrief_KEYS) Reset()         { *m = LspvInterfaceInfoBrief_KEYS{} }
func (m *LspvInterfaceInfoBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*LspvInterfaceInfoBrief_KEYS) ProtoMessage()    {}
func (*LspvInterfaceInfoBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ea21385a0ab1c6, []int{0}
}

func (m *LspvInterfaceInfoBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LspvInterfaceInfoBrief_KEYS.Unmarshal(m, b)
}
func (m *LspvInterfaceInfoBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LspvInterfaceInfoBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *LspvInterfaceInfoBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LspvInterfaceInfoBrief_KEYS.Merge(m, src)
}
func (m *LspvInterfaceInfoBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_LspvInterfaceInfoBrief_KEYS.Size(m)
}
func (m *LspvInterfaceInfoBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LspvInterfaceInfoBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LspvInterfaceInfoBrief_KEYS proto.InternalMessageInfo

func (m *LspvInterfaceInfoBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type LspvInterfaceInfoBrief struct {
	InterfaceNameXr      string   `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	State                string   `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	Mtu                  uint32   `protobuf:"varint,52,opt,name=mtu,proto3" json:"mtu,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,53,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	PrefixLengthV6       uint32   `protobuf:"varint,54,opt,name=prefix_length_v6,json=prefixLengthV6,proto3" json:"prefix_length_v6,omitempty"`
	PrimaryAddress       string   `protobuf:"bytes,55,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	PrimaryAddressV6     string   `protobuf:"bytes,56,opt,name=primary_address_v6,json=primaryAddressV6,proto3" json:"primary_address_v6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LspvInterfaceInfoBrief) Reset()         { *m = LspvInterfaceInfoBrief{} }
func (m *LspvInterfaceInfoBrief) String() string { return proto.CompactTextString(m) }
func (*LspvInterfaceInfoBrief) ProtoMessage()    {}
func (*LspvInterfaceInfoBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ea21385a0ab1c6, []int{1}
}

func (m *LspvInterfaceInfoBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LspvInterfaceInfoBrief.Unmarshal(m, b)
}
func (m *LspvInterfaceInfoBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LspvInterfaceInfoBrief.Marshal(b, m, deterministic)
}
func (m *LspvInterfaceInfoBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LspvInterfaceInfoBrief.Merge(m, src)
}
func (m *LspvInterfaceInfoBrief) XXX_Size() int {
	return xxx_messageInfo_LspvInterfaceInfoBrief.Size(m)
}
func (m *LspvInterfaceInfoBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_LspvInterfaceInfoBrief.DiscardUnknown(m)
}

var xxx_messageInfo_LspvInterfaceInfoBrief proto.InternalMessageInfo

func (m *LspvInterfaceInfoBrief) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *LspvInterfaceInfoBrief) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *LspvInterfaceInfoBrief) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *LspvInterfaceInfoBrief) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *LspvInterfaceInfoBrief) GetPrefixLengthV6() uint32 {
	if m != nil {
		return m.PrefixLengthV6
	}
	return 0
}

func (m *LspvInterfaceInfoBrief) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *LspvInterfaceInfoBrief) GetPrimaryAddressV6() string {
	if m != nil {
		return m.PrimaryAddressV6
	}
	return ""
}

func init() {
	proto.RegisterType((*LspvInterfaceInfoBrief_KEYS)(nil), "cisco_ios_xr_mpls_oam_oper.mpls_oam.interface.briefs.brief.lspv_interface_info_brief_KEYS")
	proto.RegisterType((*LspvInterfaceInfoBrief)(nil), "cisco_ios_xr_mpls_oam_oper.mpls_oam.interface.briefs.brief.lspv_interface_info_brief")
}

func init() { proto.RegisterFile("lspv_interface_info_brief.proto", fileDescriptor_97ea21385a0ab1c6) }

var fileDescriptor_97ea21385a0ab1c6 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x99, 0xa2, 0xe0, 0xc3, 0x76, 0x35, 0x78, 0x88, 0x17, 0x1d, 0x13, 0xb1, 0x88, 0xf4,
	0xe0, 0xb4, 0x8a, 0x37, 0x0f, 0xe2, 0x41, 0xf1, 0x30, 0x61, 0xe8, 0xe9, 0x91, 0x75, 0xa9, 0x06,
	0x9a, 0x26, 0xbc, 0xc4, 0x51, 0x3f, 0x88, 0xdf, 0x57, 0x96, 0xea, 0xb4, 0xc2, 0x4e, 0xc9, 0xfb,
	0xff, 0x7f, 0xef, 0x47, 0x20, 0x70, 0x50, 0x39, 0x3b, 0x47, 0x55, 0x7b, 0x49, 0xa5, 0x28, 0x24,
	0xaa, 0xba, 0x34, 0x38, 0x25, 0x25, 0xcb, 0xcc, 0x92, 0xf1, 0x86, 0x5d, 0x17, 0xca, 0x15, 0x06,
	0x95, 0x71, 0xd8, 0x10, 0x6a, 0x5b, 0x39, 0x34, 0x42, 0xa3, 0xb1, 0x92, 0xb2, 0x9f, 0x29, 0x5b,
	0xee, 0x67, 0x61, 0xd5, 0xb5, 0xc7, 0xf0, 0x0e, 0xf6, 0x57, 0xea, 0xf1, 0xfe, 0xf6, 0xe5, 0x89,
	0x1d, 0x41, 0xfc, 0x5b, 0xd6, 0x42, 0x4b, 0xde, 0x1b, 0xf4, 0xd2, 0xad, 0x71, 0xb4, 0x4c, 0x1f,
	0x85, 0x96, 0xc3, 0xcf, 0x35, 0xd8, 0x5b, 0x69, 0x62, 0x27, 0xb0, 0xd3, 0x95, 0x60, 0x43, 0xfc,
	0x2c, 0x78, 0xfa, 0x1d, 0xcf, 0x33, 0xb1, 0x5d, 0xd8, 0x70, 0x5e, 0x78, 0xc9, 0x47, 0xa1, 0x6f,
	0x07, 0x96, 0xc0, 0xba, 0xf6, 0xef, 0xfc, 0x7c, 0xd0, 0x4b, 0xa3, 0xf1, 0xe2, 0xca, 0x0e, 0x21,
	0xb2, 0x24, 0x4b, 0xd5, 0x60, 0x25, 0xeb, 0x57, 0xff, 0xc6, 0x2f, 0x42, 0xb7, 0xdd, 0x86, 0x0f,
	0x21, 0x63, 0x29, 0x24, 0x1d, 0x08, 0xe7, 0x39, 0xcf, 0x03, 0x17, 0xff, 0xe5, 0x26, 0x39, 0x3b,
	0x86, 0xbe, 0x25, 0xa5, 0x05, 0x7d, 0xa0, 0x98, 0xcd, 0x48, 0x3a, 0xc7, 0x2f, 0xc3, 0x03, 0xe2,
	0xef, 0xf8, 0xa6, 0x4d, 0xd9, 0x29, 0xb0, 0x7f, 0xe0, 0x42, 0x7a, 0x15, 0xd8, 0xa4, 0xcb, 0x4e,
	0xf2, 0xe9, 0x66, 0xf8, 0xa3, 0xd1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x4a, 0x9e, 0x43,
	0xc6, 0x01, 0x00, 0x00,
}
