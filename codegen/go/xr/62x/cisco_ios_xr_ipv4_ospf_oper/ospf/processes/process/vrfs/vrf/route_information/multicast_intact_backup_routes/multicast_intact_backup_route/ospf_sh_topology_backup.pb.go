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
// source: ospf_sh_topology_backup.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_multicast_intact_backup_routes_multicast_intact_backup_route

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

type OspfShTopologyBackup_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTopologyBackup_KEYS) Reset()         { *m = OspfShTopologyBackup_KEYS{} }
func (m *OspfShTopologyBackup_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShTopologyBackup_KEYS) ProtoMessage()    {}
func (*OspfShTopologyBackup_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{0}
}

func (m *OspfShTopologyBackup_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopologyBackup_KEYS.Unmarshal(m, b)
}
func (m *OspfShTopologyBackup_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopologyBackup_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShTopologyBackup_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopologyBackup_KEYS.Merge(m, src)
}
func (m *OspfShTopologyBackup_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShTopologyBackup_KEYS.Size(m)
}
func (m *OspfShTopologyBackup_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopologyBackup_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopologyBackup_KEYS proto.InternalMessageInfo

func (m *OspfShTopologyBackup_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopologyBackup_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShTopologyBackup_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopologyBackup_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTime struct {
	Second               uint32   `protobuf:"varint,1,opt,name=second,proto3" json:"second,omitempty"`
	Nanosecond           uint32   `protobuf:"varint,2,opt,name=nanosecond,proto3" json:"nanosecond,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTime) Reset()         { *m = OspfShTime{} }
func (m *OspfShTime) String() string { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()    {}
func (*OspfShTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{1}
}

func (m *OspfShTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTime.Unmarshal(m, b)
}
func (m *OspfShTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTime.Marshal(b, m, deterministic)
}
func (m *OspfShTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTime.Merge(m, src)
}
func (m *OspfShTime) XXX_Size() int {
	return xxx_messageInfo_OspfShTime.Size(m)
}
func (m *OspfShTime) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTime.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTime proto.InternalMessageInfo

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

type OspfShTopCommon struct {
	RouteAreaId               uint32      `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId,proto3" json:"route_area_id,omitempty"`
	RouteTeMetric             uint32      `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric,proto3" json:"route_te_metric,omitempty"`
	RouteRibVersion           uint32      `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion,proto3" json:"route_rib_version,omitempty"`
	RouteSpfVersion           uint64      `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion,proto3" json:"route_spf_version,omitempty"`
	RouteForwardDistance      uint32      `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance,proto3" json:"route_forward_distance,omitempty"`
	RouteSource               uint32      `protobuf:"varint,6,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	RouteUpdateTime           *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime,proto3" json:"route_update_time,omitempty"`
	RouteFailTime             *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime,proto3" json:"route_fail_time,omitempty"`
	RouteSpfPriority          uint32      `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority,proto3" json:"route_spf_priority,omitempty"`
	RouteAutoExcluded         bool        `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded,proto3" json:"route_auto_excluded,omitempty"`
	RouteSrtePrefixRegistered bool        `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered,proto3" json:"route_srte_prefix_registered,omitempty"`
	RouteSrteNbrRegistered    uint32      `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered,proto3" json:"route_srte_nbr_registered,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}    `json:"-"`
	XXX_unrecognized          []byte      `json:"-"`
	XXX_sizecache             int32       `json:"-"`
}

func (m *OspfShTopCommon) Reset()         { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()    {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{2}
}

func (m *OspfShTopCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopCommon.Unmarshal(m, b)
}
func (m *OspfShTopCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopCommon.Marshal(b, m, deterministic)
}
func (m *OspfShTopCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopCommon.Merge(m, src)
}
func (m *OspfShTopCommon) XXX_Size() int {
	return xxx_messageInfo_OspfShTopCommon.Size(m)
}
func (m *OspfShTopCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopCommon.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopCommon proto.InternalMessageInfo

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

type OspfShRepEl struct {
	RepairElementId      string   `protobuf:"bytes,1,opt,name=repair_element_id,json=repairElementId,proto3" json:"repair_element_id,omitempty"`
	RepairLabel          uint32   `protobuf:"varint,2,opt,name=repair_label,json=repairLabel,proto3" json:"repair_label,omitempty"`
	RepairElementType    uint32   `protobuf:"varint,3,opt,name=repair_element_type,json=repairElementType,proto3" json:"repair_element_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRepEl) Reset()         { *m = OspfShRepEl{} }
func (m *OspfShRepEl) String() string { return proto.CompactTextString(m) }
func (*OspfShRepEl) ProtoMessage()    {}
func (*OspfShRepEl) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{3}
}

func (m *OspfShRepEl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRepEl.Unmarshal(m, b)
}
func (m *OspfShRepEl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRepEl.Marshal(b, m, deterministic)
}
func (m *OspfShRepEl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRepEl.Merge(m, src)
}
func (m *OspfShRepEl) XXX_Size() int {
	return xxx_messageInfo_OspfShRepEl.Size(m)
}
func (m *OspfShRepEl) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRepEl.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRepEl proto.InternalMessageInfo

func (m *OspfShRepEl) GetRepairElementId() string {
	if m != nil {
		return m.RepairElementId
	}
	return ""
}

func (m *OspfShRepEl) GetRepairLabel() uint32 {
	if m != nil {
		return m.RepairLabel
	}
	return 0
}

func (m *OspfShRepEl) GetRepairElementType() uint32 {
	if m != nil {
		return m.RepairElementType
	}
	return 0
}

type OspfShBackupPath struct {
	BackupRouteInterfaceName  string         `protobuf:"bytes,1,opt,name=backup_route_interface_name,json=backupRouteInterfaceName,proto3" json:"backup_route_interface_name,omitempty"`
	BackupRouteNextHopAddress string         `protobuf:"bytes,2,opt,name=backup_route_next_hop_address,json=backupRouteNextHopAddress,proto3" json:"backup_route_next_hop_address,omitempty"`
	BackupRouteSource         string         `protobuf:"bytes,3,opt,name=backup_route_source,json=backupRouteSource,proto3" json:"backup_route_source,omitempty"`
	BackupMetric              uint32         `protobuf:"varint,4,opt,name=backup_metric,json=backupMetric,proto3" json:"backup_metric,omitempty"`
	PrimaryPath               bool           `protobuf:"varint,5,opt,name=primary_path,json=primaryPath,proto3" json:"primary_path,omitempty"`
	LineCardDisjoint          bool           `protobuf:"varint,6,opt,name=line_card_disjoint,json=lineCardDisjoint,proto3" json:"line_card_disjoint,omitempty"`
	Downstream                bool           `protobuf:"varint,7,opt,name=downstream,proto3" json:"downstream,omitempty"`
	NodeProtect               bool           `protobuf:"varint,8,opt,name=node_protect,json=nodeProtect,proto3" json:"node_protect,omitempty"`
	SrlgDisjoint              bool           `protobuf:"varint,9,opt,name=srlg_disjoint,json=srlgDisjoint,proto3" json:"srlg_disjoint,omitempty"`
	BackupRemoteLfa           string         `protobuf:"bytes,10,opt,name=backup_remote_lfa,json=backupRemoteLfa,proto3" json:"backup_remote_lfa,omitempty"`
	BackupRepair              []*OspfShRepEl `protobuf:"bytes,11,rep,name=backup_repair,json=backupRepair,proto3" json:"backup_repair,omitempty"`
	BackupRepairListSize      uint32         `protobuf:"varint,12,opt,name=backup_repair_list_size,json=backupRepairListSize,proto3" json:"backup_repair_list_size,omitempty"`
	BackupTunnelInterfaceName string         `protobuf:"bytes,13,opt,name=backup_tunnel_interface_name,json=backupTunnelInterfaceName,proto3" json:"backup_tunnel_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}       `json:"-"`
	XXX_unrecognized          []byte         `json:"-"`
	XXX_sizecache             int32          `json:"-"`
}

func (m *OspfShBackupPath) Reset()         { *m = OspfShBackupPath{} }
func (m *OspfShBackupPath) String() string { return proto.CompactTextString(m) }
func (*OspfShBackupPath) ProtoMessage()    {}
func (*OspfShBackupPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{4}
}

func (m *OspfShBackupPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShBackupPath.Unmarshal(m, b)
}
func (m *OspfShBackupPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShBackupPath.Marshal(b, m, deterministic)
}
func (m *OspfShBackupPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShBackupPath.Merge(m, src)
}
func (m *OspfShBackupPath) XXX_Size() int {
	return xxx_messageInfo_OspfShBackupPath.Size(m)
}
func (m *OspfShBackupPath) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShBackupPath.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShBackupPath proto.InternalMessageInfo

func (m *OspfShBackupPath) GetBackupRouteInterfaceName() string {
	if m != nil {
		return m.BackupRouteInterfaceName
	}
	return ""
}

func (m *OspfShBackupPath) GetBackupRouteNextHopAddress() string {
	if m != nil {
		return m.BackupRouteNextHopAddress
	}
	return ""
}

func (m *OspfShBackupPath) GetBackupRouteSource() string {
	if m != nil {
		return m.BackupRouteSource
	}
	return ""
}

func (m *OspfShBackupPath) GetBackupMetric() uint32 {
	if m != nil {
		return m.BackupMetric
	}
	return 0
}

func (m *OspfShBackupPath) GetPrimaryPath() bool {
	if m != nil {
		return m.PrimaryPath
	}
	return false
}

func (m *OspfShBackupPath) GetLineCardDisjoint() bool {
	if m != nil {
		return m.LineCardDisjoint
	}
	return false
}

func (m *OspfShBackupPath) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *OspfShBackupPath) GetNodeProtect() bool {
	if m != nil {
		return m.NodeProtect
	}
	return false
}

func (m *OspfShBackupPath) GetSrlgDisjoint() bool {
	if m != nil {
		return m.SrlgDisjoint
	}
	return false
}

func (m *OspfShBackupPath) GetBackupRemoteLfa() string {
	if m != nil {
		return m.BackupRemoteLfa
	}
	return ""
}

func (m *OspfShBackupPath) GetBackupRepair() []*OspfShRepEl {
	if m != nil {
		return m.BackupRepair
	}
	return nil
}

func (m *OspfShBackupPath) GetBackupRepairListSize() uint32 {
	if m != nil {
		return m.BackupRepairListSize
	}
	return 0
}

func (m *OspfShBackupPath) GetBackupTunnelInterfaceName() string {
	if m != nil {
		return m.BackupTunnelInterfaceName
	}
	return ""
}

type OspfShSrUloopPath struct {
	MicroloopRepair              []*OspfShRepEl `protobuf:"bytes,1,rep,name=microloop_repair,json=microloopRepair,proto3" json:"microloop_repair,omitempty"`
	MicroloopRepairListSize      uint32         `protobuf:"varint,2,opt,name=microloop_repair_list_size,json=microloopRepairListSize,proto3" json:"microloop_repair_list_size,omitempty"`
	MicroloopTunnelInterfaceName string         `protobuf:"bytes,3,opt,name=microloop_tunnel_interface_name,json=microloopTunnelInterfaceName,proto3" json:"microloop_tunnel_interface_name,omitempty"`
	MicroloopStrictSpf           bool           `protobuf:"varint,4,opt,name=microloop_strict_spf,json=microloopStrictSpf,proto3" json:"microloop_strict_spf,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	XXX_unrecognized             []byte         `json:"-"`
	XXX_sizecache                int32          `json:"-"`
}

func (m *OspfShSrUloopPath) Reset()         { *m = OspfShSrUloopPath{} }
func (m *OspfShSrUloopPath) String() string { return proto.CompactTextString(m) }
func (*OspfShSrUloopPath) ProtoMessage()    {}
func (*OspfShSrUloopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{5}
}

func (m *OspfShSrUloopPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShSrUloopPath.Unmarshal(m, b)
}
func (m *OspfShSrUloopPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShSrUloopPath.Marshal(b, m, deterministic)
}
func (m *OspfShSrUloopPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShSrUloopPath.Merge(m, src)
}
func (m *OspfShSrUloopPath) XXX_Size() int {
	return xxx_messageInfo_OspfShSrUloopPath.Size(m)
}
func (m *OspfShSrUloopPath) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShSrUloopPath.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShSrUloopPath proto.InternalMessageInfo

func (m *OspfShSrUloopPath) GetMicroloopRepair() []*OspfShRepEl {
	if m != nil {
		return m.MicroloopRepair
	}
	return nil
}

func (m *OspfShSrUloopPath) GetMicroloopRepairListSize() uint32 {
	if m != nil {
		return m.MicroloopRepairListSize
	}
	return 0
}

func (m *OspfShSrUloopPath) GetMicroloopTunnelInterfaceName() string {
	if m != nil {
		return m.MicroloopTunnelInterfaceName
	}
	return ""
}

func (m *OspfShSrUloopPath) GetMicroloopStrictSpf() bool {
	if m != nil {
		return m.MicroloopStrictSpf
	}
	return false
}

type OspfShTopPathBackup struct {
	RouteInterfaceName       string             `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName,proto3" json:"route_interface_name,omitempty"`
	RouteNextHopAddress      string             `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress,proto3" json:"route_next_hop_address,omitempty"`
	RouteSource              string             `protobuf:"bytes,3,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	RouteLsaid               string             `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid,proto3" json:"route_lsaid,omitempty"`
	RoutePathIsMcastIntact   bool               `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact,proto3" json:"route_path_is_mcast_intact,omitempty"`
	RoutePathIsUcmpPath      bool               `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath,proto3" json:"route_path_is_ucmp_path,omitempty"`
	RouteMetric              uint32             `protobuf:"varint,7,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	RoutePathId              uint32             `protobuf:"varint,8,opt,name=route_path_id,json=routePathId,proto3" json:"route_path_id,omitempty"`
	LsaType                  uint32             `protobuf:"varint,9,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	RouteBackupPath          *OspfShBackupPath  `protobuf:"bytes,10,opt,name=route_backup_path,json=routeBackupPath,proto3" json:"route_backup_path,omitempty"`
	SrMicroloopAvoidancePath *OspfShSrUloopPath `protobuf:"bytes,11,opt,name=sr_microloop_avoidance_path,json=srMicroloopAvoidancePath,proto3" json:"sr_microloop_avoidance_path,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *OspfShTopPathBackup) Reset()         { *m = OspfShTopPathBackup{} }
func (m *OspfShTopPathBackup) String() string { return proto.CompactTextString(m) }
func (*OspfShTopPathBackup) ProtoMessage()    {}
func (*OspfShTopPathBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{6}
}

func (m *OspfShTopPathBackup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopPathBackup.Unmarshal(m, b)
}
func (m *OspfShTopPathBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopPathBackup.Marshal(b, m, deterministic)
}
func (m *OspfShTopPathBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopPathBackup.Merge(m, src)
}
func (m *OspfShTopPathBackup) XXX_Size() int {
	return xxx_messageInfo_OspfShTopPathBackup.Size(m)
}
func (m *OspfShTopPathBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopPathBackup.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopPathBackup proto.InternalMessageInfo

func (m *OspfShTopPathBackup) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPathBackup) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPathBackup) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPathBackup) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPathBackup) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPathBackup) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPathBackup) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPathBackup) GetRoutePathId() uint32 {
	if m != nil {
		return m.RoutePathId
	}
	return 0
}

func (m *OspfShTopPathBackup) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPathBackup) GetRouteBackupPath() *OspfShBackupPath {
	if m != nil {
		return m.RouteBackupPath
	}
	return nil
}

func (m *OspfShTopPathBackup) GetSrMicroloopAvoidancePath() *OspfShSrUloopPath {
	if m != nil {
		return m.SrMicroloopAvoidancePath
	}
	return nil
}

type OspfShTopologyBackup struct {
	RoutePrefix          string                 `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
	RoutePrefixLength    uint32                 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength,proto3" json:"route_prefix_length,omitempty"`
	RouteMetric          uint32                 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	RouteType            string                 `protobuf:"bytes,53,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	RouteConnected       bool                   `protobuf:"varint,54,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	RouteInfo            *OspfShTopCommon       `protobuf:"bytes,55,opt,name=route_info,json=routeInfo,proto3" json:"route_info,omitempty"`
	RoutePath            []*OspfShTopPathBackup `protobuf:"bytes,56,rep,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *OspfShTopologyBackup) Reset()         { *m = OspfShTopologyBackup{} }
func (m *OspfShTopologyBackup) String() string { return proto.CompactTextString(m) }
func (*OspfShTopologyBackup) ProtoMessage()    {}
func (*OspfShTopologyBackup) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ea12a528fbabda, []int{7}
}

func (m *OspfShTopologyBackup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopologyBackup.Unmarshal(m, b)
}
func (m *OspfShTopologyBackup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopologyBackup.Marshal(b, m, deterministic)
}
func (m *OspfShTopologyBackup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopologyBackup.Merge(m, src)
}
func (m *OspfShTopologyBackup) XXX_Size() int {
	return xxx_messageInfo_OspfShTopologyBackup.Size(m)
}
func (m *OspfShTopologyBackup) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopologyBackup.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopologyBackup proto.InternalMessageInfo

func (m *OspfShTopologyBackup) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopologyBackup) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopologyBackup) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopologyBackup) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopologyBackup) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopologyBackup) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopologyBackup) GetRoutePath() []*OspfShTopPathBackup {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShTopologyBackup_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_topology_backup_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShBackupPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_backup_path")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopPathBackup)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_top_path_backup")
	proto.RegisterType((*OspfShTopologyBackup)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.multicast_intact_backup_routes.multicast_intact_backup_route.ospf_sh_topology_backup")
}

func init() { proto.RegisterFile("ospf_sh_topology_backup.proto", fileDescriptor_f7ea12a528fbabda) }

var fileDescriptor_f7ea12a528fbabda = []byte{
	// 1310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcd, 0x6e, 0x1c, 0xc5,
	0x13, 0xd7, 0x24, 0xf9, 0xdb, 0xbb, 0xbd, 0xde, 0x38, 0xee, 0xe4, 0x1f, 0x8f, 0xf3, 0x41, 0xcc,
	0x46, 0x02, 0x2b, 0x8a, 0x56, 0xc8, 0x71, 0xf8, 0x14, 0x02, 0x13, 0x1c, 0x61, 0xe1, 0x44, 0xd6,
	0xac, 0x83, 0xc4, 0xa9, 0xd5, 0x3b, 0x53, 0x63, 0x37, 0xcc, 0x4c, 0x8f, 0xba, 0x7b, 0x37, 0x76,
	0x1e, 0x01, 0x0e, 0xdc, 0x10, 0x70, 0x00, 0x21, 0xde, 0x03, 0x81, 0x10, 0x5c, 0xe0, 0x19, 0x90,
	0x78, 0x0a, 0xae, 0xa8, 0xab, 0x7b, 0x66, 0x67, 0x9c, 0x8f, 0xf3, 0x5e, 0x56, 0x3b, 0x55, 0xd5,
	0xd5, 0xf5, 0xab, 0xae, 0xfa, 0x75, 0x35, 0xb9, 0x2e, 0x75, 0x99, 0x32, 0x7d, 0xc4, 0x8c, 0x2c,
	0x65, 0x26, 0x0f, 0x4f, 0xd8, 0x98, 0xc7, 0x9f, 0x4f, 0xca, 0x61, 0xa9, 0xa4, 0x91, 0xf4, 0x8b,
	0x20, 0x16, 0x3a, 0x96, 0x4c, 0x48, 0xcd, 0x8e, 0x15, 0x13, 0xe5, 0x74, 0x8b, 0xe1, 0x0a, 0x59,
	0x82, 0x1a, 0xda, 0x7f, 0xd6, 0x30, 0x06, 0xad, 0x41, 0x57, 0xff, 0x86, 0x53, 0x95, 0xe2, 0xcf,
	0x50, 0xc9, 0x89, 0x01, 0x26, 0x8a, 0x54, 0xaa, 0x9c, 0x1b, 0x21, 0x8b, 0x61, 0x3e, 0xc9, 0x8c,
	0x88, 0xb9, 0x36, 0x4c, 0x14, 0x86, 0xc7, 0xc6, 0xef, 0xc6, 0xd0, 0x52, 0xbf, 0x58, 0x3d, 0xf8,
	0x36, 0x20, 0xd7, 0x9e, 0x13, 0x2e, 0xfb, 0x78, 0xe7, 0xd3, 0x11, 0x7d, 0x99, 0x2c, 0xf9, 0x18,
	0x58, 0xc1, 0x73, 0x08, 0x83, 0xf5, 0x60, 0xa3, 0x1b, 0xf5, 0xbc, 0xec, 0x21, 0xcf, 0x81, 0xae,
	0x91, 0xce, 0x54, 0xa5, 0x4e, 0x7d, 0x06, 0xd5, 0x8b, 0x53, 0x95, 0xa2, 0xea, 0x32, 0x59, 0x28,
	0x15, 0xa4, 0xe2, 0x38, 0x3c, 0x8b, 0x0a, 0xff, 0x45, 0x6f, 0x92, 0xbe, 0xfb, 0xc7, 0x32, 0x28,
	0x0e, 0xcd, 0x51, 0x78, 0x6e, 0x3d, 0xd8, 0xe8, 0x47, 0x4b, 0x4e, 0xb8, 0x87, 0xb2, 0xc1, 0x7d,
	0xb2, 0x54, 0x87, 0x26, 0x9c, 0x33, 0x0d, 0xb1, 0x2c, 0x12, 0x0c, 0xa2, 0x1f, 0xf9, 0x2f, 0xfa,
	0x12, 0x21, 0x05, 0x2f, 0xa4, 0xd7, 0x9d, 0x41, 0x5d, 0x43, 0x32, 0xf8, 0x77, 0x81, 0xd0, 0x06,
	0x46, 0x16, 0xcb, 0x3c, 0x97, 0x05, 0x1d, 0x90, 0xbe, 0x4b, 0x26, 0x57, 0xc0, 0x99, 0xa8, 0xbc,
	0xf6, 0x50, 0xb8, 0xad, 0x80, 0xef, 0x26, 0xf4, 0x15, 0xb2, 0xec, 0x6c, 0x0c, 0xb0, 0x1c, 0x8c,
	0x12, 0xb1, 0xf7, 0xef, 0x96, 0x1e, 0xc0, 0x03, 0x14, 0xd2, 0x5b, 0x64, 0xc5, 0xd9, 0x29, 0x31,
	0x66, 0x53, 0x50, 0x5a, 0xc8, 0x02, 0x21, 0xf7, 0x23, 0xe7, 0x20, 0x12, 0xe3, 0x4f, 0x9c, 0x78,
	0x66, 0x6b, 0x43, 0xaa, 0x6c, 0x2d, 0xfe, 0x73, 0xde, 0x76, 0x54, 0xa6, 0x95, 0xed, 0x16, 0xb9,
	0xec, 0x6c, 0x53, 0xa9, 0x1e, 0x73, 0x95, 0xb0, 0x44, 0x68, 0xc3, 0x8b, 0x18, 0xc2, 0xff, 0xa1,
	0xf3, 0x4b, 0xa8, 0xbd, 0xef, 0x94, 0x1f, 0x7a, 0x9d, 0x3d, 0x33, 0xbf, 0x83, 0x9c, 0xa8, 0x18,
	0xc2, 0x85, 0x06, 0xb0, 0x11, 0x8a, 0xe8, 0xef, 0x41, 0x15, 0xc5, 0xa4, 0x4c, 0xb8, 0x05, 0x28,
	0x72, 0x08, 0x17, 0xd7, 0x83, 0x8d, 0xde, 0xe6, 0x37, 0xc1, 0x70, 0x7e, 0x2a, 0x74, 0xd8, 0x2c,
	0x01, 0x9f, 0xa0, 0x47, 0x18, 0xf2, 0x81, 0xad, 0x89, 0x5f, 0x83, 0xea, 0x84, 0x52, 0x2e, 0x32,
	0x87, 0xa2, 0x33, 0xef, 0x28, 0x5c, 0xf1, 0xdc, 0xe7, 0x22, 0x43, 0x0c, 0xb7, 0x09, 0x9d, 0x15,
	0x44, 0xa9, 0x84, 0x54, 0xc2, 0x9c, 0x84, 0x5d, 0x3c, 0xb4, 0x0b, 0x55, 0x45, 0xec, 0x7b, 0x39,
	0x1d, 0x92, 0x8b, 0xbe, 0x6c, 0x27, 0x46, 0x32, 0x38, 0x8e, 0xb3, 0x49, 0x02, 0x49, 0x48, 0xd6,
	0x83, 0x8d, 0x4e, 0xe4, 0xce, 0x74, 0x7b, 0x62, 0xe4, 0x8e, 0x57, 0xd0, 0xf7, 0xc8, 0x35, 0xef,
	0x5d, 0x19, 0x60, 0xbe, 0xeb, 0x14, 0x1c, 0x0a, 0x6d, 0x40, 0x41, 0x12, 0xf6, 0x70, 0xe1, 0x9a,
	0xdb, 0x47, 0x19, 0xd8, 0x47, 0x8b, 0xa8, 0x36, 0xa0, 0x6f, 0x91, 0xb5, 0x86, 0x83, 0x62, 0xac,
	0x9a, 0xab, 0x97, 0x30, 0xca, 0xcb, 0xf5, 0xea, 0x87, 0x63, 0x35, 0x5b, 0x3a, 0xf8, 0x2a, 0x20,
	0xe7, 0x2b, 0xe4, 0x0a, 0x4a, 0x06, 0x19, 0x56, 0x3f, 0x94, 0x5c, 0x28, 0x06, 0x19, 0xe4, 0x50,
	0x98, 0xaa, 0xf3, 0xba, 0xd1, 0xb2, 0x53, 0xec, 0x38, 0xf9, 0x6e, 0x82, 0x75, 0xec, 0x6c, 0x33,
	0x3e, 0x86, 0xcc, 0xb7, 0x5e, 0xcf, 0xc9, 0xf6, 0xac, 0x08, 0xb3, 0xd1, 0x76, 0x67, 0x4e, 0x4a,
	0xf0, 0xad, 0xb7, 0xd2, 0x72, 0x78, 0x70, 0x52, 0xc2, 0xe0, 0xcb, 0x05, 0x72, 0xb1, 0x8a, 0xc8,
	0x1f, 0x50, 0xc9, 0xcd, 0x11, 0x7d, 0x97, 0x5c, 0x6d, 0x9e, 0x97, 0x3d, 0x43, 0x50, 0x29, 0x8f,
	0xa1, 0xc9, 0x7a, 0xa1, 0x33, 0x89, 0xac, 0xc5, 0x6e, 0x65, 0x80, 0x3c, 0xf7, 0x3e, 0xb9, 0xde,
	0x5a, 0x5e, 0xc0, 0xb1, 0x61, 0x47, 0xb2, 0x64, 0x3c, 0x49, 0x14, 0x68, 0xed, 0x79, 0x71, 0xad,
	0xe1, 0xe0, 0x21, 0x1c, 0x9b, 0x8f, 0x64, 0xb9, 0xed, 0x0c, 0x2c, 0x90, 0x96, 0x07, 0xdf, 0xba,
	0x8e, 0x36, 0x57, 0x1a, 0xeb, 0x7c, 0x03, 0xdf, 0x24, 0x7d, 0x6f, 0xef, 0x79, 0xc9, 0x33, 0xa8,
	0x13, 0x7a, 0x5a, 0x42, 0xf2, 0x16, 0x39, 0x57, 0x27, 0x88, 0x12, 0x49, 0xa3, 0x63, 0xc9, 0x1b,
	0x65, 0xfb, 0x16, 0xf8, 0x6d, 0x42, 0x33, 0x51, 0x00, 0x8b, 0x3d, 0xbb, 0x7c, 0x26, 0x45, 0x61,
	0x90, 0x31, 0x3a, 0xd1, 0x05, 0xab, 0xb9, 0xe7, 0x98, 0x05, 0xe5, 0x96, 0x6a, 0x13, 0xf9, 0xb8,
	0xd0, 0x46, 0x01, 0xcf, 0x91, 0x2e, 0x3a, 0x51, 0x43, 0x62, 0x37, 0x2c, 0x64, 0x62, 0xcb, 0x4c,
	0x1a, 0x88, 0x0d, 0xb6, 0x62, 0x27, 0xea, 0x59, 0xd9, 0xbe, 0x13, 0xd9, 0xc0, 0xb5, 0xca, 0x0e,
	0x67, 0x7b, 0x75, 0xd1, 0x66, 0xc9, 0x0a, 0xeb, 0x7d, 0x6e, 0x91, 0x95, 0x2a, 0x1b, 0x90, 0x4b,
	0x03, 0x2c, 0x4b, 0x39, 0x96, 0x78, 0x37, 0x5a, 0xf6, 0xb9, 0x40, 0xf9, 0x5e, 0xca, 0xe9, 0x2f,
	0x41, 0x9d, 0x0a, 0x77, 0xde, 0x61, 0x6f, 0xfd, 0xec, 0x46, 0x6f, 0xf3, 0xbb, 0xb9, 0x24, 0x00,
	0xd7, 0x06, 0xd5, 0x39, 0x45, 0x18, 0x30, 0xbd, 0x4b, 0x56, 0x5b, 0x08, 0x58, 0x26, 0xb4, 0x61,
	0x5a, 0x3c, 0x01, 0xdf, 0x60, 0x97, 0x9a, 0xe6, 0x7b, 0x42, 0x9b, 0x91, 0x78, 0x02, 0xb6, 0xb5,
	0xfd, 0x32, 0x33, 0x29, 0x0a, 0xc8, 0x4e, 0x57, 0x6d, 0xbf, 0x59, 0x74, 0x07, 0x68, 0xd2, 0x2a,
	0xdb, 0xc1, 0xd7, 0x67, 0xc9, 0xff, 0xab, 0xc0, 0xb4, 0x62, 0x93, 0x4c, 0x4a, 0xdf, 0x0f, 0x7f,
	0x04, 0xe4, 0x42, 0x2e, 0x62, 0x25, 0x51, 0xe4, 0xf3, 0x1a, 0xcc, 0x7f, 0x5e, 0x97, 0xeb, 0xa0,
	0x7d, 0x6a, 0xdf, 0x21, 0x57, 0x4e, 0xe3, 0x68, 0x64, 0xd7, 0x31, 0xca, 0xea, 0xa9, 0x45, 0x75,
	0x82, 0x77, 0xc8, 0x8d, 0xd9, 0xe2, 0x67, 0xe7, 0xd8, 0x35, 0xe8, 0xb5, 0xda, 0xec, 0x19, 0x69,
	0xa6, 0xaf, 0x91, 0x4b, 0x33, 0x37, 0xda, 0x76, 0xa6, 0xb1, 0x5c, 0x8f, 0x2d, 0xdb, 0x89, 0x68,
	0xad, 0x1b, 0xa1, 0x6a, 0x54, 0xa6, 0x83, 0x3f, 0x17, 0xc8, 0x6a, 0x73, 0x64, 0xb1, 0x67, 0xe2,
	0x71, 0x5b, 0x6f, 0x2f, 0xe0, 0x28, 0x77, 0x95, 0xb4, 0xf7, 0xbf, 0x53, 0x4d, 0x11, 0xcf, 0xa1,
	0x25, 0x77, 0xa1, 0x9c, 0x22, 0xa4, 0xd3, 0x43, 0x84, 0x03, 0xda, 0x1a, 0x22, 0x6e, 0x10, 0xf7,
	0xc9, 0x32, 0xcd, 0x45, 0x82, 0x70, 0xba, 0x11, 0x41, 0xd1, 0x9e, 0x95, 0xd0, 0xb7, 0xc9, 0x15,
	0x67, 0x80, 0xf1, 0x0b, 0xcd, 0xf2, 0xc6, 0x39, 0x7a, 0x36, 0x72, 0xa1, 0x59, 0x2e, 0xda, 0xd5,
	0x0f, 0xac, 0x7a, 0x17, 0xb5, 0x74, 0x8b, 0xac, 0xb6, 0xd7, 0x4e, 0xe2, 0xdc, 0x25, 0xc2, 0xb3,
	0xd3, 0xc5, 0xc6, 0xc2, 0x47, 0x71, 0x5e, 0x22, 0x9d, 0xd5, 0x51, 0x7b, 0x56, 0x5c, 0x6c, 0x8c,
	0x3e, 0x9e, 0x14, 0xeb, 0xb9, 0xcf, 0x39, 0x4e, 0x90, 0xa4, 0x2a, 0x1b, 0x74, 0x97, 0xd8, 0x91,
	0x36, 0xd3, 0xdc, 0xdd, 0x25, 0xee, 0x22, 0x5e, 0xcc, 0x34, 0xb7, 0x37, 0x08, 0xfd, 0xab, 0x9e,
	0x9c, 0x1a, 0xf7, 0x07, 0x72, 0x53, 0x6f, 0xf3, 0x87, 0xb9, 0x6c, 0x8d, 0x46, 0x9c, 0x7e, 0x80,
	0xfa, 0x00, 0x25, 0x98, 0xb0, 0x7f, 0x02, 0x72, 0x55, 0x2b, 0x36, 0xab, 0x4f, 0x3e, 0x95, 0x22,
	0xb1, 0x63, 0xa4, 0x03, 0xd6, 0x43, 0x60, 0x3f, 0xce, 0x25, 0xb0, 0x16, 0x65, 0x45, 0xa1, 0x56,
	0x0f, 0x2a, 0x14, 0xdb, 0x15, 0x08, 0x8b, 0x71, 0xf0, 0xf7, 0xb9, 0x56, 0x37, 0x35, 0x1f, 0x39,
	0xb3, 0x82, 0xf1, 0xef, 0x94, 0xcd, 0x46, 0x99, 0xbb, 0x51, 0x68, 0x36, 0x71, 0xb5, 0x9f, 0x2c,
	0x77, 0xfc, 0x8c, 0x31, 0xb3, 0x74, 0xef, 0x96, 0xa7, 0x6a, 0x70, 0xeb, 0xe9, 0x1a, 0xbc, 0x4e,
	0x88, 0x7f, 0x57, 0xd8, 0x0a, 0xbb, 0x8b, 0x7b, 0x76, 0xdd, 0x93, 0xc2, 0xd6, 0xd8, 0xab, 0xd5,
	0x50, 0x1b, 0xcb, 0xa2, 0x80, 0xd8, 0x40, 0x12, 0xbe, 0x8e, 0x35, 0x7f, 0x1e, 0xc5, 0xf7, 0x2a,
	0x29, 0xfd, 0x39, 0xa8, 0x1c, 0xd9, 0xe4, 0x86, 0x6f, 0xe0, 0x61, 0x7d, 0x3f, 0x9f, 0x93, 0x6f,
	0xfd, 0xf2, 0xf2, 0x48, 0x77, 0x8b, 0x54, 0xd2, 0xdf, 0x6a, 0x00, 0x58, 0x6d, 0x6f, 0xe2, 0x0d,
	0xf3, 0xd3, 0xdc, 0x02, 0x68, 0xf0, 0xb0, 0x47, 0x61, 0x0b, 0x6c, 0xbc, 0x80, 0x2f, 0xfb, 0x3b,
	0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x98, 0xd7, 0xe2, 0x3b, 0xfa, 0x0f, 0x00, 0x00,
}
