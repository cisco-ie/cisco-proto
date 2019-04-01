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
// source: bfd_mgmt_session_brief.proto

package cisco_ios_xr_ip_bfd_oper_bfd_ipv4bf_do_mplste_tail_session_briefs_ipv4bf_do_mplste_tail_session_brief

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

type BfdMgmtSessionBrief_KEYS struct {
	VrfName               string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	IncomingLabel         uint32   `protobuf:"varint,2,opt,name=incoming_label,json=incomingLabel,proto3" json:"incoming_label,omitempty"`
	FeCtype               uint32   `protobuf:"varint,3,opt,name=fe_ctype,json=feCtype,proto3" json:"fe_ctype,omitempty"`
	FecSubgroupId         uint32   `protobuf:"varint,4,opt,name=fec_subgroup_id,json=fecSubgroupId,proto3" json:"fec_subgroup_id,omitempty"`
	Feclspid              uint32   `protobuf:"varint,5,opt,name=feclspid,proto3" json:"feclspid,omitempty"`
	FecTunnelId           uint32   `protobuf:"varint,6,opt,name=fec_tunnel_id,json=fecTunnelId,proto3" json:"fec_tunnel_id,omitempty"`
	FecExtendedTunnelId   string   `protobuf:"bytes,7,opt,name=fec_extended_tunnel_id,json=fecExtendedTunnelId,proto3" json:"fec_extended_tunnel_id,omitempty"`
	FecSource             string   `protobuf:"bytes,8,opt,name=fec_source,json=fecSource,proto3" json:"fec_source,omitempty"`
	FecDestination        string   `protobuf:"bytes,9,opt,name=fec_destination,json=fecDestination,proto3" json:"fec_destination,omitempty"`
	Fecp2Mpid             uint32   `protobuf:"varint,10,opt,name=fecp2mpid,proto3" json:"fecp2mpid,omitempty"`
	FecSubgroupOriginator string   `protobuf:"bytes,11,opt,name=fec_subgroup_originator,json=fecSubgroupOriginator,proto3" json:"fec_subgroup_originator,omitempty"`
	FecCtype              uint32   `protobuf:"varint,12,opt,name=fec_ctype,json=fecCtype,proto3" json:"fec_ctype,omitempty"`
	Location              string   `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *BfdMgmtSessionBrief_KEYS) Reset()         { *m = BfdMgmtSessionBrief_KEYS{} }
func (m *BfdMgmtSessionBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSessionBrief_KEYS) ProtoMessage()    {}
func (*BfdMgmtSessionBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3591794e8a2d9099, []int{0}
}

func (m *BfdMgmtSessionBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSessionBrief_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtSessionBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSessionBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSessionBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSessionBrief_KEYS.Merge(m, src)
}
func (m *BfdMgmtSessionBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSessionBrief_KEYS.Size(m)
}
func (m *BfdMgmtSessionBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSessionBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSessionBrief_KEYS proto.InternalMessageInfo

func (m *BfdMgmtSessionBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetIncomingLabel() uint32 {
	if m != nil {
		return m.IncomingLabel
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFeCtype() uint32 {
	if m != nil {
		return m.FeCtype
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecSubgroupId() uint32 {
	if m != nil {
		return m.FecSubgroupId
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFeclspid() uint32 {
	if m != nil {
		return m.Feclspid
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecTunnelId() uint32 {
	if m != nil {
		return m.FecTunnelId
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecExtendedTunnelId() string {
	if m != nil {
		return m.FecExtendedTunnelId
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecSource() string {
	if m != nil {
		return m.FecSource
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecDestination() string {
	if m != nil {
		return m.FecDestination
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecp2Mpid() uint32 {
	if m != nil {
		return m.Fecp2Mpid
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecSubgroupOriginator() string {
	if m != nil {
		return m.FecSubgroupOriginator
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetFecCtype() uint32 {
	if m != nil {
		return m.FecCtype
	}
	return 0
}

func (m *BfdMgmtSessionBrief_KEYS) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type BfdMgmtAsyncIntervalMultiplierInformation struct {
	NegotiatedRemoteTransmitInterval uint32   `protobuf:"varint,1,opt,name=negotiated_remote_transmit_interval,json=negotiatedRemoteTransmitInterval,proto3" json:"negotiated_remote_transmit_interval,omitempty"`
	NegotiatedLocalTransmitInterval  uint32   `protobuf:"varint,2,opt,name=negotiated_local_transmit_interval,json=negotiatedLocalTransmitInterval,proto3" json:"negotiated_local_transmit_interval,omitempty"`
	DetectionTime                    uint32   `protobuf:"varint,3,opt,name=detection_time,json=detectionTime,proto3" json:"detection_time,omitempty"`
	DetectionMultiplier              uint32   `protobuf:"varint,4,opt,name=detection_multiplier,json=detectionMultiplier,proto3" json:"detection_multiplier,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *BfdMgmtAsyncIntervalMultiplierInformation) Reset() {
	*m = BfdMgmtAsyncIntervalMultiplierInformation{}
}
func (m *BfdMgmtAsyncIntervalMultiplierInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtAsyncIntervalMultiplierInformation) ProtoMessage()    {}
func (*BfdMgmtAsyncIntervalMultiplierInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3591794e8a2d9099, []int{1}
}

func (m *BfdMgmtAsyncIntervalMultiplierInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation.Unmarshal(m, b)
}
func (m *BfdMgmtAsyncIntervalMultiplierInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtAsyncIntervalMultiplierInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation.Merge(m, src)
}
func (m *BfdMgmtAsyncIntervalMultiplierInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation.Size(m)
}
func (m *BfdMgmtAsyncIntervalMultiplierInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtAsyncIntervalMultiplierInformation proto.InternalMessageInfo

func (m *BfdMgmtAsyncIntervalMultiplierInformation) GetNegotiatedRemoteTransmitInterval() uint32 {
	if m != nil {
		return m.NegotiatedRemoteTransmitInterval
	}
	return 0
}

func (m *BfdMgmtAsyncIntervalMultiplierInformation) GetNegotiatedLocalTransmitInterval() uint32 {
	if m != nil {
		return m.NegotiatedLocalTransmitInterval
	}
	return 0
}

func (m *BfdMgmtAsyncIntervalMultiplierInformation) GetDetectionTime() uint32 {
	if m != nil {
		return m.DetectionTime
	}
	return 0
}

func (m *BfdMgmtAsyncIntervalMultiplierInformation) GetDetectionMultiplier() uint32 {
	if m != nil {
		return m.DetectionMultiplier
	}
	return 0
}

type BfdMgmtEchoIntervalMultiplierInformation struct {
	NegotiatedTransmitInterval uint32   `protobuf:"varint,1,opt,name=negotiated_transmit_interval,json=negotiatedTransmitInterval,proto3" json:"negotiated_transmit_interval,omitempty"`
	DetectionTime              uint32   `protobuf:"varint,2,opt,name=detection_time,json=detectionTime,proto3" json:"detection_time,omitempty"`
	DetectionMultiplier        uint32   `protobuf:"varint,3,opt,name=detection_multiplier,json=detectionMultiplier,proto3" json:"detection_multiplier,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *BfdMgmtEchoIntervalMultiplierInformation) Reset() {
	*m = BfdMgmtEchoIntervalMultiplierInformation{}
}
func (m *BfdMgmtEchoIntervalMultiplierInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtEchoIntervalMultiplierInformation) ProtoMessage()    {}
func (*BfdMgmtEchoIntervalMultiplierInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3591794e8a2d9099, []int{2}
}

func (m *BfdMgmtEchoIntervalMultiplierInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation.Unmarshal(m, b)
}
func (m *BfdMgmtEchoIntervalMultiplierInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtEchoIntervalMultiplierInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation.Merge(m, src)
}
func (m *BfdMgmtEchoIntervalMultiplierInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation.Size(m)
}
func (m *BfdMgmtEchoIntervalMultiplierInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtEchoIntervalMultiplierInformation proto.InternalMessageInfo

func (m *BfdMgmtEchoIntervalMultiplierInformation) GetNegotiatedTransmitInterval() uint32 {
	if m != nil {
		return m.NegotiatedTransmitInterval
	}
	return 0
}

func (m *BfdMgmtEchoIntervalMultiplierInformation) GetDetectionTime() uint32 {
	if m != nil {
		return m.DetectionTime
	}
	return 0
}

func (m *BfdMgmtEchoIntervalMultiplierInformation) GetDetectionMultiplier() uint32 {
	if m != nil {
		return m.DetectionMultiplier
	}
	return 0
}

type BfdMgmtSessionStatusInformationBrief struct {
	AsyncIntervalMultiplier *BfdMgmtAsyncIntervalMultiplierInformation `protobuf:"bytes,1,opt,name=async_interval_multiplier,json=asyncIntervalMultiplier,proto3" json:"async_interval_multiplier,omitempty"`
	EchoIntervalMultiplier  *BfdMgmtEchoIntervalMultiplierInformation  `protobuf:"bytes,2,opt,name=echo_interval_multiplier,json=echoIntervalMultiplier,proto3" json:"echo_interval_multiplier,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                   `json:"-"`
	XXX_unrecognized        []byte                                     `json:"-"`
	XXX_sizecache           int32                                      `json:"-"`
}

func (m *BfdMgmtSessionStatusInformationBrief) Reset()         { *m = BfdMgmtSessionStatusInformationBrief{} }
func (m *BfdMgmtSessionStatusInformationBrief) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSessionStatusInformationBrief) ProtoMessage()    {}
func (*BfdMgmtSessionStatusInformationBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_3591794e8a2d9099, []int{3}
}

func (m *BfdMgmtSessionStatusInformationBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSessionStatusInformationBrief.Unmarshal(m, b)
}
func (m *BfdMgmtSessionStatusInformationBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSessionStatusInformationBrief.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSessionStatusInformationBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSessionStatusInformationBrief.Merge(m, src)
}
func (m *BfdMgmtSessionStatusInformationBrief) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSessionStatusInformationBrief.Size(m)
}
func (m *BfdMgmtSessionStatusInformationBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSessionStatusInformationBrief.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSessionStatusInformationBrief proto.InternalMessageInfo

func (m *BfdMgmtSessionStatusInformationBrief) GetAsyncIntervalMultiplier() *BfdMgmtAsyncIntervalMultiplierInformation {
	if m != nil {
		return m.AsyncIntervalMultiplier
	}
	return nil
}

func (m *BfdMgmtSessionStatusInformationBrief) GetEchoIntervalMultiplier() *BfdMgmtEchoIntervalMultiplierInformation {
	if m != nil {
		return m.EchoIntervalMultiplier
	}
	return nil
}

type BfdMgmtSessionBrief struct {
	NodeId                 string                                `protobuf:"bytes,50,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	State                  string                                `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	StatusBriefInformation *BfdMgmtSessionStatusInformationBrief `protobuf:"bytes,52,opt,name=status_brief_information,json=statusBriefInformation,proto3" json:"status_brief_information,omitempty"`
	SessionType            string                                `protobuf:"bytes,53,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	SessionSubtype         string                                `protobuf:"bytes,54,opt,name=session_subtype,json=sessionSubtype,proto3" json:"session_subtype,omitempty"`
	SessionFlags           uint32                                `protobuf:"varint,55,opt,name=session_flags,json=sessionFlags,proto3" json:"session_flags,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                              `json:"-"`
	XXX_unrecognized       []byte                                `json:"-"`
	XXX_sizecache          int32                                 `json:"-"`
}

func (m *BfdMgmtSessionBrief) Reset()         { *m = BfdMgmtSessionBrief{} }
func (m *BfdMgmtSessionBrief) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSessionBrief) ProtoMessage()    {}
func (*BfdMgmtSessionBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_3591794e8a2d9099, []int{4}
}

func (m *BfdMgmtSessionBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSessionBrief.Unmarshal(m, b)
}
func (m *BfdMgmtSessionBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSessionBrief.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSessionBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSessionBrief.Merge(m, src)
}
func (m *BfdMgmtSessionBrief) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSessionBrief.Size(m)
}
func (m *BfdMgmtSessionBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSessionBrief.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSessionBrief proto.InternalMessageInfo

func (m *BfdMgmtSessionBrief) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BfdMgmtSessionBrief) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BfdMgmtSessionBrief) GetStatusBriefInformation() *BfdMgmtSessionStatusInformationBrief {
	if m != nil {
		return m.StatusBriefInformation
	}
	return nil
}

func (m *BfdMgmtSessionBrief) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *BfdMgmtSessionBrief) GetSessionSubtype() string {
	if m != nil {
		return m.SessionSubtype
	}
	return ""
}

func (m *BfdMgmtSessionBrief) GetSessionFlags() uint32 {
	if m != nil {
		return m.SessionFlags
	}
	return 0
}

func init() {
	proto.RegisterType((*BfdMgmtSessionBrief_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4bf_do_mplste_tail_session_briefs.ipv4bf_do_mplste_tail_session_brief.bfd_mgmt_session_brief_KEYS")
	proto.RegisterType((*BfdMgmtAsyncIntervalMultiplierInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4bf_do_mplste_tail_session_briefs.ipv4bf_do_mplste_tail_session_brief.bfd_mgmt_async_interval_multiplier_information")
	proto.RegisterType((*BfdMgmtEchoIntervalMultiplierInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4bf_do_mplste_tail_session_briefs.ipv4bf_do_mplste_tail_session_brief.bfd_mgmt_echo_interval_multiplier_information")
	proto.RegisterType((*BfdMgmtSessionStatusInformationBrief)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4bf_do_mplste_tail_session_briefs.ipv4bf_do_mplste_tail_session_brief.bfd_mgmt_session_status_information_brief")
	proto.RegisterType((*BfdMgmtSessionBrief)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4bf_do_mplste_tail_session_briefs.ipv4bf_do_mplste_tail_session_brief.bfd_mgmt_session_brief")
}

func init() { proto.RegisterFile("bfd_mgmt_session_brief.proto", fileDescriptor_3591794e8a2d9099) }

var fileDescriptor_3591794e8a2d9099 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x4e, 0x1b, 0x49,
	0x14, 0x55, 0xe3, 0x01, 0xc3, 0x35, 0x06, 0xa9, 0x60, 0x4c, 0xf3, 0x18, 0x0d, 0x63, 0xc4, 0xc0,
	0x2c, 0xc6, 0xd2, 0x18, 0x86, 0xd9, 0x8e, 0x92, 0x10, 0xc9, 0x02, 0x12, 0xc9, 0x78, 0x93, 0x55,
	0xa9, 0x1f, 0xb7, 0x9c, 0x92, 0xba, 0xab, 0x5a, 0x55, 0x65, 0x0b, 0x7e, 0x20, 0xeb, 0x2c, 0x92,
	0x5f, 0xc8, 0xcf, 0xb0, 0x88, 0xc4, 0x87, 0xe4, 0x1b, 0xa2, 0xaa, 0x6e, 0x77, 0x1b, 0x81, 0x05,
	0x6c, 0x58, 0xde, 0x53, 0xa7, 0x8e, 0xef, 0x39, 0x7d, 0xeb, 0x1a, 0x76, 0x42, 0x16, 0xd3, 0x74,
	0x98, 0x1a, 0xaa, 0x51, 0x6b, 0x2e, 0x05, 0x0d, 0x15, 0x47, 0xd6, 0xc9, 0x94, 0x34, 0x92, 0x60,
	0xc4, 0x75, 0x24, 0x29, 0x97, 0x9a, 0x5e, 0x29, 0xca, 0x33, 0x6a, 0xd9, 0x32, 0x43, 0xd5, 0x09,
	0x59, 0xdc, 0xe1, 0xd9, 0xf8, 0x38, 0x64, 0x34, 0x96, 0x34, 0xcd, 0x12, 0x6d, 0x90, 0x9a, 0x80,
	0x27, 0x77, 0x75, 0xf4, 0x53, 0x48, 0xed, 0x1f, 0x35, 0xd8, 0x7e, 0xb8, 0x0f, 0x7a, 0x76, 0xfa,
	0xe1, 0x92, 0x6c, 0xc2, 0xe2, 0x58, 0x31, 0x2a, 0x82, 0x14, 0x7d, 0x6f, 0xd7, 0x3b, 0x5c, 0xea,
	0xd7, 0xc7, 0x8a, 0xbd, 0x0b, 0x52, 0x24, 0xfb, 0xb0, 0xc2, 0x45, 0x24, 0x53, 0x2e, 0x86, 0x34,
	0x09, 0x42, 0x4c, 0xfc, 0xb9, 0x5d, 0xef, 0xb0, 0xd9, 0x6f, 0x4e, 0xd0, 0x73, 0x0b, 0x5a, 0x05,
	0x86, 0x34, 0x32, 0xd7, 0x19, 0xfa, 0x35, 0x47, 0xa8, 0x33, 0x7c, 0x6d, 0x4b, 0xf2, 0x27, 0xac,
	0x32, 0x8c, 0xa8, 0x1e, 0x85, 0x43, 0x25, 0x47, 0x19, 0xe5, 0xb1, 0xff, 0x4b, 0x2e, 0xc1, 0x30,
	0xba, 0x2c, 0xd0, 0x5e, 0x4c, 0xb6, 0xac, 0x44, 0x94, 0xe8, 0x8c, 0xc7, 0xfe, 0xbc, 0x23, 0x94,
	0x35, 0x69, 0x83, 0x25, 0x53, 0x33, 0x12, 0x02, 0x13, 0xab, 0xb0, 0xe0, 0x08, 0x0d, 0x86, 0xd1,
	0xc0, 0x61, 0xbd, 0x98, 0x1c, 0x41, 0xcb, 0x72, 0xf0, 0xca, 0xa0, 0x88, 0x31, 0x9e, 0x22, 0xd7,
	0x9d, 0xa5, 0x35, 0x86, 0xd1, 0x69, 0x71, 0x58, 0x5e, 0xfa, 0x0d, 0xc0, 0x35, 0x27, 0x47, 0x2a,
	0x42, 0x7f, 0xd1, 0x11, 0x97, 0x6c, 0x5f, 0x0e, 0x20, 0x07, 0x79, 0xef, 0x31, 0x6a, 0xc3, 0x45,
	0x60, 0xb8, 0x14, 0xfe, 0x92, 0xe3, 0xac, 0x30, 0x8c, 0xde, 0x54, 0x28, 0xd9, 0x01, 0x7b, 0x2b,
	0xeb, 0xa6, 0xb6, 0x7b, 0x70, 0xcd, 0x55, 0x00, 0x39, 0x81, 0x8d, 0x3b, 0x11, 0x48, 0xc5, 0x87,
	0xf6, 0xa6, 0x54, 0x7e, 0xc3, 0xc9, 0xfd, 0x3a, 0x15, 0xc5, 0xfb, 0xf2, 0x90, 0x6c, 0x3b, 0xd5,
	0x22, 0xd6, 0xe5, 0x32, 0x93, 0x3c, 0xd7, 0x2d, 0x58, 0x4c, 0x64, 0x94, 0x37, 0xd5, 0x74, 0x2a,
	0x65, 0xdd, 0xfe, 0x36, 0x07, 0x9d, 0xf2, 0x83, 0x07, 0xfa, 0x5a, 0x44, 0x94, 0x0b, 0x83, 0x6a,
	0x1c, 0x24, 0x34, 0x1d, 0x25, 0x86, 0x67, 0x09, 0x47, 0x45, 0xb9, 0x60, 0x52, 0xa5, 0xb9, 0x83,
	0x0b, 0xd8, 0x13, 0x38, 0x94, 0x86, 0x07, 0x06, 0x63, 0xaa, 0x30, 0x95, 0x76, 0x96, 0x54, 0x20,
	0x74, 0xca, 0x4d, 0x79, 0xdb, 0x8d, 0x47, 0xb3, 0xbf, 0x5b, 0x51, 0xfb, 0x8e, 0x39, 0x28, 0x88,
	0xbd, 0x82, 0x47, 0xce, 0xa0, 0x3d, 0x25, 0x67, 0x1b, 0x4b, 0x1e, 0x50, 0xcb, 0x67, 0xe9, 0xf7,
	0x8a, 0x79, 0x6e, 0x89, 0xf7, 0xc4, 0xf6, 0x61, 0x25, 0x46, 0x83, 0x91, 0x6d, 0x94, 0x1a, 0x9e,
	0x4e, 0x66, 0xac, 0x59, 0xa2, 0x03, 0x9e, 0x22, 0xf9, 0x07, 0xd6, 0x2b, 0x5a, 0x65, 0xb3, 0x18,
	0xb7, 0xb5, 0xf2, 0xec, 0xa2, 0x3c, 0x6a, 0xdf, 0x78, 0xf0, 0x77, 0x19, 0x14, 0x46, 0x1f, 0xe5,
	0xa3, 0x39, 0xfd, 0x0f, 0x3b, 0x53, 0xc6, 0x66, 0x05, 0xb4, 0x55, 0x71, 0x9e, 0xe0, 0x66, 0xee,
	0x39, 0x6e, 0x6a, 0xb3, 0xdd, 0xdc, 0xd6, 0xe0, 0xaf, 0x7b, 0xef, 0x5c, 0x9b, 0xc0, 0x8c, 0xf4,
	0xb4, 0x87, 0xfc, 0xe9, 0x93, 0x5b, 0x0f, 0x36, 0x67, 0xce, 0x86, 0xf3, 0xd1, 0xe8, 0x7e, 0xf5,
	0x3a, 0x2f, 0xb2, 0xa2, 0x9e, 0x39, 0xad, 0xfd, 0x0d, 0x47, 0x9b, 0x44, 0x5a, 0x45, 0x40, 0xbe,
	0x7b, 0xe0, 0xcf, 0xfa, 0x8e, 0x2e, 0xe7, 0x46, 0xf7, 0xcb, 0x8b, 0x7b, 0x7a, 0xca, 0x60, 0xf5,
	0x5b, 0x96, 0x75, 0xdf, 0x51, 0xfb, 0x53, 0x0d, 0x5a, 0x0f, 0x2f, 0x6f, 0xb2, 0x01, 0x75, 0x21,
	0x63, 0xb4, 0x3b, 0xae, 0xeb, 0x36, 0xc0, 0x82, 0x2d, 0x7b, 0x31, 0x59, 0x87, 0x79, 0xfb, 0xd9,
	0xd1, 0x3f, 0x72, 0x70, 0x5e, 0x90, 0x1b, 0x0f, 0xfc, 0x62, 0x1a, 0xf2, 0xe5, 0x3f, 0xf5, 0xf3,
	0xfe, 0xb1, 0xcb, 0xe6, 0xf3, 0x8b, 0x67, 0xf3, 0xd8, 0x98, 0xf6, 0x5b, 0xf9, 0xc9, 0x2b, 0x5b,
	0xf4, 0xa6, 0x1e, 0xe2, 0x1f, 0xb0, 0x3c, 0xb9, 0xeb, 0xf6, 0xe3, 0xbf, 0xce, 0x6a, 0xa3, 0xc0,
	0x06, 0x76, 0x45, 0x1e, 0xc0, 0x6a, 0x29, 0x3f, 0x0a, 0x1d, 0xeb, 0x24, 0x5f, 0xdf, 0x05, 0x7c,
	0x99, 0xa3, 0x64, 0x0f, 0x9a, 0x13, 0x22, 0x4b, 0x82, 0xa1, 0xf6, 0xff, 0x73, 0x8f, 0x6c, 0xf2,
	0x03, 0x6f, 0x2d, 0x16, 0x2e, 0xb8, 0xff, 0xec, 0xa3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0xc0, 0xfa, 0xba, 0xd3, 0x07, 0x00, 0x00,
}