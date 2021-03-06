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

package cisco_ios_xr_l2vpn_oper_l2vpn_xconnect_mp2mp_ce2ces_xconnect_mp2mp_ce2ce

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
	proto.RegisterType((*L2VpnXcCe2Ce_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.xconnect_mp2mp_ce2ces.xconnect_mp2mp_ce2ce.l2vpn_xc_ce2ce_KEYS")
	proto.RegisterType((*L2VpnXcCe2Ce)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.xconnect_mp2mp_ce2ces.xconnect_mp2mp_ce2ce.l2vpn_xc_ce2ce")
}

func init() { proto.RegisterFile("l2vpn_xc_ce2ce.proto", fileDescriptor_2e1d5743ad06809a) }

var fileDescriptor_2e1d5743ad06809a = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xf4, 0x30,
	0x10, 0x86, 0xc9, 0xf7, 0x89, 0xee, 0xce, 0xa2, 0x87, 0xec, 0xa2, 0xf5, 0xb0, 0xba, 0xec, 0xa9,
	0xa7, 0x1e, 0x52, 0xfd, 0x01, 0x22, 0x0b, 0x8a, 0xe8, 0xa1, 0x9e, 0x3c, 0x0d, 0x75, 0x32, 0x94,
	0x42, 0xd3, 0x84, 0xb4, 0x2b, 0xfd, 0x21, 0x82, 0x7f, 0x57, 0x36, 0xa9, 0x87, 0x8a, 0xb7, 0xe4,
	0x7d, 0xe6, 0x99, 0x37, 0x10, 0x58, 0x35, 0xea, 0xc3, 0xb5, 0x38, 0x10, 0x12, 0x2b, 0xe2, 0xcc,
	0x79, 0xdb, 0x5b, 0xf9, 0x40, 0x75, 0x47, 0x16, 0x6b, 0xdb, 0xe1, 0xe0, 0x31, 0x8e, 0x58, 0xc7,
	0x3e, 0x0b, 0xc7, 0x6c, 0x20, 0xdb, 0xb6, 0x4c, 0x3d, 0x1a, 0xa7, 0x8c, 0x8b, 0x66, 0xf7, 0x67,
	0xba, 0xfd, 0x14, 0xb0, 0x9c, 0x56, 0xe0, 0xd3, 0xee, 0xed, 0x55, 0xae, 0x01, 0x2a, 0x6f, 0xf7,
	0x0e, 0xdb, 0xd2, 0x70, 0x22, 0x36, 0x22, 0x9d, 0x17, 0xf3, 0x90, 0xbc, 0x94, 0x86, 0xe5, 0x15,
	0x2c, 0x8c, 0x53, 0x68, 0x46, 0xfe, 0x2f, 0x72, 0xe3, 0xd4, 0x73, 0xe4, 0x6b, 0x80, 0xc6, 0x52,
	0xd9, 0x20, 0x71, 0xad, 0x93, 0xff, 0x1b, 0x91, 0x9e, 0x16, 0xf3, 0x90, 0xdc, 0x73, 0xad, 0xe5,
	0x35, 0x2c, 0x3c, 0x1b, 0xdb, 0x73, 0xe4, 0x47, 0x81, 0x43, 0x8c, 0x0e, 0x03, 0xdb, 0x2f, 0x01,
	0x67, 0xd3, 0x67, 0xc9, 0x1c, 0xce, 0xc7, 0x95, 0xfb, 0xae, 0xb7, 0x86, 0x3d, 0xb2, 0xae, 0x18,
	0x6b, 0x9d, 0xa8, 0xa0, 0x2f, 0xe3, 0xfa, 0x11, 0xee, 0x74, 0xc5, 0x8f, 0x5a, 0xde, 0xc2, 0xc5,
	0x4f, 0xd1, 0x6f, 0x2b, 0x0f, 0xd6, 0x6a, 0x2c, 0x9d, 0x6a, 0x97, 0x30, 0x23, 0xc6, 0x52, 0x6b,
	0xd6, 0xc9, 0xcd, 0x46, 0xa4, 0xb3, 0xe2, 0x84, 0xf8, 0xee, 0x70, 0x7d, 0x3f, 0x0e, 0x3f, 0x90,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xb8, 0x42, 0xf6, 0x99, 0x01, 0x00, 0x00,
}
