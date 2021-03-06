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
// source: mpls_te_tp_detail_link_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_tp_tp_links_tp_detail_links_tp_detail_link

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

type MplsTeTpDetailLinkInfo_KEYS struct {
	TpLinkId             uint32   `protobuf:"varint,1,opt,name=tp_link_id,json=tpLinkId,proto3" json:"tp_link_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpDetailLinkInfo_KEYS) Reset()         { *m = MplsTeTpDetailLinkInfo_KEYS{} }
func (m *MplsTeTpDetailLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpDetailLinkInfo_KEYS) ProtoMessage()    {}
func (*MplsTeTpDetailLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc0768c534f05b26, []int{0}
}

func (m *MplsTeTpDetailLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTpDetailLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeTpDetailLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS.Merge(m, src)
}
func (m *MplsTeTpDetailLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS.Size(m)
}
func (m *MplsTeTpDetailLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpDetailLinkInfo_KEYS proto.InternalMessageInfo

func (m *MplsTeTpDetailLinkInfo_KEYS) GetTpLinkId() uint32 {
	if m != nil {
		return m.TpLinkId
	}
	return 0
}

type MplsTeTpDetailLinkInfo struct {
	LinkId               uint32   `protobuf:"varint,50,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	Interface            string   `protobuf:"bytes,51,opt,name=interface,proto3" json:"interface,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,52,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	LinkState            string   `protobuf:"bytes,53,opt,name=link_state,json=linkState,proto3" json:"link_state,omitempty"`
	LsPs                 uint32   `protobuf:"varint,54,opt,name=ls_ps,json=lsPs,proto3" json:"ls_ps,omitempty"`
	ReservedBandwidth    uint64   `protobuf:"varint,55,opt,name=reserved_bandwidth,json=reservedBandwidth,proto3" json:"reserved_bandwidth,omitempty"`
	AvailableBandwidth   uint64   `protobuf:"varint,56,opt,name=available_bandwidth,json=availableBandwidth,proto3" json:"available_bandwidth,omitempty"`
	UnsupportedLinecard  bool     `protobuf:"varint,57,opt,name=unsupported_linecard,json=unsupportedLinecard,proto3" json:"unsupported_linecard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTpDetailLinkInfo) Reset()         { *m = MplsTeTpDetailLinkInfo{} }
func (m *MplsTeTpDetailLinkInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeTpDetailLinkInfo) ProtoMessage()    {}
func (*MplsTeTpDetailLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc0768c534f05b26, []int{1}
}

func (m *MplsTeTpDetailLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo.Unmarshal(m, b)
}
func (m *MplsTeTpDetailLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeTpDetailLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTpDetailLinkInfo.Merge(m, src)
}
func (m *MplsTeTpDetailLinkInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeTpDetailLinkInfo.Size(m)
}
func (m *MplsTeTpDetailLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTpDetailLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTpDetailLinkInfo proto.InternalMessageInfo

func (m *MplsTeTpDetailLinkInfo) GetLinkId() uint32 {
	if m != nil {
		return m.LinkId
	}
	return 0
}

func (m *MplsTeTpDetailLinkInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *MplsTeTpDetailLinkInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *MplsTeTpDetailLinkInfo) GetLinkState() string {
	if m != nil {
		return m.LinkState
	}
	return ""
}

func (m *MplsTeTpDetailLinkInfo) GetLsPs() uint32 {
	if m != nil {
		return m.LsPs
	}
	return 0
}

func (m *MplsTeTpDetailLinkInfo) GetReservedBandwidth() uint64 {
	if m != nil {
		return m.ReservedBandwidth
	}
	return 0
}

func (m *MplsTeTpDetailLinkInfo) GetAvailableBandwidth() uint64 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

func (m *MplsTeTpDetailLinkInfo) GetUnsupportedLinecard() bool {
	if m != nil {
		return m.UnsupportedLinecard
	}
	return false
}

func init() {
	proto.RegisterType((*MplsTeTpDetailLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_links.tp_detail_links.tp_detail_link.mpls_te_tp_detail_link_info_KEYS")
	proto.RegisterType((*MplsTeTpDetailLinkInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_links.tp_detail_links.tp_detail_link.mpls_te_tp_detail_link_info")
}

func init() { proto.RegisterFile("mpls_te_tp_detail_link_info.proto", fileDescriptor_dc0768c534f05b26) }

var fileDescriptor_dc0768c534f05b26 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x4b, 0x2b, 0x31,
	0x14, 0xc5, 0x99, 0x47, 0x5f, 0x5f, 0x7b, 0xe1, 0x3d, 0x9e, 0xa9, 0xe0, 0x80, 0x15, 0xc6, 0xae,
	0x66, 0xe3, 0x88, 0xd6, 0xbf, 0x3b, 0x15, 0x04, 0x8b, 0x5d, 0xc8, 0x74, 0xe5, 0xea, 0x92, 0x4e,
	0x6e, 0x69, 0x68, 0x9c, 0x84, 0x24, 0xad, 0xfd, 0x82, 0x7e, 0x2f, 0x99, 0xb4, 0x63, 0xab, 0x8b,
	0x2e, 0xef, 0x39, 0xe7, 0x77, 0x0e, 0x24, 0x70, 0xfc, 0x66, 0x94, 0x43, 0x4f, 0xe8, 0x0d, 0x0a,
	0xf2, 0x5c, 0x2a, 0x54, 0xb2, 0x9c, 0xa1, 0x2c, 0x27, 0x3a, 0x33, 0x56, 0x7b, 0xcd, 0x06, 0x85,
	0x74, 0x85, 0x46, 0xa9, 0x1d, 0x2e, 0x2d, 0xd6, 0x79, 0x6d, 0xc8, 0x66, 0xab, 0xc3, 0x64, 0xde,
	0x04, 0xca, 0x65, 0xdf, 0x5b, 0x7e, 0xde, 0xbd, 0x3b, 0x48, 0x76, 0xec, 0xe1, 0xf3, 0xe3, 0xeb,
	0x88, 0x75, 0x01, 0xd6, 0x75, 0x28, 0x45, 0x1c, 0x25, 0x51, 0xfa, 0x37, 0x6f, 0x79, 0x33, 0x94,
	0xe5, 0x6c, 0x20, 0x7a, 0x1f, 0xbf, 0xe0, 0x70, 0x47, 0x05, 0x3b, 0x80, 0x3f, 0x35, 0x7a, 0x1e,
	0xd0, 0xa6, 0x0a, 0x20, 0xeb, 0x42, 0x5b, 0x96, 0x9e, 0xec, 0x84, 0x17, 0x14, 0xf7, 0x93, 0x28,
	0x6d, 0xe7, 0x1b, 0x81, 0xa5, 0xf0, 0xbf, 0xa4, 0xa5, 0xc7, 0xa9, 0x36, 0xc8, 0x85, 0xb0, 0xe4,
	0x5c, 0x7c, 0x11, 0x42, 0xff, 0x2a, 0xfd, 0x49, 0x9b, 0xfb, 0x95, 0xca, 0x8e, 0x00, 0xc2, 0x80,
	0xf3, 0xdc, 0x53, 0x7c, 0xb9, 0x2a, 0xaa, 0x94, 0x51, 0x25, 0xb0, 0x0e, 0xfc, 0x56, 0x0e, 0x8d,
	0x8b, 0xaf, 0xc2, 0x7a, 0x43, 0xb9, 0x17, 0xc7, 0x4e, 0x80, 0x59, 0x72, 0x64, 0x17, 0x24, 0x70,
	0xcc, 0x4b, 0xf1, 0x2e, 0x85, 0x9f, 0xc6, 0xd7, 0x49, 0x94, 0x36, 0xf2, 0xbd, 0xda, 0x79, 0xa8,
	0x0d, 0x76, 0x0a, 0x1d, 0xbe, 0xe0, 0x52, 0xf1, 0xb1, 0xa2, 0xad, 0xfc, 0x4d, 0xc8, 0xb3, 0x2f,
	0x6b, 0x03, 0x9c, 0xc1, 0xfe, 0xbc, 0x74, 0x73, 0x63, 0xb4, 0xf5, 0x24, 0xaa, 0xd7, 0xa0, 0x82,
	0x5b, 0x11, 0xdf, 0x26, 0x51, 0xda, 0xca, 0x3b, 0x5b, 0xde, 0x70, 0x6d, 0x8d, 0x9b, 0xe1, 0x6f,
	0xfb, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x4e, 0x6c, 0x44, 0x00, 0x02, 0x00, 0x00,
}
