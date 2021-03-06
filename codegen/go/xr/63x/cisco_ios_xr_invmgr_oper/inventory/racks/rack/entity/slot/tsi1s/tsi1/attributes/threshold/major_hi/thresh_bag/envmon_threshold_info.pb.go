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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_attributes_threshold_major_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0x2b, 0x31,
	0x10, 0xc7, 0xd9, 0xbe, 0xf7, 0xca, 0x33, 0xa0, 0x62, 0xb4, 0x10, 0x54, 0xa4, 0xf4, 0x62, 0x2f,
	0x06, 0xb6, 0x55, 0x3f, 0x81, 0x3d, 0x09, 0x1e, 0xb6, 0x20, 0x78, 0x31, 0x64, 0xb7, 0xd3, 0x36,
	0xba, 0x9b, 0x29, 0xc9, 0x74, 0xb1, 0x9f, 0xc4, 0xaf, 0x2b, 0x9b, 0x96, 0xac, 0x4a, 0x2f, 0x93,
	0xe4, 0xf7, 0xff, 0x4d, 0x26, 0x10, 0x76, 0x01, 0xb6, 0xae, 0xd0, 0x2a, 0x5a, 0x3a, 0xf0, 0x4b,
	0x2c, 0x67, 0xca, 0xd8, 0x39, 0xca, 0x95, 0x43, 0x42, 0x5e, 0x15, 0xc6, 0x17, 0xa8, 0x0c, 0x7a,
	0xf5, 0xe1, 0x94, 0xb1, 0x75, 0xb5, 0x70, 0x0a, 0x57, 0xe0, 0xa4, 0xb1, 0x35, 0x58, 0x42, 0xb7,
	0x91, 0x4e, 0x17, 0xef, 0x3e, 0x54, 0x09, 0x96, 0x0c, 0x6d, 0xa4, 0x2f, 0x91, 0x24, 0x79, 0x93,
	0xfa, 0x50, 0xa5, 0x26, 0x72, 0x26, 0x5f, 0x13, 0x78, 0x19, 0xa7, 0xc8, 0x4a, 0xbf, 0xa1, 0x53,
	0x4b, 0xb3, 0x43, 0x2a, 0xd7, 0x8b, 0xc1, 0x2b, 0x3b, 0xdf, 0xfb, 0x1a, 0xf5, 0x38, 0x79, 0x99,
	0x72, 0xce, 0xfe, 0x5a, 0x5d, 0x81, 0x48, 0xfa, 0xc9, 0xf0, 0x20, 0x0b, 0x7b, 0xde, 0x63, 0xdd,
	0x66, 0x55, 0xa9, 0xe8, 0x04, 0xfa, 0xaf, 0x39, 0xa5, 0x11, 0x8f, 0xc4, 0x9f, 0x16, 0x8f, 0x06,
	0x9f, 0x1d, 0xd6, 0xdb, 0x3b, 0x80, 0xdf, 0x30, 0xde, 0x12, 0x0f, 0x35, 0x38, 0x43, 0x1b, 0x31,
	0xea, 0x27, 0xc3, 0xc3, 0xec, 0x24, 0x26, 0xd3, 0x5d, 0xf0, 0x53, 0x77, 0x50, 0x6a, 0x32, 0x68,
	0xc5, 0xf8, 0x97, 0x9e, 0xed, 0x02, 0x7e, 0xcd, 0x8e, 0x5b, 0xbd, 0xd6, 0xe5, 0x1a, 0xc4, 0x6d,
	0x70, 0x8f, 0x22, 0x7e, 0x6e, 0x28, 0x4f, 0xd9, 0x59, 0x2b, 0x42, 0x63, 0x6e, 0x6f, 0xbe, 0xeb,
	0x27, 0xc3, 0xff, 0xd9, 0x69, 0xcc, 0x26, 0x31, 0xe2, 0x0f, 0xec, 0xaa, 0x6d, 0xb1, 0x48, 0x66,
	0x6e, 0x8a, 0x90, 0x28, 0xb0, 0x3a, 0x2f, 0x61, 0x26, 0xee, 0x43, 0xf3, 0x65, 0xb4, 0x9e, 0xbe,
	0x49, 0x93, 0xad, 0x93, 0x77, 0xc3, 0x7f, 0x8f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x40,
	0x16, 0x60, 0x0e, 0x02, 0x00, 0x00,
}
