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
// source: l2vpn_xc_summary.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_xconnect_summary

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

type L2VpnXcSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnXcSummary_KEYS) Reset()         { *m = L2VpnXcSummary_KEYS{} }
func (m *L2VpnXcSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcSummary_KEYS) ProtoMessage()    {}
func (*L2VpnXcSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_922b751c698ee021, []int{0}
}

func (m *L2VpnXcSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcSummary_KEYS.Unmarshal(m, b)
}
func (m *L2VpnXcSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnXcSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcSummary_KEYS.Merge(m, src)
}
func (m *L2VpnXcSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcSummary_KEYS.Size(m)
}
func (m *L2VpnXcSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcSummary_KEYS proto.InternalMessageInfo

type L2VpnXcSummary struct {
	NumberGroups                                     uint32   `protobuf:"varint,50,opt,name=number_groups,json=numberGroups,proto3" json:"number_groups,omitempty"`
	NumberXconnects                                  uint32   `protobuf:"varint,51,opt,name=number_xconnects,json=numberXconnects,proto3" json:"number_xconnects,omitempty"`
	NumberXconnectsUp                                uint32   `protobuf:"varint,52,opt,name=number_xconnects_up,json=numberXconnectsUp,proto3" json:"number_xconnects_up,omitempty"`
	NumberXconnectsDown                              uint32   `protobuf:"varint,53,opt,name=number_xconnects_down,json=numberXconnectsDown,proto3" json:"number_xconnects_down,omitempty"`
	NumberXconnectsUnresolved                        uint32   `protobuf:"varint,54,opt,name=number_xconnects_unresolved,json=numberXconnectsUnresolved,proto3" json:"number_xconnects_unresolved,omitempty"`
	NumberXconnectsAdministrativelyDown              uint32   `protobuf:"varint,55,opt,name=number_xconnects_administratively_down,json=numberXconnectsAdministrativelyDown,proto3" json:"number_xconnects_administratively_down,omitempty"`
	NumberXconnectsTypeAttachmentCircuitToPseudoWire uint32   `protobuf:"varint,56,opt,name=number_xconnects_type_attachment_circuit_to_pseudo_wire,json=numberXconnectsTypeAttachmentCircuitToPseudoWire,proto3" json:"number_xconnects_type_attachment_circuit_to_pseudo_wire,omitempty"`
	NumberXconnectsTypeMonitorSessionToPseudoWire    uint32   `protobuf:"varint,57,opt,name=number_xconnects_type_monitor_session_to_pseudo_wire,json=numberXconnectsTypeMonitorSessionToPseudoWire,proto3" json:"number_xconnects_type_monitor_session_to_pseudo_wire,omitempty"`
	NumberXconnectsLocallySwitched                   uint32   `protobuf:"varint,58,opt,name=number_xconnects_locally_switched,json=numberXconnectsLocallySwitched,proto3" json:"number_xconnects_locally_switched,omitempty"`
	NumberXconnectsWithBackupPw                      uint32   `protobuf:"varint,59,opt,name=number_xconnects_with_backup_pw,json=numberXconnectsWithBackupPw,proto3" json:"number_xconnects_with_backup_pw,omitempty"`
	NumberXconnectsCurrentlyUsingBackup              uint32   `protobuf:"varint,60,opt,name=number_xconnects_currently_using_backup,json=numberXconnectsCurrentlyUsingBackup,proto3" json:"number_xconnects_currently_using_backup,omitempty"`
	DownBackupXconnects                              uint32   `protobuf:"varint,61,opt,name=down_backup_xconnects,json=downBackupXconnects,proto3" json:"down_backup_xconnects,omitempty"`
	AdminDownBackupXconnects                         uint32   `protobuf:"varint,62,opt,name=admin_down_backup_xconnects,json=adminDownBackupXconnects,proto3" json:"admin_down_backup_xconnects,omitempty"`
	UnresolvedBackupXconnects                        uint32   `protobuf:"varint,63,opt,name=unresolved_backup_xconnects,json=unresolvedBackupXconnects,proto3" json:"unresolved_backup_xconnects,omitempty"`
	StandbyBackupXconnects                           uint32   `protobuf:"varint,64,opt,name=standby_backup_xconnects,json=standbyBackupXconnects,proto3" json:"standby_backup_xconnects,omitempty"`
	StandbyReadyBackupXconnects                      uint32   `protobuf:"varint,65,opt,name=standby_ready_backup_xconnects,json=standbyReadyBackupXconnects,proto3" json:"standby_ready_backup_xconnects,omitempty"`
	OutOfMemoryState                                 uint32   `protobuf:"varint,66,opt,name=out_of_memory_state,json=outOfMemoryState,proto3" json:"out_of_memory_state,omitempty"`
	NumberXconnectsTypePseudoWireToPseudoWire        uint32   `protobuf:"varint,67,opt,name=number_xconnects_type_pseudo_wire_to_pseudo_wire,json=numberXconnectsTypePseudoWireToPseudoWire,proto3" json:"number_xconnects_type_pseudo_wire_to_pseudo_wire,omitempty"`
	NumberMp2MpXconnects                             uint32   `protobuf:"varint,68,opt,name=number_mp2mp_xconnects,json=numberMp2mpXconnects,proto3" json:"number_mp2mp_xconnects,omitempty"`
	NumberMp2MpXconnectsUp                           uint32   `protobuf:"varint,69,opt,name=number_mp2mp_xconnects_up,json=numberMp2mpXconnectsUp,proto3" json:"number_mp2mp_xconnects_up,omitempty"`
	NumberMp2MpXconnectsDown                         uint32   `protobuf:"varint,70,opt,name=number_mp2mp_xconnects_down,json=numberMp2mpXconnectsDown,proto3" json:"number_mp2mp_xconnects_down,omitempty"`
	NumberMp2MpXconnectsAdvertised                   uint32   `protobuf:"varint,71,opt,name=number_mp2mp_xconnects_advertised,json=numberMp2mpXconnectsAdvertised,proto3" json:"number_mp2mp_xconnects_advertised,omitempty"`
	NumberMp2MpXconnectssNotAdvertised               uint32   `protobuf:"varint,72,opt,name=number_mp2mp_xconnectss_not_advertised,json=numberMp2mpXconnectssNotAdvertised,proto3" json:"number_mp2mp_xconnectss_not_advertised,omitempty"`
	NumberCe2Ceconnections                           uint32   `protobuf:"varint,73,opt,name=number_ce2_ceconnections,json=numberCe2Ceconnections,proto3" json:"number_ce2_ceconnections,omitempty"`
	NumberCe2CeAdvertized                            uint32   `protobuf:"varint,74,opt,name=number_ce2ce_advertized,json=numberCe2ceAdvertized,proto3" json:"number_ce2ce_advertized,omitempty"`
	NumberCe2CeNotAdvertized                         uint32   `protobuf:"varint,75,opt,name=number_ce2ce_not_advertized,json=numberCe2ceNotAdvertized,proto3" json:"number_ce2ce_not_advertized,omitempty"`
	PartiallyProgrammedXconnects                     uint32   `protobuf:"varint,76,opt,name=partially_programmed_xconnects,json=partiallyProgrammedXconnects,proto3" json:"partially_programmed_xconnects,omitempty"`
	NumberXconnectsWithBackupInterface               uint32   `protobuf:"varint,77,opt,name=number_xconnects_with_backup_interface,json=numberXconnectsWithBackupInterface,proto3" json:"number_xconnects_with_backup_interface,omitempty"`
	NumberXconnectsCurrentlyUsingBackupInterface     uint32   `protobuf:"varint,78,opt,name=number_xconnects_currently_using_backup_interface,json=numberXconnectsCurrentlyUsingBackupInterface,proto3" json:"number_xconnects_currently_using_backup_interface,omitempty"`
	DownBackupInterfaceXconnects                     uint32   `protobuf:"varint,79,opt,name=down_backup_interface_xconnects,json=downBackupInterfaceXconnects,proto3" json:"down_backup_interface_xconnects,omitempty"`
	AdminDownBackupInterfaceXconnects                uint32   `protobuf:"varint,80,opt,name=admin_down_backup_interface_xconnects,json=adminDownBackupInterfaceXconnects,proto3" json:"admin_down_backup_interface_xconnects,omitempty"`
	UnresolvedBackupInterfaceXconnects               uint32   `protobuf:"varint,81,opt,name=unresolved_backup_interface_xconnects,json=unresolvedBackupInterfaceXconnects,proto3" json:"unresolved_backup_interface_xconnects,omitempty"`
	StandbyBackupInterfaceXconnects                  uint32   `protobuf:"varint,82,opt,name=standby_backup_interface_xconnects,json=standbyBackupInterfaceXconnects,proto3" json:"standby_backup_interface_xconnects,omitempty"`
	XXX_NoUnkeyedLiteral                             struct{} `json:"-"`
	XXX_unrecognized                                 []byte   `json:"-"`
	XXX_sizecache                                    int32    `json:"-"`
}

func (m *L2VpnXcSummary) Reset()         { *m = L2VpnXcSummary{} }
func (m *L2VpnXcSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcSummary) ProtoMessage()    {}
func (*L2VpnXcSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_922b751c698ee021, []int{1}
}

func (m *L2VpnXcSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcSummary.Unmarshal(m, b)
}
func (m *L2VpnXcSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnXcSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcSummary.Merge(m, src)
}
func (m *L2VpnXcSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcSummary.Size(m)
}
func (m *L2VpnXcSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcSummary proto.InternalMessageInfo

func (m *L2VpnXcSummary) GetNumberGroups() uint32 {
	if m != nil {
		return m.NumberGroups
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnects() uint32 {
	if m != nil {
		return m.NumberXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsUp() uint32 {
	if m != nil {
		return m.NumberXconnectsUp
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsDown() uint32 {
	if m != nil {
		return m.NumberXconnectsDown
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsUnresolved() uint32 {
	if m != nil {
		return m.NumberXconnectsUnresolved
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsAdministrativelyDown() uint32 {
	if m != nil {
		return m.NumberXconnectsAdministrativelyDown
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsTypeAttachmentCircuitToPseudoWire() uint32 {
	if m != nil {
		return m.NumberXconnectsTypeAttachmentCircuitToPseudoWire
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsTypeMonitorSessionToPseudoWire() uint32 {
	if m != nil {
		return m.NumberXconnectsTypeMonitorSessionToPseudoWire
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsLocallySwitched() uint32 {
	if m != nil {
		return m.NumberXconnectsLocallySwitched
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsWithBackupPw() uint32 {
	if m != nil {
		return m.NumberXconnectsWithBackupPw
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsCurrentlyUsingBackup() uint32 {
	if m != nil {
		return m.NumberXconnectsCurrentlyUsingBackup
	}
	return 0
}

func (m *L2VpnXcSummary) GetDownBackupXconnects() uint32 {
	if m != nil {
		return m.DownBackupXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetAdminDownBackupXconnects() uint32 {
	if m != nil {
		return m.AdminDownBackupXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetUnresolvedBackupXconnects() uint32 {
	if m != nil {
		return m.UnresolvedBackupXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetStandbyBackupXconnects() uint32 {
	if m != nil {
		return m.StandbyBackupXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetStandbyReadyBackupXconnects() uint32 {
	if m != nil {
		return m.StandbyReadyBackupXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetOutOfMemoryState() uint32 {
	if m != nil {
		return m.OutOfMemoryState
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsTypePseudoWireToPseudoWire() uint32 {
	if m != nil {
		return m.NumberXconnectsTypePseudoWireToPseudoWire
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberMp2MpXconnects() uint32 {
	if m != nil {
		return m.NumberMp2MpXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberMp2MpXconnectsUp() uint32 {
	if m != nil {
		return m.NumberMp2MpXconnectsUp
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberMp2MpXconnectsDown() uint32 {
	if m != nil {
		return m.NumberMp2MpXconnectsDown
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberMp2MpXconnectsAdvertised() uint32 {
	if m != nil {
		return m.NumberMp2MpXconnectsAdvertised
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberMp2MpXconnectssNotAdvertised() uint32 {
	if m != nil {
		return m.NumberMp2MpXconnectssNotAdvertised
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberCe2Ceconnections() uint32 {
	if m != nil {
		return m.NumberCe2Ceconnections
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberCe2CeAdvertized() uint32 {
	if m != nil {
		return m.NumberCe2CeAdvertized
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberCe2CeNotAdvertized() uint32 {
	if m != nil {
		return m.NumberCe2CeNotAdvertized
	}
	return 0
}

func (m *L2VpnXcSummary) GetPartiallyProgrammedXconnects() uint32 {
	if m != nil {
		return m.PartiallyProgrammedXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsWithBackupInterface() uint32 {
	if m != nil {
		return m.NumberXconnectsWithBackupInterface
	}
	return 0
}

func (m *L2VpnXcSummary) GetNumberXconnectsCurrentlyUsingBackupInterface() uint32 {
	if m != nil {
		return m.NumberXconnectsCurrentlyUsingBackupInterface
	}
	return 0
}

func (m *L2VpnXcSummary) GetDownBackupInterfaceXconnects() uint32 {
	if m != nil {
		return m.DownBackupInterfaceXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetAdminDownBackupInterfaceXconnects() uint32 {
	if m != nil {
		return m.AdminDownBackupInterfaceXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetUnresolvedBackupInterfaceXconnects() uint32 {
	if m != nil {
		return m.UnresolvedBackupInterfaceXconnects
	}
	return 0
}

func (m *L2VpnXcSummary) GetStandbyBackupInterfaceXconnects() uint32 {
	if m != nil {
		return m.StandbyBackupInterfaceXconnects
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnXcSummary_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.xconnect_summary.l2vpn_xc_summary_KEYS")
	proto.RegisterType((*L2VpnXcSummary)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.xconnect_summary.l2vpn_xc_summary")
}

func init() { proto.RegisterFile("l2vpn_xc_summary.proto", fileDescriptor_922b751c698ee021) }

var fileDescriptor_922b751c698ee021 = []byte{
	// 845 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xef, 0x6f, 0x1b, 0x35,
	0x1c, 0xc6, 0xc5, 0x1b, 0x5e, 0x58, 0x4c, 0x94, 0x2b, 0xed, 0x3c, 0x15, 0xb5, 0xac, 0xd3, 0x80,
	0x49, 0x2c, 0x1a, 0x59, 0xd9, 0x3a, 0x60, 0x85, 0x2c, 0x2d, 0xa3, 0x74, 0xd9, 0xb2, 0xa4, 0xd5,
	0xe0, 0x95, 0x75, 0xf1, 0xb9, 0x89, 0xb5, 0x9c, 0x7d, 0xd8, 0xbe, 0xa4, 0xb7, 0xbf, 0x87, 0x3f,
	0x14, 0xdd, 0xd7, 0xf7, 0xc3, 0xf1, 0x19, 0x69, 0xef, 0xaa, 0xf3, 0xf3, 0x7c, 0xbe, 0xce, 0xf7,
	0x9e, 0x3c, 0x0d, 0xda, 0x5d, 0xf6, 0x57, 0x99, 0x20, 0x37, 0x94, 0xe8, 0x3c, 0x4d, 0x63, 0x55,
	0xf4, 0x32, 0x25, 0x8d, 0x8c, 0x7a, 0x94, 0x6b, 0x2a, 0x09, 0x97, 0x9a, 0xdc, 0x28, 0x62, 0x45,
	0x32, 0x63, 0xaa, 0x07, 0x7f, 0xf6, 0x6e, 0xa8, 0x14, 0x82, 0x51, 0x53, 0xbb, 0x0e, 0x6f, 0xa3,
	0x1d, 0x9f, 0x44, 0x2e, 0xce, 0xfe, 0x9e, 0x1e, 0xfe, 0x1b, 0xa1, 0x2d, 0xff, 0x24, 0xba, 0x87,
	0x6e, 0x89, 0x3c, 0x9d, 0x31, 0x45, 0xe6, 0x4a, 0xe6, 0x99, 0xc6, 0xfd, 0xaf, 0x3f, 0xf9, 0xee,
	0xd6, 0xe4, 0x33, 0xfb, 0xf0, 0x25, 0x3c, 0x8b, 0x1e, 0xa0, 0xad, 0x4a, 0x54, 0x4f, 0xd3, 0xf8,
	0x31, 0xe8, 0x3e, 0xb7, 0xcf, 0xff, 0xaa, 0x1f, 0x47, 0x3d, 0xb4, 0xed, 0x4b, 0x49, 0x9e, 0xe1,
	0x23, 0x50, 0x7f, 0xe1, 0xa9, 0xaf, 0xb2, 0xa8, 0x8f, 0x76, 0x3a, 0xfa, 0x44, 0xae, 0x05, 0xfe,
	0x11, 0x1c, 0xdb, 0x9e, 0xe3, 0x54, 0xae, 0x45, 0x74, 0x82, 0xf6, 0xba, 0x33, 0x84, 0x62, 0x5a,
	0x2e, 0x57, 0x2c, 0xc1, 0x4f, 0xc0, 0x79, 0xc7, 0x9f, 0xd5, 0x08, 0xa2, 0x29, 0xfa, 0xa6, 0xe3,
	0x8f, 0x93, 0x94, 0x0b, 0xae, 0x8d, 0x8a, 0x0d, 0x5f, 0xb1, 0x65, 0x61, 0x2f, 0xf1, 0x14, 0x50,
	0xf7, 0x3c, 0xd4, 0xc0, 0xd3, 0xc2, 0xa5, 0xfe, 0x41, 0x4f, 0x3b, 0x50, 0x53, 0x64, 0x8c, 0xc4,
	0xc6, 0xc4, 0x74, 0x91, 0x32, 0x61, 0x08, 0xe5, 0x8a, 0xe6, 0xdc, 0x10, 0x23, 0x49, 0xa6, 0x59,
	0x9e, 0x48, 0xb2, 0xe6, 0x8a, 0xe1, 0x63, 0x98, 0xf2, 0xc8, 0x9b, 0x72, 0x59, 0x64, 0x6c, 0xd0,
	0x78, 0x87, 0xd6, 0x7a, 0x29, 0xc7, 0x60, 0x7c, 0xc7, 0x15, 0x8b, 0xde, 0xa3, 0xa3, 0xf0, 0xc8,
	0x54, 0x0a, 0x6e, 0xa4, 0x22, 0x9a, 0x69, 0xcd, 0xa5, 0xf0, 0xe7, 0x3d, 0x83, 0x79, 0x0f, 0x03,
	0xf3, 0x46, 0xd6, 0x38, 0xb5, 0xbe, 0x8d, 0x61, 0xe7, 0xe8, 0x6e, 0x67, 0xd8, 0x52, 0xd2, 0x78,
	0xb9, 0x2c, 0x88, 0x5e, 0x73, 0x43, 0x17, 0x2c, 0xc1, 0x3f, 0x01, 0x79, 0xdf, 0x23, 0xbf, 0xb2,
	0xb2, 0x69, 0xa5, 0x8a, 0x4e, 0xd1, 0x41, 0x07, 0xb5, 0xe6, 0x66, 0x41, 0x66, 0x31, 0x7d, 0x9f,
	0x67, 0x24, 0x5b, 0xe3, 0x9f, 0x01, 0xb4, 0xe7, 0x81, 0xde, 0x71, 0xb3, 0x78, 0x01, 0x9a, 0xf1,
	0x3a, 0xba, 0x44, 0xdf, 0x76, 0x28, 0x34, 0x57, 0x8a, 0x09, 0xb3, 0x2c, 0x48, 0xae, 0xb9, 0x98,
	0x57, 0x40, 0xfc, 0x4b, 0xf0, 0x35, 0x0e, 0x6b, 0xf1, 0x55, 0xa9, 0xb5, 0xdc, 0x32, 0x8f, 0xe5,
	0x9b, 0xaf, 0xaf, 0xd2, 0xe6, 0xfd, 0xb9, 0xcd, 0x63, 0x79, 0x68, 0xa5, 0x6d, 0xe6, 0x9f, 0xa3,
	0x3d, 0x88, 0x0f, 0x09, 0x3b, 0x4f, 0xc0, 0x89, 0x41, 0x72, 0x1a, 0xb0, 0x9f, 0xa0, 0xbd, 0x36,
	0xbd, 0x5d, 0xfb, 0xaf, 0x36, 0xce, 0xad, 0xc4, 0xf7, 0x1f, 0x23, 0xac, 0x4d, 0x2c, 0x92, 0x59,
	0xd1, 0x35, 0xff, 0x06, 0xe6, 0xdd, 0xea, 0xdc, 0x77, 0x0e, 0xd1, 0x7e, 0xed, 0x54, 0x2c, 0x4e,
	0x02, 0xfe, 0x81, 0x7d, 0x0f, 0x95, 0x6a, 0x52, 0x8a, 0x7c, 0xc8, 0x43, 0xb4, 0x2d, 0x73, 0x43,
	0xe4, 0x35, 0x49, 0x59, 0x2a, 0x55, 0x41, 0xb4, 0x89, 0x0d, 0xc3, 0x2f, 0xc0, 0xb9, 0x25, 0x73,
	0xf3, 0xe6, 0x7a, 0x04, 0x07, 0xd3, 0xf2, 0x79, 0x44, 0xd1, 0xa3, 0x70, 0x68, 0x9d, 0x84, 0xfa,
	0x81, 0x1d, 0x02, 0xeb, 0x41, 0x20, 0xb0, 0x6d, 0x40, 0x37, 0xc2, 0x7a, 0x84, 0x76, 0xab, 0x21,
	0x69, 0xd6, 0x4f, 0xdd, 0x0f, 0x74, 0x0a, 0xa8, 0x2f, 0xed, 0xe9, 0xa8, 0x3c, 0x6c, 0x3f, 0xc9,
	0x33, 0x74, 0x27, 0xec, 0x2a, 0x1b, 0xec, 0xcc, 0x6e, 0x32, 0x64, 0xbc, 0xca, 0xca, 0x08, 0xfc,
	0x8f, 0x15, 0x7a, 0xe4, 0x77, 0x1b, 0x81, 0x90, 0x19, 0xca, 0xa3, 0xfd, 0x72, 0xf9, 0xf6, 0x38,
	0x59, 0x31, 0x65, 0xb8, 0x66, 0x09, 0x7e, 0xe9, 0x7e, 0xb9, 0x36, 0x21, 0x83, 0x46, 0x15, 0x4d,
	0x9a, 0x72, 0xf3, 0x50, 0x9a, 0x08, 0x69, 0x5c, 0xde, 0x1f, 0xc0, 0x3b, 0x0c, 0xf1, 0xf4, 0x6b,
	0x69, 0x1c, 0xe6, 0x31, 0xaa, 0xae, 0x4e, 0x28, 0xeb, 0x13, 0xca, 0x2a, 0x19, 0x97, 0x42, 0xe3,
	0x73, 0x77, 0x2f, 0x43, 0xd6, 0x1f, 0xba, 0xa7, 0xd1, 0x13, 0x74, 0xbb, 0x75, 0x52, 0x56, 0x8f,
	0xff, 0xc0, 0x12, 0xfc, 0x27, 0x18, 0x77, 0x1a, 0x23, 0x65, 0x83, 0xe6, 0xd0, 0xd9, 0xa7, 0xf5,
	0x39, 0x57, 0x2f, 0xbd, 0x17, 0xee, 0x3e, 0xc1, 0xdb, 0x5e, 0xf8, 0x03, 0x34, 0xcc, 0x7e, 0x16,
	0x2b, 0xc3, 0xa1, 0x9d, 0x32, 0x25, 0xe7, 0x2a, 0x4e, 0x53, 0x96, 0x38, 0x39, 0x78, 0x05, 0x84,
	0xaf, 0x1a, 0xd5, 0xb8, 0x11, 0xb5, 0x79, 0x98, 0x04, 0xfe, 0x4f, 0xb8, 0x3d, 0xc5, 0x85, 0x61,
	0xea, 0x3a, 0xa6, 0x0c, 0x8f, 0xdc, 0x55, 0x06, 0xea, 0xea, 0xbc, 0x56, 0x46, 0x73, 0xf4, 0xc3,
	0x47, 0xb6, 0x96, 0x83, 0x7f, 0x0d, 0xf8, 0xef, 0x3f, 0xa2, 0xbf, 0xda, 0x41, 0x67, 0xe8, 0xc0,
	0xad, 0xa3, 0x06, 0xe6, 0xec, 0xe0, 0x8d, 0xdd, 0x41, 0x5b, 0x69, 0x8d, 0xbb, 0xdd, 0xc1, 0x18,
	0xdd, 0xef, 0x76, 0x5b, 0x08, 0x36, 0x06, 0xd8, 0x5d, 0xaf, 0xe5, 0x02, 0xc4, 0xb7, 0xe8, 0x7e,
	0xb7, 0xee, 0x42, 0xc4, 0xb7, 0x76, 0xa9, 0x7e, 0xf1, 0x05, 0x90, 0x17, 0xe8, 0xd0, 0x6b, 0xc0,
	0x10, 0x6f, 0x02, 0xbc, 0x83, 0x8d, 0x2e, 0xec, 0xc2, 0x66, 0x9f, 0xc2, 0xcf, 0xae, 0xc7, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xfc, 0x99, 0x8c, 0x90, 0x09, 0x00, 0x00,
}
