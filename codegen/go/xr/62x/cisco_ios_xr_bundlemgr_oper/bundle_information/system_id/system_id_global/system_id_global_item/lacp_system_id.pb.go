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
// source: lacp_system_id.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_system_id_system_id_global_system_id_global_item

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

type LacpSystemId_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpSystemId_KEYS) Reset()         { *m = LacpSystemId_KEYS{} }
func (m *LacpSystemId_KEYS) String() string { return proto.CompactTextString(m) }
func (*LacpSystemId_KEYS) ProtoMessage()    {}
func (*LacpSystemId_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5d8843499bcf3c, []int{0}
}

func (m *LacpSystemId_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpSystemId_KEYS.Unmarshal(m, b)
}
func (m *LacpSystemId_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpSystemId_KEYS.Marshal(b, m, deterministic)
}
func (m *LacpSystemId_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpSystemId_KEYS.Merge(m, src)
}
func (m *LacpSystemId_KEYS) XXX_Size() int {
	return xxx_messageInfo_LacpSystemId_KEYS.Size(m)
}
func (m *LacpSystemId_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpSystemId_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LacpSystemId_KEYS proto.InternalMessageInfo

type EtherMacaddr_ struct {
	Macaddr              []uint32 `protobuf:"varint,1,rep,packed,name=macaddr,proto3" json:"macaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtherMacaddr_) Reset()         { *m = EtherMacaddr_{} }
func (m *EtherMacaddr_) String() string { return proto.CompactTextString(m) }
func (*EtherMacaddr_) ProtoMessage()    {}
func (*EtherMacaddr_) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5d8843499bcf3c, []int{1}
}

func (m *EtherMacaddr_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtherMacaddr_.Unmarshal(m, b)
}
func (m *EtherMacaddr_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtherMacaddr_.Marshal(b, m, deterministic)
}
func (m *EtherMacaddr_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtherMacaddr_.Merge(m, src)
}
func (m *EtherMacaddr_) XXX_Size() int {
	return xxx_messageInfo_EtherMacaddr_.Size(m)
}
func (m *EtherMacaddr_) XXX_DiscardUnknown() {
	xxx_messageInfo_EtherMacaddr_.DiscardUnknown(m)
}

var xxx_messageInfo_EtherMacaddr_ proto.InternalMessageInfo

func (m *EtherMacaddr_) GetMacaddr() []uint32 {
	if m != nil {
		return m.Macaddr
	}
	return nil
}

type BmSystemIdSt struct {
	SystemPrio           uint32         `protobuf:"varint,1,opt,name=system_prio,json=systemPrio,proto3" json:"system_prio,omitempty"`
	SystemMacAddr        *EtherMacaddr_ `protobuf:"bytes,2,opt,name=system_mac_addr,json=systemMacAddr,proto3" json:"system_mac_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BmSystemIdSt) Reset()         { *m = BmSystemIdSt{} }
func (m *BmSystemIdSt) String() string { return proto.CompactTextString(m) }
func (*BmSystemIdSt) ProtoMessage()    {}
func (*BmSystemIdSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5d8843499bcf3c, []int{2}
}

func (m *BmSystemIdSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmSystemIdSt.Unmarshal(m, b)
}
func (m *BmSystemIdSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmSystemIdSt.Marshal(b, m, deterministic)
}
func (m *BmSystemIdSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmSystemIdSt.Merge(m, src)
}
func (m *BmSystemIdSt) XXX_Size() int {
	return xxx_messageInfo_BmSystemIdSt.Size(m)
}
func (m *BmSystemIdSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmSystemIdSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmSystemIdSt proto.InternalMessageInfo

func (m *BmSystemIdSt) GetSystemPrio() uint32 {
	if m != nil {
		return m.SystemPrio
	}
	return 0
}

func (m *BmSystemIdSt) GetSystemMacAddr() *EtherMacaddr_ {
	if m != nil {
		return m.SystemMacAddr
	}
	return nil
}

type LacpSystemId struct {
	SystemId             *BmSystemIdSt `protobuf:"bytes,50,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	IccpGroupId          uint32        `protobuf:"varint,51,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	SystemPriority       uint32        `protobuf:"varint,52,opt,name=system_priority,json=systemPriority,proto3" json:"system_priority,omitempty"`
	SystemMacAddress     string        `protobuf:"bytes,53,opt,name=system_mac_address,json=systemMacAddress,proto3" json:"system_mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LacpSystemId) Reset()         { *m = LacpSystemId{} }
func (m *LacpSystemId) String() string { return proto.CompactTextString(m) }
func (*LacpSystemId) ProtoMessage()    {}
func (*LacpSystemId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e5d8843499bcf3c, []int{3}
}

func (m *LacpSystemId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpSystemId.Unmarshal(m, b)
}
func (m *LacpSystemId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpSystemId.Marshal(b, m, deterministic)
}
func (m *LacpSystemId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpSystemId.Merge(m, src)
}
func (m *LacpSystemId) XXX_Size() int {
	return xxx_messageInfo_LacpSystemId.Size(m)
}
func (m *LacpSystemId) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpSystemId.DiscardUnknown(m)
}

var xxx_messageInfo_LacpSystemId proto.InternalMessageInfo

func (m *LacpSystemId) GetSystemId() *BmSystemIdSt {
	if m != nil {
		return m.SystemId
	}
	return nil
}

func (m *LacpSystemId) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

func (m *LacpSystemId) GetSystemPriority() uint32 {
	if m != nil {
		return m.SystemPriority
	}
	return 0
}

func (m *LacpSystemId) GetSystemMacAddress() string {
	if m != nil {
		return m.SystemMacAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*LacpSystemId_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.system_id.system_id_global.system_id_global_item.lacp_system_id_KEYS")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.system_id.system_id_global.system_id_global_item.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.system_id.system_id_global.system_id_global_item.bm_system_id_st")
	proto.RegisterType((*LacpSystemId)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.system_id.system_id_global.system_id_global_item.lacp_system_id")
}

func init() { proto.RegisterFile("lacp_system_id.proto", fileDescriptor_9e5d8843499bcf3c) }

var fileDescriptor_9e5d8843499bcf3c = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x65, 0xfa, 0xc1, 0xa7, 0x9d, 0x92, 0x56, 0x46, 0x85, 0xec, 0x0c, 0xd9, 0x18, 0x44, 0xb2,
	0x68, 0xf5, 0x01, 0x5c, 0x88, 0x14, 0x11, 0x24, 0xae, 0x5c, 0x5d, 0x26, 0x93, 0xb1, 0x0e, 0x64,
	0x7a, 0x87, 0x99, 0x29, 0xd8, 0x87, 0x70, 0xed, 0xc2, 0x57, 0xf2, 0xa1, 0x64, 0x92, 0x9a, 0x9a,
	0xba, 0xee, 0xee, 0x9e, 0xc3, 0xcd, 0xcd, 0xf9, 0x19, 0x7a, 0x52, 0x73, 0x61, 0xc0, 0xad, 0x9d,
	0x97, 0x1a, 0x54, 0x95, 0x1b, 0x8b, 0x1e, 0x19, 0x08, 0xe5, 0x04, 0x82, 0x42, 0x07, 0x6f, 0x16,
	0xca, 0xd5, 0xb2, 0xaa, 0xa5, 0x5e, 0x58, 0x40, 0x23, 0x6d, 0xde, 0x42, 0x50, 0xcb, 0x17, 0xb4,
	0x9a, 0x7b, 0x85, 0xcb, 0x7c, 0xfb, 0x7d, 0x37, 0xc1, 0xa2, 0xc6, 0x92, 0xd7, 0x7f, 0x08, 0x50,
	0x5e, 0xea, 0xf4, 0x94, 0x1e, 0xf7, 0x7f, 0x0c, 0xf7, 0xb7, 0xcf, 0x4f, 0xe9, 0x05, 0x1d, 0x4b,
	0xff, 0x2a, 0x2d, 0x68, 0x2e, 0x78, 0x55, 0x59, 0x60, 0x31, 0x3d, 0xd8, 0xcc, 0x31, 0x49, 0xfe,
	0x65, 0x51, 0xf1, 0x03, 0xd3, 0x2f, 0x42, 0x27, 0xa5, 0xfe, 0x75, 0xc1, 0x79, 0x76, 0x46, 0x47,
	0x1b, 0x6c, 0xac, 0xc2, 0x98, 0x24, 0x24, 0x8b, 0x0a, 0xda, 0x52, 0x8f, 0x56, 0x21, 0xfb, 0x20,
	0x74, 0xb2, 0xd9, 0xd0, 0x5c, 0x40, 0x73, 0x77, 0x90, 0x90, 0x6c, 0x34, 0xc5, 0x7c, 0xcf, 0x9e,
	0xf3, 0xbe, 0xb3, 0x22, 0x6a, 0xb7, 0x1e, 0xb8, 0xb8, 0x09, 0x76, 0x3e, 0x07, 0x74, 0xdc, 0x8f,
	0x84, 0xbd, 0x13, 0x3a, 0xec, 0x50, 0x3c, 0x6d, 0x64, 0x9a, 0xbd, 0xcb, 0xdc, 0xc9, 0xb4, 0x38,
	0x6c, 0xd1, 0xbc, 0x62, 0x29, 0x8d, 0x94, 0x10, 0x06, 0x16, 0x16, 0x57, 0x26, 0x48, 0x9a, 0x35,
	0xf9, 0x8e, 0x02, 0x79, 0x17, 0xb8, 0x79, 0xc5, 0xce, 0xbb, 0x7c, 0x43, 0x03, 0x56, 0xf9, 0x75,
	0x7c, 0xd5, 0x6c, 0x8d, 0xb7, 0x2d, 0x04, 0x96, 0x5d, 0x52, 0xb6, 0x53, 0x84, 0x74, 0x2e, 0xbe,
	0x4e, 0x48, 0x36, 0x2c, 0x8e, 0x7a, 0xd1, 0x48, 0xe7, 0xca, 0xff, 0xcd, 0xbb, 0x9c, 0x7d, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0xde, 0x4b, 0xa4, 0xaf, 0x02, 0x00, 0x00,
}
