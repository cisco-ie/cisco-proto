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
// source: mpls_lm_summary_common_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_standby_link_summary

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

type MplsLmSummaryCommonInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmSummaryCommonInfo_KEYS) Reset()         { *m = MplsLmSummaryCommonInfo_KEYS{} }
func (m *MplsLmSummaryCommonInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryCommonInfo_KEYS) ProtoMessage()    {}
func (*MplsLmSummaryCommonInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf35f21178dade80, []int{0}
}

func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Merge(m, src)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.Size(m)
}
func (m *MplsLmSummaryCommonInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryCommonInfo_KEYS proto.InternalMessageInfo

type MplsLmSummaryAreaInfo struct {
	AreaId                   string   `protobuf:"bytes,1,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Protocol                 string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	IsFlooded                bool     `protobuf:"varint,3,opt,name=is_flooded,json=isFlooded,proto3" json:"is_flooded,omitempty"`
	IsPeriodicFloodingOn     bool     `protobuf:"varint,4,opt,name=is_periodic_flooding_on,json=isPeriodicFloodingOn,proto3" json:"is_periodic_flooding_on,omitempty"`
	PeriodicFloodingInterval uint32   `protobuf:"varint,5,opt,name=periodic_flooding_interval,json=periodicFloodingInterval,proto3" json:"periodic_flooding_interval,omitempty"`
	LinksFlooded             uint32   `protobuf:"varint,6,opt,name=links_flooded,json=linksFlooded,proto3" json:"links_flooded,omitempty"`
	SystemId                 string   `protobuf:"bytes,7,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	LocalNodeRouterId        string   `protobuf:"bytes,8,opt,name=local_node_router_id,json=localNodeRouterId,proto3" json:"local_node_router_id,omitempty"`
	IgpNeighbors             uint32   `protobuf:"varint,9,opt,name=igp_neighbors,json=igpNeighbors,proto3" json:"igp_neighbors,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *MplsLmSummaryAreaInfo) Reset()         { *m = MplsLmSummaryAreaInfo{} }
func (m *MplsLmSummaryAreaInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryAreaInfo) ProtoMessage()    {}
func (*MplsLmSummaryAreaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf35f21178dade80, []int{1}
}

func (m *MplsLmSummaryAreaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Unmarshal(m, b)
}
func (m *MplsLmSummaryAreaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmSummaryAreaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryAreaInfo.Merge(m, src)
}
func (m *MplsLmSummaryAreaInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryAreaInfo.Size(m)
}
func (m *MplsLmSummaryAreaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryAreaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryAreaInfo proto.InternalMessageInfo

func (m *MplsLmSummaryAreaInfo) GetAreaId() string {
	if m != nil {
		return m.AreaId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetIsFlooded() bool {
	if m != nil {
		return m.IsFlooded
	}
	return false
}

func (m *MplsLmSummaryAreaInfo) GetIsPeriodicFloodingOn() bool {
	if m != nil {
		return m.IsPeriodicFloodingOn
	}
	return false
}

func (m *MplsLmSummaryAreaInfo) GetPeriodicFloodingInterval() uint32 {
	if m != nil {
		return m.PeriodicFloodingInterval
	}
	return 0
}

func (m *MplsLmSummaryAreaInfo) GetLinksFlooded() uint32 {
	if m != nil {
		return m.LinksFlooded
	}
	return 0
}

func (m *MplsLmSummaryAreaInfo) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetLocalNodeRouterId() string {
	if m != nil {
		return m.LocalNodeRouterId
	}
	return ""
}

func (m *MplsLmSummaryAreaInfo) GetIgpNeighbors() uint32 {
	if m != nil {
		return m.IgpNeighbors
	}
	return 0
}

type MplsLmSummaryDarkbw struct {
	IsBandwidthAccountEnabled bool     `protobuf:"varint,1,opt,name=is_bandwidth_account_enabled,json=isBandwidthAccountEnabled,proto3" json:"is_bandwidth_account_enabled,omitempty"`
	SampleInterval            uint32   `protobuf:"varint,2,opt,name=sample_interval,json=sampleInterval,proto3" json:"sample_interval,omitempty"`
	SampleTimeRemaining       uint32   `protobuf:"varint,3,opt,name=sample_time_remaining,json=sampleTimeRemaining,proto3" json:"sample_time_remaining,omitempty"`
	ApplicationInterval       uint32   `protobuf:"varint,4,opt,name=application_interval,json=applicationInterval,proto3" json:"application_interval,omitempty"`
	ApplicationTimeRemaining  uint32   `protobuf:"varint,5,opt,name=application_time_remaining,json=applicationTimeRemaining,proto3" json:"application_time_remaining,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *MplsLmSummaryDarkbw) Reset()         { *m = MplsLmSummaryDarkbw{} }
func (m *MplsLmSummaryDarkbw) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryDarkbw) ProtoMessage()    {}
func (*MplsLmSummaryDarkbw) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf35f21178dade80, []int{2}
}

func (m *MplsLmSummaryDarkbw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryDarkbw.Unmarshal(m, b)
}
func (m *MplsLmSummaryDarkbw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryDarkbw.Marshal(b, m, deterministic)
}
func (m *MplsLmSummaryDarkbw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryDarkbw.Merge(m, src)
}
func (m *MplsLmSummaryDarkbw) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryDarkbw.Size(m)
}
func (m *MplsLmSummaryDarkbw) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryDarkbw.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryDarkbw proto.InternalMessageInfo

func (m *MplsLmSummaryDarkbw) GetIsBandwidthAccountEnabled() bool {
	if m != nil {
		return m.IsBandwidthAccountEnabled
	}
	return false
}

func (m *MplsLmSummaryDarkbw) GetSampleInterval() uint32 {
	if m != nil {
		return m.SampleInterval
	}
	return 0
}

func (m *MplsLmSummaryDarkbw) GetSampleTimeRemaining() uint32 {
	if m != nil {
		return m.SampleTimeRemaining
	}
	return 0
}

func (m *MplsLmSummaryDarkbw) GetApplicationInterval() uint32 {
	if m != nil {
		return m.ApplicationInterval
	}
	return 0
}

func (m *MplsLmSummaryDarkbw) GetApplicationTimeRemaining() uint32 {
	if m != nil {
		return m.ApplicationTimeRemaining
	}
	return 0
}

type MplsLmSummaryCommonInfo struct {
	IsRoleStandby           bool                     `protobuf:"varint,50,opt,name=is_role_standby,json=isRoleStandby,proto3" json:"is_role_standby,omitempty"`
	Links                   uint32                   `protobuf:"varint,51,opt,name=links,proto3" json:"links,omitempty"`
	MaximumLinks            uint32                   `protobuf:"varint,52,opt,name=maximum_links,json=maximumLinks,proto3" json:"maximum_links,omitempty"`
	IsFloodingEnabled       bool                     `protobuf:"varint,53,opt,name=is_flooding_enabled,json=isFloodingEnabled,proto3" json:"is_flooding_enabled,omitempty"`
	AreasSummary            []*MplsLmSummaryAreaInfo `protobuf:"bytes,54,rep,name=areas_summary,json=areasSummary,proto3" json:"areas_summary,omitempty"`
	BandwidthAccountSummary *MplsLmSummaryDarkbw     `protobuf:"bytes,55,opt,name=bandwidth_account_summary,json=bandwidthAccountSummary,proto3" json:"bandwidth_account_summary,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *MplsLmSummaryCommonInfo) Reset()         { *m = MplsLmSummaryCommonInfo{} }
func (m *MplsLmSummaryCommonInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmSummaryCommonInfo) ProtoMessage()    {}
func (*MplsLmSummaryCommonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf35f21178dade80, []int{3}
}

func (m *MplsLmSummaryCommonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Unmarshal(m, b)
}
func (m *MplsLmSummaryCommonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmSummaryCommonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmSummaryCommonInfo.Merge(m, src)
}
func (m *MplsLmSummaryCommonInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmSummaryCommonInfo.Size(m)
}
func (m *MplsLmSummaryCommonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmSummaryCommonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmSummaryCommonInfo proto.InternalMessageInfo

func (m *MplsLmSummaryCommonInfo) GetIsRoleStandby() bool {
	if m != nil {
		return m.IsRoleStandby
	}
	return false
}

func (m *MplsLmSummaryCommonInfo) GetLinks() uint32 {
	if m != nil {
		return m.Links
	}
	return 0
}

func (m *MplsLmSummaryCommonInfo) GetMaximumLinks() uint32 {
	if m != nil {
		return m.MaximumLinks
	}
	return 0
}

func (m *MplsLmSummaryCommonInfo) GetIsFloodingEnabled() bool {
	if m != nil {
		return m.IsFloodingEnabled
	}
	return false
}

func (m *MplsLmSummaryCommonInfo) GetAreasSummary() []*MplsLmSummaryAreaInfo {
	if m != nil {
		return m.AreasSummary
	}
	return nil
}

func (m *MplsLmSummaryCommonInfo) GetBandwidthAccountSummary() *MplsLmSummaryDarkbw {
	if m != nil {
		return m.BandwidthAccountSummary
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsLmSummaryCommonInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.link_summary.mpls_lm_summary_common_info_KEYS")
	proto.RegisterType((*MplsLmSummaryAreaInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.link_summary.mpls_lm_summary_area_info")
	proto.RegisterType((*MplsLmSummaryDarkbw)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.link_summary.mpls_lm_summary_darkbw")
	proto.RegisterType((*MplsLmSummaryCommonInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.link_summary.mpls_lm_summary_common_info")
}

func init() { proto.RegisterFile("mpls_lm_summary_common_info.proto", fileDescriptor_bf35f21178dade80) }

var fileDescriptor_bf35f21178dade80 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x4e, 0x14, 0x4d,
	0x14, 0xcd, 0x0c, 0x7f, 0x33, 0x17, 0xe6, 0x23, 0x14, 0xf3, 0x49, 0x03, 0x9a, 0x8c, 0x63, 0xa2,
	0xb3, 0x6a, 0xe3, 0x20, 0xea, 0x82, 0xc4, 0x68, 0x02, 0xc9, 0x44, 0x83, 0xa4, 0xc7, 0x8d, 0xab,
	0x4a, 0x75, 0x57, 0xd1, 0xdc, 0x50, 0x3f, 0x9d, 0xaa, 0x1e, 0x81, 0x77, 0xf0, 0x3d, 0x7c, 0x06,
	0xdf, 0xcd, 0x85, 0xe9, 0xaa, 0x6e, 0xc0, 0x21, 0xb2, 0x30, 0x2e, 0xeb, 0x9e, 0x73, 0xfb, 0xdc,
	0xba, 0xe7, 0x54, 0xc3, 0x63, 0x55, 0x48, 0x47, 0xa5, 0xa2, 0x6e, 0xa6, 0x14, 0xb3, 0x57, 0x34,
	0x33, 0x4a, 0x19, 0x4d, 0x51, 0x9f, 0x9a, 0xb8, 0xb0, 0xa6, 0x34, 0xe4, 0x4d, 0x86, 0x2e, 0x33,
	0x14, 0x8d, 0xa3, 0x97, 0x96, 0x7a, 0x7e, 0x29, 0xa8, 0x29, 0x84, 0x8d, 0x43, 0x73, 0xc6, 0x32,
	0xea, 0x4a, 0xa6, 0x79, 0x7a, 0x15, 0x4b, 0xd4, 0xe7, 0xcd, 0xb7, 0x86, 0x43, 0x18, 0xdc, 0xf3,
	0x79, 0xfa, 0xe1, 0xf0, 0xcb, 0x74, 0xf8, 0xb3, 0x0d, 0xdb, 0xf3, 0x24, 0x66, 0x05, 0xf3, 0x14,
	0xb2, 0x05, 0x2b, 0xe1, 0xc0, 0xa3, 0xd6, 0xa0, 0x35, 0xea, 0x26, 0xcb, 0xd5, 0x71, 0xc2, 0xc9,
	0x0e, 0x74, 0xfc, 0x74, 0x99, 0x91, 0x51, 0xdb, 0x23, 0xd7, 0x67, 0xf2, 0x08, 0x00, 0x1d, 0x3d,
	0x95, 0xc6, 0x70, 0xc1, 0xa3, 0x85, 0x41, 0x6b, 0xd4, 0x49, 0xba, 0xe8, 0x8e, 0x42, 0x81, 0xec,
	0xc3, 0x16, 0x3a, 0x5a, 0x08, 0x8b, 0x86, 0x63, 0x16, 0x78, 0xa8, 0x73, 0x6a, 0x74, 0xb4, 0xe8,
	0xb9, 0x7d, 0x74, 0x27, 0x35, 0x7a, 0x54, 0x83, 0x9f, 0x34, 0x39, 0x80, 0x9d, 0xbb, 0x3d, 0xa8,
	0x4b, 0x61, 0xbf, 0x32, 0x19, 0x2d, 0x0d, 0x5a, 0xa3, 0x5e, 0x12, 0x15, 0x73, 0x7d, 0x93, 0x1a,
	0x27, 0x4f, 0xa0, 0x57, 0xad, 0xe6, 0x66, 0xac, 0x65, 0xdf, 0xb0, 0xe6, 0x8b, 0xcd, 0x64, 0xbb,
	0xd0, 0x75, 0x57, 0xae, 0x14, 0xaa, 0xba, 0xef, 0x4a, 0xb8, 0x55, 0x28, 0x4c, 0x38, 0x79, 0x0e,
	0x7d, 0x69, 0x32, 0x26, 0xa9, 0x36, 0x5c, 0x50, 0x6b, 0x66, 0xa5, 0xb0, 0x15, 0xaf, 0xe3, 0x79,
	0x1b, 0x1e, 0x3b, 0x36, 0x5c, 0x24, 0x1e, 0x99, 0xf0, 0x4a, 0x12, 0xf3, 0x82, 0x6a, 0x81, 0xf9,
	0x59, 0x6a, 0xac, 0x8b, 0xba, 0x41, 0x12, 0xf3, 0xe2, 0xb8, 0xa9, 0x0d, 0xbf, 0xb7, 0xe1, 0xc1,
	0xfc, 0xfa, 0x39, 0xb3, 0xe7, 0xe9, 0x05, 0x79, 0x0b, 0x0f, 0xd1, 0xd1, 0x94, 0x69, 0x7e, 0x81,
	0xbc, 0x3c, 0xa3, 0x2c, 0xcb, 0xcc, 0x4c, 0x97, 0x54, 0x68, 0x96, 0x4a, 0x11, 0x0c, 0xe9, 0x24,
	0xdb, 0xe8, 0xde, 0x37, 0x94, 0x77, 0x81, 0x71, 0x18, 0x08, 0xe4, 0x19, 0xac, 0x3b, 0xa6, 0x0a,
	0x29, 0x6e, 0xd6, 0xd4, 0xf6, 0x23, 0xfc, 0x17, 0xca, 0xd7, 0xcb, 0x19, 0xc3, 0xff, 0x35, 0xb1,
	0x44, 0x25, 0xa8, 0x15, 0x8a, 0xa1, 0x46, 0x9d, 0x7b, 0xef, 0x7a, 0xc9, 0x66, 0x00, 0x3f, 0xa3,
	0x12, 0x49, 0x03, 0x91, 0x17, 0xd0, 0x67, 0x45, 0x21, 0x31, 0x63, 0x25, 0xfa, 0x40, 0xd5, 0x0a,
	0x8b, 0xa1, 0xe5, 0x16, 0x76, 0x2d, 0x73, 0x00, 0x3b, 0xb7, 0x5b, 0xe6, 0xb4, 0x6a, 0x07, 0x6f,
	0x31, 0x7e, 0x13, 0x1c, 0xfe, 0x58, 0x80, 0xdd, 0x7b, 0xd2, 0x4c, 0x9e, 0xc2, 0x3a, 0x3a, 0x6a,
	0x8d, 0x14, 0xcd, 0x63, 0x88, 0xc6, 0x7e, 0x43, 0x3d, 0x74, 0x89, 0x91, 0x62, 0x1a, 0x8a, 0xa4,
	0x0f, 0x4b, 0xde, 0xf4, 0x68, 0xcf, 0x0b, 0x86, 0x43, 0x65, 0x96, 0x62, 0x97, 0xa8, 0x66, 0x8a,
	0x06, 0xf4, 0x65, 0x30, 0xab, 0x2e, 0x7e, 0xf4, 0xa4, 0x18, 0x36, 0x9b, 0x60, 0x57, 0xe1, 0x6b,
	0x8c, 0xd8, 0xf7, 0x32, 0x1b, 0x75, 0xc2, 0x51, 0xe7, 0x8d, 0x01, 0x97, 0xd0, 0xab, 0x9e, 0x8b,
	0x6b, 0xe6, 0x8d, 0x5e, 0x0d, 0x16, 0x46, 0xab, 0xe3, 0x69, 0xfc, 0xb7, 0x2f, 0x3a, 0xfe, 0xe3,
	0x4b, 0x4d, 0xd6, 0xbc, 0xd2, 0x34, 0xd4, 0xc9, 0xb7, 0x16, 0x6c, 0xdf, 0x4d, 0x4e, 0x33, 0xc6,
	0xeb, 0x41, 0x6b, 0xb4, 0x3a, 0x3e, 0xf9, 0x77, 0x63, 0x84, 0xc4, 0x26, 0x5b, 0xe9, 0x5c, 0x12,
	0xeb, 0x71, 0xd2, 0x65, 0xff, 0x6f, 0xd8, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x03, 0xe0, 0xff,
	0xa1, 0xee, 0x04, 0x00, 0x00,
}
