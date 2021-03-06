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
// source: pm_node_process_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_nodes_node_processes_process_samples_sample

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

type PmNodeProcessBag_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ProcessId            uint32   `protobuf:"varint,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	SampleId             uint32   `protobuf:"varint,3,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeProcessBag_KEYS) Reset()         { *m = PmNodeProcessBag_KEYS{} }
func (m *PmNodeProcessBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmNodeProcessBag_KEYS) ProtoMessage()    {}
func (*PmNodeProcessBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a79cfd41f3a229b3, []int{0}
}

func (m *PmNodeProcessBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeProcessBag_KEYS.Unmarshal(m, b)
}
func (m *PmNodeProcessBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeProcessBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmNodeProcessBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeProcessBag_KEYS.Merge(m, src)
}
func (m *PmNodeProcessBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmNodeProcessBag_KEYS.Size(m)
}
func (m *PmNodeProcessBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeProcessBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeProcessBag_KEYS proto.InternalMessageInfo

func (m *PmNodeProcessBag_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PmNodeProcessBag_KEYS) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *PmNodeProcessBag_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmNodeProcessBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	PeakMemory           uint32   `protobuf:"varint,51,opt,name=peak_memory,json=peakMemory,proto3" json:"peak_memory,omitempty"`
	AverageCpuUsed       uint32   `protobuf:"varint,52,opt,name=average_cpu_used,json=averageCpuUsed,proto3" json:"average_cpu_used,omitempty"`
	NoThreads            uint32   `protobuf:"varint,53,opt,name=no_threads,json=noThreads,proto3" json:"no_threads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeProcessBag) Reset()         { *m = PmNodeProcessBag{} }
func (m *PmNodeProcessBag) String() string { return proto.CompactTextString(m) }
func (*PmNodeProcessBag) ProtoMessage()    {}
func (*PmNodeProcessBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a79cfd41f3a229b3, []int{1}
}

func (m *PmNodeProcessBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeProcessBag.Unmarshal(m, b)
}
func (m *PmNodeProcessBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeProcessBag.Marshal(b, m, deterministic)
}
func (m *PmNodeProcessBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeProcessBag.Merge(m, src)
}
func (m *PmNodeProcessBag) XXX_Size() int {
	return xxx_messageInfo_PmNodeProcessBag.Size(m)
}
func (m *PmNodeProcessBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeProcessBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeProcessBag proto.InternalMessageInfo

func (m *PmNodeProcessBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmNodeProcessBag) GetPeakMemory() uint32 {
	if m != nil {
		return m.PeakMemory
	}
	return 0
}

func (m *PmNodeProcessBag) GetAverageCpuUsed() uint32 {
	if m != nil {
		return m.AverageCpuUsed
	}
	return 0
}

func (m *PmNodeProcessBag) GetNoThreads() uint32 {
	if m != nil {
		return m.NoThreads
	}
	return 0
}

func init() {
	proto.RegisterType((*PmNodeProcessBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.processes.process.samples.sample.pm_node_process_bag_KEYS")
	proto.RegisterType((*PmNodeProcessBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.processes.process.samples.sample.pm_node_process_bag")
}

func init() { proto.RegisterFile("pm_node_process_bag.proto", fileDescriptor_a79cfd41f3a229b3) }

var fileDescriptor_a79cfd41f3a229b3 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0xca, 0x6a, 0x23, 0x8a, 0xc4, 0x83, 0x11, 0x11, 0x97, 0x3d, 0xf5, 0xd4, 0x83,
	0xab, 0xbf, 0x40, 0x3c, 0x14, 0xf1, 0xd2, 0xd5, 0x83, 0xa7, 0x21, 0x6d, 0xc6, 0x1a, 0xdc, 0x74,
	0x42, 0x92, 0x8a, 0xfb, 0x67, 0xfc, 0xad, 0x92, 0xb4, 0xbd, 0xed, 0x25, 0x99, 0xf7, 0xbd, 0xe1,
	0xe5, 0x41, 0xd8, 0xb5, 0x35, 0xd0, 0x93, 0x42, 0xb0, 0x8e, 0x5a, 0xf4, 0x1e, 0x1a, 0xd9, 0x95,
	0xd6, 0x51, 0x20, 0xde, 0xb5, 0xda, 0xb7, 0x04, 0x9a, 0x3c, 0xfc, 0x3a, 0x30, 0xb2, 0x97, 0x1d,
	0xca, 0x46, 0x6f, 0x75, 0xd8, 0x81, 0x45, 0xf7, 0x69, 0x3a, 0x13, 0x80, 0x2c, 0xba, 0x32, 0x2a,
	0x88, 0x32, 0x4e, 0x9a, 0x94, 0x6e, 0xcb, 0x98, 0xe9, 0xd3, 0x59, 0x4e, 0xc9, 0xe8, 0xe7, 0xa9,
	0xf4, 0xd2, 0xd8, 0x2d, 0xce, 0xf7, 0x8a, 0x98, 0xd8, 0xd3, 0x02, 0x5e, 0x9e, 0x3f, 0x36, 0xfc,
	0x8a, 0x1d, 0x27, 0x43, 0x2b, 0x91, 0x2d, 0xb3, 0x22, 0xaf, 0x17, 0x51, 0x56, 0x8a, 0xdf, 0x32,
	0x36, 0x2f, 0x6b, 0x25, 0x0e, 0x96, 0x59, 0x71, 0x56, 0xe7, 0x13, 0xa9, 0x14, 0xbf, 0x61, 0xf9,
	0x98, 0x1e, 0xdd, 0xc3, 0xe4, 0x9e, 0x8c, 0xa0, 0x52, 0xab, 0xbf, 0x8c, 0x5d, 0xee, 0x79, 0x31,
	0x66, 0x06, 0x6d, 0x10, 0x7c, 0x90, 0xc6, 0x8a, 0xfb, 0x65, 0x56, 0x1c, 0xd5, 0x79, 0x24, 0x9b,
	0x08, 0xf8, 0x1d, 0x3b, 0xb5, 0x28, 0xbf, 0xc1, 0xa0, 0x21, 0xb7, 0x13, 0xeb, 0x94, 0xca, 0x22,
	0x7a, 0x4d, 0x84, 0x17, 0xec, 0x42, 0xfe, 0xa0, 0x93, 0x1d, 0x42, 0x6b, 0x07, 0x18, 0x3c, 0x2a,
	0xf1, 0x90, 0xb6, 0xce, 0x27, 0xfe, 0x64, 0x87, 0x77, 0x8f, 0xa9, 0x7d, 0x4f, 0x10, 0xbe, 0x1c,
	0x4a, 0xe5, 0xc5, 0xe3, 0xd8, 0xbe, 0xa7, 0xb7, 0x11, 0x34, 0x8b, 0xf4, 0x03, 0xeb, 0xff, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xce, 0x4a, 0x95, 0x03, 0x9e, 0x01, 0x00, 0x00,
}
