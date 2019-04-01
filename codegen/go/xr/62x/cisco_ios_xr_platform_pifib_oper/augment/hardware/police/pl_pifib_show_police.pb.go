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
// source: pl_pifib_show_police.proto

package cisco_ios_xr_platform_pifib_oper_augment_hardware_police

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

type PlPifibShowPolice_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlPifibShowPolice_KEYS) Reset()         { *m = PlPifibShowPolice_KEYS{} }
func (m *PlPifibShowPolice_KEYS) String() string { return proto.CompactTextString(m) }
func (*PlPifibShowPolice_KEYS) ProtoMessage()    {}
func (*PlPifibShowPolice_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab3f3d102aea346, []int{0}
}

func (m *PlPifibShowPolice_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlPifibShowPolice_KEYS.Unmarshal(m, b)
}
func (m *PlPifibShowPolice_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlPifibShowPolice_KEYS.Marshal(b, m, deterministic)
}
func (m *PlPifibShowPolice_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlPifibShowPolice_KEYS.Merge(m, src)
}
func (m *PlPifibShowPolice_KEYS) XXX_Size() int {
	return xxx_messageInfo_PlPifibShowPolice_KEYS.Size(m)
}
func (m *PlPifibShowPolice_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PlPifibShowPolice_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PlPifibShowPolice_KEYS proto.InternalMessageInfo

type PlPifibShowPolice struct {
	FlowType             uint32   `protobuf:"varint,50,opt,name=flow_type,json=flowType,proto3" json:"flow_type,omitempty"`
	Rate                 uint32   `protobuf:"varint,51,opt,name=rate,proto3" json:"rate,omitempty"`
	Accepts              uint64   `protobuf:"varint,52,opt,name=accepts,proto3" json:"accepts,omitempty"`
	Drops                uint64   `protobuf:"varint,53,opt,name=drops,proto3" json:"drops,omitempty"`
	TosPrec              uint32   `protobuf:"varint,54,opt,name=tos_prec,json=tosPrec,proto3" json:"tos_prec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlPifibShowPolice) Reset()         { *m = PlPifibShowPolice{} }
func (m *PlPifibShowPolice) String() string { return proto.CompactTextString(m) }
func (*PlPifibShowPolice) ProtoMessage()    {}
func (*PlPifibShowPolice) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ab3f3d102aea346, []int{1}
}

func (m *PlPifibShowPolice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlPifibShowPolice.Unmarshal(m, b)
}
func (m *PlPifibShowPolice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlPifibShowPolice.Marshal(b, m, deterministic)
}
func (m *PlPifibShowPolice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlPifibShowPolice.Merge(m, src)
}
func (m *PlPifibShowPolice) XXX_Size() int {
	return xxx_messageInfo_PlPifibShowPolice.Size(m)
}
func (m *PlPifibShowPolice) XXX_DiscardUnknown() {
	xxx_messageInfo_PlPifibShowPolice.DiscardUnknown(m)
}

var xxx_messageInfo_PlPifibShowPolice proto.InternalMessageInfo

func (m *PlPifibShowPolice) GetFlowType() uint32 {
	if m != nil {
		return m.FlowType
	}
	return 0
}

func (m *PlPifibShowPolice) GetRate() uint32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *PlPifibShowPolice) GetAccepts() uint64 {
	if m != nil {
		return m.Accepts
	}
	return 0
}

func (m *PlPifibShowPolice) GetDrops() uint64 {
	if m != nil {
		return m.Drops
	}
	return 0
}

func (m *PlPifibShowPolice) GetTosPrec() uint32 {
	if m != nil {
		return m.TosPrec
	}
	return 0
}

func init() {
	proto.RegisterType((*PlPifibShowPolice_KEYS)(nil), "cisco_ios_xr_platform_pifib_oper.augment.hardware.police.pl_pifib_show_police_KEYS")
	proto.RegisterType((*PlPifibShowPolice)(nil), "cisco_ios_xr_platform_pifib_oper.augment.hardware.police.pl_pifib_show_police")
}

func init() { proto.RegisterFile("pl_pifib_show_police.proto", fileDescriptor_9ab3f3d102aea346) }

var fileDescriptor_9ab3f3d102aea346 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0x4e, 0x03, 0x31,
	0x0c, 0x86, 0x75, 0x52, 0xa1, 0xc5, 0x12, 0x4b, 0xd4, 0x21, 0xa5, 0x4b, 0xd5, 0xa9, 0xd3, 0x0d,
	0x14, 0x10, 0x2f, 0xc0, 0xc4, 0x82, 0x0a, 0x0b, 0x93, 0x95, 0xa6, 0x3e, 0x1a, 0x29, 0xc5, 0x96,
	0x63, 0x74, 0xf4, 0x35, 0x78, 0x62, 0x44, 0x4a, 0xb7, 0xdb, 0xfc, 0x7f, 0xb6, 0x3f, 0xe9, 0x87,
	0x1b, 0xc9, 0x28, 0xa9, 0x4b, 0x5b, 0x2c, 0x7b, 0xee, 0x51, 0x38, 0xa7, 0x48, 0xad, 0x28, 0x1b,
	0xbb, 0xc7, 0x98, 0x4a, 0x64, 0x4c, 0x5c, 0xf0, 0x5b, 0x51, 0x72, 0xb0, 0x8e, 0xf5, 0xf0, 0x7f,
	0xce, 0x42, 0xda, 0x86, 0xaf, 0x8f, 0x03, 0x7d, 0x5a, 0xbb, 0x0f, 0xba, 0xeb, 0x83, 0x52, 0x7b,
	0xfa, 0x5f, 0xce, 0x61, 0x36, 0xe4, 0xc5, 0xe7, 0xa7, 0xf7, 0xd7, 0xe5, 0x4f, 0x03, 0xd3, 0xa1,
	0xad, 0x9b, 0xc3, 0x55, 0x97, 0xb9, 0x47, 0x3b, 0x0a, 0xf9, 0xdb, 0x45, 0xb3, 0xba, 0xde, 0x4c,
	0xfe, 0xc0, 0xdb, 0x51, 0xc8, 0x39, 0x18, 0x69, 0x30, 0xf2, 0xeb, 0xca, 0xeb, 0xec, 0x3c, 0x8c,
	0x43, 0x8c, 0x24, 0x56, 0xfc, 0xdd, 0xa2, 0x59, 0x8d, 0x36, 0xe7, 0xe8, 0xa6, 0x70, 0xb1, 0x53,
	0x96, 0xe2, 0xef, 0x2b, 0x3f, 0x05, 0x37, 0x83, 0x89, 0x71, 0x41, 0x51, 0x8a, 0xfe, 0xa1, 0x7a,
	0xc6, 0xc6, 0xe5, 0x45, 0x29, 0x6e, 0x2f, 0x6b, 0xe5, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x75, 0x32, 0xb4, 0x40, 0x10, 0x01, 0x00, 0x00,
}
