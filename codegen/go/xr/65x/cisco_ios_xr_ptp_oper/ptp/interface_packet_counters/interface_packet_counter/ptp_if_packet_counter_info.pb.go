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
// source: ptp_if_packet_counter_info.proto

package cisco_ios_xr_ptp_oper_ptp_interface_packet_counters_interface_packet_counter

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

type PtpIfPacketCounterInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfPacketCounterInfo_KEYS) Reset()         { *m = PtpIfPacketCounterInfo_KEYS{} }
func (m *PtpIfPacketCounterInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfPacketCounterInfo_KEYS) ProtoMessage()    {}
func (*PtpIfPacketCounterInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{0}
}

func (m *PtpIfPacketCounterInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfPacketCounterInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpIfPacketCounterInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfPacketCounterInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfPacketCounterInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfPacketCounterInfo_KEYS.Merge(m, src)
}
func (m *PtpIfPacketCounterInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfPacketCounterInfo_KEYS.Size(m)
}
func (m *PtpIfPacketCounterInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfPacketCounterInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfPacketCounterInfo_KEYS proto.InternalMessageInfo

func (m *PtpIfPacketCounterInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PtpBagPacketCounters struct {
	AnnounceSent                      uint32   `protobuf:"varint,1,opt,name=announce_sent,json=announceSent,proto3" json:"announce_sent,omitempty"`
	AnnounceReceived                  uint32   `protobuf:"varint,2,opt,name=announce_received,json=announceReceived,proto3" json:"announce_received,omitempty"`
	AnnounceDropped                   uint32   `protobuf:"varint,3,opt,name=announce_dropped,json=announceDropped,proto3" json:"announce_dropped,omitempty"`
	SyncSent                          uint32   `protobuf:"varint,4,opt,name=sync_sent,json=syncSent,proto3" json:"sync_sent,omitempty"`
	SyncReceived                      uint32   `protobuf:"varint,5,opt,name=sync_received,json=syncReceived,proto3" json:"sync_received,omitempty"`
	SyncDropped                       uint32   `protobuf:"varint,6,opt,name=sync_dropped,json=syncDropped,proto3" json:"sync_dropped,omitempty"`
	FollowUpSent                      uint32   `protobuf:"varint,7,opt,name=follow_up_sent,json=followUpSent,proto3" json:"follow_up_sent,omitempty"`
	FollowUpReceived                  uint32   `protobuf:"varint,8,opt,name=follow_up_received,json=followUpReceived,proto3" json:"follow_up_received,omitempty"`
	FollowUpDropped                   uint32   `protobuf:"varint,9,opt,name=follow_up_dropped,json=followUpDropped,proto3" json:"follow_up_dropped,omitempty"`
	DelayRequestSent                  uint32   `protobuf:"varint,10,opt,name=delay_request_sent,json=delayRequestSent,proto3" json:"delay_request_sent,omitempty"`
	DelayRequestReceived              uint32   `protobuf:"varint,11,opt,name=delay_request_received,json=delayRequestReceived,proto3" json:"delay_request_received,omitempty"`
	DelayRequestDropped               uint32   `protobuf:"varint,12,opt,name=delay_request_dropped,json=delayRequestDropped,proto3" json:"delay_request_dropped,omitempty"`
	DelayResponseSent                 uint32   `protobuf:"varint,13,opt,name=delay_response_sent,json=delayResponseSent,proto3" json:"delay_response_sent,omitempty"`
	DelayResponseReceived             uint32   `protobuf:"varint,14,opt,name=delay_response_received,json=delayResponseReceived,proto3" json:"delay_response_received,omitempty"`
	DelayResponseDropped              uint32   `protobuf:"varint,15,opt,name=delay_response_dropped,json=delayResponseDropped,proto3" json:"delay_response_dropped,omitempty"`
	PeerDelayRequestSent              uint32   `protobuf:"varint,16,opt,name=peer_delay_request_sent,json=peerDelayRequestSent,proto3" json:"peer_delay_request_sent,omitempty"`
	PeerDelayRequestReceived          uint32   `protobuf:"varint,17,opt,name=peer_delay_request_received,json=peerDelayRequestReceived,proto3" json:"peer_delay_request_received,omitempty"`
	PeerDelayRequestDropped           uint32   `protobuf:"varint,18,opt,name=peer_delay_request_dropped,json=peerDelayRequestDropped,proto3" json:"peer_delay_request_dropped,omitempty"`
	PeerDelayResponseSent             uint32   `protobuf:"varint,19,opt,name=peer_delay_response_sent,json=peerDelayResponseSent,proto3" json:"peer_delay_response_sent,omitempty"`
	PeerDelayResponseReceived         uint32   `protobuf:"varint,20,opt,name=peer_delay_response_received,json=peerDelayResponseReceived,proto3" json:"peer_delay_response_received,omitempty"`
	PeerDelayResponseDropped          uint32   `protobuf:"varint,21,opt,name=peer_delay_response_dropped,json=peerDelayResponseDropped,proto3" json:"peer_delay_response_dropped,omitempty"`
	PeerDelayResponseFollowUpSent     uint32   `protobuf:"varint,22,opt,name=peer_delay_response_follow_up_sent,json=peerDelayResponseFollowUpSent,proto3" json:"peer_delay_response_follow_up_sent,omitempty"`
	PeerDelayResponseFollowUpReceived uint32   `protobuf:"varint,23,opt,name=peer_delay_response_follow_up_received,json=peerDelayResponseFollowUpReceived,proto3" json:"peer_delay_response_follow_up_received,omitempty"`
	PeerDelayResponseFollowUpDropped  uint32   `protobuf:"varint,24,opt,name=peer_delay_response_follow_up_dropped,json=peerDelayResponseFollowUpDropped,proto3" json:"peer_delay_response_follow_up_dropped,omitempty"`
	SignalingSent                     uint32   `protobuf:"varint,25,opt,name=signaling_sent,json=signalingSent,proto3" json:"signaling_sent,omitempty"`
	SignalingReceived                 uint32   `protobuf:"varint,26,opt,name=signaling_received,json=signalingReceived,proto3" json:"signaling_received,omitempty"`
	SignalingDropped                  uint32   `protobuf:"varint,27,opt,name=signaling_dropped,json=signalingDropped,proto3" json:"signaling_dropped,omitempty"`
	ManagementSent                    uint32   `protobuf:"varint,28,opt,name=management_sent,json=managementSent,proto3" json:"management_sent,omitempty"`
	ManagementReceived                uint32   `protobuf:"varint,29,opt,name=management_received,json=managementReceived,proto3" json:"management_received,omitempty"`
	ManagementDropped                 uint32   `protobuf:"varint,30,opt,name=management_dropped,json=managementDropped,proto3" json:"management_dropped,omitempty"`
	OtherPacketsSent                  uint32   `protobuf:"varint,31,opt,name=other_packets_sent,json=otherPacketsSent,proto3" json:"other_packets_sent,omitempty"`
	OtherPacketsReceived              uint32   `protobuf:"varint,32,opt,name=other_packets_received,json=otherPacketsReceived,proto3" json:"other_packets_received,omitempty"`
	OtherPacketsDropped               uint32   `protobuf:"varint,33,opt,name=other_packets_dropped,json=otherPacketsDropped,proto3" json:"other_packets_dropped,omitempty"`
	TotalPacketsSent                  uint32   `protobuf:"varint,34,opt,name=total_packets_sent,json=totalPacketsSent,proto3" json:"total_packets_sent,omitempty"`
	TotalPacketsReceived              uint32   `protobuf:"varint,35,opt,name=total_packets_received,json=totalPacketsReceived,proto3" json:"total_packets_received,omitempty"`
	TotalPacketsDropped               uint32   `protobuf:"varint,36,opt,name=total_packets_dropped,json=totalPacketsDropped,proto3" json:"total_packets_dropped,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *PtpBagPacketCounters) Reset()         { *m = PtpBagPacketCounters{} }
func (m *PtpBagPacketCounters) String() string { return proto.CompactTextString(m) }
func (*PtpBagPacketCounters) ProtoMessage()    {}
func (*PtpBagPacketCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{1}
}

func (m *PtpBagPacketCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagPacketCounters.Unmarshal(m, b)
}
func (m *PtpBagPacketCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagPacketCounters.Marshal(b, m, deterministic)
}
func (m *PtpBagPacketCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagPacketCounters.Merge(m, src)
}
func (m *PtpBagPacketCounters) XXX_Size() int {
	return xxx_messageInfo_PtpBagPacketCounters.Size(m)
}
func (m *PtpBagPacketCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagPacketCounters.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagPacketCounters proto.InternalMessageInfo

func (m *PtpBagPacketCounters) GetAnnounceSent() uint32 {
	if m != nil {
		return m.AnnounceSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetAnnounceReceived() uint32 {
	if m != nil {
		return m.AnnounceReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetAnnounceDropped() uint32 {
	if m != nil {
		return m.AnnounceDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSyncSent() uint32 {
	if m != nil {
		return m.SyncSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSyncReceived() uint32 {
	if m != nil {
		return m.SyncReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSyncDropped() uint32 {
	if m != nil {
		return m.SyncDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetFollowUpSent() uint32 {
	if m != nil {
		return m.FollowUpSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetFollowUpReceived() uint32 {
	if m != nil {
		return m.FollowUpReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetFollowUpDropped() uint32 {
	if m != nil {
		return m.FollowUpDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayRequestSent() uint32 {
	if m != nil {
		return m.DelayRequestSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayRequestReceived() uint32 {
	if m != nil {
		return m.DelayRequestReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayRequestDropped() uint32 {
	if m != nil {
		return m.DelayRequestDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayResponseSent() uint32 {
	if m != nil {
		return m.DelayResponseSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayResponseReceived() uint32 {
	if m != nil {
		return m.DelayResponseReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetDelayResponseDropped() uint32 {
	if m != nil {
		return m.DelayResponseDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayRequestSent() uint32 {
	if m != nil {
		return m.PeerDelayRequestSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayRequestReceived() uint32 {
	if m != nil {
		return m.PeerDelayRequestReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayRequestDropped() uint32 {
	if m != nil {
		return m.PeerDelayRequestDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseSent() uint32 {
	if m != nil {
		return m.PeerDelayResponseSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseReceived() uint32 {
	if m != nil {
		return m.PeerDelayResponseReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseDropped() uint32 {
	if m != nil {
		return m.PeerDelayResponseDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseFollowUpSent() uint32 {
	if m != nil {
		return m.PeerDelayResponseFollowUpSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseFollowUpReceived() uint32 {
	if m != nil {
		return m.PeerDelayResponseFollowUpReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetPeerDelayResponseFollowUpDropped() uint32 {
	if m != nil {
		return m.PeerDelayResponseFollowUpDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSignalingSent() uint32 {
	if m != nil {
		return m.SignalingSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSignalingReceived() uint32 {
	if m != nil {
		return m.SignalingReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetSignalingDropped() uint32 {
	if m != nil {
		return m.SignalingDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetManagementSent() uint32 {
	if m != nil {
		return m.ManagementSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetManagementReceived() uint32 {
	if m != nil {
		return m.ManagementReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetManagementDropped() uint32 {
	if m != nil {
		return m.ManagementDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetOtherPacketsSent() uint32 {
	if m != nil {
		return m.OtherPacketsSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetOtherPacketsReceived() uint32 {
	if m != nil {
		return m.OtherPacketsReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetOtherPacketsDropped() uint32 {
	if m != nil {
		return m.OtherPacketsDropped
	}
	return 0
}

func (m *PtpBagPacketCounters) GetTotalPacketsSent() uint32 {
	if m != nil {
		return m.TotalPacketsSent
	}
	return 0
}

func (m *PtpBagPacketCounters) GetTotalPacketsReceived() uint32 {
	if m != nil {
		return m.TotalPacketsReceived
	}
	return 0
}

func (m *PtpBagPacketCounters) GetTotalPacketsDropped() uint32 {
	if m != nil {
		return m.TotalPacketsDropped
	}
	return 0
}

type PtpBagMacAddrType struct {
	Macaddr              string   `protobuf:"bytes,1,opt,name=macaddr,proto3" json:"macaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagMacAddrType) Reset()         { *m = PtpBagMacAddrType{} }
func (m *PtpBagMacAddrType) String() string { return proto.CompactTextString(m) }
func (*PtpBagMacAddrType) ProtoMessage()    {}
func (*PtpBagMacAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{2}
}

func (m *PtpBagMacAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagMacAddrType.Unmarshal(m, b)
}
func (m *PtpBagMacAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagMacAddrType.Marshal(b, m, deterministic)
}
func (m *PtpBagMacAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagMacAddrType.Merge(m, src)
}
func (m *PtpBagMacAddrType) XXX_Size() int {
	return xxx_messageInfo_PtpBagMacAddrType.Size(m)
}
func (m *PtpBagMacAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagMacAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagMacAddrType proto.InternalMessageInfo

func (m *PtpBagMacAddrType) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

type PtpBagIpv6AddrType struct {
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagIpv6AddrType) Reset()         { *m = PtpBagIpv6AddrType{} }
func (m *PtpBagIpv6AddrType) String() string { return proto.CompactTextString(m) }
func (*PtpBagIpv6AddrType) ProtoMessage()    {}
func (*PtpBagIpv6AddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{3}
}

func (m *PtpBagIpv6AddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagIpv6AddrType.Unmarshal(m, b)
}
func (m *PtpBagIpv6AddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagIpv6AddrType.Marshal(b, m, deterministic)
}
func (m *PtpBagIpv6AddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagIpv6AddrType.Merge(m, src)
}
func (m *PtpBagIpv6AddrType) XXX_Size() int {
	return xxx_messageInfo_PtpBagIpv6AddrType.Size(m)
}
func (m *PtpBagIpv6AddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagIpv6AddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagIpv6AddrType proto.InternalMessageInfo

func (m *PtpBagIpv6AddrType) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PtpBagAddress struct {
	Encapsulation        string              `protobuf:"bytes,1,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	AddressUnknown       bool                `protobuf:"varint,2,opt,name=address_unknown,json=addressUnknown,proto3" json:"address_unknown,omitempty"`
	MacAddress           *PtpBagMacAddrType  `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Ipv4Address          string              `protobuf:"bytes,4,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          *PtpBagIpv6AddrType `protobuf:"bytes,5,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PtpBagAddress) Reset()         { *m = PtpBagAddress{} }
func (m *PtpBagAddress) String() string { return proto.CompactTextString(m) }
func (*PtpBagAddress) ProtoMessage()    {}
func (*PtpBagAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{4}
}

func (m *PtpBagAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagAddress.Unmarshal(m, b)
}
func (m *PtpBagAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagAddress.Marshal(b, m, deterministic)
}
func (m *PtpBagAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagAddress.Merge(m, src)
}
func (m *PtpBagAddress) XXX_Size() int {
	return xxx_messageInfo_PtpBagAddress.Size(m)
}
func (m *PtpBagAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagAddress proto.InternalMessageInfo

func (m *PtpBagAddress) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *PtpBagAddress) GetAddressUnknown() bool {
	if m != nil {
		return m.AddressUnknown
	}
	return false
}

func (m *PtpBagAddress) GetMacAddress() *PtpBagMacAddrType {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *PtpBagAddress) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PtpBagAddress) GetIpv6Address() *PtpBagIpv6AddrType {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

type PtpBagPeerPacketCounters struct {
	Address              *PtpBagAddress        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Counters             *PtpBagPacketCounters `protobuf:"bytes,2,opt,name=counters,proto3" json:"counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpBagPeerPacketCounters) Reset()         { *m = PtpBagPeerPacketCounters{} }
func (m *PtpBagPeerPacketCounters) String() string { return proto.CompactTextString(m) }
func (*PtpBagPeerPacketCounters) ProtoMessage()    {}
func (*PtpBagPeerPacketCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{5}
}

func (m *PtpBagPeerPacketCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagPeerPacketCounters.Unmarshal(m, b)
}
func (m *PtpBagPeerPacketCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagPeerPacketCounters.Marshal(b, m, deterministic)
}
func (m *PtpBagPeerPacketCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagPeerPacketCounters.Merge(m, src)
}
func (m *PtpBagPeerPacketCounters) XXX_Size() int {
	return xxx_messageInfo_PtpBagPeerPacketCounters.Size(m)
}
func (m *PtpBagPeerPacketCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagPeerPacketCounters.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagPeerPacketCounters proto.InternalMessageInfo

func (m *PtpBagPeerPacketCounters) GetAddress() *PtpBagAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PtpBagPeerPacketCounters) GetCounters() *PtpBagPacketCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

type PtpIfPacketCounterInfo struct {
	Counters             *PtpBagPacketCounters       `protobuf:"bytes,50,opt,name=counters,proto3" json:"counters,omitempty"`
	PeerCounter          []*PtpBagPeerPacketCounters `protobuf:"bytes,51,rep,name=peer_counter,json=peerCounter,proto3" json:"peer_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PtpIfPacketCounterInfo) Reset()         { *m = PtpIfPacketCounterInfo{} }
func (m *PtpIfPacketCounterInfo) String() string { return proto.CompactTextString(m) }
func (*PtpIfPacketCounterInfo) ProtoMessage()    {}
func (*PtpIfPacketCounterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a1d9cf4e923ccf7, []int{6}
}

func (m *PtpIfPacketCounterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfPacketCounterInfo.Unmarshal(m, b)
}
func (m *PtpIfPacketCounterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfPacketCounterInfo.Marshal(b, m, deterministic)
}
func (m *PtpIfPacketCounterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfPacketCounterInfo.Merge(m, src)
}
func (m *PtpIfPacketCounterInfo) XXX_Size() int {
	return xxx_messageInfo_PtpIfPacketCounterInfo.Size(m)
}
func (m *PtpIfPacketCounterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfPacketCounterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfPacketCounterInfo proto.InternalMessageInfo

func (m *PtpIfPacketCounterInfo) GetCounters() *PtpBagPacketCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *PtpIfPacketCounterInfo) GetPeerCounter() []*PtpBagPeerPacketCounters {
	if m != nil {
		return m.PeerCounter
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpIfPacketCounterInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_if_packet_counter_info_KEYS")
	proto.RegisterType((*PtpBagPacketCounters)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_bag_packet_counters")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_bag_address")
	proto.RegisterType((*PtpBagPeerPacketCounters)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_bag_peer_packet_counters")
	proto.RegisterType((*PtpIfPacketCounterInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_packet_counters.interface_packet_counter.ptp_if_packet_counter_info")
}

func init() { proto.RegisterFile("ptp_if_packet_counter_info.proto", fileDescriptor_9a1d9cf4e923ccf7) }

var fileDescriptor_9a1d9cf4e923ccf7 = []byte{
	// 937 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x97, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0xe5, 0xa4, 0x6d, 0x92, 0xb7, 0xb1, 0xd3, 0x4c, 0x9a, 0xd8, 0x4d, 0x52, 0xea, 0x38,
	0x09, 0x14, 0xda, 0x1a, 0xe1, 0x86, 0x72, 0xa8, 0x10, 0x42, 0x14, 0x04, 0x02, 0xf1, 0x63, 0xab,
	0x1e, 0x38, 0xa0, 0xd1, 0x74, 0x3d, 0x09, 0x4b, 0xed, 0x99, 0x61, 0x77, 0xdd, 0x90, 0x23, 0x12,
	0x12, 0x27, 0x6e, 0xdc, 0xf8, 0x27, 0xf8, 0x33, 0xf8, 0xb3, 0xd0, 0xbe, 0x99, 0x79, 0x3b, 0xbb,
	0xb1, 0x7b, 0x22, 0xe2, 0xe8, 0xef, 0x7b, 0xdf, 0xf7, 0x3e, 0xef, 0xed, 0x8f, 0x59, 0x43, 0xdf,
	0x14, 0x86, 0xa7, 0xa7, 0xdc, 0x88, 0xe4, 0xa5, 0x2c, 0x78, 0xa2, 0x67, 0xaa, 0x90, 0x19, 0x4f,
	0xd5, 0xa9, 0x1e, 0x9a, 0x4c, 0x17, 0x9a, 0x7d, 0x95, 0xa4, 0x79, 0xa2, 0x79, 0xaa, 0x73, 0xfe,
	0x4b, 0xc6, 0xcb, 0x74, 0x6d, 0x64, 0x36, 0x34, 0x85, 0x19, 0xa6, 0x65, 0xee, 0xa9, 0x48, 0x64,
	0xc3, 0x9e, 0x2f, 0x8c, 0x0c, 0x3e, 0x87, 0xbb, 0x8b, 0x3b, 0xf2, 0x2f, 0x3f, 0xfd, 0xfe, 0x19,
	0x3b, 0x86, 0x4e, 0x65, 0x57, 0x62, 0x2a, 0x7b, 0xad, 0x7e, 0xeb, 0xde, 0x5a, 0xdc, 0x26, 0xf5,
	0x6b, 0x31, 0x95, 0x83, 0x7f, 0x3a, 0xd0, 0x2d, 0x4b, 0xbd, 0x10, 0x67, 0xcd, 0xf6, 0xec, 0x10,
	0xda, 0x42, 0x29, 0x3d, 0x53, 0x89, 0xe4, 0xb9, 0x54, 0x05, 0x56, 0x68, 0xc7, 0xeb, 0x5e, 0x7c,
	0x26, 0x55, 0xc1, 0xee, 0xc3, 0x26, 0x25, 0x65, 0x32, 0x91, 0xe9, 0x2b, 0x39, 0xee, 0x2d, 0x61,
	0xe2, 0x4d, 0x1f, 0x88, 0x9d, 0xce, 0xde, 0x06, 0xd2, 0xf8, 0x38, 0xd3, 0xc6, 0xc8, 0x71, 0x6f,
	0x19, 0x73, 0x37, 0xbc, 0xfe, 0xd4, 0xca, 0x6c, 0x0f, 0xd6, 0xf2, 0x0b, 0x95, 0xd8, 0xc6, 0xd7,
	0x30, 0x67, 0xb5, 0x14, 0xb0, 0xe9, 0x21, 0xb4, 0x31, 0x48, 0x0d, 0xaf, 0x5b, 0xb2, 0x52, 0xa4,
	0x66, 0x07, 0x80, 0xbf, 0xa9, 0xd1, 0x0d, 0xcc, 0x89, 0x4a, 0xcd, 0x37, 0x39, 0x82, 0xce, 0xa9,
	0x9e, 0x4c, 0xf4, 0x39, 0x9f, 0x19, 0xdb, 0x69, 0xc5, 0x16, 0xb2, 0xea, 0x73, 0x83, 0xdd, 0x1e,
	0x00, 0xab, 0xb2, 0xa8, 0xe5, 0xaa, 0x9d, 0xd1, 0x67, 0x52, 0xdb, 0x77, 0x60, 0xb3, 0xca, 0xf6,
	0xbd, 0xd7, 0xec, 0x90, 0x3e, 0xd9, 0xf7, 0x7f, 0x00, 0x6c, 0x2c, 0x27, 0xe2, 0x82, 0x67, 0xf2,
	0xe7, 0x99, 0xcc, 0x0b, 0xcb, 0x00, 0xb6, 0x32, 0x46, 0x62, 0x1b, 0x40, 0x8e, 0x13, 0xd8, 0xa9,
	0x67, 0x13, 0x4b, 0x84, 0x8e, 0x5b, 0xa1, 0x83, 0x78, 0x46, 0xb0, 0x5d, 0x77, 0x79, 0xa6, 0x75,
	0x34, 0x6d, 0x85, 0x26, 0xcf, 0x35, 0x84, 0x2d, 0xef, 0xc9, 0x8d, 0x56, 0xb9, 0xbb, 0xfe, 0x6d,
	0x74, 0x6c, 0x3a, 0x87, 0x8d, 0x20, 0xd9, 0x63, 0xe8, 0x36, 0xf2, 0x09, 0xad, 0x83, 0x9e, 0xed,
	0x9a, 0x87, 0xd8, 0x82, 0x89, 0x9c, 0xcf, 0xc3, 0x6d, 0xd4, 0x26, 0xb2, 0x41, 0x4f, 0xf7, 0x3e,
	0x74, 0x8d, 0x94, 0x19, 0x9f, 0xb3, 0xba, 0x9b, 0xd6, 0x56, 0x86, 0x9f, 0x36, 0xd7, 0xf7, 0x21,
	0xec, 0xcd, 0xb1, 0x11, 0xe8, 0x26, 0x5a, 0x7b, 0x4d, 0x2b, 0xb1, 0x3e, 0x81, 0xdd, 0x39, 0x76,
	0xcf, 0xcb, 0xd0, 0xdd, 0x6d, 0xba, 0x3d, 0xf2, 0x07, 0xd0, 0xab, 0x99, 0xc3, 0xad, 0x6e, 0xd9,
	0x0d, 0x05, 0xd6, 0x60, 0xb3, 0x1f, 0xc1, 0xfe, 0x3c, 0x23, 0x51, 0xdf, 0x42, 0xf3, 0xed, 0x4b,
	0x66, 0xc2, 0x6e, 0x4e, 0xdd, 0xd8, 0xf3, 0xf6, 0xa5, 0xa9, 0xeb, 0xbb, 0xfe, 0x02, 0x06, 0xf3,
	0xec, 0x8d, 0xa7, 0x66, 0x07, 0xab, 0xdc, 0xb9, 0x54, 0xe5, 0xb3, 0xf0, 0x31, 0xfa, 0x0e, 0xde,
	0x7c, 0x7d, 0x29, 0x1a, 0xaa, 0x8b, 0xe5, 0x0e, 0x16, 0x96, 0xa3, 0xe1, 0xbe, 0x81, 0xe3, 0xd7,
	0x97, 0xf4, 0x63, 0xf6, 0xb0, 0x62, 0x7f, 0x61, 0x45, 0x3f, 0xee, 0x31, 0x74, 0xf2, 0xf4, 0x4c,
	0x89, 0x49, 0xaa, 0xce, 0xec, 0x68, 0xb7, 0xd1, 0xd9, 0x26, 0x15, 0x47, 0x79, 0x08, 0xac, 0x4a,
	0x23, 0xec, 0x5d, 0xfb, 0x78, 0x50, 0x84, 0x30, 0xef, 0x43, 0x25, 0x12, 0xd2, 0x9e, 0x7d, 0xca,
	0x29, 0xe0, 0x11, 0xde, 0x82, 0x8d, 0xa9, 0x50, 0xe2, 0x4c, 0x4e, 0xa5, 0x72, 0x77, 0xf5, 0x3e,
	0xa6, 0x76, 0x2a, 0x19, 0x21, 0xde, 0x85, 0xad, 0x20, 0x91, 0x28, 0xee, 0x60, 0x32, 0xab, 0x42,
	0x84, 0xf1, 0x10, 0x02, 0x95, 0x38, 0xde, 0xb0, 0xd4, 0x55, 0x24, 0x78, 0x39, 0xe9, 0xe2, 0x47,
	0x99, 0xb9, 0x73, 0x21, 0xb7, 0x2c, 0x77, 0x2d, 0x36, 0x46, 0xbe, 0xb5, 0x01, 0xff, 0x72, 0xaa,
	0x67, 0x13, 0x50, 0xdf, 0x3e, 0x93, 0xa1, 0x23, 0x7c, 0x39, 0xd5, 0x5d, 0x9e, 0xea, 0xc0, 0xbe,
	0x9c, 0x42, 0x53, 0xc0, 0x55, 0xe8, 0x42, 0x4c, 0xea, 0x5c, 0x03, 0xcb, 0x85, 0x91, 0x06, 0x57,
	0x3d, 0x9b, 0xb8, 0x0e, 0x2d, 0x57, 0xe8, 0x08, 0xb9, 0xea, 0x2e, 0xcf, 0x75, 0x64, 0xb9, 0x42,
	0x93, 0xe3, 0x1a, 0xbc, 0x07, 0xdb, 0xfe, 0x24, 0x9d, 0x8a, 0x84, 0x8b, 0xf1, 0x38, 0xe3, 0xc5,
	0x85, 0x91, 0xac, 0x07, 0x2b, 0x53, 0x91, 0x94, 0xbf, 0xdd, 0x19, 0xec, 0x7f, 0x0e, 0x9e, 0xc0,
	0x8e, 0xb7, 0xa4, 0xe6, 0xd5, 0xe3, 0xc0, 0x73, 0x00, 0xeb, 0xa4, 0xc8, 0x3c, 0x77, 0xc6, 0xa8,
	0xd4, 0x3e, 0xb6, 0xd2, 0xe0, 0xcf, 0x65, 0xd8, 0xf0, 0x6e, 0x97, 0xc6, 0x8e, 0xa0, 0x2d, 0x55,
	0x22, 0x4c, 0x3e, 0x9b, 0x88, 0x22, 0xd5, 0xca, 0x1f, 0xfa, 0x35, 0xb1, 0xbc, 0xc5, 0x9c, 0x81,
	0xcf, 0xd4, 0x4b, 0xa5, 0xcf, 0x15, 0x9e, 0xd8, 0xab, 0x71, 0xc7, 0xc9, 0xcf, 0xad, 0xca, 0x7e,
	0x6b, 0x41, 0xe4, 0x67, 0x29, 0x29, 0xca, 0xb3, 0x3a, 0x1a, 0x25, 0xc3, 0xff, 0xf2, 0x63, 0x66,
	0x38, 0x77, 0x69, 0x31, 0x4c, 0x45, 0xe2, 0x26, 0x75, 0xcb, 0x38, 0x21, 0x8c, 0x6b, 0xb4, 0x8c,
	0x13, 0x9f, 0xf2, 0x7b, 0xab, 0xb1, 0xb0, 0xeb, 0x88, 0x3a, 0xbe, 0x1a, 0xd4, 0xfa, 0xc5, 0xaa,
	0x5f, 0x96, 0xbf, 0x96, 0x60, 0x9f, 0xbe, 0xa8, 0x24, 0xdd, 0xda, 0xd5, 0x67, 0xd5, 0x39, 0xac,
	0x84, 0x57, 0x35, 0x1a, 0xfd, 0x70, 0x35, 0x90, 0xae, 0x49, 0xec, 0xbb, 0xb1, 0x5f, 0x5b, 0xb0,
	0xea, 0xed, 0x78, 0xc1, 0xa3, 0x91, 0xbc, 0x9a, 0xd6, 0x0d, 0x43, 0x4c, 0x6d, 0x07, 0x7f, 0x2f,
	0xc1, 0xee, 0xe2, 0x4f, 0xd7, 0x3a, 0xe2, 0xe8, 0x7f, 0x41, 0x64, 0x7f, 0xb4, 0x60, 0x1d, 0x2f,
	0x9c, 0x53, 0x7a, 0x8f, 0xfa, 0xcb, 0xf7, 0xa2, 0xd1, 0x4f, 0x57, 0xc4, 0x31, 0xe7, 0x16, 0x89,
	0xa3, 0x52, 0xfd, 0xc4, 0xfe, 0x7a, 0x71, 0x03, 0xff, 0x41, 0x3c, 0xfa, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x29, 0xd2, 0xe3, 0x6b, 0x65, 0x0c, 0x00, 0x00,
}
