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
// source: subscriber_session_summary.proto

package cisco_ios_xr_iedge4710_oper_subscriber_session_nodes_node_ipv4_address_vrf_summaries_ipv4_address_vrf_summary

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

type SubscriberSessionSummary_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriberSessionSummary_KEYS) Reset()         { *m = SubscriberSessionSummary_KEYS{} }
func (m *SubscriberSessionSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionSummary_KEYS) ProtoMessage()    {}
func (*SubscriberSessionSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{0}
}

func (m *SubscriberSessionSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionSummary_KEYS.Unmarshal(m, b)
}
func (m *SubscriberSessionSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionSummary_KEYS.Merge(m, src)
}
func (m *SubscriberSessionSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionSummary_KEYS.Size(m)
}
func (m *SubscriberSessionSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionSummary_KEYS proto.InternalMessageInfo

func (m *SubscriberSessionSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *SubscriberSessionSummary_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SubscriberSessionSummary_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type SubscriberSessionStateCounts struct {
	InitializedSessions   uint32   `protobuf:"varint,1,opt,name=initialized_sessions,json=initializedSessions,proto3" json:"initialized_sessions,omitempty"`
	ConnectingSessions    uint32   `protobuf:"varint,2,opt,name=connecting_sessions,json=connectingSessions,proto3" json:"connecting_sessions,omitempty"`
	ConnectedSessions     uint32   `protobuf:"varint,3,opt,name=connected_sessions,json=connectedSessions,proto3" json:"connected_sessions,omitempty"`
	ActivatedSessions     uint32   `protobuf:"varint,4,opt,name=activated_sessions,json=activatedSessions,proto3" json:"activated_sessions,omitempty"`
	IdleSessions          uint32   `protobuf:"varint,5,opt,name=idle_sessions,json=idleSessions,proto3" json:"idle_sessions,omitempty"`
	DisconnectingSessions uint32   `protobuf:"varint,6,opt,name=disconnecting_sessions,json=disconnectingSessions,proto3" json:"disconnecting_sessions,omitempty"`
	EndSessions           uint32   `protobuf:"varint,7,opt,name=end_sessions,json=endSessions,proto3" json:"end_sessions,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *SubscriberSessionStateCounts) Reset()         { *m = SubscriberSessionStateCounts{} }
func (m *SubscriberSessionStateCounts) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionStateCounts) ProtoMessage()    {}
func (*SubscriberSessionStateCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{1}
}

func (m *SubscriberSessionStateCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionStateCounts.Unmarshal(m, b)
}
func (m *SubscriberSessionStateCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionStateCounts.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionStateCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionStateCounts.Merge(m, src)
}
func (m *SubscriberSessionStateCounts) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionStateCounts.Size(m)
}
func (m *SubscriberSessionStateCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionStateCounts.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionStateCounts proto.InternalMessageInfo

func (m *SubscriberSessionStateCounts) GetInitializedSessions() uint32 {
	if m != nil {
		return m.InitializedSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetConnectingSessions() uint32 {
	if m != nil {
		return m.ConnectingSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetConnectedSessions() uint32 {
	if m != nil {
		return m.ConnectedSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetActivatedSessions() uint32 {
	if m != nil {
		return m.ActivatedSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetIdleSessions() uint32 {
	if m != nil {
		return m.IdleSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetDisconnectingSessions() uint32 {
	if m != nil {
		return m.DisconnectingSessions
	}
	return 0
}

func (m *SubscriberSessionStateCounts) GetEndSessions() uint32 {
	if m != nil {
		return m.EndSessions
	}
	return 0
}

type SubscriberSessionStateSummary struct {
	Pppoe                *SubscriberSessionStateCounts `protobuf:"bytes,1,opt,name=pppoe,proto3" json:"pppoe,omitempty"`
	IpSubscriberDhcp     *SubscriberSessionStateCounts `protobuf:"bytes,2,opt,name=ip_subscriber_dhcp,json=ipSubscriberDhcp,proto3" json:"ip_subscriber_dhcp,omitempty"`
	IpSubscriberPacket   *SubscriberSessionStateCounts `protobuf:"bytes,3,opt,name=ip_subscriber_packet,json=ipSubscriberPacket,proto3" json:"ip_subscriber_packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SubscriberSessionStateSummary) Reset()         { *m = SubscriberSessionStateSummary{} }
func (m *SubscriberSessionStateSummary) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionStateSummary) ProtoMessage()    {}
func (*SubscriberSessionStateSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{2}
}

func (m *SubscriberSessionStateSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionStateSummary.Unmarshal(m, b)
}
func (m *SubscriberSessionStateSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionStateSummary.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionStateSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionStateSummary.Merge(m, src)
}
func (m *SubscriberSessionStateSummary) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionStateSummary.Size(m)
}
func (m *SubscriberSessionStateSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionStateSummary.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionStateSummary proto.InternalMessageInfo

func (m *SubscriberSessionStateSummary) GetPppoe() *SubscriberSessionStateCounts {
	if m != nil {
		return m.Pppoe
	}
	return nil
}

func (m *SubscriberSessionStateSummary) GetIpSubscriberDhcp() *SubscriberSessionStateCounts {
	if m != nil {
		return m.IpSubscriberDhcp
	}
	return nil
}

func (m *SubscriberSessionStateSummary) GetIpSubscriberPacket() *SubscriberSessionStateCounts {
	if m != nil {
		return m.IpSubscriberPacket
	}
	return nil
}

type SubscriberSessionAfCounts struct {
	InProgressSessions   uint32   `protobuf:"varint,1,opt,name=in_progress_sessions,json=inProgressSessions,proto3" json:"in_progress_sessions,omitempty"`
	Ipv4OnlySessions     uint32   `protobuf:"varint,2,opt,name=ipv4_only_sessions,json=ipv4OnlySessions,proto3" json:"ipv4_only_sessions,omitempty"`
	Ipv6OnlySessions     uint32   `protobuf:"varint,3,opt,name=ipv6_only_sessions,json=ipv6OnlySessions,proto3" json:"ipv6_only_sessions,omitempty"`
	DualPartUpSessions   uint32   `protobuf:"varint,4,opt,name=dual_part_up_sessions,json=dualPartUpSessions,proto3" json:"dual_part_up_sessions,omitempty"`
	DualUpSessions       uint32   `protobuf:"varint,5,opt,name=dual_up_sessions,json=dualUpSessions,proto3" json:"dual_up_sessions,omitempty"`
	LacSessions          uint32   `protobuf:"varint,6,opt,name=lac_sessions,json=lacSessions,proto3" json:"lac_sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriberSessionAfCounts) Reset()         { *m = SubscriberSessionAfCounts{} }
func (m *SubscriberSessionAfCounts) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionAfCounts) ProtoMessage()    {}
func (*SubscriberSessionAfCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{3}
}

func (m *SubscriberSessionAfCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionAfCounts.Unmarshal(m, b)
}
func (m *SubscriberSessionAfCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionAfCounts.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionAfCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionAfCounts.Merge(m, src)
}
func (m *SubscriberSessionAfCounts) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionAfCounts.Size(m)
}
func (m *SubscriberSessionAfCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionAfCounts.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionAfCounts proto.InternalMessageInfo

func (m *SubscriberSessionAfCounts) GetInProgressSessions() uint32 {
	if m != nil {
		return m.InProgressSessions
	}
	return 0
}

func (m *SubscriberSessionAfCounts) GetIpv4OnlySessions() uint32 {
	if m != nil {
		return m.Ipv4OnlySessions
	}
	return 0
}

func (m *SubscriberSessionAfCounts) GetIpv6OnlySessions() uint32 {
	if m != nil {
		return m.Ipv6OnlySessions
	}
	return 0
}

func (m *SubscriberSessionAfCounts) GetDualPartUpSessions() uint32 {
	if m != nil {
		return m.DualPartUpSessions
	}
	return 0
}

func (m *SubscriberSessionAfCounts) GetDualUpSessions() uint32 {
	if m != nil {
		return m.DualUpSessions
	}
	return 0
}

func (m *SubscriberSessionAfCounts) GetLacSessions() uint32 {
	if m != nil {
		return m.LacSessions
	}
	return 0
}

type SubscriberSessionAfSummary struct {
	Pppoe                *SubscriberSessionAfCounts `protobuf:"bytes,1,opt,name=pppoe,proto3" json:"pppoe,omitempty"`
	IpSubscriberDhcp     *SubscriberSessionAfCounts `protobuf:"bytes,2,opt,name=ip_subscriber_dhcp,json=ipSubscriberDhcp,proto3" json:"ip_subscriber_dhcp,omitempty"`
	IpSubscriberPacket   *SubscriberSessionAfCounts `protobuf:"bytes,3,opt,name=ip_subscriber_packet,json=ipSubscriberPacket,proto3" json:"ip_subscriber_packet,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SubscriberSessionAfSummary) Reset()         { *m = SubscriberSessionAfSummary{} }
func (m *SubscriberSessionAfSummary) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionAfSummary) ProtoMessage()    {}
func (*SubscriberSessionAfSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{4}
}

func (m *SubscriberSessionAfSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionAfSummary.Unmarshal(m, b)
}
func (m *SubscriberSessionAfSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionAfSummary.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionAfSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionAfSummary.Merge(m, src)
}
func (m *SubscriberSessionAfSummary) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionAfSummary.Size(m)
}
func (m *SubscriberSessionAfSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionAfSummary.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionAfSummary proto.InternalMessageInfo

func (m *SubscriberSessionAfSummary) GetPppoe() *SubscriberSessionAfCounts {
	if m != nil {
		return m.Pppoe
	}
	return nil
}

func (m *SubscriberSessionAfSummary) GetIpSubscriberDhcp() *SubscriberSessionAfCounts {
	if m != nil {
		return m.IpSubscriberDhcp
	}
	return nil
}

func (m *SubscriberSessionAfSummary) GetIpSubscriberPacket() *SubscriberSessionAfCounts {
	if m != nil {
		return m.IpSubscriberPacket
	}
	return nil
}

type SubscriberSessionSummary struct {
	StateXr              *SubscriberSessionStateSummary `protobuf:"bytes,50,opt,name=state_xr,json=stateXr,proto3" json:"state_xr,omitempty"`
	AddressFamilyXr      *SubscriberSessionAfSummary    `protobuf:"bytes,51,opt,name=address_family_xr,json=addressFamilyXr,proto3" json:"address_family_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SubscriberSessionSummary) Reset()         { *m = SubscriberSessionSummary{} }
func (m *SubscriberSessionSummary) String() string { return proto.CompactTextString(m) }
func (*SubscriberSessionSummary) ProtoMessage()    {}
func (*SubscriberSessionSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_1af70e55c78a4792, []int{5}
}

func (m *SubscriberSessionSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberSessionSummary.Unmarshal(m, b)
}
func (m *SubscriberSessionSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberSessionSummary.Marshal(b, m, deterministic)
}
func (m *SubscriberSessionSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberSessionSummary.Merge(m, src)
}
func (m *SubscriberSessionSummary) XXX_Size() int {
	return xxx_messageInfo_SubscriberSessionSummary.Size(m)
}
func (m *SubscriberSessionSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberSessionSummary.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberSessionSummary proto.InternalMessageInfo

func (m *SubscriberSessionSummary) GetStateXr() *SubscriberSessionStateSummary {
	if m != nil {
		return m.StateXr
	}
	return nil
}

func (m *SubscriberSessionSummary) GetAddressFamilyXr() *SubscriberSessionAfSummary {
	if m != nil {
		return m.AddressFamilyXr
	}
	return nil
}

func init() {
	proto.RegisterType((*SubscriberSessionSummary_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_summary_KEYS")
	proto.RegisterType((*SubscriberSessionStateCounts)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_state_counts")
	proto.RegisterType((*SubscriberSessionStateSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_state_summary")
	proto.RegisterType((*SubscriberSessionAfCounts)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_af_counts")
	proto.RegisterType((*SubscriberSessionAfSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_af_summary")
	proto.RegisterType((*SubscriberSessionSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.ipv4_address_vrf_summaries.ipv4_address_vrf_summary.subscriber_session_summary")
}

func init() { proto.RegisterFile("subscriber_session_summary.proto", fileDescriptor_1af70e55c78a4792) }

var fileDescriptor_1af70e55c78a4792 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xdf, 0x6e, 0xd3, 0x3c,
	0x18, 0xc6, 0x95, 0xb6, 0x5b, 0x3b, 0xb7, 0xfb, 0xbe, 0xce, 0xdb, 0x50, 0xf9, 0x27, 0xb6, 0x72,
	0xb2, 0x03, 0x08, 0xeb, 0x56, 0xc6, 0x0d, 0x00, 0x27, 0x48, 0x50, 0xb5, 0x42, 0x1a, 0x47, 0x96,
	0xeb, 0xb8, 0x9d, 0x45, 0x62, 0x5b, 0x76, 0x5a, 0xad, 0x5c, 0x00, 0x27, 0x20, 0xb8, 0x07, 0x24,
	0x8e, 0xe0, 0x06, 0x40, 0x02, 0x09, 0x2e, 0x80, 0x8b, 0xe0, 0x4a, 0x50, 0x9c, 0xc4, 0x4d, 0xb3,
	0x76, 0x9c, 0xa1, 0x9c, 0x54, 0x8a, 0x9f, 0xdf, 0x6b, 0x3f, 0xce, 0xfb, 0x3e, 0x6d, 0xc1, 0x9e,
	0x9e, 0x0c, 0x35, 0x51, 0x6c, 0x48, 0x15, 0xd2, 0x54, 0x6b, 0x26, 0x38, 0xd2, 0x93, 0x20, 0xc0,
	0x6a, 0xe6, 0x4a, 0x25, 0x42, 0x01, 0x03, 0xc2, 0x34, 0x11, 0x88, 0x09, 0x8d, 0xce, 0x15, 0x62,
	0xd4, 0x1b, 0xd3, 0xee, 0x83, 0xce, 0x21, 0x12, 0x92, 0x2a, 0x77, 0x5e, 0xed, 0x26, 0xd5, 0x2e,
	0x17, 0x1e, 0xd5, 0xe6, 0xd3, 0x65, 0x72, 0xda, 0x45, 0xd8, 0xf3, 0x14, 0xd5, 0x1a, 0x4d, 0xd5,
	0x28, 0xd9, 0x99, 0x51, 0xbd, 0x4a, 0x9a, 0xb5, 0x35, 0xb8, 0xb5, 0xda, 0x12, 0x7a, 0xf2, 0xe8,
	0xc5, 0x00, 0x5e, 0x07, 0x1b, 0xd1, 0xf6, 0x88, 0xe3, 0x80, 0xb6, 0x9c, 0x3d, 0xe7, 0x60, 0xa3,
	0x5f, 0x8b, 0x16, 0x9e, 0xe2, 0x80, 0xc2, 0xab, 0xa0, 0x16, 0x6d, 0x67, 0xb4, 0x92, 0xd1, 0xaa,
	0x53, 0x35, 0x32, 0x52, 0x0b, 0x54, 0x93, 0x13, 0x5b, 0xe5, 0x58, 0x49, 0x1e, 0xdb, 0xbf, 0x4b,
	0xcb, 0x4f, 0x0d, 0x71, 0x48, 0x11, 0x11, 0x13, 0x1e, 0x6a, 0xd8, 0x01, 0x3b, 0x8c, 0xb3, 0x90,
	0x61, 0x9f, 0xbd, 0xa2, 0x5e, 0xca, 0x68, 0x63, 0x60, 0xb3, 0xbf, 0x9d, 0xd1, 0x06, 0x89, 0x04,
	0xef, 0x81, 0x6d, 0x22, 0x38, 0xa7, 0x24, 0x64, 0x7c, 0x3c, 0xaf, 0x28, 0x99, 0x0a, 0x38, 0x97,
	0x6c, 0xc1, 0x5d, 0x90, 0xae, 0x66, 0x4f, 0x28, 0x1b, 0x7e, 0xcb, 0x2a, 0x59, 0x1c, 0x93, 0x90,
	0x4d, 0xf1, 0x02, 0x5e, 0x89, 0x71, 0xab, 0x58, 0xfc, 0x36, 0xd8, 0x64, 0x9e, 0x4f, 0xe7, 0xe4,
	0x9a, 0x21, 0x1b, 0xd1, 0xa2, 0x85, 0xee, 0x83, 0x2b, 0x5e, 0xd4, 0xf0, 0x8b, 0xb6, 0xd7, 0x0d,
	0xbd, 0xbb, 0xa0, 0xda, 0xb2, 0x7d, 0xd0, 0xa0, 0x3c, 0x63, 0xa2, 0x6a, 0xe0, 0x3a, 0xe5, 0xf6,
	0xf8, 0xf6, 0xaf, 0xca, 0xf2, 0x69, 0x33, 0x2f, 0x39, 0x69, 0x30, 0xfc, 0xe8, 0x80, 0x35, 0x29,
	0xa5, 0x88, 0x1b, 0x5b, 0x3f, 0x7a, 0xe7, 0xb8, 0xff, 0x74, 0xfe, 0xdc, 0xbf, 0x8c, 0x41, 0x3f,
	0x76, 0x07, 0xbf, 0x39, 0x00, 0x32, 0x89, 0x32, 0xb4, 0x77, 0x46, 0xa4, 0x69, 0x6d, 0x01, 0x4d,
	0x37, 0x99, 0x1c, 0x58, 0xe4, 0xe1, 0x19, 0x91, 0xf0, 0x87, 0x03, 0x76, 0x16, 0xfd, 0x4b, 0x4c,
	0x5e, 0xd2, 0xd0, 0x0c, 0x5b, 0x01, 0x6f, 0x00, 0xb3, 0x37, 0xe8, 0x19, 0xab, 0xed, 0xcf, 0x25,
	0x70, 0x63, 0x49, 0x1d, 0x1e, 0xa5, 0x91, 0x3d, 0x8c, 0x22, 0x8b, 0xa4, 0x12, 0x63, 0x73, 0x5e,
	0x2e, 0xb2, 0x90, 0xf1, 0x5e, 0x22, 0xd9, 0x31, 0xbe, 0x13, 0x75, 0x75, 0xda, 0x45, 0x82, 0xfb,
	0xb3, 0x7c, 0x60, 0x9b, 0x91, 0xf2, 0x8c, 0xfb, 0xb3, 0x1c, 0x7d, 0x92, 0xa3, 0xcb, 0x96, 0x3e,
	0x59, 0xa0, 0x3b, 0x60, 0xd7, 0x9b, 0x60, 0x1f, 0x49, 0xac, 0x42, 0x34, 0x91, 0xf9, 0xc0, 0xc2,
	0x48, 0xec, 0x61, 0x15, 0x3e, 0x97, 0xb6, 0xe4, 0x00, 0x34, 0x4d, 0x49, 0x96, 0x8e, 0x43, 0xfb,
	0x5f, 0xb4, 0x9e, 0x21, 0xf7, 0x41, 0xc3, 0xc7, 0x24, 0x1f, 0xd6, 0xba, 0x8f, 0x89, 0xcd, 0xdf,
	0xcf, 0x0a, 0xb8, 0xb9, 0xfc, 0x75, 0xa5, 0xe1, 0xfb, 0x90, 0x0b, 0xdf, 0x9b, 0x02, 0x4c, 0x81,
	0xed, 0x66, 0x9a, 0xbc, 0xaf, 0x97, 0x25, 0xaf, 0x58, 0x8e, 0x2f, 0xc6, 0xee, 0xfb, 0xe5, 0xb1,
	0x2b, 0x96, 0xfd, 0x65, 0x99, 0x7b, 0x5d, 0x06, 0xd7, 0x56, 0xff, 0x3e, 0xc3, 0x4f, 0x0e, 0xa8,
	0xc5, 0xb9, 0x3d, 0x57, 0xad, 0x23, 0x73, 0xa7, 0xf7, 0x85, 0xf9, 0x2a, 0x49, 0x80, 0x7e, 0xd5,
	0x3c, 0x9e, 0x2a, 0xf8, 0xc5, 0x01, 0x5b, 0xe9, 0x0e, 0x23, 0x1c, 0x30, 0x7f, 0x16, 0xd9, 0x3e,
	0x36, 0xb6, 0xdf, 0x16, 0xa3, 0x15, 0xa9, 0xe7, 0xff, 0x93, 0x92, 0xc7, 0xc6, 0xe6, 0xa9, 0x1a,
	0xae, 0x9b, 0x7f, 0x67, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x45, 0x3d, 0x6d, 0xc1,
	0x09, 0x00, 0x00,
}
