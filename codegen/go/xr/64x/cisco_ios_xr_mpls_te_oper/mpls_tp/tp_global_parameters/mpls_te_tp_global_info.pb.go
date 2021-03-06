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
// source: mpls_te_tp_global_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_tp_tp_global_parameters

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

type MplsTeTpGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpGlobalInfo_KEYS) Reset()         { *m = MplsTeTpGlobalInfo_KEYS{} }
func (m *MplsTeTpGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpGlobalInfo_KEYS) ProtoMessage()    {}
func (*MplsTeTpGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a1b64c77fc0e1d, []int{0}
}

func (m *MplsTeTpGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTpGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeTpGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpGlobalInfo_KEYS.Merge(m, src)
}
func (m *MplsTeTpGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpGlobalInfo_KEYS.Size(m)
}
func (m *MplsTeTpGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpGlobalInfo_KEYS proto.InternalMessageInfo

type MplsTeTpOamProtTrigger struct {
	Ais                  bool     `protobuf:"varint,1,opt,name=ais,proto3" json:"ais,omitempty"`
	Ldi                  bool     `protobuf:"varint,2,opt,name=ldi,proto3" json:"ldi,omitempty"`
	Lkr                  bool     `protobuf:"varint,3,opt,name=lkr,proto3" json:"lkr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpOamProtTrigger) Reset()         { *m = MplsTeTpOamProtTrigger{} }
func (m *MplsTeTpOamProtTrigger) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpOamProtTrigger) ProtoMessage()    {}
func (*MplsTeTpOamProtTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a1b64c77fc0e1d, []int{1}
}

func (m *MplsTeTpOamProtTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpOamProtTrigger.Unmarshal(m, b)
}
func (m *MplsTeTpOamProtTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpOamProtTrigger.Marshal(b, m, deterministic)
}
func (m *MplsTeTpOamProtTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpOamProtTrigger.Merge(m, src)
}
func (m *MplsTeTpOamProtTrigger) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpOamProtTrigger.Size(m)
}
func (m *MplsTeTpOamProtTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpOamProtTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpOamProtTrigger proto.InternalMessageInfo

func (m *MplsTeTpOamProtTrigger) GetAis() bool {
	if m != nil {
		return m.Ais
	}
	return false
}

func (m *MplsTeTpOamProtTrigger) GetLdi() bool {
	if m != nil {
		return m.Ldi
	}
	return false
}

func (m *MplsTeTpOamProtTrigger) GetLkr() bool {
	if m != nil {
		return m.Lkr
	}
	return false
}

type MplsTeTpGlobalInfo struct {
	NodeId                       string                  `protobuf:"bytes,50,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	GlobalId                     uint32                  `protobuf:"varint,51,opt,name=global_id,json=globalId,proto3" json:"global_id,omitempty"`
	OamRefresh                   uint32                  `protobuf:"varint,52,opt,name=oam_refresh,json=oamRefresh,proto3" json:"oam_refresh,omitempty"`
	BfdInterval                  uint32                  `protobuf:"varint,53,opt,name=bfd_interval,json=bfdInterval,proto3" json:"bfd_interval,omitempty"`
	BfdIntervalInMicroSec        bool                    `protobuf:"varint,54,opt,name=bfd_interval_in_micro_sec,json=bfdIntervalInMicroSec,proto3" json:"bfd_interval_in_micro_sec,omitempty"`
	BfdIntervalStandby           uint32                  `protobuf:"varint,55,opt,name=bfd_interval_standby,json=bfdIntervalStandby,proto3" json:"bfd_interval_standby,omitempty"`
	BfdIntervalStandbyInMicroSec bool                    `protobuf:"varint,56,opt,name=bfd_interval_standby_in_micro_sec,json=bfdIntervalStandbyInMicroSec,proto3" json:"bfd_interval_standby_in_micro_sec,omitempty"`
	BfdMultiplier                uint32                  `protobuf:"varint,57,opt,name=bfd_multiplier,json=bfdMultiplier,proto3" json:"bfd_multiplier,omitempty"`
	BfdMultiplierStandby         uint32                  `protobuf:"varint,58,opt,name=bfd_multiplier_standby,json=bfdMultiplierStandby,proto3" json:"bfd_multiplier_standby,omitempty"`
	WaitToRestoreInterval        uint32                  `protobuf:"varint,59,opt,name=wait_to_restore_interval,json=waitToRestoreInterval,proto3" json:"wait_to_restore_interval,omitempty"`
	OamProtectionTriggers        *MplsTeTpOamProtTrigger `protobuf:"bytes,60,opt,name=oam_protection_triggers,json=oamProtectionTriggers,proto3" json:"oam_protection_triggers,omitempty"`
	AlarmSuppression             bool                    `protobuf:"varint,61,opt,name=alarm_suppression,json=alarmSuppression,proto3" json:"alarm_suppression,omitempty"`
	SoakTime                     uint32                  `protobuf:"varint,62,opt,name=soak_time,json=soakTime,proto3" json:"soak_time,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                `json:"-"`
	XXX_unrecognized             []byte                  `json:"-"`
	XXX_sizecache                int32                   `json:"-"`
}

func (m *MplsTeTpGlobalInfo) Reset()         { *m = MplsTeTpGlobalInfo{} }
func (m *MplsTeTpGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpGlobalInfo) ProtoMessage()    {}
func (*MplsTeTpGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92a1b64c77fc0e1d, []int{2}
}

func (m *MplsTeTpGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpGlobalInfo.Unmarshal(m, b)
}
func (m *MplsTeTpGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpGlobalInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeTpGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpGlobalInfo.Merge(m, src)
}
func (m *MplsTeTpGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpGlobalInfo.Size(m)
}
func (m *MplsTeTpGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpGlobalInfo proto.InternalMessageInfo

func (m *MplsTeTpGlobalInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MplsTeTpGlobalInfo) GetGlobalId() uint32 {
	if m != nil {
		return m.GlobalId
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetOamRefresh() uint32 {
	if m != nil {
		return m.OamRefresh
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetBfdInterval() uint32 {
	if m != nil {
		return m.BfdInterval
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetBfdIntervalInMicroSec() bool {
	if m != nil {
		return m.BfdIntervalInMicroSec
	}
	return false
}

func (m *MplsTeTpGlobalInfo) GetBfdIntervalStandby() uint32 {
	if m != nil {
		return m.BfdIntervalStandby
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetBfdIntervalStandbyInMicroSec() bool {
	if m != nil {
		return m.BfdIntervalStandbyInMicroSec
	}
	return false
}

func (m *MplsTeTpGlobalInfo) GetBfdMultiplier() uint32 {
	if m != nil {
		return m.BfdMultiplier
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetBfdMultiplierStandby() uint32 {
	if m != nil {
		return m.BfdMultiplierStandby
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetWaitToRestoreInterval() uint32 {
	if m != nil {
		return m.WaitToRestoreInterval
	}
	return 0
}

func (m *MplsTeTpGlobalInfo) GetOamProtectionTriggers() *MplsTeTpOamProtTrigger {
	if m != nil {
		return m.OamProtectionTriggers
	}
	return nil
}

func (m *MplsTeTpGlobalInfo) GetAlarmSuppression() bool {
	if m != nil {
		return m.AlarmSuppression
	}
	return false
}

func (m *MplsTeTpGlobalInfo) GetSoakTime() uint32 {
	if m != nil {
		return m.SoakTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeTpGlobalInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_global_parameters.mpls_te_tp_global_info_KEYS")
	proto.RegisterType((*MplsTeTpOamProtTrigger)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_global_parameters.mpls_te_tp_oam_prot_trigger")
	proto.RegisterType((*MplsTeTpGlobalInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_global_parameters.mpls_te_tp_global_info")
}

func init() { proto.RegisterFile("mpls_te_tp_global_info.proto", fileDescriptor_92a1b64c77fc0e1d) }

var fileDescriptor_92a1b64c77fc0e1d = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdd, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0x2b, 0xeb, 0x5a, 0x65, 0x1d, 0x9d, 0x68, 0x5a, 0x8d, 0x76, 0x2c, 0x0d, 0x0c, 0x02,
	0x83, 0x30, 0xda, 0xae, 0xed, 0x7e, 0xef, 0xc6, 0x08, 0xa3, 0x30, 0xec, 0xdc, 0xec, 0xea, 0x20,
	0x5b, 0x72, 0x26, 0x62, 0x59, 0xe2, 0x48, 0xdd, 0xcf, 0x33, 0xec, 0x79, 0xf6, 0x7e, 0x43, 0x72,
	0x9c, 0xc4, 0x2c, 0xec, 0xa2, 0x77, 0xd6, 0xf7, 0x73, 0xbe, 0xc3, 0xe1, 0x33, 0x39, 0xd1, 0xb6,
	0x72, 0xe0, 0x25, 0x78, 0x0b, 0xb3, 0xca, 0xe4, 0xbc, 0x02, 0x55, 0x97, 0x66, 0x6c, 0xd1, 0x78,
	0x43, 0x2f, 0x0b, 0xe5, 0x0a, 0x03, 0xca, 0x38, 0xf8, 0x89, 0xd0, 0x4a, 0x8d, 0x95, 0x38, 0x6e,
	0x1e, 0x76, 0xbc, 0xf2, 0x59, 0x8e, 0x5c, 0x4b, 0x2f, 0xd1, 0x0d, 0x9f, 0x92, 0xe3, 0xcd, 0x73,
	0xe1, 0xf3, 0xc7, 0xaf, 0xd9, 0x30, 0xeb, 0xd0, 0x86, 0x6b, 0x08, 0x71, 0xe0, 0x51, 0xcd, 0x66,
	0x12, 0xe9, 0x3e, 0xd9, 0xe2, 0xca, 0xb1, 0x64, 0x90, 0x8c, 0x76, 0xd2, 0xf0, 0x19, 0x90, 0x4a,
	0x28, 0x76, 0xaf, 0x41, 0x2a, 0xa1, 0x22, 0x32, 0x47, 0xb6, 0xb5, 0x40, 0xe6, 0x38, 0xfc, 0x73,
	0x9f, 0x1c, 0x6e, 0x0e, 0xa5, 0x47, 0xe4, 0x41, 0x6d, 0x84, 0x04, 0x25, 0xd8, 0xd9, 0x20, 0x19,
	0xed, 0xa6, 0xdb, 0xe1, 0x39, 0x11, 0xf4, 0x98, 0xec, 0xb6, 0x3a, 0xc1, 0xce, 0x07, 0xc9, 0x68,
	0x2f, 0xdd, 0x69, 0x80, 0x89, 0xa0, 0xcf, 0x48, 0x2f, 0xac, 0x86, 0xb2, 0x44, 0xe9, 0xbe, 0xb1,
	0x8b, 0x48, 0x13, 0xc3, 0x75, 0xda, 0x20, 0xf4, 0x94, 0x3c, 0xcc, 0x4b, 0x01, 0xaa, 0xf6, 0x12,
	0xbf, 0xf3, 0x8a, 0xbd, 0x8a, 0x8a, 0x5e, 0x5e, 0x8a, 0xc9, 0x02, 0xa2, 0xd7, 0xe4, 0xc9, 0xba,
	0x04, 0x54, 0x0d, 0x5a, 0x15, 0x68, 0xc0, 0xc9, 0x82, 0x5d, 0xc6, 0xe5, 0xfb, 0x6b, 0xfa, 0x49,
	0x7d, 0x13, 0xd8, 0x4c, 0x16, 0xf4, 0x25, 0x39, 0xe8, 0x38, 0x9d, 0xe7, 0xb5, 0xc8, 0x7f, 0xb1,
	0xab, 0x18, 0x42, 0xd7, 0x4c, 0x59, 0xc3, 0xd0, 0x4f, 0xe4, 0x74, 0x93, 0xa3, 0x9b, 0x79, 0x1d,
	0x33, 0x4f, 0xfe, 0xb5, 0xaf, 0x45, 0x3f, 0x27, 0x8f, 0xc2, 0x20, 0x7d, 0x5b, 0x79, 0x65, 0x2b,
	0x25, 0x91, 0xbd, 0x8e, 0xa1, 0x7b, 0x79, 0x29, 0x6e, 0x96, 0x20, 0xbd, 0x20, 0x87, 0x5d, 0xd9,
	0x72, 0xc7, 0x37, 0x51, 0x7e, 0xd0, 0x91, 0xb7, 0x5b, 0x5e, 0x11, 0xf6, 0x83, 0x2b, 0x0f, 0xde,
	0x00, 0x4a, 0xe7, 0x0d, 0xca, 0xd5, 0x01, 0xdf, 0x46, 0x5f, 0x3f, 0xf0, 0x53, 0x93, 0x36, 0xec,
	0xf2, 0x94, 0xbf, 0x13, 0x72, 0xd4, 0x56, 0x45, 0x16, 0x5e, 0x99, 0xba, 0x2d, 0x8c, 0x63, 0xef,
	0x06, 0xc9, 0xa8, 0x77, 0x96, 0x8d, 0xef, 0x56, 0xd7, 0xf1, 0x7f, 0xca, 0x98, 0xf6, 0x0d, 0xd7,
	0x5f, 0x96, 0x91, 0xd3, 0x45, 0x22, 0x7d, 0x41, 0x1e, 0xf3, 0x8a, 0xa3, 0x06, 0x77, 0x6b, 0x2d,
	0x4a, 0xe7, 0x94, 0xa9, 0xd9, 0xfb, 0x78, 0xdc, 0xfd, 0x48, 0x64, 0x2b, 0x3c, 0xd4, 0xcc, 0x19,
	0x3e, 0x07, 0xaf, 0xb4, 0x64, 0x1f, 0x9a, 0x9a, 0x05, 0x60, 0xaa, 0xb4, 0xcc, 0xb7, 0xe3, 0xaf,
	0x76, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xaf, 0x48, 0x34, 0x8a, 0x03, 0x00, 0x00,
}
