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

package cisco_ios_xr_policy_repository_oper_routing_policy_sets_extended_community_cost_unused

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
	proto.RegisterType((*RplObjectsInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.sets.extended_community_cost.unused.rpl_objects_info_KEYS")
	proto.RegisterType((*RplObjectsInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.sets.extended_community_cost.unused.rpl_objects_info")
}

func init() { proto.RegisterFile("rpl_objects_info.proto", fileDescriptor_5524f99561c441b5) }

var fileDescriptor_5524f99561c441b5 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xce, 0xb1, 0xaa, 0xc2, 0x40,
	0x10, 0x85, 0xe1, 0xe2, 0x42, 0xe0, 0x6e, 0x25, 0x01, 0xa3, 0xa5, 0xa4, 0x12, 0x8b, 0x2d, 0xf4,
	0x19, 0xac, 0xec, 0x14, 0x04, 0xab, 0x81, 0x6c, 0x46, 0x19, 0x49, 0x76, 0x96, 0x99, 0x59, 0xc8,
	0xbe, 0xbd, 0xa0, 0xa9, 0x52, 0xfe, 0x7c, 0xc5, 0x39, 0xae, 0x91, 0x34, 0x00, 0x77, 0x6f, 0x0c,
	0xa6, 0x40, 0xf1, 0xc9, 0x3e, 0x09, 0x1b, 0xd7, 0xf7, 0x40, 0x1a, 0x18, 0x88, 0x15, 0x26, 0x81,
	0xc4, 0x03, 0x85, 0x02, 0x82, 0x89, 0x95, 0x8c, 0xa5, 0x00, 0x27, 0x14, 0x2f, 0x9c, 0x8d, 0xe2,
	0x6b, 0x66, 0xaf, 0x68, 0xea, 0x71, 0x32, 0x8c, 0x3d, 0xf6, 0x10, 0x78, 0x1c, 0x73, 0x24, 0x2b,
	0x10, 0x58, 0xcd, 0xe7, 0x98, 0x15, 0xfb, 0x76, 0xe3, 0xd6, 0xcb, 0x45, 0xb8, 0x9c, 0x1f, 0xb7,
	0xf6, 0xe0, 0x56, 0x4b, 0xa8, 0x1b, 0x57, 0xfd, 0x7a, 0x7b, 0xdc, 0xfd, 0xed, 0xff, 0xaf, 0x73,
	0x75, 0xd5, 0xf7, 0xe3, 0xe9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x96, 0x69, 0x5c, 0xbd, 0x00,
	0x00, 0x00,
}
