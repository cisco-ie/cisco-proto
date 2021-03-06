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

package cisco_ios_xr_crypto_ssh_oper_ssh_session_detail

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

type SessionDetailInfoItem struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	KeyExchange          string   `protobuf:"bytes,2,opt,name=key_exchange,json=keyExchange,proto3" json:"key_exchange,omitempty"`
	PublicKey            string   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InCipher             string   `protobuf:"bytes,4,opt,name=in_cipher,json=inCipher,proto3" json:"in_cipher,omitempty"`
	OutCipher            string   `protobuf:"bytes,5,opt,name=out_cipher,json=outCipher,proto3" json:"out_cipher,omitempty"`
	InMac                string   `protobuf:"bytes,6,opt,name=in_mac,json=inMac,proto3" json:"in_mac,omitempty"`
	OutMac               string   `protobuf:"bytes,7,opt,name=out_mac,json=outMac,proto3" json:"out_mac,omitempty"`
	StartTime            string   `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              string   `protobuf:"bytes,9,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
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

func (m *SessionDetailInfoItem) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *SessionDetailInfoItem) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

type SessionDetailInfoEntry struct {
	SessionDetailInfo    []*SessionDetailInfoItem `protobuf:"bytes,1,rep,name=session_detail_info,json=sessionDetailInfo,proto3" json:"session_detail_info,omitempty"`
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
	proto.RegisterType((*SshSessionDetail_KEYS)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.detail.ssh_session_detail_KEYS")
	proto.RegisterType((*SessionDetailInfoItem)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.detail.session_detail_info_item")
	proto.RegisterType((*SessionDetailInfoEntry)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.detail.session_detail_info_entry")
	proto.RegisterType((*SshSessionDetail)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.detail.ssh_session_detail")
}

func init() { proto.RegisterFile("ssh_session_detail.proto", fileDescriptor_2992548dd79a11f1) }

var fileDescriptor_2992548dd79a11f1 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x2d, 0xcd, 0x6e, 0xbc, 0x20, 0x51, 0x17, 0x54, 0xaf, 0x10, 0x52, 0xc8, 0x29,
	0xa7, 0x20, 0x6d, 0x1f, 0xa1, 0xf4, 0xb0, 0x54, 0x5c, 0x52, 0x2e, 0x9c, 0xac, 0xd4, 0x99, 0x6e,
	0x46, 0xdb, 0x78, 0xa2, 0xd8, 0x11, 0x9b, 0x03, 0x6f, 0xc0, 0x99, 0x3b, 0x6f, 0x8a, 0xec, 0x24,
	0x42, 0x62, 0x97, 0x43, 0xa5, 0x3d, 0xfa, 0xff, 0xe6, 0x9f, 0xdf, 0xe3, 0x91, 0x99, 0x30, 0xa6,
	0x92, 0x06, 0x8c, 0x41, 0xd2, 0xb2, 0x04, 0x5b, 0xe0, 0x53, 0xd6, 0xb4, 0x64, 0x89, 0x7f, 0x54,
	0x68, 0x14, 0x49, 0x24, 0x23, 0xf7, 0xad, 0x54, 0x6d, 0xdf, 0x58, 0x92, 0xae, 0x9a, 0x1a, 0x68,
	0x33, 0x63, 0xaa, 0x6c, 0xb4, 0x65, 0x83, 0x2d, 0x59, 0xb1, 0xab, 0xc3, 0x66, 0xf2, 0xee, 0xf6,
	0xdb, 0x7d, 0xf2, 0x7b, 0xc6, 0xc4, 0x3f, 0x3a, 0xea, 0x47, 0x92, 0x68, 0xa1, 0xe6, 0xef, 0x19,
	0x9b, 0x18, 0x96, 0x22, 0x88, 0x83, 0xf4, 0x55, 0x1e, 0x8d, 0xca, 0xa6, 0xe4, 0x1f, 0xd8, 0xcb,
	0x1d, 0xf4, 0x12, 0xf6, 0xaa, 0x2a, 0xf4, 0x16, 0xc4, 0x2c, 0x0e, 0xd2, 0x28, 0x5f, 0xee, 0xa0,
	0xbf, 0x1d, 0x25, 0xd7, 0xa1, 0xe9, 0x1e, 0x9e, 0x50, 0xc9, 0x1d, 0xf4, 0xe2, 0xcc, 0x17, 0x44,
	0x83, 0x72, 0x07, 0x3d, 0x7f, 0xc7, 0x22, 0xd4, 0x52, 0x61, 0x53, 0x41, 0x2b, 0x5e, 0x78, 0xba,
	0x40, 0x7d, 0xe3, 0xcf, 0xce, 0x4b, 0x9d, 0x9d, 0xe8, 0xf9, 0xe0, 0xa5, 0xce, 0x8e, 0xf8, 0x2d,
	0x0b, 0x51, 0xcb, 0xba, 0x50, 0x22, 0xf4, 0xe8, 0x1c, 0xf5, 0x97, 0x42, 0xf1, 0x2b, 0x36, 0x77,
	0x2e, 0xa7, 0xcf, 0xbd, 0x1e, 0x52, 0x67, 0x1d, 0x70, 0xc3, 0xd8, 0xa2, 0xb5, 0xd2, 0x62, 0x0d,
	0x62, 0x31, 0xb4, 0xf3, 0xca, 0x57, 0xac, 0x81, 0xaf, 0xd8, 0x02, 0x74, 0x39, 0xc0, 0xc8, 0xc3,
	0x39, 0xe8, 0xd2, 0xa1, 0xe4, 0x57, 0xc0, 0x56, 0xc7, 0xde, 0x08, 0xb4, 0x6d, 0x7b, 0xde, 0xb3,
	0xcb, 0x23, 0x50, 0x04, 0xf1, 0x59, 0xba, 0x5c, 0x6f, 0xb2, 0x67, 0xee, 0x2a, 0xfb, 0xdf, 0x32,
	0xf2, 0x8b, 0x91, 0x7c, 0xf2, 0x60, 0xa3, 0x1f, 0x29, 0xf9, 0x39, 0x63, 0xfc, 0x70, 0xb1, 0xfc,
	0x3b, 0xbb, 0x40, 0xad, 0xa8, 0x46, 0xbd, 0x9d, 0x90, 0x11, 0xeb, 0x38, 0x48, 0x97, 0xeb, 0xcf,
	0x27, 0xb9, 0x8f, 0x1f, 0x3c, 0x7f, 0x3d, 0x85, 0xdc, 0x8f, 0x19, 0xfc, 0x07, 0x7b, 0x43, 0x9d,
	0xdd, 0x92, 0x0b, 0x56, 0xa4, 0x35, 0x28, 0xeb, 0xb3, 0xaf, 0x4f, 0x9e, 0x7d, 0x39, 0xe5, 0xdc,
	0xfc, 0x8d, 0x79, 0x08, 0xfd, 0xf7, 0xb8, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x54, 0xfb, 0xd3,
	0x17, 0x3a, 0x03, 0x00, 0x00,
}
