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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_fantray_slot_attributes_threshold_critical_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.attributes.threshold.critical_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.fantray.slot.attributes.threshold.critical_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x99, 0xf2, 0x7d, 0x45, 0x03, 0x2a, 0x46, 0x0b, 0x83, 0x8a, 0x94, 0x6e, 0xec, 0xc6,
	0x40, 0xad, 0x7a, 0x05, 0x16, 0x17, 0x82, 0x8b, 0x29, 0x08, 0xae, 0x42, 0x66, 0x7a, 0xa6, 0x0d,
	0xa6, 0x39, 0xe5, 0xe4, 0x74, 0xb0, 0x57, 0xe2, 0xed, 0xca, 0xa4, 0x65, 0x46, 0xa5, 0x9b, 0xfc,
	0xbc, 0xcf, 0x73, 0x5e, 0x02, 0x11, 0x97, 0xe0, 0xab, 0x25, 0x7a, 0xcd, 0x0b, 0x82, 0xb0, 0x40,
	0x37, 0xd3, 0xd6, 0x97, 0xa8, 0x56, 0x84, 0x8c, 0xb2, 0x2c, 0x6c, 0x28, 0x50, 0x5b, 0x0c, 0xfa,
	0x93, 0xb4, 0xf5, 0xd5, 0x72, 0x4e, 0x1a, 0x57, 0x40, 0xca, 0xfa, 0x0a, 0x3c, 0x23, 0x6d, 0x14,
	0x99, 0xe2, 0x23, 0xc4, 0x55, 0x95, 0xc6, 0x33, 0x99, 0x8d, 0x0a, 0x0e, 0x59, 0x19, 0x66, 0xb2,
	0xf9, 0x9a, 0x21, 0xa8, 0xa6, 0x5b, 0x15, 0x64, 0xd9, 0x16, 0xc6, 0x69, 0x87, 0xbb, 0x54, 0xe7,
	0x66, 0x3e, 0x78, 0x16, 0x17, 0x7b, 0x9f, 0xa1, 0x5f, 0x26, 0xef, 0x53, 0x29, 0xc5, 0x3f, 0x6f,
	0x96, 0x90, 0x26, 0xfd, 0x64, 0x78, 0x98, 0xc5, 0xb3, 0xec, 0x89, 0x6e, 0xbd, 0xeb, 0x51, 0xda,
	0x89, 0xe9, 0xff, 0xfa, 0x36, 0x1a, 0x7c, 0x75, 0x44, 0x6f, 0x6f, 0x93, 0xbc, 0x15, 0xb2, 0x4d,
	0x02, 0x54, 0x40, 0x96, 0x37, 0xe9, 0x5d, 0x3f, 0x19, 0x1e, 0x65, 0xa7, 0x0d, 0x99, 0xee, 0xc0,
	0x6f, 0x9d, 0xc0, 0x19, 0xb6, 0xe8, 0xd3, 0xf1, 0x1f, 0x3d, 0xdb, 0x01, 0x79, 0x23, 0x4e, 0x5a,
	0xbd, 0x32, 0x6e, 0x0d, 0xe9, 0x7d, 0x74, 0x8f, 0x9b, 0xf8, 0xad, 0x4e, 0xe5, 0x48, 0x9c, 0xb7,
	0x22, 0xd4, 0xe6, 0xb6, 0xf9, 0xa1, 0x9f, 0x0c, 0x0f, 0xb2, 0xb3, 0x86, 0x4d, 0x1a, 0x24, 0x9f,
	0xc4, 0x75, 0x3b, 0xe2, 0x91, 0x6d, 0x69, 0x8b, 0x48, 0x34, 0x78, 0x93, 0x3b, 0x98, 0xa5, 0x8f,
	0x71, 0xf8, 0xaa, 0xb1, 0x5e, 0x7f, 0x48, 0x93, 0xad, 0x93, 0x77, 0xe3, 0x8f, 0x8e, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x71, 0x75, 0xda, 0xf8, 0xf0, 0x01, 0x00, 0x00,
}