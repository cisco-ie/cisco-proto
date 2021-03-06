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
// source: l2vpn_generic_if_list.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_generic_interface_lists_generic_interface_list

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

type L2VpnGenericIfList_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	InterfaceListName    string   `protobuf:"bytes,2,opt,name=interface_list_name,json=interfaceListName,proto3" json:"interface_list_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnGenericIfList_KEYS) Reset()         { *m = L2VpnGenericIfList_KEYS{} }
func (m *L2VpnGenericIfList_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnGenericIfList_KEYS) ProtoMessage()    {}
func (*L2VpnGenericIfList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_541b174b6780f4ee, []int{0}
}

func (m *L2VpnGenericIfList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGenericIfList_KEYS.Unmarshal(m, b)
}
func (m *L2VpnGenericIfList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGenericIfList_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnGenericIfList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGenericIfList_KEYS.Merge(m, src)
}
func (m *L2VpnGenericIfList_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnGenericIfList_KEYS.Size(m)
}
func (m *L2VpnGenericIfList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGenericIfList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGenericIfList_KEYS proto.InternalMessageInfo

func (m *L2VpnGenericIfList_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnGenericIfList_KEYS) GetInterfaceListName() string {
	if m != nil {
		return m.InterfaceListName
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
	return fileDescriptor_541b174b6780f4ee, []int{1}
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

type L2VpnGenericIfList struct {
	InterfaceListNameXr  string      `protobuf:"bytes,50,opt,name=interface_list_name_xr,json=interfaceListNameXr,proto3" json:"interface_list_name_xr,omitempty"`
	InterfaceListId      uint32      `protobuf:"varint,51,opt,name=interface_list_id,json=interfaceListId,proto3" json:"interface_list_id,omitempty"`
	Interface            []*IflistIf `protobuf:"bytes,52,rep,name=interface,proto3" json:"interface,omitempty"`
	NumberOfInterfaces   uint32      `protobuf:"varint,53,opt,name=number_of_interfaces,json=numberOfInterfaces,proto3" json:"number_of_interfaces,omitempty"`
	Items                uint32      `protobuf:"varint,54,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *L2VpnGenericIfList) Reset()         { *m = L2VpnGenericIfList{} }
func (m *L2VpnGenericIfList) String() string { return proto.CompactTextString(m) }
func (*L2VpnGenericIfList) ProtoMessage()    {}
func (*L2VpnGenericIfList) Descriptor() ([]byte, []int) {
	return fileDescriptor_541b174b6780f4ee, []int{2}
}

func (m *L2VpnGenericIfList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGenericIfList.Unmarshal(m, b)
}
func (m *L2VpnGenericIfList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGenericIfList.Marshal(b, m, deterministic)
}
func (m *L2VpnGenericIfList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGenericIfList.Merge(m, src)
}
func (m *L2VpnGenericIfList) XXX_Size() int {
	return xxx_messageInfo_L2VpnGenericIfList.Size(m)
}
func (m *L2VpnGenericIfList) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGenericIfList.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGenericIfList proto.InternalMessageInfo

func (m *L2VpnGenericIfList) GetInterfaceListNameXr() string {
	if m != nil {
		return m.InterfaceListNameXr
	}
	return ""
}

func (m *L2VpnGenericIfList) GetInterfaceListId() uint32 {
	if m != nil {
		return m.InterfaceListId
	}
	return 0
}

func (m *L2VpnGenericIfList) GetInterface() []*IflistIf {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *L2VpnGenericIfList) GetNumberOfInterfaces() uint32 {
	if m != nil {
		return m.NumberOfInterfaces
	}
	return 0
}

func (m *L2VpnGenericIfList) GetItems() uint32 {
	if m != nil {
		return m.Items
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnGenericIfList_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.generic_interface_lists.generic_interface_list.l2vpn_generic_if_list_KEYS")
	proto.RegisterType((*IflistIf)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.generic_interface_lists.generic_interface_list.iflist_if")
	proto.RegisterType((*L2VpnGenericIfList)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.generic_interface_lists.generic_interface_list.l2vpn_generic_if_list")
}

func init() { proto.RegisterFile("l2vpn_generic_if_list.proto", fileDescriptor_541b174b6780f4ee) }

var fileDescriptor_541b174b6780f4ee = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0xa6, 0xd3, 0x45, 0xe6, 0x58, 0x36, 0x35, 0xcc, 0xcb, 0x18, 0x08, 0xc3, 0x43,
	0xd1, 0x4e, 0x3d, 0x79, 0x54, 0x61, 0x28, 0x0a, 0xdd, 0xc5, 0x9d, 0x42, 0xdb, 0xa4, 0xe3, 0x41,
	0x9b, 0x94, 0x24, 0x9b, 0x3b, 0xfb, 0xbf, 0x79, 0xf6, 0x5f, 0x92, 0xa6, 0xae, 0xb3, 0x38, 0x8f,
	0x5e, 0x4a, 0xf3, 0xbe, 0xef, 0xf7, 0xbd, 0xf4, 0xf5, 0xa1, 0xd3, 0xc4, 0x5b, 0x66, 0x82, 0xce,
	0xb9, 0xe0, 0x0a, 0x22, 0x0a, 0x31, 0x4d, 0x40, 0x1b, 0x37, 0x53, 0xd2, 0x48, 0x3c, 0x8b, 0x40,
	0x47, 0x92, 0x82, 0xd4, 0x74, 0xa5, 0x68, 0xe1, 0x94, 0x19, 0x57, 0xae, 0x7d, 0x5d, 0x7a, 0xae,
	0x90, 0x8c, 0x6b, 0xfb, 0x74, 0xcb, 0x04, 0x61, 0xb8, 0x8a, 0x83, 0x88, 0xdb, 0x20, 0xfd, 0x47,
	0x7d, 0xc8, 0x51, 0x7f, 0x6b, 0x67, 0xfa, 0x78, 0x3f, 0x9b, 0xe2, 0x13, 0xb4, 0x97, 0x87, 0x52,
	0x60, 0xc4, 0x19, 0x38, 0xa3, 0xa6, 0xdf, 0xc8, 0x8f, 0x13, 0x86, 0x5d, 0xd4, 0xad, 0x06, 0x51,
	0x11, 0xa4, 0x9c, 0xd4, 0xac, 0xa9, 0x53, 0x4a, 0x4f, 0xa0, 0xcd, 0x73, 0x90, 0xf2, 0xe1, 0xa7,
	0x83, 0x9a, 0x10, 0x5b, 0x23, 0xc4, 0xf8, 0x0c, 0x1d, 0x6e, 0x68, 0x0b, 0x16, 0xe9, 0xad, 0xb2,
	0x9a, 0x43, 0xf8, 0x12, 0xf5, 0x32, 0x2e, 0x18, 0x88, 0x39, 0x55, 0x3c, 0x4b, 0x20, 0x0a, 0x0c,
	0x48, 0xa1, 0x6d, 0x97, 0x96, 0xdf, 0xfd, 0xd6, 0xfc, 0x1f, 0x12, 0xbe, 0x45, 0x7d, 0x21, 0x0d,
	0xd5, 0x8b, 0x2c, 0x93, 0xca, 0x70, 0x56, 0x05, 0xeb, 0x16, 0x24, 0x42, 0x9a, 0xe9, 0xda, 0x50,
	0xa1, 0xcf, 0x51, 0x07, 0x34, 0x8d, 0x21, 0xa4, 0x4c, 0xbe, 0x89, 0x44, 0x06, 0x8c, 0x33, 0xb2,
	0x33, 0x70, 0x46, 0xfb, 0x7e, 0x1b, 0xf4, 0x03, 0x84, 0x77, 0x65, 0x79, 0xf8, 0x51, 0x43, 0x47,
	0x5b, 0x27, 0x87, 0xc7, 0xe8, 0x78, 0xcb, 0x6c, 0xe8, 0x4a, 0x11, 0xcf, 0x7e, 0x65, 0xf7, 0xd7,
	0x78, 0x5e, 0x95, 0x6d, 0x5d, 0x85, 0x80, 0x91, 0xb1, 0xbd, 0x6f, 0xbb, 0xe2, 0x9f, 0x30, 0xfc,
	0x9e, 0x0f, 0x73, 0x5d, 0x23, 0x57, 0x83, 0xfa, 0xe8, 0xc0, 0x63, 0xee, 0xbf, 0xed, 0x88, 0x5b,
	0xfe, 0x38, 0x7f, 0xd3, 0x16, 0x5f, 0xa0, 0x9e, 0x58, 0xa4, 0x21, 0x57, 0x54, 0xc6, 0x1b, 0x40,
	0x93, 0x6b, 0x7b, 0x67, 0x5c, 0x68, 0x2f, 0xf1, 0xa4, 0x54, 0x70, 0x0f, 0xed, 0x82, 0xe1, 0xa9,
	0x26, 0x37, 0xd6, 0x52, 0x1c, 0xc2, 0x86, 0x5d, 0xf1, 0xf1, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x66, 0x7c, 0x4d, 0x01, 0x03, 0x00, 0x00,
}
