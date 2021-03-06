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
// source: udp_sh_table_bag.proto

package cisco_ios_xr_ip_udp_oper_udp_connection_nodes_node_pcb_details_pcb_detail

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

type UdpShTableBag_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PcbAddress           string   `protobuf:"bytes,2,opt,name=pcb_address,json=pcbAddress,proto3" json:"pcb_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UdpShTableBag_KEYS) Reset()         { *m = UdpShTableBag_KEYS{} }
func (m *UdpShTableBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*UdpShTableBag_KEYS) ProtoMessage()    {}
func (*UdpShTableBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbae815597458c95, []int{0}
}

func (m *UdpShTableBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShTableBag_KEYS.Unmarshal(m, b)
}
func (m *UdpShTableBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShTableBag_KEYS.Marshal(b, m, deterministic)
}
func (m *UdpShTableBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShTableBag_KEYS.Merge(m, src)
}
func (m *UdpShTableBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_UdpShTableBag_KEYS.Size(m)
}
func (m *UdpShTableBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShTableBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShTableBag_KEYS proto.InternalMessageInfo

func (m *UdpShTableBag_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *UdpShTableBag_KEYS) GetPcbAddress() string {
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
	return fileDescriptor_dbae815597458c95, []int{1}
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

type UdpPcbFlagsBag_ struct {
	RecvOpts             bool     `protobuf:"varint,1,opt,name=recv_opts,json=recvOpts,proto3" json:"recv_opts,omitempty"`
	RecvRetOpts          bool     `protobuf:"varint,2,opt,name=recv_ret_opts,json=recvRetOpts,proto3" json:"recv_ret_opts,omitempty"`
	RecvDestAddr         bool     `protobuf:"varint,3,opt,name=recv_dest_addr,json=recvDestAddr,proto3" json:"recv_dest_addr,omitempty"`
	HeaderInclude        bool     `protobuf:"varint,4,opt,name=header_include,json=headerInclude,proto3" json:"header_include,omitempty"`
	RecvIntf             bool     `protobuf:"varint,5,opt,name=recv_intf,json=recvIntf,proto3" json:"recv_intf,omitempty"`
	RecvHeader           bool     `protobuf:"varint,6,opt,name=recv_header,json=recvHeader,proto3" json:"recv_header,omitempty"`
	McastLoopback        bool     `protobuf:"varint,7,opt,name=mcast_loopback,json=mcastLoopback,proto3" json:"mcast_loopback,omitempty"`
	RecvL2Header         bool     `protobuf:"varint,8,opt,name=recv_l2_header,json=recvL2Header,proto3" json:"recv_l2_header,omitempty"`
	RecvPacketInfo       bool     `protobuf:"varint,9,opt,name=recv_packet_info,json=recvPacketInfo,proto3" json:"recv_packet_info,omitempty"`
	RouterAlert          bool     `protobuf:"varint,10,opt,name=router_alert,json=routerAlert,proto3" json:"router_alert,omitempty"`
	RecvHopLimit         bool     `protobuf:"varint,11,opt,name=recv_hop_limit,json=recvHopLimit,proto3" json:"recv_hop_limit,omitempty"`
	RecvRoutingHeader    bool     `protobuf:"varint,12,opt,name=recv_routing_header,json=recvRoutingHeader,proto3" json:"recv_routing_header,omitempty"`
	RecvHopHeader        bool     `protobuf:"varint,13,opt,name=recv_hop_header,json=recvHopHeader,proto3" json:"recv_hop_header,omitempty"`
	RecvDestHeader       bool     `protobuf:"varint,14,opt,name=recv_dest_header,json=recvDestHeader,proto3" json:"recv_dest_header,omitempty"`
	RecvTrafficClass     bool     `protobuf:"varint,15,opt,name=recv_traffic_class,json=recvTrafficClass,proto3" json:"recv_traffic_class,omitempty"`
	RecvIpSec            bool     `protobuf:"varint,16,opt,name=recv_ip_sec,json=recvIpSec,proto3" json:"recv_ip_sec,omitempty"`
	RecvTabelId          bool     `protobuf:"varint,17,opt,name=recv_tabel_id,json=recvTabelId,proto3" json:"recv_tabel_id,omitempty"`
	RecvPakPriority      bool     `protobuf:"varint,18,opt,name=recv_pak_priority,json=recvPakPriority,proto3" json:"recv_pak_priority,omitempty"`
	ConnLimit            bool     `protobuf:"varint,19,opt,name=conn_limit,json=connLimit,proto3" json:"conn_limit,omitempty"`
	OptHandled           bool     `protobuf:"varint,20,opt,name=opt_handled,json=optHandled,proto3" json:"opt_handled,omitempty"`
	BindLocal            bool     `protobuf:"varint,21,opt,name=bind_local,json=bindLocal,proto3" json:"bind_local,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UdpPcbFlagsBag_) Reset()         { *m = UdpPcbFlagsBag_{} }
func (m *UdpPcbFlagsBag_) String() string { return proto.CompactTextString(m) }
func (*UdpPcbFlagsBag_) ProtoMessage()    {}
func (*UdpPcbFlagsBag_) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbae815597458c95, []int{2}
}

func (m *UdpPcbFlagsBag_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpPcbFlagsBag_.Unmarshal(m, b)
}
func (m *UdpPcbFlagsBag_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpPcbFlagsBag_.Marshal(b, m, deterministic)
}
func (m *UdpPcbFlagsBag_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpPcbFlagsBag_.Merge(m, src)
}
func (m *UdpPcbFlagsBag_) XXX_Size() int {
	return xxx_messageInfo_UdpPcbFlagsBag_.Size(m)
}
func (m *UdpPcbFlagsBag_) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpPcbFlagsBag_.DiscardUnknown(m)
}

var xxx_messageInfo_UdpPcbFlagsBag_ proto.InternalMessageInfo

func (m *UdpPcbFlagsBag_) GetRecvOpts() bool {
	if m != nil {
		return m.RecvOpts
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvRetOpts() bool {
	if m != nil {
		return m.RecvRetOpts
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvDestAddr() bool {
	if m != nil {
		return m.RecvDestAddr
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetHeaderInclude() bool {
	if m != nil {
		return m.HeaderInclude
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvIntf() bool {
	if m != nil {
		return m.RecvIntf
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvHeader() bool {
	if m != nil {
		return m.RecvHeader
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetMcastLoopback() bool {
	if m != nil {
		return m.McastLoopback
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvL2Header() bool {
	if m != nil {
		return m.RecvL2Header
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvPacketInfo() bool {
	if m != nil {
		return m.RecvPacketInfo
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRouterAlert() bool {
	if m != nil {
		return m.RouterAlert
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvHopLimit() bool {
	if m != nil {
		return m.RecvHopLimit
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvRoutingHeader() bool {
	if m != nil {
		return m.RecvRoutingHeader
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvHopHeader() bool {
	if m != nil {
		return m.RecvHopHeader
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvDestHeader() bool {
	if m != nil {
		return m.RecvDestHeader
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvTrafficClass() bool {
	if m != nil {
		return m.RecvTrafficClass
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvIpSec() bool {
	if m != nil {
		return m.RecvIpSec
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvTabelId() bool {
	if m != nil {
		return m.RecvTabelId
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetRecvPakPriority() bool {
	if m != nil {
		return m.RecvPakPriority
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetConnLimit() bool {
	if m != nil {
		return m.ConnLimit
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetOptHandled() bool {
	if m != nil {
		return m.OptHandled
	}
	return false
}

func (m *UdpPcbFlagsBag_) GetBindLocal() bool {
	if m != nil {
		return m.BindLocal
	}
	return false
}

type UdpShTableBag struct {
	AfName               string           `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LocalProcessId       uint32           `protobuf:"varint,51,opt,name=local_process_id,json=localProcessId,proto3" json:"local_process_id,omitempty"`
	LocalAddress         *UdpAddressType  `protobuf:"bytes,52,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	ForeignAddress       *UdpAddressType  `protobuf:"bytes,53,opt,name=foreign_address,json=foreignAddress,proto3" json:"foreign_address,omitempty"`
	LocalPort            uint32           `protobuf:"varint,54,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	ForeignPort          uint32           `protobuf:"varint,55,opt,name=foreign_port,json=foreignPort,proto3" json:"foreign_port,omitempty"`
	ReceiveQueue         uint32           `protobuf:"varint,56,opt,name=receive_queue,json=receiveQueue,proto3" json:"receive_queue,omitempty"`
	SendQueue            uint32           `protobuf:"varint,57,opt,name=send_queue,json=sendQueue,proto3" json:"send_queue,omitempty"`
	VrfId                uint32           `protobuf:"varint,58,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	PcbFlags             *UdpPcbFlagsBag_ `protobuf:"bytes,59,opt,name=pcb_flags,json=pcbFlags,proto3" json:"pcb_flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UdpShTableBag) Reset()         { *m = UdpShTableBag{} }
func (m *UdpShTableBag) String() string { return proto.CompactTextString(m) }
func (*UdpShTableBag) ProtoMessage()    {}
func (*UdpShTableBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbae815597458c95, []int{3}
}

func (m *UdpShTableBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UdpShTableBag.Unmarshal(m, b)
}
func (m *UdpShTableBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UdpShTableBag.Marshal(b, m, deterministic)
}
func (m *UdpShTableBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UdpShTableBag.Merge(m, src)
}
func (m *UdpShTableBag) XXX_Size() int {
	return xxx_messageInfo_UdpShTableBag.Size(m)
}
func (m *UdpShTableBag) XXX_DiscardUnknown() {
	xxx_messageInfo_UdpShTableBag.DiscardUnknown(m)
}

var xxx_messageInfo_UdpShTableBag proto.InternalMessageInfo

func (m *UdpShTableBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *UdpShTableBag) GetLocalProcessId() uint32 {
	if m != nil {
		return m.LocalProcessId
	}
	return 0
}

func (m *UdpShTableBag) GetLocalAddress() *UdpAddressType {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *UdpShTableBag) GetForeignAddress() *UdpAddressType {
	if m != nil {
		return m.ForeignAddress
	}
	return nil
}

func (m *UdpShTableBag) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *UdpShTableBag) GetForeignPort() uint32 {
	if m != nil {
		return m.ForeignPort
	}
	return 0
}

func (m *UdpShTableBag) GetReceiveQueue() uint32 {
	if m != nil {
		return m.ReceiveQueue
	}
	return 0
}

func (m *UdpShTableBag) GetSendQueue() uint32 {
	if m != nil {
		return m.SendQueue
	}
	return 0
}

func (m *UdpShTableBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *UdpShTableBag) GetPcbFlags() *UdpPcbFlagsBag_ {
	if m != nil {
		return m.PcbFlags
	}
	return nil
}

func init() {
	proto.RegisterType((*UdpShTableBag_KEYS)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_details.pcb_detail.udp_sh_table_bag_KEYS")
	proto.RegisterType((*UdpAddressType)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_details.pcb_detail.udp_address_type")
	proto.RegisterType((*UdpPcbFlagsBag_)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_details.pcb_detail.udp_pcb_flags_bag_")
	proto.RegisterType((*UdpShTableBag)(nil), "cisco_ios_xr_ip_udp_oper.udp_connection.nodes.node.pcb_details.pcb_detail.udp_sh_table_bag")
}

func init() { proto.RegisterFile("udp_sh_table_bag.proto", fileDescriptor_dbae815597458c95) }

var fileDescriptor_dbae815597458c95 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0x4b, 0x8f, 0x5b, 0x35,
	0x14, 0xc7, 0x95, 0xd2, 0xa6, 0x89, 0xf3, 0x9c, 0x5b, 0x06, 0x2c, 0xa1, 0xd2, 0x21, 0x3c, 0x14,
	0x21, 0x94, 0x45, 0x5a, 0x86, 0xd7, 0xaa, 0xe2, 0xa1, 0xb9, 0x22, 0x82, 0x21, 0x2d, 0x0b, 0x84,
	0x90, 0xe5, 0x6b, 0xfb, 0x26, 0x56, 0x6e, 0xae, 0x5d, 0xdb, 0x89, 0x28, 0x2b, 0x16, 0x7c, 0x0f,
	0xbe, 0x26, 0x4b, 0x74, 0x8e, 0x9d, 0x9b, 0x34, 0xeb, 0x76, 0x33, 0xba, 0xf3, 0x3b, 0x7f, 0x9f,
	0xff, 0xf1, 0xb1, 0x8f, 0x43, 0xde, 0xd9, 0x49, 0xcb, 0xfc, 0x9a, 0x05, 0x5e, 0x54, 0x8a, 0x15,
	0x7c, 0x35, 0xb3, 0xce, 0x04, 0x93, 0xe5, 0x42, 0x7b, 0x61, 0x98, 0x36, 0x9e, 0xfd, 0xe9, 0x98,
	0xb6, 0x0c, 0x74, 0xc6, 0x2a, 0x37, 0x83, 0x0f, 0x61, 0xea, 0x5a, 0x89, 0xa0, 0x4d, 0x3d, 0xab,
	0x8d, 0x54, 0x1e, 0xff, 0xce, 0xac, 0x28, 0x98, 0x54, 0x81, 0xeb, 0xca, 0x9f, 0x7c, 0x4f, 0x7e,
	0x25, 0x97, 0xe7, 0x26, 0xec, 0xc7, 0xef, 0x7f, 0x7b, 0x96, 0xbd, 0x47, 0xba, 0xb0, 0x8e, 0xd5,
	0x7c, 0xab, 0x68, 0xeb, 0xaa, 0x35, 0xed, 0x2e, 0x3b, 0x00, 0x7e, 0xe2, 0x5b, 0x95, 0x3d, 0x22,
	0x3d, 0xc8, 0xc1, 0xa5, 0x74, 0xca, 0x7b, 0x7a, 0x07, 0xc3, 0xc4, 0x8a, 0xe2, 0x69, 0x24, 0x93,
	0x17, 0x64, 0x0c, 0x69, 0x93, 0x80, 0x85, 0x97, 0x56, 0x65, 0xef, 0x92, 0xfb, 0xbc, 0x3c, 0xcd,
	0xd7, 0xe6, 0x25, 0x66, 0xfb, 0x80, 0xf4, 0xb5, 0xdd, 0x3f, 0x39, 0x4b, 0xd7, 0x03, 0x96, 0xf2,
	0x25, 0xc9, 0x75, 0x23, 0x79, 0xab, 0x91, 0x5c, 0x1f, 0x2c, 0xff, 0x6d, 0x93, 0x0c, 0x3c, 0xa1,
	0xb0, 0xb2, 0xe2, 0x2b, 0x8f, 0x7b, 0x81, 0x7d, 0x38, 0x25, 0xf6, 0xcc, 0xd8, 0xe0, 0xd1, 0xb7,
	0xb3, 0xec, 0x00, 0xf8, 0xd9, 0x06, 0x9f, 0x4d, 0xc8, 0x00, 0x83, 0x4e, 0x85, 0x28, 0xb8, 0x83,
	0x82, 0x1e, 0xc0, 0xa5, 0x0a, 0xa8, 0xf9, 0x88, 0x0c, 0x51, 0x23, 0x95, 0x0f, 0xe8, 0x8f, 0xe6,
	0x9d, 0x65, 0x1f, 0xe8, 0x77, 0xca, 0x07, 0x28, 0x20, 0xfb, 0x98, 0x0c, 0xd7, 0x8a, 0x4b, 0xe5,
	0x98, 0xae, 0x45, 0xb5, 0x93, 0x8a, 0xde, 0x45, 0xd5, 0x20, 0xd2, 0x3c, 0xc2, 0xa6, 0x1a, 0x5d,
	0x87, 0x92, 0xde, 0x3b, 0x56, 0x93, 0xd7, 0xa1, 0x84, 0xae, 0x62, 0x30, 0x2e, 0xa1, 0x6d, 0x0c,
	0x13, 0x40, 0x37, 0x48, 0xc0, 0x64, 0x2b, 0xb8, 0x0f, 0xac, 0x32, 0xc6, 0x16, 0x5c, 0x6c, 0xe8,
	0xfd, 0x68, 0x82, 0x74, 0x91, 0x60, 0x53, 0x71, 0x35, 0x3f, 0xa4, 0xea, 0x1c, 0x2b, 0x5e, 0xcc,
	0x53, 0xb2, 0x29, 0x19, 0xa3, 0xca, 0x72, 0xb1, 0x51, 0x81, 0xe9, 0xba, 0x34, 0xb4, 0x8b, 0x3a,
	0x5c, 0x7d, 0x8b, 0x38, 0xaf, 0x4b, 0x03, 0xcd, 0x77, 0x66, 0x17, 0x94, 0x63, 0xbc, 0x52, 0x2e,
	0x50, 0x92, 0x9a, 0x84, 0xec, 0x29, 0xa0, 0xc6, 0x72, 0x6d, 0x2c, 0xab, 0xf4, 0x56, 0x07, 0xda,
	0x3b, 0x5a, 0xde, 0x18, 0xbb, 0x00, 0x96, 0xcd, 0xc8, 0x83, 0xd8, 0x6e, 0xb3, 0x0b, 0xba, 0x5e,
	0x1d, 0xaa, 0xeb, 0xa3, 0xf4, 0x02, 0x9b, 0x1e, 0x23, 0xa9, 0xc4, 0x4f, 0xc8, 0xa8, 0xc9, 0x9a,
	0xb4, 0x83, 0xb8, 0xe1, 0x94, 0xf6, 0x6c, 0x2b, 0x78, 0x44, 0x49, 0x38, 0x3c, 0x6e, 0x05, 0x0e,
	0x29, 0x29, 0x3f, 0x23, 0x19, 0x2a, 0x83, 0xe3, 0x65, 0xa9, 0x05, 0x13, 0x15, 0xf7, 0x9e, 0x8e,
	0x50, 0x8b, 0x39, 0x9e, 0xc7, 0xc0, 0xb7, 0xc0, 0xb3, 0xf7, 0xd3, 0x81, 0x68, 0xcb, 0xbc, 0x12,
	0x74, 0x8c, 0x32, 0x3c, 0xc0, 0xdc, 0x3e, 0x53, 0xa2, 0xb9, 0x3e, 0x81, 0x17, 0xaa, 0x62, 0x5a,
	0xd2, 0x8b, 0xe3, 0xf5, 0x79, 0x0e, 0x2c, 0x97, 0xd9, 0xa7, 0xe4, 0x22, 0xb5, 0x79, 0xc3, 0xac,
	0xd3, 0xc6, 0xe9, 0xf0, 0x92, 0x66, 0xa8, 0x1b, 0xc5, 0x3e, 0x6f, 0x6e, 0x13, 0xce, 0x1e, 0x12,
	0x02, 0xc3, 0x9b, 0x3a, 0xf8, 0x20, 0xda, 0x01, 0x89, 0xed, 0x7b, 0x44, 0x7a, 0xc6, 0x06, 0xb6,
	0xe6, 0xb5, 0xac, 0x94, 0xa4, 0x6f, 0xc7, 0xfb, 0x61, 0x6c, 0xb8, 0x89, 0x04, 0xd6, 0x17, 0xba,
	0x96, 0xac, 0x32, 0x82, 0x57, 0xf4, 0x32, 0xae, 0x07, 0xb2, 0x00, 0x30, 0xf9, 0xef, 0x6e, 0x9c,
	0xca, 0xd3, 0x61, 0x3f, 0x9d, 0xca, 0xf9, 0x2b, 0x53, 0x39, 0x25, 0x63, 0xcc, 0xc3, 0xac, 0x33,
	0x02, 0x86, 0x58, 0x4b, 0xfa, 0xf8, 0xaa, 0x35, 0x1d, 0x2c, 0x87, 0xc8, 0x6f, 0x23, 0xce, 0x65,
	0xf6, 0x77, 0x8b, 0x0c, 0xa2, 0xf4, 0x30, 0x9e, 0x4f, 0xae, 0x5a, 0xd3, 0xde, 0xfc, 0xf7, 0xd9,
	0x6b, 0x7b, 0xa7, 0x66, 0xe7, 0xaf, 0xc9, 0xb2, 0x8f, 0x8e, 0x87, 0xf7, 0xe1, 0x9f, 0x16, 0x19,
	0x95, 0xc6, 0x29, 0xbd, 0xaa, 0x9b, 0x22, 0x3e, 0x7f, 0xf3, 0x45, 0x0c, 0x93, 0xe7, 0xa1, 0x8c,
	0x87, 0x84, 0xa4, 0x9e, 0x19, 0x17, 0xe8, 0x35, 0x76, 0xab, 0x1b, 0xbb, 0x65, 0x5c, 0x80, 0x41,
	0x3a, 0x14, 0x89, 0x82, 0x2f, 0x50, 0xd0, 0x4b, 0x0c, 0x25, 0x1f, 0xe2, 0x95, 0x52, 0x7a, 0xaf,
	0xd8, 0x8b, 0x9d, 0xda, 0x29, 0xfa, 0x25, 0x6a, 0xfa, 0x09, 0xfe, 0x02, 0x0c, 0x6c, 0xbc, 0xaa,
	0x65, 0x52, 0x7c, 0x15, 0x6d, 0x80, 0xc4, 0xf0, 0x25, 0x69, 0xef, 0x5d, 0x09, 0xe7, 0xf5, 0x35,
	0x86, 0xee, 0xed, 0x5d, 0x99, 0xcb, 0xec, 0x2f, 0xd2, 0x6d, 0xde, 0x46, 0xfa, 0x0d, 0x36, 0xe7,
	0x8f, 0xd7, 0xdc, 0x9c, 0x57, 0xdf, 0xde, 0x65, 0xc7, 0x8a, 0xe2, 0x07, 0xf8, 0xb7, 0x68, 0xe3,
	0x0f, 0xd7, 0xe3, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x81, 0x6d, 0xb3, 0xd2, 0x06, 0x00,
	0x00,
}
