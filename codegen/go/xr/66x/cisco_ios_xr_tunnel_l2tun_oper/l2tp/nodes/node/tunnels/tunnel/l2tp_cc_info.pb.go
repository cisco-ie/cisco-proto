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
// source: l2tp_cc_info.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tp_nodes_node_tunnels_tunnel

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

type L2TpCcInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LocalTunnelId        uint32   `protobuf:"varint,2,opt,name=local_tunnel_id,json=localTunnelId,proto3" json:"local_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpCcInfo_KEYS) Reset()         { *m = L2TpCcInfo_KEYS{} }
func (m *L2TpCcInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2TpCcInfo_KEYS) ProtoMessage()    {}
func (*L2TpCcInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2f62e56b429a2f1, []int{0}
}

func (m *L2TpCcInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcInfo_KEYS.Unmarshal(m, b)
}
func (m *L2TpCcInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2TpCcInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcInfo_KEYS.Merge(m, src)
}
func (m *L2TpCcInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2TpCcInfo_KEYS.Size(m)
}
func (m *L2TpCcInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcInfo_KEYS proto.InternalMessageInfo

func (m *L2TpCcInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *L2TpCcInfo_KEYS) GetLocalTunnelId() uint32 {
	if m != nil {
		return m.LocalTunnelId
	}
	return 0
}

type L2TpCcInfo struct {
	LocalAddress                      string   `protobuf:"bytes,50,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress                     string   `protobuf:"bytes,51,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	LocalPort                         uint32   `protobuf:"varint,52,opt,name=local_port,json=localPort,proto3" json:"local_port,omitempty"`
	RemotePort                        uint32   `protobuf:"varint,53,opt,name=remote_port,json=remotePort,proto3" json:"remote_port,omitempty"`
	Protocol                          uint32   `protobuf:"varint,54,opt,name=protocol,proto3" json:"protocol,omitempty"`
	IsPmtuEnabled                     bool     `protobuf:"varint,55,opt,name=is_pmtu_enabled,json=isPmtuEnabled,proto3" json:"is_pmtu_enabled,omitempty"`
	RemoteTunnelId                    uint32   `protobuf:"varint,56,opt,name=remote_tunnel_id,json=remoteTunnelId,proto3" json:"remote_tunnel_id,omitempty"`
	LocalTunnelName                   string   `protobuf:"bytes,57,opt,name=local_tunnel_name,json=localTunnelName,proto3" json:"local_tunnel_name,omitempty"`
	RemoteTunnelName                  string   `protobuf:"bytes,58,opt,name=remote_tunnel_name,json=remoteTunnelName,proto3" json:"remote_tunnel_name,omitempty"`
	ClassName                         string   `protobuf:"bytes,59,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ActiveSessions                    uint32   `protobuf:"varint,60,opt,name=active_sessions,json=activeSessions,proto3" json:"active_sessions,omitempty"`
	RetransmitTime                    []uint32 `protobuf:"varint,61,rep,packed,name=retransmit_time,json=retransmitTime,proto3" json:"retransmit_time,omitempty"`
	SequenceNs                        uint32   `protobuf:"varint,62,opt,name=sequence_ns,json=sequenceNs,proto3" json:"sequence_ns,omitempty"`
	SequenceNr                        uint32   `protobuf:"varint,63,opt,name=sequence_nr,json=sequenceNr,proto3" json:"sequence_nr,omitempty"`
	LocalWindowSize                   uint32   `protobuf:"varint,64,opt,name=local_window_size,json=localWindowSize,proto3" json:"local_window_size,omitempty"`
	RemoteWindowSize                  uint32   `protobuf:"varint,65,opt,name=remote_window_size,json=remoteWindowSize,proto3" json:"remote_window_size,omitempty"`
	RetransmissionTime                uint32   `protobuf:"varint,66,opt,name=retransmission_time,json=retransmissionTime,proto3" json:"retransmission_time,omitempty"`
	MaximumRetransmissionTime         uint32   `protobuf:"varint,67,opt,name=maximum_retransmission_time,json=maximumRetransmissionTime,proto3" json:"maximum_retransmission_time,omitempty"`
	UnsentQueueSize                   uint32   `protobuf:"varint,68,opt,name=unsent_queue_size,json=unsentQueueSize,proto3" json:"unsent_queue_size,omitempty"`
	UnsentMaximumQueueSize            uint32   `protobuf:"varint,69,opt,name=unsent_maximum_queue_size,json=unsentMaximumQueueSize,proto3" json:"unsent_maximum_queue_size,omitempty"`
	ResendQueueSize                   uint32   `protobuf:"varint,70,opt,name=resend_queue_size,json=resendQueueSize,proto3" json:"resend_queue_size,omitempty"`
	ResendMaximumQueueSize            uint32   `protobuf:"varint,71,opt,name=resend_maximum_queue_size,json=resendMaximumQueueSize,proto3" json:"resend_maximum_queue_size,omitempty"`
	OrderQueueSize                    uint32   `protobuf:"varint,72,opt,name=order_queue_size,json=orderQueueSize,proto3" json:"order_queue_size,omitempty"`
	PacketQueueCheck                  uint32   `protobuf:"varint,73,opt,name=packet_queue_check,json=packetQueueCheck,proto3" json:"packet_queue_check,omitempty"`
	DigestSecrets                     uint32   `protobuf:"varint,74,opt,name=digest_secrets,json=digestSecrets,proto3" json:"digest_secrets,omitempty"`
	Resends                           uint32   `protobuf:"varint,75,opt,name=resends,proto3" json:"resends,omitempty"`
	ZeroLengthBodyAcknowledgementSent uint32   `protobuf:"varint,76,opt,name=zero_length_body_acknowledgement_sent,json=zeroLengthBodyAcknowledgementSent,proto3" json:"zero_length_body_acknowledgement_sent,omitempty"`
	TotalOutOfOrderDropPackets        uint32   `protobuf:"varint,77,opt,name=total_out_of_order_drop_packets,json=totalOutOfOrderDropPackets,proto3" json:"total_out_of_order_drop_packets,omitempty"`
	TotalOutOfOrderReorderPackets     uint32   `protobuf:"varint,78,opt,name=total_out_of_order_reorder_packets,json=totalOutOfOrderReorderPackets,proto3" json:"total_out_of_order_reorder_packets,omitempty"`
	TotalPeerAuthenticationFailures   uint32   `protobuf:"varint,79,opt,name=total_peer_authentication_failures,json=totalPeerAuthenticationFailures,proto3" json:"total_peer_authentication_failures,omitempty"`
	IsTunnelUp                        bool     `protobuf:"varint,80,opt,name=is_tunnel_up,json=isTunnelUp,proto3" json:"is_tunnel_up,omitempty"`
	IsCongestionControlEnabled        bool     `protobuf:"varint,81,opt,name=is_congestion_control_enabled,json=isCongestionControlEnabled,proto3" json:"is_congestion_control_enabled,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *L2TpCcInfo) Reset()         { *m = L2TpCcInfo{} }
func (m *L2TpCcInfo) String() string { return proto.CompactTextString(m) }
func (*L2TpCcInfo) ProtoMessage()    {}
func (*L2TpCcInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2f62e56b429a2f1, []int{1}
}

func (m *L2TpCcInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCcInfo.Unmarshal(m, b)
}
func (m *L2TpCcInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCcInfo.Marshal(b, m, deterministic)
}
func (m *L2TpCcInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCcInfo.Merge(m, src)
}
func (m *L2TpCcInfo) XXX_Size() int {
	return xxx_messageInfo_L2TpCcInfo.Size(m)
}
func (m *L2TpCcInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCcInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCcInfo proto.InternalMessageInfo

func (m *L2TpCcInfo) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *L2TpCcInfo) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *L2TpCcInfo) GetLocalPort() uint32 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *L2TpCcInfo) GetRemotePort() uint32 {
	if m != nil {
		return m.RemotePort
	}
	return 0
}

func (m *L2TpCcInfo) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *L2TpCcInfo) GetIsPmtuEnabled() bool {
	if m != nil {
		return m.IsPmtuEnabled
	}
	return false
}

func (m *L2TpCcInfo) GetRemoteTunnelId() uint32 {
	if m != nil {
		return m.RemoteTunnelId
	}
	return 0
}

func (m *L2TpCcInfo) GetLocalTunnelName() string {
	if m != nil {
		return m.LocalTunnelName
	}
	return ""
}

func (m *L2TpCcInfo) GetRemoteTunnelName() string {
	if m != nil {
		return m.RemoteTunnelName
	}
	return ""
}

func (m *L2TpCcInfo) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *L2TpCcInfo) GetActiveSessions() uint32 {
	if m != nil {
		return m.ActiveSessions
	}
	return 0
}

func (m *L2TpCcInfo) GetRetransmitTime() []uint32 {
	if m != nil {
		return m.RetransmitTime
	}
	return nil
}

func (m *L2TpCcInfo) GetSequenceNs() uint32 {
	if m != nil {
		return m.SequenceNs
	}
	return 0
}

func (m *L2TpCcInfo) GetSequenceNr() uint32 {
	if m != nil {
		return m.SequenceNr
	}
	return 0
}

func (m *L2TpCcInfo) GetLocalWindowSize() uint32 {
	if m != nil {
		return m.LocalWindowSize
	}
	return 0
}

func (m *L2TpCcInfo) GetRemoteWindowSize() uint32 {
	if m != nil {
		return m.RemoteWindowSize
	}
	return 0
}

func (m *L2TpCcInfo) GetRetransmissionTime() uint32 {
	if m != nil {
		return m.RetransmissionTime
	}
	return 0
}

func (m *L2TpCcInfo) GetMaximumRetransmissionTime() uint32 {
	if m != nil {
		return m.MaximumRetransmissionTime
	}
	return 0
}

func (m *L2TpCcInfo) GetUnsentQueueSize() uint32 {
	if m != nil {
		return m.UnsentQueueSize
	}
	return 0
}

func (m *L2TpCcInfo) GetUnsentMaximumQueueSize() uint32 {
	if m != nil {
		return m.UnsentMaximumQueueSize
	}
	return 0
}

func (m *L2TpCcInfo) GetResendQueueSize() uint32 {
	if m != nil {
		return m.ResendQueueSize
	}
	return 0
}

func (m *L2TpCcInfo) GetResendMaximumQueueSize() uint32 {
	if m != nil {
		return m.ResendMaximumQueueSize
	}
	return 0
}

func (m *L2TpCcInfo) GetOrderQueueSize() uint32 {
	if m != nil {
		return m.OrderQueueSize
	}
	return 0
}

func (m *L2TpCcInfo) GetPacketQueueCheck() uint32 {
	if m != nil {
		return m.PacketQueueCheck
	}
	return 0
}

func (m *L2TpCcInfo) GetDigestSecrets() uint32 {
	if m != nil {
		return m.DigestSecrets
	}
	return 0
}

func (m *L2TpCcInfo) GetResends() uint32 {
	if m != nil {
		return m.Resends
	}
	return 0
}

func (m *L2TpCcInfo) GetZeroLengthBodyAcknowledgementSent() uint32 {
	if m != nil {
		return m.ZeroLengthBodyAcknowledgementSent
	}
	return 0
}

func (m *L2TpCcInfo) GetTotalOutOfOrderDropPackets() uint32 {
	if m != nil {
		return m.TotalOutOfOrderDropPackets
	}
	return 0
}

func (m *L2TpCcInfo) GetTotalOutOfOrderReorderPackets() uint32 {
	if m != nil {
		return m.TotalOutOfOrderReorderPackets
	}
	return 0
}

func (m *L2TpCcInfo) GetTotalPeerAuthenticationFailures() uint32 {
	if m != nil {
		return m.TotalPeerAuthenticationFailures
	}
	return 0
}

func (m *L2TpCcInfo) GetIsTunnelUp() bool {
	if m != nil {
		return m.IsTunnelUp
	}
	return false
}

func (m *L2TpCcInfo) GetIsCongestionControlEnabled() bool {
	if m != nil {
		return m.IsCongestionControlEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*L2TpCcInfo_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.nodes.node.tunnels.tunnel.l2tp_cc_info_KEYS")
	proto.RegisterType((*L2TpCcInfo)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.nodes.node.tunnels.tunnel.l2tp_cc_info")
}

func init() { proto.RegisterFile("l2tp_cc_info.proto", fileDescriptor_d2f62e56b429a2f1) }

var fileDescriptor_d2f62e56b429a2f1 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x6f, 0x6f, 0x1b, 0x45,
	0x10, 0xc6, 0x15, 0x90, 0x20, 0x99, 0x26, 0x24, 0x59, 0x24, 0xb4, 0x4d, 0x15, 0xd5, 0x04, 0x15,
	0xac, 0xa8, 0x0a, 0x52, 0xca, 0xbf, 0x02, 0x2d, 0xb8, 0x6e, 0x0a, 0x21, 0x6d, 0xec, 0xda, 0x45,
	0xc0, 0xab, 0xd5, 0x65, 0x6f, 0x9c, 0xac, 0x72, 0xb7, 0x7b, 0xdd, 0xdd, 0x23, 0x6d, 0xbe, 0x2a,
	0x5f, 0x06, 0xed, 0xcc, 0x9d, 0x7d, 0x6e, 0xf3, 0x26, 0xa7, 0x7b, 0xe6, 0x37, 0xcf, 0xec, 0x3c,
	0x39, 0x2f, 0x88, 0xe2, 0x30, 0x56, 0x4a, 0x6b, 0x65, 0xec, 0xcc, 0x1d, 0x54, 0xde, 0x45, 0x27,
	0x1e, 0x69, 0x13, 0xb4, 0x53, 0xc6, 0x05, 0xf5, 0xc6, 0xab, 0x58, 0x5b, 0x8b, 0x85, 0x2a, 0x0e,
	0x63, 0x6d, 0x95, 0xab, 0xd0, 0x1f, 0xa4, 0x96, 0x03, 0xeb, 0x72, 0x0c, 0xf4, 0xf7, 0x80, 0x89,
	0xd0, 0x3c, 0xf7, 0xfe, 0x86, 0xed, 0xae, 0xa9, 0x3a, 0x39, 0xfa, 0x67, 0x2a, 0xee, 0xc0, 0x5a,
	0x62, 0x95, 0xcd, 0x4a, 0x94, 0x2b, 0xbd, 0x95, 0xfe, 0xda, 0x64, 0x35, 0x09, 0xa7, 0x59, 0x89,
	0xe2, 0x4b, 0xd8, 0x2c, 0x9c, 0xce, 0x8a, 0x76, 0x96, 0xc9, 0xe5, 0x07, 0xbd, 0x95, 0xfe, 0xc6,
	0x64, 0x83, 0xe4, 0x57, 0xa4, 0x1e, 0xe7, 0x7b, 0xff, 0xdd, 0x82, 0xf5, 0xae, 0xb5, 0xf8, 0x02,
	0x98, 0x50, 0x59, 0x9e, 0x7b, 0x0c, 0x41, 0x1e, 0x92, 0xf3, 0x3a, 0x89, 0x03, 0xd6, 0xc4, 0x3d,
	0xf8, 0xc4, 0x63, 0xe9, 0x22, 0xce, 0xa9, 0x07, 0x44, 0x6d, 0xb0, 0xda, 0x62, 0xbb, 0x00, 0xec,
	0x55, 0x39, 0x1f, 0xe5, 0x37, 0x34, 0x7f, 0x8d, 0x94, 0xb1, 0xf3, 0x51, 0xdc, 0x85, 0x5b, 0x8d,
	0x0b, 0xd5, 0xbf, 0xa5, 0x3a, 0xb0, 0x44, 0xc0, 0x0e, 0xac, 0x52, 0x7c, 0xda, 0x15, 0xf2, 0x3b,
	0xaa, 0xce, 0xdf, 0xd3, 0x82, 0x26, 0xa8, 0xaa, 0x8c, 0xb5, 0x42, 0x9b, 0x9d, 0x15, 0x98, 0xcb,
	0xef, 0x7b, 0x2b, 0xfd, 0xd5, 0xc9, 0x86, 0x09, 0xe3, 0x32, 0xd6, 0x47, 0x2c, 0x8a, 0x3e, 0x6c,
	0x35, 0x43, 0x16, 0x49, 0xfc, 0x40, 0x5e, 0xcd, 0x0a, 0x6d, 0x14, 0x62, 0x1f, 0xb6, 0x97, 0x22,
	0xa3, 0x5c, 0x1f, 0xd2, 0x5e, 0x9b, 0x9d, 0xd0, 0x28, 0xde, 0xfb, 0x20, 0x96, 0x5d, 0x09, 0xfe,
	0x91, 0xe0, 0xad, 0xae, 0x2f, 0xd1, 0xbb, 0x00, 0xba, 0xc8, 0x42, 0x60, 0xea, 0x27, 0xa2, 0xd6,
	0x48, 0xa1, 0xf2, 0x57, 0xb0, 0x99, 0xe9, 0x68, 0xfe, 0x45, 0x15, 0x30, 0x04, 0xe3, 0x6c, 0x90,
	0x3f, 0xf3, 0x09, 0x59, 0x9e, 0x36, 0x6a, 0x02, 0x3d, 0x46, 0x9f, 0xd9, 0x50, 0x9a, 0xa8, 0xa2,
	0x29, 0x51, 0x3e, 0xea, 0x7d, 0xc8, 0xab, 0xb4, 0xf2, 0x2b, 0x53, 0x62, 0x4a, 0x36, 0xe0, 0xeb,
	0x1a, 0xad, 0x46, 0x65, 0x83, 0x7c, 0xcc, 0xc9, 0xb6, 0xd2, 0x69, 0x58, 0x06, 0xbc, 0xfc, 0xe5,
	0x1d, 0xc0, 0x2f, 0xc2, 0xb8, 0x32, 0x36, 0x77, 0x57, 0x2a, 0x98, 0x6b, 0x94, 0xbf, 0x12, 0xc6,
	0x61, 0xfc, 0x45, 0xfa, 0xd4, 0x5c, 0x77, 0xc3, 0xe8, 0xc2, 0x03, 0x82, 0x9b, 0x30, 0x3a, 0xf4,
	0xd7, 0xf0, 0xe9, 0xfc, 0xb4, 0xb4, 0x17, 0x2f, 0xf2, 0x84, 0x70, 0xb1, 0x5c, 0xa2, 0x65, 0x1e,
	0xc3, 0x9d, 0x32, 0x7b, 0x63, 0xca, 0xba, 0x54, 0x37, 0x35, 0x0e, 0xa9, 0xf1, 0x76, 0x83, 0x4c,
	0xde, 0xef, 0xdf, 0x87, 0xed, 0xda, 0x06, 0xb4, 0x51, 0xbd, 0xae, 0xb1, 0x46, 0x3e, 0xdd, 0x53,
	0x5e, 0x85, 0x0b, 0x2f, 0x93, 0x4e, 0x87, 0x7b, 0x08, 0xb7, 0x1b, 0xb6, 0x1d, 0xd9, 0xe9, 0x39,
	0xa2, 0x9e, 0xcf, 0x18, 0x78, 0xc1, 0xf5, 0x45, 0xeb, 0x3e, 0x6c, 0x7b, 0x0c, 0x68, 0xf3, 0x6e,
	0xcb, 0x33, 0x1e, 0xc3, 0x85, 0xa5, 0x31, 0x0d, 0x7b, 0xc3, 0x98, 0xdf, 0x78, 0x0c, 0x03, 0xef,
	0x8d, 0xe9, 0xc3, 0x96, 0xf3, 0x39, 0xfa, 0x6e, 0xc7, 0xef, 0xfc, 0xb5, 0x90, 0xbe, 0x20, 0xef,
	0x83, 0xa8, 0x32, 0x7d, 0x89, 0xed, 0xde, 0xfa, 0x02, 0xf5, 0xa5, 0x3c, 0xe6, 0x7f, 0x0b, 0x57,
	0x08, 0x1e, 0x26, 0x3d, 0xfd, 0xa4, 0x73, 0x73, 0x8e, 0x21, 0xaa, 0x80, 0xda, 0x63, 0x0c, 0xf2,
	0x0f, 0xbe, 0x2f, 0x58, 0x9d, 0xb2, 0x28, 0x24, 0x7c, 0xcc, 0x07, 0x0b, 0xf2, 0x84, 0xea, 0xed,
	0xab, 0x18, 0xc3, 0xbd, 0x6b, 0xf4, 0x4e, 0x15, 0x68, 0xcf, 0xe3, 0x85, 0x3a, 0x73, 0xf9, 0x5b,
	0x95, 0xe9, 0x4b, 0xeb, 0xae, 0x0a, 0xcc, 0xcf, 0xb1, 0x4c, 0xa1, 0xa6, 0xe0, 0xe4, 0x73, 0xea,
	0xfb, 0x3c, 0xc1, 0xcf, 0x89, 0x7d, 0xe2, 0xf2, 0xb7, 0x83, 0x65, 0x72, 0x8a, 0x36, 0x8a, 0x21,
	0xdc, 0x8d, 0x2e, 0x66, 0x85, 0x72, 0x75, 0x54, 0x6e, 0xa6, 0x78, 0xef, 0xdc, 0xbb, 0x4a, 0xf1,
	0xf1, 0x83, 0x7c, 0x41, 0x5e, 0x3b, 0x84, 0x8d, 0xea, 0x38, 0x9a, 0x8d, 0x12, 0xf3, 0xd4, 0xbb,
	0x6a, 0xcc, 0x84, 0x38, 0x86, 0xbd, 0x1b, 0x4c, 0x3c, 0xf2, 0xb3, 0xf5, 0x39, 0x25, 0x9f, 0xdd,
	0x77, 0x7c, 0x26, 0x4c, 0xb5, 0x56, 0x27, 0xad, 0x55, 0x85, 0xe8, 0x55, 0x56, 0xc7, 0x0b, 0xb4,
	0xd1, 0xe8, 0x2c, 0xa6, 0x6f, 0x71, 0x96, 0x99, 0xa2, 0xf6, 0x18, 0xe4, 0x88, 0xac, 0xf8, 0xe4,
	0x63, 0x44, 0x3f, 0x58, 0xe2, 0x9e, 0x35, 0x98, 0xe8, 0xc1, 0xba, 0x09, 0xed, 0xed, 0x51, 0x57,
	0x72, 0x4c, 0x97, 0x17, 0x98, 0xc0, 0xf7, 0xc6, 0x9f, 0x95, 0x18, 0xc0, 0xae, 0x09, 0x4a, 0x3b,
	0x9b, 0xf2, 0x4f, 0x23, 0xb4, 0xb3, 0xd1, 0xbb, 0x62, 0x7e, 0xdf, 0xbd, 0xa4, 0x96, 0x1d, 0x13,
	0x86, 0x73, 0x66, 0xc8, 0x48, 0x73, 0xf9, 0x9d, 0x7d, 0x44, 0xd7, 0xe5, 0x83, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x28, 0xd0, 0x0a, 0x38, 0x93, 0x06, 0x00, 0x00,
}
