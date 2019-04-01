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
// source: ptp_if_unicast_peer_info.proto

package cisco_ios_xr_ptp_oper_ptp_nodes_node_node_interface_unicast_peers_node_interface_unicast_peer

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

type PtpIfUnicastPeerInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfUnicastPeerInfo_KEYS) Reset()         { *m = PtpIfUnicastPeerInfo_KEYS{} }
func (m *PtpIfUnicastPeerInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfUnicastPeerInfo_KEYS) ProtoMessage()    {}
func (*PtpIfUnicastPeerInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9d6fe8454e900c3, []int{0}
}

func (m *PtpIfUnicastPeerInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpIfUnicastPeerInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfUnicastPeerInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS.Merge(m, src)
}
func (m *PtpIfUnicastPeerInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS.Size(m)
}
func (m *PtpIfUnicastPeerInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfUnicastPeerInfo_KEYS proto.InternalMessageInfo

func (m *PtpIfUnicastPeerInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PtpIfUnicastPeerInfo_KEYS) GetInterfaceName() string {
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
	return fileDescriptor_e9d6fe8454e900c3, []int{1}
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
	return fileDescriptor_e9d6fe8454e900c3, []int{2}
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
	return fileDescriptor_e9d6fe8454e900c3, []int{3}
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
	return fileDescriptor_e9d6fe8454e900c3, []int{4}
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

type PtpBagUnicastPeer struct {
	Address              *PtpBagAddress      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AnnounceGrant        *PtpBagUnicastGrant `protobuf:"bytes,2,opt,name=announce_grant,json=announceGrant,proto3" json:"announce_grant,omitempty"`
	SyncGrant            *PtpBagUnicastGrant `protobuf:"bytes,3,opt,name=sync_grant,json=syncGrant,proto3" json:"sync_grant,omitempty"`
	DelayResponseGrant   *PtpBagUnicastGrant `protobuf:"bytes,4,opt,name=delay_response_grant,json=delayResponseGrant,proto3" json:"delay_response_grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PtpBagUnicastPeer) Reset()         { *m = PtpBagUnicastPeer{} }
func (m *PtpBagUnicastPeer) String() string { return proto.CompactTextString(m) }
func (*PtpBagUnicastPeer) ProtoMessage()    {}
func (*PtpBagUnicastPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9d6fe8454e900c3, []int{5}
}

func (m *PtpBagUnicastPeer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagUnicastPeer.Unmarshal(m, b)
}
func (m *PtpBagUnicastPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagUnicastPeer.Marshal(b, m, deterministic)
}
func (m *PtpBagUnicastPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagUnicastPeer.Merge(m, src)
}
func (m *PtpBagUnicastPeer) XXX_Size() int {
	return xxx_messageInfo_PtpBagUnicastPeer.Size(m)
}
func (m *PtpBagUnicastPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagUnicastPeer.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagUnicastPeer proto.InternalMessageInfo

func (m *PtpBagUnicastPeer) GetAddress() *PtpBagAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PtpBagUnicastPeer) GetAnnounceGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.AnnounceGrant
	}
	return nil
}

func (m *PtpBagUnicastPeer) GetSyncGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.SyncGrant
	}
	return nil
}

func (m *PtpBagUnicastPeer) GetDelayResponseGrant() *PtpBagUnicastGrant {
	if m != nil {
		return m.DelayResponseGrant
	}
	return nil
}

type PtpIfUnicastPeerInfo struct {
	Name                 string               `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	PortNumber           uint32               `protobuf:"varint,51,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	Peers                []*PtpBagUnicastPeer `protobuf:"bytes,52,rep,name=peers,proto3" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PtpIfUnicastPeerInfo) Reset()         { *m = PtpIfUnicastPeerInfo{} }
func (m *PtpIfUnicastPeerInfo) String() string { return proto.CompactTextString(m) }
func (*PtpIfUnicastPeerInfo) ProtoMessage()    {}
func (*PtpIfUnicastPeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9d6fe8454e900c3, []int{6}
}

func (m *PtpIfUnicastPeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfUnicastPeerInfo.Unmarshal(m, b)
}
func (m *PtpIfUnicastPeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfUnicastPeerInfo.Marshal(b, m, deterministic)
}
func (m *PtpIfUnicastPeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfUnicastPeerInfo.Merge(m, src)
}
func (m *PtpIfUnicastPeerInfo) XXX_Size() int {
	return xxx_messageInfo_PtpIfUnicastPeerInfo.Size(m)
}
func (m *PtpIfUnicastPeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfUnicastPeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfUnicastPeerInfo proto.InternalMessageInfo

func (m *PtpIfUnicastPeerInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PtpIfUnicastPeerInfo) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

func (m *PtpIfUnicastPeerInfo) GetPeers() []*PtpBagUnicastPeer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpIfUnicastPeerInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_if_unicast_peer_info_KEYS")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_bag_address")
	proto.RegisterType((*PtpBagUnicastGrant)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_bag_unicast_grant")
	proto.RegisterType((*PtpBagUnicastPeer)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_bag_unicast_peer")
	proto.RegisterType((*PtpIfUnicastPeerInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.node_interface_unicast_peers.node_interface_unicast_peer.ptp_if_unicast_peer_info")
}

func init() { proto.RegisterFile("ptp_if_unicast_peer_info.proto", fileDescriptor_e9d6fe8454e900c3) }

var fileDescriptor_e9d6fe8454e900c3 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x95, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x89, 0x33, 0xeb, 0xee, 0xbc, 0xb5, 0xb3, 0x12, 0x56, 0x29, 0x88, 0x3a, 0x16, 0xc5,
	0x39, 0x48, 0xc1, 0xd9, 0xc1, 0x8b, 0x27, 0x41, 0x11, 0x11, 0xf6, 0x50, 0xf1, 0xe0, 0x41, 0x42,
	0x26, 0xcd, 0x0c, 0xc5, 0x36, 0x09, 0x49, 0xba, 0x3a, 0xdf, 0x40, 0xcf, 0x22, 0x08, 0xde, 0xfd,
	0x08, 0x7e, 0x19, 0xbf, 0x8b, 0x48, 0xd2, 0x74, 0x76, 0xbb, 0xb8, 0xde, 0x76, 0xf6, 0x52, 0x26,
	0x4f, 0xdf, 0x27, 0xef, 0xef, 0xfd, 0x43, 0x07, 0xee, 0x28, 0xab, 0x48, 0xb9, 0x24, 0x8d, 0x28,
	0x19, 0x35, 0x96, 0x28, 0xce, 0x35, 0x29, 0xc5, 0x52, 0x66, 0x4a, 0x4b, 0x2b, 0xf1, 0x7b, 0x56,
	0x1a, 0x26, 0x49, 0x29, 0x0d, 0xf9, 0xa4, 0x89, 0x0b, 0x96, 0x8a, 0xeb, 0x4c, 0x59, 0x95, 0x09,
	0x59, 0x70, 0xe3, 0x9f, 0xfe, 0x41, 0x4a, 0x61, 0xb9, 0x5e, 0x52, 0xc6, 0x7b, 0x97, 0x99, 0xff,
	0xbd, 0x4c, 0x19, 0xdc, 0x3e, 0x0f, 0x80, 0xbc, 0x7e, 0xf1, 0xee, 0x0d, 0xbe, 0x05, 0x23, 0xef,
	0x17, 0xb4, 0xe6, 0x09, 0x9a, 0xa0, 0xe9, 0x28, 0xdf, 0x73, 0xc2, 0x11, 0xad, 0x39, 0x7e, 0x00,
	0xe3, 0x93, 0x7b, 0x7d, 0xc4, 0x15, 0x1f, 0x11, 0x6f, 0x54, 0x17, 0x96, 0x3e, 0x86, 0x1b, 0x2e,
	0xc9, 0x82, 0xae, 0x48, 0x4d, 0x19, 0xa1, 0x45, 0xa1, 0x89, 0x5d, 0x2b, 0x8e, 0x13, 0xd8, 0xad,
	0x29, 0x73, 0xe7, 0x70, 0x75, 0x77, 0x4c, 0x9f, 0xc2, 0xcd, 0xce, 0x52, 0xaa, 0xe3, 0x27, 0xa7,
	0x3c, 0xf7, 0xe0, 0xda, 0x46, 0xe1, 0xc6, 0x04, 0x63, 0xe4, 0xb4, 0x67, 0xad, 0x94, 0xfe, 0x1a,
	0xc0, 0x7e, 0xe7, 0x0e, 0x61, 0xf8, 0x3e, 0xc4, 0x5c, 0x30, 0xaa, 0x4c, 0x53, 0x51, 0x5b, 0x4a,
	0x11, 0x7c, 0x7d, 0x11, 0x3f, 0x84, 0xfd, 0x60, 0x20, 0x8d, 0xf8, 0x20, 0xe4, 0x47, 0xe1, 0x2b,
	0xda, 0xcb, 0xc7, 0x41, 0x7e, 0xdb, 0xaa, 0xf8, 0x1b, 0x82, 0xa8, 0xab, 0xc5, 0x51, 0x0c, 0x26,
	0x68, 0x1a, 0xcd, 0x6c, 0x76, 0xa1, 0xd3, 0xca, 0xfe, 0xd9, 0xc5, 0x1c, 0x6a, 0xca, 0x42, 0xe9,
	0xa1, 0x3b, 0xf3, 0x0d, 0xd7, 0x70, 0xd3, 0x9d, 0x79, 0x17, 0xf2, 0x1d, 0x9d, 0xe9, 0xe0, 0x8e,
	0x67, 0x6f, 0xb6, 0xc4, 0xde, 0x1f, 0x67, 0x7f, 0x70, 0xd5, 0xc9, 0xa2, 0x74, 0xde, 0x95, 0xa6,
	0xc2, 0xe2, 0x47, 0x80, 0x2b, 0xb9, 0x6a, 0x0f, 0x6d, 0x82, 0x63, 0x5a, 0x85, 0x11, 0x5e, 0xaf,
	0xe4, 0xea, 0xa5, 0x7b, 0xf1, 0x2a, 0xe8, 0x6e, 0x2d, 0xdb, 0xc8, 0xa2, 0xd1, 0xed, 0xb0, 0xdd,
	0x10, 0xe3, 0x3c, 0xf6, 0xea, 0xf3, 0x20, 0xa6, 0x7f, 0x86, 0x70, 0x70, 0x36, 0x9d, 0x43, 0xc5,
	0x9f, 0x11, 0xec, 0x9e, 0x5e, 0xaf, 0x68, 0x26, 0xb6, 0xd4, 0x9c, 0x90, 0x35, 0xef, 0xd2, 0xe3,
	0x1f, 0x08, 0xc6, 0x54, 0x08, 0xd9, 0x08, 0xc6, 0xdb, 0xf2, 0x7d, 0x2d, 0xdb, 0x5b, 0xb5, 0xde,
	0x1c, 0xf2, 0xb8, 0x63, 0xf1, 0x0d, 0xc7, 0x5f, 0x11, 0x80, 0x59, 0x0b, 0x16, 0xc8, 0x06, 0x97,
	0x48, 0x36, 0x72, 0x1c, 0x2d, 0xd5, 0x4f, 0x04, 0x07, 0x05, 0xaf, 0xe8, 0x9a, 0x68, 0x6e, 0x94,
	0x14, 0xa6, 0xeb, 0xdc, 0xf0, 0x12, 0xf9, 0xb0, 0x27, 0xca, 0x03, 0x90, 0x07, 0x4d, 0x7f, 0x23,
	0x48, 0xce, 0xfb, 0xfa, 0x62, 0x0c, 0x43, 0xff, 0x45, 0x9d, 0xf9, 0x25, 0xf7, 0xbf, 0xf1, 0x5d,
	0x88, 0x94, 0xd4, 0x96, 0x88, 0xa6, 0x5e, 0x70, 0x9d, 0x1c, 0xfa, 0xad, 0x06, 0x27, 0x1d, 0x79,
	0x05, 0x7f, 0x41, 0xb0, 0xe3, 0x09, 0x93, 0xf9, 0x64, 0x30, 0x8d, 0x66, 0x66, 0xcb, 0xb5, 0x3a,
	0x31, 0x6f, 0x09, 0x16, 0x57, 0xfd, 0x1f, 0xd8, 0xe1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x53, 0x97, 0xfb, 0xe2, 0x06, 0x00, 0x00,
}
