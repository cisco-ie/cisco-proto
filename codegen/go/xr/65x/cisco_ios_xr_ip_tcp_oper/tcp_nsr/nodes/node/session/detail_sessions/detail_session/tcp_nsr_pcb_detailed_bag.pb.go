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
// source: tcp_nsr_pcb_detailed_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_nsr_nodes_node_session_detail_sessions_detail_session

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

type TcpNsrPcbDetailedBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 string   `protobuf:"bytes,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrPcbDetailedBag_KEYS) Reset()         { *m = TcpNsrPcbDetailedBag_KEYS{} }
func (m *TcpNsrPcbDetailedBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbDetailedBag_KEYS) ProtoMessage()    {}
func (*TcpNsrPcbDetailedBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fd9544ea9546ece, []int{0}
}

func (m *TcpNsrPcbDetailedBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS.Unmarshal(m, b)
}
func (m *TcpNsrPcbDetailedBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbDetailedBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS.Merge(m, src)
}
func (m *TcpNsrPcbDetailedBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS.Size(m)
}
func (m *TcpNsrPcbDetailedBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbDetailedBag_KEYS proto.InternalMessageInfo

func (m *TcpNsrPcbDetailedBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag_KEYS) GetId_1() string {
	if m != nil {
		return m.Id_1
	}
	return ""
}

type TcpNsrSscbBriefBag struct {
	Sscb                             uint64   `protobuf:"varint,1,opt,name=sscb,proto3" json:"sscb,omitempty"`
	Pid                              uint32   `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	ClientName                       []string `protobuf:"bytes,3,rep,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	ClientInstance                   uint32   `protobuf:"varint,4,opt,name=client_instance,json=clientInstance,proto3" json:"client_instance,omitempty"`
	SetId                            uint32   `protobuf:"varint,5,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty"`
	SsoRole                          uint32   `protobuf:"varint,6,opt,name=sso_role,json=ssoRole,proto3" json:"sso_role,omitempty"`
	Mode                             uint32   `protobuf:"varint,7,opt,name=mode,proto3" json:"mode,omitempty"`
	AddressFamily                    string   `protobuf:"bytes,8,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	WellKnownPort                    uint32   `protobuf:"varint,9,opt,name=well_known_port,json=wellKnownPort,proto3" json:"well_known_port,omitempty"`
	LocalNode                        string   `protobuf:"bytes,10,opt,name=local_node,json=localNode,proto3" json:"local_node,omitempty"`
	LocalInstance                    uint32   `protobuf:"varint,11,opt,name=local_instance,json=localInstance,proto3" json:"local_instance,omitempty"`
	ProtectNode                      string   `protobuf:"bytes,12,opt,name=protect_node,json=protectNode,proto3" json:"protect_node,omitempty"`
	ProtectInstance                  uint32   `protobuf:"varint,13,opt,name=protect_instance,json=protectInstance,proto3" json:"protect_instance,omitempty"`
	NumberOfSessions                 uint32   `protobuf:"varint,14,opt,name=number_of_sessions,json=numberOfSessions,proto3" json:"number_of_sessions,omitempty"`
	NumberOfSyncedSessionsUpStream   uint32   `protobuf:"varint,15,opt,name=number_of_synced_sessions_up_stream,json=numberOfSyncedSessionsUpStream,proto3" json:"number_of_synced_sessions_up_stream,omitempty"`
	NumberOfSyncedSessionsDownStream uint32   `protobuf:"varint,16,opt,name=number_of_synced_sessions_down_stream,json=numberOfSyncedSessionsDownStream,proto3" json:"number_of_synced_sessions_down_stream,omitempty"`
	IsInitSyncInProgress             bool     `protobuf:"varint,17,opt,name=is_init_sync_in_progress,json=isInitSyncInProgress,proto3" json:"is_init_sync_in_progress,omitempty"`
	IsSscbInitSyncReady              bool     `protobuf:"varint,18,opt,name=is_sscb_init_sync_ready,json=isSscbInitSyncReady,proto3" json:"is_sscb_init_sync_ready,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *TcpNsrSscbBriefBag) Reset()         { *m = TcpNsrSscbBriefBag{} }
func (m *TcpNsrSscbBriefBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSscbBriefBag) ProtoMessage()    {}
func (*TcpNsrSscbBriefBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fd9544ea9546ece, []int{1}
}

func (m *TcpNsrSscbBriefBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Unmarshal(m, b)
}
func (m *TcpNsrSscbBriefBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrSscbBriefBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSscbBriefBag.Merge(m, src)
}
func (m *TcpNsrSscbBriefBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSscbBriefBag.Size(m)
}
func (m *TcpNsrSscbBriefBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSscbBriefBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSscbBriefBag proto.InternalMessageInfo

func (m *TcpNsrSscbBriefBag) GetSscb() uint64 {
	if m != nil {
		return m.Sscb
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetClientName() []string {
	if m != nil {
		return m.ClientName
	}
	return nil
}

func (m *TcpNsrSscbBriefBag) GetClientInstance() uint32 {
	if m != nil {
		return m.ClientInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetSetId() uint32 {
	if m != nil {
		return m.SetId
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetSsoRole() uint32 {
	if m != nil {
		return m.SsoRole
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetWellKnownPort() uint32 {
	if m != nil {
		return m.WellKnownPort
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetLocalNode() string {
	if m != nil {
		return m.LocalNode
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetLocalInstance() uint32 {
	if m != nil {
		return m.LocalInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetProtectNode() string {
	if m != nil {
		return m.ProtectNode
	}
	return ""
}

func (m *TcpNsrSscbBriefBag) GetProtectInstance() uint32 {
	if m != nil {
		return m.ProtectInstance
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSessions() uint32 {
	if m != nil {
		return m.NumberOfSessions
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSyncedSessionsUpStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsUpStream
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetNumberOfSyncedSessionsDownStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsDownStream
	}
	return 0
}

func (m *TcpNsrSscbBriefBag) GetIsInitSyncInProgress() bool {
	if m != nil {
		return m.IsInitSyncInProgress
	}
	return false
}

func (m *TcpNsrSscbBriefBag) GetIsSscbInitSyncReady() bool {
	if m != nil {
		return m.IsSscbInitSyncReady
	}
	return false
}

type TcpNsrHoldQueueNode struct {
	SequenceNumber       uint32   `protobuf:"varint,1,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	DataLength           uint32   `protobuf:"varint,2,opt,name=data_length,json=dataLength,proto3" json:"data_length,omitempty"`
	AcknoledgementNumber uint32   `protobuf:"varint,3,opt,name=acknoledgement_number,json=acknoledgementNumber,proto3" json:"acknoledgement_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrHoldQueueNode) Reset()         { *m = TcpNsrHoldQueueNode{} }
func (m *TcpNsrHoldQueueNode) String() string { return proto.CompactTextString(m) }
func (*TcpNsrHoldQueueNode) ProtoMessage()    {}
func (*TcpNsrHoldQueueNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fd9544ea9546ece, []int{2}
}

func (m *TcpNsrHoldQueueNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrHoldQueueNode.Unmarshal(m, b)
}
func (m *TcpNsrHoldQueueNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrHoldQueueNode.Marshal(b, m, deterministic)
}
func (m *TcpNsrHoldQueueNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrHoldQueueNode.Merge(m, src)
}
func (m *TcpNsrHoldQueueNode) XXX_Size() int {
	return xxx_messageInfo_TcpNsrHoldQueueNode.Size(m)
}
func (m *TcpNsrHoldQueueNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrHoldQueueNode.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrHoldQueueNode proto.InternalMessageInfo

func (m *TcpNsrHoldQueueNode) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *TcpNsrHoldQueueNode) GetDataLength() uint32 {
	if m != nil {
		return m.DataLength
	}
	return 0
}

func (m *TcpNsrHoldQueueNode) GetAcknoledgementNumber() uint32 {
	if m != nil {
		return m.AcknoledgementNumber
	}
	return 0
}

type TcpNsrPcbDetailedBag struct {
	AddressFamily                       string                 `protobuf:"bytes,50,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	Pcb                                 uint64                 `protobuf:"varint,51,opt,name=pcb,proto3" json:"pcb,omitempty"`
	Sscb                                uint64                 `protobuf:"varint,52,opt,name=sscb,proto3" json:"sscb,omitempty"`
	LocalAddress                        []string               `protobuf:"bytes,53,rep,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ForeignAddress                      []string               `protobuf:"bytes,54,rep,name=foreign_address,json=foreignAddress,proto3" json:"foreign_address,omitempty"`
	LocalPort                           uint32                 `protobuf:"varint,55,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	ForeignPort                         uint32                 `protobuf:"varint,56,opt,name=foreign_port,json=foreignPort,proto3" json:"foreign_port,omitempty"`
	VrfId                               uint32                 `protobuf:"varint,57,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	IsAdminConfiguredUp                 bool                   `protobuf:"varint,58,opt,name=is_admin_configured_up,json=isAdminConfiguredUp,proto3" json:"is_admin_configured_up,omitempty"`
	IsUsOperationalUp                   string                 `protobuf:"bytes,59,opt,name=is_us_operational_up,json=isUsOperationalUp,proto3" json:"is_us_operational_up,omitempty"`
	IsDsOperationalUp                   string                 `protobuf:"bytes,60,opt,name=is_ds_operational_up,json=isDsOperationalUp,proto3" json:"is_ds_operational_up,omitempty"`
	IsOnlyReceivePathReplication        bool                   `protobuf:"varint,61,opt,name=is_only_receive_path_replication,json=isOnlyReceivePathReplication,proto3" json:"is_only_receive_path_replication,omitempty"`
	Cookie                              uint64                 `protobuf:"varint,62,opt,name=cookie,proto3" json:"cookie,omitempty"`
	SetInformation                      *TcpNsrSscbBriefBag    `protobuf:"bytes,63,opt,name=set_information,json=setInformation,proto3" json:"set_information,omitempty"`
	IsSessionReplicated                 bool                   `protobuf:"varint,64,opt,name=is_session_replicated,json=isSessionReplicated,proto3" json:"is_session_replicated,omitempty"`
	IsSessionSynced                     bool                   `protobuf:"varint,65,opt,name=is_session_synced,json=isSessionSynced,proto3" json:"is_session_synced,omitempty"`
	FistStandbySequenceNumber           uint32                 `protobuf:"varint,66,opt,name=fist_standby_sequence_number,json=fistStandbySequenceNumber,proto3" json:"fist_standby_sequence_number,omitempty"`
	FssnOffset                          uint32                 `protobuf:"varint,67,opt,name=fssn_offset,json=fssnOffset,proto3" json:"fssn_offset,omitempty"`
	NsrDownReason                       string                 `protobuf:"bytes,68,opt,name=nsr_down_reason,json=nsrDownReason,proto3" json:"nsr_down_reason,omitempty"`
	NsrDownTime                         uint32                 `protobuf:"varint,69,opt,name=nsr_down_time,json=nsrDownTime,proto3" json:"nsr_down_time,omitempty"`
	SequenceNumberOfInitSync            uint32                 `protobuf:"varint,70,opt,name=sequence_number_of_init_sync,json=sequenceNumberOfInitSync,proto3" json:"sequence_number_of_init_sync,omitempty"`
	IsInitSyncInProgress                bool                   `protobuf:"varint,71,opt,name=is_init_sync_in_progress,json=isInitSyncInProgress,proto3" json:"is_init_sync_in_progress,omitempty"`
	IsInitSyncSecondPhase               bool                   `protobuf:"varint,72,opt,name=is_init_sync_second_phase,json=isInitSyncSecondPhase,proto3" json:"is_init_sync_second_phase,omitempty"`
	InitSyncError                       []string               `protobuf:"bytes,73,rep,name=init_sync_error,json=initSyncError,proto3" json:"init_sync_error,omitempty"`
	IsInitSyncErrorLocal                bool                   `protobuf:"varint,74,opt,name=is_init_sync_error_local,json=isInitSyncErrorLocal,proto3" json:"is_init_sync_error_local,omitempty"`
	InitSyncStartTime                   uint32                 `protobuf:"varint,75,opt,name=init_sync_start_time,json=initSyncStartTime,proto3" json:"init_sync_start_time,omitempty"`
	InitSyncEndTime                     uint32                 `protobuf:"varint,76,opt,name=init_sync_end_time,json=initSyncEndTime,proto3" json:"init_sync_end_time,omitempty"`
	InitSyncFlags                       uint32                 `protobuf:"varint,77,opt,name=init_sync_flags,json=initSyncFlags,proto3" json:"init_sync_flags,omitempty"`
	SequenceNumberOfInitSyncUpStream    uint32                 `protobuf:"varint,78,opt,name=sequence_number_of_init_sync_up_stream,json=sequenceNumberOfInitSyncUpStream,proto3" json:"sequence_number_of_init_sync_up_stream,omitempty"`
	PeerEndpHdlUpStream                 uint64                 `protobuf:"varint,79,opt,name=peer_endp_hdl_up_stream,json=peerEndpHdlUpStream,proto3" json:"peer_endp_hdl_up_stream,omitempty"`
	InitSyncStartTimeUpStream           uint32                 `protobuf:"varint,80,opt,name=init_sync_start_time_up_stream,json=initSyncStartTimeUpStream,proto3" json:"init_sync_start_time_up_stream,omitempty"`
	InitSyncEndTimeUpStream             uint32                 `protobuf:"varint,81,opt,name=init_sync_end_time_up_stream,json=initSyncEndTimeUpStream,proto3" json:"init_sync_end_time_up_stream,omitempty"`
	FistStandbySequenceNumberUpStream   uint32                 `protobuf:"varint,82,opt,name=fist_standby_sequence_number_up_stream,json=fistStandbySequenceNumberUpStream,proto3" json:"fist_standby_sequence_number_up_stream,omitempty"`
	NsrDownReasonUpStream               string                 `protobuf:"bytes,83,opt,name=nsr_down_reason_up_stream,json=nsrDownReasonUpStream,proto3" json:"nsr_down_reason_up_stream,omitempty"`
	NsrDownTimeUpStream                 uint32                 `protobuf:"varint,84,opt,name=nsr_down_time_up_stream,json=nsrDownTimeUpStream,proto3" json:"nsr_down_time_up_stream,omitempty"`
	SequenceNumberOfInitSyncDownStream  uint32                 `protobuf:"varint,85,opt,name=sequence_number_of_init_sync_down_stream,json=sequenceNumberOfInitSyncDownStream,proto3" json:"sequence_number_of_init_sync_down_stream,omitempty"`
	PeerEndpHdlDownStream               uint64                 `protobuf:"varint,86,opt,name=peer_endp_hdl_down_stream,json=peerEndpHdlDownStream,proto3" json:"peer_endp_hdl_down_stream,omitempty"`
	InitSyncStartTimeDownStream         uint32                 `protobuf:"varint,87,opt,name=init_sync_start_time_down_stream,json=initSyncStartTimeDownStream,proto3" json:"init_sync_start_time_down_stream,omitempty"`
	InitSyncEndTimeDownStream           uint32                 `protobuf:"varint,88,opt,name=init_sync_end_time_down_stream,json=initSyncEndTimeDownStream,proto3" json:"init_sync_end_time_down_stream,omitempty"`
	FistStandbySequenceNumberDownStream uint32                 `protobuf:"varint,89,opt,name=fist_standby_sequence_number_down_stream,json=fistStandbySequenceNumberDownStream,proto3" json:"fist_standby_sequence_number_down_stream,omitempty"`
	NsrDownReasonDownStream             string                 `protobuf:"bytes,90,opt,name=nsr_down_reason_down_stream,json=nsrDownReasonDownStream,proto3" json:"nsr_down_reason_down_stream,omitempty"`
	NsrDownTimeDownStream               uint32                 `protobuf:"varint,91,opt,name=nsr_down_time_down_stream,json=nsrDownTimeDownStream,proto3" json:"nsr_down_time_down_stream,omitempty"`
	PacketHoldQueue                     []*TcpNsrHoldQueueNode `protobuf:"bytes,92,rep,name=packet_hold_queue,json=packetHoldQueue,proto3" json:"packet_hold_queue,omitempty"`
	MaxNumberOfHeldPacket               int32                  `protobuf:"zigzag32,93,opt,name=max_number_of_held_packet,json=maxNumberOfHeldPacket,proto3" json:"max_number_of_held_packet,omitempty"`
	MaxNumberOfHeldPacketReachTime      uint32                 `protobuf:"varint,94,opt,name=max_number_of_held_packet_reach_time,json=maxNumberOfHeldPacketReachTime,proto3" json:"max_number_of_held_packet_reach_time,omitempty"`
	InternalAckHoldQueue                []*TcpNsrHoldQueueNode `protobuf:"bytes,95,rep,name=internal_ack_hold_queue,json=internalAckHoldQueue,proto3" json:"internal_ack_hold_queue,omitempty"`
	MaxNumberOfHeldInternalAck          int32                  `protobuf:"zigzag32,96,opt,name=max_number_of_held_internal_ack,json=maxNumberOfHeldInternalAck,proto3" json:"max_number_of_held_internal_ack,omitempty"`
	MaxNumberOfHeldInternalAckReachTime uint32                 `protobuf:"varint,97,opt,name=max_number_of_held_internal_ack_reach_time,json=maxNumberOfHeldInternalAckReachTime,proto3" json:"max_number_of_held_internal_ack_reach_time,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}               `json:"-"`
	XXX_unrecognized                    []byte                 `json:"-"`
	XXX_sizecache                       int32                  `json:"-"`
}

func (m *TcpNsrPcbDetailedBag) Reset()         { *m = TcpNsrPcbDetailedBag{} }
func (m *TcpNsrPcbDetailedBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbDetailedBag) ProtoMessage()    {}
func (*TcpNsrPcbDetailedBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fd9544ea9546ece, []int{3}
}

func (m *TcpNsrPcbDetailedBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbDetailedBag.Unmarshal(m, b)
}
func (m *TcpNsrPcbDetailedBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbDetailedBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbDetailedBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbDetailedBag.Merge(m, src)
}
func (m *TcpNsrPcbDetailedBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbDetailedBag.Size(m)
}
func (m *TcpNsrPcbDetailedBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbDetailedBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbDetailedBag proto.InternalMessageInfo

func (m *TcpNsrPcbDetailedBag) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetPcb() uint64 {
	if m != nil {
		return m.Pcb
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetSscb() uint64 {
	if m != nil {
		return m.Sscb
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetLocalAddress() []string {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetForeignAddress() []string {
	if m != nil {
		return m.ForeignAddress
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetForeignPort() uint32 {
	if m != nil {
		return m.ForeignPort
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetIsAdminConfiguredUp() bool {
	if m != nil {
		return m.IsAdminConfiguredUp
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetIsUsOperationalUp() string {
	if m != nil {
		return m.IsUsOperationalUp
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetIsDsOperationalUp() string {
	if m != nil {
		return m.IsDsOperationalUp
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetIsOnlyReceivePathReplication() bool {
	if m != nil {
		return m.IsOnlyReceivePathReplication
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetCookie() uint64 {
	if m != nil {
		return m.Cookie
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetSetInformation() *TcpNsrSscbBriefBag {
	if m != nil {
		return m.SetInformation
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetIsSessionReplicated() bool {
	if m != nil {
		return m.IsSessionReplicated
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetIsSessionSynced() bool {
	if m != nil {
		return m.IsSessionSynced
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetFistStandbySequenceNumber() uint32 {
	if m != nil {
		return m.FistStandbySequenceNumber
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetFssnOffset() uint32 {
	if m != nil {
		return m.FssnOffset
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownReason() string {
	if m != nil {
		return m.NsrDownReason
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownTime() uint32 {
	if m != nil {
		return m.NsrDownTime
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetSequenceNumberOfInitSync() uint32 {
	if m != nil {
		return m.SequenceNumberOfInitSync
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetIsInitSyncInProgress() bool {
	if m != nil {
		return m.IsInitSyncInProgress
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetIsInitSyncSecondPhase() bool {
	if m != nil {
		return m.IsInitSyncSecondPhase
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncError() []string {
	if m != nil {
		return m.InitSyncError
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetIsInitSyncErrorLocal() bool {
	if m != nil {
		return m.IsInitSyncErrorLocal
	}
	return false
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncStartTime() uint32 {
	if m != nil {
		return m.InitSyncStartTime
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncEndTime() uint32 {
	if m != nil {
		return m.InitSyncEndTime
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncFlags() uint32 {
	if m != nil {
		return m.InitSyncFlags
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetSequenceNumberOfInitSyncUpStream() uint32 {
	if m != nil {
		return m.SequenceNumberOfInitSyncUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetPeerEndpHdlUpStream() uint64 {
	if m != nil {
		return m.PeerEndpHdlUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncStartTimeUpStream() uint32 {
	if m != nil {
		return m.InitSyncStartTimeUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncEndTimeUpStream() uint32 {
	if m != nil {
		return m.InitSyncEndTimeUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetFistStandbySequenceNumberUpStream() uint32 {
	if m != nil {
		return m.FistStandbySequenceNumberUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownReasonUpStream() string {
	if m != nil {
		return m.NsrDownReasonUpStream
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownTimeUpStream() uint32 {
	if m != nil {
		return m.NsrDownTimeUpStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetSequenceNumberOfInitSyncDownStream() uint32 {
	if m != nil {
		return m.SequenceNumberOfInitSyncDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetPeerEndpHdlDownStream() uint64 {
	if m != nil {
		return m.PeerEndpHdlDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncStartTimeDownStream() uint32 {
	if m != nil {
		return m.InitSyncStartTimeDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInitSyncEndTimeDownStream() uint32 {
	if m != nil {
		return m.InitSyncEndTimeDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetFistStandbySequenceNumberDownStream() uint32 {
	if m != nil {
		return m.FistStandbySequenceNumberDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownReasonDownStream() string {
	if m != nil {
		return m.NsrDownReasonDownStream
	}
	return ""
}

func (m *TcpNsrPcbDetailedBag) GetNsrDownTimeDownStream() uint32 {
	if m != nil {
		return m.NsrDownTimeDownStream
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetPacketHoldQueue() []*TcpNsrHoldQueueNode {
	if m != nil {
		return m.PacketHoldQueue
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetMaxNumberOfHeldPacket() int32 {
	if m != nil {
		return m.MaxNumberOfHeldPacket
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetMaxNumberOfHeldPacketReachTime() uint32 {
	if m != nil {
		return m.MaxNumberOfHeldPacketReachTime
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetInternalAckHoldQueue() []*TcpNsrHoldQueueNode {
	if m != nil {
		return m.InternalAckHoldQueue
	}
	return nil
}

func (m *TcpNsrPcbDetailedBag) GetMaxNumberOfHeldInternalAck() int32 {
	if m != nil {
		return m.MaxNumberOfHeldInternalAck
	}
	return 0
}

func (m *TcpNsrPcbDetailedBag) GetMaxNumberOfHeldInternalAckReachTime() uint32 {
	if m != nil {
		return m.MaxNumberOfHeldInternalAckReachTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpNsrPcbDetailedBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.detail_sessions.detail_session.tcp_nsr_pcb_detailed_bag_KEYS")
	proto.RegisterType((*TcpNsrSscbBriefBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.detail_sessions.detail_session.tcp_nsr_sscb_brief_bag")
	proto.RegisterType((*TcpNsrHoldQueueNode)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.detail_sessions.detail_session.tcp_nsr_hold_queue_node")
	proto.RegisterType((*TcpNsrPcbDetailedBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session.detail_sessions.detail_session.tcp_nsr_pcb_detailed_bag")
}

func init() { proto.RegisterFile("tcp_nsr_pcb_detailed_bag.proto", fileDescriptor_2fd9544ea9546ece) }

var fileDescriptor_2fd9544ea9546ece = []byte{
	// 1482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdd, 0x76, 0x1b, 0x35,
	0x10, 0x3e, 0x6e, 0xd2, 0xb4, 0x51, 0xfe, 0xd5, 0x38, 0x51, 0x68, 0xda, 0xba, 0x29, 0x2d, 0xa6,
	0x70, 0xc2, 0x69, 0x52, 0x4a, 0x80, 0xfe, 0x90, 0xb6, 0x09, 0x31, 0x09, 0x49, 0xba, 0x6e, 0x28,
	0xe5, 0x4f, 0xac, 0x57, 0xda, 0x58, 0x78, 0x2d, 0x6d, 0x57, 0xeb, 0xb6, 0x79, 0x0e, 0x2e, 0xe0,
	0x9a, 0xe7, 0xe2, 0x0d, 0x78, 0x09, 0x8e, 0x46, 0xbb, 0x6b, 0xd9, 0xb1, 0xc3, 0x0d, 0x70, 0x93,
	0xb3, 0x9e, 0x99, 0xef, 0xdb, 0xd1, 0xcc, 0x7c, 0xa3, 0x0d, 0xba, 0x9a, 0x06, 0x31, 0x95, 0x3a,
	0xa1, 0x71, 0xd0, 0xa0, 0x8c, 0xa7, 0xbe, 0x88, 0x38, 0xa3, 0x0d, 0xff, 0x78, 0x35, 0x4e, 0x54,
	0xaa, 0xb0, 0x17, 0x08, 0x1d, 0x28, 0x2a, 0x94, 0xa6, 0x6f, 0x13, 0x2a, 0x62, 0x6a, 0xe2, 0x55,
	0xcc, 0x93, 0xd5, 0x0c, 0xb8, 0x2a, 0x15, 0xe3, 0x1a, 0xfe, 0xae, 0x6a, 0xae, 0xb5, 0x50, 0x72,
	0xd5, 0xf2, 0xd0, 0xec, 0xa7, 0xee, 0xfb, 0xbd, 0xf2, 0x18, 0x5d, 0x19, 0xf6, 0x56, 0xba, 0xbb,
	0xf5, 0xb2, 0x8e, 0xa7, 0xd1, 0x39, 0xc1, 0x48, 0xa9, 0x52, 0xaa, 0x8e, 0x7b, 0xe7, 0x04, 0xc3,
	0x73, 0x68, 0x54, 0x30, 0x7a, 0x87, 0x9c, 0x03, 0xcb, 0x88, 0x60, 0x77, 0x56, 0xfe, 0x3a, 0x8f,
	0x16, 0x72, 0x12, 0xad, 0x83, 0x06, 0x6d, 0x24, 0x82, 0x87, 0x86, 0x02, 0x63, 0x34, 0x6a, 0x2c,
	0x80, 0x1f, 0xf5, 0xe0, 0x19, 0xcf, 0xa2, 0x91, 0x58, 0x30, 0x20, 0x98, 0xf2, 0xcc, 0x23, 0xbe,
	0x86, 0x26, 0x82, 0x48, 0x70, 0x99, 0x52, 0xe9, 0xb7, 0x39, 0x19, 0xa9, 0x8c, 0x54, 0xc7, 0x3d,
	0x64, 0x4d, 0xfb, 0x7e, 0x9b, 0xe3, 0xf7, 0xd0, 0x4c, 0x16, 0x20, 0xa4, 0x4e, 0x7d, 0x19, 0x70,
	0x32, 0x0a, 0xf0, 0x69, 0x6b, 0xae, 0x65, 0x56, 0x5c, 0x46, 0x63, 0x9a, 0xa7, 0x54, 0x30, 0x72,
	0x1e, 0xfc, 0xe7, 0x35, 0x4f, 0x6b, 0x0c, 0x2f, 0xa1, 0x8b, 0x5a, 0x2b, 0x9a, 0xa8, 0x88, 0x93,
	0x31, 0x70, 0x5c, 0xd0, 0x5a, 0x79, 0x2a, 0xe2, 0x26, 0xc3, 0xb6, 0x62, 0x9c, 0x5c, 0x00, 0x33,
	0x3c, 0xe3, 0x9b, 0x68, 0xda, 0x67, 0x2c, 0xe1, 0x5a, 0xd3, 0xd0, 0x6f, 0x8b, 0xe8, 0x84, 0x5c,
	0x84, 0xd3, 0x4e, 0x65, 0xd6, 0x6d, 0x30, 0xe2, 0x5b, 0x68, 0xe6, 0x0d, 0x8f, 0x22, 0xda, 0x92,
	0xea, 0x8d, 0xa4, 0xb1, 0x4a, 0x52, 0x32, 0x0e, 0x2c, 0x53, 0xc6, 0xbc, 0x6b, 0xac, 0x87, 0x2a,
	0x49, 0xf1, 0x15, 0x84, 0x22, 0x15, 0xf8, 0x11, 0x35, 0x8d, 0x21, 0x08, 0xa8, 0xc6, 0xc1, 0xb2,
	0x9f, 0xbd, 0xcd, 0xba, 0x8b, 0xb3, 0x4d, 0x58, 0x16, 0xb0, 0x16, 0x47, 0xbb, 0x8e, 0x26, 0xcd,
	0x18, 0xf0, 0x20, 0xb5, 0x3c, 0x93, 0xc0, 0x33, 0x91, 0xd9, 0x80, 0xe9, 0x7d, 0x34, 0x9b, 0x87,
	0x14, 0x5c, 0x53, 0xc0, 0x35, 0x93, 0xd9, 0x0b, 0xb6, 0x0f, 0x11, 0x96, 0x9d, 0x76, 0x83, 0x27,
	0x54, 0x85, 0xc5, 0x70, 0x90, 0x69, 0x08, 0x9e, 0xb5, 0x9e, 0x83, 0xb0, 0x9e, 0xd9, 0xf1, 0x2e,
	0xba, 0xe1, 0x44, 0x9f, 0xc8, 0x80, 0xb3, 0x02, 0x44, 0x3b, 0x31, 0xd5, 0x69, 0xc2, 0xfd, 0x36,
	0x99, 0x01, 0xf8, 0xd5, 0x02, 0x0e, 0x81, 0x39, 0xc9, 0x51, 0x5c, 0x87, 0x28, 0x7c, 0x80, 0x6e,
	0x0e, 0x27, 0x63, 0xa6, 0x94, 0x19, 0xdd, 0x2c, 0xd0, 0x55, 0x06, 0xd3, 0x3d, 0x55, 0x6f, 0x64,
	0x46, 0x78, 0x0f, 0x11, 0xa1, 0xa9, 0x90, 0x22, 0x05, 0x3a, 0x2a, 0x24, 0x8d, 0x13, 0x75, 0x6c,
	0x3a, 0x45, 0xe6, 0x2a, 0xa5, 0xea, 0x45, 0x6f, 0x5e, 0xe8, 0x9a, 0x14, 0xa9, 0x61, 0xa8, 0xc9,
	0xc3, 0xcc, 0x87, 0xef, 0xa2, 0x45, 0xa1, 0xed, 0xc4, 0x76, 0xc1, 0x09, 0xf7, 0xd9, 0x09, 0xc1,
	0x00, 0xbb, 0x24, 0x74, 0x5d, 0x07, 0x8d, 0x1c, 0xea, 0x19, 0xd7, 0xca, 0xef, 0x25, 0xb4, 0x98,
	0x4f, 0x7b, 0x53, 0x45, 0x8c, 0xbe, 0xea, 0xf0, 0x0e, 0x87, 0x9e, 0x98, 0x39, 0xd5, 0xfc, 0x55,
	0x87, 0xcb, 0x80, 0x53, 0x9b, 0x36, 0x4c, 0xfe, 0x94, 0x37, 0x9d, 0x9b, 0xf7, 0xc1, 0x6a, 0x26,
	0x9e, 0xf9, 0xa9, 0x4f, 0x23, 0x2e, 0x8f, 0xd3, 0x66, 0xa6, 0x05, 0x64, 0x4c, 0x7b, 0x60, 0xc1,
	0xeb, 0xa8, 0xec, 0x07, 0x2d, 0xa9, 0x22, 0xce, 0x8e, 0x79, 0x1b, 0xa4, 0x61, 0xf9, 0x46, 0x20,
	0x74, 0xbe, 0xd7, 0x69, 0x59, 0x57, 0xfe, 0x5c, 0x40, 0x64, 0x98, 0x9a, 0x07, 0x0c, 0xf5, 0xda,
	0xa0, 0xa1, 0x36, 0xea, 0x0c, 0x1a, 0x64, 0x1d, 0x04, 0x6b, 0x1e, 0x0b, 0x0d, 0xdf, 0x75, 0x34,
	0x7c, 0x03, 0xd9, 0xe9, 0xa4, 0x19, 0x98, 0x7c, 0x0c, 0x9a, 0x9d, 0x04, 0xe3, 0xa6, 0xb5, 0x99,
	0x6a, 0x84, 0x2a, 0xe1, 0xe2, 0x58, 0x16, 0x61, 0xf7, 0x20, 0x6c, 0x3a, 0x33, 0xe7, 0x81, 0x85,
	0x40, 0x40, 0x43, 0x9f, 0xc0, 0x09, 0xad, 0x40, 0x40, 0x3f, 0xd7, 0xd1, 0x64, 0xce, 0x03, 0x01,
	0x1b, 0x10, 0x30, 0x91, 0xd9, 0x20, 0xa4, 0x8c, 0xc6, 0x5e, 0x27, 0xa1, 0xd1, 0xfd, 0xa7, 0x56,
	0xf7, 0xaf, 0x93, 0xb0, 0xc6, 0xf0, 0x3a, 0x5a, 0x10, 0x9a, 0xfa, 0xac, 0x2d, 0x24, 0x0d, 0x94,
	0x0c, 0xc5, 0x71, 0x27, 0xe1, 0x8c, 0x76, 0x62, 0xf2, 0x59, 0xde, 0xe0, 0x4d, 0xe3, 0x7c, 0x52,
	0xf8, 0x8e, 0x62, 0xfc, 0x11, 0x9a, 0x17, 0x9a, 0x76, 0x34, 0xac, 0x56, 0x3f, 0x15, 0x4a, 0xfa,
	0x91, 0x81, 0x7c, 0x0e, 0xe5, 0x9a, 0x13, 0xfa, 0x48, 0x1f, 0x74, 0x3d, 0x05, 0x80, 0x9d, 0x02,
	0xdc, 0xcf, 0x01, 0x4f, 0xfb, 0x00, 0xdb, 0xa8, 0x22, 0x34, 0x55, 0x32, 0x3a, 0xa1, 0x09, 0x0f,
	0xb8, 0x78, 0xcd, 0x69, 0xec, 0xa7, 0x4d, 0x9a, 0xf0, 0x38, 0x12, 0x01, 0xc4, 0x91, 0x07, 0x90,
	0xe0, 0xb2, 0xd0, 0x07, 0x32, 0x3a, 0xf1, 0x6c, 0xd4, 0xa1, 0x9f, 0x36, 0xbd, 0x6e, 0x0c, 0x5e,
	0x40, 0x63, 0x81, 0x52, 0x2d, 0xc1, 0xc9, 0x43, 0xe8, 0x4d, 0xf6, 0x0b, 0xff, 0x5a, 0x32, 0x73,
	0x68, 0x96, 0x40, 0xa8, 0x92, 0xb6, 0xe5, 0x7b, 0x54, 0x29, 0x55, 0x27, 0xd6, 0x7e, 0x59, 0xfd,
	0xf7, 0xef, 0x90, 0xd5, 0xc1, 0xbb, 0xdf, 0xcc, 0x7c, 0x5a, 0xeb, 0x66, 0x80, 0xd7, 0x50, 0xd9,
	0xc8, 0xcd, 0x82, 0x8a, 0xb3, 0x72, 0x46, 0xbe, 0x28, 0xc4, 0x66, 0x7d, 0x5e, 0xe1, 0xc2, 0xb7,
	0xd1, 0x9c, 0x83, 0xb1, 0xcb, 0x82, 0x6c, 0x42, 0xfc, 0x4c, 0x11, 0x6f, 0x17, 0x03, 0x7e, 0x84,
	0x96, 0x43, 0xa1, 0x53, 0x6a, 0x36, 0x1c, 0x6b, 0x9c, 0xd0, 0x7e, 0x25, 0x3e, 0x86, 0xc9, 0x58,
	0x32, 0x31, 0x75, 0x1b, 0x52, 0x3f, 0x25, 0xca, 0x50, 0x6b, 0x49, 0x55, 0x18, 0x6a, 0x9e, 0x92,
	0x27, 0x56, 0x94, 0xc6, 0x74, 0x00, 0x16, 0xb3, 0xf0, 0xcd, 0x39, 0x61, 0x47, 0x25, 0xdc, 0xd7,
	0x4a, 0x92, 0xa7, 0x56, 0x43, 0x52, 0x27, 0x66, 0x21, 0x79, 0x60, 0xc4, 0x2b, 0x68, 0xaa, 0x88,
	0x4b, 0x45, 0x9b, 0x93, 0x2d, 0x3b, 0xb1, 0x59, 0xd4, 0x73, 0xd1, 0xe6, 0xf8, 0x21, 0x5a, 0xee,
	0x4b, 0xd0, 0xac, 0xc3, 0x62, 0x0f, 0x91, 0x6d, 0x80, 0x90, 0xde, 0xbd, 0x71, 0x10, 0xe6, 0xbb,
	0xe8, 0xcc, 0xa5, 0xf7, 0xe5, 0x19, 0x4b, 0x6f, 0x03, 0x2d, 0xf5, 0xe0, 0x34, 0x0f, 0x94, 0x64,
	0x34, 0x6e, 0xfa, 0x9a, 0x93, 0x1d, 0x00, 0x96, 0xbb, 0xc0, 0x3a, 0x78, 0x0f, 0x8d, 0xd3, 0x9c,
	0xbe, 0x0b, 0xe3, 0x49, 0xa2, 0x12, 0x52, 0x03, 0x39, 0x4f, 0x89, 0x2c, 0x7a, 0xcb, 0x18, 0x4f,
	0x65, 0x06, 0xa1, 0x14, 0xe4, 0x4c, 0xbe, 0xea, 0xcf, 0x0c, 0x20, 0x7b, 0xc6, 0x07, 0x32, 0xea,
	0xa6, 0x95, 0xfa, 0x49, 0x6a, 0x8b, 0xb7, 0x0b, 0x95, 0x98, 0xcb, 0x5f, 0x52, 0x37, 0x1e, 0x28,
	0xe1, 0x07, 0x08, 0x3b, 0x6f, 0x91, 0xcc, 0x86, 0xef, 0xd9, 0x0b, 0xaf, 0xc8, 0x49, 0x32, 0x08,
	0xee, 0xc9, 0x3e, 0x8c, 0xfc, 0x63, 0x4d, 0xbe, 0xb6, 0xd7, 0x6c, 0x1e, 0xb9, 0x6d, 0x8c, 0xf8,
	0x10, 0xdd, 0x3a, 0xab, 0x2f, 0xce, 0x6d, 0xb7, 0x6f, 0xaf, 0xa7, 0x61, 0x1d, 0x2a, 0xee, 0xbb,
	0xbb, 0x68, 0x31, 0xe6, 0x3c, 0x31, 0x19, 0xc6, 0xb4, 0xc9, 0x22, 0x87, 0xe2, 0x00, 0x64, 0x7b,
	0xc9, 0xb8, 0xb7, 0x24, 0x8b, 0x77, 0x58, 0x54, 0xa0, 0x36, 0xd1, 0xd5, 0x41, 0xd5, 0x70, 0xc0,
	0x87, 0x76, 0x9e, 0x4f, 0xd5, 0xa5, 0xa0, 0x78, 0x80, 0x96, 0x4f, 0xd7, 0xc7, 0x21, 0x78, 0x06,
	0x04, 0x8b, 0x7d, 0x95, 0x2a, 0xe0, 0xcf, 0xd0, 0xad, 0xb3, 0xf4, 0xe4, 0x10, 0x79, 0x40, 0x74,
	0x7d, 0xa8, 0xb2, 0x0a, 0xca, 0x0d, 0xb4, 0xd4, 0x27, 0x20, 0x87, 0xa5, 0x0e, 0x52, 0x2a, 0xf7,
	0x48, 0xc9, 0x2d, 0x62, 0x8f, 0xa4, 0x1c, 0xdc, 0x73, 0x78, 0xfb, 0x25, 0x47, 0x5c, 0x05, 0xea,
	0x39, 0xaa, 0x9e, 0xd9, 0x4c, 0xf7, 0x6b, 0xe3, 0x08, 0x68, 0x56, 0x86, 0xb5, 0xd3, 0xf9, 0xde,
	0xd8, 0x40, 0x4b, 0xbd, 0x0d, 0x75, 0x69, 0xbe, 0x81, 0x96, 0x96, 0x9d, 0x96, 0x3a, 0xc8, 0x2d,
	0x54, 0x19, 0xd8, 0x54, 0x97, 0xe0, 0x05, 0xe4, 0x71, 0xf9, 0x54, 0x5b, 0x1d, 0x9a, 0x9e, 0xd9,
	0x28, 0x1a, 0xeb, 0x92, 0x7c, 0xdb, 0x3b, 0x1b, 0x59, 0x6b, 0x1d, 0x8a, 0x23, 0x54, 0x3d, 0xb3,
	0xb9, 0x2e, 0xd9, 0x4b, 0x20, 0xbb, 0x31, 0xb4, 0xbd, 0x0e, 0xed, 0x7d, 0x74, 0xb9, 0xbf, 0xc1,
	0x2e, 0xd3, 0x77, 0xd0, 0xe2, 0xc5, 0x9e, 0x16, 0xf7, 0x16, 0xb6, 0xb7, 0xc9, 0x2e, 0xf6, 0x7b,
	0xc8, 0xa2, 0xec, 0xb4, 0xd9, 0x41, 0xfe, 0x56, 0x42, 0x73, 0xb1, 0x1f, 0xb4, 0x78, 0xea, 0x7c,
	0x93, 0x91, 0x1f, 0x2a, 0x23, 0xd5, 0x89, 0xb5, 0xd6, 0x7f, 0x79, 0xe7, 0xf5, 0x7d, 0x01, 0x7a,
	0x33, 0x36, 0x8b, 0x1d, 0x15, 0xb1, 0x67, 0xc6, 0x6a, 0xce, 0xd4, 0xf6, 0xdf, 0x3a, 0xd3, 0xd7,
	0xe4, 0x11, 0xa3, 0x36, 0x8a, 0xfc, 0x58, 0x29, 0x55, 0xe7, 0xbc, 0x72, 0xdb, 0x7f, 0x9b, 0x8f,
	0xdb, 0x0e, 0x8f, 0xd8, 0x21, 0x38, 0xf1, 0x1e, 0x7a, 0x77, 0x28, 0xd2, 0x14, 0x37, 0x68, 0xda,
	0x85, 0xf7, 0x93, 0xfd, 0xea, 0x1e, 0x48, 0xe2, 0x99, 0x30, 0xd8, 0x7f, 0x7f, 0x94, 0xd0, 0xa2,
	0x90, 0x29, 0x4f, 0xcc, 0xc7, 0x89, 0x1f, 0xb4, 0xdc, 0x3a, 0xd1, 0xff, 0xbf, 0x4e, 0xf3, 0x79,
	0x2e, 0x9b, 0x41, 0xab, 0x5b, 0xac, 0x27, 0xe8, 0xda, 0x80, 0x23, 0xbb, 0x69, 0x93, 0x9f, 0xa1,
	0x64, 0xef, 0xf4, 0x9d, 0xb6, 0xd6, 0x65, 0xc3, 0x2f, 0xd0, 0xed, 0x7f, 0x20, 0x71, 0xab, 0xe7,
	0xdb, 0xe1, 0x1e, 0xce, 0x57, 0x94, 0xb0, 0x31, 0x06, 0xff, 0x86, 0xaf, 0xff, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x5c, 0xe0, 0x68, 0xc3, 0xa8, 0x0f, 0x00, 0x00,
}
