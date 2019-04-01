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
// source: bgp_rcv_pfxlst_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_afs_af_prefix_filters_prefix_filter

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

type BgpRcvPfxlstBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRcvPfxlstBag_KEYS) Reset()         { *m = BgpRcvPfxlstBag_KEYS{} }
func (m *BgpRcvPfxlstBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRcvPfxlstBag_KEYS) ProtoMessage()    {}
func (*BgpRcvPfxlstBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c331b238016cd4, []int{0}
}

func (m *BgpRcvPfxlstBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRcvPfxlstBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRcvPfxlstBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRcvPfxlstBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRcvPfxlstBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRcvPfxlstBag_KEYS.Merge(m, src)
}
func (m *BgpRcvPfxlstBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRcvPfxlstBag_KEYS.Size(m)
}
func (m *BgpRcvPfxlstBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRcvPfxlstBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRcvPfxlstBag_KEYS proto.InternalMessageInfo

func (m *BgpRcvPfxlstBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpRcvPfxlstBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpRcvPfxlstBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpRcvPfxlstBag_KEYS) GetNeighborAddress() string {
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
	return fileDescriptor_27c331b238016cd4, []int{1}
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
	return fileDescriptor_27c331b238016cd4, []int{2}
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
	return fileDescriptor_27c331b238016cd4, []int{3}
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
	return fileDescriptor_27c331b238016cd4, []int{4}
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
	return fileDescriptor_27c331b238016cd4, []int{5}
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
	return fileDescriptor_27c331b238016cd4, []int{6}
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

type BgpPfxlst_Item struct {
	BgpPrefix            *BgpPrefixtype `protobuf:"bytes,1,opt,name=bgp_prefix,json=bgpPrefix,proto3" json:"bgp_prefix,omitempty"`
	SequenceNumber       uint32         `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	IsPrefixGrant        bool           `protobuf:"varint,3,opt,name=is_prefix_grant,json=isPrefixGrant,proto3" json:"is_prefix_grant,omitempty"`
	MinPrefixLength      uint32         `protobuf:"varint,4,opt,name=min_prefix_length,json=minPrefixLength,proto3" json:"min_prefix_length,omitempty"`
	MaxPrefixLength      uint32         `protobuf:"varint,5,opt,name=max_prefix_length,json=maxPrefixLength,proto3" json:"max_prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpPfxlst_Item) Reset()         { *m = BgpPfxlst_Item{} }
func (m *BgpPfxlst_Item) String() string { return proto.CompactTextString(m) }
func (*BgpPfxlst_Item) ProtoMessage()    {}
func (*BgpPfxlst_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c331b238016cd4, []int{7}
}

func (m *BgpPfxlst_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPfxlst_Item.Unmarshal(m, b)
}
func (m *BgpPfxlst_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPfxlst_Item.Marshal(b, m, deterministic)
}
func (m *BgpPfxlst_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPfxlst_Item.Merge(m, src)
}
func (m *BgpPfxlst_Item) XXX_Size() int {
	return xxx_messageInfo_BgpPfxlst_Item.Size(m)
}
func (m *BgpPfxlst_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPfxlst_Item.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPfxlst_Item proto.InternalMessageInfo

func (m *BgpPfxlst_Item) GetBgpPrefix() *BgpPrefixtype {
	if m != nil {
		return m.BgpPrefix
	}
	return nil
}

func (m *BgpPfxlst_Item) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *BgpPfxlst_Item) GetIsPrefixGrant() bool {
	if m != nil {
		return m.IsPrefixGrant
	}
	return false
}

func (m *BgpPfxlst_Item) GetMinPrefixLength() uint32 {
	if m != nil {
		return m.MinPrefixLength
	}
	return 0
}

func (m *BgpPfxlst_Item) GetMaxPrefixLength() uint32 {
	if m != nil {
		return m.MaxPrefixLength
	}
	return 0
}

type BgpPfxlst_Entry struct {
	BgpPfxlst_           []*BgpPfxlst_Item `protobuf:"bytes,6,rep,name=bgp_pfxlst_,json=bgpPfxlst,proto3" json:"bgp_pfxlst_,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BgpPfxlst_Entry) Reset()         { *m = BgpPfxlst_Entry{} }
func (m *BgpPfxlst_Entry) String() string { return proto.CompactTextString(m) }
func (*BgpPfxlst_Entry) ProtoMessage()    {}
func (*BgpPfxlst_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c331b238016cd4, []int{8}
}

func (m *BgpPfxlst_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPfxlst_Entry.Unmarshal(m, b)
}
func (m *BgpPfxlst_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPfxlst_Entry.Marshal(b, m, deterministic)
}
func (m *BgpPfxlst_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPfxlst_Entry.Merge(m, src)
}
func (m *BgpPfxlst_Entry) XXX_Size() int {
	return xxx_messageInfo_BgpPfxlst_Entry.Size(m)
}
func (m *BgpPfxlst_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPfxlst_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPfxlst_Entry proto.InternalMessageInfo

func (m *BgpPfxlst_Entry) GetBgpPfxlst_() []*BgpPfxlst_Item {
	if m != nil {
		return m.BgpPfxlst_
	}
	return nil
}

type BgpRcvPfxlstBag struct {
	AfName               string           `protobuf:"bytes,50,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	NeighborAddressXr    *BgpAddrtype     `protobuf:"bytes,51,opt,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	PrefixListInfo       *BgpPfxlst_Entry `protobuf:"bytes,52,opt,name=prefix_list_info,json=prefixListInfo,proto3" json:"prefix_list_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BgpRcvPfxlstBag) Reset()         { *m = BgpRcvPfxlstBag{} }
func (m *BgpRcvPfxlstBag) String() string { return proto.CompactTextString(m) }
func (*BgpRcvPfxlstBag) ProtoMessage()    {}
func (*BgpRcvPfxlstBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_27c331b238016cd4, []int{9}
}

func (m *BgpRcvPfxlstBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRcvPfxlstBag.Unmarshal(m, b)
}
func (m *BgpRcvPfxlstBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRcvPfxlstBag.Marshal(b, m, deterministic)
}
func (m *BgpRcvPfxlstBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRcvPfxlstBag.Merge(m, src)
}
func (m *BgpRcvPfxlstBag) XXX_Size() int {
	return xxx_messageInfo_BgpRcvPfxlstBag.Size(m)
}
func (m *BgpRcvPfxlstBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRcvPfxlstBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRcvPfxlstBag proto.InternalMessageInfo

func (m *BgpRcvPfxlstBag) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpRcvPfxlstBag) GetNeighborAddressXr() *BgpAddrtype {
	if m != nil {
		return m.NeighborAddressXr
	}
	return nil
}

func (m *BgpRcvPfxlstBag) GetPrefixListInfo() *BgpPfxlst_Entry {
	if m != nil {
		return m.PrefixListInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*BgpRcvPfxlstBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_rcv_pfxlst_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_addrtype")
	proto.RegisterType((*BgpPrefixtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_prefixtype")
	proto.RegisterType((*BgpPfxlst_Item)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_pfxlst__item")
	proto.RegisterType((*BgpPfxlst_Entry)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_pfxlst__entry")
	proto.RegisterType((*BgpRcvPfxlstBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.afs.af.prefix_filters.prefix_filter.bgp_rcv_pfxlst_bag")
}

func init() { proto.RegisterFile("bgp_rcv_pfxlst_bag.proto", fileDescriptor_27c331b238016cd4) }

var fileDescriptor_27c331b238016cd4 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcf, 0x6b, 0x24, 0x45,
	0x14, 0xa6, 0x77, 0x76, 0x27, 0x49, 0xcd, 0xef, 0xce, 0x66, 0x33, 0xeb, 0x29, 0x8e, 0xe0, 0x66,
	0x65, 0x69, 0x70, 0x76, 0x68, 0x3d, 0x78, 0x11, 0x59, 0x45, 0x4c, 0x42, 0x98, 0xf5, 0xa0, 0xa7,
	0xa2, 0xa6, 0x53, 0x3d, 0x5b, 0xd0, 0xbf, 0xac, 0xaa, 0x4c, 0x26, 0x20, 0xa2, 0x07, 0x41, 0xf0,
	0xea, 0x35, 0x07, 0xef, 0xe2, 0x49, 0xf0, 0xe4, 0x7f, 0xe0, 0xc1, 0xbb, 0xff, 0x8c, 0xd4, 0xab,
	0xea, 0xea, 0xea, 0x9e, 0x51, 0xf0, 0xe2, 0x78, 0x09, 0xd3, 0xef, 0xfb, 0xde, 0xe3, 0x7d, 0xdf,
	0xab, 0xee, 0x7a, 0x41, 0xe3, 0xc5, 0xb2, 0xc0, 0x3c, 0x5a, 0xe1, 0x22, 0x5e, 0x27, 0x42, 0xe2,
	0x05, 0x59, 0x06, 0x05, 0xcf, 0x65, 0xee, 0x67, 0x11, 0x13, 0x51, 0x8e, 0x59, 0x2e, 0xf0, 0x9a,
	0x63, 0x56, 0xac, 0x66, 0x58, 0x71, 0xf3, 0x82, 0xf2, 0x60, 0xb1, 0x2c, 0x02, 0x96, 0x09, 0x49,
	0xb2, 0x88, 0x0a, 0xfb, 0xcb, 0xfe, 0xc0, 0x24, 0x92, 0x6c, 0x45, 0x83, 0x15, 0x8f, 0x85, 0xfa,
	0x13, 0x90, 0x58, 0x04, 0x24, 0x0e, 0x0a, 0x4e, 0x63, 0xb6, 0xc6, 0x31, 0x4b, 0x24, 0xe5, 0xa2,
	0xfe, 0x38, 0xb9, 0xf3, 0xd0, 0xf1, 0x66, 0x33, 0xf8, 0x93, 0x17, 0x9f, 0xbf, 0xf4, 0xdf, 0x40,
	0x3d, 0x5b, 0x3b, 0x23, 0x29, 0x1d, 0x7b, 0x27, 0xde, 0xe9, 0xc1, 0xbc, 0x5b, 0x06, 0x2f, 0x48,
	0x4a, 0xfd, 0xc7, 0x68, 0x7f, 0xc5, 0x63, 0x8d, 0xdf, 0x03, 0x7c, 0x6f, 0xc5, 0x63, 0x80, 0x8e,
	0xd1, 0x1e, 0x31, 0x48, 0x0b, 0x90, 0x36, 0xd1, 0xc0, 0x53, 0x34, 0xcc, 0x28, 0x5b, 0xbe, 0x5a,
	0xe4, 0x1c, 0x93, 0xab, 0x2b, 0x4e, 0x85, 0x18, 0xdf, 0x07, 0xc6, 0xa0, 0x8c, 0xbf, 0xaf, 0xc3,
	0x93, 0x77, 0xd0, 0x50, 0xb5, 0x97, 0x4c, 0x57, 0x45, 0x06, 0x5c, 0x2c, 0x55, 0x5f, 0xd5, 0xb3,
	0xca, 0xf5, 0x4e, 0x5a, 0xa7, 0xbd, 0x79, 0x17, 0x82, 0x65, 0xe2, 0x7b, 0xe8, 0xa8, 0x4a, 0x4c,
	0x45, 0x71, 0xf3, 0xaf, 0xb2, 0x2f, 0xf4, 0x88, 0xc0, 0x7f, 0xc1, 0x8b, 0x3c, 0x61, 0xd1, 0x6d,
	0x59, 0x60, 0x8a, 0x8e, 0x36, 0xe3, 0x55, 0xa1, 0x43, 0x05, 0xbe, 0x34, 0xd8, 0x66, 0xbd, 0xf0,
	0x6f, 0xea, 0x85, 0xff, 0x54, 0x2f, 0x6c, 0xd6, 0xfb, 0xa3, 0x87, 0xba, 0xaa, 0xa0, 0xa2, 0xca,
	0xdb, 0x82, 0xfa, 0x43, 0xd4, 0x22, 0x31, 0x33, 0x13, 0x52, 0x3f, 0xfd, 0xd7, 0x51, 0x17, 0xda,
	0x2c, 0xab, 0xe9, 0xe1, 0x74, 0x54, 0xcc, 0x54, 0xf1, 0x9f, 0x21, 0x1f, 0x28, 0x69, 0x44, 0x84,
	0xb4, 0x44, 0x3d, 0xab, 0xa1, 0x42, 0xce, 0x15, 0xd0, 0x64, 0x27, 0x64, 0x41, 0x93, 0xc6, 0xdc,
	0x80, 0x7d, 0xa6, 0x80, 0x92, 0x1d, 0x20, 0x30, 0x02, 0xcb, 0xeb, 0x2c, 0x73, 0xe8, 0x0f, 0x80,
	0x3e, 0x52, 0xd0, 0xa7, 0x80, 0x94, 0xfc, 0x27, 0x68, 0xa0, 0x82, 0xe9, 0x55, 0xd5, 0x48, 0x1b,
	0xb8, 0x7d, 0x13, 0x6e, 0x10, 0xdd, 0x09, 0xee, 0x55, 0xc4, 0x6a, 0x86, 0xfe, 0xdb, 0xe8, 0xa1,
	0x89, 0x10, 0xad, 0xd0, 0xb0, 0xf7, 0x81, 0x7d, 0x58, 0x62, 0xe7, 0x15, 0x64, 0x3c, 0x0b, 0x6d,
	0xe1, 0x03, 0xeb, 0x59, 0x58, 0x77, 0x21, 0x6c, 0x78, 0x86, 0xac, 0x0b, 0xe1, 0x16, 0xcf, 0xc2,
	0x86, 0x67, 0x9d, 0x8a, 0x5d, 0xf3, 0x4c, 0x4b, 0x0b, 0x5d, 0x69, 0x5d, 0x2b, 0x2d, 0x74, 0xa4,
	0x99, 0x23, 0x03, 0x47, 0xbb, 0xd6, 0x47, 0xcf, 0x6a, 0x53, 0x60, 0xad, 0x95, 0x9f, 0x3d, 0x34,
	0x82, 0x33, 0xbe, 0x2a, 0x12, 0x61, 0x13, 0xfa, 0x27, 0xde, 0x69, 0x67, 0xfa, 0xb5, 0x17, 0xfc,
	0xb7, 0xdf, 0x9d, 0xa0, 0xf9, 0x52, 0xcf, 0x87, 0xb6, 0x37, 0x47, 0x24, 0x97, 0x38, 0xca, 0x33,
	0x21, 0x39, 0x61, 0x59, 0x25, 0x72, 0xa0, 0x45, 0x72, 0xf9, 0x81, 0xc5, 0xca, 0x9c, 0xa7, 0x08,
	0x5c, 0x4d, 0x5d, 0x0b, 0x87, 0xfa, 0xcb, 0x52, 0xc6, 0xeb, 0xd4, 0x59, 0x8d, 0x3a, 0xb2, 0xd4,
	0x99, 0x4b, 0x7d, 0x86, 0x7c, 0xdd, 0x2b, 0x75, 0xc9, 0xbe, 0x9e, 0x22, 0x20, 0x2f, 0x1c, 0xf6,
	0x04, 0xf5, 0x12, 0x81, 0x1d, 0x8f, 0x0f, 0xf5, 0x29, 0x4a, 0xc4, 0x99, 0xd5, 0xf6, 0x8b, 0x57,
	0x96, 0xb4, 0x9f, 0x26, 0xc5, 0x7c, 0x08, 0xd3, 0xf8, 0x76, 0x87, 0xd3, 0x70, 0xbe, 0x94, 0x46,
	0xda, 0xb9, 0x28, 0x6e, 0xea, 0xe7, 0x6e, 0x86, 0xe3, 0x24, 0xbf, 0x11, 0x05, 0x8d, 0x6c, 0xe3,
	0x47, 0xd5, 0x3b, 0xf5, 0xa1, 0xc1, 0x1a, 0x67, 0x75, 0x33, 0xe7, 0x51, 0x75, 0x56, 0x9b, 0x39,
	0xef, 0xa2, 0x71, 0xf9, 0x8e, 0x6f, 0xa4, 0x1d, 0x43, 0xda, 0x23, 0x83, 0x6f, 0xcf, 0x0c, 0xb7,
	0x66, 0x8e, 0x6d, 0x66, 0xb8, 0x25, 0xf3, 0x37, 0x4f, 0x8b, 0x13, 0x1c, 0x37, 0xbe, 0xc3, 0x8f,
	0x61, 0x2a, 0xdf, 0xed, 0x64, 0x2a, 0xdb, 0x6e, 0x20, 0xed, 0xb3, 0xe0, 0x97, 0xee, 0x95, 0x50,
	0xf6, 0x1f, 0x6e, 0xf6, 0xff, 0xda, 0x6e, 0xfb, 0x0f, 0xb7, 0xf6, 0x1f, 0x36, 0xfa, 0x9f, 0xfc,
	0xee, 0xa1, 0xbe, 0xca, 0xd0, 0x35, 0xe0, 0x52, 0xfb, 0xc1, 0x43, 0x6d, 0xfd, 0x08, 0x17, 0x5b,
	0x67, 0xfa, 0xe5, 0x2e, 0x24, 0x94, 0x77, 0xec, 0xdc, 0xf4, 0xa2, 0x36, 0x08, 0xc3, 0x4a, 0x68,
	0xb6, 0x94, 0xaf, 0xe0, 0x6a, 0xed, 0xcd, 0xbb, 0x3a, 0x78, 0x06, 0xb1, 0xc9, 0x9f, 0xf7, 0xf4,
	0xe6, 0x62, 0x96, 0x2a, 0xcc, 0x24, 0x4d, 0xfd, 0x3b, 0x0f, 0xa1, 0x4a, 0xa3, 0x11, 0xf5, 0xd5,
	0x2e, 0x44, 0x55, 0x2e, 0xcf, 0x0f, 0x16, 0xcb, 0xe2, 0x52, 0x2b, 0x7b, 0x82, 0x06, 0x82, 0x7e,
	0x71, 0x4d, 0x61, 0xe3, 0xbb, 0x4e, 0x17, 0x94, 0x1b, 0x6d, 0xfd, 0x32, 0x7c, 0x01, 0x51, 0xff,
	0x4d, 0x34, 0x60, 0xc2, 0x14, 0xc1, 0x4b, 0x4e, 0x32, 0x09, 0x6b, 0xc3, 0xfe, 0xbc, 0xc7, 0x84,
	0xae, 0xf5, 0x91, 0x0a, 0xfa, 0x6f, 0xa1, 0x51, 0xca, 0x32, 0x5c, 0xb7, 0xeb, 0x3e, 0x94, 0x1c,
	0xa4, 0x2c, 0xbb, 0x74, 0x1c, 0x03, 0x2e, 0x59, 0x37, 0xb8, 0x0f, 0x0c, 0x97, 0xac, 0x5d, 0xee,
	0xe4, 0x57, 0x0f, 0x8d, 0x5c, 0x77, 0x69, 0x26, 0xf9, 0xad, 0xff, 0xa3, 0x87, 0x3a, 0x4e, 0x74,
	0xdc, 0x3e, 0x69, 0xed, 0xec, 0x72, 0x73, 0xe7, 0xae, 0x2d, 0x86, 0xc0, 0xe4, 0xfb, 0x16, 0xf2,
	0x37, 0x17, 0x6e, 0x77, 0x57, 0x9e, 0xd6, 0x76, 0xe5, 0x9f, 0x3c, 0x74, 0xd8, 0x5c, 0x96, 0xf1,
	0x9a, 0x8f, 0x9f, 0xff, 0x0f, 0x5e, 0x88, 0x51, 0x63, 0x5b, 0xff, 0x8c, 0xab, 0x76, 0x87, 0xe5,
	0x04, 0x99, 0x90, 0x98, 0x65, 0x71, 0x3e, 0x9e, 0x41, 0xaf, 0xdf, 0xec, 0x74, 0x0e, 0x70, 0x42,
	0xe6, 0x7d, 0xf3, 0x8e, 0x32, 0x21, 0x3f, 0xce, 0xe2, 0x7c, 0xd1, 0x86, 0xff, 0xba, 0x9e, 0xff,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xce, 0xe1, 0x6f, 0x91, 0x0d, 0x00, 0x00,
}
