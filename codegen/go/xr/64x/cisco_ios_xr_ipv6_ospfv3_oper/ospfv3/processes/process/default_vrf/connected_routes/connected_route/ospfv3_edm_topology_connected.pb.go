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
// source: ospfv3_edm_topology_connected.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_connected_routes_connected_route

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

type Ospfv3EdmTopologyConnected_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,3,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopologyConnected_KEYS) Reset()         { *m = Ospfv3EdmTopologyConnected_KEYS{} }
func (m *Ospfv3EdmTopologyConnected_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyConnected_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmTopologyConnected_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{0}
}

func (m *Ospfv3EdmTopologyConnected_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyConnected_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyConnected_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmTopologyConnected_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS.Size(m)
}
func (m *Ospfv3EdmTopologyConnected_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyConnected_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyConnected_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmTopologyConnected_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ospfv3EdmTopologyConnected_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type Ospfv3EdmTopology struct {
	RouteId              string   `protobuf:"bytes,1,opt,name=route_id,json=routeId,proto3" json:"route_id,omitempty"`
	RouteDistance        uint32   `protobuf:"varint,2,opt,name=route_distance,json=routeDistance,proto3" json:"route_distance,omitempty"`
	RouteCost            uint32   `protobuf:"varint,3,opt,name=route_cost,json=routeCost,proto3" json:"route_cost,omitempty"`
	RouteType            uint32   `protobuf:"varint,4,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopology) Reset()         { *m = Ospfv3EdmTopology{} }
func (m *Ospfv3EdmTopology) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopology) ProtoMessage()    {}
func (*Ospfv3EdmTopology) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{1}
}

func (m *Ospfv3EdmTopology) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopology.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopology) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopology.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopology) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopology.Merge(m, src)
}
func (m *Ospfv3EdmTopology) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopology.Size(m)
}
func (m *Ospfv3EdmTopology) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopology.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopology proto.InternalMessageInfo

func (m *Ospfv3EdmTopology) GetRouteId() string {
	if m != nil {
		return m.RouteId
	}
	return ""
}

func (m *Ospfv3EdmTopology) GetRouteDistance() uint32 {
	if m != nil {
		return m.RouteDistance
	}
	return 0
}

func (m *Ospfv3EdmTopology) GetRouteCost() uint32 {
	if m != nil {
		return m.RouteCost
	}
	return 0
}

func (m *Ospfv3EdmTopology) GetRouteType() uint32 {
	if m != nil {
		return m.RouteType
	}
	return 0
}

type Ospfv3ShBackupPath struct {
	BackupRouteInterfaceName  string   `protobuf:"bytes,1,opt,name=backup_route_interface_name,json=backupRouteInterfaceName,proto3" json:"backup_route_interface_name,omitempty"`
	BackupRouteNextHopAddress string   `protobuf:"bytes,2,opt,name=backup_route_next_hop_address,json=backupRouteNextHopAddress,proto3" json:"backup_route_next_hop_address,omitempty"`
	BackupRouteSource         string   `protobuf:"bytes,3,opt,name=backup_route_source,json=backupRouteSource,proto3" json:"backup_route_source,omitempty"`
	BackupMetric              uint32   `protobuf:"varint,4,opt,name=backup_metric,json=backupMetric,proto3" json:"backup_metric,omitempty"`
	PrimaryPath               bool     `protobuf:"varint,5,opt,name=primary_path,json=primaryPath,proto3" json:"primary_path,omitempty"`
	LineCardDisjoint          bool     `protobuf:"varint,6,opt,name=line_card_disjoint,json=lineCardDisjoint,proto3" json:"line_card_disjoint,omitempty"`
	Downstream                bool     `protobuf:"varint,7,opt,name=downstream,proto3" json:"downstream,omitempty"`
	NodeProtect               bool     `protobuf:"varint,8,opt,name=node_protect,json=nodeProtect,proto3" json:"node_protect,omitempty"`
	SrlgDisjoint              bool     `protobuf:"varint,9,opt,name=srlg_disjoint,json=srlgDisjoint,proto3" json:"srlg_disjoint,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Ospfv3ShBackupPath) Reset()         { *m = Ospfv3ShBackupPath{} }
func (m *Ospfv3ShBackupPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShBackupPath) ProtoMessage()    {}
func (*Ospfv3ShBackupPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{2}
}

func (m *Ospfv3ShBackupPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShBackupPath.Unmarshal(m, b)
}
func (m *Ospfv3ShBackupPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShBackupPath.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShBackupPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShBackupPath.Merge(m, src)
}
func (m *Ospfv3ShBackupPath) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShBackupPath.Size(m)
}
func (m *Ospfv3ShBackupPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShBackupPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShBackupPath proto.InternalMessageInfo

func (m *Ospfv3ShBackupPath) GetBackupRouteInterfaceName() string {
	if m != nil {
		return m.BackupRouteInterfaceName
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupRouteNextHopAddress() string {
	if m != nil {
		return m.BackupRouteNextHopAddress
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupRouteSource() string {
	if m != nil {
		return m.BackupRouteSource
	}
	return ""
}

func (m *Ospfv3ShBackupPath) GetBackupMetric() uint32 {
	if m != nil {
		return m.BackupMetric
	}
	return 0
}

func (m *Ospfv3ShBackupPath) GetPrimaryPath() bool {
	if m != nil {
		return m.PrimaryPath
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetLineCardDisjoint() bool {
	if m != nil {
		return m.LineCardDisjoint
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetDownstream() bool {
	if m != nil {
		return m.Downstream
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetNodeProtect() bool {
	if m != nil {
		return m.NodeProtect
	}
	return false
}

func (m *Ospfv3ShBackupPath) GetSrlgDisjoint() bool {
	if m != nil {
		return m.SrlgDisjoint
	}
	return false
}

type Ospfv3ShNnhInfo struct {
	NeighborNextHopIntfIndex uint32   `protobuf:"varint,1,opt,name=neighbor_next_hop_intf_index,json=neighborNextHopIntfIndex,proto3" json:"neighbor_next_hop_intf_index,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Ospfv3ShNnhInfo) Reset()         { *m = Ospfv3ShNnhInfo{} }
func (m *Ospfv3ShNnhInfo) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShNnhInfo) ProtoMessage()    {}
func (*Ospfv3ShNnhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{3}
}

func (m *Ospfv3ShNnhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShNnhInfo.Unmarshal(m, b)
}
func (m *Ospfv3ShNnhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShNnhInfo.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShNnhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShNnhInfo.Merge(m, src)
}
func (m *Ospfv3ShNnhInfo) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShNnhInfo.Size(m)
}
func (m *Ospfv3ShNnhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShNnhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShNnhInfo proto.InternalMessageInfo

func (m *Ospfv3ShNnhInfo) GetNeighborNextHopIntfIndex() uint32 {
	if m != nil {
		return m.NeighborNextHopIntfIndex
	}
	return 0
}

type Ospfv3EdmTopPath struct {
	InterfaceName        string              `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	RoutePathNextHop     string              `protobuf:"bytes,2,opt,name=route_path_next_hop,json=routePathNextHop,proto3" json:"route_path_next_hop,omitempty"`
	RoutePathId          uint32              `protobuf:"varint,3,opt,name=route_path_id,json=routePathId,proto3" json:"route_path_id,omitempty"`
	RouteBackupPath      *Ospfv3ShBackupPath `protobuf:"bytes,4,opt,name=route_backup_path,json=routeBackupPath,proto3" json:"route_backup_path,omitempty"`
	NeighborNextHop      []*Ospfv3ShNnhInfo  `protobuf:"bytes,5,rep,name=neighbor_next_hop,json=neighborNextHop,proto3" json:"neighbor_next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ospfv3EdmTopPath) Reset()         { *m = Ospfv3EdmTopPath{} }
func (m *Ospfv3EdmTopPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopPath) ProtoMessage()    {}
func (*Ospfv3EdmTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{4}
}

func (m *Ospfv3EdmTopPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopPath.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopPath.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopPath.Merge(m, src)
}
func (m *Ospfv3EdmTopPath) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopPath.Size(m)
}
func (m *Ospfv3EdmTopPath) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopPath.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopPath proto.InternalMessageInfo

func (m *Ospfv3EdmTopPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmTopPath) GetRoutePathNextHop() string {
	if m != nil {
		return m.RoutePathNextHop
	}
	return ""
}

func (m *Ospfv3EdmTopPath) GetRoutePathId() uint32 {
	if m != nil {
		return m.RoutePathId
	}
	return 0
}

func (m *Ospfv3EdmTopPath) GetRouteBackupPath() *Ospfv3ShBackupPath {
	if m != nil {
		return m.RouteBackupPath
	}
	return nil
}

func (m *Ospfv3EdmTopPath) GetNeighborNextHop() []*Ospfv3ShNnhInfo {
	if m != nil {
		return m.NeighborNextHop
	}
	return nil
}

type Ospfv3EdmTopologyConnected struct {
	RouteTopology        *Ospfv3EdmTopology  `protobuf:"bytes,50,opt,name=route_topology,json=routeTopology,proto3" json:"route_topology,omitempty"`
	RoutePath            []*Ospfv3EdmTopPath `protobuf:"bytes,51,rep,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ospfv3EdmTopologyConnected) Reset()         { *m = Ospfv3EdmTopologyConnected{} }
func (m *Ospfv3EdmTopologyConnected) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyConnected) ProtoMessage()    {}
func (*Ospfv3EdmTopologyConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{5}
}

func (m *Ospfv3EdmTopologyConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyConnected.Merge(m, src)
}
func (m *Ospfv3EdmTopologyConnected) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyConnected.Size(m)
}
func (m *Ospfv3EdmTopologyConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyConnected.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyConnected proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyConnected) GetRouteTopology() *Ospfv3EdmTopology {
	if m != nil {
		return m.RouteTopology
	}
	return nil
}

func (m *Ospfv3EdmTopologyConnected) GetRoutePath() []*Ospfv3EdmTopPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmTopologyConnected_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_topology_connected_KEYS")
	proto.RegisterType((*Ospfv3EdmTopology)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_topology")
	proto.RegisterType((*Ospfv3ShBackupPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_sh_backup_path")
	proto.RegisterType((*Ospfv3ShNnhInfo)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_sh_nnh_info")
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3EdmTopologyConnected)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_topology_connected")
}

func init() {
	proto.RegisterFile("ospfv3_edm_topology_connected.proto", fileDescriptor_a3518cb55007c353)
}

var fileDescriptor_a3518cb55007c353 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0xf4, 0x23, 0x4e, 0xd2, 0x0f, 0x57, 0xa0, 0xad, 0xa0, 0x28, 0x6c, 0x55, 0x29,
	0x07, 0xc8, 0xa1, 0x95, 0xb8, 0x81, 0x80, 0x16, 0x89, 0x08, 0xa8, 0xaa, 0x6d, 0x2f, 0x9c, 0xac,
	0xad, 0x3d, 0x9b, 0x18, 0x12, 0x7b, 0x65, 0x3b, 0x6d, 0xc2, 0x99, 0xbf, 0x80, 0x38, 0x72, 0x40,
	0xea, 0x9f, 0x42, 0xe2, 0xb7, 0x20, 0x7f, 0x6c, 0xb2, 0x2d, 0x55, 0x8f, 0xe5, 0xb6, 0x7e, 0xf3,
	0xec, 0x79, 0x33, 0x6f, 0x32, 0x41, 0x3b, 0x52, 0x17, 0xf9, 0xf9, 0x3e, 0x01, 0x36, 0x22, 0x46,
	0x16, 0x72, 0x28, 0xfb, 0x53, 0x42, 0xa5, 0x10, 0x40, 0x0d, 0xb0, 0x6e, 0xa1, 0xa4, 0x91, 0x98,
	0x52, 0xae, 0xa9, 0x24, 0x5c, 0x6a, 0x32, 0x51, 0x84, 0x17, 0xe7, 0xcf, 0x49, 0xb8, 0x26, 0x0b,
	0x50, 0x5d, 0xff, 0x6d, 0xb9, 0x14, 0xb4, 0x06, 0x5d, 0x7e, 0x75, 0x19, 0xe4, 0xd9, 0x78, 0x68,
	0xc8, 0xb9, 0xca, 0xbb, 0xb3, 0x37, 0x89, 0x92, 0x63, 0x03, 0xfa, 0x3a, 0x90, 0x7c, 0x8b, 0x50,
	0x72, 0xab, 0x18, 0xf2, 0xfe, 0xed, 0xa7, 0x13, 0xfc, 0x04, 0x35, 0xc3, 0xf3, 0x44, 0x64, 0x23,
	0x88, 0xa3, 0x76, 0xd4, 0xa9, 0xa7, 0x8d, 0x80, 0x1d, 0x65, 0x23, 0xc0, 0x0f, 0xd0, 0x52, 0xa1,
	0x20, 0xe7, 0x93, 0x78, 0xc1, 0x05, 0xc3, 0x09, 0xef, 0xa0, 0x96, 0xff, 0x22, 0x43, 0x10, 0x7d,
	0x33, 0x88, 0x6b, 0xed, 0xa8, 0xd3, 0x4a, 0x9b, 0x1e, 0xfc, 0xe0, 0xb0, 0xe4, 0x7b, 0x84, 0x36,
	0x6f, 0x90, 0x81, 0xb7, 0xd0, 0x8a, 0xd3, 0x49, 0x38, 0x0b, 0x39, 0x97, 0xdd, 0xb9, 0xc7, 0xf0,
	0x2e, 0x5a, 0xf5, 0x21, 0xc6, 0xb5, 0xc9, 0x04, 0x05, 0x97, 0xb7, 0x95, 0xb6, 0x1c, 0x7a, 0x18,
	0x40, 0xbc, 0x8d, 0x90, 0xa7, 0x51, 0xa9, 0x4d, 0xc8, 0x5d, 0x77, 0xc8, 0x81, 0xd4, 0x66, 0x1e,
	0x36, 0xd3, 0x02, 0xe2, 0x7b, 0x95, 0xf0, 0xe9, 0xb4, 0x80, 0xe4, 0xb2, 0x86, 0xee, 0x07, 0x5d,
	0x7a, 0x40, 0xce, 0x32, 0xfa, 0x65, 0x5c, 0x90, 0x22, 0x33, 0x03, 0xfc, 0x02, 0x3d, 0x0c, 0xc7,
	0x20, 0x50, 0x18, 0x50, 0x79, 0x46, 0xa1, 0xda, 0xa0, 0xd8, 0x53, 0x52, 0x27, 0xb9, 0x24, 0xb8,
	0x6e, 0xbd, 0x42, 0xdb, 0x57, 0xae, 0x0b, 0x98, 0x18, 0x32, 0x90, 0x05, 0xc9, 0x18, 0x53, 0xa0,
	0x75, 0x68, 0xe2, 0x56, 0xe5, 0x81, 0x23, 0x98, 0x98, 0x77, 0xb2, 0x78, 0xed, 0x09, 0xb8, 0x8b,
	0x36, 0xaf, 0xbc, 0xa0, 0xe5, 0x58, 0x51, 0x70, 0x15, 0xd6, 0xd3, 0x8d, 0xca, 0xbd, 0x13, 0x17,
	0xb0, 0x3e, 0x04, 0xfe, 0x08, 0x8c, 0xe2, 0x34, 0x14, 0xdb, 0xf4, 0xe0, 0x47, 0x87, 0x79, 0x9f,
	0xf9, 0x28, 0x53, 0x53, 0x57, 0x65, 0xbc, 0xd8, 0x8e, 0x3a, 0x2b, 0xd6, 0x67, 0x87, 0x1d, 0xdb,
	0xc2, 0x9f, 0x22, 0x3c, 0xe4, 0x02, 0x08, 0xcd, 0x14, 0xb3, 0xbd, 0xff, 0x2c, 0xb9, 0x30, 0xf1,
	0x92, 0x23, 0xae, 0xdb, 0xc8, 0x41, 0xa6, 0xd8, 0x61, 0xc0, 0xf1, 0x63, 0x84, 0x98, 0xbc, 0x10,
	0xda, 0x28, 0xc8, 0x46, 0xf1, 0xb2, 0x63, 0x55, 0x10, 0x9b, 0x50, 0x48, 0x06, 0xc4, 0x8e, 0x3c,
	0x50, 0x13, 0xaf, 0xf8, 0x84, 0x16, 0x3b, 0xf6, 0x90, 0x15, 0xae, 0xd5, 0xb0, 0x3f, 0xcf, 0x55,
	0x77, 0x9c, 0xa6, 0x05, 0xcb, 0x3c, 0xc9, 0x29, 0xc2, 0x73, 0x9f, 0x84, 0x18, 0x10, 0x2e, 0x72,
	0x89, 0x5f, 0xa2, 0x47, 0x02, 0x78, 0x7f, 0x70, 0x26, 0xd5, 0xbc, 0xc3, 0x5c, 0x98, 0x9c, 0x70,
	0xc1, 0x60, 0xe2, 0x5c, 0x6a, 0xa5, 0x71, 0xc9, 0x09, 0x1d, 0xee, 0x09, 0x93, 0xf7, 0x6c, 0x3c,
	0xf9, 0x53, 0xbb, 0x3e, 0x96, 0xde, 0xfc, 0x5d, 0xb4, 0x7a, 0xa3, 0xdf, 0x2d, 0x7e, 0xc5, 0xe4,
	0x67, 0x68, 0xd3, 0x7b, 0x63, 0x2f, 0xcd, 0x04, 0x04, 0x6b, 0xd7, 0x5d, 0xc8, 0xb6, 0x34, 0xa4,
	0xc5, 0x09, 0x6a, 0x55, 0xe8, 0x9c, 0x85, 0x69, 0x6d, 0xcc, 0x88, 0x3d, 0x86, 0x2f, 0x23, 0xb4,
	0xe1, 0x49, 0x95, 0x61, 0x74, 0x56, 0x36, 0xf6, 0xbe, 0x76, 0xef, 0x60, 0x63, 0x74, 0x6f, 0xfc,
	0x39, 0xa4, 0x6b, 0x2e, 0xf8, 0xc6, 0x21, 0x6e, 0x4c, 0x7e, 0x45, 0x68, 0xe3, 0x9f, 0xde, 0xc7,
	0x8b, 0xed, 0x5a, 0xa7, 0xb1, 0x77, 0x71, 0xc7, 0x42, 0xcb, 0x79, 0x48, 0xd7, 0xae, 0x39, 0x9d,
	0xfc, 0x5e, 0x40, 0xdb, 0xb7, 0xae, 0x3f, 0xfc, 0x33, 0x2a, 0xf7, 0x4c, 0x19, 0x8c, 0xf7, 0x5c,
	0xb7, 0x27, 0x77, 0x59, 0x44, 0x55, 0x5c, 0xd8, 0x70, 0xa7, 0xe5, 0x8e, 0xfc, 0x11, 0x95, 0x3b,
	0xcc, 0xcd, 0xc2, 0xbe, 0x6b, 0xf1, 0xff, 0x50, 0xe7, 0x27, 0xa1, 0x3e, 0x1b, 0xd7, 0xb3, 0x25,
	0xf7, 0x47, 0xb6, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0xde, 0xcb, 0xe9, 0x50, 0xef, 0x06, 0x00,
	0x00,
}