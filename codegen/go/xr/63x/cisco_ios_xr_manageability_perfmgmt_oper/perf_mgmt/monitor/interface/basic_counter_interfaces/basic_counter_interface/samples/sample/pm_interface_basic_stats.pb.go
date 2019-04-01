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

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_interface_basic_counter_interfaces_basic_counter_interface_samples_sample

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
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
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

func (m *PmInterfaceBasicStats_KEYS) GetSampleId() int32 {
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
	proto.RegisterType((*PmInterfaceBasicStats_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.basic_counter_interfaces.basic_counter_interface.samples.sample.pm_interface_basic_stats_KEYS")
	proto.RegisterType((*PmInterfaceBasicStats)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.basic_counter_interfaces.basic_counter_interface.samples.sample.pm_interface_basic_stats")
}

func init() { proto.RegisterFile("pm_interface_basic_stats.proto", fileDescriptor_c7c2f10b7a8dfea6) }

var fileDescriptor_c7c2f10b7a8dfea6 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcb, 0x4b, 0x2b, 0x31,
	0x18, 0xc5, 0x99, 0xfb, 0xa2, 0xcd, 0xe5, 0x3e, 0x3a, 0x77, 0x33, 0x70, 0xe9, 0xbd, 0xa5, 0x20,
	0x14, 0x91, 0x59, 0x58, 0xdf, 0xba, 0xb4, 0x0b, 0x11, 0x7c, 0xb4, 0x6e, 0x5c, 0x85, 0x34, 0x4d,
	0x4b, 0xb0, 0xc9, 0x17, 0x93, 0x2f, 0xa0, 0x7b, 0xff, 0x37, 0xff, 0x2d, 0x49, 0x32, 0xed, 0x8c,
	0x42, 0x57, 0xc3, 0x39, 0xe7, 0xf7, 0x9d, 0x9c, 0xc5, 0x90, 0x7f, 0x46, 0x51, 0xa9, 0x51, 0xd8,
	0x39, 0xe3, 0x82, 0x4e, 0x99, 0x93, 0x9c, 0x3a, 0x64, 0xe8, 0x4a, 0x63, 0x01, 0x21, 0x7f, 0xc9,
	0xb8, 0x74, 0x1c, 0xa8, 0x04, 0x47, 0x9f, 0x2c, 0x55, 0x4c, 0xb3, 0x85, 0x60, 0x53, 0xb9, 0x94,
	0xf8, 0x4c, 0x8d, 0xb0, 0x73, 0xb5, 0x50, 0x48, 0xc1, 0x08, 0x5b, 0x06, 0x45, 0x83, 0x2c, 0x15,
	0x68, 0x89, 0x60, 0xcb, 0x75, 0x71, 0x99, 0x8a, 0x39, 0xf8, 0xe0, 0xd4, 0x0f, 0xba, 0x4d, 0x41,
	0xe9, 0x98, 0x32, 0x4b, 0xe1, 0xaa, 0x6f, 0x9f, 0x93, 0xee, 0xa6, 0xa1, 0xf4, 0x72, 0x74, 0x3f,
	0xc9, 0xb7, 0xc8, 0xcf, 0x3a, 0xd5, 0x4c, 0x89, 0x22, 0xeb, 0x65, 0x83, 0xf6, 0xf8, 0xc7, 0xda,
	0xbd, 0x62, 0x4a, 0xe4, 0x7f, 0x49, 0x3b, 0x35, 0x52, 0x39, 0x2b, 0x3e, 0xf5, 0xb2, 0xc1, 0xd7,
	0x71, 0x2b, 0x19, 0x17, 0xb3, 0xfe, 0xeb, 0x67, 0x52, 0x6c, 0x7a, 0x25, 0xef, 0x12, 0x82, 0x52,
	0x89, 0xa0, 0x94, 0x29, 0x76, 0x7b, 0xd9, 0xe0, 0xcb, 0xb8, 0x1d, 0x9c, 0x49, 0x30, 0x42, 0x2c,
	0x35, 0x35, 0x8c, 0x3f, 0x08, 0x74, 0xc5, 0x30, 0xc5, 0x52, 0xdf, 0x24, 0x23, 0xbc, 0x2b, 0x35,
	0x05, 0x8e, 0x21, 0xdd, 0x8b, 0x69, 0x4b, 0xea, 0xeb, 0xa8, 0xf3, 0xff, 0xe4, 0x3b, 0x78, 0x5c,
	0x1f, 0xef, 0xc7, 0x98, 0x80, 0xc7, 0xd5, 0x75, 0x97, 0x04, 0xb5, 0x3a, 0x3f, 0x48, 0xe5, 0xe0,
	0xb1, 0xba, 0xdf, 0x26, 0x1d, 0xa9, 0x8d, 0x47, 0x8a, 0x80, 0x6c, 0x49, 0x67, 0x16, 0x8c, 0x2b,
	0x0e, 0x23, 0xf5, 0x2b, 0x06, 0x77, 0xc1, 0x3f, 0x0f, 0x76, 0xcd, 0x3e, 0x7a, 0xe1, 0x45, 0xc5,
	0x1e, 0x35, 0xd8, 0xdb, 0xe0, 0x27, 0x76, 0x87, 0xe4, 0xcd, 0x5e, 0x61, 0x2d, 0x58, 0x57, 0x1c,
	0x47, 0xf8, 0x77, 0x5d, 0x3c, 0x8a, 0x7e, 0xa0, 0xc1, 0xe3, 0xc7, 0x19, 0x27, 0x89, 0x4e, 0x49,
	0x63, 0x47, 0x4d, 0x37, 0x87, 0x9c, 0x36, 0xe9, 0xc6, 0x92, 0x92, 0xfc, 0x79, 0xd7, 0x5d, 0x4d,
	0x39, 0x8b, 0x78, 0xa7, 0x51, 0x9e, 0xb6, 0x4c, 0xbf, 0xc5, 0x9f, 0x77, 0xf8, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xbd, 0x38, 0x4b, 0xde, 0x02, 0x00, 0x00,
}