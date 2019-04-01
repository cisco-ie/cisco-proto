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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_iccp_sm_iccp_groups_iccp_group_iccp_ports_iccp_port

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
	proto.RegisterType((*IccpSmPort_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_KEYS")
	proto.RegisterType((*IccpSmPortInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port_info")
	proto.RegisterType((*IccpSmPort)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.iccp_sm.iccp_groups.iccp_group.iccp_ports.iccp_port.iccp_sm_port")
}

func init() { proto.RegisterFile("iccp_sm_port.proto", fileDescriptor_67f3f01b0dfead04) }

var fileDescriptor_67f3f01b0dfead04 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x8b, 0xd4, 0x30,
	0x18, 0xc5, 0xe9, 0xac, 0xee, 0x4e, 0xbf, 0x6e, 0x17, 0xcc, 0x41, 0x22, 0xba, 0x58, 0x8a, 0x42,
	0x4f, 0x3d, 0x74, 0x15, 0xbc, 0x8b, 0x0b, 0xa2, 0x88, 0x64, 0x17, 0x41, 0x3c, 0x84, 0x6c, 0x9a,
	0xae, 0x81, 0x26, 0x29, 0x49, 0xb6, 0xe8, 0x1f, 0x21, 0xde, 0xfc, 0x67, 0xbd, 0x48, 0xbe, 0x8e,
	0xb3, 0x33, 0x7a, 0x16, 0x4f, 0xfd, 0xf2, 0x7b, 0x79, 0x2f, 0x8f, 0xb4, 0x05, 0xa2, 0xa5, 0x9c,
	0x78, 0x30, 0x7c, 0x72, 0x3e, 0xb6, 0x93, 0x77, 0xd1, 0x91, 0x4f, 0x52, 0x07, 0xe9, 0xb8, 0x76,
	0x81, 0x7f, 0xf1, 0x7c, 0xec, 0xe6, 0xc9, 0x72, 0x37, 0x29, 0xdf, 0xe2, 0x38, 0x77, 0x6d, 0x88,
	0xc2, 0xf6, 0x57, 0x5f, 0xdb, 0x8d, 0x77, 0x79, 0x5e, 0x7b, 0x77, 0x33, 0x85, 0x9d, 0x79, 0x19,
	0x53, 0x6e, 0xb8, 0x1d, 0xeb, 0xb7, 0x70, 0x6f, 0xf7, 0x48, 0xfe, 0xe6, 0xd5, 0xc7, 0x0b, 0xf2,
	0x00, 0xd6, 0x68, 0xe1, 0xba, 0xa7, 0x59, 0x95, 0x35, 0x25, 0x3b, 0xc2, 0xf5, 0xeb, 0x9e, 0x3c,
	0x82, 0x5c, 0xdb, 0xa8, 0xfc, 0x20, 0xa4, 0xa2, 0xab, 0x2a, 0x6b, 0x72, 0x76, 0x0b, 0xea, 0x1f,
	0xab, 0x3f, 0xe2, 0xb4, 0x1d, 0x1c, 0x39, 0x05, 0xc0, 0x45, 0x88, 0x22, 0x2a, 0x0c, 0xcc, 0x59,
	0x9e, 0xc8, 0x45, 0x02, 0xe4, 0x09, 0x9c, 0xa0, 0x3c, 0x08, 0x3d, 0x72, 0xe9, 0xfa, 0x25, 0xb7,
	0x64, 0xc7, 0x89, 0x9e, 0x0b, 0x3d, 0xbe, 0x74, 0xbd, 0x22, 0x0f, 0x21, 0x1f, 0x82, 0xd9, 0x64,
	0x1c, 0xe0, 0x86, 0xf5, 0x10, 0xcc, 0x12, 0x71, 0x0a, 0x30, 0x8f, 0xc2, 0x6e, 0xd4, 0x3b, 0xa8,
	0xe6, 0x89, 0x2c, 0xf2, 0x63, 0x28, 0x50, 0x9e, 0x95, 0x8c, 0xce, 0xd3, 0xbb, 0xd5, 0x41, 0x53,
	0x32, 0x74, 0x7c, 0x40, 0x42, 0x9e, 0xc2, 0x89, 0x57, 0xb3, 0xf2, 0x41, 0x3b, 0xcb, 0xa3, 0x36,
	0x8a, 0x1e, 0x62, 0x46, 0xb9, 0xa5, 0x97, 0xda, 0x28, 0xf2, 0x02, 0xe8, 0xfe, 0x36, 0xee, 0x95,
	0x11, 0xda, 0x6a, 0x7b, 0x4d, 0x8f, 0xd0, 0x70, 0x7f, 0xcf, 0xc0, 0x7e, 0xab, 0xf5, 0xcf, 0x15,
	0x1c, 0xef, 0x5e, 0x4c, 0x3a, 0x71, 0x7b, 0x6d, 0xdc, 0x0a, 0xa3, 0x68, 0x87, 0xf7, 0x52, 0x6e,
	0xe9, 0x3b, 0x61, 0x14, 0xf9, 0x96, 0x01, 0x8c, 0x4e, 0x8a, 0x11, 0x5d, 0xf4, 0xac, 0xca, 0x9a,
	0xa2, 0xb3, 0xed, 0x3f, 0xfc, 0x22, 0xda, 0xbf, 0xde, 0x1f, 0xcb, 0xb1, 0xc1, 0xfb, 0x54, 0xfb,
	0x7b, 0x06, 0x85, 0x57, 0xc6, 0x45, 0xb5, 0x14, 0x7a, 0xf6, 0x5f, 0x0a, 0xc1, 0x52, 0x01, 0x1b,
	0xd5, 0x50, 0x1a, 0x21, 0xf9, 0x30, 0xde, 0x84, 0xcf, 0x3c, 0x4a, 0x4b, 0x9f, 0x57, 0x59, 0xb3,
	0x66, 0x85, 0x11, 0xf2, 0x3c, 0xb1, 0x4b, 0x69, 0xaf, 0x0e, 0xf1, 0x47, 0x3a, 0xfb, 0x15, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0xdc, 0xf5, 0xc6, 0x5e, 0x03, 0x00, 0x00,
}
