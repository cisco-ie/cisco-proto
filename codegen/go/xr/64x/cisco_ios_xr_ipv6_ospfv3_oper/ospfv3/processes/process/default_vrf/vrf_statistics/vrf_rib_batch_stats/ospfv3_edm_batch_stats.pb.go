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

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_vrf_statistics_vrf_rib_batch_stats

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
	proto.RegisterType((*Ospfv3EdmBatchStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.vrf_statistics.vrf_rib_batch_stats.ospfv3_edm_batch_stats_KEYS")
	proto.RegisterType((*Ospfv3EdmBatchStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.vrf_statistics.vrf_rib_batch_stats.ospfv3_edm_batch_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_batch_stats.proto", fileDescriptor_1109e6f43c52bb88) }

var fileDescriptor_1109e6f43c52bb88 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3d, 0x4f, 0xeb, 0x30,
	0x14, 0x40, 0xd5, 0x37, 0x3c, 0xbd, 0xba, 0x1f, 0x43, 0x86, 0xa7, 0x48, 0x05, 0x51, 0x0a, 0x48,
	0x9d, 0x32, 0x50, 0x28, 0x5f, 0x0b, 0x95, 0xda, 0x09, 0x09, 0xa1, 0x76, 0x62, 0xba, 0x72, 0xe3,
	0x1b, 0x6a, 0x29, 0x89, 0x23, 0x5f, 0x37, 0xea, 0x8f, 0xe6, 0x47, 0xa0, 0x5c, 0x27, 0x34, 0x43,
	0x37, 0xfb, 0xdc, 0xe3, 0x73, 0x33, 0x44, 0x9c, 0x19, 0x2a, 0x92, 0x72, 0x06, 0xa8, 0x32, 0xd8,
	0x4a, 0x17, 0xef, 0x80, 0x9c, 0x74, 0x14, 0x15, 0xd6, 0x38, 0x13, 0x60, 0xac, 0x29, 0x36, 0xa0,
	0x0d, 0xc1, 0xc1, 0x82, 0x2e, 0xca, 0x39, 0xd4, 0xbe, 0x29, 0xd0, 0x46, 0xfe, 0x5c, 0xb9, 0x31,
	0x12, 0x21, 0x35, 0xa7, 0x48, 0x61, 0x22, 0xf7, 0xa9, 0x83, 0xd2, 0x26, 0x51, 0x69, 0x13, 0x6e,
	0x6a, 0x72, 0x3a, 0x26, 0xbe, 0x5a, 0xbd, 0x6d, 0x2f, 0x9b, 0xbc, 0x8a, 0xd1, 0xe9, 0xcf, 0x80,
	0xb7, 0xd5, 0xe7, 0x26, 0xb8, 0x14, 0xfd, 0x3a, 0x0c, 0xb9, 0xcc, 0x30, 0xec, 0x8c, 0x3b, 0xd3,
	0xee, 0xba, 0x57, 0xb3, 0x77, 0x99, 0xe1, 0xe4, 0xfb, 0x8f, 0xf8, 0x7f, 0x3a, 0x51, 0xbd, 0xe6,
	0x2b, 0x12, 0x10, 0xe6, 0x2e, 0xbc, 0x1d, 0x77, 0xa6, 0x83, 0x75, 0xaf, 0x66, 0x1b, 0xcc, 0x5d,
	0x70, 0x21, 0x7a, 0xd6, 0xec, 0x5d, 0x63, 0xcc, 0xd8, 0x10, 0x1e, 0xb1, 0x70, 0x23, 0x86, 0xb5,
	0x50, 0x60, 0xae, 0x74, 0xfe, 0x15, 0xde, 0xb1, 0x33, 0xf0, 0xf4, 0xc3, 0xc3, 0xe0, 0x5c, 0x88,
	0x4c, 0x1e, 0xc0, 0xc3, 0xf0, 0x9e, 0x95, 0x6e, 0x26, 0x0f, 0x6b, 0x06, 0xd5, 0xb8, 0x90, 0x6e,
	0x57, 0x6f, 0x99, 0xfb, 0x31, 0x13, 0x5e, 0x72, 0x25, 0x06, 0x7e, 0xdc, 0xec, 0x78, 0x60, 0xa3,
	0xcf, 0xb0, 0x59, 0x31, 0x12, 0xfe, 0x05, 0x48, 0xa5, 0xc2, 0x47, 0x16, 0xfe, 0x31, 0x58, 0x28,
	0x15, 0x5c, 0x8b, 0xe1, 0xef, 0x10, 0xd0, 0x5a, 0x0a, 0x9f, 0x5a, 0x89, 0x85, 0x52, 0x2b, 0x6b,
	0xe9, 0x98, 0x50, 0x98, 0x86, 0xcf, 0xad, 0xc4, 0x12, 0xd3, 0x63, 0x42, 0x61, 0xea, 0x13, 0x2f,
	0xad, 0xc4, 0x12, 0xd3, 0x2a, 0xb1, 0xfd, 0xcb, 0xbf, 0xc7, 0xec, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x08, 0x46, 0x85, 0x3e, 0x02, 0x00, 0x00,
}
