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
// source: fia_edm_stats_info.proto

package cisco_ios_xr_fretta_bcm_dpa_resources_oper_dpa_stats_nodes_node_asic_statistics_asic_statistics_detail_for_npu_ids_asic_statistics_detail_for_npu_id

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

type FiaEdmStatsInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NpuId                uint32   `protobuf:"varint,2,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FiaEdmStatsInfo_KEYS) Reset()         { *m = FiaEdmStatsInfo_KEYS{} }
func (m *FiaEdmStatsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FiaEdmStatsInfo_KEYS) ProtoMessage()    {}
func (*FiaEdmStatsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57f30811988eb7b, []int{0}
}

func (m *FiaEdmStatsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmStatsInfo_KEYS.Unmarshal(m, b)
}
func (m *FiaEdmStatsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmStatsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FiaEdmStatsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmStatsInfo_KEYS.Merge(m, src)
}
func (m *FiaEdmStatsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FiaEdmStatsInfo_KEYS.Size(m)
}
func (m *FiaEdmStatsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmStatsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmStatsInfo_KEYS proto.InternalMessageInfo

func (m *FiaEdmStatsInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FiaEdmStatsInfo_KEYS) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

type FiaEdmDeviceStatsFieldInfo struct {
	FieldName            string   `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	FieldValue           uint64   `protobuf:"varint,2,opt,name=field_value,json=fieldValue,proto3" json:"field_value,omitempty"`
	IsOverflow           bool     `protobuf:"varint,3,opt,name=is_overflow,json=isOverflow,proto3" json:"is_overflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FiaEdmDeviceStatsFieldInfo) Reset()         { *m = FiaEdmDeviceStatsFieldInfo{} }
func (m *FiaEdmDeviceStatsFieldInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmDeviceStatsFieldInfo) ProtoMessage()    {}
func (*FiaEdmDeviceStatsFieldInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57f30811988eb7b, []int{1}
}

func (m *FiaEdmDeviceStatsFieldInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmDeviceStatsFieldInfo.Unmarshal(m, b)
}
func (m *FiaEdmDeviceStatsFieldInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmDeviceStatsFieldInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmDeviceStatsFieldInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmDeviceStatsFieldInfo.Merge(m, src)
}
func (m *FiaEdmDeviceStatsFieldInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmDeviceStatsFieldInfo.Size(m)
}
func (m *FiaEdmDeviceStatsFieldInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmDeviceStatsFieldInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmDeviceStatsFieldInfo proto.InternalMessageInfo

func (m *FiaEdmDeviceStatsFieldInfo) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *FiaEdmDeviceStatsFieldInfo) GetFieldValue() uint64 {
	if m != nil {
		return m.FieldValue
	}
	return 0
}

func (m *FiaEdmDeviceStatsFieldInfo) GetIsOverflow() bool {
	if m != nil {
		return m.IsOverflow
	}
	return false
}

type FiaEdmDeviceStatsBlkInfo struct {
	BlockName            string                        `protobuf:"bytes,1,opt,name=block_name,json=blockName,proto3" json:"block_name,omitempty"`
	FieldInfo            []*FiaEdmDeviceStatsFieldInfo `protobuf:"bytes,2,rep,name=field_info,json=fieldInfo,proto3" json:"field_info,omitempty"`
	NumFields            uint32                        `protobuf:"varint,3,opt,name=num_fields,json=numFields,proto3" json:"num_fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *FiaEdmDeviceStatsBlkInfo) Reset()         { *m = FiaEdmDeviceStatsBlkInfo{} }
func (m *FiaEdmDeviceStatsBlkInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmDeviceStatsBlkInfo) ProtoMessage()    {}
func (*FiaEdmDeviceStatsBlkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57f30811988eb7b, []int{2}
}

func (m *FiaEdmDeviceStatsBlkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmDeviceStatsBlkInfo.Unmarshal(m, b)
}
func (m *FiaEdmDeviceStatsBlkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmDeviceStatsBlkInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmDeviceStatsBlkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmDeviceStatsBlkInfo.Merge(m, src)
}
func (m *FiaEdmDeviceStatsBlkInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmDeviceStatsBlkInfo.Size(m)
}
func (m *FiaEdmDeviceStatsBlkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmDeviceStatsBlkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmDeviceStatsBlkInfo proto.InternalMessageInfo

func (m *FiaEdmDeviceStatsBlkInfo) GetBlockName() string {
	if m != nil {
		return m.BlockName
	}
	return ""
}

func (m *FiaEdmDeviceStatsBlkInfo) GetFieldInfo() []*FiaEdmDeviceStatsFieldInfo {
	if m != nil {
		return m.FieldInfo
	}
	return nil
}

func (m *FiaEdmDeviceStatsBlkInfo) GetNumFields() uint32 {
	if m != nil {
		return m.NumFields
	}
	return 0
}

type FiaEdmDeviceStatsAsicInfo struct {
	BlockInfo            []*FiaEdmDeviceStatsBlkInfo `protobuf:"bytes,1,rep,name=block_info,json=blockInfo,proto3" json:"block_info,omitempty"`
	NumBlocks            uint32                      `protobuf:"varint,2,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FiaEdmDeviceStatsAsicInfo) Reset()         { *m = FiaEdmDeviceStatsAsicInfo{} }
func (m *FiaEdmDeviceStatsAsicInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmDeviceStatsAsicInfo) ProtoMessage()    {}
func (*FiaEdmDeviceStatsAsicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57f30811988eb7b, []int{3}
}

func (m *FiaEdmDeviceStatsAsicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicInfo.Unmarshal(m, b)
}
func (m *FiaEdmDeviceStatsAsicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmDeviceStatsAsicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmDeviceStatsAsicInfo.Merge(m, src)
}
func (m *FiaEdmDeviceStatsAsicInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicInfo.Size(m)
}
func (m *FiaEdmDeviceStatsAsicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmDeviceStatsAsicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmDeviceStatsAsicInfo proto.InternalMessageInfo

func (m *FiaEdmDeviceStatsAsicInfo) GetBlockInfo() []*FiaEdmDeviceStatsBlkInfo {
	if m != nil {
		return m.BlockInfo
	}
	return nil
}

func (m *FiaEdmDeviceStatsAsicInfo) GetNumBlocks() uint32 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

type FiaEdmStatsInfo struct {
	Valid                bool                       `protobuf:"varint,50,opt,name=valid,proto3" json:"valid,omitempty"`
	RackNumber           uint32                     `protobuf:"varint,51,opt,name=rack_number,json=rackNumber,proto3" json:"rack_number,omitempty"`
	SlotNumber           uint32                     `protobuf:"varint,52,opt,name=slot_number,json=slotNumber,proto3" json:"slot_number,omitempty"`
	AsicInstance         uint32                     `protobuf:"varint,53,opt,name=asic_instance,json=asicInstance,proto3" json:"asic_instance,omitempty"`
	ChipVersion          uint32                     `protobuf:"varint,54,opt,name=chip_version,json=chipVersion,proto3" json:"chip_version,omitempty"`
	Statistics           *FiaEdmDeviceStatsAsicInfo `protobuf:"bytes,55,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FiaEdmStatsInfo) Reset()         { *m = FiaEdmStatsInfo{} }
func (m *FiaEdmStatsInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmStatsInfo) ProtoMessage()    {}
func (*FiaEdmStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d57f30811988eb7b, []int{4}
}

func (m *FiaEdmStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmStatsInfo.Unmarshal(m, b)
}
func (m *FiaEdmStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmStatsInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmStatsInfo.Merge(m, src)
}
func (m *FiaEdmStatsInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmStatsInfo.Size(m)
}
func (m *FiaEdmStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmStatsInfo proto.InternalMessageInfo

func (m *FiaEdmStatsInfo) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *FiaEdmStatsInfo) GetRackNumber() uint32 {
	if m != nil {
		return m.RackNumber
	}
	return 0
}

func (m *FiaEdmStatsInfo) GetSlotNumber() uint32 {
	if m != nil {
		return m.SlotNumber
	}
	return 0
}

func (m *FiaEdmStatsInfo) GetAsicInstance() uint32 {
	if m != nil {
		return m.AsicInstance
	}
	return 0
}

func (m *FiaEdmStatsInfo) GetChipVersion() uint32 {
	if m != nil {
		return m.ChipVersion
	}
	return 0
}

func (m *FiaEdmStatsInfo) GetStatistics() *FiaEdmDeviceStatsAsicInfo {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FiaEdmStatsInfo_KEYS)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_detail_for_npu_ids.asic_statistics_detail_for_npu_id.fia_edm_stats_info_KEYS")
	proto.RegisterType((*FiaEdmDeviceStatsFieldInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_detail_for_npu_ids.asic_statistics_detail_for_npu_id.fia_edm_device_stats_field_info")
	proto.RegisterType((*FiaEdmDeviceStatsBlkInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_detail_for_npu_ids.asic_statistics_detail_for_npu_id.fia_edm_device_stats_blk_info")
	proto.RegisterType((*FiaEdmDeviceStatsAsicInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_detail_for_npu_ids.asic_statistics_detail_for_npu_id.fia_edm_device_stats_asic_info")
	proto.RegisterType((*FiaEdmStatsInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_resources_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_detail_for_npu_ids.asic_statistics_detail_for_npu_id.fia_edm_stats_info")
}

func init() { proto.RegisterFile("fia_edm_stats_info.proto", fileDescriptor_d57f30811988eb7b) }

var fileDescriptor_d57f30811988eb7b = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0xc9, 0x8c, 0xbb, 0xec, 0x64, 0x76, 0x2e, 0x45, 0x31, 0x20, 0xeb, 0x8e, 0xf5, 0xd2,
	0x53, 0x0f, 0xb3, 0xfe, 0xb9, 0x0b, 0x0a, 0x83, 0xb8, 0x42, 0x85, 0x05, 0x4f, 0x2f, 0x69, 0x92,
	0x62, 0xd8, 0xb6, 0x29, 0x49, 0x5b, 0x3d, 0x7b, 0xf6, 0xa8, 0x9f, 0x44, 0xf7, 0x93, 0x88, 0xdf,
	0x47, 0xf2, 0x26, 0x5d, 0x47, 0x77, 0x41, 0x8f, 0xee, 0x65, 0x20, 0xbf, 0x3c, 0xf3, 0xf4, 0xcd,
	0x93, 0xa7, 0xa5, 0xac, 0xd2, 0x1c, 0x94, 0x6c, 0xc0, 0xf5, 0xbc, 0x77, 0xa0, 0xdb, 0xca, 0xe4,
	0x9d, 0x35, 0xbd, 0x49, 0x3e, 0x13, 0xa1, 0x9d, 0x30, 0xa0, 0x8d, 0x83, 0x0f, 0x16, 0x2a, 0xab,
	0xfa, 0x9e, 0x43, 0x29, 0x1a, 0x90, 0x1d, 0x07, 0xab, 0x9c, 0x19, 0xac, 0x50, 0x0e, 0x4c, 0xa7,
	0x6c, 0x2e, 0x3b, 0x9e, 0xa3, 0x43, 0xde, 0x1a, 0xa9, 0xc2, 0x6f, 0xce, 0x9d, 0x16, 0xe8, 0xab,
	0x5d, 0xaf, 0x85, 0xfb, 0x73, 0x0d, 0x52, 0xf5, 0x5c, 0xd7, 0x50, 0x19, 0x0b, 0x6d, 0x37, 0x80,
	0x96, 0xff, 0x20, 0x49, 0x5f, 0xd1, 0xbb, 0x57, 0x47, 0x86, 0x97, 0xcf, 0xdf, 0xbe, 0x49, 0xee,
	0xd1, 0x85, 0x7f, 0x2e, 0xb4, 0xbc, 0x51, 0x8c, 0xac, 0x49, 0xb6, 0x28, 0x0e, 0x3c, 0x38, 0xe5,
	0x8d, 0x4a, 0xee, 0xd0, 0xfd, 0xe0, 0xc0, 0x66, 0x6b, 0x92, 0xad, 0x8a, 0xbd, 0xb6, 0x1b, 0xb6,
	0x32, 0xfd, 0x48, 0xe8, 0xf1, 0xe4, 0x27, 0xd5, 0xa8, 0x85, 0x8a, 0xb6, 0x95, 0x56, 0xb5, 0x44,
	0xf3, 0xe4, 0x88, 0xd2, 0xb0, 0xda, 0x31, 0x5e, 0x20, 0x41, 0xe7, 0x63, 0xba, 0x0c, 0xdb, 0x23,
	0xaf, 0x07, 0x85, 0xf6, 0xb7, 0x8a, 0xf0, 0x8f, 0x33, 0x4f, 0xbc, 0x40, 0x3b, 0x30, 0xa3, 0xb2,
	0x55, 0x6d, 0xde, 0xb3, 0xf9, 0x9a, 0x64, 0x07, 0x05, 0xd5, 0xee, 0x75, 0x24, 0xe9, 0xc5, 0x8c,
	0x1e, 0x5d, 0x3b, 0x44, 0x59, 0x9f, 0x5f, 0x8e, 0x50, 0xd6, 0x46, 0x9c, 0xff, 0x36, 0x02, 0x12,
	0x1c, 0xe1, 0x07, 0x99, 0x46, 0xf4, 0x6a, 0x36, 0x5b, 0xcf, 0xb3, 0xe5, 0xe6, 0x82, 0xe4, 0xff,
	0xe3, 0x0d, 0xe6, 0x7f, 0x89, 0x3b, 0x46, 0xbb, 0x8d, 0xc7, 0x6e, 0x87, 0x26, 0x6c, 0x3a, 0x0c,
	0x6e, 0x55, 0x2c, 0xda, 0xa1, 0x79, 0x81, 0x20, 0xfd, 0x32, 0xa3, 0xf7, 0xaf, 0x75, 0xc3, 0x21,
	0x30, 0xb8, 0xef, 0x64, 0x4a, 0x0e, 0x93, 0x21, 0x98, 0xcc, 0xd7, 0x9b, 0x94, 0xcc, 0xd4, 0x81,
	0x78, 0xdf, 0xbb, 0xb9, 0x20, 0x70, 0xb1, 0xd0, 0x3e, 0x97, 0x67, 0x08, 0xd2, 0x4f, 0x73, 0x9a,
	0x5c, 0x7d, 0x49, 0x92, 0xdb, 0x74, 0x6f, 0xe4, 0xb5, 0x96, 0x6c, 0x83, 0x0d, 0x0c, 0x0b, 0xdf,
	0x4e, 0xcb, 0x7d, 0xb3, 0x86, 0xa6, 0x54, 0x96, 0x9d, 0xa0, 0x19, 0xf5, 0xe8, 0x14, 0x89, 0x17,
	0xb8, 0xda, 0xf4, 0x93, 0xe0, 0x51, 0x10, 0x78, 0x14, 0x05, 0x0f, 0xe9, 0x2a, 0x06, 0xee, 0x7a,
	0xde, 0x0a, 0xc5, 0x1e, 0xa3, 0xe4, 0xd0, 0xc3, 0x6d, 0x64, 0xc9, 0x03, 0x7a, 0x28, 0xde, 0xe9,
	0x0e, 0x46, 0x65, 0x9d, 0x36, 0x2d, 0x7b, 0x82, 0x9a, 0xa5, 0x67, 0x67, 0x01, 0x61, 0x8b, 0x7f,
	0xe5, 0xc3, 0x9e, 0xae, 0x49, 0xb6, 0xdc, 0x7c, 0xbb, 0x49, 0x77, 0x75, 0xd9, 0xbb, 0x62, 0xe7,
	0x20, 0xe5, 0x3e, 0x7e, 0x50, 0x4f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x30, 0x3d, 0xee,
	0x6c, 0x05, 0x00, 0x00,
}
