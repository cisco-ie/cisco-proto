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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_attributes_threshold_critical_hi_thresh_bag

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
	Name_6               string   `protobuf:"bytes,7,opt,name=name_6,json=name6,proto3" json:"name_6,omitempty"`
	Name_7               string   `protobuf:"bytes,8,opt,name=name_7,json=name7,proto3" json:"name_7,omitempty"`
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

func (m *EnvmonThresholdInfo_KEYS) GetName_6() string {
	if m != nil {
		return m.Name_6
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_7() string {
	if m != nil {
		return m.Name_7
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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6a, 0x22, 0x31,
	0x00, 0xc6, 0x19, 0x57, 0x5d, 0x37, 0xb0, 0xbb, 0x34, 0xad, 0x10, 0xda, 0x52, 0xc4, 0x4b, 0xbd,
	0x34, 0x30, 0x7f, 0x7d, 0x81, 0x7a, 0x2a, 0xf4, 0x30, 0x42, 0xa1, 0xa7, 0x90, 0x19, 0xa3, 0x86,
	0x8e, 0x89, 0x24, 0x71, 0xa8, 0x4f, 0xd2, 0x47, 0xe8, 0xf3, 0xf4, 0x8d, 0x4a, 0xa2, 0x24, 0x6d,
	0xf1, 0xf2, 0xf9, 0xf9, 0xfb, 0x7e, 0x19, 0x26, 0x30, 0xe0, 0x8a, 0x89, 0x76, 0x23, 0x05, 0x31,
	0x6b, 0xc5, 0xf4, 0x5a, 0x36, 0x0b, 0xc2, 0xc5, 0x52, 0xe2, 0xad, 0x92, 0x46, 0xc2, 0xf7, 0xa8,
	0xe6, 0xba, 0x96, 0x84, 0x4b, 0x4d, 0x5e, 0x15, 0xe1, 0xa2, 0xdd, 0xac, 0x14, 0x91, 0x5b, 0xa6,
	0x30, 0x17, 0x2d, 0x13, 0x46, 0xaa, 0x3d, 0x56, 0xb4, 0x7e, 0xd1, 0x2e, 0x31, 0x13, 0x86, 0x9b,
	0x3d, 0xd6, 0x8d, 0x34, 0xd8, 0x68, 0x1e, 0x6b, 0x97, 0x36, 0x12, 0x57, 0x13, 0x1b, 0xa9, 0xab,
	0xa9, 0x8d, 0xcc, 0xd5, 0xcc, 0x46, 0xee, 0x6a, 0x6e, 0xa3, 0x70, 0xb5, 0xc0, 0xd4, 0x18, 0xc5,
	0xab, 0x9d, 0x61, 0x1a, 0xfb, 0xb7, 0xc3, 0xb5, 0xe2, 0x86, 0xd7, 0xb4, 0x21, 0x6b, 0x7e, 0xa4,
	0xa4, 0xa2, 0xab, 0xf1, 0x47, 0x04, 0x2e, 0x4f, 0xde, 0x84, 0x3c, 0xcc, 0x9e, 0xe7, 0x10, 0x82,
	0xae, 0xa0, 0x1b, 0x86, 0xa2, 0x51, 0x34, 0xf9, 0x53, 0xba, 0x0e, 0x87, 0xa0, 0x6f, 0x7f, 0x49,
	0x8c, 0x3a, 0x8e, 0xf6, 0xec, 0xbf, 0xd8, 0xe3, 0x04, 0xfd, 0x0a, 0x38, 0xf1, 0x38, 0x45, 0xdd,
	0x80, 0x53, 0x8f, 0x33, 0xd4, 0x0b, 0x38, 0xf3, 0x38, 0x47, 0xfd, 0x80, 0x73, 0x8f, 0x0b, 0xf4,
	0x3b, 0xe0, 0xc2, 0xe3, 0x29, 0x1a, 0x04, 0x3c, 0x1d, 0xbf, 0x75, 0xc0, 0xf0, 0xe4, 0x9d, 0xe0,
	0x1d, 0x80, 0x81, 0x68, 0xd6, 0x32, 0xc5, 0xcd, 0x1e, 0x25, 0xa3, 0x68, 0xf2, 0xb7, 0x3c, 0xf3,
	0xcb, 0xfc, 0x38, 0x7c, 0xd7, 0x15, 0x6b, 0xa8, 0xe1, 0x52, 0xa0, 0xf4, 0x87, 0x5e, 0x1e, 0x07,
	0x78, 0x0b, 0xfe, 0x07, 0xbd, 0xa5, 0xcd, 0x8e, 0xa1, 0xcc, 0xb9, 0xff, 0x3c, 0x7e, 0xb2, 0x14,
	0xc6, 0xe0, 0x22, 0x88, 0xcc, 0x9a, 0x87, 0x27, 0xe7, 0xa3, 0x68, 0x32, 0x28, 0xcf, 0xfd, 0x36,
	0xf3, 0x13, 0xbc, 0x07, 0x37, 0xe1, 0x88, 0x90, 0x86, 0x2f, 0x79, 0xed, 0x16, 0xc2, 0x04, 0xad,
	0x1a, 0xb6, 0x40, 0x85, 0x3b, 0x7c, 0xed, 0xad, 0xc7, 0x2f, 0xd2, 0xec, 0xe0, 0x54, 0x7d, 0xf7,
	0x79, 0xa6, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xb7, 0xb8, 0x73, 0xbd, 0x02, 0x00, 0x00,
}
