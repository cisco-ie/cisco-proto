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
// source: pm_interface_datarate_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_interface_data_rate_interfaces_data_rate_interface_samples_sample

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

type PmInterfaceDatarateBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceDatarateBag_KEYS) Reset()         { *m = PmInterfaceDatarateBag_KEYS{} }
func (m *PmInterfaceDatarateBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceDatarateBag_KEYS) ProtoMessage()    {}
func (*PmInterfaceDatarateBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_758cf722b94180aa, []int{0}
}

func (m *PmInterfaceDatarateBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceDatarateBag_KEYS.Unmarshal(m, b)
}
func (m *PmInterfaceDatarateBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceDatarateBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmInterfaceDatarateBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceDatarateBag_KEYS.Merge(m, src)
}
func (m *PmInterfaceDatarateBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceDatarateBag_KEYS.Size(m)
}
func (m *PmInterfaceDatarateBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceDatarateBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceDatarateBag_KEYS proto.InternalMessageInfo

func (m *PmInterfaceDatarateBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PmInterfaceDatarateBag_KEYS) GetSampleId() int32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmInterfaceDatarateBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	InputDataRate        uint32   `protobuf:"varint,51,opt,name=input_data_rate,json=inputDataRate,proto3" json:"input_data_rate,omitempty"`
	InputPacketRate      uint32   `protobuf:"varint,52,opt,name=input_packet_rate,json=inputPacketRate,proto3" json:"input_packet_rate,omitempty"`
	OutputDataRate       uint32   `protobuf:"varint,53,opt,name=output_data_rate,json=outputDataRate,proto3" json:"output_data_rate,omitempty"`
	OutputPacketRate     uint32   `protobuf:"varint,54,opt,name=output_packet_rate,json=outputPacketRate,proto3" json:"output_packet_rate,omitempty"`
	InputPeakRate        uint32   `protobuf:"varint,55,opt,name=input_peak_rate,json=inputPeakRate,proto3" json:"input_peak_rate,omitempty"`
	InputPeakPkts        uint32   `protobuf:"varint,56,opt,name=input_peak_pkts,json=inputPeakPkts,proto3" json:"input_peak_pkts,omitempty"`
	OutputPeakRate       uint32   `protobuf:"varint,57,opt,name=output_peak_rate,json=outputPeakRate,proto3" json:"output_peak_rate,omitempty"`
	OutputPeakPkts       uint32   `protobuf:"varint,58,opt,name=output_peak_pkts,json=outputPeakPkts,proto3" json:"output_peak_pkts,omitempty"`
	Bandwidth            uint32   `protobuf:"varint,59,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceDatarateBag) Reset()         { *m = PmInterfaceDatarateBag{} }
func (m *PmInterfaceDatarateBag) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceDatarateBag) ProtoMessage()    {}
func (*PmInterfaceDatarateBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_758cf722b94180aa, []int{1}
}

func (m *PmInterfaceDatarateBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceDatarateBag.Unmarshal(m, b)
}
func (m *PmInterfaceDatarateBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceDatarateBag.Marshal(b, m, deterministic)
}
func (m *PmInterfaceDatarateBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceDatarateBag.Merge(m, src)
}
func (m *PmInterfaceDatarateBag) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceDatarateBag.Size(m)
}
func (m *PmInterfaceDatarateBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceDatarateBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceDatarateBag proto.InternalMessageInfo

func (m *PmInterfaceDatarateBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetInputDataRate() uint32 {
	if m != nil {
		return m.InputDataRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetInputPacketRate() uint32 {
	if m != nil {
		return m.InputPacketRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetOutputDataRate() uint32 {
	if m != nil {
		return m.OutputDataRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetOutputPacketRate() uint32 {
	if m != nil {
		return m.OutputPacketRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetInputPeakRate() uint32 {
	if m != nil {
		return m.InputPeakRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetInputPeakPkts() uint32 {
	if m != nil {
		return m.InputPeakPkts
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetOutputPeakRate() uint32 {
	if m != nil {
		return m.OutputPeakRate
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetOutputPeakPkts() uint32 {
	if m != nil {
		return m.OutputPeakPkts
	}
	return 0
}

func (m *PmInterfaceDatarateBag) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*PmInterfaceDatarateBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.data_rate_interfaces.data_rate_interface.samples.sample.pm_interface_datarate_bag_KEYS")
	proto.RegisterType((*PmInterfaceDatarateBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.data_rate_interfaces.data_rate_interface.samples.sample.pm_interface_datarate_bag")
}

func init() { proto.RegisterFile("pm_interface_datarate_bag.proto", fileDescriptor_758cf722b94180aa) }

var fileDescriptor_758cf722b94180aa = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4b, 0x4b, 0xc3, 0x40,
	0x14, 0x46, 0x89, 0x2f, 0xcc, 0x40, 0xab, 0x66, 0x15, 0xf1, 0x15, 0x0a, 0x4a, 0x10, 0xc9, 0xc2,
	0xfa, 0x76, 0xab, 0x0b, 0x11, 0xa4, 0xa4, 0x2b, 0x57, 0xc3, 0x4d, 0x72, 0x5b, 0x87, 0x74, 0x32,
	0x43, 0xe6, 0x16, 0x15, 0xfc, 0x41, 0xfe, 0x4c, 0xc9, 0xa4, 0xa4, 0xb1, 0xb4, 0xab, 0x30, 0x27,
	0x87, 0x93, 0x2f, 0x30, 0xec, 0x44, 0x4b, 0x2e, 0x0a, 0xc2, 0x72, 0x04, 0x29, 0xf2, 0x0c, 0x08,
	0x4a, 0x20, 0xe4, 0x09, 0x8c, 0x23, 0x5d, 0x2a, 0x52, 0xde, 0x4f, 0x2a, 0x4c, 0xaa, 0xb8, 0x50,
	0x86, 0x7f, 0x95, 0x5c, 0x42, 0x01, 0x63, 0x84, 0x44, 0x4c, 0x04, 0x7d, 0x73, 0x8d, 0xe5, 0x48,
	0x8e, 0x25, 0x71, 0xa5, 0xb1, 0x8c, 0xaa, 0x13, 0xaf, 0x8e, 0x91, 0x54, 0x85, 0x20, 0x55, 0x46,
	0x4d, 0x38, 0xaa, 0xc2, 0xdc, 0x96, 0x1b, 0x66, 0x96, 0xc1, 0xc8, 0x80, 0xd4, 0x13, 0x34, 0xb3,
	0x67, 0x2f, 0x63, 0xc7, 0x2b, 0x07, 0xf2, 0xd7, 0xe7, 0xf7, 0xa1, 0x77, 0xca, 0xba, 0xf3, 0xd7,
	0x05, 0x48, 0xf4, 0x9d, 0xc0, 0x09, 0xdd, 0xb8, 0xd3, 0xd0, 0x37, 0x90, 0xe8, 0x1d, 0x30, 0xb7,
	0x4e, 0x72, 0x91, 0xf9, 0x6b, 0x81, 0x13, 0x6e, 0xc6, 0xdb, 0x35, 0x78, 0xc9, 0x7a, 0xbf, 0xeb,
	0x6c, 0x7f, 0xe5, 0x67, 0xbc, 0x23, 0xc6, 0x48, 0x48, 0xe4, 0x86, 0x40, 0x6a, 0xff, 0x32, 0x70,
	0xc2, 0x8d, 0xd8, 0xad, 0xc8, 0xb0, 0x02, 0xde, 0x19, 0xdb, 0x11, 0x85, 0x9e, 0x12, 0x6f, 0x7e,
	0xc7, 0xef, 0x07, 0x4e, 0xd8, 0xa9, 0x16, 0xe8, 0x29, 0x3d, 0x01, 0x41, 0x0c, 0x84, 0xde, 0x39,
	0xdb, 0xab, 0x3d, 0x0d, 0x69, 0x8e, 0x54, 0x9b, 0x57, 0xd6, 0xac, 0x03, 0x03, 0xcb, 0xad, 0x1b,
	0xb2, 0x5d, 0x35, 0xa5, 0xff, 0xd1, 0x6b, 0xab, 0x76, 0x6b, 0xde, 0x54, 0x2f, 0x98, 0x37, 0x33,
	0xdb, 0xd9, 0x1b, 0xeb, 0xce, 0x1a, 0xad, 0x6e, 0xb3, 0x55, 0x23, 0xe4, 0xb5, 0x7a, 0xdb, 0xda,
	0x3a, 0x40, 0xc8, 0x97, 0x78, 0x3a, 0x27, 0xe3, 0xdf, 0x2d, 0x78, 0x83, 0x9c, 0x4c, 0x6b, 0xe7,
	0x3c, 0x78, 0xdf, 0xde, 0xd9, 0x14, 0x17, 0x4c, 0x9b, 0x7c, 0x58, 0x34, 0x6d, 0xf3, 0x90, 0xb9,
	0x09, 0x14, 0xd9, 0xa7, 0xc8, 0xe8, 0xc3, 0x7f, 0xb4, 0xca, 0x1c, 0x24, 0x5b, 0xf6, 0x56, 0xf6,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x78, 0x40, 0xe4, 0xb8, 0x02, 0x00, 0x00,
}
