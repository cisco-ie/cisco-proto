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

package cisco_ios_xr_policy_repository_oper_routing_policy_shadow_sets_tag_sets_set_used_by

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
	proto.RegisterType((*RplReferencesInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.tag.sets.set.used_by.rpl_references_info_KEYS")
	proto.RegisterType((*RefInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.tag.sets.set.used_by.ref_info")
	proto.RegisterType((*RplReferencesInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.tag.sets.set.used_by.rpl_references_info")
}

func init() { proto.RegisterFile("rpl_references_info.proto", fileDescriptor_ea22454312bfc5fd) }

var fileDescriptor_ea22454312bfc5fd = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xe9, 0x37, 0x30, 0x5f, 0x27, 0x2a, 0x62, 0x04, 0xe9, 0xec, 0x4a, 0xdd, 0x14, 0x17,
	0x59, 0x8c, 0x78, 0x07, 0xba, 0x12, 0x44, 0x3a, 0x2b, 0x17, 0x12, 0x3a, 0xed, 0xe9, 0x18, 0xe8,
	0xf4, 0x84, 0x73, 0x4e, 0xd1, 0xe2, 0x5d, 0x78, 0xc5, 0xd2, 0x38, 0x55, 0x17, 0xb3, 0x75, 0x97,
	0x3c, 0xef, 0x9b, 0x3c, 0xf9, 0x51, 0x4b, 0xf2, 0xad, 0x25, 0x68, 0x80, 0xa0, 0xab, 0x80, 0xad,
	0xeb, 0x1a, 0x34, 0x9e, 0x50, 0x50, 0xaf, 0x2b, 0xc7, 0x15, 0x5a, 0x87, 0x6c, 0xdf, 0xc8, 0x7a,
	0x6c, 0x5d, 0x35, 0x58, 0x02, 0x8f, 0xec, 0x04, 0x69, 0xb0, 0xe8, 0x81, 0x0c, 0x61, 0x2f, 0xae,
	0xdb, 0x4e, 0x31, 0xbf, 0x94, 0x35, 0xbe, 0x1a, 0x06, 0x61, 0x23, 0xe5, 0xf6, 0x6b, 0xc0, 0x20,
	0xa6, 0x67, 0xa8, 0xed, 0x66, 0xc8, 0x6e, 0x54, 0x72, 0xc0, 0x68, 0xef, 0xef, 0x9e, 0xd6, 0x7a,
	0xa9, 0x62, 0x06, 0xb1, 0x5d, 0xb9, 0x83, 0x24, 0x4a, 0xa3, 0x7c, 0x51, 0xfc, 0x67, 0x90, 0x87,
	0x72, 0x07, 0x19, 0xab, 0x98, 0xa0, 0x09, 0x5d, 0x7d, 0xa5, 0xce, 0x46, 0x2b, 0x4c, 0xce, 0x5f,
	0xfd, 0xd3, 0x10, 0x3c, 0x06, 0x3e, 0xae, 0xd3, 0x97, 0xea, 0x24, 0x98, 0x6b, 0x47, 0x50, 0x49,
	0x3b, 0x24, 0xff, 0xd2, 0x28, 0x8f, 0x8b, 0xe3, 0x11, 0xde, 0xee, 0x99, 0xbe, 0x50, 0x73, 0x96,
	0x52, 0x7a, 0x4e, 0x66, 0x61, 0x97, 0xfd, 0x2c, 0xfb, 0x88, 0xd4, 0xf9, 0x81, 0xc3, 0xea, 0x77,
	0xb5, 0xf8, 0x46, 0xc9, 0x2a, 0x9d, 0xe5, 0x47, 0xab, 0x67, 0xf3, 0x07, 0x8f, 0x65, 0xa6, 0x2b,
	0x17, 0x3f, 0xbe, 0xcd, 0x3c, 0x7c, 0xce, 0xf5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0xd7,
	0x54, 0x4d, 0xb9, 0x01, 0x00, 0x00,
}