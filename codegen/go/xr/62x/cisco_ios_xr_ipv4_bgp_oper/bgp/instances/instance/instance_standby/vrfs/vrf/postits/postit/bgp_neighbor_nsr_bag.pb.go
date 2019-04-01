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
// source: bgp_neighbor_nsr_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_vrfs_vrf_postits_postit

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

type BgpNeighborNsrBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpNeighborNsrBag_KEYS) Reset()         { *m = BgpNeighborNsrBag_KEYS{} }
func (m *BgpNeighborNsrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpNeighborNsrBag_KEYS) ProtoMessage()    {}
func (*BgpNeighborNsrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee9558f6ccac82e, []int{0}
}

func (m *BgpNeighborNsrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNeighborNsrBag_KEYS.Unmarshal(m, b)
}
func (m *BgpNeighborNsrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNeighborNsrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpNeighborNsrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNeighborNsrBag_KEYS.Merge(m, src)
}
func (m *BgpNeighborNsrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpNeighborNsrBag_KEYS.Size(m)
}
func (m *BgpNeighborNsrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNeighborNsrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNeighborNsrBag_KEYS proto.InternalMessageInfo

func (m *BgpNeighborNsrBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpNeighborNsrBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNeighborNsrBag_KEYS) GetNeighborAddress() string {
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
	return fileDescriptor_bee9558f6ccac82e, []int{1}
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
	return fileDescriptor_bee9558f6ccac82e, []int{2}
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
	return fileDescriptor_bee9558f6ccac82e, []int{3}
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
	return fileDescriptor_bee9558f6ccac82e, []int{4}
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
	return fileDescriptor_bee9558f6ccac82e, []int{5}
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

type BgpNotfntype_ struct {
	TimeSinceLastNotification uint32   `protobuf:"varint,1,opt,name=time_since_last_notification,json=timeSinceLastNotification,proto3" json:"time_since_last_notification,omitempty"`
	NotificationErrorCode     uint32   `protobuf:"varint,2,opt,name=notification_error_code,json=notificationErrorCode,proto3" json:"notification_error_code,omitempty"`
	NotificationErrorSubcode  uint32   `protobuf:"varint,3,opt,name=notification_error_subcode,json=notificationErrorSubcode,proto3" json:"notification_error_subcode,omitempty"`
	LastNotificationData      []uint32 `protobuf:"varint,4,rep,packed,name=last_notification_data,json=lastNotificationData,proto3" json:"last_notification_data,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *BgpNotfntype_) Reset()         { *m = BgpNotfntype_{} }
func (m *BgpNotfntype_) String() string { return proto.CompactTextString(m) }
func (*BgpNotfntype_) ProtoMessage()    {}
func (*BgpNotfntype_) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee9558f6ccac82e, []int{6}
}

func (m *BgpNotfntype_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNotfntype_.Unmarshal(m, b)
}
func (m *BgpNotfntype_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNotfntype_.Marshal(b, m, deterministic)
}
func (m *BgpNotfntype_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNotfntype_.Merge(m, src)
}
func (m *BgpNotfntype_) XXX_Size() int {
	return xxx_messageInfo_BgpNotfntype_.Size(m)
}
func (m *BgpNotfntype_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNotfntype_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNotfntype_ proto.InternalMessageInfo

func (m *BgpNotfntype_) GetTimeSinceLastNotification() uint32 {
	if m != nil {
		return m.TimeSinceLastNotification
	}
	return 0
}

func (m *BgpNotfntype_) GetNotificationErrorCode() uint32 {
	if m != nil {
		return m.NotificationErrorCode
	}
	return 0
}

func (m *BgpNotfntype_) GetNotificationErrorSubcode() uint32 {
	if m != nil {
		return m.NotificationErrorSubcode
	}
	return 0
}

func (m *BgpNotfntype_) GetLastNotificationData() []uint32 {
	if m != nil {
		return m.LastNotificationData
	}
	return nil
}

type BgpPostitInfo_ struct {
	PostitType           uint32         `protobuf:"varint,1,opt,name=postit_type,json=postitType,proto3" json:"postit_type,omitempty"`
	PostitSubtype        uint32         `protobuf:"varint,2,opt,name=postit_subtype,json=postitSubtype,proto3" json:"postit_subtype,omitempty"`
	PostitFlags          uint32         `protobuf:"varint,3,opt,name=postit_flags,json=postitFlags,proto3" json:"postit_flags,omitempty"`
	PostitAfName         string         `protobuf:"bytes,4,opt,name=postit_af_name,json=postitAfName,proto3" json:"postit_af_name,omitempty"`
	PostitTs             uint64         `protobuf:"varint,5,opt,name=postit_ts,json=postitTs,proto3" json:"postit_ts,omitempty"`
	PostitDirection      uint32         `protobuf:"varint,6,opt,name=postit_direction,json=postitDirection,proto3" json:"postit_direction,omitempty"`
	PostitId             uint32         `protobuf:"varint,7,opt,name=postit_id,json=postitId,proto3" json:"postit_id,omitempty"`
	PeerId               uint32         `protobuf:"varint,8,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	PostitPolicy         string         `protobuf:"bytes,9,opt,name=postit_policy,json=postitPolicy,proto3" json:"postit_policy,omitempty"`
	PostitReset          string         `protobuf:"bytes,10,opt,name=postit_reset,json=postitReset,proto3" json:"postit_reset,omitempty"`
	PostitNotification   *BgpNotfntype_ `protobuf:"bytes,11,opt,name=postit_notification,json=postitNotification,proto3" json:"postit_notification,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BgpPostitInfo_) Reset()         { *m = BgpPostitInfo_{} }
func (m *BgpPostitInfo_) String() string { return proto.CompactTextString(m) }
func (*BgpPostitInfo_) ProtoMessage()    {}
func (*BgpPostitInfo_) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee9558f6ccac82e, []int{7}
}

func (m *BgpPostitInfo_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPostitInfo_.Unmarshal(m, b)
}
func (m *BgpPostitInfo_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPostitInfo_.Marshal(b, m, deterministic)
}
func (m *BgpPostitInfo_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPostitInfo_.Merge(m, src)
}
func (m *BgpPostitInfo_) XXX_Size() int {
	return xxx_messageInfo_BgpPostitInfo_.Size(m)
}
func (m *BgpPostitInfo_) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPostitInfo_.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPostitInfo_ proto.InternalMessageInfo

func (m *BgpPostitInfo_) GetPostitType() uint32 {
	if m != nil {
		return m.PostitType
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitSubtype() uint32 {
	if m != nil {
		return m.PostitSubtype
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitFlags() uint32 {
	if m != nil {
		return m.PostitFlags
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitAfName() string {
	if m != nil {
		return m.PostitAfName
	}
	return ""
}

func (m *BgpPostitInfo_) GetPostitTs() uint64 {
	if m != nil {
		return m.PostitTs
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitDirection() uint32 {
	if m != nil {
		return m.PostitDirection
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitId() uint32 {
	if m != nil {
		return m.PostitId
	}
	return 0
}

func (m *BgpPostitInfo_) GetPeerId() uint32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *BgpPostitInfo_) GetPostitPolicy() string {
	if m != nil {
		return m.PostitPolicy
	}
	return ""
}

func (m *BgpPostitInfo_) GetPostitReset() string {
	if m != nil {
		return m.PostitReset
	}
	return ""
}

func (m *BgpPostitInfo_) GetPostitNotification() *BgpNotfntype_ {
	if m != nil {
		return m.PostitNotification
	}
	return nil
}

type BgpNeighborNsrBag struct {
	VrfName              string            `protobuf:"bytes,50,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	NeighborsAddress     *BgpAddrtype      `protobuf:"bytes,51,opt,name=neighbors_address,json=neighborsAddress,proto3" json:"neighbors_address,omitempty"`
	NeighborAf           []bool            `protobuf:"varint,52,rep,packed,name=neighbor_af,json=neighborAf,proto3" json:"neighbor_af,omitempty"`
	ConnectState         uint32            `protobuf:"varint,53,opt,name=connect_state,json=connectState,proto3" json:"connect_state,omitempty"`
	StartTime            uint32            `protobuf:"varint,54,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	PostitInfo           []*BgpPostitInfo_ `protobuf:"bytes,55,rep,name=postit_info,json=postitInfo,proto3" json:"postit_info,omitempty"`
	NsrState             string            `protobuf:"bytes,56,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BgpNeighborNsrBag) Reset()         { *m = BgpNeighborNsrBag{} }
func (m *BgpNeighborNsrBag) String() string { return proto.CompactTextString(m) }
func (*BgpNeighborNsrBag) ProtoMessage()    {}
func (*BgpNeighborNsrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bee9558f6ccac82e, []int{8}
}

func (m *BgpNeighborNsrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpNeighborNsrBag.Unmarshal(m, b)
}
func (m *BgpNeighborNsrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpNeighborNsrBag.Marshal(b, m, deterministic)
}
func (m *BgpNeighborNsrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpNeighborNsrBag.Merge(m, src)
}
func (m *BgpNeighborNsrBag) XXX_Size() int {
	return xxx_messageInfo_BgpNeighborNsrBag.Size(m)
}
func (m *BgpNeighborNsrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpNeighborNsrBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpNeighborNsrBag proto.InternalMessageInfo

func (m *BgpNeighborNsrBag) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *BgpNeighborNsrBag) GetNeighborsAddress() *BgpAddrtype {
	if m != nil {
		return m.NeighborsAddress
	}
	return nil
}

func (m *BgpNeighborNsrBag) GetNeighborAf() []bool {
	if m != nil {
		return m.NeighborAf
	}
	return nil
}

func (m *BgpNeighborNsrBag) GetConnectState() uint32 {
	if m != nil {
		return m.ConnectState
	}
	return 0
}

func (m *BgpNeighborNsrBag) GetStartTime() uint32 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *BgpNeighborNsrBag) GetPostitInfo() []*BgpPostitInfo_ {
	if m != nil {
		return m.PostitInfo
	}
	return nil
}

func (m *BgpNeighborNsrBag) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func init() {
	proto.RegisterType((*BgpNeighborNsrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_neighbor_nsr_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_addrtype")
	proto.RegisterType((*BgpNotfntype_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_notfntype_")
	proto.RegisterType((*BgpPostitInfo_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_postit_info_")
	proto.RegisterType((*BgpNeighborNsrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.vrfs.vrf.postits.postit.bgp_neighbor_nsr_bag")
}

func init() { proto.RegisterFile("bgp_neighbor_nsr_bag.proto", fileDescriptor_bee9558f6ccac82e) }

var fileDescriptor_bee9558f6ccac82e = []byte{
	// 1096 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x97, 0x49, 0x69, 0xd3, 0x49, 0xd2, 0xa6, 0xd3, 0x7f, 0x6e, 0x01, 0x6d, 0x09, 0x20, 0xba,
	0xd2, 0x2a, 0x12, 0xdd, 0xe0, 0xdd, 0xc3, 0x4a, 0x68, 0xb5, 0xdb, 0x95, 0x2a, 0xda, 0x0a, 0x25,
	0xbd, 0xc0, 0x65, 0xe4, 0xd8, 0xe3, 0xae, 0x91, 0x63, 0x9b, 0x99, 0x69, 0x4a, 0xbf, 0x00, 0x9c,
	0x38, 0x72, 0x40, 0x1c, 0xb8, 0xf2, 0x9d, 0xf8, 0x2a, 0x20, 0xa1, 0xf7, 0x66, 0x3c, 0xfe, 0x93,
	0x80, 0xb4, 0x87, 0xf6, 0xd2, 0xd8, 0xef, 0xf7, 0x7b, 0xcf, 0xef, 0xdf, 0xcc, 0x7b, 0x25, 0x87,
	0xd3, 0xeb, 0x9c, 0xa5, 0x3c, 0xbe, 0x7e, 0x3b, 0xcd, 0x04, 0x4b, 0xa5, 0x60, 0x53, 0xff, 0x7a,
	0x98, 0x8b, 0x4c, 0x65, 0xf4, 0xbb, 0x20, 0x96, 0x41, 0xc6, 0xe2, 0x4c, 0xb2, 0x1f, 0x05, 0x8b,
	0xf3, 0xf9, 0x88, 0x01, 0x3b, 0xcb, 0xb9, 0x18, 0x4e, 0xaf, 0xf3, 0x61, 0x9c, 0x4a, 0xe5, 0xa7,
	0x01, 0x97, 0xf6, 0xc9, 0x3e, 0x30, 0xf8, 0x09, 0xa7, 0x77, 0xc3, 0xb9, 0x88, 0x24, 0xfc, 0x19,
	0xe6, 0x99, 0x54, 0xb1, 0x92, 0xe6, 0x77, 0xf0, 0x93, 0x43, 0x0e, 0x96, 0x7d, 0x9a, 0x7d, 0x7d,
	0xfa, 0xed, 0x84, 0x7e, 0x42, 0x7a, 0xd6, 0x52, 0xea, 0xcf, 0xb8, 0xeb, 0x1c, 0x39, 0xc7, 0xeb,
	0xe3, 0x6e, 0x21, 0xbc, 0xf4, 0x67, 0x9c, 0x1e, 0x90, 0xf6, 0x5c, 0x44, 0x1a, 0x7f, 0x0f, 0xf1,
	0xb5, 0xb9, 0x88, 0x10, 0x7a, 0x4c, 0xfa, 0xd6, 0xb0, 0x1f, 0x86, 0x82, 0x4b, 0xe9, 0xb6, 0x90,
	0xb2, 0x59, 0xc8, 0x5f, 0x6a, 0xf1, 0xe0, 0x19, 0xe9, 0x83, 0x1f, 0xc9, 0xc9, 0x3c, 0x4f, 0x91,
	0xcb, 0x14, 0x7c, 0xbe, 0x7c, 0x07, 0x5d, 0xe7, 0xa8, 0x75, 0xdc, 0x1b, 0x77, 0x51, 0x58, 0x28,
	0xbe, 0x20, 0xbb, 0xa5, 0xe2, 0x4c, 0xe6, 0xb7, 0xef, 0xa4, 0x7d, 0x49, 0x5c, 0xd0, 0xc6, 0xa4,
	0x4a, 0x91, 0x67, 0x49, 0x1c, 0xdc, 0x15, 0x06, 0x4e, 0xc8, 0xee, 0xa2, 0xbc, 0x34, 0xb4, 0x0d,
	0xe0, 0xc4, 0x60, 0x8b, 0xf6, 0xbc, 0xff, 0xb0, 0xe7, 0xfd, 0x9f, 0x3d, 0xaf, 0x69, 0xef, 0xef,
	0x2e, 0xe9, 0x82, 0x41, 0xa0, 0xaa, 0xbb, 0x9c, 0xd3, 0x3e, 0x69, 0xf9, 0x51, 0x6c, 0x0a, 0x01,
	0x8f, 0xf4, 0x63, 0xd2, 0x45, 0x37, 0x0b, 0x6b, 0xba, 0x06, 0x1d, 0x90, 0x19, 0x2b, 0xf4, 0x09,
	0xa1, 0x48, 0x99, 0x05, 0xbe, 0x54, 0x8d, 0x4a, 0xf4, 0x01, 0xb9, 0x00, 0xa0, 0xc9, 0x4e, 0xfc,
	0x29, 0x4f, 0x2c, 0x7b, 0xa5, 0x64, 0x9f, 0x03, 0x50, 0xb0, 0x87, 0x04, 0x13, 0xc1, 0xd4, 0x4d,
	0x9a, 0x56, 0xe8, 0xef, 0x23, 0x7d, 0x0b, 0xa0, 0x2b, 0x44, 0x0a, 0xfe, 0xe7, 0x64, 0x13, 0x84,
	0xb3, 0xb0, 0x74, 0x64, 0x15, 0xb9, 0x1b, 0x46, 0xdc, 0x20, 0x56, 0x2b, 0xb8, 0x56, 0x12, 0xcb,
	0x1a, 0xd2, 0x2f, 0xc8, 0x8e, 0x91, 0xf8, 0x3a, 0x42, 0xc3, 0x6e, 0x23, 0x7b, 0xbb, 0xc0, 0x2e,
	0x4a, 0xc8, 0xe4, 0xcc, 0xb3, 0x86, 0xd7, 0x6d, 0xce, 0xbc, 0x7a, 0x16, 0xbc, 0x46, 0xce, 0x88,
	0xcd, 0x82, 0xb7, 0x24, 0x67, 0x5e, 0x23, 0x67, 0x9d, 0x92, 0x5d, 0xcb, 0x99, 0x0e, 0xcd, 0xab,
	0x86, 0xd6, 0xb5, 0xa1, 0x79, 0x95, 0xd0, 0x4c, 0xcb, 0x60, 0x6b, 0xd7, 0xfc, 0xe8, 0xd9, 0xd8,
	0x00, 0xac, 0xb9, 0xf2, 0x9b, 0x43, 0xb6, 0xb0, 0xc7, 0xe7, 0x79, 0x22, 0xad, 0xc2, 0xc6, 0x91,
	0x73, 0xdc, 0x39, 0x49, 0x86, 0xf7, 0x77, 0x97, 0x0c, 0x9b, 0xe7, 0x77, 0xdc, 0xb7, 0x6e, 0x54,
	0xe2, 0x11, 0x8a, 0x05, 0x59, 0x2a, 0x95, 0xf0, 0xe3, 0xb4, 0x8c, 0x67, 0x53, 0xc7, 0x23, 0xd4,
	0x2b, 0x8b, 0x15, 0x3a, 0x8f, 0x09, 0x26, 0x70, 0x56, 0xcd, 0x56, 0x5f, 0x5f, 0x22, 0x85, 0xbc,
	0x4e, 0x1d, 0xd5, 0xa8, 0x5b, 0x96, 0x3a, 0xaa, 0x52, 0x9f, 0x10, 0xaa, 0x7d, 0xe5, 0x55, 0x32,
	0xd5, 0x05, 0x43, 0xe4, 0xb4, 0xc2, 0x1e, 0x90, 0x5e, 0x22, 0x59, 0x25, 0x9d, 0xdb, 0xba, 0x61,
	0x12, 0x79, 0x6e, 0x63, 0xfb, 0xc3, 0x29, 0x4c, 0xda, 0x5b, 0x08, 0x98, 0x3b, 0x98, 0xf8, 0x1f,
	0x1e, 0x26, 0xf1, 0x95, 0xfb, 0xcf, 0x44, 0x71, 0x21, 0xf3, 0xdb, 0x7a, 0x37, 0x8d, 0x58, 0x94,
	0x64, 0xb7, 0x32, 0xe7, 0x81, 0xf5, 0x71, 0xb7, 0x3c, 0x29, 0x6f, 0x0c, 0xd6, 0xe8, 0xc0, 0x45,
	0x9d, 0xbd, 0xb2, 0x03, 0x9b, 0x3a, 0xcf, 0x89, 0x5b, 0x9c, 0xdc, 0x05, 0xb5, 0x7d, 0x54, 0xdb,
	0x33, 0xf8, 0x72, 0x4d, 0x6f, 0xa9, 0xa6, 0x6b, 0x35, 0xbd, 0x25, 0x9a, 0x7f, 0x3a, 0x3a, 0x38,
	0x29, 0x58, 0xe3, 0x76, 0x3d, 0xc0, 0x02, 0xa8, 0xfb, 0x2e, 0xc0, 0xb2, 0x11, 0xa2, 0x53, 0x2a,
	0xc5, 0x37, 0xd5, 0x3b, 0xbd, 0x70, 0xd5, 0x5b, 0x74, 0xf5, 0xf0, 0xc1, 0x5c, 0xf5, 0x96, 0xba,
	0xea, 0x35, 0x5c, 0x1d, 0xfc, 0xe3, 0x90, 0x0d, 0x5c, 0x0f, 0x32, 0x15, 0xa5, 0x30, 0x7f, 0x18,
	0xfd, 0x8a, 0x7c, 0xa8, 0xe2, 0x19, 0x67, 0x32, 0x86, 0xcf, 0x25, 0x70, 0x29, 0xa5, 0x99, 0x8a,
	0xa3, 0x38, 0xf0, 0x55, 0x9c, 0xa5, 0x38, 0x99, 0x7a, 0xe3, 0x03, 0xe0, 0x4c, 0x80, 0x72, 0xee,
	0x4b, 0x75, 0x59, 0x21, 0x50, 0x8f, 0xec, 0x57, 0x15, 0x18, 0x17, 0x22, 0x13, 0x2c, 0xc8, 0x42,
	0xbd, 0x3e, 0xf4, 0xc6, 0xbb, 0x55, 0xf8, 0x14, 0xd0, 0x57, 0x59, 0xc8, 0xe9, 0x0b, 0x72, 0xb8,
	0x44, 0x4f, 0xde, 0x4c, 0x51, 0xb5, 0x85, 0xaa, 0xee, 0x82, 0xea, 0x44, 0xe3, 0x74, 0x44, 0xf6,
	0x16, 0x7c, 0x65, 0xa1, 0xaf, 0x7c, 0x77, 0x05, 0xa7, 0xef, 0x4e, 0xd2, 0xf0, 0xf3, 0xb5, 0xaf,
	0xfc, 0xc1, 0xcf, 0x2b, 0x7a, 0x2d, 0xd1, 0xc9, 0x63, 0x71, 0x1a, 0x65, 0x8c, 0x3e, 0x22, 0x1d,
	0xf3, 0x0e, 0x19, 0x31, 0x01, 0x13, 0x2d, 0xba, 0x82, 0x19, 0xfd, 0x19, 0xd9, 0x30, 0x04, 0x79,
	0x33, 0x45, 0x8e, 0x0e, 0xac, 0xa7, 0xa5, 0x13, 0x2d, 0x84, 0x21, 0x64, 0x68, 0x51, 0xe2, 0x5f,
	0x4b, 0x13, 0x82, 0xb1, 0xfd, 0x06, 0x44, 0xf4, 0x53, 0x6b, 0xc9, 0x37, 0x1b, 0x96, 0x1e, 0xc3,
	0x46, 0xf1, 0xa5, 0x5e, 0xb3, 0x3e, 0x20, 0xeb, 0x85, 0x43, 0x7a, 0xf0, 0xae, 0x8c, 0xdb, 0xc6,
	0x1d, 0xbc, 0x13, 0x0d, 0x18, 0xc6, 0x82, 0x07, 0x58, 0xa3, 0x55, 0xfc, 0xd2, 0xa6, 0x96, 0xbf,
	0x2e, 0xc4, 0x15, 0x3b, 0x71, 0x88, 0xb3, 0xb6, 0x57, 0xd8, 0x39, 0x0b, 0xe9, 0x3e, 0x59, 0xcb,
	0x39, 0x17, 0x00, 0xb5, 0x11, 0x5a, 0x85, 0xd7, 0xb3, 0x10, 0xf6, 0x2c, 0xa3, 0xa5, 0x1b, 0xca,
	0x0c, 0x53, 0xe3, 0xa2, 0xee, 0xa7, 0x4a, 0xac, 0x82, 0x4b, 0xae, 0xcc, 0x1c, 0x35, 0xb1, 0x8e,
	0x41, 0x44, 0x7f, 0x77, 0xc8, 0xb6, 0xe1, 0xd4, 0x1a, 0xaa, 0x83, 0x87, 0xe2, 0xfb, 0xfb, 0x3e,
	0x14, 0x65, 0x8b, 0x8f, 0xa9, 0x16, 0x57, 0xbb, 0x61, 0xf0, 0x57, 0x8b, 0xec, 0x2c, 0x5b, 0x94,
	0x6b, 0xeb, 0xef, 0x49, 0x7d, 0xfd, 0xfd, 0xd5, 0x21, 0x5b, 0x05, 0xbf, 0x1c, 0x1d, 0x4f, 0x31,
	0x9e, 0xb7, 0xf7, 0x1d, 0x4f, 0xb1, 0x31, 0x8e, 0xed, 0x0a, 0x6e, 0x27, 0xd5, 0x23, 0xd2, 0x29,
	0xd7, 0xf2, 0xc8, 0x1d, 0x1d, 0xb5, 0x8e, 0xdb, 0x63, 0x62, 0x37, 0xf2, 0x08, 0x4a, 0x1a, 0x64,
	0x69, 0xca, 0x03, 0x05, 0x9f, 0x52, 0xdc, 0xfd, 0x12, 0x2b, 0xde, 0x35, 0xc2, 0x09, 0xc8, 0xe8,
	0x47, 0x84, 0x48, 0xe5, 0x0b, 0xc5, 0xe0, 0xa8, 0xbb, 0x1e, 0x32, 0xd6, 0x51, 0x72, 0x15, 0xcf,
	0x38, 0xfd, 0xc5, 0xb1, 0xc7, 0x04, 0x8e, 0x8d, 0xfb, 0xec, 0xa8, 0xf5, 0x10, 0x0b, 0x48, 0xf5,
	0xa4, 0x16, 0x87, 0xf2, 0x2c, 0x8d, 0x32, 0x68, 0x6e, 0x28, 0x99, 0x8e, 0xe7, 0x39, 0x16, 0xaa,
	0x9d, 0x4a, 0x81, 0xb1, 0x4c, 0x57, 0xf1, 0x3f, 0xad, 0xa7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x2b, 0x6d, 0xff, 0x87, 0x0d, 0x00, 0x00,
}
