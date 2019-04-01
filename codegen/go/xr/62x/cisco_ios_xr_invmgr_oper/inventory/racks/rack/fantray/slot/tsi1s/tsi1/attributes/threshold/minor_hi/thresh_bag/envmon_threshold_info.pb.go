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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_tsi1s_tsi1_attributes_threshold_minor_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.tsi1s.tsi1.attributes.threshold.minor_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0xd9, 0x7e, 0x9f, 0x45, 0x03, 0x2a, 0x46, 0x0b, 0x41, 0x45, 0x4a, 0x2f, 0xf6, 0x62,
	0x60, 0x5b, 0xf5, 0x09, 0xec, 0x49, 0xf0, 0xb0, 0x05, 0xc1, 0x8b, 0x21, 0xbb, 0x4d, 0xdb, 0xc1,
	0xdd, 0x99, 0x92, 0xa4, 0x8b, 0x7d, 0x12, 0x5f, 0x57, 0x36, 0x2d, 0x59, 0x95, 0x5e, 0x26, 0xc9,
	0xef, 0xff, 0x9b, 0x4c, 0x20, 0xec, 0xca, 0x60, 0x5d, 0x11, 0x2a, 0xbf, 0xb4, 0xc6, 0x2d, 0xa9,
	0x9c, 0x29, 0xc0, 0x39, 0xc9, 0x95, 0x25, 0x4f, 0x1c, 0x0b, 0x70, 0x05, 0x29, 0x20, 0xa7, 0x3e,
	0xad, 0x02, 0xac, 0xab, 0x85, 0x55, 0xb4, 0x32, 0x56, 0x02, 0xd6, 0x06, 0x3d, 0xd9, 0x8d, 0xb4,
	0xba, 0xf8, 0x70, 0xa1, 0xca, 0xb9, 0x46, 0x6f, 0xf5, 0x46, 0xba, 0x92, 0xbc, 0xf4, 0x0e, 0x52,
	0x17, 0xaa, 0xd4, 0xde, 0x5b, 0xc8, 0xd7, 0xde, 0x38, 0x19, 0xc7, 0xc8, 0x0a, 0x90, 0xac, 0x5a,
	0xc2, 0x0e, 0xa9, 0x5c, 0x2f, 0x06, 0xef, 0xec, 0x72, 0xef, 0x73, 0xd4, 0xf3, 0xe4, 0x6d, 0xca,
	0x39, 0xfb, 0x8f, 0xba, 0x32, 0x22, 0xe9, 0x27, 0xc3, 0xa3, 0x2c, 0xec, 0x79, 0x8f, 0x75, 0x9b,
	0x55, 0xa5, 0xa2, 0x13, 0xe8, 0x41, 0x73, 0x4a, 0x23, 0x1e, 0x89, 0x7f, 0x2d, 0x1e, 0x0d, 0xbe,
	0x3a, 0xac, 0xb7, 0x77, 0x00, 0xbf, 0x63, 0xbc, 0x25, 0xce, 0xd4, 0xc6, 0x82, 0xdf, 0x88, 0x51,
	0x3f, 0x19, 0x1e, 0x67, 0x67, 0x31, 0x99, 0xee, 0x82, 0xdf, 0xba, 0x35, 0xa5, 0xf6, 0x40, 0x28,
	0xc6, 0x7f, 0xf4, 0x6c, 0x17, 0xf0, 0x5b, 0x76, 0xda, 0xea, 0xb5, 0x2e, 0xd7, 0x46, 0xdc, 0x07,
	0xf7, 0x24, 0xe2, 0xd7, 0x86, 0xf2, 0x94, 0x5d, 0xb4, 0xa2, 0x69, 0xcc, 0xed, 0xcd, 0x0f, 0xfd,
	0x64, 0x78, 0x98, 0x9d, 0xc7, 0x6c, 0x12, 0x23, 0xfe, 0xc4, 0x6e, 0xda, 0x16, 0x24, 0x0f, 0x73,
	0x28, 0x42, 0xa2, 0x0c, 0xea, 0xbc, 0x34, 0x33, 0xf1, 0x18, 0x9a, 0xaf, 0xa3, 0xf5, 0xf2, 0x43,
	0x9a, 0x6c, 0x9d, 0xbc, 0x1b, 0x3e, 0x7c, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x22, 0xcb,
	0x0c, 0x0f, 0x02, 0x00, 0x00,
}