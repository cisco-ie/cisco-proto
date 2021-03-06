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
// source: bgp_rpc_set_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_afs_af_rpc_sets_rpc_set

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

type BgpRpcSetBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	RpcSetIndex          uint32   `protobuf:"varint,3,opt,name=rpc_set_index,json=rpcSetIndex,proto3" json:"rpc_set_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpRpcSetBag_KEYS) Reset()         { *m = BgpRpcSetBag_KEYS{} }
func (m *BgpRpcSetBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpRpcSetBag_KEYS) ProtoMessage()    {}
func (*BgpRpcSetBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_78908996c17992a5, []int{0}
}

func (m *BgpRpcSetBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpcSetBag_KEYS.Unmarshal(m, b)
}
func (m *BgpRpcSetBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpcSetBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpRpcSetBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpcSetBag_KEYS.Merge(m, src)
}
func (m *BgpRpcSetBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpRpcSetBag_KEYS.Size(m)
}
func (m *BgpRpcSetBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpcSetBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpcSetBag_KEYS proto.InternalMessageInfo

func (m *BgpRpcSetBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpRpcSetBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpRpcSetBag_KEYS) GetRpcSetIndex() uint32 {
	if m != nil {
		return m.RpcSetIndex
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
	return fileDescriptor_78908996c17992a5, []int{1}
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
	return fileDescriptor_78908996c17992a5, []int{2}
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
	return fileDescriptor_78908996c17992a5, []int{3}
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
	return fileDescriptor_78908996c17992a5, []int{4}
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
	return fileDescriptor_78908996c17992a5, []int{5}
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

type BgpRpcSetBag struct {
	NextHopXr            uint32         `protobuf:"varint,50,opt,name=next_hop_xr,json=nextHopXr,proto3" json:"next_hop_xr,omitempty"`
	NextHop              []*BgpAddrtype `protobuf:"bytes,51,rep,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Label                uint32         `protobuf:"varint,52,opt,name=label,proto3" json:"label,omitempty"`
	RefCount             uint32         `protobuf:"varint,53,opt,name=ref_count,json=refCount,proto3" json:"ref_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpRpcSetBag) Reset()         { *m = BgpRpcSetBag{} }
func (m *BgpRpcSetBag) String() string { return proto.CompactTextString(m) }
func (*BgpRpcSetBag) ProtoMessage()    {}
func (*BgpRpcSetBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_78908996c17992a5, []int{6}
}

func (m *BgpRpcSetBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpRpcSetBag.Unmarshal(m, b)
}
func (m *BgpRpcSetBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpRpcSetBag.Marshal(b, m, deterministic)
}
func (m *BgpRpcSetBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpRpcSetBag.Merge(m, src)
}
func (m *BgpRpcSetBag) XXX_Size() int {
	return xxx_messageInfo_BgpRpcSetBag.Size(m)
}
func (m *BgpRpcSetBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpRpcSetBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpRpcSetBag proto.InternalMessageInfo

func (m *BgpRpcSetBag) GetNextHopXr() uint32 {
	if m != nil {
		return m.NextHopXr
	}
	return 0
}

func (m *BgpRpcSetBag) GetNextHop() []*BgpAddrtype {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *BgpRpcSetBag) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *BgpRpcSetBag) GetRefCount() uint32 {
	if m != nil {
		return m.RefCount
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpRpcSetBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_rpc_set_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_addrtype")
	proto.RegisterType((*BgpRpcSetBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.afs.af.rpc_sets.rpc_set.bgp_rpc_set_bag")
}

func init() { proto.RegisterFile("bgp_rpc_set_bag.proto", fileDescriptor_78908996c17992a5) }

var fileDescriptor_78908996c17992a5 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcf, 0x6f, 0xd3, 0x3e,
	0x14, 0x57, 0xbe, 0xfd, 0xee, 0x47, 0xdd, 0x76, 0xed, 0xbc, 0x6e, 0x0b, 0x20, 0xa1, 0x52, 0x0e,
	0x14, 0x69, 0x8a, 0x44, 0x57, 0x0c, 0x07, 0x2e, 0x68, 0x1a, 0x02, 0xb1, 0x4d, 0xa8, 0x43, 0x08,
	0x4e, 0x96, 0x9b, 0x3a, 0x23, 0x52, 0x9a, 0x58, 0xb6, 0xd7, 0x76, 0x12, 0x12, 0xff, 0x01, 0xff,
	0x05, 0x17, 0xc4, 0x91, 0x7f, 0x8d, 0x3b, 0xb2, 0x13, 0x3b, 0x69, 0x5a, 0x90, 0xb8, 0x94, 0x9b,
	0xfd, 0x3e, 0x9f, 0xf7, 0xfc, 0xde, 0xe7, 0xbd, 0xd8, 0x01, 0xfb, 0xa3, 0x2b, 0x86, 0x39, 0xf3,
	0xb1, 0xa0, 0x12, 0x8f, 0xc8, 0x95, 0xc7, 0x78, 0x22, 0x13, 0x48, 0xfd, 0x50, 0xf8, 0x09, 0x0e,
	0x13, 0x81, 0xe7, 0x1c, 0x87, 0x6c, 0x3a, 0xc0, 0x8a, 0x98, 0x30, 0xca, 0xbd, 0xd1, 0x15, 0xf3,
	0xc2, 0x58, 0x48, 0x12, 0xfb, 0x54, 0xd8, 0x95, 0x5d, 0x60, 0xe2, 0xcb, 0x70, 0x4a, 0xbd, 0x31,
	0x0d, 0xc8, 0x75, 0x24, 0xf1, 0x94, 0x07, 0x1e, 0x09, 0x84, 0x47, 0x02, 0x2f, 0x3b, 0x47, 0x98,
	0x45, 0x77, 0x0e, 0xda, 0xa5, 0xf3, 0xf1, 0xeb, 0xd3, 0x0f, 0x97, 0xf0, 0x3e, 0x68, 0xd8, 0x70,
	0x31, 0x99, 0x50, 0xd7, 0xe9, 0x38, 0xbd, 0xea, 0xb0, 0x6e, 0x8c, 0x17, 0x64, 0x42, 0xe1, 0x21,
	0xd8, 0x22, 0x41, 0x0a, 0xff, 0xa7, 0xe1, 0x4d, 0x12, 0x68, 0xa0, 0x0b, 0x1a, 0x26, 0x62, 0x18,
	0x8f, 0xe9, 0xdc, 0xad, 0x74, 0x9c, 0x5e, 0x63, 0x58, 0xe3, 0xcc, 0xbf, 0xa4, 0xf2, 0x95, 0x32,
	0x75, 0x9f, 0x80, 0x96, 0x3a, 0x39, 0xea, 0x4f, 0x59, 0x8c, 0xc9, 0x78, 0xcc, 0xb1, 0x54, 0xa7,
	0xe6, 0x7b, 0x2a, 0x84, 0xeb, 0x74, 0x2a, 0xbd, 0xc6, 0xb0, 0xae, 0x8d, 0xcf, 0x53, 0x5b, 0xf7,
	0x59, 0x2a, 0x59, 0x4a, 0x9c, 0x08, 0x36, 0xfb, 0x2b, 0xef, 0x0b, 0xe0, 0x2a, 0x6f, 0x2d, 0xa8,
	0xe0, 0x2c, 0x89, 0x42, 0xff, 0xc6, 0x04, 0xe8, 0x83, 0xfd, 0x65, 0x7b, 0x1e, 0x68, 0x4f, 0x81,
	0x97, 0x19, 0xb6, 0x1c, 0x0f, 0xfd, 0x26, 0x1e, 0xfa, 0x53, 0x3c, 0x54, 0x8e, 0xf7, 0xbd, 0x01,
	0xea, 0x2a, 0xa0, 0xa2, 0xca, 0x1b, 0x46, 0x61, 0x0b, 0x54, 0x48, 0x10, 0x66, 0xfa, 0xab, 0x25,
	0xbc, 0x07, 0xea, 0x3a, 0x4d, 0x13, 0x2d, 0xd5, 0xbe, 0xa6, 0x6c, 0x59, 0x14, 0x78, 0x04, 0xa0,
	0xa6, 0x4c, 0x7c, 0x22, 0xa4, 0x25, 0x56, 0x34, 0xb1, 0xa5, 0x90, 0x73, 0x05, 0x94, 0xd9, 0x11,
	0x19, 0xd1, 0xc8, 0xb2, 0xff, 0xcf, 0xd9, 0x67, 0x0a, 0x30, 0x6c, 0x0f, 0x68, 0x21, 0xb0, 0xbc,
	0x8e, 0xe3, 0x02, 0x7d, 0x43, 0xd3, 0x77, 0x15, 0xf4, 0x56, 0x23, 0x86, 0xdf, 0x03, 0xad, 0x34,
	0x97, 0x71, 0x9e, 0xc9, 0xa6, 0x26, 0xef, 0xe8, 0x4c, 0xc6, 0x36, 0x8f, 0x07, 0xa0, 0xa9, 0x2c,
	0xc5, 0x16, 0x6e, 0xe5, 0xc4, 0xbc, 0x89, 0xf0, 0x11, 0x68, 0x67, 0x16, 0x92, 0x96, 0x98, 0xb1,
	0xb7, 0x35, 0x7b, 0xcf, 0x60, 0xe7, 0x39, 0x94, 0x89, 0x86, 0x6c, 0xe0, 0xaa, 0x15, 0x0d, 0x2d,
	0xca, 0x80, 0x4a, 0xa2, 0x01, 0x2b, 0x03, 0x5a, 0x21, 0x1a, 0x2a, 0x89, 0x56, 0xcb, 0xd9, 0x0b,
	0xa2, 0xa5, 0xa5, 0xa1, 0x62, 0x69, 0x75, 0x5b, 0x1a, 0x2a, 0x94, 0x96, 0xcd, 0x8c, 0x9e, 0xed,
	0x85, 0x3c, 0x1a, 0xb6, 0x36, 0x05, 0x2e, 0xa4, 0xf2, 0xd5, 0x01, 0x30, 0x9d, 0xfc, 0x29, 0x8b,
	0x84, 0xf5, 0xd8, 0xe9, 0x38, 0xbd, 0x5a, 0x7f, 0xe6, 0xad, 0xe5, 0x26, 0xf1, 0xca, 0x1f, 0xf3,
	0xb0, 0xa5, 0x77, 0xef, 0x58, 0x24, 0x0a, 0xb5, 0x71, 0x89, 0xfd, 0x24, 0x16, 0x92, 0x93, 0x30,
	0xce, 0x6b, 0x6b, 0xa6, 0xb5, 0x71, 0x79, 0x62, 0x31, 0xe3, 0xf3, 0x50, 0x4f, 0x0f, 0x9a, 0x14,
	0x95, 0x6b, 0x69, 0x7a, 0xd3, 0xd8, 0x17, 0xa9, 0x83, 0x05, 0xea, 0xae, 0xa5, 0x0e, 0x8a, 0xd4,
	0x23, 0x23, 0x18, 0x2d, 0x92, 0x61, 0xda, 0x3c, 0x8d, 0x9c, 0x16, 0xd8, 0x5d, 0xd0, 0x88, 0x04,
	0x2e, 0x28, 0xbb, 0x97, 0x0e, 0x4f, 0x24, 0xce, 0x6c, 0x6d, 0xdf, 0x6c, 0x0f, 0xec, 0x95, 0xa4,
	0x98, 0x6d, 0xdd, 0x83, 0x4f, 0x6b, 0xef, 0x41, 0xe1, 0x5e, 0xcc, 0x0a, 0x3a, 0x17, 0x6c, 0xb6,
	0x38, 0x64, 0x03, 0x1c, 0x44, 0xc9, 0x4c, 0x30, 0xea, 0xdb, 0x74, 0xf7, 0xf3, 0x0f, 0xe8, 0x45,
	0x86, 0x95, 0x06, 0x73, 0xd9, 0xe7, 0x20, 0x1f, 0xcc, 0xb2, 0xcf, 0x53, 0xe0, 0x9a, 0x0f, 0x7a,
	0xc9, 0xed, 0x50, 0xbb, 0x1d, 0x64, 0xf8, 0x6a, 0x4f, 0xb4, 0xd2, 0xd3, 0xb5, 0x9e, 0x68, 0x85,
	0xe7, 0x0f, 0x07, 0x1c, 0x64, 0xb7, 0x38, 0x2e, 0x5d, 0xbb, 0xb7, 0x74, 0x33, 0x3e, 0xaf, 0xb1,
	0x19, 0xab, 0x9e, 0x19, 0xf3, 0x8e, 0xbc, 0x29, 0xde, 0xfb, 0x26, 0x6d, 0xb4, 0x22, 0xed, 0xdb,
	0xff, 0x22, 0x6d, 0xb4, 0x32, 0x6d, 0x54, 0x4a, 0xbb, 0xfb, 0xd3, 0x01, 0xcd, 0xd2, 0x0f, 0x04,
	0xbc, 0x0b, 0x6a, 0x31, 0x9d, 0x4b, 0xfc, 0x31, 0x61, 0x78, 0xce, 0xdd, 0xbe, 0x7e, 0xfb, 0xab,
	0xca, 0xf4, 0x32, 0x61, 0xef, 0x39, 0xfc, 0xe2, 0x80, 0x6d, 0x43, 0x70, 0x8f, 0x3b, 0x95, 0x5e,
	0xad, 0x2f, 0xd6, 0x58, 0x9c, 0x79, 0x59, 0x87, 0x5b, 0x59, 0x4a, 0xb0, 0x0d, 0x36, 0xf4, 0x2d,
	0xee, 0x0e, 0x74, 0xaa, 0xe9, 0x06, 0xde, 0x01, 0x55, 0x4e, 0x03, 0xec, 0x27, 0xd7, 0xb1, 0x74,
	0x1f, 0x6b, 0x64, 0x9b, 0xd3, 0xe0, 0x44, 0xed, 0x47, 0x9b, 0xfa, 0x2f, 0xed, 0xf8, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0x2b, 0xf2, 0xf2, 0xbe, 0x09, 0x00, 0x00,
}
