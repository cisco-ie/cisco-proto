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
// source: l2vpn_evpn_tep.proto

package cisco_ios_xr_evpn_oper_evpn_standby_teps_tep

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

type L2VpnEvpnTep_KEYS struct {
	TepId                uint32   `protobuf:"varint,1,opt,name=tep_id,json=tepId,proto3" json:"tep_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnTep_KEYS) Reset()         { *m = L2VpnEvpnTep_KEYS{} }
func (m *L2VpnEvpnTep_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnTep_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnTep_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_510abc87a2dfeb81, []int{0}
}

func (m *L2VpnEvpnTep_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnTep_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnTep_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnTep_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnTep_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnTep_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnTep_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnTep_KEYS.Size(m)
}
func (m *L2VpnEvpnTep_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnTep_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnTep_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnTep_KEYS) GetTepId() uint32 {
	if m != nil {
		return m.TepId
	}
	return 0
}

type L2VpnEvpnTepInfo struct {
	EthernetVpnId        uint32   `protobuf:"varint,1,opt,name=ethernet_vpn_id,json=ethernetVpnId,proto3" json:"ethernet_vpn_id,omitempty"`
	Encapsulation        uint32   `protobuf:"varint,2,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnTepInfo) Reset()         { *m = L2VpnEvpnTepInfo{} }
func (m *L2VpnEvpnTepInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnTepInfo) ProtoMessage()    {}
func (*L2VpnEvpnTepInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_510abc87a2dfeb81, []int{1}
}

func (m *L2VpnEvpnTepInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnTepInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnTepInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnTepInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnTepInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnTepInfo.Merge(m, src)
}
func (m *L2VpnEvpnTepInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnTepInfo.Size(m)
}
func (m *L2VpnEvpnTepInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnTepInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnTepInfo proto.InternalMessageInfo

func (m *L2VpnEvpnTepInfo) GetEthernetVpnId() uint32 {
	if m != nil {
		return m.EthernetVpnId
	}
	return 0
}

func (m *L2VpnEvpnTepInfo) GetEncapsulation() uint32 {
	if m != nil {
		return m.Encapsulation
	}
	return 0
}

func (m *L2VpnEvpnTepInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type L2VpnEvpnTep struct {
	TunnelEndpointId     uint32            `protobuf:"varint,50,opt,name=tunnel_endpoint_id,json=tunnelEndpointId,proto3" json:"tunnel_endpoint_id,omitempty"`
	Type                 uint32            `protobuf:"varint,51,opt,name=type,proto3" json:"type,omitempty"`
	LocalInfo            *L2VpnEvpnTepInfo `protobuf:"bytes,52,opt,name=local_info,json=localInfo,proto3" json:"local_info,omitempty"`
	RemoteInfo           *L2VpnEvpnTepInfo `protobuf:"bytes,53,opt,name=remote_info,json=remoteInfo,proto3" json:"remote_info,omitempty"`
	UseCount             uint32            `protobuf:"varint,54,opt,name=use_count,json=useCount,proto3" json:"use_count,omitempty"`
	VrfName              string            `protobuf:"bytes,55,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	VrfTableId           uint32            `protobuf:"varint,56,opt,name=vrf_table_id,json=vrfTableId,proto3" json:"vrf_table_id,omitempty"`
	UdpPort              uint32            `protobuf:"varint,57,opt,name=udp_port,json=udpPort,proto3" json:"udp_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2VpnEvpnTep) Reset()         { *m = L2VpnEvpnTep{} }
func (m *L2VpnEvpnTep) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnTep) ProtoMessage()    {}
func (*L2VpnEvpnTep) Descriptor() ([]byte, []int) {
	return fileDescriptor_510abc87a2dfeb81, []int{2}
}

func (m *L2VpnEvpnTep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnTep.Unmarshal(m, b)
}
func (m *L2VpnEvpnTep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnTep.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnTep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnTep.Merge(m, src)
}
func (m *L2VpnEvpnTep) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnTep.Size(m)
}
func (m *L2VpnEvpnTep) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnTep.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnTep proto.InternalMessageInfo

func (m *L2VpnEvpnTep) GetTunnelEndpointId() uint32 {
	if m != nil {
		return m.TunnelEndpointId
	}
	return 0
}

func (m *L2VpnEvpnTep) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *L2VpnEvpnTep) GetLocalInfo() *L2VpnEvpnTepInfo {
	if m != nil {
		return m.LocalInfo
	}
	return nil
}

func (m *L2VpnEvpnTep) GetRemoteInfo() *L2VpnEvpnTepInfo {
	if m != nil {
		return m.RemoteInfo
	}
	return nil
}

func (m *L2VpnEvpnTep) GetUseCount() uint32 {
	if m != nil {
		return m.UseCount
	}
	return 0
}

func (m *L2VpnEvpnTep) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *L2VpnEvpnTep) GetVrfTableId() uint32 {
	if m != nil {
		return m.VrfTableId
	}
	return 0
}

func (m *L2VpnEvpnTep) GetUdpPort() uint32 {
	if m != nil {
		return m.UdpPort
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnEvpnTep_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.teps.tep.l2vpn_evpn_tep_KEYS")
	proto.RegisterType((*L2VpnEvpnTepInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.teps.tep.l2vpn_evpn_tep_info")
	proto.RegisterType((*L2VpnEvpnTep)(nil), "cisco_ios_xr_evpn_oper.evpn.standby.teps.tep.l2vpn_evpn_tep")
}

func init() { proto.RegisterFile("l2vpn_evpn_tep.proto", fileDescriptor_510abc87a2dfeb81) }

var fileDescriptor_510abc87a2dfeb81 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x8b, 0xdb, 0x30,
	0x10, 0xc5, 0x71, 0xd2, 0xe6, 0xcf, 0xa4, 0x49, 0x8b, 0xda, 0x82, 0x4a, 0x2f, 0x26, 0x94, 0x92,
	0x43, 0xf0, 0x21, 0xe9, 0xdf, 0x63, 0x29, 0x39, 0x98, 0x42, 0x29, 0x69, 0x29, 0xf4, 0xa4, 0xda,
	0xd6, 0x98, 0x0a, 0x1c, 0x49, 0x48, 0x63, 0xd3, 0x7c, 0x88, 0xfd, 0xce, 0x8b, 0xe4, 0xec, 0xb2,
	0x61, 0xf7, 0xb2, 0xb0, 0x17, 0x31, 0x9e, 0xdf, 0xcc, 0x7b, 0xcc, 0xc3, 0xf0, 0xa2, 0xd9, 0x74,
	0x56, 0x0b, 0x0c, 0x0f, 0xa1, 0xcd, 0xac, 0x33, 0x64, 0xd8, 0xba, 0x52, 0xbe, 0x32, 0x42, 0x19,
	0x2f, 0xfe, 0xbb, 0x1e, 0x1a, 0x8b, 0x2e, 0x0b, 0x55, 0xe6, 0xa9, 0xd0, 0xb2, 0x3c, 0x66, 0x84,
	0xd6, 0x87, 0x67, 0xb9, 0x86, 0xe7, 0xe7, 0x2a, 0xe2, 0xdb, 0xee, 0xcf, 0x4f, 0xf6, 0x12, 0x46,
	0xa1, 0x56, 0x92, 0x27, 0x69, 0xb2, 0x9a, 0xef, 0x1f, 0x13, 0xda, 0x5c, 0x2e, 0xfd, 0xad, 0x69,
	0xa5, 0x6b, 0xc3, 0xde, 0xc2, 0x53, 0xa4, 0x7f, 0xe8, 0x34, 0x92, 0x08, 0xe0, 0x7a, 0x6d, 0x7e,
	0xd5, 0xfe, 0x6d, 0x75, 0x2e, 0xd9, 0x1b, 0x98, 0xa3, 0xae, 0x0a, 0xeb, 0xdb, 0xa6, 0x20, 0x65,
	0x34, 0x1f, 0x9c, 0xa6, 0x6e, 0x36, 0xd9, 0x02, 0x06, 0xca, 0xf2, 0x61, 0x9a, 0xac, 0xa6, 0xfb,
	0x81, 0xb2, 0xcb, 0x8b, 0x21, 0x2c, 0xce, 0x5d, 0xd9, 0x1a, 0x18, 0xb5, 0x5a, 0x63, 0x23, 0x50,
	0x4b, 0x6b, 0x94, 0xa6, 0xe0, 0xb9, 0x89, 0x6a, 0xcf, 0x7a, 0xb2, 0x3b, 0x81, 0x5c, 0x32, 0x06,
	0x8f, 0xe8, 0x68, 0x91, 0x6f, 0x23, 0x8f, 0x35, 0xfb, 0x0b, 0xd0, 0x98, 0xaa, 0x68, 0xe2, 0x01,
	0xfc, 0x5d, 0x9a, 0xac, 0x66, 0x9b, 0x2f, 0xd9, 0x7d, 0xa2, 0xcb, 0xee, 0x48, 0x62, 0x3f, 0x8d,
	0xa2, 0x79, 0x08, 0xa5, 0x84, 0x99, 0xc3, 0x83, 0x21, 0xec, 0x2d, 0xde, 0x3f, 0x94, 0x05, 0xf4,
	0xaa, 0xd1, 0xe3, 0x35, 0x4c, 0x5b, 0x8f, 0xa2, 0x32, 0xad, 0x26, 0xfe, 0x21, 0x9e, 0x37, 0x69,
	0x3d, 0x7e, 0x0d, 0xdf, 0xec, 0x15, 0x4c, 0x3a, 0x57, 0x0b, 0x5d, 0x1c, 0x90, 0x7f, 0x8c, 0x69,
	0x8e, 0x3b, 0x57, 0x7f, 0x2f, 0x0e, 0xc8, 0x52, 0x78, 0x12, 0x10, 0x15, 0x65, 0x83, 0x21, 0xb9,
	0x4f, 0x71, 0x15, 0x3a, 0x57, 0xff, 0x0a, 0xad, 0x5c, 0x86, 0xe5, 0x56, 0x5a, 0x61, 0x8d, 0x23,
	0xfe, 0x39, 0xd2, 0x71, 0x2b, 0xed, 0x0f, 0xe3, 0xa8, 0x1c, 0xc5, 0xff, 0x6c, 0x7b, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x78, 0x3a, 0xfa, 0x76, 0x7f, 0x02, 0x00, 0x00,
}
