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
// source: pce_path_list_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_cspf_cspf_paths_cspf_path

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

type PcePathListBag_KEYS struct {
	Af                   uint32   `protobuf:"varint,1,opt,name=af,proto3" json:"af,omitempty"`
	Source1              string   `protobuf:"bytes,2,opt,name=source1,proto3" json:"source1,omitempty"`
	Destination1         string   `protobuf:"bytes,3,opt,name=destination1,proto3" json:"destination1,omitempty"`
	MetricType           uint32   `protobuf:"varint,4,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
	Source2              string   `protobuf:"bytes,5,opt,name=source2,proto3" json:"source2,omitempty"`
	Destination2         string   `protobuf:"bytes,6,opt,name=destination2,proto3" json:"destination2,omitempty"`
	DisjointLevel        uint32   `protobuf:"varint,7,opt,name=disjoint_level,json=disjointLevel,proto3" json:"disjoint_level,omitempty"`
	DisjointStrict       uint32   `protobuf:"varint,8,opt,name=disjoint_strict,json=disjointStrict,proto3" json:"disjoint_strict,omitempty"`
	ShortestPath         uint32   `protobuf:"varint,9,opt,name=shortest_path,json=shortestPath,proto3" json:"shortest_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePathListBag_KEYS) Reset()         { *m = PcePathListBag_KEYS{} }
func (m *PcePathListBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcePathListBag_KEYS) ProtoMessage()    {}
func (*PcePathListBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_38073e5b05501788, []int{0}
}

func (m *PcePathListBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePathListBag_KEYS.Unmarshal(m, b)
}
func (m *PcePathListBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePathListBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcePathListBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePathListBag_KEYS.Merge(m, src)
}
func (m *PcePathListBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcePathListBag_KEYS.Size(m)
}
func (m *PcePathListBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePathListBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcePathListBag_KEYS proto.InternalMessageInfo

func (m *PcePathListBag_KEYS) GetAf() uint32 {
	if m != nil {
		return m.Af
	}
	return 0
}

func (m *PcePathListBag_KEYS) GetSource1() string {
	if m != nil {
		return m.Source1
	}
	return ""
}

func (m *PcePathListBag_KEYS) GetDestination1() string {
	if m != nil {
		return m.Destination1
	}
	return ""
}

func (m *PcePathListBag_KEYS) GetMetricType() uint32 {
	if m != nil {
		return m.MetricType
	}
	return 0
}

func (m *PcePathListBag_KEYS) GetSource2() string {
	if m != nil {
		return m.Source2
	}
	return ""
}

func (m *PcePathListBag_KEYS) GetDestination2() string {
	if m != nil {
		return m.Destination2
	}
	return ""
}

func (m *PcePathListBag_KEYS) GetDisjointLevel() uint32 {
	if m != nil {
		return m.DisjointLevel
	}
	return 0
}

func (m *PcePathListBag_KEYS) GetDisjointStrict() uint32 {
	if m != nil {
		return m.DisjointStrict
	}
	return 0
}

func (m *PcePathListBag_KEYS) GetShortestPath() uint32 {
	if m != nil {
		return m.ShortestPath
	}
	return 0
}

type PceAddrBag struct {
	AddressFamily        uint32   `protobuf:"varint,1,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	Ipv4Prefix           string   `protobuf:"bytes,2,opt,name=ipv4_prefix,json=ipv4Prefix,proto3" json:"ipv4_prefix,omitempty"`
	Ipv6Prefix           string   `protobuf:"bytes,3,opt,name=ipv6_prefix,json=ipv6Prefix,proto3" json:"ipv6_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceAddrBag) Reset()         { *m = PceAddrBag{} }
func (m *PceAddrBag) String() string { return proto.CompactTextString(m) }
func (*PceAddrBag) ProtoMessage()    {}
func (*PceAddrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_38073e5b05501788, []int{1}
}

func (m *PceAddrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceAddrBag.Unmarshal(m, b)
}
func (m *PceAddrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceAddrBag.Marshal(b, m, deterministic)
}
func (m *PceAddrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceAddrBag.Merge(m, src)
}
func (m *PceAddrBag) XXX_Size() int {
	return xxx_messageInfo_PceAddrBag.Size(m)
}
func (m *PceAddrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceAddrBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceAddrBag proto.InternalMessageInfo

func (m *PceAddrBag) GetAddressFamily() uint32 {
	if m != nil {
		return m.AddressFamily
	}
	return 0
}

func (m *PceAddrBag) GetIpv4Prefix() string {
	if m != nil {
		return m.Ipv4Prefix
	}
	return ""
}

func (m *PceAddrBag) GetIpv6Prefix() string {
	if m != nil {
		return m.Ipv6Prefix
	}
	return ""
}

type PceIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIpAddrType) Reset()         { *m = PceIpAddrType{} }
func (m *PceIpAddrType) String() string { return proto.CompactTextString(m) }
func (*PceIpAddrType) ProtoMessage()    {}
func (*PceIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_38073e5b05501788, []int{2}
}

func (m *PceIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIpAddrType.Unmarshal(m, b)
}
func (m *PceIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIpAddrType.Marshal(b, m, deterministic)
}
func (m *PceIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIpAddrType.Merge(m, src)
}
func (m *PceIpAddrType) XXX_Size() int {
	return xxx_messageInfo_PceIpAddrType.Size(m)
}
func (m *PceIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIpAddrType proto.InternalMessageInfo

func (m *PceIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PceIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PceIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type PcePathBag struct {
	Hops                 []*PceAddrBag  `protobuf:"bytes,1,rep,name=hops,proto3" json:"hops,omitempty"`
	Cost                 uint64         `protobuf:"varint,2,opt,name=cost,proto3" json:"cost,omitempty"`
	Source               *PceIpAddrType `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *PceIpAddrType `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PcePathBag) Reset()         { *m = PcePathBag{} }
func (m *PcePathBag) String() string { return proto.CompactTextString(m) }
func (*PcePathBag) ProtoMessage()    {}
func (*PcePathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_38073e5b05501788, []int{3}
}

func (m *PcePathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePathBag.Unmarshal(m, b)
}
func (m *PcePathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePathBag.Marshal(b, m, deterministic)
}
func (m *PcePathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePathBag.Merge(m, src)
}
func (m *PcePathBag) XXX_Size() int {
	return xxx_messageInfo_PcePathBag.Size(m)
}
func (m *PcePathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePathBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePathBag proto.InternalMessageInfo

func (m *PcePathBag) GetHops() []*PceAddrBag {
	if m != nil {
		return m.Hops
	}
	return nil
}

func (m *PcePathBag) GetCost() uint64 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *PcePathBag) GetSource() *PceIpAddrType {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PcePathBag) GetDestination() *PceIpAddrType {
	if m != nil {
		return m.Destination
	}
	return nil
}

type PcePathListBag struct {
	OutputPath           []*PcePathBag `protobuf:"bytes,50,rep,name=output_path,json=outputPath,proto3" json:"output_path,omitempty"`
	HeadendsSwapped      string        `protobuf:"bytes,51,opt,name=headends_swapped,json=headendsSwapped,proto3" json:"headends_swapped,omitempty"`
	CspfResult           string        `protobuf:"bytes,52,opt,name=cspf_result,json=cspfResult,proto3" json:"cspf_result,omitempty"`
	IterationsDone       uint32        `protobuf:"varint,53,opt,name=iterations_done,json=iterationsDone,proto3" json:"iterations_done,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PcePathListBag) Reset()         { *m = PcePathListBag{} }
func (m *PcePathListBag) String() string { return proto.CompactTextString(m) }
func (*PcePathListBag) ProtoMessage()    {}
func (*PcePathListBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_38073e5b05501788, []int{4}
}

func (m *PcePathListBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePathListBag.Unmarshal(m, b)
}
func (m *PcePathListBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePathListBag.Marshal(b, m, deterministic)
}
func (m *PcePathListBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePathListBag.Merge(m, src)
}
func (m *PcePathListBag) XXX_Size() int {
	return xxx_messageInfo_PcePathListBag.Size(m)
}
func (m *PcePathListBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePathListBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePathListBag proto.InternalMessageInfo

func (m *PcePathListBag) GetOutputPath() []*PcePathBag {
	if m != nil {
		return m.OutputPath
	}
	return nil
}

func (m *PcePathListBag) GetHeadendsSwapped() string {
	if m != nil {
		return m.HeadendsSwapped
	}
	return ""
}

func (m *PcePathListBag) GetCspfResult() string {
	if m != nil {
		return m.CspfResult
	}
	return ""
}

func (m *PcePathListBag) GetIterationsDone() uint32 {
	if m != nil {
		return m.IterationsDone
	}
	return 0
}

func init() {
	proto.RegisterType((*PcePathListBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_list_bag_KEYS")
	proto.RegisterType((*PceAddrBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_addr_bag")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_ip_addr_type")
	proto.RegisterType((*PcePathBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_bag")
	proto.RegisterType((*PcePathListBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_list_bag")
}

func init() { proto.RegisterFile("pce_path_list_bag.proto", fileDescriptor_38073e5b05501788) }

var fileDescriptor_38073e5b05501788 = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8e, 0x52, 0x31,
	0x14, 0xc6, 0xc3, 0x1f, 0x41, 0x0e, 0x30, 0x33, 0x76, 0xe1, 0x74, 0x27, 0xc1, 0x18, 0x71, 0x43,
	0x32, 0x77, 0x46, 0x12, 0xf7, 0xfe, 0x59, 0x8c, 0x31, 0x93, 0x8b, 0x1b, 0xe3, 0xa2, 0xe9, 0xf4,
	0x9e, 0x2b, 0x35, 0x70, 0xdb, 0xb4, 0x65, 0x06, 0x5e, 0xcd, 0xd7, 0xf0, 0x4d, 0x7c, 0x02, 0xd3,
	0x73, 0xef, 0x05, 0x94, 0x9d, 0xce, 0x86, 0x1c, 0x7e, 0x9c, 0x7e, 0xa7, 0xfd, 0xfa, 0x15, 0x38,
	0xb7, 0x0a, 0x85, 0x95, 0x61, 0x21, 0x96, 0xda, 0x07, 0x71, 0x2b, 0xbf, 0x4d, 0xad, 0x33, 0xc1,
	0xb0, 0x37, 0x4a, 0x7b, 0x65, 0x84, 0x36, 0x5e, 0x6c, 0x9c, 0xd0, 0x45, 0xee, 0xa4, 0xd8, 0x04,
	0x25, 0x8c, 0x45, 0x37, 0xb5, 0x0a, 0xa7, 0xca, 0xdb, 0x9c, 0x3e, 0x68, 0xb9, 0xdf, 0x97, 0xe3,
	0x1f, 0x4d, 0x78, 0x7a, 0x24, 0x2b, 0xae, 0xdf, 0x7d, 0x99, 0xb3, 0x13, 0x68, 0xca, 0x9c, 0x37,
	0x46, 0x8d, 0xc9, 0x30, 0x6d, 0xca, 0x9c, 0x71, 0xe8, 0x7a, 0xb3, 0x76, 0x0a, 0x2f, 0x78, 0x73,
	0xd4, 0x98, 0xf4, 0xd2, 0xfa, 0x2b, 0x1b, 0xc3, 0x20, 0x43, 0x1f, 0x74, 0x21, 0x83, 0x36, 0xc5,
	0x05, 0x6f, 0xd1, 0xcf, 0x7f, 0x30, 0xf6, 0x0c, 0xfa, 0x2b, 0x0c, 0x4e, 0x2b, 0x11, 0xb6, 0x16,
	0x79, 0x9b, 0x64, 0xa1, 0x44, 0x9f, 0xb7, 0x16, 0xf7, 0xf2, 0x09, 0x7f, 0x74, 0x28, 0x9f, 0xfc,
	0x25, 0x9f, 0xf0, 0xce, 0x91, 0x7c, 0xc2, 0x5e, 0xc0, 0x49, 0xa6, 0xfd, 0x77, 0xa3, 0x8b, 0x20,
	0x96, 0x78, 0x87, 0x4b, 0xde, 0xa5, 0x09, 0xc3, 0x9a, 0x7e, 0x8c, 0x90, 0xbd, 0x84, 0xd3, 0x5d,
	0x9b, 0x8f, 0xa3, 0x03, 0x7f, 0x4c, 0x7d, 0xbb, 0xd5, 0x73, 0xa2, 0xec, 0x39, 0x0c, 0xfd, 0xc2,
	0xb8, 0x80, 0x3e, 0x90, 0x37, 0xbc, 0x47, 0x6d, 0x83, 0x1a, 0xde, 0x44, 0xf3, 0xee, 0x61, 0x10,
	0xbd, 0x93, 0x59, 0xe6, 0xa2, 0x6d, 0x71, 0x13, 0xb1, 0x46, 0xef, 0x45, 0x2e, 0x57, 0x7a, 0xb9,
	0xad, 0xdc, 0x1b, 0x56, 0xf4, 0x3d, 0xc1, 0x68, 0x85, 0xb6, 0x77, 0x57, 0xc2, 0x3a, 0xcc, 0xf5,
	0xa6, 0x32, 0x13, 0x22, 0xba, 0x21, 0x52, 0x35, 0xcc, 0xea, 0x86, 0xd6, 0xae, 0x61, 0x56, 0x36,
	0x8c, 0xe7, 0x70, 0x16, 0x07, 0x6b, 0x5b, 0xce, 0x8e, 0x8e, 0xb2, 0x73, 0xe8, 0xca, 0x5c, 0x14,
	0x72, 0x85, 0x34, 0xb5, 0x97, 0x76, 0x64, 0xfe, 0x49, 0xae, 0x90, 0x31, 0x68, 0x47, 0xed, 0x6a,
	0x0e, 0xd5, 0x15, 0x9b, 0x55, 0xd2, 0x54, 0x8f, 0x7f, 0x36, 0xcb, 0xe3, 0x50, 0x14, 0xe2, 0x71,
	0xbe, 0x42, 0x7b, 0x61, 0xac, 0xe7, 0x8d, 0x51, 0x6b, 0xd2, 0x4f, 0x3e, 0x4c, 0xff, 0x39, 0x65,
	0xd3, 0x43, 0x97, 0x52, 0x12, 0x8d, 0x3b, 0x50, 0xc6, 0x07, 0xda, 0x55, 0x3b, 0xa5, 0x9a, 0x29,
	0xe8, 0x94, 0x77, 0x4e, 0xfb, 0xea, 0x27, 0xd7, 0xff, 0x39, 0xf2, 0xd0, 0x9f, 0xb4, 0x92, 0x66,
	0x2b, 0xe8, 0x1f, 0x24, 0x87, 0x82, 0xf8, 0xc0, 0x93, 0x0e, 0xf5, 0xc7, 0xbf, 0x1a, 0xf0, 0xe4,
	0xe8, 0x81, 0xb1, 0x05, 0xf4, 0xcd, 0x3a, 0xd8, 0x75, 0x15, 0xae, 0xe4, 0x41, 0x1c, 0xae, 0x2f,
	0x2e, 0x85, 0x52, 0x3b, 0x66, 0x94, 0xbd, 0x82, 0xb3, 0x05, 0xca, 0x0c, 0x8b, 0xcc, 0x0b, 0x7f,
	0x2f, 0xad, 0xc5, 0x8c, 0x5f, 0xd2, 0xad, 0x9f, 0xd6, 0x7c, 0x5e, 0xe2, 0x18, 0x3b, 0x12, 0x74,
	0xe8, 0xd7, 0xcb, 0xc0, 0xaf, 0xca, 0xd8, 0x45, 0x94, 0x12, 0x89, 0xaf, 0x47, 0x07, 0x74, 0x74,
	0x30, 0x2f, 0x32, 0x53, 0x20, 0x7f, 0x5d, 0xbe, 0x9e, 0x3d, 0x7e, 0x6b, 0x0a, 0xbc, 0xed, 0xd0,
	0xff, 0xd2, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x01, 0x7f, 0x2b, 0x49, 0xb2, 0x04, 0x00,
	0x00,
}
