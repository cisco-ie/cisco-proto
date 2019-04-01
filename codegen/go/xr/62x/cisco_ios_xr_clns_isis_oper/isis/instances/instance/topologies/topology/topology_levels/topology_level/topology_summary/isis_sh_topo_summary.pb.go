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
// source: isis_sh_topo_summary.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_topology_summary

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

type IsisShTopoSummary_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	Level                string   `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTopoSummary_KEYS) Reset()         { *m = IsisShTopoSummary_KEYS{} }
func (m *IsisShTopoSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShTopoSummary_KEYS) ProtoMessage()    {}
func (*IsisShTopoSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9c80d2dfbf2207, []int{0}
}

func (m *IsisShTopoSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTopoSummary_KEYS.Unmarshal(m, b)
}
func (m *IsisShTopoSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTopoSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShTopoSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTopoSummary_KEYS.Merge(m, src)
}
func (m *IsisShTopoSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShTopoSummary_KEYS.Size(m)
}
func (m *IsisShTopoSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTopoSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTopoSummary_KEYS proto.InternalMessageInfo

func (m *IsisShTopoSummary_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTopoSummary_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShTopoSummary_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShTopoSummary_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShTopoSummary_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type IsisShTopoSummaryNodeStats struct {
	ReachableNodeCount              uint32   `protobuf:"varint,1,opt,name=reachable_node_count,json=reachableNodeCount,proto3" json:"reachable_node_count,omitempty"`
	UnreachableNodeCount            uint32   `protobuf:"varint,2,opt,name=unreachable_node_count,json=unreachableNodeCount,proto3" json:"unreachable_node_count,omitempty"`
	UnreachableParticipantNodeCount uint32   `protobuf:"varint,3,opt,name=unreachable_participant_node_count,json=unreachableParticipantNodeCount,proto3" json:"unreachable_participant_node_count,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *IsisShTopoSummaryNodeStats) Reset()         { *m = IsisShTopoSummaryNodeStats{} }
func (m *IsisShTopoSummaryNodeStats) String() string { return proto.CompactTextString(m) }
func (*IsisShTopoSummaryNodeStats) ProtoMessage()    {}
func (*IsisShTopoSummaryNodeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9c80d2dfbf2207, []int{1}
}

func (m *IsisShTopoSummaryNodeStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTopoSummaryNodeStats.Unmarshal(m, b)
}
func (m *IsisShTopoSummaryNodeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTopoSummaryNodeStats.Marshal(b, m, deterministic)
}
func (m *IsisShTopoSummaryNodeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTopoSummaryNodeStats.Merge(m, src)
}
func (m *IsisShTopoSummaryNodeStats) XXX_Size() int {
	return xxx_messageInfo_IsisShTopoSummaryNodeStats.Size(m)
}
func (m *IsisShTopoSummaryNodeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTopoSummaryNodeStats.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTopoSummaryNodeStats proto.InternalMessageInfo

func (m *IsisShTopoSummaryNodeStats) GetReachableNodeCount() uint32 {
	if m != nil {
		return m.ReachableNodeCount
	}
	return 0
}

func (m *IsisShTopoSummaryNodeStats) GetUnreachableNodeCount() uint32 {
	if m != nil {
		return m.UnreachableNodeCount
	}
	return 0
}

func (m *IsisShTopoSummaryNodeStats) GetUnreachableParticipantNodeCount() uint32 {
	if m != nil {
		return m.UnreachableParticipantNodeCount
	}
	return 0
}

type IsisShTopoSummary struct {
	RouterNodeCount      *IsisShTopoSummaryNodeStats `protobuf:"bytes,50,opt,name=router_node_count,json=routerNodeCount,proto3" json:"router_node_count,omitempty"`
	PseudonodeNodeCount  *IsisShTopoSummaryNodeStats `protobuf:"bytes,51,opt,name=pseudonode_node_count,json=pseudonodeNodeCount,proto3" json:"pseudonode_node_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *IsisShTopoSummary) Reset()         { *m = IsisShTopoSummary{} }
func (m *IsisShTopoSummary) String() string { return proto.CompactTextString(m) }
func (*IsisShTopoSummary) ProtoMessage()    {}
func (*IsisShTopoSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b9c80d2dfbf2207, []int{2}
}

func (m *IsisShTopoSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTopoSummary.Unmarshal(m, b)
}
func (m *IsisShTopoSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTopoSummary.Marshal(b, m, deterministic)
}
func (m *IsisShTopoSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTopoSummary.Merge(m, src)
}
func (m *IsisShTopoSummary) XXX_Size() int {
	return xxx_messageInfo_IsisShTopoSummary.Size(m)
}
func (m *IsisShTopoSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTopoSummary.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTopoSummary proto.InternalMessageInfo

func (m *IsisShTopoSummary) GetRouterNodeCount() *IsisShTopoSummaryNodeStats {
	if m != nil {
		return m.RouterNodeCount
	}
	return nil
}

func (m *IsisShTopoSummary) GetPseudonodeNodeCount() *IsisShTopoSummaryNodeStats {
	if m != nil {
		return m.PseudonodeNodeCount
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShTopoSummary_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.topology_summary.isis_sh_topo_summary_KEYS")
	proto.RegisterType((*IsisShTopoSummaryNodeStats)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.topology_summary.isis_sh_topo_summary_node_stats")
	proto.RegisterType((*IsisShTopoSummary)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.topology_summary.isis_sh_topo_summary")
}

func init() { proto.RegisterFile("isis_sh_topo_summary.proto", fileDescriptor_2b9c80d2dfbf2207) }

var fileDescriptor_2b9c80d2dfbf2207 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xcf, 0x4e, 0x2a, 0x31,
	0x14, 0xc6, 0x53, 0xb8, 0xc0, 0xbd, 0xe7, 0x4a, 0x8c, 0xe3, 0xa8, 0xe0, 0x06, 0x32, 0x6c, 0x58,
	0x4d, 0x0c, 0xf8, 0x06, 0xc6, 0x15, 0x09, 0x31, 0xb8, 0x72, 0xd5, 0x94, 0x52, 0xa4, 0xc9, 0x4c,
	0x3b, 0xe9, 0xe9, 0xa8, 0xbc, 0x89, 0x4f, 0xa1, 0x4f, 0xe1, 0xd2, 0xad, 0xcf, 0x63, 0xe8, 0x30,
	0x7f, 0x62, 0x26, 0x71, 0xa9, 0xbb, 0xd3, 0xf3, 0xfb, 0xce, 0xd7, 0xef, 0xcc, 0x1f, 0x38, 0x97,
	0x28, 0x91, 0xe2, 0x86, 0x5a, 0x9d, 0x68, 0x8a, 0x69, 0x1c, 0x33, 0xb3, 0x0d, 0x13, 0xa3, 0xad,
	0xf6, 0x1e, 0xb9, 0x44, 0xae, 0xa9, 0xd4, 0x48, 0x9f, 0x0c, 0xe5, 0x91, 0x42, 0xea, 0xd4, 0x3a,
	0x11, 0x26, 0xdc, 0x55, 0xa1, 0x54, 0x68, 0x99, 0xe2, 0xa2, 0xac, 0xc2, 0x9d, 0x4f, 0xa4, 0xef,
	0xa5, 0xc0, 0xbc, 0xdc, 0x16, 0x05, 0x8d, 0xc4, 0x83, 0x88, 0xf0, 0xcb, 0xb9, 0x3c, 0xee, 0xaf,
	0x0f, 0x5e, 0x09, 0xf4, 0xeb, 0x72, 0xd1, 0xd9, 0xf5, 0xdd, 0xad, 0x37, 0x82, 0x6e, 0x7e, 0x1b,
	0x55, 0x2c, 0x16, 0x3d, 0x32, 0x24, 0xe3, 0x7f, 0x8b, 0x83, 0xbc, 0x39, 0x67, 0xb1, 0xf0, 0xce,
	0xa0, 0xc3, 0xd6, 0x19, 0x6e, 0x38, 0xdc, 0x66, 0x6b, 0x07, 0xfa, 0xf0, 0x17, 0x73, 0xd2, 0x74,
	0xa4, 0x83, 0x7b, 0x34, 0x82, 0x6e, 0x11, 0xc5, 0xf1, 0x3f, 0x99, 0x71, 0xde, 0x74, 0x22, 0x1f,
	0x5a, 0x2e, 0x75, 0xaf, 0xe5, 0x60, 0x76, 0x08, 0x3e, 0x08, 0x0c, 0x6a, 0x13, 0x2b, 0xbd, 0x12,
	0x14, 0x2d, 0xb3, 0xe8, 0x5d, 0x80, 0x6f, 0x04, 0xe3, 0x1b, 0xb6, 0x8c, 0x44, 0xd6, 0xe7, 0x3a,
	0x55, 0xd6, 0xc5, 0xef, 0x2e, 0xbc, 0x82, 0xcd, 0xf5, 0x4a, 0x5c, 0xed, 0x88, 0x77, 0x09, 0xa7,
	0xa9, 0xaa, 0x9d, 0x69, 0xb8, 0x19, 0xbf, 0x42, 0xcb, 0xa9, 0x19, 0x04, 0xd5, 0xa9, 0x84, 0x19,
	0x2b, 0xb9, 0x4c, 0x98, 0xb2, 0x55, 0x87, 0xa6, 0x73, 0x18, 0x54, 0x94, 0x37, 0xa5, 0xb0, 0x30,
	0x0b, 0x5e, 0x9a, 0xe0, 0xd7, 0x2d, 0xe6, 0xbd, 0x11, 0x38, 0x32, 0x3a, 0xb5, 0xc2, 0x54, 0x5d,
	0x27, 0x43, 0x32, 0xfe, 0x3f, 0x79, 0x26, 0xe1, 0x0f, 0x7d, 0x3a, 0xe1, 0x37, 0x2f, 0x61, 0x71,
	0x98, 0x65, 0x2e, 0x9f, 0xd6, 0x3b, 0x81, 0x93, 0x04, 0x45, 0xba, 0xd2, 0x4e, 0x55, 0xd9, 0x65,
	0xfa, 0xdb, 0x77, 0x39, 0x2e, 0x73, 0x17, 0xfb, 0x2c, 0xdb, 0xee, 0xdf, 0x9d, 0x7e, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x14, 0x85, 0x35, 0xde, 0xd9, 0x03, 0x00, 0x00,
}
