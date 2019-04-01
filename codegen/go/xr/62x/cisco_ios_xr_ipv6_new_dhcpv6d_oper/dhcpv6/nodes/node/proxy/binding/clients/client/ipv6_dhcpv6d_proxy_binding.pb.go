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
// source: ipv6_dhcpv6d_proxy_binding.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_binding_clients_client

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

type Ipv6Dhcpv6DProxyBinding_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) Reset()         { *m = Ipv6Dhcpv6DProxyBinding_KEYS{} }
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyBinding_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyBinding_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{0}
}

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DProxyBinding_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyBinding_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding_KEYS) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type BagDhcpv6DAddrAttrbItem struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	LeaseTime            uint32   `protobuf:"varint,3,opt,name=lease_time,json=leaseTime,proto3" json:"lease_time,omitempty"`
	RemainingLeaseTime   uint32   `protobuf:"varint,4,opt,name=remaining_lease_time,json=remainingLeaseTime,proto3" json:"remaining_lease_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BagDhcpv6DAddrAttrbItem) Reset()         { *m = BagDhcpv6DAddrAttrbItem{} }
func (m *BagDhcpv6DAddrAttrbItem) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DAddrAttrbItem) ProtoMessage()    {}
func (*BagDhcpv6DAddrAttrbItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{1}
}

func (m *BagDhcpv6DAddrAttrbItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbItem.Unmarshal(m, b)
}
func (m *BagDhcpv6DAddrAttrbItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbItem.Marshal(b, m, deterministic)
}
func (m *BagDhcpv6DAddrAttrbItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagDhcpv6DAddrAttrbItem.Merge(m, src)
}
func (m *BagDhcpv6DAddrAttrbItem) XXX_Size() int {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbItem.Size(m)
}
func (m *BagDhcpv6DAddrAttrbItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BagDhcpv6DAddrAttrbItem.DiscardUnknown(m)
}

var xxx_messageInfo_BagDhcpv6DAddrAttrbItem proto.InternalMessageInfo

func (m *BagDhcpv6DAddrAttrbItem) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *BagDhcpv6DAddrAttrbItem) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *BagDhcpv6DAddrAttrbItem) GetLeaseTime() uint32 {
	if m != nil {
		return m.LeaseTime
	}
	return 0
}

func (m *BagDhcpv6DAddrAttrbItem) GetRemainingLeaseTime() uint32 {
	if m != nil {
		return m.RemainingLeaseTime
	}
	return 0
}

type BagDhcpv6DAddrAttrbEntry struct {
	BagDhcpv6DAddrAttrb  []*BagDhcpv6DAddrAttrbItem `protobuf:"bytes,5,rep,name=bag_dhcpv6d_addr_attrb,json=bagDhcpv6dAddrAttrb,proto3" json:"bag_dhcpv6d_addr_attrb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *BagDhcpv6DAddrAttrbEntry) Reset()         { *m = BagDhcpv6DAddrAttrbEntry{} }
func (m *BagDhcpv6DAddrAttrbEntry) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DAddrAttrbEntry) ProtoMessage()    {}
func (*BagDhcpv6DAddrAttrbEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{2}
}

func (m *BagDhcpv6DAddrAttrbEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbEntry.Unmarshal(m, b)
}
func (m *BagDhcpv6DAddrAttrbEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbEntry.Marshal(b, m, deterministic)
}
func (m *BagDhcpv6DAddrAttrbEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagDhcpv6DAddrAttrbEntry.Merge(m, src)
}
func (m *BagDhcpv6DAddrAttrbEntry) XXX_Size() int {
	return xxx_messageInfo_BagDhcpv6DAddrAttrbEntry.Size(m)
}
func (m *BagDhcpv6DAddrAttrbEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BagDhcpv6DAddrAttrbEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BagDhcpv6DAddrAttrbEntry proto.InternalMessageInfo

func (m *BagDhcpv6DAddrAttrbEntry) GetBagDhcpv6DAddrAttrb() []*BagDhcpv6DAddrAttrbItem {
	if m != nil {
		return m.BagDhcpv6DAddrAttrb
	}
	return nil
}

type BagDhcpv6DIaIdPdInfoItem struct {
	IaType               string                    `protobuf:"bytes,1,opt,name=ia_type,json=iaType,proto3" json:"ia_type,omitempty"`
	IaId                 uint32                    `protobuf:"varint,2,opt,name=ia_id,json=iaId,proto3" json:"ia_id,omitempty"`
	Flags                uint32                    `protobuf:"varint,3,opt,name=flags,proto3" json:"flags,omitempty"`
	TotalAddress         uint32                    `protobuf:"varint,4,opt,name=total_address,json=totalAddress,proto3" json:"total_address,omitempty"`
	State                string                    `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	Addresses            *BagDhcpv6DAddrAttrbEntry `protobuf:"bytes,6,opt,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *BagDhcpv6DIaIdPdInfoItem) Reset()         { *m = BagDhcpv6DIaIdPdInfoItem{} }
func (m *BagDhcpv6DIaIdPdInfoItem) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DIaIdPdInfoItem) ProtoMessage()    {}
func (*BagDhcpv6DIaIdPdInfoItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{3}
}

func (m *BagDhcpv6DIaIdPdInfoItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem.Unmarshal(m, b)
}
func (m *BagDhcpv6DIaIdPdInfoItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem.Marshal(b, m, deterministic)
}
func (m *BagDhcpv6DIaIdPdInfoItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem.Merge(m, src)
}
func (m *BagDhcpv6DIaIdPdInfoItem) XXX_Size() int {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem.Size(m)
}
func (m *BagDhcpv6DIaIdPdInfoItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem.DiscardUnknown(m)
}

var xxx_messageInfo_BagDhcpv6DIaIdPdInfoItem proto.InternalMessageInfo

func (m *BagDhcpv6DIaIdPdInfoItem) GetIaType() string {
	if m != nil {
		return m.IaType
	}
	return ""
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetIaId() uint32 {
	if m != nil {
		return m.IaId
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetTotalAddress() uint32 {
	if m != nil {
		return m.TotalAddress
	}
	return 0
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BagDhcpv6DIaIdPdInfoItem) GetAddresses() *BagDhcpv6DAddrAttrbEntry {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type BagDhcpv6DIaIdPdInfoEntry struct {
	BagDhcpv6DIaIdPdInfo []*BagDhcpv6DIaIdPdInfoItem `protobuf:"bytes,7,rep,name=bag_dhcpv6d_ia_id_pd_info,json=bagDhcpv6dIaIdPdInfo,proto3" json:"bag_dhcpv6d_ia_id_pd_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BagDhcpv6DIaIdPdInfoEntry) Reset()         { *m = BagDhcpv6DIaIdPdInfoEntry{} }
func (m *BagDhcpv6DIaIdPdInfoEntry) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DIaIdPdInfoEntry) ProtoMessage()    {}
func (*BagDhcpv6DIaIdPdInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{4}
}

func (m *BagDhcpv6DIaIdPdInfoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry.Unmarshal(m, b)
}
func (m *BagDhcpv6DIaIdPdInfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry.Marshal(b, m, deterministic)
}
func (m *BagDhcpv6DIaIdPdInfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry.Merge(m, src)
}
func (m *BagDhcpv6DIaIdPdInfoEntry) XXX_Size() int {
	return xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry.Size(m)
}
func (m *BagDhcpv6DIaIdPdInfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_BagDhcpv6DIaIdPdInfoEntry proto.InternalMessageInfo

func (m *BagDhcpv6DIaIdPdInfoEntry) GetBagDhcpv6DIaIdPdInfo() []*BagDhcpv6DIaIdPdInfoItem {
	if m != nil {
		return m.BagDhcpv6DIaIdPdInfo
	}
	return nil
}

type Ipv6Dhcpv6DProxyBinding struct {
	Duid                 string                     `protobuf:"bytes,50,opt,name=duid,proto3" json:"duid,omitempty"`
	ClientFlag           uint32                     `protobuf:"varint,51,opt,name=client_flag,json=clientFlag,proto3" json:"client_flag,omitempty"`
	SubscriberLabel      uint32                     `protobuf:"varint,52,opt,name=subscriber_label,json=subscriberLabel,proto3" json:"subscriber_label,omitempty"`
	VrfName              string                     `protobuf:"bytes,53,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	MacAddress           string                     `protobuf:"bytes,54,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	IaIdPDs              uint32                     `protobuf:"varint,55,opt,name=ia_id_p_ds,json=iaIdPDs,proto3" json:"ia_id_p_ds,omitempty"`
	IaIdPd               *BagDhcpv6DIaIdPdInfoEntry `protobuf:"bytes,56,opt,name=ia_id_pd,json=iaIdPd,proto3" json:"ia_id_pd,omitempty"`
	InterfaceName        string                     `protobuf:"bytes,57,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	AccessVrfName        string                     `protobuf:"bytes,58,opt,name=access_vrf_name,json=accessVrfName,proto3" json:"access_vrf_name,omitempty"`
	ProxyBindingTags     uint32                     `protobuf:"varint,59,opt,name=proxy_binding_tags,json=proxyBindingTags,proto3" json:"proxy_binding_tags,omitempty"`
	ProxyBindingOuterTag uint32                     `protobuf:"varint,60,opt,name=proxy_binding_outer_tag,json=proxyBindingOuterTag,proto3" json:"proxy_binding_outer_tag,omitempty"`
	ProxyBindingInnerTag uint32                     `protobuf:"varint,61,opt,name=proxy_binding_inner_tag,json=proxyBindingInnerTag,proto3" json:"proxy_binding_inner_tag,omitempty"`
	ClassName            string                     `protobuf:"bytes,62,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	PoolName             string                     `protobuf:"bytes,63,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	RxRemoteId           string                     `protobuf:"bytes,64,opt,name=rx_remote_id,json=rxRemoteId,proto3" json:"rx_remote_id,omitempty"`
	TxRemoteId           string                     `protobuf:"bytes,65,opt,name=tx_remote_id,json=txRemoteId,proto3" json:"tx_remote_id,omitempty"`
	RxInterfaceId        string                     `protobuf:"bytes,66,opt,name=rx_interface_id,json=rxInterfaceId,proto3" json:"rx_interface_id,omitempty"`
	TxInterfaceId        string                     `protobuf:"bytes,67,opt,name=tx_interface_id,json=txInterfaceId,proto3" json:"tx_interface_id,omitempty"`
	ServerIpv6Address    string                     `protobuf:"bytes,68,opt,name=server_ipv6_address,json=serverIpv6Address,proto3" json:"server_ipv6_address,omitempty"`
	ProfileName          string                     `protobuf:"bytes,69,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	FramedIpv6Prefix     string                     `protobuf:"bytes,70,opt,name=framed_ipv6_prefix,json=framedIpv6Prefix,proto3" json:"framed_ipv6_prefix,omitempty"`
	FramedPrefixLength   uint32                     `protobuf:"varint,71,opt,name=framed_prefix_length,json=framedPrefixLength,proto3" json:"framed_prefix_length,omitempty"`
	IsNakNextRenew       bool                       `protobuf:"varint,72,opt,name=is_nak_next_renew,json=isNakNextRenew,proto3" json:"is_nak_next_renew,omitempty"`
	SrgState             uint32                     `protobuf:"varint,73,opt,name=srg_state,json=srgState,proto3" json:"srg_state,omitempty"`
	SrgIntfRole          uint32                     `protobuf:"varint,74,opt,name=srg_intf_role,json=srgIntfRole,proto3" json:"srg_intf_role,omitempty"`
	Srgp2P               bool                       `protobuf:"varint,75,opt,name=srgp2p,proto3" json:"srgp2p,omitempty"`
	SrgVrfName           string                     `protobuf:"bytes,76,opt,name=srg_vrf_name,json=srgVrfName,proto3" json:"srg_vrf_name,omitempty"`
	SergState            uint32                     `protobuf:"varint,77,opt,name=serg_state,json=sergState,proto3" json:"serg_state,omitempty"`
	SergIntfRole         uint32                     `protobuf:"varint,78,opt,name=serg_intf_role,json=sergIntfRole,proto3" json:"serg_intf_role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Ipv6Dhcpv6DProxyBinding) Reset()         { *m = Ipv6Dhcpv6DProxyBinding{} }
func (m *Ipv6Dhcpv6DProxyBinding) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DProxyBinding) ProtoMessage()    {}
func (*Ipv6Dhcpv6DProxyBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5a3bcdfc78cc81a, []int{5}
}

func (m *Ipv6Dhcpv6DProxyBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DProxyBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DProxyBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DProxyBinding.Merge(m, src)
}
func (m *Ipv6Dhcpv6DProxyBinding) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DProxyBinding.Size(m)
}
func (m *Ipv6Dhcpv6DProxyBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DProxyBinding.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DProxyBinding proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DProxyBinding) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetClientFlag() uint32 {
	if m != nil {
		return m.ClientFlag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSubscriberLabel() uint32 {
	if m != nil {
		return m.SubscriberLabel
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIaIdPDs() uint32 {
	if m != nil {
		return m.IaIdPDs
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIaIdPd() *BagDhcpv6DIaIdPdInfoEntry {
	if m != nil {
		return m.IaIdPd
	}
	return nil
}

func (m *Ipv6Dhcpv6DProxyBinding) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetAccessVrfName() string {
	if m != nil {
		return m.AccessVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingTags() uint32 {
	if m != nil {
		return m.ProxyBindingTags
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingOuterTag() uint32 {
	if m != nil {
		return m.ProxyBindingOuterTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProxyBindingInnerTag() uint32 {
	if m != nil {
		return m.ProxyBindingInnerTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetRxRemoteId() string {
	if m != nil {
		return m.RxRemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetTxRemoteId() string {
	if m != nil {
		return m.TxRemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetRxInterfaceId() string {
	if m != nil {
		return m.RxInterfaceId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetTxInterfaceId() string {
	if m != nil {
		return m.TxInterfaceId
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetServerIpv6Address() string {
	if m != nil {
		return m.ServerIpv6Address
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetFramedIpv6Prefix() string {
	if m != nil {
		return m.FramedIpv6Prefix
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetFramedPrefixLength() uint32 {
	if m != nil {
		return m.FramedPrefixLength
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetIsNakNextRenew() bool {
	if m != nil {
		return m.IsNakNextRenew
	}
	return false
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgState() uint32 {
	if m != nil {
		return m.SrgState
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgIntfRole() uint32 {
	if m != nil {
		return m.SrgIntfRole
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgp2P() bool {
	if m != nil {
		return m.Srgp2P
	}
	return false
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSrgVrfName() string {
	if m != nil {
		return m.SrgVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSergState() uint32 {
	if m != nil {
		return m.SergState
	}
	return 0
}

func (m *Ipv6Dhcpv6DProxyBinding) GetSergIntfRole() uint32 {
	if m != nil {
		return m.SergIntfRole
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DProxyBinding_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.ipv6_dhcpv6d_proxy_binding_KEYS")
	proto.RegisterType((*BagDhcpv6DAddrAttrbItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_addr_attrb_item")
	proto.RegisterType((*BagDhcpv6DAddrAttrbEntry)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_addr_attrb_entry")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfoItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_ia_id_pd_info_item")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfoEntry)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.bag_dhcpv6d_ia_id_pd_info_entry")
	proto.RegisterType((*Ipv6Dhcpv6DProxyBinding)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.binding.clients.client.ipv6_dhcpv6d_proxy_binding")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_proxy_binding.proto", fileDescriptor_e5a3bcdfc78cc81a) }

var fileDescriptor_e5a3bcdfc78cc81a = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5d, 0x73, 0x1b, 0x35,
	0x17, 0x9e, 0x6d, 0x13, 0xc7, 0x3e, 0xce, 0x57, 0x15, 0x4f, 0xbb, 0x7d, 0xf3, 0x86, 0x18, 0x03,
	0x9d, 0x74, 0x86, 0xf1, 0x74, 0x52, 0x1a, 0xbe, 0x3f, 0x52, 0xd2, 0xc2, 0xd2, 0x10, 0x82, 0x9b,
	0x61, 0x86, 0xe1, 0x42, 0x23, 0xaf, 0xb4, 0x5b, 0x4d, 0x77, 0xb5, 0x8b, 0xa4, 0xb8, 0x9b, 0xff,
	0x00, 0xbf, 0x82, 0x3b, 0xb8, 0xe0, 0x17, 0x70, 0xc7, 0x35, 0xbf, 0x89, 0xd1, 0xd1, 0xae, 0xed,
	0x10, 0xcc, 0x15, 0xf4, 0x26, 0x59, 0x3d, 0xcf, 0x73, 0xce, 0x91, 0xce, 0x39, 0xd2, 0x31, 0xf4,
	0x65, 0x39, 0x39, 0xa0, 0xfc, 0x59, 0x5c, 0x4e, 0x0e, 0x38, 0x2d, 0x75, 0x51, 0x5d, 0xd0, 0xb1,
	0x54, 0x5c, 0xaa, 0x74, 0x58, 0xea, 0xc2, 0x16, 0xe4, 0xeb, 0x58, 0x9a, 0xb8, 0xa0, 0xb2, 0x30,
	0xb4, 0xd2, 0x14, 0xe5, 0x4a, 0xbc, 0x98, 0x9a, 0x14, 0xa5, 0xd0, 0x43, 0xbf, 0x18, 0xaa, 0x82,
	0x0b, 0x83, 0x7f, 0x87, 0xe8, 0x69, 0xd8, 0x78, 0x8a, 0x33, 0x29, 0x94, 0x35, 0xf5, 0xff, 0xc1,
	0x77, 0xb0, 0xbb, 0x38, 0x2c, 0x7d, 0xf2, 0xe8, 0xdb, 0xa7, 0x64, 0x1b, 0x3a, 0xce, 0x0f, 0x55,
	0x2c, 0x17, 0x61, 0xd0, 0x0f, 0xf6, 0x3a, 0xa3, 0xb6, 0x03, 0x4e, 0x58, 0x2e, 0x1c, 0xe9, 0x3d,
	0x51, 0xc9, 0xc3, 0x6b, 0x9e, 0xf4, 0x40, 0xc4, 0x07, 0xbf, 0x04, 0xb0, 0x3d, 0x66, 0xe9, 0xd4,
	0x39, 0xe3, 0x5c, 0x53, 0x66, 0xad, 0x1e, 0x53, 0x69, 0x45, 0x4e, 0x6e, 0x42, 0xab, 0xd4, 0x22,
	0x91, 0x55, 0xed, 0xb6, 0x5e, 0x91, 0xd7, 0x60, 0xcd, 0x7f, 0xd1, 0x4c, 0xa8, 0xd4, 0x3e, 0x43,
	0xc7, 0x6b, 0xa3, 0x55, 0x0f, 0x1e, 0x23, 0x46, 0x76, 0x00, 0x32, 0xc1, 0x8c, 0xa0, 0x56, 0xe6,
	0x22, 0xbc, 0x8e, 0x8a, 0x0e, 0x22, 0x67, 0x32, 0x17, 0xe4, 0x1e, 0xf4, 0xb4, 0xc8, 0x99, 0x54,
	0xee, 0x1c, 0x73, 0xc2, 0x25, 0x14, 0x92, 0x29, 0x77, 0xdc, 0x58, 0x0c, 0x7e, 0x0f, 0xe0, 0xff,
	0x0b, 0x76, 0x2b, 0x94, 0xd5, 0x17, 0xe4, 0xa7, 0x00, 0x6e, 0xfe, 0xbd, 0x20, 0x5c, 0xee, 0x5f,
	0xdf, 0xeb, 0xee, 0xab, 0xe1, 0xbf, 0x5e, 0xa0, 0xe1, 0x3f, 0xe4, 0x6f, 0xb4, 0x35, 0x66, 0xe9,
	0x91, 0xe7, 0x0e, 0x39, 0xd7, 0x87, 0x8e, 0x19, 0xfc, 0x7a, 0x0d, 0x5e, 0x99, 0x37, 0x92, 0x8c,
	0x4a, 0x4e, 0x4b, 0x4e, 0xa5, 0x4a, 0x0a, 0x9f, 0xf7, 0x5b, 0xb0, 0x22, 0x19, 0xb5, 0x17, 0x65,
	0x53, 0xcf, 0x96, 0x64, 0x67, 0x17, 0xa5, 0x20, 0x5b, 0xb0, 0x8c, 0xf2, 0x3a, 0xe1, 0x4b, 0x92,
	0x45, 0x9c, 0xf4, 0x60, 0x39, 0xc9, 0x58, 0x6a, 0xea, 0x1c, 0xfb, 0x85, 0xab, 0x91, 0x2d, 0x2c,
	0xcb, 0x70, 0x53, 0xc2, 0x98, 0x3a, 0xb1, 0xab, 0x08, 0x1e, 0x7a, 0xcc, 0x99, 0x1a, 0xcb, 0xac,
	0x08, 0x97, 0x31, 0x8c, 0x5f, 0x90, 0x1f, 0x03, 0xe8, 0xd4, 0x56, 0xc2, 0x84, 0xad, 0x7e, 0xb0,
	0xd7, 0xdd, 0x2f, 0x5e, 0x5e, 0xea, 0xb0, 0x98, 0xa3, 0xd9, 0x0e, 0x06, 0x7f, 0x04, 0xb0, 0xbb,
	0x38, 0x63, 0xbe, 0xf6, 0x3f, 0x07, 0x70, 0x7b, 0xa1, 0x26, 0x5c, 0xc1, 0xf2, 0x7f, 0xff, 0x1f,
	0x9f, 0xe1, 0x6a, 0x25, 0x47, 0xbd, 0x59, 0x07, 0x44, 0x2c, 0xe2, 0xa7, 0x3c, 0x52, 0x49, 0x31,
	0xf8, 0xad, 0x03, 0xff, 0x5b, 0x7c, 0xab, 0x09, 0x81, 0x25, 0x7e, 0x2e, 0x79, 0xb8, 0x8f, 0x45,
	0xc1, 0x6f, 0xb2, 0x0b, 0xdd, 0xfa, 0x1e, 0xbb, 0xf2, 0x86, 0xf7, 0xb1, 0x98, 0xe0, 0xa1, 0xc7,
	0x19, 0x4b, 0xc9, 0x5d, 0xd8, 0x34, 0xe7, 0x63, 0x13, 0x6b, 0x39, 0x16, 0x9a, 0x66, 0x6c, 0x2c,
	0xb2, 0xf0, 0x2d, 0x54, 0x6d, 0xcc, 0xf0, 0x63, 0x07, 0x93, 0xdb, 0xd0, 0x9e, 0xe8, 0xc4, 0xbf,
	0x17, 0x0f, 0x30, 0xc6, 0xca, 0x44, 0x27, 0xf8, 0x5c, 0xec, 0x42, 0x37, 0x67, 0xf1, 0xb4, 0x67,
	0x0e, 0x90, 0x85, 0x9c, 0xc5, 0x4d, 0xc7, 0x6c, 0x03, 0xd4, 0xc7, 0xa4, 0xdc, 0x84, 0x6f, 0x63,
	0x80, 0x15, 0xd7, 0x86, 0xa7, 0x47, 0x86, 0xfc, 0x10, 0x40, 0xbb, 0x49, 0x42, 0xf8, 0x0e, 0xf6,
	0x8d, 0x7e, 0xa9, 0x39, 0xf7, 0xad, 0xd3, 0xc2, 0xfd, 0x70, 0xf2, 0x06, 0xac, 0x4b, 0x65, 0x85,
	0x4e, 0x58, 0x5c, 0xbf, 0x8e, 0xef, 0xe2, 0x79, 0xd6, 0xa6, 0x28, 0x9e, 0xf9, 0x0e, 0x6c, 0xb0,
	0x38, 0x16, 0xc6, 0xd0, 0x69, 0x56, 0xde, 0xf3, 0x3a, 0x0f, 0x7f, 0x53, 0xe7, 0xe6, 0x4d, 0x20,
	0x97, 0x5f, 0x5f, 0xeb, 0x2e, 0xdd, 0xfb, 0x98, 0x82, 0x4d, 0x64, 0x1e, 0x7a, 0xe2, 0xcc, 0xdd,
	0xbf, 0x07, 0x70, 0xeb, 0xb2, 0xba, 0x38, 0xb7, 0x42, 0x3b, 0x9b, 0xf0, 0x03, 0x34, 0xe9, 0xcd,
	0x9b, 0x7c, 0xe5, 0xc8, 0x33, 0x96, 0x5e, 0x35, 0x93, 0x4a, 0xd5, 0x66, 0x1f, 0x5e, 0x35, 0x8b,
	0x1c, 0xe9, 0xcc, 0x76, 0x00, 0xe2, 0x8c, 0x19, 0xe3, 0xb7, 0xff, 0x11, 0x6e, 0xbf, 0x83, 0x48,
	0x33, 0x05, 0xca, 0xa2, 0xc8, 0x3c, 0xfb, 0xb1, 0x9f, 0x02, 0x0e, 0x40, 0xb2, 0x0f, 0xab, 0xba,
	0xa2, 0x5a, 0xe4, 0x85, 0x15, 0xee, 0x6d, 0xf9, 0xc4, 0x17, 0x5d, 0x57, 0x23, 0x84, 0x22, 0xee,
	0x14, 0x76, 0x5e, 0x71, 0xe8, 0x15, 0x76, 0xa6, 0xb8, 0x03, 0x1b, 0xba, 0xa2, 0xb3, 0x6c, 0x4b,
	0x1e, 0x3e, 0xf4, 0x39, 0xd4, 0x55, 0xd4, 0xa0, 0x5e, 0x67, 0xff, 0xa2, 0xfb, 0xd4, 0xeb, 0xec,
	0x25, 0xdd, 0x10, 0xb6, 0x8c, 0xd0, 0x13, 0x51, 0x77, 0x4c, 0xd3, 0x8f, 0x47, 0xa8, 0xbd, 0xe1,
	0xa9, 0xa8, 0x9c, 0x1c, 0x34, 0x6d, 0xf9, 0x2a, 0xac, 0x96, 0xba, 0x48, 0x64, 0x56, 0x17, 0xfa,
	0x11, 0x0a, 0xbb, 0x35, 0xd6, 0x94, 0x2f, 0xd1, 0x2c, 0x17, 0xdc, 0xbb, 0xac, 0x07, 0xdb, 0x63,
	0x14, 0x6e, 0x7a, 0xc6, 0x79, 0x3c, 0xf5, 0x23, 0xee, 0x1e, 0xf4, 0x6a, 0xf5, 0xe5, 0x49, 0xf7,
	0x99, 0x1f, 0x4f, 0x9e, 0x3b, 0x9d, 0x9f, 0x77, 0x77, 0xe1, 0x86, 0x74, 0xf9, 0x7f, 0x4e, 0x95,
	0xa8, 0x2c, 0xd5, 0x42, 0x89, 0x17, 0xe1, 0xe7, 0xfd, 0x60, 0xaf, 0x3d, 0x5a, 0x97, 0xe6, 0x84,
	0x3d, 0x3f, 0x11, 0x95, 0x1d, 0x39, 0xd4, 0x95, 0xc3, 0xe8, 0x94, 0xfa, 0xa7, 0x37, 0x42, 0x8f,
	0x6d, 0xa3, 0xd3, 0xa7, 0xf8, 0xfa, 0x0e, 0x60, 0xcd, 0x91, 0x52, 0xd9, 0x84, 0xea, 0x22, 0x13,
	0xe1, 0x17, 0x28, 0xe8, 0x1a, 0x9d, 0x46, 0xca, 0x26, 0xa3, 0x22, 0x13, 0x6e, 0x30, 0x1b, 0x9d,
	0x96, 0xfb, 0x65, 0xf8, 0x04, 0x03, 0xd4, 0x2b, 0x57, 0x28, 0x67, 0x3b, 0xed, 0xe3, 0x63, 0x5f,
	0x28, 0xa3, 0xd3, 0xa6, 0x89, 0x77, 0x00, 0x8c, 0x98, 0xc6, 0xfe, 0xd2, 0x4f, 0x65, 0x87, 0xf8,
	0xe0, 0xaf, 0xc3, 0x3a, 0xd2, 0xb3, 0xe8, 0x27, 0x7e, 0x6c, 0x38, 0xb4, 0x09, 0x3f, 0x6e, 0xe1,
	0xcf, 0x9d, 0xfb, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87, 0x9d, 0x54, 0xde, 0x12, 0x09, 0x00,
	0x00,
}
