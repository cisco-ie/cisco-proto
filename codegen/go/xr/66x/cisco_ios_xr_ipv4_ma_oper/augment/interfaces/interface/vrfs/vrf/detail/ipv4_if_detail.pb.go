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
// source: ipv4_if_detail.proto

package cisco_ios_xr_ipv4_ma_oper_augment_interfaces_interface_vrfs_vrf_detail

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

type Ipv4IfDetail_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4IfDetail_KEYS) Reset()         { *m = Ipv4IfDetail_KEYS{} }
func (m *Ipv4IfDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfDetail_KEYS) ProtoMessage()    {}
func (*Ipv4IfDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{0}
}

func (m *Ipv4IfDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfDetail_KEYS.Unmarshal(m, b)
}
func (m *Ipv4IfDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4IfDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfDetail_KEYS.Merge(m, src)
}
func (m *Ipv4IfDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfDetail_KEYS.Size(m)
}
func (m *Ipv4IfDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfDetail_KEYS proto.InternalMessageInfo

func (m *Ipv4IfDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ipv4IfDetail_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type McastGroup struct {
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *McastGroup) Reset()         { *m = McastGroup{} }
func (m *McastGroup) String() string { return proto.CompactTextString(m) }
func (*McastGroup) ProtoMessage()    {}
func (*McastGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{1}
}

func (m *McastGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_McastGroup.Unmarshal(m, b)
}
func (m *McastGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_McastGroup.Marshal(b, m, deterministic)
}
func (m *McastGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_McastGroup.Merge(m, src)
}
func (m *McastGroup) XXX_Size() int {
	return xxx_messageInfo_McastGroup.Size(m)
}
func (m *McastGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_McastGroup.DiscardUnknown(m)
}

var xxx_messageInfo_McastGroup proto.InternalMessageInfo

func (m *McastGroup) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

type IpAddrNode struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrefixLength         uint32   `protobuf:"varint,2,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	RouteTag             uint32   `protobuf:"varint,3,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpAddrNode) Reset()         { *m = IpAddrNode{} }
func (m *IpAddrNode) String() string { return proto.CompactTextString(m) }
func (*IpAddrNode) ProtoMessage()    {}
func (*IpAddrNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{2}
}

func (m *IpAddrNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpAddrNode.Unmarshal(m, b)
}
func (m *IpAddrNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpAddrNode.Marshal(b, m, deterministic)
}
func (m *IpAddrNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpAddrNode.Merge(m, src)
}
func (m *IpAddrNode) XXX_Size() int {
	return xxx_messageInfo_IpAddrNode.Size(m)
}
func (m *IpAddrNode) XXX_DiscardUnknown() {
	xxx_messageInfo_IpAddrNode.DiscardUnknown(m)
}

var xxx_messageInfo_IpAddrNode proto.InternalMessageInfo

func (m *IpAddrNode) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IpAddrNode) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *IpAddrNode) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

type AclConfig struct {
	Inbound              string   `protobuf:"bytes,1,opt,name=inbound,proto3" json:"inbound,omitempty"`
	Outbound             string   `protobuf:"bytes,2,opt,name=outbound,proto3" json:"outbound,omitempty"`
	CommonInBound        string   `protobuf:"bytes,3,opt,name=common_in_bound,json=commonInBound,proto3" json:"common_in_bound,omitempty"`
	CommonOutBound       string   `protobuf:"bytes,4,opt,name=common_out_bound,json=commonOutBound,proto3" json:"common_out_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AclConfig) Reset()         { *m = AclConfig{} }
func (m *AclConfig) String() string { return proto.CompactTextString(m) }
func (*AclConfig) ProtoMessage()    {}
func (*AclConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{3}
}

func (m *AclConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclConfig.Unmarshal(m, b)
}
func (m *AclConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclConfig.Marshal(b, m, deterministic)
}
func (m *AclConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclConfig.Merge(m, src)
}
func (m *AclConfig) XXX_Size() int {
	return xxx_messageInfo_AclConfig.Size(m)
}
func (m *AclConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AclConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AclConfig proto.InternalMessageInfo

func (m *AclConfig) GetInbound() string {
	if m != nil {
		return m.Inbound
	}
	return ""
}

func (m *AclConfig) GetOutbound() string {
	if m != nil {
		return m.Outbound
	}
	return ""
}

func (m *AclConfig) GetCommonInBound() string {
	if m != nil {
		return m.CommonInBound
	}
	return ""
}

func (m *AclConfig) GetCommonOutBound() string {
	if m != nil {
		return m.CommonOutBound
	}
	return ""
}

type MultiAclConfig struct {
	Inbound              []string `protobuf:"bytes,1,rep,name=inbound,proto3" json:"inbound,omitempty"`
	Outbound             []string `protobuf:"bytes,2,rep,name=outbound,proto3" json:"outbound,omitempty"`
	Common               []string `protobuf:"bytes,3,rep,name=common,proto3" json:"common,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiAclConfig) Reset()         { *m = MultiAclConfig{} }
func (m *MultiAclConfig) String() string { return proto.CompactTextString(m) }
func (*MultiAclConfig) ProtoMessage()    {}
func (*MultiAclConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{4}
}

func (m *MultiAclConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiAclConfig.Unmarshal(m, b)
}
func (m *MultiAclConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiAclConfig.Marshal(b, m, deterministic)
}
func (m *MultiAclConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiAclConfig.Merge(m, src)
}
func (m *MultiAclConfig) XXX_Size() int {
	return xxx_messageInfo_MultiAclConfig.Size(m)
}
func (m *MultiAclConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiAclConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MultiAclConfig proto.InternalMessageInfo

func (m *MultiAclConfig) GetInbound() []string {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *MultiAclConfig) GetOutbound() []string {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *MultiAclConfig) GetCommon() []string {
	if m != nil {
		return m.Common
	}
	return nil
}

type HaddrArray struct {
	AddressArray         []string `protobuf:"bytes,1,rep,name=address_array,json=addressArray,proto3" json:"address_array,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HaddrArray) Reset()         { *m = HaddrArray{} }
func (m *HaddrArray) String() string { return proto.CompactTextString(m) }
func (*HaddrArray) ProtoMessage()    {}
func (*HaddrArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{5}
}

func (m *HaddrArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HaddrArray.Unmarshal(m, b)
}
func (m *HaddrArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HaddrArray.Marshal(b, m, deterministic)
}
func (m *HaddrArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HaddrArray.Merge(m, src)
}
func (m *HaddrArray) XXX_Size() int {
	return xxx_messageInfo_HaddrArray.Size(m)
}
func (m *HaddrArray) XXX_DiscardUnknown() {
	xxx_messageInfo_HaddrArray.DiscardUnknown(m)
}

var xxx_messageInfo_HaddrArray proto.InternalMessageInfo

func (m *HaddrArray) GetAddressArray() []string {
	if m != nil {
		return m.AddressArray
	}
	return nil
}

type RpfConfig struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	AllowDefaultRoute    bool     `protobuf:"varint,2,opt,name=allow_default_route,json=allowDefaultRoute,proto3" json:"allow_default_route,omitempty"`
	AllowSelfPing        bool     `protobuf:"varint,3,opt,name=allow_self_ping,json=allowSelfPing,proto3" json:"allow_self_ping,omitempty"`
	Mode                 string   `protobuf:"bytes,4,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpfConfig) Reset()         { *m = RpfConfig{} }
func (m *RpfConfig) String() string { return proto.CompactTextString(m) }
func (*RpfConfig) ProtoMessage()    {}
func (*RpfConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{6}
}

func (m *RpfConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpfConfig.Unmarshal(m, b)
}
func (m *RpfConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpfConfig.Marshal(b, m, deterministic)
}
func (m *RpfConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpfConfig.Merge(m, src)
}
func (m *RpfConfig) XXX_Size() int {
	return xxx_messageInfo_RpfConfig.Size(m)
}
func (m *RpfConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RpfConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RpfConfig proto.InternalMessageInfo

func (m *RpfConfig) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *RpfConfig) GetAllowDefaultRoute() bool {
	if m != nil {
		return m.AllowDefaultRoute
	}
	return false
}

func (m *RpfConfig) GetAllowSelfPing() bool {
	if m != nil {
		return m.AllowSelfPing
	}
	return false
}

func (m *RpfConfig) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type BgpPaDir struct {
	Enable               bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Source               bool     `protobuf:"varint,2,opt,name=source,proto3" json:"source,omitempty"`
	Destination          bool     `protobuf:"varint,3,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpPaDir) Reset()         { *m = BgpPaDir{} }
func (m *BgpPaDir) String() string { return proto.CompactTextString(m) }
func (*BgpPaDir) ProtoMessage()    {}
func (*BgpPaDir) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{7}
}

func (m *BgpPaDir) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPaDir.Unmarshal(m, b)
}
func (m *BgpPaDir) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPaDir.Marshal(b, m, deterministic)
}
func (m *BgpPaDir) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPaDir.Merge(m, src)
}
func (m *BgpPaDir) XXX_Size() int {
	return xxx_messageInfo_BgpPaDir.Size(m)
}
func (m *BgpPaDir) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPaDir.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPaDir proto.InternalMessageInfo

func (m *BgpPaDir) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *BgpPaDir) GetSource() bool {
	if m != nil {
		return m.Source
	}
	return false
}

func (m *BgpPaDir) GetDestination() bool {
	if m != nil {
		return m.Destination
	}
	return false
}

type BgpPaConfig struct {
	Input                *BgpPaDir `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output               *BgpPaDir `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BgpPaConfig) Reset()         { *m = BgpPaConfig{} }
func (m *BgpPaConfig) String() string { return proto.CompactTextString(m) }
func (*BgpPaConfig) ProtoMessage()    {}
func (*BgpPaConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{8}
}

func (m *BgpPaConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpPaConfig.Unmarshal(m, b)
}
func (m *BgpPaConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpPaConfig.Marshal(b, m, deterministic)
}
func (m *BgpPaConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpPaConfig.Merge(m, src)
}
func (m *BgpPaConfig) XXX_Size() int {
	return xxx_messageInfo_BgpPaConfig.Size(m)
}
func (m *BgpPaConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpPaConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BgpPaConfig proto.InternalMessageInfo

func (m *BgpPaConfig) GetInput() *BgpPaDir {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *BgpPaConfig) GetOutput() *BgpPaDir {
	if m != nil {
		return m.Output
	}
	return nil
}

type TimevalEntry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimevalEntry) Reset()         { *m = TimevalEntry{} }
func (m *TimevalEntry) String() string { return proto.CompactTextString(m) }
func (*TimevalEntry) ProtoMessage()    {}
func (*TimevalEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{9}
}

func (m *TimevalEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimevalEntry.Unmarshal(m, b)
}
func (m *TimevalEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimevalEntry.Marshal(b, m, deterministic)
}
func (m *TimevalEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimevalEntry.Merge(m, src)
}
func (m *TimevalEntry) XXX_Size() int {
	return xxx_messageInfo_TimevalEntry.Size(m)
}
func (m *TimevalEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TimevalEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TimevalEntry proto.InternalMessageInfo

type Ipv4IfDetail struct {
	PrimaryAddress              string          `protobuf:"bytes,50,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	VrfId                       uint32          `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	LineState                   string          `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	PrefixLength                uint32          `protobuf:"varint,53,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	RouteTag                    uint32          `protobuf:"varint,54,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	MulticastGroup              []*McastGroup   `protobuf:"bytes,55,rep,name=multicast_group,json=multicastGroup,proto3" json:"multicast_group,omitempty"`
	Mtu                         uint32          `protobuf:"varint,56,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Unreachable                 bool            `protobuf:"varint,57,opt,name=unreachable,proto3" json:"unreachable,omitempty"`
	Redirect                    bool            `protobuf:"varint,58,opt,name=redirect,proto3" json:"redirect,omitempty"`
	DirectBroadcast             bool            `protobuf:"varint,59,opt,name=direct_broadcast,json=directBroadcast,proto3" json:"direct_broadcast,omitempty"`
	MaskReply                   bool            `protobuf:"varint,60,opt,name=mask_reply,json=maskReply,proto3" json:"mask_reply,omitempty"`
	RgIdExists                  bool            `protobuf:"varint,61,opt,name=rg_id_exists,json=rgIdExists,proto3" json:"rg_id_exists,omitempty"`
	MlacpActive                 bool            `protobuf:"varint,62,opt,name=mlacp_active,json=mlacpActive,proto3" json:"mlacp_active,omitempty"`
	UnnumberedInterfaceName     string          `protobuf:"bytes,63,opt,name=unnumbered_interface_name,json=unnumberedInterfaceName,proto3" json:"unnumbered_interface_name,omitempty"`
	NextUnnumberedInterfaceName string          `protobuf:"bytes,64,opt,name=next_unnumbered_interface_name,json=nextUnnumberedInterfaceName,proto3" json:"next_unnumbered_interface_name,omitempty"`
	SecondaryAddress            []*IpAddrNode   `protobuf:"bytes,65,rep,name=secondary_address,json=secondaryAddress,proto3" json:"secondary_address,omitempty"`
	ProxyArpDisabled            bool            `protobuf:"varint,66,opt,name=proxy_arp_disabled,json=proxyArpDisabled,proto3" json:"proxy_arp_disabled,omitempty"`
	Acl                         *AclConfig      `protobuf:"bytes,67,opt,name=acl,proto3" json:"acl,omitempty"`
	MultiAcl                    *MultiAclConfig `protobuf:"bytes,68,opt,name=multi_acl,json=multiAcl,proto3" json:"multi_acl,omitempty"`
	HelperAddress               *HaddrArray     `protobuf:"bytes,69,opt,name=helper_address,json=helperAddress,proto3" json:"helper_address,omitempty"`
	Rpf                         *RpfConfig      `protobuf:"bytes,70,opt,name=rpf,proto3" json:"rpf,omitempty"`
	BgpPa                       *BgpPaConfig    `protobuf:"bytes,71,opt,name=bgp_pa,json=bgpPa,proto3" json:"bgp_pa,omitempty"`
	FlowTagSrc                  bool            `protobuf:"varint,72,opt,name=flow_tag_src,json=flowTagSrc,proto3" json:"flow_tag_src,omitempty"`
	FlowTagDst                  bool            `protobuf:"varint,73,opt,name=flow_tag_dst,json=flowTagDst,proto3" json:"flow_tag_dst,omitempty"`
	ConfigFlags                 uint32          `protobuf:"varint,74,opt,name=config_flags,json=configFlags,proto3" json:"config_flags,omitempty"`
	OperFlags                   uint64          `protobuf:"varint,75,opt,name=oper_flags,json=operFlags,proto3" json:"oper_flags,omitempty"`
	ArmFlags                    uint32          `protobuf:"varint,76,opt,name=arm_flags,json=armFlags,proto3" json:"arm_flags,omitempty"`
	StateRecvdFrmIm             string          `protobuf:"bytes,77,opt,name=state_recvd_frm_im,json=stateRecvdFrmIm,proto3" json:"state_recvd_frm_im,omitempty"`
	CflctAddress                string          `protobuf:"bytes,78,opt,name=cflct_address,json=cflctAddress,proto3" json:"cflct_address,omitempty"`
	ClientType                  string          `protobuf:"bytes,79,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	IsOrEvent                   bool            `protobuf:"varint,80,opt,name=is_or_event,json=isOrEvent,proto3" json:"is_or_event,omitempty"`
	OrImState                   string          `protobuf:"bytes,81,opt,name=or_im_state,json=orImState,proto3" json:"or_im_state,omitempty"`
	IdbPointer                  uint64          `protobuf:"varint,82,opt,name=idb_pointer,json=idbPointer,proto3" json:"idb_pointer,omitempty"`
	PubUtime                    *TimevalEntry   `protobuf:"bytes,83,opt,name=pub_utime,json=pubUtime,proto3" json:"pub_utime,omitempty"`
	IdbUtime                    *TimevalEntry   `protobuf:"bytes,84,opt,name=idb_utime,json=idbUtime,proto3" json:"idb_utime,omitempty"`
	CapsUtime                   *TimevalEntry   `protobuf:"bytes,85,opt,name=caps_utime,json=capsUtime,proto3" json:"caps_utime,omitempty"`
	FwdEnUtime                  *TimevalEntry   `protobuf:"bytes,86,opt,name=fwd_en_utime,json=fwdEnUtime,proto3" json:"fwd_en_utime,omitempty"`
	FwdDisUtime                 *TimevalEntry   `protobuf:"bytes,87,opt,name=fwd_dis_utime,json=fwdDisUtime,proto3" json:"fwd_dis_utime,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}        `json:"-"`
	XXX_unrecognized            []byte          `json:"-"`
	XXX_sizecache               int32           `json:"-"`
}

func (m *Ipv4IfDetail) Reset()         { *m = Ipv4IfDetail{} }
func (m *Ipv4IfDetail) String() string { return proto.CompactTextString(m) }
func (*Ipv4IfDetail) ProtoMessage()    {}
func (*Ipv4IfDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed75082a60326397, []int{10}
}

func (m *Ipv4IfDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4IfDetail.Unmarshal(m, b)
}
func (m *Ipv4IfDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4IfDetail.Marshal(b, m, deterministic)
}
func (m *Ipv4IfDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4IfDetail.Merge(m, src)
}
func (m *Ipv4IfDetail) XXX_Size() int {
	return xxx_messageInfo_Ipv4IfDetail.Size(m)
}
func (m *Ipv4IfDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4IfDetail.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4IfDetail proto.InternalMessageInfo

func (m *Ipv4IfDetail) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *Ipv4IfDetail) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *Ipv4IfDetail) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *Ipv4IfDetail) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *Ipv4IfDetail) GetRouteTag() uint32 {
	if m != nil {
		return m.RouteTag
	}
	return 0
}

func (m *Ipv4IfDetail) GetMulticastGroup() []*McastGroup {
	if m != nil {
		return m.MulticastGroup
	}
	return nil
}

func (m *Ipv4IfDetail) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Ipv4IfDetail) GetUnreachable() bool {
	if m != nil {
		return m.Unreachable
	}
	return false
}

func (m *Ipv4IfDetail) GetRedirect() bool {
	if m != nil {
		return m.Redirect
	}
	return false
}

func (m *Ipv4IfDetail) GetDirectBroadcast() bool {
	if m != nil {
		return m.DirectBroadcast
	}
	return false
}

func (m *Ipv4IfDetail) GetMaskReply() bool {
	if m != nil {
		return m.MaskReply
	}
	return false
}

func (m *Ipv4IfDetail) GetRgIdExists() bool {
	if m != nil {
		return m.RgIdExists
	}
	return false
}

func (m *Ipv4IfDetail) GetMlacpActive() bool {
	if m != nil {
		return m.MlacpActive
	}
	return false
}

func (m *Ipv4IfDetail) GetUnnumberedInterfaceName() string {
	if m != nil {
		return m.UnnumberedInterfaceName
	}
	return ""
}

func (m *Ipv4IfDetail) GetNextUnnumberedInterfaceName() string {
	if m != nil {
		return m.NextUnnumberedInterfaceName
	}
	return ""
}

func (m *Ipv4IfDetail) GetSecondaryAddress() []*IpAddrNode {
	if m != nil {
		return m.SecondaryAddress
	}
	return nil
}

func (m *Ipv4IfDetail) GetProxyArpDisabled() bool {
	if m != nil {
		return m.ProxyArpDisabled
	}
	return false
}

func (m *Ipv4IfDetail) GetAcl() *AclConfig {
	if m != nil {
		return m.Acl
	}
	return nil
}

func (m *Ipv4IfDetail) GetMultiAcl() *MultiAclConfig {
	if m != nil {
		return m.MultiAcl
	}
	return nil
}

func (m *Ipv4IfDetail) GetHelperAddress() *HaddrArray {
	if m != nil {
		return m.HelperAddress
	}
	return nil
}

func (m *Ipv4IfDetail) GetRpf() *RpfConfig {
	if m != nil {
		return m.Rpf
	}
	return nil
}

func (m *Ipv4IfDetail) GetBgpPa() *BgpPaConfig {
	if m != nil {
		return m.BgpPa
	}
	return nil
}

func (m *Ipv4IfDetail) GetFlowTagSrc() bool {
	if m != nil {
		return m.FlowTagSrc
	}
	return false
}

func (m *Ipv4IfDetail) GetFlowTagDst() bool {
	if m != nil {
		return m.FlowTagDst
	}
	return false
}

func (m *Ipv4IfDetail) GetConfigFlags() uint32 {
	if m != nil {
		return m.ConfigFlags
	}
	return 0
}

func (m *Ipv4IfDetail) GetOperFlags() uint64 {
	if m != nil {
		return m.OperFlags
	}
	return 0
}

func (m *Ipv4IfDetail) GetArmFlags() uint32 {
	if m != nil {
		return m.ArmFlags
	}
	return 0
}

func (m *Ipv4IfDetail) GetStateRecvdFrmIm() string {
	if m != nil {
		return m.StateRecvdFrmIm
	}
	return ""
}

func (m *Ipv4IfDetail) GetCflctAddress() string {
	if m != nil {
		return m.CflctAddress
	}
	return ""
}

func (m *Ipv4IfDetail) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *Ipv4IfDetail) GetIsOrEvent() bool {
	if m != nil {
		return m.IsOrEvent
	}
	return false
}

func (m *Ipv4IfDetail) GetOrImState() string {
	if m != nil {
		return m.OrImState
	}
	return ""
}

func (m *Ipv4IfDetail) GetIdbPointer() uint64 {
	if m != nil {
		return m.IdbPointer
	}
	return 0
}

func (m *Ipv4IfDetail) GetPubUtime() *TimevalEntry {
	if m != nil {
		return m.PubUtime
	}
	return nil
}

func (m *Ipv4IfDetail) GetIdbUtime() *TimevalEntry {
	if m != nil {
		return m.IdbUtime
	}
	return nil
}

func (m *Ipv4IfDetail) GetCapsUtime() *TimevalEntry {
	if m != nil {
		return m.CapsUtime
	}
	return nil
}

func (m *Ipv4IfDetail) GetFwdEnUtime() *TimevalEntry {
	if m != nil {
		return m.FwdEnUtime
	}
	return nil
}

func (m *Ipv4IfDetail) GetFwdDisUtime() *TimevalEntry {
	if m != nil {
		return m.FwdDisUtime
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4IfDetail_KEYS)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.ipv4_if_detail_KEYS")
	proto.RegisterType((*McastGroup)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.mcast_group")
	proto.RegisterType((*IpAddrNode)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.ip_addr_node")
	proto.RegisterType((*AclConfig)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.acl_config")
	proto.RegisterType((*MultiAclConfig)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.multi_acl_config")
	proto.RegisterType((*HaddrArray)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.haddr_array")
	proto.RegisterType((*RpfConfig)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.rpf_config")
	proto.RegisterType((*BgpPaDir)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.bgp_pa_dir")
	proto.RegisterType((*BgpPaConfig)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.bgp_pa_config")
	proto.RegisterType((*TimevalEntry)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.timeval_entry")
	proto.RegisterType((*Ipv4IfDetail)(nil), "cisco_ios_xr_ipv4_ma_oper.augment.interfaces.interface.vrfs.vrf.detail.ipv4_if_detail")
}

func init() { proto.RegisterFile("ipv4_if_detail.proto", fileDescriptor_ed75082a60326397) }

var fileDescriptor_ed75082a60326397 = []byte{
	// 1229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5f, 0x73, 0x1b, 0x35,
	0x10, 0x1f, 0xd7, 0xad, 0x6b, 0xaf, 0xed, 0x24, 0x55, 0xa1, 0xa8, 0x74, 0x68, 0x83, 0x19, 0x20,
	0x0c, 0x8c, 0x1f, 0xd2, 0xf2, 0xaf, 0xfc, 0x4d, 0x9b, 0xa4, 0x98, 0x96, 0x36, 0x5c, 0x12, 0x0a,
	0x2f, 0x08, 0xf9, 0xa4, 0xbb, 0xa8, 0xdc, 0xe9, 0x34, 0x3a, 0x9d, 0x13, 0x33, 0xbc, 0xf0, 0x0d,
	0x78, 0xe0, 0xbb, 0xf0, 0x89, 0xf8, 0x1e, 0x8c, 0x56, 0x67, 0x27, 0x2e, 0x93, 0xe1, 0xc5, 0x79,
	0xe9, 0x9c, 0x7e, 0xfb, 0xf3, 0xfe, 0x56, 0xbb, 0xda, 0xdd, 0x06, 0x5e, 0x51, 0x66, 0x72, 0x8f,
	0xa9, 0x84, 0x09, 0xe9, 0xb8, 0xca, 0x86, 0xc6, 0x16, 0xae, 0x20, 0xbb, 0xb1, 0x2a, 0xe3, 0x82,
	0xa9, 0xa2, 0x64, 0x27, 0x96, 0x21, 0x25, 0xe7, 0xac, 0x30, 0xd2, 0x0e, 0x79, 0x95, 0xe6, 0x52,
	0xbb, 0xa1, 0xd2, 0x4e, 0xda, 0x84, 0xc7, 0xb2, 0x3c, 0xfd, 0x1c, 0x4e, 0x6c, 0x52, 0xfa, 0x7f,
	0x86, 0xc1, 0xdb, 0xe0, 0x39, 0x5c, 0x5f, 0xf4, 0xcf, 0x1e, 0xef, 0xfc, 0xb4, 0x4f, 0xde, 0x86,
	0x95, 0xf9, 0x6f, 0x98, 0xe6, 0xb9, 0xa4, 0x8d, 0xf5, 0xc6, 0x46, 0x27, 0xea, 0xcf, 0xd1, 0xa7,
	0x3c, 0x97, 0xe4, 0x26, 0xb4, 0x27, 0x36, 0x09, 0x84, 0x4b, 0x48, 0xb8, 0x3a, 0xb1, 0x89, 0x37,
	0x0d, 0x36, 0xa1, 0x9b, 0xc7, 0xbc, 0x74, 0x2c, 0xb5, 0x45, 0x65, 0xc8, 0x5b, 0xd0, 0xc7, 0x0f,
	0xc6, 0x85, 0xb0, 0xb2, 0x2c, 0x6b, 0x7f, 0x3d, 0x04, 0xb7, 0x02, 0x36, 0x78, 0x01, 0x3d, 0x15,
	0x18, 0x4c, 0x17, 0x42, 0x12, 0x0a, 0x57, 0x17, 0xe9, 0xb3, 0xa3, 0x77, 0x67, 0xac, 0x4c, 0xd4,
	0x09, 0xcb, 0xa4, 0x4e, 0xdd, 0x11, 0xaa, 0xf7, 0xa3, 0x5e, 0x00, 0x9f, 0x20, 0x46, 0x6e, 0x41,
	0xc7, 0x16, 0x95, 0x93, 0xcc, 0xf1, 0x94, 0x36, 0x91, 0xd0, 0x46, 0xe0, 0x80, 0xa7, 0x83, 0xbf,
	0x1a, 0x00, 0x3c, 0xce, 0x58, 0x5c, 0xe8, 0x44, 0xa5, 0x5e, 0x4a, 0xe9, 0x71, 0x51, 0x69, 0x31,
	0x93, 0xaa, 0x8f, 0xe4, 0x75, 0x68, 0x17, 0x95, 0x0b, 0xa6, 0x70, 0xc7, 0xf9, 0x99, 0xbc, 0x03,
	0xab, 0x71, 0x91, 0xe7, 0x85, 0x66, 0x4a, 0xb3, 0x40, 0x69, 0x86, 0x3c, 0x05, 0x78, 0xa4, 0x1f,
	0x20, 0x6f, 0x03, 0xd6, 0x6a, 0x5e, 0x51, 0xb9, 0x9a, 0x78, 0x19, 0x89, 0x2b, 0x01, 0x7f, 0x56,
	0x39, 0x64, 0x0e, 0x7e, 0x81, 0xb5, 0xbc, 0xca, 0x9c, 0x62, 0xe7, 0xc5, 0xd6, 0x3c, 0x3f, 0xb6,
	0xe6, 0x42, 0x6c, 0x37, 0xa0, 0x15, 0x7c, 0xd3, 0x26, 0x5a, 0xea, 0x93, 0x2f, 0xcc, 0x11, 0xa6,
	0x98, 0x5b, 0xcb, 0xa7, 0x3e, 0x93, 0x75, 0x52, 0x03, 0x50, 0x4b, 0xf4, 0x6a, 0x70, 0xcb, 0x63,
	0x83, 0x3f, 0x1b, 0x00, 0xd6, 0x24, 0xb3, 0x80, 0x6e, 0x40, 0x4b, 0x6a, 0x3e, 0xce, 0xc2, 0xab,
	0x68, 0x47, 0xf5, 0x89, 0x0c, 0xe1, 0x3a, 0xcf, 0xb2, 0xe2, 0x98, 0x09, 0x99, 0xf0, 0x2a, 0x73,
	0x0c, 0xb3, 0x8d, 0x59, 0x6b, 0x47, 0xd7, 0xd0, 0xb4, 0x1d, 0x2c, 0x91, 0x37, 0xf8, 0xf4, 0x05,
	0x7e, 0x29, 0xb3, 0x84, 0x19, 0xa5, 0x43, 0x99, 0xda, 0x51, 0x1f, 0xe1, 0x7d, 0x99, 0x25, 0x7b,
	0x4a, 0xa7, 0x84, 0xc0, 0xe5, 0xbc, 0x10, 0xb2, 0x4e, 0x19, 0x7e, 0x0f, 0x7e, 0x06, 0x18, 0xa7,
	0x86, 0x19, 0xce, 0x84, 0xb2, 0xe7, 0x46, 0x74, 0x03, 0x5a, 0x65, 0x51, 0xd9, 0x78, 0x16, 0x44,
	0x7d, 0x22, 0xeb, 0xd0, 0x15, 0xb2, 0x74, 0x4a, 0x73, 0xa7, 0x30, 0x43, 0xde, 0x78, 0x16, 0x1a,
	0xfc, 0xd3, 0x80, 0x7e, 0x2d, 0x50, 0xdf, 0xfa, 0x08, 0xae, 0x28, 0x6d, 0x2a, 0x87, 0x12, 0xdd,
	0xcd, 0x68, 0xb8, 0x9c, 0x16, 0x1c, 0x9e, 0x5e, 0x23, 0x0a, 0x02, 0xe4, 0x05, 0xb4, 0x8a, 0xca,
	0x79, 0xa9, 0x4b, 0x17, 0x26, 0x55, 0x2b, 0x0c, 0x56, 0xa1, 0xef, 0x54, 0x2e, 0x27, 0x3c, 0x63,
	0x52, 0x3b, 0x3b, 0x1d, 0xfc, 0x4d, 0x60, 0x65, 0x71, 0x24, 0x90, 0x77, 0x61, 0xd5, 0x58, 0x95,
	0x73, 0x3b, 0x9d, 0xb7, 0xef, 0x66, 0x78, 0xbd, 0x35, 0x5c, 0x37, 0x30, 0x79, 0x15, 0x5a, 0x7e,
	0x1e, 0x28, 0x41, 0xef, 0x62, 0xbb, 0x5d, 0x99, 0xd8, 0x64, 0x24, 0xc8, 0x1b, 0x00, 0x99, 0xd2,
	0x92, 0x95, 0x8e, 0x3b, 0x49, 0xef, 0xe1, 0x4f, 0x3b, 0x1e, 0xd9, 0xf7, 0xc0, 0x7f, 0x9b, 0xf9,
	0xc3, 0xff, 0x6b, 0xe6, 0x8f, 0x16, 0x9b, 0x99, 0xfc, 0x0e, 0xab, 0xd8, 0x35, 0xa7, 0x03, 0x87,
	0x7e, 0xbc, 0xde, 0xdc, 0xe8, 0x6e, 0xee, 0x2f, 0x2b, 0x73, 0x67, 0x66, 0x59, 0xb4, 0x32, 0xd7,
	0x7a, 0x84, 0xb3, 0x6d, 0x0d, 0x9a, 0xb9, 0xab, 0xe8, 0x27, 0x18, 0x94, 0xff, 0xf4, 0xcf, 0xab,
	0xd2, 0x56, 0xf2, 0xf8, 0x08, 0xdf, 0xe4, 0xa7, 0xe1, 0x79, 0x9d, 0x81, 0x7c, 0xe7, 0x5a, 0x29,
	0x94, 0x95, 0xb1, 0xa3, 0xf7, 0xd1, 0x3c, 0x3f, 0x93, 0xf7, 0x60, 0x2d, 0x7c, 0xb1, 0xb1, 0x2d,
	0xb8, 0xf0, 0x42, 0xf4, 0x33, 0xe4, 0xac, 0x06, 0xfc, 0xc1, 0x0c, 0xf6, 0x99, 0xcd, 0x79, 0xf9,
	0x2b, 0xb3, 0xd2, 0x64, 0x53, 0xfa, 0x39, 0x92, 0x3a, 0x1e, 0x89, 0x3c, 0x40, 0xd6, 0xa1, 0x67,
	0x53, 0xa6, 0x04, 0x93, 0x27, 0xaa, 0x74, 0x25, 0xfd, 0x02, 0x09, 0x60, 0xd3, 0x91, 0xd8, 0x41,
	0x84, 0xbc, 0x09, 0xbd, 0x3c, 0xe3, 0xb1, 0x61, 0x3c, 0x76, 0x6a, 0x22, 0xe9, 0x97, 0x21, 0x54,
	0xc4, 0xb6, 0x10, 0x22, 0xf7, 0xe1, 0x66, 0xa5, 0x75, 0x95, 0x8f, 0xa5, 0x95, 0x82, 0xbd, 0xb4,
	0x16, 0xbe, 0xc2, 0x62, 0xbe, 0x76, 0x4a, 0x18, 0x2d, 0x2c, 0x88, 0x87, 0x70, 0x5b, 0xcb, 0x13,
	0xc7, 0xce, 0x77, 0xf0, 0x35, 0x3a, 0xb8, 0xe5, 0x59, 0x87, 0xe7, 0x38, 0xf9, 0xa3, 0x01, 0xd7,
	0x4a, 0x19, 0x17, 0x5a, 0x9c, 0x7d, 0x81, 0x5b, 0x58, 0xe0, 0x83, 0x65, 0x15, 0xf8, 0xec, 0xe2,
	0x89, 0xd6, 0xe6, 0x72, 0xb3, 0x97, 0xfd, 0x01, 0x10, 0x63, 0x8b, 0x93, 0x29, 0xe3, 0xd6, 0x30,
	0xa1, 0x4a, 0x5f, 0x44, 0x41, 0x1f, 0x60, 0xb6, 0xd6, 0xd0, 0xb2, 0x65, 0xcd, 0x76, 0x8d, 0x13,
	0x01, 0x4d, 0x1e, 0x67, 0xf4, 0xe1, 0x72, 0xbb, 0xf7, 0x74, 0x25, 0x44, 0xde, 0x3d, 0xa9, 0xa0,
	0x33, 0xdf, 0x15, 0x74, 0x1b, 0xb5, 0x7e, 0x5c, 0xda, 0x7b, 0x7f, 0x69, 0x09, 0x45, 0x6d, 0x44,
	0xb6, 0xe2, 0x8c, 0xfc, 0x06, 0x2b, 0x47, 0x32, 0x33, 0xd2, 0xce, 0x4b, 0xb1, 0x83, 0xda, 0x4b,
	0xeb, 0xb5, 0x33, 0xeb, 0x29, 0xea, 0x07, 0xa9, 0x59, 0x19, 0x04, 0x34, 0xad, 0x49, 0xe8, 0xee,
	0x72, 0x13, 0x7b, 0xba, 0xda, 0x22, 0xef, 0x9e, 0x64, 0xd0, 0x0a, 0x93, 0x92, 0x3e, 0x42, 0xa1,
	0xc3, 0x25, 0xcf, 0xdf, 0x5a, 0xeb, 0xca, 0x38, 0x35, 0x7b, 0xdc, 0x37, 0x69, 0xe2, 0x97, 0xa0,
	0xe3, 0x29, 0x2b, 0x6d, 0x4c, 0xbf, 0x09, 0x4d, 0xea, 0xb1, 0x03, 0x9e, 0xee, 0xdb, 0x78, 0x81,
	0x21, 0x4a, 0x47, 0x47, 0x0b, 0x8c, 0xed, 0xd2, 0xf9, 0x36, 0x0e, 0x4e, 0x59, 0x92, 0xf1, 0xb4,
	0xa4, 0xdf, 0xe2, 0x2c, 0xea, 0x06, 0x6c, 0xd7, 0x43, 0x7e, 0x54, 0xf8, 0x80, 0x6b, 0xc2, 0xe3,
	0xf5, 0xc6, 0xc6, 0xe5, 0xa8, 0xe3, 0x91, 0x60, 0xbe, 0x05, 0x1d, 0x6e, 0xf3, 0xda, 0xfa, 0x24,
	0xcc, 0x57, 0x6e, 0xf3, 0x60, 0x7c, 0x1f, 0x08, 0xce, 0x6e, 0x66, 0x65, 0x3c, 0x11, 0x2c, 0xb1,
	0x39, 0x53, 0x39, 0xfd, 0x0e, 0x5b, 0x77, 0x15, 0x2d, 0x91, 0x37, 0xec, 0xda, 0x7c, 0x94, 0xfb,
	0x71, 0x1e, 0x27, 0x59, 0xec, 0xe6, 0xcf, 0xe3, 0x69, 0xf8, 0xaf, 0x1e, 0x82, 0xb3, 0x42, 0xde,
	0x81, 0x6e, 0x9c, 0x29, 0xa9, 0x1d, 0x73, 0x53, 0x23, 0xe9, 0x33, 0xa4, 0x40, 0x80, 0x0e, 0xa6,
	0x46, 0x92, 0xdb, 0xd0, 0x55, 0x25, 0x2b, 0x2c, 0x93, 0x13, 0xa9, 0x1d, 0xdd, 0x0b, 0xa3, 0x4d,
	0x95, 0xcf, 0xec, 0x8e, 0x07, 0xbc, 0xbd, 0xb0, 0x4c, 0xe5, 0xf5, 0x52, 0xf9, 0x3e, 0x2c, 0x95,
	0xc2, 0x8e, 0xf2, 0xb0, 0x54, 0xee, 0x40, 0x57, 0x89, 0x31, 0x33, 0x05, 0x56, 0x83, 0x46, 0x78,
	0x5f, 0x50, 0x62, 0xbc, 0x17, 0x10, 0x62, 0xa1, 0x63, 0xaa, 0x31, 0xab, 0xfc, 0xf6, 0xa3, 0xfb,
	0xcb, 0xad, 0xf3, 0xc2, 0x46, 0x8d, 0xda, 0xa6, 0x1a, 0x1f, 0x7a, 0xc8, 0x6b, 0xfa, 0xa0, 0x82,
	0xe6, 0xc1, 0x85, 0x6a, 0x2a, 0x51, 0x6b, 0x3a, 0x80, 0x98, 0x9b, 0xb2, 0x16, 0x3d, 0xbc, 0x48,
	0xd1, 0x8e, 0x17, 0x0a, 0xaa, 0xc7, 0xd0, 0x4b, 0x8e, 0x05, 0x93, 0xba, 0xd6, 0xfd, 0xe1, 0x22,
	0x75, 0x21, 0x39, 0x16, 0x3b, 0x3a, 0x08, 0x4f, 0xa1, 0xef, 0x85, 0x85, 0x9a, 0xdd, 0xf8, 0xf9,
	0x45, 0x2a, 0x77, 0x93, 0x63, 0xb1, 0xad, 0xc2, 0x9d, 0xc7, 0x2d, 0xfc, 0xd3, 0xec, 0xee, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xe1, 0x5c, 0x65, 0xb2, 0x0d, 0x00, 0x00,
}
