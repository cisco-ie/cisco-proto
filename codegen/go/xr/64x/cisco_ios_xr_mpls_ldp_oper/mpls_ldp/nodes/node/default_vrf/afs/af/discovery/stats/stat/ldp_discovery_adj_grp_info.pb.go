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
// source: ldp_discovery_adj_grp_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_afs_af_discovery_stats_stat

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

type LdpDiscoveryAdjGrpInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LsrId                string   `protobuf:"bytes,3,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,4,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpDiscoveryAdjGrpInfo_KEYS) Reset()         { *m = LdpDiscoveryAdjGrpInfo_KEYS{} }
func (m *LdpDiscoveryAdjGrpInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryAdjGrpInfo_KEYS) ProtoMessage()    {}
func (*LdpDiscoveryAdjGrpInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8555183d56035bcf, []int{0}
}

func (m *LdpDiscoveryAdjGrpInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpDiscoveryAdjGrpInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryAdjGrpInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS.Merge(m, src)
}
func (m *LdpDiscoveryAdjGrpInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS.Size(m)
}
func (m *LdpDiscoveryAdjGrpInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryAdjGrpInfo_KEYS proto.InternalMessageInfo

func (m *LdpDiscoveryAdjGrpInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpDiscoveryAdjGrpInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpDiscoveryAdjGrpInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpDiscoveryAdjGrpInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

type LdpDiscoveryAdjGrpInfo struct {
	AdjacencyGroupUpTime uint32   `protobuf:"varint,50,opt,name=adjacency_group_up_time,json=adjacencyGroupUpTime,proto3" json:"adjacency_group_up_time,omitempty"`
	TcpOpenCount         uint32   `protobuf:"varint,51,opt,name=tcp_open_count,json=tcpOpenCount,proto3" json:"tcp_open_count,omitempty"`
	TcpArbChgCount       uint32   `protobuf:"varint,52,opt,name=tcp_arb_chg_count,json=tcpArbChgCount,proto3" json:"tcp_arb_chg_count,omitempty"`
	TcpRole              uint32   `protobuf:"varint,53,opt,name=tcp_role,json=tcpRole,proto3" json:"tcp_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpDiscoveryAdjGrpInfo) Reset()         { *m = LdpDiscoveryAdjGrpInfo{} }
func (m *LdpDiscoveryAdjGrpInfo) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryAdjGrpInfo) ProtoMessage()    {}
func (*LdpDiscoveryAdjGrpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8555183d56035bcf, []int{1}
}

func (m *LdpDiscoveryAdjGrpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo.Unmarshal(m, b)
}
func (m *LdpDiscoveryAdjGrpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryAdjGrpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryAdjGrpInfo.Merge(m, src)
}
func (m *LdpDiscoveryAdjGrpInfo) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryAdjGrpInfo.Size(m)
}
func (m *LdpDiscoveryAdjGrpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryAdjGrpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryAdjGrpInfo proto.InternalMessageInfo

func (m *LdpDiscoveryAdjGrpInfo) GetAdjacencyGroupUpTime() uint32 {
	if m != nil {
		return m.AdjacencyGroupUpTime
	}
	return 0
}

func (m *LdpDiscoveryAdjGrpInfo) GetTcpOpenCount() uint32 {
	if m != nil {
		return m.TcpOpenCount
	}
	return 0
}

func (m *LdpDiscoveryAdjGrpInfo) GetTcpArbChgCount() uint32 {
	if m != nil {
		return m.TcpArbChgCount
	}
	return 0
}

func (m *LdpDiscoveryAdjGrpInfo) GetTcpRole() uint32 {
	if m != nil {
		return m.TcpRole
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpDiscoveryAdjGrpInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.stats.stat.ldp_discovery_adj_grp_info_KEYS")
	proto.RegisterType((*LdpDiscoveryAdjGrpInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.stats.stat.ldp_discovery_adj_grp_info")
}

func init() { proto.RegisterFile("ldp_discovery_adj_grp_info.proto", fileDescriptor_8555183d56035bcf) }

var fileDescriptor_8555183d56035bcf = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x59, 0x3f, 0xfa, 0x11, 0x54, 0x70, 0x51, 0xba, 0xea, 0xc1, 0x52, 0x3c, 0xd4, 0x4b,
	0x0e, 0xd6, 0x3e, 0x80, 0x14, 0x91, 0x22, 0x28, 0xb4, 0x2a, 0x78, 0x1a, 0xb2, 0xc9, 0xec, 0x76,
	0x4b, 0x36, 0x09, 0x49, 0xb6, 0xd8, 0xf7, 0xf0, 0x7d, 0x7c, 0x35, 0x49, 0xaa, 0xbd, 0xf5, 0x32,
	0x90, 0xff, 0xef, 0x37, 0x33, 0x21, 0x21, 0x7d, 0x29, 0x0c, 0x88, 0xca, 0x71, 0xbd, 0x42, 0xbb,
	0x06, 0x26, 0x96, 0x50, 0x5a, 0x03, 0x95, 0x2a, 0x34, 0x35, 0x56, 0x7b, 0x9d, 0x7e, 0xf0, 0x40,
	0xa1, 0xd2, 0x0e, 0xbe, 0x2c, 0xd4, 0x46, 0x3a, 0x08, 0x3d, 0xda, 0xa0, 0xa5, 0xff, 0x27, 0xaa,
	0xb4, 0x40, 0x17, 0x2b, 0x15, 0x58, 0xb0, 0x46, 0x7a, 0x58, 0xd9, 0x82, 0xb2, 0xc2, 0x51, 0x56,
	0xd0, 0xed, 0x7c, 0xea, 0x3c, 0xf3, 0x2e, 0xd6, 0xc1, 0x77, 0x42, 0xae, 0x77, 0x2f, 0x87, 0xe7,
	0xc7, 0xcf, 0x79, 0x7a, 0x45, 0xba, 0x61, 0x26, 0x28, 0x56, 0x63, 0x96, 0xf4, 0x93, 0x61, 0x77,
	0xd6, 0x09, 0xc1, 0x0b, 0xab, 0x31, 0xed, 0x91, 0x36, 0x2b, 0x36, 0x68, 0x2f, 0xa2, 0x16, 0x2b,
	0x22, 0x38, 0x27, 0x2d, 0xe9, 0x2c, 0x54, 0x22, 0xdb, 0x8f, 0xf9, 0xa1, 0x74, 0x76, 0x2a, 0xd2,
	0x1b, 0x72, 0x22, 0x59, 0x8e, 0x12, 0x9c, 0x61, 0x1c, 0x03, 0x3e, 0xe8, 0x27, 0xc3, 0xe3, 0xd9,
	0x51, 0x4c, 0xe7, 0x21, 0x9c, 0x8a, 0xc1, 0x4f, 0x42, 0x2e, 0x77, 0x5f, 0x2b, 0x1d, 0x93, 0x1e,
	0x13, 0x4b, 0xc6, 0x51, 0xf1, 0x35, 0x94, 0x56, 0x37, 0x06, 0x1a, 0x03, 0xbe, 0xaa, 0x31, 0xbb,
	0x8b, 0xd3, 0xce, 0xb6, 0xf8, 0x29, 0xd0, 0x77, 0xf3, 0x56, 0xd5, 0x18, 0x76, 0x7b, 0x1e, 0x1f,
	0x4d, 0x01, 0xd7, 0x8d, 0xf2, 0xd9, 0x68, 0xb3, 0xdb, 0x73, 0xf3, 0x6a, 0x50, 0x4d, 0x42, 0x96,
	0xde, 0x92, 0xd3, 0x60, 0x31, 0x9b, 0x03, 0x5f, 0x94, 0x7f, 0xe2, 0x7d, 0x14, 0x43, 0xfb, 0x83,
	0xcd, 0x27, 0x8b, 0x72, 0xa3, 0x5e, 0x90, 0x4e, 0x50, 0xad, 0x96, 0x98, 0x8d, 0xa3, 0xd1, 0xf6,
	0xdc, 0xcc, 0xb4, 0xc4, 0xbc, 0x15, 0xff, 0x6d, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x62, 0xa3,
	0x09, 0x27, 0xdb, 0x01, 0x00, 0x00,
}