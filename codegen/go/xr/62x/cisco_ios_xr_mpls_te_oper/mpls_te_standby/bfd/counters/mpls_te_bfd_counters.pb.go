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
// source: mpls_te_bfd_counters.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_bfd_counters

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

type MplsTeBfdCounters_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBfdCounters_KEYS) Reset()         { *m = MplsTeBfdCounters_KEYS{} }
func (m *MplsTeBfdCounters_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdCounters_KEYS) ProtoMessage()    {}
func (*MplsTeBfdCounters_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b51aa5f051869a0b, []int{0}
}

func (m *MplsTeBfdCounters_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Unmarshal(m, b)
}
func (m *MplsTeBfdCounters_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdCounters_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdCounters_KEYS.Merge(m, src)
}
func (m *MplsTeBfdCounters_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdCounters_KEYS.Size(m)
}
func (m *MplsTeBfdCounters_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdCounters_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdCounters_KEYS proto.InternalMessageInfo

type MplsTeBfdLspCounters struct {
	SessionCreateEvents              uint32   `protobuf:"varint,1,opt,name=session_create_events,json=sessionCreateEvents,proto3" json:"session_create_events,omitempty"`
	SessionUpEvents                  uint32   `protobuf:"varint,2,opt,name=session_up_events,json=sessionUpEvents,proto3" json:"session_up_events,omitempty"`
	SessionCreationFailedEvents      uint32   `protobuf:"varint,3,opt,name=session_creation_failed_events,json=sessionCreationFailedEvents,proto3" json:"session_creation_failed_events,omitempty"`
	SessionDownEvents                uint32   `protobuf:"varint,4,opt,name=session_down_events,json=sessionDownEvents,proto3" json:"session_down_events,omitempty"`
	SessionAdminDownEvents           uint32   `protobuf:"varint,5,opt,name=session_admin_down_events,json=sessionAdminDownEvents,proto3" json:"session_admin_down_events,omitempty"`
	SessionGracefullyDeleteEvents    uint32   `protobuf:"varint,6,opt,name=session_gracefully_delete_events,json=sessionGracefullyDeleteEvents,proto3" json:"session_gracefully_delete_events,omitempty"`
	SessionNonGracefullyDeleteEvents uint32   `protobuf:"varint,7,opt,name=session_non_gracefully_delete_events,json=sessionNonGracefullyDeleteEvents,proto3" json:"session_non_gracefully_delete_events,omitempty"`
	SessionCreateTimeoutEvents       uint32   `protobuf:"varint,8,opt,name=session_create_timeout_events,json=sessionCreateTimeoutEvents,proto3" json:"session_create_timeout_events,omitempty"`
	SessionReplayEvents              uint32   `protobuf:"varint,9,opt,name=session_replay_events,json=sessionReplayEvents,proto3" json:"session_replay_events,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *MplsTeBfdLspCounters) Reset()         { *m = MplsTeBfdLspCounters{} }
func (m *MplsTeBfdLspCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdLspCounters) ProtoMessage()    {}
func (*MplsTeBfdLspCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_b51aa5f051869a0b, []int{1}
}

func (m *MplsTeBfdLspCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdLspCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdLspCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdLspCounters.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdLspCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdLspCounters.Merge(m, src)
}
func (m *MplsTeBfdLspCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdLspCounters.Size(m)
}
func (m *MplsTeBfdLspCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdLspCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdLspCounters proto.InternalMessageInfo

func (m *MplsTeBfdLspCounters) GetSessionCreateEvents() uint32 {
	if m != nil {
		return m.SessionCreateEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionUpEvents() uint32 {
	if m != nil {
		return m.SessionUpEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionCreationFailedEvents() uint32 {
	if m != nil {
		return m.SessionCreationFailedEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionDownEvents() uint32 {
	if m != nil {
		return m.SessionDownEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionAdminDownEvents() uint32 {
	if m != nil {
		return m.SessionAdminDownEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionNonGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionNonGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionCreateTimeoutEvents() uint32 {
	if m != nil {
		return m.SessionCreateTimeoutEvents
	}
	return 0
}

func (m *MplsTeBfdLspCounters) GetSessionReplayEvents() uint32 {
	if m != nil {
		return m.SessionReplayEvents
	}
	return 0
}

type MplsTeBfdSessionCounters struct {
	SessionCreateEvents              uint32   `protobuf:"varint,1,opt,name=session_create_events,json=sessionCreateEvents,proto3" json:"session_create_events,omitempty"`
	SessionUpEvents                  uint32   `protobuf:"varint,2,opt,name=session_up_events,json=sessionUpEvents,proto3" json:"session_up_events,omitempty"`
	SessionCreationFailedEvents      uint32   `protobuf:"varint,3,opt,name=session_creation_failed_events,json=sessionCreationFailedEvents,proto3" json:"session_creation_failed_events,omitempty"`
	SessionDownEvents                uint32   `protobuf:"varint,4,opt,name=session_down_events,json=sessionDownEvents,proto3" json:"session_down_events,omitempty"`
	SessionAdminDownEvents           uint32   `protobuf:"varint,5,opt,name=session_admin_down_events,json=sessionAdminDownEvents,proto3" json:"session_admin_down_events,omitempty"`
	SessionGracefullyDeleteEvents    uint32   `protobuf:"varint,6,opt,name=session_gracefully_delete_events,json=sessionGracefullyDeleteEvents,proto3" json:"session_gracefully_delete_events,omitempty"`
	SessionNonGracefullyDeleteEvents uint32   `protobuf:"varint,7,opt,name=session_non_gracefully_delete_events,json=sessionNonGracefullyDeleteEvents,proto3" json:"session_non_gracefully_delete_events,omitempty"`
	SessionReplayEvents              uint32   `protobuf:"varint,8,opt,name=session_replay_events,json=sessionReplayEvents,proto3" json:"session_replay_events,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *MplsTeBfdSessionCounters) Reset()         { *m = MplsTeBfdSessionCounters{} }
func (m *MplsTeBfdSessionCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdSessionCounters) ProtoMessage()    {}
func (*MplsTeBfdSessionCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_b51aa5f051869a0b, []int{2}
}

func (m *MplsTeBfdSessionCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdSessionCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdSessionCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdSessionCounters.Merge(m, src)
}
func (m *MplsTeBfdSessionCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdSessionCounters.Size(m)
}
func (m *MplsTeBfdSessionCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdSessionCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdSessionCounters proto.InternalMessageInfo

func (m *MplsTeBfdSessionCounters) GetSessionCreateEvents() uint32 {
	if m != nil {
		return m.SessionCreateEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionUpEvents() uint32 {
	if m != nil {
		return m.SessionUpEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionCreationFailedEvents() uint32 {
	if m != nil {
		return m.SessionCreationFailedEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionDownEvents() uint32 {
	if m != nil {
		return m.SessionDownEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionAdminDownEvents() uint32 {
	if m != nil {
		return m.SessionAdminDownEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionNonGracefullyDeleteEvents() uint32 {
	if m != nil {
		return m.SessionNonGracefullyDeleteEvents
	}
	return 0
}

func (m *MplsTeBfdSessionCounters) GetSessionReplayEvents() uint32 {
	if m != nil {
		return m.SessionReplayEvents
	}
	return 0
}

type MplsTeBfdCounters struct {
	LastClearedTimestamp    uint32                    `protobuf:"varint,50,opt,name=last_cleared_timestamp,json=lastClearedTimestamp,proto3" json:"last_cleared_timestamp,omitempty"`
	BfdOverLspHeadCounters  *MplsTeBfdLspCounters     `protobuf:"bytes,51,opt,name=bfd_over_lsp_head_counters,json=bfdOverLspHeadCounters,proto3" json:"bfd_over_lsp_head_counters,omitempty"`
	SbfdOverLspHeadCounters *MplsTeBfdLspCounters     `protobuf:"bytes,52,opt,name=sbfd_over_lsp_head_counters,json=sbfdOverLspHeadCounters,proto3" json:"sbfd_over_lsp_head_counters,omitempty"`
	BfdOverLspTailCounters  *MplsTeBfdSessionCounters `protobuf:"bytes,53,opt,name=bfd_over_lsp_tail_counters,json=bfdOverLspTailCounters,proto3" json:"bfd_over_lsp_tail_counters,omitempty"`
	BfDoLmCounters          *MplsTeBfdSessionCounters `protobuf:"bytes,54,opt,name=bf_do_lm_counters,json=bfDoLmCounters,proto3" json:"bf_do_lm_counters,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                  `json:"-"`
	XXX_unrecognized        []byte                    `json:"-"`
	XXX_sizecache           int32                     `json:"-"`
}

func (m *MplsTeBfdCounters) Reset()         { *m = MplsTeBfdCounters{} }
func (m *MplsTeBfdCounters) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdCounters) ProtoMessage()    {}
func (*MplsTeBfdCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_b51aa5f051869a0b, []int{3}
}

func (m *MplsTeBfdCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdCounters.Unmarshal(m, b)
}
func (m *MplsTeBfdCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdCounters.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdCounters.Merge(m, src)
}
func (m *MplsTeBfdCounters) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdCounters.Size(m)
}
func (m *MplsTeBfdCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdCounters proto.InternalMessageInfo

func (m *MplsTeBfdCounters) GetLastClearedTimestamp() uint32 {
	if m != nil {
		return m.LastClearedTimestamp
	}
	return 0
}

func (m *MplsTeBfdCounters) GetBfdOverLspHeadCounters() *MplsTeBfdLspCounters {
	if m != nil {
		return m.BfdOverLspHeadCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetSbfdOverLspHeadCounters() *MplsTeBfdLspCounters {
	if m != nil {
		return m.SbfdOverLspHeadCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetBfdOverLspTailCounters() *MplsTeBfdSessionCounters {
	if m != nil {
		return m.BfdOverLspTailCounters
	}
	return nil
}

func (m *MplsTeBfdCounters) GetBfDoLmCounters() *MplsTeBfdSessionCounters {
	if m != nil {
		return m.BfDoLmCounters
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeBfdCounters_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.counters.mpls_te_bfd_counters_KEYS")
	proto.RegisterType((*MplsTeBfdLspCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.counters.mpls_te_bfd_lsp_counters")
	proto.RegisterType((*MplsTeBfdSessionCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.counters.mpls_te_bfd_session_counters")
	proto.RegisterType((*MplsTeBfdCounters)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.counters.mpls_te_bfd_counters")
}

func init() { proto.RegisterFile("mpls_te_bfd_counters.proto", fileDescriptor_b51aa5f051869a0b) }

var fileDescriptor_b51aa5f051869a0b = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x15, 0xfa, 0x43, 0x31, 0x02, 0xd4, 0xa1, 0x84, 0x34, 0xa1, 0x28, 0x8a, 0x58, 0x54,
	0x2c, 0x66, 0x91, 0x96, 0x4a, 0x2c, 0xab, 0xa4, 0x14, 0x89, 0xaa, 0xa0, 0x10, 0x16, 0xac, 0x2c,
	0xcf, 0xf8, 0x0e, 0x58, 0xf2, 0x8c, 0x2d, 0xdb, 0x49, 0xc9, 0x0a, 0x1e, 0x00, 0x21, 0x9e, 0x87,
	0x57, 0x63, 0x83, 0xe2, 0xf8, 0x26, 0x93, 0xa8, 0xd3, 0x05, 0xa2, 0x0b, 0xa4, 0xee, 0x12, 0x9d,
	0x73, 0xbe, 0x39, 0xf1, 0x75, 0xee, 0x90, 0x66, 0xae, 0xa5, 0xa5, 0x0e, 0x68, 0x92, 0x71, 0x9a,
	0xaa, 0x51, 0xe1, 0xc0, 0xd8, 0x58, 0x1b, 0xe5, 0x54, 0x74, 0x94, 0x0a, 0x9b, 0x2a, 0x2a, 0x94,
	0xa5, 0x5f, 0x0c, 0x45, 0xa3, 0xd2, 0x60, 0x62, 0xfc, 0x62, 0x1d, 0x2b, 0x78, 0x32, 0x89, 0x93,
	0x8c, 0xc7, 0x98, 0xee, 0xb4, 0xc8, 0xee, 0x65, 0x54, 0xfa, 0xe6, 0xe4, 0xe3, 0xfb, 0xce, 0xaf,
	0x75, 0xd2, 0x28, 0xab, 0xd2, 0xea, 0xb9, 0x23, 0xea, 0x92, 0x47, 0x16, 0xac, 0x15, 0xaa, 0xa0,
	0xa9, 0x01, 0xe6, 0x80, 0xc2, 0x18, 0x0a, 0x67, 0x1b, 0xb5, 0x76, 0x6d, 0xff, 0xde, 0xe0, 0x61,
	0x10, 0x7b, 0x5e, 0x3b, 0xf1, 0x52, 0xf4, 0x9c, 0x6c, 0x63, 0x66, 0xa4, 0xd1, 0x7f, 0xcb, 0xfb,
	0x1f, 0x04, 0xe1, 0x83, 0x0e, 0xde, 0x1e, 0x79, 0xba, 0xc4, 0x9f, 0x7e, 0xc8, 0x98, 0x90, 0xc0,
	0x31, 0xb8, 0xe6, 0x83, 0xad, 0xf2, 0x83, 0x84, 0x2a, 0x5e, 0x79, 0x4f, 0x80, 0xc4, 0x04, 0x7b,
	0x50, 0xae, 0x2e, 0x0a, 0x4c, 0xae, 0xfb, 0x24, 0x76, 0xe9, 0xab, 0x8b, 0x22, 0xf8, 0x5f, 0x92,
	0x5d, 0xf4, 0x33, 0x9e, 0x8b, 0xe5, 0xd4, 0x86, 0x4f, 0xd5, 0x83, 0xe1, 0x78, 0xaa, 0x97, 0xa2,
	0xa7, 0xa4, 0x8d, 0xd1, 0x4f, 0x86, 0xa5, 0x90, 0x8d, 0xa4, 0x9c, 0x50, 0x0e, 0x12, 0x16, 0x47,
	0xb3, 0xe9, 0x09, 0x7b, 0xc1, 0x77, 0x3a, 0xb7, 0xf5, 0xbd, 0x2b, 0x80, 0xce, 0xc9, 0x33, 0x04,
	0x15, 0x57, 0xc1, 0x6e, 0x7b, 0x18, 0x3e, 0xf4, 0xbc, 0x92, 0x77, 0x4c, 0xf6, 0x56, 0x06, 0xe5,
	0x44, 0x0e, 0x6a, 0xe4, 0x10, 0xb4, 0xe5, 0x41, 0xcd, 0xa5, 0x81, 0x0d, 0x67, 0x96, 0x80, 0x28,
	0xcd, 0xda, 0x80, 0x96, 0x6c, 0x82, 0xd1, 0x3b, 0x4b, 0xb3, 0x1e, 0x78, 0x6d, 0x96, 0xe9, 0xfc,
	0x5e, 0x23, 0x4f, 0xca, 0x97, 0x67, 0xde, 0xe1, 0xe6, 0x02, 0xfd, 0x67, 0x17, 0xa8, 0x72, 0xfa,
	0x5b, 0xd5, 0xd3, 0xff, 0xb6, 0x41, 0x76, 0x2e, 0x5b, 0x2c, 0xd1, 0x21, 0xa9, 0x4b, 0x66, 0x1d,
	0x4d, 0x25, 0x30, 0x03, 0xdc, 0xdf, 0x45, 0xeb, 0x58, 0xae, 0x1b, 0x5d, 0x4f, 0xdb, 0x99, 0xaa,
	0xbd, 0x99, 0x38, 0x44, 0x2d, 0xfa, 0x5e, 0x23, 0xcd, 0x29, 0x46, 0x8d, 0xc1, 0xf8, 0x35, 0xf4,
	0x19, 0xd8, 0x02, 0xda, 0x38, 0x68, 0xd7, 0xf6, 0xef, 0x76, 0xdf, 0xc5, 0x7f, 0xb7, 0x04, 0xe3,
	0xaa, 0x1d, 0x37, 0xa8, 0x27, 0x19, 0x7f, 0x3b, 0x06, 0x73, 0x66, 0xf5, 0x6b, 0x60, 0xbc, 0x87,
	0x3f, 0xe2, 0x47, 0x8d, 0xb4, 0xec, 0x15, 0x7d, 0x0e, 0xaf, 0xa9, 0xcf, 0x63, 0x5b, 0x51, 0xe8,
	0xe7, 0xea, 0xf9, 0x38, 0x26, 0xe4, 0xa2, 0xcf, 0x0b, 0xdf, 0x67, 0xf8, 0x2f, 0xfa, 0xac, 0xfe,
	0x8d, 0xcb, 0x67, 0x34, 0x64, 0x42, 0xce, 0x2b, 0x7d, 0x25, 0xdb, 0x49, 0x46, 0xb9, 0xa2, 0x32,
	0x5f, 0x14, 0x39, 0xba, 0xc6, 0x22, 0xf7, 0x93, 0xac, 0xaf, 0xce, 0x72, 0x2c, 0x90, 0x6c, 0xfa,
	0x37, 0xe3, 0xc1, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x83, 0x26, 0x1e, 0x37, 0x07, 0x00,
	0x00,
}
