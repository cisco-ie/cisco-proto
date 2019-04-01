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
// source: iflist_ifl.proto

package cisco_ios_xr_l2vpn_oper_generic_interface_list_v2_standby_generic_interface_lists_generic_interface_list

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

type IflistIfl_KEYS struct {
	GenericInterfaceListName string   `protobuf:"bytes,1,opt,name=generic_interface_list_name,json=genericInterfaceListName,proto3" json:"generic_interface_list_name,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *IflistIfl_KEYS) Reset()         { *m = IflistIfl_KEYS{} }
func (m *IflistIfl_KEYS) String() string { return proto.CompactTextString(m) }
func (*IflistIfl_KEYS) ProtoMessage()    {}
func (*IflistIfl_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e15df188a3f009b6, []int{0}
}

func (m *IflistIfl_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IflistIfl_KEYS.Unmarshal(m, b)
}
func (m *IflistIfl_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IflistIfl_KEYS.Marshal(b, m, deterministic)
}
func (m *IflistIfl_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IflistIfl_KEYS.Merge(m, src)
}
func (m *IflistIfl_KEYS) XXX_Size() int {
	return xxx_messageInfo_IflistIfl_KEYS.Size(m)
}
func (m *IflistIfl_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IflistIfl_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IflistIfl_KEYS proto.InternalMessageInfo

func (m *IflistIfl_KEYS) GetGenericInterfaceListName() string {
	if m != nil {
		return m.GenericInterfaceListName
	}
	return ""
}

type IflistIf struct {
	InterfaceName            string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	PendingReplications      uint32   `protobuf:"varint,2,opt,name=pending_replications,json=pendingReplications,proto3" json:"pending_replications,omitempty"`
	NotSupportedReplications uint32   `protobuf:"varint,3,opt,name=not_supported_replications,json=notSupportedReplications,proto3" json:"not_supported_replications,omitempty"`
	IsFibDownloaded          bool     `protobuf:"varint,4,opt,name=is_fib_downloaded,json=isFibDownloaded,proto3" json:"is_fib_downloaded,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *IflistIf) Reset()         { *m = IflistIf{} }
func (m *IflistIf) String() string { return proto.CompactTextString(m) }
func (*IflistIf) ProtoMessage()    {}
func (*IflistIf) Descriptor() ([]byte, []int) {
	return fileDescriptor_e15df188a3f009b6, []int{1}
}

func (m *IflistIf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IflistIf.Unmarshal(m, b)
}
func (m *IflistIf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IflistIf.Marshal(b, m, deterministic)
}
func (m *IflistIf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IflistIf.Merge(m, src)
}
func (m *IflistIf) XXX_Size() int {
	return xxx_messageInfo_IflistIf.Size(m)
}
func (m *IflistIf) XXX_DiscardUnknown() {
	xxx_messageInfo_IflistIf.DiscardUnknown(m)
}

var xxx_messageInfo_IflistIf proto.InternalMessageInfo

func (m *IflistIf) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IflistIf) GetPendingReplications() uint32 {
	if m != nil {
		return m.PendingReplications
	}
	return 0
}

func (m *IflistIf) GetNotSupportedReplications() uint32 {
	if m != nil {
		return m.NotSupportedReplications
	}
	return 0
}

func (m *IflistIf) GetIsFibDownloaded() bool {
	if m != nil {
		return m.IsFibDownloaded
	}
	return false
}

type IflistIfl struct {
	InterfaceListName    string      `protobuf:"bytes,50,opt,name=interface_list_name,json=interfaceListName,proto3" json:"interface_list_name,omitempty"`
	InterfaceListId      uint32      `protobuf:"varint,51,opt,name=interface_list_id,json=interfaceListId,proto3" json:"interface_list_id,omitempty"`
	Interface            []*IflistIf `protobuf:"bytes,52,rep,name=interface,proto3" json:"interface,omitempty"`
	Items                uint32      `protobuf:"varint,53,opt,name=items,proto3" json:"items,omitempty"`
	IsProvisioned        bool        `protobuf:"varint,54,opt,name=is_provisioned,json=isProvisioned,proto3" json:"is_provisioned,omitempty"`
	IsFibDownloaded      bool        `protobuf:"varint,55,opt,name=is_fib_downloaded,json=isFibDownloaded,proto3" json:"is_fib_downloaded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IflistIfl) Reset()         { *m = IflistIfl{} }
func (m *IflistIfl) String() string { return proto.CompactTextString(m) }
func (*IflistIfl) ProtoMessage()    {}
func (*IflistIfl) Descriptor() ([]byte, []int) {
	return fileDescriptor_e15df188a3f009b6, []int{2}
}

func (m *IflistIfl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IflistIfl.Unmarshal(m, b)
}
func (m *IflistIfl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IflistIfl.Marshal(b, m, deterministic)
}
func (m *IflistIfl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IflistIfl.Merge(m, src)
}
func (m *IflistIfl) XXX_Size() int {
	return xxx_messageInfo_IflistIfl.Size(m)
}
func (m *IflistIfl) XXX_DiscardUnknown() {
	xxx_messageInfo_IflistIfl.DiscardUnknown(m)
}

var xxx_messageInfo_IflistIfl proto.InternalMessageInfo

func (m *IflistIfl) GetInterfaceListName() string {
	if m != nil {
		return m.InterfaceListName
	}
	return ""
}

func (m *IflistIfl) GetInterfaceListId() uint32 {
	if m != nil {
		return m.InterfaceListId
	}
	return 0
}

func (m *IflistIfl) GetInterface() []*IflistIf {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *IflistIfl) GetItems() uint32 {
	if m != nil {
		return m.Items
	}
	return 0
}

func (m *IflistIfl) GetIsProvisioned() bool {
	if m != nil {
		return m.IsProvisioned
	}
	return false
}

func (m *IflistIfl) GetIsFibDownloaded() bool {
	if m != nil {
		return m.IsFibDownloaded
	}
	return false
}

func init() {
	proto.RegisterType((*IflistIfl_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.standby.generic_interface_lists.generic_interface_list.iflist_ifl_KEYS")
	proto.RegisterType((*IflistIf)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.standby.generic_interface_lists.generic_interface_list.iflist_if")
	proto.RegisterType((*IflistIfl)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.standby.generic_interface_lists.generic_interface_list.iflist_ifl")
}

func init() { proto.RegisterFile("iflist_ifl.proto", fileDescriptor_e15df188a3f009b6) }

var fileDescriptor_e15df188a3f009b6 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0xc5, 0x49, 0xfb, 0xff, 0x8b, 0x1d, 0xa9, 0xb5, 0xd3, 0x2e, 0x06, 0xdd, 0x84, 0x82, 0x10,
	0x5c, 0x04, 0x4c, 0xfd, 0xd8, 0xe8, 0x4e, 0x85, 0xa2, 0x48, 0x49, 0x57, 0xae, 0x86, 0x24, 0x33,
	0xa9, 0x17, 0xd2, 0x99, 0x61, 0x66, 0xac, 0xfa, 0x18, 0x3e, 0x99, 0x4f, 0xe1, 0x7b, 0x48, 0xc7,
	0x36, 0x69, 0xa4, 0xdd, 0xba, 0xcc, 0x99, 0x73, 0xee, 0xc7, 0x2f, 0x17, 0x1d, 0x40, 0x5e, 0x80,
	0xb1, 0x14, 0xf2, 0x22, 0x54, 0x5a, 0x5a, 0x89, 0x9f, 0x33, 0x30, 0x99, 0xa4, 0x20, 0x0d, 0x7d,
	0xd3, 0xb4, 0x88, 0xe6, 0x4a, 0x50, 0xa9, 0xb8, 0x0e, 0xa7, 0x5c, 0x70, 0x0d, 0x19, 0x05, 0x61,
	0xb9, 0xce, 0x93, 0x8c, 0x53, 0x97, 0x9c, 0x47, 0xa1, 0xb1, 0x89, 0x60, 0xe9, 0xfb, 0x16, 0x87,
	0xd9, 0xa2, 0x0f, 0xc6, 0xa8, 0x53, 0x75, 0xa7, 0xf7, 0xb7, 0x4f, 0x13, 0x7c, 0x8d, 0x8e, 0xb6,
	0xb4, 0x11, 0xc9, 0x8c, 0x13, 0xcf, 0xf7, 0x82, 0x56, 0x4c, 0x96, 0x96, 0xd1, 0xca, 0xf1, 0x00,
	0xc6, 0x3e, 0x26, 0x33, 0x3e, 0xf8, 0xf4, 0x50, 0xab, 0x2c, 0x89, 0x8f, 0xd1, 0x7e, 0x55, 0x64,
	0x2d, 0xdf, 0x2e, 0xd5, 0x45, 0x08, 0x9f, 0xa2, 0xbe, 0xe2, 0x82, 0x81, 0x98, 0x52, 0xcd, 0x55,
	0x01, 0x59, 0x62, 0x41, 0x0a, 0x43, 0x1a, 0xbe, 0x17, 0xb4, 0xe3, 0xde, 0xf2, 0x2d, 0x5e, 0x7b,
	0xc2, 0x57, 0xe8, 0x50, 0x48, 0x4b, 0xcd, 0x8b, 0x52, 0x52, 0x5b, 0xce, 0xea, 0xc1, 0xa6, 0x0b,
	0x12, 0x21, 0xed, 0x64, 0x65, 0xa8, 0xa5, 0x4f, 0x50, 0x17, 0x0c, 0xcd, 0x21, 0xa5, 0x4c, 0xbe,
	0x8a, 0x42, 0x26, 0x8c, 0x33, 0xf2, 0xcf, 0xf7, 0x82, 0xdd, 0xb8, 0x03, 0xe6, 0x0e, 0xd2, 0x9b,
	0x52, 0x1e, 0x7c, 0x35, 0x10, 0xaa, 0x20, 0xe1, 0x10, 0xf5, 0x36, 0x71, 0x89, 0xdc, 0x5e, 0x5d,
	0xf8, 0x0d, 0xc4, 0xb5, 0xaa, 0xfb, 0x81, 0x91, 0xa1, 0x9b, 0xaf, 0x53, 0x73, 0x8f, 0x18, 0xfe,
	0x58, 0xc0, 0x5b, 0x69, 0xe4, 0xcc, 0x6f, 0x06, 0x7b, 0x91, 0x09, 0xff, 0xea, 0x1a, 0xc2, 0x72,
	0xcb, 0xb8, 0x9a, 0x02, 0xf7, 0xd1, 0x7f, 0xb0, 0x7c, 0x66, 0xc8, 0xb9, 0x9b, 0xf9, 0xe7, 0xc3,
	0xfd, 0x58, 0x43, 0x95, 0x96, 0x73, 0x30, 0x20, 0x05, 0x67, 0xe4, 0xc2, 0xd1, 0x6b, 0x83, 0x19,
	0x57, 0xe2, 0x66, 0xce, 0x97, 0x1b, 0x39, 0xa7, 0x3b, 0xee, 0xf8, 0x87, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x03, 0xa2, 0xe8, 0x38, 0x10, 0x03, 0x00, 0x00,
}