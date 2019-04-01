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
// source: ipv4_acl_edm_ace.proto

package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_access_list_manager_accesses_access_access_list_sequences_access_list_sequence

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

type Ipv4AclEdmAce_KEYS struct {
	AccessListName       string   `protobuf:"bytes,1,opt,name=access_list_name,json=accessListName,proto3" json:"access_list_name,omitempty"`
	SequenceNumber       uint32   `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclEdmAce_KEYS) Reset()         { *m = Ipv4AclEdmAce_KEYS{} }
func (m *Ipv4AclEdmAce_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmAce_KEYS) ProtoMessage()    {}
func (*Ipv4AclEdmAce_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3b94617fae956f, []int{0}
}

func (m *Ipv4AclEdmAce_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Unmarshal(m, b)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmAce_KEYS.Merge(m, src)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmAce_KEYS.Size(m)
}
func (m *Ipv4AclEdmAce_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmAce_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmAce_KEYS proto.InternalMessageInfo

func (m *Ipv4AclEdmAce_KEYS) GetAccessListName() string {
	if m != nil {
		return m.AccessListName
	}
	return ""
}

func (m *Ipv4AclEdmAce_KEYS) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type Ipv4AclBagNhInfo struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	TrackName            string   `protobuf:"bytes,2,opt,name=track_name,json=trackName,proto3" json:"track_name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	AtStatus             string   `protobuf:"bytes,4,opt,name=at_status,json=atStatus,proto3" json:"at_status,omitempty"`
	IsAclNextHopExist    bool     `protobuf:"varint,5,opt,name=is_acl_next_hop_exist,json=isAclNextHopExist,proto3" json:"is_acl_next_hop_exist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclBagNhInfo) Reset()         { *m = Ipv4AclBagNhInfo{} }
func (m *Ipv4AclBagNhInfo) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclBagNhInfo) ProtoMessage()    {}
func (*Ipv4AclBagNhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3b94617fae956f, []int{1}
}

func (m *Ipv4AclBagNhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Unmarshal(m, b)
}
func (m *Ipv4AclBagNhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Marshal(b, m, deterministic)
}
func (m *Ipv4AclBagNhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclBagNhInfo.Merge(m, src)
}
func (m *Ipv4AclBagNhInfo) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclBagNhInfo.Size(m)
}
func (m *Ipv4AclBagNhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclBagNhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclBagNhInfo proto.InternalMessageInfo

func (m *Ipv4AclBagNhInfo) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetTrackName() string {
	if m != nil {
		return m.TrackName
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetAtStatus() string {
	if m != nil {
		return m.AtStatus
	}
	return ""
}

func (m *Ipv4AclBagNhInfo) GetIsAclNextHopExist() bool {
	if m != nil {
		return m.IsAclNextHopExist
	}
	return false
}

type Ipv4AclBagHwNhInfo struct {
	NextHop              uint32   `protobuf:"varint,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4AclBagHwNhInfo) Reset()         { *m = Ipv4AclBagHwNhInfo{} }
func (m *Ipv4AclBagHwNhInfo) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclBagHwNhInfo) ProtoMessage()    {}
func (*Ipv4AclBagHwNhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3b94617fae956f, []int{2}
}

func (m *Ipv4AclBagHwNhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Unmarshal(m, b)
}
func (m *Ipv4AclBagHwNhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Marshal(b, m, deterministic)
}
func (m *Ipv4AclBagHwNhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclBagHwNhInfo.Merge(m, src)
}
func (m *Ipv4AclBagHwNhInfo) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclBagHwNhInfo.Size(m)
}
func (m *Ipv4AclBagHwNhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclBagHwNhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclBagHwNhInfo proto.InternalMessageInfo

func (m *Ipv4AclBagHwNhInfo) GetNextHop() uint32 {
	if m != nil {
		return m.NextHop
	}
	return 0
}

func (m *Ipv4AclBagHwNhInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ipv4AclBagHwNhInfo) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type AclUdfAce struct {
	UdfName              []string `protobuf:"bytes,1,rep,name=udf_name,json=udfName,proto3" json:"udf_name,omitempty"`
	UdfValue             uint32   `protobuf:"varint,2,opt,name=udf_value,json=udfValue,proto3" json:"udf_value,omitempty"`
	UdfMask              uint32   `protobuf:"varint,3,opt,name=udf_mask,json=udfMask,proto3" json:"udf_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AclUdfAce) Reset()         { *m = AclUdfAce{} }
func (m *AclUdfAce) String() string { return proto.CompactTextString(m) }
func (*AclUdfAce) ProtoMessage()    {}
func (*AclUdfAce) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3b94617fae956f, []int{3}
}

func (m *AclUdfAce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclUdfAce.Unmarshal(m, b)
}
func (m *AclUdfAce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclUdfAce.Marshal(b, m, deterministic)
}
func (m *AclUdfAce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclUdfAce.Merge(m, src)
}
func (m *AclUdfAce) XXX_Size() int {
	return xxx_messageInfo_AclUdfAce.Size(m)
}
func (m *AclUdfAce) XXX_DiscardUnknown() {
	xxx_messageInfo_AclUdfAce.DiscardUnknown(m)
}

var xxx_messageInfo_AclUdfAce proto.InternalMessageInfo

func (m *AclUdfAce) GetUdfName() []string {
	if m != nil {
		return m.UdfName
	}
	return nil
}

func (m *AclUdfAce) GetUdfValue() uint32 {
	if m != nil {
		return m.UdfValue
	}
	return 0
}

func (m *AclUdfAce) GetUdfMask() uint32 {
	if m != nil {
		return m.UdfMask
	}
	return 0
}

type Ipv4AclEdmAce struct {
	ItemType               string              `protobuf:"bytes,50,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"`
	Sequence               uint32              `protobuf:"varint,51,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Grant                  string              `protobuf:"bytes,52,opt,name=grant,proto3" json:"grant,omitempty"`
	ProtocolOperator       uint32              `protobuf:"varint,53,opt,name=protocol_operator,json=protocolOperator,proto3" json:"protocol_operator,omitempty"`
	Protocol               uint32              `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Protocol2              uint32              `protobuf:"varint,55,opt,name=protocol2,proto3" json:"protocol2,omitempty"`
	SourceAddress          string              `protobuf:"bytes,56,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	SourceAddressMask      string              `protobuf:"bytes,57,opt,name=source_address_mask,json=sourceAddressMask,proto3" json:"source_address_mask,omitempty"`
	DestinationAddress     string              `protobuf:"bytes,58,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	DestinationAddressMask string              `protobuf:"bytes,59,opt,name=destination_address_mask,json=destinationAddressMask,proto3" json:"destination_address_mask,omitempty"`
	SourceOperator         string              `protobuf:"bytes,60,opt,name=source_operator,json=sourceOperator,proto3" json:"source_operator,omitempty"`
	SourcePort1            uint32              `protobuf:"varint,61,opt,name=source_port1,json=sourcePort1,proto3" json:"source_port1,omitempty"`
	SourcePort2            uint32              `protobuf:"varint,62,opt,name=source_port2,json=sourcePort2,proto3" json:"source_port2,omitempty"`
	SorceOperator          string              `protobuf:"bytes,63,opt,name=sorce_operator,json=sorceOperator,proto3" json:"sorce_operator,omitempty"`
	SorcePort1             uint32              `protobuf:"varint,64,opt,name=sorce_port1,json=sorcePort1,proto3" json:"sorce_port1,omitempty"`
	SorcePort2             uint32              `protobuf:"varint,65,opt,name=sorce_port2,json=sorcePort2,proto3" json:"sorce_port2,omitempty"`
	DestinationOperator    string              `protobuf:"bytes,66,opt,name=destination_operator,json=destinationOperator,proto3" json:"destination_operator,omitempty"`
	DestinationPort1       uint32              `protobuf:"varint,67,opt,name=destination_port1,json=destinationPort1,proto3" json:"destination_port1,omitempty"`
	DestinationPort2       uint32              `protobuf:"varint,68,opt,name=destination_port2,json=destinationPort2,proto3" json:"destination_port2,omitempty"`
	LogOption              string              `protobuf:"bytes,69,opt,name=log_option,json=logOption,proto3" json:"log_option,omitempty"`
	CounterName            string              `protobuf:"bytes,70,opt,name=counter_name,json=counterName,proto3" json:"counter_name,omitempty"`
	Capture                bool                `protobuf:"varint,71,opt,name=capture,proto3" json:"capture,omitempty"`
	DscpPresent            bool                `protobuf:"varint,72,opt,name=dscp_present,json=dscpPresent,proto3" json:"dscp_present,omitempty"`
	Dscp                   uint32              `protobuf:"varint,73,opt,name=dscp,proto3" json:"dscp,omitempty"`
	Dscp2                  uint32              `protobuf:"varint,74,opt,name=dscp2,proto3" json:"dscp2,omitempty"`
	DscpOperator           uint32              `protobuf:"varint,75,opt,name=dscp_operator,json=dscpOperator,proto3" json:"dscp_operator,omitempty"`
	PrecedencePresent      bool                `protobuf:"varint,76,opt,name=precedence_present,json=precedencePresent,proto3" json:"precedence_present,omitempty"`
	Precedence             uint32              `protobuf:"varint,77,opt,name=precedence,proto3" json:"precedence,omitempty"`
	TcpFlagsOperator       string              `protobuf:"bytes,78,opt,name=tcp_flags_operator,json=tcpFlagsOperator,proto3" json:"tcp_flags_operator,omitempty"`
	TcpFlags               uint32              `protobuf:"varint,79,opt,name=tcp_flags,json=tcpFlags,proto3" json:"tcp_flags,omitempty"`
	TcpFlagsMask           uint32              `protobuf:"varint,80,opt,name=tcp_flags_mask,json=tcpFlagsMask,proto3" json:"tcp_flags_mask,omitempty"`
	Fragments              uint32              `protobuf:"varint,81,opt,name=fragments,proto3" json:"fragments,omitempty"`
	PacketLengthOperator   string              `protobuf:"bytes,82,opt,name=packet_length_operator,json=packetLengthOperator,proto3" json:"packet_length_operator,omitempty"`
	PacketLength1          uint32              `protobuf:"varint,83,opt,name=packet_length1,json=packetLength1,proto3" json:"packet_length1,omitempty"`
	PacketLength2          uint32              `protobuf:"varint,84,opt,name=packet_length2,json=packetLength2,proto3" json:"packet_length2,omitempty"`
	TtlOperator            string              `protobuf:"bytes,85,opt,name=ttl_operator,json=ttlOperator,proto3" json:"ttl_operator,omitempty"`
	Ttl1                   uint32              `protobuf:"varint,86,opt,name=ttl1,proto3" json:"ttl1,omitempty"`
	Ttl2                   uint32              `protobuf:"varint,87,opt,name=ttl2,proto3" json:"ttl2,omitempty"`
	NoStats                bool                `protobuf:"varint,88,opt,name=no_stats,json=noStats,proto3" json:"no_stats,omitempty"`
	Hits                   uint64              `protobuf:"varint,89,opt,name=hits,proto3" json:"hits,omitempty"`
	IsIcmpOff              bool                `protobuf:"varint,90,opt,name=is_icmp_off,json=isIcmpOff,proto3" json:"is_icmp_off,omitempty"`
	QosGroup               uint32              `protobuf:"varint,91,opt,name=qos_group,json=qosGroup,proto3" json:"qos_group,omitempty"`
	NextHopType            string              `protobuf:"bytes,92,opt,name=next_hop_type,json=nextHopType,proto3" json:"next_hop_type,omitempty"`
	NextHopInfo            []*Ipv4AclBagNhInfo `protobuf:"bytes,93,rep,name=next_hop_info,json=nextHopInfo,proto3" json:"next_hop_info,omitempty"`
	HwNextHopInfo          *Ipv4AclBagHwNhInfo `protobuf:"bytes,94,opt,name=hw_next_hop_info,json=hwNextHopInfo,proto3" json:"hw_next_hop_info,omitempty"`
	Remark                 string              `protobuf:"bytes,95,opt,name=remark,proto3" json:"remark,omitempty"`
	Dynamic                bool                `protobuf:"varint,96,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	SourcePrefixGroup      string              `protobuf:"bytes,97,opt,name=source_prefix_group,json=sourcePrefixGroup,proto3" json:"source_prefix_group,omitempty"`
	DestinationPrefixGroup string              `protobuf:"bytes,98,opt,name=destination_prefix_group,json=destinationPrefixGroup,proto3" json:"destination_prefix_group,omitempty"`
	SourcePortGroup        string              `protobuf:"bytes,99,opt,name=source_port_group,json=sourcePortGroup,proto3" json:"source_port_group,omitempty"`
	DestinationPortGroup   string              `protobuf:"bytes,100,opt,name=destination_port_group,json=destinationPortGroup,proto3" json:"destination_port_group,omitempty"`
	AclName                string              `protobuf:"bytes,101,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	SequenceStr            string              `protobuf:"bytes,102,opt,name=sequence_str,json=sequenceStr,proto3" json:"sequence_str,omitempty"`
	FragmentOffsetOperator string              `protobuf:"bytes,103,opt,name=fragment_offset_operator,json=fragmentOffsetOperator,proto3" json:"fragment_offset_operator,omitempty"`
	FragmentOffset1        uint32              `protobuf:"varint,104,opt,name=fragment_offset1,json=fragmentOffset1,proto3" json:"fragment_offset1,omitempty"`
	FragmentOffset2        uint32              `protobuf:"varint,105,opt,name=fragment_offset2,json=fragmentOffset2,proto3" json:"fragment_offset2,omitempty"`
	Udf                    []*AclUdfAce        `protobuf:"bytes,106,rep,name=udf,proto3" json:"udf,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *Ipv4AclEdmAce) Reset()         { *m = Ipv4AclEdmAce{} }
func (m *Ipv4AclEdmAce) String() string { return proto.CompactTextString(m) }
func (*Ipv4AclEdmAce) ProtoMessage()    {}
func (*Ipv4AclEdmAce) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3b94617fae956f, []int{4}
}

func (m *Ipv4AclEdmAce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4AclEdmAce.Unmarshal(m, b)
}
func (m *Ipv4AclEdmAce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4AclEdmAce.Marshal(b, m, deterministic)
}
func (m *Ipv4AclEdmAce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4AclEdmAce.Merge(m, src)
}
func (m *Ipv4AclEdmAce) XXX_Size() int {
	return xxx_messageInfo_Ipv4AclEdmAce.Size(m)
}
func (m *Ipv4AclEdmAce) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4AclEdmAce.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4AclEdmAce proto.InternalMessageInfo

func (m *Ipv4AclEdmAce) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetGrant() string {
	if m != nil {
		return m.Grant
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetProtocolOperator() uint32 {
	if m != nil {
		return m.ProtocolOperator
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetProtocol2() uint32 {
	if m != nil {
		return m.Protocol2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourceAddressMask() string {
	if m != nil {
		return m.SourceAddressMask
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationAddressMask() string {
	if m != nil {
		return m.DestinationAddressMask
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourceOperator() string {
	if m != nil {
		return m.SourceOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourcePort1() uint32 {
	if m != nil {
		return m.SourcePort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSourcePort2() uint32 {
	if m != nil {
		return m.SourcePort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSorceOperator() string {
	if m != nil {
		return m.SorceOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSorcePort1() uint32 {
	if m != nil {
		return m.SorcePort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetSorcePort2() uint32 {
	if m != nil {
		return m.SorcePort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDestinationOperator() string {
	if m != nil {
		return m.DestinationOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPort1() uint32 {
	if m != nil {
		return m.DestinationPort1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDestinationPort2() uint32 {
	if m != nil {
		return m.DestinationPort2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetLogOption() string {
	if m != nil {
		return m.LogOption
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetCounterName() string {
	if m != nil {
		return m.CounterName
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetCapture() bool {
	if m != nil {
		return m.Capture
	}
	return false
}

func (m *Ipv4AclEdmAce) GetDscpPresent() bool {
	if m != nil {
		return m.DscpPresent
	}
	return false
}

func (m *Ipv4AclEdmAce) GetDscp() uint32 {
	if m != nil {
		return m.Dscp
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDscp2() uint32 {
	if m != nil {
		return m.Dscp2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetDscpOperator() uint32 {
	if m != nil {
		return m.DscpOperator
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPrecedencePresent() bool {
	if m != nil {
		return m.PrecedencePresent
	}
	return false
}

func (m *Ipv4AclEdmAce) GetPrecedence() uint32 {
	if m != nil {
		return m.Precedence
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTcpFlagsOperator() string {
	if m != nil {
		return m.TcpFlagsOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetTcpFlags() uint32 {
	if m != nil {
		return m.TcpFlags
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTcpFlagsMask() uint32 {
	if m != nil {
		return m.TcpFlagsMask
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetFragments() uint32 {
	if m != nil {
		return m.Fragments
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPacketLengthOperator() string {
	if m != nil {
		return m.PacketLengthOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetPacketLength1() uint32 {
	if m != nil {
		return m.PacketLength1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetPacketLength2() uint32 {
	if m != nil {
		return m.PacketLength2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTtlOperator() string {
	if m != nil {
		return m.TtlOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetTtl1() uint32 {
	if m != nil {
		return m.Ttl1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetTtl2() uint32 {
	if m != nil {
		return m.Ttl2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetNoStats() bool {
	if m != nil {
		return m.NoStats
	}
	return false
}

func (m *Ipv4AclEdmAce) GetHits() uint64 {
	if m != nil {
		return m.Hits
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetIsIcmpOff() bool {
	if m != nil {
		return m.IsIcmpOff
	}
	return false
}

func (m *Ipv4AclEdmAce) GetQosGroup() uint32 {
	if m != nil {
		return m.QosGroup
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetNextHopType() string {
	if m != nil {
		return m.NextHopType
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetNextHopInfo() []*Ipv4AclBagNhInfo {
	if m != nil {
		return m.NextHopInfo
	}
	return nil
}

func (m *Ipv4AclEdmAce) GetHwNextHopInfo() *Ipv4AclBagHwNhInfo {
	if m != nil {
		return m.HwNextHopInfo
	}
	return nil
}

func (m *Ipv4AclEdmAce) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *Ipv4AclEdmAce) GetSourcePrefixGroup() string {
	if m != nil {
		return m.SourcePrefixGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPrefixGroup() string {
	if m != nil {
		return m.DestinationPrefixGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSourcePortGroup() string {
	if m != nil {
		return m.SourcePortGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetDestinationPortGroup() string {
	if m != nil {
		return m.DestinationPortGroup
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetSequenceStr() string {
	if m != nil {
		return m.SequenceStr
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetFragmentOffsetOperator() string {
	if m != nil {
		return m.FragmentOffsetOperator
	}
	return ""
}

func (m *Ipv4AclEdmAce) GetFragmentOffset1() uint32 {
	if m != nil {
		return m.FragmentOffset1
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetFragmentOffset2() uint32 {
	if m != nil {
		return m.FragmentOffset2
	}
	return 0
}

func (m *Ipv4AclEdmAce) GetUdf() []*AclUdfAce {
	if m != nil {
		return m.Udf
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4AclEdmAce_KEYS)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_edm_ace_KEYS")
	proto.RegisterType((*Ipv4AclBagNhInfo)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_bag_nh_info")
	proto.RegisterType((*Ipv4AclBagHwNhInfo)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_bag_hw_nh_info")
	proto.RegisterType((*AclUdfAce)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.acl_udf_ace")
	proto.RegisterType((*Ipv4AclEdmAce)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.access_list_manager.accesses.access.access_list_sequences.access_list_sequence.ipv4_acl_edm_ace")
}

func init() { proto.RegisterFile("ipv4_acl_edm_ace.proto", fileDescriptor_5a3b94617fae956f) }

var fileDescriptor_5a3b94617fae956f = []byte{
	// 1222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x76, 0x1c, 0x35,
	0x13, 0x3e, 0x1d, 0xe7, 0xe2, 0xa9, 0xf9, 0xc7, 0xb1, 0x15, 0xc7, 0xbf, 0x42, 0x42, 0x30, 0x86,
	0x1c, 0x06, 0x02, 0x86, 0xe9, 0x18, 0x08, 0x77, 0x02, 0xe4, 0x46, 0x12, 0xdb, 0x8c, 0x43, 0x20,
	0xdc, 0x14, 0x59, 0xad, 0x9e, 0xe9, 0x78, 0xfa, 0xe2, 0x96, 0xc6, 0x71, 0xb6, 0xbc, 0x00, 0x5b,
	0xce, 0x61, 0xc7, 0x4b, 0xb0, 0x62, 0xc9, 0x13, 0xf1, 0x02, 0x1c, 0x55, 0xf5, 0x6d, 0xc6, 0xc3,
	0x9a, 0xac, 0x5a, 0xf5, 0xd5, 0x57, 0xaa, 0x52, 0xb5, 0xaa, 0x4a, 0xb0, 0x12, 0x65, 0x07, 0x1b,
	0x42, 0xaa, 0x91, 0xd0, 0x41, 0x2c, 0xa4, 0xd2, 0xeb, 0x59, 0x9e, 0xda, 0x94, 0xfd, 0xec, 0xa9,
	0xc8, 0xa8, 0x54, 0x44, 0xa9, 0x11, 0x87, 0xb9, 0xa8, 0x58, 0x69, 0xa6, 0xf3, 0xf5, 0x4a, 0x92,
	0x49, 0x20, 0xb2, 0x5c, 0x87, 0xd1, 0xa1, 0x18, 0x45, 0xc6, 0xae, 0x4b, 0xa5, 0xb4, 0x31, 0xb8,
	0x16, 0xb1, 0x4c, 0xe4, 0x40, 0xe7, 0x05, 0xa6, 0x4d, 0xb1, 0x98, 0xe0, 0x18, 0xbd, 0x3f, 0xd6,
	0x89, 0xd2, 0xb3, 0xd1, 0xb5, 0xc7, 0x70, 0x76, 0x3a, 0x3c, 0x71, 0xe7, 0xfa, 0xc3, 0x1d, 0xd6,
	0x85, 0xc5, 0xa6, 0x41, 0x22, 0x63, 0xcd, 0xbd, 0x55, 0xaf, 0xdb, 0xea, 0x2f, 0x10, 0x7e, 0x37,
	0x32, 0x76, 0x53, 0xc6, 0x9a, 0xbd, 0x02, 0xa7, 0xcb, 0xed, 0x44, 0x32, 0x8e, 0x77, 0x75, 0xce,
	0x8f, 0xad, 0x7a, 0xdd, 0x4e, 0x7f, 0xa1, 0x84, 0x37, 0x11, 0x5d, 0xfb, 0xc3, 0x83, 0xe5, 0xca,
	0xd9, 0xae, 0x1c, 0x88, 0x64, 0x28, 0xa2, 0x24, 0x4c, 0xd9, 0x39, 0x98, 0x4f, 0xf4, 0xa1, 0x15,
	0xc3, 0x34, 0x2b, 0x7c, 0x9c, 0x72, 0xf2, 0xad, 0x34, 0x63, 0xcf, 0x03, 0xd8, 0x5c, 0xaa, 0x3d,
	0x0a, 0xe0, 0x18, 0x2a, 0x5b, 0x88, 0xa0, 0xef, 0x15, 0x38, 0x69, 0xac, 0xb4, 0x63, 0xc3, 0xe7,
	0x50, 0x55, 0x48, 0xec, 0x3c, 0xb4, 0xa4, 0x15, 0x85, 0xea, 0x38, 0xaa, 0xe6, 0xa5, 0xdd, 0x21,
	0xe5, 0x5b, 0x70, 0x36, 0x32, 0x18, 0x44, 0xe9, 0x55, 0xe8, 0xc3, 0xc8, 0x58, 0x7e, 0x62, 0xd5,
	0xeb, 0xce, 0xf7, 0x97, 0x22, 0x73, 0x4d, 0x8d, 0x36, 0x29, 0x80, 0xeb, 0x4e, 0xb1, 0xa6, 0xe0,
	0xff, 0x13, 0x81, 0x0f, 0x9f, 0xfc, 0x6b, 0xec, 0x9d, 0x3a, 0x76, 0x06, 0xc7, 0xed, 0xd3, 0xac,
	0x8c, 0x1a, 0xd7, 0x8e, 0x7e, 0x90, 0x87, 0x74, 0x1a, 0x0a, 0xf9, 0xd4, 0x41, 0x1e, 0xba, 0xb3,
	0xac, 0x3d, 0x82, 0xb6, 0xdb, 0x7f, 0x1c, 0x84, 0xee, 0x2f, 0x38, 0xa6, 0x5b, 0x16, 0x89, 0x9f,
	0x73, 0xcc, 0x71, 0x80, 0x4c, 0x77, 0x3a, 0xa7, 0x3a, 0x90, 0xa3, 0xb1, 0x2e, 0x72, 0xed, 0xb8,
	0x0f, 0x9c, 0x5c, 0xda, 0xc5, 0xd2, 0xec, 0xa1, 0x87, 0x0e, 0xda, 0xdd, 0x93, 0x66, 0x6f, 0xed,
	0xef, 0x33, 0xb0, 0x38, 0xfd, 0xb7, 0xdd, 0x66, 0x91, 0xd5, 0xb1, 0xc0, 0x50, 0x7d, 0x4a, 0x95,
	0x03, 0xee, 0xbb, 0x70, 0x9f, 0x83, 0xf9, 0xf2, 0x27, 0xf2, 0x2b, 0xe4, 0xa8, 0x94, 0xd9, 0x32,
	0x9c, 0x18, 0xe4, 0x32, 0xb1, 0x7c, 0x03, 0x8d, 0x48, 0x60, 0x97, 0x61, 0x09, 0xaf, 0xb7, 0x4a,
	0xe9, 0x26, 0x4b, 0x9b, 0xe6, 0xfc, 0x6d, 0x34, 0x5d, 0x2c, 0x15, 0x5b, 0x05, 0xee, 0xb6, 0x2f,
	0x31, 0xfe, 0x0e, 0x6d, 0x5f, 0xca, 0xec, 0x02, 0xb4, 0xca, 0xb5, 0xcf, 0xdf, 0x45, 0x65, 0x0d,
	0xb0, 0x4b, 0xb0, 0x60, 0xd2, 0x71, 0xae, 0xb4, 0x90, 0x41, 0x90, 0x6b, 0x63, 0xf8, 0x55, 0x8c,
	0xa2, 0x43, 0xe8, 0x35, 0x02, 0xd9, 0x3a, 0x9c, 0x99, 0xa4, 0x51, 0x5e, 0xde, 0x43, 0xee, 0xd2,
	0x04, 0xd7, 0x65, 0x88, 0xbd, 0x09, 0x67, 0x02, 0x6d, 0x6c, 0x94, 0x48, 0x1b, 0xa5, 0x49, 0xb5,
	0xf7, 0xfb, 0xc8, 0x67, 0x0d, 0x55, 0xe9, 0xe0, 0x2a, 0xf0, 0x19, 0x06, 0xe4, 0xe5, 0x03, 0xb4,
	0x5a, 0x39, 0x6a, 0x85, 0xae, 0x5c, 0xd9, 0x50, 0x68, 0x55, 0x9a, 0x3e, 0xa4, 0xfa, 0x22, 0xb8,
	0x4a, 0xd2, 0x8b, 0xf0, 0xbf, 0x82, 0x98, 0xa5, 0xb9, 0xed, 0xf1, 0x8f, 0x30, 0x17, 0x6d, 0xc2,
	0xb6, 0x1d, 0x34, 0x45, 0xf1, 0xf9, 0xc7, 0xd3, 0x94, 0x22, 0x61, 0x13, 0xde, 0x3e, 0x29, 0x13,
	0xd6, 0x74, 0xf6, 0x02, 0xb4, 0x89, 0x46, 0xbe, 0x3e, 0xc5, 0x8d, 0x00, 0x21, 0x72, 0x35, 0x41,
	0xf0, 0xf9, 0xb5, 0x29, 0x82, 0xcf, 0x7a, 0xb0, 0xdc, 0xcc, 0x48, 0xe5, 0xee, 0x33, 0x74, 0xd7,
	0x4c, 0x6f, 0xe5, 0xf4, 0x32, 0x2c, 0x35, 0x4d, 0xc8, 0xf5, 0xe7, 0x74, 0x67, 0x1a, 0x0a, 0x0a,
	0x60, 0x06, 0xd9, 0xe7, 0x5f, 0xcc, 0x24, 0xfb, 0xae, 0x7d, 0x8c, 0xd2, 0x81, 0x48, 0x33, 0x07,
	0xf1, 0xeb, 0xd4, 0x3e, 0x46, 0xe9, 0x60, 0x0b, 0x01, 0x97, 0x37, 0x95, 0x8e, 0x13, 0xab, 0x73,
	0xaa, 0xb3, 0x1b, 0x48, 0x68, 0x17, 0x18, 0xd6, 0x1a, 0x87, 0x53, 0x4a, 0x66, 0x76, 0x9c, 0x6b,
	0x7e, 0x13, 0xdb, 0x43, 0x29, 0x3a, 0xe3, 0xc0, 0xa8, 0xcc, 0x75, 0x67, 0xa3, 0x13, 0xcb, 0x6f,
	0xa1, 0xba, 0xed, 0xb0, 0x6d, 0x82, 0x5c, 0x07, 0x70, 0x22, 0xbf, 0x8d, 0xe1, 0xe1, 0xda, 0x95,
	0x8d, 0xfb, 0xfa, 0xfc, 0x4b, 0x04, 0x49, 0x60, 0x2f, 0x41, 0x07, 0x37, 0xab, 0xd2, 0x75, 0x07,
	0xb5, 0xe8, 0xa1, 0xca, 0xd3, 0x1b, 0xc0, 0xb2, 0x5c, 0x2b, 0x1d, 0x60, 0xaf, 0x2d, 0xfd, 0xde,
	0xa5, 0xae, 0x55, 0x6b, 0x4a, 0xef, 0x17, 0x01, 0x6a, 0x90, 0xdf, 0xa3, 0x3f, 0x55, 0x23, 0xec,
	0x75, 0x60, 0x56, 0x65, 0x22, 0x1c, 0xc9, 0x81, 0xa9, 0x1d, 0x6f, 0x62, 0x0e, 0x16, 0xad, 0xca,
	0x6e, 0x38, 0x45, 0xe5, 0xfc, 0x3c, 0xb4, 0x2a, 0x36, 0xdf, 0xa2, 0x62, 0x2d, 0x49, 0xec, 0x65,
	0x58, 0xa8, 0xb7, 0xc2, 0xcb, 0xbf, 0x4d, 0xf1, 0x97, 0x0c, 0xbc, 0xf2, 0x17, 0xa0, 0x15, 0xe6,
	0x72, 0x10, 0xeb, 0xc4, 0x1a, 0xfe, 0x15, 0x95, 0x74, 0x05, 0xb0, 0x0d, 0x58, 0xc9, 0xa4, 0xda,
	0xd3, 0x56, 0x8c, 0x74, 0x32, 0xb0, 0xc3, 0x3a, 0xa4, 0x3e, 0x86, 0xb4, 0x4c, 0xda, 0xbb, 0xa8,
	0xac, 0xc2, 0xba, 0x04, 0x0b, 0x13, 0x56, 0x3d, 0xbe, 0x83, 0x1b, 0x77, 0x9a, 0xec, 0xde, 0x11,
	0x9a, 0xcf, 0xef, 0x1f, 0xa5, 0xf9, 0xee, 0x9f, 0x5a, 0xdb, 0x68, 0x5c, 0x5f, 0xd3, 0x85, 0xb0,
	0xb6, 0xee, 0x59, 0xae, 0xab, 0xdb, 0x51, 0x8f, 0x3f, 0xa0, 0x7f, 0xea, 0xd6, 0x05, 0xe6, 0xf3,
	0x6f, 0x2a, 0xcc, 0xc7, 0xc1, 0x90, 0xe2, 0x08, 0x32, 0xfc, 0x5b, 0xba, 0x39, 0x49, 0xea, 0x26,
	0x90, 0x71, 0xf4, 0x61, 0x64, 0x0d, 0x7f, 0xb8, 0xea, 0x75, 0x8f, 0xf7, 0x71, 0xcd, 0x2e, 0x42,
	0x3b, 0x32, 0x22, 0x52, 0x71, 0x26, 0xd2, 0x30, 0xe4, 0xdf, 0xa1, 0x45, 0x2b, 0x32, 0xb7, 0x55,
	0x9c, 0x6d, 0x85, 0xa1, 0x4b, 0xff, 0x7e, 0x6a, 0xc4, 0x20, 0x4f, 0xc7, 0x19, 0xff, 0x9e, 0xd2,
	0xbf, 0x9f, 0x9a, 0x9b, 0x4e, 0x66, 0x6b, 0xd0, 0xa9, 0x46, 0x19, 0xf6, 0xf1, 0x1f, 0x28, 0xee,
	0x62, 0x12, 0x61, 0x2b, 0xff, 0xd3, 0x6b, 0x90, 0xdc, 0xe8, 0xe2, 0x3f, 0xae, 0xce, 0x75, 0xdb,
	0xfe, 0xaf, 0xde, 0xfa, 0x7f, 0xff, 0x0e, 0x59, 0x9f, 0xf5, 0x2e, 0xa8, 0xe2, 0xbf, 0xed, 0x06,
	0xed, 0x5f, 0x1e, 0x2c, 0xba, 0xb9, 0x3b, 0x71, 0x84, 0x9f, 0x56, 0xbd, 0x6e, 0xdb, 0xff, 0xed,
	0xd9, 0x3b, 0x42, 0xfd, 0x42, 0xe8, 0x77, 0x86, 0x4f, 0x36, 0x1b, 0xe7, 0x58, 0x81, 0x93, 0xb9,
	0x8e, 0x65, 0xbe, 0xc7, 0x05, 0x3d, 0x59, 0x48, 0x72, 0x8d, 0x26, 0x78, 0x9a, 0xc8, 0x38, 0x52,
	0xfc, 0x11, 0x5d, 0x97, 0x42, 0x6c, 0x0c, 0xb1, 0x22, 0x7a, 0xba, 0x04, 0xb2, 0x39, 0xc4, 0xb6,
	0x51, 0x43, 0xb7, 0x61, 0x6a, 0x26, 0x4d, 0x18, 0xed, 0x1e, 0x99, 0x49, 0x4d, 0xcb, 0xd7, 0x60,
	0xa9, 0x31, 0x47, 0x0a, 0x13, 0x85, 0x26, 0xa7, 0xeb, 0x61, 0x42, 0xdc, 0x0d, 0x58, 0x99, 0xee,
	0xc3, 0x85, 0x41, 0x40, 0xe5, 0x3a, 0xd5, 0x8c, 0xc9, 0xea, 0x1c, 0xcc, 0xe3, 0xc3, 0xcb, 0x75,
	0x5b, 0x4d, 0xef, 0x1f, 0xa9, 0x46, 0xd8, 0x69, 0xdd, 0x10, 0x2b, 0xdf, 0x91, 0xc6, 0xe6, 0x3c,
	0xa4, 0x3b, 0x5c, 0x62, 0x3b, 0x36, 0x77, 0x27, 0x2b, 0xfb, 0x85, 0xab, 0x12, 0xa3, 0x6d, 0x5d,
	0xaa, 0x03, 0x3a, 0x59, 0xa9, 0xdf, 0x42, 0x75, 0x55, 0xb5, 0xaf, 0xc2, 0xe2, 0x94, 0x65, 0x8f,
	0x0f, 0xb1, 0x8a, 0x4e, 0x4f, 0x5a, 0xf4, 0x66, 0x50, 0x7d, 0x1e, 0xcd, 0xa2, 0xfa, 0xec, 0x77,
	0x0f, 0xe6, 0xc6, 0x41, 0xc8, 0x1f, 0x63, 0x25, 0xfd, 0xf2, 0x4c, 0x5c, 0xc3, 0xc6, 0x1b, 0xb2,
	0xef, 0x82, 0xdb, 0x3d, 0x89, 0xaf, 0xa6, 0x2b, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0x36,
	0xc1, 0x9e, 0x88, 0x0c, 0x00, 0x00,
}
