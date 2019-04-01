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

package cisco_ios_xr_policy_repository_oper_routing_policy_policies_inactive

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
	proto.RegisterType((*RplObjectsInfo_KEYS)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.inactive.rpl_objects_info_KEYS")
	proto.RegisterType((*RplObjectsInfo)(nil), "cisco_ios_xr_policy_repository_oper.routing_policy.policies.inactive.rpl_objects_info")
}

func init() { proto.RegisterFile("rpl_objects_info.proto", fileDescriptor_5524f99561c441b5) }

var fileDescriptor_5524f99561c441b5 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2a, 0xc8, 0x89,
	0xcf, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0x29, 0x8e, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x72, 0x49, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28,
	0x8a, 0x2f, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x8c, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9,
	0x2f, 0xaa, 0x8c, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x2b, 0xca, 0x2f, 0x2d, 0xc9, 0xcc, 0x4b, 0x87,
	0x4a, 0xeb, 0x81, 0xa9, 0xcc, 0xd4, 0x62, 0xbd, 0xcc, 0xbc, 0xc4, 0xe4, 0x92, 0xcc, 0xb2, 0x54,
	0x25, 0x71, 0x2e, 0x51, 0x74, 0xf3, 0xe3, 0xbd, 0x5d, 0x23, 0x83, 0x95, 0xb4, 0xb8, 0x04, 0xd0,
	0x25, 0x84, 0xc4, 0xb8, 0xd8, 0x20, 0x7c, 0x09, 0x23, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x28, 0x2f,
	0x89, 0x0d, 0xec, 0x22, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0xe0, 0x14, 0x28, 0xab,
	0x00, 0x00, 0x00,
}
