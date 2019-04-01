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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_threshold_critical_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.threshold.critical_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x99, 0xde, 0xb6, 0xdc, 0x1b, 0xb8, 0xf7, 0x62, 0xb4, 0x10, 0x54, 0xa4, 0x74, 0x63,
	0x37, 0x06, 0x66, 0xa6, 0xfa, 0x04, 0x76, 0x25, 0xb8, 0x98, 0x82, 0xe0, 0x2a, 0x64, 0xa6, 0x69,
	0x7b, 0x70, 0x9a, 0x94, 0xe4, 0x74, 0xb0, 0x2f, 0xe0, 0xca, 0xbd, 0xaf, 0x2b, 0x93, 0x96, 0x44,
	0xa5, 0x9b, 0x7f, 0xce, 0x7c, 0xff, 0x97, 0x99, 0xb3, 0x08, 0xb9, 0x50, 0xba, 0x59, 0x1b, 0x2d,
	0x70, 0x65, 0x95, 0x5b, 0x99, 0x7a, 0x2e, 0x40, 0x2f, 0x0c, 0xdf, 0x58, 0x83, 0x86, 0xbe, 0x25,
	0x15, 0xb8, 0xca, 0x08, 0x30, 0x4e, 0xbc, 0x5a, 0x01, 0xba, 0x59, 0x2f, 0xad, 0x30, 0x1b, 0x65,
	0x39, 0xe8, 0x46, 0x69, 0x34, 0x76, 0xc7, 0xad, 0xac, 0x5e, 0x9c, 0x4f, 0xae, 0x34, 0x02, 0xee,
	0xb8, 0xab, 0x0d, 0x72, 0x74, 0x90, 0x3a, 0x9f, 0x6d, 0x64, 0x7e, 0xcc, 0xda, 0xc8, 0xfd, 0x98,
	0x73, 0x89, 0x68, 0xa1, 0xdc, 0xa2, 0x72, 0x3c, 0xfc, 0x9c, 0x57, 0x16, 0x10, 0x2a, 0x59, 0x8b,
	0x15, 0x1c, 0xa8, 0x28, 0xe5, 0x72, 0xf4, 0x9e, 0x90, 0xf3, 0xa3, 0x8b, 0x8a, 0x87, 0xe9, 0xf3,
	0x8c, 0x52, 0xd2, 0xd5, 0x72, 0xad, 0x58, 0x32, 0x4c, 0xc6, 0x7f, 0x0a, 0x3f, 0xd3, 0x01, 0xe9,
	0xb7, 0x4f, 0x91, 0xb2, 0x8e, 0xa7, 0xbd, 0xf6, 0x2d, 0x0d, 0x38, 0x63, 0xbf, 0x22, 0xce, 0x02,
	0xce, 0x59, 0x37, 0xe2, 0x3c, 0xe0, 0x09, 0xeb, 0x45, 0x3c, 0x19, 0x7d, 0x74, 0xc8, 0xe0, 0xe8,
	0x3a, 0xf4, 0x86, 0xd0, 0x48, 0x9c, 0x6a, 0x94, 0x05, 0xdc, 0xb1, 0x6c, 0x98, 0x8c, 0xff, 0x16,
	0x27, 0xa1, 0x99, 0x1d, 0x8a, 0xef, 0xba, 0x55, 0xb5, 0x44, 0x30, 0x9a, 0xe5, 0x3f, 0xf4, 0xe2,
	0x50, 0xd0, 0x6b, 0xf2, 0x3f, 0xea, 0x8d, 0xac, 0xb7, 0x8a, 0x4d, 0xbc, 0xfb, 0x2f, 0xe0, 0xa7,
	0x96, 0xd2, 0x94, 0x9c, 0x45, 0x51, 0xb5, 0xe6, 0xfe, 0xcb, 0xb7, 0xc3, 0x64, 0xfc, 0xbb, 0x38,
	0x0d, 0xdd, 0x34, 0x54, 0xf4, 0x9e, 0x5c, 0xc5, 0x23, 0xda, 0x20, 0x2c, 0xa0, 0xf2, 0x8d, 0x50,
	0x5a, 0x96, 0xb5, 0x9a, 0xb3, 0x3b, 0x7f, 0xf8, 0x32, 0x58, 0x8f, 0x5f, 0xa4, 0xe9, 0xde, 0x29,
	0xfb, 0xfe, 0xe2, 0xe4, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x4b, 0x7f, 0x24, 0x57, 0x02,
	0x00, 0x00,
}
