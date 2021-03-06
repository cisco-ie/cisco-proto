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
// source: ipv4_if_brief.proto

package cisco_ios_xr_ipv4_ma_oper_augment_interfaces_interface_vrfs_vrf_brief

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

type Ipv4IfBrief_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief_KEYS) Reset()         { *m = Ipv4IfBrief_KEYS{} }
func (m *Ipv4IfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief_KEYS) ProtoMessage()    {}
func (*Ipv4IfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_306830dfda42ed3e, []int{0}
}

func (m *Ipv4IfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Unmarshal(m, b)
}
func (m *Ipv4IfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4IfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief_KEYS.Merge(m, src)
}
func (m *Ipv4IfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief_KEYS.Size(m)
}
func (m *Ipv4IfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief_KEYS proto.InternalMessageInfo

func (m *Ipv4IfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4IfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ipv4IfBrief struct {
	PrimaryAddress       string   `protobuf:"bytes,50,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	VrfId                uint32   `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	LineState            string   `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfBrief) Reset()         { *m = Ipv4IfBrief{} }
func (m *Ipv4IfBrief) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfBrief) ProtoMessage()    {}
func (*Ipv4IfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_306830dfda42ed3e, []int{1}
}

func (m *Ipv4IfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfBrief.Unmarshal(m, b)
}
func (m *Ipv4IfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfBrief.Marshal(b, m, deterministic)
}
func (m *Ipv4IfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfBrief.Merge(m, src)
}
func (m *Ipv4IfBrief) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfBrief.Size(m)
}
func (m *Ipv4IfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfBrief proto.InternalMessageInfo

func (m *Ipv4IfBrief) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *Ipv4IfBrief) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Ipv4IfBrief) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv4IfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.brief.ipv4_if_brief_KEYS")
	proto.RegisterType((*Ipv4IfBrief)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.brief.ipv4_if_brief")
}

func init() { proto.RegisterFile("ipv4_if_brief.proto", fileDescriptor_306830dfda42ed3e) }

var fileDescriptor_306830dfda42ed3e = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xc1, 0x6a, 0x07, 0xb6, 0x42, 0x44, 0x58, 0x0f, 0x42, 0x29, 0x88, 0x3d, 0xe5,
	0x60, 0xfb, 0x02, 0x1e, 0x7a, 0x10, 0xc1, 0x43, 0x0b, 0x82, 0xa7, 0x21, 0xdd, 0x9d, 0xc8, 0x80,
	0x49, 0x96, 0x49, 0x0c, 0xfa, 0xf6, 0x92, 0x54, 0x2a, 0xbd, 0x0c, 0xc3, 0x37, 0xff, 0xff, 0x1d,
	0x06, 0xae, 0x79, 0xcc, 0x6b, 0x64, 0x8b, 0x7b, 0x61, 0xb2, 0x7a, 0x94, 0x90, 0x82, 0xda, 0xf4,
	0x1c, 0xfb, 0x80, 0x1c, 0x22, 0x7e, 0x0b, 0xd6, 0x84, 0x33, 0x18, 0x46, 0x12, 0x6d, 0xbe, 0x3e,
	0x1c, 0xf9, 0xa4, 0xd9, 0x27, 0x12, 0x6b, 0x7a, 0x8a, 0xff, 0xab, 0xce, 0x62, 0x63, 0x19, 0xba,
	0xca, 0x16, 0x6f, 0xa0, 0x4e, 0xec, 0xf8, 0xb2, 0x79, 0xdf, 0xa9, 0x7b, 0x98, 0x1d, 0x1b, 0xe8,
	0x8d, 0xa3, 0xae, 0x99, 0x37, 0xcb, 0xe9, 0xb6, 0x3d, 0xd2, 0x57, 0xe3, 0x48, 0xdd, 0xc2, 0x65,
	0x16, 0x7b, 0x08, 0x9c, 0xd5, 0xc0, 0x45, 0x16, 0x5b, 0x4e, 0x0b, 0x0f, 0xed, 0x89, 0x57, 0x3d,
	0xc0, 0xd5, 0x28, 0xec, 0x8c, 0xfc, 0xa0, 0x19, 0x06, 0xa1, 0x18, 0xbb, 0xc7, 0x5a, 0x99, 0xfd,
	0xe1, 0xa7, 0x03, 0x55, 0x37, 0x30, 0x29, 0x52, 0x1e, 0xba, 0xd5, 0xbc, 0x59, 0xb6, 0xdb, 0xf3,
	0x2c, 0xf6, 0x79, 0x50, 0x77, 0x00, 0x9f, 0xec, 0x09, 0x63, 0x32, 0x89, 0xba, 0x75, 0xad, 0x4e,
	0x0b, 0xd9, 0x15, 0xb0, 0x9f, 0xd4, 0xaf, 0xac, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x22, 0xbb,
	0x9c, 0xe1, 0x2c, 0x01, 0x00, 0x00,
}
