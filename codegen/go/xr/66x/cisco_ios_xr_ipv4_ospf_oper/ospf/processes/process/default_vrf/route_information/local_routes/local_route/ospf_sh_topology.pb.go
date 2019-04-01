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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_route_information_local_routes_local_route

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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShNnhInfo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_nnh_info")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.route_information.local_routes.local_route.ospf_sh_topology")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_4169ff6e5af11206) }

var fileDescriptor_4169ff6e5af11206 = []byte{
	// 1302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6e, 0x1b, 0xc5,
	0x17, 0xd7, 0xa6, 0xf9, 0x37, 0xce, 0x38, 0x69, 0x92, 0x69, 0x9b, 0x6c, 0xfa, 0x6f, 0x69, 0x30,
	0x12, 0x58, 0x08, 0x59, 0x55, 0x9b, 0x42, 0xa1, 0x52, 0x51, 0x04, 0x89, 0xb0, 0x68, 0xab, 0x68,
	0x9d, 0x56, 0xe2, 0x6a, 0x34, 0xde, 0x9d, 0x8d, 0x47, 0xda, 0x9d, 0x19, 0xcd, 0x8c, 0x9d, 0xa4,
	0x08, 0x89, 0x07, 0x40, 0xe2, 0x1e, 0x21, 0x71, 0x01, 0x12, 0x3c, 0x06, 0x0f, 0xc0, 0xe3, 0x70,
	0xc5, 0x15, 0x9a, 0x33, 0xb3, 0xde, 0xb5, 0xdd, 0x6b, 0x7c, 0xe7, 0x3d, 0x5f, 0xf3, 0x3b, 0xdf,
	0xc7, 0x68, 0x57, 0x1a, 0x95, 0x13, 0x33, 0x22, 0x56, 0x2a, 0x59, 0xc8, 0xf3, 0xab, 0x9e, 0xd2,
	0xd2, 0x4a, 0xcc, 0x53, 0x6e, 0x52, 0x49, 0xb8, 0x34, 0xe4, 0x52, 0x13, 0xae, 0x26, 0x87, 0x04,
	0x24, 0xa5, 0x62, 0xba, 0xe7, 0x7e, 0x39, 0xb9, 0x94, 0x19, 0xc3, 0x4c, 0xf5, 0xab, 0x97, 0xb1,
	0x9c, 0x8e, 0x0b, 0x4b, 0x26, 0x3a, 0xef, 0x69, 0x39, 0xb6, 0x8c, 0x70, 0x91, 0x4b, 0x5d, 0x52,
	0xcb, 0xa5, 0xe8, 0x15, 0x32, 0xa5, 0x05, 0x01, 0xba, 0x69, 0x7e, 0x74, 0x2e, 0xd0, 0xed, 0x79,
	0x10, 0xe4, 0xeb, 0xe3, 0x6f, 0x06, 0xf8, 0x5d, 0xb4, 0x11, 0x4c, 0x13, 0x41, 0x4b, 0x16, 0x47,
	0x07, 0x51, 0x77, 0x3d, 0x69, 0x07, 0xda, 0x4b, 0x5a, 0x32, 0xbc, 0x8b, 0xae, 0x2b, 0xcd, 0x72,
	0x7e, 0x19, 0xaf, 0x00, 0x33, 0x7c, 0xe1, 0xf7, 0xd0, 0xa6, 0xff, 0x45, 0x0a, 0x26, 0xce, 0xed,
	0x28, 0xbe, 0x76, 0x10, 0x75, 0x37, 0x93, 0x0d, 0x4f, 0x7c, 0x0e, 0xb4, 0xce, 0x09, 0xda, 0x98,
	0x3e, 0xcc, 0xbd, 0x31, 0xc3, 0x52, 0x29, 0x32, 0x78, 0x69, 0x33, 0x09, 0x5f, 0xf8, 0x1d, 0x84,
	0x04, 0x15, 0x32, 0xf0, 0x56, 0x80, 0xd7, 0xa0, 0x74, 0xfe, 0x58, 0x43, 0xb8, 0xe1, 0x01, 0x49,
	0x65, 0x59, 0x4a, 0x81, 0x3b, 0x68, 0xd3, 0x47, 0x81, 0x6a, 0x46, 0x09, 0xaf, 0xac, 0xb6, 0x81,
	0x78, 0xa4, 0x19, 0xed, 0x67, 0xf8, 0x7d, 0xb4, 0xe5, 0x65, 0x2c, 0x23, 0x25, 0xb3, 0x9a, 0xa7,
	0xc1, 0xbe, 0x57, 0x3d, 0x63, 0x2f, 0x80, 0x88, 0x3f, 0x44, 0x3b, 0x5e, 0x4e, 0xf3, 0x21, 0x99,
	0x30, 0x6d, 0xb8, 0x14, 0xc1, 0x27, 0x6f, 0x20, 0xe1, 0xc3, 0xd7, 0x9e, 0x5c, 0xcb, 0x3a, 0x48,
	0x95, 0xec, 0xea, 0x41, 0xd4, 0x5d, 0x0d, 0xb2, 0x03, 0x95, 0x57, 0xb2, 0x87, 0x68, 0xd7, 0xcb,
	0xe6, 0x52, 0x5f, 0x50, 0x9d, 0x91, 0x8c, 0x1b, 0x4b, 0x45, 0xca, 0xe2, 0xff, 0x81, 0xf1, 0x5b,
	0xc0, 0x3d, 0xf1, 0xcc, 0x2f, 0x03, 0xcf, 0x25, 0x26, 0xbc, 0x20, 0xc7, 0x3a, 0x65, 0xf1, 0xf5,
	0x86, 0x63, 0x03, 0x20, 0xe1, 0x5f, 0xa3, 0x0a, 0xc5, 0x58, 0x65, 0xd4, 0x39, 0xc8, 0x4b, 0x16,
	0xaf, 0x1d, 0x44, 0xdd, 0xf6, 0xc3, 0x8b, 0xde, 0x7f, 0x56, 0x5c, 0xbd, 0x66, 0x82, 0x83, 0xfb,
	0xaf, 0x00, 0xd0, 0x99, 0xcb, 0xf8, 0x2f, 0x51, 0x15, 0xff, 0x9c, 0xf2, 0xc2, 0x63, 0x6c, 0x2d,
	0x17, 0xa3, 0x4f, 0xfc, 0x09, 0xe5, 0x05, 0x20, 0xfc, 0x08, 0xe1, 0x3a, 0x99, 0x4a, 0x73, 0xa9,
	0xb9, 0xbd, 0x8a, 0xd7, 0x21, 0xe0, 0xdb, 0x55, 0x36, 0x4f, 0x03, 0x1d, 0xf7, 0xd0, 0xcd, 0x50,
	0x72, 0x63, 0x2b, 0x09, 0xbb, 0x4c, 0x8b, 0x71, 0xc6, 0xb2, 0x18, 0x1d, 0x44, 0xdd, 0x56, 0xe2,
	0xf3, 0x71, 0x34, 0xb6, 0xf2, 0x38, 0x30, 0xf0, 0xe7, 0xe8, 0x6e, 0xb0, 0xae, 0x2d, 0x23, 0xa1,
	0x63, 0x34, 0x3b, 0xe7, 0xc6, 0x32, 0xcd, 0xb2, 0xb8, 0x0d, 0x8a, 0xfb, 0xfe, 0x1d, 0x6d, 0xd9,
	0x29, 0x48, 0x24, 0x53, 0x01, 0xfc, 0x29, 0xda, 0x6f, 0x18, 0x10, 0x43, 0xdd, 0xd4, 0xde, 0x00,
	0x94, 0xbb, 0x53, 0xed, 0x97, 0x43, 0xdd, 0x50, 0xed, 0x22, 0x8f, 0x9f, 0x14, 0x74, 0xc8, 0x0a,
	0x62, 0xaf, 0x14, 0x8b, 0x37, 0x41, 0xe3, 0x06, 0xd0, 0x9f, 0x3b, 0xf2, 0xd9, 0x95, 0x62, 0xf8,
	0x3e, 0x6a, 0x37, 0x24, 0xe3, 0x1b, 0xbe, 0x01, 0x6b, 0xa1, 0xda, 0x94, 0x71, 0x51, 0xf2, 0x52,
	0x5b, 0x0d, 0x53, 0x03, 0xa3, 0x72, 0x90, 0xec, 0xfc, 0x18, 0xa1, 0x1b, 0x55, 0xb8, 0x35, 0x53,
	0x84, 0x15, 0xd0, 0x2e, 0x4c, 0x51, 0xae, 0x09, 0x2b, 0x58, 0xc9, 0x84, 0xad, 0x5a, 0x75, 0x3d,
	0xd9, 0xf2, 0x8c, 0x63, 0x4f, 0xef, 0x67, 0x50, 0xf8, 0x5e, 0xd6, 0x3f, 0xb2, 0x12, 0x0a, 0x1f,
	0x68, 0x1e, 0x8b, 0x4b, 0xc1, 0xac, 0x39, 0xf0, 0xcc, 0xf7, 0xea, 0xce, 0x8c, 0x41, 0xe7, 0x5c,
	0xe7, 0xef, 0x95, 0x7a, 0xfc, 0x19, 0x4d, 0xc6, 0x85, 0x94, 0x8a, 0x28, 0x6a, 0x47, 0xf8, 0xb7,
	0x08, 0x6d, 0x97, 0x3c, 0xd5, 0x12, 0x48, 0x5e, 0x33, 0x8e, 0x0e, 0xae, 0x75, 0xdb, 0x0f, 0xaf,
	0x96, 0x50, 0x9d, 0x3e, 0x5c, 0xc9, 0xd6, 0x14, 0x52, 0x02, 0x88, 0xf0, 0x53, 0x74, 0x67, 0x1e,
	0x25, 0x29, 0xb8, 0xb1, 0xc4, 0xf0, 0x37, 0x2c, 0x44, 0x68, 0x6f, 0x4e, 0xe9, 0x39, 0x37, 0x76,
	0xc0, 0xdf, 0x30, 0x7c, 0x8c, 0xee, 0xd7, 0xca, 0x76, 0x2c, 0x04, 0x2b, 0x08, 0x17, 0x96, 0xe9,
	0x9c, 0xa6, 0xcc, 0x4f, 0xfd, 0x6b, 0x90, 0x8a, 0xbb, 0x53, 0xb1, 0x33, 0x90, 0xea, 0x57, 0x42,
	0xb0, 0x06, 0x1e, 0xa0, 0x5b, 0xb5, 0x19, 0xe3, 0x26, 0xa6, 0x75, 0x0d, 0x03, 0x53, 0xaf, 0x95,
	0xe0, 0x29, 0x6f, 0x00, 0xac, 0x81, 0xca, 0x3b, 0x7f, 0x45, 0x68, 0xbb, 0xf2, 0x4c, 0x88, 0x11,
	0x84, 0x02, 0x3f, 0x43, 0x77, 0x05, 0xe3, 0xe7, 0xa3, 0xa1, 0xd4, 0x44, 0xb0, 0x4b, 0x4b, 0x46,
	0x52, 0x91, 0x09, 0x2d, 0xc6, 0xcc, 0x27, 0xd1, 0x57, 0x45, 0x5c, 0xc9, 0xbc, 0x64, 0x97, 0xf6,
	0x2b, 0xa9, 0x5e, 0x3b, 0x01, 0x28, 0xd4, 0x27, 0x68, 0x7f, 0x51, 0x9f, 0x2b, 0x42, 0xb3, 0x4c,
	0x87, 0x05, 0x75, 0x7b, 0x4e, 0xb9, 0xaf, 0x8e, 0xb2, 0x4c, 0xbf, 0xfd, 0x65, 0x2e, 0x6c, 0x4e,
	0xb8, 0xc8, 0xd8, 0x65, 0x28, 0x9f, 0xf9, 0x97, 0xfb, 0xc2, 0xe6, 0x7d, 0xc7, 0xef, 0xfc, 0xd3,
	0xaa, 0xdd, 0xb1, 0x55, 0x01, 0x3d, 0x40, 0xb7, 0xaa, 0x3c, 0xcf, 0x44, 0xd4, 0xbb, 0xe1, 0xe7,
	0xca, 0x6c, 0x1c, 0x9f, 0xa2, 0x3b, 0xf3, 0x1a, 0x46, 0x94, 0x2a, 0x80, 0x08, 0xb9, 0x9c, 0xd5,
	0x1b, 0x88, 0x52, 0x01, 0x06, 0xfc, 0xa8, 0xda, 0x25, 0x53, 0x07, 0x9c, 0xdf, 0xcc, 0x98, 0x90,
	0x42, 0x3f, 0x9a, 0x02, 0xf4, 0x23, 0xcf, 0x5a, 0x58, 0x25, 0xab, 0x7e, 0xc7, 0x37, 0x57, 0x49,
	0xdd, 0xfe, 0x86, 0xf2, 0x0c, 0x16, 0xd3, 0x7a, 0xd5, 0xfe, 0x8e, 0x82, 0x3f, 0xab, 0x50, 0x3b,
	0xaf, 0x09, 0x37, 0xa4, 0x4c, 0xa9, 0xb1, 0xce, 0x07, 0x9a, 0x5a, 0x58, 0x4e, 0xad, 0x30, 0x85,
	0x4e, 0xa9, 0x1d, 0xf5, 0xcd, 0x0b, 0xc7, 0xee, 0x03, 0x17, 0x1f, 0xa2, 0xbd, 0x59, 0xdd, 0x71,
	0x5a, 0xfa, 0xf0, 0xc1, 0xb2, 0x6a, 0x05, 0xd4, 0x5e, 0xf1, 0x55, 0x5a, 0x2a, 0xf7, 0xab, 0x46,
	0x1d, 0x76, 0x76, 0xab, 0xb1, 0x00, 0xc3, 0xc6, 0xde, 0x47, 0xad, 0xc2, 0x50, 0x5f, 0x37, 0x7e,
	0x5c, 0xaf, 0x15, 0x86, 0x42, 0x99, 0xec, 0xa1, 0xb5, 0xea, 0x24, 0x40, 0xfe, 0xd0, 0xa0, 0xfe,
	0x1a, 0xb8, 0x8f, 0xda, 0xc0, 0xf0, 0x5d, 0x19, 0xa6, 0x2f, 0x72, 0xa4, 0x13, 0xa0, 0xe0, 0x3f,
	0x23, 0xf4, 0x7f, 0xa3, 0x49, 0x5d, 0xeb, 0x74, 0x22, 0x79, 0xe6, 0x76, 0xb2, 0x87, 0xbc, 0x01,
	0xbb, 0xeb, 0xfb, 0x68, 0x09, 0xe3, 0x61, 0x66, 0x76, 0x25, 0xb1, 0xd1, 0x2f, 0x2a, 0x90, 0x47,
	0x15, 0x46, 0x08, 0xdd, 0xf4, 0x2a, 0xf2, 0x01, 0xcf, 0xc2, 0xcc, 0x6f, 0xd7, 0x61, 0x76, 0x5b,
	0x65, 0x2e, 0xa1, 0xda, 0x4c, 0x14, 0x09, 0xdf, 0x30, 0xff, 0x5b, 0xc9, 0xed, 0x46, 0x5e, 0x12,
	0x33, 0x51, 0x67, 0xde, 0xfc, 0x63, 0x14, 0xcf, 0xaa, 0x1a, 0x3d, 0x55, 0xdc, 0x5a, 0x48, 0xe8,
	0x40, 0x07, 0xb5, 0x85, 0x17, 0x8d, 0x86, 0xf5, 0xe9, 0x15, 0xb7, 0x17, 0x5e, 0x1c, 0x68, 0xb7,
	0x43, 0x41, 0xf5, 0x19, 0xba, 0xf7, 0xb6, 0x17, 0x61, 0x15, 0x81, 0xf6, 0x0e, 0x68, 0xef, 0xcd,
	0x3d, 0xeb, 0x76, 0x12, 0xe8, 0xff, 0x1e, 0xa1, 0x9d, 0x85, 0xde, 0x8f, 0x31, 0xcc, 0xf9, 0x6f,
	0x97, 0x90, 0xc8, 0x6a, 0x1a, 0x26, 0x5b, 0x73, 0xd3, 0xc6, 0xdd, 0xc7, 0x17, 0x8e, 0x64, 0xe3,
	0x9b, 0xbe, 0x6c, 0xfd, 0x57, 0xe7, 0x87, 0xd5, 0x99, 0xe1, 0x03, 0x17, 0x7c, 0xdd, 0x22, 0xe1,
	0x3e, 0x7f, 0xd8, 0x68, 0x6c, 0x7f, 0x46, 0xd4, 0xd7, 0xca, 0xec, 0xa9, 0xfe, 0x28, 0xac, 0xca,
	0x5a, 0xd2, 0xdf, 0xeb, 0x0b, 0x5d, 0x77, 0xb8, 0xd8, 0x75, 0xf7, 0x10, 0x0a, 0xf7, 0xb4, 0xeb,
	0xbb, 0xc7, 0xf0, 0xe6, 0xba, 0x3f, 0xa5, 0x5d, 0xe7, 0x7d, 0x50, 0x9d, 0x7b, 0xa9, 0x14, 0x82,
	0xa5, 0x96, 0x65, 0xf1, 0xc7, 0x90, 0x1d, 0x7f, 0x27, 0x7c, 0x51, 0x51, 0xf1, 0xcf, 0x51, 0x65,
	0xc8, 0x85, 0x22, 0xfe, 0x04, 0xfa, 0xea, 0xbb, 0x65, 0xdc, 0x84, 0xd3, 0xff, 0x13, 0xc1, 0x8f,
	0xbe, 0x5b, 0x54, 0x3f, 0x4d, 0xe1, 0x41, 0x85, 0x3d, 0x59, 0x5a, 0xb1, 0x54, 0xbb, 0x26, 0x80,
	0x73, 0x05, 0x3d, 0xbc, 0x0e, 0xff, 0x20, 0x1f, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x8e,
	0xe8, 0x4a, 0x5b, 0x0e, 0x00, 0x00,
}