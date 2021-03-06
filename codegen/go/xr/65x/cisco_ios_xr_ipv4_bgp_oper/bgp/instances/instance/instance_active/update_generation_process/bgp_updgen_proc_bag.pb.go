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
// source: bgp_updgen_proc_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_update_generation_process

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

type BgpUpdgenProcBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpUpdgenProcBag_KEYS) Reset()         { *m = BgpUpdgenProcBag_KEYS{} }
func (m *BgpUpdgenProcBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenProcBag_KEYS) ProtoMessage()    {}
func (*BgpUpdgenProcBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_062fdc68ec5367c1, []int{0}
}

func (m *BgpUpdgenProcBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenProcBag_KEYS.Unmarshal(m, b)
}
func (m *BgpUpdgenProcBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenProcBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenProcBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenProcBag_KEYS.Merge(m, src)
}
func (m *BgpUpdgenProcBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenProcBag_KEYS.Size(m)
}
func (m *BgpUpdgenProcBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenProcBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenProcBag_KEYS proto.InternalMessageInfo

func (m *BgpUpdgenProcBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

type BgpUpdgenAfSummaryBag struct {
	UpdateGroupAfName             string   `protobuf:"bytes,1,opt,name=update_group_af_name,json=updateGroupAfName,proto3" json:"update_group_af_name,omitempty"`
	CurrentUpdateLimitAf          uint32   `protobuf:"varint,2,opt,name=current_update_limit_af,json=currentUpdateLimitAf,proto3" json:"current_update_limit_af,omitempty"`
	UpdateOutQueueMessages        uint32   `protobuf:"varint,3,opt,name=update_out_queue_messages,json=updateOutQueueMessages,proto3" json:"update_out_queue_messages,omitempty"`
	UpdateOutQueueSize            uint32   `protobuf:"varint,4,opt,name=update_out_queue_size,json=updateOutQueueSize,proto3" json:"update_out_queue_size,omitempty"`
	UpdateThrottled               bool     `protobuf:"varint,5,opt,name=update_throttled,json=updateThrottled,proto3" json:"update_throttled,omitempty"`
	UpdateUpdateGroupCount        uint32   `protobuf:"varint,6,opt,name=update_update_group_count,json=updateUpdateGroupCount,proto3" json:"update_update_group_count,omitempty"`
	UpdateSubGroupCount           uint32   `protobuf:"varint,7,opt,name=update_sub_group_count,json=updateSubGroupCount,proto3" json:"update_sub_group_count,omitempty"`
	SubGroupThrottledCount        uint32   `protobuf:"varint,8,opt,name=sub_group_throttled_count,json=subGroupThrottledCount,proto3" json:"sub_group_throttled_count,omitempty"`
	RefreshSubGroupCount          uint32   `protobuf:"varint,9,opt,name=refresh_sub_group_count,json=refreshSubGroupCount,proto3" json:"refresh_sub_group_count,omitempty"`
	RefreshSubGroupThrottledCount uint32   `protobuf:"varint,10,opt,name=refresh_sub_group_throttled_count,json=refreshSubGroupThrottledCount,proto3" json:"refresh_sub_group_throttled_count,omitempty"`
	FilterGroupCount              uint32   `protobuf:"varint,11,opt,name=filter_group_count,json=filterGroupCount,proto3" json:"filter_group_count,omitempty"`
	NeighborCount                 uint32   `protobuf:"varint,12,opt,name=neighbor_count,json=neighborCount,proto3" json:"neighbor_count,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *BgpUpdgenAfSummaryBag) Reset()         { *m = BgpUpdgenAfSummaryBag{} }
func (m *BgpUpdgenAfSummaryBag) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenAfSummaryBag) ProtoMessage()    {}
func (*BgpUpdgenAfSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_062fdc68ec5367c1, []int{1}
}

func (m *BgpUpdgenAfSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenAfSummaryBag.Unmarshal(m, b)
}
func (m *BgpUpdgenAfSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenAfSummaryBag.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenAfSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenAfSummaryBag.Merge(m, src)
}
func (m *BgpUpdgenAfSummaryBag) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenAfSummaryBag.Size(m)
}
func (m *BgpUpdgenAfSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenAfSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenAfSummaryBag proto.InternalMessageInfo

func (m *BgpUpdgenAfSummaryBag) GetUpdateGroupAfName() string {
	if m != nil {
		return m.UpdateGroupAfName
	}
	return ""
}

func (m *BgpUpdgenAfSummaryBag) GetCurrentUpdateLimitAf() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitAf
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetUpdateThrottled() bool {
	if m != nil {
		return m.UpdateThrottled
	}
	return false
}

func (m *BgpUpdgenAfSummaryBag) GetUpdateUpdateGroupCount() uint32 {
	if m != nil {
		return m.UpdateUpdateGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetUpdateSubGroupCount() uint32 {
	if m != nil {
		return m.UpdateSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.SubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetRefreshSubGroupCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetRefreshSubGroupThrottledCount() uint32 {
	if m != nil {
		return m.RefreshSubGroupThrottledCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetFilterGroupCount() uint32 {
	if m != nil {
		return m.FilterGroupCount
	}
	return 0
}

func (m *BgpUpdgenAfSummaryBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
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
	return fileDescriptor_062fdc68ec5367c1, []int{2}
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
	return fileDescriptor_062fdc68ec5367c1, []int{3}
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

type BgpUpdgenProcBag struct {
	UpdateOutLogging             bool                     `protobuf:"varint,50,opt,name=update_out_logging,json=updateOutLogging,proto3" json:"update_out_logging,omitempty"`
	CurrentUpdateLimitProcess    uint32                   `protobuf:"varint,51,opt,name=current_update_limit_process,json=currentUpdateLimitProcess,proto3" json:"current_update_limit_process,omitempty"`
	ConfiguredUpdateLimitProcess uint32                   `protobuf:"varint,52,opt,name=configured_update_limit_process,json=configuredUpdateLimitProcess,proto3" json:"configured_update_limit_process,omitempty"`
	UpdateOutQueueMessages       uint32                   `protobuf:"varint,53,opt,name=update_out_queue_messages,json=updateOutQueueMessages,proto3" json:"update_out_queue_messages,omitempty"`
	UpdateOutQueueSize           uint32                   `protobuf:"varint,54,opt,name=update_out_queue_size,json=updateOutQueueSize,proto3" json:"update_out_queue_size,omitempty"`
	UpdateThrottled              bool                     `protobuf:"varint,55,opt,name=update_throttled,json=updateThrottled,proto3" json:"update_throttled,omitempty"`
	UpdateAddressFamily          []*BgpUpdgenAfSummaryBag `protobuf:"bytes,56,rep,name=update_address_family,json=updateAddressFamily,proto3" json:"update_address_family,omitempty"`
	UpdateStatistics             *BgpUpdgenStatsBag       `protobuf:"bytes,57,opt,name=update_statistics,json=updateStatistics,proto3" json:"update_statistics,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                 `json:"-"`
	XXX_unrecognized             []byte                   `json:"-"`
	XXX_sizecache                int32                    `json:"-"`
}

func (m *BgpUpdgenProcBag) Reset()         { *m = BgpUpdgenProcBag{} }
func (m *BgpUpdgenProcBag) String() string { return proto.CompactTextString(m) }
func (*BgpUpdgenProcBag) ProtoMessage()    {}
func (*BgpUpdgenProcBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_062fdc68ec5367c1, []int{4}
}

func (m *BgpUpdgenProcBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpUpdgenProcBag.Unmarshal(m, b)
}
func (m *BgpUpdgenProcBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpUpdgenProcBag.Marshal(b, m, deterministic)
}
func (m *BgpUpdgenProcBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpUpdgenProcBag.Merge(m, src)
}
func (m *BgpUpdgenProcBag) XXX_Size() int {
	return xxx_messageInfo_BgpUpdgenProcBag.Size(m)
}
func (m *BgpUpdgenProcBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpUpdgenProcBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpUpdgenProcBag proto.InternalMessageInfo

func (m *BgpUpdgenProcBag) GetUpdateOutLogging() bool {
	if m != nil {
		return m.UpdateOutLogging
	}
	return false
}

func (m *BgpUpdgenProcBag) GetCurrentUpdateLimitProcess() uint32 {
	if m != nil {
		return m.CurrentUpdateLimitProcess
	}
	return 0
}

func (m *BgpUpdgenProcBag) GetConfiguredUpdateLimitProcess() uint32 {
	if m != nil {
		return m.ConfiguredUpdateLimitProcess
	}
	return 0
}

func (m *BgpUpdgenProcBag) GetUpdateOutQueueMessages() uint32 {
	if m != nil {
		return m.UpdateOutQueueMessages
	}
	return 0
}

func (m *BgpUpdgenProcBag) GetUpdateOutQueueSize() uint32 {
	if m != nil {
		return m.UpdateOutQueueSize
	}
	return 0
}

func (m *BgpUpdgenProcBag) GetUpdateThrottled() bool {
	if m != nil {
		return m.UpdateThrottled
	}
	return false
}

func (m *BgpUpdgenProcBag) GetUpdateAddressFamily() []*BgpUpdgenAfSummaryBag {
	if m != nil {
		return m.UpdateAddressFamily
	}
	return nil
}

func (m *BgpUpdgenProcBag) GetUpdateStatistics() *BgpUpdgenStatsBag {
	if m != nil {
		return m.UpdateStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpUpdgenProcBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.update_generation_process.bgp_updgen_proc_bag_KEYS")
	proto.RegisterType((*BgpUpdgenAfSummaryBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.update_generation_process.bgp_updgen_af_summary_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.update_generation_process.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.update_generation_process.bgp_updgen_stats_bag")
	proto.RegisterType((*BgpUpdgenProcBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.update_generation_process.bgp_updgen_proc_bag")
}

func init() { proto.RegisterFile("bgp_updgen_proc_bag.proto", fileDescriptor_062fdc68ec5367c1) }

var fileDescriptor_062fdc68ec5367c1 = []byte{
	// 1030 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x69, 0xd3, 0xa4, 0xe3, 0xb8, 0x49, 0x26, 0x4e, 0xbb, 0x4e, 0x93, 0xc6, 0x75, 0xf9,
	0x09, 0x52, 0x65, 0x44, 0xd2, 0x50, 0xc2, 0x4d, 0x31, 0x6d, 0x43, 0x05, 0x69, 0xa1, 0x4e, 0x7b,
	0x81, 0xb8, 0x18, 0x8d, 0x77, 0xc7, 0xeb, 0x91, 0x76, 0x77, 0xb6, 0x33, 0xb3, 0x11, 0xe9, 0xbb,
	0xc0, 0x03, 0x20, 0xc4, 0x05, 0x4f, 0xc0, 0x25, 0x8f, 0x85, 0xe6, 0x67, 0x7f, 0xec, 0x1d, 0x47,
	0x50, 0xa9, 0xc9, 0x5d, 0x32, 0xe7, 0x7c, 0xdf, 0xf9, 0xce, 0xe7, 0x99, 0xe3, 0x63, 0xd0, 0x19,
	0x85, 0x29, 0xca, 0xd2, 0x20, 0x24, 0x09, 0x4a, 0x39, 0xf3, 0xd1, 0x08, 0x87, 0xfd, 0x94, 0x33,
	0xc9, 0xe0, 0xcf, 0x3e, 0x15, 0x3e, 0x43, 0x94, 0x09, 0xf4, 0x0b, 0x47, 0x34, 0x3d, 0x7d, 0x80,
	0x54, 0x32, 0x4b, 0x09, 0xef, 0x8f, 0xc2, 0xb4, 0x4f, 0x13, 0x21, 0x71, 0xe2, 0x13, 0x51, 0xfc,
	0x55, 0xfc, 0x81, 0xb0, 0x2f, 0xe9, 0x29, 0xe9, 0x67, 0x69, 0x80, 0x25, 0x41, 0x21, 0x49, 0x08,
	0xc7, 0x92, 0x32, 0x53, 0x83, 0x08, 0xd1, 0x7b, 0x04, 0x3c, 0x47, 0x65, 0xf4, 0xfd, 0xd3, 0x9f,
	0x4e, 0xe0, 0x3d, 0xd0, 0x2a, 0x88, 0x12, 0x1c, 0x13, 0xaf, 0xd1, 0x6d, 0xec, 0x5e, 0x1f, 0x2e,
	0xe7, 0x87, 0x2f, 0x70, 0x4c, 0x7a, 0xbf, 0x2e, 0x4c, 0x69, 0xc7, 0x63, 0x24, 0xb2, 0x38, 0xc6,
	0xfc, 0x4c, 0xf1, 0xc0, 0xcf, 0x40, 0x3b, 0xaf, 0xcd, 0x59, 0x96, 0xaa, 0x70, 0x85, 0x69, 0xcd,
	0xc4, 0xbe, 0x55, 0xa1, 0xc1, 0x58, 0xd1, 0xc1, 0x03, 0x70, 0xcb, 0xcf, 0x38, 0x27, 0x89, 0x44,
	0x16, 0x18, 0xd1, 0x98, 0x4a, 0x84, 0xc7, 0xde, 0x07, 0xdd, 0xc6, 0x6e, 0x6b, 0xd8, 0xb6, 0xe1,
	0xd7, 0x3a, 0x7a, 0xac, 0x82, 0x83, 0x31, 0x3c, 0x04, 0x1d, 0x9b, 0xce, 0x32, 0x89, 0xde, 0x64,
	0x24, 0x23, 0x28, 0x26, 0x42, 0xe0, 0x90, 0x08, 0xef, 0x8a, 0x06, 0xde, 0x34, 0x09, 0x3f, 0x64,
	0xf2, 0xa5, 0x0a, 0x3f, 0xb7, 0x51, 0xf8, 0x39, 0xd8, 0xa8, 0x41, 0x05, 0x7d, 0x4b, 0xbc, 0xab,
	0x1a, 0x06, 0xa7, 0x61, 0x27, 0xf4, 0x2d, 0x81, 0x9f, 0x82, 0x55, 0x0b, 0x91, 0x13, 0xce, 0xa4,
	0x8c, 0x48, 0xe0, 0x2d, 0x74, 0x1b, 0xbb, 0x4b, 0xc3, 0x15, 0x73, 0xfe, 0x2a, 0x3f, 0xae, 0x08,
	0x9b, 0xf2, 0xc1, 0x67, 0x59, 0x22, 0xbd, 0x6b, 0x55, 0x61, 0xaf, 0x4b, 0x2f, 0x1e, 0xab, 0x28,
	0xdc, 0x07, 0x36, 0x82, 0x44, 0x36, 0x9a, 0xc2, 0x2d, 0x6a, 0xdc, 0xba, 0x89, 0x9e, 0x64, 0xa3,
	0x0a, 0xe8, 0x10, 0x74, 0xca, 0xec, 0x42, 0x9d, 0xc5, 0x2d, 0x99, 0x7a, 0xc2, 0x22, 0x0a, 0x95,
	0x06, 0x7a, 0x00, 0x6e, 0x71, 0x32, 0xe6, 0x44, 0x4c, 0x6a, 0x05, 0xaf, 0x1b, 0xeb, 0x6d, 0x78,
	0xba, 0xe2, 0x33, 0x70, 0xb7, 0x0e, 0x9b, 0xad, 0x0c, 0x34, 0xc1, 0xf6, 0x0c, 0xc1, 0x8c, 0x80,
	0xfb, 0x00, 0x8e, 0x69, 0x24, 0x09, 0x9f, 0xaa, 0xdd, 0xd4, 0xd0, 0x55, 0x13, 0xa9, 0xd4, 0xfd,
	0x08, 0xdc, 0x48, 0x08, 0x0d, 0x27, 0x23, 0xc6, 0x6d, 0xe6, 0xb2, 0xce, 0x6c, 0xe5, 0xa7, 0x3a,
	0xad, 0xf7, 0x1d, 0x58, 0x56, 0xd7, 0x53, 0xd2, 0x98, 0x88, 0x94, 0xf8, 0xd0, 0x03, 0x8b, 0x82,
	0xf8, 0x2c, 0x09, 0x84, 0xbe, 0x84, 0xad, 0x61, 0xfe, 0x2f, 0xec, 0x82, 0x66, 0x82, 0x13, 0x96,
	0x47, 0xcd, 0x75, 0xab, 0x1e, 0xf5, 0xfe, 0x5e, 0x01, 0xed, 0xca, 0x5d, 0x17, 0x12, 0x4b, 0xa1,
	0xaf, 0xf9, 0x37, 0xe0, 0xce, 0xdc, 0xeb, 0x87, 0x26, 0x34, 0x9c, 0xd8, 0x5a, 0x9b, 0xee, 0x3b,
	0xf8, 0x8c, 0x86, 0x13, 0xf8, 0x02, 0x7c, 0x38, 0x9f, 0xc3, 0xcf, 0xe2, 0x2c, 0xc2, 0xea, 0x2d,
	0x5b, 0x5d, 0x5d, 0x37, 0xd3, 0xe3, 0x22, 0x0f, 0x1e, 0x83, 0x7b, 0xf3, 0xf9, 0x02, 0x2a, 0x7c,
	0xcc, 0x03, 0x12, 0xd8, 0xc7, 0xb1, 0xe3, 0xa6, 0x7b, 0x92, 0xa7, 0xa9, 0x4f, 0xf9, 0x1c, 0x75,
	0x11, 0xc1, 0x9c, 0x04, 0xf6, 0xc5, 0x6c, 0xcf, 0x91, 0x66, 0x92, 0xe0, 0x57, 0x60, 0xd3, 0xf9,
	0xde, 0x8c, 0x4f, 0x0b, 0xae, 0xb7, 0xaa, 0x1e, 0x9d, 0xf6, 0xe8, 0x08, 0x74, 0xdd, 0xd8, 0x8a,
	0x3f, 0xea, 0x51, 0x5d, 0x1d, 0x6e, 0xd5, 0x19, 0x2a, 0xde, 0x3c, 0x01, 0x3b, 0x6e, 0x9e, 0xd2,
	0x97, 0x45, 0x4d, 0x73, 0xbb, 0x4e, 0x53, 0x7a, 0xf2, 0x35, 0xd8, 0x9e, 0xa3, 0xc6, 0xfa, 0xb1,
	0xa4, 0x39, 0x3a, 0x0e, 0x29, 0xd6, 0x8b, 0x3f, 0x1a, 0x60, 0x3b, 0xc2, 0xa2, 0x98, 0x75, 0xb6,
	0xbc, 0xb9, 0xad, 0x12, 0xc7, 0xa9, 0x7e, 0x79, 0xcd, 0x3d, 0xda, 0x7f, 0x8f, 0xdf, 0x01, 0xfd,
	0xea, 0xfb, 0x18, 0x6e, 0x2a, 0x3d, 0x66, 0x16, 0xd9, 0x4e, 0x5f, 0xe5, 0x62, 0xd4, 0x84, 0x70,
	0xa9, 0xc5, 0x21, 0xb1, 0x0f, 0xbc, 0x5d, 0x03, 0x0f, 0x42, 0x52, 0xeb, 0xd2, 0xda, 0x53, 0xe9,
	0xb2, 0x79, 0x89, 0x5d, 0xda, 0xcf, 0xa2, 0xec, 0x72, 0x1f, 0xdc, 0xac, 0xa9, 0x35, 0x4d, 0x9a,
	0x01, 0xb3, 0x3e, 0x83, 0xd5, 0x3d, 0xee, 0x15, 0xdf, 0x22, 0xf9, 0xe8, 0xb3, 0x43, 0xa9, 0x55,
	0x9d, 0xd5, 0xf9, 0xc0, 0x33, 0x13, 0xec, 0xcf, 0x06, 0xb8, 0x53, 0xad, 0x54, 0x20, 0x4b, 0x63,
	0x6e, 0x5c, 0xb4, 0x31, 0xb7, 0xcb, 0xe6, 0x72, 0xb1, 0xa5, 0x33, 0x0f, 0x81, 0xe7, 0xd4, 0xab,
	0xbc, 0x59, 0xd1, 0x7d, 0x6e, 0xd4, 0xe1, 0xd3, 0xee, 0x70, 0xe2, 0xb3, 0x53, 0xc2, 0xcf, 0xac,
	0x3b, 0xab, 0x55, 0x77, 0x86, 0x36, 0xe6, 0x76, 0xa7, 0x40, 0x96, 0xee, 0xac, 0x5d, 0xa2, 0x3b,
	0xb9, 0xd8, 0xb9, 0xee, 0x14, 0x7a, 0x95, 0x3b, 0x70, 0xd6, 0x9d, 0x1c, 0xae, 0xdc, 0x29, 0x27,
	0x75, 0x4c, 0x62, 0xa6, 0x10, 0x51, 0xc4, 0x7c, 0x23, 0x63, 0x8c, 0x69, 0x64, 0xbd, 0x5a, 0xaf,
	0x4e, 0xea, 0xe7, 0x3a, 0x73, 0x50, 0x24, 0x1e, 0x61, 0x1a, 0x19, 0xdf, 0xfe, 0x69, 0x80, 0xfb,
	0x55, 0x1d, 0x73, 0x38, 0x4b, 0x17, 0xdb, 0x17, 0xed, 0xe2, 0x27, 0xa5, 0x0d, 0xae, 0x36, 0x4a,
	0x47, 0x5f, 0x82, 0x8f, 0xff, 0x43, 0x27, 0xca, 0xdf, 0x0d, 0xed, 0xcd, 0xdd, 0xf3, 0x89, 0x07,
	0x21, 0xe9, 0xfd, 0xb5, 0x00, 0xd6, 0x1d, 0x0b, 0xaf, 0xda, 0x3d, 0x2a, 0xb3, 0x3c, 0x62, 0x61,
	0x48, 0x93, 0xd0, 0xdb, 0xd3, 0x4b, 0xdd, 0x6a, 0x31, 0xc0, 0x8f, 0xcd, 0x39, 0x7c, 0x04, 0xb6,
	0x9c, 0x5b, 0xaa, 0x6d, 0xd7, 0xdb, 0xd7, 0x72, 0x3a, 0xf5, 0x55, 0xf5, 0x47, 0x93, 0x00, 0x9f,
	0x82, 0x1d, 0x9f, 0x25, 0x63, 0x1a, 0x66, 0x6a, 0x10, 0x3a, 0x39, 0x1e, 0x68, 0x8e, 0xad, 0x32,
	0xcd, 0x41, 0x73, 0xee, 0xda, 0x7b, 0xf0, 0x6e, 0x6b, 0xef, 0x17, 0xff, 0x6b, 0xed, 0x7d, 0xe8,
	0x5e, 0x7b, 0x7f, 0x6f, 0x14, 0xf4, 0x38, 0x08, 0x38, 0x11, 0x02, 0x8d, 0x71, 0x4c, 0xa3, 0x33,
	0xef, 0xcb, 0xee, 0x95, 0xdd, 0xe6, 0xde, 0xe9, 0x7b, 0xbf, 0x6d, 0xce, 0xdf, 0x23, 0xf9, 0xa4,
	0x19, 0x18, 0x4d, 0x47, 0x5a, 0x12, 0xfc, 0xad, 0x01, 0xd6, 0xf2, 0x4d, 0x5b, 0x62, 0x49, 0x85,
	0xa4, 0xbe, 0xf0, 0x0e, 0xf5, 0xb3, 0x78, 0x73, 0x51, 0x42, 0x8b, 0x65, 0x32, 0xbf, 0x6e, 0x27,
	0x85, 0x94, 0xd1, 0x35, 0xfd, 0x43, 0x70, 0xff, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x18,
	0xa8, 0xef, 0x25, 0x0e, 0x00, 0x00,
}
