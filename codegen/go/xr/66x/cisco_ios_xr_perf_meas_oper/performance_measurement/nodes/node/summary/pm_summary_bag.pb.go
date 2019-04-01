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
// source: pm_summary_bag.proto

package cisco_ios_xr_perf_meas_oper_performance_measurement_nodes_node_summary

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

type PmSummaryBag_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PmSummaryBag_KEYS) Reset()         { *m = PmSummaryBag_KEYS{} }
func (m *PmSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PmSummaryBag_KEYS) ProtoMessage()    {}
func (*PmSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{0}
}

func (m *PmSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *PmSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PmSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSummaryBag_KEYS.Merge(m, src)
}
func (m *PmSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PmSummaryBag_KEYS.Size(m)
}
func (m *PmSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PmSummaryBag_KEYS proto.InternalMessageInfo

func (m *PmSummaryBag_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

type PmDelayProfileBag struct {
	ProbeInterval                          uint32   `protobuf:"varint,1,opt,name=probe_interval,json=probeInterval,proto3" json:"probe_interval,omitempty"`
	BurstInterval                          uint32   `protobuf:"varint,2,opt,name=burst_interval,json=burstInterval,proto3" json:"burst_interval,omitempty"`
	BurstCount                             uint32   `protobuf:"varint,3,opt,name=burst_count,json=burstCount,proto3" json:"burst_count,omitempty"`
	DelayMeasurementMode                   string   `protobuf:"bytes,4,opt,name=delay_measurement_mode,json=delayMeasurementMode,proto3" json:"delay_measurement_mode,omitempty"`
	PeriodicAdvertisementEnabled           bool     `protobuf:"varint,5,opt,name=periodic_advertisement_enabled,json=periodicAdvertisementEnabled,proto3" json:"periodic_advertisement_enabled,omitempty"`
	PeriodicAdvertisementInterval          uint32   `protobuf:"varint,6,opt,name=periodic_advertisement_interval,json=periodicAdvertisementInterval,proto3" json:"periodic_advertisement_interval,omitempty"`
	EffectivePeriodicAdvertisementInterval uint32   `protobuf:"varint,7,opt,name=effective_periodic_advertisement_interval,json=effectivePeriodicAdvertisementInterval,proto3" json:"effective_periodic_advertisement_interval,omitempty"`
	PeriodicAdvertisementThreshold         uint32   `protobuf:"varint,8,opt,name=periodic_advertisement_threshold,json=periodicAdvertisementThreshold,proto3" json:"periodic_advertisement_threshold,omitempty"`
	PeriodicAdvertisementMinimumChange     uint32   `protobuf:"varint,9,opt,name=periodic_advertisement_minimum_change,json=periodicAdvertisementMinimumChange,proto3" json:"periodic_advertisement_minimum_change,omitempty"`
	AcceleratedAdvertisementThreshold      uint32   `protobuf:"varint,10,opt,name=accelerated_advertisement_threshold,json=acceleratedAdvertisementThreshold,proto3" json:"accelerated_advertisement_threshold,omitempty"`
	AcceleratedAdvertisementMinimumChange  uint32   `protobuf:"varint,11,opt,name=accelerated_advertisement_minimum_change,json=acceleratedAdvertisementMinimumChange,proto3" json:"accelerated_advertisement_minimum_change,omitempty"`
	AcceleratedAdvertisementEnabled        bool     `protobuf:"varint,12,opt,name=accelerated_advertisement_enabled,json=acceleratedAdvertisementEnabled,proto3" json:"accelerated_advertisement_enabled,omitempty"`
	XXX_NoUnkeyedLiteral                   struct{} `json:"-"`
	XXX_unrecognized                       []byte   `json:"-"`
	XXX_sizecache                          int32    `json:"-"`
}

func (m *PmDelayProfileBag) Reset()         { *m = PmDelayProfileBag{} }
func (m *PmDelayProfileBag) String() string { return proto.CompactTextString(m) }
func (*PmDelayProfileBag) ProtoMessage()    {}
func (*PmDelayProfileBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{1}
}

func (m *PmDelayProfileBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmDelayProfileBag.Unmarshal(m, b)
}
func (m *PmDelayProfileBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmDelayProfileBag.Marshal(b, m, deterministic)
}
func (m *PmDelayProfileBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmDelayProfileBag.Merge(m, src)
}
func (m *PmDelayProfileBag) XXX_Size() int {
	return xxx_messageInfo_PmDelayProfileBag.Size(m)
}
func (m *PmDelayProfileBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmDelayProfileBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmDelayProfileBag proto.InternalMessageInfo

func (m *PmDelayProfileBag) GetProbeInterval() uint32 {
	if m != nil {
		return m.ProbeInterval
	}
	return 0
}

func (m *PmDelayProfileBag) GetBurstInterval() uint32 {
	if m != nil {
		return m.BurstInterval
	}
	return 0
}

func (m *PmDelayProfileBag) GetBurstCount() uint32 {
	if m != nil {
		return m.BurstCount
	}
	return 0
}

func (m *PmDelayProfileBag) GetDelayMeasurementMode() string {
	if m != nil {
		return m.DelayMeasurementMode
	}
	return ""
}

func (m *PmDelayProfileBag) GetPeriodicAdvertisementEnabled() bool {
	if m != nil {
		return m.PeriodicAdvertisementEnabled
	}
	return false
}

func (m *PmDelayProfileBag) GetPeriodicAdvertisementInterval() uint32 {
	if m != nil {
		return m.PeriodicAdvertisementInterval
	}
	return 0
}

func (m *PmDelayProfileBag) GetEffectivePeriodicAdvertisementInterval() uint32 {
	if m != nil {
		return m.EffectivePeriodicAdvertisementInterval
	}
	return 0
}

func (m *PmDelayProfileBag) GetPeriodicAdvertisementThreshold() uint32 {
	if m != nil {
		return m.PeriodicAdvertisementThreshold
	}
	return 0
}

func (m *PmDelayProfileBag) GetPeriodicAdvertisementMinimumChange() uint32 {
	if m != nil {
		return m.PeriodicAdvertisementMinimumChange
	}
	return 0
}

func (m *PmDelayProfileBag) GetAcceleratedAdvertisementThreshold() uint32 {
	if m != nil {
		return m.AcceleratedAdvertisementThreshold
	}
	return 0
}

func (m *PmDelayProfileBag) GetAcceleratedAdvertisementMinimumChange() uint32 {
	if m != nil {
		return m.AcceleratedAdvertisementMinimumChange
	}
	return 0
}

func (m *PmDelayProfileBag) GetAcceleratedAdvertisementEnabled() bool {
	if m != nil {
		return m.AcceleratedAdvertisementEnabled
	}
	return false
}

type PmQuerierDelayCountersBag struct {
	QueryPacketsSent                         uint64   `protobuf:"varint,1,opt,name=query_packets_sent,json=queryPacketsSent,proto3" json:"query_packets_sent,omitempty"`
	QueryPacketSentErrors                    uint64   `protobuf:"varint,2,opt,name=query_packet_sent_errors,json=queryPacketSentErrors,proto3" json:"query_packet_sent_errors,omitempty"`
	QueryPacketSentErrorNoIpAddress          uint64   `protobuf:"varint,3,opt,name=query_packet_sent_error_no_ip_address,json=queryPacketSentErrorNoIpAddress,proto3" json:"query_packet_sent_error_no_ip_address,omitempty"`
	QueryPacketsReceived                     uint64   `protobuf:"varint,4,opt,name=query_packets_received,json=queryPacketsReceived,proto3" json:"query_packets_received,omitempty"`
	ReceivedPacketErrorNegativeDelay         uint64   `protobuf:"varint,5,opt,name=received_packet_error_negative_delay,json=receivedPacketErrorNegativeDelay,proto3" json:"received_packet_error_negative_delay,omitempty"`
	ReceivedPacketErrorDelayExceedsThreshold uint64   `protobuf:"varint,6,opt,name=received_packet_error_delay_exceeds_threshold,json=receivedPacketErrorDelayExceedsThreshold,proto3" json:"received_packet_error_delay_exceeds_threshold,omitempty"`
	ReceivedPacketErrorMissingTxTimestamp    uint64   `protobuf:"varint,7,opt,name=received_packet_error_missing_tx_timestamp,json=receivedPacketErrorMissingTxTimestamp,proto3" json:"received_packet_error_missing_tx_timestamp,omitempty"`
	ReceivedPacketErrorMissingRxTimestamp    uint64   `protobuf:"varint,8,opt,name=received_packet_error_missing_rx_timestamp,json=receivedPacketErrorMissingRxTimestamp,proto3" json:"received_packet_error_missing_rx_timestamp,omitempty"`
	ReceivedPacketErrorProbeFull             uint64   `protobuf:"varint,9,opt,name=received_packet_error_probe_full,json=receivedPacketErrorProbeFull,proto3" json:"received_packet_error_probe_full,omitempty"`
	ReceivedPacketErrorProbeNotStarted       uint64   `protobuf:"varint,10,opt,name=received_packet_error_probe_not_started,json=receivedPacketErrorProbeNotStarted,proto3" json:"received_packet_error_probe_not_started,omitempty"`
	ReceivedPacketControlCodeError           uint64   `protobuf:"varint,11,opt,name=received_packet_control_code_error,json=receivedPacketControlCodeError,proto3" json:"received_packet_control_code_error,omitempty"`
	ReceivedPacketControlCodeNotification    uint64   `protobuf:"varint,12,opt,name=received_packet_control_code_notification,json=receivedPacketControlCodeNotification,proto3" json:"received_packet_control_code_notification,omitempty"`
	ProbesStarted                            uint64   `protobuf:"varint,13,opt,name=probes_started,json=probesStarted,proto3" json:"probes_started,omitempty"`
	ProbesComplete                           uint64   `protobuf:"varint,14,opt,name=probes_complete,json=probesComplete,proto3" json:"probes_complete,omitempty"`
	ProbesIncomplete                         uint64   `protobuf:"varint,15,opt,name=probes_incomplete,json=probesIncomplete,proto3" json:"probes_incomplete,omitempty"`
	Advertisement                            uint64   `protobuf:"varint,16,opt,name=advertisement,proto3" json:"advertisement,omitempty"`
	XXX_NoUnkeyedLiteral                     struct{} `json:"-"`
	XXX_unrecognized                         []byte   `json:"-"`
	XXX_sizecache                            int32    `json:"-"`
}

func (m *PmQuerierDelayCountersBag) Reset()         { *m = PmQuerierDelayCountersBag{} }
func (m *PmQuerierDelayCountersBag) String() string { return proto.CompactTextString(m) }
func (*PmQuerierDelayCountersBag) ProtoMessage()    {}
func (*PmQuerierDelayCountersBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{2}
}

func (m *PmQuerierDelayCountersBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmQuerierDelayCountersBag.Unmarshal(m, b)
}
func (m *PmQuerierDelayCountersBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmQuerierDelayCountersBag.Marshal(b, m, deterministic)
}
func (m *PmQuerierDelayCountersBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmQuerierDelayCountersBag.Merge(m, src)
}
func (m *PmQuerierDelayCountersBag) XXX_Size() int {
	return xxx_messageInfo_PmQuerierDelayCountersBag.Size(m)
}
func (m *PmQuerierDelayCountersBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmQuerierDelayCountersBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmQuerierDelayCountersBag proto.InternalMessageInfo

func (m *PmQuerierDelayCountersBag) GetQueryPacketsSent() uint64 {
	if m != nil {
		return m.QueryPacketsSent
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetQueryPacketSentErrors() uint64 {
	if m != nil {
		return m.QueryPacketSentErrors
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetQueryPacketSentErrorNoIpAddress() uint64 {
	if m != nil {
		return m.QueryPacketSentErrorNoIpAddress
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetQueryPacketsReceived() uint64 {
	if m != nil {
		return m.QueryPacketsReceived
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorNegativeDelay() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorNegativeDelay
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorDelayExceedsThreshold() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorDelayExceedsThreshold
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorMissingTxTimestamp() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorMissingTxTimestamp
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorMissingRxTimestamp() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorMissingRxTimestamp
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorProbeFull() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorProbeFull
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketErrorProbeNotStarted() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorProbeNotStarted
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketControlCodeError() uint64 {
	if m != nil {
		return m.ReceivedPacketControlCodeError
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetReceivedPacketControlCodeNotification() uint64 {
	if m != nil {
		return m.ReceivedPacketControlCodeNotification
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetProbesStarted() uint64 {
	if m != nil {
		return m.ProbesStarted
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetProbesComplete() uint64 {
	if m != nil {
		return m.ProbesComplete
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetProbesIncomplete() uint64 {
	if m != nil {
		return m.ProbesIncomplete
	}
	return 0
}

func (m *PmQuerierDelayCountersBag) GetAdvertisement() uint64 {
	if m != nil {
		return m.Advertisement
	}
	return 0
}

type PmIntfExclusiveDelayCountersInfoBag struct {
	QueryPacketSentErrorInterfaceDown uint64   `protobuf:"varint,1,opt,name=query_packet_sent_error_interface_down,json=queryPacketSentErrorInterfaceDown,proto3" json:"query_packet_sent_error_interface_down,omitempty"`
	QueryPacketSentErrorNoMplsCaps    uint64   `protobuf:"varint,2,opt,name=query_packet_sent_error_no_mpls_caps,json=queryPacketSentErrorNoMplsCaps,proto3" json:"query_packet_sent_error_no_mpls_caps,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *PmIntfExclusiveDelayCountersInfoBag) Reset()         { *m = PmIntfExclusiveDelayCountersInfoBag{} }
func (m *PmIntfExclusiveDelayCountersInfoBag) String() string { return proto.CompactTextString(m) }
func (*PmIntfExclusiveDelayCountersInfoBag) ProtoMessage()    {}
func (*PmIntfExclusiveDelayCountersInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{3}
}

func (m *PmIntfExclusiveDelayCountersInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag.Unmarshal(m, b)
}
func (m *PmIntfExclusiveDelayCountersInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag.Marshal(b, m, deterministic)
}
func (m *PmIntfExclusiveDelayCountersInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag.Merge(m, src)
}
func (m *PmIntfExclusiveDelayCountersInfoBag) XXX_Size() int {
	return xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag.Size(m)
}
func (m *PmIntfExclusiveDelayCountersInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmIntfExclusiveDelayCountersInfoBag proto.InternalMessageInfo

func (m *PmIntfExclusiveDelayCountersInfoBag) GetQueryPacketSentErrorInterfaceDown() uint64 {
	if m != nil {
		return m.QueryPacketSentErrorInterfaceDown
	}
	return 0
}

func (m *PmIntfExclusiveDelayCountersInfoBag) GetQueryPacketSentErrorNoMplsCaps() uint64 {
	if m != nil {
		return m.QueryPacketSentErrorNoMplsCaps
	}
	return 0
}

type PmExclusiveDelayCountersInfoUnion struct {
	Type                       string                               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	InterfaceExclusiveCounters *PmIntfExclusiveDelayCountersInfoBag `protobuf:"bytes,2,opt,name=interface_exclusive_counters,json=interfaceExclusiveCounters,proto3" json:"interface_exclusive_counters,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                             `json:"-"`
	XXX_unrecognized           []byte                               `json:"-"`
	XXX_sizecache              int32                                `json:"-"`
}

func (m *PmExclusiveDelayCountersInfoUnion) Reset()         { *m = PmExclusiveDelayCountersInfoUnion{} }
func (m *PmExclusiveDelayCountersInfoUnion) String() string { return proto.CompactTextString(m) }
func (*PmExclusiveDelayCountersInfoUnion) ProtoMessage()    {}
func (*PmExclusiveDelayCountersInfoUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{4}
}

func (m *PmExclusiveDelayCountersInfoUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmExclusiveDelayCountersInfoUnion.Unmarshal(m, b)
}
func (m *PmExclusiveDelayCountersInfoUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmExclusiveDelayCountersInfoUnion.Marshal(b, m, deterministic)
}
func (m *PmExclusiveDelayCountersInfoUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmExclusiveDelayCountersInfoUnion.Merge(m, src)
}
func (m *PmExclusiveDelayCountersInfoUnion) XXX_Size() int {
	return xxx_messageInfo_PmExclusiveDelayCountersInfoUnion.Size(m)
}
func (m *PmExclusiveDelayCountersInfoUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_PmExclusiveDelayCountersInfoUnion.DiscardUnknown(m)
}

var xxx_messageInfo_PmExclusiveDelayCountersInfoUnion proto.InternalMessageInfo

func (m *PmExclusiveDelayCountersInfoUnion) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PmExclusiveDelayCountersInfoUnion) GetInterfaceExclusiveCounters() *PmIntfExclusiveDelayCountersInfoBag {
	if m != nil {
		return m.InterfaceExclusiveCounters
	}
	return nil
}

type PmDelayCountersInfoBag struct {
	GenericCounters      *PmQuerierDelayCountersBag         `protobuf:"bytes,1,opt,name=generic_counters,json=genericCounters,proto3" json:"generic_counters,omitempty"`
	ExclusiveCounters    *PmExclusiveDelayCountersInfoUnion `protobuf:"bytes,2,opt,name=exclusive_counters,json=exclusiveCounters,proto3" json:"exclusive_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *PmDelayCountersInfoBag) Reset()         { *m = PmDelayCountersInfoBag{} }
func (m *PmDelayCountersInfoBag) String() string { return proto.CompactTextString(m) }
func (*PmDelayCountersInfoBag) ProtoMessage()    {}
func (*PmDelayCountersInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{5}
}

func (m *PmDelayCountersInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmDelayCountersInfoBag.Unmarshal(m, b)
}
func (m *PmDelayCountersInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmDelayCountersInfoBag.Marshal(b, m, deterministic)
}
func (m *PmDelayCountersInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmDelayCountersInfoBag.Merge(m, src)
}
func (m *PmDelayCountersInfoBag) XXX_Size() int {
	return xxx_messageInfo_PmDelayCountersInfoBag.Size(m)
}
func (m *PmDelayCountersInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmDelayCountersInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmDelayCountersInfoBag proto.InternalMessageInfo

func (m *PmDelayCountersInfoBag) GetGenericCounters() *PmQuerierDelayCountersBag {
	if m != nil {
		return m.GenericCounters
	}
	return nil
}

func (m *PmDelayCountersInfoBag) GetExclusiveCounters() *PmExclusiveDelayCountersInfoUnion {
	if m != nil {
		return m.ExclusiveCounters
	}
	return nil
}

type PmDelayTransportSummaryBag struct {
	DelayProfile           *PmDelayProfileBag      `protobuf:"bytes,1,opt,name=delay_profile,json=delayProfile,proto3" json:"delay_profile,omitempty"`
	TotalDelaySessions     uint32                  `protobuf:"varint,2,opt,name=total_delay_sessions,json=totalDelaySessions,proto3" json:"total_delay_sessions,omitempty"`
	DelayTransportCounters *PmDelayCountersInfoBag `protobuf:"bytes,3,opt,name=delay_transport_counters,json=delayTransportCounters,proto3" json:"delay_transport_counters,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *PmDelayTransportSummaryBag) Reset()         { *m = PmDelayTransportSummaryBag{} }
func (m *PmDelayTransportSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PmDelayTransportSummaryBag) ProtoMessage()    {}
func (*PmDelayTransportSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{6}
}

func (m *PmDelayTransportSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmDelayTransportSummaryBag.Unmarshal(m, b)
}
func (m *PmDelayTransportSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmDelayTransportSummaryBag.Marshal(b, m, deterministic)
}
func (m *PmDelayTransportSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmDelayTransportSummaryBag.Merge(m, src)
}
func (m *PmDelayTransportSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PmDelayTransportSummaryBag.Size(m)
}
func (m *PmDelayTransportSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmDelayTransportSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmDelayTransportSummaryBag proto.InternalMessageInfo

func (m *PmDelayTransportSummaryBag) GetDelayProfile() *PmDelayProfileBag {
	if m != nil {
		return m.DelayProfile
	}
	return nil
}

func (m *PmDelayTransportSummaryBag) GetTotalDelaySessions() uint32 {
	if m != nil {
		return m.TotalDelaySessions
	}
	return 0
}

func (m *PmDelayTransportSummaryBag) GetDelayTransportCounters() *PmDelayCountersInfoBag {
	if m != nil {
		return m.DelayTransportCounters
	}
	return nil
}

type PmDelayGlobalCountersBag struct {
	QueryPacketsSent                    uint64   `protobuf:"varint,1,opt,name=query_packets_sent,json=queryPacketsSent,proto3" json:"query_packets_sent,omitempty"`
	QueryPacketsReceived                uint64   `protobuf:"varint,2,opt,name=query_packets_received,json=queryPacketsReceived,proto3" json:"query_packets_received,omitempty"`
	ReceivedPacketErrorInvalidSessionId uint64   `protobuf:"varint,3,opt,name=received_packet_error_invalid_session_id,json=receivedPacketErrorInvalidSessionId,proto3" json:"received_packet_error_invalid_session_id,omitempty"`
	ReceivedPacketErrorNoSession        uint64   `protobuf:"varint,4,opt,name=received_packet_error_no_session,json=receivedPacketErrorNoSession,proto3" json:"received_packet_error_no_session,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *PmDelayGlobalCountersBag) Reset()         { *m = PmDelayGlobalCountersBag{} }
func (m *PmDelayGlobalCountersBag) String() string { return proto.CompactTextString(m) }
func (*PmDelayGlobalCountersBag) ProtoMessage()    {}
func (*PmDelayGlobalCountersBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{7}
}

func (m *PmDelayGlobalCountersBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmDelayGlobalCountersBag.Unmarshal(m, b)
}
func (m *PmDelayGlobalCountersBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmDelayGlobalCountersBag.Marshal(b, m, deterministic)
}
func (m *PmDelayGlobalCountersBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmDelayGlobalCountersBag.Merge(m, src)
}
func (m *PmDelayGlobalCountersBag) XXX_Size() int {
	return xxx_messageInfo_PmDelayGlobalCountersBag.Size(m)
}
func (m *PmDelayGlobalCountersBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmDelayGlobalCountersBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmDelayGlobalCountersBag proto.InternalMessageInfo

func (m *PmDelayGlobalCountersBag) GetQueryPacketsSent() uint64 {
	if m != nil {
		return m.QueryPacketsSent
	}
	return 0
}

func (m *PmDelayGlobalCountersBag) GetQueryPacketsReceived() uint64 {
	if m != nil {
		return m.QueryPacketsReceived
	}
	return 0
}

func (m *PmDelayGlobalCountersBag) GetReceivedPacketErrorInvalidSessionId() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorInvalidSessionId
	}
	return 0
}

func (m *PmDelayGlobalCountersBag) GetReceivedPacketErrorNoSession() uint64 {
	if m != nil {
		return m.ReceivedPacketErrorNoSession
	}
	return 0
}

type PmDelaySummaryBag struct {
	InterfaceDelaySummary *PmDelayTransportSummaryBag `protobuf:"bytes,1,opt,name=interface_delay_summary,json=interfaceDelaySummary,proto3" json:"interface_delay_summary,omitempty"`
	DelayGlobalCounters   *PmDelayGlobalCountersBag   `protobuf:"bytes,2,opt,name=delay_global_counters,json=delayGlobalCounters,proto3" json:"delay_global_counters,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                    `json:"-"`
	XXX_unrecognized      []byte                      `json:"-"`
	XXX_sizecache         int32                       `json:"-"`
}

func (m *PmDelaySummaryBag) Reset()         { *m = PmDelaySummaryBag{} }
func (m *PmDelaySummaryBag) String() string { return proto.CompactTextString(m) }
func (*PmDelaySummaryBag) ProtoMessage()    {}
func (*PmDelaySummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{8}
}

func (m *PmDelaySummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmDelaySummaryBag.Unmarshal(m, b)
}
func (m *PmDelaySummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmDelaySummaryBag.Marshal(b, m, deterministic)
}
func (m *PmDelaySummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmDelaySummaryBag.Merge(m, src)
}
func (m *PmDelaySummaryBag) XXX_Size() int {
	return xxx_messageInfo_PmDelaySummaryBag.Size(m)
}
func (m *PmDelaySummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmDelaySummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmDelaySummaryBag proto.InternalMessageInfo

func (m *PmDelaySummaryBag) GetInterfaceDelaySummary() *PmDelayTransportSummaryBag {
	if m != nil {
		return m.InterfaceDelaySummary
	}
	return nil
}

func (m *PmDelaySummaryBag) GetDelayGlobalCounters() *PmDelayGlobalCountersBag {
	if m != nil {
		return m.DelayGlobalCounters
	}
	return nil
}

type PmSummaryBag struct {
	TotalInterfaces      uint32             `protobuf:"varint,50,opt,name=total_interfaces,json=totalInterfaces,proto3" json:"total_interfaces,omitempty"`
	DelaySummary         *PmDelaySummaryBag `protobuf:"bytes,51,opt,name=delay_summary,json=delaySummary,proto3" json:"delay_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PmSummaryBag) Reset()         { *m = PmSummaryBag{} }
func (m *PmSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PmSummaryBag) ProtoMessage()    {}
func (*PmSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_28e0a2aa83a9d7d3, []int{9}
}

func (m *PmSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PmSummaryBag.Unmarshal(m, b)
}
func (m *PmSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PmSummaryBag.Marshal(b, m, deterministic)
}
func (m *PmSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PmSummaryBag.Merge(m, src)
}
func (m *PmSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PmSummaryBag.Size(m)
}
func (m *PmSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PmSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PmSummaryBag proto.InternalMessageInfo

func (m *PmSummaryBag) GetTotalInterfaces() uint32 {
	if m != nil {
		return m.TotalInterfaces
	}
	return 0
}

func (m *PmSummaryBag) GetDelaySummary() *PmDelaySummaryBag {
	if m != nil {
		return m.DelaySummary
	}
	return nil
}

func init() {
	proto.RegisterType((*PmSummaryBag_KEYS)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_summary_bag_KEYS")
	proto.RegisterType((*PmDelayProfileBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_delay_profile_bag")
	proto.RegisterType((*PmQuerierDelayCountersBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_querier_delay_counters_bag")
	proto.RegisterType((*PmIntfExclusiveDelayCountersInfoBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_intf_exclusive_delay_counters_info_bag")
	proto.RegisterType((*PmExclusiveDelayCountersInfoUnion)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_exclusive_delay_counters_info_union")
	proto.RegisterType((*PmDelayCountersInfoBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_delay_counters_info_bag")
	proto.RegisterType((*PmDelayTransportSummaryBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_delay_transport_summary_bag")
	proto.RegisterType((*PmDelayGlobalCountersBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_delay_global_counters_bag")
	proto.RegisterType((*PmDelaySummaryBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_delay_summary_bag")
	proto.RegisterType((*PmSummaryBag)(nil), "cisco_ios_xr_perf_meas_oper.performance_measurement.nodes.node.summary.pm_summary_bag")
}

func init() { proto.RegisterFile("pm_summary_bag.proto", fileDescriptor_28e0a2aa83a9d7d3) }

var fileDescriptor_28e0a2aa83a9d7d3 = []byte{
	// 1239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdf, 0x6f, 0xdc, 0xc4,
	0x13, 0x97, 0x53, 0x7f, 0xf3, 0x4d, 0xa7, 0x4d, 0x93, 0x6e, 0x53, 0x38, 0x55, 0x69, 0x92, 0x5e,
	0x9b, 0x36, 0xe1, 0x47, 0x84, 0x5a, 0x24, 0x9e, 0xab, 0x6b, 0x02, 0x29, 0xe4, 0x94, 0x3a, 0x41,
	0x50, 0x09, 0x69, 0xe5, 0xb3, 0xe7, 0xd2, 0x15, 0xf6, 0xae, 0xbb, 0xbb, 0x97, 0x26, 0x8f, 0xfc,
	0x07, 0x3c, 0xf2, 0x00, 0x48, 0xfc, 0x0d, 0x48, 0xbc, 0xf2, 0x0f, 0x20, 0x21, 0xf1, 0xc6, 0xbf,
	0x81, 0xc4, 0x3b, 0xf2, 0xac, 0xed, 0xf3, 0x35, 0x76, 0x52, 0x89, 0xe3, 0x25, 0xba, 0xec, 0x7c,
	0xe6, 0xe3, 0xcf, 0x8c, 0x67, 0x66, 0xc7, 0xb0, 0x94, 0xa5, 0xdc, 0x8c, 0xd2, 0x34, 0xd4, 0xa7,
	0x7c, 0x10, 0x1e, 0x6d, 0x65, 0x5a, 0x59, 0xc5, 0x76, 0x22, 0x61, 0x22, 0xc5, 0x85, 0x32, 0xfc,
	0x44, 0xf3, 0x0c, 0xf5, 0x90, 0xa7, 0x18, 0x1a, 0xae, 0x32, 0xd4, 0x5b, 0xf9, 0xbf, 0x4a, 0xa7,
	0xa1, 0x8c, 0x90, 0x4e, 0x47, 0x1a, 0x53, 0x94, 0x76, 0x4b, 0xaa, 0x18, 0x0d, 0xfd, 0xdd, 0x2a,
	0x18, 0xbb, 0x9b, 0x70, 0x63, 0x92, 0x9f, 0x7f, 0xba, 0xfd, 0xfc, 0x80, 0x31, 0xf0, 0x73, 0x58,
	0xc7, 0x5b, 0xf3, 0x36, 0x2e, 0x07, 0xf4, 0xbb, 0xfb, 0xcb, 0x2c, 0x69, 0x89, 0x31, 0x09, 0x4f,
	0x79, 0xa6, 0xd5, 0x50, 0x24, 0x98, 0x7b, 0xb0, 0x75, 0xb8, 0x96, 0x69, 0x35, 0x40, 0x2e, 0xa4,
	0x45, 0x7d, 0x1c, 0x26, 0xe4, 0x36, 0x1f, 0xcc, 0xd3, 0xe9, 0x6e, 0x71, 0x98, 0xc3, 0x06, 0x23,
	0x6d, 0xec, 0x18, 0x36, 0xe3, 0x60, 0x74, 0x5a, 0xc1, 0x56, 0xe1, 0x8a, 0x83, 0x45, 0x6a, 0x24,
	0x6d, 0xe7, 0x12, 0x61, 0x80, 0x8e, 0x7a, 0xf9, 0x09, 0xfb, 0x10, 0xde, 0x72, 0x1a, 0x6a, 0xa1,
	0xf1, 0x34, 0x57, 0xeb, 0x93, 0xda, 0x25, 0xb2, 0xee, 0x8d, 0x8d, 0x7b, 0x2a, 0x46, 0xf6, 0x04,
	0x56, 0x32, 0xd4, 0x42, 0xc5, 0x22, 0xe2, 0x61, 0x7c, 0x8c, 0xda, 0x0a, 0xe3, 0x5c, 0x51, 0x86,
	0x83, 0x04, 0xe3, 0xce, 0xff, 0xd6, 0xbc, 0x8d, 0xb9, 0x60, 0xb9, 0x44, 0x3d, 0xae, 0x83, 0xb6,
	0x1d, 0x86, 0xed, 0xc0, 0x6a, 0x0b, 0x4b, 0x15, 0xd4, 0x2c, 0x09, 0xbe, 0xdd, 0x48, 0x53, 0x05,
	0xf9, 0x1c, 0x36, 0x71, 0x38, 0xc4, 0xc8, 0x8a, 0x63, 0xe4, 0x17, 0x31, 0xfe, 0x9f, 0x18, 0xef,
	0x57, 0x0e, 0xfb, 0xe7, 0x52, 0x7f, 0x02, 0x6b, 0x2d, 0x84, 0xf6, 0x85, 0x46, 0xf3, 0x42, 0x25,
	0x71, 0x67, 0x8e, 0x18, 0x57, 0x1a, 0x35, 0x1e, 0x96, 0x28, 0xf6, 0x0c, 0xd6, 0x5b, 0x98, 0x52,
	0x21, 0x45, 0x3a, 0x4a, 0x79, 0xf4, 0x22, 0x94, 0x47, 0xd8, 0xb9, 0x4c, 0x74, 0xdd, 0x46, 0xba,
	0x3d, 0x07, 0xed, 0x11, 0x92, 0xf5, 0xe1, 0x6e, 0x18, 0x45, 0x98, 0xa0, 0x0e, 0x2d, 0xc6, 0xad,
	0xfa, 0x80, 0x08, 0xef, 0xd4, 0xa0, 0x2d, 0x12, 0xbf, 0x80, 0x8d, 0x76, 0xbe, 0xd7, 0x54, 0x5e,
	0x21, 0xd2, 0xf5, 0x36, 0xd2, 0x49, 0xa1, 0x4f, 0xe1, 0x4e, 0x3b, 0x71, 0x59, 0x31, 0x57, 0xa9,
	0x62, 0x56, 0xdb, 0x18, 0x8b, 0xa2, 0xe9, 0xfe, 0x39, 0x07, 0xb7, 0xb3, 0x94, 0xbf, 0x1c, 0xa1,
	0x16, 0xa8, 0x8b, 0x06, 0xa2, 0xea, 0x46, 0x6d, 0xa8, 0x83, 0xde, 0x03, 0x96, 0x5b, 0x4f, 0x79,
	0x16, 0x46, 0x5f, 0xa3, 0x35, 0xdc, 0xa0, 0xb4, 0xd4, 0x45, 0x7e, 0xb0, 0x48, 0x96, 0x7d, 0x67,
	0x38, 0x40, 0x69, 0xd9, 0x47, 0xd0, 0xa9, 0xa3, 0x09, 0xcc, 0x51, 0x6b, 0xa5, 0x0d, 0xb5, 0x94,
	0x1f, 0xdc, 0xac, 0xf9, 0xe4, 0x2e, 0xdb, 0x64, 0x64, 0x7d, 0x58, 0x6f, 0x71, 0xe4, 0x52, 0x71,
	0x91, 0xf1, 0x30, 0x8e, 0x35, 0x1a, 0x43, 0x4d, 0xe7, 0x07, 0xab, 0x4d, 0x2c, 0x7d, 0xb5, 0x9b,
	0x3d, 0x76, 0xb0, 0xbc, 0x13, 0x27, 0x65, 0x6b, 0x8c, 0x50, 0x1c, 0x63, 0x4c, 0x9d, 0xe8, 0x07,
	0x4b, 0x75, 0xe9, 0x41, 0x61, 0x63, 0x7d, 0xb8, 0x57, 0xe2, 0x4a, 0x21, 0x85, 0x06, 0x3c, 0x0a,
	0xa9, 0x21, 0x28, 0x43, 0xd4, 0x8f, 0x7e, 0xb0, 0x56, 0x62, 0x1d, 0x8d, 0xd3, 0x50, 0x00, 0x9f,
	0xe4, 0x38, 0xc6, 0xe1, 0xfd, 0x66, 0x3e, 0x97, 0x68, 0x3c, 0x89, 0x10, 0x63, 0x53, 0xab, 0xae,
	0x59, 0x22, 0xde, 0x68, 0x20, 0x26, 0xc2, 0x6d, 0xe7, 0x30, 0x2e, 0xb2, 0xe7, 0xf0, 0x4e, 0xf3,
	0x03, 0x52, 0x61, 0x8c, 0x90, 0x47, 0xdc, 0x9e, 0x70, 0x2b, 0x52, 0x34, 0x36, 0x4c, 0x33, 0xea,
	0x56, 0x3f, 0x58, 0x6f, 0x60, 0xdf, 0x73, 0xf0, 0xc3, 0x93, 0xc3, 0x12, 0x7c, 0x31, 0xb5, 0xae,
	0x53, 0xcf, 0x5d, 0x44, 0x1d, 0xd4, 0xa8, 0x77, 0x60, 0xad, 0x99, 0xda, 0xcd, 0xea, 0xe1, 0x28,
	0x49, 0xa8, 0x71, 0xfd, 0x60, 0xb9, 0x81, 0x70, 0x3f, 0x07, 0xed, 0x8c, 0x92, 0x84, 0x1d, 0xc0,
	0x83, 0xf3, 0x78, 0xa4, 0xb2, 0xdc, 0xd8, 0x50, 0x5b, 0x74, 0x6d, 0xeb, 0x07, 0xdd, 0x36, 0xba,
	0xbe, 0xb2, 0x07, 0x0e, 0xc9, 0x9e, 0x42, 0xf7, 0x75, 0xd2, 0x48, 0x49, 0xab, 0x55, 0xc2, 0x23,
	0x15, 0xa3, 0x7b, 0x02, 0x75, 0xac, 0x1f, 0xac, 0x4c, 0xf2, 0xf5, 0x1c, 0xae, 0xa7, 0x62, 0x24,
	0x6a, 0xf6, 0x25, 0x6c, 0x9e, 0xcb, 0x25, 0x95, 0x15, 0x43, 0x11, 0x85, 0x56, 0x28, 0x49, 0x2d,
	0x7b, 0x26, 0x85, 0x35, 0xca, 0x7e, 0x0d, 0x5c, 0x5d, 0x6c, 0xa6, 0x8a, 0x70, 0x9e, 0xdc, 0xdd,
	0xc5, 0x66, 0xca, 0x60, 0x1e, 0xc0, 0x42, 0x01, 0x8b, 0x54, 0x9a, 0x25, 0x68, 0xb1, 0x73, 0x8d,
	0x70, 0x85, 0x77, 0xaf, 0x38, 0x65, 0xef, 0xc2, 0xf5, 0x02, 0x28, 0x64, 0x05, 0x5d, 0x70, 0x5d,
	0xee, 0x0c, 0xbb, 0xd5, 0x39, 0xbb, 0x07, 0xf3, 0x13, 0x53, 0xa7, 0xb3, 0xe8, 0x9e, 0x3d, 0x71,
	0xd8, 0xfd, 0xcd, 0x83, 0xcd, 0x2c, 0xcd, 0xef, 0x8a, 0x61, 0x5e, 0xe9, 0xc9, 0xc8, 0x54, 0x0d,
	0x34, 0x1e, 0x31, 0x42, 0x0e, 0x15, 0xcd, 0x99, 0x67, 0x70, 0xbf, 0x6d, 0x00, 0xd0, 0x6d, 0x33,
	0x0c, 0x23, 0xe4, 0xb1, 0x7a, 0x25, 0x8b, 0xd9, 0x73, 0xa7, 0x69, 0x02, 0xec, 0x96, 0xc8, 0x27,
	0xea, 0x95, 0x64, 0x9f, 0xc1, 0xbd, 0x73, 0x66, 0x4a, 0x9a, 0x25, 0x86, 0x47, 0x61, 0x56, 0x0e,
	0xa6, 0x95, 0xe6, 0x91, 0xb2, 0x97, 0x25, 0xa6, 0x17, 0x66, 0xa6, 0xfb, 0x97, 0x07, 0xf7, 0xb3,
	0xf4, 0x82, 0x48, 0x46, 0x32, 0x7f, 0x39, 0x0c, 0x7c, 0x7b, 0x9a, 0x55, 0x2b, 0x4a, 0xfe, 0x9b,
	0xfd, 0xec, 0xc1, 0xf2, 0x38, 0x90, 0x31, 0x4b, 0xe9, 0x4f, 0x2a, 0xae, 0x3c, 0x7c, 0xb9, 0x35,
	0x9d, 0xed, 0x69, 0xeb, 0x8d, 0x33, 0x1f, 0xdc, 0xaa, 0x64, 0x6d, 0x97, 0xd8, 0x5e, 0x01, 0xea,
	0xfe, 0x3e, 0x03, 0xb7, 0xaa, 0xc5, 0xea, 0xec, 0x4b, 0xfb, 0xd6, 0x83, 0xc5, 0x23, 0x94, 0xa8,
	0x45, 0x34, 0x0e, 0xc4, 0xa3, 0x40, 0x70, 0x8a, 0x81, 0xb4, 0x5f, 0x4f, 0xc1, 0x42, 0xf1, 0xf8,
	0x52, 0x31, 0xfb, 0xc1, 0x03, 0xd6, 0x9a, 0x5d, 0x39, 0x45, 0x51, 0x6f, 0x50, 0x08, 0xc1, 0x75,
	0x3c, 0x93, 0xd1, 0xbf, 0x67, 0x60, 0xa5, 0xca, 0xa8, 0xd5, 0xa1, 0x34, 0x99, 0xd2, 0xb6, 0xbe,
	0xe6, 0xb2, 0x6f, 0x3c, 0x98, 0x9f, 0x58, 0x65, 0x8b, 0x94, 0x7e, 0x35, 0x45, 0xf5, 0x67, 0x56,
	0xe5, 0xe0, 0x2a, 0x1d, 0xed, 0xbb, 0x13, 0xf6, 0x01, 0x2c, 0x59, 0x65, 0xc3, 0xa4, 0x00, 0x1a,
	0x34, 0x46, 0x28, 0x69, 0x8a, 0xbd, 0x98, 0x91, 0x8d, 0xae, 0xa4, 0x83, 0xc2, 0xc2, 0xbe, 0xf7,
	0xa0, 0xf3, 0x7a, 0x54, 0x55, 0xfa, 0x2f, 0x51, 0x00, 0x83, 0xa9, 0x07, 0x70, 0xb6, 0x9a, 0xdd,
	0x02, 0x7e, 0x58, 0x4a, 0xa8, 0xf2, 0xfe, 0xd3, 0x0c, 0x2c, 0x57, 0x6e, 0x47, 0x89, 0x1a, 0x84,
	0xc9, 0xbf, 0x59, 0x74, 0xda, 0xf7, 0x8b, 0x99, 0x73, 0xf6, 0x8b, 0xcf, 0x61, 0xa3, 0xf9, 0xc2,
	0x12, 0xf2, 0x38, 0x4c, 0x44, 0x5c, 0xe6, 0x99, 0x8b, 0xb8, 0x58, 0x74, 0xee, 0x36, 0xdc, 0x58,
	0xbb, 0x0e, 0x5c, 0x64, 0x7e, 0x37, 0x6e, 0xbf, 0x4f, 0xa5, 0x2a, 0x19, 0x8b, 0xb5, 0xa7, 0xe9,
	0x3e, 0xed, 0xab, 0x82, 0xa9, 0xfb, 0xc7, 0x4c, 0xed, 0x33, 0xaa, 0x5e, 0x91, 0x3f, 0x7a, 0xf0,
	0x76, 0x6d, 0x0a, 0xd7, 0xed, 0x45, 0x6d, 0x0e, 0xa7, 0xfe, 0x6a, 0x1b, 0x7b, 0x23, 0xb8, 0x59,
	0xc9, 0x70, 0xd5, 0xe7, 0x4c, 0xec, 0x3b, 0x0f, 0x6e, 0x36, 0xbe, 0xda, 0xa2, 0xf1, 0xe3, 0xa9,
	0xcb, 0x6b, 0x28, 0xa1, 0xe0, 0x06, 0x99, 0x3e, 0x26, 0x4b, 0x55, 0x78, 0xbf, 0x7a, 0x70, 0x6d,
	0xf2, 0x3b, 0x96, 0x6d, 0xc2, 0xa2, 0x6b, 0xae, 0x2a, 0x18, 0xd3, 0x79, 0x48, 0x8d, 0xb5, 0x40,
	0xe7, 0xd5, 0x35, 0x66, 0x6a, 0xb3, 0xa0, 0xcc, 0xf7, 0xa3, 0xff, 0x68, 0x16, 0xd4, 0xb3, 0xec,
	0x66, 0x41, 0x91, 0xdc, 0xc1, 0x2c, 0x7d, 0xd7, 0x3f, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xf4,
	0x96, 0x3d, 0x6f, 0xef, 0x0f, 0x00, 0x00,
}
