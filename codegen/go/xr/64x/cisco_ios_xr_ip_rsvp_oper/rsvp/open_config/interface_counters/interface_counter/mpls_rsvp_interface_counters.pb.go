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
// source: mpls_rsvp_interface_counters.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_open_config_interface_counters_interface_counter

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

type MplsRsvpInterfaceCounters_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsRsvpInterfaceCounters_KEYS) Reset()         { *m = MplsRsvpInterfaceCounters_KEYS{} }
func (m *MplsRsvpInterfaceCounters_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsRsvpInterfaceCounters_KEYS) ProtoMessage()    {}
func (*MplsRsvpInterfaceCounters_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eaa1309eb95d18a, []int{0}
}

func (m *MplsRsvpInterfaceCounters_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS.Unmarshal(m, b)
}
func (m *MplsRsvpInterfaceCounters_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsRsvpInterfaceCounters_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS.Merge(m, src)
}
func (m *MplsRsvpInterfaceCounters_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS.Size(m)
}
func (m *MplsRsvpInterfaceCounters_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsRsvpInterfaceCounters_KEYS proto.InternalMessageInfo

func (m *MplsRsvpInterfaceCounters_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsRsvpInterfaceCounters struct {
	InterfaceNameXr             string   `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	InPathMessages              uint64   `protobuf:"varint,51,opt,name=in_path_messages,json=inPathMessages,proto3" json:"in_path_messages,omitempty"`
	InPathErrorMessages         uint64   `protobuf:"varint,52,opt,name=in_path_error_messages,json=inPathErrorMessages,proto3" json:"in_path_error_messages,omitempty"`
	InPathTearMessages          uint64   `protobuf:"varint,53,opt,name=in_path_tear_messages,json=inPathTearMessages,proto3" json:"in_path_tear_messages,omitempty"`
	InReservationMessages       uint64   `protobuf:"varint,54,opt,name=in_reservation_messages,json=inReservationMessages,proto3" json:"in_reservation_messages,omitempty"`
	InReservationErrorMessages  uint64   `protobuf:"varint,55,opt,name=in_reservation_error_messages,json=inReservationErrorMessages,proto3" json:"in_reservation_error_messages,omitempty"`
	InReservationTearMessages   uint64   `protobuf:"varint,56,opt,name=in_reservation_tear_messages,json=inReservationTearMessages,proto3" json:"in_reservation_tear_messages,omitempty"`
	InHelloMessages             uint64   `protobuf:"varint,57,opt,name=in_hello_messages,json=inHelloMessages,proto3" json:"in_hello_messages,omitempty"`
	InSrefreshMessages          uint64   `protobuf:"varint,58,opt,name=in_srefresh_messages,json=inSrefreshMessages,proto3" json:"in_srefresh_messages,omitempty"`
	InAckMessages               uint64   `protobuf:"varint,59,opt,name=in_ack_messages,json=inAckMessages,proto3" json:"in_ack_messages,omitempty"`
	OutPathMessages             uint64   `protobuf:"varint,60,opt,name=out_path_messages,json=outPathMessages,proto3" json:"out_path_messages,omitempty"`
	OutPathErrorMessages        uint64   `protobuf:"varint,61,opt,name=out_path_error_messages,json=outPathErrorMessages,proto3" json:"out_path_error_messages,omitempty"`
	OutPathTearMessages         uint64   `protobuf:"varint,62,opt,name=out_path_tear_messages,json=outPathTearMessages,proto3" json:"out_path_tear_messages,omitempty"`
	OutReservationMessages      uint64   `protobuf:"varint,63,opt,name=out_reservation_messages,json=outReservationMessages,proto3" json:"out_reservation_messages,omitempty"`
	OutReservationErrorMessages uint64   `protobuf:"varint,64,opt,name=out_reservation_error_messages,json=outReservationErrorMessages,proto3" json:"out_reservation_error_messages,omitempty"`
	OutReservationTearMessages  uint64   `protobuf:"varint,65,opt,name=out_reservation_tear_messages,json=outReservationTearMessages,proto3" json:"out_reservation_tear_messages,omitempty"`
	OutHelloMessages            uint64   `protobuf:"varint,66,opt,name=out_hello_messages,json=outHelloMessages,proto3" json:"out_hello_messages,omitempty"`
	OutSrefreshMessages         uint64   `protobuf:"varint,67,opt,name=out_srefresh_messages,json=outSrefreshMessages,proto3" json:"out_srefresh_messages,omitempty"`
	OutAckMessages              uint64   `protobuf:"varint,68,opt,name=out_ack_messages,json=outAckMessages,proto3" json:"out_ack_messages,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MplsRsvpInterfaceCounters) Reset()         { *m = MplsRsvpInterfaceCounters{} }
func (m *MplsRsvpInterfaceCounters) String() string { return proto.CompactTextString(m) }
func (*MplsRsvpInterfaceCounters) ProtoMessage()    {}
func (*MplsRsvpInterfaceCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_7eaa1309eb95d18a, []int{1}
}

func (m *MplsRsvpInterfaceCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsRsvpInterfaceCounters.Unmarshal(m, b)
}
func (m *MplsRsvpInterfaceCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsRsvpInterfaceCounters.Marshal(b, m, deterministic)
}
func (m *MplsRsvpInterfaceCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsRsvpInterfaceCounters.Merge(m, src)
}
func (m *MplsRsvpInterfaceCounters) XXX_Size() int {
	return xxx_messageInfo_MplsRsvpInterfaceCounters.Size(m)
}
func (m *MplsRsvpInterfaceCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsRsvpInterfaceCounters.DiscardUnknown(m)
}

var xxx_messageInfo_MplsRsvpInterfaceCounters proto.InternalMessageInfo

func (m *MplsRsvpInterfaceCounters) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *MplsRsvpInterfaceCounters) GetInPathMessages() uint64 {
	if m != nil {
		return m.InPathMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInPathErrorMessages() uint64 {
	if m != nil {
		return m.InPathErrorMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInPathTearMessages() uint64 {
	if m != nil {
		return m.InPathTearMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInReservationMessages() uint64 {
	if m != nil {
		return m.InReservationMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInReservationErrorMessages() uint64 {
	if m != nil {
		return m.InReservationErrorMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInReservationTearMessages() uint64 {
	if m != nil {
		return m.InReservationTearMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInHelloMessages() uint64 {
	if m != nil {
		return m.InHelloMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInSrefreshMessages() uint64 {
	if m != nil {
		return m.InSrefreshMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetInAckMessages() uint64 {
	if m != nil {
		return m.InAckMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutPathMessages() uint64 {
	if m != nil {
		return m.OutPathMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutPathErrorMessages() uint64 {
	if m != nil {
		return m.OutPathErrorMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutPathTearMessages() uint64 {
	if m != nil {
		return m.OutPathTearMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutReservationMessages() uint64 {
	if m != nil {
		return m.OutReservationMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutReservationErrorMessages() uint64 {
	if m != nil {
		return m.OutReservationErrorMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutReservationTearMessages() uint64 {
	if m != nil {
		return m.OutReservationTearMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutHelloMessages() uint64 {
	if m != nil {
		return m.OutHelloMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutSrefreshMessages() uint64 {
	if m != nil {
		return m.OutSrefreshMessages
	}
	return 0
}

func (m *MplsRsvpInterfaceCounters) GetOutAckMessages() uint64 {
	if m != nil {
		return m.OutAckMessages
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsRsvpInterfaceCounters_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.open_config.interface_counters.interface_counter.mpls_rsvp_interface_counters_KEYS")
	proto.RegisterType((*MplsRsvpInterfaceCounters)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.open_config.interface_counters.interface_counter.mpls_rsvp_interface_counters")
}

func init() { proto.RegisterFile("mpls_rsvp_interface_counters.proto", fileDescriptor_7eaa1309eb95d18a) }

var fileDescriptor_7eaa1309eb95d18a = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5f, 0x6b, 0x13, 0x41,
	0x14, 0xc5, 0x09, 0x88, 0xe8, 0x85, 0x36, 0x75, 0x6d, 0xda, 0xa8, 0xad, 0xd4, 0x80, 0x12, 0x44,
	0x82, 0x36, 0xb6, 0xd6, 0xbf, 0x35, 0xd6, 0x82, 0x28, 0xfe, 0x21, 0xf5, 0x41, 0x9f, 0x86, 0x71,
	0xb9, 0x69, 0x86, 0x66, 0x67, 0x96, 0x99, 0xd9, 0xd2, 0x0f, 0xe0, 0x07, 0x2f, 0x33, 0xbb, 0x99,
	0xdd, 0x3b, 0x09, 0x79, 0x5b, 0x66, 0xce, 0xef, 0xe4, 0xe4, 0x9c, 0x4d, 0xa0, 0x97, 0xe5, 0x33,
	0xc3, 0xb4, 0xb9, 0xcc, 0x99, 0x90, 0x16, 0xf5, 0x84, 0xa7, 0xc8, 0x52, 0x55, 0xb8, 0x67, 0x33,
	0xc8, 0xb5, 0xb2, 0x2a, 0xf9, 0x99, 0x0a, 0x93, 0x2a, 0x26, 0x94, 0x61, 0x57, 0x9a, 0x89, 0xbc,
	0x94, 0xab, 0x1c, 0xf5, 0xc0, 0x3d, 0x0d, 0x54, 0x8e, 0x92, 0xa5, 0x4a, 0x4e, 0xc4, 0xf9, 0x60,
	0x89, 0xc9, 0xc2, 0x51, 0xef, 0x2b, 0x3c, 0x5a, 0xf5, 0xb1, 0xec, 0xdb, 0xe9, 0xdf, 0xb3, 0xe4,
	0x31, 0xac, 0xd7, 0x57, 0x92, 0x67, 0xd8, 0x6d, 0xed, 0xb5, 0xfa, 0xb7, 0xc7, 0x6b, 0xe1, 0xf4,
	0x07, 0xcf, 0xb0, 0xf7, 0xff, 0x16, 0xec, 0xac, 0x32, 0x4b, 0x9e, 0xc2, 0x1d, 0xea, 0xc3, 0xae,
	0x74, 0x77, 0xdf, 0x5b, 0xb5, 0x89, 0xd5, 0x1f, 0x9d, 0xf4, 0x61, 0x43, 0x48, 0x96, 0x73, 0x3b,
	0x65, 0x19, 0x1a, 0xc3, 0xcf, 0xd1, 0x74, 0x87, 0x7b, 0xad, 0xfe, 0x8d, 0xf1, 0xba, 0x90, 0xbf,
	0xb8, 0x9d, 0x7e, 0xaf, 0x4e, 0x93, 0x21, 0x6c, 0xcd, 0x95, 0xa8, 0xb5, 0xd2, 0xb5, 0xfe, 0xa5,
	0xd7, 0xdf, 0x2d, 0xf5, 0xa7, 0xee, 0x2e, 0x40, 0x2f, 0xa0, 0x33, 0x87, 0x2c, 0xf2, 0x06, 0x73,
	0xe0, 0x99, 0xa4, 0x64, 0x7e, 0x23, 0xaf, 0x91, 0x43, 0xd8, 0x16, 0x92, 0x69, 0x34, 0xa8, 0x2f,
	0xb9, 0x15, 0x4a, 0xd6, 0xd0, 0xa1, 0x87, 0x3a, 0x42, 0x8e, 0xeb, 0xdb, 0xc0, 0x8d, 0x60, 0x37,
	0xe2, 0xa2, 0x98, 0xaf, 0x3c, 0x7d, 0x9f, 0xd0, 0x34, 0xed, 0x31, 0xec, 0x44, 0x16, 0x34, 0xf4,
	0x91, 0x77, 0xb8, 0x47, 0x1c, 0x48, 0x76, 0xdf, 0x3c, 0x9b, 0xe2, 0x6c, 0xa6, 0x6a, 0xea, 0xb5,
	0xa7, 0xda, 0x42, 0x7e, 0x71, 0xe7, 0x41, 0xfb, 0x1c, 0x36, 0x85, 0x64, 0x46, 0xe3, 0x44, 0xa3,
	0x69, 0xb4, 0xff, 0x66, 0xde, 0xcc, 0x59, 0x75, 0x15, 0x88, 0x27, 0xd0, 0x16, 0x92, 0xf1, 0xf4,
	0xa2, 0x16, 0xbf, 0xf5, 0xe2, 0x35, 0x21, 0x47, 0xe9, 0x45, 0x33, 0x85, 0x2a, 0x6c, 0x34, 0xea,
	0xbb, 0x32, 0x85, 0x2a, 0x2c, 0x59, 0xf5, 0x00, 0xb6, 0x83, 0x36, 0xea, 0xeb, 0xbd, 0x27, 0x36,
	0x2b, 0x82, 0x36, 0x35, 0x84, 0xad, 0x80, 0xd1, 0x8e, 0x3e, 0x94, 0x2f, 0x43, 0x45, 0x91, 0x76,
	0x8e, 0xa0, 0xeb, 0xa0, 0xa5, 0xd3, 0x1e, 0x7b, 0xcc, 0x99, 0x2e, 0xdb, 0xf6, 0x04, 0x1e, 0xc6,
	0x64, 0x14, 0xf6, 0xa3, 0xe7, 0x1f, 0x50, 0x9e, 0x66, 0x1e, 0xc1, 0x6e, 0x6c, 0x42, 0xa3, 0x8f,
	0xca, 0x17, 0x84, 0x7a, 0x90, 0x6f, 0xf0, 0x0c, 0x12, 0x67, 0x11, 0x0d, 0xfc, 0xc9, 0x73, 0x1b,
	0xaa, 0xb0, 0x74, 0xe1, 0x7d, 0xe8, 0x38, 0xf5, 0xe2, 0xc4, 0x27, 0xa1, 0xa3, 0x85, 0x8d, 0xfb,
	0xe0, 0x7c, 0xe8, 0xc8, 0x9f, 0xcb, 0xdf, 0xa3, 0x2a, 0x6c, 0x63, 0xe5, 0x7f, 0x37, 0xfd, 0x5f,
	0xd5, 0xf0, 0x3a, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x83, 0x38, 0x1a, 0xd0, 0x04, 0x00, 0x00,
}
