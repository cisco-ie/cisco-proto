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
// source: bgp_static_routes_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_sourced_networks_sourced_network

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

type BgpStaticRoutesBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Network              string   `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpStaticRoutesBag_KEYS) Reset()         { *m = BgpStaticRoutesBag_KEYS{} }
func (m *BgpStaticRoutesBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpStaticRoutesBag_KEYS) ProtoMessage()    {}
func (*BgpStaticRoutesBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f6a59aaa51735, []int{0}
}

func (m *BgpStaticRoutesBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpStaticRoutesBag_KEYS.Unmarshal(m, b)
}
func (m *BgpStaticRoutesBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpStaticRoutesBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpStaticRoutesBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpStaticRoutesBag_KEYS.Merge(m, src)
}
func (m *BgpStaticRoutesBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpStaticRoutesBag_KEYS.Size(m)
}
func (m *BgpStaticRoutesBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpStaticRoutesBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpStaticRoutesBag_KEYS proto.InternalMessageInfo

func (m *BgpStaticRoutesBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpStaticRoutesBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpStaticRoutesBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpStaticRoutesBag_KEYS) GetNetwork() string {
	if m != nil {
		return m.Network
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
	return fileDescriptor_681f6a59aaa51735, []int{1}
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
	return fileDescriptor_681f6a59aaa51735, []int{2}
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
	return fileDescriptor_681f6a59aaa51735, []int{3}
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
	return fileDescriptor_681f6a59aaa51735, []int{4}
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
	Ipv4MdtAddress         string                `protobuf:"bytes,6,opt,name=ipv4mdt_address,json=ipv4mdtAddress,proto3" json:"ipv4mdt_address,omitempty"`
	Ipv4VpnAddress         string                `protobuf:"bytes,7,opt,name=ipv4vpn_address,json=ipv4vpnAddress,proto3" json:"ipv4vpn_address,omitempty"`
	Ipv4VpnaMcastddress    string                `protobuf:"bytes,8,opt,name=ipv4vpna_mcastddress,json=ipv4vpnaMcastddress,proto3" json:"ipv4vpna_mcastddress,omitempty"`
	Ipv6Address            string                `protobuf:"bytes,9,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	Ipv6McastAddress       string                `protobuf:"bytes,10,opt,name=ipv6_mcast_address,json=ipv6McastAddress,proto3" json:"ipv6_mcast_address,omitempty"`
	Ipv6LabelAddress       string                `protobuf:"bytes,11,opt,name=ipv6_label_address,json=ipv6LabelAddress,proto3" json:"ipv6_label_address,omitempty"`
	Ipv6VpnAddress         string                `protobuf:"bytes,12,opt,name=ipv6vpn_address,json=ipv6vpnAddress,proto3" json:"ipv6vpn_address,omitempty"`
	Ipv6VpnMcastAddress    string                `protobuf:"bytes,13,opt,name=ipv6vpn_mcast_address,json=ipv6vpnMcastAddress,proto3" json:"ipv6vpn_mcast_address,omitempty"`
	L2VpnvplsAddress       *BgpL2VpnAddrT        `protobuf:"bytes,14,opt,name=l2vpnvpls_address,json=l2vpnvplsAddress,proto3" json:"l2vpnvpls_address,omitempty"`
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
	Ipv4SrPolicyAddress    *BgpIpv4SrpolicyAddrT `protobuf:"bytes,25,opt,name=ipv4sr_policy_address,json=ipv4srPolicyAddress,proto3" json:"ipv4sr_policy_address,omitempty"`
	Ipv6SrPolicyAddress    *BgpIpv6SrpolicyAddrT `protobuf:"bytes,26,opt,name=ipv6sr_policy_address,json=ipv6srPolicyAddress,proto3" json:"ipv6sr_policy_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}              `json:"-"`
	XXX_unrecognized       []byte                `json:"-"`
	XXX_sizecache          int32                 `json:"-"`
}

func (m *BgpAddrtype) Reset()         { *m = BgpAddrtype{} }
func (m *BgpAddrtype) String() string { return proto.CompactTextString(m) }
func (*BgpAddrtype) ProtoMessage()    {}
func (*BgpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f6a59aaa51735, []int{5}
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

func (m *BgpAddrtype) GetL2VpnvplsAddress() *BgpL2VpnAddrT {
	if m != nil {
		return m.L2VpnvplsAddress
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

type BgpPrefixtype struct {
	Prefix               *BgpAddrtype `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLength         uint32       `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BgpPrefixtype) Reset()         { *m = BgpPrefixtype{} }
func (m *BgpPrefixtype) String() string { return proto.CompactTextString(m) }
func (*BgpPrefixtype) ProtoMessage()    {}
func (*BgpPrefixtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f6a59aaa51735, []int{6}
}

func (m *BgpPrefixtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPrefixtype.Unmarshal(m, b)
}
func (m *BgpPrefixtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPrefixtype.Marshal(b, m, deterministic)
}
func (m *BgpPrefixtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPrefixtype.Merge(m, src)
}
func (m *BgpPrefixtype) XXX_Size() int {
	return xxx_messageInfo_BgpPrefixtype.Size(m)
}
func (m *BgpPrefixtype) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPrefixtype.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPrefixtype proto.InternalMessageInfo

func (m *BgpPrefixtype) GetPrefix() *BgpAddrtype {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *BgpPrefixtype) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type BgpStaticRoutesBag struct {
	AfName               string         `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SourcedPrefix        *BgpPrefixtype `protobuf:"bytes,51,opt,name=sourced_prefix,json=sourcedPrefix,proto3" json:"sourced_prefix,omitempty"`
	IsBackdoor           bool           `protobuf:"varint,52,opt,name=is_backdoor,json=isBackdoor,proto3" json:"is_backdoor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpStaticRoutesBag) Reset()         { *m = BgpStaticRoutesBag{} }
func (m *BgpStaticRoutesBag) String() string { return proto.CompactTextString(m) }
func (*BgpStaticRoutesBag) ProtoMessage()    {}
func (*BgpStaticRoutesBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_681f6a59aaa51735, []int{7}
}

func (m *BgpStaticRoutesBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpStaticRoutesBag.Unmarshal(m, b)
}
func (m *BgpStaticRoutesBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpStaticRoutesBag.Marshal(b, m, deterministic)
}
func (m *BgpStaticRoutesBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpStaticRoutesBag.Merge(m, src)
}
func (m *BgpStaticRoutesBag) XXX_Size() int {
	return xxx_messageInfo_BgpStaticRoutesBag.Size(m)
}
func (m *BgpStaticRoutesBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpStaticRoutesBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpStaticRoutesBag proto.InternalMessageInfo

func (m *BgpStaticRoutesBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpStaticRoutesBag) GetSourcedPrefix() *BgpPrefixtype {
	if m != nil {
		return m.SourcedPrefix
	}
	return nil
}

func (m *BgpStaticRoutesBag) GetIsBackdoor() bool {
	if m != nil {
		return m.IsBackdoor
	}
	return false
}

func init() {
	proto.RegisterType((*BgpStaticRoutesBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_static_routes_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_addrtype")
	proto.RegisterType((*BgpPrefixtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_prefixtype")
	proto.RegisterType((*BgpStaticRoutesBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.sourced_networks.sourced_network.bgp_static_routes_bag")
}

func init() { proto.RegisterFile("bgp_static_routes_bag.proto", fileDescriptor_681f6a59aaa51735) }

var fileDescriptor_681f6a59aaa51735 = []byte{
	// 816 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xbb, 0x6e, 0x13, 0x4d,
	0x14, 0xd6, 0xfe, 0xc9, 0x6f, 0x27, 0xe3, 0x4b, 0x9c, 0xcd, 0x6d, 0x13, 0x0a, 0x8c, 0x29, 0x30,
	0x52, 0xb4, 0x12, 0x8e, 0x35, 0x50, 0xd0, 0x00, 0x0a, 0x0d, 0x49, 0x14, 0x39, 0x34, 0x54, 0xa3,
	0xf1, 0x7a, 0xd6, 0xac, 0xb2, 0xde, 0x5d, 0xcd, 0x4c, 0xec, 0xa4, 0x43, 0x88, 0x06, 0x09, 0x89,
	0x9a, 0x57, 0xa0, 0xa4, 0xa2, 0xe5, 0x11, 0x78, 0x03, 0x9e, 0x04, 0xa1, 0xb9, 0xec, 0xec, 0xc5,
	0x06, 0x89, 0x06, 0x37, 0xd1, 0xf8, 0x7c, 0xdf, 0x39, 0x39, 0xdf, 0x77, 0x66, 0x67, 0x06, 0xdc,
	0x1a, 0x8e, 0x13, 0xc4, 0x38, 0xe6, 0x81, 0x87, 0x68, 0x7c, 0xc5, 0x09, 0x43, 0x43, 0x3c, 0x76,
	0x13, 0x1a, 0xf3, 0xd8, 0xa6, 0x5e, 0xc0, 0xbc, 0x18, 0x05, 0x31, 0x43, 0xd7, 0x14, 0x05, 0xc9,
	0xb4, 0x8f, 0x04, 0x3d, 0x4e, 0x08, 0x75, 0x87, 0xe3, 0xc4, 0x0d, 0x22, 0xc6, 0x71, 0xe4, 0x11,
	0x66, 0x56, 0x66, 0x81, 0xb0, 0xc7, 0x83, 0x29, 0x71, 0xa7, 0xd4, 0x67, 0xe2, 0x8f, 0x8b, 0x7d,
	0xe6, 0x62, 0xdf, 0x65, 0xf1, 0x15, 0xf5, 0xc8, 0x08, 0x45, 0x84, 0xcf, 0x62, 0x7a, 0xc9, 0xca,
	0x81, 0xce, 0x47, 0x0b, 0x1c, 0x2c, 0xec, 0x09, 0xbd, 0x38, 0x7e, 0x75, 0x61, 0xdf, 0x05, 0x0d,
	0xf3, 0x2f, 0x22, 0x3c, 0x21, 0x8e, 0xd5, 0xb6, 0xba, 0xeb, 0x83, 0x7a, 0x1a, 0x3c, 0xc3, 0x13,
	0x62, 0xef, 0x83, 0xb5, 0x29, 0xf5, 0x15, 0xfe, 0x9f, 0xc4, 0xab, 0x53, 0xea, 0x4b, 0x68, 0x0f,
	0x54, 0xb1, 0x46, 0x56, 0x24, 0x52, 0xc1, 0x0a, 0x70, 0x40, 0x55, 0xb7, 0xe0, 0xac, 0xaa, 0x94,
	0xb4, 0xa3, 0x87, 0xa0, 0x25, 0x1a, 0x0a, 0x7b, 0xd3, 0x24, 0x42, 0x78, 0x34, 0xa2, 0x88, 0x8b,
	0x36, 0xb2, 0xdf, 0x84, 0x31, 0xc7, 0x6a, 0xaf, 0x74, 0x1b, 0x83, 0xba, 0x0c, 0x3e, 0x51, 0xb1,
	0xce, 0x63, 0xb0, 0x93, 0x25, 0x4e, 0x58, 0x32, 0xfb, 0xab, 0xec, 0x33, 0xe0, 0x88, 0x6c, 0xe9,
	0x3a, 0xa3, 0x49, 0x1c, 0x06, 0xde, 0x4d, 0x5a, 0xa0, 0x07, 0x76, 0xe6, 0xe3, 0x59, 0xa1, 0x2d,
	0x01, 0x5e, 0x68, 0x6c, 0xbe, 0x1e, 0xfc, 0x4d, 0x3d, 0xf8, 0xa7, 0x7a, 0xb0, 0x5c, 0xef, 0x47,
	0x03, 0xd4, 0x45, 0x41, 0x41, 0xe5, 0x37, 0x09, 0xb1, 0x5b, 0x60, 0x05, 0xfb, 0x81, 0x1e, 0x88,
	0x58, 0xda, 0x77, 0x40, 0x5d, 0xb6, 0x99, 0x56, 0x53, 0xb3, 0xa8, 0x89, 0x98, 0xae, 0x62, 0x1f,
	0x02, 0x5b, 0x52, 0x26, 0x1e, 0x66, 0xdc, 0x10, 0xd5, 0x68, 0x5a, 0x02, 0x39, 0x15, 0x40, 0x99,
	0x1d, 0xe2, 0x21, 0x09, 0x0d, 0x7b, 0x35, 0x63, 0x9f, 0x08, 0x20, 0x65, 0xbb, 0x40, 0x1a, 0x81,
	0xf8, 0x55, 0x14, 0xe5, 0xe8, 0xff, 0x4b, 0xfa, 0xa6, 0x80, 0x5e, 0x4a, 0x24, 0xe5, 0xdf, 0x03,
	0x1b, 0x22, 0x38, 0x19, 0x65, 0x8d, 0x54, 0x24, 0xb7, 0xa9, 0xc3, 0x25, 0x62, 0x7e, 0x82, 0xd5,
	0x8c, 0x98, 0xcd, 0xd0, 0x7e, 0x00, 0xb6, 0x75, 0x04, 0x2b, 0x85, 0x9a, 0xbd, 0x26, 0xd9, 0x5b,
	0x29, 0x76, 0x9a, 0x41, 0xda, 0x33, 0x68, 0x0a, 0xaf, 0x1b, 0xcf, 0x60, 0xd1, 0x05, 0x58, 0xf2,
	0x0c, 0x18, 0x17, 0xe0, 0x02, 0xcf, 0x60, 0xc9, 0xb3, 0x5a, 0xc6, 0x2e, 0x78, 0xa6, 0xa4, 0xc1,
	0xbc, 0xb4, 0xba, 0x91, 0x06, 0x73, 0xd2, 0xf4, 0x96, 0x91, 0x5b, 0xbb, 0xd0, 0x47, 0xc3, 0x68,
	0x13, 0x60, 0xa1, 0x95, 0x2f, 0x16, 0xd8, 0x94, 0x7b, 0x7c, 0x9a, 0x84, 0xcc, 0x24, 0x34, 0xdb,
	0x56, 0xb7, 0xd6, 0x7b, 0x67, 0xb9, 0xff, 0xfe, 0xb4, 0x71, 0xcb, 0x1f, 0xf6, 0xa0, 0x65, 0xfa,
	0xcb, 0x09, 0xa5, 0x1c, 0x79, 0x71, 0xc4, 0x38, 0xc5, 0x41, 0x94, 0x09, 0xdd, 0x50, 0x42, 0x29,
	0x7f, 0x66, 0xb0, 0x34, 0xe7, 0x3e, 0x90, 0xce, 0x4e, 0xf2, 0x36, 0xb6, 0x24, 0x7d, 0x23, 0x8d,
	0x17, 0xa9, 0xfd, 0x02, 0x75, 0xd3, 0x50, 0xfb, 0x79, 0xea, 0x21, 0xb0, 0x55, 0xaf, 0x24, 0x4f,
	0xb6, 0xd5, 0x24, 0x25, 0x72, 0x9c, 0x63, 0x77, 0x40, 0x23, 0x64, 0x28, 0xe7, 0xf3, 0x96, 0xda,
	0x49, 0x21, 0x3b, 0x31, 0xda, 0xbe, 0x5a, 0x69, 0x49, 0x73, 0x3c, 0x09, 0xe6, 0xb6, 0x9c, 0xc8,
	0xfb, 0x25, 0x4f, 0x24, 0x77, 0x62, 0x6a, 0x79, 0xa7, 0x2c, 0x99, 0x15, 0xf7, 0x5f, 0x1f, 0xf9,
	0x61, 0x3c, 0x63, 0x09, 0xf1, 0x4c, 0xf3, 0x3b, 0xd9, 0xb7, 0xf5, 0x5c, 0x63, 0xa5, 0x3d, 0x3b,
	0x9f, 0xb3, 0x9b, 0xed, 0xd9, 0x72, 0xce, 0x23, 0xe0, 0xa4, 0xdf, 0xfa, 0x5c, 0xda, 0x9e, 0x4c,
	0xdb, 0xd5, 0xf8, 0xe2, 0x4c, 0xb8, 0x30, 0xd3, 0x31, 0x99, 0x70, 0x41, 0xe6, 0x37, 0x4b, 0x89,
	0x63, 0x14, 0x95, 0xce, 0xe3, 0x7d, 0x39, 0x99, 0x0f, 0x4b, 0x9b, 0xcc, 0xa2, 0xdb, 0x48, 0x79,
	0xcd, 0xe8, 0x79, 0xfe, 0x7a, 0x48, 0x35, 0xc0, 0x79, 0x0d, 0x07, 0xcb, 0xd7, 0x00, 0x17, 0x6a,
	0x80, 0x25, 0x0d, 0x9d, 0xef, 0x16, 0x68, 0x8a, 0x8c, 0x84, 0x12, 0x3f, 0xb8, 0x96, 0x97, 0xdc,
	0x27, 0x0b, 0x54, 0xd4, 0x4f, 0x79, 0xd1, 0xd5, 0x7a, 0x6f, 0x96, 0xa6, 0x23, 0xbd, 0x78, 0x07,
	0xba, 0x21, 0xf1, 0xac, 0x50, 0x2b, 0x14, 0x92, 0x68, 0xcc, 0x5f, 0xcb, 0xfb, 0xb6, 0x31, 0xa8,
	0xab, 0xe0, 0x89, 0x8c, 0x75, 0x7e, 0x5a, 0xea, 0x55, 0x32, 0xf7, 0xbe, 0xca, 0x3f, 0x8d, 0x7a,
	0x85, 0xa7, 0xd1, 0x67, 0x0b, 0x34, 0xd3, 0x06, 0xb4, 0xf6, 0x23, 0xa9, 0xfd, 0xed, 0xd2, 0xb4,
	0x67, 0x13, 0x19, 0x34, 0x34, 0x7e, 0xae, 0x4c, 0xb8, 0x0d, 0x6a, 0x81, 0xd0, 0xe3, 0x5d, 0x8e,
	0xe2, 0x98, 0x3a, 0xfd, 0xb6, 0xd5, 0x5d, 0x1b, 0x80, 0x80, 0x3d, 0xd5, 0x91, 0x61, 0x45, 0xbe,
	0x6d, 0x8f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x19, 0x61, 0xed, 0xfa, 0x0a, 0x00, 0x00,
}
