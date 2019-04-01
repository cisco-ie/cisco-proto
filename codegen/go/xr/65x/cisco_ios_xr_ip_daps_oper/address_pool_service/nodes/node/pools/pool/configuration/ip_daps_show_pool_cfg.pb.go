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
// source: ip_daps_show_pool_cfg.proto

package cisco_ios_xr_ip_daps_oper_address_pool_service_nodes_node_pools_pool_configuration

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

type IpDapsShowPoolCfg_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PoolName             string   `protobuf:"bytes,2,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpDapsShowPoolCfg_KEYS) Reset()         { *m = IpDapsShowPoolCfg_KEYS{} }
func (m *IpDapsShowPoolCfg_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpDapsShowPoolCfg_KEYS) ProtoMessage()    {}
func (*IpDapsShowPoolCfg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_85df0f33f2660e35, []int{0}
}

func (m *IpDapsShowPoolCfg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpDapsShowPoolCfg_KEYS.Unmarshal(m, b)
}
func (m *IpDapsShowPoolCfg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpDapsShowPoolCfg_KEYS.Marshal(b, m, deterministic)
}
func (m *IpDapsShowPoolCfg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpDapsShowPoolCfg_KEYS.Merge(m, src)
}
func (m *IpDapsShowPoolCfg_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpDapsShowPoolCfg_KEYS.Size(m)
}
func (m *IpDapsShowPoolCfg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpDapsShowPoolCfg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpDapsShowPoolCfg_KEYS proto.InternalMessageInfo

func (m *IpDapsShowPoolCfg_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *IpDapsShowPoolCfg_KEYS) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

type DapsPoolSergInfo struct {
	PreferredRole        string   `protobuf:"bytes,1,opt,name=preferred_role,json=preferredRole,proto3" json:"preferred_role,omitempty"`
	PeerDown             bool     `protobuf:"varint,2,opt,name=peer_down,json=peerDown,proto3" json:"peer_down,omitempty"`
	VerifyPend           bool     `protobuf:"varint,3,opt,name=verify_pend,json=verifyPend,proto3" json:"verify_pend,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DapsPoolSergInfo) Reset()         { *m = DapsPoolSergInfo{} }
func (m *DapsPoolSergInfo) String() string { return proto.CompactTextString(m) }
func (*DapsPoolSergInfo) ProtoMessage()    {}
func (*DapsPoolSergInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_85df0f33f2660e35, []int{1}
}

func (m *DapsPoolSergInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DapsPoolSergInfo.Unmarshal(m, b)
}
func (m *DapsPoolSergInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DapsPoolSergInfo.Marshal(b, m, deterministic)
}
func (m *DapsPoolSergInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DapsPoolSergInfo.Merge(m, src)
}
func (m *DapsPoolSergInfo) XXX_Size() int {
	return xxx_messageInfo_DapsPoolSergInfo.Size(m)
}
func (m *DapsPoolSergInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DapsPoolSergInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DapsPoolSergInfo proto.InternalMessageInfo

func (m *DapsPoolSergInfo) GetPreferredRole() string {
	if m != nil {
		return m.PreferredRole
	}
	return ""
}

func (m *DapsPoolSergInfo) GetPeerDown() bool {
	if m != nil {
		return m.PeerDown
	}
	return false
}

func (m *DapsPoolSergInfo) GetVerifyPend() bool {
	if m != nil {
		return m.VerifyPend
	}
	return false
}

type IpDapsShowPoolCfg struct {
	PoolName             string            `protobuf:"bytes,50,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	PoolId               uint32            `protobuf:"varint,51,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	VrfName              string            `protobuf:"bytes,52,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	PoolScope            string            `protobuf:"bytes,53,opt,name=pool_scope,json=poolScope,proto3" json:"pool_scope,omitempty"`
	PoolPrefixLength     uint32            `protobuf:"varint,54,opt,name=pool_prefix_length,json=poolPrefixLength,proto3" json:"pool_prefix_length,omitempty"`
	HighUtilizationMark  uint32            `protobuf:"varint,55,opt,name=high_utilization_mark,json=highUtilizationMark,proto3" json:"high_utilization_mark,omitempty"`
	LowUtilizationMark   uint32            `protobuf:"varint,56,opt,name=low_utilization_mark,json=lowUtilizationMark,proto3" json:"low_utilization_mark,omitempty"`
	CurrentUtilization   uint32            `protobuf:"varint,57,opt,name=current_utilization,json=currentUtilization,proto3" json:"current_utilization,omitempty"`
	UtilizationHighCount uint32            `protobuf:"varint,58,opt,name=utilization_high_count,json=utilizationHighCount,proto3" json:"utilization_high_count,omitempty"`
	UtilizationLowCount  uint32            `protobuf:"varint,59,opt,name=utilization_low_count,json=utilizationLowCount,proto3" json:"utilization_low_count,omitempty"`
	SergInfo             *DapsPoolSergInfo `protobuf:"bytes,60,opt,name=serg_info,json=sergInfo,proto3" json:"serg_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IpDapsShowPoolCfg) Reset()         { *m = IpDapsShowPoolCfg{} }
func (m *IpDapsShowPoolCfg) String() string { return proto.CompactTextString(m) }
func (*IpDapsShowPoolCfg) ProtoMessage()    {}
func (*IpDapsShowPoolCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_85df0f33f2660e35, []int{2}
}

func (m *IpDapsShowPoolCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpDapsShowPoolCfg.Unmarshal(m, b)
}
func (m *IpDapsShowPoolCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpDapsShowPoolCfg.Marshal(b, m, deterministic)
}
func (m *IpDapsShowPoolCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpDapsShowPoolCfg.Merge(m, src)
}
func (m *IpDapsShowPoolCfg) XXX_Size() int {
	return xxx_messageInfo_IpDapsShowPoolCfg.Size(m)
}
func (m *IpDapsShowPoolCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_IpDapsShowPoolCfg.DiscardUnknown(m)
}

var xxx_messageInfo_IpDapsShowPoolCfg proto.InternalMessageInfo

func (m *IpDapsShowPoolCfg) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *IpDapsShowPoolCfg) GetPoolId() uint32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IpDapsShowPoolCfg) GetPoolScope() string {
	if m != nil {
		return m.PoolScope
	}
	return ""
}

func (m *IpDapsShowPoolCfg) GetPoolPrefixLength() uint32 {
	if m != nil {
		return m.PoolPrefixLength
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetHighUtilizationMark() uint32 {
	if m != nil {
		return m.HighUtilizationMark
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetLowUtilizationMark() uint32 {
	if m != nil {
		return m.LowUtilizationMark
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetCurrentUtilization() uint32 {
	if m != nil {
		return m.CurrentUtilization
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetUtilizationHighCount() uint32 {
	if m != nil {
		return m.UtilizationHighCount
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetUtilizationLowCount() uint32 {
	if m != nil {
		return m.UtilizationLowCount
	}
	return 0
}

func (m *IpDapsShowPoolCfg) GetSergInfo() *DapsPoolSergInfo {
	if m != nil {
		return m.SergInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*IpDapsShowPoolCfg_KEYS)(nil), "cisco_ios_xr_ip_daps_oper.address_pool_service.nodes.node.pools.pool.configuration.ip_daps_show_pool_cfg_KEYS")
	proto.RegisterType((*DapsPoolSergInfo)(nil), "cisco_ios_xr_ip_daps_oper.address_pool_service.nodes.node.pools.pool.configuration.daps_pool_serg_info")
	proto.RegisterType((*IpDapsShowPoolCfg)(nil), "cisco_ios_xr_ip_daps_oper.address_pool_service.nodes.node.pools.pool.configuration.ip_daps_show_pool_cfg")
}

func init() { proto.RegisterFile("ip_daps_show_pool_cfg.proto", fileDescriptor_85df0f33f2660e35) }

var fileDescriptor_85df0f33f2660e35 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0x39, 0x95, 0x36, 0xb7, 0xa5, 0x22, 0x7b, 0xad, 0x46, 0x8b, 0x78, 0x1c, 0x08, 0xf7,
	0x20, 0x51, 0xae, 0xf5, 0xb7, 0x6f, 0x2a, 0x58, 0xac, 0x52, 0x52, 0x14, 0x7c, 0x5a, 0x62, 0x76,
	0x92, 0x2c, 0xcd, 0xed, 0x84, 0xd9, 0xdc, 0xa5, 0xf6, 0xd9, 0x7f, 0xcb, 0xff, 0x4d, 0x76, 0xd6,
	0xab, 0xa9, 0xde, 0x63, 0x5f, 0x96, 0xe4, 0xfb, 0xd9, 0xcf, 0xec, 0x0c, 0xd9, 0x88, 0x3d, 0xd3,
	0x28, 0x9d, 0x35, 0x4e, 0xb9, 0x0a, 0x3b, 0xd5, 0x20, 0xd6, 0x2a, 0x2f, 0xca, 0xa4, 0x21, 0x6c,
	0x51, 0xa6, 0xb9, 0x71, 0x39, 0x2a, 0x83, 0x4e, 0x9d, 0x91, 0x5a, 0xed, 0xc4, 0x06, 0x28, 0xc9,
	0xb4, 0x26, 0x70, 0x2e, 0x18, 0x0e, 0x68, 0x69, 0x72, 0x48, 0x2c, 0x6a, 0x70, 0xbc, 0x26, 0x3e,
	0x77, 0xbc, 0x26, 0x39, 0xda, 0xc2, 0x94, 0x0b, 0xca, 0x5a, 0x83, 0x76, 0xf2, 0x55, 0xdc, 0x5b,
	0x7b, 0xa4, 0xfa, 0xf8, 0xfe, 0xdb, 0x89, 0xdc, 0x13, 0x43, 0xaf, 0x2b, 0x9b, 0xcd, 0x21, 0x1e,
	0x8c, 0x07, 0xd3, 0x61, 0x1a, 0xf9, 0xe0, 0x73, 0x36, 0x07, 0x0f, 0x79, 0x37, 0xc3, 0x6b, 0x01,
	0xfa, 0xc0, 0xc3, 0xc9, 0xb9, 0x18, 0x71, 0xd1, 0x55, 0x43, 0xa5, 0x32, 0xb6, 0x40, 0xf9, 0x50,
	0xdc, 0x6c, 0x08, 0x0a, 0x20, 0x02, 0xad, 0x08, 0xeb, 0x55, 0xd5, 0xed, 0x8b, 0x34, 0xc5, 0x3a,
	0x94, 0x06, 0x20, 0xa5, 0xb1, 0xb3, 0x5c, 0x3a, 0x4a, 0x23, 0x1f, 0xbc, 0xc3, 0xce, 0xca, 0x07,
	0x62, 0x6b, 0x09, 0x64, 0x8a, 0x1f, 0xaa, 0x01, 0xab, 0xe3, 0xeb, 0x8c, 0x45, 0x88, 0x8e, 0xc1,
	0xea, 0xc9, 0xaf, 0x1b, 0x62, 0x77, 0xed, 0x50, 0x97, 0x5b, 0x9e, 0x5d, 0x6e, 0x59, 0xde, 0x11,
	0x9b, 0x0c, 0x8d, 0x8e, 0xf7, 0xc7, 0x83, 0xe9, 0x76, 0xba, 0xe1, 0x5f, 0x0f, 0xb5, 0xbc, 0x2b,
	0xa2, 0x25, 0x15, 0x41, 0x3a, 0x60, 0x69, 0x73, 0x49, 0x05, 0x3b, 0xf7, 0x85, 0x08, 0x13, 0xe6,
	0xd8, 0x40, 0xfc, 0x94, 0x21, 0x1f, 0x71, 0xe2, 0x03, 0xf9, 0x48, 0x48, 0xc6, 0x7e, 0x3a, 0x73,
	0xa6, 0x6a, 0xb0, 0x65, 0x5b, 0xc5, 0xcf, 0xb8, 0xfa, 0x2d, 0x4f, 0x8e, 0x19, 0x1c, 0x71, 0x2e,
	0x67, 0x62, 0xb7, 0x32, 0x65, 0xa5, 0x16, 0xad, 0xa9, 0xcd, 0x39, 0x7f, 0x1f, 0x35, 0xcf, 0xe8,
	0x34, 0x7e, 0xce, 0xc2, 0xc8, 0xc3, 0x2f, 0x7f, 0xd9, 0xa7, 0x8c, 0x4e, 0xe5, 0x13, 0xb1, 0x53,
	0x63, 0xf7, 0xbf, 0xf2, 0x82, 0x15, 0x59, 0x63, 0xf7, 0xaf, 0xf1, 0x58, 0x8c, 0xf2, 0x05, 0x11,
	0xd8, 0xb6, 0x6f, 0xc5, 0x2f, 0x83, 0xf0, 0x07, 0xf5, 0x24, 0x79, 0x20, 0x6e, 0xf7, 0xcb, 0x73,
	0x8b, 0x39, 0x2e, 0x6c, 0x1b, 0xbf, 0x62, 0x67, 0xa7, 0x47, 0x3f, 0x98, 0xb2, 0x7a, 0xeb, 0x99,
	0x1f, 0xa6, 0x6f, 0xf9, 0x26, 0x83, 0xf4, 0x3a, 0x0c, 0xd3, 0x83, 0x47, 0xd8, 0x05, 0xe7, 0xe7,
	0x40, 0x0c, 0x2f, 0xee, 0x4a, 0xfc, 0x66, 0x3c, 0x98, 0x6e, 0xcd, 0xca, 0xe4, 0xea, 0x6f, 0x7d,
	0xb2, 0xe6, 0x6a, 0xa6, 0x91, 0x7f, 0x3c, 0xb4, 0x05, 0x7e, 0xdf, 0xe0, 0xdf, 0x6d, 0xff, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x49, 0xdd, 0x20, 0x8d, 0x03, 0x00, 0x00,
}