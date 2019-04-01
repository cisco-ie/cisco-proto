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
// source: ac_buffer_info_bag.proto

package cisco_ios_xr_infra_correlator_oper_correlator_buffer_status

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

type AcBufferInfoBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcBufferInfoBag_KEYS) Reset()         { *m = AcBufferInfoBag_KEYS{} }
func (m *AcBufferInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AcBufferInfoBag_KEYS) ProtoMessage()    {}
func (*AcBufferInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e7c6e6034d694b, []int{0}
}

func (m *AcBufferInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcBufferInfoBag_KEYS.Unmarshal(m, b)
}
func (m *AcBufferInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcBufferInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AcBufferInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcBufferInfoBag_KEYS.Merge(m, src)
}
func (m *AcBufferInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AcBufferInfoBag_KEYS.Size(m)
}
func (m *AcBufferInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AcBufferInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AcBufferInfoBag_KEYS proto.InternalMessageInfo

type AcBufferInfoBag struct {
	CurrentSize          uint32   `protobuf:"varint,50,opt,name=current_size,json=currentSize,proto3" json:"current_size,omitempty"`
	ConfiguredSize       uint32   `protobuf:"varint,51,opt,name=configured_size,json=configuredSize,proto3" json:"configured_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcBufferInfoBag) Reset()         { *m = AcBufferInfoBag{} }
func (m *AcBufferInfoBag) String() string { return proto.CompactTextString(m) }
func (*AcBufferInfoBag) ProtoMessage()    {}
func (*AcBufferInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_73e7c6e6034d694b, []int{1}
}

func (m *AcBufferInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcBufferInfoBag.Unmarshal(m, b)
}
func (m *AcBufferInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcBufferInfoBag.Marshal(b, m, deterministic)
}
func (m *AcBufferInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcBufferInfoBag.Merge(m, src)
}
func (m *AcBufferInfoBag) XXX_Size() int {
	return xxx_messageInfo_AcBufferInfoBag.Size(m)
}
func (m *AcBufferInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AcBufferInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_AcBufferInfoBag proto.InternalMessageInfo

func (m *AcBufferInfoBag) GetCurrentSize() uint32 {
	if m != nil {
		return m.CurrentSize
	}
	return 0
}

func (m *AcBufferInfoBag) GetConfiguredSize() uint32 {
	if m != nil {
		return m.ConfiguredSize
	}
	return 0
}

func init() {
	proto.RegisterType((*AcBufferInfoBag_KEYS)(nil), "cisco_ios_xr_infra_correlator_oper.correlator.buffer_status.ac_buffer_info_bag_KEYS")
	proto.RegisterType((*AcBufferInfoBag)(nil), "cisco_ios_xr_infra_correlator_oper.correlator.buffer_status.ac_buffer_info_bag")
}

func init() { proto.RegisterFile("ac_buffer_info_bag.proto", fileDescriptor_73e7c6e6034d694b) }

var fileDescriptor_73e7c6e6034d694b = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x3f, 0x0b, 0xc2, 0x30,
	0x10, 0x47, 0x71, 0x71, 0x88, 0xff, 0x20, 0x8b, 0x75, 0xd3, 0x2e, 0x3a, 0x75, 0xb0, 0xa3, 0xb3,
	0x93, 0x9b, 0x9d, 0x9c, 0xce, 0x34, 0x5e, 0xca, 0x81, 0xf4, 0xca, 0x25, 0x01, 0xe9, 0xa7, 0x17,
	0x63, 0xa1, 0x43, 0xc7, 0xfb, 0xbd, 0xc7, 0x83, 0x53, 0x99, 0xb1, 0x50, 0x47, 0xe7, 0x50, 0x80,
	0x5a, 0xc7, 0x50, 0x9b, 0xa6, 0xe8, 0x84, 0x03, 0xeb, 0x8b, 0x25, 0x6f, 0x19, 0x88, 0x3d, 0x7c,
	0x12, 0x14, 0x03, 0x96, 0x45, 0xf0, 0x6d, 0x02, 0x0b, 0x70, 0x87, 0x52, 0x8c, 0x77, 0x31, 0x44,
	0x7c, 0x30, 0x21, 0xfa, 0x7c, 0xa7, 0xb6, 0xd3, 0x30, 0xdc, 0xae, 0x8f, 0x2a, 0x7f, 0x2a, 0x3d,
	0x45, 0xfa, 0xa0, 0x96, 0x36, 0x8a, 0x60, 0x1b, 0xc0, 0x53, 0x8f, 0xd9, 0x79, 0x3f, 0x3b, 0xad,
	0xee, 0x8b, 0x61, 0xab, 0xa8, 0x47, 0x7d, 0x54, 0x1b, 0xcb, 0xad, 0xa3, 0x26, 0x0a, 0xbe, 0xfe,
	0x56, 0x99, 0xac, 0xf5, 0x38, 0xff, 0xc4, 0x7a, 0x9e, 0x1e, 0x28, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0x14, 0xe4, 0x85, 0xdc, 0x00, 0x00, 0x00,
}
