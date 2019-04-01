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
// source: bgp_nbr_brief_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_sessions_session

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

type BgpNbrBriefBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpNbrBriefBag_KEYS) Reset()         { *m = BgpNbrBriefBag_KEYS{} }
func (m *BgpNbrBriefBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag_KEYS) ProtoMessage()    {}
func (*BgpNbrBriefBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{0}
}

func (m *BgpNbrBriefBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNbrBriefBag_KEYS.Unmarshal(m, b)
}
func (m *BgpNbrBriefBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNbrBriefBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpNbrBriefBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNbrBriefBag_KEYS.Merge(m, src)
}
func (m *BgpNbrBriefBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpNbrBriefBag_KEYS.Size(m)
}
func (m *BgpNbrBriefBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNbrBriefBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNbrBriefBag_KEYS proto.InternalMessageInfo

func (m *BgpNbrBriefBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNbrBriefBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type BgpL2VpnAddrT struct {
	L2VpnAddress         []uint32 `protobuf:"varint,1,rep,packed,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpL2VpnAddrT) Reset()         { *m = BgpL2VpnAddrT{} }
func (m *BgpL2VpnAddrT) String() string { return proto.CompactTextString(m) }
func (*BgpL2VpnAddrT) ProtoMessage()    {}
func (*BgpL2VpnAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{1}
}

func (m *BgpL2VpnAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpL2VpnAddrT.Unmarshal(m, b)
}
func (m *BgpL2VpnAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpL2VpnAddrT.Marshal(b, m, deterministic)
}
func (m *BgpL2VpnAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpL2VpnAddrT.Merge(m, src)
}
func (m *BgpL2VpnAddrT) XXX_Size() int {
	return xxx_messageInfo_BgpL2VpnAddrT.Size(m)
}
func (m *BgpL2VpnAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpL2VpnAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_BgpL2VpnAddrT proto.InternalMessageInfo

func (m *BgpL2VpnAddrT) GetL2VpnAddress() []uint32 {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

type BgpL2VpnMspwAddrT struct {
	L2VpnAddress         []uint32 `protobuf:"varint,1,rep,packed,name=l2vpn_address,json=l2vpnAddress,proto3" json:"l2vpn_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpL2VpnMspwAddrT) Reset()         { *m = BgpL2VpnMspwAddrT{} }
func (m *BgpL2VpnMspwAddrT) String() string { return proto.CompactTextString(m) }
func (*BgpL2VpnMspwAddrT) ProtoMessage()    {}
func (*BgpL2VpnMspwAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{2}
}

func (m *BgpL2VpnMspwAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpL2VpnMspwAddrT.Unmarshal(m, b)
}
func (m *BgpL2VpnMspwAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpL2VpnMspwAddrT.Marshal(b, m, deterministic)
}
func (m *BgpL2VpnMspwAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpL2VpnMspwAddrT.Merge(m, src)
}
func (m *BgpL2VpnMspwAddrT) XXX_Size() int {
	return xxx_messageInfo_BgpL2VpnMspwAddrT.Size(m)
}
func (m *BgpL2VpnMspwAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpL2VpnMspwAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_BgpL2VpnMspwAddrT proto.InternalMessageInfo

func (m *BgpL2VpnMspwAddrT) GetL2VpnAddress() []uint32 {
	if m != nil {
		return m.L2VpnAddress
	}
	return nil
}

type BgpIpv4SrpolicyAddrT struct {
	Ipv4SrpolicyAddress  []uint32 `protobuf:"varint,1,rep,packed,name=ipv4_srpolicy_address,json=ipv4SrpolicyAddress,proto3" json:"ipv4_srpolicy_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpIpv4SrpolicyAddrT) Reset()         { *m = BgpIpv4SrpolicyAddrT{} }
func (m *BgpIpv4SrpolicyAddrT) String() string { return proto.CompactTextString(m) }
func (*BgpIpv4SrpolicyAddrT) ProtoMessage()    {}
func (*BgpIpv4SrpolicyAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{3}
}

func (m *BgpIpv4SrpolicyAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpIpv4SrpolicyAddrT.Unmarshal(m, b)
}
func (m *BgpIpv4SrpolicyAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpIpv4SrpolicyAddrT.Marshal(b, m, deterministic)
}
func (m *BgpIpv4SrpolicyAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpIpv4SrpolicyAddrT.Merge(m, src)
}
func (m *BgpIpv4SrpolicyAddrT) XXX_Size() int {
	return xxx_messageInfo_BgpIpv4SrpolicyAddrT.Size(m)
}
func (m *BgpIpv4SrpolicyAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpIpv4SrpolicyAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_BgpIpv4SrpolicyAddrT proto.InternalMessageInfo

func (m *BgpIpv4SrpolicyAddrT) GetIpv4SrpolicyAddress() []uint32 {
	if m != nil {
		return m.Ipv4SrpolicyAddress
	}
	return nil
}

type BgpIpv6SrpolicyAddrT struct {
	Ipv6SrpolicyAddress  []uint32 `protobuf:"varint,1,rep,packed,name=ipv6_srpolicy_address,json=ipv6SrpolicyAddress,proto3" json:"ipv6_srpolicy_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpIpv6SrpolicyAddrT) Reset()         { *m = BgpIpv6SrpolicyAddrT{} }
func (m *BgpIpv6SrpolicyAddrT) String() string { return proto.CompactTextString(m) }
func (*BgpIpv6SrpolicyAddrT) ProtoMessage()    {}
func (*BgpIpv6SrpolicyAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{4}
}

func (m *BgpIpv6SrpolicyAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpIpv6SrpolicyAddrT.Unmarshal(m, b)
}
func (m *BgpIpv6SrpolicyAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpIpv6SrpolicyAddrT.Marshal(b, m, deterministic)
}
func (m *BgpIpv6SrpolicyAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpIpv6SrpolicyAddrT.Merge(m, src)
}
func (m *BgpIpv6SrpolicyAddrT) XXX_Size() int {
	return xxx_messageInfo_BgpIpv6SrpolicyAddrT.Size(m)
}
func (m *BgpIpv6SrpolicyAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpIpv6SrpolicyAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_BgpIpv6SrpolicyAddrT proto.InternalMessageInfo

func (m *BgpIpv6SrpolicyAddrT) GetIpv6SrpolicyAddress() []uint32 {
	if m != nil {
		return m.Ipv6SrpolicyAddress
	}
	return nil
}

type BgpAddrtype struct {
	Afi                    string                `protobuf:"bytes,1,opt,name=afi,proto3" json:"afi,omitempty"`
	Ipv4Address            string                `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv4McastAddress       string                `protobuf:"bytes,3,opt,name=ipv4_mcast_address,json=ipv4McastAddress,proto3" json:"ipv4_mcast_address,omitempty"`
	Ipv4LabelAddress       string                `protobuf:"bytes,4,opt,name=ipv4_label_address,json=ipv4LabelAddress,proto3" json:"ipv4_label_address,omitempty"`
	Ipv4TunnelAddress      string                `protobuf:"bytes,5,opt,name=ipv4_tunnel_address,json=ipv4TunnelAddress,proto3" json:"ipv4_tunnel_address,omitempty"`
	Ipv4MdtAddress         string                `protobuf:"bytes,6,opt,name=ipv4_mdt_address,json=ipv4MdtAddress,proto3" json:"ipv4_mdt_address,omitempty"`
	Ipv4VpnAddress         string                `protobuf:"bytes,7,opt,name=ipv4vpn_address,json=ipv4vpnAddress,proto3" json:"ipv4vpn_address,omitempty"`
	Ipv4VpnaMcastddress    string                `protobuf:"bytes,8,opt,name=ipv4vpna_mcastddress,json=ipv4vpnaMcastddress,proto3" json:"ipv4vpna_mcastddress,omitempty"`
	Ipv6Address            string                `protobuf:"bytes,9,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	Ipv6McastAddress       string                `protobuf:"bytes,10,opt,name=ipv6_mcast_address,json=ipv6McastAddress,proto3" json:"ipv6_mcast_address,omitempty"`
	Ipv6LabelAddress       string                `protobuf:"bytes,11,opt,name=ipv6_label_address,json=ipv6LabelAddress,proto3" json:"ipv6_label_address,omitempty"`
	Ipv6VpnAddress         string                `protobuf:"bytes,12,opt,name=ipv6vpn_address,json=ipv6vpnAddress,proto3" json:"ipv6vpn_address,omitempty"`
	Ipv6VpnMcastAddress    string                `protobuf:"bytes,13,opt,name=ipv6vpn_mcast_address,json=ipv6vpnMcastAddress,proto3" json:"ipv6vpn_mcast_address,omitempty"`
	L2VpnVplsAddress       *BgpL2VpnAddrT        `protobuf:"bytes,14,opt,name=l2vpn_vpls_address,json=l2vpnVplsAddress,proto3" json:"l2vpn_vpls_address,omitempty"`
	RtConstraintAddress    string                `protobuf:"bytes,15,opt,name=rt_constraint_address,json=rtConstraintAddress,proto3" json:"rt_constraint_address,omitempty"`
	Ipv6MvpnAddress        string                `protobuf:"bytes,16,opt,name=ipv6mvpn_address,json=ipv6mvpnAddress,proto3" json:"ipv6mvpn_address,omitempty"`
	Ipv4MvpnAddress        string                `protobuf:"bytes,17,opt,name=ipv4mvpn_address,json=ipv4mvpnAddress,proto3" json:"ipv4mvpn_address,omitempty"`
	L2VpnEvpnAddress       string                `protobuf:"bytes,18,opt,name=l2vpn_evpn_address,json=l2vpnEvpnAddress,proto3" json:"l2vpn_evpn_address,omitempty"`
	LsLsAddress            string                `protobuf:"bytes,19,opt,name=ls_ls_address,json=lsLsAddress,proto3" json:"ls_ls_address,omitempty"`
	L2VpnMspwAddress       *BgpL2VpnMspwAddrT    `protobuf:"bytes,20,opt,name=l2vpn_mspw_address,json=l2vpnMspwAddress,proto3" json:"l2vpn_mspw_address,omitempty"`
	Ipv4FlowspecAddress    string                `protobuf:"bytes,21,opt,name=ipv4_flowspec_address,json=ipv4FlowspecAddress,proto3" json:"ipv4_flowspec_address,omitempty"`
	Ipv6FlowspecAddress    string                `protobuf:"bytes,22,opt,name=ipv6_flowspec_address,json=ipv6FlowspecAddress,proto3" json:"ipv6_flowspec_address,omitempty"`
	Ipv4VpnFlowspecAddress string                `protobuf:"bytes,23,opt,name=ipv4vpn_flowspec_address,json=ipv4vpnFlowspecAddress,proto3" json:"ipv4vpn_flowspec_address,omitempty"`
	Ipv6VpnFlowspecAddress string                `protobuf:"bytes,24,opt,name=ipv6vpn_flowspec_address,json=ipv6vpnFlowspecAddress,proto3" json:"ipv6vpn_flowspec_address,omitempty"`
	Ipv4SrPolicyAddress    *BgpIpv4SrpolicyAddrT `protobuf:"bytes,25,opt,name=ipv4_sr_policy_address,json=ipv4SrPolicyAddress,proto3" json:"ipv4_sr_policy_address,omitempty"`
	Ipv6SrPolicyAddress    *BgpIpv6SrpolicyAddrT `protobuf:"bytes,26,opt,name=ipv6_sr_policy_address,json=ipv6SrPolicyAddress,proto3" json:"ipv6_sr_policy_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *BgpAddrtype) Reset()         { *m = BgpAddrtype{} }
func (m *BgpAddrtype) String() string { return proto.CompactTextString(m) }
func (*BgpAddrtype) ProtoMessage()    {}
func (*BgpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{5}
}

func (m *BgpAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpAddrtype.Unmarshal(m, b)
}
func (m *BgpAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpAddrtype.Marshal(b, m, deterministic)
}
func (m *BgpAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpAddrtype.Merge(m, src)
}
func (m *BgpAddrtype) XXX_Size() int {
	return xxx_messageInfo_BgpAddrtype.Size(m)
}
func (m *BgpAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_BgpAddrtype proto.InternalMessageInfo

func (m *BgpAddrtype) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4McastAddress() string {
	if m != nil {
		return m.Ipv4McastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4LabelAddress() string {
	if m != nil {
		return m.Ipv4LabelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4TunnelAddress() string {
	if m != nil {
		return m.Ipv4TunnelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4MdtAddress() string {
	if m != nil {
		return m.Ipv4MdtAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnAddress() string {
	if m != nil {
		return m.Ipv4VpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnaMcastddress() string {
	if m != nil {
		return m.Ipv4VpnaMcastddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6McastAddress() string {
	if m != nil {
		return m.Ipv6McastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6LabelAddress() string {
	if m != nil {
		return m.Ipv6LabelAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6VpnAddress() string {
	if m != nil {
		return m.Ipv6VpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6VpnMcastAddress() string {
	if m != nil {
		return m.Ipv6VpnMcastAddress
	}
	return ""
}

func (m *BgpAddrtype) GetL2VpnVplsAddress() *BgpL2VpnAddrT {
	if m != nil {
		return m.L2VpnVplsAddress
	}
	return nil
}

func (m *BgpAddrtype) GetRtConstraintAddress() string {
	if m != nil {
		return m.RtConstraintAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6MvpnAddress() string {
	if m != nil {
		return m.Ipv6MvpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4MvpnAddress() string {
	if m != nil {
		return m.Ipv4MvpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetL2VpnEvpnAddress() string {
	if m != nil {
		return m.L2VpnEvpnAddress
	}
	return ""
}

func (m *BgpAddrtype) GetLsLsAddress() string {
	if m != nil {
		return m.LsLsAddress
	}
	return ""
}

func (m *BgpAddrtype) GetL2VpnMspwAddress() *BgpL2VpnMspwAddrT {
	if m != nil {
		return m.L2VpnMspwAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv4FlowspecAddress() string {
	if m != nil {
		return m.Ipv4FlowspecAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6FlowspecAddress() string {
	if m != nil {
		return m.Ipv6FlowspecAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4VpnFlowspecAddress() string {
	if m != nil {
		return m.Ipv4VpnFlowspecAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv6VpnFlowspecAddress() string {
	if m != nil {
		return m.Ipv6VpnFlowspecAddress
	}
	return ""
}

func (m *BgpAddrtype) GetIpv4SrPolicyAddress() *BgpIpv4SrpolicyAddrT {
	if m != nil {
		return m.Ipv4SrPolicyAddress
	}
	return nil
}

func (m *BgpAddrtype) GetIpv6SrPolicyAddress() *BgpIpv6SrpolicyAddrT {
	if m != nil {
		return m.Ipv6SrPolicyAddress
	}
	return nil
}

type BgpNbrBriefBag struct {
	SpeakerId                uint32       `protobuf:"varint,50,opt,name=speaker_id,json=speakerId,proto3" json:"speaker_id,omitempty"`
	Description              string       `protobuf:"bytes,51,opt,name=description,proto3" json:"description,omitempty"`
	LocalAs                  uint32       `protobuf:"varint,52,opt,name=local_as,json=localAs,proto3" json:"local_as,omitempty"`
	RemoteAs                 uint32       `protobuf:"varint,53,opt,name=remote_as,json=remoteAs,proto3" json:"remote_as,omitempty"`
	MessagesQueuedIn         uint32       `protobuf:"varint,54,opt,name=messages_queued_in,json=messagesQueuedIn,proto3" json:"messages_queued_in,omitempty"`
	MessagesQueuedOut        uint32       `protobuf:"varint,55,opt,name=messages_queued_out,json=messagesQueuedOut,proto3" json:"messages_queued_out,omitempty"`
	ConnectionState          string       `protobuf:"bytes,56,opt,name=connection_state,json=connectionState,proto3" json:"connection_state,omitempty"`
	ConnectionLocalAddress   *BgpAddrtype `protobuf:"bytes,57,opt,name=connection_local_address,json=connectionLocalAddress,proto3" json:"connection_local_address,omitempty"`
	IsLocalAddressConfigured bool         `protobuf:"varint,58,opt,name=is_local_address_configured,json=isLocalAddressConfigured,proto3" json:"is_local_address_configured,omitempty"`
	ConnectionRemoteAddress  *BgpAddrtype `protobuf:"bytes,59,opt,name=connection_remote_address,json=connectionRemoteAddress,proto3" json:"connection_remote_address,omitempty"`
	VrfName                  string       `protobuf:"bytes,60,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	NsrEnabled               bool         `protobuf:"varint,61,opt,name=nsr_enabled,json=nsrEnabled,proto3" json:"nsr_enabled,omitempty"`
	NsrState                 string       `protobuf:"bytes,62,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	PostitPending            bool         `protobuf:"varint,63,opt,name=postit_pending,json=postitPending,proto3" json:"postit_pending,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}     `json:"-"`
	XXX_unrecognized         []byte       `json:"-"`
	XXX_sizecache            int32        `json:"-"`
}

func (m *BgpNbrBriefBag) Reset()         { *m = BgpNbrBriefBag{} }
func (m *BgpNbrBriefBag) String() string { return proto.CompactTextString(m) }
func (*BgpNbrBriefBag) ProtoMessage()    {}
func (*BgpNbrBriefBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_50c9d2128fae86fe, []int{6}
}

func (m *BgpNbrBriefBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNbrBriefBag.Unmarshal(m, b)
}
func (m *BgpNbrBriefBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNbrBriefBag.Marshal(b, m, deterministic)
}
func (m *BgpNbrBriefBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNbrBriefBag.Merge(m, src)
}
func (m *BgpNbrBriefBag) XXX_Size() int {
	return xxx_messageInfo_BgpNbrBriefBag.Size(m)
}
func (m *BgpNbrBriefBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNbrBriefBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNbrBriefBag proto.InternalMessageInfo

func (m *BgpNbrBriefBag) GetSpeakerId() uint32 {
	if m != nil {
		return m.SpeakerId
	}
	return 0
}

func (m *BgpNbrBriefBag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BgpNbrBriefBag) GetLocalAs() uint32 {
	if m != nil {
		return m.LocalAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetRemoteAs() uint32 {
	if m != nil {
		return m.RemoteAs
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedIn() uint32 {
	if m != nil {
		return m.MessagesQueuedIn
	}
	return 0
}

func (m *BgpNbrBriefBag) GetMessagesQueuedOut() uint32 {
	if m != nil {
		return m.MessagesQueuedOut
	}
	return 0
}

func (m *BgpNbrBriefBag) GetConnectionState() string {
	if m != nil {
		return m.ConnectionState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetConnectionLocalAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionLocalAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetIsLocalAddressConfigured() bool {
	if m != nil {
		return m.IsLocalAddressConfigured
	}
	return false
}

func (m *BgpNbrBriefBag) GetConnectionRemoteAddress() *BgpAddrtype {
	if m != nil {
		return m.ConnectionRemoteAddress
	}
	return nil
}

func (m *BgpNbrBriefBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNbrBriefBag) GetNsrEnabled() bool {
	if m != nil {
		return m.NsrEnabled
	}
	return false
}

func (m *BgpNbrBriefBag) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *BgpNbrBriefBag) GetPostitPending() bool {
	if m != nil {
		return m.PostitPending
	}
	return false
}

func init() {
	proto.RegisterType((*BgpNbrBriefBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_nbr_brief_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_addrtype")
	proto.RegisterType((*BgpNbrBriefBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.sessions.session.bgp_nbr_brief_bag")
}

func init() { proto.RegisterFile("bgp_nbr_brief_bag.proto", fileDescriptor_50c9d2128fae86fe) }

var fileDescriptor_50c9d2128fae86fe = []byte{
	// 968 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xd7, 0x11, 0x48, 0x9c, 0xb5, 0x9d, 0x38, 0x97, 0xc6, 0xb9, 0x50, 0x21, 0x4c, 0x10, 0xc2,
	0x95, 0x2a, 0x4b, 0xa4, 0xe6, 0x5a, 0xa0, 0x05, 0x45, 0x55, 0x90, 0x2a, 0x92, 0x52, 0x1c, 0x84,
	0xc4, 0x0b, 0xab, 0xf5, 0xdd, 0xda, 0x5d, 0x71, 0xde, 0x5b, 0x76, 0xd7, 0x4e, 0xfb, 0x2d, 0xf8,
	0x04, 0xbc, 0x20, 0xa1, 0xbe, 0xf0, 0x1d, 0xd1, 0xce, 0xde, 0xee, 0xfd, 0xb1, 0x41, 0xe2, 0x81,
	0xf0, 0x76, 0x37, 0xbf, 0xdf, 0xcc, 0xcd, 0xfc, 0x66, 0x76, 0x6e, 0xd1, 0xf1, 0x74, 0x2e, 0x30,
	0x9f, 0x4a, 0x3c, 0x95, 0x8c, 0xce, 0xf0, 0x94, 0xcc, 0x47, 0x42, 0xe6, 0x3a, 0x0f, 0x7f, 0x4a,
	0x98, 0x4a, 0x72, 0xcc, 0x72, 0x85, 0x5f, 0x49, 0xcc, 0xc4, 0x6a, 0x8c, 0x0d, 0x35, 0x17, 0x54,
	0x8e, 0xa6, 0x73, 0x31, 0x62, 0x5c, 0x69, 0xc2, 0x13, 0xaa, 0xfc, 0x93, 0x7f, 0xc0, 0x24, 0xd1,
	0x6c, 0x45, 0x47, 0x29, 0x9d, 0x91, 0x65, 0xa6, 0xf1, 0x4a, 0xce, 0x46, 0x8a, 0x2a, 0xc5, 0x72,
	0xae, 0xdc, 0xc3, 0xe9, 0x4b, 0xd4, 0x5f, 0xfb, 0x34, 0xfe, 0xe6, 0xe2, 0xc7, 0xeb, 0xf0, 0x43,
	0xd4, 0xf5, 0x91, 0x38, 0x59, 0xd0, 0x28, 0x18, 0x04, 0xc3, 0xdd, 0x49, 0xc7, 0x19, 0x9f, 0x93,
	0x05, 0x0d, 0xef, 0xa1, 0x1e, 0xa7, 0x6c, 0xfe, 0x72, 0x9a, 0x4b, 0x4c, 0xd2, 0x54, 0x52, 0xa5,
	0xa2, 0xb7, 0x80, 0xb7, 0xef, 0xec, 0xe7, 0xd6, 0x7c, 0xfa, 0x10, 0xf5, 0xcc, 0x97, 0xb2, 0xb3,
	0x95, 0xe0, 0xc0, 0xc5, 0xda, 0x7c, 0xa3, 0x7c, 0x37, 0xbe, 0xc1, 0x60, 0x6b, 0xd8, 0x9d, 0x74,
	0xc0, 0xe8, 0x1c, 0x1f, 0xa3, 0xa3, 0xd2, 0x71, 0xa1, 0xc4, 0xcd, 0xbf, 0xf2, 0x7e, 0x8e, 0x22,
	0xe3, 0x0d, 0xca, 0x29, 0x29, 0xf2, 0x8c, 0x25, 0xaf, 0x5d, 0x80, 0x33, 0x74, 0xb4, 0x6e, 0x2f,
	0x03, 0x1d, 0x1a, 0xf0, 0xba, 0xc0, 0xd6, 0xe3, 0xc5, 0x7f, 0x13, 0x2f, 0xfe, 0xa7, 0x78, 0x71,
	0x33, 0xde, 0xaf, 0x5d, 0xd4, 0x31, 0x01, 0x0d, 0x55, 0xbf, 0x16, 0x34, 0xec, 0xa1, 0x2d, 0x32,
	0x63, 0x85, 0xda, 0xe6, 0x31, 0xfc, 0x00, 0x75, 0x20, 0xcd, 0xba, 0xc0, 0x6d, 0x63, 0x2b, 0xa2,
	0x84, 0xf7, 0x51, 0x08, 0x94, 0x45, 0x42, 0x94, 0xf6, 0xc4, 0x2d, 0x20, 0xf6, 0x0c, 0x72, 0x65,
	0x80, 0x26, 0x3b, 0x23, 0x53, 0x9a, 0x79, 0xf6, 0xdb, 0x25, 0xfb, 0xd2, 0x00, 0x8e, 0x3d, 0x42,
	0x20, 0x04, 0xd6, 0x4b, 0xce, 0x2b, 0xf4, 0x77, 0x80, 0x7e, 0x60, 0xa0, 0xef, 0x01, 0x71, 0xfc,
	0x21, 0xea, 0xd9, 0x5c, 0xd2, 0x32, 0x93, 0x6d, 0x20, 0xef, 0x41, 0x26, 0xa9, 0xcf, 0xe3, 0x63,
	0xb4, 0x6f, 0x2c, 0xd5, 0x16, 0xee, 0x94, 0xc4, 0xb2, 0x89, 0xe1, 0x27, 0xe8, 0x4e, 0x61, 0x21,
	0xb6, 0xc4, 0x82, 0xdd, 0x02, 0xf6, 0xa1, 0xc3, 0xae, 0x4a, 0xa8, 0x10, 0x2d, 0xf6, 0x81, 0x77,
	0xbd, 0x68, 0x71, 0x5d, 0x86, 0xb8, 0x21, 0x1a, 0xf2, 0x32, 0xc4, 0x1b, 0x44, 0x8b, 0x1b, 0xa2,
	0xb5, 0x4b, 0x76, 0x4d, 0x34, 0x5b, 0x5a, 0x5c, 0x2d, 0xad, 0xe3, 0x4b, 0x8b, 0x2b, 0xa5, 0x15,
	0x33, 0x03, 0xb3, 0x5d, 0xcb, 0xa3, 0xeb, 0x6b, 0x33, 0x60, 0x2d, 0x95, 0xdf, 0x02, 0x14, 0xda,
	0xc9, 0x5f, 0x89, 0x4c, 0x79, 0x8f, 0xbd, 0x41, 0x30, 0x6c, 0x9f, 0x89, 0xd1, 0x7f, 0xbb, 0x32,
	0x46, 0xcd, 0x53, 0x3c, 0xe9, 0xc1, 0xdb, 0x0f, 0x22, 0x53, 0x95, 0xa2, 0xa4, 0xc6, 0x49, 0xce,
	0x95, 0x96, 0x84, 0xf1, 0xb2, 0xa8, 0x7d, 0x5b, 0x94, 0xd4, 0x4f, 0x3d, 0xe6, 0x7c, 0xee, 0xc1,
	0xd8, 0xc4, 0x8b, 0xaa, 0x64, 0x3d, 0xbb, 0x4a, 0x9c, 0xbd, 0x4e, 0x1d, 0xd7, 0xa8, 0x07, 0x9e,
	0x3a, 0xae, 0x52, 0xef, 0x3b, 0xa5, 0x68, 0x95, 0x1c, 0xda, 0xae, 0x01, 0x72, 0x51, 0x61, 0x9f,
	0xa2, 0x6e, 0xa6, 0x70, 0x45, 0xd2, 0x43, 0x3b, 0x35, 0x99, 0xba, 0xf4, 0xb5, 0xfd, 0xee, 0xc5,
	0xf7, 0xbb, 0xc8, 0x30, 0xef, 0x80, 0xf8, 0xcb, 0xdb, 0x13, 0xbf, 0xb2, 0x09, 0x8b, 0x4a, 0xae,
	0x94, 0xb8, 0xa9, 0x8f, 0xd5, 0x18, 0xcf, 0xb2, 0xfc, 0x46, 0x09, 0x9a, 0xf8, 0x3c, 0x8f, 0xca,
	0x23, 0xf3, 0x75, 0x81, 0x35, 0x46, 0x71, 0xdd, 0xa7, 0x5f, 0x8e, 0x62, 0xd3, 0xe7, 0x11, 0x8a,
	0xdc, 0x11, 0x5e, 0x73, 0x3b, 0x06, 0xb7, 0x7e, 0x81, 0x6f, 0xf6, 0x8c, 0x37, 0x7a, 0x46, 0xde,
	0x33, 0xde, 0xe0, 0xf9, 0x67, 0x80, 0xfa, 0xc5, 0xde, 0xc6, 0x8d, 0x45, 0x7b, 0x02, 0x5d, 0x78,
	0x75, 0x1b, 0x5d, 0xd8, 0xf4, 0x47, 0x71, 0xbf, 0x8c, 0x17, 0xd5, 0x15, 0xef, 0xf2, 0x8d, 0x37,
	0xe4, 0xfb, 0xee, 0xad, 0xe6, 0x1b, 0x6f, 0xcc, 0x37, 0x6e, 0xe4, 0x7b, 0xfa, 0x66, 0x1b, 0x1d,
	0xac, 0x5d, 0x0a, 0xc2, 0xf7, 0x10, 0x52, 0x82, 0x92, 0x9f, 0xa9, 0xc4, 0x2c, 0x8d, 0xce, 0x06,
	0xc1, 0xb0, 0x3b, 0xd9, 0x2d, 0x2c, 0xcf, 0xd2, 0x70, 0x80, 0xda, 0x29, 0x55, 0x89, 0x64, 0x42,
	0xb3, 0x9c, 0x47, 0x0f, 0xec, 0xc1, 0xa9, 0x98, 0xc2, 0x13, 0xd4, 0xca, 0xf2, 0x84, 0x64, 0x98,
	0xa8, 0x68, 0x0c, 0xee, 0x3b, 0xf0, 0x7e, 0xae, 0xc2, 0xbb, 0x68, 0x57, 0xd2, 0x45, 0xae, 0xa9,
	0xc1, 0x3e, 0x05, 0xac, 0x65, 0x0d, 0xe7, 0x70, 0x84, 0x17, 0x54, 0x29, 0x32, 0xa7, 0x0a, 0xff,
	0xb2, 0xa4, 0x4b, 0x9a, 0x62, 0xc6, 0xa3, 0x18, 0x58, 0x3d, 0x87, 0x7c, 0x07, 0xc0, 0x33, 0x6e,
	0xfe, 0x56, 0x4d, 0x76, 0xbe, 0xd4, 0xd1, 0x43, 0xa0, 0x1f, 0xd4, 0xe9, 0xdf, 0x2e, 0xb5, 0xd9,
	0x25, 0x49, 0xce, 0x39, 0x4d, 0x4c, 0x8e, 0x58, 0x69, 0xa2, 0x69, 0xf4, 0xc8, 0xee, 0x92, 0xd2,
	0x7e, 0x6d, 0xcc, 0xe1, 0x1f, 0x01, 0x8a, 0x2a, 0xdc, 0xa2, 0x98, 0xa2, 0x93, 0x9f, 0x41, 0x27,
	0xb3, 0xdb, 0xe8, 0xa4, 0xbb, 0x2a, 0x4c, 0xfa, 0x65, 0x36, 0x97, 0xa0, 0x64, 0x31, 0x70, 0x4f,
	0xd0, 0x5d, 0xa6, 0xea, 0xf9, 0x99, 0x65, 0x3c, 0x63, 0xf3, 0xa5, 0xa4, 0x69, 0xf4, 0xf9, 0x20,
	0x18, 0xb6, 0x26, 0x11, 0x53, 0x55, 0xa7, 0xa7, 0x1e, 0x0f, 0xdf, 0x04, 0xe8, 0xa4, 0x52, 0xa7,
	0xeb, 0x4c, 0x51, 0xe8, 0x17, 0xff, 0x43, 0xa1, 0xc7, 0x65, 0x3a, 0x13, 0x3b, 0x16, 0x45, 0xa5,
	0x27, 0xa8, 0xb5, 0x92, 0x33, 0x7b, 0x3f, 0x7d, 0x0c, 0x5d, 0xdb, 0x59, 0xc9, 0x19, 0x5c, 0x4d,
	0xdf, 0x47, 0x6d, 0xae, 0x24, 0xa6, 0x9c, 0x4c, 0x33, 0x9a, 0x46, 0x4f, 0xa0, 0x68, 0xc4, 0x95,
	0xbc, 0xb0, 0x16, 0x33, 0x74, 0x86, 0x60, 0x5b, 0xfe, 0x25, 0x38, 0xb7, 0xb8, 0x92, 0xb6, 0xd7,
	0x1f, 0xa1, 0x3d, 0x91, 0x2b, 0xcd, 0x34, 0x16, 0x94, 0xa7, 0x8c, 0xcf, 0xa3, 0xaf, 0x20, 0x40,
	0xd7, 0x5a, 0x5f, 0x58, 0xe3, 0x74, 0x1b, 0x6e, 0xe9, 0x0f, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x8c, 0xf5, 0x41, 0xae, 0xc0, 0x0b, 0x00, 0x00,
}