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

package cisco_ios_xr_infra_policymgr_oper_policy_manager_global_class_map_class_map_types_class_map_type_class_map_details_class_map_detail

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
	ClassMapName         string   `protobuf:"bytes,2,opt,name=class_map_name,json=classMapName,proto3" json:"class_map_name,omitempty"`
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

func (m *PlmgrAppDbBg_KEYS) GetClassMapName() string {
	if m != nil {
		return m.ClassMapName
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
	proto.RegisterType((*PlmgrAppDbBg_KEYS)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_app_db_bg_KEYS")
	proto.RegisterType((*PlmgrObjRefPmapBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_obj_ref_pmap_bg")
	proto.RegisterType((*PlmgrObjRefSummaryBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_obj_ref_summary_bg")
	proto.RegisterType((*PlmgrPmapList)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_pmap_list")
	proto.RegisterType((*PlmgrObjCmapDetailBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_obj_cmap_detail_bg")
	proto.RegisterType((*PlmgrAppDbBg)(nil), "cisco_ios_xr_infra_policymgr_oper.policy_manager.global.class_map.class_map_types.class_map_type.class_map_details.class_map_detail.plmgr_app_db_bg")
}

func init() { proto.RegisterFile("plmgr_app_db_bg.proto", fileDescriptor_6574dbf8cbffff90) }

var fileDescriptor_6574dbf8cbffff90 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xbd, 0x8e, 0x13, 0x31,
	0x14, 0x85, 0xe5, 0x2c, 0x7f, 0xb9, 0xd9, 0xec, 0x82, 0xd9, 0x15, 0x53, 0x80, 0x36, 0x0a, 0x08,
	0xa5, 0x9a, 0x62, 0x17, 0x96, 0x82, 0x12, 0x01, 0x42, 0x88, 0x1f, 0x19, 0x1a, 0x68, 0x2c, 0xcf,
	0xac, 0x13, 0x0d, 0xf2, 0x8c, 0x2d, 0xdb, 0x08, 0x52, 0xf3, 0x0a, 0x54, 0x08, 0x9e, 0x80, 0x17,
	0x80, 0x82, 0x9a, 0x8e, 0x67, 0x42, 0xb6, 0xe3, 0xf1, 0x64, 0xa2, 0xed, 0xd3, 0x8d, 0x8f, 0xcf,
	0xb5, 0xcf, 0xfd, 0x6c, 0x27, 0x70, 0xa8, 0x44, 0xbd, 0xd0, 0x94, 0x29, 0x45, 0xcf, 0x0a, 0x5a,
	0x2c, 0x72, 0xa5, 0xa5, 0x95, 0xf8, 0x0b, 0x2a, 0x2b, 0x53, 0x4a, 0x5a, 0x49, 0x43, 0x3f, 0x6b,
	0x5a, 0x35, 0x73, 0xcd, 0xa8, 0x92, 0xa2, 0x2a, 0x97, 0xce, 0x2e, 0x15, 0xd7, 0x79, 0x18, 0xd2,
	0x9a, 0x35, 0x6c, 0xc1, 0x75, 0xbe, 0x10, 0xb2, 0x60, 0x22, 0x2f, 0x05, 0x33, 0x86, 0xd6, 0x4c,
	0xa5, 0x2f, 0x6a, 0x97, 0x8a, 0x9b, 0xde, 0xb8, 0x33, 0x3c, 0xe3, 0x96, 0x55, 0xc2, 0x6c, 0x28,
	0xd3, 0xd7, 0x70, 0xd0, 0x8b, 0x47, 0x9f, 0x3f, 0x7e, 0xf7, 0x06, 0x63, 0xb8, 0xe0, 0x96, 0xc8,
	0xd0, 0x04, 0xcd, 0x86, 0xc4, 0x7f, 0xe3, 0x3b, 0xb0, 0x97, 0xea, 0x1b, 0x56, 0xf3, 0x6c, 0xe0,
	0x67, 0x77, 0xbd, 0xfa, 0x82, 0xa9, 0x97, 0xac, 0xe6, 0xd3, 0x9f, 0x28, 0x76, 0x2c, 0x8b, 0x0f,
	0x54, 0xf3, 0x39, 0x55, 0xce, 0x5f, 0x2c, 0xf0, 0x53, 0x98, 0x58, 0x69, 0x99, 0xa0, 0x55, 0x63,
	0xb9, 0x6e, 0x98, 0x70, 0xb3, 0x5c, 0xf3, 0xa6, 0xe4, 0xce, 0xcb, 0x4b, 0x6b, 0xfc, 0x7e, 0x63,
	0x72, 0xcb, 0xfb, 0x9e, 0xad, 0x6c, 0x24, 0xba, 0x5e, 0x05, 0x13, 0x9e, 0xc1, 0xd5, 0xb0, 0x50,
	0x1b, 0xc7, 0xf8, 0x28, 0x63, 0xb2, 0xe7, 0xf5, 0x47, 0xab, 0x3c, 0x06, 0x1f, 0xc1, 0x28, 0x38,
	0xe7, 0x42, 0x7e, 0x32, 0xd9, 0x8e, 0x37, 0x81, 0x97, 0x9e, 0x38, 0x65, 0xfa, 0x10, 0xb2, 0xf5,
	0xb0, 0xe6, 0x63, 0x5d, 0x33, 0xbd, 0x74, 0x79, 0x8f, 0x60, 0x14, 0x62, 0x85, 0x66, 0x03, 0x0a,
	0x08, 0x92, 0x6f, 0xf5, 0x3d, 0xec, 0x87, 0x62, 0xdf, 0xa1, 0xa8, 0x8c, 0xed, 0xd4, 0x74, 0xf0,
	0xad, 0x6a, 0xde, 0x3a, 0x88, 0x77, 0x61, 0xbf, 0x3d, 0xd5, 0x35, 0x8a, 0xe3, 0x20, 0x47, 0x8c,
	0xdf, 0x07, 0xdd, 0x64, 0x65, 0x3a, 0x32, 0x97, 0xec, 0x14, 0x6e, 0x84, 0xb6, 0xce, 0x03, 0x78,
	0xe8, 0xa7, 0x37, 0xc0, 0xfd, 0x45, 0x70, 0xd0, 0xd9, 0xbd, 0xad, 0xce, 0x06, 0x93, 0x9d, 0xd9,
	0xe8, 0xf8, 0x2b, 0xca, 0xb7, 0xe0, 0x4e, 0xe6, 0x3d, 0xa6, 0x04, 0xb7, 0x64, 0xda, 0x86, 0xa6,
	0x7f, 0x2e, 0x46, 0xf6, 0xed, 0xc5, 0xed, 0xb3, 0x3f, 0xde, 0x60, 0x7f, 0x1b, 0xc6, 0x01, 0x5b,
	0x84, 0x75, 0xe2, 0x61, 0xed, 0x7a, 0x31, 0x32, 0xfa, 0x8d, 0x60, 0x98, 0xc0, 0xdc, 0xf3, 0x60,
	0xbe, 0x6d, 0x13, 0x98, 0xde, 0xb3, 0x22, 0x29, 0x2d, 0xfe, 0x85, 0xe0, 0x72, 0xec, 0xed, 0xbe,
	0x4f, 0xfe, 0x63, 0x1b, 0x93, 0xa7, 0x37, 0x46, 0x62, 0x5c, 0xfc, 0x0f, 0xc1, 0xf5, 0x54, 0x96,
	0x0e, 0xe0, 0x74, 0x3b, 0xdb, 0x58, 0x7f, 0x90, 0xe4, 0x5a, 0xfc, 0x09, 0x6c, 0x6f, 0x28, 0xbe,
	0x09, 0x43, 0xab, 0x59, 0x63, 0x2a, 0xde, 0xd8, 0xec, 0xc1, 0x04, 0xcd, 0xae, 0x90, 0x24, 0x14,
	0x97, 0xfc, 0x9f, 0xc0, 0xc9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xae, 0xd6, 0x5d, 0x1d,
	0x06, 0x00, 0x00,
}
