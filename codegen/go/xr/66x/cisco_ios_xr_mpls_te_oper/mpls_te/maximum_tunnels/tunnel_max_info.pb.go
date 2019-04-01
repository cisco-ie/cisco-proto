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
// source: tunnel_max_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_maximum_tunnels

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

type TunnelMaxInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TunnelMaxInfo_KEYS) Reset()         { *m = TunnelMaxInfo_KEYS{} }
func (m *TunnelMaxInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*TunnelMaxInfo_KEYS) ProtoMessage()    {}
func (*TunnelMaxInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ca435800aabbc5, []int{0}
}

func (m *TunnelMaxInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelMaxInfo_KEYS.Unmarshal(m, b)
}
func (m *TunnelMaxInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelMaxInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *TunnelMaxInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelMaxInfo_KEYS.Merge(m, src)
}
func (m *TunnelMaxInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_TunnelMaxInfo_KEYS.Size(m)
}
func (m *TunnelMaxInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelMaxInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelMaxInfo_KEYS proto.InternalMessageInfo

type TunnelMaxInfo struct {
	CurrentMaximumTunnels                 uint32   `protobuf:"varint,50,opt,name=current_maximum_tunnels,json=currentMaximumTunnels,proto3" json:"current_maximum_tunnels,omitempty"`
	CurrentMaximumDestinations            uint32   `protobuf:"varint,51,opt,name=current_maximum_destinations,json=currentMaximumDestinations,proto3" json:"current_maximum_destinations,omitempty"`
	CurrentMaximumP2MpTunnels             uint32   `protobuf:"varint,52,opt,name=current_maximum_p2mp_tunnels,json=currentMaximumP2mpTunnels,proto3" json:"current_maximum_p2mp_tunnels,omitempty"`
	CurrentMaximumP2MpDestnationPerTunnel uint32   `protobuf:"varint,53,opt,name=current_maximum_p2mp_destnation_per_tunnel,json=currentMaximumP2mpDestnationPerTunnel,proto3" json:"current_maximum_p2mp_destnation_per_tunnel,omitempty"`
	AbsoluteMaximum                       uint32   `protobuf:"varint,54,opt,name=absolute_maximum,json=absoluteMaximum,proto3" json:"absolute_maximum,omitempty"`
	CurrentMaximumAutobackups             uint32   `protobuf:"varint,55,opt,name=current_maximum_autobackups,json=currentMaximumAutobackups,proto3" json:"current_maximum_autobackups,omitempty"`
	CurrentMaximumAutoMesh                uint32   `protobuf:"varint,56,opt,name=current_maximum_auto_mesh,json=currentMaximumAutoMesh,proto3" json:"current_maximum_auto_mesh,omitempty"`
	CurrentTunnelCount                    uint32   `protobuf:"varint,57,opt,name=current_tunnel_count,json=currentTunnelCount,proto3" json:"current_tunnel_count,omitempty"`
	CurrentDestinationCount               uint32   `protobuf:"varint,58,opt,name=current_destination_count,json=currentDestinationCount,proto3" json:"current_destination_count,omitempty"`
	CurrentP2MpTunnelCount                uint32   `protobuf:"varint,59,opt,name=current_p2mp_tunnel_count,json=currentP2mpTunnelCount,proto3" json:"current_p2mp_tunnel_count,omitempty"`
	CurrentP2MpDestnationPerTunnelCount   uint32   `protobuf:"varint,60,opt,name=current_p2mp_destnation_per_tunnel_count,json=currentP2mpDestnationPerTunnelCount,proto3" json:"current_p2mp_destnation_per_tunnel_count,omitempty"`
	IsAutobackupRangeConfigured           bool     `protobuf:"varint,61,opt,name=is_autobackup_range_configured,json=isAutobackupRangeConfigured,proto3" json:"is_autobackup_range_configured,omitempty"`
	CurrentAutobackups                    uint32   `protobuf:"varint,62,opt,name=current_autobackups,json=currentAutobackups,proto3" json:"current_autobackups,omitempty"`
	IsAutoMeshRangeConfigured             bool     `protobuf:"varint,63,opt,name=is_auto_mesh_range_configured,json=isAutoMeshRangeConfigured,proto3" json:"is_auto_mesh_range_configured,omitempty"`
	CurrentAutoMesh                       uint32   `protobuf:"varint,64,opt,name=current_auto_mesh,json=currentAutoMesh,proto3" json:"current_auto_mesh,omitempty"`
	CurrentMaximumGmplsUnitunnels         uint32   `protobuf:"varint,65,opt,name=current_maximum_gmpls_unitunnels,json=currentMaximumGmplsUnitunnels,proto3" json:"current_maximum_gmpls_unitunnels,omitempty"`
	CurrentGmplsUni                       uint32   `protobuf:"varint,66,opt,name=current_gmpls_uni,json=currentGmplsUni,proto3" json:"current_gmpls_uni,omitempty"`
	XXX_NoUnkeyedLiteral                  struct{} `json:"-"`
	XXX_unrecognized                      []byte   `json:"-"`
	XXX_sizecache                         int32    `json:"-"`
}

func (m *TunnelMaxInfo) Reset()         { *m = TunnelMaxInfo{} }
func (m *TunnelMaxInfo) String() string { return proto.CompactTextString(m) }
func (*TunnelMaxInfo) ProtoMessage()    {}
func (*TunnelMaxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_23ca435800aabbc5, []int{1}
}

func (m *TunnelMaxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TunnelMaxInfo.Unmarshal(m, b)
}
func (m *TunnelMaxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TunnelMaxInfo.Marshal(b, m, deterministic)
}
func (m *TunnelMaxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TunnelMaxInfo.Merge(m, src)
}
func (m *TunnelMaxInfo) XXX_Size() int {
	return xxx_messageInfo_TunnelMaxInfo.Size(m)
}
func (m *TunnelMaxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TunnelMaxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TunnelMaxInfo proto.InternalMessageInfo

func (m *TunnelMaxInfo) GetCurrentMaximumTunnels() uint32 {
	if m != nil {
		return m.CurrentMaximumTunnels
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumDestinations() uint32 {
	if m != nil {
		return m.CurrentMaximumDestinations
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumP2MpTunnels() uint32 {
	if m != nil {
		return m.CurrentMaximumP2MpTunnels
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumP2MpDestnationPerTunnel() uint32 {
	if m != nil {
		return m.CurrentMaximumP2MpDestnationPerTunnel
	}
	return 0
}

func (m *TunnelMaxInfo) GetAbsoluteMaximum() uint32 {
	if m != nil {
		return m.AbsoluteMaximum
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumAutobackups() uint32 {
	if m != nil {
		return m.CurrentMaximumAutobackups
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumAutoMesh() uint32 {
	if m != nil {
		return m.CurrentMaximumAutoMesh
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentTunnelCount() uint32 {
	if m != nil {
		return m.CurrentTunnelCount
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentDestinationCount() uint32 {
	if m != nil {
		return m.CurrentDestinationCount
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentP2MpTunnelCount() uint32 {
	if m != nil {
		return m.CurrentP2MpTunnelCount
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentP2MpDestnationPerTunnelCount() uint32 {
	if m != nil {
		return m.CurrentP2MpDestnationPerTunnelCount
	}
	return 0
}

func (m *TunnelMaxInfo) GetIsAutobackupRangeConfigured() bool {
	if m != nil {
		return m.IsAutobackupRangeConfigured
	}
	return false
}

func (m *TunnelMaxInfo) GetCurrentAutobackups() uint32 {
	if m != nil {
		return m.CurrentAutobackups
	}
	return 0
}

func (m *TunnelMaxInfo) GetIsAutoMeshRangeConfigured() bool {
	if m != nil {
		return m.IsAutoMeshRangeConfigured
	}
	return false
}

func (m *TunnelMaxInfo) GetCurrentAutoMesh() uint32 {
	if m != nil {
		return m.CurrentAutoMesh
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentMaximumGmplsUnitunnels() uint32 {
	if m != nil {
		return m.CurrentMaximumGmplsUnitunnels
	}
	return 0
}

func (m *TunnelMaxInfo) GetCurrentGmplsUni() uint32 {
	if m != nil {
		return m.CurrentGmplsUni
	}
	return 0
}

func init() {
	proto.RegisterType((*TunnelMaxInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.maximum_tunnels.tunnel_max_info_KEYS")
	proto.RegisterType((*TunnelMaxInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.maximum_tunnels.tunnel_max_info")
}

func init() { proto.RegisterFile("tunnel_max_info.proto", fileDescriptor_23ca435800aabbc5) }

var fileDescriptor_23ca435800aabbc5 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x4b, 0x6f, 0x13, 0x31,
	0x14, 0x85, 0xd5, 0x0d, 0xaa, 0xae, 0x84, 0x0a, 0xa6, 0x2d, 0x0e, 0xa5, 0x28, 0x2a, 0x42, 0x0a,
	0x5d, 0x04, 0x48, 0xa1, 0xd0, 0x02, 0x6d, 0x4a, 0x8a, 0xba, 0x40, 0x95, 0xaa, 0x40, 0x17, 0x5d,
	0x59, 0x93, 0xa9, 0x9b, 0x58, 0x64, 0x6c, 0xcb, 0x0f, 0x29, 0xbf, 0x88, 0xdf, 0x89, 0xc6, 0x8f,
	0x89, 0xe7, 0xb1, 0x4c, 0xee, 0x3d, 0xdf, 0xb9, 0xf6, 0xb9, 0x1e, 0xd8, 0x31, 0x96, 0x73, 0xba,
	0x24, 0x45, 0xb6, 0x22, 0x8c, 0x3f, 0x88, 0xa1, 0x54, 0xc2, 0x08, 0xf4, 0x21, 0x67, 0x3a, 0x17,
	0x84, 0x09, 0x4d, 0x56, 0x8a, 0x14, 0x72, 0xa9, 0x89, 0xa1, 0x44, 0x48, 0xaa, 0x86, 0xe1, 0xc7,
	0xb0, 0xc8, 0x56, 0xac, 0xb0, 0x05, 0xf1, 0x00, 0x7d, 0xb0, 0x0b, 0xdb, 0x0d, 0x16, 0xf9, 0xf5,
	0xf3, 0xee, 0xf7, 0xc1, 0xbf, 0x4d, 0xd8, 0x6a, 0x14, 0xd0, 0x31, 0x3c, 0xcf, 0xad, 0x52, 0x94,
	0x1b, 0xd2, 0xc0, 0xe0, 0x51, 0x7f, 0x63, 0xf0, 0x78, 0xba, 0x13, 0xca, 0xd7, 0xbe, 0xfa, 0xc7,
	0x17, 0xd1, 0x18, 0x5e, 0x36, 0x75, 0xf7, 0x54, 0x1b, 0xc6, 0x33, 0xc3, 0x04, 0xd7, 0xf8, 0xc8,
	0x89, 0x5f, 0xd4, 0xc5, 0x97, 0x49, 0x07, 0x3a, 0x6f, 0x13, 0xe4, 0xa8, 0x90, 0x95, 0xfd, 0x47,
	0x47, 0xe8, 0xd5, 0x09, 0x37, 0xa3, 0x42, 0xc6, 0x11, 0xee, 0xe0, 0xb0, 0x13, 0x50, 0xce, 0xe1,
	0x4d, 0x88, 0xa4, 0x2a, 0xf0, 0xf0, 0x27, 0x87, 0x7b, 0xd3, 0xc6, 0x5d, 0x56, 0xed, 0x37, 0x54,
	0x79, 0x36, 0x7a, 0x0b, 0x4f, 0xb2, 0x99, 0x16, 0x4b, 0x6b, 0x68, 0x64, 0xe3, 0x63, 0x07, 0xd8,
	0x8a, 0xff, 0x07, 0x02, 0x3a, 0x83, 0xbd, 0xe6, 0x14, 0x99, 0x35, 0x62, 0x96, 0xe5, 0x7f, 0xad,
	0xd4, 0xf8, 0x73, 0xd7, 0x29, 0x2e, 0xd6, 0x0d, 0xe8, 0x04, 0x7a, 0x5d, 0x7a, 0x52, 0x50, 0xbd,
	0xc0, 0x5f, 0x9c, 0x7a, 0xb7, 0xad, 0xbe, 0xa6, 0x7a, 0x81, 0xde, 0xc3, 0x76, 0x94, 0x86, 0x58,
	0x73, 0x61, 0xb9, 0xc1, 0x27, 0x4e, 0x85, 0x42, 0xcd, 0x1f, 0x69, 0x52, 0x56, 0xd0, 0xe9, 0xda,
	0x2c, 0x49, 0x2b, 0xc8, 0x4e, 0x9d, 0x2c, 0xae, 0x43, 0x92, 0x95, 0xd7, 0x26, 0x83, 0x26, 0x39,
	0x05, 0xed, 0xd7, 0xda, 0xa0, 0xeb, 0x94, 0xbc, 0xf4, 0x16, 0x06, 0x35, 0x69, 0x67, 0x42, 0x81,
	0xf4, 0xcd, 0x91, 0x5e, 0x27, 0xa4, 0x8e, 0x80, 0x3c, 0x76, 0x02, 0xaf, 0x98, 0x4e, 0x6e, 0x9b,
	0xa8, 0x8c, 0xcf, 0x29, 0xc9, 0x05, 0x7f, 0x60, 0x73, 0xab, 0xe8, 0x3d, 0xfe, 0xde, 0xdf, 0x18,
	0x6c, 0x4e, 0xf7, 0x98, 0x5e, 0xdf, 0xf8, 0xb4, 0xec, 0x99, 0x54, 0x2d, 0xe8, 0x1d, 0x3c, 0x8b,
	0xb3, 0xa5, 0xb9, 0x9d, 0xd5, 0xee, 0x30, 0x0d, 0x6c, 0x0c, 0xfb, 0xc1, 0xd5, 0x65, 0xd4, 0x36,
	0x3d, 0x77, 0xa6, 0x3d, 0x6f, 0x5a, 0x06, 0xd5, 0xb4, 0x3c, 0x84, 0xa7, 0xa9, 0xa5, 0x8f, 0x7a,
	0xec, 0xd7, 0x2b, 0x31, 0x74, 0x19, 0x5f, 0x41, 0xbf, 0xb9, 0x1e, 0x73, 0xf7, 0xee, 0x2d, 0x67,
	0xf1, 0xa5, 0x5c, 0x38, 0xe9, 0x7e, 0x7d, 0x4b, 0xae, 0xca, 0xae, 0xdb, 0xaa, 0x29, 0x35, 0xad,
	0x00, 0xf8, 0x47, 0xcd, 0x34, 0x4a, 0x66, 0x8f, 0xdc, 0xa7, 0xe7, 0xe8, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1f, 0x02, 0x26, 0x1f, 0x93, 0x04, 0x00, 0x00,
}
