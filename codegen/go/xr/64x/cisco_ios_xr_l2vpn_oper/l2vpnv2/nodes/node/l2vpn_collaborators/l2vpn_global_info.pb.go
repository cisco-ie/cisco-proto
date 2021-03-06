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
// source: l2vpn_global_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_l2vpn_collaborators

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

type L2VpnGlobalInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnGlobalInfo_KEYS) Reset()         { *m = L2VpnGlobalInfo_KEYS{} }
func (m *L2VpnGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnGlobalInfo_KEYS) ProtoMessage()    {}
func (*L2VpnGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_575868185b4a3972, []int{0}
}

func (m *L2VpnGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *L2VpnGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGlobalInfo_KEYS.Merge(m, src)
}
func (m *L2VpnGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnGlobalInfo_KEYS.Size(m)
}
func (m *L2VpnGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGlobalInfo_KEYS proto.InternalMessageInfo

func (m *L2VpnGlobalInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type L2VPNCollabCntr struct {
	CollaboratorName     string   `protobuf:"bytes,1,opt,name=collaborator_name,json=collaboratorName,proto3" json:"collaborator_name,omitempty"`
	Up                   uint32   `protobuf:"varint,2,opt,name=up,proto3" json:"up,omitempty"`
	Down                 uint32   `protobuf:"varint,3,opt,name=down,proto3" json:"down,omitempty"`
	IsUp                 bool     `protobuf:"varint,4,opt,name=is_up,json=isUp,proto3" json:"is_up,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VPNCollabCntr) Reset()         { *m = L2VPNCollabCntr{} }
func (m *L2VPNCollabCntr) String() string { return proto.CompactTextString(m) }
func (*L2VPNCollabCntr) ProtoMessage()    {}
func (*L2VPNCollabCntr) Descriptor() ([]byte, []int) {
	return fileDescriptor_575868185b4a3972, []int{1}
}

func (m *L2VPNCollabCntr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VPNCollabCntr.Unmarshal(m, b)
}
func (m *L2VPNCollabCntr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VPNCollabCntr.Marshal(b, m, deterministic)
}
func (m *L2VPNCollabCntr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VPNCollabCntr.Merge(m, src)
}
func (m *L2VPNCollabCntr) XXX_Size() int {
	return xxx_messageInfo_L2VPNCollabCntr.Size(m)
}
func (m *L2VPNCollabCntr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VPNCollabCntr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VPNCollabCntr proto.InternalMessageInfo

func (m *L2VPNCollabCntr) GetCollaboratorName() string {
	if m != nil {
		return m.CollaboratorName
	}
	return ""
}

func (m *L2VPNCollabCntr) GetUp() uint32 {
	if m != nil {
		return m.Up
	}
	return 0
}

func (m *L2VPNCollabCntr) GetDown() uint32 {
	if m != nil {
		return m.Down
	}
	return 0
}

func (m *L2VPNCollabCntr) GetIsUp() bool {
	if m != nil {
		return m.IsUp
	}
	return false
}

type L2VPNCollabStats struct {
	Count                *L2VPNCollabCntr `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VPNCollabStats) Reset()         { *m = L2VPNCollabStats{} }
func (m *L2VPNCollabStats) String() string { return proto.CompactTextString(m) }
func (*L2VPNCollabStats) ProtoMessage()    {}
func (*L2VPNCollabStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_575868185b4a3972, []int{2}
}

func (m *L2VPNCollabStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VPNCollabStats.Unmarshal(m, b)
}
func (m *L2VPNCollabStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VPNCollabStats.Marshal(b, m, deterministic)
}
func (m *L2VPNCollabStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VPNCollabStats.Merge(m, src)
}
func (m *L2VPNCollabStats) XXX_Size() int {
	return xxx_messageInfo_L2VPNCollabStats.Size(m)
}
func (m *L2VPNCollabStats) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VPNCollabStats.DiscardUnknown(m)
}

var xxx_messageInfo_L2VPNCollabStats proto.InternalMessageInfo

func (m *L2VPNCollabStats) GetCount() *L2VPNCollabCntr {
	if m != nil {
		return m.Count
	}
	return nil
}

type L2VpnGlobalInfo struct {
	CollaboratorStatistics *L2VPNCollabStats `protobuf:"bytes,50,opt,name=collaborator_statistics,json=collaboratorStatistics,proto3" json:"collaborator_statistics,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *L2VpnGlobalInfo) Reset()         { *m = L2VpnGlobalInfo{} }
func (m *L2VpnGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnGlobalInfo) ProtoMessage()    {}
func (*L2VpnGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_575868185b4a3972, []int{3}
}

func (m *L2VpnGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGlobalInfo.Unmarshal(m, b)
}
func (m *L2VpnGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGlobalInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGlobalInfo.Merge(m, src)
}
func (m *L2VpnGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnGlobalInfo.Size(m)
}
func (m *L2VpnGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGlobalInfo proto.InternalMessageInfo

func (m *L2VpnGlobalInfo) GetCollaboratorStatistics() *L2VPNCollabStats {
	if m != nil {
		return m.CollaboratorStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnGlobalInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_collaborators.l2vpn_global_info_KEYS")
	proto.RegisterType((*L2VPNCollabCntr)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_collaborators.L2VPNCollabCntr")
	proto.RegisterType((*L2VPNCollabStats)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_collaborators.L2VPNCollabStats")
	proto.RegisterType((*L2VpnGlobalInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.l2vpn_collaborators.l2vpn_global_info")
}

func init() { proto.RegisterFile("l2vpn_global_info.proto", fileDescriptor_575868185b4a3972) }

var fileDescriptor_575868185b4a3972 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0x49, 0xed, 0xae, 0x3a, 0xa2, 0xee, 0x46, 0xd8, 0xf6, 0x58, 0x7a, 0x2a, 0x08, 0x05,
	0xeb, 0xdd, 0xcb, 0xe2, 0x41, 0x94, 0x75, 0x69, 0x51, 0xf0, 0x14, 0xd2, 0x3f, 0x4a, 0xa0, 0x9b,
	0x09, 0x49, 0xba, 0xab, 0x8f, 0xe0, 0x53, 0xf8, 0xaa, 0xb2, 0xa9, 0x48, 0x75, 0x8f, 0x7a, 0x09,
	0x93, 0x99, 0xcc, 0xef, 0x9b, 0x6f, 0x08, 0x04, 0x6d, 0xb6, 0x56, 0x92, 0xbd, 0xb4, 0x58, 0xf2,
	0x96, 0x09, 0xf9, 0x8c, 0xa9, 0xd2, 0x68, 0x91, 0x5e, 0x55, 0xc2, 0x54, 0xc8, 0x04, 0x1a, 0xf6,
	0xaa, 0x59, 0xff, 0x0a, 0x55, 0xa3, 0x53, 0x17, 0xae, 0xb3, 0x54, 0x62, 0xdd, 0x18, 0x77, 0xf6,
	0x29, 0x56, 0x61, 0xdb, 0xf2, 0x12, 0x35, 0xb7, 0xa8, 0x4d, 0x7c, 0x01, 0xb3, 0x1d, 0x34, 0xbb,
	0xbd, 0x7e, 0x2a, 0x68, 0x00, 0xfb, 0xdb, 0x2e, 0x26, 0xea, 0x90, 0x44, 0x24, 0x39, 0xcc, 0xc7,
	0xdb, 0xeb, 0x4d, 0x1d, 0x6f, 0xe0, 0xf4, 0x2e, 0x7b, 0x5c, 0x2e, 0xe6, 0x0e, 0x34, 0x97, 0x56,
	0xd3, 0x73, 0x98, 0x0e, 0xb1, 0x4c, 0xf2, 0x55, 0xf3, 0xd5, 0x35, 0x19, 0x16, 0x16, 0x7c, 0xd5,
	0xd0, 0x13, 0xf0, 0x3a, 0x15, 0x7a, 0x11, 0x49, 0x8e, 0x73, 0xaf, 0x53, 0x94, 0x82, 0x5f, 0xe3,
	0x46, 0x86, 0x7b, 0x2e, 0xe3, 0x62, 0x7a, 0x06, 0x23, 0x61, 0x58, 0xa7, 0x42, 0x3f, 0x22, 0xc9,
	0x41, 0xee, 0x0b, 0xf3, 0xa0, 0xe2, 0x37, 0x98, 0x0c, 0x84, 0x0b, 0xcb, 0xad, 0xa1, 0x0d, 0x8c,
	0x2a, 0xec, 0xa4, 0x75, 0x6a, 0x47, 0xd9, 0x7d, 0xfa, 0xb7, 0x7d, 0xa4, 0xbf, 0x9c, 0xe5, 0x3d,
	0x3d, 0xfe, 0x20, 0x30, 0xdd, 0xd9, 0x13, 0x7d, 0x27, 0x10, 0xfc, 0xf0, 0x6d, 0x2c, 0xb7, 0xc2,
	0x58, 0x51, 0x99, 0x30, 0x73, 0xf3, 0x2c, 0xff, 0x71, 0x1e, 0x67, 0x38, 0x9f, 0x0d, 0xeb, 0xc5,
	0xb7, 0x5e, 0x39, 0x76, 0xff, 0xe1, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xa0, 0x79, 0x9b,
	0x2a, 0x02, 0x00, 0x00,
}
