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
// source: plmgr_pmap_str_targets_bg.proto

package cisco_ios_xr_infra_policymgr_oper_policy_manager_nodes_node_target_policy_map_types_target_policy_map_type_policy_map_targets_policy_map_target

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

type PlmgrPmapStrTargetsBg_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	PolicyMapName        string   `protobuf:"bytes,3,opt,name=policy_map_name,json=policyMapName,proto3" json:"policy_map_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmgrPmapStrTargetsBg_KEYS) Reset()         { *m = PlmgrPmapStrTargetsBg_KEYS{} }
func (m *PlmgrPmapStrTargetsBg_KEYS) String() string { return proto.CompactTextString(m) }
func (*PlmgrPmapStrTargetsBg_KEYS) ProtoMessage()    {}
func (*PlmgrPmapStrTargetsBg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b4b8a59f0baafb, []int{0}
}

func (m *PlmgrPmapStrTargetsBg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS.Unmarshal(m, b)
}
func (m *PlmgrPmapStrTargetsBg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS.Marshal(b, m, deterministic)
}
func (m *PlmgrPmapStrTargetsBg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS.Merge(m, src)
}
func (m *PlmgrPmapStrTargetsBg_KEYS) XXX_Size() int {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS.Size(m)
}
func (m *PlmgrPmapStrTargetsBg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrPmapStrTargetsBg_KEYS proto.InternalMessageInfo

func (m *PlmgrPmapStrTargetsBg_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PlmgrPmapStrTargetsBg_KEYS) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlmgrPmapStrTargetsBg_KEYS) GetPolicyMapName() string {
	if m != nil {
		return m.PolicyMapName
	}
	return ""
}

type PlmgrPmapStrTargetsBg struct {
	Targets              []string `protobuf:"bytes,50,rep,name=targets,proto3" json:"targets,omitempty"`
	ParentTargets        []string `protobuf:"bytes,51,rep,name=parent_targets,json=parentTargets,proto3" json:"parent_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmgrPmapStrTargetsBg) Reset()         { *m = PlmgrPmapStrTargetsBg{} }
func (m *PlmgrPmapStrTargetsBg) String() string { return proto.CompactTextString(m) }
func (*PlmgrPmapStrTargetsBg) ProtoMessage()    {}
func (*PlmgrPmapStrTargetsBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1b4b8a59f0baafb, []int{1}
}

func (m *PlmgrPmapStrTargetsBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg.Unmarshal(m, b)
}
func (m *PlmgrPmapStrTargetsBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg.Marshal(b, m, deterministic)
}
func (m *PlmgrPmapStrTargetsBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrPmapStrTargetsBg.Merge(m, src)
}
func (m *PlmgrPmapStrTargetsBg) XXX_Size() int {
	return xxx_messageInfo_PlmgrPmapStrTargetsBg.Size(m)
}
func (m *PlmgrPmapStrTargetsBg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrPmapStrTargetsBg.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrPmapStrTargetsBg proto.InternalMessageInfo

func (m *PlmgrPmapStrTargetsBg) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *PlmgrPmapStrTargetsBg) GetParentTargets() []string {
	if m != nil {
		return m.ParentTargets
	}
	return nil
}

func init() {
	proto.RegisterType((*PlmgrPmapStrTargetsBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.target_policy_map_types.target_policy_map_type.policy_map_targets.policy_map_target.plmgr_pmap_str_targets_bg_KEYS")
	proto.RegisterType((*PlmgrPmapStrTargetsBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.target_policy_map_types.target_policy_map_type.policy_map_targets.policy_map_target.plmgr_pmap_str_targets_bg")
}

func init() { proto.RegisterFile("plmgr_pmap_str_targets_bg.proto", fileDescriptor_f1b4b8a59f0baafb) }

var fileDescriptor_f1b4b8a59f0baafb = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0x2b, 0xea, 0x06, 0xaa, 0x90, 0x53, 0x44, 0xd0, 0x52, 0x50, 0x7a, 0xca, 0xc1,
	0xfe, 0x06, 0x4f, 0xa2, 0x87, 0xea, 0x45, 0x10, 0x86, 0xe9, 0x3a, 0x2e, 0x0b, 0x4d, 0x32, 0x4c,
	0x72, 0x70, 0x7f, 0x85, 0x7f, 0x59, 0x92, 0x58, 0x29, 0xc8, 0x5e, 0x42, 0xde, 0x97, 0xf7, 0xf2,
	0x86, 0x51, 0x37, 0xbc, 0x73, 0xbd, 0x00, 0x3b, 0x64, 0x88, 0x49, 0x20, 0xa1, 0xf4, 0x94, 0x22,
	0x6c, 0x7b, 0xcb, 0x12, 0x52, 0xd0, 0xdf, 0x4d, 0x37, 0xc4, 0x2e, 0xc0, 0x10, 0x22, 0x7c, 0x09,
	0x0c, 0xfe, 0x53, 0x10, 0x38, 0xec, 0x86, 0x6e, 0xcc, 0xc1, 0xc0, 0x24, 0xb6, 0x4a, 0x70, 0xe8,
	0xb1, 0x27, 0xb1, 0x3e, 0x7c, 0x50, 0x2c, 0xa7, 0xad, 0xdf, 0xc1, 0x9f, 0x81, 0x21, 0x8d, 0x4c,
	0x71, 0x82, 0xdb, 0x43, 0x5d, 0x07, 0xf9, 0x8f, 0x96, 0xa3, 0xba, 0x9e, 0x1c, 0x1a, 0x1e, 0x1f,
	0xde, 0x5e, 0xf4, 0x95, 0x6a, 0x73, 0x3b, 0x78, 0x74, 0x64, 0x9a, 0x45, 0xb3, 0x6a, 0x37, 0x67,
	0x19, 0x3c, 0xa3, 0x23, 0xad, 0xd5, 0x71, 0xee, 0x32, 0x47, 0x85, 0x97, 0xbb, 0xbe, 0x53, 0x17,
	0x07, 0x3d, 0x25, 0x36, 0x2b, 0xcf, 0xf3, 0x8a, 0x9f, 0x90, 0x73, 0x76, 0xf9, 0xae, 0x2e, 0x27,
	0xab, 0xb5, 0x51, 0xa7, 0xbf, 0xca, 0xdc, 0x2f, 0x66, 0xab, 0x76, 0xb3, 0x97, 0xfa, 0x56, 0x9d,
	0x33, 0x0a, 0xf9, 0xb4, 0xb7, 0x9b, 0x75, 0x31, 0xcc, 0x2b, 0x7d, 0xad, 0x70, 0x7b, 0x52, 0x36,
	0xbe, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x06, 0x5b, 0x51, 0x94, 0x01, 0x00, 0x00,
}