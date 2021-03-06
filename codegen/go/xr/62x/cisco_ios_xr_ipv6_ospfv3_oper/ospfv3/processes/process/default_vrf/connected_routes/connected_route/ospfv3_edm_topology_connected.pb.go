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

type Ospfv3EdmTopPath struct {
	InterfaceName        string              `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	RoutePathNextHop     string              `protobuf:"bytes,2,opt,name=route_path_next_hop,json=routePathNextHop,proto3" json:"route_path_next_hop,omitempty"`
	RoutePathId          uint32              `protobuf:"varint,3,opt,name=route_path_id,json=routePathId,proto3" json:"route_path_id,omitempty"`
	RouteBackupPath      *Ospfv3ShBackupPath `protobuf:"bytes,4,opt,name=route_backup_path,json=routeBackupPath,proto3" json:"route_backup_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Ospfv3EdmTopPath) Reset()         { *m = Ospfv3EdmTopPath{} }
func (m *Ospfv3EdmTopPath) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopPath) ProtoMessage()    {}
func (*Ospfv3EdmTopPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3518cb55007c353, []int{3}
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
	return fileDescriptor_a3518cb55007c353, []int{4}
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
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3EdmTopologyConnected)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.connected_routes.connected_route.ospfv3_edm_topology_connected")
}

func init() {
	proto.RegisterFile("ospfv3_edm_topology_connected.proto", fileDescriptor_a3518cb55007c353)
}

var fileDescriptor_a3518cb55007c353 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x4e, 0x54, 0x3d,
	0x14, 0xce, 0x85, 0xff, 0x07, 0xa6, 0xcc, 0x20, 0x94, 0x68, 0x2e, 0x31, 0x98, 0xf1, 0x12, 0x12,
	0x16, 0x7a, 0x17, 0x90, 0xb8, 0x33, 0x51, 0xc1, 0x44, 0xa2, 0x12, 0x32, 0xb0, 0x71, 0xd5, 0x94,
	0xf6, 0x0c, 0x53, 0x9d, 0xb9, 0x6d, 0xda, 0x0e, 0xce, 0xb8, 0xf6, 0x15, 0x8c, 0x4b, 0x77, 0x3c,
	0x83, 0xef, 0xe2, 0xcb, 0x98, 0x9e, 0xf6, 0x0e, 0x17, 0x42, 0x58, 0xe2, 0xae, 0xfd, 0xce, 0xd7,
	0x9e, 0xef, 0x9c, 0xef, 0xb4, 0x64, 0x4b, 0x3b, 0xd3, 0xbf, 0xd8, 0x63, 0x20, 0x47, 0xcc, 0x6b,
	0xa3, 0x87, 0xfa, 0x7c, 0xca, 0x84, 0xae, 0x2a, 0x10, 0x1e, 0x64, 0x69, 0xac, 0xf6, 0x9a, 0x0a,
	0xa1, 0x9c, 0xd0, 0x4c, 0x69, 0xc7, 0x26, 0x96, 0x29, 0x73, 0xf1, 0x82, 0xa5, 0x63, 0xda, 0x80,
	0x2d, 0xe3, 0x3a, 0x70, 0x05, 0x38, 0x07, 0xae, 0x5e, 0x95, 0x12, 0xfa, 0x7c, 0x3c, 0xf4, 0xec,
	0xc2, 0xf6, 0xcb, 0xd9, 0x9d, 0xcc, 0xea, 0xb1, 0x07, 0x77, 0x13, 0x28, 0xbe, 0x67, 0xa4, 0xb8,
	0x53, 0x0c, 0x7b, 0xff, 0xf6, 0xd3, 0x09, 0x7d, 0x4a, 0xda, 0xe9, 0x7a, 0x56, 0xf1, 0x11, 0xe4,
	0x59, 0x37, 0xdb, 0x69, 0xf5, 0x96, 0x13, 0x76, 0xc4, 0x47, 0x40, 0x1f, 0x91, 0x05, 0x63, 0xa1,
	0xaf, 0x26, 0xf9, 0x1c, 0x06, 0xd3, 0x8e, 0x6e, 0x91, 0x4e, 0x5c, 0xb1, 0x21, 0x54, 0xe7, 0x7e,
	0x90, 0xcf, 0x77, 0xb3, 0x9d, 0x4e, 0xaf, 0x1d, 0xc1, 0x0f, 0x88, 0x15, 0x3f, 0x32, 0xb2, 0x7e,
	0x8b, 0x0c, 0xba, 0x41, 0x96, 0x50, 0x27, 0x53, 0x32, 0xe5, 0x5c, 0xc4, 0xfd, 0xa1, 0xa4, 0xdb,
	0x64, 0x25, 0x86, 0xa4, 0x72, 0x9e, 0x57, 0x02, 0x30, 0x6f, 0xa7, 0xd7, 0x41, 0xf4, 0x20, 0x81,
	0x74, 0x93, 0x90, 0x48, 0x13, 0xda, 0xf9, 0x94, 0xbb, 0x85, 0xc8, 0xbe, 0x76, 0xfe, 0x2a, 0xec,
	0xa7, 0x06, 0xf2, 0xff, 0x1a, 0xe1, 0xd3, 0xa9, 0x81, 0xe2, 0x72, 0x9e, 0x3c, 0x4c, 0xba, 0xdc,
	0x80, 0x9d, 0x71, 0xf1, 0x65, 0x6c, 0x98, 0xe1, 0x7e, 0x40, 0x5f, 0x92, 0xc7, 0x69, 0x9b, 0x04,
	0x56, 0x1e, 0x6c, 0x9f, 0x0b, 0x68, 0x36, 0x28, 0x8f, 0x94, 0x1e, 0x4a, 0xae, 0x09, 0xd8, 0xad,
	0x57, 0x64, 0xf3, 0xda, 0xf1, 0x0a, 0x26, 0x9e, 0x0d, 0xb4, 0x61, 0x5c, 0x4a, 0x0b, 0xce, 0xa5,
	0x26, 0x6e, 0x34, 0x2e, 0x38, 0x82, 0x89, 0x7f, 0xa7, 0xcd, 0xeb, 0x48, 0xa0, 0x25, 0x59, 0xbf,
	0x76, 0x83, 0xd3, 0x63, 0x2b, 0x00, 0x2b, 0x6c, 0xf5, 0xd6, 0x1a, 0xe7, 0x4e, 0x30, 0x10, 0x7c,
	0x48, 0xfc, 0x11, 0x78, 0xab, 0x44, 0x2a, 0xb6, 0x1d, 0xc1, 0x8f, 0x88, 0x45, 0x9f, 0xd5, 0x88,
	0xdb, 0x29, 0x56, 0x99, 0xff, 0xdf, 0xcd, 0x76, 0x96, 0x82, 0xcf, 0x88, 0x1d, 0x87, 0xc2, 0x9f,
	0x11, 0x3a, 0x54, 0x15, 0x30, 0xc1, 0xad, 0x0c, 0xbd, 0xff, 0xac, 0x55, 0xe5, 0xf3, 0x05, 0x24,
	0xae, 0x86, 0xc8, 0x3e, 0xb7, 0xf2, 0x20, 0xe1, 0xf4, 0x09, 0x21, 0x52, 0x7f, 0xad, 0x9c, 0xb7,
	0xc0, 0x47, 0xf9, 0x22, 0xb2, 0x1a, 0x48, 0x48, 0x58, 0x69, 0x09, 0x2c, 0x8c, 0x3c, 0x08, 0x9f,
	0x2f, 0xc5, 0x84, 0x01, 0x3b, 0x8e, 0x50, 0x10, 0xee, 0xec, 0xf0, 0xfc, 0x2a, 0x57, 0x0b, 0x39,
	0xed, 0x00, 0xd6, 0x79, 0x8a, 0xdf, 0x73, 0x37, 0x07, 0x28, 0xda, 0xb4, 0x4d, 0x56, 0x6e, 0x75,
	0xa6, 0xa3, 0xae, 0xd9, 0xf1, 0x9c, 0xac, 0xc7, 0x2e, 0x86, 0x43, 0x33, 0x33, 0x92, 0x09, 0xab,
	0x18, 0x0a, 0xc5, 0x27, 0x0b, 0x68, 0x41, 0x3a, 0x0d, 0xba, 0x92, 0x69, 0xae, 0x96, 0x67, 0xc4,
	0x43, 0x49, 0x2f, 0x33, 0xb2, 0x16, 0x49, 0x8d, 0xb1, 0xc1, 0xa6, 0x2f, 0xef, 0x7e, 0x2b, 0xef,
	0xe1, 0x6d, 0x97, 0xb7, 0x0e, 0x6e, 0xef, 0x01, 0x06, 0xdf, 0x20, 0x12, 0xa4, 0x16, 0x7f, 0xe6,
	0xc8, 0xe6, 0x9d, 0x5f, 0x00, 0xfd, 0x95, 0xd5, 0x6f, 0xad, 0x0e, 0xe6, 0xbb, 0x58, 0xc7, 0xe4,
	0x3e, 0xeb, 0x68, 0x8a, 0x4b, 0xaf, 0xfc, 0xb4, 0xfe, 0x27, 0x7e, 0x66, 0xf5, 0x3b, 0xc6, 0x2e,
	0xef, 0x75, 0xe7, 0xff, 0x91, 0xba, 0xd8, 0xe3, 0xd6, 0x6c, 0x10, 0xce, 0x16, 0xf0, 0x33, 0xdf,
	0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd4, 0x0b, 0x6a, 0xf3, 0x05, 0x00, 0x00,
}
