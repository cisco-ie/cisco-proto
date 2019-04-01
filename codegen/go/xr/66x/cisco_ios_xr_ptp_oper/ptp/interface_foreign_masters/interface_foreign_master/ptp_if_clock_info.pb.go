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
// source: ptp_if_clock_info.proto

package cisco_ios_xr_ptp_oper_ptp_interface_foreign_masters_interface_foreign_master

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

type PtpIfClockInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfClockInfo_KEYS) Reset()         { *m = PtpIfClockInfo_KEYS{} }
func (m *PtpIfClockInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfClockInfo_KEYS) ProtoMessage()    {}
func (*PtpIfClockInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{0}
}

func (m *PtpIfClockInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfClockInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpIfClockInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfClockInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfClockInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfClockInfo_KEYS.Merge(m, src)
}
func (m *PtpIfClockInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfClockInfo_KEYS.Size(m)
}
func (m *PtpIfClockInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfClockInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfClockInfo_KEYS proto.InternalMessageInfo

func (m *PtpIfClockInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PtpBagClockUtcOffset struct {
	CurrentOffset        string   `protobuf:"bytes,1,opt,name=current_offset,json=currentOffset,proto3" json:"current_offset,omitempty"`
	OffsetValid          bool     `protobuf:"varint,2,opt,name=offset_valid,json=offsetValid,proto3" json:"offset_valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagClockUtcOffset) Reset()         { *m = PtpBagClockUtcOffset{} }
func (m *PtpBagClockUtcOffset) String() string { return proto.CompactTextString(m) }
func (*PtpBagClockUtcOffset) ProtoMessage()    {}
func (*PtpBagClockUtcOffset) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{1}
}

func (m *PtpBagClockUtcOffset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagClockUtcOffset.Unmarshal(m, b)
}
func (m *PtpBagClockUtcOffset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagClockUtcOffset.Marshal(b, m, deterministic)
}
func (m *PtpBagClockUtcOffset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagClockUtcOffset.Merge(m, src)
}
func (m *PtpBagClockUtcOffset) XXX_Size() int {
	return xxx_messageInfo_PtpBagClockUtcOffset.Size(m)
}
func (m *PtpBagClockUtcOffset) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagClockUtcOffset.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagClockUtcOffset proto.InternalMessageInfo

func (m *PtpBagClockUtcOffset) GetCurrentOffset() string {
	if m != nil {
		return m.CurrentOffset
	}
	return ""
}

func (m *PtpBagClockUtcOffset) GetOffsetValid() bool {
	if m != nil {
		return m.OffsetValid
	}
	return false
}

type PtpBagPortId struct {
	ClockId              uint64   `protobuf:"varint,1,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	PortNumber           uint32   `protobuf:"varint,2,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagPortId) Reset()         { *m = PtpBagPortId{} }
func (m *PtpBagPortId) String() string { return proto.CompactTextString(m) }
func (*PtpBagPortId) ProtoMessage()    {}
func (*PtpBagPortId) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{2}
}

func (m *PtpBagPortId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagPortId.Unmarshal(m, b)
}
func (m *PtpBagPortId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagPortId.Marshal(b, m, deterministic)
}
func (m *PtpBagPortId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagPortId.Merge(m, src)
}
func (m *PtpBagPortId) XXX_Size() int {
	return xxx_messageInfo_PtpBagPortId.Size(m)
}
func (m *PtpBagPortId) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagPortId.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagPortId proto.InternalMessageInfo

func (m *PtpBagPortId) GetClockId() uint64 {
	if m != nil {
		return m.ClockId
	}
	return 0
}

func (m *PtpBagPortId) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

type PtpBagClock struct {
	ClockId              uint64                `protobuf:"varint,1,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	Priority1            uint32                `protobuf:"varint,2,opt,name=priority1,proto3" json:"priority1,omitempty"`
	Priority2            uint32                `protobuf:"varint,3,opt,name=priority2,proto3" json:"priority2,omitempty"`
	Class                uint32                `protobuf:"varint,4,opt,name=class,proto3" json:"class,omitempty"`
	Accuracy             uint32                `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	OffsetLogVariance    uint32                `protobuf:"varint,6,opt,name=offset_log_variance,json=offsetLogVariance,proto3" json:"offset_log_variance,omitempty"`
	StepsRemoved         uint32                `protobuf:"varint,7,opt,name=steps_removed,json=stepsRemoved,proto3" json:"steps_removed,omitempty"`
	TimeSource           string                `protobuf:"bytes,8,opt,name=time_source,json=timeSource,proto3" json:"time_source,omitempty"`
	UtcOffset            *PtpBagClockUtcOffset `protobuf:"bytes,9,opt,name=utc_offset,json=utcOffset,proto3" json:"utc_offset,omitempty"`
	FrequencyTraceable   bool                  `protobuf:"varint,10,opt,name=frequency_traceable,json=frequencyTraceable,proto3" json:"frequency_traceable,omitempty"`
	TimeTraceable        bool                  `protobuf:"varint,11,opt,name=time_traceable,json=timeTraceable,proto3" json:"time_traceable,omitempty"`
	Timescale            string                `protobuf:"bytes,12,opt,name=timescale,proto3" json:"timescale,omitempty"`
	LeapSeconds          string                `protobuf:"bytes,13,opt,name=leap_seconds,json=leapSeconds,proto3" json:"leap_seconds,omitempty"`
	Receiver             *PtpBagPortId         `protobuf:"bytes,14,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Sender               *PtpBagPortId         `protobuf:"bytes,15,opt,name=sender,proto3" json:"sender,omitempty"`
	Local                bool                  `protobuf:"varint,16,opt,name=local,proto3" json:"local,omitempty"`
	ConfiguredClockClass uint32                `protobuf:"varint,17,opt,name=configured_clock_class,json=configuredClockClass,proto3" json:"configured_clock_class,omitempty"`
	ConfiguredPriority   uint32                `protobuf:"varint,18,opt,name=configured_priority,json=configuredPriority,proto3" json:"configured_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpBagClock) Reset()         { *m = PtpBagClock{} }
func (m *PtpBagClock) String() string { return proto.CompactTextString(m) }
func (*PtpBagClock) ProtoMessage()    {}
func (*PtpBagClock) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{3}
}

func (m *PtpBagClock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagClock.Unmarshal(m, b)
}
func (m *PtpBagClock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagClock.Marshal(b, m, deterministic)
}
func (m *PtpBagClock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagClock.Merge(m, src)
}
func (m *PtpBagClock) XXX_Size() int {
	return xxx_messageInfo_PtpBagClock.Size(m)
}
func (m *PtpBagClock) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagClock.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagClock proto.InternalMessageInfo

func (m *PtpBagClock) GetClockId() uint64 {
	if m != nil {
		return m.ClockId
	}
	return 0
}

func (m *PtpBagClock) GetPriority1() uint32 {
	if m != nil {
		return m.Priority1
	}
	return 0
}

func (m *PtpBagClock) GetPriority2() uint32 {
	if m != nil {
		return m.Priority2
	}
	return 0
}

func (m *PtpBagClock) GetClass() uint32 {
	if m != nil {
		return m.Class
	}
	return 0
}

func (m *PtpBagClock) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *PtpBagClock) GetOffsetLogVariance() uint32 {
	if m != nil {
		return m.OffsetLogVariance
	}
	return 0
}

func (m *PtpBagClock) GetStepsRemoved() uint32 {
	if m != nil {
		return m.StepsRemoved
	}
	return 0
}

func (m *PtpBagClock) GetTimeSource() string {
	if m != nil {
		return m.TimeSource
	}
	return ""
}

func (m *PtpBagClock) GetUtcOffset() *PtpBagClockUtcOffset {
	if m != nil {
		return m.UtcOffset
	}
	return nil
}

func (m *PtpBagClock) GetFrequencyTraceable() bool {
	if m != nil {
		return m.FrequencyTraceable
	}
	return false
}

func (m *PtpBagClock) GetTimeTraceable() bool {
	if m != nil {
		return m.TimeTraceable
	}
	return false
}

func (m *PtpBagClock) GetTimescale() string {
	if m != nil {
		return m.Timescale
	}
	return ""
}

func (m *PtpBagClock) GetLeapSeconds() string {
	if m != nil {
		return m.LeapSeconds
	}
	return ""
}

func (m *PtpBagClock) GetReceiver() *PtpBagPortId {
	if m != nil {
		return m.Receiver
	}
	return nil
}

func (m *PtpBagClock) GetSender() *PtpBagPortId {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *PtpBagClock) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *PtpBagClock) GetConfiguredClockClass() uint32 {
	if m != nil {
		return m.ConfiguredClockClass
	}
	return 0
}

func (m *PtpBagClock) GetConfiguredPriority() uint32 {
	if m != nil {
		return m.ConfiguredPriority
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
	return fileDescriptor_51135ae6cdc5ffb0, []int{4}
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
	return fileDescriptor_51135ae6cdc5ffb0, []int{5}
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
	return fileDescriptor_51135ae6cdc5ffb0, []int{6}
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

type PtpBagUnicastGrant struct {
	LogGrantInterval     string   `protobuf:"bytes,1,opt,name=log_grant_interval,json=logGrantInterval,proto3" json:"log_grant_interval,omitempty"`
	GrantDuration        uint32   `protobuf:"varint,2,opt,name=grant_duration,json=grantDuration,proto3" json:"grant_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagUnicastGrant) Reset()         { *m = PtpBagUnicastGrant{} }
func (m *PtpBagUnicastGrant) String() string { return proto.CompactTextString(m) }
func (*PtpBagUnicastGrant) ProtoMessage()    {}
func (*PtpBagUnicastGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{7}
}

func (m *PtpBagUnicastGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagUnicastGrant.Unmarshal(m, b)
}
func (m *PtpBagUnicastGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagUnicastGrant.Marshal(b, m, deterministic)
}
func (m *PtpBagUnicastGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagUnicastGrant.Merge(m, src)
}
func (m *PtpBagUnicastGrant) XXX_Size() int {
	return xxx_messageInfo_PtpBagUnicastGrant.Size(m)
}
func (m *PtpBagUnicastGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagUnicastGrant.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagUnicastGrant proto.InternalMessageInfo

func (m *PtpBagUnicastGrant) GetLogGrantInterval() string {
	if m != nil {
		return m.LogGrantInterval
	}
	return ""
}

func (m *PtpBagUnicastGrant) GetGrantDuration() uint32 {
	if m != nil {
		return m.GrantDuration
	}
	return 0
}

type PtpBagForeignClock struct {
	ForeignClock         *PtpBagClock        `protobuf:"bytes,1,opt,name=foreign_clock,json=foreignClock,proto3" json:"foreign_clock,omitempty"`
	IsQualified          bool                `protobuf:"varint,2,opt,name=is_qualified,json=isQualified,proto3" json:"is_qualified,omitempty"`
	IsGrandmaster        bool                `protobuf:"varint,3,opt,name=is_grandmaster,json=isGrandmaster,proto3" json:"is_grandmaster,omitempty"`
	CommunicationModel   string              `protobuf:"bytes,4,opt,name=communication_model,json=communicationModel,proto3" json:"communication_model,omitempty"`
	IsKnown              bool                `protobuf:"varint,5,opt,name=is_known,json=isKnown,proto3" json:"is_known,omitempty"`
	TimeKnownFor         uint32              `protobuf:"varint,6,opt,name=time_known_for,json=timeKnownFor,proto3" json:"time_known_for,omitempty"`
	ForeignDomain        uint32              `protobuf:"varint,7,opt,name=foreign_domain,json=foreignDomain,proto3" json:"foreign_domain,omitempty"`
	Address              *PtpBagAddress      `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	ConfiguredPriority   uint32              `protobuf:"varint,9,opt,name=configured_priority,json=configuredPriority,proto3" json:"configured_priority,omitempty"`
	ConfiguredClockClass uint32              `protobuf:"varint,10,opt,name=configured_clock_class,json=configuredClockClass,proto3" json:"configured_clock_class,omitempty"`
	DelayAsymmetry       int32               `protobuf:"zigzag32,11,opt,name=delay_asymmetry,json=delayAsymmetry,proto3" json:"delay_asymmetry,omitempty"`
	AnnounceGrant        *PtpBagUnicastGrant `protobuf:"bytes,12,opt,name=announce_grant,json=announceGrant,proto3" json:"announce_grant,omitempty"`
	SyncGrant            *PtpBagUnicastGrant `protobuf:"bytes,13,opt,name=sync_grant,json=syncGrant,proto3" json:"sync_grant,omitempty"`
	DelayResponseGrant   *PtpBagUnicastGrant `protobuf:"bytes,14,opt,name=delay_response_grant,json=delayResponseGrant,proto3" json:"delay_response_grant,omitempty"`
	PtsfLossAnnounce     bool                `protobuf:"varint,15,opt,name=ptsf_loss_announce,json=ptsfLossAnnounce,proto3" json:"ptsf_loss_announce,omitempty"`
	PtsfLossSync         bool                `protobuf:"varint,16,opt,name=ptsf_loss_sync,json=ptsfLossSync,proto3" json:"ptsf_loss_sync,omitempty"`
	IsDnu                bool                `protobuf:"varint,17,opt,name=is_dnu,json=isDnu,proto3" json:"is_dnu,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PtpBagForeignClock) Reset()         { *m = PtpBagForeignClock{} }
func (m *PtpBagForeignClock) String() string { return proto.CompactTextString(m) }
func (*PtpBagForeignClock) ProtoMessage()    {}
func (*PtpBagForeignClock) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{8}
}

func (m *PtpBagForeignClock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagForeignClock.Unmarshal(m, b)
}
func (m *PtpBagForeignClock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagForeignClock.Marshal(b, m, deterministic)
}
func (m *PtpBagForeignClock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagForeignClock.Merge(m, src)
}
func (m *PtpBagForeignClock) XXX_Size() int {
	return xxx_messageInfo_PtpBagForeignClock.Size(m)
}
func (m *PtpBagForeignClock) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagForeignClock.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagForeignClock proto.InternalMessageInfo

func (m *PtpBagForeignClock) GetForeignClock() *PtpBagClock {
	if m != nil {
		return m.ForeignClock
	}
	return nil
}

func (m *PtpBagForeignClock) GetIsQualified() bool {
	if m != nil {
		return m.IsQualified
	}
	return false
}

func (m *PtpBagForeignClock) GetIsGrandmaster() bool {
	if m != nil {
		return m.IsGrandmaster
	}
	return false
}

func (m *PtpBagForeignClock) GetCommunicationModel() string {
	if m != nil {
		return m.CommunicationModel
	}
	return ""
}

func (m *PtpBagForeignClock) GetIsKnown() bool {
	if m != nil {
		return m.IsKnown
	}
	return false
}

func (m *PtpBagForeignClock) GetTimeKnownFor() uint32 {
	if m != nil {
		return m.TimeKnownFor
	}
	return 0
}

func (m *PtpBagForeignClock) GetForeignDomain() uint32 {
	if m != nil {
		return m.ForeignDomain
	}
	return 0
}

func (m *PtpBagForeignClock) GetAddress() *PtpBagAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PtpBagForeignClock) GetConfiguredPriority() uint32 {
	if m != nil {
		return m.ConfiguredPriority
	}
	return 0
}

func (m *PtpBagForeignClock) GetConfiguredClockClass() uint32 {
	if m != nil {
		return m.ConfiguredClockClass
	}
	return 0
}

func (m *PtpBagForeignClock) GetDelayAsymmetry() int32 {
	if m != nil {
		return m.DelayAsymmetry
	}
	return 0
}

func (m *PtpBagForeignClock) GetAnnounceGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.AnnounceGrant
	}
	return nil
}

func (m *PtpBagForeignClock) GetSyncGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.SyncGrant
	}
	return nil
}

func (m *PtpBagForeignClock) GetDelayResponseGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.DelayResponseGrant
	}
	return nil
}

func (m *PtpBagForeignClock) GetPtsfLossAnnounce() bool {
	if m != nil {
		return m.PtsfLossAnnounce
	}
	return false
}

func (m *PtpBagForeignClock) GetPtsfLossSync() bool {
	if m != nil {
		return m.PtsfLossSync
	}
	return false
}

func (m *PtpBagForeignClock) GetIsDnu() bool {
	if m != nil {
		return m.IsDnu
	}
	return false
}

type PtpIfClockInfo struct {
	PortNumber           uint32                `protobuf:"varint,50,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	ForeignClock         []*PtpBagForeignClock `protobuf:"bytes,51,rep,name=foreign_clock,json=foreignClock,proto3" json:"foreign_clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpIfClockInfo) Reset()         { *m = PtpIfClockInfo{} }
func (m *PtpIfClockInfo) String() string { return proto.CompactTextString(m) }
func (*PtpIfClockInfo) ProtoMessage()    {}
func (*PtpIfClockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_51135ae6cdc5ffb0, []int{9}
}

func (m *PtpIfClockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfClockInfo.Unmarshal(m, b)
}
func (m *PtpIfClockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfClockInfo.Marshal(b, m, deterministic)
}
func (m *PtpIfClockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfClockInfo.Merge(m, src)
}
func (m *PtpIfClockInfo) XXX_Size() int {
	return xxx_messageInfo_PtpIfClockInfo.Size(m)
}
func (m *PtpIfClockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfClockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfClockInfo proto.InternalMessageInfo

func (m *PtpIfClockInfo) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

func (m *PtpIfClockInfo) GetForeignClock() []*PtpBagForeignClock {
	if m != nil {
		return m.ForeignClock
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpIfClockInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_if_clock_info_KEYS")
	proto.RegisterType((*PtpBagClockUtcOffset)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_clock_utc_offset")
	proto.RegisterType((*PtpBagPortId)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_port_id")
	proto.RegisterType((*PtpBagClock)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_clock")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_address")
	proto.RegisterType((*PtpBagUnicastGrant)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_unicast_grant")
	proto.RegisterType((*PtpBagForeignClock)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_bag_foreign_clock")
	proto.RegisterType((*PtpIfClockInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_foreign_masters.interface_foreign_master.ptp_if_clock_info")
}

func init() { proto.RegisterFile("ptp_if_clock_info.proto", fileDescriptor_51135ae6cdc5ffb0) }

var fileDescriptor_51135ae6cdc5ffb0 = []byte{
	// 1077 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0xd6, 0x91, 0x26, 0xb6, 0xc7, 0x3e, 0xa7, 0xd9, 0xa4, 0xe5, 0x40, 0x95, 0x48, 0x4d, 0x51,
	0xf3, 0x80, 0x5c, 0x35, 0x8d, 0x78, 0xe1, 0x01, 0x45, 0x0d, 0x44, 0x55, 0xd3, 0x02, 0x17, 0xa8,
	0x84, 0x10, 0x5a, 0x6d, 0xf6, 0xd6, 0xd6, 0xaa, 0x77, 0xbb, 0xd7, 0xdd, 0x3b, 0x17, 0xbf, 0x21,
	0x04, 0x02, 0xf1, 0x0c, 0x4f, 0x48, 0xfc, 0x15, 0xfe, 0x1a, 0xda, 0xd9, 0x3d, 0x5f, 0x9c, 0x28,
	0x79, 0xc2, 0x79, 0xdc, 0xef, 0x9b, 0x9d, 0x99, 0x9d, 0x99, 0xfb, 0xc6, 0x86, 0x77, 0xcb, 0xaa,
	0xa4, 0x72, 0x42, 0x79, 0xae, 0xf9, 0x6b, 0x2a, 0xd5, 0x44, 0x8f, 0x4b, 0xa3, 0x2b, 0x4d, 0x4e,
	0xb8, 0xb4, 0x5c, 0x53, 0xa9, 0x2d, 0xfd, 0xd1, 0x50, 0x67, 0xa5, 0x4b, 0x61, 0xc6, 0x65, 0x55,
	0x8e, 0xa5, 0xaa, 0x84, 0x99, 0x30, 0x2e, 0xe8, 0x44, 0x1b, 0x21, 0xa7, 0x8a, 0x16, 0xcc, 0x56,
	0xc2, 0xd8, 0x2b, 0x99, 0xd1, 0x67, 0x70, 0xf7, 0x52, 0x20, 0xfa, 0xfc, 0xf3, 0xef, 0x4e, 0xc9,
	0x47, 0x30, 0x6c, 0x6f, 0x29, 0x56, 0x88, 0x24, 0xda, 0x8d, 0xf6, 0x7a, 0x69, 0xbc, 0x40, 0x5f,
	0xb2, 0x42, 0x8c, 0x32, 0x48, 0x9c, 0x83, 0x33, 0x36, 0x0d, 0x1e, 0xea, 0x8a, 0x53, 0x3d, 0x99,
	0x58, 0x51, 0x39, 0x17, 0xbc, 0x36, 0x46, 0xa8, 0x2a, 0x20, 0x8d, 0x8b, 0x80, 0x7e, 0xe9, 0xcd,
	0xee, 0xc3, 0xc0, 0xd3, 0x74, 0xc6, 0x72, 0x99, 0x25, 0xef, 0xec, 0x46, 0x7b, 0xdd, 0xb4, 0xef,
	0xb1, 0x57, 0x0e, 0x1a, 0xbd, 0x80, 0xcd, 0x26, 0x4a, 0xa9, 0x4d, 0x45, 0x65, 0x46, 0xde, 0x83,
	0x6e, 0x48, 0x39, 0x43, 0xb7, 0xb7, 0xd2, 0x0e, 0x9e, 0x9f, 0x65, 0xe4, 0x03, 0xe8, 0xa3, 0x95,
	0xaa, 0x8b, 0x33, 0x61, 0xd0, 0x5f, 0x9c, 0x82, 0x83, 0x5e, 0x22, 0x32, 0xfa, 0xa7, 0x03, 0xf1,
	0x52, 0xd6, 0xd7, 0x79, 0xbb, 0x07, 0xbd, 0xd2, 0x48, 0x6d, 0x64, 0x35, 0x7f, 0x1c, 0x7c, 0xb5,
	0xc0, 0x79, 0x76, 0x3f, 0x59, 0x5b, 0x66, 0xf7, 0xc9, 0x0e, 0xac, 0xf3, 0x9c, 0x59, 0x9b, 0xdc,
	0x42, 0xc6, 0x1f, 0xc8, 0xfb, 0xd0, 0x65, 0x9c, 0xd7, 0x86, 0xf1, 0x79, 0xb2, 0x8e, 0xc4, 0xe2,
	0x4c, 0xc6, 0xb0, 0x1d, 0x8a, 0x91, 0xeb, 0x29, 0x9d, 0x31, 0x23, 0x99, 0xe2, 0x22, 0xd9, 0x40,
	0xb3, 0x2d, 0x4f, 0x9d, 0xe8, 0xe9, 0xab, 0x40, 0x90, 0x0f, 0x21, 0xb6, 0x95, 0x28, 0x2d, 0x35,
	0xa2, 0xd0, 0x33, 0x91, 0x25, 0x1d, 0xb4, 0x1c, 0x20, 0x98, 0x7a, 0xcc, 0x15, 0xa4, 0x92, 0x85,
	0xa0, 0x56, 0xd7, 0x86, 0x8b, 0xa4, 0x8b, 0x5d, 0x00, 0x07, 0x9d, 0x22, 0x42, 0x7e, 0x8d, 0x00,
	0xda, 0xc6, 0x25, 0xbd, 0xdd, 0x68, 0xaf, 0xbf, 0x3f, 0x19, 0xff, 0x9f, 0xa3, 0x36, 0xbe, 0x6a,
	0x4c, 0xd2, 0x5e, 0x5d, 0xf1, 0x30, 0x0a, 0x8f, 0x60, 0x7b, 0x62, 0xc4, 0x9b, 0x5a, 0x28, 0x3e,
	0xa7, 0x95, 0x61, 0x5c, 0xb0, 0xb3, 0x5c, 0x24, 0x80, 0x13, 0x41, 0x16, 0xd4, 0x37, 0x0d, 0xe3,
	0x46, 0x0c, 0x5f, 0xd6, 0xda, 0xf6, 0xd1, 0x36, 0x76, 0x68, 0x6b, 0x76, 0x0f, 0x7a, 0x0e, 0xb0,
	0x9c, 0xe5, 0x22, 0x19, 0xe0, 0xf3, 0x5b, 0xc0, 0x0d, 0x60, 0x2e, 0x58, 0x49, 0xad, 0xe0, 0x5a,
	0x65, 0x36, 0x89, 0xd1, 0xa0, 0xef, 0xb0, 0x53, 0x0f, 0x91, 0x39, 0x74, 0x8d, 0xe0, 0x42, 0xce,
	0x84, 0x49, 0x86, 0x58, 0x9d, 0x1f, 0x56, 0x53, 0x9d, 0x30, 0xde, 0xe9, 0x22, 0x1c, 0xa9, 0x61,
	0xc3, 0x0a, 0x95, 0x09, 0x93, 0x6c, 0xde, 0x44, 0xe0, 0x10, 0xcc, 0x8d, 0x6e, 0xae, 0x39, 0xcb,
	0x93, 0xdb, 0x58, 0x50, 0x7f, 0x20, 0x07, 0x70, 0x97, 0x6b, 0x35, 0x91, 0xd3, 0xda, 0x88, 0x2c,
	0xb4, 0xd2, 0x4f, 0xf8, 0x16, 0xce, 0xdd, 0x4e, 0xcb, 0x3e, 0x75, 0xe4, 0x53, 0x1c, 0xf8, 0x47,
	0xb0, 0x7d, 0xee, 0x56, 0xf3, 0x79, 0x24, 0x04, 0xaf, 0x90, 0x96, 0xfa, 0x2a, 0x30, 0xa3, 0xc7,
	0x70, 0xa7, 0xc9, 0xab, 0x60, 0x9c, 0xb2, 0x2c, 0x33, 0xb4, 0x9a, 0x97, 0x82, 0x24, 0xd0, 0x29,
	0x18, 0x77, 0xe7, 0xa0, 0x25, 0xcd, 0x71, 0xf4, 0xa9, 0x57, 0x32, 0x77, 0x45, 0x96, 0xb3, 0x4f,
	0xce, 0xdd, 0xb9, 0x0f, 0x83, 0x05, 0x22, 0xac, 0x0d, 0x17, 0xfb, 0x0e, 0x3b, 0xf4, 0xd0, 0xe8,
	0xcf, 0xb5, 0x56, 0x60, 0x82, 0x19, 0x79, 0x00, 0xb1, 0x50, 0x9c, 0x95, 0xb6, 0xce, 0x59, 0x25,
	0xb5, 0x6a, 0xc4, 0x6b, 0x09, 0x24, 0x0f, 0x61, 0x33, 0x5c, 0xa0, 0xb5, 0x7a, 0xad, 0xf4, 0x5b,
	0x15, 0xf4, 0x6b, 0x18, 0xe0, 0x6f, 0x3d, 0x4a, 0x7e, 0x89, 0xa0, 0xdf, 0xbc, 0xc5, 0x65, 0xb1,
	0x86, 0xcd, 0xe4, 0xab, 0x69, 0xe6, 0x52, 0xd1, 0x52, 0x28, 0x18, 0x0f, 0x2f, 0x0d, 0xc5, 0x38,
	0x58, 0xa4, 0x71, 0x6b, 0x51, 0x8c, 0x83, 0xc6, 0xe4, 0xb7, 0xe8, 0x42, 0xc1, 0xd6, 0x31, 0xd5,
	0x6c, 0x35, 0xa9, 0x2e, 0x37, 0x6b, 0xb9, 0x2d, 0x79, 0x3b, 0x06, 0xb5, 0x92, 0x9c, 0xd9, 0x8a,
	0x4e, 0x0d, 0x53, 0x15, 0xf9, 0x18, 0x88, 0x93, 0x47, 0x3c, 0x50, 0x0c, 0x31, 0x63, 0x79, 0x68,
	0xd0, 0xed, 0x5c, 0x4f, 0x8f, 0x1d, 0xf1, 0x2c, 0xe0, 0x4e, 0x24, 0xbc, 0x65, 0x56, 0x1b, 0xdf,
	0x4a, 0x2f, 0xe3, 0x31, 0xa2, 0x47, 0x01, 0x1c, 0xfd, 0xdd, 0x6b, 0xc3, 0x35, 0xc9, 0xfa, 0xed,
	0xf0, 0x53, 0x04, 0xf1, 0x12, 0x82, 0xa1, 0xfa, 0xfb, 0xdf, 0xaf, 0x50, 0x21, 0xd3, 0x41, 0x60,
	0xf1, 0x3b, 0xc2, 0xbe, 0x59, 0xfa, 0xa6, 0x66, 0xb9, 0x9c, 0x48, 0xb1, 0x58, 0x92, 0xd2, 0x7e,
	0xdd, 0x40, 0xb8, 0xb1, 0x2d, 0xd6, 0x24, 0xf3, 0x0e, 0x71, 0xc6, 0xba, 0x69, 0x2c, 0xed, 0x71,
	0x0b, 0xfa, 0x8f, 0xb1, 0x28, 0xb0, 0xa0, 0xee, 0xdd, 0xb4, 0xd0, 0x99, 0xc8, 0xc3, 0x20, 0x90,
	0x25, 0xea, 0x85, 0x63, 0xdc, 0x6e, 0x94, 0x96, 0xfa, 0xd9, 0x5e, 0x47, 0x8f, 0x1d, 0x69, 0x9f,
	0xe3, 0x50, 0x3f, 0x08, 0xf2, 0x8b, 0xa4, 0x7b, 0x4e, 0x58, 0x54, 0x03, 0x87, 0xa2, 0xc9, 0x17,
	0xda, 0xb8, 0xc4, 0x9a, 0x97, 0x66, 0xba, 0x60, 0x52, 0x85, 0x25, 0xd5, 0xd4, 0xf4, 0x08, 0x41,
	0xf2, 0x16, 0x3a, 0xcd, 0xc4, 0x75, 0x57, 0xa9, 0x74, 0x21, 0x48, 0xda, 0x44, 0xbb, 0x4a, 0x9e,
	0x7a, 0x57, 0xc9, 0xd3, 0x35, 0x2a, 0x08, 0xd7, 0xa8, 0xe0, 0x43, 0xd8, 0xcc, 0x44, 0xce, 0xe6,
	0x94, 0xd9, 0x79, 0x51, 0x88, 0xca, 0xcc, 0x71, 0x59, 0x6d, 0xa5, 0x43, 0x84, 0x0f, 0x1b, 0x94,
	0xfc, 0x11, 0xc1, 0x90, 0x29, 0xa5, 0x6b, 0xc5, 0x85, 0x9f, 0x71, 0xdc, 0x59, 0x2b, 0x53, 0x8b,
	0xa5, 0x6f, 0x2b, 0x8d, 0x9b, 0xd0, 0xf8, 0x11, 0x91, 0x9f, 0x23, 0x00, 0x3b, 0x57, 0x3c, 0x24,
	0x12, 0xdf, 0x5c, 0x22, 0x3d, 0x17, 0xd6, 0x27, 0xf1, 0x57, 0x04, 0x3b, 0xbe, 0x76, 0x46, 0xd8,
	0x52, 0x2b, 0xdb, 0xd4, 0x65, 0x78, 0x73, 0xe9, 0x10, 0x4c, 0x20, 0x0d, 0xf1, 0x8f, 0x1b, 0x1d,
	0x2a, 0x2b, 0x3b, 0xa1, 0xb9, 0xb6, 0x96, 0x36, 0x75, 0xc3, 0x3d, 0xdd, 0x4d, 0x6f, 0x3b, 0xe6,
	0x44, 0x5b, 0x7b, 0x18, 0x70, 0xf7, 0xb5, 0xb4, 0xd6, 0xee, 0x71, 0x61, 0xb7, 0x0e, 0x1a, 0xcb,
	0xd3, 0xb9, 0xe2, 0xe4, 0x0e, 0x6c, 0x48, 0x4b, 0x33, 0x55, 0xe3, 0x4a, 0xed, 0xa6, 0xeb, 0xd2,
	0x1e, 0xa9, 0x7a, 0xf4, 0x6f, 0x04, 0x5b, 0x97, 0x7e, 0xaa, 0x5f, 0xfc, 0xa9, 0xbb, 0x7f, 0xf1,
	0xa7, 0x2e, 0xf9, 0xfd, 0x92, 0x74, 0x3d, 0xd9, 0x5d, 0x5b, 0x5d, 0xc9, 0x96, 0x42, 0x2d, 0x4b,
	0xd8, 0xd9, 0x06, 0xfe, 0x81, 0x79, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x08, 0x40,
	0x99, 0xdb, 0x0c, 0x00, 0x00,
}
