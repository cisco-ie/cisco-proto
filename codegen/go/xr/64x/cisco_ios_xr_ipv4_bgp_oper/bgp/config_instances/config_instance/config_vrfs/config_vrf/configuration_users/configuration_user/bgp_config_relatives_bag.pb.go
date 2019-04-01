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
// source: bgp_config_relatives_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_config_instances_config_instance_config_vrfs_config_vrf_configuration_users_configuration_user

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

type BgpConfigRelativesBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	EntityType           string   `protobuf:"bytes,3,opt,name=entity_type,json=entityType,proto3" json:"entity_type,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	EntityName           string   `protobuf:"bytes,5,opt,name=entity_name,json=entityName,proto3" json:"entity_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpConfigRelativesBag_KEYS) Reset()         { *m = BgpConfigRelativesBag_KEYS{} }
func (m *BgpConfigRelativesBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpConfigRelativesBag_KEYS) ProtoMessage()    {}
func (*BgpConfigRelativesBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_83d6410b1bbf056e, []int{0}
}

func (m *BgpConfigRelativesBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConfigRelativesBag_KEYS.Unmarshal(m, b)
}
func (m *BgpConfigRelativesBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConfigRelativesBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpConfigRelativesBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConfigRelativesBag_KEYS.Merge(m, src)
}
func (m *BgpConfigRelativesBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpConfigRelativesBag_KEYS.Size(m)
}
func (m *BgpConfigRelativesBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConfigRelativesBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConfigRelativesBag_KEYS proto.InternalMessageInfo

func (m *BgpConfigRelativesBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpConfigRelativesBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpConfigRelativesBag_KEYS) GetEntityType() string {
	if m != nil {
		return m.EntityType
	}
	return ""
}

func (m *BgpConfigRelativesBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *BgpConfigRelativesBag_KEYS) GetEntityName() string {
	if m != nil {
		return m.EntityName
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
	return fileDescriptor_83d6410b1bbf056e, []int{1}
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
	return fileDescriptor_83d6410b1bbf056e, []int{2}
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
	return fileDescriptor_83d6410b1bbf056e, []int{3}
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
	return fileDescriptor_83d6410b1bbf056e, []int{4}
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
	return fileDescriptor_83d6410b1bbf056e, []int{5}
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

type BgpConfigEntid_Item struct {
	AddressFamilyIdentifier uint32       `protobuf:"varint,1,opt,name=address_family_identifier,json=addressFamilyIdentifier,proto3" json:"address_family_identifier,omitempty"`
	ConfigurationType       string       `protobuf:"bytes,2,opt,name=configuration_type,json=configurationType,proto3" json:"configuration_type,omitempty"`
	NeighborAddressXr       *BgpAddrtype `protobuf:"bytes,3,opt,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	GroupName               string       `protobuf:"bytes,4,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}     `json:"-"`
	XXX_unrecognized        []byte       `json:"-"`
	XXX_sizecache           int32        `json:"-"`
}

func (m *BgpConfigEntid_Item) Reset()         { *m = BgpConfigEntid_Item{} }
func (m *BgpConfigEntid_Item) String() string { return proto.CompactTextString(m) }
func (*BgpConfigEntid_Item) ProtoMessage()    {}
func (*BgpConfigEntid_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_83d6410b1bbf056e, []int{6}
}

func (m *BgpConfigEntid_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConfigEntid_Item.Unmarshal(m, b)
}
func (m *BgpConfigEntid_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConfigEntid_Item.Marshal(b, m, deterministic)
}
func (m *BgpConfigEntid_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConfigEntid_Item.Merge(m, src)
}
func (m *BgpConfigEntid_Item) XXX_Size() int {
	return xxx_messageInfo_BgpConfigEntid_Item.Size(m)
}
func (m *BgpConfigEntid_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConfigEntid_Item.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConfigEntid_Item proto.InternalMessageInfo

func (m *BgpConfigEntid_Item) GetAddressFamilyIdentifier() uint32 {
	if m != nil {
		return m.AddressFamilyIdentifier
	}
	return 0
}

func (m *BgpConfigEntid_Item) GetConfigurationType() string {
	if m != nil {
		return m.ConfigurationType
	}
	return ""
}

func (m *BgpConfigEntid_Item) GetNeighborAddressXr() *BgpAddrtype {
	if m != nil {
		return m.NeighborAddressXr
	}
	return nil
}

func (m *BgpConfigEntid_Item) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type BgpConfigEntid_Entry struct {
	BgpConfigEntid_      []*BgpConfigEntid_Item `protobuf:"bytes,1,rep,name=bgp_config_entid_,json=bgpConfigEntid,proto3" json:"bgp_config_entid_,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BgpConfigEntid_Entry) Reset()         { *m = BgpConfigEntid_Entry{} }
func (m *BgpConfigEntid_Entry) String() string { return proto.CompactTextString(m) }
func (*BgpConfigEntid_Entry) ProtoMessage()    {}
func (*BgpConfigEntid_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_83d6410b1bbf056e, []int{7}
}

func (m *BgpConfigEntid_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConfigEntid_Entry.Unmarshal(m, b)
}
func (m *BgpConfigEntid_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConfigEntid_Entry.Marshal(b, m, deterministic)
}
func (m *BgpConfigEntid_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConfigEntid_Entry.Merge(m, src)
}
func (m *BgpConfigEntid_Entry) XXX_Size() int {
	return xxx_messageInfo_BgpConfigEntid_Entry.Size(m)
}
func (m *BgpConfigEntid_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConfigEntid_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConfigEntid_Entry proto.InternalMessageInfo

func (m *BgpConfigEntid_Entry) GetBgpConfigEntid_() []*BgpConfigEntid_Item {
	if m != nil {
		return m.BgpConfigEntid_
	}
	return nil
}

type BgpConfigRelativesBag struct {
	NeighborAddressXr       *BgpAddrtype            `protobuf:"bytes,50,opt,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	GroupName               string                  `protobuf:"bytes,51,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	ConfigurationType       string                  `protobuf:"bytes,52,opt,name=configuration_type,json=configurationType,proto3" json:"configuration_type,omitempty"`
	AddressFamilyIdentifier uint32                  `protobuf:"varint,53,opt,name=address_family_identifier,json=addressFamilyIdentifier,proto3" json:"address_family_identifier,omitempty"`
	AfIndependentRelatives  *BgpConfigEntid_Entry   `protobuf:"bytes,54,opt,name=af_independent_relatives,json=afIndependentRelatives,proto3" json:"af_independent_relatives,omitempty"`
	AfDependentRelative     []*BgpConfigEntid_Entry `protobuf:"bytes,55,rep,name=af_dependent_relative,json=afDependentRelative,proto3" json:"af_dependent_relative,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                `json:"-"`
	XXX_unrecognized        []byte                  `json:"-"`
	XXX_sizecache           int32                   `json:"-"`
}

func (m *BgpConfigRelativesBag) Reset()         { *m = BgpConfigRelativesBag{} }
func (m *BgpConfigRelativesBag) String() string { return proto.CompactTextString(m) }
func (*BgpConfigRelativesBag) ProtoMessage()    {}
func (*BgpConfigRelativesBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_83d6410b1bbf056e, []int{8}
}

func (m *BgpConfigRelativesBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpConfigRelativesBag.Unmarshal(m, b)
}
func (m *BgpConfigRelativesBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpConfigRelativesBag.Marshal(b, m, deterministic)
}
func (m *BgpConfigRelativesBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpConfigRelativesBag.Merge(m, src)
}
func (m *BgpConfigRelativesBag) XXX_Size() int {
	return xxx_messageInfo_BgpConfigRelativesBag.Size(m)
}
func (m *BgpConfigRelativesBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpConfigRelativesBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpConfigRelativesBag proto.InternalMessageInfo

func (m *BgpConfigRelativesBag) GetNeighborAddressXr() *BgpAddrtype {
	if m != nil {
		return m.NeighborAddressXr
	}
	return nil
}

func (m *BgpConfigRelativesBag) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *BgpConfigRelativesBag) GetConfigurationType() string {
	if m != nil {
		return m.ConfigurationType
	}
	return ""
}

func (m *BgpConfigRelativesBag) GetAddressFamilyIdentifier() uint32 {
	if m != nil {
		return m.AddressFamilyIdentifier
	}
	return 0
}

func (m *BgpConfigRelativesBag) GetAfIndependentRelatives() *BgpConfigEntid_Entry {
	if m != nil {
		return m.AfIndependentRelatives
	}
	return nil
}

func (m *BgpConfigRelativesBag) GetAfDependentRelative() []*BgpConfigEntid_Entry {
	if m != nil {
		return m.AfDependentRelative
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpConfigRelativesBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_config_relatives_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_addrtype")
	proto.RegisterType((*BgpConfigEntid_Item)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_config_entid__item")
	proto.RegisterType((*BgpConfigEntid_Entry)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_config_entid__entry")
	proto.RegisterType((*BgpConfigRelativesBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.config_instances.config_instance.config_vrfs.config_vrf.configuration_users.configuration_user.bgp_config_relatives_bag")
}

func init() { proto.RegisterFile("bgp_config_relatives_bag.proto", fileDescriptor_83d6410b1bbf056e) }

var fileDescriptor_83d6410b1bbf056e = []byte{
	// 965 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0xd6, 0x64, 0x61, 0x93, 0x2d, 0xdb, 0xbb, 0xf6, 0x38, 0xeb, 0x9d, 0x45, 0x0a, 0x2c, 0xe6,
	0xc0, 0x46, 0x0a, 0x96, 0x70, 0xcc, 0x04, 0x21, 0x2e, 0x28, 0x6c, 0xa4, 0x88, 0x6c, 0x84, 0x9c,
	0x08, 0xc1, 0xa9, 0xd5, 0xb6, 0x7b, 0x4c, 0x4b, 0xe3, 0x99, 0x56, 0xf7, 0xd8, 0x1b, 0x1f, 0x38,
	0x72, 0xe2, 0x00, 0x07, 0x7e, 0x5e, 0x06, 0x21, 0x21, 0x8e, 0xc0, 0x81, 0x07, 0xe0, 0xc0, 0x9b,
	0xa0, 0xae, 0x9e, 0x9e, 0x3f, 0x8f, 0x57, 0xe2, 0x84, 0x73, 0xf3, 0xd4, 0xf7, 0x55, 0x4d, 0x7d,
	0x55, 0x35, 0xdd, 0x65, 0x78, 0x7d, 0x32, 0x17, 0x64, 0x1a, 0x47, 0x01, 0x9f, 0x13, 0xc9, 0x42,
	0x9a, 0xf0, 0x15, 0x53, 0x64, 0x42, 0xe7, 0x03, 0x21, 0xe3, 0x24, 0x76, 0xbf, 0x9a, 0x72, 0x35,
	0x8d, 0x09, 0x8f, 0x15, 0x79, 0x21, 0x09, 0x17, 0xab, 0x11, 0xd1, 0x1e, 0xb1, 0x60, 0x72, 0x30,
	0x99, 0x8b, 0x41, 0xea, 0xca, 0x23, 0x95, 0xd0, 0x68, 0xca, 0x54, 0xd5, 0x60, 0x9f, 0x57, 0x32,
	0x50, 0x85, 0xdf, 0xe9, 0xcf, 0xa5, 0xa4, 0x09, 0x8f, 0x23, 0xb2, 0x54, 0x4c, 0xaa, 0x1a, 0x5b,
	0xff, 0x0f, 0x07, 0xee, 0x6c, 0xcb, 0x90, 0x7c, 0x72, 0xf1, 0xc5, 0x33, 0xf7, 0x2d, 0x68, 0xd9,
	0x77, 0x91, 0x88, 0x2e, 0x98, 0xe7, 0x9c, 0x39, 0xe7, 0x07, 0xe3, 0xa6, 0x35, 0x3e, 0xa5, 0x0b,
	0xe6, 0x9e, 0xc2, 0xad, 0x95, 0x0c, 0x0c, 0x7e, 0x03, 0xf1, 0x9b, 0x2b, 0x19, 0x20, 0xf4, 0x06,
	0x34, 0x58, 0x94, 0xf0, 0x64, 0x4d, 0x92, 0xb5, 0x60, 0xde, 0x1e, 0xa2, 0x60, 0x4c, 0xcf, 0xd7,
	0x82, 0xb9, 0x77, 0xa1, 0x1d, 0x31, 0x3e, 0xff, 0x72, 0x12, 0x4b, 0x42, 0x67, 0x33, 0xc9, 0x94,
	0xf2, 0x5e, 0x41, 0xd6, 0x91, 0xb5, 0x7f, 0x64, 0xcc, 0x85, 0x58, 0xf8, 0xa6, 0x57, 0x8b, 0xb1,
	0xf4, 0xcb, 0xfa, 0x0f, 0xa0, 0xad, 0xd5, 0x84, 0xc3, 0x95, 0x88, 0x30, 0x18, 0x49, 0xb4, 0x80,
	0xfc, 0x59, 0x07, 0x77, 0xce, 0xf6, 0xce, 0x5b, 0xe3, 0x26, 0x1a, 0xd3, 0xc8, 0xfd, 0x0f, 0xe1,
	0x38, 0x77, 0x5c, 0x28, 0x71, 0xf5, 0x9f, 0xbc, 0x9f, 0x82, 0xa7, 0xbd, 0xb1, 0x7b, 0x4a, 0x8a,
	0x38, 0xe4, 0xd3, 0xb5, 0x0d, 0x30, 0x84, 0xe3, 0x4d, 0x7b, 0x1e, 0xa8, 0xab, 0xc1, 0x67, 0x29,
	0xb6, 0x19, 0xcf, 0xdf, 0x12, 0xcf, 0xbf, 0x2e, 0x9e, 0x5f, 0x8d, 0xf7, 0xdd, 0x21, 0x34, 0x75,
	0x40, 0x4d, 0xd5, 0x5d, 0x70, 0xdb, 0xb0, 0x47, 0x03, 0x9e, 0xb6, 0x52, 0xff, 0x74, 0xdf, 0x84,
	0x26, 0xa6, 0x69, 0xa3, 0x99, 0x2e, 0x36, 0xb4, 0xcd, 0x56, 0xff, 0x1e, 0xb8, 0x48, 0x59, 0x4c,
	0xa9, 0x4a, 0x32, 0xa2, 0x69, 0x68, 0x5b, 0x23, 0x97, 0x1a, 0xa8, 0xb2, 0x43, 0x3a, 0x61, 0x61,
	0xa5, 0xb1, 0xc8, 0x7e, 0xa2, 0x01, 0xcb, 0x1e, 0x00, 0x16, 0x82, 0x24, 0xcb, 0x28, 0x2a, 0xd0,
	0x4d, 0x87, 0x3b, 0x1a, 0x7a, 0x8e, 0x88, 0xe5, 0x9f, 0x43, 0xdb, 0xe4, 0x32, 0xcb, 0x33, 0xd9,
	0x47, 0xf2, 0x21, 0x66, 0x32, 0xcb, 0xf2, 0x78, 0x1b, 0x8e, 0xb4, 0xa5, 0xd8, 0xc2, 0x9b, 0x39,
	0x31, 0x6f, 0xa2, 0xfb, 0x2e, 0xdc, 0x4e, 0x2d, 0xd4, 0x48, 0x4c, 0xd9, 0xb7, 0x90, 0xdd, 0xb5,
	0xd8, 0x65, 0x0e, 0xa5, 0x45, 0xf3, 0xb3, 0xc0, 0x07, 0x59, 0xd1, 0xfc, 0x72, 0x19, 0xfc, 0x4a,
	0xd1, 0x20, 0x2b, 0x83, 0x5f, 0x53, 0x34, 0xbf, 0x52, 0xb4, 0x46, 0xce, 0x2e, 0x15, 0xcd, 0x48,
	0xf3, 0x8b, 0xd2, 0x9a, 0x99, 0x34, 0xbf, 0x20, 0x2d, 0x9d, 0x19, 0x9c, 0xed, 0x52, 0x1e, 0xad,
	0x4c, 0x9b, 0x06, 0x4b, 0xa9, 0xfc, 0xe2, 0x80, 0x6b, 0x26, 0x7f, 0x25, 0x42, 0x95, 0x79, 0x1c,
	0x9e, 0x39, 0xe7, 0x8d, 0xe1, 0xb7, 0xce, 0xe0, 0x7f, 0x3d, 0xb7, 0x06, 0xd5, 0xaf, 0x7c, 0xdc,
	0xc6, 0xa7, 0xcf, 0x44, 0xa8, 0x0a, 0xa2, 0x65, 0xa2, 0x0f, 0x36, 0x95, 0x48, 0xca, 0xa3, 0x5c,
	0xf4, 0x91, 0x11, 0x2d, 0x93, 0x87, 0x19, 0x66, 0x7d, 0xee, 0xe2, 0x58, 0xf9, 0x8b, 0x62, 0x49,
	0xdb, 0xe6, 0x2c, 0xb2, 0xf6, 0x32, 0x75, 0x54, 0xa2, 0x76, 0x32, 0xea, 0xa8, 0x48, 0xbd, 0x67,
	0x2b, 0xc9, 0x8a, 0x64, 0xd7, 0x74, 0x15, 0x91, 0x8b, 0x02, 0xbb, 0x0f, 0xad, 0x50, 0x91, 0x42,
	0xc9, 0xbb, 0x66, 0xaa, 0x42, 0xf5, 0x24, 0xd3, 0xf6, 0x5b, 0xd6, 0x9c, 0xec, 0xac, 0xd2, 0xcc,
	0xdb, 0xd8, 0x9c, 0xef, 0x77, 0xa7, 0x39, 0x85, 0x93, 0x34, 0x55, 0x7a, 0xa9, 0xc4, 0x55, 0x79,
	0x2c, 0x47, 0x24, 0x08, 0xe3, 0x2b, 0x25, 0xd8, 0x34, 0xd3, 0x71, 0x9c, 0x7f, 0x72, 0x8f, 0x52,
	0xac, 0x32, 0xca, 0x9b, 0x3e, 0xbd, 0x7c, 0x94, 0xab, 0x3e, 0xef, 0x83, 0x67, 0x8f, 0x80, 0x0d,
	0xb7, 0x13, 0x74, 0xeb, 0xa5, 0x78, 0xbd, 0xa7, 0x5f, 0xeb, 0xe9, 0x65, 0x9e, 0x7e, 0x8d, 0xe7,
	0x9f, 0x0e, 0xf4, 0xd2, 0x73, 0x9f, 0x54, 0x0e, 0xea, 0x53, 0xec, 0xd2, 0x4f, 0xbb, 0xd0, 0xa5,
	0xba, 0x1b, 0xcb, 0x5e, 0x49, 0x9f, 0x16, 0xaf, 0x10, 0xab, 0xc7, 0xaf, 0xd1, 0xf3, 0xda, 0x4e,
	0xe9, 0xf1, 0x6b, 0xf5, 0xf8, 0x15, 0x3d, 0xfd, 0x7f, 0x6e, 0x40, 0xaf, 0xb0, 0xf8, 0xe8, 0x1d,
	0x62, 0x46, 0x08, 0x4f, 0xd8, 0xc2, 0xfd, 0x00, 0x4e, 0x53, 0x69, 0x24, 0xa0, 0x0b, 0x1e, 0xae,
	0x09, 0x9f, 0x69, 0x3c, 0xe0, 0x4c, 0xe2, 0x95, 0xd9, 0x1a, 0x9f, 0xa4, 0x84, 0x47, 0x88, 0x3f,
	0xce, 0x60, 0xf7, 0x1d, 0x70, 0xcb, 0xa9, 0xe1, 0xd2, 0x63, 0x2e, 0xd3, 0x4e, 0x09, 0xc1, 0xdd,
	0xe7, 0x67, 0x07, 0xba, 0xd5, 0xe5, 0x87, 0xbc, 0x90, 0x78, 0xa9, 0x36, 0x86, 0xdf, 0xec, 0x42,
	0x49, 0xed, 0xce, 0x30, 0xee, 0x54, 0xb6, 0xb1, 0xcf, 0xa5, 0x7b, 0x07, 0x60, 0x2e, 0xe3, 0xa5,
	0x30, 0xeb, 0x98, 0xb9, 0xdb, 0x0f, 0xd0, 0x82, 0xdb, 0xd8, 0xdf, 0x0e, 0x9c, 0x6c, 0xd6, 0x98,
	0x45, 0x89, 0x5c, 0xbb, 0xbf, 0x3a, 0xd0, 0xd9, 0xc0, 0x70, 0x87, 0x69, 0x0c, 0x7f, 0xd8, 0x05,
	0xdd, 0x9b, 0x83, 0x31, 0x3e, 0x9c, 0xcc, 0xc5, 0x43, 0x34, 0x5f, 0x68, 0x6b, 0xff, 0xeb, 0x7d,
	0xb3, 0xa7, 0xd5, 0x2d, 0xcf, 0x5b, 0x5b, 0x3b, 0x7c, 0x19, 0x5b, 0x7b, 0xbf, 0xd2, 0xda, 0x2d,
	0x73, 0x3e, 0xda, 0x36, 0xe7, 0xd7, 0x7e, 0x52, 0xef, 0x5d, 0xff, 0x49, 0xfd, 0xe5, 0x80, 0x47,
	0x03, 0xc2, 0xa3, 0x19, 0x13, 0x2c, 0xd2, 0xf6, 0xbc, 0xd2, 0x9e, 0x8f, 0xd5, 0xfc, 0x71, 0xf7,
	0x06, 0x06, 0xa7, 0x7c, 0xdc, 0xa3, 0xc1, 0xe3, 0x3c, 0xef, 0xb1, 0x4d, 0xdb, 0xfd, 0xdd, 0x81,
	0x63, 0x1a, 0x90, 0x4d, 0x45, 0xde, 0x03, 0xfc, 0x02, 0x76, 0x56, 0x50, 0x97, 0x06, 0x1f, 0x57,
	0xe5, 0x4c, 0xf6, 0xf1, 0xaf, 0xec, 0xfd, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x33, 0xaf,
	0x2c, 0xec, 0x0e, 0x00, 0x00,
}
