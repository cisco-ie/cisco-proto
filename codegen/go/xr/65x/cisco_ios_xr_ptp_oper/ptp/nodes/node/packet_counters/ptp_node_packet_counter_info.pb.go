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
// source: ptp_node_packet_counter_info.proto

package cisco_ios_xr_ptp_oper_ptp_nodes_node_packet_counters

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

type PtpNodePacketCounterInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpNodePacketCounterInfo_KEYS) Reset()         { *m = PtpNodePacketCounterInfo_KEYS{} }
func (m *PtpNodePacketCounterInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpNodePacketCounterInfo_KEYS) ProtoMessage()    {}
func (*PtpNodePacketCounterInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_786518142ae33ccc, []int{0}
}

func (m *PtpNodePacketCounterInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpNodePacketCounterInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpNodePacketCounterInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpNodePacketCounterInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpNodePacketCounterInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpNodePacketCounterInfo_KEYS.Merge(m, src)
}
func (m *PtpNodePacketCounterInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpNodePacketCounterInfo_KEYS.Size(m)
}
func (m *PtpNodePacketCounterInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpNodePacketCounterInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpNodePacketCounterInfo_KEYS proto.InternalMessageInfo

func (m *PtpNodePacketCounterInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
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
	return fileDescriptor_786518142ae33ccc, []int{1}
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

type PtpBagDropReasons struct {
	NotReady             uint32   `protobuf:"varint,1,opt,name=not_ready,json=notReady,proto3" json:"not_ready,omitempty"`
	WrongDomain          uint32   `protobuf:"varint,2,opt,name=wrong_domain,json=wrongDomain,proto3" json:"wrong_domain,omitempty"`
	TooShort             uint32   `protobuf:"varint,3,opt,name=too_short,json=tooShort,proto3" json:"too_short,omitempty"`
	LoopedSamePort       uint32   `protobuf:"varint,4,opt,name=looped_same_port,json=loopedSamePort,proto3" json:"looped_same_port,omitempty"`
	LoopedHigherPort     uint32   `protobuf:"varint,5,opt,name=looped_higher_port,json=loopedHigherPort,proto3" json:"looped_higher_port,omitempty"`
	LoopedLowerPort      uint32   `protobuf:"varint,6,opt,name=looped_lower_port,json=loopedLowerPort,proto3" json:"looped_lower_port,omitempty"`
	NoTimestamp          uint32   `protobuf:"varint,7,opt,name=no_timestamp,json=noTimestamp,proto3" json:"no_timestamp,omitempty"`
	ZeroTimestamp        uint32   `protobuf:"varint,8,opt,name=zero_timestamp,json=zeroTimestamp,proto3" json:"zero_timestamp,omitempty"`
	InvalidTlVs          uint32   `protobuf:"varint,9,opt,name=invalid_tl_vs,json=invalidTlVs,proto3" json:"invalid_tl_vs,omitempty"`
	NotForUs             uint32   `protobuf:"varint,10,opt,name=not_for_us,json=notForUs,proto3" json:"not_for_us,omitempty"`
	NotListening         uint32   `protobuf:"varint,11,opt,name=not_listening,json=notListening,proto3" json:"not_listening,omitempty"`
	WrongMaster          uint32   `protobuf:"varint,12,opt,name=wrong_master,json=wrongMaster,proto3" json:"wrong_master,omitempty"`
	UnknownMaster        uint32   `protobuf:"varint,13,opt,name=unknown_master,json=unknownMaster,proto3" json:"unknown_master,omitempty"`
	NotMaster            uint32   `protobuf:"varint,14,opt,name=not_master,json=notMaster,proto3" json:"not_master,omitempty"`
	NotSlave             uint32   `protobuf:"varint,15,opt,name=not_slave,json=notSlave,proto3" json:"not_slave,omitempty"`
	NotGranted           uint32   `protobuf:"varint,16,opt,name=not_granted,json=notGranted,proto3" json:"not_granted,omitempty"`
	TooSlow              uint32   `protobuf:"varint,17,opt,name=too_slow,json=tooSlow,proto3" json:"too_slow,omitempty"`
	InvalidPacket        uint32   `protobuf:"varint,18,opt,name=invalid_packet,json=invalidPacket,proto3" json:"invalid_packet,omitempty"`
	WrongSequenceId      uint32   `protobuf:"varint,19,opt,name=wrong_sequence_id,json=wrongSequenceId,proto3" json:"wrong_sequence_id,omitempty"`
	NoOffloadSession     uint32   `protobuf:"varint,20,opt,name=no_offload_session,json=noOffloadSession,proto3" json:"no_offload_session,omitempty"`
	NotSupported         uint32   `protobuf:"varint,21,opt,name=not_supported,json=notSupported,proto3" json:"not_supported,omitempty"`
	MinClockClass        uint32   `protobuf:"varint,22,opt,name=min_clock_class,json=minClockClass,proto3" json:"min_clock_class,omitempty"`
	BadClockClass        uint32   `protobuf:"varint,23,opt,name=bad_clock_class,json=badClockClass,proto3" json:"bad_clock_class,omitempty"`
	StepsRemoved         uint32   `protobuf:"varint,24,opt,name=steps_removed,json=stepsRemoved,proto3" json:"steps_removed,omitempty"`
	ReservedClockId      uint32   `protobuf:"varint,25,opt,name=reserved_clock_id,json=reservedClockId,proto3" json:"reserved_clock_id,omitempty"`
	G8265_1Incompatible  uint32   `protobuf:"varint,26,opt,name=g8265_1_incompatible,json=g82651Incompatible,proto3" json:"g8265_1_incompatible,omitempty"`
	G8275_1Incompatible  uint32   `protobuf:"varint,27,opt,name=g8275_1_incompatible,json=g82751Incompatible,proto3" json:"g8275_1_incompatible,omitempty"`
	G8275_2Incompatible  uint32   `protobuf:"varint,28,opt,name=g8275_2_incompatible,json=g82752Incompatible,proto3" json:"g8275_2_incompatible,omitempty"`
	IncorrectAddress     uint32   `protobuf:"varint,29,opt,name=incorrect_address,json=incorrectAddress,proto3" json:"incorrect_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagDropReasons) Reset()         { *m = PtpBagDropReasons{} }
func (m *PtpBagDropReasons) String() string { return proto.CompactTextString(m) }
func (*PtpBagDropReasons) ProtoMessage()    {}
func (*PtpBagDropReasons) Descriptor() ([]byte, []int) {
	return fileDescriptor_786518142ae33ccc, []int{2}
}

func (m *PtpBagDropReasons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagDropReasons.Unmarshal(m, b)
}
func (m *PtpBagDropReasons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagDropReasons.Marshal(b, m, deterministic)
}
func (m *PtpBagDropReasons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagDropReasons.Merge(m, src)
}
func (m *PtpBagDropReasons) XXX_Size() int {
	return xxx_messageInfo_PtpBagDropReasons.Size(m)
}
func (m *PtpBagDropReasons) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagDropReasons.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagDropReasons proto.InternalMessageInfo

func (m *PtpBagDropReasons) GetNotReady() uint32 {
	if m != nil {
		return m.NotReady
	}
	return 0
}

func (m *PtpBagDropReasons) GetWrongDomain() uint32 {
	if m != nil {
		return m.WrongDomain
	}
	return 0
}

func (m *PtpBagDropReasons) GetTooShort() uint32 {
	if m != nil {
		return m.TooShort
	}
	return 0
}

func (m *PtpBagDropReasons) GetLoopedSamePort() uint32 {
	if m != nil {
		return m.LoopedSamePort
	}
	return 0
}

func (m *PtpBagDropReasons) GetLoopedHigherPort() uint32 {
	if m != nil {
		return m.LoopedHigherPort
	}
	return 0
}

func (m *PtpBagDropReasons) GetLoopedLowerPort() uint32 {
	if m != nil {
		return m.LoopedLowerPort
	}
	return 0
}

func (m *PtpBagDropReasons) GetNoTimestamp() uint32 {
	if m != nil {
		return m.NoTimestamp
	}
	return 0
}

func (m *PtpBagDropReasons) GetZeroTimestamp() uint32 {
	if m != nil {
		return m.ZeroTimestamp
	}
	return 0
}

func (m *PtpBagDropReasons) GetInvalidTlVs() uint32 {
	if m != nil {
		return m.InvalidTlVs
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotForUs() uint32 {
	if m != nil {
		return m.NotForUs
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotListening() uint32 {
	if m != nil {
		return m.NotListening
	}
	return 0
}

func (m *PtpBagDropReasons) GetWrongMaster() uint32 {
	if m != nil {
		return m.WrongMaster
	}
	return 0
}

func (m *PtpBagDropReasons) GetUnknownMaster() uint32 {
	if m != nil {
		return m.UnknownMaster
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotMaster() uint32 {
	if m != nil {
		return m.NotMaster
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotSlave() uint32 {
	if m != nil {
		return m.NotSlave
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotGranted() uint32 {
	if m != nil {
		return m.NotGranted
	}
	return 0
}

func (m *PtpBagDropReasons) GetTooSlow() uint32 {
	if m != nil {
		return m.TooSlow
	}
	return 0
}

func (m *PtpBagDropReasons) GetInvalidPacket() uint32 {
	if m != nil {
		return m.InvalidPacket
	}
	return 0
}

func (m *PtpBagDropReasons) GetWrongSequenceId() uint32 {
	if m != nil {
		return m.WrongSequenceId
	}
	return 0
}

func (m *PtpBagDropReasons) GetNoOffloadSession() uint32 {
	if m != nil {
		return m.NoOffloadSession
	}
	return 0
}

func (m *PtpBagDropReasons) GetNotSupported() uint32 {
	if m != nil {
		return m.NotSupported
	}
	return 0
}

func (m *PtpBagDropReasons) GetMinClockClass() uint32 {
	if m != nil {
		return m.MinClockClass
	}
	return 0
}

func (m *PtpBagDropReasons) GetBadClockClass() uint32 {
	if m != nil {
		return m.BadClockClass
	}
	return 0
}

func (m *PtpBagDropReasons) GetStepsRemoved() uint32 {
	if m != nil {
		return m.StepsRemoved
	}
	return 0
}

func (m *PtpBagDropReasons) GetReservedClockId() uint32 {
	if m != nil {
		return m.ReservedClockId
	}
	return 0
}

func (m *PtpBagDropReasons) GetG8265_1Incompatible() uint32 {
	if m != nil {
		return m.G8265_1Incompatible
	}
	return 0
}

func (m *PtpBagDropReasons) GetG8275_1Incompatible() uint32 {
	if m != nil {
		return m.G8275_1Incompatible
	}
	return 0
}

func (m *PtpBagDropReasons) GetG8275_2Incompatible() uint32 {
	if m != nil {
		return m.G8275_2Incompatible
	}
	return 0
}

func (m *PtpBagDropReasons) GetIncorrectAddress() uint32 {
	if m != nil {
		return m.IncorrectAddress
	}
	return 0
}

type PtpNodePacketCounterInfo struct {
	Counters             *PtpBagPacketCounters `protobuf:"bytes,50,opt,name=counters,proto3" json:"counters,omitempty"`
	DropReasons          *PtpBagDropReasons    `protobuf:"bytes,51,opt,name=drop_reasons,json=dropReasons,proto3" json:"drop_reasons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpNodePacketCounterInfo) Reset()         { *m = PtpNodePacketCounterInfo{} }
func (m *PtpNodePacketCounterInfo) String() string { return proto.CompactTextString(m) }
func (*PtpNodePacketCounterInfo) ProtoMessage()    {}
func (*PtpNodePacketCounterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_786518142ae33ccc, []int{3}
}

func (m *PtpNodePacketCounterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpNodePacketCounterInfo.Unmarshal(m, b)
}
func (m *PtpNodePacketCounterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpNodePacketCounterInfo.Marshal(b, m, deterministic)
}
func (m *PtpNodePacketCounterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpNodePacketCounterInfo.Merge(m, src)
}
func (m *PtpNodePacketCounterInfo) XXX_Size() int {
	return xxx_messageInfo_PtpNodePacketCounterInfo.Size(m)
}
func (m *PtpNodePacketCounterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpNodePacketCounterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpNodePacketCounterInfo proto.InternalMessageInfo

func (m *PtpNodePacketCounterInfo) GetCounters() *PtpBagPacketCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *PtpNodePacketCounterInfo) GetDropReasons() *PtpBagDropReasons {
	if m != nil {
		return m.DropReasons
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpNodePacketCounterInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.packet_counters.ptp_node_packet_counter_info_KEYS")
	proto.RegisterType((*PtpBagPacketCounters)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.packet_counters.ptp_bag_packet_counters")
	proto.RegisterType((*PtpBagDropReasons)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.packet_counters.ptp_bag_drop_reasons")
	proto.RegisterType((*PtpNodePacketCounterInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.packet_counters.ptp_node_packet_counter_info")
}

func init() { proto.RegisterFile("ptp_node_packet_counter_info.proto", fileDescriptor_786518142ae33ccc) }

var fileDescriptor_786518142ae33ccc = []byte{
	// 1248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0x5b, 0x93, 0x1b, 0x35,
	0x13, 0xad, 0xfd, 0x3e, 0x48, 0x9c, 0xf6, 0x6d, 0x3d, 0xbb, 0xc9, 0x3a, 0xd9, 0x0d, 0xd9, 0x75,
	0x2e, 0x84, 0x5c, 0x0c, 0x71, 0x6e, 0x54, 0x51, 0x14, 0x50, 0x09, 0x81, 0x85, 0x84, 0x04, 0x3b,
	0xa1, 0x8a, 0x27, 0x95, 0x76, 0x46, 0xeb, 0x4c, 0x65, 0x46, 0x1a, 0x46, 0xb2, 0x4d, 0xf8, 0x85,
	0xfc, 0x20, 0x1e, 0x78, 0xa4, 0xd4, 0x2d, 0xc9, 0x33, 0xde, 0xcb, 0x03, 0xbc, 0xb8, 0x6a, 0x4e,
	0x9f, 0xd3, 0x3a, 0xdd, 0x1a, 0xa9, 0x3d, 0x30, 0x28, 0x4c, 0xc1, 0xa4, 0x4a, 0x04, 0x2b, 0x78,
	0xfc, 0x4e, 0x18, 0x16, 0xab, 0x99, 0x34, 0xa2, 0x64, 0xa9, 0x3c, 0x54, 0xc3, 0xa2, 0x54, 0x46,
	0x45, 0x0f, 0xe2, 0x54, 0xc7, 0x8a, 0xa5, 0x4a, 0xb3, 0xdf, 0x4b, 0x66, 0x05, 0xaa, 0x10, 0xe5,
	0xb0, 0x30, 0xc5, 0xd0, 0x2a, 0x35, 0xfe, 0x0e, 0xeb, 0x7a, 0x3d, 0xf8, 0x1a, 0xf6, 0x4e, 0xcb,
	0xcd, 0x7e, 0xfc, 0xf6, 0xd7, 0x49, 0xb4, 0x0d, 0xe7, 0x90, 0x20, 0x79, 0x2e, 0xfa, 0x6b, 0xbb,
	0x6b, 0x37, 0xcf, 0x8d, 0x1b, 0x16, 0xf8, 0x89, 0xe7, 0x62, 0xf0, 0x67, 0x07, 0xb6, 0x6c, 0x8a,
	0x03, 0x3e, 0x5d, 0xc9, 0xa0, 0xa3, 0xab, 0xd0, 0xe6, 0x52, 0xaa, 0x99, 0x8c, 0x05, 0xd3, 0x42,
	0x1a, 0x14, 0xb7, 0xc7, 0x2d, 0x0f, 0x4e, 0x84, 0x34, 0xd1, 0x6d, 0xe8, 0x05, 0x52, 0x29, 0x62,
	0x91, 0xce, 0x45, 0xd2, 0xff, 0x1f, 0x12, 0xd7, 0x7d, 0x60, 0xec, 0xf0, 0xe8, 0x13, 0x08, 0x18,
	0x4b, 0x4a, 0x55, 0x14, 0x22, 0xe9, 0xff, 0x1f, 0xb9, 0x5d, 0x8f, 0x3f, 0x25, 0xd8, 0xba, 0xd6,
	0xef, 0x65, 0x4c, 0x0b, 0x7f, 0x80, 0x9c, 0x86, 0x05, 0x70, 0xd1, 0xab, 0xd0, 0xc6, 0x60, 0x58,
	0xf0, 0x43, 0x72, 0x66, 0xc1, 0xb0, 0xd8, 0x1e, 0xe0, 0x73, 0x58, 0xe8, 0x0c, 0x72, 0x9a, 0x16,
	0xf3, 0x8b, 0x5c, 0x83, 0xce, 0xa1, 0xca, 0x32, 0xb5, 0x60, 0xb3, 0x82, 0x56, 0x3a, 0x4b, 0x89,
	0x08, 0x7d, 0x53, 0xe0, 0x6a, 0x77, 0x20, 0x5a, 0xb2, 0xc2, 0x92, 0x0d, 0xaa, 0xd1, 0x33, 0xc3,
	0xb2, 0xb7, 0xa0, 0xb7, 0x64, 0xfb, 0xb5, 0xcf, 0x51, 0x91, 0x9e, 0xec, 0xd7, 0xbf, 0x03, 0x51,
	0x22, 0x32, 0xfe, 0x9e, 0x95, 0xe2, 0xb7, 0x99, 0xd0, 0x86, 0x3c, 0x00, 0x65, 0xc6, 0xc8, 0x98,
	0x02, 0xe8, 0xe3, 0x01, 0x5c, 0xa8, 0xb3, 0x83, 0x97, 0x26, 0x2a, 0x36, 0xab, 0x8a, 0xe0, 0x67,
	0x04, 0xe7, 0xeb, 0x2a, 0xef, 0xa9, 0x85, 0xa2, 0x8d, 0xaa, 0xc8, 0xfb, 0x1a, 0xc2, 0x86, 0xd7,
	0xe8, 0x42, 0x49, 0xed, 0xf6, 0xbf, 0x8d, 0x8a, 0x9e, 0x53, 0x50, 0x04, 0x9d, 0x3d, 0x82, 0xad,
	0x15, 0x7e, 0xb0, 0xd6, 0x41, 0xcd, 0xf9, 0x9a, 0x26, 0x78, 0xab, 0x54, 0xe4, 0x74, 0xde, 0x5c,
	0xb7, 0x56, 0x11, 0x05, 0xbd, 0xbb, 0x87, 0xb0, 0x55, 0x08, 0x51, 0xb2, 0x63, 0x5a, 0xb7, 0x4e,
	0x32, 0x1b, 0x7e, 0xba, 0xda, 0xbe, 0x2f, 0x61, 0xfb, 0x18, 0x59, 0x30, 0xda, 0x43, 0x69, 0x7f,
	0x55, 0x1a, 0xbc, 0x7e, 0x01, 0x97, 0x8e, 0x91, 0x7b, 0xbf, 0x11, 0xaa, 0xb7, 0x56, 0xd5, 0xde,
	0xf2, 0x63, 0xe8, 0xd7, 0xc4, 0xd5, 0xae, 0x6e, 0x50, 0x87, 0x2a, 0xd2, 0x4a, 0x67, 0xbf, 0x82,
	0x9d, 0xe3, 0x84, 0xc1, 0xf5, 0x26, 0x8a, 0x2f, 0x1e, 0x11, 0x07, 0xdb, 0xab, 0x55, 0xaf, 0xf4,
	0xf9, 0xfc, 0x91, 0xaa, 0xeb, 0xbd, 0xde, 0x87, 0xc1, 0x71, 0xf2, 0x95, 0x53, 0x73, 0x01, 0xb3,
	0x5c, 0x3e, 0x92, 0xe5, 0x59, 0xf5, 0x18, 0xfd, 0x0c, 0x37, 0x4e, 0x4f, 0x15, 0x8a, 0xda, 0xc2,
	0x74, 0x7b, 0x27, 0xa6, 0x0b, 0xc5, 0xbd, 0x84, 0xeb, 0xa7, 0xa7, 0xf4, 0x65, 0xf6, 0x31, 0xe3,
	0xee, 0x89, 0x19, 0x7d, 0xb9, 0xd7, 0xa1, 0xa3, 0xd3, 0xa9, 0xe4, 0x59, 0x2a, 0xa7, 0x54, 0xda,
	0x45, 0x54, 0xb6, 0x03, 0x8a, 0xa5, 0xdc, 0x85, 0x68, 0x49, 0x0b, 0xb6, 0x2f, 0xd1, 0xf1, 0x08,
	0x91, 0x60, 0xf3, 0x36, 0x2c, 0xc1, 0x60, 0x69, 0x9b, 0x4e, 0x79, 0x08, 0x78, 0x0b, 0x1f, 0x43,
	0x37, 0xe7, 0x92, 0x4f, 0x45, 0x2e, 0xa4, 0x7b, 0xab, 0x77, 0x90, 0xda, 0x59, 0xc2, 0x68, 0xe2,
	0x53, 0xd8, 0xa8, 0x10, 0x83, 0x8b, 0xcb, 0x48, 0x8e, 0x96, 0xa1, 0x60, 0xe3, 0x2e, 0x54, 0xd0,
	0xe0, 0xe3, 0x23, 0x72, 0xbd, 0x8c, 0x54, 0x2e, 0x27, 0x65, 0xde, 0x8a, 0xd2, 0xcd, 0x05, 0x4d,
	0x5e, 0xae, 0x90, 0x6d, 0x8c, 0xbc, 0xa2, 0x80, 0xbf, 0x9c, 0xea, 0xec, 0x60, 0x68, 0x97, 0xce,
	0x64, 0x55, 0x51, 0xbd, 0x9c, 0xea, 0x2a, 0xef, 0x6a, 0x8f, 0x2e, 0xa7, 0xaa, 0xa8, 0xe2, 0xcb,
	0x28, 0xc3, 0xb3, 0xba, 0xaf, 0x01, 0xf9, 0xc2, 0xc8, 0x8a, 0xaf, 0x3a, 0x3b, 0xf8, 0xba, 0x4a,
	0xbe, 0xaa, 0x8a, 0xaa, 0xaf, 0xba, 0xca, 0xfb, 0xba, 0x46, 0xbe, 0xaa, 0x22, 0xe7, 0x6b, 0xf0,
	0x57, 0x03, 0x36, 0xfd, 0x28, 0xb5, 0x74, 0x56, 0x0a, 0xae, 0x95, 0xd4, 0x34, 0x80, 0xed, 0x0e,
	0xf1, 0xe4, 0xbd, 0x9b, 0xa1, 0x0d, 0xa9, 0xcc, 0xd8, 0x3e, 0xdb, 0x29, 0xb5, 0x28, 0x95, 0x7d,
	0x2f, 0x54, 0xce, 0x53, 0xe9, 0x46, 0x67, 0x13, 0xb1, 0xa7, 0x08, 0x59, 0xbd, 0x51, 0x8a, 0xe9,
	0xb7, 0xaa, 0x34, 0x6e, 0x5c, 0x36, 0x8c, 0x52, 0x13, 0xfb, 0x1c, 0xdd, 0x84, 0xf5, 0x4c, 0xa9,
	0x42, 0x24, 0x4c, 0xf3, 0x5c, 0xb0, 0xc2, 0x72, 0x68, 0x5c, 0x76, 0x08, 0x9f, 0xf0, 0x5c, 0xbc,
	0xb2, 0xcc, 0x3b, 0x10, 0x39, 0xe6, 0xdb, 0x74, 0x8a, 0x3d, 0xb7, 0x5c, 0x9a, 0x9c, 0x2e, 0xc7,
	0xf7, 0x18, 0x40, 0xf6, 0x2d, 0xe8, 0x39, 0x76, 0xa6, 0x16, 0x9e, 0x4c, 0x23, 0xb4, 0x4b, 0x81,
	0xe7, 0x16, 0x47, 0xee, 0x1e, 0xb4, 0xa4, 0x62, 0x26, 0xcd, 0x85, 0x36, 0x3c, 0x2f, 0xdc, 0x10,
	0x6d, 0x4a, 0xf5, 0xda, 0x43, 0xf6, 0x60, 0xfd, 0x21, 0xca, 0x2a, 0x89, 0xe6, 0x67, 0xdb, 0xa2,
	0x4b, 0xda, 0x00, 0xda, 0xa9, 0x9c, 0xf3, 0x2c, 0x4d, 0x98, 0xc9, 0xd8, 0x5c, 0xbb, 0xc1, 0xd9,
	0x74, 0xe0, 0xeb, 0xec, 0x17, 0x1d, 0xed, 0x00, 0xd8, 0x76, 0x1e, 0xaa, 0x92, 0xcd, 0xb4, 0x1b,
	0x96, 0xb6, 0x9f, 0xcf, 0x54, 0xf9, 0x06, 0xff, 0xb4, 0xd8, 0x68, 0x96, 0x6a, 0x23, 0x64, 0x2a,
	0xa7, 0x6e, 0x36, 0xb6, 0xa4, 0x32, 0xcf, 0x3d, 0xb6, 0x6c, 0x7a, 0xce, 0xb5, 0x11, 0xa5, 0x1b,
	0x85, 0xd4, 0xf4, 0x17, 0x08, 0x59, 0xc3, 0x33, 0xf9, 0x4e, 0xaa, 0x85, 0xf4, 0x24, 0x9a, 0x7e,
	0x6d, 0x87, 0x3a, 0xda, 0x65, 0x32, 0xe3, 0x28, 0x34, 0xec, 0xec, 0x6e, 0xbb, 0xb0, 0xdb, 0x7a,
	0x9d, 0xf1, 0xb9, 0x70, 0x33, 0xcd, 0x5a, 0x9d, 0xd8, 0xe7, 0xe8, 0x0a, 0x34, 0x6d, 0x70, 0x5a,
	0x72, 0x69, 0x44, 0xe2, 0x66, 0x97, 0x4d, 0xf7, 0x1d, 0x21, 0xd1, 0x45, 0x68, 0xe0, 0xc6, 0x67,
	0x6a, 0xe1, 0xc6, 0xd3, 0x59, 0xbb, 0xef, 0x99, 0x5a, 0x58, 0x7b, 0xbe, 0x51, 0xf4, 0x8a, 0xba,
	0x09, 0xe4, 0xdb, 0x47, 0xef, 0xa6, 0xdd, 0x45, 0x2a, 0x54, 0xdb, 0x79, 0x64, 0xff, 0x76, 0xa5,
	0x89, 0x1b, 0x38, 0x5d, 0x0c, 0x4c, 0x1c, 0xbe, 0x8f, 0xe7, 0x4a, 0x2a, 0xa6, 0x0e, 0x0f, 0x33,
	0xc5, 0x13, 0xa6, 0x85, 0xd6, 0xa9, 0x92, 0x6e, 0xc0, 0xac, 0x4b, 0xf5, 0x92, 0x02, 0x13, 0xc2,
	0x7d, 0x9f, 0xf5, 0xac, 0xb0, 0x6f, 0x46, 0x98, 0x24, 0xb6, 0xcf, 0x13, 0x8f, 0x45, 0x37, 0xa0,
	0x9b, 0xa7, 0x92, 0xc5, 0x99, 0x8a, 0xdf, 0xb1, 0x38, 0xe3, 0x5a, 0xbb, 0x51, 0xd1, 0xce, 0x53,
	0xf9, 0xc4, 0xa2, 0x4f, 0x2c, 0x68, 0x79, 0x07, 0x3c, 0xa9, 0xf1, 0x68, 0x06, 0xb4, 0x0f, 0x78,
	0x52, 0xe1, 0xd9, 0xff, 0x7d, 0x46, 0x14, 0xf6, 0x10, 0xe7, 0x6a, 0x1e, 0xee, 0xf5, 0x16, 0x82,
	0x63, 0xc2, 0x6c, 0xcd, 0xa5, 0xd0, 0xa2, 0x9c, 0x0b, 0x9f, 0x31, 0x4d, 0xdc, 0x35, 0xde, 0xf5,
	0x01, 0xcc, 0xb9, 0x9f, 0x44, 0x9f, 0xc1, 0xe6, 0xf4, 0xf3, 0xd1, 0xa3, 0x87, 0xec, 0x1e, 0x4b,
	0x65, 0xac, 0xf2, 0x82, 0x9b, 0xf4, 0x20, 0x13, 0xee, 0x2a, 0x8f, 0x30, 0x76, 0x6f, 0xbf, 0x12,
	0x71, 0x8a, 0xc7, 0x47, 0x14, 0xdb, 0x41, 0xf1, 0xf8, 0x24, 0xc5, 0xa8, 0xae, 0xd8, 0xa9, 0x28,
	0x46, 0x35, 0xc5, 0x6d, 0xe8, 0x59, 0x66, 0x59, 0x8a, 0xd8, 0x30, 0x9e, 0x24, 0xa5, 0xd0, 0xda,
	0xdd, 0xeb, 0xeb, 0x21, 0xf0, 0x0d, 0xe1, 0x83, 0xbf, 0xd7, 0x60, 0xe7, 0xb4, 0x8f, 0x80, 0x28,
	0x85, 0x86, 0xff, 0x4b, 0xdf, 0x1f, 0xed, 0xae, 0xdd, 0x6c, 0x8e, 0x5e, 0x0c, 0xff, 0xcd, 0xd7,
	0xc6, 0xf0, 0x84, 0xef, 0x84, 0x71, 0x48, 0x1f, 0xe5, 0xd0, 0xaa, 0xde, 0x7c, 0xfd, 0xfb, 0xb8,
	0xdc, 0x0f, 0xff, 0x6d, 0xb9, 0x6a, 0xc6, 0x71, 0xd3, 0x3e, 0x8d, 0xe9, 0xe1, 0xe0, 0x0c, 0x7e,
	0x3b, 0xdd, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xf3, 0xf1, 0xbc, 0x61, 0x0d, 0x00, 0x00,
}
