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
// source: l2vpn_evpn_ead.proto

package cisco_ios_xr_evpn_oper_evpn_nodes_node_evi_detail_evi_children_ethernet_auto_discoveries_ethernet_auto_discovery

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

type L2VpnEvpnEad_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Evi                  uint32   `protobuf:"varint,2,opt,name=evi,proto3" json:"evi,omitempty"`
	Esi1                 string   `protobuf:"bytes,3,opt,name=esi1,proto3" json:"esi1,omitempty"`
	Esi2                 string   `protobuf:"bytes,4,opt,name=esi2,proto3" json:"esi2,omitempty"`
	Esi3                 string   `protobuf:"bytes,5,opt,name=esi3,proto3" json:"esi3,omitempty"`
	Esi4                 string   `protobuf:"bytes,6,opt,name=esi4,proto3" json:"esi4,omitempty"`
	Esi5                 string   `protobuf:"bytes,7,opt,name=esi5,proto3" json:"esi5,omitempty"`
	EthernetTag          uint32   `protobuf:"varint,8,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEad_KEYS) Reset()         { *m = L2VpnEvpnEad_KEYS{} }
func (m *L2VpnEvpnEad_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEad_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnEad_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b249f33df3bc066, []int{0}
}

func (m *L2VpnEvpnEad_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEad_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnEad_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEad_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEad_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEad_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnEad_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEad_KEYS.Size(m)
}
func (m *L2VpnEvpnEad_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEad_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEad_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnEad_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnEad_KEYS) GetEsi1() string {
	if m != nil {
		return m.Esi1
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEsi2() string {
	if m != nil {
		return m.Esi2
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEsi3() string {
	if m != nil {
		return m.Esi3
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEsi4() string {
	if m != nil {
		return m.Esi4
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEsi5() string {
	if m != nil {
		return m.Esi5
	}
	return ""
}

func (m *L2VpnEvpnEad_KEYS) GetEthernetTag() uint32 {
	if m != nil {
		return m.EthernetTag
	}
	return 0
}

type L2VpnLabelPathBuffer struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	OutputLabel          uint32   `protobuf:"varint,2,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLabelPathBuffer) Reset()         { *m = L2VpnLabelPathBuffer{} }
func (m *L2VpnLabelPathBuffer) String() string { return proto.CompactTextString(m) }
func (*L2VpnLabelPathBuffer) ProtoMessage()    {}
func (*L2VpnLabelPathBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b249f33df3bc066, []int{1}
}

func (m *L2VpnLabelPathBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Unmarshal(m, b)
}
func (m *L2VpnLabelPathBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Marshal(b, m, deterministic)
}
func (m *L2VpnLabelPathBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLabelPathBuffer.Merge(m, src)
}
func (m *L2VpnLabelPathBuffer) XXX_Size() int {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Size(m)
}
func (m *L2VpnLabelPathBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLabelPathBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLabelPathBuffer proto.InternalMessageInfo

func (m *L2VpnLabelPathBuffer) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnLabelPathBuffer) GetOutputLabel() uint32 {
	if m != nil {
		return m.OutputLabel
	}
	return 0
}

type L2VpnEvpnEad struct {
	EthernetVpnid             uint32                  `protobuf:"varint,50,opt,name=ethernet_vpnid,json=ethernetVpnid,proto3" json:"ethernet_vpnid,omitempty"`
	Type                      string                  `protobuf:"bytes,51,opt,name=type,proto3" json:"type,omitempty"`
	EthernetSegmentIdentifier []uint32                `protobuf:"varint,52,rep,packed,name=ethernet_segment_identifier,json=ethernetSegmentIdentifier,proto3" json:"ethernet_segment_identifier,omitempty"`
	EthernetTagXr             uint32                  `protobuf:"varint,53,opt,name=ethernet_tag_xr,json=ethernetTagXr,proto3" json:"ethernet_tag_xr,omitempty"`
	LocalNextHop              string                  `protobuf:"bytes,54,opt,name=local_next_hop,json=localNextHop,proto3" json:"local_next_hop,omitempty"`
	LocalLabel                uint32                  `protobuf:"varint,55,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	IsLocalEad                bool                    `protobuf:"varint,56,opt,name=is_local_ead,json=isLocalEad,proto3" json:"is_local_ead,omitempty"`
	Encap                     uint32                  `protobuf:"varint,57,opt,name=encap,proto3" json:"encap,omitempty"`
	RedundancySingleActive    bool                    `protobuf:"varint,58,opt,name=redundancy_single_active,json=redundancySingleActive,proto3" json:"redundancy_single_active,omitempty"`
	NumPaths                  uint32                  `protobuf:"varint,59,opt,name=num_paths,json=numPaths,proto3" json:"num_paths,omitempty"`
	PathBuffer                []*L2VpnLabelPathBuffer `protobuf:"bytes,60,rep,name=path_buffer,json=pathBuffer,proto3" json:"path_buffer,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                `json:"-"`
	XXX_unrecognized          []byte                  `json:"-"`
	XXX_sizecache             int32                   `json:"-"`
}

func (m *L2VpnEvpnEad) Reset()         { *m = L2VpnEvpnEad{} }
func (m *L2VpnEvpnEad) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEad) ProtoMessage()    {}
func (*L2VpnEvpnEad) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b249f33df3bc066, []int{2}
}

func (m *L2VpnEvpnEad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEad.Unmarshal(m, b)
}
func (m *L2VpnEvpnEad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEad.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEad.Merge(m, src)
}
func (m *L2VpnEvpnEad) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEad.Size(m)
}
func (m *L2VpnEvpnEad) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEad.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEad proto.InternalMessageInfo

func (m *L2VpnEvpnEad) GetEthernetVpnid() uint32 {
	if m != nil {
		return m.EthernetVpnid
	}
	return 0
}

func (m *L2VpnEvpnEad) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *L2VpnEvpnEad) GetEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.EthernetSegmentIdentifier
	}
	return nil
}

func (m *L2VpnEvpnEad) GetEthernetTagXr() uint32 {
	if m != nil {
		return m.EthernetTagXr
	}
	return 0
}

func (m *L2VpnEvpnEad) GetLocalNextHop() string {
	if m != nil {
		return m.LocalNextHop
	}
	return ""
}

func (m *L2VpnEvpnEad) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *L2VpnEvpnEad) GetIsLocalEad() bool {
	if m != nil {
		return m.IsLocalEad
	}
	return false
}

func (m *L2VpnEvpnEad) GetEncap() uint32 {
	if m != nil {
		return m.Encap
	}
	return 0
}

func (m *L2VpnEvpnEad) GetRedundancySingleActive() bool {
	if m != nil {
		return m.RedundancySingleActive
	}
	return false
}

func (m *L2VpnEvpnEad) GetNumPaths() uint32 {
	if m != nil {
		return m.NumPaths
	}
	return 0
}

func (m *L2VpnEvpnEad) GetPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.PathBuffer
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnEvpnEad_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.ethernet_auto_discoveries.ethernet_auto_discovery.l2vpn_evpn_ead_KEYS")
	proto.RegisterType((*L2VpnLabelPathBuffer)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.ethernet_auto_discoveries.ethernet_auto_discovery.l2vpn_label_path_buffer")
	proto.RegisterType((*L2VpnEvpnEad)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.evi_detail.evi_children.ethernet_auto_discoveries.ethernet_auto_discovery.l2vpn_evpn_ead")
}

func init() { proto.RegisterFile("l2vpn_evpn_ead.proto", fileDescriptor_5b249f33df3bc066) }

var fileDescriptor_5b249f33df3bc066 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xdb, 0x6e, 0x13, 0x3d,
	0x10, 0xc7, 0xb5, 0x5f, 0xda, 0x26, 0x9d, 0x1c, 0x3e, 0x64, 0x2a, 0xea, 0xaa, 0x17, 0x2c, 0x11,
	0xa0, 0x5c, 0xad, 0x44, 0x0e, 0x50, 0x0e, 0x42, 0x02, 0xa9, 0x12, 0x15, 0x15, 0x42, 0x49, 0xc5,
	0xe1, 0xca, 0x72, 0xd6, 0x93, 0xc4, 0xd2, 0xc6, 0x6b, 0x79, 0xbd, 0xab, 0xe4, 0x11, 0x78, 0x16,
	0x1e, 0x80, 0x47, 0xe0, 0xb5, 0x90, 0xbd, 0x9b, 0xa5, 0xb9, 0xe8, 0x2d, 0x37, 0xab, 0x99, 0xdf,
	0xfc, 0x67, 0xec, 0xbf, 0xbd, 0x86, 0x93, 0x64, 0x58, 0x68, 0xc5, 0xd0, 0x7f, 0xb8, 0x88, 0xb4,
	0x49, 0x6d, 0x4a, 0x74, 0x2c, 0xb3, 0x38, 0x65, 0x32, 0xcd, 0xd8, 0xc6, 0x94, 0xc5, 0x54, 0xa3,
	0x89, 0x5c, 0x14, 0xa9, 0x54, 0x60, 0xe6, 0xbf, 0x11, 0x16, 0x92, 0x09, 0xb4, 0x5c, 0x26, 0x3e,
	0x8c, 0x57, 0x32, 0x11, 0x06, 0x55, 0x84, 0x76, 0x85, 0x46, 0xa1, 0x65, 0x3c, 0xb7, 0x29, 0x13,
	0x6e, 0x58, 0x81, 0x46, 0x62, 0x76, 0x47, 0x65, 0xdb, 0xff, 0x1d, 0xc0, 0xfd, 0xfd, 0xad, 0xb0,
	0x8f, 0x97, 0xdf, 0x67, 0xe4, 0x14, 0x9a, 0x6e, 0x25, 0x26, 0x05, 0x0d, 0xc2, 0x60, 0x70, 0x3c,
	0x3d, 0x72, 0xe9, 0x95, 0x20, 0xf7, 0xa0, 0x81, 0x85, 0xa4, 0xff, 0x85, 0xc1, 0xa0, 0x3b, 0x75,
	0x21, 0x21, 0x70, 0x80, 0x99, 0x7c, 0x46, 0x1b, 0x5e, 0xe7, 0xe3, 0x8a, 0x0d, 0xe9, 0x41, 0xcd,
	0x86, 0x15, 0x1b, 0xd1, 0xc3, 0x9a, 0x8d, 0x2a, 0x36, 0xa6, 0x47, 0x35, 0x1b, 0x57, 0x6c, 0x42,
	0x9b, 0x35, 0x9b, 0x90, 0x47, 0xd0, 0xa9, 0x1d, 0x58, 0xbe, 0xa4, 0x2d, 0xbf, 0x7c, 0x7b, 0xc7,
	0x6e, 0xf8, 0xb2, 0xff, 0x15, 0x4e, 0x4b, 0x23, 0x09, 0x9f, 0x63, 0xc2, 0x34, 0xb7, 0x2b, 0x36,
	0xcf, 0x17, 0x0b, 0x34, 0xe4, 0x0c, 0x5a, 0x0a, 0x37, 0x96, 0xad, 0x52, 0x5d, 0xb9, 0x69, 0xba,
	0xfc, 0x43, 0xaa, 0xdd, 0xe0, 0x34, 0xb7, 0x3a, 0xb7, 0x65, 0x5b, 0xe5, 0xab, 0x5d, 0xb2, 0x6b,
	0x87, 0xfa, 0xbf, 0x0e, 0xa0, 0xb7, 0x7f, 0x44, 0xe4, 0x09, 0xf4, 0xea, 0xed, 0x14, 0x5a, 0x49,
	0x41, 0x87, 0xbe, 0xaf, 0xbb, 0xa3, 0x5f, 0x1c, 0x74, 0x4e, 0xec, 0x56, 0x23, 0x1d, 0x95, 0x4e,
	0x5c, 0x4c, 0xde, 0xc2, 0x79, 0xdd, 0x9a, 0xe1, 0x72, 0x8d, 0xca, 0x32, 0x29, 0x50, 0x59, 0xb9,
	0x90, 0x68, 0xe8, 0x38, 0x6c, 0x0c, 0xba, 0xd3, 0xb3, 0x9d, 0x64, 0x56, 0x2a, 0xae, 0x6a, 0x01,
	0x79, 0x0a, 0xff, 0xdf, 0x3e, 0x09, 0xb6, 0x31, 0x74, 0xb2, 0xbf, 0xf6, 0x0d, 0x5f, 0x7e, 0x33,
	0xe4, 0x31, 0xf4, 0x92, 0x34, 0xe6, 0x09, 0xab, 0x9d, 0x3f, 0xf7, 0xbb, 0xe8, 0x78, 0xfa, 0xa9,
	0xb2, 0xff, 0x10, 0xda, 0xa5, 0xaa, 0x74, 0xff, 0xc2, 0x4f, 0x02, 0x8f, 0xbc, 0x79, 0x12, 0x42,
	0x47, 0x66, 0xac, 0xd4, 0x20, 0x17, 0xf4, 0x22, 0x0c, 0x06, 0xad, 0x29, 0xc8, 0xec, 0xda, 0xa1,
	0x4b, 0x2e, 0xc8, 0x09, 0x1c, 0xa2, 0x8a, 0xb9, 0xa6, 0x2f, 0x7d, 0x73, 0x99, 0x90, 0x0b, 0xa0,
	0x06, 0x45, 0xae, 0x04, 0x57, 0xf1, 0x96, 0x65, 0x52, 0x2d, 0x13, 0x64, 0x3c, 0xb6, 0xb2, 0x40,
	0xfa, 0xca, 0xcf, 0x78, 0xf0, 0xb7, 0x3e, 0xf3, 0xe5, 0x77, 0xbe, 0x4a, 0xce, 0xe1, 0x58, 0xe5,
	0x6b, 0x7f, 0x7f, 0x19, 0x7d, 0xed, 0x67, 0xb6, 0x54, 0xbe, 0xfe, 0xec, 0x72, 0xf2, 0x33, 0x80,
	0xf6, 0xad, 0x9b, 0xa5, 0x6f, 0xc2, 0xc6, 0xa0, 0x3d, 0xfc, 0x11, 0x44, 0xff, 0xfa, 0xe1, 0x44,
	0x77, 0xfc, 0x6b, 0x53, 0x70, 0xc9, 0x7b, 0x1f, 0xcf, 0x8f, 0xfc, 0xab, 0x1e, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xc2, 0xf7, 0x23, 0x2b, 0xed, 0x03, 0x00, 0x00,
}
