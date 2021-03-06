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

package cisco_ios_xr_l2vpn_oper_generic_interface_list_v2_nodes_node_generic_interface_lists_generic_interface_list

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
	NodeId                   string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	GenericInterfaceListName string   `protobuf:"bytes,2,opt,name=generic_interface_list_name,json=genericInterfaceListName,proto3" json:"generic_interface_list_name,omitempty"`
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

func (m *IflistIfl_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

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
	proto.RegisterType((*IflistIfl_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.nodes.node.generic_interface_lists.generic_interface_list.iflist_ifl_KEYS")
	proto.RegisterType((*IflistIf)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.nodes.node.generic_interface_lists.generic_interface_list.iflist_if")
	proto.RegisterType((*IflistIfl)(nil), "cisco_ios_xr_l2vpn_oper.generic_interface_list_v2.nodes.node.generic_interface_lists.generic_interface_list.iflist_ifl")
}

func init() { proto.RegisterFile("iflist_ifl.proto", fileDescriptor_e15df188a3f009b6) }

var fileDescriptor_e15df188a3f009b6 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0xa5, 0x9b, 0x4e, 0x17, 0x99, 0x73, 0xd9, 0xc0, 0xa0, 0x97, 0x31, 0x10, 0x86, 0x87, 0x82,
	0x9d, 0x7f, 0x2e, 0x7a, 0x53, 0x61, 0x28, 0x22, 0xdd, 0xc9, 0x53, 0xe8, 0x9a, 0x74, 0xfc, 0xb0,
	0x4b, 0x42, 0x12, 0xab, 0x1f, 0xc4, 0xef, 0xe5, 0xd7, 0xf0, 0x63, 0x48, 0xb3, 0xad, 0xb5, 0xb2,
	0x5d, 0xbd, 0x14, 0xf2, 0xf2, 0xde, 0xef, 0xbd, 0xdf, 0x6b, 0x8b, 0x0e, 0x20, 0x49, 0xc1, 0x58,
	0x0a, 0x49, 0xea, 0x2b, 0x2d, 0xad, 0xc4, 0xaf, 0x31, 0x98, 0x58, 0x52, 0x90, 0x86, 0x7e, 0x68,
	0x9a, 0x06, 0x99, 0x12, 0x54, 0x2a, 0xae, 0xfd, 0x19, 0x17, 0x5c, 0x43, 0x4c, 0x41, 0x58, 0xae,
	0x93, 0x28, 0xe6, 0xd4, 0x29, 0xb3, 0xc0, 0x17, 0x92, 0x71, 0xe3, 0x9e, 0x1b, 0x48, 0x66, 0x03,
	0x3e, 0x00, 0xd4, 0x2e, 0x03, 0xd0, 0x87, 0xbb, 0x97, 0x09, 0x3e, 0x44, 0x3b, 0xf9, 0x24, 0x0a,
	0x8c, 0x78, 0x7d, 0x6f, 0xd8, 0x0c, 0x1b, 0xf9, 0x71, 0xcc, 0xf0, 0x0d, 0x3a, 0xde, 0x10, 0x41,
	0x44, 0x73, 0x4e, 0x6a, 0x8e, 0x4c, 0x96, 0x94, 0xf1, 0x8a, 0xf1, 0x08, 0xc6, 0x3e, 0x45, 0x73,
	0x3e, 0xf8, 0xf2, 0x50, 0xb3, 0xf0, 0xc2, 0x27, 0x68, 0xbf, 0x1c, 0xe2, 0xf4, 0x0b, 0xb3, 0x56,
	0x81, 0xe6, 0x22, 0x7c, 0x86, 0x7a, 0x8a, 0x0b, 0x06, 0x62, 0x46, 0x35, 0x57, 0x29, 0xc4, 0x91,
	0x05, 0x29, 0x8c, 0x33, 0x6b, 0x85, 0xdd, 0xe5, 0x5d, 0xf8, 0xeb, 0x0a, 0x5f, 0xa3, 0x23, 0x21,
	0x2d, 0x35, 0x6f, 0x4a, 0x49, 0x6d, 0x39, 0xab, 0x0a, 0xeb, 0x4e, 0x48, 0x84, 0xb4, 0x93, 0x15,
	0xa1, 0xa2, 0x3e, 0x45, 0x1d, 0x30, 0x34, 0x81, 0x29, 0x65, 0xf2, 0x5d, 0xa4, 0x32, 0x62, 0x9c,
	0x91, 0xad, 0xbe, 0x37, 0xdc, 0x0d, 0xdb, 0x60, 0xee, 0x61, 0x7a, 0x5b, 0xc0, 0x83, 0xef, 0x1a,
	0x42, 0x65, 0x7b, 0xd8, 0x47, 0xdd, 0x75, 0xbd, 0x04, 0x6e, 0xaf, 0x0e, 0xfc, 0x2d, 0xc4, 0x59,
	0x55, 0xf9, 0xc0, 0xc8, 0xc8, 0xe5, 0x6b, 0x57, 0xd8, 0x63, 0x86, 0x3f, 0xf3, 0xf2, 0x56, 0x18,
	0x39, 0xef, 0xd7, 0x87, 0x7b, 0x41, 0xe6, 0xff, 0xe3, 0x97, 0xe2, 0x17, 0x8b, 0x86, 0x65, 0x10,
	0xdc, 0x43, 0xdb, 0x60, 0xf9, 0xdc, 0x90, 0x0b, 0x17, 0x7b, 0x71, 0x70, 0xef, 0xd6, 0x50, 0xa5,
	0x65, 0x06, 0x06, 0xa4, 0xe0, 0x8c, 0x5c, 0xba, 0x02, 0x5b, 0x60, 0x9e, 0x4b, 0x70, 0x7d, 0xd5,
	0x57, 0x6b, 0xab, 0x9e, 0x36, 0xdc, 0xbf, 0x31, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x77, 0x69,
	0x62, 0xad, 0x2f, 0x03, 0x00, 0x00,
}
