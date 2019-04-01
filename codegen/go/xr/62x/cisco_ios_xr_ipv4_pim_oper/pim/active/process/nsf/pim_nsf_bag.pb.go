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
// source: pim_nsf_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_process_nsf

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

type PimNsfBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimNsfBag_KEYS) Reset()         { *m = PimNsfBag_KEYS{} }
func (m *PimNsfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimNsfBag_KEYS) ProtoMessage()    {}
func (*PimNsfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_786656eb2778ccff, []int{0}
}

func (m *PimNsfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNsfBag_KEYS.Unmarshal(m, b)
}
func (m *PimNsfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNsfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimNsfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNsfBag_KEYS.Merge(m, src)
}
func (m *PimNsfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimNsfBag_KEYS.Size(m)
}
func (m *PimNsfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNsfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimNsfBag_KEYS proto.InternalMessageInfo

type PimNsfBag struct {
	ConfiguredState      bool     `protobuf:"varint,50,opt,name=configured_state,json=configuredState,proto3" json:"configured_state,omitempty"`
	NsfState             bool     `protobuf:"varint,51,opt,name=nsf_state,json=nsfState,proto3" json:"nsf_state,omitempty"`
	NsfTimeout           uint32   `protobuf:"varint,52,opt,name=nsf_timeout,json=nsfTimeout,proto3" json:"nsf_timeout,omitempty"`
	NsfTimeLeft          uint32   `protobuf:"varint,53,opt,name=nsf_time_left,json=nsfTimeLeft,proto3" json:"nsf_time_left,omitempty"`
	WaitingTimer         bool     `protobuf:"varint,54,opt,name=waiting_timer,json=waitingTimer,proto3" json:"waiting_timer,omitempty"`
	WaitingMembership    bool     `protobuf:"varint,55,opt,name=waiting_membership,json=waitingMembership,proto3" json:"waiting_membership,omitempty"`
	RespawnCount         uint32   `protobuf:"varint,56,opt,name=respawn_count,json=respawnCount,proto3" json:"respawn_count,omitempty"`
	LastNsfOn            string   `protobuf:"bytes,57,opt,name=last_nsf_on,json=lastNsfOn,proto3" json:"last_nsf_on,omitempty"`
	LastNsfOff           string   `protobuf:"bytes,58,opt,name=last_nsf_off,json=lastNsfOff,proto3" json:"last_nsf_off,omitempty"`
	LastNsfOnSec         int32    `protobuf:"zigzag32,59,opt,name=last_nsf_on_sec,json=lastNsfOnSec,proto3" json:"last_nsf_on_sec,omitempty"`
	LastNsfOffSec        int32    `protobuf:"zigzag32,60,opt,name=last_nsf_off_sec,json=lastNsfOffSec,proto3" json:"last_nsf_off_sec,omitempty"`
	LastIcdNotifRecv     string   `protobuf:"bytes,61,opt,name=last_icd_notif_recv,json=lastIcdNotifRecv,proto3" json:"last_icd_notif_recv,omitempty"`
	LastIcdNotifRecvSec  int32    `protobuf:"zigzag32,62,opt,name=last_icd_notif_recv_sec,json=lastIcdNotifRecvSec,proto3" json:"last_icd_notif_recv_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimNsfBag) Reset()         { *m = PimNsfBag{} }
func (m *PimNsfBag) String() string { return proto.CompactTextString(m) }
func (*PimNsfBag) ProtoMessage()    {}
func (*PimNsfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_786656eb2778ccff, []int{1}
}

func (m *PimNsfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNsfBag.Unmarshal(m, b)
}
func (m *PimNsfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNsfBag.Marshal(b, m, deterministic)
}
func (m *PimNsfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNsfBag.Merge(m, src)
}
func (m *PimNsfBag) XXX_Size() int {
	return xxx_messageInfo_PimNsfBag.Size(m)
}
func (m *PimNsfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNsfBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimNsfBag proto.InternalMessageInfo

func (m *PimNsfBag) GetConfiguredState() bool {
	if m != nil {
		return m.ConfiguredState
	}
	return false
}

func (m *PimNsfBag) GetNsfState() bool {
	if m != nil {
		return m.NsfState
	}
	return false
}

func (m *PimNsfBag) GetNsfTimeout() uint32 {
	if m != nil {
		return m.NsfTimeout
	}
	return 0
}

func (m *PimNsfBag) GetNsfTimeLeft() uint32 {
	if m != nil {
		return m.NsfTimeLeft
	}
	return 0
}

func (m *PimNsfBag) GetWaitingTimer() bool {
	if m != nil {
		return m.WaitingTimer
	}
	return false
}

func (m *PimNsfBag) GetWaitingMembership() bool {
	if m != nil {
		return m.WaitingMembership
	}
	return false
}

func (m *PimNsfBag) GetRespawnCount() uint32 {
	if m != nil {
		return m.RespawnCount
	}
	return 0
}

func (m *PimNsfBag) GetLastNsfOn() string {
	if m != nil {
		return m.LastNsfOn
	}
	return ""
}

func (m *PimNsfBag) GetLastNsfOff() string {
	if m != nil {
		return m.LastNsfOff
	}
	return ""
}

func (m *PimNsfBag) GetLastNsfOnSec() int32 {
	if m != nil {
		return m.LastNsfOnSec
	}
	return 0
}

func (m *PimNsfBag) GetLastNsfOffSec() int32 {
	if m != nil {
		return m.LastNsfOffSec
	}
	return 0
}

func (m *PimNsfBag) GetLastIcdNotifRecv() string {
	if m != nil {
		return m.LastIcdNotifRecv
	}
	return ""
}

func (m *PimNsfBag) GetLastIcdNotifRecvSec() int32 {
	if m != nil {
		return m.LastIcdNotifRecvSec
	}
	return 0
}

func init() {
	proto.RegisterType((*PimNsfBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.nsf.pim_nsf_bag_KEYS")
	proto.RegisterType((*PimNsfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.nsf.pim_nsf_bag")
}

func init() { proto.RegisterFile("pim_nsf_bag.proto", fileDescriptor_786656eb2778ccff) }

var fileDescriptor_786656eb2778ccff = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0x06, 0x70, 0x0a, 0x22, 0x3b, 0x6f, 0x5b, 0x76, 0x26, 0x7b, 0x30, 0x20, 0x68, 0x19, 0x11,
	0xeb, 0x61, 0x0b, 0xba, 0xe3, 0xff, 0x3f, 0x17, 0xf1, 0x20, 0xea, 0x0a, 0xad, 0x17, 0x4f, 0xa1,
	0x93, 0x79, 0x33, 0x06, 0xb6, 0x49, 0x49, 0x32, 0x5d, 0x3f, 0x83, 0x9f, 0x5a, 0xf2, 0x4e, 0x67,
	0xa7, 0xc8, 0x5e, 0x9f, 0xe7, 0x97, 0xe7, 0x85, 0x52, 0x58, 0xf4, 0xba, 0x13, 0xc6, 0x2b, 0xb1,
	0x6e, 0xb7, 0x55, 0xef, 0x6c, 0xb0, 0xec, 0x99, 0xd4, 0x5e, 0x5a, 0xa1, 0xad, 0x17, 0x7f, 0x9c,
	0xd0, 0xfd, 0xb0, 0x12, 0x11, 0xd9, 0x1e, 0x5d, 0xd5, 0xeb, 0xae, 0x6a, 0x65, 0xd0, 0x03, 0x46,
	0x2c, 0xd1, 0xfb, 0xca, 0x78, 0xb5, 0x64, 0x30, 0x9f, 0xec, 0x88, 0xaf, 0x9f, 0x7f, 0x35, 0xcb,
	0xbf, 0x77, 0x20, 0x9d, 0x84, 0xec, 0x29, 0xcc, 0xa5, 0x35, 0x4a, 0x6f, 0x77, 0x0e, 0x37, 0xc2,
	0x87, 0x36, 0x20, 0x7f, 0x5e, 0x24, 0xe5, 0x49, 0x7d, 0x7a, 0xcc, 0x9b, 0x18, 0xb3, 0xfb, 0x30,
	0x8b, 0xaf, 0xf6, 0xe6, 0x82, 0xcc, 0x89, 0xf1, 0x6a, 0x5f, 0x3e, 0x84, 0x34, 0x96, 0x41, 0x77,
	0x68, 0x77, 0x81, 0xaf, 0x8a, 0xa4, 0xcc, 0x6b, 0x30, 0x5e, 0xfd, 0xdc, 0x27, 0x6c, 0x09, 0xf9,
	0x01, 0x88, 0x2b, 0x54, 0x81, 0xbf, 0x20, 0x92, 0x8e, 0xe4, 0x1b, 0xaa, 0xc0, 0x1e, 0x41, 0x7e,
	0xdd, 0xea, 0xa0, 0xcd, 0x96, 0x9c, 0xe3, 0x2f, 0xe9, 0x4a, 0x36, 0x86, 0xd1, 0x39, 0x76, 0x0e,
	0xec, 0x80, 0x3a, 0xec, 0xd6, 0xe8, 0xfc, 0x6f, 0xdd, 0xf3, 0x57, 0x24, 0x17, 0x63, 0xf3, 0xfd,
	0xa6, 0x88, 0x9b, 0x0e, 0x7d, 0xdf, 0x5e, 0x1b, 0x21, 0xed, 0xce, 0x04, 0xfe, 0x9a, 0xee, 0x66,
	0x63, 0xf8, 0x29, 0x66, 0xec, 0x01, 0xa4, 0x57, 0xad, 0x0f, 0xf4, 0x55, 0xac, 0xe1, 0x6f, 0x8a,
	0xa4, 0x9c, 0xd5, 0xb3, 0x18, 0x5d, 0x7a, 0xf5, 0xc3, 0xb0, 0x02, 0xb2, 0x63, 0xaf, 0x14, 0x7f,
	0x4b, 0x00, 0x0e, 0x40, 0x29, 0xf6, 0x18, 0x4e, 0x27, 0x0b, 0xc2, 0xa3, 0xe4, 0xef, 0x8a, 0xa4,
	0x5c, 0xd4, 0xd9, 0xcd, 0x4a, 0x83, 0x92, 0x3d, 0x81, 0xf9, 0x74, 0x88, 0xdc, 0x7b, 0x72, 0xf9,
	0x71, 0x2c, 0xc2, 0x73, 0x38, 0x23, 0xa8, 0xe5, 0x46, 0x18, 0x1b, 0xb4, 0x12, 0x0e, 0xe5, 0xc0,
	0x3f, 0xd0, 0x61, 0xda, 0xf8, 0x22, 0x37, 0x97, 0xb1, 0xa8, 0x51, 0x0e, 0x6c, 0x05, 0xf7, 0x6e,
	0xe1, 0x34, 0xff, 0x91, 0xe6, 0xcf, 0xfe, 0x7f, 0xd2, 0xa0, 0x5c, 0xdf, 0xa5, 0x5f, 0xeb, 0xe2,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0xd5, 0x62, 0x9a, 0x6f, 0x02, 0x00, 0x00,
}