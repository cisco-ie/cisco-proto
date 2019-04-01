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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_bidir_df_winners_bidir_df_winner

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
	RpAddress            string   `protobuf:"bytes,1,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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
	proto.RegisterType((*PimBidirDfBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bidir_df_winners.bidir_df_winner.pim_bidir_df_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bidir_df_winners.bidir_df_winner.pim_addrtype")
	proto.RegisterType((*PimBidirDfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bidir_df_winners.bidir_df_winner.pim_bidir_df_bag")
}

func init() { proto.RegisterFile("pim_bidir_df_bag.proto", fileDescriptor_f8d8ee746cdb1057) }

var fileDescriptor_f8d8ee746cdb1057 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xc1, 0x6b, 0xd4, 0x40,
	0x14, 0xc6, 0x19, 0xab, 0xdb, 0xee, 0xdb, 0xae, 0xd4, 0x81, 0xd6, 0x41, 0x10, 0xe2, 0x82, 0x10,
	0x50, 0x72, 0x68, 0xeb, 0xaa, 0x47, 0x41, 0x0f, 0xa2, 0x88, 0xc4, 0xc3, 0xea, 0x41, 0x1e, 0xb3,
	0x99, 0x37, 0x32, 0xd0, 0x4c, 0x86, 0x97, 0xa9, 0xdd, 0x5e, 0xbd, 0xf9, 0x67, 0xfa, 0x9f, 0xc8,
	0x24, 0xdb, 0xb8, 0xe4, 0x6e, 0x6f, 0x99, 0xef, 0xfd, 0xf8, 0xde, 0xfb, 0x3e, 0x08, 0x9c, 0x04,
	0x57, 0xe3, 0xda, 0x19, 0xc7, 0x68, 0x2c, 0xae, 0xf5, 0x8f, 0x22, 0x70, 0x13, 0x1b, 0xb9, 0xaa,
	0x5c, 0x5b, 0x35, 0xe8, 0x9a, 0x16, 0x37, 0x8c, 0x2e, 0xfc, 0x3c, 0xc7, 0x44, 0x36, 0x81, 0xb8,
	0x08, 0xae, 0x2e, 0xda, 0xa8, 0xbd, 0x59, 0x5f, 0x17, 0x86, 0xac, 0xbe, 0xbc, 0x88, 0x58, 0x35,
	0x3e, 0xd2, 0x26, 0x16, 0x83, 0xd5, 0x95, 0xf3, 0x9e, 0xb8, 0x1d, 0x0b, 0x8b, 0xef, 0x70, 0x3c,
	0x5e, 0x89, 0x1f, 0xde, 0x7d, 0xfb, 0x22, 0x1f, 0x03, 0x70, 0x40, 0x6d, 0x0c, 0x53, 0xdb, 0x2a,
	0x91, 0x89, 0x7c, 0x5a, 0x4e, 0x39, 0xbc, 0xe9, 0x05, 0xf9, 0x14, 0xee, 0x3b, 0x1f, 0x89, 0xad,
	0xae, 0x08, 0xbd, 0xae, 0x49, 0xdd, 0xe9, 0x90, 0xf9, 0xa0, 0x7e, 0xd2, 0x35, 0x2d, 0x6a, 0x38,
	0x4c, 0xf6, 0xc9, 0x26, 0x5e, 0x07, 0x92, 0x0f, 0x61, 0x5f, 0xdb, 0x9e, 0xef, 0x2d, 0x27, 0xda,
	0x26, 0x50, 0x3e, 0x81, 0xc3, 0x2e, 0xd5, 0xcd, 0xc2, 0xde, 0x6d, 0x96, 0xb4, 0x9b, 0x95, 0x3d,
	0xb2, 0x1c, 0x90, 0xbd, 0x01, 0x59, 0x6e, 0x91, 0xc5, 0x9f, 0x3d, 0x38, 0x1a, 0xc7, 0x91, 0xbf,
	0x05, 0xcc, 0xff, 0x45, 0xc1, 0x0d, 0xab, 0xd3, 0x4c, 0xe4, 0xb3, 0x53, 0x2a, 0xfe, 0x53, 0xa9,
	0xc5, 0x6e, 0xe4, 0x72, 0x36, 0x94, 0xf6, 0x95, 0xe5, 0x73, 0x90, 0x69, 0x38, 0xaa, 0xee, 0xac,
	0x4b, 0x92, 0x2e, 0x7f, 0xbf, 0xdb, 0x9e, 0xfc, 0x25, 0x60, 0x3a, 0xb8, 0xaa, 0xf3, 0xdb, 0xbc,
	0xfa, 0xc0, 0xd8, 0x55, 0x27, 0xcb, 0x47, 0x30, 0xd5, 0x4c, 0x78, 0x45, 0x68, 0xac, 0x7a, 0x91,
	0x89, 0xfc, 0xa0, 0xdc, 0xd7, 0x4c, 0x2b, 0x7a, 0x6b, 0xe5, 0x31, 0x4c, 0x38, 0xe0, 0x85, 0xf6,
	0x6a, 0xd9, 0x0d, 0xee, 0x71, 0xf8, 0xa8, 0xbd, 0x3c, 0x81, 0x49, 0x4d, 0x91, 0x5d, 0xa5, 0x5e,
	0x66, 0x22, 0x9f, 0x97, 0xdb, 0x97, 0x7c, 0x06, 0x0f, 0xfa, 0x2f, 0x0c, 0x4c, 0x96, 0x98, 0x7c,
	0x45, 0xea, 0x55, 0x87, 0x1c, 0xf5, 0x83, 0xcf, 0x83, 0x9e, 0x4c, 0x2e, 0x43, 0x74, 0x35, 0xa9,
	0xd7, 0x99, 0xc8, 0xef, 0x96, 0xdb, 0xd7, 0x7a, 0xd2, 0xfd, 0x11, 0x67, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x2c, 0xfa, 0x54, 0x2b, 0x03, 0x00, 0x00,
}
