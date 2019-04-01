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
// source: rpl_references_info.proto

package cisco_ios_xr_policy_repository_oper_routing_policy_shadow_sets_rd_sets_set_used_by

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

type RplReferencesInfo_KEYS struct {
	SetName              string   `protobuf:"bytes,1,opt,name=set_name,json=setName,proto3" json:"set_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RplReferencesInfo_KEYS) Reset()         { *m = RplReferencesInfo_KEYS{} }
func (m *RplReferencesInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RplReferencesInfo_KEYS) ProtoMessage()    {}
func (*RplReferencesInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea22454312bfc5fd, []int{0}
}

func (m *RplReferencesInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplReferencesInfo_KEYS.Unmarshal(m, b)
}
func (m *RplReferencesInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplReferencesInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RplReferencesInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplReferencesInfo_KEYS.Merge(m, src)
}
func (m *RplReferencesInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RplReferencesInfo_KEYS.Size(m)
}
func (m *RplReferencesInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RplReferencesInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RplReferencesInfo_KEYS proto.InternalMessageInfo

func (m *RplReferencesInfo_KEYS) GetSetName() string {
	if m != nil {
		return m.SetName
	}
	return ""
}

type RefInfo struct {
	RoutePolicyName      string   `protobuf:"bytes,1,opt,name=route_policy_name,json=routePolicyName,proto3" json:"route_policy_name,omitempty"`
	UsedDirectly         bool     `protobuf:"varint,2,opt,name=used_directly,json=usedDirectly,proto3" json:"used_directly,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefInfo) Reset()         { *m = RefInfo{} }
func (m *RefInfo) String() string { return proto.CompactTextString(m) }
func (*RefInfo) ProtoMessage()    {}
func (*RefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea22454312bfc5fd, []int{1}
}

func (m *RefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefInfo.Unmarshal(m, b)
}
func (m *RefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefInfo.Marshal(b, m, deterministic)
}
func (m *RefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefInfo.Merge(m, src)
}
func (m *RefInfo) XXX_Size() int {
	return xxx_messageInfo_RefInfo.Size(m)
}
func (m *RefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RefInfo proto.InternalMessageInfo

func (m *RefInfo) GetRoutePolicyName() string {
	if m != nil {
		return m.RoutePolicyName
	}
	return ""
}

func (m *RefInfo) GetUsedDirectly() bool {
	if m != nil {
		return m.UsedDirectly
	}
	return false
}

func (m *RefInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RplReferencesInfo struct {
	Reference            []*RefInfo `protobuf:"bytes,50,rep,name=reference,proto3" json:"reference,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RplReferencesInfo) Reset()         { *m = RplReferencesInfo{} }
func (m *RplReferencesInfo) String() string { return proto.CompactTextString(m) }
func (*RplReferencesInfo) ProtoMessage()    {}
func (*RplReferencesInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea22454312bfc5fd, []int{2}
}

func (m *RplReferencesInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplReferencesInfo.Unmarshal(m, b)
}
func (m *RplReferencesInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplReferencesInfo.Marshal(b, m, deterministic)
}
func (m *RplReferencesInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplReferencesInfo.Merge(m, src)
}
func (m *RplReferencesInfo) XXX_Size() int {
	return xxx_messageInfo_RplReferencesInfo.Size(m)
}
func (m *RplReferencesInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RplReferencesInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RplReferencesInfo proto.InternalMessageInfo

func (m *RplReferencesInfo) GetReference() []*RefInfo {
	if m != nil {
		return m.Reference
	}
	return nil
}

func init() {
	proto.RegisterType((*RplReferencesInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.rd.sets.set.used_by.rpl_references_info_KEYS")
	proto.RegisterType((*RefInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.rd.sets.set.used_by.ref_info")
	proto.RegisterType((*RplReferencesInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.rd.sets.set.used_by.rpl_references_info")
}

func init() { proto.RegisterFile("rpl_references_info.proto", fileDescriptor_ea22454312bfc5fd) }

var fileDescriptor_ea22454312bfc5fd = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0x0b, 0x6b, 0x37, 0x2a, 0x62, 0x04, 0xe9, 0xde, 0x4a, 0xbd, 0x14, 0x0f, 0x39,
	0xac, 0xf8, 0x06, 0x7a, 0x12, 0x44, 0xea, 0x49, 0x10, 0x42, 0xb7, 0x9d, 0x6a, 0xa0, 0xdb, 0x09,
	0x33, 0x53, 0xb4, 0xbe, 0x85, 0x6f, 0x2c, 0x8d, 0x5b, 0xf5, 0xb0, 0xc7, 0x3d, 0x85, 0xfc, 0xbe,
	0x2f, 0xf9, 0xe5, 0x8f, 0x5a, 0x92, 0x6f, 0x2d, 0x41, 0x03, 0x04, 0x5d, 0x05, 0x6c, 0x5d, 0xd7,
	0xa0, 0xf1, 0x84, 0x82, 0xba, 0xa8, 0x1c, 0x57, 0x68, 0x1d, 0xb2, 0xfd, 0x20, 0xeb, 0xb1, 0x75,
	0xd5, 0x60, 0x09, 0x3c, 0xb2, 0x13, 0xa4, 0xc1, 0xa2, 0x07, 0x32, 0x84, 0xbd, 0xb8, 0xee, 0x75,
	0x8a, 0xf9, 0xad, 0xac, 0xf1, 0xdd, 0x30, 0x08, 0x1b, 0xaa, 0x7f, 0x46, 0x06, 0x31, 0x3d, 0x43,
	0x6d, 0xd7, 0x43, 0x76, 0xa3, 0x92, 0x1d, 0x42, 0x7b, 0x7f, 0xf7, 0xfc, 0xa4, 0x97, 0x2a, 0x66,
	0x10, 0xdb, 0x95, 0x1b, 0x48, 0xa2, 0x34, 0xca, 0x17, 0xc5, 0x21, 0x83, 0x3c, 0x94, 0x1b, 0xc8,
	0x58, 0xc5, 0x04, 0x4d, 0xe8, 0xea, 0x2b, 0x75, 0x36, 0x4a, 0x61, 0x52, 0xfe, 0xeb, 0x9f, 0x86,
	0xe0, 0x31, 0xf0, 0x71, 0x9d, 0xbe, 0x54, 0x27, 0xc1, 0x5c, 0x3b, 0x82, 0x4a, 0xda, 0x21, 0x39,
	0x48, 0xa3, 0x3c, 0x2e, 0x8e, 0x47, 0x78, 0xbb, 0x65, 0xfa, 0x42, 0xcd, 0x59, 0x4a, 0xe9, 0x39,
	0x99, 0x85, 0x5d, 0xb6, 0xb3, 0xec, 0x2b, 0x52, 0xe7, 0x3b, 0x0e, 0xab, 0x3f, 0xd5, 0xe2, 0x17,
	0x25, 0xab, 0x74, 0x96, 0x1f, 0xad, 0x5e, 0xcc, 0xfe, 0xdf, 0xca, 0x4c, 0x37, 0x2e, 0xfe, 0x74,
	0xeb, 0x79, 0xf8, 0x9a, 0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x41, 0xed, 0xc8, 0xb7,
	0x01, 0x00, 0x00,
}
