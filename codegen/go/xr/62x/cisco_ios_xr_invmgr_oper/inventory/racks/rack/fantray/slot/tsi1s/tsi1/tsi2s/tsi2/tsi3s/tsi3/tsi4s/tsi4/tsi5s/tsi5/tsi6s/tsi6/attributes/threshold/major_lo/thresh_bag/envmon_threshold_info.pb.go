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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_tsi4s_tsi4_tsi5s_tsi5_tsi6s_tsi6_attributes_threshold_major_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.tsi4s.tsi4.tsi5s.tsi5.tsi6s.tsi6.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6a, 0x22, 0x31,
	0x00, 0xc6, 0x19, 0x57, 0x5d, 0x37, 0xb0, 0xbb, 0x6c, 0x76, 0x85, 0xb0, 0xbb, 0x14, 0xf1, 0x52,
	0x2f, 0x0d, 0xcc, 0x5f, 0x5f, 0xa0, 0x9e, 0x0a, 0x3d, 0x8c, 0x50, 0xe8, 0x29, 0x64, 0xc6, 0xa8,
	0x69, 0x67, 0x12, 0x49, 0xe2, 0x50, 0x9f, 0xa4, 0x4f, 0xd0, 0x07, 0xea, 0x1b, 0x95, 0x44, 0x49,
	0xda, 0xe2, 0xe5, 0xf3, 0xf3, 0xf7, 0xfd, 0x32, 0x4c, 0x60, 0xc0, 0x3f, 0x26, 0xba, 0x56, 0x0a,
	0x62, 0xb6, 0x8a, 0xe9, 0xad, 0x6c, 0x56, 0x84, 0x8b, 0xb5, 0xc4, 0x3b, 0x25, 0x8d, 0x84, 0x2f,
	0x51, 0xcd, 0x75, 0x2d, 0x09, 0x97, 0x9a, 0x3c, 0x29, 0xc2, 0x45, 0xd7, 0x6e, 0x14, 0x91, 0x3b,
	0xa6, 0x30, 0x17, 0x1d, 0x13, 0x46, 0xaa, 0x03, 0x56, 0xb4, 0x7e, 0xd4, 0x2e, 0xf1, 0x9a, 0x0a,
	0xa3, 0xe8, 0x01, 0xeb, 0x46, 0x1a, 0x6c, 0x34, 0x8f, 0xb5, 0x4b, 0x1b, 0x89, 0xab, 0x89, 0x8d,
	0xd4, 0xd5, 0xd4, 0x46, 0xe6, 0x6a, 0x66, 0x23, 0x77, 0x35, 0xb7, 0x51, 0xb8, 0x5a, 0x60, 0x6a,
	0x8c, 0xe2, 0xd5, 0xde, 0x30, 0x8d, 0xfd, 0xeb, 0xe1, 0x96, 0x3e, 0x48, 0x45, 0x1a, 0x79, 0x42,
	0xa4, 0xa2, 0x9b, 0xe9, 0x6b, 0x04, 0xfe, 0x9e, 0xbd, 0x07, 0xb9, 0x59, 0xdc, 0x2f, 0x21, 0x04,
	0x7d, 0x41, 0x5b, 0x86, 0xa2, 0x49, 0x34, 0xfb, 0x56, 0xba, 0x0e, 0xc7, 0x60, 0x68, 0x7f, 0x49,
	0x8c, 0x7a, 0x8e, 0x0e, 0xec, 0xbf, 0xd8, 0xe3, 0x04, 0x7d, 0x09, 0x38, 0xf1, 0x38, 0x45, 0xfd,
	0x80, 0x53, 0x8f, 0x33, 0x34, 0x08, 0x38, 0xf3, 0x38, 0x47, 0xc3, 0x80, 0x73, 0x8f, 0x0b, 0xf4,
	0x35, 0xe0, 0xc2, 0xe3, 0x39, 0x1a, 0x05, 0x3c, 0x9f, 0x3e, 0xf7, 0xc0, 0xf8, 0xec, 0x9d, 0xe0,
	0x15, 0x80, 0x81, 0x68, 0xd6, 0x31, 0xc5, 0xcd, 0x01, 0x25, 0x93, 0x68, 0xf6, 0xbd, 0xfc, 0xe5,
	0x97, 0xe5, 0x69, 0xf8, 0xa8, 0x2b, 0xd6, 0x50, 0xc3, 0xa5, 0x40, 0xe9, 0x27, 0xbd, 0x3c, 0x0d,
	0xf0, 0x12, 0xfc, 0x0c, 0x7a, 0x47, 0x9b, 0x3d, 0x43, 0x99, 0x73, 0x7f, 0x78, 0x7c, 0x67, 0x29,
	0x8c, 0xc1, 0x9f, 0x20, 0x32, 0x6b, 0x1e, 0x9f, 0x9c, 0x4f, 0xa2, 0xd9, 0xa8, 0xfc, 0xed, 0xb7,
	0x85, 0x9f, 0xe0, 0x35, 0xb8, 0x08, 0x47, 0x84, 0x34, 0x7c, 0xcd, 0x6b, 0xb7, 0x10, 0x26, 0x68,
	0xd5, 0xb0, 0x15, 0x2a, 0xdc, 0xe1, 0xff, 0xde, 0xba, 0x7d, 0x27, 0x2d, 0x8e, 0x4e, 0x35, 0x74,
	0x1f, 0x67, 0xfa, 0x16, 0x00, 0x00, 0xff, 0xff, 0x12, 0x12, 0x68, 0xff, 0xbb, 0x02, 0x00, 0x00,
}
