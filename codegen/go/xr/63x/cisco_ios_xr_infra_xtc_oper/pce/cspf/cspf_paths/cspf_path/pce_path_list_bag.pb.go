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

func init() {
	proto.RegisterType((*PcePathListBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_list_bag_KEYS")
	proto.RegisterType((*PceAddrBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_addr_bag")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_ip_addr_type")
	proto.RegisterType((*PcePathBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_bag")
	proto.RegisterType((*PcePathListBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.cspf.cspf_paths.cspf_path.pce_path_list_bag")
}

func init() { proto.RegisterFile("pce_path_list_bag.proto", fileDescriptor_38073e5b05501788) }

var fileDescriptor_38073e5b05501788 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd5, 0x3f, 0xb4, 0x74, 0xda, 0xee, 0x2e, 0x3e, 0xb0, 0xbe, 0x51, 0x15, 0x21, 0xca,
	0x25, 0xd2, 0x66, 0x57, 0x95, 0x78, 0x00, 0xe0, 0xb0, 0x08, 0xad, 0x52, 0x2e, 0x88, 0x83, 0xe5,
	0x75, 0x1c, 0x62, 0x94, 0xc4, 0x96, 0xed, 0xee, 0xb6, 0xaf, 0xc6, 0x1b, 0x70, 0xe6, 0x85, 0x90,
	0x27, 0x49, 0x37, 0xd0, 0x1b, 0xec, 0xa5, 0x9a, 0xfe, 0x3a, 0xf9, 0x66, 0xfc, 0xf9, 0x6b, 0xe0,
	0xdc, 0x08, 0xc9, 0x0c, 0xf7, 0x39, 0x2b, 0x94, 0xf3, 0xec, 0x96, 0x7f, 0x8b, 0x8c, 0xd5, 0x5e,
	0x93, 0xb7, 0x42, 0x39, 0xa1, 0x99, 0xd2, 0x8e, 0xed, 0x2c, 0x53, 0x55, 0x66, 0x39, 0xdb, 0x79,
	0xc1, 0xb4, 0x91, 0x36, 0x32, 0x42, 0x46, 0xc2, 0x99, 0x0c, 0x3f, 0xf0, 0x71, 0xf7, 0x50, 0x2e,
	0x7f, 0xf4, 0xe1, 0xf9, 0x91, 0x2c, 0xbb, 0x7e, 0xf7, 0x65, 0x43, 0x4e, 0xa0, 0xcf, 0x33, 0xda,
	0x5b, 0xf4, 0x56, 0xf3, 0xa4, 0xcf, 0x33, 0x42, 0x61, 0xec, 0xf4, 0xd6, 0x0a, 0x79, 0x41, 0xfb,
	0x8b, 0xde, 0x6a, 0x92, 0xb4, 0x5f, 0xc9, 0x12, 0x66, 0xa9, 0x74, 0x5e, 0x55, 0xdc, 0x2b, 0x5d,
	0x5d, 0xd0, 0x01, 0xfe, 0xfc, 0x07, 0x23, 0x2f, 0x60, 0x5a, 0x4a, 0x6f, 0x95, 0x60, 0x7e, 0x6f,
	0x24, 0x1d, 0xa2, 0x2c, 0xd4, 0xe8, 0xf3, 0xde, 0xc8, 0x07, 0xf9, 0x98, 0x3e, 0xe9, 0xca, 0xc7,
	0x7f, 0xc9, 0xc7, 0x74, 0x74, 0x24, 0x1f, 0x93, 0x57, 0x70, 0x92, 0x2a, 0xf7, 0x5d, 0xab, 0xca,
	0xb3, 0x42, 0xde, 0xc9, 0x82, 0x8e, 0x71, 0xc2, 0xbc, 0xa5, 0x1f, 0x03, 0x24, 0xaf, 0xe1, 0xf4,
	0xd0, 0xe6, 0xc2, 0x68, 0x4f, 0x9f, 0x62, 0xdf, 0xe1, 0xe9, 0x0d, 0x52, 0xf2, 0x12, 0xe6, 0x2e,
	0xd7, 0xd6, 0x4b, 0xe7, 0xd1, 0x1b, 0x3a, 0xc1, 0xb6, 0x59, 0x0b, 0x6f, 0x82, 0x79, 0xf7, 0x30,
	0x0b, 0xde, 0xf1, 0x34, 0xb5, 0xc1, 0xb6, 0xb0, 0x44, 0xa8, 0xa5, 0x73, 0x2c, 0xe3, 0xa5, 0x2a,
	0xf6, 0x8d, 0x7b, 0xf3, 0x86, 0xbe, 0x47, 0x18, 0xac, 0x50, 0xe6, 0xee, 0x8a, 0x19, 0x2b, 0x33,
	0xb5, 0x6b, 0xcc, 0x84, 0x80, 0x6e, 0x90, 0x34, 0x0d, 0xeb, 0xb6, 0x61, 0x70, 0x68, 0x58, 0xd7,
	0x0d, 0xcb, 0x0d, 0x9c, 0x85, 0xc1, 0xca, 0xd4, 0xb3, 0x83, 0xa3, 0xe4, 0x1c, 0xc6, 0x3c, 0x63,
	0x15, 0x2f, 0x25, 0x4e, 0x9d, 0x24, 0x23, 0x9e, 0x7d, 0xe2, 0xa5, 0x24, 0x04, 0x86, 0x41, 0xbb,
	0x99, 0x83, 0x75, 0xc3, 0xd6, 0x8d, 0x34, 0xd6, 0xcb, 0x5f, 0xfd, 0xfa, 0x38, 0x18, 0x85, 0x70,
	0x9c, 0xaf, 0x30, 0xcc, 0xb5, 0x71, 0xb4, 0xb7, 0x18, 0xac, 0xa6, 0xf1, 0x87, 0xe8, 0x9f, 0x53,
	0x16, 0x75, 0x5d, 0x4a, 0x50, 0x34, 0x6c, 0x20, 0xb4, 0xf3, 0xb8, 0xd5, 0x30, 0xc1, 0x9a, 0x08,
	0x18, 0xd5, 0x77, 0x8e, 0x7b, 0x4d, 0xe3, 0xeb, 0xff, 0x1c, 0xd9, 0xf5, 0x27, 0x69, 0xa4, 0x49,
	0x09, 0xd3, 0x4e, 0x72, 0x30, 0x88, 0x8f, 0x3c, 0xa9, 0xab, 0xbf, 0xfc, 0xd9, 0x83, 0x67, 0x47,
	0x7f, 0x30, 0x92, 0xc3, 0x54, 0x6f, 0xbd, 0xd9, 0x36, 0xe1, 0x8a, 0x1f, 0xc5, 0xe1, 0xf6, 0xe2,
	0x12, 0xa8, 0xb5, 0x43, 0x46, 0xc9, 0x1b, 0x38, 0xcb, 0x25, 0x4f, 0x65, 0x95, 0x3a, 0xe6, 0xee,
	0xb9, 0x31, 0x32, 0xa5, 0x97, 0x78, 0xeb, 0xa7, 0x2d, 0xdf, 0xd4, 0x38, 0xc4, 0x0e, 0x05, 0xad,
	0x74, 0xdb, 0xc2, 0xd3, 0xab, 0x3a, 0x76, 0x01, 0x25, 0x48, 0x6e, 0x47, 0xf8, 0xba, 0xb9, 0xfc,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xf9, 0xe1, 0xbd, 0x89, 0x04, 0x00, 0x00,
}
