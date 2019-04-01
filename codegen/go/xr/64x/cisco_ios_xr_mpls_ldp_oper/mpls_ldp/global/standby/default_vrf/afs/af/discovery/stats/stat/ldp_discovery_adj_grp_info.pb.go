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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_discovery_stats_stat

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
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LsrId                string   `protobuf:"bytes,2,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,3,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
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
	proto.RegisterType((*LdpDiscoveryAdjGrpInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.discovery.stats.stat.ldp_discovery_adj_grp_info_KEYS")
	proto.RegisterType((*LdpDiscoveryAdjGrpInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.discovery.stats.stat.ldp_discovery_adj_grp_info")
}

func init() { proto.RegisterFile("ldp_discovery_adj_grp_info.proto", fileDescriptor_8555183d56035bcf) }

var fileDescriptor_8555183d56035bcf = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x19, 0xc5, 0x56, 0x83, 0x0a, 0x0e, 0x4a, 0x47, 0x37, 0x96, 0xe2, 0xa2, 0x6e, 0xb2,
	0xb0, 0xf6, 0x01, 0xa4, 0x88, 0x14, 0x41, 0xa1, 0xd5, 0x85, 0x6e, 0x2e, 0x77, 0xf2, 0x33, 0x9d,
	0x92, 0x49, 0x42, 0x92, 0xa9, 0xf6, 0xe9, 0x7c, 0x35, 0x49, 0xaa, 0xdd, 0x75, 0x13, 0xc8, 0x39,
	0xdf, 0xcd, 0x07, 0xb9, 0xa4, 0xaf, 0xb8, 0x05, 0x5e, 0x7b, 0x66, 0x56, 0xc2, 0xad, 0x01, 0xf9,
	0x12, 0x2a, 0x67, 0xa1, 0xd6, 0xd2, 0x50, 0xeb, 0x4c, 0x30, 0xf9, 0x27, 0x8b, 0x2d, 0xd4, 0xc6,
	0xc3, 0xb7, 0x83, 0xc6, 0x2a, 0x0f, 0x71, 0xc6, 0x58, 0xe1, 0xe8, 0xff, 0x8d, 0x56, 0xca, 0x94,
	0xa8, 0xa8, 0x0f, 0xa8, 0x79, 0xb9, 0xa6, 0x5c, 0x48, 0x6c, 0x55, 0x80, 0x95, 0x93, 0x14, 0xa5,
	0xa7, 0x28, 0xe9, 0xd6, 0x11, 0xa1, 0xe0, 0xd3, 0x39, 0xf8, 0x22, 0xd7, 0xbb, 0xfd, 0xf0, 0xfc,
	0xf8, 0x31, 0xcf, 0x7b, 0xa4, 0x8b, 0x12, 0x34, 0x36, 0xa2, 0xc8, 0xfa, 0xd9, 0xf0, 0x68, 0xd6,
	0x41, 0xf9, 0x82, 0x8d, 0xc8, 0x2f, 0x48, 0x47, 0x79, 0x07, 0x35, 0x2f, 0xf6, 0x52, 0x7e, 0xa0,
	0xbc, 0x9b, 0xf2, 0xfc, 0x86, 0x9c, 0x2a, 0x2c, 0x85, 0x02, 0x6f, 0x91, 0x89, 0x58, 0xef, 0xf7,
	0xb3, 0xe1, 0xc9, 0xec, 0x38, 0xa5, 0xf3, 0x18, 0x4e, 0xf9, 0xe0, 0x27, 0x23, 0x57, 0xbb, 0xcd,
	0xf9, 0x98, 0xf4, 0x90, 0x2f, 0x91, 0x09, 0xcd, 0xd6, 0x50, 0x39, 0xd3, 0x5a, 0x68, 0x2d, 0x84,
	0xba, 0x11, 0xc5, 0x5d, 0x7a, 0xed, 0x7c, 0x5b, 0x3f, 0xc5, 0xf6, 0xdd, 0xbe, 0xd5, 0x8d, 0x88,
	0xee, 0xc0, 0xd2, 0xd7, 0x68, 0x60, 0xa6, 0xd5, 0xa1, 0x18, 0x6d, 0xdc, 0x81, 0xd9, 0x57, 0x2b,
	0xf4, 0x24, 0x66, 0xf9, 0x2d, 0x39, 0x8b, 0x14, 0xba, 0x12, 0xd8, 0xa2, 0xfa, 0x03, 0xef, 0x13,
	0x18, 0xc7, 0x1f, 0x5c, 0x39, 0x59, 0x54, 0x1b, 0xf4, 0x92, 0x1c, 0x46, 0xd4, 0x19, 0x25, 0x8a,
	0x71, 0x22, 0xba, 0x81, 0xd9, 0x99, 0x51, 0xa2, 0xec, 0xa4, 0xed, 0x8c, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xbc, 0x7d, 0x06, 0xbd, 0xc1, 0x01, 0x00, 0x00,
}
