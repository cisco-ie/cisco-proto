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
// source: bfd_mgmt_client_brief_information.proto

package cisco_ios_xr_ip_bfd_oper_bfd_client_briefs_client_brief

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

type BfdMgmtClientBriefInformation_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtClientBriefInformation_KEYS) Reset()         { *m = BfdMgmtClientBriefInformation_KEYS{} }
func (m *BfdMgmtClientBriefInformation_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtClientBriefInformation_KEYS) ProtoMessage()    {}
func (*BfdMgmtClientBriefInformation_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c5589975c0de305, []int{0}
}

func (m *BfdMgmtClientBriefInformation_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtClientBriefInformation_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtClientBriefInformation_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS.Merge(m, src)
}
func (m *BfdMgmtClientBriefInformation_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS.Size(m)
}
func (m *BfdMgmtClientBriefInformation_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtClientBriefInformation_KEYS proto.InternalMessageInfo

func (m *BfdMgmtClientBriefInformation_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type BfdMgmtClientBriefInformation struct {
	NameXr               string   `protobuf:"bytes,50,opt,name=name_xr,json=nameXr,proto3" json:"name_xr,omitempty"`
	NodeId               string   `protobuf:"bytes,51,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	SessionCount         uint32   `protobuf:"varint,52,opt,name=session_count,json=sessionCount,proto3" json:"session_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtClientBriefInformation) Reset()         { *m = BfdMgmtClientBriefInformation{} }
func (m *BfdMgmtClientBriefInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtClientBriefInformation) ProtoMessage()    {}
func (*BfdMgmtClientBriefInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c5589975c0de305, []int{1}
}

func (m *BfdMgmtClientBriefInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtClientBriefInformation.Unmarshal(m, b)
}
func (m *BfdMgmtClientBriefInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtClientBriefInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtClientBriefInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtClientBriefInformation.Merge(m, src)
}
func (m *BfdMgmtClientBriefInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtClientBriefInformation.Size(m)
}
func (m *BfdMgmtClientBriefInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtClientBriefInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtClientBriefInformation proto.InternalMessageInfo

func (m *BfdMgmtClientBriefInformation) GetNameXr() string {
	if m != nil {
		return m.NameXr
	}
	return ""
}

func (m *BfdMgmtClientBriefInformation) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BfdMgmtClientBriefInformation) GetSessionCount() uint32 {
	if m != nil {
		return m.SessionCount
	}
	return 0
}

func init() {
	proto.RegisterType((*BfdMgmtClientBriefInformation_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.client_briefs.client_brief.bfd_mgmt_client_brief_information_KEYS")
	proto.RegisterType((*BfdMgmtClientBriefInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.client_briefs.client_brief.bfd_mgmt_client_brief_information")
}

func init() {
	proto.RegisterFile("bfd_mgmt_client_brief_information.proto", fileDescriptor_7c5589975c0de305)
}

var fileDescriptor_7c5589975c0de305 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x29, 0x48, 0xc5, 0x60, 0x97, 0x2c, 0x66, 0xac, 0x15, 0xb4, 0x53, 0x07, 0x2b, 0xb8,
	0xb8, 0x89, 0x83, 0xb8, 0xd5, 0x45, 0xa7, 0x4b, 0xf3, 0x4f, 0x2e, 0x98, 0xdc, 0x92, 0xe4, 0x41,
	0x79, 0x9f, 0xfe, 0x91, 0xd0, 0xe1, 0xbd, 0xa9, 0xdb, 0x3d, 0xe7, 0xdc, 0x1f, 0x87, 0xc3, 0x9e,
	0xa4, 0xd5, 0xe0, 0xfe, 0x5c, 0x02, 0xf5, 0x8f, 0xc6, 0x27, 0x90, 0x01, 0x8d, 0x05, 0xf4, 0x96,
	0x82, 0x9b, 0x13, 0x92, 0x1f, 0x96, 0x40, 0x89, 0xf8, 0xab, 0xc2, 0xa8, 0x08, 0x90, 0x22, 0xac,
	0x01, 0x70, 0x81, 0x0c, 0xd2, 0x62, 0xc2, 0x20, 0xad, 0x1e, 0xce, 0xe1, 0x78, 0xa1, 0xba, 0x37,
	0xf6, 0xb8, 0xdb, 0x01, 0x5f, 0x1f, 0xbf, 0xdf, 0x9c, 0xb3, 0x2b, 0x3f, 0x3b, 0x23, 0xaa, 0xb6,
	0xea, 0x6f, 0xa6, 0x72, 0x77, 0x47, 0x76, 0xbf, 0x4b, 0xf3, 0x3b, 0x76, 0x9d, 0x9f, 0x61, 0x0d,
	0xe2, 0xb9, 0xb0, 0x75, 0x96, 0x3f, 0xa1, 0x04, 0xa4, 0x0d, 0xa0, 0x16, 0xe3, 0x16, 0x90, 0x36,
	0x9f, 0x9a, 0x3f, 0xb0, 0x26, 0x9a, 0x18, 0x73, 0xb5, 0xa2, 0x83, 0x4f, 0xe2, 0xa5, 0xad, 0xfa,
	0x66, 0xba, 0xdd, 0xcc, 0xf7, 0xec, 0xc9, 0xba, 0x2c, 0x1f, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x7c, 0x53, 0xec, 0x24, 0x01, 0x00, 0x00,
}
