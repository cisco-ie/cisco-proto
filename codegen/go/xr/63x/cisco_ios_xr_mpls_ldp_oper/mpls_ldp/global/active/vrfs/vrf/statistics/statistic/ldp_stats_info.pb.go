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
// source: ldp_stats_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_vrfs_vrf_statistics_statistic

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

type LdpStatsInfo_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LsrId                string   `protobuf:"bytes,2,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,3,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpStatsInfo_KEYS) Reset()         { *m = LdpStatsInfo_KEYS{} }
func (m *LdpStatsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpStatsInfo_KEYS) ProtoMessage()    {}
func (*LdpStatsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8321582c74958171, []int{0}
}

func (m *LdpStatsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpStatsInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpStatsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpStatsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpStatsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpStatsInfo_KEYS.Merge(m, src)
}
func (m *LdpStatsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpStatsInfo_KEYS.Size(m)
}
func (m *LdpStatsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpStatsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpStatsInfo_KEYS proto.InternalMessageInfo

func (m *LdpStatsInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpStatsInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpStatsInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

type LdpMsgCounters struct {
	TotalCount             uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	InitCount              uint32   `protobuf:"varint,2,opt,name=init_count,json=initCount,proto3" json:"init_count,omitempty"`
	AddressCount           uint32   `protobuf:"varint,3,opt,name=address_count,json=addressCount,proto3" json:"address_count,omitempty"`
	AddressWithdrawCount   uint32   `protobuf:"varint,4,opt,name=address_withdraw_count,json=addressWithdrawCount,proto3" json:"address_withdraw_count,omitempty"`
	LabelMapCount          uint32   `protobuf:"varint,5,opt,name=label_map_count,json=labelMapCount,proto3" json:"label_map_count,omitempty"`
	LabelWithdrawCount     uint32   `protobuf:"varint,6,opt,name=label_withdraw_count,json=labelWithdrawCount,proto3" json:"label_withdraw_count,omitempty"`
	LabelReleaseCount      uint32   `protobuf:"varint,7,opt,name=label_release_count,json=labelReleaseCount,proto3" json:"label_release_count,omitempty"`
	LabelRequestCount      uint32   `protobuf:"varint,8,opt,name=label_request_count,json=labelRequestCount,proto3" json:"label_request_count,omitempty"`
	LabelAbortRequestCount uint32   `protobuf:"varint,9,opt,name=label_abort_request_count,json=labelAbortRequestCount,proto3" json:"label_abort_request_count,omitempty"`
	NotificationCount      uint32   `protobuf:"varint,10,opt,name=notification_count,json=notificationCount,proto3" json:"notification_count,omitempty"`
	KeepAliveCount         uint32   `protobuf:"varint,11,opt,name=keep_alive_count,json=keepAliveCount,proto3" json:"keep_alive_count,omitempty"`
	IccpRgConnCount        uint32   `protobuf:"varint,12,opt,name=iccp_rg_conn_count,json=iccpRgConnCount,proto3" json:"iccp_rg_conn_count,omitempty"`
	IccpRgDisconnCount     uint32   `protobuf:"varint,13,opt,name=iccp_rg_disconn_count,json=iccpRgDisconnCount,proto3" json:"iccp_rg_disconn_count,omitempty"`
	IccpRgNotifCount       uint32   `protobuf:"varint,14,opt,name=iccp_rg_notif_count,json=iccpRgNotifCount,proto3" json:"iccp_rg_notif_count,omitempty"`
	IccpRgAppDataCount     uint32   `protobuf:"varint,15,opt,name=iccp_rg_app_data_count,json=iccpRgAppDataCount,proto3" json:"iccp_rg_app_data_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *LdpMsgCounters) Reset()         { *m = LdpMsgCounters{} }
func (m *LdpMsgCounters) String() string { return proto.CompactTextString(m) }
func (*LdpMsgCounters) ProtoMessage()    {}
func (*LdpMsgCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_8321582c74958171, []int{1}
}

func (m *LdpMsgCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpMsgCounters.Unmarshal(m, b)
}
func (m *LdpMsgCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpMsgCounters.Marshal(b, m, deterministic)
}
func (m *LdpMsgCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpMsgCounters.Merge(m, src)
}
func (m *LdpMsgCounters) XXX_Size() int {
	return xxx_messageInfo_LdpMsgCounters.Size(m)
}
func (m *LdpMsgCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpMsgCounters.DiscardUnknown(m)
}

var xxx_messageInfo_LdpMsgCounters proto.InternalMessageInfo

func (m *LdpMsgCounters) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *LdpMsgCounters) GetInitCount() uint32 {
	if m != nil {
		return m.InitCount
	}
	return 0
}

func (m *LdpMsgCounters) GetAddressCount() uint32 {
	if m != nil {
		return m.AddressCount
	}
	return 0
}

func (m *LdpMsgCounters) GetAddressWithdrawCount() uint32 {
	if m != nil {
		return m.AddressWithdrawCount
	}
	return 0
}

func (m *LdpMsgCounters) GetLabelMapCount() uint32 {
	if m != nil {
		return m.LabelMapCount
	}
	return 0
}

func (m *LdpMsgCounters) GetLabelWithdrawCount() uint32 {
	if m != nil {
		return m.LabelWithdrawCount
	}
	return 0
}

func (m *LdpMsgCounters) GetLabelReleaseCount() uint32 {
	if m != nil {
		return m.LabelReleaseCount
	}
	return 0
}

func (m *LdpMsgCounters) GetLabelRequestCount() uint32 {
	if m != nil {
		return m.LabelRequestCount
	}
	return 0
}

func (m *LdpMsgCounters) GetLabelAbortRequestCount() uint32 {
	if m != nil {
		return m.LabelAbortRequestCount
	}
	return 0
}

func (m *LdpMsgCounters) GetNotificationCount() uint32 {
	if m != nil {
		return m.NotificationCount
	}
	return 0
}

func (m *LdpMsgCounters) GetKeepAliveCount() uint32 {
	if m != nil {
		return m.KeepAliveCount
	}
	return 0
}

func (m *LdpMsgCounters) GetIccpRgConnCount() uint32 {
	if m != nil {
		return m.IccpRgConnCount
	}
	return 0
}

func (m *LdpMsgCounters) GetIccpRgDisconnCount() uint32 {
	if m != nil {
		return m.IccpRgDisconnCount
	}
	return 0
}

func (m *LdpMsgCounters) GetIccpRgNotifCount() uint32 {
	if m != nil {
		return m.IccpRgNotifCount
	}
	return 0
}

func (m *LdpMsgCounters) GetIccpRgAppDataCount() uint32 {
	if m != nil {
		return m.IccpRgAppDataCount
	}
	return 0
}

type LdpStatsInfo struct {
	IccpEnabled          bool            `protobuf:"varint,50,opt,name=iccp_enabled,json=iccpEnabled,proto3" json:"iccp_enabled,omitempty"`
	MessageOut           *LdpMsgCounters `protobuf:"bytes,51,opt,name=message_out,json=messageOut,proto3" json:"message_out,omitempty"`
	MessageIn            *LdpMsgCounters `protobuf:"bytes,52,opt,name=message_in,json=messageIn,proto3" json:"message_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LdpStatsInfo) Reset()         { *m = LdpStatsInfo{} }
func (m *LdpStatsInfo) String() string { return proto.CompactTextString(m) }
func (*LdpStatsInfo) ProtoMessage()    {}
func (*LdpStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8321582c74958171, []int{2}
}

func (m *LdpStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpStatsInfo.Unmarshal(m, b)
}
func (m *LdpStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpStatsInfo.Marshal(b, m, deterministic)
}
func (m *LdpStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpStatsInfo.Merge(m, src)
}
func (m *LdpStatsInfo) XXX_Size() int {
	return xxx_messageInfo_LdpStatsInfo.Size(m)
}
func (m *LdpStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpStatsInfo proto.InternalMessageInfo

func (m *LdpStatsInfo) GetIccpEnabled() bool {
	if m != nil {
		return m.IccpEnabled
	}
	return false
}

func (m *LdpStatsInfo) GetMessageOut() *LdpMsgCounters {
	if m != nil {
		return m.MessageOut
	}
	return nil
}

func (m *LdpStatsInfo) GetMessageIn() *LdpMsgCounters {
	if m != nil {
		return m.MessageIn
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpStatsInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.vrfs.vrf.statistics.statistic.ldp_stats_info_KEYS")
	proto.RegisterType((*LdpMsgCounters)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.vrfs.vrf.statistics.statistic.ldp_msg_counters")
	proto.RegisterType((*LdpStatsInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.vrfs.vrf.statistics.statistic.ldp_stats_info")
}

func init() { proto.RegisterFile("ldp_stats_info.proto", fileDescriptor_8321582c74958171) }

var fileDescriptor_8321582c74958171 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4b, 0x6f, 0x13, 0x31,
	0x14, 0x85, 0x35, 0x85, 0x3e, 0x72, 0xd3, 0x3c, 0x70, 0xd2, 0x28, 0x5d, 0x20, 0x42, 0x40, 0x28,
	0x12, 0x62, 0x04, 0x69, 0x37, 0x2c, 0xa3, 0xb6, 0x8b, 0x08, 0xd1, 0x4a, 0xd3, 0x05, 0x62, 0x65,
	0xdd, 0xcc, 0x38, 0xc1, 0xc2, 0x33, 0x36, 0xb6, 0x93, 0xb2, 0x44, 0xfc, 0x2d, 0x24, 0x7e, 0x1b,
	0xf2, 0x63, 0xa0, 0x09, 0x5b, 0xc4, 0x26, 0xca, 0x9c, 0xf3, 0x9d, 0x7b, 0x6c, 0x27, 0x1e, 0xe8,
	0x8b, 0x42, 0x51, 0x63, 0xd1, 0x1a, 0xca, 0xab, 0xa5, 0x4c, 0x95, 0x96, 0x56, 0x92, 0x9b, 0x9c,
	0x9b, 0x5c, 0x52, 0x2e, 0x0d, 0xfd, 0xaa, 0x69, 0xa9, 0x84, 0xa1, 0x8e, 0x93, 0x8a, 0xe9, 0xb4,
	0x7e, 0x4a, 0x57, 0x42, 0x2e, 0x50, 0xa4, 0x98, 0x5b, 0xbe, 0x61, 0xe9, 0x46, 0x2f, 0x8d, 0xfb,
	0x48, 0xdd, 0x30, 0x6e, 0x2c, 0xcf, 0xcd, 0x9f, 0xaf, 0xe3, 0x12, 0x7a, 0xdb, 0x45, 0xf4, 0xdd,
	0xd5, 0xc7, 0x5b, 0x72, 0x0a, 0x47, 0x1b, 0xbd, 0xa4, 0x15, 0x96, 0x6c, 0x98, 0x8c, 0x92, 0x49,
	0x23, 0x3b, 0xdc, 0xe8, 0xe5, 0x35, 0x96, 0x8c, 0x9c, 0xc0, 0x81, 0x30, 0x9a, 0xf2, 0x62, 0xb8,
	0xe7, 0x8d, 0x7d, 0x61, 0xf4, 0xbc, 0x20, 0xcf, 0xa1, 0x2d, 0x70, 0xc1, 0x04, 0x35, 0x0a, 0x73,
	0xe6, 0xec, 0x07, 0xa3, 0x64, 0xd2, 0xca, 0x8e, 0xbd, 0x7a, 0xeb, 0xc4, 0x79, 0x31, 0xfe, 0xb9,
	0x0f, 0x5d, 0xd7, 0x57, 0x9a, 0x15, 0xcd, 0xe5, 0xba, 0xb2, 0x4c, 0x1b, 0xf2, 0x04, 0x9a, 0x56,
	0x5a, 0x14, 0x41, 0xf1, 0x7d, 0xad, 0x0c, 0xbc, 0x74, 0xe1, 0x14, 0xf2, 0x18, 0x80, 0x57, 0xdc,
	0x46, 0x7f, 0xcf, 0xfb, 0x0d, 0xa7, 0x04, 0xfb, 0x19, 0xb4, 0xb0, 0x28, 0x34, 0x33, 0x26, 0x12,
	0xb1, 0x39, 0x8a, 0x01, 0x3a, 0x87, 0x41, 0x0d, 0xdd, 0x71, 0xfb, 0xa9, 0xd0, 0x78, 0x17, 0xe9,
	0x87, 0x9e, 0xee, 0x47, 0xf7, 0x43, 0x34, 0x43, 0xea, 0x05, 0x74, 0xc2, 0xae, 0x4a, 0x54, 0x11,
	0xdf, 0xf7, 0x78, 0xcb, 0xcb, 0xef, 0x51, 0x05, 0xee, 0x35, 0xf4, 0x03, 0xb7, 0x33, 0xfb, 0xc0,
	0xc3, 0xc4, 0x7b, 0xdb, 0x93, 0x53, 0xe8, 0x85, 0x84, 0x66, 0x82, 0xa1, 0x61, 0x31, 0x70, 0xe8,
	0x03, 0x8f, 0xbc, 0x95, 0x05, 0xe7, 0x2f, 0xfe, 0xcb, 0x9a, 0x99, 0xfa, 0x30, 0x8e, 0xb6, 0x78,
	0xef, 0x04, 0xfe, 0x2d, 0x9c, 0x06, 0x1e, 0x17, 0x52, 0xdb, 0x9d, 0x54, 0xc3, 0xa7, 0x06, 0x1e,
	0x98, 0x39, 0x7f, 0x2b, 0xfa, 0x0a, 0x48, 0x25, 0x2d, 0x5f, 0xf2, 0x1c, 0x2d, 0x97, 0x55, 0xcc,
	0x40, 0x68, 0xba, 0xef, 0x04, 0x7c, 0x02, 0xdd, 0xcf, 0x8c, 0x29, 0x8a, 0x82, 0x6f, 0xea, 0x6d,
	0x34, 0x3d, 0xdc, 0x76, 0xfa, 0xcc, 0xc9, 0x81, 0x7c, 0x09, 0x84, 0xe7, 0xb9, 0xa2, 0xda, 0xfd,
	0xf8, 0x55, 0x3d, 0xf8, 0xd8, 0xb3, 0x1d, 0xe7, 0x64, 0xab, 0x0b, 0x59, 0xc5, 0xb1, 0x6f, 0xe0,
	0xa4, 0x86, 0x0b, 0xf7, 0xa7, 0xff, 0xcd, 0xb7, 0xc2, 0x99, 0x06, 0xfe, 0x32, 0x58, 0xf5, 0xc2,
	0x7b, 0x75, 0xc4, 0x2f, 0x33, 0x06, 0xda, 0x3e, 0xd0, 0x0d, 0x81, 0x6b, 0x67, 0x04, 0x7c, 0x0a,
	0x83, 0x1a, 0x47, 0xa5, 0x68, 0x81, 0x16, 0x63, 0xa2, 0x73, 0xbf, 0x62, 0xa6, 0xd4, 0x25, 0x5a,
	0xf4, 0x99, 0xf1, 0x8f, 0x3d, 0x68, 0x6f, 0x5f, 0x18, 0xf2, 0x14, 0x8e, 0xfd, 0x18, 0x56, 0xe1,
	0x42, 0xb0, 0x62, 0x38, 0x1d, 0x25, 0x93, 0xa3, 0xac, 0xe9, 0xb4, 0xab, 0x20, 0x91, 0xef, 0x09,
	0x34, 0x4b, 0x66, 0x0c, 0xae, 0x18, 0x95, 0x6b, 0x3b, 0x3c, 0x1b, 0x25, 0x93, 0xe6, 0x14, 0xd3,
	0x7f, 0x7c, 0x9b, 0xd3, 0xdd, 0xab, 0x95, 0x41, 0x6c, 0xbd, 0x59, 0x5b, 0xf2, 0x2d, 0x81, 0xfa,
	0x91, 0xf2, 0x6a, 0x78, 0xfe, 0xbf, 0xd6, 0xd0, 0x88, 0xa5, 0xf3, 0x6a, 0x71, 0xe0, 0xdf, 0x62,
	0x67, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xc2, 0xc6, 0x99, 0xdd, 0x04, 0x00, 0x00,
}
