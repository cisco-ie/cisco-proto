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
// source: bgp_updgen_af_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_update_generation_address_family

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

type BgpUpdgenAfBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpUpdgenAfBag_KEYS) Reset()         { *m = BgpUpdgenAfBag_KEYS{} }
func (m *BgpUpdgenAfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenAfBag_KEYS) ProtoMessage()    {}
func (*BgpUpdgenAfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0755445336bd1ac1, []int{0}
}

func (m *BgpUpdgenAfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenAfBag_KEYS.Unmarshal(m, b)
}
func (m *BgpUpdgenAfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenAfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenAfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenAfBag_KEYS.Merge(m, src)
}
func (m *BgpUpdgenAfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenAfBag_KEYS.Size(m)
}
func (m *BgpUpdgenAfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenAfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenAfBag_KEYS proto.InternalMessageInfo

func (m *BgpUpdgenAfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpUpdgenAfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpUpdgenAfBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
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
	return fileDescriptor_0755445336bd1ac1, []int{1}
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
	return fileDescriptor_0755445336bd1ac1, []int{2}
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

type BgpUpdgenAfBag struct {
	UpdateGroupAfName                  string             `protobuf:"bytes,50,opt,name=update_group_af_name,json=updateGroupAfName,proto3" json:"update_group_af_name,omitempty"`
	UpdateMainTableVersion             uint32             `protobuf:"varint,51,opt,name=update_main_table_version,json=updateMainTableVersion,proto3" json:"update_main_table_version,omitempty"`
	UpdateRibVersion                   uint32             `protobuf:"varint,52,opt,name=update_rib_version,json=updateRibVersion,proto3" json:"update_rib_version,omitempty"`
	UpdateMinNeighborVersion           uint32             `protobuf:"varint,53,opt,name=update_min_neighbor_version,json=updateMinNeighborVersion,proto3" json:"update_min_neighbor_version,omitempty"`
	CurrentUpdateLimitAf               uint32             `protobuf:"varint,54,opt,name=current_update_limit_af,json=currentUpdateLimitAf,proto3" json:"current_update_limit_af,omitempty"`
	ConfiguredUpdateLimitAf            uint32             `protobuf:"varint,55,opt,name=configured_update_limit_af,json=configuredUpdateLimitAf,proto3" json:"configured_update_limit_af,omitempty"`
	CurrentUpdateLimitSubgrpEbgp       uint32             `protobuf:"varint,56,opt,name=current_update_limit_subgrp_ebgp,json=currentUpdateLimitSubgrpEbgp,proto3" json:"current_update_limit_subgrp_ebgp,omitempty"`
	ConfiguredUpdateLimitSubgrpEbgp    uint32             `protobuf:"varint,57,opt,name=configured_update_limit_subgrp_ebgp,json=configuredUpdateLimitSubgrpEbgp,proto3" json:"configured_update_limit_subgrp_ebgp,omitempty"`
	CurrentUpdateLimitSubGroupIbgp     uint32             `protobuf:"varint,58,opt,name=current_update_limit_sub_group_ibgp,json=currentUpdateLimitSubGroupIbgp,proto3" json:"current_update_limit_sub_group_ibgp,omitempty"`
	ConfiguredUpdateLimitSubGroupIbgp  uint32             `protobuf:"varint,59,opt,name=configured_update_limit_sub_group_ibgp,json=configuredUpdateLimitSubGroupIbgp,proto3" json:"configured_update_limit_sub_group_ibgp,omitempty"`
	UpdateOutQueueMessages             uint32             `protobuf:"varint,60,opt,name=update_out_queue_messages,json=updateOutQueueMessages,proto3" json:"update_out_queue_messages,omitempty"`
	UpdateOutQueueSize                 uint32             `protobuf:"varint,61,opt,name=update_out_queue_size,json=updateOutQueueSize,proto3" json:"update_out_queue_size,omitempty"`
	UpdateThrottled                    bool               `protobuf:"varint,62,opt,name=update_throttled,json=updateThrottled,proto3" json:"update_throttled,omitempty"`
	UpdateUpdateGroupCount             uint32             `protobuf:"varint,63,opt,name=update_update_group_count,json=updateUpdateGroupCount,proto3" json:"update_update_group_count,omitempty"`
	UpdateSubGroupCount                uint32             `protobuf:"varint,64,opt,name=update_sub_group_count,json=updateSubGroupCount,proto3" json:"update_sub_group_count,omitempty"`
	SubGroupThrottledCount             uint32             `protobuf:"varint,65,opt,name=sub_group_throttled_count,json=subGroupThrottledCount,proto3" json:"sub_group_throttled_count,omitempty"`
	RefreshSubGroupCount               uint32             `protobuf:"varint,66,opt,name=refresh_sub_group_count,json=refreshSubGroupCount,proto3" json:"refresh_sub_group_count,omitempty"`
	RefreshSubGroupThrottledCount      uint32             `protobuf:"varint,67,opt,name=refresh_sub_group_throttled_count,json=refreshSubGroupThrottledCount,proto3" json:"refresh_sub_group_throttled_count,omitempty"`
	FilterGroupCount                   uint32             `protobuf:"varint,68,opt,name=filter_group_count,json=filterGroupCount,proto3" json:"filter_group_count,omitempty"`
	NeighborCount                      uint32             `protobuf:"varint,69,opt,name=neighbor_count,json=neighborCount,proto3" json:"neighbor_count,omitempty"`
	UpdateTableVrfName                 string             `protobuf:"bytes,70,opt,name=update_table_vrf_name,json=updateTableVrfName,proto3" json:"update_table_vrf_name,omitempty"`
	UpdateVrfafName                    uint32             `protobuf:"varint,71,opt,name=update_vrfaf_name,json=updateVrfafName,proto3" json:"update_vrfaf_name,omitempty"`
	UpdateVrfRdVersion                 uint32             `protobuf:"varint,72,opt,name=update_vrf_rd_version,json=updateVrfRdVersion,proto3" json:"update_vrf_rd_version,omitempty"`
	UpdateVrfTableRibVersion           uint32             `protobuf:"varint,73,opt,name=update_vrf_table_rib_version,json=updateVrfTableRibVersion,proto3" json:"update_vrf_table_rib_version,omitempty"`
	TableUpdateGroupCount              uint32             `protobuf:"varint,74,opt,name=table_update_group_count,json=tableUpdateGroupCount,proto3" json:"table_update_group_count,omitempty"`
	UpdateTableSubGroupCount           uint32             `protobuf:"varint,75,opt,name=update_table_sub_group_count,json=updateTableSubGroupCount,proto3" json:"update_table_sub_group_count,omitempty"`
	TableSubGroupThrottledCount        uint32             `protobuf:"varint,76,opt,name=table_sub_group_throttled_count,json=tableSubGroupThrottledCount,proto3" json:"table_sub_group_throttled_count,omitempty"`
	TableRefreshSubGroupCount          uint32             `protobuf:"varint,77,opt,name=table_refresh_sub_group_count,json=tableRefreshSubGroupCount,proto3" json:"table_refresh_sub_group_count,omitempty"`
	TableRefreshSubGroupThrottledCount uint32             `protobuf:"varint,78,opt,name=table_refresh_sub_group_throttled_count,json=tableRefreshSubGroupThrottledCount,proto3" json:"table_refresh_sub_group_throttled_count,omitempty"`
	UpdateTableFilterGroupCount        uint32             `protobuf:"varint,79,opt,name=update_table_filter_group_count,json=updateTableFilterGroupCount,proto3" json:"update_table_filter_group_count,omitempty"`
	TableNeighborCount                 uint32             `protobuf:"varint,80,opt,name=table_neighbor_count,json=tableNeighborCount,proto3" json:"table_neighbor_count,omitempty"`
	UpdateStatistics                   *BgpUpdgenStatsBag `protobuf:"bytes,81,opt,name=update_statistics,json=updateStatistics,proto3" json:"update_statistics,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}           `json:"-"`
	XXX_unrecognized                   []byte             `json:"-"`
	XXX_sizecache                      int32              `json:"-"`
}

func (m *BgpUpdgenAfBag) Reset()         { *m = BgpUpdgenAfBag{} }
func (m *BgpUpdgenAfBag) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenAfBag) ProtoMessage()    {}
func (*BgpUpdgenAfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0755445336bd1ac1, []int{3}
}

func (m *BgpUpdgenAfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenAfBag.Unmarshal(m, b)
}
func (m *BgpUpdgenAfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenAfBag.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenAfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenAfBag.Merge(m, src)
}
func (m *BgpUpdgenAfBag) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenAfBag.Size(m)
}
func (m *BgpUpdgenAfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenAfBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenAfBag proto.InternalMessageInfo

func (m *BgpUpdgenAfBag) GetUpdateGroupAfName() string {
	if m != nil {
		return m.UpdateGroupAfName
	}
	return ""
}

func (m *BgpUpdgenAfBag) GetUpdateMainTableVersion() uint32 {
	if m != nil {
		return m.UpdateMainTableVersion
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateRibVersion() uint32 {
	if m != nil {
		return m.UpdateRibVersion
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateMinNeighborVersion() uint32 {
	if m != nil {
		return m.UpdateMinNeighborVersion
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetCurrentUpdateLimitAf() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitAf
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetConfiguredUpdateLimitAf() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitAf
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetCurrentUpdateLimitSubgrpEbgp() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitSubgrpEbgp
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetConfiguredUpdateLimitSubgrpEbgp() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitSubgrpEbgp
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetCurrentUpdateLimitSubGroupIbgp() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitSubGroupIbgp
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetConfiguredUpdateLimitSubGroupIbgp() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitSubGroupIbgp
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateThrottled() bool {
	if m != nil {
		return m.UpdateThrottled
	}
	return false
}

func (m *BgpUpdgenAfBag) GetUpdateUpdateGroupCount() uint32 {
	if m != nil {
		return m.UpdateUpdateGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateSubGroupCount() uint32 {
	if m != nil {
		return m.UpdateSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.SubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetRefreshSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetFilterGroupCount() uint32 {
	if m != nil {
		return m.FilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateTableVrfName() string {
	if m != nil {
		return m.UpdateTableVrfName
	}
	return ""
}

func (m *BgpUpdgenAfBag) GetUpdateVrfafName() uint32 {
	if m != nil {
		return m.UpdateVrfafName
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateVrfRdVersion() uint32 {
	if m != nil {
		return m.UpdateVrfRdVersion
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateVrfTableRibVersion() uint32 {
	if m != nil {
		return m.UpdateVrfTableRibVersion
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetTableUpdateGroupCount() uint32 {
	if m != nil {
		return m.TableUpdateGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateTableSubGroupCount() uint32 {
	if m != nil {
		return m.UpdateTableSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetTableSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.TableSubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetTableRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.TableRefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetTableRefreshSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.TableRefreshSubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateTableFilterGroupCount() uint32 {
	if m != nil {
		return m.UpdateTableFilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetTableNeighborCount() uint32 {
	if m != nil {
		return m.TableNeighborCount
	}
	return 0
}

func (m *BgpUpdgenAfBag) GetUpdateStatistics() *BgpUpdgenStatsBag {
	if m != nil {
		return m.UpdateStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpUpdgenAfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_address_family.bgp_updgen_af_bag_KEYS")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_address_family.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_address_family.bgp_updgen_stats_bag")
	proto.RegisterType((*BgpUpdgenAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.update_generation_address_family.bgp_updgen_af_bag")
}

func init() { proto.RegisterFile("bgp_updgen_af_bag.proto", fileDescriptor_0755445336bd1ac1) }

var fileDescriptor_0755445336bd1ac1 = []byte{
	// 1296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0xcd, 0x52, 0xdc, 0x46,
	0x10, 0xae, 0x75, 0x1c, 0x63, 0xb7, 0x8d, 0x8d, 0x65, 0xec, 0x15, 0x60, 0xcc, 0x02, 0x89, 0x4d,
	0x52, 0xd4, 0x26, 0x01, 0x13, 0x82, 0x1d, 0x3b, 0xac, 0xf9, 0x31, 0x36, 0x3f, 0x0e, 0x0b, 0xa6,
	0x2a, 0x27, 0xd5, 0x48, 0x1a, 0x69, 0xa7, 0x4a, 0x2b, 0xc9, 0x33, 0xd2, 0x56, 0xf0, 0x2d, 0xc7,
	0xe4, 0x85, 0x72, 0xca, 0x31, 0x95, 0xe7, 0xc8, 0x9b, 0xa4, 0xe6, 0x47, 0xd2, 0xac, 0xa4, 0xa5,
	0x72, 0x33, 0x17, 0xd8, 0x55, 0xf7, 0xf7, 0x75, 0xf7, 0x37, 0x3d, 0x3d, 0xa3, 0x85, 0xa6, 0xed,
	0xc7, 0x56, 0x1a, 0xbb, 0x3e, 0x0e, 0x2d, 0xe4, 0x59, 0x36, 0xf2, 0xdb, 0x31, 0x8d, 0x92, 0xc8,
	0x60, 0x0e, 0x61, 0x4e, 0x64, 0x91, 0x88, 0x59, 0xbf, 0x52, 0x8b, 0xc4, 0x83, 0xa7, 0x16, 0x77,
	0x8d, 0x62, 0x4c, 0xdb, 0xb6, 0x1f, 0xb7, 0x49, 0xc8, 0x12, 0x14, 0x3a, 0x98, 0xe5, 0x9f, 0xf2,
	0x0f, 0x16, 0xff, 0xe7, 0xda, 0xe7, 0xed, 0x01, 0xf5, 0x18, 0xff, 0xd3, 0x46, 0x1e, 0x6b, 0x23,
	0xaf, 0x9d, 0xc6, 0x2e, 0x4a, 0xb0, 0xe5, 0xe3, 0x10, 0x53, 0x94, 0x90, 0x28, 0xb4, 0x90, 0xeb,
	0x52, 0xcc, 0x98, 0xe5, 0xa1, 0x3e, 0x09, 0xce, 0x17, 0x3e, 0xc0, 0x83, 0x4a, 0x3e, 0xd6, 0xfe,
	0xce, 0x2f, 0x27, 0xc6, 0x22, 0x8c, 0xe7, 0xf4, 0x21, 0xea, 0x63, 0xb3, 0xd1, 0x6a, 0x2c, 0xdd,
	0xe8, 0xde, 0xca, 0x1e, 0x1e, 0xa1, 0x3e, 0x36, 0xa6, 0xe0, 0xfa, 0x80, 0x7a, 0xd2, 0x7e, 0x45,
	0xd8, 0xc7, 0x06, 0xd4, 0x13, 0xa6, 0x26, 0x8c, 0x21, 0x65, 0xf9, 0x4c, 0x58, 0xae, 0x21, 0x61,
	0x58, 0x78, 0x0b, 0xb7, 0x78, 0xc8, 0x84, 0xf4, 0x31, 0x8b, 0xb1, 0x63, 0x98, 0x30, 0xc6, 0xb0,
	0x13, 0x85, 0x2e, 0x13, 0x21, 0xc6, 0xbb, 0xd9, 0x57, 0xa3, 0x05, 0x37, 0x43, 0x14, 0x46, 0x99,
	0xf5, 0x8a, 0xb0, 0xea, 0x8f, 0x16, 0xfe, 0x9a, 0x80, 0x49, 0x2d, 0x7f, 0x96, 0xa0, 0x84, 0xf1,
	0x12, 0x8c, 0x57, 0xf0, 0x48, 0xd5, 0x1e, 0xa5, 0x89, 0xf5, 0x21, 0xc5, 0x29, 0xb6, 0xfa, 0x98,
	0x31, 0xe4, 0x63, 0x66, 0xf5, 0x88, 0xdf, 0x53, 0xb1, 0xa6, 0xa5, 0xd7, 0xbb, 0x34, 0x39, 0xe6,
	0x3e, 0x87, 0xca, 0x65, 0x8f, 0xf8, 0x3d, 0xe3, 0x08, 0xbe, 0x18, 0xcd, 0xe1, 0xa4, 0xfd, 0x34,
	0x40, 0x09, 0x19, 0x60, 0x95, 0x57, 0xab, 0x9e, 0x69, 0x2b, 0xf7, 0x33, 0x0e, 0x60, 0x71, 0x34,
	0x9f, 0x4b, 0x98, 0x83, 0xa8, 0x8b, 0x5d, 0xa1, 0xd6, 0x78, 0x77, 0xae, 0x9e, 0x6e, 0x3b, 0x73,
	0x33, 0xf6, 0x60, 0xfe, 0x82, 0xec, 0x02, 0x8c, 0x28, 0x76, 0xcd, 0xab, 0x82, 0x6b, 0x76, 0x44,
	0x6a, 0xd2, 0xc9, 0x78, 0x06, 0xd3, 0x15, 0x26, 0x46, 0x3e, 0x62, 0xa9, 0xd3, 0xe7, 0x82, 0xe2,
	0xc1, 0x30, 0xc5, 0x09, 0xf9, 0x88, 0x85, 0x46, 0xbb, 0xd0, 0xaa, 0xc7, 0x6a, 0xfa, 0x5c, 0x6b,
	0x35, 0x96, 0xae, 0x76, 0x1f, 0x56, 0x19, 0x34, 0x6d, 0xb6, 0x61, 0xae, 0x9e, 0xa7, 0xd0, 0x65,
	0x4c, 0xd0, 0xcc, 0x54, 0x69, 0x0a, 0x4d, 0x36, 0x61, 0x76, 0x44, 0x36, 0x4a, 0x8f, 0xeb, 0x82,
	0x63, 0xaa, 0x26, 0x15, 0xa5, 0xc5, 0xdf, 0x0d, 0x98, 0x0d, 0x10, 0x4b, 0x2c, 0xc5, 0xa3, 0xc2,
	0xcb, 0x6e, 0x4d, 0x50, 0x3f, 0x36, 0x6f, 0xb4, 0x1a, 0x4b, 0x37, 0x57, 0x7e, 0x6b, 0xb4, 0x3f,
	0xc1, 0x76, 0x6d, 0xeb, 0x1b, 0xa7, 0x3b, 0xcd, 0x13, 0x7d, 0x2f, 0x20, 0x4a, 0x82, 0xd3, 0x2c,
	0x4b, 0x63, 0x0d, 0x9a, 0x75, 0x65, 0x20, 0x1f, 0x9b, 0x20, 0x16, 0x74, 0xb2, 0x02, 0xee, 0xf8,
	0xb8, 0x52, 0xbe, 0xd2, 0x4d, 0x2b, 0xff, 0xe6, 0x65, 0x2c, 0x5f, 0xad, 0x5e, 0x51, 0xfe, 0x2a,
	0x3c, 0xa8, 0x94, 0x21, 0xab, 0xbf, 0x25, 0xaa, 0xbf, 0x57, 0xc2, 0x8a, 0xe2, 0x57, 0xe0, 0xbe,
	0xf2, 0x4f, 0x7a, 0x34, 0x4a, 0x92, 0x00, 0x5b, 0x4e, 0x94, 0x86, 0x89, 0x39, 0x2e, 0x31, 0xd2,
	0x78, 0xaa, 0x6c, 0x5b, 0xdc, 0x64, 0xfc, 0xd3, 0x80, 0x47, 0x7a, 0xa4, 0x1c, 0x59, 0x28, 0x76,
	0xfb, 0xd2, 0x28, 0x36, 0x53, 0x54, 0x9d, 0x55, 0x51, 0x48, 0xb6, 0x0e, 0x66, 0x6d, 0x21, 0x5c,
	0xb4, 0x3b, 0x42, 0x80, 0xfb, 0x55, 0xf8, 0xb0, 0x6c, 0x14, 0x3b, 0xd1, 0x00, 0xd3, 0x73, 0x25,
	0xdb, 0x84, 0x2e, 0x5b, 0x57, 0xd9, 0xea, 0x65, 0xcb, 0x91, 0x85, 0x6c, 0x77, 0x2f, 0xa3, 0x6c,
	0x59, 0x15, 0x23, 0x65, 0xcb, 0x0b, 0xe1, 0xb2, 0x19, 0x65, 0xd9, 0x32, 0x38, 0x97, 0xad, 0x38,
	0x0d, 0xfa, 0xb8, 0x1f, 0x71, 0x44, 0x10, 0x44, 0x8e, 0xcc, 0xc6, 0x43, 0x24, 0x50, 0x22, 0xde,
	0xd3, 0x4f, 0x83, 0x43, 0xe1, 0xd9, 0xc9, 0x1d, 0x77, 0x11, 0x09, 0xa4, 0xa0, 0xff, 0x36, 0x60,
	0x59, 0xcf, 0x63, 0x04, 0x67, 0x21, 0xef, 0xe4, 0xa5, 0x91, 0xf7, 0x49, 0xa1, 0x4f, 0x5d, 0x7d,
	0x85, 0xd4, 0xc7, 0xf0, 0xf8, 0x7f, 0x94, 0xc8, 0x85, 0xbf, 0x2f, 0x44, 0x9b, 0xbf, 0x98, 0xb8,
	0xe3, 0xe3, 0x85, 0x3f, 0x26, 0xe0, 0x6e, 0xe5, 0xfe, 0x63, 0x7c, 0x03, 0x93, 0x59, 0x09, 0x34,
	0x4a, 0x63, 0x2b, 0xbb, 0xc7, 0xac, 0x88, 0x7b, 0xcc, 0x5d, 0x69, 0x7b, 0xcd, 0x4d, 0x1d, 0x79,
	0xd7, 0xd9, 0x80, 0xa9, 0x2c, 0x29, 0x44, 0x42, 0x2b, 0x41, 0x76, 0x80, 0xad, 0x01, 0xa6, 0x8c,
	0x44, 0xa1, 0xb9, 0xaa, 0x1f, 0xa0, 0x87, 0x88, 0x84, 0xa7, 0xdc, 0x7c, 0x26, 0xad, 0xc6, 0x32,
	0x18, 0x59, 0xeb, 0x10, 0x3b, 0xc7, 0x3c, 0x15, 0x98, 0x09, 0xb5, 0x75, 0x88, 0x9d, 0x79, 0xbf,
	0x80, 0x99, 0x2c, 0x10, 0x09, 0xad, 0x10, 0x13, 0xbf, 0x67, 0x47, 0x34, 0x87, 0xad, 0x09, 0x98,
	0xa9, 0x42, 0x91, 0xf0, 0x48, 0x39, 0x64, 0xf0, 0x35, 0x68, 0x3a, 0x29, 0xa5, 0x38, 0xcc, 0x45,
	0x0c, 0x48, 0x9f, 0x24, 0x16, 0xf2, 0xcc, 0xef, 0xe5, 0xa9, 0xa0, 0xcc, 0x52, 0xb5, 0x03, 0x6e,
	0xec, 0x78, 0xc6, 0x73, 0x98, 0x76, 0xa2, 0xd0, 0x23, 0x7e, 0xca, 0xcf, 0x82, 0x32, 0x72, 0x5d,
	0x20, 0x9b, 0x85, 0xc7, 0x30, 0x78, 0x17, 0x5a, 0xb5, 0x31, 0x59, 0x6a, 0xfb, 0x34, 0xb6, 0xb0,
	0xed, 0xc7, 0xe6, 0x0f, 0x82, 0xe2, 0x61, 0x35, 0xf8, 0x89, 0x70, 0xda, 0xb1, 0xfd, 0x98, 0xef,
	0x97, 0x51, 0x49, 0xe8, 0x54, 0x1b, 0x72, 0xbf, 0xd4, 0x66, 0xa3, 0xb1, 0xed, 0xc3, 0xe2, 0xa8,
	0xac, 0xd4, 0xa2, 0x13, 0xce, 0xf6, 0x4c, 0xb0, 0x3d, 0xaa, 0x4d, 0x4c, 0x34, 0xc0, 0x1b, 0x4e,
	0x76, 0x0c, 0x8f, 0x2f, 0x48, 0x4d, 0xe7, 0x7b, 0x2e, 0x1b, 0x73, 0x54, 0x76, 0x05, 0x65, 0xd1,
	0x51, 0xd5, 0xdb, 0x9d, 0xf9, 0x63, 0xdd, 0x95, 0x2c, 0xbb, 0xd5, 0x19, 0xdf, 0xe5, 0xf3, 0x78,
	0xf8, 0x12, 0x64, 0xbe, 0x10, 0x30, 0xa3, 0x7a, 0xf9, 0x31, 0xbe, 0x82, 0x89, 0xd2, 0xd8, 0x77,
	0xcd, 0x97, 0xad, 0xc6, 0xd2, 0xf5, 0xee, 0x9d, 0xe1, 0x43, 0xcf, 0xd5, 0x12, 0x1b, 0xda, 0x22,
	0x72, 0x58, 0xfd, 0xa4, 0x27, 0xf6, 0xbe, 0xd8, 0x26, 0x72, 0x46, 0xad, 0x82, 0xb2, 0x68, 0xaa,
	0x48, 0xdc, 0xa6, 0x7e, 0x52, 0x64, 0x3a, 0x48, 0xd0, 0x06, 0x4c, 0x15, 0xde, 0x79, 0x76, 0x0a,
	0xd7, 0x91, 0xf1, 0x98, 0x42, 0xe4, 0x59, 0x4a, 0xe8, 0x1a, 0x34, 0x29, 0xf6, 0x28, 0x66, 0xbd,
	0x4a, 0xc0, 0x57, 0xb2, 0xdb, 0x95, 0x79, 0x38, 0xe2, 0x1e, 0xcc, 0x57, 0x61, 0xe5, 0xc8, 0x5b,
	0xf2, 0x62, 0x5d, 0x22, 0x28, 0x25, 0xb0, 0x0c, 0x86, 0x47, 0x82, 0x04, 0xd3, 0xa1, 0xd8, 0xdb,
	0x72, 0x6f, 0x4b, 0x8b, 0x16, 0xf7, 0x4b, 0xb8, 0x9d, 0x6f, 0x68, 0xe9, 0xb9, 0x23, 0x3c, 0xc7,
	0xb3, 0xa7, 0xd2, 0xad, 0x58, 0x5e, 0x35, 0x66, 0xb2, 0xf7, 0xaf, 0x5d, 0x31, 0x9d, 0xd4, 0xf2,
	0xca, 0x19, 0xa3, 0x5e, 0xc5, 0xbe, 0x06, 0x35, 0xb3, 0xb8, 0x73, 0x36, 0xcc, 0x5e, 0x0b, 0x72,
	0xb5, 0xbe, 0x67, 0xfc, 0xb9, 0xf0, 0x2d, 0xe8, 0x39, 0x31, 0x75, 0xf3, 0xd9, 0xb2, 0xa7, 0x77,
	0xcf, 0x19, 0xf5, 0xba, 0x6e, 0x36, 0x55, 0x5e, 0xc2, 0x43, 0x0d, 0x22, 0xb3, 0xd2, 0x87, 0xd9,
	0x1b, 0x7d, 0x2a, 0x9d, 0x51, 0x4f, 0xe4, 0xa6, 0x0d, 0xb5, 0x75, 0x30, 0x25, 0xa8, 0xa6, 0xa3,
	0xde, 0xca, 0x23, 0x54, 0xd8, 0x2b, 0x0d, 0x55, 0x04, 0x96, 0xf8, 0xf2, 0x2a, 0xef, 0xeb, 0x81,
	0x45, 0xd4, 0xe1, 0x95, 0xde, 0x86, 0xb9, 0x32, 0xb0, 0xbc, 0xce, 0x07, 0x82, 0x62, 0x26, 0xd1,
	0xc1, 0xa5, 0x55, 0xde, 0x84, 0x59, 0x55, 0xf3, 0x88, 0x66, 0x3b, 0x14, 0x1c, 0x53, 0xc2, 0xa9,
	0x5b, 0xd7, 0x71, 0x27, 0xf0, 0x64, 0x14, 0x43, 0x39, 0x9f, 0x23, 0xc1, 0xb5, 0x50, 0xc7, 0x55,
	0x4a, 0xab, 0x78, 0xa3, 0x92, 0xdc, 0x35, 0x9d, 0xf8, 0x4e, 0x16, 0xa7, 0xe9, 0xb3, 0x5b, 0x6e,
	0xca, 0x6f, 0x61, 0x52, 0xc2, 0x4b, 0xad, 0xf9, 0xb3, 0xec, 0x06, 0x61, 0x3b, 0x1a, 0xea, 0xcf,
	0x3f, 0x1b, 0x79, 0xb7, 0xf1, 0xd7, 0x71, 0xc2, 0x12, 0xe2, 0x30, 0xf3, 0x58, 0x5c, 0x37, 0x7e,
	0xff, 0x74, 0xd7, 0x8d, 0xf2, 0x2f, 0x04, 0xd9, 0xe1, 0x7a, 0x92, 0xe7, 0x68, 0x5f, 0x13, 0xbf,
	0xc3, 0xac, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x38, 0x37, 0x49, 0x41, 0xa2, 0x11, 0x00, 0x00,
}
