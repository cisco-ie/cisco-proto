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

package cisco_ios_xr_infra_policymgr_oper_policy_manager_global_policy_map_policy_map_types_policy_map_type_policy_map_details_policy_map_detail

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
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	PolicyMapName        string   `protobuf:"bytes,2,opt,name=policy_map_name,json=policyMapName,proto3" json:"policy_map_name,omitempty"`
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
	Object               []*PlmgrObjRefSummaryBg `protobuf:"bytes,53,rep,name=object,proto3" json:"object,omitempty"`
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

func (m *PlmgrAppDbBg) GetObject() []*PlmgrObjRefSummaryBg {
	if m != nil {
		return m.Object
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
	proto.RegisterType((*PlmgrAppDbBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_app_db_bg_KEYS")
	proto.RegisterType((*PlmgrObjRefPmapBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_obj_ref_pmap_bg")
	proto.RegisterType((*PlmgrObjRefSummaryBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_obj_ref_summary_bg")
	proto.RegisterType((*PlmgrPmapList)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_pmap_list")
	proto.RegisterType((*PlmgrObjCmapDetailBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_obj_cmap_detail_bg")
	proto.RegisterType((*PlmgrAppDbBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.policy_map.policy_map_types.policy_map_type.policy_map_details.policy_map_detail.plmgr_app_db_bg")
}

func init() { proto.RegisterFile("plmgr_app_db_bg.proto", fileDescriptor_6574dbf8cbffff90) }

var fileDescriptor_6574dbf8cbffff90 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xbd, 0x8e, 0x13, 0x31,
	0x10, 0x96, 0x73, 0x70, 0x22, 0x13, 0xc2, 0x81, 0xb9, 0x13, 0x5b, 0x80, 0x2e, 0x0a, 0x12, 0x4a,
	0xb5, 0xc5, 0x1d, 0x1c, 0x05, 0x25, 0x02, 0x84, 0xd0, 0x81, 0x64, 0x68, 0xa0, 0xb1, 0xbc, 0x7b,
	0xce, 0x6a, 0x91, 0x77, 0x6d, 0xd9, 0x46, 0x90, 0x37, 0xe0, 0x15, 0xa8, 0x10, 0xb4, 0x3c, 0x02,
	0x12, 0x0f, 0x80, 0xe8, 0x78, 0x21, 0xe4, 0x71, 0xf6, 0xe7, 0x36, 0xa2, 0xb8, 0x2e, 0xdd, 0xee,
	0xe7, 0x6f, 0xbc, 0xdf, 0xf7, 0xcd, 0x4c, 0x02, 0x07, 0x46, 0x55, 0x85, 0xe5, 0xc2, 0x18, 0x7e,
	0x96, 0xf1, 0xac, 0x48, 0x8d, 0xd5, 0x5e, 0xd3, 0xcf, 0x24, 0x2f, 0x5d, 0xae, 0x79, 0xa9, 0x1d,
	0xff, 0x64, 0x79, 0x59, 0x2f, 0xad, 0xe0, 0x46, 0xab, 0x32, 0x5f, 0x05, 0xba, 0x36, 0xd2, 0xa6,
	0xf1, 0x95, 0x57, 0xa2, 0x16, 0x85, 0xb4, 0x69, 0xa1, 0x74, 0x26, 0x54, 0x87, 0x9a, 0xde, 0x23,
	0xf7, 0x2b, 0x23, 0xdd, 0x10, 0xe8, 0xbf, 0x9f, 0x49, 0x2f, 0x4a, 0xe5, 0x36, 0xa1, 0x39, 0x83,
	0xfd, 0x81, 0x46, 0xfe, 0xe2, 0xc9, 0xdb, 0xd7, 0x94, 0xc2, 0xa5, 0x70, 0x49, 0x42, 0x66, 0x64,
	0x31, 0x66, 0xf8, 0x4c, 0xef, 0xc1, 0x5e, 0xef, 0x82, 0x5a, 0x54, 0x32, 0x19, 0xe1, 0xf1, 0x34,
	0xc2, 0xa7, 0xc2, 0xbc, 0x14, 0x95, 0x9c, 0xff, 0x20, 0x8d, 0x71, 0x9d, 0xbd, 0xe7, 0x56, 0x2e,
	0xb9, 0x09, 0x05, 0x59, 0x41, 0x9f, 0xc1, 0xcc, 0x6b, 0x2f, 0x14, 0x2f, 0x6b, 0x2f, 0x6d, 0x2d,
	0x54, 0x38, 0x95, 0x56, 0xd6, 0xb9, 0x0c, 0x5c, 0x99, 0x7b, 0x87, 0x5f, 0x9c, 0xb2, 0x3b, 0xc8,
	0x7b, 0xbe, 0xa6, 0xb1, 0x86, 0xf5, 0x2a, 0x92, 0xe8, 0x02, 0xae, 0xc7, 0x8b, 0x72, 0x25, 0x9c,
	0x0b, 0x7a, 0x1c, 0x6a, 0x99, 0xb2, 0x6b, 0x88, 0x3f, 0x0e, 0xf0, 0xa9, 0x30, 0x8e, 0x1e, 0xc2,
	0x24, 0x32, 0x97, 0x4a, 0x7f, 0x74, 0xc9, 0x0e, 0x92, 0x00, 0xa1, 0xa7, 0x01, 0x99, 0x3f, 0x82,
	0xe4, 0xbc, 0x58, 0xf7, 0xa1, 0xaa, 0x84, 0x5d, 0x05, 0xbd, 0x87, 0x30, 0x89, 0xb2, 0xa2, 0xdb,
	0x18, 0x06, 0x44, 0x08, 0xad, 0xbe, 0x83, 0xbd, 0x58, 0x8c, 0x0e, 0x55, 0xe9, 0x7c, 0xaf, 0xa6,
	0x17, 0xe0, 0xba, 0xe6, 0xcd, 0x45, 0x62, 0xfc, 0x36, 0xea, 0x2b, 0xcb, 0xbb, 0xa6, 0x05, 0x65,
	0x27, 0x70, 0x2b, 0xda, 0xfa, 0x5f, 0x80, 0x07, 0x78, 0xbc, 0x11, 0xdc, 0x1f, 0x02, 0xfb, 0xbd,
	0xaf, 0xb7, 0xd5, 0xc9, 0x68, 0xb6, 0xb3, 0x98, 0x1c, 0x7d, 0x21, 0xe9, 0xb6, 0x8c, 0x66, 0x3a,
	0x08, 0x96, 0xd1, 0x36, 0x9e, 0xd6, 0xd5, 0xfc, 0xf7, 0xe5, 0xa6, 0x01, 0xed, 0xfc, 0x0e, 0x1b,
	0x70, 0xb4, 0xd1, 0x80, 0xbb, 0x30, 0x8d, 0xd9, 0x35, 0x89, 0x1d, 0x63, 0x62, 0x57, 0x11, 0x6c,
	0x82, 0xfa, 0x45, 0x60, 0xdc, 0xa5, 0x73, 0x1f, 0xd3, 0xf9, 0xba, 0x75, 0xe9, 0x0c, 0x16, 0x8c,
	0x75, 0x92, 0xe9, 0x4f, 0x02, 0xbb, 0xd1, 0x60, 0xf2, 0x00, 0xd5, 0x7f, 0xdf, 0x5a, 0xf5, 0xdd,
	0xc6, 0xb1, 0xb5, 0x64, 0xfa, 0x97, 0xc0, 0xcd, 0x76, 0xb7, 0x7b, 0x63, 0x7a, 0xb2, 0xc5, 0x56,
	0xce, 0xaf, 0x28, 0xbb, 0x91, 0xaf, 0x7f, 0x84, 0xda, 0x71, 0xa5, 0xb7, 0x61, 0xec, 0xad, 0xa8,
	0x5d, 0x29, 0x6b, 0x9f, 0x3c, 0x9c, 0x91, 0xc5, 0x15, 0xd6, 0x01, 0xd9, 0x2e, 0xfe, 0x3b, 0x1c,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x23, 0xea, 0xd4, 0x36, 0x06, 0x00, 0x00,
}
