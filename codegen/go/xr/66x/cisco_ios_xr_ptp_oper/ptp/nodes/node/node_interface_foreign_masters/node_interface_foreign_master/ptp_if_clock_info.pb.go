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

package cisco_ios_xr_ptp_oper_ptp_nodes_node_node_interface_foreign_masters_node_interface_foreign_master

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
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *PtpIfClockInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

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
	proto.RegisterType((*PtpIfClockInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_if_clock_info_KEYS")
	proto.RegisterType((*PtpBagClockUtcOffset)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_clock_utc_offset")
	proto.RegisterType((*PtpBagPortId)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_port_id")
	proto.RegisterType((*PtpBagClock)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_clock")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_address")
	proto.RegisterType((*PtpBagUnicastGrant)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_unicast_grant")
	proto.RegisterType((*PtpBagForeignClock)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_bag_foreign_clock")
	proto.RegisterType((*PtpIfClockInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_foreign_masters.node_interface_foreign_master.ptp_if_clock_info")
}

func init() { proto.RegisterFile("ptp_if_clock_info.proto", fileDescriptor_51135ae6cdc5ffb0) }

var fileDescriptor_51135ae6cdc5ffb0 = []byte{
	// 1104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4f, 0x6f, 0x1c, 0xb5,
	0x1b, 0xd6, 0xfc, 0xd2, 0x64, 0x77, 0xdf, 0xdd, 0xd9, 0x24, 0x4e, 0xda, 0xdf, 0x00, 0x95, 0x48,
	0x97, 0xa2, 0xf6, 0x80, 0xb6, 0x6a, 0x1a, 0x71, 0xe1, 0x14, 0x35, 0x10, 0x55, 0xfd, 0x03, 0x4c,
	0xa0, 0x12, 0x12, 0x92, 0xe5, 0x78, 0xbc, 0x2b, 0xab, 0x33, 0xf6, 0xd4, 0x9e, 0x59, 0xba, 0xe2,
	0xcc, 0x05, 0x21, 0x6e, 0x08, 0x81, 0x40, 0xdc, 0xf9, 0x10, 0x7c, 0x0a, 0xee, 0x7c, 0x15, 0xe4,
	0xd7, 0x9e, 0x9d, 0x6c, 0xa2, 0xe4, 0xc6, 0xe6, 0xb2, 0x92, 0x9f, 0xe7, 0xb5, 0xfd, 0xf8, 0xfd,
	0x3b, 0x0b, 0xff, 0x2f, 0xab, 0x92, 0xca, 0x09, 0xe5, 0xb9, 0xe6, 0xaf, 0xa8, 0x54, 0x13, 0x3d,
	0x2e, 0x8d, 0xae, 0x34, 0x61, 0x5c, 0x5a, 0xae, 0xa9, 0xd4, 0x96, 0xbe, 0x31, 0xd4, 0x59, 0xe9,
	0x52, 0x98, 0x71, 0x59, 0x95, 0x63, 0xa5, 0x33, 0x61, 0xf1, 0x17, 0x7f, 0xa8, 0x54, 0x95, 0x30,
	0x13, 0xc6, 0x05, 0x9d, 0x68, 0x23, 0xe4, 0x54, 0xd1, 0x82, 0xd9, 0x4a, 0x18, 0x7b, 0x35, 0x3d,
	0xfa, 0x1a, 0x6e, 0x5d, 0xb8, 0x9d, 0x3e, 0xfd, 0xf8, 0xab, 0x13, 0xf2, 0x0e, 0xf4, 0x70, 0xab,
	0x62, 0x85, 0x48, 0xa2, 0xbd, 0xe8, 0x7e, 0x2f, 0xed, 0x3a, 0xe0, 0x05, 0x2b, 0x04, 0x79, 0x1f,
	0x86, 0xed, 0x91, 0x68, 0xf1, 0x3f, 0xb4, 0x88, 0x17, 0xa8, 0x33, 0x1b, 0x65, 0x90, 0xb8, 0xd3,
	0x4f, 0xd9, 0x34, 0x1c, 0x5f, 0x57, 0x9c, 0xea, 0xc9, 0xc4, 0x8a, 0xca, 0x1d, 0xc1, 0x6b, 0x63,
	0x84, 0xaa, 0x02, 0x12, 0x2e, 0x89, 0x03, 0xfa, 0xa9, 0x37, 0xbb, 0x03, 0x03, 0x4f, 0xd3, 0x19,
	0xcb, 0x65, 0x86, 0xf7, 0x74, 0xd3, 0xbe, 0xc7, 0x5e, 0x3a, 0x68, 0xf4, 0x1c, 0x36, 0x9b, 0x5b,
	0x4a, 0x6d, 0x2a, 0x2a, 0x33, 0xf2, 0x16, 0x74, 0xc3, 0x7b, 0x32, 0x3c, 0xf6, 0x46, 0xda, 0xc1,
	0xf5, 0x93, 0x8c, 0xbc, 0x0b, 0x7d, 0xb4, 0x52, 0x75, 0x71, 0x2a, 0x0c, 0x9e, 0x17, 0xa7, 0xe0,
	0xa0, 0x17, 0x88, 0x8c, 0xfe, 0xe9, 0x40, 0xbc, 0xa4, 0xfa, 0xaa, 0xd3, 0x6e, 0x43, 0xaf, 0x34,
	0x52, 0x1b, 0x59, 0xcd, 0x1f, 0x86, 0xb3, 0x5a, 0xe0, 0x2c, 0xbb, 0x9f, 0xac, 0x2d, 0xb3, 0xfb,
	0x64, 0x17, 0xd6, 0x79, 0xce, 0xac, 0x4d, 0x6e, 0x20, 0xe3, 0x17, 0xe4, 0x6d, 0xe8, 0x32, 0xce,
	0x6b, 0xc3, 0xf8, 0x3c, 0x59, 0x47, 0x62, 0xb1, 0x26, 0x63, 0xd8, 0x09, 0xce, 0xc8, 0xf5, 0x94,
	0xce, 0x98, 0x91, 0x4c, 0x71, 0x91, 0x6c, 0xa0, 0xd9, 0xb6, 0xa7, 0x9e, 0xe9, 0xe9, 0xcb, 0x40,
	0x90, 0xf7, 0x20, 0xb6, 0x95, 0x28, 0x2d, 0x35, 0xa2, 0xd0, 0x33, 0x91, 0x25, 0x1d, 0xb4, 0x1c,
	0x20, 0x98, 0x7a, 0xcc, 0x39, 0xa4, 0x92, 0x85, 0xa0, 0x56, 0xd7, 0x86, 0x8b, 0xa4, 0x8b, 0x51,
	0x00, 0x07, 0x9d, 0x20, 0x42, 0x7e, 0x8d, 0x00, 0xda, 0xc0, 0x25, 0xbd, 0xbd, 0xe8, 0x7e, 0x7f,
	0xff, 0xdb, 0xf1, 0x7f, 0x9e, 0x9c, 0xe3, 0xcb, 0x72, 0x27, 0xed, 0xd5, 0x15, 0x0f, 0xf9, 0xf1,
	0x00, 0x76, 0x26, 0x46, 0xbc, 0xae, 0x85, 0xe2, 0x73, 0x5a, 0x19, 0xc6, 0x05, 0x3b, 0xcd, 0x45,
	0x02, 0x98, 0x26, 0x64, 0x41, 0x7d, 0xd1, 0x30, 0x2e, 0xef, 0xf0, 0xb9, 0xad, 0x6d, 0x1f, 0x6d,
	0x63, 0x87, 0xb6, 0x66, 0xb7, 0xa1, 0xe7, 0x00, 0xcb, 0x59, 0x2e, 0x92, 0x01, 0xfa, 0xa4, 0x05,
	0x5c, 0x56, 0xe6, 0x82, 0x95, 0xd4, 0x0a, 0xae, 0x55, 0x66, 0x93, 0x18, 0x0d, 0xfa, 0x0e, 0x3b,
	0xf1, 0x10, 0xf9, 0x31, 0x82, 0xae, 0x11, 0x5c, 0xc8, 0x99, 0x30, 0xc9, 0x10, 0x7d, 0x66, 0x56,
	0xe8, 0xb3, 0x50, 0x09, 0xe9, 0x42, 0x03, 0xf9, 0x3e, 0x82, 0x0d, 0x2b, 0x54, 0x26, 0x4c, 0xb2,
	0x79, 0x6d, 0x72, 0x82, 0x02, 0x97, 0xfb, 0xb9, 0xe6, 0x2c, 0x4f, 0xb6, 0xd0, 0xf9, 0x7e, 0x41,
	0x0e, 0xe0, 0x16, 0xd7, 0x6a, 0x22, 0xa7, 0xb5, 0x11, 0x59, 0x08, 0xbb, 0x2f, 0x91, 0x6d, 0x4c,
	0xdc, 0xdd, 0x96, 0x7d, 0xec, 0xc8, 0xc7, 0x58, 0x31, 0x0f, 0x60, 0xe7, 0xcc, 0xae, 0xa6, 0xbe,
	0x12, 0x82, 0x5b, 0x48, 0x4b, 0x7d, 0x16, 0x98, 0xd1, 0x43, 0xb8, 0xd9, 0xe8, 0x2a, 0x18, 0xa7,
	0x2c, 0xcb, 0x0c, 0xad, 0xe6, 0xa5, 0x20, 0x09, 0x74, 0x0a, 0xc6, 0xdd, 0x3a, 0x34, 0xa3, 0x66,
	0x39, 0xfa, 0xc8, 0xf7, 0x49, 0xb7, 0x45, 0x96, 0xb3, 0x0f, 0xcf, 0xec, 0xb9, 0x03, 0x83, 0x05,
	0x22, 0xac, 0x0d, 0x1b, 0xfb, 0x0e, 0x3b, 0xf4, 0xd0, 0xe8, 0xaf, 0xb5, 0xb6, 0x43, 0x05, 0x33,
	0x72, 0x17, 0x62, 0xa1, 0x38, 0x2b, 0x6d, 0x9d, 0xb3, 0x4a, 0x6a, 0xd5, 0x74, 0xbf, 0x25, 0x90,
	0xdc, 0x83, 0xcd, 0xb0, 0x81, 0xd6, 0xea, 0x95, 0xd2, 0xdf, 0xa8, 0xd0, 0x00, 0x87, 0x01, 0xfe,
	0xd2, 0xa3, 0xe4, 0x97, 0x08, 0xfa, 0xcd, 0x5b, 0x9c, 0x8a, 0x35, 0x8c, 0xf0, 0x9b, 0x15, 0x46,
	0x78, 0xc9, 0x93, 0x29, 0x14, 0x8c, 0x87, 0xe7, 0x07, 0x0f, 0x1d, 0x2c, 0xb4, 0xdd, 0x58, 0x78,
	0xe8, 0xa0, 0x31, 0xf9, 0x2d, 0x3a, 0xe7, 0xc5, 0x75, 0xd4, 0x3f, 0x5f, 0xa1, 0xfe, 0xe5, 0xb0,
	0x2e, 0x07, 0x30, 0x6f, 0x13, 0xa6, 0x56, 0x92, 0x33, 0x5b, 0xd1, 0xa9, 0x61, 0xaa, 0x22, 0x1f,
	0x00, 0x71, 0x9d, 0x18, 0x17, 0xfe, 0x8a, 0x19, 0xcb, 0x43, 0x28, 0xb7, 0x72, 0x3d, 0x3d, 0x76,
	0xc4, 0x93, 0x80, 0xbb, 0xd6, 0xe3, 0x2d, 0xb3, 0xda, 0xf8, 0xa0, 0xfb, 0x89, 0x11, 0x23, 0x7a,
	0x14, 0xc0, 0xd1, 0x77, 0xd0, 0x5e, 0xd7, 0x88, 0xf5, 0x83, 0xe8, 0xa7, 0x08, 0xe2, 0x25, 0x04,
	0xaf, 0xea, 0xef, 0x97, 0xab, 0x6e, 0xc6, 0xe9, 0x20, 0xb0, 0x58, 0x86, 0x18, 0x61, 0x4b, 0x5f,
	0xd7, 0x2c, 0x97, 0x13, 0x29, 0x16, 0x43, 0x5a, 0xda, 0xcf, 0x1b, 0x08, 0xbf, 0x18, 0x2c, 0x3a,
	0x2a, 0xf3, 0x07, 0x62, 0x8a, 0x76, 0xd3, 0x58, 0xda, 0xe3, 0x16, 0xf4, 0xb5, 0x5c, 0x14, 0xe8,
	0x65, 0xe7, 0x0c, 0x5a, 0xe8, 0x4c, 0xe4, 0x21, 0x65, 0xc8, 0x12, 0xf5, 0xdc, 0x31, 0x6e, 0x36,
	0x4b, 0x4b, 0x7d, 0x69, 0xac, 0xe3, 0x89, 0x1d, 0x69, 0x9f, 0x62, 0x4d, 0xdc, 0x0d, 0x9d, 0x1e,
	0x49, 0xf7, 0x9c, 0x30, 0x28, 0x07, 0x0e, 0x45, 0x93, 0x4f, 0xb4, 0x71, 0xc2, 0x9a, 0x97, 0x66,
	0xba, 0x60, 0x52, 0x85, 0x21, 0xd9, 0x38, 0xfa, 0x08, 0x41, 0xf2, 0x43, 0x04, 0x9d, 0x26, 0x39,
	0xbb, 0x2b, 0x6f, 0x9f, 0xe1, 0xe6, 0xb4, 0x91, 0x70, 0x59, 0xcf, 0xeb, 0x5d, 0xd6, 0xf3, 0xae,
	0x68, 0xad, 0x70, 0x45, 0x6b, 0xbd, 0x07, 0x9b, 0x99, 0xc8, 0xd9, 0x9c, 0x32, 0x3b, 0x2f, 0x0a,
	0x51, 0x99, 0x39, 0x4e, 0xcb, 0xed, 0x74, 0x88, 0xf0, 0x61, 0x83, 0x92, 0x3f, 0x22, 0x18, 0x32,
	0xa5, 0x74, 0xad, 0xb8, 0xf0, 0xe5, 0x80, 0x43, 0x73, 0xb5, 0x2d, 0x68, 0xa9, 0x36, 0xd3, 0xb8,
	0xd1, 0x83, 0x45, 0x48, 0x7e, 0x8e, 0x00, 0xec, 0x5c, 0xf1, 0xa0, 0x2e, 0xbe, 0x66, 0x75, 0x3d,
	0xa7, 0xc5, 0x2b, 0xfb, 0x33, 0x82, 0x5d, 0xef, 0x65, 0x23, 0x6c, 0xa9, 0x95, 0x6d, 0x3c, 0x38,
	0xbc, 0x66, 0x8d, 0x04, 0x55, 0xa5, 0x41, 0xd4, 0x71, 0xd3, 0xf1, 0xca, 0xca, 0x4e, 0x68, 0xae,
	0xad, 0xa5, 0x8d, 0x87, 0xf1, 0x83, 0xa2, 0x9b, 0x6e, 0x39, 0xe6, 0x99, 0xb6, 0xf6, 0x30, 0xe0,
	0xae, 0x04, 0x5b, 0x6b, 0xf7, 0xe2, 0x30, 0xef, 0x07, 0x8d, 0xe5, 0xc9, 0x5c, 0x71, 0x72, 0x13,
	0x36, 0xa4, 0xa5, 0x99, 0xaa, 0x71, 0xcc, 0x77, 0xd3, 0x75, 0x69, 0x8f, 0x54, 0x3d, 0xfa, 0x3b,
	0x82, 0xed, 0x0b, 0x7f, 0x4e, 0xce, 0x7f, 0xbf, 0xef, 0x9f, 0xff, 0x7e, 0x27, 0xbf, 0x5f, 0x68,
	0x92, 0x8f, 0xf6, 0xd6, 0x56, 0xec, 0xc7, 0xa5, 0xfb, 0x97, 0x9b, 0xe5, 0xe9, 0x06, 0xfe, 0xb9,
	0x7b, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xc9, 0xaf, 0xa0, 0xf7, 0x0d, 0x00, 0x00,
}
