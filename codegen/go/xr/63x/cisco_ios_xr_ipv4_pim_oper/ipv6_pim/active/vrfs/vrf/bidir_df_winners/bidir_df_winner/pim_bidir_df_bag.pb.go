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
// source: pim_bidir_df_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_bidir_df_winners_bidir_df_winner

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

type PimBidirDfBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RpAddress            string   `protobuf:"bytes,2,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBidirDfBag_KEYS) Reset()         { *m = PimBidirDfBag_KEYS{} }
func (m *PimBidirDfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBidirDfBag_KEYS) ProtoMessage()    {}
func (*PimBidirDfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d8ee746cdb1057, []int{0}
}

func (m *PimBidirDfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBidirDfBag_KEYS.Unmarshal(m, b)
}
func (m *PimBidirDfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBidirDfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBidirDfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBidirDfBag_KEYS.Merge(m, src)
}
func (m *PimBidirDfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBidirDfBag_KEYS.Size(m)
}
func (m *PimBidirDfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBidirDfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBidirDfBag_KEYS proto.InternalMessageInfo

func (m *PimBidirDfBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimBidirDfBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

func (m *PimBidirDfBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
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
	return fileDescriptor_f8d8ee746cdb1057, []int{1}
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

type PimBidirDfBag struct {
	RpAddressXr          *PimAddrtype `protobuf:"bytes,50,opt,name=rp_address_xr,json=rpAddressXr,proto3" json:"rp_address_xr,omitempty"`
	PimInterfaceName     string       `protobuf:"bytes,51,opt,name=pim_interface_name,json=pimInterfaceName,proto3" json:"pim_interface_name,omitempty"`
	DfWinner             *PimAddrtype `protobuf:"bytes,52,opt,name=df_winner,json=dfWinner,proto3" json:"df_winner,omitempty"`
	AreWeDf              bool         `protobuf:"varint,53,opt,name=are_we_df,json=areWeDf,proto3" json:"are_we_df,omitempty"`
	RpLan                bool         `protobuf:"varint,54,opt,name=rp_lan,json=rpLan,proto3" json:"rp_lan,omitempty"`
	Metric               uint32       `protobuf:"varint,55,opt,name=metric,proto3" json:"metric,omitempty"`
	MetricPreference     uint32       `protobuf:"varint,56,opt,name=metric_preference,json=metricPreference,proto3" json:"metric_preference,omitempty"`
	Uptime               uint64       `protobuf:"varint,57,opt,name=uptime,proto3" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimBidirDfBag) Reset()         { *m = PimBidirDfBag{} }
func (m *PimBidirDfBag) String() string { return proto.CompactTextString(m) }
func (*PimBidirDfBag) ProtoMessage()    {}
func (*PimBidirDfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8d8ee746cdb1057, []int{2}
}

func (m *PimBidirDfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBidirDfBag.Unmarshal(m, b)
}
func (m *PimBidirDfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBidirDfBag.Marshal(b, m, deterministic)
}
func (m *PimBidirDfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBidirDfBag.Merge(m, src)
}
func (m *PimBidirDfBag) XXX_Size() int {
	return xxx_messageInfo_PimBidirDfBag.Size(m)
}
func (m *PimBidirDfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBidirDfBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBidirDfBag proto.InternalMessageInfo

func (m *PimBidirDfBag) GetRpAddressXr() *PimAddrtype {
	if m != nil {
		return m.RpAddressXr
	}
	return nil
}

func (m *PimBidirDfBag) GetPimInterfaceName() string {
	if m != nil {
		return m.PimInterfaceName
	}
	return ""
}

func (m *PimBidirDfBag) GetDfWinner() *PimAddrtype {
	if m != nil {
		return m.DfWinner
	}
	return nil
}

func (m *PimBidirDfBag) GetAreWeDf() bool {
	if m != nil {
		return m.AreWeDf
	}
	return false
}

func (m *PimBidirDfBag) GetRpLan() bool {
	if m != nil {
		return m.RpLan
	}
	return false
}

func (m *PimBidirDfBag) GetMetric() uint32 {
	if m != nil {
		return m.Metric
	}
	return 0
}

func (m *PimBidirDfBag) GetMetricPreference() uint32 {
	if m != nil {
		return m.MetricPreference
	}
	return 0
}

func (m *PimBidirDfBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func init() {
	proto.RegisterType((*PimBidirDfBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.bidir_df_winners.bidir_df_winner.pim_bidir_df_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.bidir_df_winners.bidir_df_winner.pim_addrtype")
	proto.RegisterType((*PimBidirDfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.bidir_df_winners.bidir_df_winner.pim_bidir_df_bag")
}

func init() { proto.RegisterFile("pim_bidir_df_bag.proto", fileDescriptor_f8d8ee746cdb1057) }

var fileDescriptor_f8d8ee746cdb1057 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x41, 0xcb, 0xd3, 0x40,
	0x10, 0x25, 0x56, 0xd3, 0x66, 0xfa, 0x55, 0x3e, 0x17, 0xbe, 0xba, 0x0a, 0x42, 0x2c, 0x08, 0x01,
	0x25, 0x87, 0xb6, 0x46, 0x3d, 0x0a, 0x7a, 0x10, 0x45, 0x24, 0x0a, 0xd5, 0xd3, 0xb2, 0x49, 0x66,
	0x65, 0xc1, 0x24, 0xcb, 0x24, 0xa6, 0xf5, 0xe4, 0xcd, 0x1f, 0xe9, 0xaf, 0x91, 0xdd, 0xd4, 0xd8,
	0xe6, 0xee, 0x77, 0x09, 0x99, 0x37, 0x2f, 0xef, 0xcd, 0x7b, 0x10, 0x58, 0x1a, 0x5d, 0x8a, 0x4c,
	0x17, 0x9a, 0x44, 0xa1, 0x44, 0x26, 0xbf, 0xc6, 0x86, 0xea, 0xb6, 0x66, 0x9f, 0x72, 0xdd, 0xe4,
	0xb5, 0xd0, 0x75, 0x23, 0x0e, 0x24, 0xb4, 0xe9, 0xb6, 0xc2, 0x32, 0x6b, 0x83, 0x14, 0x6b, 0xd3,
	0x25, 0x76, 0x8a, 0x65, 0xde, 0xea, 0x0e, 0xe3, 0x8e, 0x54, 0x63, 0x1f, 0xf1, 0xa0, 0xb3, 0xd7,
	0x55, 0x85, 0xd4, 0x8c, 0x81, 0xd5, 0x01, 0xae, 0xc6, 0x7e, 0xe2, 0xed, 0xeb, 0x2f, 0x1f, 0xd9,
	0x3d, 0x98, 0x75, 0xa4, 0x44, 0x25, 0x4b, 0xe4, 0x5e, 0xe8, 0x45, 0x41, 0x3a, 0xed, 0x48, 0xbd,
	0x97, 0x25, 0xb2, 0x07, 0x00, 0x64, 0x84, 0x2c, 0x0a, 0xc2, 0xa6, 0xe1, 0x37, 0xdc, 0x32, 0x20,
	0xf3, 0xb2, 0x07, 0xd8, 0x23, 0xb8, 0xad, 0xab, 0x16, 0x49, 0xc9, 0x1c, 0xfb, 0xef, 0x27, 0x8e,
	0xb2, 0x18, 0x50, 0xab, 0xb2, 0x2a, 0xe1, 0xc2, 0x3a, 0x5b, 0x99, 0xf6, 0x87, 0x41, 0x76, 0x17,
	0xa6, 0xf2, 0xcc, 0xcf, 0x97, 0xbd, 0xdd, 0x43, 0xb8, 0x70, 0x69, 0xcf, 0x0d, 0xe7, 0x16, 0xfb,
	0x6b, 0xd9, 0x53, 0x92, 0x81, 0x32, 0x19, 0x28, 0xc9, 0x91, 0xb2, 0xfa, 0x3d, 0x81, 0xcb, 0x71,
	0x52, 0xf6, 0xcb, 0x83, 0xc5, 0xbf, 0x28, 0xe2, 0x40, 0x7c, 0x1d, 0x7a, 0xd1, 0x7c, 0x9d, 0xc5,
	0xff, 0xa3, 0xec, 0xf8, 0x34, 0x6f, 0x3a, 0x1f, 0x1a, 0xfb, 0x4c, 0xec, 0x09, 0x30, 0xbb, 0x1c,
	0xf5, 0xb6, 0x71, 0x31, 0xec, 0xd9, 0x6f, 0x4e, 0xab, 0x63, 0x3f, 0x21, 0x18, 0x44, 0xf9, 0xf6,
	0xda, 0x2e, 0x9e, 0x15, 0x6a, 0xe7, 0x60, 0x76, 0x1f, 0x02, 0x49, 0x28, 0xf6, 0x28, 0x0a, 0xc5,
	0x9f, 0x86, 0x5e, 0x34, 0x4b, 0xa7, 0x92, 0x70, 0x87, 0xaf, 0x14, 0xbb, 0x02, 0x9f, 0x8c, 0xf8,
	0x26, 0x2b, 0x9e, 0xb8, 0xc5, 0x2d, 0x32, 0xef, 0x64, 0xc5, 0x96, 0xe0, 0x97, 0xd8, 0x92, 0xce,
	0xf9, 0xb3, 0xd0, 0x8b, 0x16, 0xe9, 0x71, 0x62, 0x8f, 0xe1, 0x4e, 0xff, 0x26, 0x0c, 0xa1, 0x42,
	0xc2, 0x2a, 0x47, 0xfe, 0xdc, 0x51, 0x2e, 0xfb, 0xc5, 0x87, 0x01, 0xb7, 0x22, 0xdf, 0x4d, 0xab,
	0x4b, 0xe4, 0x2f, 0x42, 0x2f, 0xba, 0x99, 0x1e, 0xa7, 0xcc, 0x77, 0xbf, 0xc8, 0xe6, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbe, 0x2c, 0xac, 0x66, 0x3c, 0x03, 0x00, 0x00,
}