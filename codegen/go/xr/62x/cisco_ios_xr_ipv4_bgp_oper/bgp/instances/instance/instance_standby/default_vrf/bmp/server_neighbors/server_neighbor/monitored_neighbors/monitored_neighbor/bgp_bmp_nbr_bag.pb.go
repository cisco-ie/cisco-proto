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
// source: bgp_bmp_nbr_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_bmp_server_neighbors_server_neighbor_monitored_neighbors_monitored_neighbor

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

type BgpBmpNbrBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	ServerId             uint32   `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpBmpNbrBag_KEYS) Reset()         { *m = BgpBmpNbrBag_KEYS{} }
func (m *BgpBmpNbrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpBmpNbrBag_KEYS) ProtoMessage()    {}
func (*BgpBmpNbrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3893b48767fb191b, []int{0}
}

func (m *BgpBmpNbrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpBmpNbrBag_KEYS.Unmarshal(m, b)
}
func (m *BgpBmpNbrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpBmpNbrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpBmpNbrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpBmpNbrBag_KEYS.Merge(m, src)
}
func (m *BgpBmpNbrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpBmpNbrBag_KEYS.Size(m)
}
func (m *BgpBmpNbrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpBmpNbrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpBmpNbrBag_KEYS proto.InternalMessageInfo

func (m *BgpBmpNbrBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpBmpNbrBag_KEYS) GetServerId() uint32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *BgpBmpNbrBag_KEYS) GetNeighborAddress() string {
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
	return fileDescriptor_3893b48767fb191b, []int{1}
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
	return fileDescriptor_3893b48767fb191b, []int{2}
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
	return fileDescriptor_3893b48767fb191b, []int{3}
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
	return fileDescriptor_3893b48767fb191b, []int{4}
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
	return fileDescriptor_3893b48767fb191b, []int{5}
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

type BgpBmpNbrBag struct {
	BmpNeighborAddress             *BgpAddrtype `protobuf:"bytes,50,opt,name=bmp_neighbor_address,json=bmpNeighborAddress,proto3" json:"bmp_neighbor_address,omitempty"`
	BmpNeighborMsgPending          uint32       `protobuf:"varint,51,opt,name=bmp_neighbor_msg_pending,json=bmpNeighborMsgPending,proto3" json:"bmp_neighbor_msg_pending,omitempty"`
	BmpNeighborMsgDropped          uint32       `protobuf:"varint,52,opt,name=bmp_neighbor_msg_dropped,json=bmpNeighborMsgDropped,proto3" json:"bmp_neighbor_msg_dropped,omitempty"`
	BmpNeighborPeerUpSent          uint32       `protobuf:"varint,53,opt,name=bmp_neighbor_peer_up_sent,json=bmpNeighborPeerUpSent,proto3" json:"bmp_neighbor_peer_up_sent,omitempty"`
	BmpNeighborPeerDownSent        uint32       `protobuf:"varint,54,opt,name=bmp_neighbor_peer_down_sent,json=bmpNeighborPeerDownSent,proto3" json:"bmp_neighbor_peer_down_sent,omitempty"`
	BmpNeighborRouteMonitorSent    uint32       `protobuf:"varint,55,opt,name=bmp_neighbor_route_monitor_sent,json=bmpNeighborRouteMonitorSent,proto3" json:"bmp_neighbor_route_monitor_sent,omitempty"`
	BmpNeighborRouteMonitorEorSent uint32       `protobuf:"varint,56,opt,name=bmp_neighbor_route_monitor_eor_sent,json=bmpNeighborRouteMonitorEorSent,proto3" json:"bmp_neighbor_route_monitor_eor_sent,omitempty"`
	BmpNeighborPathUpdateSent      uint32       `protobuf:"varint,57,opt,name=bmp_neighbor_path_update_sent,json=bmpNeighborPathUpdateSent,proto3" json:"bmp_neighbor_path_update_sent,omitempty"`
	BmpNeighborPathWithdrawSent    uint32       `protobuf:"varint,58,opt,name=bmp_neighbor_path_withdraw_sent,json=bmpNeighborPathWithdrawSent,proto3" json:"bmp_neighbor_path_withdraw_sent,omitempty"`
	BmpNeighborPathUpdateDrop      uint32       `protobuf:"varint,59,opt,name=bmp_neighbor_path_update_drop,json=bmpNeighborPathUpdateDrop,proto3" json:"bmp_neighbor_path_update_drop,omitempty"`
	BmpNeighborPathWithdrawDrop    uint32       `protobuf:"varint,60,opt,name=bmp_neighbor_path_withdraw_drop,json=bmpNeighborPathWithdrawDrop,proto3" json:"bmp_neighbor_path_withdraw_drop,omitempty"`
	BmpNeighborUpdMsgSent          uint32       `protobuf:"varint,61,opt,name=bmp_neighbor_upd_msg_sent,json=bmpNeighborUpdMsgSent,proto3" json:"bmp_neighbor_upd_msg_sent,omitempty"`
	BmpNeighborWdrawMsgSent        uint32       `protobuf:"varint,62,opt,name=bmp_neighbor_wdraw_msg_sent,json=bmpNeighborWdrawMsgSent,proto3" json:"bmp_neighbor_wdraw_msg_sent,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}     `json:"-"`
	XXX_unrecognized               []byte       `json:"-"`
	XXX_sizecache                  int32        `json:"-"`
}

func (m *BgpBmpNbrBag) Reset()         { *m = BgpBmpNbrBag{} }
func (m *BgpBmpNbrBag) String() string { return proto.CompactTextString(m) }
func (*BgpBmpNbrBag) ProtoMessage()    {}
func (*BgpBmpNbrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3893b48767fb191b, []int{6}
}

func (m *BgpBmpNbrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpBmpNbrBag.Unmarshal(m, b)
}
func (m *BgpBmpNbrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpBmpNbrBag.Marshal(b, m, deterministic)
}
func (m *BgpBmpNbrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpBmpNbrBag.Merge(m, src)
}
func (m *BgpBmpNbrBag) XXX_Size() int {
	return xxx_messageInfo_BgpBmpNbrBag.Size(m)
}
func (m *BgpBmpNbrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpBmpNbrBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpBmpNbrBag proto.InternalMessageInfo

func (m *BgpBmpNbrBag) GetBmpNeighborAddress() *BgpAddrtype {
	if m != nil {
		return m.BmpNeighborAddress
	}
	return nil
}

func (m *BgpBmpNbrBag) GetBmpNeighborMsgPending() uint32 {
	if m != nil {
		return m.BmpNeighborMsgPending
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborMsgDropped() uint32 {
	if m != nil {
		return m.BmpNeighborMsgDropped
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPeerUpSent() uint32 {
	if m != nil {
		return m.BmpNeighborPeerUpSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPeerDownSent() uint32 {
	if m != nil {
		return m.BmpNeighborPeerDownSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborRouteMonitorSent() uint32 {
	if m != nil {
		return m.BmpNeighborRouteMonitorSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborRouteMonitorEorSent() uint32 {
	if m != nil {
		return m.BmpNeighborRouteMonitorEorSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPathUpdateSent() uint32 {
	if m != nil {
		return m.BmpNeighborPathUpdateSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPathWithdrawSent() uint32 {
	if m != nil {
		return m.BmpNeighborPathWithdrawSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPathUpdateDrop() uint32 {
	if m != nil {
		return m.BmpNeighborPathUpdateDrop
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborPathWithdrawDrop() uint32 {
	if m != nil {
		return m.BmpNeighborPathWithdrawDrop
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborUpdMsgSent() uint32 {
	if m != nil {
		return m.BmpNeighborUpdMsgSent
	}
	return 0
}

func (m *BgpBmpNbrBag) GetBmpNeighborWdrawMsgSent() uint32 {
	if m != nil {
		return m.BmpNeighborWdrawMsgSent
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpBmpNbrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_bmp_nbr_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_addrtype")
	proto.RegisterType((*BgpBmpNbrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_bmp_nbr_bag")
}

func init() { proto.RegisterFile("bgp_bmp_nbr_bag.proto", fileDescriptor_3893b48767fb191b) }

var fileDescriptor_3893b48767fb191b = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0x41, 0x6f, 0xdc, 0x44,
	0x14, 0x96, 0x1b, 0x68, 0x9b, 0x49, 0xb6, 0xbb, 0x71, 0x92, 0xd6, 0x21, 0x02, 0x42, 0x7a, 0x20,
	0x95, 0x2a, 0x4b, 0xa4, 0x8b, 0x13, 0x20, 0x20, 0x10, 0x09, 0x12, 0x6a, 0x37, 0x8a, 0x36, 0x44,
	0x15, 0xa7, 0x91, 0xbd, 0x9e, 0x6c, 0x2c, 0xd9, 0xe3, 0xd1, 0xcc, 0xec, 0x2e, 0xb9, 0x21, 0xf1,
	0x2f, 0xf8, 0x05, 0x1c, 0xb8, 0x20, 0xf1, 0x6b, 0xe0, 0xc2, 0x95, 0x3b, 0x3f, 0x00, 0xcd, 0x1b,
	0x7b, 0x6c, 0x8f, 0xdd, 0x54, 0xbd, 0xa5, 0xa7, 0xec, 0xbe, 0xf7, 0x7d, 0x9f, 0xdf, 0xf7, 0xde,
	0xf3, 0xec, 0x04, 0x6d, 0x46, 0x53, 0x86, 0xa3, 0x8c, 0x61, 0x1a, 0x71, 0x1c, 0x85, 0x53, 0x9f,
	0xf1, 0x5c, 0xe6, 0xee, 0xaf, 0xce, 0x24, 0x11, 0x93, 0x1c, 0x27, 0xb9, 0xc0, 0x3f, 0x71, 0x9c,
	0xb0, 0xf9, 0x10, 0x2b, 0x64, 0xce, 0x08, 0xf7, 0xa3, 0x29, 0xf3, 0x13, 0x2a, 0x64, 0x48, 0x27,
	0x44, 0x98, 0x4f, 0xe6, 0x03, 0x56, 0x7f, 0xe2, 0xe8, 0xda, 0x8f, 0xc9, 0x65, 0x38, 0x4b, 0x25,
	0x9e, 0xf3, 0x4b, 0x3f, 0xca, 0x98, 0x2f, 0x08, 0x9f, 0x13, 0x8e, 0x29, 0x49, 0xa6, 0x57, 0x51,
	0xce, 0x85, 0x1d, 0xf0, 0xb3, 0x9c, 0x26, 0x32, 0xe7, 0x24, 0xae, 0x61, 0xda, 0xb1, 0xdd, 0x5f,
	0x1c, 0xb4, 0x61, 0x95, 0x8d, 0x9f, 0x9f, 0xfc, 0x78, 0xee, 0x3e, 0x46, 0x3d, 0x53, 0x04, 0x0d,
	0x33, 0xe2, 0x39, 0x3b, 0xce, 0xde, 0xf2, 0x78, 0xb5, 0x0c, 0x9e, 0x86, 0x19, 0x71, 0xb7, 0xd1,
	0x72, 0xf1, 0xdc, 0x24, 0xf6, 0xee, 0xec, 0x38, 0x7b, 0xbd, 0xf1, 0x7d, 0x1d, 0xf8, 0x3e, 0x76,
	0x9f, 0xa0, 0x41, 0xf9, 0x18, 0x1c, 0xc6, 0x31, 0x27, 0x42, 0x78, 0x4b, 0x20, 0xd2, 0x2f, 0xe3,
	0xdf, 0xe8, 0xf0, 0xee, 0x01, 0x1a, 0xa8, 0x22, 0xd2, 0xfd, 0x39, 0xa3, 0x80, 0xc5, 0x52, 0x15,
	0x50, 0x7d, 0x57, 0x5c, 0x67, 0x67, 0x69, 0xaf, 0x37, 0x5e, 0x85, 0x60, 0x49, 0x3c, 0xd2, 0x4d,
	0xd7, 0xc0, 0x4c, 0xb0, 0xc5, 0x1b, 0xb1, 0x4f, 0x91, 0xa7, 0xd8, 0x30, 0x11, 0xc1, 0x59, 0x9e,
	0x26, 0x93, 0xeb, 0x52, 0x60, 0x1f, 0x6d, 0xb6, 0xe3, 0x95, 0xd0, 0xba, 0x4a, 0x9e, 0x17, 0xb9,
	0xb6, 0x5e, 0xf0, 0x0a, 0xbd, 0xe0, 0x26, 0xbd, 0xc0, 0xd6, 0xfb, 0xb9, 0x8f, 0x56, 0x95, 0xa0,
	0x82, 0xca, 0x6b, 0x46, 0xdc, 0x01, 0x5a, 0x0a, 0x2f, 0x93, 0x62, 0x14, 0xea, 0xa3, 0xfb, 0x11,
	0x5a, 0x85, 0x32, 0x4b, 0xb5, 0x3b, 0x90, 0x5a, 0x51, 0xb1, 0x42, 0xc5, 0x7d, 0x8a, 0x5c, 0x80,
	0x64, 0x93, 0x50, 0x48, 0x6b, 0x12, 0x03, 0x95, 0x19, 0xa9, 0x84, 0x8d, 0x4e, 0xc3, 0x88, 0xa4,
	0x06, 0xfd, 0x4e, 0x85, 0x7e, 0xa1, 0x12, 0x25, 0xda, 0x47, 0xd0, 0x08, 0x2c, 0x67, 0x94, 0xd6,
	0xe0, 0xef, 0x02, 0x7c, 0x4d, 0xa5, 0x7e, 0x80, 0x4c, 0x89, 0xff, 0x18, 0xf5, 0x55, 0x30, 0x8b,
	0xab, 0x42, 0xee, 0x02, 0xf6, 0x41, 0x11, 0xb6, 0x80, 0xf5, 0x09, 0xde, 0xab, 0x80, 0xd5, 0x0c,
	0xdd, 0x4f, 0xd0, 0x46, 0x11, 0x09, 0xb5, 0xc3, 0x02, 0x7d, 0x1f, 0xd0, 0xeb, 0x65, 0x6e, 0x54,
	0xa5, 0x8a, 0x9e, 0x05, 0x46, 0x78, 0xd9, 0xf4, 0x2c, 0x68, 0x76, 0x21, 0xb0, 0x7a, 0x86, 0x4c,
	0x17, 0x82, 0x8e, 0x9e, 0x05, 0x56, 0xcf, 0x56, 0x2a, 0x74, 0xa3, 0x67, 0xda, 0x5a, 0x50, 0xb7,
	0xb6, 0x6a, 0xac, 0x05, 0x35, 0x6b, 0xc5, 0xca, 0xc0, 0x6a, 0x37, 0xea, 0xe8, 0x19, 0x6f, 0x2a,
	0xd9, 0x28, 0xe5, 0x2f, 0x07, 0xad, 0xc1, 0x8e, 0xcf, 0x59, 0x2a, 0x0c, 0xe1, 0xc1, 0x8e, 0xb3,
	0xb7, 0xb2, 0xff, 0xbb, 0xe3, 0xdf, 0xde, 0x93, 0xc8, 0xb7, 0x0f, 0x80, 0xf1, 0xc0, 0xf8, 0xa8,
	0x35, 0x84, 0x4b, 0x3c, 0xc9, 0xa9, 0x90, 0x3c, 0x4c, 0x68, 0xd5, 0x90, 0xbe, 0x6e, 0x08, 0x97,
	0xdf, 0x9a, 0x5c, 0xc9, 0x79, 0x82, 0x60, 0x02, 0x59, 0xbd, 0xdd, 0x03, 0x7d, 0x0a, 0x95, 0xf1,
	0x26, 0x74, 0xd8, 0x80, 0xae, 0x19, 0xe8, 0xb0, 0x0e, 0x7d, 0x8a, 0x5c, 0x5d, 0x2b, 0xa9, 0x83,
	0x5d, 0x3d, 0x71, 0xc8, 0x9c, 0xd4, 0xd0, 0xbb, 0xa8, 0x97, 0x0a, 0x5c, 0x9b, 0xc7, 0xba, 0xde,
	0xb8, 0x54, 0xbc, 0x30, 0xde, 0xfe, 0x71, 0x4a, 0x49, 0x73, 0x8c, 0x29, 0xe4, 0x06, 0x4c, 0xee,
	0x8f, 0xb7, 0x64, 0x72, 0xb5, 0x13, 0xb8, 0x68, 0xc3, 0x48, 0xb0, 0x45, 0x73, 0x9f, 0x87, 0xf8,
	0x32, 0xcd, 0x17, 0x82, 0x91, 0x89, 0x31, 0xb9, 0x59, 0xbd, 0xab, 0xdf, 0x15, 0x39, 0xeb, 0x1d,
	0x68, 0x73, 0x1e, 0x56, 0xef, 0x80, 0xcd, 0x39, 0x44, 0x5e, 0x79, 0x76, 0xb4, 0x68, 0x8f, 0x80,
	0xf6, 0xb0, 0xc8, 0x77, 0x33, 0x83, 0x4e, 0xa6, 0x67, 0x98, 0x41, 0x07, 0xf3, 0x5f, 0x47, 0x9b,
	0x13, 0x1c, 0x5b, 0xe7, 0xfb, 0x16, 0x4c, 0xf0, 0xcf, 0x5b, 0x3f, 0xc1, 0xae, 0x5f, 0x41, 0x3d,
	0x13, 0xc1, 0xcf, 0xea, 0x3f, 0x4b, 0xa5, 0xd7, 0xa0, 0xed, 0xf5, 0xbd, 0xb7, 0xc7, 0x6b, 0xd0,
	0xe9, 0x35, 0xb0, 0xbc, 0xee, 0xfe, 0x77, 0x0f, 0xf5, 0xad, 0xfb, 0x91, 0xfb, 0xb7, 0xba, 0x33,
	0xa9, 0xef, 0xf6, 0xed, 0x66, 0x1f, 0xec, 0xff, 0x76, 0xeb, 0xed, 0x97, 0xf7, 0x89, 0xb1, 0x1b,
	0x65, 0xec, 0xb4, 0x79, 0x17, 0x73, 0x0f, 0x90, 0xd7, 0x30, 0x97, 0x89, 0x29, 0x66, 0x84, 0xc6,
	0x09, 0x9d, 0x7a, 0xcf, 0xe0, 0x8a, 0xb7, 0x59, 0x63, 0x8d, 0xc4, 0xf4, 0x4c, 0x27, 0x3b, 0x89,
	0x31, 0xcf, 0x19, 0x23, 0xb1, 0x37, 0xec, 0x22, 0x1e, 0xeb, 0xa4, 0x7b, 0x88, 0xb6, 0x1a, 0x44,
	0x46, 0x08, 0xc7, 0x33, 0x86, 0x05, 0xa1, 0xd2, 0xfb, 0xb4, 0xc5, 0x3c, 0x23, 0x84, 0x5f, 0xb0,
	0x73, 0x42, 0xa5, 0x7b, 0x84, 0xb6, 0xdb, 0xcc, 0x38, 0x5f, 0x50, 0xcd, 0x0d, 0x80, 0xfb, 0xc8,
	0xe2, 0x1e, 0xe7, 0x0b, 0x0a, 0xec, 0x63, 0xf4, 0x61, 0x83, 0xcd, 0xf3, 0x99, 0x24, 0xb8, 0xe8,
	0x99, 0x56, 0x38, 0x00, 0x85, 0xed, 0x9a, 0xc2, 0x58, 0x81, 0x46, 0x1a, 0x03, 0x2a, 0xcf, 0xd1,
	0xe3, 0x1b, 0x54, 0x48, 0xa9, 0x74, 0x08, 0x4a, 0x1f, 0xbc, 0x42, 0xe9, 0xa4, 0x10, 0xfb, 0x1a,
	0xbd, 0xdf, 0x34, 0x14, 0xca, 0x2b, 0x3c, 0x63, 0x71, 0x28, 0x89, 0x96, 0xf9, 0x0c, 0x64, 0xb6,
	0xea, 0x96, 0x42, 0x79, 0x75, 0x01, 0x88, 0x4e, 0x53, 0xa0, 0xb0, 0x48, 0xe4, 0x55, 0xcc, 0xc3,
	0x85, 0xd6, 0xf8, 0xbc, 0x65, 0x4a, 0x69, 0xbc, 0x2c, 0x30, 0xaf, 0xad, 0x43, 0xcd, 0xd4, 0xfb,
	0xe2, 0x86, 0x3a, 0xd4, 0x5c, 0x5f, 0x53, 0x07, 0x68, 0x1c, 0xdd, 0x58, 0x07, 0xa8, 0xd8, 0xab,
	0x31, 0x63, 0x31, 0xec, 0x15, 0xf8, 0xf8, 0xb2, 0xb5, 0x1a, 0x17, 0x2c, 0x1e, 0x89, 0x69, 0xe7,
	0x6a, 0x2c, 0xe0, 0xb9, 0x86, 0xfb, 0x55, 0x6b, 0x35, 0x5e, 0x2a, 0x40, 0xc1, 0x8e, 0xee, 0xc2,
	0xbf, 0x6e, 0xcf, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x63, 0xa4, 0xa0, 0xd3, 0x0d, 0x00,
	0x00,
}
