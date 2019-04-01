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

package cisco_ios_xr_infra_policymgr_oper_policy_manager_global_transient_policy_map_targets_types_transient_policy_map_targets_type_transient_policy_map_targets_transient_policy_map_target

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
	proto.RegisterType((*PlmgrPmapStrTargetsBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.transient_policy_map_targets_types.transient_policy_map_targets_type.transient_policy_map_targets.transient_policy_map_target.plmgr_pmap_str_targets_bg_KEYS")
	proto.RegisterType((*PlmgrPmapStrTargetsBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.transient_policy_map_targets_types.transient_policy_map_targets_type.transient_policy_map_targets.transient_policy_map_target.plmgr_pmap_str_targets_bg")
}

func init() { proto.RegisterFile("plmgr_pmap_str_targets_bg.proto", fileDescriptor_f1b4b8a59f0baafb) }

var fileDescriptor_f1b4b8a59f0baafb = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0x15, 0xa5, 0x81, 0x2a, 0xe4, 0x14, 0x2f, 0x5a, 0x0a, 0x4a, 0x4f, 0x39, 0xd8,
	0xcf, 0xe0, 0x49, 0xf4, 0x50, 0xbd, 0x08, 0x85, 0x61, 0x76, 0x89, 0x21, 0x90, 0x3f, 0xc3, 0x24,
	0x07, 0xfb, 0xc1, 0xfc, 0x7e, 0xb2, 0x89, 0x2d, 0x5e, 0xba, 0xbd, 0xe5, 0xbd, 0xf7, 0xcb, 0x4b,
	0x66, 0xc4, 0x3d, 0xf9, 0x60, 0x19, 0x28, 0x20, 0x41, 0x2e, 0x0c, 0x05, 0xd9, 0x9a, 0x92, 0xa1,
	0xb7, 0x9a, 0x38, 0x95, 0x24, 0x7f, 0xba, 0xc1, 0xe5, 0x21, 0x81, 0x4b, 0x19, 0xbe, 0x19, 0x5c,
	0xfc, 0x62, 0x04, 0x4a, 0xde, 0x0d, 0xbb, 0xf1, 0x62, 0x22, 0xc3, 0xba, 0x49, 0x08, 0x18, 0xd1,
	0x1a, 0xd6, 0xd6, 0xa7, 0x1e, 0xbd, 0x2e, 0x8c, 0x31, 0x3b, 0x13, 0x0b, 0x1c, 0x72, 0x3a, 0xf4,
	0x97, 0x1d, 0x99, 0x7c, 0x1a, 0x99, 0x24, 0xa6, 0xc2, 0xe5, 0x56, 0xdc, 0x1d, 0x1d, 0x0d, 0x5e,
	0x9e, 0x3f, 0xdf, 0xa5, 0x14, 0x17, 0xe3, 0x13, 0xaa, 0x5b, 0x74, 0xab, 0xd9, 0xa6, 0x9e, 0xe5,
	0xa3, 0xb8, 0xf9, 0x57, 0x15, 0x31, 0x18, 0x75, 0x56, 0xe3, 0x79, 0xb3, 0x5f, 0x91, 0xde, 0x30,
	0x98, 0xe5, 0x56, 0xdc, 0x1e, 0x6d, 0x97, 0x4a, 0x5c, 0xfd, 0x29, 0xf5, 0xb4, 0x38, 0x5f, 0xcd,
	0x36, 0x7b, 0x29, 0x1f, 0xc4, 0x35, 0x21, 0x8f, 0x1f, 0xde, 0x03, 0xeb, 0x0a, 0xcc, 0x9b, 0xfb,
	0xd1, 0xcc, 0xfe, 0xb2, 0xae, 0x7e, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x43, 0x85, 0x62, 0x38,
	0x9d, 0x01, 0x00, 0x00,
}
