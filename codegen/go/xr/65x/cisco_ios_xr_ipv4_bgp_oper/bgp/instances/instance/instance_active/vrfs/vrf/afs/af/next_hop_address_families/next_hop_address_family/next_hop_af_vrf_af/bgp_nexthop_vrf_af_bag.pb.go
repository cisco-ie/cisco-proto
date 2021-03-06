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

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_next_hop_address_families_next_hop_address_family_next_hop_af_vrf_af

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	NextHopAfName        string   `protobuf:"bytes,4,opt,name=next_hop_af_name,json=nextHopAfName,proto3" json:"next_hop_af_name,omitempty"`
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

func (m *BgpNexthopVrfAfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*BgpNexthopVrfAfBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_vrf_af_bag_KEYS")
	proto.RegisterType((*BgpNexthopPerfVrfAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_perf_vrf_af_bag")
	proto.RegisterType((*BgpNexthopVrfAfBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.next_hop_address_families.next_hop_address_family.next_hop_af_vrf_af.bgp_nexthop_vrf_af_bag")
}

func init() { proto.RegisterFile("bgp_nexthop_vrf_af_bag.proto", fileDescriptor_85b73b3380d3d6da) }

var fileDescriptor_85b73b3380d3d6da = []byte{
	// 1142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xed, 0x6e, 0x1c, 0x35,
	0x17, 0xc7, 0xb5, 0x7d, 0x1e, 0xf2, 0xe2, 0x64, 0xd3, 0x76, 0xd2, 0x6c, 0xdc, 0xb4, 0xaa, 0x42,
	0x00, 0x11, 0x5e, 0xb4, 0x40, 0x12, 0x4a, 0x4b, 0x78, 0x69, 0xb2, 0x61, 0xa1, 0xa2, 0x5a, 0xa1,
	0xc9, 0x16, 0xc4, 0x27, 0xcb, 0x33, 0xe3, 0x99, 0xb1, 0x32, 0x7b, 0x3c, 0xb2, 0xbd, 0x9b, 0xe6,
	0x1b, 0x57, 0xc1, 0x57, 0x2e, 0x82, 0x8b, 0xe0, 0x62, 0xb8, 0x09, 0xe4, 0xe3, 0x99, 0xc9, 0x6c,
	0xb4, 0xe1, 0x33, 0x5f, 0x56, 0xab, 0xf3, 0xff, 0xff, 0xec, 0x73, 0x8e, 0xc7, 0x2f, 0xe4, 0x71,
	0x94, 0x95, 0x0c, 0xc4, 0x1b, 0x9b, 0xab, 0x92, 0xcd, 0x74, 0xca, 0x78, 0xca, 0x22, 0x9e, 0xf5,
	0x4b, 0xad, 0xac, 0x0a, 0x7e, 0xef, 0xc4, 0xd2, 0xc4, 0x8a, 0x49, 0x65, 0xd8, 0x1b, 0xcd, 0x64,
	0x39, 0x3b, 0x62, 0x0e, 0x50, 0xa5, 0xd0, 0xfd, 0x28, 0x2b, 0xfb, 0x12, 0x8c, 0xe5, 0x10, 0x0b,
	0xd3, 0xfc, 0x6b, 0xfe, 0x30, 0x1e, 0x5b, 0x39, 0x13, 0xfd, 0x99, 0x4e, 0x8d, 0xfb, 0xe9, 0xf3,
	0xd4, 0xf4, 0x79, 0xda, 0x77, 0x93, 0x31, 0x37, 0x1b, 0x4f, 0x12, 0x2d, 0x8c, 0x61, 0x29, 0x9f,
	0xc8, 0x42, 0x0a, 0x73, 0x8b, 0x72, 0xd5, 0x8a, 0xa7, 0x55, 0x8a, 0x7b, 0x7f, 0x74, 0xc8, 0xa3,
	0xc5, 0x99, 0xb3, 0x1f, 0xbf, 0xfb, 0xf5, 0x3c, 0x78, 0x87, 0x74, 0x9b, 0x3c, 0x80, 0x4f, 0x04,
	0xed, 0xec, 0x76, 0xf6, 0x57, 0xc3, 0xf5, 0x3a, 0x38, 0xe2, 0x13, 0x11, 0x3c, 0x24, 0x2b, 0x8e,
	0x43, 0xfd, 0x0e, 0xea, 0xcb, 0x33, 0x9d, 0xa2, 0xb4, 0x4d, 0x96, 0x79, 0xa5, 0xfc, 0x0f, 0x95,
	0x25, 0xee, 0x85, 0xf7, 0xc9, 0xbd, 0x76, 0x3a, 0xe8, 0xf8, 0x3f, 0x3a, 0xba, 0x2e, 0xfe, 0x83,
	0x2a, 0x4f, 0xd0, 0xb8, 0xf7, 0xdb, 0xfa, 0x7c, 0x86, 0xa5, 0xd0, 0x69, 0x2b, 0xcd, 0xa0, 0x4f,
	0x36, 0x63, 0x2d, 0xad, 0x8c, 0x79, 0xc1, 0x40, 0xd9, 0x94, 0xc5, 0x6a, 0x0a, 0x16, 0xf3, 0xec,
	0x86, 0xf7, 0x6b, 0x69, 0xa4, 0x6c, 0x3a, 0x70, 0x42, 0x70, 0x44, 0x7a, 0xa0, 0x60, 0x11, 0x72,
	0x07, 0x91, 0x07, 0x2d, 0xf5, 0x9a, 0x3a, 0x26, 0x3b, 0x05, 0x37, 0xd6, 0xdb, 0x23, 0x61, 0x6c,
	0xc9, 0x6d, 0xce, 0x12, 0x51, 0x08, 0x2b, 0x0c, 0x96, 0xd6, 0x0d, 0xb7, 0x9d, 0xc3, 0x21, 0xa7,
	0x95, 0x7e, 0xe6, 0xe5, 0x5b, 0xe0, 0x38, 0xe7, 0x90, 0x09, 0x83, 0x55, 0x2f, 0x80, 0x07, 0x5e,
	0x0e, 0x3e, 0x20, 0xf7, 0x21, 0x67, 0xe6, 0x0a, 0x62, 0xa6, 0x45, 0xc6, 0x62, 0x5e, 0x14, 0x86,
	0xbe, 0x85, 0xcc, 0x06, 0xe4, 0xe7, 0x57, 0x10, 0x87, 0x22, 0x1b, 0xb8, 0x68, 0xf0, 0x11, 0x09,
	0x20, 0xe7, 0x37, 0xbd, 0x4b, 0xe8, 0xbd, 0x0b, 0x39, 0x9f, 0x33, 0x7f, 0x4a, 0xb6, 0x1a, 0xf3,
	0x14, 0x5a, 0xfe, 0x65, 0xdf, 0xb9, 0xca, 0xff, 0x1a, 0x1a, 0xe2, 0x19, 0xa1, 0x90, 0xb3, 0x52,
	0x40, 0x22, 0x21, 0x73, 0x80, 0x34, 0x56, 0x73, 0x2b, 0x15, 0x18, 0xba, 0x82, 0x50, 0x0f, 0xf2,
	0x9f, 0xbc, 0x1c, 0xb6, 0xd5, 0xe0, 0xc0, 0xcd, 0xc5, 0x4a, 0xc1, 0x2f, 0x6e, 0x60, 0xab, 0x88,
	0x6d, 0x3a, 0x8c, 0x5f, 0xcc, 0x33, 0x9f, 0x21, 0x13, 0x71, 0x1b, 0xe7, 0x2c, 0x95, 0x20, 0x4d,
	0x5e, 0xe5, 0x47, 0x90, 0x09, 0x20, 0x3f, 0x75, 0xda, 0x10, 0x25, 0x9f, 0xe0, 0x27, 0xe4, 0x01,
	0xe4, 0x2c, 0x2d, 0xa6, 0x26, 0x67, 0x56, 0x4e, 0x84, 0xae, 0x88, 0xb5, 0xba, 0xa2, 0xa1, 0x93,
	0xc6, 0x4e, 0xf1, 0xc0, 0xc7, 0x64, 0x13, 0x72, 0x86, 0x6b, 0xd3, 0x34, 0xcd, 0x1a, 0xba, 0x5e,
	0x77, 0xec, 0x15, 0x37, 0xb6, 0x6a, 0xda, 0xd8, 0x0d, 0xbf, 0x55, 0xbb, 0xf9, 0x9c, 0xbf, 0x8b,
	0xfe, 0x7b, 0xde, 0x7f, 0x72, 0x0d, 0x1c, 0x61, 0xc3, 0xda, 0x40, 0xd5, 0x68, 0x6b, 0xe8, 0x46,
	0x5d, 0x45, 0xc3, 0x60, 0xa7, 0x91, 0xda, 0xae, 0xa9, 0xb9, 0xea, 0xad, 0xa1, 0x77, 0xeb, 0x76,
	0x39, 0xa8, 0x55, 0xfe, 0xd8, 0xb5, 0xb8, 0x57, 0x53, 0xed, 0x06, 0x58, 0x43, 0xef, 0xb5, 0x67,
	0xba, 0xee, 0xc0, 0xd8, 0x95, 0x1f, 0x40, 0xae, 0x65, 0xc4, 0xa6, 0x65, 0xc2, 0xad, 0xa8, 0xba,
	0x75, 0xbf, 0xae, 0x46, 0xcb, 0xe8, 0x35, 0x0a, 0xbe, 0x59, 0x1f, 0xba, 0x0f, 0xb1, 0xe5, 0x76,
	0x13, 0xd0, 0xa0, 0x6e, 0x55, 0x63, 0x76, 0x83, 0xe3, 0x4a, 0xf8, 0xfd, 0x6a, 0x98, 0x11, 0x60,
	0x99, 0x55, 0x4c, 0xcb, 0x88, 0x6e, 0x56, 0x2b, 0x51, 0x69, 0xe7, 0x02, 0xec, 0x58, 0x85, 0x32,
	0x0a, 0x0e, 0x49, 0xaf, 0x01, 0xb4, 0x68, 0x23, 0x0f, 0xaa, 0x9a, 0x2b, 0x35, 0x44, 0xd1, 0x43,
	0xcf, 0xc9, 0xc3, 0x16, 0x34, 0x51, 0x33, 0x91, 0xb0, 0x54, 0xab, 0x09, 0x72, 0x5b, 0xd5, 0x17,
	0xd9, 0x70, 0xa8, 0x0f, 0xb5, 0x9a, 0x54, 0xa8, 0x2b, 0xa5, 0x5e, 0x45, 0x69, 0xac, 0xd0, 0xee,
	0x9c, 0x94, 0x85, 0x48, 0x68, 0xcf, 0xa3, 0x5a, 0x46, 0xd5, 0x5a, 0x7a, 0x79, 0x88, 0xaa, 0xdb,
	0xcd, 0x5a, 0x46, 0xfc, 0x16, 0x76, 0xdb, 0xef, 0x66, 0xe7, 0x58, 0x04, 0xbf, 0x20, 0x4f, 0xb4,
	0x8c, 0x20, 0xc7, 0xcf, 0xc0, 0x23, 0x2c, 0x55, 0x9a, 0x81, 0x62, 0x36, 0x2a, 0x98, 0x4c, 0x28,
	0xc5, 0x01, 0x28, 0xba, 0x42, 0x91, 0x79, 0x6e, 0xa8, 0xf4, 0x48, 0x8d, 0xa3, 0xe2, 0x65, 0x12,
	0x9c, 0x91, 0xdd, 0xdb, 0x46, 0x70, 0x15, 0xc5, 0x0a, 0x80, 0x3e, 0xc4, 0x31, 0x76, 0x16, 0x8c,
	0x11, 0xca, 0x68, 0xa0, 0x00, 0x82, 0xa7, 0x84, 0xda, 0x29, 0x80, 0x28, 0xd8, 0x25, 0x2f, 0x2e,
	0x18, 0xe4, 0xee, 0x70, 0x62, 0xa9, 0x9a, 0x42, 0x42, 0x77, 0xfc, 0x39, 0xe8, 0xf5, 0x5f, 0x78,
	0x71, 0x31, 0xca, 0x47, 0xca, 0x0e, 0x9d, 0x76, 0x93, 0xcb, 0x2e, 0x5b, 0xdc, 0xa3, 0x9b, 0xdc,
	0xf7, 0x97, 0x0d, 0x77, 0x4c, 0x76, 0xda, 0x5c, 0xc4, 0xe3, 0x8b, 0x69, 0xc9, 0x62, 0x2d, 0xb8,
	0x3b, 0x3f, 0x1f, 0xfb, 0xa6, 0x5d, 0x93, 0xa7, 0xa8, 0x0f, 0xbc, 0xbc, 0xf7, 0xd7, 0x12, 0xe9,
	0x2d, 0xbe, 0xa4, 0xda, 0xf7, 0xcb, 0xc1, 0xdc, 0xfd, 0xf2, 0x88, 0xac, 0x66, 0x97, 0xb5, 0x74,
	0x88, 0xd2, 0x8a, 0x0b, 0xa0, 0xf8, 0x1e, 0xd9, 0xb0, 0xca, 0xba, 0xd3, 0xbf, 0xfa, 0x3a, 0xe8,
	0x11, 0x66, 0xd0, 0xc5, 0xe8, 0xa8, 0x0a, 0xba, 0xab, 0xa2, 0xb9, 0x27, 0xac, 0x96, 0x59, 0x26,
	0xb4, 0x3b, 0xf2, 0xf9, 0x15, 0xfd, 0xdc, 0x97, 0x5a, 0xab, 0x63, 0x2f, 0x9e, 0x39, 0xcd, 0x95,
	0x0a, 0x0a, 0xd8, 0x2d, 0xe4, 0x53, 0x5f, 0x2a, 0x28, 0x18, 0x2c, 0x82, 0xdf, 0x26, 0xeb, 0x96,
	0x47, 0x45, 0x7d, 0xe9, 0xd3, 0x2f, 0x76, 0x3b, 0xfb, 0x2b, 0xe1, 0x1a, 0xc6, 0x4e, 0x30, 0x14,
	0xec, 0x90, 0x55, 0xb7, 0xfa, 0xb8, 0x11, 0xe9, 0x33, 0xd4, 0x97, 0x21, 0x0f, 0xdd, 0xfe, 0x0b,
	0xde, 0x25, 0x1b, 0x95, 0x36, 0x13, 0xda, 0x48, 0x05, 0xf4, 0x39, 0xce, 0xb7, 0x8e, 0x86, 0x9f,
	0x7d, 0xcc, 0xef, 0xfb, 0xeb, 0x6e, 0x56, 0xce, 0x2f, 0xeb, 0x7d, 0x5f, 0xd5, 0x5f, 0xbb, 0x9f,
	0x90, 0x35, 0xc8, 0x99, 0xcf, 0x4a, 0x26, 0xf4, 0x18, 0x6d, 0xab, 0x90, 0x8f, 0x5d, 0xe4, 0x65,
	0xe2, 0xce, 0x05, 0x51, 0x8a, 0xca, 0x50, 0x0f, 0xf6, 0x95, 0x3f, 0x17, 0x44, 0x29, 0xd0, 0x56,
	0x8f, 0x55, 0x79, 0x0b, 0x1e, 0x89, 0xa2, 0xf1, 0x7e, 0xdd, 0x78, 0x5f, 0xb9, 0x78, 0xed, 0x3d,
	0x22, 0x3d, 0xe7, 0x4d, 0xd4, 0x25, 0x14, 0x8a, 0x27, 0x22, 0x69, 0x80, 0x6f, 0x7c, 0xf7, 0x45,
	0x29, 0xce, 0x1a, 0xb1, 0xa6, 0xfa, 0x64, 0xd3, 0x51, 0xee, 0x75, 0x92, 0x44, 0x57, 0x0d, 0xf2,
	0xad, 0x3f, 0x78, 0x44, 0x29, 0xce, 0xbd, 0x52, 0xfb, 0xff, 0xee, 0x90, 0x9e, 0x7b, 0x52, 0x28,
	0x3d, 0xc1, 0x47, 0x8e, 0xb1, 0xdc, 0x4a, 0x63, 0x65, 0x6c, 0xe8, 0x8b, 0xdd, 0xce, 0xfe, 0xda,
	0xc1, 0x9f, 0x9d, 0xfe, 0x7f, 0xf3, 0xed, 0xd6, 0xff, 0x97, 0x57, 0x51, 0xb8, 0xd5, 0xaa, 0xe9,
	0xbc, 0x29, 0x29, 0x5a, 0xc2, 0xe7, 0xe8, 0xe1, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xba,
	0xe8, 0x99, 0xae, 0x0a, 0x00, 0x00,
}
