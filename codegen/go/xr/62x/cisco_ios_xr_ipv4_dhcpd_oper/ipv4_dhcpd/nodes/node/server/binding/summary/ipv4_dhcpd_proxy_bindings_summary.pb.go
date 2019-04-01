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
// source: ipv4_dhcpd_proxy_bindings_summary.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_server_binding_summary

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

type Ipv4DhcpdProxyBindingsSummary_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) Reset()         { *m = Ipv4DhcpdProxyBindingsSummary_KEYS{} }
func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyBindingsSummary_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdProxyBindingsSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_95e5a5b343731df9, []int{0}
}

func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS.Size(m)
}
func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyBindingsSummary_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

type Ipv4DhcpdProxyBindingsSummary struct {
	Clients                 uint32   `protobuf:"varint,50,opt,name=clients,proto3" json:"clients,omitempty"`
	InitializingClients     uint32   `protobuf:"varint,51,opt,name=initializing_clients,json=initializingClients,proto3" json:"initializing_clients,omitempty"`
	WaitingForDpmInit       uint32   `protobuf:"varint,52,opt,name=waiting_for_dpm_init,json=waitingForDpmInit,proto3" json:"waiting_for_dpm_init,omitempty"`
	WaitingForDapsInit      uint32   `protobuf:"varint,53,opt,name=waiting_for_daps_init,json=waitingForDapsInit,proto3" json:"waiting_for_daps_init,omitempty"`
	SelectingClients        uint32   `protobuf:"varint,54,opt,name=selecting_clients,json=selectingClients,proto3" json:"selecting_clients,omitempty"`
	OfferSentForClient      uint32   `protobuf:"varint,55,opt,name=offer_sent_for_client,json=offerSentForClient,proto3" json:"offer_sent_for_client,omitempty"`
	RequestingClients       uint32   `protobuf:"varint,56,opt,name=requesting_clients,json=requestingClients,proto3" json:"requesting_clients,omitempty"`
	RequestWaitingForDpm    uint32   `protobuf:"varint,57,opt,name=request_waiting_for_dpm,json=requestWaitingForDpm,proto3" json:"request_waiting_for_dpm,omitempty"`
	AckWaitingForDpm        uint32   `protobuf:"varint,58,opt,name=ack_waiting_for_dpm,json=ackWaitingForDpm,proto3" json:"ack_waiting_for_dpm,omitempty"`
	BoundClients            uint32   `protobuf:"varint,59,opt,name=bound_clients,json=boundClients,proto3" json:"bound_clients,omitempty"`
	RenewingClients         uint32   `protobuf:"varint,60,opt,name=renewing_clients,json=renewingClients,proto3" json:"renewing_clients,omitempty"`
	InformingClients        uint32   `protobuf:"varint,61,opt,name=informing_clients,json=informingClients,proto3" json:"informing_clients,omitempty"`
	ReauthorizingClients    uint32   `protobuf:"varint,62,opt,name=reauthorizing_clients,json=reauthorizingClients,proto3" json:"reauthorizing_clients,omitempty"`
	WaitingForDpmDisconnect uint32   `protobuf:"varint,63,opt,name=waiting_for_dpm_disconnect,json=waitingForDpmDisconnect,proto3" json:"waiting_for_dpm_disconnect,omitempty"`
	WaitingForDpmAddrChange uint32   `protobuf:"varint,64,opt,name=waiting_for_dpm_addr_change,json=waitingForDpmAddrChange,proto3" json:"waiting_for_dpm_addr_change,omitempty"`
	DeletingClientsD        uint32   `protobuf:"varint,65,opt,name=deleting_clients_d,json=deletingClientsD,proto3" json:"deleting_clients_d,omitempty"`
	DisconnectedClients     uint32   `protobuf:"varint,66,opt,name=disconnected_clients,json=disconnectedClients,proto3" json:"disconnected_clients,omitempty"`
	RestartingClients       uint32   `protobuf:"varint,67,opt,name=restarting_clients,json=restartingClients,proto3" json:"restarting_clients,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyBindingsSummary) Reset()         { *m = Ipv4DhcpdProxyBindingsSummary{} }
func (m *Ipv4DhcpdProxyBindingsSummary) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyBindingsSummary) ProtoMessage()    {}
func (*Ipv4DhcpdProxyBindingsSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_95e5a5b343731df9, []int{1}
}

func (m *Ipv4DhcpdProxyBindingsSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyBindingsSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyBindingsSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary.Merge(m, src)
}
func (m *Ipv4DhcpdProxyBindingsSummary) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary.Size(m)
}
func (m *Ipv4DhcpdProxyBindingsSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyBindingsSummary proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyBindingsSummary) GetClients() uint32 {
	if m != nil {
		return m.Clients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetInitializingClients() uint32 {
	if m != nil {
		return m.InitializingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetWaitingForDpmInit() uint32 {
	if m != nil {
		return m.WaitingForDpmInit
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetWaitingForDapsInit() uint32 {
	if m != nil {
		return m.WaitingForDapsInit
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetSelectingClients() uint32 {
	if m != nil {
		return m.SelectingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetOfferSentForClient() uint32 {
	if m != nil {
		return m.OfferSentForClient
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetRequestingClients() uint32 {
	if m != nil {
		return m.RequestingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetRequestWaitingForDpm() uint32 {
	if m != nil {
		return m.RequestWaitingForDpm
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetAckWaitingForDpm() uint32 {
	if m != nil {
		return m.AckWaitingForDpm
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetBoundClients() uint32 {
	if m != nil {
		return m.BoundClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetRenewingClients() uint32 {
	if m != nil {
		return m.RenewingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetInformingClients() uint32 {
	if m != nil {
		return m.InformingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetReauthorizingClients() uint32 {
	if m != nil {
		return m.ReauthorizingClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetWaitingForDpmDisconnect() uint32 {
	if m != nil {
		return m.WaitingForDpmDisconnect
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetWaitingForDpmAddrChange() uint32 {
	if m != nil {
		return m.WaitingForDpmAddrChange
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetDeletingClientsD() uint32 {
	if m != nil {
		return m.DeletingClientsD
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetDisconnectedClients() uint32 {
	if m != nil {
		return m.DisconnectedClients
	}
	return 0
}

func (m *Ipv4DhcpdProxyBindingsSummary) GetRestartingClients() uint32 {
	if m != nil {
		return m.RestartingClients
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4DhcpdProxyBindingsSummary_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.binding.summary.ipv4_dhcpd_proxy_bindings_summary_KEYS")
	proto.RegisterType((*Ipv4DhcpdProxyBindingsSummary)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.binding.summary.ipv4_dhcpd_proxy_bindings_summary")
}

func init() {
	proto.RegisterFile("ipv4_dhcpd_proxy_bindings_summary.proto", fileDescriptor_95e5a5b343731df9)
}

var fileDescriptor_95e5a5b343731df9 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x5d, 0x6f, 0xd3, 0x3e,
	0x14, 0xc6, 0xd5, 0x9b, 0x4d, 0xff, 0xa3, 0xff, 0xc4, 0xea, 0xb5, 0x2c, 0x82, 0x9b, 0x31, 0x24,
	0x18, 0x1a, 0x0b, 0x1a, 0xdd, 0x78, 0xdb, 0x80, 0x8d, 0x96, 0x49, 0x13, 0x77, 0xdb, 0x05, 0xe2,
	0xca, 0x72, 0xe3, 0xd3, 0xd6, 0x5a, 0x6b, 0x87, 0x63, 0x77, 0x2f, 0x7c, 0x00, 0x3e, 0x37, 0x8a,
	0x13, 0xa7, 0x4e, 0x41, 0xda, 0x4d, 0x25, 0xfb, 0x79, 0x7e, 0xe7, 0xa5, 0x79, 0x12, 0x78, 0xae,
	0xf2, 0xeb, 0x03, 0x2e, 0x27, 0x59, 0x2e, 0x79, 0x4e, 0xe6, 0xf6, 0x8e, 0x0f, 0x95, 0x96, 0x4a,
	0x8f, 0x2d, 0xb7, 0xf3, 0xd9, 0x4c, 0xd0, 0x5d, 0x9a, 0x93, 0x71, 0x86, 0x9d, 0x67, 0xca, 0x66,
	0x86, 0x2b, 0x63, 0xf9, 0x2d, 0xf1, 0x88, 0x32, 0x39, 0x52, 0xba, 0x38, 0xa7, 0xda, 0x48, 0xb4,
	0xfe, 0x37, 0xb5, 0x48, 0xd7, 0x48, 0x69, 0x55, 0x31, 0xad, 0x0a, 0x6e, 0x9f, 0xc0, 0xb3, 0x7b,
	0xbb, 0xf2, 0x6f, 0x5f, 0x7f, 0x5c, 0xb2, 0x87, 0xb0, 0x52, 0x14, 0x52, 0x32, 0x69, 0x6d, 0xb5,
	0x76, 0xfe, 0xbb, 0xa8, 0x4e, 0xdb, 0xbf, 0x57, 0xe1, 0xc9, 0xbd, 0x25, 0x58, 0x02, 0xab, 0xd9,
	0x54, 0xa1, 0x76, 0x36, 0x79, 0xbd, 0xd5, 0xda, 0x59, 0xbb, 0x08, 0x47, 0xb6, 0x0f, 0x1d, 0xa5,
	0x95, 0x53, 0x62, 0xaa, 0x7e, 0x29, 0x3d, 0xe6, 0xc1, 0xd6, 0xf3, 0xb6, 0x8d, 0x58, 0xeb, 0x57,
	0xc8, 0x2b, 0xe8, 0xdc, 0x08, 0xe5, 0x0a, 0xf7, 0xc8, 0x10, 0x97, 0xf9, 0x8c, 0x17, 0xb6, 0xe4,
	0xc0, 0x23, 0xed, 0x4a, 0x3b, 0x33, 0x34, 0xc8, 0x67, 0xe7, 0x5a, 0x39, 0xb6, 0x0f, 0xdd, 0x06,
	0x20, 0x72, 0x5b, 0x12, 0x87, 0x9e, 0x60, 0x11, 0x21, 0x72, 0xeb, 0x91, 0x5d, 0x68, 0x5b, 0x9c,
	0x62, 0xe6, 0xe2, 0x99, 0xde, 0x78, 0xfb, 0x7a, 0x2d, 0xf4, 0xeb, 0x1d, 0xba, 0x66, 0x34, 0x42,
	0xe2, 0x16, 0xb5, 0xf3, 0x2d, 0x4a, 0x22, 0x79, 0x5b, 0xd6, 0xf7, 0xe2, 0x25, 0x6a, 0x77, 0x66,
	0xa8, 0x64, 0xd8, 0x1e, 0x30, 0xc2, 0x9f, 0x73, 0xb4, 0x8d, 0x06, 0xef, 0xca, 0x0d, 0x16, 0x4a,
	0xe8, 0x70, 0x08, 0x9b, 0xd5, 0x25, 0x5f, 0x5a, 0x3d, 0x79, 0xef, 0x99, 0x4e, 0x25, 0x7f, 0x8f,
	0x97, 0x67, 0x7b, 0xb0, 0x21, 0xb2, 0xab, 0xbf, 0x90, 0x0f, 0xe5, 0x1e, 0x22, 0xbb, 0x6a, 0xda,
	0x9f, 0xc2, 0xda, 0xd0, 0xcc, 0xb5, 0xac, 0xe7, 0x39, 0xf2, 0xc6, 0xff, 0xfd, 0x65, 0x18, 0xe5,
	0x05, 0xac, 0x13, 0x6a, 0xbc, 0x89, 0xe7, 0x3e, 0xf6, 0xbe, 0x07, 0xe1, 0x3e, 0x58, 0x77, 0xa1,
	0xad, 0xf4, 0xc8, 0xd0, 0x2c, 0xf6, 0x7e, 0x2c, 0x9b, 0xd7, 0x42, 0x30, 0xf7, 0xa0, 0x4b, 0x28,
	0xe6, 0x6e, 0x62, 0xa8, 0x99, 0x84, 0x4f, 0x61, 0xc1, 0x48, 0x0c, 0xd0, 0x11, 0x3c, 0x5a, 0x8e,
	0x82, 0x2c, 0x5e, 0x0e, 0xad, 0x31, 0x73, 0xc9, 0x67, 0x4f, 0x6e, 0x36, 0x02, 0x31, 0xa8, 0x65,
	0x76, 0x0c, 0x8f, 0x97, 0x61, 0x21, 0x25, 0xf1, 0x6c, 0x22, 0xf4, 0x18, 0x93, 0x93, 0x7f, 0xd0,
	0xa7, 0x52, 0x52, 0xdf, 0xcb, 0xec, 0x25, 0x30, 0x89, 0x53, 0x8c, 0x9f, 0x1f, 0x97, 0xc9, 0x69,
	0xb9, 0x5d, 0x50, 0xaa, 0x39, 0x07, 0x45, 0xcc, 0x17, 0x83, 0xe1, 0xe2, 0x1f, 0xfe, 0x52, 0xc6,
	0x3c, 0xd6, 0xc2, 0x6e, 0x3e, 0x22, 0xd6, 0x09, 0x6a, 0x44, 0xa4, 0x1f, 0x22, 0x12, 0x94, 0xca,
	0x3e, 0x5c, 0xf1, 0x1f, 0x87, 0xde, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0xc6, 0xf8, 0xd1,
	0x47, 0x04, 0x00, 0x00,
}
