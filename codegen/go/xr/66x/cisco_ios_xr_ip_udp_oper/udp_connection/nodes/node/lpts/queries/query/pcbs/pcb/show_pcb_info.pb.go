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
// source: show_pcb_info.proto

package cisco_ios_xr_ip_udp_oper_udp_connection_nodes_node_lpts_queries_query_pcbs_pcb

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

type ShowPcbInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	QueryName            string   `protobuf:"bytes,2,opt,name=query_name,json=queryName,proto3" json:"query_name,omitempty"`
	PcbAddress           string   `protobuf:"bytes,3,opt,name=pcb_address,json=pcbAddress,proto3" json:"pcb_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowPcbInfo_KEYS) Reset()         { *m = ShowPcbInfo_KEYS{} }
func (m *ShowPcbInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ShowPcbInfo_KEYS) ProtoMessage()    {}
func (*ShowPcbInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{0}
}

func (m *ShowPcbInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowPcbInfo_KEYS.Unmarshal(m, b)
}
func (m *ShowPcbInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowPcbInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ShowPcbInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowPcbInfo_KEYS.Merge(m, src)
}
func (m *ShowPcbInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ShowPcbInfo_KEYS.Size(m)
}
func (m *ShowPcbInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowPcbInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ShowPcbInfo_KEYS proto.InternalMessageInfo

func (m *ShowPcbInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ShowPcbInfo_KEYS) GetQueryName() string {
	if m != nil {
		return m.QueryName
	}
	return ""
}

func (m *ShowPcbInfo_KEYS) GetPcbAddress() string {
	if m != nil {
		return m.PcbAddress
	}
	return ""
}

type IpAddr struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpAddr) Reset()         { *m = IpAddr{} }
func (m *IpAddr) String() string { return proto.CompactTextString(m) }
func (*IpAddr) ProtoMessage()    {}
func (*IpAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{1}
}

func (m *IpAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpAddr.Unmarshal(m, b)
}
func (m *IpAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpAddr.Marshal(b, m, deterministic)
}
func (m *IpAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpAddr.Merge(m, src)
}
func (m *IpAddr) XXX_Size() int {
	return xxx_messageInfo_IpAddr.Size(m)
}
func (m *IpAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_IpAddr.DiscardUnknown(m)
}

var xxx_messageInfo_IpAddr proto.InternalMessageInfo

func (m *IpAddr) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IpAddr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IpAddr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type LptsOptions struct {
	IsReceiveFilter      bool     `protobuf:"varint,1,opt,name=is_receive_filter,json=isReceiveFilter,proto3" json:"is_receive_filter,omitempty"`
	IsIpSla              bool     `protobuf:"varint,2,opt,name=is_ip_sla,json=isIpSla,proto3" json:"is_ip_sla,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LptsOptions) Reset()         { *m = LptsOptions{} }
func (m *LptsOptions) String() string { return proto.CompactTextString(m) }
func (*LptsOptions) ProtoMessage()    {}
func (*LptsOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{2}
}

func (m *LptsOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LptsOptions.Unmarshal(m, b)
}
func (m *LptsOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LptsOptions.Marshal(b, m, deterministic)
}
func (m *LptsOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LptsOptions.Merge(m, src)
}
func (m *LptsOptions) XXX_Size() int {
	return xxx_messageInfo_LptsOptions.Size(m)
}
func (m *LptsOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_LptsOptions.DiscardUnknown(m)
}

var xxx_messageInfo_LptsOptions proto.InternalMessageInfo

func (m *LptsOptions) GetIsReceiveFilter() bool {
	if m != nil {
		return m.IsReceiveFilter
	}
	return false
}

func (m *LptsOptions) GetIsIpSla() bool {
	if m != nil {
		return m.IsIpSla
	}
	return false
}

type LptsFlags struct {
	IsPcbBound           bool     `protobuf:"varint,1,opt,name=is_pcb_bound,json=isPcbBound,proto3" json:"is_pcb_bound,omitempty"`
	IsLocalAddressIgnore bool     `protobuf:"varint,2,opt,name=is_local_address_ignore,json=isLocalAddressIgnore,proto3" json:"is_local_address_ignore,omitempty"`
	IsIgnoreVrfFilter    bool     `protobuf:"varint,3,opt,name=is_ignore_vrf_filter,json=isIgnoreVrfFilter,proto3" json:"is_ignore_vrf_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LptsFlags) Reset()         { *m = LptsFlags{} }
func (m *LptsFlags) String() string { return proto.CompactTextString(m) }
func (*LptsFlags) ProtoMessage()    {}
func (*LptsFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{3}
}

func (m *LptsFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LptsFlags.Unmarshal(m, b)
}
func (m *LptsFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LptsFlags.Marshal(b, m, deterministic)
}
func (m *LptsFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LptsFlags.Merge(m, src)
}
func (m *LptsFlags) XXX_Size() int {
	return xxx_messageInfo_LptsFlags.Size(m)
}
func (m *LptsFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_LptsFlags.DiscardUnknown(m)
}

var xxx_messageInfo_LptsFlags proto.InternalMessageInfo

func (m *LptsFlags) GetIsPcbBound() bool {
	if m != nil {
		return m.IsPcbBound
	}
	return false
}

func (m *LptsFlags) GetIsLocalAddressIgnore() bool {
	if m != nil {
		return m.IsLocalAddressIgnore
	}
	return false
}

func (m *LptsFlags) GetIsIgnoreVrfFilter() bool {
	if m != nil {
		return m.IsIgnoreVrfFilter
	}
	return false
}

type LptsAcceptMask struct {
	IsInterface          bool     `protobuf:"varint,1,opt,name=is_interface,json=isInterface,proto3" json:"is_interface,omitempty"`
	IsPacketType         bool     `protobuf:"varint,2,opt,name=is_packet_type,json=isPacketType,proto3" json:"is_packet_type,omitempty"`
	IsRemoteAddress      bool     `protobuf:"varint,3,opt,name=is_remote_address,json=isRemoteAddress,proto3" json:"is_remote_address,omitempty"`
	IsRemotePort         bool     `protobuf:"varint,4,opt,name=is_remote_port,json=isRemotePort,proto3" json:"is_remote_port,omitempty"`
	IsLocalAddress       bool     `protobuf:"varint,5,opt,name=is_local_address,json=isLocalAddress,proto3" json:"is_local_address,omitempty"`
	IsLocalPort          bool     `protobuf:"varint,6,opt,name=is_local_port,json=isLocalPort,proto3" json:"is_local_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LptsAcceptMask) Reset()         { *m = LptsAcceptMask{} }
func (m *LptsAcceptMask) String() string { return proto.CompactTextString(m) }
func (*LptsAcceptMask) ProtoMessage()    {}
func (*LptsAcceptMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{4}
}

func (m *LptsAcceptMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LptsAcceptMask.Unmarshal(m, b)
}
func (m *LptsAcceptMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LptsAcceptMask.Marshal(b, m, deterministic)
}
func (m *LptsAcceptMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LptsAcceptMask.Merge(m, src)
}
func (m *LptsAcceptMask) XXX_Size() int {
	return xxx_messageInfo_LptsAcceptMask.Size(m)
}
func (m *LptsAcceptMask) XXX_DiscardUnknown() {
	xxx_messageInfo_LptsAcceptMask.DiscardUnknown(m)
}

var xxx_messageInfo_LptsAcceptMask proto.InternalMessageInfo

func (m *LptsAcceptMask) GetIsInterface() bool {
	if m != nil {
		return m.IsInterface
	}
	return false
}

func (m *LptsAcceptMask) GetIsPacketType() bool {
	if m != nil {
		return m.IsPacketType
	}
	return false
}

func (m *LptsAcceptMask) GetIsRemoteAddress() bool {
	if m != nil {
		return m.IsRemoteAddress
	}
	return false
}

func (m *LptsAcceptMask) GetIsRemotePort() bool {
	if m != nil {
		return m.IsRemotePort
	}
	return false
}

func (m *LptsAcceptMask) GetIsLocalAddress() bool {
	if m != nil {
		return m.IsLocalAddress
	}
	return false
}

func (m *LptsAcceptMask) GetIsLocalPort() bool {
	if m != nil {
		return m.IsLocalPort
	}
	return false
}

type PktTypeEn struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	IcmpMessageType      string   `protobuf:"bytes,2,opt,name=icmp_message_type,json=icmpMessageType,proto3" json:"icmp_message_type,omitempty"`
	IcmPv6MessageType    string   `protobuf:"bytes,3,opt,name=icm_pv6_message_type,json=icmPv6MessageType,proto3" json:"icm_pv6_message_type,omitempty"`
	IgmpMessageType      string   `protobuf:"bytes,4,opt,name=igmp_message_type,json=igmpMessageType,proto3" json:"igmp_message_type,omitempty"`
	MessageId            uint32   `protobuf:"varint,5,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PktTypeEn) Reset()         { *m = PktTypeEn{} }
func (m *PktTypeEn) String() string { return proto.CompactTextString(m) }
func (*PktTypeEn) ProtoMessage()    {}
func (*PktTypeEn) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{5}
}

func (m *PktTypeEn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PktTypeEn.Unmarshal(m, b)
}
func (m *PktTypeEn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PktTypeEn.Marshal(b, m, deterministic)
}
func (m *PktTypeEn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PktTypeEn.Merge(m, src)
}
func (m *PktTypeEn) XXX_Size() int {
	return xxx_messageInfo_PktTypeEn.Size(m)
}
func (m *PktTypeEn) XXX_DiscardUnknown() {
	xxx_messageInfo_PktTypeEn.DiscardUnknown(m)
}

var xxx_messageInfo_PktTypeEn proto.InternalMessageInfo

func (m *PktTypeEn) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PktTypeEn) GetIcmpMessageType() string {
	if m != nil {
		return m.IcmpMessageType
	}
	return ""
}

func (m *PktTypeEn) GetIcmPv6MessageType() string {
	if m != nil {
		return m.IcmPv6MessageType
	}
	return ""
}

func (m *PktTypeEn) GetIgmpMessageType() string {
	if m != nil {
		return m.IgmpMessageType
	}
	return ""
}

func (m *PktTypeEn) GetMessageId() uint32 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type LptsRxFilter struct {
	InterfaceName        string     `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	PacketType           *PktTypeEn `protobuf:"bytes,2,opt,name=packet_type,json=packetType,proto3" json:"packet_type,omitempty"`
	RemoteAddress        *IpAddr    `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	LocalAddress         *IpAddr    `protobuf:"bytes,4,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteLength         uint32     `protobuf:"varint,5,opt,name=remote_length,json=remoteLength,proto3" json:"remote_length,omitempty"`
	LocalLength          uint32     `protobuf:"varint,6,opt,name=local_length,json=localLength,proto3" json:"local_length,omitempty"`
	ReceiveRemotePort    uint32     `protobuf:"varint,7,opt,name=receive_remote_port,json=receiveRemotePort,proto3" json:"receive_remote_port,omitempty"`
	ReceiveLocalPort     uint32     `protobuf:"varint,8,opt,name=receive_local_port,json=receiveLocalPort,proto3" json:"receive_local_port,omitempty"`
	Priority             uint32     `protobuf:"varint,9,opt,name=priority,proto3" json:"priority,omitempty"`
	Ttl                  uint32     `protobuf:"varint,10,opt,name=ttl,proto3" json:"ttl,omitempty"`
	FlowTypesInfo        uint32     `protobuf:"varint,11,opt,name=flow_types_info,json=flowTypesInfo,proto3" json:"flow_types_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LptsRxFilter) Reset()         { *m = LptsRxFilter{} }
func (m *LptsRxFilter) String() string { return proto.CompactTextString(m) }
func (*LptsRxFilter) ProtoMessage()    {}
func (*LptsRxFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{6}
}

func (m *LptsRxFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LptsRxFilter.Unmarshal(m, b)
}
func (m *LptsRxFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LptsRxFilter.Marshal(b, m, deterministic)
}
func (m *LptsRxFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LptsRxFilter.Merge(m, src)
}
func (m *LptsRxFilter) XXX_Size() int {
	return xxx_messageInfo_LptsRxFilter.Size(m)
}
func (m *LptsRxFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_LptsRxFilter.DiscardUnknown(m)
}

var xxx_messageInfo_LptsRxFilter proto.InternalMessageInfo

func (m *LptsRxFilter) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *LptsRxFilter) GetPacketType() *PktTypeEn {
	if m != nil {
		return m.PacketType
	}
	return nil
}

func (m *LptsRxFilter) GetRemoteAddress() *IpAddr {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

func (m *LptsRxFilter) GetLocalAddress() *IpAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *LptsRxFilter) GetRemoteLength() uint32 {
	if m != nil {
		return m.RemoteLength
	}
	return 0
}

func (m *LptsRxFilter) GetLocalLength() uint32 {
	if m != nil {
		return m.LocalLength
	}
	return 0
}

func (m *LptsRxFilter) GetReceiveRemotePort() uint32 {
	if m != nil {
		return m.ReceiveRemotePort
	}
	return 0
}

func (m *LptsRxFilter) GetReceiveLocalPort() uint32 {
	if m != nil {
		return m.ReceiveLocalPort
	}
	return 0
}

func (m *LptsRxFilter) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LptsRxFilter) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *LptsRxFilter) GetFlowTypesInfo() uint32 {
	if m != nil {
		return m.FlowTypesInfo
	}
	return 0
}

type LptsPcb struct {
	Options              *LptsOptions    `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	LptsFlags            *LptsFlags      `protobuf:"bytes,2,opt,name=lpts_flags,json=lptsFlags,proto3" json:"lpts_flags,omitempty"`
	AcceptMask           *LptsAcceptMask `protobuf:"bytes,3,opt,name=accept_mask,json=acceptMask,proto3" json:"accept_mask,omitempty"`
	Filter               []*LptsRxFilter `protobuf:"bytes,4,rep,name=filter,proto3" json:"filter,omitempty"`
	Ttl                  uint32          `protobuf:"varint,5,opt,name=ttl,proto3" json:"ttl,omitempty"`
	FlowTypesInfo        uint32          `protobuf:"varint,6,opt,name=flow_types_info,json=flowTypesInfo,proto3" json:"flow_types_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LptsPcb) Reset()         { *m = LptsPcb{} }
func (m *LptsPcb) String() string { return proto.CompactTextString(m) }
func (*LptsPcb) ProtoMessage()    {}
func (*LptsPcb) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{7}
}

func (m *LptsPcb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LptsPcb.Unmarshal(m, b)
}
func (m *LptsPcb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LptsPcb.Marshal(b, m, deterministic)
}
func (m *LptsPcb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LptsPcb.Merge(m, src)
}
func (m *LptsPcb) XXX_Size() int {
	return xxx_messageInfo_LptsPcb.Size(m)
}
func (m *LptsPcb) XXX_DiscardUnknown() {
	xxx_messageInfo_LptsPcb.DiscardUnknown(m)
}

var xxx_messageInfo_LptsPcb proto.InternalMessageInfo

func (m *LptsPcb) GetOptions() *LptsOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LptsPcb) GetLptsFlags() *LptsFlags {
	if m != nil {
		return m.LptsFlags
	}
	return nil
}

func (m *LptsPcb) GetAcceptMask() *LptsAcceptMask {
	if m != nil {
		return m.AcceptMask
	}
	return nil
}

func (m *LptsPcb) GetFilter() []*LptsRxFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *LptsPcb) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *LptsPcb) GetFlowTypesInfo() uint32 {
	if m != nil {
		return m.FlowTypesInfo
	}
	return 0
}

type CommonPcbInfo struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LptsPcb              *LptsPcb `protobuf:"bytes,2,opt,name=lpts_pcb,json=lptsPcb,proto3" json:"lpts_pcb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonPcbInfo) Reset()         { *m = CommonPcbInfo{} }
func (m *CommonPcbInfo) String() string { return proto.CompactTextString(m) }
func (*CommonPcbInfo) ProtoMessage()    {}
func (*CommonPcbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{8}
}

func (m *CommonPcbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonPcbInfo.Unmarshal(m, b)
}
func (m *CommonPcbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonPcbInfo.Marshal(b, m, deterministic)
}
func (m *CommonPcbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonPcbInfo.Merge(m, src)
}
func (m *CommonPcbInfo) XXX_Size() int {
	return xxx_messageInfo_CommonPcbInfo.Size(m)
}
func (m *CommonPcbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonPcbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CommonPcbInfo proto.InternalMessageInfo

func (m *CommonPcbInfo) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *CommonPcbInfo) GetLptsPcb() *LptsPcb {
	if m != nil {
		return m.LptsPcb
	}
	return nil
}

type ShowPcbInfo struct {
	L4Protocol           uint32         `protobuf:"varint,50,opt,name=l4_protocol,json=l4Protocol,proto3" json:"l4_protocol,omitempty"`
	LocalPort            uint32         `protobuf:"varint,51,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	ForeignPort          uint32         `protobuf:"varint,52,opt,name=foreign_port,json=foreignPort,proto3" json:"foreign_port,omitempty"`
	LocalAddress         *IpAddr        `protobuf:"bytes,53,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ForeignAddress       *IpAddr        `protobuf:"bytes,54,opt,name=foreign_address,json=foreignAddress,proto3" json:"foreign_address,omitempty"`
	Common               *CommonPcbInfo `protobuf:"bytes,55,opt,name=common,proto3" json:"common,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ShowPcbInfo) Reset()         { *m = ShowPcbInfo{} }
func (m *ShowPcbInfo) String() string { return proto.CompactTextString(m) }
func (*ShowPcbInfo) ProtoMessage()    {}
func (*ShowPcbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_477fa41837dd68d6, []int{9}
}

func (m *ShowPcbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowPcbInfo.Unmarshal(m, b)
}
func (m *ShowPcbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowPcbInfo.Marshal(b, m, deterministic)
}
func (m *ShowPcbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowPcbInfo.Merge(m, src)
}
func (m *ShowPcbInfo) XXX_Size() int {
	return xxx_messageInfo_ShowPcbInfo.Size(m)
}
func (m *ShowPcbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowPcbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ShowPcbInfo proto.InternalMessageInfo

func (m *ShowPcbInfo) GetL4Protocol() uint32 {
	if m != nil {
		return m.L4Protocol
	}
	return 0
}

func (m *ShowPcbInfo) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *ShowPcbInfo) GetForeignPort() uint32 {
	if m != nil {
		return m.ForeignPort
	}
	return 0
}

func (m *ShowPcbInfo) GetLocalAddress() *IpAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *ShowPcbInfo) GetForeignAddress() *IpAddr {
	if m != nil {
		return m.ForeignAddress
	}
	return nil
}

func (m *ShowPcbInfo) GetCommon() *CommonPcbInfo {
	if m != nil {
		return m.Common
	}
	return nil
}

func init() {
	proto.RegisterType((*ShowPcbInfo_KEYS)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.show_pcb_info_KEYS")
	proto.RegisterType((*IpAddr)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.ipAddr")
	proto.RegisterType((*LptsOptions)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.lpts_options")
	proto.RegisterType((*LptsFlags)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.lpts_flags")
	proto.RegisterType((*LptsAcceptMask)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.lpts_accept_mask")
	proto.RegisterType((*PktTypeEn)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.pkt_type_en")
	proto.RegisterType((*LptsRxFilter)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.lptsRxFilter")
	proto.RegisterType((*LptsPcb)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.lpts_pcb")
	proto.RegisterType((*CommonPcbInfo)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.common_pcb_info")
	proto.RegisterType((*ShowPcbInfo)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.lpts.queries.query.pcbs.pcb.show_pcb_info")
}

func init() { proto.RegisterFile("show_pcb_info.proto", fileDescriptor_477fa41837dd68d6) }

var fileDescriptor_477fa41837dd68d6 = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x56, 0x48, 0x36, 0x3f, 0xc7, 0xf9, 0xd9, 0x9d, 0x56, 0x6a, 0x54, 0x54, 0xb1, 0x0d, 0x3f,
	0xaa, 0x10, 0x32, 0xd2, 0x76, 0xbb, 0x5c, 0x83, 0x44, 0xa5, 0x88, 0xb6, 0x8a, 0x5c, 0xb4, 0xe2,
	0x4f, 0x1a, 0xec, 0xc9, 0x38, 0x1d, 0xad, 0xed, 0x99, 0x7a, 0xbc, 0xd9, 0x46, 0x80, 0x10, 0x2f,
	0xc1, 0x2d, 0x0f, 0xc0, 0x15, 0x8f, 0xc1, 0xa3, 0xf0, 0x0c, 0xdc, 0xa0, 0x39, 0x33, 0xe3, 0x75,
	0x42, 0xe1, 0x6a, 0xd3, 0x9b, 0xc4, 0xfe, 0xce, 0xe7, 0xf3, 0x9d, 0x73, 0xe6, 0x9c, 0x63, 0xc3,
	0x2d, 0xfd, 0x42, 0x5e, 0x51, 0xc5, 0x12, 0x2a, 0x8a, 0x54, 0x86, 0xaa, 0x94, 0x95, 0x24, 0xcf,
	0x98, 0xd0, 0x4c, 0x52, 0x21, 0x35, 0x7d, 0x55, 0x52, 0xa1, 0xe8, 0xe5, 0x52, 0x51, 0xa9, 0x78,
	0x19, 0x9a, 0x0b, 0x26, 0x8b, 0x82, 0xb3, 0x4a, 0xc8, 0x22, 0x2c, 0xe4, 0x92, 0x6b, 0xfc, 0x0d,
	0x33, 0x55, 0xe9, 0xf0, 0xe5, 0x25, 0x2f, 0x05, 0xb7, 0xff, 0x9b, 0x50, 0xb1, 0x44, 0x9b, 0x9f,
	0xd9, 0x4b, 0x20, 0x5b, 0x32, 0xf4, 0x8b, 0xcf, 0xbf, 0x7e, 0x4e, 0xde, 0x86, 0x81, 0x79, 0x92,
	0x16, 0x71, 0xce, 0xa7, 0xad, 0xe3, 0xd6, 0x83, 0x41, 0xd4, 0x37, 0xc0, 0xb3, 0x38, 0xe7, 0xe4,
	0x1e, 0x00, 0x3a, 0xb1, 0xd6, 0xb7, 0xd0, 0x3a, 0x40, 0x04, 0xcd, 0xef, 0x40, 0x60, 0x9c, 0xc5,
	0xcb, 0x65, 0xc9, 0xb5, 0x9e, 0xb6, 0xd1, 0x0e, 0x8a, 0x25, 0x9f, 0x5a, 0x64, 0xb6, 0x82, 0xae,
	0x50, 0xe6, 0x86, 0xdc, 0x81, 0x5e, 0x9c, 0x36, 0x45, 0xba, 0x71, 0x8a, 0x3e, 0xee, 0xc3, 0x50,
	0xa8, 0xf5, 0x69, 0xed, 0xc4, 0x8a, 0x04, 0x06, 0x73, 0x5e, 0x1c, 0xe5, 0x6c, 0x47, 0xc7, 0x50,
	0xce, 0xbc, 0xd0, 0x39, 0x0c, 0x4d, 0xea, 0x54, 0x2a, 0x53, 0x11, 0x4d, 0x3e, 0x84, 0x23, 0xa1,
	0x69, 0xc9, 0x19, 0x17, 0x6b, 0x4e, 0x53, 0x91, 0x55, 0xbc, 0x44, 0xe1, 0x7e, 0x34, 0x11, 0x3a,
	0xb2, 0xf8, 0x63, 0x84, 0xc9, 0x5d, 0x18, 0x08, 0x6d, 0xea, 0xab, 0xb3, 0x18, 0xe5, 0xfb, 0x51,
	0x4f, 0xe8, 0xb9, 0x7a, 0x9e, 0xc5, 0xb3, 0x5f, 0x5b, 0x00, 0xe8, 0x38, 0xcd, 0xe2, 0x95, 0x26,
	0xc7, 0x30, 0x14, 0x1a, 0x0b, 0x98, 0xc8, 0xcb, 0x62, 0xe9, 0x3c, 0x82, 0xd0, 0x0b, 0x96, 0x7c,
	0x66, 0x10, 0xf2, 0x08, 0xee, 0x08, 0x4d, 0x33, 0xc9, 0xe2, 0xcc, 0xc7, 0x4b, 0xc5, 0xaa, 0x90,
	0x25, 0x77, 0xae, 0x6f, 0x0b, 0xfd, 0xc4, 0x58, 0x5d, 0xe4, 0x73, 0xb4, 0x91, 0x8f, 0xe1, 0xb6,
	0xf0, 0x44, 0xba, 0x2e, 0x53, 0x1f, 0x72, 0x1b, 0x9f, 0x39, 0x12, 0x8e, 0x77, 0x5e, 0xa6, 0x36,
	0xe8, 0xd9, 0xdf, 0x2d, 0x38, 0xc4, 0xc0, 0x62, 0xc6, 0xb8, 0xaa, 0x68, 0x1e, 0xeb, 0x0b, 0x2c,
	0x94, 0xa6, 0xa2, 0xa8, 0x78, 0x99, 0xc6, 0x8c, 0xbb, 0xf0, 0x02, 0xa1, 0xe7, 0x1e, 0x22, 0xef,
	0xc1, 0xd8, 0x64, 0x10, 0xb3, 0x0b, 0x5e, 0xd1, 0x6a, 0xa3, 0x7c, 0x58, 0x43, 0xa1, 0x17, 0x08,
	0x7e, 0xb9, 0x51, 0xbc, 0x2e, 0x5f, 0x2e, 0x2b, 0xbe, 0x55, 0x76, 0x57, 0x3e, 0x83, 0xfb, 0xd3,
	0xb1, 0x1e, 0x1d, 0x57, 0xc9, 0xb2, 0x9a, 0x76, 0xbc, 0x47, 0x4b, 0x5c, 0xc8, 0xb2, 0x22, 0x0f,
	0xe0, 0x70, 0xb7, 0x2e, 0xd3, 0x03, 0xe4, 0x8d, 0xb7, 0x0b, 0x42, 0x66, 0x30, 0xaa, 0x99, 0xe8,
	0xae, 0xeb, 0xb3, 0x40, 0x9a, 0xf1, 0x36, 0xfb, 0xb3, 0x05, 0x81, 0xba, 0xb0, 0x09, 0x50, 0x5e,
	0x10, 0x02, 0x1d, 0xcc, 0xc5, 0xb6, 0x16, 0x5e, 0x63, 0x0e, 0x2c, 0x57, 0x34, 0xe7, 0x5a, 0xc7,
	0x2b, 0x7e, 0x9d, 0xec, 0x20, 0x9a, 0x18, 0xc3, 0x53, 0x8b, 0x63, 0xbe, 0xa6, 0xfc, 0x2c, 0xa7,
	0xa6, 0xc9, 0xb6, 0xe8, 0xb6, 0xd3, 0x8c, 0x9f, 0xc5, 0xfa, 0xac, 0xf9, 0x80, 0x71, 0xbe, 0xda,
	0x75, 0xde, 0x71, 0xce, 0x57, 0xdb, 0xce, 0xef, 0x01, 0x78, 0x9a, 0x58, 0x62, 0xd2, 0xa3, 0x68,
	0xe0, 0x90, 0xf9, 0x72, 0xf6, 0xfb, 0x81, 0xed, 0xdd, 0xe8, 0x95, 0xeb, 0xc7, 0xf7, 0x61, 0x5c,
	0x1f, 0x61, 0x73, 0x62, 0x46, 0x35, 0x8a, 0x83, 0xf3, 0x23, 0x04, 0xbb, 0xc7, 0x18, 0x9c, 0x7c,
	0x1b, 0xde, 0xec, 0xd2, 0x08, 0x1b, 0x55, 0x8e, 0x40, 0x5d, 0x77, 0xc8, 0x4f, 0x30, 0x7e, 0x4d,
	0x7b, 0x04, 0x27, 0xe7, 0x37, 0x1d, 0x80, 0xdd, 0x1f, 0xd1, 0xa8, 0xdc, 0x6a, 0xba, 0x1f, 0x60,
	0xb4, 0xdd, 0x4b, 0x9d, 0xbd, 0xaa, 0x0f, 0xb3, 0x66, 0x87, 0xbe, 0x0b, 0x2e, 0x1a, 0x9a, 0xf1,
	0x62, 0x55, 0xbd, 0x70, 0x67, 0x3a, 0xb4, 0xe0, 0x13, 0xc4, 0xcc, 0x2c, 0xda, 0x08, 0x1d, 0xa7,
	0x8b, 0x9c, 0x00, 0x31, 0x47, 0x09, 0xe1, 0x96, 0xdf, 0x50, 0xcd, 0xf1, 0xe9, 0x21, 0xf3, 0xc8,
	0x99, 0x1a, 0x33, 0xf4, 0x11, 0x10, 0xcf, 0x6f, 0x8c, 0x47, 0x1f, 0xe9, 0x87, 0xce, 0x52, 0xcf,
	0x08, 0xb9, 0x0b, 0x7d, 0x55, 0x0a, 0x59, 0x8a, 0x6a, 0x33, 0x1d, 0x20, 0xa7, 0xbe, 0x27, 0x87,
	0xd0, 0xae, 0xaa, 0x6c, 0x0a, 0x08, 0x9b, 0x4b, 0xf2, 0x01, 0x4c, 0xd2, 0x4c, 0x5e, 0xe1, 0x59,
	0x6b, 0x7c, 0x3d, 0x4c, 0x03, 0xb4, 0x8e, 0x0c, 0x6c, 0x8e, 0x5c, 0xcf, 0x8b, 0x54, 0xce, 0xfe,
	0xe8, 0x40, 0x1f, 0xf7, 0x8e, 0x62, 0x09, 0x59, 0x43, 0xcf, 0x2d, 0x5c, 0x6c, 0xd1, 0xe0, 0xe4,
	0xbb, 0x9b, 0xae, 0x7f, 0x73, 0xa9, 0x47, 0x5e, 0x8c, 0x6c, 0x9a, 0x4b, 0xd9, 0x75, 0xfe, 0x37,
	0x7b, 0x91, 0x46, 0x85, 0x68, 0x60, 0xae, 0x1f, 0xe3, 0x1b, 0xe0, 0x97, 0x16, 0x04, 0x8d, 0x95,
	0xeb, 0xba, 0xfe, 0xfb, 0xbd, 0x88, 0x37, 0x74, 0x22, 0xb0, 0x37, 0x4f, 0xcd, 0x9a, 0xaf, 0xa0,
	0xeb, 0x5e, 0x0f, 0x9d, 0xe3, 0xf6, 0xbe, 0xaa, 0xee, 0xd7, 0x51, 0xe4, 0xb4, 0x7c, 0xcf, 0x1c,
	0xfc, 0x6f, 0xcf, 0x74, 0x5f, 0xd7, 0x33, 0xbf, 0xb5, 0x60, 0xc2, 0x64, 0x9e, 0xcb, 0xa2, 0xfe,
	0xf6, 0xf8, 0xef, 0xef, 0x01, 0x7d, 0xdd, 0x5f, 0xee, 0x64, 0xbf, 0xda, 0x4b, 0x71, 0x15, 0x4b,
	0xa2, 0x9e, 0xb9, 0x5a, 0xb0, 0x64, 0xf6, 0x57, 0x1b, 0x46, 0x5b, 0xdf, 0x46, 0xe6, 0xd3, 0x26,
	0x3b, 0xa5, 0xf8, 0x21, 0xc6, 0x64, 0x36, 0x3d, 0xc1, 0xbc, 0x20, 0x3b, 0x5d, 0x38, 0xc4, 0x6c,
	0xf5, 0xc6, 0x10, 0x3e, 0xb4, 0x5b, 0x3d, 0xab, 0xa7, 0xef, 0x3e, 0x0c, 0x53, 0x59, 0x72, 0xb1,
	0x2a, 0x2c, 0xe1, 0xd4, 0x8e, 0xbf, 0xc3, 0x90, 0xf2, 0xaf, 0x1d, 0xf6, 0xe8, 0x0d, 0xee, 0xb0,
	0x9f, 0x61, 0xe2, 0xe3, 0xf3, 0xf2, 0x67, 0x7b, 0x95, 0x1f, 0x3b, 0x39, 0x1f, 0xc0, 0x15, 0x74,
	0x6d, 0x4f, 0x4c, 0x3f, 0x41, 0x5d, 0x7a, 0xd3, 0xba, 0x3b, 0x1d, 0x17, 0x39, 0xb9, 0xa4, 0x8b,
	0x87, 0xfa, 0xf0, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xd3, 0xb6, 0x79, 0x74, 0x0b, 0x00,
	0x00,
}
