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
// source: pm_interface_generic_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_interface_generic_counter_interfaces_generic_counter_interface_samples_sample

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

type PmInterfaceGenericBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceGenericBag_KEYS) Reset()         { *m = PmInterfaceGenericBag_KEYS{} }
func (m *PmInterfaceGenericBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceGenericBag_KEYS) ProtoMessage()    {}
func (*PmInterfaceGenericBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99bf44af982d4e7, []int{0}
}

func (m *PmInterfaceGenericBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceGenericBag_KEYS.Unmarshal(m, b)
}
func (m *PmInterfaceGenericBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceGenericBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmInterfaceGenericBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceGenericBag_KEYS.Merge(m, src)
}
func (m *PmInterfaceGenericBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceGenericBag_KEYS.Size(m)
}
func (m *PmInterfaceGenericBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceGenericBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceGenericBag_KEYS proto.InternalMessageInfo

func (m *PmInterfaceGenericBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PmInterfaceGenericBag_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmInterfaceGenericBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	InPackets            uint64   `protobuf:"varint,51,opt,name=in_packets,json=inPackets,proto3" json:"in_packets,omitempty"`
	InOctets             uint64   `protobuf:"varint,52,opt,name=in_octets,json=inOctets,proto3" json:"in_octets,omitempty"`
	OutPackets           uint64   `protobuf:"varint,53,opt,name=out_packets,json=outPackets,proto3" json:"out_packets,omitempty"`
	OutOctets            uint64   `protobuf:"varint,54,opt,name=out_octets,json=outOctets,proto3" json:"out_octets,omitempty"`
	InUcastPkts          uint64   `protobuf:"varint,55,opt,name=in_ucast_pkts,json=inUcastPkts,proto3" json:"in_ucast_pkts,omitempty"`
	InMulticastPkts      uint64   `protobuf:"varint,56,opt,name=in_multicast_pkts,json=inMulticastPkts,proto3" json:"in_multicast_pkts,omitempty"`
	InBroadcastPkts      uint64   `protobuf:"varint,57,opt,name=in_broadcast_pkts,json=inBroadcastPkts,proto3" json:"in_broadcast_pkts,omitempty"`
	OutUcastPkts         uint64   `protobuf:"varint,58,opt,name=out_ucast_pkts,json=outUcastPkts,proto3" json:"out_ucast_pkts,omitempty"`
	OutMulticastPkts     uint64   `protobuf:"varint,59,opt,name=out_multicast_pkts,json=outMulticastPkts,proto3" json:"out_multicast_pkts,omitempty"`
	OutBroadcastPkts     uint64   `protobuf:"varint,60,opt,name=out_broadcast_pkts,json=outBroadcastPkts,proto3" json:"out_broadcast_pkts,omitempty"`
	OutputTotalDrops     uint32   `protobuf:"varint,61,opt,name=output_total_drops,json=outputTotalDrops,proto3" json:"output_total_drops,omitempty"`
	InputTotalDrops      uint32   `protobuf:"varint,62,opt,name=input_total_drops,json=inputTotalDrops,proto3" json:"input_total_drops,omitempty"`
	InputQueueDrops      uint32   `protobuf:"varint,63,opt,name=input_queue_drops,json=inputQueueDrops,proto3" json:"input_queue_drops,omitempty"`
	InputUnknownProto    uint32   `protobuf:"varint,64,opt,name=input_unknown_proto,json=inputUnknownProto,proto3" json:"input_unknown_proto,omitempty"`
	OutputTotalErrors    uint32   `protobuf:"varint,65,opt,name=output_total_errors,json=outputTotalErrors,proto3" json:"output_total_errors,omitempty"`
	OutputUnderrun       uint32   `protobuf:"varint,66,opt,name=output_underrun,json=outputUnderrun,proto3" json:"output_underrun,omitempty"`
	InputTotalErrors     uint32   `protobuf:"varint,67,opt,name=input_total_errors,json=inputTotalErrors,proto3" json:"input_total_errors,omitempty"`
	InputCrc             uint32   `protobuf:"varint,68,opt,name=input_crc,json=inputCrc,proto3" json:"input_crc,omitempty"`
	InputOverrun         uint32   `protobuf:"varint,69,opt,name=input_overrun,json=inputOverrun,proto3" json:"input_overrun,omitempty"`
	InputFrame           uint32   `protobuf:"varint,70,opt,name=input_frame,json=inputFrame,proto3" json:"input_frame,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmInterfaceGenericBag) Reset()         { *m = PmInterfaceGenericBag{} }
func (m *PmInterfaceGenericBag) String() string { return proto.CompactTextString(m) }
func (*PmInterfaceGenericBag) ProtoMessage()    {}
func (*PmInterfaceGenericBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c99bf44af982d4e7, []int{1}
}

func (m *PmInterfaceGenericBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmInterfaceGenericBag.Unmarshal(m, b)
}
func (m *PmInterfaceGenericBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmInterfaceGenericBag.Marshal(b, m, deterministic)
}
func (m *PmInterfaceGenericBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmInterfaceGenericBag.Merge(m, src)
}
func (m *PmInterfaceGenericBag) XXX_Size() int {
	return xxx_messageInfo_PmInterfaceGenericBag.Size(m)
}
func (m *PmInterfaceGenericBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmInterfaceGenericBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmInterfaceGenericBag proto.InternalMessageInfo

func (m *PmInterfaceGenericBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInPackets() uint64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInOctets() uint64 {
	if m != nil {
		return m.InOctets
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutPackets() uint64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutOctets() uint64 {
	if m != nil {
		return m.OutOctets
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInUcastPkts() uint64 {
	if m != nil {
		return m.InUcastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInMulticastPkts() uint64 {
	if m != nil {
		return m.InMulticastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInBroadcastPkts() uint64 {
	if m != nil {
		return m.InBroadcastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutUcastPkts() uint64 {
	if m != nil {
		return m.OutUcastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutMulticastPkts() uint64 {
	if m != nil {
		return m.OutMulticastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutBroadcastPkts() uint64 {
	if m != nil {
		return m.OutBroadcastPkts
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutputTotalDrops() uint32 {
	if m != nil {
		return m.OutputTotalDrops
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputTotalDrops() uint32 {
	if m != nil {
		return m.InputTotalDrops
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputQueueDrops() uint32 {
	if m != nil {
		return m.InputQueueDrops
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputUnknownProto() uint32 {
	if m != nil {
		return m.InputUnknownProto
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutputTotalErrors() uint32 {
	if m != nil {
		return m.OutputTotalErrors
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetOutputUnderrun() uint32 {
	if m != nil {
		return m.OutputUnderrun
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputTotalErrors() uint32 {
	if m != nil {
		return m.InputTotalErrors
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputCrc() uint32 {
	if m != nil {
		return m.InputCrc
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputOverrun() uint32 {
	if m != nil {
		return m.InputOverrun
	}
	return 0
}

func (m *PmInterfaceGenericBag) GetInputFrame() uint32 {
	if m != nil {
		return m.InputFrame
	}
	return 0
}

func init() {
	proto.RegisterType((*PmInterfaceGenericBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.generic_counter_interfaces.generic_counter_interface.samples.sample.pm_interface_generic_bag_KEYS")
	proto.RegisterType((*PmInterfaceGenericBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.interface.generic_counter_interfaces.generic_counter_interface.samples.sample.pm_interface_generic_bag")
}

func init() { proto.RegisterFile("pm_interface_generic_bag.proto", fileDescriptor_c99bf44af982d4e7) }

var fileDescriptor_c99bf44af982d4e7 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xc7, 0x95, 0x4f, 0x9f, 0xaa, 0xd6, 0x6d, 0x93, 0x32, 0xdd, 0x58, 0x42, 0x81, 0x28, 0x80,
	0x88, 0x10, 0x9a, 0x05, 0xe5, 0x7e, 0xa7, 0x6d, 0x2a, 0x21, 0x04, 0x0d, 0x29, 0x59, 0xb0, 0xb2,
	0x1c, 0xc7, 0x89, 0xac, 0xc4, 0xf6, 0xe0, 0x0b, 0x97, 0x37, 0xe0, 0xfd, 0x78, 0x21, 0x74, 0x8e,
	0x67, 0x32, 0x49, 0x44, 0x56, 0x91, 0x7f, 0xff, 0x9f, 0xcf, 0x39, 0xe3, 0xd8, 0xe4, 0x46, 0xa1,
	0x99, 0x32, 0x41, 0xba, 0x29, 0x17, 0x92, 0xcd, 0xa4, 0x91, 0x4e, 0x09, 0x36, 0xe6, 0xb3, 0xbc,
	0x70, 0x36, 0xd8, 0xec, 0x77, 0x43, 0x28, 0x2f, 0x2c, 0x53, 0xd6, 0xb3, 0x9f, 0x8e, 0x69, 0x6e,
	0xf8, 0x4c, 0xf2, 0xb1, 0x5a, 0xa8, 0xf0, 0x8b, 0x15, 0xd2, 0x4d, 0xf5, 0x4c, 0x07, 0x66, 0x0b,
	0xe9, 0x72, 0x58, 0x31, 0x58, 0xe6, 0xda, 0x1a, 0x15, 0xac, 0xcb, 0x97, 0x85, 0xf3, 0xaa, 0xb0,
	0xb0, 0x11, 0x58, 0xdd, 0xd2, 0x6f, 0x8f, 0x72, 0xcf, 0x75, 0xb1, 0x90, 0xbe, 0xfc, 0xed, 0x0a,
	0xd2, 0xde, 0x36, 0x2c, 0xfb, 0xd0, 0xff, 0x7a, 0x95, 0xdd, 0x21, 0xcd, 0x3a, 0x35, 0x5c, 0x4b,
	0xda, 0xe8, 0x34, 0x7a, 0x7b, 0xc3, 0xc3, 0x25, 0xfd, 0xc4, 0xb5, 0xcc, 0xae, 0x93, 0xbd, 0x54,
	0x91, 0xa9, 0x09, 0xfd, 0xaf, 0xd3, 0xe8, 0x1d, 0x0e, 0x77, 0x13, 0x78, 0x3f, 0xe9, 0xfe, 0xd9,
	0x21, 0x74, 0x5b, 0x97, 0xac, 0x4d, 0x48, 0x50, 0x5a, 0x32, 0x1f, 0xb8, 0x2e, 0xe8, 0x83, 0x4e,
	0xa3, 0xf7, 0xff, 0x70, 0x0f, 0xc8, 0x15, 0x00, 0x88, 0x95, 0x61, 0x05, 0x17, 0x73, 0x19, 0x3c,
	0x3d, 0x49, 0xb1, 0x32, 0x83, 0x04, 0xa0, 0xaf, 0x32, 0xcc, 0x8a, 0x00, 0xe9, 0x43, 0x4c, 0x77,
	0x95, 0xb9, 0xc4, 0x75, 0x76, 0x93, 0xec, 0xdb, 0x18, 0x96, 0x9b, 0x1f, 0x61, 0x4c, 0x6c, 0x0c,
	0xd5, 0xee, 0x36, 0x81, 0x55, 0xb5, 0xfd, 0x71, 0x2a, 0x6e, 0x63, 0x28, 0xf7, 0x77, 0xc9, 0xa1,
	0x32, 0x2c, 0x0a, 0xee, 0x03, 0x2b, 0xe6, 0xc1, 0xd3, 0x27, 0x68, 0xec, 0x2b, 0x33, 0x02, 0x36,
	0x98, 0x07, 0x9f, 0xdd, 0x23, 0xd7, 0x94, 0x61, 0x3a, 0x2e, 0x82, 0xaa, 0xbd, 0xa7, 0xe8, 0xb5,
	0x94, 0xf9, 0x58, 0xf1, 0x15, 0x77, 0xec, 0x2c, 0x9f, 0xd4, 0xee, 0xb3, 0xca, 0x3d, 0xad, 0x38,
	0xba, 0xb7, 0x49, 0x13, 0x46, 0x5b, 0x69, 0xfe, 0x1c, 0xc5, 0x03, 0x1b, 0x43, 0xdd, 0xfd, 0x3e,
	0xc9, 0xc0, 0xda, 0x68, 0xff, 0x02, 0xcd, 0x23, 0x1b, 0xc3, 0x7a, 0xff, 0xd2, 0xde, 0x18, 0xe0,
	0xe5, 0xd2, 0x5e, 0x9f, 0x20, 0xd9, 0x45, 0x0c, 0x2c, 0xd8, 0xc0, 0x17, 0x6c, 0xe2, 0x6c, 0xe1,
	0xe9, 0x2b, 0xfc, 0x6f, 0x8f, 0x52, 0xf2, 0x05, 0x82, 0x73, 0xe0, 0xe9, 0xdb, 0x36, 0xe5, 0xd7,
	0x28, 0xb7, 0x30, 0xf8, 0x97, 0xfb, 0x2d, 0xca, 0x28, 0x4b, 0xf7, 0xcd, 0x8a, 0xfb, 0x19, 0x78,
	0x72, 0x73, 0x72, 0x9c, 0xdc, 0x68, 0xe6, 0xc6, 0xfe, 0x30, 0x0c, 0x9f, 0x10, 0x7d, 0x8b, 0x76,
	0x2a, 0x33, 0x4a, 0xc9, 0x00, 0xdf, 0x56, 0x4e, 0x8e, 0xd7, 0xa6, 0x96, 0xce, 0x59, 0xe7, 0xe9,
	0xbb, 0xe4, 0xaf, 0x8c, 0xdd, 0xc7, 0x20, 0xbb, 0x4b, 0x5a, 0xa5, 0x1f, 0xcd, 0x44, 0x3a, 0x17,
	0x0d, 0x3d, 0x45, 0xb7, 0x99, 0xf0, 0xa8, 0xa4, 0x70, 0x1c, 0xab, 0x1f, 0x58, 0xd6, 0x3d, 0x4b,
	0xc7, 0x51, 0x7f, 0x61, 0x59, 0x16, 0xef, 0x25, 0xd8, 0xc2, 0x09, 0x7a, 0x9e, 0xde, 0x03, 0x82,
	0x33, 0x27, 0xb2, 0x5b, 0x70, 0xaf, 0x20, 0xb4, 0xdf, 0x53, 0xc7, 0x3e, 0x0a, 0x07, 0x08, 0x2f,
	0x13, 0x83, 0xcb, 0x9b, 0xa4, 0xa9, 0x83, 0x57, 0x77, 0x81, 0x0a, 0x41, 0x74, 0x01, 0x64, 0xbc,
	0x83, 0x27, 0x71, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x10, 0xaf, 0x90, 0xef, 0x6e, 0x04, 0x00,
	0x00,
}
