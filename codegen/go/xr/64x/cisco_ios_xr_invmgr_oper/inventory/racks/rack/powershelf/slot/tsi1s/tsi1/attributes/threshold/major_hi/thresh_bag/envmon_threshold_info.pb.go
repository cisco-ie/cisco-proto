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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_attributes_threshold_major_hi_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.attributes.threshold.major_hi.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xd4, 0xa1, 0x01, 0x15, 0xa3, 0x83, 0xa0, 0x22, 0x63, 0x17, 0x77, 0x31, 0xd0,
	0x4d, 0xfd, 0x04, 0xee, 0x24, 0x78, 0xe8, 0x40, 0xf0, 0x62, 0x48, 0xbb, 0xb7, 0x35, 0xda, 0xe6,
	0xcd, 0x24, 0xab, 0xee, 0x93, 0xf8, 0x75, 0xa5, 0x59, 0x49, 0x55, 0x76, 0x79, 0x49, 0x7e, 0xff,
	0xdf, 0xcb, 0x0b, 0x84, 0x5c, 0x80, 0xae, 0x4a, 0xd4, 0xc2, 0xe5, 0x06, 0x6c, 0x8e, 0xc5, 0x4c,
	0x28, 0x3d, 0x47, 0xbe, 0x34, 0xe8, 0x90, 0x7e, 0x64, 0xca, 0x66, 0x28, 0x14, 0x5a, 0xf1, 0x65,
	0x84, 0xd2, 0x55, 0xb9, 0x30, 0x02, 0x97, 0x60, 0xb8, 0xd2, 0x15, 0x68, 0x87, 0x66, 0xcd, 0x8d,
	0xcc, 0xde, 0xad, 0xaf, 0x7c, 0x89, 0x9f, 0x60, 0x6c, 0x0e, 0xc5, 0x9c, 0xdb, 0x02, 0x1d, 0x77,
	0x56, 0xc5, 0xd6, 0x57, 0x2e, 0x9d, 0x33, 0x2a, 0x5d, 0x39, 0xb0, 0x3c, 0x4c, 0xe2, 0xa5, 0x7c,
	0x43, 0x23, 0x72, 0xd5, 0x20, 0x91, 0xca, 0xc5, 0xe0, 0x95, 0x9c, 0x6f, 0x7d, 0x91, 0x78, 0x9c,
	0xbc, 0x4c, 0x29, 0x25, 0xbb, 0x5a, 0x96, 0xc0, 0xa2, 0x7e, 0x34, 0x3c, 0x48, 0xfc, 0x9e, 0xf6,
	0x48, 0xb7, 0x5e, 0x45, 0xcc, 0x3a, 0x9e, 0xee, 0xd5, 0xa7, 0x38, 0xe0, 0x11, 0xdb, 0x69, 0xf1,
	0x68, 0xf0, 0xdd, 0x21, 0xbd, 0xad, 0x03, 0xe8, 0x0d, 0xa1, 0x2d, 0xb1, 0x50, 0x81, 0x51, 0x6e,
	0xcd, 0x46, 0xfd, 0x68, 0x78, 0x98, 0x9c, 0x84, 0x64, 0xda, 0x04, 0x7f, 0x75, 0x03, 0x85, 0x74,
	0x0a, 0x35, 0x1b, 0xff, 0xd3, 0x93, 0x26, 0xa0, 0xd7, 0xe4, 0xb8, 0xd5, 0x2b, 0x59, 0xac, 0x80,
	0xdd, 0x7a, 0xf7, 0x28, 0xe0, 0xe7, 0x9a, 0xd2, 0x98, 0x9c, 0xb5, 0x22, 0xd4, 0xe6, 0xe6, 0xe6,
	0xbb, 0x7e, 0x34, 0xdc, 0x4f, 0x4e, 0x43, 0x36, 0x09, 0x11, 0x7d, 0x20, 0x57, 0x6d, 0x8b, 0x46,
	0xa7, 0xe6, 0x2a, 0xf3, 0x89, 0x00, 0x2d, 0xd3, 0x02, 0x66, 0xec, 0xde, 0x37, 0x5f, 0x06, 0xeb,
	0xe9, 0x97, 0x34, 0xd9, 0x38, 0x69, 0xd7, 0xff, 0xf9, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x86, 0x5c, 0xee, 0x12, 0x02, 0x00, 0x00,
}
