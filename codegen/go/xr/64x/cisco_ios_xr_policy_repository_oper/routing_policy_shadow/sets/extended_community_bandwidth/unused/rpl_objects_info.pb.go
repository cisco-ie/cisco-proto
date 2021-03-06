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

package cisco_ios_xr_policy_repository_oper_routing_policy_shadow_sets_extended_community_bandwidth_unused

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
	proto.RegisterType((*RplObjectsInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.extended_community_bandwidth.unused.rpl_objects_info_KEYS")
	proto.RegisterType((*RplObjectsInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy_shadow.sets.extended_community_bandwidth.unused.rpl_objects_info")
}

func init() { proto.RegisterFile("rpl_objects_info.proto", fileDescriptor_5524f99561c441b5) }

var fileDescriptor_5524f99561c441b5 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0xce, 0x3d, 0x6b, 0xc3, 0x30,
	0x10, 0xc6, 0xf1, 0xa1, 0x60, 0xa8, 0xa6, 0x62, 0xa8, 0xdb, 0xb1, 0x78, 0x2a, 0x19, 0x34, 0x24,
	0x9f, 0x21, 0x53, 0xb6, 0x64, 0xca, 0x74, 0x58, 0xd2, 0x25, 0xbe, 0x60, 0xeb, 0x84, 0xee, 0x84,
	0xed, 0x6f, 0x1f, 0xc8, 0xcb, 0xe2, 0xf1, 0xcf, 0x6f, 0x78, 0x1e, 0xd3, 0xe4, 0x34, 0x00, 0xbb,
	0x1b, 0x7a, 0x15, 0xa0, 0x78, 0x61, 0x9b, 0x32, 0x2b, 0xd7, 0xce, 0x93, 0x78, 0x06, 0x62, 0x81,
	0x39, 0x43, 0xe2, 0x81, 0xfc, 0x02, 0x19, 0x13, 0x0b, 0x29, 0xe7, 0x05, 0x38, 0x61, 0xb6, 0x99,
	0x8b, 0x52, 0xbc, 0xbe, 0x59, 0xfa, 0x2e, 0xf0, 0x64, 0x05, 0x55, 0x2c, 0xce, 0x8a, 0x31, 0x60,
	0x00, 0xcf, 0xe3, 0x58, 0x22, 0xe9, 0x02, 0xae, 0x8b, 0x61, 0xa2, 0xa0, 0xbd, 0x2d, 0xb1, 0x08,
	0x86, 0xf6, 0xc7, 0x7c, 0xaf, 0xd7, 0xe1, 0xb0, 0x3f, 0x9f, 0xda, 0x8d, 0xf9, 0x5a, 0x43, 0xdd,
	0x98, 0xea, 0xd9, 0xbf, 0xdb, 0xbf, 0x8f, 0xff, 0xcf, 0xe3, 0xab, 0x5c, 0xf5, 0xf8, 0xbb, 0xbb,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x58, 0x0b, 0x3a, 0xc9, 0x00, 0x00, 0x00,
}
