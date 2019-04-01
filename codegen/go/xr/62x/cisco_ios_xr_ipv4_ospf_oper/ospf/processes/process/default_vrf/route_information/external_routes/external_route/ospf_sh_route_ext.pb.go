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
// source: ospf_sh_route_ext.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_external_routes_external_route

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

type OspfShRouteExt_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRouteExt_KEYS) Reset()         { *m = OspfShRouteExt_KEYS{} }
func (m *OspfShRouteExt_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteExt_KEYS) ProtoMessage()    {}
func (*OspfShRouteExt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5115045e98cfb91c, []int{0}
}

func (m *OspfShRouteExt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteExt_KEYS.Unmarshal(m, b)
}
func (m *OspfShRouteExt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteExt_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRouteExt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteExt_KEYS.Merge(m, src)
}
func (m *OspfShRouteExt_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteExt_KEYS.Size(m)
}
func (m *OspfShRouteExt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteExt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteExt_KEYS proto.InternalMessageInfo

func (m *OspfShRouteExt_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRouteExt_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShRouteExt_KEYS) GetPrefixLength() uint32 {
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
	return fileDescriptor_5115045e98cfb91c, []int{1}
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
	return fileDescriptor_5115045e98cfb91c, []int{2}
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
	return fileDescriptor_5115045e98cfb91c, []int{3}
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
	return fileDescriptor_5115045e98cfb91c, []int{4}
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
	return fileDescriptor_5115045e98cfb91c, []int{5}
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
	RoutePrefix          string           `protobuf:"bytes,1,opt,name=route_prefix,json=routePrefix,proto3" json:"route_prefix,omitempty"`
	RoutePrefixLength    uint32           `protobuf:"varint,2,opt,name=route_prefix_length,json=routePrefixLength,proto3" json:"route_prefix_length,omitempty"`
	RouteMetric          uint32           `protobuf:"varint,3,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	RouteType            string           `protobuf:"bytes,4,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	RouteConnected       bool             `protobuf:"varint,5,opt,name=route_connected,json=routeConnected,proto3" json:"route_connected,omitempty"`
	RouteInfo            *OspfShTopCommon `protobuf:"bytes,6,opt,name=route_info,json=routeInfo,proto3" json:"route_info,omitempty"`
	RoutePath            []*OspfShTopPath `protobuf:"bytes,7,rep,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *OspfShTopology) Reset()         { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()    {}
func (*OspfShTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_5115045e98cfb91c, []int{6}
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

type OspfShRedistProto struct {
	ProtocolType         string   `protobuf:"bytes,1,opt,name=protocol_type,json=protocolType,proto3" json:"protocol_type,omitempty"`
	IsisInstanceId       string   `protobuf:"bytes,2,opt,name=isis_instance_id,json=isisInstanceId,proto3" json:"isis_instance_id,omitempty"`
	OspfProcessId        string   `protobuf:"bytes,3,opt,name=ospf_process_id,json=ospfProcessId,proto3" json:"ospf_process_id,omitempty"`
	BgpAsNumber          string   `protobuf:"bytes,4,opt,name=bgp_as_number,json=bgpAsNumber,proto3" json:"bgp_as_number,omitempty"`
	EigrpAsNumber        string   `protobuf:"bytes,5,opt,name=eigrp_as_number,json=eigrpAsNumber,proto3" json:"eigrp_as_number,omitempty"`
	ConnectedInstance    string   `protobuf:"bytes,6,opt,name=connected_instance,json=connectedInstance,proto3" json:"connected_instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRedistProto) Reset()         { *m = OspfShRedistProto{} }
func (m *OspfShRedistProto) String() string { return proto.CompactTextString(m) }
func (*OspfShRedistProto) ProtoMessage()    {}
func (*OspfShRedistProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5115045e98cfb91c, []int{7}
}

func (m *OspfShRedistProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRedistProto.Unmarshal(m, b)
}
func (m *OspfShRedistProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRedistProto.Marshal(b, m, deterministic)
}
func (m *OspfShRedistProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRedistProto.Merge(m, src)
}
func (m *OspfShRedistProto) XXX_Size() int {
	return xxx_messageInfo_OspfShRedistProto.Size(m)
}
func (m *OspfShRedistProto) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRedistProto.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRedistProto proto.InternalMessageInfo

func (m *OspfShRedistProto) GetProtocolType() string {
	if m != nil {
		return m.ProtocolType
	}
	return ""
}

func (m *OspfShRedistProto) GetIsisInstanceId() string {
	if m != nil {
		return m.IsisInstanceId
	}
	return ""
}

func (m *OspfShRedistProto) GetOspfProcessId() string {
	if m != nil {
		return m.OspfProcessId
	}
	return ""
}

func (m *OspfShRedistProto) GetBgpAsNumber() string {
	if m != nil {
		return m.BgpAsNumber
	}
	return ""
}

func (m *OspfShRedistProto) GetEigrpAsNumber() string {
	if m != nil {
		return m.EigrpAsNumber
	}
	return ""
}

func (m *OspfShRedistProto) GetConnectedInstance() string {
	if m != nil {
		return m.ConnectedInstance
	}
	return ""
}

type OspfShRouteExtendedComm struct {
	ExtendedCommunityDomainIdValue []uint32 `protobuf:"varint,1,rep,packed,name=extended_community_domain_id_value,json=extendedCommunityDomainIdValue,proto3" json:"extended_community_domain_id_value,omitempty"`
	ExtendedCommunitylDomainIdType uint32   `protobuf:"varint,2,opt,name=extended_communityl_domain_id_type,json=extendedCommunitylDomainIdType,proto3" json:"extended_communityl_domain_id_type,omitempty"`
	ExtendedCommunityAreaId        uint32   `protobuf:"varint,3,opt,name=extended_community_area_id,json=extendedCommunityAreaId,proto3" json:"extended_community_area_id,omitempty"`
	ExtendedCommunityRouterId      string   `protobuf:"bytes,4,opt,name=extended_community_router_id,json=extendedCommunityRouterId,proto3" json:"extended_community_router_id,omitempty"`
	ExtendedCommunityRouteType     uint32   `protobuf:"varint,5,opt,name=extended_community_route_type,json=extendedCommunityRouteType,proto3" json:"extended_community_route_type,omitempty"`
	ExtendedCommunityOptions       uint32   `protobuf:"varint,6,opt,name=extended_community_options,json=extendedCommunityOptions,proto3" json:"extended_community_options,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *OspfShRouteExtendedComm) Reset()         { *m = OspfShRouteExtendedComm{} }
func (m *OspfShRouteExtendedComm) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteExtendedComm) ProtoMessage()    {}
func (*OspfShRouteExtendedComm) Descriptor() ([]byte, []int) {
	return fileDescriptor_5115045e98cfb91c, []int{8}
}

func (m *OspfShRouteExtendedComm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteExtendedComm.Unmarshal(m, b)
}
func (m *OspfShRouteExtendedComm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteExtendedComm.Marshal(b, m, deterministic)
}
func (m *OspfShRouteExtendedComm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteExtendedComm.Merge(m, src)
}
func (m *OspfShRouteExtendedComm) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteExtendedComm.Size(m)
}
func (m *OspfShRouteExtendedComm) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteExtendedComm.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteExtendedComm proto.InternalMessageInfo

func (m *OspfShRouteExtendedComm) GetExtendedCommunityDomainIdValue() []uint32 {
	if m != nil {
		return m.ExtendedCommunityDomainIdValue
	}
	return nil
}

func (m *OspfShRouteExtendedComm) GetExtendedCommunitylDomainIdType() uint32 {
	if m != nil {
		return m.ExtendedCommunitylDomainIdType
	}
	return 0
}

func (m *OspfShRouteExtendedComm) GetExtendedCommunityAreaId() uint32 {
	if m != nil {
		return m.ExtendedCommunityAreaId
	}
	return 0
}

func (m *OspfShRouteExtendedComm) GetExtendedCommunityRouterId() string {
	if m != nil {
		return m.ExtendedCommunityRouterId
	}
	return ""
}

func (m *OspfShRouteExtendedComm) GetExtendedCommunityRouteType() uint32 {
	if m != nil {
		return m.ExtendedCommunityRouteType
	}
	return 0
}

func (m *OspfShRouteExtendedComm) GetExtendedCommunityOptions() uint32 {
	if m != nil {
		return m.ExtendedCommunityOptions
	}
	return 0
}

type OspfShRouteExt struct {
	RouteInformation       *OspfShTopology          `protobuf:"bytes,50,opt,name=route_information,json=routeInformation,proto3" json:"route_information,omitempty"`
	ProtocolName           *OspfShRedistProto       `protobuf:"bytes,51,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	RouteExtendedCommunity *OspfShRouteExtendedComm `protobuf:"bytes,52,opt,name=route_extended_community,json=routeExtendedCommunity,proto3" json:"route_extended_community,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *OspfShRouteExt) Reset()         { *m = OspfShRouteExt{} }
func (m *OspfShRouteExt) String() string { return proto.CompactTextString(m) }
func (*OspfShRouteExt) ProtoMessage()    {}
func (*OspfShRouteExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_5115045e98cfb91c, []int{9}
}

func (m *OspfShRouteExt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRouteExt.Unmarshal(m, b)
}
func (m *OspfShRouteExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRouteExt.Marshal(b, m, deterministic)
}
func (m *OspfShRouteExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRouteExt.Merge(m, src)
}
func (m *OspfShRouteExt) XXX_Size() int {
	return xxx_messageInfo_OspfShRouteExt.Size(m)
}
func (m *OspfShRouteExt) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRouteExt.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRouteExt proto.InternalMessageInfo

func (m *OspfShRouteExt) GetRouteInformation() *OspfShTopology {
	if m != nil {
		return m.RouteInformation
	}
	return nil
}

func (m *OspfShRouteExt) GetProtocolName() *OspfShRedistProto {
	if m != nil {
		return m.ProtocolName
	}
	return nil
}

func (m *OspfShRouteExt) GetRouteExtendedCommunity() *OspfShRouteExtendedComm {
	if m != nil {
		return m.RouteExtendedCommunity
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRouteExt_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_route_ext_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_topology")
	proto.RegisterType((*OspfShRedistProto)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_redist_proto")
	proto.RegisterType((*OspfShRouteExtendedComm)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_route_extended_comm")
	proto.RegisterType((*OspfShRouteExt)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.external_routes.external_route.ospf_sh_route_ext")
}

func init() { proto.RegisterFile("ospf_sh_route_ext.proto", fileDescriptor_5115045e98cfb91c) }

var fileDescriptor_5115045e98cfb91c = []byte{
	// 1377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6f, 0x24, 0xb5,
	0x12, 0x56, 0x27, 0xd9, 0x24, 0xe3, 0x49, 0x76, 0x13, 0xbf, 0xbc, 0xa4, 0x93, 0xfd, 0x95, 0x37,
	0x4f, 0x82, 0x11, 0x82, 0xd1, 0x2a, 0x9b, 0x0b, 0x2c, 0xd2, 0x2a, 0xda, 0x4d, 0xc4, 0x40, 0x36,
	0x44, 0x9d, 0xec, 0x4a, 0x9c, 0x2c, 0x4f, 0xb7, 0x27, 0xb1, 0xd4, 0xdd, 0x6e, 0xd9, 0x9e, 0x30,
	0x59, 0x09, 0x21, 0x84, 0x40, 0x1c, 0x90, 0xb8, 0x73, 0x43, 0x88, 0x03, 0x70, 0xe0, 0xc2, 0x15,
	0x71, 0x83, 0x7f, 0x87, 0x33, 0x27, 0xe4, 0x2a, 0x77, 0xf7, 0x4c, 0x26, 0x5c, 0x19, 0x6e, 0xdd,
	0x5f, 0x95, 0xcb, 0x9f, 0x5d, 0x55, 0x9f, 0x8b, 0x6c, 0x28, 0x53, 0xf4, 0x99, 0x39, 0x67, 0x5a,
	0x0d, 0xac, 0x60, 0x62, 0x68, 0x3b, 0x85, 0x56, 0x56, 0x51, 0x15, 0x4b, 0x13, 0x2b, 0x26, 0x95,
	0x61, 0x43, 0xcd, 0x64, 0x71, 0xb1, 0xcb, 0xc0, 0x55, 0x15, 0x42, 0x77, 0xdc, 0x97, 0xf3, 0x8b,
	0x85, 0x31, 0xc2, 0x94, 0x5f, 0x9d, 0x44, 0xf4, 0xf9, 0x20, 0xb5, 0xec, 0x42, 0xf7, 0x3b, 0x18,
	0x4f, 0xe6, 0x7d, 0xa5, 0x33, 0x6e, 0xa5, 0xca, 0x3b, 0x62, 0x68, 0x85, 0xce, 0x79, 0x8a, 0x5b,
	0x99, 0x2b, 0xff, 0xad, 0x21, 0x59, 0x9f, 0xe0, 0xc2, 0xde, 0xdb, 0xff, 0xe0, 0x84, 0xfe, 0x8f,
	0x2c, 0xf9, 0x1d, 0x58, 0xce, 0x33, 0x11, 0x06, 0xdb, 0x41, 0xbb, 0x11, 0x35, 0x3d, 0x76, 0xc4,
	0x33, 0x41, 0xd7, 0xc9, 0x7c, 0xa1, 0x45, 0x5f, 0x0e, 0xc3, 0x19, 0x30, 0xfa, 0x3f, 0xfa, 0x7f,
	0xb2, 0x8c, 0x5f, 0x2c, 0x15, 0xf9, 0x99, 0x3d, 0x0f, 0x67, 0xb7, 0x83, 0xf6, 0x72, 0xb4, 0x84,
	0xe0, 0x21, 0x60, 0xad, 0x03, 0xb2, 0x54, 0xee, 0x6c, 0x25, 0x06, 0x33, 0x22, 0x56, 0x79, 0x02,
	0x3b, 0x2d, 0x47, 0xfe, 0x8f, 0xde, 0x23, 0x24, 0xe7, 0xb9, 0xf2, 0xb6, 0x19, 0xb0, 0x8d, 0x20,
	0xad, 0x5f, 0xe7, 0x09, 0xad, 0x02, 0xa9, 0x82, 0xc5, 0x2a, 0xcb, 0x54, 0x4e, 0x5b, 0x64, 0x19,
	0x0f, 0xc4, 0xb5, 0xe0, 0x4c, 0x96, 0x51, 0x9b, 0x00, 0xee, 0x69, 0xc1, 0xbb, 0x09, 0x7d, 0x85,
	0xdc, 0x42, 0x1f, 0x2b, 0x58, 0x26, 0xac, 0x96, 0xb1, 0x8f, 0x8f, 0x4b, 0x4f, 0xc5, 0x33, 0x00,
	0xe9, 0x6b, 0x64, 0x15, 0xfd, 0xb4, 0xec, 0xb1, 0x0b, 0xa1, 0x8d, 0x54, 0xb9, 0x3f, 0x13, 0x06,
	0x88, 0x64, 0xef, 0x05, 0xc2, 0xb5, 0xaf, 0xa3, 0x54, 0xfa, 0xce, 0x6d, 0x07, 0xed, 0x39, 0xef,
	0x7b, 0x52, 0xf4, 0x4b, 0xdf, 0x5d, 0xb2, 0x8e, 0xbe, 0x7d, 0xa5, 0x3f, 0xe4, 0x3a, 0x61, 0x89,
	0x34, 0x96, 0xe7, 0xb1, 0x08, 0x6f, 0x40, 0xf0, 0x35, 0xb0, 0x1e, 0xa0, 0xf1, 0xa9, 0xb7, 0xb9,
	0xc4, 0xf8, 0x1d, 0xd4, 0x40, 0xc7, 0x22, 0x9c, 0x1f, 0x39, 0xd8, 0x09, 0x40, 0xf4, 0xfb, 0xa0,
	0x64, 0x31, 0x28, 0x12, 0xee, 0x0e, 0x28, 0x33, 0x11, 0x2e, 0x6c, 0x07, 0xed, 0xe6, 0xce, 0x47,
	0x9d, 0x7f, 0xb8, 0xc6, 0x3a, 0xa3, 0x69, 0xf6, 0x97, 0xf0, 0x1c, 0x68, 0x9d, 0xba, 0xbc, 0x7f,
	0x17, 0x94, 0x59, 0xe8, 0x73, 0x99, 0x22, 0xd3, 0xc5, 0x7f, 0x03, 0x53, 0x2c, 0x82, 0x03, 0x2e,
	0x53, 0xe0, 0xf9, 0x3a, 0xa1, 0x75, 0x62, 0x0b, 0x2d, 0x95, 0x96, 0xf6, 0x32, 0x6c, 0xc0, 0xe5,
	0xaf, 0x94, 0x99, 0x3d, 0xf6, 0x38, 0xed, 0x90, 0xff, 0xf8, 0xf2, 0x1b, 0x58, 0xc5, 0xc4, 0x30,
	0x4e, 0x07, 0x89, 0x48, 0x42, 0xb2, 0x1d, 0xb4, 0x17, 0x23, 0xcc, 0xcd, 0xde, 0xc0, 0xaa, 0x7d,
	0x6f, 0xa0, 0x8f, 0xc9, 0x1d, 0x1f, 0x5d, 0x5b, 0xc1, 0x7c, 0xf7, 0x68, 0x71, 0x26, 0x8d, 0x15,
	0x5a, 0x24, 0x61, 0x13, 0x16, 0x6e, 0xe2, 0x3e, 0xda, 0x8a, 0x63, 0xf0, 0x88, 0x2a, 0x07, 0xfa,
	0x26, 0xd9, 0x1c, 0x09, 0x90, 0xf7, 0xf4, 0xe8, 0xea, 0x25, 0x60, 0xb9, 0x5e, 0xad, 0x3e, 0xea,
	0xe9, 0x7a, 0x69, 0xeb, 0xab, 0x80, 0xdc, 0xac, 0x44, 0x40, 0x14, 0x4c, 0xa4, 0x50, 0xc5, 0xa2,
	0xe0, 0x52, 0x33, 0x91, 0x8a, 0x4c, 0xe4, 0xb6, 0xec, 0xa0, 0x46, 0x74, 0x0b, 0x0d, 0xfb, 0x88,
	0x77, 0x13, 0xa8, 0x47, 0xf4, 0x4d, 0x79, 0x4f, 0xa4, 0xbe, 0x85, 0x9a, 0x88, 0x1d, 0x3a, 0x08,
	0x6e, 0x63, 0x3c, 0x9c, 0xbd, 0x2c, 0x84, 0x6f, 0xa1, 0xd5, 0xb1, 0x80, 0xa7, 0x97, 0x85, 0x68,
	0xfd, 0x39, 0x43, 0xfe, 0x5b, 0x32, 0x32, 0x9a, 0x0d, 0x52, 0xa5, 0x0a, 0x56, 0x70, 0x7b, 0x4e,
	0x7f, 0x08, 0xc8, 0x4a, 0x26, 0x63, 0xad, 0x00, 0xc2, 0x95, 0x61, 0xb0, 0x3d, 0xdb, 0x6e, 0xee,
	0x7c, 0x3c, 0xb5, 0x72, 0xc1, 0x4b, 0x8b, 0x6e, 0x55, 0xc4, 0x22, 0xe0, 0x45, 0x1f, 0x91, 0xad,
	0xab, 0x5c, 0x59, 0x2a, 0x8d, 0x65, 0x46, 0xbe, 0x14, 0xfe, 0x9e, 0x36, 0xae, 0x2c, 0x3a, 0x94,
	0xc6, 0x9e, 0xc8, 0x97, 0x82, 0xee, 0x93, 0xfb, 0xf5, 0x62, 0x3b, 0xc8, 0x73, 0x91, 0x32, 0x99,
	0x5b, 0xa1, 0xfb, 0x3c, 0x16, 0x28, 0xc9, 0xb3, 0x90, 0x90, 0x3b, 0x95, 0xdb, 0x29, 0x78, 0x75,
	0x4b, 0x27, 0xd0, 0xe8, 0x07, 0x64, 0xad, 0x0e, 0x63, 0x9c, 0x9c, 0x59, 0x57, 0xc1, 0x20, 0x49,
	0x8b, 0x11, 0xad, 0x6c, 0x27, 0x60, 0x3a, 0x29, 0xfa, 0xad, 0x3f, 0xe6, 0xc8, 0xca, 0xa8, 0xa0,
	0xc2, 0xbd, 0x3f, 0x20, 0x6b, 0xe5, 0xf5, 0x8c, 0x51, 0xc0, 0x9a, 0xc0, 0xce, 0x18, 0xdf, 0xf8,
	0x61, 0x29, 0x6e, 0xb9, 0x7b, 0x52, 0xce, 0x55, 0xc1, 0x78, 0x92, 0x68, 0x61, 0x8c, 0x7f, 0x2c,
	0xb0, 0x3f, 0x8e, 0xc4, 0xd0, 0xbe, 0xa3, 0x8a, 0x3d, 0x34, 0x4d, 0x68, 0x1b, 0x9e, 0x70, 0x4c,
	0xdb, 0xee, 0x13, 0xfc, 0x65, 0xa9, 0xe1, 0x32, 0x81, 0x73, 0x34, 0x22, 0x02, 0xd0, 0xa1, 0x43,
	0xe8, 0x5b, 0x64, 0x0b, 0x1d, 0x1c, 0x71, 0x26, 0x0d, 0xcb, 0x62, 0x6e, 0xac, 0x23, 0xce, 0x63,
	0x0b, 0xca, 0xba, 0xe8, 0x5b, 0xe1, 0x98, 0xdb, 0xf3, 0xae, 0x79, 0xe6, 0xcc, 0x5d, 0xb0, 0xd2,
	0x5d, 0xb2, 0x31, 0xbe, 0x76, 0x10, 0x67, 0x78, 0x03, 0x20, 0xb3, 0x8b, 0x9e, 0x35, 0x2e, 0x7c,
	0x1e, 0x67, 0x85, 0xfb, 0xaa, 0x59, 0xfb, 0x47, 0x64, 0x61, 0x44, 0x91, 0xfd, 0x13, 0xb2, 0x49,
	0x16, 0x53, 0xc3, 0xb1, 0xec, 0x17, 0xc1, 0xbc, 0x90, 0x1a, 0xee, 0x8a, 0x9d, 0x6e, 0x90, 0x85,
	0xf2, 0x8d, 0x42, 0x35, 0x99, 0xe7, 0xf8, 0x3c, 0xdd, 0x27, 0x4d, 0x30, 0x60, 0x3d, 0x7a, 0xed,
	0x20, 0x0e, 0x3a, 0x00, 0x84, 0xfe, 0x16, 0x90, 0xdb, 0x46, 0xb3, 0x3a, 0xbf, 0xfc, 0x42, 0xc9,
	0xc4, 0x3d, 0x12, 0x48, 0xb9, 0x09, 0x32, 0xfa, 0x79, 0x30, 0xb5, 0xc6, 0x18, 0xeb, 0xdd, 0x28,
	0x34, 0xfa, 0x59, 0x49, 0x75, 0xaf, 0x64, 0xea, 0x2e, 0xb0, 0xf5, 0xf5, 0x78, 0xc9, 0xa9, 0x54,
	0x9d, 0x5d, 0xd6, 0xb7, 0xea, 0x67, 0x8c, 0x60, 0xa4, 0x16, 0x50, 0xfe, 0x6a, 0x95, 0x1d, 0x1f,
	0x37, 0x66, 0xbc, 0xae, 0xd4, 0x9e, 0x38, 0x73, 0x4c, 0x24, 0x6a, 0x76, 0x32, 0x51, 0x77, 0x09,
	0xf1, 0x33, 0x81, 0x4b, 0x15, 0x56, 0x57, 0x03, 0xc7, 0x01, 0x97, 0xac, 0x57, 0xcb, 0xc7, 0x2a,
	0x56, 0x79, 0x2e, 0x62, 0x2b, 0x12, 0x5f, 0x51, 0x37, 0x01, 0x7e, 0x52, 0xa2, 0xf4, 0xdb, 0xa0,
	0x0c, 0xe4, 0xee, 0x0d, 0xaa, 0xa7, 0xb9, 0xf3, 0xe9, 0xf4, 0x52, 0x51, 0x8f, 0x46, 0xfe, 0x38,
	0xdd, 0xbc, 0xaf, 0xe8, 0x37, 0x15, 0x4b, 0x28, 0x98, 0x05, 0x10, 0xd2, 0x4f, 0xa6, 0xcb, 0x12,
	0x6a, 0xa5, 0x51, 0xf5, 0x59, 0xeb, 0x8b, 0x19, 0xb2, 0x56, 0x2b, 0xad, 0x1b, 0x91, 0x18, 0x0e,
	0xcb, 0x30, 0x66, 0x2a, 0xab, 0x62, 0x95, 0x62, 0xb6, 0xb0, 0x42, 0x96, 0x4a, 0x10, 0x12, 0xd6,
	0x26, 0x2b, 0xd2, 0x48, 0xc3, 0x64, 0x8e, 0xe3, 0x93, 0x6b, 0x33, 0x14, 0xa0, 0x9b, 0x0e, 0xef,
	0x7a, 0x18, 0xa7, 0x41, 0x85, 0x6f, 0x3b, 0x4e, 0xbd, 0x32, 0xf1, 0xf2, 0xb3, 0xec, 0xe0, 0x63,
	0x44, 0xbb, 0x89, 0x9b, 0x2c, 0x7b, 0x67, 0x05, 0xe3, 0x86, 0xe5, 0x83, 0xac, 0x27, 0xb4, 0x2f,
	0x92, 0x66, 0xef, 0xac, 0xd8, 0x33, 0x47, 0x00, 0xb9, 0x58, 0x42, 0x9e, 0xe9, 0x51, 0xaf, 0x1b,
	0x18, 0x0b, 0xe0, 0xca, 0xef, 0x0d, 0x42, 0xab, 0x42, 0xaa, 0x28, 0x42, 0xb1, 0x34, 0xa2, 0xd5,
	0xca, 0x52, 0x92, 0x6c, 0xfd, 0x3c, 0x4b, 0x6e, 0x4f, 0x8c, 0xeb, 0x22, 0x4f, 0x44, 0x02, 0xb9,
	0xa5, 0xef, 0x92, 0xd6, 0x18, 0x30, 0xc8, 0xa5, 0xbd, 0x64, 0x89, 0xca, 0xb8, 0xcc, 0x99, 0x4c,
	0xd8, 0x05, 0x4f, 0x07, 0x02, 0x9e, 0xcb, 0xe5, 0xe8, 0x5e, 0xe9, 0xf9, 0xa4, 0x74, 0x7c, 0x0a,
	0x7e, 0xdd, 0xe4, 0x85, 0xf3, 0xba, 0x3e, 0x56, 0x3a, 0x12, 0x0c, 0xae, 0x1c, 0x5b, 0x6d, 0x32,
	0x56, 0x5a, 0x06, 0x83, 0x24, 0x3c, 0x22, 0x5b, 0xd7, 0xf0, 0x2a, 0x55, 0x0f, 0xbb, 0x70, 0x63,
	0x22, 0x86, 0x9f, 0xd2, 0x1f, 0x93, 0x3b, 0xd7, 0x2c, 0x86, 0xe3, 0x6b, 0x56, 0xbd, 0x00, 0x9b,
	0x13, 0xcb, 0x23, 0xf0, 0xe8, 0x26, 0x74, 0x8f, 0xdc, 0xfd, 0xbb, 0x00, 0x78, 0x08, 0x9c, 0xb6,
	0xb7, 0xae, 0x8f, 0x00, 0x07, 0x78, 0xfb, 0xda, 0x03, 0xa8, 0xc2, 0xd5, 0xb6, 0xf1, 0x13, 0x78,
	0x38, 0xb1, 0xfe, 0x7d, 0xb4, 0xb7, 0x7e, 0x99, 0x23, 0xab, 0x13, 0x69, 0xa3, 0x3f, 0x55, 0x43,
	0xfa, 0x48, 0xa3, 0x84, 0x3b, 0x20, 0x14, 0xd3, 0x6d, 0x41, 0xd0, 0x5f, 0x3f, 0xd4, 0x76, 0xeb,
	0x08, 0xf4, 0xc7, 0x60, 0xa4, 0xe3, 0xe0, 0xf9, 0x7f, 0x08, 0x6c, 0x3f, 0x0b, 0xa6, 0x38, 0x7a,
	0xd5, 0x82, 0x50, 0x77, 0x3e, 0x0c, 0x20, 0xbf, 0x07, 0x24, 0xbc, 0xa6, 0x49, 0x20, 0x33, 0xe1,
	0x2e, 0x10, 0xff, 0x72, 0x8a, 0xc4, 0x27, 0x99, 0xf9, 0xa9, 0x64, 0xff, 0x6a, 0x19, 0xf5, 0xe6,
	0xe1, 0x5c, 0x0f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x44, 0xc3, 0xb9, 0x9d, 0x37, 0x10, 0x00,
	0x00,
}