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
// source: fsync_if_info_v2.proto

package cisco_ios_xr_ncs4k_freqsync_oper_frequency_synchronization_interface_datas_interface_data

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

type FsyncIfInfoV2_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncIfInfoV2_KEYS) Reset()         { *m = FsyncIfInfoV2_KEYS{} }
func (m *FsyncIfInfoV2_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncIfInfoV2_KEYS) ProtoMessage()    {}
func (*FsyncIfInfoV2_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{0}
}

func (m *FsyncIfInfoV2_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncIfInfoV2_KEYS.Unmarshal(m, b)
}
func (m *FsyncIfInfoV2_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncIfInfoV2_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncIfInfoV2_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncIfInfoV2_KEYS.Merge(m, src)
}
func (m *FsyncIfInfoV2_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncIfInfoV2_KEYS.Size(m)
}
func (m *FsyncIfInfoV2_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncIfInfoV2_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncIfInfoV2_KEYS proto.InternalMessageInfo

func (m *FsyncIfInfoV2_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type FsyncBagClockId struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ClockName            string   `protobuf:"bytes,3,opt,name=clock_name,json=clockName,proto3" json:"clock_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncBagClockId) Reset()         { *m = FsyncBagClockId{} }
func (m *FsyncBagClockId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagClockId) ProtoMessage()    {}
func (*FsyncBagClockId) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{1}
}

func (m *FsyncBagClockId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagClockId.Unmarshal(m, b)
}
func (m *FsyncBagClockId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagClockId.Marshal(b, m, deterministic)
}
func (m *FsyncBagClockId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagClockId.Merge(m, src)
}
func (m *FsyncBagClockId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagClockId.Size(m)
}
func (m *FsyncBagClockId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagClockId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagClockId proto.InternalMessageInfo

func (m *FsyncBagClockId) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *FsyncBagClockId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FsyncBagClockId) GetClockName() string {
	if m != nil {
		return m.ClockName
	}
	return ""
}

type FsyncBagSourceId struct {
	SourceClass              string           `protobuf:"bytes,1,opt,name=source_class,json=sourceClass,proto3" json:"source_class,omitempty"`
	EthernetInterface        string           `protobuf:"bytes,2,opt,name=ethernet_interface,json=ethernetInterface,proto3" json:"ethernet_interface,omitempty"`
	SonetInterface           string           `protobuf:"bytes,3,opt,name=sonet_interface,json=sonetInterface,proto3" json:"sonet_interface,omitempty"`
	ClockId                  *FsyncBagClockId `protobuf:"bytes,4,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	InternalClockId          *FsyncBagClockId `protobuf:"bytes,5,opt,name=internal_clock_id,json=internalClockId,proto3" json:"internal_clock_id,omitempty"`
	PtpNode                  string           `protobuf:"bytes,6,opt,name=ptp_node,json=ptpNode,proto3" json:"ptp_node,omitempty"`
	SatelliteAccessInterface string           `protobuf:"bytes,7,opt,name=satellite_access_interface,json=satelliteAccessInterface,proto3" json:"satellite_access_interface,omitempty"`
	NtpNode                  string           `protobuf:"bytes,8,opt,name=ntp_node,json=ntpNode,proto3" json:"ntp_node,omitempty"`
	GnssReceiverId           *FsyncBagClockId `protobuf:"bytes,9,opt,name=gnss_receiver_id,json=gnssReceiverId,proto3" json:"gnss_receiver_id,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}         `json:"-"`
	XXX_unrecognized         []byte           `json:"-"`
	XXX_sizecache            int32            `json:"-"`
}

func (m *FsyncBagSourceId) Reset()         { *m = FsyncBagSourceId{} }
func (m *FsyncBagSourceId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSourceId) ProtoMessage()    {}
func (*FsyncBagSourceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{2}
}

func (m *FsyncBagSourceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSourceId.Unmarshal(m, b)
}
func (m *FsyncBagSourceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSourceId.Marshal(b, m, deterministic)
}
func (m *FsyncBagSourceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSourceId.Merge(m, src)
}
func (m *FsyncBagSourceId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSourceId.Size(m)
}
func (m *FsyncBagSourceId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSourceId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSourceId proto.InternalMessageInfo

func (m *FsyncBagSourceId) GetSourceClass() string {
	if m != nil {
		return m.SourceClass
	}
	return ""
}

func (m *FsyncBagSourceId) GetEthernetInterface() string {
	if m != nil {
		return m.EthernetInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetSonetInterface() string {
	if m != nil {
		return m.SonetInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetClockId() *FsyncBagClockId {
	if m != nil {
		return m.ClockId
	}
	return nil
}

func (m *FsyncBagSourceId) GetInternalClockId() *FsyncBagClockId {
	if m != nil {
		return m.InternalClockId
	}
	return nil
}

func (m *FsyncBagSourceId) GetPtpNode() string {
	if m != nil {
		return m.PtpNode
	}
	return ""
}

func (m *FsyncBagSourceId) GetSatelliteAccessInterface() string {
	if m != nil {
		return m.SatelliteAccessInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetNtpNode() string {
	if m != nil {
		return m.NtpNode
	}
	return ""
}

func (m *FsyncBagSourceId) GetGnssReceiverId() *FsyncBagClockId {
	if m != nil {
		return m.GnssReceiverId
	}
	return nil
}

type FsyncBagQl struct {
	QualityLevelOption      string   `protobuf:"bytes,1,opt,name=quality_level_option,json=qualityLevelOption,proto3" json:"quality_level_option,omitempty"`
	Option1Value            string   `protobuf:"bytes,2,opt,name=option1_value,json=option1Value,proto3" json:"option1_value,omitempty"`
	Option2Generation1Value string   `protobuf:"bytes,3,opt,name=option2_generation1_value,json=option2Generation1Value,proto3" json:"option2_generation1_value,omitempty"`
	Option2Generation2Value string   `protobuf:"bytes,4,opt,name=option2_generation2_value,json=option2Generation2Value,proto3" json:"option2_generation2_value,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *FsyncBagQl) Reset()         { *m = FsyncBagQl{} }
func (m *FsyncBagQl) String() string { return proto.CompactTextString(m) }
func (*FsyncBagQl) ProtoMessage()    {}
func (*FsyncBagQl) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{3}
}

func (m *FsyncBagQl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagQl.Unmarshal(m, b)
}
func (m *FsyncBagQl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagQl.Marshal(b, m, deterministic)
}
func (m *FsyncBagQl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagQl.Merge(m, src)
}
func (m *FsyncBagQl) XXX_Size() int {
	return xxx_messageInfo_FsyncBagQl.Size(m)
}
func (m *FsyncBagQl) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagQl.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagQl proto.InternalMessageInfo

func (m *FsyncBagQl) GetQualityLevelOption() string {
	if m != nil {
		return m.QualityLevelOption
	}
	return ""
}

func (m *FsyncBagQl) GetOption1Value() string {
	if m != nil {
		return m.Option1Value
	}
	return ""
}

func (m *FsyncBagQl) GetOption2Generation1Value() string {
	if m != nil {
		return m.Option2Generation1Value
	}
	return ""
}

func (m *FsyncBagQl) GetOption2Generation2Value() string {
	if m != nil {
		return m.Option2Generation2Value
	}
	return ""
}

type FsyncBagSpInfo struct {
	SelectionPoint            uint32   `protobuf:"varint,1,opt,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	SelectionPointDescription string   `protobuf:"bytes,2,opt,name=selection_point_description,json=selectionPointDescription,proto3" json:"selection_point_description,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *FsyncBagSpInfo) Reset()         { *m = FsyncBagSpInfo{} }
func (m *FsyncBagSpInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSpInfo) ProtoMessage()    {}
func (*FsyncBagSpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{4}
}

func (m *FsyncBagSpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSpInfo.Unmarshal(m, b)
}
func (m *FsyncBagSpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSpInfo.Marshal(b, m, deterministic)
}
func (m *FsyncBagSpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSpInfo.Merge(m, src)
}
func (m *FsyncBagSpInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSpInfo.Size(m)
}
func (m *FsyncBagSpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSpInfo proto.InternalMessageInfo

func (m *FsyncBagSpInfo) GetSelectionPoint() uint32 {
	if m != nil {
		return m.SelectionPoint
	}
	return 0
}

func (m *FsyncBagSpInfo) GetSelectionPointDescription() string {
	if m != nil {
		return m.SelectionPointDescription
	}
	return ""
}

type FsyncBagTimestamp struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          uint32   `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncBagTimestamp) Reset()         { *m = FsyncBagTimestamp{} }
func (m *FsyncBagTimestamp) String() string { return proto.CompactTextString(m) }
func (*FsyncBagTimestamp) ProtoMessage()    {}
func (*FsyncBagTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{5}
}

func (m *FsyncBagTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagTimestamp.Unmarshal(m, b)
}
func (m *FsyncBagTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagTimestamp.Marshal(b, m, deterministic)
}
func (m *FsyncBagTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagTimestamp.Merge(m, src)
}
func (m *FsyncBagTimestamp) XXX_Size() int {
	return xxx_messageInfo_FsyncBagTimestamp.Size(m)
}
func (m *FsyncBagTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagTimestamp proto.InternalMessageInfo

func (m *FsyncBagTimestamp) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *FsyncBagTimestamp) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type FsyncBagEthPeerInfo struct {
	PeerState            string             `protobuf:"bytes,1,opt,name=peer_state,json=peerState,proto3" json:"peer_state,omitempty"`
	PeerStateTime        *FsyncBagTimestamp `protobuf:"bytes,2,opt,name=peer_state_time,json=peerStateTime,proto3" json:"peer_state_time,omitempty"`
	LastSsm              *FsyncBagTimestamp `protobuf:"bytes,3,opt,name=last_ssm,json=lastSsm,proto3" json:"last_ssm,omitempty"`
	PeerUpCount          uint32             `protobuf:"varint,4,opt,name=peer_up_count,json=peerUpCount,proto3" json:"peer_up_count,omitempty"`
	PeerTimeoutCount     uint32             `protobuf:"varint,5,opt,name=peer_timeout_count,json=peerTimeoutCount,proto3" json:"peer_timeout_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FsyncBagEthPeerInfo) Reset()         { *m = FsyncBagEthPeerInfo{} }
func (m *FsyncBagEthPeerInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncBagEthPeerInfo) ProtoMessage()    {}
func (*FsyncBagEthPeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{6}
}

func (m *FsyncBagEthPeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagEthPeerInfo.Unmarshal(m, b)
}
func (m *FsyncBagEthPeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagEthPeerInfo.Marshal(b, m, deterministic)
}
func (m *FsyncBagEthPeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagEthPeerInfo.Merge(m, src)
}
func (m *FsyncBagEthPeerInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncBagEthPeerInfo.Size(m)
}
func (m *FsyncBagEthPeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagEthPeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagEthPeerInfo proto.InternalMessageInfo

func (m *FsyncBagEthPeerInfo) GetPeerState() string {
	if m != nil {
		return m.PeerState
	}
	return ""
}

func (m *FsyncBagEthPeerInfo) GetPeerStateTime() *FsyncBagTimestamp {
	if m != nil {
		return m.PeerStateTime
	}
	return nil
}

func (m *FsyncBagEthPeerInfo) GetLastSsm() *FsyncBagTimestamp {
	if m != nil {
		return m.LastSsm
	}
	return nil
}

func (m *FsyncBagEthPeerInfo) GetPeerUpCount() uint32 {
	if m != nil {
		return m.PeerUpCount
	}
	return 0
}

func (m *FsyncBagEthPeerInfo) GetPeerTimeoutCount() uint32 {
	if m != nil {
		return m.PeerTimeoutCount
	}
	return 0
}

type FsyncBagEthEsmcStats struct {
	EsmcEventsSent        uint32   `protobuf:"varint,1,opt,name=esmc_events_sent,json=esmcEventsSent,proto3" json:"esmc_events_sent,omitempty"`
	EsmcEventsReceived    uint32   `protobuf:"varint,2,opt,name=esmc_events_received,json=esmcEventsReceived,proto3" json:"esmc_events_received,omitempty"`
	EsmcInfosSent         uint32   `protobuf:"varint,3,opt,name=esmc_infos_sent,json=esmcInfosSent,proto3" json:"esmc_infos_sent,omitempty"`
	EsmcInfosReceived     uint32   `protobuf:"varint,4,opt,name=esmc_infos_received,json=esmcInfosReceived,proto3" json:"esmc_infos_received,omitempty"`
	EsmcDnUsSent          uint32   `protobuf:"varint,5,opt,name=esmc_dn_us_sent,json=esmcDnUsSent,proto3" json:"esmc_dn_us_sent,omitempty"`
	EsmcDnUsReceived      uint32   `protobuf:"varint,6,opt,name=esmc_dn_us_received,json=esmcDnUsReceived,proto3" json:"esmc_dn_us_received,omitempty"`
	EsmcMalformedReceived uint32   `protobuf:"varint,7,opt,name=esmc_malformed_received,json=esmcMalformedReceived,proto3" json:"esmc_malformed_received,omitempty"`
	EsmcReceivedError     uint32   `protobuf:"varint,8,opt,name=esmc_received_error,json=esmcReceivedError,proto3" json:"esmc_received_error,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *FsyncBagEthEsmcStats) Reset()         { *m = FsyncBagEthEsmcStats{} }
func (m *FsyncBagEthEsmcStats) String() string { return proto.CompactTextString(m) }
func (*FsyncBagEthEsmcStats) ProtoMessage()    {}
func (*FsyncBagEthEsmcStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{7}
}

func (m *FsyncBagEthEsmcStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagEthEsmcStats.Unmarshal(m, b)
}
func (m *FsyncBagEthEsmcStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagEthEsmcStats.Marshal(b, m, deterministic)
}
func (m *FsyncBagEthEsmcStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagEthEsmcStats.Merge(m, src)
}
func (m *FsyncBagEthEsmcStats) XXX_Size() int {
	return xxx_messageInfo_FsyncBagEthEsmcStats.Size(m)
}
func (m *FsyncBagEthEsmcStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagEthEsmcStats.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagEthEsmcStats proto.InternalMessageInfo

func (m *FsyncBagEthEsmcStats) GetEsmcEventsSent() uint32 {
	if m != nil {
		return m.EsmcEventsSent
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcEventsReceived() uint32 {
	if m != nil {
		return m.EsmcEventsReceived
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcInfosSent() uint32 {
	if m != nil {
		return m.EsmcInfosSent
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcInfosReceived() uint32 {
	if m != nil {
		return m.EsmcInfosReceived
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcDnUsSent() uint32 {
	if m != nil {
		return m.EsmcDnUsSent
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcDnUsReceived() uint32 {
	if m != nil {
		return m.EsmcDnUsReceived
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcMalformedReceived() uint32 {
	if m != nil {
		return m.EsmcMalformedReceived
	}
	return 0
}

func (m *FsyncBagEthEsmcStats) GetEsmcReceivedError() uint32 {
	if m != nil {
		return m.EsmcReceivedError
	}
	return 0
}

type FsyncIfInfoV2 struct {
	Source                      *FsyncBagSourceId     `protobuf:"bytes,50,opt,name=source,proto3" json:"source,omitempty"`
	Name                        string                `protobuf:"bytes,51,opt,name=name,proto3" json:"name,omitempty"`
	State                       string                `protobuf:"bytes,52,opt,name=state,proto3" json:"state,omitempty"`
	SsmEnabled                  bool                  `protobuf:"varint,53,opt,name=ssm_enabled,json=ssmEnabled,proto3" json:"ssm_enabled,omitempty"`
	Squelched                   bool                  `protobuf:"varint,54,opt,name=squelched,proto3" json:"squelched,omitempty"`
	SelectionInput              bool                  `protobuf:"varint,55,opt,name=selection_input,json=selectionInput,proto3" json:"selection_input,omitempty"`
	Priority                    uint32                `protobuf:"varint,56,opt,name=priority,proto3" json:"priority,omitempty"`
	TimeOfDayPriority           uint32                `protobuf:"varint,57,opt,name=time_of_day_priority,json=timeOfDayPriority,proto3" json:"time_of_day_priority,omitempty"`
	SelectedSource              *FsyncBagSourceId     `protobuf:"bytes,58,opt,name=selected_source,json=selectedSource,proto3" json:"selected_source,omitempty"`
	DampingState                string                `protobuf:"bytes,59,opt,name=damping_state,json=dampingState,proto3" json:"damping_state,omitempty"`
	DampingTime                 uint32                `protobuf:"varint,60,opt,name=damping_time,json=dampingTime,proto3" json:"damping_time,omitempty"`
	WaitToRestoreTime           uint32                `protobuf:"varint,61,opt,name=wait_to_restore_time,json=waitToRestoreTime,proto3" json:"wait_to_restore_time,omitempty"`
	SupportsFrequency           bool                  `protobuf:"varint,62,opt,name=supports_frequency,json=supportsFrequency,proto3" json:"supports_frequency,omitempty"`
	SupportsTimeOfDay           bool                  `protobuf:"varint,63,opt,name=supports_time_of_day,json=supportsTimeOfDay,proto3" json:"supports_time_of_day,omitempty"`
	QualityLevelReceived        *FsyncBagQl           `protobuf:"bytes,64,opt,name=quality_level_received,json=qualityLevelReceived,proto3" json:"quality_level_received,omitempty"`
	QualityLevelDamped          *FsyncBagQl           `protobuf:"bytes,65,opt,name=quality_level_damped,json=qualityLevelDamped,proto3" json:"quality_level_damped,omitempty"`
	QualityLevelEffectiveInput  *FsyncBagQl           `protobuf:"bytes,66,opt,name=quality_level_effective_input,json=qualityLevelEffectiveInput,proto3" json:"quality_level_effective_input,omitempty"`
	QualityLevelEffectiveOutput *FsyncBagQl           `protobuf:"bytes,67,opt,name=quality_level_effective_output,json=qualityLevelEffectiveOutput,proto3" json:"quality_level_effective_output,omitempty"`
	QualityLevelSelectedSource  *FsyncBagQl           `protobuf:"bytes,68,opt,name=quality_level_selected_source,json=qualityLevelSelectedSource,proto3" json:"quality_level_selected_source,omitempty"`
	SpaSelectionPoint           []*FsyncBagSpInfo     `protobuf:"bytes,69,rep,name=spa_selection_point,json=spaSelectionPoint,proto3" json:"spa_selection_point,omitempty"`
	NodeSelectionPoint          []*FsyncBagSpInfo     `protobuf:"bytes,70,rep,name=node_selection_point,json=nodeSelectionPoint,proto3" json:"node_selection_point,omitempty"`
	EthernetPeerInformation     *FsyncBagEthPeerInfo  `protobuf:"bytes,71,opt,name=ethernet_peer_information,json=ethernetPeerInformation,proto3" json:"ethernet_peer_information,omitempty"`
	EsmcStatistics              *FsyncBagEthEsmcStats `protobuf:"bytes,72,opt,name=esmc_statistics,json=esmcStatistics,proto3" json:"esmc_statistics,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}              `json:"-"`
	XXX_unrecognized            []byte                `json:"-"`
	XXX_sizecache               int32                 `json:"-"`
}

func (m *FsyncIfInfoV2) Reset()         { *m = FsyncIfInfoV2{} }
func (m *FsyncIfInfoV2) String() string { return proto.CompactTextString(m) }
func (*FsyncIfInfoV2) ProtoMessage()    {}
func (*FsyncIfInfoV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_84da25c7bb20e3b5, []int{8}
}

func (m *FsyncIfInfoV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncIfInfoV2.Unmarshal(m, b)
}
func (m *FsyncIfInfoV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncIfInfoV2.Marshal(b, m, deterministic)
}
func (m *FsyncIfInfoV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncIfInfoV2.Merge(m, src)
}
func (m *FsyncIfInfoV2) XXX_Size() int {
	return xxx_messageInfo_FsyncIfInfoV2.Size(m)
}
func (m *FsyncIfInfoV2) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncIfInfoV2.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncIfInfoV2 proto.InternalMessageInfo

func (m *FsyncIfInfoV2) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FsyncIfInfoV2) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FsyncIfInfoV2) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *FsyncIfInfoV2) GetSsmEnabled() bool {
	if m != nil {
		return m.SsmEnabled
	}
	return false
}

func (m *FsyncIfInfoV2) GetSquelched() bool {
	if m != nil {
		return m.Squelched
	}
	return false
}

func (m *FsyncIfInfoV2) GetSelectionInput() bool {
	if m != nil {
		return m.SelectionInput
	}
	return false
}

func (m *FsyncIfInfoV2) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *FsyncIfInfoV2) GetTimeOfDayPriority() uint32 {
	if m != nil {
		return m.TimeOfDayPriority
	}
	return 0
}

func (m *FsyncIfInfoV2) GetSelectedSource() *FsyncBagSourceId {
	if m != nil {
		return m.SelectedSource
	}
	return nil
}

func (m *FsyncIfInfoV2) GetDampingState() string {
	if m != nil {
		return m.DampingState
	}
	return ""
}

func (m *FsyncIfInfoV2) GetDampingTime() uint32 {
	if m != nil {
		return m.DampingTime
	}
	return 0
}

func (m *FsyncIfInfoV2) GetWaitToRestoreTime() uint32 {
	if m != nil {
		return m.WaitToRestoreTime
	}
	return 0
}

func (m *FsyncIfInfoV2) GetSupportsFrequency() bool {
	if m != nil {
		return m.SupportsFrequency
	}
	return false
}

func (m *FsyncIfInfoV2) GetSupportsTimeOfDay() bool {
	if m != nil {
		return m.SupportsTimeOfDay
	}
	return false
}

func (m *FsyncIfInfoV2) GetQualityLevelReceived() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelReceived
	}
	return nil
}

func (m *FsyncIfInfoV2) GetQualityLevelDamped() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelDamped
	}
	return nil
}

func (m *FsyncIfInfoV2) GetQualityLevelEffectiveInput() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelEffectiveInput
	}
	return nil
}

func (m *FsyncIfInfoV2) GetQualityLevelEffectiveOutput() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelEffectiveOutput
	}
	return nil
}

func (m *FsyncIfInfoV2) GetQualityLevelSelectedSource() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelSelectedSource
	}
	return nil
}

func (m *FsyncIfInfoV2) GetSpaSelectionPoint() []*FsyncBagSpInfo {
	if m != nil {
		return m.SpaSelectionPoint
	}
	return nil
}

func (m *FsyncIfInfoV2) GetNodeSelectionPoint() []*FsyncBagSpInfo {
	if m != nil {
		return m.NodeSelectionPoint
	}
	return nil
}

func (m *FsyncIfInfoV2) GetEthernetPeerInformation() *FsyncBagEthPeerInfo {
	if m != nil {
		return m.EthernetPeerInformation
	}
	return nil
}

func (m *FsyncIfInfoV2) GetEsmcStatistics() *FsyncBagEthEsmcStats {
	if m != nil {
		return m.EsmcStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncIfInfoV2_KEYS)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_if_info_v2_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagQl)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_ql")
	proto.RegisterType((*FsyncBagSpInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_sp_info")
	proto.RegisterType((*FsyncBagTimestamp)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_timestamp")
	proto.RegisterType((*FsyncBagEthPeerInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_eth_peer_info")
	proto.RegisterType((*FsyncBagEthEsmcStats)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_bag_eth_esmc_stats")
	proto.RegisterType((*FsyncIfInfoV2)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.interface_datas.interface_data.fsync_if_info_v2")
}

func init() { proto.RegisterFile("fsync_if_info_v2.proto", fileDescriptor_84da25c7bb20e3b5) }

var fileDescriptor_84da25c7bb20e3b5 = []byte{
	// 1303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x58, 0xd1, 0x6e, 0x1c, 0x35,
	0x17, 0xd6, 0x34, 0x4d, 0xb2, 0x39, 0x9b, 0x4d, 0xb2, 0x6e, 0xda, 0x4c, 0xda, 0xbf, 0xff, 0x9f,
	0x6e, 0xd5, 0x9f, 0x5c, 0xd0, 0x14, 0xb6, 0xa5, 0x40, 0x29, 0x85, 0x92, 0xa4, 0x25, 0x02, 0xda,
	0x32, 0x9b, 0x82, 0x7a, 0x65, 0xb9, 0x33, 0xde, 0x64, 0xd4, 0x19, 0x7b, 0x32, 0xf6, 0x06, 0x82,
	0xb8, 0x46, 0x20, 0x71, 0x4d, 0x55, 0x09, 0x78, 0x00, 0x84, 0xe0, 0x1d, 0x90, 0x78, 0x0c, 0xde,
	0x05, 0xf9, 0xd8, 0x9e, 0xd9, 0xdd, 0xb6, 0xdc, 0xb1, 0xb9, 0x8b, 0xbf, 0xf3, 0x9d, 0xe3, 0xcf,
	0xc7, 0xc7, 0x67, 0xce, 0x06, 0xce, 0xf4, 0xd5, 0x91, 0x88, 0x69, 0xda, 0xa7, 0xa9, 0xe8, 0x4b,
	0x7a, 0xd8, 0xdd, 0x28, 0x4a, 0xa9, 0x25, 0x79, 0x14, 0xa7, 0x2a, 0x96, 0x34, 0x95, 0x8a, 0x7e,
	0x59, 0x52, 0x11, 0xab, 0x6b, 0x4f, 0x68, 0xbf, 0xe4, 0x07, 0xc8, 0x96, 0x05, 0x2f, 0x37, 0xcc,
	0x6a, 0xc0, 0x45, 0x7c, 0x44, 0x0d, 0xb6, 0x5f, 0x4a, 0x91, 0x7e, 0xc5, 0x74, 0x2a, 0xc5, 0x46,
	0x2a, 0x34, 0x2f, 0xfb, 0x2c, 0xe6, 0x34, 0x61, 0x9a, 0xa9, 0xb1, 0x75, 0xe7, 0x16, 0x9c, 0x1e,
	0xdf, 0x94, 0x7e, 0xb4, 0xfd, 0xa8, 0x47, 0x2e, 0xc1, 0x42, 0x4d, 0x15, 0x2c, 0xe7, 0x61, 0xb0,
	0x16, 0xac, 0xcf, 0x45, 0xad, 0x0a, 0xbd, 0xc7, 0x72, 0xde, 0xf9, 0x1c, 0x88, 0xf5, 0x7f, 0xcc,
	0xf6, 0x68, 0x9c, 0xc9, 0xf8, 0x09, 0x4d, 0x13, 0x42, 0xe0, 0xa4, 0x90, 0x89, 0x77, 0xc1, 0xbf,
	0xc9, 0x02, 0x9c, 0x48, 0x93, 0xf0, 0xc4, 0x5a, 0xb0, 0xde, 0x8a, 0x4e, 0xa4, 0x09, 0x39, 0x0f,
	0x60, 0xf9, 0x18, 0x7c, 0x0a, 0x99, 0x73, 0x88, 0x60, 0xe0, 0x3f, 0xa7, 0xe1, 0x54, 0x1d, 0x59,
	0xc9, 0x41, 0x19, 0x73, 0x13, 0xfa, 0x02, 0xcc, 0xbb, 0x45, 0x9c, 0x31, 0xa5, 0xdc, 0x16, 0x4d,
	0x8b, 0x6d, 0x1a, 0x88, 0x5c, 0x06, 0xc2, 0xf5, 0x3e, 0x2f, 0x05, 0xd7, 0xb4, 0x52, 0x8b, 0x3b,
	0xcf, 0x45, 0x6d, 0x6f, 0xd9, 0xf1, 0x06, 0xf2, 0x0a, 0x2c, 0x2a, 0x39, 0xca, 0xb5, 0x6a, 0x16,
	0x10, 0xae, 0x89, 0xdf, 0x06, 0xd0, 0xf0, 0x47, 0x0c, 0x4f, 0xae, 0x05, 0xeb, 0xcd, 0x6e, 0xbe,
	0xf1, 0xaf, 0x5d, 0xcd, 0xc6, 0xf3, 0x79, 0x8d, 0x66, 0xf1, 0xaf, 0x9d, 0x84, 0x3c, 0x0b, 0xa0,
	0x8d, 0x74, 0xc1, 0xb2, 0xca, 0x1c, 0x4e, 0x1f, 0x87, 0xa6, 0x45, 0xaf, 0x63, 0xd3, 0x69, 0x5b,
	0x85, 0x46, 0xa1, 0x0b, 0x8a, 0x05, 0x30, 0x83, 0x89, 0x9c, 0x2d, 0x74, 0x71, 0xcf, 0xd4, 0xc0,
	0x4d, 0x38, 0xab, 0x98, 0xe6, 0x59, 0x96, 0x6a, 0x4e, 0x59, 0x1c, 0x73, 0xa5, 0x86, 0xb2, 0x3e,
	0x8b, 0xe4, 0xb0, 0x62, 0xdc, 0x46, 0x42, 0x9d, 0xff, 0x55, 0x68, 0x08, 0x1f, 0xb8, 0x61, 0x03,
	0x0b, 0x17, 0xf8, 0x69, 0x00, 0x4b, 0x7b, 0x42, 0x29, 0x5a, 0xf2, 0x98, 0xa7, 0x87, 0xbc, 0x34,
	0xe9, 0x98, 0x3b, 0x8e, 0x74, 0x2c, 0x18, 0x19, 0x91, 0x53, 0xb1, 0x93, 0x74, 0xfe, 0x0a, 0x60,
	0xbe, 0xa6, 0x1d, 0x64, 0xe4, 0x35, 0x58, 0x3e, 0x18, 0xb0, 0x2c, 0xd5, 0x47, 0x34, 0xe3, 0x87,
	0x3c, 0xa3, 0xb2, 0x30, 0x5b, 0xb9, 0x42, 0x26, 0xce, 0xf6, 0xb1, 0x31, 0xdd, 0x47, 0x0b, 0xb9,
	0x08, 0x2d, 0xcb, 0x79, 0x9d, 0x1e, 0xb2, 0x6c, 0xe0, 0x4b, 0x79, 0xde, 0x81, 0x9f, 0x19, 0x8c,
	0xdc, 0x80, 0x55, 0xbb, 0xee, 0xd2, 0x3d, 0x2e, 0x78, 0xc9, 0x86, 0x1d, 0x6c, 0x3d, 0xaf, 0x38,
	0xc2, 0xdd, 0xda, 0xfe, 0x0f, 0xbe, 0x5d, 0xe7, 0x7b, 0xf2, 0x25, 0xbe, 0x5d, 0xf4, 0xed, 0x7c,
	0x0d, 0xed, 0xa1, 0x67, 0x5a, 0x60, 0x13, 0xc1, 0x27, 0xc5, 0x33, 0x1e, 0x1b, 0x1e, 0x2d, 0x64,
	0x2a, 0x34, 0x1e, 0xaf, 0x15, 0x2d, 0x54, 0xf0, 0x03, 0x83, 0x92, 0x5b, 0x70, 0x6e, 0x8c, 0x48,
	0x13, 0xae, 0xe2, 0x32, 0xb5, 0x39, 0xb1, 0x07, 0x5d, 0x1d, 0x75, 0xda, 0xaa, 0x09, 0x9d, 0x4f,
	0x87, 0x9b, 0x84, 0x4e, 0x73, 0xae, 0x34, 0xcb, 0x0b, 0x12, 0xc2, 0xac, 0xe2, 0xb1, 0x14, 0x89,
	0x72, 0xfb, 0xfa, 0x25, 0x59, 0x83, 0xa6, 0x60, 0x42, 0x7a, 0xab, 0x6d, 0x47, 0xc3, 0x50, 0xe7,
	0xf7, 0x29, 0x58, 0xa9, 0x63, 0x72, 0xbd, 0x4f, 0x0b, 0x6e, 0xea, 0xc9, 0x9c, 0xeb, 0x3c, 0x00,
	0x2e, 0x94, 0x66, 0xda, 0x77, 0xb7, 0x39, 0x83, 0xf4, 0x0c, 0x40, 0x7e, 0x08, 0x60, 0xb1, 0xb6,
	0xa3, 0x1e, 0xdc, 0xa1, 0xd9, 0x15, 0x13, 0x29, 0xc2, 0x2a, 0x01, 0x51, 0xab, 0x12, 0xb5, 0x9b,
	0xe6, 0x9c, 0x7c, 0x17, 0x40, 0x23, 0x63, 0x4a, 0x53, 0xa5, 0x72, 0x2c, 0x86, 0xc9, 0x2b, 0x9a,
	0x35, 0xfb, 0xf7, 0x54, 0x4e, 0x3a, 0x80, 0xe2, 0xe8, 0xa0, 0xa0, 0xb1, 0x1c, 0x08, 0x8d, 0x05,
	0xd6, 0x8a, 0x9a, 0x06, 0x7c, 0x58, 0x6c, 0x1a, 0x88, 0xbc, 0x0a, 0x04, 0x39, 0xc6, 0x5d, 0x0e,
	0xb4, 0x23, 0x4e, 0x23, 0x71, 0xc9, 0x58, 0x76, 0xad, 0x01, 0xd9, 0x9d, 0xef, 0xa7, 0x20, 0x1c,
	0xbd, 0x31, 0xae, 0xf2, 0x18, 0x2f, 0x41, 0x91, 0x75, 0x58, 0xc2, 0x15, 0x3f, 0xe4, 0x42, 0x2b,
	0xaa, 0x78, 0x5d, 0x8b, 0x06, 0xdf, 0x46, 0xb8, 0xc7, 0x85, 0x36, 0x0f, 0x73, 0x98, 0xe9, 0x3a,
	0x89, 0xff, 0x64, 0x91, 0x9a, 0xed, 0x5e, 0x77, 0x42, 0xfe, 0x0f, 0x8b, 0xe8, 0x61, 0x6a, 0xc3,
	0x85, 0x9e, 0x42, 0x72, 0xcb, 0xc0, 0x3b, 0x06, 0xc5, 0xc8, 0x1b, 0x70, 0x6a, 0x88, 0x57, 0x05,
	0xb6, 0x07, 0x6f, 0x57, 0xdc, 0x2a, 0xee, 0x25, 0x17, 0x37, 0x11, 0x74, 0xe0, 0xe2, 0xda, 0xb3,
	0xcf, 0x1b, 0x78, 0x4b, 0x3c, 0xb4, 0x61, 0x2f, 0xbb, 0xb0, 0x96, 0x56, 0x85, 0x9d, 0xb1, 0x69,
	0xf2, 0xd4, 0x2a, 0xea, 0x75, 0x58, 0x41, 0x7a, 0xce, 0xb2, 0xbe, 0x2c, 0x73, 0x9e, 0xd4, 0x2e,
	0xb3, 0xe8, 0x72, 0xda, 0x98, 0x3f, 0xf1, 0xd6, 0xca, 0xcf, 0xab, 0xf7, 0x6c, 0xca, 0xcb, 0x52,
	0x96, 0xd8, 0x81, 0x9d, 0x7a, 0x4f, 0xdd, 0x36, 0x86, 0xce, 0x1f, 0x6d, 0x58, 0x1a, 0x9f, 0x29,
	0xc8, 0x37, 0x01, 0xcc, 0xd8, 0x6f, 0x74, 0xd8, 0x9d, 0x60, 0xfd, 0x55, 0x73, 0x43, 0xe4, 0x76,
	0xc7, 0xd1, 0xc4, 0x0c, 0x1c, 0x57, 0xdd, 0x68, 0xc2, 0x72, 0x4e, 0x96, 0x61, 0xda, 0xbe, 0xe8,
	0x6b, 0x08, 0xda, 0x05, 0xf9, 0x1f, 0x34, 0x95, 0xca, 0x29, 0x17, 0xec, 0x71, 0xc6, 0x93, 0xf0,
	0x8d, 0xb5, 0x60, 0xbd, 0x11, 0x81, 0x52, 0xf9, 0xb6, 0x45, 0xc8, 0x7f, 0x60, 0x4e, 0x1d, 0x0c,
	0x78, 0x16, 0xef, 0xf3, 0x24, 0xbc, 0x8e, 0xe6, 0x1a, 0x18, 0xed, 0x81, 0xa9, 0x28, 0x06, 0x3a,
	0x7c, 0x13, 0x39, 0x75, 0x0f, 0xdc, 0x31, 0x28, 0x39, 0x0b, 0x8d, 0xa2, 0x4c, 0x65, 0x99, 0xea,
	0xa3, 0xf0, 0x2d, 0x4c, 0x6a, 0xb5, 0x26, 0x57, 0x60, 0xd9, 0xbc, 0x01, 0x2a, 0xfb, 0x34, 0x61,
	0x47, 0xb4, 0xe2, 0xbd, 0x6d, 0x93, 0x6f, 0x6c, 0xf7, 0xfb, 0x5b, 0xec, 0xe8, 0x81, 0x77, 0x78,
	0x1a, 0xf8, 0x6d, 0x79, 0xe2, 0x4e, 0x1f, 0xde, 0x38, 0x96, 0x84, 0x2f, 0x78, 0x19, 0x3d, 0x9b,
	0xf8, 0x8b, 0xd0, 0x4a, 0x58, 0x5e, 0xa4, 0x62, 0xcf, 0xb5, 0xcf, 0x77, 0xec, 0x57, 0xcc, 0x81,
	0xb6, 0x83, 0x5e, 0x00, 0xbf, 0xb6, 0xdd, 0xf3, 0xa6, 0xed, 0x0d, 0x0e, 0xc3, 0x5e, 0x76, 0x05,
	0x96, 0xbf, 0x60, 0xa9, 0xa6, 0x5a, 0xd2, 0x92, 0x2b, 0x2d, 0x4b, 0xd7, 0x68, 0xdf, 0xb5, 0x29,
	0x31, 0xb6, 0x5d, 0x19, 0x59, 0x0b, 0x3a, 0x5c, 0x06, 0xa2, 0x06, 0x45, 0x21, 0x4b, 0xad, 0x68,
	0x75, 0xc2, 0xf0, 0x16, 0xde, 0x45, 0xdb, 0x5b, 0xee, 0x78, 0x83, 0x89, 0x5f, 0xd1, 0x87, 0x72,
	0x1f, 0xbe, 0x37, 0xea, 0xb0, 0xeb, 0x53, 0x4f, 0x7e, 0x0a, 0xe0, 0xcc, 0xe8, 0x17, 0xbd, 0x7a,
	0x57, 0xef, 0x63, 0xe6, 0xf7, 0x26, 0x92, 0xf9, 0x83, 0x2c, 0x5a, 0x1e, 0x1e, 0x1e, 0xaa, 0xf7,
	0xfb, 0x2c, 0x18, 0x9f, 0x38, 0x4c, 0x3a, 0x79, 0x12, 0xde, 0x9e, 0xac, 0xba, 0x91, 0xd1, 0x66,
	0x0b, 0x25, 0x90, 0x5f, 0x02, 0x38, 0x3f, 0xaa, 0x8d, 0xf7, 0xfb, 0xe6, 0x71, 0x1c, 0x72, 0xf7,
	0x66, 0x3e, 0x98, 0xac, 0xc8, 0xb3, 0xc3, 0x22, 0xb7, 0xbd, 0x16, 0xfb, 0x50, 0x7f, 0x0d, 0xe0,
	0xbf, 0x2f, 0x13, 0x2b, 0x07, 0xda, 0xa8, 0xdd, 0x9c, 0xac, 0xda, 0x73, 0x2f, 0x54, 0x7b, 0x1f,
	0xb5, 0xbc, 0x20, 0xb7, 0xe3, 0x8d, 0x61, 0xeb, 0x18, 0x73, 0xdb, 0x1b, 0xed, 0x0e, 0x3f, 0x06,
	0x70, 0x4a, 0x15, 0x8c, 0x8e, 0x8f, 0x8d, 0xdb, 0x6b, 0x53, 0xeb, 0xcd, 0x6e, 0x36, 0x99, 0xde,
	0x65, 0xa7, 0xd7, 0xa8, 0xad, 0x0a, 0xd6, 0x1b, 0x9d, 0x53, 0x7f, 0x0e, 0x60, 0xd9, 0xfc, 0xee,
	0x78, 0x4e, 0xdf, 0x9d, 0x63, 0xd0, 0x47, 0x8c, 0x92, 0x31, 0x81, 0xbf, 0x05, 0xb0, 0x5a, 0xfd,
	0xe8, 0xad, 0x06, 0xd6, 0x32, 0xc7, 0x2d, 0xc2, 0xbb, 0x78, 0xd1, 0xe5, 0x44, 0x54, 0x8e, 0x4c,
	0xcc, 0xd1, 0x8a, 0x17, 0xf5, 0x80, 0xf3, 0x72, 0xa7, 0x96, 0x64, 0x2e, 0x7c, 0xb1, 0x1a, 0xd3,
	0x52, 0xa5, 0xd3, 0x58, 0x85, 0x1f, 0xa2, 0x4c, 0x35, 0x31, 0x99, 0xf5, 0x98, 0x68, 0x87, 0xc1,
	0x5e, 0x25, 0xe5, 0xf1, 0x0c, 0xfe, 0xe7, 0xe5, 0xea, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x05,
	0x54, 0x1e, 0x1c, 0x93, 0x11, 0x00, 0x00,
}
