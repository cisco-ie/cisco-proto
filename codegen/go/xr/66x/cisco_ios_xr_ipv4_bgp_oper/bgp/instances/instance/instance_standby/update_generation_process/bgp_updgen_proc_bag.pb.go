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

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_update_generation_process

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
	proto.RegisterType((*BgpUpdgenProcBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.update_generation_process.bgp_updgen_proc_bag_KEYS")
	proto.RegisterType((*BgpUpdgenAfSummaryBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.update_generation_process.bgp_updgen_af_summary_bag")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.update_generation_process.bgp_timespec")
	proto.RegisterType((*BgpUpdgenStatsBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.update_generation_process.bgp_updgen_stats_bag")
	proto.RegisterType((*BgpUpdgenProcBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.update_generation_process.bgp_updgen_proc_bag")
}

func init() { proto.RegisterFile("bgp_updgen_proc_bag.proto", fileDescriptor_062fdc68ec5367c1) }

var fileDescriptor_062fdc68ec5367c1 = []byte{
	// 1032 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x69, 0xd3, 0xa4, 0xe3, 0xb8, 0x49, 0x26, 0x4e, 0xbb, 0x4e, 0x93, 0xc6, 0x75, 0xf9,
	0x09, 0x52, 0x65, 0x44, 0xd2, 0x50, 0xc2, 0x4d, 0x31, 0x6d, 0x43, 0x05, 0x69, 0xa1, 0x4e, 0x7b,
	0x81, 0x84, 0x34, 0x1a, 0xef, 0x8e, 0xd7, 0x83, 0x76, 0x77, 0x96, 0x99, 0xd9, 0x40, 0xfa, 0x2e,
	0x88, 0x07, 0x40, 0x88, 0x1b, 0x5e, 0x81, 0x0b, 0xde, 0x0a, 0xcd, 0xcf, 0xfe, 0xd8, 0x3b, 0x8e,
	0x00, 0xa9, 0xc9, 0x55, 0xd3, 0x39, 0xe7, 0xfb, 0xce, 0x77, 0x3e, 0xcf, 0x1c, 0x1f, 0x83, 0xce,
	0x28, 0x4c, 0x51, 0x96, 0x06, 0x21, 0x49, 0x50, 0xca, 0x99, 0x8f, 0x46, 0x38, 0xec, 0xa7, 0x9c,
	0x49, 0x06, 0xbf, 0xf7, 0xa9, 0xf0, 0x19, 0xa2, 0x4c, 0xa0, 0x9f, 0x39, 0xa2, 0xe9, 0xe9, 0x03,
	0xa4, 0x92, 0x59, 0x4a, 0x78, 0x7f, 0x14, 0xa6, 0x7d, 0x9a, 0x08, 0x89, 0x13, 0x9f, 0x88, 0xe2,
	0xaf, 0xe2, 0x0f, 0xa4, 0xfe, 0x09, 0x46, 0x67, 0xfd, 0x2c, 0x0d, 0xb0, 0x24, 0x28, 0x24, 0x09,
	0xe1, 0x58, 0x52, 0x66, 0x8a, 0x10, 0x21, 0x7a, 0x8f, 0x80, 0xe7, 0x28, 0x8d, 0xbe, 0x7e, 0xfa,
	0xdd, 0x09, 0xbc, 0x07, 0x5a, 0x05, 0x53, 0x82, 0x63, 0xe2, 0x35, 0xba, 0x8d, 0xdd, 0xeb, 0xc3,
	0xe5, 0xfc, 0xf0, 0x05, 0x8e, 0x49, 0xef, 0x97, 0x85, 0x29, 0xf1, 0x78, 0x8c, 0x44, 0x16, 0xc7,
	0x98, 0x9f, 0x29, 0x1e, 0xf8, 0x11, 0x68, 0xe7, 0xb5, 0x39, 0xcb, 0x52, 0x15, 0xae, 0x30, 0xad,
	0x99, 0xd8, 0x97, 0x2a, 0x34, 0x18, 0x2b, 0x3a, 0x78, 0x00, 0x6e, 0xf9, 0x19, 0xe7, 0x24, 0x91,
	0xc8, 0x02, 0x23, 0x1a, 0x53, 0x89, 0xf0, 0xd8, 0x7b, 0xa7, 0xdb, 0xd8, 0x6d, 0x0d, 0xdb, 0x36,
	0xfc, 0x5a, 0x47, 0x8f, 0x55, 0x70, 0x30, 0x86, 0x87, 0xa0, 0x63, 0xd3, 0x59, 0x26, 0xd1, 0x8f,
	0x19, 0xc9, 0x08, 0x8a, 0x89, 0x10, 0x38, 0x24, 0xc2, 0xbb, 0xa2, 0x81, 0x37, 0x4d, 0xc2, 0x37,
	0x99, 0x7c, 0xa9, 0xc2, 0xcf, 0x6d, 0x14, 0x7e, 0x0c, 0x36, 0x6a, 0x50, 0x41, 0xdf, 0x10, 0xef,
	0xaa, 0x86, 0xc1, 0x69, 0xd8, 0x09, 0x7d, 0x43, 0xe0, 0x87, 0x60, 0xd5, 0x42, 0xe4, 0x84, 0x33,
	0x29, 0x23, 0x12, 0x78, 0x0b, 0xdd, 0xc6, 0xee, 0xd2, 0x70, 0xc5, 0x9c, 0xbf, 0xca, 0x8f, 0x2b,
	0xc2, 0xa6, 0x7c, 0xf0, 0x59, 0x96, 0x48, 0xef, 0x5a, 0x55, 0xd8, 0xeb, 0xd2, 0x8b, 0xc7, 0x2a,
	0x0a, 0xf7, 0x81, 0x8d, 0x20, 0x91, 0x8d, 0xa6, 0x70, 0x8b, 0x1a, 0xb7, 0x6e, 0xa2, 0x27, 0xd9,
	0xa8, 0x02, 0x3a, 0x04, 0x9d, 0x32, 0xbb, 0x50, 0x67, 0x71, 0x4b, 0xa6, 0x9e, 0xb0, 0x88, 0x42,
	0xa5, 0x81, 0x1e, 0x80, 0x5b, 0x9c, 0x8c, 0x39, 0x11, 0x93, 0x5a, 0xc1, 0xeb, 0xc6, 0x7a, 0x1b,
	0x9e, 0xae, 0xf8, 0x0c, 0xdc, 0xad, 0xc3, 0x66, 0x2b, 0x03, 0x4d, 0xb0, 0x3d, 0x43, 0x30, 0x23,
	0xe0, 0x3e, 0x80, 0x63, 0x1a, 0x49, 0xc2, 0xa7, 0x6a, 0x37, 0x35, 0x74, 0xd5, 0x44, 0x2a, 0x75,
	0xdf, 0x03, 0x37, 0x12, 0x42, 0xc3, 0xc9, 0x88, 0x71, 0x9b, 0xb9, 0xac, 0x33, 0x5b, 0xf9, 0xa9,
	0x4e, 0xeb, 0x7d, 0x05, 0x96, 0xd5, 0xf5, 0x94, 0x34, 0x26, 0x22, 0x25, 0x3e, 0xf4, 0xc0, 0xa2,
	0x20, 0x3e, 0x4b, 0x02, 0xa1, 0x2f, 0x61, 0x6b, 0x98, 0xff, 0x17, 0x76, 0x41, 0x33, 0xc1, 0x09,
	0xcb, 0xa3, 0xe6, 0xba, 0x55, 0x8f, 0x7a, 0x7f, 0xad, 0x80, 0x76, 0xe5, 0xae, 0x0b, 0x89, 0xa5,
	0xd0, 0xd7, 0xfc, 0x0b, 0x70, 0x67, 0xee, 0xf5, 0x43, 0x13, 0x1a, 0x4e, 0x6c, 0xad, 0x4d, 0xf7,
	0x1d, 0x7c, 0x46, 0xc3, 0x09, 0x7c, 0x01, 0xde, 0x9d, 0xcf, 0xe1, 0x67, 0x71, 0x16, 0x61, 0x49,
	0x4f, 0x89, 0xd5, 0xd5, 0x75, 0x33, 0x3d, 0x2e, 0xf2, 0xe0, 0x31, 0xb8, 0x37, 0x9f, 0x2f, 0xa0,
	0xc2, 0xc7, 0x3c, 0x20, 0x81, 0x7d, 0x1c, 0x3b, 0x6e, 0xba, 0x27, 0x79, 0x9a, 0xfa, 0x94, 0xcf,
	0x51, 0x17, 0x11, 0xcc, 0x49, 0x60, 0x5f, 0xcc, 0xf6, 0x1c, 0x69, 0x26, 0x09, 0x7e, 0x06, 0x36,
	0x9d, 0xef, 0xcd, 0xf8, 0xb4, 0xe0, 0x7a, 0xab, 0xea, 0xd1, 0x69, 0x8f, 0x8e, 0x40, 0xd7, 0x8d,
	0xad, 0xf8, 0xa3, 0x1e, 0xd5, 0xd5, 0xe1, 0x56, 0x9d, 0xa1, 0xe2, 0xcd, 0x13, 0xb0, 0xe3, 0xe6,
	0x29, 0x7d, 0x59, 0xd4, 0x34, 0xb7, 0xeb, 0x34, 0xa5, 0x27, 0x9f, 0x83, 0xed, 0x39, 0x6a, 0xac,
	0x1f, 0x4b, 0x9a, 0xa3, 0xe3, 0x90, 0x62, 0xbd, 0xf8, 0xbd, 0x01, 0xb6, 0x23, 0x2c, 0x8a, 0x59,
	0x67, 0xcb, 0x9b, 0xdb, 0x2a, 0x71, 0x9c, 0xea, 0x97, 0xd7, 0xdc, 0xfb, 0xa1, 0xff, 0x36, 0xbf,
	0x04, 0xfa, 0xd5, 0x07, 0x32, 0xdc, 0x54, 0x82, 0xcc, 0x30, 0xb2, 0xad, 0xbe, 0xca, 0xd5, 0xa8,
	0x11, 0xe1, 0x92, 0x8b, 0x43, 0x62, 0x5f, 0x78, 0xbb, 0x06, 0x1e, 0x84, 0xa4, 0xd6, 0xa6, 0xf5,
	0xa7, 0xd2, 0x66, 0xf3, 0x32, 0xdb, 0xb4, 0x9f, 0x46, 0xd9, 0xe6, 0x3e, 0xb8, 0x59, 0x93, 0x6b,
	0xba, 0x34, 0x23, 0x66, 0x7d, 0x06, 0xab, 0x9b, 0xdc, 0x2b, 0xbe, 0x47, 0xf2, 0xe1, 0x67, 0xc7,
	0x52, 0xab, 0x3a, 0xad, 0xf3, 0x91, 0x67, 0x66, 0xd8, 0x1f, 0x0d, 0x70, 0xa7, 0x5a, 0xa9, 0x40,
	0x96, 0xce, 0xdc, 0xb8, 0x70, 0x67, 0x6e, 0x97, 0xdd, 0xe5, 0x6a, 0x4b, 0x6b, 0x1e, 0x02, 0xcf,
	0x29, 0x58, 0x99, 0xb3, 0xa2, 0x1b, 0xdd, 0xa8, 0xc3, 0xa7, 0xed, 0xe1, 0xc4, 0x67, 0xa7, 0x84,
	0x9f, 0x59, 0x7b, 0x56, 0xab, 0xf6, 0x0c, 0x6d, 0xcc, 0x6d, 0x4f, 0x81, 0x2c, 0xed, 0x59, 0xbb,
	0x4c, 0x7b, 0x72, 0xb5, 0x73, 0xed, 0x29, 0x04, 0x2b, 0x7b, 0xe0, 0xac, 0x3d, 0x39, 0x5c, 0xd9,
	0x53, 0x4e, 0xeb, 0x98, 0xc4, 0x4c, 0x21, 0xa2, 0x88, 0xf9, 0x46, 0xc6, 0x18, 0xd3, 0xc8, 0x9a,
	0xb5, 0x5e, 0x9d, 0xd6, 0xcf, 0x75, 0xe6, 0xa0, 0x48, 0x3c, 0xc2, 0x34, 0x32, 0xc6, 0xfd, 0xdd,
	0x00, 0xf7, 0xab, 0x3a, 0xe6, 0x70, 0x96, 0x36, 0xb6, 0x2f, 0xdc, 0xc6, 0x0f, 0x4a, 0x1f, 0x5c,
	0x7d, 0x94, 0x96, 0xbe, 0x04, 0xef, 0xff, 0x8b, 0x56, 0x94, 0xc1, 0x1b, 0xda, 0x9c, 0xbb, 0xe7,
	0x13, 0x0f, 0x42, 0xd2, 0xfb, 0x73, 0x01, 0xac, 0x3b, 0xb6, 0x5e, 0xb5, 0x80, 0x54, 0x06, 0x7a,
	0xc4, 0xc2, 0x90, 0x26, 0xa1, 0xb7, 0xa7, 0x37, 0xbb, 0xd5, 0x62, 0x8a, 0x1f, 0x9b, 0x73, 0xf8,
	0x08, 0x6c, 0x39, 0x57, 0x55, 0xdb, 0xae, 0xb7, 0xaf, 0xe5, 0x74, 0xea, 0xfb, 0xea, 0xb7, 0x26,
	0x01, 0x3e, 0x05, 0x3b, 0x3e, 0x4b, 0xc6, 0x34, 0xcc, 0xd4, 0x30, 0x74, 0x72, 0x3c, 0xd0, 0x1c,
	0x5b, 0x65, 0x9a, 0x83, 0xe6, 0xdc, 0xdd, 0xf7, 0xe0, 0xff, 0xed, 0xbe, 0x9f, 0xfc, 0xa7, 0xdd,
	0xf7, 0xa1, 0x7b, 0xf7, 0xfd, 0xad, 0x51, 0xd0, 0xe3, 0x20, 0xe0, 0x44, 0x08, 0x34, 0xc6, 0x31,
	0x8d, 0xce, 0xbc, 0x4f, 0xbb, 0x57, 0x76, 0x9b, 0x7b, 0x3f, 0xbd, 0xfd, 0xeb, 0xe6, 0xfc, 0x55,
	0x92, 0x0f, 0x9b, 0x81, 0x11, 0x75, 0xa4, 0x35, 0xc1, 0x5f, 0x1b, 0x60, 0x2d, 0xdf, 0xb7, 0x25,
	0x96, 0x54, 0x48, 0xea, 0x0b, 0xef, 0x50, 0x3f, 0x0c, 0x7e, 0x61, 0x4a, 0x8b, 0x9d, 0x32, 0xbf,
	0x70, 0x27, 0x85, 0x96, 0xd1, 0x35, 0xfd, 0x83, 0x70, 0xff, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0xb5, 0x7d, 0x53, 0x2d, 0x0e, 0x00, 0x00,
}
