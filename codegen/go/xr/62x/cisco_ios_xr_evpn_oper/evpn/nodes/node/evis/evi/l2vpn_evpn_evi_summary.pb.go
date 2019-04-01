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
// source: l2vpn_evpn_evi_summary.proto

package cisco_ios_xr_evpn_oper_evpn_nodes_node_evis_evi

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

type L2VpnEvpnEviSummary_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Evi                  uint32   `protobuf:"varint,2,opt,name=evi,proto3" json:"evi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEviSummary_KEYS) Reset()         { *m = L2VpnEvpnEviSummary_KEYS{} }
func (m *L2VpnEvpnEviSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviSummary_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnEviSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd70d64a7ad712df, []int{0}
}

func (m *L2VpnEvpnEviSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviSummary_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviSummary_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnEviSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviSummary_KEYS.Size(m)
}
func (m *L2VpnEvpnEviSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviSummary_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnEviSummary_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnEvpnEviSummary_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

type L2VpnEvpnEviSummary struct {
	EviXr                uint32   `protobuf:"varint,50,opt,name=evi_xr,json=eviXr,proto3" json:"evi_xr,omitempty"`
	BdName               string   `protobuf:"bytes,51,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	Type                 string   `protobuf:"bytes,52,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEviSummary) Reset()         { *m = L2VpnEvpnEviSummary{} }
func (m *L2VpnEvpnEviSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviSummary) ProtoMessage()    {}
func (*L2VpnEvpnEviSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd70d64a7ad712df, []int{1}
}

func (m *L2VpnEvpnEviSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviSummary.Merge(m, src)
}
func (m *L2VpnEvpnEviSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Size(m)
}
func (m *L2VpnEvpnEviSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviSummary proto.InternalMessageInfo

func (m *L2VpnEvpnEviSummary) GetEviXr() uint32 {
	if m != nil {
		return m.EviXr
	}
	return 0
}

func (m *L2VpnEvpnEviSummary) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

func (m *L2VpnEvpnEviSummary) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnEvpnEviSummary_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evis.evi.l2vpn_evpn_evi_summary_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evis.evi.l2vpn_evpn_evi_summary")
}

func init() { proto.RegisterFile("l2vpn_evpn_evi_summary.proto", fileDescriptor_dd70d64a7ad712df) }

var fileDescriptor_dd70d64a7ad712df = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x85, 0xe9, 0xff, 0x6b, 0xc5, 0x80, 0x20, 0x01, 0x35, 0xa0, 0x43, 0xe9, 0xd4, 0x29, 0x42,
	0xeb, 0x2b, 0x08, 0x8a, 0xe0, 0x50, 0x17, 0x05, 0x21, 0xb4, 0xcd, 0x1d, 0x02, 0xa6, 0x09, 0x49,
	0x0d, 0xed, 0xdb, 0x4b, 0xd2, 0xb5, 0xcb, 0xe1, 0x9c, 0xcb, 0xc7, 0x07, 0x17, 0x1d, 0x3e, 0xb9,
	0xd3, 0x2d, 0x83, 0x31, 0x04, 0xb3, 0x5f, 0x29, 0x2b, 0x33, 0x50, 0x6d, 0x54, 0xa7, 0xf0, 0xb1,
	0x11, 0xb6, 0x51, 0x4c, 0x28, 0xcb, 0x7a, 0x33, 0x42, 0x4a, 0x83, 0xa1, 0xbe, 0xd1, 0x56, 0x71,
	0xb0, 0x21, 0x29, 0x38, 0x61, 0x7d, 0xa4, 0x17, 0xb4, 0x9f, 0x16, 0xb2, 0xdb, 0xf9, 0xf5, 0xc0,
	0x3b, 0xb4, 0xf0, 0x3c, 0x13, 0x9c, 0x44, 0x49, 0x94, 0x2d, 0xcb, 0xd8, 0xcf, 0x2b, 0xc7, 0x6b,
	0xf4, 0x0f, 0x4e, 0x90, 0xbf, 0x24, 0xca, 0x56, 0xa5, 0xaf, 0xe9, 0x1b, 0x6d, 0xa7, 0x4d, 0x78,
	0x83, 0x62, 0x3f, 0x7b, 0x43, 0xf2, 0x80, 0xcf, 0xc1, 0x89, 0xa7, 0xf1, 0xee, 0x9a, 0xb3, 0xb6,
	0x92, 0x40, 0x8a, 0xd1, 0x5d, 0xf3, 0x7b, 0x25, 0x01, 0x63, 0x34, 0xeb, 0x06, 0x0d, 0xe4, 0x14,
	0xae, 0xa1, 0xd7, 0x71, 0xf8, 0xaf, 0xf8, 0x05, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x77, 0xf3, 0xfa,
	0xff, 0x00, 0x00, 0x00,
}
