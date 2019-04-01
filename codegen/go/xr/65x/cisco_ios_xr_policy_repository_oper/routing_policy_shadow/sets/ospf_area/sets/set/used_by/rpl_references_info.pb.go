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

package cisco_ios_xr_policy_repository_oper_routing_policy_shadow_sets_ospf_area_sets_set_used_by

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
	proto.RegisterType((*RplReferencesInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.ospf_area.sets.set.used_by.rpl_references_info_KEYS")
	proto.RegisterType((*RefInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.ospf_area.sets.set.used_by.ref_info")
	proto.RegisterType((*RplReferencesInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.ospf_area.sets.set.used_by.rpl_references_info")
}

func init() { proto.RegisterFile("rpl_references_info.proto", fileDescriptor_ea22454312bfc5fd) }

var fileDescriptor_ea22454312bfc5fd = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x2a, 0x95, 0xd4, 0x80, 0x10, 0x46, 0x42, 0xe9, 0x16, 0x85, 0x25, 0x62, 0xf0,
	0x50, 0xc4, 0x1b, 0xc0, 0x84, 0x84, 0x50, 0x98, 0x3a, 0x59, 0xae, 0x73, 0x01, 0x4b, 0x69, 0xce,
	0xba, 0x73, 0x04, 0x19, 0x79, 0x0c, 0xde, 0x16, 0xc5, 0x34, 0xc0, 0xd0, 0xb5, 0xe3, 0xfd, 0xbe,
	0xcf, 0xfe, 0xf9, 0x8f, 0x58, 0x92, 0x6f, 0x35, 0x41, 0x03, 0x04, 0x9d, 0x05, 0xd6, 0xae, 0x6b,
	0x50, 0x79, 0xc2, 0x80, 0x72, 0x6d, 0x1d, 0x5b, 0xd4, 0x0e, 0x59, 0x7f, 0x90, 0xf6, 0xd8, 0x3a,
	0x3b, 0x68, 0x02, 0x8f, 0xec, 0x02, 0xd2, 0xa0, 0xd1, 0x03, 0x29, 0xc2, 0x3e, 0xb8, 0xee, 0x75,
	0x8a, 0xf9, 0xcd, 0xd4, 0xf8, 0xae, 0x18, 0x02, 0x2b, 0x64, 0xdf, 0x68, 0x43, 0x60, 0x7e, 0x46,
	0x86, 0xa0, 0x7a, 0x86, 0x5a, 0x6f, 0x86, 0xe2, 0x4e, 0x64, 0x7b, 0xbc, 0xfa, 0xf1, 0x61, 0xfd,
	0x22, 0x97, 0x22, 0x65, 0x08, 0xba, 0x33, 0x5b, 0xc8, 0x92, 0x3c, 0x29, 0x17, 0xd5, 0x31, 0x43,
	0x78, 0x32, 0x5b, 0x28, 0x58, 0xa4, 0x04, 0x4d, 0xec, 0xca, 0x1b, 0x71, 0x31, 0xba, 0x61, 0x32,
	0xff, 0xeb, 0x9f, 0xc7, 0xe0, 0x39, 0xf2, 0x71, 0x9d, 0xbc, 0x16, 0x67, 0xd1, 0x5c, 0x3b, 0x02,
	0x1b, 0xda, 0x21, 0x3b, 0xca, 0x93, 0x32, 0xad, 0x4e, 0x47, 0x78, 0xbf, 0x63, 0xf2, 0x4a, 0xcc,
	0x39, 0x98, 0xd0, 0x73, 0x36, 0x8b, 0xbb, 0xec, 0xa6, 0xe2, 0x2b, 0x11, 0x97, 0x7b, 0x0e, 0x2b,
	0x3f, 0x13, 0xb1, 0xf8, 0x65, 0xd9, 0x2a, 0x9f, 0x95, 0x27, 0x2b, 0xab, 0x0e, 0xf6, 0x66, 0x6a,
	0xba, 0x79, 0xf5, 0x67, 0xdd, 0xcc, 0xe3, 0x4f, 0xdd, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x60, 0x84, 0x1a, 0xc6, 0x01, 0x00, 0x00,
}