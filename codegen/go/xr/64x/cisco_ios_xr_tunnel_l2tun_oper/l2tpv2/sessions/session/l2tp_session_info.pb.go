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
// source: l2tp_session_info.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tpv2_sessions_session

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

type L2TpSessionInfo_KEYS struct {
	LocalTunnelId        uint32   `protobuf:"varint,1,opt,name=local_tunnel_id,json=localTunnelId,proto3" json:"local_tunnel_id,omitempty"`
	LocalSessionId       uint32   `protobuf:"varint,2,opt,name=local_session_id,json=localSessionId,proto3" json:"local_session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpSessionInfo_KEYS) Reset()         { *m = L2TpSessionInfo_KEYS{} }
func (m *L2TpSessionInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2TpSessionInfo_KEYS) ProtoMessage()    {}
func (*L2TpSessionInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8422bc409d0e143e, []int{0}
}

func (m *L2TpSessionInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpSessionInfo_KEYS.Unmarshal(m, b)
}
func (m *L2TpSessionInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpSessionInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2TpSessionInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpSessionInfo_KEYS.Merge(m, src)
}
func (m *L2TpSessionInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2TpSessionInfo_KEYS.Size(m)
}
func (m *L2TpSessionInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpSessionInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpSessionInfo_KEYS proto.InternalMessageInfo

func (m *L2TpSessionInfo_KEYS) GetLocalTunnelId() uint32 {
	if m != nil {
		return m.LocalTunnelId
	}
	return 0
}

func (m *L2TpSessionInfo_KEYS) GetLocalSessionId() uint32 {
	if m != nil {
		return m.LocalSessionId
	}
	return 0
}

type L2TpSessionXconnect struct {
	CircuitName            string   `protobuf:"bytes,1,opt,name=circuit_name,json=circuitName,proto3" json:"circuit_name,omitempty"`
	SessionvcId            uint32   `protobuf:"varint,2,opt,name=sessionvc_id,json=sessionvcId,proto3" json:"sessionvc_id,omitempty"`
	IsCircuitStateUp       bool     `protobuf:"varint,3,opt,name=is_circuit_state_up,json=isCircuitStateUp,proto3" json:"is_circuit_state_up,omitempty"`
	IsLocalCircuitStateUp  bool     `protobuf:"varint,4,opt,name=is_local_circuit_state_up,json=isLocalCircuitStateUp,proto3" json:"is_local_circuit_state_up,omitempty"`
	IsRemoteCircuitStateUp bool     `protobuf:"varint,5,opt,name=is_remote_circuit_state_up,json=isRemoteCircuitStateUp,proto3" json:"is_remote_circuit_state_up,omitempty"`
	Ipv6ProtocolTunneling  bool     `protobuf:"varint,6,opt,name=ipv6_protocol_tunneling,json=ipv6ProtocolTunneling,proto3" json:"ipv6_protocol_tunneling,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *L2TpSessionXconnect) Reset()         { *m = L2TpSessionXconnect{} }
func (m *L2TpSessionXconnect) String() string { return proto.CompactTextString(m) }
func (*L2TpSessionXconnect) ProtoMessage()    {}
func (*L2TpSessionXconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8422bc409d0e143e, []int{1}
}

func (m *L2TpSessionXconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpSessionXconnect.Unmarshal(m, b)
}
func (m *L2TpSessionXconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpSessionXconnect.Marshal(b, m, deterministic)
}
func (m *L2TpSessionXconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpSessionXconnect.Merge(m, src)
}
func (m *L2TpSessionXconnect) XXX_Size() int {
	return xxx_messageInfo_L2TpSessionXconnect.Size(m)
}
func (m *L2TpSessionXconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpSessionXconnect.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpSessionXconnect proto.InternalMessageInfo

func (m *L2TpSessionXconnect) GetCircuitName() string {
	if m != nil {
		return m.CircuitName
	}
	return ""
}

func (m *L2TpSessionXconnect) GetSessionvcId() uint32 {
	if m != nil {
		return m.SessionvcId
	}
	return 0
}

func (m *L2TpSessionXconnect) GetIsCircuitStateUp() bool {
	if m != nil {
		return m.IsCircuitStateUp
	}
	return false
}

func (m *L2TpSessionXconnect) GetIsLocalCircuitStateUp() bool {
	if m != nil {
		return m.IsLocalCircuitStateUp
	}
	return false
}

func (m *L2TpSessionXconnect) GetIsRemoteCircuitStateUp() bool {
	if m != nil {
		return m.IsRemoteCircuitStateUp
	}
	return false
}

func (m *L2TpSessionXconnect) GetIpv6ProtocolTunneling() bool {
	if m != nil {
		return m.Ipv6ProtocolTunneling
	}
	return false
}

type L2TpSessionVpdn struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpSessionVpdn) Reset()         { *m = L2TpSessionVpdn{} }
func (m *L2TpSessionVpdn) String() string { return proto.CompactTextString(m) }
func (*L2TpSessionVpdn) ProtoMessage()    {}
func (*L2TpSessionVpdn) Descriptor() ([]byte, []int) {
	return fileDescriptor_8422bc409d0e143e, []int{2}
}

func (m *L2TpSessionVpdn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpSessionVpdn.Unmarshal(m, b)
}
func (m *L2TpSessionVpdn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpSessionVpdn.Marshal(b, m, deterministic)
}
func (m *L2TpSessionVpdn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpSessionVpdn.Merge(m, src)
}
func (m *L2TpSessionVpdn) XXX_Size() int {
	return xxx_messageInfo_L2TpSessionVpdn.Size(m)
}
func (m *L2TpSessionVpdn) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpSessionVpdn.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpSessionVpdn proto.InternalMessageInfo

func (m *L2TpSessionVpdn) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *L2TpSessionVpdn) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type SessionApp struct {
	L2TpShSessAppType    uint32               `protobuf:"varint,1,opt,name=l2tp_sh_sess_app_type,json=l2tpShSessAppType,proto3" json:"l2tp_sh_sess_app_type,omitempty"`
	Xconnect             *L2TpSessionXconnect `protobuf:"bytes,2,opt,name=xconnect,proto3" json:"xconnect,omitempty"`
	Vpdn                 *L2TpSessionVpdn     `protobuf:"bytes,3,opt,name=vpdn,proto3" json:"vpdn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SessionApp) Reset()         { *m = SessionApp{} }
func (m *SessionApp) String() string { return proto.CompactTextString(m) }
func (*SessionApp) ProtoMessage()    {}
func (*SessionApp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8422bc409d0e143e, []int{3}
}

func (m *SessionApp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionApp.Unmarshal(m, b)
}
func (m *SessionApp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionApp.Marshal(b, m, deterministic)
}
func (m *SessionApp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionApp.Merge(m, src)
}
func (m *SessionApp) XXX_Size() int {
	return xxx_messageInfo_SessionApp.Size(m)
}
func (m *SessionApp) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionApp.DiscardUnknown(m)
}

var xxx_messageInfo_SessionApp proto.InternalMessageInfo

func (m *SessionApp) GetL2TpShSessAppType() uint32 {
	if m != nil {
		return m.L2TpShSessAppType
	}
	return 0
}

func (m *SessionApp) GetXconnect() *L2TpSessionXconnect {
	if m != nil {
		return m.Xconnect
	}
	return nil
}

func (m *SessionApp) GetVpdn() *L2TpSessionVpdn {
	if m != nil {
		return m.Vpdn
	}
	return nil
}

type L2TpSessionInfo struct {
	LocalIpAddress              string      `protobuf:"bytes,50,opt,name=local_ip_address,json=localIpAddress,proto3" json:"local_ip_address,omitempty"`
	RemoteIpAddress             string      `protobuf:"bytes,51,opt,name=remote_ip_address,json=remoteIpAddress,proto3" json:"remote_ip_address,omitempty"`
	L2TpShSessUdpLport          uint32      `protobuf:"varint,52,opt,name=l2tp_sh_sess_udp_lport,json=l2tpShSessUdpLport,proto3" json:"l2tp_sh_sess_udp_lport,omitempty"`
	L2TpShSessUdpRport          uint32      `protobuf:"varint,53,opt,name=l2tp_sh_sess_udp_rport,json=l2tpShSessUdpRport,proto3" json:"l2tp_sh_sess_udp_rport,omitempty"`
	Protocol                    uint32      `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	RemoteTunnelId              uint32      `protobuf:"varint,55,opt,name=remote_tunnel_id,json=remoteTunnelId,proto3" json:"remote_tunnel_id,omitempty"`
	CallSerialNumber            uint32      `protobuf:"varint,56,opt,name=call_serial_number,json=callSerialNumber,proto3" json:"call_serial_number,omitempty"`
	LocalTunnelName             string      `protobuf:"bytes,57,opt,name=local_tunnel_name,json=localTunnelName,proto3" json:"local_tunnel_name,omitempty"`
	RemoteTunnelName            string      `protobuf:"bytes,58,opt,name=remote_tunnel_name,json=remoteTunnelName,proto3" json:"remote_tunnel_name,omitempty"`
	RemoteSessionId             uint32      `protobuf:"varint,59,opt,name=remote_session_id,json=remoteSessionId,proto3" json:"remote_session_id,omitempty"`
	L2TpShSessTieBreakerEnabled uint32      `protobuf:"varint,60,opt,name=l2tp_sh_sess_tie_breaker_enabled,json=l2tpShSessTieBreakerEnabled,proto3" json:"l2tp_sh_sess_tie_breaker_enabled,omitempty"`
	L2TpShSessTieBreaker        uint64      `protobuf:"varint,61,opt,name=l2tp_sh_sess_tie_breaker,json=l2tpShSessTieBreaker,proto3" json:"l2tp_sh_sess_tie_breaker,omitempty"`
	IsSessionManual             bool        `protobuf:"varint,62,opt,name=is_session_manual,json=isSessionManual,proto3" json:"is_session_manual,omitempty"`
	IsSessionUp                 bool        `protobuf:"varint,63,opt,name=is_session_up,json=isSessionUp,proto3" json:"is_session_up,omitempty"`
	IsUdpChecksumEnabled        bool        `protobuf:"varint,64,opt,name=is_udp_checksum_enabled,json=isUdpChecksumEnabled,proto3" json:"is_udp_checksum_enabled,omitempty"`
	IsSequencingOn              bool        `protobuf:"varint,65,opt,name=is_sequencing_on,json=isSequencingOn,proto3" json:"is_sequencing_on,omitempty"`
	IsSessionStateEstablished   bool        `protobuf:"varint,66,opt,name=is_session_state_established,json=isSessionStateEstablished,proto3" json:"is_session_state_established,omitempty"`
	IsSessionLocallyInitiated   bool        `protobuf:"varint,67,opt,name=is_session_locally_initiated,json=isSessionLocallyInitiated,proto3" json:"is_session_locally_initiated,omitempty"`
	IsConditionalDebugEnabled   bool        `protobuf:"varint,68,opt,name=is_conditional_debug_enabled,json=isConditionalDebugEnabled,proto3" json:"is_conditional_debug_enabled,omitempty"`
	UniqueId                    uint32      `protobuf:"varint,69,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	InterfaceName               string      `protobuf:"bytes,70,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SessionApplicationData      *SessionApp `protobuf:"bytes,71,opt,name=session_application_data,json=sessionApplicationData,proto3" json:"session_application_data,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}    `json:"-"`
	XXX_unrecognized            []byte      `json:"-"`
	XXX_sizecache               int32       `json:"-"`
}

func (m *L2TpSessionInfo) Reset()         { *m = L2TpSessionInfo{} }
func (m *L2TpSessionInfo) String() string { return proto.CompactTextString(m) }
func (*L2TpSessionInfo) ProtoMessage()    {}
func (*L2TpSessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8422bc409d0e143e, []int{4}
}

func (m *L2TpSessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpSessionInfo.Unmarshal(m, b)
}
func (m *L2TpSessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpSessionInfo.Marshal(b, m, deterministic)
}
func (m *L2TpSessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpSessionInfo.Merge(m, src)
}
func (m *L2TpSessionInfo) XXX_Size() int {
	return xxx_messageInfo_L2TpSessionInfo.Size(m)
}
func (m *L2TpSessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpSessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpSessionInfo proto.InternalMessageInfo

func (m *L2TpSessionInfo) GetLocalIpAddress() string {
	if m != nil {
		return m.LocalIpAddress
	}
	return ""
}

func (m *L2TpSessionInfo) GetRemoteIpAddress() string {
	if m != nil {
		return m.RemoteIpAddress
	}
	return ""
}

func (m *L2TpSessionInfo) GetL2TpShSessUdpLport() uint32 {
	if m != nil {
		return m.L2TpShSessUdpLport
	}
	return 0
}

func (m *L2TpSessionInfo) GetL2TpShSessUdpRport() uint32 {
	if m != nil {
		return m.L2TpShSessUdpRport
	}
	return 0
}

func (m *L2TpSessionInfo) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *L2TpSessionInfo) GetRemoteTunnelId() uint32 {
	if m != nil {
		return m.RemoteTunnelId
	}
	return 0
}

func (m *L2TpSessionInfo) GetCallSerialNumber() uint32 {
	if m != nil {
		return m.CallSerialNumber
	}
	return 0
}

func (m *L2TpSessionInfo) GetLocalTunnelName() string {
	if m != nil {
		return m.LocalTunnelName
	}
	return ""
}

func (m *L2TpSessionInfo) GetRemoteTunnelName() string {
	if m != nil {
		return m.RemoteTunnelName
	}
	return ""
}

func (m *L2TpSessionInfo) GetRemoteSessionId() uint32 {
	if m != nil {
		return m.RemoteSessionId
	}
	return 0
}

func (m *L2TpSessionInfo) GetL2TpShSessTieBreakerEnabled() uint32 {
	if m != nil {
		return m.L2TpShSessTieBreakerEnabled
	}
	return 0
}

func (m *L2TpSessionInfo) GetL2TpShSessTieBreaker() uint64 {
	if m != nil {
		return m.L2TpShSessTieBreaker
	}
	return 0
}

func (m *L2TpSessionInfo) GetIsSessionManual() bool {
	if m != nil {
		return m.IsSessionManual
	}
	return false
}

func (m *L2TpSessionInfo) GetIsSessionUp() bool {
	if m != nil {
		return m.IsSessionUp
	}
	return false
}

func (m *L2TpSessionInfo) GetIsUdpChecksumEnabled() bool {
	if m != nil {
		return m.IsUdpChecksumEnabled
	}
	return false
}

func (m *L2TpSessionInfo) GetIsSequencingOn() bool {
	if m != nil {
		return m.IsSequencingOn
	}
	return false
}

func (m *L2TpSessionInfo) GetIsSessionStateEstablished() bool {
	if m != nil {
		return m.IsSessionStateEstablished
	}
	return false
}

func (m *L2TpSessionInfo) GetIsSessionLocallyInitiated() bool {
	if m != nil {
		return m.IsSessionLocallyInitiated
	}
	return false
}

func (m *L2TpSessionInfo) GetIsConditionalDebugEnabled() bool {
	if m != nil {
		return m.IsConditionalDebugEnabled
	}
	return false
}

func (m *L2TpSessionInfo) GetUniqueId() uint32 {
	if m != nil {
		return m.UniqueId
	}
	return 0
}

func (m *L2TpSessionInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2TpSessionInfo) GetSessionApplicationData() *SessionApp {
	if m != nil {
		return m.SessionApplicationData
	}
	return nil
}

func init() {
	proto.RegisterType((*L2TpSessionInfo_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.sessions.session.l2tp_session_info_KEYS")
	proto.RegisterType((*L2TpSessionXconnect)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.sessions.session.l2tp_session_xconnect")
	proto.RegisterType((*L2TpSessionVpdn)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.sessions.session.l2tp_session_vpdn")
	proto.RegisterType((*SessionApp)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.sessions.session.session_app")
	proto.RegisterType((*L2TpSessionInfo)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.sessions.session.l2tp_session_info")
}

func init() { proto.RegisterFile("l2tp_session_info.proto", fileDescriptor_8422bc409d0e143e) }

var fileDescriptor_8422bc409d0e143e = []byte{
	// 855 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6f, 0x1b, 0x45,
	0x10, 0x57, 0x42, 0xa8, 0x92, 0x35, 0x49, 0xdc, 0xa5, 0x4d, 0xb6, 0x29, 0x0f, 0xc1, 0x12, 0xc8,
	0x42, 0xc5, 0x42, 0x2e, 0x35, 0xa5, 0x7c, 0x94, 0xd4, 0x31, 0xc8, 0x22, 0x2d, 0xe8, 0xec, 0x20,
	0xf1, 0x80, 0x56, 0xeb, 0xdb, 0x6d, 0x3c, 0xf4, 0xbc, 0xb7, 0xbd, 0xdd, 0xb3, 0x9a, 0x07, 0x1e,
	0x78, 0xe4, 0x5f, 0xe2, 0xaf, 0x43, 0x3b, 0x7b, 0x5f, 0x76, 0xcc, 0x4b, 0xd5, 0xb7, 0xbb, 0x99,
	0xdf, 0x6f, 0x66, 0xf6, 0x37, 0x1f, 0xe4, 0x38, 0xe9, 0x3b, 0xc3, 0xad, 0xb2, 0x16, 0x52, 0xcd,
	0x41, 0xbf, 0x4c, 0x7b, 0x26, 0x4b, 0x5d, 0x4a, 0x07, 0x31, 0xd8, 0x38, 0xe5, 0x90, 0x5a, 0xfe,
	0x26, 0xe3, 0x2e, 0xd7, 0x5a, 0x25, 0x3c, 0xe9, 0xbb, 0x5c, 0xf3, 0xd4, 0xa8, 0xac, 0xe7, 0x79,
	0xcb, 0x7e, 0xaf, 0x60, 0xda, 0xf2, 0xa3, 0xf3, 0x27, 0x39, 0xba, 0x11, 0x92, 0xff, 0x3c, 0xfa,
	0x7d, 0x42, 0x3f, 0x25, 0x87, 0x49, 0x1a, 0x8b, 0xa4, 0x0c, 0x06, 0x92, 0x6d, 0x9d, 0x6e, 0x75,
	0xf7, 0xa3, 0x7d, 0x34, 0x4f, 0xd1, 0x3a, 0x96, 0xb4, 0x4b, 0xda, 0x01, 0x57, 0x85, 0x90, 0x6c,
	0x1b, 0x81, 0x07, 0x68, 0x9f, 0x04, 0xf3, 0x58, 0x76, 0xfe, 0xdd, 0x26, 0x77, 0x57, 0x92, 0xbd,
	0x89, 0x53, 0xad, 0x55, 0xec, 0xe8, 0xc7, 0xe4, 0x83, 0x18, 0xb2, 0x38, 0x07, 0xc7, 0xb5, 0x58,
	0x28, 0x4c, 0xb4, 0x17, 0xb5, 0x0a, 0xdb, 0x0b, 0xb1, 0x50, 0x1e, 0x52, 0xd0, 0x96, 0x71, 0x9d,
	0xa2, 0x55, 0xd9, 0xc6, 0x92, 0x7e, 0x4e, 0x3e, 0x04, 0xcb, 0xcb, 0x40, 0xd6, 0x09, 0xa7, 0x78,
	0x6e, 0xd8, 0x7b, 0xa7, 0x5b, 0xdd, 0xdd, 0xa8, 0x0d, 0x76, 0x18, 0x3c, 0x13, 0xef, 0xb8, 0x34,
	0xf4, 0x31, 0xb9, 0x07, 0x96, 0x87, 0xda, 0x6f, 0x90, 0x76, 0x90, 0x74, 0x17, 0xec, 0x85, 0xf7,
	0xaf, 0x31, 0x9f, 0x90, 0x13, 0xb0, 0x3c, 0x53, 0x8b, 0xd4, 0xa9, 0x9b, 0xd4, 0xf7, 0x91, 0x7a,
	0x04, 0x36, 0x42, 0xc0, 0x1a, 0x77, 0x40, 0x8e, 0xc1, 0x2c, 0x07, 0x1c, 0xdb, 0x16, 0xa7, 0xa5,
	0xbc, 0xa0, 0xaf, 0xd8, 0xad, 0x22, 0xa7, 0x59, 0x0e, 0x7e, 0x2d, 0xbc, 0xd3, 0xd2, 0xd9, 0xf9,
	0x8d, 0xdc, 0x5e, 0xd1, 0x6e, 0x69, 0xa4, 0xa6, 0x27, 0x64, 0x37, 0xb7, 0x2a, 0x6b, 0x68, 0x56,
	0xfd, 0xd3, 0x4f, 0xc8, 0x01, 0x68, 0xa7, 0xb2, 0x97, 0x22, 0x56, 0x41, 0xd5, 0x6d, 0x44, 0xec,
	0x57, 0x56, 0xaf, 0x6b, 0xe7, 0x9f, 0x6d, 0x52, 0x8a, 0xc8, 0x85, 0x31, 0xf4, 0x8b, 0xb2, 0x47,
	0x73, 0x4c, 0xe5, 0x6d, 0xdc, 0x5d, 0x1b, 0x55, 0x34, 0x1f, 0x8b, 0x98, 0xcc, 0x7d, 0x53, 0xcf,
	0x8c, 0x99, 0x5e, 0x1b, 0x45, 0x81, 0xec, 0x96, 0x8d, 0xc4, 0x14, 0xad, 0xfe, 0xf3, 0xde, 0xdb,
	0x4d, 0x63, 0x6f, 0xe3, 0x74, 0x44, 0x55, 0x78, 0xfa, 0x07, 0xd9, 0xf1, 0xef, 0xc6, 0x96, 0xb6,
	0xfa, 0xe3, 0x77, 0x92, 0xc6, 0x07, 0x8c, 0x30, 0x6c, 0xe7, 0xef, 0xbd, 0x35, 0x91, 0xfd, 0x36,
	0xd4, 0x03, 0x0e, 0x86, 0x0b, 0x29, 0x33, 0x65, 0x2d, 0xeb, 0xa3, 0x94, 0x61, 0xc0, 0xc7, 0xe6,
	0x2c, 0x58, 0xe9, 0x67, 0xe4, 0x76, 0x31, 0x14, 0x0d, 0xe8, 0x43, 0x84, 0x1e, 0x06, 0x47, 0x8d,
	0xed, 0x97, 0x8b, 0x57, 0xe8, 0x9c, 0x4b, 0xc3, 0x13, 0x93, 0x66, 0x8e, 0x7d, 0x89, 0x42, 0xd3,
	0x5a, 0xe8, 0x4b, 0x69, 0x2e, 0xbc, 0x67, 0x23, 0x27, 0x43, 0xce, 0xa3, 0x0d, 0x9c, 0x08, 0x39,
	0x27, 0x64, 0xb7, 0x1c, 0x35, 0x36, 0x40, 0x54, 0xf5, 0xef, 0x5f, 0x56, 0xd4, 0x5b, 0xef, 0xf8,
	0x57, 0x61, 0x75, 0x83, 0xbd, 0x5a, 0xf2, 0x07, 0x84, 0xc6, 0x22, 0xf1, 0x3b, 0x9e, 0x81, 0x48,
	0xb8, 0xce, 0x17, 0x33, 0x95, 0xb1, 0xc7, 0x88, 0x6d, 0x7b, 0xcf, 0x04, 0x1d, 0x2f, 0xd0, 0xee,
	0x75, 0x58, 0x39, 0x1d, 0x38, 0x7d, 0x5f, 0x07, 0x1d, 0x1a, 0xc7, 0x03, 0xf7, 0xfa, 0x01, 0xa1,
	0xab, 0x35, 0x20, 0xf8, 0x09, 0x82, 0xdb, 0xcd, 0x2a, 0x10, 0x5d, 0x2b, 0xdc, 0xb8, 0x36, 0xdf,
	0x60, 0x19, 0x85, 0xc2, 0xd5, 0xb9, 0xa1, 0x23, 0x72, 0xba, 0xa2, 0x96, 0x03, 0xc5, 0x67, 0x99,
	0x12, 0xaf, 0x54, 0xc6, 0x95, 0x16, 0xb3, 0x44, 0x49, 0xf6, 0x2d, 0x52, 0xef, 0xd7, 0xba, 0x4d,
	0x41, 0x3d, 0x0b, 0x98, 0x51, 0x80, 0xd0, 0x01, 0x61, 0xff, 0x17, 0x86, 0x7d, 0x77, 0xba, 0xd5,
	0xdd, 0x89, 0xee, 0x6c, 0xa2, 0xfb, 0x52, 0xc1, 0x56, 0x65, 0x2e, 0x84, 0xce, 0x45, 0xc2, 0xbe,
	0xc7, 0x15, 0x3f, 0x04, 0x5b, 0x94, 0xf9, 0x1c, 0xcd, 0xb4, 0x43, 0xf6, 0x1b, 0xd8, 0xdc, 0xb0,
	0xa7, 0x88, 0x6b, 0x55, 0xb8, 0x4b, 0x43, 0x1f, 0x91, 0x63, 0x08, 0x2d, 0x8f, 0xe7, 0x2a, 0x7e,
	0x65, 0xf3, 0x45, 0xf5, 0x8a, 0x1f, 0x10, 0x7d, 0x07, 0x7c, 0xd7, 0x87, 0x85, 0xb3, 0x2c, 0xbf,
	0x4b, 0xda, 0x18, 0xfa, 0x75, 0xae, 0x74, 0x0c, 0xfa, 0x8a, 0xa7, 0x9a, 0x9d, 0x21, 0xfe, 0xc0,
	0x47, 0x2f, 0xcd, 0xbf, 0x68, 0xfa, 0x94, 0x7c, 0xd4, 0x28, 0x22, 0x9c, 0x33, 0x65, 0x9d, 0x98,
	0x25, 0x60, 0xe7, 0x4a, 0xb2, 0x67, 0xc8, 0xba, 0x57, 0xd5, 0x84, 0x17, 0x6d, 0x54, 0x03, 0xd6,
	0x02, 0x60, 0xa3, 0x93, 0x6b, 0x0e, 0x1a, 0x1c, 0x08, 0xa7, 0x24, 0x1b, 0xae, 0x05, 0xb8, 0x08,
	0x88, 0x71, 0x09, 0x28, 0x02, 0xc4, 0xa9, 0x96, 0xe0, 0x20, 0xd5, 0x22, 0xe1, 0x52, 0xcd, 0xf2,
	0xab, 0xea, 0x9d, 0xe7, 0x65, 0x80, 0x61, 0x0d, 0x39, 0xf7, 0x88, 0xf2, 0xb1, 0xf7, 0xc9, 0x5e,
	0xae, 0xe1, 0x75, 0xae, 0xfc, 0x58, 0x8c, 0xc2, 0xb4, 0x07, 0xc3, 0x58, 0x6e, 0x38, 0x88, 0x3f,
	0x6e, 0x38, 0x88, 0xf4, 0x2f, 0xc2, 0x1a, 0xf7, 0x30, 0x81, 0x58, 0xf8, 0x44, 0x5c, 0x0a, 0x27,
	0xd8, 0x4f, 0x78, 0x77, 0x86, 0x6f, 0x7b, 0x77, 0x1a, 0x71, 0xa3, 0xa3, 0xe2, 0xe7, 0xac, 0xce,
	0x71, 0x2e, 0x9c, 0x98, 0xdd, 0xc2, 0xed, 0x7c, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1,
	0x94, 0xd0, 0xf6, 0xea, 0x07, 0x00, 0x00,
}
