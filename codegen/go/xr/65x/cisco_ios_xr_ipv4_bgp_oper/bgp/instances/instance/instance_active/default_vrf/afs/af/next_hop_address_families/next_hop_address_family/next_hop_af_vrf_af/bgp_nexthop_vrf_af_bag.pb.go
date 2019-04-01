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
// source: bgp_nexthop_vrf_af_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_next_hop_address_families_next_hop_address_family_next_hop_af_vrf_af

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

type BgpNexthopVrfAfBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	NextHopAfName        string   `protobuf:"bytes,3,opt,name=next_hop_af_name,json=nextHopAfName,proto3" json:"next_hop_af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpNexthopVrfAfBag_KEYS) Reset()         { *m = BgpNexthopVrfAfBag_KEYS{} }
func (m *BgpNexthopVrfAfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpNexthopVrfAfBag_KEYS) ProtoMessage()    {}
func (*BgpNexthopVrfAfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_85b73b3380d3d6da, []int{0}
}

func (m *BgpNexthopVrfAfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNexthopVrfAfBag_KEYS.Unmarshal(m, b)
}
func (m *BgpNexthopVrfAfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNexthopVrfAfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpNexthopVrfAfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNexthopVrfAfBag_KEYS.Merge(m, src)
}
func (m *BgpNexthopVrfAfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpNexthopVrfAfBag_KEYS.Size(m)
}
func (m *BgpNexthopVrfAfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNexthopVrfAfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNexthopVrfAfBag_KEYS proto.InternalMessageInfo

func (m *BgpNexthopVrfAfBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNexthopVrfAfBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpNexthopVrfAfBag_KEYS) GetNextHopAfName() string {
	if m != nil {
		return m.NextHopAfName
	}
	return ""
}

type BgpNexthopPerfVrfAfBag struct {
	CriticalNotfCount          uint32   `protobuf:"varint,1,opt,name=critical_notf_count,json=criticalNotfCount,proto3" json:"critical_notf_count,omitempty"`
	NoncriticalNotfCount       uint32   `protobuf:"varint,2,opt,name=noncritical_notf_count,json=noncriticalNotfCount,proto3" json:"noncritical_notf_count,omitempty"`
	LastNotfBestpathDeletes    uint32   `protobuf:"varint,3,opt,name=last_notf_bestpath_deletes,json=lastNotfBestpathDeletes,proto3" json:"last_notf_bestpath_deletes,omitempty"`
	LastNotfBestpathChanges    uint32   `protobuf:"varint,4,opt,name=last_notf_bestpath_changes,json=lastNotfBestpathChanges,proto3" json:"last_notf_bestpath_changes,omitempty"`
	NhSyncRegCalls             uint32   `protobuf:"varint,5,opt,name=nh_sync_reg_calls,json=nhSyncRegCalls,proto3" json:"nh_sync_reg_calls,omitempty"`
	NhaSyncRegCalls            uint32   `protobuf:"varint,6,opt,name=nha_sync_reg_calls,json=nhaSyncRegCalls,proto3" json:"nha_sync_reg_calls,omitempty"`
	NhaSyncUnRegCalls          uint32   `protobuf:"varint,7,opt,name=nha_sync_un_reg_calls,json=nhaSyncUnRegCalls,proto3" json:"nha_sync_un_reg_calls,omitempty"`
	NhPendingRegistrations     uint32   `protobuf:"varint,8,opt,name=nh_pending_registrations,json=nhPendingRegistrations,proto3" json:"nh_pending_registrations,omitempty"`
	NhPeakRegistrations        uint32   `protobuf:"varint,9,opt,name=nh_peak_registrations,json=nhPeakRegistrations,proto3" json:"nh_peak_registrations,omitempty"`
	NhBatchFinishCalls         uint32   `protobuf:"varint,10,opt,name=nh_batch_finish_calls,json=nhBatchFinishCalls,proto3" json:"nh_batch_finish_calls,omitempty"`
	NhFlushTimerCalls          uint32   `protobuf:"varint,11,opt,name=nh_flush_timer_calls,json=nhFlushTimerCalls,proto3" json:"nh_flush_timer_calls,omitempty"`
	NhLastSyncRegTs            uint32   `protobuf:"varint,12,opt,name=nh_last_sync_reg_ts,json=nhLastSyncRegTs,proto3" json:"nh_last_sync_reg_ts,omitempty"`
	NhLastASyncRegTs           uint32   `protobuf:"varint,13,opt,name=nh_last_a_sync_reg_ts,json=nhLastASyncRegTs,proto3" json:"nh_last_a_sync_reg_ts,omitempty"`
	NhLastASyncUnRegTs         uint32   `protobuf:"varint,14,opt,name=nh_last_a_sync_un_reg_ts,json=nhLastASyncUnRegTs,proto3" json:"nh_last_a_sync_un_reg_ts,omitempty"`
	NhLastBatchFinishTs        uint32   `protobuf:"varint,15,opt,name=nh_last_batch_finish_ts,json=nhLastBatchFinishTs,proto3" json:"nh_last_batch_finish_ts,omitempty"`
	NhLastFlushTimerTs         uint32   `protobuf:"varint,16,opt,name=nh_last_flush_timer_ts,json=nhLastFlushTimerTs,proto3" json:"nh_last_flush_timer_ts,omitempty"`
	NhribUpdateCalls           uint32   `protobuf:"varint,17,opt,name=nhrib_update_calls,json=nhribUpdateCalls,proto3" json:"nhrib_update_calls,omitempty"`
	NhribUpdateTime            uint32   `protobuf:"varint,18,opt,name=nhrib_update_time,json=nhribUpdateTime,proto3" json:"nhrib_update_time,omitempty"`
	NexthopsSentToRib          uint32   `protobuf:"varint,19,opt,name=nexthops_sent_to_rib,json=nexthopsSentToRib,proto3" json:"nexthops_sent_to_rib,omitempty"`
	NexthopsResentToRib        uint32   `protobuf:"varint,20,opt,name=nexthops_resent_to_rib,json=nexthopsResentToRib,proto3" json:"nexthops_resent_to_rib,omitempty"`
	NexthopsRemovedFromRib     uint32   `protobuf:"varint,21,opt,name=nexthops_removed_from_rib,json=nexthopsRemovedFromRib,proto3" json:"nexthops_removed_from_rib,omitempty"`
	RibSyncRegistersFailed     uint32   `protobuf:"varint,22,opt,name=rib_sync_registers_failed,json=ribSyncRegistersFailed,proto3" json:"rib_sync_registers_failed,omitempty"`
	RibaSyncRegistersFailed    uint32   `protobuf:"varint,23,opt,name=riba_sync_registers_failed,json=ribaSyncRegistersFailed,proto3" json:"riba_sync_registers_failed,omitempty"`
	RibnhRegFailedForNoTblId   uint32   `protobuf:"varint,24,opt,name=ribnh_reg_failed_for_no_tbl_id,json=ribnhRegFailedForNoTblId,proto3" json:"ribnh_reg_failed_for_no_tbl_id,omitempty"`
	RibnhRegFailedForNoRibConn uint32   `protobuf:"varint,25,opt,name=ribnh_reg_failed_for_no_rib_conn,json=ribnhRegFailedForNoRibConn,proto3" json:"ribnh_reg_failed_for_no_rib_conn,omitempty"`
	TunnelWalkNhNotFound       uint32   `protobuf:"varint,26,opt,name=tunnel_walk_nh_not_found,json=tunnelWalkNhNotFound,proto3" json:"tunnel_walk_nh_not_found,omitempty"`
	TunnelWalkGwNotFound       uint32   `protobuf:"varint,27,opt,name=tunnel_walk_gw_not_found,json=tunnelWalkGwNotFound,proto3" json:"tunnel_walk_gw_not_found,omitempty"`
	TunnelWalkBackupCreates    uint32   `protobuf:"varint,28,opt,name=tunnel_walk_backup_creates,json=tunnelWalkBackupCreates,proto3" json:"tunnel_walk_backup_creates,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *BgpNexthopPerfVrfAfBag) Reset()         { *m = BgpNexthopPerfVrfAfBag{} }
func (m *BgpNexthopPerfVrfAfBag) String() string { return proto.CompactTextString(m) }
func (*BgpNexthopPerfVrfAfBag) ProtoMessage()    {}
func (*BgpNexthopPerfVrfAfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_85b73b3380d3d6da, []int{1}
}

func (m *BgpNexthopPerfVrfAfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNexthopPerfVrfAfBag.Unmarshal(m, b)
}
func (m *BgpNexthopPerfVrfAfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNexthopPerfVrfAfBag.Marshal(b, m, deterministic)
}
func (m *BgpNexthopPerfVrfAfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNexthopPerfVrfAfBag.Merge(m, src)
}
func (m *BgpNexthopPerfVrfAfBag) XXX_Size() int {
	return xxx_messageInfo_BgpNexthopPerfVrfAfBag.Size(m)
}
func (m *BgpNexthopPerfVrfAfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNexthopPerfVrfAfBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNexthopPerfVrfAfBag proto.InternalMessageInfo

func (m *BgpNexthopPerfVrfAfBag) GetCriticalNotfCount() uint32 {
	if m != nil {
		return m.CriticalNotfCount
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNoncriticalNotfCount() uint32 {
	if m != nil {
		return m.NoncriticalNotfCount
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetLastNotfBestpathDeletes() uint32 {
	if m != nil {
		return m.LastNotfBestpathDeletes
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetLastNotfBestpathChanges() uint32 {
	if m != nil {
		return m.LastNotfBestpathChanges
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhSyncRegCalls() uint32 {
	if m != nil {
		return m.NhSyncRegCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhaSyncRegCalls() uint32 {
	if m != nil {
		return m.NhaSyncRegCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhaSyncUnRegCalls() uint32 {
	if m != nil {
		return m.NhaSyncUnRegCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhPendingRegistrations() uint32 {
	if m != nil {
		return m.NhPendingRegistrations
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhPeakRegistrations() uint32 {
	if m != nil {
		return m.NhPeakRegistrations
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhBatchFinishCalls() uint32 {
	if m != nil {
		return m.NhBatchFinishCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhFlushTimerCalls() uint32 {
	if m != nil {
		return m.NhFlushTimerCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhLastSyncRegTs() uint32 {
	if m != nil {
		return m.NhLastSyncRegTs
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhLastASyncRegTs() uint32 {
	if m != nil {
		return m.NhLastASyncRegTs
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhLastASyncUnRegTs() uint32 {
	if m != nil {
		return m.NhLastASyncUnRegTs
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhLastBatchFinishTs() uint32 {
	if m != nil {
		return m.NhLastBatchFinishTs
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhLastFlushTimerTs() uint32 {
	if m != nil {
		return m.NhLastFlushTimerTs
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhribUpdateCalls() uint32 {
	if m != nil {
		return m.NhribUpdateCalls
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNhribUpdateTime() uint32 {
	if m != nil {
		return m.NhribUpdateTime
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNexthopsSentToRib() uint32 {
	if m != nil {
		return m.NexthopsSentToRib
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNexthopsResentToRib() uint32 {
	if m != nil {
		return m.NexthopsResentToRib
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetNexthopsRemovedFromRib() uint32 {
	if m != nil {
		return m.NexthopsRemovedFromRib
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetRibSyncRegistersFailed() uint32 {
	if m != nil {
		return m.RibSyncRegistersFailed
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetRibaSyncRegistersFailed() uint32 {
	if m != nil {
		return m.RibaSyncRegistersFailed
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetRibnhRegFailedForNoTblId() uint32 {
	if m != nil {
		return m.RibnhRegFailedForNoTblId
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetRibnhRegFailedForNoRibConn() uint32 {
	if m != nil {
		return m.RibnhRegFailedForNoRibConn
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetTunnelWalkNhNotFound() uint32 {
	if m != nil {
		return m.TunnelWalkNhNotFound
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetTunnelWalkGwNotFound() uint32 {
	if m != nil {
		return m.TunnelWalkGwNotFound
	}
	return 0
}

func (m *BgpNexthopPerfVrfAfBag) GetTunnelWalkBackupCreates() uint32 {
	if m != nil {
		return m.TunnelWalkBackupCreates
	}
	return 0
}

type BgpNexthopVrfAfBag struct {
	AfName                  string                  `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	GwafName                string                  `protobuf:"bytes,51,opt,name=gwaf_name,json=gwafName,proto3" json:"gwaf_name,omitempty"`
	TotalNexthops           uint32                  `protobuf:"varint,52,opt,name=total_nexthops,json=totalNexthops,proto3" json:"total_nexthops,omitempty"`
	CriticalTriggerDelay    uint32                  `protobuf:"varint,53,opt,name=critical_trigger_delay,json=criticalTriggerDelay,proto3" json:"critical_trigger_delay,omitempty"`
	NonCriticalTriggerDelay uint32                  `protobuf:"varint,54,opt,name=non_critical_trigger_delay,json=nonCriticalTriggerDelay,proto3" json:"non_critical_trigger_delay,omitempty"`
	TableActive             bool                    `protobuf:"varint,55,opt,name=table_active,json=tableActive,proto3" json:"table_active,omitempty"`
	NhRibUp                 bool                    `protobuf:"varint,56,opt,name=nh_rib_up,json=nhRibUp,proto3" json:"nh_rib_up,omitempty"`
	NhRibVersion            uint32                  `protobuf:"varint,57,opt,name=nh_rib_version,json=nhRibVersion,proto3" json:"nh_rib_version,omitempty"`
	NhNexthopVersion        uint32                  `protobuf:"varint,58,opt,name=nh_nexthop_version,json=nhNexthopVersion,proto3" json:"nh_nexthop_version,omitempty"`
	NhTableId               uint32                  `protobuf:"varint,59,opt,name=nh_table_id,json=nhTableId,proto3" json:"nh_table_id,omitempty"`
	EpeTableVersion         uint32                  `protobuf:"varint,60,opt,name=epe_table_version,json=epeTableVersion,proto3" json:"epe_table_version,omitempty"`
	EpeLabelVersion         uint32                  `protobuf:"varint,61,opt,name=epe_label_version,json=epeLabelVersion,proto3" json:"epe_label_version,omitempty"`
	EpeDownloadedVersion    uint32                  `protobuf:"varint,62,opt,name=epe_downloaded_version,json=epeDownloadedVersion,proto3" json:"epe_downloaded_version,omitempty"`
	EpeStandbyVersion       uint32                  `protobuf:"varint,63,opt,name=epe_standby_version,json=epeStandbyVersion,proto3" json:"epe_standby_version,omitempty"`
	PerformanceStatistics   *BgpNexthopPerfVrfAfBag `protobuf:"bytes,64,opt,name=performance_statistics,json=performanceStatistics,proto3" json:"performance_statistics,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                `json:"-"`
	XXX_unrecognized        []byte                  `json:"-"`
	XXX_sizecache           int32                   `json:"-"`
}

func (m *BgpNexthopVrfAfBag) Reset()         { *m = BgpNexthopVrfAfBag{} }
func (m *BgpNexthopVrfAfBag) String() string { return proto.CompactTextString(m) }
func (*BgpNexthopVrfAfBag) ProtoMessage()    {}
func (*BgpNexthopVrfAfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_85b73b3380d3d6da, []int{2}
}

func (m *BgpNexthopVrfAfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNexthopVrfAfBag.Unmarshal(m, b)
}
func (m *BgpNexthopVrfAfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNexthopVrfAfBag.Marshal(b, m, deterministic)
}
func (m *BgpNexthopVrfAfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNexthopVrfAfBag.Merge(m, src)
}
func (m *BgpNexthopVrfAfBag) XXX_Size() int {
	return xxx_messageInfo_BgpNexthopVrfAfBag.Size(m)
}
func (m *BgpNexthopVrfAfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNexthopVrfAfBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNexthopVrfAfBag proto.InternalMessageInfo

func (m *BgpNexthopVrfAfBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpNexthopVrfAfBag) GetGwafName() string {
	if m != nil {
		return m.GwafName
	}
	return ""
}

func (m *BgpNexthopVrfAfBag) GetTotalNexthops() uint32 {
	if m != nil {
		return m.TotalNexthops
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetCriticalTriggerDelay() uint32 {
	if m != nil {
		return m.CriticalTriggerDelay
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetNonCriticalTriggerDelay() uint32 {
	if m != nil {
		return m.NonCriticalTriggerDelay
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetTableActive() bool {
	if m != nil {
		return m.TableActive
	}
	return false
}

func (m *BgpNexthopVrfAfBag) GetNhRibUp() bool {
	if m != nil {
		return m.NhRibUp
	}
	return false
}

func (m *BgpNexthopVrfAfBag) GetNhRibVersion() uint32 {
	if m != nil {
		return m.NhRibVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetNhNexthopVersion() uint32 {
	if m != nil {
		return m.NhNexthopVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetNhTableId() uint32 {
	if m != nil {
		return m.NhTableId
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetEpeTableVersion() uint32 {
	if m != nil {
		return m.EpeTableVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetEpeLabelVersion() uint32 {
	if m != nil {
		return m.EpeLabelVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetEpeDownloadedVersion() uint32 {
	if m != nil {
		return m.EpeDownloadedVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetEpeStandbyVersion() uint32 {
	if m != nil {
		return m.EpeStandbyVersion
	}
	return 0
}

func (m *BgpNexthopVrfAfBag) GetPerformanceStatistics() *BgpNexthopPerfVrfAfBag {
	if m != nil {
		return m.PerformanceStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpNexthopVrfAfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_vrf_af_bag_KEYS")
	proto.RegisterType((*BgpNexthopPerfVrfAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_perf_vrf_af_bag")
	proto.RegisterType((*BgpNexthopVrfAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_vrf_af_bag")
}

func init() { proto.RegisterFile("bgp_nexthop_vrf_af_bag.proto", fileDescriptor_85b73b3380d3d6da) }

var fileDescriptor_85b73b3380d3d6da = []byte{
	// 1132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xdf, 0x6e, 0xe4, 0xb4,
	0x17, 0xc7, 0x35, 0xfb, 0xfb, 0xd1, 0xdd, 0xba, 0x9d, 0x6e, 0x9b, 0xb6, 0x53, 0x6f, 0xbb, 0x5a,
	0x95, 0x02, 0xa2, 0xfc, 0xd1, 0x00, 0x6d, 0x59, 0x76, 0x29, 0x7f, 0xb6, 0x9d, 0x32, 0xb0, 0x62,
	0x35, 0x42, 0xe9, 0x2c, 0x88, 0x2b, 0xcb, 0x49, 0x9c, 0xc4, 0x6a, 0xe6, 0x38, 0xb2, 0x3d, 0xed,
	0xf6, 0x8e, 0x0b, 0x5e, 0x82, 0x17, 0xe1, 0x21, 0x78, 0x18, 0x9e, 0x01, 0xf9, 0x38, 0x49, 0x33,
	0xd5, 0x94, 0x6b, 0xee, 0xa2, 0xf3, 0xfd, 0x7e, 0xec, 0x73, 0x8e, 0xe3, 0x3f, 0xe4, 0x71, 0x94,
	0x95, 0x0c, 0xc4, 0x1b, 0x9b, 0xab, 0x92, 0x5d, 0xea, 0x94, 0xf1, 0x94, 0x45, 0x3c, 0xeb, 0x97,
	0x5a, 0x59, 0x15, 0xfc, 0xd1, 0x89, 0xa5, 0x89, 0x15, 0x93, 0xca, 0xb0, 0x37, 0x9a, 0xc9, 0xf2,
	0xf2, 0x88, 0x39, 0x40, 0x95, 0x42, 0xf7, 0xa3, 0xac, 0xec, 0x4b, 0x30, 0x96, 0x43, 0x2c, 0x4c,
	0xf3, 0xd5, 0x7c, 0x30, 0x1e, 0x5b, 0x79, 0x29, 0xfa, 0x89, 0x48, 0xf9, 0xb4, 0xb0, 0x6e, 0xf0,
	0x3e, 0x4f, 0x4d, 0x9f, 0xa7, 0x7d, 0x37, 0x1f, 0x73, 0x13, 0xf2, 0x24, 0xd1, 0xc2, 0x18, 0x96,
	0xf2, 0x89, 0x2c, 0xa4, 0x30, 0x77, 0x28, 0xd7, 0xad, 0x78, 0x5a, 0x65, 0xb9, 0xf7, 0x7b, 0x87,
	0xec, 0xcc, 0x4f, 0x9e, 0xfd, 0xf8, 0xdd, 0xaf, 0xe7, 0xc1, 0x3b, 0xa4, 0xdb, 0xa4, 0x02, 0x7c,
	0x22, 0x68, 0x67, 0xb7, 0xb3, 0xbf, 0x18, 0x2e, 0xd7, 0xc1, 0x11, 0x9f, 0x88, 0x60, 0x8b, 0xdc,
	0xe7, 0xa9, 0x97, 0xef, 0xa1, 0xbc, 0xc0, 0x53, 0x14, 0xde, 0x27, 0xab, 0xed, 0x39, 0xd1, 0xf1,
	0x3f, 0x74, 0x74, 0x5d, 0xfc, 0x07, 0x55, 0x9e, 0xa0, 0x71, 0xef, 0xb7, 0xe5, 0xd9, 0x34, 0x4a,
	0xa1, 0xd3, 0x56, 0x2e, 0x41, 0x9f, 0xac, 0xc7, 0x5a, 0x5a, 0x19, 0xf3, 0x82, 0x81, 0xb2, 0x29,
	0x8b, 0xd5, 0x14, 0x2c, 0x26, 0xd3, 0x0d, 0xd7, 0x6a, 0x69, 0xa4, 0x6c, 0x3a, 0x70, 0x42, 0x70,
	0x44, 0x7a, 0xa0, 0x60, 0x1e, 0x72, 0x0f, 0x91, 0x8d, 0x96, 0x7a, 0x43, 0x1d, 0x93, 0xed, 0x82,
	0x1b, 0xeb, 0xed, 0x91, 0x30, 0xb6, 0xe4, 0x36, 0x67, 0x89, 0x28, 0x84, 0x15, 0x06, 0x13, 0xef,
	0x86, 0x5b, 0xce, 0xe1, 0x90, 0xd3, 0x4a, 0x3f, 0xf3, 0xf2, 0x1d, 0x70, 0x9c, 0x73, 0xc8, 0x84,
	0xa1, 0xff, 0x9f, 0x0f, 0x0f, 0xbc, 0x1c, 0x7c, 0x40, 0xd6, 0x20, 0x67, 0xe6, 0x1a, 0x62, 0xa6,
	0x45, 0xc6, 0x62, 0x5e, 0x14, 0x86, 0xbe, 0x85, 0xcc, 0x0a, 0xe4, 0xe7, 0xd7, 0x10, 0x87, 0x22,
	0x1b, 0xb8, 0x68, 0xf0, 0x11, 0x09, 0x20, 0xe7, 0xb7, 0xbd, 0x0b, 0xe8, 0x7d, 0x08, 0x39, 0x9f,
	0x31, 0x7f, 0x4a, 0x36, 0x1b, 0xf3, 0x14, 0x5a, 0xfe, 0xfb, 0xbe, 0x73, 0x95, 0xff, 0x35, 0x34,
	0xc4, 0x33, 0x42, 0x21, 0x67, 0xa5, 0x80, 0x44, 0x42, 0xe6, 0x00, 0x69, 0xac, 0xe6, 0x56, 0x2a,
	0x30, 0xf4, 0x01, 0x42, 0x3d, 0xc8, 0x7f, 0xf2, 0x72, 0xd8, 0x56, 0x83, 0x03, 0x37, 0x17, 0x2b,
	0x05, 0xbf, 0xb8, 0x85, 0x2d, 0x22, 0xb6, 0xee, 0x30, 0x7e, 0x31, 0xcb, 0x7c, 0x86, 0x4c, 0xc4,
	0x6d, 0x9c, 0xb3, 0x54, 0x82, 0x34, 0x79, 0x95, 0x1f, 0x41, 0x26, 0x80, 0xfc, 0xd4, 0x69, 0x43,
	0x94, 0x7c, 0x82, 0x9f, 0x90, 0x0d, 0xc8, 0x59, 0x5a, 0x4c, 0x4d, 0xce, 0xac, 0x9c, 0x08, 0x5d,
	0x11, 0x4b, 0x75, 0x45, 0x43, 0x27, 0x8d, 0x9d, 0xe2, 0x81, 0x8f, 0xc9, 0x3a, 0xe4, 0x0c, 0xd7,
	0xa6, 0x69, 0x9a, 0x35, 0x74, 0xb9, 0xee, 0xd8, 0x2b, 0x6e, 0x6c, 0xd5, 0xb4, 0xb1, 0x1b, 0x7e,
	0xb3, 0x76, 0xf3, 0x19, 0x7f, 0x17, 0xfd, 0xab, 0xde, 0x7f, 0x72, 0x03, 0x1c, 0x61, 0xc3, 0xda,
	0x40, 0xd5, 0x68, 0x6b, 0xe8, 0x4a, 0x5d, 0x45, 0xc3, 0x60, 0xa7, 0x91, 0xda, 0xaa, 0xa9, 0x99,
	0xea, 0xad, 0xa1, 0x0f, 0xeb, 0x76, 0x39, 0xa8, 0x55, 0xfe, 0xd8, 0xb5, 0xb8, 0x57, 0x53, 0xed,
	0x06, 0x58, 0x43, 0x57, 0xdb, 0x33, 0xdd, 0x74, 0x60, 0xec, 0xca, 0x0f, 0x20, 0xd7, 0x32, 0x62,
	0xd3, 0x32, 0xe1, 0x56, 0x54, 0xdd, 0x5a, 0xab, 0xab, 0xd1, 0x32, 0x7a, 0x8d, 0x82, 0x6f, 0xd6,
	0x87, 0xee, 0x47, 0x6c, 0xb9, 0xdd, 0x04, 0x34, 0xa8, 0x5b, 0xd5, 0x98, 0xdd, 0xe0, 0xb8, 0x12,
	0x7e, 0xbf, 0x1a, 0x66, 0x04, 0x58, 0x66, 0x15, 0xd3, 0x32, 0xa2, 0xeb, 0xd5, 0x4a, 0x54, 0xda,
	0xb9, 0x00, 0x3b, 0x56, 0xa1, 0x8c, 0x82, 0x43, 0xd2, 0x6b, 0x00, 0x2d, 0xda, 0xc8, 0x46, 0x55,
	0x73, 0xa5, 0x86, 0x28, 0x7a, 0xe8, 0x39, 0x79, 0xd4, 0x82, 0x26, 0xea, 0x52, 0x24, 0x2c, 0xd5,
	0x6a, 0x82, 0xdc, 0x66, 0xf5, 0x47, 0x36, 0x1c, 0xea, 0x43, 0xad, 0x26, 0x15, 0xea, 0x4a, 0xa9,
	0x57, 0x51, 0x1a, 0x2b, 0xb4, 0x3b, 0x0c, 0x65, 0x21, 0x12, 0xda, 0xf3, 0xa8, 0x96, 0x51, 0xb5,
	0x96, 0x5e, 0x1e, 0xa2, 0xea, 0x76, 0xb3, 0x96, 0x11, 0xbf, 0x83, 0xdd, 0xf2, 0xbb, 0xd9, 0x39,
	0xe6, 0xc1, 0x2f, 0xc8, 0x13, 0x2d, 0x23, 0xc8, 0xf1, 0x37, 0xf0, 0x08, 0x4b, 0x95, 0x66, 0xa0,
	0x98, 0x8d, 0x0a, 0x26, 0x13, 0x4a, 0x71, 0x00, 0x8a, 0xae, 0x50, 0x64, 0x9e, 0x1b, 0x2a, 0x3d,
	0x52, 0xe3, 0xa8, 0x78, 0x99, 0x04, 0x67, 0x64, 0xf7, 0xae, 0x11, 0x5c, 0x45, 0xb1, 0x02, 0xa0,
	0x8f, 0x70, 0x8c, 0xed, 0x39, 0x63, 0x84, 0x32, 0x1a, 0x28, 0x80, 0xe0, 0x29, 0xa1, 0x76, 0x0a,
	0x20, 0x0a, 0x76, 0xc5, 0x8b, 0x0b, 0x06, 0xb9, 0x3b, 0x9c, 0x58, 0xaa, 0xa6, 0x90, 0xd0, 0x6d,
	0x7f, 0x0e, 0x7a, 0xfd, 0x17, 0x5e, 0x5c, 0x8c, 0xf2, 0x91, 0xb2, 0x43, 0xa7, 0xdd, 0xe6, 0xb2,
	0xab, 0x16, 0xb7, 0x73, 0x9b, 0xfb, 0xfe, 0xaa, 0xe1, 0x8e, 0xc9, 0x76, 0x9b, 0x8b, 0x78, 0x7c,
	0x31, 0x2d, 0x59, 0xac, 0x05, 0x77, 0xe7, 0xe7, 0x63, 0xdf, 0xb4, 0x1b, 0xf2, 0x14, 0xf5, 0x81,
	0x97, 0xf7, 0xfe, 0x5a, 0x20, 0xbd, 0xf9, 0x37, 0x51, 0xfb, 0x7e, 0x39, 0x98, 0xb9, 0x5f, 0x76,
	0xc8, 0x62, 0x76, 0x55, 0x4b, 0x87, 0x28, 0x3d, 0x70, 0x01, 0x14, 0xdf, 0x23, 0x2b, 0x56, 0x59,
	0x77, 0xfa, 0x57, 0x7f, 0x07, 0x3d, 0xc2, 0x0c, 0xba, 0x18, 0x1d, 0x55, 0x41, 0x77, 0x55, 0x34,
	0xf7, 0x84, 0xd5, 0x32, 0xcb, 0x84, 0x76, 0x47, 0x3e, 0xbf, 0xa6, 0x9f, 0xfb, 0x52, 0x6b, 0x75,
	0xec, 0xc5, 0x33, 0xa7, 0xb9, 0x52, 0x41, 0x01, 0xbb, 0x83, 0x7c, 0xea, 0x4b, 0x05, 0x05, 0x83,
	0x79, 0xf0, 0xdb, 0x64, 0xd9, 0xf2, 0xa8, 0xa8, 0x2f, 0x77, 0xfa, 0xc5, 0x6e, 0x67, 0xff, 0x41,
	0xb8, 0x84, 0xb1, 0x13, 0x0c, 0x05, 0xdb, 0x64, 0xd1, 0xad, 0x3e, 0x6e, 0x44, 0xfa, 0x0c, 0xf5,
	0xfb, 0x90, 0x87, 0x6e, 0xff, 0x05, 0xef, 0x92, 0x95, 0x4a, 0xbb, 0x14, 0xda, 0x48, 0x05, 0xf4,
	0x39, 0xce, 0xb7, 0x8c, 0x86, 0x9f, 0x7d, 0xcc, 0xef, 0xfb, 0x9b, 0x6e, 0x56, 0xce, 0x2f, 0xeb,
	0x7d, 0x5f, 0xd5, 0x5f, 0xbb, 0x9f, 0x90, 0x25, 0xc8, 0x99, 0xcf, 0x4a, 0x26, 0xf4, 0x18, 0x6d,
	0x8b, 0x90, 0x8f, 0x5d, 0xe4, 0x65, 0xe2, 0xce, 0x05, 0x51, 0x8a, 0xca, 0x50, 0x0f, 0xf6, 0x95,
	0x3f, 0x17, 0x44, 0x29, 0xd0, 0x56, 0x8f, 0x55, 0x79, 0x0b, 0x1e, 0x89, 0xa2, 0xf1, 0x7e, 0xdd,
	0x78, 0x5f, 0xb9, 0x78, 0xed, 0x3d, 0x22, 0x3d, 0xe7, 0x4d, 0xd4, 0x15, 0x14, 0x8a, 0x27, 0x22,
	0x69, 0x80, 0x6f, 0x7c, 0xf7, 0x45, 0x29, 0xce, 0x1a, 0xb1, 0xa6, 0xfa, 0x64, 0xdd, 0x51, 0xee,
	0x09, 0x92, 0x44, 0xd7, 0x0d, 0xf2, 0xad, 0x3f, 0x78, 0x44, 0x29, 0xce, 0xbd, 0x52, 0xfb, 0xff,
	0xee, 0x90, 0x9e, 0x7b, 0x52, 0x28, 0x3d, 0xc1, 0x97, 0x8c, 0xb1, 0xdc, 0x4a, 0x63, 0x65, 0x6c,
	0xe8, 0x8b, 0xdd, 0xce, 0xfe, 0xd2, 0xc1, 0x9f, 0x9d, 0xfe, 0x7f, 0xf6, 0x8d, 0xd6, 0xff, 0x97,
	0x87, 0x51, 0xb8, 0xd9, 0x2a, 0xeb, 0xbc, 0xa9, 0x2a, 0x5a, 0xc0, 0x97, 0xe7, 0xe1, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0xa5, 0x4b, 0x0c, 0x99, 0x0a, 0x00, 0x00,
}
