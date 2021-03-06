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
// source: pm_node_cpu_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_periodic_nodes_node_sample_xr_sample

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

type PmNodeCpuBag_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	SampleId             uint32   `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeCpuBag_KEYS) Reset()         { *m = PmNodeCpuBag_KEYS{} }
func (m *PmNodeCpuBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmNodeCpuBag_KEYS) ProtoMessage()    {}
func (*PmNodeCpuBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdeece9996eaf6dc, []int{0}
}

func (m *PmNodeCpuBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeCpuBag_KEYS.Unmarshal(m, b)
}
func (m *PmNodeCpuBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeCpuBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmNodeCpuBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeCpuBag_KEYS.Merge(m, src)
}
func (m *PmNodeCpuBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmNodeCpuBag_KEYS.Size(m)
}
func (m *PmNodeCpuBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeCpuBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeCpuBag_KEYS proto.InternalMessageInfo

func (m *PmNodeCpuBag_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PmNodeCpuBag_KEYS) GetSampleId() uint32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmNodeCpuBag struct {
	TimeStamp            uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	NoProcesses          uint32   `protobuf:"varint,51,opt,name=no_processes,json=noProcesses,proto3" json:"no_processes,omitempty"`
	AverageCpuUsed       uint32   `protobuf:"varint,52,opt,name=average_cpu_used,json=averageCpuUsed,proto3" json:"average_cpu_used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmNodeCpuBag) Reset()         { *m = PmNodeCpuBag{} }
func (m *PmNodeCpuBag) String() string { return proto.CompactTextString(m) }
func (*PmNodeCpuBag) ProtoMessage()    {}
func (*PmNodeCpuBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdeece9996eaf6dc, []int{1}
}

func (m *PmNodeCpuBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmNodeCpuBag.Unmarshal(m, b)
}
func (m *PmNodeCpuBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmNodeCpuBag.Marshal(b, m, deterministic)
}
func (m *PmNodeCpuBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmNodeCpuBag.Merge(m, src)
}
func (m *PmNodeCpuBag) XXX_Size() int {
	return xxx_messageInfo_PmNodeCpuBag.Size(m)
}
func (m *PmNodeCpuBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmNodeCpuBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmNodeCpuBag proto.InternalMessageInfo

func (m *PmNodeCpuBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmNodeCpuBag) GetNoProcesses() uint32 {
	if m != nil {
		return m.NoProcesses
	}
	return 0
}

func (m *PmNodeCpuBag) GetAverageCpuUsed() uint32 {
	if m != nil {
		return m.AverageCpuUsed
	}
	return 0
}

func init() {
	proto.RegisterType((*PmNodeCpuBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.sample_xr.sample.pm_node_cpu_bag_KEYS")
	proto.RegisterType((*PmNodeCpuBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.periodic.nodes.node.sample_xr.sample.pm_node_cpu_bag")
}

func init() { proto.RegisterFile("pm_node_cpu_bag.proto", fileDescriptor_bdeece9996eaf6dc) }

var fileDescriptor_bdeece9996eaf6dc = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xc8, 0xea, 0xc6, 0x4f, 0x82, 0x62, 0x41, 0x84, 0xba, 0xa7, 0x9e, 0x7a, 0x70,
	0xfd, 0x07, 0xe2, 0x61, 0xd1, 0x83, 0x74, 0x11, 0xf1, 0x34, 0xa4, 0xc9, 0x58, 0x02, 0x9b, 0x4e,
	0xc8, 0xa4, 0xb2, 0x1e, 0xfc, 0xef, 0x92, 0x6c, 0xbd, 0xec, 0x65, 0x78, 0xdf, 0x87, 0xe1, 0x19,
	0x18, 0x71, 0xed, 0x1d, 0x0c, 0x64, 0x10, 0xb4, 0x1f, 0xa1, 0x53, 0x7d, 0xe3, 0x03, 0x45, 0x92,
	0x1f, 0xda, 0xb2, 0x26, 0xb0, 0xc4, 0xb0, 0x0d, 0xe0, 0xd4, 0xa0, 0x7a, 0x54, 0x9d, 0xdd, 0xd8,
	0xf8, 0x03, 0x1e, 0xc3, 0x97, 0xeb, 0x5d, 0x04, 0xf2, 0x18, 0x9a, 0xd4, 0x20, 0xd5, 0x94, 0x2c,
	0x19, 0xab, 0x9b, 0xe4, 0xe3, 0x3c, 0x1b, 0x56, 0xce, 0x6f, 0x10, 0xb6, 0x61, 0x4a, 0x8b, 0x57,
	0x71, 0xb5, 0x77, 0x11, 0x5e, 0x9e, 0x3f, 0xd7, 0xf2, 0x46, 0x1c, 0x65, 0x68, 0x4d, 0x59, 0x54,
	0x45, 0x3d, 0x6f, 0x67, 0xa9, 0xae, 0x8c, 0xbc, 0x15, 0xf3, 0x49, 0x62, 0x4d, 0x79, 0x50, 0x15,
	0xf5, 0x59, 0x7b, 0xbc, 0x03, 0x2b, 0xb3, 0xf8, 0x15, 0x17, 0x7b, 0x36, 0x79, 0x27, 0x44, 0xb4,
	0x0e, 0x81, 0xa3, 0x72, 0xbe, 0x7c, 0xa8, 0x8a, 0xfa, 0xb0, 0x9d, 0x27, 0xb2, 0x4e, 0x40, 0xde,
	0x8b, 0xd3, 0x81, 0xc0, 0x07, 0xd2, 0xc8, 0x8c, 0x5c, 0x2e, 0xb3, 0xf1, 0x64, 0xa0, 0xb7, 0x7f,
	0x24, 0x6b, 0x71, 0xa9, 0xbe, 0x31, 0xa8, 0x7e, 0x27, 0x1d, 0x19, 0x4d, 0xf9, 0x98, 0xd7, 0xce,
	0x27, 0xfe, 0xe4, 0xc7, 0x77, 0x46, 0xd3, 0xcd, 0xf2, 0xb3, 0x96, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x01, 0xe7, 0x86, 0x45, 0x45, 0x01, 0x00, 0x00,
}
