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
// source: mpls_te_bfd_summary.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_bfd_summary

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

type MplsTeBfdSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBfdSummary_KEYS) Reset()         { *m = MplsTeBfdSummary_KEYS{} }
func (m *MplsTeBfdSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdSummary_KEYS) ProtoMessage()    {}
func (*MplsTeBfdSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_40674d286562ef44, []int{0}
}

func (m *MplsTeBfdSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdSummary_KEYS.Unmarshal(m, b)
}
func (m *MplsTeBfdSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdSummary_KEYS.Merge(m, src)
}
func (m *MplsTeBfdSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdSummary_KEYS.Size(m)
}
func (m *MplsTeBfdSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdSummary_KEYS proto.InternalMessageInfo

type MplsTeBfdSummary struct {
	TunnelCountTotal               uint32   `protobuf:"varint,50,opt,name=tunnel_count_total,json=tunnelCountTotal,proto3" json:"tunnel_count_total,omitempty"`
	TunnelCountBfdEnabled          uint32   `protobuf:"varint,51,opt,name=tunnel_count_bfd_enabled,json=tunnelCountBfdEnabled,proto3" json:"tunnel_count_bfd_enabled,omitempty"`
	TunnelCountSessionUp           uint32   `protobuf:"varint,52,opt,name=tunnel_count_session_up,json=tunnelCountSessionUp,proto3" json:"tunnel_count_session_up,omitempty"`
	TunnelCountSbfdEnabled         uint32   `protobuf:"varint,53,opt,name=tunnel_count_sbfd_enabled,json=tunnelCountSbfdEnabled,proto3" json:"tunnel_count_sbfd_enabled,omitempty"`
	TunnelCountSbfdSessionUp       uint32   `protobuf:"varint,54,opt,name=tunnel_count_sbfd_session_up,json=tunnelCountSbfdSessionUp,proto3" json:"tunnel_count_sbfd_session_up,omitempty"`
	HeadLspCountSessionCreated     uint32   `protobuf:"varint,55,opt,name=head_lsp_count_session_created,json=headLspCountSessionCreated,proto3" json:"head_lsp_count_session_created,omitempty"`
	HeadLspCountUp                 uint32   `protobuf:"varint,56,opt,name=head_lsp_count_up,json=headLspCountUp,proto3" json:"head_lsp_count_up,omitempty"`
	TailLspCountSessionCreated     uint32   `protobuf:"varint,57,opt,name=tail_lsp_count_session_created,json=tailLspCountSessionCreated,proto3" json:"tail_lsp_count_session_created,omitempty"`
	TailLspCountSessionUp          uint32   `protobuf:"varint,58,opt,name=tail_lsp_count_session_up,json=tailLspCountSessionUp,proto3" json:"tail_lsp_count_session_up,omitempty"`
	HeadLspCountSbfdSessionCreated uint32   `protobuf:"varint,59,opt,name=head_lsp_count_sbfd_session_created,json=headLspCountSbfdSessionCreated,proto3" json:"head_lsp_count_sbfd_session_created,omitempty"`
	HeadLspCountSbfdUp             uint32   `protobuf:"varint,60,opt,name=head_lsp_count_sbfd_up,json=headLspCountSbfdUp,proto3" json:"head_lsp_count_sbfd_up,omitempty"`
	LinkCountBfdEnabled            uint32   `protobuf:"varint,61,opt,name=link_count_bfd_enabled,json=linkCountBfdEnabled,proto3" json:"link_count_bfd_enabled,omitempty"`
	LinkCountSessionCreated        uint32   `protobuf:"varint,62,opt,name=link_count_session_created,json=linkCountSessionCreated,proto3" json:"link_count_session_created,omitempty"`
	LinkCountSessionUp             uint32   `protobuf:"varint,63,opt,name=link_count_session_up,json=linkCountSessionUp,proto3" json:"link_count_session_up,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *MplsTeBfdSummary) Reset()         { *m = MplsTeBfdSummary{} }
func (m *MplsTeBfdSummary) String() string { return proto.CompactTextString(m) }
func (*MplsTeBfdSummary) ProtoMessage()    {}
func (*MplsTeBfdSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_40674d286562ef44, []int{1}
}

func (m *MplsTeBfdSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBfdSummary.Unmarshal(m, b)
}
func (m *MplsTeBfdSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBfdSummary.Marshal(b, m, deterministic)
}
func (m *MplsTeBfdSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBfdSummary.Merge(m, src)
}
func (m *MplsTeBfdSummary) XXX_Size() int {
	return xxx_messageInfo_MplsTeBfdSummary.Size(m)
}
func (m *MplsTeBfdSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBfdSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBfdSummary proto.InternalMessageInfo

func (m *MplsTeBfdSummary) GetTunnelCountTotal() uint32 {
	if m != nil {
		return m.TunnelCountTotal
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTunnelCountBfdEnabled() uint32 {
	if m != nil {
		return m.TunnelCountBfdEnabled
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTunnelCountSessionUp() uint32 {
	if m != nil {
		return m.TunnelCountSessionUp
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTunnelCountSbfdEnabled() uint32 {
	if m != nil {
		return m.TunnelCountSbfdEnabled
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTunnelCountSbfdSessionUp() uint32 {
	if m != nil {
		return m.TunnelCountSbfdSessionUp
	}
	return 0
}

func (m *MplsTeBfdSummary) GetHeadLspCountSessionCreated() uint32 {
	if m != nil {
		return m.HeadLspCountSessionCreated
	}
	return 0
}

func (m *MplsTeBfdSummary) GetHeadLspCountUp() uint32 {
	if m != nil {
		return m.HeadLspCountUp
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTailLspCountSessionCreated() uint32 {
	if m != nil {
		return m.TailLspCountSessionCreated
	}
	return 0
}

func (m *MplsTeBfdSummary) GetTailLspCountSessionUp() uint32 {
	if m != nil {
		return m.TailLspCountSessionUp
	}
	return 0
}

func (m *MplsTeBfdSummary) GetHeadLspCountSbfdSessionCreated() uint32 {
	if m != nil {
		return m.HeadLspCountSbfdSessionCreated
	}
	return 0
}

func (m *MplsTeBfdSummary) GetHeadLspCountSbfdUp() uint32 {
	if m != nil {
		return m.HeadLspCountSbfdUp
	}
	return 0
}

func (m *MplsTeBfdSummary) GetLinkCountBfdEnabled() uint32 {
	if m != nil {
		return m.LinkCountBfdEnabled
	}
	return 0
}

func (m *MplsTeBfdSummary) GetLinkCountSessionCreated() uint32 {
	if m != nil {
		return m.LinkCountSessionCreated
	}
	return 0
}

func (m *MplsTeBfdSummary) GetLinkCountSessionUp() uint32 {
	if m != nil {
		return m.LinkCountSessionUp
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeBfdSummary_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.summary.mpls_te_bfd_summary_KEYS")
	proto.RegisterType((*MplsTeBfdSummary)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.bfd.summary.mpls_te_bfd_summary")
}

func init() { proto.RegisterFile("mpls_te_bfd_summary.proto", fileDescriptor_40674d286562ef44) }

var fileDescriptor_40674d286562ef44 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0x6f, 0xe2, 0x30,
	0x14, 0xc4, 0xb5, 0x17, 0x0e, 0x4f, 0xda, 0xd5, 0xae, 0x59, 0xc0, 0xa0, 0x15, 0x5a, 0xd1, 0x4b,
	0x2b, 0x55, 0x91, 0x0a, 0xa5, 0x40, 0x69, 0xa9, 0x04, 0xe2, 0x44, 0x4f, 0xa5, 0x39, 0xf4, 0x64,
	0xe5, 0x8f, 0x51, 0xa3, 0x86, 0xd8, 0x8a, 0x6d, 0xa9, 0x7c, 0xbc, 0x7e, 0xb3, 0x2a, 0x0e, 0x29,
	0x8e, 0x09, 0x3d, 0x26, 0xf3, 0x7e, 0x33, 0x63, 0xe7, 0x05, 0xda, 0x5b, 0x1e, 0x0b, 0x22, 0x29,
	0xf1, 0x37, 0x21, 0x11, 0x6a, 0xbb, 0xf5, 0xd2, 0x9d, 0xc3, 0x53, 0x26, 0x19, 0x1a, 0x06, 0x91,
	0x08, 0x18, 0x89, 0x98, 0x20, 0xef, 0x29, 0x29, 0xe6, 0x18, 0xa7, 0xa9, 0x53, 0x3c, 0x08, 0xe9,
	0x25, 0xa1, 0xbf, 0x73, 0xfc, 0x4d, 0xe8, 0xec, 0xe1, 0x5e, 0x07, 0x70, 0x85, 0x27, 0x59, 0x2d,
	0x5f, 0xd6, 0xbd, 0x8f, 0x1a, 0xd4, 0x2b, 0x44, 0x74, 0x09, 0x48, 0xaa, 0x24, 0xa1, 0x31, 0x09,
	0x98, 0x4a, 0x24, 0x91, 0x4c, 0x7a, 0x31, 0xee, 0xff, 0xff, 0x71, 0xfe, 0xf3, 0xe9, 0x77, 0xae,
	0x2c, 0x32, 0xe1, 0x39, 0x7b, 0x8f, 0x46, 0x80, 0x4b, 0xd3, 0x99, 0x13, 0x4d, 0x3c, 0x3f, 0xa6,
	0x21, 0x1e, 0x68, 0xa6, 0x61, 0x30, 0xf3, 0x4d, 0xb8, 0xcc, 0x45, 0x34, 0x84, 0x56, 0x09, 0x14,
	0x54, 0x88, 0x88, 0x25, 0x44, 0x71, 0x7c, 0xad, 0xb9, 0xbf, 0x06, 0xb7, 0xce, 0x45, 0x97, 0xa3,
	0x09, 0xb4, 0xcb, 0x98, 0x19, 0x38, 0xd4, 0x60, 0xd3, 0x04, 0xfd, 0x43, 0xe2, 0x0c, 0xfe, 0x1d,
	0xa3, 0x46, 0xec, 0x8d, 0xa6, 0xb1, 0x45, 0x1f, 0xa2, 0xe7, 0xd0, 0x7d, 0xa5, 0x5e, 0x48, 0x62,
	0xc1, 0xad, 0xce, 0x41, 0x4a, 0x3d, 0x49, 0x43, 0x3c, 0xd2, 0x0e, 0x9d, 0x6c, 0xea, 0x51, 0x70,
	0xb3, 0xf9, 0x22, 0x9f, 0x40, 0x17, 0xf0, 0xc7, 0xf2, 0x50, 0x1c, 0x8f, 0x35, 0xf6, 0xcb, 0xc4,
	0xf2, 0x38, 0xe9, 0x45, 0xf1, 0x37, 0x71, 0x93, 0x3c, 0x2e, 0x9b, 0x3a, 0x11, 0x37, 0x86, 0xf6,
	0x09, 0x0f, 0xc5, 0xf1, 0xed, 0xfe, 0xf3, 0x1c, 0xe3, 0x2e, 0x47, 0x2b, 0x38, 0xb3, 0x0f, 0x6b,
	0x5e, 0x57, 0x51, 0x61, 0xaa, 0x3d, 0xba, 0xa5, 0x13, 0x1f, 0x2e, 0xad, 0xa8, 0xd1, 0x87, 0x66,
	0x95, 0x99, 0xe2, 0xf8, 0x4e, 0xf3, 0xc8, 0xe6, 0x5d, 0x8e, 0x06, 0xd0, 0x8c, 0xa3, 0xe4, 0xad,
	0x62, 0xad, 0xee, 0x35, 0x53, 0xcf, 0x54, 0x7b, 0xa9, 0xa6, 0xd0, 0x31, 0x20, 0xbb, 0xec, 0x4c,
	0x83, 0xad, 0x2f, 0xd0, 0x6a, 0x79, 0x05, 0x8d, 0x0a, 0x58, 0x71, 0xfc, 0x90, 0x97, 0xb4, 0x39,
	0x97, 0xfb, 0x35, 0xfd, 0x77, 0x0e, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62, 0x5c, 0x67, 0xb1,
	0xba, 0x03, 0x00, 0x00,
}
