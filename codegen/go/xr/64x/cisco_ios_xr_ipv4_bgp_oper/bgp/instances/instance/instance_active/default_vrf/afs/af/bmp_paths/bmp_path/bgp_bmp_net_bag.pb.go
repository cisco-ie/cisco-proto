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
// source: bgp_bmp_net_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_bmp_paths_bmp_path

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

type BgpBmpNetBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Network              string   `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,4,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpBmpNetBag_KEYS) Reset()         { *m = BgpBmpNetBag_KEYS{} }
func (m *BgpBmpNetBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpBmpNetBag_KEYS) ProtoMessage()    {}
func (*BgpBmpNetBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b62712d7052d9c5, []int{0}
}

func (m *BgpBmpNetBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpBmpNetBag_KEYS.Unmarshal(m, b)
}
func (m *BgpBmpNetBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpBmpNetBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpBmpNetBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpBmpNetBag_KEYS.Merge(m, src)
}
func (m *BgpBmpNetBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpBmpNetBag_KEYS.Size(m)
}
func (m *BgpBmpNetBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpBmpNetBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpBmpNetBag_KEYS proto.InternalMessageInfo

func (m *BgpBmpNetBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpBmpNetBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpBmpNetBag_KEYS) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *BgpBmpNetBag_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
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
	return fileDescriptor_3b62712d7052d9c5, []int{1}
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
	return fileDescriptor_3b62712d7052d9c5, []int{2}
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
	return fileDescriptor_3b62712d7052d9c5, []int{3}
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
	return fileDescriptor_3b62712d7052d9c5, []int{4}
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
	return fileDescriptor_3b62712d7052d9c5, []int{5}
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
	return fileDescriptor_3b62712d7052d9c5, []int{6}
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

type BgpTimespec struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanoseconds          uint32   `protobuf:"varint,2,opt,name=nanoseconds,proto3" json:"nanoseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpTimespec) Reset()         { *m = BgpTimespec{} }
func (m *BgpTimespec) String() string { return proto.CompactTextString(m) }
func (*BgpTimespec) ProtoMessage()    {}
func (*BgpTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b62712d7052d9c5, []int{7}
}

func (m *BgpTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpTimespec.Unmarshal(m, b)
}
func (m *BgpTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpTimespec.Marshal(b, m, deterministic)
}
func (m *BgpTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpTimespec.Merge(m, src)
}
func (m *BgpTimespec) XXX_Size() int {
	return xxx_messageInfo_BgpTimespec.Size(m)
}
func (m *BgpTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_BgpTimespec proto.InternalMessageInfo

func (m *BgpTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *BgpTimespec) GetNanoseconds() uint32 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

type BgpBmpPathBag struct {
	NeighborAddress      *BgpAddrtype `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	BmpPathFlags         uint32       `protobuf:"varint,2,opt,name=bmp_path_flags,json=bmpPathFlags,proto3" json:"bmp_path_flags,omitempty"`
	BmpPathAdvBitMap     uint32       `protobuf:"varint,3,opt,name=bmp_path_adv_bit_map,json=bmpPathAdvBitMap,proto3" json:"bmp_path_adv_bit_map,omitempty"`
	BmpPathSndBitMap     uint32       `protobuf:"varint,4,opt,name=bmp_path_snd_bit_map,json=bmpPathSndBitMap,proto3" json:"bmp_path_snd_bit_map,omitempty"`
	BmpNbrBitMap         uint32       `protobuf:"varint,5,opt,name=bmp_nbr_bit_map,json=bmpNbrBitMap,proto3" json:"bmp_nbr_bit_map,omitempty"`
	LocalPath            bool         `protobuf:"varint,6,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	ReceivedLabel        []uint32     `protobuf:"varint,7,rep,packed,name=received_label,json=receivedLabel,proto3" json:"received_label,omitempty"`
	BpathPointer         uint64       `protobuf:"varint,8,opt,name=bpath_pointer,json=bpathPointer,proto3" json:"bpath_pointer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BgpBmpPathBag) Reset()         { *m = BgpBmpPathBag{} }
func (m *BgpBmpPathBag) String() string { return proto.CompactTextString(m) }
func (*BgpBmpPathBag) ProtoMessage()    {}
func (*BgpBmpPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b62712d7052d9c5, []int{8}
}

func (m *BgpBmpPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpBmpPathBag.Unmarshal(m, b)
}
func (m *BgpBmpPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpBmpPathBag.Marshal(b, m, deterministic)
}
func (m *BgpBmpPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpBmpPathBag.Merge(m, src)
}
func (m *BgpBmpPathBag) XXX_Size() int {
	return xxx_messageInfo_BgpBmpPathBag.Size(m)
}
func (m *BgpBmpPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpBmpPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpBmpPathBag proto.InternalMessageInfo

func (m *BgpBmpPathBag) GetNeighborAddress() *BgpAddrtype {
	if m != nil {
		return m.NeighborAddress
	}
	return nil
}

func (m *BgpBmpPathBag) GetBmpPathFlags() uint32 {
	if m != nil {
		return m.BmpPathFlags
	}
	return 0
}

func (m *BgpBmpPathBag) GetBmpPathAdvBitMap() uint32 {
	if m != nil {
		return m.BmpPathAdvBitMap
	}
	return 0
}

func (m *BgpBmpPathBag) GetBmpPathSndBitMap() uint32 {
	if m != nil {
		return m.BmpPathSndBitMap
	}
	return 0
}

func (m *BgpBmpPathBag) GetBmpNbrBitMap() uint32 {
	if m != nil {
		return m.BmpNbrBitMap
	}
	return 0
}

func (m *BgpBmpPathBag) GetLocalPath() bool {
	if m != nil {
		return m.LocalPath
	}
	return false
}

func (m *BgpBmpPathBag) GetReceivedLabel() []uint32 {
	if m != nil {
		return m.ReceivedLabel
	}
	return nil
}

func (m *BgpBmpPathBag) GetBpathPointer() uint64 {
	if m != nil {
		return m.BpathPointer
	}
	return 0
}

type BgpBmpNetBag struct {
	AfName               string           `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	BgpPrefix            *BgpPrefixtype   `protobuf:"bytes,51,opt,name=bgp_prefix,json=bgpPrefix,proto3" json:"bgp_prefix,omitempty"`
	BmpNetFlags          uint32           `protobuf:"varint,52,opt,name=bmp_net_flags,json=bmpNetFlags,proto3" json:"bmp_net_flags,omitempty"`
	BmpNetVersion        uint32           `protobuf:"varint,53,opt,name=bmp_net_version,json=bmpNetVersion,proto3" json:"bmp_net_version,omitempty"`
	NumOfPath            uint32           `protobuf:"varint,54,opt,name=num_of_path,json=numOfPath,proto3" json:"num_of_path,omitempty"`
	VersionTimestamp     *BgpTimespec     `protobuf:"bytes,55,opt,name=version_timestamp,json=versionTimestamp,proto3" json:"version_timestamp,omitempty"`
	VersionAge           *BgpTimespec     `protobuf:"bytes,56,opt,name=version_age,json=versionAge,proto3" json:"version_age,omitempty"`
	HasLocalLabel        bool             `protobuf:"varint,57,opt,name=has_local_label,json=hasLocalLabel,proto3" json:"has_local_label,omitempty"`
	NetLocalLabel        uint32           `protobuf:"varint,58,opt,name=net_local_label,json=netLocalLabel,proto3" json:"net_local_label,omitempty"`
	BmpPath              []*BgpBmpPathBag `protobuf:"bytes,59,rep,name=bmp_path,json=bmpPath,proto3" json:"bmp_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BgpBmpNetBag) Reset()         { *m = BgpBmpNetBag{} }
func (m *BgpBmpNetBag) String() string { return proto.CompactTextString(m) }
func (*BgpBmpNetBag) ProtoMessage()    {}
func (*BgpBmpNetBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b62712d7052d9c5, []int{9}
}

func (m *BgpBmpNetBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpBmpNetBag.Unmarshal(m, b)
}
func (m *BgpBmpNetBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpBmpNetBag.Marshal(b, m, deterministic)
}
func (m *BgpBmpNetBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpBmpNetBag.Merge(m, src)
}
func (m *BgpBmpNetBag) XXX_Size() int {
	return xxx_messageInfo_BgpBmpNetBag.Size(m)
}
func (m *BgpBmpNetBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpBmpNetBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpBmpNetBag proto.InternalMessageInfo

func (m *BgpBmpNetBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpBmpNetBag) GetBgpPrefix() *BgpPrefixtype {
	if m != nil {
		return m.BgpPrefix
	}
	return nil
}

func (m *BgpBmpNetBag) GetBmpNetFlags() uint32 {
	if m != nil {
		return m.BmpNetFlags
	}
	return 0
}

func (m *BgpBmpNetBag) GetBmpNetVersion() uint32 {
	if m != nil {
		return m.BmpNetVersion
	}
	return 0
}

func (m *BgpBmpNetBag) GetNumOfPath() uint32 {
	if m != nil {
		return m.NumOfPath
	}
	return 0
}

func (m *BgpBmpNetBag) GetVersionTimestamp() *BgpTimespec {
	if m != nil {
		return m.VersionTimestamp
	}
	return nil
}

func (m *BgpBmpNetBag) GetVersionAge() *BgpTimespec {
	if m != nil {
		return m.VersionAge
	}
	return nil
}

func (m *BgpBmpNetBag) GetHasLocalLabel() bool {
	if m != nil {
		return m.HasLocalLabel
	}
	return false
}

func (m *BgpBmpNetBag) GetNetLocalLabel() uint32 {
	if m != nil {
		return m.NetLocalLabel
	}
	return 0
}

func (m *BgpBmpNetBag) GetBmpPath() []*BgpBmpPathBag {
	if m != nil {
		return m.BmpPath
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpBmpNetBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_bmp_net_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_addrtype")
	proto.RegisterType((*BgpPrefixtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_prefixtype")
	proto.RegisterType((*BgpTimespec)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_timespec")
	proto.RegisterType((*BgpBmpPathBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_bmp_path_bag")
	proto.RegisterType((*BgpBmpNetBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.bmp_paths.bmp_path.bgp_bmp_net_bag")
}

func init() { proto.RegisterFile("bgp_bmp_net_bag.proto", fileDescriptor_3b62712d7052d9c5) }

var fileDescriptor_3b62712d7052d9c5 = []byte{
	// 1135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4b, 0x6f, 0xe4, 0xc4,
	0x13, 0x97, 0x37, 0xd9, 0x3c, 0x7a, 0x66, 0x32, 0x13, 0xe7, 0xb1, 0xfe, 0xff, 0x25, 0xd0, 0x30,
	0xb0, 0x4b, 0x90, 0x56, 0x96, 0xc8, 0x06, 0xef, 0xf2, 0xb8, 0x04, 0xb4, 0x7b, 0x80, 0x24, 0x44,
	0xce, 0x6a, 0x25, 0x4e, 0xad, 0xb6, 0xdd, 0x9e, 0xb1, 0xb0, 0xdb, 0x2d, 0x77, 0x67, 0xb2, 0xb9,
	0x20, 0xf1, 0x01, 0x38, 0x81, 0x38, 0x23, 0x2e, 0x5c, 0x38, 0xc3, 0x27, 0xe0, 0xc8, 0x8d, 0x0f,
	0x84, 0xba, 0xda, 0xdd, 0xf6, 0x38, 0x03, 0x12, 0x97, 0xe4, 0xe6, 0xae, 0xfa, 0x55, 0xcd, 0xaf,
	0x1e, 0xdd, 0x55, 0x83, 0xf6, 0xa2, 0x29, 0xc7, 0x51, 0xc1, 0x31, 0xa3, 0x12, 0x47, 0x64, 0xea,
	0xf3, 0xaa, 0x94, 0xa5, 0x3b, 0x8d, 0x33, 0x11, 0x97, 0x38, 0x2b, 0x05, 0x7e, 0x5d, 0xe1, 0x8c,
	0xcf, 0x8f, 0xb0, 0x02, 0x96, 0x9c, 0x56, 0x7e, 0x34, 0xe5, 0x7e, 0xc6, 0x84, 0x24, 0x2c, 0xa6,
	0xc2, 0x7e, 0xd9, 0x0f, 0x4c, 0x62, 0x99, 0xcd, 0xa9, 0x9f, 0xd0, 0x94, 0x5c, 0xe6, 0x12, 0xcf,
	0xab, 0xd4, 0x27, 0xa9, 0xf0, 0x49, 0xea, 0xab, 0xdf, 0xe1, 0x44, 0xce, 0x84, 0xfd, 0x9a, 0x7c,
	0xef, 0xa0, 0xdd, 0x0e, 0x05, 0xfc, 0xc5, 0xf3, 0xaf, 0x2e, 0xdc, 0xb7, 0xd1, 0xc0, 0x7a, 0x64,
	0xa4, 0xa0, 0x9e, 0x33, 0x76, 0x0e, 0x36, 0xc3, 0xbe, 0x11, 0x9e, 0x91, 0x82, 0xba, 0x0f, 0xd0,
	0x3a, 0x49, 0xb5, 0xfa, 0x1e, 0xa8, 0xd7, 0x48, 0x0a, 0x0a, 0x0f, 0xad, 0x33, 0x2a, 0xaf, 0xca,
	0xea, 0x6b, 0x6f, 0x05, 0x14, 0xe6, 0xa8, 0xfc, 0xf2, 0x8a, 0xa6, 0xd9, 0x6b, 0x9c, 0x53, 0x36,
	0x95, 0x33, 0x6f, 0x75, 0xec, 0x1c, 0x0c, 0xc2, 0xbe, 0x16, 0x9e, 0x80, 0x6c, 0xf2, 0x14, 0x8d,
	0x14, 0xa9, 0xfc, 0x70, 0xce, 0x19, 0x26, 0x49, 0x52, 0x61, 0xa9, 0x0c, 0x9b, 0x33, 0x15, 0xc2,
	0x73, 0xc6, 0x2b, 0xca, 0x10, 0x84, 0xc7, 0x5a, 0x36, 0xf9, 0x44, 0x27, 0x54, 0x03, 0x0b, 0xc1,
	0xaf, 0xfe, 0x93, 0xf5, 0x19, 0xf2, 0x94, 0x35, 0xa4, 0x5b, 0x54, 0xbc, 0xcc, 0xb3, 0xf8, 0xda,
	0x38, 0x38, 0x44, 0x7b, 0x37, 0xe5, 0x8d, 0xa3, 0x1d, 0xa5, 0xbc, 0xa8, 0x75, 0x37, 0xfd, 0x05,
	0xff, 0xe0, 0x2f, 0xf8, 0x37, 0x7f, 0x41, 0xd7, 0xdf, 0x6f, 0x03, 0xd4, 0x57, 0x0e, 0x15, 0x54,
	0x5e, 0x73, 0xea, 0x8e, 0xd0, 0x0a, 0x49, 0xb3, 0xba, 0x34, 0xea, 0xd3, 0x7d, 0x0b, 0xf5, 0x81,
	0xa6, 0xf1, 0xa6, 0xcb, 0xd2, 0x53, 0xb2, 0xda, 0x8b, 0xfb, 0x18, 0xb9, 0x00, 0x29, 0x62, 0x22,
	0xa4, 0x05, 0xea, 0x32, 0x8d, 0x94, 0xe6, 0x54, 0x29, 0xba, 0xe8, 0x9c, 0x44, 0x34, 0xb7, 0xe8,
	0xd5, 0x06, 0x7d, 0xa2, 0x14, 0x06, 0xed, 0x23, 0x48, 0x04, 0x96, 0x97, 0x8c, 0xb5, 0xe0, 0xf7,
	0x01, 0xbe, 0xad, 0x54, 0x2f, 0x41, 0x63, 0xf0, 0x07, 0x68, 0xa4, 0xb9, 0x24, 0x0d, 0x93, 0x35,
	0x00, 0x6f, 0x01, 0x93, 0xc4, 0xf2, 0x78, 0x17, 0x0d, 0x95, 0xa4, 0x5d, 0xc2, 0xf5, 0x06, 0xd8,
	0x14, 0xd1, 0x7d, 0x1f, 0xed, 0xd6, 0x12, 0xa2, 0x43, 0xac, 0xd1, 0x1b, 0x80, 0xde, 0x31, 0xba,
	0xd3, 0x46, 0x55, 0x27, 0x2d, 0xb0, 0x8e, 0x37, 0x6d, 0xd2, 0x82, 0xc5, 0x34, 0x04, 0x9d, 0xa4,
	0x21, 0x9b, 0x86, 0x60, 0x49, 0xd2, 0x82, 0x4e, 0xd2, 0x7a, 0x0d, 0x7a, 0x21, 0x69, 0x3a, 0xb4,
	0xa0, 0x1d, 0x5a, 0xdf, 0x86, 0x16, 0xb4, 0x42, 0xab, 0x7b, 0x06, 0x7a, 0x7b, 0x81, 0xc7, 0xc0,
	0xc6, 0xa6, 0x94, 0x0b, 0x54, 0x7e, 0x71, 0x90, 0xab, 0x3b, 0x7f, 0xce, 0x73, 0x61, 0x2d, 0xb6,
	0xc6, 0xce, 0x41, 0xef, 0xf0, 0xda, 0xbf, 0xa5, 0x77, 0xc6, 0xef, 0x5e, 0xe7, 0x70, 0x04, 0xa7,
	0x57, 0x3c, 0x17, 0xad, 0xe8, 0x2a, 0x89, 0xe3, 0x92, 0x09, 0x59, 0x91, 0x8c, 0x35, 0xd1, 0x0d,
	0x75, 0x74, 0x95, 0xfc, 0xcc, 0xea, 0x8c, 0xcd, 0x7b, 0xd0, 0x3f, 0x41, 0xd1, 0xce, 0xdd, 0x08,
	0xe0, 0x43, 0x23, 0x5f, 0x84, 0x1e, 0x2d, 0x40, 0xb7, 0x2d, 0xf4, 0xa8, 0x0d, 0x7d, 0x6c, 0x52,
	0x46, 0xdb, 0x60, 0x57, 0x97, 0x0f, 0x34, 0xcf, 0x5b, 0xe8, 0x09, 0x1a, 0xe4, 0x02, 0xb7, 0x72,
	0xbb, 0xa3, 0xdb, 0x27, 0x17, 0x27, 0x36, 0xb6, 0x5f, 0x6d, 0x15, 0xec, 0xa3, 0xa4, 0x90, 0xbb,
	0x50, 0x85, 0x6f, 0xee, 0xa0, 0x0a, 0xad, 0xb7, 0xb1, 0x0e, 0xe9, 0x54, 0xf0, 0xab, 0xc5, 0x46,
	0x3b, 0xc2, 0x69, 0x5e, 0x5e, 0x09, 0x4e, 0x63, 0x4b, 0x78, 0xaf, 0xb9, 0x44, 0x2f, 0x6a, 0x5d,
	0xa7, 0x39, 0x6f, 0xda, 0xec, 0x37, 0xcd, 0xd9, 0xb5, 0x79, 0x86, 0x3c, 0x73, 0xa9, 0x6f, 0x98,
	0x3d, 0x00, 0xb3, 0xfd, 0x5a, 0xbf, 0xdc, 0x32, 0x58, 0x6a, 0xe9, 0x59, 0xcb, 0x60, 0x89, 0xe5,
	0xef, 0x0e, 0xda, 0xaf, 0x5f, 0x72, 0xdc, 0x79, 0x7a, 0xff, 0x07, 0xe5, 0xf8, 0xd6, 0xb9, 0xd5,
	0x7a, 0x2c, 0x9b, 0x36, 0x66, 0x9c, 0x9c, 0xb7, 0x9f, 0x7f, 0xc3, 0x3c, 0x58, 0xc2, 0xfc, 0xff,
	0x77, 0xc4, 0x3c, 0x58, 0xca, 0x3c, 0xe8, 0x30, 0x9f, 0xfc, 0xe1, 0xa0, 0x2d, 0x65, 0xa1, 0x87,
	0x3c, 0x8c, 0xae, 0xef, 0x1c, 0xb4, 0xa6, 0x8f, 0x30, 0xbe, 0x7a, 0x87, 0x97, 0xb7, 0xca, 0xdd,
	0x8c, 0xd0, 0xb0, 0x26, 0x71, 0x73, 0x2f, 0xb9, 0xb7, 0x64, 0x2f, 0xf9, 0x5c, 0xcf, 0x5f, 0x99,
	0x15, 0x54, 0xb5, 0x94, 0x5a, 0x73, 0x04, 0x8d, 0x4b, 0x96, 0x08, 0x08, 0x62, 0x10, 0x9a, 0xa3,
	0x3b, 0x46, 0x3d, 0x46, 0x58, 0x69, 0xb4, 0xda, 0x59, 0x5b, 0x34, 0xf9, 0x73, 0x45, 0x2f, 0x39,
	0x86, 0x96, 0x5a, 0xbd, 0xdc, 0x9f, 0x1c, 0x34, 0x62, 0x34, 0x9b, 0xce, 0xa2, 0xb2, 0x6a, 0x6d,
	0x04, 0x77, 0x98, 0x9f, 0xa1, 0xa1, 0x63, 0xba, 0xf0, 0x1d, 0xb4, 0x65, 0x29, 0xa7, 0x39, 0x99,
	0x9a, 0xe0, 0xfa, 0x51, 0xc1, 0xcf, 0x89, 0x9c, 0xbd, 0x50, 0x32, 0xd7, 0x47, 0xbb, 0x16, 0x45,
	0x92, 0x39, 0x8e, 0x32, 0x89, 0x0b, 0xc2, 0x61, 0xcd, 0x18, 0x84, 0xa3, 0x1a, 0x7b, 0x9c, 0xcc,
	0x3f, 0xcd, 0xe4, 0x29, 0xe1, 0x0b, 0x78, 0xc1, 0x12, 0x8b, 0x5f, 0x5d, 0xc0, 0x5f, 0xb0, 0xa4,
	0xc6, 0x3f, 0x44, 0x43, 0x58, 0x59, 0xa3, 0xca, 0x42, 0xef, 0x5b, 0x1a, 0x67, 0x51, 0x55, 0xc3,
	0xde, 0x40, 0x28, 0x2f, 0x63, 0x92, 0x83, 0x63, 0xd8, 0x2c, 0x36, 0xc2, 0x4d, 0x90, 0x28, 0x77,
	0xee, 0x43, 0xb4, 0x55, 0xd1, 0x98, 0x66, 0x73, 0x9a, 0xe8, 0x59, 0xed, 0xad, 0xc3, 0xf6, 0x35,
	0x30, 0x52, 0x98, 0xd3, 0xaa, 0x37, 0x22, 0x60, 0xc6, 0xcb, 0x8c, 0x49, 0x5a, 0xc1, 0x2e, 0xb1,
	0x1a, 0xf6, 0x41, 0x78, 0xae, 0x65, 0x93, 0xbf, 0xd6, 0xd0, 0xb0, 0xb3, 0x49, 0xb7, 0xf7, 0xe3,
	0xc3, 0x85, 0xfd, 0xf8, 0x47, 0x07, 0xa1, 0xe6, 0x42, 0x78, 0x4f, 0xa0, 0xc2, 0x57, 0xb7, 0x5a,
	0xe1, 0xe6, 0x2e, 0x86, 0x9b, 0xd1, 0x94, 0x9f, 0xeb, 0x6b, 0x30, 0x41, 0x03, 0x13, 0x80, 0x2e,
	0xee, 0x91, 0xee, 0x5c, 0x95, 0x55, 0x2a, 0x75, 0x6d, 0x1f, 0xd5, 0xb9, 0xa7, 0x12, 0xcf, 0x69,
	0x25, 0xb2, 0x92, 0x79, 0x1f, 0x00, 0x6a, 0xa0, 0x51, 0xaf, 0xb4, 0xd0, 0x7d, 0x13, 0xf5, 0xd8,
	0x65, 0x81, 0xcb, 0x54, 0x67, 0x3f, 0x00, 0xcc, 0x26, 0xbb, 0x2c, 0xbe, 0x4c, 0x21, 0xfb, 0x3f,
	0x3b, 0x68, 0xbb, 0x76, 0xa0, 0xaf, 0x94, 0x24, 0x05, 0xf7, 0x9e, 0xde, 0x41, 0xb7, 0x9b, 0x0b,
	0x1d, 0x8e, 0x6a, 0x3e, 0x2f, 0x0d, 0x1d, 0x55, 0xa9, 0x9e, 0x21, 0x49, 0xa6, 0xd4, 0x7b, 0x76,
	0x97, 0xf4, 0x50, 0xcd, 0xe4, 0x78, 0x4a, 0x55, 0x15, 0x66, 0x44, 0x60, 0xdd, 0xde, 0xba, 0x79,
	0x3f, 0x84, 0xfe, 0x1e, 0xcc, 0x88, 0x38, 0x51, 0x52, 0xdd, 0xbc, 0x8f, 0xd0, 0x50, 0x55, 0xaa,
	0x8d, 0xfb, 0x48, 0x57, 0x8b, 0x51, 0xd9, 0xc2, 0xfd, 0xe0, 0xa0, 0x0d, 0xf3, 0xab, 0xde, 0xc7,
	0xe3, 0x95, 0x5b, 0x5f, 0x0f, 0xdb, 0x0f, 0x61, 0xb8, 0x5e, 0xdf, 0xf8, 0x68, 0x0d, 0xfe, 0x10,
	0x3f, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x78, 0x58, 0xfb, 0x29, 0x0f, 0x00, 0x00,
}