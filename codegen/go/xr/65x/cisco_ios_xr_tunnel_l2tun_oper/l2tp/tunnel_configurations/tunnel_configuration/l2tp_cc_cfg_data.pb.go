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
// source: l2tp_cc_cfg_data.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tp_tunnel_configurations_tunnel_configuration

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

type L2TpCcCfgData_KEYS struct {
	LocalTunnelId        uint32   `protobuf:"varint,1,opt,name=local_tunnel_id,json=localTunnelId,proto3" json:"local_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpCcCfgData_KEYS) Reset()         { *m = L2TpCcCfgData_KEYS{} }
func (m *L2TpCcCfgData_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2TpCcCfgData_KEYS) ProtoMessage()    {}
func (*L2TpCcCfgData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_660d2a9abf46f742, []int{0}
}

func (m *L2TpCcCfgData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcCfgData_KEYS.Unmarshal(m, b)
}
func (m *L2TpCcCfgData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcCfgData_KEYS.Marshal(b, m, deterministic)
}
func (m *L2TpCcCfgData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcCfgData_KEYS.Merge(m, src)
}
func (m *L2TpCcCfgData_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2TpCcCfgData_KEYS.Size(m)
}
func (m *L2TpCcCfgData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcCfgData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcCfgData_KEYS proto.InternalMessageInfo

func (m *L2TpCcCfgData_KEYS) GetLocalTunnelId() uint32 {
	if m != nil {
		return m.LocalTunnelId
	}
	return 0
}

type L2TpCcClassData struct {
	IpTos                           uint32   `protobuf:"varint,1,opt,name=ip_tos,json=ipTos,proto3" json:"ip_tos,omitempty"`
	VrfName                         string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ReceiveWindowSize               uint32   `protobuf:"varint,3,opt,name=receive_window_size,json=receiveWindowSize,proto3" json:"receive_window_size,omitempty"`
	ClassNameXr                     string   `protobuf:"bytes,4,opt,name=class_name_xr,json=classNameXr,proto3" json:"class_name_xr,omitempty"`
	DigestHash                      string   `protobuf:"bytes,5,opt,name=digest_hash,json=digestHash,proto3" json:"digest_hash,omitempty"`
	Password                        string   `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	EncodedPassword                 string   `protobuf:"bytes,7,opt,name=encoded_password,json=encodedPassword,proto3" json:"encoded_password,omitempty"`
	HostName                        string   `protobuf:"bytes,8,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	AccountingMethodList            string   `protobuf:"bytes,9,opt,name=accounting_method_list,json=accountingMethodList,proto3" json:"accounting_method_list,omitempty"`
	HelloTimeout                    uint32   `protobuf:"varint,10,opt,name=hello_timeout,json=helloTimeout,proto3" json:"hello_timeout,omitempty"`
	SetupTimeout                    uint32   `protobuf:"varint,11,opt,name=setup_timeout,json=setupTimeout,proto3" json:"setup_timeout,omitempty"`
	RetransmitMinimumTimeout        uint32   `protobuf:"varint,12,opt,name=retransmit_minimum_timeout,json=retransmitMinimumTimeout,proto3" json:"retransmit_minimum_timeout,omitempty"`
	RetransmitMaximumTimeout        uint32   `protobuf:"varint,13,opt,name=retransmit_maximum_timeout,json=retransmitMaximumTimeout,proto3" json:"retransmit_maximum_timeout,omitempty"`
	InitialRetransmitMinimumTimeout uint32   `protobuf:"varint,14,opt,name=initial_retransmit_minimum_timeout,json=initialRetransmitMinimumTimeout,proto3" json:"initial_retransmit_minimum_timeout,omitempty"`
	InitialRetransmitMaximumTimeout uint32   `protobuf:"varint,15,opt,name=initial_retransmit_maximum_timeout,json=initialRetransmitMaximumTimeout,proto3" json:"initial_retransmit_maximum_timeout,omitempty"`
	TimeoutNoUser                   uint32   `protobuf:"varint,16,opt,name=timeout_no_user,json=timeoutNoUser,proto3" json:"timeout_no_user,omitempty"`
	RetransmitRetries               uint32   `protobuf:"varint,17,opt,name=retransmit_retries,json=retransmitRetries,proto3" json:"retransmit_retries,omitempty"`
	InitialRetransmitRetries        uint32   `protobuf:"varint,18,opt,name=initial_retransmit_retries,json=initialRetransmitRetries,proto3" json:"initial_retransmit_retries,omitempty"`
	IsAuthenticationEnabled         bool     `protobuf:"varint,19,opt,name=is_authentication_enabled,json=isAuthenticationEnabled,proto3" json:"is_authentication_enabled,omitempty"`
	IsHidden                        bool     `protobuf:"varint,20,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden,omitempty"`
	IsDigestEnabled                 bool     `protobuf:"varint,21,opt,name=is_digest_enabled,json=isDigestEnabled,proto3" json:"is_digest_enabled,omitempty"`
	IsDigestCheckEnabled            bool     `protobuf:"varint,22,opt,name=is_digest_check_enabled,json=isDigestCheckEnabled,proto3" json:"is_digest_check_enabled,omitempty"`
	IsCongestionControlEnabled      bool     `protobuf:"varint,23,opt,name=is_congestion_control_enabled,json=isCongestionControlEnabled,proto3" json:"is_congestion_control_enabled,omitempty"`
	IsPeerAddressChecked            bool     `protobuf:"varint,24,opt,name=is_peer_address_checked,json=isPeerAddressChecked,proto3" json:"is_peer_address_checked,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *L2TpCcClassData) Reset()         { *m = L2TpCcClassData{} }
func (m *L2TpCcClassData) String() string { return proto.CompactTextString(m) }
func (*L2TpCcClassData) ProtoMessage()    {}
func (*L2TpCcClassData) Descriptor() ([]byte, []int) {
	return fileDescriptor_660d2a9abf46f742, []int{1}
}

func (m *L2TpCcClassData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcClassData.Unmarshal(m, b)
}
func (m *L2TpCcClassData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcClassData.Marshal(b, m, deterministic)
}
func (m *L2TpCcClassData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcClassData.Merge(m, src)
}
func (m *L2TpCcClassData) XXX_Size() int {
	return xxx_messageInfo_L2TpCcClassData.Size(m)
}
func (m *L2TpCcClassData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcClassData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcClassData proto.InternalMessageInfo

func (m *L2TpCcClassData) GetIpTos() uint32 {
	if m != nil {
		return m.IpTos
	}
	return 0
}

func (m *L2TpCcClassData) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *L2TpCcClassData) GetReceiveWindowSize() uint32 {
	if m != nil {
		return m.ReceiveWindowSize
	}
	return 0
}

func (m *L2TpCcClassData) GetClassNameXr() string {
	if m != nil {
		return m.ClassNameXr
	}
	return ""
}

func (m *L2TpCcClassData) GetDigestHash() string {
	if m != nil {
		return m.DigestHash
	}
	return ""
}

func (m *L2TpCcClassData) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *L2TpCcClassData) GetEncodedPassword() string {
	if m != nil {
		return m.EncodedPassword
	}
	return ""
}

func (m *L2TpCcClassData) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *L2TpCcClassData) GetAccountingMethodList() string {
	if m != nil {
		return m.AccountingMethodList
	}
	return ""
}

func (m *L2TpCcClassData) GetHelloTimeout() uint32 {
	if m != nil {
		return m.HelloTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetSetupTimeout() uint32 {
	if m != nil {
		return m.SetupTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitMinimumTimeout() uint32 {
	if m != nil {
		return m.RetransmitMinimumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitMaximumTimeout() uint32 {
	if m != nil {
		return m.RetransmitMaximumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitMinimumTimeout() uint32 {
	if m != nil {
		return m.InitialRetransmitMinimumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitMaximumTimeout() uint32 {
	if m != nil {
		return m.InitialRetransmitMaximumTimeout
	}
	return 0
}

func (m *L2TpCcClassData) GetTimeoutNoUser() uint32 {
	if m != nil {
		return m.TimeoutNoUser
	}
	return 0
}

func (m *L2TpCcClassData) GetRetransmitRetries() uint32 {
	if m != nil {
		return m.RetransmitRetries
	}
	return 0
}

func (m *L2TpCcClassData) GetInitialRetransmitRetries() uint32 {
	if m != nil {
		return m.InitialRetransmitRetries
	}
	return 0
}

func (m *L2TpCcClassData) GetIsAuthenticationEnabled() bool {
	if m != nil {
		return m.IsAuthenticationEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsHidden() bool {
	if m != nil {
		return m.IsHidden
	}
	return false
}

func (m *L2TpCcClassData) GetIsDigestEnabled() bool {
	if m != nil {
		return m.IsDigestEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsDigestCheckEnabled() bool {
	if m != nil {
		return m.IsDigestCheckEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsCongestionControlEnabled() bool {
	if m != nil {
		return m.IsCongestionControlEnabled
	}
	return false
}

func (m *L2TpCcClassData) GetIsPeerAddressChecked() bool {
	if m != nil {
		return m.IsPeerAddressChecked
	}
	return false
}

type L2TpCcCfgData struct {
	RemoteTunnelId       uint32           `protobuf:"varint,50,opt,name=remote_tunnel_id,json=remoteTunnelId,proto3" json:"remote_tunnel_id,omitempty"`
	L2TpClass            *L2TpCcClassData `protobuf:"bytes,51,opt,name=l2tp_class,json=l2tpClass,proto3" json:"l2tp_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2TpCcCfgData) Reset()         { *m = L2TpCcCfgData{} }
func (m *L2TpCcCfgData) String() string { return proto.CompactTextString(m) }
func (*L2TpCcCfgData) ProtoMessage()    {}
func (*L2TpCcCfgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_660d2a9abf46f742, []int{2}
}

func (m *L2TpCcCfgData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcCfgData.Unmarshal(m, b)
}
func (m *L2TpCcCfgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcCfgData.Marshal(b, m, deterministic)
}
func (m *L2TpCcCfgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcCfgData.Merge(m, src)
}
func (m *L2TpCcCfgData) XXX_Size() int {
	return xxx_messageInfo_L2TpCcCfgData.Size(m)
}
func (m *L2TpCcCfgData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcCfgData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcCfgData proto.InternalMessageInfo

func (m *L2TpCcCfgData) GetRemoteTunnelId() uint32 {
	if m != nil {
		return m.RemoteTunnelId
	}
	return 0
}

func (m *L2TpCcCfgData) GetL2TpClass() *L2TpCcClassData {
	if m != nil {
		return m.L2TpClass
	}
	return nil
}

func init() {
	proto.RegisterType((*L2TpCcCfgData_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.tunnel_configurations.tunnel_configuration.l2tp_cc_cfg_data_KEYS")
	proto.RegisterType((*L2TpCcClassData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.tunnel_configurations.tunnel_configuration.l2tp_cc_class_data")
	proto.RegisterType((*L2TpCcCfgData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.tunnel_configurations.tunnel_configuration.l2tp_cc_cfg_data")
}

func init() { proto.RegisterFile("l2tp_cc_cfg_data.proto", fileDescriptor_660d2a9abf46f742) }

var fileDescriptor_660d2a9abf46f742 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x4f, 0xdb, 0x3a,
	0x14, 0x57, 0xef, 0xbd, 0x40, 0xeb, 0x52, 0x5a, 0xcc, 0x3f, 0xc3, 0xd5, 0x44, 0xd5, 0x49, 0xa8,
	0x9b, 0xb4, 0x3e, 0x94, 0xed, 0x65, 0x9a, 0x34, 0xa1, 0x0e, 0x89, 0x89, 0x81, 0x50, 0x61, 0xda,
	0xf6, 0x64, 0x99, 0xe4, 0xb4, 0x39, 0x5a, 0x62, 0x47, 0xb6, 0x03, 0x88, 0xa7, 0xed, 0x9b, 0xed,
	0xa3, 0x4d, 0xb1, 0x93, 0x40, 0x81, 0xee, 0x69, 0x6f, 0xf1, 0xef, 0xdf, 0x39, 0x39, 0xb2, 0x0f,
	0xd9, 0x8c, 0x87, 0x36, 0xe5, 0x41, 0xc0, 0x83, 0xc9, 0x94, 0x87, 0xc2, 0x8a, 0x41, 0xaa, 0x95,
	0x55, 0xf4, 0x34, 0x40, 0x13, 0x28, 0x8e, 0xca, 0xf0, 0x1b, 0xcd, 0x6d, 0x26, 0x25, 0xc4, 0x3c,
	0x1e, 0xda, 0x4c, 0x72, 0x95, 0x82, 0x1e, 0xe4, 0xb6, 0x41, 0x01, 0x07, 0x4a, 0x4e, 0x70, 0x9a,
	0x69, 0x61, 0x51, 0x49, 0xf3, 0x24, 0xda, 0x7b, 0x4f, 0x36, 0x1e, 0x56, 0xe2, 0xc7, 0x87, 0xdf,
	0xce, 0xe9, 0x1e, 0x69, 0xc7, 0x2a, 0x10, 0x71, 0x59, 0x03, 0x43, 0x56, 0xeb, 0xd6, 0xfa, 0xad,
	0x71, 0xcb, 0xc1, 0x17, 0x0e, 0xfd, 0x18, 0xf6, 0x7e, 0x34, 0x08, 0xad, 0x12, 0x62, 0x61, 0x8c,
	0xcb, 0xa0, 0x1b, 0x64, 0x11, 0x53, 0x6e, 0x95, 0x29, 0x5c, 0x0b, 0x98, 0x5e, 0x28, 0x43, 0xb7,
	0x49, 0xfd, 0x4a, 0x4f, 0xb8, 0x14, 0x09, 0xb0, 0x7f, 0xba, 0xb5, 0x7e, 0x63, 0xbc, 0x74, 0xa5,
	0x27, 0xa7, 0x22, 0x01, 0x3a, 0x20, 0x6b, 0x1a, 0x02, 0xc0, 0x2b, 0xe0, 0xd7, 0x28, 0x43, 0x75,
	0xcd, 0x0d, 0xde, 0x02, 0xfb, 0xd7, 0xd9, 0x57, 0x0b, 0xea, 0x8b, 0x63, 0xce, 0xf1, 0x16, 0x68,
	0x8f, 0xb4, 0x7c, 0xbd, 0x3c, 0x8c, 0xdf, 0x68, 0xf6, 0x9f, 0xcb, 0x6b, 0x3a, 0x30, 0x4f, 0xfc,
	0xaa, 0xe9, 0x2e, 0x69, 0x86, 0x38, 0x05, 0x63, 0x79, 0x24, 0x4c, 0xc4, 0x16, 0x9c, 0x82, 0x78,
	0xe8, 0x48, 0x98, 0x88, 0xee, 0x90, 0x7a, 0x2a, 0x8c, 0xb9, 0x56, 0x3a, 0x64, 0x8b, 0x8e, 0xad,
	0xce, 0xf4, 0x05, 0xe9, 0x80, 0x0c, 0x54, 0x08, 0x21, 0xaf, 0x34, 0x4b, 0x4e, 0xd3, 0x2e, 0xf0,
	0xb3, 0x52, 0xfa, 0x3f, 0x69, 0x44, 0xca, 0x58, 0xff, 0x5f, 0x75, 0x9f, 0x93, 0x03, 0xee, 0xc7,
	0x5e, 0x93, 0x4d, 0x11, 0x04, 0x2a, 0x93, 0x16, 0xe5, 0x94, 0x27, 0x60, 0x23, 0x15, 0xf2, 0x18,
	0x8d, 0x65, 0x0d, 0xa7, 0x5c, 0xbf, 0x63, 0x4f, 0x1c, 0xf9, 0x09, 0x8d, 0xa5, 0xcf, 0x49, 0x2b,
	0x82, 0x38, 0x56, 0xdc, 0x62, 0x02, 0x2a, 0xb3, 0x8c, 0xb8, 0x41, 0x2c, 0x3b, 0xf0, 0xc2, 0x63,
	0xb9, 0xc8, 0x80, 0xcd, 0xd2, 0x4a, 0xd4, 0xf4, 0x22, 0x07, 0x96, 0xa2, 0x77, 0x64, 0x47, 0x83,
	0xd5, 0x42, 0x9a, 0x04, 0x2d, 0x4f, 0x50, 0x62, 0x92, 0x25, 0x95, 0x63, 0xd9, 0x39, 0xd8, 0x9d,
	0xe2, 0xc4, 0x0b, 0xe6, 0xb8, 0xc5, 0xcd, 0x8c, 0xbb, 0xf5, 0xc8, 0xed, 0x05, 0xa5, 0xfb, 0x98,
	0xf4, 0x50, 0xa2, 0x45, 0x11, 0xf3, 0x3f, 0xf4, 0xb0, 0xe2, 0x52, 0x76, 0x0b, 0xe5, 0x78, 0x5e,
	0x2b, 0x73, 0xc2, 0x1e, 0xb4, 0xd4, 0x9e, 0x17, 0x36, 0xdb, 0xd9, 0x1e, 0x69, 0x17, 0x0e, 0x2e,
	0x15, 0xcf, 0x0c, 0x68, 0xd6, 0xf1, 0xf7, 0xbb, 0x80, 0x4f, 0xd5, 0x67, 0x03, 0x9a, 0xbe, 0x22,
	0xf4, 0x5e, 0xb1, 0xfc, 0x13, 0xc1, 0xb0, 0xd5, 0xf2, 0x56, 0x96, 0xcc, 0xd8, 0x13, 0xf9, 0xb8,
	0x9e, 0xe8, 0xb1, 0xb4, 0x51, 0x3f, 0xae, 0x47, 0xbd, 0x95, 0xee, 0xb7, 0x64, 0x1b, 0x0d, 0x17,
	0x99, 0x8d, 0x40, 0x5a, 0x0c, 0xdc, 0x13, 0xe5, 0x20, 0xc5, 0x65, 0x0c, 0x21, 0x5b, 0xeb, 0xd6,
	0xfa, 0xf5, 0xf1, 0x16, 0x9a, 0x83, 0x19, 0xfe, 0xd0, 0xd3, 0xf9, 0x1d, 0x44, 0xc3, 0x23, 0x0c,
	0x43, 0x90, 0x6c, 0xdd, 0x69, 0xeb, 0x68, 0x8e, 0xdc, 0x99, 0xbe, 0x24, 0xab, 0x68, 0x78, 0xf1,
	0x16, 0xca, 0xc0, 0x0d, 0x27, 0x6a, 0xa3, 0xf9, 0xe0, 0xf0, 0x32, 0xe8, 0x0d, 0xd9, 0xba, 0xd3,
	0x06, 0x11, 0x04, 0xdf, 0x2b, 0xc7, 0xa6, 0x73, 0xac, 0x97, 0x8e, 0x51, 0x4e, 0x96, 0xb6, 0x03,
	0xf2, 0x0c, 0x4d, 0xbe, 0x5d, 0x72, 0x26, 0xef, 0x3b, 0x50, 0xd2, 0x6a, 0x15, 0x57, 0xe6, 0x2d,
	0x67, 0xde, 0x41, 0x33, 0xaa, 0x34, 0x23, 0x2f, 0x99, 0xad, 0x9c, 0x02, 0x68, 0x2e, 0xc2, 0x50,
	0x83, 0x31, 0xbe, 0x3e, 0x84, 0x8c, 0x95, 0x95, 0xcf, 0x00, 0xf4, 0x81, 0x27, 0x47, 0x9e, 0xeb,
	0xfd, 0xaa, 0x91, 0xce, 0xc3, 0x25, 0x46, 0xfb, 0xa4, 0xa3, 0x21, 0x51, 0x16, 0xee, 0x2d, 0xb0,
	0xa1, 0x1b, 0xff, 0x8a, 0xc7, 0xcb, 0x0d, 0x46, 0x7f, 0xd6, 0x08, 0xf1, 0xf6, 0x7c, 0x73, 0xb0,
	0xfd, 0x6e, 0xad, 0xdf, 0x1c, 0x5e, 0x0e, 0xfe, 0xee, 0xa2, 0x1d, 0x3c, 0xde, 0x91, 0xe3, 0x46,
	0x8e, 0x8d, 0xf2, 0xf3, 0xe5, 0xa2, 0xdb, 0xee, 0xfb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x0a, 0x7e, 0x67, 0xf7, 0x05, 0x00, 0x00,
}
