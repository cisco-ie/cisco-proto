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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_attributes_threshold_minor_hi_thresh_bag

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
	Name_8               string   `protobuf:"bytes,9,opt,name=name_8,json=name8,proto3" json:"name_8,omitempty"`
	Name_9               string   `protobuf:"bytes,10,opt,name=name_9,json=name9,proto3" json:"name_9,omitempty"`
	Name_10              string   `protobuf:"bytes,11,opt,name=name_10,json=name10,proto3" json:"name_10,omitempty"`
	Name_11              string   `protobuf:"bytes,12,opt,name=name_11,json=name11,proto3" json:"name_11,omitempty"`
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

func (m *EnvmonThresholdInfo_KEYS) GetName_8() string {
	if m != nil {
		return m.Name_8
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_9() string {
	if m != nil {
		return m.Name_9
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_10() string {
	if m != nil {
		return m.Name_10
	}
	return ""
}

func (m *EnvmonThresholdInfo_KEYS) GetName_11() string {
	if m != nil {
		return m.Name_11
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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x35, 0xa1, 0x4d, 0x5b, 0xf3, 0x27, 0x0c, 0x15, 0x57, 0x80, 0x50, 0xd4, 0x0d, 0xd9,
	0x60, 0x65, 0x7e, 0x93, 0xec, 0xc9, 0x0a, 0x89, 0x45, 0x2a, 0x21, 0xb1, 0xb2, 0x66, 0xa6, 0x6e,
	0x63, 0x31, 0x63, 0x57, 0xb6, 0x3b, 0x22, 0x4f, 0xc2, 0x73, 0xf0, 0x2a, 0x3c, 0x11, 0xf2, 0x9d,
	0xc8, 0x2e, 0x55, 0x37, 0x67, 0xce, 0x7c, 0xe7, 0xf8, 0x8e, 0xee, 0x48, 0x26, 0xef, 0x85, 0x1a,
	0x7a, 0xad, 0xb8, 0xdb, 0x19, 0x61, 0x77, 0xba, 0xbb, 0xe2, 0x52, 0x5d, 0x6b, 0x76, 0x6b, 0xb4,
	0xd3, 0xf4, 0x6f, 0xd2, 0x4a, 0xdb, 0x6a, 0x2e, 0xb5, 0xe5, 0xbf, 0x0c, 0x97, 0x6a, 0xe8, 0x6f,
	0x0c, 0xd7, 0xb7, 0xc2, 0x30, 0xa9, 0x06, 0xa1, 0x9c, 0x36, 0x7b, 0x66, 0xea, 0xf6, 0xa7, 0x45,
	0x65, 0x42, 0x39, 0xe9, 0xf6, 0xcc, 0x76, 0xda, 0x31, 0x67, 0x65, 0x6a, 0x51, 0xbd, 0x64, 0x68,
	0x33, 0x2f, 0x39, 0xda, 0xdc, 0x4b, 0x81, 0xb6, 0xf0, 0x52, 0xa2, 0x2d, 0xbd, 0x54, 0x68, 0x2b,
	0x2f, 0x4b, 0xb4, 0x4b, 0x2f, 0x2b, 0xb4, 0x2b, 0x2f, 0x6b, 0xb4, 0x6b, 0x1c, 0xbe, 0x18, 0xbf,
	0xb1, 0x60, 0xb5, 0x73, 0x46, 0x36, 0x77, 0x4e, 0x58, 0x16, 0x76, 0x61, 0xbd, 0x54, 0xda, 0xf0,
	0x9d, 0x3c, 0x20, 0xde, 0xd4, 0x37, 0x17, 0x7f, 0x26, 0xe4, 0xdd, 0xa3, 0x4b, 0xf3, 0xaf, 0x9b,
	0x1f, 0x97, 0x94, 0x92, 0x23, 0x55, 0xf7, 0x02, 0x92, 0x59, 0x32, 0x3f, 0xdb, 0xa2, 0xa7, 0xe7,
	0x64, 0xea, 0x9f, 0x3c, 0x85, 0x09, 0xd2, 0x63, 0xff, 0x96, 0x06, 0x9c, 0xc1, 0x93, 0x88, 0xb3,
	0x80, 0x73, 0x38, 0x8a, 0x38, 0x0f, 0xb8, 0x80, 0xe3, 0x88, 0x8b, 0x80, 0x4b, 0x98, 0x46, 0x5c,
	0x06, 0x5c, 0xc1, 0x49, 0xc4, 0x55, 0xc0, 0x4b, 0x38, 0x8d, 0x78, 0x19, 0xf0, 0x0a, 0xce, 0x22,
	0x5e, 0x05, 0xbc, 0x06, 0x12, 0xf1, 0x9a, 0xbe, 0x25, 0x27, 0xe3, 0x3a, 0x0b, 0x78, 0x8a, 0x1c,
	0x5b, 0xe9, 0x22, 0x06, 0x29, 0x3c, 0xbb, 0x17, 0xa4, 0x17, 0xbf, 0x27, 0xe4, 0xfc, 0xd1, 0x7f,
	0x46, 0x3f, 0x13, 0x1a, 0x89, 0x15, 0x83, 0x30, 0xd2, 0xed, 0x21, 0x9b, 0x25, 0xf3, 0xe7, 0xdb,
	0x57, 0x21, 0xb9, 0x3c, 0x04, 0xff, 0xd7, 0x8d, 0xe8, 0x6a, 0x27, 0xb5, 0x82, 0xfc, 0x41, 0x7d,
	0x7b, 0x08, 0xe8, 0x27, 0xf2, 0x32, 0xd6, 0x87, 0xba, 0xbb, 0x13, 0x50, 0x60, 0xf7, 0x45, 0xc0,
	0xdf, 0x3d, 0xa5, 0x29, 0x79, 0x13, 0x8b, 0xc2, 0x37, 0xc7, 0xc9, 0xe5, 0x2c, 0x99, 0x9f, 0x6e,
	0x5f, 0x87, 0x6c, 0x13, 0x22, 0xfa, 0x85, 0x7c, 0x8c, 0x47, 0x94, 0x76, 0xf2, 0x5a, 0xb6, 0x98,
	0x70, 0xa1, 0xea, 0xa6, 0x13, 0x57, 0x50, 0xe1, 0xe1, 0x0f, 0xa1, 0xf5, 0xed, 0x5e, 0x69, 0x33,
	0x76, 0x9a, 0x29, 0xde, 0x94, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x77, 0x38, 0x31,
	0x48, 0x03, 0x00, 0x00,
}