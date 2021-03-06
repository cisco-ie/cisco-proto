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
// source: ospfv3_edm_batch_stats.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_vrf_statistics_vrf_rib_batch_stats

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

type Ospfv3EdmBatchStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmBatchStats_KEYS) Reset()         { *m = Ospfv3EdmBatchStats_KEYS{} }
func (m *Ospfv3EdmBatchStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmBatchStats_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmBatchStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1109e6f43c52bb88, []int{0}
}

func (m *Ospfv3EdmBatchStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmBatchStats_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmBatchStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmBatchStats_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmBatchStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmBatchStats_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmBatchStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmBatchStats_KEYS.Size(m)
}
func (m *Ospfv3EdmBatchStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmBatchStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmBatchStats_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmBatchStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmBatchStats_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type Ospfv3EdmBatchStats struct {
	BatchesSent          uint32   `protobuf:"varint,50,opt,name=batches_sent,json=batchesSent,proto3" json:"batches_sent,omitempty"`
	RoutesSent           uint32   `protobuf:"varint,51,opt,name=routes_sent,json=routesSent,proto3" json:"routes_sent,omitempty"`
	RoutesPending        uint32   `protobuf:"varint,52,opt,name=routes_pending,json=routesPending,proto3" json:"routes_pending,omitempty"`
	MaxRoutes            uint32   `protobuf:"varint,53,opt,name=max_routes,json=maxRoutes,proto3" json:"max_routes,omitempty"`
	PathsSent            uint32   `protobuf:"varint,54,opt,name=paths_sent,json=pathsSent,proto3" json:"paths_sent,omitempty"`
	PathsPending         uint32   `protobuf:"varint,55,opt,name=paths_pending,json=pathsPending,proto3" json:"paths_pending,omitempty"`
	PathsAdd             uint32   `protobuf:"varint,56,opt,name=paths_add,json=pathsAdd,proto3" json:"paths_add,omitempty"`
	PathsAddErrs         uint32   `protobuf:"varint,57,opt,name=paths_add_errs,json=pathsAddErrs,proto3" json:"paths_add_errs,omitempty"`
	PathsDel             uint32   `protobuf:"varint,58,opt,name=paths_del,json=pathsDel,proto3" json:"paths_del,omitempty"`
	PathsDelErrs         uint32   `protobuf:"varint,59,opt,name=paths_del_errs,json=pathsDelErrs,proto3" json:"paths_del_errs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmBatchStats) Reset()         { *m = Ospfv3EdmBatchStats{} }
func (m *Ospfv3EdmBatchStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmBatchStats) ProtoMessage()    {}
func (*Ospfv3EdmBatchStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1109e6f43c52bb88, []int{1}
}

func (m *Ospfv3EdmBatchStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmBatchStats.Unmarshal(m, b)
}
func (m *Ospfv3EdmBatchStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmBatchStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmBatchStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmBatchStats.Merge(m, src)
}
func (m *Ospfv3EdmBatchStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmBatchStats.Size(m)
}
func (m *Ospfv3EdmBatchStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmBatchStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmBatchStats proto.InternalMessageInfo

func (m *Ospfv3EdmBatchStats) GetBatchesSent() uint32 {
	if m != nil {
		return m.BatchesSent
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetRoutesSent() uint32 {
	if m != nil {
		return m.RoutesSent
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetRoutesPending() uint32 {
	if m != nil {
		return m.RoutesPending
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetMaxRoutes() uint32 {
	if m != nil {
		return m.MaxRoutes
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsSent() uint32 {
	if m != nil {
		return m.PathsSent
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsPending() uint32 {
	if m != nil {
		return m.PathsPending
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsAdd() uint32 {
	if m != nil {
		return m.PathsAdd
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsAddErrs() uint32 {
	if m != nil {
		return m.PathsAddErrs
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsDel() uint32 {
	if m != nil {
		return m.PathsDel
	}
	return 0
}

func (m *Ospfv3EdmBatchStats) GetPathsDelErrs() uint32 {
	if m != nil {
		return m.PathsDelErrs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmBatchStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.vrf_rib_batch_stats.ospfv3_edm_batch_stats_KEYS")
	proto.RegisterType((*Ospfv3EdmBatchStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.vrf_statistics.vrf_rib_batch_stats.ospfv3_edm_batch_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_batch_stats.proto", fileDescriptor_1109e6f43c52bb88) }

var fileDescriptor_1109e6f43c52bb88 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3d, 0x6f, 0xea, 0x30,
	0x14, 0x86, 0x05, 0xc3, 0xbd, 0x60, 0x3e, 0x86, 0x0c, 0x57, 0xbe, 0xe2, 0x5e, 0x95, 0xd2, 0x56,
	0x62, 0x62, 0x28, 0x2d, 0xfd, 0x9a, 0x90, 0x60, 0xaa, 0x54, 0x55, 0x30, 0x55, 0x1d, 0x8e, 0x9c,
	0xf8, 0x50, 0x2c, 0x25, 0x71, 0x64, 0xbb, 0x11, 0x3f, 0xba, 0x3f, 0xa2, 0xf2, 0x71, 0x52, 0x32,
	0x30, 0x24, 0x72, 0x9e, 0xf7, 0xcd, 0xf3, 0x66, 0x08, 0xfb, 0xa7, 0x6d, 0xb1, 0x2b, 0xe7, 0x80,
	0x32, 0x83, 0x58, 0xb8, 0x64, 0x0f, 0xd6, 0x09, 0x67, 0x67, 0x85, 0xd1, 0x4e, 0x47, 0x71, 0xa2,
	0x6c, 0xa2, 0x41, 0x69, 0x0b, 0x07, 0x03, 0xaa, 0x28, 0x17, 0x50, 0xf5, 0x75, 0x81, 0x66, 0x16,
	0xce, 0xbe, 0x9b, 0xa0, 0xb5, 0x68, 0xeb, 0xd3, 0xac, 0x34, 0x3b, 0xba, 0xf9, 0x8b, 0x84, 0xca,
	0x3a, 0x95, 0x10, 0x02, 0xa3, 0xe2, 0xe6, 0xd2, 0xe4, 0x9d, 0x8d, 0x4e, 0x7f, 0x03, 0x3c, 0xaf,
	0xdf, 0xb6, 0xd1, 0x39, 0xeb, 0x57, 0x56, 0xc8, 0x45, 0x86, 0xbc, 0x35, 0x6e, 0x4d, 0xbb, 0x9b,
	0x5e, 0xc5, 0x5e, 0x44, 0x86, 0xd1, 0x5f, 0xd6, 0xf1, 0x62, 0x8a, 0xdb, 0x14, 0xff, 0x2e, 0xcd,
	0xce, 0x47, 0x93, 0xaf, 0x36, 0xfb, 0x73, 0xda, 0xee, 0xc5, 0xf4, 0x88, 0x16, 0x2c, 0xe6, 0x8e,
	0x5f, 0x8f, 0x5b, 0xd3, 0xc1, 0xa6, 0x57, 0xb1, 0x2d, 0xe6, 0x2e, 0x3a, 0x63, 0x3d, 0xa3, 0x3f,
	0x5d, 0xdd, 0x98, 0x53, 0x83, 0x05, 0x44, 0x85, 0x2b, 0x36, 0xac, 0x0a, 0x05, 0xe6, 0x52, 0xe5,
	0x1f, 0xfc, 0x86, 0x3a, 0x83, 0x40, 0x5f, 0x03, 0x8c, 0xfe, 0x33, 0x96, 0x89, 0x03, 0x04, 0xc8,
	0x6f, 0xa9, 0xd2, 0xcd, 0xc4, 0x61, 0x43, 0xc0, 0xc7, 0x85, 0x70, 0xfb, 0x6a, 0x65, 0x11, 0x62,
	0x22, 0x34, 0x72, 0xc1, 0x06, 0x21, 0xae, 0x37, 0xee, 0xa8, 0xd1, 0x27, 0x58, 0x4f, 0x8c, 0x58,
	0x78, 0x03, 0x84, 0x94, 0xfc, 0x9e, 0x0a, 0x1d, 0x02, 0x4b, 0x29, 0xa3, 0x4b, 0x36, 0xfc, 0x09,
	0x01, 0x8d, 0xb1, 0xfc, 0xa1, 0xa1, 0x58, 0x4a, 0xb9, 0x36, 0xc6, 0x1e, 0x15, 0x12, 0x53, 0xfe,
	0xd8, 0x50, 0xac, 0x30, 0x3d, 0x2a, 0x24, 0xa6, 0x41, 0xf1, 0xd4, 0x50, 0xac, 0x30, 0xf5, 0x8a,
	0xf8, 0x17, 0xfd, 0x36, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x8c, 0xf0, 0xb9, 0x56,
	0x02, 0x00, 0x00,
}
