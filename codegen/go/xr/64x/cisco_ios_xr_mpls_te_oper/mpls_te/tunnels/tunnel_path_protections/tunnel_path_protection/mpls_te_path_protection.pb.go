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
// source: mpls_te_path_protection.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_tunnels_tunnel_path_protections_tunnel_path_protection

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

type MplsTePathProtection_KEYS struct {
	Ctype                string   `protobuf:"bytes,1,opt,name=ctype,proto3" json:"ctype,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	TunnelName           string   `protobuf:"bytes,3,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTePathProtection_KEYS) Reset()         { *m = MplsTePathProtection_KEYS{} }
func (m *MplsTePathProtection_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtection_KEYS) ProtoMessage()    {}
func (*MplsTePathProtection_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{0}
}

func (m *MplsTePathProtection_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Unmarshal(m, b)
}
func (m *MplsTePathProtection_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTePathProtection_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtection_KEYS.Merge(m, src)
}
func (m *MplsTePathProtection_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtection_KEYS.Size(m)
}
func (m *MplsTePathProtection_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtection_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtection_KEYS proto.InternalMessageInfo

func (m *MplsTePathProtection_KEYS) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *MplsTePathProtection_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTePathProtection_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

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
	return fileDescriptor_baf87df8ae158ee7, []int{1}
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

type MplsTePpInfo struct {
	TimeOfSwitchoverSec                 uint32             `protobuf:"varint,1,opt,name=time_of_switchover_sec,json=timeOfSwitchoverSec,proto3" json:"time_of_switchover_sec,omitempty"`
	SwitchoverTotal                     uint32             `protobuf:"varint,2,opt,name=switchover_total,json=switchoverTotal,proto3" json:"switchover_total,omitempty"`
	SwitchoverReady                     uint32             `protobuf:"varint,3,opt,name=switchover_ready,json=switchoverReady,proto3" json:"switchover_ready,omitempty"`
	StandbyReoptimizedNumber            uint32             `protobuf:"varint,4,opt,name=standby_reoptimized_number,json=standbyReoptimizedNumber,proto3" json:"standby_reoptimized_number,omitempty"`
	SwitchoverReason                    uint32             `protobuf:"varint,5,opt,name=switchover_reason,json=switchoverReason,proto3" json:"switchover_reason,omitempty"`
	DiversityType                       string             `protobuf:"bytes,6,opt,name=diversity_type,json=diversityType,proto3" json:"diversity_type,omitempty"`
	IsPathProtectConfigured             bool               `protobuf:"varint,7,opt,name=is_path_protect_configured,json=isPathProtectConfigured,proto3" json:"is_path_protect_configured,omitempty"`
	PathProtectionProtectedById         uint32             `protobuf:"varint,8,opt,name=path_protection_protected_by_id,json=pathProtectionProtectedById,proto3" json:"path_protection_protected_by_id,omitempty"`
	ValidPathProtectionPathOptionExists bool               `protobuf:"varint,9,opt,name=valid_path_protection_path_option_exists,json=validPathProtectionPathOptionExists,proto3" json:"valid_path_protection_path_option_exists,omitempty"`
	IsPathProtectSwitchOverUnderway     bool               `protobuf:"varint,10,opt,name=is_path_protect_switch_over_underway,json=isPathProtectSwitchOverUnderway,proto3" json:"is_path_protect_switch_over_underway,omitempty"`
	Switchover                          *TePpSwLogEntryBag `protobuf:"bytes,11,opt,name=switchover,proto3" json:"switchover,omitempty"`
	ReoptimizationTimeRemaining         uint32             `protobuf:"varint,12,opt,name=reoptimization_time_remaining,json=reoptimizationTimeRemaining,proto3" json:"reoptimization_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}           `json:"-"`
	XXX_unrecognized                    []byte             `json:"-"`
	XXX_sizecache                       int32              `json:"-"`
}

func (m *MplsTePpInfo) Reset()         { *m = MplsTePpInfo{} }
func (m *MplsTePpInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTePpInfo) ProtoMessage()    {}
func (*MplsTePpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{2}
}

func (m *MplsTePpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePpInfo.Unmarshal(m, b)
}
func (m *MplsTePpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePpInfo.Marshal(b, m, deterministic)
}
func (m *MplsTePpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePpInfo.Merge(m, src)
}
func (m *MplsTePpInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTePpInfo.Size(m)
}
func (m *MplsTePpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePpInfo proto.InternalMessageInfo

func (m *MplsTePpInfo) GetTimeOfSwitchoverSec() uint32 {
	if m != nil {
		return m.TimeOfSwitchoverSec
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverTotal() uint32 {
	if m != nil {
		return m.SwitchoverTotal
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverReady() uint32 {
	if m != nil {
		return m.SwitchoverReady
	}
	return 0
}

func (m *MplsTePpInfo) GetStandbyReoptimizedNumber() uint32 {
	if m != nil {
		return m.StandbyReoptimizedNumber
	}
	return 0
}

func (m *MplsTePpInfo) GetSwitchoverReason() uint32 {
	if m != nil {
		return m.SwitchoverReason
	}
	return 0
}

func (m *MplsTePpInfo) GetDiversityType() string {
	if m != nil {
		return m.DiversityType
	}
	return ""
}

func (m *MplsTePpInfo) GetIsPathProtectConfigured() bool {
	if m != nil {
		return m.IsPathProtectConfigured
	}
	return false
}

func (m *MplsTePpInfo) GetPathProtectionProtectedById() uint32 {
	if m != nil {
		return m.PathProtectionProtectedById
	}
	return 0
}

func (m *MplsTePpInfo) GetValidPathProtectionPathOptionExists() bool {
	if m != nil {
		return m.ValidPathProtectionPathOptionExists
	}
	return false
}

func (m *MplsTePpInfo) GetIsPathProtectSwitchOverUnderway() bool {
	if m != nil {
		return m.IsPathProtectSwitchOverUnderway
	}
	return false
}

func (m *MplsTePpInfo) GetSwitchover() *TePpSwLogEntryBag {
	if m != nil {
		return m.Switchover
	}
	return nil
}

func (m *MplsTePpInfo) GetReoptimizationTimeRemaining() uint32 {
	if m != nil {
		return m.ReoptimizationTimeRemaining
	}
	return 0
}

type MplsTeS2LSrOutgoingFwdInfo struct {
	LspOutputInterface   string   `protobuf:"bytes,1,opt,name=lsp_output_interface,json=lspOutputInterface,proto3" json:"lsp_output_interface,omitempty"`
	LspOutputLabel       uint32   `protobuf:"varint,2,opt,name=lsp_output_label,json=lspOutputLabel,proto3" json:"lsp_output_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeS2LSrOutgoingFwdInfo) Reset()         { *m = MplsTeS2LSrOutgoingFwdInfo{} }
func (m *MplsTeS2LSrOutgoingFwdInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeS2LSrOutgoingFwdInfo) ProtoMessage()    {}
func (*MplsTeS2LSrOutgoingFwdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{3}
}

func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Unmarshal(m, b)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Merge(m, src)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.Size(m)
}
func (m *MplsTeS2LSrOutgoingFwdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeS2LSrOutgoingFwdInfo proto.InternalMessageInfo

func (m *MplsTeS2LSrOutgoingFwdInfo) GetLspOutputInterface() string {
	if m != nil {
		return m.LspOutputInterface
	}
	return ""
}

func (m *MplsTeS2LSrOutgoingFwdInfo) GetLspOutputLabel() uint32 {
	if m != nil {
		return m.LspOutputLabel
	}
	return 0
}

type RsvpMgmtEroIpv4Subobj struct {
	IsStrictRoute        bool     `protobuf:"varint,1,opt,name=is_strict_route,json=isStrictRoute,proto3" json:"is_strict_route,omitempty"`
	EroAddress           string   `protobuf:"bytes,2,opt,name=ero_address,json=eroAddress,proto3" json:"ero_address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEroIpv4Subobj) Reset()         { *m = RsvpMgmtEroIpv4Subobj{} }
func (m *RsvpMgmtEroIpv4Subobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroIpv4Subobj) ProtoMessage()    {}
func (*RsvpMgmtEroIpv4Subobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{4}
}

func (m *RsvpMgmtEroIpv4Subobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Merge(m, src)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroIpv4Subobj.Size(m)
}
func (m *RsvpMgmtEroIpv4Subobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroIpv4Subobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroIpv4Subobj proto.InternalMessageInfo

func (m *RsvpMgmtEroIpv4Subobj) GetIsStrictRoute() bool {
	if m != nil {
		return m.IsStrictRoute
	}
	return false
}

func (m *RsvpMgmtEroIpv4Subobj) GetEroAddress() string {
	if m != nil {
		return m.EroAddress
	}
	return ""
}

func (m *RsvpMgmtEroIpv4Subobj) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type RsvpMgmtEroUnnumSubobj struct {
	IsStrictRoute        bool     `protobuf:"varint,1,opt,name=is_strict_route,json=isStrictRoute,proto3" json:"is_strict_route,omitempty"`
	EroInterfaceId       uint32   `protobuf:"varint,2,opt,name=ero_interface_id,json=eroInterfaceId,proto3" json:"ero_interface_id,omitempty"`
	EroRouterId          string   `protobuf:"bytes,3,opt,name=ero_router_id,json=eroRouterId,proto3" json:"ero_router_id,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEroUnnumSubobj) Reset()         { *m = RsvpMgmtEroUnnumSubobj{} }
func (m *RsvpMgmtEroUnnumSubobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroUnnumSubobj) ProtoMessage()    {}
func (*RsvpMgmtEroUnnumSubobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{5}
}

func (m *RsvpMgmtEroUnnumSubobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Merge(m, src)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroUnnumSubobj.Size(m)
}
func (m *RsvpMgmtEroUnnumSubobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroUnnumSubobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroUnnumSubobj proto.InternalMessageInfo

func (m *RsvpMgmtEroUnnumSubobj) GetIsStrictRoute() bool {
	if m != nil {
		return m.IsStrictRoute
	}
	return false
}

func (m *RsvpMgmtEroUnnumSubobj) GetEroInterfaceId() uint32 {
	if m != nil {
		return m.EroInterfaceId
	}
	return 0
}

func (m *RsvpMgmtEroUnnumSubobj) GetEroRouterId() string {
	if m != nil {
		return m.EroRouterId
	}
	return ""
}

func (m *RsvpMgmtEroUnnumSubobj) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type RsvpMgmtEroSubobj struct {
	EroType                string                  `protobuf:"bytes,1,opt,name=ero_type,json=eroType,proto3" json:"ero_type,omitempty"`
	Ipv4EroSubObject       *RsvpMgmtEroIpv4Subobj  `protobuf:"bytes,2,opt,name=ipv4ero_sub_object,json=ipv4eroSubObject,proto3" json:"ipv4ero_sub_object,omitempty"`
	UnnumberedEroSubObject *RsvpMgmtEroUnnumSubobj `protobuf:"bytes,3,opt,name=unnumbered_ero_sub_object,json=unnumberedEroSubObject,proto3" json:"unnumbered_ero_sub_object,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *RsvpMgmtEroSubobj) Reset()         { *m = RsvpMgmtEroSubobj{} }
func (m *RsvpMgmtEroSubobj) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEroSubobj) ProtoMessage()    {}
func (*RsvpMgmtEroSubobj) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{6}
}

func (m *RsvpMgmtEroSubobj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Unmarshal(m, b)
}
func (m *RsvpMgmtEroSubobj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtEroSubobj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEroSubobj.Merge(m, src)
}
func (m *RsvpMgmtEroSubobj) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEroSubobj.Size(m)
}
func (m *RsvpMgmtEroSubobj) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEroSubobj.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEroSubobj proto.InternalMessageInfo

func (m *RsvpMgmtEroSubobj) GetEroType() string {
	if m != nil {
		return m.EroType
	}
	return ""
}

func (m *RsvpMgmtEroSubobj) GetIpv4EroSubObject() *RsvpMgmtEroIpv4Subobj {
	if m != nil {
		return m.Ipv4EroSubObject
	}
	return nil
}

func (m *RsvpMgmtEroSubobj) GetUnnumberedEroSubObject() *RsvpMgmtEroUnnumSubobj {
	if m != nil {
		return m.UnnumberedEroSubObject
	}
	return nil
}

type MplsTePathProtectionLsp struct {
	LspUptime            uint32                        `protobuf:"varint,1,opt,name=lsp_uptime,json=lspUptime,proto3" json:"lsp_uptime,omitempty"`
	PathProtectionLspId  uint32                        `protobuf:"varint,2,opt,name=path_protection_lsp_id,json=pathProtectionLspId,proto3" json:"path_protection_lsp_id,omitempty"`
	LspLocalLabel        uint32                        `protobuf:"varint,3,opt,name=lsp_local_label,json=lspLocalLabel,proto3" json:"lsp_local_label,omitempty"`
	SrlspOutgoingInfo    []*MplsTeS2LSrOutgoingFwdInfo `protobuf:"bytes,4,rep,name=srlsp_outgoing_info,json=srlspOutgoingInfo,proto3" json:"srlsp_outgoing_info,omitempty"`
	LspOutputInterface   string                        `protobuf:"bytes,5,opt,name=lsp_output_interface,json=lspOutputInterface,proto3" json:"lsp_output_interface,omitempty"`
	LspOutputLabel       uint32                        `protobuf:"varint,6,opt,name=lsp_output_label,json=lspOutputLabel,proto3" json:"lsp_output_label,omitempty"`
	LspHop               []*RsvpMgmtEroSubobj          `protobuf:"bytes,7,rep,name=lsp_hop,json=lspHop,proto3" json:"lsp_hop,omitempty"`
	LspState             string                        `protobuf:"bytes,8,opt,name=lsp_state,json=lspState,proto3" json:"lsp_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MplsTePathProtectionLsp) Reset()         { *m = MplsTePathProtectionLsp{} }
func (m *MplsTePathProtectionLsp) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtectionLsp) ProtoMessage()    {}
func (*MplsTePathProtectionLsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{7}
}

func (m *MplsTePathProtectionLsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtectionLsp.Unmarshal(m, b)
}
func (m *MplsTePathProtectionLsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtectionLsp.Marshal(b, m, deterministic)
}
func (m *MplsTePathProtectionLsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtectionLsp.Merge(m, src)
}
func (m *MplsTePathProtectionLsp) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtectionLsp.Size(m)
}
func (m *MplsTePathProtectionLsp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtectionLsp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtectionLsp proto.InternalMessageInfo

func (m *MplsTePathProtectionLsp) GetLspUptime() uint32 {
	if m != nil {
		return m.LspUptime
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetPathProtectionLspId() uint32 {
	if m != nil {
		return m.PathProtectionLspId
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetLspLocalLabel() uint32 {
	if m != nil {
		return m.LspLocalLabel
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetSrlspOutgoingInfo() []*MplsTeS2LSrOutgoingFwdInfo {
	if m != nil {
		return m.SrlspOutgoingInfo
	}
	return nil
}

func (m *MplsTePathProtectionLsp) GetLspOutputInterface() string {
	if m != nil {
		return m.LspOutputInterface
	}
	return ""
}

func (m *MplsTePathProtectionLsp) GetLspOutputLabel() uint32 {
	if m != nil {
		return m.LspOutputLabel
	}
	return 0
}

func (m *MplsTePathProtectionLsp) GetLspHop() []*RsvpMgmtEroSubobj {
	if m != nil {
		return m.LspHop
	}
	return nil
}

func (m *MplsTePathProtectionLsp) GetLspState() string {
	if m != nil {
		return m.LspState
	}
	return ""
}

type MplsTePathProtection struct {
	IsTunnelUp           bool                     `protobuf:"varint,50,opt,name=is_tunnel_up,json=isTunnelUp,proto3" json:"is_tunnel_up,omitempty"`
	SourceAddress        string                   `protobuf:"bytes,51,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string                   `protobuf:"bytes,52,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	ExtendedTunnelId     string                   `protobuf:"bytes,53,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	PathProtection       *MplsTePpInfo            `protobuf:"bytes,54,opt,name=path_protection,json=pathProtection,proto3" json:"path_protection,omitempty"`
	CurrentLsp           *MplsTePathProtectionLsp `protobuf:"bytes,55,opt,name=current_lsp,json=currentLsp,proto3" json:"current_lsp,omitempty"`
	StandbyLsp           *MplsTePathProtectionLsp `protobuf:"bytes,56,opt,name=standby_lsp,json=standbyLsp,proto3" json:"standby_lsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MplsTePathProtection) Reset()         { *m = MplsTePathProtection{} }
func (m *MplsTePathProtection) String() string { return proto.CompactTextString(m) }
func (*MplsTePathProtection) ProtoMessage()    {}
func (*MplsTePathProtection) Descriptor() ([]byte, []int) {
	return fileDescriptor_baf87df8ae158ee7, []int{8}
}

func (m *MplsTePathProtection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTePathProtection.Unmarshal(m, b)
}
func (m *MplsTePathProtection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTePathProtection.Marshal(b, m, deterministic)
}
func (m *MplsTePathProtection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTePathProtection.Merge(m, src)
}
func (m *MplsTePathProtection) XXX_Size() int {
	return xxx_messageInfo_MplsTePathProtection.Size(m)
}
func (m *MplsTePathProtection) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTePathProtection.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTePathProtection proto.InternalMessageInfo

func (m *MplsTePathProtection) GetIsTunnelUp() bool {
	if m != nil {
		return m.IsTunnelUp
	}
	return false
}

func (m *MplsTePathProtection) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTePathProtection) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTePathProtection) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTePathProtection) GetPathProtection() *MplsTePpInfo {
	if m != nil {
		return m.PathProtection
	}
	return nil
}

func (m *MplsTePathProtection) GetCurrentLsp() *MplsTePathProtectionLsp {
	if m != nil {
		return m.CurrentLsp
	}
	return nil
}

func (m *MplsTePathProtection) GetStandbyLsp() *MplsTePathProtectionLsp {
	if m != nil {
		return m.StandbyLsp
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTePathProtection_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection_KEYS")
	proto.RegisterType((*TePpSwLogEntryBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.te_pp_sw_log_entry_bag")
	proto.RegisterType((*MplsTePpInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_pp_info")
	proto.RegisterType((*MplsTeS2LSrOutgoingFwdInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_s2l_sr_outgoing_fwd_info")
	proto.RegisterType((*RsvpMgmtEroIpv4Subobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_ipv4_subobj")
	proto.RegisterType((*RsvpMgmtEroUnnumSubobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_unnum_subobj")
	proto.RegisterType((*RsvpMgmtEroSubobj)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.rsvp_mgmt_ero_subobj")
	proto.RegisterType((*MplsTePathProtectionLsp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection_lsp")
	proto.RegisterType((*MplsTePathProtection)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.tunnels.tunnel_path_protections.tunnel_path_protection.mpls_te_path_protection")
}

func init() { proto.RegisterFile("mpls_te_path_protection.proto", fileDescriptor_baf87df8ae158ee7) }

var fileDescriptor_baf87df8ae158ee7 = []byte{
	// 1263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4d, 0x6f, 0x1b, 0x37,
	0x13, 0xc6, 0xbe, 0x8a, 0x65, 0x69, 0x14, 0xd9, 0x0a, 0x6d, 0x38, 0x6b, 0xe7, 0x4b, 0xaf, 0xd2,
	0x06, 0x2a, 0x5a, 0xb8, 0x85, 0x93, 0xf4, 0x03, 0xed, 0xa1, 0x4d, 0xe3, 0xa2, 0x42, 0x95, 0xc8,
	0x58, 0xd9, 0x40, 0x7b, 0x22, 0x56, 0x4b, 0x4a, 0x66, 0xb0, 0xbb, 0x24, 0x48, 0xae, 0x6c, 0xe5,
	0xd0, 0x6b, 0xda, 0x5b, 0xd1, 0x43, 0xef, 0xbd, 0xe6, 0xd0, 0xfe, 0x8e, 0xfe, 0x9a, 0xfe, 0x85,
	0x82, 0x5c, 0xae, 0xbe, 0x62, 0xa5, 0xc8, 0xc1, 0xe9, 0xc9, 0xe2, 0xcc, 0xc3, 0xd9, 0x87, 0x33,
	0xf3, 0x0c, 0x69, 0xb8, 0x95, 0x88, 0x58, 0x61, 0x4d, 0xb1, 0x08, 0xf5, 0x29, 0x16, 0x92, 0x6b,
	0x1a, 0x69, 0xc6, 0xd3, 0x7d, 0xf3, 0x93, 0xa3, 0xef, 0x23, 0xa6, 0x22, 0x8e, 0x19, 0x57, 0xf8,
	0x5c, 0xe2, 0x02, 0xcb, 0x05, 0x95, 0xfb, 0x6e, 0xb1, 0xaf, 0xb3, 0x34, 0xa5, 0xb1, 0x72, 0x7f,
	0x97, 0xe3, 0xac, 0xb2, 0xb7, 0x24, 0xdc, 0x5c, 0xf1, 0x69, 0xfc, 0xdd, 0xe1, 0x0f, 0x7d, 0xb4,
	0x0d, 0x6b, 0x91, 0x9e, 0x08, 0xea, 0x7b, 0x4d, 0xaf, 0x5d, 0x0d, 0xf2, 0x05, 0xba, 0x01, 0x55,
	0x17, 0x8f, 0x11, 0xff, 0x7f, 0x4d, 0xaf, 0x5d, 0x0f, 0x2a, 0xb9, 0xa1, 0x43, 0xd0, 0x1d, 0xa8,
	0x39, 0x67, 0x1a, 0x26, 0xd4, 0x2f, 0xd9, 0x8d, 0x90, 0x9b, 0x9e, 0x86, 0x09, 0x6d, 0xfd, 0x5d,
	0x82, 0x1d, 0xf3, 0x3d, 0x81, 0xd5, 0x19, 0x8e, 0xf9, 0x08, 0xd3, 0x54, 0xcb, 0x09, 0x1e, 0x84,
	0x23, 0x14, 0xc0, 0xbd, 0x65, 0x1a, 0xea, 0x8c, 0xe9, 0xe8, 0x94, 0x8f, 0xa9, 0xc4, 0x74, 0x4c,
	0x53, 0x8d, 0x59, 0x4a, 0xe8, 0xb9, 0xe5, 0x53, 0x0f, 0x5a, 0x06, 0x7d, 0x34, 0x05, 0xf7, 0xa7,
	0xd8, 0x43, 0x03, 0xed, 0x18, 0x24, 0xfa, 0x0c, 0x76, 0x97, 0x63, 0x2e, 0x93, 0xdf, 0x59, 0x0c,
	0x73, 0x5c, 0x1c, 0xe5, 0x36, 0xd4, 0x86, 0x92, 0x27, 0x38, 0x56, 0xc2, 0x80, 0x4b, 0x16, 0x5c,
	0x35, 0xa6, 0xae, 0x12, 0x1d, 0x82, 0xf6, 0xa0, 0xaa, 0x79, 0xe1, 0xbd, 0x62, 0xbd, 0xeb, 0x9a,
	0xe7, 0xbe, 0x87, 0x70, 0x9d, 0x84, 0xa6, 0x4a, 0x43, 0x4c, 0xa5, 0xe4, 0x12, 0x13, 0xea, 0xc2,
	0xfb, 0x6b, 0x16, 0xb9, 0x6d, 0xdc, 0xbd, 0xe1, 0xa1, 0x71, 0x3e, 0x2e, 0x7c, 0xe8, 0x1b, 0x68,
	0xae, 0xd8, 0x86, 0x13, 0x16, 0xc7, 0x4c, 0xd1, 0xc8, 0x2f, 0xdb, 0xfd, 0x37, 0x2f, 0xda, 0xff,
	0xc4, 0x61, 0xd0, 0x97, 0x70, 0x73, 0x2e, 0x73, 0x24, 0x93, 0xe1, 0x62, 0x8c, 0x75, 0x1b, 0x63,
	0x6f, 0x86, 0x79, 0xec, 0x20, 0xd3, 0x08, 0x1d, 0xf8, 0xff, 0x6b, 0x6a, 0x21, 0x69, 0xa8, 0x78,
	0xea, 0x57, 0x6c, 0x75, 0x6f, 0xaf, 0x2a, 0x43, 0x60, 0x51, 0xad, 0x3f, 0xca, 0xb0, 0x39, 0x6d,
	0x33, 0x81, 0x59, 0x3a, 0xe4, 0xe8, 0x3e, 0xec, 0x68, 0x96, 0xd8, 0x83, 0xce, 0x85, 0x35, 0xd4,
	0xf2, 0xd2, 0x6e, 0x19, 0x6f, 0x6f, 0x38, 0x8b, 0xd5, 0xa7, 0x11, 0x7a, 0x0f, 0x1a, 0x73, 0x60,
	0xcd, 0x75, 0x18, 0xbb, 0x12, 0x6e, 0xce, 0xec, 0xc7, 0xc6, 0xbc, 0x04, 0x95, 0x34, 0x24, 0x13,
	0x57, 0xc0, 0x39, 0x68, 0x60, 0xcc, 0xe8, 0x0b, 0xd8, 0x53, 0x3a, 0x4c, 0xc9, 0x60, 0x82, 0x25,
	0xe5, 0x42, 0xb3, 0x84, 0x3d, 0xa7, 0x04, 0xa7, 0x59, 0x32, 0xa0, 0xd2, 0xd5, 0xd5, 0x77, 0x88,
	0x60, 0x06, 0x78, 0x6a, 0xfd, 0xe8, 0x7d, 0xb8, 0xf6, 0x6a, 0x5e, 0xf2, 0x12, 0x37, 0xd4, 0x52,
	0x26, 0xd0, 0xbb, 0xb0, 0x41, 0xd8, 0x98, 0x4a, 0xc5, 0xf4, 0x04, 0x5b, 0x61, 0x95, 0x6d, 0x06,
	0xeb, 0x53, 0xeb, 0xb1, 0x11, 0xd8, 0xe7, 0xb0, 0xc7, 0xd4, 0x82, 0x22, 0x71, 0xc4, 0xd3, 0x21,
	0x1b, 0x65, 0x92, 0x12, 0x5b, 0xbb, 0x4a, 0x70, 0x9d, 0xa9, 0xa3, 0x59, 0xda, 0xbf, 0x9e, 0xba,
	0xd1, 0x63, 0xb8, 0xb3, 0x5c, 0x38, 0xf7, 0x93, 0x12, 0x3c, 0x98, 0x98, 0x5e, 0xad, 0x58, 0x7a,
	0x37, 0x16, 0xcb, 0x76, 0x54, 0x80, 0x1e, 0x4d, 0x3a, 0x04, 0x9d, 0x40, 0x7b, 0x1c, 0xc6, 0x8c,
	0xbc, 0x32, 0x17, 0xec, 0xda, 0x24, 0x81, 0xa7, 0x98, 0x9e, 0x33, 0xa5, 0x95, 0x5f, 0xb5, 0x84,
	0xee, 0x5a, 0xfc, 0xd1, 0x62, 0xcc, 0x50, 0x9f, 0xf6, 0x2c, 0xf6, 0xd0, 0x42, 0xd1, 0x13, 0x78,
	0x67, 0xf9, 0x64, 0x79, 0x92, 0xb0, 0x4d, 0x5f, 0x96, 0x12, 0x2a, 0xcf, 0xc2, 0x89, 0x0f, 0x36,
	0xe4, 0x9d, 0x85, 0x33, 0xe6, 0xbd, 0xd0, 0x1b, 0x53, 0x79, 0xe2, 0x60, 0xe8, 0x17, 0x0f, 0x60,
	0x96, 0x64, 0xbf, 0xd6, 0xf4, 0xda, 0xb5, 0x03, 0xb1, 0x7f, 0x59, 0xf3, 0x72, 0xff, 0xe2, 0xb9,
	0x15, 0xcc, 0x71, 0x40, 0x8f, 0xe0, 0xd6, 0xb4, 0x8b, 0x72, 0xd1, 0xd9, 0x3e, 0x97, 0x34, 0x09,
	0x59, 0xca, 0xd2, 0x91, 0x7f, 0x35, 0x4f, 0xfe, 0x22, 0xe8, 0x98, 0x25, 0x34, 0x28, 0x20, 0xad,
	0x1f, 0xa1, 0x59, 0xb0, 0x56, 0x07, 0x31, 0x56, 0x12, 0xf3, 0x4c, 0x8f, 0x38, 0x4b, 0x47, 0x78,
	0x78, 0x46, 0x72, 0x01, 0x7d, 0x04, 0xdb, 0x66, 0xf2, 0xf0, 0x4c, 0x8b, 0xcc, 0x4c, 0x45, 0x4d,
	0xe5, 0x30, 0x8c, 0x8a, 0x49, 0x8d, 0x62, 0x25, 0x7a, 0xd6, 0xd5, 0x29, 0x3c, 0xa8, 0x0d, 0x8d,
	0xb9, 0x1d, 0x71, 0x38, 0xa0, 0x85, 0x7a, 0x36, 0xa6, 0xe8, 0xae, 0xb1, 0xb6, 0x7e, 0xf6, 0x60,
	0x57, 0xaa, 0xb1, 0xc0, 0xc9, 0x28, 0xd1, 0x98, 0x4a, 0x8e, 0x99, 0x18, 0x3f, 0xc0, 0x2a, 0x1b,
	0xf0, 0xc1, 0x33, 0x74, 0x0f, 0x36, 0x99, 0xc2, 0x4a, 0x4b, 0x16, 0x69, 0x2c, 0x79, 0xa6, 0xf3,
	0x8f, 0x56, 0x82, 0x3a, 0x53, 0x7d, 0x6b, 0x0d, 0x8c, 0xd1, 0xdc, 0x04, 0x66, 0x6b, 0x48, 0x88,
	0xa4, 0x4a, 0xd9, 0x4f, 0x55, 0x03, 0xa0, 0x92, 0x7f, 0x95, 0x5b, 0xd0, 0x5d, 0xa8, 0x0b, 0x49,
	0x87, 0xec, 0x1c, 0xc7, 0x34, 0x1d, 0xe9, 0x53, 0x27, 0xd0, 0xab, 0xb9, 0xb1, 0x6b, 0x6d, 0xad,
	0x97, 0x1e, 0xec, 0x2d, 0x72, 0xc9, 0xd2, 0x34, 0x4b, 0xde, 0x94, 0x4c, 0x1b, 0x1a, 0xf6, 0x1c,
	0x45, 0x36, 0x66, 0xd3, 0x7f, 0x83, 0x4a, 0x3e, 0x4d, 0x52, 0x87, 0xa0, 0x16, 0xd4, 0x0d, 0xd2,
	0xc6, 0x92, 0xc5, 0xdc, 0xaf, 0x06, 0xe6, 0x2c, 0x36, 0x94, 0xec, 0x10, 0xb4, 0x03, 0x65, 0xa5,
	0x43, 0x9d, 0x29, 0x3b, 0x1e, 0xaa, 0x81, 0x5b, 0xb5, 0x5e, 0x94, 0x60, 0x7b, 0x91, 0xac, 0xa3,
	0xb9, 0x0b, 0x15, 0xb3, 0x9a, 0xbb, 0x4b, 0xd7, 0xa9, 0xe4, 0x56, 0xec, 0xbf, 0x7b, 0x80, 0x4c,
	0x7a, 0x1d, 0x1a, 0xf3, 0xc1, 0x33, 0x1a, 0x69, 0x4b, 0xae, 0x76, 0xa0, 0x2e, 0xaf, 0x97, 0x57,
	0x16, 0x38, 0x68, 0x38, 0x3a, 0xfd, 0x6c, 0xd0, 0xb3, 0x64, 0xd0, 0x9f, 0x1e, 0xec, 0xda, 0xb4,
	0x0f, 0xa8, 0xa4, 0x04, 0x2f, 0x51, 0x2d, 0x59, 0xaa, 0xfa, 0x6d, 0x51, 0x9d, 0xaf, 0x7f, 0xb0,
	0x33, 0xa3, 0x75, 0x38, 0xc7, 0xb8, 0xf5, 0xd7, 0x15, 0xb8, 0xb1, 0xea, 0x69, 0x13, 0x2b, 0x81,
	0x6e, 0x01, 0x18, 0x31, 0x64, 0x46, 0x83, 0xd4, 0xdd, 0x39, 0xd5, 0x58, 0x89, 0x13, 0x6b, 0x30,
	0xd7, 0xd3, 0x05, 0xbb, 0x66, 0x4d, 0xb3, 0xb5, 0x38, 0x3b, 0xf3, 0x3b, 0xff, 0x1e, 0x6c, 0x1a,
	0x50, 0xcc, 0xa3, 0x30, 0x76, 0xfa, 0xca, 0x3b, 0xba, 0x1e, 0x2b, 0xd1, 0x35, 0x56, 0x2b, 0x2f,
	0xf4, 0xd2, 0x83, 0x2d, 0x25, 0x9d, 0x16, 0x73, 0x59, 0x1b, 0x49, 0xfb, 0x57, 0x9a, 0xa5, 0x76,
	0xed, 0xe0, 0xf9, 0xe5, 0xe5, 0xf1, 0xdf, 0x86, 0x4a, 0x70, 0xcd, 0xd2, 0xea, 0x39, 0x7b, 0xe7,
	0x75, 0x73, 0x66, 0xed, 0x8d, 0xe6, 0x4c, 0xf9, 0xa2, 0x39, 0x83, 0x5e, 0x78, 0xb0, 0x6e, 0xa0,
	0xa7, 0x5c, 0xf8, 0xeb, 0xf6, 0xf0, 0xe9, 0xdb, 0x6a, 0x22, 0xd7, 0x3e, 0xe5, 0x58, 0x89, 0x6f,
	0xb9, 0x30, 0x4f, 0x5a, 0x43, 0xc4, 0xc8, 0x98, 0xba, 0x57, 0x4d, 0x25, 0x56, 0xa2, 0x6f, 0xd6,
	0xad, 0x9f, 0xd6, 0xe0, 0xfa, 0x8a, 0x5e, 0x42, 0x4d, 0xb8, 0xca, 0x54, 0xf1, 0xa2, 0xcc, 0x84,
	0x7f, 0x60, 0x87, 0x0f, 0x30, 0x95, 0xbf, 0x22, 0x4f, 0x84, 0xb9, 0xf3, 0x15, 0xcf, 0x64, 0x44,
	0xa7, 0x93, 0xf0, 0x7e, 0x7e, 0xe7, 0xe7, 0xd6, 0x62, 0x18, 0x7e, 0x08, 0x5b, 0x84, 0x2a, 0xcd,
	0xd2, 0xfc, 0xd2, 0x28, 0xb0, 0x0f, 0xf2, 0x34, 0xcf, 0xb9, 0x8a, 0x0d, 0x1f, 0x00, 0xa2, 0xe7,
	0x9a, 0xa6, 0x84, 0x92, 0xb9, 0x17, 0xed, 0x43, 0x8b, 0x6f, 0x14, 0x9e, 0xe9, 0x5b, 0xf6, 0x57,
	0x0f, 0x36, 0x97, 0xb8, 0xfb, 0x1f, 0x5b, 0xdd, 0xb2, 0xcb, 0xef, 0x37, 0xf7, 0xe8, 0x0b, 0x36,
	0x16, 0x55, 0x83, 0x7e, 0xf3, 0xa0, 0x16, 0x65, 0x52, 0x9a, 0x77, 0x7d, 0xac, 0x84, 0xff, 0x89,
	0x25, 0x94, 0xbd, 0x05, 0x42, 0xaf, 0x6a, 0x3b, 0x00, 0xc7, 0xa4, 0xab, 0x84, 0x25, 0x56, 0xbc,
	0x09, 0x0d, 0xb1, 0x4f, 0xff, 0x53, 0x62, 0x8e, 0x49, 0x57, 0x89, 0x41, 0xd9, 0xfe, 0x47, 0x78,
	0xff, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0xd5, 0x12, 0x98, 0x32, 0x0e, 0x00, 0x00,
}
