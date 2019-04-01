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
// source: pm_interface_basic_stats.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_interface_basic_counter_interfaces_basic_counter_interface_samples_sample

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

type PmInterfaceBasicStats_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceBasicStats_KEYS) Reset()         { *m = PmInterfaceBasicStats_KEYS{} }
func (m *PmInterfaceBasicStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceBasicStats_KEYS) ProtoMessage()    {}
func (*PmInterfaceBasicStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7c2f10b7a8dfea6, []int{0}
}

func (m *PmInterfaceBasicStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceBasicStats_KEYS.Unmarshal(m, b)
}
func (m *PmInterfaceBasicStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceBasicStats_KEYS.Marshal(b, m, deterministic)
}
func (m *PmInterfaceBasicStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceBasicStats_KEYS.Merge(m, src)
}
func (m *PmInterfaceBasicStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceBasicStats_KEYS.Size(m)
}
func (m *PmInterfaceBasicStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceBasicStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceBasicStats_KEYS proto.InternalMessageInfo

func (m *PmInterfaceBasicStats_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PmInterfaceBasicStats_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmInterfaceBasicStats struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	InPackets            uint64   `protobuf:"varint,51,opt,name=in_packets,json=inPackets,proto3" json:"in_packets,omitempty"`
	InOctets             uint64   `protobuf:"varint,52,opt,name=in_octets,json=inOctets,proto3" json:"in_octets,omitempty"`
	OutPackets           uint64   `protobuf:"varint,53,opt,name=out_packets,json=outPackets,proto3" json:"out_packets,omitempty"`
	OutOctets            uint64   `protobuf:"varint,54,opt,name=out_octets,json=outOctets,proto3" json:"out_octets,omitempty"`
	InputTotalDrops      uint64   `protobuf:"varint,55,opt,name=input_total_drops,json=inputTotalDrops,proto3" json:"input_total_drops,omitempty"`
	InputQueueDrops      uint64   `protobuf:"varint,56,opt,name=input_queue_drops,json=inputQueueDrops,proto3" json:"input_queue_drops,omitempty"`
	InputTotalErrors     uint64   `protobuf:"varint,57,opt,name=input_total_errors,json=inputTotalErrors,proto3" json:"input_total_errors,omitempty"`
	OutputTotalDrops     uint64   `protobuf:"varint,58,opt,name=output_total_drops,json=outputTotalDrops,proto3" json:"output_total_drops,omitempty"`
	OutputQueueDrops     uint64   `protobuf:"varint,59,opt,name=output_queue_drops,json=outputQueueDrops,proto3" json:"output_queue_drops,omitempty"`
	OutputTotalErrors    uint64   `protobuf:"varint,60,opt,name=output_total_errors,json=outputTotalErrors,proto3" json:"output_total_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceBasicStats) Reset()         { *m = PmInterfaceBasicStats{} }
func (m *PmInterfaceBasicStats) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceBasicStats) ProtoMessage()    {}
func (*PmInterfaceBasicStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7c2f10b7a8dfea6, []int{1}
}

func (m *PmInterfaceBasicStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceBasicStats.Unmarshal(m, b)
}
func (m *PmInterfaceBasicStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceBasicStats.Marshal(b, m, deterministic)
}
func (m *PmInterfaceBasicStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceBasicStats.Merge(m, src)
}
func (m *PmInterfaceBasicStats) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceBasicStats.Size(m)
}
func (m *PmInterfaceBasicStats) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceBasicStats.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceBasicStats proto.InternalMessageInfo

func (m *PmInterfaceBasicStats) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetInPackets() uint64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetInOctets() uint64 {
	if m != nil {
		return m.InOctets
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetOutPackets() uint64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetOutOctets() uint64 {
	if m != nil {
		return m.OutOctets
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetInputTotalDrops() uint64 {
	if m != nil {
		return m.InputTotalDrops
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetInputQueueDrops() uint64 {
	if m != nil {
		return m.InputQueueDrops
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetInputTotalErrors() uint64 {
	if m != nil {
		return m.InputTotalErrors
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetOutputTotalDrops() uint64 {
	if m != nil {
		return m.OutputTotalDrops
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetOutputQueueDrops() uint64 {
	if m != nil {
		return m.OutputQueueDrops
	}
	return 0
}

func (m *PmInterfaceBasicStats) GetOutputTotalErrors() uint64 {
	if m != nil {
		return m.OutputTotalErrors
	}
	return 0
}

func init() {
	proto.RegisterType((*PmInterfaceBasicStats_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.interface.basic_counter_interfaces.basic_counter_interface.samples.sample.pm_interface_basic_stats_KEYS")
	proto.RegisterType((*PmInterfaceBasicStats)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.interface.basic_counter_interfaces.basic_counter_interface.samples.sample.pm_interface_basic_stats")
}

func init() { proto.RegisterFile("pm_interface_basic_stats.proto", fileDescriptor_c7c2f10b7a8dfea6) }

var fileDescriptor_c7c2f10b7a8dfea6 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0x99, 0xff, 0x5f, 0xa4, 0x8d, 0x54, 0xed, 0xb8, 0x19, 0x90, 0x6a, 0x29, 0x08, 0x45,
	0x64, 0x16, 0xd6, 0x6f, 0x5d, 0xda, 0x85, 0x08, 0x7e, 0xb4, 0x6e, 0x5c, 0x85, 0x34, 0x93, 0x96,
	0x60, 0xf3, 0x61, 0x72, 0x03, 0xfa, 0x00, 0x3e, 0x9b, 0xaf, 0x25, 0x49, 0xa6, 0x9d, 0x51, 0xe8,
	0x6a, 0xe6, 0x9e, 0xf3, 0xbb, 0xe7, 0x9e, 0x45, 0xd0, 0x9e, 0x16, 0x98, 0x4b, 0x60, 0x66, 0x4a,
	0x28, 0xc3, 0x13, 0x62, 0x39, 0xc5, 0x16, 0x08, 0xd8, 0x5c, 0x1b, 0x05, 0x2a, 0xfd, 0x4a, 0x28,
	0xb7, 0x54, 0x61, 0xae, 0x2c, 0xfe, 0x30, 0x58, 0x10, 0x49, 0x66, 0x8c, 0x4c, 0xf8, 0x9c, 0xc3,
	0x27, 0xd6, 0xcc, 0x4c, 0xc5, 0x4c, 0x00, 0x56, 0x9a, 0x99, 0xdc, 0x4f, 0xd8, 0x8f, 0xfe, 0x8f,
	0xab, 0x82, 0xd3, 0x7c, 0x99, 0x9c, 0xc7, 0x64, 0xaa, 0x9c, 0x57, 0xaa, 0x8b, 0x76, 0x95, 0x91,
	0x5b, 0x22, 0xf4, 0x9c, 0xd9, 0xf2, 0xdb, 0xa3, 0xa8, 0xb3, 0xaa, 0x29, 0xbe, 0x1f, 0xbe, 0x8e,
	0xd3, 0x03, 0xb4, 0x59, 0xb9, 0x92, 0x08, 0x96, 0x25, 0xdd, 0xa4, 0xdf, 0x1c, 0xb5, 0x96, 0xea,
	0x03, 0x11, 0x2c, 0xdd, 0x45, 0xcd, 0x98, 0x88, 0x79, 0x91, 0xfd, 0xeb, 0x26, 0xfd, 0xd6, 0xa8,
	0x11, 0x85, 0xbb, 0xa2, 0xf7, 0xfd, 0x1f, 0x65, 0xab, 0xae, 0xa4, 0x1d, 0x84, 0x80, 0x0b, 0xe6,
	0x27, 0xa1, 0xb3, 0xe3, 0x6e, 0xd2, 0x5f, 0x1b, 0x35, 0xbd, 0x32, 0xf6, 0x82, 0xb7, 0xb9, 0xc4,
	0x9a, 0xd0, 0x37, 0x06, 0x36, 0x1b, 0x44, 0x9b, 0xcb, 0xa7, 0x28, 0xf8, 0xbb, 0x5c, 0x62, 0x45,
	0xc1, 0xbb, 0x27, 0xc1, 0x6d, 0x70, 0xf9, 0x18, 0xe6, 0x74, 0x1f, 0x6d, 0x28, 0x07, 0xcb, 0xe5,
	0xd3, 0x60, 0x23, 0xe5, 0x60, 0xb1, 0xdd, 0x41, 0x7e, 0x5a, 0xac, 0x9f, 0xc5, 0x70, 0xe5, 0xa0,
	0xdc, 0x3f, 0x44, 0x6d, 0x2e, 0xb5, 0x03, 0x0c, 0x0a, 0xc8, 0x1c, 0x17, 0x46, 0x69, 0x9b, 0x9d,
	0x07, 0x6a, 0x2b, 0x18, 0x2f, 0x5e, 0xbf, 0xf5, 0x72, 0xc5, 0xbe, 0x3b, 0xe6, 0x58, 0xc9, 0x5e,
	0xd4, 0xd8, 0x67, 0xaf, 0x47, 0xf6, 0x08, 0xa5, 0xf5, 0x5c, 0x66, 0x8c, 0x32, 0x36, 0xbb, 0x0c,
	0xf0, 0x76, 0x15, 0x3c, 0x0c, 0xba, 0xa7, 0x95, 0x83, 0xbf, 0x35, 0xae, 0x22, 0x1d, 0x9d, 0x5a,
	0x8f, 0x8a, 0xae, 0x17, 0xb9, 0xae, 0xd3, 0xb5, 0x26, 0x39, 0xda, 0xf9, 0x95, 0x5d, 0x56, 0xb9,
	0x09, 0x78, 0xbb, 0x16, 0x1e, 0xbb, 0x4c, 0xd6, 0xc3, 0xeb, 0x1d, 0xfc, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x24, 0xa7, 0x05, 0x0b, 0xdf, 0x02, 0x00, 0x00,
}