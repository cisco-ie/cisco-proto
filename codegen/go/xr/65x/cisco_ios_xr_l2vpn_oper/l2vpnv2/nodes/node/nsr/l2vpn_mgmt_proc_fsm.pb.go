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
// source: l2vpn_mgmt_proc_fsm.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_nsr

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

type L2VpnMgmtProcFsm_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMgmtProcFsm_KEYS) Reset()         { *m = L2VpnMgmtProcFsm_KEYS{} }
func (m *L2VpnMgmtProcFsm_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsm_KEYS) ProtoMessage()    {}
func (*L2VpnMgmtProcFsm_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{0}
}

func (m *L2VpnMgmtProcFsm_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMgmtProcFsm_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMgmtProcFsm_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMgmtProcFsm_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMgmtProcFsm_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMgmtProcFsm_KEYS.Merge(m, src)
}
func (m *L2VpnMgmtProcFsm_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMgmtProcFsm_KEYS.Size(m)
}
func (m *L2VpnMgmtProcFsm_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMgmtProcFsm_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMgmtProcFsm_KEYS proto.InternalMessageInfo

func (m *L2VpnMgmtProcFsm_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type L2VpnMgmtProcFsmXidInfo struct {
	AppType              string   `protobuf:"bytes,1,opt,name=app_type,json=appType,proto3" json:"app_type,omitempty"`
	SentIds              uint32   `protobuf:"varint,2,opt,name=sent_ids,json=sentIds,proto3" json:"sent_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMgmtProcFsmXidInfo) Reset()         { *m = L2VpnMgmtProcFsmXidInfo{} }
func (m *L2VpnMgmtProcFsmXidInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsmXidInfo) ProtoMessage()    {}
func (*L2VpnMgmtProcFsmXidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{1}
}

func (m *L2VpnMgmtProcFsmXidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMgmtProcFsmXidInfo.Unmarshal(m, b)
}
func (m *L2VpnMgmtProcFsmXidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMgmtProcFsmXidInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnMgmtProcFsmXidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMgmtProcFsmXidInfo.Merge(m, src)
}
func (m *L2VpnMgmtProcFsmXidInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnMgmtProcFsmXidInfo.Size(m)
}
func (m *L2VpnMgmtProcFsmXidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMgmtProcFsmXidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMgmtProcFsmXidInfo proto.InternalMessageInfo

func (m *L2VpnMgmtProcFsmXidInfo) GetAppType() string {
	if m != nil {
		return m.AppType
	}
	return ""
}

func (m *L2VpnMgmtProcFsmXidInfo) GetSentIds() uint32 {
	if m != nil {
		return m.SentIds
	}
	return 0
}

type L2VpnFailoverStatus struct {
	TriggeredTime        uint32   `protobuf:"varint,1,opt,name=triggered_time,json=triggeredTime,proto3" json:"triggered_time,omitempty"`
	StartTime            uint32   `protobuf:"varint,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	MasterTime           uint32   `protobuf:"varint,3,opt,name=master_time,json=masterTime,proto3" json:"master_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnFailoverStatus) Reset()         { *m = L2VpnFailoverStatus{} }
func (m *L2VpnFailoverStatus) String() string { return proto.CompactTextString(m) }
func (*L2VpnFailoverStatus) ProtoMessage()    {}
func (*L2VpnFailoverStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{2}
}

func (m *L2VpnFailoverStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnFailoverStatus.Unmarshal(m, b)
}
func (m *L2VpnFailoverStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnFailoverStatus.Marshal(b, m, deterministic)
}
func (m *L2VpnFailoverStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnFailoverStatus.Merge(m, src)
}
func (m *L2VpnFailoverStatus) XXX_Size() int {
	return xxx_messageInfo_L2VpnFailoverStatus.Size(m)
}
func (m *L2VpnFailoverStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnFailoverStatus.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnFailoverStatus proto.InternalMessageInfo

func (m *L2VpnFailoverStatus) GetTriggeredTime() uint32 {
	if m != nil {
		return m.TriggeredTime
	}
	return 0
}

func (m *L2VpnFailoverStatus) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *L2VpnFailoverStatus) GetMasterTime() uint32 {
	if m != nil {
		return m.MasterTime
	}
	return 0
}

type L2VpnIdtStatus struct {
	SyncStatus           string   `protobuf:"bytes,1,opt,name=sync_status,json=syncStatus,proto3" json:"sync_status,omitempty"`
	NotReadyReason       string   `protobuf:"bytes,2,opt,name=not_ready_reason,json=notReadyReason,proto3" json:"not_ready_reason,omitempty"`
	IdtStartTime         uint32   `protobuf:"varint,3,opt,name=idt_start_time,json=idtStartTime,proto3" json:"idt_start_time,omitempty"`
	IdtEndTime           uint32   `protobuf:"varint,4,opt,name=idt_end_time,json=idtEndTime,proto3" json:"idt_end_time,omitempty"`
	DeclareTime          uint32   `protobuf:"varint,5,opt,name=declare_time,json=declareTime,proto3" json:"declare_time,omitempty"`
	WithdrawTime         uint32   `protobuf:"varint,6,opt,name=withdraw_time,json=withdrawTime,proto3" json:"withdraw_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnIdtStatus) Reset()         { *m = L2VpnIdtStatus{} }
func (m *L2VpnIdtStatus) String() string { return proto.CompactTextString(m) }
func (*L2VpnIdtStatus) ProtoMessage()    {}
func (*L2VpnIdtStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{3}
}

func (m *L2VpnIdtStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnIdtStatus.Unmarshal(m, b)
}
func (m *L2VpnIdtStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnIdtStatus.Marshal(b, m, deterministic)
}
func (m *L2VpnIdtStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnIdtStatus.Merge(m, src)
}
func (m *L2VpnIdtStatus) XXX_Size() int {
	return xxx_messageInfo_L2VpnIdtStatus.Size(m)
}
func (m *L2VpnIdtStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnIdtStatus.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnIdtStatus proto.InternalMessageInfo

func (m *L2VpnIdtStatus) GetSyncStatus() string {
	if m != nil {
		return m.SyncStatus
	}
	return ""
}

func (m *L2VpnIdtStatus) GetNotReadyReason() string {
	if m != nil {
		return m.NotReadyReason
	}
	return ""
}

func (m *L2VpnIdtStatus) GetIdtStartTime() uint32 {
	if m != nil {
		return m.IdtStartTime
	}
	return 0
}

func (m *L2VpnIdtStatus) GetIdtEndTime() uint32 {
	if m != nil {
		return m.IdtEndTime
	}
	return 0
}

func (m *L2VpnIdtStatus) GetDeclareTime() uint32 {
	if m != nil {
		return m.DeclareTime
	}
	return 0
}

func (m *L2VpnIdtStatus) GetWithdrawTime() uint32 {
	if m != nil {
		return m.WithdrawTime
	}
	return 0
}

type L2VpnNsrIssuStatusInfo struct {
	NsrRole              uint32          `protobuf:"varint,1,opt,name=nsr_role,json=nsrRole,proto3" json:"nsr_role,omitempty"`
	IssuRole             uint32          `protobuf:"varint,2,opt,name=issu_role,json=issuRole,proto3" json:"issu_role,omitempty"`
	IdtStatus            *L2VpnIdtStatus `protobuf:"bytes,3,opt,name=idt_status,json=idtStatus,proto3" json:"idt_status,omitempty"`
	PreviOusIdtStatus    *L2VpnIdtStatus `protobuf:"bytes,4,opt,name=previ_ous_idt_status,json=previOusIdtStatus,proto3" json:"previ_ous_idt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *L2VpnNsrIssuStatusInfo) Reset()         { *m = L2VpnNsrIssuStatusInfo{} }
func (m *L2VpnNsrIssuStatusInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnNsrIssuStatusInfo) ProtoMessage()    {}
func (*L2VpnNsrIssuStatusInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{4}
}

func (m *L2VpnNsrIssuStatusInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnNsrIssuStatusInfo.Unmarshal(m, b)
}
func (m *L2VpnNsrIssuStatusInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnNsrIssuStatusInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnNsrIssuStatusInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnNsrIssuStatusInfo.Merge(m, src)
}
func (m *L2VpnNsrIssuStatusInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnNsrIssuStatusInfo.Size(m)
}
func (m *L2VpnNsrIssuStatusInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnNsrIssuStatusInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnNsrIssuStatusInfo proto.InternalMessageInfo

func (m *L2VpnNsrIssuStatusInfo) GetNsrRole() uint32 {
	if m != nil {
		return m.NsrRole
	}
	return 0
}

func (m *L2VpnNsrIssuStatusInfo) GetIssuRole() uint32 {
	if m != nil {
		return m.IssuRole
	}
	return 0
}

func (m *L2VpnNsrIssuStatusInfo) GetIdtStatus() *L2VpnIdtStatus {
	if m != nil {
		return m.IdtStatus
	}
	return nil
}

func (m *L2VpnNsrIssuStatusInfo) GetPreviOusIdtStatus() *L2VpnIdtStatus {
	if m != nil {
		return m.PreviOusIdtStatus
	}
	return nil
}

type L2VpnMgmtProcFsmReportCard struct {
	CollaboratorIsConnected bool     `protobuf:"varint,1,opt,name=collaborator_is_connected,json=collaboratorIsConnected,proto3" json:"collaborator_is_connected,omitempty"`
	ConnectionChangeTime    uint32   `protobuf:"varint,2,opt,name=connection_change_time,json=connectionChangeTime,proto3" json:"connection_change_time,omitempty"`
	CollaboratorIdtDone     bool     `protobuf:"varint,3,opt,name=collaborator_idt_done,json=collaboratorIdtDone,proto3" json:"collaborator_idt_done,omitempty"`
	IdtTime                 uint32   `protobuf:"varint,4,opt,name=idt_time,json=idtTime,proto3" json:"idt_time,omitempty"`
	CollaboratorSkipped     bool     `protobuf:"varint,5,opt,name=collaborator_skipped,json=collaboratorSkipped,proto3" json:"collaborator_skipped,omitempty"`
	ExpectIdt               bool     `protobuf:"varint,6,opt,name=expect_idt,json=expectIdt,proto3" json:"expect_idt,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *L2VpnMgmtProcFsmReportCard) Reset()         { *m = L2VpnMgmtProcFsmReportCard{} }
func (m *L2VpnMgmtProcFsmReportCard) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsmReportCard) ProtoMessage()    {}
func (*L2VpnMgmtProcFsmReportCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{5}
}

func (m *L2VpnMgmtProcFsmReportCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMgmtProcFsmReportCard.Unmarshal(m, b)
}
func (m *L2VpnMgmtProcFsmReportCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMgmtProcFsmReportCard.Marshal(b, m, deterministic)
}
func (m *L2VpnMgmtProcFsmReportCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMgmtProcFsmReportCard.Merge(m, src)
}
func (m *L2VpnMgmtProcFsmReportCard) XXX_Size() int {
	return xxx_messageInfo_L2VpnMgmtProcFsmReportCard.Size(m)
}
func (m *L2VpnMgmtProcFsmReportCard) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMgmtProcFsmReportCard.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMgmtProcFsmReportCard proto.InternalMessageInfo

func (m *L2VpnMgmtProcFsmReportCard) GetCollaboratorIsConnected() bool {
	if m != nil {
		return m.CollaboratorIsConnected
	}
	return false
}

func (m *L2VpnMgmtProcFsmReportCard) GetConnectionChangeTime() uint32 {
	if m != nil {
		return m.ConnectionChangeTime
	}
	return 0
}

func (m *L2VpnMgmtProcFsmReportCard) GetCollaboratorIdtDone() bool {
	if m != nil {
		return m.CollaboratorIdtDone
	}
	return false
}

func (m *L2VpnMgmtProcFsmReportCard) GetIdtTime() uint32 {
	if m != nil {
		return m.IdtTime
	}
	return 0
}

func (m *L2VpnMgmtProcFsmReportCard) GetCollaboratorSkipped() bool {
	if m != nil {
		return m.CollaboratorSkipped
	}
	return false
}

func (m *L2VpnMgmtProcFsmReportCard) GetExpectIdt() bool {
	if m != nil {
		return m.ExpectIdt
	}
	return false
}

type L2VpnMgmtProcFsm struct {
	HaRole               uint32                        `protobuf:"varint,50,opt,name=ha_role,json=haRole,proto3" json:"ha_role,omitempty"`
	IssuRole             uint32                        `protobuf:"varint,51,opt,name=issu_role,json=issuRole,proto3" json:"issu_role,omitempty"`
	SyncFlags            uint32                        `protobuf:"varint,52,opt,name=sync_flags,json=syncFlags,proto3" json:"sync_flags,omitempty"`
	SwInstallInProgress  bool                          `protobuf:"varint,53,opt,name=sw_install_in_progress,json=swInstallInProgress,proto3" json:"sw_install_in_progress,omitempty"`
	XidInfo              []*L2VpnMgmtProcFsmXidInfo    `protobuf:"bytes,54,rep,name=xid_info,json=xidInfo,proto3" json:"xid_info,omitempty"`
	FailoverStatus       *L2VpnFailoverStatus          `protobuf:"bytes,55,opt,name=failover_status,json=failoverStatus,proto3" json:"failover_status,omitempty"`
	NsrStatus            *L2VpnNsrIssuStatusInfo       `protobuf:"bytes,56,opt,name=nsr_status,json=nsrStatus,proto3" json:"nsr_status,omitempty"`
	IssuStatus           *L2VpnNsrIssuStatusInfo       `protobuf:"bytes,57,opt,name=issu_status,json=issuStatus,proto3" json:"issu_status,omitempty"`
	ReportCard           []*L2VpnMgmtProcFsmReportCard `protobuf:"bytes,58,rep,name=report_card,json=reportCard,proto3" json:"report_card,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *L2VpnMgmtProcFsm) Reset()         { *m = L2VpnMgmtProcFsm{} }
func (m *L2VpnMgmtProcFsm) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsm) ProtoMessage()    {}
func (*L2VpnMgmtProcFsm) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{6}
}

func (m *L2VpnMgmtProcFsm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMgmtProcFsm.Unmarshal(m, b)
}
func (m *L2VpnMgmtProcFsm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMgmtProcFsm.Marshal(b, m, deterministic)
}
func (m *L2VpnMgmtProcFsm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMgmtProcFsm.Merge(m, src)
}
func (m *L2VpnMgmtProcFsm) XXX_Size() int {
	return xxx_messageInfo_L2VpnMgmtProcFsm.Size(m)
}
func (m *L2VpnMgmtProcFsm) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMgmtProcFsm.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMgmtProcFsm proto.InternalMessageInfo

func (m *L2VpnMgmtProcFsm) GetHaRole() uint32 {
	if m != nil {
		return m.HaRole
	}
	return 0
}

func (m *L2VpnMgmtProcFsm) GetIssuRole() uint32 {
	if m != nil {
		return m.IssuRole
	}
	return 0
}

func (m *L2VpnMgmtProcFsm) GetSyncFlags() uint32 {
	if m != nil {
		return m.SyncFlags
	}
	return 0
}

func (m *L2VpnMgmtProcFsm) GetSwInstallInProgress() bool {
	if m != nil {
		return m.SwInstallInProgress
	}
	return false
}

func (m *L2VpnMgmtProcFsm) GetXidInfo() []*L2VpnMgmtProcFsmXidInfo {
	if m != nil {
		return m.XidInfo
	}
	return nil
}

func (m *L2VpnMgmtProcFsm) GetFailoverStatus() *L2VpnFailoverStatus {
	if m != nil {
		return m.FailoverStatus
	}
	return nil
}

func (m *L2VpnMgmtProcFsm) GetNsrStatus() *L2VpnNsrIssuStatusInfo {
	if m != nil {
		return m.NsrStatus
	}
	return nil
}

func (m *L2VpnMgmtProcFsm) GetIssuStatus() *L2VpnNsrIssuStatusInfo {
	if m != nil {
		return m.IssuStatus
	}
	return nil
}

func (m *L2VpnMgmtProcFsm) GetReportCard() []*L2VpnMgmtProcFsmReportCard {
	if m != nil {
		return m.ReportCard
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnMgmtProcFsm_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_mgmt_proc_fsm_KEYS")
	proto.RegisterType((*L2VpnMgmtProcFsmXidInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_mgmt_proc_fsm_xid_info")
	proto.RegisterType((*L2VpnFailoverStatus)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_failover_status")
	proto.RegisterType((*L2VpnIdtStatus)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_idt_status")
	proto.RegisterType((*L2VpnNsrIssuStatusInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_nsr_issu_status_info")
	proto.RegisterType((*L2VpnMgmtProcFsmReportCard)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_mgmt_proc_fsm_report_card")
	proto.RegisterType((*L2VpnMgmtProcFsm)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.nsr.l2vpn_mgmt_proc_fsm")
}

func init() { proto.RegisterFile("l2vpn_mgmt_proc_fsm.proto", fileDescriptor_19aa22cfcf96c09f) }

var fileDescriptor_19aa22cfcf96c09f = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4b, 0x6e, 0x1b, 0x39,
	0x10, 0x85, 0xfc, 0xd1, 0xa7, 0x64, 0x69, 0x3c, 0xed, 0x5f, 0x7b, 0x3e, 0xb0, 0x46, 0x33, 0x03,
	0x68, 0x25, 0x60, 0x24, 0xcf, 0x4c, 0xe2, 0x55, 0x00, 0xc7, 0x01, 0x3a, 0x09, 0xe0, 0x80, 0xf2,
	0x26, 0x2b, 0x82, 0x6e, 0x52, 0x12, 0xe1, 0x16, 0xc9, 0x90, 0xb4, 0x3e, 0x9b, 0x9c, 0x27, 0x07,
	0xc8, 0x09, 0x72, 0x9f, 0xdc, 0x21, 0x20, 0xd9, 0x6d, 0xcb, 0x8a, 0x36, 0x31, 0xbc, 0x11, 0xcc,
	0x57, 0xaf, 0xeb, 0x55, 0x15, 0x5f, 0xd1, 0x70, 0x9c, 0xf5, 0xa6, 0x4a, 0xe0, 0xc9, 0x68, 0x62,
	0xb1, 0xd2, 0x32, 0xc5, 0x43, 0x33, 0xe9, 0x2a, 0x2d, 0xad, 0x8c, 0xba, 0x29, 0x37, 0xa9, 0xc4,
	0x5c, 0x1a, 0x3c, 0xd7, 0x38, 0xf0, 0xa4, 0x62, 0xba, 0xeb, 0xff, 0x9c, 0xf6, 0xba, 0x42, 0x52,
	0x66, 0xfc, 0x6f, 0x57, 0x18, 0xdd, 0xee, 0x43, 0xbc, 0x26, 0x19, 0x7e, 0x73, 0xf1, 0x7e, 0x10,
	0x1d, 0x41, 0xc5, 0xf1, 0x30, 0xa7, 0x71, 0xa9, 0x55, 0xea, 0xd4, 0x50, 0xd9, 0x1d, 0x13, 0xda,
	0xbe, 0x82, 0xdf, 0xd6, 0x7d, 0x34, 0xe7, 0x14, 0x73, 0x31, 0x94, 0xd1, 0x31, 0x54, 0x89, 0x52,
	0xd8, 0x2e, 0x14, 0xcb, 0xbf, 0xac, 0x10, 0xa5, 0xae, 0x16, 0x8a, 0xb9, 0x90, 0x61, 0xc2, 0x62,
	0x4e, 0x4d, 0xbc, 0xd1, 0x2a, 0x75, 0x1a, 0xa8, 0xe2, 0xce, 0x09, 0x35, 0xed, 0x8f, 0x70, 0x10,
	0xb2, 0x0e, 0x09, 0xcf, 0xe4, 0x94, 0x69, 0x6c, 0x2c, 0xb1, 0xb7, 0x26, 0xfa, 0x1b, 0x9a, 0x56,
	0xf3, 0xd1, 0x88, 0x69, 0x46, 0xb1, 0xe5, 0x93, 0x90, 0xb4, 0x81, 0x1a, 0x77, 0xe8, 0x15, 0x9f,
	0xb0, 0xe8, 0x77, 0x00, 0x63, 0x89, 0xb6, 0x81, 0x12, 0x92, 0xd7, 0x3c, 0xe2, 0xc3, 0x27, 0x50,
	0x9f, 0x10, 0x63, 0x99, 0x0e, 0xf1, 0x4d, 0x1f, 0x87, 0x00, 0x39, 0x42, 0xfb, 0x6b, 0x09, 0x76,
	0x43, 0x01, 0x9c, 0xda, 0x42, 0xfb, 0x04, 0xea, 0x66, 0x21, 0xd2, 0xfc, 0x98, 0x77, 0x03, 0x0e,
	0x1a, 0x04, 0x42, 0x07, 0x76, 0x85, 0xb4, 0x58, 0x33, 0x42, 0x17, 0xee, 0xd7, 0x48, 0xe1, 0xb5,
	0x6b, 0xa8, 0x29, 0xa4, 0x45, 0x0e, 0x46, 0x1e, 0x8d, 0xfe, 0x82, 0x66, 0x9e, 0xb8, 0xa8, 0x31,
	0xd4, 0xb0, 0xc3, 0xa9, 0x1d, 0xdc, 0x95, 0xd9, 0x02, 0x77, 0xc6, 0x4c, 0xe4, 0xad, 0x6e, 0x85,
	0x3a, 0x39, 0xb5, 0x17, 0x22, 0xf4, 0xf9, 0x07, 0xec, 0x50, 0x96, 0x66, 0x44, 0xb3, 0xc0, 0xd8,
	0xf6, 0x8c, 0x7a, 0x8e, 0x79, 0xca, 0x9f, 0xd0, 0x98, 0x71, 0x3b, 0xa6, 0x9a, 0xcc, 0x02, 0xa7,
	0x1c, 0x94, 0x0a, 0xd0, 0xf7, 0xfb, 0x69, 0x03, 0x7e, 0x09, 0xfd, 0x0a, 0xa3, 0x31, 0x37, 0xe6,
	0x36, 0xef, 0xf2, 0xee, 0x12, 0x1d, 0xae, 0x65, 0x56, 0xcc, 0xbb, 0x22, 0x8c, 0x46, 0x32, 0x63,
	0xd1, 0xaf, 0x50, 0xf3, 0x74, 0x1f, 0x0b, 0x83, 0xae, 0x3a, 0xc0, 0x07, 0x31, 0xc0, 0xfd, 0xfc,
	0x7c, 0x8b, 0xf5, 0xde, 0x8b, 0x1f, 0xb4, 0x65, 0x77, 0xf5, 0x1e, 0x50, 0x2d, 0x0c, 0xc9, 0x4d,
	0xfc, 0x03, 0xec, 0x2b, 0xcd, 0xa6, 0x1c, 0x4b, 0x57, 0xea, 0xbd, 0xd4, 0xd6, 0x13, 0x49, 0xfd,
	0xec, 0xb3, 0x5f, 0xde, 0x9a, 0xa4, 0x90, 0x6c, 0x7f, 0xde, 0x80, 0x93, 0x75, 0x8e, 0xd7, 0x4c,
	0x49, 0x6d, 0x71, 0x4a, 0x34, 0x8d, 0xce, 0xe0, 0x38, 0x95, 0x59, 0x46, 0xae, 0xa5, 0x26, 0x56,
	0xba, 0x81, 0xe2, 0x54, 0x0a, 0xc1, 0x52, 0xcb, 0xc2, 0xfe, 0x54, 0xd1, 0xd1, 0x32, 0x21, 0x31,
	0xe7, 0x45, 0x38, 0x3a, 0x85, 0xc3, 0x9c, 0xcb, 0xa5, 0xc0, 0xe9, 0x98, 0x88, 0x11, 0x5b, 0xb6,
	0xf1, 0xfe, 0x7d, 0xf4, 0xdc, 0x07, 0xfd, 0x2d, 0xf7, 0xe0, 0xe0, 0xa1, 0x22, 0xb5, 0x98, 0x4a,
	0x11, 0x7c, 0x55, 0x45, 0x7b, 0x0f, 0xd4, 0xa8, 0x7d, 0x29, 0x85, 0xdf, 0x3f, 0x47, 0x5b, 0xb2,
	0x56, 0x85, 0xd3, 0xe0, 0xbc, 0x7f, 0x60, 0xff, 0x41, 0x3a, 0x73, 0xc3, 0x95, 0x62, 0xd4, 0xfb,
	0x6b, 0x25, 0xdb, 0x20, 0x84, 0xdc, 0xca, 0xb1, 0xb9, 0x62, 0xa9, 0xdb, 0x67, 0xeb, 0x4d, 0x56,
	0x45, 0xb5, 0x80, 0x24, 0xd4, 0xb6, 0xbf, 0x6c, 0xc3, 0xde, 0x9a, 0xb1, 0xb9, 0x87, 0x65, 0x4c,
	0x82, 0x7b, 0x7a, 0xbe, 0x86, 0xf2, 0x98, 0x7c, 0x6f, 0xac, 0xfe, 0x8a, 0xb1, 0xdc, 0x7e, 0xbb,
	0x55, 0x1c, 0x66, 0x64, 0x64, 0xe2, 0xd3, 0x7c, 0xbf, 0x17, 0x22, 0x7d, 0xe5, 0x80, 0xa8, 0x0f,
	0x87, 0x66, 0x86, 0xb9, 0x30, 0x96, 0x64, 0x19, 0xe6, 0xc2, 0xc9, 0x8d, 0x34, 0x33, 0x26, 0xfe,
	0x37, 0x34, 0x60, 0x66, 0x49, 0x08, 0x26, 0xe2, 0x5d, 0x1e, 0x8a, 0x46, 0x50, 0x2d, 0x5e, 0xad,
	0xf8, 0xbf, 0xd6, 0x66, 0xa7, 0xde, 0x7b, 0xfb, 0x38, 0xff, 0xac, 0x7f, 0x09, 0x51, 0x65, 0xce,
	0x69, 0xe2, 0xb6, 0x49, 0xc0, 0x4f, 0x2b, 0xcf, 0x5a, 0xfc, 0xbf, 0xf7, 0xeb, 0xc5, 0xe3, 0xf4,
	0x56, 0x92, 0xa1, 0x66, 0x01, 0xe4, 0x4b, 0xc2, 0x01, 0xdc, 0xf6, 0xe6, 0x52, 0xcf, 0xbc, 0xd4,
	0xeb, 0xc7, 0x49, 0xad, 0x7b, 0x1d, 0x50, 0x4d, 0x98, 0x42, 0xea, 0x06, 0xea, 0x4b, 0xe1, 0xf8,
	0xf9, 0x93, 0x6b, 0x81, 0x43, 0x72, 0x31, 0x05, 0xf5, 0xa5, 0xa5, 0x8b, 0xcf, 0xfc, 0x9d, 0x5d,
	0x3e, 0xc5, 0x9d, 0x2d, 0xa5, 0x45, 0x10, 0x0e, 0xe7, 0x44, 0xd3, 0xeb, 0xb2, 0xff, 0xc7, 0xda,
	0xff, 0x16, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xb0, 0x64, 0xf1, 0x75, 0x07, 0x00, 0x00,
}
