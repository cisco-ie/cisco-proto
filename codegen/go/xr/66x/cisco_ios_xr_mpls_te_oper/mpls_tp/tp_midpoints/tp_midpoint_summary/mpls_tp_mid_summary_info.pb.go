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
// source: mpls_tp_mid_summary_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_tp_tp_midpoints_tp_midpoint_summary

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

type MplsTpMidSummaryInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTpMidSummaryInfo_KEYS) Reset()         { *m = MplsTpMidSummaryInfo_KEYS{} }
func (m *MplsTpMidSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTpMidSummaryInfo_KEYS) ProtoMessage()    {}
func (*MplsTpMidSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_40a4f428d99bf413, []int{0}
}

func (m *MplsTpMidSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpMidSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTpMidSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpMidSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTpMidSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpMidSummaryInfo_KEYS.Merge(m, src)
}
func (m *MplsTpMidSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTpMidSummaryInfo_KEYS.Size(m)
}
func (m *MplsTpMidSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpMidSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpMidSummaryInfo_KEYS proto.InternalMessageInfo

type MplsTpMidSummaryInfo struct {
	MidPoints            uint32   `protobuf:"varint,50,opt,name=mid_points,json=midPoints,proto3" json:"mid_points,omitempty"`
	UpForwardLsPs        uint32   `protobuf:"varint,51,opt,name=up_forward_ls_ps,json=upForwardLsPs,proto3" json:"up_forward_ls_ps,omitempty"`
	DownForwardLsPs      uint32   `protobuf:"varint,52,opt,name=down_forward_ls_ps,json=downForwardLsPs,proto3" json:"down_forward_ls_ps,omitempty"`
	UpReverseLsPs        uint32   `protobuf:"varint,53,opt,name=up_reverse_ls_ps,json=upReverseLsPs,proto3" json:"up_reverse_ls_ps,omitempty"`
	DownReverseLsPs      uint32   `protobuf:"varint,54,opt,name=down_reverse_ls_ps,json=downReverseLsPs,proto3" json:"down_reverse_ls_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTpMidSummaryInfo) Reset()         { *m = MplsTpMidSummaryInfo{} }
func (m *MplsTpMidSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTpMidSummaryInfo) ProtoMessage()    {}
func (*MplsTpMidSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_40a4f428d99bf413, []int{1}
}

func (m *MplsTpMidSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTpMidSummaryInfo.Unmarshal(m, b)
}
func (m *MplsTpMidSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTpMidSummaryInfo.Marshal(b, m, deterministic)
}
func (m *MplsTpMidSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTpMidSummaryInfo.Merge(m, src)
}
func (m *MplsTpMidSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTpMidSummaryInfo.Size(m)
}
func (m *MplsTpMidSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTpMidSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTpMidSummaryInfo proto.InternalMessageInfo

func (m *MplsTpMidSummaryInfo) GetMidPoints() uint32 {
	if m != nil {
		return m.MidPoints
	}
	return 0
}

func (m *MplsTpMidSummaryInfo) GetUpForwardLsPs() uint32 {
	if m != nil {
		return m.UpForwardLsPs
	}
	return 0
}

func (m *MplsTpMidSummaryInfo) GetDownForwardLsPs() uint32 {
	if m != nil {
		return m.DownForwardLsPs
	}
	return 0
}

func (m *MplsTpMidSummaryInfo) GetUpReverseLsPs() uint32 {
	if m != nil {
		return m.UpReverseLsPs
	}
	return 0
}

func (m *MplsTpMidSummaryInfo) GetDownReverseLsPs() uint32 {
	if m != nil {
		return m.DownReverseLsPs
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTpMidSummaryInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_midpoint_summary.mpls_tp_mid_summary_info_KEYS")
	proto.RegisterType((*MplsTpMidSummaryInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_tp.tp_midpoints.tp_midpoint_summary.mpls_tp_mid_summary_info")
}

func init() { proto.RegisterFile("mpls_tp_mid_summary_info.proto", fileDescriptor_40a4f428d99bf413) }

var fileDescriptor_40a4f428d99bf413 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xd9, 0x8b, 0x60, 0x60, 0x51, 0x72, 0xea, 0x65, 0x55, 0xf6, 0xa2, 0x20, 0xf4, 0xe0,
	0xaa, 0x0f, 0x20, 0xe8, 0x45, 0x0f, 0xa5, 0x9e, 0x3c, 0x0d, 0xb5, 0x49, 0x21, 0xd0, 0x74, 0x86,
	0x99, 0xd6, 0xea, 0x03, 0xfb, 0x1e, 0xd2, 0xa4, 0xc5, 0x76, 0xa1, 0xc7, 0x7c, 0xf9, 0xfe, 0x2f,
	0x10, 0x75, 0xe1, 0xa9, 0x16, 0x68, 0x09, 0xbc, 0x33, 0x20, 0x9d, 0xf7, 0x05, 0xff, 0x80, 0x6b,
	0x2a, 0x4c, 0x89, 0xb1, 0x45, 0xfd, 0x54, 0x3a, 0x29, 0x11, 0x1c, 0x0a, 0x7c, 0x33, 0x44, 0xd9,
	0x02, 0x92, 0xe5, 0x74, 0x5c, 0xa6, 0x71, 0x4c, 0xe8, 0x9a, 0x56, 0xe6, 0x87, 0x29, 0xb7, 0xbf,
	0x54, 0xbb, 0xb5, 0x57, 0xe0, 0xf5, 0xf9, 0xe3, 0x7d, 0xff, 0xbb, 0x51, 0xc9, 0x9a, 0xa1, 0x77,
	0x4a, 0x0d, 0x2c, 0xf6, 0x93, 0xbb, 0xab, 0xcd, 0xcd, 0x36, 0x3f, 0xf5, 0xce, 0x64, 0x01, 0xe8,
	0x6b, 0x75, 0xde, 0x11, 0x54, 0xc8, 0x7d, 0xc1, 0x06, 0x6a, 0x01, 0x92, 0xe4, 0x10, 0xa4, 0x6d,
	0x47, 0x2f, 0x11, 0xbf, 0x49, 0x26, 0xfa, 0x56, 0x69, 0x83, 0x7d, 0x73, 0xa4, 0xde, 0x07, 0xf5,
	0x6c, 0xb8, 0x99, 0xcb, 0xb1, 0xca, 0xf6, 0xcb, 0xb2, 0xd8, 0x51, 0x7d, 0x98, 0xaa, 0x79, 0xc4,
	0x8b, 0xea, 0x52, 0x7d, 0xfc, 0xaf, 0xce, 0xe4, 0xcf, 0x93, 0xf0, 0xa7, 0x87, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0x23, 0xb2, 0xa3, 0x75, 0x01, 0x00, 0x00,
}