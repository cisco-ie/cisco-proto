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
// source: ssh_session_rekey.proto

package cisco_ios_xr_crypto_ssh_oper_ssh_session_rekey

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

type SshSessionRekey_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshSessionRekey_KEYS) Reset()         { *m = SshSessionRekey_KEYS{} }
func (m *SshSessionRekey_KEYS) String() string { return proto.CompactTextString(m) }
func (*SshSessionRekey_KEYS) ProtoMessage()    {}
func (*SshSessionRekey_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19f1c57ea0ab067, []int{0}
}

func (m *SshSessionRekey_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshSessionRekey_KEYS.Unmarshal(m, b)
}
func (m *SshSessionRekey_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshSessionRekey_KEYS.Marshal(b, m, deterministic)
}
func (m *SshSessionRekey_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshSessionRekey_KEYS.Merge(m, src)
}
func (m *SshSessionRekey_KEYS) XXX_Size() int {
	return xxx_messageInfo_SshSessionRekey_KEYS.Size(m)
}
func (m *SshSessionRekey_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SshSessionRekey_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SshSessionRekey_KEYS proto.InternalMessageInfo

type SessionRekeyInfoItem struct {
	SessionId            uint32   `protobuf:"varint,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionRekeyCount    uint32   `protobuf:"varint,2,opt,name=session_rekey_count,json=sessionRekeyCount,proto3" json:"session_rekey_count,omitempty"`
	TimeToRekey          string   `protobuf:"bytes,3,opt,name=time_to_rekey,json=timeToRekey,proto3" json:"time_to_rekey,omitempty"`
	VolumeToRekey        string   `protobuf:"bytes,4,opt,name=volume_to_rekey,json=volumeToRekey,proto3" json:"volume_to_rekey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionRekeyInfoItem) Reset()         { *m = SessionRekeyInfoItem{} }
func (m *SessionRekeyInfoItem) String() string { return proto.CompactTextString(m) }
func (*SessionRekeyInfoItem) ProtoMessage()    {}
func (*SessionRekeyInfoItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19f1c57ea0ab067, []int{1}
}

func (m *SessionRekeyInfoItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRekeyInfoItem.Unmarshal(m, b)
}
func (m *SessionRekeyInfoItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRekeyInfoItem.Marshal(b, m, deterministic)
}
func (m *SessionRekeyInfoItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRekeyInfoItem.Merge(m, src)
}
func (m *SessionRekeyInfoItem) XXX_Size() int {
	return xxx_messageInfo_SessionRekeyInfoItem.Size(m)
}
func (m *SessionRekeyInfoItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRekeyInfoItem.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRekeyInfoItem proto.InternalMessageInfo

func (m *SessionRekeyInfoItem) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SessionRekeyInfoItem) GetSessionRekeyCount() uint32 {
	if m != nil {
		return m.SessionRekeyCount
	}
	return 0
}

func (m *SessionRekeyInfoItem) GetTimeToRekey() string {
	if m != nil {
		return m.TimeToRekey
	}
	return ""
}

func (m *SessionRekeyInfoItem) GetVolumeToRekey() string {
	if m != nil {
		return m.VolumeToRekey
	}
	return ""
}

type SessionRekeyInfoEntry struct {
	SessionRekeyInfo     []*SessionRekeyInfoItem `protobuf:"bytes,1,rep,name=session_rekey_info,json=sessionRekeyInfo,proto3" json:"session_rekey_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SessionRekeyInfoEntry) Reset()         { *m = SessionRekeyInfoEntry{} }
func (m *SessionRekeyInfoEntry) String() string { return proto.CompactTextString(m) }
func (*SessionRekeyInfoEntry) ProtoMessage()    {}
func (*SessionRekeyInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19f1c57ea0ab067, []int{2}
}

func (m *SessionRekeyInfoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionRekeyInfoEntry.Unmarshal(m, b)
}
func (m *SessionRekeyInfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionRekeyInfoEntry.Marshal(b, m, deterministic)
}
func (m *SessionRekeyInfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionRekeyInfoEntry.Merge(m, src)
}
func (m *SessionRekeyInfoEntry) XXX_Size() int {
	return xxx_messageInfo_SessionRekeyInfoEntry.Size(m)
}
func (m *SessionRekeyInfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionRekeyInfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SessionRekeyInfoEntry proto.InternalMessageInfo

func (m *SessionRekeyInfoEntry) GetSessionRekeyInfo() []*SessionRekeyInfoItem {
	if m != nil {
		return m.SessionRekeyInfo
	}
	return nil
}

type SshSessionRekey struct {
	IncomingSessions     *SessionRekeyInfoEntry `protobuf:"bytes,50,opt,name=incoming_sessions,json=incomingSessions,proto3" json:"incoming_sessions,omitempty"`
	OutgoingConnections  *SessionRekeyInfoEntry `protobuf:"bytes,51,opt,name=outgoing_connections,json=outgoingConnections,proto3" json:"outgoing_connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SshSessionRekey) Reset()         { *m = SshSessionRekey{} }
func (m *SshSessionRekey) String() string { return proto.CompactTextString(m) }
func (*SshSessionRekey) ProtoMessage()    {}
func (*SshSessionRekey) Descriptor() ([]byte, []int) {
	return fileDescriptor_f19f1c57ea0ab067, []int{3}
}

func (m *SshSessionRekey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshSessionRekey.Unmarshal(m, b)
}
func (m *SshSessionRekey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshSessionRekey.Marshal(b, m, deterministic)
}
func (m *SshSessionRekey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshSessionRekey.Merge(m, src)
}
func (m *SshSessionRekey) XXX_Size() int {
	return xxx_messageInfo_SshSessionRekey.Size(m)
}
func (m *SshSessionRekey) XXX_DiscardUnknown() {
	xxx_messageInfo_SshSessionRekey.DiscardUnknown(m)
}

var xxx_messageInfo_SshSessionRekey proto.InternalMessageInfo

func (m *SshSessionRekey) GetIncomingSessions() *SessionRekeyInfoEntry {
	if m != nil {
		return m.IncomingSessions
	}
	return nil
}

func (m *SshSessionRekey) GetOutgoingConnections() *SessionRekeyInfoEntry {
	if m != nil {
		return m.OutgoingConnections
	}
	return nil
}

func init() {
	proto.RegisterType((*SshSessionRekey_KEYS)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.rekey.ssh_session_rekey_KEYS")
	proto.RegisterType((*SessionRekeyInfoItem)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.rekey.session_rekey_info_item")
	proto.RegisterType((*SessionRekeyInfoEntry)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.rekey.session_rekey_info_entry")
	proto.RegisterType((*SshSessionRekey)(nil), "cisco_ios_xr_crypto_ssh_oper.ssh.session.rekey.ssh_session_rekey")
}

func init() { proto.RegisterFile("ssh_session_rekey.proto", fileDescriptor_f19f1c57ea0ab067) }

var fileDescriptor_f19f1c57ea0ab067 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x99, 0xf6, 0x72, 0xa1, 0xa7, 0x14, 0xdb, 0xa9, 0xd8, 0xd9, 0x08, 0x21, 0x0b, 0xc9,
	0x6a, 0x16, 0xe9, 0x23, 0x14, 0xd1, 0xe2, 0x2e, 0x75, 0xe3, 0x6a, 0xc0, 0xe9, 0xb4, 0x1d, 0x34,
	0x73, 0xc2, 0xcc, 0x44, 0x0c, 0x3e, 0x80, 0x6b, 0xdf, 0xc5, 0x07, 0x94, 0x4c, 0x13, 0x6c, 0x89,
	0x2e, 0x84, 0x6e, 0xff, 0xff, 0x3b, 0xf9, 0xe6, 0x1c, 0x02, 0x33, 0xe7, 0x76, 0xc2, 0x29, 0xe7,
	0x34, 0x1a, 0x61, 0xd5, 0x93, 0xaa, 0x78, 0x61, 0xd1, 0x23, 0xe5, 0x52, 0x3b, 0x89, 0x42, 0xa3,
	0x13, 0xaf, 0x56, 0x48, 0x5b, 0x15, 0x1e, 0x45, 0x0d, 0x63, 0xa1, 0x2c, 0x77, 0x6e, 0xc7, 0x9b,
	0x29, 0x1e, 0xa6, 0x62, 0x06, 0x17, 0x9d, 0x4f, 0x89, 0xbb, 0xeb, 0x87, 0x55, 0xfc, 0x49, 0x60,
	0x76, 0x1c, 0x6b, 0xb3, 0x41, 0xa1, 0xbd, 0xca, 0xe9, 0x25, 0x40, 0x5b, 0xe9, 0x35, 0x23, 0x11,
	0x49, 0x46, 0xd9, 0xa0, 0x49, 0x96, 0x6b, 0xca, 0x61, 0x7a, 0x3c, 0x29, 0xb1, 0x34, 0x9e, 0xf5,
	0x02, 0x37, 0x69, 0xaa, 0xac, 0x6e, 0x16, 0x75, 0x41, 0x63, 0x18, 0x79, 0x9d, 0x2b, 0xe1, 0x71,
	0xcf, 0xb3, 0x7e, 0x44, 0x92, 0x41, 0x36, 0xac, 0xc3, 0x7b, 0x0c, 0x20, 0xbd, 0x82, 0xb3, 0x17,
	0x7c, 0x2e, 0x0f, 0xa9, 0x7f, 0x81, 0x1a, 0xed, 0xe3, 0x86, 0x8b, 0x3f, 0x08, 0xb0, 0x1f, 0x9e,
	0xad, 0x8c, 0xb7, 0x15, 0x2d, 0x81, 0x76, 0x3b, 0x46, 0xa2, 0x7e, 0x32, 0x4c, 0x6f, 0xfe, 0x78,
	0x3a, 0xfe, 0xcb, 0x71, 0xb2, 0xf1, 0xe1, 0x82, 0x4b, 0xb3, 0xc1, 0xf8, 0xbd, 0x07, 0x93, 0xce,
	0x95, 0x69, 0x09, 0x13, 0x6d, 0x24, 0xe6, 0xda, 0x6c, 0xdb, 0xc6, 0xb1, 0x34, 0x22, 0xc9, 0x30,
	0xbd, 0x3d, 0xc1, 0x5b, 0xc2, 0xc6, 0xd9, 0xb8, 0x55, 0xac, 0x1a, 0x03, 0x7d, 0x83, 0x73, 0x2c,
	0xfd, 0x16, 0x6b, 0xad, 0x44, 0x63, 0x94, 0xf4, 0xc1, 0x3c, 0x3f, 0xb1, 0x79, 0xda, 0x5a, 0x16,
	0xdf, 0x92, 0xc7, 0xff, 0xe1, 0x2f, 0x9d, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x82, 0xf9, 0xbe,
	0xa4, 0xc0, 0x02, 0x00, 0x00,
}