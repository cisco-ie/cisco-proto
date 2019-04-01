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
// source: pm_node_mem_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_nodes_node_samples_sample

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

type PmNodeMemBag_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeMemBag_KEYS) Reset()         { *m = PmNodeMemBag_KEYS{} }
func (m *PmNodeMemBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmNodeMemBag_KEYS) ProtoMessage()    {}
func (*PmNodeMemBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f4228fa296845a8, []int{0}
}

func (m *PmNodeMemBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeMemBag_KEYS.Unmarshal(m, b)
}
func (m *PmNodeMemBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeMemBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmNodeMemBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeMemBag_KEYS.Merge(m, src)
}
func (m *PmNodeMemBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmNodeMemBag_KEYS.Size(m)
}
func (m *PmNodeMemBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeMemBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeMemBag_KEYS proto.InternalMessageInfo

func (m *PmNodeMemBag_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PmNodeMemBag_KEYS) GetSampleId() int32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmNodeMemBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	CurrMemory           uint32   `protobuf:"varint,51,opt,name=curr_memory,json=currMemory,proto3" json:"curr_memory,omitempty"`
	PeakMemory           uint32   `protobuf:"varint,52,opt,name=peak_memory,json=peakMemory,proto3" json:"peak_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeMemBag) Reset()         { *m = PmNodeMemBag{} }
func (m *PmNodeMemBag) String() string { return proto.CompactTextString(m) }
func (*PmNodeMemBag) ProtoMessage()    {}
func (*PmNodeMemBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f4228fa296845a8, []int{1}
}

func (m *PmNodeMemBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeMemBag.Unmarshal(m, b)
}
func (m *PmNodeMemBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeMemBag.Marshal(b, m, deterministic)
}
func (m *PmNodeMemBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeMemBag.Merge(m, src)
}
func (m *PmNodeMemBag) XXX_Size() int {
	return xxx_messageInfo_PmNodeMemBag.Size(m)
}
func (m *PmNodeMemBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeMemBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeMemBag proto.InternalMessageInfo

func (m *PmNodeMemBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmNodeMemBag) GetCurrMemory() uint32 {
	if m != nil {
		return m.CurrMemory
	}
	return 0
}

func (m *PmNodeMemBag) GetPeakMemory() uint32 {
	if m != nil {
		return m.PeakMemory
	}
	return 0
}

func init() {
	proto.RegisterType((*PmNodeMemBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.samples.sample.pm_node_mem_bag_KEYS")
	proto.RegisterType((*PmNodeMemBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.samples.sample.pm_node_mem_bag")
}

func init() { proto.RegisterFile("pm_node_mem_bag.proto", fileDescriptor_1f4228fa296845a8) }

var fileDescriptor_1f4228fa296845a8 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0x68, 0x75, 0x47, 0x44, 0x08, 0x8a, 0x0b, 0x22, 0x2e, 0x3d, 0xed, 0x69, 0x0f,
	0xd6, 0x57, 0xf0, 0x50, 0xd4, 0xcb, 0x16, 0x0f, 0x9e, 0x86, 0xec, 0x26, 0x2e, 0xc1, 0x4e, 0x13,
	0x26, 0x11, 0xec, 0xdb, 0xcb, 0x6c, 0xec, 0xa5, 0x97, 0x24, 0xff, 0x97, 0x2f, 0x7f, 0x60, 0xe0,
	0x36, 0x12, 0xee, 0x82, 0x75, 0x48, 0x8e, 0x70, 0x30, 0x53, 0x17, 0x39, 0xe4, 0xa0, 0x3f, 0x46,
	0x9f, 0xc6, 0x80, 0x3e, 0x24, 0xfc, 0x65, 0x24, 0xb3, 0x33, 0x93, 0x33, 0x83, 0xdf, 0xfa, 0xbc,
	0xc7, 0xe8, 0xf8, 0x8b, 0x26, 0xca, 0x18, 0xa2, 0xe3, 0x4e, 0x12, 0x4a, 0x94, 0x93, 0x0f, 0xd6,
	0x8f, 0x9d, 0xf4, 0xa5, 0x79, 0xed, 0x92, 0xa1, 0xb8, 0x75, 0xe9, 0x7f, 0x5f, 0xbe, 0xc1, 0xcd,
	0xd1, 0x7f, 0xf8, 0xfa, 0xf2, 0xb9, 0xd1, 0x77, 0x70, 0x3e, 0x43, 0x6f, 0x6b, 0xd5, 0xa8, 0xb6,
	0xea, 0x17, 0x12, 0xd7, 0x56, 0xdf, 0x43, 0x55, 0x9e, 0xca, 0xd5, 0x49, 0xa3, 0xda, 0xb3, 0xfe,
	0xa2, 0x80, 0xb5, 0x5d, 0x32, 0x5c, 0x1f, 0xb5, 0xe9, 0x07, 0x80, 0xec, 0xc9, 0x61, 0xca, 0x86,
	0x62, 0xfd, 0xd4, 0xa8, 0xf6, 0xb4, 0xaf, 0x84, 0x6c, 0x04, 0xe8, 0x47, 0xb8, 0x1c, 0x7f, 0x98,
	0x45, 0x0f, 0xbc, 0xaf, 0x57, 0x8d, 0x6a, 0xaf, 0x7a, 0x10, 0xf4, 0x3e, 0x13, 0x11, 0xa2, 0x33,
	0xdf, 0x07, 0xe1, 0xb9, 0x08, 0x82, 0x8a, 0x30, 0x2c, 0xe6, 0xf9, 0xac, 0xfe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0xaa, 0x80, 0x75, 0x38, 0x01, 0x00, 0x00,
}
