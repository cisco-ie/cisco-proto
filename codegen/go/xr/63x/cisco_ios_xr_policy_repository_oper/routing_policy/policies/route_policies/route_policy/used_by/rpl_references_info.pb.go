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

package cisco_ios_xr_policy_repository_oper_routing_policy_policies_route_policies_route_policy_used_by

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
	RoutePolicyName      string   `protobuf:"bytes,1,opt,name=route_policy_name,json=routePolicyName,proto3" json:"route_policy_name,omitempty"`
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

func (m *RplReferencesInfo_KEYS) GetRoutePolicyName() string {
	if m != nil {
		return m.RoutePolicyName
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
	proto.RegisterType((*RplReferencesInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.used_by.rpl_references_info_KEYS")
	proto.RegisterType((*RefInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.used_by.ref_info")
	proto.RegisterType((*RplReferencesInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.used_by.rpl_references_info")
}

func init() { proto.RegisterFile("rpl_references_info.proto", fileDescriptor_ea22454312bfc5fd) }

var fileDescriptor_ea22454312bfc5fd = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xa9, 0x0b, 0xcb, 0x6e, 0x54, 0xc4, 0x08, 0x12, 0x6f, 0xa5, 0x5e, 0x8a, 0x87, 0x1c,
	0xd6, 0x57, 0x50, 0x2f, 0x82, 0x48, 0x3d, 0x79, 0x1a, 0xba, 0xd9, 0xa9, 0x04, 0xba, 0x9d, 0x30,
	0x93, 0x82, 0x79, 0x02, 0x1f, 0xc3, 0x57, 0x15, 0x43, 0xfd, 0x73, 0xe8, 0xc5, 0x83, 0xa7, 0x30,
	0xbf, 0x8f, 0x2f, 0xdf, 0x97, 0x8c, 0xba, 0xe0, 0xd0, 0x03, 0x63, 0x87, 0x8c, 0x83, 0x43, 0x01,
	0x3f, 0x74, 0x64, 0x03, 0x53, 0x24, 0x0d, 0xce, 0x8b, 0x23, 0xf0, 0x24, 0xf0, 0xca, 0x10, 0xa8,
	0xf7, 0x2e, 0x01, 0x63, 0x20, 0xf1, 0x91, 0x38, 0x01, 0x05, 0x64, 0xcb, 0x34, 0x46, 0x3f, 0xbc,
	0x4c, 0xb2, 0xcd, 0x87, 0x47, 0xc9, 0x1c, 0x61, 0x6e, 0x4c, 0x76, 0x14, 0xdc, 0xc1, 0x36, 0x55,
	0x77, 0xca, 0xcc, 0xa4, 0xc3, 0xfd, 0xed, 0xf3, 0x93, 0xbe, 0x52, 0xa7, 0xbf, 0x3d, 0x30, 0xb4,
	0x7b, 0x34, 0x45, 0x59, 0xd4, 0xeb, 0xe6, 0x24, 0x0b, 0x8f, 0x99, 0x3f, 0xb4, 0x7b, 0xac, 0x44,
	0xad, 0x18, 0xbb, 0x6c, 0xfe, 0x8b, 0x4f, 0x5f, 0xaa, 0xe3, 0x5c, 0x65, 0xe7, 0x19, 0x5d, 0xec,
	0x93, 0x39, 0x28, 0x8b, 0x7a, 0xd5, 0x1c, 0x7d, 0xc2, 0x9b, 0x89, 0xe9, 0x73, 0xb5, 0x94, 0xd8,
	0xc6, 0x51, 0xcc, 0x22, 0xdf, 0x32, 0x4d, 0xd5, 0x7b, 0xa1, 0xce, 0x66, 0xda, 0xeb, 0xb7, 0x42,
	0xad, 0xbf, 0x99, 0xd9, 0x94, 0x8b, 0xfa, 0x70, 0xe3, 0xed, 0x3f, 0x7f, 0xa5, 0xfd, 0x7a, 0x7f,
	0xf3, 0x93, 0xbd, 0x5d, 0xe6, 0x35, 0x5e, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x3a, 0x18,
	0xde, 0xe3, 0x01, 0x00, 0x00,
}
