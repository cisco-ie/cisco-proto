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
// source: bgp_updgen_subgrp_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_update_generation_sub_groups_update_generation_sub_group

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

type BgpUpdgenSubgrpBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	UpdateGroupIndex     uint32   `protobuf:"varint,4,opt,name=update_group_index,json=updateGroupIndex,proto3" json:"update_group_index,omitempty"`
	SubGroupIndex        uint32   `protobuf:"varint,5,opt,name=sub_group_index,json=subGroupIndex,proto3" json:"sub_group_index,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,6,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpUpdgenSubgrpBag_KEYS) Reset()         { *m = BgpUpdgenSubgrpBag_KEYS{} }
func (m *BgpUpdgenSubgrpBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenSubgrpBag_KEYS) ProtoMessage()    {}
func (*BgpUpdgenSubgrpBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ba9998b0fc000c, []int{0}
}

func (m *BgpUpdgenSubgrpBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS.Unmarshal(m, b)
}
func (m *BgpUpdgenSubgrpBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenSubgrpBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS.Merge(m, src)
}
func (m *BgpUpdgenSubgrpBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS.Size(m)
}
func (m *BgpUpdgenSubgrpBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenSubgrpBag_KEYS proto.InternalMessageInfo

func (m *BgpUpdgenSubgrpBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetUpdateGroupIndex() uint32 {
	if m != nil {
		return m.UpdateGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetSubGroupIndex() uint32 {
	if m != nil {
		return m.SubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

type BgpTimespec struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          uint32   `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpTimespec) Reset()         { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()    {}
func (*BgpTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ba9998b0fc000c, []int{1}
}

func (m *BgpTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpTimespec.Unmarshal(m, b)
}
func (m *BgpTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpTimespec.Marshal(b, m, deterministic)
}
func (m *BgpTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpTimespec.Merge(m, src)
}
func (m *BgpTimespec) XXX_Size() int {
	return xxx_messageInfo_BgpTimespec.Size(m)
}
func (m *BgpTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_BgpTimespec proto.InternalMessageInfo

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type BgpUpdgenStatsBag struct {
	UpdateOutQueueMessagesHigh              uint32       `protobuf:"varint,1,opt,name=update_out_queue_messages_high,json=updateOutQueueMessagesHigh,proto3" json:"update_out_queue_messages_high,omitempty"`
	UpdateOutQueueMessagesCumulative        uint32       `protobuf:"varint,2,opt,name=update_out_queue_messages_cumulative,json=updateOutQueueMessagesCumulative,proto3" json:"update_out_queue_messages_cumulative,omitempty"`
	UpdateOutQueueMessagesDiscarded         uint32       `protobuf:"varint,3,opt,name=update_out_queue_messages_discarded,json=updateOutQueueMessagesDiscarded,proto3" json:"update_out_queue_messages_discarded,omitempty"`
	UpdateOutQueueMessagesCleared           uint32       `protobuf:"varint,4,opt,name=update_out_queue_messages_cleared,json=updateOutQueueMessagesCleared,proto3" json:"update_out_queue_messages_cleared,omitempty"`
	UpdateOutQueueSizeHigh                  uint32       `protobuf:"varint,5,opt,name=update_out_queue_size_high,json=updateOutQueueSizeHigh,proto3" json:"update_out_queue_size_high,omitempty"`
	UpdateOutQueueSizeCumulative            uint64       `protobuf:"varint,6,opt,name=update_out_queue_size_cumulative,json=updateOutQueueSizeCumulative,proto3" json:"update_out_queue_size_cumulative,omitempty"`
	UpdateOutQueueSizeDiscarded             uint64       `protobuf:"varint,7,opt,name=update_out_queue_size_discarded,json=updateOutQueueSizeDiscarded,proto3" json:"update_out_queue_size_discarded,omitempty"`
	UpdateOutQueueSizeCleared               uint64       `protobuf:"varint,8,opt,name=update_out_queue_size_cleared,json=updateOutQueueSizeCleared,proto3" json:"update_out_queue_size_cleared,omitempty"`
	LastUpdateDiscardTimestamp              *BgpTimespec `protobuf:"bytes,9,opt,name=last_update_discard_timestamp,json=lastUpdateDiscardTimestamp,proto3" json:"last_update_discard_timestamp,omitempty"`
	LastUpdateDiscardAge                    uint32       `protobuf:"varint,10,opt,name=last_update_discard_age,json=lastUpdateDiscardAge,proto3" json:"last_update_discard_age,omitempty"`
	LastUpdateClearedTimestamp              *BgpTimespec `protobuf:"bytes,11,opt,name=last_update_cleared_timestamp,json=lastUpdateClearedTimestamp,proto3" json:"last_update_cleared_timestamp,omitempty"`
	LastUpdateCleardAge                     uint32       `protobuf:"varint,12,opt,name=last_update_cleard_age,json=lastUpdateCleardAge,proto3" json:"last_update_cleard_age,omitempty"`
	UpdateThrottleCount                     uint32       `protobuf:"varint,13,opt,name=update_throttle_count,json=updateThrottleCount,proto3" json:"update_throttle_count,omitempty"`
	LastUpdateThrottleTimestamp             *BgpTimespec `protobuf:"bytes,14,opt,name=last_update_throttle_timestamp,json=lastUpdateThrottleTimestamp,proto3" json:"last_update_throttle_timestamp,omitempty"`
	LastUpdateThrottleAge                   uint32       `protobuf:"varint,15,opt,name=last_update_throttle_age,json=lastUpdateThrottleAge,proto3" json:"last_update_throttle_age,omitempty"`
	UpdateRecoveryCount                     uint32       `protobuf:"varint,16,opt,name=update_recovery_count,json=updateRecoveryCount,proto3" json:"update_recovery_count,omitempty"`
	LastUpdateRecoveryTimestamp             *BgpTimespec `protobuf:"bytes,17,opt,name=last_update_recovery_timestamp,json=lastUpdateRecoveryTimestamp,proto3" json:"last_update_recovery_timestamp,omitempty"`
	LastUpdateRecoveryAge                   uint32       `protobuf:"varint,18,opt,name=last_update_recovery_age,json=lastUpdateRecoveryAge,proto3" json:"last_update_recovery_age,omitempty"`
	UpdateMemoryAllocationFailCount         uint32       `protobuf:"varint,19,opt,name=update_memory_allocation_fail_count,json=updateMemoryAllocationFailCount,proto3" json:"update_memory_allocation_fail_count,omitempty"`
	LastUpdateMemoryAllocationFailTimestamp *BgpTimespec `protobuf:"bytes,20,opt,name=last_update_memory_allocation_fail_timestamp,json=lastUpdateMemoryAllocationFailTimestamp,proto3" json:"last_update_memory_allocation_fail_timestamp,omitempty"`
	LastUpdateMemoryAllocationFailAge       uint32       `protobuf:"varint,21,opt,name=last_update_memory_allocation_fail_age,json=lastUpdateMemoryAllocationFailAge,proto3" json:"last_update_memory_allocation_fail_age,omitempty"`
	XXX_NoUnkeyedLiteral                    struct{}     `json:"-"`
	XXX_unrecognized                        []byte       `json:"-"`
	XXX_sizecache                           int32        `json:"-"`
}

func (m *BgpUpdgenStatsBag) Reset()         { *m = BgpUpdgenStatsBag{} }
func (m *BgpUpdgenStatsBag) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenStatsBag) ProtoMessage()    {}
func (*BgpUpdgenStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ba9998b0fc000c, []int{2}
}

func (m *BgpUpdgenStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenStatsBag.Unmarshal(m, b)
}
func (m *BgpUpdgenStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenStatsBag.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenStatsBag.Merge(m, src)
}
func (m *BgpUpdgenStatsBag) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenStatsBag.Size(m)
}
func (m *BgpUpdgenStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenStatsBag proto.InternalMessageInfo

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCumulative() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesDiscarded() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueMessagesCleared() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessagesCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeHigh() uint32 {
	if m != nil {
		return m.UpdateOutQueueSizeHigh
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCumulative() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCumulative
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeDiscarded() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeDiscarded
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateOutQueueSizeCleared() uint64 {
	if m != nil {
		return m.UpdateOutQueueSizeCleared
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateDiscardTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateDiscardAge() uint32 {
	if m != nil {
		return m.LastUpdateDiscardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateClearedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateClearedTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateCleardAge() uint32 {
	if m != nil {
		return m.LastUpdateCleardAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateThrottleCount() uint32 {
	if m != nil {
		return m.UpdateThrottleCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateThrottleTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateThrottleAge() uint32 {
	if m != nil {
		return m.LastUpdateThrottleAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateRecoveryCount() uint32 {
	if m != nil {
		return m.UpdateRecoveryCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateRecoveryTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateRecoveryAge() uint32 {
	if m != nil {
		return m.LastUpdateRecoveryAge
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetUpdateMemoryAllocationFailCount() uint32 {
	if m != nil {
		return m.UpdateMemoryAllocationFailCount
	}
	return 0
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailTimestamp
	}
	return nil
}

func (m *BgpUpdgenStatsBag) GetLastUpdateMemoryAllocationFailAge() uint32 {
	if m != nil {
		return m.LastUpdateMemoryAllocationFailAge
	}
	return 0
}

type BgpUpdgenSubgrpBag struct {
	ProcessId                     uint32             `protobuf:"varint,50,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	UpdateVrfName                 string             `protobuf:"bytes,51,opt,name=update_vrf_name,json=updateVrfName,proto3" json:"update_vrf_name,omitempty"`
	UpdateGroupAfName             string             `protobuf:"bytes,52,opt,name=update_group_af_name,json=updateGroupAfName,proto3" json:"update_group_af_name,omitempty"`
	SubGroupIndexXr               uint32             `protobuf:"varint,53,opt,name=sub_group_index_xr,json=subGroupIndexXr,proto3" json:"sub_group_index_xr,omitempty"`
	SubGroupIdXr                  uint32             `protobuf:"varint,54,opt,name=sub_group_id_xr,json=subGroupIdXr,proto3" json:"sub_group_id_xr,omitempty"`
	ParentSubGroupIndex           uint32             `protobuf:"varint,55,opt,name=parent_sub_group_index,json=parentSubGroupIndex,proto3" json:"parent_sub_group_index,omitempty"`
	ParentSubGroupId              uint32             `protobuf:"varint,56,opt,name=parent_sub_group_id,json=parentSubGroupId,proto3" json:"parent_sub_group_id,omitempty"`
	UpdateGroupIndexXr            uint32             `protobuf:"varint,57,opt,name=update_group_index_xr,json=updateGroupIndexXr,proto3" json:"update_group_index_xr,omitempty"`
	UpdateMainTableVersion        uint32             `protobuf:"varint,58,opt,name=update_main_table_version,json=updateMainTableVersion,proto3" json:"update_main_table_version,omitempty"`
	UpdateVrfTableRibVersion      uint32             `protobuf:"varint,59,opt,name=update_vrf_table_rib_version,json=updateVrfTableRibVersion,proto3" json:"update_vrf_table_rib_version,omitempty"`
	CurrentUpdateLimitSubGroup    uint32             `protobuf:"varint,60,opt,name=current_update_limit_sub_group,json=currentUpdateLimitSubGroup,proto3" json:"current_update_limit_sub_group,omitempty"`
	ConfiguredUpdateLimitSubGroup uint32             `protobuf:"varint,61,opt,name=configured_update_limit_sub_group,json=configuredUpdateLimitSubGroup,proto3" json:"configured_update_limit_sub_group,omitempty"`
	UpdateOutQueueMessages        uint32             `protobuf:"varint,62,opt,name=update_out_queue_messages,json=updateOutQueueMessages,proto3" json:"update_out_queue_messages,omitempty"`
	UpdateOutQueueSize            uint32             `protobuf:"varint,63,opt,name=update_out_queue_size,json=updateOutQueueSize,proto3" json:"update_out_queue_size,omitempty"`
	UpdateThrottled               bool               `protobuf:"varint,64,opt,name=update_throttled,json=updateThrottled,proto3" json:"update_throttled,omitempty"`
	RefreshSubGroupCount          uint32             `protobuf:"varint,65,opt,name=refresh_sub_group_count,json=refreshSubGroupCount,proto3" json:"refresh_sub_group_count,omitempty"`
	FilterGroupCount              uint32             `protobuf:"varint,66,opt,name=filter_group_count,json=filterGroupCount,proto3" json:"filter_group_count,omitempty"`
	NeighborCount                 uint32             `protobuf:"varint,67,opt,name=neighbor_count,json=neighborCount,proto3" json:"neighbor_count,omitempty"`
	Version                       uint32             `protobuf:"varint,68,opt,name=version,proto3" json:"version,omitempty"`
	NsrVersion                    uint32             `protobuf:"varint,69,opt,name=nsr_version,json=nsrVersion,proto3" json:"nsr_version,omitempty"`
	PendingTargetVersion          uint32             `protobuf:"varint,70,opt,name=pending_target_version,json=pendingTargetVersion,proto3" json:"pending_target_version,omitempty"`
	NextResumeVersion             uint32             `protobuf:"varint,71,opt,name=next_resume_version,json=nextResumeVersion,proto3" json:"next_resume_version,omitempty"`
	UpdateRefreshTargetVersion    uint32             `protobuf:"varint,72,opt,name=update_refresh_target_version,json=updateRefreshTargetVersion,proto3" json:"update_refresh_target_version,omitempty"`
	IoWriteEventPending           bool               `protobuf:"varint,73,opt,name=io_write_event_pending,json=ioWriteEventPending,proto3" json:"io_write_event_pending,omitempty"`
	UpdateDefaultRouteSent        bool               `protobuf:"varint,74,opt,name=update_default_route_sent,json=updateDefaultRouteSent,proto3" json:"update_default_route_sent,omitempty"`
	CreationTimestamp             *BgpTimespec       `protobuf:"bytes,75,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	MergeCount                    uint32             `protobuf:"varint,76,opt,name=merge_count,json=mergeCount,proto3" json:"merge_count,omitempty"`
	LastMergeTimestamp            *BgpTimespec       `protobuf:"bytes,77,opt,name=last_merge_timestamp,json=lastMergeTimestamp,proto3" json:"last_merge_timestamp,omitempty"`
	LastMergedSubGroupIndex       uint32             `protobuf:"varint,78,opt,name=last_merged_sub_group_index,json=lastMergedSubGroupIndex,proto3" json:"last_merged_sub_group_index,omitempty"`
	EoRAttempted                  bool               `protobuf:"varint,79,opt,name=eo_r_attempted,json=eoRAttempted,proto3" json:"eo_r_attempted,omitempty"`
	FirstUpdateWalkStartTimestamp *BgpTimespec       `protobuf:"bytes,80,opt,name=first_update_walk_start_timestamp,json=firstUpdateWalkStartTimestamp,proto3" json:"first_update_walk_start_timestamp,omitempty"`
	FirstUpdateWalkEndTimestamp   *BgpTimespec       `protobuf:"bytes,81,opt,name=first_update_walk_end_timestamp,json=firstUpdateWalkEndTimestamp,proto3" json:"first_update_walk_end_timestamp,omitempty"`
	LastUpdateWalkStartTimestamp  *BgpTimespec       `protobuf:"bytes,82,opt,name=last_update_walk_start_timestamp,json=lastUpdateWalkStartTimestamp,proto3" json:"last_update_walk_start_timestamp,omitempty"`
	LastUpdateWalkEndTimestamp    *BgpTimespec       `protobuf:"bytes,83,opt,name=last_update_walk_end_timestamp,json=lastUpdateWalkEndTimestamp,proto3" json:"last_update_walk_end_timestamp,omitempty"`
	LastUpdateWalkEndAge          uint32             `protobuf:"varint,84,opt,name=last_update_walk_end_age,json=lastUpdateWalkEndAge,proto3" json:"last_update_walk_end_age,omitempty"`
	LastUpdateQueuedTimestamp     *BgpTimespec       `protobuf:"bytes,85,opt,name=last_update_queued_timestamp,json=lastUpdateQueuedTimestamp,proto3" json:"last_update_queued_timestamp,omitempty"`
	LastUpdateQueuedAge           uint32             `protobuf:"varint,86,opt,name=last_update_queued_age,json=lastUpdateQueuedAge,proto3" json:"last_update_queued_age,omitempty"`
	UpdateStatistics              *BgpUpdgenStatsBag `protobuf:"bytes,87,opt,name=update_statistics,json=updateStatistics,proto3" json:"update_statistics,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}           `json:"-"`
	XXX_unrecognized              []byte             `json:"-"`
	XXX_sizecache                 int32              `json:"-"`
}

func (m *BgpUpdgenSubgrpBag) Reset()         { *m = BgpUpdgenSubgrpBag{} }
func (m *BgpUpdgenSubgrpBag) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenSubgrpBag) ProtoMessage()    {}
func (*BgpUpdgenSubgrpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ba9998b0fc000c, []int{3}
}

func (m *BgpUpdgenSubgrpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenSubgrpBag.Unmarshal(m, b)
}
func (m *BgpUpdgenSubgrpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenSubgrpBag.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenSubgrpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenSubgrpBag.Merge(m, src)
}
func (m *BgpUpdgenSubgrpBag) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenSubgrpBag.Size(m)
}
func (m *BgpUpdgenSubgrpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenSubgrpBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenSubgrpBag proto.InternalMessageInfo

func (m *BgpUpdgenSubgrpBag) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateVrfName() string {
	if m != nil {
		return m.UpdateVrfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag) GetUpdateGroupAfName() string {
	if m != nil {
		return m.UpdateGroupAfName
	}
	return ""
}

func (m *BgpUpdgenSubgrpBag) GetSubGroupIndexXr() uint32 {
	if m != nil {
		return m.SubGroupIndexXr
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetSubGroupIdXr() uint32 {
	if m != nil {
		return m.SubGroupIdXr
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetParentSubGroupIndex() uint32 {
	if m != nil {
		return m.ParentSubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetParentSubGroupId() uint32 {
	if m != nil {
		return m.ParentSubGroupId
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateGroupIndexXr() uint32 {
	if m != nil {
		return m.UpdateGroupIndexXr
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateMainTableVersion() uint32 {
	if m != nil {
		return m.UpdateMainTableVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateVrfTableRibVersion() uint32 {
	if m != nil {
		return m.UpdateVrfTableRibVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetCurrentUpdateLimitSubGroup() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitSubGroup
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetConfiguredUpdateLimitSubGroup() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitSubGroup
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateThrottled() bool {
	if m != nil {
		return m.UpdateThrottled
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetFilterGroupCount() uint32 {
	if m != nil {
		return m.FilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNsrVersion() uint32 {
	if m != nil {
		return m.NsrVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetPendingTargetVersion() uint32 {
	if m != nil {
		return m.PendingTargetVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetNextResumeVersion() uint32 {
	if m != nil {
		return m.NextResumeVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateRefreshTargetVersion() uint32 {
	if m != nil {
		return m.UpdateRefreshTargetVersion
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetIoWriteEventPending() bool {
	if m != nil {
		return m.IoWriteEventPending
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetUpdateDefaultRouteSent() bool {
	if m != nil {
		return m.UpdateDefaultRouteSent
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetCreationTimestamp() *BgpTimespec {
	if m != nil {
		return m.CreationTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetMergeCount() uint32 {
	if m != nil {
		return m.MergeCount
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetLastMergeTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastMergeTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastMergedSubGroupIndex() uint32 {
	if m != nil {
		return m.LastMergedSubGroupIndex
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetEoRAttempted() bool {
	if m != nil {
		return m.EoRAttempted
	}
	return false
}

func (m *BgpUpdgenSubgrpBag) GetFirstUpdateWalkStartTimestamp() *BgpTimespec {
	if m != nil {
		return m.FirstUpdateWalkStartTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetFirstUpdateWalkEndTimestamp() *BgpTimespec {
	if m != nil {
		return m.FirstUpdateWalkEndTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkStartTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateWalkStartTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkEndTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateWalkEndTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateWalkEndAge() uint32 {
	if m != nil {
		return m.LastUpdateWalkEndAge
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateQueuedTimestamp() *BgpTimespec {
	if m != nil {
		return m.LastUpdateQueuedTimestamp
	}
	return nil
}

func (m *BgpUpdgenSubgrpBag) GetLastUpdateQueuedAge() uint32 {
	if m != nil {
		return m.LastUpdateQueuedAge
	}
	return 0
}

func (m *BgpUpdgenSubgrpBag) GetUpdateStatistics() *BgpUpdgenStatsBag {
	if m != nil {
		return m.UpdateStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpUpdgenSubgrpBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_subgrp_bag_KEYS")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_stats_bag")
	proto.RegisterType((*BgpUpdgenSubgrpBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.update_generation_sub_groups.update_generation_sub_group.bgp_updgen_subgrp_bag")
}

func init() { proto.RegisterFile("bgp_updgen_subgrp_bag.proto", fileDescriptor_c6ba9998b0fc000c) }

var fileDescriptor_c6ba9998b0fc000c = []byte{
	// 1511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0x49, 0x73, 0x1b, 0x37,
	0x16, 0xae, 0xf6, 0x78, 0x2c, 0xfb, 0x49, 0xb2, 0x2c, 0x68, 0x71, 0x6b, 0xb3, 0x68, 0x79, 0x19,
	0x4d, 0x8d, 0x87, 0x53, 0x63, 0x79, 0x89, 0x1d, 0xc7, 0xb1, 0x6c, 0xc9, 0xbb, 0xbc, 0x50, 0xf2,
	0x92, 0x13, 0x0a, 0xec, 0x06, 0x5b, 0x28, 0x93, 0xdd, 0x1d, 0x00, 0x4d, 0x2b, 0xfe, 0x09, 0xf9,
	0x05, 0x39, 0xe6, 0x27, 0xe4, 0x2f, 0xe4, 0x96, 0xaa, 0x54, 0xe5, 0x07, 0xa4, 0xb2, 0x54, 0x2a,
	0xe7, 0xdc, 0x73, 0x4c, 0x61, 0xe9, 0x6e, 0x90, 0x6c, 0xba, 0x72, 0xa4, 0x2f, 0x2c, 0x12, 0xef,
	0x7b, 0x0f, 0xef, 0x7d, 0x0f, 0x04, 0x3e, 0x00, 0x96, 0x9a, 0x51, 0x8a, 0xb3, 0x34, 0x8c, 0x68,
	0x8c, 0x45, 0xd6, 0x8c, 0x78, 0x8a, 0x9b, 0x24, 0xaa, 0xa7, 0x3c, 0x91, 0x09, 0xfa, 0xd2, 0x0b,
	0x98, 0x08, 0x12, 0xcc, 0x12, 0x81, 0x0f, 0x38, 0x66, 0x69, 0xf7, 0x12, 0x56, 0xf8, 0x24, 0xa5,
	0xbc, 0xde, 0x8c, 0xd2, 0x3a, 0x8b, 0x85, 0x24, 0x71, 0x40, 0x45, 0xf1, 0xad, 0xf8, 0x82, 0x49,
	0x20, 0x59, 0x97, 0xd6, 0xbb, 0xbc, 0x25, 0xd4, 0x47, 0x9d, 0xb4, 0x44, 0x9d, 0xb4, 0xea, 0x59,
	0x1a, 0x12, 0x49, 0x71, 0x44, 0x63, 0xca, 0x89, 0x64, 0x89, 0x9e, 0x16, 0x47, 0x3c, 0xc9, 0x52,
	0xf1, 0x3e, 0xe3, 0xda, 0x1f, 0x1e, 0x2c, 0x56, 0x26, 0x8b, 0x1f, 0x6d, 0x7f, 0xb6, 0x8b, 0xce,
	0xc0, 0x64, 0x31, 0x75, 0x4c, 0x3a, 0xd4, 0xf7, 0x6a, 0xde, 0xfa, 0xb1, 0xc6, 0x44, 0x3e, 0xf8,
	0x84, 0x74, 0x28, 0x5a, 0x80, 0xa3, 0x5d, 0xde, 0x32, 0xf6, 0x43, 0xda, 0x3e, 0xd6, 0xe5, 0x2d,
	0x6d, 0x3a, 0x09, 0x63, 0xc4, 0x5a, 0xfe, 0xa1, 0x2d, 0x47, 0x88, 0x31, 0x5c, 0x00, 0x94, 0xa7,
	0xa5, 0xf2, 0xc0, 0x2c, 0x0e, 0xe9, 0x81, 0x7f, 0xb8, 0xe6, 0xad, 0x4f, 0x36, 0x4e, 0x18, 0xcb,
	0x3d, 0x65, 0x78, 0xa0, 0xc6, 0xd1, 0x79, 0x98, 0x2a, 0x52, 0xb6, 0xd0, 0x7f, 0x6a, 0xe8, 0xa4,
	0xc8, 0x9a, 0x0e, 0xae, 0x06, 0x13, 0x0e, 0x2e, 0xf4, 0x8f, 0x68, 0x10, 0x14, 0xa0, 0x70, 0xed,
	0x21, 0x4c, 0xa8, 0x72, 0x25, 0xeb, 0x50, 0x91, 0xd2, 0x00, 0xf9, 0x30, 0x26, 0x68, 0x90, 0xc4,
	0xa1, 0xd0, 0xa5, 0x4d, 0x36, 0xf2, 0x9f, 0xa8, 0x06, 0xe3, 0x31, 0x89, 0x93, 0xdc, 0x7a, 0x48,
	0x5b, 0xdd, 0xa1, 0xb5, 0x6f, 0xa6, 0x61, 0xd6, 0xe5, 0x4e, 0x12, 0x29, 0x14, 0x75, 0xe8, 0x36,
	0x9c, 0xb2, 0xc5, 0x25, 0x99, 0xc4, 0x9f, 0x67, 0x34, 0xa3, 0xb8, 0x43, 0x85, 0x20, 0x11, 0x15,
	0x78, 0x9f, 0x45, 0xfb, 0x76, 0xae, 0x45, 0x83, 0x7a, 0x9a, 0xc9, 0xe7, 0x0a, 0xb3, 0x63, 0x21,
	0xf7, 0x59, 0xb4, 0x8f, 0x9e, 0xc0, 0xd9, 0xe1, 0x31, 0x82, 0xac, 0x93, 0xb5, 0x89, 0x5a, 0x09,
	0x36, 0xaf, 0x5a, 0x75, 0xa4, 0x3b, 0x05, 0x0e, 0x3d, 0x86, 0x33, 0xc3, 0xe3, 0x85, 0x4c, 0x04,
	0x84, 0x87, 0x34, 0xd4, 0x5d, 0x9a, 0x6c, 0xac, 0x56, 0x87, 0xdb, 0xca, 0x61, 0xe8, 0x3e, 0x9c,
	0x7e, 0x4f, 0x76, 0x6d, 0x4a, 0x38, 0x0d, 0x6d, 0x37, 0x57, 0x86, 0xa4, 0x66, 0x40, 0xe8, 0x3a,
	0x2c, 0x0e, 0x44, 0x12, 0xec, 0x1d, 0x35, 0x3c, 0x99, 0x2e, 0xcf, 0xf7, 0x86, 0xd8, 0x65, 0xef,
	0xa8, 0xe6, 0xe8, 0x2e, 0xd4, 0xaa, 0x7d, 0x1d, 0x7e, 0xd4, 0x12, 0x38, 0xdc, 0x58, 0x1e, 0x8c,
	0xe0, 0x70, 0xb3, 0x05, 0xab, 0xd5, 0x71, 0x4a, 0x5e, 0xc6, 0x74, 0x98, 0xa5, 0xc1, 0x30, 0x25,
	0x27, 0xb7, 0x60, 0x65, 0x48, 0x36, 0x96, 0x8f, 0xa3, 0x3a, 0xc6, 0x42, 0x45, 0x2a, 0x96, 0x8b,
	0x9f, 0x3c, 0x58, 0x69, 0x13, 0x21, 0xb1, 0x8d, 0x63, 0xa7, 0x37, 0xab, 0x55, 0x92, 0x4e, 0xea,
	0x1f, 0xab, 0x79, 0xeb, 0xe3, 0x17, 0xbf, 0xf2, 0xea, 0xa3, 0xb3, 0x85, 0xd4, 0xdd, 0xff, 0x53,
	0x63, 0x51, 0xe5, 0xff, 0x42, 0xa3, 0x2d, 0x33, 0x7b, 0x79, 0xf2, 0xe8, 0x32, 0x9c, 0xac, 0xaa,
	0x8e, 0x44, 0xd4, 0x07, 0xdd, 0xe7, 0xd9, 0x01, 0xe7, 0xcd, 0x88, 0x0e, 0xb0, 0x62, 0xe9, 0x74,
	0x58, 0x19, 0xff, 0x80, 0x58, 0xb1, 0xbd, 0x2e, 0x59, 0xd9, 0x80, 0xf9, 0x81, 0xea, 0x0c, 0x29,
	0x13, 0x9a, 0x94, 0x99, 0x3e, 0x5f, 0xcd, 0xc9, 0x45, 0x98, 0xb3, 0x78, 0xb9, 0xcf, 0x13, 0x29,
	0xdb, 0x14, 0x07, 0x49, 0x16, 0x4b, 0x7f, 0xd2, 0xf8, 0x18, 0xe3, 0x9e, 0xb5, 0xdd, 0x51, 0x26,
	0xf4, 0x8b, 0x07, 0xa7, 0xdc, 0x99, 0x0a, 0xcf, 0x92, 0xc8, 0xe3, 0xa3, 0x4e, 0xe4, 0x52, 0x49,
	0x46, 0x5e, 0x5c, 0xc9, 0xe4, 0x55, 0xf0, 0x2b, 0xeb, 0x53, 0x5c, 0x4e, 0x69, 0x5e, 0xe6, 0x06,
	0xdd, 0x7b, 0xd9, 0xe4, 0x34, 0x48, 0xba, 0x94, 0x7f, 0x61, 0xd9, 0x3c, 0xe1, 0xb2, 0xd9, 0xb0,
	0xb6, 0x6a, 0x36, 0x0b, 0xcf, 0x92, 0xcd, 0xe9, 0x0f, 0x88, 0xcd, 0xbc, 0xb8, 0xa1, 0x6c, 0x16,
	0xf5, 0x29, 0x36, 0x51, 0x3f, 0x9b, 0xb9, 0xbb, 0x62, 0xb3, 0x3c, 0x69, 0x3a, 0xb4, 0x93, 0x28,
	0x8f, 0x76, 0x3b, 0x09, 0x4c, 0x22, 0x2d, 0xc2, 0xda, 0x96, 0xdb, 0x19, 0xf7, 0xa4, 0xd9, 0xd1,
	0xc8, 0xcd, 0x02, 0x78, 0x97, 0xb0, 0xb6, 0xe1, 0xf9, 0x4f, 0x0f, 0x2e, 0xb8, 0x79, 0x0c, 0x89,
	0x59, 0xb2, 0x3e, 0x3b, 0xea, 0xac, 0xff, 0xab, 0xa4, 0xad, 0xaa, 0xec, 0xb2, 0x03, 0xcf, 0xe1,
	0xfc, 0xdf, 0xa8, 0x5c, 0xf5, 0x63, 0x4e, 0x73, 0x79, 0xfa, 0xfd, 0x81, 0x37, 0x23, 0xba, 0xf6,
	0xed, 0x02, 0xcc, 0x55, 0xca, 0x3d, 0xb4, 0x02, 0x90, 0xf2, 0x24, 0xa0, 0x42, 0x28, 0xe1, 0x74,
	0x51, 0x07, 0x3c, 0x66, 0x47, 0x1e, 0x84, 0x4a, 0x81, 0xd9, 0x34, 0x0a, 0xa9, 0xb7, 0xa1, 0x05,
	0xdd, 0xa4, 0x19, 0x7e, 0x69, 0x05, 0xdf, 0xff, 0x60, 0xb6, 0x47, 0xd7, 0xe5, 0xea, 0xef, 0x92,
	0x06, 0x4f, 0x3b, 0xca, 0x6e, 0xd3, 0x38, 0xfc, 0x07, 0x50, 0x9f, 0xb4, 0xc3, 0x07, 0xdc, 0xbf,
	0xac, 0xe7, 0x9f, 0xea, 0x51, 0x77, 0xaf, 0x39, 0x3a, 0xd7, 0xa3, 0x03, 0x43, 0x85, 0xbc, 0xa2,
	0x91, 0x13, 0xa5, 0xc4, 0x7b, 0xcd, 0xd5, 0x96, 0x9a, 0x12, 0x4e, 0x63, 0x89, 0xfb, 0x55, 0xe3,
	0x55, 0xf3, 0x87, 0x36, 0xd6, 0xdd, 0x1e, 0xed, 0xf8, 0x5f, 0x98, 0x19, 0x74, 0x0a, 0xfd, 0x8f,
	0x8c, 0x24, 0xed, 0xf3, 0x08, 0xd1, 0xff, 0x8b, 0x3d, 0xa3, 0x2f, 0xf5, 0x6b, 0xda, 0x01, 0xf5,
	0x6b, 0xd8, 0xd7, 0x1c, 0x5d, 0x83, 0x85, 0xbc, 0x95, 0x84, 0xc5, 0x58, 0x92, 0x66, 0x9b, 0xe2,
	0x2e, 0xe5, 0x82, 0x25, 0xb1, 0x7f, 0xdd, 0x55, 0x3a, 0x3b, 0x84, 0xc5, 0x7b, 0xca, 0xfc, 0xd2,
	0x58, 0xd1, 0x4d, 0x58, 0x76, 0xe8, 0x37, 0x9e, 0x9c, 0x35, 0x0b, 0xef, 0x8f, 0xb5, 0xb7, 0x5f,
	0xf4, 0x42, 0x3b, 0x37, 0x58, 0x33, 0xf7, 0xbf, 0x0d, 0xa7, 0x82, 0x8c, 0xeb, 0xea, 0x6c, 0x9c,
	0x36, 0xeb, 0x30, 0xa7, 0x54, 0xff, 0x86, 0x51, 0xa4, 0x16, 0x65, 0x56, 0xd1, 0x63, 0x85, 0xc9,
	0x6b, 0x56, 0x9a, 0x2f, 0x48, 0xe2, 0x16, 0x8b, 0x32, 0x75, 0xfa, 0x0e, 0x09, 0xf3, 0x89, 0xd1,
	0x7c, 0x25, 0xb0, 0x2a, 0x52, 0x49, 0xc4, 0xa0, 0x7a, 0xf4, 0x6f, 0x56, 0x49, 0xbe, 0x5c, 0x35,
	0x3a, 0xb4, 0xf7, 0x8a, 0x2c, 0xff, 0x53, 0x97, 0x76, 0x57, 0x5c, 0xa1, 0x7f, 0xc3, 0x89, 0xbe,
	0x13, 0x21, 0xf4, 0x6f, 0xd5, 0xbc, 0xf5, 0xa3, 0x8d, 0xa9, 0xde, 0x63, 0x32, 0x54, 0x0a, 0x85,
	0xd3, 0x16, 0xa7, 0x62, 0xdf, 0x59, 0x04, 0x66, 0xbb, 0xda, 0x34, 0x0a, 0xc5, 0x9a, 0xf3, 0x52,
	0xcc, 0x1e, 0x75, 0x01, 0x50, 0x8b, 0xb5, 0x25, 0xe5, 0x3d, 0x1e, 0xb7, 0xcd, 0xca, 0x31, 0x16,
	0x07, 0x7d, 0x0e, 0x8e, 0xc7, 0x94, 0x45, 0xfb, 0xcd, 0x84, 0x5b, 0xe4, 0x1d, 0x73, 0x97, 0xc9,
	0x47, 0x0d, 0xcc, 0x87, 0xb1, 0xbc, 0xbb, 0x5b, 0xe6, 0x66, 0x62, 0x7f, 0xa2, 0x55, 0x18, 0x8f,
	0x05, 0x2f, 0x7a, 0xbf, 0x6d, 0x2e, 0x39, 0xb1, 0xe0, 0x79, 0xb7, 0x2f, 0xc1, 0x7c, 0x4a, 0xe3,
	0x90, 0xc5, 0x11, 0x96, 0x84, 0x47, 0x54, 0x16, 0xd8, 0xbb, 0xa6, 0x0a, 0x6b, 0xdd, 0xd3, 0xc6,
	0xdc, 0xab, 0x0e, 0x33, 0x31, 0x3d, 0x90, 0x98, 0x53, 0x91, 0x75, 0xca, 0x85, 0x79, 0x4f, 0xbb,
	0x4c, 0x2b, 0x53, 0x43, 0x5b, 0x72, 0xfc, 0x66, 0xa1, 0x77, 0x73, 0xce, 0xfa, 0x26, 0xbb, 0xef,
	0x5e, 0x72, 0x1a, 0x06, 0xd3, 0x3b, 0xe5, 0x06, 0xcc, 0xb3, 0x04, 0xbf, 0xe5, 0x4c, 0x52, 0x4c,
	0xbb, 0x6a, 0x75, 0xda, 0xcc, 0xfc, 0x07, 0xba, 0x41, 0x33, 0x2c, 0x79, 0xa5, 0x8c, 0xdb, 0xca,
	0xf6, 0xcc, 0x98, 0x9c, 0xd5, 0x13, 0xd2, 0x16, 0xc9, 0xda, 0x12, 0xf3, 0x24, 0x93, 0x14, 0x0b,
	0x1a, 0x4b, 0xff, 0xa1, 0xf6, 0xb3, 0xab, 0x67, 0xcb, 0xd8, 0x1b, 0xca, 0xbc, 0x4b, 0x63, 0x89,
	0xbe, 0xf3, 0x00, 0x05, 0x9c, 0x9a, 0xdd, 0xb3, 0x3c, 0x32, 0x1e, 0x8d, 0xfa, 0x91, 0x31, 0x9d,
	0x27, 0x5d, 0x1e, 0x0e, 0xab, 0x30, 0xde, 0xa1, 0x3c, 0xca, 0x75, 0xdf, 0x63, 0xb3, 0x08, 0xf4,
	0x90, 0x59, 0x3f, 0xdf, 0x7b, 0xa0, 0xf5, 0x34, 0x36, 0xb0, 0xb2, 0xda, 0x9d, 0x51, 0xaf, 0x16,
	0xa9, 0xb4, 0x77, 0x54, 0xd6, 0x65, 0xb9, 0x37, 0x60, 0xa9, 0x2c, 0x26, 0x1c, 0xd8, 0xd7, 0x9f,
	0xe8, 0xf2, 0x4f, 0x16, 0x8e, 0x61, 0xef, 0xde, 0x7e, 0x16, 0x8e, 0xd3, 0x04, 0x73, 0x4c, 0xa4,
	0xa4, 0x9d, 0x54, 0xd2, 0xd0, 0x7f, 0xaa, 0xd7, 0xc9, 0x04, 0x4d, 0x1a, 0x9b, 0xf9, 0x18, 0xfa,
	0xdd, 0x83, 0xd3, 0x2d, 0xc6, 0xcb, 0x13, 0xf7, 0x2d, 0x69, 0xbf, 0x51, 0xd7, 0x7a, 0x2e, 0x1d,
	0xfa, 0x9e, 0x8d, 0x3a, 0x7d, 0x2b, 0xba, 0x06, 0xb3, 0xef, 0xbe, 0x22, 0xed, 0x37, 0xbb, 0xaa,
	0x80, 0x92, 0xc9, 0x5f, 0x3d, 0x58, 0x1d, 0xac, 0x92, 0xc6, 0xee, 0x85, 0xea, 0xf9, 0xc8, 0x2b,
	0xd7, 0xbe, 0x1a, 0xb7, 0x63, 0xe7, 0x46, 0xf5, 0x9b, 0x07, 0x35, 0x57, 0x38, 0x55, 0xb6, 0xb1,
	0x31, 0xea, 0x25, 0x2e, 0x97, 0x6a, 0xae, 0xa2, 0x8b, 0x3f, 0xf7, 0x5d, 0x3f, 0x2a, 0x9a, 0xb8,
	0xfb, 0x01, 0xdd, 0x8a, 0x07, 0x7a, 0x78, 0xa5, 0xf7, 0xf6, 0x51, 0x94, 0xa7, 0xd4, 0xee, 0x5e,
	0xff, 0x63, 0x81, 0xf5, 0x56, 0x97, 0x8f, 0x1f, 0x3d, 0x58, 0x76, 0x1d, 0xb5, 0x42, 0x70, 0x59,
	0x79, 0x31, 0xea, 0xac, 0x2c, 0x94, 0x75, 0x69, 0x0d, 0x33, 0xfc, 0xa9, 0xc0, 0xd6, 0xa6, 0x28,
	0x79, 0xd9, 0xff, 0x54, 0x60, 0x5c, 0x15, 0x23, 0x3f, 0x78, 0x60, 0x65, 0xb7, 0x7e, 0xa1, 0x64,
	0x42, 0xb2, 0x40, 0xf8, 0xaf, 0x34, 0x0d, 0x5f, 0x8f, 0x1c, 0x0d, 0xfd, 0x6f, 0xa9, 0xf9, 0x63,
	0xf0, 0x6e, 0x91, 0x7a, 0xf3, 0x88, 0x7e, 0x46, 0xdf, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x50,
	0x21, 0xb9, 0xd0, 0x65, 0x17, 0x00, 0x00,
}
