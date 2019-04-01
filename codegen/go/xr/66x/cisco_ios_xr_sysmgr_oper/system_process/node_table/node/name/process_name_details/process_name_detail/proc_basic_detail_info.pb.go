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
// source: proc_basic_detail_info.proto

package cisco_ios_xr_sysmgr_oper_system_process_node_table_node_name_process_name_details_process_name_detail

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

type ProcBasicDetailInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProcName             string   `protobuf:"bytes,2,opt,name=proc_name,json=procName,proto3" json:"proc_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcBasicDetailInfo_KEYS) Reset()         { *m = ProcBasicDetailInfo_KEYS{} }
func (m *ProcBasicDetailInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ProcBasicDetailInfo_KEYS) ProtoMessage()    {}
func (*ProcBasicDetailInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{0}
}

func (m *ProcBasicDetailInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBasicDetailInfo_KEYS.Unmarshal(m, b)
}
func (m *ProcBasicDetailInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBasicDetailInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ProcBasicDetailInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBasicDetailInfo_KEYS.Merge(m, src)
}
func (m *ProcBasicDetailInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ProcBasicDetailInfo_KEYS.Size(m)
}
func (m *ProcBasicDetailInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBasicDetailInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBasicDetailInfo_KEYS proto.InternalMessageInfo

func (m *ProcBasicDetailInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ProcBasicDetailInfo_KEYS) GetProcName() string {
	if m != nil {
		return m.ProcName
	}
	return ""
}

type TupleInfo struct {
	Tuple                string   `protobuf:"bytes,1,opt,name=tuple,proto3" json:"tuple,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TupleInfo) Reset()         { *m = TupleInfo{} }
func (m *TupleInfo) String() string { return proto.CompactTextString(m) }
func (*TupleInfo) ProtoMessage()    {}
func (*TupleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{1}
}

func (m *TupleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TupleInfo.Unmarshal(m, b)
}
func (m *TupleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TupleInfo.Marshal(b, m, deterministic)
}
func (m *TupleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TupleInfo.Merge(m, src)
}
func (m *TupleInfo) XXX_Size() int {
	return xxx_messageInfo_TupleInfo.Size(m)
}
func (m *TupleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TupleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TupleInfo proto.InternalMessageInfo

func (m *TupleInfo) GetTuple() string {
	if m != nil {
		return m.Tuple
	}
	return ""
}

type ProcCpuTime struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	System               string   `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
	Total                string   `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcCpuTime) Reset()         { *m = ProcCpuTime{} }
func (m *ProcCpuTime) String() string { return proto.CompactTextString(m) }
func (*ProcCpuTime) ProtoMessage()    {}
func (*ProcCpuTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{2}
}

func (m *ProcCpuTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcCpuTime.Unmarshal(m, b)
}
func (m *ProcCpuTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcCpuTime.Marshal(b, m, deterministic)
}
func (m *ProcCpuTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcCpuTime.Merge(m, src)
}
func (m *ProcCpuTime) XXX_Size() int {
	return xxx_messageInfo_ProcCpuTime.Size(m)
}
func (m *ProcCpuTime) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcCpuTime.DiscardUnknown(m)
}

var xxx_messageInfo_ProcCpuTime proto.InternalMessageInfo

func (m *ProcCpuTime) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ProcCpuTime) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *ProcCpuTime) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

type ProcBasicInfo struct {
	JobIdXr              uint32       `protobuf:"varint,1,opt,name=job_id_xr,json=jobIdXr,proto3" json:"job_id_xr,omitempty"`
	ProcessId            uint32       `protobuf:"varint,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	ProcessName          string       `protobuf:"bytes,3,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Executable           string       `protobuf:"bytes,4,opt,name=executable,proto3" json:"executable,omitempty"`
	ActivePath           string       `protobuf:"bytes,5,opt,name=active_path,json=activePath,proto3" json:"active_path,omitempty"`
	InstanceId           int32        `protobuf:"zigzag32,6,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Args                 string       `protobuf:"bytes,7,opt,name=args,proto3" json:"args,omitempty"`
	VersionId            string       `protobuf:"bytes,8,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	Respawn              string       `protobuf:"bytes,9,opt,name=respawn,proto3" json:"respawn,omitempty"`
	RespawnCount         int32        `protobuf:"zigzag32,10,opt,name=respawn_count,json=respawnCount,proto3" json:"respawn_count,omitempty"`
	LastStarted          string       `protobuf:"bytes,11,opt,name=last_started,json=lastStarted,proto3" json:"last_started,omitempty"`
	ProcessState         string       `protobuf:"bytes,12,opt,name=process_state,json=processState,proto3" json:"process_state,omitempty"`
	LastExitStatus       int32        `protobuf:"zigzag32,13,opt,name=last_exit_status,json=lastExitStatus,proto3" json:"last_exit_status,omitempty"`
	LastExitReason       string       `protobuf:"bytes,14,opt,name=last_exit_reason,json=lastExitReason,proto3" json:"last_exit_reason,omitempty"`
	PackageState         string       `protobuf:"bytes,15,opt,name=package_state,json=packageState,proto3" json:"package_state,omitempty"`
	StartedOnConfig      string       `protobuf:"bytes,16,opt,name=started_on_config,json=startedOnConfig,proto3" json:"started_on_config,omitempty"`
	RegisteredItem       []*TupleInfo `protobuf:"bytes,17,rep,name=registered_item,json=registeredItem,proto3" json:"registered_item,omitempty"`
	FeatureName          string       `protobuf:"bytes,18,opt,name=feature_name,json=featureName,proto3" json:"feature_name,omitempty"`
	Tag                  string       `protobuf:"bytes,19,opt,name=tag,proto3" json:"tag,omitempty"`
	Group                string       `protobuf:"bytes,20,opt,name=group,proto3" json:"group,omitempty"`
	Core                 string       `protobuf:"bytes,21,opt,name=core,proto3" json:"core,omitempty"`
	MaxCore              int32        `protobuf:"zigzag32,22,opt,name=max_core,json=maxCore,proto3" json:"max_core,omitempty"`
	Level                string       `protobuf:"bytes,23,opt,name=level,proto3" json:"level,omitempty"`
	Mandatory            bool         `protobuf:"varint,24,opt,name=mandatory,proto3" json:"mandatory,omitempty"`
	MaintModeProc        bool         `protobuf:"varint,25,opt,name=maint_mode_proc,json=maintModeProc,proto3" json:"maint_mode_proc,omitempty"`
	PlacementState       string       `protobuf:"bytes,26,opt,name=placement_state,json=placementState,proto3" json:"placement_state,omitempty"`
	StartUpPath          string       `protobuf:"bytes,27,opt,name=start_up_path,json=startUpPath,proto3" json:"start_up_path,omitempty"`
	MemoryLimit          uint32       `protobuf:"varint,28,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	Ready                string       `protobuf:"bytes,29,opt,name=ready,proto3" json:"ready,omitempty"`
	Available            string       `protobuf:"bytes,30,opt,name=available,proto3" json:"available,omitempty"`
	ProcCpuTime          *ProcCpuTime `protobuf:"bytes,31,opt,name=proc_cpu_time,json=procCpuTime,proto3" json:"proc_cpu_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProcBasicInfo) Reset()         { *m = ProcBasicInfo{} }
func (m *ProcBasicInfo) String() string { return proto.CompactTextString(m) }
func (*ProcBasicInfo) ProtoMessage()    {}
func (*ProcBasicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{3}
}

func (m *ProcBasicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBasicInfo.Unmarshal(m, b)
}
func (m *ProcBasicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBasicInfo.Marshal(b, m, deterministic)
}
func (m *ProcBasicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBasicInfo.Merge(m, src)
}
func (m *ProcBasicInfo) XXX_Size() int {
	return xxx_messageInfo_ProcBasicInfo.Size(m)
}
func (m *ProcBasicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBasicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBasicInfo proto.InternalMessageInfo

func (m *ProcBasicInfo) GetJobIdXr() uint32 {
	if m != nil {
		return m.JobIdXr
	}
	return 0
}

func (m *ProcBasicInfo) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *ProcBasicInfo) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *ProcBasicInfo) GetExecutable() string {
	if m != nil {
		return m.Executable
	}
	return ""
}

func (m *ProcBasicInfo) GetActivePath() string {
	if m != nil {
		return m.ActivePath
	}
	return ""
}

func (m *ProcBasicInfo) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *ProcBasicInfo) GetArgs() string {
	if m != nil {
		return m.Args
	}
	return ""
}

func (m *ProcBasicInfo) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

func (m *ProcBasicInfo) GetRespawn() string {
	if m != nil {
		return m.Respawn
	}
	return ""
}

func (m *ProcBasicInfo) GetRespawnCount() int32 {
	if m != nil {
		return m.RespawnCount
	}
	return 0
}

func (m *ProcBasicInfo) GetLastStarted() string {
	if m != nil {
		return m.LastStarted
	}
	return ""
}

func (m *ProcBasicInfo) GetProcessState() string {
	if m != nil {
		return m.ProcessState
	}
	return ""
}

func (m *ProcBasicInfo) GetLastExitStatus() int32 {
	if m != nil {
		return m.LastExitStatus
	}
	return 0
}

func (m *ProcBasicInfo) GetLastExitReason() string {
	if m != nil {
		return m.LastExitReason
	}
	return ""
}

func (m *ProcBasicInfo) GetPackageState() string {
	if m != nil {
		return m.PackageState
	}
	return ""
}

func (m *ProcBasicInfo) GetStartedOnConfig() string {
	if m != nil {
		return m.StartedOnConfig
	}
	return ""
}

func (m *ProcBasicInfo) GetRegisteredItem() []*TupleInfo {
	if m != nil {
		return m.RegisteredItem
	}
	return nil
}

func (m *ProcBasicInfo) GetFeatureName() string {
	if m != nil {
		return m.FeatureName
	}
	return ""
}

func (m *ProcBasicInfo) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *ProcBasicInfo) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *ProcBasicInfo) GetCore() string {
	if m != nil {
		return m.Core
	}
	return ""
}

func (m *ProcBasicInfo) GetMaxCore() int32 {
	if m != nil {
		return m.MaxCore
	}
	return 0
}

func (m *ProcBasicInfo) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *ProcBasicInfo) GetMandatory() bool {
	if m != nil {
		return m.Mandatory
	}
	return false
}

func (m *ProcBasicInfo) GetMaintModeProc() bool {
	if m != nil {
		return m.MaintModeProc
	}
	return false
}

func (m *ProcBasicInfo) GetPlacementState() string {
	if m != nil {
		return m.PlacementState
	}
	return ""
}

func (m *ProcBasicInfo) GetStartUpPath() string {
	if m != nil {
		return m.StartUpPath
	}
	return ""
}

func (m *ProcBasicInfo) GetMemoryLimit() uint32 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *ProcBasicInfo) GetReady() string {
	if m != nil {
		return m.Ready
	}
	return ""
}

func (m *ProcBasicInfo) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

func (m *ProcBasicInfo) GetProcCpuTime() *ProcCpuTime {
	if m != nil {
		return m.ProcCpuTime
	}
	return nil
}

type ProcDetailInfo struct {
	RunningPath            string   `protobuf:"bytes,1,opt,name=running_path,json=runningPath,proto3" json:"running_path,omitempty"`
	PackagePath            string   `protobuf:"bytes,2,opt,name=package_path,json=packagePath,proto3" json:"package_path,omitempty"`
	JobIdLink              int32    `protobuf:"zigzag32,3,opt,name=job_id_link,json=jobIdLink,proto3" json:"job_id_link,omitempty"`
	GroupJid               string   `protobuf:"bytes,4,opt,name=group_jid,json=groupJid,proto3" json:"group_jid,omitempty"`
	FailCount              uint32   `protobuf:"varint,5,opt,name=fail_count,json=failCount,proto3" json:"fail_count,omitempty"`
	RestartNeeded          bool     `protobuf:"varint,6,opt,name=restart_needed,json=restartNeeded,proto3" json:"restart_needed,omitempty"`
	InitProcess            bool     `protobuf:"varint,7,opt,name=init_process,json=initProcess,proto3" json:"init_process,omitempty"`
	LastOnline             string   `protobuf:"bytes,8,opt,name=last_online,json=lastOnline,proto3" json:"last_online,omitempty"`
	ThisPcb                string   `protobuf:"bytes,9,opt,name=this_pcb,json=thisPcb,proto3" json:"this_pcb,omitempty"`
	NextPcb                string   `protobuf:"bytes,10,opt,name=next_pcb,json=nextPcb,proto3" json:"next_pcb,omitempty"`
	Envs                   string   `protobuf:"bytes,11,opt,name=envs,proto3" json:"envs,omitempty"`
	WaitFor                string   `protobuf:"bytes,12,opt,name=wait_for,json=waitFor,proto3" json:"wait_for,omitempty"`
	JobIdOnRp              int32    `protobuf:"zigzag32,13,opt,name=job_id_on_rp,json=jobIdOnRp,proto3" json:"job_id_on_rp,omitempty"`
	IsStandbyCapable       bool     `protobuf:"varint,14,opt,name=is_standby_capable,json=isStandbyCapable,proto3" json:"is_standby_capable,omitempty"`
	DisableKill            bool     `protobuf:"varint,15,opt,name=disable_kill,json=disableKill,proto3" json:"disable_kill,omitempty"`
	SendAvail              bool     `protobuf:"varint,16,opt,name=send_avail,json=sendAvail,proto3" json:"send_avail,omitempty"`
	NodeEventCliInfo       int32    `protobuf:"zigzag32,17,opt,name=node_event_cli_info,json=nodeEventCliInfo,proto3" json:"node_event_cli_info,omitempty"`
	NodeRedundancyState    string   `protobuf:"bytes,18,opt,name=node_redundancy_state,json=nodeRedundancyState,proto3" json:"node_redundancy_state,omitempty"`
	RoleEventCliInfo       int32    `protobuf:"zigzag32,19,opt,name=role_event_cli_info,json=roleEventCliInfo,proto3" json:"role_event_cli_info,omitempty"`
	ProcRoleState          string   `protobuf:"bytes,20,opt,name=proc_role_state,json=procRoleState,proto3" json:"proc_role_state,omitempty"`
	StandbyEventCliInfo    int32    `protobuf:"zigzag32,21,opt,name=standby_event_cli_info,json=standbyEventCliInfo,proto3" json:"standby_event_cli_info,omitempty"`
	CleanupEventCliInfo    int32    `protobuf:"zigzag32,22,opt,name=cleanup_event_cli_info,json=cleanupEventCliInfo,proto3" json:"cleanup_event_cli_info,omitempty"`
	BandReadyEventCliInfo  int32    `protobuf:"zigzag32,23,opt,name=band_ready_event_cli_info,json=bandReadyEventCliInfo,proto3" json:"band_ready_event_cli_info,omitempty"`
	LrEventCliInfo         int32    `protobuf:"zigzag32,24,opt,name=lr_event_cli_info,json=lrEventCliInfo,proto3" json:"lr_event_cli_info,omitempty"`
	PlaneReadyEventCliInfo int32    `protobuf:"zigzag32,25,opt,name=plane_ready_event_cli_info,json=planeReadyEventCliInfo,proto3" json:"plane_ready_event_cli_info,omitempty"`
	MdrIsDoneCliInfo       int32    `protobuf:"zigzag32,26,opt,name=mdr_is_done_cli_info,json=mdrIsDoneCliInfo,proto3" json:"mdr_is_done_cli_info,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *ProcDetailInfo) Reset()         { *m = ProcDetailInfo{} }
func (m *ProcDetailInfo) String() string { return proto.CompactTextString(m) }
func (*ProcDetailInfo) ProtoMessage()    {}
func (*ProcDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{4}
}

func (m *ProcDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcDetailInfo.Unmarshal(m, b)
}
func (m *ProcDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcDetailInfo.Marshal(b, m, deterministic)
}
func (m *ProcDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcDetailInfo.Merge(m, src)
}
func (m *ProcDetailInfo) XXX_Size() int {
	return xxx_messageInfo_ProcDetailInfo.Size(m)
}
func (m *ProcDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcDetailInfo proto.InternalMessageInfo

func (m *ProcDetailInfo) GetRunningPath() string {
	if m != nil {
		return m.RunningPath
	}
	return ""
}

func (m *ProcDetailInfo) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

func (m *ProcDetailInfo) GetJobIdLink() int32 {
	if m != nil {
		return m.JobIdLink
	}
	return 0
}

func (m *ProcDetailInfo) GetGroupJid() string {
	if m != nil {
		return m.GroupJid
	}
	return ""
}

func (m *ProcDetailInfo) GetFailCount() uint32 {
	if m != nil {
		return m.FailCount
	}
	return 0
}

func (m *ProcDetailInfo) GetRestartNeeded() bool {
	if m != nil {
		return m.RestartNeeded
	}
	return false
}

func (m *ProcDetailInfo) GetInitProcess() bool {
	if m != nil {
		return m.InitProcess
	}
	return false
}

func (m *ProcDetailInfo) GetLastOnline() string {
	if m != nil {
		return m.LastOnline
	}
	return ""
}

func (m *ProcDetailInfo) GetThisPcb() string {
	if m != nil {
		return m.ThisPcb
	}
	return ""
}

func (m *ProcDetailInfo) GetNextPcb() string {
	if m != nil {
		return m.NextPcb
	}
	return ""
}

func (m *ProcDetailInfo) GetEnvs() string {
	if m != nil {
		return m.Envs
	}
	return ""
}

func (m *ProcDetailInfo) GetWaitFor() string {
	if m != nil {
		return m.WaitFor
	}
	return ""
}

func (m *ProcDetailInfo) GetJobIdOnRp() int32 {
	if m != nil {
		return m.JobIdOnRp
	}
	return 0
}

func (m *ProcDetailInfo) GetIsStandbyCapable() bool {
	if m != nil {
		return m.IsStandbyCapable
	}
	return false
}

func (m *ProcDetailInfo) GetDisableKill() bool {
	if m != nil {
		return m.DisableKill
	}
	return false
}

func (m *ProcDetailInfo) GetSendAvail() bool {
	if m != nil {
		return m.SendAvail
	}
	return false
}

func (m *ProcDetailInfo) GetNodeEventCliInfo() int32 {
	if m != nil {
		return m.NodeEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetNodeRedundancyState() string {
	if m != nil {
		return m.NodeRedundancyState
	}
	return ""
}

func (m *ProcDetailInfo) GetRoleEventCliInfo() int32 {
	if m != nil {
		return m.RoleEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetProcRoleState() string {
	if m != nil {
		return m.ProcRoleState
	}
	return ""
}

func (m *ProcDetailInfo) GetStandbyEventCliInfo() int32 {
	if m != nil {
		return m.StandbyEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetCleanupEventCliInfo() int32 {
	if m != nil {
		return m.CleanupEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetBandReadyEventCliInfo() int32 {
	if m != nil {
		return m.BandReadyEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetLrEventCliInfo() int32 {
	if m != nil {
		return m.LrEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetPlaneReadyEventCliInfo() int32 {
	if m != nil {
		return m.PlaneReadyEventCliInfo
	}
	return 0
}

func (m *ProcDetailInfo) GetMdrIsDoneCliInfo() int32 {
	if m != nil {
		return m.MdrIsDoneCliInfo
	}
	return 0
}

type ProcBasicDetailInfo struct {
	BasicInfo            *ProcBasicInfo  `protobuf:"bytes,50,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	DetailInfo           *ProcDetailInfo `protobuf:"bytes,51,opt,name=detail_info,json=detailInfo,proto3" json:"detail_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ProcBasicDetailInfo) Reset()         { *m = ProcBasicDetailInfo{} }
func (m *ProcBasicDetailInfo) String() string { return proto.CompactTextString(m) }
func (*ProcBasicDetailInfo) ProtoMessage()    {}
func (*ProcBasicDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8b8fc8a2af706761, []int{5}
}

func (m *ProcBasicDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBasicDetailInfo.Unmarshal(m, b)
}
func (m *ProcBasicDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBasicDetailInfo.Marshal(b, m, deterministic)
}
func (m *ProcBasicDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBasicDetailInfo.Merge(m, src)
}
func (m *ProcBasicDetailInfo) XXX_Size() int {
	return xxx_messageInfo_ProcBasicDetailInfo.Size(m)
}
func (m *ProcBasicDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBasicDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBasicDetailInfo proto.InternalMessageInfo

func (m *ProcBasicDetailInfo) GetBasicInfo() *ProcBasicInfo {
	if m != nil {
		return m.BasicInfo
	}
	return nil
}

func (m *ProcBasicDetailInfo) GetDetailInfo() *ProcDetailInfo {
	if m != nil {
		return m.DetailInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ProcBasicDetailInfo_KEYS)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.proc_basic_detail_info_KEYS")
	proto.RegisterType((*TupleInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.tuple_info")
	proto.RegisterType((*ProcCpuTime)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.proc_cpu_time")
	proto.RegisterType((*ProcBasicInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.proc_basic_info")
	proto.RegisterType((*ProcDetailInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.proc_detail_info")
	proto.RegisterType((*ProcBasicDetailInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.name.process_name_details.process_name_detail.proc_basic_detail_info")
}

func init() { proto.RegisterFile("proc_basic_detail_info.proto", fileDescriptor_8b8fc8a2af706761) }

var fileDescriptor_8b8fc8a2af706761 = []byte{
	// 1240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0x86, 0x92, 0x38, 0x96, 0x46, 0x96, 0x7f, 0xd6, 0x3f, 0xa1, 0x9d, 0x3f, 0x47, 0x45, 0x53,
	0xb7, 0x68, 0x74, 0xb0, 0x2f, 0x45, 0x6f, 0x85, 0x9b, 0x02, 0x6a, 0xd2, 0xc4, 0xdd, 0xb4, 0x68,
	0x7b, 0x22, 0xa8, 0x25, 0xad, 0x30, 0xde, 0x25, 0xb7, 0x24, 0x57, 0x96, 0x5e, 0xa1, 0xf7, 0x02,
	0x41, 0x5f, 0xa9, 0x6f, 0xd1, 0x27, 0x29, 0x66, 0x48, 0xd9, 0x96, 0xe3, 0xb3, 0x6f, 0x9c, 0x6f,
	0xfe, 0x38, 0xc3, 0x99, 0x6f, 0x17, 0x1e, 0xd5, 0xce, 0x16, 0x7c, 0x24, 0xbc, 0x2e, 0xb8, 0x54,
	0x41, 0xe8, 0x92, 0x6b, 0x73, 0x6a, 0x07, 0xb5, 0xb3, 0xc1, 0x66, 0xaa, 0xd0, 0xbe, 0xb0, 0x5c,
	0x5b, 0xcf, 0xa7, 0x8e, 0xfb, 0x99, 0xaf, 0xc6, 0x8e, 0xdb, 0x5a, 0xb9, 0x81, 0x9f, 0xf9, 0xa0,
	0x2a, 0x8e, 0xde, 0xca, 0xfb, 0x81, 0xb1, 0x52, 0xf1, 0x20, 0x46, 0xa5, 0xa2, 0xe3, 0xc0, 0x88,
	0x4a, 0x0d, 0x92, 0x92, 0xa3, 0x90, 0x82, 0xfb, 0x9b, 0xc0, 0xfe, 0x6f, 0xf0, 0xf0, 0xe6, 0x6b,
	0xf0, 0x57, 0x2f, 0xff, 0x78, 0x97, 0x3d, 0x84, 0x0e, 0xc5, 0x47, 0x17, 0xd6, 0xda, 0x6f, 0x1d,
	0x74, 0xf2, 0x36, 0x02, 0x6f, 0x44, 0xa5, 0x50, 0x49, 0xbe, 0xa4, 0xbc, 0x13, 0x95, 0x08, 0xa0,
	0xb2, 0xdf, 0x07, 0x08, 0x4d, 0x5d, 0x2a, 0x0a, 0x96, 0x6d, 0xc1, 0x12, 0x49, 0x29, 0x46, 0x14,
	0xfa, 0x3f, 0x43, 0x8f, 0x02, 0x14, 0x75, 0xc3, 0x83, 0xae, 0x54, 0x96, 0xc1, 0xbd, 0xc6, 0x2b,
	0x97, 0xac, 0xe8, 0x9c, 0xed, 0xc0, 0xfd, 0x58, 0x71, 0x4a, 0x91, 0x24, 0x0a, 0x69, 0x83, 0x28,
	0xd9, 0xdd, 0x14, 0x12, 0x85, 0xfe, 0x7f, 0x1d, 0x58, 0xbb, 0x52, 0x10, 0x25, 0xdf, 0x83, 0xce,
	0x07, 0x3b, 0xe2, 0x5a, 0xf2, 0x69, 0x0c, 0xdd, 0xcb, 0x97, 0x3f, 0xd8, 0xd1, 0x50, 0xfe, 0xee,
	0xb2, 0xc7, 0x00, 0xf3, 0xb6, 0x68, 0x49, 0x19, 0x7a, 0x79, 0x27, 0x21, 0x43, 0x99, 0x3d, 0x83,
	0x95, 0xab, 0x5d, 0x4b, 0xb9, 0xba, 0x09, 0xa3, 0x2e, 0x3c, 0x01, 0x50, 0x53, 0x55, 0x34, 0xf4,
	0x04, 0xec, 0x1e, 0x19, 0x5c, 0x41, 0xb2, 0xa7, 0xd0, 0x15, 0x45, 0xd0, 0x13, 0xc5, 0x6b, 0x11,
	0xde, 0xb3, 0xa5, 0x68, 0x10, 0xa1, 0x13, 0x11, 0xde, 0xa3, 0x81, 0x36, 0x3e, 0x08, 0x53, 0x28,
	0xbc, 0xc3, 0xfd, 0xfd, 0xd6, 0xc1, 0x46, 0x0e, 0x73, 0x68, 0x28, 0xb1, 0x2b, 0xc2, 0x8d, 0x3d,
	0x5b, 0x8e, 0x5d, 0xc1, 0x33, 0xde, 0x7b, 0xa2, 0x9c, 0xd7, 0xd6, 0xa0, 0x4f, 0x9b, 0x34, 0x9d,
	0x84, 0x0c, 0x65, 0xc6, 0x60, 0xd9, 0x29, 0x5f, 0x8b, 0x73, 0xc3, 0x3a, 0xa4, 0x9b, 0x8b, 0xd9,
	0x67, 0xd0, 0x4b, 0x47, 0x5e, 0xd8, 0xc6, 0x04, 0x06, 0x94, 0x6f, 0x25, 0x81, 0xc7, 0x88, 0x61,
	0xd9, 0xa5, 0xf0, 0x81, 0xfb, 0x20, 0x5c, 0x50, 0x92, 0x75, 0x63, 0xd9, 0x88, 0xbd, 0x8b, 0x10,
	0xc6, 0x99, 0x77, 0xc6, 0x07, 0x11, 0x14, 0x5b, 0x21, 0x9b, 0x79, 0xbb, 0xde, 0x21, 0x96, 0x1d,
	0xc0, 0x3a, 0xc5, 0x51, 0x53, 0x4d, 0xc1, 0x42, 0xe3, 0x59, 0x8f, 0xf2, 0xad, 0x22, 0xfe, 0x72,
	0xaa, 0x31, 0x5e, 0x68, 0xfc, 0xa2, 0xa5, 0x53, 0xc2, 0x5b, 0xc3, 0x56, 0x29, 0xe2, 0x85, 0x65,
	0x4e, 0x28, 0x25, 0x16, 0xc5, 0x99, 0x18, 0xab, 0x94, 0x78, 0x2d, 0x25, 0x8e, 0x60, 0x4c, 0xfc,
	0x15, 0x6c, 0xa4, 0xbb, 0x73, 0x8b, 0x85, 0x9a, 0x53, 0x3d, 0x66, 0xeb, 0x64, 0xb8, 0x96, 0x14,
	0x6f, 0xcd, 0x31, 0xc1, 0xd9, 0x3f, 0x2d, 0x58, 0x73, 0x6a, 0xac, 0x7d, 0x50, 0x4e, 0x49, 0xae,
	0x71, 0xd4, 0x36, 0xf6, 0xef, 0x1e, 0x74, 0x0f, 0xff, 0x1c, 0xdc, 0xca, 0x12, 0x0e, 0x2e, 0x17,
	0x25, 0x5f, 0xbd, 0xbc, 0xc9, 0x10, 0xa7, 0xfc, 0x19, 0xac, 0x9c, 0x2a, 0x11, 0x1a, 0x97, 0x76,
	0x30, 0x8b, 0x2f, 0x91, 0x30, 0x1a, 0xc0, 0x75, 0xb8, 0x1b, 0xc4, 0x98, 0x6d, 0x92, 0x06, 0x8f,
	0xb8, 0x1a, 0x63, 0x67, 0x9b, 0x9a, 0x6d, 0xc5, 0xd5, 0x20, 0x01, 0xc7, 0xa8, 0xb0, 0x4e, 0xb1,
	0xed, 0x38, 0x46, 0x78, 0xce, 0x76, 0xa1, 0x5d, 0x89, 0x29, 0x27, 0x7c, 0x87, 0x1e, 0x66, 0xb9,
	0x12, 0xd3, 0x63, 0x54, 0x6d, 0xc1, 0x52, 0xa9, 0x26, 0xaa, 0x64, 0x0f, 0x62, 0x10, 0x12, 0xb2,
	0x47, 0xd0, 0xa9, 0x84, 0x91, 0x22, 0x58, 0x37, 0x63, 0x6c, 0xbf, 0x75, 0xd0, 0xce, 0x2f, 0x81,
	0xec, 0x39, 0xac, 0x55, 0x42, 0x9b, 0xc0, 0x2b, 0xec, 0x07, 0x96, 0xca, 0x76, 0xc9, 0xa6, 0x47,
	0xf0, 0x4f, 0x56, 0xaa, 0x13, 0x67, 0x8b, 0xec, 0x0b, 0x58, 0xab, 0x4b, 0x51, 0xa8, 0x4a, 0x99,
	0x90, 0x5e, 0x71, 0x2f, 0x3e, 0xf6, 0x05, 0x1c, 0xdf, 0xb1, 0x0f, 0x3d, 0x7a, 0x2e, 0xde, 0xd4,
	0x71, 0x7d, 0x1e, 0xc6, 0xfa, 0x09, 0xfc, 0xb5, 0xa6, 0xfd, 0x79, 0x06, 0x2b, 0x95, 0xaa, 0xac,
	0x9b, 0xf1, 0x52, 0x57, 0x3a, 0xb0, 0x47, 0xb4, 0xc4, 0xdd, 0x88, 0xbd, 0x46, 0x08, 0x6b, 0x71,
	0x4a, 0xc8, 0x19, 0x7b, 0x1c, 0x6b, 0x21, 0x01, 0x6b, 0x11, 0x13, 0xa1, 0x4b, 0x5a, 0xdc, 0x27,
	0x71, 0x85, 0x2e, 0x80, 0xec, 0x63, 0xeb, 0x1a, 0x3b, 0xb1, 0xa7, 0xfb, 0xad, 0x83, 0xee, 0x61,
	0xb8, 0xa5, 0xa1, 0x58, 0xc8, 0x1d, 0x29, 0xe7, 0xb8, 0x6e, 0x7e, 0xd1, 0x95, 0xea, 0xff, 0xd5,
	0x86, 0x75, 0x52, 0x5f, 0xe1, 0x6b, 0x6c, 0x83, 0x6b, 0x8c, 0xd1, 0x66, 0x1c, 0x3b, 0x15, 0x39,
	0xb4, 0x9b, 0xb0, 0x79, 0xa7, 0xe6, 0xab, 0x43, 0x26, 0x77, 0x12, 0x9b, 0x45, 0x8c, 0x4c, 0x9e,
	0x40, 0x37, 0x71, 0x65, 0xa9, 0xcd, 0x19, 0xf1, 0xdd, 0x46, 0xde, 0x21, 0xb6, 0x7c, 0xad, 0xcd,
	0x19, 0x72, 0x3e, 0x4d, 0x13, 0xff, 0xa0, 0x65, 0x22, 0xbb, 0x36, 0x01, 0x3f, 0x6a, 0x89, 0xa4,
	0x74, 0x8a, 0xf7, 0x89, 0xc4, 0xb2, 0x14, 0xc9, 0x14, 0x91, 0xc8, 0x2a, 0x9f, 0xc3, 0xaa, 0x53,
	0xf1, 0x39, 0x8d, 0x52, 0x52, 0x45, 0xae, 0x6b, 0xe7, 0xbd, 0x84, 0xbe, 0x21, 0x10, 0x6f, 0xa9,
	0x8d, 0x0e, 0xf3, 0x36, 0x12, 0xed, 0xb5, 0xf3, 0x2e, 0x62, 0x27, 0x11, 0x42, 0xca, 0x24, 0xb6,
	0xb0, 0xa6, 0xd4, 0x46, 0x25, 0xfa, 0x03, 0x84, 0xde, 0x12, 0x82, 0x73, 0x1d, 0xde, 0x6b, 0xcf,
	0xeb, 0x62, 0x34, 0x27, 0x40, 0x94, 0x4f, 0x8a, 0x11, 0xaa, 0x8c, 0x9a, 0x06, 0x52, 0x41, 0x54,
	0xa1, 0x8c, 0xaa, 0x0c, 0xee, 0x29, 0x33, 0xf1, 0x89, 0xee, 0xe8, 0x8c, 0xe6, 0xe7, 0x42, 0x07,
	0x7e, 0x6a, 0x5d, 0xa2, 0xb8, 0x65, 0x94, 0x7f, 0xb0, 0x2e, 0x7b, 0x0a, 0x2b, 0xa9, 0x57, 0xd6,
	0x70, 0x57, 0x27, 0x66, 0x8b, 0xcd, 0x7a, 0x6b, 0xf2, 0x3a, 0xfb, 0x1a, 0x32, 0x4d, 0xf4, 0x68,
	0xe4, 0x68, 0xc6, 0x0b, 0x51, 0xd3, 0xa4, 0xad, 0x52, 0x3d, 0xeb, 0x1a, 0x39, 0x12, 0x15, 0xc7,
	0x11, 0xc7, 0xba, 0xa5, 0xf6, 0x78, 0xe4, 0x67, 0xba, 0x2c, 0x89, 0xd7, 0xda, 0x79, 0x37, 0x61,
	0xaf, 0x74, 0x59, 0x62, 0x83, 0xbd, 0x32, 0x92, 0xd3, 0x94, 0x12, 0x9f, 0xb5, 0xf3, 0x0e, 0x22,
	0xdf, 0x21, 0x90, 0xbd, 0x80, 0x4d, 0x9a, 0x39, 0x35, 0xc1, 0xbd, 0x2a, 0x4a, 0x4d, 0x93, 0xc1,
	0x36, 0xe8, 0x5e, 0xeb, 0xa8, 0x7a, 0x89, 0x9a, 0xe3, 0x52, 0x0f, 0x71, 0x62, 0x0e, 0x61, 0x9b,
	0xcc, 0x9d, 0x92, 0x8d, 0x91, 0xc2, 0x14, 0xb3, 0xb4, 0x8b, 0x91, 0x64, 0x28, 0x56, 0x7e, 0xa1,
	0x8b, 0x0b, 0xf9, 0x02, 0x36, 0x9d, 0x2d, 0x3f, 0x49, 0xb1, 0x19, 0x53, 0xa0, 0x6a, 0x21, 0xc5,
	0xf3, 0xf4, 0x35, 0x26, 0x9f, 0x18, 0x3c, 0x72, 0x12, 0xad, 0x56, 0x6e, 0xcb, 0xc4, 0xd7, 0x47,
	0xb0, 0x33, 0x6f, 0xd3, 0xb5, 0xc8, 0xdb, 0x14, 0x79, 0x33, 0x69, 0x17, 0x82, 0x1f, 0xc1, 0x4e,
	0x51, 0x2a, 0x61, 0x9a, 0xfa, 0xba, 0x53, 0xa4, 0xb2, 0xcd, 0xa4, 0x5d, 0x70, 0xfa, 0x06, 0x76,
	0x47, 0xc2, 0x48, 0x4e, 0x14, 0x70, 0xdd, 0xef, 0x01, 0xf9, 0x6d, 0xa3, 0x41, 0x8e, 0xfa, 0x05,
	0xcf, 0x2f, 0x61, 0xa3, 0x74, 0xd7, 0x3d, 0x58, 0xfa, 0x9a, 0xb9, 0x05, 0xd3, 0x6f, 0x61, 0xaf,
	0x2e, 0x85, 0x51, 0x37, 0x67, 0xd9, 0x25, 0x9f, 0x1d, 0xb2, 0xf8, 0x34, 0xcd, 0x00, 0xb6, 0x2a,
	0xe9, 0xb8, 0xf6, 0x5c, 0x5a, 0xa3, 0x2e, 0xbd, 0xf6, 0x62, 0x8b, 0x2b, 0xe9, 0x86, 0xfe, 0x7b,
	0x6b, 0x54, 0xb2, 0xef, 0xff, 0x7b, 0x07, 0x76, 0x6e, 0xfe, 0x85, 0xcb, 0xfe, 0x6e, 0x01, 0x5c,
	0xfe, 0x07, 0xb1, 0x43, 0xe2, 0xaf, 0xc9, 0x6d, 0xf2, 0xd7, 0x65, 0xf6, 0xbc, 0x43, 0x67, 0x2a,
	0xf1, 0x63, 0x0b, 0xba, 0x57, 0xee, 0xc9, 0x8e, 0xe8, 0x62, 0xe7, 0xb7, 0x79, 0xb1, 0x2b, 0xe9,
	0x73, 0x88, 0x02, 0x5e, 0x6d, 0x74, 0x9f, 0xfe, 0xbe, 0x8f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x81, 0x25, 0xb9, 0x9d, 0x0b, 0x00, 0x00,
}
