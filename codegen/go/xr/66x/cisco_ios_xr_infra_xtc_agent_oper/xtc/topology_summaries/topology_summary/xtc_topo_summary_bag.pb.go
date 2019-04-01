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
// source: xtc_topo_summary_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_topology_summaries_topology_summary

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

type XtcTopoSummaryBag_KEYS struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcTopoSummaryBag_KEYS) Reset()         { *m = XtcTopoSummaryBag_KEYS{} }
func (m *XtcTopoSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcTopoSummaryBag_KEYS) ProtoMessage()    {}
func (*XtcTopoSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a951e2c3239240d, []int{0}
}

func (m *XtcTopoSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcTopoSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *XtcTopoSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcTopoSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcTopoSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcTopoSummaryBag_KEYS.Merge(m, src)
}
func (m *XtcTopoSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcTopoSummaryBag_KEYS.Size(m)
}
func (m *XtcTopoSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcTopoSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcTopoSummaryBag_KEYS proto.InternalMessageInfo

func (m *XtcTopoSummaryBag_KEYS) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *XtcTopoSummaryBag_KEYS) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type XtcTopoSummaryBag struct {
	Nodes                uint32   `protobuf:"varint,50,opt,name=nodes,proto3" json:"nodes,omitempty"`
	Prefixes             uint32   `protobuf:"varint,51,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	PrefixSids           uint32   `protobuf:"varint,52,opt,name=prefix_sids,json=prefixSids,proto3" json:"prefix_sids,omitempty"`
	Links                uint32   `protobuf:"varint,53,opt,name=links,proto3" json:"links,omitempty"`
	AdjacencySids        uint32   `protobuf:"varint,54,opt,name=adjacency_sids,json=adjacencySids,proto3" json:"adjacency_sids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcTopoSummaryBag) Reset()         { *m = XtcTopoSummaryBag{} }
func (m *XtcTopoSummaryBag) String() string { return proto.CompactTextString(m) }
func (*XtcTopoSummaryBag) ProtoMessage()    {}
func (*XtcTopoSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a951e2c3239240d, []int{1}
}

func (m *XtcTopoSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcTopoSummaryBag.Unmarshal(m, b)
}
func (m *XtcTopoSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcTopoSummaryBag.Marshal(b, m, deterministic)
}
func (m *XtcTopoSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcTopoSummaryBag.Merge(m, src)
}
func (m *XtcTopoSummaryBag) XXX_Size() int {
	return xxx_messageInfo_XtcTopoSummaryBag.Size(m)
}
func (m *XtcTopoSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcTopoSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcTopoSummaryBag proto.InternalMessageInfo

func (m *XtcTopoSummaryBag) GetNodes() uint32 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

func (m *XtcTopoSummaryBag) GetPrefixes() uint32 {
	if m != nil {
		return m.Prefixes
	}
	return 0
}

func (m *XtcTopoSummaryBag) GetPrefixSids() uint32 {
	if m != nil {
		return m.PrefixSids
	}
	return 0
}

func (m *XtcTopoSummaryBag) GetLinks() uint32 {
	if m != nil {
		return m.Links
	}
	return 0
}

func (m *XtcTopoSummaryBag) GetAdjacencySids() uint32 {
	if m != nil {
		return m.AdjacencySids
	}
	return 0
}

func init() {
	proto.RegisterType((*XtcTopoSummaryBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_summaries.topology_summary.xtc_topo_summary_bag_KEYS")
	proto.RegisterType((*XtcTopoSummaryBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.topology_summaries.topology_summary.xtc_topo_summary_bag")
}

func init() { proto.RegisterFile("xtc_topo_summary_bag.proto", fileDescriptor_9a951e2c3239240d) }

var fileDescriptor_9a951e2c3239240d = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xd9, 0x05, 0x45, 0x47, 0xda, 0x43, 0xe8, 0x21, 0xf6, 0x62, 0x29, 0x08, 0x3d, 0xed,
	0xc1, 0xaa, 0xff, 0x40, 0x44, 0xbc, 0xb5, 0x27, 0x4f, 0x43, 0x9a, 0xcd, 0x2e, 0xd1, 0x6d, 0x66,
	0xc9, 0x44, 0xd8, 0xfd, 0x43, 0xfe, 0x4e, 0x49, 0x22, 0x7b, 0x90, 0xbd, 0xe5, 0xbd, 0x97, 0xf7,
	0xcd, 0x83, 0xf5, 0x10, 0x34, 0x06, 0xea, 0x09, 0xf9, 0xfb, 0x7c, 0x56, 0x7e, 0xc4, 0x93, 0x6a,
	0xab, 0xde, 0x53, 0x20, 0xf1, 0xa6, 0x2d, 0x6b, 0x42, 0x4b, 0x8c, 0x83, 0x47, 0xeb, 0x1a, 0xaf,
	0x30, 0x7e, 0x57, 0xad, 0x71, 0x01, 0xa9, 0x37, 0xbe, 0x1a, 0x82, 0xae, 0x62, 0xbb, 0xa3, 0x76,
	0xfc, 0x23, 0x58, 0xc3, 0xff, 0xad, 0x71, 0xfb, 0x0a, 0xb7, 0x73, 0x87, 0xf0, 0xfd, 0xe5, 0xe3,
	0x28, 0x96, 0x50, 0xaa, 0x46, 0x16, 0x9b, 0x62, 0x77, 0x7d, 0x28, 0x55, 0x23, 0xd6, 0x70, 0x95,
	0x06, 0x68, 0xea, 0x64, 0x99, 0xdc, 0x49, 0x6f, 0x7f, 0x0a, 0x58, 0xcd, 0x91, 0xc4, 0x0a, 0x2e,
	0x1c, 0xd5, 0x86, 0xe5, 0xc3, 0xa6, 0xd8, 0x2d, 0x0e, 0x59, 0x64, 0x94, 0x69, 0xec, 0x60, 0x58,
	0xee, 0x53, 0x30, 0x69, 0x71, 0x07, 0x37, 0xf9, 0x8d, 0x6c, 0x6b, 0x96, 0x8f, 0x29, 0x86, 0x6c,
	0x1d, 0x6d, 0xcd, 0x11, 0xd9, 0x59, 0xf7, 0xc5, 0xf2, 0x29, 0x23, 0x93, 0x10, 0xf7, 0xb0, 0x54,
	0xf5, 0xa7, 0xd2, 0xc6, 0xe9, 0x31, 0x37, 0x9f, 0x53, 0xbc, 0x98, 0xdc, 0x58, 0x3e, 0x5d, 0xa6,
	0xc9, 0xfb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xbb, 0x00, 0xf9, 0x61, 0x01, 0x00, 0x00,
}