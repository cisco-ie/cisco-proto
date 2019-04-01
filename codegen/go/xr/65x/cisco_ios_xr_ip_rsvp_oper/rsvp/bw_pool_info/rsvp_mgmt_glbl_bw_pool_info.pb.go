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
// source: rsvp_mgmt_glbl_bw_pool_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_bw_pool_info

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

type RsvpMgmtGlblBwPoolInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtGlblBwPoolInfo_KEYS) Reset()         { *m = RsvpMgmtGlblBwPoolInfo_KEYS{} }
func (m *RsvpMgmtGlblBwPoolInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblBwPoolInfo_KEYS) ProtoMessage()    {}
func (*RsvpMgmtGlblBwPoolInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bec3003a9fdf089, []int{0}
}

func (m *RsvpMgmtGlblBwPoolInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblBwPoolInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtGlblBwPoolInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS.Merge(m, src)
}
func (m *RsvpMgmtGlblBwPoolInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS.Size(m)
}
func (m *RsvpMgmtGlblBwPoolInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblBwPoolInfo_KEYS proto.InternalMessageInfo

type RsvpMgmtGlblBwPoolInfo struct {
	MaxResPoolPercent             uint32   `protobuf:"varint,50,opt,name=max_res_pool_percent,json=maxResPoolPercent,proto3" json:"max_res_pool_percent,omitempty"`
	Bc0Percent                    uint32   `protobuf:"varint,51,opt,name=bc0_percent,json=bc0Percent,proto3" json:"bc0_percent,omitempty"`
	Bc1Percent                    uint32   `protobuf:"varint,52,opt,name=bc1_percent,json=bc1Percent,proto3" json:"bc1_percent,omitempty"`
	IsMaxResPoolPercentConfigured bool     `protobuf:"varint,53,opt,name=is_max_res_pool_percent_configured,json=isMaxResPoolPercentConfigured,proto3" json:"is_max_res_pool_percent_configured,omitempty"`
	IsBc0PercentConfigured        bool     `protobuf:"varint,54,opt,name=is_bc0_percent_configured,json=isBc0PercentConfigured,proto3" json:"is_bc0_percent_configured,omitempty"`
	IsBc1PercentConfigured        bool     `protobuf:"varint,55,opt,name=is_bc1_percent_configured,json=isBc1PercentConfigured,proto3" json:"is_bc1_percent_configured,omitempty"`
	BandwidthConfigurationModel   string   `protobuf:"bytes,56,opt,name=bandwidth_configuration_model,json=bandwidthConfigurationModel,proto3" json:"bandwidth_configuration_model,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *RsvpMgmtGlblBwPoolInfo) Reset()         { *m = RsvpMgmtGlblBwPoolInfo{} }
func (m *RsvpMgmtGlblBwPoolInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtGlblBwPoolInfo) ProtoMessage()    {}
func (*RsvpMgmtGlblBwPoolInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bec3003a9fdf089, []int{1}
}

func (m *RsvpMgmtGlblBwPoolInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtGlblBwPoolInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtGlblBwPoolInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtGlblBwPoolInfo.Merge(m, src)
}
func (m *RsvpMgmtGlblBwPoolInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtGlblBwPoolInfo.Size(m)
}
func (m *RsvpMgmtGlblBwPoolInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtGlblBwPoolInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtGlblBwPoolInfo proto.InternalMessageInfo

func (m *RsvpMgmtGlblBwPoolInfo) GetMaxResPoolPercent() uint32 {
	if m != nil {
		return m.MaxResPoolPercent
	}
	return 0
}

func (m *RsvpMgmtGlblBwPoolInfo) GetBc0Percent() uint32 {
	if m != nil {
		return m.Bc0Percent
	}
	return 0
}

func (m *RsvpMgmtGlblBwPoolInfo) GetBc1Percent() uint32 {
	if m != nil {
		return m.Bc1Percent
	}
	return 0
}

func (m *RsvpMgmtGlblBwPoolInfo) GetIsMaxResPoolPercentConfigured() bool {
	if m != nil {
		return m.IsMaxResPoolPercentConfigured
	}
	return false
}

func (m *RsvpMgmtGlblBwPoolInfo) GetIsBc0PercentConfigured() bool {
	if m != nil {
		return m.IsBc0PercentConfigured
	}
	return false
}

func (m *RsvpMgmtGlblBwPoolInfo) GetIsBc1PercentConfigured() bool {
	if m != nil {
		return m.IsBc1PercentConfigured
	}
	return false
}

func (m *RsvpMgmtGlblBwPoolInfo) GetBandwidthConfigurationModel() string {
	if m != nil {
		return m.BandwidthConfigurationModel
	}
	return ""
}

func init() {
	proto.RegisterType((*RsvpMgmtGlblBwPoolInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.bw_pool_info.rsvp_mgmt_glbl_bw_pool_info_KEYS")
	proto.RegisterType((*RsvpMgmtGlblBwPoolInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp.bw_pool_info.rsvp_mgmt_glbl_bw_pool_info")
}

func init() { proto.RegisterFile("rsvp_mgmt_glbl_bw_pool_info.proto", fileDescriptor_4bec3003a9fdf089) }

var fileDescriptor_4bec3003a9fdf089 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0x71, 0x8a, 0x20, 0x1a, 0xf1, 0x60, 0x11, 0xa9, 0x8c, 0x61, 0xed, 0xa9, 0x20, 0xd4,
	0xd5, 0xf9, 0xf3, 0xba, 0xe1, 0x41, 0x64, 0x30, 0xea, 0xc9, 0xd3, 0xa3, 0x49, 0xb3, 0xfa, 0xa0,
	0xed, 0x0b, 0x49, 0x75, 0xfd, 0x47, 0xfc, 0x7f, 0x65, 0x1d, 0x8b, 0xd5, 0x95, 0xdd, 0x4a, 0xf3,
	0xf9, 0xbe, 0x3c, 0x08, 0xbb, 0xd4, 0xe6, 0x4b, 0x41, 0x99, 0x97, 0x35, 0xe4, 0x05, 0x2f, 0x80,
	0x2f, 0x41, 0x11, 0x15, 0x80, 0xd5, 0x82, 0x22, 0xa5, 0xa9, 0x26, 0xf7, 0x4a, 0xa0, 0x11, 0x04,
	0x48, 0x06, 0x1a, 0x0d, 0xa8, 0xa0, 0x4d, 0x48, 0x49, 0x1d, 0xad, 0xbe, 0xa2, 0x6e, 0x12, 0x04,
	0xcc, 0xdf, 0x31, 0x11, 0x5e, 0x9f, 0xdf, 0xdf, 0x82, 0xef, 0x3d, 0x36, 0xd8, 0x81, 0xdc, 0x6b,
	0x76, 0x5a, 0xa6, 0x0d, 0x68, 0x69, 0xd6, 0x3f, 0x95, 0xd4, 0x42, 0x56, 0xb5, 0x77, 0xe3, 0x3b,
	0xe1, 0x71, 0x72, 0x52, 0xa6, 0x4d, 0x22, 0xcd, 0x9c, 0xa8, 0x98, 0xaf, 0x0f, 0xdc, 0x0b, 0x76,
	0xc4, 0xc5, 0xc8, 0xba, 0x71, 0xeb, 0x18, 0x17, 0xa3, 0x3f, 0x20, 0xb6, 0xe0, 0x76, 0x03, 0xe2,
	0x0d, 0x78, 0x61, 0x01, 0x1a, 0xe8, 0xbb, 0x15, 0x04, 0x55, 0x0b, 0xcc, 0x3f, 0xb5, 0xcc, 0xbc,
	0x3b, 0xdf, 0x09, 0x0f, 0x92, 0x21, 0x9a, 0xd9, 0xff, 0x15, 0xa6, 0x16, 0xb9, 0x4f, 0xec, 0x1c,
	0x0d, 0x74, 0xf6, 0xe9, 0x4e, 0xb8, 0x6f, 0x27, 0x9c, 0xa1, 0x99, 0xd8, 0xe5, 0x7a, 0xd2, 0xb8,
	0x2f, 0x7d, 0xf8, 0x4d, 0xe3, 0xed, 0x74, 0xc2, 0x86, 0x3c, 0xad, 0xb2, 0x25, 0x66, 0xf5, 0x87,
	0xad, 0xd2, 0x1a, 0xa9, 0x82, 0x92, 0x32, 0x59, 0x78, 0x8f, 0xbe, 0x13, 0x1e, 0x26, 0x03, 0x8b,
	0xa6, 0x5d, 0x33, 0x5b, 0x11, 0xbe, 0xdf, 0xbe, 0xf7, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xfd, 0x1d, 0x2d, 0x14, 0x02, 0x00, 0x00,
}
