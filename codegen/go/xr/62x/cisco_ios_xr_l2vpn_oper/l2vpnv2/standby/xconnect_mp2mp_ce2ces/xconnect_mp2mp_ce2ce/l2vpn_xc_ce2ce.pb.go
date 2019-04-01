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
// source: l2vpn_xc_ce2ce.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_xconnect_mp2mp_ce2ces_xconnect_mp2mp_ce2ce

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

type L2VpnXcCe2Ce_KEYS struct {
	GroupName            string   `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Mp2MpName            string   `protobuf:"bytes,2,opt,name=mp2_mp_name,json=mp2MpName,proto3" json:"mp2_mp_name,omitempty"`
	LocalCeid            uint32   `protobuf:"varint,3,opt,name=local_ceid,json=localCeid,proto3" json:"local_ceid,omitempty"`
	RemoteCeid           uint32   `protobuf:"varint,4,opt,name=remote_ceid,json=remoteCeid,proto3" json:"remote_ceid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnXcCe2Ce_KEYS) Reset()         { *m = L2VpnXcCe2Ce_KEYS{} }
func (m *L2VpnXcCe2Ce_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcCe2Ce_KEYS) ProtoMessage()    {}
func (*L2VpnXcCe2Ce_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1d5743ad06809a, []int{0}
}

func (m *L2VpnXcCe2Ce_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcCe2Ce_KEYS.Unmarshal(m, b)
}
func (m *L2VpnXcCe2Ce_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcCe2Ce_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnXcCe2Ce_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcCe2Ce_KEYS.Merge(m, src)
}
func (m *L2VpnXcCe2Ce_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcCe2Ce_KEYS.Size(m)
}
func (m *L2VpnXcCe2Ce_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcCe2Ce_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcCe2Ce_KEYS proto.InternalMessageInfo

func (m *L2VpnXcCe2Ce_KEYS) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *L2VpnXcCe2Ce_KEYS) GetMp2MpName() string {
	if m != nil {
		return m.Mp2MpName
	}
	return ""
}

func (m *L2VpnXcCe2Ce_KEYS) GetLocalCeid() uint32 {
	if m != nil {
		return m.LocalCeid
	}
	return 0
}

func (m *L2VpnXcCe2Ce_KEYS) GetRemoteCeid() uint32 {
	if m != nil {
		return m.RemoteCeid
	}
	return 0
}

type L2VpnXcCe2Ce struct {
	LocalCustomerEdgeId  uint32   `protobuf:"varint,50,opt,name=local_customer_edge_id,json=localCustomerEdgeId,proto3" json:"local_customer_edge_id,omitempty"`
	RemoteCustomerEdgeId uint32   `protobuf:"varint,51,opt,name=remote_customer_edge_id,json=remoteCustomerEdgeId,proto3" json:"remote_customer_edge_id,omitempty"`
	CeAdded              bool     `protobuf:"varint,52,opt,name=ce_added,json=ceAdded,proto3" json:"ce_added,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnXcCe2Ce) Reset()         { *m = L2VpnXcCe2Ce{} }
func (m *L2VpnXcCe2Ce) String() string { return proto.CompactTextString(m) }
func (*L2VpnXcCe2Ce) ProtoMessage()    {}
func (*L2VpnXcCe2Ce) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e1d5743ad06809a, []int{1}
}

func (m *L2VpnXcCe2Ce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnXcCe2Ce.Unmarshal(m, b)
}
func (m *L2VpnXcCe2Ce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnXcCe2Ce.Marshal(b, m, deterministic)
}
func (m *L2VpnXcCe2Ce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnXcCe2Ce.Merge(m, src)
}
func (m *L2VpnXcCe2Ce) XXX_Size() int {
	return xxx_messageInfo_L2VpnXcCe2Ce.Size(m)
}
func (m *L2VpnXcCe2Ce) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnXcCe2Ce.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnXcCe2Ce proto.InternalMessageInfo

func (m *L2VpnXcCe2Ce) GetLocalCustomerEdgeId() uint32 {
	if m != nil {
		return m.LocalCustomerEdgeId
	}
	return 0
}

func (m *L2VpnXcCe2Ce) GetRemoteCustomerEdgeId() uint32 {
	if m != nil {
		return m.RemoteCustomerEdgeId
	}
	return 0
}

func (m *L2VpnXcCe2Ce) GetCeAdded() bool {
	if m != nil {
		return m.CeAdded
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnXcCe2Ce_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mp_ce2ces.xconnect_mp2mp_ce2ce.l2vpn_xc_ce2ce_KEYS")
	proto.RegisterType((*L2VpnXcCe2Ce)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.xconnect_mp2mp_ce2ces.xconnect_mp2mp_ce2ce.l2vpn_xc_ce2ce")
}

func init() { proto.RegisterFile("l2vpn_xc_ce2ce.proto", fileDescriptor_2e1d5743ad06809a) }

var fileDescriptor_2e1d5743ad06809a = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x89, 0x8a, 0xae, 0xef, 0xd0, 0x43, 0x36, 0xb4, 0x1e, 0xa6, 0x65, 0xa7, 0x9e, 0x7a,
	0x48, 0xf5, 0x03, 0x88, 0xec, 0x20, 0xa2, 0x87, 0x7a, 0xf2, 0xf4, 0xd2, 0x25, 0x2f, 0xa5, 0xd0,
	0x34, 0x21, 0xed, 0x46, 0xfd, 0x1e, 0x82, 0x5f, 0x57, 0x96, 0xd4, 0x43, 0x65, 0xb7, 0xe4, 0xf9,
	0x3d, 0x7f, 0x02, 0x81, 0x65, 0x23, 0xf6, 0xb6, 0xc5, 0x41, 0xa2, 0x24, 0x21, 0x29, 0xb3, 0xce,
	0xf4, 0x86, 0x17, 0xb2, 0xee, 0xa4, 0xc1, 0xda, 0x74, 0x38, 0x38, 0x0c, 0x16, 0x63, 0xc9, 0x65,
	0xfe, 0xb8, 0x17, 0x59, 0xd7, 0x97, 0xad, 0xda, 0x7e, 0x65, 0x83, 0x34, 0x6d, 0x4b, 0xb2, 0x47,
	0x6d, 0x85, 0xb6, 0xa1, 0xa3, 0x3b, 0xaa, 0xae, 0xbf, 0x19, 0x2c, 0xa6, 0x63, 0xf8, 0xba, 0xf9,
	0xfc, 0xe0, 0x2b, 0x80, 0xca, 0x99, 0x9d, 0xc5, 0xb6, 0xd4, 0x14, 0xb3, 0x84, 0xa5, 0x51, 0x11,
	0x79, 0xe5, 0xbd, 0xd4, 0xc4, 0xef, 0x60, 0xae, 0xad, 0x40, 0x3d, 0xf2, 0x93, 0xc0, 0xb5, 0x15,
	0x6f, 0x81, 0xaf, 0x00, 0x1a, 0x23, 0xcb, 0x06, 0x25, 0xd5, 0x2a, 0x3e, 0x4d, 0x58, 0x7a, 0x59,
	0x44, 0x5e, 0x79, 0xa6, 0x5a, 0xf1, 0x7b, 0x98, 0x3b, 0xd2, 0xa6, 0xa7, 0xc0, 0xcf, 0x3c, 0x87,
	0x20, 0x1d, 0x0c, 0xeb, 0x1f, 0x06, 0x57, 0xd3, 0x67, 0xf1, 0x1c, 0xae, 0xc7, 0xca, 0x5d, 0xd7,
	0x1b, 0x4d, 0x0e, 0x49, 0x55, 0x84, 0xb5, 0x8a, 0x85, 0x8f, 0x2f, 0x42, 0xfd, 0x08, 0x37, 0xaa,
	0xa2, 0x17, 0xc5, 0x1f, 0xe1, 0xe6, 0x6f, 0xe8, 0x7f, 0x2a, 0xf7, 0xa9, 0xe5, 0x38, 0x3a, 0x8d,
	0xdd, 0xc2, 0x4c, 0x12, 0x96, 0x4a, 0x91, 0x8a, 0x1f, 0x12, 0x96, 0xce, 0x8a, 0x0b, 0x49, 0x4f,
	0x87, 0xeb, 0xf6, 0xdc, 0xff, 0x45, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x42, 0xa9, 0x2c,
	0xa3, 0x01, 0x00, 0x00,
}