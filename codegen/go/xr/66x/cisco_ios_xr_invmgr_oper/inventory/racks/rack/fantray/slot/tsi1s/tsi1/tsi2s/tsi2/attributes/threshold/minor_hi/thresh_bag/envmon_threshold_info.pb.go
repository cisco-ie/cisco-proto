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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_threshold_minor_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xe9, 0xbe, 0x7d, 0x43, 0x03, 0x2a, 0x46, 0x07, 0x41, 0x45, 0xc6, 0x6e, 0xdc, 0x8d,
	0x81, 0xb6, 0xea, 0x13, 0xb8, 0x2b, 0xc1, 0x8b, 0x0e, 0x04, 0xaf, 0x42, 0xda, 0x65, 0xdb, 0xc1,
	0x36, 0x67, 0x24, 0x59, 0xb1, 0x4f, 0xe2, 0xeb, 0x4a, 0xb3, 0x92, 0xaa, 0xec, 0xe6, 0xf4, 0xf4,
	0xf7, 0xff, 0x9d, 0x9c, 0x5c, 0x84, 0x5c, 0x2b, 0x5d, 0x57, 0xa8, 0x85, 0xdb, 0x18, 0x65, 0x37,
	0x58, 0x2e, 0x05, 0xe8, 0x15, 0xf2, 0xad, 0x41, 0x87, 0xb4, 0x29, 0xc0, 0x16, 0x28, 0x00, 0xad,
	0xf8, 0x34, 0x02, 0x74, 0x5d, 0xad, 0x8d, 0xc0, 0xad, 0x32, 0x1c, 0x74, 0xad, 0xb4, 0x43, 0xd3,
	0x70, 0x23, 0x8b, 0x0f, 0xeb, 0x2b, 0x5f, 0x49, 0xed, 0x8c, 0x6c, 0xb8, 0x2d, 0xd1, 0x71, 0x67,
	0x21, 0xb6, 0xbe, 0xb6, 0x25, 0xf1, 0x6d, 0xc2, 0xa5, 0x73, 0x06, 0xf2, 0x9d, 0x53, 0x96, 0x87,
	0x8d, 0xbc, 0x02, 0x8d, 0x46, 0x6c, 0xa0, 0x43, 0x22, 0x97, 0xeb, 0xe9, 0x8e, 0x5c, 0x1d, 0xbc,
	0x99, 0x78, 0x99, 0xbf, 0x2f, 0x28, 0x25, 0x43, 0x2d, 0x2b, 0xc5, 0xa2, 0x49, 0x34, 0x3b, 0xce,
	0x7c, 0x4f, 0xc7, 0x64, 0xd4, 0x7e, 0x45, 0xcc, 0x06, 0x9e, 0xfe, 0x6f, 0xff, 0xe2, 0x80, 0x13,
	0xf6, 0xaf, 0xc7, 0x49, 0xc0, 0x29, 0x1b, 0xf6, 0x38, 0x9d, 0x7e, 0x0d, 0xc8, 0xf8, 0xe0, 0x5e,
	0x7a, 0x4f, 0x68, 0x4f, 0xac, 0xaa, 0x95, 0x01, 0xd7, 0xb0, 0x64, 0x12, 0xcd, 0x4e, 0xb2, 0xf3,
	0x90, 0x2c, 0xba, 0xe0, 0xb7, 0x6e, 0x54, 0x29, 0x1d, 0xa0, 0x66, 0xe9, 0x1f, 0x3d, 0xeb, 0x02,
	0x7a, 0x47, 0xce, 0x7a, 0xbd, 0x96, 0xe5, 0x4e, 0xb1, 0x07, 0xef, 0x9e, 0x06, 0xfc, 0xd6, 0x52,
	0x1a, 0x93, 0xcb, 0x5e, 0x54, 0xad, 0xb9, 0x3f, 0xf9, 0x71, 0x12, 0xcd, 0x8e, 0xb2, 0x8b, 0x90,
	0xcd, 0x43, 0x44, 0x9f, 0xc9, 0x6d, 0x3f, 0xa2, 0xd1, 0xc1, 0x0a, 0x0a, 0x9f, 0x08, 0xa5, 0x65,
	0x5e, 0xaa, 0x25, 0x7b, 0xf2, 0xc3, 0x37, 0xc1, 0x7a, 0xfd, 0x21, 0xcd, 0xf7, 0x4e, 0x3e, 0xf2,
	0x4f, 0x22, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xa8, 0x54, 0x16, 0x31, 0x02, 0x00, 0x00,
}