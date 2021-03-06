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
// source: bgp_epe_set_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_afs_af_epes_epe

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

type BgpEpeSetBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,3,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	EpeKeyLength         uint32   `protobuf:"varint,4,opt,name=epe_key_length,json=epeKeyLength,proto3" json:"epe_key_length,omitempty"`
	EpeSetKey            string   `protobuf:"bytes,5,opt,name=epe_set_key,json=epeSetKey,proto3" json:"epe_set_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpEpeSetBag_KEYS) Reset()         { *m = BgpEpeSetBag_KEYS{} }
func (m *BgpEpeSetBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpEpeSetBag_KEYS) ProtoMessage()    {}
func (*BgpEpeSetBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f4e09ec364f645, []int{0}
}

func (m *BgpEpeSetBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpEpeSetBag_KEYS.Unmarshal(m, b)
}
func (m *BgpEpeSetBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpEpeSetBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpEpeSetBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpEpeSetBag_KEYS.Merge(m, src)
}
func (m *BgpEpeSetBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpEpeSetBag_KEYS.Size(m)
}
func (m *BgpEpeSetBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpEpeSetBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpEpeSetBag_KEYS proto.InternalMessageInfo

func (m *BgpEpeSetBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *BgpEpeSetBag_KEYS) GetEpeKeyLength() uint32 {
	if m != nil {
		return m.EpeKeyLength
	}
	return 0
}

func (m *BgpEpeSetBag_KEYS) GetEpeSetKey() string {
	if m != nil {
		return m.EpeSetKey
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
	return fileDescriptor_a8f4e09ec364f645, []int{1}
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
	return fileDescriptor_a8f4e09ec364f645, []int{2}
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
	return fileDescriptor_a8f4e09ec364f645, []int{3}
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
	return fileDescriptor_a8f4e09ec364f645, []int{4}
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
	return fileDescriptor_a8f4e09ec364f645, []int{5}
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

type BgpEpeSetBag struct {
	EpeKey               []uint32       `protobuf:"varint,50,rep,packed,name=epe_key,json=epeKey,proto3" json:"epe_key,omitempty"`
	EpeKeyLengthXr       uint32         `protobuf:"varint,51,opt,name=epe_key_length_xr,json=epeKeyLengthXr,proto3" json:"epe_key_length_xr,omitempty"`
	EpeVersion           uint32         `protobuf:"varint,52,opt,name=epe_version,json=epeVersion,proto3" json:"epe_version,omitempty"`
	EpeFlags             uint32         `protobuf:"varint,53,opt,name=epe_flags,json=epeFlags,proto3" json:"epe_flags,omitempty"`
	EpeLocalAsn          uint32         `protobuf:"varint,54,opt,name=epe_local_asn,json=epeLocalAsn,proto3" json:"epe_local_asn,omitempty"`
	EpeRemoteAsn         uint32         `protobuf:"varint,55,opt,name=epe_remote_asn,json=epeRemoteAsn,proto3" json:"epe_remote_asn,omitempty"`
	EpeRemoteRouterId    uint32         `protobuf:"varint,56,opt,name=epe_remote_router_id,json=epeRemoteRouterId,proto3" json:"epe_remote_router_id,omitempty"`
	EpeLocalRouterId     uint32         `protobuf:"varint,57,opt,name=epe_local_router_id,json=epeLocalRouterId,proto3" json:"epe_local_router_id,omitempty"`
	EpeLocalAddress      *BgpAddrtype   `protobuf:"bytes,58,opt,name=epe_local_address,json=epeLocalAddress,proto3" json:"epe_local_address,omitempty"`
	EpeNextHop           *BgpAddrtype   `protobuf:"bytes,59,opt,name=epe_next_hop,json=epeNextHop,proto3" json:"epe_next_hop,omitempty"`
	FirstHop             []*BgpAddrtype `protobuf:"bytes,60,rep,name=first_hop,json=firstHop,proto3" json:"first_hop,omitempty"`
	NexthopId            []uint32       `protobuf:"varint,61,rep,packed,name=nexthop_id,json=nexthopId,proto3" json:"nexthop_id,omitempty"`
	Ifhandle             []uint32       `protobuf:"varint,62,rep,packed,name=ifhandle,proto3" json:"ifhandle,omitempty"`
	Label                uint32         `protobuf:"varint,63,opt,name=label,proto3" json:"label,omitempty"`
	RefCount             uint32         `protobuf:"varint,64,opt,name=ref_count,json=refCount,proto3" json:"ref_count,omitempty"`
	RpcSetObjectId       uint32         `protobuf:"varint,65,opt,name=rpc_set_object_id,json=rpcSetObjectId,proto3" json:"rpc_set_object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpEpeSetBag) Reset()         { *m = BgpEpeSetBag{} }
func (m *BgpEpeSetBag) String() string { return proto.CompactTextString(m) }
func (*BgpEpeSetBag) ProtoMessage()    {}
func (*BgpEpeSetBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8f4e09ec364f645, []int{6}
}

func (m *BgpEpeSetBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpEpeSetBag.Unmarshal(m, b)
}
func (m *BgpEpeSetBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpEpeSetBag.Marshal(b, m, deterministic)
}
func (m *BgpEpeSetBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpEpeSetBag.Merge(m, src)
}
func (m *BgpEpeSetBag) XXX_Size() int {
	return xxx_messageInfo_BgpEpeSetBag.Size(m)
}
func (m *BgpEpeSetBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpEpeSetBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpEpeSetBag proto.InternalMessageInfo

func (m *BgpEpeSetBag) GetEpeKey() []uint32 {
	if m != nil {
		return m.EpeKey
	}
	return nil
}

func (m *BgpEpeSetBag) GetEpeKeyLengthXr() uint32 {
	if m != nil {
		return m.EpeKeyLengthXr
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeVersion() uint32 {
	if m != nil {
		return m.EpeVersion
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeFlags() uint32 {
	if m != nil {
		return m.EpeFlags
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalAsn() uint32 {
	if m != nil {
		return m.EpeLocalAsn
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeRemoteAsn() uint32 {
	if m != nil {
		return m.EpeRemoteAsn
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeRemoteRouterId() uint32 {
	if m != nil {
		return m.EpeRemoteRouterId
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalRouterId() uint32 {
	if m != nil {
		return m.EpeLocalRouterId
	}
	return 0
}

func (m *BgpEpeSetBag) GetEpeLocalAddress() *BgpAddrtype {
	if m != nil {
		return m.EpeLocalAddress
	}
	return nil
}

func (m *BgpEpeSetBag) GetEpeNextHop() *BgpAddrtype {
	if m != nil {
		return m.EpeNextHop
	}
	return nil
}

func (m *BgpEpeSetBag) GetFirstHop() []*BgpAddrtype {
	if m != nil {
		return m.FirstHop
	}
	return nil
}

func (m *BgpEpeSetBag) GetNexthopId() []uint32 {
	if m != nil {
		return m.NexthopId
	}
	return nil
}

func (m *BgpEpeSetBag) GetIfhandle() []uint32 {
	if m != nil {
		return m.Ifhandle
	}
	return nil
}

func (m *BgpEpeSetBag) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *BgpEpeSetBag) GetRefCount() uint32 {
	if m != nil {
		return m.RefCount
	}
	return 0
}

func (m *BgpEpeSetBag) GetRpcSetObjectId() uint32 {
	if m != nil {
		return m.RpcSetObjectId
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpEpeSetBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_epe_set_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_addrtype")
	proto.RegisterType((*BgpEpeSetBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.afs.af.epes.epe.bgp_epe_set_bag")
}

func init() { proto.RegisterFile("bgp_epe_set_bag.proto", fileDescriptor_a8f4e09ec364f645) }

var fileDescriptor_a8f4e09ec364f645 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0xe9, 0x6e, 0x9b, 0x4c, 0x93, 0x26, 0x75, 0xd3, 0xae, 0x77, 0x11, 0x10, 0x02, 0x12,
	0x5d, 0x69, 0x31, 0xa2, 0x1b, 0xbc, 0x0b, 0x2c, 0x3f, 0xaa, 0xd5, 0xae, 0xa8, 0xfa, 0x03, 0x94,
	0xa2, 0x15, 0x88, 0xc3, 0xc8, 0xb1, 0x9f, 0x53, 0x53, 0xc7, 0x1e, 0xcd, 0x4c, 0xd3, 0xe6, 0x0e,
	0x5c, 0x38, 0x73, 0xe2, 0xc2, 0x99, 0x3b, 0xff, 0x1f, 0x9a, 0x37, 0xf6, 0xd8, 0x71, 0x03, 0xd2,
	0x1e, 0xda, 0x4b, 0xea, 0x79, 0xdf, 0xf7, 0x9e, 0xbf, 0x79, 0xef, 0xcd, 0xf8, 0x95, 0x6c, 0x8f,
	0x27, 0x8c, 0x02, 0x03, 0x2a, 0x40, 0xd2, 0xb1, 0x3f, 0x71, 0x19, 0xcf, 0x64, 0x66, 0xff, 0x14,
	0xc4, 0x22, 0xc8, 0x68, 0x9c, 0x09, 0x7a, 0xc5, 0x69, 0xcc, 0x66, 0x43, 0xaa, 0x88, 0x19, 0x03,
	0xee, 0x8e, 0x27, 0xcc, 0x8d, 0x53, 0x21, 0xfd, 0x34, 0x00, 0x61, 0x9e, 0xcc, 0x03, 0x55, 0x7f,
	0xc2, 0xf1, 0xdc, 0x9d, 0xf1, 0x48, 0xa8, 0x1f, 0xd7, 0x8f, 0x84, 0xeb, 0x47, 0x2e, 0x30, 0x10,
	0xea, 0x67, 0xf0, 0x8f, 0x45, 0x7a, 0xb5, 0xd7, 0xd2, 0xc3, 0x17, 0x3f, 0x9e, 0xda, 0xef, 0x91,
	0xb6, 0x89, 0x92, 0xfa, 0x53, 0x70, 0xac, 0xbe, 0xb5, 0xdb, 0x1c, 0xb5, 0x0a, 0xe3, 0x89, 0x3f,
	0x05, 0xfb, 0x3e, 0x69, 0xcc, 0x78, 0xa4, 0xf1, 0x37, 0x10, 0x5f, 0x9b, 0xf1, 0x08, 0xa1, 0x7b,
	0x64, 0xcd, 0xcf, 0x91, 0x15, 0x44, 0x56, 0x7d, 0x0d, 0xbc, 0x4f, 0x36, 0xd4, 0xcb, 0xce, 0x61,
	0x4e, 0x13, 0x48, 0x27, 0xf2, 0xcc, 0xb9, 0xd3, 0xb7, 0x76, 0xdb, 0xa3, 0x16, 0x30, 0x38, 0x84,
	0xf9, 0x11, 0xda, 0xec, 0xb7, 0xc9, 0x7a, 0x21, 0xe9, 0x1c, 0xe6, 0xce, 0x5d, 0x0c, 0xd1, 0x04,
	0x06, 0xa7, 0x20, 0x0f, 0x61, 0x3e, 0x78, 0x42, 0xba, 0x4a, 0x76, 0xb2, 0x37, 0x63, 0x29, 0xf5,
	0xc3, 0x90, 0x53, 0xa9, 0x24, 0x97, 0x6b, 0x10, 0xc2, 0xb1, 0xfa, 0x2b, 0x2a, 0x30, 0x1a, 0xf7,
	0xb5, 0x6d, 0xf0, 0x4c, 0xa7, 0x59, 0x13, 0xa7, 0x82, 0x5d, 0xbe, 0x96, 0xf7, 0x09, 0x71, 0x94,
	0x37, 0x16, 0x41, 0x70, 0x96, 0x25, 0x71, 0x30, 0x2f, 0x02, 0xec, 0x91, 0xed, 0xeb, 0xf6, 0x32,
	0xd0, 0x96, 0x02, 0x4f, 0x73, 0xec, 0x7a, 0x3c, 0xef, 0x3f, 0xe2, 0x79, 0xff, 0x17, 0xcf, 0xab,
	0xc7, 0xfb, 0xa5, 0x4d, 0x5a, 0x2a, 0xa0, 0xa2, 0xca, 0x39, 0x03, 0xbb, 0x4b, 0x56, 0xfc, 0x28,
	0xce, 0x8b, 0xa7, 0x1e, 0xed, 0x77, 0x49, 0x0b, 0x65, 0x16, 0xd1, 0x74, 0xdd, 0xd6, 0x95, 0x2d,
	0x8f, 0x62, 0x3f, 0x22, 0x36, 0x52, 0xa6, 0x81, 0x2f, 0xa4, 0x21, 0xea, 0x32, 0x76, 0x15, 0x72,
	0xac, 0x80, 0x3a, 0x3b, 0xf1, 0xc7, 0x90, 0x18, 0xf6, 0x9d, 0x92, 0x7d, 0xa4, 0x80, 0x82, 0xed,
	0x12, 0x4c, 0x04, 0x95, 0x17, 0x69, 0x5a, 0xa1, 0xeb, 0x02, 0x6f, 0x2a, 0xe8, 0x7b, 0x44, 0x0a,
	0xfe, 0x2e, 0xe9, 0x6a, 0x2d, 0x61, 0xa9, 0x64, 0x15, 0xc9, 0x1b, 0xa8, 0x24, 0x34, 0x3a, 0x3e,
	0x20, 0x1d, 0x65, 0xa9, 0x96, 0x70, 0xad, 0x24, 0x96, 0x45, 0xb4, 0x3f, 0x26, 0xbd, 0xdc, 0xe2,
	0xeb, 0x2d, 0xe6, 0xec, 0x06, 0xb2, 0xb7, 0x0a, 0xec, 0xb8, 0x84, 0xf2, 0xa4, 0x79, 0x26, 0x70,
	0xd3, 0x24, 0xcd, 0x5b, 0x4c, 0x83, 0x57, 0x4b, 0x1a, 0x31, 0x69, 0xf0, 0x96, 0x24, 0xcd, 0xab,
	0x25, 0x6d, 0xbd, 0x64, 0x2f, 0x24, 0x4d, 0x6f, 0xcd, 0xab, 0x6e, 0xad, 0x65, 0xb6, 0xe6, 0x55,
	0xb6, 0x96, 0xf7, 0x0c, 0xf6, 0xf6, 0x82, 0x8e, 0xb6, 0xd9, 0x9b, 0x02, 0x17, 0xa4, 0xfc, 0x69,
	0x11, 0x5b, 0x77, 0xfe, 0x8c, 0x25, 0xc2, 0x78, 0x6c, 0xf4, 0xad, 0xdd, 0xf5, 0xbd, 0xa9, 0x7b,
	0x83, 0xb7, 0x8f, 0x5b, 0x3f, 0xc2, 0xa3, 0x2e, 0xae, 0x5e, 0xb1, 0x44, 0x54, 0x76, 0xc4, 0x25,
	0x0d, 0xb2, 0x54, 0x48, 0xee, 0xc7, 0x69, 0xb9, 0xa3, 0x8e, 0xde, 0x11, 0x97, 0xcf, 0x0d, 0x56,
	0xf8, 0x3c, 0xc4, 0x9e, 0xf1, 0xa6, 0xd5, 0x7c, 0x75, 0x91, 0xde, 0x29, 0xec, 0x8b, 0xd4, 0xe1,
	0x02, 0x75, 0xd3, 0x50, 0x87, 0x55, 0xea, 0xa3, 0x22, 0x4d, 0x50, 0x25, 0xdb, 0xba, 0x64, 0x88,
	0xbc, 0xa8, 0xb0, 0x07, 0xa4, 0x9d, 0x08, 0x5a, 0xc9, 0xe7, 0x96, 0x6e, 0x99, 0x44, 0x1c, 0x99,
	0xbd, 0xfd, 0x65, 0x32, 0x6f, 0x2e, 0x22, 0xc5, 0xec, 0x61, 0xe6, 0xf9, 0x2d, 0x65, 0xbe, 0x72,
	0x07, 0xe6, 0xdb, 0x38, 0x16, 0xec, 0x72, 0xb1, 0xa1, 0x86, 0x34, 0x4a, 0xb2, 0x4b, 0xc1, 0x20,
	0x30, 0x22, 0xb7, 0xcb, 0xc3, 0xf2, 0x32, 0xc7, 0x6a, 0x4d, 0x78, 0xdd, 0x67, 0xa7, 0x6c, 0xc2,
	0xba, 0xcf, 0x53, 0xe2, 0x14, 0x87, 0xf7, 0x9a, 0xdb, 0x3d, 0x74, 0xdb, 0xc9, 0xf1, 0xe5, 0x9e,
	0xde, 0x52, 0x4f, 0xc7, 0x78, 0x7a, 0x4b, 0x3c, 0xff, 0xb6, 0xc8, 0x4e, 0x7e, 0x63, 0xd3, 0xda,
	0x15, 0x7b, 0x1f, 0x4b, 0x70, 0x71, 0xe3, 0x25, 0x58, 0xf6, 0x21, 0x29, 0xbe, 0x14, 0xdf, 0x55,
	0x6f, 0xf6, 0x42, 0xac, 0xb7, 0x44, 0xec, 0x83, 0xdb, 0x13, 0xeb, 0x2d, 0x15, 0xeb, 0xd5, 0xc4,
	0x0e, 0x7e, 0x5d, 0x23, 0x9d, 0xda, 0x54, 0xa1, 0x06, 0x82, 0xfc, 0xbb, 0xef, 0xec, 0xe1, 0x07,
	0x6c, 0x55, 0x7f, 0xf0, 0xed, 0x87, 0x64, 0x73, 0x71, 0x20, 0xa0, 0x57, 0xdc, 0x79, 0x8c, 0x33,
	0xc1, 0x46, 0x75, 0x26, 0xf8, 0x81, 0xdb, 0xef, 0xe8, 0xa9, 0x60, 0x06, 0x5c, 0xc4, 0x59, 0xea,
	0x0c, 0x91, 0x44, 0x80, 0xc1, 0x2b, 0x6d, 0xb1, 0xdf, 0x24, 0x6a, 0x46, 0xa0, 0x51, 0xe2, 0x4f,
	0x84, 0xf3, 0x09, 0xc2, 0x0d, 0x60, 0xf0, 0x52, 0xad, 0xd5, 0x91, 0x54, 0x60, 0x92, 0x05, 0x7e,
	0x42, 0x7d, 0x91, 0x3a, 0x1e, 0x12, 0x54, 0xc8, 0x23, 0x65, 0xdb, 0x17, 0x69, 0x31, 0x9d, 0x70,
	0x98, 0x66, 0x12, 0x90, 0xf4, 0xc4, 0x4c, 0x27, 0x23, 0x34, 0x2a, 0xd6, 0x47, 0xa4, 0x57, 0x61,
	0xf1, 0xec, 0x42, 0x02, 0xa7, 0x71, 0xe8, 0x3c, 0x45, 0xee, 0xa6, 0xe1, 0x8e, 0x10, 0x39, 0x08,
	0xed, 0x0f, 0xc9, 0x56, 0xf9, 0xea, 0x92, 0xff, 0x29, 0xf2, 0xbb, 0x85, 0x00, 0x43, 0xff, 0xc3,
	0xd2, 0x39, 0xc9, 0xa5, 0xe6, 0x75, 0xfe, 0x0c, 0xeb, 0x1c, 0xdf, 0x78, 0x9d, 0x8b, 0xe1, 0x61,
	0xd4, 0x31, 0x99, 0xc9, 0x9b, 0xf0, 0x77, 0x8b, 0xa8, 0x44, 0xd0, 0x14, 0xae, 0x24, 0x3d, 0xcb,
	0x98, 0xf3, 0xf9, 0x6d, 0x4b, 0x52, 0xc5, 0x3e, 0x81, 0x2b, 0xf9, 0x4d, 0xc6, 0xec, 0xdf, 0x2c,
	0xd2, 0x8c, 0x62, 0x2e, 0xb4, 0x94, 0x67, 0xfd, 0x95, 0xdb, 0x95, 0xd2, 0xc0, 0x77, 0x2b, 0x21,
	0x6f, 0x11, 0xa2, 0x32, 0x72, 0x96, 0x31, 0x55, 0xd4, 0x2f, 0xb0, 0xbb, 0x9b, 0xb9, 0xe5, 0x20,
	0xb4, 0x1f, 0x90, 0x46, 0x1c, 0x9d, 0xf9, 0x69, 0x98, 0x80, 0xf3, 0x25, 0x82, 0x66, 0x6d, 0xf7,
	0xc8, 0x5d, 0x1c, 0x01, 0x9c, 0xaf, 0xb0, 0x15, 0xf4, 0x42, 0xb5, 0x31, 0x87, 0x88, 0x06, 0xd9,
	0x45, 0x2a, 0x9d, 0xaf, 0x75, 0x1b, 0x73, 0x88, 0x9e, 0xab, 0xb5, 0x3a, 0x2f, 0x9c, 0x05, 0x78,
	0xae, 0xb2, 0xf1, 0xcf, 0x10, 0x48, 0xf5, 0xd2, 0x7d, 0x7d, 0x5e, 0x38, 0x0b, 0x4e, 0x41, 0x7e,
	0x8b, 0xe6, 0x83, 0x70, 0xbc, 0x8a, 0xff, 0x41, 0x3c, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0x54, 0xee, 0xe7, 0x5a, 0x0c, 0x00, 0x00,
}
