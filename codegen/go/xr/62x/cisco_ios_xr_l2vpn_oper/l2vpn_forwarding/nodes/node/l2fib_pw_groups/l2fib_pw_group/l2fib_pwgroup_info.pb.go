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
// source: l2fib_pwgroup_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_pw_groups_l2fib_pw_group

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

type L2FibPwgroupInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PeerAddress          string   `protobuf:"bytes,2,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	GroupId              uint32   `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	TunnelIfHandle       uint32   `protobuf:"varint,4,opt,name=tunnel_if_handle,json=tunnelIfHandle,proto3" json:"tunnel_if_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPwgroupInfo_KEYS) Reset()         { *m = L2FibPwgroupInfo_KEYS{} }
func (m *L2FibPwgroupInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibPwgroupInfo_KEYS) ProtoMessage()    {}
func (*L2FibPwgroupInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_72bcb4bffcb6332e, []int{0}
}

func (m *L2FibPwgroupInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwgroupInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibPwgroupInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwgroupInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibPwgroupInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwgroupInfo_KEYS.Merge(m, src)
}
func (m *L2FibPwgroupInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibPwgroupInfo_KEYS.Size(m)
}
func (m *L2FibPwgroupInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwgroupInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwgroupInfo_KEYS proto.InternalMessageInfo

func (m *L2FibPwgroupInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibPwgroupInfo_KEYS) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *L2FibPwgroupInfo_KEYS) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *L2FibPwgroupInfo_KEYS) GetTunnelIfHandle() uint32 {
	if m != nil {
		return m.TunnelIfHandle
	}
	return 0
}

type L2FibPwgroupInfo struct {
	GroupState           string   `protobuf:"bytes,50,opt,name=group_state,json=groupState,proto3" json:"group_state,omitempty"`
	PwListCount          uint32   `protobuf:"varint,51,opt,name=pw_list_count,json=pwListCount,proto3" json:"pw_list_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPwgroupInfo) Reset()         { *m = L2FibPwgroupInfo{} }
func (m *L2FibPwgroupInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPwgroupInfo) ProtoMessage()    {}
func (*L2FibPwgroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_72bcb4bffcb6332e, []int{1}
}

func (m *L2FibPwgroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwgroupInfo.Unmarshal(m, b)
}
func (m *L2FibPwgroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwgroupInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPwgroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwgroupInfo.Merge(m, src)
}
func (m *L2FibPwgroupInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPwgroupInfo.Size(m)
}
func (m *L2FibPwgroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwgroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwgroupInfo proto.InternalMessageInfo

func (m *L2FibPwgroupInfo) GetGroupState() string {
	if m != nil {
		return m.GroupState
	}
	return ""
}

func (m *L2FibPwgroupInfo) GetPwListCount() uint32 {
	if m != nil {
		return m.PwListCount
	}
	return 0
}

func init() {
	proto.RegisterType((*L2FibPwgroupInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_pw_groups.l2fib_pw_group.l2fib_pwgroup_info_KEYS")
	proto.RegisterType((*L2FibPwgroupInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_pw_groups.l2fib_pw_group.l2fib_pwgroup_info")
}

func init() { proto.RegisterFile("l2fib_pwgroup_info.proto", fileDescriptor_72bcb4bffcb6332e) }

var fileDescriptor_72bcb4bffcb6332e = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x15, 0x40, 0x2d, 0x5c, 0x28, 0x42, 0x5e, 0x6a, 0x26, 0x4a, 0xa6, 0x4c, 0x19, 0xd2,
	0x2f, 0x40, 0x08, 0x89, 0x08, 0xa6, 0x74, 0xea, 0x74, 0x4a, 0x63, 0xa7, 0x58, 0x8a, 0x7c, 0x96,
	0xed, 0x10, 0xbe, 0x85, 0xaf, 0x45, 0xb9, 0xc0, 0xd2, 0x2e, 0x96, 0xdf, 0xb3, 0xf5, 0xce, 0x32,
	0xc8, 0xbe, 0xec, 0xcc, 0x01, 0xdd, 0x78, 0xf4, 0x34, 0x38, 0x34, 0xb6, 0xa3, 0xc2, 0x79, 0x8a,
	0x24, 0xea, 0xd6, 0x84, 0x96, 0xd0, 0x50, 0xc0, 0x6f, 0x8f, 0x7d, 0xf9, 0xe5, 0x2c, 0x92, 0xd3,
	0xbe, 0x98, 0xb7, 0x1d, 0xf9, 0xb1, 0xf1, 0xca, 0xd8, 0x63, 0x61, 0x49, 0xe9, 0xc0, 0x6b, 0xf1,
	0x5f, 0x43, 0xce, 0x85, 0x13, 0xce, 0x7e, 0x12, 0x58, 0x9f, 0x0f, 0xc4, 0xf7, 0xd7, 0xfd, 0x4e,
	0xac, 0x61, 0x39, 0x25, 0xd0, 0x28, 0x99, 0x6c, 0x92, 0xfc, 0xa6, 0x5e, 0x4c, 0x58, 0x29, 0xf1,
	0x04, 0xb7, 0x4e, 0x6b, 0x8f, 0x8d, 0x52, 0x5e, 0x87, 0x20, 0x2f, 0xf8, 0x34, 0x9d, 0xdc, 0xf3,
	0xac, 0xc4, 0x03, 0x5c, 0xff, 0xe5, 0x94, 0xbc, 0xdc, 0x24, 0xf9, 0xaa, 0x5e, 0x32, 0x57, 0x4a,
	0xe4, 0x70, 0x1f, 0x07, 0x6b, 0x75, 0x8f, 0xa6, 0xc3, 0xcf, 0xc6, 0xaa, 0x5e, 0xcb, 0x2b, 0xbe,
	0x72, 0x37, 0xfb, 0xaa, 0x7b, 0x63, 0x9b, 0xed, 0x41, 0x9c, 0xbf, 0x4d, 0x3c, 0x42, 0x3a, 0x53,
	0x88, 0x4d, 0xd4, 0xb2, 0xe4, 0xe1, 0xc0, 0x6a, 0x37, 0x19, 0x91, 0xc1, 0xca, 0x8d, 0xd8, 0x9b,
	0x10, 0xb1, 0xa5, 0xc1, 0x46, 0xb9, 0xe5, 0x7a, 0xea, 0xc6, 0x0f, 0x13, 0xe2, 0xcb, 0xa4, 0x0e,
	0x0b, 0xfe, 0xd2, 0xed, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xb0, 0xa2, 0x95, 0x6e, 0x01,
	0x00, 0x00,
}
