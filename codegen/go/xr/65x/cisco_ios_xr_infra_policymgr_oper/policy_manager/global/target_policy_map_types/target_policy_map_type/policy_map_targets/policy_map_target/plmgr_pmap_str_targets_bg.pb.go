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

package cisco_ios_xr_infra_policymgr_oper_policy_manager_global_target_policy_map_types_target_policy_map_type_policy_map_targets_policy_map_target

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
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PolicyMapName        string   `protobuf:"bytes,2,opt,name=policy_map_name,json=policyMapName,proto3" json:"policy_map_name,omitempty"`
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
	proto.RegisterType((*PlmgrPmapStrTargetsBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.target_policy_map_types.target_policy_map_type.policy_map_targets.policy_map_target.plmgr_pmap_str_targets_bg_KEYS")
	proto.RegisterType((*PlmgrPmapStrTargetsBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.target_policy_map_types.target_policy_map_type.policy_map_targets.policy_map_target.plmgr_pmap_str_targets_bg")
}

func init() { proto.RegisterFile("plmgr_pmap_str_targets_bg.proto", fileDescriptor_f1b4b8a59f0baafb) }

var fileDescriptor_f1b4b8a59f0baafb = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x59, 0x15, 0xa5, 0x81, 0x2a, 0xe4, 0x14, 0x2f, 0x5a, 0x0a, 0x4a, 0x4f, 0x39, 0xd8,
	0x67, 0xf0, 0x24, 0x7a, 0xa8, 0x5e, 0x84, 0xc2, 0x30, 0xbb, 0xc4, 0x10, 0xc8, 0x9f, 0x61, 0x92,
	0x83, 0x7d, 0x06, 0x5f, 0x5a, 0x9a, 0xb8, 0xb2, 0x50, 0xf6, 0x96, 0xef, 0x37, 0x93, 0x1f, 0x1f,
	0x23, 0xee, 0xc9, 0x07, 0xcb, 0x40, 0x01, 0x09, 0x72, 0x61, 0x28, 0xc8, 0xd6, 0x94, 0x0c, 0xbd,
	0xd5, 0xc4, 0xa9, 0x24, 0xf9, 0xd3, 0x0d, 0x2e, 0x0f, 0x09, 0x5c, 0xca, 0xf0, 0xcd, 0xe0, 0xe2,
	0x17, 0x23, 0x50, 0xf2, 0x6e, 0x38, 0x1c, 0x3f, 0x26, 0x32, 0xac, 0x5b, 0x84, 0x80, 0x11, 0xad,
	0x61, 0x6d, 0x7d, 0xea, 0xd1, 0xeb, 0xa6, 0x82, 0xff, 0x21, 0x41, 0x39, 0x90, 0xc9, 0x33, 0x5c,
	0x4f, 0x73, 0x2b, 0x71, 0x8a, 0xd6, 0x7b, 0x71, 0x37, 0x5b, 0x18, 0x5e, 0x9e, 0x3f, 0xdf, 0xa5,
	0x14, 0x17, 0x47, 0x9d, 0xea, 0x56, 0xdd, 0x66, 0xb1, 0xab, 0x6f, 0xf9, 0x28, 0x6e, 0x26, 0xaa,
	0x88, 0xc1, 0xa8, 0xb3, 0x3a, 0x5e, 0x36, 0xfc, 0x8a, 0xf4, 0x86, 0xc1, 0xac, 0xf7, 0xe2, 0x76,
	0xd6, 0x2e, 0x95, 0xb8, 0xfa, 0x4b, 0xea, 0x69, 0x75, 0xbe, 0x59, 0xec, 0xc6, 0x28, 0x1f, 0xc4,
	0x35, 0x21, 0x9b, 0x58, 0xc6, 0x75, 0xb5, 0xad, 0x0b, 0xcb, 0x46, 0x3f, 0x1a, 0xec, 0x2f, 0xeb,
	0x41, 0xb7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x77, 0x23, 0xe3, 0x73, 0x01, 0x00, 0x00,
}