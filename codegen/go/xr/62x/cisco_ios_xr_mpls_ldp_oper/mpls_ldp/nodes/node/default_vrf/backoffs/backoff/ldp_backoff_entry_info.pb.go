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
// source: ldp_backoff_entry_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_backoffs_backoff

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

type LdpBackoffEntryInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LsrId                string   `protobuf:"bytes,2,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,3,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffEntryInfo_KEYS) Reset()         { *m = LdpBackoffEntryInfo_KEYS{} }
func (m *LdpBackoffEntryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffEntryInfo_KEYS) ProtoMessage()    {}
func (*LdpBackoffEntryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b1b2b369a41567, []int{0}
}

func (m *LdpBackoffEntryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Merge(m, src)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffEntryInfo_KEYS.Size(m)
}
func (m *LdpBackoffEntryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffEntryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffEntryInfo_KEYS proto.InternalMessageInfo

func (m *LdpBackoffEntryInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpBackoffEntryInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpBackoffEntryInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

type LdpBackoffEntryInfo struct {
	BackoffSeconds       uint32   `protobuf:"varint,50,opt,name=backoff_seconds,json=backoffSeconds,proto3" json:"backoff_seconds,omitempty"`
	WaitingSeconds       uint32   `protobuf:"varint,51,opt,name=waiting_seconds,json=waitingSeconds,proto3" json:"waiting_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBackoffEntryInfo) Reset()         { *m = LdpBackoffEntryInfo{} }
func (m *LdpBackoffEntryInfo) String() string { return proto.CompactTextString(m) }
func (*LdpBackoffEntryInfo) ProtoMessage()    {}
func (*LdpBackoffEntryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b1b2b369a41567, []int{1}
}

func (m *LdpBackoffEntryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBackoffEntryInfo.Unmarshal(m, b)
}
func (m *LdpBackoffEntryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBackoffEntryInfo.Marshal(b, m, deterministic)
}
func (m *LdpBackoffEntryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBackoffEntryInfo.Merge(m, src)
}
func (m *LdpBackoffEntryInfo) XXX_Size() int {
	return xxx_messageInfo_LdpBackoffEntryInfo.Size(m)
}
func (m *LdpBackoffEntryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBackoffEntryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBackoffEntryInfo proto.InternalMessageInfo

func (m *LdpBackoffEntryInfo) GetBackoffSeconds() uint32 {
	if m != nil {
		return m.BackoffSeconds
	}
	return 0
}

func (m *LdpBackoffEntryInfo) GetWaitingSeconds() uint32 {
	if m != nil {
		return m.WaitingSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpBackoffEntryInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.backoffs.backoff.ldp_backoff_entry_info_KEYS")
	proto.RegisterType((*LdpBackoffEntryInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.backoffs.backoff.ldp_backoff_entry_info")
}

func init() { proto.RegisterFile("ldp_backoff_entry_info.proto", fileDescriptor_91b1b2b369a41567) }

var fileDescriptor_91b1b2b369a41567 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x19, 0xc5, 0x62, 0x83, 0x56, 0x08, 0x28, 0x03, 0x75, 0x51, 0x8a, 0x60, 0x57, 0x59,
	0xd8, 0x67, 0x70, 0x51, 0x0a, 0x2e, 0x3a, 0x2b, 0x57, 0x97, 0x4c, 0x7e, 0x24, 0x9a, 0xc9, 0x0d,
	0xb9, 0xa3, 0xd5, 0xb7, 0x97, 0xa4, 0x33, 0x5d, 0xcd, 0x26, 0x24, 0xdf, 0xf9, 0x38, 0x07, 0xc2,
	0x1e, 0xbd, 0x8e, 0xd0, 0x4a, 0xf5, 0x85, 0xd6, 0x82, 0x09, 0x7d, 0xfa, 0x03, 0x17, 0x2c, 0x8a,
	0x98, 0xb0, 0x47, 0xbe, 0x57, 0x8e, 0x14, 0x82, 0x43, 0x82, 0xdf, 0x04, 0x5d, 0xf4, 0x04, 0xd9,
	0xc7, 0x68, 0x92, 0x18, 0x5f, 0x22, 0xa0, 0x36, 0x54, 0x4e, 0xa1, 0x8d, 0x95, 0xdf, 0xbe, 0x87,
	0x9f, 0x64, 0xc5, 0x50, 0x4a, 0xe3, 0x65, 0x7d, 0x64, 0xcb, 0xe9, 0x31, 0xd8, 0xbf, 0xbe, 0x37,
	0x7c, 0xc9, 0xe6, 0xb9, 0x03, 0x82, 0xec, 0x4c, 0x5d, 0xad, 0xaa, 0xcd, 0xfc, 0x70, 0x9d, 0xc1,
	0x9b, 0xec, 0x0c, 0xbf, 0x67, 0x33, 0x4f, 0x09, 0x9c, 0xae, 0x2f, 0x4a, 0x72, 0xe5, 0x29, 0xed,
	0x34, 0x7f, 0x62, 0x0b, 0x2f, 0x5b, 0xe3, 0x81, 0xa2, 0x54, 0x26, 0xc7, 0x97, 0xab, 0x6a, 0x73,
	0x7b, 0xb8, 0x29, 0xb4, 0xc9, 0x70, 0xa7, 0xd7, 0x9f, 0xec, 0x61, 0x7a, 0x98, 0x3f, 0xb3, 0xbb,
	0x91, 0x92, 0x51, 0x18, 0x34, 0xd5, 0x2f, 0xa5, 0x60, 0x31, 0xe0, 0xe6, 0x44, 0xb3, 0x78, 0x94,
	0xae, 0x77, 0xe1, 0xe3, 0x2c, 0x6e, 0x4f, 0xe2, 0x80, 0x07, 0xb1, 0x9d, 0x95, 0x8f, 0xdb, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x4c, 0x37, 0x15, 0x58, 0x01, 0x00, 0x00,
}
