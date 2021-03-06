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
// source: te_pp_sw_log_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_path_protection_switchover_log

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

type TePpSwLogBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TePpSwLogBag_KEYS) Reset()         { *m = TePpSwLogBag_KEYS{} }
func (m *TePpSwLogBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogBag_KEYS) ProtoMessage()    {}
func (*TePpSwLogBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e2d89a4301bad7c, []int{0}
}

func (m *TePpSwLogBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Unmarshal(m, b)
}
func (m *TePpSwLogBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TePpSwLogBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogBag_KEYS.Merge(m, src)
}
func (m *TePpSwLogBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogBag_KEYS.Size(m)
}
func (m *TePpSwLogBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogBag_KEYS proto.InternalMessageInfo

type TePpSwLogEntryBag struct {
	PathProtectionSwitchoverEventIndex uint32   `protobuf:"varint,1,opt,name=path_protection_switchover_event_index,json=pathProtectionSwitchoverEventIndex,proto3" json:"path_protection_switchover_event_index,omitempty"`
	PathProtectionTunnelId             uint32   `protobuf:"varint,2,opt,name=path_protection_tunnel_id,json=pathProtectionTunnelId,proto3" json:"path_protection_tunnel_id,omitempty"`
	FromLspId                          uint32   `protobuf:"varint,3,opt,name=from_lsp_id,json=fromLspId,proto3" json:"from_lsp_id,omitempty"`
	ToLspId                            uint32   `protobuf:"varint,4,opt,name=to_lsp_id,json=toLspId,proto3" json:"to_lsp_id,omitempty"`
	DateOfErrorDetection               uint32   `protobuf:"varint,5,opt,name=date_of_error_detection,json=dateOfErrorDetection,proto3" json:"date_of_error_detection,omitempty"`
	DateOfErrorDetectionMillisec       uint32   `protobuf:"varint,6,opt,name=date_of_error_detection_millisec,json=dateOfErrorDetectionMillisec,proto3" json:"date_of_error_detection_millisec,omitempty"`
	SwitchoverDurationMillisec         uint32   `protobuf:"varint,7,opt,name=switchover_duration_millisec,json=switchoverDurationMillisec,proto3" json:"switchover_duration_millisec,omitempty"`
	PathProtectionSwitchoverReason     string   `protobuf:"bytes,8,opt,name=path_protection_switchover_reason,json=pathProtectionSwitchoverReason,proto3" json:"path_protection_switchover_reason,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *TePpSwLogEntryBag) Reset()         { *m = TePpSwLogEntryBag{} }
func (m *TePpSwLogEntryBag) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogEntryBag) ProtoMessage()    {}
func (*TePpSwLogEntryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e2d89a4301bad7c, []int{1}
}

func (m *TePpSwLogEntryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogEntryBag.Unmarshal(m, b)
}
func (m *TePpSwLogEntryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogEntryBag.Marshal(b, m, deterministic)
}
func (m *TePpSwLogEntryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogEntryBag.Merge(m, src)
}
func (m *TePpSwLogEntryBag) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogEntryBag.Size(m)
}
func (m *TePpSwLogEntryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogEntryBag.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogEntryBag proto.InternalMessageInfo

func (m *TePpSwLogEntryBag) GetPathProtectionSwitchoverEventIndex() uint32 {
	if m != nil {
		return m.PathProtectionSwitchoverEventIndex
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetPathProtectionTunnelId() uint32 {
	if m != nil {
		return m.PathProtectionTunnelId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetFromLspId() uint32 {
	if m != nil {
		return m.FromLspId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetToLspId() uint32 {
	if m != nil {
		return m.ToLspId
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetDateOfErrorDetection() uint32 {
	if m != nil {
		return m.DateOfErrorDetection
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetDateOfErrorDetectionMillisec() uint32 {
	if m != nil {
		return m.DateOfErrorDetectionMillisec
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetSwitchoverDurationMillisec() uint32 {
	if m != nil {
		return m.SwitchoverDurationMillisec
	}
	return 0
}

func (m *TePpSwLogEntryBag) GetPathProtectionSwitchoverReason() string {
	if m != nil {
		return m.PathProtectionSwitchoverReason
	}
	return ""
}

type TePpSwLogBag struct {
	PathProtectionSwitchovers     uint32               `protobuf:"varint,50,opt,name=path_protection_switchovers,json=pathProtectionSwitchovers,proto3" json:"path_protection_switchovers,omitempty"`
	MaximumSwitchoverMillisec     uint32               `protobuf:"varint,51,opt,name=maximum_switchover_millisec,json=maximumSwitchoverMillisec,proto3" json:"maximum_switchover_millisec,omitempty"`
	AverageSwitchoverMillisec     uint32               `protobuf:"varint,52,opt,name=average_switchover_millisec,json=averageSwitchoverMillisec,proto3" json:"average_switchover_millisec,omitempty"`
	PathProtectionSwitchoverEntry []*TePpSwLogEntryBag `protobuf:"bytes,53,rep,name=path_protection_switchover_entry,json=pathProtectionSwitchoverEntry,proto3" json:"path_protection_switchover_entry,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *TePpSwLogBag) Reset()         { *m = TePpSwLogBag{} }
func (m *TePpSwLogBag) String() string { return proto.CompactTextString(m) }
func (*TePpSwLogBag) ProtoMessage()    {}
func (*TePpSwLogBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e2d89a4301bad7c, []int{2}
}

func (m *TePpSwLogBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TePpSwLogBag.Unmarshal(m, b)
}
func (m *TePpSwLogBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TePpSwLogBag.Marshal(b, m, deterministic)
}
func (m *TePpSwLogBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TePpSwLogBag.Merge(m, src)
}
func (m *TePpSwLogBag) XXX_Size() int {
	return xxx_messageInfo_TePpSwLogBag.Size(m)
}
func (m *TePpSwLogBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TePpSwLogBag.DiscardUnknown(m)
}

var xxx_messageInfo_TePpSwLogBag proto.InternalMessageInfo

func (m *TePpSwLogBag) GetPathProtectionSwitchovers() uint32 {
	if m != nil {
		return m.PathProtectionSwitchovers
	}
	return 0
}

func (m *TePpSwLogBag) GetMaximumSwitchoverMillisec() uint32 {
	if m != nil {
		return m.MaximumSwitchoverMillisec
	}
	return 0
}

func (m *TePpSwLogBag) GetAverageSwitchoverMillisec() uint32 {
	if m != nil {
		return m.AverageSwitchoverMillisec
	}
	return 0
}

func (m *TePpSwLogBag) GetPathProtectionSwitchoverEntry() []*TePpSwLogEntryBag {
	if m != nil {
		return m.PathProtectionSwitchoverEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*TePpSwLogBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.path_protection.switchover_log.te_pp_sw_log_bag_KEYS")
	proto.RegisterType((*TePpSwLogEntryBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.path_protection.switchover_log.te_pp_sw_log_entry_bag")
	proto.RegisterType((*TePpSwLogBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.path_protection.switchover_log.te_pp_sw_log_bag")
}

func init() { proto.RegisterFile("te_pp_sw_log_bag.proto", fileDescriptor_3e2d89a4301bad7c) }

var fileDescriptor_3e2d89a4301bad7c = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xd5, 0x15, 0x36, 0xea, 0x09, 0x09, 0x59, 0xb0, 0x65, 0x2f, 0x4c, 0xa1, 0x07, 0xd4,
	0x53, 0x0e, 0x1b, 0x3b, 0x70, 0x99, 0x38, 0xac, 0x88, 0x08, 0x10, 0x28, 0xe3, 0xc2, 0xe9, 0xc1,
	0x8d, 0x9f, 0x76, 0x96, 0x12, 0xdb, 0xb2, 0xdd, 0xae, 0xfb, 0x48, 0x7c, 0x3a, 0xbe, 0x01, 0x42,
	0x76, 0xd3, 0xac, 0x8b, 0x92, 0x9e, 0x38, 0xb6, 0xff, 0xdf, 0xff, 0x17, 0xbf, 0x3d, 0xe4, 0xc0,
	0x21, 0x68, 0x0d, 0xf6, 0x0e, 0x0a, 0x35, 0x83, 0x09, 0x9b, 0x25, 0xda, 0x28, 0xa7, 0xe8, 0xa7,
	0x5c, 0xd8, 0x5c, 0x81, 0x50, 0x16, 0x96, 0x06, 0x4a, 0x5d, 0x58, 0x70, 0x08, 0x4a, 0xa3, 0x49,
	0xd6, 0x3f, 0xac, 0x63, 0x92, 0x4f, 0xee, 0x13, 0xcd, 0xdc, 0x2d, 0xf8, 0x16, 0xe6, 0x4e, 0x28,
	0x99, 0xd8, 0x3b, 0xe1, 0xf2, 0x5b, 0xb5, 0x40, 0xe3, 0x9d, 0xc3, 0x43, 0xf2, 0xaa, 0xf9, 0x0d,
	0xf8, 0x3c, 0xfe, 0x79, 0x33, 0xfc, 0xd3, 0x6f, 0x7c, 0x1d, 0xa5, 0x33, 0xf7, 0x3e, 0xa7, 0x19,
	0x79, 0xdb, 0xb0, 0xc2, 0x86, 0x15, 0x17, 0x28, 0x1d, 0x08, 0xc9, 0x71, 0x19, 0xf5, 0xe2, 0xde,
	0xe8, 0x79, 0x36, 0xf4, 0xf4, 0xf7, 0x1a, 0xbe, 0xa9, 0xd9, 0xb1, 0x47, 0x53, 0x4f, 0xd2, 0xf7,
	0xe4, 0xa8, 0xe9, 0x74, 0x73, 0x29, 0xb1, 0x00, 0xc1, 0xa3, 0x9d, 0xa0, 0x39, 0x78, 0xac, 0xf9,
	0x11, 0xe2, 0x94, 0xd3, 0x33, 0xb2, 0x3f, 0x35, 0xaa, 0x84, 0xc2, 0x6a, 0x0f, 0xf7, 0x03, 0x3c,
	0xf0, 0x7f, 0x7d, 0xb1, 0x3a, 0xe5, 0xf4, 0x98, 0x0c, 0x9c, 0x5a, 0xa7, 0x4f, 0x42, 0xba, 0xe7,
	0xd4, 0x2a, 0xbb, 0x24, 0x87, 0x9c, 0xf9, 0xa3, 0x9b, 0x02, 0x1a, 0xa3, 0x0c, 0x70, 0xac, 0xf4,
	0xd1, 0xd3, 0x40, 0xbe, 0xf4, 0xf1, 0xb7, 0xe9, 0xd8, 0x87, 0xd7, 0xeb, 0x8c, 0x7e, 0x24, 0x71,
	0x47, 0x0d, 0x4a, 0x51, 0x14, 0xc2, 0x62, 0x1e, 0xed, 0x86, 0xfe, 0x69, 0x5b, 0xff, 0x6b, 0xc5,
	0xd0, 0x0f, 0xe4, 0x74, 0xe3, 0xe4, 0xf8, 0xdc, 0xb0, 0xc7, 0x8e, 0xbd, 0xe0, 0x38, 0x7e, 0x60,
	0xae, 0x2b, 0xa4, 0x36, 0xa4, 0xe4, 0xcd, 0x96, 0xbb, 0x30, 0xc8, 0xac, 0x92, 0xd1, 0xb3, 0xb8,
	0x37, 0x1a, 0x64, 0x67, 0x5d, 0xd7, 0x90, 0x05, 0x6a, 0xf8, 0x77, 0x87, 0xbc, 0x68, 0xbe, 0x05,
	0x7a, 0x45, 0x4e, 0xba, 0xfd, 0x36, 0x3a, 0x0f, 0x0b, 0x3c, 0xea, 0x32, 0x5b, 0xdf, 0x2f, 0xd9,
	0x52, 0x94, 0xf3, 0x72, 0x73, 0x5d, 0xf5, 0x06, 0x2f, 0x56, 0xfd, 0x0a, 0x79, 0x28, 0xd6, 0xfb,
	0xbb, 0x22, 0x27, 0x6c, 0x81, 0x86, 0xcd, 0xb0, 0xb5, 0xff, 0x6e, 0xd5, 0xaf, 0x90, 0x96, 0xfe,
	0xef, 0x1e, 0x89, 0xb7, 0x3d, 0x56, 0xff, 0xa8, 0xa3, 0xcb, 0xb8, 0x3f, 0xda, 0x3f, 0xff, 0x95,
	0xfc, 0xaf, 0xa9, 0x4a, 0xda, 0x07, 0x27, 0x7b, 0xdd, 0x39, 0x08, 0x1e, 0x99, 0xec, 0x86, 0xe1,
	0xbe, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x72, 0x07, 0xfb, 0xb2, 0xf6, 0x03, 0x00, 0x00,
}
