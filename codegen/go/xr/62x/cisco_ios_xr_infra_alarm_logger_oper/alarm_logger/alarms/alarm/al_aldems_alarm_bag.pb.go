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
// source: al_aldems_alarm_bag.proto

package cisco_ios_xr_infra_alarm_logger_oper_alarm_logger_alarms_alarm

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

type AlAldemsAlarmBag_KEYS struct {
	EventId              uint32   `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlAldemsAlarmBag_KEYS) Reset()         { *m = AlAldemsAlarmBag_KEYS{} }
func (m *AlAldemsAlarmBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlAldemsAlarmBag_KEYS) ProtoMessage()    {}
func (*AlAldemsAlarmBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6190d2674cff3f, []int{0}
}

func (m *AlAldemsAlarmBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlAldemsAlarmBag_KEYS.Unmarshal(m, b)
}
func (m *AlAldemsAlarmBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlAldemsAlarmBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AlAldemsAlarmBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlAldemsAlarmBag_KEYS.Merge(m, src)
}
func (m *AlAldemsAlarmBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlAldemsAlarmBag_KEYS.Size(m)
}
func (m *AlAldemsAlarmBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlAldemsAlarmBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlAldemsAlarmBag_KEYS proto.InternalMessageInfo

func (m *AlAldemsAlarmBag_KEYS) GetEventId() uint32 {
	if m != nil {
		return m.EventId
	}
	return 0
}

type AlAldemsAlarmBag struct {
	SourceId             string   `protobuf:"bytes,50,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Timestamp            uint64   `protobuf:"varint,51,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Category             string   `protobuf:"bytes,52,opt,name=category,proto3" json:"category,omitempty"`
	Group                string   `protobuf:"bytes,53,opt,name=group,proto3" json:"group,omitempty"`
	Code                 string   `protobuf:"bytes,54,opt,name=code,proto3" json:"code,omitempty"`
	Severity             string   `protobuf:"bytes,55,opt,name=severity,proto3" json:"severity,omitempty"`
	State                string   `protobuf:"bytes,56,opt,name=state,proto3" json:"state,omitempty"`
	CorrelationId        uint32   `protobuf:"varint,57,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	IsAdmin              bool     `protobuf:"varint,58,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	AdditionalText       string   `protobuf:"bytes,59,opt,name=additional_text,json=additionalText,proto3" json:"additional_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlAldemsAlarmBag) Reset()         { *m = AlAldemsAlarmBag{} }
func (m *AlAldemsAlarmBag) String() string { return proto.CompactTextString(m) }
func (*AlAldemsAlarmBag) ProtoMessage()    {}
func (*AlAldemsAlarmBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6190d2674cff3f, []int{1}
}

func (m *AlAldemsAlarmBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlAldemsAlarmBag.Unmarshal(m, b)
}
func (m *AlAldemsAlarmBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlAldemsAlarmBag.Marshal(b, m, deterministic)
}
func (m *AlAldemsAlarmBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlAldemsAlarmBag.Merge(m, src)
}
func (m *AlAldemsAlarmBag) XXX_Size() int {
	return xxx_messageInfo_AlAldemsAlarmBag.Size(m)
}
func (m *AlAldemsAlarmBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AlAldemsAlarmBag.DiscardUnknown(m)
}

var xxx_messageInfo_AlAldemsAlarmBag proto.InternalMessageInfo

func (m *AlAldemsAlarmBag) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AlAldemsAlarmBag) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *AlAldemsAlarmBag) GetCorrelationId() uint32 {
	if m != nil {
		return m.CorrelationId
	}
	return 0
}

func (m *AlAldemsAlarmBag) GetIsAdmin() bool {
	if m != nil {
		return m.IsAdmin
	}
	return false
}

func (m *AlAldemsAlarmBag) GetAdditionalText() string {
	if m != nil {
		return m.AdditionalText
	}
	return ""
}

func init() {
	proto.RegisterType((*AlAldemsAlarmBag_KEYS)(nil), "cisco_ios_xr_infra_alarm_logger_oper.alarm_logger.alarms.alarm.al_aldems_alarm_bag_KEYS")
	proto.RegisterType((*AlAldemsAlarmBag)(nil), "cisco_ios_xr_infra_alarm_logger_oper.alarm_logger.alarms.alarm.al_aldems_alarm_bag")
}

func init() { proto.RegisterFile("al_aldems_alarm_bag.proto", fileDescriptor_3c6190d2674cff3f) }

var fileDescriptor_3c6190d2674cff3f = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x54, 0x68, 0x6a, 0xa9, 0x45, 0x32, 0x0c, 0x2e, 0x30, 0x44, 0x95, 0x10, 0x99,
	0x3a, 0x50, 0xca, 0xa7, 0x84, 0xc4, 0xc0, 0x50, 0xb1, 0x05, 0x16, 0x26, 0xeb, 0x1a, 0x1f, 0x91,
	0xa5, 0x24, 0x8e, 0xec, 0x6b, 0xd5, 0xfe, 0x2e, 0xfe, 0x20, 0xb2, 0x13, 0xd4, 0x0e, 0x9d, 0x92,
	0xe7, 0xf1, 0xdd, 0x7b, 0xfe, 0x60, 0x63, 0x28, 0x25, 0x94, 0x0a, 0x2b, 0x27, 0xa1, 0x04, 0x5b,
	0xc9, 0x25, 0x14, 0xd3, 0xc6, 0x1a, 0x32, 0xfc, 0x35, 0xd7, 0x2e, 0x37, 0x52, 0x1b, 0x27, 0x37,
	0x56, 0xea, 0xfa, 0xc7, 0x42, 0x57, 0x53, 0x9a, 0xa2, 0x40, 0x2b, 0x4d, 0x83, 0x76, 0xba, 0x6f,
	0x5a, 0x70, 0xed, 0x67, 0x32, 0x67, 0xe2, 0x40, 0xb8, 0xfc, 0x78, 0xff, 0xfe, 0xe4, 0x63, 0x16,
	0xe3, 0x1a, 0x6b, 0x92, 0x5a, 0x89, 0x28, 0x89, 0xd2, 0x61, 0xd6, 0x0f, 0xbc, 0x50, 0x93, 0xdf,
	0x23, 0x76, 0x76, 0xa0, 0x8f, 0x5f, 0xb2, 0x81, 0x33, 0x2b, 0x9b, 0xa3, 0xef, 0xb9, 0x4d, 0xa2,
	0x74, 0x90, 0xc5, 0xad, 0x58, 0x28, 0x7e, 0xc5, 0x06, 0xa4, 0x2b, 0x74, 0x04, 0x55, 0x23, 0x66,
	0x49, 0x94, 0xf6, 0xb2, 0x9d, 0xe0, 0x17, 0x2c, 0xce, 0x81, 0xb0, 0x30, 0x76, 0x2b, 0xee, 0xda,
	0xce, 0x7f, 0xe6, 0xe7, 0xec, 0xb8, 0xb0, 0x66, 0xd5, 0x88, 0x79, 0x58, 0x68, 0x81, 0x73, 0xd6,
	0xcb, 0x8d, 0x42, 0x71, 0x1f, 0x64, 0xf8, 0xf7, 0x29, 0x0e, 0xd7, 0x68, 0x35, 0x6d, 0xc5, 0x43,
	0x37, 0xbf, 0x63, 0x9f, 0xe2, 0x08, 0x08, 0xc5, 0x63, 0x9b, 0x12, 0x80, 0x5f, 0xb3, 0x51, 0x6e,
	0xac, 0xc5, 0x12, 0x48, 0x9b, 0xda, 0xef, 0xfb, 0x29, 0x9c, 0x75, 0xb8, 0x67, 0x17, 0xca, 0x5f,
	0x86, 0x76, 0x12, 0x54, 0xa5, 0x6b, 0xf1, 0x9c, 0x44, 0x69, 0x9c, 0xf5, 0xb5, 0x7b, 0xf3, 0xc8,
	0x6f, 0xd8, 0x29, 0x28, 0xa5, 0x7d, 0x21, 0x94, 0x92, 0x70, 0x43, 0xe2, 0x25, 0x4c, 0x18, 0xed,
	0xf4, 0x17, 0x6e, 0x68, 0x79, 0x12, 0xde, 0x6c, 0xf6, 0x17, 0x00, 0x00, 0xff, 0xff, 0x81, 0x74,
	0x5d, 0x1c, 0xd0, 0x01, 0x00, 0x00,
}
