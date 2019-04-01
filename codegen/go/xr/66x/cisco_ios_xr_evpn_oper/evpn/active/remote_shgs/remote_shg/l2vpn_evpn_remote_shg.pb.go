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
// source: l2vpn_evpn_remote_shg.proto

package cisco_ios_xr_evpn_oper_evpn_active_remote_shgs_remote_shg

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

type L2VpnEvpnRemoteShg_KEYS struct {
	Evi                  uint32   `protobuf:"varint,1,opt,name=evi,proto3" json:"evi,omitempty"`
	Esi1                 string   `protobuf:"bytes,2,opt,name=esi1,proto3" json:"esi1,omitempty"`
	Esi2                 string   `protobuf:"bytes,3,opt,name=esi2,proto3" json:"esi2,omitempty"`
	Esi3                 string   `protobuf:"bytes,4,opt,name=esi3,proto3" json:"esi3,omitempty"`
	Esi4                 string   `protobuf:"bytes,5,opt,name=esi4,proto3" json:"esi4,omitempty"`
	Esi5                 string   `protobuf:"bytes,6,opt,name=esi5,proto3" json:"esi5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnRemoteShg_KEYS) Reset()         { *m = L2VpnEvpnRemoteShg_KEYS{} }
func (m *L2VpnEvpnRemoteShg_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnRemoteShg_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnRemoteShg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93b708cb223c0df, []int{0}
}

func (m *L2VpnEvpnRemoteShg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnRemoteShg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnRemoteShg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnRemoteShg_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS.Size(m)
}
func (m *L2VpnEvpnRemoteShg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnRemoteShg_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnRemoteShg_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnRemoteShg_KEYS) GetEsi1() string {
	if m != nil {
		return m.Esi1
	}
	return ""
}

func (m *L2VpnEvpnRemoteShg_KEYS) GetEsi2() string {
	if m != nil {
		return m.Esi2
	}
	return ""
}

func (m *L2VpnEvpnRemoteShg_KEYS) GetEsi3() string {
	if m != nil {
		return m.Esi3
	}
	return ""
}

func (m *L2VpnEvpnRemoteShg_KEYS) GetEsi4() string {
	if m != nil {
		return m.Esi4
	}
	return ""
}

func (m *L2VpnEvpnRemoteShg_KEYS) GetEsi5() string {
	if m != nil {
		return m.Esi5
	}
	return ""
}

type L2VpnEvpnRemoteShgInfo struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Label                uint32   `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnRemoteShgInfo) Reset()         { *m = L2VpnEvpnRemoteShgInfo{} }
func (m *L2VpnEvpnRemoteShgInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnRemoteShgInfo) ProtoMessage()    {}
func (*L2VpnEvpnRemoteShgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93b708cb223c0df, []int{1}
}

func (m *L2VpnEvpnRemoteShgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Merge(m, src)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Size(m)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnRemoteShgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnRemoteShgInfo proto.InternalMessageInfo

func (m *L2VpnEvpnRemoteShgInfo) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnEvpnRemoteShgInfo) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type L2VpnEvpnRemoteShg struct {
	EthernetSegmentIdentifier    []uint32                  `protobuf:"varint,50,rep,packed,name=ethernet_segment_identifier,json=ethernetSegmentIdentifier,proto3" json:"ethernet_segment_identifier,omitempty"`
	EthernetVpnId                uint32                    `protobuf:"varint,51,opt,name=ethernet_vpn_id,json=ethernetVpnId,proto3" json:"ethernet_vpn_id,omitempty"`
	Encapsulation                uint32                    `protobuf:"varint,52,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	RemoteSplitHorizonGroupLabel []*L2VpnEvpnRemoteShgInfo `protobuf:"bytes,53,rep,name=remote_split_horizon_group_label,json=remoteSplitHorizonGroupLabel,proto3" json:"remote_split_horizon_group_label,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                  `json:"-"`
	XXX_unrecognized             []byte                    `json:"-"`
	XXX_sizecache                int32                     `json:"-"`
}

func (m *L2VpnEvpnRemoteShg) Reset()         { *m = L2VpnEvpnRemoteShg{} }
func (m *L2VpnEvpnRemoteShg) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnRemoteShg) ProtoMessage()    {}
func (*L2VpnEvpnRemoteShg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d93b708cb223c0df, []int{2}
}

func (m *L2VpnEvpnRemoteShg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnRemoteShg.Unmarshal(m, b)
}
func (m *L2VpnEvpnRemoteShg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnRemoteShg.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnRemoteShg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnRemoteShg.Merge(m, src)
}
func (m *L2VpnEvpnRemoteShg) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnRemoteShg.Size(m)
}
func (m *L2VpnEvpnRemoteShg) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnRemoteShg.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnRemoteShg proto.InternalMessageInfo

func (m *L2VpnEvpnRemoteShg) GetEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.EthernetSegmentIdentifier
	}
	return nil
}

func (m *L2VpnEvpnRemoteShg) GetEthernetVpnId() uint32 {
	if m != nil {
		return m.EthernetVpnId
	}
	return 0
}

func (m *L2VpnEvpnRemoteShg) GetEncapsulation() uint32 {
	if m != nil {
		return m.Encapsulation
	}
	return 0
}

func (m *L2VpnEvpnRemoteShg) GetRemoteSplitHorizonGroupLabel() []*L2VpnEvpnRemoteShgInfo {
	if m != nil {
		return m.RemoteSplitHorizonGroupLabel
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnEvpnRemoteShg_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.remote_shgs.remote_shg.l2vpn_evpn_remote_shg_KEYS")
	proto.RegisterType((*L2VpnEvpnRemoteShgInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.active.remote_shgs.remote_shg.l2vpn_evpn_remote_shg_info")
	proto.RegisterType((*L2VpnEvpnRemoteShg)(nil), "cisco_ios_xr_evpn_oper.evpn.active.remote_shgs.remote_shg.l2vpn_evpn_remote_shg")
}

func init() { proto.RegisterFile("l2vpn_evpn_remote_shg.proto", fileDescriptor_d93b708cb223c0df) }

var fileDescriptor_d93b708cb223c0df = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0xba, 0xed, 0xff, 0x5f, 0xa4, 0x28, 0x41, 0x21, 0x73, 0x1e, 0xca, 0x10, 0xe9,
	0xa9, 0x60, 0xbb, 0x1d, 0xbc, 0x78, 0x13, 0x37, 0xd4, 0x4b, 0x87, 0x82, 0xa7, 0xd0, 0x75, 0xef,
	0xb6, 0x40, 0x97, 0x84, 0x24, 0x1b, 0xc3, 0x8f, 0xe1, 0xdd, 0x8f, 0xe3, 0xf7, 0x92, 0xb4, 0x5d,
	0xdd, 0x61, 0xbb, 0x78, 0x29, 0x4f, 0x7e, 0x3c, 0x79, 0x9f, 0xe7, 0x6d, 0x8b, 0x7a, 0x79, 0xb4,
	0x91, 0x9c, 0x82, 0x7d, 0x28, 0x58, 0x09, 0x03, 0x54, 0x2f, 0x17, 0xa1, 0x54, 0xc2, 0x08, 0x7c,
	0x97, 0x31, 0x9d, 0x09, 0xca, 0x84, 0xa6, 0x5b, 0x55, 0x7a, 0x84, 0x04, 0x15, 0x5a, 0x15, 0xa6,
	0x99, 0x61, 0x1b, 0x08, 0x7f, 0x2f, 0xe9, 0x3d, 0xdd, 0xff, 0x74, 0xd0, 0xe5, 0xc1, 0xd1, 0xf4,
	0xe9, 0xe1, 0x7d, 0x82, 0xcf, 0x90, 0x0b, 0x1b, 0x46, 0x1c, 0xdf, 0x09, 0xbc, 0xc4, 0x4a, 0x8c,
	0x51, 0x13, 0x34, 0xbb, 0x25, 0x0d, 0xdf, 0x09, 0x3a, 0x49, 0xa1, 0x2b, 0x16, 0x11, 0xb7, 0x66,
	0x51, 0xc5, 0x62, 0xd2, 0xac, 0x59, 0x5c, 0xb1, 0x01, 0x69, 0xd5, 0x6c, 0x50, 0xb1, 0x21, 0x69,
	0xd7, 0x6c, 0xd8, 0x7f, 0x39, 0xd6, 0x89, 0xf1, 0xb9, 0xc0, 0x5d, 0xf4, 0x9f, 0xc3, 0xd6, 0xd0,
	0xa5, 0x90, 0x45, 0xb1, 0x4e, 0xf2, 0xcf, 0x9e, 0x47, 0x42, 0xe2, 0x73, 0xd4, 0xca, 0xd3, 0x29,
	0xe4, 0x45, 0x3b, 0x2f, 0x29, 0x0f, 0xfd, 0xef, 0x06, 0xba, 0x38, 0x38, 0x0f, 0xdf, 0xa3, 0x1e,
	0x98, 0x25, 0x28, 0x0e, 0x86, 0x6a, 0x58, 0xac, 0x80, 0x1b, 0xca, 0x66, 0xc0, 0x0d, 0x9b, 0x33,
	0x50, 0x24, 0xf2, 0xdd, 0xc0, 0x4b, 0xba, 0x3b, 0xcb, 0xa4, 0x74, 0x8c, 0x6b, 0x03, 0xbe, 0x41,
	0xa7, 0xf5, 0x7d, 0x3b, 0x9a, 0xcd, 0x48, 0x5c, 0x24, 0x7b, 0x3b, 0xfc, 0x26, 0xf9, 0x78, 0x86,
	0xaf, 0x91, 0x07, 0x3c, 0x4b, 0xa5, 0x5e, 0xe7, 0xa9, 0x61, 0x82, 0x93, 0x41, 0xe5, 0xda, 0x87,
	0xf8, 0xcb, 0x41, 0xfe, 0xae, 0x9c, 0xcc, 0x99, 0xdd, 0x50, 0xb1, 0x0f, 0xc1, 0xe9, 0x42, 0x89,
	0xb5, 0xa4, 0xe5, 0x66, 0x43, 0xdf, 0x0d, 0x4e, 0xa2, 0xd7, 0xf0, 0xcf, 0x9f, 0x3c, 0x3c, 0xfe,
	0x6a, 0x93, 0xab, 0x12, 0x4c, 0x6c, 0xfa, 0xa8, 0x0c, 0x7f, 0xb4, 0xd9, 0xcf, 0x36, 0x7a, 0xda,
	0x2e, 0xfe, 0xb6, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x40, 0xb9, 0x07, 0xcc, 0x8c, 0x02, 0x00,
	0x00,
}
