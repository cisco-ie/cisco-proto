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
// source: udp_sh_brief_bag.proto

package cisco_ios_xr_ip_udp_oper_udp_connection_nodes_node_pcb_briefs_pcb_brief

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

type UdpShBriefBag_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PcbAddress           string   `protobuf:"bytes,2,opt,name=pcb_address,json=pcbAddress,proto3" json:"pcb_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UdpShBriefBag_KEYS) Reset()         { *m = UdpShBriefBag_KEYS{} }
func (m *UdpShBriefBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*UdpShBriefBag_KEYS) ProtoMessage()    {}
func (*UdpShBriefBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4b431cfae5125a, []int{0}
}

func (m *UdpShBriefBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShBriefBag_KEYS.Unmarshal(m, b)
}
func (m *UdpShBriefBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShBriefBag_KEYS.Marshal(b, m, deterministic)
}
func (m *UdpShBriefBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShBriefBag_KEYS.Merge(m, src)
}
func (m *UdpShBriefBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_UdpShBriefBag_KEYS.Size(m)
}
func (m *UdpShBriefBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShBriefBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShBriefBag_KEYS proto.InternalMessageInfo

func (m *UdpShBriefBag_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *UdpShBriefBag_KEYS) GetPcbAddress() string {
	if m != nil {
		return m.PcbAddress
	}
	return ""
}

type UdpAddressType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UdpAddressType) Reset()         { *m = UdpAddressType{} }
func (m *UdpAddressType) String() string { return proto.CompactTextString(m) }
func (*UdpAddressType) ProtoMessage()    {}
func (*UdpAddressType) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4b431cfae5125a, []int{1}
}

func (m *UdpAddressType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpAddressType.Unmarshal(m, b)
}
func (m *UdpAddressType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpAddressType.Marshal(b, m, deterministic)
}
func (m *UdpAddressType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpAddressType.Merge(m, src)
}
func (m *UdpAddressType) XXX_Size() int {
	return xxx_messageInfo_UdpAddressType.Size(m)
}
func (m *UdpAddressType) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpAddressType.DiscardUnknown(m)
}

var xxx_messageInfo_UdpAddressType proto.InternalMessageInfo

func (m *UdpAddressType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *UdpAddressType) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *UdpAddressType) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type UdpShBriefBag struct {
	AfName               string          `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LocalAddress         *UdpAddressType `protobuf:"bytes,51,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ForeignAddress       *UdpAddressType `protobuf:"bytes,52,opt,name=foreign_address,json=foreignAddress,proto3" json:"foreign_address,omitempty"`
	LocalPort            uint32          `protobuf:"varint,53,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	ForeignPort          uint32          `protobuf:"varint,54,opt,name=foreign_port,json=foreignPort,proto3" json:"foreign_port,omitempty"`
	ReceiveQueue         uint32          `protobuf:"varint,55,opt,name=receive_queue,json=receiveQueue,proto3" json:"receive_queue,omitempty"`
	SendQueue            uint32          `protobuf:"varint,56,opt,name=send_queue,json=sendQueue,proto3" json:"send_queue,omitempty"`
	VrfId                uint32          `protobuf:"varint,57,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UdpShBriefBag) Reset()         { *m = UdpShBriefBag{} }
func (m *UdpShBriefBag) String() string { return proto.CompactTextString(m) }
func (*UdpShBriefBag) ProtoMessage()    {}
func (*UdpShBriefBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f4b431cfae5125a, []int{2}
}

func (m *UdpShBriefBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShBriefBag.Unmarshal(m, b)
}
func (m *UdpShBriefBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShBriefBag.Marshal(b, m, deterministic)
}
func (m *UdpShBriefBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShBriefBag.Merge(m, src)
}
func (m *UdpShBriefBag) XXX_Size() int {
	return xxx_messageInfo_UdpShBriefBag.Size(m)
}
func (m *UdpShBriefBag) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShBriefBag.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShBriefBag proto.InternalMessageInfo

func (m *UdpShBriefBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *UdpShBriefBag) GetLocalAddress() *UdpAddressType {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *UdpShBriefBag) GetForeignAddress() *UdpAddressType {
	if m != nil {
		return m.ForeignAddress
	}
	return nil
}

func (m *UdpShBriefBag) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *UdpShBriefBag) GetForeignPort() uint32 {
	if m != nil {
		return m.ForeignPort
	}
	return 0
}

func (m *UdpShBriefBag) GetReceiveQueue() uint32 {
	if m != nil {
		return m.ReceiveQueue
	}
	return 0
}

func (m *UdpShBriefBag) GetSendQueue() uint32 {
	if m != nil {
		return m.SendQueue
	}
	return 0
}

func (m *UdpShBriefBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func init() {
	proto.RegisterType((*UdpShBriefBag_KEYS)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_briefs.pcb_brief.udp_sh_brief_bag_KEYS")
	proto.RegisterType((*UdpAddressType)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_briefs.pcb_brief.udp_address_type")
	proto.RegisterType((*UdpShBriefBag)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_briefs.pcb_brief.udp_sh_brief_bag")
}

func init() { proto.RegisterFile("udp_sh_brief_bag.proto", fileDescriptor_4f4b431cfae5125a) }

var fileDescriptor_4f4b431cfae5125a = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4b, 0xeb, 0x40,
	0x10, 0xc6, 0xc9, 0x2b, 0xaf, 0xef, 0x75, 0xd2, 0xbe, 0x27, 0x81, 0x6a, 0x40, 0xc4, 0x5a, 0x2f,
	0x3d, 0xe5, 0xd0, 0xd6, 0xaa, 0x47, 0x0f, 0x22, 0x22, 0x88, 0x56, 0x3c, 0xf4, 0xb4, 0x24, 0x9b,
	0x49, 0x5d, 0x68, 0x77, 0xb7, 0xbb, 0x69, 0xd0, 0x8b, 0x07, 0xff, 0x3f, 0xff, 0x27, 0xd9, 0xcd,
	0x26, 0xd6, 0x9e, 0xf5, 0x12, 0x86, 0x6f, 0x7e, 0xf9, 0xbe, 0x99, 0x61, 0x61, 0x77, 0x9d, 0x4a,
	0xa2, 0x9f, 0x48, 0xa2, 0x18, 0x66, 0x24, 0x89, 0xe7, 0x91, 0x54, 0x22, 0x17, 0xc1, 0x15, 0x65,
	0x9a, 0x0a, 0xc2, 0x84, 0x26, 0xcf, 0x8a, 0x30, 0x49, 0x0c, 0x27, 0x24, 0xaa, 0xc8, 0x14, 0x54,
	0x70, 0x8e, 0x34, 0x67, 0x82, 0x47, 0x5c, 0xa4, 0xa8, 0xed, 0x37, 0x92, 0x34, 0x29, 0x7d, 0xf4,
	0x67, 0xd9, 0x7f, 0x84, 0xee, 0x76, 0x04, 0xb9, 0xb9, 0x9c, 0x3d, 0x04, 0xfb, 0xd0, 0x32, 0x7f,
	0x11, 0x1e, 0x2f, 0x31, 0xf4, 0x7a, 0xde, 0xa0, 0x35, 0xfd, 0x6b, 0x84, 0xdb, 0x78, 0x89, 0xc1,
	0x21, 0xf8, 0xc6, 0x22, 0x4e, 0x53, 0x85, 0x5a, 0x87, 0xbf, 0x6c, 0x1b, 0x24, 0x4d, 0x2e, 0x4a,
	0xa5, 0xbf, 0x82, 0x1d, 0x63, 0xeb, 0x00, 0x92, 0xbf, 0x48, 0x0c, 0xf6, 0xe0, 0x4f, 0x9c, 0x6d,
	0xfa, 0x35, 0xe3, 0xcc, 0xba, 0x1d, 0x41, 0x9b, 0xc9, 0x62, 0xbc, 0x65, 0xe7, 0x1b, 0xcd, 0xf9,
	0x39, 0x64, 0x52, 0x23, 0x8d, 0x1a, 0x99, 0x54, 0x91, 0xef, 0x8d, 0x32, 0x73, 0x73, 0x95, 0xcd,
	0xcc, 0xe1, 0x97, 0xcc, 0x57, 0xe8, 0x2c, 0x04, 0x8d, 0x17, 0xb5, 0xe3, 0xa8, 0xe7, 0x0d, 0xfc,
	0xe1, 0x2c, 0xfa, 0xa6, 0xc3, 0x46, 0xdb, 0xeb, 0x4f, 0xdb, 0x36, 0xaf, 0x5a, 0xe8, 0xcd, 0x83,
	0xff, 0x99, 0x50, 0xc8, 0xe6, 0xbc, 0x1e, 0x61, 0xfc, 0xd3, 0x23, 0xfc, 0x73, 0x89, 0xd5, 0x10,
	0x07, 0x00, 0xe5, 0x11, 0xa4, 0x50, 0x79, 0x78, 0xd2, 0xf3, 0x06, 0x9d, 0x69, 0xcb, 0x2a, 0x77,
	0x42, 0xe5, 0xe6, 0xe8, 0xd5, 0x88, 0x16, 0x98, 0x58, 0xc0, 0x77, 0x9a, 0x45, 0x8e, 0xa1, 0xa3,
	0x90, 0x22, 0x2b, 0x90, 0xac, 0xd6, 0xb8, 0xc6, 0xf0, 0xd4, 0x32, 0x6d, 0x27, 0xde, 0x1b, 0xcd,
	0xc4, 0x68, 0xe4, 0xa9, 0x23, 0xce, 0xca, 0x18, 0xa3, 0x94, 0xed, 0x2e, 0x34, 0x0b, 0x95, 0x11,
	0x96, 0x86, 0xe7, 0xb6, 0xf5, 0xbb, 0x50, 0xd9, 0x75, 0x9a, 0x34, 0xed, 0x4b, 0x1f, 0x7d, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x93, 0xa4, 0x12, 0x7d, 0x03, 0x03, 0x00, 0x00,
}