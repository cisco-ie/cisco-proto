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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_attributes_threshold_minor_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.attributes.threshold.minor_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.attributes.threshold.minor_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xa2, 0x45, 0x03, 0x2a, 0x46, 0x0b, 0x83, 0x8a, 0x94, 0x6e, 0xec, 0xc6, 0x40,
	0xad, 0xfa, 0x04, 0x16, 0x17, 0x82, 0x8b, 0x29, 0x08, 0xae, 0x42, 0x66, 0x7a, 0xa7, 0x0d, 0x66,
	0x72, 0xcb, 0x4d, 0x3a, 0xda, 0x27, 0xf1, 0x75, 0x65, 0xd2, 0x92, 0x51, 0xe9, 0x26, 0x3f, 0xe7,
	0xfb, 0xee, 0x21, 0x10, 0x76, 0x09, 0xb6, 0xae, 0xd0, 0x4a, 0xbf, 0x20, 0x70, 0x0b, 0x34, 0x33,
	0xa9, 0x6d, 0x89, 0x62, 0x49, 0xe8, 0x91, 0x97, 0x85, 0x76, 0x05, 0x4a, 0x8d, 0x4e, 0x7e, 0x91,
	0xd4, 0xb6, 0xae, 0xe6, 0x24, 0x71, 0x09, 0x24, 0xb4, 0xad, 0xc1, 0x7a, 0xa4, 0xb5, 0x20, 0x55,
	0x7c, 0xb8, 0xb0, 0x8a, 0x25, 0x7e, 0x02, 0xb9, 0x05, 0x98, 0x52, 0x38, 0x83, 0x5e, 0x28, 0xef,
	0x49, 0xe7, 0x2b, 0x0f, 0x4e, 0xc4, 0x7a, 0x51, 0x69, 0x8b, 0x24, 0x0d, 0x6e, 0x23, 0x99, 0xab,
	0xf9, 0xe0, 0x99, 0x5d, 0xec, 0x7c, 0x86, 0x7c, 0x99, 0xbc, 0x4f, 0x39, 0x67, 0x7b, 0x56, 0x55,
	0x90, 0x26, 0xfd, 0x64, 0x78, 0x98, 0x85, 0x33, 0xef, 0xb1, 0x6e, 0xb3, 0xcb, 0x51, 0xda, 0x09,
	0xe9, 0x7e, 0x73, 0x1b, 0x0d, 0xbe, 0x3b, 0xac, 0xb7, 0xb3, 0x89, 0xdf, 0x32, 0xde, 0x26, 0x0e,
	0x6a, 0x20, 0xed, 0xd7, 0xe9, 0x5d, 0x3f, 0x19, 0x1e, 0x65, 0xa7, 0x91, 0x4c, 0xb7, 0xe0, 0xaf,
	0x4e, 0x60, 0x94, 0xd7, 0x68, 0xd3, 0xf1, 0x3f, 0x3d, 0xdb, 0x02, 0x7e, 0xc3, 0x4e, 0x5a, 0xbd,
	0x56, 0x66, 0x05, 0xe9, 0x7d, 0x70, 0x8f, 0x63, 0xfc, 0xd6, 0xa4, 0x7c, 0xc4, 0xce, 0x5b, 0x11,
	0x1a, 0x73, 0xd3, 0xfc, 0xd0, 0x4f, 0x86, 0x07, 0xd9, 0x59, 0x64, 0x93, 0x88, 0xf8, 0x13, 0xbb,
	0x6e, 0x47, 0x2c, 0x7a, 0x5d, 0xea, 0x22, 0x10, 0x09, 0x56, 0xe5, 0x06, 0x66, 0xe9, 0x63, 0x18,
	0xbe, 0x8a, 0xd6, 0xeb, 0x2f, 0x69, 0xb2, 0x71, 0xf2, 0x6e, 0xf8, 0xd1, 0xf1, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x82, 0x1a, 0x78, 0xbc, 0xf0, 0x01, 0x00, 0x00,
}
