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
// source: bm_bundle_data.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_bundle_bundle_bundles_bundle_bundle_bundle_bundle_item

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

type BmBundleData_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmBundleData_KEYS) Reset()         { *m = BmBundleData_KEYS{} }
func (m *BmBundleData_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmBundleData_KEYS) ProtoMessage()    {}
func (*BmBundleData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeef7a164208cb2, []int{0}
}

func (m *BmBundleData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmBundleData_KEYS.Unmarshal(m, b)
}
func (m *BmBundleData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmBundleData_KEYS.Marshal(b, m, deterministic)
}
func (m *BmBundleData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmBundleData_KEYS.Merge(m, src)
}
func (m *BmBundleData_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmBundleData_KEYS.Size(m)
}
func (m *BmBundleData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmBundleData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmBundleData_KEYS proto.InternalMessageInfo

func (m *BmBundleData_KEYS) GetBundleInterface() string {
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
	return fileDescriptor_faeef7a164208cb2, []int{1}
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
	return fileDescriptor_faeef7a164208cb2, []int{2}
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
	return fileDescriptor_faeef7a164208cb2, []int{3}
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
	BundleInterfaceName             string             `protobuf:"bytes,50,opt,name=bundle_interface_name,json=bundleInterfaceName,proto3" json:"bundle_interface_name,omitempty"`
	MacAddress                      *BmMacAddrSt       `protobuf:"bytes,51,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	AvailableBandwidth              uint32             `protobuf:"varint,52,opt,name=available_bandwidth,json=availableBandwidth,proto3" json:"available_bandwidth,omitempty"`
	EffectiveBandwidth              uint32             `protobuf:"varint,53,opt,name=effective_bandwidth,json=effectiveBandwidth,proto3" json:"effective_bandwidth,omitempty"`
	ConfiguredBandwidth             uint32             `protobuf:"varint,54,opt,name=configured_bandwidth,json=configuredBandwidth,proto3" json:"configured_bandwidth,omitempty"`
	MinimumActiveLinks              uint32             `protobuf:"varint,55,opt,name=minimum_active_links,json=minimumActiveLinks,proto3" json:"minimum_active_links,omitempty"`
	MaximumActiveLinks              uint32             `protobuf:"varint,56,opt,name=maximum_active_links,json=maximumActiveLinks,proto3" json:"maximum_active_links,omitempty"`
	MaximumActiveLinksSource        string             `protobuf:"bytes,57,opt,name=maximum_active_links_source,json=maximumActiveLinksSource,proto3" json:"maximum_active_links_source,omitempty"`
	MinimumBandwidth                uint32             `protobuf:"varint,58,opt,name=minimum_bandwidth,json=minimumBandwidth,proto3" json:"minimum_bandwidth,omitempty"`
	PrimaryMember                   string             `protobuf:"bytes,59,opt,name=primary_member,json=primaryMember,proto3" json:"primary_member,omitempty"`
	BfdConfig                       []*BmBundleBfdData `protobuf:"bytes,60,rep,name=bfd_config,json=bfdConfig,proto3" json:"bfd_config,omitempty"`
	BundleStatus                    string             `protobuf:"bytes,61,opt,name=bundle_status,json=bundleStatus,proto3" json:"bundle_status,omitempty"`
	ActiveMemberCount               uint32             `protobuf:"varint,62,opt,name=active_member_count,json=activeMemberCount,proto3" json:"active_member_count,omitempty"`
	StandbyMemberCount              uint32             `protobuf:"varint,63,opt,name=standby_member_count,json=standbyMemberCount,proto3" json:"standby_member_count,omitempty"`
	ConfiguredMemberCount           uint32             `protobuf:"varint,64,opt,name=configured_member_count,json=configuredMemberCount,proto3" json:"configured_member_count,omitempty"`
	MacSource                       string             `protobuf:"bytes,65,opt,name=mac_source,json=macSource,proto3" json:"mac_source,omitempty"`
	MacSourceMember                 string             `protobuf:"bytes,66,opt,name=mac_source_member,json=macSourceMember,proto3" json:"mac_source_member,omitempty"`
	InterChassis                    uint32             `protobuf:"varint,67,opt,name=inter_chassis,json=interChassis,proto3" json:"inter_chassis,omitempty"`
	IsActive                        uint32             `protobuf:"varint,68,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	LacpStatus                      string             `protobuf:"bytes,69,opt,name=lacp_status,json=lacpStatus,proto3" json:"lacp_status,omitempty"`
	MlacpStatus                     string             `protobuf:"bytes,70,opt,name=mlacp_status,json=mlacpStatus,proto3" json:"mlacp_status,omitempty"`
	Ipv4BfdStatus                   string             `protobuf:"bytes,71,opt,name=ipv4bfd_status,json=ipv4bfdStatus,proto3" json:"ipv4bfd_status,omitempty"`
	LinkOrderStatus                 string             `protobuf:"bytes,72,opt,name=link_order_status,json=linkOrderStatus,proto3" json:"link_order_status,omitempty"`
	Ipv6BfdStatus                   string             `protobuf:"bytes,73,opt,name=ipv6bfd_status,json=ipv6bfdStatus,proto3" json:"ipv6bfd_status,omitempty"`
	LoadBalanceHashType             string             `protobuf:"bytes,74,opt,name=load_balance_hash_type,json=loadBalanceHashType,proto3" json:"load_balance_hash_type,omitempty"`
	LoadBalanceLocalityThreshold    uint32             `protobuf:"varint,75,opt,name=load_balance_locality_threshold,json=loadBalanceLocalityThreshold,proto3" json:"load_balance_locality_threshold,omitempty"`
	SuppressionTimer                uint32             `protobuf:"varint,76,opt,name=suppression_timer,json=suppressionTimer,proto3" json:"suppression_timer,omitempty"`
	WaitWhileTimer                  uint32             `protobuf:"varint,77,opt,name=wait_while_timer,json=waitWhileTimer,proto3" json:"wait_while_timer,omitempty"`
	CollectorMaxDelay               uint32             `protobuf:"varint,78,opt,name=collector_max_delay,json=collectorMaxDelay,proto3" json:"collector_max_delay,omitempty"`
	CiscoExtensions                 uint32             `protobuf:"varint,79,opt,name=cisco_extensions,json=ciscoExtensions,proto3" json:"cisco_extensions,omitempty"`
	LacpNonrevertive                uint32             `protobuf:"varint,80,opt,name=lacp_nonrevertive,json=lacpNonrevertive,proto3" json:"lacp_nonrevertive,omitempty"`
	IccpGroupId                     uint32             `protobuf:"varint,81,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	ActiveForeignMemberCount        uint32             `protobuf:"varint,82,opt,name=active_foreign_member_count,json=activeForeignMemberCount,proto3" json:"active_foreign_member_count,omitempty"`
	ConfiguredForeignMemberCount    uint32             `protobuf:"varint,83,opt,name=configured_foreign_member_count,json=configuredForeignMemberCount,proto3" json:"configured_foreign_member_count,omitempty"`
	SwitchoverType                  string             `protobuf:"bytes,84,opt,name=switchover_type,json=switchoverType,proto3" json:"switchover_type,omitempty"`
	MaximizeThresholdValueLinks     uint32             `protobuf:"varint,85,opt,name=maximize_threshold_value_links,json=maximizeThresholdValueLinks,proto3" json:"maximize_threshold_value_links,omitempty"`
	MaximizeThresholdValueBandWidth uint32             `protobuf:"varint,86,opt,name=maximize_threshold_value_band_width,json=maximizeThresholdValueBandWidth,proto3" json:"maximize_threshold_value_band_width,omitempty"`
	MlacpMode                       string             `protobuf:"bytes,87,opt,name=mlacp_mode,json=mlacpMode,proto3" json:"mlacp_mode,omitempty"`
	RecoveryDelay                   uint32             `protobuf:"varint,88,opt,name=recovery_delay,json=recoveryDelay,proto3" json:"recovery_delay,omitempty"`
	Singleton                       uint32             `protobuf:"varint,89,opt,name=singleton,proto3" json:"singleton,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}           `json:"-"`
	XXX_unrecognized                []byte             `json:"-"`
	XXX_sizecache                   int32              `json:"-"`
}

func (m *BmBundleData) Reset()         { *m = BmBundleData{} }
func (m *BmBundleData) String() string { return proto.CompactTextString(m) }
func (*BmBundleData) ProtoMessage()    {}
func (*BmBundleData) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeef7a164208cb2, []int{4}
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

func init() {
	proto.RegisterType((*BmBundleData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bundle.bundle_bundles.bundle_bundle.bundle_bundle_item.bm_bundle_data_KEYS")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bundle.bundle_bundles.bundle_bundle.bundle_bundle_item.bm_mac_addr_st")
	proto.RegisterType((*BmAddr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bundle.bundle_bundles.bundle_bundle.bundle_bundle_item.bm_addr")
	proto.RegisterType((*BmBundleBfdData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bundle.bundle_bundles.bundle_bundle.bundle_bundle_item.bm_bundle_bfd_data")
	proto.RegisterType((*BmBundleData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.bundle.bundle_bundles.bundle_bundle.bundle_bundle_item.bm_bundle_data")
}

func init() { proto.RegisterFile("bm_bundle_data.proto", fileDescriptor_faeef7a164208cb2) }

var fileDescriptor_faeef7a164208cb2 = []byte{
	// 1182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x72, 0x13, 0xb7,
	0x17, 0x9f, 0x10, 0xfe, 0x40, 0x8e, 0x49, 0x48, 0xd6, 0xc0, 0x5f, 0x33, 0xd0, 0x86, 0x86, 0x61,
	0x48, 0xd3, 0x8e, 0xdb, 0x26, 0xd4, 0xfd, 0xa4, 0x25, 0x04, 0x03, 0x29, 0x31, 0xb4, 0x4e, 0x80,
	0x72, 0xa5, 0xd1, 0xee, 0x6a, 0x63, 0x4d, 0x77, 0xa5, 0x1d, 0x49, 0xeb, 0xc4, 0x7d, 0x90, 0x4e,
	0xef, 0xfa, 0x10, 0x7d, 0xb3, 0x3e, 0x41, 0x47, 0x47, 0x5a, 0xef, 0x1a, 0xd2, 0x5b, 0xae, 0x92,
	0xfd, 0x7d, 0x48, 0x3a, 0x3a, 0x1f, 0x32, 0x5c, 0x8d, 0x0b, 0x1a, 0x57, 0x32, 0xcd, 0x39, 0x4d,
	0x99, 0x65, 0xbd, 0x52, 0x2b, 0xab, 0x22, 0x9e, 0x08, 0x93, 0x28, 0x2a, 0x94, 0xa1, 0xa7, 0x3a,
	0xf0, 0xc5, 0xb1, 0xa6, 0xaa, 0xe4, 0xba, 0x17, 0xe4, 0x42, 0x66, 0x4a, 0x17, 0xcc, 0x0a, 0x25,
	0x03, 0x54, 0x33, 0xfe, 0x8f, 0x99, 0xff, 0x9c, 0xff, 0xa2, 0xc2, 0xf2, 0x62, 0xe3, 0x01, 0x74,
	0xe7, 0xb7, 0xa7, 0xcf, 0x06, 0x6f, 0x0e, 0xa3, 0x8f, 0x61, 0x75, 0xb6, 0x87, 0xe5, 0x3a, 0x63,
	0x09, 0x27, 0x0b, 0xb7, 0x16, 0x36, 0x97, 0x46, 0x57, 0x3c, 0xbe, 0x5f, 0xc3, 0x1b, 0x5b, 0xb0,
	0x12, 0x17, 0xb4, 0x60, 0x09, 0x65, 0x69, 0xaa, 0xa9, 0xb1, 0x11, 0x81, 0x8b, 0xee, 0x5f, 0x6e,
	0x4c, 0xf0, 0xd4, 0x9f, 0x1b, 0xbb, 0x70, 0x31, 0x2e, 0x50, 0x17, 0xad, 0xc0, 0x39, 0x96, 0x05,
	0xfe, 0x1c, 0xcb, 0xa2, 0x08, 0xce, 0x8b, 0x72, 0x72, 0x8f, 0x9c, 0x43, 0x04, 0xff, 0x0f, 0x58,
	0x9f, 0x2c, 0xce, 0xb0, 0xfe, 0xc6, 0x3f, 0x8b, 0x10, 0x35, 0x27, 0x8e, 0xb3, 0x14, 0x4f, 0x1d,
	0xdd, 0x86, 0xe5, 0x00, 0x19, 0xcb, 0x6c, 0x55, 0xef, 0x7c, 0xd9, 0x83, 0x87, 0x88, 0x45, 0xeb,
	0xd0, 0x31, 0x96, 0x69, 0x4b, 0xad, 0x28, 0xb8, 0xc6, 0xad, 0x96, 0x47, 0x80, 0xd0, 0x91, 0x43,
	0xa2, 0x4f, 0x21, 0x92, 0xb1, 0xa6, 0x95, 0x4c, 0x94, 0xcc, 0xc4, 0x71, 0xd0, 0x2d, 0xa2, 0x6e,
	0x55, 0xc6, 0xfa, 0x65, 0x20, 0xbc, 0xfa, 0x2e, 0x5c, 0x29, 0x35, 0xcf, 0x68, 0x51, 0xe5, 0x56,
	0x94, 0xb9, 0xe0, 0x9a, 0x9c, 0x47, 0xe9, 0x8a, 0x83, 0x87, 0x33, 0x34, 0xda, 0x82, 0x35, 0x2f,
	0x14, 0xd2, 0xdf, 0xe7, 0x84, 0xe5, 0xe4, 0x7f, 0x28, 0xc5, 0x15, 0x86, 0x42, 0xee, 0x07, 0x38,
	0xda, 0x81, 0xeb, 0xa8, 0xe5, 0xc9, 0x58, 0xcd, 0x1b, 0x2e, 0xa0, 0xa1, 0xeb, 0xd8, 0x41, 0x32,
	0x56, 0x6d, 0xd3, 0x3a, 0x74, 0x32, 0x66, 0x2c, 0x4d, 0xb9, 0xe5, 0x89, 0x25, 0x17, 0x7d, 0x60,
	0x0e, 0x7a, 0x84, 0x48, 0xf4, 0xd7, 0x02, 0x74, 0x53, 0x6e, 0xac, 0x90, 0x58, 0x2d, 0xb4, 0xce,
	0xcf, 0xa5, 0x5b, 0x0b, 0x9b, 0x9d, 0x6d, 0xd9, 0x7b, 0x2f, 0xc5, 0xd6, 0x0b, 0xb9, 0x1f, 0x45,
	0xad, 0xa3, 0xec, 0xfa, 0x93, 0x44, 0x37, 0x60, 0xa9, 0x50, 0xa9, 0xdf, 0x86, 0x2c, 0x61, 0x00,
	0x97, 0x1c, 0xb0, 0x2f, 0x33, 0xb5, 0xf1, 0xf7, 0x2a, 0x16, 0x59, 0xab, 0x4c, 0xa3, 0x6d, 0xb8,
	0xf6, 0x76, 0x85, 0x52, 0xc9, 0x0a, 0x4e, 0xb6, 0x31, 0xf1, 0xdd, 0xb7, 0xca, 0xf4, 0x39, 0x2b,
	0x78, 0xf4, 0xc7, 0x02, 0x74, 0xea, 0x42, 0x75, 0xd1, 0xef, 0x60, 0xf4, 0xd5, 0xfb, 0x8b, 0xbe,
	0xd5, 0x25, 0x23, 0x28, 0x58, 0x52, 0x07, 0xff, 0x19, 0x74, 0xd9, 0x84, 0x89, 0x9c, 0xc5, 0xce,
	0xc3, 0x64, 0x7a, 0x22, 0x52, 0x3b, 0x26, 0xf7, 0xf0, 0x1a, 0xa2, 0x19, 0xf5, 0xb0, 0x66, 0x9c,
	0x81, 0x67, 0x19, 0x4f, 0xac, 0x98, 0xb4, 0x0d, 0x5f, 0x7a, 0xc3, 0x8c, 0x6a, 0x0c, 0x5f, 0xc0,
	0x55, 0x5f, 0xba, 0x95, 0xe6, 0x69, 0xcb, 0xd1, 0xf7, 0x45, 0xd5, 0x70, 0x8d, 0xe5, 0x73, 0xb8,
	0x5a, 0x08, 0x29, 0x8a, 0xaa, 0xa0, 0xcc, 0x6f, 0x94, 0x0b, 0xf9, 0x9b, 0x21, 0x5f, 0xf9, 0x4d,
	0x02, 0xb7, 0x8b, 0xd4, 0x81, 0x63, 0xd0, 0xc1, 0x4e, 0xdf, 0x75, 0x7c, 0x1d, 0x1c, 0x9e, 0x6b,
	0x3b, 0xee, 0xc3, 0x8d, 0xb3, 0x1c, 0xd4, 0xa8, 0x4a, 0x27, 0x9c, 0x7c, 0x83, 0xb9, 0x24, 0xef,
	0x1a, 0x0f, 0x91, 0x8f, 0x3e, 0x81, 0xb5, 0xfa, 0x88, 0x4d, 0x48, 0xdf, 0xfa, 0x76, 0x0d, 0x44,
	0x13, 0xcf, 0x1d, 0x58, 0x29, 0xb5, 0x28, 0x98, 0x9e, 0xd2, 0x82, 0x17, 0x31, 0xd7, 0xe4, 0x3b,
	0x5c, 0x7e, 0x39, 0xa0, 0x43, 0x04, 0xa3, 0x3f, 0x17, 0x00, 0xdc, 0x58, 0xf1, 0x57, 0x42, 0xbe,
	0xbf, 0xb5, 0xb8, 0xd9, 0xd9, 0x9e, 0xbe, 0xbf, 0x1a, 0x79, 0x6b, 0xb2, 0x8d, 0x96, 0xe2, 0x2c,
	0xdd, 0xc3, 0xb3, 0xbc, 0x3b, 0xe4, 0xee, 0x9f, 0x31, 0xe4, 0x7a, 0xd0, 0x0d, 0x57, 0xe9, 0xa3,
	0xa4, 0x89, 0xaa, 0xa4, 0x25, 0x3f, 0xe0, 0xad, 0xac, 0x79, 0xca, 0x87, 0xba, 0xe7, 0x08, 0x97,
	0x34, 0x63, 0x99, 0x4c, 0xe3, 0xe9, 0xbc, 0xe1, 0x47, 0x9f, 0xb4, 0xc0, 0xb5, 0x1d, 0x7d, 0xf8,
	0x7f, 0xab, 0x96, 0xe6, 0x4c, 0x0f, 0xd0, 0x74, 0xad, 0xa1, 0xdb, 0xbe, 0x0f, 0xc0, 0xd5, 0x7c,
	0x9d, 0xdb, 0x5d, 0x3c, 0xfb, 0x52, 0xc1, 0x92, 0x90, 0xcc, 0x2d, 0x58, 0x6b, 0xe8, 0x3a, 0x45,
	0x0f, 0xfd, 0xa3, 0x33, 0x53, 0x85, 0x24, 0xdd, 0x86, 0x65, 0x6c, 0x7b, 0x9a, 0x8c, 0x99, 0x31,
	0xc2, 0x90, 0x3d, 0xdc, 0xf8, 0x32, 0x82, 0x7b, 0x1e, 0x73, 0x23, 0x45, 0x98, 0x50, 0x57, 0xe4,
	0x91, 0x1f, 0x29, 0xc2, 0xf8, 0x2a, 0x72, 0x23, 0x33, 0x67, 0x49, 0x59, 0xdf, 0xe4, 0x00, 0xf7,
	0x01, 0x07, 0x85, 0x7b, 0xfc, 0x08, 0x2e, 0x17, 0x6d, 0xc5, 0x63, 0x54, 0x74, 0x8a, 0x96, 0xe4,
	0x0e, 0xac, 0xb8, 0x77, 0xca, 0xa5, 0x2a, 0x88, 0x9e, 0xf8, 0x8a, 0x0a, 0x68, 0x90, 0x6d, 0xc1,
	0x9a, 0xab, 0x6a, 0xaa, 0x74, 0xca, 0x75, 0xad, 0x7c, 0xea, 0x03, 0x73, 0xc4, 0x0b, 0x87, 0xcf,
	0x2d, 0xd9, 0x6f, 0x2d, 0xb9, 0x3f, 0x5b, 0xb2, 0xdf, 0x2c, 0xb9, 0x03, 0xd7, 0x73, 0xc5, 0x5c,
	0x23, 0xe7, 0x4c, 0x26, 0x9c, 0x8e, 0x99, 0x19, 0x53, 0x3b, 0x2d, 0x39, 0xf9, 0xc9, 0x8f, 0x3f,
	0xc7, 0x3e, 0xf4, 0xe4, 0x53, 0x66, 0xc6, 0x47, 0xd3, 0x92, 0x47, 0x03, 0x58, 0x9f, 0x33, 0xe5,
	0x2a, 0x61, 0xb9, 0xb0, 0x53, 0x6a, 0xc7, 0x9a, 0x9b, 0xb1, 0xca, 0x53, 0xf2, 0x0c, 0x6f, 0xe9,
	0x66, 0xcb, 0x7d, 0x10, 0x44, 0x47, 0xb5, 0xc6, 0x35, 0x9d, 0xa9, 0xca, 0xd2, 0x0d, 0x2e, 0xf7,
	0x94, 0xf8, 0x37, 0xf2, 0xc0, 0x37, 0x5d, 0x8b, 0xf0, 0x6f, 0xe4, 0x26, 0xac, 0x9e, 0x30, 0x61,
	0xe9, 0xc9, 0x58, 0xe4, 0x3c, 0x68, 0x87, 0xfe, 0x91, 0x74, 0xf8, 0x6b, 0x07, 0x7b, 0x65, 0x0f,
	0xba, 0x89, 0xca, 0x73, 0x9e, 0x58, 0xa5, 0x69, 0xc1, 0x4e, 0x69, 0xca, 0x73, 0x36, 0x25, 0xcf,
	0x7d, 0xdd, 0xce, 0xa8, 0x21, 0x3b, 0x7d, 0xe4, 0x08, 0xf7, 0x13, 0xc5, 0xf7, 0x24, 0x3f, 0xb5,
	0x5c, 0xba, 0x1d, 0x0d, 0x79, 0xe1, 0xdf, 0x54, 0xc4, 0x07, 0x33, 0xd8, 0x9d, 0x18, 0x33, 0x29,
	0x95, 0xd4, 0x7c, 0xc2, 0x35, 0x16, 0xc4, 0xcf, 0xfe, 0xc4, 0x8e, 0x78, 0xde, 0xc2, 0xa3, 0x0d,
	0x58, 0x16, 0x49, 0x52, 0xd2, 0x63, 0xad, 0xaa, 0x92, 0x8a, 0x94, 0xfc, 0x82, 0xc2, 0x8e, 0x03,
	0x9f, 0x38, 0x6c, 0x3f, 0x75, 0x63, 0x2b, 0xf4, 0x58, 0xa6, 0x34, 0x17, 0xc7, 0x72, 0xbe, 0x0b,
	0x46, 0xe8, 0x20, 0x5e, 0xf2, 0xd8, 0x2b, 0xda, 0x8d, 0x30, 0x80, 0xf5, 0x56, 0x03, 0x9d, 0xb9,
	0xc4, 0xa1, 0x4f, 0x44, 0x23, 0x3b, 0x63, 0x99, 0xbb, 0x70, 0xc5, 0x9c, 0x08, 0x9b, 0x8c, 0xd5,
	0x84, 0x6b, 0x9f, 0xfd, 0x23, 0xcc, 0xfe, 0x4a, 0x03, 0x63, 0xe2, 0xf7, 0xe0, 0x43, 0x1c, 0xa1,
	0xe2, 0x77, 0xde, 0xe4, 0x9a, 0x4e, 0x58, 0x5e, 0xd5, 0x13, 0xfa, 0x25, 0x6e, 0x77, 0xa3, 0x56,
	0xcd, 0x92, 0xfd, 0xca, 0x69, 0xfc, 0xa8, 0x3e, 0x80, 0xdb, 0xff, 0xb9, 0x88, 0x1b, 0xbe, 0xd4,
	0x4f, 0xdf, 0x57, 0xb8, 0xd2, 0xfa, 0xd9, 0x2b, 0xb9, 0x61, 0xfc, 0x1a, 0x87, 0xb1, 0x9b, 0x05,
	0x98, 0x13, 0xf7, 0xc6, 0x93, 0xd7, 0x61, 0x16, 0x38, 0x64, 0xa8, 0x52, 0xee, 0xda, 0x40, 0xf3,
	0xc4, 0x05, 0x30, 0x0d, 0x75, 0xf0, 0x2b, 0xae, 0xbb, 0x5c, 0xa3, 0xbe, 0x06, 0x6e, 0xc2, 0x92,
	0x11, 0xf2, 0x38, 0xe7, 0x56, 0x49, 0xf2, 0x06, 0x15, 0x0d, 0x10, 0x5f, 0xc0, 0x5f, 0xd2, 0x3b,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x55, 0x4a, 0xaf, 0x7a, 0x61, 0x0b, 0x00, 0x00,
}
