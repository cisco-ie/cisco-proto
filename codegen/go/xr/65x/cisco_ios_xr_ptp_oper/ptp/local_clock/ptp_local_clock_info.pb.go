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
// source: ptp_local_clock_info.proto

package cisco_ios_xr_ptp_oper_ptp_local_clock

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

type PtpLocalClockInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpLocalClockInfo_KEYS) Reset()         { *m = PtpLocalClockInfo_KEYS{} }
func (m *PtpLocalClockInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpLocalClockInfo_KEYS) ProtoMessage()    {}
func (*PtpLocalClockInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{0}
}

func (m *PtpLocalClockInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpLocalClockInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpLocalClockInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpLocalClockInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpLocalClockInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpLocalClockInfo_KEYS.Merge(m, src)
}
func (m *PtpLocalClockInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpLocalClockInfo_KEYS.Size(m)
}
func (m *PtpLocalClockInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpLocalClockInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpLocalClockInfo_KEYS proto.InternalMessageInfo

type PtpBagClockUtcOffset struct {
	CurrentOffset        string   `protobuf:"bytes,1,opt,name=current_offset,json=currentOffset,proto3" json:"current_offset,omitempty"`
	OffsetValid          bool     `protobuf:"varint,2,opt,name=offset_valid,json=offsetValid,proto3" json:"offset_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagClockUtcOffset) Reset()         { *m = PtpBagClockUtcOffset{} }
func (m *PtpBagClockUtcOffset) String() string { return proto.CompactTextString(m) }
func (*PtpBagClockUtcOffset) ProtoMessage()    {}
func (*PtpBagClockUtcOffset) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{1}
}

func (m *PtpBagClockUtcOffset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagClockUtcOffset.Unmarshal(m, b)
}
func (m *PtpBagClockUtcOffset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagClockUtcOffset.Marshal(b, m, deterministic)
}
func (m *PtpBagClockUtcOffset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagClockUtcOffset.Merge(m, src)
}
func (m *PtpBagClockUtcOffset) XXX_Size() int {
	return xxx_messageInfo_PtpBagClockUtcOffset.Size(m)
}
func (m *PtpBagClockUtcOffset) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagClockUtcOffset.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagClockUtcOffset proto.InternalMessageInfo

func (m *PtpBagClockUtcOffset) GetCurrentOffset() string {
	if m != nil {
		return m.CurrentOffset
	}
	return ""
}

func (m *PtpBagClockUtcOffset) GetOffsetValid() bool {
	if m != nil {
		return m.OffsetValid
	}
	return false
}

type PtpBagPortId struct {
	ClockId              uint64   `protobuf:"varint,1,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	PortNumber           uint32   `protobuf:"varint,2,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagPortId) Reset()         { *m = PtpBagPortId{} }
func (m *PtpBagPortId) String() string { return proto.CompactTextString(m) }
func (*PtpBagPortId) ProtoMessage()    {}
func (*PtpBagPortId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{2}
}

func (m *PtpBagPortId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagPortId.Unmarshal(m, b)
}
func (m *PtpBagPortId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagPortId.Marshal(b, m, deterministic)
}
func (m *PtpBagPortId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagPortId.Merge(m, src)
}
func (m *PtpBagPortId) XXX_Size() int {
	return xxx_messageInfo_PtpBagPortId.Size(m)
}
func (m *PtpBagPortId) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagPortId.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagPortId proto.InternalMessageInfo

func (m *PtpBagPortId) GetClockId() uint64 {
	if m != nil {
		return m.ClockId
	}
	return 0
}

func (m *PtpBagPortId) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type PtpBagClock struct {
	ClockId              uint64                `protobuf:"varint,1,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	Priority1            uint32                `protobuf:"varint,2,opt,name=priority1,proto3" json:"priority1,omitempty"`
	Priority2            uint32                `protobuf:"varint,3,opt,name=priority2,proto3" json:"priority2,omitempty"`
	Class                uint32                `protobuf:"varint,4,opt,name=class,proto3" json:"class,omitempty"`
	Accuracy             uint32                `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	OffsetLogVariance    uint32                `protobuf:"varint,6,opt,name=offset_log_variance,json=offsetLogVariance,proto3" json:"offset_log_variance,omitempty"`
	StepsRemoved         uint32                `protobuf:"varint,7,opt,name=steps_removed,json=stepsRemoved,proto3" json:"steps_removed,omitempty"`
	TimeSource           string                `protobuf:"bytes,8,opt,name=time_source,json=timeSource,proto3" json:"time_source,omitempty"`
	UtcOffset            *PtpBagClockUtcOffset `protobuf:"bytes,9,opt,name=utc_offset,json=utcOffset,proto3" json:"utc_offset,omitempty"`
	FrequencyTraceable   bool                  `protobuf:"varint,10,opt,name=frequency_traceable,json=frequencyTraceable,proto3" json:"frequency_traceable,omitempty"`
	TimeTraceable        bool                  `protobuf:"varint,11,opt,name=time_traceable,json=timeTraceable,proto3" json:"time_traceable,omitempty"`
	Timescale            string                `protobuf:"bytes,12,opt,name=timescale,proto3" json:"timescale,omitempty"`
	LeapSeconds          string                `protobuf:"bytes,13,opt,name=leap_seconds,json=leapSeconds,proto3" json:"leap_seconds,omitempty"`
	Receiver             *PtpBagPortId         `protobuf:"bytes,14,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender               *PtpBagPortId         `protobuf:"bytes,15,opt,name=sender,proto3" json:"sender,omitempty"`
	Local                bool                  `protobuf:"varint,16,opt,name=local,proto3" json:"local,omitempty"`
	ConfiguredClockClass uint32                `protobuf:"varint,17,opt,name=configured_clock_class,json=configuredClockClass,proto3" json:"configured_clock_class,omitempty"`
	ConfiguredPriority   uint32                `protobuf:"varint,18,opt,name=configured_priority,json=configuredPriority,proto3" json:"configured_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpBagClock) Reset()         { *m = PtpBagClock{} }
func (m *PtpBagClock) String() string { return proto.CompactTextString(m) }
func (*PtpBagClock) ProtoMessage()    {}
func (*PtpBagClock) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{3}
}

func (m *PtpBagClock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagClock.Unmarshal(m, b)
}
func (m *PtpBagClock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagClock.Marshal(b, m, deterministic)
}
func (m *PtpBagClock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagClock.Merge(m, src)
}
func (m *PtpBagClock) XXX_Size() int {
	return xxx_messageInfo_PtpBagClock.Size(m)
}
func (m *PtpBagClock) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagClock.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagClock proto.InternalMessageInfo

func (m *PtpBagClock) GetClockId() uint64 {
	if m != nil {
		return m.ClockId
	}
	return 0
}

func (m *PtpBagClock) GetPriority1() uint32 {
	if m != nil {
		return m.Priority1
	}
	return 0
}

func (m *PtpBagClock) GetPriority2() uint32 {
	if m != nil {
		return m.Priority2
	}
	return 0
}

func (m *PtpBagClock) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *PtpBagClock) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *PtpBagClock) GetOffsetLogVariance() uint32 {
	if m != nil {
		return m.OffsetLogVariance
	}
	return 0
}

func (m *PtpBagClock) GetStepsRemoved() uint32 {
	if m != nil {
		return m.StepsRemoved
	}
	return 0
}

func (m *PtpBagClock) GetTimeSource() string {
	if m != nil {
		return m.TimeSource
	}
	return ""
}

func (m *PtpBagClock) GetUtcOffset() *PtpBagClockUtcOffset {
	if m != nil {
		return m.UtcOffset
	}
	return nil
}

func (m *PtpBagClock) GetFrequencyTraceable() bool {
	if m != nil {
		return m.FrequencyTraceable
	}
	return false
}

func (m *PtpBagClock) GetTimeTraceable() bool {
	if m != nil {
		return m.TimeTraceable
	}
	return false
}

func (m *PtpBagClock) GetTimescale() string {
	if m != nil {
		return m.Timescale
	}
	return ""
}

func (m *PtpBagClock) GetLeapSeconds() string {
	if m != nil {
		return m.LeapSeconds
	}
	return ""
}

func (m *PtpBagClock) GetReceiver() *PtpBagPortId {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *PtpBagClock) GetSender() *PtpBagPortId {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PtpBagClock) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *PtpBagClock) GetConfiguredClockClass() uint32 {
	if m != nil {
		return m.ConfiguredClockClass
	}
	return 0
}

func (m *PtpBagClock) GetConfiguredPriority() uint32 {
	if m != nil {
		return m.ConfiguredPriority
	}
	return 0
}

type PtpBagVirtualPort struct {
	Configured           bool     `protobuf:"varint,1,opt,name=configured,proto3" json:"configured,omitempty"`
	Connected            bool     `protobuf:"varint,2,opt,name=connected,proto3" json:"connected,omitempty"`
	Priority1            uint32   `protobuf:"varint,3,opt,name=priority1,proto3" json:"priority1,omitempty"`
	Priority2            uint32   `protobuf:"varint,4,opt,name=priority2,proto3" json:"priority2,omitempty"`
	Class                uint32   `protobuf:"varint,5,opt,name=class,proto3" json:"class,omitempty"`
	Accuracy             uint32   `protobuf:"varint,6,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	OffsetLogVariance    uint32   `protobuf:"varint,7,opt,name=offset_log_variance,json=offsetLogVariance,proto3" json:"offset_log_variance,omitempty"`
	LocalPriority        uint32   `protobuf:"varint,8,opt,name=local_priority,json=localPriority,proto3" json:"local_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagVirtualPort) Reset()         { *m = PtpBagVirtualPort{} }
func (m *PtpBagVirtualPort) String() string { return proto.CompactTextString(m) }
func (*PtpBagVirtualPort) ProtoMessage()    {}
func (*PtpBagVirtualPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{4}
}

func (m *PtpBagVirtualPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagVirtualPort.Unmarshal(m, b)
}
func (m *PtpBagVirtualPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagVirtualPort.Marshal(b, m, deterministic)
}
func (m *PtpBagVirtualPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagVirtualPort.Merge(m, src)
}
func (m *PtpBagVirtualPort) XXX_Size() int {
	return xxx_messageInfo_PtpBagVirtualPort.Size(m)
}
func (m *PtpBagVirtualPort) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagVirtualPort.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagVirtualPort proto.InternalMessageInfo

func (m *PtpBagVirtualPort) GetConfigured() bool {
	if m != nil {
		return m.Configured
	}
	return false
}

func (m *PtpBagVirtualPort) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *PtpBagVirtualPort) GetPriority1() uint32 {
	if m != nil {
		return m.Priority1
	}
	return 0
}

func (m *PtpBagVirtualPort) GetPriority2() uint32 {
	if m != nil {
		return m.Priority2
	}
	return 0
}

func (m *PtpBagVirtualPort) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *PtpBagVirtualPort) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *PtpBagVirtualPort) GetOffsetLogVariance() uint32 {
	if m != nil {
		return m.OffsetLogVariance
	}
	return 0
}

func (m *PtpBagVirtualPort) GetLocalPriority() uint32 {
	if m != nil {
		return m.LocalPriority
	}
	return 0
}

type PtpLocalClockInfo struct {
	Domain               uint32             `protobuf:"varint,50,opt,name=domain,proto3" json:"domain,omitempty"`
	ClockProperties      *PtpBagClock       `protobuf:"bytes,51,opt,name=clock_properties,json=clockProperties,proto3" json:"clock_properties,omitempty"`
	VirtualPort          *PtpBagVirtualPort `protobuf:"bytes,52,opt,name=virtual_port,json=virtualPort,proto3" json:"virtual_port,omitempty"`
	Holdover             bool               `protobuf:"varint,53,opt,name=holdover,proto3" json:"holdover,omitempty"`
	HoldoverClockClass   uint32             `protobuf:"varint,54,opt,name=holdover_clock_class,json=holdoverClockClass,proto3" json:"holdover_clock_class,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PtpLocalClockInfo) Reset()         { *m = PtpLocalClockInfo{} }
func (m *PtpLocalClockInfo) String() string { return proto.CompactTextString(m) }
func (*PtpLocalClockInfo) ProtoMessage()    {}
func (*PtpLocalClockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab8a038ddedaca8b, []int{5}
}

func (m *PtpLocalClockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpLocalClockInfo.Unmarshal(m, b)
}
func (m *PtpLocalClockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpLocalClockInfo.Marshal(b, m, deterministic)
}
func (m *PtpLocalClockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpLocalClockInfo.Merge(m, src)
}
func (m *PtpLocalClockInfo) XXX_Size() int {
	return xxx_messageInfo_PtpLocalClockInfo.Size(m)
}
func (m *PtpLocalClockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpLocalClockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpLocalClockInfo proto.InternalMessageInfo

func (m *PtpLocalClockInfo) GetDomain() uint32 {
	if m != nil {
		return m.Domain
	}
	return 0
}

func (m *PtpLocalClockInfo) GetClockProperties() *PtpBagClock {
	if m != nil {
		return m.ClockProperties
	}
	return nil
}

func (m *PtpLocalClockInfo) GetVirtualPort() *PtpBagVirtualPort {
	if m != nil {
		return m.VirtualPort
	}
	return nil
}

func (m *PtpLocalClockInfo) GetHoldover() bool {
	if m != nil {
		return m.Holdover
	}
	return false
}

func (m *PtpLocalClockInfo) GetHoldoverClockClass() uint32 {
	if m != nil {
		return m.HoldoverClockClass
	}
	return 0
}

func init() {
	proto.RegisterType((*PtpLocalClockInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_local_clock_info_KEYS")
	proto.RegisterType((*PtpBagClockUtcOffset)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_bag_clock_utc_offset")
	proto.RegisterType((*PtpBagPortId)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_bag_port_id")
	proto.RegisterType((*PtpBagClock)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_bag_clock")
	proto.RegisterType((*PtpBagVirtualPort)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_bag_virtual_port")
	proto.RegisterType((*PtpLocalClockInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.local_clock.ptp_local_clock_info")
}

func init() { proto.RegisterFile("ptp_local_clock_info.proto", fileDescriptor_ab8a038ddedaca8b) }

var fileDescriptor_ab8a038ddedaca8b = []byte{
	// 685 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0xfa, 0x91, 0x6e, 0x26, 0x49, 0x3f, 0xdc, 0xa8, 0x72, 0x4b, 0x05, 0x25, 0xa8, 0x52,
	0x4f, 0x29, 0xa4, 0xa5, 0x17, 0x0e, 0x1c, 0x2a, 0x0e, 0x08, 0x28, 0xd5, 0x16, 0x55, 0xe2, 0x52,
	0xcb, 0xf1, 0x3a, 0xc1, 0x62, 0xb3, 0x5e, 0xbc, 0xde, 0x88, 0xfe, 0x14, 0xfe, 0x01, 0x7f, 0x82,
	0xff, 0x86, 0x3c, 0xeb, 0xec, 0x26, 0x88, 0x54, 0x2d, 0x9c, 0xda, 0x79, 0xef, 0xf9, 0xc5, 0x9e,
	0x79, 0x93, 0xc0, 0x5e, 0x6a, 0x53, 0x16, 0x6b, 0xc1, 0x63, 0x26, 0x62, 0x2d, 0xbe, 0x32, 0x95,
	0x0c, 0x75, 0x2f, 0x35, 0xda, 0x6a, 0x72, 0x28, 0x54, 0x26, 0x34, 0x53, 0x3a, 0x63, 0xdf, 0x0d,
	0x73, 0x42, 0x9d, 0x4a, 0xd3, 0x4b, 0x6d, 0xda, 0x9b, 0x39, 0xd1, 0x7d, 0x04, 0xbb, 0x7f, 0x33,
	0x61, 0xef, 0xde, 0x7c, 0xbe, 0xea, 0x46, 0x40, 0x1d, 0x39, 0xe0, 0x23, 0x4f, 0xe5, 0x56, 0x30,
	0x3d, 0x1c, 0x66, 0xd2, 0x92, 0x43, 0x58, 0x17, 0xb9, 0x31, 0x32, 0xb1, 0x1e, 0xa1, 0xb5, 0x83,
	0xda, 0x51, 0x23, 0x6c, 0x7b, 0xf4, 0x63, 0x21, 0x7b, 0x0a, 0xad, 0x82, 0x66, 0x13, 0x1e, 0xab,
	0x88, 0x2e, 0x1d, 0xd4, 0x8e, 0x82, 0xb0, 0x59, 0x60, 0xd7, 0x0e, 0xea, 0x7e, 0x80, 0x8d, 0xe9,
	0xa7, 0xa4, 0xda, 0x58, 0xa6, 0x22, 0xb2, 0x0b, 0x81, 0xbf, 0x4b, 0x84, 0xb6, 0x2b, 0xe1, 0x1a,
	0xd6, 0x6f, 0x23, 0xf2, 0x04, 0x9a, 0xa8, 0x4a, 0xf2, 0xf1, 0x40, 0x1a, 0xf4, 0x6b, 0x87, 0xe0,
	0xa0, 0x0b, 0x44, 0xba, 0x3f, 0xeb, 0xd0, 0x9e, 0xbb, 0xf5, 0x5d, 0x6e, 0xfb, 0xd0, 0x48, 0x8d,
	0xd2, 0x46, 0xd9, 0xdb, 0x17, 0xde, 0xab, 0x02, 0x66, 0xd9, 0x3e, 0x5d, 0x9e, 0x67, 0xfb, 0xa4,
	0x03, 0xab, 0x22, 0xe6, 0x59, 0x46, 0x57, 0x90, 0x29, 0x0a, 0xb2, 0x07, 0x01, 0x17, 0x22, 0x37,
	0x5c, 0xdc, 0xd2, 0x55, 0x24, 0xca, 0x9a, 0xf4, 0x60, 0xdb, 0x37, 0x23, 0xd6, 0x23, 0x36, 0xe1,
	0x46, 0xf1, 0x44, 0x48, 0x5a, 0x47, 0xd9, 0x56, 0x41, 0xbd, 0xd7, 0xa3, 0x6b, 0x4f, 0x90, 0x67,
	0xd0, 0xce, 0xac, 0x4c, 0x33, 0x66, 0xe4, 0x58, 0x4f, 0x64, 0x44, 0xd7, 0x50, 0xd9, 0x42, 0x30,
	0x2c, 0x30, 0xd7, 0x10, 0xab, 0xc6, 0x92, 0x65, 0x3a, 0x37, 0x42, 0xd2, 0x00, 0xa7, 0x00, 0x0e,
	0xba, 0x42, 0x84, 0xdc, 0x00, 0x54, 0x73, 0xa3, 0x8d, 0x83, 0xda, 0x51, 0xb3, 0xff, 0xba, 0x77,
	0xaf, 0x78, 0xf4, 0x16, 0x8d, 0x3f, 0x6c, 0xe4, 0x56, 0xf8, 0x11, 0x1f, 0xc3, 0xf6, 0xd0, 0xc8,
	0x6f, 0xb9, 0x4c, 0xc4, 0x2d, 0xb3, 0x86, 0x0b, 0xc9, 0x07, 0xb1, 0xa4, 0x80, 0x93, 0x26, 0x25,
	0xf5, 0x69, 0xca, 0xb8, 0xe8, 0xe0, 0x8d, 0x2b, 0x6d, 0x13, 0xb5, 0x6d, 0x87, 0x56, 0xb2, 0x7d,
	0x68, 0x38, 0x20, 0x13, 0x3c, 0x96, 0xb4, 0x85, 0xcf, 0xaa, 0x00, 0x17, 0xac, 0x58, 0xf2, 0x94,
	0x65, 0x52, 0xe8, 0x24, 0xca, 0x68, 0x1b, 0x05, 0x4d, 0x87, 0x5d, 0x15, 0x10, 0x09, 0x21, 0x30,
	0x52, 0x48, 0x35, 0x91, 0x86, 0xae, 0xe3, 0xb3, 0xcf, 0x1e, 0xf8, 0x6c, 0x9f, 0xc7, 0xb0, 0xf4,
	0x21, 0x17, 0x50, 0xcf, 0x64, 0x12, 0x49, 0x43, 0x37, 0xfe, 0xcb, 0xd1, 0xbb, 0xb8, 0x10, 0xa1,
	0x8c, 0x6e, 0x62, 0x0b, 0x8a, 0x82, 0x9c, 0xc2, 0x8e, 0xd0, 0xc9, 0x50, 0x8d, 0x72, 0x23, 0x23,
	0xdf, 0xfc, 0x22, 0x6b, 0x5b, 0x98, 0x80, 0x4e, 0xc5, 0x9e, 0x3b, 0xf2, 0x1c, 0xa3, 0x77, 0x0c,
	0xdb, 0x33, 0xa7, 0xa6, 0x41, 0xa5, 0x04, 0x8f, 0x90, 0x8a, 0xba, 0xf4, 0x4c, 0xf7, 0xc7, 0x12,
	0x74, 0xa6, 0x17, 0x9b, 0x28, 0x63, 0x73, 0x1e, 0xe3, 0x05, 0xc9, 0x63, 0x80, 0x4a, 0x8e, 0x3b,
	0x13, 0x84, 0x33, 0x88, 0x1b, 0x8d, 0xd0, 0x49, 0x22, 0x85, 0x95, 0xd3, 0x95, 0xae, 0x80, 0xf9,
	0xa5, 0x5a, 0xbe, 0x73, 0xa9, 0x56, 0x16, 0x2e, 0xd5, 0xea, 0xa2, 0xa5, 0xaa, 0xdf, 0x6f, 0xa9,
	0xd6, 0x16, 0x2d, 0xd5, 0x21, 0xac, 0x17, 0x83, 0x29, 0x1b, 0x14, 0xa0, 0xb4, 0x8d, 0x68, 0xd9,
	0x9b, 0x5f, 0xbe, 0x37, 0x7f, 0x7e, 0x33, 0x92, 0x1d, 0xa8, 0x47, 0x7a, 0xcc, 0x55, 0x42, 0xfb,
	0x78, 0xce, 0x57, 0x84, 0xc1, 0x66, 0xa1, 0x4a, 0x8d, 0x4b, 0x80, 0x55, 0x32, 0xa3, 0x27, 0x98,
	0x91, 0xd3, 0x7f, 0x59, 0xb6, 0x70, 0x03, 0xff, 0x5c, 0x96, 0x66, 0xe4, 0x06, 0x5a, 0xb3, 0x43,
	0xa2, 0xa7, 0x68, 0xfe, 0xea, 0x81, 0xe6, 0xb3, 0x16, 0x61, 0xd3, 0x57, 0x97, 0x6e, 0xe8, 0x7b,
	0x10, 0x7c, 0xd1, 0x71, 0xa4, 0xdd, 0xba, 0xbc, 0xc4, 0x99, 0x96, 0x35, 0x79, 0x0e, 0x9d, 0xe9,
	0xff, 0x73, 0x71, 0x3c, 0x2b, 0xb2, 0x35, 0xe5, 0xaa, 0x30, 0x0e, 0xea, 0xf8, 0x33, 0x74, 0xf2,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x74, 0x57, 0x40, 0x30, 0xa4, 0x06, 0x00, 0x00,
}
