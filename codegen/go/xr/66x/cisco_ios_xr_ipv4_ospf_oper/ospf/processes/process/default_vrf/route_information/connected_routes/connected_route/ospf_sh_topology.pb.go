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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_connected_routes_connected_route

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
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
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
	RouteLabelType            uint32      `protobuf:"varint,13,opt,name=route_label_type,json=routeLabelType,proto3" json:"route_label_type,omitempty"`
	RouteLabel                uint32      `protobuf:"varint,14,opt,name=route_label,json=routeLabel,proto3" json:"route_label,omitempty"`
	RouteSspfLabel            uint32      `protobuf:"varint,15,opt,name=route_sspf_label,json=routeSspfLabel,proto3" json:"route_sspf_label,omitempty"`
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

func (m *OspfShTopCommon) GetRouteLabelType() uint32 {
	if m != nil {
		return m.RouteLabelType
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteLabel() uint32 {
	if m != nil {
		return m.RouteLabel
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSspfLabel() uint32 {
	if m != nil {
		return m.RouteSspfLabel
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

type OspfShNnhInfo struct {
	NeighborNextHopValueType string   `protobuf:"bytes,1,opt,name=neighbor_next_hop_value_type,json=neighborNextHopValueType,proto3" json:"neighbor_next_hop_value_type,omitempty"`
	NeighborNextHopIpAddr    string   `protobuf:"bytes,2,opt,name=neighbor_next_hop_ip_addr,json=neighborNextHopIpAddr,proto3" json:"neighbor_next_hop_ip_addr,omitempty"`
	NeighborNextHopIntfIndex uint32   `protobuf:"varint,3,opt,name=neighbor_next_hop_intf_index,json=neighborNextHopIntfIndex,proto3" json:"neighbor_next_hop_intf_index,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *OspfShNnhInfo) Reset()         { *m = OspfShNnhInfo{} }
func (m *OspfShNnhInfo) String() string { return proto.CompactTextString(m) }
func (*OspfShNnhInfo) ProtoMessage()    {}
func (*OspfShNnhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4169ff6e5af11206, []int{5}
}

func (m *OspfShNnhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNnhInfo.Unmarshal(m, b)
}
func (m *OspfShNnhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNnhInfo.Marshal(b, m, deterministic)
}
func (m *OspfShNnhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNnhInfo.Merge(m, src)
}
func (m *OspfShNnhInfo) XXX_Size() int {
	return xxx_messageInfo_OspfShNnhInfo.Size(m)
}
func (m *OspfShNnhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNnhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNnhInfo proto.InternalMessageInfo

func (m *OspfShNnhInfo) GetNeighborNextHopValueType() string {
	if m != nil {
		return m.NeighborNextHopValueType
	}
	return ""
}

func (m *OspfShNnhInfo) GetNeighborNextHopIpAddr() string {
	if m != nil {
		return m.NeighborNextHopIpAddr
	}
	return ""
}

func (m *OspfShNnhInfo) GetNeighborNextHopIntfIndex() uint32 {
	if m != nil {
		return m.NeighborNextHopIntfIndex
	}
	return 0
}

type OspfShTopPath struct {
	RouteInterfaceName       string             `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName,proto3" json:"route_interface_name,omitempty"`
	RouteInterfaceSnmpIndex  uint32             `protobuf:"varint,2,opt,name=route_interface_snmp_index,json=routeInterfaceSnmpIndex,proto3" json:"route_interface_snmp_index,omitempty"`
	RouteNextHopAddress      string             `protobuf:"bytes,3,opt,name=route_next_hop_address,json=routeNextHopAddress,proto3" json:"route_next_hop_address,omitempty"`
	RouteSource              string             `protobuf:"bytes,4,opt,name=route_source,json=routeSource,proto3" json:"route_source,omitempty"`
	RouteLsaid               string             `protobuf:"bytes,5,opt,name=route_lsaid,json=routeLsaid,proto3" json:"route_lsaid,omitempty"`
	RoutePathIsMcastIntact   bool               `protobuf:"varint,6,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact,proto3" json:"route_path_is_mcast_intact,omitempty"`
	RoutePathIsUcmpPath      bool               `protobuf:"varint,7,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath,proto3" json:"route_path_is_ucmp_path,omitempty"`
	RouteMetric              uint32             `protobuf:"varint,8,opt,name=route_metric,json=routeMetric,proto3" json:"route_metric,omitempty"`
	LsaType                  uint32             `protobuf:"varint,9,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	AreaId                   uint32             `protobuf:"varint,10,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaFormat               bool               `protobuf:"varint,11,opt,name=area_format,json=areaFormat,proto3" json:"area_format,omitempty"`
	SrMicroloopAvoidancePath *OspfShSrUloopPath `protobuf:"bytes,12,opt,name=sr_microloop_avoidance_path,json=srMicroloopAvoidancePath,proto3" json:"sr_microloop_avoidance_path,omitempty"`
	RoutePathId              uint32             `protobuf:"varint,13,opt,name=route_path_id,json=routePathId,proto3" json:"route_path_id,omitempty"`
	RoutePathIsRsvpTePath    bool               `protobuf:"varint,14,opt,name=route_path_is_rsvp_te_path,json=routePathIsRsvpTePath,proto3" json:"route_path_is_rsvp_te_path,omitempty"`
	RoutePathIsSrTePath      bool               `protobuf:"varint,15,opt,name=route_path_is_sr_te_path,json=routePathIsSrTePath,proto3" json:"route_path_is_sr_te_path,omitempty"`
	RoutePathIsSrExclPath    bool               `protobuf:"varint,16,opt,name=route_path_is_sr_excl_path,json=routePathIsSrExclPath,proto3" json:"route_path_is_sr_excl_path,omitempty"`
	RoutePathIsSrTeSspfPath  bool               `protobuf:"varint,17,opt,name=route_path_is_sr_te_sspf_path,json=routePathIsSrTeSspfPath,proto3" json:"route_path_is_sr_te_sspf_path,omitempty"`
	NeighborNextHop          []*OspfShNnhInfo   `protobuf:"bytes,18,rep,name=neighbor_next_hop,json=neighborNextHop,proto3" json:"neighbor_next_hop,omitempty"`
	Weight                   uint32             `protobuf:"varint,19,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}           `json:"-"`
	XXX_unrecognized         []byte             `json:"-"`
	XXX_sizecache            int32              `json:"-"`
}

func (m *OspfShTopPath) Reset()         { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()    {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_4169ff6e5af11206, []int{6}
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

func (m *OspfShTopPath) GetRouteInterfaceSnmpIndex() uint32 {
	if m != nil {
		return m.RouteInterfaceSnmpIndex
	}
	return 0
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

func (m *OspfShTopPath) GetRoutePathId() uint32 {
	if m != nil {
		return m.RoutePathId
	}
	return 0
}

func (m *OspfShTopPath) GetRoutePathIsRsvpTePath() bool {
	if m != nil {
		return m.RoutePathIsRsvpTePath
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsSrTePath() bool {
	if m != nil {
		return m.RoutePathIsSrTePath
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsSrExclPath() bool {
	if m != nil {
		return m.RoutePathIsSrExclPath
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsSrTeSspfPath() bool {
	if m != nil {
		return m.RoutePathIsSrTeSspfPath
	}
	return false
}

func (m *OspfShTopPath) GetNeighborNextHop() []*OspfShNnhInfo {
	if m != nil {
		return m.NeighborNextHop
	}
	return nil
}

func (m *OspfShTopPath) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
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
	return fileDescriptor_4169ff6e5af11206, []int{7}
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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShNnhInfo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_nnh_info")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.connected_routes.connected_route.ospf_sh_topology")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_4169ff6e5af11206) }

var fileDescriptor_4169ff6e5af11206 = []byte{
	// 1306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4d, 0x6f, 0x1b, 0xc5,
	0x1b, 0xd7, 0xa6, 0xf9, 0x37, 0xce, 0x38, 0x69, 0x92, 0x69, 0x9b, 0x6c, 0xfa, 0x6f, 0x69, 0x30,
	0x12, 0x58, 0x08, 0x59, 0x55, 0x9b, 0x42, 0xa1, 0x52, 0x51, 0x04, 0x89, 0xb0, 0x68, 0xab, 0x6a,
	0x9d, 0x56, 0xe2, 0x34, 0x1a, 0xef, 0xce, 0xc6, 0x23, 0xed, 0xce, 0x0c, 0x33, 0x63, 0xd7, 0xe9,
	0x05, 0x21, 0x84, 0x84, 0xc4, 0x81, 0x6f, 0xc0, 0x05, 0x21, 0x0e, 0x08, 0x89, 0x23, 0x1f, 0x81,
	0x03, 0x5f, 0x09, 0x09, 0xcd, 0x33, 0xb3, 0xde, 0xb5, 0xdd, 0x33, 0xe1, 0xe6, 0x7d, 0xde, 0xe6,
	0x79, 0xff, 0x3d, 0x46, 0xbb, 0xd2, 0xa8, 0x9c, 0x98, 0x11, 0xb1, 0x52, 0xc9, 0x42, 0x9e, 0x9d,
	0xf7, 0x94, 0x96, 0x56, 0xe2, 0x2f, 0x53, 0x6e, 0x52, 0x49, 0xb8, 0x34, 0x64, 0xaa, 0x09, 0x57,
	0x93, 0x43, 0x02, 0x92, 0x52, 0x31, 0xdd, 0x73, 0xbf, 0x9c, 0x5c, 0xca, 0x8c, 0x61, 0xa6, 0xfa,
	0xd5, 0xcb, 0x58, 0x4e, 0xc7, 0x85, 0x25, 0x13, 0x9d, 0xf7, 0xb4, 0x1c, 0x5b, 0x46, 0xb8, 0xc8,
	0xa5, 0x2e, 0xa9, 0xe5, 0x52, 0xf4, 0x52, 0x29, 0x04, 0x4b, 0x2d, 0xcb, 0x08, 0xf0, 0xcc, 0x22,
	0xa1, 0xf3, 0x12, 0x5d, 0x5f, 0x74, 0x86, 0x7c, 0x7e, 0xfc, 0xc5, 0x00, 0xbf, 0x89, 0x36, 0xc2,
	0x13, 0x44, 0xd0, 0x92, 0xc5, 0xd1, 0x41, 0xd4, 0x5d, 0x4f, 0xda, 0x81, 0xf6, 0x94, 0x96, 0x0c,
	0xef, 0xa2, 0xcb, 0x4a, 0xb3, 0x9c, 0x4f, 0xe3, 0x15, 0x60, 0x86, 0x2f, 0xfc, 0x16, 0xda, 0xf4,
	0xbf, 0x48, 0xc1, 0xc4, 0x99, 0x1d, 0xc5, 0x97, 0x0e, 0xa2, 0xee, 0x66, 0xb2, 0xe1, 0x89, 0x8f,
	0x81, 0xd6, 0x39, 0x41, 0x1b, 0xb3, 0x87, 0xb9, 0x37, 0x66, 0x58, 0x2a, 0x45, 0x06, 0x2f, 0x6d,
	0x26, 0xe1, 0x0b, 0xbf, 0x81, 0x90, 0xa0, 0x42, 0x06, 0xde, 0x0a, 0xf0, 0x1a, 0x94, 0xce, 0x1f,
	0x6b, 0x08, 0x37, 0x22, 0x20, 0xa9, 0x2c, 0x4b, 0x29, 0x70, 0x07, 0x6d, 0xfa, 0x6c, 0x50, 0xcd,
	0x28, 0xe1, 0x95, 0xd5, 0x36, 0x10, 0x8f, 0x34, 0xa3, 0xfd, 0x0c, 0xbf, 0x8d, 0xb6, 0xbc, 0x8c,
	0x65, 0xa4, 0x64, 0x56, 0xf3, 0x34, 0xd8, 0xf7, 0xaa, 0xa7, 0xec, 0x09, 0x10, 0xf1, 0xbb, 0x68,
	0xc7, 0xcb, 0x69, 0x3e, 0x24, 0x13, 0xa6, 0x0d, 0x97, 0x22, 0xc4, 0xe4, 0x0d, 0x24, 0x7c, 0xf8,
	0xc2, 0x93, 0x6b, 0x59, 0xe7, 0x52, 0x25, 0xbb, 0x7a, 0x10, 0x75, 0x57, 0x83, 0xec, 0x40, 0xe5,
	0x95, 0xec, 0x21, 0xda, 0xf5, 0xb2, 0xb9, 0xd4, 0x2f, 0xa9, 0xce, 0x48, 0xc6, 0x8d, 0xa5, 0x22,
	0x65, 0xf1, 0xff, 0xc0, 0xf8, 0x35, 0xe0, 0x9e, 0x78, 0xe6, 0xa7, 0x81, 0xe7, 0x0a, 0x13, 0x5e,
	0x90, 0x63, 0x9d, 0xb2, 0xf8, 0x72, 0x23, 0xb0, 0x01, 0x90, 0xf0, 0xaf, 0x51, 0xe5, 0xc5, 0x58,
	0x65, 0xd4, 0x05, 0xc8, 0x4b, 0x16, 0xaf, 0x1d, 0x44, 0xdd, 0xf6, 0xdd, 0xaf, 0x7a, 0xff, 0x7a,
	0x93, 0xf5, 0x9a, 0x85, 0x0e, 0x69, 0x78, 0x0e, 0x8e, 0x9d, 0xba, 0xca, 0xff, 0x12, 0x55, 0x75,
	0xc8, 0x29, 0x2f, 0xbc, 0xaf, 0xad, 0xff, 0x86, 0xaf, 0xbe, 0x11, 0x4e, 0x28, 0x2f, 0xc0, 0xd3,
	0xf7, 0x10, 0xae, 0x8b, 0xab, 0x34, 0x97, 0x9a, 0xdb, 0xf3, 0x78, 0x1d, 0x0a, 0xb0, 0x5d, 0x55,
	0xf7, 0x59, 0xa0, 0xe3, 0x1e, 0xba, 0x1a, 0x5a, 0x70, 0x6c, 0x25, 0x61, 0xd3, 0xb4, 0x18, 0x67,
	0x2c, 0x8b, 0xd1, 0x41, 0xd4, 0x6d, 0x25, 0xbe, 0x3e, 0x47, 0x63, 0x2b, 0x8f, 0x03, 0x03, 0x7f,
	0x8c, 0x6e, 0x06, 0xeb, 0xda, 0x32, 0x12, 0x26, 0x48, 0xb3, 0x33, 0x6e, 0x2c, 0xd3, 0x2c, 0x8b,
	0xdb, 0xa0, 0xb8, 0xef, 0xdf, 0xd1, 0x96, 0x3d, 0x03, 0x89, 0x64, 0x26, 0x80, 0x3f, 0x44, 0xfb,
	0x0d, 0x03, 0x62, 0xa8, 0x9b, 0xda, 0x1b, 0xe0, 0xe5, 0xee, 0x4c, 0xfb, 0xe9, 0x50, 0x37, 0x54,
	0xbb, 0xc8, 0xfb, 0x4f, 0x0a, 0x3a, 0x64, 0x05, 0xb1, 0xe7, 0x8a, 0xc5, 0x9b, 0xa0, 0x71, 0x05,
	0xe8, 0x8f, 0x1d, 0xf9, 0xf4, 0x5c, 0x31, 0x7c, 0x1b, 0xb5, 0x1b, 0x92, 0xf1, 0x15, 0x3f, 0x90,
	0xb5, 0x50, 0x6d, 0xca, 0xb8, 0x2c, 0x79, 0xa9, 0xad, 0x86, 0xa9, 0x81, 0x51, 0x39, 0x48, 0x76,
	0x7e, 0x88, 0xd0, 0x95, 0x2a, 0xdd, 0x9a, 0x29, 0xc2, 0x0a, 0x18, 0x1f, 0xa6, 0x28, 0xd7, 0x84,
	0x15, 0xac, 0x64, 0xc2, 0x56, 0xa3, 0xbb, 0x9e, 0x6c, 0x79, 0xc6, 0xb1, 0xa7, 0xf7, 0x33, 0x18,
	0x04, 0x2f, 0xeb, 0x1f, 0x59, 0x09, 0x83, 0x00, 0x34, 0xef, 0x8b, 0x2b, 0xc1, 0xbc, 0x39, 0x88,
	0xcc, 0xcf, 0xee, 0xce, 0x9c, 0x41, 0x17, 0x5c, 0xe7, 0xef, 0x95, 0x7a, 0x1d, 0x1a, 0x4d, 0xc6,
	0x85, 0x94, 0x8a, 0x28, 0x6a, 0x47, 0xf8, 0xb7, 0x08, 0x6d, 0x97, 0x3c, 0xd5, 0x12, 0x48, 0x5e,
	0x33, 0x8e, 0x0e, 0x2e, 0x75, 0xdb, 0x77, 0xbf, 0x8e, 0x2e, 0xb0, 0x4d, 0x7d, 0xde, 0x92, 0xad,
	0x99, 0x6f, 0x09, 0xb8, 0x86, 0x1f, 0xa2, 0x1b, 0x8b, 0xee, 0x92, 0x82, 0x1b, 0x4b, 0x0c, 0x7f,
	0xc5, 0x42, 0xaa, 0xf6, 0x16, 0x94, 0x1e, 0x73, 0x63, 0x07, 0xfc, 0x15, 0xc3, 0xc7, 0xe8, 0x76,
	0xad, 0x6c, 0xc7, 0x42, 0xb0, 0x82, 0x70, 0x61, 0x99, 0xce, 0x69, 0xca, 0x3c, 0x1c, 0x5c, 0x82,
	0x9a, 0xdc, 0x9c, 0x89, 0x9d, 0x82, 0x54, 0xbf, 0x12, 0x02, 0x7c, 0xb8, 0x83, 0xae, 0xd5, 0x66,
	0x8c, 0x5b, 0xa5, 0xd6, 0x4d, 0x0e, 0xac, 0xc3, 0x56, 0x82, 0x67, 0xbc, 0x01, 0xb0, 0x06, 0x2a,
	0xef, 0xfc, 0x15, 0xa1, 0xed, 0x2a, 0x32, 0x21, 0x46, 0x90, 0x12, 0xfc, 0x08, 0xdd, 0x14, 0x8c,
	0x9f, 0x8d, 0x86, 0x52, 0x13, 0xc1, 0xa6, 0x96, 0x8c, 0xa4, 0x22, 0x13, 0x5a, 0x8c, 0x99, 0xaf,
	0xa6, 0x6f, 0x8f, 0xb8, 0x92, 0x79, 0xca, 0xa6, 0xf6, 0x33, 0xa9, 0x5e, 0x38, 0x01, 0xe8, 0xd8,
	0x07, 0x68, 0x7f, 0x59, 0x9f, 0x2b, 0x42, 0xb3, 0x4c, 0x07, 0xe4, 0xba, 0xbe, 0xa0, 0xdc, 0x57,
	0x47, 0x59, 0xa6, 0x5f, 0xff, 0x32, 0x17, 0x36, 0x27, 0x5c, 0x64, 0x6c, 0x1a, 0xfa, 0x68, 0xf1,
	0xe5, 0xbe, 0xb0, 0x79, 0xdf, 0xf1, 0x3b, 0xdf, 0xaf, 0xd7, 0xe1, 0xd8, 0xaa, 0x93, 0xee, 0xa0,
	0x6b, 0x55, 0xbd, 0xe7, 0x32, 0xea, 0xc3, 0xf0, 0x0b, 0x66, 0x3e, 0x8f, 0x0f, 0xd1, 0x8d, 0x45,
	0x0d, 0x23, 0x4a, 0x15, 0x9c, 0x08, 0xb5, 0x9c, 0xd7, 0x1b, 0x88, 0x52, 0x81, 0x0f, 0xf8, 0x5e,
	0x05, 0x32, 0xb3, 0x00, 0x5c, 0xdc, 0xcc, 0x98, 0x50, 0x42, 0xbf, 0xa3, 0x82, 0xeb, 0x47, 0x9e,
	0xb5, 0x84, 0x31, 0xab, 0x1e, 0xfc, 0x9b, 0x18, 0x53, 0xef, 0x01, 0x43, 0x79, 0x06, 0x88, 0xb5,
	0x5e, 0xed, 0x01, 0x47, 0xc1, 0x1f, 0x55, 0x5e, 0xbb, 0xa8, 0x09, 0x37, 0xa4, 0x4c, 0xa9, 0xb1,
	0x2e, 0x06, 0x9a, 0x5a, 0x40, 0xad, 0x56, 0x58, 0x47, 0xcf, 0xa8, 0x1d, 0xf5, 0xcd, 0x13, 0xc7,
	0xee, 0x03, 0x17, 0x1f, 0xa2, 0xbd, 0x79, 0xdd, 0x71, 0x5a, 0xfa, 0xf4, 0x01, 0x8a, 0xb5, 0x82,
	0xd7, 0x5e, 0xf1, 0x79, 0x5a, 0x2a, 0xf7, 0xab, 0xf6, 0x3a, 0x80, 0x79, 0xab, 0x81, 0x8c, 0x01,
	0xca, 0xf7, 0x51, 0xab, 0x30, 0xd4, 0xf7, 0x8d, 0xdf, 0xdb, 0x6b, 0x85, 0xa1, 0xd0, 0x26, 0x7b,
	0x68, 0xad, 0xba, 0x15, 0x90, 0xbf, 0x40, 0xa8, 0x3f, 0x13, 0x6e, 0xa3, 0x36, 0x30, 0xfc, 0x74,
	0x86, 0x35, 0x8c, 0x1c, 0xe9, 0x04, 0x28, 0xf8, 0xcf, 0x08, 0xfd, 0xdf, 0x68, 0x52, 0xf7, 0x3a,
	0x9d, 0x48, 0x9e, 0x39, 0xb0, 0xf6, 0x2e, 0x6f, 0x00, 0x98, 0x7d, 0x77, 0x91, 0x6b, 0x62, 0x6e,
	0x99, 0x25, 0xb1, 0xd1, 0x4f, 0x2a, 0x67, 0x8f, 0x2a, 0x5f, 0x21, 0x85, 0xb3, 0xb3, 0xc9, 0x27,
	0x3e, 0x0b, 0x20, 0xd0, 0xae, 0xd3, 0xed, 0x60, 0x66, 0xa1, 0xb0, 0xda, 0x4c, 0x14, 0x09, 0xdf,
	0x00, 0x08, 0xad, 0xe4, 0x7a, 0xa3, 0x3e, 0x89, 0x99, 0xa8, 0x53, 0x6f, 0xfe, 0x3e, 0x8a, 0xe7,
	0x55, 0x8d, 0x9e, 0x29, 0x6e, 0x2d, 0x15, 0x76, 0xa0, 0x83, 0xda, 0xd2, 0x8b, 0x46, 0x03, 0x9e,
	0x7a, 0xc5, 0xed, 0xa5, 0x17, 0x07, 0xda, 0x81, 0x2a, 0xa8, 0x3e, 0x42, 0xb7, 0x5e, 0xf7, 0x22,
	0x60, 0x13, 0x68, 0xef, 0x80, 0xf6, 0xde, 0xc2, 0xb3, 0x0e, 0xa4, 0x40, 0xff, 0xf7, 0x08, 0xed,
	0x2c, 0xed, 0x80, 0x18, 0xc3, 0xe2, 0xff, 0xe6, 0x22, 0x2b, 0x5a, 0xad, 0xc7, 0x64, 0x6b, 0x61,
	0xfd, 0xb8, 0x4b, 0xfa, 0xa5, 0x23, 0xd9, 0xf8, 0xaa, 0xef, 0x63, 0xff, 0xd5, 0xf9, 0x71, 0x75,
	0x6e, 0x1b, 0xc1, 0xad, 0x5f, 0xcf, 0x4c, 0xb8, 0xe4, 0xef, 0x36, 0x26, 0xdd, 0x1f, 0x18, 0xf5,
	0x1d, 0x33, 0x7f, 0xd4, 0xdf, 0x0b, 0x20, 0x5a, 0x4b, 0xfa, 0xcb, 0x7e, 0x69, 0x0c, 0x0f, 0x97,
	0xc7, 0xf0, 0x16, 0x42, 0xe1, 0xf2, 0x76, 0x83, 0x78, 0x1f, 0xde, 0x5c, 0xf7, 0x47, 0xb7, 0x1b,
	0xc5, 0x77, 0xaa, 0x83, 0x70, 0x16, 0x7c, 0xfc, 0x3e, 0x94, 0xc9, 0x5f, 0x10, 0x9f, 0x54, 0x54,
	0xfc, 0x73, 0x54, 0x19, 0x72, 0xa9, 0x88, 0x3f, 0x80, 0x41, 0xfb, 0xf6, 0x22, 0xcb, 0x52, 0xff,
	0x05, 0x09, 0x01, 0xf5, 0x1d, 0x84, 0xfd, 0x34, 0xf3, 0x13, 0x7a, 0xee, 0xc1, 0xc5, 0xb7, 0x4f,
	0x05, 0x47, 0xc1, 0x4b, 0xd7, 0xeb, 0xc3, 0xcb, 0xf0, 0x2f, 0xf4, 0xde, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc1, 0xe7, 0x37, 0xb6, 0x9f, 0x0e, 0x00, 0x00,
}
