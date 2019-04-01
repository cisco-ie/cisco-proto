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
// source: pim_mroute_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_mib_databases_mib_database

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

type PimMrouteBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         string   `protobuf:"bytes,2,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	SourceNetmask        uint32   `protobuf:"varint,3,opt,name=source_netmask,json=sourceNetmask,proto3" json:"source_netmask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimMrouteBag_KEYS) Reset()         { *m = PimMrouteBag_KEYS{} }
func (m *PimMrouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimMrouteBag_KEYS) ProtoMessage()    {}
func (*PimMrouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8098ca7299bdc29, []int{0}
}

func (m *PimMrouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMrouteBag_KEYS.Unmarshal(m, b)
}
func (m *PimMrouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMrouteBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimMrouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMrouteBag_KEYS.Merge(m, src)
}
func (m *PimMrouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimMrouteBag_KEYS.Size(m)
}
func (m *PimMrouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMrouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimMrouteBag_KEYS proto.InternalMessageInfo

func (m *PimMrouteBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimMrouteBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *PimMrouteBag_KEYS) GetSourceNetmask() uint32 {
	if m != nil {
		return m.SourceNetmask
	}
	return 0
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8098ca7299bdc29, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimMrouteBag struct {
	SourceAddressXr        *PimAddrtype `protobuf:"bytes,50,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	GroupAddressXr         *PimAddrtype `protobuf:"bytes,51,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	UpstreamAssertTimer    int32        `protobuf:"zigzag32,52,opt,name=upstream_assert_timer,json=upstreamAssertTimer,proto3" json:"upstream_assert_timer,omitempty"`
	AssertMetric           uint32       `protobuf:"varint,53,opt,name=assert_metric,json=assertMetric,proto3" json:"assert_metric,omitempty"`
	AssertMetricPreference uint32       `protobuf:"varint,54,opt,name=assert_metric_preference,json=assertMetricPreference,proto3" json:"assert_metric_preference,omitempty"`
	AssertRptBit           bool         `protobuf:"varint,55,opt,name=assert_rpt_bit,json=assertRptBit,proto3" json:"assert_rpt_bit,omitempty"`
	SptBit                 bool         `protobuf:"varint,56,opt,name=spt_bit,json=sptBit,proto3" json:"spt_bit,omitempty"`
	RpfNeighbor            *PimAddrtype `protobuf:"bytes,57,opt,name=rpf_neighbor,json=rpfNeighbor,proto3" json:"rpf_neighbor,omitempty"`
	RpfRoot                *PimAddrtype `protobuf:"bytes,58,opt,name=rpf_root,json=rpfRoot,proto3" json:"rpf_root,omitempty"`
	RpfMask                uint32       `protobuf:"varint,59,opt,name=rpf_mask,json=rpfMask,proto3" json:"rpf_mask,omitempty"`
	RpfSafi                uint32       `protobuf:"varint,60,opt,name=rpf_safi,json=rpfSafi,proto3" json:"rpf_safi,omitempty"`
	RpfTableName           string       `protobuf:"bytes,61,opt,name=rpf_table_name,json=rpfTableName,proto3" json:"rpf_table_name,omitempty"`
	RpfDrop                bool         `protobuf:"varint,62,opt,name=rpf_drop,json=rpfDrop,proto3" json:"rpf_drop,omitempty"`
	RpfExtranet            bool         `protobuf:"varint,63,opt,name=rpf_extranet,json=rpfExtranet,proto3" json:"rpf_extranet,omitempty"`
	RpfInterfaceName       string       `protobuf:"bytes,64,opt,name=rpf_interface_name,json=rpfInterfaceName,proto3" json:"rpf_interface_name,omitempty"`
	RpfVrfName             string       `protobuf:"bytes,65,opt,name=rpf_vrf_name,json=rpfVrfName,proto3" json:"rpf_vrf_name,omitempty"`
	BidirectionalRoute     bool         `protobuf:"varint,66,opt,name=bidirectional_route,json=bidirectionalRoute,proto3" json:"bidirectional_route,omitempty"`
	Uptime                 uint64       `protobuf:"varint,67,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Protocol               string       `protobuf:"bytes,68,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}     `json:"-"`
	XXX_unrecognized       []byte       `json:"-"`
	XXX_sizecache          int32        `json:"-"`
}

func (m *PimMrouteBag) Reset()         { *m = PimMrouteBag{} }
func (m *PimMrouteBag) String() string { return proto.CompactTextString(m) }
func (*PimMrouteBag) ProtoMessage()    {}
func (*PimMrouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8098ca7299bdc29, []int{2}
}

func (m *PimMrouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMrouteBag.Unmarshal(m, b)
}
func (m *PimMrouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMrouteBag.Marshal(b, m, deterministic)
}
func (m *PimMrouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMrouteBag.Merge(m, src)
}
func (m *PimMrouteBag) XXX_Size() int {
	return xxx_messageInfo_PimMrouteBag.Size(m)
}
func (m *PimMrouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMrouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimMrouteBag proto.InternalMessageInfo

func (m *PimMrouteBag) GetSourceAddressXr() *PimAddrtype {
	if m != nil {
		return m.SourceAddressXr
	}
	return nil
}

func (m *PimMrouteBag) GetGroupAddressXr() *PimAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *PimMrouteBag) GetUpstreamAssertTimer() int32 {
	if m != nil {
		return m.UpstreamAssertTimer
	}
	return 0
}

func (m *PimMrouteBag) GetAssertMetric() uint32 {
	if m != nil {
		return m.AssertMetric
	}
	return 0
}

func (m *PimMrouteBag) GetAssertMetricPreference() uint32 {
	if m != nil {
		return m.AssertMetricPreference
	}
	return 0
}

func (m *PimMrouteBag) GetAssertRptBit() bool {
	if m != nil {
		return m.AssertRptBit
	}
	return false
}

func (m *PimMrouteBag) GetSptBit() bool {
	if m != nil {
		return m.SptBit
	}
	return false
}

func (m *PimMrouteBag) GetRpfNeighbor() *PimAddrtype {
	if m != nil {
		return m.RpfNeighbor
	}
	return nil
}

func (m *PimMrouteBag) GetRpfRoot() *PimAddrtype {
	if m != nil {
		return m.RpfRoot
	}
	return nil
}

func (m *PimMrouteBag) GetRpfMask() uint32 {
	if m != nil {
		return m.RpfMask
	}
	return 0
}

func (m *PimMrouteBag) GetRpfSafi() uint32 {
	if m != nil {
		return m.RpfSafi
	}
	return 0
}

func (m *PimMrouteBag) GetRpfTableName() string {
	if m != nil {
		return m.RpfTableName
	}
	return ""
}

func (m *PimMrouteBag) GetRpfDrop() bool {
	if m != nil {
		return m.RpfDrop
	}
	return false
}

func (m *PimMrouteBag) GetRpfExtranet() bool {
	if m != nil {
		return m.RpfExtranet
	}
	return false
}

func (m *PimMrouteBag) GetRpfInterfaceName() string {
	if m != nil {
		return m.RpfInterfaceName
	}
	return ""
}

func (m *PimMrouteBag) GetRpfVrfName() string {
	if m != nil {
		return m.RpfVrfName
	}
	return ""
}

func (m *PimMrouteBag) GetBidirectionalRoute() bool {
	if m != nil {
		return m.BidirectionalRoute
	}
	return false
}

func (m *PimMrouteBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *PimMrouteBag) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*PimMrouteBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.mib_databases.mib_database.pim_mroute_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.mib_databases.mib_database.pim_addrtype")
	proto.RegisterType((*PimMrouteBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.mib_databases.mib_database.pim_mroute_bag")
}

func init() { proto.RegisterFile("pim_mroute_bag.proto", fileDescriptor_f8098ca7299bdc29) }

var fileDescriptor_f8098ca7299bdc29 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcf, 0x6f, 0xd4, 0x3a,
	0x10, 0x56, 0x5e, 0x9f, 0xb6, 0xfb, 0xdc, 0xed, 0xbe, 0xd6, 0x85, 0x62, 0x38, 0x85, 0x02, 0xd2,
	0x1e, 0x50, 0x90, 0xda, 0xb2, 0x94, 0xdf, 0xb4, 0xb4, 0x07, 0x84, 0x5a, 0xa1, 0xb4, 0xaa, 0xca,
	0xc9, 0x72, 0x92, 0x71, 0xb1, 0xba, 0x89, 0xad, 0xb1, 0xb7, 0xda, 0x9e, 0x38, 0xc2, 0x09, 0xf1,
	0x27, 0xf2, 0xa7, 0x20, 0x3b, 0xc9, 0xd2, 0xfc, 0x01, 0xf4, 0x16, 0x7f, 0xdf, 0x97, 0xf9, 0xc6,
	0x33, 0x9e, 0x21, 0xb7, 0x8c, 0x2a, 0x79, 0x89, 0x7a, 0xea, 0x80, 0x67, 0xe2, 0x3c, 0x31, 0xa8,
	0x9d, 0xa6, 0xa7, 0xb9, 0xb2, 0xb9, 0xe6, 0x4a, 0x5b, 0x3e, 0x43, 0xae, 0xcc, 0xe5, 0x36, 0xf7,
	0x3a, 0x6d, 0x00, 0x13, 0x65, 0x2e, 0xc7, 0xfe, 0x94, 0x58, 0x27, 0xaa, 0x22, 0xbb, 0x4a, 0x0a,
	0x90, 0x62, 0x3a, 0x71, 0x3c, 0xd7, 0x95, 0x83, 0x99, 0x4b, 0x4a, 0x95, 0xf1, 0x42, 0x38, 0x91,
	0x09, 0x0b, 0xb6, 0x73, 0xda, 0xf8, 0x1e, 0x91, 0xb5, 0xae, 0x21, 0xff, 0x78, 0xf0, 0xf9, 0x98,
	0x3e, 0x22, 0x43, 0xab, 0xa7, 0x98, 0x03, 0x17, 0x45, 0x81, 0x60, 0x2d, 0x8b, 0xe2, 0x68, 0xf4,
	0x5f, 0xba, 0x5c, 0xa3, 0xbb, 0x35, 0x48, 0x1f, 0x90, 0xe5, 0x73, 0xd4, 0x53, 0x33, 0x57, 0xfd,
	0x13, 0x54, 0x83, 0x00, 0xb6, 0xa2, 0x3f, 0xb1, 0x2a, 0x70, 0xa5, 0xb0, 0x17, 0x6c, 0x21, 0x8e,
	0x46, 0xcb, 0x6d, 0xac, 0xa3, 0x1a, 0xdc, 0x28, 0xc9, 0xc0, 0x67, 0xe2, 0x23, 0xb9, 0x2b, 0x03,
	0xf4, 0x0e, 0x59, 0x14, 0x92, 0x57, 0xa2, 0x84, 0xc6, 0xbb, 0x27, 0xe4, 0x91, 0x28, 0x81, 0xde,
	0x27, 0x83, 0x50, 0x80, 0xae, 0xe7, 0x92, 0xc7, 0x5a, 0xcb, 0x5a, 0x32, 0x9e, 0x4b, 0x16, 0xe6,
	0x92, 0x71, 0x23, 0xd9, 0xf8, 0xd5, 0x27, 0xc3, 0xee, 0xcd, 0xe9, 0xcf, 0x88, 0xac, 0x76, 0x6f,
	0xcd, 0x67, 0xc8, 0x36, 0xe3, 0x68, 0xb4, 0xb4, 0x59, 0x24, 0x7f, 0xa7, 0x03, 0xc9, 0xf5, 0x3b,
	0xa7, 0xff, 0x77, 0xca, 0x7b, 0x86, 0xf4, 0x47, 0x44, 0x56, 0x3a, 0x15, 0xf6, 0x19, 0x6d, 0xdd,
	0x60, 0x46, 0xc3, 0xeb, 0xad, 0x3c, 0x43, 0xba, 0x49, 0x6e, 0x4f, 0x8d, 0x75, 0x08, 0xa2, 0xe4,
	0xc2, 0x5a, 0x40, 0xc7, 0x9d, 0x2a, 0x01, 0xd9, 0x76, 0x1c, 0x8d, 0x56, 0xd3, 0xb5, 0x96, 0xdc,
	0x0d, 0xdc, 0x89, 0xa7, 0xfc, 0x2b, 0x69, 0xa4, 0x25, 0x38, 0x54, 0x39, 0x7b, 0x1a, 0xfa, 0x3f,
	0xa8, 0xc1, 0xc3, 0x80, 0xd1, 0x1d, 0xc2, 0x3a, 0x22, 0x6e, 0x10, 0x24, 0x20, 0x54, 0x39, 0xb0,
	0x71, 0xd0, 0xaf, 0x5f, 0xd7, 0x7f, 0x9a, 0xb3, 0xf4, 0x21, 0x19, 0x36, 0x7f, 0xa2, 0x71, 0x3c,
	0x53, 0x8e, 0x3d, 0x8b, 0xa3, 0x51, 0xbf, 0x8d, 0x9f, 0x1a, 0xb7, 0xa7, 0x9c, 0x7f, 0x4e, 0xb6,
	0xa1, 0x77, 0x02, 0xdd, 0xb3, 0x35, 0xf1, 0x2d, 0x22, 0x03, 0x34, 0x92, 0x57, 0xa0, 0xce, 0xbf,
	0x64, 0x1a, 0xd9, 0xf3, 0x1b, 0x2c, 0xef, 0x12, 0x1a, 0x79, 0xd4, 0x18, 0xd3, 0xaf, 0xa4, 0xef,
	0x13, 0x41, 0xad, 0x1d, 0x7b, 0x71, 0x83, 0x49, 0x2c, 0xa2, 0x91, 0xa9, 0xd6, 0x8e, 0xde, 0xad,
	0x13, 0x08, 0x33, 0xfa, 0x32, 0xd4, 0xdc, 0x53, 0x87, 0xc2, 0x5e, 0xb4, 0x94, 0x15, 0x52, 0xb1,
	0x57, 0x73, 0xea, 0x58, 0x48, 0xe5, 0xeb, 0xef, 0x29, 0x27, 0xb2, 0x09, 0xd4, 0xf3, 0xfa, 0xba,
	0xde, 0x02, 0x68, 0xe4, 0x89, 0x07, 0xc3, 0xd4, 0x36, 0x01, 0x0a, 0xd4, 0x86, 0xbd, 0x09, 0x0d,
	0xf0, 0x01, 0xf6, 0x51, 0x1b, 0x3f, 0xad, 0x9e, 0x82, 0x99, 0x43, 0x51, 0x81, 0x63, 0x6f, 0x03,
	0xed, 0x4b, 0x73, 0xd0, 0x40, 0xf4, 0x31, 0xa1, 0x5e, 0xa2, 0x2a, 0x07, 0x28, 0x45, 0xde, 0xf8,
	0xbc, 0x0b, 0x3e, 0x2b, 0x68, 0xe4, 0x87, 0x96, 0x08, 0x5e, 0x71, 0x1d, 0xf0, 0x12, 0x9b, 0xfd,
	0xb1, 0x1b, 0x74, 0x04, 0x8d, 0x3c, 0xc5, 0x7a, 0x87, 0x3c, 0x21, 0x6b, 0x99, 0x2a, 0x14, 0x42,
	0xee, 0x94, 0xae, 0xc4, 0x84, 0x87, 0x2d, 0xc0, 0xf6, 0x82, 0x33, 0xed, 0x50, 0xa9, 0x67, 0xe8,
	0x3a, 0xe9, 0x4d, 0x8d, 0x7f, 0xe9, 0xec, 0x7d, 0x1c, 0x8d, 0xfe, 0x4d, 0x9b, 0x13, 0xbd, 0x47,
	0xfa, 0x61, 0x43, 0xe7, 0x7a, 0xc2, 0xf6, 0x83, 0xcd, 0xfc, 0x9c, 0xf5, 0xc2, 0xd7, 0xd6, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x24, 0xd1, 0xf7, 0xd3, 0x05, 0x00, 0x00,
}
