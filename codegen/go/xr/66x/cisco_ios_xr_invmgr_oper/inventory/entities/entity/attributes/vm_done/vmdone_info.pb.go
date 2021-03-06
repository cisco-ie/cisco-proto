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
// source: vmdone_info.proto

package cisco_ios_xr_invmgr_oper_inventory_entities_entity_attributes_vm_done

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

type VmdoneInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmdoneInfo_KEYS) Reset()         { *m = VmdoneInfo_KEYS{} }
func (m *VmdoneInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*VmdoneInfo_KEYS) ProtoMessage()    {}
func (*VmdoneInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0d8296b129b66, []int{0}
}

func (m *VmdoneInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmdoneInfo_KEYS.Unmarshal(m, b)
}
func (m *VmdoneInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmdoneInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *VmdoneInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmdoneInfo_KEYS.Merge(m, src)
}
func (m *VmdoneInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_VmdoneInfo_KEYS.Size(m)
}
func (m *VmdoneInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VmdoneInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VmdoneInfo_KEYS proto.InternalMessageInfo

func (m *VmdoneInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VmdoneInfo struct {
	VmNodeDone           uint32   `protobuf:"varint,50,opt,name=vm_node_done,json=vmNodeDone,proto3" json:"vm_node_done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmdoneInfo) Reset()         { *m = VmdoneInfo{} }
func (m *VmdoneInfo) String() string { return proto.CompactTextString(m) }
func (*VmdoneInfo) ProtoMessage()    {}
func (*VmdoneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_66d0d8296b129b66, []int{1}
}

func (m *VmdoneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmdoneInfo.Unmarshal(m, b)
}
func (m *VmdoneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmdoneInfo.Marshal(b, m, deterministic)
}
func (m *VmdoneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmdoneInfo.Merge(m, src)
}
func (m *VmdoneInfo) XXX_Size() int {
	return xxx_messageInfo_VmdoneInfo.Size(m)
}
func (m *VmdoneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VmdoneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VmdoneInfo proto.InternalMessageInfo

func (m *VmdoneInfo) GetVmNodeDone() uint32 {
	if m != nil {
		return m.VmNodeDone
	}
	return 0
}

func init() {
	proto.RegisterType((*VmdoneInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.attributes.vm_done.vmdone_info_KEYS")
	proto.RegisterType((*VmdoneInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.attributes.vm_done.vmdone_info")
}

func init() { proto.RegisterFile("vmdone_info.proto", fileDescriptor_66d0d8296b129b66) }

var fileDescriptor_66d0d8296b129b66 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x3d, 0x8b, 0xc2, 0x40,
	0x10, 0xc6, 0x71, 0x02, 0xc7, 0xc1, 0xed, 0x29, 0xe8, 0x56, 0x29, 0x43, 0x0a, 0x49, 0xb5, 0x82,
	0x7e, 0x05, 0x53, 0x09, 0x16, 0xb1, 0xb2, 0x1a, 0xf2, 0x32, 0xca, 0x14, 0x3b, 0x13, 0x76, 0xc7,
	0xc5, 0x7c, 0x7b, 0x31, 0x36, 0xe9, 0xfe, 0xc5, 0xc3, 0x8f, 0xc7, 0x6c, 0x93, 0x1f, 0x84, 0x11,
	0x88, 0xef, 0xe2, 0xc6, 0x20, 0x2a, 0xb6, 0xee, 0x29, 0xf6, 0x02, 0x24, 0x11, 0x5e, 0x01, 0x88,
	0x93, 0x7f, 0x04, 0x90, 0x11, 0x83, 0x23, 0x4e, 0xc8, 0x2a, 0x61, 0x72, 0xc8, 0x4a, 0x4a, 0x18,
	0xbf, 0x31, 0xb9, 0x56, 0x35, 0x50, 0xf7, 0x54, 0x8c, 0x2e, 0x79, 0xf8, 0x88, 0xe5, 0xce, 0x6c,
	0x16, 0x36, 0x9c, 0xeb, 0xdb, 0xd5, 0x5a, 0xf3, 0xc3, 0xad, 0xc7, 0x3c, 0x2b, 0xb2, 0xea, 0xaf,
	0x99, 0xbb, 0xdc, 0x9b, 0xff, 0xc5, 0xce, 0x16, 0x66, 0x95, 0x3c, 0xb0, 0x0c, 0x38, 0x33, 0xf9,
	0xa1, 0xc8, 0xaa, 0x75, 0x63, 0x92, 0xbf, 0xc8, 0x80, 0x27, 0x61, 0xec, 0x7e, 0xe7, 0x9b, 0xc7,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x63, 0xd5, 0x4c, 0xbb, 0x00, 0x00, 0x00,
}
