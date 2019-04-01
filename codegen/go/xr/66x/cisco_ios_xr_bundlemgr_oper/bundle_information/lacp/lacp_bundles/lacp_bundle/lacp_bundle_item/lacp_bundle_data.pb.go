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
// source: lacp_bundle_data.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_lacp_lacp_bundles_lacp_bundle_lacp_bundle_item

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

type LacpBundleData_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpBundleData_KEYS) Reset()         { *m = LacpBundleData_KEYS{} }
func (m *LacpBundleData_KEYS) String() string { return proto.CompactTextString(m) }
func (*LacpBundleData_KEYS) ProtoMessage()    {}
func (*LacpBundleData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{0}
}

func (m *LacpBundleData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpBundleData_KEYS.Unmarshal(m, b)
}
func (m *LacpBundleData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpBundleData_KEYS.Marshal(b, m, deterministic)
}
func (m *LacpBundleData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpBundleData_KEYS.Merge(m, src)
}
func (m *LacpBundleData_KEYS) XXX_Size() int {
	return xxx_messageInfo_LacpBundleData_KEYS.Size(m)
}
func (m *LacpBundleData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpBundleData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LacpBundleData_KEYS proto.InternalMessageInfo

func (m *LacpBundleData_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

type BmMacAddrSt struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmMacAddrSt) Reset()         { *m = BmMacAddrSt{} }
func (m *BmMacAddrSt) String() string { return proto.CompactTextString(m) }
func (*BmMacAddrSt) ProtoMessage()    {}
func (*BmMacAddrSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{1}
}

func (m *BmMacAddrSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMacAddrSt.Unmarshal(m, b)
}
func (m *BmMacAddrSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMacAddrSt.Marshal(b, m, deterministic)
}
func (m *BmMacAddrSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMacAddrSt.Merge(m, src)
}
func (m *BmMacAddrSt) XXX_Size() int {
	return xxx_messageInfo_BmMacAddrSt.Size(m)
}
func (m *BmMacAddrSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMacAddrSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmMacAddrSt proto.InternalMessageInfo

func (m *BmMacAddrSt) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type BmAddr struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmAddr) Reset()         { *m = BmAddr{} }
func (m *BmAddr) String() string { return proto.CompactTextString(m) }
func (*BmAddr) ProtoMessage()    {}
func (*BmAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{2}
}

func (m *BmAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmAddr.Unmarshal(m, b)
}
func (m *BmAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmAddr.Marshal(b, m, deterministic)
}
func (m *BmAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmAddr.Merge(m, src)
}
func (m *BmAddr) XXX_Size() int {
	return xxx_messageInfo_BmAddr.Size(m)
}
func (m *BmAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_BmAddr.DiscardUnknown(m)
}

var xxx_messageInfo_BmAddr proto.InternalMessageInfo

func (m *BmAddr) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *BmAddr) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *BmAddr) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type BmBundleBfdData struct {
	BundleStatus         string   `protobuf:"bytes,1,opt,name=bundle_status,json=bundleStatus,proto3" json:"bundle_status,omitempty"`
	StartTimer           uint32   `protobuf:"varint,2,opt,name=start_timer,json=startTimer,proto3" json:"start_timer,omitempty"`
	NbrUnconfigTimer     uint32   `protobuf:"varint,3,opt,name=nbr_unconfig_timer,json=nbrUnconfigTimer,proto3" json:"nbr_unconfig_timer,omitempty"`
	PrefMultiplier       uint32   `protobuf:"varint,4,opt,name=pref_multiplier,json=prefMultiplier,proto3" json:"pref_multiplier,omitempty"`
	PrefMinInterval      uint32   `protobuf:"varint,5,opt,name=pref_min_interval,json=prefMinInterval,proto3" json:"pref_min_interval,omitempty"`
	PrefEchoMinInterval  uint32   `protobuf:"varint,6,opt,name=pref_echo_min_interval,json=prefEchoMinInterval,proto3" json:"pref_echo_min_interval,omitempty"`
	FastDetect           uint32   `protobuf:"varint,7,opt,name=fast_detect,json=fastDetect,proto3" json:"fast_detect,omitempty"`
	DestinationAddress   *BmAddr  `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	ModeInfo             uint32   `protobuf:"varint,9,opt,name=mode_info,json=modeInfo,proto3" json:"mode_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmBundleBfdData) Reset()         { *m = BmBundleBfdData{} }
func (m *BmBundleBfdData) String() string { return proto.CompactTextString(m) }
func (*BmBundleBfdData) ProtoMessage()    {}
func (*BmBundleBfdData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{3}
}

func (m *BmBundleBfdData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmBundleBfdData.Unmarshal(m, b)
}
func (m *BmBundleBfdData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmBundleBfdData.Marshal(b, m, deterministic)
}
func (m *BmBundleBfdData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmBundleBfdData.Merge(m, src)
}
func (m *BmBundleBfdData) XXX_Size() int {
	return xxx_messageInfo_BmBundleBfdData.Size(m)
}
func (m *BmBundleBfdData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmBundleBfdData.DiscardUnknown(m)
}

var xxx_messageInfo_BmBundleBfdData proto.InternalMessageInfo

func (m *BmBundleBfdData) GetBundleStatus() string {
	if m != nil {
		return m.BundleStatus
	}
	return ""
}

func (m *BmBundleBfdData) GetStartTimer() uint32 {
	if m != nil {
		return m.StartTimer
	}
	return 0
}

func (m *BmBundleBfdData) GetNbrUnconfigTimer() uint32 {
	if m != nil {
		return m.NbrUnconfigTimer
	}
	return 0
}

func (m *BmBundleBfdData) GetPrefMultiplier() uint32 {
	if m != nil {
		return m.PrefMultiplier
	}
	return 0
}

func (m *BmBundleBfdData) GetPrefMinInterval() uint32 {
	if m != nil {
		return m.PrefMinInterval
	}
	return 0
}

func (m *BmBundleBfdData) GetPrefEchoMinInterval() uint32 {
	if m != nil {
		return m.PrefEchoMinInterval
	}
	return 0
}

func (m *BmBundleBfdData) GetFastDetect() uint32 {
	if m != nil {
		return m.FastDetect
	}
	return 0
}

func (m *BmBundleBfdData) GetDestinationAddress() *BmAddr {
	if m != nil {
		return m.DestinationAddress
	}
	return nil
}

func (m *BmBundleBfdData) GetModeInfo() uint32 {
	if m != nil {
		return m.ModeInfo
	}
	return 0
}

type BmBundleData struct {
	BundleInterfaceName             string             `protobuf:"bytes,1,opt,name=bundle_interface_name,json=bundleInterfaceName,proto3" json:"bundle_interface_name,omitempty"`
	MacAddress                      *BmMacAddrSt       `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	AvailableBandwidth              uint32             `protobuf:"varint,3,opt,name=available_bandwidth,json=availableBandwidth,proto3" json:"available_bandwidth,omitempty"`
	EffectiveBandwidth              uint32             `protobuf:"varint,4,opt,name=effective_bandwidth,json=effectiveBandwidth,proto3" json:"effective_bandwidth,omitempty"`
	ConfiguredBandwidth             uint32             `protobuf:"varint,5,opt,name=configured_bandwidth,json=configuredBandwidth,proto3" json:"configured_bandwidth,omitempty"`
	MinimumActiveLinks              uint32             `protobuf:"varint,6,opt,name=minimum_active_links,json=minimumActiveLinks,proto3" json:"minimum_active_links,omitempty"`
	MaximumActiveLinks              uint32             `protobuf:"varint,7,opt,name=maximum_active_links,json=maximumActiveLinks,proto3" json:"maximum_active_links,omitempty"`
	MaximumActiveLinksSource        string             `protobuf:"bytes,8,opt,name=maximum_active_links_source,json=maximumActiveLinksSource,proto3" json:"maximum_active_links_source,omitempty"`
	MinimumBandwidth                uint32             `protobuf:"varint,9,opt,name=minimum_bandwidth,json=minimumBandwidth,proto3" json:"minimum_bandwidth,omitempty"`
	PrimaryMember                   string             `protobuf:"bytes,10,opt,name=primary_member,json=primaryMember,proto3" json:"primary_member,omitempty"`
	BfdConfig                       []*BmBundleBfdData `protobuf:"bytes,11,rep,name=bfd_config,json=bfdConfig,proto3" json:"bfd_config,omitempty"`
	BundleStatus                    string             `protobuf:"bytes,12,opt,name=bundle_status,json=bundleStatus,proto3" json:"bundle_status,omitempty"`
	ActiveMemberCount               uint32             `protobuf:"varint,13,opt,name=active_member_count,json=activeMemberCount,proto3" json:"active_member_count,omitempty"`
	StandbyMemberCount              uint32             `protobuf:"varint,14,opt,name=standby_member_count,json=standbyMemberCount,proto3" json:"standby_member_count,omitempty"`
	ConfiguredMemberCount           uint32             `protobuf:"varint,15,opt,name=configured_member_count,json=configuredMemberCount,proto3" json:"configured_member_count,omitempty"`
	MacSource                       string             `protobuf:"bytes,16,opt,name=mac_source,json=macSource,proto3" json:"mac_source,omitempty"`
	MacSourceMember                 string             `protobuf:"bytes,17,opt,name=mac_source_member,json=macSourceMember,proto3" json:"mac_source_member,omitempty"`
	InterChassis                    uint32             `protobuf:"varint,18,opt,name=inter_chassis,json=interChassis,proto3" json:"inter_chassis,omitempty"`
	IsActive                        uint32             `protobuf:"varint,19,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	LacpStatus                      string             `protobuf:"bytes,20,opt,name=lacp_status,json=lacpStatus,proto3" json:"lacp_status,omitempty"`
	MlacpStatus                     string             `protobuf:"bytes,21,opt,name=mlacp_status,json=mlacpStatus,proto3" json:"mlacp_status,omitempty"`
	Ipv4BfdStatus                   string             `protobuf:"bytes,22,opt,name=ipv4bfd_status,json=ipv4bfdStatus,proto3" json:"ipv4bfd_status,omitempty"`
	LinkOrderStatus                 string             `protobuf:"bytes,23,opt,name=link_order_status,json=linkOrderStatus,proto3" json:"link_order_status,omitempty"`
	Ipv6BfdStatus                   string             `protobuf:"bytes,24,opt,name=ipv6bfd_status,json=ipv6bfdStatus,proto3" json:"ipv6bfd_status,omitempty"`
	LoadBalanceHashType             string             `protobuf:"bytes,25,opt,name=load_balance_hash_type,json=loadBalanceHashType,proto3" json:"load_balance_hash_type,omitempty"`
	LoadBalanceLocalityThreshold    uint32             `protobuf:"varint,26,opt,name=load_balance_locality_threshold,json=loadBalanceLocalityThreshold,proto3" json:"load_balance_locality_threshold,omitempty"`
	SuppressionTimer                uint32             `protobuf:"varint,27,opt,name=suppression_timer,json=suppressionTimer,proto3" json:"suppression_timer,omitempty"`
	WaitWhileTimer                  uint32             `protobuf:"varint,28,opt,name=wait_while_timer,json=waitWhileTimer,proto3" json:"wait_while_timer,omitempty"`
	CollectorMaxDelay               uint32             `protobuf:"varint,29,opt,name=collector_max_delay,json=collectorMaxDelay,proto3" json:"collector_max_delay,omitempty"`
	CiscoExtensions                 uint32             `protobuf:"varint,30,opt,name=cisco_extensions,json=ciscoExtensions,proto3" json:"cisco_extensions,omitempty"`
	LacpNonrevertive                uint32             `protobuf:"varint,31,opt,name=lacp_nonrevertive,json=lacpNonrevertive,proto3" json:"lacp_nonrevertive,omitempty"`
	IccpGroupId                     uint32             `protobuf:"varint,32,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	ActiveForeignMemberCount        uint32             `protobuf:"varint,33,opt,name=active_foreign_member_count,json=activeForeignMemberCount,proto3" json:"active_foreign_member_count,omitempty"`
	ConfiguredForeignMemberCount    uint32             `protobuf:"varint,34,opt,name=configured_foreign_member_count,json=configuredForeignMemberCount,proto3" json:"configured_foreign_member_count,omitempty"`
	SwitchoverType                  string             `protobuf:"bytes,35,opt,name=switchover_type,json=switchoverType,proto3" json:"switchover_type,omitempty"`
	MaximizeThresholdValueLinks     uint32             `protobuf:"varint,36,opt,name=maximize_threshold_value_links,json=maximizeThresholdValueLinks,proto3" json:"maximize_threshold_value_links,omitempty"`
	MaximizeThresholdValueBandWidth uint32             `protobuf:"varint,37,opt,name=maximize_threshold_value_band_width,json=maximizeThresholdValueBandWidth,proto3" json:"maximize_threshold_value_band_width,omitempty"`
	MlacpMode                       string             `protobuf:"bytes,38,opt,name=mlacp_mode,json=mlacpMode,proto3" json:"mlacp_mode,omitempty"`
	RecoveryDelay                   uint32             `protobuf:"varint,39,opt,name=recovery_delay,json=recoveryDelay,proto3" json:"recovery_delay,omitempty"`
	Singleton                       uint32             `protobuf:"varint,40,opt,name=singleton,proto3" json:"singleton,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}           `json:"-"`
	XXX_unrecognized                []byte             `json:"-"`
	XXX_sizecache                   int32              `json:"-"`
}

func (m *BmBundleData) Reset()         { *m = BmBundleData{} }
func (m *BmBundleData) String() string { return proto.CompactTextString(m) }
func (*BmBundleData) ProtoMessage()    {}
func (*BmBundleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{4}
}

func (m *BmBundleData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmBundleData.Unmarshal(m, b)
}
func (m *BmBundleData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmBundleData.Marshal(b, m, deterministic)
}
func (m *BmBundleData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmBundleData.Merge(m, src)
}
func (m *BmBundleData) XXX_Size() int {
	return xxx_messageInfo_BmBundleData.Size(m)
}
func (m *BmBundleData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmBundleData.DiscardUnknown(m)
}

var xxx_messageInfo_BmBundleData proto.InternalMessageInfo

func (m *BmBundleData) GetBundleInterfaceName() string {
	if m != nil {
		return m.BundleInterfaceName
	}
	return ""
}

func (m *BmBundleData) GetMacAddress() *BmMacAddrSt {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *BmBundleData) GetAvailableBandwidth() uint32 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

func (m *BmBundleData) GetEffectiveBandwidth() uint32 {
	if m != nil {
		return m.EffectiveBandwidth
	}
	return 0
}

func (m *BmBundleData) GetConfiguredBandwidth() uint32 {
	if m != nil {
		return m.ConfiguredBandwidth
	}
	return 0
}

func (m *BmBundleData) GetMinimumActiveLinks() uint32 {
	if m != nil {
		return m.MinimumActiveLinks
	}
	return 0
}

func (m *BmBundleData) GetMaximumActiveLinks() uint32 {
	if m != nil {
		return m.MaximumActiveLinks
	}
	return 0
}

func (m *BmBundleData) GetMaximumActiveLinksSource() string {
	if m != nil {
		return m.MaximumActiveLinksSource
	}
	return ""
}

func (m *BmBundleData) GetMinimumBandwidth() uint32 {
	if m != nil {
		return m.MinimumBandwidth
	}
	return 0
}

func (m *BmBundleData) GetPrimaryMember() string {
	if m != nil {
		return m.PrimaryMember
	}
	return ""
}

func (m *BmBundleData) GetBfdConfig() []*BmBundleBfdData {
	if m != nil {
		return m.BfdConfig
	}
	return nil
}

func (m *BmBundleData) GetBundleStatus() string {
	if m != nil {
		return m.BundleStatus
	}
	return ""
}

func (m *BmBundleData) GetActiveMemberCount() uint32 {
	if m != nil {
		return m.ActiveMemberCount
	}
	return 0
}

func (m *BmBundleData) GetStandbyMemberCount() uint32 {
	if m != nil {
		return m.StandbyMemberCount
	}
	return 0
}

func (m *BmBundleData) GetConfiguredMemberCount() uint32 {
	if m != nil {
		return m.ConfiguredMemberCount
	}
	return 0
}

func (m *BmBundleData) GetMacSource() string {
	if m != nil {
		return m.MacSource
	}
	return ""
}

func (m *BmBundleData) GetMacSourceMember() string {
	if m != nil {
		return m.MacSourceMember
	}
	return ""
}

func (m *BmBundleData) GetInterChassis() uint32 {
	if m != nil {
		return m.InterChassis
	}
	return 0
}

func (m *BmBundleData) GetIsActive() uint32 {
	if m != nil {
		return m.IsActive
	}
	return 0
}

func (m *BmBundleData) GetLacpStatus() string {
	if m != nil {
		return m.LacpStatus
	}
	return ""
}

func (m *BmBundleData) GetMlacpStatus() string {
	if m != nil {
		return m.MlacpStatus
	}
	return ""
}

func (m *BmBundleData) GetIpv4BfdStatus() string {
	if m != nil {
		return m.Ipv4BfdStatus
	}
	return ""
}

func (m *BmBundleData) GetLinkOrderStatus() string {
	if m != nil {
		return m.LinkOrderStatus
	}
	return ""
}

func (m *BmBundleData) GetIpv6BfdStatus() string {
	if m != nil {
		return m.Ipv6BfdStatus
	}
	return ""
}

func (m *BmBundleData) GetLoadBalanceHashType() string {
	if m != nil {
		return m.LoadBalanceHashType
	}
	return ""
}

func (m *BmBundleData) GetLoadBalanceLocalityThreshold() uint32 {
	if m != nil {
		return m.LoadBalanceLocalityThreshold
	}
	return 0
}

func (m *BmBundleData) GetSuppressionTimer() uint32 {
	if m != nil {
		return m.SuppressionTimer
	}
	return 0
}

func (m *BmBundleData) GetWaitWhileTimer() uint32 {
	if m != nil {
		return m.WaitWhileTimer
	}
	return 0
}

func (m *BmBundleData) GetCollectorMaxDelay() uint32 {
	if m != nil {
		return m.CollectorMaxDelay
	}
	return 0
}

func (m *BmBundleData) GetCiscoExtensions() uint32 {
	if m != nil {
		return m.CiscoExtensions
	}
	return 0
}

func (m *BmBundleData) GetLacpNonrevertive() uint32 {
	if m != nil {
		return m.LacpNonrevertive
	}
	return 0
}

func (m *BmBundleData) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

func (m *BmBundleData) GetActiveForeignMemberCount() uint32 {
	if m != nil {
		return m.ActiveForeignMemberCount
	}
	return 0
}

func (m *BmBundleData) GetConfiguredForeignMemberCount() uint32 {
	if m != nil {
		return m.ConfiguredForeignMemberCount
	}
	return 0
}

func (m *BmBundleData) GetSwitchoverType() string {
	if m != nil {
		return m.SwitchoverType
	}
	return ""
}

func (m *BmBundleData) GetMaximizeThresholdValueLinks() uint32 {
	if m != nil {
		return m.MaximizeThresholdValueLinks
	}
	return 0
}

func (m *BmBundleData) GetMaximizeThresholdValueBandWidth() uint32 {
	if m != nil {
		return m.MaximizeThresholdValueBandWidth
	}
	return 0
}

func (m *BmBundleData) GetMlacpMode() string {
	if m != nil {
		return m.MlacpMode
	}
	return ""
}

func (m *BmBundleData) GetRecoveryDelay() uint32 {
	if m != nil {
		return m.RecoveryDelay
	}
	return 0
}

func (m *BmBundleData) GetSingleton() uint32 {
	if m != nil {
		return m.Singleton
	}
	return 0
}

type EtherMacaddr_ struct {
	Macaddr              []uint32 `protobuf:"varint,1,rep,packed,name=macaddr,proto3" json:"macaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtherMacaddr_) Reset()         { *m = EtherMacaddr_{} }
func (m *EtherMacaddr_) String() string { return proto.CompactTextString(m) }
func (*EtherMacaddr_) ProtoMessage()    {}
func (*EtherMacaddr_) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{5}
}

func (m *EtherMacaddr_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtherMacaddr_.Unmarshal(m, b)
}
func (m *EtherMacaddr_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtherMacaddr_.Marshal(b, m, deterministic)
}
func (m *EtherMacaddr_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtherMacaddr_.Merge(m, src)
}
func (m *EtherMacaddr_) XXX_Size() int {
	return xxx_messageInfo_EtherMacaddr_.Size(m)
}
func (m *EtherMacaddr_) XXX_DiscardUnknown() {
	xxx_messageInfo_EtherMacaddr_.DiscardUnknown(m)
}

var xxx_messageInfo_EtherMacaddr_ proto.InternalMessageInfo

func (m *EtherMacaddr_) GetMacaddr() []uint32 {
	if m != nil {
		return m.Macaddr
	}
	return nil
}

type BmSystemIdSt struct {
	SystemPrio           uint32         `protobuf:"varint,1,opt,name=system_prio,json=systemPrio,proto3" json:"system_prio,omitempty"`
	SystemMacAddr        *EtherMacaddr_ `protobuf:"bytes,2,opt,name=system_mac_addr,json=systemMacAddr,proto3" json:"system_mac_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BmSystemIdSt) Reset()         { *m = BmSystemIdSt{} }
func (m *BmSystemIdSt) String() string { return proto.CompactTextString(m) }
func (*BmSystemIdSt) ProtoMessage()    {}
func (*BmSystemIdSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{6}
}

func (m *BmSystemIdSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmSystemIdSt.Unmarshal(m, b)
}
func (m *BmSystemIdSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmSystemIdSt.Marshal(b, m, deterministic)
}
func (m *BmSystemIdSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmSystemIdSt.Merge(m, src)
}
func (m *BmSystemIdSt) XXX_Size() int {
	return xxx_messageInfo_BmSystemIdSt.Size(m)
}
func (m *BmSystemIdSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmSystemIdSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmSystemIdSt proto.InternalMessageInfo

func (m *BmSystemIdSt) GetSystemPrio() uint32 {
	if m != nil {
		return m.SystemPrio
	}
	return 0
}

func (m *BmSystemIdSt) GetSystemMacAddr() *EtherMacaddr_ {
	if m != nil {
		return m.SystemMacAddr
	}
	return nil
}

type LacpBundleData struct {
	ActorBundleData         *BmBundleData `protobuf:"bytes,50,opt,name=actor_bundle_data,json=actorBundleData,proto3" json:"actor_bundle_data,omitempty"`
	BundleSystemId          *BmSystemIdSt `protobuf:"bytes,51,opt,name=bundle_system_id,json=bundleSystemId,proto3" json:"bundle_system_id,omitempty"`
	ActorOperationalKey     uint32        `protobuf:"varint,52,opt,name=actor_operational_key,json=actorOperationalKey,proto3" json:"actor_operational_key,omitempty"`
	PartnerSystemPriority   uint32        `protobuf:"varint,53,opt,name=partner_system_priority,json=partnerSystemPriority,proto3" json:"partner_system_priority,omitempty"`
	PartnerSystemMacAddress string        `protobuf:"bytes,54,opt,name=partner_system_mac_address,json=partnerSystemMacAddress,proto3" json:"partner_system_mac_address,omitempty"`
	PartnerOperationalKey   uint32        `protobuf:"varint,55,opt,name=partner_operational_key,json=partnerOperationalKey,proto3" json:"partner_operational_key,omitempty"`
	AppliedLacpMode         string        `protobuf:"bytes,56,opt,name=applied_lacp_mode,json=appliedLacpMode,proto3" json:"applied_lacp_mode,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}      `json:"-"`
	XXX_unrecognized        []byte        `json:"-"`
	XXX_sizecache           int32         `json:"-"`
}

func (m *LacpBundleData) Reset()         { *m = LacpBundleData{} }
func (m *LacpBundleData) String() string { return proto.CompactTextString(m) }
func (*LacpBundleData) ProtoMessage()    {}
func (*LacpBundleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0f92e1e33ba2c22, []int{7}
}

func (m *LacpBundleData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpBundleData.Unmarshal(m, b)
}
func (m *LacpBundleData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpBundleData.Marshal(b, m, deterministic)
}
func (m *LacpBundleData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpBundleData.Merge(m, src)
}
func (m *LacpBundleData) XXX_Size() int {
	return xxx_messageInfo_LacpBundleData.Size(m)
}
func (m *LacpBundleData) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpBundleData.DiscardUnknown(m)
}

var xxx_messageInfo_LacpBundleData proto.InternalMessageInfo

func (m *LacpBundleData) GetActorBundleData() *BmBundleData {
	if m != nil {
		return m.ActorBundleData
	}
	return nil
}

func (m *LacpBundleData) GetBundleSystemId() *BmSystemIdSt {
	if m != nil {
		return m.BundleSystemId
	}
	return nil
}

func (m *LacpBundleData) GetActorOperationalKey() uint32 {
	if m != nil {
		return m.ActorOperationalKey
	}
	return 0
}

func (m *LacpBundleData) GetPartnerSystemPriority() uint32 {
	if m != nil {
		return m.PartnerSystemPriority
	}
	return 0
}

func (m *LacpBundleData) GetPartnerSystemMacAddress() string {
	if m != nil {
		return m.PartnerSystemMacAddress
	}
	return ""
}

func (m *LacpBundleData) GetPartnerOperationalKey() uint32 {
	if m != nil {
		return m.PartnerOperationalKey
	}
	return 0
}

func (m *LacpBundleData) GetAppliedLacpMode() string {
	if m != nil {
		return m.AppliedLacpMode
	}
	return ""
}

func init() {
	proto.RegisterType((*LacpBundleData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.lacp_bundle_data_KEYS")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.bm_mac_addr_st")
	proto.RegisterType((*BmAddr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.bm_addr")
	proto.RegisterType((*BmBundleBfdData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.bm_bundle_bfd_data")
	proto.RegisterType((*BmBundleData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.bm_bundle_data")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.bm_system_id_st")
	proto.RegisterType((*LacpBundleData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_item.lacp_bundle_data")
}

func init() { proto.RegisterFile("lacp_bundle_data.proto", fileDescriptor_b0f92e1e33ba2c22) }

var fileDescriptor_b0f92e1e33ba2c22 = []byte{
	// 1413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xd1, 0x52, 0x1b, 0xc7,
	0x12, 0x2d, 0x81, 0xaf, 0x31, 0x2d, 0x24, 0x60, 0x64, 0xf0, 0x5e, 0x63, 0x1b, 0x2c, 0xae, 0xaf,
	0xb9, 0xdc, 0x94, 0x92, 0x80, 0x43, 0x52, 0x95, 0xca, 0x83, 0xc1, 0x24, 0xa1, 0x0c, 0x76, 0x4a,
	0x38, 0x71, 0xe5, 0x21, 0x35, 0x35, 0xda, 0x1d, 0xa1, 0x29, 0xef, 0xee, 0x6c, 0x66, 0x47, 0x80,
	0xf2, 0x11, 0xa9, 0xbc, 0xa4, 0x92, 0x3c, 0xe7, 0x6f, 0xf2, 0x01, 0xf9, 0x88, 0x7c, 0x45, 0xaa,
	0x7b, 0x66, 0xb5, 0x2b, 0x99, 0x3c, 0xe2, 0x37, 0xed, 0x39, 0xdd, 0x3d, 0x3d, 0xd3, 0xa7, 0x7b,
	0x46, 0xb0, 0x1a, 0x8b, 0x30, 0xe3, 0xbd, 0x61, 0x1a, 0xc5, 0x92, 0x47, 0xc2, 0x8a, 0x4e, 0x66,
	0xb4, 0xd5, 0xec, 0xbb, 0x50, 0xe5, 0xa1, 0xe6, 0x4a, 0xe7, 0xfc, 0xd2, 0x78, 0x3e, 0x39, 0x33,
	0x5c, 0x67, 0xd2, 0x74, 0xbc, 0xb9, 0x4a, 0xfb, 0xda, 0x24, 0xc2, 0x2a, 0x9d, 0x76, 0x30, 0x4c,
	0xa7, 0x12, 0x2b, 0xaf, 0x7e, 0x54, 0x7f, 0x73, 0x65, 0x65, 0xd2, 0xde, 0x87, 0x95, 0xe9, 0x85,
	0xf9, 0xf3, 0xc3, 0x6f, 0x4f, 0xd9, 0xff, 0x60, 0x69, 0x1c, 0xdd, 0x4a, 0xd3, 0x17, 0xa1, 0x0c,
	0x6a, 0x1b, 0xb5, 0xad, 0xf9, 0xee, 0xa2, 0xc3, 0x8f, 0x0a, 0xb8, 0xbd, 0x0d, 0xcd, 0x5e, 0xc2,
	0x13, 0x11, 0x72, 0x11, 0x45, 0x86, 0xe7, 0x96, 0x05, 0x30, 0x87, 0x3f, 0x65, 0x9e, 0x7b, 0x9f,
	0xe2, 0xb3, 0xfd, 0x14, 0xe6, 0x7a, 0x09, 0xd9, 0xb1, 0x26, 0xcc, 0x88, 0xbe, 0xe7, 0x67, 0x44,
	0x9f, 0x31, 0xb8, 0xa1, 0xb2, 0xf3, 0x27, 0xc1, 0x0c, 0x21, 0xf4, 0xdb, 0x63, 0x7b, 0xc1, 0xec,
	0x18, 0xdb, 0x6b, 0xff, 0x35, 0x0b, 0xac, 0x97, 0x14, 0x19, 0xf7, 0xfa, 0x11, 0x65, 0xcd, 0x36,
	0xa1, 0xe1, 0xa1, 0xdc, 0x0a, 0x3b, 0x2c, 0x56, 0x5e, 0x70, 0xe0, 0x29, 0x61, 0x6c, 0x1d, 0xea,
	0xb9, 0x15, 0xc6, 0x72, 0xab, 0x12, 0x69, 0x68, 0xa9, 0x46, 0x17, 0x08, 0x7a, 0x85, 0x08, 0x7b,
	0x0f, 0x58, 0xda, 0x33, 0x7c, 0x98, 0x86, 0x3a, 0xed, 0xab, 0x33, 0x6f, 0x37, 0x4b, 0x76, 0x4b,
	0x69, 0xcf, 0x7c, 0xed, 0x09, 0x67, 0xfd, 0x18, 0x16, 0x33, 0x23, 0xfb, 0x3c, 0x19, 0xc6, 0x56,
	0x65, 0xb1, 0x92, 0x26, 0xb8, 0x41, 0xa6, 0x4d, 0x84, 0x4f, 0xc6, 0x28, 0xdb, 0x86, 0x65, 0x67,
	0xa8, 0x52, 0x77, 0x9e, 0xe7, 0x22, 0x0e, 0xfe, 0x45, 0xa6, 0x14, 0xe1, 0x44, 0xa5, 0x47, 0x1e,
	0x66, 0xbb, 0xb0, 0x4a, 0xb6, 0x32, 0x1c, 0xe8, 0x49, 0x87, 0x9b, 0xe4, 0xd0, 0x42, 0xf6, 0x30,
	0x1c, 0xe8, 0xaa, 0xd3, 0x3a, 0xd4, 0xfb, 0x22, 0xb7, 0x3c, 0x92, 0x56, 0x86, 0x36, 0x98, 0x73,
	0x1b, 0x43, 0xe8, 0x19, 0x21, 0xec, 0x97, 0x1a, 0xb4, 0x22, 0x99, 0x5b, 0x95, 0x92, 0x4e, 0x78,
	0x51, 0x9f, 0x5b, 0x1b, 0xb5, 0xad, 0xfa, 0x4e, 0xbf, 0x73, 0xad, 0x32, 0xeb, 0xf8, 0x9a, 0x77,
	0x59, 0x25, 0x85, 0xa7, 0x2e, 0x03, 0xb6, 0x06, 0xf3, 0x89, 0x8e, 0x5c, 0xf8, 0x60, 0x9e, 0x12,
	0xbf, 0x85, 0xc0, 0x51, 0xda, 0xd7, 0xed, 0xdf, 0x97, 0x48, 0x5c, 0x15, 0x79, 0xb2, 0x1d, 0x58,
	0x99, 0x56, 0x26, 0x4f, 0x45, 0x52, 0xc8, 0xb3, 0x35, 0x25, 0xcf, 0x17, 0x22, 0x91, 0xec, 0xc7,
	0x1a, 0xd4, 0x0b, 0x81, 0xe2, 0xae, 0x67, 0x68, 0xd7, 0xc9, 0xf5, 0xef, 0xba, 0xd2, 0x15, 0x5d,
	0x48, 0x44, 0x58, 0x6c, 0xfa, 0x7d, 0x68, 0x89, 0x73, 0xa1, 0x62, 0xd1, 0x43, 0x0d, 0x8b, 0x34,
	0xba, 0x50, 0x91, 0x1d, 0x78, 0xa1, 0xb1, 0x31, 0xb5, 0x5f, 0x30, 0xe8, 0x20, 0xfb, 0x7d, 0x19,
	0x5a, 0x75, 0x5e, 0x75, 0x70, 0x72, 0x63, 0x63, 0xaa, 0x74, 0xf8, 0x10, 0x6e, 0x3b, 0xa9, 0x0e,
	0x8d, 0x8c, 0x2a, 0x1e, 0x4e, 0x75, 0xad, 0x92, 0x2b, 0x5d, 0x3e, 0x80, 0xdb, 0x89, 0x4a, 0x55,
	0x32, 0x4c, 0xb8, 0x70, 0x0b, 0xc5, 0x2a, 0x7d, 0x93, 0x7b, 0xdd, 0x31, 0xcf, 0x3d, 0x25, 0xea,
	0x18, 0x19, 0xf2, 0x10, 0x97, 0x6f, 0x7b, 0xcc, 0x79, 0x0f, 0xc7, 0x55, 0x3d, 0x3e, 0x83, 0xb5,
	0xab, 0x3c, 0x78, 0xae, 0x87, 0x26, 0x94, 0x24, 0xc7, 0xf9, 0x6e, 0xf0, 0xb6, 0xe3, 0x29, 0xf1,
	0xec, 0xff, 0xb0, 0x5c, 0xa4, 0x58, 0x6e, 0xc9, 0x89, 0x66, 0xc9, 0x13, 0xe5, 0x7e, 0x1e, 0x41,
	0x33, 0x33, 0x2a, 0x11, 0x66, 0xc4, 0x13, 0x99, 0xf4, 0xa4, 0x09, 0x80, 0xc2, 0x37, 0x3c, 0x7a,
	0x42, 0x20, 0xfb, 0xa9, 0x06, 0x80, 0x63, 0xc4, 0x1d, 0x49, 0x50, 0xdf, 0x98, 0xdd, 0xaa, 0xef,
	0x7c, 0x7f, 0xfd, 0xda, 0x98, 0x9a, 0x60, 0xdd, 0xf9, 0x5e, 0x3f, 0x3a, 0xa0, 0x1c, 0xde, 0x1e,
	0x66, 0x0b, 0x57, 0x0c, 0xb3, 0x0e, 0xb4, 0xfc, 0x11, 0xba, 0xdd, 0xf1, 0x50, 0x0f, 0x53, 0x1b,
	0x34, 0xe8, 0x34, 0x96, 0x1d, 0xe5, 0xb6, 0x78, 0x80, 0x04, 0x16, 0x2b, 0xb7, 0x22, 0x8d, 0x7a,
	0xa3, 0x49, 0x87, 0xa6, 0x2b, 0x96, 0xe7, 0xaa, 0x1e, 0x7b, 0x70, 0xa7, 0xa2, 0xa1, 0x09, 0xa7,
	0x45, 0x72, 0x5a, 0x29, 0xe9, 0xaa, 0xdf, 0x7d, 0x40, 0xad, 0x17, 0x35, 0x5d, 0xa2, 0xdc, 0xe7,
	0x13, 0x11, 0xfa, 0x22, 0x6e, 0xc3, 0x72, 0x49, 0x17, 0xa5, 0x59, 0x76, 0x97, 0xcb, 0xd8, 0xca,
	0x17, 0x67, 0x13, 0x1a, 0xd4, 0xe6, 0x3c, 0x1c, 0x88, 0x3c, 0x57, 0x79, 0xc0, 0x68, 0xe1, 0x05,
	0x02, 0x0f, 0x1c, 0x86, 0x23, 0x44, 0xe5, 0x5e, 0x4f, 0x41, 0xcb, 0x8d, 0x10, 0x95, 0x3b, 0xf5,
	0xe0, 0x68, 0xa4, 0xd3, 0xf7, 0x27, 0x79, 0x9b, 0xd6, 0x01, 0x84, 0xfc, 0x39, 0x3e, 0x84, 0x85,
	0xa4, 0x6a, 0xb1, 0x42, 0x16, 0xf5, 0xa4, 0x62, 0xf2, 0x08, 0x9a, 0x78, 0x1f, 0x61, 0xa9, 0xbc,
	0xd1, 0xaa, 0x53, 0x92, 0x47, 0xbd, 0xd9, 0x36, 0x2c, 0xa3, 0x9a, 0xb9, 0x36, 0x91, 0x34, 0x85,
	0xe5, 0x1d, 0xb7, 0x31, 0x24, 0x5e, 0x22, 0x3e, 0x11, 0x72, 0xaf, 0x12, 0x32, 0x18, 0x87, 0xdc,
	0x2b, 0x43, 0xee, 0xc2, 0x6a, 0xac, 0x05, 0x36, 0x70, 0x2c, 0xd2, 0x50, 0xf2, 0x81, 0xc8, 0x07,
	0xdc, 0x8e, 0x32, 0x19, 0xfc, 0xdb, 0x8d, 0x3b, 0x64, 0xf7, 0x1d, 0xf9, 0xa5, 0xc8, 0x07, 0xaf,
	0x46, 0x99, 0x64, 0x87, 0xb0, 0x3e, 0xe1, 0x14, 0xeb, 0x50, 0xc4, 0xca, 0x8e, 0xb8, 0x1d, 0x18,
	0x99, 0x0f, 0x74, 0x1c, 0x05, 0x77, 0xe9, 0x94, 0xee, 0x55, 0xbc, 0x8f, 0xbd, 0xd1, 0xab, 0xc2,
	0x06, 0x9b, 0x2d, 0x1f, 0x66, 0x19, 0x0e, 0x2c, 0xbc, 0x32, 0xdc, 0x5d, 0xb8, 0xe6, 0x9a, 0xad,
	0x42, 0xb8, 0xbb, 0x70, 0x0b, 0x96, 0x2e, 0x84, 0xb2, 0xfc, 0x62, 0xa0, 0x62, 0xe9, 0x6d, 0xef,
	0xb9, 0xcb, 0x10, 0xf1, 0xd7, 0x08, 0x3b, 0xcb, 0x0e, 0xb4, 0x42, 0x1d, 0xc7, 0x32, 0xb4, 0xda,
	0xf0, 0x44, 0x5c, 0xf2, 0x48, 0xc6, 0x62, 0x14, 0xdc, 0x77, 0xba, 0x1d, 0x53, 0x27, 0xe2, 0xf2,
	0x19, 0x12, 0xf8, 0x14, 0x71, 0xbd, 0x28, 0x2f, 0xad, 0x4c, 0x71, 0xc5, 0x3c, 0x78, 0xe0, 0xee,
	0x4e, 0xc2, 0x0f, 0xc7, 0x30, 0x66, 0x4c, 0x95, 0x4c, 0x75, 0x6a, 0xe4, 0xb9, 0x34, 0x24, 0x88,
	0x75, 0x97, 0x31, 0x12, 0x2f, 0x2a, 0x38, 0x6b, 0x43, 0x43, 0x85, 0x61, 0xc6, 0xcf, 0x8c, 0x1e,
	0x66, 0x5c, 0x45, 0xc1, 0x06, 0x19, 0xd6, 0x11, 0xfc, 0x02, 0xb1, 0xa3, 0x08, 0xc7, 0x95, 0xef,
	0xb1, 0xbe, 0x36, 0x52, 0x9d, 0xa5, 0x93, 0x5d, 0xf0, 0x90, 0x3c, 0x02, 0x67, 0xf2, 0xb9, 0xb3,
	0xa8, 0x36, 0xc2, 0x21, 0xac, 0x57, 0x1a, 0xe8, 0xca, 0x10, 0x6d, 0x57, 0x88, 0xd2, 0xec, 0x8a,
	0x30, 0x8f, 0x61, 0x31, 0xbf, 0x50, 0x36, 0x1c, 0xe8, 0x73, 0x69, 0x5c, 0xf5, 0x37, 0xa9, 0xfa,
	0xcd, 0x12, 0xa6, 0xc2, 0x1f, 0xc0, 0x03, 0x1a, 0x9d, 0xea, 0x07, 0x59, 0xd6, 0x9a, 0x9f, 0x8b,
	0x78, 0x58, 0x4c, 0xe6, 0xff, 0xd0, 0x72, 0x6b, 0x85, 0xd5, 0xb8, 0xd8, 0xdf, 0xa0, 0x8d, 0x1b,
	0xd1, 0xc7, 0xb0, 0xf9, 0x8f, 0x41, 0x70, 0xe8, 0x72, 0x37, 0x75, 0x1f, 0x51, 0xa4, 0xf5, 0xab,
	0x23, 0xe1, 0x10, 0x7e, 0x4d, 0x43, 0x18, 0x67, 0x01, 0xd5, 0x04, 0xef, 0xf4, 0xe0, 0xbf, 0x7e,
	0x16, 0x20, 0x72, 0xa2, 0x23, 0x89, 0x6d, 0x60, 0x64, 0x88, 0x1b, 0x18, 0x79, 0x1d, 0x3c, 0xa6,
	0xb8, 0x8d, 0x02, 0x75, 0x1a, 0xb8, 0x07, 0xf3, 0xb9, 0x4a, 0xcf, 0x62, 0x69, 0x75, 0x1a, 0x6c,
	0x91, 0x45, 0x09, 0xe0, 0x0b, 0x54, 0xda, 0x81, 0x44, 0x35, 0x85, 0x74, 0xdb, 0xe2, 0x0b, 0xd4,
	0xff, 0x0e, 0x6a, 0x1b, 0xb3, 0x5b, 0x8d, 0x6e, 0xf1, 0xd9, 0xfe, 0xa3, 0x06, 0x8b, 0xbd, 0x84,
	0xe7, 0xa3, 0xdc, 0xca, 0x84, 0x2b, 0xec, 0x3e, 0x7a, 0x16, 0xba, 0xef, 0xcc, 0x28, 0x4d, 0x0f,
	0x09, 0x7c, 0x16, 0x12, 0xf4, 0x95, 0x51, 0x9a, 0xfd, 0x5c, 0x83, 0x45, 0x6f, 0x51, 0xdc, 0xe8,
	0xef, 0xe8, 0x0d, 0x31, 0xb9, 0xaf, 0x6e, 0xc3, 0x65, 0x71, 0xe2, 0x5e, 0x12, 0xed, 0x3f, 0x6f,
	0xc0, 0xd2, 0xf4, 0xfb, 0x9d, 0xfd, 0x56, 0x03, 0x9c, 0xfe, 0xda, 0x54, 0xd1, 0x60, 0xe7, 0x5d,
	0x3d, 0x79, 0x2a, 0x8b, 0x76, 0x17, 0x29, 0x8f, 0x7d, 0x42, 0x9e, 0x61, 0x6e, 0xbf, 0xd6, 0xc6,
	0xff, 0x2b, 0xc6, 0x15, 0x08, 0x76, 0x29, 0xb5, 0xf4, 0xfa, 0x53, 0xab, 0x16, 0xbd, 0xdb, 0xf4,
	0x97, 0x29, 0x61, 0x47, 0x11, 0xbe, 0x2b, 0xdd, 0xa9, 0xe1, 0x7a, 0xb4, 0x82, 0x88, 0xf9, 0x1b,
	0x39, 0x0a, 0x9e, 0xb8, 0x17, 0x13, 0x91, 0x2f, 0x4b, 0xee, 0xb9, 0x1c, 0xe1, 0x05, 0x99, 0x09,
	0x63, 0x53, 0x9c, 0xf6, 0xa5, 0x80, 0x8c, 0xb2, 0xa3, 0xe0, 0x23, 0x77, 0x41, 0x7a, 0xfa, 0x74,
	0xac, 0x25, 0x24, 0xd9, 0xa7, 0x70, 0x77, 0xca, 0xaf, 0xfa, 0x3a, 0xdd, 0xa3, 0x26, 0xb9, 0x33,
	0xe1, 0x7a, 0x52, 0xbe, 0x1d, 0x2b, 0x8b, 0x4e, 0xa7, 0xfa, 0xf1, 0xc4, 0xa2, 0x53, 0xc9, 0x6e,
	0xc3, 0xb2, 0xc8, 0xf0, 0xff, 0x48, 0xc4, 0xcb, 0x86, 0xfc, 0xc4, 0xdd, 0x4e, 0x9e, 0x38, 0xf6,
	0x6d, 0xd9, 0xbb, 0x49, 0xff, 0x3e, 0x77, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xb3, 0x25,
	0xce, 0x97, 0x0e, 0x00, 0x00,
}