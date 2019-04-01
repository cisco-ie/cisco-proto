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
// source: rsvp_mgmt_event_count_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_counters_interface_events_interface_event

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

type RsvpMgmtEventCountInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEventCountInfo_KEYS) Reset()         { *m = RsvpMgmtEventCountInfo_KEYS{} }
func (m *RsvpMgmtEventCountInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEventCountInfo_KEYS) ProtoMessage()    {}
func (*RsvpMgmtEventCountInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_68d9c8f8dcc21863, []int{0}
}

func (m *RsvpMgmtEventCountInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtEventCountInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtEventCountInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS.Merge(m, src)
}
func (m *RsvpMgmtEventCountInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS.Size(m)
}
func (m *RsvpMgmtEventCountInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEventCountInfo_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtEventCountInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsvpMgmtEventCountInfo struct {
	ExpiredPaths         uint32   `protobuf:"varint,50,opt,name=expired_paths,json=expiredPaths,proto3" json:"expired_paths,omitempty"`
	ExpiredReservations  uint32   `protobuf:"varint,51,opt,name=expired_reservations,json=expiredReservations,proto3" json:"expired_reservations,omitempty"`
	NacKs                uint32   `protobuf:"varint,52,opt,name=nac_ks,json=nacKs,proto3" json:"nac_ks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtEventCountInfo) Reset()         { *m = RsvpMgmtEventCountInfo{} }
func (m *RsvpMgmtEventCountInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtEventCountInfo) ProtoMessage()    {}
func (*RsvpMgmtEventCountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68d9c8f8dcc21863, []int{1}
}

func (m *RsvpMgmtEventCountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtEventCountInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtEventCountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtEventCountInfo.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtEventCountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtEventCountInfo.Merge(m, src)
}
func (m *RsvpMgmtEventCountInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtEventCountInfo.Size(m)
}
func (m *RsvpMgmtEventCountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtEventCountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtEventCountInfo proto.InternalMessageInfo

func (m *RsvpMgmtEventCountInfo) GetExpiredPaths() uint32 {
	if m != nil {
		return m.ExpiredPaths
	}
	return 0
}

func (m *RsvpMgmtEventCountInfo) GetExpiredReservations() uint32 {
	if m != nil {
		return m.ExpiredReservations
	}
	return 0
}

func (m *RsvpMgmtEventCountInfo) GetNacKs() uint32 {
	if m != nil {
		return m.NacKs
	}
	return 0
}

func init() {
	proto.RegisterType((*RsvpMgmtEventCountInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.interface_events.interface_event.rsvp_mgmt_event_count_info_KEYS")
	proto.RegisterType((*RsvpMgmtEventCountInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.counters.interface_events.interface_event.rsvp_mgmt_event_count_info")
}

func init() { proto.RegisterFile("rsvp_mgmt_event_count_info.proto", fileDescriptor_68d9c8f8dcc21863) }

var fileDescriptor_68d9c8f8dcc21863 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0xd9, 0x83, 0x05, 0x83, 0xeb, 0x21, 0x2a, 0x04, 0x2f, 0x2e, 0x15, 0xa1, 0xa7, 0x80,
	0xd6, 0x57, 0x10, 0x0a, 0x05, 0x91, 0xf5, 0xe4, 0x69, 0x88, 0x71, 0xaa, 0x41, 0x92, 0x09, 0x33,
	0x71, 0xe9, 0x3b, 0xf8, 0xd2, 0xb2, 0xb1, 0x7e, 0x40, 0xd9, 0xd3, 0x0c, 0xbf, 0xff, 0xc7, 0xc0,
	0xa8, 0x8e, 0x65, 0xc8, 0x10, 0x5f, 0x63, 0x01, 0x1c, 0x30, 0x15, 0xf0, 0xf4, 0x91, 0x0a, 0x84,
	0xb4, 0x21, 0x9b, 0x99, 0x0a, 0xe9, 0x95, 0x0f, 0xe2, 0x09, 0x02, 0x09, 0x6c, 0x19, 0x42, 0x86,
	0x9a, 0xa0, 0x8c, 0x6c, 0xc7, 0xcd, 0xd6, 0x00, 0xb2, 0xd8, 0x30, 0x8e, 0x8d, 0xf3, 0xf8, 0xdd,
	0xb4, 0x07, 0xe6, 0x2b, 0x75, 0x31, 0x7d, 0x0d, 0xd6, 0x77, 0x4f, 0x8f, 0xfa, 0x4a, 0x1d, 0xff,
	0xa5, 0x92, 0x8b, 0x68, 0x9a, 0xae, 0x59, 0x1c, 0xf6, 0xed, 0x2f, 0xbd, 0x77, 0x11, 0xe7, 0x9f,
	0x8d, 0x3a, 0x9f, 0xae, 0xd2, 0x97, 0xaa, 0xc5, 0x6d, 0x0e, 0x8c, 0x2f, 0x90, 0x5d, 0x79, 0x13,
	0x73, 0xd3, 0x35, 0x8b, 0xb6, 0x3f, 0xda, 0xc1, 0x87, 0x91, 0xe9, 0x6b, 0x75, 0xfa, 0x63, 0x62,
	0x14, 0xe4, 0xc1, 0x95, 0x40, 0x49, 0xcc, 0xb2, 0x7a, 0x4f, 0x76, 0x5a, 0xff, 0x4f, 0xd2, 0x67,
	0x6a, 0x96, 0x9c, 0x87, 0x77, 0x31, 0xb7, 0xd5, 0x74, 0x90, 0x9c, 0x5f, 0xcb, 0xf3, 0xac, 0x3e,
	0x6a, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x5a, 0x81, 0xa0, 0x4c, 0x01, 0x00, 0x00,
}