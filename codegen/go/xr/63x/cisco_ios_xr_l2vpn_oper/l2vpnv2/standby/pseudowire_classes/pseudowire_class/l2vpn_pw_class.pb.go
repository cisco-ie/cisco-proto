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
// source: l2vpn_pw_class.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_pseudowire_classes_pseudowire_class

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

type L2VpnPwClass_KEYS struct {
	PseudowireClassName  string   `protobuf:"bytes,1,opt,name=pseudowire_class_name,json=pseudowireClassName,proto3" json:"pseudowire_class_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwClass_KEYS) Reset()         { *m = L2VpnPwClass_KEYS{} }
func (m *L2VpnPwClass_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwClass_KEYS) ProtoMessage()    {}
func (*L2VpnPwClass_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{0}
}

func (m *L2VpnPwClass_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwClass_KEYS.Unmarshal(m, b)
}
func (m *L2VpnPwClass_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwClass_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnPwClass_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwClass_KEYS.Merge(m, src)
}
func (m *L2VpnPwClass_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwClass_KEYS.Size(m)
}
func (m *L2VpnPwClass_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwClass_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwClass_KEYS proto.InternalMessageInfo

func (m *L2VpnPwClass_KEYS) GetPseudowireClassName() string {
	if m != nil {
		return m.PseudowireClassName
	}
	return ""
}

type L2VpnL2Tpv3PwClass struct {
	L2TpClassName        string   `protobuf:"bytes,1,opt,name=l2tp_class_name,json=l2tpClassName,proto3" json:"l2tp_class_name,omitempty"`
	Ipv4SourceAddress    string   `protobuf:"bytes,2,opt,name=ipv4_source_address,json=ipv4SourceAddress,proto3" json:"ipv4_source_address,omitempty"`
	PathMtuEnabled       bool     `protobuf:"varint,3,opt,name=path_mtu_enabled,json=pathMtuEnabled,proto3" json:"path_mtu_enabled,omitempty"`
	PathMtuMaxValue      uint32   `protobuf:"varint,4,opt,name=path_mtu_max_value,json=pathMtuMaxValue,proto3" json:"path_mtu_max_value,omitempty"`
	DontFragmentBit      bool     `protobuf:"varint,5,opt,name=dont_fragment_bit,json=dontFragmentBit,proto3" json:"dont_fragment_bit,omitempty"`
	TosMode              string   `protobuf:"bytes,6,opt,name=tos_mode,json=tosMode,proto3" json:"tos_mode,omitempty"`
	Tos                  uint32   `protobuf:"varint,7,opt,name=tos,proto3" json:"tos,omitempty"`
	Ttl                  uint32   `protobuf:"varint,8,opt,name=ttl,proto3" json:"ttl,omitempty"`
	CookieSize           uint32   `protobuf:"varint,9,opt,name=cookie_size,json=cookieSize,proto3" json:"cookie_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnL2Tpv3PwClass) Reset()         { *m = L2VpnL2Tpv3PwClass{} }
func (m *L2VpnL2Tpv3PwClass) String() string { return proto.CompactTextString(m) }
func (*L2VpnL2Tpv3PwClass) ProtoMessage()    {}
func (*L2VpnL2Tpv3PwClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{1}
}

func (m *L2VpnL2Tpv3PwClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnL2Tpv3PwClass.Unmarshal(m, b)
}
func (m *L2VpnL2Tpv3PwClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnL2Tpv3PwClass.Marshal(b, m, deterministic)
}
func (m *L2VpnL2Tpv3PwClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnL2Tpv3PwClass.Merge(m, src)
}
func (m *L2VpnL2Tpv3PwClass) XXX_Size() int {
	return xxx_messageInfo_L2VpnL2Tpv3PwClass.Size(m)
}
func (m *L2VpnL2Tpv3PwClass) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnL2Tpv3PwClass.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnL2Tpv3PwClass proto.InternalMessageInfo

func (m *L2VpnL2Tpv3PwClass) GetL2TpClassName() string {
	if m != nil {
		return m.L2TpClassName
	}
	return ""
}

func (m *L2VpnL2Tpv3PwClass) GetIpv4SourceAddress() string {
	if m != nil {
		return m.Ipv4SourceAddress
	}
	return ""
}

func (m *L2VpnL2Tpv3PwClass) GetPathMtuEnabled() bool {
	if m != nil {
		return m.PathMtuEnabled
	}
	return false
}

func (m *L2VpnL2Tpv3PwClass) GetPathMtuMaxValue() uint32 {
	if m != nil {
		return m.PathMtuMaxValue
	}
	return 0
}

func (m *L2VpnL2Tpv3PwClass) GetDontFragmentBit() bool {
	if m != nil {
		return m.DontFragmentBit
	}
	return false
}

func (m *L2VpnL2Tpv3PwClass) GetTosMode() string {
	if m != nil {
		return m.TosMode
	}
	return ""
}

func (m *L2VpnL2Tpv3PwClass) GetTos() uint32 {
	if m != nil {
		return m.Tos
	}
	return 0
}

func (m *L2VpnL2Tpv3PwClass) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *L2VpnL2Tpv3PwClass) GetCookieSize() uint32 {
	if m != nil {
		return m.CookieSize
	}
	return 0
}

type L2VpnEncapPwClass struct {
	Encapsulation        string              `protobuf:"bytes,1,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	L2Tpv3               *L2VpnL2Tpv3PwClass `protobuf:"bytes,2,opt,name=l2tpv3,proto3" json:"l2tpv3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2VpnEncapPwClass) Reset()         { *m = L2VpnEncapPwClass{} }
func (m *L2VpnEncapPwClass) String() string { return proto.CompactTextString(m) }
func (*L2VpnEncapPwClass) ProtoMessage()    {}
func (*L2VpnEncapPwClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{2}
}

func (m *L2VpnEncapPwClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEncapPwClass.Unmarshal(m, b)
}
func (m *L2VpnEncapPwClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEncapPwClass.Marshal(b, m, deterministic)
}
func (m *L2VpnEncapPwClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEncapPwClass.Merge(m, src)
}
func (m *L2VpnEncapPwClass) XXX_Size() int {
	return xxx_messageInfo_L2VpnEncapPwClass.Size(m)
}
func (m *L2VpnEncapPwClass) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEncapPwClass.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEncapPwClass proto.InternalMessageInfo

func (m *L2VpnEncapPwClass) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *L2VpnEncapPwClass) GetL2Tpv3() *L2VpnL2Tpv3PwClass {
	if m != nil {
		return m.L2Tpv3
	}
	return nil
}

type L2VpnMplsPrefpath struct {
	PolicyName           string   `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMplsPrefpath) Reset()         { *m = L2VpnMplsPrefpath{} }
func (m *L2VpnMplsPrefpath) String() string { return proto.CompactTextString(m) }
func (*L2VpnMplsPrefpath) ProtoMessage()    {}
func (*L2VpnMplsPrefpath) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{3}
}

func (m *L2VpnMplsPrefpath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMplsPrefpath.Unmarshal(m, b)
}
func (m *L2VpnMplsPrefpath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMplsPrefpath.Marshal(b, m, deterministic)
}
func (m *L2VpnMplsPrefpath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMplsPrefpath.Merge(m, src)
}
func (m *L2VpnMplsPrefpath) XXX_Size() int {
	return xxx_messageInfo_L2VpnMplsPrefpath.Size(m)
}
func (m *L2VpnMplsPrefpath) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMplsPrefpath.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMplsPrefpath proto.InternalMessageInfo

func (m *L2VpnMplsPrefpath) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

type L2VpnPwMplsPrefpath struct {
	Option                  string             `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	NextHopIp               uint32             `protobuf:"varint,2,opt,name=next_hop_ip,json=nextHopIp,proto3" json:"next_hop_ip,omitempty"`
	TeTunnelInterfaceNumber uint32             `protobuf:"varint,3,opt,name=te_tunnel_interface_number,json=teTunnelInterfaceNumber,proto3" json:"te_tunnel_interface_number,omitempty"`
	IpTunnelInterfaceNumber uint32             `protobuf:"varint,4,opt,name=ip_tunnel_interface_number,json=ipTunnelInterfaceNumber,proto3" json:"ip_tunnel_interface_number,omitempty"`
	TpTunnelInterfaceNumber uint32             `protobuf:"varint,5,opt,name=tp_tunnel_interface_number,json=tpTunnelInterfaceNumber,proto3" json:"tp_tunnel_interface_number,omitempty"`
	SrtePolicy              *L2VpnMplsPrefpath `protobuf:"bytes,6,opt,name=srte_policy,json=srtePolicy,proto3" json:"srte_policy,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}           `json:"-"`
	XXX_unrecognized        []byte             `json:"-"`
	XXX_sizecache           int32              `json:"-"`
}

func (m *L2VpnPwMplsPrefpath) Reset()         { *m = L2VpnPwMplsPrefpath{} }
func (m *L2VpnPwMplsPrefpath) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwMplsPrefpath) ProtoMessage()    {}
func (*L2VpnPwMplsPrefpath) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{4}
}

func (m *L2VpnPwMplsPrefpath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwMplsPrefpath.Unmarshal(m, b)
}
func (m *L2VpnPwMplsPrefpath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwMplsPrefpath.Marshal(b, m, deterministic)
}
func (m *L2VpnPwMplsPrefpath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwMplsPrefpath.Merge(m, src)
}
func (m *L2VpnPwMplsPrefpath) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwMplsPrefpath.Size(m)
}
func (m *L2VpnPwMplsPrefpath) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwMplsPrefpath.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwMplsPrefpath proto.InternalMessageInfo

func (m *L2VpnPwMplsPrefpath) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *L2VpnPwMplsPrefpath) GetNextHopIp() uint32 {
	if m != nil {
		return m.NextHopIp
	}
	return 0
}

func (m *L2VpnPwMplsPrefpath) GetTeTunnelInterfaceNumber() uint32 {
	if m != nil {
		return m.TeTunnelInterfaceNumber
	}
	return 0
}

func (m *L2VpnPwMplsPrefpath) GetIpTunnelInterfaceNumber() uint32 {
	if m != nil {
		return m.IpTunnelInterfaceNumber
	}
	return 0
}

func (m *L2VpnPwMplsPrefpath) GetTpTunnelInterfaceNumber() uint32 {
	if m != nil {
		return m.TpTunnelInterfaceNumber
	}
	return 0
}

func (m *L2VpnPwMplsPrefpath) GetSrtePolicy() *L2VpnMplsPrefpath {
	if m != nil {
		return m.SrtePolicy
	}
	return nil
}

type L2VpnSourceAddress struct {
	Configuration        string   `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnSourceAddress) Reset()         { *m = L2VpnSourceAddress{} }
func (m *L2VpnSourceAddress) String() string { return proto.CompactTextString(m) }
func (*L2VpnSourceAddress) ProtoMessage()    {}
func (*L2VpnSourceAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{5}
}

func (m *L2VpnSourceAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnSourceAddress.Unmarshal(m, b)
}
func (m *L2VpnSourceAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnSourceAddress.Marshal(b, m, deterministic)
}
func (m *L2VpnSourceAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnSourceAddress.Merge(m, src)
}
func (m *L2VpnSourceAddress) XXX_Size() int {
	return xxx_messageInfo_L2VpnSourceAddress.Size(m)
}
func (m *L2VpnSourceAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnSourceAddress.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnSourceAddress proto.InternalMessageInfo

func (m *L2VpnSourceAddress) GetConfiguration() string {
	if m != nil {
		return m.Configuration
	}
	return ""
}

func (m *L2VpnSourceAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type L2VpnPwClass struct {
	PwClassName                  string               `protobuf:"bytes,50,opt,name=pw_class_name,json=pwClassName,proto3" json:"pw_class_name,omitempty"`
	ControlWord                  string               `protobuf:"bytes,51,opt,name=control_word,json=controlWord,proto3" json:"control_word,omitempty"`
	TransportMode                string               `protobuf:"bytes,52,opt,name=transport_mode,json=transportMode,proto3" json:"transport_mode,omitempty"`
	SequencingType               string               `protobuf:"bytes,53,opt,name=sequencing_type,json=sequencingType,proto3" json:"sequencing_type,omitempty"`
	ResyncEnabled                bool                 `protobuf:"varint,54,opt,name=resync_enabled,json=resyncEnabled,proto3" json:"resync_enabled,omitempty"`
	ResyncThreshold              uint32               `protobuf:"varint,55,opt,name=resync_threshold,json=resyncThreshold,proto3" json:"resync_threshold,omitempty"`
	Protocol                     string               `protobuf:"bytes,56,opt,name=protocol,proto3" json:"protocol,omitempty"`
	EncapsulationInfo            *L2VpnEncapPwClass   `protobuf:"bytes,57,opt,name=encapsulation_info,json=encapsulationInfo,proto3" json:"encapsulation_info,omitempty"`
	DisableNever                 bool                 `protobuf:"varint,58,opt,name=disable_never,json=disableNever,proto3" json:"disable_never,omitempty"`
	DisableDelay                 uint32               `protobuf:"varint,59,opt,name=disable_delay,json=disableDelay,proto3" json:"disable_delay,omitempty"`
	BackupMacWithdraw            bool                 `protobuf:"varint,60,opt,name=backup_mac_withdraw,json=backupMacWithdraw,proto3" json:"backup_mac_withdraw,omitempty"`
	TagRewrite                   uint32               `protobuf:"varint,61,opt,name=tag_rewrite,json=tagRewrite,proto3" json:"tag_rewrite,omitempty"`
	PreferredPath                *L2VpnPwMplsPrefpath `protobuf:"bytes,62,opt,name=preferred_path,json=preferredPath,proto3" json:"preferred_path,omitempty"`
	PreferredPathDisableFallback bool                 `protobuf:"varint,63,opt,name=preferred_path_disable_fallback,json=preferredPathDisableFallback,proto3" json:"preferred_path_disable_fallback,omitempty"`
	LoadBalance                  string               `protobuf:"bytes,64,opt,name=load_balance,json=loadBalance,proto3" json:"load_balance,omitempty"`
	LocalSourceAddress           *L2VpnSourceAddress  `protobuf:"bytes,65,opt,name=local_source_address,json=localSourceAddress,proto3" json:"local_source_address,omitempty"`
	PwFlowLabelTypeCfg           string               `protobuf:"bytes,66,opt,name=pw_flow_label_type_cfg,json=pwFlowLabelTypeCfg,proto3" json:"pw_flow_label_type_cfg,omitempty"`
	PwFlowLabelCode17Disabled    bool                 `protobuf:"varint,67,opt,name=pw_flow_label_code17_disabled,json=pwFlowLabelCode17Disabled,proto3" json:"pw_flow_label_code17_disabled,omitempty"`
	IsFlowLabelStatic            bool                 `protobuf:"varint,68,opt,name=is_flow_label_static,json=isFlowLabelStatic,proto3" json:"is_flow_label_static,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}             `json:"-"`
	XXX_unrecognized             []byte               `json:"-"`
	XXX_sizecache                int32                `json:"-"`
}

func (m *L2VpnPwClass) Reset()         { *m = L2VpnPwClass{} }
func (m *L2VpnPwClass) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwClass) ProtoMessage()    {}
func (*L2VpnPwClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_40f9358f05810a55, []int{6}
}

func (m *L2VpnPwClass) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwClass.Unmarshal(m, b)
}
func (m *L2VpnPwClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwClass.Marshal(b, m, deterministic)
}
func (m *L2VpnPwClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwClass.Merge(m, src)
}
func (m *L2VpnPwClass) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwClass.Size(m)
}
func (m *L2VpnPwClass) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwClass.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwClass proto.InternalMessageInfo

func (m *L2VpnPwClass) GetPwClassName() string {
	if m != nil {
		return m.PwClassName
	}
	return ""
}

func (m *L2VpnPwClass) GetControlWord() string {
	if m != nil {
		return m.ControlWord
	}
	return ""
}

func (m *L2VpnPwClass) GetTransportMode() string {
	if m != nil {
		return m.TransportMode
	}
	return ""
}

func (m *L2VpnPwClass) GetSequencingType() string {
	if m != nil {
		return m.SequencingType
	}
	return ""
}

func (m *L2VpnPwClass) GetResyncEnabled() bool {
	if m != nil {
		return m.ResyncEnabled
	}
	return false
}

func (m *L2VpnPwClass) GetResyncThreshold() uint32 {
	if m != nil {
		return m.ResyncThreshold
	}
	return 0
}

func (m *L2VpnPwClass) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *L2VpnPwClass) GetEncapsulationInfo() *L2VpnEncapPwClass {
	if m != nil {
		return m.EncapsulationInfo
	}
	return nil
}

func (m *L2VpnPwClass) GetDisableNever() bool {
	if m != nil {
		return m.DisableNever
	}
	return false
}

func (m *L2VpnPwClass) GetDisableDelay() uint32 {
	if m != nil {
		return m.DisableDelay
	}
	return 0
}

func (m *L2VpnPwClass) GetBackupMacWithdraw() bool {
	if m != nil {
		return m.BackupMacWithdraw
	}
	return false
}

func (m *L2VpnPwClass) GetTagRewrite() uint32 {
	if m != nil {
		return m.TagRewrite
	}
	return 0
}

func (m *L2VpnPwClass) GetPreferredPath() *L2VpnPwMplsPrefpath {
	if m != nil {
		return m.PreferredPath
	}
	return nil
}

func (m *L2VpnPwClass) GetPreferredPathDisableFallback() bool {
	if m != nil {
		return m.PreferredPathDisableFallback
	}
	return false
}

func (m *L2VpnPwClass) GetLoadBalance() string {
	if m != nil {
		return m.LoadBalance
	}
	return ""
}

func (m *L2VpnPwClass) GetLocalSourceAddress() *L2VpnSourceAddress {
	if m != nil {
		return m.LocalSourceAddress
	}
	return nil
}

func (m *L2VpnPwClass) GetPwFlowLabelTypeCfg() string {
	if m != nil {
		return m.PwFlowLabelTypeCfg
	}
	return ""
}

func (m *L2VpnPwClass) GetPwFlowLabelCode17Disabled() bool {
	if m != nil {
		return m.PwFlowLabelCode17Disabled
	}
	return false
}

func (m *L2VpnPwClass) GetIsFlowLabelStatic() bool {
	if m != nil {
		return m.IsFlowLabelStatic
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnPwClass_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_pw_class_KEYS")
	proto.RegisterType((*L2VpnL2Tpv3PwClass)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_l2tpv3_pw_class")
	proto.RegisterType((*L2VpnEncapPwClass)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_encap_pw_class")
	proto.RegisterType((*L2VpnMplsPrefpath)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_mpls_prefpath")
	proto.RegisterType((*L2VpnPwMplsPrefpath)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_pw_mpls_prefpath")
	proto.RegisterType((*L2VpnSourceAddress)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_source_address")
	proto.RegisterType((*L2VpnPwClass)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.pseudowire_classes.pseudowire_class.l2vpn_pw_class")
}

func init() { proto.RegisterFile("l2vpn_pw_class.proto", fileDescriptor_40f9358f05810a55) }

var fileDescriptor_40f9358f05810a55 = []byte{
	// 994 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x73, 0x1c, 0x35,
	0x17, 0xad, 0x89, 0xbf, 0xf8, 0xa1, 0xc9, 0xf8, 0x21, 0x27, 0x8e, 0x92, 0xfa, 0x20, 0x66, 0x78,
	0x19, 0xa8, 0x1a, 0x8a, 0x71, 0x48, 0x80, 0xf0, 0x48, 0xfc, 0x2a, 0x5c, 0xc1, 0xae, 0xd4, 0xd8,
	0x95, 0x14, 0x2b, 0xa1, 0x51, 0xdf, 0x19, 0xab, 0xa2, 0x91, 0x84, 0xa4, 0x9e, 0xf6, 0x78, 0xc7,
	0x16, 0x36, 0x54, 0xf1, 0x67, 0xf8, 0x1d, 0xfc, 0x19, 0xb6, 0x94, 0xa4, 0xee, 0xb1, 0xdb, 0xc1,
	0x3b, 0xbc, 0xeb, 0x3e, 0xf7, 0x5c, 0xdd, 0xbe, 0x57, 0xe7, 0x9e, 0x46, 0xb7, 0x65, 0x77, 0x6c,
	0x14, 0x35, 0x05, 0xe5, 0x92, 0x39, 0xd7, 0x31, 0x56, 0x7b, 0x8d, 0x9f, 0x73, 0xe1, 0xb8, 0xa6,
	0x42, 0x3b, 0x7a, 0x6a, 0x69, 0xa2, 0x68, 0x03, 0xb6, 0x13, 0x1f, 0xc7, 0xdd, 0x8e, 0xf3, 0x4c,
	0x65, 0xfd, 0x49, 0xc7, 0x38, 0xc8, 0x33, 0x5d, 0x08, 0x0b, 0x29, 0x1f, 0xdc, 0x1b, 0x50, 0x7b,
	0x1f, 0xad, 0xd6, 0x8b, 0xd0, 0xe7, 0xbb, 0x3f, 0x1e, 0xe1, 0x2e, 0xba, 0x73, 0x99, 0x4a, 0x15,
	0x1b, 0x01, 0x69, 0xac, 0x37, 0x36, 0x16, 0x7a, 0xab, 0xe7, 0xc1, 0xed, 0x10, 0x3b, 0x64, 0x23,
	0x68, 0xff, 0x75, 0x03, 0xdd, 0x49, 0x67, 0xc9, 0xae, 0x37, 0xe3, 0xcd, 0xe9, 0x91, 0xf8, 0x03,
	0xb4, 0x14, 0xa0, 0x37, 0xcf, 0x69, 0x05, 0x78, 0x7a, 0x02, 0xee, 0xa0, 0x55, 0x61, 0xc6, 0x0f,
	0xa9, 0xd3, 0xb9, 0xe5, 0x40, 0x59, 0x96, 0x59, 0x70, 0x8e, 0xdc, 0x88, 0xdc, 0x95, 0x10, 0x3a,
	0x8a, 0x91, 0x67, 0x29, 0x80, 0x37, 0xd0, 0xb2, 0x61, 0xfe, 0x84, 0x8e, 0x7c, 0x4e, 0x41, 0xb1,
	0xbe, 0x84, 0x8c, 0xcc, 0xac, 0x37, 0x36, 0xe6, 0x7b, 0x8b, 0x01, 0x3f, 0xf0, 0xf9, 0x6e, 0x42,
	0xf1, 0x27, 0x08, 0x4f, 0x99, 0x23, 0x76, 0x4a, 0xc7, 0x4c, 0xe6, 0x40, 0xfe, 0xb7, 0xde, 0xd8,
	0x68, 0xf5, 0x96, 0x4a, 0xee, 0x01, 0x3b, 0x7d, 0x19, 0x60, 0xfc, 0x31, 0x5a, 0xc9, 0xb4, 0xf2,
	0x74, 0x60, 0xd9, 0x70, 0x04, 0xca, 0xd3, 0xbe, 0xf0, 0xe4, 0x66, 0x3c, 0x77, 0x29, 0x04, 0xf6,
	0x4a, 0x7c, 0x4b, 0x78, 0x7c, 0x0f, 0xcd, 0x7b, 0xed, 0xe8, 0x48, 0x67, 0x40, 0x66, 0xe3, 0x77,
	0xce, 0x79, 0xed, 0x0e, 0x74, 0x06, 0x78, 0x19, 0xcd, 0x78, 0xed, 0xc8, 0x5c, 0x2c, 0x12, 0x1e,
	0x23, 0xe2, 0x25, 0x99, 0x2f, 0x11, 0x2f, 0xf1, 0x03, 0xd4, 0xe4, 0x5a, 0xbf, 0x16, 0x40, 0x9d,
	0x38, 0x03, 0xb2, 0x10, 0x23, 0x28, 0x41, 0x47, 0xe2, 0x0c, 0xda, 0x7f, 0x36, 0x2a, 0x15, 0x80,
	0xe2, 0xcc, 0x9c, 0xcf, 0xf4, 0x3d, 0xd4, 0x8a, 0x88, 0xcb, 0x25, 0xf3, 0x42, 0xab, 0x6a, 0xa2,
	0x35, 0x10, 0x9f, 0xa1, 0xd9, 0x74, 0x19, 0x71, 0x88, 0xcd, 0x6e, 0xbf, 0xf3, 0x1f, 0x8a, 0xa7,
	0xf3, 0xaf, 0xb7, 0xdd, 0x2b, 0x2b, 0xb6, 0x1f, 0x55, 0xd2, 0x1a, 0x19, 0xe9, 0xa8, 0xb1, 0x30,
	0x08, 0x83, 0x0e, 0x2d, 0x1b, 0x2d, 0x05, 0x9f, 0x5c, 0x14, 0x02, 0x4a, 0x50, 0xd4, 0xd1, 0x6f,
	0x33, 0x68, 0x6d, 0xaa, 0xc9, 0x7a, 0xee, 0x1a, 0x9a, 0xd5, 0xe6, 0x42, 0xb7, 0xe5, 0x1b, 0x7e,
	0x1b, 0x35, 0x15, 0x9c, 0x7a, 0x7a, 0xa2, 0x0d, 0x15, 0x26, 0xf6, 0xda, 0xea, 0x2d, 0x04, 0xe8,
	0x7b, 0x6d, 0xf6, 0x0d, 0x7e, 0x82, 0xee, 0x7b, 0xa0, 0x3e, 0x57, 0x0a, 0x24, 0x15, 0xca, 0x83,
	0x1d, 0x30, 0x0e, 0x54, 0xe5, 0xa3, 0x3e, 0xd8, 0x28, 0x99, 0x56, 0xef, 0xae, 0x87, 0xe3, 0x48,
	0xd8, 0xaf, 0xe2, 0x87, 0x31, 0x1c, 0x92, 0x85, 0xb9, 0x32, 0x39, 0x69, 0xe8, 0xae, 0x30, 0x57,
	0x26, 0xfb, 0xab, 0x93, 0x6f, 0x96, 0x95, 0xaf, 0x48, 0xfe, 0xa5, 0x81, 0x9a, 0xce, 0x7a, 0xa0,
	0x69, 0x3a, 0x51, 0x60, 0xcd, 0xee, 0x4f, 0xd7, 0x70, 0x87, 0xb5, 0x31, 0xf7, 0x50, 0x28, 0xfa,
	0x22, 0xd6, 0x6c, 0xbf, 0xac, 0xf4, 0x57, 0x5f, 0xca, 0xa0, 0x3f, 0xae, 0xd5, 0x40, 0x0c, 0x73,
	0x5b, 0xd3, 0x5f, 0x0d, 0xc4, 0x04, 0xcd, 0xd5, 0xb7, 0xb8, 0x7a, 0x6d, 0xff, 0x3d, 0x8f, 0x16,
	0xeb, 0xce, 0x83, 0xdb, 0xa8, 0x35, 0x75, 0xa1, 0xa8, 0x8d, 0x6e, 0x4c, 0x69, 0x9a, 0xe2, 0xdc,
	0x22, 0xde, 0x41, 0xb7, 0xb8, 0x56, 0xde, 0x6a, 0x49, 0x0b, 0x6d, 0x33, 0xb2, 0x99, 0x28, 0x25,
	0xf6, 0x4a, 0xdb, 0x0c, 0xbf, 0x8f, 0x16, 0xbd, 0x65, 0xca, 0x19, 0x6d, 0x7d, 0x5a, 0xcc, 0x87,
	0xe9, 0xd3, 0xa6, 0x68, 0x5c, 0xcf, 0x0f, 0xd1, 0x92, 0x83, 0x9f, 0x73, 0x50, 0x5c, 0xa8, 0x21,
	0xf5, 0x13, 0x03, 0xe4, 0xf3, 0xc8, 0x5b, 0x3c, 0x87, 0x8f, 0x27, 0x06, 0xc2, 0x79, 0x16, 0xdc,
	0x44, 0xf1, 0xa9, 0xc7, 0x3c, 0x8a, 0x5e, 0xd0, 0x4a, 0x68, 0x65, 0x31, 0x1f, 0xa1, 0xe5, 0x92,
	0xe6, 0x4f, 0x2c, 0xb8, 0x13, 0x2d, 0x33, 0xf2, 0x38, 0x19, 0x4c, 0xc2, 0x8f, 0x2b, 0x18, 0xdf,
	0x47, 0xf3, 0xd1, 0xca, 0xb9, 0x96, 0xe4, 0x8b, 0x58, 0x73, 0xfa, 0x8e, 0x7f, 0x6f, 0x20, 0x5c,
	0xdb, 0x61, 0x2a, 0xd4, 0x40, 0x93, 0x2f, 0xe3, 0xd5, 0xb3, 0x6b, 0xb8, 0xfa, 0xba, 0xaf, 0xf4,
	0x56, 0x6a, 0xc5, 0xf7, 0xd5, 0x40, 0xe3, 0x77, 0x51, 0x2b, 0x13, 0x2e, 0x74, 0x49, 0x15, 0x8c,
	0xc1, 0x92, 0xaf, 0x62, 0xff, 0xb7, 0x4a, 0xf0, 0x30, 0x60, 0x17, 0x49, 0x19, 0x48, 0x36, 0x21,
	0x4f, 0x62, 0xef, 0x15, 0x69, 0x27, 0x60, 0xc1, 0xe0, 0xfb, 0x8c, 0xbf, 0xce, 0x0d, 0x1d, 0x31,
	0x4e, 0x0b, 0xe1, 0x4f, 0x32, 0xcb, 0x0a, 0xf2, 0x75, 0x3c, 0x6f, 0x25, 0x85, 0x0e, 0x18, 0x7f,
	0x55, 0x06, 0x82, 0x57, 0x78, 0x36, 0xa4, 0x16, 0x0a, 0x2b, 0x3c, 0x90, 0x6f, 0x92, 0x3d, 0x7a,
	0x36, 0xec, 0x25, 0x04, 0xff, 0xda, 0x40, 0x8b, 0x41, 0xb6, 0x60, 0x2d, 0x64, 0x34, 0x88, 0x97,
	0x7c, 0x1b, 0x27, 0xc5, 0xaf, 0x61, 0x52, 0x97, 0xed, 0xa8, 0xd7, 0x9a, 0x96, 0x7e, 0x11, 0xdc,
	0x69, 0x17, 0x3d, 0xa8, 0x7f, 0x0b, 0xad, 0x26, 0x32, 0x60, 0x52, 0x86, 0xee, 0xc8, 0x77, 0xb1,
	0xd3, 0xff, 0xd7, 0xf2, 0x76, 0x12, 0x69, 0xaf, 0xe4, 0x04, 0x89, 0x4b, 0xcd, 0x32, 0xda, 0x67,
	0x92, 0x29, 0x0e, 0xe4, 0x69, 0x92, 0x78, 0xc0, 0xb6, 0x12, 0x84, 0xff, 0x08, 0x7f, 0x05, 0xcd,
	0x99, 0xbc, 0xfc, 0xab, 0x7c, 0x76, 0x6d, 0x32, 0xa9, 0x17, 0xea, 0xe1, 0x58, 0xbe, 0xfe, 0x3b,
	0xee, 0xa2, 0x35, 0x53, 0xd0, 0x81, 0xd4, 0x05, 0x95, 0xac, 0x0f, 0x32, 0x2e, 0x15, 0xe5, 0x83,
	0x21, 0xd9, 0x8a, 0x2d, 0x60, 0x53, 0xec, 0x49, 0x5d, 0xfc, 0x10, 0x62, 0x61, 0xb3, 0xb6, 0x07,
	0x43, 0xfc, 0x14, 0xbd, 0x55, 0xcf, 0xe1, 0x3a, 0x83, 0xcf, 0x1e, 0x57, 0x93, 0xcb, 0xc8, 0x76,
	0x9c, 0xd8, 0xbd, 0x0b, 0xa9, 0xdb, 0x91, 0x51, 0x4e, 0x2d, 0xc3, 0x9f, 0xa2, 0xdb, 0xc2, 0x5d,
	0x3c, 0xc1, 0x79, 0xe6, 0x05, 0x27, 0x3b, 0x49, 0x54, 0xc2, 0x4d, 0x13, 0x8f, 0x62, 0xa0, 0x3f,
	0x1b, 0x77, 0x6d, 0xf3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x13, 0x6f, 0xf5, 0x5e, 0x09,
	0x00, 0x00,
}
