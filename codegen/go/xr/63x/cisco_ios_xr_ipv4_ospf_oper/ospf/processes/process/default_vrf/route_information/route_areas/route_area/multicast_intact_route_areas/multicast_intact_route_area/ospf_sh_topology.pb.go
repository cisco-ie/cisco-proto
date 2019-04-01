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
// source: ospf_sh_topology.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_multicast_intact_route_areas_multicast_intact_route_area

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

type OspfShTopology_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShTopology_KEYS) Reset()         { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()    {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4169ff6e5af11206, []int{0}
}

func (m *OspfShTopology_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopology_KEYS.Unmarshal(m, b)
}
func (m *OspfShTopology_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopology_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShTopology_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopology_KEYS.Merge(m, src)
}
func (m *OspfShTopology_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShTopology_KEYS.Size(m)
}
func (m *OspfShTopology_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopology_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopology_KEYS proto.InternalMessageInfo

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
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
	return fileDescriptor_4169ff6e5af11206, []int{1}
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
	return fileDescriptor_4169ff6e5af11206, []int{2}
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
	return fileDescriptor_4169ff6e5af11206, []int{3}
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
	return fileDescriptor_4169ff6e5af11206, []int{4}
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

type OspfShTopPath struct {
	RouteInterfaceName       string             `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName,proto3" json:"route_interface_name,omitempty"`
	RouteNextHopAddress      string             `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress,proto3" json:"route_next_hop_address,omitempty"`
	RouteSource              string             `protobuf:"bytes,3,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	RouteLsaid               string             `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid,proto3" json:"route_lsaid,omitempty"`
	RoutePathIsMcastIntact   bool               `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact,proto3" json:"route_path_is_mcast_intact,omitempty"`
	RoutePathIsUcmpPath      bool               `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath,proto3" json:"route_path_is_ucmp_path,omitempty"`
	RouteMetric              uint32             `protobuf:"varint,7,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	LsaType                  uint32             `protobuf:"varint,8,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	AreaId                   uint32             `protobuf:"varint,9,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaFormat               bool               `protobuf:"varint,10,opt,name=area_format,json=areaFormat,proto3" json:"area_format,omitempty"`
	SrMicroloopAvoidancePath *OspfShSrUloopPath `protobuf:"bytes,11,opt,name=sr_microloop_avoidance_path,json=srMicroloopAvoidancePath,proto3" json:"sr_microloop_avoidance_path,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *OspfShTopPath) Reset()         { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()    {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_4169ff6e5af11206, []int{5}
}

func (m *OspfShTopPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopPath.Unmarshal(m, b)
}
func (m *OspfShTopPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopPath.Marshal(b, m, deterministic)
}
func (m *OspfShTopPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopPath.Merge(m, src)
}
func (m *OspfShTopPath) XXX_Size() int {
	return xxx_messageInfo_OspfShTopPath.Size(m)
}
func (m *OspfShTopPath) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopPath.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopPath proto.InternalMessageInfo

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func (m *OspfShTopPath) GetSrMicroloopAvoidancePath() *OspfShSrUloopPath {
	if m != nil {
		return m.SrMicroloopAvoidancePath
	}
	return nil
}

type OspfShTopology struct {
	RoutePrefix          string           `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
	RoutePrefixLength    uint32           `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength,proto3" json:"route_prefix_length,omitempty"`
	RouteMetric          uint32           `protobuf:"varint,52,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	RouteType            string           `protobuf:"bytes,53,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	RouteConnected       bool             `protobuf:"varint,54,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	RouteInfo            *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo,proto3" json:"route_info,omitempty"`
	RoutePath            []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OspfShTopology) Reset()         { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()    {}
func (*OspfShTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_4169ff6e5af11206, []int{6}
}

func (m *OspfShTopology) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShTopology.Unmarshal(m, b)
}
func (m *OspfShTopology) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShTopology.Marshal(b, m, deterministic)
}
func (m *OspfShTopology) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShTopology.Merge(m, src)
}
func (m *OspfShTopology) XXX_Size() int {
	return xxx_messageInfo_OspfShTopology.Size(m)
}
func (m *OspfShTopology) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShTopology.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShTopology proto.InternalMessageInfo

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePath() []*OspfShTopPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_topology")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_4169ff6e5af11206) }

var fileDescriptor_4169ff6e5af11206 = []byte{
	// 1071 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x4e, 0x2c, 0x45,
	0x14, 0x4e, 0x0b, 0x17, 0x66, 0x6a, 0xe0, 0xc2, 0x2d, 0x11, 0x9a, 0xeb, 0x55, 0x70, 0x4c, 0x94,
	0x18, 0x33, 0xb9, 0x01, 0xfc, 0x5f, 0x18, 0xa2, 0x10, 0x27, 0x02, 0x21, 0x3d, 0x5c, 0x13, 0x57,
	0x95, 0x9a, 0xee, 0xd3, 0x50, 0x49, 0x77, 0x57, 0xa7, 0xaa, 0x1a, 0xe1, 0xbe, 0x84, 0x3e, 0x82,
	0xc6, 0xad, 0x31, 0xc6, 0x98, 0xf8, 0x1e, 0xc6, 0xc4, 0x85, 0xbe, 0x87, 0x1b, 0x17, 0xa6, 0x4e,
	0x55, 0x4f, 0x77, 0x43, 0xe2, 0x5a, 0x76, 0x35, 0xdf, 0xf9, 0xe9, 0xf3, 0xd5, 0xf9, 0xab, 0x21,
	0xeb, 0x52, 0x97, 0x29, 0xd3, 0x97, 0xcc, 0xc8, 0x52, 0x66, 0xf2, 0xe2, 0x66, 0x54, 0x2a, 0x69,
	0x24, 0xfd, 0x36, 0x88, 0x85, 0x8e, 0x25, 0x13, 0x52, 0xb3, 0x6b, 0xc5, 0x44, 0x79, 0xb5, 0xcf,
	0x50, 0x55, 0x96, 0xa0, 0x46, 0xf6, 0x64, 0x15, 0x63, 0xd0, 0x1a, 0x74, 0x7d, 0x1a, 0x25, 0x90,
	0xf2, 0x2a, 0x33, 0xec, 0x4a, 0xa5, 0x23, 0x25, 0x2b, 0x03, 0x4c, 0x14, 0xa9, 0x54, 0x39, 0x37,
	0x42, 0x16, 0x1e, 0xe1, 0x0a, 0xb8, 0x6e, 0x9d, 0x47, 0x79, 0x95, 0x19, 0x11, 0x73, 0x6d, 0x98,
	0x28, 0x0c, 0x8f, 0x0d, 0x6b, 0xeb, 0xfd, 0x87, 0x70, 0xf8, 0x4d, 0x40, 0x5e, 0xba, 0x1d, 0x3d,
	0xfb, 0xfc, 0xf0, 0xcb, 0x09, 0x7d, 0x8d, 0x2c, 0xf9, 0x90, 0x58, 0xc1, 0x73, 0x08, 0x83, 0xed,
	0x60, 0xa7, 0x1f, 0x0d, 0x3c, 0x76, 0xca, 0x73, 0xa0, 0x1b, 0x64, 0xd1, 0x3a, 0x61, 0x22, 0x09,
	0x5f, 0xd8, 0x0e, 0x76, 0x96, 0xa3, 0x05, 0xfb, 0x73, 0x9c, 0xd0, 0x75, 0xb2, 0x50, 0x2a, 0x48,
	0xc5, 0x75, 0x38, 0x87, 0x56, 0xfe, 0x17, 0x7d, 0x9d, 0x2c, 0xbb, 0x13, 0xcb, 0xa0, 0xb8, 0x30,
	0x97, 0xe1, 0x3c, 0x9a, 0x2d, 0x39, 0xf0, 0x18, 0xb1, 0xe1, 0x11, 0x59, 0x9a, 0x45, 0x24, 0x72,
	0xb0, 0xce, 0x34, 0xc4, 0xb2, 0x48, 0x30, 0x84, 0xe5, 0xc8, 0xff, 0xa2, 0xaf, 0x12, 0x52, 0xf0,
	0x42, 0x7a, 0x99, 0x0b, 0xa0, 0x85, 0x0c, 0xbf, 0x5f, 0x24, 0xb4, 0x45, 0x8d, 0xc5, 0x32, 0xcf,
	0x65, 0x41, 0x87, 0x64, 0xb9, 0xe1, 0x6f, 0x43, 0x77, 0x5e, 0x07, 0x08, 0x1e, 0xb8, 0xf8, 0xdf,
	0x20, 0x2b, 0x4e, 0xc7, 0x00, 0xcb, 0xc1, 0x28, 0x11, 0x7b, 0xff, 0xce, 0xf4, 0x1c, 0x4e, 0x10,
	0xa4, 0x6f, 0x91, 0x47, 0x4e, 0x4f, 0x89, 0x29, 0xbb, 0x02, 0xa5, 0x85, 0x2c, 0x90, 0xf2, 0x72,
	0xe4, 0x1c, 0x44, 0x62, 0xfa, 0x85, 0x83, 0x1b, 0x5d, 0x1b, 0x52, 0xad, 0x6b, 0xf9, 0xcf, 0x7b,
	0xdd, 0x49, 0x99, 0xd6, 0xba, 0xfb, 0x64, 0xdd, 0xe9, 0xa6, 0x52, 0x7d, 0xc5, 0x55, 0xc2, 0x12,
	0xa1, 0x0d, 0x2f, 0x62, 0x08, 0x1f, 0xa0, 0xf3, 0x35, 0x94, 0x1e, 0x39, 0xe1, 0xa7, 0x5e, 0x66,
	0x33, 0xe6, 0xbf, 0x20, 0x2b, 0x15, 0x43, 0xb8, 0xd0, 0x22, 0x36, 0x41, 0x88, 0xfe, 0x19, 0xd4,
	0x51, 0x54, 0x65, 0xc2, 0x2d, 0x41, 0x91, 0x43, 0xb8, 0xb8, 0x1d, 0xec, 0x0c, 0x76, 0x7f, 0x0c,
	0x46, 0xff, 0xf7, 0x72, 0x1d, 0xb5, 0x0b, 0xc3, 0x5f, 0xdb, 0x33, 0x24, 0x72, 0x6e, 0x2b, 0xe5,
	0x8f, 0xa0, 0xce, 0x5b, 0xca, 0x45, 0xe6, 0xb8, 0xf5, 0xee, 0x27, 0x37, 0x57, 0x68, 0x47, 0x5c,
	0x64, 0xc8, 0xec, 0x6d, 0x42, 0x9b, 0xe2, 0x29, 0x95, 0x90, 0x4a, 0x98, 0x9b, 0xb0, 0x8f, 0x09,
	0x5e, 0xad, 0xab, 0xe7, 0xcc, 0xe3, 0x74, 0x44, 0x5e, 0xf4, 0x7e, 0x2b, 0x23, 0x19, 0x5c, 0xc7,
	0x59, 0x95, 0x40, 0x12, 0x92, 0xed, 0x60, 0xa7, 0x17, 0xb9, 0xfc, 0x1f, 0x54, 0x46, 0x1e, 0x7a,
	0x01, 0xfd, 0x98, 0x3c, 0xf1, 0xde, 0x95, 0x01, 0xe6, 0x3b, 0x54, 0xc1, 0x85, 0xd0, 0x06, 0x14,
	0x24, 0xe1, 0x00, 0x0d, 0x37, 0xdd, 0x77, 0x94, 0x81, 0x33, 0xd4, 0x88, 0x66, 0x0a, 0xf4, 0x03,
	0xb2, 0xd9, 0x72, 0x50, 0x4c, 0x55, 0xdb, 0x7a, 0x09, 0xa3, 0x5c, 0x9f, 0x59, 0x9f, 0x4e, 0x55,
	0x63, 0x3a, 0xfc, 0x3a, 0x20, 0x0f, 0x6b, 0xe6, 0x0a, 0x4a, 0x06, 0x19, 0x76, 0x0a, 0x94, 0x5c,
	0x28, 0x06, 0x19, 0xe4, 0x50, 0x98, 0xba, 0x4b, 0xfb, 0xd1, 0x8a, 0x13, 0x1c, 0x3a, 0x7c, 0x9c,
	0x60, 0xcd, 0x3b, 0xdd, 0x8c, 0x4f, 0x21, 0xf3, 0x6d, 0x3a, 0x70, 0xd8, 0xb1, 0x85, 0xf0, 0x36,
	0xba, 0xee, 0xcc, 0x4d, 0x09, 0xbe, 0x4d, 0x1f, 0x75, 0x1c, 0x9e, 0xdf, 0x94, 0x30, 0xfc, 0x61,
	0xae, 0x19, 0x89, 0x5a, 0xb1, 0x2a, 0x93, 0xb2, 0x64, 0x25, 0x37, 0x97, 0xf4, 0xaf, 0x80, 0xac,
	0xe6, 0x22, 0x56, 0x12, 0x21, 0x67, 0x19, 0x06, 0xdb, 0x73, 0x3b, 0x83, 0xdd, 0x9f, 0xee, 0x51,
	0x81, 0xb9, 0x6b, 0x8e, 0x56, 0x66, 0x54, 0x22, 0x64, 0x42, 0x3f, 0x22, 0x8f, 0x6f, 0xb3, 0x63,
	0x99, 0xd0, 0x86, 0x69, 0xf1, 0x1c, 0xfc, 0xcd, 0x6e, 0xdc, 0x32, 0x3a, 0x16, 0xda, 0x4c, 0xc4,
	0x73, 0xa0, 0x87, 0x64, 0xab, 0x31, 0x36, 0x55, 0x51, 0x40, 0x66, 0x43, 0x00, 0x95, 0xf2, 0x18,
	0xdc, 0x06, 0x71, 0xbb, 0xe0, 0xc9, 0x4c, 0xed, 0x1c, 0xb5, 0xc6, 0xb5, 0x12, 0xae, 0x94, 0xa7,
	0x64, 0xad, 0x71, 0xa3, 0xed, 0x90, 0x35, 0xb6, 0xe6, 0x71, 0x50, 0xf6, 0x22, 0x3a, 0x93, 0x4d,
	0x50, 0x34, 0x29, 0xd3, 0xe1, 0x77, 0x0f, 0xc8, 0x6a, 0x7b, 0xcc, 0x63, 0xa6, 0x9e, 0x92, 0xb5,
	0xfa, 0x3e, 0x3b, 0x21, 0xb8, 0x2a, 0x72, 0xbd, 0xd4, 0xfd, 0xf0, 0x5e, 0x3d, 0x72, 0x0b, 0xb8,
	0x36, 0xec, 0x52, 0x96, 0x8c, 0x27, 0x89, 0x02, 0xad, 0x91, 0x78, 0x3f, 0x72, 0x1d, 0x75, 0x0a,
	0xd7, 0xe6, 0x33, 0x59, 0x1e, 0x38, 0xd1, 0x9d, 0x89, 0xeb, 0x18, 0x76, 0x26, 0xee, 0x16, 0x71,
	0x3f, 0x59, 0xa6, 0xb9, 0x48, 0x90, 0x47, 0x3f, 0x22, 0x08, 0x1d, 0x5b, 0x84, 0x7e, 0x48, 0x1e,
	0x3b, 0x05, 0x1b, 0x38, 0x13, 0x9a, 0xe5, 0xad, 0xf4, 0xe1, 0xbc, 0xef, 0xf9, 0xe6, 0x39, 0xe3,
	0xe6, 0x72, 0xac, 0x4f, 0xac, 0x78, 0x8c, 0x52, 0xba, 0x4f, 0x36, 0xba, 0xb6, 0x55, 0x9c, 0xbb,
	0x1b, 0xc0, 0xe1, 0xdf, 0xf3, 0x51, 0x3b, 0xc3, 0x67, 0x71, 0x5e, 0xda, 0x53, 0x13, 0xb5, 0x5f,
	0x6d, 0x8b, 0xad, 0x3d, 0xe1, 0x17, 0xdb, 0x26, 0xe9, 0x65, 0x9a, 0xbb, 0x46, 0xe9, 0xa1, 0x78,
	0x31, 0xd3, 0xdc, 0xb6, 0x47, 0x7b, 0xe9, 0xf7, 0x3b, 0x4b, 0x7f, 0x8b, 0x0c, 0x50, 0xe0, 0x0a,
	0xd8, 0x4f, 0x1b, 0x62, 0xa1, 0x23, 0x44, 0xe8, 0xdf, 0x01, 0x79, 0x59, 0x2b, 0xd6, 0xe4, 0x97,
	0x5f, 0x49, 0x91, 0xd8, 0xd5, 0xe5, 0x42, 0x1e, 0xe0, 0xa8, 0xfe, 0xf5, 0x1e, 0x75, 0x52, 0x67,
	0x3c, 0x44, 0xa1, 0x56, 0x27, 0x35, 0xb7, 0x83, 0x9a, 0x9a, 0xbd, 0xf1, 0xe1, 0x3f, 0xf3, 0x9d,
	0x1a, 0xc5, 0x57, 0x56, 0x93, 0x06, 0xff, 0x54, 0xda, 0x6d, 0x15, 0x8f, 0x9b, 0xb0, 0xcd, 0x20,
	0xef, 0xbe, 0x9a, 0xf6, 0xfc, 0xe8, 0x6a, 0x34, 0xdd, 0xd3, 0xe9, 0x4e, 0x66, 0xf7, 0xef, 0x66,
	0xf6, 0x15, 0x42, 0xfc, 0xd3, 0xc6, 0xe6, 0xf6, 0x1d, 0xfc, 0x66, 0xdf, 0xbd, 0x6a, 0x6c, 0x76,
	0xdf, 0xac, 0x37, 0x68, 0x2c, 0x8b, 0x02, 0x62, 0x03, 0x49, 0xf8, 0x2e, 0x26, 0xf2, 0x21, 0xc2,
	0x9f, 0xd4, 0x28, 0xfd, 0x3d, 0xa8, 0x1d, 0xd9, 0x8b, 0x0e, 0xdf, 0xc3, 0xdc, 0xfd, 0x72, 0x9f,
	0xd6, 0xec, 0xec, 0x49, 0xe8, 0xf9, 0x8f, 0x8b, 0x54, 0xd2, 0xdf, 0x66, 0xb4, 0xb0, 0x24, 0xdf,
	0xc7, 0xe1, 0xfe, 0xf3, 0x3d, 0xa3, 0x85, 0xd5, 0xd8, 0x9f, 0xb5, 0xfe, 0x74, 0x01, 0xff, 0x8e,
	0xec, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x94, 0x4b, 0xaf, 0xa8, 0x0c, 0x00, 0x00,
}
