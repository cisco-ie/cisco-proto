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
// source: ip_traffic.proto

package cisco_ios_xr_ipv4_io_oper_ipv4_network_nodes_node_statistics_traffic

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

type IpTraffic_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpTraffic_KEYS) Reset()         { *m = IpTraffic_KEYS{} }
func (m *IpTraffic_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpTraffic_KEYS) ProtoMessage()    {}
func (*IpTraffic_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a2a8e22f29b44, []int{0}
}

func (m *IpTraffic_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpTraffic_KEYS.Unmarshal(m, b)
}
func (m *IpTraffic_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpTraffic_KEYS.Marshal(b, m, deterministic)
}
func (m *IpTraffic_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpTraffic_KEYS.Merge(m, src)
}
func (m *IpTraffic_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpTraffic_KEYS.Size(m)
}
func (m *IpTraffic_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpTraffic_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpTraffic_KEYS proto.InternalMessageInfo

func (m *IpTraffic_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv4IoTraffic struct {
	InputPackets            uint32   `protobuf:"varint,1,opt,name=input_packets,json=inputPackets,proto3" json:"input_packets,omitempty"`
	ReceivedPackets         uint32   `protobuf:"varint,2,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	FormatErrors            uint32   `protobuf:"varint,3,opt,name=format_errors,json=formatErrors,proto3" json:"format_errors,omitempty"`
	BadHopCount             uint32   `protobuf:"varint,4,opt,name=bad_hop_count,json=badHopCount,proto3" json:"bad_hop_count,omitempty"`
	BadSourceAddress        uint32   `protobuf:"varint,5,opt,name=bad_source_address,json=badSourceAddress,proto3" json:"bad_source_address,omitempty"`
	BadHeader               uint32   `protobuf:"varint,6,opt,name=bad_header,json=badHeader,proto3" json:"bad_header,omitempty"`
	NoProtocol              uint32   `protobuf:"varint,7,opt,name=no_protocol,json=noProtocol,proto3" json:"no_protocol,omitempty"`
	NoGateway               uint32   `protobuf:"varint,8,opt,name=no_gateway,json=noGateway,proto3" json:"no_gateway,omitempty"`
	ReassembleInput         uint32   `protobuf:"varint,9,opt,name=reassemble_input,json=reassembleInput,proto3" json:"reassemble_input,omitempty"`
	Reassembled             uint32   `protobuf:"varint,10,opt,name=reassembled,proto3" json:"reassembled,omitempty"`
	ReassembleTimeout       uint32   `protobuf:"varint,11,opt,name=reassemble_timeout,json=reassembleTimeout,proto3" json:"reassemble_timeout,omitempty"`
	ReassembleMaxDrop       uint32   `protobuf:"varint,12,opt,name=reassemble_max_drop,json=reassembleMaxDrop,proto3" json:"reassemble_max_drop,omitempty"`
	ReassembleFailed        uint32   `protobuf:"varint,13,opt,name=reassemble_failed,json=reassembleFailed,proto3" json:"reassemble_failed,omitempty"`
	OptionsPresent          uint32   `protobuf:"varint,14,opt,name=options_present,json=optionsPresent,proto3" json:"options_present,omitempty"`
	BadOption               uint32   `protobuf:"varint,15,opt,name=bad_option,json=badOption,proto3" json:"bad_option,omitempty"`
	UnknownOption           uint32   `protobuf:"varint,16,opt,name=unknown_option,json=unknownOption,proto3" json:"unknown_option,omitempty"`
	BadSecurityOption       uint32   `protobuf:"varint,17,opt,name=bad_security_option,json=badSecurityOption,proto3" json:"bad_security_option,omitempty"`
	BasicSecurityOption     uint32   `protobuf:"varint,18,opt,name=basic_security_option,json=basicSecurityOption,proto3" json:"basic_security_option,omitempty"`
	ExtendedSecurityOption  uint32   `protobuf:"varint,19,opt,name=extended_security_option,json=extendedSecurityOption,proto3" json:"extended_security_option,omitempty"`
	CipsoOption             uint32   `protobuf:"varint,20,opt,name=cipso_option,json=cipsoOption,proto3" json:"cipso_option,omitempty"`
	StrictSourceRouteOption uint32   `protobuf:"varint,21,opt,name=strict_source_route_option,json=strictSourceRouteOption,proto3" json:"strict_source_route_option,omitempty"`
	LooseSourceRouteOption  uint32   `protobuf:"varint,22,opt,name=loose_source_route_option,json=looseSourceRouteOption,proto3" json:"loose_source_route_option,omitempty"`
	RecordRouteOption       uint32   `protobuf:"varint,23,opt,name=record_route_option,json=recordRouteOption,proto3" json:"record_route_option,omitempty"`
	SidOption               uint32   `protobuf:"varint,24,opt,name=sid_option,json=sidOption,proto3" json:"sid_option,omitempty"`
	TimestampOption         uint32   `protobuf:"varint,25,opt,name=timestamp_option,json=timestampOption,proto3" json:"timestamp_option,omitempty"`
	RouterAlertOption       uint32   `protobuf:"varint,26,opt,name=router_alert_option,json=routerAlertOption,proto3" json:"router_alert_option,omitempty"`
	NoopOption              uint32   `protobuf:"varint,27,opt,name=noop_option,json=noopOption,proto3" json:"noop_option,omitempty"`
	EndOption               uint32   `protobuf:"varint,28,opt,name=end_option,json=endOption,proto3" json:"end_option,omitempty"`
	PacketsOutput           uint32   `protobuf:"varint,29,opt,name=packets_output,json=packetsOutput,proto3" json:"packets_output,omitempty"`
	PacketsForwarded        uint32   `protobuf:"varint,30,opt,name=packets_forwarded,json=packetsForwarded,proto3" json:"packets_forwarded,omitempty"`
	PacketsFragmented       uint32   `protobuf:"varint,31,opt,name=packets_fragmented,json=packetsFragmented,proto3" json:"packets_fragmented,omitempty"`
	FragmentCount           uint32   `protobuf:"varint,32,opt,name=fragment_count,json=fragmentCount,proto3" json:"fragment_count,omitempty"`
	EncapsulationFailed     uint32   `protobuf:"varint,33,opt,name=encapsulation_failed,json=encapsulationFailed,proto3" json:"encapsulation_failed,omitempty"`
	NoRouter                uint32   `protobuf:"varint,34,opt,name=no_router,json=noRouter,proto3" json:"no_router,omitempty"`
	PacketTooBig            uint32   `protobuf:"varint,35,opt,name=packet_too_big,json=packetTooBig,proto3" json:"packet_too_big,omitempty"`
	MulticastIn             uint32   `protobuf:"varint,36,opt,name=multicast_in,json=multicastIn,proto3" json:"multicast_in,omitempty"`
	MulticastOut            uint32   `protobuf:"varint,37,opt,name=multicast_out,json=multicastOut,proto3" json:"multicast_out,omitempty"`
	BroadcastIn             uint32   `protobuf:"varint,38,opt,name=broadcast_in,json=broadcastIn,proto3" json:"broadcast_in,omitempty"`
	BroadcastOut            uint32   `protobuf:"varint,39,opt,name=broadcast_out,json=broadcastOut,proto3" json:"broadcast_out,omitempty"`
	LispV4Encap             uint32   `protobuf:"varint,40,opt,name=lisp_v4_encap,json=lispV4Encap,proto3" json:"lisp_v4_encap,omitempty"`
	LispV4Decap             uint32   `protobuf:"varint,41,opt,name=lisp_v4_decap,json=lispV4Decap,proto3" json:"lisp_v4_decap,omitempty"`
	LispV6Encap             uint32   `protobuf:"varint,42,opt,name=lisp_v6_encap,json=lispV6Encap,proto3" json:"lisp_v6_encap,omitempty"`
	LispV6Decap             uint32   `protobuf:"varint,43,opt,name=lisp_v6_decap,json=lispV6Decap,proto3" json:"lisp_v6_decap,omitempty"`
	LispEncapError          uint32   `protobuf:"varint,44,opt,name=lisp_encap_error,json=lispEncapError,proto3" json:"lisp_encap_error,omitempty"`
	LispDecapError          uint32   `protobuf:"varint,45,opt,name=lisp_decap_error,json=lispDecapError,proto3" json:"lisp_decap_error,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ipv4IoTraffic) Reset()         { *m = Ipv4IoTraffic{} }
func (m *Ipv4IoTraffic) String() string { return proto.CompactTextString(m) }
func (*Ipv4IoTraffic) ProtoMessage()    {}
func (*Ipv4IoTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a2a8e22f29b44, []int{1}
}

func (m *Ipv4IoTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IoTraffic.Unmarshal(m, b)
}
func (m *Ipv4IoTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IoTraffic.Marshal(b, m, deterministic)
}
func (m *Ipv4IoTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IoTraffic.Merge(m, src)
}
func (m *Ipv4IoTraffic) XXX_Size() int {
	return xxx_messageInfo_Ipv4IoTraffic.Size(m)
}
func (m *Ipv4IoTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IoTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IoTraffic proto.InternalMessageInfo

func (m *Ipv4IoTraffic) GetInputPackets() uint32 {
	if m != nil {
		return m.InputPackets
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *Ipv4IoTraffic) GetFormatErrors() uint32 {
	if m != nil {
		return m.FormatErrors
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadHopCount() uint32 {
	if m != nil {
		return m.BadHopCount
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadSourceAddress() uint32 {
	if m != nil {
		return m.BadSourceAddress
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadHeader() uint32 {
	if m != nil {
		return m.BadHeader
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoProtocol() uint32 {
	if m != nil {
		return m.NoProtocol
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoGateway() uint32 {
	if m != nil {
		return m.NoGateway
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleInput() uint32 {
	if m != nil {
		return m.ReassembleInput
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembled() uint32 {
	if m != nil {
		return m.Reassembled
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleTimeout() uint32 {
	if m != nil {
		return m.ReassembleTimeout
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleMaxDrop() uint32 {
	if m != nil {
		return m.ReassembleMaxDrop
	}
	return 0
}

func (m *Ipv4IoTraffic) GetReassembleFailed() uint32 {
	if m != nil {
		return m.ReassembleFailed
	}
	return 0
}

func (m *Ipv4IoTraffic) GetOptionsPresent() uint32 {
	if m != nil {
		return m.OptionsPresent
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadOption() uint32 {
	if m != nil {
		return m.BadOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetUnknownOption() uint32 {
	if m != nil {
		return m.UnknownOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBadSecurityOption() uint32 {
	if m != nil {
		return m.BadSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBasicSecurityOption() uint32 {
	if m != nil {
		return m.BasicSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetExtendedSecurityOption() uint32 {
	if m != nil {
		return m.ExtendedSecurityOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetCipsoOption() uint32 {
	if m != nil {
		return m.CipsoOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetStrictSourceRouteOption() uint32 {
	if m != nil {
		return m.StrictSourceRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLooseSourceRouteOption() uint32 {
	if m != nil {
		return m.LooseSourceRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetRecordRouteOption() uint32 {
	if m != nil {
		return m.RecordRouteOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetSidOption() uint32 {
	if m != nil {
		return m.SidOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetTimestampOption() uint32 {
	if m != nil {
		return m.TimestampOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetRouterAlertOption() uint32 {
	if m != nil {
		return m.RouterAlertOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoopOption() uint32 {
	if m != nil {
		return m.NoopOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetEndOption() uint32 {
	if m != nil {
		return m.EndOption
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsOutput() uint32 {
	if m != nil {
		return m.PacketsOutput
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsForwarded() uint32 {
	if m != nil {
		return m.PacketsForwarded
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketsFragmented() uint32 {
	if m != nil {
		return m.PacketsFragmented
	}
	return 0
}

func (m *Ipv4IoTraffic) GetFragmentCount() uint32 {
	if m != nil {
		return m.FragmentCount
	}
	return 0
}

func (m *Ipv4IoTraffic) GetEncapsulationFailed() uint32 {
	if m != nil {
		return m.EncapsulationFailed
	}
	return 0
}

func (m *Ipv4IoTraffic) GetNoRouter() uint32 {
	if m != nil {
		return m.NoRouter
	}
	return 0
}

func (m *Ipv4IoTraffic) GetPacketTooBig() uint32 {
	if m != nil {
		return m.PacketTooBig
	}
	return 0
}

func (m *Ipv4IoTraffic) GetMulticastIn() uint32 {
	if m != nil {
		return m.MulticastIn
	}
	return 0
}

func (m *Ipv4IoTraffic) GetMulticastOut() uint32 {
	if m != nil {
		return m.MulticastOut
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBroadcastIn() uint32 {
	if m != nil {
		return m.BroadcastIn
	}
	return 0
}

func (m *Ipv4IoTraffic) GetBroadcastOut() uint32 {
	if m != nil {
		return m.BroadcastOut
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV4Encap() uint32 {
	if m != nil {
		return m.LispV4Encap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV4Decap() uint32 {
	if m != nil {
		return m.LispV4Decap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV6Encap() uint32 {
	if m != nil {
		return m.LispV6Encap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispV6Decap() uint32 {
	if m != nil {
		return m.LispV6Decap
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispEncapError() uint32 {
	if m != nil {
		return m.LispEncapError
	}
	return 0
}

func (m *Ipv4IoTraffic) GetLispDecapError() uint32 {
	if m != nil {
		return m.LispDecapError
	}
	return 0
}

type Ipv4IoIcmpTraffic struct {
	Received                    uint32   `protobuf:"varint,1,opt,name=received,proto3" json:"received,omitempty"`
	ChecksumError               uint32   `protobuf:"varint,2,opt,name=checksum_error,json=checksumError,proto3" json:"checksum_error,omitempty"`
	Unknown                     uint32   `protobuf:"varint,3,opt,name=unknown,proto3" json:"unknown,omitempty"`
	Output                      uint32   `protobuf:"varint,4,opt,name=output,proto3" json:"output,omitempty"`
	AdminUnreachableSent        uint32   `protobuf:"varint,5,opt,name=admin_unreachable_sent,json=adminUnreachableSent,proto3" json:"admin_unreachable_sent,omitempty"`
	NetworkUnreachableSent      uint32   `protobuf:"varint,6,opt,name=network_unreachable_sent,json=networkUnreachableSent,proto3" json:"network_unreachable_sent,omitempty"`
	HostUnreachableSent         uint32   `protobuf:"varint,7,opt,name=host_unreachable_sent,json=hostUnreachableSent,proto3" json:"host_unreachable_sent,omitempty"`
	ProtocolUnreachableSent     uint32   `protobuf:"varint,8,opt,name=protocol_unreachable_sent,json=protocolUnreachableSent,proto3" json:"protocol_unreachable_sent,omitempty"`
	PortUnreachableSent         uint32   `protobuf:"varint,9,opt,name=port_unreachable_sent,json=portUnreachableSent,proto3" json:"port_unreachable_sent,omitempty"`
	FragmentUnreachableSent     uint32   `protobuf:"varint,10,opt,name=fragment_unreachable_sent,json=fragmentUnreachableSent,proto3" json:"fragment_unreachable_sent,omitempty"`
	AdminUnreachableReceived    uint32   `protobuf:"varint,11,opt,name=admin_unreachable_received,json=adminUnreachableReceived,proto3" json:"admin_unreachable_received,omitempty"`
	NetworkUnreachableReceived  uint32   `protobuf:"varint,12,opt,name=network_unreachable_received,json=networkUnreachableReceived,proto3" json:"network_unreachable_received,omitempty"`
	HostUnreachableReceived     uint32   `protobuf:"varint,13,opt,name=host_unreachable_received,json=hostUnreachableReceived,proto3" json:"host_unreachable_received,omitempty"`
	ProtocolUnreachableReceived uint32   `protobuf:"varint,14,opt,name=protocol_unreachable_received,json=protocolUnreachableReceived,proto3" json:"protocol_unreachable_received,omitempty"`
	PortUnreachableReceived     uint32   `protobuf:"varint,15,opt,name=port_unreachable_received,json=portUnreachableReceived,proto3" json:"port_unreachable_received,omitempty"`
	FragmentUnreachableReceived uint32   `protobuf:"varint,16,opt,name=fragment_unreachable_received,json=fragmentUnreachableReceived,proto3" json:"fragment_unreachable_received,omitempty"`
	HopcountSent                uint32   `protobuf:"varint,17,opt,name=hopcount_sent,json=hopcountSent,proto3" json:"hopcount_sent,omitempty"`
	ReassemblySent              uint32   `protobuf:"varint,18,opt,name=reassembly_sent,json=reassemblySent,proto3" json:"reassembly_sent,omitempty"`
	HopcountReceived            uint32   `protobuf:"varint,19,opt,name=hopcount_received,json=hopcountReceived,proto3" json:"hopcount_received,omitempty"`
	ReasseblyReceived           uint32   `protobuf:"varint,20,opt,name=reassebly_received,json=reasseblyReceived,proto3" json:"reassebly_received,omitempty"`
	ParamErrorReceived          uint32   `protobuf:"varint,21,opt,name=param_error_received,json=paramErrorReceived,proto3" json:"param_error_received,omitempty"`
	ParamErrorSend              uint32   `protobuf:"varint,22,opt,name=param_error_send,json=paramErrorSend,proto3" json:"param_error_send,omitempty"`
	EchoRequestSent             uint32   `protobuf:"varint,23,opt,name=echo_request_sent,json=echoRequestSent,proto3" json:"echo_request_sent,omitempty"`
	EchoRequestReceived         uint32   `protobuf:"varint,24,opt,name=echo_request_received,json=echoRequestReceived,proto3" json:"echo_request_received,omitempty"`
	EchoReplySent               uint32   `protobuf:"varint,25,opt,name=echo_reply_sent,json=echoReplySent,proto3" json:"echo_reply_sent,omitempty"`
	EchoReplyReceived           uint32   `protobuf:"varint,26,opt,name=echo_reply_received,json=echoReplyReceived,proto3" json:"echo_reply_received,omitempty"`
	MaskRequestSent             uint32   `protobuf:"varint,27,opt,name=mask_request_sent,json=maskRequestSent,proto3" json:"mask_request_sent,omitempty"`
	MaskRequestReceived         uint32   `protobuf:"varint,28,opt,name=mask_request_received,json=maskRequestReceived,proto3" json:"mask_request_received,omitempty"`
	MaskReplySent               uint32   `protobuf:"varint,29,opt,name=mask_reply_sent,json=maskReplySent,proto3" json:"mask_reply_sent,omitempty"`
	MaskReplyReceived           uint32   `protobuf:"varint,30,opt,name=mask_reply_received,json=maskReplyReceived,proto3" json:"mask_reply_received,omitempty"`
	SourceQuenchReceived        uint32   `protobuf:"varint,31,opt,name=source_quench_received,json=sourceQuenchReceived,proto3" json:"source_quench_received,omitempty"`
	RedirectReceived            uint32   `protobuf:"varint,32,opt,name=redirect_received,json=redirectReceived,proto3" json:"redirect_received,omitempty"`
	RedirectSend                uint32   `protobuf:"varint,33,opt,name=redirect_send,json=redirectSend,proto3" json:"redirect_send,omitempty"`
	TimestampReceived           uint32   `protobuf:"varint,34,opt,name=timestamp_received,json=timestampReceived,proto3" json:"timestamp_received,omitempty"`
	TimestampReplyReceived      uint32   `protobuf:"varint,35,opt,name=timestamp_reply_received,json=timestampReplyReceived,proto3" json:"timestamp_reply_received,omitempty"`
	RouterAdvertReceived        uint32   `protobuf:"varint,36,opt,name=router_advert_received,json=routerAdvertReceived,proto3" json:"router_advert_received,omitempty"`
	RouterSolicitReceived       uint32   `protobuf:"varint,37,opt,name=router_solicit_received,json=routerSolicitReceived,proto3" json:"router_solicit_received,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *Ipv4IoIcmpTraffic) Reset()         { *m = Ipv4IoIcmpTraffic{} }
func (m *Ipv4IoIcmpTraffic) String() string { return proto.CompactTextString(m) }
func (*Ipv4IoIcmpTraffic) ProtoMessage()    {}
func (*Ipv4IoIcmpTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a2a8e22f29b44, []int{2}
}

func (m *Ipv4IoIcmpTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Unmarshal(m, b)
}
func (m *Ipv4IoIcmpTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Marshal(b, m, deterministic)
}
func (m *Ipv4IoIcmpTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IoIcmpTraffic.Merge(m, src)
}
func (m *Ipv4IoIcmpTraffic) XXX_Size() int {
	return xxx_messageInfo_Ipv4IoIcmpTraffic.Size(m)
}
func (m *Ipv4IoIcmpTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IoIcmpTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IoIcmpTraffic proto.InternalMessageInfo

func (m *Ipv4IoIcmpTraffic) GetReceived() uint32 {
	if m != nil {
		return m.Received
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetChecksumError() uint32 {
	if m != nil {
		return m.ChecksumError
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetUnknown() uint32 {
	if m != nil {
		return m.Unknown
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetOutput() uint32 {
	if m != nil {
		return m.Output
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetAdminUnreachableSent() uint32 {
	if m != nil {
		return m.AdminUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetNetworkUnreachableSent() uint32 {
	if m != nil {
		return m.NetworkUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHostUnreachableSent() uint32 {
	if m != nil {
		return m.HostUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetProtocolUnreachableSent() uint32 {
	if m != nil {
		return m.ProtocolUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetPortUnreachableSent() uint32 {
	if m != nil {
		return m.PortUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetFragmentUnreachableSent() uint32 {
	if m != nil {
		return m.FragmentUnreachableSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetAdminUnreachableReceived() uint32 {
	if m != nil {
		return m.AdminUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetNetworkUnreachableReceived() uint32 {
	if m != nil {
		return m.NetworkUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHostUnreachableReceived() uint32 {
	if m != nil {
		return m.HostUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetProtocolUnreachableReceived() uint32 {
	if m != nil {
		return m.ProtocolUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetPortUnreachableReceived() uint32 {
	if m != nil {
		return m.PortUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetFragmentUnreachableReceived() uint32 {
	if m != nil {
		return m.FragmentUnreachableReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHopcountSent() uint32 {
	if m != nil {
		return m.HopcountSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetReassemblySent() uint32 {
	if m != nil {
		return m.ReassemblySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetHopcountReceived() uint32 {
	if m != nil {
		return m.HopcountReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetReasseblyReceived() uint32 {
	if m != nil {
		return m.ReasseblyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetParamErrorReceived() uint32 {
	if m != nil {
		return m.ParamErrorReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetParamErrorSend() uint32 {
	if m != nil {
		return m.ParamErrorSend
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoRequestSent() uint32 {
	if m != nil {
		return m.EchoRequestSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoRequestReceived() uint32 {
	if m != nil {
		return m.EchoRequestReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoReplySent() uint32 {
	if m != nil {
		return m.EchoReplySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetEchoReplyReceived() uint32 {
	if m != nil {
		return m.EchoReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskRequestSent() uint32 {
	if m != nil {
		return m.MaskRequestSent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskRequestReceived() uint32 {
	if m != nil {
		return m.MaskRequestReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskReplySent() uint32 {
	if m != nil {
		return m.MaskReplySent
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetMaskReplyReceived() uint32 {
	if m != nil {
		return m.MaskReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetSourceQuenchReceived() uint32 {
	if m != nil {
		return m.SourceQuenchReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRedirectReceived() uint32 {
	if m != nil {
		return m.RedirectReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRedirectSend() uint32 {
	if m != nil {
		return m.RedirectSend
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetTimestampReceived() uint32 {
	if m != nil {
		return m.TimestampReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetTimestampReplyReceived() uint32 {
	if m != nil {
		return m.TimestampReplyReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRouterAdvertReceived() uint32 {
	if m != nil {
		return m.RouterAdvertReceived
	}
	return 0
}

func (m *Ipv4IoIcmpTraffic) GetRouterSolicitReceived() uint32 {
	if m != nil {
		return m.RouterSolicitReceived
	}
	return 0
}

type IpTraffic struct {
	Ipv4Stats            *Ipv4IoTraffic     `protobuf:"bytes,50,opt,name=ipv4_stats,json=ipv4Stats,proto3" json:"ipv4_stats,omitempty"`
	IcmpStats            *Ipv4IoIcmpTraffic `protobuf:"bytes,51,opt,name=icmp_stats,json=icmpStats,proto3" json:"icmp_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IpTraffic) Reset()         { *m = IpTraffic{} }
func (m *IpTraffic) String() string { return proto.CompactTextString(m) }
func (*IpTraffic) ProtoMessage()    {}
func (*IpTraffic) Descriptor() ([]byte, []int) {
	return fileDescriptor_e13a2a8e22f29b44, []int{3}
}

func (m *IpTraffic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpTraffic.Unmarshal(m, b)
}
func (m *IpTraffic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpTraffic.Marshal(b, m, deterministic)
}
func (m *IpTraffic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpTraffic.Merge(m, src)
}
func (m *IpTraffic) XXX_Size() int {
	return xxx_messageInfo_IpTraffic.Size(m)
}
func (m *IpTraffic) XXX_DiscardUnknown() {
	xxx_messageInfo_IpTraffic.DiscardUnknown(m)
}

var xxx_messageInfo_IpTraffic proto.InternalMessageInfo

func (m *IpTraffic) GetIpv4Stats() *Ipv4IoTraffic {
	if m != nil {
		return m.Ipv4Stats
	}
	return nil
}

func (m *IpTraffic) GetIcmpStats() *Ipv4IoIcmpTraffic {
	if m != nil {
		return m.IcmpStats
	}
	return nil
}

func init() {
	proto.RegisterType((*IpTraffic_KEYS)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ip_traffic_KEYS")
	proto.RegisterType((*Ipv4IoTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ipv4_io_traffic")
	proto.RegisterType((*Ipv4IoIcmpTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ipv4_io_icmp_traffic")
	proto.RegisterType((*IpTraffic)(nil), "cisco_ios_xr_ipv4_io_oper.ipv4_network.nodes.node.statistics.traffic.ip_traffic")
}

func init() { proto.RegisterFile("ip_traffic.proto", fileDescriptor_e13a2a8e22f29b44) }

var fileDescriptor_e13a2a8e22f29b44 = []byte{
	// 1486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0x5b, 0x4f, 0x1c, 0x39,
	0x16, 0x16, 0x7b, 0x49, 0xe0, 0x40, 0xd3, 0x4d, 0x75, 0x03, 0x05, 0x84, 0x0d, 0x81, 0x5c, 0xc8,
	0xad, 0xb5, 0x4b, 0x10, 0xca, 0x66, 0xf7, 0x61, 0x93, 0x4d, 0xb2, 0x1b, 0x8d, 0x66, 0x92, 0x81,
	0x64, 0xa4, 0x99, 0x17, 0xcb, 0x5d, 0x65, 0xc0, 0xa2, 0xbb, 0x5c, 0xb1, 0x5d, 0x5c, 0xfe, 0xd5,
	0x3c, 0xcd, 0xbf, 0x1b, 0x69, 0xe4, 0x73, 0x6c, 0x57, 0xf5, 0xe5, 0x31, 0x2f, 0x48, 0xfd, 0x9d,
	0xef, 0xfb, 0x7c, 0xea, 0xf8, 0x72, 0x0e, 0xd0, 0x91, 0x25, 0xb3, 0x9a, 0x9f, 0x9e, 0xca, 0xac,
	0x5f, 0x6a, 0x65, 0x55, 0xf2, 0x36, 0x93, 0x26, 0x53, 0x4c, 0x2a, 0xc3, 0xae, 0x35, 0x93, 0xe5,
	0xe5, 0x21, 0x93, 0x8a, 0xa9, 0x52, 0xe8, 0x3e, 0xfe, 0x28, 0x84, 0xbd, 0x52, 0xfa, 0xa2, 0x5f,
	0xa8, 0x5c, 0x18, 0xfc, 0xdb, 0x37, 0x96, 0x5b, 0x69, 0xac, 0xcc, 0x4c, 0xdf, 0x7b, 0xed, 0xf6,
	0xa1, 0x5d, 0x3b, 0xb3, 0xef, 0xde, 0xfd, 0x7c, 0x92, 0x6c, 0xc1, 0x82, 0x63, 0xb3, 0x82, 0x8f,
	0x44, 0x3a, 0xb7, 0x33, 0xb7, 0xbf, 0x70, 0x3c, 0xef, 0x80, 0x1f, 0xf8, 0x48, 0xec, 0xfe, 0xb6,
	0xec, 0x04, 0xb4, 0x96, 0x57, 0x25, 0x7b, 0xd0, 0x92, 0x45, 0x59, 0x59, 0x56, 0xf2, 0xec, 0x42,
	0x58, 0x83, 0xa2, 0xd6, 0xf1, 0x12, 0x82, 0x9f, 0x08, 0x4b, 0x1e, 0x43, 0x47, 0x8b, 0x4c, 0xc8,
	0x4b, 0x91, 0x47, 0xde, 0x9f, 0x90, 0xd7, 0x0e, 0x78, 0xa0, 0xee, 0x41, 0xeb, 0x54, 0xe9, 0x11,
	0xb7, 0x4c, 0x68, 0xad, 0xb4, 0x49, 0xff, 0x4c, 0x7e, 0x04, 0xbe, 0x43, 0x2c, 0xd9, 0x85, 0xd6,
	0x80, 0xe7, 0xec, 0x5c, 0x95, 0x2c, 0x53, 0x55, 0x61, 0xd3, 0xbf, 0x20, 0x69, 0x71, 0xc0, 0xf3,
	0xff, 0xab, 0xf2, 0xbf, 0x0e, 0x4a, 0x9e, 0x41, 0xe2, 0x38, 0x46, 0x55, 0x3a, 0x13, 0x8c, 0xe7,
	0xb9, 0x16, 0xc6, 0xa4, 0x7f, 0x45, 0x62, 0x67, 0xc0, 0xf3, 0x13, 0x0c, 0xbc, 0x26, 0x3c, 0xd9,
	0x06, 0x40, 0x47, 0xc1, 0x73, 0xa1, 0xd3, 0x5b, 0xc8, 0x5a, 0x70, 0x76, 0x08, 0x24, 0x77, 0x61,
	0xb1, 0x50, 0x0c, 0x6b, 0x9f, 0xa9, 0x61, 0x7a, 0x1b, 0xe3, 0x50, 0xa8, 0x4f, 0x1e, 0x71, 0xfa,
	0x42, 0xb1, 0x33, 0x6e, 0xc5, 0x15, 0xbf, 0x49, 0xe7, 0x49, 0x5f, 0xa8, 0xff, 0x11, 0x40, 0x05,
	0xe0, 0xc6, 0x88, 0xd1, 0x60, 0x28, 0x18, 0xd6, 0x26, 0x5d, 0x08, 0x05, 0x08, 0xf8, 0x07, 0x07,
	0x27, 0x3b, 0xb0, 0x58, 0x43, 0x79, 0x0a, 0xf4, 0x65, 0x0d, 0x28, 0x79, 0x0e, 0x49, 0xc3, 0xcc,
	0xca, 0x91, 0x50, 0x95, 0x4d, 0x17, 0x91, 0xb8, 0x52, 0x47, 0x3e, 0x53, 0x20, 0xe9, 0x43, 0xb7,
	0x41, 0x1f, 0xf1, 0x6b, 0x96, 0x6b, 0x55, 0xa6, 0x4b, 0x93, 0xfc, 0xef, 0xf9, 0xf5, 0x5b, 0xad,
	0xca, 0xe4, 0x29, 0x34, 0x40, 0x76, 0xca, 0xa5, 0x4b, 0xa3, 0x45, 0x75, 0xab, 0x03, 0xef, 0x11,
	0x4f, 0x1e, 0x41, 0x5b, 0x95, 0x56, 0xaa, 0xc2, 0xb0, 0x52, 0x0b, 0x23, 0x0a, 0x9b, 0x2e, 0x23,
	0x75, 0xd9, 0xc3, 0x9f, 0x08, 0x0d, 0x05, 0x26, 0x34, 0x6d, 0xc7, 0x02, 0x7f, 0x44, 0x20, 0x79,
	0x00, 0xcb, 0x55, 0x71, 0x51, 0xa8, 0xab, 0x22, 0x50, 0x3a, 0x48, 0x69, 0x79, 0xd4, 0xd3, 0xfa,
	0xd0, 0xc5, 0x4d, 0x15, 0x59, 0xa5, 0xa5, 0xbd, 0x09, 0xdc, 0x15, 0xfa, 0x16, 0xb7, 0xab, 0x3e,
	0xe2, 0xf9, 0x07, 0xb0, 0x3a, 0xe0, 0x46, 0x66, 0x53, 0x8a, 0x04, 0x15, 0x5d, 0x0c, 0x4e, 0x68,
	0x5e, 0x42, 0x2a, 0xae, 0xad, 0x28, 0x72, 0x31, 0xbd, 0x50, 0x17, 0x65, 0x6b, 0x21, 0x3e, 0xa1,
	0xbc, 0x07, 0x4b, 0x99, 0x2c, 0x8d, 0x0a, 0xec, 0x1e, 0xed, 0x1d, 0x62, 0x9e, 0xf2, 0x2f, 0xd8,
	0x34, 0x56, 0xcb, 0xcc, 0x86, 0x83, 0xa9, 0x55, 0x65, 0x45, 0x10, 0xac, 0xa2, 0x60, 0x9d, 0x18,
	0x74, 0x40, 0x8f, 0x5d, 0xdc, 0x8b, 0xff, 0x09, 0x1b, 0x43, 0xa5, 0x8c, 0x98, 0xa9, 0x5d, 0xa3,
	0xd4, 0x90, 0x30, 0x2d, 0xc5, 0x43, 0x90, 0x29, 0x9d, 0x8f, 0x8b, 0xd6, 0xc3, 0x21, 0x70, 0xa1,
	0x26, 0x7f, 0x1b, 0xc0, 0xc8, 0xb8, 0x5d, 0x29, 0x6d, 0x97, 0x91, 0x61, 0xbb, 0x1e, 0x43, 0xc7,
	0x9d, 0x3b, 0x63, 0xf9, 0xa8, 0x0c, 0xa4, 0x0d, 0x3a, 0xcf, 0x11, 0x6f, 0xac, 0xec, 0x8c, 0x35,
	0xe3, 0x43, 0xa1, 0x6d, 0x60, 0x6f, 0xfa, 0x95, 0x31, 0xf4, 0xda, 0x45, 0x3c, 0x1f, 0xaf, 0x9a,
	0x8a, 0xae, 0x5b, 0xe1, 0xaa, 0xa9, 0xb2, 0x4e, 0x4d, 0x14, 0x31, 0xb5, 0x3b, 0x94, 0x9a, 0x28,
	0x1a, 0x27, 0xc9, 0x3f, 0x31, 0x4c, 0x55, 0xd6, 0x5d, 0xb4, 0x6d, 0x3a, 0x49, 0x1e, 0xfd, 0x88,
	0xa0, 0x3b, 0xe5, 0x81, 0x76, 0xaa, 0xf4, 0x15, 0xd7, 0xb9, 0xc8, 0xd3, 0xbf, 0xd1, 0x29, 0xf7,
	0x81, 0xf7, 0x01, 0x77, 0x37, 0x2e, 0x92, 0x35, 0x3f, 0x1b, 0x89, 0xc2, 0x8a, 0x3c, 0xbd, 0x4b,
	0x9f, 0x10, 0xd8, 0x31, 0xe0, 0x52, 0x08, 0x34, 0xff, 0x3e, 0xed, 0x50, 0x0a, 0x01, 0xa5, 0x17,
	0xea, 0x1f, 0xd0, 0x13, 0x45, 0xc6, 0x4b, 0x53, 0x0d, 0xb9, 0x4b, 0x3d, 0xdc, 0xb5, 0x7b, 0x74,
	0x36, 0xc7, 0x62, 0xfe, 0xba, 0xe1, 0xf3, 0x4c, 0x5b, 0xa8, 0xd3, 0x5d, 0xe4, 0xcd, 0x17, 0x0a,
	0x37, 0x4e, 0x27, 0xf7, 0xc3, 0x97, 0x33, 0xab, 0x14, 0x1b, 0xc8, 0xb3, 0x74, 0x8f, 0xde, 0x4e,
	0x42, 0x3f, 0x2b, 0xf5, 0x46, 0x9e, 0xb9, 0x43, 0x3a, 0xaa, 0x86, 0x56, 0x66, 0xdc, 0x58, 0x26,
	0x8b, 0xf4, 0x3e, 0x1d, 0xd2, 0x88, 0x7d, 0x28, 0xdc, 0x1b, 0x5c, 0x53, 0xdc, 0xdb, 0xf2, 0x80,
	0x7c, 0x22, 0xf8, 0xb1, 0xb2, 0xce, 0x67, 0xa0, 0x15, 0xcf, 0x83, 0xcf, 0x43, 0xff, 0x04, 0x07,
	0x8c, 0x7c, 0x6a, 0x8a, 0xf3, 0x79, 0x44, 0x3e, 0x11, 0x74, 0x3e, 0xbb, 0xd0, 0x1a, 0x4a, 0x53,
	0xb2, 0xcb, 0x43, 0x86, 0x5f, 0x9c, 0xee, 0x93, 0x91, 0x03, 0x7f, 0x3a, 0x7c, 0xe7, 0xa0, 0x26,
	0x27, 0x17, 0x8e, 0xf3, 0xb8, 0xc9, 0x79, 0x2b, 0xc6, 0x38, 0x47, 0xde, 0xe7, 0x49, 0x83, 0x73,
	0x34, 0xe1, 0x73, 0xe4, 0x7d, 0x9e, 0x36, 0x39, 0xe4, 0xb3, 0x0f, 0x1d, 0xe4, 0xa0, 0x09, 0x35,
	0xa1, 0xf4, 0x19, 0x3d, 0x69, 0x0e, 0x47, 0x23, 0x6c, 0x43, 0x91, 0x89, 0x56, 0x9e, 0xf9, 0xbc,
	0x66, 0xa2, 0x1d, 0x32, 0x77, 0x7f, 0x5d, 0x86, 0x5e, 0x68, 0x9c, 0x32, 0x1b, 0xc5, 0x9e, 0x9b,
	0x6c, 0xc2, 0x7c, 0x68, 0x80, 0xbe, 0x71, 0xc6, 0xdf, 0xee, 0x14, 0x65, 0xe7, 0x22, 0xbb, 0x30,
	0xd5, 0xc8, 0x9b, 0x53, 0xcb, 0x6c, 0x05, 0x94, 0xb2, 0x48, 0xe1, 0xb6, 0x7f, 0x23, 0x7d, 0xab,
	0x0c, 0x3f, 0x93, 0x35, 0xb8, 0xe5, 0x6f, 0x00, 0xb5, 0x47, 0xff, 0x2b, 0x39, 0x84, 0x35, 0x9e,
	0x8f, 0x64, 0xc1, 0xaa, 0x42, 0x0b, 0x9e, 0x9d, 0x73, 0xf7, 0xce, 0xe3, 0xd3, 0x4d, 0xdd, 0xb1,
	0x87, 0xd1, 0x2f, 0x75, 0xf0, 0xc4, 0x3d, 0xe0, 0x2f, 0x21, 0xf5, 0x53, 0xc5, 0xb4, 0x8e, 0xfa,
	0xe5, 0x9a, 0x8f, 0x4f, 0x2a, 0x0f, 0x60, 0xf5, 0x5c, 0x19, 0x3b, 0x2d, 0xa3, 0x36, 0xda, 0x75,
	0xc1, 0x49, 0xcd, 0x2b, 0xd8, 0x08, 0xdd, 0x76, 0x5a, 0x47, 0xed, 0x75, 0x3d, 0x10, 0x66, 0xac,
	0x57, 0x2a, 0x3d, 0x63, 0x3d, 0xea, 0xb8, 0x5d, 0x17, 0x9c, 0xb1, 0x5e, 0xbc, 0xb2, 0x53, 0x3a,
	0xea, 0xc1, 0xeb, 0x81, 0x30, 0xa9, 0xfd, 0x37, 0x6c, 0x4e, 0xd7, 0x33, 0x6e, 0x2b, 0xf5, 0xe5,
	0x74, 0xb2, 0xa6, 0xc7, 0x61, 0x9b, 0xff, 0x03, 0x77, 0x66, 0xd5, 0x35, 0xea, 0xa9, 0x4f, 0x6f,
	0x4e, 0xd7, 0x36, 0x3a, 0xbc, 0x82, 0x8d, 0xa9, 0xfa, 0x46, 0x39, 0x35, 0xee, 0xf5, 0x89, 0x1a,
	0x47, 0xed, 0x1b, 0xd8, 0x9e, 0x59, 0xe7, 0xa8, 0xa7, 0x6e, 0xbe, 0x35, 0xa3, 0xd6, 0xcd, 0xf5,
	0xa7, 0xea, 0x1d, 0xf5, 0x6d, 0xbf, 0x57, 0xe3, 0x35, 0x6f, 0xae, 0x3f, 0xb3, 0xee, 0x51, 0x4f,
	0x63, 0xc0, 0xd6, 0x8c, 0xda, 0x47, 0x8f, 0x3d, 0x68, 0x9d, 0xab, 0x12, 0x1f, 0x5a, 0xda, 0x2f,
	0x1a, 0x07, 0x96, 0x02, 0x88, 0x9b, 0xf4, 0x08, 0xea, 0x49, 0xeb, 0x86, 0x68, 0x34, 0x03, 0x2c,
	0xd7, 0x30, 0x12, 0x9f, 0xc2, 0x4a, 0x74, 0x8b, 0x59, 0x50, 0xdf, 0xef, 0x84, 0x40, 0x5c, 0x3a,
	0x8e, 0x62, 0xce, 0x34, 0xb2, 0x7b, 0xcd, 0xd1, 0x6a, 0x30, 0xbc, 0x89, 0xf4, 0xbf, 0x43, 0xaf,
	0xe4, 0x9a, 0xfb, 0xfb, 0x5c, 0x0b, 0xa8, 0xef, 0x27, 0x18, 0xc3, 0x5b, 0x1d, 0x15, 0xfb, 0xd0,
	0x69, 0x2a, 0x8c, 0x28, 0x72, 0xdf, 0xe9, 0x97, 0x6b, 0xf6, 0x89, 0x28, 0xf2, 0xe4, 0x09, 0xac,
	0x88, 0xec, 0x5c, 0x31, 0x2d, 0xbe, 0x56, 0xc2, 0xf8, 0x4a, 0x50, 0x7f, 0x6f, 0xbb, 0xc0, 0x31,
	0xe1, 0xe1, 0x86, 0x8c, 0x71, 0x63, 0x22, 0xa9, 0x6f, 0x3d, 0x35, 0x3f, 0x66, 0xf2, 0x10, 0xda,
	0x5e, 0x53, 0x86, 0x02, 0x52, 0xc7, 0x6f, 0x11, 0xbb, 0xf4, 0xf5, 0xeb, 0x43, 0xb7, 0xc1, 0x8b,
	0xce, 0xbe, 0xdf, 0x47, 0x6e, 0xf4, 0x7d, 0x02, 0x2b, 0x23, 0x6e, 0x2e, 0xc6, 0xf3, 0xa6, 0xae,
	0xdf, 0x76, 0x81, 0x89, 0xbc, 0xc7, 0xb8, 0xd1, 0x9d, 0xa6, 0x80, 0x6e, 0x83, 0xdf, 0xcc, 0xdb,
	0x6b, 0x62, 0xde, 0x7e, 0x20, 0x20, 0x76, 0x23, 0xef, 0x06, 0x2f, 0x3a, 0xd3, 0x48, 0xb0, 0x12,
	0xb9, 0xd1, 0xf7, 0x10, 0xd6, 0xfc, 0x18, 0xf6, 0xb5, 0x12, 0x45, 0x76, 0x5e, 0x4b, 0x68, 0x2e,
	0xe8, 0x51, 0xf4, 0x47, 0x0c, 0x46, 0x15, 0x0e, 0xd7, 0xb9, 0xd4, 0x22, 0x6b, 0x64, 0xbf, 0x13,
	0x86, 0x6b, 0x0a, 0x34, 0x0f, 0x76, 0x24, 0xe3, 0xce, 0xd3, 0x64, 0xb0, 0x14, 0x40, 0xdc, 0xf7,
	0xe7, 0x90, 0xd4, 0xa3, 0x58, 0xb4, 0xa4, 0xd9, 0x60, 0x25, 0x46, 0xa2, 0xe7, 0x4b, 0x48, 0x9b,
	0xf4, 0xb1, 0x6f, 0xa5, 0x71, 0x61, 0xad, 0x21, 0x9a, 0xf8, 0xe0, 0x30, 0xc8, 0xe5, 0x97, 0x6e,
	0x92, 0x8b, 0x3a, 0x1a, 0x21, 0x7a, 0x7e, 0x96, 0xc3, 0x60, 0x54, 0x1d, 0xc1, 0xba, 0x57, 0x19,
	0x35, 0x94, 0x99, 0x6c, 0xc8, 0x68, 0xaa, 0x58, 0xa5, 0xf0, 0x09, 0x45, 0x83, 0x6e, 0xf7, 0xf7,
	0x39, 0x80, 0xfa, 0x9f, 0xd3, 0xc4, 0xba, 0x5f, 0x97, 0x87, 0xcc, 0xfd, 0x17, 0x6b, 0xd2, 0x83,
	0x9d, 0xb9, 0xfd, 0xc5, 0x83, 0x2f, 0xfd, 0x6f, 0xf1, 0x5f, 0x70, 0x7f, 0xe2, 0x3f, 0xda, 0xe3,
	0x05, 0x07, 0x9c, 0xb8, 0x75, 0x92, 0x1b, 0x00, 0x6c, 0xd7, 0xb4, 0xea, 0x0b, 0x5c, 0xf5, 0x97,
	0x6f, 0xbb, 0x6a, 0x73, 0x1c, 0x38, 0x5e, 0x70, 0xbf, 0x70, 0xe9, 0xc1, 0x2d, 0x7c, 0x71, 0x5f,
	0xfc, 0x11, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x0f, 0x9f, 0xb0, 0xfc, 0x0f, 0x00, 0x00,
}
