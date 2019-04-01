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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_threshold_major_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xe9, 0xbe, 0x7d, 0x43, 0x03, 0x2a, 0x46, 0x07, 0x41, 0x45, 0xc6, 0x6e, 0xdc, 0x8d,
	0x81, 0xb6, 0xea, 0x13, 0xb8, 0x2b, 0xc1, 0x8b, 0x0e, 0x04, 0xaf, 0x42, 0xda, 0x9d, 0x6d, 0xd1,
	0x36, 0x67, 0x24, 0x59, 0xd9, 0x9e, 0xc4, 0xd7, 0x95, 0x66, 0x25, 0x55, 0xd9, 0xcd, 0xe9, 0xe9,
	0xef, 0xff, 0x3b, 0x39, 0xb9, 0x08, 0xb9, 0x06, 0x5d, 0x57, 0xa8, 0x85, 0x5b, 0x19, 0xb0, 0x2b,
	0x2c, 0xe7, 0x42, 0xe9, 0x05, 0xf2, 0xb5, 0x41, 0x87, 0x74, 0x5b, 0x28, 0x5b, 0xa0, 0x50, 0x68,
	0xc5, 0xd6, 0x08, 0xa5, 0xeb, 0x6a, 0x69, 0x04, 0xae, 0xc1, 0x70, 0xa5, 0x6b, 0xd0, 0x0e, 0xcd,
	0x8e, 0x1b, 0x59, 0x7c, 0x5a, 0x5f, 0x39, 0x68, 0xa7, 0xdc, 0x8e, 0xdb, 0x12, 0x1d, 0x77, 0x56,
	0xc5, 0xd6, 0xd7, 0xa6, 0x24, 0xbe, 0x4d, 0xb8, 0x74, 0xce, 0xa8, 0x7c, 0xe3, 0xc0, 0xf2, 0xb0,
	0x90, 0x57, 0xf2, 0x03, 0x8d, 0x28, 0xb1, 0x45, 0x22, 0x97, 0xcb, 0xf1, 0x86, 0x5c, 0x1d, 0xbc,
	0x98, 0x78, 0x99, 0xbe, 0xcf, 0x28, 0x25, 0x7d, 0x2d, 0x2b, 0x60, 0xd1, 0x28, 0x9a, 0x1c, 0x67,
	0xbe, 0xa7, 0x43, 0x32, 0x68, 0xbe, 0x22, 0x66, 0x3d, 0x4f, 0xff, 0x37, 0x7f, 0x71, 0xc0, 0x09,
	0xfb, 0xd7, 0xe1, 0x24, 0xe0, 0x94, 0xf5, 0x3b, 0x9c, 0x8e, 0xbf, 0x7a, 0x64, 0x78, 0x70, 0x2f,
	0xbd, 0x27, 0xb4, 0x23, 0x16, 0x6a, 0x30, 0xca, 0xed, 0x58, 0x32, 0x8a, 0x26, 0x27, 0xd9, 0x79,
	0x48, 0x66, 0x6d, 0xf0, 0x5b, 0x37, 0x50, 0x4a, 0xa7, 0x50, 0xb3, 0xf4, 0x8f, 0x9e, 0xb5, 0x01,
	0xbd, 0x23, 0x67, 0x9d, 0x5e, 0xcb, 0x72, 0x03, 0xec, 0xc1, 0xbb, 0xa7, 0x01, 0xbf, 0x35, 0x94,
	0xc6, 0xe4, 0xb2, 0x13, 0xa1, 0x31, 0xf7, 0x27, 0x3f, 0x8e, 0xa2, 0xc9, 0x51, 0x76, 0x11, 0xb2,
	0x69, 0x88, 0xe8, 0x33, 0xb9, 0xed, 0x46, 0x34, 0x3a, 0xb5, 0x50, 0x85, 0x4f, 0x04, 0x68, 0x99,
	0x97, 0x30, 0x67, 0x4f, 0x7e, 0xf8, 0x26, 0x58, 0xaf, 0x3f, 0xa4, 0xe9, 0xde, 0xc9, 0x07, 0xfe,
	0x45, 0xa4, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xbe, 0x30, 0x18, 0x30, 0x02, 0x00, 0x00,
}
