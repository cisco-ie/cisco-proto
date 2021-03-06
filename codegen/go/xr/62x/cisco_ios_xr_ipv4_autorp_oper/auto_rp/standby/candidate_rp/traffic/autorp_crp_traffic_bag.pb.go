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
// source: autorp_crp_traffic_bag.proto

package cisco_ios_xr_ipv4_autorp_oper_auto_rp_standby_candidate_rp_traffic

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

type AutorpCrpTrafficBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutorpCrpTrafficBag_KEYS) Reset()         { *m = AutorpCrpTrafficBag_KEYS{} }
func (m *AutorpCrpTrafficBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AutorpCrpTrafficBag_KEYS) ProtoMessage()    {}
func (*AutorpCrpTrafficBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_443c66d29a49c00c, []int{0}
}

func (m *AutorpCrpTrafficBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpCrpTrafficBag_KEYS.Unmarshal(m, b)
}
func (m *AutorpCrpTrafficBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpCrpTrafficBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AutorpCrpTrafficBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpCrpTrafficBag_KEYS.Merge(m, src)
}
func (m *AutorpCrpTrafficBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AutorpCrpTrafficBag_KEYS.Size(m)
}
func (m *AutorpCrpTrafficBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpCrpTrafficBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpCrpTrafficBag_KEYS proto.InternalMessageInfo

type AutorpCrpTrafficBag struct {
	ActiveSentPackets    uint32   `protobuf:"varint,50,opt,name=active_sent_packets,json=activeSentPackets,proto3" json:"active_sent_packets,omitempty"`
	StandbySentPackets   uint32   `protobuf:"varint,51,opt,name=standby_sent_packets,json=standbySentPackets,proto3" json:"standby_sent_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutorpCrpTrafficBag) Reset()         { *m = AutorpCrpTrafficBag{} }
func (m *AutorpCrpTrafficBag) String() string { return proto.CompactTextString(m) }
func (*AutorpCrpTrafficBag) ProtoMessage()    {}
func (*AutorpCrpTrafficBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_443c66d29a49c00c, []int{1}
}

func (m *AutorpCrpTrafficBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpCrpTrafficBag.Unmarshal(m, b)
}
func (m *AutorpCrpTrafficBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpCrpTrafficBag.Marshal(b, m, deterministic)
}
func (m *AutorpCrpTrafficBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpCrpTrafficBag.Merge(m, src)
}
func (m *AutorpCrpTrafficBag) XXX_Size() int {
	return xxx_messageInfo_AutorpCrpTrafficBag.Size(m)
}
func (m *AutorpCrpTrafficBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpCrpTrafficBag.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpCrpTrafficBag proto.InternalMessageInfo

func (m *AutorpCrpTrafficBag) GetActiveSentPackets() uint32 {
	if m != nil {
		return m.ActiveSentPackets
	}
	return 0
}

func (m *AutorpCrpTrafficBag) GetStandbySentPackets() uint32 {
	if m != nil {
		return m.StandbySentPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*AutorpCrpTrafficBag_KEYS)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.standby.candidate_rp.traffic.autorp_crp_traffic_bag_KEYS")
	proto.RegisterType((*AutorpCrpTrafficBag)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.standby.candidate_rp.traffic.autorp_crp_traffic_bag")
}

func init() { proto.RegisterFile("autorp_crp_traffic_bag.proto", fileDescriptor_443c66d29a49c00c) }

var fileDescriptor_443c66d29a49c00c = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0x86, 0xf9, 0x96, 0x6f, 0x08, 0x38, 0x18, 0x45, 0x04, 0x15, 0xa4, 0x93, 0x53, 0x10, 0xeb,
	0x2f, 0x10, 0x9c, 0x5c, 0xc4, 0x4e, 0x4e, 0xc7, 0x35, 0x4d, 0x25, 0x08, 0x49, 0xb8, 0x9c, 0x45,
	0xfd, 0xf5, 0xd2, 0x36, 0x83, 0x42, 0xb7, 0xe3, 0x9e, 0x7b, 0xee, 0xee, 0x15, 0x4b, 0x7c, 0xb0,
	0xa7, 0x00, 0x9a, 0x02, 0x30, 0x61, 0x5d, 0x5b, 0x0d, 0x25, 0xde, 0x54, 0x20, 0xcf, 0x5e, 0x1e,
	0xb4, 0x8d, 0xda, 0x83, 0xf5, 0x11, 0x9e, 0x04, 0x36, 0x34, 0x7b, 0x48, 0xf3, 0x3e, 0x18, 0x52,
	0x6d, 0x0d, 0x14, 0x54, 0x64, 0x74, 0x55, 0xf9, 0x52, 0x1a, 0x5d, 0x65, 0x2b, 0x64, 0xd3, 0x36,
	0xd3, 0xb6, 0x6c, 0x25, 0x16, 0xc3, 0x37, 0xe0, 0x74, 0xbc, 0x16, 0xd9, 0x5b, 0xcc, 0x86, 0xb1,
	0x54, 0x62, 0x82, 0x9a, 0x6d, 0x63, 0x20, 0x1a, 0xc7, 0x10, 0x50, 0xdf, 0x0d, 0xc7, 0xf9, 0x6e,
	0xfd, 0xb7, 0x19, 0x5d, 0xc6, 0x3d, 0x2a, 0x8c, 0xe3, 0x73, 0x0f, 0xe4, 0x56, 0x4c, 0xd3, 0x23,
	0xbf, 0x42, 0xde, 0x09, 0x32, 0xb1, 0x2f, 0xa3, 0xfc, 0xef, 0x52, 0xe6, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x27, 0xc8, 0xc6, 0x05, 0x01, 0x00, 0x00,
}
