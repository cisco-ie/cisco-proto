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
// source: l2fib_bridge_mac_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fibmac_details_l2fibmac_detail

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

type L2FibBridgeMacInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	GroupName            string   `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBridgeMacInfo_KEYS) Reset()         { *m = L2FibBridgeMacInfo_KEYS{} }
func (m *L2FibBridgeMacInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgeMacInfo_KEYS) ProtoMessage()    {}
func (*L2FibBridgeMacInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{0}
}

func (m *L2FibBridgeMacInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgeMacInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibBridgeMacInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgeMacInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibBridgeMacInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgeMacInfo_KEYS.Merge(m, src)
}
func (m *L2FibBridgeMacInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgeMacInfo_KEYS.Size(m)
}
func (m *L2FibBridgeMacInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgeMacInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgeMacInfo_KEYS proto.InternalMessageInfo

func (m *L2FibBridgeMacInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibBridgeMacInfo_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *L2FibBridgeMacInfo_KEYS) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *L2FibBridgeMacInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type L2FibBaseInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBaseInfo) Reset()         { *m = L2FibBaseInfo{} }
func (m *L2FibBaseInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBaseInfo) ProtoMessage()    {}
func (*L2FibBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{1}
}

func (m *L2FibBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBaseInfo.Unmarshal(m, b)
}
func (m *L2FibBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBaseInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBaseInfo.Merge(m, src)
}
func (m *L2FibBaseInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBaseInfo.Size(m)
}
func (m *L2FibBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBaseInfo proto.InternalMessageInfo

type L2FibAcKeyInfo struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	VspVlan              uint32   `protobuf:"varint,2,opt,name=vsp_vlan,json=vspVlan,proto3" json:"vsp_vlan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibAcKeyInfo) Reset()         { *m = L2FibAcKeyInfo{} }
func (m *L2FibAcKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibAcKeyInfo) ProtoMessage()    {}
func (*L2FibAcKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{2}
}

func (m *L2FibAcKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibAcKeyInfo.Unmarshal(m, b)
}
func (m *L2FibAcKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibAcKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibAcKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibAcKeyInfo.Merge(m, src)
}
func (m *L2FibAcKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibAcKeyInfo.Size(m)
}
func (m *L2FibAcKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibAcKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibAcKeyInfo proto.InternalMessageInfo

func (m *L2FibAcKeyInfo) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2FibAcKeyInfo) GetVspVlan() uint32 {
	if m != nil {
		return m.VspVlan
	}
	return 0
}

type L2FibPbbKeyInfo struct {
	Xcid                 uint32   `protobuf:"varint,1,opt,name=xcid,proto3" json:"xcid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPbbKeyInfo) Reset()         { *m = L2FibPbbKeyInfo{} }
func (m *L2FibPbbKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPbbKeyInfo) ProtoMessage()    {}
func (*L2FibPbbKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{3}
}

func (m *L2FibPbbKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPbbKeyInfo.Unmarshal(m, b)
}
func (m *L2FibPbbKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPbbKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPbbKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPbbKeyInfo.Merge(m, src)
}
func (m *L2FibPbbKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPbbKeyInfo.Size(m)
}
func (m *L2FibPbbKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPbbKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPbbKeyInfo proto.InternalMessageInfo

func (m *L2FibPbbKeyInfo) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

type L2FibVniKeyInfo struct {
	Xcid                 uint32   `protobuf:"varint,1,opt,name=xcid,proto3" json:"xcid,omitempty"`
	ParentIf             string   `protobuf:"bytes,2,opt,name=parent_if,json=parentIf,proto3" json:"parent_if,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibVniKeyInfo) Reset()         { *m = L2FibVniKeyInfo{} }
func (m *L2FibVniKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibVniKeyInfo) ProtoMessage()    {}
func (*L2FibVniKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{4}
}

func (m *L2FibVniKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibVniKeyInfo.Unmarshal(m, b)
}
func (m *L2FibVniKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibVniKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibVniKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibVniKeyInfo.Merge(m, src)
}
func (m *L2FibVniKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibVniKeyInfo.Size(m)
}
func (m *L2FibVniKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibVniKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibVniKeyInfo proto.InternalMessageInfo

func (m *L2FibVniKeyInfo) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

func (m *L2FibVniKeyInfo) GetParentIf() string {
	if m != nil {
		return m.ParentIf
	}
	return ""
}

type L2FibEvpnKeyInfo struct {
	Xcid                 uint32   `protobuf:"varint,1,opt,name=xcid,proto3" json:"xcid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibEvpnKeyInfo) Reset()         { *m = L2FibEvpnKeyInfo{} }
func (m *L2FibEvpnKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibEvpnKeyInfo) ProtoMessage()    {}
func (*L2FibEvpnKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{5}
}

func (m *L2FibEvpnKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibEvpnKeyInfo.Unmarshal(m, b)
}
func (m *L2FibEvpnKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibEvpnKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibEvpnKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibEvpnKeyInfo.Merge(m, src)
}
func (m *L2FibEvpnKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibEvpnKeyInfo.Size(m)
}
func (m *L2FibEvpnKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibEvpnKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibEvpnKeyInfo proto.InternalMessageInfo

func (m *L2FibEvpnKeyInfo) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

type L2FibPwKeyInfo struct {
	PwId                 uint64   `protobuf:"varint,1,opt,name=pw_id,json=pwId,proto3" json:"pw_id,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	PseudoWireIdType     string   `protobuf:"bytes,3,opt,name=pseudo_wire_id_type,json=pseudoWireIdType,proto3" json:"pseudo_wire_id_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPwKeyInfo) Reset()         { *m = L2FibPwKeyInfo{} }
func (m *L2FibPwKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPwKeyInfo) ProtoMessage()    {}
func (*L2FibPwKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{6}
}

func (m *L2FibPwKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwKeyInfo.Unmarshal(m, b)
}
func (m *L2FibPwKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPwKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwKeyInfo.Merge(m, src)
}
func (m *L2FibPwKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPwKeyInfo.Size(m)
}
func (m *L2FibPwKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwKeyInfo proto.InternalMessageInfo

func (m *L2FibPwKeyInfo) GetPwId() uint64 {
	if m != nil {
		return m.PwId
	}
	return 0
}

func (m *L2FibPwKeyInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *L2FibPwKeyInfo) GetPseudoWireIdType() string {
	if m != nil {
		return m.PseudoWireIdType
	}
	return ""
}

type L2FibBridgePortSegKeyUn struct {
	DataType             string            `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Ac                   *L2FibAcKeyInfo   `protobuf:"bytes,2,opt,name=ac,proto3" json:"ac,omitempty"`
	Pbb                  *L2FibPbbKeyInfo  `protobuf:"bytes,3,opt,name=pbb,proto3" json:"pbb,omitempty"`
	Vni                  *L2FibVniKeyInfo  `protobuf:"bytes,4,opt,name=vni,proto3" json:"vni,omitempty"`
	Evpn                 *L2FibEvpnKeyInfo `protobuf:"bytes,5,opt,name=evpn,proto3" json:"evpn,omitempty"`
	Pw                   *L2FibPwKeyInfo   `protobuf:"bytes,6,opt,name=pw,proto3" json:"pw,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2FibBridgePortSegKeyUn) Reset()         { *m = L2FibBridgePortSegKeyUn{} }
func (m *L2FibBridgePortSegKeyUn) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgePortSegKeyUn) ProtoMessage()    {}
func (*L2FibBridgePortSegKeyUn) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{7}
}

func (m *L2FibBridgePortSegKeyUn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgePortSegKeyUn.Unmarshal(m, b)
}
func (m *L2FibBridgePortSegKeyUn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgePortSegKeyUn.Marshal(b, m, deterministic)
}
func (m *L2FibBridgePortSegKeyUn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgePortSegKeyUn.Merge(m, src)
}
func (m *L2FibBridgePortSegKeyUn) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgePortSegKeyUn.Size(m)
}
func (m *L2FibBridgePortSegKeyUn) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgePortSegKeyUn.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgePortSegKeyUn proto.InternalMessageInfo

func (m *L2FibBridgePortSegKeyUn) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *L2FibBridgePortSegKeyUn) GetAc() *L2FibAcKeyInfo {
	if m != nil {
		return m.Ac
	}
	return nil
}

func (m *L2FibBridgePortSegKeyUn) GetPbb() *L2FibPbbKeyInfo {
	if m != nil {
		return m.Pbb
	}
	return nil
}

func (m *L2FibBridgePortSegKeyUn) GetVni() *L2FibVniKeyInfo {
	if m != nil {
		return m.Vni
	}
	return nil
}

func (m *L2FibBridgePortSegKeyUn) GetEvpn() *L2FibEvpnKeyInfo {
	if m != nil {
		return m.Evpn
	}
	return nil
}

func (m *L2FibBridgePortSegKeyUn) GetPw() *L2FibPwKeyInfo {
	if m != nil {
		return m.Pw
	}
	return nil
}

type L2VpnEvpnMoiInfo struct {
	TunnelEndpointId     uint32   `protobuf:"varint,1,opt,name=tunnel_endpoint_id,json=tunnelEndpointId,proto3" json:"tunnel_endpoint_id,omitempty"`
	NextHopIpv6Addr      string   `protobuf:"bytes,2,opt,name=next_hop_ipv6_addr,json=nextHopIpv6Addr,proto3" json:"next_hop_ipv6_addr,omitempty"`
	McastLabel           uint32   `protobuf:"varint,3,opt,name=mcast_label,json=mcastLabel,proto3" json:"mcast_label,omitempty"`
	McastEncapsulation   uint32   `protobuf:"varint,4,opt,name=mcast_encapsulation,json=mcastEncapsulation,proto3" json:"mcast_encapsulation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnMoiInfo) Reset()         { *m = L2VpnEvpnMoiInfo{} }
func (m *L2VpnEvpnMoiInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnMoiInfo) ProtoMessage()    {}
func (*L2VpnEvpnMoiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{8}
}

func (m *L2VpnEvpnMoiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnMoiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnMoiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnMoiInfo.Merge(m, src)
}
func (m *L2VpnEvpnMoiInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnMoiInfo.Size(m)
}
func (m *L2VpnEvpnMoiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnMoiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnMoiInfo proto.InternalMessageInfo

func (m *L2VpnEvpnMoiInfo) GetTunnelEndpointId() uint32 {
	if m != nil {
		return m.TunnelEndpointId
	}
	return 0
}

func (m *L2VpnEvpnMoiInfo) GetNextHopIpv6Addr() string {
	if m != nil {
		return m.NextHopIpv6Addr
	}
	return ""
}

func (m *L2VpnEvpnMoiInfo) GetMcastLabel() uint32 {
	if m != nil {
		return m.McastLabel
	}
	return 0
}

func (m *L2VpnEvpnMoiInfo) GetMcastEncapsulation() uint32 {
	if m != nil {
		return m.McastEncapsulation
	}
	return 0
}

type L2FibBridgeMacEvpnCtxUn struct {
	DataType             string            `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	EsiId                uint32            `protobuf:"varint,2,opt,name=esi_id,json=esiId,proto3" json:"esi_id,omitempty"`
	LocalLabel           uint32            `protobuf:"varint,3,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	McastOle             *L2VpnEvpnMoiInfo `protobuf:"bytes,4,opt,name=mcast_ole,json=mcastOle,proto3" json:"mcast_ole,omitempty"`
	BpIfh                string            `protobuf:"bytes,5,opt,name=bp_ifh,json=bpIfh,proto3" json:"bp_ifh,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2FibBridgeMacEvpnCtxUn) Reset()         { *m = L2FibBridgeMacEvpnCtxUn{} }
func (m *L2FibBridgeMacEvpnCtxUn) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgeMacEvpnCtxUn) ProtoMessage()    {}
func (*L2FibBridgeMacEvpnCtxUn) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{9}
}

func (m *L2FibBridgeMacEvpnCtxUn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgeMacEvpnCtxUn.Unmarshal(m, b)
}
func (m *L2FibBridgeMacEvpnCtxUn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgeMacEvpnCtxUn.Marshal(b, m, deterministic)
}
func (m *L2FibBridgeMacEvpnCtxUn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgeMacEvpnCtxUn.Merge(m, src)
}
func (m *L2FibBridgeMacEvpnCtxUn) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgeMacEvpnCtxUn.Size(m)
}
func (m *L2FibBridgeMacEvpnCtxUn) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgeMacEvpnCtxUn.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgeMacEvpnCtxUn proto.InternalMessageInfo

func (m *L2FibBridgeMacEvpnCtxUn) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *L2FibBridgeMacEvpnCtxUn) GetEsiId() uint32 {
	if m != nil {
		return m.EsiId
	}
	return 0
}

func (m *L2FibBridgeMacEvpnCtxUn) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *L2FibBridgeMacEvpnCtxUn) GetMcastOle() *L2VpnEvpnMoiInfo {
	if m != nil {
		return m.McastOle
	}
	return nil
}

func (m *L2FibBridgeMacEvpnCtxUn) GetBpIfh() string {
	if m != nil {
		return m.BpIfh
	}
	return ""
}

type L2FibNhopInfo struct {
	Base                  *L2FibBaseInfo `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	NextHopAddress        string         `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	NextHopAddressV6      string         `protobuf:"bytes,3,opt,name=next_hop_address_v6,json=nextHopAddressV6,proto3" json:"next_hop_address_v6,omitempty"`
	NextHopInternalLabel  uint32         `protobuf:"varint,4,opt,name=next_hop_internal_label,json=nextHopInternalLabel,proto3" json:"next_hop_internal_label,omitempty"`
	EcdPlaformtDataValid  bool           `protobuf:"varint,5,opt,name=ecd_plaformt_data_valid,json=ecdPlaformtDataValid,proto3" json:"ecd_plaformt_data_valid,omitempty"`
	EcdPlatformDataLength uint32         `protobuf:"varint,6,opt,name=ecd_platform_data_length,json=ecdPlatformDataLength,proto3" json:"ecd_platform_data_length,omitempty"`
	ChildrenCount         uint32         `protobuf:"varint,7,opt,name=children_count,json=childrenCount,proto3" json:"children_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}       `json:"-"`
	XXX_unrecognized      []byte         `json:"-"`
	XXX_sizecache         int32          `json:"-"`
}

func (m *L2FibNhopInfo) Reset()         { *m = L2FibNhopInfo{} }
func (m *L2FibNhopInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibNhopInfo) ProtoMessage()    {}
func (*L2FibNhopInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{10}
}

func (m *L2FibNhopInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibNhopInfo.Unmarshal(m, b)
}
func (m *L2FibNhopInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibNhopInfo.Marshal(b, m, deterministic)
}
func (m *L2FibNhopInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibNhopInfo.Merge(m, src)
}
func (m *L2FibNhopInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibNhopInfo.Size(m)
}
func (m *L2FibNhopInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibNhopInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibNhopInfo proto.InternalMessageInfo

func (m *L2FibNhopInfo) GetBase() *L2FibBaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *L2FibNhopInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *L2FibNhopInfo) GetNextHopAddressV6() string {
	if m != nil {
		return m.NextHopAddressV6
	}
	return ""
}

func (m *L2FibNhopInfo) GetNextHopInternalLabel() uint32 {
	if m != nil {
		return m.NextHopInternalLabel
	}
	return 0
}

func (m *L2FibNhopInfo) GetEcdPlaformtDataValid() bool {
	if m != nil {
		return m.EcdPlaformtDataValid
	}
	return false
}

func (m *L2FibNhopInfo) GetEcdPlatformDataLength() uint32 {
	if m != nil {
		return m.EcdPlatformDataLength
	}
	return 0
}

func (m *L2FibNhopInfo) GetChildrenCount() uint32 {
	if m != nil {
		return m.ChildrenCount
	}
	return 0
}

type L2FibBridgeMacInfo struct {
	Base                  *L2FibBaseInfo           `protobuf:"bytes,50,opt,name=base,proto3" json:"base,omitempty"`
	PlatformIsHwLearn     uint32                   `protobuf:"varint,51,opt,name=platform_is_hw_learn,json=platformIsHwLearn,proto3" json:"platform_is_hw_learn,omitempty"`
	Segment               *L2FibBridgePortSegKeyUn `protobuf:"bytes,52,opt,name=segment,proto3" json:"segment,omitempty"`
	NodeId                string                   `protobuf:"bytes,53,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	BridgeDomainName      string                   `protobuf:"bytes,54,opt,name=bridge_domain_name,json=bridgeDomainName,proto3" json:"bridge_domain_name,omitempty"`
	BridgeId              uint32                   `protobuf:"varint,55,opt,name=bridge_id,json=bridgeId,proto3" json:"bridge_id,omitempty"`
	BmacConfigured        bool                     `protobuf:"varint,56,opt,name=bmac_configured,json=bmacConfigured,proto3" json:"bmac_configured,omitempty"`
	VniL3Flag             bool                     `protobuf:"varint,57,opt,name=vni_l3_flag,json=vniL3Flag,proto3" json:"vni_l3_flag,omitempty"`
	PbbBmac               string                   `protobuf:"bytes,58,opt,name=pbb_bmac,json=pbbBmac,proto3" json:"pbb_bmac,omitempty"`
	L3EncapsulationvlanId uint32                   `protobuf:"varint,59,opt,name=l3_encapsulationvlan_id,json=l3EncapsulationvlanId,proto3" json:"l3_encapsulationvlan_id,omitempty"`
	EvpnCtx               *L2FibBridgeMacEvpnCtxUn `protobuf:"bytes,60,opt,name=evpn_ctx,json=evpnCtx,proto3" json:"evpn_ctx,omitempty"`
	NextHopValid          bool                     `protobuf:"varint,61,opt,name=next_hop_valid,json=nextHopValid,proto3" json:"next_hop_valid,omitempty"`
	NextHop               *L2FibNhopInfo           `protobuf:"bytes,62,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	PeerVtepIp            string                   `protobuf:"bytes,63,opt,name=peer_vtep_ip,json=peerVtepIp,proto3" json:"peer_vtep_ip,omitempty"`
	FlagExtension         uint32                   `protobuf:"varint,64,opt,name=flag_extension,json=flagExtension,proto3" json:"flag_extension,omitempty"`
	Vni                   uint32                   `protobuf:"varint,65,opt,name=vni,proto3" json:"vni,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *L2FibBridgeMacInfo) Reset()         { *m = L2FibBridgeMacInfo{} }
func (m *L2FibBridgeMacInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBridgeMacInfo) ProtoMessage()    {}
func (*L2FibBridgeMacInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_778f8bd0df822369, []int{11}
}

func (m *L2FibBridgeMacInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBridgeMacInfo.Unmarshal(m, b)
}
func (m *L2FibBridgeMacInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBridgeMacInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBridgeMacInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBridgeMacInfo.Merge(m, src)
}
func (m *L2FibBridgeMacInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBridgeMacInfo.Size(m)
}
func (m *L2FibBridgeMacInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBridgeMacInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBridgeMacInfo proto.InternalMessageInfo

func (m *L2FibBridgeMacInfo) GetBase() *L2FibBaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *L2FibBridgeMacInfo) GetPlatformIsHwLearn() uint32 {
	if m != nil {
		return m.PlatformIsHwLearn
	}
	return 0
}

func (m *L2FibBridgeMacInfo) GetSegment() *L2FibBridgePortSegKeyUn {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *L2FibBridgeMacInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibBridgeMacInfo) GetBridgeDomainName() string {
	if m != nil {
		return m.BridgeDomainName
	}
	return ""
}

func (m *L2FibBridgeMacInfo) GetBridgeId() uint32 {
	if m != nil {
		return m.BridgeId
	}
	return 0
}

func (m *L2FibBridgeMacInfo) GetBmacConfigured() bool {
	if m != nil {
		return m.BmacConfigured
	}
	return false
}

func (m *L2FibBridgeMacInfo) GetVniL3Flag() bool {
	if m != nil {
		return m.VniL3Flag
	}
	return false
}

func (m *L2FibBridgeMacInfo) GetPbbBmac() string {
	if m != nil {
		return m.PbbBmac
	}
	return ""
}

func (m *L2FibBridgeMacInfo) GetL3EncapsulationvlanId() uint32 {
	if m != nil {
		return m.L3EncapsulationvlanId
	}
	return 0
}

func (m *L2FibBridgeMacInfo) GetEvpnCtx() *L2FibBridgeMacEvpnCtxUn {
	if m != nil {
		return m.EvpnCtx
	}
	return nil
}

func (m *L2FibBridgeMacInfo) GetNextHopValid() bool {
	if m != nil {
		return m.NextHopValid
	}
	return false
}

func (m *L2FibBridgeMacInfo) GetNextHop() *L2FibNhopInfo {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *L2FibBridgeMacInfo) GetPeerVtepIp() string {
	if m != nil {
		return m.PeerVtepIp
	}
	return ""
}

func (m *L2FibBridgeMacInfo) GetFlagExtension() uint32 {
	if m != nil {
		return m.FlagExtension
	}
	return 0
}

func (m *L2FibBridgeMacInfo) GetVni() uint32 {
	if m != nil {
		return m.Vni
	}
	return 0
}

func init() {
	proto.RegisterType((*L2FibBridgeMacInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_bridge_mac_info_KEYS")
	proto.RegisterType((*L2FibBaseInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_base_info")
	proto.RegisterType((*L2FibAcKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_ac_key_info")
	proto.RegisterType((*L2FibPbbKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_pbb_key_info")
	proto.RegisterType((*L2FibVniKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_vni_key_info")
	proto.RegisterType((*L2FibEvpnKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_evpn_key_info")
	proto.RegisterType((*L2FibPwKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_pw_key_info")
	proto.RegisterType((*L2FibBridgePortSegKeyUn)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_bridge_port_seg_key_un")
	proto.RegisterType((*L2VpnEvpnMoiInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2vpn_evpn_moi_info")
	proto.RegisterType((*L2FibBridgeMacEvpnCtxUn)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_bridge_mac_evpn_ctx_un")
	proto.RegisterType((*L2FibNhopInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_nhop_info")
	proto.RegisterType((*L2FibBridgeMacInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fibmac_details.l2fibmac_detail.l2fib_bridge_mac_info")
}

func init() { proto.RegisterFile("l2fib_bridge_mac_info.proto", fileDescriptor_778f8bd0df822369) }

var fileDescriptor_778f8bd0df822369 = []byte{
	// 1119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xdd, 0x6e, 0x23, 0x45,
	0x13, 0x95, 0xb3, 0x4e, 0x6c, 0x57, 0xd6, 0xd9, 0x6c, 0x67, 0xa3, 0xcc, 0xf7, 0x85, 0x9f, 0x95,
	0x05, 0x22, 0x88, 0xc5, 0x2b, 0xd9, 0xc4, 0xcb, 0x3f, 0x2c, 0x59, 0xa3, 0x1d, 0x11, 0x01, 0x1a,
	0x96, 0x20, 0xae, 0x5a, 0x3d, 0xd3, 0x6d, 0xbb, 0xc5, 0xb8, 0xa7, 0x35, 0x33, 0x1e, 0x3b, 0x48,
	0x48, 0xc0, 0x05, 0x57, 0x70, 0xc5, 0x2b, 0xf0, 0x06, 0xbc, 0x01, 0x0f, 0xc3, 0x73, 0xa0, 0xaa,
	0x9e, 0x71, 0xe2, 0x6c, 0xb4, 0x88, 0x0b, 0xef, 0x4d, 0x34, 0xae, 0x53, 0xdd, 0x55, 0xd5, 0x75,
	0x4e, 0x57, 0x07, 0x0e, 0xe3, 0xde, 0x48, 0x87, 0x3c, 0x4c, 0xb5, 0x1c, 0x2b, 0x3e, 0x15, 0x11,
	0xd7, 0x66, 0x94, 0x74, 0x6d, 0x9a, 0xe4, 0x09, 0x7b, 0x12, 0xe9, 0x2c, 0x4a, 0xb8, 0x4e, 0x32,
	0xbe, 0x48, 0x79, 0xdc, 0x2b, 0xac, 0xe1, 0x89, 0x55, 0x69, 0xd7, 0x7d, 0x8e, 0x92, 0x74, 0x2e,
	0x52, 0xa9, 0xcd, 0xb8, 0x6b, 0x12, 0xa9, 0x32, 0xfa, 0xdb, 0xa5, 0x0d, 0x71, 0x23, 0xa9, 0x72,
	0xa1, 0xe3, 0xec, 0xaa, 0xa1, 0xf3, 0x73, 0x0d, 0xfe, 0x7f, 0x6d, 0x54, 0xfe, 0xd9, 0xf0, 0xdb,
	0xaf, 0xd8, 0x01, 0x34, 0x70, 0x1f, 0xae, 0xa5, 0x57, 0xbb, 0x5b, 0x3b, 0x6a, 0x05, 0x5b, 0xf8,
	0xd3, 0x97, 0xcc, 0x83, 0x86, 0x90, 0x32, 0x55, 0x59, 0xe6, 0x6d, 0x10, 0x50, 0xfd, 0x64, 0x2f,
	0x02, 0x8c, 0xd3, 0x64, 0x66, 0xb9, 0x11, 0x53, 0xe5, 0xdd, 0x20, 0xb0, 0x45, 0x96, 0xcf, 0xc5,
	0x54, 0x31, 0x06, 0x75, 0x02, 0xea, 0x04, 0xd0, 0x77, 0xe7, 0x36, 0xdc, 0x2a, 0x73, 0x10, 0x99,
	0xa2, 0xe8, 0x9d, 0xaf, 0xe1, 0xb6, 0x33, 0x89, 0x88, 0x7f, 0xa7, 0xce, 0xc9, 0xc8, 0x5e, 0x85,
	0x1d, 0x6d, 0x72, 0x95, 0x8e, 0x44, 0xa4, 0xdc, 0xf6, 0x2e, 0xa9, 0xf6, 0xd2, 0x4a, 0x21, 0xfe,
	0x07, 0xcd, 0x22, 0xb3, 0xbc, 0x88, 0x85, 0xa1, 0xe4, 0xda, 0x41, 0xa3, 0xc8, 0xec, 0x59, 0x2c,
	0x4c, 0xe7, 0x08, 0x98, 0xdb, 0xd6, 0x86, 0xe1, 0xc5, 0xbe, 0x0c, 0xea, 0x8b, 0xa8, 0x2c, 0xb1,
	0x1d, 0xd0, 0x77, 0x67, 0x58, 0x79, 0x16, 0x46, 0x3f, 0xd3, 0x93, 0x1d, 0x42, 0xcb, 0x8a, 0x54,
	0x99, 0x9c, 0xeb, 0x51, 0x79, 0x18, 0x4d, 0x67, 0xf0, 0x47, 0x9d, 0xd7, 0x61, 0xcf, 0x6d, 0xa3,
	0xb0, 0x49, 0xcf, 0x8c, 0xf8, 0x53, 0xad, 0xaa, 0xd9, 0xce, 0x2f, 0x3c, 0xf7, 0x60, 0xd3, 0xce,
	0xab, 0xf3, 0xaf, 0x07, 0x75, 0x3b, 0xf7, 0x25, 0x3b, 0x82, 0x5d, 0xa3, 0x16, 0x39, 0x9f, 0x24,
	0x96, 0xaf, 0xb6, 0x61, 0x07, 0xed, 0x8f, 0x13, 0xfb, 0xb0, 0xec, 0xc6, 0x9b, 0xb0, 0x67, 0x33,
	0x35, 0x93, 0x09, 0x9f, 0xeb, 0x14, 0xfb, 0xc8, 0xf3, 0x73, 0x5b, 0xb5, 0x65, 0xd7, 0x41, 0xdf,
	0xe8, 0x54, 0xf9, 0xf2, 0xc9, 0xb9, 0x55, 0x9d, 0x3f, 0x36, 0xe1, 0x85, 0x15, 0x3a, 0xd8, 0x24,
	0xcd, 0x79, 0xa6, 0xc6, 0x94, 0xd0, 0xcc, 0x60, 0xb1, 0x52, 0xe4, 0xc2, 0xed, 0xe2, 0x4e, 0xbf,
	0x89, 0x06, 0x5c, 0xcd, 0xe6, 0xb0, 0x21, 0x22, 0x4a, 0x64, 0xbb, 0x37, 0xee, 0xae, 0x83, 0xaf,
	0xdd, 0xa7, 0x48, 0x11, 0x6c, 0x88, 0x88, 0x7d, 0x0f, 0x37, 0x6c, 0x18, 0x52, 0x55, 0xdb, 0xbd,
	0xc9, 0x3a, 0x23, 0x5f, 0xe6, 0x4d, 0x80, 0x41, 0x31, 0x76, 0x61, 0x34, 0xf1, 0x79, 0xcd, 0xb1,
	0x2f, 0x33, 0x31, 0xc0, 0xa0, 0xec, 0x07, 0xa8, 0x23, 0xaf, 0xbc, 0x4d, 0x0a, 0xae, 0xd7, 0x19,
	0x7c, 0x85, 0xbf, 0x01, 0x85, 0xc5, 0x7e, 0xdb, 0xb9, 0xb7, 0xb5, 0xfe, 0x7e, 0x5f, 0x12, 0x44,
	0xb0, 0x61, 0xe7, 0x9d, 0xbf, 0x6a, 0x28, 0x2b, 0xdc, 0x96, 0xd2, 0x9a, 0x26, 0xda, 0x89, 0xe5,
	0x1e, 0xb0, 0x7c, 0x66, 0x8c, 0x8a, 0xb9, 0x32, 0xd2, 0x26, 0x1a, 0x35, 0x59, 0x89, 0x6c, 0xd7,
	0x21, 0xc3, 0x12, 0xf0, 0x25, 0x7b, 0x03, 0xd8, 0x52, 0x45, 0xda, 0x16, 0x03, 0x92, 0x52, 0xa9,
	0xa3, 0x5b, 0xa5, 0x8e, 0x7c, 0x5b, 0x0c, 0x50, 0x4b, 0xec, 0x65, 0xd8, 0x9e, 0x46, 0x22, 0xcb,
	0x79, 0x2c, 0x42, 0x15, 0x13, 0xd5, 0xda, 0x01, 0x90, 0xe9, 0x14, 0x2d, 0xec, 0x3e, 0xec, 0x39,
	0x07, 0x65, 0x22, 0x61, 0xb3, 0x59, 0x2c, 0x72, 0x9d, 0x18, 0xe2, 0x45, 0x3b, 0x60, 0x04, 0x0d,
	0x2f, 0x23, 0x9d, 0xdf, 0x37, 0xae, 0x68, 0x0d, 0xcb, 0xa6, 0x7a, 0xa2, 0x7c, 0xf1, 0xaf, 0x5a,
	0xdb, 0x87, 0x2d, 0x95, 0x69, 0x2c, 0xcf, 0x5d, 0x71, 0x9b, 0x2a, 0xd3, 0xbe, 0xc4, 0x34, 0xe3,
	0x24, 0x12, 0xf1, 0x6a, 0x9a, 0x64, 0x72, 0x69, 0xfe, 0x52, 0x83, 0x96, 0xcb, 0x33, 0x89, 0x55,
	0xc9, 0xda, 0xb5, 0x11, 0xe7, 0xa9, 0x0e, 0x05, 0x4d, 0x8a, 0xfd, 0x45, 0x4c, 0x05, 0x84, 0x96,
	0xeb, 0xd1, 0x84, 0xd8, 0xdb, 0x0a, 0x36, 0x43, 0xeb, 0x8f, 0x26, 0x9d, 0x3f, 0x6f, 0x54, 0xc3,
	0xc0, 0x50, 0x5f, 0xb0, 0xad, 0xe7, 0x50, 0xc7, 0xc9, 0x40, 0x67, 0xb0, 0xdd, 0x53, 0xeb, 0x64,
	0xda, 0x72, 0x02, 0x05, 0x14, 0xf2, 0xbf, 0xdd, 0xb4, 0x57, 0x3d, 0x79, 0x31, 0xa8, 0x6e, 0xda,
	0x55, 0xe7, 0xb3, 0x01, 0x3b, 0x86, 0x83, 0x0b, 0xf2, 0xe1, 0xf8, 0x32, 0xcb, 0xa6, 0x39, 0xca,
	0xdc, 0xa9, 0x18, 0x58, 0x82, 0xae, 0x7d, 0xc7, 0x70, 0xa0, 0x22, 0xc9, 0x6d, 0x2c, 0x46, 0x49,
	0x3a, 0xcd, 0x39, 0x11, 0xa4, 0x10, 0xb1, 0x96, 0x74, 0x8c, 0xcd, 0xe0, 0x8e, 0x8a, 0xe4, 0x97,
	0x25, 0xfa, 0x48, 0xe4, 0xe2, 0x0c, 0x31, 0xf6, 0x00, 0xbc, 0x72, 0x59, 0x8e, 0x88, 0x5b, 0x16,
	0x2b, 0x33, 0xce, 0x27, 0xa4, 0xdf, 0x76, 0xb0, 0xef, 0xd6, 0x11, 0x8c, 0xeb, 0x4e, 0x09, 0xc4,
	0x91, 0x1b, 0x4d, 0x74, 0x2c, 0x53, 0x65, 0x78, 0x94, 0xcc, 0x4c, 0xee, 0x35, 0xc8, 0xbd, 0x5d,
	0x59, 0x4f, 0xd0, 0xd8, 0xf9, 0xbb, 0x01, 0xfb, 0xd7, 0x3e, 0x23, 0x96, 0xbd, 0xeb, 0x3d, 0xff,
	0xde, 0xdd, 0x87, 0x3b, 0xcb, 0x82, 0x75, 0xc6, 0x27, 0x73, 0x1e, 0x2b, 0x91, 0x1a, 0xaf, 0x4f,
	0x15, 0xdc, 0xae, 0x30, 0x3f, 0x7b, 0x3c, 0x3f, 0x45, 0x80, 0xfd, 0x5a, 0x83, 0x46, 0xa6, 0xc6,
	0x53, 0x65, 0x72, 0xef, 0x2d, 0xca, 0x37, 0x5d, 0x6b, 0xbe, 0xd7, 0x8e, 0xd8, 0xa0, 0x4a, 0xe1,
	0xf2, 0xe3, 0xeb, 0x78, 0xe5, 0xf1, 0x75, 0x0f, 0x58, 0xb9, 0x56, 0x26, 0x53, 0xa1, 0x8d, 0x7b,
	0x0b, 0x0d, 0x1c, 0xd3, 0x1c, 0xf2, 0x88, 0x00, 0x7a, 0x0e, 0x1d, 0x42, 0xab, 0xf4, 0xd6, 0xd2,
	0x7b, 0x40, 0xb5, 0x37, 0x9d, 0xc1, 0x97, 0xec, 0x35, 0xb8, 0x45, 0x89, 0x45, 0x89, 0x19, 0xe9,
	0xf1, 0x2c, 0x55, 0xd2, 0x7b, 0x9b, 0x78, 0xb4, 0x83, 0xe6, 0x93, 0xa5, 0x95, 0xbd, 0x04, 0xdb,
	0x38, 0x7f, 0xe2, 0x3e, 0x1f, 0xc5, 0x62, 0xec, 0xbd, 0x43, 0x4e, 0xad, 0xc2, 0xe8, 0xd3, 0xfe,
	0xa7, 0xb1, 0x18, 0xe3, 0xa3, 0x0b, 0x67, 0x23, 0xae, 0xf2, 0xde, 0x75, 0x2f, 0x42, 0x1b, 0x86,
	0x9f, 0x4c, 0x45, 0xc4, 0x06, 0x70, 0x10, 0xf7, 0x57, 0xaf, 0x45, 0x7c, 0x9b, 0x61, 0x3a, 0xef,
	0x39, 0xee, 0xc5, 0xfd, 0xe1, 0x55, 0xd4, 0x97, 0xec, 0xb7, 0x1a, 0x34, 0xab, 0xfb, 0xd0, 0x7b,
	0xff, 0xb9, 0xf5, 0xe3, 0xca, 0x35, 0x1c, 0x34, 0xf0, 0xc7, 0x49, 0xbe, 0x60, 0xaf, 0xc0, 0xce,
	0x52, 0xb2, 0x4e, 0x72, 0x1f, 0xd0, 0x29, 0xdc, 0x2c, 0x95, 0xea, 0xa4, 0xf6, 0x63, 0x0d, 0x9a,
	0x95, 0x9b, 0xf7, 0xe1, 0xfa, 0x59, 0xbf, 0xbc, 0x26, 0x83, 0x46, 0x99, 0x07, 0xbb, 0x0b, 0x37,
	0xad, 0x52, 0x29, 0x2f, 0x72, 0x85, 0x93, 0xcd, 0xfb, 0x88, 0xfa, 0x01, 0x68, 0x3b, 0xcb, 0x95,
	0xf5, 0x2d, 0xca, 0x1a, 0xdb, 0xc8, 0xd5, 0x22, 0x57, 0x26, 0xc3, 0x39, 0xf5, 0xb1, 0x93, 0x35,
	0x5a, 0x87, 0x95, 0x91, 0xed, 0xba, 0xb7, 0xcd, 0x43, 0xc2, 0xf0, 0x33, 0xdc, 0xa2, 0x7f, 0x46,
	0xfa, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x05, 0x5f, 0x02, 0xdf, 0xab, 0x0c, 0x00, 0x00,
}
