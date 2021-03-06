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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_xconnect_summary

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
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
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

func (m *L2VpnXcSummary_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

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
	proto.RegisterType((*L2VpnXcSummary_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.xconnect_summary.l2vpn_xc_summary_KEYS")
	proto.RegisterType((*L2VpnXcSummary)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.xconnect_summary.l2vpn_xc_summary")
}

func init() { proto.RegisterFile("l2vpn_xc_summary.proto", fileDescriptor_922b751c698ee021) }

var fileDescriptor_922b751c698ee021 = []byte{
	// 872 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xef, 0x4f, 0x1b, 0x37,
	0x1c, 0xc6, 0xc5, 0x9b, 0x4e, 0xb3, 0x56, 0x8d, 0x99, 0x41, 0x5d, 0x31, 0x41, 0x4b, 0xd5, 0xad,
	0x95, 0xd6, 0x88, 0xa5, 0xac, 0xa5, 0xeb, 0xca, 0x96, 0x06, 0xd6, 0x31, 0x4a, 0x9b, 0x26, 0xa0,
	0x6e, 0xaf, 0xac, 0xc3, 0x67, 0xc0, 0x6a, 0xce, 0xbe, 0xd9, 0xbe, 0x84, 0xf4, 0xef, 0xd9, 0x1f,
	0x3a, 0xdd, 0xd7, 0xf7, 0xc3, 0xf8, 0x5c, 0xa9, 0x6f, 0x10, 0xd8, 0xcf, 0xf3, 0xf9, 0x3a, 0xbe,
	0xe7, 0x1e, 0x82, 0xd6, 0xa6, 0xfd, 0x59, 0x2e, 0xe9, 0x15, 0xa3, 0xa6, 0xc8, 0xb2, 0x44, 0x2f,
	0x7a, 0xb9, 0x56, 0x56, 0xe1, 0xe7, 0x4c, 0x18, 0xa6, 0xa8, 0x50, 0x86, 0x5e, 0x69, 0xea, 0x44,
	0x2a, 0xe7, 0xba, 0x07, 0xbf, 0xce, 0xfa, 0x3d, 0xa9, 0x52, 0x6e, 0xe0, 0x67, 0xef, 0x8a, 0x29,
	0x29, 0x39, 0xb3, 0x35, 0x62, 0x6b, 0x1b, 0xad, 0x86, 0x58, 0x7a, 0x74, 0xf0, 0xcf, 0x04, 0xdf,
	0x42, 0x5f, 0x94, 0x0e, 0x2a, 0x52, 0xb2, 0x74, 0x67, 0xe9, 0xc1, 0x97, 0xe3, 0x1b, 0xe5, 0x9f,
	0x87, 0xe9, 0xd6, 0x7f, 0x18, 0x2d, 0x87, 0x16, 0x7c, 0x0f, 0xdd, 0x94, 0x45, 0x76, 0xc6, 0x35,
	0xbd, 0xd0, 0xaa, 0xc8, 0x0d, 0xe9, 0xdf, 0x59, 0x7a, 0x70, 0x73, 0xfc, 0x95, 0x5b, 0x7c, 0x05,
	0x6b, 0xf8, 0x21, 0x5a, 0xae, 0x44, 0xf5, 0x31, 0x0c, 0x79, 0x0c, 0xba, 0xaf, 0xdd, 0xfa, 0xdf,
	0xf5, 0x32, 0xee, 0xa1, 0x95, 0x50, 0x4a, 0x8b, 0x9c, 0xec, 0x80, 0xfa, 0x9b, 0x40, 0x7d, 0x9a,
	0xe3, 0x3e, 0x5a, 0xed, 0xe8, 0x53, 0x35, 0x97, 0xe4, 0x67, 0x70, 0xac, 0x04, 0x8e, 0x7d, 0x35,
	0x97, 0x78, 0x0f, 0xad, 0x77, 0x67, 0x48, 0xcd, 0x8d, 0x9a, 0xce, 0x78, 0x4a, 0x9e, 0x80, 0xf3,
	0x76, 0x38, 0xab, 0x11, 0xe0, 0x09, 0xfa, 0xbe, 0xe3, 0x4f, 0xd2, 0x4c, 0x48, 0x61, 0xac, 0x4e,
	0xac, 0x98, 0xf1, 0xe9, 0xc2, 0x1d, 0xe2, 0x29, 0xa0, 0xee, 0x05, 0xa8, 0x41, 0xa0, 0x85, 0x43,
	0xfd, 0x8b, 0x9e, 0x76, 0xa0, 0x76, 0x91, 0x73, 0x9a, 0x58, 0x9b, 0xb0, 0xcb, 0x8c, 0x4b, 0x4b,
	0x99, 0xd0, 0xac, 0x10, 0x96, 0x5a, 0x45, 0x73, 0xc3, 0x8b, 0x54, 0xd1, 0xb9, 0xd0, 0x9c, 0xec,
	0xc2, 0x94, 0xed, 0x60, 0xca, 0xc9, 0x22, 0xe7, 0x83, 0xc6, 0x3b, 0x74, 0xd6, 0x13, 0x35, 0x02,
	0xe3, 0x7b, 0xa1, 0x39, 0xfe, 0x80, 0x76, 0xe2, 0x23, 0x33, 0x25, 0x85, 0x55, 0x9a, 0x1a, 0x6e,
	0x8c, 0x50, 0x32, 0x9c, 0xf7, 0x0c, 0xe6, 0x3d, 0x8a, 0xcc, 0x3b, 0x76, 0xc6, 0x89, 0xf3, 0x5d,
	0x1b, 0x76, 0x88, 0xee, 0x76, 0x86, 0x4d, 0x15, 0x4b, 0xa6, 0xd3, 0x05, 0x35, 0x73, 0x61, 0xd9,
	0x25, 0x4f, 0xc9, 0x2f, 0x40, 0xde, 0x08, 0xc8, 0xaf, 0x9d, 0x6c, 0x52, 0xa9, 0xf0, 0x3e, 0xda,
	0xec, 0xa0, 0xe6, 0xc2, 0x5e, 0xd2, 0xb3, 0x84, 0x7d, 0x28, 0x72, 0x9a, 0xcf, 0xc9, 0x73, 0x00,
	0xad, 0x07, 0xa0, 0xf7, 0xc2, 0x5e, 0xbe, 0x04, 0xcd, 0x68, 0x8e, 0x4f, 0xd0, 0x0f, 0x1d, 0x0a,
	0x2b, 0xb4, 0xe6, 0xd2, 0x4e, 0x17, 0xb4, 0x30, 0x42, 0x5e, 0x54, 0x40, 0xf2, 0x6b, 0xf4, 0x31,
	0x0e, 0x6b, 0xf1, 0x69, 0xa9, 0x75, 0xdc, 0x32, 0x8f, 0xe5, 0x93, 0xaf, 0x8f, 0xd2, 0xe6, 0xfd,
	0x85, 0xcb, 0x63, 0xb9, 0xe9, 0xa4, 0x6d, 0xe6, 0x5f, 0xa0, 0x75, 0x88, 0x0f, 0x8d, 0x3b, 0xf7,
	0xc0, 0x49, 0x40, 0xb2, 0x1f, 0xb1, 0xef, 0xa1, 0xf5, 0x36, 0xbd, 0x5d, 0xfb, 0x6f, 0x2e, 0xce,
	0xad, 0x24, 0xf4, 0xef, 0x22, 0x62, 0x6c, 0x22, 0xd3, 0xb3, 0x45, 0xd7, 0xfc, 0x3b, 0x98, 0xd7,
	0xaa, 0xfd, 0xd0, 0x39, 0x44, 0x1b, 0xb5, 0x53, 0xf3, 0x24, 0x8d, 0xf8, 0x07, 0xee, 0x39, 0x54,
	0xaa, 0x71, 0x29, 0x0a, 0x21, 0x8f, 0xd0, 0x8a, 0x2a, 0x2c, 0x55, 0xe7, 0x34, 0xe3, 0x99, 0xd2,
	0x0b, 0x6a, 0x6c, 0x62, 0x39, 0x79, 0x09, 0xce, 0x65, 0x55, 0xd8, 0xb7, 0xe7, 0xc7, 0xb0, 0x31,
	0x29, 0xd7, 0x31, 0x43, 0xdb, 0xf1, 0xd0, 0x7a, 0x09, 0x0d, 0x03, 0x3b, 0x04, 0xd6, 0xc3, 0x48,
	0x60, 0xdb, 0x80, 0x5e, 0x0b, 0xeb, 0x0e, 0x5a, 0xab, 0x86, 0x64, 0x79, 0x3f, 0xf3, 0x3f, 0xd0,
	0x3e, 0xa0, 0xbe, 0x75, 0xbb, 0xc7, 0xe5, 0x66, 0xfb, 0x49, 0x9e, 0xa1, 0xdb, 0x71, 0x57, 0xd9,
	0x60, 0x07, 0xee, 0x26, 0x63, 0xc6, 0xd3, 0xbc, 0x8c, 0xc0, 0x27, 0xac, 0xd0, 0x23, 0x7f, 0xb8,
	0x08, 0xc4, 0xcc, 0x50, 0x1e, 0xed, 0xcb, 0x15, 0xda, 0x93, 0x74, 0xc6, 0xb5, 0x15, 0x86, 0xa7,
	0xe4, 0x95, 0xff, 0x72, 0x5d, 0x87, 0x0c, 0x1a, 0x15, 0x1e, 0x37, 0xe5, 0x16, 0xa0, 0x0c, 0x95,
	0xca, 0xfa, 0xbc, 0x3f, 0x81, 0xb7, 0x15, 0xe3, 0x99, 0x37, 0xca, 0x7a, 0xcc, 0x5d, 0x54, 0x1d,
	0x9d, 0x32, 0xde, 0xa7, 0x8c, 0x57, 0x32, 0xa1, 0xa4, 0x21, 0x87, 0xfe, 0xbd, 0x0c, 0x79, 0x7f,
	0xe8, 0xef, 0xe2, 0x27, 0xe8, 0x56, 0xeb, 0x64, 0xbc, 0x1e, 0xff, 0x91, 0xa7, 0xe4, 0x2f, 0x30,
	0xae, 0x36, 0x46, 0xc6, 0x07, 0xcd, 0xa6, 0x77, 0x9f, 0xce, 0xe7, 0x1d, 0xbd, 0xf4, 0x1e, 0xf9,
	0xf7, 0x09, 0xde, 0xf6, 0xc0, 0x1f, 0xa1, 0x61, 0x36, 0xf2, 0x44, 0x5b, 0x01, 0xed, 0x94, 0x6b,
	0x75, 0xa1, 0x93, 0x2c, 0xe3, 0xa9, 0x97, 0x83, 0xd7, 0x40, 0xf8, 0xae, 0x51, 0x8d, 0x1a, 0x51,
	0x9b, 0x87, 0x71, 0xe4, 0xff, 0x84, 0xdf, 0x53, 0x42, 0x5a, 0xae, 0xcf, 0x13, 0xc6, 0xc9, 0xb1,
	0x7f, 0x95, 0x91, 0xba, 0x3a, 0xac, 0x95, 0xf8, 0x02, 0xfd, 0xf4, 0x99, 0xad, 0xe5, 0xe1, 0xdf,
	0x00, 0xfe, 0xc7, 0xcf, 0xe8, 0xaf, 0x76, 0xd0, 0x01, 0xda, 0xf4, 0xeb, 0xa8, 0x81, 0x79, 0x77,
	0xf0, 0xd6, 0xdd, 0x41, 0x5b, 0x69, 0x8d, 0xbb, 0xbd, 0x83, 0x11, 0xba, 0xdf, 0xed, 0xb6, 0x18,
	0x6c, 0x04, 0xb0, 0xbb, 0x41, 0xcb, 0x45, 0x88, 0xef, 0xd0, 0xfd, 0x6e, 0xdd, 0xc5, 0x88, 0xef,
	0xdc, 0xa5, 0x86, 0xc5, 0x17, 0x41, 0x1e, 0xa1, 0xad, 0xa0, 0x01, 0x63, 0xbc, 0x31, 0xf0, 0x36,
	0xaf, 0x75, 0x61, 0x17, 0x76, 0x76, 0x03, 0xbe, 0x9c, 0x3d, 0xfe, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0x0d, 0x3a, 0x2f, 0xb3, 0xb6, 0x09, 0x00, 0x00,
}
