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
// source: isis_sh_frr_summary.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_frr_summary

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

type IsisShFrrSummary_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,3,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShFrrSummary_KEYS) Reset()         { *m = IsisShFrrSummary_KEYS{} }
func (m *IsisShFrrSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShFrrSummary_KEYS) ProtoMessage()    {}
func (*IsisShFrrSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c98dc24990f20193, []int{0}
}

func (m *IsisShFrrSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShFrrSummary_KEYS.Unmarshal(m, b)
}
func (m *IsisShFrrSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShFrrSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShFrrSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShFrrSummary_KEYS.Merge(m, src)
}
func (m *IsisShFrrSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShFrrSummary_KEYS.Size(m)
}
func (m *IsisShFrrSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShFrrSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShFrrSummary_KEYS proto.InternalMessageInfo

func (m *IsisShFrrSummary_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShFrrSummary_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShFrrSummary_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShFrrSummary_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

type IsisPerPriorityCounts struct {
	Critical             uint32   `protobuf:"varint,1,opt,name=critical,proto3" json:"critical,omitempty"`
	High                 uint32   `protobuf:"varint,2,opt,name=high,proto3" json:"high,omitempty"`
	Medium               uint32   `protobuf:"varint,3,opt,name=medium,proto3" json:"medium,omitempty"`
	Low                  uint32   `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisPerPriorityCounts) Reset()         { *m = IsisPerPriorityCounts{} }
func (m *IsisPerPriorityCounts) String() string { return proto.CompactTextString(m) }
func (*IsisPerPriorityCounts) ProtoMessage()    {}
func (*IsisPerPriorityCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_c98dc24990f20193, []int{1}
}

func (m *IsisPerPriorityCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisPerPriorityCounts.Unmarshal(m, b)
}
func (m *IsisPerPriorityCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisPerPriorityCounts.Marshal(b, m, deterministic)
}
func (m *IsisPerPriorityCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisPerPriorityCounts.Merge(m, src)
}
func (m *IsisPerPriorityCounts) XXX_Size() int {
	return xxx_messageInfo_IsisPerPriorityCounts.Size(m)
}
func (m *IsisPerPriorityCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisPerPriorityCounts.DiscardUnknown(m)
}

var xxx_messageInfo_IsisPerPriorityCounts proto.InternalMessageInfo

func (m *IsisPerPriorityCounts) GetCritical() uint32 {
	if m != nil {
		return m.Critical
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetHigh() uint32 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetMedium() uint32 {
	if m != nil {
		return m.Medium
	}
	return 0
}

func (m *IsisPerPriorityCounts) GetLow() uint32 {
	if m != nil {
		return m.Low
	}
	return 0
}

type IsisShFrrLevelSummary struct {
	AllPathsProtected    *IsisPerPriorityCounts `protobuf:"bytes,1,opt,name=all_paths_protected,json=allPathsProtected,proto3" json:"all_paths_protected,omitempty"`
	SomePathsProtected   *IsisPerPriorityCounts `protobuf:"bytes,2,opt,name=some_paths_protected,json=somePathsProtected,proto3" json:"some_paths_protected,omitempty"`
	Unprotected          *IsisPerPriorityCounts `protobuf:"bytes,3,opt,name=unprotected,proto3" json:"unprotected,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IsisShFrrLevelSummary) Reset()         { *m = IsisShFrrLevelSummary{} }
func (m *IsisShFrrLevelSummary) String() string { return proto.CompactTextString(m) }
func (*IsisShFrrLevelSummary) ProtoMessage()    {}
func (*IsisShFrrLevelSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c98dc24990f20193, []int{2}
}

func (m *IsisShFrrLevelSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShFrrLevelSummary.Unmarshal(m, b)
}
func (m *IsisShFrrLevelSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShFrrLevelSummary.Marshal(b, m, deterministic)
}
func (m *IsisShFrrLevelSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShFrrLevelSummary.Merge(m, src)
}
func (m *IsisShFrrLevelSummary) XXX_Size() int {
	return xxx_messageInfo_IsisShFrrLevelSummary.Size(m)
}
func (m *IsisShFrrLevelSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShFrrLevelSummary.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShFrrLevelSummary proto.InternalMessageInfo

func (m *IsisShFrrLevelSummary) GetAllPathsProtected() *IsisPerPriorityCounts {
	if m != nil {
		return m.AllPathsProtected
	}
	return nil
}

func (m *IsisShFrrLevelSummary) GetSomePathsProtected() *IsisPerPriorityCounts {
	if m != nil {
		return m.SomePathsProtected
	}
	return nil
}

func (m *IsisShFrrLevelSummary) GetUnprotected() *IsisPerPriorityCounts {
	if m != nil {
		return m.Unprotected
	}
	return nil
}

type IsisShFrrSummary struct {
	Level1Prefixes       *IsisShFrrLevelSummary `protobuf:"bytes,50,opt,name=level1_prefixes,json=level1Prefixes,proto3" json:"level1_prefixes,omitempty"`
	Level2Prefixes       *IsisShFrrLevelSummary `protobuf:"bytes,51,opt,name=level2_prefixes,json=level2Prefixes,proto3" json:"level2_prefixes,omitempty"`
	UnreachablePrefixes  *IsisPerPriorityCounts `protobuf:"bytes,52,opt,name=unreachable_prefixes,json=unreachablePrefixes,proto3" json:"unreachable_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IsisShFrrSummary) Reset()         { *m = IsisShFrrSummary{} }
func (m *IsisShFrrSummary) String() string { return proto.CompactTextString(m) }
func (*IsisShFrrSummary) ProtoMessage()    {}
func (*IsisShFrrSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c98dc24990f20193, []int{3}
}

func (m *IsisShFrrSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShFrrSummary.Unmarshal(m, b)
}
func (m *IsisShFrrSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShFrrSummary.Marshal(b, m, deterministic)
}
func (m *IsisShFrrSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShFrrSummary.Merge(m, src)
}
func (m *IsisShFrrSummary) XXX_Size() int {
	return xxx_messageInfo_IsisShFrrSummary.Size(m)
}
func (m *IsisShFrrSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShFrrSummary.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShFrrSummary proto.InternalMessageInfo

func (m *IsisShFrrSummary) GetLevel1Prefixes() *IsisShFrrLevelSummary {
	if m != nil {
		return m.Level1Prefixes
	}
	return nil
}

func (m *IsisShFrrSummary) GetLevel2Prefixes() *IsisShFrrLevelSummary {
	if m != nil {
		return m.Level2Prefixes
	}
	return nil
}

func (m *IsisShFrrSummary) GetUnreachablePrefixes() *IsisPerPriorityCounts {
	if m != nil {
		return m.UnreachablePrefixes
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShFrrSummary_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.frr_summary.isis_sh_frr_summary_KEYS")
	proto.RegisterType((*IsisPerPriorityCounts)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.frr_summary.isis_per_priority_counts")
	proto.RegisterType((*IsisShFrrLevelSummary)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.frr_summary.isis_sh_frr_level_summary")
	proto.RegisterType((*IsisShFrrSummary)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.frr_summary.isis_sh_frr_summary")
}

func init() { proto.RegisterFile("isis_sh_frr_summary.proto", fileDescriptor_c98dc24990f20193) }

var fileDescriptor_c98dc24990f20193 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0xd4, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x07, 0x70, 0xb9, 0xa9, 0xb6, 0x65, 0x4a, 0xf8, 0xf0, 0x56, 0xb0, 0xe5, 0x84, 0xb6, 0x17,
	0x4e, 0x91, 0x48, 0x79, 0x05, 0x4e, 0x48, 0x68, 0x95, 0x9e, 0x38, 0x59, 0xae, 0x3b, 0x69, 0x2c,
	0xf9, 0x23, 0xb2, 0x1d, 0xe8, 0x3e, 0x05, 0x47, 0xb8, 0x20, 0x78, 0x40, 0x24, 0x5e, 0x01, 0xc5,
	0x89, 0x9b, 0x05, 0x96, 0x63, 0x95, 0xdb, 0xd8, 0x7f, 0x6b, 0xfc, 0x93, 0x67, 0x37, 0x70, 0x26,
	0xbd, 0xf4, 0xcc, 0x37, 0xac, 0x76, 0x8e, 0xf9, 0x4e, 0x6b, 0xee, 0xb6, 0x45, 0xeb, 0x6c, 0xb0,
	0xf4, 0x52, 0x48, 0x2f, 0x2c, 0x93, 0xd6, 0xb3, 0x5b, 0xc7, 0x84, 0x32, 0x9e, 0xc5, 0xc3, 0xb6,
	0x45, 0x57, 0xf4, 0x55, 0x21, 0x8d, 0x0f, 0xdc, 0x08, 0x9c, 0xaa, 0x22, 0xd8, 0xd6, 0x2a, 0x7b,
	0x23, 0xd1, 0xa7, 0x72, 0x5b, 0xec, 0xb4, 0x5e, 0x7f, 0x25, 0xb0, 0xda, 0x73, 0x25, 0x7b, 0xf7,
	0xf6, 0xc3, 0x25, 0x3d, 0x87, 0x3c, 0x35, 0x62, 0x86, 0x6b, 0x5c, 0x91, 0x97, 0xe4, 0xd5, 0x83,
	0xea, 0x61, 0xda, 0x7c, 0xcf, 0x35, 0xd2, 0xe7, 0x70, 0xc4, 0xeb, 0x21, 0x3e, 0x88, 0xf1, 0x82,
	0xd7, 0x31, 0x38, 0x83, 0x63, 0x9f, 0x92, 0x2c, 0x26, 0x47, 0x7e, 0x8c, 0xce, 0x21, 0x4f, 0x9a,
	0x21, 0x3f, 0x1c, 0x1a, 0xa7, 0xcd, 0xfe, 0xd0, 0x3a, 0x8c, 0xb2, 0x16, 0x1d, 0x6b, 0x9d, 0xb4,
	0x4e, 0x86, 0x2d, 0x13, 0xb6, 0x33, 0xc1, 0xd3, 0x17, 0x70, 0x2c, 0x9c, 0x0c, 0x52, 0x70, 0x15,
	0x51, 0x79, 0x75, 0xb7, 0xa6, 0x14, 0x0e, 0x1b, 0x79, 0xd3, 0x44, 0x4d, 0x5e, 0xc5, 0x9a, 0x3e,
	0x83, 0x85, 0xc6, 0x6b, 0xd9, 0xe9, 0x28, 0xc9, 0xab, 0x71, 0x45, 0x9f, 0x40, 0xa6, 0xec, 0xa7,
	0x78, 0x7d, 0x5e, 0xf5, 0xe5, 0xfa, 0x57, 0xf6, 0xe7, 0x0c, 0x14, 0x7e, 0x44, 0x95, 0x9e, 0x85,
	0x7e, 0x23, 0xb0, 0xe4, 0x4a, 0xb1, 0x96, 0x87, 0xc6, 0xb3, 0x7e, 0x2e, 0x28, 0x02, 0x5e, 0x47,
	0xc3, 0x49, 0xa9, 0x8b, 0x7b, 0x18, 0x51, 0xf1, 0xbf, 0x47, 0xa8, 0x9e, 0x72, 0xa5, 0x36, 0x3d,
	0x64, 0x93, 0x1c, 0xf4, 0x3b, 0x81, 0x53, 0x6f, 0x35, 0xfe, 0x03, 0x3c, 0x98, 0x03, 0x48, 0x7b,
	0xca, 0x5f, 0xc2, 0xcf, 0x04, 0x4e, 0x3a, 0x33, 0xc1, 0xb2, 0x39, 0x60, 0xbb, 0x82, 0xf5, 0xcf,
	0x0c, 0x96, 0x7b, 0xfe, 0x02, 0xf4, 0x0b, 0x81, 0xc7, 0x71, 0xfa, 0xaf, 0x59, 0xeb, 0xb0, 0x96,
	0xb7, 0xe8, 0x57, 0x65, 0xd4, 0x9a, 0xfb, 0xd3, 0xee, 0xfb, 0xd5, 0x55, 0x8f, 0x06, 0xc6, 0x66,
	0x54, 0x4c, 0xb2, 0x72, 0x92, 0x5d, 0xcc, 0x28, 0x2b, 0xef, 0x64, 0x3f, 0x08, 0x9c, 0x76, 0xc6,
	0x21, 0x17, 0x0d, 0xbf, 0x52, 0x38, 0xf1, 0xde, 0xcc, 0x31, 0xe6, 0xe5, 0x0e, 0x25, 0x11, 0xaf,
	0x16, 0xf1, 0x6b, 0x7a, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x45, 0x5c, 0x5d, 0x65, 0x6a, 0x05,
	0x00, 0x00,
}
