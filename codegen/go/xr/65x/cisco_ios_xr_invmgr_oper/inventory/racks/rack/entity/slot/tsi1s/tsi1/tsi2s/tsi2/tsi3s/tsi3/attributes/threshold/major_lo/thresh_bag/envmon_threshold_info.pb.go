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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_entity_slot_tsi1s_tsi1_tsi2s_tsi2_tsi3s_tsi3_attributes_threshold_major_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.entity.slot.tsi1s.tsi1.tsi2s.tsi2.tsi3s.tsi3.attributes.threshold.major_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x4b, 0xf3, 0x30,
	0x18, 0xc7, 0xe9, 0xde, 0x6d, 0xbc, 0x6f, 0xe0, 0x55, 0x8c, 0x0e, 0x82, 0x8a, 0x8c, 0x5d, 0xdc,
	0xc5, 0x40, 0xdb, 0xe9, 0x27, 0x70, 0x27, 0xc1, 0x43, 0x07, 0x82, 0xa7, 0x90, 0x76, 0xcf, 0xb6,
	0x68, 0x9b, 0x67, 0x24, 0x59, 0x71, 0x67, 0xaf, 0xde, 0xfd, 0xba, 0xd2, 0x6c, 0xa4, 0x2a, 0xbb,
	0xfc, 0xfb, 0xf4, 0xf7, 0xff, 0xa5, 0x7d, 0x0e, 0x21, 0x17, 0xa0, 0xeb, 0x0a, 0xb5, 0x70, 0x2b,
	0x03, 0x76, 0x85, 0xe5, 0x5c, 0x28, 0xbd, 0x40, 0xbe, 0x36, 0xe8, 0x90, 0xbe, 0x47, 0x85, 0xb2,
	0x05, 0x0a, 0x85, 0x56, 0xbc, 0x19, 0xa1, 0x74, 0x5d, 0x2d, 0x8d, 0xc0, 0x35, 0x18, 0xae, 0x74,
	0x0d, 0xda, 0xa1, 0xd9, 0x72, 0x23, 0x8b, 0x57, 0xeb, 0x93, 0x83, 0x76, 0xca, 0x6d, 0xb9, 0x2d,
	0xd1, 0x71, 0x67, 0x55, 0x6c, 0x7d, 0x36, 0x91, 0xf8, 0x31, 0x69, 0x22, 0xf5, 0x63, 0xca, 0xa5,
	0x73, 0x46, 0xe5, 0x1b, 0x07, 0x96, 0x87, 0x9f, 0xf3, 0x4a, 0xbe, 0xa0, 0x11, 0x25, 0xee, 0x91,
	0xc8, 0xe5, 0x72, 0xf4, 0x11, 0x91, 0xf3, 0x83, 0x5b, 0x8a, 0x87, 0xe9, 0xf3, 0x8c, 0x52, 0xd2,
	0xd5, 0xb2, 0x02, 0x16, 0x0d, 0xa3, 0xf1, 0xbf, 0xcc, 0xcf, 0x74, 0x40, 0xfa, 0xcd, 0x53, 0xc4,
	0xac, 0xe3, 0x69, 0xaf, 0x79, 0x8b, 0x03, 0x4e, 0xd8, 0x9f, 0x16, 0x27, 0x01, 0xa7, 0xac, 0xdb,
	0xe2, 0x34, 0xe0, 0x09, 0xeb, 0xb5, 0x78, 0x32, 0xfa, 0xec, 0x90, 0xc1, 0xc1, 0x75, 0xe8, 0x0d,
	0xa1, 0x2d, 0xb1, 0x50, 0x83, 0x51, 0x6e, 0xcb, 0x92, 0x61, 0x34, 0xfe, 0x9f, 0x9d, 0x84, 0x66,
	0xb6, 0x2f, 0x7e, 0xea, 0x06, 0x4a, 0xe9, 0x14, 0x6a, 0x96, 0xfe, 0xd2, 0xb3, 0x7d, 0x41, 0xaf,
	0xc9, 0x71, 0xab, 0xd7, 0xb2, 0xdc, 0x00, 0x9b, 0x78, 0xf7, 0x28, 0xe0, 0xa7, 0x86, 0xd2, 0x98,
	0x9c, 0xb5, 0x22, 0x34, 0xe6, 0xee, 0xcb, 0xb7, 0xc3, 0x68, 0xfc, 0x37, 0x3b, 0x0d, 0xdd, 0x34,
	0x54, 0xf4, 0x9e, 0x5c, 0xb5, 0x47, 0x34, 0x3a, 0xb5, 0x50, 0x85, 0x6f, 0x04, 0x68, 0x99, 0x97,
	0x30, 0x67, 0x77, 0xfe, 0xf0, 0x65, 0xb0, 0x1e, 0xbf, 0x49, 0xd3, 0x9d, 0x93, 0xf7, 0xfd, 0xad,
	0x49, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0x12, 0xdb, 0x82, 0x54, 0x02, 0x00, 0x00,
}