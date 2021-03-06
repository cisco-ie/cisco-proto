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
// source: pm_bgp_nbr_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_bgp_bgp_neighbors_bgp_neighbor_samples_sample

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

type PmBgpNbrBag_KEYS struct {
	IpAddress            string   `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmBgpNbrBag_KEYS) Reset()         { *m = PmBgpNbrBag_KEYS{} }
func (m *PmBgpNbrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmBgpNbrBag_KEYS) ProtoMessage()    {}
func (*PmBgpNbrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e88bee685a3a839, []int{0}
}

func (m *PmBgpNbrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmBgpNbrBag_KEYS.Unmarshal(m, b)
}
func (m *PmBgpNbrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmBgpNbrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmBgpNbrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmBgpNbrBag_KEYS.Merge(m, src)
}
func (m *PmBgpNbrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmBgpNbrBag_KEYS.Size(m)
}
func (m *PmBgpNbrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmBgpNbrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmBgpNbrBag_KEYS proto.InternalMessageInfo

func (m *PmBgpNbrBag_KEYS) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *PmBgpNbrBag_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmBgpNbrBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	InputMessages        uint32   `protobuf:"varint,51,opt,name=input_messages,json=inputMessages,proto3" json:"input_messages,omitempty"`
	OutputMessages       uint32   `protobuf:"varint,52,opt,name=output_messages,json=outputMessages,proto3" json:"output_messages,omitempty"`
	InputUpdateMessages  uint32   `protobuf:"varint,53,opt,name=input_update_messages,json=inputUpdateMessages,proto3" json:"input_update_messages,omitempty"`
	OutputUpdateMessages uint32   `protobuf:"varint,54,opt,name=output_update_messages,json=outputUpdateMessages,proto3" json:"output_update_messages,omitempty"`
	ConnEstablished      uint32   `protobuf:"varint,55,opt,name=conn_established,json=connEstablished,proto3" json:"conn_established,omitempty"`
	ConnDropped          uint32   `protobuf:"varint,56,opt,name=conn_dropped,json=connDropped,proto3" json:"conn_dropped,omitempty"`
	ErrorsReceived       uint32   `protobuf:"varint,57,opt,name=errors_received,json=errorsReceived,proto3" json:"errors_received,omitempty"`
	ErrorsSent           uint32   `protobuf:"varint,58,opt,name=errors_sent,json=errorsSent,proto3" json:"errors_sent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmBgpNbrBag) Reset()         { *m = PmBgpNbrBag{} }
func (m *PmBgpNbrBag) String() string { return proto.CompactTextString(m) }
func (*PmBgpNbrBag) ProtoMessage()    {}
func (*PmBgpNbrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e88bee685a3a839, []int{1}
}

func (m *PmBgpNbrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmBgpNbrBag.Unmarshal(m, b)
}
func (m *PmBgpNbrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmBgpNbrBag.Marshal(b, m, deterministic)
}
func (m *PmBgpNbrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmBgpNbrBag.Merge(m, src)
}
func (m *PmBgpNbrBag) XXX_Size() int {
	return xxx_messageInfo_PmBgpNbrBag.Size(m)
}
func (m *PmBgpNbrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmBgpNbrBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmBgpNbrBag proto.InternalMessageInfo

func (m *PmBgpNbrBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmBgpNbrBag) GetInputMessages() uint32 {
	if m != nil {
		return m.InputMessages
	}
	return 0
}

func (m *PmBgpNbrBag) GetOutputMessages() uint32 {
	if m != nil {
		return m.OutputMessages
	}
	return 0
}

func (m *PmBgpNbrBag) GetInputUpdateMessages() uint32 {
	if m != nil {
		return m.InputUpdateMessages
	}
	return 0
}

func (m *PmBgpNbrBag) GetOutputUpdateMessages() uint32 {
	if m != nil {
		return m.OutputUpdateMessages
	}
	return 0
}

func (m *PmBgpNbrBag) GetConnEstablished() uint32 {
	if m != nil {
		return m.ConnEstablished
	}
	return 0
}

func (m *PmBgpNbrBag) GetConnDropped() uint32 {
	if m != nil {
		return m.ConnDropped
	}
	return 0
}

func (m *PmBgpNbrBag) GetErrorsReceived() uint32 {
	if m != nil {
		return m.ErrorsReceived
	}
	return 0
}

func (m *PmBgpNbrBag) GetErrorsSent() uint32 {
	if m != nil {
		return m.ErrorsSent
	}
	return 0
}

func init() {
	proto.RegisterType((*PmBgpNbrBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.bgp.bgp_neighbors.bgp_neighbor.samples.sample.pm_bgp_nbr_bag_KEYS")
	proto.RegisterType((*PmBgpNbrBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.bgp.bgp_neighbors.bgp_neighbor.samples.sample.pm_bgp_nbr_bag")
}

func init() { proto.RegisterFile("pm_bgp_nbr_bag.proto", fileDescriptor_1e88bee685a3a839) }

var fileDescriptor_1e88bee685a3a839 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4d, 0xab, 0x13, 0x31,
	0x14, 0x86, 0x19, 0x15, 0x71, 0x52, 0xdb, 0xca, 0xb4, 0x4a, 0x40, 0xc4, 0x5a, 0x10, 0xeb, 0x66,
	0x16, 0x6d, 0xfd, 0xdc, 0x09, 0x76, 0x21, 0xe2, 0xc2, 0x29, 0x2e, 0x5c, 0x1d, 0x92, 0xc9, 0x71,
	0x1a, 0x68, 0x3e, 0x48, 0x52, 0xb9, 0x77, 0x79, 0xff, 0xf9, 0x25, 0xc9, 0xd0, 0x32, 0x5d, 0x0d,
	0xe7, 0x39, 0xcf, 0xfb, 0x32, 0x07, 0x42, 0xe6, 0x56, 0x01, 0xef, 0x2c, 0x68, 0xee, 0x80, 0xb3,
	0xae, 0xb6, 0xce, 0x04, 0x53, 0x1d, 0x5a, 0xe9, 0x5b, 0x03, 0xd2, 0x78, 0xb8, 0x71, 0xa0, 0x98,
	0x66, 0x1d, 0x32, 0x2e, 0x8f, 0x32, 0xdc, 0x82, 0x45, 0xf7, 0x4f, 0x75, 0x2a, 0x80, 0xb1, 0xe8,
	0xea, 0x38, 0x41, 0x1c, 0x6b, 0x65, 0xb4, 0x0c, 0xc6, 0xd5, 0xbc, 0xb3, 0x75, 0x6a, 0x44, 0xd9,
	0x1d, 0xb8, 0x71, 0x7e, 0x30, 0xd5, 0x9e, 0x29, 0x7b, 0x44, 0xdf, 0x7f, 0x97, 0xbf, 0xc9, 0x6c,
	0xf8, 0x07, 0xf0, 0x73, 0xf7, 0x77, 0x5f, 0xbd, 0x22, 0x44, 0x5a, 0x60, 0x42, 0x38, 0xf4, 0x9e,
	0x16, 0x8b, 0x62, 0x55, 0x36, 0xa5, 0xb4, 0xdf, 0x32, 0xa8, 0x5e, 0x92, 0x32, 0xe7, 0x41, 0x0a,
	0xfa, 0x60, 0x51, 0xac, 0xc6, 0xcd, 0x93, 0x0c, 0x7e, 0x88, 0xe5, 0xdd, 0x43, 0x32, 0x19, 0x76,
	0xc6, 0xba, 0x20, 0x15, 0x82, 0x0f, 0x4c, 0x59, 0xba, 0x5e, 0x14, 0xab, 0x47, 0x4d, 0x19, 0xc9,
	0x3e, 0x82, 0xea, 0x2d, 0x99, 0x48, 0x6d, 0x4f, 0x01, 0x14, 0x7a, 0xcf, 0x3a, 0xf4, 0x74, 0x93,
	0x3a, 0xc7, 0x89, 0xfe, 0xea, 0x61, 0xf5, 0x8e, 0x4c, 0xcd, 0x29, 0x0c, 0xbc, 0x6d, 0xf2, 0x26,
	0x19, 0x9f, 0xc5, 0x35, 0x79, 0x9e, 0xfb, 0x4e, 0x56, 0xb0, 0x80, 0x17, 0xfd, 0x43, 0xd2, 0x67,
	0x69, 0xf9, 0x27, 0xed, 0xce, 0x99, 0x2d, 0x79, 0xd1, 0x97, 0x5f, 0x87, 0x3e, 0xa6, 0xd0, 0x3c,
	0x6f, 0xaf, 0x52, 0xef, 0xc9, 0xb3, 0xd6, 0x68, 0x0d, 0xe8, 0x03, 0xe3, 0x47, 0xe9, 0x0f, 0x28,
	0xe8, 0xa7, 0xe4, 0x4f, 0x23, 0xdf, 0x5d, 0x70, 0xf5, 0x86, 0x3c, 0x4d, 0xaa, 0x70, 0xc6, 0x5a,
	0x14, 0xf4, 0x73, 0xd2, 0x46, 0x91, 0x7d, 0xcf, 0x28, 0x1e, 0x88, 0xce, 0x19, 0xe7, 0xc1, 0x61,
	0x8b, 0xf2, 0x3f, 0x0a, 0xfa, 0x25, 0x1f, 0x98, 0x71, 0xd3, 0xd3, 0xea, 0x35, 0x19, 0xf5, 0xa2,
	0x47, 0x1d, 0xe8, 0xd7, 0x24, 0x91, 0x8c, 0xf6, 0xa8, 0x03, 0x7f, 0x9c, 0xde, 0xd1, 0xe6, 0x3e,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x8e, 0x21, 0x62, 0x5f, 0x02, 0x00, 0x00,
}
