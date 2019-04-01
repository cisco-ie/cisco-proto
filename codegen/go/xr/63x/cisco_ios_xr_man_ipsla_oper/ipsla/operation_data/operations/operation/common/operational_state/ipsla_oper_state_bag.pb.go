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
// source: ipsla_oper_state_bag.proto

package cisco_ios_xr_man_ipsla_oper_ipsla_operation_data_operations_operation_common_operational_state

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

type IpslaOperStateBag_KEYS struct {
	OperationId          int32    `protobuf:"varint,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaOperStateBag_KEYS) Reset()         { *m = IpslaOperStateBag_KEYS{} }
func (m *IpslaOperStateBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpslaOperStateBag_KEYS) ProtoMessage()    {}
func (*IpslaOperStateBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_32675a9b4d21a2a3, []int{0}
}

func (m *IpslaOperStateBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaOperStateBag_KEYS.Unmarshal(m, b)
}
func (m *IpslaOperStateBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaOperStateBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IpslaOperStateBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaOperStateBag_KEYS.Merge(m, src)
}
func (m *IpslaOperStateBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpslaOperStateBag_KEYS.Size(m)
}
func (m *IpslaOperStateBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaOperStateBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaOperStateBag_KEYS proto.InternalMessageInfo

func (m *IpslaOperStateBag_KEYS) GetOperationId() int32 {
	if m != nil {
		return m.OperationId
	}
	return 0
}

type IpslaOperStateBag struct {
	ModificationTime         uint64   `protobuf:"varint,50,opt,name=modification_time,json=modificationTime,proto3" json:"modification_time,omitempty"`
	StartTime                uint64   `protobuf:"varint,51,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	AttemptCount             uint32   `protobuf:"varint,52,opt,name=attempt_count,json=attemptCount,proto3" json:"attempt_count,omitempty"`
	SkippedCount             uint32   `protobuf:"varint,53,opt,name=skipped_count,json=skippedCount,proto3" json:"skipped_count,omitempty"`
	LifeRemaining            uint32   `protobuf:"varint,54,opt,name=life_remaining,json=lifeRemaining,proto3" json:"life_remaining,omitempty"`
	Frequency                uint32   `protobuf:"varint,55,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Recurring                bool     `protobuf:"varint,56,opt,name=recurring,proto3" json:"recurring,omitempty"`
	OperationalState         string   `protobuf:"bytes,57,opt,name=operational_state,json=operationalState,proto3" json:"operational_state,omitempty"`
	Flags                    uint32   `protobuf:"varint,58,opt,name=flags,proto3" json:"flags,omitempty"`
	LocalPort                uint32   `protobuf:"varint,59,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	UnexpectedPackets        uint32   `protobuf:"varint,60,opt,name=unexpected_packets,json=unexpectedPackets,proto3" json:"unexpected_packets,omitempty"`
	UnexpectedControlPackets uint32   `protobuf:"varint,61,opt,name=unexpected_control_packets,json=unexpectedControlPackets,proto3" json:"unexpected_control_packets,omitempty"`
	OperationTime            uint64   `protobuf:"varint,62,opt,name=operation_time,json=operationTime,proto3" json:"operation_time,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *IpslaOperStateBag) Reset()         { *m = IpslaOperStateBag{} }
func (m *IpslaOperStateBag) String() string { return proto.CompactTextString(m) }
func (*IpslaOperStateBag) ProtoMessage()    {}
func (*IpslaOperStateBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_32675a9b4d21a2a3, []int{1}
}

func (m *IpslaOperStateBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaOperStateBag.Unmarshal(m, b)
}
func (m *IpslaOperStateBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaOperStateBag.Marshal(b, m, deterministic)
}
func (m *IpslaOperStateBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaOperStateBag.Merge(m, src)
}
func (m *IpslaOperStateBag) XXX_Size() int {
	return xxx_messageInfo_IpslaOperStateBag.Size(m)
}
func (m *IpslaOperStateBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaOperStateBag.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaOperStateBag proto.InternalMessageInfo

func (m *IpslaOperStateBag) GetModificationTime() uint64 {
	if m != nil {
		return m.ModificationTime
	}
	return 0
}

func (m *IpslaOperStateBag) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *IpslaOperStateBag) GetAttemptCount() uint32 {
	if m != nil {
		return m.AttemptCount
	}
	return 0
}

func (m *IpslaOperStateBag) GetSkippedCount() uint32 {
	if m != nil {
		return m.SkippedCount
	}
	return 0
}

func (m *IpslaOperStateBag) GetLifeRemaining() uint32 {
	if m != nil {
		return m.LifeRemaining
	}
	return 0
}

func (m *IpslaOperStateBag) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *IpslaOperStateBag) GetRecurring() bool {
	if m != nil {
		return m.Recurring
	}
	return false
}

func (m *IpslaOperStateBag) GetOperationalState() string {
	if m != nil {
		return m.OperationalState
	}
	return ""
}

func (m *IpslaOperStateBag) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *IpslaOperStateBag) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *IpslaOperStateBag) GetUnexpectedPackets() uint32 {
	if m != nil {
		return m.UnexpectedPackets
	}
	return 0
}

func (m *IpslaOperStateBag) GetUnexpectedControlPackets() uint32 {
	if m != nil {
		return m.UnexpectedControlPackets
	}
	return 0
}

func (m *IpslaOperStateBag) GetOperationTime() uint64 {
	if m != nil {
		return m.OperationTime
	}
	return 0
}

func init() {
	proto.RegisterType((*IpslaOperStateBag_KEYS)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.common.operational_state.ipsla_oper_state_bag_KEYS")
	proto.RegisterType((*IpslaOperStateBag)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.operation_data.operations.operation.common.operational_state.ipsla_oper_state_bag")
}

func init() { proto.RegisterFile("ipsla_oper_state_bag.proto", fileDescriptor_32675a9b4d21a2a3) }

var fileDescriptor_32675a9b4d21a2a3 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6b, 0xdc, 0x30,
	0x10, 0x86, 0x31, 0x24, 0xa5, 0xab, 0xc6, 0x21, 0x11, 0x39, 0xa8, 0xa1, 0x05, 0x37, 0x25, 0x60,
	0x28, 0xf5, 0xa1, 0xe9, 0x77, 0xd3, 0x5c, 0x42, 0x0f, 0xa5, 0x97, 0xc5, 0xdb, 0x4b, 0x2f, 0x15,
	0x5a, 0x59, 0x5e, 0xc4, 0x5a, 0x1f, 0x95, 0xc6, 0xb0, 0xfd, 0x01, 0xfd, 0xdf, 0x45, 0xd2, 0xee,
	0xca, 0x90, 0xbd, 0x8d, 0x9f, 0xf7, 0x19, 0x0d, 0xd6, 0x08, 0x5d, 0x4a, 0xeb, 0x07, 0x46, 0x8d,
	0x15, 0x8e, 0x7a, 0x60, 0x20, 0xe8, 0x92, 0xad, 0x1a, 0xeb, 0x0c, 0x18, 0xfc, 0x9b, 0x4b, 0xcf,
	0x0d, 0x95, 0xc6, 0xd3, 0x8d, 0xa3, 0x8a, 0x69, 0x9a, 0xe5, 0x26, 0x96, 0x4d, 0x28, 0x19, 0x48,
	0xa3, 0x69, 0xc7, 0x60, 0xf2, 0xe9, 0x73, 0xd9, 0x70, 0xa3, 0x94, 0xd1, 0x19, 0xb0, 0x21, 0x4d,
	0xba, 0xba, 0x43, 0x4f, 0x0f, 0x4d, 0xa7, 0x3f, 0xbe, 0xfd, 0x5a, 0xe0, 0x17, 0xe8, 0x24, 0x1f,
	0x2e, 0x3b, 0x52, 0x54, 0x45, 0x7d, 0xdc, 0x3e, 0xd9, 0xb3, 0xef, 0xdd, 0xd5, 0xbf, 0x23, 0x74,
	0x71, 0xe8, 0x00, 0xfc, 0x0a, 0x9d, 0x2b, 0xd3, 0xc9, 0x5e, 0xf2, 0xd4, 0x0e, 0x52, 0x09, 0xf2,
	0xa6, 0x2a, 0xea, 0xa3, 0xf6, 0x6c, 0x1a, 0xfc, 0x94, 0x4a, 0xe0, 0xe7, 0x08, 0x79, 0x60, 0x0e,
	0x92, 0x75, 0x13, 0xad, 0x59, 0x24, 0x31, 0x7e, 0x89, 0x4a, 0x06, 0x20, 0x94, 0x05, 0xca, 0xcd,
	0xa8, 0x81, 0xbc, 0xad, 0x8a, 0xba, 0x6c, 0x4f, 0xb6, 0xf0, 0x3e, 0xb0, 0x20, 0xf9, 0xb5, 0xb4,
	0x56, 0x74, 0x5b, 0xe9, 0x5d, 0x92, 0xb6, 0x30, 0x49, 0xd7, 0xe8, 0x74, 0x90, 0xbd, 0xa0, 0x4e,
	0x28, 0x26, 0xb5, 0xd4, 0x2b, 0xf2, 0x3e, 0x5a, 0x65, 0xa0, 0xed, 0x0e, 0xe2, 0x67, 0x68, 0xd6,
	0x3b, 0xf1, 0x67, 0x14, 0x9a, 0xff, 0x25, 0x1f, 0xa2, 0x91, 0x41, 0x48, 0x9d, 0xe0, 0xa3, 0x73,
	0xa1, 0xff, 0x63, 0x55, 0xd4, 0x8f, 0xdb, 0x0c, 0xc2, 0x8f, 0x3f, 0xb8, 0x66, 0xf2, 0xa9, 0x2a,
	0xea, 0x59, 0x7b, 0x36, 0x09, 0x16, 0x81, 0xe3, 0x0b, 0x74, 0xdc, 0x0f, 0x6c, 0xe5, 0xc9, 0xe7,
	0x38, 0x24, 0x7d, 0x84, 0xeb, 0x18, 0x0c, 0x67, 0x03, 0xb5, 0xc6, 0x01, 0xf9, 0x92, 0xe6, 0x47,
	0x32, 0x37, 0x0e, 0xf0, 0x6b, 0x84, 0x47, 0x2d, 0x36, 0x56, 0x70, 0x10, 0x1d, 0xb5, 0x8c, 0xaf,
	0x05, 0x78, 0x72, 0x1b, 0xb5, 0xf3, 0x9c, 0xcc, 0x53, 0x80, 0x6f, 0xd1, 0xe5, 0x44, 0xe7, 0x46,
	0x83, 0x33, 0xc3, 0xbe, 0xed, 0x6b, 0x6c, 0x23, 0xd9, 0xb8, 0x4f, 0xc2, 0xae, 0xfb, 0x1a, 0x9d,
	0xe6, 0x37, 0x10, 0xd7, 0x73, 0x17, 0xd7, 0x53, 0xee, 0x69, 0x58, 0xd1, 0xf2, 0x51, 0x7c, 0xae,
	0x37, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x48, 0x3c, 0xe3, 0x7e, 0xcc, 0x02, 0x00, 0x00,
}