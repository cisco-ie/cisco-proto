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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_indexes_index

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
	PoolId               string   `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
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
	proto.RegisterType((*L2VpnIndexInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.indexes.index.l2vpn_index_info_KEYS")
	proto.RegisterType((*L2VpnIndexInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.indexes.index.l2vpn_index_info")
}

func init() { proto.RegisterFile("l2vpn_index_info.proto", fileDescriptor_c514637533b1d1a8) }

var fileDescriptor_c514637533b1d1a8 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0x33, 0x31,
	0x14, 0xc5, 0xe9, 0xb7, 0xe8, 0x67, 0x6f, 0xad, 0xd6, 0x80, 0x1a, 0x50, 0xb0, 0xd4, 0x4d, 0xdd,
	0x0c, 0xd2, 0x5a, 0xf5, 0x05, 0x5c, 0x0c, 0xa2, 0x8b, 0x56, 0x41, 0x57, 0x97, 0x74, 0x12, 0xcb,
	0x85, 0xc9, 0x1f, 0x92, 0x69, 0x99, 0xf6, 0x1d, 0x7c, 0x67, 0x99, 0xcc, 0x60, 0xa5, 0xab, 0xe4,
	0x1e, 0x7e, 0xe7, 0xdc, 0xc3, 0x85, 0xb3, 0x7c, 0xbc, 0x76, 0x06, 0xc9, 0x48, 0x55, 0x22, 0x99,
	0x2f, 0x9b, 0x38, 0x6f, 0x0b, 0xcb, 0xa6, 0x19, 0x85, 0xcc, 0x22, 0xd9, 0x80, 0xa5, 0xc7, 0x1a,
	0xb2, 0x4e, 0xf9, 0x24, 0x7e, 0xd7, 0xe3, 0x24, 0x14, 0xc2, 0xc8, 0xc5, 0x26, 0x89, 0x4e, 0x15,
	0xea, 0x77, 0x78, 0x0b, 0xa7, 0xfb, 0x81, 0xf8, 0xfc, 0xf4, 0x39, 0x67, 0xe7, 0xf0, 0xdf, 0x59,
	0x9b, 0x23, 0x49, 0xde, 0x1a, 0xb4, 0x46, 0x9d, 0x59, 0xbb, 0x1a, 0x53, 0x39, 0xfc, 0xfe, 0x07,
	0xfd, 0x7d, 0x0b, 0xbb, 0x04, 0x68, 0x68, 0x2c, 0x3d, 0x1f, 0x0f, 0x5a, 0xa3, 0xde, 0xec, 0xa0,
	0x36, 0x7c, 0x78, 0x76, 0x03, 0x7d, 0xe1, 0x5c, 0x4e, 0x99, 0x28, 0xc8, 0x1a, 0x2c, 0x36, 0x4e,
	0xf1, 0x49, 0x0c, 0x3d, 0xfe, 0xa3, 0xbf, 0x6d, 0x9c, 0x62, 0xd7, 0xd0, 0x13, 0x79, 0x6e, 0x33,
	0x51, 0x28, 0x89, 0x24, 0x03, 0xbf, 0x8b, 0x59, 0x87, 0xbf, 0x62, 0x2a, 0x03, 0xbb, 0x82, 0xee,
	0xd6, 0xea, 0x05, 0x35, 0xc8, 0x34, 0x22, 0xd0, 0x48, 0x15, 0x70, 0x01, 0x9d, 0x58, 0x27, 0xd0,
	0x56, 0xf1, 0xfb, 0x5d, 0x9b, 0x39, 0x6d, 0xab, 0x15, 0x47, 0x5a, 0x94, 0x68, 0x56, 0xba, 0xaa,
	0xab, 0x97, 0x9e, 0x3f, 0x44, 0xa2, 0xab, 0x45, 0xf9, 0xba, 0xd2, 0xa9, 0x7c, 0x59, 0x7a, 0x36,
	0x82, 0x93, 0x1d, 0x80, 0x64, 0x70, 0x15, 0x14, 0x7f, 0x8c, 0x5c, 0xcf, 0x34, 0x50, 0x6a, 0xde,
	0x83, 0x5a, 0xb4, 0xe3, 0xfd, 0x27, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x05, 0x79, 0x1a,
	0x99, 0x01, 0x00, 0x00,
}
