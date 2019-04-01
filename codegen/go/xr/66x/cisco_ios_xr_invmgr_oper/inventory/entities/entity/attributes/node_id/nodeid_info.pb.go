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
// source: nodeid_info.proto

package cisco_ios_xr_invmgr_oper_inventory_entities_entity_attributes_node_id

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

type NodeidInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeidInfo_KEYS) Reset()         { *m = NodeidInfo_KEYS{} }
func (m *NodeidInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*NodeidInfo_KEYS) ProtoMessage()    {}
func (*NodeidInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_663c418946bbd2ad, []int{0}
}

func (m *NodeidInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeidInfo_KEYS.Unmarshal(m, b)
}
func (m *NodeidInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeidInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *NodeidInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeidInfo_KEYS.Merge(m, src)
}
func (m *NodeidInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_NodeidInfo_KEYS.Size(m)
}
func (m *NodeidInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeidInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_NodeidInfo_KEYS proto.InternalMessageInfo

func (m *NodeidInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NodeidInfo struct {
	NodeId               uint32   `protobuf:"varint,50,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeidInfo) Reset()         { *m = NodeidInfo{} }
func (m *NodeidInfo) String() string { return proto.CompactTextString(m) }
func (*NodeidInfo) ProtoMessage()    {}
func (*NodeidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_663c418946bbd2ad, []int{1}
}

func (m *NodeidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeidInfo.Unmarshal(m, b)
}
func (m *NodeidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeidInfo.Marshal(b, m, deterministic)
}
func (m *NodeidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeidInfo.Merge(m, src)
}
func (m *NodeidInfo) XXX_Size() int {
	return xxx_messageInfo_NodeidInfo.Size(m)
}
func (m *NodeidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeidInfo proto.InternalMessageInfo

func (m *NodeidInfo) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func init() {
	proto.RegisterType((*NodeidInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.attributes.node_id.nodeid_info_KEYS")
	proto.RegisterType((*NodeidInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.entities.entity.attributes.node_id.nodeid_info")
}

func init() { proto.RegisterFile("nodeid_info.proto", fileDescriptor_663c418946bbd2ad) }

var fileDescriptor_663c418946bbd2ad = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0xcb, 0x4f, 0x49,
	0xcd, 0x4c, 0x89, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x72, 0x4d,
	0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0xcc, 0x2b, 0xcb,
	0x4d, 0x2f, 0x8a, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0xcb, 0xcc, 0x2b, 0x4b, 0xcd, 0x2b, 0xc9, 0x2f,
	0xaa, 0xd4, 0x4b, 0xcd, 0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0x86, 0x30, 0x2a, 0xf5, 0x12, 0x4b,
	0x4a, 0x8a, 0x32, 0x93, 0x4a, 0x4b, 0x52, 0x8b, 0xf5, 0x40, 0x06, 0xc6, 0x67, 0xa6, 0x28, 0xa9,
	0x71, 0x09, 0x20, 0x99, 0x1d, 0xef, 0xed, 0x1a, 0x19, 0x2c, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98,
	0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0xa9, 0x71, 0x71, 0x23, 0xa9,
	0x13, 0x12, 0xe7, 0x62, 0x87, 0x9a, 0x20, 0x61, 0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0x06, 0xe2,
	0x7a, 0xa6, 0x24, 0xb1, 0x81, 0x5d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0c, 0x3e,
	0x48, 0xb2, 0x00, 0x00, 0x00,
}