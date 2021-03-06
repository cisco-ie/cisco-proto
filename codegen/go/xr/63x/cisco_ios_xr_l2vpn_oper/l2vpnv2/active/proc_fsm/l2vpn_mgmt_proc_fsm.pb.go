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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_proc_fsm

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
	Role                 string          `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	IdtStatus            *L2VpnIdtStatus `protobuf:"bytes,2,opt,name=idt_status,json=idtStatus,proto3" json:"idt_status,omitempty"`
	PreviOusIdtStatus    *L2VpnIdtStatus `protobuf:"bytes,3,opt,name=previ_ous_idt_status,json=previOusIdtStatus,proto3" json:"previ_ous_idt_status,omitempty"`
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

func (m *L2VpnNsrIssuStatusInfo) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
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

type L2VpnMgmtProcFsmEventHistory struct {
	PreviousState        uint32   `protobuf:"varint,1,opt,name=previous_state,json=previousState,proto3" json:"previous_state,omitempty"`
	NewState             uint32   `protobuf:"varint,2,opt,name=new_state,json=newState,proto3" json:"new_state,omitempty"`
	ProcessEvent         uint32   `protobuf:"varint,3,opt,name=process_event,json=processEvent,proto3" json:"process_event,omitempty"`
	ProcessCollaborator  uint32   `protobuf:"varint,4,opt,name=process_collaborator,json=processCollaborator,proto3" json:"process_collaborator,omitempty"`
	EventTimestamp       uint32   `protobuf:"varint,5,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMgmtProcFsmEventHistory) Reset()         { *m = L2VpnMgmtProcFsmEventHistory{} }
func (m *L2VpnMgmtProcFsmEventHistory) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsmEventHistory) ProtoMessage()    {}
func (*L2VpnMgmtProcFsmEventHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{6}
}

func (m *L2VpnMgmtProcFsmEventHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMgmtProcFsmEventHistory.Unmarshal(m, b)
}
func (m *L2VpnMgmtProcFsmEventHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMgmtProcFsmEventHistory.Marshal(b, m, deterministic)
}
func (m *L2VpnMgmtProcFsmEventHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMgmtProcFsmEventHistory.Merge(m, src)
}
func (m *L2VpnMgmtProcFsmEventHistory) XXX_Size() int {
	return xxx_messageInfo_L2VpnMgmtProcFsmEventHistory.Size(m)
}
func (m *L2VpnMgmtProcFsmEventHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMgmtProcFsmEventHistory.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMgmtProcFsmEventHistory proto.InternalMessageInfo

func (m *L2VpnMgmtProcFsmEventHistory) GetPreviousState() uint32 {
	if m != nil {
		return m.PreviousState
	}
	return 0
}

func (m *L2VpnMgmtProcFsmEventHistory) GetNewState() uint32 {
	if m != nil {
		return m.NewState
	}
	return 0
}

func (m *L2VpnMgmtProcFsmEventHistory) GetProcessEvent() uint32 {
	if m != nil {
		return m.ProcessEvent
	}
	return 0
}

func (m *L2VpnMgmtProcFsmEventHistory) GetProcessCollaborator() uint32 {
	if m != nil {
		return m.ProcessCollaborator
	}
	return 0
}

func (m *L2VpnMgmtProcFsmEventHistory) GetEventTimestamp() uint32 {
	if m != nil {
		return m.EventTimestamp
	}
	return 0
}

type L2VpnMgmtProcFsm struct {
	ProcessRole          uint32                          `protobuf:"varint,50,opt,name=process_role,json=processRole,proto3" json:"process_role,omitempty"`
	ProcessState         uint32                          `protobuf:"varint,51,opt,name=process_state,json=processState,proto3" json:"process_state,omitempty"`
	SyncFlags            uint32                          `protobuf:"varint,52,opt,name=sync_flags,json=syncFlags,proto3" json:"sync_flags,omitempty"`
	SwInstallInProgress  bool                            `protobuf:"varint,53,opt,name=sw_install_in_progress,json=swInstallInProgress,proto3" json:"sw_install_in_progress,omitempty"`
	XidInfo              []*L2VpnMgmtProcFsmXidInfo      `protobuf:"bytes,54,rep,name=xid_info,json=xidInfo,proto3" json:"xid_info,omitempty"`
	FailoverStatus       *L2VpnFailoverStatus            `protobuf:"bytes,55,opt,name=failover_status,json=failoverStatus,proto3" json:"failover_status,omitempty"`
	StateTransitionTime  []uint32                        `protobuf:"varint,56,rep,packed,name=state_transition_time,json=stateTransitionTime,proto3" json:"state_transition_time,omitempty"`
	NsrStatus            *L2VpnNsrIssuStatusInfo         `protobuf:"bytes,57,opt,name=nsr_status,json=nsrStatus,proto3" json:"nsr_status,omitempty"`
	IssuStatus           *L2VpnNsrIssuStatusInfo         `protobuf:"bytes,58,opt,name=issu_status,json=issuStatus,proto3" json:"issu_status,omitempty"`
	ReportCard           []*L2VpnMgmtProcFsmReportCard   `protobuf:"bytes,59,rep,name=report_card,json=reportCard,proto3" json:"report_card,omitempty"`
	EventHistory         []*L2VpnMgmtProcFsmEventHistory `protobuf:"bytes,60,rep,name=event_history,json=eventHistory,proto3" json:"event_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *L2VpnMgmtProcFsm) Reset()         { *m = L2VpnMgmtProcFsm{} }
func (m *L2VpnMgmtProcFsm) String() string { return proto.CompactTextString(m) }
func (*L2VpnMgmtProcFsm) ProtoMessage()    {}
func (*L2VpnMgmtProcFsm) Descriptor() ([]byte, []int) {
	return fileDescriptor_19aa22cfcf96c09f, []int{7}
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

func (m *L2VpnMgmtProcFsm) GetProcessRole() uint32 {
	if m != nil {
		return m.ProcessRole
	}
	return 0
}

func (m *L2VpnMgmtProcFsm) GetProcessState() uint32 {
	if m != nil {
		return m.ProcessState
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

func (m *L2VpnMgmtProcFsm) GetStateTransitionTime() []uint32 {
	if m != nil {
		return m.StateTransitionTime
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

func (m *L2VpnMgmtProcFsm) GetEventHistory() []*L2VpnMgmtProcFsmEventHistory {
	if m != nil {
		return m.EventHistory
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnMgmtProcFsm_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_mgmt_proc_fsm_KEYS")
	proto.RegisterType((*L2VpnMgmtProcFsmXidInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_mgmt_proc_fsm_xid_info")
	proto.RegisterType((*L2VpnFailoverStatus)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_failover_status")
	proto.RegisterType((*L2VpnIdtStatus)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_idt_status")
	proto.RegisterType((*L2VpnNsrIssuStatusInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_nsr_issu_status_info")
	proto.RegisterType((*L2VpnMgmtProcFsmReportCard)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_mgmt_proc_fsm_report_card")
	proto.RegisterType((*L2VpnMgmtProcFsmEventHistory)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_mgmt_proc_fsm_event_history")
	proto.RegisterType((*L2VpnMgmtProcFsm)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.proc_fsm.l2vpn_mgmt_proc_fsm")
}

func init() { proto.RegisterFile("l2vpn_mgmt_proc_fsm.proto", fileDescriptor_19aa22cfcf96c09f) }

var fileDescriptor_19aa22cfcf96c09f = []byte{
	// 864 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x73, 0x1b, 0x35,
	0x14, 0x1f, 0x27, 0xa5, 0xb6, 0x9f, 0x63, 0xb7, 0x6c, 0x92, 0xb2, 0x29, 0x30, 0x71, 0x5d, 0x18,
	0x72, 0x32, 0x83, 0x53, 0xfe, 0x15, 0x2e, 0x4c, 0x48, 0x07, 0x4f, 0x87, 0xa1, 0x23, 0xfb, 0xc2,
	0x49, 0xa8, 0x2b, 0xd9, 0x16, 0xac, 0xa5, 0x45, 0x52, 0xec, 0xf8, 0xc2, 0xb1, 0xdf, 0x80, 0xef,
	0xc8, 0x85, 0xef, 0xc0, 0x48, 0x4f, 0xeb, 0xac, 0x3d, 0xbe, 0x90, 0xf1, 0x25, 0xb3, 0xfa, 0xfd,
	0x5e, 0xde, 0xfb, 0xbd, 0xbf, 0x63, 0x38, 0xcb, 0x07, 0x8b, 0x42, 0xd1, 0xf9, 0x74, 0xee, 0x68,
	0x61, 0x74, 0x46, 0x27, 0x76, 0xde, 0x2f, 0x8c, 0x76, 0x3a, 0xf9, 0x3c, 0x93, 0x36, 0xd3, 0x54,
	0x6a, 0x4b, 0x6f, 0x0d, 0x45, 0x3b, 0x5d, 0x08, 0xd3, 0x0f, 0x9f, 0x8b, 0x41, 0x9f, 0x65, 0x4e,
	0x2e, 0x44, 0xbf, 0xfc, 0xb7, 0xde, 0x53, 0x48, 0x77, 0x78, 0xa3, 0xaf, 0xaf, 0x7f, 0x1d, 0xf5,
	0xc6, 0xf0, 0xd1, 0x2e, 0xee, 0x56, 0x72, 0x2a, 0xd5, 0x44, 0x27, 0x67, 0xd0, 0x60, 0x45, 0x41,
	0xdd, 0xaa, 0x10, 0x69, 0xad, 0x5b, 0xbb, 0x68, 0x92, 0x3a, 0x2b, 0x8a, 0xf1, 0xaa, 0x10, 0x9e,
	0xb2, 0x42, 0x39, 0x2a, 0xb9, 0x4d, 0x0f, 0xba, 0xb5, 0x8b, 0x36, 0xa9, 0xfb, 0xf7, 0x90, 0xdb,
	0xde, 0x5f, 0x70, 0x8a, 0x5e, 0x27, 0x4c, 0xe6, 0x7a, 0x21, 0x0c, 0xb5, 0x8e, 0xb9, 0x1b, 0x9b,
	0x7c, 0x0a, 0x1d, 0x67, 0xe4, 0x74, 0x2a, 0x8c, 0xe0, 0xd4, 0xc9, 0x39, 0x3a, 0x6d, 0x93, 0xf6,
	0x1a, 0x1d, 0xcb, 0xb9, 0x48, 0x3e, 0x06, 0xb0, 0x8e, 0x19, 0x87, 0x26, 0xe8, 0xbc, 0x19, 0x90,
	0x40, 0x9f, 0x43, 0x6b, 0xce, 0xac, 0x13, 0x06, 0xf9, 0xc3, 0xc0, 0x03, 0x42, 0xde, 0xa0, 0xf7,
	0x6f, 0x0d, 0x1e, 0xa3, 0x00, 0xc9, 0x5d, 0x19, 0xfb, 0x1c, 0x5a, 0x76, 0xa5, 0xb2, 0xf8, 0x8c,
	0xd9, 0x80, 0x87, 0x46, 0x68, 0x70, 0x01, 0x8f, 0x95, 0x76, 0xd4, 0x08, 0xc6, 0x57, 0xfe, 0xaf,
	0xd5, 0x2a, 0xc4, 0x6e, 0x92, 0x8e, 0xd2, 0x8e, 0x78, 0x98, 0x04, 0x34, 0xf9, 0x04, 0x3a, 0xd1,
	0x71, 0xa9, 0x11, 0x35, 0x1c, 0x49, 0xee, 0x46, 0x6b, 0x99, 0x5d, 0xf0, 0x6f, 0x2a, 0x54, 0x4c,
	0xf5, 0x01, 0xea, 0x94, 0xdc, 0x5d, 0x2b, 0xcc, 0xf3, 0x19, 0x1c, 0x71, 0x91, 0xe5, 0xcc, 0x08,
	0xb4, 0x78, 0x2f, 0x58, 0xb4, 0x22, 0x16, 0x4c, 0x9e, 0x43, 0x7b, 0x29, 0xdd, 0x8c, 0x1b, 0xb6,
	0x44, 0x9b, 0x87, 0x18, 0xa9, 0x04, 0x43, 0xbe, 0xef, 0x0e, 0xe0, 0x29, 0xe6, 0xab, 0xac, 0xa1,
	0xd2, 0xda, 0x9b, 0x98, 0x25, 0x36, 0x31, 0x81, 0x07, 0x46, 0xe7, 0x65, 0x03, 0xc3, 0x77, 0xf2,
	0x1b, 0xc0, 0x5d, 0x6d, 0x42, 0x9a, 0xad, 0xc1, 0x0f, 0xfd, 0xff, 0x39, 0x5a, 0xfd, 0xed, 0x22,
	0x93, 0x26, 0x56, 0xc0, 0x97, 0xd3, 0xc0, 0x49, 0x61, 0xc4, 0x42, 0x52, 0xed, 0x75, 0xdc, 0xc5,
	0x3a, 0xdc, 0x57, 0xac, 0xf7, 0x83, 0xfb, 0x5f, 0x6e, 0xec, 0xb0, 0x8c, 0xd9, 0xfb, 0xfb, 0x00,
	0xce, 0x77, 0xcd, 0xb3, 0x11, 0x85, 0x36, 0x8e, 0x66, 0xcc, 0xf0, 0xe4, 0x25, 0x9c, 0x65, 0x3a,
	0xcf, 0xd9, 0x5b, 0x6d, 0x98, 0xd3, 0xbe, 0x5c, 0x34, 0xd3, 0x4a, 0x89, 0xcc, 0x09, 0x1e, 0x4a,
	0xd4, 0x20, 0x1f, 0x54, 0x0d, 0x86, 0xf6, 0xaa, 0xa4, 0x93, 0x17, 0xf0, 0x24, 0xda, 0x4a, 0xad,
	0x68, 0x36, 0x63, 0x6a, 0x2a, 0xaa, 0x43, 0x7a, 0x72, 0xc7, 0x5e, 0x05, 0x32, 0xf4, 0x70, 0x00,
	0xa7, 0x9b, 0x11, 0xb9, 0xa3, 0x5c, 0x2b, 0x9c, 0x9a, 0x06, 0x39, 0xde, 0x88, 0xc6, 0xdd, 0x8f,
	0x5a, 0x85, 0xed, 0xf2, 0x66, 0x95, 0xc1, 0xa9, 0x4b, 0x8e, 0x73, 0xf5, 0x05, 0x9c, 0x6c, 0xb8,
	0xb3, 0x7f, 0xc8, 0xa2, 0x10, 0x3c, 0x4c, 0xcf, 0x96, 0xb7, 0x11, 0x52, 0xbd, 0x7f, 0x6a, 0xf0,
	0x6c, 0x57, 0x5d, 0xc4, 0xc2, 0x6f, 0xf0, 0x4c, 0x5a, 0xa7, 0xcd, 0xca, 0x6f, 0x67, 0x28, 0xa9,
	0x6f, 0x98, 0xaf, 0xf1, 0x7a, 0x3b, 0x4b, 0xd4, 0x57, 0x59, 0x24, 0x1f, 0x42, 0x53, 0x89, 0x65,
	0xb4, 0xc0, 0xbc, 0x1b, 0x4a, 0x2c, 0x91, 0x7c, 0x0e, 0x6d, 0xef, 0x5d, 0x58, 0x8b, 0xce, 0xcb,
	0xcd, 0x88, 0xe0, 0xb5, 0xc7, 0x7c, 0x06, 0xa5, 0x51, 0x55, 0x6d, 0x4c, 0xf4, 0x38, 0x72, 0x57,
	0x15, 0x2a, 0xf9, 0x0c, 0x1e, 0xa1, 0x58, 0x5f, 0x11, 0xeb, 0xd8, 0xbc, 0x88, 0xdb, 0xd2, 0x09,
	0xf0, 0xb8, 0x44, 0x7b, 0xef, 0xea, 0x70, 0xbc, 0x23, 0x55, 0xbf, 0x6b, 0x65, 0xcc, 0xb0, 0x0c,
	0x03, 0xdc, 0xb5, 0x88, 0x11, 0xbf, 0x13, 0x15, 0xed, 0x98, 0xdc, 0xe5, 0x86, 0x76, 0x4c, 0xd0,
	0xdf, 0x26, 0x7f, 0x46, 0x26, 0x39, 0x9b, 0xda, 0xf4, 0x45, 0xbc, 0x4d, 0x2b, 0x95, 0xbd, 0xf2,
	0x40, 0x72, 0x09, 0x4f, 0xec, 0x92, 0x4a, 0x65, 0x1d, 0xcb, 0x73, 0x2a, 0x95, 0x17, 0x30, 0x35,
	0xc2, 0xda, 0xf4, 0x4b, 0x6c, 0x8f, 0x5d, 0x0e, 0x91, 0x1c, 0xaa, 0x37, 0x91, 0x4a, 0x66, 0xd0,
	0x28, 0x2f, 0x6e, 0xfa, 0x55, 0xf7, 0xf0, 0xa2, 0x35, 0xf8, 0xf9, 0x9e, 0xeb, 0xb1, 0xfb, 0x8c,
	0x93, 0xfa, 0xad, 0xe4, 0x43, 0x7f, 0x0a, 0x34, 0x3c, 0xda, 0xba, 0xc9, 0xe9, 0xd7, 0x61, 0x1f,
	0x5f, 0xdd, 0x33, 0xe0, 0x96, 0x37, 0xd2, 0x29, 0x81, 0x78, 0x05, 0x06, 0x70, 0x1a, 0x6a, 0x49,
	0x9d, 0x61, 0xca, 0xca, 0xb0, 0x37, 0x61, 0xa8, 0xbf, 0xe9, 0x1e, 0xfa, 0x5e, 0x07, 0x72, 0xbc,
	0xe6, 0xc2, 0x80, 0xff, 0x0e, 0xe0, 0xef, 0x58, 0xd4, 0xf7, 0x6d, 0xd0, 0xf7, 0xfa, 0x9e, 0xfa,
	0x76, 0x1d, 0x44, 0xd2, 0x54, 0xb6, 0xd4, 0x97, 0x43, 0xab, 0x42, 0xa7, 0x2f, 0xf7, 0x1f, 0x0c,
	0x3c, 0x12, 0xa3, 0xfd, 0x09, 0xad, 0xca, 0x29, 0x4a, 0xbf, 0x0b, 0xbd, 0x7e, 0xb3, 0x97, 0x5e,
	0x57, 0xfc, 0x12, 0xc0, 0xc7, 0x95, 0x3f, 0x77, 0x4b, 0x68, 0x6f, 0x6c, 0x79, 0xfa, 0x7d, 0x08,
	0x4a, 0xf6, 0x12, 0x74, 0xc3, 0x33, 0x39, 0x0a, 0xcf, 0x9f, 0xf0, 0xf5, 0xf6, 0x61, 0xf8, 0xb9,
	0x72, 0xf9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xd1, 0x10, 0x58, 0xcb, 0x08, 0x00, 0x00,
}
