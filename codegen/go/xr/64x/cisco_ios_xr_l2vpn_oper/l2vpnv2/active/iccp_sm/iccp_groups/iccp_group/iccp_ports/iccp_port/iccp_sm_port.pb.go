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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_iccp_sm_iccp_groups_iccp_group_iccp_ports_iccp_port

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
	GroupId              uint32   `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
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
	proto.RegisterType((*IccpSmPort_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_KEYS")
	proto.RegisterType((*IccpSmPortInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_info")
	proto.RegisterType((*IccpSmPort)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port")
}

func init() { proto.RegisterFile("iccp_sm_port.proto", fileDescriptor_67f3f01b0dfead04) }

var fileDescriptor_67f3f01b0dfead04 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x8b, 0xd4, 0x30,
	0x14, 0xc6, 0xe9, 0xac, 0xee, 0x4e, 0xdf, 0x6c, 0x17, 0xcc, 0x41, 0x22, 0xba, 0x58, 0x06, 0x85,
	0x9e, 0x7a, 0x98, 0x55, 0xf0, 0x2e, 0x2e, 0x88, 0x22, 0x92, 0x5d, 0x04, 0xbd, 0x84, 0x98, 0xbe,
	0xae, 0x81, 0x26, 0x29, 0x49, 0xb6, 0xf8, 0x47, 0x88, 0x47, 0xff, 0x57, 0x6f, 0x92, 0xd7, 0x71,
	0x76, 0x56, 0xcf, 0xe2, 0xa9, 0x2f, 0xbf, 0x2f, 0xdf, 0x97, 0x8f, 0xb4, 0x05, 0x66, 0xb4, 0x1e,
	0x65, 0xb4, 0x72, 0xf4, 0x21, 0xb5, 0x63, 0xf0, 0xc9, 0xb3, 0x4f, 0xda, 0x44, 0xed, 0xa5, 0xf1,
	0x51, 0x7e, 0x0d, 0x72, 0xd8, 0x4c, 0xa3, 0x93, 0x7e, 0xc4, 0xd0, 0xd2, 0x38, 0x6d, 0x5a, 0xa5,
	0x93, 0x99, 0xb0, 0xdd, 0x5a, 0xe7, 0xe7, 0x55, 0xf0, 0xd7, 0x63, 0xdc, 0x9b, 0xe7, 0x31, 0xc7,
	0xc6, 0x9b, 0x71, 0xfd, 0x16, 0xee, 0xed, 0x9f, 0x28, 0xdf, 0xbc, 0xfa, 0x78, 0xc1, 0x1e, 0xc0,
	0x92, 0x2c, 0xd2, 0x74, 0xbc, 0xa8, 0x8b, 0xa6, 0x12, 0x47, 0xb4, 0x7e, 0xdd, 0xb1, 0x47, 0x50,
	0x1a, 0x97, 0x30, 0xf4, 0x4a, 0x23, 0x5f, 0xd4, 0x45, 0x53, 0x8a, 0x1b, 0xb0, 0xfe, 0xb1, 0xf8,
	0x23, 0xce, 0xb8, 0xde, 0xb3, 0x53, 0x00, 0x5a, 0xc4, 0xa4, 0x12, 0x52, 0x60, 0x29, 0xca, 0x4c,
	0x2e, 0x32, 0x60, 0x4f, 0xe0, 0x84, 0xe4, 0x5e, 0x99, 0x41, 0x6a, 0xdf, 0xcd, 0xb9, 0x95, 0x38,
	0xce, 0xf4, 0x5c, 0x99, 0xe1, 0xa5, 0xef, 0x90, 0x3d, 0x84, 0xb2, 0x8f, 0x76, 0x9b, 0x71, 0x40,
	0x1b, 0x96, 0x7d, 0xb4, 0x73, 0xc4, 0x29, 0xc0, 0x34, 0x28, 0xb7, 0x55, 0xef, 0x90, 0x5a, 0x66,
	0x32, 0xcb, 0x8f, 0x61, 0x45, 0xf2, 0x84, 0x3a, 0xf9, 0xc0, 0xef, 0xd6, 0x07, 0x4d, 0x25, 0xc8,
	0xf1, 0x81, 0x08, 0x7b, 0x0a, 0x27, 0x01, 0x27, 0x0c, 0xd1, 0x78, 0x27, 0x93, 0xb1, 0xc8, 0x0f,
	0x29, 0xa3, 0xda, 0xd1, 0x4b, 0x63, 0x91, 0xbd, 0x00, 0x7e, 0x7b, 0x9b, 0x0c, 0x68, 0x95, 0x71,
	0xc6, 0x5d, 0xf1, 0x23, 0x32, 0xdc, 0xbf, 0x65, 0x10, 0xbf, 0xd5, 0xf5, 0xcf, 0x05, 0x1c, 0xef,
	0x5f, 0x4c, 0x3e, 0x71, 0x77, 0x6d, 0xd2, 0x29, 0x8b, 0x7c, 0x43, 0xf7, 0x52, 0xed, 0xe8, 0x3b,
	0x65, 0x91, 0x7d, 0x2b, 0x00, 0x06, 0xaf, 0xd5, 0x40, 0x2e, 0x7e, 0x56, 0x17, 0xcd, 0x6a, 0x63,
	0xdb, 0x7f, 0xf7, 0x41, 0xb4, 0x7f, 0xbd, 0x3e, 0x51, 0x52, 0x81, 0xf7, 0xb9, 0xf5, 0xf7, 0x02,
	0x56, 0x01, 0xad, 0x4f, 0x38, 0xf7, 0x79, 0xf6, 0x3f, 0xfa, 0xc0, 0xdc, 0x80, 0x0a, 0xad, 0xa1,
	0xb2, 0x4a, 0xcb, 0x7e, 0xb8, 0x8e, 0x5f, 0x64, 0xd2, 0x8e, 0x3f, 0xaf, 0x8b, 0x66, 0x29, 0x56,
	0x56, 0xe9, 0xf3, 0xcc, 0x2e, 0xb5, 0xfb, 0x7c, 0x48, 0x7f, 0xd1, 0xd9, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x53, 0xc2, 0x19, 0x7c, 0x5b, 0x03, 0x00, 0x00,
}