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
// source: pm_mpls_ldp_bag.proto

package cisco_ios_xr_manageability_perfmgmt_oper_perf_mgmt_monitor_mpls_ldp_neighbors_ldp_neighbor_samples_sample

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

type PmMplsLdpBag_KEYS struct {
	Nbr                  string   `protobuf:"bytes,1,opt,name=nbr,proto3" json:"nbr,omitempty"`
	SampleId             int32    `protobuf:"varint,2,opt,name=sample_id,json=sampleId,proto3" json:"sample_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmMplsLdpBag_KEYS) Reset()         { *m = PmMplsLdpBag_KEYS{} }
func (m *PmMplsLdpBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmMplsLdpBag_KEYS) ProtoMessage()    {}
func (*PmMplsLdpBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22717b8ac374911, []int{0}
}

func (m *PmMplsLdpBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmMplsLdpBag_KEYS.Unmarshal(m, b)
}
func (m *PmMplsLdpBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmMplsLdpBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmMplsLdpBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmMplsLdpBag_KEYS.Merge(m, src)
}
func (m *PmMplsLdpBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmMplsLdpBag_KEYS.Size(m)
}
func (m *PmMplsLdpBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmMplsLdpBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmMplsLdpBag_KEYS proto.InternalMessageInfo

func (m *PmMplsLdpBag_KEYS) GetNbr() string {
	if m != nil {
		return m.Nbr
	}
	return ""
}

func (m *PmMplsLdpBag_KEYS) GetSampleId() int32 {
	if m != nil {
		return m.SampleId
	}
	return 0
}

type PmMplsLdpBag struct {
	TimeStamp               uint64   `protobuf:"varint,50,opt,name=time_stamp,json=timeStamp,proto3" json:"time_stamp,omitempty"`
	TotalMsgsSent           uint32   `protobuf:"varint,51,opt,name=total_msgs_sent,json=totalMsgsSent,proto3" json:"total_msgs_sent,omitempty"`
	TotalMsgsRcvd           uint32   `protobuf:"varint,52,opt,name=total_msgs_rcvd,json=totalMsgsRcvd,proto3" json:"total_msgs_rcvd,omitempty"`
	InitMsgsSent            uint32   `protobuf:"varint,53,opt,name=init_msgs_sent,json=initMsgsSent,proto3" json:"init_msgs_sent,omitempty"`
	InitMsgsRcvd            uint32   `protobuf:"varint,54,opt,name=init_msgs_rcvd,json=initMsgsRcvd,proto3" json:"init_msgs_rcvd,omitempty"`
	AddressMsgsSent         uint32   `protobuf:"varint,55,opt,name=address_msgs_sent,json=addressMsgsSent,proto3" json:"address_msgs_sent,omitempty"`
	AddressMsgsRcvd         uint32   `protobuf:"varint,56,opt,name=address_msgs_rcvd,json=addressMsgsRcvd,proto3" json:"address_msgs_rcvd,omitempty"`
	AddressWithdrawMsgsSent uint32   `protobuf:"varint,57,opt,name=address_withdraw_msgs_sent,json=addressWithdrawMsgsSent,proto3" json:"address_withdraw_msgs_sent,omitempty"`
	AddressWithdrawMsgsRcvd uint32   `protobuf:"varint,58,opt,name=address_withdraw_msgs_rcvd,json=addressWithdrawMsgsRcvd,proto3" json:"address_withdraw_msgs_rcvd,omitempty"`
	LabelMappingMsgsSent    uint32   `protobuf:"varint,59,opt,name=label_mapping_msgs_sent,json=labelMappingMsgsSent,proto3" json:"label_mapping_msgs_sent,omitempty"`
	LabelMappingMsgsRcvd    uint32   `protobuf:"varint,60,opt,name=label_mapping_msgs_rcvd,json=labelMappingMsgsRcvd,proto3" json:"label_mapping_msgs_rcvd,omitempty"`
	LabelWithdrawMsgsSent   uint32   `protobuf:"varint,61,opt,name=label_withdraw_msgs_sent,json=labelWithdrawMsgsSent,proto3" json:"label_withdraw_msgs_sent,omitempty"`
	LabelWithdrawMsgsRcvd   uint32   `protobuf:"varint,62,opt,name=label_withdraw_msgs_rcvd,json=labelWithdrawMsgsRcvd,proto3" json:"label_withdraw_msgs_rcvd,omitempty"`
	LabelReleaseMsgsSent    uint32   `protobuf:"varint,63,opt,name=label_release_msgs_sent,json=labelReleaseMsgsSent,proto3" json:"label_release_msgs_sent,omitempty"`
	LabelReleaseMsgsRcvd    uint32   `protobuf:"varint,64,opt,name=label_release_msgs_rcvd,json=labelReleaseMsgsRcvd,proto3" json:"label_release_msgs_rcvd,omitempty"`
	NotificationMsgsSent    uint32   `protobuf:"varint,65,opt,name=notification_msgs_sent,json=notificationMsgsSent,proto3" json:"notification_msgs_sent,omitempty"`
	NotificationMsgsRcvd    uint32   `protobuf:"varint,66,opt,name=notification_msgs_rcvd,json=notificationMsgsRcvd,proto3" json:"notification_msgs_rcvd,omitempty"`
	KeepaliveMsgsSent       uint32   `protobuf:"varint,67,opt,name=keepalive_msgs_sent,json=keepaliveMsgsSent,proto3" json:"keepalive_msgs_sent,omitempty"`
	KeepaliveMsgsRcvd       uint32   `protobuf:"varint,68,opt,name=keepalive_msgs_rcvd,json=keepaliveMsgsRcvd,proto3" json:"keepalive_msgs_rcvd,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *PmMplsLdpBag) Reset()         { *m = PmMplsLdpBag{} }
func (m *PmMplsLdpBag) String() string { return proto.CompactTextString(m) }
func (*PmMplsLdpBag) ProtoMessage()    {}
func (*PmMplsLdpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e22717b8ac374911, []int{1}
}

func (m *PmMplsLdpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmMplsLdpBag.Unmarshal(m, b)
}
func (m *PmMplsLdpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmMplsLdpBag.Marshal(b, m, deterministic)
}
func (m *PmMplsLdpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmMplsLdpBag.Merge(m, src)
}
func (m *PmMplsLdpBag) XXX_Size() int {
	return xxx_messageInfo_PmMplsLdpBag.Size(m)
}
func (m *PmMplsLdpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmMplsLdpBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmMplsLdpBag proto.InternalMessageInfo

func (m *PmMplsLdpBag) GetTimeStamp() uint64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *PmMplsLdpBag) GetTotalMsgsSent() uint32 {
	if m != nil {
		return m.TotalMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetTotalMsgsRcvd() uint32 {
	if m != nil {
		return m.TotalMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetInitMsgsSent() uint32 {
	if m != nil {
		return m.InitMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetInitMsgsRcvd() uint32 {
	if m != nil {
		return m.InitMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetAddressMsgsSent() uint32 {
	if m != nil {
		return m.AddressMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetAddressMsgsRcvd() uint32 {
	if m != nil {
		return m.AddressMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetAddressWithdrawMsgsSent() uint32 {
	if m != nil {
		return m.AddressWithdrawMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetAddressWithdrawMsgsRcvd() uint32 {
	if m != nil {
		return m.AddressWithdrawMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelMappingMsgsSent() uint32 {
	if m != nil {
		return m.LabelMappingMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelMappingMsgsRcvd() uint32 {
	if m != nil {
		return m.LabelMappingMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelWithdrawMsgsSent() uint32 {
	if m != nil {
		return m.LabelWithdrawMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelWithdrawMsgsRcvd() uint32 {
	if m != nil {
		return m.LabelWithdrawMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelReleaseMsgsSent() uint32 {
	if m != nil {
		return m.LabelReleaseMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetLabelReleaseMsgsRcvd() uint32 {
	if m != nil {
		return m.LabelReleaseMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetNotificationMsgsSent() uint32 {
	if m != nil {
		return m.NotificationMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetNotificationMsgsRcvd() uint32 {
	if m != nil {
		return m.NotificationMsgsRcvd
	}
	return 0
}

func (m *PmMplsLdpBag) GetKeepaliveMsgsSent() uint32 {
	if m != nil {
		return m.KeepaliveMsgsSent
	}
	return 0
}

func (m *PmMplsLdpBag) GetKeepaliveMsgsRcvd() uint32 {
	if m != nil {
		return m.KeepaliveMsgsRcvd
	}
	return 0
}

func init() {
	proto.RegisterType((*PmMplsLdpBag_KEYS)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.mpls.ldp_neighbors.ldp_neighbor.samples.sample.pm_mpls_ldp_bag_KEYS")
	proto.RegisterType((*PmMplsLdpBag)(nil), "cisco_ios_xr_manageability_perfmgmt_oper.perf_mgmt.monitor.mpls.ldp_neighbors.ldp_neighbor.samples.sample.pm_mpls_ldp_bag")
}

func init() { proto.RegisterFile("pm_mpls_ldp_bag.proto", fileDescriptor_e22717b8ac374911) }

var fileDescriptor_e22717b8ac374911 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0xbe, 0xb3, 0xa2, 0x84, 0x9a, 0x96, 0x5a, 0x20, 0xa4, 0xa8, 0x42, 0x28, 0xe2,
	0xe0, 0x03, 0x6d, 0x29, 0x50, 0xbe, 0xa1, 0x07, 0x84, 0x7a, 0x71, 0x0e, 0x88, 0xd3, 0x6a, 0x6d,
	0x6f, 0xdd, 0x11, 0xfb, 0xa5, 0xdd, 0x55, 0x02, 0xff, 0x8c, 0x9f, 0x87, 0x76, 0x1c, 0xac, 0xb5,
	0xe3, 0x9c, 0xda, 0x9d, 0x79, 0x9f, 0xf7, 0xf1, 0x5c, 0x42, 0xf6, 0x8d, 0xa4, 0xd2, 0x08, 0x47,
	0x45, 0x6d, 0x68, 0xc9, 0x9a, 0xdc, 0x58, 0xed, 0x75, 0x0a, 0x15, 0xb8, 0x4a, 0x53, 0xd0, 0x8e,
	0xfe, 0xb6, 0x54, 0x32, 0xc5, 0x1a, 0xce, 0x4a, 0x10, 0xe0, 0xff, 0x50, 0xc3, 0xed, 0xa5, 0x6c,
	0xa4, 0xa7, 0xda, 0x70, 0x9b, 0x87, 0x17, 0x0d, 0xcf, 0x5c, 0x6a, 0x05, 0x5e, 0xdb, 0x3c, 0xd4,
	0xe5, 0xa1, 0x4e, 0x71, 0x68, 0xae, 0x4a, 0x6d, 0xfb, 0xaf, 0xdc, 0x31, 0x69, 0x04, 0x77, 0xeb,
	0xbf, 0x87, 0xe7, 0x64, 0x6f, 0xf0, 0x0d, 0xf4, 0xfb, 0xf9, 0xcf, 0x45, 0x7a, 0x9f, 0x5c, 0x57,
	0xa5, 0xcd, 0x92, 0x59, 0x32, 0x9f, 0x14, 0xe1, 0xdf, 0xf4, 0x31, 0x99, 0xb4, 0x0c, 0x85, 0x3a,
	0xbb, 0x36, 0x4b, 0xe6, 0x37, 0x8b, 0x3b, 0xed, 0xe0, 0x5b, 0x7d, 0xf8, 0xf7, 0x36, 0x99, 0x0e,
	0x7a, 0xd2, 0x27, 0x84, 0x78, 0x90, 0x9c, 0x3a, 0xcf, 0xa4, 0xc9, 0x5e, 0xcc, 0x92, 0xf9, 0x8d,
	0x62, 0x12, 0x26, 0x8b, 0x30, 0x48, 0x9f, 0x91, 0xa9, 0xd7, 0x9e, 0x09, 0x2a, 0x5d, 0xe3, 0xa8,
	0xe3, 0xca, 0x67, 0x47, 0xb3, 0x64, 0xbe, 0x53, 0xec, 0xe0, 0xf8, 0xc2, 0x35, 0x6e, 0xc1, 0x95,
	0x1f, 0xe4, 0x6c, 0xb5, 0xac, 0xb3, 0xe3, 0x41, 0xae, 0xa8, 0x96, 0x75, 0xfa, 0x94, 0xdc, 0x03,
	0x05, 0x3e, 0xaa, 0x3b, 0xc1, 0xd8, 0xdd, 0x30, 0xed, 0xda, 0x7a, 0x29, 0x2c, 0x7b, 0xd9, 0x4f,
	0x61, 0xd7, 0x73, 0xb2, 0xcb, 0xea, 0xda, 0x72, 0xe7, 0xa2, 0xba, 0x53, 0x0c, 0x4e, 0xd7, 0x8b,
	0xae, 0x71, 0x98, 0xc5, 0xd2, 0x57, 0x1b, 0x59, 0xec, 0x3d, 0x23, 0x8f, 0xfe, 0x67, 0x57, 0xe0,
	0xaf, 0x6a, 0xcb, 0x56, 0x91, 0xe0, 0x35, 0x42, 0x07, 0xeb, 0xc4, 0x8f, 0x75, 0xa0, 0x13, 0x6d,
	0x85, 0xd1, 0xf8, 0x66, 0x2b, 0x8c, 0xe6, 0x13, 0x72, 0x20, 0x58, 0xc9, 0x05, 0x95, 0xcc, 0x18,
	0x50, 0x4d, 0xa4, 0x3d, 0x43, 0x72, 0x0f, 0xd7, 0x17, 0xed, 0xb6, 0x73, 0x8e, 0x63, 0x28, 0x7c,
	0x3b, 0x8e, 0xa1, 0xed, 0x94, 0x64, 0x2d, 0x36, 0x72, 0xe5, 0x3b, 0xe4, 0xf6, 0x71, 0xbf, 0x71,
	0xe3, 0x16, 0x10, 0x85, 0xef, 0xb7, 0x80, 0xfd, 0xfb, 0x2c, 0x17, 0x9c, 0x39, 0x1e, 0x09, 0x3f,
	0x44, 0x1f, 0x5a, 0xb4, 0xdb, 0xcd, 0xfb, 0x7a, 0x18, 0xea, 0x3e, 0x8e, 0x63, 0x68, 0x3b, 0x26,
	0x0f, 0x95, 0xf6, 0x70, 0x09, 0x15, 0xf3, 0xa0, 0x55, 0x24, 0xfb, 0xd4, 0x52, 0xf1, 0xb6, 0x93,
	0x8d, 0x52, 0xe8, 0xfa, 0x3c, 0x4e, 0xa1, 0x2b, 0x27, 0x0f, 0x7e, 0x71, 0x6e, 0x98, 0x80, 0x65,
	0x7c, 0xd5, 0x17, 0x44, 0x76, 0xbb, 0x55, 0x67, 0xd9, 0xcc, 0xa3, 0xe2, 0xeb, 0x48, 0x3e, 0xf4,
	0x97, 0xb7, 0xf0, 0x37, 0xe7, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x03, 0x67, 0x41,
	0x8c, 0x04, 0x00, 0x00,
}
