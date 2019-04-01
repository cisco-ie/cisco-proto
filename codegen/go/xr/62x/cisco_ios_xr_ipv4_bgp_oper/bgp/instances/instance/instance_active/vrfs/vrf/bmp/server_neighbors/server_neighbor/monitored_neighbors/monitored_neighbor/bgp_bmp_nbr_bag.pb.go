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

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_bmp_server_neighbors_server_neighbor_monitored_neighbors_monitored_neighbor

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	ServerId             uint32   `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
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

func (m *BgpBmpNbrBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*BgpBmpNbrBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_bmp_nbr_bag_KEYS")
	proto.RegisterType((*BgpL2VpnAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_l2vpn_addr_t")
	proto.RegisterType((*BgpL2VpnMspwAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_l2vpn_mspw_addr_t")
	proto.RegisterType((*BgpIpv4SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_ipv4_srpolicy_addr_t")
	proto.RegisterType((*BgpIpv6SrpolicyAddrT)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_ipv6_srpolicy_addr_t")
	proto.RegisterType((*BgpAddrtype)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_addrtype")
	proto.RegisterType((*BgpBmpNbrBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.bmp.server_neighbors.server_neighbor.monitored_neighbors.monitored_neighbor.bgp_bmp_nbr_bag")
}

func init() { proto.RegisterFile("bgp_bmp_nbr_bag.proto", fileDescriptor_3893b48767fb191b) }

var fileDescriptor_3893b48767fb191b = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0x29, 0x6c, 0xdb, 0x69, 0xd2, 0xa6, 0x6e, 0xbb, 0xeb, 0x52, 0x01, 0xa5, 0x7b, 0xa0,
	0x2b, 0xad, 0x2c, 0xd1, 0x0d, 0x6e, 0x81, 0x82, 0x40, 0xb4, 0x48, 0x68, 0x37, 0x55, 0x95, 0x52,
	0xad, 0x38, 0x8d, 0xec, 0x78, 0x92, 0x5a, 0x8a, 0xc7, 0xa3, 0x99, 0x49, 0xc2, 0xfe, 0x11, 0x8e,
	0x9c, 0x38, 0x70, 0x41, 0xe2, 0xd7, 0xc0, 0x85, 0x2b, 0xf7, 0xfd, 0x09, 0x68, 0xde, 0xd8, 0x63,
	0x7b, 0xec, 0xed, 0x8a, 0x5b, 0x7b, 0xa9, 0x9c, 0xf7, 0xbe, 0xef, 0xeb, 0xfb, 0xde, 0x7b, 0x19,
	0x4f, 0xd0, 0x4e, 0x34, 0x61, 0x38, 0x4a, 0x19, 0xa6, 0x11, 0xc7, 0x51, 0x38, 0xf1, 0x19, 0xcf,
	0x64, 0xe6, 0xfe, 0xe2, 0x8c, 0x12, 0x31, 0xca, 0x70, 0x92, 0x09, 0xfc, 0x33, 0xc7, 0x09, 0x9b,
	0xf7, 0xb1, 0x42, 0x66, 0x8c, 0x70, 0x3f, 0x9a, 0x30, 0x3f, 0xa1, 0x42, 0x86, 0x74, 0x44, 0x84,
	0x79, 0x32, 0x0f, 0x38, 0x1c, 0xc9, 0x64, 0x4e, 0xfc, 0x39, 0x1f, 0x0b, 0xf5, 0xc7, 0x8f, 0x52,
	0xe6, 0x0b, 0xc2, 0xe7, 0x84, 0x63, 0x4a, 0x92, 0xc9, 0x4d, 0x94, 0x71, 0x61, 0x07, 0xfc, 0x34,
	0xa3, 0x89, 0xcc, 0x38, 0x89, 0x2b, 0x98, 0x66, 0xec, 0xe0, 0x57, 0x07, 0x6d, 0x5b, 0x25, 0xe3,
	0xe7, 0xe7, 0x3f, 0x5d, 0xb9, 0x8f, 0x51, 0xd7, 0x14, 0x40, 0xc3, 0x94, 0x78, 0xce, 0xbe, 0x73,
	0xb8, 0x3a, 0xec, 0x14, 0xc1, 0x8b, 0x30, 0x25, 0xee, 0x2e, 0x5a, 0x99, 0xf3, 0xb1, 0xce, 0xbf,
	0x03, 0xf9, 0xe5, 0x39, 0x1f, 0x43, 0x6a, 0x0f, 0xad, 0xe6, 0x25, 0x25, 0xb1, 0xb7, 0xb4, 0xef,
	0x1c, 0x76, 0x87, 0x2b, 0x3a, 0xf0, 0x43, 0xec, 0x3e, 0x41, 0xbd, 0xa2, 0x02, 0x1c, 0xc6, 0x31,
	0x27, 0x42, 0x78, 0xef, 0x02, 0x7f, 0xa3, 0x88, 0x7f, 0xab, 0xc3, 0x07, 0xc7, 0xa8, 0xa7, 0xea,
	0x9b, 0x1e, 0xcd, 0x19, 0x05, 0x2c, 0x96, 0xaa, 0xb6, 0xf2, 0xb3, 0xe2, 0x3a, 0xfb, 0x4b, 0x87,
	0xdd, 0x61, 0x07, 0x82, 0x05, 0xf1, 0x54, 0xcf, 0x42, 0x03, 0x53, 0xc1, 0x16, 0xff, 0x8b, 0x7d,
	0x81, 0x3c, 0xc5, 0x86, 0x41, 0x09, 0xce, 0xb2, 0x69, 0x32, 0x7a, 0x55, 0x08, 0x1c, 0xa1, 0x9d,
	0x66, 0xbc, 0x14, 0xda, 0x52, 0xc9, 0xab, 0x3c, 0xd7, 0xd4, 0x0b, 0xde, 0xa0, 0x17, 0xdc, 0xa6,
	0x17, 0xd8, 0x7a, 0xaf, 0xd7, 0x51, 0x47, 0x09, 0x2a, 0xa8, 0x7c, 0xc5, 0x88, 0xdb, 0x43, 0x4b,
	0xe1, 0x38, 0xc9, 0xa7, 0xa4, 0x1e, 0xdd, 0x8f, 0x51, 0x07, 0xca, 0x2c, 0xd4, 0xf4, 0x80, 0xd6,
	0x54, 0x2c, 0x57, 0x71, 0x9f, 0x22, 0x17, 0x20, 0xe9, 0x28, 0x14, 0xd2, 0x00, 0x97, 0x00, 0xd8,
	0x53, 0x99, 0x81, 0x4a, 0xd8, 0xe8, 0x69, 0x18, 0x91, 0xa9, 0x35, 0x37, 0x40, 0xbf, 0x50, 0x89,
	0x02, 0xed, 0x23, 0x68, 0x04, 0x96, 0x33, 0x4a, 0x2b, 0xf0, 0xf7, 0x00, 0xbe, 0xa9, 0x52, 0x3f,
	0x42, 0xa6, 0xc0, 0x7f, 0x82, 0x36, 0x54, 0x30, 0x8d, 0xcb, 0x42, 0x1e, 0x00, 0x76, 0x3d, 0x0f,
	0x5b, 0xc0, 0xea, 0x04, 0x97, 0x4b, 0x60, 0x39, 0x43, 0xf7, 0x53, 0xb4, 0x9d, 0x47, 0x42, 0xed,
	0x30, 0x47, 0xaf, 0x00, 0x7a, 0xab, 0xc8, 0x0d, 0xca, 0x54, 0xde, 0xb3, 0xc0, 0x08, 0xaf, 0x9a,
	0x9e, 0x05, 0xf5, 0x2e, 0x04, 0x56, 0xcf, 0x90, 0xe9, 0x42, 0xd0, 0xd2, 0xb3, 0xc0, 0xea, 0xd9,
	0x5a, 0x89, 0xae, 0xf5, 0x4c, 0x5b, 0x0b, 0xaa, 0xd6, 0x3a, 0xc6, 0x5a, 0x50, 0xb1, 0x96, 0xaf,
	0x0c, 0xac, 0x76, 0xad, 0x8e, 0xae, 0xf1, 0xa6, 0x92, 0xb5, 0x52, 0xfe, 0x72, 0xd0, 0x26, 0xec,
	0xf8, 0x9c, 0x4d, 0x85, 0x21, 0xac, 0xef, 0x3b, 0x87, 0x6b, 0x47, 0xbf, 0x3b, 0xfe, 0xdd, 0x3c,
	0xa0, 0x7c, 0xfb, 0xcb, 0x3f, 0xec, 0x19, 0x0f, 0x95, 0x66, 0x70, 0x89, 0x47, 0x19, 0x15, 0x92,
	0x87, 0x09, 0x2d, 0x9b, 0xb1, 0xa1, 0x9b, 0xc1, 0xe5, 0x77, 0x26, 0x57, 0x70, 0x9e, 0x20, 0xe8,
	0x7e, 0x5a, 0x6d, 0x75, 0x4f, 0x9f, 0x40, 0x45, 0xbc, 0x0e, 0xed, 0xd7, 0xa0, 0x9b, 0x06, 0xda,
	0xaf, 0x42, 0x9f, 0x22, 0x57, 0xd7, 0x4a, 0xaa, 0x60, 0x57, 0x4f, 0x1b, 0x32, 0xe7, 0x15, 0xf4,
	0x01, 0xea, 0x4e, 0x05, 0xae, 0xcc, 0x62, 0x4b, 0x6f, 0xdb, 0x54, 0xbc, 0x30, 0xde, 0xfe, 0x71,
	0x0a, 0x49, 0x73, 0x84, 0x29, 0xe4, 0x36, 0x4c, 0xed, 0x8f, 0x7b, 0x30, 0xb5, 0xca, 0xc9, 0x9b,
	0xb7, 0x60, 0x20, 0xd8, 0xa2, 0xbe, 0xc7, 0x7d, 0x3c, 0x9e, 0x66, 0x0b, 0xc1, 0xc8, 0xc8, 0x18,
	0xdc, 0x29, 0xbf, 0xa3, 0xdf, 0xe7, 0x39, 0x6b, 0xf7, 0x9b, 0x9c, 0x87, 0xe5, 0xee, 0xdb, 0x9c,
	0x13, 0xe4, 0x15, 0x67, 0x46, 0x83, 0xf6, 0x08, 0x68, 0x0f, 0xf3, 0x7c, 0x3b, 0x33, 0x68, 0x65,
	0x7a, 0x86, 0x19, 0xb4, 0x30, 0xff, 0x75, 0xb4, 0x39, 0xc1, 0xb1, 0x75, 0xae, 0xef, 0xc2, 0xf4,
	0xfe, 0xbc, 0xd3, 0xd3, 0x6b, 0x7b, 0xf3, 0xe9, 0x79, 0x08, 0x7e, 0x59, 0x7d, 0x15, 0x15, 0x3e,
	0x83, 0xa6, 0xcf, 0xf7, 0xef, 0x87, 0xcf, 0xa0, 0xd5, 0x67, 0x60, 0xf9, 0x3c, 0x78, 0xbd, 0x8c,
	0x36, 0xac, 0xab, 0x92, 0xfb, 0xb7, 0xba, 0x3e, 0xa9, 0xcf, 0xf6, 0x6d, 0xe6, 0x08, 0xac, 0xff,
	0x76, 0xa7, 0xad, 0x17, 0x77, 0x87, 0xa1, 0x1b, 0xa5, 0xec, 0xa2, 0x7e, 0xef, 0x72, 0x8f, 0x91,
	0x57, 0x33, 0x96, 0x8a, 0x09, 0x66, 0x84, 0xc6, 0x09, 0x9d, 0x78, 0xcf, 0xe0, 0x3a, 0xb7, 0x53,
	0x61, 0x0d, 0xc4, 0xe4, 0x52, 0x27, 0x5b, 0x89, 0x31, 0xcf, 0x18, 0x23, 0xb1, 0xd7, 0x6f, 0x23,
	0x9e, 0xe9, 0xa4, 0x7b, 0x82, 0x76, 0x6b, 0x44, 0x46, 0x08, 0xc7, 0x33, 0x86, 0x05, 0xa1, 0xd2,
	0xfb, 0xac, 0xc1, 0xbc, 0x24, 0x84, 0x5f, 0xb3, 0x2b, 0x42, 0xa5, 0x7b, 0x8a, 0xf6, 0x9a, 0xcc,
	0x38, 0x5b, 0x50, 0xcd, 0x0d, 0x80, 0xfb, 0xc8, 0xe2, 0x9e, 0x65, 0x0b, 0x0a, 0xec, 0x33, 0xf4,
	0x51, 0x8d, 0xcd, 0xb3, 0x99, 0x24, 0x38, 0xef, 0x99, 0x56, 0x38, 0x06, 0x85, 0xbd, 0x8a, 0xc2,
	0x50, 0x81, 0x06, 0x1a, 0x03, 0x2a, 0xcf, 0xd1, 0xe3, 0x5b, 0x54, 0x48, 0xa1, 0x74, 0x02, 0x4a,
	0x1f, 0xbe, 0x41, 0xe9, 0x3c, 0x17, 0xfb, 0x06, 0x7d, 0x50, 0x37, 0x14, 0xca, 0x1b, 0x3c, 0x63,
	0x71, 0x28, 0x89, 0x96, 0xf9, 0x1c, 0x64, 0x76, 0xab, 0x96, 0x42, 0x79, 0x73, 0x0d, 0x88, 0x56,
	0x53, 0xa0, 0xb0, 0x48, 0xe4, 0x4d, 0xcc, 0xc3, 0x85, 0xd6, 0xf8, 0xa2, 0x61, 0x4a, 0x69, 0xbc,
	0xcc, 0x31, 0x6f, 0xad, 0x43, 0xcd, 0xd4, 0xfb, 0xf2, 0x96, 0x3a, 0xd4, 0x5c, 0xdf, 0x52, 0x07,
	0x68, 0x9c, 0xde, 0x5a, 0x07, 0xa8, 0xd8, 0xab, 0x31, 0x63, 0x31, 0xec, 0x15, 0xf8, 0xf8, 0xaa,
	0xb1, 0x1a, 0xd7, 0x2c, 0x1e, 0x88, 0x49, 0xeb, 0x6a, 0x2c, 0xe0, 0xff, 0x1a, 0xee, 0xd7, 0x8d,
	0xd5, 0x78, 0xa9, 0x00, 0x39, 0x3b, 0x7a, 0x00, 0xbf, 0xde, 0x9e, 0xfd, 0x17, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0xb4, 0x2f, 0xfc, 0xd6, 0x0d, 0x00, 0x00,
}