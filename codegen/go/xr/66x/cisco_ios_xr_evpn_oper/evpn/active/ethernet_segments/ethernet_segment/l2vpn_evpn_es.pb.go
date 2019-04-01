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
// source: l2vpn_evpn_es.proto

package cisco_ios_xr_evpn_oper_evpn_active_ethernet_segments_ethernet_segment

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

type L2VpnEvpnEs_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Esi1                 string   `protobuf:"bytes,2,opt,name=esi1,proto3" json:"esi1,omitempty"`
	Esi2                 string   `protobuf:"bytes,3,opt,name=esi2,proto3" json:"esi2,omitempty"`
	Esi3                 string   `protobuf:"bytes,4,opt,name=esi3,proto3" json:"esi3,omitempty"`
	Esi4                 string   `protobuf:"bytes,5,opt,name=esi4,proto3" json:"esi4,omitempty"`
	Esi5                 string   `protobuf:"bytes,6,opt,name=esi5,proto3" json:"esi5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEs_KEYS) Reset()         { *m = L2VpnEvpnEs_KEYS{} }
func (m *L2VpnEvpnEs_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEs_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnEs_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_95dd9c21af48b1f6, []int{0}
}

func (m *L2VpnEvpnEs_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEs_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnEs_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEs_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEs_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEs_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnEs_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEs_KEYS.Size(m)
}
func (m *L2VpnEvpnEs_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEs_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEs_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnEs_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2VpnEvpnEs_KEYS) GetEsi1() string {
	if m != nil {
		return m.Esi1
	}
	return ""
}

func (m *L2VpnEvpnEs_KEYS) GetEsi2() string {
	if m != nil {
		return m.Esi2
	}
	return ""
}

func (m *L2VpnEvpnEs_KEYS) GetEsi3() string {
	if m != nil {
		return m.Esi3
	}
	return ""
}

func (m *L2VpnEvpnEs_KEYS) GetEsi4() string {
	if m != nil {
		return m.Esi4
	}
	return ""
}

func (m *L2VpnEvpnEs_KEYS) GetEsi5() string {
	if m != nil {
		return m.Esi5
	}
	return ""
}

type L2VpnEvpnEsNhBuffer struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	DfDontPrempt         bool     `protobuf:"varint,2,opt,name=df_dont_prempt,json=dfDontPrempt,proto3" json:"df_dont_prempt,omitempty"`
	DfType               uint32   `protobuf:"varint,3,opt,name=df_type,json=dfType,proto3" json:"df_type,omitempty"`
	DfPref               uint32   `protobuf:"varint,4,opt,name=df_pref,json=dfPref,proto3" json:"df_pref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEsNhBuffer) Reset()         { *m = L2VpnEvpnEsNhBuffer{} }
func (m *L2VpnEvpnEsNhBuffer) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEsNhBuffer) ProtoMessage()    {}
func (*L2VpnEvpnEsNhBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_95dd9c21af48b1f6, []int{1}
}

func (m *L2VpnEvpnEsNhBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEsNhBuffer.Unmarshal(m, b)
}
func (m *L2VpnEvpnEsNhBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEsNhBuffer.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEsNhBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEsNhBuffer.Merge(m, src)
}
func (m *L2VpnEvpnEsNhBuffer) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEsNhBuffer.Size(m)
}
func (m *L2VpnEvpnEsNhBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEsNhBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEsNhBuffer proto.InternalMessageInfo

func (m *L2VpnEvpnEsNhBuffer) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnEvpnEsNhBuffer) GetDfDontPrempt() bool {
	if m != nil {
		return m.DfDontPrempt
	}
	return false
}

func (m *L2VpnEvpnEsNhBuffer) GetDfType() uint32 {
	if m != nil {
		return m.DfType
	}
	return 0
}

func (m *L2VpnEvpnEsNhBuffer) GetDfPref() uint32 {
	if m != nil {
		return m.DfPref
	}
	return 0
}

type L2VpnEvpnEadServiceInfo struct {
	VpnId                uint32   `protobuf:"varint,1,opt,name=vpn_id,json=vpnId,proto3" json:"vpn_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	EthernetTag          uint32   `protobuf:"varint,3,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEadServiceInfo) Reset()         { *m = L2VpnEvpnEadServiceInfo{} }
func (m *L2VpnEvpnEadServiceInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEadServiceInfo) ProtoMessage()    {}
func (*L2VpnEvpnEadServiceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_95dd9c21af48b1f6, []int{2}
}

func (m *L2VpnEvpnEadServiceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEadServiceInfo.Unmarshal(m, b)
}
func (m *L2VpnEvpnEadServiceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEadServiceInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEadServiceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEadServiceInfo.Merge(m, src)
}
func (m *L2VpnEvpnEadServiceInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEadServiceInfo.Size(m)
}
func (m *L2VpnEvpnEadServiceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEadServiceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEadServiceInfo proto.InternalMessageInfo

func (m *L2VpnEvpnEadServiceInfo) GetVpnId() uint32 {
	if m != nil {
		return m.VpnId
	}
	return 0
}

func (m *L2VpnEvpnEadServiceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *L2VpnEvpnEadServiceInfo) GetEthernetTag() uint32 {
	if m != nil {
		return m.EthernetTag
	}
	return 0
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
	return fileDescriptor_95dd9c21af48b1f6, []int{3}
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

type L2VpnEvpnEs struct {
	EthernetSegmentIdentifier          []uint32                   `protobuf:"varint,50,rep,packed,name=ethernet_segment_identifier,json=ethernetSegmentIdentifier,proto3" json:"ethernet_segment_identifier,omitempty"`
	EsiType                            string                     `protobuf:"bytes,51,opt,name=esi_type,json=esiType,proto3" json:"esi_type,omitempty"`
	EsiSystemIdentifier                string                     `protobuf:"bytes,52,opt,name=esi_system_identifier,json=esiSystemIdentifier,proto3" json:"esi_system_identifier,omitempty"`
	EsiPortKey                         uint32                     `protobuf:"varint,53,opt,name=esi_port_key,json=esiPortKey,proto3" json:"esi_port_key,omitempty"`
	EsiSystemPriority                  uint32                     `protobuf:"varint,54,opt,name=esi_system_priority,json=esiSystemPriority,proto3" json:"esi_system_priority,omitempty"`
	EthernetSegmentName                string                     `protobuf:"bytes,55,opt,name=ethernet_segment_name,json=ethernetSegmentName,proto3" json:"ethernet_segment_name,omitempty"`
	EthernetSegmentState               string                     `protobuf:"bytes,56,opt,name=ethernet_segment_state,json=ethernetSegmentState,proto3" json:"ethernet_segment_state,omitempty"`
	IfHandle                           string                     `protobuf:"bytes,57,opt,name=if_handle,json=ifHandle,proto3" json:"if_handle,omitempty"`
	MainPortRole                       string                     `protobuf:"bytes,58,opt,name=main_port_role,json=mainPortRole,proto3" json:"main_port_role,omitempty"`
	MainPortMac                        string                     `protobuf:"bytes,59,opt,name=main_port_mac,json=mainPortMac,proto3" json:"main_port_mac,omitempty"`
	NumUpPWs                           uint32                     `protobuf:"varint,60,opt,name=num_up_p_ws,json=numUpPWs,proto3" json:"num_up_p_ws,omitempty"`
	RouteTarget                        string                     `protobuf:"bytes,61,opt,name=route_target,json=routeTarget,proto3" json:"route_target,omitempty"`
	RtOrigin                           string                     `protobuf:"bytes,62,opt,name=rt_origin,json=rtOrigin,proto3" json:"rt_origin,omitempty"`
	EsBgpGates                         string                     `protobuf:"bytes,63,opt,name=es_bgp_gates,json=esBgpGates,proto3" json:"es_bgp_gates,omitempty"`
	EsL2FibGates                       string                     `protobuf:"bytes,64,opt,name=es_l2fib_gates,json=esL2fibGates,proto3" json:"es_l2fib_gates,omitempty"`
	MacFlushingModeConfig              string                     `protobuf:"bytes,65,opt,name=mac_flushing_mode_config,json=macFlushingModeConfig,proto3" json:"mac_flushing_mode_config,omitempty"`
	LoadBalanceModeConfig              string                     `protobuf:"bytes,66,opt,name=load_balance_mode_config,json=loadBalanceModeConfig,proto3" json:"load_balance_mode_config,omitempty"`
	LoadBalanceModeIsDefault           bool                       `protobuf:"varint,67,opt,name=load_balance_mode_is_default,json=loadBalanceModeIsDefault,proto3" json:"load_balance_mode_is_default,omitempty"`
	LoadBalanceModeOper                string                     `protobuf:"bytes,68,opt,name=load_balance_mode_oper,json=loadBalanceModeOper,proto3" json:"load_balance_mode_oper,omitempty"`
	ForceSingleHome                    bool                       `protobuf:"varint,69,opt,name=force_single_home,json=forceSingleHome,proto3" json:"force_single_home,omitempty"`
	SourceMacOper                      string                     `protobuf:"bytes,70,opt,name=source_mac_oper,json=sourceMacOper,proto3" json:"source_mac_oper,omitempty"`
	SourceMacOrigin                    string                     `protobuf:"bytes,71,opt,name=source_mac_origin,json=sourceMacOrigin,proto3" json:"source_mac_origin,omitempty"`
	PeeringTimer                       uint32                     `protobuf:"varint,72,opt,name=peering_timer,json=peeringTimer,proto3" json:"peering_timer,omitempty"`
	PeeringTimerLeft                   uint32                     `protobuf:"varint,73,opt,name=peering_timer_left,json=peeringTimerLeft,proto3" json:"peering_timer_left,omitempty"`
	RecoveryTimer                      uint32                     `protobuf:"varint,74,opt,name=recovery_timer,json=recoveryTimer,proto3" json:"recovery_timer,omitempty"`
	RecoveryTimerLeft                  uint32                     `protobuf:"varint,75,opt,name=recovery_timer_left,json=recoveryTimerLeft,proto3" json:"recovery_timer_left,omitempty"`
	CarvingTimer                       uint32                     `protobuf:"varint,76,opt,name=carving_timer,json=carvingTimer,proto3" json:"carving_timer,omitempty"`
	CarvingTimerLeft                   uint32                     `protobuf:"varint,77,opt,name=carving_timer_left,json=carvingTimerLeft,proto3" json:"carving_timer_left,omitempty"`
	ServiceCarvingMode                 string                     `protobuf:"bytes,78,opt,name=service_carving_mode,json=serviceCarvingMode,proto3" json:"service_carving_mode,omitempty"`
	PrimaryService                     []uint32                   `protobuf:"varint,79,rep,packed,name=primary_service,json=primaryService,proto3" json:"primary_service,omitempty"`
	PrimaryServicesInput               string                     `protobuf:"bytes,80,opt,name=primary_services_input,json=primaryServicesInput,proto3" json:"primary_services_input,omitempty"`
	SecondaryService                   []uint32                   `protobuf:"varint,81,rep,packed,name=secondary_service,json=secondaryService,proto3" json:"secondary_service,omitempty"`
	SecondaryServicesInput             string                     `protobuf:"bytes,82,opt,name=secondary_services_input,json=secondaryServicesInput,proto3" json:"secondary_services_input,omitempty"`
	NextHop                            []*L2VpnEvpnEsNhBuffer     `protobuf:"bytes,83,rep,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	ForwarderPorts                     uint32                     `protobuf:"varint,84,opt,name=forwarder_ports,json=forwarderPorts,proto3" json:"forwarder_ports,omitempty"`
	PermanentForwarderPorts            uint32                     `protobuf:"varint,85,opt,name=permanent_forwarder_ports,json=permanentForwarderPorts,proto3" json:"permanent_forwarder_ports,omitempty"`
	ElectedForwarderPorts              uint32                     `protobuf:"varint,86,opt,name=elected_forwarder_ports,json=electedForwarderPorts,proto3" json:"elected_forwarder_ports,omitempty"`
	NotElectedForwarderPorts           uint32                     `protobuf:"varint,87,opt,name=not_elected_forwarder_ports,json=notElectedForwarderPorts,proto3" json:"not_elected_forwarder_ports,omitempty"`
	NotConfigForwarderPorts            uint32                     `protobuf:"varint,88,opt,name=not_config_forwarder_ports,json=notConfigForwarderPorts,proto3" json:"not_config_forwarder_ports,omitempty"`
	MpProtected                        bool                       `protobuf:"varint,89,opt,name=mp_protected,json=mpProtected,proto3" json:"mp_protected,omitempty"`
	NveAnycastVtep                     bool                       `protobuf:"varint,90,opt,name=nve_anycast_vtep,json=nveAnycastVtep,proto3" json:"nve_anycast_vtep,omitempty"`
	NveIngressReplication              bool                       `protobuf:"varint,91,opt,name=nve_ingress_replication,json=nveIngressReplication,proto3" json:"nve_ingress_replication,omitempty"`
	ServiceCarvingVpwsPermanentResult  []*L2VpnEvpnEadServiceInfo `protobuf:"bytes,92,rep,name=service_carving_vpws_permanent_result,json=serviceCarvingVpwsPermanentResult,proto3" json:"service_carving_vpws_permanent_result,omitempty"`
	ServiceCarvingISidelectedResult    []uint32                   `protobuf:"varint,93,rep,packed,name=service_carving_i_sidelected_result,json=serviceCarvingISidelectedResult,proto3" json:"service_carving_i_sidelected_result,omitempty"`
	ServiceCarvingIsidNotElectedResult []uint32                   `protobuf:"varint,94,rep,packed,name=service_carving_isid_not_elected_result,json=serviceCarvingIsidNotElectedResult,proto3" json:"service_carving_isid_not_elected_result,omitempty"`
	ServiceCarvingEviElectedResult     []uint32                   `protobuf:"varint,95,rep,packed,name=service_carving_evi_elected_result,json=serviceCarvingEviElectedResult,proto3" json:"service_carving_evi_elected_result,omitempty"`
	ServiceCarvingEviNotElectedResult  []uint32                   `protobuf:"varint,96,rep,packed,name=service_carving_evi_not_elected_result,json=serviceCarvingEviNotElectedResult,proto3" json:"service_carving_evi_not_elected_result,omitempty"`
	ServiceCarvingVniElectedResult     []uint32                   `protobuf:"varint,97,rep,packed,name=service_carving_vni_elected_result,json=serviceCarvingVniElectedResult,proto3" json:"service_carving_vni_elected_result,omitempty"`
	ServiceCarvingVniNotElectedResult  []uint32                   `protobuf:"varint,98,rep,packed,name=service_carving_vni_not_elected_result,json=serviceCarvingVniNotElectedResult,proto3" json:"service_carving_vni_not_elected_result,omitempty"`
	LocalSplitHorizonGroupLabelValid   bool                       `protobuf:"varint,99,opt,name=local_split_horizon_group_label_valid,json=localSplitHorizonGroupLabelValid,proto3" json:"local_split_horizon_group_label_valid,omitempty"`
	LocalSplitHorizonGroupLabel        uint32                     `protobuf:"varint,100,opt,name=local_split_horizon_group_label,json=localSplitHorizonGroupLabel,proto3" json:"local_split_horizon_group_label,omitempty"`
	RemoteSplitHorizonGroupLabel       []*L2VpnEvpnRemoteShgInfo  `protobuf:"bytes,101,rep,name=remote_split_horizon_group_label,json=remoteSplitHorizonGroupLabel,proto3" json:"remote_split_horizon_group_label,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}                   `json:"-"`
	XXX_unrecognized                   []byte                     `json:"-"`
	XXX_sizecache                      int32                      `json:"-"`
}

func (m *L2VpnEvpnEs) Reset()         { *m = L2VpnEvpnEs{} }
func (m *L2VpnEvpnEs) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEs) ProtoMessage()    {}
func (*L2VpnEvpnEs) Descriptor() ([]byte, []int) {
	return fileDescriptor_95dd9c21af48b1f6, []int{4}
}

func (m *L2VpnEvpnEs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEs.Unmarshal(m, b)
}
func (m *L2VpnEvpnEs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEs.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEs.Merge(m, src)
}
func (m *L2VpnEvpnEs) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEs.Size(m)
}
func (m *L2VpnEvpnEs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEs proto.InternalMessageInfo

func (m *L2VpnEvpnEs) GetEthernetSegmentIdentifier() []uint32 {
	if m != nil {
		return m.EthernetSegmentIdentifier
	}
	return nil
}

func (m *L2VpnEvpnEs) GetEsiType() string {
	if m != nil {
		return m.EsiType
	}
	return ""
}

func (m *L2VpnEvpnEs) GetEsiSystemIdentifier() string {
	if m != nil {
		return m.EsiSystemIdentifier
	}
	return ""
}

func (m *L2VpnEvpnEs) GetEsiPortKey() uint32 {
	if m != nil {
		return m.EsiPortKey
	}
	return 0
}

func (m *L2VpnEvpnEs) GetEsiSystemPriority() uint32 {
	if m != nil {
		return m.EsiSystemPriority
	}
	return 0
}

func (m *L2VpnEvpnEs) GetEthernetSegmentName() string {
	if m != nil {
		return m.EthernetSegmentName
	}
	return ""
}

func (m *L2VpnEvpnEs) GetEthernetSegmentState() string {
	if m != nil {
		return m.EthernetSegmentState
	}
	return ""
}

func (m *L2VpnEvpnEs) GetIfHandle() string {
	if m != nil {
		return m.IfHandle
	}
	return ""
}

func (m *L2VpnEvpnEs) GetMainPortRole() string {
	if m != nil {
		return m.MainPortRole
	}
	return ""
}

func (m *L2VpnEvpnEs) GetMainPortMac() string {
	if m != nil {
		return m.MainPortMac
	}
	return ""
}

func (m *L2VpnEvpnEs) GetNumUpPWs() uint32 {
	if m != nil {
		return m.NumUpPWs
	}
	return 0
}

func (m *L2VpnEvpnEs) GetRouteTarget() string {
	if m != nil {
		return m.RouteTarget
	}
	return ""
}

func (m *L2VpnEvpnEs) GetRtOrigin() string {
	if m != nil {
		return m.RtOrigin
	}
	return ""
}

func (m *L2VpnEvpnEs) GetEsBgpGates() string {
	if m != nil {
		return m.EsBgpGates
	}
	return ""
}

func (m *L2VpnEvpnEs) GetEsL2FibGates() string {
	if m != nil {
		return m.EsL2FibGates
	}
	return ""
}

func (m *L2VpnEvpnEs) GetMacFlushingModeConfig() string {
	if m != nil {
		return m.MacFlushingModeConfig
	}
	return ""
}

func (m *L2VpnEvpnEs) GetLoadBalanceModeConfig() string {
	if m != nil {
		return m.LoadBalanceModeConfig
	}
	return ""
}

func (m *L2VpnEvpnEs) GetLoadBalanceModeIsDefault() bool {
	if m != nil {
		return m.LoadBalanceModeIsDefault
	}
	return false
}

func (m *L2VpnEvpnEs) GetLoadBalanceModeOper() string {
	if m != nil {
		return m.LoadBalanceModeOper
	}
	return ""
}

func (m *L2VpnEvpnEs) GetForceSingleHome() bool {
	if m != nil {
		return m.ForceSingleHome
	}
	return false
}

func (m *L2VpnEvpnEs) GetSourceMacOper() string {
	if m != nil {
		return m.SourceMacOper
	}
	return ""
}

func (m *L2VpnEvpnEs) GetSourceMacOrigin() string {
	if m != nil {
		return m.SourceMacOrigin
	}
	return ""
}

func (m *L2VpnEvpnEs) GetPeeringTimer() uint32 {
	if m != nil {
		return m.PeeringTimer
	}
	return 0
}

func (m *L2VpnEvpnEs) GetPeeringTimerLeft() uint32 {
	if m != nil {
		return m.PeeringTimerLeft
	}
	return 0
}

func (m *L2VpnEvpnEs) GetRecoveryTimer() uint32 {
	if m != nil {
		return m.RecoveryTimer
	}
	return 0
}

func (m *L2VpnEvpnEs) GetRecoveryTimerLeft() uint32 {
	if m != nil {
		return m.RecoveryTimerLeft
	}
	return 0
}

func (m *L2VpnEvpnEs) GetCarvingTimer() uint32 {
	if m != nil {
		return m.CarvingTimer
	}
	return 0
}

func (m *L2VpnEvpnEs) GetCarvingTimerLeft() uint32 {
	if m != nil {
		return m.CarvingTimerLeft
	}
	return 0
}

func (m *L2VpnEvpnEs) GetServiceCarvingMode() string {
	if m != nil {
		return m.ServiceCarvingMode
	}
	return ""
}

func (m *L2VpnEvpnEs) GetPrimaryService() []uint32 {
	if m != nil {
		return m.PrimaryService
	}
	return nil
}

func (m *L2VpnEvpnEs) GetPrimaryServicesInput() string {
	if m != nil {
		return m.PrimaryServicesInput
	}
	return ""
}

func (m *L2VpnEvpnEs) GetSecondaryService() []uint32 {
	if m != nil {
		return m.SecondaryService
	}
	return nil
}

func (m *L2VpnEvpnEs) GetSecondaryServicesInput() string {
	if m != nil {
		return m.SecondaryServicesInput
	}
	return ""
}

func (m *L2VpnEvpnEs) GetNextHop() []*L2VpnEvpnEsNhBuffer {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *L2VpnEvpnEs) GetForwarderPorts() uint32 {
	if m != nil {
		return m.ForwarderPorts
	}
	return 0
}

func (m *L2VpnEvpnEs) GetPermanentForwarderPorts() uint32 {
	if m != nil {
		return m.PermanentForwarderPorts
	}
	return 0
}

func (m *L2VpnEvpnEs) GetElectedForwarderPorts() uint32 {
	if m != nil {
		return m.ElectedForwarderPorts
	}
	return 0
}

func (m *L2VpnEvpnEs) GetNotElectedForwarderPorts() uint32 {
	if m != nil {
		return m.NotElectedForwarderPorts
	}
	return 0
}

func (m *L2VpnEvpnEs) GetNotConfigForwarderPorts() uint32 {
	if m != nil {
		return m.NotConfigForwarderPorts
	}
	return 0
}

func (m *L2VpnEvpnEs) GetMpProtected() bool {
	if m != nil {
		return m.MpProtected
	}
	return false
}

func (m *L2VpnEvpnEs) GetNveAnycastVtep() bool {
	if m != nil {
		return m.NveAnycastVtep
	}
	return false
}

func (m *L2VpnEvpnEs) GetNveIngressReplication() bool {
	if m != nil {
		return m.NveIngressReplication
	}
	return false
}

func (m *L2VpnEvpnEs) GetServiceCarvingVpwsPermanentResult() []*L2VpnEvpnEadServiceInfo {
	if m != nil {
		return m.ServiceCarvingVpwsPermanentResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingISidelectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingISidelectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingIsidNotElectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingIsidNotElectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingEviElectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingEviElectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingEviNotElectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingEviNotElectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingVniElectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingVniElectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetServiceCarvingVniNotElectedResult() []uint32 {
	if m != nil {
		return m.ServiceCarvingVniNotElectedResult
	}
	return nil
}

func (m *L2VpnEvpnEs) GetLocalSplitHorizonGroupLabelValid() bool {
	if m != nil {
		return m.LocalSplitHorizonGroupLabelValid
	}
	return false
}

func (m *L2VpnEvpnEs) GetLocalSplitHorizonGroupLabel() uint32 {
	if m != nil {
		return m.LocalSplitHorizonGroupLabel
	}
	return 0
}

func (m *L2VpnEvpnEs) GetRemoteSplitHorizonGroupLabel() []*L2VpnEvpnRemoteShgInfo {
	if m != nil {
		return m.RemoteSplitHorizonGroupLabel
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnEvpnEs_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ethernet_segments.ethernet_segment.l2vpn_evpn_es_KEYS")
	proto.RegisterType((*L2VpnEvpnEsNhBuffer)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ethernet_segments.ethernet_segment.l2vpn_evpn_es_nh_buffer")
	proto.RegisterType((*L2VpnEvpnEadServiceInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ethernet_segments.ethernet_segment.l2vpn_evpn_ead_service_info")
	proto.RegisterType((*L2VpnEvpnRemoteShgInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ethernet_segments.ethernet_segment.l2vpn_evpn_remote_shg_info")
	proto.RegisterType((*L2VpnEvpnEs)(nil), "cisco_ios_xr_evpn_oper.evpn.active.ethernet_segments.ethernet_segment.l2vpn_evpn_es")
}

func init() { proto.RegisterFile("l2vpn_evpn_es.proto", fileDescriptor_95dd9c21af48b1f6) }

var fileDescriptor_95dd9c21af48b1f6 = []byte{
	// 1413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0x5d, 0x53, 0x14, 0xcd,
	0x19, 0x2d, 0xa2, 0x22, 0x36, 0x2c, 0xe8, 0x20, 0xd2, 0x8a, 0x89, 0xb0, 0x46, 0xa5, 0x92, 0xd4,
	0x56, 0x02, 0xf8, 0x11, 0x8d, 0x26, 0x2a, 0x28, 0x28, 0xc8, 0xba, 0x8b, 0x18, 0xf3, 0x61, 0xa7,
	0x77, 0xe6, 0x99, 0xa5, 0xcb, 0x99, 0xee, 0xae, 0xee, 0xde, 0xc5, 0xcd, 0x6f, 0xc8, 0x7f, 0xc8,
	0x75, 0xee, 0x73, 0xf1, 0xfe, 0xbc, 0xb7, 0xfa, 0xe9, 0xd9, 0x65, 0x77, 0x16, 0xf5, 0xbd, 0xf0,
	0xbd, 0xa1, 0x66, 0xce, 0x73, 0xce, 0xe9, 0xd3, 0x5f, 0xcf, 0x0e, 0x64, 0x3e, 0x5b, 0xeb, 0x6a,
	0xc9, 0x00, 0xff, 0xd8, 0x9a, 0x36, 0xca, 0xa9, 0x68, 0x2b, 0x16, 0x36, 0x56, 0x4c, 0x28, 0xcb,
	0x3e, 0x9b, 0x50, 0x53, 0x1a, 0x4c, 0xcd, 0x3f, 0xd5, 0x78, 0xec, 0x44, 0x17, 0x6a, 0xe0, 0x8e,
	0xc0, 0x48, 0x70, 0xcc, 0x42, 0x3b, 0x07, 0xe9, 0xec, 0x18, 0x52, 0xfd, 0xef, 0x04, 0x89, 0x46,
	0xec, 0xd9, 0xeb, 0xad, 0x0f, 0xcd, 0xe8, 0x16, 0x99, 0x15, 0xd2, 0x81, 0x49, 0x79, 0x0c, 0x4c,
	0xf2, 0x1c, 0xe8, 0xc4, 0xf2, 0xc4, 0xea, 0x85, 0x46, 0x65, 0x80, 0xbe, 0xe1, 0x39, 0x44, 0x11,
	0x39, 0x0b, 0x56, 0xfc, 0x81, 0xfe, 0x02, 0x8b, 0xf8, 0x5c, 0x60, 0x6b, 0xf4, 0xcc, 0x00, 0x5b,
	0x2b, 0xb0, 0x75, 0x7a, 0x76, 0x80, 0xad, 0x17, 0xd8, 0x06, 0x3d, 0x37, 0xc0, 0x36, 0x0a, 0xec,
	0x2e, 0x9d, 0x1c, 0x60, 0x77, 0xab, 0xff, 0x99, 0x20, 0x8b, 0xa3, 0x09, 0xe5, 0x11, 0x6b, 0x75,
	0xd2, 0x14, 0x4c, 0x74, 0x95, 0x4c, 0x49, 0xf8, 0xec, 0xd8, 0x91, 0xd2, 0x45, 0xc0, 0xf3, 0xfe,
	0x7d, 0x5b, 0xe9, 0xe8, 0xd7, 0x64, 0x36, 0x49, 0x59, 0xa2, 0xa4, 0x63, 0xda, 0x40, 0xae, 0x1d,
	0x86, 0x9c, 0x6a, 0xcc, 0x24, 0xe9, 0xa6, 0x92, 0xae, 0x8e, 0x58, 0xb4, 0x48, 0xce, 0x27, 0x29,
	0x73, 0x3d, 0x0d, 0x98, 0xb7, 0xd2, 0x98, 0x4c, 0xd2, 0x83, 0x9e, 0x86, 0xa2, 0xa0, 0x0d, 0xa4,
	0x18, 0x1a, 0x0b, 0x75, 0x03, 0x69, 0xf5, 0x13, 0x59, 0x1a, 0x4e, 0xc3, 0x13, 0x66, 0xc1, 0x74,
	0x45, 0x0c, 0x4c, 0xc8, 0x54, 0x45, 0x0b, 0x64, 0xd2, 0xe3, 0x22, 0xc1, 0x3c, 0x95, 0xc6, 0xb9,
	0xae, 0x96, 0x3b, 0x89, 0x9f, 0x18, 0x0e, 0x52, 0x2c, 0x94, 0x7f, 0x8e, 0x56, 0xc8, 0xcc, 0x60,
	0x3b, 0x1c, 0x6f, 0x17, 0x01, 0xa6, 0xfb, 0xd8, 0x01, 0x6f, 0x57, 0xf7, 0xc8, 0xb5, 0xa1, 0xc1,
	0x0c, 0xe4, 0xca, 0x01, 0xb3, 0x47, 0xed, 0x30, 0xd6, 0x57, 0x66, 0x7f, 0x99, 0x9c, 0xcb, 0x78,
	0x0b, 0x32, 0x1c, 0xb0, 0xd2, 0x08, 0x2f, 0xd5, 0x1f, 0x16, 0x49, 0x65, 0x64, 0x29, 0xa3, 0x27,
	0x64, 0xa9, 0x7c, 0x24, 0x98, 0x48, 0x40, 0x3a, 0x91, 0x0a, 0x30, 0x74, 0x6d, 0xf9, 0xcc, 0x6a,
	0xa5, 0x71, 0xb5, 0x4f, 0x69, 0x06, 0xc6, 0xce, 0x80, 0xe0, 0x23, 0x80, 0x15, 0x61, 0x01, 0xd7,
	0x43, 0x04, 0xb0, 0x02, 0x57, 0x70, 0x8d, 0x2c, 0xf8, 0x92, 0xed, 0x59, 0x07, 0xf9, 0xb0, 0xe9,
	0x06, 0xf2, 0xe6, 0xc1, 0x8a, 0x26, 0xd6, 0x86, 0xec, 0x96, 0xc9, 0x8c, 0xd7, 0x68, 0x65, 0x1c,
	0xfb, 0x04, 0x3d, 0x7a, 0x17, 0xd3, 0x13, 0xb0, 0xa2, 0xae, 0x8c, 0x7b, 0x0d, 0xbd, 0xa8, 0x46,
	0xe6, 0x87, 0x5c, 0xb5, 0x11, 0xca, 0x08, 0xd7, 0xa3, 0xf7, 0x90, 0x78, 0x69, 0xe0, 0x59, 0x2f,
	0x0a, 0x98, 0xa2, 0x3c, 0x41, 0x3c, 0xcf, 0xf7, 0x8b, 0x14, 0xa3, 0x53, 0xc3, 0x53, 0xbd, 0x41,
	0xae, 0x8c, 0x69, 0xac, 0xe3, 0x0e, 0xe8, 0x03, 0x14, 0x5d, 0x2e, 0x89, 0x9a, 0xbe, 0x16, 0x2d,
	0x91, 0x0b, 0x22, 0x65, 0x47, 0x5c, 0x26, 0x19, 0xd0, 0x3f, 0x22, 0x71, 0x4a, 0xa4, 0xdb, 0xf8,
	0xee, 0x4f, 0x63, 0xce, 0x85, 0x0c, 0x33, 0x33, 0x2a, 0x03, 0xfa, 0x10, 0x19, 0x33, 0x1e, 0xf5,
	0x73, 0x6b, 0xa8, 0x0c, 0xa2, 0x2a, 0xa9, 0x9c, 0xb0, 0x72, 0x1e, 0xd3, 0x47, 0x48, 0x9a, 0xee,
	0x93, 0xf6, 0x78, 0x1c, 0xfd, 0x92, 0x4c, 0xcb, 0x4e, 0xce, 0x3a, 0x9a, 0x69, 0x76, 0x6c, 0xe9,
	0x9f, 0x70, 0xe2, 0x53, 0xb2, 0x93, 0xbf, 0xd3, 0xf5, 0xf7, 0xd6, 0x1f, 0x2a, 0xa3, 0x3a, 0x0e,
	0x98, 0xe3, 0xa6, 0x0d, 0x8e, 0x3e, 0x0e, 0x0e, 0x88, 0x1d, 0x20, 0xe4, 0x83, 0x1a, 0xc7, 0x94,
	0x11, 0x6d, 0x21, 0xe9, 0x93, 0x10, 0xd4, 0xb8, 0x7d, 0x7c, 0x0f, 0x3b, 0xc0, 0x5a, 0x6d, 0xcd,
	0xda, 0xdc, 0x81, 0xa5, 0x7f, 0xc6, 0x3a, 0x01, 0xfb, 0xac, 0xad, 0x5f, 0x7a, 0xc4, 0x4f, 0x05,
	0x2c, 0xcb, 0xd6, 0x52, 0xd1, 0x2a, 0x38, 0x7f, 0x09, 0x53, 0x01, 0xbb, 0xeb, 0xc1, 0xc0, 0xba,
	0x4f, 0x68, 0xce, 0x63, 0x96, 0x66, 0x1d, 0x7b, 0x24, 0x64, 0x9b, 0xe5, 0x2a, 0x01, 0x16, 0x2b,
	0x99, 0x8a, 0x36, 0x7d, 0x8a, 0xfc, 0x85, 0x9c, 0xc7, 0x2f, 0x8a, 0xf2, 0x9e, 0x4a, 0xe0, 0x39,
	0x16, 0xbd, 0x30, 0x53, 0x3c, 0x61, 0x2d, 0x9e, 0x71, 0x19, 0xc3, 0x88, 0xf0, 0x59, 0x10, 0xfa,
	0xfa, 0xb3, 0x50, 0x1e, 0x12, 0x3e, 0x21, 0xd7, 0xc7, 0x85, 0xc2, 0xb2, 0x04, 0x52, 0xde, 0xc9,
	0x1c, 0x7d, 0x8e, 0xd7, 0x9f, 0x96, 0xc4, 0x3b, 0x76, 0x33, 0xd4, 0xa3, 0x75, 0x72, 0x65, 0x5c,
	0xef, 0x5b, 0x2a, 0xdd, 0x0c, 0x47, 0xa5, 0xa4, 0xdc, 0xd7, 0x60, 0xa2, 0xdf, 0x90, 0x4b, 0xa9,
	0x32, 0x31, 0x30, 0x2b, 0x64, 0x3b, 0x03, 0x76, 0xa4, 0x72, 0xa0, 0x5b, 0x38, 0xd2, 0x1c, 0x16,
	0x9a, 0x88, 0x6f, 0xab, 0x1c, 0xa2, 0xdb, 0x64, 0xce, 0xaa, 0x8e, 0x27, 0xfb, 0x95, 0x41, 0xe7,
	0x17, 0xa1, 0xa9, 0x06, 0x78, 0x8f, 0xc7, 0x7d, 0xcf, 0x61, 0x5e, 0xd8, 0xa7, 0x97, 0xc8, 0x9c,
	0x3b, 0x61, 0x86, 0xed, 0xba, 0x49, 0x2a, 0x1a, 0xc0, 0xf8, 0x15, 0x76, 0x22, 0x07, 0x43, 0xb7,
	0xf1, 0x3c, 0xcc, 0x14, 0xe0, 0x81, 0xc7, 0xa2, 0xdf, 0x91, 0x68, 0x84, 0xc4, 0x32, 0x48, 0x1d,
	0xdd, 0x41, 0xe6, 0xc5, 0x61, 0xe6, 0x2e, 0xa4, 0xce, 0xb7, 0x7e, 0x03, 0xb1, 0xea, 0x82, 0xe9,
	0x15, 0x9e, 0xaf, 0x90, 0x59, 0xe9, 0xa3, 0xc1, 0xb4, 0x46, 0xe6, 0x47, 0x69, 0xc1, 0xf5, 0x75,
	0xb8, 0x88, 0x23, 0x5c, 0xb4, 0xbd, 0x49, 0x2a, 0x31, 0x37, 0xdd, 0x93, 0xa4, 0xbb, 0x21, 0x69,
	0x01, 0x0e, 0x92, 0x8e, 0x90, 0x82, 0xe7, 0x5e, 0x48, 0x3a, 0xcc, 0x44, 0xcb, 0xdf, 0x93, 0xcb,
	0xfd, 0xde, 0xdb, 0x57, 0xf9, 0x4d, 0xa3, 0x6f, 0x70, 0xad, 0xa2, 0xa2, 0xf6, 0x3c, 0x94, 0xfc,
	0x96, 0x45, 0x77, 0xc8, 0x9c, 0x36, 0x22, 0xe7, 0xa6, 0xd7, 0xef, 0xda, 0x74, 0x1f, 0x5b, 0xdc,
	0x6c, 0x01, 0x37, 0x03, 0xea, 0x5b, 0x40, 0x89, 0x68, 0x99, 0x90, 0xba, 0xe3, 0x68, 0x3d, 0xb4,
	0x80, 0x51, 0xbe, 0xdd, 0xf1, 0xb5, 0xe8, 0xb7, 0xe4, 0x92, 0x85, 0x58, 0xc9, 0x64, 0x78, 0x80,
	0xb7, 0x38, 0xc0, 0xc5, 0x41, 0xa1, 0x3f, 0xc4, 0x03, 0x42, 0xc7, 0xc8, 0xfd, 0x41, 0x1a, 0x38,
	0xc8, 0x95, 0xb2, 0xa6, 0x18, 0xa6, 0x37, 0xd4, 0xf7, 0x9b, 0xcb, 0x67, 0x56, 0xa7, 0xd7, 0x3e,
	0xd6, 0xbe, 0xcb, 0xd7, 0x40, 0xed, 0x0b, 0xbf, 0xb3, 0x27, 0xbf, 0x2b, 0x77, 0x88, 0x3f, 0xd6,
	0xc7, 0xdc, 0x24, 0x60, 0xb0, 0x4d, 0x59, 0x7a, 0x80, 0xbb, 0x33, 0x3b, 0x80, 0x7d, 0xa3, 0xb2,
	0xd1, 0x43, 0x72, 0x55, 0x83, 0xc9, 0xb9, 0xf4, 0xcd, 0xb3, 0x2c, 0x79, 0x87, 0x92, 0xc5, 0x01,
	0xe1, 0xc5, 0xa8, 0xf6, 0x1e, 0x59, 0x84, 0x0c, 0x62, 0x07, 0xc9, 0x98, 0xf2, 0x10, 0x95, 0x0b,
	0x45, 0xb9, 0xa4, 0x7b, 0x4c, 0x96, 0xa4, 0x72, 0xec, 0x4b, 0xda, 0xf7, 0xa8, 0xa5, 0x52, 0xb9,
	0xad, 0x53, 0xe5, 0x8f, 0xc8, 0x35, 0x2f, 0x0f, 0xbd, 0x66, 0x4c, 0xfd, 0xd7, 0x90, 0x59, 0x2a,
	0x17, 0xfa, 0x4d, 0x49, 0xbc, 0x42, 0x66, 0x72, 0xcd, 0xfc, 0xa7, 0x19, 0x5a, 0xd3, 0x0f, 0xd8,
	0x03, 0xa6, 0x73, 0x5d, 0xef, 0x43, 0xd1, 0x2a, 0xb9, 0x28, 0xbb, 0xc0, 0xb8, 0xec, 0xc5, 0xdc,
	0x3a, 0xd6, 0x75, 0xa0, 0xe9, 0xdf, 0x90, 0x36, 0x2b, 0xbb, 0xf0, 0x34, 0xc0, 0x87, 0x0e, 0xb4,
	0x5f, 0x00, 0xcf, 0x14, 0xb2, 0x6d, 0xc0, 0x5a, 0x66, 0x40, 0x67, 0x22, 0xe6, 0x4e, 0x28, 0x49,
	0xff, 0x8e, 0x82, 0x05, 0xd9, 0x85, 0x9d, 0x50, 0x6d, 0x9c, 0x14, 0xa3, 0xff, 0x4f, 0x90, 0x5b,
	0xe5, 0x1b, 0xd1, 0xd5, 0xc7, 0x96, 0x9d, 0x6c, 0x85, 0x01, 0xeb, 0x9b, 0xe1, 0x3f, 0xf0, 0xd8,
	0xb4, 0x7e, 0x86, 0x63, 0x53, 0xfa, 0x20, 0x6a, 0xac, 0x8c, 0x5e, 0xc3, 0x43, 0x7d, 0x6c, 0xeb,
	0xfd, 0x34, 0x0d, 0x0c, 0x13, 0xed, 0x92, 0x9b, 0xe5, 0xd4, 0x82, 0x59, 0x91, 0xf4, 0x37, 0xb2,
	0xc8, 0xfc, 0x4f, 0xbc, 0x48, 0x37, 0x46, 0xfd, 0x76, 0x9a, 0x03, 0x5e, 0xe1, 0xd6, 0x24, 0x77,
	0xc6, 0xdc, 0xac, 0x48, 0xd8, 0xf0, 0xd1, 0x28, 0x1c, 0x3f, 0xa2, 0x63, 0xb5, 0xe4, 0x68, 0x45,
	0xf2, 0x66, 0x70, 0x46, 0x0a, 0xd3, 0x57, 0xa4, 0x5a, 0x36, 0x85, 0xae, 0x28, 0xfb, 0x31, 0xf4,
	0xfb, 0xd5, 0xa8, 0xdf, 0x56, 0x57, 0x8c, 0x7a, 0xbd, 0x25, 0xb7, 0x4f, 0xf3, 0x3a, 0x25, 0xdf,
	0xbf, 0xd0, 0x6f, 0x65, 0xcc, 0xef, 0xa7, 0xc4, 0xeb, 0xca, 0xb1, 0x78, 0xfc, 0xb4, 0x78, 0x87,
	0xf2, 0xdb, 0xf1, 0xbc, 0xd7, 0x29, 0xf1, 0x5a, 0xa7, 0xc5, 0x3b, 0x94, 0xe3, 0xf1, 0xf6, 0xc9,
	0xad, 0x4c, 0xc5, 0x3c, 0x63, 0x56, 0x67, 0xc2, 0xf7, 0x2d, 0x23, 0xfe, 0xad, 0x24, 0x6b, 0x1b,
	0xd5, 0xd1, 0x0c, 0x3f, 0x4d, 0x59, 0x97, 0x67, 0x22, 0xa1, 0x31, 0x9e, 0xee, 0x65, 0x24, 0x37,
	0x3d, 0x77, 0x3b, 0x50, 0x5f, 0x7a, 0xe6, 0xae, 0x27, 0x1e, 0x7a, 0x5e, 0xb4, 0x49, 0x6e, 0x7c,
	0xc3, 0x90, 0x26, 0x78, 0x5f, 0x97, 0xbe, 0x62, 0x15, 0xfd, 0x6f, 0x82, 0x2c, 0xf7, 0xbf, 0xa9,
	0xbf, 0xe8, 0x03, 0x78, 0x53, 0xf8, 0xf7, 0xbf, 0x29, 0xa5, 0xaf, 0xf9, 0xc6, 0xf5, 0x00, 0x9c,
	0x9e, 0xb5, 0x35, 0x89, 0xff, 0xf5, 0xad, 0xff, 0x18, 0x00, 0x00, 0xff, 0xff, 0x70, 0x8b, 0xc9,
	0xe0, 0x0c, 0x0e, 0x00, 0x00,
}