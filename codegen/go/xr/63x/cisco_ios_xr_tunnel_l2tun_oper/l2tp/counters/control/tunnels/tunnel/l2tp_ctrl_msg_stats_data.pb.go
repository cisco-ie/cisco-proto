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
// source: l2tp_ctrl_msg_stats_data.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tp_counters_control_tunnels_tunnel

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

type L2TpCtrlMsgStatsData_KEYS struct {
	TunnelId             uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpCtrlMsgStatsData_KEYS) Reset()         { *m = L2TpCtrlMsgStatsData_KEYS{} }
func (m *L2TpCtrlMsgStatsData_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2TpCtrlMsgStatsData_KEYS) ProtoMessage()    {}
func (*L2TpCtrlMsgStatsData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e922b457f47aeae, []int{0}
}

func (m *L2TpCtrlMsgStatsData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS.Unmarshal(m, b)
}
func (m *L2TpCtrlMsgStatsData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS.Marshal(b, m, deterministic)
}
func (m *L2TpCtrlMsgStatsData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS.Merge(m, src)
}
func (m *L2TpCtrlMsgStatsData_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS.Size(m)
}
func (m *L2TpCtrlMsgStatsData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCtrlMsgStatsData_KEYS proto.InternalMessageInfo

func (m *L2TpCtrlMsgStatsData_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

type L2TpCtlMsgBriefData struct {
	RemoteTunnelId       uint32   `protobuf:"varint,1,opt,name=remote_tunnel_id,json=remoteTunnelId,proto3" json:"remote_tunnel_id,omitempty"`
	LocalAddress         string   `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        string   `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpCtlMsgBriefData) Reset()         { *m = L2TpCtlMsgBriefData{} }
func (m *L2TpCtlMsgBriefData) String() string { return proto.CompactTextString(m) }
func (*L2TpCtlMsgBriefData) ProtoMessage()    {}
func (*L2TpCtlMsgBriefData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e922b457f47aeae, []int{1}
}

func (m *L2TpCtlMsgBriefData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCtlMsgBriefData.Unmarshal(m, b)
}
func (m *L2TpCtlMsgBriefData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCtlMsgBriefData.Marshal(b, m, deterministic)
}
func (m *L2TpCtlMsgBriefData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCtlMsgBriefData.Merge(m, src)
}
func (m *L2TpCtlMsgBriefData) XXX_Size() int {
	return xxx_messageInfo_L2TpCtlMsgBriefData.Size(m)
}
func (m *L2TpCtlMsgBriefData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCtlMsgBriefData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCtlMsgBriefData proto.InternalMessageInfo

func (m *L2TpCtlMsgBriefData) GetRemoteTunnelId() uint32 {
	if m != nil {
		return m.RemoteTunnelId
	}
	return 0
}

func (m *L2TpCtlMsgBriefData) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *L2TpCtlMsgBriefData) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

type L2TpCtrlMsgCounts struct {
	UnknownPackets                      uint32   `protobuf:"varint,1,opt,name=unknown_packets,json=unknownPackets,proto3" json:"unknown_packets,omitempty"`
	ZeroLengthBodyPackets               uint32   `protobuf:"varint,2,opt,name=zero_length_body_packets,json=zeroLengthBodyPackets,proto3" json:"zero_length_body_packets,omitempty"`
	StartControlConnectionRequests      uint32   `protobuf:"varint,3,opt,name=start_control_connection_requests,json=startControlConnectionRequests,proto3" json:"start_control_connection_requests,omitempty"`
	StartControlConnectionReplies       uint32   `protobuf:"varint,4,opt,name=start_control_connection_replies,json=startControlConnectionReplies,proto3" json:"start_control_connection_replies,omitempty"`
	StartControlConnectionNotifications uint32   `protobuf:"varint,5,opt,name=start_control_connection_notifications,json=startControlConnectionNotifications,proto3" json:"start_control_connection_notifications,omitempty"`
	StopControlConnectionNotifications  uint32   `protobuf:"varint,6,opt,name=stop_control_connection_notifications,json=stopControlConnectionNotifications,proto3" json:"stop_control_connection_notifications,omitempty"`
	HelloPackets                        uint32   `protobuf:"varint,7,opt,name=hello_packets,json=helloPackets,proto3" json:"hello_packets,omitempty"`
	OutgoingCallRequests                uint32   `protobuf:"varint,8,opt,name=outgoing_call_requests,json=outgoingCallRequests,proto3" json:"outgoing_call_requests,omitempty"`
	OutgoingCallReplies                 uint32   `protobuf:"varint,9,opt,name=outgoing_call_replies,json=outgoingCallReplies,proto3" json:"outgoing_call_replies,omitempty"`
	OutgoingCallConnectedPackets        uint32   `protobuf:"varint,10,opt,name=outgoing_call_connected_packets,json=outgoingCallConnectedPackets,proto3" json:"outgoing_call_connected_packets,omitempty"`
	IncomingCallRequests                uint32   `protobuf:"varint,11,opt,name=incoming_call_requests,json=incomingCallRequests,proto3" json:"incoming_call_requests,omitempty"`
	IncomingCallReplies                 uint32   `protobuf:"varint,12,opt,name=incoming_call_replies,json=incomingCallReplies,proto3" json:"incoming_call_replies,omitempty"`
	IncomingCallConnectedPackets        uint32   `protobuf:"varint,13,opt,name=incoming_call_connected_packets,json=incomingCallConnectedPackets,proto3" json:"incoming_call_connected_packets,omitempty"`
	CallDisconnectNotifyPackets         uint32   `protobuf:"varint,14,opt,name=call_disconnect_notify_packets,json=callDisconnectNotifyPackets,proto3" json:"call_disconnect_notify_packets,omitempty"`
	WanErrorNotifyPackets               uint32   `protobuf:"varint,15,opt,name=wan_error_notify_packets,json=wanErrorNotifyPackets,proto3" json:"wan_error_notify_packets,omitempty"`
	SetLinkInfoPackets                  uint32   `protobuf:"varint,16,opt,name=set_link_info_packets,json=setLinkInfoPackets,proto3" json:"set_link_info_packets,omitempty"`
	ServiceRelayRequests                uint32   `protobuf:"varint,17,opt,name=service_relay_requests,json=serviceRelayRequests,proto3" json:"service_relay_requests,omitempty"`
	ServiceRelayReplies                 uint32   `protobuf:"varint,18,opt,name=service_relay_replies,json=serviceRelayReplies,proto3" json:"service_relay_replies,omitempty"`
	AcknowledgementPackets              uint32   `protobuf:"varint,19,opt,name=acknowledgement_packets,json=acknowledgementPackets,proto3" json:"acknowledgement_packets,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *L2TpCtrlMsgCounts) Reset()         { *m = L2TpCtrlMsgCounts{} }
func (m *L2TpCtrlMsgCounts) String() string { return proto.CompactTextString(m) }
func (*L2TpCtrlMsgCounts) ProtoMessage()    {}
func (*L2TpCtrlMsgCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e922b457f47aeae, []int{2}
}

func (m *L2TpCtrlMsgCounts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCtrlMsgCounts.Unmarshal(m, b)
}
func (m *L2TpCtrlMsgCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCtrlMsgCounts.Marshal(b, m, deterministic)
}
func (m *L2TpCtrlMsgCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCtrlMsgCounts.Merge(m, src)
}
func (m *L2TpCtrlMsgCounts) XXX_Size() int {
	return xxx_messageInfo_L2TpCtrlMsgCounts.Size(m)
}
func (m *L2TpCtrlMsgCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCtrlMsgCounts.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCtrlMsgCounts proto.InternalMessageInfo

func (m *L2TpCtrlMsgCounts) GetUnknownPackets() uint32 {
	if m != nil {
		return m.UnknownPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetZeroLengthBodyPackets() uint32 {
	if m != nil {
		return m.ZeroLengthBodyPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetStartControlConnectionRequests() uint32 {
	if m != nil {
		return m.StartControlConnectionRequests
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetStartControlConnectionReplies() uint32 {
	if m != nil {
		return m.StartControlConnectionReplies
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetStartControlConnectionNotifications() uint32 {
	if m != nil {
		return m.StartControlConnectionNotifications
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetStopControlConnectionNotifications() uint32 {
	if m != nil {
		return m.StopControlConnectionNotifications
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetHelloPackets() uint32 {
	if m != nil {
		return m.HelloPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetOutgoingCallRequests() uint32 {
	if m != nil {
		return m.OutgoingCallRequests
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetOutgoingCallReplies() uint32 {
	if m != nil {
		return m.OutgoingCallReplies
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetOutgoingCallConnectedPackets() uint32 {
	if m != nil {
		return m.OutgoingCallConnectedPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetIncomingCallRequests() uint32 {
	if m != nil {
		return m.IncomingCallRequests
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetIncomingCallReplies() uint32 {
	if m != nil {
		return m.IncomingCallReplies
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetIncomingCallConnectedPackets() uint32 {
	if m != nil {
		return m.IncomingCallConnectedPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetCallDisconnectNotifyPackets() uint32 {
	if m != nil {
		return m.CallDisconnectNotifyPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetWanErrorNotifyPackets() uint32 {
	if m != nil {
		return m.WanErrorNotifyPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetSetLinkInfoPackets() uint32 {
	if m != nil {
		return m.SetLinkInfoPackets
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetServiceRelayRequests() uint32 {
	if m != nil {
		return m.ServiceRelayRequests
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetServiceRelayReplies() uint32 {
	if m != nil {
		return m.ServiceRelayReplies
	}
	return 0
}

func (m *L2TpCtrlMsgCounts) GetAcknowledgementPackets() uint32 {
	if m != nil {
		return m.AcknowledgementPackets
	}
	return 0
}

type L2TpCtrlMsgStatsGlobalData struct {
	TotalTransmit        uint32             `protobuf:"varint,1,opt,name=total_transmit,json=totalTransmit,proto3" json:"total_transmit,omitempty"`
	TotalRetransmit      uint32             `protobuf:"varint,2,opt,name=total_retransmit,json=totalRetransmit,proto3" json:"total_retransmit,omitempty"`
	TotalReceived        uint32             `protobuf:"varint,3,opt,name=total_received,json=totalReceived,proto3" json:"total_received,omitempty"`
	TotalDrop            uint32             `protobuf:"varint,4,opt,name=total_drop,json=totalDrop,proto3" json:"total_drop,omitempty"`
	Transmit             *L2TpCtrlMsgCounts `protobuf:"bytes,5,opt,name=transmit,proto3" json:"transmit,omitempty"`
	Retransmit           *L2TpCtrlMsgCounts `protobuf:"bytes,6,opt,name=retransmit,proto3" json:"retransmit,omitempty"`
	Received             *L2TpCtrlMsgCounts `protobuf:"bytes,7,opt,name=received,proto3" json:"received,omitempty"`
	Drop                 *L2TpCtrlMsgCounts `protobuf:"bytes,8,opt,name=drop,proto3" json:"drop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2TpCtrlMsgStatsGlobalData) Reset()         { *m = L2TpCtrlMsgStatsGlobalData{} }
func (m *L2TpCtrlMsgStatsGlobalData) String() string { return proto.CompactTextString(m) }
func (*L2TpCtrlMsgStatsGlobalData) ProtoMessage()    {}
func (*L2TpCtrlMsgStatsGlobalData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e922b457f47aeae, []int{3}
}

func (m *L2TpCtrlMsgStatsGlobalData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCtrlMsgStatsGlobalData.Unmarshal(m, b)
}
func (m *L2TpCtrlMsgStatsGlobalData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCtrlMsgStatsGlobalData.Marshal(b, m, deterministic)
}
func (m *L2TpCtrlMsgStatsGlobalData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCtrlMsgStatsGlobalData.Merge(m, src)
}
func (m *L2TpCtrlMsgStatsGlobalData) XXX_Size() int {
	return xxx_messageInfo_L2TpCtrlMsgStatsGlobalData.Size(m)
}
func (m *L2TpCtrlMsgStatsGlobalData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCtrlMsgStatsGlobalData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCtrlMsgStatsGlobalData proto.InternalMessageInfo

func (m *L2TpCtrlMsgStatsGlobalData) GetTotalTransmit() uint32 {
	if m != nil {
		return m.TotalTransmit
	}
	return 0
}

func (m *L2TpCtrlMsgStatsGlobalData) GetTotalRetransmit() uint32 {
	if m != nil {
		return m.TotalRetransmit
	}
	return 0
}

func (m *L2TpCtrlMsgStatsGlobalData) GetTotalReceived() uint32 {
	if m != nil {
		return m.TotalReceived
	}
	return 0
}

func (m *L2TpCtrlMsgStatsGlobalData) GetTotalDrop() uint32 {
	if m != nil {
		return m.TotalDrop
	}
	return 0
}

func (m *L2TpCtrlMsgStatsGlobalData) GetTransmit() *L2TpCtrlMsgCounts {
	if m != nil {
		return m.Transmit
	}
	return nil
}

func (m *L2TpCtrlMsgStatsGlobalData) GetRetransmit() *L2TpCtrlMsgCounts {
	if m != nil {
		return m.Retransmit
	}
	return nil
}

func (m *L2TpCtrlMsgStatsGlobalData) GetReceived() *L2TpCtrlMsgCounts {
	if m != nil {
		return m.Received
	}
	return nil
}

func (m *L2TpCtrlMsgStatsGlobalData) GetDrop() *L2TpCtrlMsgCounts {
	if m != nil {
		return m.Drop
	}
	return nil
}

type L2TpCtrlMsgStatsData struct {
	Brief                *L2TpCtlMsgBriefData        `protobuf:"bytes,50,opt,name=brief,proto3" json:"brief,omitempty"`
	Global               *L2TpCtrlMsgStatsGlobalData `protobuf:"bytes,51,opt,name=global,proto3" json:"global,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *L2TpCtrlMsgStatsData) Reset()         { *m = L2TpCtrlMsgStatsData{} }
func (m *L2TpCtrlMsgStatsData) String() string { return proto.CompactTextString(m) }
func (*L2TpCtrlMsgStatsData) ProtoMessage()    {}
func (*L2TpCtrlMsgStatsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e922b457f47aeae, []int{4}
}

func (m *L2TpCtrlMsgStatsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpCtrlMsgStatsData.Unmarshal(m, b)
}
func (m *L2TpCtrlMsgStatsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpCtrlMsgStatsData.Marshal(b, m, deterministic)
}
func (m *L2TpCtrlMsgStatsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpCtrlMsgStatsData.Merge(m, src)
}
func (m *L2TpCtrlMsgStatsData) XXX_Size() int {
	return xxx_messageInfo_L2TpCtrlMsgStatsData.Size(m)
}
func (m *L2TpCtrlMsgStatsData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpCtrlMsgStatsData.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpCtrlMsgStatsData proto.InternalMessageInfo

func (m *L2TpCtrlMsgStatsData) GetBrief() *L2TpCtlMsgBriefData {
	if m != nil {
		return m.Brief
	}
	return nil
}

func (m *L2TpCtrlMsgStatsData) GetGlobal() *L2TpCtrlMsgStatsGlobalData {
	if m != nil {
		return m.Global
	}
	return nil
}

func init() {
	proto.RegisterType((*L2TpCtrlMsgStatsData_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.counters.control.tunnels.tunnel.l2tp_ctrl_msg_stats_data_KEYS")
	proto.RegisterType((*L2TpCtlMsgBriefData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.counters.control.tunnels.tunnel.l2tp_ctl_msg_brief_data")
	proto.RegisterType((*L2TpCtrlMsgCounts)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.counters.control.tunnels.tunnel.l2tp_ctrl_msg_counts")
	proto.RegisterType((*L2TpCtrlMsgStatsGlobalData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.counters.control.tunnels.tunnel.l2tp_ctrl_msg_stats_global_data")
	proto.RegisterType((*L2TpCtrlMsgStatsData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tp.counters.control.tunnels.tunnel.l2tp_ctrl_msg_stats_data")
}

func init() { proto.RegisterFile("l2tp_ctrl_msg_stats_data.proto", fileDescriptor_2e922b457f47aeae) }

var fileDescriptor_2e922b457f47aeae = []byte{
	// 797 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x4f, 0x2b, 0x37,
	0x14, 0x55, 0x28, 0x84, 0xe4, 0xc2, 0x00, 0x35, 0x5f, 0x23, 0x51, 0x3e, 0x1a, 0x44, 0x4b, 0x37,
	0x91, 0x1a, 0x2a, 0xb1, 0xe9, 0xa6, 0x0d, 0xa8, 0x42, 0x45, 0x55, 0x3b, 0xb0, 0x41, 0xaa, 0x64,
	0x39, 0x33, 0x4e, 0xb0, 0xe2, 0xd8, 0x53, 0xdb, 0x81, 0xe6, 0xe9, 0xfd, 0x84, 0x27, 0xa1, 0xf7,
	0x8f, 0x9f, 0xc6, 0x76, 0x26, 0x93, 0x09, 0x79, 0x6c, 0xc8, 0x2a, 0xca, 0xb9, 0xe7, 0xdc, 0x7b,
	0x7c, 0x26, 0xbe, 0x13, 0x38, 0xe2, 0x2d, 0x93, 0xe2, 0xd8, 0x28, 0x8e, 0x07, 0xba, 0x87, 0xb5,
	0x21, 0x46, 0xe3, 0x84, 0x18, 0xd2, 0x4c, 0x95, 0x34, 0x12, 0xb5, 0x63, 0xa6, 0x63, 0x89, 0x99,
	0xd4, 0xf8, 0x7f, 0x85, 0xcd, 0x50, 0x08, 0xca, 0x31, 0x6f, 0x99, 0xa1, 0xc0, 0x32, 0xa5, 0xaa,
	0x99, 0xc9, 0x9b, 0xb1, 0x1c, 0x0a, 0x43, 0x95, 0x6e, 0xc6, 0x52, 0x18, 0x25, 0x79, 0xd3, 0xf1,
	0xb4, 0xff, 0x6c, 0xfc, 0x0a, 0x87, 0xf3, 0xc6, 0xe0, 0x3f, 0xaf, 0x1f, 0xee, 0xd0, 0x01, 0xd4,
	0x7d, 0x6b, 0x96, 0x84, 0x95, 0x93, 0xca, 0x79, 0x10, 0xd5, 0x1c, 0x70, 0x93, 0x34, 0x5e, 0x2a,
	0xb0, 0xef, 0xe5, 0x4e, 0xdd, 0x51, 0x8c, 0x76, 0xad, 0x1a, 0x9d, 0xc3, 0x96, 0xa2, 0x03, 0x69,
	0x28, 0x2e, 0xeb, 0x37, 0x1c, 0x7e, 0xef, 0xbb, 0xa0, 0x53, 0x08, 0xb8, 0x8c, 0x09, 0xc7, 0x24,
	0x49, 0x14, 0xd5, 0x3a, 0x5c, 0x3a, 0xa9, 0x9c, 0xd7, 0xa3, 0x75, 0x0b, 0xfe, 0xe6, 0x30, 0x74,
	0x06, 0x5e, 0x96, 0xb3, 0xbe, 0xb1, 0xac, 0xc0, 0xa1, 0x9e, 0xd6, 0x78, 0xa9, 0xc3, 0xce, 0xf4,
	0x81, 0x6c, 0x02, 0x1a, 0xfd, 0x08, 0x9b, 0x43, 0xd1, 0x17, 0xf2, 0x59, 0xe0, 0x94, 0xc4, 0x7d,
	0x6a, 0xf4, 0xd8, 0x8d, 0x87, 0xff, 0x76, 0x28, 0xba, 0x84, 0xf0, 0x03, 0x55, 0x12, 0x73, 0x2a,
	0x7a, 0xe6, 0x11, 0x77, 0x64, 0x32, 0xca, 0x15, 0x4b, 0x56, 0xb1, 0x9b, 0xd5, 0x6f, 0x6d, 0xf9,
	0x77, 0x99, 0x8c, 0xc6, 0xc2, 0x1b, 0xf8, 0x5e, 0x1b, 0xa2, 0x0c, 0xf6, 0x51, 0x67, 0x9f, 0x82,
	0xc6, 0x86, 0x49, 0x81, 0x15, 0xfd, 0x6f, 0x48, 0xb5, 0x71, 0xa6, 0x83, 0xe8, 0xc8, 0x12, 0xdb,
	0x8e, 0xd7, 0xce, 0x69, 0x91, 0x67, 0xa1, 0x3f, 0xe0, 0xe4, 0x2b, 0xad, 0x52, 0xce, 0xa8, 0x0e,
	0x97, 0x6d, 0xa7, 0xc3, 0x79, 0x9d, 0x2c, 0x09, 0xdd, 0xc1, 0x0f, 0x73, 0x1b, 0x09, 0x69, 0x58,
	0x97, 0xc5, 0x24, 0xfb, 0xa2, 0xc3, 0x15, 0xdb, 0xee, 0xf4, 0xf5, 0x76, 0x7f, 0x15, 0xa9, 0xe8,
	0x1f, 0x38, 0xd3, 0x46, 0xa6, 0x6f, 0xf7, 0xac, 0xda, 0x9e, 0x8d, 0x8c, 0xfc, 0x46, 0xcb, 0x53,
	0x08, 0x1e, 0x29, 0xe7, 0x32, 0x4f, 0x7a, 0xd5, 0x4a, 0xd7, 0x2d, 0x38, 0x0e, 0xf8, 0x17, 0xd8,
	0x93, 0x43, 0xd3, 0x93, 0x4c, 0xf4, 0x70, 0x4c, 0x38, 0x9f, 0xa4, 0x5a, 0xb3, 0xec, 0x9d, 0x71,
	0xb5, 0x4d, 0x38, 0xcf, 0xb3, 0x6c, 0xc1, 0x6e, 0x59, 0xe5, 0x02, 0xac, 0x5b, 0xd1, 0xf6, 0xb4,
	0xc8, 0xc5, 0x76, 0x0d, 0xc7, 0xd3, 0x1a, 0x7f, 0x44, 0x9a, 0xe4, 0x06, 0xc1, 0xaa, 0xbf, 0x2b,
	0xaa, 0xdb, 0x63, 0x52, 0xc1, 0x30, 0x13, 0xb1, 0x1c, 0xcc, 0x1a, 0x5e, 0x73, 0x86, 0xc7, 0xd5,
	0xb2, 0xe1, 0xb2, 0xca, 0x19, 0x5e, 0x77, 0x86, 0xa7, 0x45, 0xb9, 0xe1, 0x69, 0xcd, 0xac, 0xe1,
	0xc0, 0x19, 0x2e, 0xaa, 0x67, 0x0c, 0xb7, 0xe1, 0xc8, 0xaa, 0x93, 0x6c, 0xb3, 0xd8, 0x9a, 0x7b,
	0xa2, 0x93, 0x1b, 0xb0, 0x61, 0xbb, 0x1c, 0x64, 0xac, 0xab, 0x9c, 0x64, 0x9f, 0xe5, 0xa8, 0x70,
	0x81, 0x9e, 0x89, 0xc0, 0x54, 0x29, 0xa9, 0xca, 0xf2, 0x4d, 0x77, 0x81, 0x9e, 0x89, 0xb8, 0xce,
	0xca, 0xd3, 0xc2, 0x9f, 0x61, 0x57, 0x53, 0x83, 0x39, 0x13, 0x7d, 0xcc, 0x44, 0x77, 0xf2, 0x63,
	0xd8, 0xb2, 0x2a, 0xa4, 0xa9, 0xb9, 0x65, 0xa2, 0x7f, 0x23, 0xba, 0xc5, 0x9f, 0x84, 0xa6, 0xea,
	0x89, 0xc5, 0x14, 0x2b, 0xca, 0xc9, 0x68, 0x92, 0xf0, 0xb7, 0x2e, 0x61, 0x5f, 0x8d, 0xb2, 0x62,
	0x31, 0xe1, 0xb2, 0xca, 0x25, 0x8c, 0x5c, 0xc2, 0xd3, 0x22, 0x97, 0xf0, 0x25, 0xec, 0x93, 0x38,
	0x5b, 0x14, 0x9c, 0x26, 0x3d, 0x3a, 0xa0, 0xc2, 0xe4, 0xf6, 0xb6, 0xad, 0x6a, 0xaf, 0x54, 0xf6,
	0x16, 0x1b, 0x9f, 0x57, 0xe0, 0xf8, 0xb5, 0x15, 0xdb, 0xe3, 0xb2, 0x43, 0xb8, 0xdb, 0x95, 0x67,
	0xb0, 0x61, 0xa4, 0x21, 0x1c, 0x1b, 0x45, 0x84, 0x1e, 0x30, 0xe3, 0x77, 0x53, 0x60, 0xd1, 0x7b,
	0x0f, 0xa2, 0x9f, 0x60, 0xcb, 0xd1, 0x14, 0xcd, 0x89, 0x6e, 0x25, 0x6d, 0x5a, 0x3c, 0xca, 0xe1,
	0x49, 0x47, 0x45, 0x63, 0xca, 0x9e, 0x68, 0xe2, 0x37, 0x4f, 0xe0, 0x89, 0x0e, 0x44, 0x87, 0x00,
	0x8e, 0x96, 0x28, 0x99, 0xfa, 0x95, 0x52, 0xb7, 0xc8, 0x95, 0x92, 0x29, 0x1a, 0x42, 0x2d, 0x1f,
	0x94, 0x2d, 0x88, 0xb5, 0xd6, 0x43, 0xf3, 0x1d, 0xde, 0x3a, 0xcd, 0xd7, 0x36, 0x74, 0x94, 0x8f,
	0x42, 0x23, 0x80, 0xc2, 0x09, 0xab, 0x8b, 0x1e, 0x5c, 0x18, 0x96, 0x9d, 0x38, 0x4f, 0x6c, 0x75,
	0xe1, 0x27, 0x1e, 0x8f, 0x42, 0x03, 0x58, 0xb6, 0x4f, 0xa0, 0xb6, 0xe8, 0x91, 0x76, 0x4c, 0xe3,
	0xd3, 0x12, 0x84, 0xf3, 0x5e, 0xfb, 0x48, 0xc1, 0x8a, 0x7d, 0x8d, 0x87, 0x2d, 0x6b, 0xe6, 0xdf,
	0xf7, 0x34, 0x53, 0xfe, 0x97, 0x10, 0xb9, 0x51, 0xe8, 0x23, 0x54, 0xdd, 0x7d, 0x08, 0x2f, 0xec,
	0xd0, 0x64, 0x01, 0x09, 0xcc, 0x5c, 0xbb, 0xc8, 0xcf, 0xec, 0x54, 0xed, 0x1f, 0xaa, 0x8b, 0x2f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x49, 0xf5, 0xfc, 0x72, 0x09, 0x00, 0x00,
}
