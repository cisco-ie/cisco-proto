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

package cisco_ios_xr_evpn_oper_evpn_nodes_node_remote_shgs_remote_shg

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
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
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

func (m *L2VpnEvpnRemoteShg_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
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
	RemoteSplitHorizonGroupLabel []*L2VpnEvpnRemoteShgInfo `protobuf:"bytes,51,rep,name=remote_split_horizon_group_label,json=remoteSplitHorizonGroupLabel,proto3" json:"remote_split_horizon_group_label,omitempty"`
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

func (m *L2VpnEvpnRemoteShg) GetRemoteSplitHorizonGroupLabel() []*L2VpnEvpnRemoteShgInfo {
	if m != nil {
		return m.RemoteSplitHorizonGroupLabel
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnEvpnRemoteShg_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.remote_shgs.remote_shg.l2vpn_evpn_remote_shg_KEYS")
	proto.RegisterType((*L2VpnEvpnRemoteShgInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.remote_shgs.remote_shg.l2vpn_evpn_remote_shg_info")
	proto.RegisterType((*L2VpnEvpnRemoteShg)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.remote_shgs.remote_shg.l2vpn_evpn_remote_shg")
}

func init() { proto.RegisterFile("l2vpn_evpn_remote_shg.proto", fileDescriptor_d93b708cb223c0df) }

var fileDescriptor_d93b708cb223c0df = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x15, 0xfa, 0x0f, 0x8c, 0xba, 0x58, 0x20, 0x5c, 0xca, 0x10, 0x75, 0xea, 0x64, 0x89,
	0xa4, 0x1d, 0x61, 0x43, 0xb4, 0x02, 0x96, 0x74, 0xea, 0x64, 0xd1, 0xe6, 0xda, 0x5a, 0x4a, 0x7d,
	0x96, 0x6d, 0x50, 0xc5, 0x9b, 0x30, 0xf2, 0x9c, 0x2c, 0xc8, 0x71, 0x1a, 0x18, 0xda, 0x89, 0xc5,
	0xba, 0xfb, 0xf4, 0xf3, 0xdd, 0xa7, 0xc4, 0xa4, 0x5f, 0x24, 0xef, 0x5a, 0x09, 0xf0, 0x87, 0x81,
	0x2d, 0x3a, 0x10, 0x76, 0xb3, 0xe6, 0xda, 0xa0, 0x43, 0x7a, 0xb7, 0x94, 0x76, 0x89, 0x42, 0xa2,
	0x15, 0x3b, 0x13, 0x32, 0xa8, 0xc1, 0x70, 0x5f, 0x71, 0x85, 0x39, 0xd8, 0xf2, 0xe4, 0xbf, 0x17,
	0xed, 0x9f, 0x7a, 0xf0, 0x19, 0x91, 0xeb, 0x83, 0xe3, 0xc5, 0xd3, 0xc3, 0x7c, 0x46, 0xaf, 0x48,
	0xc7, 0xdf, 0x16, 0x32, 0x67, 0x51, 0x1c, 0x0d, 0xcf, 0xb2, 0xb6, 0x6f, 0xa7, 0x39, 0xa5, 0xa4,
	0x09, 0x56, 0xde, 0xb2, 0x93, 0x92, 0x96, 0x75, 0xc5, 0x12, 0xd6, 0xa8, 0x59, 0x52, 0xb1, 0x94,
	0x35, 0x6b, 0x96, 0x56, 0x6c, 0xc4, 0x5a, 0x35, 0x1b, 0x55, 0x6c, 0xcc, 0xda, 0x35, 0x1b, 0x0f,
	0x5e, 0x8e, 0xa9, 0x49, 0xb5, 0x42, 0xda, 0x23, 0xa7, 0x0a, 0x76, 0x4e, 0x6c, 0x50, 0x57, 0x6e,
	0x1d, 0xdf, 0x4f, 0x50, 0xd3, 0x0b, 0xd2, 0x2a, 0x5e, 0x17, 0x50, 0x94, 0x76, 0xdd, 0x2c, 0x34,
	0x83, 0xef, 0x88, 0x5c, 0x1e, 0x9c, 0x47, 0xef, 0x49, 0x1f, 0xdc, 0x06, 0x8c, 0x02, 0x27, 0x2c,
	0xac, 0xb7, 0xa0, 0x9c, 0x90, 0x39, 0x28, 0x27, 0x57, 0x12, 0x0c, 0x4b, 0xe2, 0xc6, 0xb0, 0x9b,
	0xf5, 0xf6, 0x91, 0x59, 0x48, 0x4c, 0xeb, 0x00, 0xfd, 0x8a, 0x48, 0xbc, 0x1f, 0xa7, 0x0b, 0xe9,
	0x9d, 0x8c, 0xfc, 0x40, 0x25, 0xd6, 0x06, 0xdf, 0xb4, 0x08, 0x2e, 0x69, 0xdc, 0x18, 0x9e, 0x27,
	0x73, 0xfe, 0xaf, 0xff, 0xc5, 0x8f, 0x7f, 0x90, 0xec, 0x26, 0x80, 0x99, 0x37, 0x98, 0x04, 0x81,
	0x47, 0xbf, 0xff, 0xd9, 0xaf, 0x5f, 0xb4, 0xcb, 0xe7, 0x92, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x55, 0x23, 0x85, 0x00, 0x4d, 0x02, 0x00, 0x00,
}