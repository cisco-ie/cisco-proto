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
// source: ssh_session_detail.proto

package cisco_ios_xr_crypto_ssh_oper_ssh1_kex_nodes_node

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

type SshSessionDetail_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshSessionDetail_KEYS) Reset()         { *m = SshSessionDetail_KEYS{} }
func (m *SshSessionDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*SshSessionDetail_KEYS) ProtoMessage()    {}
func (*SshSessionDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2992548dd79a11f1, []int{0}
}

func (m *SshSessionDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshSessionDetail_KEYS.Unmarshal(m, b)
}
func (m *SshSessionDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshSessionDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *SshSessionDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshSessionDetail_KEYS.Merge(m, src)
}
func (m *SshSessionDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_SshSessionDetail_KEYS.Size(m)
}
func (m *SshSessionDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SshSessionDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SshSessionDetail_KEYS proto.InternalMessageInfo

func (m *SshSessionDetail_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type SessionDetailInfoItem struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	KeyExchange          string   `protobuf:"bytes,2,opt,name=key_exchange,json=keyExchange,proto3" json:"key_exchange,omitempty"`
	PublicKey            string   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InCipher             string   `protobuf:"bytes,4,opt,name=in_cipher,json=inCipher,proto3" json:"in_cipher,omitempty"`
	OutCipher            string   `protobuf:"bytes,5,opt,name=out_cipher,json=outCipher,proto3" json:"out_cipher,omitempty"`
	InMac                string   `protobuf:"bytes,6,opt,name=in_mac,json=inMac,proto3" json:"in_mac,omitempty"`
	OutMac               string   `protobuf:"bytes,7,opt,name=out_mac,json=outMac,proto3" json:"out_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionDetailInfoItem) Reset()         { *m = SessionDetailInfoItem{} }
func (m *SessionDetailInfoItem) String() string { return proto.CompactTextString(m) }
func (*SessionDetailInfoItem) ProtoMessage()    {}
func (*SessionDetailInfoItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2992548dd79a11f1, []int{1}
}

func (m *SessionDetailInfoItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDetailInfoItem.Unmarshal(m, b)
}
func (m *SessionDetailInfoItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDetailInfoItem.Marshal(b, m, deterministic)
}
func (m *SessionDetailInfoItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDetailInfoItem.Merge(m, src)
}
func (m *SessionDetailInfoItem) XXX_Size() int {
	return xxx_messageInfo_SessionDetailInfoItem.Size(m)
}
func (m *SessionDetailInfoItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDetailInfoItem.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDetailInfoItem proto.InternalMessageInfo

func (m *SessionDetailInfoItem) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SessionDetailInfoItem) GetKeyExchange() string {
	if m != nil {
		return m.KeyExchange
	}
	return ""
}

func (m *SessionDetailInfoItem) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *SessionDetailInfoItem) GetInCipher() string {
	if m != nil {
		return m.InCipher
	}
	return ""
}

func (m *SessionDetailInfoItem) GetOutCipher() string {
	if m != nil {
		return m.OutCipher
	}
	return ""
}

func (m *SessionDetailInfoItem) GetInMac() string {
	if m != nil {
		return m.InMac
	}
	return ""
}

func (m *SessionDetailInfoItem) GetOutMac() string {
	if m != nil {
		return m.OutMac
	}
	return ""
}

type SessionDetailInfoEntry struct {
	SessionDetailInfo    []*SessionDetailInfoItem `protobuf:"bytes,8,rep,name=session_detail_info,json=sessionDetailInfo,proto3" json:"session_detail_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SessionDetailInfoEntry) Reset()         { *m = SessionDetailInfoEntry{} }
func (m *SessionDetailInfoEntry) String() string { return proto.CompactTextString(m) }
func (*SessionDetailInfoEntry) ProtoMessage()    {}
func (*SessionDetailInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2992548dd79a11f1, []int{2}
}

func (m *SessionDetailInfoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionDetailInfoEntry.Unmarshal(m, b)
}
func (m *SessionDetailInfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionDetailInfoEntry.Marshal(b, m, deterministic)
}
func (m *SessionDetailInfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionDetailInfoEntry.Merge(m, src)
}
func (m *SessionDetailInfoEntry) XXX_Size() int {
	return xxx_messageInfo_SessionDetailInfoEntry.Size(m)
}
func (m *SessionDetailInfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionDetailInfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SessionDetailInfoEntry proto.InternalMessageInfo

func (m *SessionDetailInfoEntry) GetSessionDetailInfo() []*SessionDetailInfoItem {
	if m != nil {
		return m.SessionDetailInfo
	}
	return nil
}

type SshSessionDetail struct {
	IncomingSessions     *SessionDetailInfoEntry `protobuf:"bytes,50,opt,name=incoming_sessions,json=incomingSessions,proto3" json:"incoming_sessions,omitempty"`
	OutgoingConnections  *SessionDetailInfoEntry `protobuf:"bytes,51,opt,name=outgoing_connections,json=outgoingConnections,proto3" json:"outgoing_connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SshSessionDetail) Reset()         { *m = SshSessionDetail{} }
func (m *SshSessionDetail) String() string { return proto.CompactTextString(m) }
func (*SshSessionDetail) ProtoMessage()    {}
func (*SshSessionDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2992548dd79a11f1, []int{3}
}

func (m *SshSessionDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshSessionDetail.Unmarshal(m, b)
}
func (m *SshSessionDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshSessionDetail.Marshal(b, m, deterministic)
}
func (m *SshSessionDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshSessionDetail.Merge(m, src)
}
func (m *SshSessionDetail) XXX_Size() int {
	return xxx_messageInfo_SshSessionDetail.Size(m)
}
func (m *SshSessionDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_SshSessionDetail.DiscardUnknown(m)
}

var xxx_messageInfo_SshSessionDetail proto.InternalMessageInfo

func (m *SshSessionDetail) GetIncomingSessions() *SessionDetailInfoEntry {
	if m != nil {
		return m.IncomingSessions
	}
	return nil
}

func (m *SshSessionDetail) GetOutgoingConnections() *SessionDetailInfoEntry {
	if m != nil {
		return m.OutgoingConnections
	}
	return nil
}

func init() {
	proto.RegisterType((*SshSessionDetail_KEYS)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh1.kex.nodes.node.ssh_session_detail_KEYS")
	proto.RegisterType((*SessionDetailInfoItem)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh1.kex.nodes.node.session_detail_info_item")
	proto.RegisterType((*SessionDetailInfoEntry)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh1.kex.nodes.node.session_detail_info_entry")
	proto.RegisterType((*SshSessionDetail)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh1.kex.nodes.node.ssh_session_detail")
}

func init() { proto.RegisterFile("ssh_session_detail.proto", fileDescriptor_2992548dd79a11f1) }

var fileDescriptor_2992548dd79a11f1 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0xc9, 0xbd, 0xde, 0xdc, 0x66, 0xaa, 0x60, 0xa7, 0x4a, 0x47, 0x44, 0xa8, 0x59, 0x75,
	0x15, 0xb4, 0x05, 0x5f, 0xa0, 0x76, 0x51, 0x4b, 0x5d, 0xa4, 0x2b, 0x57, 0x43, 0x3a, 0x39, 0x6d,
	0x86, 0x34, 0x73, 0x42, 0x66, 0x02, 0x89, 0x0b, 0x1f, 0xc1, 0xad, 0x8f, 0xe8, 0x6b, 0xc8, 0x4c,
	0x13, 0x04, 0x5b, 0x17, 0x17, 0xba, 0x09, 0xe4, 0xff, 0xce, 0x7f, 0xfe, 0x73, 0x72, 0x08, 0x61,
	0x5a, 0x67, 0x5c, 0x83, 0xd6, 0x12, 0x15, 0x4f, 0xc1, 0x24, 0xf2, 0x14, 0x95, 0x15, 0x1a, 0xa4,
	0x1f, 0x84, 0xd4, 0x02, 0xb9, 0x44, 0xcd, 0x9b, 0x8a, 0x8b, 0xaa, 0x2d, 0x0d, 0x72, 0x5b, 0x8d,
	0x25, 0x54, 0x91, 0xd6, 0xd9, 0xc7, 0x28, 0x87, 0x26, 0x52, 0x98, 0x82, 0x76, 0xcf, 0xf0, 0x13,
	0x99, 0x5c, 0x76, 0xe3, 0x9b, 0xd5, 0xb7, 0x1d, 0x7d, 0x4b, 0x02, 0x5b, 0xc2, 0x55, 0x52, 0x00,
	0xf3, 0xa6, 0xde, 0x2c, 0x88, 0x07, 0x56, 0xf8, 0x9a, 0x14, 0x10, 0xfe, 0xf6, 0x08, 0xfb, 0xc7,
	0x24, 0xd5, 0x01, 0xb9, 0x34, 0x50, 0xd0, 0x77, 0x84, 0xf4, 0x4c, 0xa6, 0xce, 0xfa, 0x22, 0x0e,
	0x3a, 0x65, 0x9d, 0xd2, 0xf7, 0xe4, 0x79, 0x0e, 0x2d, 0x87, 0x46, 0x64, 0x89, 0x3a, 0x02, 0xbb,
	0x73, 0xbd, 0x87, 0x39, 0xb4, 0xab, 0x4e, 0xb2, 0x1d, 0xca, 0x7a, 0x7f, 0x92, 0x82, 0xe7, 0xd0,
	0xb2, 0x7b, 0x57, 0x10, 0x9c, 0x95, 0x0d, 0xb4, 0x76, 0x34, 0xa9, 0xb8, 0x90, 0x65, 0x06, 0x15,
	0x7b, 0x76, 0x1e, 0x4d, 0xaa, 0xa5, 0x7b, 0xb7, 0x5e, 0xac, 0x4d, 0x4f, 0x1f, 0xce, 0x5e, 0xac,
	0x4d, 0x87, 0x5f, 0x13, 0x5f, 0x2a, 0x5e, 0x24, 0x82, 0xf9, 0x0e, 0x3d, 0x48, 0xb5, 0x4d, 0x04,
	0x9d, 0x90, 0x47, 0xeb, 0xb2, 0xfa, 0xa3, 0xd3, 0x7d, 0xac, 0xcd, 0x36, 0x11, 0xe1, 0x2f, 0x8f,
	0xbc, 0xb9, 0xb6, 0x29, 0x28, 0x53, 0xb5, 0xf4, 0x3b, 0x19, 0x5f, 0x81, 0x6c, 0x30, 0xbd, 0x9f,
	0x0d, 0xe7, 0x5f, 0xa2, 0xa7, 0xde, 0x23, 0xfa, 0xdf, 0x37, 0x8d, 0x47, 0x1d, 0xf9, 0xec, 0xc0,
	0x5a, 0x1d, 0x30, 0xfc, 0x79, 0x47, 0xe8, 0xe5, 0xf1, 0x68, 0x43, 0x46, 0x52, 0x09, 0x2c, 0xa4,
	0x3a, 0xf6, 0x48, 0xb3, 0xf9, 0xd4, 0x9b, 0x0d, 0xe7, 0x9b, 0xdb, 0x0c, 0xe4, 0x56, 0x8f, 0x5f,
	0xf6, 0x29, 0xbb, 0x2e, 0x84, 0xfe, 0x20, 0xaf, 0xb0, 0x36, 0x47, 0xb4, 0xc9, 0x02, 0x95, 0x02,
	0x61, 0x5c, 0xf8, 0xe2, 0xf6, 0xe1, 0xe3, 0x3e, 0x68, 0xf9, 0x37, 0x67, 0xef, 0xbb, 0xbf, 0x60,
	0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x36, 0x50, 0x34, 0x90, 0x21, 0x03, 0x00, 0x00,
}