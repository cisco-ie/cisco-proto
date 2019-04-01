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
// source: pim_issu_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_process_issu

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

type PimIssuBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimIssuBag_KEYS) Reset()         { *m = PimIssuBag_KEYS{} }
func (m *PimIssuBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimIssuBag_KEYS) ProtoMessage()    {}
func (*PimIssuBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da6e4ae69281e34, []int{0}
}

func (m *PimIssuBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIssuBag_KEYS.Unmarshal(m, b)
}
func (m *PimIssuBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIssuBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimIssuBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIssuBag_KEYS.Merge(m, src)
}
func (m *PimIssuBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimIssuBag_KEYS.Size(m)
}
func (m *PimIssuBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIssuBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimIssuBag_KEYS proto.InternalMessageInfo

type PimIssuBag struct {
	Informationvalid                  bool     `protobuf:"varint,50,opt,name=informationvalid,proto3" json:"informationvalid,omitempty"`
	RoleHa                            int32    `protobuf:"zigzag32,51,opt,name=role_ha,json=roleHa,proto3" json:"role_ha,omitempty"`
	RoleIssu                          int32    `protobuf:"zigzag32,52,opt,name=role_issu,json=roleIssu,proto3" json:"role_issu,omitempty"`
	PhaseIssu                         int32    `protobuf:"zigzag32,53,opt,name=phase_issu,json=phaseIssu,proto3" json:"phase_issu,omitempty"`
	LastHaRoleNotificationReceived    uint64   `protobuf:"varint,54,opt,name=last_ha_role_notification_received,json=lastHaRoleNotificationReceived,proto3" json:"last_ha_role_notification_received,omitempty"`
	LastIssuRoleNotificationReceived  uint64   `protobuf:"varint,55,opt,name=last_issu_role_notification_received,json=lastIssuRoleNotificationReceived,proto3" json:"last_issu_role_notification_received,omitempty"`
	LastIssuPhaseNotificationReceived uint64   `protobuf:"varint,56,opt,name=last_issu_phase_notification_received,json=lastIssuPhaseNotificationReceived,proto3" json:"last_issu_phase_notification_received,omitempty"`
	IsEocReceived                     bool     `protobuf:"varint,57,opt,name=is_eoc_received,json=isEocReceived,proto3" json:"is_eoc_received,omitempty"`
	EocReceivedTimestamp              uint64   `protobuf:"varint,58,opt,name=eoc_received_timestamp,json=eocReceivedTimestamp,proto3" json:"eoc_received_timestamp,omitempty"`
	IsIhmsDoneReceived                bool     `protobuf:"varint,59,opt,name=is_ihms_done_received,json=isIhmsDoneReceived,proto3" json:"is_ihms_done_received,omitempty"`
	IhmsReceivedTimestamp             uint64   `protobuf:"varint,60,opt,name=ihms_received_timestamp,json=ihmsReceivedTimestamp,proto3" json:"ihms_received_timestamp,omitempty"`
	IsRibSyncReceived                 bool     `protobuf:"varint,61,opt,name=is_rib_sync_received,json=isRibSyncReceived,proto3" json:"is_rib_sync_received,omitempty"`
	RibSyncReceivedTimestamp          uint64   `protobuf:"varint,62,opt,name=rib_sync_received_timestamp,json=ribSyncReceivedTimestamp,proto3" json:"rib_sync_received_timestamp,omitempty"`
	IsNbrSyncReceived                 bool     `protobuf:"varint,63,opt,name=is_nbr_sync_received,json=isNbrSyncReceived,proto3" json:"is_nbr_sync_received,omitempty"`
	NbrSyncReceivedTimestamp          uint64   `protobuf:"varint,64,opt,name=nbr_sync_received_timestamp,json=nbrSyncReceivedTimestamp,proto3" json:"nbr_sync_received_timestamp,omitempty"`
	IsCheckpointIdtDone               bool     `protobuf:"varint,65,opt,name=is_checkpoint_idt_done,json=isCheckpointIdtDone,proto3" json:"is_checkpoint_idt_done,omitempty"`
	CheckpointIdtTimestamp            uint64   `protobuf:"varint,66,opt,name=checkpoint_idt_timestamp,json=checkpointIdtTimestamp,proto3" json:"checkpoint_idt_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *PimIssuBag) Reset()         { *m = PimIssuBag{} }
func (m *PimIssuBag) String() string { return proto.CompactTextString(m) }
func (*PimIssuBag) ProtoMessage()    {}
func (*PimIssuBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_5da6e4ae69281e34, []int{1}
}

func (m *PimIssuBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIssuBag.Unmarshal(m, b)
}
func (m *PimIssuBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIssuBag.Marshal(b, m, deterministic)
}
func (m *PimIssuBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIssuBag.Merge(m, src)
}
func (m *PimIssuBag) XXX_Size() int {
	return xxx_messageInfo_PimIssuBag.Size(m)
}
func (m *PimIssuBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIssuBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimIssuBag proto.InternalMessageInfo

func (m *PimIssuBag) GetInformationvalid() bool {
	if m != nil {
		return m.Informationvalid
	}
	return false
}

func (m *PimIssuBag) GetRoleHa() int32 {
	if m != nil {
		return m.RoleHa
	}
	return 0
}

func (m *PimIssuBag) GetRoleIssu() int32 {
	if m != nil {
		return m.RoleIssu
	}
	return 0
}

func (m *PimIssuBag) GetPhaseIssu() int32 {
	if m != nil {
		return m.PhaseIssu
	}
	return 0
}

func (m *PimIssuBag) GetLastHaRoleNotificationReceived() uint64 {
	if m != nil {
		return m.LastHaRoleNotificationReceived
	}
	return 0
}

func (m *PimIssuBag) GetLastIssuRoleNotificationReceived() uint64 {
	if m != nil {
		return m.LastIssuRoleNotificationReceived
	}
	return 0
}

func (m *PimIssuBag) GetLastIssuPhaseNotificationReceived() uint64 {
	if m != nil {
		return m.LastIssuPhaseNotificationReceived
	}
	return 0
}

func (m *PimIssuBag) GetIsEocReceived() bool {
	if m != nil {
		return m.IsEocReceived
	}
	return false
}

func (m *PimIssuBag) GetEocReceivedTimestamp() uint64 {
	if m != nil {
		return m.EocReceivedTimestamp
	}
	return 0
}

func (m *PimIssuBag) GetIsIhmsDoneReceived() bool {
	if m != nil {
		return m.IsIhmsDoneReceived
	}
	return false
}

func (m *PimIssuBag) GetIhmsReceivedTimestamp() uint64 {
	if m != nil {
		return m.IhmsReceivedTimestamp
	}
	return 0
}

func (m *PimIssuBag) GetIsRibSyncReceived() bool {
	if m != nil {
		return m.IsRibSyncReceived
	}
	return false
}

func (m *PimIssuBag) GetRibSyncReceivedTimestamp() uint64 {
	if m != nil {
		return m.RibSyncReceivedTimestamp
	}
	return 0
}

func (m *PimIssuBag) GetIsNbrSyncReceived() bool {
	if m != nil {
		return m.IsNbrSyncReceived
	}
	return false
}

func (m *PimIssuBag) GetNbrSyncReceivedTimestamp() uint64 {
	if m != nil {
		return m.NbrSyncReceivedTimestamp
	}
	return 0
}

func (m *PimIssuBag) GetIsCheckpointIdtDone() bool {
	if m != nil {
		return m.IsCheckpointIdtDone
	}
	return false
}

func (m *PimIssuBag) GetCheckpointIdtTimestamp() uint64 {
	if m != nil {
		return m.CheckpointIdtTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*PimIssuBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.issu.pim_issu_bag_KEYS")
	proto.RegisterType((*PimIssuBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.issu.pim_issu_bag")
}

func init() { proto.RegisterFile("pim_issu_bag.proto", fileDescriptor_5da6e4ae69281e34) }

var fileDescriptor_5da6e4ae69281e34 = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x86, 0x09, 0x48, 0x3f, 0x06, 0x45, 0x33, 0x6d, 0xd3, 0x85, 0xa2, 0xc4, 0xa0, 0x12, 0xbc,
	0x88, 0xd8, 0xc4, 0x5a, 0x3f, 0xea, 0x77, 0xa1, 0x51, 0x08, 0x65, 0xeb, 0x8d, 0x57, 0x87, 0xd9,
	0xc9, 0xd4, 0x3d, 0x98, 0x9d, 0x59, 0xe6, 0x4c, 0x83, 0xfd, 0x21, 0xfe, 0x5f, 0x99, 0xd3, 0x26,
	0xbb, 0x4d, 0xa2, 0x97, 0x39, 0xef, 0x73, 0xde, 0x27, 0x67, 0x60, 0x85, 0x2c, 0xb1, 0x00, 0x24,
	0xba, 0x80, 0x4c, 0xfd, 0xec, 0x95, 0xde, 0x05, 0x27, 0xf7, 0x35, 0x92, 0x76, 0x80, 0x8e, 0xe0,
	0xb7, 0x07, 0x2c, 0xa7, 0x03, 0x88, 0x94, 0x2b, 0x8d, 0xef, 0x95, 0x58, 0xf4, 0x94, 0x0e, 0x38,
	0x35, 0x11, 0xd6, 0x86, 0xa8, 0x17, 0xb7, 0x3b, 0x5b, 0xa2, 0x59, 0x6f, 0x82, 0x6f, 0xc7, 0x3f,
	0xce, 0x3a, 0x7f, 0xd6, 0xc5, 0xed, 0xfa, 0x54, 0x3e, 0x15, 0xf7, 0xd0, 0x9e, 0x3b, 0x5f, 0xa8,
	0x80, 0xce, 0x4e, 0xd5, 0x04, 0xc7, 0xc9, 0x7e, 0xbb, 0xd1, 0xdd, 0x48, 0x97, 0xe6, 0x72, 0x57,
	0xac, 0x7b, 0x37, 0x31, 0x90, 0xab, 0xa4, 0xdf, 0x6e, 0x74, 0x9b, 0xe9, 0x5a, 0xfc, 0x79, 0xa2,
	0xe4, 0x9e, 0xd8, 0xe4, 0x20, 0xb6, 0x26, 0x03, 0x8e, 0x36, 0xe2, 0x60, 0x48, 0x74, 0x21, 0xef,
	0x0b, 0x51, 0xe6, 0x8a, 0xae, 0xd3, 0x17, 0x9c, 0x6e, 0xf2, 0x84, 0xe3, 0xaf, 0xa2, 0x33, 0x51,
	0x14, 0x20, 0x57, 0xc0, 0x1d, 0xd6, 0x05, 0x3c, 0x47, 0xcd, 0x5a, 0xf0, 0x46, 0x1b, 0x9c, 0x9a,
	0x71, 0x72, 0xd0, 0x6e, 0x74, 0x6f, 0xa5, 0x0f, 0x22, 0x79, 0xa2, 0x52, 0x37, 0x31, 0xa3, 0x1a,
	0x96, 0x5e, 0x53, 0x72, 0x24, 0x1e, 0x71, 0x17, 0x5f, 0xf7, 0x9f, 0xb6, 0x97, 0xdc, 0xd6, 0x8e,
	0x6c, 0xfc, 0x0f, 0xff, 0xec, 0x3b, 0x15, 0x8f, 0xab, 0xbe, 0xab, 0x23, 0x56, 0x17, 0x1e, 0x72,
	0xe1, 0xc3, 0x59, 0xe1, 0x69, 0x44, 0x57, 0x36, 0x3e, 0x11, 0x77, 0x91, 0xc0, 0x38, 0x5d, 0xed,
	0xbe, 0xe2, 0xd7, 0xbe, 0x83, 0x74, 0xec, 0xf4, 0x9c, 0x1b, 0x88, 0x56, 0x1d, 0x82, 0x80, 0x85,
	0xa1, 0xa0, 0x8a, 0x32, 0x79, 0xcd, 0xaa, 0x6d, 0x53, 0xc1, 0xdf, 0x67, 0x99, 0x7c, 0x2e, 0x76,
	0x90, 0x00, 0xf3, 0x82, 0x60, 0xec, 0xac, 0xa9, 0x1c, 0x6f, 0xd8, 0x21, 0x91, 0x86, 0x79, 0x41,
	0x5f, 0x9c, 0x35, 0x73, 0xd1, 0x81, 0xd8, 0x65, 0x7e, 0x85, 0xe9, 0x2d, 0x9b, 0x76, 0x62, 0xbc,
	0xac, 0x7a, 0x26, 0xb6, 0x91, 0xc0, 0x63, 0x06, 0x74, 0x69, 0x6b, 0xd7, 0x1c, 0xb1, 0xa9, 0x89,
	0x94, 0x62, 0x76, 0x76, 0x69, 0xab, 0x8b, 0x8e, 0xc4, 0xde, 0x12, 0x5d, 0x93, 0xbd, 0x63, 0x59,
	0xe2, 0x6f, 0x6e, 0x2d, 0xfa, 0x6c, 0xe6, 0x17, 0x7c, 0xef, 0x67, 0xbe, 0x51, 0xe6, 0x17, 0x7d,
	0x4b, 0x74, 0xcd, 0xf7, 0xe1, 0xca, 0x67, 0x6f, 0x6e, 0x55, 0xbe, 0xbe, 0x68, 0x21, 0x81, 0xce,
	0x8d, 0xfe, 0x55, 0x3a, 0xb4, 0x01, 0x70, 0x1c, 0xf8, 0x51, 0x93, 0x8f, 0x6c, 0xdc, 0x42, 0xfa,
	0x3c, 0x0f, 0x87, 0xe3, 0x10, 0x1f, 0x55, 0x1e, 0x8a, 0x64, 0x61, 0xa3, 0x12, 0x7e, 0x62, 0x61,
	0x4b, 0xd7, 0x97, 0xe6, 0xba, 0x6c, 0x8d, 0xbf, 0xf3, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd8, 0x1d, 0x67, 0xfd, 0xfd, 0x03, 0x00, 0x00,
}