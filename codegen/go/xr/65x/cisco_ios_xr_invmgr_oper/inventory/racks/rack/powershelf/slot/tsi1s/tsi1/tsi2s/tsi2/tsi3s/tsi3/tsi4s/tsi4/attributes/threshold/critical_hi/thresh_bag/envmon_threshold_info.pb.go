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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_attributes_threshold_critical_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x99, 0x7e, 0x6d, 0xf9, 0x0c, 0xa8, 0x18, 0x2d, 0x04, 0x15, 0x29, 0xdd, 0xd8, 0x8d,
	0x81, 0xf9, 0xa9, 0x57, 0x60, 0x57, 0x82, 0x8b, 0x29, 0x08, 0xae, 0x42, 0x66, 0x7a, 0xda, 0x09,
	0x4e, 0x93, 0x92, 0xa4, 0xa3, 0xbd, 0x10, 0xf1, 0x02, 0xbc, 0x51, 0x99, 0xb4, 0x24, 0x2a, 0xdd,
	0xbc, 0x73, 0xe6, 0x79, 0x9f, 0x33, 0x9c, 0xc5, 0xa0, 0x2b, 0x90, 0xcd, 0x4a, 0x49, 0x66, 0x2b,
	0x0d, 0xa6, 0x52, 0xf5, 0x9c, 0x09, 0xb9, 0x50, 0x74, 0xad, 0x95, 0x55, 0xf8, 0x23, 0x2a, 0x85,
	0x29, 0x15, 0x13, 0xca, 0xb0, 0x77, 0xcd, 0x84, 0x6c, 0x56, 0x4b, 0xcd, 0xd4, 0x1a, 0x34, 0x15,
	0xb2, 0x01, 0x69, 0x95, 0xde, 0x52, 0xcd, 0xcb, 0x57, 0xe3, 0x92, 0xae, 0xd5, 0x1b, 0x68, 0x53,
	0x41, 0xbd, 0xa0, 0xa6, 0x56, 0x96, 0x5a, 0x23, 0x62, 0xe3, 0xb2, 0x8d, 0xc4, 0x8d, 0x49, 0x1b,
	0xa9, 0x1b, 0xd3, 0x36, 0x32, 0x37, 0x66, 0x94, 0x5b, 0xab, 0x45, 0xb1, 0xb1, 0x60, 0xa8, 0xbf,
	0x85, 0x96, 0x5a, 0x58, 0x51, 0xf2, 0x9a, 0x55, 0x62, 0x4f, 0x59, 0xc1, 0x97, 0xa3, 0xaf, 0x08,
	0x5d, 0x1e, 0xbc, 0x9b, 0x3d, 0x4e, 0x5f, 0x66, 0x18, 0xa3, 0xae, 0xe4, 0x2b, 0x20, 0xd1, 0x30,
	0x1a, 0x1f, 0xe5, 0x6e, 0xc6, 0x03, 0xd4, 0x6f, 0x9f, 0x2c, 0x26, 0x1d, 0x47, 0x7b, 0xed, 0x5b,
	0xec, 0x71, 0x42, 0xfe, 0x05, 0x9c, 0x78, 0x9c, 0x92, 0x6e, 0xc0, 0xa9, 0xc7, 0x19, 0xe9, 0x05,
	0x9c, 0x79, 0x3c, 0x21, 0xfd, 0x80, 0x27, 0xa3, 0xcf, 0x0e, 0x1a, 0x1c, 0xbc, 0x12, 0xdf, 0x21,
	0x1c, 0x88, 0x81, 0x06, 0xb4, 0xb0, 0x5b, 0x92, 0x0c, 0xa3, 0xf1, 0x71, 0x7e, 0xe6, 0x9b, 0xd9,
	0xbe, 0xf8, 0xad, 0x6b, 0xa8, 0xb9, 0x15, 0x4a, 0x92, 0xf4, 0x8f, 0x9e, 0xef, 0x0b, 0x7c, 0x8b,
	0x4e, 0x83, 0xde, 0xf0, 0x7a, 0x03, 0x24, 0x73, 0xee, 0x89, 0xc7, 0xcf, 0x2d, 0xc5, 0x31, 0xba,
	0x08, 0x22, 0xb4, 0xe6, 0xee, 0xcb, 0x93, 0x61, 0x34, 0xfe, 0x9f, 0x9f, 0xfb, 0x6e, 0xea, 0x2b,
	0xfc, 0x80, 0x6e, 0xc2, 0x8a, 0x54, 0x56, 0x2c, 0x44, 0xe9, 0x1a, 0x06, 0x92, 0x17, 0x35, 0xcc,
	0xc9, 0xbd, 0x5b, 0xbe, 0xf6, 0xd6, 0xd3, 0x0f, 0x69, 0xba, 0x73, 0x8a, 0xbe, 0xfb, 0xbd, 0xd2,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x78, 0x18, 0xd4, 0x7d, 0x02, 0x00, 0x00,
}
