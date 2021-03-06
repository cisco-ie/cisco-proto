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
// source: pim_ma_summary.proto

package cisco_ios_xr_ipv4_pim_oper_pim_ma_standby_pim_ma_summary

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

type PimMaSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimMaSummary_KEYS) Reset()         { *m = PimMaSummary_KEYS{} }
func (m *PimMaSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimMaSummary_KEYS) ProtoMessage()    {}
func (*PimMaSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5164df2bfba2c7, []int{0}
}

func (m *PimMaSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMaSummary_KEYS.Unmarshal(m, b)
}
func (m *PimMaSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMaSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *PimMaSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMaSummary_KEYS.Merge(m, src)
}
func (m *PimMaSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimMaSummary_KEYS.Size(m)
}
func (m *PimMaSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMaSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimMaSummary_KEYS proto.InternalMessageInfo

type PimMaSummary struct {
	IsImConnectionOpen         bool     `protobuf:"varint,50,opt,name=is_im_connection_open,json=isImConnectionOpen,proto3" json:"is_im_connection_open,omitempty"`
	IsNetioConnectionOpen      bool     `protobuf:"varint,51,opt,name=is_netio_connection_open,json=isNetioConnectionOpen,proto3" json:"is_netio_connection_open,omitempty"`
	IsEdmConnectionOpen        bool     `protobuf:"varint,52,opt,name=is_edm_connection_open,json=isEdmConnectionOpen,proto3" json:"is_edm_connection_open,omitempty"`
	IsStandbyEdmConnectionOpen bool     `protobuf:"varint,53,opt,name=is_standby_edm_connection_open,json=isStandbyEdmConnectionOpen,proto3" json:"is_standby_edm_connection_open,omitempty"`
	EncapInterfaceCount        uint32   `protobuf:"varint,54,opt,name=encap_interface_count,json=encapInterfaceCount,proto3" json:"encap_interface_count,omitempty"`
	DecapInterfaceCount        uint32   `protobuf:"varint,55,opt,name=decap_interface_count,json=decapInterfaceCount,proto3" json:"decap_interface_count,omitempty"`
	MdtInterfaceCount          uint32   `protobuf:"varint,56,opt,name=mdt_interface_count,json=mdtInterfaceCount,proto3" json:"mdt_interface_count,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *PimMaSummary) Reset()         { *m = PimMaSummary{} }
func (m *PimMaSummary) String() string { return proto.CompactTextString(m) }
func (*PimMaSummary) ProtoMessage()    {}
func (*PimMaSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e5164df2bfba2c7, []int{1}
}

func (m *PimMaSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMaSummary.Unmarshal(m, b)
}
func (m *PimMaSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMaSummary.Marshal(b, m, deterministic)
}
func (m *PimMaSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMaSummary.Merge(m, src)
}
func (m *PimMaSummary) XXX_Size() int {
	return xxx_messageInfo_PimMaSummary.Size(m)
}
func (m *PimMaSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMaSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PimMaSummary proto.InternalMessageInfo

func (m *PimMaSummary) GetIsImConnectionOpen() bool {
	if m != nil {
		return m.IsImConnectionOpen
	}
	return false
}

func (m *PimMaSummary) GetIsNetioConnectionOpen() bool {
	if m != nil {
		return m.IsNetioConnectionOpen
	}
	return false
}

func (m *PimMaSummary) GetIsEdmConnectionOpen() bool {
	if m != nil {
		return m.IsEdmConnectionOpen
	}
	return false
}

func (m *PimMaSummary) GetIsStandbyEdmConnectionOpen() bool {
	if m != nil {
		return m.IsStandbyEdmConnectionOpen
	}
	return false
}

func (m *PimMaSummary) GetEncapInterfaceCount() uint32 {
	if m != nil {
		return m.EncapInterfaceCount
	}
	return 0
}

func (m *PimMaSummary) GetDecapInterfaceCount() uint32 {
	if m != nil {
		return m.DecapInterfaceCount
	}
	return 0
}

func (m *PimMaSummary) GetMdtInterfaceCount() uint32 {
	if m != nil {
		return m.MdtInterfaceCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PimMaSummary_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim_ma.standby.pim_ma_summary.pim_ma_summary_KEYS")
	proto.RegisterType((*PimMaSummary)(nil), "cisco_ios_xr_ipv4_pim_oper.pim_ma.standby.pim_ma_summary.pim_ma_summary")
}

func init() { proto.RegisterFile("pim_ma_summary.proto", fileDescriptor_4e5164df2bfba2c7) }

var fileDescriptor_4e5164df2bfba2c7 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd1, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0x71, 0x86, 0x20, 0x12, 0x50, 0xb0, 0xb5, 0x52, 0x3c, 0xc8, 0xd8, 0x69, 0xa7, 0x82,
	0xeb, 0x74, 0x3b, 0x3b, 0x76, 0x18, 0x82, 0xc2, 0x76, 0xf2, 0xf4, 0xe8, 0x92, 0x08, 0xef, 0x90,
	0x97, 0x90, 0x97, 0x89, 0xfb, 0x10, 0x7e, 0x67, 0x49, 0xaa, 0x42, 0xdb, 0x5d, 0xdf, 0xff, 0xfd,
	0xfa, 0x28, 0x11, 0x37, 0x0e, 0x0d, 0x98, 0x06, 0xf8, 0x60, 0x4c, 0xe3, 0x8f, 0x95, 0xf3, 0x36,
	0xd8, 0x6c, 0x29, 0x91, 0xa5, 0x05, 0xb4, 0x0c, 0x5f, 0x1e, 0xd0, 0x7d, 0xce, 0x21, 0xee, 0x59,
	0xa7, 0x7d, 0xd5, 0x82, 0x8a, 0x43, 0x43, 0x6a, 0x7f, 0xac, 0xba, 0x7e, 0x52, 0x88, 0xbc, 0x3b,
	0x81, 0x97, 0xf5, 0xfb, 0x6e, 0xf2, 0x7d, 0x26, 0xae, 0xba, 0xf3, 0xec, 0x41, 0x14, 0xc8, 0x80,
	0x06, 0xa4, 0x25, 0xd2, 0x32, 0xa0, 0xa5, 0x78, 0x80, 0xca, 0xd9, 0x78, 0x34, 0xbd, 0xd8, 0x66,
	0xc8, 0x1b, 0xb3, 0xfa, 0x4f, 0x6f, 0x4e, 0x53, 0xb6, 0x10, 0x25, 0x32, 0x90, 0x0e, 0x68, 0x07,
	0xaa, 0x4e, 0xaa, 0x40, 0x7e, 0x8d, 0xb9, 0x07, 0x6b, 0x71, 0x8b, 0x0c, 0x5a, 0x0d, 0x8f, 0xcd,
	0x13, 0xcb, 0x91, 0xd7, 0xaa, 0x7f, 0xed, 0x59, 0xdc, 0x23, 0xc3, 0xef, 0x7f, 0x9e, 0xc4, 0x8f,
	0x09, 0xdf, 0x21, 0xef, 0xda, 0xa5, 0xe1, 0x37, 0x66, 0xa2, 0xd0, 0x24, 0x1b, 0x07, 0x48, 0x41,
	0xfb, 0x8f, 0x46, 0x6a, 0x90, 0xf6, 0x40, 0xa1, 0x7c, 0x1a, 0x8f, 0xa6, 0x97, 0xdb, 0x3c, 0xc5,
	0xcd, 0x5f, 0x5b, 0xc5, 0x14, 0x8d, 0xd2, 0xa7, 0xcc, 0xa2, 0x35, 0x29, 0xf6, 0x4c, 0x25, 0x72,
	0xa3, 0xc2, 0x40, 0x2c, 0x93, 0xb8, 0x36, 0x2a, 0x74, 0xf7, 0xf7, 0xe7, 0xe9, 0x9d, 0xeb, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0xf6, 0xf5, 0xdc, 0xff, 0x01, 0x00, 0x00,
}
