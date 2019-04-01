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
// source: pce_lsp_summary_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_lsp_data_lsp_summary

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

type PceLspSummaryBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceLspSummaryBag_KEYS) Reset()         { *m = PceLspSummaryBag_KEYS{} }
func (m *PceLspSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceLspSummaryBag_KEYS) ProtoMessage()    {}
func (*PceLspSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b906580c59cc32, []int{0}
}

func (m *PceLspSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceLspSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *PceLspSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceLspSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceLspSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceLspSummaryBag_KEYS.Merge(m, src)
}
func (m *PceLspSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceLspSummaryBag_KEYS.Size(m)
}
func (m *PceLspSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceLspSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceLspSummaryBag_KEYS proto.InternalMessageInfo

type PceLspCountsBag struct {
	AllLsPs              uint32   `protobuf:"varint,1,opt,name=all_ls_ps,json=allLsPs,proto3" json:"all_ls_ps,omitempty"`
	UpLsPs               uint32   `protobuf:"varint,2,opt,name=up_ls_ps,json=upLsPs,proto3" json:"up_ls_ps,omitempty"`
	AdminUpLsPs          uint32   `protobuf:"varint,3,opt,name=admin_up_ls_ps,json=adminUpLsPs,proto3" json:"admin_up_ls_ps,omitempty"`
	SrLsPs               uint32   `protobuf:"varint,4,opt,name=sr_ls_ps,json=srLsPs,proto3" json:"sr_ls_ps,omitempty"`
	RsvpLsPs             uint32   `protobuf:"varint,5,opt,name=rsvp_ls_ps,json=rsvpLsPs,proto3" json:"rsvp_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceLspCountsBag) Reset()         { *m = PceLspCountsBag{} }
func (m *PceLspCountsBag) String() string { return proto.CompactTextString(m) }
func (*PceLspCountsBag) ProtoMessage()    {}
func (*PceLspCountsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b906580c59cc32, []int{1}
}

func (m *PceLspCountsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceLspCountsBag.Unmarshal(m, b)
}
func (m *PceLspCountsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceLspCountsBag.Marshal(b, m, deterministic)
}
func (m *PceLspCountsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceLspCountsBag.Merge(m, src)
}
func (m *PceLspCountsBag) XXX_Size() int {
	return xxx_messageInfo_PceLspCountsBag.Size(m)
}
func (m *PceLspCountsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceLspCountsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceLspCountsBag proto.InternalMessageInfo

func (m *PceLspCountsBag) GetAllLsPs() uint32 {
	if m != nil {
		return m.AllLsPs
	}
	return 0
}

func (m *PceLspCountsBag) GetUpLsPs() uint32 {
	if m != nil {
		return m.UpLsPs
	}
	return 0
}

func (m *PceLspCountsBag) GetAdminUpLsPs() uint32 {
	if m != nil {
		return m.AdminUpLsPs
	}
	return 0
}

func (m *PceLspCountsBag) GetSrLsPs() uint32 {
	if m != nil {
		return m.SrLsPs
	}
	return 0
}

func (m *PceLspCountsBag) GetRsvpLsPs() uint32 {
	if m != nil {
		return m.RsvpLsPs
	}
	return 0
}

type PceIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIpAddrType) Reset()         { *m = PceIpAddrType{} }
func (m *PceIpAddrType) String() string { return proto.CompactTextString(m) }
func (*PceIpAddrType) ProtoMessage()    {}
func (*PceIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b906580c59cc32, []int{2}
}

func (m *PceIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIpAddrType.Unmarshal(m, b)
}
func (m *PceIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIpAddrType.Marshal(b, m, deterministic)
}
func (m *PceIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIpAddrType.Merge(m, src)
}
func (m *PceIpAddrType) XXX_Size() int {
	return xxx_messageInfo_PceIpAddrType.Size(m)
}
func (m *PceIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIpAddrType proto.InternalMessageInfo

func (m *PceIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PceIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PceIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type PceLspPerPeerSummaryBag struct {
	LspSummary           *PceLspCountsBag `protobuf:"bytes,1,opt,name=lsp_summary,json=lspSummary,proto3" json:"lsp_summary,omitempty"`
	PeerAddress          *PceIpAddrType   `protobuf:"bytes,2,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PceLspPerPeerSummaryBag) Reset()         { *m = PceLspPerPeerSummaryBag{} }
func (m *PceLspPerPeerSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PceLspPerPeerSummaryBag) ProtoMessage()    {}
func (*PceLspPerPeerSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b906580c59cc32, []int{3}
}

func (m *PceLspPerPeerSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceLspPerPeerSummaryBag.Unmarshal(m, b)
}
func (m *PceLspPerPeerSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceLspPerPeerSummaryBag.Marshal(b, m, deterministic)
}
func (m *PceLspPerPeerSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceLspPerPeerSummaryBag.Merge(m, src)
}
func (m *PceLspPerPeerSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PceLspPerPeerSummaryBag.Size(m)
}
func (m *PceLspPerPeerSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceLspPerPeerSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceLspPerPeerSummaryBag proto.InternalMessageInfo

func (m *PceLspPerPeerSummaryBag) GetLspSummary() *PceLspCountsBag {
	if m != nil {
		return m.LspSummary
	}
	return nil
}

func (m *PceLspPerPeerSummaryBag) GetPeerAddress() *PceIpAddrType {
	if m != nil {
		return m.PeerAddress
	}
	return nil
}

type PceLspSummaryBag struct {
	PeerLsPsInfo         []*PceLspPerPeerSummaryBag `protobuf:"bytes,50,rep,name=peer_ls_ps_info,json=peerLsPsInfo,proto3" json:"peer_ls_ps_info,omitempty"`
	AllLsPs              *PceLspCountsBag           `protobuf:"bytes,51,opt,name=all_ls_ps,json=allLsPs,proto3" json:"all_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PceLspSummaryBag) Reset()         { *m = PceLspSummaryBag{} }
func (m *PceLspSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PceLspSummaryBag) ProtoMessage()    {}
func (*PceLspSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b906580c59cc32, []int{4}
}

func (m *PceLspSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceLspSummaryBag.Unmarshal(m, b)
}
func (m *PceLspSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceLspSummaryBag.Marshal(b, m, deterministic)
}
func (m *PceLspSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceLspSummaryBag.Merge(m, src)
}
func (m *PceLspSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PceLspSummaryBag.Size(m)
}
func (m *PceLspSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceLspSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceLspSummaryBag proto.InternalMessageInfo

func (m *PceLspSummaryBag) GetPeerLsPsInfo() []*PceLspPerPeerSummaryBag {
	if m != nil {
		return m.PeerLsPsInfo
	}
	return nil
}

func (m *PceLspSummaryBag) GetAllLsPs() *PceLspCountsBag {
	if m != nil {
		return m.AllLsPs
	}
	return nil
}

func init() {
	proto.RegisterType((*PceLspSummaryBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce_lsp_data.lsp_summary.pce_lsp_summary_bag_KEYS")
	proto.RegisterType((*PceLspCountsBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_lsp_data.lsp_summary.pce_lsp_counts_bag")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce_lsp_data.lsp_summary.pce_ip_addr_type")
	proto.RegisterType((*PceLspPerPeerSummaryBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_lsp_data.lsp_summary.pce_lsp_per_peer_summary_bag")
	proto.RegisterType((*PceLspSummaryBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_lsp_data.lsp_summary.pce_lsp_summary_bag")
}

func init() { proto.RegisterFile("pce_lsp_summary_bag.proto", fileDescriptor_69b906580c59cc32) }

var fileDescriptor_69b906580c59cc32 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0x51, 0xed, 0xfa, 0xcf, 0xa8, 0xff, 0xd8, 0x1e, 0xaa, 0x1a, 0x1f, 0x8c, 0x7a, 0xf1,
	0x49, 0x07, 0xd9, 0xf8, 0xde, 0x43, 0x4b, 0x4b, 0x4b, 0x08, 0x32, 0x39, 0xe4, 0x34, 0xac, 0xa5,
	0x55, 0x58, 0x90, 0xb4, 0xcb, 0xae, 0x64, 0xec, 0x4f, 0x94, 0x2f, 0x18, 0x72, 0x0e, 0x3b, 0xb2,
	0x82, 0x20, 0x3e, 0x39, 0xb9, 0x89, 0xf9, 0xf1, 0x76, 0xde, 0xbc, 0x19, 0xc1, 0x77, 0x9d, 0x0a,
	0x2c, 0xac, 0x46, 0xdb, 0x94, 0x25, 0x37, 0x47, 0xdc, 0xf1, 0xbb, 0x48, 0x1b, 0x55, 0x2b, 0xb6,
	0x4e, 0xa5, 0x4d, 0x15, 0x4a, 0x65, 0xf1, 0x60, 0x50, 0x56, 0xb9, 0xe1, 0x78, 0xa8, 0x53, 0x54,
	0x5a, 0x98, 0xa8, 0x93, 0x65, 0xbc, 0xe6, 0x51, 0x4f, 0x1f, 0xce, 0x20, 0x38, 0xf3, 0x24, 0xfe,
	0xfb, 0x75, 0xbb, 0x0d, 0xef, 0x3d, 0x60, 0x1d, 0x4c, 0x55, 0x53, 0xd5, 0xd6, 0x31, 0x36, 0x83,
	0x29, 0x2f, 0x0a, 0x2c, 0x2c, 0x6a, 0x1b, 0x78, 0x0b, 0x6f, 0xf9, 0x31, 0x19, 0xf3, 0xa2, 0xf8,
	0x6f, 0xaf, 0x2d, 0x0b, 0x60, 0xd2, 0xe8, 0x13, 0x7a, 0x47, 0x68, 0xd4, 0x68, 0x22, 0x3f, 0xe0,
	0x13, 0xcf, 0x4a, 0x59, 0xe1, 0x33, 0x1f, 0x10, 0xf7, 0xa9, 0x7a, 0xa3, 0x3b, 0xb9, 0x35, 0x27,
	0x3c, 0x6c, 0xe5, 0xd6, 0x10, 0x99, 0x03, 0x18, 0xbb, 0xef, 0xa4, 0xef, 0x89, 0x4d, 0x5c, 0xc5,
	0xd1, 0x70, 0x0b, 0x5f, 0x9c, 0x51, 0xa9, 0x91, 0x67, 0x99, 0xc1, 0xfa, 0xa8, 0x05, 0xfb, 0x06,
	0x63, 0x9e, 0x63, 0xc5, 0x4b, 0x41, 0x26, 0xa7, 0xc9, 0x88, 0xe7, 0x57, 0xbc, 0x14, 0x8c, 0xc1,
	0x50, 0xea, 0xfd, 0x9a, 0xfc, 0x4d, 0x13, 0xfa, 0x3e, 0xd5, 0x36, 0xe4, 0xa9, 0xad, 0x6d, 0xc2,
	0x07, 0x0f, 0xe6, 0xdd, 0xf8, 0x5a, 0x18, 0xd4, 0x42, 0x98, 0x7e, 0x48, 0x4c, 0x82, 0xdf, 0xcb,
	0x8d, 0xba, 0xf8, 0xf1, 0x9f, 0xe8, 0x92, 0x3d, 0x44, 0x2f, 0x73, 0x4e, 0xa0, 0xb0, 0x7a, 0xdb,
	0x62, 0x26, 0xe1, 0x03, 0xb5, 0x77, 0xe3, 0x09, 0xdb, 0x66, 0xeb, 0xc7, 0xbf, 0x2f, 0xef, 0xd5,
	0x8f, 0x2a, 0xf1, 0xdd, 0xdb, 0x3f, 0xdb, 0xa7, 0xc3, 0x47, 0x0f, 0xbe, 0x9e, 0x39, 0x09, 0x76,
	0x84, 0xcf, 0x64, 0x81, 0x36, 0xe0, 0x7a, 0xa9, 0x20, 0x5e, 0x0c, 0x96, 0x7e, 0x9c, 0xbc, 0x6e,
	0xe2, 0x73, 0xd1, 0x26, 0x34, 0xad, 0x5b, 0xed, 0xdf, 0x2a, 0x57, 0x2c, 0xeb, 0x5f, 0xdc, 0xea,
	0x8d, 0x63, 0xee, 0x6e, 0x77, 0x37, 0xa2, 0xff, 0x68, 0xf5, 0x14, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0xe2, 0x8a, 0x91, 0x64, 0x03, 0x00, 0x00,
}
