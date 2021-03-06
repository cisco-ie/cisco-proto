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
// source: ifmib_ifname_ifindex_entry.proto

package cisco_ios_xr_snmp_ifmib_oper_augment_interface_mib_interfaces_interface

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

type IfmibIfnameIfindexEntry_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfmibIfnameIfindexEntry_KEYS) Reset()         { *m = IfmibIfnameIfindexEntry_KEYS{} }
func (m *IfmibIfnameIfindexEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*IfmibIfnameIfindexEntry_KEYS) ProtoMessage()    {}
func (*IfmibIfnameIfindexEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbaa5c30e96b5ba8, []int{0}
}

func (m *IfmibIfnameIfindexEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS.Unmarshal(m, b)
}
func (m *IfmibIfnameIfindexEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *IfmibIfnameIfindexEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS.Merge(m, src)
}
func (m *IfmibIfnameIfindexEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS.Size(m)
}
func (m *IfmibIfnameIfindexEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IfmibIfnameIfindexEntry_KEYS proto.InternalMessageInfo

func (m *IfmibIfnameIfindexEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IfmibIfnameIfindexEntry struct {
	IfIndex              uint32   `protobuf:"varint,50,opt,name=if_index,json=ifIndex,proto3" json:"if_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfmibIfnameIfindexEntry) Reset()         { *m = IfmibIfnameIfindexEntry{} }
func (m *IfmibIfnameIfindexEntry) String() string { return proto.CompactTextString(m) }
func (*IfmibIfnameIfindexEntry) ProtoMessage()    {}
func (*IfmibIfnameIfindexEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbaa5c30e96b5ba8, []int{1}
}

func (m *IfmibIfnameIfindexEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfmibIfnameIfindexEntry.Unmarshal(m, b)
}
func (m *IfmibIfnameIfindexEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfmibIfnameIfindexEntry.Marshal(b, m, deterministic)
}
func (m *IfmibIfnameIfindexEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfmibIfnameIfindexEntry.Merge(m, src)
}
func (m *IfmibIfnameIfindexEntry) XXX_Size() int {
	return xxx_messageInfo_IfmibIfnameIfindexEntry.Size(m)
}
func (m *IfmibIfnameIfindexEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IfmibIfnameIfindexEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IfmibIfnameIfindexEntry proto.InternalMessageInfo

func (m *IfmibIfnameIfindexEntry) GetIfIndex() uint32 {
	if m != nil {
		return m.IfIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*IfmibIfnameIfindexEntry_KEYS)(nil), "cisco_ios_xr_snmp_ifmib_oper.augment.interface_mib.interfaces.interface.ifmib_ifname_ifindex_entry_KEYS")
	proto.RegisterType((*IfmibIfnameIfindexEntry)(nil), "cisco_ios_xr_snmp_ifmib_oper.augment.interface_mib.interfaces.interface.ifmib_ifname_ifindex_entry")
}

func init() { proto.RegisterFile("ifmib_ifname_ifindex_entry.proto", fileDescriptor_fbaa5c30e96b5ba8) }

var fileDescriptor_fbaa5c30e96b5ba8 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4c, 0xcb, 0xcd,
	0x4c, 0x8a, 0xcf, 0x4c, 0xcb, 0x4b, 0xcc, 0x4d, 0x8d, 0xcf, 0x4c, 0xcb, 0xcc, 0x4b, 0x49, 0xad,
	0x88, 0x4f, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x4f, 0xce,
	0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x2f, 0xce, 0xcb, 0x2d, 0x88,
	0x87, 0xe8, 0xc9, 0x2f, 0x48, 0x2d, 0xd2, 0x4b, 0x2c, 0x4d, 0xcf, 0x4d, 0xcd, 0x2b, 0xd1, 0xcb,
	0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0x8d, 0xcf, 0xcd, 0x4c, 0x42, 0xf0, 0x8a, 0x11,
	0x4c, 0x25, 0x0f, 0x2e, 0x79, 0xdc, 0x96, 0xc5, 0x7b, 0xbb, 0x46, 0x06, 0x0b, 0xa9, 0x72, 0xf1,
	0x21, 0x0c, 0x02, 0x29, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x85, 0x8b, 0xfa, 0x25,
	0xe6, 0xa6, 0x2a, 0x99, 0x73, 0x49, 0xe1, 0x36, 0x49, 0x48, 0x92, 0x8b, 0x23, 0x33, 0x2d, 0x1e,
	0x2c, 0x22, 0x61, 0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0x9e, 0x99, 0xe6, 0x09, 0xe2, 0x26, 0xb1,
	0x81, 0xbd, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x45, 0x57, 0xe8, 0x77, 0xf6, 0x00, 0x00,
	0x00,
}
