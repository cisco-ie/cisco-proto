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
// source: ipv6_dhcpv6d_server_binding.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_binding_clients_client

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

type Ipv6Dhcpv6DServerBinding_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ClientId             string   `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBinding_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerBinding_KEYS{} }
func (m *Ipv6Dhcpv6DServerBinding_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBinding_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBinding_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9338fecc6ed048, []int{0}
}

func (m *Ipv6Dhcpv6DServerBinding_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBinding_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBinding_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBinding_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerBinding_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBinding_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBinding_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding_KEYS) GetClientId() string {
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
	return fileDescriptor_0c9338fecc6ed048, []int{1}
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
	BagDhcpv6DAddrAttrb  []*BagDhcpv6DAddrAttrbItem `protobuf:"bytes,1,rep,name=bag_dhcpv6d_addr_attrb,json=bagDhcpv6dAddrAttrb,proto3" json:"bag_dhcpv6d_addr_attrb,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *BagDhcpv6DAddrAttrbEntry) Reset()         { *m = BagDhcpv6DAddrAttrbEntry{} }
func (m *BagDhcpv6DAddrAttrbEntry) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DAddrAttrbEntry) ProtoMessage()    {}
func (*BagDhcpv6DAddrAttrbEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9338fecc6ed048, []int{2}
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
	return fileDescriptor_0c9338fecc6ed048, []int{3}
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
	BagDhcpv6DIaIdPdInfo []*BagDhcpv6DIaIdPdInfoItem `protobuf:"bytes,1,rep,name=bag_dhcpv6d_ia_id_pd_info,json=bagDhcpv6dIaIdPdInfo,proto3" json:"bag_dhcpv6d_ia_id_pd_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BagDhcpv6DIaIdPdInfoEntry) Reset()         { *m = BagDhcpv6DIaIdPdInfoEntry{} }
func (m *BagDhcpv6DIaIdPdInfoEntry) String() string { return proto.CompactTextString(m) }
func (*BagDhcpv6DIaIdPdInfoEntry) ProtoMessage()    {}
func (*BagDhcpv6DIaIdPdInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9338fecc6ed048, []int{4}
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

type Ipv6Dhcpv6DServerBinding struct {
	Duid                  string                     `protobuf:"bytes,50,opt,name=duid,proto3" json:"duid,omitempty"`
	ClientIdXr            uint32                     `protobuf:"varint,51,opt,name=client_id_xr,json=clientIdXr,proto3" json:"client_id_xr,omitempty"`
	ClientFlag            uint32                     `protobuf:"varint,52,opt,name=client_flag,json=clientFlag,proto3" json:"client_flag,omitempty"`
	SubscriberLabel       uint32                     `protobuf:"varint,53,opt,name=subscriber_label,json=subscriberLabel,proto3" json:"subscriber_label,omitempty"`
	VrfName               string                     `protobuf:"bytes,54,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	MacAddress            string                     `protobuf:"bytes,55,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	IaIdPDs               uint32                     `protobuf:"varint,56,opt,name=ia_id_p_ds,json=iaIdPDs,proto3" json:"ia_id_p_ds,omitempty"`
	IaIdPd                *BagDhcpv6DIaIdPdInfoEntry `protobuf:"bytes,57,opt,name=ia_id_pd,json=iaIdPd,proto3" json:"ia_id_pd,omitempty"`
	LinkLocalAddress      string                     `protobuf:"bytes,58,opt,name=link_local_address,json=linkLocalAddress,proto3" json:"link_local_address,omitempty"`
	InterfaceName         string                     `protobuf:"bytes,59,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	AccessVrfName         string                     `protobuf:"bytes,60,opt,name=access_vrf_name,json=accessVrfName,proto3" json:"access_vrf_name,omitempty"`
	ServerBindingTags     uint32                     `protobuf:"varint,61,opt,name=server_binding_tags,json=serverBindingTags,proto3" json:"server_binding_tags,omitempty"`
	ServerBindingOuterTag uint32                     `protobuf:"varint,62,opt,name=server_binding_outer_tag,json=serverBindingOuterTag,proto3" json:"server_binding_outer_tag,omitempty"`
	ServerBindingInnerTag uint32                     `protobuf:"varint,63,opt,name=server_binding_inner_tag,json=serverBindingInnerTag,proto3" json:"server_binding_inner_tag,omitempty"`
	PoolName              string                     `protobuf:"bytes,64,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	ProfileName           string                     `protobuf:"bytes,65,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	SelecetedProfileName  string                     `protobuf:"bytes,66,opt,name=seleceted_profile_name,json=selecetedProfileName,proto3" json:"seleceted_profile_name,omitempty"`
	FramedIpv6Prefix      string                     `protobuf:"bytes,67,opt,name=framed_ipv6_prefix,json=framedIpv6Prefix,proto3" json:"framed_ipv6_prefix,omitempty"`
	FramedPrefixLength    uint32                     `protobuf:"varint,68,opt,name=framed_prefix_length,json=framedPrefixLength,proto3" json:"framed_prefix_length,omitempty"`
	ClassName             string                     `protobuf:"bytes,69,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	RxRemoteId            string                     `protobuf:"bytes,70,opt,name=rx_remote_id,json=rxRemoteId,proto3" json:"rx_remote_id,omitempty"`
	RxInterfaceId         string                     `protobuf:"bytes,71,opt,name=rx_interface_id,json=rxInterfaceId,proto3" json:"rx_interface_id,omitempty"`
	PrefixPoolName        string                     `protobuf:"bytes,72,opt,name=prefix_pool_name,json=prefixPoolName,proto3" json:"prefix_pool_name,omitempty"`
	AddressPoolName       string                     `protobuf:"bytes,73,opt,name=address_pool_name,json=addressPoolName,proto3" json:"address_pool_name,omitempty"`
	DnsServerCount        uint32                     `protobuf:"varint,74,opt,name=dns_server_count,json=dnsServerCount,proto3" json:"dns_server_count,omitempty"`
	IsNakNextRenew        bool                       `protobuf:"varint,75,opt,name=is_nak_next_renew,json=isNakNextRenew,proto3" json:"is_nak_next_renew,omitempty"`
	SrgState              uint32                     `protobuf:"varint,76,opt,name=srg_state,json=srgState,proto3" json:"srg_state,omitempty"`
	SrgIntfRole           uint32                     `protobuf:"varint,77,opt,name=srg_intf_role,json=srgIntfRole,proto3" json:"srg_intf_role,omitempty"`
	Srgp2P                bool                       `protobuf:"varint,78,opt,name=srgp2p,proto3" json:"srgp2p,omitempty"`
	SrgVrfName            string                     `protobuf:"bytes,79,opt,name=srg_vrf_name,json=srgVrfName,proto3" json:"srg_vrf_name,omitempty"`
	SesrgState            uint32                     `protobuf:"varint,80,opt,name=sesrg_state,json=sesrgState,proto3" json:"sesrg_state,omitempty"`
	SergIntfRole          uint32                     `protobuf:"varint,81,opt,name=serg_intf_role,json=sergIntfRole,proto3" json:"serg_intf_role,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                   `json:"-"`
	XXX_unrecognized      []byte                     `json:"-"`
	XXX_sizecache         int32                      `json:"-"`
}

func (m *Ipv6Dhcpv6DServerBinding) Reset()         { *m = Ipv6Dhcpv6DServerBinding{} }
func (m *Ipv6Dhcpv6DServerBinding) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerBinding) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerBinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c9338fecc6ed048, []int{5}
}

func (m *Ipv6Dhcpv6DServerBinding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerBinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerBinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerBinding.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerBinding) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerBinding.Size(m)
}
func (m *Ipv6Dhcpv6DServerBinding) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerBinding.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerBinding proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerBinding) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetClientIdXr() uint32 {
	if m != nil {
		return m.ClientIdXr
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetClientFlag() uint32 {
	if m != nil {
		return m.ClientFlag
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetSubscriberLabel() uint32 {
	if m != nil {
		return m.SubscriberLabel
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetIaIdPDs() uint32 {
	if m != nil {
		return m.IaIdPDs
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetIaIdPd() *BagDhcpv6DIaIdPdInfoEntry {
	if m != nil {
		return m.IaIdPd
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerBinding) GetLinkLocalAddress() string {
	if m != nil {
		return m.LinkLocalAddress
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetAccessVrfName() string {
	if m != nil {
		return m.AccessVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetServerBindingTags() uint32 {
	if m != nil {
		return m.ServerBindingTags
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetServerBindingOuterTag() uint32 {
	if m != nil {
		return m.ServerBindingOuterTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetServerBindingInnerTag() uint32 {
	if m != nil {
		return m.ServerBindingInnerTag
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetSelecetedProfileName() string {
	if m != nil {
		return m.SelecetedProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetFramedIpv6Prefix() string {
	if m != nil {
		return m.FramedIpv6Prefix
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetFramedPrefixLength() uint32 {
	if m != nil {
		return m.FramedPrefixLength
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetRxRemoteId() string {
	if m != nil {
		return m.RxRemoteId
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetRxInterfaceId() string {
	if m != nil {
		return m.RxInterfaceId
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetPrefixPoolName() string {
	if m != nil {
		return m.PrefixPoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetAddressPoolName() string {
	if m != nil {
		return m.AddressPoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetDnsServerCount() uint32 {
	if m != nil {
		return m.DnsServerCount
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetIsNakNextRenew() bool {
	if m != nil {
		return m.IsNakNextRenew
	}
	return false
}

func (m *Ipv6Dhcpv6DServerBinding) GetSrgState() uint32 {
	if m != nil {
		return m.SrgState
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetSrgIntfRole() uint32 {
	if m != nil {
		return m.SrgIntfRole
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetSrgp2P() bool {
	if m != nil {
		return m.Srgp2P
	}
	return false
}

func (m *Ipv6Dhcpv6DServerBinding) GetSrgVrfName() string {
	if m != nil {
		return m.SrgVrfName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerBinding) GetSesrgState() uint32 {
	if m != nil {
		return m.SesrgState
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerBinding) GetSergIntfRole() uint32 {
	if m != nil {
		return m.SergIntfRole
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerBinding_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.ipv6_dhcpv6d_server_binding_KEYS")
	proto.RegisterType((*BagDhcpv6DAddrAttrbItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.bag_dhcpv6d_addr_attrb_item")
	proto.RegisterType((*BagDhcpv6DAddrAttrbEntry)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.bag_dhcpv6d_addr_attrb_entry")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfoItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.bag_dhcpv6d_ia_id_pd_info_item")
	proto.RegisterType((*BagDhcpv6DIaIdPdInfoEntry)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.bag_dhcpv6d_ia_id_pd_info_entry")
	proto.RegisterType((*Ipv6Dhcpv6DServerBinding)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.binding.clients.client.ipv6_dhcpv6d_server_binding")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_server_binding.proto", fileDescriptor_0c9338fecc6ed048) }

var fileDescriptor_0c9338fecc6ed048 = []byte{
	// 1028 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcb, 0x6e, 0x1b, 0x37,
	0x17, 0xc6, 0x24, 0xbe, 0xe9, 0xc8, 0x57, 0xda, 0xbf, 0xff, 0x09, 0xdc, 0xd6, 0x8a, 0x7a, 0x81,
	0x52, 0x04, 0x42, 0xe0, 0xa4, 0x4e, 0xef, 0xad, 0x13, 0x27, 0xed, 0x34, 0xae, 0xa3, 0xca, 0x46,
	0xd0, 0x02, 0x05, 0x08, 0x6a, 0x48, 0x4d, 0x08, 0x8f, 0x38, 0x03, 0x92, 0x56, 0xe4, 0x87, 0x68,
	0x9f, 0xa2, 0xbb, 0x6c, 0xfa, 0x12, 0x5d, 0x77, 0xd1, 0x17, 0x2a, 0x78, 0xc8, 0x91, 0x64, 0xd7,
	0xf6, 0xaa, 0xcd, 0xc6, 0x16, 0xbf, 0xef, 0x3b, 0xbc, 0x7c, 0xe7, 0x90, 0x67, 0xe0, 0xb6, 0x2c,
	0x87, 0xbb, 0x94, 0xbf, 0x4c, 0xcb, 0xe1, 0x2e, 0xa7, 0x46, 0xe8, 0xa1, 0xd0, 0xb4, 0x27, 0x15,
	0x97, 0x2a, 0x6b, 0x97, 0xba, 0xb0, 0x05, 0xe9, 0xa6, 0xd2, 0xa4, 0x05, 0x95, 0x85, 0xa1, 0x23,
	0x4d, 0x51, 0xaf, 0xc4, 0xab, 0x71, 0x4c, 0x51, 0x0a, 0xdd, 0xf6, 0x83, 0xb6, 0x2a, 0xb8, 0x30,
	0xf8, 0xb7, 0xed, 0xa7, 0x6a, 0x57, 0x53, 0xa5, 0xb9, 0x14, 0xca, 0x9a, 0xf0, 0xbf, 0xf9, 0x33,
	0x34, 0xae, 0x59, 0x98, 0x3e, 0x7b, 0xf2, 0xd3, 0x11, 0xd9, 0x82, 0x9a, 0x9b, 0x89, 0x2a, 0x36,
	0x10, 0x71, 0xd4, 0x88, 0x5a, 0xb5, 0xee, 0x82, 0x03, 0x0e, 0xd9, 0x40, 0x38, 0xd2, 0x4f, 0x45,
	0x25, 0x8f, 0x6f, 0x78, 0xd2, 0x03, 0x09, 0x6f, 0xbe, 0x8e, 0x60, 0xab, 0xc7, 0xb2, 0xf1, 0xec,
	0x8c, 0x73, 0x4d, 0x99, 0xb5, 0xba, 0x47, 0xa5, 0x15, 0x03, 0xb2, 0x09, 0x73, 0xa5, 0x16, 0x7d,
	0x39, 0x0a, 0xd3, 0x86, 0x11, 0x79, 0x17, 0x96, 0xfc, 0x2f, 0x9a, 0x0b, 0x95, 0xd9, 0x97, 0x38,
	0xf1, 0x52, 0x77, 0xd1, 0x83, 0x07, 0x88, 0x91, 0xb7, 0x01, 0x72, 0xc1, 0x8c, 0xa0, 0x56, 0x0e,
	0x44, 0x7c, 0x13, 0x15, 0x35, 0x44, 0x8e, 0xe5, 0x40, 0x90, 0x7b, 0xb0, 0xa1, 0xc5, 0x80, 0x49,
	0xe5, 0xce, 0x31, 0x25, 0x9c, 0x41, 0x21, 0x19, 0x73, 0x07, 0x55, 0x44, 0xf3, 0x8f, 0x08, 0xde,
	0xba, 0x62, 0xb7, 0x42, 0x59, 0x7d, 0x46, 0x7e, 0x8b, 0x60, 0xf3, 0x72, 0x41, 0x1c, 0x35, 0x6e,
	0xb6, 0xea, 0x3b, 0x45, 0xfb, 0xdf, 0x4f, 0x51, 0xfb, 0x1a, 0x03, 0xbb, 0xeb, 0x3d, 0x96, 0xed,
	0x7b, 0x6e, 0x8f, 0x73, 0xbd, 0xe7, 0x98, 0xe6, 0xef, 0x37, 0xe0, 0x9d, 0xe9, 0x20, 0xc9, 0xa8,
	0xe4, 0xb4, 0xe4, 0x54, 0xaa, 0x7e, 0xe1, 0x8d, 0xff, 0x3f, 0xcc, 0x4b, 0x46, 0xed, 0x59, 0x59,
	0x25, 0x74, 0x4e, 0xb2, 0xe3, 0xb3, 0x52, 0x90, 0x75, 0x98, 0x45, 0x79, 0x70, 0x7c, 0x46, 0xb2,
	0x84, 0x93, 0x0d, 0x98, 0xed, 0xe7, 0x2c, 0x33, 0xc1, 0x64, 0x3f, 0x70, 0x49, 0xb2, 0x85, 0x65,
	0x39, 0x6e, 0x4a, 0x18, 0x13, 0x9c, 0x5d, 0x44, 0x70, 0xcf, 0x63, 0x2e, 0xd4, 0x58, 0x66, 0x45,
	0x3c, 0x8b, 0xcb, 0xf8, 0x01, 0xf9, 0x35, 0x82, 0x5a, 0x88, 0x12, 0x26, 0x9e, 0x6b, 0x44, 0xad,
	0xfa, 0x4e, 0xf9, 0x06, 0xbd, 0xc3, 0x74, 0x76, 0x27, 0x5b, 0x68, 0xfe, 0x19, 0xc1, 0xf6, 0xd5,
	0x96, 0xf9, 0xec, 0xbf, 0x8e, 0xe0, 0xd6, 0x95, 0x9a, 0x50, 0x00, 0xfa, 0xbf, 0x3e, 0xc4, 0x3f,
	0x73, 0xd9, 0xdd, 0x98, 0xd4, 0x40, 0xc2, 0x12, 0xde, 0xe1, 0x89, 0xea, 0x17, 0xcd, 0xbf, 0x00,
	0xb6, 0xae, 0xb9, 0xd9, 0x84, 0xc0, 0x0c, 0x3f, 0x95, 0x3c, 0xde, 0xc1, 0xbc, 0xe0, 0x6f, 0xd2,
	0x80, 0xc5, 0xf1, 0x5d, 0xa6, 0x23, 0x1d, 0xdf, 0xc7, 0x84, 0x42, 0x75, 0x9d, 0x7f, 0xd4, 0x64,
	0x1b, 0xea, 0x41, 0xe1, 0x6a, 0x20, 0x7e, 0x30, 0x2d, 0x78, 0x9a, 0xb3, 0x8c, 0xdc, 0x81, 0x55,
	0x73, 0xda, 0x33, 0xa9, 0x96, 0x3d, 0xa1, 0x69, 0xce, 0x7a, 0x22, 0x8f, 0x3f, 0x42, 0xd5, 0xca,
	0x04, 0x3f, 0x70, 0x30, 0xb9, 0x05, 0x0b, 0x43, 0xdd, 0xf7, 0xaf, 0xca, 0x2e, 0xee, 0x62, 0x7e,
	0xa8, 0xfb, 0xf8, 0xa8, 0x6c, 0x43, 0x7d, 0xc0, 0xd2, 0x71, 0x61, 0x3d, 0x44, 0x16, 0x06, 0x2c,
	0xad, 0xca, 0x6a, 0x0b, 0x20, 0x38, 0x41, 0xb9, 0x89, 0x3f, 0xc6, 0x05, 0xe6, 0x5d, 0xad, 0x76,
	0xf6, 0x0d, 0xf9, 0x25, 0x82, 0x85, 0xca, 0xa7, 0xf8, 0x13, 0x2c, 0x2e, 0xf3, 0x66, 0xf3, 0xe2,
	0xeb, 0x6b, 0x0e, 0x37, 0xc4, 0xc9, 0x5d, 0x20, 0xb9, 0x54, 0x27, 0x34, 0x2f, 0xd2, 0xa9, 0xdb,
	0xf2, 0x29, 0x1e, 0x6a, 0xd5, 0x31, 0x07, 0x8e, 0xa8, 0x8e, 0xf6, 0x3e, 0x2c, 0x4b, 0x65, 0x85,
	0xee, 0xb3, 0x34, 0x3c, 0xb9, 0x9f, 0xa1, 0x72, 0x69, 0x8c, 0xa2, 0x45, 0x1f, 0xc0, 0x0a, 0x4b,
	0x53, 0x61, 0x0c, 0x1d, 0x9b, 0xf8, 0xb9, 0xd7, 0x79, 0xf8, 0x45, 0xb0, 0xb2, 0x0d, 0xeb, 0x17,
	0xde, 0x74, 0xeb, 0x6e, 0xf2, 0x17, 0x68, 0xd9, 0x9a, 0xa7, 0x1e, 0x79, 0xe6, 0xd8, 0xdd, 0xea,
	0x87, 0x10, 0x5f, 0xd0, 0x17, 0xa7, 0x56, 0x68, 0x17, 0x15, 0x7f, 0x89, 0x41, 0xff, 0x3b, 0x17,
	0xf4, 0xdc, 0xb1, 0xc7, 0x2c, 0xbb, 0x24, 0x50, 0x2a, 0x15, 0x02, 0xbf, 0xba, 0x24, 0x30, 0x71,
	0xac, 0x0b, 0xdc, 0x82, 0x5a, 0x59, 0x14, 0xb9, 0x3f, 0xc3, 0xd7, 0xbe, 0x83, 0x38, 0x00, 0xb7,
	0x7f, 0x1b, 0x16, 0x4b, 0x5d, 0xf4, 0x65, 0x1e, 0xbc, 0xd8, 0x43, 0xbe, 0x1e, 0x30, 0x94, 0x3c,
	0x80, 0x4d, 0x23, 0x72, 0x91, 0x0a, 0x2b, 0x38, 0x3d, 0x27, 0x7e, 0x84, 0xe2, 0x8d, 0x31, 0xdb,
	0x99, 0x8a, 0xba, 0x0b, 0xa4, 0xaf, 0xd9, 0x40, 0x70, 0x5f, 0x0c, 0xa1, 0x0d, 0x3d, 0xf6, 0x49,
	0xf1, 0x4c, 0x52, 0x0e, 0x77, 0x3b, 0xbe, 0x21, 0xdd, 0x83, 0x8d, 0xa0, 0x3e, 0xdf, 0x97, 0xf6,
	0x7d, 0x33, 0xf1, 0x5c, 0xe7, 0x42, 0x77, 0x4a, 0x73, 0x66, 0x8c, 0xdf, 0xc9, 0x13, 0x9c, 0xb7,
	0x86, 0x08, 0x2e, 0xdf, 0x80, 0x45, 0x3d, 0xa2, 0x5a, 0x0c, 0x0a, 0x2b, 0xdc, 0x73, 0xfb, 0xd4,
	0x97, 0xb8, 0x1e, 0x75, 0x11, 0x4a, 0xb8, 0x4b, 0xb0, 0x1e, 0xd1, 0x49, 0x29, 0x48, 0x1e, 0x7f,
	0xe3, 0x13, 0xac, 0x47, 0x49, 0x85, 0x26, 0x9c, 0xb4, 0x60, 0x35, 0xec, 0x69, 0xe2, 0xe2, 0xb7,
	0x28, 0x5c, 0xf6, 0x78, 0xa7, 0xf2, 0xf2, 0x43, 0x58, 0x0b, 0xc5, 0x37, 0x25, 0x4d, 0x50, 0xba,
	0x12, 0x88, 0xb1, 0xb6, 0x05, 0xab, 0x5c, 0x99, 0xea, 0xd1, 0x48, 0x8b, 0x53, 0x65, 0xe3, 0xef,
	0xf0, 0xb0, 0xcb, 0x5c, 0x99, 0x23, 0x84, 0x1f, 0x3b, 0x94, 0xdc, 0x81, 0x35, 0xe9, 0x4e, 0x79,
	0x42, 0x95, 0x18, 0x59, 0xaa, 0x85, 0x12, 0xaf, 0xe2, 0x67, 0x8d, 0xa8, 0xb5, 0xd0, 0x5d, 0x96,
	0xe6, 0x90, 0x9d, 0x1c, 0x8a, 0x91, 0xed, 0x3a, 0xd4, 0x65, 0xda, 0xe8, 0x8c, 0xfa, 0x86, 0x70,
	0x80, 0xb3, 0x2d, 0x18, 0x9d, 0x1d, 0x61, 0x4f, 0x68, 0xc2, 0x92, 0x23, 0xa5, 0xb2, 0x7d, 0xaa,
	0x8b, 0x5c, 0xc4, 0xdf, 0xa3, 0xa0, 0x6e, 0x74, 0x96, 0x28, 0xdb, 0xef, 0x16, 0xb9, 0x70, 0xdf,
	0x0b, 0x46, 0x67, 0xe5, 0x4e, 0x19, 0x1f, 0xe2, 0x02, 0x61, 0xe4, 0xdc, 0x74, 0xb1, 0xe3, 0x9b,
	0xf0, 0xdc, 0xbb, 0x69, 0x74, 0xf6, 0x62, 0xf2, 0xa2, 0x18, 0x31, 0x59, 0xbc, 0xe3, 0x1f, 0x2e,
	0x84, 0xfc, 0xf2, 0xef, 0xc1, 0xb2, 0x11, 0xe7, 0xd6, 0xff, 0xc1, 0xb7, 0x33, 0x87, 0x56, 0x1b,
	0xe8, 0xcd, 0xe1, 0x97, 0xd8, 0xfd, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbf, 0xbe, 0xd4, 0x55,
	0xae, 0x09, 0x00, 0x00,
}
