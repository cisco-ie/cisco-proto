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
// source: pce_node_summary_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_topology_topology_summary

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

type PceNodeSummaryBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceNodeSummaryBag_KEYS) Reset()         { *m = PceNodeSummaryBag_KEYS{} }
func (m *PceNodeSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceNodeSummaryBag_KEYS) ProtoMessage()    {}
func (*PceNodeSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0129059af41579e5, []int{0}
}

func (m *PceNodeSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNodeSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *PceNodeSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNodeSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceNodeSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNodeSummaryBag_KEYS.Merge(m, src)
}
func (m *PceNodeSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceNodeSummaryBag_KEYS.Size(m)
}
func (m *PceNodeSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNodeSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceNodeSummaryBag_KEYS proto.InternalMessageInfo

type PceTopoStatsBag struct {
	NumNodesAdded        uint32   `protobuf:"varint,1,opt,name=num_nodes_added,json=numNodesAdded,proto3" json:"num_nodes_added,omitempty"`
	NumNodesDeleted      uint32   `protobuf:"varint,2,opt,name=num_nodes_deleted,json=numNodesDeleted,proto3" json:"num_nodes_deleted,omitempty"`
	NumLinksAdded        uint32   `protobuf:"varint,3,opt,name=num_links_added,json=numLinksAdded,proto3" json:"num_links_added,omitempty"`
	NumLinksDeleted      uint32   `protobuf:"varint,4,opt,name=num_links_deleted,json=numLinksDeleted,proto3" json:"num_links_deleted,omitempty"`
	NumPrefixesAdded     uint32   `protobuf:"varint,5,opt,name=num_prefixes_added,json=numPrefixesAdded,proto3" json:"num_prefixes_added,omitempty"`
	NumPrefixesDeleted   uint32   `protobuf:"varint,6,opt,name=num_prefixes_deleted,json=numPrefixesDeleted,proto3" json:"num_prefixes_deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceTopoStatsBag) Reset()         { *m = PceTopoStatsBag{} }
func (m *PceTopoStatsBag) String() string { return proto.CompactTextString(m) }
func (*PceTopoStatsBag) ProtoMessage()    {}
func (*PceTopoStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0129059af41579e5, []int{1}
}

func (m *PceTopoStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTopoStatsBag.Unmarshal(m, b)
}
func (m *PceTopoStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTopoStatsBag.Marshal(b, m, deterministic)
}
func (m *PceTopoStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTopoStatsBag.Merge(m, src)
}
func (m *PceTopoStatsBag) XXX_Size() int {
	return xxx_messageInfo_PceTopoStatsBag.Size(m)
}
func (m *PceTopoStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTopoStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceTopoStatsBag proto.InternalMessageInfo

func (m *PceTopoStatsBag) GetNumNodesAdded() uint32 {
	if m != nil {
		return m.NumNodesAdded
	}
	return 0
}

func (m *PceTopoStatsBag) GetNumNodesDeleted() uint32 {
	if m != nil {
		return m.NumNodesDeleted
	}
	return 0
}

func (m *PceTopoStatsBag) GetNumLinksAdded() uint32 {
	if m != nil {
		return m.NumLinksAdded
	}
	return 0
}

func (m *PceTopoStatsBag) GetNumLinksDeleted() uint32 {
	if m != nil {
		return m.NumLinksDeleted
	}
	return 0
}

func (m *PceTopoStatsBag) GetNumPrefixesAdded() uint32 {
	if m != nil {
		return m.NumPrefixesAdded
	}
	return 0
}

func (m *PceTopoStatsBag) GetNumPrefixesDeleted() uint32 {
	if m != nil {
		return m.NumPrefixesDeleted
	}
	return 0
}

type PceNodeSummaryBag struct {
	Nodes                    uint32           `protobuf:"varint,50,opt,name=nodes,proto3" json:"nodes,omitempty"`
	LookupNodes              uint32           `protobuf:"varint,51,opt,name=lookup_nodes,json=lookupNodes,proto3" json:"lookup_nodes,omitempty"`
	Prefixes                 uint32           `protobuf:"varint,52,opt,name=prefixes,proto3" json:"prefixes,omitempty"`
	PrefixSids               uint32           `protobuf:"varint,53,opt,name=prefix_sids,json=prefixSids,proto3" json:"prefix_sids,omitempty"`
	RegularPrefixSids        uint32           `protobuf:"varint,54,opt,name=regular_prefix_sids,json=regularPrefixSids,proto3" json:"regular_prefix_sids,omitempty"`
	StrictPrefixSids         uint32           `protobuf:"varint,55,opt,name=strict_prefix_sids,json=strictPrefixSids,proto3" json:"strict_prefix_sids,omitempty"`
	Links                    uint32           `protobuf:"varint,56,opt,name=links,proto3" json:"links,omitempty"`
	EpeLinks                 uint32           `protobuf:"varint,57,opt,name=epe_links,json=epeLinks,proto3" json:"epe_links,omitempty"`
	AdjacencySids            uint32           `protobuf:"varint,58,opt,name=adjacency_sids,json=adjacencySids,proto3" json:"adjacency_sids,omitempty"`
	Epesids                  uint32           `protobuf:"varint,59,opt,name=epesids,proto3" json:"epesids,omitempty"`
	ProtectedAdjacencySids   uint32           `protobuf:"varint,60,opt,name=protected_adjacency_sids,json=protectedAdjacencySids,proto3" json:"protected_adjacency_sids,omitempty"`
	UnProtectedAdjacencySids uint32           `protobuf:"varint,61,opt,name=un_protected_adjacency_sids,json=unProtectedAdjacencySids,proto3" json:"un_protected_adjacency_sids,omitempty"`
	StatsTopologyUpdate      *PceTopoStatsBag `protobuf:"bytes,62,opt,name=stats_topology_update,json=statsTopologyUpdate,proto3" json:"stats_topology_update,omitempty"`
	TopologyConsistent       bool             `protobuf:"varint,63,opt,name=topology_consistent,json=topologyConsistent,proto3" json:"topology_consistent,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}         `json:"-"`
	XXX_unrecognized         []byte           `json:"-"`
	XXX_sizecache            int32            `json:"-"`
}

func (m *PceNodeSummaryBag) Reset()         { *m = PceNodeSummaryBag{} }
func (m *PceNodeSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PceNodeSummaryBag) ProtoMessage()    {}
func (*PceNodeSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0129059af41579e5, []int{2}
}

func (m *PceNodeSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceNodeSummaryBag.Unmarshal(m, b)
}
func (m *PceNodeSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceNodeSummaryBag.Marshal(b, m, deterministic)
}
func (m *PceNodeSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceNodeSummaryBag.Merge(m, src)
}
func (m *PceNodeSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PceNodeSummaryBag.Size(m)
}
func (m *PceNodeSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceNodeSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceNodeSummaryBag proto.InternalMessageInfo

func (m *PceNodeSummaryBag) GetNodes() uint32 {
	if m != nil {
		return m.Nodes
	}
	return 0
}

func (m *PceNodeSummaryBag) GetLookupNodes() uint32 {
	if m != nil {
		return m.LookupNodes
	}
	return 0
}

func (m *PceNodeSummaryBag) GetPrefixes() uint32 {
	if m != nil {
		return m.Prefixes
	}
	return 0
}

func (m *PceNodeSummaryBag) GetPrefixSids() uint32 {
	if m != nil {
		return m.PrefixSids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetRegularPrefixSids() uint32 {
	if m != nil {
		return m.RegularPrefixSids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetStrictPrefixSids() uint32 {
	if m != nil {
		return m.StrictPrefixSids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetLinks() uint32 {
	if m != nil {
		return m.Links
	}
	return 0
}

func (m *PceNodeSummaryBag) GetEpeLinks() uint32 {
	if m != nil {
		return m.EpeLinks
	}
	return 0
}

func (m *PceNodeSummaryBag) GetAdjacencySids() uint32 {
	if m != nil {
		return m.AdjacencySids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetEpesids() uint32 {
	if m != nil {
		return m.Epesids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetProtectedAdjacencySids() uint32 {
	if m != nil {
		return m.ProtectedAdjacencySids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetUnProtectedAdjacencySids() uint32 {
	if m != nil {
		return m.UnProtectedAdjacencySids
	}
	return 0
}

func (m *PceNodeSummaryBag) GetStatsTopologyUpdate() *PceTopoStatsBag {
	if m != nil {
		return m.StatsTopologyUpdate
	}
	return nil
}

func (m *PceNodeSummaryBag) GetTopologyConsistent() bool {
	if m != nil {
		return m.TopologyConsistent
	}
	return false
}

func init() {
	proto.RegisterType((*PceNodeSummaryBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.topology_summary.pce_node_summary_bag_KEYS")
	proto.RegisterType((*PceTopoStatsBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.topology_summary.pce_topo_stats_bag")
	proto.RegisterType((*PceNodeSummaryBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce_topology.topology_summary.pce_node_summary_bag")
}

func init() { proto.RegisterFile("pce_node_summary_bag.proto", fileDescriptor_0129059af41579e5) }

var fileDescriptor_0129059af41579e5 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0x41, 0xc7, 0x78, 0x65, 0x8c, 0xb9, 0x05, 0x99, 0xf5, 0x40, 0xa9, 0x04, 0xaa,
	0xd0, 0x14, 0xd0, 0xc6, 0x8f, 0x0d, 0x18, 0x68, 0x02, 0x4e, 0xfc, 0x50, 0xd5, 0xc1, 0x81, 0x93,
	0x95, 0xc5, 0x6f, 0x55, 0x58, 0x62, 0x5b, 0xb1, 0x23, 0xb5, 0x47, 0xee, 0xfc, 0x09, 0xfc, 0xb1,
	0x28, 0xcf, 0x75, 0xd6, 0x42, 0xb9, 0xec, 0xe6, 0xf7, 0xbe, 0x9f, 0xef, 0x37, 0xc9, 0xb3, 0x1d,
	0xd8, 0x31, 0x29, 0x0a, 0xa5, 0x25, 0x0a, 0x5b, 0x15, 0x45, 0x52, 0xce, 0xc4, 0x69, 0x32, 0x89,
	0x4d, 0xa9, 0x9d, 0x66, 0x87, 0x69, 0x66, 0x53, 0x2d, 0x32, 0x6d, 0xc5, 0xb4, 0x14, 0x99, 0x3a,
	0x2b, 0x13, 0x31, 0x75, 0xa9, 0xd0, 0x06, 0xcb, 0xb8, 0xf6, 0x39, 0x6d, 0x74, 0xae, 0x27, 0xb3,
	0x38, 0x2c, 0x42, 0xc8, 0xa0, 0x07, 0x77, 0x57, 0x05, 0x8b, 0x8f, 0x1f, 0xbe, 0x9f, 0x0c, 0x7e,
	0xaf, 0x01, 0x0b, 0x76, 0x61, 0x5d, 0xe2, 0x6c, 0xad, 0xb1, 0x87, 0xb0, 0xa5, 0xaa, 0x82, 0x3c,
	0x56, 0x24, 0x52, 0xa2, 0xe4, 0x51, 0x3f, 0x1a, 0x6e, 0x8e, 0x37, 0x55, 0x55, 0x7c, 0xa9, 0xbb,
	0xc7, 0x75, 0x93, 0x3d, 0x82, 0xed, 0x0b, 0x4e, 0x62, 0x8e, 0x0e, 0x25, 0x5f, 0x23, 0x72, 0x2b,
	0x90, 0xef, 0x7d, 0x3b, 0x64, 0xe6, 0x99, 0x3a, 0x0f, 0x99, 0x57, 0x9a, 0xcc, 0x4f, 0x75, 0x77,
	0x29, 0xd3, 0x73, 0x21, 0xf3, 0x6a, 0x93, 0x49, 0x64, 0xc8, 0xdc, 0x05, 0x56, 0xb3, 0xa6, 0xc4,
	0xb3, 0x6c, 0xda, 0xbc, 0x6a, 0x8b, 0xe0, 0x5b, 0xaa, 0x2a, 0x46, 0x73, 0xc1, 0x27, 0x3f, 0x81,
	0xee, 0x12, 0x1d, 0xc2, 0xd7, 0x89, 0x67, 0x0b, 0xfc, 0x3c, 0x7f, 0xf0, 0xab, 0x05, 0xdd, 0x55,
	0xc3, 0x63, 0x5d, 0x68, 0xd1, 0x47, 0xf3, 0x3d, 0xf2, 0xfa, 0x82, 0xdd, 0x87, 0x1b, 0xb9, 0xd6,
	0xe7, 0x95, 0xf1, 0x13, 0xe1, 0xfb, 0x24, 0xb6, 0x7d, 0x8f, 0x86, 0xc1, 0x76, 0x60, 0x23, 0x3c,
	0x9f, 0x3f, 0x25, 0xb9, 0xa9, 0xd9, 0x3d, 0x68, 0xfb, 0xb5, 0xb0, 0x99, 0xb4, 0xfc, 0x19, 0xc9,
	0xe0, 0x5b, 0x27, 0x99, 0xb4, 0x2c, 0x86, 0x4e, 0x89, 0x93, 0x2a, 0x4f, 0x4a, 0xb1, 0x08, 0x3e,
	0x27, 0x70, 0x7b, 0x2e, 0x8d, 0x2e, 0xf8, 0x5d, 0x60, 0xd6, 0x95, 0x59, 0xea, 0x96, 0xf0, 0x17,
	0x7e, 0x3c, 0x5e, 0x59, 0xa0, 0xbb, 0xd0, 0xa2, 0xa1, 0xf3, 0x03, 0xff, 0x4d, 0x54, 0xb0, 0x1e,
	0x5c, 0x47, 0x83, 0x7e, 0x3b, 0xf8, 0xa1, 0x7f, 0x63, 0x34, 0x48, 0xdb, 0xc0, 0x1e, 0xc0, 0xcd,
	0x44, 0xfe, 0x48, 0x52, 0x54, 0xe9, 0xcc, 0x87, 0xbf, 0xf4, 0x5b, 0xda, 0x74, 0x29, 0x99, 0xc3,
	0x35, 0x34, 0x48, 0xfa, 0x2b, 0xd2, 0x43, 0xc9, 0x0e, 0x80, 0xd7, 0x07, 0x1c, 0x53, 0x87, 0x52,
	0xfc, 0x15, 0xf5, 0x9a, 0xd0, 0x3b, 0x8d, 0x7e, 0xbc, 0x94, 0x79, 0x04, 0xbd, 0x4a, 0x89, 0xff,
	0x9a, 0x8f, 0xc8, 0xcc, 0x2b, 0x35, 0x5a, 0x6d, 0xff, 0x19, 0xc1, 0x6d, 0x7f, 0xde, 0x9b, 0x0b,
	0x53, 0x19, 0x99, 0x38, 0xe4, 0x6f, 0xfa, 0xd1, 0xb0, 0xbd, 0xf7, 0x39, 0xbe, 0xf4, 0x8d, 0x8b,
	0xff, 0xbd, 0x50, 0xe3, 0x0e, 0x2d, 0xbf, 0xce, 0xc9, 0x6f, 0xf4, 0x24, 0xf6, 0x18, 0x3a, 0x8d,
	0x37, 0xd5, 0xca, 0x66, 0xd6, 0xa1, 0x72, 0xfc, 0x6d, 0x3f, 0x1a, 0x6e, 0x8c, 0x59, 0x90, 0xde,
	0x35, 0xca, 0xe9, 0x3a, 0xfd, 0x0c, 0xf6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0xef, 0x3e,
	0x73, 0x2a, 0x04, 0x00, 0x00,
}
