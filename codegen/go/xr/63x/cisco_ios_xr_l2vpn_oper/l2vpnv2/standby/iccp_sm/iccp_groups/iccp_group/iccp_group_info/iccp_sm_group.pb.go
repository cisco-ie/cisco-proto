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
// source: iccp_sm_group.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_iccp_sm_iccp_groups_iccp_group_iccp_group_info

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

type IccpSmGroup_KEYS struct {
	GroupId              uint32   `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IccpSmGroup_KEYS) Reset()         { *m = IccpSmGroup_KEYS{} }
func (m *IccpSmGroup_KEYS) String() string { return proto.CompactTextString(m) }
func (*IccpSmGroup_KEYS) ProtoMessage()    {}
func (*IccpSmGroup_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_00bba1a912396317, []int{0}
}

func (m *IccpSmGroup_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpSmGroup_KEYS.Unmarshal(m, b)
}
func (m *IccpSmGroup_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpSmGroup_KEYS.Marshal(b, m, deterministic)
}
func (m *IccpSmGroup_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpSmGroup_KEYS.Merge(m, src)
}
func (m *IccpSmGroup_KEYS) XXX_Size() int {
	return xxx_messageInfo_IccpSmGroup_KEYS.Size(m)
}
func (m *IccpSmGroup_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpSmGroup_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IccpSmGroup_KEYS proto.InternalMessageInfo

func (m *IccpSmGroup_KEYS) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
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
	return fileDescriptor_00bba1a912396317, []int{1}
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
	InterfaceName        string          `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	LocalPort            *IccpSmPortInfo `protobuf:"bytes,2,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemotePort           *IccpSmPortInfo `protobuf:"bytes,3,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
	MacFlushTcn          bool            `protobuf:"varint,4,opt,name=mac_flush_tcn,json=macFlushTcn,proto3" json:"mac_flush_tcn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IccpSmPort) Reset()         { *m = IccpSmPort{} }
func (m *IccpSmPort) String() string { return proto.CompactTextString(m) }
func (*IccpSmPort) ProtoMessage()    {}
func (*IccpSmPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_00bba1a912396317, []int{2}
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

type IccpSmGroup struct {
	GroupId              uint32        `protobuf:"varint,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	LocalNodeId          uint32        `protobuf:"varint,51,opt,name=local_node_id,json=localNodeId,proto3" json:"local_node_id,omitempty"`
	RemoteNodeId         uint32        `protobuf:"varint,52,opt,name=remote_node_id,json=remoteNodeId,proto3" json:"remote_node_id,omitempty"`
	State                string        `protobuf:"bytes,53,opt,name=state,proto3" json:"state,omitempty"`
	Ports                []*IccpSmPort `protobuf:"bytes,54,rep,name=ports,proto3" json:"ports,omitempty"`
	IccpTransportUp      bool          `protobuf:"varint,55,opt,name=iccp_transport_up,json=iccpTransportUp,proto3" json:"iccp_transport_up,omitempty"`
	IccpMemberUp         bool          `protobuf:"varint,56,opt,name=iccp_member_up,json=iccpMemberUp,proto3" json:"iccp_member_up,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IccpSmGroup) Reset()         { *m = IccpSmGroup{} }
func (m *IccpSmGroup) String() string { return proto.CompactTextString(m) }
func (*IccpSmGroup) ProtoMessage()    {}
func (*IccpSmGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_00bba1a912396317, []int{3}
}

func (m *IccpSmGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IccpSmGroup.Unmarshal(m, b)
}
func (m *IccpSmGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IccpSmGroup.Marshal(b, m, deterministic)
}
func (m *IccpSmGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IccpSmGroup.Merge(m, src)
}
func (m *IccpSmGroup) XXX_Size() int {
	return xxx_messageInfo_IccpSmGroup.Size(m)
}
func (m *IccpSmGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_IccpSmGroup.DiscardUnknown(m)
}

var xxx_messageInfo_IccpSmGroup proto.InternalMessageInfo

func (m *IccpSmGroup) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *IccpSmGroup) GetLocalNodeId() uint32 {
	if m != nil {
		return m.LocalNodeId
	}
	return 0
}

func (m *IccpSmGroup) GetRemoteNodeId() uint32 {
	if m != nil {
		return m.RemoteNodeId
	}
	return 0
}

func (m *IccpSmGroup) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *IccpSmGroup) GetPorts() []*IccpSmPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *IccpSmGroup) GetIccpTransportUp() bool {
	if m != nil {
		return m.IccpTransportUp
	}
	return false
}

func (m *IccpSmGroup) GetIccpMemberUp() bool {
	if m != nil {
		return m.IccpMemberUp
	}
	return false
}

func init() {
	proto.RegisterType((*IccpSmGroup_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_group_info.iccp_sm_group_KEYS")
	proto.RegisterType((*IccpSmPortInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_group_info.iccp_sm_port_info")
	proto.RegisterType((*IccpSmPort)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_group_info.iccp_sm_port")
	proto.RegisterType((*IccpSmGroup)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_group_info.iccp_sm_group")
}

func init() { proto.RegisterFile("iccp_sm_group.proto", fileDescriptor_00bba1a912396317) }

var fileDescriptor_00bba1a912396317 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6a, 0x13, 0x41,
	0x14, 0x26, 0x89, 0x69, 0x93, 0xb3, 0xdd, 0x8a, 0xa3, 0xc8, 0x8a, 0x14, 0x43, 0x50, 0x08, 0x5e,
	0xac, 0x90, 0xfa, 0xd3, 0x7b, 0xb1, 0x50, 0xc4, 0x22, 0xdb, 0xb4, 0xe0, 0xd5, 0x30, 0x99, 0x9d,
	0xad, 0x03, 0x3b, 0x33, 0xcb, 0xcc, 0x24, 0xa8, 0x4f, 0x20, 0x3e, 0x80, 0xcf, 0xe6, 0x6b, 0xf8,
	0x06, 0x32, 0x67, 0x76, 0x4b, 0x16, 0xaf, 0xdb, 0xab, 0xcc, 0x7c, 0xdf, 0x77, 0x7e, 0xe6, 0x3b,
	0x27, 0x0b, 0x0f, 0x25, 0xe7, 0x0d, 0x75, 0x8a, 0x5e, 0x5b, 0xb3, 0x69, 0xf2, 0xc6, 0x1a, 0x6f,
	0xc8, 0x15, 0x97, 0x8e, 0x1b, 0x2a, 0x8d, 0xa3, 0xdf, 0x2c, 0xad, 0x97, 0xdb, 0x46, 0x53, 0xd3,
	0x08, 0x9b, 0xe3, 0x71, 0xbb, 0xcc, 0x9d, 0x67, 0xba, 0x5c, 0x7f, 0xcf, 0xdb, 0xe0, 0xf8, 0x8b,
	0x19, 0xdc, 0xce, 0x79, 0xe7, 0x48, 0xa5, 0xae, 0xcc, 0xfc, 0x15, 0x90, 0x5e, 0x39, 0xfa, 0xf1,
	0xc3, 0x97, 0x0b, 0xf2, 0x04, 0x26, 0xad, 0xa6, 0xcc, 0x06, 0xb3, 0xc1, 0x22, 0x2d, 0xf6, 0xf1,
	0x7e, 0x56, 0xce, 0x7f, 0x0f, 0xe1, 0x41, 0x17, 0xd1, 0x18, 0xeb, 0x31, 0x0d, 0x39, 0x02, 0xc0,
	0x8b, 0xf3, 0xcc, 0x0b, 0x0c, 0x99, 0x16, 0xd3, 0x80, 0x5c, 0x04, 0x80, 0x3c, 0x87, 0x43, 0xa4,
	0x2b, 0x26, 0x6b, 0xca, 0x4d, 0x29, 0xb2, 0x21, 0x66, 0x3d, 0x08, 0xe8, 0x29, 0x93, 0xf5, 0x7b,
	0x53, 0x0a, 0xf2, 0x14, 0xa6, 0x95, 0x53, 0x6d, 0x8e, 0x11, 0x0a, 0x26, 0x95, 0x53, 0x31, 0xc5,
	0x11, 0xc0, 0xb6, 0x66, 0xba, 0x65, 0xef, 0x21, 0x3b, 0x0d, 0x48, 0xa4, 0x9f, 0x41, 0x82, 0xf4,
	0x56, 0x70, 0x6f, 0x6c, 0x36, 0x9e, 0x8d, 0x16, 0x69, 0x81, 0x11, 0x57, 0x88, 0x90, 0x17, 0x70,
	0x68, 0xc5, 0x56, 0x58, 0x27, 0x8d, 0xa6, 0x5e, 0x2a, 0x91, 0xed, 0x61, 0x8e, 0xf4, 0x06, 0x5d,
	0x49, 0x25, 0xc8, 0x09, 0x64, 0x7d, 0x19, 0xb5, 0x42, 0x31, 0xa9, 0xa5, 0xbe, 0xce, 0xf6, 0x31,
	0xe0, 0x71, 0x2f, 0xa0, 0xe8, 0xd8, 0xf9, 0xdf, 0x21, 0x1c, 0xec, 0x1a, 0x13, 0x2a, 0x4a, 0xed,
	0x85, 0xad, 0x18, 0x17, 0x54, 0x33, 0xd5, 0xf9, 0x92, 0xde, 0xa0, 0xe7, 0x4c, 0x09, 0xf2, 0x73,
	0x00, 0x50, 0x1b, 0xce, 0x6a, 0x8c, 0x42, 0x63, 0x92, 0xa5, 0xcc, 0x6f, 0x67, 0xde, 0xf9, 0x7f,
	0xa3, 0x2b, 0xa6, 0x58, 0xfc, 0x73, 0xe8, 0xf8, 0xd7, 0x00, 0x12, 0x2b, 0x94, 0xf1, 0x22, 0xf6,
	0x32, 0xba, 0xeb, 0x5e, 0x20, 0x56, 0xc7, 0x66, 0xe6, 0x90, 0x2a, 0xc6, 0x69, 0x55, 0x6f, 0xdc,
	0x57, 0xea, 0xb9, 0xc6, 0x99, 0x4f, 0x8a, 0x44, 0x31, 0x7e, 0x1a, 0xb0, 0x15, 0xd7, 0xf3, 0x3f,
	0x43, 0x48, 0x7b, 0xeb, 0xdb, 0xdb, 0xdc, 0x65, 0x6f, 0x73, 0x43, 0xc2, 0xe8, 0xb3, 0x36, 0xa5,
	0x08, 0xfc, 0x31, 0xf2, 0x09, 0x82, 0xe7, 0xa6, 0x14, 0x67, 0x65, 0x58, 0xd4, 0xd6, 0x80, 0x4e,
	0xf4, 0x3a, 0x2e, 0x6a, 0x44, 0x5b, 0xd5, 0x23, 0x18, 0xc7, 0x35, 0x7c, 0x83, 0x03, 0x8d, 0x17,
	0xf2, 0x03, 0xc6, 0xe1, 0x25, 0x2e, 0x7b, 0x3b, 0x1b, 0x2d, 0x92, 0x65, 0x79, 0x17, 0xb6, 0x15,
	0xb1, 0x24, 0x79, 0xd9, 0xfe, 0x29, 0xbd, 0x65, 0xda, 0xa1, 0x9f, 0x9b, 0x26, 0x7b, 0x87, 0x86,
	0xdd, 0x0f, 0xc4, 0xaa, 0xc3, 0x2f, 0x9b, 0xf0, 0x46, 0xd4, 0x2a, 0xa1, 0xd6, 0xc2, 0x06, 0xe1,
	0x09, 0x0a, 0x71, 0x7b, 0x3f, 0x21, 0x78, 0xd9, 0xac, 0xf7, 0xf0, 0xbb, 0x73, 0xfc, 0x2f, 0x00,
	0x00, 0xff, 0xff, 0xe4, 0xfb, 0x0e, 0x89, 0x8e, 0x04, 0x00, 0x00,
}