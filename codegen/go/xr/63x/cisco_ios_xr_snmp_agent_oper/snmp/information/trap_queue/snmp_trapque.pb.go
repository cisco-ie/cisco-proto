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
// source: snmp_trapque.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_trap_queue

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

type SnmpTrapque_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpTrapque_KEYS) Reset()         { *m = SnmpTrapque_KEYS{} }
func (m *SnmpTrapque_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapque_KEYS) ProtoMessage()    {}
func (*SnmpTrapque_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_20735a7a5bda6daf, []int{0}
}

func (m *SnmpTrapque_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapque_KEYS.Unmarshal(m, b)
}
func (m *SnmpTrapque_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapque_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpTrapque_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapque_KEYS.Merge(m, src)
}
func (m *SnmpTrapque_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapque_KEYS.Size(m)
}
func (m *SnmpTrapque_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapque_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapque_KEYS proto.InternalMessageInfo

type SnmpQueue_ struct {
	Min                  uint32   `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Avg                  uint32   `protobuf:"varint,2,opt,name=avg,proto3" json:"avg,omitempty"`
	Max                  uint32   `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpQueue_) Reset()         { *m = SnmpQueue_{} }
func (m *SnmpQueue_) String() string { return proto.CompactTextString(m) }
func (*SnmpQueue_) ProtoMessage()    {}
func (*SnmpQueue_) Descriptor() ([]byte, []int) {
	return fileDescriptor_20735a7a5bda6daf, []int{1}
}

func (m *SnmpQueue_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpQueue_.Unmarshal(m, b)
}
func (m *SnmpQueue_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpQueue_.Marshal(b, m, deterministic)
}
func (m *SnmpQueue_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpQueue_.Merge(m, src)
}
func (m *SnmpQueue_) XXX_Size() int {
	return xxx_messageInfo_SnmpQueue_.Size(m)
}
func (m *SnmpQueue_) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpQueue_.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpQueue_ proto.InternalMessageInfo

func (m *SnmpQueue_) GetMin() uint32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *SnmpQueue_) GetAvg() uint32 {
	if m != nil {
		return m.Avg
	}
	return 0
}

func (m *SnmpQueue_) GetMax() uint32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type SnmpTrapque struct {
	TrapMin              uint32        `protobuf:"varint,50,opt,name=trap_min,json=trapMin,proto3" json:"trap_min,omitempty"`
	TrapAvg              uint32        `protobuf:"varint,51,opt,name=trap_avg,json=trapAvg,proto3" json:"trap_avg,omitempty"`
	TrapMax              uint32        `protobuf:"varint,52,opt,name=trap_max,json=trapMax,proto3" json:"trap_max,omitempty"`
	TrapQ                []*SnmpQueue_ `protobuf:"bytes,53,rep,name=trap_q,json=trapQ,proto3" json:"trap_q,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SnmpTrapque) Reset()         { *m = SnmpTrapque{} }
func (m *SnmpTrapque) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapque) ProtoMessage()    {}
func (*SnmpTrapque) Descriptor() ([]byte, []int) {
	return fileDescriptor_20735a7a5bda6daf, []int{2}
}

func (m *SnmpTrapque) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapque.Unmarshal(m, b)
}
func (m *SnmpTrapque) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapque.Marshal(b, m, deterministic)
}
func (m *SnmpTrapque) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapque.Merge(m, src)
}
func (m *SnmpTrapque) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapque.Size(m)
}
func (m *SnmpTrapque) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapque.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapque proto.InternalMessageInfo

func (m *SnmpTrapque) GetTrapMin() uint32 {
	if m != nil {
		return m.TrapMin
	}
	return 0
}

func (m *SnmpTrapque) GetTrapAvg() uint32 {
	if m != nil {
		return m.TrapAvg
	}
	return 0
}

func (m *SnmpTrapque) GetTrapMax() uint32 {
	if m != nil {
		return m.TrapMax
	}
	return 0
}

func (m *SnmpTrapque) GetTrapQ() []*SnmpQueue_ {
	if m != nil {
		return m.TrapQ
	}
	return nil
}

func init() {
	proto.RegisterType((*SnmpTrapque_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_queue.snmp_trapque_KEYS")
	proto.RegisterType((*SnmpQueue_)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_queue.snmp_queue_")
	proto.RegisterType((*SnmpTrapque)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_queue.snmp_trapque")
}

func init() { proto.RegisterFile("snmp_trapque.proto", fileDescriptor_20735a7a5bda6daf) }

var fileDescriptor_20735a7a5bda6daf = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0xce, 0xcb, 0x2d,
	0x88, 0x2f, 0x29, 0x4a, 0x2c, 0x28, 0x2c, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2,
	0x48, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x07, 0x2b, 0x48,
	0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xf5, 0x32, 0xf3, 0xd2,
	0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x40, 0x9a, 0xe3, 0x0b, 0x4b, 0x53, 0x4b,
	0x53, 0x95, 0x84, 0xb9, 0x04, 0x91, 0xcd, 0x8b, 0xf7, 0x76, 0x8d, 0x0c, 0x56, 0x72, 0xe6, 0xe2,
	0x06, 0x0b, 0x82, 0x95, 0xc4, 0x0b, 0x09, 0x70, 0x31, 0xe7, 0x66, 0xe6, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0xf0, 0x06, 0x81, 0x98, 0x20, 0x91, 0xc4, 0xb2, 0x74, 0x09, 0x26, 0x88, 0x48, 0x62, 0x59,
	0x3a, 0x58, 0x4d, 0x62, 0x85, 0x04, 0x33, 0x54, 0x4d, 0x62, 0x85, 0xd2, 0x5e, 0x46, 0x2e, 0x1e,
	0x64, 0xa3, 0x85, 0x24, 0xb9, 0x38, 0xc0, 0x16, 0x83, 0xcc, 0x32, 0x02, 0xab, 0x63, 0x07, 0xf1,
	0x7d, 0x33, 0xf3, 0xe0, 0x52, 0x20, 0x43, 0x8d, 0x11, 0x52, 0x8e, 0x65, 0xe9, 0x08, 0x5d, 0x89,
	0x15, 0x12, 0x26, 0x48, 0xba, 0x12, 0x2b, 0x84, 0x62, 0xb8, 0xd8, 0x20, 0x3e, 0x91, 0x30, 0x55,
	0x60, 0xd6, 0xe0, 0x36, 0x72, 0xd5, 0x23, 0x37, 0x18, 0xf4, 0x90, 0xbc, 0x1b, 0xc4, 0x0a, 0x12,
	0x0f, 0x4c, 0x62, 0x03, 0x07, 0xad, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x86, 0xd9, 0x53, 0x5a,
	0x70, 0x01, 0x00, 0x00,
}
