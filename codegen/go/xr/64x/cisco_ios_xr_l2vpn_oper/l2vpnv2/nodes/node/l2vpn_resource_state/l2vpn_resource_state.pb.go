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
// source: l2vpn_resource_state.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_l2vpn_resource_state

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

type L2VpnResourceState_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnResourceState_KEYS) Reset()         { *m = L2VpnResourceState_KEYS{} }
func (m *L2VpnResourceState_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnResourceState_KEYS) ProtoMessage()    {}
func (*L2VpnResourceState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c34676918da1de, []int{0}
}

func (m *L2VpnResourceState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Unmarshal(m, b)
}
func (m *L2VpnResourceState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnResourceState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnResourceState_KEYS.Merge(m, src)
}
func (m *L2VpnResourceState_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Size(m)
}
func (m *L2VpnResourceState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnResourceState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnResourceState_KEYS proto.InternalMessageInfo

func (m *L2VpnResourceState_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type L2VpnResourceState struct {
	ResourceOutOfMemoryState uint32   `protobuf:"varint,50,opt,name=resource_out_of_memory_state,json=resourceOutOfMemoryState,proto3" json:"resource_out_of_memory_state,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *L2VpnResourceState) Reset()         { *m = L2VpnResourceState{} }
func (m *L2VpnResourceState) String() string { return proto.CompactTextString(m) }
func (*L2VpnResourceState) ProtoMessage()    {}
func (*L2VpnResourceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c34676918da1de, []int{1}
}

func (m *L2VpnResourceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnResourceState.Unmarshal(m, b)
}
func (m *L2VpnResourceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnResourceState.Marshal(b, m, deterministic)
}
func (m *L2VpnResourceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnResourceState.Merge(m, src)
}
func (m *L2VpnResourceState) XXX_Size() int {
	return xxx_messageInfo_L2VpnResourceState.Size(m)
}
func (m *L2VpnResourceState) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnResourceState.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnResourceState proto.InternalMessageInfo

func (m *L2VpnResourceState) GetResourceOutOfMemoryState() uint32 {
	if m != nil {
		return m.ResourceOutOfMemoryState
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnResourceState_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_resource_state.l2vpn_resource_state_KEYS")
	proto.RegisterType((*L2VpnResourceState)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_resource_state.l2vpn_resource_state")
}

func init() { proto.RegisterFile("l2vpn_resource_state.proto", fileDescriptor_b2c34676918da1de) }

var fileDescriptor_b2c34676918da1de = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x31, 0x2a, 0x2b,
	0xc8, 0x8b, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4f, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc,
	0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x87, 0x28, 0xcc, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0x33, 0xcb, 0x8c,
	0xf4, 0xf2, 0xf2, 0x53, 0x52, 0x8b, 0xc1, 0xa4, 0x1e, 0x36, 0x63, 0x94, 0x4c, 0xb8, 0x24, 0xb1,
	0x89, 0xc7, 0x7b, 0xbb, 0x46, 0x06, 0x0b, 0x89, 0x73, 0xb1, 0x83, 0x74, 0xc6, 0x67, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x4a, 0x61, 0x5c, 0x22, 0xd8,
	0x74, 0x09, 0xd9, 0x71, 0xc9, 0xc0, 0x45, 0xf2, 0x4b, 0x4b, 0xe2, 0xf3, 0xd3, 0xe2, 0x73, 0x53,
	0x73, 0xf3, 0x8b, 0x2a, 0x21, 0xf2, 0x12, 0x46, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x12, 0x30, 0x35,
	0xfe, 0xa5, 0x25, 0xfe, 0x69, 0xbe, 0x60, 0x05, 0xc1, 0x20, 0xf9, 0x24, 0x36, 0xb0, 0xaf, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x78, 0x10, 0xfd, 0xf3, 0x00, 0x00, 0x00,
}