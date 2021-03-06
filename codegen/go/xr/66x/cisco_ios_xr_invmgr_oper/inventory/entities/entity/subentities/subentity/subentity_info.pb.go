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
// source: subentity_info.proto

package cisco_ios_xr_invmgr_oper_inventory_entities_entity_subentities_subentity

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

type SubentityInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubentityInfo_KEYS) Reset()         { *m = SubentityInfo_KEYS{} }
func (m *SubentityInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*SubentityInfo_KEYS) ProtoMessage()    {}
func (*SubentityInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_621af4b2ba3b22f9, []int{0}
}

func (m *SubentityInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubentityInfo_KEYS.Unmarshal(m, b)
}
func (m *SubentityInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubentityInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *SubentityInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubentityInfo_KEYS.Merge(m, src)
}
func (m *SubentityInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_SubentityInfo_KEYS.Size(m)
}
func (m *SubentityInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SubentityInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SubentityInfo_KEYS proto.InternalMessageInfo

func (m *SubentityInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SubentityInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

type SubentityInfo struct {
	SubentityName        string   `protobuf:"bytes,50,opt,name=subentity_name,json=subentityName,proto3" json:"subentity_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubentityInfo) Reset()         { *m = SubentityInfo{} }
func (m *SubentityInfo) String() string { return proto.CompactTextString(m) }
func (*SubentityInfo) ProtoMessage()    {}
func (*SubentityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_621af4b2ba3b22f9, []int{1}
}

func (m *SubentityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubentityInfo.Unmarshal(m, b)
}
func (m *SubentityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubentityInfo.Marshal(b, m, deterministic)
}
func (m *SubentityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubentityInfo.Merge(m, src)
}
func (m *SubentityInfo) XXX_Size() int {
	return xxx_messageInfo_SubentityInfo.Size(m)
}
func (m *SubentityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SubentityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SubentityInfo proto.InternalMessageInfo

func (m *SubentityInfo) GetSubentityName() string {
	if m != nil {
		return m.SubentityName
	}
	return ""
}

func init() {
	proto.RegisterType((*SubentityInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.subentities.subentity.subentity_info_KEYS")
	proto.RegisterType((*SubentityInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.subentities.subentity.subentity_info")
}

func init() { proto.RegisterFile("subentity_info.proto", fileDescriptor_621af4b2ba3b22f9) }

var fileDescriptor_621af4b2ba3b22f9 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0x4d, 0x4a,
	0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x8c, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xf2, 0x48, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf,
	0xcc, 0x2b, 0xcb, 0x4d, 0x2f, 0x8a, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0xcb, 0xcc, 0x2b, 0x4b, 0xcd,
	0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0x03, 0xeb, 0xca, 0x4c, 0x2d, 0x86, 0x30, 0x2a, 0xf5, 0x60, 0x06,
	0x81, 0x84, 0xe0, 0x86, 0x2a, 0x39, 0x70, 0x09, 0xa3, 0xda, 0x10, 0xef, 0xed, 0x1a, 0x19, 0x2c,
	0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66,
	0x0b, 0x89, 0x72, 0xb1, 0x81, 0xe8, 0x78, 0x43, 0x09, 0x26, 0xb0, 0x28, 0x2b, 0x88, 0x67, 0xa8,
	0x64, 0xce, 0xc5, 0x87, 0x6a, 0x82, 0x90, 0x2a, 0xb2, 0x08, 0xd8, 0x18, 0x23, 0xb0, 0x06, 0x5e,
	0xb8, 0xa8, 0x5f, 0x62, 0x6e, 0x6a, 0x12, 0x1b, 0xd8, 0x2f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0x14, 0x00, 0x91, 0xe3, 0x00, 0x00, 0x00,
}
