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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_net_io_tunnels_net_io_tunnel

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	TunnelName           string   `protobuf:"bytes,2,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
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

func (m *PimTunnelInfoBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimTunnelInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_addrtype")
	proto.RegisterType((*PimTunnelInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag")
}

func init() { proto.RegisterFile("pim_tunnel_info_bag.proto", fileDescriptor_0876aa1a98c122bd) }

var fileDescriptor_0876aa1a98c122bd = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0x96, 0xdf, 0x17, 0xb5, 0xf4, 0x5a, 0x90, 0x70, 0x91, 0x68, 0x27, 0x4a, 0xa7, 0x4e, 0x1e,
	0xda, 0xd2, 0x9d, 0x81, 0x09, 0xd4, 0xa1, 0x48, 0x48, 0xb0, 0x58, 0x6e, 0xe2, 0x54, 0x96, 0x88,
	0x6d, 0xd9, 0x4e, 0x44, 0x46, 0x56, 0xc4, 0x3f, 0xe3, 0x4f, 0x21, 0x27, 0x69, 0x54, 0xa3, 0x8e,
	0x64, 0x89, 0x94, 0xbb, 0xe7, 0x9e, 0x8f, 0xd3, 0x19, 0xc6, 0x5a, 0xa4, 0xd4, 0x65, 0x52, 0xf2,
	0x37, 0x2a, 0x64, 0xa2, 0xe8, 0x96, 0xed, 0x88, 0x36, 0xca, 0x29, 0xfc, 0x18, 0x09, 0x1b, 0x29,
	0x2a, 0x94, 0xa5, 0xef, 0x86, 0x0a, 0x9d, 0x2f, 0xa9, 0x07, 0x2b, 0xcd, 0x0d, 0xd1, 0x22, 0x25,
	0xd6, 0x31, 0x19, 0x6f, 0x0b, 0x92, 0x9b, 0xc4, 0xfa, 0x0f, 0x91, 0xdc, 0x51, 0xa1, 0x6a, 0x36,
	0x1b, 0xfe, 0x4e, 0x9f, 0x61, 0x74, 0x44, 0x8a, 0x3e, 0xdc, 0xbf, 0x3c, 0xe1, 0x31, 0x9c, 0xe6,
	0x26, 0xa1, 0x92, 0xa5, 0x7c, 0x84, 0x26, 0x68, 0xd6, 0xdb, 0x74, 0x73, 0x93, 0xac, 0x59, 0xca,
	0xf1, 0x35, 0xf4, 0xeb, 0x91, 0xb2, 0xfb, 0xaf, 0xec, 0x42, 0x55, 0xf2, 0x80, 0x69, 0x0a, 0x03,
	0xcf, 0xcb, 0xe2, 0xd8, 0xb8, 0x42, 0x73, 0x7c, 0x05, 0x5d, 0x16, 0x50, 0x75, 0x58, 0xc5, 0x74,
	0x03, 0x83, 0x32, 0x83, 0x47, 0x72, 0x6b, 0x6b, 0xaa, 0xbe, 0xaf, 0xdd, 0x55, 0xa5, 0x1a, 0xb2,
	0x6a, 0x20, 0xff, 0x1b, 0xc8, 0xaa, 0x86, 0x4c, 0xbf, 0x4f, 0x60, 0x78, 0x24, 0x07, 0xfe, 0x40,
	0x70, 0x6e, 0x55, 0x66, 0x22, 0xde, 0x4c, 0xcf, 0x27, 0x68, 0xd6, 0x9f, 0xbf, 0x92, 0xbf, 0x5c,
	0x23, 0x39, 0xcc, 0xba, 0x39, 0xab, 0x14, 0xf7, 0xf6, 0x0b, 0x00, 0xa3, 0x1b, 0xf9, 0x45, 0xeb,
	0xf2, 0x3d, 0xa3, 0xf7, 0xd2, 0x5f, 0x08, 0x2e, 0xc3, 0xf8, 0x54, 0x72, 0x27, 0xd4, 0x68, 0xd9,
	0xba, 0x0b, 0x1c, 0x2c, 0x61, 0xed, 0x55, 0xf1, 0x27, 0x82, 0xe1, 0xce, 0xa8, 0x4c, 0xff, 0x72,
	0x73, 0xdb, 0xba, 0x9b, 0x8b, 0x52, 0x36, 0x30, 0x73, 0x78, 0xdd, 0xab, 0xe0, 0xba, 0xb7, 0x9d,
	0xf2, 0xa5, 0x2d, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x72, 0xf7, 0xc6, 0x60, 0x86, 0x03, 0x00,
	0x00,
}
