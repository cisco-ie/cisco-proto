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
// source: pce_peer_detail_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_peer_detail_infos_peer_detail_info

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

type PcePeerDetailBag_KEYS struct {
	PeerAddress          string   `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePeerDetailBag_KEYS) Reset()         { *m = PcePeerDetailBag_KEYS{} }
func (m *PcePeerDetailBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PcePeerDetailBag_KEYS) ProtoMessage()    {}
func (*PcePeerDetailBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3afadee1dc9c91, []int{0}
}

func (m *PcePeerDetailBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerDetailBag_KEYS.Unmarshal(m, b)
}
func (m *PcePeerDetailBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerDetailBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PcePeerDetailBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerDetailBag_KEYS.Merge(m, src)
}
func (m *PcePeerDetailBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PcePeerDetailBag_KEYS.Size(m)
}
func (m *PcePeerDetailBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerDetailBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerDetailBag_KEYS proto.InternalMessageInfo

func (m *PcePeerDetailBag_KEYS) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

type PcePcepBag struct {
	PcepState                string   `protobuf:"bytes,1,opt,name=pcep_state,json=pcepState,proto3" json:"pcep_state,omitempty"`
	Stateful                 bool     `protobuf:"varint,2,opt,name=stateful,proto3" json:"stateful,omitempty"`
	CapabilityUpdate         bool     `protobuf:"varint,3,opt,name=capability_update,json=capabilityUpdate,proto3" json:"capability_update,omitempty"`
	CapabilityInstantiate    bool     `protobuf:"varint,4,opt,name=capability_instantiate,json=capabilityInstantiate,proto3" json:"capability_instantiate,omitempty"`
	CapabilitySegmentRouting bool     `protobuf:"varint,5,opt,name=capability_segment_routing,json=capabilitySegmentRouting,proto3" json:"capability_segment_routing,omitempty"`
	CapabilityTriggeredSync  bool     `protobuf:"varint,6,opt,name=capability_triggered_sync,json=capabilityTriggeredSync,proto3" json:"capability_triggered_sync,omitempty"`
	CapabilityDbVersion      bool     `protobuf:"varint,7,opt,name=capability_db_version,json=capabilityDbVersion,proto3" json:"capability_db_version,omitempty"`
	CapabilityDeltaSync      bool     `protobuf:"varint,8,opt,name=capability_delta_sync,json=capabilityDeltaSync,proto3" json:"capability_delta_sync,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *PcePcepBag) Reset()         { *m = PcePcepBag{} }
func (m *PcePcepBag) String() string { return proto.CompactTextString(m) }
func (*PcePcepBag) ProtoMessage()    {}
func (*PcePcepBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3afadee1dc9c91, []int{1}
}

func (m *PcePcepBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePcepBag.Unmarshal(m, b)
}
func (m *PcePcepBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePcepBag.Marshal(b, m, deterministic)
}
func (m *PcePcepBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePcepBag.Merge(m, src)
}
func (m *PcePcepBag) XXX_Size() int {
	return xxx_messageInfo_PcePcepBag.Size(m)
}
func (m *PcePcepBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePcepBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePcepBag proto.InternalMessageInfo

func (m *PcePcepBag) GetPcepState() string {
	if m != nil {
		return m.PcepState
	}
	return ""
}

func (m *PcePcepBag) GetStateful() bool {
	if m != nil {
		return m.Stateful
	}
	return false
}

func (m *PcePcepBag) GetCapabilityUpdate() bool {
	if m != nil {
		return m.CapabilityUpdate
	}
	return false
}

func (m *PcePcepBag) GetCapabilityInstantiate() bool {
	if m != nil {
		return m.CapabilityInstantiate
	}
	return false
}

func (m *PcePcepBag) GetCapabilitySegmentRouting() bool {
	if m != nil {
		return m.CapabilitySegmentRouting
	}
	return false
}

func (m *PcePcepBag) GetCapabilityTriggeredSync() bool {
	if m != nil {
		return m.CapabilityTriggeredSync
	}
	return false
}

func (m *PcePcepBag) GetCapabilityDbVersion() bool {
	if m != nil {
		return m.CapabilityDbVersion
	}
	return false
}

func (m *PcePcepBag) GetCapabilityDeltaSync() bool {
	if m != nil {
		return m.CapabilityDeltaSync
	}
	return false
}

type PcePcepErrorBag struct {
	PcErrorType          uint32   `protobuf:"varint,1,opt,name=pc_error_type,json=pcErrorType,proto3" json:"pc_error_type,omitempty"`
	PcErrorValue         uint32   `protobuf:"varint,2,opt,name=pc_error_value,json=pcErrorValue,proto3" json:"pc_error_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PcePcepErrorBag) Reset()         { *m = PcePcepErrorBag{} }
func (m *PcePcepErrorBag) String() string { return proto.CompactTextString(m) }
func (*PcePcepErrorBag) ProtoMessage()    {}
func (*PcePcepErrorBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3afadee1dc9c91, []int{2}
}

func (m *PcePcepErrorBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePcepErrorBag.Unmarshal(m, b)
}
func (m *PcePcepErrorBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePcepErrorBag.Marshal(b, m, deterministic)
}
func (m *PcePcepErrorBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePcepErrorBag.Merge(m, src)
}
func (m *PcePcepErrorBag) XXX_Size() int {
	return xxx_messageInfo_PcePcepErrorBag.Size(m)
}
func (m *PcePcepErrorBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePcepErrorBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePcepErrorBag proto.InternalMessageInfo

func (m *PcePcepErrorBag) GetPcErrorType() uint32 {
	if m != nil {
		return m.PcErrorType
	}
	return 0
}

func (m *PcePcepErrorBag) GetPcErrorValue() uint32 {
	if m != nil {
		return m.PcErrorValue
	}
	return 0
}

type PcePcepDetailBag struct {
	BriefPcepInformation      *PcePcepBag      `protobuf:"bytes,1,opt,name=brief_pcep_information,json=briefPcepInformation,proto3" json:"brief_pcep_information,omitempty"`
	Error                     string           `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	SpeakerId                 string           `protobuf:"bytes,3,opt,name=speaker_id,json=speakerId,proto3" json:"speaker_id,omitempty"`
	PcepUpTime                uint32           `protobuf:"varint,4,opt,name=pcep_up_time,json=pcepUpTime,proto3" json:"pcep_up_time,omitempty"`
	Keepalives                uint32           `protobuf:"varint,5,opt,name=keepalives,proto3" json:"keepalives,omitempty"`
	Md5Enabled                bool             `protobuf:"varint,6,opt,name=md5_enabled,json=md5Enabled,proto3" json:"md5_enabled,omitempty"`
	KeychainEnabled           bool             `protobuf:"varint,7,opt,name=keychain_enabled,json=keychainEnabled,proto3" json:"keychain_enabled,omitempty"`
	NegotiatedLocalKeepalive  uint32           `protobuf:"varint,8,opt,name=negotiated_local_keepalive,json=negotiatedLocalKeepalive,proto3" json:"negotiated_local_keepalive,omitempty"`
	NegotiatedRemoteKeepalive uint32           `protobuf:"varint,9,opt,name=negotiated_remote_keepalive,json=negotiatedRemoteKeepalive,proto3" json:"negotiated_remote_keepalive,omitempty"`
	NegotiatedDeadTime        uint32           `protobuf:"varint,10,opt,name=negotiated_dead_time,json=negotiatedDeadTime,proto3" json:"negotiated_dead_time,omitempty"`
	PceRequestRx              uint32           `protobuf:"varint,11,opt,name=pce_request_rx,json=pceRequestRx,proto3" json:"pce_request_rx,omitempty"`
	PceRequestTx              uint32           `protobuf:"varint,12,opt,name=pce_request_tx,json=pceRequestTx,proto3" json:"pce_request_tx,omitempty"`
	PceReplyRx                uint32           `protobuf:"varint,13,opt,name=pce_reply_rx,json=pceReplyRx,proto3" json:"pce_reply_rx,omitempty"`
	PceReplyTx                uint32           `protobuf:"varint,14,opt,name=pce_reply_tx,json=pceReplyTx,proto3" json:"pce_reply_tx,omitempty"`
	PceErrorRx                uint32           `protobuf:"varint,15,opt,name=pce_error_rx,json=pceErrorRx,proto3" json:"pce_error_rx,omitempty"`
	PceErrorTx                uint32           `protobuf:"varint,16,opt,name=pce_error_tx,json=pceErrorTx,proto3" json:"pce_error_tx,omitempty"`
	PceOpenTx                 uint32           `protobuf:"varint,17,opt,name=pce_open_tx,json=pceOpenTx,proto3" json:"pce_open_tx,omitempty"`
	PceOpenRx                 uint32           `protobuf:"varint,18,opt,name=pce_open_rx,json=pceOpenRx,proto3" json:"pce_open_rx,omitempty"`
	PceReportRx               uint32           `protobuf:"varint,19,opt,name=pce_report_rx,json=pceReportRx,proto3" json:"pce_report_rx,omitempty"`
	PceReportTx               uint32           `protobuf:"varint,20,opt,name=pce_report_tx,json=pceReportTx,proto3" json:"pce_report_tx,omitempty"`
	PceUpdateRx               uint32           `protobuf:"varint,21,opt,name=pce_update_rx,json=pceUpdateRx,proto3" json:"pce_update_rx,omitempty"`
	PceUpdateTx               uint32           `protobuf:"varint,22,opt,name=pce_update_tx,json=pceUpdateTx,proto3" json:"pce_update_tx,omitempty"`
	PceInitiateRx             uint32           `protobuf:"varint,23,opt,name=pce_initiate_rx,json=pceInitiateRx,proto3" json:"pce_initiate_rx,omitempty"`
	PceInitiateTx             uint32           `protobuf:"varint,24,opt,name=pce_initiate_tx,json=pceInitiateTx,proto3" json:"pce_initiate_tx,omitempty"`
	PceKeepaliveTx            uint64           `protobuf:"varint,25,opt,name=pce_keepalive_tx,json=pceKeepaliveTx,proto3" json:"pce_keepalive_tx,omitempty"`
	PceKeepaliveRx            uint64           `protobuf:"varint,26,opt,name=pce_keepalive_rx,json=pceKeepaliveRx,proto3" json:"pce_keepalive_rx,omitempty"`
	LocalSessionId            uint32           `protobuf:"varint,27,opt,name=local_session_id,json=localSessionId,proto3" json:"local_session_id,omitempty"`
	RemoteSessionId           uint32           `protobuf:"varint,28,opt,name=remote_session_id,json=remoteSessionId,proto3" json:"remote_session_id,omitempty"`
	MinimumKeepaliveInterval  uint32           `protobuf:"varint,29,opt,name=minimum_keepalive_interval,json=minimumKeepaliveInterval,proto3" json:"minimum_keepalive_interval,omitempty"`
	MaximumDeadInterval       uint32           `protobuf:"varint,30,opt,name=maximum_dead_interval,json=maximumDeadInterval,proto3" json:"maximum_dead_interval,omitempty"`
	LastErrorRx               *PcePcepErrorBag `protobuf:"bytes,31,opt,name=last_error_rx,json=lastErrorRx,proto3" json:"last_error_rx,omitempty"`
	LastErrorTx               *PcePcepErrorBag `protobuf:"bytes,32,opt,name=last_error_tx,json=lastErrorTx,proto3" json:"last_error_tx,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}         `json:"-"`
	XXX_unrecognized          []byte           `json:"-"`
	XXX_sizecache             int32            `json:"-"`
}

func (m *PcePcepDetailBag) Reset()         { *m = PcePcepDetailBag{} }
func (m *PcePcepDetailBag) String() string { return proto.CompactTextString(m) }
func (*PcePcepDetailBag) ProtoMessage()    {}
func (*PcePcepDetailBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3afadee1dc9c91, []int{3}
}

func (m *PcePcepDetailBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePcepDetailBag.Unmarshal(m, b)
}
func (m *PcePcepDetailBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePcepDetailBag.Marshal(b, m, deterministic)
}
func (m *PcePcepDetailBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePcepDetailBag.Merge(m, src)
}
func (m *PcePcepDetailBag) XXX_Size() int {
	return xxx_messageInfo_PcePcepDetailBag.Size(m)
}
func (m *PcePcepDetailBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePcepDetailBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePcepDetailBag proto.InternalMessageInfo

func (m *PcePcepDetailBag) GetBriefPcepInformation() *PcePcepBag {
	if m != nil {
		return m.BriefPcepInformation
	}
	return nil
}

func (m *PcePcepDetailBag) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *PcePcepDetailBag) GetSpeakerId() string {
	if m != nil {
		return m.SpeakerId
	}
	return ""
}

func (m *PcePcepDetailBag) GetPcepUpTime() uint32 {
	if m != nil {
		return m.PcepUpTime
	}
	return 0
}

func (m *PcePcepDetailBag) GetKeepalives() uint32 {
	if m != nil {
		return m.Keepalives
	}
	return 0
}

func (m *PcePcepDetailBag) GetMd5Enabled() bool {
	if m != nil {
		return m.Md5Enabled
	}
	return false
}

func (m *PcePcepDetailBag) GetKeychainEnabled() bool {
	if m != nil {
		return m.KeychainEnabled
	}
	return false
}

func (m *PcePcepDetailBag) GetNegotiatedLocalKeepalive() uint32 {
	if m != nil {
		return m.NegotiatedLocalKeepalive
	}
	return 0
}

func (m *PcePcepDetailBag) GetNegotiatedRemoteKeepalive() uint32 {
	if m != nil {
		return m.NegotiatedRemoteKeepalive
	}
	return 0
}

func (m *PcePcepDetailBag) GetNegotiatedDeadTime() uint32 {
	if m != nil {
		return m.NegotiatedDeadTime
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceRequestRx() uint32 {
	if m != nil {
		return m.PceRequestRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceRequestTx() uint32 {
	if m != nil {
		return m.PceRequestTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceReplyRx() uint32 {
	if m != nil {
		return m.PceReplyRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceReplyTx() uint32 {
	if m != nil {
		return m.PceReplyTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceErrorRx() uint32 {
	if m != nil {
		return m.PceErrorRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceErrorTx() uint32 {
	if m != nil {
		return m.PceErrorTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceOpenTx() uint32 {
	if m != nil {
		return m.PceOpenTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceOpenRx() uint32 {
	if m != nil {
		return m.PceOpenRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceReportRx() uint32 {
	if m != nil {
		return m.PceReportRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceReportTx() uint32 {
	if m != nil {
		return m.PceReportTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceUpdateRx() uint32 {
	if m != nil {
		return m.PceUpdateRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceUpdateTx() uint32 {
	if m != nil {
		return m.PceUpdateTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceInitiateRx() uint32 {
	if m != nil {
		return m.PceInitiateRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceInitiateTx() uint32 {
	if m != nil {
		return m.PceInitiateTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceKeepaliveTx() uint64 {
	if m != nil {
		return m.PceKeepaliveTx
	}
	return 0
}

func (m *PcePcepDetailBag) GetPceKeepaliveRx() uint64 {
	if m != nil {
		return m.PceKeepaliveRx
	}
	return 0
}

func (m *PcePcepDetailBag) GetLocalSessionId() uint32 {
	if m != nil {
		return m.LocalSessionId
	}
	return 0
}

func (m *PcePcepDetailBag) GetRemoteSessionId() uint32 {
	if m != nil {
		return m.RemoteSessionId
	}
	return 0
}

func (m *PcePcepDetailBag) GetMinimumKeepaliveInterval() uint32 {
	if m != nil {
		return m.MinimumKeepaliveInterval
	}
	return 0
}

func (m *PcePcepDetailBag) GetMaximumDeadInterval() uint32 {
	if m != nil {
		return m.MaximumDeadInterval
	}
	return 0
}

func (m *PcePcepDetailBag) GetLastErrorRx() *PcePcepErrorBag {
	if m != nil {
		return m.LastErrorRx
	}
	return nil
}

func (m *PcePcepDetailBag) GetLastErrorTx() *PcePcepErrorBag {
	if m != nil {
		return m.LastErrorTx
	}
	return nil
}

type PcePeerDetailBag struct {
	PeerAddressXr         string            `protobuf:"bytes,50,opt,name=peer_address_xr,json=peerAddressXr,proto3" json:"peer_address_xr,omitempty"`
	PeerProtocol          string            `protobuf:"bytes,51,opt,name=peer_protocol,json=peerProtocol,proto3" json:"peer_protocol,omitempty"`
	DetailPcepInformation *PcePcepDetailBag `protobuf:"bytes,52,opt,name=detail_pcep_information,json=detailPcepInformation,proto3" json:"detail_pcep_information,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *PcePeerDetailBag) Reset()         { *m = PcePeerDetailBag{} }
func (m *PcePeerDetailBag) String() string { return proto.CompactTextString(m) }
func (*PcePeerDetailBag) ProtoMessage()    {}
func (*PcePeerDetailBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f3afadee1dc9c91, []int{4}
}

func (m *PcePeerDetailBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PcePeerDetailBag.Unmarshal(m, b)
}
func (m *PcePeerDetailBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PcePeerDetailBag.Marshal(b, m, deterministic)
}
func (m *PcePeerDetailBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PcePeerDetailBag.Merge(m, src)
}
func (m *PcePeerDetailBag) XXX_Size() int {
	return xxx_messageInfo_PcePeerDetailBag.Size(m)
}
func (m *PcePeerDetailBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PcePeerDetailBag.DiscardUnknown(m)
}

var xxx_messageInfo_PcePeerDetailBag proto.InternalMessageInfo

func (m *PcePeerDetailBag) GetPeerAddressXr() string {
	if m != nil {
		return m.PeerAddressXr
	}
	return ""
}

func (m *PcePeerDetailBag) GetPeerProtocol() string {
	if m != nil {
		return m.PeerProtocol
	}
	return ""
}

func (m *PcePeerDetailBag) GetDetailPcepInformation() *PcePcepDetailBag {
	if m != nil {
		return m.DetailPcepInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*PcePeerDetailBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_detail_infos.peer_detail_info.pce_peer_detail_bag_KEYS")
	proto.RegisterType((*PcePcepBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_detail_infos.peer_detail_info.pce_pcep_bag")
	proto.RegisterType((*PcePcepErrorBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_detail_infos.peer_detail_info.pce_pcep_error_bag")
	proto.RegisterType((*PcePcepDetailBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_detail_infos.peer_detail_info.pce_pcep_detail_bag")
	proto.RegisterType((*PcePeerDetailBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.peer_detail_infos.peer_detail_info.pce_peer_detail_bag")
}

func init() { proto.RegisterFile("pce_peer_detail_bag.proto", fileDescriptor_2f3afadee1dc9c91) }

var fileDescriptor_2f3afadee1dc9c91 = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x5b, 0x6f, 0x23, 0x35,
	0x14, 0xc7, 0x15, 0xf6, 0x42, 0xe3, 0x34, 0xbd, 0xb8, 0x37, 0xb7, 0xcb, 0x76, 0x4b, 0x40, 0xa8,
	0x80, 0x14, 0xa1, 0x2e, 0x7d, 0x41, 0x05, 0x09, 0xb4, 0x7d, 0x88, 0x16, 0x89, 0xca, 0xc9, 0x16,
	0x78, 0xc1, 0x72, 0x66, 0x4e, 0x83, 0xd5, 0xb9, 0x18, 0x8f, 0x53, 0x39, 0x7c, 0x00, 0x5e, 0xf7,
	0x99, 0xef, 0xc9, 0x07, 0x40, 0x3e, 0x9e, 0xcc, 0x4c, 0x26, 0xfb, 0xc6, 0x8a, 0xb7, 0xe4, 0x7f,
	0x7e, 0x7f, 0x1f, 0xcf, 0xf8, 0x6f, 0x8f, 0xc9, 0xb1, 0x8e, 0x40, 0x68, 0x00, 0x23, 0x62, 0xb0,
	0x52, 0x25, 0x62, 0x2a, 0x67, 0x43, 0x6d, 0x72, 0x9b, 0xd3, 0x1f, 0x22, 0x55, 0x44, 0xb9, 0x50,
	0x79, 0x21, 0x9c, 0x11, 0x2a, 0xbb, 0x33, 0x52, 0x38, 0x1b, 0x89, 0x5c, 0x83, 0x19, 0xea, 0x08,
	0x86, 0x4d, 0x9b, 0xca, 0xee, 0xf2, 0x62, 0x4d, 0x19, 0x7c, 0x4b, 0xd8, 0x3b, 0x1a, 0x88, 0xd7,
	0xd7, 0xbf, 0x8e, 0xe9, 0xc7, 0x64, 0x13, 0x75, 0x19, 0xc7, 0x06, 0x8a, 0x82, 0x75, 0xce, 0x3a,
	0xe7, 0x5d, 0xde, 0xf3, 0xda, 0xf7, 0x41, 0x1a, 0xbc, 0x7d, 0x44, 0x36, 0xd1, 0x1f, 0x81, 0xf6,
	0x46, 0xfa, 0x9c, 0x10, 0xfc, 0x5d, 0x58, 0x69, 0xa1, 0x74, 0x74, 0xbd, 0x32, 0xf6, 0x02, 0x3d,
	0x21, 0x1b, 0x58, 0xb9, 0x9b, 0x27, 0xec, 0x83, 0xb3, 0xce, 0xf9, 0x06, 0xaf, 0xfe, 0xd3, 0x2f,
	0xc9, 0x6e, 0x24, 0xb5, 0x9c, 0xaa, 0x44, 0xd9, 0x85, 0x98, 0xeb, 0xd8, 0x8f, 0xf0, 0x08, 0xa1,
	0x9d, 0xba, 0xf0, 0x06, 0x75, 0x7a, 0x49, 0x0e, 0x1b, 0xb0, 0xca, 0x0a, 0x2b, 0x33, 0xab, 0xbc,
	0xe3, 0x31, 0x3a, 0x0e, 0xea, 0xea, 0xa8, 0x2e, 0xd2, 0x2b, 0x72, 0xd2, 0xb0, 0x15, 0x30, 0x4b,
	0x21, 0xb3, 0xc2, 0xe4, 0x73, 0xab, 0xb2, 0x19, 0x7b, 0x82, 0x56, 0x56, 0x13, 0xe3, 0x00, 0xf0,
	0x50, 0xa7, 0xdf, 0x90, 0xe3, 0x86, 0xdb, 0x1a, 0x35, 0x9b, 0x81, 0x81, 0x58, 0x14, 0x8b, 0x2c,
	0x62, 0x4f, 0xd1, 0x7c, 0x54, 0x03, 0x93, 0x65, 0x7d, 0xbc, 0xc8, 0x22, 0x7a, 0x41, 0x1a, 0x53,
	0x12, 0xf1, 0x54, 0x3c, 0x80, 0x29, 0x54, 0x9e, 0xb1, 0x0f, 0xd1, 0xb7, 0x57, 0x17, 0x5f, 0x4d,
	0x6f, 0x43, 0xa9, 0xed, 0x81, 0xc4, 0xca, 0xd0, 0x6b, 0x63, 0xcd, 0xe3, 0x6b, 0xbe, 0xcf, 0xe0,
	0x37, 0x42, 0xab, 0x05, 0x01, 0x63, 0x72, 0x83, 0xcb, 0x32, 0x20, 0x7d, 0x1d, 0x95, 0xff, 0xed,
	0x42, 0x87, 0x95, 0xe9, 0xf3, 0x9e, 0x8e, 0xae, 0xbd, 0x36, 0x59, 0x68, 0xa0, 0x9f, 0x92, 0xad,
	0x8a, 0x79, 0x90, 0xc9, 0x1c, 0x70, 0x85, 0xfa, 0x7c, 0xb3, 0x84, 0x6e, 0xbd, 0x36, 0xf8, 0x7b,
	0x93, 0xec, 0x55, 0x0d, 0xea, 0xc4, 0xd0, 0xbf, 0x3a, 0xe4, 0x70, 0x6a, 0x14, 0xdc, 0x85, 0x8a,
	0x0f, 0x97, 0x49, 0xa5, 0xf5, 0x4f, 0xe8, 0x7b, 0xf5, 0x2e, 0x6e, 0x86, 0xff, 0x3d, 0xae, 0xc3,
	0x66, 0xd6, 0xf8, 0x3e, 0xf6, 0xbb, 0x89, 0x40, 0x8f, 0xea, 0x6e, 0x74, 0x9f, 0x3c, 0xc1, 0x67,
	0xc0, 0xd9, 0x77, 0x79, 0xf8, 0xe3, 0x73, 0x59, 0x68, 0x90, 0xf7, 0x60, 0x84, 0x8a, 0x31, 0x55,
	0x5d, 0xde, 0x2d, 0x95, 0x51, 0x4c, 0xcf, 0x30, 0xc6, 0x5a, 0xcc, 0xb5, 0xb0, 0x2a, 0x0d, 0x21,
	0xea, 0x73, 0x8c, 0xf2, 0x1b, 0x3d, 0x51, 0x29, 0xd0, 0x53, 0x42, 0xee, 0x01, 0xb4, 0x4c, 0xd4,
	0x03, 0x14, 0x98, 0x94, 0x3e, 0x6f, 0x28, 0xf4, 0x05, 0xe9, 0xa5, 0xf1, 0xa5, 0x80, 0x4c, 0x4e,
	0x13, 0x88, 0xcb, 0x34, 0x90, 0x34, 0xbe, 0xbc, 0x0e, 0x0a, 0xfd, 0x9c, 0xec, 0xdc, 0xc3, 0x22,
	0xfa, 0x5d, 0xaa, 0xac, 0xa2, 0xc2, 0xda, 0x6f, 0x2f, 0xf5, 0x25, 0x7a, 0x45, 0x4e, 0x32, 0x98,
	0xe5, 0x18, 0xd9, 0x58, 0x24, 0x79, 0x24, 0x13, 0x51, 0xb5, 0xc2, 0xc5, 0xef, 0x73, 0x56, 0x13,
	0x3f, 0x7a, 0xe0, 0xf5, 0xb2, 0x4e, 0xbf, 0x23, 0xcf, 0x1a, 0x6e, 0x03, 0x69, 0x6e, 0xa1, 0x61,
	0xef, 0xa2, 0xfd, 0xb8, 0x46, 0x38, 0x12, 0xb5, 0xff, 0x2b, 0xb2, 0xdf, 0xf0, 0xc7, 0x20, 0xe3,
	0xf0, 0x4e, 0x08, 0x1a, 0x69, 0x5d, 0x7b, 0x05, 0x32, 0xc6, 0x77, 0x83, 0xc9, 0x01, 0x61, 0xe0,
	0x8f, 0x39, 0x14, 0x56, 0x18, 0xc7, 0x7a, 0xcb, 0xe4, 0x00, 0x0f, 0x22, 0x77, 0x6d, 0xca, 0x3a,
	0xb6, 0xd9, 0xa6, 0x26, 0xae, 0x5c, 0x09, 0x61, 0x40, 0x27, 0x0b, 0x3f, 0x52, 0xbf, 0x5a, 0x09,
	0xee, 0x25, 0xde, 0x22, 0xac, 0x63, 0x5b, 0xab, 0x44, 0x3d, 0x46, 0x88, 0xb2, 0x71, 0x6c, 0xbb,
	0x22, 0x30, 0xc8, 0xbc, 0x45, 0x58, 0xc7, 0x76, 0x56, 0x89, 0x89, 0xa3, 0xa7, 0xa4, 0xe7, 0x89,
	0x5c, 0x43, 0xe6, 0x81, 0x5d, 0x04, 0xfc, 0x49, 0xf6, 0x93, 0x86, 0xac, 0x55, 0x37, 0x8e, 0xd1,
	0x95, 0x3a, 0x77, 0x61, 0xc7, 0xe1, 0x2c, 0x73, 0x83, 0xaf, 0x64, 0x6f, 0xb9, 0xe3, 0xfc, 0x34,
	0x73, 0x63, 0xd7, 0x18, 0xeb, 0xd8, 0x7e, 0x8b, 0x99, 0x54, 0x4c, 0x38, 0x0e, 0xfd, 0x38, 0x07,
	0x15, 0x13, 0x8e, 0x42, 0xde, 0x66, 0xac, 0x63, 0x87, 0x2d, 0x66, 0xe2, 0xe8, 0x67, 0x64, 0xdb,
	0x33, 0x2a, 0x53, 0xb8, 0x78, 0x7e, 0xa4, 0x23, 0xa4, 0xbc, 0x75, 0x54, 0xaa, 0x7c, 0x9d, 0xb3,
	0x8e, 0xb1, 0x35, 0x6e, 0xe2, 0xe8, 0x39, 0xd9, 0xf1, 0x5c, 0x95, 0x2b, 0x0f, 0x1e, 0x9f, 0x75,
	0xce, 0x1f, 0x73, 0xbf, 0xca, 0x55, 0x9a, 0xde, 0x45, 0x1a, 0xc7, 0x4e, 0xd6, 0x49, 0x8e, 0x64,
	0x08, 0x7b, 0x01, 0x85, 0x3f, 0x00, 0xfd, 0x56, 0x7d, 0x86, 0xcd, 0xb7, 0x50, 0x1f, 0x07, 0x79,
	0x14, 0xd3, 0x2f, 0xc8, 0x6e, 0x19, 0xec, 0x06, 0xfa, 0x11, 0xa2, 0xdb, 0xa1, 0x50, 0xb3, 0x57,
	0xe4, 0x24, 0x55, 0x99, 0x4a, 0xe7, 0x69, 0x63, 0x0e, 0x2a, 0xb3, 0x60, 0x1e, 0x64, 0xc2, 0x9e,
	0x87, 0xdd, 0x54, 0x12, 0xd5, 0x6c, 0x46, 0x65, 0xdd, 0x9f, 0xc1, 0xa9, 0x74, 0xe8, 0xc6, 0xad,
	0x50, 0x19, 0x4f, 0xd1, 0xb8, 0x57, 0x16, 0xfd, 0x5e, 0xa8, 0x3c, 0x7f, 0x92, 0x7e, 0x22, 0x0b,
	0x5b, 0x07, 0xf0, 0x05, 0x9e, 0x80, 0xb7, 0xef, 0xf5, 0x04, 0xac, 0x0e, 0x77, 0xde, 0xf3, 0xcd,
	0x96, 0xc9, 0x5e, 0xed, 0x6d, 0x1d, 0x3b, 0xfb, 0x9f, 0x7a, 0x4f, 0xdc, 0xe0, 0x9f, 0x4e, 0xf9,
	0x6d, 0x58, 0xbd, 0x4d, 0x60, 0xa6, 0x1a, 0x17, 0x09, 0xe1, 0x0c, 0xbb, 0xc0, 0x13, 0xb8, 0xdf,
	0xb8, 0x4b, 0xfc, 0x62, 0xe8, 0x27, 0x04, 0x05, 0x81, 0xd7, 0x9b, 0x28, 0x4f, 0xd8, 0x4b, 0xa4,
	0xf0, 0x16, 0x72, 0x53, 0x6a, 0xf4, 0x6d, 0x87, 0x1c, 0x95, 0x63, 0xaf, 0x7d, 0x69, 0xbe, 0xc6,
	0x67, 0xfd, 0xf9, 0xbd, 0x3e, 0x6b, 0xfd, 0x1c, 0xfc, 0x20, 0xfc, 0x6e, 0x7d, 0x71, 0xa6, 0x4f,
	0x71, 0xbe, 0x2f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xc7, 0x2c, 0x85, 0xab, 0x09, 0x00,
	0x00,
}
