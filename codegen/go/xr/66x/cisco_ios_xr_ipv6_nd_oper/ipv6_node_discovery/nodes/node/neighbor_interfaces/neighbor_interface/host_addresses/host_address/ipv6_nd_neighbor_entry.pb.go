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
// source: ipv6_nd_neighbor_entry.proto

package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_neighbor_interfaces_neighbor_interface_host_addresses_host_address

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

type Ipv6NdNeighborEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	HostAddress          string   `protobuf:"bytes,3,opt,name=host_address,json=hostAddress,proto3" json:"host_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdNeighborEntry_KEYS) Reset()         { *m = Ipv6NdNeighborEntry_KEYS{} }
func (m *Ipv6NdNeighborEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntry_KEYS) ProtoMessage()    {}
func (*Ipv6NdNeighborEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a853b2f12f0146e, []int{0}
}

func (m *Ipv6NdNeighborEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdNeighborEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv6NdNeighborEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdNeighborEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6NdNeighborEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdNeighborEntry_KEYS.Merge(m, src)
}
func (m *Ipv6NdNeighborEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdNeighborEntry_KEYS.Size(m)
}
func (m *Ipv6NdNeighborEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdNeighborEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdNeighborEntry_KEYS proto.InternalMessageInfo

func (m *Ipv6NdNeighborEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry_KEYS) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

type BagTimespec struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BagTimespec) Reset()         { *m = BagTimespec{} }
func (m *BagTimespec) String() string { return proto.CompactTextString(m) }
func (*BagTimespec) ProtoMessage()    {}
func (*BagTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a853b2f12f0146e, []int{1}
}

func (m *BagTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagTimespec.Unmarshal(m, b)
}
func (m *BagTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagTimespec.Marshal(b, m, deterministic)
}
func (m *BagTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagTimespec.Merge(m, src)
}
func (m *BagTimespec) XXX_Size() int {
	return xxx_messageInfo_BagTimespec.Size(m)
}
func (m *BagTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_BagTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_BagTimespec proto.InternalMessageInfo

func (m *BagTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

type Ipv6NdNeighborEntry struct {
	LastReachedTime       *BagTimespec `protobuf:"bytes,50,opt,name=last_reached_time,json=lastReachedTime,proto3" json:"last_reached_time,omitempty"`
	ReachabilityState     string       `protobuf:"bytes,51,opt,name=reachability_state,json=reachabilityState,proto3" json:"reachability_state,omitempty"`
	LinkLayerAddress      string       `protobuf:"bytes,52,opt,name=link_layer_address,json=linkLayerAddress,proto3" json:"link_layer_address,omitempty"`
	Encapsulation         string       `protobuf:"bytes,53,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	SelectedEncapsulation string       `protobuf:"bytes,54,opt,name=selected_encapsulation,json=selectedEncapsulation,proto3" json:"selected_encapsulation,omitempty"`
	OriginEncapsulation   string       `protobuf:"bytes,55,opt,name=origin_encapsulation,json=originEncapsulation,proto3" json:"origin_encapsulation,omitempty"`
	InterfaceName         string       `protobuf:"bytes,56,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Location              string       `protobuf:"bytes,57,opt,name=location,proto3" json:"location,omitempty"`
	IsRouter              bool         `protobuf:"varint,58,opt,name=is_router,json=isRouter,proto3" json:"is_router,omitempty"`
	SergFlags             uint32       `protobuf:"varint,59,opt,name=serg_flags,json=sergFlags,proto3" json:"serg_flags,omitempty"`
	Vrfid                 uint32       `protobuf:"varint,60,opt,name=vrfid,proto3" json:"vrfid,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *Ipv6NdNeighborEntry) Reset()         { *m = Ipv6NdNeighborEntry{} }
func (m *Ipv6NdNeighborEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntry) ProtoMessage()    {}
func (*Ipv6NdNeighborEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a853b2f12f0146e, []int{2}
}

func (m *Ipv6NdNeighborEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdNeighborEntry.Unmarshal(m, b)
}
func (m *Ipv6NdNeighborEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdNeighborEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6NdNeighborEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdNeighborEntry.Merge(m, src)
}
func (m *Ipv6NdNeighborEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdNeighborEntry.Size(m)
}
func (m *Ipv6NdNeighborEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdNeighborEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdNeighborEntry proto.InternalMessageInfo

func (m *Ipv6NdNeighborEntry) GetLastReachedTime() *BagTimespec {
	if m != nil {
		return m.LastReachedTime
	}
	return nil
}

func (m *Ipv6NdNeighborEntry) GetReachabilityState() string {
	if m != nil {
		return m.ReachabilityState
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetLinkLayerAddress() string {
	if m != nil {
		return m.LinkLayerAddress
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetSelectedEncapsulation() string {
	if m != nil {
		return m.SelectedEncapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetOriginEncapsulation() string {
	if m != nil {
		return m.OriginEncapsulation
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Ipv6NdNeighborEntry) GetIsRouter() bool {
	if m != nil {
		return m.IsRouter
	}
	return false
}

func (m *Ipv6NdNeighborEntry) GetSergFlags() uint32 {
	if m != nil {
		return m.SergFlags
	}
	return 0
}

func (m *Ipv6NdNeighborEntry) GetVrfid() uint32 {
	if m != nil {
		return m.Vrfid
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6NdNeighborEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.ipv6_nd_neighbor_entry_KEYS")
	proto.RegisterType((*BagTimespec)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.bag_timespec")
	proto.RegisterType((*Ipv6NdNeighborEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_interfaces.neighbor_interface.host_addresses.host_address.ipv6_nd_neighbor_entry")
}

func init() { proto.RegisterFile("ipv6_nd_neighbor_entry.proto", fileDescriptor_7a853b2f12f0146e) }

var fileDescriptor_7a853b2f12f0146e = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xf2, 0x9b, 0x4c, 0x1b, 0xa0, 0xa6, 0x54, 0x16, 0x05, 0x29, 0x54, 0x20, 0xe5, 0x00,
	0x2b, 0xd1, 0xd2, 0xf2, 0x7b, 0xe1, 0x50, 0x2e, 0x20, 0x0e, 0x5b, 0x2e, 0x9c, 0x2c, 0xc7, 0x9e,
	0x6c, 0x2c, 0x36, 0xf6, 0xca, 0xe3, 0x46, 0x44, 0x5c, 0xb9, 0xf0, 0x42, 0xbc, 0x08, 0x2f, 0x84,
	0x6c, 0x37, 0x51, 0x96, 0xe6, 0xce, 0x65, 0xa5, 0xef, 0x67, 0x66, 0xbe, 0x9d, 0x9d, 0x85, 0x07,
	0xa6, 0x9d, 0x9f, 0x08, 0xab, 0x85, 0x45, 0x53, 0x4f, 0xc7, 0xce, 0x0b, 0xb4, 0xc1, 0x2f, 0xca,
	0xd6, 0xbb, 0xe0, 0xd8, 0x0f, 0x65, 0x48, 0x39, 0x61, 0x1c, 0x89, 0xef, 0x5e, 0x2c, 0xad, 0xae,
	0x45, 0x5f, 0x66, 0xe0, 0x34, 0x0a, 0x1d, 0x3d, 0x73, 0xf4, 0x8b, 0x32, 0x42, 0x4a, 0xcf, 0x72,
	0xd5, 0xce, 0xd8, 0x80, 0x7e, 0x22, 0x55, 0x14, 0x2e, 0x71, 0xe5, 0xd4, 0x51, 0x10, 0x52, 0x6b,
	0x8f, 0x44, 0x48, 0x1d, 0x78, 0xf0, 0xb3, 0x80, 0xfd, 0xcd, 0xe9, 0xc4, 0xc7, 0xd3, 0xaf, 0x67,
	0x6c, 0x1f, 0xfa, 0x69, 0xbe, 0x95, 0x33, 0xe4, 0xc5, 0xb0, 0x18, 0xf5, 0xab, 0x5e, 0x24, 0x3e,
	0xcb, 0x19, 0xb2, 0x27, 0x70, 0x6b, 0x35, 0x25, 0x3b, 0xae, 0x24, 0xc7, 0x60, 0xc5, 0x26, 0xdb,
	0x23, 0xd8, 0x5e, 0x9f, 0xc9, 0xaf, 0x26, 0xd3, 0x56, 0xe4, 0xde, 0x5f, 0xc4, 0x18, 0xc1, 0xf6,
	0x58, 0xd6, 0x22, 0x98, 0x19, 0x52, 0x8b, 0x8a, 0x71, 0xb8, 0x49, 0xa8, 0x9c, 0xd5, 0x94, 0x86,
	0x0e, 0xaa, 0x25, 0x3c, 0xf8, 0x73, 0x0d, 0xf6, 0x36, 0x07, 0x66, 0xbf, 0x0b, 0xd8, 0x69, 0x24,
	0x05, 0xe1, 0x51, 0xaa, 0x29, 0xea, 0xd4, 0x8e, 0x1f, 0x0e, 0x8b, 0xd1, 0xd6, 0xe1, 0xaf, 0xa2,
	0xfc, 0x8f, 0x6b, 0x2e, 0xd7, 0x5f, 0xae, 0xba, 0x1d, 0x43, 0x56, 0x39, 0xe3, 0x17, 0x33, 0x43,
	0xf6, 0x0c, 0x58, 0x8a, 0x2c, 0xc7, 0xa6, 0x31, 0x61, 0x21, 0x28, 0xc8, 0x80, 0xfc, 0x28, 0xad,
	0x69, 0x67, 0x5d, 0x39, 0x8b, 0x02, 0x7b, 0x0a, 0xac, 0x31, 0xf6, 0x9b, 0x68, 0xe4, 0x02, 0xfd,
	0x6a, 0xab, 0x2f, 0x92, 0xfd, 0x4e, 0x54, 0x3e, 0x45, 0xe1, 0x62, 0xb5, 0xec, 0x31, 0x0c, 0xd0,
	0x2a, 0xd9, 0xd2, 0x79, 0x23, 0x83, 0x71, 0x96, 0x1f, 0xe7, 0x6f, 0xd4, 0x21, 0xd9, 0x31, 0xec,
	0x11, 0x36, 0xa8, 0x02, 0x6a, 0xd1, 0xb5, 0x9f, 0x24, 0xfb, 0xbd, 0xa5, 0x7a, 0xda, 0x29, 0x7b,
	0x0e, 0xbb, 0xce, 0x9b, 0xda, 0xd8, 0x7f, 0x8a, 0x5e, 0xa6, 0xa2, 0xbb, 0x59, 0xeb, 0x96, 0x5c,
	0x3e, 0x9a, 0x57, 0x9b, 0x8e, 0xe6, 0x3e, 0xf4, 0x1a, 0xa7, 0x72, 0xb7, 0xd7, 0xf9, 0xee, 0x96,
	0x38, 0x1e, 0xa5, 0x21, 0xe1, 0xdd, 0x79, 0x40, 0xcf, 0xdf, 0x0c, 0x8b, 0x51, 0xaf, 0xea, 0x19,
	0xaa, 0x12, 0x66, 0x0f, 0x01, 0x08, 0x7d, 0x2d, 0x26, 0x8d, 0xac, 0x89, 0xbf, 0x4d, 0xd7, 0xd3,
	0x8f, 0xcc, 0x87, 0x48, 0xb0, 0x5d, 0xb8, 0x3e, 0xf7, 0x13, 0xa3, 0xf9, 0xbb, 0xa4, 0x64, 0x30,
	0xbe, 0x91, 0x7e, 0xc5, 0xa3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x50, 0x6b, 0xc6, 0xaa,
	0x03, 0x00, 0x00,
}
