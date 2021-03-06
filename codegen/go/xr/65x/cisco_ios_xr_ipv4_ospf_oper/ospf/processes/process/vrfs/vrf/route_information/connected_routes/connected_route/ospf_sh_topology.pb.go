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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_connected_routes_connected_route

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
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

func (m *OspfShTopology_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_top_common")
	proto.RegisterType((*OspfShRepEl)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_rep_el")
	proto.RegisterType((*OspfShSrUloopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_sr_uloop_path")
	proto.RegisterType((*OspfShNnhInfo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_nnh_info")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_top_path")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.connected_routes.connected_route.ospf_sh_topology")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor_4169ff6e5af11206) }

var fileDescriptor_4169ff6e5af11206 = []byte{
	// 1304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4d, 0x6f, 0x1c, 0xb5,
	0x1b, 0xd7, 0xa4, 0xf9, 0x37, 0x1b, 0x6f, 0xd2, 0x24, 0xf3, 0x6f, 0x93, 0x49, 0x69, 0x69, 0x58,
	0x24, 0x88, 0x10, 0x5a, 0x55, 0x6d, 0x0a, 0x85, 0x4a, 0x45, 0x11, 0x24, 0x62, 0x45, 0x5b, 0x45,
	0xb3, 0x69, 0x25, 0x4e, 0x96, 0x77, 0xc6, 0x93, 0xb5, 0x34, 0x63, 0x5b, 0xb6, 0x77, 0x49, 0x2a,
	0x84, 0x10, 0x12, 0x67, 0x24, 0x4e, 0x1c, 0x11, 0xe2, 0x00, 0x07, 0x2e, 0x7c, 0x07, 0x6e, 0x7c,
	0x15, 0x3e, 0x00, 0x37, 0xe4, 0xe7, 0xf1, 0xec, 0xec, 0x4b, 0xcf, 0x84, 0x4b, 0x34, 0xfb, 0xbc,
	0xf9, 0xf7, 0xbc, 0x3f, 0x21, 0xdb, 0xca, 0xea, 0x82, 0xda, 0x21, 0x75, 0x4a, 0xab, 0x52, 0x9d,
	0x5d, 0x74, 0xb5, 0x51, 0x4e, 0xc5, 0x32, 0x13, 0x36, 0x53, 0x54, 0x28, 0x4b, 0xcf, 0x0d, 0x15,
	0x7a, 0x7c, 0x40, 0x41, 0x52, 0x69, 0x6e, 0xba, 0xfe, 0xcb, 0xcb, 0x65, 0xdc, 0x5a, 0x6e, 0xeb,
	0xaf, 0xee, 0xd8, 0x14, 0xf0, 0xa7, 0x6b, 0xd4, 0xc8, 0x71, 0x2a, 0x64, 0xa1, 0x4c, 0xc5, 0x9c,
	0x50, 0xb2, 0x9b, 0x29, 0x29, 0x79, 0xe6, 0x78, 0x4e, 0x81, 0x67, 0xe7, 0x09, 0x9d, 0xef, 0x23,
	0x72, 0x63, 0x1e, 0x0a, 0xfd, 0xec, 0xe8, 0xf3, 0x7e, 0xfc, 0x06, 0x59, 0x0b, 0x0f, 0x50, 0xc9,
	0x2a, 0x9e, 0x44, 0x7b, 0xd1, 0xfe, 0x6a, 0xda, 0x0e, 0xb4, 0x67, 0xac, 0xe2, 0xf1, 0x2e, 0x69,
	0x8d, 0x4d, 0x81, 0xec, 0x25, 0x60, 0xaf, 0x8c, 0x4d, 0x01, 0xac, 0x6d, 0x72, 0x55, 0x1b, 0x5e,
	0x88, 0xf3, 0xe4, 0x0a, 0x30, 0xc2, 0xaf, 0xf8, 0x4d, 0xb2, 0x8e, 0x5f, 0xb4, 0xe4, 0xf2, 0xcc,
	0x0d, 0x93, 0xe5, 0xbd, 0x68, 0x7f, 0x3d, 0x5d, 0x43, 0xe2, 0x13, 0xa0, 0x75, 0x8e, 0xc9, 0xda,
	0x04, 0x93, 0x40, 0x63, 0x96, 0x67, 0x4a, 0xe6, 0x00, 0x62, 0x3d, 0x0d, 0xbf, 0xe2, 0xd7, 0x09,
	0x91, 0x4c, 0xaa, 0xc0, 0x5b, 0x02, 0xde, 0x14, 0xa5, 0xf3, 0xfb, 0x0a, 0x89, 0xa7, 0x9c, 0xa3,
	0x99, 0xaa, 0x2a, 0x25, 0xe3, 0x0e, 0x59, 0xc7, 0x48, 0x31, 0xc3, 0x19, 0x15, 0xb5, 0xd5, 0x36,
	0x10, 0x0f, 0x0d, 0x67, 0xbd, 0x3c, 0x7e, 0x8b, 0x6c, 0xa0, 0x8c, 0xe3, 0xb4, 0xe2, 0xce, 0x88,
	0x2c, 0xd8, 0x47, 0xd5, 0x53, 0xfe, 0x14, 0x88, 0xf1, 0x3b, 0x64, 0x0b, 0xe5, 0x8c, 0x18, 0xd0,
	0x31, 0x37, 0x56, 0x28, 0x09, 0x2e, 0xaf, 0xa7, 0x68, 0x20, 0x15, 0x83, 0x17, 0x48, 0x6e, 0x64,
	0x3d, 0xa4, 0x5a, 0xd6, 0xfb, 0xbf, 0x1c, 0x64, 0xfb, 0xba, 0xa8, 0x65, 0x0f, 0xc8, 0x36, 0xca,
	0x16, 0xca, 0x7c, 0xc1, 0x4c, 0x4e, 0x73, 0x61, 0x1d, 0x93, 0x19, 0x4f, 0xfe, 0x07, 0xc6, 0xaf,
	0x03, 0xf7, 0x18, 0x99, 0x9f, 0x04, 0x9e, 0xcf, 0x59, 0x78, 0x41, 0x8d, 0x4c, 0xc6, 0x93, 0xab,
	0x53, 0x8e, 0xf5, 0x81, 0x14, 0xff, 0x12, 0xd5, 0x28, 0x46, 0x3a, 0x67, 0xde, 0x41, 0x51, 0xf1,
	0x64, 0x65, 0x2f, 0xda, 0x6f, 0xdf, 0xfb, 0xb2, 0xfb, 0xef, 0x56, 0x5f, 0x77, 0x3a, 0xcb, 0x21,
	0x06, 0xcf, 0x01, 0xd5, 0xa9, 0x4f, 0xfb, 0xcf, 0x51, 0x9d, 0x84, 0x82, 0x89, 0x12, 0x81, 0xb6,
	0xfe, 0x03, 0x40, 0xb1, 0x04, 0x8e, 0x99, 0x28, 0x01, 0xe6, 0xbb, 0x24, 0x6e, 0xd2, 0xaa, 0x8d,
	0x50, 0x46, 0xb8, 0x8b, 0x64, 0x15, 0x42, 0xbf, 0x59, 0xe7, 0xf5, 0x24, 0xd0, 0xe3, 0x2e, 0xf9,
	0x7f, 0x28, 0xbe, 0x91, 0x53, 0x94, 0x9f, 0x67, 0xe5, 0x28, 0xe7, 0x79, 0x42, 0xf6, 0xa2, 0xfd,
	0x56, 0x8a, 0x99, 0x39, 0x1c, 0x39, 0x75, 0x14, 0x18, 0xf1, 0x47, 0xe4, 0x56, 0xb0, 0x6e, 0x1c,
	0xa7, 0xa1, 0x77, 0x0c, 0x3f, 0x13, 0xd6, 0x71, 0xc3, 0xf3, 0xa4, 0x0d, 0x8a, 0xbb, 0xf8, 0x8e,
	0x71, 0xfc, 0x04, 0x24, 0xd2, 0x89, 0x40, 0xfc, 0x01, 0xd9, 0x9d, 0x32, 0x20, 0x07, 0x66, 0x5a,
	0x7b, 0x0d, 0x50, 0x6e, 0x4f, 0xb4, 0x9f, 0x0d, 0xcc, 0x94, 0xea, 0x3e, 0x41, 0xfc, 0xb4, 0x64,
	0x03, 0x5e, 0x52, 0x77, 0xa1, 0x79, 0xb2, 0x0e, 0x1a, 0xd7, 0x80, 0xfe, 0xc4, 0x93, 0x4f, 0x2f,
	0x34, 0x8f, 0xef, 0x90, 0xf6, 0x94, 0x64, 0x72, 0x0d, 0x5b, 0xb1, 0x11, 0x6a, 0x4c, 0x59, 0x1f,
	0x25, 0x94, 0xda, 0x98, 0x32, 0xd5, 0xb7, 0xba, 0x00, 0xc9, 0xce, 0x77, 0x11, 0xb9, 0x56, 0x87,
	0xdb, 0x70, 0x4d, 0x79, 0x09, 0x8d, 0xc3, 0x35, 0x13, 0x86, 0xf2, 0x92, 0x57, 0x5c, 0xba, 0xba,
	0x69, 0x57, 0xd3, 0x0d, 0x64, 0x1c, 0x21, 0xbd, 0x97, 0x43, 0x0b, 0xa0, 0x2c, 0x3e, 0xb2, 0x14,
	0x5a, 0x00, 0x68, 0x88, 0xc5, 0xa7, 0x60, 0xd6, 0x1c, 0x78, 0x86, 0x5d, 0xbb, 0x35, 0x63, 0xd0,
	0x3b, 0xd7, 0xf9, 0x7b, 0xa9, 0x99, 0x91, 0xd6, 0xd0, 0x51, 0xa9, 0x94, 0xa6, 0x9a, 0xb9, 0x61,
	0xfc, 0x6b, 0x44, 0x36, 0x2b, 0x91, 0x19, 0x05, 0x24, 0xd4, 0x4c, 0xa2, 0xbd, 0x2b, 0xfb, 0xed,
	0x7b, 0x5f, 0x5d, 0x56, 0x89, 0x62, 0xcc, 0xd2, 0x8d, 0x09, 0xae, 0x14, 0x60, 0xc5, 0x8f, 0xc8,
	0xcd, 0x79, 0xa8, 0xb4, 0x14, 0xd6, 0x51, 0x2b, 0x5e, 0xf2, 0x10, 0xa6, 0x9d, 0x39, 0xa5, 0x27,
	0xc2, 0xba, 0xbe, 0x78, 0xc9, 0xe3, 0x23, 0x72, 0xa7, 0x51, 0x76, 0x23, 0x29, 0x79, 0x49, 0x85,
	0x74, 0xdc, 0x14, 0x2c, 0xe3, 0xb8, 0x00, 0x70, 0xce, 0xdf, 0x9a, 0x88, 0x9d, 0x82, 0x54, 0xaf,
	0x16, 0x82, 0xad, 0x70, 0x97, 0x5c, 0x6f, 0xcc, 0x58, 0x3f, 0x40, 0x9d, 0xef, 0x1a, 0x18, 0x82,
	0xad, 0x34, 0x9e, 0xf0, 0xfa, 0xc0, 0xea, 0xeb, 0xa2, 0xf3, 0x67, 0x44, 0x36, 0x6b, 0xcf, 0xa4,
	0x1c, 0x42, 0x48, 0xe2, 0xc7, 0xe4, 0x96, 0xe4, 0xe2, 0x6c, 0x38, 0x50, 0x86, 0x4a, 0x7e, 0xee,
	0xe8, 0x50, 0x69, 0x3a, 0x66, 0xe5, 0x88, 0x63, 0x26, 0xb1, 0x34, 0x92, 0x5a, 0xe6, 0x19, 0x3f,
	0x77, 0x9f, 0x2a, 0xfd, 0xc2, 0x0b, 0x40, 0xb5, 0x3e, 0x24, 0xbb, 0x8b, 0xfa, 0x42, 0x53, 0x96,
	0xe7, 0x26, 0x2c, 0xb2, 0x1b, 0x73, 0xca, 0x3d, 0x7d, 0x98, 0xe7, 0xe6, 0xd5, 0x2f, 0x0b, 0xe9,
	0x0a, 0x2a, 0x64, 0xce, 0xcf, 0x43, 0x0d, 0xcd, 0xbf, 0xdc, 0x93, 0xae, 0xe8, 0x79, 0x7e, 0xe7,
	0xaf, 0x56, 0xe3, 0x8e, 0xab, 0xab, 0xe8, 0x2e, 0xb9, 0x5e, 0xe7, 0x7b, 0x26, 0xa2, 0xe8, 0x06,
	0x0e, 0x97, 0xd9, 0x38, 0x3e, 0x22, 0x37, 0xe7, 0x35, 0xac, 0xac, 0x74, 0x00, 0x11, 0x72, 0x39,
	0xab, 0xd7, 0x97, 0x95, 0x06, 0x0c, 0xf1, 0xfd, 0x7a, 0xb5, 0x4c, 0x1c, 0xf0, 0x7e, 0x73, 0x6b,
	0x43, 0x0a, 0x71, 0x3e, 0x05, 0xe8, 0x87, 0xc8, 0x5a, 0xd8, 0x2c, 0xcb, 0x78, 0x0d, 0x4c, 0x6f,
	0x96, 0x66, 0x06, 0x58, 0x26, 0x72, 0xd8, 0x53, 0xab, 0xf5, 0x0c, 0xf0, 0x94, 0xf8, 0xc3, 0x1a,
	0xb5, 0xf7, 0x9a, 0x0a, 0x4b, 0xab, 0x8c, 0x59, 0xe7, 0x7d, 0x60, 0x99, 0x83, 0x5d, 0xd5, 0x0a,
	0xa3, 0xe8, 0x84, 0xb9, 0x61, 0xcf, 0x3e, 0xf5, 0xec, 0x1e, 0x70, 0xe3, 0x03, 0xb2, 0x33, 0xab,
	0x3b, 0xca, 0x2a, 0x0c, 0x1f, 0xec, 0xae, 0x56, 0x40, 0x8d, 0x8a, 0xcf, 0xb3, 0x4a, 0xfb, 0xaf,
	0x06, 0x75, 0x58, 0xe1, 0xad, 0xa9, 0x7d, 0x18, 0x16, 0xf8, 0x2e, 0x69, 0x95, 0x96, 0x61, 0xdd,
	0xe0, 0xcc, 0x5e, 0x29, 0x2d, 0x83, 0x32, 0xd9, 0x21, 0x2b, 0xf5, 0x85, 0x40, 0xf0, 0xee, 0x60,
	0x78, 0x1c, 0xdc, 0x21, 0x6d, 0x60, 0x60, 0x77, 0x86, 0x11, 0x4c, 0x3c, 0xe9, 0x18, 0x28, 0xf1,
	0x1f, 0x11, 0x79, 0xcd, 0x1a, 0xda, 0xd4, 0x3a, 0x1b, 0x2b, 0x91, 0xfb, 0x15, 0x8d, 0x90, 0xd7,
	0x60, 0x8b, 0x7d, 0x1b, 0x5d, 0xd6, 0x8c, 0x98, 0x99, 0x62, 0x69, 0x62, 0xcd, 0xd3, 0x1a, 0xe9,
	0x61, 0x0d, 0x14, 0xe2, 0x37, 0xb9, 0x94, 0x30, 0xea, 0x79, 0x98, 0xfe, 0xed, 0x26, 0xd6, 0x7e,
	0xbf, 0xcc, 0x65, 0xd5, 0xd8, 0xb1, 0xa6, 0xe1, 0x37, 0x6c, 0x82, 0x56, 0x7a, 0x63, 0x2a, 0x39,
	0xa9, 0x1d, 0xeb, 0x53, 0x34, 0xff, 0x80, 0x24, 0xb3, 0xaa, 0xd6, 0x4c, 0x14, 0x37, 0x16, 0xb2,
	0xda, 0x37, 0x41, 0x6d, 0xe1, 0x45, 0x6b, 0x60, 0x91, 0xa2, 0xe2, 0xe6, 0xc2, 0x8b, 0x7d, 0xe3,
	0xb7, 0x29, 0xa8, 0x3e, 0x26, 0xb7, 0x5f, 0xf5, 0x22, 0x2c, 0x25, 0xd0, 0xde, 0x02, 0xed, 0x9d,
	0xb9, 0x67, 0xfd, 0x76, 0x02, 0xfd, 0xdf, 0x22, 0xb2, 0xb5, 0x30, 0x00, 0x92, 0x18, 0x26, 0xfe,
	0xd7, 0x97, 0x96, 0xce, 0x7a, 0x30, 0xa6, 0x1b, 0x73, 0x83, 0xa7, 0xf3, 0xc3, 0xf2, 0xcc, 0xbc,
	0x81, 0xf3, 0xbe, 0xe9, 0x8a, 0x70, 0xa1, 0xdf, 0x9b, 0xea, 0x65, 0x3c, 0x1f, 0x9a, 0x2b, 0x65,
	0xf6, 0x58, 0xbf, 0x1f, 0x56, 0x64, 0x23, 0x89, 0x17, 0xfb, 0x42, 0xa3, 0x1d, 0x2c, 0x36, 0xda,
	0x6d, 0x42, 0xc2, 0x45, 0xed, 0x5b, 0xed, 0x01, 0xbc, 0xb9, 0x8a, 0xc7, 0xb4, 0x6f, 0xb6, 0xb7,
	0xeb, 0x5b, 0x6f, 0xe2, 0x64, 0xf2, 0x1e, 0xe4, 0x02, 0xef, 0x83, 0x8f, 0x6b, 0x6a, 0xfc, 0x53,
	0x54, 0x1b, 0xf2, 0x2e, 0x27, 0xef, 0x43, 0x2b, 0x7d, 0x73, 0x69, 0xb1, 0x6f, 0xfe, 0xaf, 0x08,
	0xde, 0xf4, 0xfc, 0x86, 0xfa, 0x71, 0x02, 0x12, 0xaa, 0xea, 0xe1, 0x25, 0x17, 0x48, 0xbd, 0x6a,
	0x02, 0x44, 0x5f, 0xca, 0x83, 0xab, 0xf0, 0x0f, 0xe7, 0xfd, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x45, 0xc5, 0xd3, 0xdc, 0x8a, 0x0e, 0x00, 0x00,
}
