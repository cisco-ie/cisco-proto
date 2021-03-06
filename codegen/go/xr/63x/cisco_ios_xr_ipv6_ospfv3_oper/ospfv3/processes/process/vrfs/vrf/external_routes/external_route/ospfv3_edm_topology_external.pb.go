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
// source: ospfv3_edm_topology_external.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_external_routes_external_route

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

type Ospfv3EdmTopologyExternal_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmTopologyExternal_KEYS) Reset()         { *m = Ospfv3EdmTopologyExternal_KEYS{} }
func (m *Ospfv3EdmTopologyExternal_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyExternal_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmTopologyExternal_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{0}
}

func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.Size(m)
}
func (m *Ospfv3EdmTopologyExternal_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyExternal_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyExternal_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Ospfv3EdmTopologyExternal_KEYS) GetPrefixLength() uint32 {
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
	return fileDescriptor_b36ab0046955ae4e, []int{1}
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
	return fileDescriptor_b36ab0046955ae4e, []int{2}
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
	return fileDescriptor_b36ab0046955ae4e, []int{3}
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

type Ospfv3ShRouteExtendedComm struct {
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

func (m *Ospfv3ShRouteExtendedComm) Reset()         { *m = Ospfv3ShRouteExtendedComm{} }
func (m *Ospfv3ShRouteExtendedComm) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShRouteExtendedComm) ProtoMessage()    {}
func (*Ospfv3ShRouteExtendedComm) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{4}
}

func (m *Ospfv3ShRouteExtendedComm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Unmarshal(m, b)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShRouteExtendedComm.Merge(m, src)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShRouteExtendedComm.Size(m)
}
func (m *Ospfv3ShRouteExtendedComm) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShRouteExtendedComm.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShRouteExtendedComm proto.InternalMessageInfo

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityDomainIdValue() []uint32 {
	if m != nil {
		return m.ExtendedCommunityDomainIdValue
	}
	return nil
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunitylDomainIdType() uint32 {
	if m != nil {
		return m.ExtendedCommunitylDomainIdType
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityAreaId() uint32 {
	if m != nil {
		return m.ExtendedCommunityAreaId
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityRouterId() string {
	if m != nil {
		return m.ExtendedCommunityRouterId
	}
	return ""
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityRouteType() uint32 {
	if m != nil {
		return m.ExtendedCommunityRouteType
	}
	return 0
}

func (m *Ospfv3ShRouteExtendedComm) GetExtendedCommunityOptions() uint32 {
	if m != nil {
		return m.ExtendedCommunityOptions
	}
	return 0
}

type Ospfv3EdmTopologyExternal struct {
	RouteTopology          *Ospfv3EdmTopology         `protobuf:"bytes,50,opt,name=route_topology,json=routeTopology,proto3" json:"route_topology,omitempty"`
	RoutePath              []*Ospfv3EdmTopPath        `protobuf:"bytes,51,rep,name=route_path,json=routePath,proto3" json:"route_path,omitempty"`
	RouteExtendedCommunity *Ospfv3ShRouteExtendedComm `protobuf:"bytes,52,opt,name=route_extended_community,json=routeExtendedCommunity,proto3" json:"route_extended_community,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                   `json:"-"`
	XXX_unrecognized       []byte                     `json:"-"`
	XXX_sizecache          int32                      `json:"-"`
}

func (m *Ospfv3EdmTopologyExternal) Reset()         { *m = Ospfv3EdmTopologyExternal{} }
func (m *Ospfv3EdmTopologyExternal) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmTopologyExternal) ProtoMessage()    {}
func (*Ospfv3EdmTopologyExternal) Descriptor() ([]byte, []int) {
	return fileDescriptor_b36ab0046955ae4e, []int{5}
}

func (m *Ospfv3EdmTopologyExternal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Unmarshal(m, b)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmTopologyExternal.Merge(m, src)
}
func (m *Ospfv3EdmTopologyExternal) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmTopologyExternal.Size(m)
}
func (m *Ospfv3EdmTopologyExternal) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmTopologyExternal.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmTopologyExternal proto.InternalMessageInfo

func (m *Ospfv3EdmTopologyExternal) GetRouteTopology() *Ospfv3EdmTopology {
	if m != nil {
		return m.RouteTopology
	}
	return nil
}

func (m *Ospfv3EdmTopologyExternal) GetRoutePath() []*Ospfv3EdmTopPath {
	if m != nil {
		return m.RoutePath
	}
	return nil
}

func (m *Ospfv3EdmTopologyExternal) GetRouteExtendedCommunity() *Ospfv3ShRouteExtendedComm {
	if m != nil {
		return m.RouteExtendedCommunity
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmTopologyExternal_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology_external_KEYS")
	proto.RegisterType((*Ospfv3EdmTopology)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology")
	proto.RegisterType((*Ospfv3ShBackupPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_sh_backup_path")
	proto.RegisterType((*Ospfv3EdmTopPath)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_top_path")
	proto.RegisterType((*Ospfv3ShRouteExtendedComm)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_sh_route_extended_comm")
	proto.RegisterType((*Ospfv3EdmTopologyExternal)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.external_routes.external_route.ospfv3_edm_topology_external")
}

func init() { proto.RegisterFile("ospfv3_edm_topology_external.proto", fileDescriptor_b36ab0046955ae4e) }

var fileDescriptor_b36ab0046955ae4e = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0x2b, 0x35,
	0x14, 0xd6, 0xdc, 0xf4, 0xb6, 0x8d, 0xdb, 0xb9, 0xb4, 0xae, 0x28, 0xd3, 0xd2, 0x54, 0xe9, 0xa0,
	0x4a, 0x59, 0x40, 0x16, 0x2d, 0x62, 0x03, 0x08, 0x4a, 0x5b, 0x89, 0xf0, 0x53, 0xaa, 0x69, 0x85,
	0xc4, 0x06, 0xcb, 0x1d, 0x3b, 0x8d, 0x61, 0x66, 0x3c, 0xb2, 0x9d, 0x90, 0x2c, 0x78, 0x01, 0x96,
	0x48, 0x6c, 0x58, 0x20, 0xb1, 0xe1, 0x05, 0x90, 0xe0, 0xf5, 0x90, 0x8f, 0x3d, 0xcd, 0xa4, 0x19,
	0xba, 0xbb, 0xdd, 0x44, 0xe3, 0xef, 0x7c, 0xe7, 0xf8, 0x3b, 0x7f, 0x56, 0x50, 0x2c, 0x75, 0x39,
	0x9c, 0x9c, 0x12, 0xce, 0x72, 0x62, 0x64, 0x29, 0x33, 0x79, 0x3f, 0x23, 0x7c, 0x6a, 0xb8, 0x2a,
	0x68, 0xd6, 0x2f, 0x95, 0x34, 0x12, 0x7f, 0x9f, 0x0a, 0x9d, 0x4a, 0x22, 0xa4, 0x26, 0x53, 0x45,
	0x44, 0x39, 0xf9, 0x80, 0x78, 0x2f, 0x59, 0x72, 0xd5, 0x77, 0xdf, 0x96, 0x9b, 0x72, 0xad, 0xb9,
	0xae, 0xbe, 0xfa, 0x13, 0x35, 0x84, 0x9f, 0x7e, 0x15, 0x8f, 0x28, 0x39, 0x36, 0x5c, 0x3f, 0x3a,
	0xc7, 0x7f, 0x04, 0xe8, 0xe8, 0x29, 0x19, 0xe4, 0xcb, 0xcb, 0xef, 0x6e, 0xf0, 0x11, 0xda, 0xf4,
	0x81, 0x49, 0x41, 0x73, 0x1e, 0x05, 0xdd, 0xa0, 0xd7, 0x4e, 0x36, 0x3c, 0x76, 0x45, 0x73, 0x8e,
	0xf7, 0xd0, 0xfa, 0x44, 0x0d, 0x9d, 0xf9, 0x05, 0x98, 0xd7, 0x26, 0x6a, 0x08, 0xa6, 0x5d, 0xb4,
	0x5a, 0x2a, 0x3e, 0x14, 0xd3, 0xa8, 0x05, 0x06, 0x7f, 0xc2, 0xef, 0xa0, 0xd0, 0x7d, 0x91, 0x8c,
	0x17, 0xf7, 0x66, 0x14, 0xad, 0x74, 0x83, 0x5e, 0x98, 0x6c, 0x3a, 0xf0, 0x2b, 0xc0, 0xe2, 0xdf,
	0x02, 0xb4, 0xd3, 0x20, 0xd0, 0xde, 0x07, 0x19, 0x10, 0xc1, 0xbc, 0x9c, 0x35, 0x38, 0x0f, 0x18,
	0x3e, 0x46, 0xaf, 0x9c, 0x89, 0x09, 0x6d, 0x68, 0x91, 0x3a, 0x41, 0x61, 0x12, 0x02, 0x7a, 0xe1,
	0x41, 0xdc, 0x41, 0xc8, 0xd1, 0x52, 0xa9, 0x0d, 0x48, 0x0b, 0x93, 0x36, 0x20, 0xe7, 0x52, 0x9b,
	0xb9, 0xd9, 0xcc, 0x4a, 0xee, 0xa5, 0x39, 0xf3, 0xed, 0xac, 0xe4, 0xf1, 0x5f, 0x2d, 0xf4, 0xa6,
	0xd7, 0xa5, 0x47, 0xe4, 0x8e, 0xa6, 0x3f, 0x8e, 0x4b, 0x52, 0x52, 0x33, 0xc2, 0x1f, 0xa3, 0xb7,
	0xfd, 0xd1, 0x0b, 0x2c, 0x0c, 0x57, 0x43, 0x9a, 0xf2, 0x7a, 0xed, 0x22, 0x47, 0x49, 0x40, 0x72,
	0x45, 0x80, 0x6a, 0x7d, 0x8a, 0x3a, 0x0b, 0xee, 0x05, 0x9f, 0x1a, 0x32, 0x92, 0x25, 0xa1, 0x8c,
	0x29, 0xae, 0xb5, 0xaf, 0xee, 0x5e, 0x2d, 0xc0, 0x15, 0x9f, 0x9a, 0xcf, 0x65, 0x79, 0xe6, 0x08,
	0xb8, 0x8f, 0x76, 0x16, 0x22, 0x68, 0x39, 0x56, 0x29, 0xf7, 0xc5, 0xdf, 0xae, 0xf9, 0xdd, 0x80,
	0xc1, 0xf6, 0xc1, 0xf3, 0x73, 0x6e, 0x94, 0x48, 0xab, 0x3e, 0x38, 0xf0, 0x6b, 0xc0, 0xdc, 0x08,
	0x88, 0x9c, 0xaa, 0x19, 0x64, 0x19, 0xbd, 0xec, 0x06, 0xbd, 0x75, 0x3b, 0x02, 0x80, 0x5d, 0xdb,
	0xc4, 0xdf, 0x45, 0x38, 0x13, 0x05, 0x27, 0x29, 0x55, 0xcc, 0xd6, 0xfe, 0x07, 0x29, 0x0a, 0x13,
	0xad, 0x02, 0x71, 0xcb, 0x5a, 0xce, 0xa9, 0x62, 0x17, 0x1e, 0xc7, 0x87, 0x08, 0x31, 0xf9, 0x53,
	0xa1, 0x8d, 0xe2, 0x34, 0x8f, 0xd6, 0x80, 0x55, 0x43, 0xec, 0x85, 0x85, 0x64, 0x9c, 0xd8, 0x3d,
	0xe0, 0xa9, 0x89, 0xd6, 0xdd, 0x85, 0x16, 0xbb, 0x76, 0x90, 0x15, 0xae, 0x55, 0x76, 0x3f, 0xbf,
	0xab, 0x0d, 0x9c, 0x4d, 0x0b, 0x56, 0xf7, 0xc4, 0x7f, 0xbf, 0x78, 0x3c, 0x40, 0xae, 0x4d, 0xc7,
	0xe8, 0x55, 0x63, 0x67, 0x42, 0xb1, 0xd0, 0x8e, 0xf7, 0xd0, 0x8e, 0xab, 0xa2, 0x75, 0x7a, 0x68,
	0x86, 0x6f, 0xc2, 0x16, 0x98, 0x6c, 0xf2, 0xbe, 0x05, 0x38, 0x46, 0x61, 0x8d, 0x2e, 0x98, 0x9f,
	0xab, 0x8d, 0x07, 0xe2, 0x80, 0xe1, 0x3f, 0x03, 0xb4, 0xed, 0x48, 0xb5, 0xb1, 0x81, 0xa2, 0x6f,
	0x9c, 0x8c, 0xfb, 0xaf, 0x77, 0xe1, 0xfb, 0x8d, 0x33, 0x9b, 0xbc, 0x01, 0xc6, 0xcf, 0x00, 0xb1,
	0x2a, 0xe3, 0x7f, 0x5a, 0xa8, 0x33, 0xa7, 0x3a, 0xb5, 0x36, 0x52, 0xc1, 0x38, 0x23, 0xa9, 0xcc,
	0x73, 0xfc, 0x05, 0x8a, 0x17, 0x80, 0x71, 0x21, 0xcc, 0x8c, 0x30, 0x99, 0x53, 0x51, 0x10, 0xc1,
	0xc8, 0x84, 0x66, 0x63, 0x5b, 0xd3, 0x56, 0x2f, 0x4c, 0x0e, 0x2b, 0xe6, 0x79, 0x45, 0xbc, 0x00,
	0xde, 0x80, 0x7d, 0x6b, 0x59, 0xcd, 0xb1, 0xb2, 0x5a, 0x30, 0xd8, 0x41, 0xb7, 0xc5, 0xcb, 0xb1,
	0xb2, 0x2a, 0x98, 0x5d, 0x4c, 0xfc, 0x21, 0xda, 0x6f, 0xd0, 0x45, 0x15, 0xa7, 0xf3, 0x76, 0xbc,
	0xb5, 0x14, 0xe3, 0x4c, 0x71, 0x3a, 0x60, 0xf8, 0x13, 0x74, 0xd0, 0xe0, 0x0c, 0xe9, 0x2b, 0xeb,
	0xbe, 0xe2, 0x76, 0x6f, 0xc9, 0x1d, 0xd6, 0x49, 0x0d, 0x18, 0x3e, 0x43, 0x9d, 0xff, 0x0b, 0xe0,
	0x92, 0x78, 0x09, 0x02, 0xf6, 0x9b, 0x23, 0x40, 0x02, 0x1f, 0x35, 0x26, 0x20, 0x4b, 0x23, 0x64,
	0xa1, 0x61, 0x9d, 0xc2, 0x24, 0x5a, 0xf2, 0xff, 0xc6, 0xd9, 0xe3, 0x5f, 0x56, 0xd0, 0xc1, 0x53,
	0x0f, 0x3a, 0xfe, 0x3d, 0xa8, 0x9e, 0xc7, 0xca, 0x16, 0x9d, 0xc0, 0xe8, 0xe9, 0x67, 0x1a, 0xbd,
	0xba, 0x2c, 0xff, 0x26, 0xdf, 0x56, 0xaf, 0xfa, 0xaf, 0x41, 0xf5, 0xea, 0xc2, 0x4e, 0x9c, 0x76,
	0x5b, 0xcf, 0x2f, 0xcc, 0x6d, 0x44, 0xfb, 0x61, 0x63, 0xf1, 0xbf, 0x01, 0x8a, 0x1a, 0x36, 0x00,
	0xca, 0x1e, 0xbd, 0x0f, 0xb5, 0xfb, 0xf9, 0xf9, 0xd6, 0xb6, 0x41, 0x49, 0xb2, 0x0b, 0xe0, 0xe5,
	0xe3, 0x99, 0xb8, 0x5b, 0x85, 0x3f, 0x11, 0xa7, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x36, 0xb6,
	0xdc, 0x7a, 0x6a, 0x08, 0x00, 0x00,
}
