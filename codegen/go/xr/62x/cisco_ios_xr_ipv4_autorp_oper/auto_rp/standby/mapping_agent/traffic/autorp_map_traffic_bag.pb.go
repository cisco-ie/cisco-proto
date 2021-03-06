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
// source: autorp_map_traffic_bag.proto

package cisco_ios_xr_ipv4_autorp_oper_auto_rp_standby_mapping_agent_traffic

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

type AutorpMapTrafficBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutorpMapTrafficBag_KEYS) Reset()         { *m = AutorpMapTrafficBag_KEYS{} }
func (m *AutorpMapTrafficBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AutorpMapTrafficBag_KEYS) ProtoMessage()    {}
func (*AutorpMapTrafficBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e7c89c2e2fba98, []int{0}
}

func (m *AutorpMapTrafficBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpMapTrafficBag_KEYS.Unmarshal(m, b)
}
func (m *AutorpMapTrafficBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpMapTrafficBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AutorpMapTrafficBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpMapTrafficBag_KEYS.Merge(m, src)
}
func (m *AutorpMapTrafficBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AutorpMapTrafficBag_KEYS.Size(m)
}
func (m *AutorpMapTrafficBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpMapTrafficBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpMapTrafficBag_KEYS proto.InternalMessageInfo

type AutorpMapTrafficBag struct {
	ActiveSentPackets      uint32   `protobuf:"varint,50,opt,name=active_sent_packets,json=activeSentPackets,proto3" json:"active_sent_packets,omitempty"`
	StandbySentPackets     uint32   `protobuf:"varint,51,opt,name=standby_sent_packets,json=standbySentPackets,proto3" json:"standby_sent_packets,omitempty"`
	ActiveReceivedPackets  uint32   `protobuf:"varint,52,opt,name=active_received_packets,json=activeReceivedPackets,proto3" json:"active_received_packets,omitempty"`
	StandbyReceivedPackets uint32   `protobuf:"varint,53,opt,name=standby_received_packets,json=standbyReceivedPackets,proto3" json:"standby_received_packets,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *AutorpMapTrafficBag) Reset()         { *m = AutorpMapTrafficBag{} }
func (m *AutorpMapTrafficBag) String() string { return proto.CompactTextString(m) }
func (*AutorpMapTrafficBag) ProtoMessage()    {}
func (*AutorpMapTrafficBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_08e7c89c2e2fba98, []int{1}
}

func (m *AutorpMapTrafficBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutorpMapTrafficBag.Unmarshal(m, b)
}
func (m *AutorpMapTrafficBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutorpMapTrafficBag.Marshal(b, m, deterministic)
}
func (m *AutorpMapTrafficBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutorpMapTrafficBag.Merge(m, src)
}
func (m *AutorpMapTrafficBag) XXX_Size() int {
	return xxx_messageInfo_AutorpMapTrafficBag.Size(m)
}
func (m *AutorpMapTrafficBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AutorpMapTrafficBag.DiscardUnknown(m)
}

var xxx_messageInfo_AutorpMapTrafficBag proto.InternalMessageInfo

func (m *AutorpMapTrafficBag) GetActiveSentPackets() uint32 {
	if m != nil {
		return m.ActiveSentPackets
	}
	return 0
}

func (m *AutorpMapTrafficBag) GetStandbySentPackets() uint32 {
	if m != nil {
		return m.StandbySentPackets
	}
	return 0
}

func (m *AutorpMapTrafficBag) GetActiveReceivedPackets() uint32 {
	if m != nil {
		return m.ActiveReceivedPackets
	}
	return 0
}

func (m *AutorpMapTrafficBag) GetStandbyReceivedPackets() uint32 {
	if m != nil {
		return m.StandbyReceivedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*AutorpMapTrafficBag_KEYS)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.standby.mapping_agent.traffic.autorp_map_traffic_bag_KEYS")
	proto.RegisterType((*AutorpMapTrafficBag)(nil), "cisco_ios_xr_ipv4_autorp_oper.auto_rp.standby.mapping_agent.traffic.autorp_map_traffic_bag")
}

func init() { proto.RegisterFile("autorp_map_traffic_bag.proto", fileDescriptor_08e7c89c2e2fba98) }

var fileDescriptor_08e7c89c2e2fba98 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc7, 0x71, 0xf6, 0xe2, 0x21, 0xe0, 0xc1, 0xa8, 0xeb, 0x82, 0x0a, 0xb2, 0x27, 0x4f, 0x41,
	0xdc, 0x55, 0xbc, 0x8b, 0x27, 0x2f, 0xb2, 0x3d, 0x79, 0x1a, 0xa6, 0xe9, 0xb4, 0x04, 0x69, 0x32,
	0x24, 0x63, 0xd1, 0x77, 0xf6, 0x21, 0xc4, 0x34, 0x15, 0xff, 0xf4, 0x56, 0xf8, 0xf6, 0x93, 0x5f,
	0x88, 0x3a, 0xc3, 0x57, 0x09, 0x91, 0xa1, 0x47, 0x06, 0x89, 0xd8, 0xb6, 0xce, 0x42, 0x8d, 0x9d,
	0xe1, 0x18, 0x24, 0xe8, 0x7b, 0xeb, 0x92, 0x0d, 0xe0, 0x42, 0x82, 0xb7, 0x08, 0x8e, 0x87, 0x2d,
	0x94, 0xff, 0x03, 0x53, 0x34, 0x5f, 0xdf, 0x10, 0xd9, 0x24, 0x41, 0xdf, 0xd4, 0xef, 0xa6, 0x47,
	0x66, 0xe7, 0x3b, 0xc0, 0x8e, 0xbc, 0x98, 0x72, 0xdc, 0xfa, 0x5c, 0x9d, 0xce, 0x8f, 0xc0, 0xe3,
	0xc3, 0x73, 0xb5, 0xfe, 0x58, 0xa8, 0xe5, 0x7c, 0xd7, 0x46, 0x1d, 0xa2, 0x15, 0x37, 0x10, 0x24,
	0xf2, 0x02, 0x8c, 0xf6, 0x85, 0x24, 0xad, 0xae, 0x2f, 0x16, 0x97, 0xfb, 0xbb, 0x83, 0x31, 0x55,
	0xe4, 0xe5, 0x69, 0x0c, 0xfa, 0x4a, 0x1d, 0x95, 0xab, 0xfc, 0x06, 0x9b, 0x0c, 0x74, 0x69, 0x3f,
	0xc5, 0xad, 0x3a, 0x29, 0x0b, 0x91, 0x2c, 0xb9, 0x81, 0x9a, 0x6f, 0xb4, 0xcd, 0xe8, 0x78, 0xcc,
	0xbb, 0x52, 0x27, 0x77, 0xa7, 0x56, 0xd3, 0xd2, 0x3f, 0x78, 0x93, 0xe1, 0xb2, 0xf4, 0x3f, 0xb2,
	0xde, 0xcb, 0x2f, 0xbb, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xed, 0x7f, 0x6b, 0x30, 0x79, 0x01,
	0x00, 0x00,
}
