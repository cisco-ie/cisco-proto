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
// source: eigrp_show_interfaces_bd.proto

package cisco_ios_xr_eigrp_oper_eigrp_processes_process_vrfs_vrf_afs_af_ases_as_interfaces_interface

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

type EigrpShowInterfacesBd_KEYS struct {
	ProcessId            string   `protobuf:"bytes,1,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Asn                  uint32   `protobuf:"varint,4,opt,name=asn,proto3" json:"asn,omitempty"`
	InterfaceName        string   `protobuf:"bytes,5,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EigrpShowInterfacesBd_KEYS) Reset()         { *m = EigrpShowInterfacesBd_KEYS{} }
func (m *EigrpShowInterfacesBd_KEYS) String() string { return proto.CompactTextString(m) }
func (*EigrpShowInterfacesBd_KEYS) ProtoMessage()    {}
func (*EigrpShowInterfacesBd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06608898067dd8f, []int{0}
}

func (m *EigrpShowInterfacesBd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EigrpShowInterfacesBd_KEYS.Unmarshal(m, b)
}
func (m *EigrpShowInterfacesBd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EigrpShowInterfacesBd_KEYS.Marshal(b, m, deterministic)
}
func (m *EigrpShowInterfacesBd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EigrpShowInterfacesBd_KEYS.Merge(m, src)
}
func (m *EigrpShowInterfacesBd_KEYS) XXX_Size() int {
	return xxx_messageInfo_EigrpShowInterfacesBd_KEYS.Size(m)
}
func (m *EigrpShowInterfacesBd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_EigrpShowInterfacesBd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_EigrpShowInterfacesBd_KEYS proto.InternalMessageInfo

func (m *EigrpShowInterfacesBd_KEYS) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *EigrpShowInterfacesBd_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *EigrpShowInterfacesBd_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *EigrpShowInterfacesBd_KEYS) GetAsn() uint32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *EigrpShowInterfacesBd_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type EigrpAddrBd struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,2,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EigrpAddrBd) Reset()         { *m = EigrpAddrBd{} }
func (m *EigrpAddrBd) String() string { return proto.CompactTextString(m) }
func (*EigrpAddrBd) ProtoMessage()    {}
func (*EigrpAddrBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06608898067dd8f, []int{1}
}

func (m *EigrpAddrBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EigrpAddrBd.Unmarshal(m, b)
}
func (m *EigrpAddrBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EigrpAddrBd.Marshal(b, m, deterministic)
}
func (m *EigrpAddrBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EigrpAddrBd.Merge(m, src)
}
func (m *EigrpAddrBd) XXX_Size() int {
	return xxx_messageInfo_EigrpAddrBd.Size(m)
}
func (m *EigrpAddrBd) XXX_DiscardUnknown() {
	xxx_messageInfo_EigrpAddrBd.DiscardUnknown(m)
}

var xxx_messageInfo_EigrpAddrBd proto.InternalMessageInfo

func (m *EigrpAddrBd) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *EigrpAddrBd) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type EigrpShowInterfacesBd struct {
	Afi                       uint32         `protobuf:"varint,50,opt,name=afi,proto3" json:"afi,omitempty"`
	AsNumber                  uint32         `protobuf:"varint,51,opt,name=as_number,json=asNumber,proto3" json:"as_number,omitempty"`
	PeerCount                 uint32         `protobuf:"varint,52,opt,name=peer_count,json=peerCount,proto3" json:"peer_count,omitempty"`
	ClassicPeerCount          uint32         `protobuf:"varint,53,opt,name=classic_peer_count,json=classicPeerCount,proto3" json:"classic_peer_count,omitempty"`
	WidePeerCount             uint32         `protobuf:"varint,54,opt,name=wide_peer_count,json=widePeerCount,proto3" json:"wide_peer_count,omitempty"`
	UnreliableTransmits       uint32         `protobuf:"varint,55,opt,name=unreliable_transmits,json=unreliableTransmits,proto3" json:"unreliable_transmits,omitempty"`
	ReliableTransmits         uint32         `protobuf:"varint,56,opt,name=reliable_transmits,json=reliableTransmits,proto3" json:"reliable_transmits,omitempty"`
	TotalSrtt                 uint32         `protobuf:"varint,57,opt,name=total_srtt,json=totalSrtt,proto3" json:"total_srtt,omitempty"`
	UnreliableSendInterval    uint32         `protobuf:"varint,58,opt,name=unreliable_send_interval,json=unreliableSendInterval,proto3" json:"unreliable_send_interval,omitempty"`
	ReliableSendInterval      uint32         `protobuf:"varint,59,opt,name=reliable_send_interval,json=reliableSendInterval,proto3" json:"reliable_send_interval,omitempty"`
	LastMcFlowDelay           uint32         `protobuf:"varint,60,opt,name=last_mc_flow_delay,json=lastMcFlowDelay,proto3" json:"last_mc_flow_delay,omitempty"`
	PendingRoutes             uint32         `protobuf:"varint,61,opt,name=pending_routes,json=pendingRoutes,proto3" json:"pending_routes,omitempty"`
	HelloInterval             uint32         `protobuf:"varint,62,opt,name=hello_interval,json=helloInterval,proto3" json:"hello_interval,omitempty"`
	Holdtime                  uint32         `protobuf:"varint,63,opt,name=holdtime,proto3" json:"holdtime,omitempty"`
	BfdEnabled                bool           `protobuf:"varint,64,opt,name=bfd_enabled,json=bfdEnabled,proto3" json:"bfd_enabled,omitempty"`
	BfdInterval               uint32         `protobuf:"varint,65,opt,name=bfd_interval,json=bfdInterval,proto3" json:"bfd_interval,omitempty"`
	BfdMultiplier             uint32         `protobuf:"varint,66,opt,name=bfd_multiplier,json=bfdMultiplier,proto3" json:"bfd_multiplier,omitempty"`
	SerialNumberPresent       bool           `protobuf:"varint,67,opt,name=serial_number_present,json=serialNumberPresent,proto3" json:"serial_number_present,omitempty"`
	TransmitSerialNumber      uint64         `protobuf:"varint,68,opt,name=transmit_serial_number,json=transmitSerialNumber,proto3" json:"transmit_serial_number,omitempty"`
	PacketizePending          bool           `protobuf:"varint,69,opt,name=packetize_pending,json=packetizePending,proto3" json:"packetize_pending,omitempty"`
	UnreliableMulticastSent   uint32         `protobuf:"varint,70,opt,name=unreliable_multicast_sent,json=unreliableMulticastSent,proto3" json:"unreliable_multicast_sent,omitempty"`
	ReliableMulticastSent     uint32         `protobuf:"varint,71,opt,name=reliable_multicast_sent,json=reliableMulticastSent,proto3" json:"reliable_multicast_sent,omitempty"`
	UnreliableUnicastSent     uint32         `protobuf:"varint,72,opt,name=unreliable_unicast_sent,json=unreliableUnicastSent,proto3" json:"unreliable_unicast_sent,omitempty"`
	ReliableUnicastSent       uint32         `protobuf:"varint,73,opt,name=reliable_unicast_sent,json=reliableUnicastSent,proto3" json:"reliable_unicast_sent,omitempty"`
	MulticastExceptionsSent   uint32         `protobuf:"varint,74,opt,name=multicast_exceptions_sent,json=multicastExceptionsSent,proto3" json:"multicast_exceptions_sent,omitempty"`
	CrPacketsSent             uint32         `protobuf:"varint,75,opt,name=cr_packets_sent,json=crPacketsSent,proto3" json:"cr_packets_sent,omitempty"`
	AcksSuppressed            uint32         `protobuf:"varint,76,opt,name=acks_suppressed,json=acksSuppressed,proto3" json:"acks_suppressed,omitempty"`
	RetransmissionsSent       uint32         `protobuf:"varint,77,opt,name=retransmissions_sent,json=retransmissionsSent,proto3" json:"retransmissions_sent,omitempty"`
	OutOfSequenceReceived     uint32         `protobuf:"varint,78,opt,name=out_of_sequence_received,json=outOfSequenceReceived,proto3" json:"out_of_sequence_received,omitempty"`
	StubInterface             bool           `protobuf:"varint,79,opt,name=stub_interface,json=stubInterface,proto3" json:"stub_interface,omitempty"`
	NextHopSelfEnabled        bool           `protobuf:"varint,80,opt,name=next_hop_self_enabled,json=nextHopSelfEnabled,proto3" json:"next_hop_self_enabled,omitempty"`
	SplitHorizonEnabled       bool           `protobuf:"varint,81,opt,name=split_horizon_enabled,json=splitHorizonEnabled,proto3" json:"split_horizon_enabled,omitempty"`
	PassiveInterface          bool           `protobuf:"varint,82,opt,name=passive_interface,json=passiveInterface,proto3" json:"passive_interface,omitempty"`
	BandwidthPercent          uint32         `protobuf:"varint,83,opt,name=bandwidth_percent,json=bandwidthPercent,proto3" json:"bandwidth_percent,omitempty"`
	StaticNeighbor            []*EigrpAddrBd `protobuf:"bytes,84,rep,name=static_neighbor,json=staticNeighbor,proto3" json:"static_neighbor,omitempty"`
	SiteOfOriginType          string         `protobuf:"bytes,85,opt,name=site_of_origin_type,json=siteOfOriginType,proto3" json:"site_of_origin_type,omitempty"`
	SiteOfOrigin              string         `protobuf:"bytes,86,opt,name=site_of_origin,json=siteOfOrigin,proto3" json:"site_of_origin,omitempty"`
	AuthMode                  uint32         `protobuf:"varint,87,opt,name=auth_mode,json=authMode,proto3" json:"auth_mode,omitempty"`
	AuthKeychain              string         `protobuf:"bytes,88,opt,name=auth_keychain,json=authKeychain,proto3" json:"auth_keychain,omitempty"`
	AuthKeyExists             bool           `protobuf:"varint,89,opt,name=auth_key_exists,json=authKeyExists,proto3" json:"auth_key_exists,omitempty"`
	AuthKeyMd5                bool           `protobuf:"varint,90,opt,name=auth_key_md5,json=authKeyMd5,proto3" json:"auth_key_md5,omitempty"`
	AuthKeyId                 uint64         `protobuf:"varint,91,opt,name=auth_key_id,json=authKeyId,proto3" json:"auth_key_id,omitempty"`
	TotalPktRecvd             uint32         `protobuf:"varint,92,opt,name=total_pkt_recvd,json=totalPktRecvd,proto3" json:"total_pkt_recvd,omitempty"`
	PktDropWrongKc            uint32         `protobuf:"varint,93,opt,name=pkt_drop_wrong_kc,json=pktDropWrongKc,proto3" json:"pkt_drop_wrong_kc,omitempty"`
	PktDropNoAuth             uint32         `protobuf:"varint,94,opt,name=pkt_drop_no_auth,json=pktDropNoAuth,proto3" json:"pkt_drop_no_auth,omitempty"`
	PktDropInvalidAuth        uint32         `protobuf:"varint,95,opt,name=pkt_drop_invalid_auth,json=pktDropInvalidAuth,proto3" json:"pkt_drop_invalid_auth,omitempty"`
	PktAcceptedValidAuth      uint32         `protobuf:"varint,96,opt,name=pkt_accepted_valid_auth,json=pktAcceptedValidAuth,proto3" json:"pkt_accepted_valid_auth,omitempty"`
	Bandwidth                 uint32         `protobuf:"varint,97,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	Bandwidth64               uint64         `protobuf:"varint,98,opt,name=bandwidth64,proto3" json:"bandwidth64,omitempty"`
	Delay                     uint32         `protobuf:"varint,99,opt,name=delay,proto3" json:"delay,omitempty"`
	Delay64                   uint64         `protobuf:"varint,100,opt,name=delay64,proto3" json:"delay64,omitempty"`
	DelayUnit                 string         `protobuf:"bytes,101,opt,name=delay_unit,json=delayUnit,proto3" json:"delay_unit,omitempty"`
	Reliability               uint32         `protobuf:"varint,102,opt,name=reliability,proto3" json:"reliability,omitempty"`
	Load                      uint32         `protobuf:"varint,103,opt,name=load,proto3" json:"load,omitempty"`
	Mtu                       uint32         `protobuf:"varint,104,opt,name=mtu,proto3" json:"mtu,omitempty"`
	ConfiguredBandwidth       uint32         `protobuf:"varint,105,opt,name=configured_bandwidth,json=configuredBandwidth,proto3" json:"configured_bandwidth,omitempty"`
	ConfiguredBandwidth64     uint64         `protobuf:"varint,106,opt,name=configured_bandwidth64,json=configuredBandwidth64,proto3" json:"configured_bandwidth64,omitempty"`
	ConfiguredDelay           uint32         `protobuf:"varint,107,opt,name=configured_delay,json=configuredDelay,proto3" json:"configured_delay,omitempty"`
	ConfiguredDelay64         uint64         `protobuf:"varint,108,opt,name=configured_delay64,json=configuredDelay64,proto3" json:"configured_delay64,omitempty"`
	ConfiguredDelayUnit       string         `protobuf:"bytes,109,opt,name=configured_delay_unit,json=configuredDelayUnit,proto3" json:"configured_delay_unit,omitempty"`
	ConfiguredReliability     uint32         `protobuf:"varint,110,opt,name=configured_reliability,json=configuredReliability,proto3" json:"configured_reliability,omitempty"`
	ConfiguredLoad            uint32         `protobuf:"varint,111,opt,name=configured_load,json=configuredLoad,proto3" json:"configured_load,omitempty"`
	ConfiguredBandwidthFlag   bool           `protobuf:"varint,112,opt,name=configured_bandwidth_flag,json=configuredBandwidthFlag,proto3" json:"configured_bandwidth_flag,omitempty"`
	ConfiguredDelayFlag       bool           `protobuf:"varint,113,opt,name=configured_delay_flag,json=configuredDelayFlag,proto3" json:"configured_delay_flag,omitempty"`
	ConfiguredReliabilityFlag bool           `protobuf:"varint,114,opt,name=configured_reliability_flag,json=configuredReliabilityFlag,proto3" json:"configured_reliability_flag,omitempty"`
	ConfiguredLoadFlag        bool           `protobuf:"varint,115,opt,name=configured_load_flag,json=configuredLoadFlag,proto3" json:"configured_load_flag,omitempty"`
	Up                        bool           `protobuf:"varint,116,opt,name=up,proto3" json:"up,omitempty"`
	TypeSupported             bool           `protobuf:"varint,117,opt,name=type_supported,json=typeSupported,proto3" json:"type_supported,omitempty"`
	ItalRecordFound           bool           `protobuf:"varint,118,opt,name=ital_record_found,json=italRecordFound,proto3" json:"ital_record_found,omitempty"`
	Configured                bool           `protobuf:"varint,119,opt,name=configured,proto3" json:"configured,omitempty"`
	MulticastEnabled          bool           `protobuf:"varint,120,opt,name=multicast_enabled,json=multicastEnabled,proto3" json:"multicast_enabled,omitempty"`
	SocketSetup               bool           `protobuf:"varint,121,opt,name=socket_setup,json=socketSetup,proto3" json:"socket_setup,omitempty"`
	LptsSocketSetup           bool           `protobuf:"varint,122,opt,name=lpts_socket_setup,json=lptsSocketSetup,proto3" json:"lpts_socket_setup,omitempty"`
	PrimaryIpv4Address        string         `protobuf:"bytes,123,opt,name=primary_ipv4_address,json=primaryIpv4Address,proto3" json:"primary_ipv4_address,omitempty"`
	Ipv6LinkLocalAddr         string         `protobuf:"bytes,124,opt,name=ipv6_link_local_addr,json=ipv6LinkLocalAddr,proto3" json:"ipv6_link_local_addr,omitempty"`
	PrimaryPrefixLength       uint32         `protobuf:"varint,125,opt,name=primary_prefix_length,json=primaryPrefixLength,proto3" json:"primary_prefix_length,omitempty"`
	InterfaceHandle           uint32         `protobuf:"varint,126,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	InterfaceType             uint32         `protobuf:"varint,127,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	ConfiguredItems           uint32         `protobuf:"varint,128,opt,name=configured_items,json=configuredItems,proto3" json:"configured_items,omitempty"`
	IsPassiveEnabled          bool           `protobuf:"varint,129,opt,name=is_passive_enabled,json=isPassiveEnabled,proto3" json:"is_passive_enabled,omitempty"`
	IsPassiveDisabled         bool           `protobuf:"varint,130,opt,name=is_passive_disabled,json=isPassiveDisabled,proto3" json:"is_passive_disabled,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}       `json:"-"`
	XXX_unrecognized          []byte         `json:"-"`
	XXX_sizecache             int32          `json:"-"`
}

func (m *EigrpShowInterfacesBd) Reset()         { *m = EigrpShowInterfacesBd{} }
func (m *EigrpShowInterfacesBd) String() string { return proto.CompactTextString(m) }
func (*EigrpShowInterfacesBd) ProtoMessage()    {}
func (*EigrpShowInterfacesBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06608898067dd8f, []int{2}
}

func (m *EigrpShowInterfacesBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EigrpShowInterfacesBd.Unmarshal(m, b)
}
func (m *EigrpShowInterfacesBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EigrpShowInterfacesBd.Marshal(b, m, deterministic)
}
func (m *EigrpShowInterfacesBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EigrpShowInterfacesBd.Merge(m, src)
}
func (m *EigrpShowInterfacesBd) XXX_Size() int {
	return xxx_messageInfo_EigrpShowInterfacesBd.Size(m)
}
func (m *EigrpShowInterfacesBd) XXX_DiscardUnknown() {
	xxx_messageInfo_EigrpShowInterfacesBd.DiscardUnknown(m)
}

var xxx_messageInfo_EigrpShowInterfacesBd proto.InternalMessageInfo

func (m *EigrpShowInterfacesBd) GetAfi() uint32 {
	if m != nil {
		return m.Afi
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetAsNumber() uint32 {
	if m != nil {
		return m.AsNumber
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPeerCount() uint32 {
	if m != nil {
		return m.PeerCount
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetClassicPeerCount() uint32 {
	if m != nil {
		return m.ClassicPeerCount
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetWidePeerCount() uint32 {
	if m != nil {
		return m.WidePeerCount
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetUnreliableTransmits() uint32 {
	if m != nil {
		return m.UnreliableTransmits
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetReliableTransmits() uint32 {
	if m != nil {
		return m.ReliableTransmits
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetTotalSrtt() uint32 {
	if m != nil {
		return m.TotalSrtt
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetUnreliableSendInterval() uint32 {
	if m != nil {
		return m.UnreliableSendInterval
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetReliableSendInterval() uint32 {
	if m != nil {
		return m.ReliableSendInterval
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetLastMcFlowDelay() uint32 {
	if m != nil {
		return m.LastMcFlowDelay
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPendingRoutes() uint32 {
	if m != nil {
		return m.PendingRoutes
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetHoldtime() uint32 {
	if m != nil {
		return m.Holdtime
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetBfdEnabled() bool {
	if m != nil {
		return m.BfdEnabled
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetBfdInterval() uint32 {
	if m != nil {
		return m.BfdInterval
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetBfdMultiplier() uint32 {
	if m != nil {
		return m.BfdMultiplier
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetSerialNumberPresent() bool {
	if m != nil {
		return m.SerialNumberPresent
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetTransmitSerialNumber() uint64 {
	if m != nil {
		return m.TransmitSerialNumber
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPacketizePending() bool {
	if m != nil {
		return m.PacketizePending
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetUnreliableMulticastSent() uint32 {
	if m != nil {
		return m.UnreliableMulticastSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetReliableMulticastSent() uint32 {
	if m != nil {
		return m.ReliableMulticastSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetUnreliableUnicastSent() uint32 {
	if m != nil {
		return m.UnreliableUnicastSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetReliableUnicastSent() uint32 {
	if m != nil {
		return m.ReliableUnicastSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetMulticastExceptionsSent() uint32 {
	if m != nil {
		return m.MulticastExceptionsSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetCrPacketsSent() uint32 {
	if m != nil {
		return m.CrPacketsSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetAcksSuppressed() uint32 {
	if m != nil {
		return m.AcksSuppressed
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetRetransmissionsSent() uint32 {
	if m != nil {
		return m.RetransmissionsSent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetOutOfSequenceReceived() uint32 {
	if m != nil {
		return m.OutOfSequenceReceived
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetStubInterface() bool {
	if m != nil {
		return m.StubInterface
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetNextHopSelfEnabled() bool {
	if m != nil {
		return m.NextHopSelfEnabled
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetSplitHorizonEnabled() bool {
	if m != nil {
		return m.SplitHorizonEnabled
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetPassiveInterface() bool {
	if m != nil {
		return m.PassiveInterface
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetBandwidthPercent() uint32 {
	if m != nil {
		return m.BandwidthPercent
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetStaticNeighbor() []*EigrpAddrBd {
	if m != nil {
		return m.StaticNeighbor
	}
	return nil
}

func (m *EigrpShowInterfacesBd) GetSiteOfOriginType() string {
	if m != nil {
		return m.SiteOfOriginType
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetSiteOfOrigin() string {
	if m != nil {
		return m.SiteOfOrigin
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetAuthMode() uint32 {
	if m != nil {
		return m.AuthMode
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetAuthKeychain() string {
	if m != nil {
		return m.AuthKeychain
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetAuthKeyExists() bool {
	if m != nil {
		return m.AuthKeyExists
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetAuthKeyMd5() bool {
	if m != nil {
		return m.AuthKeyMd5
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetAuthKeyId() uint64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetTotalPktRecvd() uint32 {
	if m != nil {
		return m.TotalPktRecvd
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPktDropWrongKc() uint32 {
	if m != nil {
		return m.PktDropWrongKc
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPktDropNoAuth() uint32 {
	if m != nil {
		return m.PktDropNoAuth
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPktDropInvalidAuth() uint32 {
	if m != nil {
		return m.PktDropInvalidAuth
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetPktAcceptedValidAuth() uint32 {
	if m != nil {
		return m.PktAcceptedValidAuth
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetBandwidth64() uint64 {
	if m != nil {
		return m.Bandwidth64
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetDelay() uint32 {
	if m != nil {
		return m.Delay
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetDelay64() uint64 {
	if m != nil {
		return m.Delay64
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetDelayUnit() string {
	if m != nil {
		return m.DelayUnit
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetReliability() uint32 {
	if m != nil {
		return m.Reliability
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetLoad() uint32 {
	if m != nil {
		return m.Load
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredBandwidth() uint32 {
	if m != nil {
		return m.ConfiguredBandwidth
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredBandwidth64() uint64 {
	if m != nil {
		return m.ConfiguredBandwidth64
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredDelay() uint32 {
	if m != nil {
		return m.ConfiguredDelay
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredDelay64() uint64 {
	if m != nil {
		return m.ConfiguredDelay64
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredDelayUnit() string {
	if m != nil {
		return m.ConfiguredDelayUnit
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetConfiguredReliability() uint32 {
	if m != nil {
		return m.ConfiguredReliability
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredLoad() uint32 {
	if m != nil {
		return m.ConfiguredLoad
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredBandwidthFlag() bool {
	if m != nil {
		return m.ConfiguredBandwidthFlag
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetConfiguredDelayFlag() bool {
	if m != nil {
		return m.ConfiguredDelayFlag
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetConfiguredReliabilityFlag() bool {
	if m != nil {
		return m.ConfiguredReliabilityFlag
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetConfiguredLoadFlag() bool {
	if m != nil {
		return m.ConfiguredLoadFlag
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetUp() bool {
	if m != nil {
		return m.Up
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetTypeSupported() bool {
	if m != nil {
		return m.TypeSupported
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetItalRecordFound() bool {
	if m != nil {
		return m.ItalRecordFound
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetConfigured() bool {
	if m != nil {
		return m.Configured
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetMulticastEnabled() bool {
	if m != nil {
		return m.MulticastEnabled
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetSocketSetup() bool {
	if m != nil {
		return m.SocketSetup
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetLptsSocketSetup() bool {
	if m != nil {
		return m.LptsSocketSetup
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetPrimaryIpv4Address() string {
	if m != nil {
		return m.PrimaryIpv4Address
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetIpv6LinkLocalAddr() string {
	if m != nil {
		return m.Ipv6LinkLocalAddr
	}
	return ""
}

func (m *EigrpShowInterfacesBd) GetPrimaryPrefixLength() uint32 {
	if m != nil {
		return m.PrimaryPrefixLength
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetInterfaceHandle() uint32 {
	if m != nil {
		return m.InterfaceHandle
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetInterfaceType() uint32 {
	if m != nil {
		return m.InterfaceType
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetConfiguredItems() uint32 {
	if m != nil {
		return m.ConfiguredItems
	}
	return 0
}

func (m *EigrpShowInterfacesBd) GetIsPassiveEnabled() bool {
	if m != nil {
		return m.IsPassiveEnabled
	}
	return false
}

func (m *EigrpShowInterfacesBd) GetIsPassiveDisabled() bool {
	if m != nil {
		return m.IsPassiveDisabled
	}
	return false
}

func init() {
	proto.RegisterType((*EigrpShowInterfacesBd_KEYS)(nil), "cisco_ios_xr_eigrp_oper.eigrp.processes.process.vrfs.vrf.afs.af.ases.as.interfaces.interface.eigrp_show_interfaces_bd_KEYS")
	proto.RegisterType((*EigrpAddrBd)(nil), "cisco_ios_xr_eigrp_oper.eigrp.processes.process.vrfs.vrf.afs.af.ases.as.interfaces.interface.eigrp_addr_bd")
	proto.RegisterType((*EigrpShowInterfacesBd)(nil), "cisco_ios_xr_eigrp_oper.eigrp.processes.process.vrfs.vrf.afs.af.ases.as.interfaces.interface.eigrp_show_interfaces_bd")
}

func init() { proto.RegisterFile("eigrp_show_interfaces_bd.proto", fileDescriptor_c06608898067dd8f) }

var fileDescriptor_c06608898067dd8f = []byte{
	// 1699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdb, 0x7b, 0x1b, 0x47,
	0x15, 0xff, 0x9c, 0xa4, 0x4d, 0x3c, 0x8e, 0x6f, 0x63, 0x3b, 0x1e, 0xd3, 0x26, 0x98, 0x40, 0x1b,
	0xb7, 0x21, 0x6a, 0x9b, 0xca, 0x6a, 0x09, 0x50, 0x48, 0x1b, 0x87, 0x98, 0xd8, 0x8e, 0x90, 0x92,
	0x96, 0x42, 0x61, 0x18, 0xed, 0xcc, 0x4a, 0x83, 0x56, 0x3b, 0xdb, 0x99, 0x59, 0xd9, 0x0e, 0x77,
	0x9e, 0x79, 0xe4, 0xbf, 0xe0, 0x9f, 0xe4, 0x3b, 0x67, 0xf6, 0x26, 0x23, 0x3f, 0xf2, 0xa2, 0x6f,
	0xf5, 0xbb, 0xec, 0x9c, 0x39, 0x7b, 0xe6, 0x9c, 0x5d, 0x72, 0x47, 0xe9, 0xa1, 0xcd, 0xb8, 0x1b,
	0x99, 0x53, 0xae, 0x53, 0xaf, 0x6c, 0x2c, 0x22, 0xe5, 0xf8, 0x40, 0xb6, 0x32, 0x6b, 0xbc, 0xa1,
	0xdf, 0x44, 0xda, 0x45, 0x86, 0x6b, 0xe3, 0xf8, 0x99, 0xe5, 0x41, 0x6c, 0x32, 0x65, 0x5b, 0x78,
	0x09, 0xa2, 0x48, 0x39, 0xa7, 0x5c, 0x79, 0xd5, 0x9a, 0xda, 0x18, 0x7f, 0x5a, 0x22, 0x76, 0x2d,
	0x11, 0xb7, 0x04, 0x90, 0xc2, 0xb5, 0xea, 0x9b, 0xd7, 0x97, 0x77, 0xff, 0xb3, 0x40, 0x6e, 0x5f,
	0x16, 0x00, 0x7f, 0x7e, 0xf0, 0x75, 0x9f, 0xde, 0x26, 0xa4, 0xb8, 0x33, 0xd7, 0x92, 0x2d, 0xec,
	0x2e, 0xec, 0x2d, 0xf6, 0x16, 0x0b, 0xe4, 0x50, 0xd2, 0x1d, 0x72, 0x63, 0x6a, 0x63, 0x9e, 0x8a,
	0x89, 0x62, 0x57, 0x90, 0xbc, 0x3e, 0xb5, 0xf1, 0x89, 0x98, 0x28, 0xba, 0x4d, 0xae, 0x8b, 0x82,
	0xb9, 0x8a, 0xcc, 0x9b, 0x22, 0x10, 0x6b, 0xe4, 0xaa, 0x70, 0x29, 0xbb, 0xb6, 0xbb, 0xb0, 0xb7,
	0xdc, 0x83, 0x4b, 0xfa, 0x0e, 0x59, 0xa9, 0x96, 0x0e, 0x8e, 0x37, 0xd0, 0xb1, 0x5c, 0xa1, 0x60,
	0xbc, 0xfb, 0x8a, 0x2c, 0x87, 0x60, 0x85, 0x94, 0x96, 0x0f, 0x24, 0xfd, 0x1e, 0xb9, 0xa9, 0xb3,
	0x69, 0x1b, 0xff, 0x2b, 0xe7, 0x8a, 0xf0, 0x96, 0x00, 0x7b, 0x1c, 0xa0, 0x42, 0xd2, 0xa9, 0x24,
	0x57, 0x2a, 0x49, 0xa7, 0x90, 0xdc, 0xfd, 0xd7, 0xdb, 0x84, 0x5d, 0x96, 0x04, 0x0c, 0x36, 0xd6,
	0xec, 0x61, 0x11, 0x6c, 0xac, 0xe9, 0x5b, 0x64, 0x51, 0x38, 0x9e, 0xe6, 0x93, 0x81, 0xb2, 0xec,
	0x63, 0xc4, 0x6f, 0x08, 0x77, 0x82, 0xff, 0x31, 0x5d, 0x4a, 0x59, 0x1e, 0x99, 0x3c, 0xf5, 0xac,
	0x8d, 0xec, 0x22, 0x20, 0x5f, 0x00, 0x40, 0x7f, 0x48, 0x68, 0x94, 0x08, 0xe7, 0x74, 0xc4, 0x1b,
	0xb2, 0x7d, 0x94, 0xad, 0x15, 0x4c, 0xb7, 0x52, 0xbf, 0x4b, 0x56, 0x4f, 0xb5, 0x54, 0x4d, 0x69,
	0x07, 0xa5, 0xcb, 0x00, 0xd7, 0xba, 0x8f, 0xc8, 0x66, 0x9e, 0x5a, 0x95, 0x68, 0x31, 0x48, 0x14,
	0xf7, 0x56, 0xa4, 0x6e, 0xa2, 0xbd, 0x63, 0x9f, 0xa0, 0x78, 0xa3, 0xe6, 0x5e, 0x96, 0x14, 0x7d,
	0x40, 0xe8, 0x1c, 0xc3, 0xa7, 0x68, 0x58, 0xff, 0x5f, 0xf9, 0x6d, 0x42, 0xbc, 0xf1, 0x22, 0xe1,
	0xce, 0x7a, 0xcf, 0x7e, 0x14, 0xb6, 0x85, 0x48, 0xdf, 0x7a, 0x4f, 0x3f, 0x25, 0xac, 0x11, 0x80,
	0x53, 0xa9, 0x0c, 0x59, 0x9c, 0x8a, 0x84, 0x3d, 0x42, 0xf1, 0xad, 0x9a, 0xef, 0xab, 0x54, 0x1e,
	0x16, 0x2c, 0x6d, 0x93, 0x5b, 0x97, 0xf8, 0x7e, 0x8c, 0xbe, 0xcd, 0xb9, 0xae, 0xfb, 0x84, 0x26,
	0xc2, 0x79, 0x3e, 0x89, 0x78, 0x9c, 0x98, 0x53, 0x2e, 0x55, 0x22, 0xce, 0xd9, 0x4f, 0xd0, 0xb1,
	0x0a, 0xcc, 0x71, 0xf4, 0x34, 0x31, 0xa7, 0x4f, 0x00, 0x86, 0xe2, 0xca, 0x54, 0x2a, 0x75, 0x3a,
	0xe4, 0xd6, 0xe4, 0x5e, 0x39, 0xf6, 0xd3, 0x90, 0xc4, 0x02, 0xed, 0x21, 0x08, 0xb2, 0x91, 0x4a,
	0x12, 0x53, 0x47, 0xf0, 0x59, 0x90, 0x21, 0x5a, 0x2d, 0xfd, 0x1d, 0x72, 0x63, 0x64, 0x12, 0xe9,
	0xf5, 0x44, 0xb1, 0x9f, 0x85, 0x87, 0x5f, 0xfe, 0xa7, 0xdf, 0x25, 0x4b, 0x83, 0x58, 0x72, 0x95,
	0x42, 0xc0, 0x92, 0xfd, 0x7c, 0x77, 0x61, 0xef, 0x46, 0x8f, 0x0c, 0x62, 0x79, 0x10, 0x10, 0x28,
	0x46, 0x10, 0x54, 0x2b, 0x3c, 0xc6, 0x1b, 0x80, 0xa9, 0xba, 0xff, 0x3b, 0x64, 0x05, 0x24, 0x93,
	0x3c, 0xf1, 0x3a, 0x4b, 0xb4, 0xb2, 0xec, 0xf3, 0x10, 0xc6, 0x20, 0x96, 0xc7, 0x15, 0x48, 0x1f,
	0x92, 0x2d, 0xa7, 0xac, 0x16, 0x49, 0x51, 0x88, 0x3c, 0xb3, 0xca, 0xa9, 0xd4, 0xb3, 0x2f, 0x70,
	0xd1, 0x8d, 0x40, 0x86, 0xa2, 0xec, 0x06, 0x0a, 0x72, 0x5d, 0x3e, 0x6a, 0x3e, 0x63, 0x66, 0x4f,
	0x76, 0x17, 0xf6, 0xae, 0xf5, 0x36, 0x4b, 0xb6, 0xdf, 0x30, 0xd3, 0xfb, 0x64, 0x3d, 0x13, 0xd1,
	0x58, 0x79, 0xfd, 0x1a, 0x2a, 0x11, 0x53, 0xc6, 0x0e, 0x70, 0x95, 0xb5, 0x8a, 0xe8, 0x06, 0x9c,
	0x3e, 0x22, 0x3b, 0x8d, 0x42, 0xc0, 0x4d, 0x44, 0xf0, 0xa0, 0x30, 0xb4, 0xa7, 0xb8, 0x91, 0xed,
	0x5a, 0x70, 0x5c, 0xf2, 0x7d, 0x08, 0xaf, 0x43, 0xb6, 0x2f, 0x73, 0xfe, 0x02, 0x9d, 0x5b, 0x97,
	0xfa, 0x1a, 0x6b, 0xe6, 0x69, 0xc3, 0xf7, 0x2c, 0xf8, 0x6a, 0xfa, 0x55, 0x5a, 0xfb, 0x1e, 0x92,
	0xad, 0xf9, 0xae, 0xc3, 0x70, 0x6c, 0xe6, 0x79, 0x1e, 0x91, 0x9d, 0x3a, 0x34, 0x75, 0x16, 0xa9,
	0xcc, 0x6b, 0x93, 0xba, 0xe0, 0xfb, 0x65, 0xd8, 0x5f, 0x25, 0x38, 0xa8, 0x78, 0xf4, 0xbe, 0x4b,
	0x56, 0x23, 0xcb, 0x43, 0xca, 0x0a, 0xc7, 0xf3, 0xf0, 0x68, 0x23, 0xdb, 0x0d, 0x28, 0xea, 0xee,
	0x91, 0x55, 0x11, 0x8d, 0x1d, 0x77, 0x79, 0x06, 0x0f, 0xd5, 0x29, 0xc9, 0x8e, 0x50, 0xb7, 0x02,
	0x70, 0xbf, 0x42, 0xe1, 0xd8, 0x5b, 0x55, 0x3c, 0x33, 0xe7, 0xea, 0x38, 0x8e, 0xcb, 0xf8, 0x67,
	0x38, 0xbc, 0xf7, 0x27, 0x84, 0x99, 0xdc, 0x73, 0x13, 0x73, 0xa7, 0xbe, 0xcd, 0x55, 0x1a, 0x29,
	0x6e, 0x55, 0xa4, 0xf4, 0x54, 0x49, 0x76, 0x12, 0x92, 0x65, 0x72, 0xff, 0x22, 0xee, 0x17, 0x6c,
	0xaf, 0x20, 0xa1, 0x2c, 0x9d, 0xcf, 0x07, 0x75, 0x73, 0x64, 0x2f, 0xb0, 0x04, 0x96, 0x01, 0x3d,
	0x2c, 0x41, 0xfa, 0x11, 0xd9, 0x4a, 0xd5, 0x99, 0xe7, 0x23, 0x93, 0x71, 0xa7, 0x92, 0xb8, 0x3a,
	0x0b, 0x5d, 0x54, 0x53, 0x20, 0x9f, 0x99, 0xac, 0xaf, 0x92, 0xb8, 0x3c, 0x13, 0x50, 0xc9, 0x59,
	0xa2, 0xc1, 0x63, 0xf5, 0x6b, 0x93, 0x56, 0x96, 0x5f, 0x15, 0x95, 0x0c, 0xe4, 0xb3, 0xc0, 0x95,
	0x1e, 0xac, 0x49, 0xe7, 0xf4, 0x54, 0x35, 0x02, 0xea, 0x95, 0x35, 0x89, 0x44, 0x1d, 0xd3, 0x7d,
	0xb2, 0x3e, 0x10, 0xa9, 0x3c, 0xd5, 0xd2, 0x8f, 0x78, 0xa6, 0x6c, 0x04, 0x39, 0xea, 0x87, 0x96,
	0x5b, 0x11, 0xdd, 0x80, 0xd3, 0x7f, 0x2f, 0x90, 0x55, 0xe7, 0x85, 0xd7, 0x11, 0x4f, 0x95, 0x1e,
	0x8e, 0x06, 0xc6, 0xb2, 0x97, 0xbb, 0x57, 0xf7, 0x96, 0x1e, 0x8e, 0x5b, 0xff, 0xcf, 0x49, 0xdc,
	0x9a, 0x19, 0x6c, 0xbd, 0x95, 0x10, 0xc3, 0x49, 0x11, 0x02, 0x7d, 0x40, 0x36, 0x9c, 0xf6, 0x0a,
	0x1e, 0x9c, 0xb1, 0x7a, 0xa8, 0x53, 0xee, 0xcf, 0x33, 0xc5, 0x5e, 0xe1, 0x30, 0x5b, 0x03, 0xea,
	0x45, 0xfc, 0x02, 0x89, 0x97, 0xe7, 0x99, 0xa2, 0x3f, 0x20, 0x2b, 0xb3, 0x72, 0xf6, 0x25, 0x2a,
	0x6f, 0x36, 0x95, 0x38, 0xc8, 0x72, 0x3f, 0xe2, 0x13, 0x23, 0x15, 0xfb, 0xaa, 0x18, 0x64, 0xb9,
	0x1f, 0x1d, 0x1b, 0xa9, 0xe8, 0xf7, 0xc9, 0x32, 0x92, 0x63, 0x75, 0x1e, 0x8d, 0x84, 0x4e, 0xd9,
	0xaf, 0xc3, 0x1d, 0x00, 0x7c, 0x5e, 0x60, 0x50, 0xd2, 0xa5, 0x88, 0xab, 0x33, 0xed, 0xbc, 0x63,
	0x5f, 0x87, 0xb2, 0x28, 0x64, 0x07, 0x08, 0xd2, 0x5d, 0x72, 0xb3, 0xd2, 0x4d, 0xe4, 0x3e, 0xfb,
	0x4d, 0xe8, 0x8c, 0x85, 0xe8, 0x58, 0xee, 0xd3, 0x3b, 0x64, 0xa9, 0x52, 0x68, 0xc9, 0x7e, 0x8b,
	0x0d, 0x69, 0xb1, 0x10, 0x1c, 0x4a, 0x58, 0x29, 0x0c, 0xa0, 0x6c, 0xec, 0xa1, 0x64, 0xa7, 0x92,
	0x7d, 0x13, 0x0e, 0x0f, 0xc2, 0xdd, 0xb1, 0xef, 0x01, 0x48, 0xdf, 0x23, 0xeb, 0xa0, 0x90, 0xd6,
	0x64, 0xfc, 0xd4, 0x9a, 0x74, 0xc8, 0xc7, 0x11, 0xfb, 0x5d, 0x38, 0x3e, 0xd9, 0xd8, 0x3f, 0xb1,
	0x26, 0xfb, 0x0a, 0xe0, 0xe7, 0x11, 0xbd, 0x47, 0xd6, 0x2a, 0x69, 0x6a, 0x38, 0xac, 0xc5, 0x7e,
	0x5f, 0x4c, 0x86, 0xa0, 0x3c, 0x31, 0x8f, 0x73, 0x3f, 0x82, 0xa2, 0xae, 0x84, 0x3a, 0x9d, 0x8a,
	0x44, 0xcb, 0xa0, 0xe6, 0xa8, 0xa6, 0x85, 0xfa, 0x30, 0x50, 0x68, 0xd9, 0x27, 0xdb, 0x60, 0x11,
	0x11, 0x74, 0x00, 0x25, 0x79, 0xc3, 0xf4, 0x87, 0x30, 0xd7, 0xb2, 0xb1, 0x7f, 0x5c, 0xb0, 0x5f,
	0x56, 0xb6, 0xb7, 0xc9, 0x62, 0x55, 0x91, 0x4c, 0x84, 0x29, 0x5b, 0x01, 0x74, 0x97, 0x2c, 0x55,
	0x7f, 0x3a, 0x6d, 0x36, 0xc0, 0x1c, 0x35, 0x21, 0xba, 0x49, 0xde, 0x08, 0xa3, 0x30, 0x42, 0x6f,
	0xf8, 0x43, 0x19, 0xb9, 0x8e, 0x17, 0x9d, 0x36, 0x93, 0xe8, 0x29, 0xff, 0xc2, 0x58, 0xc7, 0x4b,
	0xe8, 0x7f, 0x9e, 0xa9, 0xf0, 0x72, 0x87, 0xc8, 0xab, 0x54, 0x7b, 0x58, 0x30, 0x34, 0x41, 0x9d,
	0x68, 0x7f, 0xce, 0xe2, 0x30, 0xad, 0x1a, 0x10, 0xa5, 0xe4, 0x5a, 0x62, 0x84, 0x64, 0x43, 0xa4,
	0xf0, 0x1a, 0xde, 0x98, 0x26, 0x3e, 0x67, 0xa3, 0xf0, 0xc6, 0x34, 0xf1, 0x39, 0x34, 0xaa, 0xc8,
	0xa4, 0xb1, 0x1e, 0xe6, 0x56, 0x49, 0x5e, 0xef, 0x50, 0x87, 0x46, 0x55, 0x73, 0x9f, 0x57, 0x7b,
	0xdd, 0x27, 0xb7, 0xe6, 0x59, 0x3a, 0x6d, 0xf6, 0x47, 0xdc, 0xc2, 0xd6, 0x1c, 0x53, 0xa7, 0x4d,
	0xdf, 0x23, 0x6b, 0x0d, 0x5b, 0xc8, 0xc5, 0x38, 0xbc, 0x16, 0xd4, 0x78, 0x78, 0x2d, 0x78, 0x40,
	0xe8, 0x45, 0x69, 0xa7, 0xcd, 0x12, 0xbc, 0xfb, 0xfa, 0x05, 0x71, 0xa7, 0x0d, 0x6d, 0xea, 0xa2,
	0x3c, 0x64, 0x6d, 0x82, 0x59, 0xdb, 0xb8, 0xe0, 0xc0, 0xfc, 0xcd, 0x6e, 0xa2, 0x99, 0xca, 0x34,
	0xf4, 0xda, 0x9a, 0xed, 0x35, 0x92, 0x7a, 0x8f, 0x34, 0x82, 0xe5, 0x98, 0x5f, 0x13, 0x2a, 0xb8,
	0x86, 0x8f, 0x20, 0xd3, 0x8f, 0xc8, 0xce, 0xbc, 0x24, 0xf1, 0x38, 0x11, 0x43, 0x96, 0xe1, 0x19,
	0xdb, 0x9e, 0x93, 0xa7, 0xa7, 0x89, 0x18, 0xce, 0xdd, 0x0f, 0xfa, 0xbe, 0x0d, 0x6d, 0xf7, 0xc2,
	0x7e, 0xd0, 0xf3, 0x19, 0x79, 0x6b, 0xfe, 0x7e, 0x82, 0xd3, 0xa2, 0x73, 0x67, 0xee, 0xa6, 0xd0,
	0xff, 0xe1, 0x4c, 0x1d, 0xc0, 0xc6, 0x82, 0xd1, 0x85, 0xe1, 0x30, 0xbb, 0x3b, 0x74, 0xac, 0x90,
	0x2b, 0x79, 0xc6, 0x3c, 0xf2, 0x57, 0xf2, 0x0c, 0xc6, 0x10, 0x34, 0x3e, 0x9c, 0x8d, 0xc6, 0x7a,
	0x25, 0x59, 0x1e, 0xfa, 0x0d, 0xa0, 0xfd, 0x12, 0xa4, 0xef, 0x93, 0x75, 0x0d, 0xcd, 0xc2, 0xaa,
	0xc8, 0x58, 0xc9, 0x63, 0x93, 0xa7, 0x92, 0x4d, 0x51, 0xb9, 0x0a, 0x44, 0x0f, 0xf1, 0xa7, 0x00,
	0xd3, 0x3b, 0x84, 0xd4, 0x0b, 0xb3, 0xd3, 0xd0, 0x99, 0x6a, 0x04, 0xc6, 0x47, 0x63, 0xe4, 0x17,
	0xb3, 0xe9, 0x2c, 0xcc, 0x9a, 0x7a, 0xd4, 0xd7, 0x2f, 0x78, 0xce, 0xc0, 0x28, 0xe7, 0x4e, 0xf9,
	0x3c, 0x63, 0xe7, 0xa8, 0x5b, 0x0a, 0x58, 0x1f, 0x20, 0x88, 0x2d, 0xc9, 0xe0, 0x05, 0xa0, 0xa9,
	0x7b, 0x1d, 0x62, 0x03, 0xa2, 0xdf, 0xd0, 0x7e, 0x48, 0x36, 0x33, 0xab, 0x27, 0xc2, 0x9e, 0xf3,
	0x99, 0xef, 0x9c, 0x3f, 0x61, 0xcd, 0xd1, 0x82, 0x3b, 0x6c, 0x7c, 0xee, 0x7c, 0x40, 0x36, 0xf1,
	0x73, 0x27, 0xd1, 0xe9, 0x98, 0x27, 0x26, 0x12, 0x09, 0x9a, 0xd8, 0x9f, 0xd1, 0xb1, 0x0e, 0xdc,
	0x91, 0x4e, 0xc7, 0x47, 0xc0, 0x80, 0x07, 0xea, 0xa0, 0x5c, 0x22, 0xb3, 0x2a, 0xd6, 0x67, 0x3c,
	0x51, 0xe9, 0xd0, 0x8f, 0xd8, 0x5f, 0xc2, 0xe1, 0x2c, 0xc8, 0x2e, 0x72, 0x47, 0x48, 0xc1, 0x29,
	0xab, 0x3f, 0xd7, 0x46, 0x22, 0x95, 0x89, 0x62, 0x7f, 0x0d, 0xa7, 0xac, 0xc2, 0x9f, 0x21, 0x3c,
	0xfb, 0x65, 0x87, 0x33, 0xeb, 0x6f, 0xa1, 0xc5, 0x56, 0x28, 0x0e, 0xac, 0xf7, 0x67, 0xce, 0xad,
	0xf6, 0x6a, 0xe2, 0xd8, 0xdf, 0x17, 0x2e, 0x1e, 0xdc, 0x43, 0xc0, 0xe1, 0xe0, 0x6a, 0xc7, 0xcb,
	0xf9, 0x5f, 0x3e, 0x91, 0x7f, 0x2c, 0x84, 0x47, 0xa2, 0x5d, 0x37, 0x30, 0xe5, 0x23, 0xf9, 0x80,
	0x6c, 0x34, 0xe4, 0x52, 0xbb, 0xa0, 0xff, 0x67, 0xd0, 0xaf, 0x57, 0xfa, 0x27, 0x05, 0x33, 0x78,
	0x13, 0x3f, 0xbc, 0x3f, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x42, 0x46, 0xf6, 0x9a,
	0x0f, 0x00, 0x00,
}