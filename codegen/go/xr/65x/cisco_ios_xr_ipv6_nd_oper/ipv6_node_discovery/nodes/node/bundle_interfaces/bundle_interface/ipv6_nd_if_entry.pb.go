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
// source: ipv6_nd_if_entry.proto

package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_bundle_interfaces_bundle_interface

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

type Ipv6NdIfEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdIfEntry_KEYS) Reset()         { *m = Ipv6NdIfEntry_KEYS{} }
func (m *Ipv6NdIfEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfEntry_KEYS) ProtoMessage()    {}
func (*Ipv6NdIfEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_abee8767647ed8d3, []int{0}
}

func (m *Ipv6NdIfEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfEntry_KEYS.Merge(m, src)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfEntry_KEYS.Size(m)
}
func (m *Ipv6NdIfEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfEntry_KEYS proto.InternalMessageInfo

func (m *Ipv6NdIfEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdIfEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6NdIfParams struct {
	IsDadEnabled            bool     `protobuf:"varint,1,opt,name=is_dad_enabled,json=isDadEnabled,proto3" json:"is_dad_enabled,omitempty"`
	DadAttempts             uint32   `protobuf:"varint,2,opt,name=dad_attempts,json=dadAttempts,proto3" json:"dad_attempts,omitempty"`
	IsIcmPv6Redirect        bool     `protobuf:"varint,3,opt,name=is_icm_pv6_redirect,json=isIcmPv6Redirect,proto3" json:"is_icm_pv6_redirect,omitempty"`
	IsDhcpManaged           bool     `protobuf:"varint,4,opt,name=is_dhcp_managed,json=isDhcpManaged,proto3" json:"is_dhcp_managed,omitempty"`
	IsRouteAddressManaged   bool     `protobuf:"varint,5,opt,name=is_route_address_managed,json=isRouteAddressManaged,proto3" json:"is_route_address_managed,omitempty"`
	IsSuppressed            bool     `protobuf:"varint,6,opt,name=is_suppressed,json=isSuppressed,proto3" json:"is_suppressed,omitempty"`
	SendUnicastRa           bool     `protobuf:"varint,7,opt,name=send_unicast_ra,json=sendUnicastRa,proto3" json:"send_unicast_ra,omitempty"`
	NdRetransmitInterval    uint32   `protobuf:"varint,8,opt,name=nd_retransmit_interval,json=ndRetransmitInterval,proto3" json:"nd_retransmit_interval,omitempty"`
	NdMinTransmitInterval   uint32   `protobuf:"varint,9,opt,name=nd_min_transmit_interval,json=ndMinTransmitInterval,proto3" json:"nd_min_transmit_interval,omitempty"`
	NdMaxTransmitInterval   uint32   `protobuf:"varint,10,opt,name=nd_max_transmit_interval,json=ndMaxTransmitInterval,proto3" json:"nd_max_transmit_interval,omitempty"`
	NdAdvertisementLifetime uint32   `protobuf:"varint,11,opt,name=nd_advertisement_lifetime,json=ndAdvertisementLifetime,proto3" json:"nd_advertisement_lifetime,omitempty"`
	NdReachableTime         uint32   `protobuf:"varint,12,opt,name=nd_reachable_time,json=ndReachableTime,proto3" json:"nd_reachable_time,omitempty"`
	NdCacheLimit            uint32   `protobuf:"varint,13,opt,name=nd_cache_limit,json=ndCacheLimit,proto3" json:"nd_cache_limit,omitempty"`
	CompleteProtocolCount   uint32   `protobuf:"varint,14,opt,name=complete_protocol_count,json=completeProtocolCount,proto3" json:"complete_protocol_count,omitempty"`
	CompleteGleanCount      uint32   `protobuf:"varint,15,opt,name=complete_glean_count,json=completeGleanCount,proto3" json:"complete_glean_count,omitempty"`
	IncompleteProtocolCount uint32   `protobuf:"varint,16,opt,name=incomplete_protocol_count,json=incompleteProtocolCount,proto3" json:"incomplete_protocol_count,omitempty"`
	IncompleteGleanCount    uint32   `protobuf:"varint,17,opt,name=incomplete_glean_count,json=incompleteGleanCount,proto3" json:"incomplete_glean_count,omitempty"`
	DroppedProtocolReqCount uint32   `protobuf:"varint,18,opt,name=dropped_protocol_req_count,json=droppedProtocolReqCount,proto3" json:"dropped_protocol_req_count,omitempty"`
	DroppedGleanReqCount    uint32   `protobuf:"varint,19,opt,name=dropped_glean_req_count,json=droppedGleanReqCount,proto3" json:"dropped_glean_req_count,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ipv6NdIfParams) Reset()         { *m = Ipv6NdIfParams{} }
func (m *Ipv6NdIfParams) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfParams) ProtoMessage()    {}
func (*Ipv6NdIfParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_abee8767647ed8d3, []int{1}
}

func (m *Ipv6NdIfParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfParams.Unmarshal(m, b)
}
func (m *Ipv6NdIfParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfParams.Marshal(b, m, deterministic)
}
func (m *Ipv6NdIfParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfParams.Merge(m, src)
}
func (m *Ipv6NdIfParams) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfParams.Size(m)
}
func (m *Ipv6NdIfParams) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfParams.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfParams proto.InternalMessageInfo

func (m *Ipv6NdIfParams) GetIsDadEnabled() bool {
	if m != nil {
		return m.IsDadEnabled
	}
	return false
}

func (m *Ipv6NdIfParams) GetDadAttempts() uint32 {
	if m != nil {
		return m.DadAttempts
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIsIcmPv6Redirect() bool {
	if m != nil {
		return m.IsIcmPv6Redirect
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsDhcpManaged() bool {
	if m != nil {
		return m.IsDhcpManaged
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsRouteAddressManaged() bool {
	if m != nil {
		return m.IsRouteAddressManaged
	}
	return false
}

func (m *Ipv6NdIfParams) GetIsSuppressed() bool {
	if m != nil {
		return m.IsSuppressed
	}
	return false
}

func (m *Ipv6NdIfParams) GetSendUnicastRa() bool {
	if m != nil {
		return m.SendUnicastRa
	}
	return false
}

func (m *Ipv6NdIfParams) GetNdRetransmitInterval() uint32 {
	if m != nil {
		return m.NdRetransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdMinTransmitInterval() uint32 {
	if m != nil {
		return m.NdMinTransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdMaxTransmitInterval() uint32 {
	if m != nil {
		return m.NdMaxTransmitInterval
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdAdvertisementLifetime() uint32 {
	if m != nil {
		return m.NdAdvertisementLifetime
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdReachableTime() uint32 {
	if m != nil {
		return m.NdReachableTime
	}
	return 0
}

func (m *Ipv6NdIfParams) GetNdCacheLimit() uint32 {
	if m != nil {
		return m.NdCacheLimit
	}
	return 0
}

func (m *Ipv6NdIfParams) GetCompleteProtocolCount() uint32 {
	if m != nil {
		return m.CompleteProtocolCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetCompleteGleanCount() uint32 {
	if m != nil {
		return m.CompleteGleanCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIncompleteProtocolCount() uint32 {
	if m != nil {
		return m.IncompleteProtocolCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetIncompleteGleanCount() uint32 {
	if m != nil {
		return m.IncompleteGleanCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetDroppedProtocolReqCount() uint32 {
	if m != nil {
		return m.DroppedProtocolReqCount
	}
	return 0
}

func (m *Ipv6NdIfParams) GetDroppedGleanReqCount() uint32 {
	if m != nil {
		return m.DroppedGleanReqCount
	}
	return 0
}

type Ipv6NdAddr struct {
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	ValidLifetime        uint32   `protobuf:"varint,2,opt,name=valid_lifetime,json=validLifetime,proto3" json:"valid_lifetime,omitempty"`
	PrefLifetime         uint32   `protobuf:"varint,3,opt,name=pref_lifetime,json=prefLifetime,proto3" json:"pref_lifetime,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	Flags                uint32   `protobuf:"varint,5,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdAddr) Reset()         { *m = Ipv6NdAddr{} }
func (m *Ipv6NdAddr) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdAddr) ProtoMessage()    {}
func (*Ipv6NdAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_abee8767647ed8d3, []int{2}
}

func (m *Ipv6NdAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdAddr.Unmarshal(m, b)
}
func (m *Ipv6NdAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdAddr.Marshal(b, m, deterministic)
}
func (m *Ipv6NdAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdAddr.Merge(m, src)
}
func (m *Ipv6NdAddr) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdAddr.Size(m)
}
func (m *Ipv6NdAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdAddr.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdAddr proto.InternalMessageInfo

func (m *Ipv6NdAddr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *Ipv6NdAddr) GetValidLifetime() uint32 {
	if m != nil {
		return m.ValidLifetime
	}
	return 0
}

func (m *Ipv6NdAddr) GetPrefLifetime() uint32 {
	if m != nil {
		return m.PrefLifetime
	}
	return 0
}

func (m *Ipv6NdAddr) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv6NdAddr) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type Ipv6NdGspnode struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	TotalLinks           uint32   `protobuf:"varint,2,opt,name=total_links,json=totalLinks,proto3" json:"total_links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdGspnode) Reset()         { *m = Ipv6NdGspnode{} }
func (m *Ipv6NdGspnode) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdGspnode) ProtoMessage()    {}
func (*Ipv6NdGspnode) Descriptor() ([]byte, []int) {
	return fileDescriptor_abee8767647ed8d3, []int{3}
}

func (m *Ipv6NdGspnode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdGspnode.Unmarshal(m, b)
}
func (m *Ipv6NdGspnode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdGspnode.Marshal(b, m, deterministic)
}
func (m *Ipv6NdGspnode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdGspnode.Merge(m, src)
}
func (m *Ipv6NdGspnode) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdGspnode.Size(m)
}
func (m *Ipv6NdGspnode) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdGspnode.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdGspnode proto.InternalMessageInfo

func (m *Ipv6NdGspnode) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdGspnode) GetTotalLinks() uint32 {
	if m != nil {
		return m.TotalLinks
	}
	return 0
}

type Ipv6NdIfEntry struct {
	ParentInterfaceName  string           `protobuf:"bytes,50,opt,name=parent_interface_name,json=parentInterfaceName,proto3" json:"parent_interface_name,omitempty"`
	Iftype               uint32           `protobuf:"varint,51,opt,name=iftype,proto3" json:"iftype,omitempty"`
	Mtu                  uint32           `protobuf:"varint,52,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Etype                uint32           `protobuf:"varint,53,opt,name=etype,proto3" json:"etype,omitempty"`
	VlanTag              uint32           `protobuf:"varint,54,opt,name=vlan_tag,json=vlanTag,proto3" json:"vlan_tag,omitempty"`
	MacAddrSize          uint32           `protobuf:"varint,55,opt,name=mac_addr_size,json=macAddrSize,proto3" json:"mac_addr_size,omitempty"`
	MacAddr              string           `protobuf:"bytes,56,opt,name=mac_addr,json=macAddr,proto3" json:"mac_addr,omitempty"`
	IsInterfaceEnabled   bool             `protobuf:"varint,57,opt,name=is_interface_enabled,json=isInterfaceEnabled,proto3" json:"is_interface_enabled,omitempty"`
	IsIpv6Enabled        bool             `protobuf:"varint,58,opt,name=is_ipv6_enabled,json=isIpv6Enabled,proto3" json:"is_ipv6_enabled,omitempty"`
	IsMplsEnabled        bool             `protobuf:"varint,59,opt,name=is_mpls_enabled,json=isMplsEnabled,proto3" json:"is_mpls_enabled,omitempty"`
	NdParameters         *Ipv6NdIfParams  `protobuf:"bytes,60,opt,name=nd_parameters,json=ndParameters,proto3" json:"nd_parameters,omitempty"`
	GlobalAddress        []*Ipv6NdAddr    `protobuf:"bytes,61,rep,name=global_address,json=globalAddress,proto3" json:"global_address,omitempty"`
	LocalAddress         *Ipv6NdAddr      `protobuf:"bytes,62,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	MemberLink           []uint32         `protobuf:"varint,63,rep,packed,name=member_link,json=memberLink,proto3" json:"member_link,omitempty"`
	MemberNode           []*Ipv6NdGspnode `protobuf:"bytes,64,rep,name=member_node,json=memberNode,proto3" json:"member_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Ipv6NdIfEntry) Reset()         { *m = Ipv6NdIfEntry{} }
func (m *Ipv6NdIfEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdIfEntry) ProtoMessage()    {}
func (*Ipv6NdIfEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_abee8767647ed8d3, []int{4}
}

func (m *Ipv6NdIfEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdIfEntry.Unmarshal(m, b)
}
func (m *Ipv6NdIfEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdIfEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6NdIfEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdIfEntry.Merge(m, src)
}
func (m *Ipv6NdIfEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdIfEntry.Size(m)
}
func (m *Ipv6NdIfEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdIfEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdIfEntry proto.InternalMessageInfo

func (m *Ipv6NdIfEntry) GetParentInterfaceName() string {
	if m != nil {
		return m.ParentInterfaceName
	}
	return ""
}

func (m *Ipv6NdIfEntry) GetIftype() uint32 {
	if m != nil {
		return m.Iftype
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetEtype() uint32 {
	if m != nil {
		return m.Etype
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetVlanTag() uint32 {
	if m != nil {
		return m.VlanTag
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMacAddrSize() uint32 {
	if m != nil {
		return m.MacAddrSize
	}
	return 0
}

func (m *Ipv6NdIfEntry) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *Ipv6NdIfEntry) GetIsInterfaceEnabled() bool {
	if m != nil {
		return m.IsInterfaceEnabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetIsIpv6Enabled() bool {
	if m != nil {
		return m.IsIpv6Enabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetIsMplsEnabled() bool {
	if m != nil {
		return m.IsMplsEnabled
	}
	return false
}

func (m *Ipv6NdIfEntry) GetNdParameters() *Ipv6NdIfParams {
	if m != nil {
		return m.NdParameters
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetGlobalAddress() []*Ipv6NdAddr {
	if m != nil {
		return m.GlobalAddress
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetLocalAddress() *Ipv6NdAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetMemberLink() []uint32 {
	if m != nil {
		return m.MemberLink
	}
	return nil
}

func (m *Ipv6NdIfEntry) GetMemberNode() []*Ipv6NdGspnode {
	if m != nil {
		return m.MemberNode
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv6NdIfEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_entry_KEYS")
	proto.RegisterType((*Ipv6NdIfParams)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_params")
	proto.RegisterType((*Ipv6NdAddr)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_addr")
	proto.RegisterType((*Ipv6NdGspnode)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_gspnode")
	proto.RegisterType((*Ipv6NdIfEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_interfaces.bundle_interface.ipv6_nd_if_entry")
}

func init() { proto.RegisterFile("ipv6_nd_if_entry.proto", fileDescriptor_abee8767647ed8d3) }

var fileDescriptor_abee8767647ed8d3 = []byte{
	// 988 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x5d, 0x6f, 0x1c, 0x35,
	0x14, 0xd5, 0x12, 0xda, 0x24, 0xde, 0x9d, 0x7c, 0x38, 0x1f, 0x9d, 0xc0, 0x43, 0xc3, 0xb6, 0xa0,
	0x08, 0x89, 0x15, 0x4a, 0xd3, 0x14, 0x52, 0xbe, 0xa2, 0xb6, 0x42, 0x11, 0x49, 0x89, 0x26, 0xe1,
	0x01, 0xf5, 0xc1, 0xf2, 0x8e, 0xef, 0xee, 0x5a, 0x78, 0x3c, 0x53, 0xdb, 0xbb, 0xda, 0xf4, 0x3f,
	0x80, 0x90, 0xf8, 0x33, 0xbc, 0xf1, 0x0b, 0xf8, 0x4f, 0xc8, 0xd7, 0xe3, 0xd9, 0x6d, 0x42, 0x79,
	0x23, 0x2f, 0xab, 0xf5, 0xb9, 0xe7, 0xf8, 0x5c, 0x8f, 0xe7, 0x9e, 0x21, 0xdb, 0xb2, 0x9a, 0x1c,
	0x32, 0x2d, 0x98, 0x1c, 0x30, 0xd0, 0xce, 0x5c, 0xf5, 0x2a, 0x53, 0xba, 0x92, 0xbe, 0xca, 0xa5,
	0xcd, 0x4b, 0x26, 0x4b, 0xcb, 0xa6, 0x86, 0x45, 0x52, 0x59, 0x81, 0xe9, 0x85, 0x45, 0x29, 0x80,
	0x09, 0xcf, 0x99, 0x80, 0xb9, 0xea, 0xf9, 0xa5, 0xc5, 0xdf, 0x5e, 0x7f, 0xac, 0x85, 0x02, 0x26,
	0xb5, 0x03, 0x33, 0xe0, 0x39, 0xd8, 0x1b, 0x48, 0xf7, 0x15, 0xd9, 0xba, 0x6e, 0xcb, 0x7e, 0x78,
	0xf1, 0xf3, 0x05, 0xfd, 0x90, 0x2c, 0xe3, 0xc6, 0x9a, 0x17, 0x90, 0xb6, 0x76, 0x5b, 0x7b, 0xcb,
	0xd9, 0x92, 0x07, 0x5e, 0xf2, 0x02, 0xe8, 0xc7, 0x64, 0xa5, 0xd9, 0x22, 0x30, 0xde, 0x43, 0x46,
	0xd2, 0xa0, 0x9e, 0xd6, 0xfd, 0x6b, 0x91, 0xac, 0xcf, 0xed, 0x5e, 0x71, 0xc3, 0x0b, 0x4b, 0x1f,
	0x92, 0x15, 0x69, 0x99, 0xe0, 0x82, 0x81, 0xe6, 0x7d, 0x05, 0x02, 0xb7, 0x5f, 0xca, 0x3a, 0xd2,
	0x3e, 0xe7, 0xe2, 0x45, 0xc0, 0xe8, 0x47, 0xa4, 0xe3, 0x29, 0xdc, 0x39, 0x28, 0x2a, 0x67, 0xd1,
	0x20, 0xc9, 0xda, 0x82, 0x8b, 0xe3, 0x1a, 0xa2, 0x9f, 0x91, 0x0d, 0x69, 0x99, 0xcc, 0x0b, 0xe6,
	0x3d, 0x0c, 0x08, 0x69, 0x20, 0x77, 0xe9, 0x02, 0xee, 0xb6, 0x26, 0xed, 0x49, 0x5e, 0x9c, 0x4f,
	0x0e, 0xb3, 0x1a, 0xa7, 0x9f, 0x90, 0x55, 0xef, 0x3b, 0xca, 0x2b, 0x56, 0x70, 0xcd, 0x87, 0x20,
	0xd2, 0xf7, 0x91, 0x9a, 0x48, 0xfb, 0x7c, 0x94, 0x57, 0x67, 0x01, 0xa4, 0x4f, 0x48, 0x2a, 0x2d,
	0x33, 0xe5, 0xd8, 0x01, 0xe3, 0x42, 0x18, 0xb0, 0xb6, 0x11, 0xdc, 0x41, 0xc1, 0x96, 0xb4, 0x99,
	0x2f, 0x1f, 0x87, 0x6a, 0x14, 0x3e, 0x20, 0x89, 0xb4, 0xcc, 0x8e, 0xab, 0xca, 0xa3, 0x20, 0xd2,
	0xbb, 0xf1, 0x5c, 0x17, 0x0d, 0xe6, 0xbb, 0xb0, 0xa0, 0x05, 0x1b, 0x6b, 0x99, 0x73, 0xeb, 0x98,
	0xe1, 0xe9, 0x62, 0xe8, 0xc2, 0xc3, 0x3f, 0x05, 0x34, 0xe3, 0xf4, 0x80, 0x6c, 0x6b, 0xc1, 0x0c,
	0x38, 0xc3, 0xb5, 0x2d, 0xa4, 0x0b, 0x77, 0x36, 0xe1, 0x2a, 0x5d, 0xc2, 0x27, 0xb1, 0xa9, 0x45,
	0xd6, 0x14, 0x4f, 0xea, 0x9a, 0xef, 0x5d, 0x0b, 0x56, 0x48, 0xcd, 0x6e, 0xea, 0x96, 0x51, 0xb7,
	0xa5, 0xc5, 0x99, 0xd4, 0x97, 0xef, 0x10, 0xf2, 0xe9, 0xbf, 0x08, 0x49, 0x23, 0xe4, 0xd3, 0x1b,
	0xc2, 0x23, 0xb2, 0xa3, 0x05, 0xe3, 0x62, 0x02, 0xc6, 0x49, 0x0b, 0x05, 0x68, 0xc7, 0x94, 0x1c,
	0x80, 0x93, 0x05, 0xa4, 0x6d, 0x54, 0xde, 0xd3, 0xe2, 0x78, 0xbe, 0x7e, 0x5a, 0x97, 0xe9, 0xa7,
	0x64, 0x1d, 0xcf, 0xc8, 0xf3, 0x91, 0xbf, 0x74, 0x86, 0x9a, 0x0e, 0x6a, 0x56, 0xfd, 0xf1, 0x6a,
	0xfc, 0xd2, 0x73, 0x1f, 0x92, 0x15, 0x2d, 0x58, 0xce, 0xf3, 0x11, 0x30, 0x25, 0x0b, 0xe9, 0xd2,
	0x04, 0x89, 0x1d, 0x2d, 0x9e, 0x79, 0xf0, 0xd4, 0x63, 0xf4, 0x90, 0xdc, 0xcb, 0xcb, 0xa2, 0x52,
	0xe0, 0x80, 0xe1, 0xf4, 0xe4, 0xa5, 0x62, 0x79, 0x39, 0xd6, 0x2e, 0x5d, 0x09, 0xa7, 0x88, 0xe5,
	0xf3, 0xba, 0xfa, 0xcc, 0x17, 0xe9, 0xe7, 0x64, 0xb3, 0xd1, 0x0d, 0x15, 0x70, 0x5d, 0x8b, 0x56,
	0x51, 0x44, 0x63, 0xed, 0x7b, 0x5f, 0x0a, 0x8a, 0x23, 0xb2, 0x23, 0xf5, 0xbb, 0xbc, 0xd6, 0xc2,
	0xb9, 0x67, 0x84, 0xb7, 0xdd, 0x0e, 0xc8, 0xf6, 0x9c, 0x76, 0xde, 0x6f, 0x3d, 0xdc, 0xed, 0xac,
	0x3a, 0xe7, 0xf8, 0x94, 0x7c, 0x20, 0x4c, 0x59, 0x55, 0x20, 0x66, 0x76, 0x06, 0x5e, 0xd7, 0x4a,
	0x1a, 0x2c, 0x6b, 0x46, 0xf4, 0xcb, 0xe0, 0x75, 0x10, 0x3f, 0x26, 0xb1, 0x54, 0xfb, 0xcd, 0x94,
	0x1b, 0xc1, 0xb3, 0x2e, 0xa3, 0x61, 0x94, 0x75, 0xff, 0x6c, 0x91, 0x4e, 0x9c, 0x60, 0x3f, 0x0b,
	0x7e, 0x2c, 0x71, 0x5d, 0x0f, 0x46, 0x9d, 0x0c, 0x6d, 0x8f, 0xd5, 0xd3, 0xe0, 0xc3, 0x61, 0xc2,
	0x95, 0x14, 0xb3, 0xd7, 0x20, 0xcc, 0x6e, 0x82, 0x68, 0x73, 0xf9, 0x0f, 0x48, 0x52, 0x19, 0x18,
	0xcc, 0x58, 0x0b, 0xe1, 0x3e, 0x3d, 0x78, 0x9d, 0x24, 0xa7, 0x4c, 0x81, 0x1e, 0xba, 0x11, 0x4e,
	0x6c, 0x4d, 0x92, 0xd3, 0x53, 0xc4, 0xe8, 0x26, 0xb9, 0x33, 0x50, 0x7c, 0x68, 0x71, 0x3a, 0x93,
	0x2c, 0x2c, 0xba, 0x3f, 0x92, 0xd5, 0xd8, 0xf9, 0xd0, 0x56, 0x3e, 0xba, 0xfe, 0x3b, 0xd3, 0xee,
	0x93, 0xb6, 0x2b, 0x1d, 0x57, 0x4c, 0x49, 0xfd, 0x4b, 0xcc, 0x1b, 0x82, 0xd0, 0xa9, 0x47, 0xba,
	0x7f, 0x2f, 0x92, 0xb5, 0xeb, 0x59, 0x49, 0xf7, 0xc9, 0x56, 0xc5, 0x8d, 0x7f, 0xe9, 0xaf, 0x05,
	0xe2, 0x3e, 0x6e, 0xbf, 0x11, 0x8a, 0x27, 0xf3, 0xb1, 0x48, 0xb7, 0xc9, 0x5d, 0x39, 0x70, 0x57,
	0x15, 0xa4, 0x8f, 0xd0, 0xa4, 0x5e, 0xd1, 0x35, 0xb2, 0x50, 0xb8, 0x71, 0x7a, 0x80, 0xa0, 0xff,
	0xeb, 0x4f, 0x06, 0x48, 0x7c, 0x1c, 0x4e, 0x86, 0x0b, 0xba, 0x43, 0x96, 0x26, 0x8a, 0x6b, 0xe6,
	0xf8, 0x30, 0x3d, 0xc4, 0xc2, 0xa2, 0x5f, 0x5f, 0xf2, 0x21, 0xed, 0x92, 0xa4, 0xe0, 0x39, 0xde,
	0x0e, 0xb3, 0xf2, 0x0d, 0xa4, 0x4f, 0x42, 0x6c, 0x16, 0x3c, 0xf7, 0xd7, 0x73, 0x21, 0xdf, 0xa0,
	0x3c, 0x72, 0xd2, 0x2f, 0xb0, 0xcb, 0xc5, 0xba, 0xec, 0xc7, 0xc0, 0x27, 0x6a, 0x73, 0x92, 0x18,
	0xd0, 0x5f, 0x62, 0x42, 0x51, 0x69, 0x9b, 0x83, 0xc4, 0x98, 0x0e, 0xa1, 0x8a, 0x8f, 0x25, 0x92,
	0x8f, 0x62, 0xa8, 0x9e, 0x54, 0x93, 0xc3, 0xb7, 0x79, 0x45, 0xa5, 0x6c, 0xc3, 0x7b, 0x1a, 0x79,
	0x67, 0x95, 0xb2, 0x91, 0xf7, 0x47, 0x8b, 0x24, 0x5a, 0x84, 0x4f, 0x05, 0x38, 0x30, 0x36, 0xfd,
	0x6a, 0xb7, 0xb5, 0xd7, 0xde, 0xd7, 0xbd, 0xff, 0xf1, 0x2b, 0xd8, 0xbb, 0xf1, 0x91, 0xf2, 0xb1,
	0x72, 0xde, 0xf4, 0x40, 0x7f, 0x6f, 0x91, 0x95, 0xa1, 0x2a, 0xfb, 0x5c, 0x35, 0x2f, 0xfe, 0xd7,
	0xbb, 0x0b, 0x7b, 0xed, 0x7d, 0x79, 0x2b, 0x6d, 0x79, 0xcf, 0x2c, 0x09, 0x0d, 0xc4, 0x29, 0xfb,
	0xad, 0x45, 0x12, 0x55, 0xe6, 0x73, 0x1d, 0x7d, 0x83, 0x0f, 0xea, 0x16, 0x3b, 0xea, 0xa0, 0x7f,
	0x6c, 0xe8, 0x3e, 0x69, 0x17, 0x50, 0xf4, 0xc1, 0xe0, 0x00, 0xa5, 0xdf, 0xee, 0x2e, 0xf8, 0xf9,
	0x09, 0x90, 0x1f, 0x20, 0xfa, 0x6b, 0xab, 0x61, 0x78, 0xbb, 0xf4, 0x3b, 0x7c, 0x82, 0xea, 0x56,
	0xfa, 0xad, 0x13, 0x20, 0xf6, 0xf3, 0xb2, 0x14, 0xd0, 0xbf, 0x8b, 0x29, 0xfa, 0xe8, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x0d, 0x6a, 0x36, 0x78, 0x09, 0x00, 0x00,
}
