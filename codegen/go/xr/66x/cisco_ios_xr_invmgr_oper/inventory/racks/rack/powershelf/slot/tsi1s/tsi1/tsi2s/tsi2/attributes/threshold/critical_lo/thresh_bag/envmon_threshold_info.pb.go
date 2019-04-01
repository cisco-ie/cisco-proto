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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_powershelf_slot_tsi1s_tsi1_tsi2s_tsi2_attributes_threshold_critical_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.critical_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.powershelf.slot.tsi1s.tsi1.tsi2s.tsi2.attributes.threshold.critical_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x4a, 0x33, 0x31,
	0x10, 0xc7, 0xd9, 0x7e, 0xfd, 0x8a, 0x06, 0x54, 0x8c, 0x16, 0x82, 0x8a, 0x94, 0x5e, 0xec, 0xc5,
	0xc0, 0xee, 0xaa, 0x4f, 0x60, 0x4f, 0x82, 0x87, 0x2d, 0x08, 0x9e, 0x42, 0x76, 0x3b, 0x6d, 0x83,
	0x69, 0xa6, 0x24, 0xe9, 0x6a, 0x4f, 0x3e, 0x86, 0xaf, 0x2b, 0x9b, 0x96, 0xac, 0x4a, 0x2f, 0xb3,
	0xb3, 0xbf, 0xff, 0x6f, 0x32, 0x39, 0x84, 0x5c, 0x82, 0xa9, 0x97, 0x68, 0x84, 0x5f, 0x58, 0x70,
	0x0b, 0xd4, 0x53, 0xa1, 0xcc, 0x0c, 0xf9, 0xca, 0xa2, 0x47, 0xfa, 0x59, 0x29, 0x57, 0xa1, 0x50,
	0xe8, 0xc4, 0x87, 0x15, 0xca, 0xd4, 0xcb, 0xb9, 0x15, 0xb8, 0x02, 0xcb, 0x95, 0xa9, 0xc1, 0x78,
	0xb4, 0x1b, 0x6e, 0x65, 0xf5, 0xe6, 0x42, 0xe5, 0x2b, 0x7c, 0x07, 0xeb, 0x16, 0xa0, 0x67, 0xdc,
	0x69, 0xf4, 0xdc, 0x3b, 0x95, 0xba, 0x50, 0x9b, 0x92, 0x85, 0x36, 0xe3, 0xd2, 0x7b, 0xab, 0xca,
	0xb5, 0x07, 0xc7, 0xe3, 0x52, 0x5e, 0x59, 0xe5, 0x55, 0x25, 0xb5, 0xd0, 0xb8, 0xa3, 0xa2, 0x94,
	0xf3, 0xe1, 0x9a, 0x5c, 0xec, 0xbd, 0x9f, 0x78, 0x1a, 0xbf, 0x4e, 0x28, 0x25, 0x5d, 0x23, 0x97,
	0xc0, 0x92, 0x41, 0x32, 0x3a, 0x2c, 0x42, 0x4f, 0xfb, 0xa4, 0xd7, 0x7c, 0x45, 0xca, 0x3a, 0x81,
	0xfe, 0x6f, 0xfe, 0xd2, 0x88, 0x33, 0xf6, 0xaf, 0xc5, 0x59, 0xc4, 0x39, 0xeb, 0xb6, 0x38, 0x1f,
	0x7e, 0x75, 0x48, 0x7f, 0xef, 0x5e, 0x7a, 0x4b, 0x68, 0x4b, 0x1c, 0xd4, 0x60, 0x95, 0xdf, 0xb0,
	0x6c, 0x90, 0x8c, 0x8e, 0x8a, 0xd3, 0x98, 0x4c, 0x76, 0xc1, 0x6f, 0xdd, 0x82, 0x96, 0x5e, 0xa1,
	0x61, 0xf9, 0x1f, 0xbd, 0xd8, 0x05, 0xf4, 0x86, 0x9c, 0xb4, 0x7a, 0x2d, 0xf5, 0x1a, 0xd8, 0x5d,
	0x70, 0x8f, 0x23, 0x7e, 0x69, 0x28, 0x4d, 0xc9, 0x79, 0x2b, 0x42, 0x63, 0x6e, 0x4f, 0xbe, 0x1f,
	0x24, 0xa3, 0x83, 0xe2, 0x2c, 0x66, 0xe3, 0x18, 0xd1, 0x47, 0x72, 0xdd, 0x8e, 0x18, 0xf4, 0x6a,
	0xa6, 0xaa, 0x90, 0x08, 0x30, 0xb2, 0xd4, 0x30, 0x65, 0x0f, 0x61, 0xf8, 0x2a, 0x5a, 0xcf, 0x3f,
	0xa4, 0xf1, 0xd6, 0x29, 0x7b, 0xe1, 0x61, 0xe4, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xea, 0x83,
	0x9f, 0x43, 0x37, 0x02, 0x00, 0x00,
}
