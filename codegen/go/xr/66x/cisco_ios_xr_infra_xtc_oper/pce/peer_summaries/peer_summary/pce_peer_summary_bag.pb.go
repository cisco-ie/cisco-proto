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
// source: pce_peer_summary_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_peer_summaries_peer_summary

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

type PcePeerSummaryBag_KEYS struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePeerSummaryBag_KEYS) Reset()         { *m = PcePeerSummaryBag_KEYS{} }
func (m *PcePeerSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcePeerSummaryBag_KEYS) ProtoMessage()    {}
func (*PcePeerSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8325859fc081eb3c, []int{0}
}

func (m *PcePeerSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *PcePeerSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcePeerSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerSummaryBag_KEYS.Merge(m, src)
}
func (m *PcePeerSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcePeerSummaryBag_KEYS.Size(m)
}
func (m *PcePeerSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerSummaryBag_KEYS proto.InternalMessageInfo

func (m *PcePeerSummaryBag_KEYS) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

type PcePeerProtoSummaryBag struct {
	PeerCountUp          uint32   `protobuf:"varint,1,opt,name=peer_count_up,json=peerCountUp,proto3" json:"peer_count_up,omitempty"`
	PeerCountDown        uint32   `protobuf:"varint,2,opt,name=peer_count_down,json=peerCountDown,proto3" json:"peer_count_down,omitempty"`
	PeerCountAll         uint32   `protobuf:"varint,3,opt,name=peer_count_all,json=peerCountAll,proto3" json:"peer_count_all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePeerProtoSummaryBag) Reset()         { *m = PcePeerProtoSummaryBag{} }
func (m *PcePeerProtoSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PcePeerProtoSummaryBag) ProtoMessage()    {}
func (*PcePeerProtoSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8325859fc081eb3c, []int{1}
}

func (m *PcePeerProtoSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerProtoSummaryBag.Unmarshal(m, b)
}
func (m *PcePeerProtoSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerProtoSummaryBag.Marshal(b, m, deterministic)
}
func (m *PcePeerProtoSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerProtoSummaryBag.Merge(m, src)
}
func (m *PcePeerProtoSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PcePeerProtoSummaryBag.Size(m)
}
func (m *PcePeerProtoSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerProtoSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerProtoSummaryBag proto.InternalMessageInfo

func (m *PcePeerProtoSummaryBag) GetPeerCountUp() uint32 {
	if m != nil {
		return m.PeerCountUp
	}
	return 0
}

func (m *PcePeerProtoSummaryBag) GetPeerCountDown() uint32 {
	if m != nil {
		return m.PeerCountDown
	}
	return 0
}

func (m *PcePeerProtoSummaryBag) GetPeerCountAll() uint32 {
	if m != nil {
		return m.PeerCountAll
	}
	return 0
}

type PcePeerSummaryBag struct {
	PcepPeers            *PcePeerProtoSummaryBag `protobuf:"bytes,50,opt,name=pcep_peers,json=pcepPeers,proto3" json:"pcep_peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PcePeerSummaryBag) Reset()         { *m = PcePeerSummaryBag{} }
func (m *PcePeerSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PcePeerSummaryBag) ProtoMessage()    {}
func (*PcePeerSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8325859fc081eb3c, []int{2}
}

func (m *PcePeerSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerSummaryBag.Unmarshal(m, b)
}
func (m *PcePeerSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerSummaryBag.Marshal(b, m, deterministic)
}
func (m *PcePeerSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerSummaryBag.Merge(m, src)
}
func (m *PcePeerSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PcePeerSummaryBag.Size(m)
}
func (m *PcePeerSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerSummaryBag proto.InternalMessageInfo

func (m *PcePeerSummaryBag) GetPcepPeers() *PcePeerProtoSummaryBag {
	if m != nil {
		return m.PcepPeers
	}
	return nil
}

func init() {
	proto.RegisterType((*PcePeerSummaryBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_summaries.peer_summary.pce_peer_summary_bag_KEYS")
	proto.RegisterType((*PcePeerProtoSummaryBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_summaries.peer_summary.pce_peer_proto_summary_bag")
	proto.RegisterType((*PcePeerSummaryBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_summaries.peer_summary.pce_peer_summary_bag")
}

func init() { proto.RegisterFile("pce_peer_summary_bag.proto", fileDescriptor_8325859fc081eb3c) }

var fileDescriptor_8325859fc081eb3c = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x48, 0x4e, 0x8d,
	0x2f, 0x48, 0x4d, 0x2d, 0x8a, 0x2f, 0x2e, 0xcd, 0xcd, 0x4d, 0x2c, 0xaa, 0x8c, 0x4f, 0x4a, 0x4c,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4e, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc,
	0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0xcc, 0x4b, 0x2b, 0x4a, 0x8c, 0xaf, 0x28, 0x49, 0x8e, 0xcf,
	0x2f, 0x48, 0x2d, 0xd2, 0x2b, 0x48, 0x4e, 0xd5, 0x43, 0xd2, 0x97, 0x99, 0x5a, 0x8c, 0xcc, 0xad,
	0x54, 0xd2, 0xe6, 0x92, 0xc4, 0x66, 0x74, 0xbc, 0xb7, 0x6b, 0x64, 0xb0, 0x10, 0x1f, 0x17, 0x53,
	0x62, 0x9a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x62, 0x9a, 0x52, 0x1f, 0x23, 0x92,
	0x43, 0xc0, 0x96, 0x23, 0xeb, 0x11, 0x52, 0xe2, 0xe2, 0x05, 0xcb, 0x24, 0xe7, 0x97, 0xe6, 0x95,
	0xc4, 0x97, 0x16, 0x80, 0x75, 0xf2, 0x06, 0x71, 0x83, 0x04, 0x9d, 0x41, 0x62, 0xa1, 0x05, 0x42,
	0x6a, 0x5c, 0xfc, 0x48, 0x6a, 0x52, 0xf2, 0xcb, 0xf3, 0x24, 0x98, 0xc0, 0xaa, 0x78, 0xe1, 0xaa,
	0x5c, 0xf2, 0xcb, 0xf3, 0x84, 0x54, 0xb8, 0xf8, 0x90, 0xd4, 0x25, 0xe6, 0xe4, 0x48, 0x30, 0x83,
	0x95, 0xf1, 0xc0, 0x95, 0x39, 0xe6, 0xe4, 0x80, 0x1c, 0x24, 0x82, 0xcd, 0xf9, 0x42, 0x65, 0x5c,
	0x5c, 0x05, 0xc9, 0xa9, 0x05, 0x60, 0x89, 0x62, 0x09, 0x23, 0x05, 0x46, 0x0d, 0x6e, 0xa3, 0x70,
	0x3d, 0x0a, 0x02, 0x4a, 0x0f, 0xb7, 0xbf, 0x83, 0x38, 0x41, 0x56, 0x05, 0x80, 0x6c, 0x4a, 0x62,
	0x03, 0xcb, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x86, 0x4e, 0x60, 0x41, 0xb0, 0x01, 0x00,
	0x00,
}
