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
// source: rsvp_mgmt_counters_oor.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_counters_out_of_resource_interfaces_interfaces_interface

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

type RsvpMgmtCountersOor_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountersOor_KEYS) Reset()         { *m = RsvpMgmtCountersOor_KEYS{} }
func (m *RsvpMgmtCountersOor_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountersOor_KEYS) ProtoMessage()    {}
func (*RsvpMgmtCountersOor_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_07bd3afb892c36d0, []int{0}
}

func (m *RsvpMgmtCountersOor_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountersOor_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtCountersOor_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountersOor_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtCountersOor_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountersOor_KEYS.Merge(m, src)
}
func (m *RsvpMgmtCountersOor_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountersOor_KEYS.Size(m)
}
func (m *RsvpMgmtCountersOor_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountersOor_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountersOor_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtCountersOor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsvpMgmtCountersOor struct {
	DroppedPathMessages  uint32   `protobuf:"varint,50,opt,name=dropped_path_messages,json=droppedPathMessages,proto3" json:"dropped_path_messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountersOor) Reset()         { *m = RsvpMgmtCountersOor{} }
func (m *RsvpMgmtCountersOor) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountersOor) ProtoMessage()    {}
func (*RsvpMgmtCountersOor) Descriptor() ([]byte, []int) {
	return fileDescriptor_07bd3afb892c36d0, []int{1}
}

func (m *RsvpMgmtCountersOor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountersOor.Unmarshal(m, b)
}
func (m *RsvpMgmtCountersOor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountersOor.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtCountersOor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountersOor.Merge(m, src)
}
func (m *RsvpMgmtCountersOor) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountersOor.Size(m)
}
func (m *RsvpMgmtCountersOor) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountersOor.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountersOor proto.InternalMessageInfo

func (m *RsvpMgmtCountersOor) GetDroppedPathMessages() uint32 {
	if m != nil {
		return m.DroppedPathMessages
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtCountersOor_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.out_of_resource.interfaces.interfaces.interface.rsvp_mgmt_counters_oor_KEYS")
	proto.RegisterType((*RsvpMgmtCountersOor)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.out_of_resource.interfaces.interfaces.interface.rsvp_mgmt_counters_oor")
}

func init() { proto.RegisterFile("rsvp_mgmt_counters_oor.proto", fileDescriptor_07bd3afb892c36d0) }

var fileDescriptor_07bd3afb892c36d0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3d, 0x6b, 0x84, 0x40,
	0x10, 0x86, 0xb1, 0x09, 0x64, 0xc1, 0x14, 0x86, 0x04, 0x21, 0x29, 0x44, 0x08, 0x58, 0x6d, 0x61,
	0xfe, 0x42, 0x52, 0xe5, 0x83, 0x60, 0xaa, 0xab, 0x86, 0x75, 0x1d, 0x75, 0x8b, 0x75, 0x96, 0x99,
	0xf5, 0xb8, 0xfb, 0xf7, 0xc7, 0x89, 0x67, 0x65, 0xf7, 0xf2, 0x3e, 0xcf, 0x0c, 0x33, 0xea, 0x95,
	0xe5, 0x18, 0xc0, 0x0f, 0x3e, 0x82, 0xa5, 0x79, 0x8a, 0xc8, 0x02, 0x44, 0xac, 0x03, 0x53, 0xa4,
	0x0c, 0xac, 0x13, 0x4b, 0xe0, 0x48, 0xe0, 0xc4, 0xe0, 0x02, 0x2c, 0x36, 0x05, 0x64, 0xbd, 0x24,
	0x89, 0x66, 0xea, 0xda, 0xb3, 0xbe, 0x8d, 0x6a, 0x9a, 0x23, 0x50, 0x0f, 0x8c, 0x42, 0x33, 0x5b,
	0xd4, 0xee, 0x5a, 0xf7, 0xc6, 0xa2, 0xec, 0xc6, 0xf2, 0x43, 0xbd, 0xec, 0x1f, 0x00, 0x5f, 0x9f,
	0x87, 0xff, 0xec, 0x4d, 0x3d, 0x6c, 0x2e, 0x4c, 0xc6, 0x63, 0x9e, 0x14, 0x49, 0x75, 0xdf, 0xa4,
	0x5b, 0xfb, 0x6b, 0x3c, 0x96, 0xdf, 0xea, 0x79, 0x7f, 0x4b, 0x56, 0xab, 0xa7, 0x8e, 0x29, 0x04,
	0xec, 0x20, 0x98, 0x38, 0x82, 0x47, 0x11, 0x33, 0xa0, 0xe4, 0x75, 0x91, 0x54, 0x69, 0xf3, 0xb8,
	0xc2, 0x3f, 0x13, 0xc7, 0x9f, 0x15, 0xb5, 0x77, 0xcb, 0xef, 0xef, 0x97, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x38, 0x49, 0x01, 0x9e, 0x1b, 0x01, 0x00, 0x00,
}
