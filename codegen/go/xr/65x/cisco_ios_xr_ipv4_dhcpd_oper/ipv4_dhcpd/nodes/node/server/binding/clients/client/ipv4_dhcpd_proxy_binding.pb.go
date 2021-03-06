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
// source: ipv4_dhcpd_proxy_binding.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_server_binding_clients_client

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

type Ipv4DhcpdProxyBinding_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyBinding_KEYS) Reset()         { *m = Ipv4DhcpdProxyBinding_KEYS{} }
func (m *Ipv4DhcpdProxyBinding_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyBinding_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdProxyBinding_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_eef22acb6c976c44, []int{0}
}

func (m *Ipv4DhcpdProxyBinding_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyBinding_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyBinding_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdProxyBinding_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS.Size(m)
}
func (m *Ipv4DhcpdProxyBinding_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyBinding_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyBinding_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding_KEYS) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type Ipv4DhcpdProxyBinding struct {
	ClientIdXr              string   `protobuf:"bytes,50,opt,name=client_id_xr,json=clientIdXr,proto3" json:"client_id_xr,omitempty"`
	MacAddress              string   `protobuf:"bytes,51,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	VrfName                 string   `protobuf:"bytes,52,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ServerVrfName           string   `protobuf:"bytes,53,opt,name=server_vrf_name,json=serverVrfName,proto3" json:"server_vrf_name,omitempty"`
	IpAddress               string   `protobuf:"bytes,54,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	ClientGiAddr            string   `protobuf:"bytes,55,opt,name=client_gi_addr,json=clientGiAddr,proto3" json:"client_gi_addr,omitempty"`
	ToServerGiAddr          string   `protobuf:"bytes,56,opt,name=to_server_gi_addr,json=toServerGiAddr,proto3" json:"to_server_gi_addr,omitempty"`
	ServerIpAddress         string   `protobuf:"bytes,57,opt,name=server_ip_address,json=serverIpAddress,proto3" json:"server_ip_address,omitempty"`
	ReplyServerIpAddress    string   `protobuf:"bytes,58,opt,name=reply_server_ip_address,json=replyServerIpAddress,proto3" json:"reply_server_ip_address,omitempty"`
	LeaseTime               uint32   `protobuf:"varint,59,opt,name=lease_time,json=leaseTime,proto3" json:"lease_time,omitempty"`
	RemainingLeaseTime      uint32   `protobuf:"varint,60,opt,name=remaining_lease_time,json=remainingLeaseTime,proto3" json:"remaining_lease_time,omitempty"`
	State                   string   `protobuf:"bytes,61,opt,name=state,proto3" json:"state,omitempty"`
	InterfaceName           string   `protobuf:"bytes,62,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	AccessVrfName           string   `protobuf:"bytes,63,opt,name=access_vrf_name,json=accessVrfName,proto3" json:"access_vrf_name,omitempty"`
	ProxyBindingOuterTag    uint32   `protobuf:"varint,64,opt,name=proxy_binding_outer_tag,json=proxyBindingOuterTag,proto3" json:"proxy_binding_outer_tag,omitempty"`
	ProxyBindingInnerTag    uint32   `protobuf:"varint,65,opt,name=proxy_binding_inner_tag,json=proxyBindingInnerTag,proto3" json:"proxy_binding_inner_tag,omitempty"`
	ProfileName             string   `protobuf:"bytes,66,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	IsNakNextRenew          bool     `protobuf:"varint,67,opt,name=is_nak_next_renew,json=isNakNextRenew,proto3" json:"is_nak_next_renew,omitempty"`
	SubscriberLabel         uint32   `protobuf:"varint,68,opt,name=subscriber_label,json=subscriberLabel,proto3" json:"subscriber_label,omitempty"`
	OldSubscriberLabel      uint32   `protobuf:"varint,69,opt,name=old_subscriber_label,json=oldSubscriberLabel,proto3" json:"old_subscriber_label,omitempty"`
	SubscriberInterfaceName string   `protobuf:"bytes,70,opt,name=subscriber_interface_name,json=subscriberInterfaceName,proto3" json:"subscriber_interface_name,omitempty"`
	RxCircuitId             string   `protobuf:"bytes,71,opt,name=rx_circuit_id,json=rxCircuitId,proto3" json:"rx_circuit_id,omitempty"`
	TxCircuitId             string   `protobuf:"bytes,72,opt,name=tx_circuit_id,json=txCircuitId,proto3" json:"tx_circuit_id,omitempty"`
	RxRemoteId              string   `protobuf:"bytes,73,opt,name=rx_remote_id,json=rxRemoteId,proto3" json:"rx_remote_id,omitempty"`
	TxRemoteId              string   `protobuf:"bytes,74,opt,name=tx_remote_id,json=txRemoteId,proto3" json:"tx_remote_id,omitempty"`
	RxVsiso                 string   `protobuf:"bytes,75,opt,name=rx_vsiso,json=rxVsiso,proto3" json:"rx_vsiso,omitempty"`
	TxVsiso                 string   `protobuf:"bytes,76,opt,name=tx_vsiso,json=txVsiso,proto3" json:"tx_vsiso,omitempty"`
	IsAuthReceived          bool     `protobuf:"varint,77,opt,name=is_auth_received,json=isAuthReceived,proto3" json:"is_auth_received,omitempty"`
	IsMblSubscriber         bool     `protobuf:"varint,78,opt,name=is_mbl_subscriber,json=isMblSubscriber,proto3" json:"is_mbl_subscriber,omitempty"`
	ParamRequest            string   `protobuf:"bytes,79,opt,name=param_request,json=paramRequest,proto3" json:"param_request,omitempty"`
	ParamResponse           string   `protobuf:"bytes,80,opt,name=param_response,json=paramResponse,proto3" json:"param_response,omitempty"`
	SessionStartTimeEpoch   uint64   `protobuf:"varint,81,opt,name=session_start_time_epoch,json=sessionStartTimeEpoch,proto3" json:"session_start_time_epoch,omitempty"`
	SrgState                uint32   `protobuf:"varint,82,opt,name=srg_state,json=srgState,proto3" json:"srg_state,omitempty"`
	EventHistory            []uint32 `protobuf:"varint,83,rep,packed,name=event_history,json=eventHistory,proto3" json:"event_history,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ipv4DhcpdProxyBinding) Reset()         { *m = Ipv4DhcpdProxyBinding{} }
func (m *Ipv4DhcpdProxyBinding) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdProxyBinding) ProtoMessage()    {}
func (*Ipv4DhcpdProxyBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_eef22acb6c976c44, []int{1}
}

func (m *Ipv4DhcpdProxyBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding.Unmarshal(m, b)
}
func (m *Ipv4DhcpdProxyBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdProxyBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdProxyBinding.Merge(m, src)
}
func (m *Ipv4DhcpdProxyBinding) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdProxyBinding.Size(m)
}
func (m *Ipv4DhcpdProxyBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdProxyBinding.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdProxyBinding proto.InternalMessageInfo

func (m *Ipv4DhcpdProxyBinding) GetClientIdXr() string {
	if m != nil {
		return m.ClientIdXr
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetServerVrfName() string {
	if m != nil {
		return m.ServerVrfName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetClientGiAddr() string {
	if m != nil {
		return m.ClientGiAddr
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetToServerGiAddr() string {
	if m != nil {
		return m.ToServerGiAddr
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetServerIpAddress() string {
	if m != nil {
		return m.ServerIpAddress
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetReplyServerIpAddress() string {
	if m != nil {
		return m.ReplyServerIpAddress
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetLeaseTime() uint32 {
	if m != nil {
		return m.LeaseTime
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetRemainingLeaseTime() uint32 {
	if m != nil {
		return m.RemainingLeaseTime
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetAccessVrfName() string {
	if m != nil {
		return m.AccessVrfName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetProxyBindingOuterTag() uint32 {
	if m != nil {
		return m.ProxyBindingOuterTag
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetProxyBindingInnerTag() uint32 {
	if m != nil {
		return m.ProxyBindingInnerTag
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetIsNakNextRenew() bool {
	if m != nil {
		return m.IsNakNextRenew
	}
	return false
}

func (m *Ipv4DhcpdProxyBinding) GetSubscriberLabel() uint32 {
	if m != nil {
		return m.SubscriberLabel
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetOldSubscriberLabel() uint32 {
	if m != nil {
		return m.OldSubscriberLabel
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetSubscriberInterfaceName() string {
	if m != nil {
		return m.SubscriberInterfaceName
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetRxCircuitId() string {
	if m != nil {
		return m.RxCircuitId
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetTxCircuitId() string {
	if m != nil {
		return m.TxCircuitId
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetRxRemoteId() string {
	if m != nil {
		return m.RxRemoteId
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetTxRemoteId() string {
	if m != nil {
		return m.TxRemoteId
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetRxVsiso() string {
	if m != nil {
		return m.RxVsiso
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetTxVsiso() string {
	if m != nil {
		return m.TxVsiso
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetIsAuthReceived() bool {
	if m != nil {
		return m.IsAuthReceived
	}
	return false
}

func (m *Ipv4DhcpdProxyBinding) GetIsMblSubscriber() bool {
	if m != nil {
		return m.IsMblSubscriber
	}
	return false
}

func (m *Ipv4DhcpdProxyBinding) GetParamRequest() string {
	if m != nil {
		return m.ParamRequest
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetParamResponse() string {
	if m != nil {
		return m.ParamResponse
	}
	return ""
}

func (m *Ipv4DhcpdProxyBinding) GetSessionStartTimeEpoch() uint64 {
	if m != nil {
		return m.SessionStartTimeEpoch
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetSrgState() uint32 {
	if m != nil {
		return m.SrgState
	}
	return 0
}

func (m *Ipv4DhcpdProxyBinding) GetEventHistory() []uint32 {
	if m != nil {
		return m.EventHistory
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4DhcpdProxyBinding_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.binding.clients.client.ipv4_dhcpd_proxy_binding_KEYS")
	proto.RegisterType((*Ipv4DhcpdProxyBinding)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.server.binding.clients.client.ipv4_dhcpd_proxy_binding")
}

func init() { proto.RegisterFile("ipv4_dhcpd_proxy_binding.proto", fileDescriptor_eef22acb6c976c44) }

var fileDescriptor_eef22acb6c976c44 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x95, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0xc7, 0x65, 0x1e, 0x8a, 0x3d, 0xb5, 0x9d, 0xe6, 0x64, 0xc8, 0x46, 0x28, 0x60, 0x52, 0x40,
	0x0e, 0x2f, 0x2c, 0x44, 0x1b, 0x0a, 0xe5, 0x31, 0x2d, 0xa1, 0x35, 0x4d, 0xdd, 0x72, 0x8e, 0x2a,
	0x78, 0xb5, 0x5a, 0xdf, 0x8d, 0xed, 0x55, 0xef, 0x6e, 0x8f, 0xd9, 0xb5, 0xb9, 0x7c, 0x58, 0xbe,
	0x0b, 0xda, 0x87, 0xb3, 0x9d, 0x58, 0x79, 0x93, 0x68, 0xff, 0xff, 0xdf, 0xec, 0xcc, 0xcd, 0xdc,
	0x9c, 0xe1, 0x13, 0x59, 0xae, 0x1e, 0xf2, 0x74, 0x91, 0x94, 0x29, 0x2f, 0x49, 0x55, 0x57, 0x7c,
	0x2a, 0x8b, 0x54, 0x16, 0xf3, 0x61, 0x49, 0xca, 0xa8, 0xe8, 0x75, 0x22, 0x75, 0xa2, 0xb8, 0x54,
	0x9a, 0x57, 0xc4, 0xb7, 0x60, 0x55, 0x22, 0x0d, 0x37, 0xe7, 0x61, 0xa1, 0x52, 0xd4, 0xee, 0xef,
	0x50, 0x23, 0xad, 0x90, 0x86, 0xf5, 0x45, 0x49, 0x26, 0xb1, 0x30, 0x3a, 0xfc, 0x3f, 0xbe, 0x84,
	0xa3, 0xdb, 0x72, 0xf2, 0x17, 0xe7, 0x7f, 0x4f, 0xa2, 0x8f, 0xe0, 0x8e, 0xbd, 0x46, 0xa6, 0xac,
	0xd1, 0x6f, 0x0c, 0x5a, 0x71, 0x38, 0x45, 0x1f, 0x43, 0xcb, 0x5f, 0xc1, 0x65, 0xca, 0xde, 0x71,
	0x56, 0xd3, 0x0b, 0xa3, 0xf4, 0xf8, 0x3f, 0x00, 0x76, 0xdb, 0xb5, 0x51, 0x1f, 0xda, 0xeb, 0x48,
	0x5e, 0x11, 0xfb, 0xc6, 0x05, 0x43, 0x1d, 0xfc, 0x17, 0x45, 0x9f, 0xc2, 0xdd, 0x5c, 0x24, 0x5c,
	0xa4, 0x29, 0xa1, 0xd6, 0xec, 0x81, 0x07, 0x72, 0x91, 0x9c, 0x79, 0x25, 0x3a, 0x84, 0xe6, 0x8a,
	0x66, 0xbc, 0x10, 0x39, 0xb2, 0x87, 0xce, 0xfd, 0x60, 0x45, 0xb3, 0xb1, 0xc8, 0x31, 0xfa, 0x12,
	0xf6, 0xfc, 0x13, 0xf3, 0x35, 0x71, 0xea, 0x88, 0x8e, 0x97, 0xdf, 0x04, 0xee, 0x08, 0x40, 0x96,
	0xeb, 0x14, 0xdf, 0x3a, 0xa4, 0x25, 0xcb, 0x3a, 0xc3, 0xe7, 0xd0, 0x0d, 0x45, 0xce, 0xa5, 0xa3,
	0xd8, 0x23, 0x87, 0x84, 0xd2, 0x9f, 0x49, 0x0b, 0x46, 0x27, 0xb0, 0x6f, 0x14, 0x0f, 0xf9, 0x6a,
	0xf0, 0x3b, 0x07, 0x76, 0x8d, 0x9a, 0x38, 0x3d, 0xa0, 0x5f, 0xc1, 0x7e, 0xe0, 0xb6, 0xd2, 0x7e,
	0xef, 0xd0, 0x50, 0xf0, 0x68, 0x9d, 0xfc, 0x14, 0x0e, 0x08, 0xcb, 0xec, 0x8a, 0xef, 0x46, 0x3c,
	0x76, 0x11, 0x3d, 0x67, 0x4f, 0x6e, 0x84, 0x1d, 0x01, 0x64, 0x28, 0x34, 0x72, 0x23, 0x73, 0x64,
	0x3f, 0xf4, 0x1b, 0x83, 0x4e, 0xdc, 0x72, 0xca, 0xa5, 0xcc, 0x31, 0xfa, 0x1a, 0x7a, 0x84, 0xb9,
	0x90, 0x85, 0x9d, 0xed, 0x16, 0xf8, 0xa3, 0x03, 0xa3, 0xb5, 0x77, 0xb1, 0x8e, 0xe8, 0xc1, 0xfb,
	0xda, 0x08, 0x83, 0xec, 0x27, 0x97, 0xd5, 0x1f, 0xa2, 0x2f, 0xa0, 0x2b, 0x0b, 0x83, 0x34, 0x13,
	0x09, 0xfa, 0x06, 0xff, 0xec, 0x1b, 0xbc, 0x56, 0xeb, 0x41, 0x88, 0x24, 0x41, 0xad, 0x37, 0x83,
	0xf8, 0xc5, 0x73, 0x5e, 0xae, 0x07, 0x71, 0x0a, 0x07, 0xd7, 0x5f, 0x3b, 0xb5, 0x34, 0x48, 0xdc,
	0x88, 0x39, 0xfb, 0xd5, 0x55, 0xd6, 0x73, 0xf6, 0x13, 0xef, 0xbe, 0xb2, 0xe6, 0xa5, 0x98, 0xef,
	0x86, 0xc9, 0xa2, 0x08, 0x61, 0x67, 0xbb, 0x61, 0x23, 0x6b, 0xda, 0xb0, 0xcf, 0xa0, 0x5d, 0x92,
	0x9a, 0xc9, 0x2c, 0x94, 0xfe, 0xc4, 0x95, 0x74, 0x37, 0x68, 0xae, 0xa0, 0x13, 0xd8, 0x97, 0x9a,
	0x17, 0xe2, 0x2d, 0x2f, 0xb0, 0x32, 0x9c, 0xb0, 0xc0, 0x7f, 0xd9, 0xd3, 0x7e, 0x63, 0xd0, 0x8c,
	0xbb, 0x52, 0x8f, 0xc5, 0xdb, 0x31, 0x56, 0x26, 0xb6, 0x6a, 0x74, 0x02, 0xf7, 0xf4, 0x72, 0xaa,
	0x13, 0x92, 0x53, 0x24, 0x9e, 0x89, 0x29, 0x66, 0xec, 0x37, 0x97, 0x7d, 0x6f, 0xa3, 0x5f, 0x58,
	0xd9, 0x76, 0x5f, 0x65, 0x29, 0xdf, 0xc1, 0xcf, 0x7d, 0xf7, 0x55, 0x96, 0x4e, 0x6e, 0x44, 0x3c,
	0x86, 0xc3, 0x2d, 0xfa, 0x46, 0xcb, 0x7f, 0x77, 0x75, 0x1f, 0x6c, 0x80, 0xd1, 0xb5, 0xe6, 0x1f,
	0x43, 0x87, 0x2a, 0x9e, 0x48, 0x4a, 0x96, 0xd2, 0x6d, 0xe8, 0x33, 0xff, 0x9c, 0x54, 0x3d, 0xf5,
	0xda, 0x28, 0xb5, 0x8c, 0xb9, 0xc6, 0x3c, 0xf7, 0x8c, 0xd9, 0x62, 0xfa, 0xd0, 0xa6, 0x8a, 0x13,
	0xe6, 0xca, 0xa0, 0x45, 0x46, 0x7e, 0x15, 0xa9, 0x8a, 0x9d, 0xe4, 0x09, 0xb3, 0x4d, 0xfc, 0xe1,
	0x09, 0xb3, 0x21, 0x0e, 0xa1, 0x49, 0x15, 0x5f, 0x69, 0xa9, 0x15, 0x7b, 0xe1, 0x97, 0x95, 0xaa,
	0x37, 0xf6, 0x68, 0x2d, 0x53, 0x5b, 0x17, 0xde, 0x32, 0xc1, 0x1a, 0xc0, 0x3d, 0xa9, 0xb9, 0x58,
	0x9a, 0x05, 0x27, 0x4c, 0x50, 0xae, 0x30, 0x65, 0x2f, 0xeb, 0x21, 0x9c, 0x2d, 0xcd, 0x22, 0x0e,
	0xaa, 0xdd, 0x2c, 0xa9, 0x79, 0x3e, 0xcd, 0xb6, 0x9a, 0xcb, 0xc6, 0x0e, 0xdd, 0x93, 0xfa, 0xe5,
	0x34, 0xdb, 0x34, 0x36, 0xba, 0x0f, 0x9d, 0x52, 0x90, 0xc8, 0x39, 0xe1, 0x3f, 0x4b, 0xd4, 0x86,
	0xbd, 0xf2, 0x5b, 0xed, 0xc4, 0xd8, 0x6b, 0xf6, 0x05, 0xaf, 0x21, 0x5d, 0xaa, 0x42, 0x23, 0x7b,
	0xed, 0x5f, 0xdc, 0x40, 0x79, 0x31, 0x7a, 0x04, 0x4c, 0xa3, 0xd6, 0x52, 0x15, 0x5c, 0x1b, 0x41,
	0xc6, 0x6d, 0x13, 0xc7, 0x52, 0x25, 0x0b, 0xf6, 0x67, 0xbf, 0x31, 0x78, 0x2f, 0xfe, 0x30, 0xf8,
	0x13, 0x6b, 0xdb, 0x8d, 0x3a, 0xb7, 0xa6, 0xfd, 0x74, 0x6a, 0x9a, 0x73, 0xbf, 0x5a, 0xb1, 0x9b,
	0x7f, 0x53, 0xd3, 0x7c, 0xe2, 0xb6, 0xeb, 0x3e, 0x74, 0x70, 0x65, 0xbf, 0x3b, 0x0b, 0xa9, 0x8d,
	0xa2, 0x2b, 0x36, 0xe9, 0xbf, 0x3b, 0xe8, 0xc4, 0x6d, 0x27, 0x3e, 0xf7, 0xda, 0xf4, 0x8e, 0xfb,
	0x39, 0x78, 0xf0, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf1, 0xf1, 0x30, 0x30, 0x06, 0x00,
	0x00,
}
