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
// source: l2fib_main_interface_detail_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_main_interfaces_l2fib_main_interface_l2fib_main_interface_detail_info

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

type L2FibMainInterfaceDetailInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	MainInterfaceId      uint32   `protobuf:"varint,2,opt,name=main_interface_id,json=mainInterfaceId,proto3" json:"main_interface_id,omitempty"`
	MainInterfaceType    string   `protobuf:"bytes,3,opt,name=main_interface_type,json=mainInterfaceType,proto3" json:"main_interface_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceDetailInfo_KEYS) Reset()         { *m = L2FibMainInterfaceDetailInfo_KEYS{} }
func (m *L2FibMainInterfaceDetailInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceDetailInfo_KEYS) ProtoMessage()    {}
func (*L2FibMainInterfaceDetailInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0294706f50afcd, []int{0}
}

func (m *L2FibMainInterfaceDetailInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceDetailInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceDetailInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS.Merge(m, src)
}
func (m *L2FibMainInterfaceDetailInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS.Size(m)
}
func (m *L2FibMainInterfaceDetailInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceDetailInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMainInterfaceDetailInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMainInterfaceDetailInfo_KEYS) GetMainInterfaceId() uint32 {
	if m != nil {
		return m.MainInterfaceId
	}
	return 0
}

func (m *L2FibMainInterfaceDetailInfo_KEYS) GetMainInterfaceType() string {
	if m != nil {
		return m.MainInterfaceType
	}
	return ""
}

type L2VpnEvpnRemoteShgInfo struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	Label                uint32   `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnRemoteShgInfo) Reset()         { *m = L2VpnEvpnRemoteShgInfo{} }
func (m *L2VpnEvpnRemoteShgInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnRemoteShgInfo) ProtoMessage()    {}
func (*L2VpnEvpnRemoteShgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0294706f50afcd, []int{1}
}

func (m *L2VpnEvpnRemoteShgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Merge(m, src)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnRemoteShgInfo.Size(m)
}
func (m *L2VpnEvpnRemoteShgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnRemoteShgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnRemoteShgInfo proto.InternalMessageInfo

func (m *L2VpnEvpnRemoteShgInfo) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnEvpnRemoteShgInfo) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type L2FibEvpnRemoteShgInfo struct {
	RemoteSplitHorizonGroupLabel *L2VpnEvpnRemoteShgInfo `protobuf:"bytes,1,opt,name=remote_split_horizon_group_label,json=remoteSplitHorizonGroupLabel,proto3" json:"remote_split_horizon_group_label,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                `json:"-"`
	XXX_unrecognized             []byte                  `json:"-"`
	XXX_sizecache                int32                   `json:"-"`
}

func (m *L2FibEvpnRemoteShgInfo) Reset()         { *m = L2FibEvpnRemoteShgInfo{} }
func (m *L2FibEvpnRemoteShgInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnRemoteShgInfo) ProtoMessage()    {}
func (*L2FibEvpnRemoteShgInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0294706f50afcd, []int{2}
}

func (m *L2FibEvpnRemoteShgInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnRemoteShgInfo.Unmarshal(m, b)
}
func (m *L2FibEvpnRemoteShgInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnRemoteShgInfo.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnRemoteShgInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnRemoteShgInfo.Merge(m, src)
}
func (m *L2FibEvpnRemoteShgInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnRemoteShgInfo.Size(m)
}
func (m *L2FibEvpnRemoteShgInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnRemoteShgInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnRemoteShgInfo proto.InternalMessageInfo

func (m *L2FibEvpnRemoteShgInfo) GetRemoteSplitHorizonGroupLabel() *L2VpnEvpnRemoteShgInfo {
	if m != nil {
		return m.RemoteSplitHorizonGroupLabel
	}
	return nil
}

type L2FibMainInterfaceEsNhInfo struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceEsNhInfo) Reset()         { *m = L2FibMainInterfaceEsNhInfo{} }
func (m *L2FibMainInterfaceEsNhInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceEsNhInfo) ProtoMessage()    {}
func (*L2FibMainInterfaceEsNhInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0294706f50afcd, []int{3}
}

func (m *L2FibMainInterfaceEsNhInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceEsNhInfo.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceEsNhInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceEsNhInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceEsNhInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceEsNhInfo.Merge(m, src)
}
func (m *L2FibMainInterfaceEsNhInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceEsNhInfo.Size(m)
}
func (m *L2FibMainInterfaceEsNhInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceEsNhInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceEsNhInfo proto.InternalMessageInfo

func (m *L2FibMainInterfaceEsNhInfo) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type L2FibMainInterfaceDetailInfo struct {
	MainInterface             string                        `protobuf:"bytes,50,opt,name=main_interface,json=mainInterface,proto3" json:"main_interface,omitempty"`
	MainIfType                uint32                        `protobuf:"varint,51,opt,name=main_if_type,json=mainIfType,proto3" json:"main_if_type,omitempty"`
	VirtualIfName             string                        `protobuf:"bytes,52,opt,name=virtual_if_name,json=virtualIfName,proto3" json:"virtual_if_name,omitempty"`
	InstanceId                []uint32                      `protobuf:"varint,53,rep,packed,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstancesState            []string                      `protobuf:"bytes,54,rep,name=instances_state,json=instancesState,proto3" json:"instances_state,omitempty"`
	InstancesProvisioned      []bool                        `protobuf:"varint,55,rep,packed,name=instances_provisioned,json=instancesProvisioned,proto3" json:"instances_provisioned,omitempty"`
	BridgePortCountInstance   []uint32                      `protobuf:"varint,56,rep,packed,name=bridge_port_count_instance,json=bridgePortCountInstance,proto3" json:"bridge_port_count_instance,omitempty"`
	Mac                       string                        `protobuf:"bytes,57,opt,name=mac,proto3" json:"mac,omitempty"`
	EsiId                     uint32                        `protobuf:"varint,58,opt,name=esi_id,json=esiId,proto3" json:"esi_id,omitempty"`
	EthernetSegmentIdentifier []uint32                      `protobuf:"varint,59,rep,packed,name=ethernet_segment_identifier,json=ethernetSegmentIdentifier,proto3" json:"ethernet_segment_identifier,omitempty"`
	EvpnShgLocalLabel         uint32                        `protobuf:"varint,60,opt,name=evpn_shg_local_label,json=evpnShgLocalLabel,proto3" json:"evpn_shg_local_label,omitempty"`
	EvpnShgRemoteInfo         []*L2FibEvpnRemoteShgInfo     `protobuf:"bytes,61,rep,name=evpn_shg_remote_info,json=evpnShgRemoteInfo,proto3" json:"evpn_shg_remote_info,omitempty"`
	EvpnRid                   string                        `protobuf:"bytes,62,opt,name=evpn_rid,json=evpnRid,proto3" json:"evpn_rid,omitempty"`
	EsNextHop                 []*L2FibMainInterfaceEsNhInfo `protobuf:"bytes,63,rep,name=es_next_hop,json=esNextHop,proto3" json:"es_next_hop,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                      `json:"-"`
	XXX_unrecognized          []byte                        `json:"-"`
	XXX_sizecache             int32                         `json:"-"`
}

func (m *L2FibMainInterfaceDetailInfo) Reset()         { *m = L2FibMainInterfaceDetailInfo{} }
func (m *L2FibMainInterfaceDetailInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceDetailInfo) ProtoMessage()    {}
func (*L2FibMainInterfaceDetailInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0294706f50afcd, []int{4}
}

func (m *L2FibMainInterfaceDetailInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceDetailInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceDetailInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceDetailInfo.Merge(m, src)
}
func (m *L2FibMainInterfaceDetailInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceDetailInfo.Size(m)
}
func (m *L2FibMainInterfaceDetailInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceDetailInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceDetailInfo proto.InternalMessageInfo

func (m *L2FibMainInterfaceDetailInfo) GetMainInterface() string {
	if m != nil {
		return m.MainInterface
	}
	return ""
}

func (m *L2FibMainInterfaceDetailInfo) GetMainIfType() uint32 {
	if m != nil {
		return m.MainIfType
	}
	return 0
}

func (m *L2FibMainInterfaceDetailInfo) GetVirtualIfName() string {
	if m != nil {
		return m.VirtualIfName
	}
	return ""
}

func (m *L2FibMainInterfaceDetailInfo) GetInstanceId() []uint32 {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetInstancesState() []string {
	if m != nil {
		return m.InstancesState
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetInstancesProvisioned() []bool {
	if m != nil {
		return m.InstancesProvisioned
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetBridgePortCountInstance() []uint32 {
	if m != nil {
		return m.BridgePortCountInstance
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *L2FibMainInterfaceDetailInfo) GetEsiId() uint32 {
	if m != nil {
		return m.EsiId
	}
	return 0
}

func (m *L2FibMainInterfaceDetailInfo) GetEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.EthernetSegmentIdentifier
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetEvpnShgLocalLabel() uint32 {
	if m != nil {
		return m.EvpnShgLocalLabel
	}
	return 0
}

func (m *L2FibMainInterfaceDetailInfo) GetEvpnShgRemoteInfo() []*L2FibEvpnRemoteShgInfo {
	if m != nil {
		return m.EvpnShgRemoteInfo
	}
	return nil
}

func (m *L2FibMainInterfaceDetailInfo) GetEvpnRid() string {
	if m != nil {
		return m.EvpnRid
	}
	return ""
}

func (m *L2FibMainInterfaceDetailInfo) GetEsNextHop() []*L2FibMainInterfaceEsNhInfo {
	if m != nil {
		return m.EsNextHop
	}
	return nil
}

func init() {
	proto.RegisterType((*L2FibMainInterfaceDetailInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_detail_info.l2fib_main_interface_detail_info_KEYS")
	proto.RegisterType((*L2VpnEvpnRemoteShgInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_detail_info.l2vpn_evpn_remote_shg_info")
	proto.RegisterType((*L2FibEvpnRemoteShgInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_detail_info.l2fib_evpn_remote_shg_info")
	proto.RegisterType((*L2FibMainInterfaceEsNhInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_detail_info.l2fib_main_interface_es_nh_info")
	proto.RegisterType((*L2FibMainInterfaceDetailInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_detail_info.l2fib_main_interface_detail_info")
}

func init() {
	proto.RegisterFile("l2fib_main_interface_detail_info.proto", fileDescriptor_cf0294706f50afcd)
}

var fileDescriptor_cf0294706f50afcd = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x4e, 0x14, 0x31,
	0x14, 0xce, 0xb0, 0x91, 0x85, 0x22, 0x20, 0x15, 0x42, 0x41, 0x13, 0x26, 0x9b, 0x80, 0x1b, 0x2f,
	0xc6, 0x64, 0xf1, 0x1f, 0xc5, 0x0b, 0x63, 0x64, 0x22, 0x12, 0x32, 0xeb, 0x8d, 0x57, 0x4d, 0x77,
	0x7b, 0x66, 0xb7, 0xc9, 0x6c, 0x3b, 0x69, 0x0b, 0x82, 0x37, 0x5e, 0xf8, 0x12, 0x26, 0x7a, 0xe1,
	0xab, 0xf8, 0x00, 0xbe, 0x85, 0x0f, 0x62, 0xda, 0xce, 0x2c, 0x42, 0x96, 0x70, 0x29, 0x37, 0x93,
	0x39, 0xe7, 0x7c, 0xf3, 0x9d, 0xbf, 0xaf, 0x1d, 0xb4, 0x55, 0x74, 0x72, 0xd1, 0xa3, 0x23, 0x26,
	0x24, 0x15, 0xd2, 0x82, 0xce, 0x59, 0x1f, 0x28, 0x07, 0xcb, 0x44, 0x41, 0x85, 0xcc, 0x55, 0x52,
	0x6a, 0x65, 0x15, 0xfe, 0xd2, 0x17, 0xa6, 0xaf, 0xa8, 0x50, 0x86, 0x9e, 0x68, 0x5a, 0x74, 0x8e,
	0x4b, 0x49, 0x55, 0x09, 0x3a, 0x09, 0xaf, 0xb9, 0xd2, 0x9f, 0x98, 0xe6, 0x42, 0x0e, 0x12, 0xa9,
	0x38, 0x18, 0xff, 0x4c, 0x26, 0x71, 0x9b, 0x89, 0xde, 0xe4, 0xaa, 0x32, 0x5a, 0x3f, 0x22, 0xb4,
	0x79, 0x15, 0x88, 0xbe, 0x7b, 0xf3, 0xb1, 0x8b, 0x57, 0x51, 0xd3, 0x65, 0xa7, 0x82, 0x93, 0x28,
	0x8e, 0xda, 0xb3, 0xd9, 0xb4, 0x33, 0x53, 0x8e, 0xef, 0xa3, 0xa5, 0x0b, 0xdf, 0x0a, 0x4e, 0xa6,
	0xe2, 0xa8, 0x3d, 0x9f, 0x2d, 0xba, 0x40, 0x5a, 0xfb, 0x53, 0x8e, 0x13, 0x74, 0xfb, 0x02, 0xd6,
	0x9e, 0x96, 0x40, 0x1a, 0x9e, 0x70, 0xe9, 0x1c, 0xfa, 0xc3, 0x69, 0x09, 0xad, 0xf7, 0x68, 0x3d,
	0x4c, 0x02, 0xdc, 0x43, 0xc3, 0x48, 0x59, 0xa0, 0x66, 0x38, 0xf0, 0x75, 0xe1, 0x35, 0x34, 0x23,
	0xe1, 0xc4, 0xd2, 0xa1, 0x2a, 0xab, 0x9a, 0x9a, 0xce, 0xde, 0x53, 0x25, 0x5e, 0x46, 0x37, 0x0a,
	0xd6, 0x83, 0xa2, 0x2a, 0x24, 0x18, 0xad, 0xaf, 0x53, 0x8e, 0xcf, 0x75, 0x3b, 0x91, 0xef, 0x4f,
	0x84, 0xe2, 0xda, 0x57, 0x16, 0xc2, 0x11, 0x6b, 0xf1, 0x59, 0x49, 0x3a, 0xd0, 0xea, 0xa8, 0xa4,
	0x81, 0xd0, 0x25, 0x9a, 0xeb, 0x7c, 0x8f, 0x92, 0xff, 0xbc, 0xba, 0xe4, 0xf2, 0xc1, 0x64, 0x77,
	0x83, 0xa3, 0xeb, 0x9a, 0xd8, 0x0b, 0x3d, 0xbc, 0x75, 0x2d, 0xec, 0xfb, 0x29, 0xec, 0xa0, 0x8d,
	0x89, 0xe4, 0x60, 0xa8, 0x1c, 0x86, 0x49, 0x10, 0xd4, 0x64, 0x9c, 0x6b, 0x30, 0xa6, 0x1e, 0x6c,
	0x65, 0xb6, 0xbe, 0x35, 0x51, 0x7c, 0x55, 0x69, 0x78, 0x13, 0x2d, 0x9c, 0x8f, 0x92, 0x8e, 0x67,
	0x99, 0x3f, 0xb7, 0x61, 0x1c, 0xa3, 0x9b, 0x01, 0x96, 0x07, 0x19, 0x6c, 0xfb, 0x5d, 0x21, 0x0f,
	0xca, 0xdd, 0xfe, 0xf1, 0x16, 0x5a, 0x3c, 0x16, 0xda, 0x1e, 0xb1, 0xc2, 0x81, 0x24, 0x1b, 0x01,
	0x79, 0x18, 0x98, 0x2a, 0x77, 0x9a, 0x1f, 0xb0, 0x11, 0xe0, 0x0d, 0x34, 0x27, 0xa4, 0xb1, 0x4c,
	0x06, 0xf5, 0x3d, 0x8a, 0x1b, 0x8e, 0xa8, 0x76, 0xa5, 0x1c, 0xdf, 0x43, 0x8b, 0xb5, 0x65, 0xa8,
	0xb1, 0xcc, 0x02, 0x79, 0x1c, 0x37, 0xda, 0xb3, 0xd9, 0xc2, 0xd8, 0xdd, 0x75, 0x5e, 0xbc, 0x8d,
	0x56, 0xce, 0x80, 0xa5, 0x56, 0xc7, 0xc2, 0x08, 0x25, 0x81, 0x93, 0x27, 0x71, 0xa3, 0x3d, 0x93,
	0x2d, 0x8f, 0x83, 0x87, 0x67, 0x31, 0xbc, 0x83, 0xd6, 0x7b, 0x5a, 0xf0, 0x01, 0xd0, 0x52, 0x69,
	0x4b, 0xfb, 0xea, 0x48, 0x5a, 0x5a, 0x23, 0xc9, 0x53, 0x5f, 0xcd, 0x6a, 0x40, 0x1c, 0x2a, 0x6d,
	0x5f, 0xbb, 0x78, 0x5a, 0x85, 0xf1, 0x2d, 0xd4, 0x18, 0xb1, 0x3e, 0x79, 0xe6, 0xfb, 0x72, 0xaf,
	0x78, 0x05, 0x4d, 0x83, 0x11, 0xae, 0x91, 0xe7, 0x41, 0xbd, 0x60, 0x44, 0xca, 0xf1, 0x2e, 0xba,
	0x03, 0x76, 0x08, 0x5a, 0x82, 0xa5, 0x06, 0x06, 0x23, 0x70, 0x49, 0x38, 0x48, 0x2b, 0x72, 0x01,
	0x9a, 0xec, 0xf8, 0x34, 0x6b, 0x35, 0xa4, 0x1b, 0x10, 0xe9, 0x18, 0x80, 0x1f, 0xa0, 0x65, 0xaf,
	0x16, 0x27, 0x93, 0x42, 0xf5, 0x59, 0x51, 0x29, 0xfa, 0x85, 0x4f, 0xb2, 0xe4, 0x62, 0xdd, 0xe1,
	0x60, 0xdf, 0x45, 0xbc, 0x50, 0xf0, 0xef, 0xe8, 0x9f, 0x2f, 0x2a, 0x8d, 0xb9, 0xfd, 0x92, 0x97,
	0x71, 0xe3, 0xba, 0x9c, 0x81, 0xcb, 0x0e, 0xf3, 0xb8, 0x9f, 0xcc, 0xfb, 0xd3, 0xea, 0xbe, 0x08,
	0x50, 0xc1, 0xc9, 0x6e, 0x90, 0xb5, 0xb3, 0x33, 0xc1, 0xf1, 0xaf, 0x08, 0xcd, 0x39, 0xfd, 0xd7,
	0xd7, 0xc9, 0x2b, 0xdf, 0xe1, 0xcf, 0xeb, 0xd2, 0xe1, 0xa5, 0x27, 0x35, 0x9b, 0x05, 0x73, 0x10,
	0xee, 0xbc, 0xde, 0xb4, 0xff, 0xa7, 0x6c, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x56, 0x1e,
	0xf6, 0x7d, 0x06, 0x00, 0x00,
}