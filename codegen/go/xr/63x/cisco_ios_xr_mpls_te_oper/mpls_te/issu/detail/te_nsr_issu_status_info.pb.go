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
// source: te_nsr_issu_status_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_issu_detail

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

type TeNsrIssuStatusInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeNsrIssuStatusInfo_KEYS) Reset()         { *m = TeNsrIssuStatusInfo_KEYS{} }
func (m *TeNsrIssuStatusInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeNsrIssuStatusInfo_KEYS) ProtoMessage()    {}
func (*TeNsrIssuStatusInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{0}
}

func (m *TeNsrIssuStatusInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Unmarshal(m, b)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Merge(m, src)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.Size(m)
}
func (m *TeNsrIssuStatusInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeNsrIssuStatusInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeNsrIssuStatusInfo_KEYS proto.InternalMessageInfo

type TeIdtStatus struct {
	IsReadyStatus        bool     `protobuf:"varint,1,opt,name=is_ready_status,json=isReadyStatus,proto3" json:"is_ready_status,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	IdtStartTime         uint32   `protobuf:"varint,3,opt,name=idt_start_time,json=idtStartTime,proto3" json:"idt_start_time,omitempty"`
	IdtEndTime           uint32   `protobuf:"varint,4,opt,name=idt_end_time,json=idtEndTime,proto3" json:"idt_end_time,omitempty"`
	DeclareTime          uint32   `protobuf:"varint,5,opt,name=declare_time,json=declareTime,proto3" json:"declare_time,omitempty"`
	WithdrawTime         uint32   `protobuf:"varint,6,opt,name=withdraw_time,json=withdrawTime,proto3" json:"withdraw_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeIdtStatus) Reset()         { *m = TeIdtStatus{} }
func (m *TeIdtStatus) String() string { return proto.CompactTextString(m) }
func (*TeIdtStatus) ProtoMessage()    {}
func (*TeIdtStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{1}
}

func (m *TeIdtStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeIdtStatus.Unmarshal(m, b)
}
func (m *TeIdtStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeIdtStatus.Marshal(b, m, deterministic)
}
func (m *TeIdtStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeIdtStatus.Merge(m, src)
}
func (m *TeIdtStatus) XXX_Size() int {
	return xxx_messageInfo_TeIdtStatus.Size(m)
}
func (m *TeIdtStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TeIdtStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TeIdtStatus proto.InternalMessageInfo

func (m *TeIdtStatus) GetIsReadyStatus() bool {
	if m != nil {
		return m.IsReadyStatus
	}
	return false
}

func (m *TeIdtStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TeIdtStatus) GetIdtStartTime() uint32 {
	if m != nil {
		return m.IdtStartTime
	}
	return 0
}

func (m *TeIdtStatus) GetIdtEndTime() uint32 {
	if m != nil {
		return m.IdtEndTime
	}
	return 0
}

func (m *TeIdtStatus) GetDeclareTime() uint32 {
	if m != nil {
		return m.DeclareTime
	}
	return 0
}

func (m *TeIdtStatus) GetWithdrawTime() uint32 {
	if m != nil {
		return m.WithdrawTime
	}
	return 0
}

type TeSyncIdtInfo struct {
	CurrentIdtInfo       *TeIdtStatus `protobuf:"bytes,1,opt,name=current_idt_info,json=currentIdtInfo,proto3" json:"current_idt_info,omitempty"`
	PreviousIdtStatus    *TeIdtStatus `protobuf:"bytes,2,opt,name=previous_idt_status,json=previousIdtStatus,proto3" json:"previous_idt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TeSyncIdtInfo) Reset()         { *m = TeSyncIdtInfo{} }
func (m *TeSyncIdtInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncIdtInfo) ProtoMessage()    {}
func (*TeSyncIdtInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{2}
}

func (m *TeSyncIdtInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncIdtInfo.Unmarshal(m, b)
}
func (m *TeSyncIdtInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncIdtInfo.Marshal(b, m, deterministic)
}
func (m *TeSyncIdtInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncIdtInfo.Merge(m, src)
}
func (m *TeSyncIdtInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncIdtInfo.Size(m)
}
func (m *TeSyncIdtInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncIdtInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncIdtInfo proto.InternalMessageInfo

func (m *TeSyncIdtInfo) GetCurrentIdtInfo() *TeIdtStatus {
	if m != nil {
		return m.CurrentIdtInfo
	}
	return nil
}

func (m *TeSyncIdtInfo) GetPreviousIdtStatus() *TeIdtStatus {
	if m != nil {
		return m.PreviousIdtStatus
	}
	return nil
}

type TeVifPendingInfo struct {
	PendingReason        string   `protobuf:"bytes,1,opt,name=pending_reason,json=pendingReason,proto3" json:"pending_reason,omitempty"`
	TunnelName           string   `protobuf:"bytes,2,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeVifPendingInfo) Reset()         { *m = TeVifPendingInfo{} }
func (m *TeVifPendingInfo) String() string { return proto.CompactTextString(m) }
func (*TeVifPendingInfo) ProtoMessage()    {}
func (*TeVifPendingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{3}
}

func (m *TeVifPendingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeVifPendingInfo.Unmarshal(m, b)
}
func (m *TeVifPendingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeVifPendingInfo.Marshal(b, m, deterministic)
}
func (m *TeVifPendingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeVifPendingInfo.Merge(m, src)
}
func (m *TeVifPendingInfo) XXX_Size() int {
	return xxx_messageInfo_TeVifPendingInfo.Size(m)
}
func (m *TeVifPendingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeVifPendingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeVifPendingInfo proto.InternalMessageInfo

func (m *TeVifPendingInfo) GetPendingReason() string {
	if m != nil {
		return m.PendingReason
	}
	return ""
}

func (m *TeVifPendingInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

type TeS2LPendingInfo struct {
	PendingReason        string   `protobuf:"bytes,1,opt,name=pending_reason,json=pendingReason,proto3" json:"pending_reason,omitempty"`
	SignaledName         string   `protobuf:"bytes,2,opt,name=signaled_name,json=signaledName,proto3" json:"signaled_name,omitempty"`
	S2LRole              string   `protobuf:"bytes,3,opt,name=s2l_role,json=s2lRole,proto3" json:"s2l_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeS2LPendingInfo) Reset()         { *m = TeS2LPendingInfo{} }
func (m *TeS2LPendingInfo) String() string { return proto.CompactTextString(m) }
func (*TeS2LPendingInfo) ProtoMessage()    {}
func (*TeS2LPendingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{4}
}

func (m *TeS2LPendingInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeS2LPendingInfo.Unmarshal(m, b)
}
func (m *TeS2LPendingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeS2LPendingInfo.Marshal(b, m, deterministic)
}
func (m *TeS2LPendingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeS2LPendingInfo.Merge(m, src)
}
func (m *TeS2LPendingInfo) XXX_Size() int {
	return xxx_messageInfo_TeS2LPendingInfo.Size(m)
}
func (m *TeS2LPendingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeS2LPendingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeS2LPendingInfo proto.InternalMessageInfo

func (m *TeS2LPendingInfo) GetPendingReason() string {
	if m != nil {
		return m.PendingReason
	}
	return ""
}

func (m *TeS2LPendingInfo) GetSignaledName() string {
	if m != nil {
		return m.SignaledName
	}
	return ""
}

func (m *TeS2LPendingInfo) GetS2LRole() string {
	if m != nil {
		return m.S2LRole
	}
	return ""
}

type TeSyncStatusSlaveInfo struct {
	Idt                  *TeSyncIdtInfo      `protobuf:"bytes,1,opt,name=idt,proto3" json:"idt,omitempty"`
	VifPending           []*TeVifPendingInfo `protobuf:"bytes,2,rep,name=vif_pending,json=vifPending,proto3" json:"vif_pending,omitempty"`
	S2LPending           []*TeS2LPendingInfo `protobuf:"bytes,3,rep,name=s2l_pending,json=s2lPending,proto3" json:"s2l_pending,omitempty"`
	InsyncTunnels        uint32              `protobuf:"varint,4,opt,name=insync_tunnels,json=insyncTunnels,proto3" json:"insync_tunnels,omitempty"`
	InsyncSubLsPs        uint32              `protobuf:"varint,5,opt,name=insync_sub_ls_ps,json=insyncSubLsPs,proto3" json:"insync_sub_ls_ps,omitempty"`
	PendingTunnels       uint32              `protobuf:"varint,6,opt,name=pending_tunnels,json=pendingTunnels,proto3" json:"pending_tunnels,omitempty"`
	PendingSubLsPs       uint32              `protobuf:"varint,7,opt,name=pending_sub_ls_ps,json=pendingSubLsPs,proto3" json:"pending_sub_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TeSyncStatusSlaveInfo) Reset()         { *m = TeSyncStatusSlaveInfo{} }
func (m *TeSyncStatusSlaveInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusSlaveInfo) ProtoMessage()    {}
func (*TeSyncStatusSlaveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{5}
}

func (m *TeSyncStatusSlaveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusSlaveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Marshal(b, m, deterministic)
}
func (m *TeSyncStatusSlaveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusSlaveInfo.Merge(m, src)
}
func (m *TeSyncStatusSlaveInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusSlaveInfo.Size(m)
}
func (m *TeSyncStatusSlaveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusSlaveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusSlaveInfo proto.InternalMessageInfo

func (m *TeSyncStatusSlaveInfo) GetIdt() *TeSyncIdtInfo {
	if m != nil {
		return m.Idt
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetVifPending() []*TeVifPendingInfo {
	if m != nil {
		return m.VifPending
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetS2LPending() []*TeS2LPendingInfo {
	if m != nil {
		return m.S2LPending
	}
	return nil
}

func (m *TeSyncStatusSlaveInfo) GetInsyncTunnels() uint32 {
	if m != nil {
		return m.InsyncTunnels
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetInsyncSubLsPs() uint32 {
	if m != nil {
		return m.InsyncSubLsPs
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetPendingTunnels() uint32 {
	if m != nil {
		return m.PendingTunnels
	}
	return 0
}

func (m *TeSyncStatusSlaveInfo) GetPendingSubLsPs() uint32 {
	if m != nil {
		return m.PendingSubLsPs
	}
	return 0
}

type TeSyncStatusMasterInfo struct {
	Idt                  *TeSyncIdtInfo `protobuf:"bytes,1,opt,name=idt,proto3" json:"idt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TeSyncStatusMasterInfo) Reset()         { *m = TeSyncStatusMasterInfo{} }
func (m *TeSyncStatusMasterInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusMasterInfo) ProtoMessage()    {}
func (*TeSyncStatusMasterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{6}
}

func (m *TeSyncStatusMasterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusMasterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Marshal(b, m, deterministic)
}
func (m *TeSyncStatusMasterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusMasterInfo.Merge(m, src)
}
func (m *TeSyncStatusMasterInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusMasterInfo.Size(m)
}
func (m *TeSyncStatusMasterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusMasterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusMasterInfo proto.InternalMessageInfo

func (m *TeSyncStatusMasterInfo) GetIdt() *TeSyncIdtInfo {
	if m != nil {
		return m.Idt
	}
	return nil
}

type TeSyncStatusInfo struct {
	SyncShowType          string                  `protobuf:"bytes,1,opt,name=sync_show_type,json=syncShowType,proto3" json:"sync_show_type,omitempty"`
	SlaveSyncInformation  *TeSyncStatusSlaveInfo  `protobuf:"bytes,2,opt,name=slave_sync_information,json=slaveSyncInformation,proto3" json:"slave_sync_information,omitempty"`
	MasterSyncInformation *TeSyncStatusMasterInfo `protobuf:"bytes,3,opt,name=master_sync_information,json=masterSyncInformation,proto3" json:"master_sync_information,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *TeSyncStatusInfo) Reset()         { *m = TeSyncStatusInfo{} }
func (m *TeSyncStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TeSyncStatusInfo) ProtoMessage()    {}
func (*TeSyncStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{7}
}

func (m *TeSyncStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeSyncStatusInfo.Unmarshal(m, b)
}
func (m *TeSyncStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeSyncStatusInfo.Marshal(b, m, deterministic)
}
func (m *TeSyncStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeSyncStatusInfo.Merge(m, src)
}
func (m *TeSyncStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TeSyncStatusInfo.Size(m)
}
func (m *TeSyncStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeSyncStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeSyncStatusInfo proto.InternalMessageInfo

func (m *TeSyncStatusInfo) GetSyncShowType() string {
	if m != nil {
		return m.SyncShowType
	}
	return ""
}

func (m *TeSyncStatusInfo) GetSlaveSyncInformation() *TeSyncStatusSlaveInfo {
	if m != nil {
		return m.SlaveSyncInformation
	}
	return nil
}

func (m *TeSyncStatusInfo) GetMasterSyncInformation() *TeSyncStatusMasterInfo {
	if m != nil {
		return m.MasterSyncInformation
	}
	return nil
}

type TeNsrIssuStatusInfo struct {
	Role                  string            `protobuf:"bytes,50,opt,name=role,proto3" json:"role,omitempty"`
	SyncStatusInformation *TeSyncStatusInfo `protobuf:"bytes,51,opt,name=sync_status_information,json=syncStatusInformation,proto3" json:"sync_status_information,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *TeNsrIssuStatusInfo) Reset()         { *m = TeNsrIssuStatusInfo{} }
func (m *TeNsrIssuStatusInfo) String() string { return proto.CompactTextString(m) }
func (*TeNsrIssuStatusInfo) ProtoMessage()    {}
func (*TeNsrIssuStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_42ec59026909b271, []int{8}
}

func (m *TeNsrIssuStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Unmarshal(m, b)
}
func (m *TeNsrIssuStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Marshal(b, m, deterministic)
}
func (m *TeNsrIssuStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeNsrIssuStatusInfo.Merge(m, src)
}
func (m *TeNsrIssuStatusInfo) XXX_Size() int {
	return xxx_messageInfo_TeNsrIssuStatusInfo.Size(m)
}
func (m *TeNsrIssuStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TeNsrIssuStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TeNsrIssuStatusInfo proto.InternalMessageInfo

func (m *TeNsrIssuStatusInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TeNsrIssuStatusInfo) GetSyncStatusInformation() *TeSyncStatusInfo {
	if m != nil {
		return m.SyncStatusInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*TeNsrIssuStatusInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_nsr_issu_status_info_KEYS")
	proto.RegisterType((*TeIdtStatus)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_idt_status")
	proto.RegisterType((*TeSyncIdtInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_idt_info")
	proto.RegisterType((*TeVifPendingInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_vif_pending_info")
	proto.RegisterType((*TeS2LPendingInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_s2l_pending_info")
	proto.RegisterType((*TeSyncStatusSlaveInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_slave_info")
	proto.RegisterType((*TeSyncStatusMasterInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_master_info")
	proto.RegisterType((*TeSyncStatusInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_sync_status_info")
	proto.RegisterType((*TeNsrIssuStatusInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.issu.detail.te_nsr_issu_status_info")
}

func init() { proto.RegisterFile("te_nsr_issu_status_info.proto", fileDescriptor_42ec59026909b271) }

var fileDescriptor_42ec59026909b271 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x18, 0x95, 0x09, 0x97, 0x9f, 0x2f, 0x71, 0x00, 0x73, 0x2f, 0x84, 0xab, 0xfe, 0xa4, 0x06, 0xda,
	0x74, 0xd1, 0x2c, 0xcc, 0xb6, 0x52, 0x25, 0x24, 0xa4, 0x46, 0xad, 0x2a, 0xea, 0xb0, 0xe9, 0xa2,
	0x1a, 0x19, 0xfb, 0x0b, 0x8c, 0x64, 0xcf, 0x58, 0x33, 0xe3, 0xa4, 0xa9, 0xaa, 0x4a, 0x7d, 0x9c,
	0x3e, 0x45, 0x5f, 0xa4, 0x52, 0xb7, 0x7d, 0x8c, 0x6a, 0x7e, 0x0c, 0x21, 0x88, 0x05, 0xa0, 0xee,
	0x32, 0x87, 0xc3, 0x39, 0xdf, 0x9c, 0xf3, 0x4d, 0x02, 0x0f, 0x15, 0x12, 0x26, 0x05, 0xa1, 0x52,
	0x56, 0x44, 0xaa, 0x44, 0x55, 0x92, 0x50, 0x36, 0xe2, 0xfd, 0x52, 0x70, 0xc5, 0x83, 0x17, 0x29,
	0x95, 0x29, 0x27, 0x94, 0x4b, 0xf2, 0x49, 0x90, 0xa2, 0xcc, 0x25, 0x51, 0x48, 0x78, 0x89, 0xa2,
	0xef, 0x0e, 0x7d, 0xfd, 0x9f, 0xfd, 0x0c, 0x55, 0x42, 0xf3, 0xf0, 0x11, 0x3c, 0xb8, 0x41, 0x8f,
	0xbc, 0x39, 0xfa, 0x30, 0x0c, 0x7f, 0x79, 0xe0, 0x2b, 0x24, 0x34, 0x53, 0xee, 0x6f, 0xc1, 0x53,
	0x58, 0xa3, 0x92, 0x08, 0x4c, 0xb2, 0xa9, 0x83, 0x3a, 0x5e, 0xd7, 0xeb, 0xad, 0xc4, 0x3e, 0x95,
	0xb1, 0x46, 0x87, 0x96, 0xb7, 0x05, 0x4b, 0x02, 0x13, 0xc9, 0x59, 0x67, 0xa1, 0xeb, 0xf5, 0x56,
	0x63, 0x77, 0x0a, 0xf6, 0xa0, 0xed, 0xd4, 0x84, 0x22, 0x8a, 0x16, 0xd8, 0x69, 0x74, 0xbd, 0x9e,
	0x1f, 0xb7, 0x68, 0xa6, 0x86, 0x1a, 0x3c, 0xa1, 0x05, 0x06, 0x5d, 0xd0, 0x67, 0x82, 0x2c, 0xb3,
	0x9c, 0x45, 0xc3, 0x01, 0x9a, 0xa9, 0x23, 0x96, 0x19, 0xc6, 0x13, 0x68, 0x65, 0x98, 0xe6, 0x89,
	0x40, 0xcb, 0xf8, 0xc7, 0x30, 0x9a, 0x0e, 0x33, 0x94, 0x5d, 0xf0, 0x27, 0x54, 0x9d, 0x67, 0x22,
	0x99, 0x58, 0xce, 0x92, 0x75, 0xaa, 0x41, 0x4d, 0x0a, 0x7f, 0x7b, 0xb0, 0xae, 0x90, 0xc8, 0x29,
	0x4b, 0xcd, 0x35, 0xf5, 0xdd, 0x83, 0x11, 0xac, 0xa7, 0x95, 0x10, 0xc8, 0xd4, 0x05, 0x66, 0x6e,
	0xd9, 0x8c, 0x5e, 0xf6, 0x6f, 0x15, 0x70, 0xff, 0x4a, 0x78, 0x71, 0xdb, 0xa9, 0x0e, 0x32, 0x35,
	0xd0, 0x3e, 0x39, 0x6c, 0x96, 0x02, 0xc7, 0x94, 0xeb, 0xd0, 0x2f, 0x68, 0x26, 0xb1, 0xfb, 0x5a,
	0x6d, 0xd4, 0xc2, 0x03, 0x93, 0xab, 0xaa, 0x64, 0xf8, 0x11, 0x36, 0x15, 0x92, 0x31, 0x1d, 0x91,
	0x12, 0x59, 0x46, 0xd9, 0x99, 0xbd, 0xec, 0x3e, 0xb4, 0xeb, 0xb3, 0x6b, 0xcc, 0x33, 0x8d, 0xf9,
	0x0e, 0x8d, 0x6d, 0x71, 0x8f, 0xa1, 0xa9, 0x2a, 0xc6, 0x30, 0x27, 0x2c, 0x29, 0xd0, 0xb5, 0x0a,
	0x16, 0x7a, 0x97, 0x14, 0x18, 0x7e, 0x31, 0xf2, 0x32, 0xca, 0xef, 0x24, 0xbf, 0x0b, 0xbe, 0xa4,
	0x67, 0x2c, 0xc9, 0x31, 0x9b, 0x35, 0x68, 0xd5, 0xa0, 0xb6, 0x08, 0x76, 0x60, 0x45, 0xeb, 0x0b,
	0x9e, 0xdb, 0xb5, 0x59, 0x8d, 0x97, 0x65, 0x94, 0xc7, 0x3c, 0xc7, 0xf0, 0x67, 0x03, 0x76, 0xea,
	0x1e, 0xdd, 0x1a, 0xcb, 0x3c, 0x19, 0xa3, 0x1d, 0xe2, 0x3d, 0x34, 0x68, 0xa6, 0x5c, 0x87, 0xaf,
	0x6e, 0x1f, 0xec, 0x95, 0xf5, 0x88, 0xb5, 0x56, 0x90, 0x42, 0x73, 0x26, 0xca, 0xce, 0x42, 0xb7,
	0xd1, 0x6b, 0x46, 0x87, 0xb7, 0x97, 0x9e, 0xef, 0x23, 0x86, 0x31, 0x1d, 0x1d, 0x5b, 0x40, 0x9b,
	0xcc, 0x04, 0xda, 0x69, 0xdc, 0xd5, 0x64, 0xbe, 0x95, 0x18, 0x64, 0x94, 0xd7, 0x26, 0xfb, 0xd0,
	0xa6, 0xcc, 0xdc, 0xd0, 0xb6, 0x29, 0xdd, 0x73, 0xf3, 0x2d, 0x7a, 0x62, 0xc1, 0xe0, 0x19, 0xac,
	0x3b, 0x9a, 0xac, 0x4e, 0x49, 0x2e, 0x49, 0x29, 0xdd, 0xab, 0x73, 0xc4, 0x61, 0x75, 0xfa, 0x56,
	0x1e, 0x6b, 0xe2, 0x5a, 0xed, 0x55, 0x0b, 0xda, 0x97, 0x57, 0x2f, 0x42, 0xad, 0xf8, 0x1c, 0x36,
	0x6a, 0xe2, 0xa5, 0xe4, 0xf2, 0x15, 0xaa, 0xd3, 0x0c, 0x39, 0xfc, 0x3f, 0xd7, 0x6e, 0x91, 0x48,
	0x85, 0xe2, 0x6f, 0xd5, 0x1b, 0xfe, 0x58, 0xb0, 0xeb, 0x3c, 0xe3, 0x68, 0xac, 0xf6, 0xa0, 0x6d,
	0xb1, 0x73, 0x3e, 0x21, 0x6a, 0x5a, 0xa2, 0x5b, 0xe7, 0x96, 0x49, 0xe0, 0x9c, 0x4f, 0x4e, 0xa6,
	0x25, 0x06, 0x5f, 0x61, 0xcb, 0x6e, 0x9f, 0x55, 0x66, 0x23, 0x2e, 0x8a, 0x44, 0x51, 0xf7, 0x6d,
	0xd8, 0x8c, 0x5e, 0xdf, 0x71, 0xc6, 0x6b, 0x9b, 0x1d, 0xff, 0x6b, 0x3e, 0x0f, 0xa7, 0x2c, 0x1d,
	0x5c, 0xba, 0x04, 0xdf, 0x3c, 0xd8, 0x76, 0x01, 0x5d, 0x9b, 0xa0, 0x61, 0x26, 0x18, 0xdc, 0x6f,
	0x82, 0x99, 0xf4, 0xe3, 0xff, 0xec, 0x61, 0x6e, 0x86, 0xf0, 0xbb, 0x07, 0xdb, 0x37, 0xfc, 0xb8,
	0x04, 0x01, 0x2c, 0x9a, 0x47, 0x1c, 0x99, 0xec, 0xcc, 0xe7, 0xe0, 0x33, 0x6c, 0xcf, 0xa7, 0x5d,
	0x8f, 0x7c, 0x60, 0x46, 0x3e, 0xbc, 0xdf, 0xc8, 0x76, 0x56, 0x53, 0x93, 0x01, 0x66, 0x66, 0x3d,
	0x5d, 0x32, 0xbf, 0x9e, 0x07, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0x87, 0xa0, 0xe5, 0x5e,
	0x07, 0x00, 0x00,
}
