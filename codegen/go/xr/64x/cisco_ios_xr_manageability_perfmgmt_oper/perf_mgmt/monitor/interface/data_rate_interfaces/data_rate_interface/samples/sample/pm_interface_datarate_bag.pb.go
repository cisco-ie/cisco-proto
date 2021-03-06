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
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
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

func (m *PmInterfaceDatarateBag_KEYS) GetSampleId() uint32 {
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
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x46, 0xc9, 0xfb, 0x8a, 0x98, 0x81, 0x56, 0xcd, 0x2a, 0xe2, 0x57, 0x29, 0x28, 0x41, 0x24,
	0x0b, 0xeb, 0xb7, 0x5b, 0x5d, 0x88, 0x20, 0x25, 0x5d, 0xb9, 0x1a, 0x6e, 0x92, 0xdb, 0x3a, 0xa4,
	0x93, 0x19, 0x32, 0xb7, 0xa8, 0xe0, 0x0f, 0xf2, 0x67, 0x4a, 0x26, 0x65, 0x1a, 0x4b, 0xbb, 0x0a,
	0x73, 0x72, 0x38, 0x79, 0x02, 0xc3, 0x8e, 0xb5, 0xe4, 0xa2, 0x24, 0xac, 0xc6, 0x90, 0x21, 0xcf,
	0x81, 0xa0, 0x02, 0x42, 0x9e, 0xc2, 0x24, 0xd6, 0x95, 0x22, 0x15, 0x7c, 0x67, 0xc2, 0x64, 0x8a,
	0x0b, 0x65, 0xf8, 0x67, 0xc5, 0x25, 0x94, 0x30, 0x41, 0x48, 0xc5, 0x54, 0xd0, 0x17, 0xd7, 0x58,
	0x8d, 0xe5, 0x44, 0x12, 0x57, 0x1a, 0xab, 0xb8, 0x3e, 0xf1, 0xfa, 0x18, 0x4b, 0x55, 0x0a, 0x52,
	0x55, 0xec, 0xc2, 0x71, 0x1d, 0xe6, 0xb6, 0xec, 0x98, 0x59, 0x05, 0x63, 0x03, 0x52, 0x4f, 0xd1,
	0xcc, 0x9f, 0xfd, 0x9c, 0x1d, 0xad, 0x1d, 0xc8, 0x5f, 0x9e, 0xde, 0x46, 0xc1, 0x09, 0xeb, 0x2e,
	0x5e, 0x97, 0x20, 0x31, 0xf4, 0x7a, 0x5e, 0xe4, 0x27, 0x1d, 0x47, 0x5f, 0x41, 0x62, 0xb0, 0xcf,
	0xfc, 0x26, 0xc9, 0x45, 0x1e, 0xfe, 0xeb, 0x79, 0x51, 0x27, 0xd9, 0x6a, 0xc0, 0x73, 0xde, 0xff,
	0xf9, 0xcf, 0xf6, 0xd6, 0x7e, 0x26, 0x38, 0x64, 0x8c, 0x84, 0x44, 0x6e, 0x08, 0xa4, 0x0e, 0x2f,
	0x7a, 0x5e, 0xb4, 0x91, 0xf8, 0x35, 0x19, 0xd5, 0x20, 0x38, 0x65, 0xdb, 0xa2, 0xd4, 0x33, 0xe2,
	0xee, 0x77, 0xc2, 0x81, 0xed, 0x77, 0x2c, 0x7e, 0x04, 0x82, 0x04, 0x08, 0x83, 0x33, 0xb6, 0xdb,
	0x78, 0x1a, 0xb2, 0x02, 0xa9, 0x31, 0x2f, 0xad, 0xd9, 0x04, 0x86, 0x96, 0x5b, 0x37, 0x62, 0x3b,
	0x6a, 0x46, 0x7f, 0xa3, 0x57, 0x56, 0xed, 0x36, 0xdc, 0x55, 0xcf, 0x59, 0x30, 0x37, 0xdb, 0xd9,
	0x6b, 0xeb, 0xce, 0x1b, 0xad, 0xae, 0xdb, 0xaa, 0x11, 0x8a, 0x46, 0xbd, 0x69, 0x6d, 0x1d, 0x22,
	0x14, 0x2b, 0x3c, 0x5d, 0x90, 0x09, 0x6f, 0x97, 0xbc, 0x61, 0x41, 0xa6, 0xb5, 0x73, 0x11, 0xbc,
	0x6b, 0xef, 0x74, 0xc5, 0x25, 0xd3, 0x26, 0xef, 0x97, 0x4d, 0xdb, 0x3c, 0x60, 0x7e, 0x0a, 0x65,
	0xfe, 0x21, 0x72, 0x7a, 0x0f, 0x1f, 0xac, 0xb2, 0x00, 0xe9, 0xa6, 0xbd, 0x95, 0x83, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x72, 0xb8, 0xe4, 0x7a, 0xb8, 0x02, 0x00, 0x00,
}
