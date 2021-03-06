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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_route_areas_route_area_route_area_informations_route_area_information

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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.route_areas.route_area.route_area_informations.route_area_information.ospf_sh_topology")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_4169ff6e5af11206) }

var fileDescriptor_4169ff6e5af11206 = []byte{
	// 1064 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xdd, 0x4e, 0x24, 0xc5,
	0x17, 0x4f, 0xff, 0x61, 0x61, 0xa6, 0x06, 0x16, 0xb6, 0xfe, 0x08, 0xcd, 0xba, 0x0a, 0x8e, 0x89,
	0x12, 0x63, 0x26, 0x1b, 0xc0, 0xef, 0x0b, 0x43, 0x14, 0xe2, 0x44, 0x20, 0xa4, 0x87, 0x35, 0xf1,
	0xaa, 0x52, 0xd3, 0x7d, 0x1a, 0x2a, 0xe9, 0xee, 0xea, 0x54, 0x55, 0x23, 0xec, 0x4b, 0xac, 0x57,
	0x7a, 0x6f, 0x8c, 0x31, 0x26, 0xbe, 0x87, 0x31, 0xf1, 0x42, 0x9f, 0xc2, 0x1b, 0xdf, 0xc1, 0xd4,
	0xa9, 0xea, 0xe9, 0x6e, 0x88, 0xf7, 0xec, 0x5d, 0xf5, 0xef, 0x7c, 0xd4, 0xf9, 0x3e, 0xd5, 0x64,
	0x5d, 0xea, 0x32, 0x65, 0xfa, 0x92, 0x19, 0x59, 0xca, 0x4c, 0x5e, 0xdc, 0x8c, 0x4a, 0x25, 0x8d,
	0xa4, 0xdf, 0x05, 0xb1, 0xd0, 0xb1, 0x64, 0x42, 0x6a, 0x76, 0xad, 0x98, 0x28, 0xaf, 0xf6, 0x19,
	0xb2, 0xca, 0x12, 0xd4, 0xc8, 0x9e, 0x2c, 0x63, 0x0c, 0x5a, 0x83, 0xae, 0x4f, 0xa3, 0x04, 0x52,
	0x5e, 0x65, 0x86, 0x5d, 0xa9, 0x74, 0xa4, 0x64, 0x65, 0x80, 0x89, 0x22, 0x95, 0x2a, 0xe7, 0x46,
	0xc8, 0xc2, 0x23, 0x5c, 0x01, 0xd7, 0xad, 0x73, 0xeb, 0xd8, 0xe6, 0xd6, 0xff, 0x81, 0x0f, 0xbf,
	0x0d, 0xc8, 0x2b, 0xb7, 0x6d, 0x66, 0x5f, 0x1e, 0x7e, 0x3d, 0xa1, 0x6f, 0x90, 0x25, 0x6f, 0x08,
	0x2b, 0x78, 0x0e, 0x61, 0xb0, 0x1d, 0xec, 0xf4, 0xa3, 0x81, 0xc7, 0x4e, 0x79, 0x0e, 0x74, 0x83,
	0x2c, 0x3a, 0x85, 0x49, 0xf8, 0xbf, 0xed, 0x60, 0x67, 0x39, 0x5a, 0xb0, 0x9f, 0xe3, 0x84, 0xae,
	0x93, 0x85, 0x52, 0x41, 0x2a, 0xae, 0xc3, 0x39, 0x94, 0xf2, 0x5f, 0xf4, 0x4d, 0xb2, 0xec, 0x4e,
	0x2c, 0x83, 0xe2, 0xc2, 0x5c, 0x86, 0xf3, 0x28, 0xb6, 0xe4, 0xc0, 0x63, 0xc4, 0x86, 0x47, 0x64,
	0x69, 0x66, 0x91, 0xc8, 0xc1, 0x2a, 0xd3, 0x10, 0xcb, 0x22, 0x41, 0x13, 0x96, 0x23, 0xff, 0x45,
	0x5f, 0x27, 0xa4, 0xe0, 0x85, 0xf4, 0x34, 0x67, 0x40, 0x0b, 0x19, 0xbe, 0x58, 0x24, 0xb4, 0xe5,
	0x1a, 0x8b, 0x65, 0x9e, 0xcb, 0x82, 0x0e, 0xc9, 0x72, 0x3b, 0x16, 0xb5, 0xd6, 0x01, 0x82, 0x07,
	0xce, 0xfe, 0xb7, 0xc8, 0x8a, 0xe3, 0x31, 0xc0, 0x72, 0x30, 0x4a, 0xc4, 0x5e, 0xbf, 0x13, 0x3d,
	0x87, 0x13, 0x04, 0xe9, 0x3b, 0xe4, 0x91, 0xe3, 0x53, 0x62, 0xca, 0xae, 0x40, 0x69, 0x21, 0x0b,
	0x74, 0x79, 0x39, 0x72, 0x0a, 0x22, 0x31, 0xfd, 0xca, 0xc1, 0x0d, 0xaf, 0x35, 0xa9, 0xe6, 0xb5,
	0xfe, 0xcf, 0x7b, 0xde, 0x49, 0x99, 0xd6, 0xbc, 0xfb, 0x64, 0xdd, 0xf1, 0xa6, 0x52, 0x7d, 0xc3,
	0x55, 0xc2, 0x12, 0xa1, 0x0d, 0x2f, 0x62, 0x08, 0x1f, 0xa0, 0xf2, 0x35, 0xa4, 0x1e, 0x39, 0xe2,
	0xe7, 0x9e, 0x66, 0x33, 0xe6, 0x6f, 0x90, 0x95, 0x8a, 0x21, 0x5c, 0x68, 0x39, 0x36, 0x41, 0x88,
	0xfe, 0x19, 0xd4, 0x56, 0x54, 0x65, 0xc2, 0xad, 0x83, 0x22, 0x87, 0x70, 0x71, 0x3b, 0xd8, 0x19,
	0xec, 0xfe, 0x18, 0x8c, 0xee, 0x67, 0x91, 0x8e, 0xda, 0xe5, 0xe0, 0x83, 0xf5, 0x0c, 0xcd, 0x3f,
	0xb7, 0xf5, 0xf1, 0x47, 0x50, 0x67, 0x2b, 0xe5, 0x22, 0x73, 0x1e, 0xf5, 0x5e, 0x26, 0x8f, 0x5c,
	0x51, 0x1d, 0x71, 0x91, 0xa1, 0x3f, 0xef, 0x12, 0xda, 0x14, 0x4a, 0xa9, 0x84, 0x54, 0xc2, 0xdc,
	0x84, 0x7d, 0x4c, 0xe6, 0x6a, 0x5d, 0x29, 0x67, 0x1e, 0xa7, 0x23, 0xf2, 0x7f, 0x7f, 0x47, 0x65,
	0x24, 0x83, 0xeb, 0x38, 0xab, 0x12, 0x48, 0x42, 0xb2, 0x1d, 0xec, 0xf4, 0x22, 0x97, 0xeb, 0x83,
	0xca, 0xc8, 0x43, 0x4f, 0xa0, 0x9f, 0x92, 0x27, 0x5e, 0xbb, 0x32, 0xc0, 0x7c, 0x37, 0x2a, 0xb8,
	0x10, 0xda, 0x80, 0x82, 0x24, 0x1c, 0xa0, 0xe0, 0xa6, 0xbb, 0x47, 0x19, 0x38, 0x43, 0x8e, 0x68,
	0xc6, 0x40, 0x3f, 0x22, 0x9b, 0x2d, 0x05, 0xc5, 0x54, 0xb5, 0xa5, 0x97, 0xd0, 0xca, 0xf5, 0x99,
	0xf4, 0xe9, 0x54, 0x35, 0xa2, 0xc3, 0x17, 0x01, 0x79, 0x58, 0x7b, 0xae, 0xa0, 0x64, 0x90, 0x61,
	0x57, 0x40, 0xc9, 0x85, 0x62, 0x90, 0x41, 0x0e, 0x85, 0xa9, 0x3b, 0xb2, 0x1f, 0xad, 0x38, 0xc2,
	0xa1, 0xc3, 0xc7, 0x09, 0xd6, 0xb7, 0xe3, 0xcd, 0xf8, 0x14, 0x32, 0xdf, 0x92, 0x03, 0x87, 0x1d,
	0x5b, 0x08, 0xa3, 0xd1, 0x55, 0x67, 0x6e, 0x4a, 0xf0, 0x2d, 0xf9, 0xa8, 0xa3, 0xf0, 0xfc, 0xa6,
	0x84, 0xe1, 0x0f, 0x73, 0xcd, 0xf8, 0xd3, 0x8a, 0x55, 0x99, 0x94, 0x25, 0x2b, 0xb9, 0xb9, 0xa4,
	0x7f, 0x05, 0x64, 0x35, 0x17, 0xb1, 0x92, 0x08, 0x39, 0xc9, 0x30, 0xd8, 0x9e, 0xdb, 0x19, 0xec,
	0xfe, 0x74, 0xef, 0xcb, 0xca, 0x05, 0x37, 0x5a, 0x99, 0x39, 0x10, 0xa1, 0xfd, 0xf4, 0x13, 0xf2,
	0xf8, 0xb6, 0x4f, 0x2c, 0x13, 0xda, 0x30, 0x2d, 0x9e, 0x83, 0x8f, 0xe7, 0xc6, 0x2d, 0xa1, 0x63,
	0xa1, 0xcd, 0x44, 0x3c, 0x07, 0x7a, 0x48, 0xb6, 0x1a, 0x61, 0x53, 0x15, 0x05, 0x64, 0x4c, 0x14,
	0x06, 0x54, 0xca, 0x63, 0x70, 0x3b, 0xc2, 0x4d, 0xfb, 0x27, 0x33, 0xb6, 0x73, 0xe4, 0x1a, 0xd7,
	0x4c, 0xb8, 0x34, 0x9e, 0x92, 0xb5, 0x46, 0x8d, 0xb6, 0x63, 0xd4, 0xd8, 0x4a, 0xc7, 0x51, 0xd8,
	0x8b, 0xe8, 0x8c, 0x36, 0x41, 0xd2, 0xa4, 0x4c, 0x87, 0xdf, 0x3f, 0x20, 0xab, 0xed, 0x41, 0x8e,
	0xf9, 0x79, 0x4a, 0xd6, 0xea, 0x28, 0x76, 0x4c, 0x70, 0xb5, 0xe3, 0x3a, 0xa8, 0x7b, 0xf1, 0x5e,
	0x3d, 0x54, 0x0b, 0xb8, 0x36, 0xec, 0x52, 0x96, 0x8c, 0x27, 0x89, 0x02, 0xad, 0xd1, 0xf1, 0x7e,
	0xe4, 0xfa, 0xe8, 0x14, 0xae, 0xcd, 0x17, 0xb2, 0x3c, 0x70, 0xa4, 0x3b, 0x33, 0xd5, 0x79, 0xd8,
	0x99, 0xa9, 0x5b, 0xc4, 0x7d, 0xb2, 0x4c, 0x73, 0x91, 0xa0, 0x1f, 0xfd, 0x88, 0x20, 0x74, 0x6c,
	0x11, 0xfa, 0x31, 0x79, 0xec, 0x18, 0xac, 0xe1, 0x4c, 0x68, 0x96, 0xc7, 0x5c, 0x1b, 0x6b, 0x38,
	0x8f, 0x0d, 0x4e, 0xf4, 0x9e, 0x6f, 0x99, 0x33, 0x6e, 0x2e, 0xc7, 0xfa, 0xc4, 0x92, 0xc7, 0x48,
	0xa5, 0xfb, 0x64, 0xa3, 0x2b, 0x5b, 0xc5, 0xb9, 0x8b, 0x00, 0x8e, 0xf7, 0x9e, 0xb7, 0xda, 0x09,
	0x3e, 0x8b, 0xf3, 0xd2, 0x9e, 0x1a, 0xab, 0xfd, 0xf2, 0x5a, 0x6c, 0x6d, 0x02, 0xbf, 0xba, 0x36,
	0x49, 0x2f, 0xd3, 0xdc, 0xb5, 0x47, 0x0f, 0xc9, 0x8b, 0x99, 0xe6, 0xb6, 0x29, 0xda, 0x6b, 0xbd,
	0xdf, 0x59, 0xeb, 0x5b, 0x64, 0x80, 0x04, 0x57, 0x70, 0x7e, 0xc6, 0x10, 0x0b, 0x1d, 0x21, 0x42,
	0xff, 0x09, 0xc8, 0xab, 0x5a, 0xb1, 0x26, 0xbf, 0xfc, 0x4a, 0x8a, 0xc4, 0x2e, 0x27, 0x67, 0xf2,
	0x00, 0xc7, 0xf2, 0xaf, 0xf7, 0xbe, 0x7f, 0x3a, 0xa3, 0x20, 0x0a, 0xb5, 0x3a, 0xa9, 0x3d, 0x3a,
	0xa8, 0x1d, 0xb2, 0x71, 0x1e, 0xfe, 0x3d, 0xdf, 0xa9, 0x4c, 0x7c, 0x3d, 0x35, 0xc1, 0xf7, 0x4f,
	0xa0, 0xdd, 0x56, 0xc9, 0xb8, 0x69, 0xda, 0x0c, 0xed, 0xee, 0x6b, 0x68, 0xcf, 0x8f, 0xa9, 0x86,
	0xd3, 0x3d, 0x89, 0xee, 0xe4, 0x73, 0xff, 0x6e, 0x3e, 0x5f, 0x23, 0xc4, 0x3f, 0x59, 0x6c, 0x46,
	0xdf, 0xc3, 0x3b, 0xfb, 0xee, 0xb5, 0x62, 0x73, 0xfa, 0x76, 0xbd, 0x23, 0x63, 0x59, 0x14, 0x10,
	0x1b, 0x48, 0xc2, 0xf7, 0x31, 0x7d, 0x0f, 0x11, 0xfe, 0xac, 0x46, 0xe9, 0xef, 0x41, 0xad, 0xc8,
	0x06, 0x26, 0xfc, 0x00, 0x33, 0xf6, 0xcb, 0xfd, 0x5f, 0xa4, 0xb3, 0x07, 0x9e, 0xf7, 0x7a, 0x5c,
	0xa4, 0x92, 0xfe, 0x36, 0x73, 0x06, 0xcb, 0xef, 0x43, 0x1c, 0xdf, 0x3f, 0xbf, 0x14, 0xce, 0x60,
	0xe5, 0xf5, 0x67, 0xcd, 0x3d, 0x5d, 0xc0, 0x1f, 0x89, 0xbd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0xca, 0xab, 0x11, 0x62, 0x0c, 0x00, 0x00,
}
