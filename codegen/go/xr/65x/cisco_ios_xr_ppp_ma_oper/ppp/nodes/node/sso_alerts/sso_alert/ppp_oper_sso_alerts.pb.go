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
// source: ppp_oper_sso_alerts.proto

package cisco_ios_xr_ppp_ma_oper_ppp_nodes_node_sso_alerts_sso_alert

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

type PppOperSsoAlerts_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PppOperSsoAlerts_KEYS) Reset()         { *m = PppOperSsoAlerts_KEYS{} }
func (m *PppOperSsoAlerts_KEYS) String() string { return proto.CompactTextString(m) }
func (*PppOperSsoAlerts_KEYS) ProtoMessage()    {}
func (*PppOperSsoAlerts_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d53424786a566, []int{0}
}

func (m *PppOperSsoAlerts_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppOperSsoAlerts_KEYS.Unmarshal(m, b)
}
func (m *PppOperSsoAlerts_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppOperSsoAlerts_KEYS.Marshal(b, m, deterministic)
}
func (m *PppOperSsoAlerts_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppOperSsoAlerts_KEYS.Merge(m, src)
}
func (m *PppOperSsoAlerts_KEYS) XXX_Size() int {
	return xxx_messageInfo_PppOperSsoAlerts_KEYS.Size(m)
}
func (m *PppOperSsoAlerts_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PppOperSsoAlerts_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PppOperSsoAlerts_KEYS proto.InternalMessageInfo

func (m *PppOperSsoAlerts_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PppOperSsoAlerts_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

type PppOperSsoError struct {
	IsError              bool     `protobuf:"varint,1,opt,name=is_error,json=isError,proto3" json:"is_error,omitempty"`
	Error                uint32   `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	Context              uint32   `protobuf:"varint,3,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PppOperSsoError) Reset()         { *m = PppOperSsoError{} }
func (m *PppOperSsoError) String() string { return proto.CompactTextString(m) }
func (*PppOperSsoError) ProtoMessage()    {}
func (*PppOperSsoError) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d53424786a566, []int{1}
}

func (m *PppOperSsoError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppOperSsoError.Unmarshal(m, b)
}
func (m *PppOperSsoError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppOperSsoError.Marshal(b, m, deterministic)
}
func (m *PppOperSsoError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppOperSsoError.Merge(m, src)
}
func (m *PppOperSsoError) XXX_Size() int {
	return xxx_messageInfo_PppOperSsoError.Size(m)
}
func (m *PppOperSsoError) XXX_DiscardUnknown() {
	xxx_messageInfo_PppOperSsoError.DiscardUnknown(m)
}

var xxx_messageInfo_PppOperSsoError proto.InternalMessageInfo

func (m *PppOperSsoError) GetIsError() bool {
	if m != nil {
		return m.IsError
	}
	return false
}

func (m *PppOperSsoError) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *PppOperSsoError) GetContext() uint32 {
	if m != nil {
		return m.Context
	}
	return 0
}

type PppOperSsoAlerts struct {
	LcpError             *PppOperSsoError `protobuf:"bytes,50,opt,name=lcp_error,json=lcpError,proto3" json:"lcp_error,omitempty"`
	OfUsAuthError        *PppOperSsoError `protobuf:"bytes,51,opt,name=of_us_auth_error,json=ofUsAuthError,proto3" json:"of_us_auth_error,omitempty"`
	OfPeerAuthError      *PppOperSsoError `protobuf:"bytes,52,opt,name=of_peer_auth_error,json=ofPeerAuthError,proto3" json:"of_peer_auth_error,omitempty"`
	IpcpError            *PppOperSsoError `protobuf:"bytes,53,opt,name=ipcp_error,json=ipcpError,proto3" json:"ipcp_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PppOperSsoAlerts) Reset()         { *m = PppOperSsoAlerts{} }
func (m *PppOperSsoAlerts) String() string { return proto.CompactTextString(m) }
func (*PppOperSsoAlerts) ProtoMessage()    {}
func (*PppOperSsoAlerts) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d53424786a566, []int{2}
}

func (m *PppOperSsoAlerts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PppOperSsoAlerts.Unmarshal(m, b)
}
func (m *PppOperSsoAlerts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PppOperSsoAlerts.Marshal(b, m, deterministic)
}
func (m *PppOperSsoAlerts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PppOperSsoAlerts.Merge(m, src)
}
func (m *PppOperSsoAlerts) XXX_Size() int {
	return xxx_messageInfo_PppOperSsoAlerts.Size(m)
}
func (m *PppOperSsoAlerts) XXX_DiscardUnknown() {
	xxx_messageInfo_PppOperSsoAlerts.DiscardUnknown(m)
}

var xxx_messageInfo_PppOperSsoAlerts proto.InternalMessageInfo

func (m *PppOperSsoAlerts) GetLcpError() *PppOperSsoError {
	if m != nil {
		return m.LcpError
	}
	return nil
}

func (m *PppOperSsoAlerts) GetOfUsAuthError() *PppOperSsoError {
	if m != nil {
		return m.OfUsAuthError
	}
	return nil
}

func (m *PppOperSsoAlerts) GetOfPeerAuthError() *PppOperSsoError {
	if m != nil {
		return m.OfPeerAuthError
	}
	return nil
}

func (m *PppOperSsoAlerts) GetIpcpError() *PppOperSsoError {
	if m != nil {
		return m.IpcpError
	}
	return nil
}

func init() {
	proto.RegisterType((*PppOperSsoAlerts_KEYS)(nil), "cisco_ios_xr_ppp_ma_oper.ppp.nodes.node.sso_alerts.sso_alert.ppp_oper_sso_alerts_KEYS")
	proto.RegisterType((*PppOperSsoError)(nil), "cisco_ios_xr_ppp_ma_oper.ppp.nodes.node.sso_alerts.sso_alert.ppp_oper_sso_error")
	proto.RegisterType((*PppOperSsoAlerts)(nil), "cisco_ios_xr_ppp_ma_oper.ppp.nodes.node.sso_alerts.sso_alert.ppp_oper_sso_alerts")
}

func init() { proto.RegisterFile("ppp_oper_sso_alerts.proto", fileDescriptor_751d53424786a566) }

var fileDescriptor_751d53424786a566 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xb1, 0x4e, 0xc3, 0x40,
	0x0c, 0x86, 0x55, 0x2a, 0x68, 0x62, 0x54, 0x81, 0x0e, 0x86, 0x54, 0x30, 0x54, 0x9d, 0x3a, 0x65,
	0x68, 0x61, 0x63, 0x61, 0xe8, 0x84, 0x84, 0xaa, 0xa0, 0x0e, 0x4c, 0x56, 0x38, 0x1c, 0xf5, 0xa4,
	0x26, 0xb6, 0xee, 0xae, 0x52, 0x19, 0x78, 0x66, 0x5e, 0x01, 0xe5, 0x52, 0xda, 0x22, 0xba, 0xa1,
	0x2c, 0x91, 0xfd, 0xdb, 0xfe, 0x3f, 0xc7, 0x3a, 0x18, 0x88, 0x08, 0xb2, 0x90, 0x45, 0xe7, 0x18,
	0xf3, 0x15, 0x59, 0xef, 0x52, 0xb1, 0xec, 0x59, 0x3d, 0x68, 0xe3, 0x34, 0xa3, 0x61, 0x87, 0x1b,
	0x8b, 0x75, 0x5f, 0x99, 0x87, 0xd6, 0x54, 0x44, 0xd2, 0x8a, 0xdf, 0xc9, 0x85, 0x6f, 0x7a, 0x30,
	0xb9, 0x0b, 0x47, 0x0b, 0x48, 0x8e, 0x58, 0xe3, 0xd3, 0xec, 0xf5, 0x45, 0xdd, 0x40, 0x5c, 0x0f,
	0x62, 0x95, 0x97, 0x94, 0x74, 0x86, 0x9d, 0x71, 0x9c, 0x45, 0xb5, 0xf0, 0x9c, 0x97, 0xa4, 0x6e,
	0x21, 0x36, 0x95, 0x27, 0x5b, 0xe4, 0x9a, 0x92, 0x93, 0x50, 0xdc, 0x0b, 0x23, 0x04, 0xf5, 0xcb,
	0x96, 0xac, 0x65, 0xab, 0x06, 0x10, 0x19, 0xd7, 0xc4, 0xc1, 0x2f, 0xca, 0x7a, 0xc6, 0xcd, 0x42,
	0xe9, 0x1a, 0x4e, 0x1b, 0xbd, 0xb6, 0xea, 0x67, 0x4d, 0xa2, 0x12, 0xe8, 0x69, 0xae, 0x3c, 0x6d,
	0x7c, 0xd2, 0x0d, 0xfa, 0x4f, 0x3a, 0xfa, 0xea, 0xc2, 0xd5, 0x91, 0xc5, 0x55, 0x09, 0xf1, 0x4a,
	0xcb, 0x96, 0x31, 0x19, 0x76, 0xc6, 0xe7, 0x93, 0x79, 0xfa, 0x9f, 0x0b, 0xa5, 0x7f, 0xff, 0x23,
	0x8b, 0x56, 0x5a, 0x9a, 0xb5, 0x3f, 0xe0, 0x92, 0x0b, 0x5c, 0x3b, 0xcc, 0xd7, 0x7e, 0xb9, 0xa5,
	0x4e, 0x5b, 0xa2, 0xf6, 0xb9, 0x58, 0xb8, 0xc7, 0xb5, 0x5f, 0x36, 0xe8, 0x4f, 0x50, 0x5c, 0xa0,
	0x10, 0xd9, 0x43, 0xf8, 0x5d, 0x4b, 0xf0, 0x0b, 0x2e, 0xe6, 0x44, 0x76, 0x8f, 0x67, 0x00, 0x23,
	0xbb, 0x4b, 0xdf, 0xb7, 0x84, 0x8d, 0x6b, 0x46, 0x00, 0xbe, 0x9d, 0x85, 0xe7, 0x3e, 0xfd, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xd6, 0xbc, 0xf1, 0x57, 0x0b, 0x03, 0x00, 0x00,
}
