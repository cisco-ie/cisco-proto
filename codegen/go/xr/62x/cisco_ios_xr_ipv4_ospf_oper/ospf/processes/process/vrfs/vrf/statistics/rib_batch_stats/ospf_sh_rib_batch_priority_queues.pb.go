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
// source: ospf_sh_rib_batch_priority_queues.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_statistics_rib_batch_stats

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

type OspfShRibBatchPriorityQueues_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRibBatchPriorityQueues_KEYS) Reset()         { *m = OspfShRibBatchPriorityQueues_KEYS{} }
func (m *OspfShRibBatchPriorityQueues_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRibBatchPriorityQueues_KEYS) ProtoMessage()    {}
func (*OspfShRibBatchPriorityQueues_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e167d1b65f82749, []int{0}
}

func (m *OspfShRibBatchPriorityQueues_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS.Unmarshal(m, b)
}
func (m *OspfShRibBatchPriorityQueues_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRibBatchPriorityQueues_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS.Merge(m, src)
}
func (m *OspfShRibBatchPriorityQueues_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS.Size(m)
}
func (m *OspfShRibBatchPriorityQueues_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRibBatchPriorityQueues_KEYS proto.InternalMessageInfo

func (m *OspfShRibBatchPriorityQueues_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRibBatchPriorityQueues_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShRibBatch struct {
	BatchesSent          uint32   `protobuf:"varint,1,opt,name=batches_sent,json=batchesSent,proto3" json:"batches_sent,omitempty"`
	RoutesSent           uint32   `protobuf:"varint,2,opt,name=routes_sent,json=routesSent,proto3" json:"routes_sent,omitempty"`
	RoutesPending        uint32   `protobuf:"varint,3,opt,name=routes_pending,json=routesPending,proto3" json:"routes_pending,omitempty"`
	MaxRoutes            uint32   `protobuf:"varint,4,opt,name=max_routes,json=maxRoutes,proto3" json:"max_routes,omitempty"`
	PathsSent            uint32   `protobuf:"varint,5,opt,name=paths_sent,json=pathsSent,proto3" json:"paths_sent,omitempty"`
	PathsPending         uint32   `protobuf:"varint,6,opt,name=paths_pending,json=pathsPending,proto3" json:"paths_pending,omitempty"`
	PathsAdd             uint32   `protobuf:"varint,7,opt,name=paths_add,json=pathsAdd,proto3" json:"paths_add,omitempty"`
	PathsAddErrs         uint32   `protobuf:"varint,8,opt,name=paths_add_errs,json=pathsAddErrs,proto3" json:"paths_add_errs,omitempty"`
	PathsDel             uint32   `protobuf:"varint,9,opt,name=paths_del,json=pathsDel,proto3" json:"paths_del,omitempty"`
	PathsDelErrs         uint32   `protobuf:"varint,10,opt,name=paths_del_errs,json=pathsDelErrs,proto3" json:"paths_del_errs,omitempty"`
	LfasPending          uint32   `protobuf:"varint,11,opt,name=lfas_pending,json=lfasPending,proto3" json:"lfas_pending,omitempty"`
	LfasAdd              uint32   `protobuf:"varint,12,opt,name=lfas_add,json=lfasAdd,proto3" json:"lfas_add,omitempty"`
	LfasDel              uint32   `protobuf:"varint,13,opt,name=lfas_del,json=lfasDel,proto3" json:"lfas_del,omitempty"`
	LfasSent             uint32   `protobuf:"varint,14,opt,name=lfas_sent,json=lfasSent,proto3" json:"lfas_sent,omitempty"`
	PriorityLevel        string   `protobuf:"bytes,15,opt,name=priority_level,json=priorityLevel,proto3" json:"priority_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRibBatch) Reset()         { *m = OspfShRibBatch{} }
func (m *OspfShRibBatch) String() string { return proto.CompactTextString(m) }
func (*OspfShRibBatch) ProtoMessage()    {}
func (*OspfShRibBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e167d1b65f82749, []int{1}
}

func (m *OspfShRibBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRibBatch.Unmarshal(m, b)
}
func (m *OspfShRibBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRibBatch.Marshal(b, m, deterministic)
}
func (m *OspfShRibBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRibBatch.Merge(m, src)
}
func (m *OspfShRibBatch) XXX_Size() int {
	return xxx_messageInfo_OspfShRibBatch.Size(m)
}
func (m *OspfShRibBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRibBatch.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRibBatch proto.InternalMessageInfo

func (m *OspfShRibBatch) GetBatchesSent() uint32 {
	if m != nil {
		return m.BatchesSent
	}
	return 0
}

func (m *OspfShRibBatch) GetRoutesSent() uint32 {
	if m != nil {
		return m.RoutesSent
	}
	return 0
}

func (m *OspfShRibBatch) GetRoutesPending() uint32 {
	if m != nil {
		return m.RoutesPending
	}
	return 0
}

func (m *OspfShRibBatch) GetMaxRoutes() uint32 {
	if m != nil {
		return m.MaxRoutes
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsSent() uint32 {
	if m != nil {
		return m.PathsSent
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsPending() uint32 {
	if m != nil {
		return m.PathsPending
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsAdd() uint32 {
	if m != nil {
		return m.PathsAdd
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsAddErrs() uint32 {
	if m != nil {
		return m.PathsAddErrs
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsDel() uint32 {
	if m != nil {
		return m.PathsDel
	}
	return 0
}

func (m *OspfShRibBatch) GetPathsDelErrs() uint32 {
	if m != nil {
		return m.PathsDelErrs
	}
	return 0
}

func (m *OspfShRibBatch) GetLfasPending() uint32 {
	if m != nil {
		return m.LfasPending
	}
	return 0
}

func (m *OspfShRibBatch) GetLfasAdd() uint32 {
	if m != nil {
		return m.LfasAdd
	}
	return 0
}

func (m *OspfShRibBatch) GetLfasDel() uint32 {
	if m != nil {
		return m.LfasDel
	}
	return 0
}

func (m *OspfShRibBatch) GetLfasSent() uint32 {
	if m != nil {
		return m.LfasSent
	}
	return 0
}

func (m *OspfShRibBatch) GetPriorityLevel() string {
	if m != nil {
		return m.PriorityLevel
	}
	return ""
}

type OspfShRibBatchPriorityQueues struct {
	PriorityBatch        []*OspfShRibBatch `protobuf:"bytes,50,rep,name=priority_batch,json=priorityBatch,proto3" json:"priority_batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OspfShRibBatchPriorityQueues) Reset()         { *m = OspfShRibBatchPriorityQueues{} }
func (m *OspfShRibBatchPriorityQueues) String() string { return proto.CompactTextString(m) }
func (*OspfShRibBatchPriorityQueues) ProtoMessage()    {}
func (*OspfShRibBatchPriorityQueues) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e167d1b65f82749, []int{2}
}

func (m *OspfShRibBatchPriorityQueues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues.Unmarshal(m, b)
}
func (m *OspfShRibBatchPriorityQueues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues.Marshal(b, m, deterministic)
}
func (m *OspfShRibBatchPriorityQueues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRibBatchPriorityQueues.Merge(m, src)
}
func (m *OspfShRibBatchPriorityQueues) XXX_Size() int {
	return xxx_messageInfo_OspfShRibBatchPriorityQueues.Size(m)
}
func (m *OspfShRibBatchPriorityQueues) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRibBatchPriorityQueues.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRibBatchPriorityQueues proto.InternalMessageInfo

func (m *OspfShRibBatchPriorityQueues) GetPriorityBatch() []*OspfShRibBatch {
	if m != nil {
		return m.PriorityBatch
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRibBatchPriorityQueues_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.rib_batch_stats.ospf_sh_rib_batch_priority_queues_KEYS")
	proto.RegisterType((*OspfShRibBatch)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.rib_batch_stats.ospf_sh_rib_batch")
	proto.RegisterType((*OspfShRibBatchPriorityQueues)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.rib_batch_stats.ospf_sh_rib_batch_priority_queues")
}

func init() {
	proto.RegisterFile("ospf_sh_rib_batch_priority_queues.proto", fileDescriptor_5e167d1b65f82749)
}

var fileDescriptor_5e167d1b65f82749 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xbd, 0x6e, 0xdb, 0x30,
	0x14, 0x85, 0xe1, 0x24, 0x8d, 0xed, 0x2b, 0xcb, 0x45, 0x35, 0x31, 0x08, 0x8a, 0xc6, 0xee, 0x5f,
	0x26, 0x0d, 0x69, 0x5f, 0x20, 0x45, 0x32, 0xb5, 0x28, 0x0a, 0x07, 0x28, 0xd0, 0x89, 0xa0, 0xc5,
	0xab, 0x9a, 0x80, 0x2c, 0xaa, 0x24, 0x2d, 0xb8, 0x7b, 0x1f, 0xa0, 0x4f, 0xd3, 0xe7, 0x2b, 0xee,
	0xa5, 0x14, 0x19, 0xc8, 0x90, 0x29, 0x8b, 0x41, 0x7d, 0xe7, 0xf0, 0xdc, 0x03, 0x9a, 0x84, 0xf7,
	0xd6, 0x37, 0xa5, 0xf4, 0x1b, 0xe9, 0xcc, 0x5a, 0xae, 0x55, 0x28, 0x36, 0xb2, 0x71, 0xc6, 0x3a,
	0x13, 0x7e, 0xcb, 0x5f, 0x3b, 0xdc, 0xa1, 0xcf, 0x1b, 0x67, 0x83, 0xcd, 0xbe, 0x17, 0xc6, 0x17,
	0x56, 0x1a, 0xeb, 0xe5, 0xde, 0x49, 0xd3, 0xb4, 0x1f, 0x25, 0x6f, 0xb5, 0x0d, 0xba, 0x9c, 0x56,
	0xe4, 0x2b, 0xd0, 0xfb, 0xb8, 0x83, 0x56, 0x79, 0xeb, 0x4a, 0xfe, 0xc9, 0x7d, 0x50, 0xc1, 0xf8,
	0x60, 0x0a, 0x9f, 0x0f, 0x73, 0x08, 0xfa, 0x65, 0x09, 0xef, 0x1e, 0xad, 0x20, 0x3f, 0xdf, 0xfe,
	0xb8, 0xcb, 0x16, 0x30, 0xeb, 0x82, 0x65, 0xad, 0xb6, 0x28, 0x46, 0x17, 0xa3, 0xcb, 0xe9, 0x2a,
	0xe9, 0xd8, 0x57, 0xb5, 0xc5, 0xec, 0x0c, 0x26, 0xad, 0x2b, 0xa3, 0x7c, 0xc4, 0xf2, 0xb8, 0x75,
	0x25, 0x49, 0xcb, 0x3f, 0x27, 0xf0, 0xe2, 0xc1, 0x20, 0xca, 0xe4, 0x05, 0x7a, 0xe9, 0xb1, 0x0e,
	0x9c, 0x99, 0xae, 0x92, 0x8e, 0xdd, 0x61, 0x1d, 0xb2, 0x57, 0x90, 0x38, 0xbb, 0x0b, 0xbd, 0xe3,
	0x88, 0x1d, 0x10, 0x11, 0x1b, 0xde, 0xc2, 0xbc, 0x33, 0x34, 0x58, 0x6b, 0x53, 0xff, 0x14, 0xc7,
	0xec, 0x49, 0x23, 0xfd, 0x16, 0x61, 0xf6, 0x12, 0x60, 0xab, 0xf6, 0x32, 0x42, 0x71, 0xc2, 0x96,
	0xe9, 0x56, 0xed, 0x57, 0x0c, 0x48, 0x6e, 0x54, 0xd8, 0x74, 0x53, 0x9e, 0x45, 0x99, 0x09, 0x0f,
	0x79, 0x0d, 0x69, 0x94, 0xfb, 0x19, 0xa7, 0xec, 0x98, 0x31, 0xec, 0x47, 0x9c, 0x43, 0xdc, 0x21,
	0x95, 0xd6, 0x62, 0xcc, 0x86, 0x09, 0x83, 0x6b, 0xad, 0xb3, 0x37, 0x30, 0xbf, 0x17, 0x25, 0x3a,
	0xe7, 0xc5, 0xe4, 0x20, 0xe2, 0x5a, 0xeb, 0x5b, 0xe7, 0xfc, 0x10, 0xa1, 0xb1, 0x12, 0xd3, 0x83,
	0x88, 0x1b, 0xac, 0x86, 0x08, 0x8d, 0x55, 0x8c, 0x80, 0x83, 0x88, 0x1b, 0xac, 0x38, 0x62, 0x01,
	0xb3, 0xaa, 0x54, 0x43, 0xd3, 0x24, 0x9e, 0x29, 0xb1, 0xbe, 0xe8, 0x19, 0x4c, 0xd8, 0x42, 0x3d,
	0x67, 0x2c, 0x8f, 0xe9, 0x9b, 0x6a, 0xf6, 0x12, 0xcd, 0x4f, 0x07, 0x89, 0xc6, 0x9f, 0xc3, 0x94,
	0x25, 0x3e, 0xa1, 0x79, 0xec, 0x46, 0xa0, 0xff, 0x17, 0xee, 0x6f, 0x4d, 0x85, 0x2d, 0x56, 0xe2,
	0x39, 0x5f, 0x80, 0xb4, 0xa7, 0x5f, 0x08, 0x2e, 0xff, 0x8d, 0x60, 0xf1, 0xe8, 0x7d, 0xcb, 0xfe,
	0x8e, 0x0e, 0xd2, 0xd8, 0x22, 0xae, 0x2e, 0x8e, 0x2f, 0x93, 0x2b, 0x93, 0x3f, 0xcd, 0x33, 0xc8,
	0x1f, 0x74, 0x1a, 0x8a, 0x7f, 0xa2, 0xcf, 0xf5, 0x29, 0x3f, 0xc3, 0x0f, 0xff, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0x99, 0xe6, 0x7f, 0xb1, 0x03, 0x00, 0x00,
}
