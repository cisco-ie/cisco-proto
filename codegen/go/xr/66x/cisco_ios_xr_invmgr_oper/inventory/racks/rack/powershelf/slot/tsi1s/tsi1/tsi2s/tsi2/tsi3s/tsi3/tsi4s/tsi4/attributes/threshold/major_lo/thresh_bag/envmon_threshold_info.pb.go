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
// source: envmon_threshold_info.proto

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_attributes_threshold_major_lo_thresh_bag

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

type EnvmonThresholdInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	Name_2               string   `protobuf:"bytes,3,opt,name=name_2,json=name2,proto3" json:"name_2,omitempty"`
	Name_3               string   `protobuf:"bytes,4,opt,name=name_3,json=name3,proto3" json:"name_3,omitempty"`
	Name_4               string   `protobuf:"bytes,5,opt,name=name_4,json=name4,proto3" json:"name_4,omitempty"`
	Name_5               string   `protobuf:"bytes,6,opt,name=name_5,json=name5,proto3" json:"name_5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvmonThresholdInfo_KEYS) Reset()         { *m = EnvmonThresholdInfo_KEYS{} }
func (m *EnvmonThresholdInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*EnvmonThresholdInfo_KEYS) ProtoMessage()    {}
func (*EnvmonThresholdInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7410abf38d6220ea, []int{0}
}

func (m *EnvmonThresholdInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonThresholdInfo_KEYS.Unmarshal(m, b)
}
func (m *EnvmonThresholdInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonThresholdInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *EnvmonThresholdInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonThresholdInfo_KEYS.Merge(m, src)
}
func (m *EnvmonThresholdInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_EnvmonThresholdInfo_KEYS.Size(m)
}
func (m *EnvmonThresholdInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonThresholdInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonThresholdInfo_KEYS proto.InternalMessageInfo

func (m *EnvmonThresholdInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_2() string {
	if m != nil {
		return m.Name_2
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_3() string {
	if m != nil {
		return m.Name_3
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_4() string {
	if m != nil {
		return m.Name_4
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_5() string {
	if m != nil {
		return m.Name_5
	}
	return ""
}

type EnvmonThresholdInfo struct {
	ThresholdSeverity            uint32   `protobuf:"varint,50,opt,name=threshold_severity,json=thresholdSeverity,proto3" json:"threshold_severity,omitempty"`
	ThresholdRelation            uint32   `protobuf:"varint,51,opt,name=threshold_relation,json=thresholdRelation,proto3" json:"threshold_relation,omitempty"`
	ThresholdValue               uint32   `protobuf:"varint,52,opt,name=threshold_value,json=thresholdValue,proto3" json:"threshold_value,omitempty"`
	ThresholdEvaluation          bool     `protobuf:"varint,53,opt,name=threshold_evaluation,json=thresholdEvaluation,proto3" json:"threshold_evaluation,omitempty"`
	ThresholdNotificationEnabled bool     `protobuf:"varint,54,opt,name=threshold_notification_enabled,json=thresholdNotificationEnabled,proto3" json:"threshold_notification_enabled,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *EnvmonThresholdInfo) Reset()         { *m = EnvmonThresholdInfo{} }
func (m *EnvmonThresholdInfo) String() string { return proto.CompactTextString(m) }
func (*EnvmonThresholdInfo) ProtoMessage()    {}
func (*EnvmonThresholdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7410abf38d6220ea, []int{1}
}

func (m *EnvmonThresholdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvmonThresholdInfo.Unmarshal(m, b)
}
func (m *EnvmonThresholdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvmonThresholdInfo.Marshal(b, m, deterministic)
}
func (m *EnvmonThresholdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvmonThresholdInfo.Merge(m, src)
}
func (m *EnvmonThresholdInfo) XXX_Size() int {
	return xxx_messageInfo_EnvmonThresholdInfo.Size(m)
}
func (m *EnvmonThresholdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvmonThresholdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EnvmonThresholdInfo proto.InternalMessageInfo

func (m *EnvmonThresholdInfo) GetThresholdSeverity() uint32 {
	if m != nil {
		return m.ThresholdSeverity
	}
	return 0
}

func (m *EnvmonThresholdInfo) GetThresholdRelation() uint32 {
	if m != nil {
		return m.ThresholdRelation
	}
	return 0
}

func (m *EnvmonThresholdInfo) GetThresholdValue() uint32 {
	if m != nil {
		return m.ThresholdValue
	}
	return 0
}

func (m *EnvmonThresholdInfo) GetThresholdEvaluation() bool {
	if m != nil {
		return m.ThresholdEvaluation
	}
	return false
}

func (m *EnvmonThresholdInfo) GetThresholdNotificationEnabled() bool {
	if m != nil {
		return m.ThresholdNotificationEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x4a, 0x2b, 0x31,
	0x14, 0xc6, 0x99, 0xde, 0xb6, 0xdc, 0x1b, 0xb8, 0x8a, 0xd1, 0x42, 0x50, 0x91, 0xd2, 0x8d, 0xdd,
	0x18, 0x98, 0x3f, 0xf5, 0x09, 0xec, 0x4a, 0x70, 0x31, 0x05, 0xc1, 0x55, 0xc8, 0x4c, 0xd3, 0x36,
	0x3a, 0x93, 0x53, 0x92, 0x74, 0xb4, 0xaf, 0xe1, 0xc6, 0x07, 0xf0, 0x45, 0x65, 0xd2, 0x92, 0xa8,
	0x74, 0xf3, 0xcd, 0x99, 0xdf, 0xf7, 0x3b, 0xc3, 0x59, 0x0c, 0xba, 0x10, 0xaa, 0xa9, 0x41, 0x31,
	0xbb, 0xd2, 0xc2, 0xac, 0xa0, 0x9a, 0x33, 0xa9, 0x16, 0x40, 0xd7, 0x1a, 0x2c, 0xe0, 0xf7, 0xa8,
	0x94, 0xa6, 0x04, 0x26, 0xc1, 0xb0, 0x37, 0xcd, 0xa4, 0x6a, 0xea, 0xa5, 0x66, 0xb0, 0x16, 0x9a,
	0x4a, 0xd5, 0x08, 0x65, 0x41, 0x6f, 0xa9, 0xe6, 0xe5, 0x8b, 0x71, 0x49, 0xd7, 0xf0, 0x2a, 0xb4,
	0x59, 0x89, 0x6a, 0x41, 0x4d, 0x05, 0x96, 0x5a, 0x23, 0x63, 0xe3, 0xb2, 0x8d, 0xc4, 0x8d, 0x49,
	0x1b, 0xa9, 0x1b, 0xd3, 0x36, 0x32, 0x37, 0x66, 0x94, 0x5b, 0xab, 0x65, 0xb1, 0xb1, 0xc2, 0x50,
	0x7f, 0x0b, 0xad, 0xf9, 0x33, 0x68, 0x56, 0xc1, 0x1e, 0xb1, 0x82, 0x2f, 0x47, 0x9f, 0x11, 0x3a,
	0x3f, 0x78, 0x34, 0xbb, 0x9f, 0x3e, 0xcd, 0x30, 0x46, 0x5d, 0xc5, 0x6b, 0x41, 0xa2, 0x61, 0x34,
	0xfe, 0x97, 0xbb, 0x19, 0x0f, 0x50, 0xbf, 0x7d, 0xb2, 0x98, 0x74, 0x1c, 0xed, 0xb5, 0x6f, 0xb1,
	0xc7, 0x09, 0xf9, 0x13, 0x70, 0xe2, 0x71, 0x4a, 0xba, 0x01, 0xa7, 0x1e, 0x67, 0xa4, 0x17, 0x70,
	0xe6, 0xf1, 0x84, 0xf4, 0x03, 0x9e, 0x8c, 0x3e, 0x3a, 0x68, 0x70, 0xf0, 0x4a, 0x7c, 0x83, 0x70,
	0x20, 0x46, 0x34, 0x42, 0x4b, 0xbb, 0x25, 0xc9, 0x30, 0x1a, 0xff, 0xcf, 0x4f, 0x7c, 0x33, 0xdb,
	0x17, 0x3f, 0x75, 0x2d, 0x2a, 0x6e, 0x25, 0x28, 0x92, 0xfe, 0xd2, 0xf3, 0x7d, 0x81, 0xaf, 0xd1,
	0x71, 0xd0, 0x1b, 0x5e, 0x6d, 0x04, 0xc9, 0x9c, 0x7b, 0xe4, 0xf1, 0x63, 0x4b, 0x71, 0x8c, 0xce,
	0x82, 0x28, 0x5a, 0x73, 0xf7, 0xe5, 0xc9, 0x30, 0x1a, 0xff, 0xcd, 0x4f, 0x7d, 0x37, 0xf5, 0x15,
	0xbe, 0x43, 0x57, 0x61, 0x45, 0x81, 0x95, 0x0b, 0x59, 0xba, 0x86, 0x09, 0xc5, 0x8b, 0x4a, 0xcc,
	0xc9, 0xad, 0x5b, 0xbe, 0xf4, 0xd6, 0xc3, 0x37, 0x69, 0xba, 0x73, 0x8a, 0xbe, 0xfb, 0xb7, 0xd2,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0xc8, 0x95, 0x44, 0x7a, 0x02, 0x00, 0x00,
}
