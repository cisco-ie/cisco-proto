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
// source: mpls_lm_bandwidth_account_link_tel_det_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_standby_bandwidth_account_bandwidth_account_links_bandwidth_account_link

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

type MplsLmBandwidthAccountLinkTelDetInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) Reset() {
	*m = MplsLmBandwidthAccountLinkTelDetInfo_KEYS{}
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmBandwidthAccountLinkTelDetInfo_KEYS) ProtoMessage()    {}
func (*MplsLmBandwidthAccountLinkTelDetInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{0}
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS.Merge(m, src)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS.Size(m)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo_KEYS proto.InternalMessageInfo

func (m *MplsLmBandwidthAccountLinkTelDetInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsTeBandwidthAccountBwUtilRsvp struct {
	TotalLinkBandwidthUtilization      uint64   `protobuf:"varint,1,opt,name=total_link_bandwidth_utilization,json=totalLinkBandwidthUtilization,proto3" json:"total_link_bandwidth_utilization,omitempty"`
	RsvpTeBandwidthUtilization         uint64   `protobuf:"varint,2,opt,name=rsvp_te_bandwidth_utilization,json=rsvpTeBandwidthUtilization,proto3" json:"rsvp_te_bandwidth_utilization,omitempty"`
	NonRsvpTeBandwidthUtilization      uint64   `protobuf:"varint,3,opt,name=non_rsvp_te_bandwidth_utilization,json=nonRsvpTeBandwidthUtilization,proto3" json:"non_rsvp_te_bandwidth_utilization,omitempty"`
	RsvpTeAdjustedBandwidthUtilization uint64   `protobuf:"varint,4,opt,name=rsvp_te_adjusted_bandwidth_utilization,json=rsvpTeAdjustedBandwidthUtilization,proto3" json:"rsvp_te_adjusted_bandwidth_utilization,omitempty"`
	RsvpTeEnforcedBandwidthUtilization uint64   `protobuf:"varint,5,opt,name=rsvp_te_enforced_bandwidth_utilization,json=rsvpTeEnforcedBandwidthUtilization,proto3" json:"rsvp_te_enforced_bandwidth_utilization,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) Reset()         { *m = MplsTeBandwidthAccountBwUtilRsvp{} }
func (m *MplsTeBandwidthAccountBwUtilRsvp) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountBwUtilRsvp) ProtoMessage()    {}
func (*MplsTeBandwidthAccountBwUtilRsvp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{1}
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountBwUtilRsvp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountBwUtilRsvp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp.Merge(m, src)
}
func (m *MplsTeBandwidthAccountBwUtilRsvp) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp.Size(m)
}
func (m *MplsTeBandwidthAccountBwUtilRsvp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountBwUtilRsvp proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountBwUtilRsvp) GetTotalLinkBandwidthUtilization() uint64 {
	if m != nil {
		return m.TotalLinkBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) GetRsvpTeBandwidthUtilization() uint64 {
	if m != nil {
		return m.RsvpTeBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) GetNonRsvpTeBandwidthUtilization() uint64 {
	if m != nil {
		return m.NonRsvpTeBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) GetRsvpTeAdjustedBandwidthUtilization() uint64 {
	if m != nil {
		return m.RsvpTeAdjustedBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilRsvp) GetRsvpTeEnforcedBandwidthUtilization() uint64 {
	if m != nil {
		return m.RsvpTeEnforcedBandwidthUtilization
	}
	return 0
}

type MplsTeBandwidthAccountBwUtilSr struct {
	SrBandwidthUtilization         uint64   `protobuf:"varint,1,opt,name=sr_bandwidth_utilization,json=srBandwidthUtilization,proto3" json:"sr_bandwidth_utilization,omitempty"`
	SrAdjustedBandwidthUtilization uint64   `protobuf:"varint,2,opt,name=sr_adjusted_bandwidth_utilization,json=srAdjustedBandwidthUtilization,proto3" json:"sr_adjusted_bandwidth_utilization,omitempty"`
	SrEnforcedBandwidthUtilization uint64   `protobuf:"varint,3,opt,name=sr_enforced_bandwidth_utilization,json=srEnforcedBandwidthUtilization,proto3" json:"sr_enforced_bandwidth_utilization,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *MplsTeBandwidthAccountBwUtilSr) Reset()         { *m = MplsTeBandwidthAccountBwUtilSr{} }
func (m *MplsTeBandwidthAccountBwUtilSr) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountBwUtilSr) ProtoMessage()    {}
func (*MplsTeBandwidthAccountBwUtilSr) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{2}
}

func (m *MplsTeBandwidthAccountBwUtilSr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountBwUtilSr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountBwUtilSr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr.Merge(m, src)
}
func (m *MplsTeBandwidthAccountBwUtilSr) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr.Size(m)
}
func (m *MplsTeBandwidthAccountBwUtilSr) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountBwUtilSr proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountBwUtilSr) GetSrBandwidthUtilization() uint64 {
	if m != nil {
		return m.SrBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilSr) GetSrAdjustedBandwidthUtilization() uint64 {
	if m != nil {
		return m.SrAdjustedBandwidthUtilization
	}
	return 0
}

func (m *MplsTeBandwidthAccountBwUtilSr) GetSrEnforcedBandwidthUtilization() uint64 {
	if m != nil {
		return m.SrEnforcedBandwidthUtilization
	}
	return 0
}

type MplsLmBandwidthAccountLinkTelemetryInfo struct {
	IsBandwidthAccountEnabled            bool                              `protobuf:"varint,1,opt,name=is_bandwidth_account_enabled,json=isBandwidthAccountEnabled,proto3" json:"is_bandwidth_account_enabled,omitempty"`
	ApplicationEnforced                  bool                              `protobuf:"varint,2,opt,name=application_enforced,json=applicationEnforced,proto3" json:"application_enforced,omitempty"`
	CollectionType                       string                            `protobuf:"bytes,3,opt,name=collection_type,json=collectionType,proto3" json:"collection_type,omitempty"`
	SampleTimeRemainingTimestampNanosec  uint64                            `protobuf:"varint,4,opt,name=sample_time_remaining_timestamp_nanosec,json=sampleTimeRemainingTimestampNanosec,proto3" json:"sample_time_remaining_timestamp_nanosec,omitempty"`
	LastSampleCollectionTimestampNanosec uint64                            `protobuf:"varint,5,opt,name=last_sample_collection_timestamp_nanosec,json=lastSampleCollectionTimestampNanosec,proto3" json:"last_sample_collection_timestamp_nanosec,omitempty"`
	NextSampleCollectionNanosec          uint64                            `protobuf:"varint,6,opt,name=next_sample_collection_nanosec,json=nextSampleCollectionNanosec,proto3" json:"next_sample_collection_nanosec,omitempty"`
	ApplicationTimeRemainingNanosec      uint64                            `protobuf:"varint,7,opt,name=application_time_remaining_nanosec,json=applicationTimeRemainingNanosec,proto3" json:"application_time_remaining_nanosec,omitempty"`
	LastApplicationTimestampNanosec      uint64                            `protobuf:"varint,8,opt,name=last_application_timestamp_nanosec,json=lastApplicationTimestampNanosec,proto3" json:"last_application_timestamp_nanosec,omitempty"`
	RsvpTeBandwidthUtilization           *MplsTeBandwidthAccountBwUtilRsvp `protobuf:"bytes,9,opt,name=rsvp_te_bandwidth_utilization,json=rsvpTeBandwidthUtilization,proto3" json:"rsvp_te_bandwidth_utilization,omitempty"`
	SrBandwidthUtilization               *MplsTeBandwidthAccountBwUtilSr   `protobuf:"bytes,10,opt,name=sr_bandwidth_utilization,json=srBandwidthUtilization,proto3" json:"sr_bandwidth_utilization,omitempty"`
	NextApplicationTimestampNanosec      uint64                            `protobuf:"varint,11,opt,name=next_application_timestamp_nanosec,json=nextApplicationTimestampNanosec,proto3" json:"next_application_timestamp_nanosec,omitempty"`
	EffectiveMaximumReservableBandwidth  uint64                            `protobuf:"varint,12,opt,name=effective_maximum_reservable_bandwidth,json=effectiveMaximumReservableBandwidth,proto3" json:"effective_maximum_reservable_bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{}                          `json:"-"`
	XXX_unrecognized                     []byte                            `json:"-"`
	XXX_sizecache                        int32                             `json:"-"`
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) Reset() {
	*m = MplsLmBandwidthAccountLinkTelemetryInfo{}
}
func (m *MplsLmBandwidthAccountLinkTelemetryInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBandwidthAccountLinkTelemetryInfo) ProtoMessage()    {}
func (*MplsLmBandwidthAccountLinkTelemetryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{3}
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo.Unmarshal(m, b)
}
func (m *MplsLmBandwidthAccountLinkTelemetryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmBandwidthAccountLinkTelemetryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo.Merge(m, src)
}
func (m *MplsLmBandwidthAccountLinkTelemetryInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo.Size(m)
}
func (m *MplsLmBandwidthAccountLinkTelemetryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBandwidthAccountLinkTelemetryInfo proto.InternalMessageInfo

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetIsBandwidthAccountEnabled() bool {
	if m != nil {
		return m.IsBandwidthAccountEnabled
	}
	return false
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetApplicationEnforced() bool {
	if m != nil {
		return m.ApplicationEnforced
	}
	return false
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetCollectionType() string {
	if m != nil {
		return m.CollectionType
	}
	return ""
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetSampleTimeRemainingTimestampNanosec() uint64 {
	if m != nil {
		return m.SampleTimeRemainingTimestampNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetLastSampleCollectionTimestampNanosec() uint64 {
	if m != nil {
		return m.LastSampleCollectionTimestampNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetNextSampleCollectionNanosec() uint64 {
	if m != nil {
		return m.NextSampleCollectionNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetApplicationTimeRemainingNanosec() uint64 {
	if m != nil {
		return m.ApplicationTimeRemainingNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetLastApplicationTimestampNanosec() uint64 {
	if m != nil {
		return m.LastApplicationTimestampNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetRsvpTeBandwidthUtilization() *MplsTeBandwidthAccountBwUtilRsvp {
	if m != nil {
		return m.RsvpTeBandwidthUtilization
	}
	return nil
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetSrBandwidthUtilization() *MplsTeBandwidthAccountBwUtilSr {
	if m != nil {
		return m.SrBandwidthUtilization
	}
	return nil
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetNextApplicationTimestampNanosec() uint64 {
	if m != nil {
		return m.NextApplicationTimestampNanosec
	}
	return 0
}

func (m *MplsLmBandwidthAccountLinkTelemetryInfo) GetEffectiveMaximumReservableBandwidth() uint64 {
	if m != nil {
		return m.EffectiveMaximumReservableBandwidth
	}
	return 0
}

type MplsTeBandwidthAccountSampleRsvp struct {
	TimestampNanosec     uint64   `protobuf:"varint,1,opt,name=timestamp_nanosec,json=timestampNanosec,proto3" json:"timestamp_nanosec,omitempty"`
	TotalRate            uint64   `protobuf:"varint,2,opt,name=total_rate,json=totalRate,proto3" json:"total_rate,omitempty"`
	TotalPacketRate      uint64   `protobuf:"varint,3,opt,name=total_packet_rate,json=totalPacketRate,proto3" json:"total_packet_rate,omitempty"`
	RsvpTeRate           uint64   `protobuf:"varint,4,opt,name=rsvp_te_rate,json=rsvpTeRate,proto3" json:"rsvp_te_rate,omitempty"`
	RsvpTePacketRate     uint64   `protobuf:"varint,5,opt,name=rsvp_te_packet_rate,json=rsvpTePacketRate,proto3" json:"rsvp_te_packet_rate,omitempty"`
	NonRsvpTeRate        uint64   `protobuf:"varint,6,opt,name=non_rsvp_te_rate,json=nonRsvpTeRate,proto3" json:"non_rsvp_te_rate,omitempty"`
	NonRsvpTePacketRate  uint64   `protobuf:"varint,7,opt,name=non_rsvp_te_packet_rate,json=nonRsvpTePacketRate,proto3" json:"non_rsvp_te_packet_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBandwidthAccountSampleRsvp) Reset()         { *m = MplsTeBandwidthAccountSampleRsvp{} }
func (m *MplsTeBandwidthAccountSampleRsvp) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountSampleRsvp) ProtoMessage()    {}
func (*MplsTeBandwidthAccountSampleRsvp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{4}
}

func (m *MplsTeBandwidthAccountSampleRsvp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountSampleRsvp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountSampleRsvp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp.Merge(m, src)
}
func (m *MplsTeBandwidthAccountSampleRsvp) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp.Size(m)
}
func (m *MplsTeBandwidthAccountSampleRsvp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountSampleRsvp proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountSampleRsvp) GetTimestampNanosec() uint64 {
	if m != nil {
		return m.TimestampNanosec
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetTotalRate() uint64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetTotalPacketRate() uint64 {
	if m != nil {
		return m.TotalPacketRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetRsvpTeRate() uint64 {
	if m != nil {
		return m.RsvpTeRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetRsvpTePacketRate() uint64 {
	if m != nil {
		return m.RsvpTePacketRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetNonRsvpTeRate() uint64 {
	if m != nil {
		return m.NonRsvpTeRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleRsvp) GetNonRsvpTePacketRate() uint64 {
	if m != nil {
		return m.NonRsvpTePacketRate
	}
	return 0
}

type MplsTeBandwidthAccountHistoryRsvp struct {
	RsvpTeActiveIntervalSample   []*MplsTeBandwidthAccountSampleRsvp `protobuf:"bytes,1,rep,name=rsvp_te_active_interval_sample,json=rsvpTeActiveIntervalSample,proto3" json:"rsvp_te_active_interval_sample,omitempty"`
	RsvpTePreviousIntervalSample []*MplsTeBandwidthAccountSampleRsvp `protobuf:"bytes,2,rep,name=rsvp_te_previous_interval_sample,json=rsvpTePreviousIntervalSample,proto3" json:"rsvp_te_previous_interval_sample,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                            `json:"-"`
	XXX_unrecognized             []byte                              `json:"-"`
	XXX_sizecache                int32                               `json:"-"`
}

func (m *MplsTeBandwidthAccountHistoryRsvp) Reset()         { *m = MplsTeBandwidthAccountHistoryRsvp{} }
func (m *MplsTeBandwidthAccountHistoryRsvp) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountHistoryRsvp) ProtoMessage()    {}
func (*MplsTeBandwidthAccountHistoryRsvp) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{5}
}

func (m *MplsTeBandwidthAccountHistoryRsvp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountHistoryRsvp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountHistoryRsvp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp.Merge(m, src)
}
func (m *MplsTeBandwidthAccountHistoryRsvp) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp.Size(m)
}
func (m *MplsTeBandwidthAccountHistoryRsvp) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountHistoryRsvp proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountHistoryRsvp) GetRsvpTeActiveIntervalSample() []*MplsTeBandwidthAccountSampleRsvp {
	if m != nil {
		return m.RsvpTeActiveIntervalSample
	}
	return nil
}

func (m *MplsTeBandwidthAccountHistoryRsvp) GetRsvpTePreviousIntervalSample() []*MplsTeBandwidthAccountSampleRsvp {
	if m != nil {
		return m.RsvpTePreviousIntervalSample
	}
	return nil
}

type MplsTeBandwidthAccountSampleSr struct {
	TimestampNanosec     uint64   `protobuf:"varint,1,opt,name=timestamp_nanosec,json=timestampNanosec,proto3" json:"timestamp_nanosec,omitempty"`
	SrRate               uint64   `protobuf:"varint,2,opt,name=sr_rate,json=srRate,proto3" json:"sr_rate,omitempty"`
	SrPacketRate         uint64   `protobuf:"varint,3,opt,name=sr_packet_rate,json=srPacketRate,proto3" json:"sr_packet_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeBandwidthAccountSampleSr) Reset()         { *m = MplsTeBandwidthAccountSampleSr{} }
func (m *MplsTeBandwidthAccountSampleSr) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountSampleSr) ProtoMessage()    {}
func (*MplsTeBandwidthAccountSampleSr) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{6}
}

func (m *MplsTeBandwidthAccountSampleSr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleSr.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountSampleSr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleSr.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountSampleSr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountSampleSr.Merge(m, src)
}
func (m *MplsTeBandwidthAccountSampleSr) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountSampleSr.Size(m)
}
func (m *MplsTeBandwidthAccountSampleSr) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountSampleSr.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountSampleSr proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountSampleSr) GetTimestampNanosec() uint64 {
	if m != nil {
		return m.TimestampNanosec
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleSr) GetSrRate() uint64 {
	if m != nil {
		return m.SrRate
	}
	return 0
}

func (m *MplsTeBandwidthAccountSampleSr) GetSrPacketRate() uint64 {
	if m != nil {
		return m.SrPacketRate
	}
	return 0
}

type MplsTeBandwidthAccountHistorySr struct {
	SrActiveIntervalSample   []*MplsTeBandwidthAccountSampleSr `protobuf:"bytes,1,rep,name=sr_active_interval_sample,json=srActiveIntervalSample,proto3" json:"sr_active_interval_sample,omitempty"`
	SrPreviousIntervalSample []*MplsTeBandwidthAccountSampleSr `protobuf:"bytes,2,rep,name=sr_previous_interval_sample,json=srPreviousIntervalSample,proto3" json:"sr_previous_interval_sample,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                          `json:"-"`
	XXX_unrecognized         []byte                            `json:"-"`
	XXX_sizecache            int32                             `json:"-"`
}

func (m *MplsTeBandwidthAccountHistorySr) Reset()         { *m = MplsTeBandwidthAccountHistorySr{} }
func (m *MplsTeBandwidthAccountHistorySr) String() string { return proto.CompactTextString(m) }
func (*MplsTeBandwidthAccountHistorySr) ProtoMessage()    {}
func (*MplsTeBandwidthAccountHistorySr) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{7}
}

func (m *MplsTeBandwidthAccountHistorySr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeBandwidthAccountHistorySr.Unmarshal(m, b)
}
func (m *MplsTeBandwidthAccountHistorySr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeBandwidthAccountHistorySr.Marshal(b, m, deterministic)
}
func (m *MplsTeBandwidthAccountHistorySr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeBandwidthAccountHistorySr.Merge(m, src)
}
func (m *MplsTeBandwidthAccountHistorySr) XXX_Size() int {
	return xxx_messageInfo_MplsTeBandwidthAccountHistorySr.Size(m)
}
func (m *MplsTeBandwidthAccountHistorySr) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeBandwidthAccountHistorySr.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeBandwidthAccountHistorySr proto.InternalMessageInfo

func (m *MplsTeBandwidthAccountHistorySr) GetSrActiveIntervalSample() []*MplsTeBandwidthAccountSampleSr {
	if m != nil {
		return m.SrActiveIntervalSample
	}
	return nil
}

func (m *MplsTeBandwidthAccountHistorySr) GetSrPreviousIntervalSample() []*MplsTeBandwidthAccountSampleSr {
	if m != nil {
		return m.SrPreviousIntervalSample
	}
	return nil
}

type MplsLmBandwidthAccountLinkTelDetInfo struct {
	LinkId               string                                   `protobuf:"bytes,50,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	CommonInfo           *MplsLmBandwidthAccountLinkTelemetryInfo `protobuf:"bytes,51,opt,name=common_info,json=commonInfo,proto3" json:"common_info,omitempty"`
	RsvpTeSampleHistory  *MplsTeBandwidthAccountHistoryRsvp       `protobuf:"bytes,52,opt,name=rsvp_te_sample_history,json=rsvpTeSampleHistory,proto3" json:"rsvp_te_sample_history,omitempty"`
	SrSampleHistory      *MplsTeBandwidthAccountHistorySr         `protobuf:"bytes,53,opt,name=sr_sample_history,json=srSampleHistory,proto3" json:"sr_sample_history,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo) Reset()         { *m = MplsLmBandwidthAccountLinkTelDetInfo{} }
func (m *MplsLmBandwidthAccountLinkTelDetInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBandwidthAccountLinkTelDetInfo) ProtoMessage()    {}
func (*MplsLmBandwidthAccountLinkTelDetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_da5850e9ff76b703, []int{8}
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo.Unmarshal(m, b)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo.Merge(m, src)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo.Size(m)
}
func (m *MplsLmBandwidthAccountLinkTelDetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBandwidthAccountLinkTelDetInfo proto.InternalMessageInfo

func (m *MplsLmBandwidthAccountLinkTelDetInfo) GetLinkId() string {
	if m != nil {
		return m.LinkId
	}
	return ""
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo) GetCommonInfo() *MplsLmBandwidthAccountLinkTelemetryInfo {
	if m != nil {
		return m.CommonInfo
	}
	return nil
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo) GetRsvpTeSampleHistory() *MplsTeBandwidthAccountHistoryRsvp {
	if m != nil {
		return m.RsvpTeSampleHistory
	}
	return nil
}

func (m *MplsLmBandwidthAccountLinkTelDetInfo) GetSrSampleHistory() *MplsTeBandwidthAccountHistorySr {
	if m != nil {
		return m.SrSampleHistory
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsLmBandwidthAccountLinkTelDetInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_lm_bandwidth_account_link_tel_det_info_KEYS")
	proto.RegisterType((*MplsTeBandwidthAccountBwUtilRsvp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_bw_util_rsvp")
	proto.RegisterType((*MplsTeBandwidthAccountBwUtilSr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_bw_util_sr")
	proto.RegisterType((*MplsLmBandwidthAccountLinkTelemetryInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_lm_bandwidth_account_link_telemetry_info")
	proto.RegisterType((*MplsTeBandwidthAccountSampleRsvp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_sample_rsvp")
	proto.RegisterType((*MplsTeBandwidthAccountHistoryRsvp)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_history_rsvp")
	proto.RegisterType((*MplsTeBandwidthAccountSampleSr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_sample_sr")
	proto.RegisterType((*MplsTeBandwidthAccountHistorySr)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_te_bandwidth_account_history_sr")
	proto.RegisterType((*MplsLmBandwidthAccountLinkTelDetInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac_standby.bandwidth_account.bandwidth_account_links.bandwidth_account_link.mpls_lm_bandwidth_account_link_tel_det_info")
}

func init() {
	proto.RegisterFile("mpls_lm_bandwidth_account_link_tel_det_info.proto", fileDescriptor_da5850e9ff76b703)
}

var fileDescriptor_da5850e9ff76b703 = []byte{
	// 1022 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x05, 0xa3, 0xc4, 0xb1, 0xae, 0x1d, 0x3f, 0xe8, 0x22, 0x51, 0x9a, 0xd8, 0x55, 0x14, 0xc7,
	0x36, 0x1a, 0x44, 0x68, 0x9c, 0x14, 0xe8, 0xae, 0x50, 0x03, 0xa3, 0x31, 0xdc, 0x06, 0x05, 0xad,
	0x16, 0xc8, 0x6a, 0x30, 0x22, 0x47, 0xcd, 0xd4, 0xe4, 0x0c, 0x31, 0x33, 0x52, 0xac, 0xfe, 0x45,
	0xfa, 0x58, 0x75, 0x9f, 0xbf, 0x28, 0xba, 0x0f, 0x50, 0x14, 0xdd, 0x65, 0xd1, 0x55, 0x3f, 0xa2,
	0xeb, 0x82, 0x77, 0x48, 0x8a, 0x7a, 0x51, 0xf1, 0xa6, 0xd5, 0x52, 0xc3, 0xc3, 0xc3, 0x7b, 0xee,
	0x3d, 0x73, 0x66, 0x6c, 0x78, 0x18, 0xc5, 0xa1, 0x26, 0x61, 0x44, 0x3a, 0x54, 0x04, 0x2f, 0x79,
	0x60, 0x5e, 0x10, 0xea, 0xfb, 0xb2, 0x27, 0x0c, 0x09, 0xb9, 0x38, 0x23, 0x86, 0x85, 0x24, 0x60,
	0x86, 0x70, 0xd1, 0x95, 0xcd, 0x58, 0x49, 0x23, 0xdd, 0xd0, 0xe7, 0xda, 0x97, 0x84, 0x4b, 0x4d,
	0xce, 0x15, 0xc1, 0xf7, 0x0d, 0x23, 0x32, 0x66, 0xaa, 0x69, 0xc9, 0x7c, 0xea, 0x13, 0x6d, 0xa8,
	0x08, 0x3a, 0x83, 0xe6, 0x04, 0x6d, 0x73, 0xfa, 0x87, 0xf4, 0x8c, 0xf5, 0xc6, 0x73, 0xf8, 0xe8,
	0x02, 0x25, 0x92, 0x93, 0xa3, 0xe7, 0xa7, 0xee, 0x3d, 0x58, 0xe3, 0xc2, 0x30, 0xd5, 0xa5, 0x3e,
	0x23, 0x82, 0x46, 0xac, 0xe6, 0xd4, 0x9d, 0x83, 0xaa, 0x77, 0x2d, 0x5f, 0x7d, 0x46, 0x23, 0xd6,
	0x78, 0x5d, 0x81, 0xbd, 0xac, 0xfc, 0x49, 0xee, 0xce, 0x4b, 0xd2, 0x33, 0x3c, 0x24, 0x4a, 0xf7,
	0x63, 0xf7, 0x73, 0xa8, 0x1b, 0x69, 0x68, 0x68, 0xbf, 0x38, 0x04, 0x27, 0x08, 0xfe, 0x3d, 0x35,
	0x5c, 0x0a, 0xfc, 0xc6, 0x65, 0x6f, 0x1b, 0x71, 0x5f, 0x70, 0x71, 0xf6, 0x59, 0x86, 0xfa, 0x7a,
	0x08, 0x72, 0x5b, 0xb0, 0x9d, 0x10, 0x8e, 0x7e, 0xb2, 0xc8, 0x72, 0x09, 0x59, 0xde, 0x4f, 0x40,
	0x6d, 0x36, 0x95, 0xe2, 0x29, 0xdc, 0x11, 0x52, 0x90, 0x72, 0x9a, 0x8a, 0x2d, 0x46, 0x48, 0xe1,
	0xcd, 0x66, 0xf2, 0x60, 0x2f, 0x63, 0xa1, 0xc1, 0x77, 0x3d, 0x6d, 0x58, 0x30, 0x83, 0xee, 0x32,
	0xd2, 0x35, 0x6c, 0x55, 0xad, 0x14, 0x3b, 0x8f, 0x93, 0x89, 0xae, 0x54, 0xfe, 0x4c, 0xce, 0x2b,
	0x45, 0xce, 0xa3, 0x14, 0x3b, 0x8d, 0xb3, 0xf1, 0x8f, 0x03, 0xbb, 0xf3, 0x07, 0xa5, 0x95, 0xfb,
	0x09, 0xd4, 0xb4, 0x2a, 0x1d, 0xcf, 0x75, 0xad, 0xa6, 0x96, 0x7d, 0x0c, 0x77, 0xb4, 0x9a, 0xd7,
	0x05, 0x3b, 0x9b, 0x1d, 0xad, 0x4a, 0x3b, 0x60, 0xa9, 0xe6, 0x88, 0xaf, 0x64, 0x54, 0xa5, 0xc2,
	0x5f, 0x57, 0xe1, 0xc1, 0x7c, 0xf7, 0xb3, 0x88, 0x19, 0x35, 0x40, 0xff, 0xbb, 0x9f, 0xc2, 0x6d,
	0xae, 0xa7, 0x60, 0x99, 0xa0, 0x9d, 0x90, 0x05, 0xd8, 0x85, 0x65, 0xef, 0x26, 0xd7, 0xf9, 0xf7,
	0x5a, 0x16, 0x71, 0x64, 0x01, 0xee, 0x43, 0x78, 0x8f, 0xc6, 0x71, 0xc8, 0x7d, 0xac, 0x20, 0x97,
	0x81, 0xda, 0x97, 0xbd, 0xad, 0xc2, 0xb3, 0xac, 0x72, 0x77, 0x1f, 0xd6, 0x7d, 0x19, 0x86, 0xcc,
	0xc7, 0x37, 0xcc, 0x20, 0x66, 0x28, 0xaf, 0xea, 0xad, 0x0d, 0x97, 0xdb, 0x83, 0x98, 0xb9, 0x6d,
	0xd8, 0xd7, 0x34, 0x8a, 0x43, 0x46, 0x0c, 0x8f, 0x18, 0x51, 0x2c, 0xa2, 0x5c, 0x70, 0xf1, 0x2d,
	0xfe, 0xd4, 0x86, 0x46, 0x31, 0x11, 0x54, 0x48, 0xcd, 0xfc, 0xd4, 0x70, 0x77, 0x2d, 0xbc, 0xcd,
	0x23, 0xe6, 0x65, 0xe0, 0x76, 0x86, 0x7d, 0x66, 0xa1, 0xee, 0x37, 0x70, 0x10, 0x52, 0x6d, 0x48,
	0x4a, 0x5d, 0x2c, 0x65, 0x82, 0xd6, 0x7a, 0x6e, 0x37, 0xc1, 0x9f, 0x22, 0xfc, 0xc9, 0xb0, 0xc2,
	0x71, 0xde, 0x27, 0xb0, 0x23, 0xd8, 0xf9, 0x34, 0xde, 0x8c, 0x6d, 0x09, 0xd9, 0x6e, 0x25, 0xa8,
	0x71, 0xb6, 0x8c, 0xe4, 0x04, 0x1a, 0xc5, 0x76, 0x8e, 0xe9, 0xce, 0x88, 0xae, 0x22, 0xd1, 0x07,
	0x05, 0xe4, 0x88, 0xe4, 0x02, 0x19, 0x2a, 0x1d, 0x67, 0x1c, 0xd5, 0xb8, 0x6c, 0xc9, 0x12, 0x64,
	0x6b, 0x94, 0x70, 0x44, 0xde, 0x5b, 0x67, 0x5e, 0x14, 0x55, 0xeb, 0xce, 0xc1, 0xca, 0xe1, 0x4f,
	0x4e, 0xf3, 0xbf, 0x0c, 0xfc, 0xe6, 0xbb, 0x25, 0x72, 0x69, 0x42, 0xfe, 0xee, 0x94, 0xe4, 0x00,
	0xa0, 0xaa, 0x1f, 0x16, 0x4e, 0x95, 0x56, 0x33, 0xc3, 0xe9, 0x04, 0x1a, 0xe8, 0xc4, 0xf2, 0xb9,
	0xaf, 0xd8, 0xb9, 0x27, 0xc8, 0xb2, 0xb9, 0x9f, 0xc2, 0x1e, 0xeb, 0x76, 0x13, 0x97, 0xf6, 0x19,
	0x89, 0xe8, 0x39, 0x8f, 0x7a, 0x11, 0x51, 0x4c, 0x33, 0xd5, 0x4f, 0x12, 0x60, 0x58, 0x61, 0x6d,
	0xd5, 0xee, 0xc1, 0x1c, 0xfd, 0xa5, 0x05, 0x7b, 0x39, 0x36, 0x2f, 0xb6, 0xf1, 0xc7, 0x25, 0xb8,
	0x37, 0x5b, 0x62, 0xba, 0x83, 0xf0, 0x24, 0xbd, 0x0f, 0x9b, 0x93, 0xa5, 0xdb, 0x6c, 0xde, 0x30,
	0xe3, 0xb5, 0x6e, 0x03, 0xd8, 0x63, 0x57, 0x51, 0xc3, 0xd2, 0xf8, 0xad, 0xe2, 0x8a, 0x47, 0x0d,
	0x73, 0x3f, 0x84, 0x4d, 0xfb, 0x38, 0xa6, 0xfe, 0x19, 0x33, 0x16, 0x65, 0x93, 0x75, 0x1d, 0x1f,
	0x7c, 0x85, 0xeb, 0x88, 0xad, 0xc3, 0x6a, 0xe6, 0x76, 0x84, 0xd9, 0x80, 0x01, 0xeb, 0x22, 0x44,
	0x3c, 0x80, 0xad, 0x0c, 0x51, 0xe4, 0xb3, 0x91, 0xb1, 0x61, 0x81, 0x05, 0xc2, 0x7d, 0xd8, 0x28,
	0x1e, 0xc3, 0x88, 0xb5, 0x81, 0x70, 0x2d, 0x3f, 0x75, 0x11, 0xf8, 0x18, 0x6e, 0x14, 0x81, 0x45,
	0x6e, 0xbb, 0xef, 0xb7, 0x72, 0xfc, 0x90, 0xbe, 0xf1, 0x77, 0xe9, 0xe5, 0xe4, 0x05, 0xd7, 0x46,
	0xaa, 0x81, 0x6d, 0xe9, 0x5b, 0x07, 0x76, 0xf2, 0x73, 0xdc, 0xce, 0x15, 0x2f, 0x3a, 0x7d, 0x1a,
	0xa6, 0xad, 0xaf, 0x39, 0xf5, 0xca, 0xc1, 0xca, 0xe1, 0x8f, 0x0b, 0x63, 0xfa, 0x82, 0x23, 0xb2,
	0x9d, 0xdc, 0xc2, 0xca, 0x8f, 0xd3, 0xc2, 0x6d, 0x9e, 0xba, 0x7f, 0x39, 0x50, 0xcf, 0x1b, 0xa7,
	0x58, 0x9f, 0xcb, 0x9e, 0x9e, 0x10, 0x77, 0x69, 0x81, 0xc5, 0xdd, 0x4e, 0x7d, 0x93, 0xd6, 0x3e,
	0x2a, 0xaf, 0xf1, 0xca, 0x81, 0xbb, 0x73, 0x79, 0xb4, 0xba, 0xd8, 0xa6, 0xb9, 0x01, 0x57, 0xb5,
	0x2a, 0xee, 0x98, 0x25, 0xad, 0xd0, 0x88, 0xbb, 0xb0, 0xa6, 0xd5, 0x94, 0xbd, 0xb2, 0xaa, 0x55,
	0xc1, 0x78, 0x6f, 0x2a, 0x65, 0x97, 0xad, 0xcc, 0x78, 0x5a, 0x25, 0x29, 0x7b, 0x33, 0xb9, 0x33,
	0x95, 0x39, 0xee, 0xd5, 0xa2, 0x0d, 0xc5, 0xa6, 0xec, 0x54, 0xaf, 0xfd, 0xe9, 0xc0, 0xad, 0xa4,
	0x3f, 0xe5, 0x36, 0x5b, 0x44, 0x45, 0x35, 0xad, 0x66, 0x18, 0xec, 0xe7, 0x2b, 0x70, 0xff, 0x02,
	0x7f, 0x3e, 0x25, 0xde, 0xc1, 0x45, 0x1e, 0xd4, 0x0e, 0xf1, 0x0a, 0xb7, 0x94, 0xfc, 0x3c, 0x0e,
	0xdc, 0x5f, 0x1d, 0x58, 0xf1, 0x65, 0x14, 0x49, 0x81, 0xc0, 0xda, 0x23, 0x3c, 0x45, 0x7f, 0xf9,
	0x3f, 0x9a, 0xf1, 0xae, 0x77, 0x61, 0x0f, 0x6c, 0xc1, 0xc7, 0x89, 0xb0, 0x37, 0x0e, 0x5c, 0xcf,
	0x82, 0x24, 0x6d, 0x5c, 0x6a, 0xe5, 0xda, 0xe3, 0x05, 0xbb, 0xe6, 0x14, 0xb3, 0xdd, 0xdb, 0xb2,
	0xf9, 0x61, 0xc7, 0xf9, 0xd4, 0x3e, 0x71, 0x7f, 0x73, 0x60, 0x53, 0xab, 0x71, 0x1d, 0x1f, 0x2f,
	0xd8, 0xc5, 0x66, 0x18, 0x15, 0xde, 0xba, 0x56, 0x23, 0x0a, 0x3a, 0x4b, 0xf8, 0xaf, 0x84, 0x47,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x62, 0xea, 0x96, 0x8f, 0x7f, 0x10, 0x00, 0x00,
}