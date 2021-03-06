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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_vrfs_vrf_net_io_tunnels_net_io_tunnel

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
	proto.RegisterType((*PimTunnelInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_addrtype")
	proto.RegisterType((*PimTunnelInfoBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.vrfs.vrf.net_io_tunnels.net_io_tunnel.pim_tunnel_info_bag")
}

func init() { proto.RegisterFile("pim_tunnel_info_bag.proto", fileDescriptor_0876aa1a98c122bd) }

var fileDescriptor_0876aa1a98c122bd = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0x26, 0xef, 0x2b, 0xad, 0x9d, 0x56, 0xc1, 0x54, 0xb0, 0x3d, 0x59, 0x7b, 0xea, 0x29, 0x87,
	0xb6, 0xee, 0xdd, 0x83, 0x27, 0xa1, 0x60, 0x05, 0xc1, 0xd3, 0x90, 0xb6, 0xd9, 0x12, 0x70, 0x93,
	0x90, 0xa4, 0x8b, 0xbd, 0x78, 0xf4, 0x2a, 0xfe, 0x42, 0xff, 0x8a, 0x64, 0x77, 0xbb, 0x34, 0xd2,
	0xa3, 0xf4, 0xb2, 0x90, 0x99, 0x67, 0x9e, 0x8f, 0x61, 0x16, 0xfa, 0x46, 0x66, 0xe8, 0x37, 0x4a,
	0x89, 0x57, 0x94, 0x2a, 0xd5, 0xb8, 0xe0, 0x6b, 0x66, 0xac, 0xf6, 0x9a, 0x3e, 0x2e, 0xa5, 0x5b,
	0x6a, 0x94, 0xda, 0xe1, 0x9b, 0x45, 0x69, 0xf2, 0x29, 0x06, 0xb0, 0x36, 0xc2, 0x32, 0x69, 0xf2,
	0x24, 0xbc, 0x98, 0xf3, 0x5c, 0xad, 0x16, 0x5b, 0x96, 0xdb, 0xd4, 0x85, 0x0f, 0x53, 0xc2, 0xa3,
	0xd4, 0x15, 0xa5, 0x8b, 0x9f, 0xc3, 0x67, 0xe8, 0x1d, 0xd0, 0xc3, 0x87, 0xfb, 0x97, 0x27, 0xda,
	0x87, 0xd3, 0xdc, 0xa6, 0xa8, 0x78, 0x26, 0x7a, 0x64, 0x40, 0x46, 0xad, 0x79, 0x33, 0xb7, 0xe9,
	0x8c, 0x67, 0x82, 0x5e, 0x43, 0xbb, 0x1a, 0x29, 0xba, 0xff, 0x8a, 0x2e, 0x94, 0xa5, 0x00, 0x18,
	0x66, 0xd0, 0x09, 0xbc, 0x7c, 0xb5, 0xb2, 0x7e, 0x6b, 0x04, 0xbd, 0x82, 0x26, 0x8f, 0xa8, 0x1a,
	0xbc, 0x64, 0xba, 0x81, 0x4e, 0x11, 0x24, 0x20, 0x85, 0x73, 0x15, 0x55, 0x3b, 0xd4, 0xee, 0xca,
	0x52, 0x05, 0x49, 0x6a, 0xc8, 0xff, 0x1a, 0x92, 0x54, 0x90, 0xe1, 0xf7, 0x09, 0x74, 0x0f, 0xe4,
	0xa0, 0x1f, 0x04, 0xce, 0x9d, 0xde, 0xd8, 0xa5, 0xa8, 0xa7, 0xc7, 0x03, 0x32, 0x6a, 0x8f, 0x91,
	0xfd, 0xf9, 0x2e, 0xd9, 0x7e, 0xe0, 0xf9, 0x59, 0x29, 0xbb, 0xcb, 0xf0, 0x0e, 0x60, 0x4d, 0xed,
	0x61, 0x72, 0x1c, 0x0f, 0x2d, 0x6b, 0x76, 0xfa, 0x5f, 0x04, 0x2e, 0xe3, 0x45, 0xa0, 0x12, 0x5e,
	0xea, 0xde, 0xf4, 0x38, 0x56, 0x68, 0xb4, 0x8e, 0x59, 0x90, 0xa6, 0x9f, 0x04, 0xba, 0x6b, 0xab,
	0x37, 0xe6, 0x97, 0xa5, 0xdb, 0xe3, 0x58, 0xba, 0x28, 0xb4, 0x23, 0x47, 0xfb, 0x17, 0x9f, 0x44,
	0x17, 0xbf, 0x68, 0x14, 0xbf, 0xe0, 0xe4, 0x27, 0x00, 0x00, 0xff, 0xff, 0x93, 0xd1, 0x79, 0x48,
	0x9f, 0x03, 0x00, 0x00,
}
