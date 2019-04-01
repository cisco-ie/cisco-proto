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
// source: pim_tunnel_info_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_net_io_tunnels_net_io_tunnel

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

type PimTunnelInfoBag_KEYS struct {
	TunnelName           string   `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimTunnelInfoBag_KEYS) Reset()         { *m = PimTunnelInfoBag_KEYS{} }
func (m *PimTunnelInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimTunnelInfoBag_KEYS) ProtoMessage()    {}
func (*PimTunnelInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0876aa1a98c122bd, []int{0}
}

func (m *PimTunnelInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTunnelInfoBag_KEYS.Unmarshal(m, b)
}
func (m *PimTunnelInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTunnelInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimTunnelInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTunnelInfoBag_KEYS.Merge(m, src)
}
func (m *PimTunnelInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimTunnelInfoBag_KEYS.Size(m)
}
func (m *PimTunnelInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTunnelInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimTunnelInfoBag_KEYS proto.InternalMessageInfo

func (m *PimTunnelInfoBag_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
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
	return fileDescriptor_0876aa1a98c122bd, []int{1}
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

type PimTunnelInfoBag struct {
	SourceAddress        *PimAddrtype `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	RpAddress            *PimAddrtype `protobuf:"bytes,51,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	SourceAddressNetio   *PimAddrtype `protobuf:"bytes,52,opt,name=source_address_netio,json=sourceAddressNetio,proto3" json:"source_address_netio,omitempty"`
	GroupAddressNetio    *PimAddrtype `protobuf:"bytes,53,opt,name=group_address_netio,json=groupAddressNetio,proto3" json:"group_address_netio,omitempty"`
	VrfName              string       `protobuf:"bytes,54,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimTunnelInfoBag) Reset()         { *m = PimTunnelInfoBag{} }
func (m *PimTunnelInfoBag) String() string { return proto.CompactTextString(m) }
func (*PimTunnelInfoBag) ProtoMessage()    {}
func (*PimTunnelInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0876aa1a98c122bd, []int{2}
}

func (m *PimTunnelInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTunnelInfoBag.Unmarshal(m, b)
}
func (m *PimTunnelInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTunnelInfoBag.Marshal(b, m, deterministic)
}
func (m *PimTunnelInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTunnelInfoBag.Merge(m, src)
}
func (m *PimTunnelInfoBag) XXX_Size() int {
	return xxx_messageInfo_PimTunnelInfoBag.Size(m)
}
func (m *PimTunnelInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTunnelInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimTunnelInfoBag proto.InternalMessageInfo

func (m *PimTunnelInfoBag) GetSourceAddress() *PimAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *PimTunnelInfoBag) GetRpAddress() *PimAddrtype {
	if m != nil {
		return m.RpAddress
	}
	return nil
}

func (m *PimTunnelInfoBag) GetSourceAddressNetio() *PimAddrtype {
	if m != nil {
		return m.SourceAddressNetio
	}
	return nil
}

func (m *PimTunnelInfoBag) GetGroupAddressNetio() *PimAddrtype {
	if m != nil {
		return m.GroupAddressNetio
	}
	return nil
}

func (m *PimTunnelInfoBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func init() {
	proto.RegisterType((*PimTunnelInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.net_io_tunnels.net_io_tunnel.pim_addrtype")
	proto.RegisterType((*PimTunnelInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag")
}

func init() { proto.RegisterFile("pim_tunnel_info_bag.proto", fileDescriptor_0876aa1a98c122bd) }

var fileDescriptor_0876aa1a98c122bd = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xb1, 0x4f, 0x3a, 0x31,
	0x14, 0xc7, 0xd3, 0xdf, 0xcf, 0x80, 0x3c, 0xd0, 0xc4, 0x62, 0xe2, 0x31, 0x89, 0x4c, 0x4c, 0x1d,
	0x00, 0x6f, 0x71, 0x72, 0x70, 0x32, 0x61, 0xc0, 0x45, 0xa7, 0xa6, 0xdc, 0xf5, 0x48, 0x13, 0xae,
	0x6d, 0xda, 0xde, 0x05, 0xfe, 0x00, 0x47, 0x57, 0xe3, 0x1f, 0xeb, 0x60, 0x5a, 0x8e, 0x0b, 0x67,
	0x58, 0x65, 0xec, 0xf7, 0x3e, 0xbd, 0xf7, 0x79, 0x2f, 0xaf, 0x30, 0xd0, 0x22, 0xa7, 0xae, 0x90,
	0x92, 0xaf, 0xa9, 0x90, 0x99, 0xa2, 0x4b, 0xb6, 0x22, 0xda, 0x28, 0xa7, 0xf0, 0x6b, 0x22, 0x6c,
	0xa2, 0xa8, 0x50, 0x96, 0x6e, 0x0c, 0x15, 0xba, 0x9c, 0x51, 0x0f, 0x2b, 0xcd, 0x0d, 0x11, 0xba,
	0x8c, 0xfd, 0x89, 0x58, 0xc7, 0x64, 0xba, 0xdc, 0x92, 0x94, 0x67, 0xac, 0x58, 0x3b, 0x9a, 0x28,
	0xe9, 0xf8, 0xc6, 0x11, 0xc9, 0x1d, 0x15, 0xaa, 0xfa, 0xb3, 0x6d, 0x1e, 0x47, 0x0f, 0x10, 0x1d,
	0x29, 0x4b, 0x9f, 0x9f, 0xde, 0x5e, 0xf0, 0x2d, 0x74, 0xab, 0x5c, 0xb2, 0x9c, 0x47, 0x68, 0x88,
	0xc6, 0x9d, 0x05, 0xec, 0xa2, 0x39, 0xcb, 0xf9, 0x28, 0x87, 0x9e, 0xbf, 0xcc, 0xd2, 0xd4, 0xb8,
	0xad, 0xe6, 0xf8, 0x06, 0xda, 0x2c, 0x3b, 0x84, 0x5b, 0x2c, 0xf3, 0x20, 0xbe, 0x83, 0x5e, 0x90,
	0xf6, 0x24, 0xb7, 0x36, 0xfa, 0x17, 0xbe, 0x76, 0x7d, 0xf6, 0xb8, 0x8b, 0x2a, 0x24, 0xae, 0x91,
	0xff, 0x35, 0x12, 0x57, 0xc8, 0xe8, 0xfb, 0x0c, 0xfa, 0x47, 0x64, 0xf1, 0x07, 0x82, 0x4b, 0xab,
	0x0a, 0x93, 0xf0, 0xfa, 0xf6, 0x64, 0x88, 0xc6, 0xdd, 0x49, 0x46, 0xfe, 0x6a, 0x6e, 0xe4, 0xb0,
	0xef, 0xc5, 0xc5, 0xae, 0xfa, 0xbe, 0x95, 0x77, 0x04, 0x60, 0x74, 0xed, 0x32, 0x3d, 0xa9, 0x4b,
	0xc7, 0xe8, 0xbd, 0xc7, 0x17, 0x82, 0xeb, 0xe6, 0x5c, 0xa8, 0xe4, 0x4e, 0xa8, 0x68, 0x76, 0x52,
	0x23, 0xdc, 0x98, 0xce, 0xdc, 0x1b, 0xe0, 0x4f, 0x04, 0xfd, 0x95, 0x51, 0x85, 0xfe, 0x65, 0x76,
	0x7f, 0x52, 0xb3, 0xab, 0xa0, 0xd0, 0x10, 0x1b, 0xc0, 0x79, 0x69, 0xaa, 0x1d, 0x8e, 0xc3, 0x0a,
	0xb6, 0x4b, 0x13, 0x96, 0x78, 0xd9, 0x0a, 0x6f, 0x71, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0x7c, 0x91, 0xf2, 0xa8, 0x03, 0x00, 0x00,
}
