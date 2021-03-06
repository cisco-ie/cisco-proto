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

package cisco_ios_xr_ip_bfd_oper_bfd_ipv6_multi_hop_session_briefs_ipv6_multi_hop_session_brief

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
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Location             string   `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	VrfName              string   `protobuf:"bytes,4,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *BfdMgmtSessionBrief_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *BfdMgmtSessionBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*BfdMgmtSessionBrief_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv6_multi_hop_session_briefs.ipv6_multi_hop_session_brief.bfd_mgmt_session_brief_KEYS")
	proto.RegisterType((*BfdMgmtAsyncIntervalMultiplierInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv6_multi_hop_session_briefs.ipv6_multi_hop_session_brief.bfd_mgmt_async_interval_multiplier_information")
	proto.RegisterType((*BfdMgmtEchoIntervalMultiplierInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv6_multi_hop_session_briefs.ipv6_multi_hop_session_brief.bfd_mgmt_echo_interval_multiplier_information")
	proto.RegisterType((*BfdMgmtSessionStatusInformationBrief)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv6_multi_hop_session_briefs.ipv6_multi_hop_session_brief.bfd_mgmt_session_status_information_brief")
	proto.RegisterType((*BfdMgmtSessionBrief)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv6_multi_hop_session_briefs.ipv6_multi_hop_session_brief.bfd_mgmt_session_brief")
}

func init() { proto.RegisterFile("bfd_mgmt_session_brief.proto", fileDescriptor_3591794e8a2d9099) }

var fileDescriptor_3591794e8a2d9099 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xe5, 0x14, 0xda, 0xf2, 0x82, 0x83, 0x74, 0xad, 0x52, 0xb7, 0x54, 0xa2, 0xa4, 0xaa,
	0x80, 0x81, 0x20, 0x52, 0x28, 0x2b, 0x20, 0x81, 0x14, 0x95, 0x32, 0xb8, 0x91, 0x10, 0xd3, 0xc9,
	0xb1, 0x9f, 0xd3, 0x93, 0x6c, 0x9f, 0x75, 0x77, 0x89, 0xc8, 0xca, 0x00, 0xff, 0x09, 0x13, 0x33,
	0x1b, 0x03, 0x12, 0x1b, 0xff, 0x10, 0x23, 0xf2, 0x73, 0xfc, 0x03, 0x25, 0xa1, 0xed, 0x40, 0xc7,
	0xfb, 0xbe, 0xcf, 0xbd, 0xcb, 0xf7, 0x7b, 0xef, 0x1c, 0xd8, 0x1d, 0x86, 0x01, 0x8f, 0x47, 0xb1,
	0xe1, 0x1a, 0xb5, 0x16, 0x32, 0xe1, 0x43, 0x25, 0x30, 0xec, 0xa6, 0x4a, 0x1a, 0xc9, 0xde, 0xf9,
	0x42, 0xfb, 0x92, 0x0b, 0xa9, 0xf9, 0x07, 0xc5, 0x45, 0xca, 0x33, 0x5a, 0xa6, 0xa8, 0xba, 0xc3,
	0x30, 0xe8, 0x8a, 0x74, 0x72, 0xc4, 0xe3, 0x71, 0x64, 0x04, 0x3f, 0x93, 0xe9, 0xdf, 0x0d, 0xf4,
	0x3f, 0xab, 0x9d, 0xaf, 0x16, 0xdc, 0x5e, 0x7c, 0x32, 0x3f, 0x7e, 0xf5, 0xfe, 0x94, 0x1d, 0x40,
	0x4b, 0xcb, 0xb1, 0xf2, 0x91, 0x7b, 0x41, 0xa0, 0x50, 0x6b, 0xc7, 0xda, 0xb3, 0xee, 0xdf, 0x70,
	0xed, 0x5c, 0x7d, 0x91, 0x8b, 0xec, 0x11, 0x6c, 0x04, 0xa8, 0x8d, 0x48, 0x3c, 0x93, 0x35, 0x28,
	0xd8, 0x06, 0xb1, 0xac, 0x56, 0x2a, 0x36, 0xec, 0xc0, 0x7a, 0x24, 0x7d, 0x92, 0x9c, 0x15, 0xa2,
	0xca, 0x35, 0xdb, 0x86, 0xf5, 0x89, 0x0a, 0x79, 0xe2, 0xc5, 0xe8, 0x5c, 0xa3, 0xda, 0xda, 0x44,
	0x85, 0x6f, 0xbd, 0x18, 0x3b, 0x5f, 0x1a, 0xd0, 0x2d, 0x7f, 0xae, 0xa7, 0xa7, 0x89, 0xcf, 0x45,
	0x62, 0x50, 0x4d, 0xbc, 0x28, 0xb7, 0x98, 0x46, 0x02, 0x15, 0x17, 0x49, 0x28, 0x55, 0x9c, 0x77,
	0x3b, 0x81, 0xfd, 0x04, 0x47, 0xd2, 0x08, 0xcf, 0x60, 0xc0, 0x15, 0xc6, 0xd2, 0x20, 0x37, 0xca,
	0x4b, 0x74, 0x2c, 0x4c, 0xb9, 0x9b, 0x6c, 0xd9, 0xee, 0x5e, 0x85, 0xba, 0x44, 0x0e, 0x66, 0x60,
	0x7f, 0xc6, 0xb1, 0x63, 0xe8, 0xd4, 0xda, 0x65, 0xbf, 0x39, 0x5a, 0xd0, 0xad, 0x41, 0xdd, 0xee,
	0x54, 0xe4, 0x9b, 0x0c, 0x9c, 0x6b, 0x76, 0x00, 0xad, 0x00, 0x0d, 0xfa, 0x14, 0x9a, 0x11, 0x31,
	0x52, 0x16, 0xb6, 0x6b, 0x97, 0xea, 0x40, 0xc4, 0xc8, 0x1e, 0xc3, 0x66, 0x85, 0x55, 0x36, 0x29,
	0x1c, 0xdb, 0xdd, 0x28, 0x6b, 0x27, 0x65, 0xa9, 0xf3, 0xcb, 0x82, 0x87, 0x65, 0x50, 0xe8, 0x9f,
	0xc9, 0x73, 0x73, 0x7a, 0x0e, 0xbb, 0x35, 0x63, 0xcb, 0x02, 0xda, 0xa9, 0x98, 0x0b, 0xb8, 0x69,
	0x5c, 0xc6, 0xcd, 0xca, 0x72, 0x37, 0xdf, 0x56, 0xe0, 0xc1, 0xdc, 0x94, 0x6a, 0xe3, 0x99, 0xb1,
	0xae, 0x7b, 0xc8, 0x07, 0x97, 0xfd, 0xb4, 0x60, 0x7b, 0xe9, 0x6c, 0x90, 0x8f, 0x66, 0xef, 0xb3,
	0xd5, 0xfd, 0x4f, 0x4f, 0xea, 0x92, 0xf3, 0xe9, 0x6e, 0x11, 0x56, 0x84, 0x58, 0x99, 0x66, 0x3f,
	0x2c, 0x70, 0x96, 0xdd, 0x1c, 0x25, 0xdb, 0xec, 0x7d, 0xba, 0x02, 0x17, 0x17, 0x19, 0x1e, 0xb7,
	0x9d, 0x51, 0xf3, 0x1e, 0x3a, 0xbf, 0x1b, 0xd0, 0x5e, 0xfc, 0x79, 0x61, 0x5b, 0xb0, 0x96, 0xc8,
	0x00, 0xb9, 0x08, 0x9c, 0x1e, 0x3d, 0xf2, 0xd5, 0x6c, 0xd9, 0x0f, 0xd8, 0x26, 0x5c, 0xcf, 0xae,
	0x16, 0x9d, 0x43, 0x92, 0xf3, 0x05, 0xfb, 0x6e, 0x81, 0x33, 0xbb, 0xf1, 0xfc, 0xf3, 0x54, 0x3b,
	0xde, 0x79, 0x42, 0x69, 0x7c, 0xbc, 0x82, 0x34, 0xce, 0x1b, 0x3e, 0xb7, 0x9d, 0x57, 0x5e, 0x66,
	0x8b, 0x7e, 0xed, 0x79, 0xdd, 0x85, 0x9b, 0xc5, 0x5e, 0x33, 0x4d, 0xd1, 0x79, 0x4a, 0xe6, 0x9a,
	0x33, 0x6d, 0x30, 0x4d, 0x91, 0xdd, 0x83, 0x5b, 0x65, 0xfb, 0xf1, 0x90, 0xa8, 0x23, 0xa2, 0x5a,
	0x33, 0xf9, 0x34, 0x57, 0xd9, 0x3e, 0xd8, 0x05, 0x18, 0x46, 0xde, 0x48, 0x3b, 0xcf, 0xe8, 0xe9,
	0x14, 0x07, 0xbc, 0xce, 0xb4, 0xe1, 0x2a, 0xfd, 0x73, 0x1c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x03, 0x05, 0x82, 0xdd, 0x59, 0x06, 0x00, 0x00,
}
