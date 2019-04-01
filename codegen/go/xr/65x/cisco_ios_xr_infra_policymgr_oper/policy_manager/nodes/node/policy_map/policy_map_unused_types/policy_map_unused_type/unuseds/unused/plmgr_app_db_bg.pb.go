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
// source: plmgr_app_db_bg.proto

package cisco_ios_xr_infra_policymgr_oper_policy_manager_nodes_node_policy_map_policy_map_unused_types_policy_map_unused_type_unuseds_unused

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

type PlmgrAppDbBg_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	PolicyMapName        string   `protobuf:"bytes,3,opt,name=policy_map_name,json=policyMapName,proto3" json:"policy_map_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmgrAppDbBg_KEYS) Reset()         { *m = PlmgrAppDbBg_KEYS{} }
func (m *PlmgrAppDbBg_KEYS) String() string { return proto.CompactTextString(m) }
func (*PlmgrAppDbBg_KEYS) ProtoMessage()    {}
func (*PlmgrAppDbBg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{0}
}

func (m *PlmgrAppDbBg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrAppDbBg_KEYS.Unmarshal(m, b)
}
func (m *PlmgrAppDbBg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrAppDbBg_KEYS.Marshal(b, m, deterministic)
}
func (m *PlmgrAppDbBg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrAppDbBg_KEYS.Merge(m, src)
}
func (m *PlmgrAppDbBg_KEYS) XXX_Size() int {
	return xxx_messageInfo_PlmgrAppDbBg_KEYS.Size(m)
}
func (m *PlmgrAppDbBg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrAppDbBg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrAppDbBg_KEYS proto.InternalMessageInfo

func (m *PlmgrAppDbBg_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PlmgrAppDbBg_KEYS) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlmgrAppDbBg_KEYS) GetPolicyMapName() string {
	if m != nil {
		return m.PolicyMapName
	}
	return ""
}

type PlmgrObjRefPmapBg struct {
	TotalInternalReferenceObjects uint32   `protobuf:"varint,1,opt,name=total_internal_reference_objects,json=totalInternalReferenceObjects,proto3" json:"total_internal_reference_objects,omitempty"`
	TotalClassMaps                uint32   `protobuf:"varint,2,opt,name=total_class_maps,json=totalClassMaps,proto3" json:"total_class_maps,omitempty"`
	TotalFlows                    uint32   `protobuf:"varint,3,opt,name=total_flows,json=totalFlows,proto3" json:"total_flows,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *PlmgrObjRefPmapBg) Reset()         { *m = PlmgrObjRefPmapBg{} }
func (m *PlmgrObjRefPmapBg) String() string { return proto.CompactTextString(m) }
func (*PlmgrObjRefPmapBg) ProtoMessage()    {}
func (*PlmgrObjRefPmapBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{1}
}

func (m *PlmgrObjRefPmapBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrObjRefPmapBg.Unmarshal(m, b)
}
func (m *PlmgrObjRefPmapBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrObjRefPmapBg.Marshal(b, m, deterministic)
}
func (m *PlmgrObjRefPmapBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrObjRefPmapBg.Merge(m, src)
}
func (m *PlmgrObjRefPmapBg) XXX_Size() int {
	return xxx_messageInfo_PlmgrObjRefPmapBg.Size(m)
}
func (m *PlmgrObjRefPmapBg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrObjRefPmapBg.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrObjRefPmapBg proto.InternalMessageInfo

func (m *PlmgrObjRefPmapBg) GetTotalInternalReferenceObjects() uint32 {
	if m != nil {
		return m.TotalInternalReferenceObjects
	}
	return 0
}

func (m *PlmgrObjRefPmapBg) GetTotalClassMaps() uint32 {
	if m != nil {
		return m.TotalClassMaps
	}
	return 0
}

func (m *PlmgrObjRefPmapBg) GetTotalFlows() uint32 {
	if m != nil {
		return m.TotalFlows
	}
	return 0
}

type PlmgrObjRefSummaryBg struct {
	ObjectName           string   `protobuf:"bytes,1,opt,name=object_name,json=objectName,proto3" json:"object_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmgrObjRefSummaryBg) Reset()         { *m = PlmgrObjRefSummaryBg{} }
func (m *PlmgrObjRefSummaryBg) String() string { return proto.CompactTextString(m) }
func (*PlmgrObjRefSummaryBg) ProtoMessage()    {}
func (*PlmgrObjRefSummaryBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{2}
}

func (m *PlmgrObjRefSummaryBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrObjRefSummaryBg.Unmarshal(m, b)
}
func (m *PlmgrObjRefSummaryBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrObjRefSummaryBg.Marshal(b, m, deterministic)
}
func (m *PlmgrObjRefSummaryBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrObjRefSummaryBg.Merge(m, src)
}
func (m *PlmgrObjRefSummaryBg) XXX_Size() int {
	return xxx_messageInfo_PlmgrObjRefSummaryBg.Size(m)
}
func (m *PlmgrObjRefSummaryBg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrObjRefSummaryBg.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrObjRefSummaryBg proto.InternalMessageInfo

func (m *PlmgrObjRefSummaryBg) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

type PlmgrPmapList struct {
	ObjectType           string   `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	PolicyMapName        string   `protobuf:"bytes,2,opt,name=policy_map_name,json=policyMapName,proto3" json:"policy_map_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlmgrPmapList) Reset()         { *m = PlmgrPmapList{} }
func (m *PlmgrPmapList) String() string { return proto.CompactTextString(m) }
func (*PlmgrPmapList) ProtoMessage()    {}
func (*PlmgrPmapList) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{3}
}

func (m *PlmgrPmapList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrPmapList.Unmarshal(m, b)
}
func (m *PlmgrPmapList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrPmapList.Marshal(b, m, deterministic)
}
func (m *PlmgrPmapList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrPmapList.Merge(m, src)
}
func (m *PlmgrPmapList) XXX_Size() int {
	return xxx_messageInfo_PlmgrPmapList.Size(m)
}
func (m *PlmgrPmapList) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrPmapList.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrPmapList proto.InternalMessageInfo

func (m *PlmgrPmapList) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *PlmgrPmapList) GetPolicyMapName() string {
	if m != nil {
		return m.PolicyMapName
	}
	return ""
}

type PlmgrObjCmapDetailBg struct {
	TotalReferenceObjects uint32           `protobuf:"varint,1,opt,name=total_reference_objects,json=totalReferenceObjects,proto3" json:"total_reference_objects,omitempty"`
	PolicyMapReference    []*PlmgrPmapList `protobuf:"bytes,2,rep,name=policy_map_reference,json=policyMapReference,proto3" json:"policy_map_reference,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *PlmgrObjCmapDetailBg) Reset()         { *m = PlmgrObjCmapDetailBg{} }
func (m *PlmgrObjCmapDetailBg) String() string { return proto.CompactTextString(m) }
func (*PlmgrObjCmapDetailBg) ProtoMessage()    {}
func (*PlmgrObjCmapDetailBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{4}
}

func (m *PlmgrObjCmapDetailBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrObjCmapDetailBg.Unmarshal(m, b)
}
func (m *PlmgrObjCmapDetailBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrObjCmapDetailBg.Marshal(b, m, deterministic)
}
func (m *PlmgrObjCmapDetailBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrObjCmapDetailBg.Merge(m, src)
}
func (m *PlmgrObjCmapDetailBg) XXX_Size() int {
	return xxx_messageInfo_PlmgrObjCmapDetailBg.Size(m)
}
func (m *PlmgrObjCmapDetailBg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrObjCmapDetailBg.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrObjCmapDetailBg proto.InternalMessageInfo

func (m *PlmgrObjCmapDetailBg) GetTotalReferenceObjects() uint32 {
	if m != nil {
		return m.TotalReferenceObjects
	}
	return 0
}

func (m *PlmgrObjCmapDetailBg) GetPolicyMapReference() []*PlmgrPmapList {
	if m != nil {
		return m.PolicyMapReference
	}
	return nil
}

type PlmgrAppDbBg struct {
	ObjectType           string                  `protobuf:"bytes,50,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	TotalObjects         uint32                  `protobuf:"varint,51,opt,name=total_objects,json=totalObjects,proto3" json:"total_objects,omitempty"`
	Reference            []*PlmgrObjRefPmapBg    `protobuf:"bytes,52,rep,name=reference,proto3" json:"reference,omitempty"`
	Objects              []*PlmgrObjRefSummaryBg `protobuf:"bytes,53,rep,name=objects,proto3" json:"objects,omitempty"`
	ClassMapReference    []*PlmgrObjCmapDetailBg `protobuf:"bytes,54,rep,name=class_map_reference,json=classMapReference,proto3" json:"class_map_reference,omitempty"`
	Transient            bool                    `protobuf:"varint,55,opt,name=transient,proto3" json:"transient,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PlmgrAppDbBg) Reset()         { *m = PlmgrAppDbBg{} }
func (m *PlmgrAppDbBg) String() string { return proto.CompactTextString(m) }
func (*PlmgrAppDbBg) ProtoMessage()    {}
func (*PlmgrAppDbBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_6574dbf8cbffff90, []int{5}
}

func (m *PlmgrAppDbBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlmgrAppDbBg.Unmarshal(m, b)
}
func (m *PlmgrAppDbBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlmgrAppDbBg.Marshal(b, m, deterministic)
}
func (m *PlmgrAppDbBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlmgrAppDbBg.Merge(m, src)
}
func (m *PlmgrAppDbBg) XXX_Size() int {
	return xxx_messageInfo_PlmgrAppDbBg.Size(m)
}
func (m *PlmgrAppDbBg) XXX_DiscardUnknown() {
	xxx_messageInfo_PlmgrAppDbBg.DiscardUnknown(m)
}

var xxx_messageInfo_PlmgrAppDbBg proto.InternalMessageInfo

func (m *PlmgrAppDbBg) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *PlmgrAppDbBg) GetTotalObjects() uint32 {
	if m != nil {
		return m.TotalObjects
	}
	return 0
}

func (m *PlmgrAppDbBg) GetReference() []*PlmgrObjRefPmapBg {
	if m != nil {
		return m.Reference
	}
	return nil
}

func (m *PlmgrAppDbBg) GetObjects() []*PlmgrObjRefSummaryBg {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *PlmgrAppDbBg) GetClassMapReference() []*PlmgrObjCmapDetailBg {
	if m != nil {
		return m.ClassMapReference
	}
	return nil
}

func (m *PlmgrAppDbBg) GetTransient() bool {
	if m != nil {
		return m.Transient
	}
	return false
}

func init() {
	proto.RegisterType((*PlmgrAppDbBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_app_db_bg_KEYS")
	proto.RegisterType((*PlmgrObjRefPmapBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_obj_ref_pmap_bg")
	proto.RegisterType((*PlmgrObjRefSummaryBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_obj_ref_summary_bg")
	proto.RegisterType((*PlmgrPmapList)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_pmap_list")
	proto.RegisterType((*PlmgrObjCmapDetailBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_obj_cmap_detail_bg")
	proto.RegisterType((*PlmgrAppDbBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.nodes.node.policy_map.policy_map_unused_types.policy_map_unused_type.unuseds.unused.plmgr_app_db_bg")
}

func init() { proto.RegisterFile("plmgr_app_db_bg.proto", fileDescriptor_6574dbf8cbffff90) }

var fileDescriptor_6574dbf8cbffff90 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xbf, 0x8e, 0x13, 0x3d,
	0x14, 0xc5, 0xe5, 0xec, 0xf7, 0xc1, 0xe6, 0x86, 0xb0, 0x60, 0x76, 0xc5, 0x48, 0x80, 0x36, 0x0a,
	0x12, 0x4a, 0x35, 0xc5, 0x2e, 0x2c, 0x05, 0x25, 0x02, 0x84, 0xd0, 0x82, 0x64, 0x68, 0xa0, 0xb1,
	0x3c, 0x13, 0x27, 0x1a, 0x34, 0x63, 0x5b, 0xb6, 0x23, 0x48, 0xcf, 0x2b, 0x50, 0x21, 0x78, 0x01,
	0x9e, 0x00, 0x2a, 0x4a, 0x3a, 0x5e, 0x09, 0xf9, 0x3a, 0xf3, 0x27, 0x59, 0xd2, 0xa7, 0x49, 0x9c,
	0xe3, 0x73, 0xc7, 0xe7, 0xfe, 0xe4, 0x9b, 0x81, 0x23, 0x53, 0x56, 0x73, 0xcb, 0x85, 0x31, 0x7c,
	0x9a, 0xf1, 0x6c, 0x9e, 0x1a, 0xab, 0xbd, 0xa6, 0x9f, 0x48, 0x5e, 0xb8, 0x5c, 0xf3, 0x42, 0x3b,
	0xfe, 0xd1, 0xf2, 0x42, 0xcd, 0xac, 0xe0, 0x46, 0x97, 0x45, 0xbe, 0x0c, 0x76, 0x6d, 0xa4, 0x4d,
	0xe3, 0x4f, 0x5e, 0x09, 0x25, 0xe6, 0xd2, 0xa6, 0x4a, 0x4f, 0xa5, 0xc3, 0xcf, 0x76, 0xc7, 0x74,
	0x96, 0x7c, 0xa1, 0x16, 0x4e, 0x4e, 0xb9, 0x5f, 0x1a, 0xe9, 0xb6, 0xe8, 0x69, 0x5c, 0xbb, 0xd5,
	0xf7, 0x58, 0xc3, 0xe1, 0x46, 0x3e, 0xfe, 0xe2, 0xc9, 0xdb, 0xd7, 0xf4, 0x16, 0xf4, 0xc3, 0x41,
	0x5c, 0x89, 0x4a, 0x26, 0x64, 0x44, 0x26, 0x7d, 0xb6, 0x1f, 0x84, 0x97, 0xa2, 0x92, 0x94, 0xc2,
	0x7f, 0xe1, 0x59, 0x49, 0x0f, 0x75, 0x5c, 0xd3, 0x7b, 0x70, 0xd0, 0x39, 0x11, 0xcb, 0xf6, 0x70,
	0x7b, 0x18, 0xe5, 0x73, 0x61, 0x42, 0xed, 0xf8, 0x3b, 0xa9, 0x89, 0xe8, 0xec, 0x3d, 0xb7, 0x72,
	0xc6, 0x4d, 0x28, 0xc8, 0xe6, 0xf4, 0x19, 0x8c, 0xbc, 0xf6, 0xa2, 0xe4, 0x85, 0xf2, 0xd2, 0x2a,
	0x51, 0x86, 0x5d, 0x69, 0xa5, 0xca, 0x65, 0xf0, 0xca, 0xdc, 0x3b, 0x4c, 0x32, 0x64, 0x77, 0xd0,
	0xf7, 0x7c, 0x65, 0x63, 0xb5, 0xeb, 0x55, 0x34, 0xd1, 0x09, 0x5c, 0x8b, 0x0f, 0xca, 0x4b, 0xe1,
	0x5c, 0xc8, 0xe3, 0x30, 0xea, 0x90, 0x5d, 0x45, 0xfd, 0x71, 0x90, 0xcf, 0x85, 0x71, 0xf4, 0x18,
	0x06, 0xd1, 0x39, 0x2b, 0xf5, 0x07, 0x87, 0x81, 0x87, 0x0c, 0x50, 0x7a, 0x1a, 0x94, 0xf1, 0x23,
	0x48, 0xd6, 0xc3, 0xba, 0x45, 0x55, 0x09, 0xbb, 0x0c, 0x79, 0x8f, 0x61, 0x10, 0x63, 0x75, 0x21,
	0x41, 0x94, 0xb0, 0xd5, 0x77, 0x70, 0x10, 0x8b, 0xb1, 0xc3, 0xb2, 0x70, 0xbe, 0x53, 0x83, 0x00,
	0xd7, 0x6a, 0xde, 0x6c, 0xc1, 0xd8, 0xfb, 0x17, 0xc6, 0xaf, 0xbd, 0x6e, 0xb2, 0x3c, 0x98, 0xa7,
	0xd2, 0x8b, 0xa2, 0x0c, 0xc9, 0xce, 0xe0, 0x66, 0x6c, 0x6b, 0x1b, 0xc0, 0x23, 0xdc, 0xbe, 0x00,
	0xee, 0x37, 0x81, 0xc3, 0xce, 0xe9, 0x4d, 0x75, 0xd2, 0x1b, 0xed, 0x4d, 0x06, 0x27, 0x9f, 0x49,
	0xba, 0x0b, 0x77, 0x36, 0xdd, 0x80, 0xca, 0x68, 0x83, 0xa6, 0xe9, 0x68, 0xfc, 0xeb, 0xff, 0x1a,
	0x7e, 0x73, 0xb1, 0x37, 0xe1, 0x9f, 0x5c, 0x80, 0x7f, 0x17, 0x86, 0x91, 0x5b, 0x4d, 0xeb, 0x14,
	0x69, 0x5d, 0x41, 0xb1, 0x86, 0xf4, 0x93, 0x40, 0xbf, 0x25, 0x73, 0x1f, 0xc9, 0x7c, 0xd9, 0x29,
	0x32, 0x1b, 0x83, 0xc5, 0xda, 0xb8, 0xf4, 0x07, 0x81, 0xcb, 0x75, 0x73, 0x0f, 0x30, 0xfa, 0xb7,
	0x9d, 0x8c, 0xde, 0x8e, 0x19, 0xab, 0xf3, 0xd2, 0x3f, 0x04, 0x6e, 0x34, 0x13, 0xdd, 0xb9, 0x9c,
	0x67, 0x3b, 0xda, 0xc7, 0xfa, 0x50, 0xb2, 0xeb, 0xf9, 0xea, 0x6f, 0xa7, 0xb9, 0xa4, 0xf4, 0x36,
	0xf4, 0xbd, 0x15, 0xca, 0x15, 0x52, 0xf9, 0xe4, 0xe1, 0x88, 0x4c, 0xf6, 0x59, 0x2b, 0x64, 0x97,
	0xf0, 0x45, 0x71, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x5a, 0xcb, 0xf3, 0x41, 0x06, 0x00,
	0x00,
}