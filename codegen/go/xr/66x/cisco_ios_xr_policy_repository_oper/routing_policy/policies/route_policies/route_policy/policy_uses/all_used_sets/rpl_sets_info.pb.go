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
// source: rpl_sets_info.proto

package cisco_ios_xr_policy_repository_oper_routing_policy_policies_route_policies_route_policy_policy_uses_all_used_sets

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

type RplSetsInfo_KEYS struct {
	RoutePolicyName      string   `protobuf:"bytes,1,opt,name=route_policy_name,json=routePolicyName,proto3" json:"route_policy_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RplSetsInfo_KEYS) Reset()         { *m = RplSetsInfo_KEYS{} }
func (m *RplSetsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RplSetsInfo_KEYS) ProtoMessage()    {}
func (*RplSetsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b34d7b359685036, []int{0}
}

func (m *RplSetsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplSetsInfo_KEYS.Unmarshal(m, b)
}
func (m *RplSetsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplSetsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RplSetsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplSetsInfo_KEYS.Merge(m, src)
}
func (m *RplSetsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RplSetsInfo_KEYS.Size(m)
}
func (m *RplSetsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RplSetsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RplSetsInfo_KEYS proto.InternalMessageInfo

func (m *RplSetsInfo_KEYS) GetRoutePolicyName() string {
	if m != nil {
		return m.RoutePolicyName
	}
	return ""
}

type SetInfo struct {
	SetDomain            string   `protobuf:"bytes,1,opt,name=set_domain,json=setDomain,proto3" json:"set_domain,omitempty"`
	SetName              []string `protobuf:"bytes,2,rep,name=set_name,json=setName,proto3" json:"set_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetInfo) Reset()         { *m = SetInfo{} }
func (m *SetInfo) String() string { return proto.CompactTextString(m) }
func (*SetInfo) ProtoMessage()    {}
func (*SetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b34d7b359685036, []int{1}
}

func (m *SetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetInfo.Unmarshal(m, b)
}
func (m *SetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetInfo.Marshal(b, m, deterministic)
}
func (m *SetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetInfo.Merge(m, src)
}
func (m *SetInfo) XXX_Size() int {
	return xxx_messageInfo_SetInfo.Size(m)
}
func (m *SetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SetInfo proto.InternalMessageInfo

func (m *SetInfo) GetSetDomain() string {
	if m != nil {
		return m.SetDomain
	}
	return ""
}

func (m *SetInfo) GetSetName() []string {
	if m != nil {
		return m.SetName
	}
	return nil
}

type RplSetsInfo struct {
	Sets                 []*SetInfo `protobuf:"bytes,50,rep,name=sets,proto3" json:"sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RplSetsInfo) Reset()         { *m = RplSetsInfo{} }
func (m *RplSetsInfo) String() string { return proto.CompactTextString(m) }
func (*RplSetsInfo) ProtoMessage()    {}
func (*RplSetsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b34d7b359685036, []int{2}
}

func (m *RplSetsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplSetsInfo.Unmarshal(m, b)
}
func (m *RplSetsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplSetsInfo.Marshal(b, m, deterministic)
}
func (m *RplSetsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplSetsInfo.Merge(m, src)
}
func (m *RplSetsInfo) XXX_Size() int {
	return xxx_messageInfo_RplSetsInfo.Size(m)
}
func (m *RplSetsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RplSetsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RplSetsInfo proto.InternalMessageInfo

func (m *RplSetsInfo) GetSets() []*SetInfo {
	if m != nil {
		return m.Sets
	}
	return nil
}

func init() {
	proto.RegisterType((*RplSetsInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.policy_uses.all_used_sets.rpl_sets_info_KEYS")
	proto.RegisterType((*SetInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.policy_uses.all_used_sets.set_info")
	proto.RegisterType((*RplSetsInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.route_policies.route_policy.policy_uses.all_used_sets.rpl_sets_info")
}

func init() { proto.RegisterFile("rpl_sets_info.proto", fileDescriptor_4b34d7b359685036) }

var fileDescriptor_4b34d7b359685036 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x51, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x25, 0x9e, 0xa8, 0x19, 0x11, 0x71, 0x6d, 0xce, 0x42, 0x08, 0x57, 0x05, 0x8b, 0x2d, 0xce,
	0x1f, 0xb0, 0x38, 0x2b, 0x41, 0x24, 0x56, 0x56, 0x43, 0xcc, 0x8d, 0x32, 0x90, 0x64, 0xd6, 0x9d,
	0x3d, 0x30, 0xf8, 0x11, 0x7e, 0x83, 0x7f, 0x2a, 0x19, 0x23, 0x78, 0x60, 0x7d, 0xd5, 0xce, 0xbc,
	0x79, 0xef, 0xed, 0xdb, 0x1d, 0x38, 0x8f, 0xa1, 0x45, 0xa5, 0xa4, 0xc8, 0xfd, 0x8b, 0xf8, 0x10,
	0x25, 0x89, 0x7b, 0x6b, 0x58, 0x1b, 0x41, 0x16, 0xc5, 0xf7, 0x88, 0x41, 0x5a, 0x6e, 0x06, 0x8c,
	0x14, 0x44, 0x39, 0x49, 0x1c, 0x50, 0x02, 0x45, 0x1f, 0x65, 0x93, 0xb8, 0x7f, 0x9d, 0xc6, 0xde,
	0x0e, 0x26, 0x35, 0x9c, 0xf0, 0xbf, 0x76, 0x22, 0x0d, 0xb8, 0x51, 0x52, 0x5f, 0xb7, 0xed, 0x58,
	0xac, 0xed, 0xf6, 0xc5, 0x0d, 0xb8, 0xad, 0x24, 0x78, 0x77, 0xfb, 0xf4, 0xe8, 0xae, 0xe0, 0xec,
	0xaf, 0x1e, 0xfb, 0xba, 0xa3, 0x79, 0x56, 0x64, 0x65, 0x5e, 0x9d, 0xda, 0xe0, 0xc1, 0xf0, 0xfb,
	0xba, 0xa3, 0xc5, 0x0a, 0x8e, 0x94, 0x92, 0x89, 0xdd, 0x25, 0xc0, 0x58, 0xaf, 0xa5, 0xab, 0xb9,
	0x9f, 0x04, 0xb9, 0x52, 0x5a, 0x19, 0xe0, 0x2e, 0x7e, 0xa8, 0xe6, 0xb6, 0x57, 0xcc, 0xca, 0xbc,
	0x3a, 0x54, 0x4a, 0xe6, 0xf2, 0x95, 0xc1, 0xc9, 0x56, 0x10, 0xf7, 0x99, 0xc1, 0xfe, 0xd8, 0xcd,
	0x97, 0xc5, 0xac, 0x3c, 0x5e, 0x7e, 0xf8, 0x9d, 0x7f, 0x8e, 0xff, 0x7d, 0x57, 0x65, 0x41, 0x9e,
	0x0f, 0x6c, 0x4b, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x89, 0x07, 0x14, 0x83, 0xbc, 0x01,
	0x00, 0x00,
}
