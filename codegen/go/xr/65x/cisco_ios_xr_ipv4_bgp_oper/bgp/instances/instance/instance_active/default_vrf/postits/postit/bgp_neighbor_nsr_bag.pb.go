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

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_default_vrf_postits_postit

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
	NeighborAddress      string   `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
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
	proto.RegisterType((*BgpNeighborNsrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_neighbor_nsr_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_addrtype")
	proto.RegisterType((*BgpNotfntype_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_notfntype_")
	proto.RegisterType((*BgpPostitInfo_)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_postit_info_")
	proto.RegisterType((*BgpNeighborNsrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.default_vrf.postits.postit.bgp_neighbor_nsr_bag")
}

func init() { proto.RegisterFile("bgp_neighbor_nsr_bag.proto", fileDescriptor_bee9558f6ccac82e) }

var fileDescriptor_bee9558f6ccac82e = []byte{
	// 1100 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x97, 0x49, 0x69, 0xd3, 0x49, 0xd2, 0xa6, 0xd3, 0x7f, 0x6e, 0x01, 0x6d, 0x08, 0x20, 0xb2,
	0xd2, 0x2a, 0x12, 0xdd, 0xe0, 0xdd, 0xc3, 0x4a, 0x68, 0xb5, 0xdb, 0x95, 0x2a, 0xda, 0x0a, 0x39,
	0x15, 0x12, 0x12, 0xd2, 0x68, 0x62, 0x8f, 0x83, 0xc1, 0xb1, 0xad, 0x99, 0x69, 0xca, 0x7e, 0x03,
	0x38, 0x71, 0xe4, 0x02, 0x07, 0xce, 0x7c, 0x28, 0x3e, 0x09, 0x12, 0x9a, 0x37, 0xe3, 0xb1, 0xe3,
	0x18, 0x24, 0x0e, 0xed, 0xa9, 0xf1, 0x7b, 0xbf, 0xdf, 0x9b, 0xf7, 0x6f, 0xe6, 0xbd, 0xa2, 0xd3,
	0xd9, 0x3c, 0x27, 0x29, 0x8b, 0xe7, 0xdf, 0xcd, 0x32, 0x4e, 0x52, 0xc1, 0xc9, 0x8c, 0xce, 0xc7,
	0x39, 0xcf, 0x64, 0x86, 0xbf, 0x0d, 0x62, 0x11, 0x64, 0x24, 0xce, 0x04, 0xf9, 0x91, 0x93, 0x38,
	0x5f, 0x4e, 0x88, 0x42, 0x67, 0x39, 0xe3, 0xe3, 0xd9, 0x3c, 0x1f, 0xc7, 0xa9, 0x90, 0x34, 0x0d,
	0x98, 0xb0, 0xbf, 0xec, 0x0f, 0x42, 0x03, 0x19, 0x2f, 0xd9, 0x38, 0x64, 0x11, 0xbd, 0x4d, 0x24,
	0x59, 0xf2, 0x68, 0x9c, 0x67, 0x42, 0xc6, 0x52, 0x98, 0xbf, 0xc3, 0x1f, 0xd0, 0x49, 0xd3, 0xd9,
	0xe4, 0xcb, 0xf3, 0x6f, 0xa6, 0xf8, 0x23, 0xd4, 0xb3, 0xa6, 0x52, 0xba, 0x60, 0xae, 0x33, 0x70,
	0x46, 0xdb, 0x7e, 0xb7, 0x10, 0x5e, 0xd3, 0x05, 0xc3, 0x8f, 0x51, 0xdf, 0xb2, 0x69, 0x18, 0x72,
	0x26, 0x84, 0xfb, 0x0e, 0xe0, 0x76, 0x0b, 0xf9, 0x4b, 0x2d, 0x1e, 0x3e, 0x43, 0x7d, 0x75, 0x58,
	0x72, 0xb6, 0xcc, 0x53, 0xc0, 0x12, 0xa9, 0xce, 0x28, 0xbf, 0x15, 0xd7, 0x19, 0xb4, 0x46, 0x3d,
	0xbf, 0x0b, 0xc2, 0x82, 0xf8, 0x02, 0x1d, 0x96, 0xc4, 0x85, 0xc8, 0xef, 0xfe, 0x17, 0xfb, 0x1a,
	0xb9, 0x8a, 0x0d, 0xa9, 0x13, 0x3c, 0xcf, 0x92, 0x38, 0x78, 0x5b, 0x18, 0x38, 0x43, 0x87, 0xeb,
	0xf2, 0xd2, 0xd0, 0xbe, 0x52, 0x4e, 0x8d, 0x6e, 0xdd, 0x9e, 0xf7, 0x2f, 0xf6, 0xbc, 0xff, 0xb2,
	0xe7, 0xd5, 0xed, 0xfd, 0xd4, 0x43, 0x5d, 0x65, 0x50, 0x41, 0xe5, 0xdb, 0x9c, 0xe1, 0x3e, 0x6a,
	0xd1, 0x28, 0x36, 0xd9, 0x56, 0x3f, 0xf1, 0x87, 0xa8, 0x0b, 0x6e, 0xae, 0x26, 0xb8, 0xa3, 0x64,
	0xc6, 0x0a, 0x7e, 0x82, 0x30, 0x40, 0x16, 0x01, 0x15, 0xd2, 0x02, 0x5b, 0x00, 0xec, 0x2b, 0xcd,
	0x95, 0x52, 0xd4, 0xd1, 0x09, 0x9d, 0xb1, 0xc4, 0xa2, 0x37, 0x4a, 0xf4, 0xa5, 0x52, 0x14, 0xe8,
	0x31, 0x82, 0x44, 0x10, 0x79, 0x9b, 0xa6, 0x15, 0xf8, 0xbb, 0x00, 0xdf, 0x53, 0xaa, 0x1b, 0xd0,
	0x14, 0xf8, 0x11, 0xea, 0x6b, 0x5f, 0xc2, 0xd2, 0x93, 0x4d, 0x00, 0xef, 0x80, 0x27, 0xa1, 0xf5,
	0xe3, 0x53, 0xb4, 0xab, 0x24, 0xd5, 0x12, 0x6e, 0x95, 0xc0, 0xb2, 0x88, 0xf8, 0x33, 0x74, 0x60,
	0x24, 0x54, 0x87, 0x68, 0xd0, 0x6d, 0x40, 0xef, 0x17, 0xba, 0xab, 0x52, 0x65, 0x92, 0xe6, 0x59,
	0xc3, 0xdb, 0x36, 0x69, 0xde, 0x6a, 0x1a, 0xbc, 0x5a, 0xd2, 0x90, 0x4d, 0x83, 0xd7, 0x90, 0x34,
	0xaf, 0x96, 0xb4, 0x4e, 0x89, 0x5e, 0x49, 0x9a, 0x0e, 0xcd, 0xab, 0x86, 0xd6, 0xb5, 0xa1, 0x79,
	0x95, 0xd0, 0x4c, 0xcf, 0x40, 0x6f, 0xaf, 0xf8, 0xd1, 0xb3, 0xb1, 0x29, 0xe5, 0x8a, 0x2b, 0xbf,
	0x39, 0x08, 0xeb, 0xce, 0x5f, 0xe6, 0x89, 0xb0, 0x8c, 0x9d, 0x81, 0x33, 0xea, 0x9c, 0xa5, 0xe3,
	0xfb, 0x7c, 0x33, 0xc6, 0xf5, 0x3b, 0xec, 0xf7, 0xe1, 0xeb, 0xeb, 0x3c, 0x11, 0x95, 0x90, 0xb8,
	0x24, 0x41, 0x96, 0x0a, 0xc9, 0x69, 0x9c, 0x96, 0x21, 0xed, 0xea, 0x90, 0xb8, 0x7c, 0x65, 0x75,
	0x05, 0xe7, 0x31, 0x34, 0x8d, 0xb7, 0xa8, 0x26, 0xac, 0xaf, 0x1f, 0x92, 0x42, 0xbe, 0x0a, 0x9d,
	0xac, 0x40, 0xf7, 0x2c, 0x74, 0x52, 0x85, 0x3e, 0x29, 0xf2, 0xc4, 0xaa, 0x60, 0xac, 0x6b, 0x06,
	0x9a, 0xf3, 0x0a, 0x7a, 0x88, 0x7a, 0x89, 0x20, 0x95, 0x84, 0xee, 0xeb, 0x9e, 0x49, 0xc4, 0xa5,
	0x8d, 0xed, 0x0f, 0x9b, 0x7a, 0xfb, 0x12, 0x29, 0xe4, 0x01, 0xa4, 0x5e, 0x3c, 0x54, 0xea, 0x2b,
	0xaf, 0xa0, 0x89, 0xe3, 0x4a, 0xe4, 0x77, 0xab, 0x2d, 0x35, 0x21, 0x51, 0x92, 0xdd, 0x89, 0x9c,
	0x05, 0xd6, 0xcb, 0xc3, 0xf2, 0xba, 0xbc, 0x31, 0xba, 0x5a, 0x1b, 0xae, 0x73, 0x8e, 0xca, 0x36,
	0xac, 0x73, 0x9e, 0x23, 0xb7, 0xb8, 0xbe, 0x6b, 0xb4, 0x63, 0xa0, 0x1d, 0x19, 0x7d, 0x33, 0xd3,
	0x6b, 0x64, 0xba, 0x96, 0xe9, 0x35, 0x30, 0xff, 0x74, 0xd0, 0x91, 0x79, 0xb3, 0x49, 0xed, 0x91,
	0x3d, 0x81, 0x1a, 0x2c, 0xef, 0xbf, 0x06, 0x4d, 0xb3, 0xa4, 0x18, 0x16, 0x5f, 0x55, 0x1f, 0xf7,
	0xc2, 0x5b, 0xaf, 0xc1, 0xdb, 0xd3, 0x07, 0xf4, 0xd6, 0x6b, 0xf4, 0xd6, 0xab, 0x79, 0x3b, 0xfc,
	0xdb, 0x41, 0x3b, 0xb0, 0x0f, 0x64, 0x32, 0x4a, 0xd5, 0x2c, 0x22, 0xf8, 0x0b, 0xf4, 0xbe, 0x8c,
	0x17, 0x8c, 0x88, 0x58, 0x1d, 0x98, 0xa8, 0xf7, 0x29, 0xcd, 0x64, 0x1c, 0xc5, 0x01, 0x95, 0x71,
	0x96, 0xc2, 0x94, 0xea, 0xf9, 0x27, 0x0a, 0x33, 0x55, 0x90, 0x4b, 0x2a, 0xe4, 0x75, 0x05, 0x80,
	0x3d, 0x74, 0x5c, 0x25, 0x10, 0xc6, 0x79, 0xc6, 0x49, 0x90, 0x85, 0x0c, 0xc6, 0x58, 0xcf, 0x3f,
	0xac, 0xaa, 0xcf, 0x95, 0xf6, 0x55, 0x16, 0x32, 0xfc, 0x02, 0x9d, 0x36, 0xf0, 0xc4, 0xed, 0x0c,
	0xa8, 0x2d, 0xa0, 0xba, 0x6b, 0xd4, 0xa9, 0xd6, 0xe3, 0x09, 0x3a, 0x5a, 0xf3, 0x95, 0x84, 0x54,
	0x52, 0x77, 0x03, 0x26, 0xf1, 0x41, 0x52, 0xf3, 0xf3, 0x35, 0x95, 0x74, 0xf8, 0xf3, 0x86, 0x5e,
	0x51, 0x74, 0xf2, 0x48, 0x9c, 0x46, 0x19, 0xc1, 0x8f, 0x50, 0xc7, 0x7c, 0xab, 0x8c, 0x98, 0x80,
	0x91, 0x16, 0xdd, 0xa8, 0x79, 0xfd, 0x09, 0xda, 0x31, 0x00, 0x71, 0x3b, 0x03, 0x8c, 0x0e, 0xac,
	0xa7, 0xa5, 0x53, 0x2d, 0x54, 0xf3, 0xc8, 0xc0, 0xa2, 0x84, 0xce, 0x85, 0x09, 0xc1, 0xd8, 0x7e,
	0xa3, 0x44, 0xf8, 0x63, 0x6b, 0x89, 0x46, 0x7a, 0xe5, 0xd2, 0x23, 0xd9, 0x10, 0x5f, 0x46, 0xb0,
	0x72, 0xbd, 0x87, 0xb6, 0x0b, 0x87, 0xf4, 0x10, 0xde, 0xf0, 0xdb, 0xc6, 0x1d, 0x78, 0x1b, 0x8d,
	0x32, 0x8c, 0x39, 0x0b, 0xa0, 0x46, 0x9b, 0x70, 0xd2, 0xae, 0x96, 0xbf, 0x2e, 0xc4, 0x15, 0x3b,
	0x71, 0x08, 0x63, 0xb7, 0x57, 0xd8, 0xb9, 0x08, 0xf1, 0x31, 0xda, 0xca, 0x19, 0xe3, 0x4a, 0xd5,
	0x06, 0xd5, 0xa6, 0xfa, 0xbc, 0x08, 0xd5, 0xce, 0x65, 0x58, 0xba, 0xa1, 0xcc, 0x5c, 0x35, 0x2e,
	0xea, 0x7e, 0xaa, 0xc4, 0xca, 0x99, 0x60, 0xd2, 0x8c, 0x54, 0x13, 0xab, 0xaf, 0x44, 0xf8, 0x77,
	0x07, 0xed, 0x1b, 0xcc, 0x4a, 0x43, 0x75, 0xe0, 0x5a, 0x24, 0xf7, 0x7f, 0x2d, 0xca, 0x26, 0xf7,
	0xb1, 0x16, 0x57, 0xfb, 0x61, 0xf8, 0x57, 0x0b, 0x1d, 0x34, 0xed, 0xc6, 0xf8, 0x04, 0xb5, 0x97,
	0xdc, 0x94, 0xe7, 0x0c, 0xe2, 0xda, 0x5a, 0x72, 0x5d, 0x99, 0x5f, 0x1d, 0xb4, 0x57, 0xe0, 0xcb,
	0x21, 0xf2, 0x14, 0x22, 0xfa, 0xfe, 0xfe, 0x23, 0x2a, 0x36, 0x48, 0xdf, 0xae, 0xe4, 0x76, 0x6a,
	0x3d, 0x42, 0x9d, 0x72, 0x4d, 0x8f, 0xdc, 0xc9, 0xa0, 0x35, 0x6a, 0xfb, 0xc8, 0x6e, 0xe8, 0x91,
	0x2a, 0x6b, 0x90, 0xa5, 0x29, 0x0b, 0x24, 0x11, 0x92, 0x4a, 0xe6, 0x7e, 0x0e, 0x55, 0xef, 0x1a,
	0xe1, 0x54, 0xc9, 0xf0, 0x07, 0x08, 0x09, 0x49, 0xb9, 0x24, 0xea, 0xba, 0xbb, 0x1e, 0x20, 0xb6,
	0x41, 0x72, 0x13, 0x2f, 0x18, 0xfe, 0xc5, 0xb1, 0x57, 0x45, 0x5d, 0x1d, 0xf7, 0xd9, 0xa0, 0xf5,
	0x30, 0xeb, 0x48, 0xf5, 0xbe, 0x16, 0x57, 0xf3, 0x22, 0x8d, 0x32, 0xd5, 0xe2, 0xaa, 0x6c, 0x3a,
	0xa2, 0xe7, 0x50, 0xac, 0x76, 0x2a, 0x38, 0x44, 0x33, 0xdb, 0x84, 0xff, 0xb0, 0x9e, 0xfe, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x84, 0x74, 0x32, 0x7f, 0x0d, 0x00, 0x00,
}
