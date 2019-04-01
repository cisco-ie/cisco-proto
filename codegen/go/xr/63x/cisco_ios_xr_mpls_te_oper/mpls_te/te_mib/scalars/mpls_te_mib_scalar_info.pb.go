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
// source: mpls_te_mib_scalar_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_te_mib_scalars

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

type MplsTeMibScalarInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeMibScalarInfo_KEYS) Reset()         { *m = MplsTeMibScalarInfo_KEYS{} }
func (m *MplsTeMibScalarInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeMibScalarInfo_KEYS) ProtoMessage()    {}
func (*MplsTeMibScalarInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb7d7db37913ebf0, []int{0}
}

func (m *MplsTeMibScalarInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeMibScalarInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeMibScalarInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeMibScalarInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeMibScalarInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeMibScalarInfo_KEYS.Merge(m, src)
}
func (m *MplsTeMibScalarInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeMibScalarInfo_KEYS.Size(m)
}
func (m *MplsTeMibScalarInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeMibScalarInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeMibScalarInfo_KEYS proto.InternalMessageInfo

type MplsTeMibScalarInfo struct {
	MplsTunnelConfigured          uint32   `protobuf:"varint,50,opt,name=mpls_tunnel_configured,json=mplsTunnelConfigured,proto3" json:"mpls_tunnel_configured,omitempty"`
	MplsTunnelActive              uint32   `protobuf:"varint,51,opt,name=mpls_tunnel_active,json=mplsTunnelActive,proto3" json:"mpls_tunnel_active,omitempty"`
	MplsTunnelTeDistProto         uint32   `protobuf:"varint,52,opt,name=mpls_tunnel_te_dist_proto,json=mplsTunnelTeDistProto,proto3" json:"mpls_tunnel_te_dist_proto,omitempty"`
	MplsTunnelMaxHops             uint32   `protobuf:"varint,53,opt,name=mpls_tunnel_max_hops,json=mplsTunnelMaxHops,proto3" json:"mpls_tunnel_max_hops,omitempty"`
	MplsTunnelNotificationMaxRate uint32   `protobuf:"varint,54,opt,name=mpls_tunnel_notification_max_rate,json=mplsTunnelNotificationMaxRate,proto3" json:"mpls_tunnel_notification_max_rate,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *MplsTeMibScalarInfo) Reset()         { *m = MplsTeMibScalarInfo{} }
func (m *MplsTeMibScalarInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeMibScalarInfo) ProtoMessage()    {}
func (*MplsTeMibScalarInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb7d7db37913ebf0, []int{1}
}

func (m *MplsTeMibScalarInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeMibScalarInfo.Unmarshal(m, b)
}
func (m *MplsTeMibScalarInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeMibScalarInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeMibScalarInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeMibScalarInfo.Merge(m, src)
}
func (m *MplsTeMibScalarInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeMibScalarInfo.Size(m)
}
func (m *MplsTeMibScalarInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeMibScalarInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeMibScalarInfo proto.InternalMessageInfo

func (m *MplsTeMibScalarInfo) GetMplsTunnelConfigured() uint32 {
	if m != nil {
		return m.MplsTunnelConfigured
	}
	return 0
}

func (m *MplsTeMibScalarInfo) GetMplsTunnelActive() uint32 {
	if m != nil {
		return m.MplsTunnelActive
	}
	return 0
}

func (m *MplsTeMibScalarInfo) GetMplsTunnelTeDistProto() uint32 {
	if m != nil {
		return m.MplsTunnelTeDistProto
	}
	return 0
}

func (m *MplsTeMibScalarInfo) GetMplsTunnelMaxHops() uint32 {
	if m != nil {
		return m.MplsTunnelMaxHops
	}
	return 0
}

func (m *MplsTeMibScalarInfo) GetMplsTunnelNotificationMaxRate() uint32 {
	if m != nil {
		return m.MplsTunnelNotificationMaxRate
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeMibScalarInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.te_mib.scalars.mpls_te_mib_scalar_info_KEYS")
	proto.RegisterType((*MplsTeMibScalarInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.te_mib.scalars.mpls_te_mib_scalar_info")
}

func init() { proto.RegisterFile("mpls_te_mib_scalar_info.proto", fileDescriptor_bb7d7db37913ebf0) }

var fileDescriptor_bb7d7db37913ebf0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0xf4, 0xe0, 0x21, 0x20, 0x68, 0xa8, 0xba, 0x82, 0x15, 0xed, 0xc9, 0x83, 0xac,
	0x62, 0xab, 0x78, 0x15, 0x15, 0x0a, 0x52, 0x91, 0xda, 0x8b, 0xa7, 0x21, 0x4d, 0x53, 0x1d, 0xd8,
	0xcd, 0x84, 0x64, 0x2a, 0xfb, 0x5c, 0x3e, 0xa1, 0x38, 0xab, 0x6e, 0x2e, 0x3d, 0x26, 0xff, 0xff,
	0x85, 0xcc, 0xa8, 0x7e, 0x1d, 0xaa, 0x04, 0xec, 0xa0, 0xc6, 0x39, 0x24, 0x6b, 0x2a, 0x13, 0x01,
	0xfd, 0x92, 0xca, 0x10, 0x89, 0x49, 0x5f, 0x5a, 0x4c, 0x96, 0x00, 0x29, 0x41, 0x13, 0xe1, 0xaf,
	0x4b, 0xc1, 0xc5, 0xf2, 0xf7, 0x50, 0xb6, 0xb6, 0x6c, 0x6d, 0x1a, 0x1c, 0xab, 0xa3, 0x35, 0x4f,
	0xc2, 0xd3, 0xe3, 0xdb, 0xeb, 0xe0, 0x6b, 0x53, 0x1d, 0xac, 0x29, 0xe8, 0x91, 0xda, 0x6f, 0xa3,
	0x95, 0xf7, 0xae, 0x02, 0x4b, 0x7e, 0x89, 0xef, 0xab, 0xe8, 0x16, 0xc5, 0xd5, 0xc9, 0xc6, 0xd9,
	0xf6, 0xb4, 0xf7, 0x93, 0xce, 0x24, 0xbc, 0xff, 0xcf, 0xf4, 0xb9, 0xd2, 0xb9, 0x32, 0x96, 0xf1,
	0xd3, 0x15, 0x43, 0x11, 0x3b, 0x9d, 0xb8, 0x93, 0x7b, 0x7d, 0xab, 0x0e, 0xf3, 0x36, 0x3b, 0x58,
	0x60, 0x62, 0x90, 0x71, 0x8b, 0x91, 0xa0, 0xbd, 0x0e, 0xcd, 0xdc, 0x03, 0x26, 0x7e, 0x91, 0x5d,
	0x5c, 0xa8, 0x5e, 0x2e, 0x6b, 0xd3, 0xc0, 0x07, 0x85, 0x54, 0x5c, 0x0b, 0xda, 0xed, 0xd0, 0xc4,
	0x34, 0x63, 0x0a, 0x49, 0x8f, 0xd5, 0x69, 0x0e, 0x3c, 0x31, 0x2e, 0xd1, 0x1a, 0x46, 0xf2, 0xa2,
	0xa3, 0x61, 0x57, 0xdc, 0x88, 0xee, 0x77, 0xfa, 0x39, 0xab, 0x4d, 0x4c, 0x33, 0x35, 0xec, 0xe6,
	0x5b, 0xf2, 0xbd, 0xe1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x4f, 0x21, 0xff, 0xae, 0x01,
	0x00, 0x00,
}