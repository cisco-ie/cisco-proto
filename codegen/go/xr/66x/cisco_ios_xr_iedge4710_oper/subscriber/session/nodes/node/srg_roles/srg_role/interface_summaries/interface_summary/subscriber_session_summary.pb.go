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

package cisco_ios_xr_iedge4710_oper_subscriber_session_nodes_node_srg_roles_srg_role_interface_summaries_interface_summary

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
	Srg                  string   `protobuf:"bytes,2,opt,name=srg,proto3" json:"srg,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *SubscriberSessionSummary_KEYS) GetSrg() string {
	if m != nil {
		return m.Srg
	}
	return ""
}

func (m *SubscriberSessionSummary_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	proto.RegisterType((*SubscriberSessionSummary_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_summary_KEYS")
	proto.RegisterType((*SubscriberSessionStateCounts)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_state_counts")
	proto.RegisterType((*SubscriberSessionStateSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_state_summary")
	proto.RegisterType((*SubscriberSessionAfCounts)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_af_counts")
	proto.RegisterType((*SubscriberSessionAfSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_af_summary")
	proto.RegisterType((*SubscriberSessionSummary)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.session.nodes.node.srg_roles.srg_role.interface_summaries.interface_summary.subscriber_session_summary")
}

func init() { proto.RegisterFile("subscriber_session_summary.proto", fileDescriptor_1af70e55c78a4792) }

var fileDescriptor_1af70e55c78a4792 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x24, 0xfd, 0x75, 0x93, 0x7e, 0xa9, 0xdb, 0x7e, 0x8a, 0xf8, 0x11, 0x6d, 0x10, 0x52,
	0x17, 0x30, 0x34, 0x6d, 0x29, 0x2f, 0x00, 0x6c, 0x90, 0x20, 0x4a, 0x84, 0x54, 0x56, 0x96, 0xe3,
	0x71, 0x52, 0x8b, 0x89, 0x6d, 0xd9, 0x4e, 0x69, 0x78, 0x8a, 0x22, 0x78, 0x07, 0x24, 0xc4, 0x82,
	0x27, 0x00, 0x09, 0x24, 0xf6, 0x3c, 0x00, 0x1b, 0x9e, 0x04, 0xd9, 0x99, 0x71, 0x26, 0x93, 0xb4,
	0xec, 0xd0, 0x6c, 0xa2, 0xc9, 0x3d, 0xe7, 0xe8, 0x1e, 0xcf, 0xbd, 0xc7, 0x09, 0xd8, 0xd1, 0xc3,
	0xae, 0x26, 0x8a, 0x75, 0xa9, 0x42, 0x9a, 0x6a, 0xcd, 0x04, 0x47, 0x7a, 0x38, 0x18, 0x60, 0x35,
	0x0a, 0xa5, 0x12, 0x46, 0x40, 0x45, 0x98, 0x26, 0x02, 0x31, 0xa1, 0xd1, 0xb9, 0x42, 0x8c, 0x46,
	0x7d, 0x7a, 0xf4, 0xb0, 0xb9, 0x8f, 0x84, 0xa4, 0x2a, 0x9c, 0xa8, 0xc3, 0x44, 0x1d, 0x72, 0x11,
	0x51, 0xed, 0x3e, 0x43, 0xad, 0xfa, 0x48, 0x89, 0x98, 0x6a, 0xff, 0x14, 0x32, 0x6e, 0xa8, 0xea,
	0x61, 0x42, 0x93, 0x2e, 0x8c, 0xea, 0x99, 0xda, 0xa8, 0xf1, 0x1a, 0xdc, 0xba, 0xdc, 0x17, 0x7a,
	0xfa, 0xf8, 0x65, 0x07, 0x5e, 0x07, 0xab, 0xb6, 0x07, 0xe2, 0x78, 0x40, 0xeb, 0xc1, 0x4e, 0xb0,
	0xb7, 0xda, 0x5e, 0xb1, 0x85, 0x67, 0x78, 0x40, 0x61, 0x0d, 0x94, 0xb5, 0xea, 0xd7, 0x4b, 0xae,
	0x6c, 0x1f, 0xe1, 0x1d, 0xb0, 0x3e, 0x69, 0xe3, 0x34, 0x65, 0x07, 0x56, 0x7d, 0xd5, 0x0a, 0x1b,
	0xbf, 0x4b, 0xf3, 0x3b, 0x1b, 0x6c, 0x28, 0x22, 0x62, 0xc8, 0x8d, 0x86, 0x4d, 0xb0, 0xc5, 0x38,
	0x33, 0x0c, 0xc7, 0xec, 0x0d, 0x8d, 0x52, 0x8e, 0x76, 0x26, 0xaa, 0xed, 0xcd, 0x0c, 0xd6, 0x49,
	0x20, 0x78, 0x1f, 0x6c, 0x12, 0xc1, 0x39, 0x25, 0x86, 0xf1, 0xfe, 0x44, 0x51, 0x72, 0x0a, 0x38,
	0x81, 0xbc, 0xe0, 0x1e, 0x48, 0xab, 0xd9, 0x0e, 0x65, 0xc7, 0xdf, 0xf0, 0x48, 0x96, 0x8e, 0x89,
	0x61, 0x67, 0x78, 0x8a, 0xbe, 0x30, 0xa6, 0x7b, 0xc4, 0xd3, 0x6f, 0x83, 0x2a, 0x8b, 0x62, 0x3a,
	0x61, 0x2e, 0x3a, 0x66, 0xc5, 0x16, 0x3d, 0xe9, 0x01, 0xf8, 0x3f, 0xb2, 0x93, 0x9f, 0xb5, 0xbd,
	0xe4, 0xd8, 0xdb, 0x53, 0xa8, 0x97, 0xed, 0x82, 0x0a, 0xe5, 0x19, 0x13, 0xcb, 0x8e, 0xbc, 0x46,
	0xb9, 0x6f, 0xdf, 0xf8, 0xb5, 0x30, 0x7f, 0xed, 0xdc, 0x4b, 0x4e, 0x86, 0x0c, 0x3f, 0x06, 0x60,
	0x51, 0x4a, 0x29, 0xc6, 0xc3, 0x5d, 0x3b, 0x78, 0x17, 0x84, 0xff, 0x7e, 0x11, 0xc3, 0xbf, 0xec,
	0x42, 0x7b, 0x6c, 0x11, 0x7e, 0x0b, 0x00, 0x64, 0x12, 0x65, 0xd8, 0xd1, 0x29, 0x91, 0x6e, 0xbe,
	0x45, 0x75, 0x5e, 0x63, 0xb2, 0xe3, 0x29, 0x8f, 0x4e, 0x89, 0x84, 0x3f, 0x02, 0xb0, 0x35, 0x7d,
	0x08, 0x89, 0xc9, 0x2b, 0x6a, 0xdc, 0xda, 0x15, 0xf5, 0x18, 0x30, 0x7b, 0x8c, 0x96, 0xf3, 0xdb,
	0xf8, 0x54, 0x02, 0x37, 0xe6, 0xe8, 0x70, 0x2f, 0x4d, 0xf0, 0xbe, 0x4d, 0x30, 0x92, 0x4a, 0xf4,
	0x15, 0xd5, 0x3a, 0x9f, 0x60, 0xc8, 0x78, 0x2b, 0x81, 0xfc, 0x56, 0xdf, 0xb5, 0xf3, 0x3d, 0x3b,
	0x42, 0x82, 0xc7, 0xa3, 0x7c, 0x7e, 0x6b, 0x16, 0x79, 0xce, 0xe3, 0x51, 0x8e, 0x7d, 0x9c, 0x63,
	0x97, 0x3d, 0xfb, 0x78, 0x8a, 0xdd, 0x04, 0xdb, 0xd1, 0x10, 0xc7, 0x48, 0x62, 0x65, 0xd0, 0x50,
	0xe6, 0xf3, 0x0b, 0x2d, 0xd8, 0xc2, 0xca, 0xbc, 0x90, 0x5e, 0xb2, 0x07, 0x6a, 0x4e, 0x92, 0x65,
	0x8f, 0x33, 0xbc, 0x6e, 0xeb, 0x19, 0xe6, 0x2e, 0xa8, 0xc4, 0x98, 0xe4, 0xb3, 0xbb, 0x16, 0x63,
	0xe2, 0xe3, 0xf8, 0x73, 0x01, 0xdc, 0x9c, 0xff, 0xba, 0xd2, 0x2c, 0x7e, 0xc8, 0x65, 0xf1, 0xa2,
	0x28, 0xab, 0xe0, 0x47, 0x9a, 0x06, 0xf1, 0xeb, 0x55, 0x41, 0x2c, 0xa0, 0xed, 0xd9, 0x14, 0x7e,
	0xbf, 0x3a, 0x85, 0x05, 0x3c, 0xc3, 0xbc, 0x08, 0x5e, 0x94, 0xc1, 0xb5, 0xcb, 0x7f, 0xc1, 0xe1,
	0xe7, 0x00, 0xac, 0x8c, 0x63, 0x7c, 0xae, 0xea, 0x07, 0xee, 0x60, 0xef, 0x8b, 0x75, 0xbd, 0x24,
	0x84, 0xf6, 0xb2, 0xfb, 0x7a, 0xa2, 0xe0, 0x97, 0x00, 0x6c, 0xe0, 0x28, 0x72, 0x17, 0x46, 0x0f,
	0x0f, 0x58, 0x3c, 0xb2, 0xde, 0x0f, 0x9d, 0xf7, 0xb7, 0x05, 0x1a, 0x4a, 0x6a, 0xfc, 0xbf, 0xc4,
	0xec, 0x13, 0xe7, 0xf5, 0x44, 0x75, 0x97, 0xdc, 0xdf, 0xb9, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0x77, 0x10, 0x9b, 0xf2, 0x09, 0x00, 0x00,
}
