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
// source: ptp_if_info.proto

package cisco_ios_xr_ptp_oper_ptp_nodes_node_node_interfaces_node_interface

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

type PtpIfInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfInfo_KEYS) Reset()         { *m = PtpIfInfo_KEYS{} }
func (m *PtpIfInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfInfo_KEYS) ProtoMessage()    {}
func (*PtpIfInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_044daa3b11182e32, []int{0}
}

func (m *PtpIfInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpIfInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfInfo_KEYS.Merge(m, src)
}
func (m *PtpIfInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfInfo_KEYS.Size(m)
}
func (m *PtpIfInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfInfo_KEYS proto.InternalMessageInfo

func (m *PtpIfInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PtpIfInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
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
	return fileDescriptor_044daa3b11182e32, []int{1}
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
	return fileDescriptor_044daa3b11182e32, []int{2}
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
	return fileDescriptor_044daa3b11182e32, []int{3}
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

type PtpBagMasterTableEntry struct {
	Address              *PtpBagAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CommunicationModel   string         `protobuf:"bytes,2,opt,name=communication_model,json=communicationModel,proto3" json:"communication_model,omitempty"`
	Priority             uint32         `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Known                bool           `protobuf:"varint,4,opt,name=known,proto3" json:"known,omitempty"`
	Qualified            bool           `protobuf:"varint,5,opt,name=qualified,proto3" json:"qualified,omitempty"`
	IsGrandmaster        bool           `protobuf:"varint,6,opt,name=is_grandmaster,json=isGrandmaster,proto3" json:"is_grandmaster,omitempty"`
	PtsfLossAnnounce     uint32         `protobuf:"varint,7,opt,name=ptsf_loss_announce,json=ptsfLossAnnounce,proto3" json:"ptsf_loss_announce,omitempty"`
	PtsfLossSync         uint32         `protobuf:"varint,8,opt,name=ptsf_loss_sync,json=ptsfLossSync,proto3" json:"ptsf_loss_sync,omitempty"`
	IsNonnegotiated      bool           `protobuf:"varint,9,opt,name=is_nonnegotiated,json=isNonnegotiated,proto3" json:"is_nonnegotiated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PtpBagMasterTableEntry) Reset()         { *m = PtpBagMasterTableEntry{} }
func (m *PtpBagMasterTableEntry) String() string { return proto.CompactTextString(m) }
func (*PtpBagMasterTableEntry) ProtoMessage()    {}
func (*PtpBagMasterTableEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_044daa3b11182e32, []int{4}
}

func (m *PtpBagMasterTableEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagMasterTableEntry.Unmarshal(m, b)
}
func (m *PtpBagMasterTableEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagMasterTableEntry.Marshal(b, m, deterministic)
}
func (m *PtpBagMasterTableEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagMasterTableEntry.Merge(m, src)
}
func (m *PtpBagMasterTableEntry) XXX_Size() int {
	return xxx_messageInfo_PtpBagMasterTableEntry.Size(m)
}
func (m *PtpBagMasterTableEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagMasterTableEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagMasterTableEntry proto.InternalMessageInfo

func (m *PtpBagMasterTableEntry) GetAddress() *PtpBagAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PtpBagMasterTableEntry) GetCommunicationModel() string {
	if m != nil {
		return m.CommunicationModel
	}
	return ""
}

func (m *PtpBagMasterTableEntry) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *PtpBagMasterTableEntry) GetKnown() bool {
	if m != nil {
		return m.Known
	}
	return false
}

func (m *PtpBagMasterTableEntry) GetQualified() bool {
	if m != nil {
		return m.Qualified
	}
	return false
}

func (m *PtpBagMasterTableEntry) GetIsGrandmaster() bool {
	if m != nil {
		return m.IsGrandmaster
	}
	return false
}

func (m *PtpBagMasterTableEntry) GetPtsfLossAnnounce() uint32 {
	if m != nil {
		return m.PtsfLossAnnounce
	}
	return 0
}

func (m *PtpBagMasterTableEntry) GetPtsfLossSync() uint32 {
	if m != nil {
		return m.PtsfLossSync
	}
	return 0
}

func (m *PtpBagMasterTableEntry) GetIsNonnegotiated() bool {
	if m != nil {
		return m.IsNonnegotiated
	}
	return false
}

type PtpIfInfo struct {
	PortState                  string                    `protobuf:"bytes,50,opt,name=port_state,json=portState,proto3" json:"port_state,omitempty"`
	PortNumber                 uint32                    `protobuf:"varint,51,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	LineState                  string                    `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	Encapsulation              string                    `protobuf:"bytes,53,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	Ipv6Address                string                    `protobuf:"bytes,54,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	Ipv4Address                string                    `protobuf:"bytes,55,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	MacAddress                 *PtpBagMacAddrType        `protobuf:"bytes,56,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	TwoStep                    bool                      `protobuf:"varint,57,opt,name=two_step,json=twoStep,proto3" json:"two_step,omitempty"`
	CommunicationModel         string                    `protobuf:"bytes,58,opt,name=communication_model,json=communicationModel,proto3" json:"communication_model,omitempty"`
	LogSyncInterval            int32                     `protobuf:"zigzag32,59,opt,name=log_sync_interval,json=logSyncInterval,proto3" json:"log_sync_interval,omitempty"`
	LogAnnounceInterval        int32                     `protobuf:"zigzag32,60,opt,name=log_announce_interval,json=logAnnounceInterval,proto3" json:"log_announce_interval,omitempty"`
	AnnounceTimeout            uint32                    `protobuf:"varint,61,opt,name=announce_timeout,json=announceTimeout,proto3" json:"announce_timeout,omitempty"`
	LogMinDelayRequestInterval int32                     `protobuf:"zigzag32,62,opt,name=log_min_delay_request_interval,json=logMinDelayRequestInterval,proto3" json:"log_min_delay_request_interval,omitempty"`
	MasterTable                []*PtpBagMasterTableEntry `protobuf:"bytes,63,rep,name=master_table,json=masterTable,proto3" json:"master_table,omitempty"`
	ConfiguredPortState        string                    `protobuf:"bytes,64,opt,name=configured_port_state,json=configuredPortState,proto3" json:"configured_port_state,omitempty"`
	SupportsUnicast            bool                      `protobuf:"varint,65,opt,name=supports_unicast,json=supportsUnicast,proto3" json:"supports_unicast,omitempty"`
	SupportsMaster             bool                      `protobuf:"varint,66,opt,name=supports_master,json=supportsMaster,proto3" json:"supports_master,omitempty"`
	SupportsOneStep            bool                      `protobuf:"varint,67,opt,name=supports_one_step,json=supportsOneStep,proto3" json:"supports_one_step,omitempty"`
	SupportsTwoStep            bool                      `protobuf:"varint,68,opt,name=supports_two_step,json=supportsTwoStep,proto3" json:"supports_two_step,omitempty"`
	SupportsEthernet           bool                      `protobuf:"varint,69,opt,name=supports_ethernet,json=supportsEthernet,proto3" json:"supports_ethernet,omitempty"`
	SupportsMulticast          bool                      `protobuf:"varint,70,opt,name=supports_multicast,json=supportsMulticast,proto3" json:"supports_multicast,omitempty"`
	SupportsIpv4               bool                      `protobuf:"varint,71,opt,name=supports_ipv4,json=supportsIpv4,proto3" json:"supports_ipv4,omitempty"`
	SupportsIpv6               bool                      `protobuf:"varint,72,opt,name=supports_ipv6,json=supportsIpv6,proto3" json:"supports_ipv6,omitempty"`
	SupportsSlave              bool                      `protobuf:"varint,73,opt,name=supports_slave,json=supportsSlave,proto3" json:"supports_slave,omitempty"`
	SupportsSourceIp           bool                      `protobuf:"varint,74,opt,name=supports_source_ip,json=supportsSourceIp,proto3" json:"supports_source_ip,omitempty"`
	MaxSyncRate                uint32                    `protobuf:"varint,75,opt,name=max_sync_rate,json=maxSyncRate,proto3" json:"max_sync_rate,omitempty"`
	EventCos                   uint32                    `protobuf:"varint,76,opt,name=event_cos,json=eventCos,proto3" json:"event_cos,omitempty"`
	GeneralCos                 uint32                    `protobuf:"varint,77,opt,name=general_cos,json=generalCos,proto3" json:"general_cos,omitempty"`
	EventDscp                  uint32                    `protobuf:"varint,78,opt,name=event_dscp,json=eventDscp,proto3" json:"event_dscp,omitempty"`
	GeneralDscp                uint32                    `protobuf:"varint,79,opt,name=general_dscp,json=generalDscp,proto3" json:"general_dscp,omitempty"`
	UnicastPeers               uint32                    `protobuf:"varint,80,opt,name=unicast_peers,json=unicastPeers,proto3" json:"unicast_peers,omitempty"`
	LocalPriority              uint32                    `protobuf:"varint,81,opt,name=local_priority,json=localPriority,proto3" json:"local_priority,omitempty"`
	SignalFail                 bool                      `protobuf:"varint,82,opt,name=signal_fail,json=signalFail,proto3" json:"signal_fail,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                  `json:"-"`
	XXX_unrecognized           []byte                    `json:"-"`
	XXX_sizecache              int32                     `json:"-"`
}

func (m *PtpIfInfo) Reset()         { *m = PtpIfInfo{} }
func (m *PtpIfInfo) String() string { return proto.CompactTextString(m) }
func (*PtpIfInfo) ProtoMessage()    {}
func (*PtpIfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_044daa3b11182e32, []int{5}
}

func (m *PtpIfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfInfo.Unmarshal(m, b)
}
func (m *PtpIfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfInfo.Marshal(b, m, deterministic)
}
func (m *PtpIfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfInfo.Merge(m, src)
}
func (m *PtpIfInfo) XXX_Size() int {
	return xxx_messageInfo_PtpIfInfo.Size(m)
}
func (m *PtpIfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfInfo proto.InternalMessageInfo

func (m *PtpIfInfo) GetPortState() string {
	if m != nil {
		return m.PortState
	}
	return ""
}

func (m *PtpIfInfo) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

func (m *PtpIfInfo) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *PtpIfInfo) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *PtpIfInfo) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *PtpIfInfo) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PtpIfInfo) GetMacAddress() *PtpBagMacAddrType {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *PtpIfInfo) GetTwoStep() bool {
	if m != nil {
		return m.TwoStep
	}
	return false
}

func (m *PtpIfInfo) GetCommunicationModel() string {
	if m != nil {
		return m.CommunicationModel
	}
	return ""
}

func (m *PtpIfInfo) GetLogSyncInterval() int32 {
	if m != nil {
		return m.LogSyncInterval
	}
	return 0
}

func (m *PtpIfInfo) GetLogAnnounceInterval() int32 {
	if m != nil {
		return m.LogAnnounceInterval
	}
	return 0
}

func (m *PtpIfInfo) GetAnnounceTimeout() uint32 {
	if m != nil {
		return m.AnnounceTimeout
	}
	return 0
}

func (m *PtpIfInfo) GetLogMinDelayRequestInterval() int32 {
	if m != nil {
		return m.LogMinDelayRequestInterval
	}
	return 0
}

func (m *PtpIfInfo) GetMasterTable() []*PtpBagMasterTableEntry {
	if m != nil {
		return m.MasterTable
	}
	return nil
}

func (m *PtpIfInfo) GetConfiguredPortState() string {
	if m != nil {
		return m.ConfiguredPortState
	}
	return ""
}

func (m *PtpIfInfo) GetSupportsUnicast() bool {
	if m != nil {
		return m.SupportsUnicast
	}
	return false
}

func (m *PtpIfInfo) GetSupportsMaster() bool {
	if m != nil {
		return m.SupportsMaster
	}
	return false
}

func (m *PtpIfInfo) GetSupportsOneStep() bool {
	if m != nil {
		return m.SupportsOneStep
	}
	return false
}

func (m *PtpIfInfo) GetSupportsTwoStep() bool {
	if m != nil {
		return m.SupportsTwoStep
	}
	return false
}

func (m *PtpIfInfo) GetSupportsEthernet() bool {
	if m != nil {
		return m.SupportsEthernet
	}
	return false
}

func (m *PtpIfInfo) GetSupportsMulticast() bool {
	if m != nil {
		return m.SupportsMulticast
	}
	return false
}

func (m *PtpIfInfo) GetSupportsIpv4() bool {
	if m != nil {
		return m.SupportsIpv4
	}
	return false
}

func (m *PtpIfInfo) GetSupportsIpv6() bool {
	if m != nil {
		return m.SupportsIpv6
	}
	return false
}

func (m *PtpIfInfo) GetSupportsSlave() bool {
	if m != nil {
		return m.SupportsSlave
	}
	return false
}

func (m *PtpIfInfo) GetSupportsSourceIp() bool {
	if m != nil {
		return m.SupportsSourceIp
	}
	return false
}

func (m *PtpIfInfo) GetMaxSyncRate() uint32 {
	if m != nil {
		return m.MaxSyncRate
	}
	return 0
}

func (m *PtpIfInfo) GetEventCos() uint32 {
	if m != nil {
		return m.EventCos
	}
	return 0
}

func (m *PtpIfInfo) GetGeneralCos() uint32 {
	if m != nil {
		return m.GeneralCos
	}
	return 0
}

func (m *PtpIfInfo) GetEventDscp() uint32 {
	if m != nil {
		return m.EventDscp
	}
	return 0
}

func (m *PtpIfInfo) GetGeneralDscp() uint32 {
	if m != nil {
		return m.GeneralDscp
	}
	return 0
}

func (m *PtpIfInfo) GetUnicastPeers() uint32 {
	if m != nil {
		return m.UnicastPeers
	}
	return 0
}

func (m *PtpIfInfo) GetLocalPriority() uint32 {
	if m != nil {
		return m.LocalPriority
	}
	return 0
}

func (m *PtpIfInfo) GetSignalFail() bool {
	if m != nil {
		return m.SignalFail
	}
	return false
}

func init() {
	proto.RegisterType((*PtpIfInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_if_info_KEYS")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_bag_address")
	proto.RegisterType((*PtpBagMasterTableEntry)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_bag_master_table_entry")
	proto.RegisterType((*PtpIfInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interfaces.node_interface.ptp_if_info")
}

func init() { proto.RegisterFile("ptp_if_info.proto", fileDescriptor_044daa3b11182e32) }

var fileDescriptor_044daa3b11182e32 = []byte{
	// 1042 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0x6b, 0x6f, 0x1c, 0x35,
	0x17, 0xc7, 0x95, 0x27, 0x6d, 0xb3, 0x7b, 0x36, 0x9b, 0x8b, 0xf3, 0x14, 0x99, 0x94, 0xd2, 0xb0,
	0x2d, 0x22, 0x29, 0x10, 0x44, 0x1a, 0xc2, 0xa5, 0xdc, 0xd2, 0x24, 0x2d, 0xa1, 0x4d, 0x1a, 0x66,
	0x53, 0x24, 0xe0, 0x85, 0xe5, 0xcc, 0x7a, 0x17, 0x8b, 0x19, 0xdb, 0x1d, 0x7b, 0x92, 0xac, 0x90,
	0x10, 0xe2, 0x2b, 0xf1, 0x89, 0xf8, 0x26, 0xc8, 0xc7, 0x9e, 0xd9, 0xdd, 0xa4, 0x7d, 0x17, 0xf1,
	0x66, 0xa5, 0xf9, 0x9d, 0xff, 0xf1, 0xe5, 0xdc, 0xbc, 0xb0, 0x68, 0x9c, 0x61, 0xb2, 0xcf, 0xa4,
	0xea, 0xeb, 0x75, 0x53, 0x68, 0xa7, 0xc9, 0x4e, 0x2a, 0x6d, 0xaa, 0x99, 0xd4, 0x96, 0x9d, 0x17,
	0xcc, 0xdb, 0xb5, 0x11, 0xc5, 0xba, 0x71, 0x66, 0x5d, 0xe9, 0x9e, 0xb0, 0xf8, 0x8b, 0x3f, 0x4c,
	0x2a, 0x27, 0x8a, 0x3e, 0x4f, 0x23, 0x1c, 0x7d, 0x77, 0x7e, 0x84, 0x85, 0xb1, 0x95, 0xd9, 0xd3,
	0xbd, 0x9f, 0xba, 0xe4, 0x16, 0x34, 0x51, 0xa5, 0x78, 0x2e, 0xe8, 0xd4, 0xca, 0xd4, 0x6a, 0x33,
	0x69, 0x78, 0x70, 0xc8, 0x73, 0x41, 0xde, 0x85, 0xb9, 0xda, 0x3b, 0x28, 0xfe, 0x87, 0x8a, 0x76,
	0x4d, 0xbd, 0xac, 0xf3, 0x31, 0xdc, 0xf4, 0xeb, 0x9e, 0xf0, 0x01, 0xcb, 0x79, 0xca, 0x78, 0xaf,
	0x57, 0x30, 0x37, 0x34, 0x82, 0x50, 0x98, 0xc9, 0x79, 0xea, 0xbf, 0xe3, 0xd2, 0xd5, 0x67, 0xe7,
	0x21, 0xbc, 0x51, 0xb9, 0x48, 0x73, 0xba, 0x35, 0xe6, 0xf3, 0x0e, 0xcc, 0xd6, 0x44, 0x58, 0x1b,
	0x1d, 0x5b, 0x9e, 0x6d, 0x07, 0xd4, 0xf9, 0x73, 0x1a, 0xe6, 0x2b, 0xef, 0x28, 0x23, 0xf7, 0xa0,
	0x2d, 0x54, 0xca, 0x8d, 0x2d, 0x33, 0xee, 0xa4, 0x56, 0xd1, 0x6f, 0x12, 0x92, 0xf7, 0x60, 0x3e,
	0x3a, 0xb0, 0x52, 0xfd, 0xa6, 0xf4, 0x99, 0xc2, 0x1b, 0x35, 0x92, 0xb9, 0x88, 0x5f, 0x04, 0x4a,
	0x7e, 0x87, 0x56, 0x75, 0x15, 0x7f, 0x88, 0xe9, 0x95, 0xa9, 0xd5, 0xd6, 0xc6, 0xcf, 0xeb, 0x57,
	0x90, 0x85, 0xf5, 0x57, 0x86, 0x2a, 0x81, 0x9c, 0xa7, 0xf1, 0x7e, 0x31, 0x04, 0x9b, 0xf5, 0xee,
	0xd7, 0xea, 0x10, 0x6c, 0x56, 0x92, 0x3f, 0x2e, 0x44, 0xe9, 0x3a, 0x1e, 0xf0, 0x97, 0x2b, 0x3d,
	0xe0, 0x64, 0x62, 0x26, 0x53, 0xf0, 0xf7, 0x34, 0x2c, 0x8f, 0x2e, 0x62, 0x9d, 0x28, 0x98, 0xe3,
	0x27, 0x99, 0x60, 0x42, 0xb9, 0x62, 0x48, 0x14, 0xcc, 0x8c, 0xe7, 0xaf, 0xb5, 0x71, 0x7c, 0xa5,
	0x27, 0x8b, 0x6b, 0x27, 0xd5, 0x26, 0xe4, 0x23, 0x58, 0x4a, 0x75, 0x9e, 0x97, 0x4a, 0xa6, 0x98,
	0x68, 0x96, 0xeb, 0x9e, 0xc8, 0x62, 0xb5, 0x92, 0x09, 0xd3, 0x81, 0xb7, 0x90, 0x65, 0x68, 0x98,
	0x42, 0xea, 0x42, 0xba, 0x21, 0x26, 0xb7, 0x9d, 0xd4, 0xdf, 0xe4, 0xff, 0x70, 0x3d, 0x94, 0xc6,
	0x35, 0x2c, 0x8d, 0xf0, 0x41, 0xde, 0x82, 0xe6, 0xcb, 0x92, 0x67, 0xb2, 0x2f, 0x45, 0x0f, 0xc3,
	0xdd, 0x48, 0x46, 0x00, 0x3b, 0xc5, 0xb2, 0x41, 0xc1, 0x55, 0x2f, 0x84, 0x83, 0xde, 0x40, 0x49,
	0x5b, 0xda, 0x27, 0x23, 0x48, 0x3e, 0x00, 0x62, 0x9c, 0xed, 0xb3, 0x4c, 0x5b, 0xcb, 0xb8, 0x52,
	0xba, 0x54, 0xa9, 0xa0, 0x33, 0x78, 0x80, 0x05, 0x6f, 0x79, 0xa6, 0xad, 0xdd, 0x8e, 0x9c, 0xdc,
	0x83, 0xb9, 0x91, 0xda, 0x0e, 0x55, 0x4a, 0x1b, 0xa8, 0x9c, 0xad, 0x94, 0xdd, 0xa1, 0x4a, 0xc9,
	0x1a, 0x2c, 0x48, 0xcb, 0x94, 0x56, 0x4a, 0x0c, 0xb4, 0x93, 0xdc, 0x89, 0x1e, 0x6d, 0xe2, 0xe6,
	0xf3, 0xd2, 0x1e, 0x8e, 0xe3, 0xce, 0x3f, 0x2d, 0x68, 0x8d, 0x4d, 0x00, 0x72, 0x1b, 0xc0, 0xe8,
	0xc2, 0x31, 0xeb, 0xb8, 0x13, 0x74, 0x03, 0xa3, 0xd5, 0xf4, 0xa4, 0xeb, 0x01, 0xb9, 0x03, 0x2d,
	0x34, 0xab, 0x32, 0x3f, 0x11, 0x05, 0x7d, 0x80, 0x9b, 0xa3, 0xc7, 0x21, 0x12, 0xef, 0x9f, 0x49,
	0x25, 0xa2, 0xff, 0x66, 0xf0, 0xf7, 0x24, 0xf8, 0x5f, 0xea, 0xc9, 0x4f, 0x5e, 0xd5, 0x93, 0x17,
	0x1b, 0x7e, 0xeb, 0x52, 0xc3, 0x5f, 0x6a, 0x88, 0x4f, 0x2f, 0x37, 0xc4, 0x85, 0x86, 0xfd, 0xec,
	0x3f, 0x6d, 0xd8, 0x37, 0xa1, 0xe1, 0xce, 0x34, 0xb3, 0x4e, 0x18, 0xfa, 0x39, 0x86, 0x7e, 0xc6,
	0x9d, 0xe9, 0xae, 0x13, 0xe6, 0x75, 0x95, 0xf9, 0xc5, 0x6b, 0x2b, 0xf3, 0x3e, 0x2c, 0x66, 0x7a,
	0x80, 0xe9, 0x0e, 0x67, 0x38, 0xe5, 0x19, 0x7d, 0xb8, 0x32, 0xb5, 0xba, 0x98, 0xcc, 0x67, 0x7a,
	0xe0, 0x53, 0xbe, 0x1f, 0x31, 0xd9, 0x80, 0x9b, 0x5e, 0x5b, 0x15, 0xd2, 0x48, 0xff, 0x25, 0xea,
	0x97, 0x32, 0x3d, 0xa8, 0x8a, 0xa9, 0xf6, 0x59, 0x83, 0x85, 0x5a, 0xef, 0x64, 0x2e, 0x74, 0xe9,
	0xe8, 0x57, 0x98, 0xd9, 0xf9, 0x8a, 0x1f, 0x07, 0x4c, 0x1e, 0xc1, 0xdb, 0x7e, 0xf9, 0x5c, 0x2a,
	0xd6, 0x13, 0x19, 0x1f, 0xb2, 0x42, 0xbc, 0x2c, 0x85, 0x75, 0xa3, 0x7d, 0xbe, 0xc6, 0x7d, 0x96,
	0x33, 0x3d, 0x38, 0x90, 0x6a, 0xd7, 0x6b, 0x92, 0x20, 0xa9, 0xb7, 0xfb, 0x6b, 0x0a, 0x66, 0xc7,
	0x07, 0x04, 0xfd, 0x66, 0x65, 0x7a, 0xb5, 0xb5, 0xc1, 0xae, 0x38, 0x33, 0x17, 0x27, 0x50, 0xd2,
	0x0a, 0xec, 0xd8, 0x23, 0x1f, 0xa7, 0x54, 0xab, 0xbe, 0x1c, 0x94, 0x85, 0xe8, 0xb1, 0xb1, 0x92,
	0xff, 0x16, 0xd3, 0xb0, 0x34, 0x32, 0x1e, 0xd5, 0xc5, 0xbf, 0x06, 0x0b, 0xb6, 0x34, 0x5e, 0xeb,
	0xdf, 0x0a, 0x99, 0x72, 0xeb, 0xe8, 0x76, 0x68, 0xab, 0x8a, 0xbf, 0x08, 0xd8, 0xbf, 0x2a, 0xb5,
	0x34, 0x76, 0xff, 0xa3, 0xf0, 0xaa, 0x54, 0xf8, 0x20, 0xb4, 0xff, 0x7d, 0x58, 0xac, 0x85, 0x1a,
	0xfb, 0x46, 0x18, 0xba, 0x33, 0xb9, 0xe8, 0x73, 0xdf, 0x3d, 0xc2, 0x4c, 0x68, 0xeb, 0xe2, 0xda,
	0x9d, 0xd4, 0x1e, 0xc7, 0x22, 0x7b, 0x7f, 0x4c, 0x2b, 0xdc, 0xaf, 0xa2, 0x50, 0xc2, 0xd1, 0x3d,
	0xd4, 0xd6, 0x97, 0xd8, 0x8b, 0x9c, 0x7c, 0x08, 0x64, 0x74, 0xda, 0x32, 0x73, 0xe1, 0x6a, 0x8f,
	0x51, 0x5d, 0x2f, 0x73, 0x50, 0x19, 0xc8, 0x5d, 0x68, 0xd7, 0x72, 0xdf, 0x70, 0xf4, 0x09, 0x2a,
	0x67, 0x2b, 0xb8, 0x6f, 0x4e, 0x37, 0x2f, 0x8a, 0xb6, 0xe8, 0x77, 0x97, 0x44, 0x5b, 0x7e, 0x46,
	0xd6, 0x22, 0x9b, 0xf1, 0x53, 0x41, 0xf7, 0xc3, 0x8c, 0xac, 0x68, 0xd7, 0x43, 0x3f, 0x23, 0x47,
	0x32, 0x5d, 0x16, 0xbe, 0xae, 0x0d, 0xfd, 0x7e, 0xf2, 0x36, 0x5d, 0x34, 0xec, 0x1b, 0xd2, 0x81,
	0x76, 0xce, 0xcf, 0x43, 0xbb, 0x14, 0x3e, 0xa5, 0x4f, 0xb1, 0x96, 0x5b, 0x39, 0x3f, 0xf7, 0xad,
	0x92, 0xf8, 0x54, 0xde, 0x82, 0xa6, 0x38, 0x15, 0xca, 0xb1, 0x54, 0x5b, 0xfa, 0x2c, 0x4c, 0x7b,
	0x04, 0x3b, 0xda, 0xfa, 0x21, 0x37, 0x10, 0x4a, 0x14, 0x3c, 0x43, 0xf3, 0x41, 0x18, 0x72, 0x11,
	0x79, 0xc1, 0x6d, 0x80, 0xe0, 0xdd, 0xb3, 0xa9, 0xa1, 0x87, 0x68, 0x0f, 0xeb, 0xed, 0xda, 0xd4,
	0xf8, 0xd9, 0x54, 0xf9, 0xa3, 0xe0, 0x79, 0xd8, 0x3f, 0x32, 0x94, 0xdc, 0x85, 0x76, 0xac, 0x20,
	0x66, 0x84, 0x28, 0x2c, 0x3d, 0x0a, 0x63, 0x3c, 0xc2, 0x23, 0xcf, 0x7c, 0x74, 0x32, 0x9d, 0xf2,
	0x8c, 0xd5, 0xef, 0xd2, 0x0f, 0xa8, 0x6a, 0x23, 0x3d, 0xaa, 0x1e, 0xa7, 0x3b, 0xd0, 0xb2, 0x72,
	0xa0, 0x78, 0xc6, 0xfa, 0x5c, 0x66, 0x34, 0xc1, 0xb0, 0x40, 0x40, 0x8f, 0xb9, 0xcc, 0x4e, 0x6e,
	0xe0, 0x1f, 0xc6, 0x07, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x49, 0x4c, 0x53, 0x45, 0x0a,
	0x00, 0x00,
}