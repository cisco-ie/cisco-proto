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

package cisco_ios_xr_invmgr_oper_inventory_racks_rack_attributes_threshold_minor_lo_thresh_bag

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
	proto.RegisterType((*EnvmonThresholdInfo_KEYS)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.threshold.minor_lo.thresh_bag.envmon_threshold_info_KEYS")
	proto.RegisterType((*EnvmonThresholdInfo)(nil), "cisco_ios_xr_invmgr_oper.inventory.racks.rack.attributes.threshold.minor_lo.thresh_bag.envmon_threshold_info")
}

func init() { proto.RegisterFile("envmon_threshold_info.proto", fileDescriptor_7410abf38d6220ea) }

var fileDescriptor_7410abf38d6220ea = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x99, 0x9f, 0x1f, 0xd1, 0x80, 0x8a, 0x51, 0x61, 0x50, 0x91, 0xd2, 0x8d, 0xb3, 0x31,
	0xa8, 0x55, 0x9f, 0xc0, 0x59, 0x09, 0x2e, 0xa6, 0x50, 0x70, 0x15, 0x32, 0xd3, 0xdb, 0x36, 0x38,
	0x73, 0x6f, 0xb9, 0x49, 0x07, 0xfb, 0x24, 0xbe, 0xae, 0x34, 0x1d, 0x32, 0x2a, 0xdd, 0x84, 0x70,
	0xbe, 0xef, 0x9c, 0x2c, 0x22, 0x2e, 0x01, 0xdb, 0x86, 0x50, 0xfb, 0x05, 0x83, 0x5b, 0x50, 0x3d,
	0xd5, 0x16, 0x67, 0xa4, 0x96, 0x4c, 0x9e, 0xe4, 0xa4, 0xb2, 0xae, 0x22, 0x6d, 0xc9, 0xe9, 0x4f,
	0xd6, 0x16, 0xdb, 0x66, 0xce, 0x9a, 0x96, 0xc0, 0xca, 0x62, 0x0b, 0xe8, 0x89, 0xd7, 0x8a, 0x4d,
	0xf5, 0xe1, 0xc2, 0xa9, 0x8c, 0xf7, 0x6c, 0xcb, 0x95, 0x07, 0xa7, 0xe2, 0x9c, 0x6a, 0x2c, 0x12,
	0xeb, 0x9a, 0xba, 0x48, 0x97, 0x66, 0x3e, 0xbc, 0x13, 0x17, 0x3b, 0x9f, 0xd5, 0xaf, 0xf9, 0xfb,
	0x58, 0x4a, 0xf1, 0x1f, 0x4d, 0x03, 0x69, 0x32, 0x48, 0xb2, 0x83, 0x22, 0xdc, 0x87, 0x5f, 0xff,
	0xc4, 0xf9, 0xce, 0x8a, 0xbc, 0x15, 0xb2, 0x4f, 0x1c, 0xb4, 0xc0, 0xd6, 0xaf, 0xd3, 0x87, 0x41,
	0x92, 0x1d, 0x16, 0x27, 0x91, 0x8c, 0x3b, 0xf0, 0x5b, 0x67, 0xa8, 0x8d, 0xb7, 0x84, 0xe9, 0xe8,
	0x8f, 0x5e, 0x74, 0x40, 0xde, 0x88, 0xe3, 0x5e, 0x6f, 0x4d, 0xbd, 0x82, 0xf4, 0x31, 0xb8, 0x47,
	0x31, 0x9e, 0x6c, 0x52, 0x79, 0x2f, 0xce, 0x7a, 0x11, 0x36, 0xe6, 0x76, 0xf9, 0x69, 0x90, 0x64,
	0xfb, 0xc5, 0x69, 0x64, 0x79, 0x44, 0xf2, 0x45, 0x5c, 0xf7, 0x15, 0x24, 0x6f, 0x67, 0xb6, 0x0a,
	0x44, 0x03, 0x9a, 0xb2, 0x86, 0x69, 0xfa, 0x1c, 0xca, 0x57, 0xd1, 0x7a, 0xfb, 0x21, 0xe5, 0x5b,
	0xa7, 0xdc, 0x0b, 0x5f, 0x35, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x9e, 0x08, 0x41, 0xc9,
	0x01, 0x00, 0x00,
}
