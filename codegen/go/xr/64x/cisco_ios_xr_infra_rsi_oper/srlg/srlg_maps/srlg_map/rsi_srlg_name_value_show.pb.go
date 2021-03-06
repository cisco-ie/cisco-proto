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
// source: rsi_srlg_name_value_show.proto

package cisco_ios_xr_infra_rsi_oper_srlg_srlg_maps_srlg_map

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

type RsiSrlgNameValueShow_KEYS struct {
	SrlgName             string   `protobuf:"bytes,1,opt,name=srlg_name,json=srlgName,proto3" json:"srlg_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgNameValueShow_KEYS) Reset()         { *m = RsiSrlgNameValueShow_KEYS{} }
func (m *RsiSrlgNameValueShow_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgNameValueShow_KEYS) ProtoMessage()    {}
func (*RsiSrlgNameValueShow_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1204a3faf78226f2, []int{0}
}

func (m *RsiSrlgNameValueShow_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgNameValueShow_KEYS.Unmarshal(m, b)
}
func (m *RsiSrlgNameValueShow_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgNameValueShow_KEYS.Marshal(b, m, deterministic)
}
func (m *RsiSrlgNameValueShow_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgNameValueShow_KEYS.Merge(m, src)
}
func (m *RsiSrlgNameValueShow_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgNameValueShow_KEYS.Size(m)
}
func (m *RsiSrlgNameValueShow_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgNameValueShow_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgNameValueShow_KEYS proto.InternalMessageInfo

func (m *RsiSrlgNameValueShow_KEYS) GetSrlgName() string {
	if m != nil {
		return m.SrlgName
	}
	return ""
}

type RsiSrlgNameValueShow struct {
	SrlgValue            uint32   `protobuf:"varint,50,opt,name=srlg_value,json=srlgValue,proto3" json:"srlg_value,omitempty"`
	SrlgNameXr           string   `protobuf:"bytes,51,opt,name=srlg_name_xr,json=srlgNameXr,proto3" json:"srlg_name_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgNameValueShow) Reset()         { *m = RsiSrlgNameValueShow{} }
func (m *RsiSrlgNameValueShow) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgNameValueShow) ProtoMessage()    {}
func (*RsiSrlgNameValueShow) Descriptor() ([]byte, []int) {
	return fileDescriptor_1204a3faf78226f2, []int{1}
}

func (m *RsiSrlgNameValueShow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgNameValueShow.Unmarshal(m, b)
}
func (m *RsiSrlgNameValueShow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgNameValueShow.Marshal(b, m, deterministic)
}
func (m *RsiSrlgNameValueShow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgNameValueShow.Merge(m, src)
}
func (m *RsiSrlgNameValueShow) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgNameValueShow.Size(m)
}
func (m *RsiSrlgNameValueShow) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgNameValueShow.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgNameValueShow proto.InternalMessageInfo

func (m *RsiSrlgNameValueShow) GetSrlgValue() uint32 {
	if m != nil {
		return m.SrlgValue
	}
	return 0
}

func (m *RsiSrlgNameValueShow) GetSrlgNameXr() string {
	if m != nil {
		return m.SrlgNameXr
	}
	return ""
}

func init() {
	proto.RegisterType((*RsiSrlgNameValueShow_KEYS)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.srlg_maps.srlg_map.rsi_srlg_name_value_show_KEYS")
	proto.RegisterType((*RsiSrlgNameValueShow)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.srlg_maps.srlg_map.rsi_srlg_name_value_show")
}

func init() { proto.RegisterFile("rsi_srlg_name_value_show.proto", fileDescriptor_1204a3faf78226f2) }

var fileDescriptor_1204a3faf78226f2 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2a, 0xce, 0x8c,
	0x2f, 0x2e, 0xca, 0x49, 0x8f, 0xcf, 0x4b, 0xcc, 0x4d, 0x8d, 0x2f, 0x4b, 0xcc, 0x29, 0x4d, 0x8d,
	0x2f, 0xce, 0xc8, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x4e, 0xce, 0x2c, 0x4e,
	0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0xcc, 0x4b, 0x2b, 0x4a, 0x8c, 0x07,
	0x69, 0xc9, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xe9, 0x03, 0x13, 0xf1, 0xb9, 0x89, 0x05, 0xc5, 0x70,
	0x96, 0x92, 0x0d, 0x97, 0x2c, 0x2e, 0x63, 0xe3, 0xbd, 0x5d, 0x23, 0x83, 0x85, 0xa4, 0xb9, 0x38,
	0xe1, 0x92, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54,
	0xa5, 0x68, 0x2e, 0x09, 0x5c, 0xba, 0x85, 0x64, 0xb9, 0xb8, 0xc0, 0xe2, 0x60, 0x21, 0x09, 0x23,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0xb0, 0x51, 0x61, 0x20, 0x01, 0x21, 0x05, 0x2e, 0x1e, 0x84, 0xb6,
	0x8a, 0x22, 0x09, 0x63, 0xb0, 0xd1, 0x5c, 0x30, 0xa3, 0x23, 0x8a, 0x92, 0xd8, 0xc0, 0xde, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xac, 0x8f, 0xd6, 0xf8, 0x00, 0x00, 0x00,
}
