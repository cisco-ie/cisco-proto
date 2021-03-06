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
// source: rpl_objects_info.proto

package cisco_ios_xr_policy_repository_oper_routing_policy_shadow_sets_large_community_active

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

type RplObjectsInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RplObjectsInfo_KEYS) Reset()         { *m = RplObjectsInfo_KEYS{} }
func (m *RplObjectsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RplObjectsInfo_KEYS) ProtoMessage()    {}
func (*RplObjectsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5524f99561c441b5, []int{0}
}

func (m *RplObjectsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplObjectsInfo_KEYS.Unmarshal(m, b)
}
func (m *RplObjectsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplObjectsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RplObjectsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplObjectsInfo_KEYS.Merge(m, src)
}
func (m *RplObjectsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RplObjectsInfo_KEYS.Size(m)
}
func (m *RplObjectsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RplObjectsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RplObjectsInfo_KEYS proto.InternalMessageInfo

type RplObjectsInfo struct {
	Object               []string `protobuf:"bytes,50,rep,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RplObjectsInfo) Reset()         { *m = RplObjectsInfo{} }
func (m *RplObjectsInfo) String() string { return proto.CompactTextString(m) }
func (*RplObjectsInfo) ProtoMessage()    {}
func (*RplObjectsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5524f99561c441b5, []int{1}
}

func (m *RplObjectsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RplObjectsInfo.Unmarshal(m, b)
}
func (m *RplObjectsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RplObjectsInfo.Marshal(b, m, deterministic)
}
func (m *RplObjectsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RplObjectsInfo.Merge(m, src)
}
func (m *RplObjectsInfo) XXX_Size() int {
	return xxx_messageInfo_RplObjectsInfo.Size(m)
}
func (m *RplObjectsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RplObjectsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RplObjectsInfo proto.InternalMessageInfo

func (m *RplObjectsInfo) GetObject() []string {
	if m != nil {
		return m.Object
	}
	return nil
}

func init() {
	proto.RegisterType((*RplObjectsInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.large_community.active.rpl_objects_info_KEYS")
	proto.RegisterType((*RplObjectsInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.large_community.active.rpl_objects_info")
}

func init() { proto.RegisterFile("rpl_objects_info.proto", fileDescriptor_5524f99561c441b5) }

var fileDescriptor_5524f99561c441b5 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xce, 0xbf, 0x8b, 0xc2, 0x40,
	0x10, 0xc5, 0xf1, 0xe2, 0x20, 0x70, 0x5b, 0x1d, 0x81, 0xcb, 0x5d, 0x29, 0xa9, 0xc4, 0x62, 0x0b,
	0xfd, 0x1b, 0xac, 0xec, 0x14, 0x0b, 0xab, 0x21, 0x59, 0xd7, 0x38, 0x92, 0xe4, 0x2d, 0x33, 0x13,
	0x35, 0xff, 0xbd, 0xe0, 0x8f, 0x26, 0xe5, 0x97, 0x4f, 0xf1, 0x9e, 0x2b, 0x24, 0xb5, 0x84, 0xfa,
	0x12, 0x83, 0x29, 0x71, 0x7f, 0x82, 0x4f, 0x02, 0x43, 0xbe, 0x0f, 0xac, 0x01, 0xc4, 0x50, 0xba,
	0x0b, 0x25, 0xb4, 0x1c, 0x46, 0x92, 0x98, 0xa0, 0x6c, 0x90, 0x91, 0x90, 0xa2, 0x78, 0xc1, 0x60,
	0xdc, 0x37, 0x1f, 0xd6, 0x73, 0x75, 0xc4, 0xcd, 0x6b, 0x34, 0xf5, 0x6d, 0x25, 0x4d, 0xa4, 0x80,
	0xae, 0x1b, 0x7a, 0xb6, 0xd1, 0x57, 0xc1, 0xf8, 0x1a, 0xcb, 0x3f, 0xf7, 0x3b, 0x1d, 0xa4, 0xcd,
	0xfa, 0xb0, 0x2b, 0x17, 0xee, 0x67, 0x0a, 0x79, 0xe1, 0xb2, 0x57, 0xff, 0x2f, 0x67, 0x5f, 0xf3,
	0xef, 0xed, 0xbb, 0xea, 0xec, 0x79, 0x71, 0xf5, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xc8, 0xe0,
	0x65, 0xbc, 0x00, 0x00, 0x00,
}
