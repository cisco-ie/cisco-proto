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
// source: rsvp_mgmt_frr_states_compact.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_frr_summary

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

type RsvpMgmtFrrStatesCompact_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtFrrStatesCompact_KEYS) Reset()         { *m = RsvpMgmtFrrStatesCompact_KEYS{} }
func (m *RsvpMgmtFrrStatesCompact_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesCompact_KEYS) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesCompact_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e569263222c28bc, []int{0}
}

func (m *RsvpMgmtFrrStatesCompact_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesCompact_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtFrrStatesCompact_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS.Merge(m, src)
}
func (m *RsvpMgmtFrrStatesCompact_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS.Size(m)
}
func (m *RsvpMgmtFrrStatesCompact_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesCompact_KEYS proto.InternalMessageInfo

type RsvpMgmtFrrStateStatusCounters struct {
	Total                uint32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	ActiveInstances      uint32   `protobuf:"varint,2,opt,name=active_instances,json=activeInstances,proto3" json:"active_instances,omitempty"`
	ReadyInstances       uint32   `protobuf:"varint,3,opt,name=ready_instances,json=readyInstances,proto3" json:"ready_instances,omitempty"`
	ActiveWaitInstances  uint32   `protobuf:"varint,4,opt,name=active_wait_instances,json=activeWaitInstances,proto3" json:"active_wait_instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtFrrStateStatusCounters) Reset()         { *m = RsvpMgmtFrrStateStatusCounters{} }
func (m *RsvpMgmtFrrStateStatusCounters) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStateStatusCounters) ProtoMessage()    {}
func (*RsvpMgmtFrrStateStatusCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e569263222c28bc, []int{1}
}

func (m *RsvpMgmtFrrStateStatusCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStateStatusCounters.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStateStatusCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStateStatusCounters.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtFrrStateStatusCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStateStatusCounters.Merge(m, src)
}
func (m *RsvpMgmtFrrStateStatusCounters) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStateStatusCounters.Size(m)
}
func (m *RsvpMgmtFrrStateStatusCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStateStatusCounters.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStateStatusCounters proto.InternalMessageInfo

func (m *RsvpMgmtFrrStateStatusCounters) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RsvpMgmtFrrStateStatusCounters) GetActiveInstances() uint32 {
	if m != nil {
		return m.ActiveInstances
	}
	return 0
}

func (m *RsvpMgmtFrrStateStatusCounters) GetReadyInstances() uint32 {
	if m != nil {
		return m.ReadyInstances
	}
	return 0
}

func (m *RsvpMgmtFrrStateStatusCounters) GetActiveWaitInstances() uint32 {
	if m != nil {
		return m.ActiveWaitInstances
	}
	return 0
}

type RsvpMgmtFrrStatesCompact struct {
	PathStates           *RsvpMgmtFrrStateStatusCounters `protobuf:"bytes,50,opt,name=path_states,json=pathStates,proto3" json:"path_states,omitempty"`
	ReservationStates    *RsvpMgmtFrrStateStatusCounters `protobuf:"bytes,51,opt,name=reservation_states,json=reservationStates,proto3" json:"reservation_states,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RsvpMgmtFrrStatesCompact) Reset()         { *m = RsvpMgmtFrrStatesCompact{} }
func (m *RsvpMgmtFrrStatesCompact) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesCompact) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesCompact) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e569263222c28bc, []int{2}
}

func (m *RsvpMgmtFrrStatesCompact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesCompact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtFrrStatesCompact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesCompact.Merge(m, src)
}
func (m *RsvpMgmtFrrStatesCompact) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesCompact.Size(m)
}
func (m *RsvpMgmtFrrStatesCompact) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesCompact.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesCompact proto.InternalMessageInfo

func (m *RsvpMgmtFrrStatesCompact) GetPathStates() *RsvpMgmtFrrStateStatusCounters {
	if m != nil {
		return m.PathStates
	}
	return nil
}

func (m *RsvpMgmtFrrStatesCompact) GetReservationStates() *RsvpMgmtFrrStateStatusCounters {
	if m != nil {
		return m.ReservationStates
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtFrrStatesCompact_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frr_summary.rsvp_mgmt_frr_states_compact_KEYS")
	proto.RegisterType((*RsvpMgmtFrrStateStatusCounters)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frr_summary.rsvp_mgmt_frr_state_status_counters")
	proto.RegisterType((*RsvpMgmtFrrStatesCompact)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.frr_summary.rsvp_mgmt_frr_states_compact")
}

func init() { proto.RegisterFile("rsvp_mgmt_frr_states_compact.proto", fileDescriptor_6e569263222c28bc) }

var fileDescriptor_6e569263222c28bc = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0xf2, 0x67, 0xb8, 0x0a, 0x0a, 0x06, 0xa4, 0x0e, 0x0c, 0x25, 0x1d, 0x28, 0x0c,
	0x19, 0xd2, 0x67, 0x60, 0x40, 0x0c, 0x48, 0xed, 0x80, 0x98, 0x4e, 0x87, 0x31, 0x60, 0x89, 0xc4,
	0xd6, 0xf9, 0x12, 0xe8, 0xc2, 0x33, 0xf0, 0x4e, 0xbc, 0x18, 0x8a, 0xd3, 0x42, 0x06, 0x54, 0x75,
	0x61, 0xb1, 0xec, 0x9f, 0xbf, 0xbb, 0x4f, 0x3e, 0x19, 0x52, 0x0e, 0xb5, 0xc7, 0xe2, 0xb9, 0x10,
	0x7c, 0x62, 0xc6, 0x20, 0x24, 0x26, 0xa0, 0x76, 0x85, 0x27, 0x2d, 0x99, 0x67, 0x27, 0x4e, 0x5d,
	0x6a, 0x1b, 0xb4, 0x43, 0xeb, 0x02, 0xbe, 0x33, 0x5a, 0x8f, 0xb1, 0xc6, 0x79, 0xc3, 0x59, 0xb3,
	0xcb, 0x62, 0x61, 0x55, 0x14, 0xc4, 0x8b, 0x74, 0x0c, 0x67, 0xeb, 0x3a, 0xe2, 0xcd, 0xd5, 0xfd,
	0x3c, 0xfd, 0x4a, 0x60, 0xfc, 0x07, 0x15, 0xd7, 0xaa, 0x61, 0xab, 0x52, 0x0c, 0x07, 0x75, 0x0c,
	0x3b, 0xe2, 0x84, 0x5e, 0x87, 0xc9, 0x28, 0x99, 0xec, 0xcd, 0xda, 0x83, 0xba, 0x80, 0x03, 0xd2,
	0x62, 0x6b, 0x83, 0xb6, 0x0c, 0x42, 0xa5, 0x36, 0x61, 0xd8, 0x8b, 0xc0, 0xa0, 0xcd, 0xaf, 0x57,
	0xb1, 0x3a, 0x87, 0x01, 0x1b, 0x7a, 0x5c, 0x74, 0xc8, 0xad, 0x48, 0xee, 0xc7, 0xf8, 0x17, 0xcc,
	0xe1, 0x64, 0xd9, 0xf3, 0x8d, 0xac, 0x74, 0xf0, 0xed, 0x88, 0x1f, 0xb5, 0x97, 0x77, 0x64, 0xe5,
	0xa7, 0x26, 0xfd, 0xec, 0xc1, 0xe9, 0xba, 0xb7, 0x2a, 0x0f, 0x7d, 0x4f, 0xf2, 0xb2, 0x8c, 0x87,
	0xf9, 0x28, 0x99, 0xf4, 0xf3, 0xdb, 0x6c, 0xf3, 0x69, 0x66, 0x1b, 0x0c, 0x69, 0x06, 0x8d, 0x63,
	0x1e, 0x15, 0xea, 0x03, 0x14, 0x9b, 0x60, 0xb8, 0x26, 0xb1, 0xae, 0x5c, 0x89, 0xa7, 0xff, 0x23,
	0x3e, 0xec, 0xa8, 0x5a, 0xff, 0xc3, 0x6e, 0xfc, 0x30, 0xd3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x88, 0x83, 0x27, 0xee, 0x56, 0x02, 0x00, 0x00,
}