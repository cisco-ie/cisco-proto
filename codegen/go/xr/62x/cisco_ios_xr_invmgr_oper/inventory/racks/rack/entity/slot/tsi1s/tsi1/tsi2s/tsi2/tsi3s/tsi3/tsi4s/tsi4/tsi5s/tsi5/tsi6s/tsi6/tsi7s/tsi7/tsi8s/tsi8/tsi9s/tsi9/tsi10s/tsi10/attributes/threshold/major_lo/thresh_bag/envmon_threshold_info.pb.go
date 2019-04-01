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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_tsi7s_tsi7_tsi8s_tsi8_tsi9s_tsi9_tsi10s_tsi10_attributes_threshold_major_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.tsi7s.tsi7.tsi8s.tsi8.tsi9s.tsi9.tsi10s.tsi10.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x35, 0xa1, 0x4d, 0x5b, 0xf3, 0x27, 0x0c, 0x15, 0x57, 0x80, 0x50, 0xd4, 0x0d, 0xd9,
	0x60, 0x65, 0x7e, 0x93, 0xec, 0xc9, 0x0a, 0x89, 0x45, 0x2a, 0x21, 0xb1, 0xb2, 0x9c, 0xd4, 0x6d,
	0x0d, 0x33, 0x76, 0x65, 0xbb, 0x23, 0xf2, 0x24, 0x3c, 0x07, 0xaf, 0xc2, 0x13, 0x21, 0xdf, 0x89,
	0xec, 0x82, 0xb2, 0x39, 0x73, 0xe6, 0x3b, 0xc7, 0x77, 0x74, 0x47, 0x32, 0x79, 0x2b, 0x75, 0xdf,
	0x19, 0xcd, 0xfd, 0xad, 0x95, 0xee, 0xd6, 0xb4, 0x57, 0x5c, 0xe9, 0x6b, 0xc3, 0xee, 0xac, 0xf1,
	0x86, 0xfe, 0xc9, 0xb6, 0xca, 0x6d, 0x0d, 0x57, 0xc6, 0xf1, 0x9f, 0x96, 0x2b, 0xdd, 0x77, 0x37,
	0x96, 0x9b, 0x3b, 0x69, 0x99, 0xd2, 0xbd, 0xd4, 0xde, 0xd8, 0x1d, 0xb3, 0x62, 0xfb, 0xc3, 0xa1,
	0x32, 0xa9, 0xbd, 0xf2, 0x3b, 0xe6, 0x5a, 0xe3, 0x99, 0x77, 0x2a, 0x77, 0xa8, 0x41, 0x0a, 0xb4,
	0x45, 0x90, 0x12, 0x6d, 0x19, 0xa4, 0x42, 0x5b, 0x05, 0xa9, 0xd1, 0xd6, 0x41, 0x1a, 0xb4, 0x4d,
	0x90, 0x39, 0xda, 0x79, 0x90, 0x05, 0xda, 0x45, 0x90, 0x25, 0xda, 0x25, 0x0e, 0x9f, 0x0d, 0xdf,
	0x98, 0x31, 0xe1, 0xbd, 0x55, 0x9b, 0x7b, 0x2f, 0x1d, 0x8b, 0xbb, 0xb0, 0x4e, 0x7c, 0x37, 0x96,
	0xb7, 0x66, 0x8f, 0xf8, 0x46, 0xdc, 0x5c, 0xfc, 0x1e, 0x91, 0x37, 0x07, 0x97, 0xe6, 0x9f, 0x57,
	0xdf, 0x2e, 0x29, 0x25, 0x47, 0x5a, 0x74, 0x12, 0xb2, 0x49, 0x36, 0x3d, 0x5b, 0xa3, 0xa7, 0xe7,
	0x64, 0x1c, 0x9e, 0x3c, 0x87, 0x11, 0xd2, 0xe3, 0xf0, 0x96, 0x47, 0x5c, 0xc0, 0xa3, 0x84, 0x8b,
	0x88, 0x4b, 0x38, 0x4a, 0xb8, 0x8c, 0xb8, 0x82, 0xe3, 0x84, 0xab, 0x88, 0x6b, 0x18, 0x27, 0x5c,
	0x47, 0xdc, 0xc0, 0x49, 0xc2, 0x4d, 0xc4, 0x73, 0x38, 0x4d, 0x78, 0x1e, 0xf1, 0x02, 0xce, 0x12,
	0x5e, 0x44, 0xbc, 0x04, 0x92, 0xf0, 0x92, 0xbe, 0x26, 0x27, 0xc3, 0x3a, 0x33, 0x78, 0x8c, 0x1c,
	0x5b, 0xf9, 0x2c, 0x05, 0x39, 0x3c, 0x79, 0x10, 0xe4, 0x17, 0xbf, 0x46, 0xe4, 0xfc, 0xe0, 0x3f,
	0xa3, 0x1f, 0x09, 0x4d, 0xc4, 0xc9, 0x5e, 0x5a, 0xe5, 0x77, 0x50, 0x4c, 0xb2, 0xe9, 0xd3, 0xf5,
	0x8b, 0x98, 0x5c, 0xee, 0x83, 0x7f, 0xeb, 0x56, 0xb6, 0xc2, 0x2b, 0xa3, 0xa1, 0xfc, 0xaf, 0xbe,
	0xde, 0x07, 0xf4, 0x03, 0x79, 0x9e, 0xea, 0xbd, 0x68, 0xef, 0x25, 0x54, 0xd8, 0x7d, 0x16, 0xf1,
	0xd7, 0x40, 0x69, 0x4e, 0x5e, 0xa5, 0xa2, 0x0c, 0xcd, 0x61, 0x72, 0x3d, 0xc9, 0xa6, 0xa7, 0xeb,
	0x97, 0x31, 0x5b, 0xc5, 0x88, 0x7e, 0x22, 0xef, 0xd3, 0x11, 0x6d, 0xbc, 0xba, 0x56, 0x5b, 0x4c,
	0xb8, 0xd4, 0x62, 0xd3, 0xca, 0x2b, 0x68, 0xf0, 0xf0, 0xbb, 0xd8, 0xfa, 0xf2, 0xa0, 0xb4, 0x1a,
	0x3a, 0x9b, 0x31, 0xde, 0x94, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x13, 0xda, 0xc1,
	0x48, 0x03, 0x00, 0x00,
}
