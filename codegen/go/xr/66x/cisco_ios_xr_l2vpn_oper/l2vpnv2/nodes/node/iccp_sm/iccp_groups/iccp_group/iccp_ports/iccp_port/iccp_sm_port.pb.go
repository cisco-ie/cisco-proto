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
// source: iccp_sm_port.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_iccp_sm_iccp_groups_iccp_group_iccp_ports_iccp_port

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

type IccpSmPort_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	GroupId              uint32   `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Interface            string   `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpSmPort_KEYS) Reset()         { *m = IccpSmPort_KEYS{} }
func (m *IccpSmPort_KEYS) String() string { return proto.CompactTextString(m) }
func (*IccpSmPort_KEYS) ProtoMessage()    {}
func (*IccpSmPort_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_67f3f01b0dfead04, []int{0}
}

func (m *IccpSmPort_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpSmPort_KEYS.Unmarshal(m, b)
}
func (m *IccpSmPort_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpSmPort_KEYS.Marshal(b, m, deterministic)
}
func (m *IccpSmPort_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpSmPort_KEYS.Merge(m, src)
}
func (m *IccpSmPort_KEYS) XXX_Size() int {
	return xxx_messageInfo_IccpSmPort_KEYS.Size(m)
}
func (m *IccpSmPort_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpSmPort_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IccpSmPort_KEYS proto.InternalMessageInfo

func (m *IccpSmPort_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *IccpSmPort_KEYS) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IccpSmPort_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

type IccpSmPortInfo struct {
	PortState              string   `protobuf:"bytes,1,opt,name=port_state,json=portState,proto3" json:"port_state,omitempty"`
	PortFailCode           uint32   `protobuf:"varint,2,opt,name=port_fail_code,json=portFailCode,proto3" json:"port_fail_code,omitempty"`
	FsmState               uint32   `protobuf:"varint,3,opt,name=fsm_state,json=fsmState,proto3" json:"fsm_state,omitempty"`
	VlanState              uint32   `protobuf:"varint,4,opt,name=vlan_state,json=vlanState,proto3" json:"vlan_state,omitempty"`
	VlanVector             []uint32 `protobuf:"varint,5,rep,packed,name=vlan_vector,json=vlanVector,proto3" json:"vlan_vector,omitempty"`
	ReversionTime          uint32   `protobuf:"varint,6,opt,name=reversion_time,json=reversionTime,proto3" json:"reversion_time,omitempty"`
	ReversionTimeRemaining uint32   `protobuf:"varint,7,opt,name=reversion_time_remaining,json=reversionTimeRemaining,proto3" json:"reversion_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *IccpSmPortInfo) Reset()         { *m = IccpSmPortInfo{} }
func (m *IccpSmPortInfo) String() string { return proto.CompactTextString(m) }
func (*IccpSmPortInfo) ProtoMessage()    {}
func (*IccpSmPortInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67f3f01b0dfead04, []int{1}
}

func (m *IccpSmPortInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpSmPortInfo.Unmarshal(m, b)
}
func (m *IccpSmPortInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpSmPortInfo.Marshal(b, m, deterministic)
}
func (m *IccpSmPortInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpSmPortInfo.Merge(m, src)
}
func (m *IccpSmPortInfo) XXX_Size() int {
	return xxx_messageInfo_IccpSmPortInfo.Size(m)
}
func (m *IccpSmPortInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpSmPortInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IccpSmPortInfo proto.InternalMessageInfo

func (m *IccpSmPortInfo) GetPortState() string {
	if m != nil {
		return m.PortState
	}
	return ""
}

func (m *IccpSmPortInfo) GetPortFailCode() uint32 {
	if m != nil {
		return m.PortFailCode
	}
	return 0
}

func (m *IccpSmPortInfo) GetFsmState() uint32 {
	if m != nil {
		return m.FsmState
	}
	return 0
}

func (m *IccpSmPortInfo) GetVlanState() uint32 {
	if m != nil {
		return m.VlanState
	}
	return 0
}

func (m *IccpSmPortInfo) GetVlanVector() []uint32 {
	if m != nil {
		return m.VlanVector
	}
	return nil
}

func (m *IccpSmPortInfo) GetReversionTime() uint32 {
	if m != nil {
		return m.ReversionTime
	}
	return 0
}

func (m *IccpSmPortInfo) GetReversionTimeRemaining() uint32 {
	if m != nil {
		return m.ReversionTimeRemaining
	}
	return 0
}

type IccpSmPort struct {
	InterfaceName        string          `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	LocalPort            *IccpSmPortInfo `protobuf:"bytes,51,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemotePort           *IccpSmPortInfo `protobuf:"bytes,52,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
	MacFlushTcn          bool            `protobuf:"varint,53,opt,name=mac_flush_tcn,json=macFlushTcn,proto3" json:"mac_flush_tcn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IccpSmPort) Reset()         { *m = IccpSmPort{} }
func (m *IccpSmPort) String() string { return proto.CompactTextString(m) }
func (*IccpSmPort) ProtoMessage()    {}
func (*IccpSmPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_67f3f01b0dfead04, []int{2}
}

func (m *IccpSmPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpSmPort.Unmarshal(m, b)
}
func (m *IccpSmPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpSmPort.Marshal(b, m, deterministic)
}
func (m *IccpSmPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpSmPort.Merge(m, src)
}
func (m *IccpSmPort) XXX_Size() int {
	return xxx_messageInfo_IccpSmPort.Size(m)
}
func (m *IccpSmPort) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpSmPort.DiscardUnknown(m)
}

var xxx_messageInfo_IccpSmPort proto.InternalMessageInfo

func (m *IccpSmPort) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IccpSmPort) GetLocalPort() *IccpSmPortInfo {
	if m != nil {
		return m.LocalPort
	}
	return nil
}

func (m *IccpSmPort) GetRemotePort() *IccpSmPortInfo {
	if m != nil {
		return m.RemotePort
	}
	return nil
}

func (m *IccpSmPort) GetMacFlushTcn() bool {
	if m != nil {
		return m.MacFlushTcn
	}
	return false
}

func init() {
	proto.RegisterType((*IccpSmPort_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_KEYS")
	proto.RegisterType((*IccpSmPortInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_info")
	proto.RegisterType((*IccpSmPort)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port")
}

func init() { proto.RegisterFile("iccp_sm_port.proto", fileDescriptor_67f3f01b0dfead04) }

var fileDescriptor_67f3f01b0dfead04 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xa6, 0x3b, 0xba, 0x33, 0x7d, 0xb3, 0x5d, 0x30, 0x07, 0x8d, 0xe8, 0x62, 0x19, 0x14, 0x7a,
	0xea, 0x61, 0x56, 0xc1, 0xbb, 0xb8, 0xb0, 0x08, 0x22, 0xdd, 0x45, 0xf0, 0x62, 0x88, 0xe9, 0xeb,
	0x1a, 0x68, 0x92, 0x9a, 0x64, 0x8b, 0x57, 0xff, 0x81, 0x78, 0xf0, 0xf7, 0x4a, 0x5e, 0xeb, 0xec,
	0xac, 0x9e, 0xc5, 0x4b, 0xfb, 0xde, 0xf7, 0xe5, 0xfb, 0xde, 0xc7, 0x6b, 0x03, 0x4c, 0x2b, 0x35,
	0x88, 0x60, 0xc4, 0xe0, 0x7c, 0xac, 0x07, 0xef, 0xa2, 0x63, 0x1f, 0x95, 0x0e, 0xca, 0x09, 0xed,
	0x82, 0xf8, 0xea, 0x45, 0xbf, 0x1d, 0x07, 0x2b, 0xdc, 0x80, 0xbe, 0xa6, 0x72, 0xdc, 0xd6, 0xd6,
	0xb5, 0x18, 0xe8, 0x59, 0xcf, 0xf2, 0xe9, 0x7d, 0xe5, 0xdd, 0xf5, 0x10, 0xf6, 0xea, 0xa9, 0x4c,
	0xd6, 0xe1, 0xa6, 0xdc, 0x20, 0xdc, 0xdb, 0x9f, 0x2a, 0xde, 0xbc, 0xfe, 0x70, 0xc1, 0x1e, 0xc0,
	0x32, 0x19, 0x0a, 0xdd, 0xf2, 0xac, 0xcc, 0xaa, 0xbc, 0x39, 0x4c, 0xed, 0x79, 0xcb, 0x1e, 0xc2,
	0x8a, 0xbc, 0x12, 0x73, 0x50, 0x66, 0x55, 0xd1, 0x2c, 0xa9, 0x3f, 0x6f, 0xd9, 0x63, 0xc8, 0xb5,
	0x8d, 0xe8, 0x3b, 0xa9, 0x90, 0x2f, 0x48, 0x75, 0x03, 0x6c, 0x7e, 0x1e, 0xfc, 0x31, 0x47, 0xdb,
	0xce, 0xb1, 0x13, 0x00, 0x6a, 0x42, 0x94, 0x11, 0xe7, 0x51, 0x79, 0x42, 0x2e, 0x12, 0xc0, 0x9e,
	0xc2, 0x31, 0xd1, 0x9d, 0xd4, 0xbd, 0x50, 0xae, 0xc5, 0x79, 0xe6, 0x51, 0x42, 0xcf, 0xa4, 0xee,
	0x5f, 0xb9, 0x16, 0xd9, 0x23, 0xc8, 0xbb, 0x60, 0x66, 0x8f, 0x05, 0x1d, 0x58, 0x75, 0xc1, 0x4c,
	0x16, 0x27, 0x00, 0x63, 0x2f, 0xed, 0xcc, 0xde, 0x21, 0x36, 0x4f, 0xc8, 0x44, 0x3f, 0x81, 0x35,
	0xd1, 0x23, 0xaa, 0xe8, 0x3c, 0xbf, 0x5b, 0x2e, 0xaa, 0xa2, 0x21, 0xc5, 0x7b, 0x42, 0xd8, 0x33,
	0x38, 0xf6, 0x38, 0xa2, 0x0f, 0xda, 0x59, 0x11, 0xb5, 0x41, 0x7e, 0x48, 0x1e, 0xc5, 0x0e, 0xbd,
	0xd4, 0x06, 0xd9, 0x4b, 0xe0, 0xb7, 0x8f, 0x09, 0x8f, 0x46, 0x6a, 0xab, 0xed, 0x15, 0x5f, 0x92,
	0xe0, 0xfe, 0x2d, 0x41, 0xf3, 0x9b, 0xdd, 0x7c, 0x5b, 0xc0, 0xd1, 0xfe, 0x62, 0xd2, 0xc4, 0xdd,
	0xda, 0x84, 0x95, 0x06, 0xf9, 0x96, 0xf6, 0x52, 0xec, 0xd0, 0xb7, 0xd2, 0x20, 0xfb, 0x9e, 0x01,
	0xf4, 0x4e, 0xc9, 0x9e, 0x54, 0xfc, 0xb4, 0xcc, 0xaa, 0xf5, 0xf6, 0x4b, 0xfd, 0x6f, 0xff, 0x96,
	0xfa, 0xaf, 0x4f, 0xd8, 0xe4, 0x14, 0xe2, 0x5d, 0x4a, 0xfe, 0x23, 0x83, 0xb5, 0x47, 0xe3, 0x22,
	0x4e, 0x99, 0x9e, 0xff, 0xaf, 0x4c, 0x30, 0xa5, 0xa0, 0x50, 0x1b, 0x28, 0x8c, 0x54, 0xa2, 0xeb,
	0xaf, 0xc3, 0x67, 0x11, 0x95, 0xe5, 0x2f, 0xca, 0xac, 0x5a, 0x35, 0x6b, 0x23, 0xd5, 0x59, 0xc2,
	0x2e, 0x95, 0xfd, 0x74, 0x48, 0x57, 0xed, 0xf4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xc7,
	0xd3, 0x3a, 0x80, 0x03, 0x00, 0x00,
}