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
// source: ptp_parent_ds_info.proto

package cisco_ios_xr_ptp_oper_ptp_dataset_parent_ds

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

type PtpParentDsInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpParentDsInfo_KEYS) Reset()         { *m = PtpParentDsInfo_KEYS{} }
func (m *PtpParentDsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpParentDsInfo_KEYS) ProtoMessage()    {}
func (*PtpParentDsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe55ecc32e12e119, []int{0}
}

func (m *PtpParentDsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpParentDsInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpParentDsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpParentDsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpParentDsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpParentDsInfo_KEYS.Merge(m, src)
}
func (m *PtpParentDsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpParentDsInfo_KEYS.Size(m)
}
func (m *PtpParentDsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpParentDsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpParentDsInfo_KEYS proto.InternalMessageInfo

type PtpParentDsInfo struct {
	ParentClockId                      uint64   `protobuf:"varint,50,opt,name=parent_clock_id,json=parentClockId,proto3" json:"parent_clock_id,omitempty"`
	ParentPortNumber                   uint32   `protobuf:"varint,51,opt,name=parent_port_number,json=parentPortNumber,proto3" json:"parent_port_number,omitempty"`
	ParentStats                        bool     `protobuf:"varint,52,opt,name=parent_stats,json=parentStats,proto3" json:"parent_stats,omitempty"`
	ObservedParentOslv                 uint32   `protobuf:"varint,53,opt,name=observed_parent_oslv,json=observedParentOslv,proto3" json:"observed_parent_oslv,omitempty"`
	ObservedParentClockPhaseChangeRate uint32   `protobuf:"varint,54,opt,name=observed_parent_clock_phase_change_rate,json=observedParentClockPhaseChangeRate,proto3" json:"observed_parent_clock_phase_change_rate,omitempty"`
	GmClockId                          uint64   `protobuf:"varint,55,opt,name=gm_clock_id,json=gmClockId,proto3" json:"gm_clock_id,omitempty"`
	GmClockClass                       uint32   `protobuf:"varint,56,opt,name=gm_clock_class,json=gmClockClass,proto3" json:"gm_clock_class,omitempty"`
	GmClockAccuracy                    uint32   `protobuf:"varint,57,opt,name=gm_clock_accuracy,json=gmClockAccuracy,proto3" json:"gm_clock_accuracy,omitempty"`
	Gmoslv                             uint32   `protobuf:"varint,58,opt,name=gmoslv,proto3" json:"gmoslv,omitempty"`
	GmPriority1                        uint32   `protobuf:"varint,59,opt,name=gm_priority1,json=gmPriority1,proto3" json:"gm_priority1,omitempty"`
	GmPriority2                        uint32   `protobuf:"varint,60,opt,name=gm_priority2,json=gmPriority2,proto3" json:"gm_priority2,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *PtpParentDsInfo) Reset()         { *m = PtpParentDsInfo{} }
func (m *PtpParentDsInfo) String() string { return proto.CompactTextString(m) }
func (*PtpParentDsInfo) ProtoMessage()    {}
func (*PtpParentDsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe55ecc32e12e119, []int{1}
}

func (m *PtpParentDsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpParentDsInfo.Unmarshal(m, b)
}
func (m *PtpParentDsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpParentDsInfo.Marshal(b, m, deterministic)
}
func (m *PtpParentDsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpParentDsInfo.Merge(m, src)
}
func (m *PtpParentDsInfo) XXX_Size() int {
	return xxx_messageInfo_PtpParentDsInfo.Size(m)
}
func (m *PtpParentDsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpParentDsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpParentDsInfo proto.InternalMessageInfo

func (m *PtpParentDsInfo) GetParentClockId() uint64 {
	if m != nil {
		return m.ParentClockId
	}
	return 0
}

func (m *PtpParentDsInfo) GetParentPortNumber() uint32 {
	if m != nil {
		return m.ParentPortNumber
	}
	return 0
}

func (m *PtpParentDsInfo) GetParentStats() bool {
	if m != nil {
		return m.ParentStats
	}
	return false
}

func (m *PtpParentDsInfo) GetObservedParentOslv() uint32 {
	if m != nil {
		return m.ObservedParentOslv
	}
	return 0
}

func (m *PtpParentDsInfo) GetObservedParentClockPhaseChangeRate() uint32 {
	if m != nil {
		return m.ObservedParentClockPhaseChangeRate
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmClockId() uint64 {
	if m != nil {
		return m.GmClockId
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmClockClass() uint32 {
	if m != nil {
		return m.GmClockClass
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmClockAccuracy() uint32 {
	if m != nil {
		return m.GmClockAccuracy
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmoslv() uint32 {
	if m != nil {
		return m.Gmoslv
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmPriority1() uint32 {
	if m != nil {
		return m.GmPriority1
	}
	return 0
}

func (m *PtpParentDsInfo) GetGmPriority2() uint32 {
	if m != nil {
		return m.GmPriority2
	}
	return 0
}

func init() {
	proto.RegisterType((*PtpParentDsInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.dataset.parent_ds.ptp_parent_ds_info_KEYS")
	proto.RegisterType((*PtpParentDsInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.dataset.parent_ds.ptp_parent_ds_info")
}

func init() { proto.RegisterFile("ptp_parent_ds_info.proto", fileDescriptor_fe55ecc32e12e119) }

var fileDescriptor_fe55ecc32e12e119 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x1c, 0x86, 0x19, 0xca, 0xd0, 0x6c, 0x73, 0x1a, 0x44, 0xe3, 0x45, 0xe6, 0x10, 0x1d, 0x2a, 0x45,
	0x37, 0xff, 0xeb, 0x45, 0x86, 0x07, 0x11, 0xb4, 0x74, 0x27, 0x4f, 0x21, 0x4b, 0x63, 0x57, 0x6c,
	0x9b, 0x90, 0x64, 0xc3, 0x7d, 0x5f, 0x3f, 0x88, 0xf4, 0xd7, 0x74, 0x62, 0x77, 0xec, 0xf3, 0x3e,
	0x79, 0xc9, 0x5b, 0x82, 0x88, 0xb2, 0x8a, 0x2a, 0xa6, 0x45, 0x66, 0x69, 0x68, 0x68, 0x9c, 0x7d,
	0x4a, 0x4f, 0x69, 0x69, 0x25, 0x3e, 0xe5, 0xb1, 0xe1, 0x92, 0xc6, 0xd2, 0xd0, 0x6f, 0x4d, 0x73,
	0x4d, 0x2a, 0xa1, 0x3d, 0x65, 0x95, 0x17, 0x32, 0xcb, 0x8c, 0xb0, 0xde, 0xe2, 0x5c, 0x77, 0x0f,
	0xed, 0x2e, 0x17, 0xd1, 0xd7, 0xe7, 0x8f, 0x51, 0xf7, 0x67, 0x05, 0xe1, 0xe5, 0x0c, 0x1f, 0xa1,
	0xb6, 0x23, 0x3c, 0x91, 0xfc, 0x8b, 0xc6, 0x21, 0xe9, 0x77, 0x6a, 0xbd, 0xd5, 0xa0, 0x55, 0xe0,
	0x61, 0x4e, 0x5f, 0x42, 0x7c, 0x86, 0xb0, 0xf3, 0x94, 0xd4, 0x96, 0x66, 0xd3, 0x74, 0x2c, 0x34,
	0x19, 0x74, 0x6a, 0xbd, 0x56, 0xb0, 0x59, 0x24, 0xbe, 0xd4, 0xf6, 0x0d, 0x38, 0x3e, 0x40, 0x4d,
	0x67, 0x1b, 0xcb, 0xac, 0x21, 0x97, 0x9d, 0x5a, 0x6f, 0x2d, 0x68, 0x14, 0x6c, 0x94, 0x23, 0x7c,
	0x8e, 0xb6, 0xe5, 0xd8, 0x08, 0x3d, 0x13, 0x61, 0x79, 0x27, 0x69, 0x92, 0x19, 0xb9, 0x82, 0x4a,
	0x5c, 0x66, 0x3e, 0x44, 0xef, 0x26, 0x99, 0xe1, 0x11, 0x3a, 0xae, 0x9e, 0x28, 0xee, 0xac, 0x26,
	0xcc, 0x08, 0xca, 0x27, 0x2c, 0x8b, 0x04, 0xd5, 0xcc, 0x0a, 0x72, 0x0d, 0x25, 0xdd, 0xff, 0x25,
	0x30, 0xc5, 0xcf, 0xdd, 0x21, 0xa8, 0x01, 0xb3, 0x02, 0xef, 0xa3, 0x46, 0x94, 0xfe, 0x6d, 0xbf,
	0x81, 0xed, 0xeb, 0x51, 0x5a, 0xee, 0x3e, 0x44, 0x1b, 0x8b, 0x9c, 0x27, 0xcc, 0x18, 0x72, 0x0b,
	0xdd, 0x4d, 0xa7, 0x0c, 0x73, 0x86, 0x4f, 0xd0, 0xd6, 0xc2, 0x62, 0x9c, 0x4f, 0x35, 0xe3, 0x73,
	0x72, 0x07, 0x62, 0xdb, 0x89, 0x4f, 0x0e, 0xe3, 0x1d, 0x54, 0x8f, 0x52, 0x98, 0x7a, 0x0f, 0x82,
	0xfb, 0xca, 0xff, 0x59, 0x94, 0x52, 0xa5, 0x63, 0xa9, 0x63, 0x3b, 0xbf, 0x20, 0x0f, 0x90, 0x36,
	0xa2, 0xd4, 0x2f, 0x51, 0x45, 0xe9, 0x93, 0xc7, 0xaa, 0xd2, 0x1f, 0xd7, 0xe1, 0xd5, 0x0c, 0x7e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x2e, 0x94, 0x34, 0x51, 0x02, 0x00, 0x00,
}