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
// source: fs_mgr_nlri.proto

package cisco_ios_xr_flowspec_oper_flow_spec_vrfs_vrf_afs_af_nlris_nlri

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

type FsMgrNlri_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	NlriBytes            string   `protobuf:"bytes,3,opt,name=nlri_bytes,json=nlriBytes,proto3" json:"nlri_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrNlri_KEYS) Reset()         { *m = FsMgrNlri_KEYS{} }
func (m *FsMgrNlri_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsMgrNlri_KEYS) ProtoMessage()    {}
func (*FsMgrNlri_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e1a81ecaa011991, []int{0}
}

func (m *FsMgrNlri_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrNlri_KEYS.Unmarshal(m, b)
}
func (m *FsMgrNlri_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrNlri_KEYS.Marshal(b, m, deterministic)
}
func (m *FsMgrNlri_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrNlri_KEYS.Merge(m, src)
}
func (m *FsMgrNlri_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsMgrNlri_KEYS.Size(m)
}
func (m *FsMgrNlri_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrNlri_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrNlri_KEYS proto.InternalMessageInfo

func (m *FsMgrNlri_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FsMgrNlri_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *FsMgrNlri_KEYS) GetNlriBytes() string {
	if m != nil {
		return m.NlriBytes
	}
	return ""
}

type FsMgrFlowStatsCtrs struct {
	Packets              uint64   `protobuf:"varint,1,opt,name=packets,proto3" json:"packets,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrFlowStatsCtrs) Reset()         { *m = FsMgrFlowStatsCtrs{} }
func (m *FsMgrFlowStatsCtrs) String() string { return proto.CompactTextString(m) }
func (*FsMgrFlowStatsCtrs) ProtoMessage()    {}
func (*FsMgrFlowStatsCtrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e1a81ecaa011991, []int{1}
}

func (m *FsMgrFlowStatsCtrs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrFlowStatsCtrs.Unmarshal(m, b)
}
func (m *FsMgrFlowStatsCtrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrFlowStatsCtrs.Marshal(b, m, deterministic)
}
func (m *FsMgrFlowStatsCtrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrFlowStatsCtrs.Merge(m, src)
}
func (m *FsMgrFlowStatsCtrs) XXX_Size() int {
	return xxx_messageInfo_FsMgrFlowStatsCtrs.Size(m)
}
func (m *FsMgrFlowStatsCtrs) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrFlowStatsCtrs.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrFlowStatsCtrs proto.InternalMessageInfo

func (m *FsMgrFlowStatsCtrs) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *FsMgrFlowStatsCtrs) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

type FsMgrFlowStats struct {
	Classified           *FsMgrFlowStatsCtrs `protobuf:"bytes,1,opt,name=classified,proto3" json:"classified,omitempty"`
	Dropped              *FsMgrFlowStatsCtrs `protobuf:"bytes,2,opt,name=dropped,proto3" json:"dropped,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FsMgrFlowStats) Reset()         { *m = FsMgrFlowStats{} }
func (m *FsMgrFlowStats) String() string { return proto.CompactTextString(m) }
func (*FsMgrFlowStats) ProtoMessage()    {}
func (*FsMgrFlowStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e1a81ecaa011991, []int{2}
}

func (m *FsMgrFlowStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrFlowStats.Unmarshal(m, b)
}
func (m *FsMgrFlowStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrFlowStats.Marshal(b, m, deterministic)
}
func (m *FsMgrFlowStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrFlowStats.Merge(m, src)
}
func (m *FsMgrFlowStats) XXX_Size() int {
	return xxx_messageInfo_FsMgrFlowStats.Size(m)
}
func (m *FsMgrFlowStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrFlowStats.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrFlowStats proto.InternalMessageInfo

func (m *FsMgrFlowStats) GetClassified() *FsMgrFlowStatsCtrs {
	if m != nil {
		return m.Classified
	}
	return nil
}

func (m *FsMgrFlowStats) GetDropped() *FsMgrFlowStatsCtrs {
	if m != nil {
		return m.Dropped
	}
	return nil
}

type FsMgrNlri struct {
	FlowStatistics       *FsMgrFlowStats `protobuf:"bytes,50,opt,name=flow_statistics,json=flowStatistics,proto3" json:"flow_statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FsMgrNlri) Reset()         { *m = FsMgrNlri{} }
func (m *FsMgrNlri) String() string { return proto.CompactTextString(m) }
func (*FsMgrNlri) ProtoMessage()    {}
func (*FsMgrNlri) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e1a81ecaa011991, []int{3}
}

func (m *FsMgrNlri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrNlri.Unmarshal(m, b)
}
func (m *FsMgrNlri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrNlri.Marshal(b, m, deterministic)
}
func (m *FsMgrNlri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrNlri.Merge(m, src)
}
func (m *FsMgrNlri) XXX_Size() int {
	return xxx_messageInfo_FsMgrNlri.Size(m)
}
func (m *FsMgrNlri) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrNlri.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrNlri proto.InternalMessageInfo

func (m *FsMgrNlri) GetFlowStatistics() *FsMgrFlowStats {
	if m != nil {
		return m.FlowStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FsMgrNlri_KEYS)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.nlris.nlri.fs_mgr_nlri_KEYS")
	proto.RegisterType((*FsMgrFlowStatsCtrs)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.nlris.nlri.fs_mgr_flow_stats_ctrs")
	proto.RegisterType((*FsMgrFlowStats)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.nlris.nlri.fs_mgr_flow_stats")
	proto.RegisterType((*FsMgrNlri)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.nlris.nlri.fs_mgr_nlri")
}

func init() { proto.RegisterFile("fs_mgr_nlri.proto", fileDescriptor_9e1a81ecaa011991) }

var fileDescriptor_9e1a81ecaa011991 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x14, 0x64, 0xab, 0x76, 0xed, 0x2b, 0xf8, 0x11, 0x44, 0xd7, 0x83, 0x20, 0x7b, 0xf2, 0x94, 0x43,
	0xfd, 0x01, 0x82, 0x20, 0x08, 0x82, 0x87, 0xf4, 0x20, 0x9e, 0x42, 0x9a, 0x4d, 0x24, 0xb8, 0xdb,
	0xc4, 0xbc, 0xb0, 0x55, 0xfc, 0x07, 0xfe, 0x69, 0x25, 0xd9, 0x76, 0xdd, 0x83, 0x37, 0xf5, 0xf2,
	0xc8, 0xcc, 0x84, 0x99, 0xe1, 0x25, 0x70, 0xa8, 0x91, 0x37, 0x4f, 0x9e, 0x2f, 0x6b, 0x6f, 0xa8,
	0xf3, 0x36, 0x58, 0x72, 0x25, 0x0d, 0x4a, 0xcb, 0x8d, 0x45, 0xfe, 0xea, 0xb9, 0xae, 0xed, 0x0a,
	0x9d, 0x92, 0xdc, 0x3a, 0xe5, 0x69, 0x44, 0x3c, 0x42, 0xda, 0x7a, 0x8d, 0x71, 0x50, 0xa1, 0x91,
	0x0a, 0x4d, 0xa3, 0x01, 0xa6, 0x59, 0x2a, 0x38, 0x18, 0xb8, 0xf2, 0xbb, 0x9b, 0xc7, 0x39, 0x39,
	0x85, 0xdd, 0xd6, 0x6b, 0xbe, 0x14, 0x8d, 0x2a, 0xb2, 0xf3, 0xec, 0x62, 0xc2, 0xf2, 0xd6, 0xeb,
	0x7b, 0xd1, 0x28, 0x72, 0x02, 0xb9, 0x58, 0x2b, 0xa3, 0xa4, 0x8c, 0x45, 0x27, 0x9c, 0x01, 0x24,
	0x83, 0xc5, 0x5b, 0x50, 0x58, 0x6c, 0x25, 0x6d, 0x12, 0x99, 0xeb, 0x48, 0x94, 0xb7, 0x70, 0xbc,
	0x8e, 0xe9, 0x5a, 0x05, 0x11, 0x90, 0xcb, 0xe0, 0x91, 0x14, 0x90, 0x3b, 0x21, 0x9f, 0x55, 0xc0,
	0x94, 0xb5, 0xcd, 0x36, 0x90, 0x1c, 0xc1, 0x4e, 0xe7, 0x36, 0x4a, 0x7c, 0x07, 0xca, 0xcf, 0xac,
	0xdf, 0xc3, 0xb7, 0x15, 0x59, 0x01, 0xc8, 0x5a, 0x20, 0x1a, 0x6d, 0x54, 0x95, 0x8c, 0xa6, 0xb3,
	0x07, 0xfa, 0xcb, 0xe5, 0xd0, 0x9f, 0x2b, 0xb3, 0x41, 0x14, 0x79, 0x81, 0xbc, 0xf2, 0xd6, 0x39,
	0x55, 0xa5, 0x9a, 0xff, 0x98, 0xba, 0xc9, 0x29, 0x3f, 0x32, 0x98, 0x0e, 0xde, 0x8c, 0xbc, 0xc3,
	0x7e, 0x7f, 0xd7, 0x60, 0x30, 0x12, 0x8b, 0x59, 0xaa, 0xc2, 0xfe, 0xbe, 0x0a, 0xdb, 0x8b, 0xe7,
	0x79, 0x9f, 0xb4, 0x18, 0xa7, 0x7f, 0x78, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x19, 0x66, 0x79,
	0x58, 0x9c, 0x02, 0x00, 0x00,
}
