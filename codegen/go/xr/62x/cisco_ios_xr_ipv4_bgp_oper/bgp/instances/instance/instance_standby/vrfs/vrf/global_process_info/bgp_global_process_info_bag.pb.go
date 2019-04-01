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
// source: bgp_global_process_info_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_global_process_info

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

type BgpGlobalProcessInfoBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpGlobalProcessInfoBag_KEYS) Reset()         { *m = BgpGlobalProcessInfoBag_KEYS{} }
func (m *BgpGlobalProcessInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoBag_KEYS) ProtoMessage()    {}
func (*BgpGlobalProcessInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_99316a6dde1b0219, []int{0}
}

func (m *BgpGlobalProcessInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS.Unmarshal(m, b)
}
func (m *BgpGlobalProcessInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpGlobalProcessInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS.Merge(m, src)
}
func (m *BgpGlobalProcessInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS.Size(m)
}
func (m *BgpGlobalProcessInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpGlobalProcessInfoBag_KEYS proto.InternalMessageInfo

func (m *BgpGlobalProcessInfoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpGlobalProcessInfoBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type ClusterIdBag_ struct {
	ClusterIdVal         uint32   `protobuf:"varint,1,opt,name=cluster_id_val,json=clusterIdVal,proto3" json:"cluster_id_val,omitempty"`
	ClusterIdType        uint32   `protobuf:"varint,2,opt,name=cluster_id_type,json=clusterIdType,proto3" json:"cluster_id_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterIdBag_) Reset()         { *m = ClusterIdBag_{} }
func (m *ClusterIdBag_) String() string { return proto.CompactTextString(m) }
func (*ClusterIdBag_) ProtoMessage()    {}
func (*ClusterIdBag_) Descriptor() ([]byte, []int) {
	return fileDescriptor_99316a6dde1b0219, []int{1}
}

func (m *ClusterIdBag_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterIdBag_.Unmarshal(m, b)
}
func (m *ClusterIdBag_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterIdBag_.Marshal(b, m, deterministic)
}
func (m *ClusterIdBag_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterIdBag_.Merge(m, src)
}
func (m *ClusterIdBag_) XXX_Size() int {
	return xxx_messageInfo_ClusterIdBag_.Size(m)
}
func (m *ClusterIdBag_) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterIdBag_.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterIdBag_ proto.InternalMessageInfo

func (m *ClusterIdBag_) GetClusterIdVal() uint32 {
	if m != nil {
		return m.ClusterIdVal
	}
	return 0
}

func (m *ClusterIdBag_) GetClusterIdType() uint32 {
	if m != nil {
		return m.ClusterIdType
	}
	return 0
}

type BgpGlobalProcessInfoGbl_ struct {
	InStandaloneMode          bool             `protobuf:"varint,1,opt,name=in_standalone_mode,json=inStandaloneMode,proto3" json:"in_standalone_mode,omitempty"`
	LocalAs                   uint32           `protobuf:"varint,2,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	InstanceName              string           `protobuf:"bytes,3,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	RestartCount              uint32           `protobuf:"varint,4,opt,name=restart_count,json=restartCount,proto3" json:"restart_count,omitempty"`
	UpdateDelay               uint32           `protobuf:"varint,5,opt,name=update_delay,json=updateDelay,proto3" json:"update_delay,omitempty"`
	GenericScanPeriod         uint32           `protobuf:"varint,6,opt,name=generic_scan_period,json=genericScanPeriod,proto3" json:"generic_scan_period,omitempty"`
	ConfederationId           uint32           `protobuf:"varint,7,opt,name=confederation_id,json=confederationId,proto3" json:"confederation_id,omitempty"`
	ClusterId                 uint32           `protobuf:"varint,8,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	ConfiguredClusterId       uint32           `protobuf:"varint,9,opt,name=configured_cluster_id,json=configuredClusterId,proto3" json:"configured_cluster_id,omitempty"`
	IsClusterIdSpecifiedAsIp  bool             `protobuf:"varint,10,opt,name=is_cluster_id_specified_as_ip,json=isClusterIdSpecifiedAsIp,proto3" json:"is_cluster_id_specified_as_ip,omitempty"`
	ClusterIdEntry            []*ClusterIdBag_ `protobuf:"bytes,11,rep,name=cluster_id_entry,json=clusterIdEntry,proto3" json:"cluster_id_entry,omitempty"`
	AsnFormat                 uint32           `protobuf:"varint,12,opt,name=asn_format,json=asnFormat,proto3" json:"asn_format,omitempty"`
	SrgbStartConfigured       uint32           `protobuf:"varint,13,opt,name=srgb_start_configured,json=srgbStartConfigured,proto3" json:"srgb_start_configured,omitempty"`
	SrgbEndConfigured         uint32           `protobuf:"varint,14,opt,name=srgb_end_configured,json=srgbEndConfigured,proto3" json:"srgb_end_configured,omitempty"`
	SrgbStart                 uint32           `protobuf:"varint,15,opt,name=srgb_start,json=srgbStart,proto3" json:"srgb_start,omitempty"`
	SrgbEnd                   uint32           `protobuf:"varint,16,opt,name=srgb_end,json=srgbEnd,proto3" json:"srgb_end,omitempty"`
	GracefulMaintenance       bool             `protobuf:"varint,17,opt,name=graceful_maintenance,json=gracefulMaintenance,proto3" json:"graceful_maintenance,omitempty"`
	GracefulMaintAllNbrs      bool             `protobuf:"varint,18,opt,name=graceful_maint_all_nbrs,json=gracefulMaintAllNbrs,proto3" json:"graceful_maint_all_nbrs,omitempty"`
	GracefulMaintRetainRoutes bool             `protobuf:"varint,19,opt,name=graceful_maint_retain_routes,json=gracefulMaintRetainRoutes,proto3" json:"graceful_maint_retain_routes,omitempty"`
	ProcessRlimit             uint64           `protobuf:"varint,20,opt,name=process_rlimit,json=processRlimit,proto3" json:"process_rlimit,omitempty"`
	BmpMaximumBufferSize      uint64           `protobuf:"varint,21,opt,name=bmp_maximum_buffer_size,json=bmpMaximumBufferSize,proto3" json:"bmp_maximum_buffer_size,omitempty"`
	BmpDefaultBufferSize      uint64           `protobuf:"varint,22,opt,name=bmp_default_buffer_size,json=bmpDefaultBufferSize,proto3" json:"bmp_default_buffer_size,omitempty"`
	BmpCurrentBufferSize      uint64           `protobuf:"varint,23,opt,name=bmp_current_buffer_size,json=bmpCurrentBufferSize,proto3" json:"bmp_current_buffer_size,omitempty"`
	BmpCurMaximumBufferSize   uint64           `protobuf:"varint,24,opt,name=bmp_cur_maximum_buffer_size,json=bmpCurMaximumBufferSize,proto3" json:"bmp_cur_maximum_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}         `json:"-"`
	XXX_unrecognized          []byte           `json:"-"`
	XXX_sizecache             int32            `json:"-"`
}

func (m *BgpGlobalProcessInfoGbl_) Reset()         { *m = BgpGlobalProcessInfoGbl_{} }
func (m *BgpGlobalProcessInfoGbl_) String() string { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoGbl_) ProtoMessage()    {}
func (*BgpGlobalProcessInfoGbl_) Descriptor() ([]byte, []int) {
	return fileDescriptor_99316a6dde1b0219, []int{2}
}

func (m *BgpGlobalProcessInfoGbl_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpGlobalProcessInfoGbl_.Unmarshal(m, b)
}
func (m *BgpGlobalProcessInfoGbl_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpGlobalProcessInfoGbl_.Marshal(b, m, deterministic)
}
func (m *BgpGlobalProcessInfoGbl_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpGlobalProcessInfoGbl_.Merge(m, src)
}
func (m *BgpGlobalProcessInfoGbl_) XXX_Size() int {
	return xxx_messageInfo_BgpGlobalProcessInfoGbl_.Size(m)
}
func (m *BgpGlobalProcessInfoGbl_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpGlobalProcessInfoGbl_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpGlobalProcessInfoGbl_ proto.InternalMessageInfo

func (m *BgpGlobalProcessInfoGbl_) GetInStandaloneMode() bool {
	if m != nil {
		return m.InStandaloneMode
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpGlobalProcessInfoGbl_) GetRestartCount() uint32 {
	if m != nil {
		return m.RestartCount
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetUpdateDelay() uint32 {
	if m != nil {
		return m.UpdateDelay
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetGenericScanPeriod() uint32 {
	if m != nil {
		return m.GenericScanPeriod
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetConfederationId() uint32 {
	if m != nil {
		return m.ConfederationId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetClusterId() uint32 {
	if m != nil {
		return m.ClusterId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetConfiguredClusterId() uint32 {
	if m != nil {
		return m.ConfiguredClusterId
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetIsClusterIdSpecifiedAsIp() bool {
	if m != nil {
		return m.IsClusterIdSpecifiedAsIp
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetClusterIdEntry() []*ClusterIdBag_ {
	if m != nil {
		return m.ClusterIdEntry
	}
	return nil
}

func (m *BgpGlobalProcessInfoGbl_) GetAsnFormat() uint32 {
	if m != nil {
		return m.AsnFormat
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbStartConfigured() uint32 {
	if m != nil {
		return m.SrgbStartConfigured
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbEndConfigured() uint32 {
	if m != nil {
		return m.SrgbEndConfigured
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbStart() uint32 {
	if m != nil {
		return m.SrgbStart
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetSrgbEnd() uint32 {
	if m != nil {
		return m.SrgbEnd
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintenance() bool {
	if m != nil {
		return m.GracefulMaintenance
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintAllNbrs() bool {
	if m != nil {
		return m.GracefulMaintAllNbrs
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetGracefulMaintRetainRoutes() bool {
	if m != nil {
		return m.GracefulMaintRetainRoutes
	}
	return false
}

func (m *BgpGlobalProcessInfoGbl_) GetProcessRlimit() uint64 {
	if m != nil {
		return m.ProcessRlimit
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpMaximumBufferSize() uint64 {
	if m != nil {
		return m.BmpMaximumBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpDefaultBufferSize() uint64 {
	if m != nil {
		return m.BmpDefaultBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpCurrentBufferSize() uint64 {
	if m != nil {
		return m.BmpCurrentBufferSize
	}
	return 0
}

func (m *BgpGlobalProcessInfoGbl_) GetBmpCurMaximumBufferSize() uint64 {
	if m != nil {
		return m.BmpCurMaximumBufferSize
	}
	return 0
}

type BgpGlobalProcessInfoVrf_ struct {
	VrfIsActive                                  bool     `protobuf:"varint,1,opt,name=vrf_is_active,json=vrfIsActive,proto3" json:"vrf_is_active,omitempty"`
	RouteDistinguisher                           []uint32 `protobuf:"varint,2,rep,packed,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	RouterId                                     string   `protobuf:"bytes,3,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	ConfiguredRouterId                           string   `protobuf:"bytes,4,opt,name=configured_router_id,json=configuredRouterId,proto3" json:"configured_router_id,omitempty"`
	IsRedistributeIbgpToIgPsEnabled              bool     `protobuf:"varint,5,opt,name=is_redistribute_ibgp_to_ig_ps_enabled,json=isRedistributeIbgpToIgPsEnabled,proto3" json:"is_redistribute_ibgp_to_ig_ps_enabled,omitempty"`
	IsFastExternalFalloverEnabled                bool     `protobuf:"varint,6,opt,name=is_fast_external_fallover_enabled,json=isFastExternalFalloverEnabled,proto3" json:"is_fast_external_fallover_enabled,omitempty"`
	IsBestpathMissingMedIsWorstEnabled           bool     `protobuf:"varint,7,opt,name=is_bestpath_missing_med_is_worst_enabled,json=isBestpathMissingMedIsWorstEnabled,proto3" json:"is_bestpath_missing_med_is_worst_enabled,omitempty"`
	IsBestpathAlwaysCompareMedEnabled            bool     `protobuf:"varint,8,opt,name=is_bestpath_always_compare_med_enabled,json=isBestpathAlwaysCompareMedEnabled,proto3" json:"is_bestpath_always_compare_med_enabled,omitempty"`
	IsBestpathIgnoreAsPathEnabled                bool     `protobuf:"varint,9,opt,name=is_bestpath_ignore_as_path_enabled,json=isBestpathIgnoreAsPathEnabled,proto3" json:"is_bestpath_ignore_as_path_enabled,omitempty"`
	IsBestpathAsPathMpathRelaxEnabled            bool     `protobuf:"varint,10,opt,name=is_bestpath_as_path_mpath_relax_enabled,json=isBestpathAsPathMpathRelaxEnabled,proto3" json:"is_bestpath_as_path_mpath_relax_enabled,omitempty"`
	IsBestpathCompareMedFromConfedPeerEnabled    bool     `protobuf:"varint,11,opt,name=is_bestpath_compare_med_from_confed_peer_enabled,json=isBestpathCompareMedFromConfedPeerEnabled,proto3" json:"is_bestpath_compare_med_from_confed_peer_enabled,omitempty"`
	IsBestpathCompareRouterIdForEbgpPeersEnabled bool     `protobuf:"varint,12,opt,name=is_bestpath_compare_router_id_for_ebgp_peers_enabled,json=isBestpathCompareRouterIdForEbgpPeersEnabled,proto3" json:"is_bestpath_compare_router_id_for_ebgp_peers_enabled,omitempty"`
	IsBestpathAigpIgnoreEnabled                  bool     `protobuf:"varint,13,opt,name=is_bestpath_aigp_ignore_enabled,json=isBestpathAigpIgnoreEnabled,proto3" json:"is_bestpath_aigp_ignore_enabled,omitempty"`
	IsMultipathAsPathIgnoreOnwardsEnabled        bool     `protobuf:"varint,14,opt,name=is_multipath_as_path_ignore_onwards_enabled,json=isMultipathAsPathIgnoreOnwardsEnabled,proto3" json:"is_multipath_as_path_ignore_onwards_enabled,omitempty"`
	IsEnforceFirstAsEnabled                      bool     `protobuf:"varint,15,opt,name=is_enforce_first_as_enabled,json=isEnforceFirstAsEnabled,proto3" json:"is_enforce_first_as_enabled,omitempty"`
	DefaultLocalPreference                       uint32   `protobuf:"varint,16,opt,name=default_local_preference,json=defaultLocalPreference,proto3" json:"default_local_preference,omitempty"`
	KeepAliveTime                                uint32   `protobuf:"varint,17,opt,name=keep_alive_time,json=keepAliveTime,proto3" json:"keep_alive_time,omitempty"`
	HoldTime                                     uint32   `protobuf:"varint,18,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	MinAcceptableHoldTime                        uint32   `protobuf:"varint,19,opt,name=min_acceptable_hold_time,json=minAcceptableHoldTime,proto3" json:"min_acceptable_hold_time,omitempty"`
	IsNeighborLogging                            bool     `protobuf:"varint,20,opt,name=is_neighbor_logging,json=isNeighborLogging,proto3" json:"is_neighbor_logging,omitempty"`
	IsDefaultMetricConfigured                    bool     `protobuf:"varint,21,opt,name=is_default_metric_configured,json=isDefaultMetricConfigured,proto3" json:"is_default_metric_configured,omitempty"`
	DefaultMetric                                uint32   `protobuf:"varint,22,opt,name=default_metric,json=defaultMetric,proto3" json:"default_metric,omitempty"`
	IsDefaultOriginateConfigured                 bool     `protobuf:"varint,23,opt,name=is_default_originate_configured,json=isDefaultOriginateConfigured,proto3" json:"is_default_originate_configured,omitempty"`
	IsGracefulRestart                            bool     `protobuf:"varint,24,opt,name=is_graceful_restart,json=isGracefulRestart,proto3" json:"is_graceful_restart,omitempty"`
	IsNsr                                        bool     `protobuf:"varint,25,opt,name=is_nsr,json=isNsr,proto3" json:"is_nsr,omitempty"`
	RestartTime                                  uint32   `protobuf:"varint,26,opt,name=restart_time,json=restartTime,proto3" json:"restart_time,omitempty"`
	StalePathTime                                uint32   `protobuf:"varint,27,opt,name=stale_path_time,json=stalePathTime,proto3" json:"stale_path_time,omitempty"`
	RibPurgeTimeout                              uint32   `protobuf:"varint,28,opt,name=rib_purge_timeout,json=ribPurgeTimeout,proto3" json:"rib_purge_timeout,omitempty"`
	XXX_NoUnkeyedLiteral                         struct{} `json:"-"`
	XXX_unrecognized                             []byte   `json:"-"`
	XXX_sizecache                                int32    `json:"-"`
}

func (m *BgpGlobalProcessInfoVrf_) Reset()         { *m = BgpGlobalProcessInfoVrf_{} }
func (m *BgpGlobalProcessInfoVrf_) String() string { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoVrf_) ProtoMessage()    {}
func (*BgpGlobalProcessInfoVrf_) Descriptor() ([]byte, []int) {
	return fileDescriptor_99316a6dde1b0219, []int{3}
}

func (m *BgpGlobalProcessInfoVrf_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpGlobalProcessInfoVrf_.Unmarshal(m, b)
}
func (m *BgpGlobalProcessInfoVrf_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpGlobalProcessInfoVrf_.Marshal(b, m, deterministic)
}
func (m *BgpGlobalProcessInfoVrf_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpGlobalProcessInfoVrf_.Merge(m, src)
}
func (m *BgpGlobalProcessInfoVrf_) XXX_Size() int {
	return xxx_messageInfo_BgpGlobalProcessInfoVrf_.Size(m)
}
func (m *BgpGlobalProcessInfoVrf_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpGlobalProcessInfoVrf_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpGlobalProcessInfoVrf_ proto.InternalMessageInfo

func (m *BgpGlobalProcessInfoVrf_) GetVrfIsActive() bool {
	if m != nil {
		return m.VrfIsActive
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetRouteDistinguisher() []uint32 {
	if m != nil {
		return m.RouteDistinguisher
	}
	return nil
}

func (m *BgpGlobalProcessInfoVrf_) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *BgpGlobalProcessInfoVrf_) GetConfiguredRouterId() string {
	if m != nil {
		return m.ConfiguredRouterId
	}
	return ""
}

func (m *BgpGlobalProcessInfoVrf_) GetIsRedistributeIbgpToIgPsEnabled() bool {
	if m != nil {
		return m.IsRedistributeIbgpToIgPsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsFastExternalFalloverEnabled() bool {
	if m != nil {
		return m.IsFastExternalFalloverEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathMissingMedIsWorstEnabled() bool {
	if m != nil {
		return m.IsBestpathMissingMedIsWorstEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAlwaysCompareMedEnabled() bool {
	if m != nil {
		return m.IsBestpathAlwaysCompareMedEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathIgnoreAsPathEnabled() bool {
	if m != nil {
		return m.IsBestpathIgnoreAsPathEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAsPathMpathRelaxEnabled() bool {
	if m != nil {
		return m.IsBestpathAsPathMpathRelaxEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathCompareMedFromConfedPeerEnabled() bool {
	if m != nil {
		return m.IsBestpathCompareMedFromConfedPeerEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathCompareRouterIdForEbgpPeersEnabled() bool {
	if m != nil {
		return m.IsBestpathCompareRouterIdForEbgpPeersEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsBestpathAigpIgnoreEnabled() bool {
	if m != nil {
		return m.IsBestpathAigpIgnoreEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsMultipathAsPathIgnoreOnwardsEnabled() bool {
	if m != nil {
		return m.IsMultipathAsPathIgnoreOnwardsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsEnforceFirstAsEnabled() bool {
	if m != nil {
		return m.IsEnforceFirstAsEnabled
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetDefaultLocalPreference() uint32 {
	if m != nil {
		return m.DefaultLocalPreference
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetKeepAliveTime() uint32 {
	if m != nil {
		return m.KeepAliveTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetMinAcceptableHoldTime() uint32 {
	if m != nil {
		return m.MinAcceptableHoldTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetIsNeighborLogging() bool {
	if m != nil {
		return m.IsNeighborLogging
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsDefaultMetricConfigured() bool {
	if m != nil {
		return m.IsDefaultMetricConfigured
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetDefaultMetric() uint32 {
	if m != nil {
		return m.DefaultMetric
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetIsDefaultOriginateConfigured() bool {
	if m != nil {
		return m.IsDefaultOriginateConfigured
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsGracefulRestart() bool {
	if m != nil {
		return m.IsGracefulRestart
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetIsNsr() bool {
	if m != nil {
		return m.IsNsr
	}
	return false
}

func (m *BgpGlobalProcessInfoVrf_) GetRestartTime() uint32 {
	if m != nil {
		return m.RestartTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetStalePathTime() uint32 {
	if m != nil {
		return m.StalePathTime
	}
	return 0
}

func (m *BgpGlobalProcessInfoVrf_) GetRibPurgeTimeout() uint32 {
	if m != nil {
		return m.RibPurgeTimeout
	}
	return 0
}

type BgpGlobalProcessInfoBag struct {
	VrfName              string                    `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Vrfid                uint32                    `protobuf:"varint,51,opt,name=vrfid,proto3" json:"vrfid,omitempty"`
	Global               *BgpGlobalProcessInfoGbl_ `protobuf:"bytes,52,opt,name=global,proto3" json:"global,omitempty"`
	Vrf                  *BgpGlobalProcessInfoVrf_ `protobuf:"bytes,53,opt,name=vrf,proto3" json:"vrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BgpGlobalProcessInfoBag) Reset()         { *m = BgpGlobalProcessInfoBag{} }
func (m *BgpGlobalProcessInfoBag) String() string { return proto.CompactTextString(m) }
func (*BgpGlobalProcessInfoBag) ProtoMessage()    {}
func (*BgpGlobalProcessInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_99316a6dde1b0219, []int{4}
}

func (m *BgpGlobalProcessInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpGlobalProcessInfoBag.Unmarshal(m, b)
}
func (m *BgpGlobalProcessInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpGlobalProcessInfoBag.Marshal(b, m, deterministic)
}
func (m *BgpGlobalProcessInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpGlobalProcessInfoBag.Merge(m, src)
}
func (m *BgpGlobalProcessInfoBag) XXX_Size() int {
	return xxx_messageInfo_BgpGlobalProcessInfoBag.Size(m)
}
func (m *BgpGlobalProcessInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpGlobalProcessInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpGlobalProcessInfoBag proto.InternalMessageInfo

func (m *BgpGlobalProcessInfoBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpGlobalProcessInfoBag) GetVrfid() uint32 {
	if m != nil {
		return m.Vrfid
	}
	return 0
}

func (m *BgpGlobalProcessInfoBag) GetGlobal() *BgpGlobalProcessInfoGbl_ {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *BgpGlobalProcessInfoBag) GetVrf() *BgpGlobalProcessInfoVrf_ {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpGlobalProcessInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.global_process_info.bgp_global_process_info_bag_KEYS")
	proto.RegisterType((*ClusterIdBag_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.global_process_info.cluster_id_bag_")
	proto.RegisterType((*BgpGlobalProcessInfoGbl_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.global_process_info.bgp_global_process_info_gbl_")
	proto.RegisterType((*BgpGlobalProcessInfoVrf_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.global_process_info.bgp_global_process_info_vrf_")
	proto.RegisterType((*BgpGlobalProcessInfoBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.global_process_info.bgp_global_process_info_bag")
}

func init() { proto.RegisterFile("bgp_global_process_info_bag.proto", fileDescriptor_99316a6dde1b0219) }

var fileDescriptor_99316a6dde1b0219 = []byte{
	// 1514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x5f, 0x53, 0xdc, 0xba,
	0x15, 0x1f, 0x42, 0x02, 0x8b, 0x60, 0xf9, 0x63, 0x20, 0x38, 0x05, 0x26, 0xb0, 0x29, 0x29, 0x69,
	0x33, 0xdb, 0x94, 0x24, 0xd3, 0x3e, 0x64, 0x26, 0xb3, 0x81, 0xa5, 0xd9, 0x69, 0x96, 0x50, 0xc3,
	0xb4, 0xd3, 0xbe, 0x68, 0x64, 0x5b, 0x36, 0xa7, 0xb5, 0x25, 0x8f, 0xe4, 0xdd, 0x40, 0x66, 0xfa,
	0x19, 0xfa, 0xd4, 0x97, 0x7e, 0x84, 0xfb, 0x71, 0xee, 0xa7, 0xb9, 0x8f, 0x77, 0x74, 0x64, 0x79,
	0xbd, 0x81, 0xe4, 0x2d, 0x73, 0x5f, 0x76, 0xd7, 0xe7, 0xf7, 0xe7, 0x1c, 0x1d, 0x69, 0x25, 0x99,
	0xec, 0x87, 0x69, 0x41, 0xd3, 0x4c, 0x86, 0x2c, 0xa3, 0x85, 0x92, 0x11, 0xd7, 0x9a, 0x82, 0x48,
	0x24, 0x0d, 0x59, 0xda, 0x2d, 0x94, 0x2c, 0xa5, 0x47, 0x23, 0xd0, 0x91, 0xa4, 0x20, 0x35, 0xbd,
	0x56, 0x14, 0x8a, 0xf1, 0x2b, 0x6a, 0x44, 0xb2, 0xe0, 0xaa, 0x1b, 0xa6, 0x45, 0x17, 0x84, 0x2e,
	0x99, 0x88, 0xb8, 0xae, 0x7f, 0xd5, 0x3f, 0xa8, 0xf9, 0x8a, 0xc3, 0x9b, 0xee, 0x58, 0x25, 0xda,
	0x7c, 0x74, 0xef, 0x48, 0xd5, 0x09, 0xc9, 0xde, 0x37, 0xaa, 0xa0, 0x7f, 0xe9, 0xff, 0xe3, 0xc2,
	0x7b, 0x42, 0xda, 0xb5, 0xa9, 0x60, 0x39, 0xf7, 0x67, 0xf6, 0x66, 0x0e, 0x17, 0x82, 0x25, 0x17,
	0x3c, 0x63, 0x39, 0xf7, 0x1e, 0x91, 0xd6, 0x58, 0x25, 0x16, 0xbf, 0x87, 0xf8, 0xfc, 0x58, 0x25,
	0x06, 0xea, 0x50, 0xb2, 0x12, 0x65, 0x23, 0x5d, 0x72, 0x45, 0x21, 0x46, 0x5b, 0xef, 0xd7, 0x64,
	0xb9, 0x11, 0x1a, 0xb3, 0x0c, 0x3d, 0xdb, 0xc1, 0x52, 0x15, 0x1d, 0xc4, 0x7f, 0x63, 0x99, 0xf7,
	0x74, 0x4a, 0x58, 0xde, 0x14, 0xd6, 0xba, 0x1d, 0xb4, 0x6b, 0xda, 0xe5, 0x4d, 0xc1, 0x3b, 0x3f,
	0x2e, 0x90, 0x9d, 0xaf, 0x8d, 0x22, 0x0d, 0x33, 0xea, 0x3d, 0x27, 0x1e, 0x08, 0xdb, 0x10, 0x96,
	0x49, 0xc1, 0x69, 0x2e, 0x63, 0x3b, 0x8c, 0x56, 0xb0, 0x0a, 0xe2, 0xa2, 0x06, 0x86, 0x32, 0xc6,
	0xa1, 0x64, 0x32, 0x62, 0x19, 0x65, 0xba, 0xca, 0x37, 0x8f, 0xcf, 0x3d, 0x7d, 0xbb, 0x15, 0xb3,
	0x77, 0xb4, 0xe2, 0x09, 0x69, 0x2b, 0xae, 0x4b, 0xa6, 0x4a, 0x1a, 0xc9, 0x91, 0x28, 0xfd, 0xfb,
	0x76, 0x6c, 0x55, 0xf0, 0xd8, 0xc4, 0xbc, 0x7d, 0xb2, 0x34, 0x2a, 0x62, 0x56, 0x72, 0x1a, 0xf3,
	0x8c, 0xdd, 0xf8, 0x0f, 0x90, 0xb3, 0x68, 0x63, 0x27, 0x26, 0xe4, 0x75, 0xc9, 0x7a, 0xca, 0x05,
	0x57, 0x10, 0x51, 0x1d, 0x31, 0x41, 0x0b, 0xae, 0x40, 0xc6, 0xfe, 0x1c, 0x32, 0xd7, 0x2a, 0xe8,
	0x22, 0x62, 0xe2, 0x1c, 0x01, 0xef, 0x19, 0x59, 0x8d, 0xa4, 0x48, 0x78, 0xcc, 0x15, 0x2b, 0x41,
	0x0a, 0x0a, 0xb1, 0x3f, 0x8f, 0xe4, 0x95, 0xa9, 0xf8, 0x20, 0xf6, 0x76, 0x09, 0x99, 0x74, 0xd6,
	0x6f, 0x21, 0x69, 0xa1, 0x6e, 0xaa, 0x77, 0x44, 0x36, 0x8d, 0x02, 0xd2, 0x91, 0xe2, 0x31, 0x6d,
	0x30, 0x17, 0x90, 0xb9, 0x3e, 0x01, 0x8f, 0x6b, 0xcd, 0x5b, 0xb2, 0x0b, 0xba, 0xc1, 0xa5, 0xba,
	0xe0, 0x11, 0x24, 0xc0, 0x63, 0xca, 0x34, 0x85, 0xc2, 0x27, 0xd8, 0x6e, 0x1f, 0x74, 0xad, 0xb9,
	0x70, 0x8c, 0x9e, 0x1e, 0x14, 0xde, 0xff, 0x67, 0xc8, 0x6a, 0x43, 0xce, 0x45, 0xa9, 0x6e, 0xfc,
	0xc5, 0xbd, 0xd9, 0xc3, 0xc5, 0xa3, 0xa2, 0xfb, 0x9d, 0xff, 0x07, 0xdd, 0x2f, 0x16, 0x68, 0xb0,
	0x5c, 0x37, 0xa3, 0x6f, 0xea, 0x30, 0x0d, 0x63, 0x5a, 0xd0, 0x44, 0xaa, 0x9c, 0x95, 0xfe, 0x92,
	0x6d, 0x18, 0xd3, 0xe2, 0x14, 0x03, 0xa6, 0x61, 0x5a, 0xa5, 0x21, 0x75, 0xb3, 0xee, 0xda, 0xe3,
	0xb7, 0x6d, 0xc3, 0x0c, 0x78, 0x61, 0x27, 0xdf, 0x41, 0x66, 0x7a, 0x51, 0xc3, 0x45, 0xdc, 0x54,
	0x2c, 0xdb, 0xe9, 0x35, 0x50, 0x5f, 0xc4, 0x0d, 0xfe, 0x2e, 0x21, 0x93, 0x1c, 0xfe, 0x8a, 0x2d,
	0xa1, 0x36, 0x36, 0xab, 0xd6, 0xd9, 0xf9, 0xab, 0x76, 0xd5, 0x56, 0x1e, 0xde, 0x1f, 0xc8, 0x46,
	0xaa, 0x58, 0xc4, 0x93, 0x51, 0x46, 0x73, 0x06, 0xa2, 0xe4, 0xc2, 0x74, 0xc6, 0x5f, 0xc3, 0x19,
	0x59, 0x77, 0xd8, 0x70, 0x02, 0x79, 0xaf, 0xc9, 0xd6, 0xb4, 0x84, 0xb2, 0x2c, 0xa3, 0x22, 0x54,
	0xda, 0xf7, 0x50, 0xb5, 0x31, 0xa5, 0xea, 0x65, 0xd9, 0x59, 0xa8, 0xb4, 0xf7, 0x96, 0xec, 0x7c,
	0x21, 0x53, 0xbc, 0x64, 0x20, 0xa8, 0x92, 0xa3, 0x92, 0x6b, 0x7f, 0x1d, 0xb5, 0x8f, 0xa6, 0xb4,
	0x01, 0x32, 0x02, 0x24, 0x78, 0x07, 0x64, 0xd9, 0xcd, 0x8b, 0xca, 0x20, 0x87, 0xd2, 0xdf, 0xd8,
	0x9b, 0x39, 0xbc, 0x1f, 0xb4, 0xab, 0x68, 0x80, 0x41, 0x53, 0x5e, 0x98, 0x17, 0x34, 0x67, 0xd7,
	0x90, 0x8f, 0x72, 0x1a, 0x8e, 0x92, 0x84, 0x2b, 0xaa, 0xe1, 0x33, 0xf7, 0x37, 0x91, 0xbf, 0x11,
	0xe6, 0xc5, 0xd0, 0xa2, 0xef, 0x10, 0xbc, 0x80, 0xcf, 0xdc, 0xc9, 0x62, 0x9e, 0xb0, 0x51, 0x56,
	0x4e, 0xc9, 0x1e, 0xd6, 0xb2, 0x13, 0x8b, 0xde, 0x96, 0x45, 0x23, 0xa5, 0xb8, 0x98, 0x96, 0x6d,
	0xd5, 0xb2, 0x63, 0x8b, 0x36, 0x64, 0x6f, 0xc8, 0x76, 0x25, 0xbb, 0xb3, 0x50, 0x1f, 0xa5, 0x5b,
	0x56, 0x7a, 0xab, 0xd6, 0xce, 0x0f, 0xed, 0xaf, 0x6f, 0x6a, 0x66, 0xa7, 0xf5, 0x3a, 0xa4, 0x6d,
	0xbe, 0x41, 0x53, 0x16, 0x95, 0x30, 0x76, 0xfb, 0xd9, 0xe2, 0x58, 0x25, 0x03, 0xdd, 0xc3, 0x90,
	0xf7, 0x7b, 0xb2, 0x8e, 0x9d, 0xa7, 0x31, 0xe8, 0x12, 0x44, 0x3a, 0x02, 0x7d, 0xc5, 0x95, 0x7f,
	0x6f, 0x6f, 0xf6, 0xb0, 0x1d, 0x78, 0x08, 0x9d, 0x34, 0x11, 0x6f, 0x9b, 0x2c, 0x60, 0x14, 0xff,
	0xed, 0x76, 0x73, 0x6b, 0xd9, 0xc0, 0x20, 0xf6, 0x5e, 0x90, 0x8d, 0xc6, 0xb6, 0x30, 0xe1, 0xdd,
	0x47, 0x9e, 0x37, 0xc1, 0x02, 0xa7, 0x38, 0x23, 0x07, 0xa0, 0xa9, 0xe2, 0x26, 0xbd, 0x82, 0xd0,
	0x54, 0x02, 0x66, 0x50, 0xa5, 0xa4, 0x90, 0xd2, 0x42, 0x53, 0x2e, 0x58, 0x98, 0xf1, 0x18, 0xb7,
	0xbf, 0x56, 0xf0, 0x18, 0x74, 0xd0, 0xe0, 0x0e, 0xc2, 0xb4, 0xb8, 0x94, 0x83, 0xf4, 0x5c, 0xf7,
	0x2d, 0xcd, 0x7b, 0x4f, 0xf6, 0x41, 0xd3, 0x84, 0xe9, 0x92, 0xf2, 0xeb, 0x92, 0x2b, 0xc1, 0x32,
	0x9a, 0xb0, 0x2c, 0x93, 0x63, 0xae, 0x6a, 0xaf, 0x39, 0xf4, 0xda, 0x05, 0x7d, 0xca, 0x74, 0xd9,
	0xaf, 0x68, 0xa7, 0x15, 0xcb, 0x39, 0x5d, 0x92, 0x43, 0xd0, 0x34, 0xe4, 0xba, 0x2c, 0x58, 0x79,
	0x45, 0x73, 0xd0, 0x1a, 0x44, 0x4a, 0x73, 0x1e, 0x9b, 0x8e, 0x7e, 0x92, 0xca, 0xa4, 0xa8, 0x0c,
	0xe7, 0xd1, 0xb0, 0x03, 0xfa, 0x5d, 0x45, 0x1f, 0x5a, 0xf6, 0x90, 0xc7, 0x03, 0xfd, 0x77, 0x43,
	0x75, 0xae, 0x7f, 0x25, 0x4f, 0x9b, 0xae, 0x2c, 0xfb, 0xc4, 0x6e, 0x34, 0x8d, 0x64, 0x5e, 0x30,
	0xc5, 0xd1, 0xdc, 0x79, 0xb6, 0xd0, 0x73, 0x7f, 0xe2, 0xd9, 0x43, 0xee, 0xb1, 0xa5, 0x0e, 0x79,
	0xec, 0x2c, 0x07, 0xa4, 0xd3, 0xb4, 0x84, 0x54, 0x48, 0xc5, 0xcd, 0x96, 0x8a, 0x8f, 0xce, 0x6e,
	0xc1, 0x8d, 0xd9, 0xd9, 0x0d, 0x90, 0xd7, 0xd3, 0xe7, 0xac, 0xbc, 0x72, 0x56, 0x01, 0xf9, 0xcd,
	0x54, 0x75, 0x95, 0x47, 0x8e, 0x9f, 0x8a, 0x67, 0xec, 0xba, 0xf6, 0x23, 0xb7, 0xca, 0x43, 0xa7,
	0xa1, 0xf9, 0x19, 0x18, 0xa6, 0xf3, 0x8c, 0xc8, 0x8b, 0xa6, 0x67, 0x73, 0xa8, 0x89, 0x92, 0x39,
	0xb5, 0x27, 0x0f, 0x2d, 0x78, 0x63, 0x82, 0x16, 0xd1, 0xfc, 0xd9, 0xc4, 0x7c, 0x32, 0xea, 0x53,
	0x25, 0xf3, 0x63, 0x94, 0x9c, 0xf3, 0xc9, 0x64, 0xfd, 0x8b, 0xbc, 0xba, 0x2b, 0x49, 0xbd, 0x02,
	0xcd, 0xbe, 0x4c, 0xb9, 0x59, 0x58, 0x26, 0xcf, 0x64, 0x55, 0x2d, 0x61, 0xa2, 0xe7, 0xb7, 0x12,
	0xb9, 0xf5, 0x79, 0x2a, 0x55, 0x3f, 0x4c, 0x0b, 0x93, 0xaa, 0x5e, 0x62, 0x27, 0xe4, 0xf1, 0x54,
	0x93, 0x20, 0x2d, 0x5c, 0xd3, 0x9d, 0x6d, 0x1b, 0x6d, 0xb7, 0x1b, 0xcd, 0x81, 0xb4, 0xb0, 0x0d,
	0x77, 0x2e, 0xff, 0x24, 0xbf, 0x03, 0x4d, 0xf3, 0x51, 0x56, 0xc2, 0x54, 0xaf, 0x2b, 0x27, 0x29,
	0x3e, 0x31, 0x15, 0x4f, 0x0a, 0x5d, 0x46, 0xc7, 0x03, 0xd0, 0x43, 0xa7, 0xb0, 0xfd, 0xb6, 0xa6,
	0x1f, 0x2d, 0xdb, 0x79, 0xbf, 0x21, 0xdb, 0x60, 0xa4, 0x89, 0x54, 0x11, 0xa7, 0x09, 0x98, 0xa5,
	0xca, 0x26, 0x5e, 0x2b, 0xe8, 0xb5, 0x05, 0xba, 0x6f, 0x19, 0xa7, 0x86, 0xd0, 0xab, 0xd5, 0x7f,
	0x22, 0xbe, 0xdb, 0xff, 0xec, 0x2d, 0xa7, 0x50, 0x3c, 0xe1, 0x8a, 0x9b, 0x03, 0xc1, 0x9e, 0x1b,
	0x0f, 0x2b, 0xfc, 0x83, 0x81, 0xcf, 0x6b, 0xd4, 0x5c, 0xc7, 0xfe, 0xcd, 0x79, 0x41, 0x59, 0x06,
	0x63, 0x4e, 0x4b, 0xc8, 0xed, 0x09, 0xd2, 0x0e, 0xda, 0x26, 0xdc, 0x33, 0xd1, 0x4b, 0xc8, 0xb9,
	0xd9, 0x43, 0xae, 0x64, 0x16, 0x5b, 0x86, 0x87, 0x8c, 0x96, 0x09, 0x20, 0xf8, 0x47, 0xe2, 0xe7,
	0x20, 0x28, 0x8b, 0x22, 0x5e, 0x94, 0xa6, 0x22, 0x3a, 0xe1, 0xae, 0x23, 0x77, 0x33, 0x07, 0xd1,
	0xab, 0xe1, 0xf7, 0x4e, 0xd8, 0x25, 0xeb, 0xa0, 0xa9, 0xe0, 0x90, 0x5e, 0x85, 0x52, 0xd1, 0x4c,
	0xa6, 0x29, 0x88, 0x14, 0x8f, 0x87, 0x56, 0xb0, 0x06, 0xfa, 0xac, 0x42, 0x3e, 0x58, 0xc0, 0x1c,
	0x45, 0xa0, 0xeb, 0xad, 0x3e, 0xe7, 0xa5, 0xb9, 0x47, 0x35, 0xce, 0xd9, 0x4d, 0x7b, 0x14, 0x81,
	0xae, 0xf6, 0xfb, 0x21, 0x32, 0x1a, 0xe7, 0xed, 0x01, 0x59, 0x9e, 0x56, 0xe3, 0x19, 0xd1, 0x0e,
	0xda, 0x71, 0x53, 0xe0, 0xf5, 0x71, 0xbd, 0x38, 0xa6, 0x54, 0x90, 0x82, 0x30, 0xd7, 0xba, 0x46,
	0xaa, 0x2d, 0x4c, 0xb5, 0x53, 0xa7, 0xfa, 0xe8, 0x48, 0xd3, 0xb7, 0x01, 0xd0, 0xb4, 0x3e, 0x3c,
	0xab, 0xbb, 0x22, 0x1e, 0x12, 0x38, 0xbc, 0x3f, 0x57, 0x48, 0x60, 0x01, 0x6f, 0x93, 0xcc, 0x99,
	0x76, 0x68, 0xe5, 0x3f, 0x42, 0xca, 0x03, 0xd0, 0x67, 0x5a, 0x99, 0x6b, 0xa5, 0xbb, 0x7b, 0x62,
	0x4b, 0x7f, 0x65, 0xaf, 0x95, 0x55, 0x0c, 0x1b, 0xf9, 0x94, 0xac, 0xe8, 0x92, 0x65, 0xdc, 0x2e,
	0x48, 0x64, 0x6d, 0xdb, 0x81, 0x61, 0xd8, 0x2c, 0x3b, 0xe4, 0xfd, 0x96, 0xac, 0x29, 0x08, 0x69,
	0x31, 0x52, 0xa9, 0x9d, 0x6d, 0x39, 0x2a, 0xfd, 0x1d, 0x7b, 0x9f, 0x54, 0x10, 0x9e, 0x9b, 0xf8,
	0xa5, 0x0d, 0x77, 0x7e, 0xba, 0x47, 0xb6, 0xbf, 0xf1, 0x1e, 0x31, 0xf5, 0x76, 0x70, 0x34, 0xf5,
	0x76, 0xe0, 0x6d, 0x90, 0x07, 0x63, 0x95, 0x40, 0xec, 0xbf, 0x44, 0x6b, 0xfb, 0xe0, 0xfd, 0x6f,
	0x86, 0xcc, 0x59, 0x33, 0xff, 0xd5, 0xde, 0xcc, 0xe1, 0xe2, 0xd1, 0x7f, 0xbe, 0xfb, 0x15, 0xf0,
	0x5b, 0x6f, 0x10, 0x41, 0x55, 0x8c, 0xf7, 0xdf, 0x19, 0x32, 0x3b, 0x56, 0x89, 0xff, 0xfa, 0x17,
	0x2e, 0xca, 0x74, 0x33, 0x30, 0x95, 0x84, 0x73, 0xf8, 0xa6, 0xf8, 0xf2, 0xe7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x35, 0xee, 0xce, 0x4e, 0x0e, 0x00, 0x00,
}
