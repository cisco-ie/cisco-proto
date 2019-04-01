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
// source: ip_arp_statistics.proto

package cisco_ios_xr_ipv4_arp_oper_arp_nodes_node_traffic_interfaces_traffic_interface

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

type IpArpStatistics_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArpStatistics_KEYS) Reset()         { *m = IpArpStatistics_KEYS{} }
func (m *IpArpStatistics_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpArpStatistics_KEYS) ProtoMessage()    {}
func (*IpArpStatistics_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_394adc8d84b7147b, []int{0}
}

func (m *IpArpStatistics_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArpStatistics_KEYS.Unmarshal(m, b)
}
func (m *IpArpStatistics_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArpStatistics_KEYS.Marshal(b, m, deterministic)
}
func (m *IpArpStatistics_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArpStatistics_KEYS.Merge(m, src)
}
func (m *IpArpStatistics_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpArpStatistics_KEYS.Size(m)
}
func (m *IpArpStatistics_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArpStatistics_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpArpStatistics_KEYS proto.InternalMessageInfo

func (m *IpArpStatistics_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *IpArpStatistics_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IpArpStatistics struct {
	RequestsReceived              uint32   `protobuf:"varint,50,opt,name=requests_received,json=requestsReceived,proto3" json:"requests_received,omitempty"`
	RepliesReceived               uint32   `protobuf:"varint,51,opt,name=replies_received,json=repliesReceived,proto3" json:"replies_received,omitempty"`
	RequestsSent                  uint32   `protobuf:"varint,52,opt,name=requests_sent,json=requestsSent,proto3" json:"requests_sent,omitempty"`
	RepliesSent                   uint32   `protobuf:"varint,53,opt,name=replies_sent,json=repliesSent,proto3" json:"replies_sent,omitempty"`
	ProxyRepliesSent              uint32   `protobuf:"varint,54,opt,name=proxy_replies_sent,json=proxyRepliesSent,proto3" json:"proxy_replies_sent,omitempty"`
	SubscrRequestsReceived        uint32   `protobuf:"varint,55,opt,name=subscr_requests_received,json=subscrRequestsReceived,proto3" json:"subscr_requests_received,omitempty"`
	SubscrRepliesSent             uint32   `protobuf:"varint,56,opt,name=subscr_replies_sent,json=subscrRepliesSent,proto3" json:"subscr_replies_sent,omitempty"`
	SubscrRepliesGratgSent        uint32   `protobuf:"varint,57,opt,name=subscr_replies_gratg_sent,json=subscrRepliesGratgSent,proto3" json:"subscr_replies_gratg_sent,omitempty"`
	LocalProxyRepliesSent         uint32   `protobuf:"varint,58,opt,name=local_proxy_replies_sent,json=localProxyRepliesSent,proto3" json:"local_proxy_replies_sent,omitempty"`
	GratuitousRepliesSent         uint32   `protobuf:"varint,59,opt,name=gratuitous_replies_sent,json=gratuitousRepliesSent,proto3" json:"gratuitous_replies_sent,omitempty"`
	ResolutionRequestsReceived    uint32   `protobuf:"varint,60,opt,name=resolution_requests_received,json=resolutionRequestsReceived,proto3" json:"resolution_requests_received,omitempty"`
	ResolutionRepliesReceived     uint32   `protobuf:"varint,61,opt,name=resolution_replies_received,json=resolutionRepliesReceived,proto3" json:"resolution_replies_received,omitempty"`
	ResolutionRequestsDropped     uint32   `protobuf:"varint,62,opt,name=resolution_requests_dropped,json=resolutionRequestsDropped,proto3" json:"resolution_requests_dropped,omitempty"`
	OutOfMemoryErrors             uint32   `protobuf:"varint,63,opt,name=out_of_memory_errors,json=outOfMemoryErrors,proto3" json:"out_of_memory_errors,omitempty"`
	NoBufferErrors                uint32   `protobuf:"varint,64,opt,name=no_buffer_errors,json=noBufferErrors,proto3" json:"no_buffer_errors,omitempty"`
	TotalEntries                  uint32   `protobuf:"varint,65,opt,name=total_entries,json=totalEntries,proto3" json:"total_entries,omitempty"`
	DynamicEntries                uint32   `protobuf:"varint,66,opt,name=dynamic_entries,json=dynamicEntries,proto3" json:"dynamic_entries,omitempty"`
	StaticEntries                 uint32   `protobuf:"varint,67,opt,name=static_entries,json=staticEntries,proto3" json:"static_entries,omitempty"`
	AliasEntries                  uint32   `protobuf:"varint,68,opt,name=alias_entries,json=aliasEntries,proto3" json:"alias_entries,omitempty"`
	InterfaceEntries              uint32   `protobuf:"varint,69,opt,name=interface_entries,json=interfaceEntries,proto3" json:"interface_entries,omitempty"`
	StandbyEntries                uint32   `protobuf:"varint,70,opt,name=standby_entries,json=standbyEntries,proto3" json:"standby_entries,omitempty"`
	DhcpEntries                   uint32   `protobuf:"varint,71,opt,name=dhcp_entries,json=dhcpEntries,proto3" json:"dhcp_entries,omitempty"`
	VxlanEntries                  uint32   `protobuf:"varint,72,opt,name=vxlan_entries,json=vxlanEntries,proto3" json:"vxlan_entries,omitempty"`
	DropAdjacencyEntries          uint32   `protobuf:"varint,73,opt,name=drop_adjacency_entries,json=dropAdjacencyEntries,proto3" json:"drop_adjacency_entries,omitempty"`
	IpPacketsDroppedNode          uint32   `protobuf:"varint,74,opt,name=ip_packets_dropped_node,json=ipPacketsDroppedNode,proto3" json:"ip_packets_dropped_node,omitempty"`
	ArpPacketNodeOutOfSubnet      uint32   `protobuf:"varint,75,opt,name=arp_packet_node_out_of_subnet,json=arpPacketNodeOutOfSubnet,proto3" json:"arp_packet_node_out_of_subnet,omitempty"`
	IpPacketsDroppedInterface     uint32   `protobuf:"varint,76,opt,name=ip_packets_dropped_interface,json=ipPacketsDroppedInterface,proto3" json:"ip_packets_dropped_interface,omitempty"`
	ArpPacketInterfaceOutOfSubnet uint32   `protobuf:"varint,77,opt,name=arp_packet_interface_out_of_subnet,json=arpPacketInterfaceOutOfSubnet,proto3" json:"arp_packet_interface_out_of_subnet,omitempty"`
	ArpPacketUnsolicitedPacket    uint32   `protobuf:"varint,78,opt,name=arp_packet_unsolicited_packet,json=arpPacketUnsolicitedPacket,proto3" json:"arp_packet_unsolicited_packet,omitempty"`
	IdbStructures                 uint32   `protobuf:"varint,79,opt,name=idb_structures,json=idbStructures,proto3" json:"idb_structures,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *IpArpStatistics) Reset()         { *m = IpArpStatistics{} }
func (m *IpArpStatistics) String() string { return proto.CompactTextString(m) }
func (*IpArpStatistics) ProtoMessage()    {}
func (*IpArpStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_394adc8d84b7147b, []int{1}
}

func (m *IpArpStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArpStatistics.Unmarshal(m, b)
}
func (m *IpArpStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArpStatistics.Marshal(b, m, deterministic)
}
func (m *IpArpStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArpStatistics.Merge(m, src)
}
func (m *IpArpStatistics) XXX_Size() int {
	return xxx_messageInfo_IpArpStatistics.Size(m)
}
func (m *IpArpStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArpStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_IpArpStatistics proto.InternalMessageInfo

func (m *IpArpStatistics) GetRequestsReceived() uint32 {
	if m != nil {
		return m.RequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetRepliesReceived() uint32 {
	if m != nil {
		return m.RepliesReceived
	}
	return 0
}

func (m *IpArpStatistics) GetRequestsSent() uint32 {
	if m != nil {
		return m.RequestsSent
	}
	return 0
}

func (m *IpArpStatistics) GetRepliesSent() uint32 {
	if m != nil {
		return m.RepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetProxyRepliesSent() uint32 {
	if m != nil {
		return m.ProxyRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRequestsReceived() uint32 {
	if m != nil {
		return m.SubscrRequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRepliesSent() uint32 {
	if m != nil {
		return m.SubscrRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetSubscrRepliesGratgSent() uint32 {
	if m != nil {
		return m.SubscrRepliesGratgSent
	}
	return 0
}

func (m *IpArpStatistics) GetLocalProxyRepliesSent() uint32 {
	if m != nil {
		return m.LocalProxyRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetGratuitousRepliesSent() uint32 {
	if m != nil {
		return m.GratuitousRepliesSent
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRequestsReceived() uint32 {
	if m != nil {
		return m.ResolutionRequestsReceived
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRepliesReceived() uint32 {
	if m != nil {
		return m.ResolutionRepliesReceived
	}
	return 0
}

func (m *IpArpStatistics) GetResolutionRequestsDropped() uint32 {
	if m != nil {
		return m.ResolutionRequestsDropped
	}
	return 0
}

func (m *IpArpStatistics) GetOutOfMemoryErrors() uint32 {
	if m != nil {
		return m.OutOfMemoryErrors
	}
	return 0
}

func (m *IpArpStatistics) GetNoBufferErrors() uint32 {
	if m != nil {
		return m.NoBufferErrors
	}
	return 0
}

func (m *IpArpStatistics) GetTotalEntries() uint32 {
	if m != nil {
		return m.TotalEntries
	}
	return 0
}

func (m *IpArpStatistics) GetDynamicEntries() uint32 {
	if m != nil {
		return m.DynamicEntries
	}
	return 0
}

func (m *IpArpStatistics) GetStaticEntries() uint32 {
	if m != nil {
		return m.StaticEntries
	}
	return 0
}

func (m *IpArpStatistics) GetAliasEntries() uint32 {
	if m != nil {
		return m.AliasEntries
	}
	return 0
}

func (m *IpArpStatistics) GetInterfaceEntries() uint32 {
	if m != nil {
		return m.InterfaceEntries
	}
	return 0
}

func (m *IpArpStatistics) GetStandbyEntries() uint32 {
	if m != nil {
		return m.StandbyEntries
	}
	return 0
}

func (m *IpArpStatistics) GetDhcpEntries() uint32 {
	if m != nil {
		return m.DhcpEntries
	}
	return 0
}

func (m *IpArpStatistics) GetVxlanEntries() uint32 {
	if m != nil {
		return m.VxlanEntries
	}
	return 0
}

func (m *IpArpStatistics) GetDropAdjacencyEntries() uint32 {
	if m != nil {
		return m.DropAdjacencyEntries
	}
	return 0
}

func (m *IpArpStatistics) GetIpPacketsDroppedNode() uint32 {
	if m != nil {
		return m.IpPacketsDroppedNode
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketNodeOutOfSubnet() uint32 {
	if m != nil {
		return m.ArpPacketNodeOutOfSubnet
	}
	return 0
}

func (m *IpArpStatistics) GetIpPacketsDroppedInterface() uint32 {
	if m != nil {
		return m.IpPacketsDroppedInterface
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketInterfaceOutOfSubnet() uint32 {
	if m != nil {
		return m.ArpPacketInterfaceOutOfSubnet
	}
	return 0
}

func (m *IpArpStatistics) GetArpPacketUnsolicitedPacket() uint32 {
	if m != nil {
		return m.ArpPacketUnsolicitedPacket
	}
	return 0
}

func (m *IpArpStatistics) GetIdbStructures() uint32 {
	if m != nil {
		return m.IdbStructures
	}
	return 0
}

func init() {
	proto.RegisterType((*IpArpStatistics_KEYS)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.traffic_interfaces.traffic_interface.ip_arp_statistics_KEYS")
	proto.RegisterType((*IpArpStatistics)(nil), "cisco_ios_xr_ipv4_arp_oper.arp.nodes.node.traffic_interfaces.traffic_interface.ip_arp_statistics")
}

func init() { proto.RegisterFile("ip_arp_statistics.proto", fileDescriptor_394adc8d84b7147b) }

var fileDescriptor_394adc8d84b7147b = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0xdb, 0x4f, 0xdb, 0x48,
	0x14, 0xc6, 0xc5, 0x3e, 0xac, 0x96, 0x03, 0x01, 0xe2, 0x65, 0xc1, 0x2c, 0x20, 0x71, 0x11, 0x5a,
	0x56, 0xac, 0xb2, 0x52, 0xb9, 0xf7, 0xc2, 0xad, 0xa4, 0x94, 0x52, 0x02, 0x4a, 0xd4, 0x87, 0x4a,
	0x95, 0x46, 0x93, 0xf1, 0x84, 0x4e, 0xeb, 0xcc, 0xb8, 0x33, 0x63, 0x44, 0xfe, 0xeb, 0xfe, 0x09,
	0x95, 0xcf, 0xd8, 0x63, 0x27, 0xe1, 0x25, 0x0f, 0xdf, 0xf7, 0xfb, 0xce, 0x39, 0xe3, 0x39, 0x76,
	0x60, 0x51, 0x24, 0x84, 0xea, 0x84, 0x18, 0x4b, 0xad, 0x30, 0x56, 0x30, 0xd3, 0x48, 0xb4, 0xb2,
	0x2a, 0x68, 0x31, 0x61, 0x98, 0x22, 0x42, 0x19, 0xf2, 0xa4, 0x89, 0x48, 0x1e, 0xf7, 0x90, 0x53,
	0x09, 0xd7, 0x0d, 0xaa, 0x93, 0x86, 0x54, 0x11, 0x37, 0xf8, 0xdb, 0xb0, 0x9a, 0xf6, 0x7a, 0x82,
	0x11, 0x21, 0x2d, 0xd7, 0x3d, 0xca, 0xb8, 0x19, 0x97, 0x36, 0xbe, 0xc0, 0xc2, 0x58, 0x2b, 0x72,
	0xd3, 0xfc, 0xdc, 0x09, 0x96, 0x61, 0x32, 0x2b, 0x43, 0x24, 0xed, 0xf3, 0x70, 0x62, 0x6d, 0x62,
	0x7b, 0xb2, 0xfd, 0x47, 0x26, 0xb4, 0x68, 0x9f, 0x07, 0x5b, 0x30, 0xe3, 0x6b, 0x38, 0xe2, 0x37,
	0x24, 0x6a, 0x5e, 0xcd, 0xb0, 0x8d, 0x9f, 0x53, 0x50, 0x1f, 0x2b, 0x1f, 0xec, 0x40, 0x5d, 0xf3,
	0x1f, 0x29, 0x37, 0xd6, 0x10, 0xcd, 0x19, 0x17, 0x8f, 0x3c, 0x0a, 0x5f, 0xac, 0x4d, 0x6c, 0xd7,
	0xda, 0x73, 0x85, 0xd1, 0xce, 0xf5, 0xe0, 0x5f, 0x98, 0xd3, 0x3c, 0x89, 0x05, 0xaf, 0xb0, 0xbb,
	0xc8, 0xce, 0xe6, 0xba, 0x47, 0x37, 0xa1, 0xe6, 0xeb, 0x1a, 0x2e, 0x6d, 0xb8, 0x87, 0xdc, 0x74,
	0x21, 0x76, 0xb8, 0xb4, 0xc1, 0x3a, 0x4c, 0x17, 0xf5, 0x90, 0xd9, 0x47, 0x66, 0x2a, 0xd7, 0x10,
	0xf9, 0x0f, 0x82, 0x44, 0xab, 0xa7, 0x01, 0x19, 0x02, 0x0f, 0xdc, 0x80, 0xe8, 0xb4, 0x2b, 0xf4,
	0x11, 0x84, 0x26, 0xed, 0x1a, 0xa6, 0xc9, 0xf8, 0xa1, 0x0e, 0x31, 0xb3, 0xe0, 0xfc, 0xf6, 0xe8,
	0xd1, 0x1a, 0xf0, 0xa7, 0x4f, 0x56, 0x1a, 0x1d, 0x61, 0xa8, 0x5e, 0x84, 0xca, 0x4e, 0xc7, 0xb0,
	0x34, 0xc2, 0x3f, 0x68, 0x6a, 0x1f, 0x5c, 0xea, 0x78, 0xb8, 0x15, 0xfa, 0x57, 0x99, 0x8d, 0xd1,
	0x43, 0x08, 0x63, 0xc5, 0x68, 0x4c, 0x9e, 0x39, 0xd8, 0x4b, 0x4c, 0xfe, 0x85, 0xfe, 0xfd, 0xe8,
	0xe9, 0x0e, 0x60, 0x31, 0x6b, 0x92, 0x0a, 0xab, 0x52, 0x33, 0x9c, 0x7b, 0xe5, 0x72, 0xa5, 0x5d,
	0xcd, 0x9d, 0xc1, 0x8a, 0xe6, 0x46, 0xc5, 0xa9, 0x15, 0x4a, 0x3e, 0xf3, 0x64, 0x5e, 0x63, 0xf8,
	0xef, 0x92, 0x19, 0x7b, 0x3a, 0x27, 0xb0, 0x3c, 0x54, 0x61, 0x64, 0x07, 0xde, 0x60, 0x81, 0xa5,
	0x6a, 0x81, 0xe1, 0x6d, 0x18, 0xcd, 0xe7, 0x13, 0x44, 0x5a, 0x25, 0x09, 0x8f, 0xc2, 0x93, 0xf1,
	0xbc, 0x23, 0x2e, 0x1d, 0x10, 0xfc, 0x0f, 0xf3, 0x2a, 0xb5, 0x44, 0xf5, 0x48, 0x9f, 0xf7, 0x95,
	0x1e, 0x10, 0xae, 0xb5, 0xd2, 0x26, 0x3c, 0x75, 0xd7, 0xa3, 0x52, 0x7b, 0xd7, 0xbb, 0x45, 0xa7,
	0x89, 0x46, 0xb0, 0x0d, 0x73, 0x52, 0x91, 0x6e, 0xda, 0xeb, 0x71, 0x5d, 0xc0, 0x67, 0x08, 0xcf,
	0x48, 0x75, 0x81, 0x72, 0x4e, 0x6e, 0x42, 0xcd, 0x2a, 0x4b, 0x63, 0xc2, 0xa5, 0xd5, 0x82, 0x9b,
	0xf0, 0xdc, 0x2d, 0x2a, 0x8a, 0x4d, 0xa7, 0x05, 0xff, 0xc0, 0x6c, 0x34, 0x90, 0xb4, 0x2f, 0x98,
	0xc7, 0x2e, 0x5c, 0xb5, 0x5c, 0x2e, 0xc0, 0x2d, 0x98, 0xc1, 0x97, 0xab, 0xe4, 0xde, 0x22, 0x57,
	0x73, 0x6a, 0x81, 0x6d, 0x42, 0x8d, 0xc6, 0x82, 0x1a, 0x4f, 0x5d, 0xba, 0xa6, 0x28, 0x16, 0xd0,
	0x0e, 0xd4, 0xcb, 0xf7, 0xba, 0x00, 0x9b, 0x6e, 0xf3, 0xbd, 0x51, 0x99, 0xd0, 0x58, 0x2a, 0xa3,
	0xee, 0xc0, 0xa3, 0xef, 0xdc, 0x84, 0xb9, 0x5c, 0x80, 0xeb, 0x30, 0x1d, 0x7d, 0x65, 0x89, 0xa7,
	0xae, 0xdc, 0x3b, 0x97, 0x69, 0x95, 0xe9, 0x1e, 0x9f, 0x62, 0x2a, 0x3d, 0xf3, 0xde, 0x4d, 0x87,
	0x62, 0x01, 0xed, 0xc1, 0x42, 0x76, 0x7d, 0x84, 0x46, 0xdf, 0x28, 0xe3, 0x92, 0x95, 0x7d, 0xaf,
	0x91, 0x9e, 0xcf, 0xdc, 0xf3, 0xc2, 0x2c, 0x52, 0xfb, 0xf8, 0x35, 0x4d, 0x28, 0xfb, 0xce, 0xcb,
	0xfb, 0x27, 0xd9, 0xa7, 0x2c, 0xfc, 0xe0, 0x62, 0x22, 0xb9, 0x77, 0x6e, 0x7e, 0xf7, 0x2d, 0x15,
	0xf1, 0xe0, 0x14, 0x56, 0xb3, 0xef, 0x96, 0xcb, 0x21, 0x4e, 0xf2, 0x7d, 0x30, 0x69, 0x57, 0x72,
	0x1b, 0xde, 0x60, 0x38, 0xa4, 0x3a, 0x4f, 0x67, 0xa9, 0xbb, 0x6c, 0x2b, 0x3a, 0xe8, 0x07, 0xa7,
	0xb0, 0xf2, 0x4c, 0x5f, 0xff, 0x14, 0xc3, 0x8f, 0x6e, 0x03, 0x47, 0x9b, 0x5f, 0x17, 0x40, 0x70,
	0x0d, 0x1b, 0x95, 0x09, 0xca, 0x7b, 0x19, 0x1e, 0xe3, 0x16, 0xcb, 0xac, 0xfa, 0x31, 0x7c, 0xbe,
	0x3a, 0xcb, 0xf9, 0xd0, 0x61, 0x52, 0x69, 0x54, 0x2c, 0x98, 0xb0, 0x3c, 0xca, 0xa5, 0xb0, 0xe5,
	0xde, 0x47, 0x5f, 0xe5, 0x53, 0x89, 0x38, 0x01, 0x3f, 0xf9, 0x51, 0x97, 0x18, 0xab, 0x53, 0x66,
	0x53, 0xcd, 0x4d, 0x78, 0xe7, 0xd6, 0x4c, 0x44, 0xdd, 0x8e, 0x17, 0xbb, 0xbf, 0xe3, 0xff, 0xd4,
	0xee, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x8f, 0xeb, 0xf2, 0xc2, 0x06, 0x00, 0x00,
}
