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
// source: install_log_entry_4.proto

package cisco_ios_xr_installmgr_admin_oper_install_logs_log

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

type InstallLogEntry_4_KEYS struct {
	RequestId            uint32   `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallLogEntry_4_KEYS) Reset()         { *m = InstallLogEntry_4_KEYS{} }
func (m *InstallLogEntry_4_KEYS) String() string { return proto.CompactTextString(m) }
func (*InstallLogEntry_4_KEYS) ProtoMessage()    {}
func (*InstallLogEntry_4_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{0}
}

func (m *InstallLogEntry_4_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallLogEntry_4_KEYS.Unmarshal(m, b)
}
func (m *InstallLogEntry_4_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallLogEntry_4_KEYS.Marshal(b, m, deterministic)
}
func (m *InstallLogEntry_4_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallLogEntry_4_KEYS.Merge(m, src)
}
func (m *InstallLogEntry_4_KEYS) XXX_Size() int {
	return xxx_messageInfo_InstallLogEntry_4_KEYS.Size(m)
}
func (m *InstallLogEntry_4_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallLogEntry_4_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_InstallLogEntry_4_KEYS proto.InternalMessageInfo

func (m *InstallLogEntry_4_KEYS) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type InstmgrBagLogEntryUserMsgScope_ struct {
	AdminRead            bool     `protobuf:"varint,1,opt,name=admin_read,json=adminRead,proto3" json:"admin_read,omitempty"`
	AffectedSdRs         uint32   `protobuf:"varint,2,opt,name=affected_sd_rs,json=affectedSdRs,proto3" json:"affected_sd_rs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstmgrBagLogEntryUserMsgScope_) Reset()         { *m = InstmgrBagLogEntryUserMsgScope_{} }
func (m *InstmgrBagLogEntryUserMsgScope_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagLogEntryUserMsgScope_) ProtoMessage()    {}
func (*InstmgrBagLogEntryUserMsgScope_) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{1}
}

func (m *InstmgrBagLogEntryUserMsgScope_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_.Unmarshal(m, b)
}
func (m *InstmgrBagLogEntryUserMsgScope_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagLogEntryUserMsgScope_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_.Merge(m, src)
}
func (m *InstmgrBagLogEntryUserMsgScope_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_.Size(m)
}
func (m *InstmgrBagLogEntryUserMsgScope_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagLogEntryUserMsgScope_ proto.InternalMessageInfo

func (m *InstmgrBagLogEntryUserMsgScope_) GetAdminRead() bool {
	if m != nil {
		return m.AdminRead
	}
	return false
}

func (m *InstmgrBagLogEntryUserMsgScope_) GetAffectedSdRs() uint32 {
	if m != nil {
		return m.AffectedSdRs
	}
	return 0
}

type InstmgrBagLogEntryV3_ struct {
	Category             string                           `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Scope                *InstmgrBagLogEntryUserMsgScope_ `protobuf:"bytes,2,opt,name=scope,proto3" json:"scope,omitempty"`
	Message              string                           `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *InstmgrBagLogEntryV3_) Reset()         { *m = InstmgrBagLogEntryV3_{} }
func (m *InstmgrBagLogEntryV3_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagLogEntryV3_) ProtoMessage()    {}
func (*InstmgrBagLogEntryV3_) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{2}
}

func (m *InstmgrBagLogEntryV3_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagLogEntryV3_.Unmarshal(m, b)
}
func (m *InstmgrBagLogEntryV3_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagLogEntryV3_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagLogEntryV3_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagLogEntryV3_.Merge(m, src)
}
func (m *InstmgrBagLogEntryV3_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagLogEntryV3_.Size(m)
}
func (m *InstmgrBagLogEntryV3_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagLogEntryV3_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagLogEntryV3_ proto.InternalMessageInfo

func (m *InstmgrBagLogEntryV3_) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *InstmgrBagLogEntryV3_) GetScope() *InstmgrBagLogEntryUserMsgScope_ {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *InstmgrBagLogEntryV3_) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type MsgContents struct {
	Version              uint32                 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	V3                   *InstmgrBagLogEntryV3_ `protobuf:"bytes,2,opt,name=v3,proto3" json:"v3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MsgContents) Reset()         { *m = MsgContents{} }
func (m *MsgContents) String() string { return proto.CompactTextString(m) }
func (*MsgContents) ProtoMessage()    {}
func (*MsgContents) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{3}
}

func (m *MsgContents) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgContents.Unmarshal(m, b)
}
func (m *MsgContents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgContents.Marshal(b, m, deterministic)
}
func (m *MsgContents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgContents.Merge(m, src)
}
func (m *MsgContents) XXX_Size() int {
	return xxx_messageInfo_MsgContents.Size(m)
}
func (m *MsgContents) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgContents.DiscardUnknown(m)
}

var xxx_messageInfo_MsgContents proto.InternalMessageInfo

func (m *MsgContents) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MsgContents) GetV3() *InstmgrBagLogEntryV3_ {
	if m != nil {
		return m.V3
	}
	return nil
}

type InstmgrBagLogEntryUserMsg_3_ struct {
	LogContents          *MsgContents `protobuf:"bytes,1,opt,name=log_contents,json=logContents,proto3" json:"log_contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InstmgrBagLogEntryUserMsg_3_) Reset()         { *m = InstmgrBagLogEntryUserMsg_3_{} }
func (m *InstmgrBagLogEntryUserMsg_3_) String() string { return proto.CompactTextString(m) }
func (*InstmgrBagLogEntryUserMsg_3_) ProtoMessage()    {}
func (*InstmgrBagLogEntryUserMsg_3_) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{4}
}

func (m *InstmgrBagLogEntryUserMsg_3_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_.Unmarshal(m, b)
}
func (m *InstmgrBagLogEntryUserMsg_3_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_.Marshal(b, m, deterministic)
}
func (m *InstmgrBagLogEntryUserMsg_3_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_.Merge(m, src)
}
func (m *InstmgrBagLogEntryUserMsg_3_) XXX_Size() int {
	return xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_.Size(m)
}
func (m *InstmgrBagLogEntryUserMsg_3_) XXX_DiscardUnknown() {
	xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_.DiscardUnknown(m)
}

var xxx_messageInfo_InstmgrBagLogEntryUserMsg_3_ proto.InternalMessageInfo

func (m *InstmgrBagLogEntryUserMsg_3_) GetLogContents() *MsgContents {
	if m != nil {
		return m.LogContents
	}
	return nil
}

type InstallLogEntry_4 struct {
	Header               []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,50,rep,name=header,proto3" json:"header,omitempty"`
	Summary              []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,51,rep,name=summary,proto3" json:"summary,omitempty"`
	Message              []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,52,rep,name=message,proto3" json:"message,omitempty"`
	Change               []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,53,rep,name=change,proto3" json:"change,omitempty"`
	Detail               []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,54,rep,name=detail,proto3" json:"detail,omitempty"`
	Communication        []*InstmgrBagLogEntryUserMsg_3_ `protobuf:"bytes,55,rep,name=communication,proto3" json:"communication,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InstallLogEntry_4) Reset()         { *m = InstallLogEntry_4{} }
func (m *InstallLogEntry_4) String() string { return proto.CompactTextString(m) }
func (*InstallLogEntry_4) ProtoMessage()    {}
func (*InstallLogEntry_4) Descriptor() ([]byte, []int) {
	return fileDescriptor_a31645df27a2433b, []int{5}
}

func (m *InstallLogEntry_4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallLogEntry_4.Unmarshal(m, b)
}
func (m *InstallLogEntry_4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallLogEntry_4.Marshal(b, m, deterministic)
}
func (m *InstallLogEntry_4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallLogEntry_4.Merge(m, src)
}
func (m *InstallLogEntry_4) XXX_Size() int {
	return xxx_messageInfo_InstallLogEntry_4.Size(m)
}
func (m *InstallLogEntry_4) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallLogEntry_4.DiscardUnknown(m)
}

var xxx_messageInfo_InstallLogEntry_4 proto.InternalMessageInfo

func (m *InstallLogEntry_4) GetHeader() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *InstallLogEntry_4) GetSummary() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Summary
	}
	return nil
}

func (m *InstallLogEntry_4) GetMessage() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *InstallLogEntry_4) GetChange() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Change
	}
	return nil
}

func (m *InstallLogEntry_4) GetDetail() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *InstallLogEntry_4) GetCommunication() []*InstmgrBagLogEntryUserMsg_3_ {
	if m != nil {
		return m.Communication
	}
	return nil
}

func init() {
	proto.RegisterType((*InstallLogEntry_4_KEYS)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.install_log_entry_4_KEYS")
	proto.RegisterType((*InstmgrBagLogEntryUserMsgScope_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.instmgr_bag_log_entry_user_msg_scope_")
	proto.RegisterType((*InstmgrBagLogEntryV3_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.instmgr_bag_log_entry_v3_")
	proto.RegisterType((*MsgContents)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.msg_contents")
	proto.RegisterType((*InstmgrBagLogEntryUserMsg_3_)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.instmgr_bag_log_entry_user_msg_3_")
	proto.RegisterType((*InstallLogEntry_4)(nil), "cisco_ios_xr_installmgr_admin_oper.install.logs.log.install_log_entry_4")
}

func init() { proto.RegisterFile("install_log_entry_4.proto", fileDescriptor_a31645df27a2433b) }

var fileDescriptor_a31645df27a2433b = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x8b, 0xd5, 0x30,
	0x10, 0xc7, 0xe9, 0x5b, 0xf6, 0xc7, 0x9b, 0xf7, 0xd6, 0x43, 0xbc, 0x54, 0x41, 0x58, 0x8b, 0xc2,
	0x9e, 0x7a, 0x78, 0x59, 0x15, 0x8f, 0x22, 0x1e, 0x44, 0xf0, 0x90, 0x05, 0x41, 0x0f, 0x0e, 0xd9,
	0x64, 0x36, 0x5b, 0x68, 0x93, 0x9a, 0xe4, 0x15, 0x1f, 0xf8, 0x07, 0xe8, 0xdf, 0xe5, 0xc1, 0x7f,
	0x4b, 0xd2, 0x1f, 0xb2, 0xe2, 0x82, 0xb2, 0xd8, 0x4b, 0x61, 0xa6, 0x93, 0xf9, 0x7c, 0xe7, 0xdb,
	0x69, 0xe0, 0x5e, 0x65, 0x43, 0x94, 0x75, 0x8d, 0xb5, 0x33, 0x48, 0x36, 0xfa, 0x1d, 0x9e, 0x95,
	0xad, 0x77, 0xd1, 0x31, 0xae, 0xaa, 0xa0, 0x1c, 0x56, 0x2e, 0xe0, 0x67, 0x8f, 0x63, 0x5d, 0x63,
	0x3c, 0x4a, 0xdd, 0x54, 0x16, 0x5d, 0x4b, 0xbe, 0x1c, 0xb3, 0x65, 0xed, 0x4c, 0x48, 0x8f, 0xe2,
	0x39, 0xe4, 0x37, 0x74, 0xc4, 0x37, 0xaf, 0xde, 0x9f, 0xb3, 0x07, 0x00, 0x9e, 0x3e, 0x6d, 0x29,
	0x44, 0xac, 0x74, 0x9e, 0x9d, 0x64, 0xa7, 0xc7, 0x62, 0x39, 0x66, 0x5e, 0xeb, 0xa2, 0x86, 0xc7,
	0xe9, 0x68, 0x22, 0x5c, 0x48, 0x73, 0xed, 0xf8, 0x36, 0x90, 0xc7, 0x26, 0x18, 0x0c, 0xca, 0xb5,
	0x84, 0xa9, 0xcf, 0x20, 0xc1, 0x93, 0x1c, 0xfa, 0x1c, 0x89, 0x65, 0x9f, 0x11, 0x24, 0x35, 0x7b,
	0x04, 0x77, 0xe4, 0xe5, 0x25, 0xa9, 0x48, 0x1a, 0x83, 0x46, 0x1f, 0xf2, 0x45, 0x8f, 0x5a, 0x4f,
	0xd9, 0x73, 0x2d, 0x42, 0xf1, 0x3d, 0x1b, 0x66, 0xff, 0x13, 0xd7, 0x71, 0x64, 0xf7, 0xe1, 0x48,
	0xc9, 0x48, 0xc6, 0xf9, 0x5d, 0x0f, 0x58, 0x8a, 0x5f, 0x31, 0x6b, 0x61, 0xbf, 0x17, 0xd2, 0xb7,
	0x5d, 0x6d, 0x3e, 0x94, 0xb7, 0xf0, 0xa9, 0xfc, 0xa7, 0x49, 0xc5, 0x00, 0x62, 0x39, 0x1c, 0x36,
	0x14, 0x82, 0x34, 0x94, 0xef, 0xf5, 0x62, 0xa6, 0xb0, 0xf8, 0x9a, 0xc1, 0x3a, 0xd5, 0x2b, 0x67,
	0x23, 0xd9, 0x18, 0x52, 0x69, 0x47, 0x3e, 0x54, 0xce, 0x8e, 0x06, 0x4f, 0x21, 0xfb, 0x08, 0x8b,
	0x8e, 0x8f, 0x9a, 0xdf, 0xfe, 0x47, 0xcd, 0x1d, 0x47, 0xb1, 0xe8, 0x78, 0xf1, 0x2d, 0x83, 0x87,
	0x7f, 0x99, 0x8a, 0x23, 0xd3, 0xb0, 0x4e, 0x2f, 0x26, 0xbd, 0xbd, 0xc8, 0xd5, 0xe6, 0xc5, 0xad,
	0xf4, 0x5c, 0x1f, 0x5c, 0xac, 0x6a, 0x67, 0x5e, 0x8e, 0x41, 0xf1, 0x63, 0x1f, 0xee, 0xde, 0xb0,
	0x86, 0xcc, 0xc2, 0xc1, 0x15, 0x49, 0x4d, 0x3e, 0xdf, 0x9c, 0xec, 0x9d, 0xae, 0x36, 0xef, 0xe6,
	0xf8, 0x76, 0x1c, 0xc5, 0x48, 0x61, 0x2d, 0x1c, 0x86, 0x6d, 0xd3, 0x48, 0xbf, 0xcb, 0xf9, 0xac,
	0xc0, 0x09, 0x93, 0x88, 0xd3, 0xaa, 0x9c, 0xcd, 0x4b, 0x1c, 0x31, 0xc9, 0x53, 0x75, 0x25, 0xad,
	0xa1, 0xfc, 0xc9, 0xbc, 0x9e, 0x0e, 0x94, 0xc4, 0xd3, 0x14, 0x65, 0x55, 0xe7, 0x4f, 0xe7, 0xe5,
	0x0d, 0x14, 0xf6, 0x05, 0x8e, 0x95, 0x6b, 0x9a, 0xad, 0xad, 0x94, 0x8c, 0xe9, 0xbf, 0x7a, 0x36,
	0x2b, 0xf6, 0x77, 0xd8, 0xc5, 0x41, 0x7f, 0x17, 0xf3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18,
	0xbb, 0xeb, 0xf1, 0xa8, 0x05, 0x00, 0x00,
}
