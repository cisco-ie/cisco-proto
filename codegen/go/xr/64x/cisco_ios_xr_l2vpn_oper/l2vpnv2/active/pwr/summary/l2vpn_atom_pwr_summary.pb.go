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
// source: l2vpn_atom_pwr_summary.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_pwr_summary

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

type L2VpnAtomPwrSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnAtomPwrSummary_KEYS) Reset()         { *m = L2VpnAtomPwrSummary_KEYS{} }
func (m *L2VpnAtomPwrSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnAtomPwrSummary_KEYS) ProtoMessage()    {}
func (*L2VpnAtomPwrSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{0}
}

func (m *L2VpnAtomPwrSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnAtomPwrSummary_KEYS.Unmarshal(m, b)
}
func (m *L2VpnAtomPwrSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnAtomPwrSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnAtomPwrSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnAtomPwrSummary_KEYS.Merge(m, src)
}
func (m *L2VpnAtomPwrSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnAtomPwrSummary_KEYS.Size(m)
}
func (m *L2VpnAtomPwrSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnAtomPwrSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnAtomPwrSummary_KEYS proto.InternalMessageInfo

type L2VpnRdAuto struct {
	RouterId             string   `protobuf:"bytes,1,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	AutoIndex            uint32   `protobuf:"varint,2,opt,name=auto_index,json=autoIndex,proto3" json:"auto_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRdAuto) Reset()         { *m = L2VpnRdAuto{} }
func (m *L2VpnRdAuto) String() string { return proto.CompactTextString(m) }
func (*L2VpnRdAuto) ProtoMessage()    {}
func (*L2VpnRdAuto) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{1}
}

func (m *L2VpnRdAuto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRdAuto.Unmarshal(m, b)
}
func (m *L2VpnRdAuto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRdAuto.Marshal(b, m, deterministic)
}
func (m *L2VpnRdAuto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRdAuto.Merge(m, src)
}
func (m *L2VpnRdAuto) XXX_Size() int {
	return xxx_messageInfo_L2VpnRdAuto.Size(m)
}
func (m *L2VpnRdAuto) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRdAuto.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRdAuto proto.InternalMessageInfo

func (m *L2VpnRdAuto) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *L2VpnRdAuto) GetAutoIndex() uint32 {
	if m != nil {
		return m.AutoIndex
	}
	return 0
}

type L2VpnRd_2ByteAs struct {
	TwoByteAs            uint32   `protobuf:"varint,1,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteIndex        uint32   `protobuf:"varint,2,opt,name=four_byte_index,json=fourByteIndex,proto3" json:"four_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRd_2ByteAs) Reset()         { *m = L2VpnRd_2ByteAs{} }
func (m *L2VpnRd_2ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd_2ByteAs) ProtoMessage()    {}
func (*L2VpnRd_2ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{2}
}

func (m *L2VpnRd_2ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRd_2ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRd_2ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd_2ByteAs.Merge(m, src)
}
func (m *L2VpnRd_2ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd_2ByteAs.Size(m)
}
func (m *L2VpnRd_2ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd_2ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd_2ByteAs proto.InternalMessageInfo

func (m *L2VpnRd_2ByteAs) GetTwoByteAs() uint32 {
	if m != nil {
		return m.TwoByteAs
	}
	return 0
}

func (m *L2VpnRd_2ByteAs) GetFourByteIndex() uint32 {
	if m != nil {
		return m.FourByteIndex
	}
	return 0
}

type L2VpnRd_4ByteAs struct {
	FourByteAs           uint32   `protobuf:"varint,1,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRd_4ByteAs) Reset()         { *m = L2VpnRd_4ByteAs{} }
func (m *L2VpnRd_4ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd_4ByteAs) ProtoMessage()    {}
func (*L2VpnRd_4ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{3}
}

func (m *L2VpnRd_4ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRd_4ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRd_4ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd_4ByteAs.Merge(m, src)
}
func (m *L2VpnRd_4ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd_4ByteAs.Size(m)
}
func (m *L2VpnRd_4ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd_4ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd_4ByteAs proto.InternalMessageInfo

func (m *L2VpnRd_4ByteAs) GetFourByteAs() uint32 {
	if m != nil {
		return m.FourByteAs
	}
	return 0
}

func (m *L2VpnRd_4ByteAs) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRdV4Addr struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRdV4Addr) Reset()         { *m = L2VpnRdV4Addr{} }
func (m *L2VpnRdV4Addr) String() string { return proto.CompactTextString(m) }
func (*L2VpnRdV4Addr) ProtoMessage()    {}
func (*L2VpnRdV4Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{4}
}

func (m *L2VpnRdV4Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRdV4Addr.Unmarshal(m, b)
}
func (m *L2VpnRdV4Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRdV4Addr.Marshal(b, m, deterministic)
}
func (m *L2VpnRdV4Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRdV4Addr.Merge(m, src)
}
func (m *L2VpnRdV4Addr) XXX_Size() int {
	return xxx_messageInfo_L2VpnRdV4Addr.Size(m)
}
func (m *L2VpnRdV4Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRdV4Addr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRdV4Addr proto.InternalMessageInfo

func (m *L2VpnRdV4Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *L2VpnRdV4Addr) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRd struct {
	Rd                   string           `protobuf:"bytes,1,opt,name=rd,proto3" json:"rd,omitempty"`
	Auto                 *L2VpnRdAuto     `protobuf:"bytes,2,opt,name=auto,proto3" json:"auto,omitempty"`
	TwoByteAs            *L2VpnRd_2ByteAs `protobuf:"bytes,3,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteAs           *L2VpnRd_4ByteAs `protobuf:"bytes,4,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	V4Addr               *L2VpnRdV4Addr   `protobuf:"bytes,5,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnRd) Reset()         { *m = L2VpnRd{} }
func (m *L2VpnRd) String() string { return proto.CompactTextString(m) }
func (*L2VpnRd) ProtoMessage()    {}
func (*L2VpnRd) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{5}
}

func (m *L2VpnRd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRd.Unmarshal(m, b)
}
func (m *L2VpnRd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRd.Marshal(b, m, deterministic)
}
func (m *L2VpnRd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRd.Merge(m, src)
}
func (m *L2VpnRd) XXX_Size() int {
	return xxx_messageInfo_L2VpnRd.Size(m)
}
func (m *L2VpnRd) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRd.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRd proto.InternalMessageInfo

func (m *L2VpnRd) GetRd() string {
	if m != nil {
		return m.Rd
	}
	return ""
}

func (m *L2VpnRd) GetAuto() *L2VpnRdAuto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *L2VpnRd) GetTwoByteAs() *L2VpnRd_2ByteAs {
	if m != nil {
		return m.TwoByteAs
	}
	return nil
}

func (m *L2VpnRd) GetFourByteAs() *L2VpnRd_4ByteAs {
	if m != nil {
		return m.FourByteAs
	}
	return nil
}

func (m *L2VpnRd) GetV4Addr() *L2VpnRdV4Addr {
	if m != nil {
		return m.V4Addr
	}
	return nil
}

type L2VpnAtomPwrSummary struct {
	BgpRouterId          string   `protobuf:"bytes,50,opt,name=bgp_router_id,json=bgpRouterId,proto3" json:"bgp_router_id,omitempty"`
	CfgRouterId          string   `protobuf:"bytes,51,opt,name=cfg_router_id,json=cfgRouterId,proto3" json:"cfg_router_id,omitempty"`
	BgpAs                uint32   `protobuf:"varint,52,opt,name=bgp_as,json=bgpAs,proto3" json:"bgp_as,omitempty"`
	CfgGlobalId          uint32   `protobuf:"varint,53,opt,name=cfg_global_id,json=cfgGlobalId,proto3" json:"cfg_global_id,omitempty"`
	RdAuto               *L2VpnRd `protobuf:"bytes,54,opt,name=rd_auto,json=rdAuto,proto3" json:"rd_auto,omitempty"`
	RdConfigured         *L2VpnRd `protobuf:"bytes,55,opt,name=rd_configured,json=rdConfigured,proto3" json:"rd_configured,omitempty"`
	L2VpnHasBgpEod       bool     `protobuf:"varint,56,opt,name=l2vpn_has_bgp_eod,json=l2vpnHasBgpEod,proto3" json:"l2vpn_has_bgp_eod,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnAtomPwrSummary) Reset()         { *m = L2VpnAtomPwrSummary{} }
func (m *L2VpnAtomPwrSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnAtomPwrSummary) ProtoMessage()    {}
func (*L2VpnAtomPwrSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a6e1ee7d75ddfe3, []int{6}
}

func (m *L2VpnAtomPwrSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnAtomPwrSummary.Unmarshal(m, b)
}
func (m *L2VpnAtomPwrSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnAtomPwrSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnAtomPwrSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnAtomPwrSummary.Merge(m, src)
}
func (m *L2VpnAtomPwrSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnAtomPwrSummary.Size(m)
}
func (m *L2VpnAtomPwrSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnAtomPwrSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnAtomPwrSummary proto.InternalMessageInfo

func (m *L2VpnAtomPwrSummary) GetBgpRouterId() string {
	if m != nil {
		return m.BgpRouterId
	}
	return ""
}

func (m *L2VpnAtomPwrSummary) GetCfgRouterId() string {
	if m != nil {
		return m.CfgRouterId
	}
	return ""
}

func (m *L2VpnAtomPwrSummary) GetBgpAs() uint32 {
	if m != nil {
		return m.BgpAs
	}
	return 0
}

func (m *L2VpnAtomPwrSummary) GetCfgGlobalId() uint32 {
	if m != nil {
		return m.CfgGlobalId
	}
	return 0
}

func (m *L2VpnAtomPwrSummary) GetRdAuto() *L2VpnRd {
	if m != nil {
		return m.RdAuto
	}
	return nil
}

func (m *L2VpnAtomPwrSummary) GetRdConfigured() *L2VpnRd {
	if m != nil {
		return m.RdConfigured
	}
	return nil
}

func (m *L2VpnAtomPwrSummary) GetL2VpnHasBgpEod() bool {
	if m != nil {
		return m.L2VpnHasBgpEod
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnAtomPwrSummary_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_atom_pwr_summary_KEYS")
	proto.RegisterType((*L2VpnRdAuto)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_rd_auto")
	proto.RegisterType((*L2VpnRd_2ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_rd_2byte_as")
	proto.RegisterType((*L2VpnRd_4ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_rd_4byte_as")
	proto.RegisterType((*L2VpnRdV4Addr)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_rd_v4addr")
	proto.RegisterType((*L2VpnRd)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_rd")
	proto.RegisterType((*L2VpnAtomPwrSummary)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.pwr.summary.l2vpn_atom_pwr_summary")
}

func init() { proto.RegisterFile("l2vpn_atom_pwr_summary.proto", fileDescriptor_0a6e1ee7d75ddfe3) }

var fileDescriptor_0a6e1ee7d75ddfe3 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4d, 0x8f, 0xda, 0x30,
	0x10, 0x86, 0x05, 0xec, 0xb2, 0x30, 0x10, 0x56, 0xb5, 0xd4, 0x2a, 0xd2, 0x76, 0x2b, 0x1a, 0x55,
	0x15, 0xbd, 0xe4, 0x90, 0xa5, 0x1f, 0x87, 0x5e, 0xb2, 0x2b, 0xd4, 0xa2, 0xbd, 0xa5, 0xea, 0xa1,
	0x1f, 0x92, 0xe5, 0xc4, 0x26, 0x8d, 0x04, 0xd8, 0xb2, 0x93, 0xb0, 0xdc, 0xfa, 0x4f, 0xfa, 0x57,
	0x2b, 0xdb, 0x49, 0x16, 0xda, 0x3d, 0xb4, 0x82, 0x1b, 0x1e, 0xcf, 0xfb, 0x0c, 0x9e, 0x77, 0x26,
	0xf0, 0x74, 0x19, 0x94, 0x62, 0x8d, 0x49, 0xce, 0x57, 0x58, 0x6c, 0x24, 0x56, 0xc5, 0x6a, 0x45,
	0xe4, 0xd6, 0x17, 0x92, 0xe7, 0x1c, 0x05, 0x49, 0xa6, 0x12, 0x8e, 0x33, 0xae, 0xf0, 0x9d, 0xc4,
	0x36, 0x95, 0x0b, 0x26, 0x7d, 0xf3, 0xb3, 0x0c, 0x7c, 0x92, 0xe4, 0x59, 0xc9, 0x7c, 0xb1, 0x91,
	0x7e, 0xa5, 0xf4, 0x2e, 0xe1, 0xe2, 0x61, 0x26, 0xbe, 0x9d, 0x7d, 0xf9, 0xe4, 0xdd, 0x82, 0x63,
	0xaf, 0x25, 0xc5, 0xa4, 0xc8, 0x39, 0xba, 0x80, 0xbe, 0xe4, 0x45, 0xce, 0x24, 0xce, 0xa8, 0xdb,
	0x1a, 0xb7, 0x26, 0xfd, 0xa8, 0x67, 0x03, 0x73, 0x8a, 0x2e, 0x01, 0x74, 0x12, 0xce, 0xd6, 0x94,
	0xdd, 0xb9, 0xed, 0x71, 0x6b, 0xe2, 0x44, 0x7d, 0x1d, 0x99, 0xeb, 0x80, 0xf7, 0x0d, 0x1e, 0x35,
	0xb0, 0x20, 0xde, 0xe6, 0x0c, 0x13, 0x85, 0x9e, 0xc1, 0x20, 0xdf, 0x70, 0x5c, 0x1d, 0x0d, 0xd2,
	0x89, 0xfa, 0xf9, 0x86, 0x5f, 0x6f, 0x73, 0x16, 0x2a, 0xf4, 0x12, 0xce, 0x17, 0xbc, 0x90, 0x36,
	0x61, 0x17, 0xec, 0xe8, 0xb0, 0x4e, 0xfa, 0x1b, 0x3e, 0xad, 0xe1, 0x63, 0x18, 0xde, 0x8b, 0x1b,
	0x3a, 0xd4, 0xca, 0x50, 0xa1, 0x17, 0x30, 0x6a, 0xca, 0xef, 0xd2, 0x87, 0xd5, 0x3f, 0xb0, 0xf0,
	0xaf, 0x70, 0xde, 0xc0, 0xcb, 0x29, 0xa1, 0x54, 0xa2, 0xe7, 0x30, 0xcc, 0x44, 0x39, 0xc5, 0xfa,
	0xc0, 0x94, 0xaa, 0x7a, 0x31, 0xd0, 0xb1, 0xd0, 0x86, 0xfe, 0x91, 0xfd, 0xab, 0x03, 0xbd, 0x1a,
	0x8e, 0x46, 0xd0, 0x96, 0x75, 0x5f, 0xdb, 0x92, 0xa2, 0xcf, 0x70, 0xa2, 0xfb, 0x67, 0x84, 0x83,
	0x20, 0xf4, 0xff, 0xdf, 0x61, 0x7f, 0xcf, 0xbf, 0xc8, 0xe0, 0x10, 0xdb, 0x6f, 0x7a, 0xc7, 0xd0,
	0x67, 0x07, 0xd1, 0x6b, 0x43, 0x77, 0xbd, 0x4b, 0xff, 0x68, 0xff, 0xc9, 0x11, 0xea, 0xd4, 0xde,
	0xee, 0xb9, 0xf8, 0x1d, 0xce, 0x2a, 0x2b, 0xdc, 0x53, 0x53, 0xe3, 0xe6, 0xa0, 0x1a, 0xd6, 0xe2,
	0xa8, 0x6b, 0xad, 0xf4, 0x7e, 0x76, 0xe0, 0xc9, 0xc3, 0x4b, 0x82, 0x3c, 0x70, 0xe2, 0x54, 0xe0,
	0xfb, 0x95, 0x08, 0xec, 0x18, 0xc4, 0xa9, 0x88, 0xea, 0xad, 0xf0, 0xc0, 0x49, 0x16, 0xe9, 0x4e,
	0xce, 0x95, 0xcd, 0x49, 0x16, 0x69, 0x93, 0xf3, 0x18, 0xba, 0x9a, 0x43, 0x94, 0x3b, 0x35, 0x23,
	0x72, 0x1a, 0xa7, 0x22, 0x54, 0xb5, 0x34, 0x5d, 0xf2, 0x98, 0x2c, 0xb5, 0xf4, 0xb5, 0xb9, 0xd5,
	0xd2, 0x0f, 0x26, 0x36, 0xd7, 0x23, 0x72, 0x56, 0x99, 0xeb, 0xbe, 0x31, 0x6f, 0x7f, 0x7f, 0xc8,
	0xdb, 0xa3, 0xae, 0xa4, 0xa1, 0x1e, 0x11, 0x02, 0x8e, 0xa4, 0x38, 0xe1, 0xeb, 0x45, 0x96, 0x16,
	0x92, 0x51, 0xf7, 0xed, 0x11, 0xe0, 0x43, 0x49, 0x6f, 0x1a, 0x22, 0x7a, 0x55, 0xaf, 0xec, 0x0f,
	0xa2, 0xb0, 0x7e, 0x3e, 0xe3, 0xd4, 0x7d, 0x37, 0x6e, 0x4d, 0x7a, 0xd1, 0xc8, 0x5c, 0x7c, 0x24,
	0xea, 0x3a, 0x15, 0x33, 0x4e, 0xe3, 0xae, 0xf9, 0xc2, 0x5d, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x63, 0x45, 0xe1, 0xe3, 0x01, 0x05, 0x00, 0x00,
}