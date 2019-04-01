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
// source: mpls_io_show_intf_info.proto

package cisco_ios_xr_mpls_io_oper_mpls_ma_nodes_node_interfaces_interface

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

type MplsIoShowIntfInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsIoShowIntfInfo_KEYS) Reset()         { *m = MplsIoShowIntfInfo_KEYS{} }
func (m *MplsIoShowIntfInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsIoShowIntfInfo_KEYS) ProtoMessage()    {}
func (*MplsIoShowIntfInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_575f8f26baae8c2c, []int{0}
}

func (m *MplsIoShowIntfInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsIoShowIntfInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsIoShowIntfInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsIoShowIntfInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsIoShowIntfInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsIoShowIntfInfo_KEYS.Merge(m, src)
}
func (m *MplsIoShowIntfInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsIoShowIntfInfo_KEYS.Size(m)
}
func (m *MplsIoShowIntfInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsIoShowIntfInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsIoShowIntfInfo_KEYS proto.InternalMessageInfo

func (m *MplsIoShowIntfInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsIoShowIntfInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsIoShowIntfInfo struct {
	Mtu                  uint32   `protobuf:"varint,50,opt,name=mtu,proto3" json:"mtu,omitempty"`
	BkpLabelStackDepth   uint32   `protobuf:"varint,51,opt,name=bkp_label_stack_depth,json=bkpLabelStackDepth,proto3" json:"bkp_label_stack_depth,omitempty"`
	SrteLabelStackDepth  uint32   `protobuf:"varint,52,opt,name=srte_label_stack_depth,json=srteLabelStackDepth,proto3" json:"srte_label_stack_depth,omitempty"`
	PriLabelStackDepth   uint32   `protobuf:"varint,53,opt,name=pri_label_stack_depth,json=priLabelStackDepth,proto3" json:"pri_label_stack_depth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsIoShowIntfInfo) Reset()         { *m = MplsIoShowIntfInfo{} }
func (m *MplsIoShowIntfInfo) String() string { return proto.CompactTextString(m) }
func (*MplsIoShowIntfInfo) ProtoMessage()    {}
func (*MplsIoShowIntfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_575f8f26baae8c2c, []int{1}
}

func (m *MplsIoShowIntfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsIoShowIntfInfo.Unmarshal(m, b)
}
func (m *MplsIoShowIntfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsIoShowIntfInfo.Marshal(b, m, deterministic)
}
func (m *MplsIoShowIntfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsIoShowIntfInfo.Merge(m, src)
}
func (m *MplsIoShowIntfInfo) XXX_Size() int {
	return xxx_messageInfo_MplsIoShowIntfInfo.Size(m)
}
func (m *MplsIoShowIntfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsIoShowIntfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsIoShowIntfInfo proto.InternalMessageInfo

func (m *MplsIoShowIntfInfo) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *MplsIoShowIntfInfo) GetBkpLabelStackDepth() uint32 {
	if m != nil {
		return m.BkpLabelStackDepth
	}
	return 0
}

func (m *MplsIoShowIntfInfo) GetSrteLabelStackDepth() uint32 {
	if m != nil {
		return m.SrteLabelStackDepth
	}
	return 0
}

func (m *MplsIoShowIntfInfo) GetPriLabelStackDepth() uint32 {
	if m != nil {
		return m.PriLabelStackDepth
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsIoShowIntfInfo_KEYS)(nil), "cisco_ios_xr_mpls_io_oper.mpls_ma.nodes.node.interfaces.interface.mpls_io_show_intf_info_KEYS")
	proto.RegisterType((*MplsIoShowIntfInfo)(nil), "cisco_ios_xr_mpls_io_oper.mpls_ma.nodes.node.interfaces.interface.mpls_io_show_intf_info")
}

func init() { proto.RegisterFile("mpls_io_show_intf_info.proto", fileDescriptor_575f8f26baae8c2c) }

var fileDescriptor_575f8f26baae8c2c = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x05, 0xb1, 0x03, 0x15, 0x89, 0xb4, 0x2c, 0xd4, 0x43, 0x29, 0x08, 0x3d, 0x2d,
	0xd4, 0xd5, 0x07, 0x10, 0xf4, 0xa4, 0x78, 0x68, 0x4f, 0x9e, 0x86, 0xec, 0x76, 0x96, 0x86, 0xdd,
	0x64, 0x42, 0x12, 0xd1, 0x17, 0xf4, 0xbd, 0x24, 0x29, 0x56, 0x58, 0x72, 0x09, 0x93, 0xf9, 0xbf,
	0x2f, 0x3f, 0x04, 0x6e, 0xb5, 0x1d, 0x3c, 0x2a, 0x46, 0x7f, 0xe0, 0x2f, 0x54, 0x26, 0x74, 0xa8,
	0x4c, 0xc7, 0x95, 0x75, 0x1c, 0x58, 0x3c, 0xb5, 0xca, 0xb7, 0x8c, 0x8a, 0x3d, 0x7e, 0x3b, 0xfc,
	0x43, 0xd9, 0x92, 0xab, 0xd2, 0x45, 0xcb, 0xca, 0xf0, 0x9e, 0x7c, 0x3a, 0x2b, 0x65, 0x02, 0xb9,
	0x4e, 0xb6, 0xe4, 0xff, 0xc7, 0x95, 0x84, 0x45, 0xbe, 0x02, 0x5f, 0x5f, 0x3e, 0x76, 0x62, 0x01,
	0x93, 0xe8, 0xa2, 0x91, 0x9a, 0xca, 0x62, 0x59, 0xac, 0x27, 0xdb, 0xcb, 0xb8, 0x78, 0x97, 0x9a,
	0xc4, 0x1d, 0x5c, 0x9d, 0x1e, 0x3a, 0x12, 0x67, 0x89, 0x98, 0x9e, 0xb6, 0x11, 0x5b, 0xfd, 0x14,
	0x30, 0xcf, 0x77, 0x88, 0x6b, 0x38, 0xd7, 0xe1, 0xb3, 0xbc, 0x5f, 0x16, 0xeb, 0xe9, 0x36, 0x8e,
	0x62, 0x03, 0xb3, 0xa6, 0xb7, 0x38, 0xc8, 0x86, 0x06, 0xf4, 0x41, 0xb6, 0x3d, 0xee, 0xc9, 0x86,
	0x43, 0x59, 0x27, 0x46, 0x34, 0xbd, 0x7d, 0x8b, 0xd9, 0x2e, 0x46, 0xcf, 0x31, 0x11, 0x35, 0xcc,
	0xbd, 0x0b, 0x94, 0x71, 0x1e, 0x92, 0x73, 0x13, 0xd3, 0xb1, 0xb4, 0x81, 0x99, 0x75, 0x2a, 0xe3,
	0x3c, 0x1e, 0x7b, 0xac, 0x53, 0x23, 0xa5, 0xb9, 0x48, 0x9f, 0x5e, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xda, 0xc6, 0x94, 0xad, 0x94, 0x01, 0x00, 0x00,
}
