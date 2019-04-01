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
	PrimaryAddress          string          `protobuf:"bytes,50,opt,name=primary_address,json=primaryAddress,proto3" json:"primary_address,omitempty"`
	VrfId                   uint32          `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	LineState               string          `protobuf:"bytes,52,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	PrefixLength            uint32          `protobuf:"varint,53,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	RouteTag                uint32          `protobuf:"varint,54,opt,name=route_tag,json=routeTag,proto3" json:"route_tag,omitempty"`
	MulticastGroup          []*McastGroup   `protobuf:"bytes,55,rep,name=multicast_group,json=multicastGroup,proto3" json:"multicast_group,omitempty"`
	Mtu                     uint32          `protobuf:"varint,56,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Unreachable             bool            `protobuf:"varint,57,opt,name=unreachable,proto3" json:"unreachable,omitempty"`
	Redirect                bool            `protobuf:"varint,58,opt,name=redirect,proto3" json:"redirect,omitempty"`
	DirectBroadcast         bool            `protobuf:"varint,59,opt,name=direct_broadcast,json=directBroadcast,proto3" json:"direct_broadcast,omitempty"`
	MaskReply               bool            `protobuf:"varint,60,opt,name=mask_reply,json=maskReply,proto3" json:"mask_reply,omitempty"`
	RgIdExists              bool            `protobuf:"varint,61,opt,name=rg_id_exists,json=rgIdExists,proto3" json:"rg_id_exists,omitempty"`
	MlacpActive             bool            `protobuf:"varint,62,opt,name=mlacp_active,json=mlacpActive,proto3" json:"mlacp_active,omitempty"`
	UnnumberedInterfaceName string          `protobuf:"bytes,63,opt,name=unnumbered_interface_name,json=unnumberedInterfaceName,proto3" json:"unnumbered_interface_name,omitempty"`
	SecondaryAddress        []*IpAddrNode   `protobuf:"bytes,64,rep,name=secondary_address,json=secondaryAddress,proto3" json:"secondary_address,omitempty"`
	ProxyArpDisabled        bool            `protobuf:"varint,65,opt,name=proxy_arp_disabled,json=proxyArpDisabled,proto3" json:"proxy_arp_disabled,omitempty"`
	Acl                     *AclConfig      `protobuf:"bytes,66,opt,name=acl,proto3" json:"acl,omitempty"`
	MultiAcl                *MultiAclConfig `protobuf:"bytes,67,opt,name=multi_acl,json=multiAcl,proto3" json:"multi_acl,omitempty"`
	HelperAddress           *HaddrArray     `protobuf:"bytes,68,opt,name=helper_address,json=helperAddress,proto3" json:"helper_address,omitempty"`
	Rpf                     *RpfConfig      `protobuf:"bytes,69,opt,name=rpf,proto3" json:"rpf,omitempty"`
	BgpPa                   *BgpPaConfig    `protobuf:"bytes,70,opt,name=bgp_pa,json=bgpPa,proto3" json:"bgp_pa,omitempty"`
	FlowTagSrc              bool            `protobuf:"varint,71,opt,name=flow_tag_src,json=flowTagSrc,proto3" json:"flow_tag_src,omitempty"`
	FlowTagDst              bool            `protobuf:"varint,72,opt,name=flow_tag_dst,json=flowTagDst,proto3" json:"flow_tag_dst,omitempty"`
	PubUtime                *TimevalEntry   `protobuf:"bytes,73,opt,name=pub_utime,json=pubUtime,proto3" json:"pub_utime,omitempty"`
	IdbUtime                *TimevalEntry   `protobuf:"bytes,74,opt,name=idb_utime,json=idbUtime,proto3" json:"idb_utime,omitempty"`
	CapsUtime               *TimevalEntry   `protobuf:"bytes,75,opt,name=caps_utime,json=capsUtime,proto3" json:"caps_utime,omitempty"`
	FwdEnUtime              *TimevalEntry   `protobuf:"bytes,76,opt,name=fwd_en_utime,json=fwdEnUtime,proto3" json:"fwd_en_utime,omitempty"`
	FwdDisUtime             *TimevalEntry   `protobuf:"bytes,77,opt,name=fwd_dis_utime,json=fwdDisUtime,proto3" json:"fwd_dis_utime,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}        `json:"-"`
	XXX_unrecognized        []byte          `json:"-"`
	XXX_sizecache           int32           `json:"-"`
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
	// 1050 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0x1e, 0xd7, 0x8d, 0x6b, 0x9f, 0xd8, 0xb1, 0xab, 0x42, 0xd8, 0xc2, 0x30, 0x13, 0xcc, 0x00,
	0x61, 0x86, 0xf1, 0x45, 0x5a, 0xfe, 0xca, 0xaf, 0x43, 0xd2, 0x12, 0x5a, 0xa0, 0xb3, 0x49, 0x07,
	0xb8, 0x41, 0xc8, 0x2b, 0xed, 0x46, 0x65, 0x57, 0xab, 0xd1, 0x6a, 0x9d, 0x98, 0xe1, 0x86, 0x37,
	0xe0, 0x82, 0xf7, 0xe1, 0x69, 0x78, 0x0f, 0x46, 0x47, 0xf2, 0x5f, 0x99, 0x0c, 0x37, 0xce, 0x8d,
	0x47, 0xfa, 0xce, 0xb7, 0xfa, 0x8e, 0x8e, 0xce, 0xcf, 0x18, 0x5e, 0x92, 0x7a, 0x7a, 0x9f, 0xca,
	0x94, 0x72, 0x61, 0x99, 0xcc, 0x47, 0xda, 0x94, 0xb6, 0x24, 0x0f, 0x13, 0x59, 0x25, 0x25, 0x95,
	0x65, 0x45, 0x2f, 0x0d, 0x45, 0x4a, 0xc1, 0x68, 0xa9, 0x85, 0x19, 0xb1, 0x3a, 0x2b, 0x84, 0xb2,
	0x23, 0xa9, 0xac, 0x30, 0x29, 0x4b, 0x44, 0xb5, 0x5c, 0x8e, 0xa6, 0x26, 0xad, 0xdc, 0xcf, 0xc8,
	0x9f, 0x36, 0xfc, 0x01, 0xee, 0xac, 0x9f, 0x4f, 0x1f, 0x1f, 0xff, 0x74, 0x4a, 0xde, 0x82, 0x9d,
	0xc5, 0x37, 0x54, 0xb1, 0x42, 0x44, 0x8d, 0xbd, 0xc6, 0x7e, 0x27, 0xee, 0x2d, 0xd0, 0xef, 0x58,
	0x21, 0xc8, 0x5d, 0x68, 0x4f, 0x4d, 0xea, 0x09, 0x37, 0x90, 0x70, 0x6b, 0x6a, 0x52, 0x67, 0x1a,
	0x1e, 0xc0, 0x76, 0x91, 0xb0, 0xca, 0xd2, 0xcc, 0x94, 0xb5, 0x26, 0x6f, 0x42, 0x0f, 0x17, 0x94,
	0x71, 0x6e, 0x44, 0x55, 0x85, 0xf3, 0xba, 0x08, 0x8e, 0x3d, 0x36, 0x7c, 0x0e, 0x5d, 0xe9, 0x19,
	0x54, 0x95, 0x5c, 0x90, 0x08, 0x6e, 0xad, 0xd3, 0xe7, 0x5b, 0x77, 0x9c, 0x36, 0x22, 0x95, 0x97,
	0x34, 0x17, 0x2a, 0xb3, 0xe7, 0xa8, 0xde, 0x8b, 0xbb, 0x1e, 0x7c, 0x82, 0x18, 0x79, 0x0d, 0x3a,
	0xa6, 0xac, 0xad, 0xa0, 0x96, 0x65, 0x51, 0x13, 0x09, 0x6d, 0x04, 0xce, 0x58, 0x36, 0xfc, 0xab,
	0x01, 0xc0, 0x92, 0x9c, 0x26, 0xa5, 0x4a, 0x65, 0xe6, 0xa4, 0xa4, 0x9a, 0x94, 0xb5, 0xe2, 0x73,
	0xa9, 0xb0, 0x25, 0xaf, 0x42, 0xbb, 0xac, 0xad, 0x37, 0xf9, 0x3b, 0x2e, 0xf6, 0xe4, 0x6d, 0xe8,
	0x27, 0x65, 0x51, 0x94, 0x8a, 0x4a, 0x45, 0x3d, 0xa5, 0xe9, 0xe3, 0xe4, 0xe1, 0x13, 0x75, 0x88,
	0xbc, 0x7d, 0x18, 0x04, 0x5e, 0x59, 0xdb, 0x40, 0xbc, 0x89, 0xc4, 0x1d, 0x8f, 0x7f, 0x5f, 0x5b,
	0x64, 0x0e, 0x7f, 0x81, 0x41, 0x51, 0xe7, 0x56, 0xd2, 0xab, 0x7c, 0x6b, 0x5e, 0xed, 0x5b, 0x73,
	0xcd, 0xb7, 0x5d, 0x68, 0xf9, 0xb3, 0xa3, 0x26, 0x5a, 0xc2, 0xce, 0x3d, 0xcc, 0x39, 0x86, 0x98,
	0x19, 0xc3, 0x66, 0x2e, 0x92, 0x21, 0xa8, 0x1e, 0x08, 0x12, 0xdd, 0x00, 0x8e, 0x1d, 0x36, 0xfc,
	0xb3, 0x01, 0x60, 0x74, 0x3a, 0x77, 0x68, 0x17, 0x5a, 0x42, 0xb1, 0x49, 0xee, 0xb3, 0xa2, 0x1d,
	0x87, 0x1d, 0x19, 0xc1, 0x1d, 0x96, 0xe7, 0xe5, 0x05, 0xe5, 0x22, 0x65, 0x75, 0x6e, 0x29, 0x46,
	0x1b, 0xa3, 0xd6, 0x8e, 0x6f, 0xa3, 0xe9, 0xc8, 0x5b, 0x62, 0x67, 0x70, 0xe1, 0xf3, 0xfc, 0x4a,
	0xe4, 0x29, 0xd5, 0x52, 0xf9, 0x67, 0x6a, 0xc7, 0x3d, 0x84, 0x4f, 0x45, 0x9e, 0x3e, 0x95, 0x2a,
	0x23, 0x04, 0x6e, 0x16, 0x25, 0x17, 0x21, 0x64, 0xb8, 0x1e, 0xfe, 0x0c, 0x30, 0xc9, 0x34, 0xd5,
	0x8c, 0x72, 0x69, 0xae, 0xf4, 0x68, 0x17, 0x5a, 0x55, 0x59, 0x9b, 0x64, 0xee, 0x44, 0xd8, 0x91,
	0x3d, 0xd8, 0xe6, 0xa2, 0xb2, 0x52, 0x31, 0x2b, 0x31, 0x42, 0xce, 0xb8, 0x0a, 0x0d, 0xff, 0x69,
	0x40, 0x2f, 0x08, 0x84, 0x5b, 0x9f, 0xc3, 0x96, 0x54, 0xba, 0xb6, 0x28, 0xb1, 0x7d, 0x10, 0x8f,
	0x36, 0x53, 0x82, 0xa3, 0xe5, 0x35, 0x62, 0x2f, 0x40, 0x9e, 0x43, 0xab, 0xac, 0xad, 0x93, 0xba,
	0x71, 0x6d, 0x52, 0x41, 0x61, 0xd8, 0x87, 0x9e, 0x95, 0x85, 0x98, 0xb2, 0x9c, 0x0a, 0x65, 0xcd,
	0x6c, 0xf8, 0x77, 0x1f, 0x76, 0xd6, 0x5b, 0x02, 0x79, 0x07, 0xfa, 0xda, 0xc8, 0x82, 0x99, 0xd9,
	0xa2, 0x7c, 0x0f, 0x7c, 0xf6, 0x06, 0x38, 0x14, 0x30, 0x79, 0x19, 0x5a, 0xae, 0x1f, 0x48, 0x1e,
	0xdd, 0xc3, 0x72, 0xdb, 0x9a, 0x9a, 0xf4, 0x84, 0x93, 0xd7, 0x01, 0x72, 0xa9, 0x04, 0xad, 0x2c,
	0xb3, 0x22, 0xba, 0x8f, 0x9f, 0x76, 0x1c, 0x72, 0xea, 0x80, 0xff, 0x16, 0xf3, 0xfb, 0xff, 0x57,
	0xcc, 0x1f, 0xac, 0x17, 0x33, 0xf9, 0x1d, 0xfa, 0x58, 0x35, 0xcb, 0x86, 0x13, 0x7d, 0xb8, 0xd7,
	0xdc, 0xdf, 0x3e, 0x38, 0xdd, 0x54, 0xe4, 0x56, 0x7a, 0x59, 0xbc, 0xb3, 0xd0, 0x7a, 0x84, 0xbd,
	0x6d, 0x00, 0xcd, 0xc2, 0xd6, 0xd1, 0x47, 0xe8, 0x94, 0x5b, 0xba, 0xf4, 0xaa, 0x95, 0x11, 0x2c,
	0x39, 0xc7, 0x9c, 0xfc, 0xd8, 0xa7, 0xd7, 0x0a, 0xe4, 0x2a, 0xd7, 0x08, 0x2e, 0x8d, 0x48, 0x6c,
	0xf4, 0x00, 0xcd, 0x8b, 0x3d, 0x79, 0x17, 0x06, 0x7e, 0x45, 0x27, 0xa6, 0x64, 0xdc, 0x09, 0x45,
	0x9f, 0x20, 0xa7, 0xef, 0xf1, 0xc3, 0x39, 0xec, 0x22, 0x5b, 0xb0, 0xea, 0x57, 0x6a, 0x84, 0xce,
	0x67, 0xd1, 0xa7, 0x48, 0xea, 0x38, 0x24, 0x76, 0x00, 0xd9, 0x83, 0xae, 0xc9, 0xa8, 0xe4, 0x54,
	0x5c, 0xca, 0xca, 0x56, 0xd1, 0x67, 0x48, 0x00, 0x93, 0x9d, 0xf0, 0x63, 0x44, 0xc8, 0x1b, 0xd0,
	0x2d, 0x72, 0x96, 0x68, 0xca, 0x12, 0x2b, 0xa7, 0x22, 0xfa, 0xdc, 0xbb, 0x8a, 0xd8, 0x18, 0x21,
	0xf2, 0x00, 0xee, 0xd6, 0x4a, 0xd5, 0xc5, 0x44, 0x18, 0xc1, 0xe9, 0x0b, 0x63, 0xe1, 0x0b, 0x7c,
	0xcc, 0x57, 0x96, 0x84, 0x93, 0xb5, 0x01, 0xf1, 0x47, 0x03, 0x6e, 0x57, 0x22, 0x29, 0x15, 0x5f,
	0x4d, 0x9e, 0x2f, 0xf1, 0x6d, 0xce, 0x36, 0xf5, 0x36, 0xab, 0x33, 0x23, 0x1e, 0x2c, 0xe4, 0xe6,
	0x49, 0xf9, 0x1e, 0x10, 0x6d, 0xca, 0xcb, 0x19, 0x65, 0x46, 0x53, 0x2e, 0x2b, 0x17, 0x7f, 0x1e,
	0x8d, 0xf1, 0xa2, 0x03, 0xb4, 0x8c, 0x8d, 0x3e, 0x0a, 0x38, 0xe1, 0xd0, 0x64, 0x49, 0x1e, 0x1d,
	0x6e, 0xb6, 0xf0, 0x96, 0xdd, 0x3c, 0x76, 0xc7, 0x93, 0x1a, 0x3a, 0x8b, 0x36, 0x1f, 0x7d, 0x85,
	0x5a, 0x3f, 0x6e, 0x2c, 0x55, 0x5f, 0x98, 0x1f, 0x71, 0x1b, 0x91, 0x71, 0x92, 0x93, 0xdf, 0x60,
	0xe7, 0x5c, 0xe4, 0x5a, 0x98, 0xc5, 0x53, 0x1c, 0xa1, 0xf6, 0xc6, 0xca, 0x64, 0x65, 0xb2, 0xc4,
	0x3d, 0x2f, 0x35, 0x7f, 0x06, 0x0e, 0x4d, 0xa3, 0xd3, 0xe8, 0x78, 0xb3, 0x81, 0x5d, 0x4e, 0xa5,
	0xd8, 0x1d, 0x4f, 0x72, 0x68, 0xf9, 0x26, 0x17, 0x3d, 0x44, 0xa1, 0x67, 0x1b, 0x6e, 0x9d, 0x41,
	0x6b, 0x6b, 0x92, 0xe9, 0xa7, 0xcc, 0xd5, 0x57, 0xea, 0xe6, 0x97, 0x65, 0x19, 0xad, 0x4c, 0x12,
	0x3d, 0xf2, 0xf5, 0xe5, 0xb0, 0x33, 0x96, 0x9d, 0x9a, 0x64, 0x8d, 0xc1, 0x2b, 0x1b, 0x7d, 0xbd,
	0xc6, 0x38, 0xaa, 0x2c, 0x31, 0xd0, 0xd1, 0xf5, 0x84, 0xd6, 0xae, 0x0b, 0x47, 0x27, 0x9b, 0x75,
	0x7a, 0xad, 0xb3, 0xc7, 0x6d, 0x5d, 0x4f, 0x9e, 0x39, 0xc8, 0x69, 0x4a, 0x3e, 0xd7, 0xfc, 0xe6,
	0x5a, 0x35, 0x25, 0x0f, 0x9a, 0x16, 0x20, 0x61, 0xba, 0x0a, 0xa2, 0x8f, 0xaf, 0x53, 0xb4, 0xe3,
	0x84, 0xbc, 0xea, 0x05, 0x74, 0xd3, 0x0b, 0x4e, 0x85, 0x0a, 0xba, 0x4f, 0xae, 0x53, 0x17, 0xd2,
	0x0b, 0x7e, 0xac, 0xbc, 0xf0, 0x0c, 0x7a, 0x4e, 0x98, 0xcb, 0xf9, 0x8d, 0xbf, 0xbd, 0x4e, 0xe5,
	0xed, 0xf4, 0x82, 0x1f, 0x49, 0x7f, 0xe7, 0x49, 0x0b, 0xff, 0x22, 0xdc, 0xfb, 0x37, 0x00, 0x00,
	0xff, 0xff, 0xad, 0x64, 0x16, 0x33, 0x3a, 0x0c, 0x00, 0x00,
}