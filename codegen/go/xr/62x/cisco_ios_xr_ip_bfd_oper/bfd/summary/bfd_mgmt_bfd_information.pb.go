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
// source: bfd_mgmt_bfd_information.proto

package cisco_ios_xr_ip_bfd_oper_bfd_summary

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

type BfdMgmtBfdInformation_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtBfdInformation_KEYS) Reset()         { *m = BfdMgmtBfdInformation_KEYS{} }
func (m *BfdMgmtBfdInformation_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtBfdInformation_KEYS) ProtoMessage()    {}
func (*BfdMgmtBfdInformation_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa4a1ce300f3c92, []int{0}
}

func (m *BfdMgmtBfdInformation_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtBfdInformation_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtBfdInformation_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtBfdInformation_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtBfdInformation_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtBfdInformation_KEYS.Merge(m, src)
}
func (m *BfdMgmtBfdInformation_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtBfdInformation_KEYS.Size(m)
}
func (m *BfdMgmtBfdInformation_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtBfdInformation_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtBfdInformation_KEYS proto.InternalMessageInfo

type BfdMgmtSessionStateInformation struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	DownCount            uint32   `protobuf:"varint,2,opt,name=down_count,json=downCount,proto3" json:"down_count,omitempty"`
	UpCount              uint32   `protobuf:"varint,3,opt,name=up_count,json=upCount,proto3" json:"up_count,omitempty"`
	UnknownCount         uint32   `protobuf:"varint,4,opt,name=unknown_count,json=unknownCount,proto3" json:"unknown_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtSessionStateInformation) Reset()         { *m = BfdMgmtSessionStateInformation{} }
func (m *BfdMgmtSessionStateInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSessionStateInformation) ProtoMessage()    {}
func (*BfdMgmtSessionStateInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa4a1ce300f3c92, []int{1}
}

func (m *BfdMgmtSessionStateInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSessionStateInformation.Unmarshal(m, b)
}
func (m *BfdMgmtSessionStateInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSessionStateInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSessionStateInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSessionStateInformation.Merge(m, src)
}
func (m *BfdMgmtSessionStateInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSessionStateInformation.Size(m)
}
func (m *BfdMgmtSessionStateInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSessionStateInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSessionStateInformation proto.InternalMessageInfo

func (m *BfdMgmtSessionStateInformation) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *BfdMgmtSessionStateInformation) GetDownCount() uint32 {
	if m != nil {
		return m.DownCount
	}
	return 0
}

func (m *BfdMgmtSessionStateInformation) GetUpCount() uint32 {
	if m != nil {
		return m.UpCount
	}
	return 0
}

func (m *BfdMgmtSessionStateInformation) GetUnknownCount() uint32 {
	if m != nil {
		return m.UnknownCount
	}
	return 0
}

type BfdMgmtBfdInformation struct {
	SessionState         *BfdMgmtSessionStateInformation `protobuf:"bytes,50,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *BfdMgmtBfdInformation) Reset()         { *m = BfdMgmtBfdInformation{} }
func (m *BfdMgmtBfdInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtBfdInformation) ProtoMessage()    {}
func (*BfdMgmtBfdInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa4a1ce300f3c92, []int{2}
}

func (m *BfdMgmtBfdInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtBfdInformation.Unmarshal(m, b)
}
func (m *BfdMgmtBfdInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtBfdInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtBfdInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtBfdInformation.Merge(m, src)
}
func (m *BfdMgmtBfdInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtBfdInformation.Size(m)
}
func (m *BfdMgmtBfdInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtBfdInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtBfdInformation proto.InternalMessageInfo

func (m *BfdMgmtBfdInformation) GetSessionState() *BfdMgmtSessionStateInformation {
	if m != nil {
		return m.SessionState
	}
	return nil
}

func init() {
	proto.RegisterType((*BfdMgmtBfdInformation_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.summary.bfd_mgmt_bfd_information_KEYS")
	proto.RegisterType((*BfdMgmtSessionStateInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.summary.bfd_mgmt_session_state_information")
	proto.RegisterType((*BfdMgmtBfdInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.summary.bfd_mgmt_bfd_information")
}

func init() { proto.RegisterFile("bfd_mgmt_bfd_information.proto", fileDescriptor_0fa4a1ce300f3c92) }

var fileDescriptor_0fa4a1ce300f3c92 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xbb, 0x4e, 0x04, 0x21,
	0x14, 0x06, 0xe0, 0xa0, 0xc6, 0xcb, 0xd9, 0x9d, 0x66, 0xaa, 0xb1, 0x58, 0x77, 0x83, 0x16, 0x5b,
	0x51, 0xac, 0x8f, 0x60, 0x4c, 0x4c, 0xec, 0x76, 0x2b, 0x2b, 0xc2, 0xdc, 0x0c, 0x51, 0x38, 0x04,
	0x0e, 0x51, 0x1f, 0xc1, 0xf7, 0xf0, 0x41, 0x0d, 0x2c, 0xf1, 0x52, 0x4c, 0xb4, 0x23, 0xff, 0xf7,
	0x43, 0x38, 0x00, 0x17, 0xed, 0xd8, 0x4b, 0xf3, 0x68, 0x48, 0xa6, 0x85, 0xb6, 0x23, 0x7a, 0xa3,
	0x48, 0xa3, 0x15, 0xce, 0x23, 0x61, 0x7d, 0xd5, 0xe9, 0xd0, 0xa1, 0xd4, 0x18, 0xe4, 0xab, 0x97,
	0xda, 0xe5, 0x1a, 0xba, 0xc1, 0x8b, 0x76, 0xec, 0x45, 0x88, 0xc6, 0x28, 0xff, 0xc6, 0x97, 0xb0,
	0x98, 0x3a, 0x47, 0xde, 0xdf, 0x3e, 0xec, 0xf8, 0x07, 0x03, 0xfe, 0xd5, 0x08, 0x43, 0x08, 0x49,
	0x02, 0x29, 0x1a, 0x7e, 0x76, 0xeb, 0x25, 0xcc, 0x08, 0x49, 0x3d, 0xcb, 0x0e, 0xa3, 0xa5, 0x86,
	0xad, 0xd8, 0xba, 0xda, 0x42, 0x8e, 0x6e, 0x52, 0x52, 0x2f, 0x00, 0x7a, 0x7c, 0xb1, 0xc5, 0x0f,
	0xb2, 0x9f, 0xa5, 0x64, 0xcf, 0xe7, 0x70, 0x1a, 0x5d, 0xc1, 0xc3, 0x8c, 0x27, 0xd1, 0xed, 0xe9,
	0x12, 0xaa, 0x68, 0x9f, 0xec, 0xf7, 0xe6, 0xa3, 0xec, 0xf3, 0x12, 0xe6, 0x12, 0x7f, 0x67, 0xd0,
	0x4c, 0x0d, 0x52, 0x1b, 0xa8, 0x7e, 0xdd, 0xbc, 0xd9, 0xac, 0xd8, 0x7a, 0xb6, 0xb9, 0x13, 0xff,
	0x79, 0x22, 0xf1, 0xf7, 0xf4, 0xdb, 0x79, 0xa1, 0x5d, 0x92, 0xf6, 0x38, 0x7f, 0xc0, 0xf5, 0x67,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x26, 0x17, 0x26, 0xa2, 0x01, 0x00, 0x00,
}
