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
// source: elod_running_config.proto

package cisco_ios_xr_ethernet_link_oam_oper_ether_link_oam_running_config_interfaces_running_config_interface

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

type ElodRunningConfig_KEYS struct {
	MemberInterface      string   `protobuf:"bytes,1,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElodRunningConfig_KEYS) Reset()         { *m = ElodRunningConfig_KEYS{} }
func (m *ElodRunningConfig_KEYS) String() string { return proto.CompactTextString(m) }
func (*ElodRunningConfig_KEYS) ProtoMessage()    {}
func (*ElodRunningConfig_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0612a046b6621eb, []int{0}
}

func (m *ElodRunningConfig_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElodRunningConfig_KEYS.Unmarshal(m, b)
}
func (m *ElodRunningConfig_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElodRunningConfig_KEYS.Marshal(b, m, deterministic)
}
func (m *ElodRunningConfig_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElodRunningConfig_KEYS.Merge(m, src)
}
func (m *ElodRunningConfig_KEYS) XXX_Size() int {
	return xxx_messageInfo_ElodRunningConfig_KEYS.Size(m)
}
func (m *ElodRunningConfig_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ElodRunningConfig_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ElodRunningConfig_KEYS proto.InternalMessageInfo

func (m *ElodRunningConfig_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
}

type ElodRunningConfig struct {
	FastHelloIntervalEnabled             bool     `protobuf:"varint,50,opt,name=fast_hello_interval_enabled,json=fastHelloIntervalEnabled,proto3" json:"fast_hello_interval_enabled,omitempty"`
	LinkMonitorEnabled                   bool     `protobuf:"varint,51,opt,name=link_monitor_enabled,json=linkMonitorEnabled,proto3" json:"link_monitor_enabled,omitempty"`
	RemoteLoopbackEnabled                bool     `protobuf:"varint,52,opt,name=remote_loopback_enabled,json=remoteLoopbackEnabled,proto3" json:"remote_loopback_enabled,omitempty"`
	MibRetrievalEnabled                  bool     `protobuf:"varint,53,opt,name=mib_retrieval_enabled,json=mibRetrievalEnabled,proto3" json:"mib_retrieval_enabled,omitempty"`
	UdlfEnabled                          bool     `protobuf:"varint,54,opt,name=udlf_enabled,json=udlfEnabled,proto3" json:"udlf_enabled,omitempty"`
	Mode                                 string   `protobuf:"bytes,55,opt,name=mode,proto3" json:"mode,omitempty"`
	ConnectionTimeout                    uint32   `protobuf:"varint,56,opt,name=connection_timeout,json=connectionTimeout,proto3" json:"connection_timeout,omitempty"`
	SymbolPeriodWindow                   uint32   `protobuf:"varint,57,opt,name=symbol_period_window,json=symbolPeriodWindow,proto3" json:"symbol_period_window,omitempty"`
	SymbolPeriodWindowUnits              uint32   `protobuf:"varint,58,opt,name=symbol_period_window_units,json=symbolPeriodWindowUnits,proto3" json:"symbol_period_window_units,omitempty"`
	SymbolPeriodWindowMultiplier         uint32   `protobuf:"varint,59,opt,name=symbol_period_window_multiplier,json=symbolPeriodWindowMultiplier,proto3" json:"symbol_period_window_multiplier,omitempty"`
	SymbolPeriodThresholdLow             uint64   `protobuf:"varint,60,opt,name=symbol_period_threshold_low,json=symbolPeriodThresholdLow,proto3" json:"symbol_period_threshold_low,omitempty"`
	SymbolPeriodThresholdHigh            uint64   `protobuf:"varint,61,opt,name=symbol_period_threshold_high,json=symbolPeriodThresholdHigh,proto3" json:"symbol_period_threshold_high,omitempty"`
	SymbolPeriodThresholdUnits           uint32   `protobuf:"varint,62,opt,name=symbol_period_threshold_units,json=symbolPeriodThresholdUnits,proto3" json:"symbol_period_threshold_units,omitempty"`
	SymbolPeriodThresholdLowMultiplier   uint32   `protobuf:"varint,63,opt,name=symbol_period_threshold_low_multiplier,json=symbolPeriodThresholdLowMultiplier,proto3" json:"symbol_period_threshold_low_multiplier,omitempty"`
	SymbolPeriodThresholdHighMultiplier  uint32   `protobuf:"varint,64,opt,name=symbol_period_threshold_high_multiplier,json=symbolPeriodThresholdHighMultiplier,proto3" json:"symbol_period_threshold_high_multiplier,omitempty"`
	FrameWindow                          uint32   `protobuf:"varint,65,opt,name=frame_window,json=frameWindow,proto3" json:"frame_window,omitempty"`
	FrameThresholdLow                    uint64   `protobuf:"varint,66,opt,name=frame_threshold_low,json=frameThresholdLow,proto3" json:"frame_threshold_low,omitempty"`
	FrameThresholdHigh                   uint64   `protobuf:"varint,67,opt,name=frame_threshold_high,json=frameThresholdHigh,proto3" json:"frame_threshold_high,omitempty"`
	FrameThresholdLowMultiplier          uint32   `protobuf:"varint,68,opt,name=frame_threshold_low_multiplier,json=frameThresholdLowMultiplier,proto3" json:"frame_threshold_low_multiplier,omitempty"`
	FrameThresholdHighMultiplier         uint32   `protobuf:"varint,69,opt,name=frame_threshold_high_multiplier,json=frameThresholdHighMultiplier,proto3" json:"frame_threshold_high_multiplier,omitempty"`
	FramePeriodWindow                    uint32   `protobuf:"varint,70,opt,name=frame_period_window,json=framePeriodWindow,proto3" json:"frame_period_window,omitempty"`
	FramePeriodWindowUnits               uint32   `protobuf:"varint,71,opt,name=frame_period_window_units,json=framePeriodWindowUnits,proto3" json:"frame_period_window_units,omitempty"`
	FramePeriodWindowMultiplier          uint32   `protobuf:"varint,72,opt,name=frame_period_window_multiplier,json=framePeriodWindowMultiplier,proto3" json:"frame_period_window_multiplier,omitempty"`
	FramePeriodThresholdLow              uint64   `protobuf:"varint,73,opt,name=frame_period_threshold_low,json=framePeriodThresholdLow,proto3" json:"frame_period_threshold_low,omitempty"`
	FramePeriodThresholdHigh             uint64   `protobuf:"varint,74,opt,name=frame_period_threshold_high,json=framePeriodThresholdHigh,proto3" json:"frame_period_threshold_high,omitempty"`
	FramePeriodThresholdUnits            uint32   `protobuf:"varint,75,opt,name=frame_period_threshold_units,json=framePeriodThresholdUnits,proto3" json:"frame_period_threshold_units,omitempty"`
	FramePeriodThresholdLowMultiplier    uint32   `protobuf:"varint,76,opt,name=frame_period_threshold_low_multiplier,json=framePeriodThresholdLowMultiplier,proto3" json:"frame_period_threshold_low_multiplier,omitempty"`
	FramePeriodThresholdHighMultiplier   uint32   `protobuf:"varint,77,opt,name=frame_period_threshold_high_multiplier,json=framePeriodThresholdHighMultiplier,proto3" json:"frame_period_threshold_high_multiplier,omitempty"`
	FrameSecondsWindow                   uint32   `protobuf:"varint,78,opt,name=frame_seconds_window,json=frameSecondsWindow,proto3" json:"frame_seconds_window,omitempty"`
	FrameSecondsThresholdLow             uint64   `protobuf:"varint,79,opt,name=frame_seconds_threshold_low,json=frameSecondsThresholdLow,proto3" json:"frame_seconds_threshold_low,omitempty"`
	FrameSecondsThresholdHigh            uint64   `protobuf:"varint,80,opt,name=frame_seconds_threshold_high,json=frameSecondsThresholdHigh,proto3" json:"frame_seconds_threshold_high,omitempty"`
	HighThresholdAction                  string   `protobuf:"bytes,81,opt,name=high_threshold_action,json=highThresholdAction,proto3" json:"high_threshold_action,omitempty"`
	LinkFaultAction                      string   `protobuf:"bytes,82,opt,name=link_fault_action,json=linkFaultAction,proto3" json:"link_fault_action,omitempty"`
	DyingGaspAction                      string   `protobuf:"bytes,83,opt,name=dying_gasp_action,json=dyingGaspAction,proto3" json:"dying_gasp_action,omitempty"`
	CriticalEventAction                  string   `protobuf:"bytes,84,opt,name=critical_event_action,json=criticalEventAction,proto3" json:"critical_event_action,omitempty"`
	DiscoveryTimeoutAction               string   `protobuf:"bytes,85,opt,name=discovery_timeout_action,json=discoveryTimeoutAction,proto3" json:"discovery_timeout_action,omitempty"`
	CapabilitiesConflictAction           string   `protobuf:"bytes,86,opt,name=capabilities_conflict_action,json=capabilitiesConflictAction,proto3" json:"capabilities_conflict_action,omitempty"`
	WiringConflictAction                 string   `protobuf:"bytes,87,opt,name=wiring_conflict_action,json=wiringConflictAction,proto3" json:"wiring_conflict_action,omitempty"`
	SessionUpAction                      string   `protobuf:"bytes,88,opt,name=session_up_action,json=sessionUpAction,proto3" json:"session_up_action,omitempty"`
	SessionDownAction                    string   `protobuf:"bytes,89,opt,name=session_down_action,json=sessionDownAction,proto3" json:"session_down_action,omitempty"`
	RemoteLoopbackAction                 string   `protobuf:"bytes,90,opt,name=remote_loopback_action,json=remoteLoopbackAction,proto3" json:"remote_loopback_action,omitempty"`
	RequireRemoteMode                    string   `protobuf:"bytes,91,opt,name=require_remote_mode,json=requireRemoteMode,proto3" json:"require_remote_mode,omitempty"`
	RequireRemoteMibRetrieval            bool     `protobuf:"varint,92,opt,name=require_remote_mib_retrieval,json=requireRemoteMibRetrieval,proto3" json:"require_remote_mib_retrieval,omitempty"`
	RequireLoopback                      bool     `protobuf:"varint,93,opt,name=require_loopback,json=requireLoopback,proto3" json:"require_loopback,omitempty"`
	RequireLinkMonitoring                bool     `protobuf:"varint,94,opt,name=require_link_monitoring,json=requireLinkMonitoring,proto3" json:"require_link_monitoring,omitempty"`
	FastHelloIntervalEnabledOverridden   bool     `protobuf:"varint,95,opt,name=fast_hello_interval_enabled_overridden,json=fastHelloIntervalEnabledOverridden,proto3" json:"fast_hello_interval_enabled_overridden,omitempty"`
	LinkMonitoringEnabledOverridden      bool     `protobuf:"varint,96,opt,name=link_monitoring_enabled_overridden,json=linkMonitoringEnabledOverridden,proto3" json:"link_monitoring_enabled_overridden,omitempty"`
	RemoteLoopbackEnabledOverridden      bool     `protobuf:"varint,97,opt,name=remote_loopback_enabled_overridden,json=remoteLoopbackEnabledOverridden,proto3" json:"remote_loopback_enabled_overridden,omitempty"`
	MibRetrievalEnabledOverridden        bool     `protobuf:"varint,98,opt,name=mib_retrieval_enabled_overridden,json=mibRetrievalEnabledOverridden,proto3" json:"mib_retrieval_enabled_overridden,omitempty"`
	UdlfEnabledOverridden                bool     `protobuf:"varint,99,opt,name=udlf_enabled_overridden,json=udlfEnabledOverridden,proto3" json:"udlf_enabled_overridden,omitempty"`
	ModeOverridden                       bool     `protobuf:"varint,100,opt,name=mode_overridden,json=modeOverridden,proto3" json:"mode_overridden,omitempty"`
	ConnectionTimeoutOverridden          bool     `protobuf:"varint,101,opt,name=connection_timeout_overridden,json=connectionTimeoutOverridden,proto3" json:"connection_timeout_overridden,omitempty"`
	SymbolPeriodWindowOverridden         bool     `protobuf:"varint,102,opt,name=symbol_period_window_overridden,json=symbolPeriodWindowOverridden,proto3" json:"symbol_period_window_overridden,omitempty"`
	SymbolPeriodThresholdLowOverridden   bool     `protobuf:"varint,103,opt,name=symbol_period_threshold_low_overridden,json=symbolPeriodThresholdLowOverridden,proto3" json:"symbol_period_threshold_low_overridden,omitempty"`
	SymbolPeriodThresholdHighOverridden  bool     `protobuf:"varint,104,opt,name=symbol_period_threshold_high_overridden,json=symbolPeriodThresholdHighOverridden,proto3" json:"symbol_period_threshold_high_overridden,omitempty"`
	FrameWindowOverridden                bool     `protobuf:"varint,105,opt,name=frame_window_overridden,json=frameWindowOverridden,proto3" json:"frame_window_overridden,omitempty"`
	FrameThresholdLowOverridden          bool     `protobuf:"varint,106,opt,name=frame_threshold_low_overridden,json=frameThresholdLowOverridden,proto3" json:"frame_threshold_low_overridden,omitempty"`
	FrameThresholdHighOverridden         bool     `protobuf:"varint,107,opt,name=frame_threshold_high_overridden,json=frameThresholdHighOverridden,proto3" json:"frame_threshold_high_overridden,omitempty"`
	FramePeriodWindowOverridden          bool     `protobuf:"varint,108,opt,name=frame_period_window_overridden,json=framePeriodWindowOverridden,proto3" json:"frame_period_window_overridden,omitempty"`
	FramePeriodThresholdLowOverridden    bool     `protobuf:"varint,109,opt,name=frame_period_threshold_low_overridden,json=framePeriodThresholdLowOverridden,proto3" json:"frame_period_threshold_low_overridden,omitempty"`
	FramePeriodThresholdHighOverridden   bool     `protobuf:"varint,110,opt,name=frame_period_threshold_high_overridden,json=framePeriodThresholdHighOverridden,proto3" json:"frame_period_threshold_high_overridden,omitempty"`
	FrameSecondsWindowOverridden         bool     `protobuf:"varint,111,opt,name=frame_seconds_window_overridden,json=frameSecondsWindowOverridden,proto3" json:"frame_seconds_window_overridden,omitempty"`
	FrameSecondsThresholdLowOverridden   bool     `protobuf:"varint,112,opt,name=frame_seconds_threshold_low_overridden,json=frameSecondsThresholdLowOverridden,proto3" json:"frame_seconds_threshold_low_overridden,omitempty"`
	FrameSecondsThresholdHighOverridden  bool     `protobuf:"varint,113,opt,name=frame_seconds_threshold_high_overridden,json=frameSecondsThresholdHighOverridden,proto3" json:"frame_seconds_threshold_high_overridden,omitempty"`
	HighThresholdActionOverridden        bool     `protobuf:"varint,114,opt,name=high_threshold_action_overridden,json=highThresholdActionOverridden,proto3" json:"high_threshold_action_overridden,omitempty"`
	LinkFaultActionOverridden            bool     `protobuf:"varint,115,opt,name=link_fault_action_overridden,json=linkFaultActionOverridden,proto3" json:"link_fault_action_overridden,omitempty"`
	DyingGaspActionOverridden            bool     `protobuf:"varint,116,opt,name=dying_gasp_action_overridden,json=dyingGaspActionOverridden,proto3" json:"dying_gasp_action_overridden,omitempty"`
	CriticalEventActionOverridden        bool     `protobuf:"varint,117,opt,name=critical_event_action_overridden,json=criticalEventActionOverridden,proto3" json:"critical_event_action_overridden,omitempty"`
	DiscoveryTimeoutActionOverridden     bool     `protobuf:"varint,118,opt,name=discovery_timeout_action_overridden,json=discoveryTimeoutActionOverridden,proto3" json:"discovery_timeout_action_overridden,omitempty"`
	CapabilitiesConflictActionOverridden bool     `protobuf:"varint,119,opt,name=capabilities_conflict_action_overridden,json=capabilitiesConflictActionOverridden,proto3" json:"capabilities_conflict_action_overridden,omitempty"`
	WiringConflictActionOverridden       bool     `protobuf:"varint,120,opt,name=wiring_conflict_action_overridden,json=wiringConflictActionOverridden,proto3" json:"wiring_conflict_action_overridden,omitempty"`
	SessionDownActionOverridden          bool     `protobuf:"varint,121,opt,name=session_down_action_overridden,json=sessionDownActionOverridden,proto3" json:"session_down_action_overridden,omitempty"`
	SessionUpActionOverridden            bool     `protobuf:"varint,122,opt,name=session_up_action_overridden,json=sessionUpActionOverridden,proto3" json:"session_up_action_overridden,omitempty"`
	RemoteLoopbackActionOverridden       bool     `protobuf:"varint,123,opt,name=remote_loopback_action_overridden,json=remoteLoopbackActionOverridden,proto3" json:"remote_loopback_action_overridden,omitempty"`
	RequireModeOverridden                bool     `protobuf:"varint,124,opt,name=require_mode_overridden,json=requireModeOverridden,proto3" json:"require_mode_overridden,omitempty"`
	RequireMibRetrievalOverridden        bool     `protobuf:"varint,125,opt,name=require_mib_retrieval_overridden,json=requireMibRetrievalOverridden,proto3" json:"require_mib_retrieval_overridden,omitempty"`
	RequireLoopbackOverridden            bool     `protobuf:"varint,126,opt,name=require_loopback_overridden,json=requireLoopbackOverridden,proto3" json:"require_loopback_overridden,omitempty"`
	RequireLinkMonitoringOverridden      bool     `protobuf:"varint,127,opt,name=require_link_monitoring_overridden,json=requireLinkMonitoringOverridden,proto3" json:"require_link_monitoring_overridden,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{} `json:"-"`
	XXX_unrecognized                     []byte   `json:"-"`
	XXX_sizecache                        int32    `json:"-"`
}

func (m *ElodRunningConfig) Reset()         { *m = ElodRunningConfig{} }
func (m *ElodRunningConfig) String() string { return proto.CompactTextString(m) }
func (*ElodRunningConfig) ProtoMessage()    {}
func (*ElodRunningConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0612a046b6621eb, []int{1}
}

func (m *ElodRunningConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElodRunningConfig.Unmarshal(m, b)
}
func (m *ElodRunningConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElodRunningConfig.Marshal(b, m, deterministic)
}
func (m *ElodRunningConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElodRunningConfig.Merge(m, src)
}
func (m *ElodRunningConfig) XXX_Size() int {
	return xxx_messageInfo_ElodRunningConfig.Size(m)
}
func (m *ElodRunningConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ElodRunningConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ElodRunningConfig proto.InternalMessageInfo

func (m *ElodRunningConfig) GetFastHelloIntervalEnabled() bool {
	if m != nil {
		return m.FastHelloIntervalEnabled
	}
	return false
}

func (m *ElodRunningConfig) GetLinkMonitorEnabled() bool {
	if m != nil {
		return m.LinkMonitorEnabled
	}
	return false
}

func (m *ElodRunningConfig) GetRemoteLoopbackEnabled() bool {
	if m != nil {
		return m.RemoteLoopbackEnabled
	}
	return false
}

func (m *ElodRunningConfig) GetMibRetrievalEnabled() bool {
	if m != nil {
		return m.MibRetrievalEnabled
	}
	return false
}

func (m *ElodRunningConfig) GetUdlfEnabled() bool {
	if m != nil {
		return m.UdlfEnabled
	}
	return false
}

func (m *ElodRunningConfig) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ElodRunningConfig) GetConnectionTimeout() uint32 {
	if m != nil {
		return m.ConnectionTimeout
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodWindow() uint32 {
	if m != nil {
		return m.SymbolPeriodWindow
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodWindowUnits() uint32 {
	if m != nil {
		return m.SymbolPeriodWindowUnits
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodWindowMultiplier() uint32 {
	if m != nil {
		return m.SymbolPeriodWindowMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdLow() uint64 {
	if m != nil {
		return m.SymbolPeriodThresholdLow
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdHigh() uint64 {
	if m != nil {
		return m.SymbolPeriodThresholdHigh
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdUnits() uint32 {
	if m != nil {
		return m.SymbolPeriodThresholdUnits
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdLowMultiplier() uint32 {
	if m != nil {
		return m.SymbolPeriodThresholdLowMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdHighMultiplier() uint32 {
	if m != nil {
		return m.SymbolPeriodThresholdHighMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameWindow() uint32 {
	if m != nil {
		return m.FrameWindow
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameThresholdLow() uint64 {
	if m != nil {
		return m.FrameThresholdLow
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameThresholdHigh() uint64 {
	if m != nil {
		return m.FrameThresholdHigh
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameThresholdLowMultiplier() uint32 {
	if m != nil {
		return m.FrameThresholdLowMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameThresholdHighMultiplier() uint32 {
	if m != nil {
		return m.FrameThresholdHighMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodWindow() uint32 {
	if m != nil {
		return m.FramePeriodWindow
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodWindowUnits() uint32 {
	if m != nil {
		return m.FramePeriodWindowUnits
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodWindowMultiplier() uint32 {
	if m != nil {
		return m.FramePeriodWindowMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodThresholdLow() uint64 {
	if m != nil {
		return m.FramePeriodThresholdLow
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodThresholdHigh() uint64 {
	if m != nil {
		return m.FramePeriodThresholdHigh
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodThresholdUnits() uint32 {
	if m != nil {
		return m.FramePeriodThresholdUnits
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodThresholdLowMultiplier() uint32 {
	if m != nil {
		return m.FramePeriodThresholdLowMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFramePeriodThresholdHighMultiplier() uint32 {
	if m != nil {
		return m.FramePeriodThresholdHighMultiplier
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameSecondsWindow() uint32 {
	if m != nil {
		return m.FrameSecondsWindow
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameSecondsThresholdLow() uint64 {
	if m != nil {
		return m.FrameSecondsThresholdLow
	}
	return 0
}

func (m *ElodRunningConfig) GetFrameSecondsThresholdHigh() uint64 {
	if m != nil {
		return m.FrameSecondsThresholdHigh
	}
	return 0
}

func (m *ElodRunningConfig) GetHighThresholdAction() string {
	if m != nil {
		return m.HighThresholdAction
	}
	return ""
}

func (m *ElodRunningConfig) GetLinkFaultAction() string {
	if m != nil {
		return m.LinkFaultAction
	}
	return ""
}

func (m *ElodRunningConfig) GetDyingGaspAction() string {
	if m != nil {
		return m.DyingGaspAction
	}
	return ""
}

func (m *ElodRunningConfig) GetCriticalEventAction() string {
	if m != nil {
		return m.CriticalEventAction
	}
	return ""
}

func (m *ElodRunningConfig) GetDiscoveryTimeoutAction() string {
	if m != nil {
		return m.DiscoveryTimeoutAction
	}
	return ""
}

func (m *ElodRunningConfig) GetCapabilitiesConflictAction() string {
	if m != nil {
		return m.CapabilitiesConflictAction
	}
	return ""
}

func (m *ElodRunningConfig) GetWiringConflictAction() string {
	if m != nil {
		return m.WiringConflictAction
	}
	return ""
}

func (m *ElodRunningConfig) GetSessionUpAction() string {
	if m != nil {
		return m.SessionUpAction
	}
	return ""
}

func (m *ElodRunningConfig) GetSessionDownAction() string {
	if m != nil {
		return m.SessionDownAction
	}
	return ""
}

func (m *ElodRunningConfig) GetRemoteLoopbackAction() string {
	if m != nil {
		return m.RemoteLoopbackAction
	}
	return ""
}

func (m *ElodRunningConfig) GetRequireRemoteMode() string {
	if m != nil {
		return m.RequireRemoteMode
	}
	return ""
}

func (m *ElodRunningConfig) GetRequireRemoteMibRetrieval() bool {
	if m != nil {
		return m.RequireRemoteMibRetrieval
	}
	return false
}

func (m *ElodRunningConfig) GetRequireLoopback() bool {
	if m != nil {
		return m.RequireLoopback
	}
	return false
}

func (m *ElodRunningConfig) GetRequireLinkMonitoring() bool {
	if m != nil {
		return m.RequireLinkMonitoring
	}
	return false
}

func (m *ElodRunningConfig) GetFastHelloIntervalEnabledOverridden() bool {
	if m != nil {
		return m.FastHelloIntervalEnabledOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetLinkMonitoringEnabledOverridden() bool {
	if m != nil {
		return m.LinkMonitoringEnabledOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRemoteLoopbackEnabledOverridden() bool {
	if m != nil {
		return m.RemoteLoopbackEnabledOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetMibRetrievalEnabledOverridden() bool {
	if m != nil {
		return m.MibRetrievalEnabledOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetUdlfEnabledOverridden() bool {
	if m != nil {
		return m.UdlfEnabledOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetModeOverridden() bool {
	if m != nil {
		return m.ModeOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetConnectionTimeoutOverridden() bool {
	if m != nil {
		return m.ConnectionTimeoutOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetSymbolPeriodWindowOverridden() bool {
	if m != nil {
		return m.SymbolPeriodWindowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdLowOverridden() bool {
	if m != nil {
		return m.SymbolPeriodThresholdLowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetSymbolPeriodThresholdHighOverridden() bool {
	if m != nil {
		return m.SymbolPeriodThresholdHighOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameWindowOverridden() bool {
	if m != nil {
		return m.FrameWindowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameThresholdLowOverridden() bool {
	if m != nil {
		return m.FrameThresholdLowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameThresholdHighOverridden() bool {
	if m != nil {
		return m.FrameThresholdHighOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFramePeriodWindowOverridden() bool {
	if m != nil {
		return m.FramePeriodWindowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFramePeriodThresholdLowOverridden() bool {
	if m != nil {
		return m.FramePeriodThresholdLowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFramePeriodThresholdHighOverridden() bool {
	if m != nil {
		return m.FramePeriodThresholdHighOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameSecondsWindowOverridden() bool {
	if m != nil {
		return m.FrameSecondsWindowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameSecondsThresholdLowOverridden() bool {
	if m != nil {
		return m.FrameSecondsThresholdLowOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetFrameSecondsThresholdHighOverridden() bool {
	if m != nil {
		return m.FrameSecondsThresholdHighOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetHighThresholdActionOverridden() bool {
	if m != nil {
		return m.HighThresholdActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetLinkFaultActionOverridden() bool {
	if m != nil {
		return m.LinkFaultActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetDyingGaspActionOverridden() bool {
	if m != nil {
		return m.DyingGaspActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetCriticalEventActionOverridden() bool {
	if m != nil {
		return m.CriticalEventActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetDiscoveryTimeoutActionOverridden() bool {
	if m != nil {
		return m.DiscoveryTimeoutActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetCapabilitiesConflictActionOverridden() bool {
	if m != nil {
		return m.CapabilitiesConflictActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetWiringConflictActionOverridden() bool {
	if m != nil {
		return m.WiringConflictActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetSessionDownActionOverridden() bool {
	if m != nil {
		return m.SessionDownActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetSessionUpActionOverridden() bool {
	if m != nil {
		return m.SessionUpActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRemoteLoopbackActionOverridden() bool {
	if m != nil {
		return m.RemoteLoopbackActionOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRequireModeOverridden() bool {
	if m != nil {
		return m.RequireModeOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRequireMibRetrievalOverridden() bool {
	if m != nil {
		return m.RequireMibRetrievalOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRequireLoopbackOverridden() bool {
	if m != nil {
		return m.RequireLoopbackOverridden
	}
	return false
}

func (m *ElodRunningConfig) GetRequireLinkMonitoringOverridden() bool {
	if m != nil {
		return m.RequireLinkMonitoringOverridden
	}
	return false
}

func init() {
	proto.RegisterType((*ElodRunningConfig_KEYS)(nil), "cisco_ios_xr_ethernet_link_oam_oper.ether_link_oam.running_config_interfaces.running_config_interface.elod_running_config_KEYS")
	proto.RegisterType((*ElodRunningConfig)(nil), "cisco_ios_xr_ethernet_link_oam_oper.ether_link_oam.running_config_interfaces.running_config_interface.elod_running_config")
}

func init() { proto.RegisterFile("elod_running_config.proto", fileDescriptor_a0612a046b6621eb) }

var fileDescriptor_a0612a046b6621eb = []byte{
	// 1418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x98, 0xed, 0x53, 0xdb, 0x46,
	0x10, 0xc6, 0x87, 0x99, 0x4c, 0xa7, 0xbd, 0xa4, 0x4d, 0x11, 0x01, 0x04, 0x98, 0x60, 0x48, 0x0b,
	0x34, 0x33, 0x65, 0x32, 0x24, 0xa5, 0x49, 0x69, 0x42, 0x08, 0x71, 0x80, 0x82, 0x1b, 0x6a, 0xa0,
	0x69, 0xfa, 0xa6, 0xca, 0xd2, 0xd9, 0xbe, 0x22, 0xe9, 0x1c, 0x49, 0xc6, 0xa1, 0xaf, 0xff, 0x77,
	0x3f, 0x75, 0xb4, 0xa7, 0x93, 0xf7, 0xa4, 0x3b, 0xe7, 0x1b, 0xa3, 0xdd, 0xe7, 0xe7, 0xdd, 0xbd,
	0xb7, 0x67, 0x20, 0x73, 0x34, 0xe0, 0xbe, 0x13, 0x0f, 0xa2, 0x88, 0x45, 0x5d, 0xc7, 0xe3, 0x51,
	0x87, 0x75, 0x37, 0xfa, 0x31, 0x4f, 0xb9, 0x45, 0x3d, 0x96, 0x78, 0xdc, 0x61, 0x3c, 0x71, 0xde,
	0xc6, 0x0e, 0x4d, 0x7b, 0x34, 0x8e, 0x68, 0xea, 0x04, 0x2c, 0xba, 0x70, 0xb8, 0x1b, 0x3a, 0xbc,
	0x4f, 0xe3, 0x0d, 0xf8, 0x5c, 0x7c, 0xdb, 0x50, 0x41, 0x0e, 0x8b, 0x52, 0x1a, 0x77, 0x5c, 0x8f,
	0x26, 0xc6, 0xc8, 0x4a, 0x83, 0xd8, 0x9a, 0x1a, 0x9c, 0xa3, 0xc6, 0xeb, 0x53, 0xeb, 0x33, 0xf2,
	0x71, 0x48, 0xc3, 0x36, 0x8d, 0x47, 0xf9, 0xf6, 0x44, 0x7d, 0x62, 0xfd, 0x83, 0xd6, 0x4d, 0xf1,
	0xfd, 0xb0, 0xc0, 0xfc, 0xb7, 0x4a, 0xa6, 0x34, 0x1c, 0xeb, 0x31, 0x59, 0xe8, 0xb8, 0x49, 0xea,
	0xf4, 0x68, 0x10, 0x70, 0x81, 0xb9, 0x74, 0x03, 0x87, 0x46, 0x6e, 0x3b, 0xa0, 0xbe, 0xbd, 0x59,
	0x9f, 0x58, 0x7f, 0xbf, 0x65, 0x67, 0x29, 0x07, 0x59, 0xc6, 0x61, 0x9e, 0xd0, 0x10, 0x71, 0xeb,
	0x1e, 0xb9, 0x05, 0xcd, 0x85, 0x3c, 0x62, 0x29, 0x8f, 0x0b, 0xdd, 0x7d, 0xd0, 0x59, 0x59, 0xac,
	0x29, 0x42, 0x52, 0xb1, 0x45, 0x66, 0x63, 0x1a, 0xf2, 0x94, 0x3a, 0x01, 0xe7, 0xfd, 0xb6, 0xeb,
	0x5d, 0x14, 0xa2, 0x07, 0x20, 0x9a, 0x16, 0xe1, 0xe3, 0x3c, 0x2a, 0x75, 0x9b, 0x64, 0x3a, 0x64,
	0x6d, 0x27, 0xa6, 0x69, 0xcc, 0x28, 0x2e, 0xf1, 0x0b, 0x50, 0x4d, 0x85, 0xac, 0xdd, 0x92, 0x31,
	0xa9, 0x59, 0x26, 0x37, 0x06, 0x7e, 0xd0, 0x29, 0x52, 0xb7, 0x20, 0xf5, 0x7a, 0xf6, 0x4d, 0xa6,
	0x58, 0xe4, 0x5a, 0xc8, 0x7d, 0x6a, 0x7f, 0x09, 0x63, 0x83, 0xbf, 0xad, 0xcf, 0x89, 0xe5, 0xf1,
	0x28, 0xa2, 0x5e, 0xca, 0x78, 0xe4, 0xa4, 0x2c, 0xa4, 0x7c, 0x90, 0xda, 0x0f, 0xeb, 0x13, 0xeb,
	0x1f, 0xb6, 0x26, 0x47, 0x91, 0x33, 0x11, 0xc8, 0x66, 0x90, 0x5c, 0x85, 0x6d, 0x1e, 0x38, 0x7d,
	0x1a, 0x33, 0xee, 0x3b, 0x43, 0x16, 0xf9, 0x7c, 0x68, 0x3f, 0x02, 0x81, 0x25, 0x62, 0x27, 0x10,
	0x7a, 0x05, 0x11, 0x6b, 0x9b, 0xcc, 0xeb, 0x14, 0xce, 0x20, 0x62, 0x69, 0x62, 0x7f, 0x05, 0xba,
	0xd9, 0xaa, 0xee, 0x3c, 0x0b, 0x5b, 0x0d, 0xb2, 0xa4, 0x15, 0x87, 0x83, 0x20, 0x65, 0xfd, 0x80,
	0xd1, 0xd8, 0xde, 0x06, 0x42, 0xad, 0x4a, 0x68, 0x16, 0x39, 0xd9, 0xc2, 0xab, 0x98, 0xb4, 0x17,
	0xd3, 0xa4, 0xc7, 0x03, 0xdf, 0x09, 0xf8, 0xd0, 0xfe, 0xba, 0x3e, 0xb1, 0x7e, 0xad, 0x65, 0x63,
	0xc4, 0x99, 0x4c, 0x38, 0xe6, 0x43, 0x6b, 0x87, 0xd4, 0x4c, 0xf2, 0x1e, 0xeb, 0xf6, 0xec, 0xc7,
	0xa0, 0x9f, 0xd3, 0xea, 0x0f, 0x58, 0xb7, 0x67, 0xed, 0x92, 0x45, 0x13, 0x40, 0x8c, 0xe1, 0x09,
	0x34, 0x31, 0xaf, 0x25, 0x88, 0x49, 0xb4, 0xc8, 0xea, 0x98, 0x16, 0xf0, 0x40, 0x76, 0x80, 0xb5,
	0x62, 0xea, 0x06, 0x8d, 0xe5, 0x8c, 0xac, 0x8d, 0xeb, 0x0b, 0x43, 0x9f, 0x02, 0xf4, 0x8e, 0xb1,
	0x45, 0x44, 0x5d, 0x26, 0x37, 0x3a, 0xb1, 0x1b, 0x52, 0xb9, 0x35, 0x76, 0x41, 0x7a, 0x1d, 0xbe,
	0xe5, 0x7b, 0x62, 0x83, 0x4c, 0x89, 0x14, 0x75, 0x1d, 0x9e, 0xc1, 0x1c, 0x27, 0x21, 0xa4, 0x2c,
	0xc0, 0x3d, 0x72, 0xab, 0x9c, 0x0f, 0x83, 0xdf, 0x03, 0x81, 0xa5, 0x0a, 0x60, 0xe2, 0x7b, 0xe4,
	0xb6, 0xe6, 0x17, 0x70, 0x47, 0xcf, 0xa1, 0xac, 0x85, 0xca, 0x8f, 0xa1, 0x4e, 0x1a, 0x64, 0x49,
	0xf7, 0xb3, 0x98, 0xd2, 0x10, 0xbb, 0xaf, 0x5a, 0x01, 0xc2, 0x14, 0xdd, 0xaa, 0x47, 0xe6, 0x85,
	0x38, 0x63, 0x10, 0x52, 0x4e, 0xcc, 0x23, 0x32, 0xa7, 0xc9, 0xcf, 0x77, 0xca, 0x3e, 0xa8, 0x66,
	0x2a, 0x2a, 0xb1, 0x4b, 0x8a, 0xb6, 0x8d, 0xc7, 0xe5, 0x00, 0xb5, 0x6d, 0x38, 0x2d, 0xdb, 0x64,
	0x5e, 0x81, 0xa8, 0x8b, 0x74, 0x08, 0x33, 0x9f, 0x45, 0x00, 0x65, 0xa9, 0xb2, 0x3b, 0x56, 0x2f,
	0x86, 0x15, 0xfb, 0x46, 0x1c, 0x35, 0x9d, 0x1a, 0xd6, 0x6d, 0x87, 0xd4, 0x0c, 0x72, 0xd1, 0xfe,
	0x11, 0x94, 0x3f, 0xa7, 0xd3, 0x8b, 0x09, 0x9c, 0x90, 0x4f, 0xcd, 0xc5, 0xe3, 0x41, 0x1c, 0x03,
	0x69, 0xd9, 0xd0, 0x07, 0x1a, 0x47, 0x8b, 0xac, 0x8e, 0xe9, 0x08, 0x23, 0x9b, 0xe2, 0xe4, 0x99,
	0x9a, 0x43, 0xcc, 0x62, 0x43, 0x27, 0xd4, 0xe3, 0x91, 0x9f, 0xc8, 0x3d, 0xf1, 0xad, 0xb8, 0x46,
	0x21, 0x76, 0x2a, 0x42, 0xf9, 0xa6, 0x28, 0xe6, 0x2a, 0x15, 0xea, 0xaa, 0xbc, 0x44, 0x73, 0xcd,
	0x85, 0xe5, 0x2b, 0xcc, 0x24, 0x87, 0x75, 0x39, 0x11, 0x57, 0x98, 0x56, 0x0f, 0x0b, 0xb3, 0x49,
	0xa6, 0xa1, 0xdd, 0x91, 0xce, 0x85, 0x87, 0xc1, 0xfe, 0x0e, 0x1e, 0x93, 0xa9, 0x2c, 0x58, 0x28,
	0x76, 0x21, 0x64, 0xdd, 0x25, 0x93, 0xf0, 0x60, 0x76, 0xdc, 0x41, 0x90, 0xca, 0xfc, 0x96, 0x78,
	0xb3, 0xb3, 0xc0, 0x8b, 0xec, 0xfb, 0x28, 0xd7, 0xbf, 0xca, 0xde, 0xea, 0xae, 0x9b, 0xf4, 0x65,
	0xee, 0xa9, 0xc8, 0x85, 0xc0, 0xbe, 0x9b, 0xf4, 0xf3, 0xdc, 0x4d, 0x32, 0xed, 0xc5, 0x2c, 0x65,
	0x5e, 0xf6, 0x32, 0x5e, 0xd2, 0xa8, 0x60, 0x9f, 0x89, 0x5a, 0x64, 0xb0, 0x91, 0xc5, 0x72, 0xcd,
	0x43, 0x62, 0xfb, 0x99, 0x87, 0xb9, 0xa4, 0xf1, 0x95, 0x7c, 0xe6, 0xa4, 0xec, 0x1c, 0x64, 0x33,
	0x45, 0x3c, 0x7f, 0xec, 0x72, 0xe5, 0x53, 0x52, 0xf3, 0xdc, 0xbe, 0xdb, 0x66, 0x01, 0x4b, 0x19,
	0x4d, 0xc0, 0x4c, 0x04, 0xcc, 0x2b, 0xd4, 0xdf, 0x83, 0x7a, 0x1e, 0xe7, 0xec, 0xe5, 0x29, 0x39,
	0xe1, 0x01, 0x99, 0x19, 0xb2, 0x58, 0x1a, 0x11, 0xac, 0x7d, 0x05, 0xda, 0x5b, 0x22, 0x5a, 0x52,
	0xdd, 0x25, 0x93, 0x09, 0x4d, 0x92, 0xec, 0x59, 0x1e, 0x14, 0x13, 0xf9, 0x41, 0x4c, 0x24, 0x0f,
	0x9c, 0xcb, 0x89, 0x6c, 0x90, 0x29, 0x99, 0xeb, 0xf3, 0x61, 0x24, 0xb3, 0x5f, 0x43, 0xb6, 0xc4,
	0x3c, 0xe7, 0xc3, 0x68, 0x54, 0x51, 0xd9, 0x98, 0xe4, 0x92, 0x1f, 0x45, 0x45, 0xaa, 0x2f, 0x19,
	0xfd, 0x4a, 0x4c, 0xdf, 0x0c, 0x58, 0x4c, 0x9d, 0x5c, 0x0d, 0x76, 0xe2, 0x27, 0xf1, 0x2b, 0x79,
	0xa8, 0x05, 0x91, 0x66, 0xe6, 0x2d, 0x76, 0x48, 0xad, 0x9c, 0x8f, 0x5d, 0x8d, 0xfd, 0x33, 0x58,
	0x94, 0x39, 0x55, 0x88, 0xac, 0x4d, 0xe6, 0xf9, 0x24, 0x40, 0xd6, 0x69, 0xff, 0x02, 0xa2, 0x9b,
	0xf9, 0x77, 0x59, 0xa1, 0xb0, 0x5a, 0x79, 0x2a, 0x32, 0x69, 0x2c, 0xea, 0xda, 0xbf, 0x4a, 0xab,
	0x25, 0x14, 0x23, 0x9b, 0xc6, 0xa2, 0x2e, 0x9c, 0x6e, 0xb3, 0x27, 0x74, 0xb2, 0xfd, 0x10, 0x33,
	0xdf, 0xa7, 0x91, 0xed, 0x00, 0x66, 0xc5, 0x64, 0x0f, 0x5f, 0x16, 0x99, 0xd6, 0x11, 0x59, 0x29,
	0xd5, 0xa0, 0xe3, 0xfd, 0x06, 0xbc, 0xa5, 0x40, 0xa9, 0x47, 0x0b, 0x33, 0x78, 0x48, 0x0c, 0x73,
	0x05, 0x4c, 0x6b, 0x27, 0x11, 0x6c, 0x9f, 0xd4, 0xb5, 0xc6, 0x12, 0xa3, 0xda, 0x80, 0x5a, 0xd4,
	0x78, 0x4c, 0x04, 0xda, 0x22, 0xb3, 0xd8, 0x6d, 0x62, 0xbd, 0x27, 0xc6, 0x8d, 0x8c, 0x27, 0xd2,
	0xad, 0x91, 0x9b, 0xd9, 0x9e, 0xc1, 0xf9, 0x3e, 0xe4, 0x7f, 0x94, 0x7d, 0x46, 0x89, 0xcf, 0xc8,
	0x62, 0xd5, 0x97, 0x62, 0x19, 0x05, 0xd9, 0x42, 0xc5, 0xa2, 0x22, 0x86, 0xc9, 0x3d, 0x22, 0x4a,
	0x07, 0x28, 0x1a, 0xf7, 0x88, 0x30, 0xef, 0xb0, 0x5e, 0x88, 0xd6, 0x15, 0x5b, 0xc4, 0x64, 0xbd,
	0x10, 0xf3, 0x5d, 0xd6, 0x0b, 0x41, 0x7b, 0x00, 0x35, 0x5b, 0x2f, 0x75, 0x55, 0xb0, 0xf5, 0xc2,
	0x14, 0x26, 0x56, 0x05, 0xb9, 0x30, 0xa4, 0x33, 0xb8, 0x25, 0x24, 0xff, 0x5d, 0x4c, 0xbb, 0xe2,
	0x96, 0xd4, 0x69, 0x6b, 0xdd, 0x12, 0xa2, 0x5c, 0x88, 0x69, 0x57, 0xdd, 0x92, 0xae, 0x16, 0xe3,
	0x9a, 0x05, 0xa8, 0x16, 0xc3, 0x92, 0x8d, 0x77, 0x01, 0x88, 0x15, 0x02, 0xcb, 0xe4, 0x02, 0xd4,
	0x4d, 0x30, 0xce, 0x05, 0x20, 0x64, 0x94, 0xdf, 0x13, 0x06, 0x17, 0xa0, 0x9b, 0x98, 0xea, 0x02,
	0x30, 0x8c, 0xa3, 0x89, 0x29, 0x86, 0x40, 0x57, 0x9a, 0xd6, 0x1a, 0x60, 0x5a, 0x1f, 0x95, 0xa6,
	0x71, 0x09, 0xea, 0xfe, 0x1c, 0xe7, 0x17, 0x30, 0xf4, 0x8d, 0xd8, 0x9f, 0x46, 0xeb, 0xa0, 0x5e,
	0x3f, 0x5a, 0x13, 0x81, 0x71, 0xb1, 0xb8, 0x7e, 0x34, 0x7e, 0x02, 0x81, 0x76, 0x48, 0xad, 0xe2,
	0x2c, 0x30, 0x24, 0x11, 0x2f, 0x4b, 0xc9, 0x64, 0xa8, 0x80, 0x8a, 0xdd, 0xc0, 0x80, 0x54, 0x00,
	0x4a, 0xce, 0x43, 0x6d, 0x45, 0xeb, 0x41, 0x30, 0x64, 0x20, 0x5a, 0xd1, 0xd8, 0x11, 0x04, 0x6a,
	0x92, 0x3b, 0x26, 0x63, 0x82, 0x59, 0x97, 0xc0, 0xaa, 0xeb, 0x3d, 0x0a, 0xc2, 0x9d, 0x93, 0xb5,
	0x71, 0x6e, 0x05, 0x23, 0x87, 0x80, 0xfc, 0xc4, 0x6c, 0x5c, 0x10, 0xf6, 0x90, 0x2c, 0xeb, 0x2d,
	0x0c, 0x06, 0xbe, 0x05, 0xe0, 0x6d, 0x9d, 0x9b, 0x51, 0x0f, 0xb8, 0xc6, 0xab, 0x60, 0xce, 0x95,
	0x38, 0xe0, 0x15, 0xdb, 0xa2, 0xae, 0x5f, 0xc5, 0x1c, 0x61, 0xc4, 0x1f, 0x62, 0xfd, 0x4a, 0x3e,
	0x49, 0x6d, 0x48, 0xef, 0x80, 0x30, 0xe5, 0x4f, 0xd1, 0x90, 0xce, 0x0c, 0xa9, 0xb7, 0xae, 0xb4,
	0x1e, 0xe5, 0xb7, 0xed, 0x2f, 0xc5, 0x7a, 0x34, 0xd5, 0x27, 0x6e, 0x9f, 0xd4, 0x0b, 0x9d, 0xf2,
	0x28, 0x23, 0xc0, 0xdf, 0x62, 0x0b, 0x49, 0x00, 0x7a, 0x93, 0x11, 0xe8, 0x09, 0x59, 0x28, 0xdb,
	0x24, 0xcc, 0xf8, 0x47, 0xb1, 0x59, 0xb2, 0x8d, 0xb2, 0xc5, 0xd0, 0x7a, 0x27, 0x8c, 0xf9, 0x57,
	0x5a, 0x0c, 0x8d, 0x8d, 0x1a, 0xc1, 0xda, 0xef, 0xc1, 0x7f, 0x0c, 0xef, 0xff, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0x8f, 0xe4, 0x94, 0x4e, 0x14, 0x00, 0x00,
}
