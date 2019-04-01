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

package cisco_ios_xr_ip_tcp_oper_tcp_connection_nodes_node_extended_information_display_types_display_type_connection_id

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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DispType             string   `protobuf:"bytes,2,opt,name=disp_type,json=dispType,proto3" json:"disp_type,omitempty"`
	PcbId                string   `protobuf:"bytes,3,opt,name=pcb_id,json=pcbId,proto3" json:"pcb_id,omitempty"`
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

func (m *ShowPcbInfo_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ShowPcbInfo_KEYS) GetDispType() string {
	if m != nil {
		return m.DispType
	}
	return ""
}

func (m *ShowPcbInfo_KEYS) GetPcbId() string {
	if m != nil {
		return m.PcbId
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
	proto.RegisterType((*ShowPcbInfo_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.show_pcb_info_KEYS")
	proto.RegisterType((*IpAddr)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.ipAddr")
	proto.RegisterType((*LptsOptions)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.lpts_options")
	proto.RegisterType((*LptsFlags)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.lpts_flags")
	proto.RegisterType((*LptsAcceptMask)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.lpts_accept_mask")
	proto.RegisterType((*PktTypeEn)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.pkt_type_en")
	proto.RegisterType((*LptsRxFilter)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.lptsRxFilter")
	proto.RegisterType((*LptsPcb)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.lpts_pcb")
	proto.RegisterType((*CommonPcbInfo)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.common_pcb_info")
	proto.RegisterType((*ShowPcbInfo)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.extended_information.display_types.display_type.connection_id.show_pcb_info")
}

func init() { proto.RegisterFile("show_pcb_info.proto", fileDescriptor_477fa41837dd68d6) }

var fileDescriptor_477fa41837dd68d6 = []byte{
	// 987 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x5d, 0x8b, 0xe4, 0x44,
	0x14, 0x25, 0xd3, 0x3d, 0xe9, 0xee, 0x9b, 0xfe, 0x98, 0xa9, 0x5d, 0xd9, 0xb0, 0x22, 0xce, 0xc6,
	0x0f, 0x16, 0x91, 0x16, 0x66, 0x67, 0xc7, 0xe7, 0x15, 0x5c, 0x68, 0xdc, 0x95, 0x26, 0x2b, 0x8b,
	0x3e, 0x15, 0xe9, 0x4a, 0xa5, 0xb7, 0x98, 0x24, 0x55, 0xa4, 0xe2, 0xec, 0x8c, 0xa8, 0xa0, 0xbf,
	0x41, 0x17, 0x51, 0x54, 0x04, 0xff, 0x86, 0x2f, 0xbe, 0xf9, 0xec, 0xcf, 0xf1, 0x45, 0xea, 0x56,
	0xa5, 0x27, 0x19, 0x47, 0x1f, 0xa7, 0x7d, 0x69, 0x2a, 0xa7, 0x4e, 0x6e, 0x9f, 0x7b, 0xeb, 0xdc,
	0x5b, 0x81, 0x1b, 0xfa, 0x99, 0x7c, 0x4e, 0x15, 0x5b, 0x51, 0x51, 0x66, 0x72, 0xae, 0x2a, 0x59,
	0x4b, 0xa2, 0x98, 0xd0, 0x4c, 0x52, 0x21, 0x35, 0x3d, 0xab, 0xa8, 0x50, 0xb4, 0x66, 0x8a, 0x4a,
	0xc5, 0xab, 0xb9, 0x59, 0x30, 0x59, 0x96, 0x9c, 0xd5, 0x42, 0x96, 0xf3, 0x52, 0xa6, 0x5c, 0xe3,
	0xef, 0x9c, 0x9f, 0xd5, 0xbc, 0x4c, 0x79, 0x8a, 0x71, 0xaa, 0x22, 0xc1, 0xfd, 0x54, 0x68, 0x95,
	0x27, 0xe7, 0xb4, 0x3e, 0x57, 0x5c, 0x77, 0x9e, 0xe6, 0x17, 0x51, 0xa8, 0x48, 0xa3, 0x8f, 0x81,
	0x74, 0x84, 0xd0, 0x0f, 0xde, 0xff, 0xe4, 0x09, 0x99, 0xc2, 0x8e, 0x48, 0x43, 0xef, 0xc0, 0xbb,
	0x3b, 0x8a, 0x77, 0x44, 0x4a, 0x5e, 0x86, 0x91, 0x89, 0x81, 0x01, 0xc2, 0x1d, 0x84, 0x87, 0x06,
	0xf8, 0xe8, 0x5c, 0x71, 0xf2, 0x12, 0xf8, 0xf8, 0x76, 0x1a, 0xf6, 0x70, 0x67, 0x57, 0xb1, 0xd5,
	0x22, 0x8d, 0xd6, 0xe0, 0x0b, 0xf5, 0x20, 0x4d, 0x2b, 0x72, 0x0b, 0x06, 0x49, 0x46, 0xcb, 0xa4,
	0xe0, 0x2e, 0xa4, 0x9f, 0x64, 0x1f, 0x26, 0x05, 0x27, 0x77, 0x60, 0x2c, 0xd4, 0xe9, 0x11, 0x4d,
	0xd2, 0xb4, 0xe2, 0x5a, 0xbb, 0xc8, 0x81, 0xc1, 0x1e, 0x58, 0xc8, 0x51, 0x8e, 0x37, 0x94, 0xde,
	0x86, 0x72, 0xec, 0x28, 0xd1, 0x53, 0x18, 0xe7, 0xaa, 0xd6, 0x54, 0x2a, 0x93, 0x94, 0x26, 0x6f,
	0xc1, 0xbe, 0xd0, 0xb4, 0xe2, 0x8c, 0x8b, 0x53, 0x4e, 0x33, 0x91, 0xd7, 0xbc, 0xc2, 0x3f, 0x1e,
	0xc6, 0x33, 0xa1, 0x63, 0x8b, 0x3f, 0x44, 0x98, 0xdc, 0x86, 0x91, 0xd0, 0xa6, 0xd0, 0x3a, 0x4f,
	0xf0, 0xef, 0x87, 0xf1, 0x40, 0xe8, 0x85, 0x7a, 0x92, 0x27, 0xd1, 0x0b, 0x0f, 0x00, 0x03, 0x67,
	0x79, 0xb2, 0xd6, 0xe4, 0x00, 0xc6, 0x42, 0x63, 0x9d, 0x56, 0xf2, 0xd3, 0x32, 0x75, 0x11, 0x41,
	0xe8, 0x25, 0x5b, 0xbd, 0x67, 0x10, 0x72, 0x1f, 0x6e, 0x09, 0x4d, 0x73, 0xc9, 0x92, 0xbc, 0xd1,
	0x4b, 0xc5, 0xba, 0x94, 0x15, 0x77, 0xa1, 0x6f, 0x0a, 0xfd, 0xc8, 0xec, 0x3a, 0xe5, 0x0b, 0xdc,
	0x23, 0xef, 0xc0, 0x4d, 0xd1, 0x10, 0xe9, 0x69, 0x95, 0x35, 0x92, 0x7b, 0xf8, 0xce, 0xbe, 0x70,
	0xbc, 0xa7, 0x55, 0x66, 0x45, 0x47, 0x7f, 0x79, 0xb0, 0x87, 0xc2, 0x12, 0xc6, 0xb8, 0xaa, 0x69,
	0x91, 0xe8, 0x13, 0x2c, 0x94, 0xa6, 0xa2, 0xac, 0x79, 0x95, 0x25, 0x8c, 0x3b, 0x79, 0x81, 0xd0,
	0x8b, 0x06, 0x22, 0xaf, 0xc3, 0xd4, 0x64, 0x90, 0xb0, 0x13, 0x5e, 0x5f, 0x1c, 0xe5, 0x30, 0x1e,
	0x0b, 0xbd, 0x44, 0x10, 0x8f, 0xb3, 0x29, 0x5f, 0x21, 0x6b, 0xde, 0x29, 0xbb, 0x2b, 0x9f, 0xc1,
	0x9b, 0xd3, 0xb1, 0x11, 0x1d, 0x57, 0xc9, 0xaa, 0x0e, 0xfb, 0x4d, 0x44, 0x4b, 0x5c, 0xca, 0xaa,
	0x26, 0x77, 0x61, 0xef, 0x72, 0x5d, 0xc2, 0x5d, 0xe4, 0x4d, 0xbb, 0x05, 0x21, 0x11, 0x4c, 0x36,
	0x4c, 0x0c, 0xe7, 0x37, 0x59, 0x20, 0xcd, 0x44, 0x8b, 0xfe, 0xf0, 0x20, 0x50, 0x27, 0x36, 0x01,
	0xca, 0x4b, 0x42, 0xa0, 0x8f, 0xb9, 0x58, 0x6b, 0xe1, 0x1a, 0x73, 0x60, 0x85, 0xa2, 0x05, 0xd7,
	0x3a, 0x59, 0xf3, 0xb6, 0x6f, 0x67, 0x66, 0xe3, 0xb1, 0xc5, 0x31, 0x5f, 0x53, 0x7e, 0x56, 0x50,
	0x63, 0xb2, 0x0e, 0xdd, 0x3a, 0xcd, 0xc4, 0x59, 0x9e, 0x1e, 0xb7, 0x5f, 0x30, 0xc1, 0xd7, 0x97,
	0x83, 0xf7, 0x5d, 0xf0, 0x75, 0x37, 0xf8, 0x2b, 0x00, 0x0d, 0x4d, 0xa4, 0x98, 0xf4, 0x24, 0x1e,
	0x39, 0x64, 0x91, 0x46, 0xdf, 0xf8, 0xd6, 0xbb, 0xf1, 0x99, 0xf3, 0xe3, 0x1b, 0x30, 0xdd, 0x1c,
	0x61, 0xbb, 0x63, 0x26, 0x1b, 0x14, 0x1b, 0xe7, 0x27, 0x53, 0x83, 0x4b, 0xe7, 0x18, 0x1c, 0x7e,
	0x31, 0xbf, 0xee, 0xf1, 0x31, 0x6f, 0x1d, 0x44, 0x0c, 0xea, 0xc2, 0x44, 0x3f, 0x7b, 0x30, 0xbd,
	0xc2, 0x42, 0xc1, 0xe1, 0xd9, 0xf5, 0x6b, 0xb4, 0x53, 0x28, 0x9e, 0x54, 0x1d, 0xeb, 0xfe, 0xe8,
	0xc1, 0xa4, 0x6b, 0xc9, 0xfe, 0x96, 0x05, 0x8e, 0xf3, 0x76, 0x2b, 0xbc, 0x06, 0x4e, 0x30, 0xcd,
	0x79, 0xb9, 0xae, 0x9f, 0x39, 0xf3, 0x8c, 0x2d, 0xf8, 0x08, 0x31, 0xd3, 0xf4, 0x36, 0x07, 0xc7,
	0xf1, 0x91, 0x13, 0x20, 0xe6, 0x28, 0x73, 0xb8, 0xd1, 0x8c, 0xc2, 0x76, 0x9f, 0x0e, 0x90, 0xb9,
	0xef, 0xb6, 0x5a, 0xcd, 0xfa, 0x36, 0x90, 0x86, 0xdf, 0xea, 0xc3, 0x21, 0xd2, 0xf7, 0xdc, 0xce,
	0xa6, 0x19, 0xc9, 0x6d, 0x18, 0xaa, 0x4a, 0xc8, 0x4a, 0xd4, 0xe7, 0xe1, 0x08, 0x39, 0x9b, 0x67,
	0xb2, 0x07, 0xbd, 0xba, 0xce, 0x43, 0x40, 0xd8, 0x2c, 0xc9, 0x9b, 0x30, 0xcb, 0x72, 0xf9, 0xdc,
	0x56, 0x04, 0x6b, 0x15, 0x06, 0xb8, 0x3b, 0x31, 0xb0, 0x31, 0x8e, 0x5e, 0x94, 0x99, 0x8c, 0x7e,
	0xdb, 0x85, 0x21, 0x0e, 0x38, 0xc5, 0x56, 0xe4, 0x3b, 0x0f, 0x06, 0x6e, 0xb4, 0x63, 0x33, 0x04,
	0x87, 0x5f, 0x5e, 0xff, 0x11, 0xb5, 0x2f, 0x98, 0xb8, 0x91, 0x43, 0x7e, 0xe8, 0xdc, 0x10, 0xae,
	0x0b, 0x3f, 0xdf, 0x92, 0x3a, 0xd4, 0x10, 0x8f, 0xcc, 0xfa, 0x21, 0x5e, 0x58, 0xbf, 0x7a, 0x10,
	0xb4, 0x6e, 0x08, 0xd7, 0x80, 0x5f, 0x7b, 0x5b, 0xd2, 0xd7, 0x92, 0x12, 0x83, 0x7d, 0x78, 0x6c,
	0x2e, 0xae, 0x17, 0x1e, 0xf8, 0xee, 0xc6, 0xeb, 0x1f, 0xf4, 0xb6, 0x77, 0xbc, 0xcd, 0x0c, 0x8e,
	0x9d, 0x9a, 0xc6, 0xbf, 0xbb, 0xff, 0xe9, 0x5f, 0xff, 0x2a, 0xff, 0xfe, 0xee, 0xc1, 0x8c, 0xc9,
	0xa2, 0x90, 0xe5, 0xe6, 0xbb, 0xea, 0xdf, 0x3f, 0x82, 0xbe, 0xf5, 0x2e, 0xcc, 0xee, 0x2c, 0xf4,
	0xd9, 0x96, 0x8e, 0x48, 0xb1, 0x55, 0x3c, 0x30, 0xab, 0x25, 0x5b, 0x45, 0x7f, 0xf6, 0x61, 0xd2,
	0xf9, 0x34, 0x24, 0xaf, 0x42, 0x90, 0x1f, 0x51, 0xfc, 0x52, 0x65, 0x32, 0x0f, 0x0f, 0x31, 0x75,
	0xc8, 0x8f, 0x96, 0x0e, 0x31, 0xb7, 0x5d, 0x6b, 0x66, 0xdc, 0xb3, 0xb7, 0x5d, 0xbe, 0x19, 0x16,
	0x77, 0x60, 0x9c, 0xc9, 0x8a, 0x8b, 0x75, 0x69, 0x09, 0x47, 0x76, 0x5a, 0x39, 0x0c, 0x29, 0xff,
	0x9c, 0xca, 0xf7, 0xff, 0x57, 0x53, 0xf9, 0x17, 0x0f, 0x66, 0x4d, 0x0e, 0x8d, 0xc2, 0xe3, 0x2d,
	0x2b, 0x9c, 0x3a, 0x41, 0x8d, 0xc6, 0xef, 0x3d, 0xf0, 0xad, 0xfb, 0xc2, 0x77, 0x51, 0xda, 0x57,
	0x5b, 0x68, 0xf9, 0x4b, 0xf6, 0x8f, 0x9d, 0xa2, 0x95, 0x8f, 0xf6, 0xb9, 0xf7, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0x75, 0x58, 0x30, 0xff, 0x0c, 0x00, 0x00,
}