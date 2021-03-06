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
// source: l2vpn_index_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_indexes_index

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

type L2VpnIndexInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PoolId               string   `protobuf:"bytes,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnIndexInfo_KEYS) Reset()         { *m = L2VpnIndexInfo_KEYS{} }
func (m *L2VpnIndexInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnIndexInfo_KEYS) ProtoMessage()    {}
func (*L2VpnIndexInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c514637533b1d1a8, []int{0}
}

func (m *L2VpnIndexInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnIndexInfo_KEYS.Unmarshal(m, b)
}
func (m *L2VpnIndexInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnIndexInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnIndexInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnIndexInfo_KEYS.Merge(m, src)
}
func (m *L2VpnIndexInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnIndexInfo_KEYS.Size(m)
}
func (m *L2VpnIndexInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnIndexInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnIndexInfo_KEYS proto.InternalMessageInfo

func (m *L2VpnIndexInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnIndexInfo_KEYS) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

type L2VpnIndexInfo struct {
	PoolIdXr             uint32   `protobuf:"varint,50,opt,name=pool_id_xr,json=poolIdXr,proto3" json:"pool_id_xr,omitempty"`
	ApplicationType      string   `protobuf:"bytes,51,opt,name=application_type,json=applicationType,proto3" json:"application_type,omitempty"`
	AllocatedIds         uint32   `protobuf:"varint,52,opt,name=allocated_ids,json=allocatedIds,proto3" json:"allocated_ids,omitempty"`
	ZombiedIds           uint32   `protobuf:"varint,53,opt,name=zombied_ids,json=zombiedIds,proto3" json:"zombied_ids,omitempty"`
	PoolSize             uint32   `protobuf:"varint,54,opt,name=pool_size,json=poolSize,proto3" json:"pool_size,omitempty"`
	MaxNumIdMgr          uint32   `protobuf:"varint,55,opt,name=max_num_id_mgr,json=maxNumIdMgr,proto3" json:"max_num_id_mgr,omitempty"`
	NumIdMgrInUse        uint32   `protobuf:"varint,56,opt,name=num_id_mgr_in_use,json=numIdMgrInUse,proto3" json:"num_id_mgr_in_use,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnIndexInfo) Reset()         { *m = L2VpnIndexInfo{} }
func (m *L2VpnIndexInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnIndexInfo) ProtoMessage()    {}
func (*L2VpnIndexInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c514637533b1d1a8, []int{1}
}

func (m *L2VpnIndexInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnIndexInfo.Unmarshal(m, b)
}
func (m *L2VpnIndexInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnIndexInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnIndexInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnIndexInfo.Merge(m, src)
}
func (m *L2VpnIndexInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnIndexInfo.Size(m)
}
func (m *L2VpnIndexInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnIndexInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnIndexInfo proto.InternalMessageInfo

func (m *L2VpnIndexInfo) GetPoolIdXr() uint32 {
	if m != nil {
		return m.PoolIdXr
	}
	return 0
}

func (m *L2VpnIndexInfo) GetApplicationType() string {
	if m != nil {
		return m.ApplicationType
	}
	return ""
}

func (m *L2VpnIndexInfo) GetAllocatedIds() uint32 {
	if m != nil {
		return m.AllocatedIds
	}
	return 0
}

func (m *L2VpnIndexInfo) GetZombiedIds() uint32 {
	if m != nil {
		return m.ZombiedIds
	}
	return 0
}

func (m *L2VpnIndexInfo) GetPoolSize() uint32 {
	if m != nil {
		return m.PoolSize
	}
	return 0
}

func (m *L2VpnIndexInfo) GetMaxNumIdMgr() uint32 {
	if m != nil {
		return m.MaxNumIdMgr
	}
	return 0
}

func (m *L2VpnIndexInfo) GetNumIdMgrInUse() uint32 {
	if m != nil {
		return m.NumIdMgrInUse
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnIndexInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.indexes.index.l2vpn_index_info_KEYS")
	proto.RegisterType((*L2VpnIndexInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.indexes.index.l2vpn_index_info")
}

func init() { proto.RegisterFile("l2vpn_index_info.proto", fileDescriptor_c514637533b1d1a8) }

var fileDescriptor_c514637533b1d1a8 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x03, 0x0b, 0xbe, 0x8f, 0x22, 0x8a, 0x4d, 0xd4, 0x26, 0x9a, 0x48, 0x70, 0x83, 0x9b,
	0x59, 0x80, 0x7f, 0x78, 0x01, 0x17, 0x13, 0xa3, 0x0b, 0xd0, 0x44, 0x57, 0x37, 0xc3, 0xb4, 0x92,
	0x9b, 0x4c, 0x7b, 0x9b, 0x76, 0x20, 0x03, 0xef, 0xe0, 0x3b, 0x9b, 0x76, 0x26, 0x62, 0xd8, 0xcc,
	0xf4, 0x9e, 0xf3, 0xbb, 0xa7, 0x27, 0x29, 0x3b, 0x2f, 0x26, 0x1b, 0x6b, 0x00, 0x8d, 0x54, 0x15,
	0xa0, 0xf9, 0xa2, 0xc4, 0x3a, 0x2a, 0x89, 0xcf, 0x72, 0xf4, 0x39, 0x01, 0x92, 0x87, 0xca, 0x41,
	0x0d, 0x91, 0x55, 0x2e, 0x89, 0xc7, 0xcd, 0x24, 0x31, 0x24, 0x95, 0x8f, 0xdf, 0x24, 0x2e, 0x2b,
	0x5f, 0xff, 0x47, 0x29, 0x3b, 0x3b, 0xcc, 0x84, 0xe7, 0xa7, 0xcf, 0x05, 0xbf, 0x60, 0xff, 0x02,
	0x0e, 0x28, 0x45, 0x6b, 0xd8, 0x1a, 0x77, 0xe7, 0x9d, 0x30, 0xa6, 0x32, 0x18, 0x96, 0xa8, 0x08,
	0x46, 0xbb, 0x36, 0xc2, 0x98, 0xca, 0xd1, 0x77, 0x9b, 0x0d, 0x0e, 0xb3, 0xf8, 0x15, 0x63, 0x0d,
	0x0d, 0x95, 0x13, 0x93, 0x61, 0x6b, 0xdc, 0x9f, 0xff, 0xaf, 0x17, 0x3e, 0x1c, 0xbf, 0x65, 0x83,
	0xcc, 0xda, 0x02, 0xf3, 0xac, 0x44, 0x32, 0x50, 0x6e, 0xad, 0x12, 0xd3, 0x18, 0x7a, 0xf2, 0x47,
	0x7f, 0xdb, 0x5a, 0xc5, 0x6f, 0x58, 0x3f, 0x2b, 0x0a, 0xca, 0xb3, 0x52, 0x49, 0x40, 0xe9, 0xc5,
	0x5d, 0xcc, 0x3a, 0xfa, 0x15, 0x53, 0xe9, 0xf9, 0x35, 0xeb, 0xed, 0x48, 0x2f, 0xb1, 0x41, 0xee,
	0x23, 0xc2, 0x1a, 0x29, 0x00, 0x97, 0xac, 0x1b, 0xeb, 0x78, 0xdc, 0x29, 0xf1, 0xb0, 0x6f, 0xb3,
	0xc0, 0x5d, 0xb8, 0xe2, 0x58, 0x67, 0x15, 0x98, 0xb5, 0x0e, 0x75, 0xf5, 0xca, 0x89, 0xc7, 0x48,
	0xf4, 0x74, 0x56, 0xbd, 0xae, 0x75, 0x2a, 0x5f, 0x56, 0x8e, 0x8f, 0xd9, 0xe9, 0x1e, 0x00, 0x34,
	0xb0, 0xf6, 0x4a, 0xcc, 0x22, 0xd7, 0x37, 0x0d, 0x94, 0x9a, 0x77, 0xaf, 0x96, 0x9d, 0xf8, 0x36,
	0xd3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x30, 0x6d, 0x6e, 0xb5, 0x01, 0x00, 0x00,
}
