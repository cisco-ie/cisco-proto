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
// source: rcmd_show_intf_event.proto

package cisco_ios_xr_infra_rcmd_oper_rcmd_intf_events_event

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

type RcmdShowIntfEvent_KEYS struct {
	EventNo              uint32   `protobuf:"varint,1,opt,name=event_no,json=eventNo,proto3" json:"event_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowIntfEvent_KEYS) Reset()         { *m = RcmdShowIntfEvent_KEYS{} }
func (m *RcmdShowIntfEvent_KEYS) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIntfEvent_KEYS) ProtoMessage()    {}
func (*RcmdShowIntfEvent_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b7ad2ba81813b9, []int{0}
}

func (m *RcmdShowIntfEvent_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIntfEvent_KEYS.Unmarshal(m, b)
}
func (m *RcmdShowIntfEvent_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIntfEvent_KEYS.Marshal(b, m, deterministic)
}
func (m *RcmdShowIntfEvent_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIntfEvent_KEYS.Merge(m, src)
}
func (m *RcmdShowIntfEvent_KEYS) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIntfEvent_KEYS.Size(m)
}
func (m *RcmdShowIntfEvent_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIntfEvent_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIntfEvent_KEYS proto.InternalMessageInfo

func (m *RcmdShowIntfEvent_KEYS) GetEventNo() uint32 {
	if m != nil {
		return m.EventNo
	}
	return 0
}

type RcmdShowIntfEvent struct {
	SequenceNo           uint32   `protobuf:"varint,50,opt,name=sequence_no,json=sequenceNo,proto3" json:"sequence_no,omitempty"`
	InterfaceName        string   `protobuf:"bytes,51,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Component            string   `protobuf:"bytes,52,opt,name=component,proto3" json:"component,omitempty"`
	EventType            string   `protobuf:"bytes,53,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EventTime            string   `protobuf:"bytes,54,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	PrimaryAddress       string   `protobuf:"bytes,55,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowIntfEvent) Reset()         { *m = RcmdShowIntfEvent{} }
func (m *RcmdShowIntfEvent) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIntfEvent) ProtoMessage()    {}
func (*RcmdShowIntfEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_70b7ad2ba81813b9, []int{1}
}

func (m *RcmdShowIntfEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIntfEvent.Unmarshal(m, b)
}
func (m *RcmdShowIntfEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIntfEvent.Marshal(b, m, deterministic)
}
func (m *RcmdShowIntfEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIntfEvent.Merge(m, src)
}
func (m *RcmdShowIntfEvent) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIntfEvent.Size(m)
}
func (m *RcmdShowIntfEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIntfEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIntfEvent proto.InternalMessageInfo

func (m *RcmdShowIntfEvent) GetSequenceNo() uint32 {
	if m != nil {
		return m.SequenceNo
	}
	return 0
}

func (m *RcmdShowIntfEvent) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RcmdShowIntfEvent) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *RcmdShowIntfEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *RcmdShowIntfEvent) GetEventTime() string {
	if m != nil {
		return m.EventTime
	}
	return ""
}

func (m *RcmdShowIntfEvent) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*RcmdShowIntfEvent_KEYS)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.intf.events.event.rcmd_show_intf_event_KEYS")
	proto.RegisterType((*RcmdShowIntfEvent)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.intf.events.event.rcmd_show_intf_event")
}

func init() { proto.RegisterFile("rcmd_show_intf_event.proto", fileDescriptor_70b7ad2ba81813b9) }

var fileDescriptor_70b7ad2ba81813b9 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbb, 0x4e, 0x2b, 0x31,
	0x10, 0x86, 0xb5, 0xcd, 0x39, 0x64, 0x50, 0x82, 0x64, 0x51, 0x38, 0x08, 0x44, 0x14, 0x09, 0x91,
	0x6a, 0x0b, 0x16, 0x42, 0x4d, 0x41, 0x85, 0xb4, 0x45, 0xa0, 0xa1, 0xb2, 0x8c, 0x77, 0x56, 0xb8,
	0xb0, 0xc7, 0x8c, 0xcd, 0x65, 0x9f, 0x97, 0x17, 0x41, 0x6b, 0x73, 0x49, 0x91, 0xca, 0x9e, 0xef,
	0xff, 0x3f, 0xcb, 0x1a, 0x38, 0x62, 0xe3, 0x3a, 0x15, 0x9f, 0xe9, 0x5d, 0x59, 0x9f, 0x7a, 0x85,
	0x6f, 0xe8, 0x53, 0x1d, 0x98, 0x12, 0x89, 0xc6, 0xd8, 0x68, 0x48, 0x59, 0x8a, 0xea, 0x83, 0x95,
	0xf5, 0x3d, 0x6b, 0x95, 0xeb, 0x14, 0x90, 0xeb, 0xf1, 0x56, 0x8f, 0x4e, 0x9d, 0x9d, 0x58, 0x8e,
	0xe5, 0x1a, 0xe6, 0xbb, 0x9e, 0x54, 0x77, 0xb7, 0x8f, 0xf7, 0x62, 0x0e, 0x7b, 0x65, 0xf2, 0x24,
	0xab, 0x45, 0xb5, 0x9a, 0x6e, 0xfe, 0xe7, 0xb9, 0xa5, 0xe5, 0x67, 0x05, 0x87, 0xbb, 0x44, 0x71,
	0x0a, 0xfb, 0x11, 0x5f, 0x5e, 0xd1, 0x1b, 0x1c, 0xb5, 0x8b, 0xac, 0xc1, 0x0f, 0x6a, 0x49, 0x9c,
	0xc1, 0xcc, 0xfa, 0x84, 0xdc, 0xeb, 0xb1, 0xa1, 0x1d, 0xca, 0x66, 0x51, 0xad, 0x26, 0x9b, 0xe9,
	0x2f, 0x6d, 0xb5, 0x43, 0x71, 0x0c, 0x13, 0x43, 0x2e, 0x90, 0x47, 0x9f, 0xe4, 0x65, 0x6e, 0xfc,
	0x01, 0x71, 0x02, 0x50, 0x7e, 0x96, 0x86, 0x80, 0xf2, 0xaa, 0xc4, 0x99, 0x3c, 0x0c, 0x01, 0xb7,
	0x62, 0xeb, 0x50, 0xae, 0xb7, 0x63, 0xeb, 0x50, 0x9c, 0xc3, 0x41, 0x60, 0xeb, 0x34, 0x0f, 0x4a,
	0x77, 0x1d, 0x63, 0x8c, 0xf2, 0x3a, 0x77, 0x66, 0xdf, 0xf8, 0xa6, 0xd0, 0xa7, 0x7f, 0x79, 0xb3,
	0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xff, 0xaa, 0xd6, 0x77, 0x01, 0x00, 0x00,
}
