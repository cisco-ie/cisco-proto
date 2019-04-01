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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_tsi11s_tsi11_attributes_threshold_major_hi_thresh_bag

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
	Name_12              string   `protobuf:"bytes,13,opt,name=name_12,json=name12,proto3" json:"name_12,omitempty"`
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

func (m *EnvmonThresholdInfo_KEYS) GetName_12() string {
	if m != nil {
		return m.Name_12
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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.tsi11s.tsi11.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x35, 0xa1, 0x4d, 0x5b, 0x43, 0x41, 0x18, 0x2a, 0xae, 0x00, 0xa1, 0xa8, 0x1b, 0xb2,
	0xc1, 0xca, 0xfc, 0x26, 0xd9, 0x93, 0x15, 0x12, 0x8b, 0x54, 0x42, 0x62, 0x65, 0x39, 0xa9, 0xd3,
	0x18, 0x26, 0x76, 0x65, 0xbb, 0x23, 0xf2, 0x24, 0xbc, 0x13, 0x4f, 0xc1, 0xa3, 0x20, 0xdf, 0x44,
	0x76, 0xa8, 0xb2, 0x39, 0x73, 0xe6, 0x3b, 0xc7, 0xb6, 0x7c, 0x25, 0x93, 0x77, 0x52, 0x77, 0x1b,
	0xa3, 0xb9, 0x5f, 0x5b, 0xe9, 0xd6, 0xa6, 0xbd, 0xe5, 0x4a, 0xaf, 0x0c, 0xbb, 0xb7, 0xc6, 0x1b,
	0xfa, 0x37, 0x5b, 0x2a, 0xb7, 0x34, 0x5c, 0x19, 0xc7, 0x7f, 0x59, 0xae, 0x74, 0xb7, 0xb9, 0xb3,
	0xdc, 0xdc, 0x4b, 0xcb, 0x94, 0xee, 0xa4, 0xf6, 0xc6, 0x6e, 0x99, 0x15, 0xcb, 0x9f, 0x0e, 0x95,
	0xad, 0x84, 0xf6, 0x56, 0x6c, 0x99, 0x6b, 0x8d, 0x67, 0xde, 0xa9, 0xdc, 0xa1, 0x06, 0x29, 0xd0,
	0x16, 0x41, 0x4a, 0xb4, 0x65, 0x90, 0x0a, 0x6d, 0x15, 0xa4, 0x46, 0x5b, 0x07, 0x69, 0xd0, 0x36,
	0x41, 0xc6, 0x68, 0xc7, 0x41, 0x26, 0x68, 0x27, 0x41, 0xa6, 0x68, 0xa7, 0xb8, 0xf9, 0x68, 0x77,
	0xc6, 0x08, 0x75, 0x7f, 0x60, 0xce, 0x84, 0xf7, 0x56, 0x2d, 0x1e, 0xbc, 0x74, 0x2c, 0xde, 0x8c,
	0x6d, 0xc4, 0x0f, 0x63, 0xf9, 0x5a, 0xed, 0x11, 0x5f, 0x88, 0xbb, 0xeb, 0x3f, 0x3d, 0xf2, 0xf6,
	0xe8, 0x08, 0xf8, 0x97, 0xd9, 0xf7, 0x1b, 0x4a, 0xc9, 0x89, 0x16, 0x1b, 0x09, 0xd9, 0x20, 0x1b,
	0x5e, 0xcc, 0xd1, 0xd3, 0x2b, 0xd2, 0x0f, 0x5f, 0x9e, 0x43, 0x0f, 0xe9, 0x69, 0xf8, 0xcb, 0x23,
	0x2e, 0xe0, 0x49, 0xc2, 0x45, 0xc4, 0x25, 0x9c, 0x24, 0x5c, 0x46, 0x5c, 0xc1, 0x69, 0xc2, 0x55,
	0xc4, 0x35, 0xf4, 0x13, 0xae, 0x23, 0x6e, 0xe0, 0x2c, 0xe1, 0x26, 0xe2, 0x31, 0x9c, 0x27, 0x3c,
	0x8e, 0x78, 0x02, 0x17, 0x09, 0x4f, 0x22, 0x9e, 0x02, 0x49, 0x78, 0x4a, 0xdf, 0x90, 0xb3, 0xdd,
	0x75, 0x46, 0xf0, 0x14, 0x39, 0xb6, 0xf2, 0x51, 0x0a, 0x72, 0x78, 0x76, 0x10, 0xe4, 0x29, 0x28,
	0xe0, 0xf2, 0x20, 0x28, 0xae, 0x7f, 0xf7, 0xc8, 0xd5, 0xd1, 0x61, 0xd2, 0x4f, 0x84, 0x26, 0xe2,
	0x64, 0x27, 0xad, 0xf2, 0x5b, 0x28, 0x06, 0xd9, 0xf0, 0x72, 0xfe, 0x32, 0x26, 0x37, 0xfb, 0xe0,
	0xff, 0xba, 0x95, 0xad, 0xf0, 0xca, 0x68, 0x28, 0x1f, 0xd5, 0xe7, 0xfb, 0x80, 0x7e, 0x24, 0x2f,
	0x52, 0xbd, 0x13, 0xed, 0x83, 0x84, 0x0a, 0xbb, 0xcf, 0x23, 0xfe, 0x16, 0x28, 0xcd, 0xc9, 0xeb,
	0x54, 0x94, 0xa1, 0xb9, 0xdb, 0xb9, 0x1e, 0x64, 0xc3, 0xf3, 0xf9, 0xab, 0x98, 0xcd, 0x62, 0x44,
	0x3f, 0x93, 0x0f, 0x69, 0x89, 0x36, 0x5e, 0xad, 0xd4, 0x12, 0x13, 0x2e, 0xb5, 0x58, 0xb4, 0xf2,
	0x16, 0x1a, 0x5c, 0xfc, 0x3e, 0xb6, 0xbe, 0x1e, 0x94, 0x66, 0xbb, 0xce, 0xa2, 0x8f, 0x0f, 0xaa,
	0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x30, 0xc9, 0x7c, 0x6f, 0x03, 0x00, 0x00,
}
