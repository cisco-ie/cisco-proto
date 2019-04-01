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
// source: evpn_grp_info.proto

package cisco_ios_xr_evpn_oper_evpn_nodes_node_evpn_groups_evpn_group

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

type EvpnGrpInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	GroupNumber          uint32   `protobuf:"varint,2,opt,name=group_number,json=groupNumber,proto3" json:"group_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvpnGrpInfo_KEYS) Reset()         { *m = EvpnGrpInfo_KEYS{} }
func (m *EvpnGrpInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*EvpnGrpInfo_KEYS) ProtoMessage()    {}
func (*EvpnGrpInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2947777f5470ef38, []int{0}
}

func (m *EvpnGrpInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvpnGrpInfo_KEYS.Unmarshal(m, b)
}
func (m *EvpnGrpInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvpnGrpInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *EvpnGrpInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvpnGrpInfo_KEYS.Merge(m, src)
}
func (m *EvpnGrpInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_EvpnGrpInfo_KEYS.Size(m)
}
func (m *EvpnGrpInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EvpnGrpInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EvpnGrpInfo_KEYS proto.InternalMessageInfo

func (m *EvpnGrpInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *EvpnGrpInfo_KEYS) GetGroupNumber() uint32 {
	if m != nil {
		return m.GroupNumber
	}
	return 0
}

type EvpnGrpIntfInfo struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EvpnGrpIntfInfo) Reset()         { *m = EvpnGrpIntfInfo{} }
func (m *EvpnGrpIntfInfo) String() string { return proto.CompactTextString(m) }
func (*EvpnGrpIntfInfo) ProtoMessage()    {}
func (*EvpnGrpIntfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2947777f5470ef38, []int{1}
}

func (m *EvpnGrpIntfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvpnGrpIntfInfo.Unmarshal(m, b)
}
func (m *EvpnGrpIntfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvpnGrpIntfInfo.Marshal(b, m, deterministic)
}
func (m *EvpnGrpIntfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvpnGrpIntfInfo.Merge(m, src)
}
func (m *EvpnGrpIntfInfo) XXX_Size() int {
	return xxx_messageInfo_EvpnGrpIntfInfo.Size(m)
}
func (m *EvpnGrpIntfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EvpnGrpIntfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EvpnGrpIntfInfo proto.InternalMessageInfo

func (m *EvpnGrpIntfInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *EvpnGrpIntfInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type EvpnGrpInfo struct {
	GroupId              uint32             `protobuf:"varint,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	State                string             `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	CoreInterface        []*EvpnGrpIntfInfo `protobuf:"bytes,52,rep,name=core_interface,json=coreInterface,proto3" json:"core_interface,omitempty"`
	AccessInterface      []*EvpnGrpIntfInfo `protobuf:"bytes,53,rep,name=access_interface,json=accessInterface,proto3" json:"access_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EvpnGrpInfo) Reset()         { *m = EvpnGrpInfo{} }
func (m *EvpnGrpInfo) String() string { return proto.CompactTextString(m) }
func (*EvpnGrpInfo) ProtoMessage()    {}
func (*EvpnGrpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2947777f5470ef38, []int{2}
}

func (m *EvpnGrpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvpnGrpInfo.Unmarshal(m, b)
}
func (m *EvpnGrpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvpnGrpInfo.Marshal(b, m, deterministic)
}
func (m *EvpnGrpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvpnGrpInfo.Merge(m, src)
}
func (m *EvpnGrpInfo) XXX_Size() int {
	return xxx_messageInfo_EvpnGrpInfo.Size(m)
}
func (m *EvpnGrpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EvpnGrpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EvpnGrpInfo proto.InternalMessageInfo

func (m *EvpnGrpInfo) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *EvpnGrpInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *EvpnGrpInfo) GetCoreInterface() []*EvpnGrpIntfInfo {
	if m != nil {
		return m.CoreInterface
	}
	return nil
}

func (m *EvpnGrpInfo) GetAccessInterface() []*EvpnGrpIntfInfo {
	if m != nil {
		return m.AccessInterface
	}
	return nil
}

func init() {
	proto.RegisterType((*EvpnGrpInfo_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evpn_groups.evpn_group.evpn_grp_info_KEYS")
	proto.RegisterType((*EvpnGrpIntfInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evpn_groups.evpn_group.evpn_grp_intf_info")
	proto.RegisterType((*EvpnGrpInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evpn_groups.evpn_group.evpn_grp_info")
}

func init() { proto.RegisterFile("evpn_grp_info.proto", fileDescriptor_2947777f5470ef38) }

var fileDescriptor_2947777f5470ef38 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xd5, 0x22, 0x5a, 0xea, 0x92, 0x82, 0x0c, 0x12, 0x61, 0x0b, 0x91, 0x90, 0x32, 0x79,
	0x68, 0x61, 0x64, 0x64, 0x88, 0x90, 0x2a, 0x1a, 0x26, 0xa6, 0x53, 0xea, 0x5c, 0x90, 0x87, 0xd8,
	0x91, 0xed, 0xa2, 0x0e, 0x3c, 0x18, 0x8f, 0x87, 0x72, 0x29, 0x21, 0x11, 0x23, 0x62, 0x89, 0xee,
	0x4e, 0xb9, 0xef, 0xfb, 0x6d, 0x99, 0x5d, 0xe0, 0x7b, 0xad, 0xe1, 0xcd, 0xd6, 0xa0, 0x74, 0x69,
	0x44, 0x6d, 0x8d, 0x37, 0xfc, 0x41, 0x2a, 0x27, 0x0d, 0x28, 0xe3, 0x60, 0x6f, 0x81, 0xfe, 0x30,
	0x35, 0x5a, 0xd1, 0x54, 0x42, 0x9b, 0x02, 0x1d, 0x7d, 0xc5, 0x61, 0xd7, 0xec, 0x6a, 0xd7, 0xab,
	0xe3, 0x67, 0xc6, 0x07, 0x54, 0x78, 0x7a, 0x7c, 0x7d, 0xe1, 0x57, 0x6c, 0xda, 0x2c, 0x81, 0x2a,
	0xc2, 0x51, 0x34, 0x4a, 0x66, 0xd9, 0xa4, 0x69, 0xd3, 0x82, 0xdf, 0xb0, 0x53, 0xda, 0x03, 0xbd,
	0xab, 0xb6, 0x68, 0xc3, 0x71, 0x34, 0x4a, 0x82, 0x6c, 0x4e, 0xb3, 0x35, 0x8d, 0xe2, 0xcd, 0x80,
	0xe8, 0x4b, 0xc2, 0xf2, 0x5b, 0xb6, 0x50, 0xda, 0xa3, 0x2d, 0x73, 0x89, 0xa0, 0xf3, 0x0a, 0x0f,
	0xe0, 0xa0, 0x9b, 0xae, 0xf3, 0x0a, 0xf9, 0x25, 0x3b, 0x76, 0x3e, 0xf7, 0x48, 0xe0, 0x59, 0xd6,
	0x36, 0xf1, 0xe7, 0x98, 0x05, 0x83, 0x94, 0xfc, 0x9a, 0x9d, 0xb4, 0x39, 0x54, 0x11, 0x2e, 0x29,
	0xc3, 0x94, 0xfa, 0xb4, 0xf8, 0x41, 0xac, 0x7a, 0x08, 0xbe, 0x67, 0x0b, 0x69, 0x2c, 0x42, 0xa7,
	0x0b, 0xef, 0xa2, 0xa3, 0x64, 0xbe, 0xdc, 0x88, 0x3f, 0xdd, 0x9f, 0xf8, 0x7d, 0xd4, 0x2c, 0x68,
	0x44, 0xe9, 0xb7, 0x87, 0x7f, 0xb0, 0xf3, 0x5c, 0x4a, 0x74, 0xae, 0xe7, 0xbe, 0xff, 0x2f, 0xf7,
	0x59, 0xab, 0xea, 0xec, 0xdb, 0x09, 0xbd, 0x92, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0xad, 0xcc, 0x3d, 0x3c, 0x02, 0x00, 0x00,
}