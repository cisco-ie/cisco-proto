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

package cisco_ios_xr_evpn_oper_evpn_nodes_node_ethernet_segments_ethernet_segment

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
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Esi1                 string   `protobuf:"bytes,3,opt,name=esi1,proto3" json:"esi1,omitempty"`
	Esi2                 string   `protobuf:"bytes,4,opt,name=esi2,proto3" json:"esi2,omitempty"`
	Esi3                 string   `protobuf:"bytes,5,opt,name=esi3,proto3" json:"esi3,omitempty"`
	Esi4                 string   `protobuf:"bytes,6,opt,name=esi4,proto3" json:"esi4,omitempty"`
	Esi5                 string   `protobuf:"bytes,7,opt,name=esi5,proto3" json:"esi5,omitempty"`
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

func (m *L2VpnEvpnEs_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

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
	EthernetSegmentState               uint32                     `protobuf:"varint,56,opt,name=ethernet_segment_state,json=ethernetSegmentState,proto3" json:"ethernet_segment_state,omitempty"`
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

func (m *L2VpnEvpnEs) GetEthernetSegmentState() uint32 {
	if m != nil {
		return m.EthernetSegmentState
	}
	return 0
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
	proto.RegisterType((*L2VpnEvpnEs_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ethernet_segments.ethernet_segment.l2vpn_evpn_es_KEYS")
	proto.RegisterType((*L2VpnEvpnEsNhBuffer)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ethernet_segments.ethernet_segment.l2vpn_evpn_es_nh_buffer")
	proto.RegisterType((*L2VpnEvpnEadServiceInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ethernet_segments.ethernet_segment.l2vpn_evpn_ead_service_info")
	proto.RegisterType((*L2VpnEvpnRemoteShgInfo)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ethernet_segments.ethernet_segment.l2vpn_evpn_remote_shg_info")
	proto.RegisterType((*L2VpnEvpnEs)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.ethernet_segments.ethernet_segment.l2vpn_evpn_es")
}

func init() { proto.RegisterFile("l2vpn_evpn_es.proto", fileDescriptor_95dd9c21af48b1f6) }

var fileDescriptor_95dd9c21af48b1f6 = []byte{
	// 1428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xed, 0x52, 0x14, 0x47,
	0x1b, 0x2d, 0x5e, 0x15, 0xb1, 0x61, 0x41, 0x07, 0x90, 0x56, 0x7c, 0x5f, 0x61, 0x7d, 0x55, 0x2a,
	0x49, 0x6d, 0x25, 0x80, 0x1f, 0xd1, 0x68, 0xa2, 0x82, 0xb2, 0x0a, 0xb2, 0xee, 0x22, 0xc6, 0x7c,
	0x75, 0x66, 0x67, 0x9e, 0x59, 0xba, 0x9c, 0xe9, 0xee, 0xea, 0xee, 0x5d, 0xdc, 0x54, 0xe5, 0x0e,
	0x72, 0x2b, 0xf9, 0xef, 0x05, 0xe4, 0xc2, 0x52, 0xfd, 0xf4, 0xec, 0xb2, 0x3b, 0x8b, 0x9a, 0x1f,
	0xf1, 0x0f, 0x35, 0x73, 0x9e, 0x73, 0x4e, 0x9f, 0xfe, 0x7a, 0x76, 0x20, 0xb3, 0xe9, 0x6a, 0x47,
	0x09, 0x06, 0xf8, 0xc7, 0x54, 0x94, 0x96, 0x56, 0x06, 0xd5, 0x88, 0x9b, 0x48, 0x32, 0x2e, 0x0d,
	0x7b, 0xab, 0x7d, 0x4d, 0x2a, 0xd0, 0x15, 0xf7, 0x54, 0x11, 0x32, 0x06, 0x83, 0x7f, 0x2b, 0x60,
	0x0f, 0x40, 0x0b, 0xb0, 0xcc, 0x40, 0x2b, 0x03, 0x61, 0xcd, 0x08, 0x52, 0x7e, 0x37, 0x46, 0x82,
	0xa1, 0x21, 0xd8, 0xb3, 0xcd, 0xd7, 0x8d, 0x60, 0x81, 0x9c, 0x76, 0x0e, 0x8c, 0xc7, 0x74, 0x6c,
	0x69, 0x6c, 0xe5, 0x4c, 0x7d, 0xdc, 0xbd, 0x56, 0xe3, 0xe0, 0x2a, 0x99, 0xe6, 0xc2, 0x82, 0x4e,
	0xc2, 0x08, 0x98, 0x08, 0x33, 0xa0, 0xff, 0xc1, 0x7a, 0xa9, 0x8f, 0x3e, 0x0f, 0x33, 0x08, 0x02,
	0x72, 0x12, 0x0c, 0xff, 0x8a, 0x9e, 0xc0, 0x22, 0x3e, 0xe7, 0xd8, 0x2a, 0x3d, 0xd9, 0xc7, 0x56,
	0x73, 0x6c, 0x8d, 0x9e, 0xea, 0x63, 0x6b, 0x39, 0xb6, 0x4e, 0xc7, 0xfb, 0xd8, 0x7a, 0x8e, 0xdd,
	0xa0, 0xa7, 0xfb, 0xd8, 0x8d, 0xf2, 0x1f, 0x63, 0x64, 0x61, 0x38, 0xba, 0x38, 0x60, 0xcd, 0x76,
	0x92, 0x80, 0x0e, 0x2e, 0x90, 0x09, 0x01, 0x6f, 0x2d, 0x3b, 0x90, 0x2a, 0x9f, 0xc0, 0x69, 0xf7,
	0xbe, 0x25, 0x55, 0xf0, 0x7f, 0x32, 0x1d, 0x27, 0x2c, 0x96, 0xc2, 0x32, 0xa5, 0x21, 0x53, 0x16,
	0x67, 0x30, 0x51, 0x9f, 0x8a, 0x93, 0x0d, 0x29, 0x6c, 0x0d, 0x31, 0xb7, 0x00, 0x71, 0xc2, 0x6c,
	0x57, 0x01, 0xce, 0xa1, 0x54, 0x1f, 0x8f, 0x93, 0xbd, 0xae, 0x82, 0xbc, 0xa0, 0x34, 0x24, 0x38,
	0x11, 0x2c, 0xd4, 0x34, 0x24, 0xe5, 0x37, 0x64, 0x71, 0x30, 0x4d, 0x18, 0x33, 0x03, 0xba, 0xc3,
	0x23, 0x60, 0x5c, 0x24, 0x32, 0x98, 0x27, 0xe3, 0x0e, 0xcf, 0x17, 0xb4, 0x54, 0x3f, 0xd5, 0x51,
	0xa2, 0x1a, 0xbb, 0x89, 0xe1, 0x20, 0x7e, 0x15, 0xf1, 0x39, 0x58, 0x26, 0x53, 0xfd, 0x7d, 0xb2,
	0x61, 0x2b, 0x0f, 0x30, 0xd9, 0xc3, 0xf6, 0xc2, 0x56, 0x79, 0x87, 0x5c, 0x1c, 0x18, 0x4c, 0x43,
	0x26, 0x2d, 0x30, 0x73, 0xd0, 0xf2, 0x63, 0x7d, 0x60, 0xf6, 0x73, 0xe4, 0x54, 0x1a, 0x36, 0x21,
	0xc5, 0x01, 0x4b, 0x75, 0xff, 0x52, 0xfe, 0x6b, 0x81, 0x94, 0x86, 0x96, 0x32, 0xb8, 0x4f, 0x16,
	0x8b, 0x67, 0x85, 0xf1, 0x18, 0x84, 0xe5, 0x09, 0x07, 0x4d, 0x57, 0x97, 0x4e, 0xac, 0x94, 0xea,
	0x17, 0x7a, 0x94, 0x86, 0x67, 0x54, 0xfb, 0x04, 0x17, 0x01, 0x0c, 0xf7, 0x0b, 0xb8, 0xe6, 0x23,
	0x80, 0xe1, 0xb8, 0x82, 0xab, 0x64, 0xde, 0x95, 0x4c, 0xd7, 0x58, 0xc8, 0x06, 0x4d, 0xd7, 0x91,
	0x37, 0x0b, 0x86, 0x37, 0xb0, 0x36, 0x60, 0xb7, 0x44, 0xa6, 0x9c, 0x46, 0x49, 0x6d, 0xd9, 0x1b,
	0xe8, 0xd2, 0x1b, 0x98, 0x9e, 0x80, 0xe1, 0x35, 0xa9, 0xed, 0x33, 0xe8, 0x06, 0x15, 0x32, 0x3b,
	0xe0, 0xaa, 0x34, 0x97, 0x9a, 0xdb, 0x2e, 0xbd, 0x89, 0xc4, 0x73, 0x7d, 0xcf, 0x5a, 0x5e, 0xc0,
	0x14, 0xc5, 0x09, 0xe2, 0x79, 0xbe, 0x95, 0xa7, 0x18, 0x9e, 0x1a, 0x9e, 0xea, 0x75, 0x72, 0x7e,
	0x44, 0x63, 0x6c, 0x68, 0x81, 0xde, 0xc6, 0x61, 0xe6, 0x0a, 0xa2, 0x86, 0xab, 0x05, 0x8b, 0xe4,
	0x0c, 0x4f, 0xd8, 0x41, 0x28, 0xe2, 0x14, 0xe8, 0xd7, 0xe8, 0x3e, 0xc1, 0x93, 0x2d, 0x7c, 0x77,
	0xa7, 0x31, 0x0b, 0xb9, 0xf0, 0x33, 0xd3, 0x32, 0x05, 0x7a, 0x07, 0x19, 0x53, 0x0e, 0x75, 0x73,
	0xab, 0xcb, 0x14, 0x82, 0x32, 0x29, 0x1d, 0xb1, 0xb2, 0x30, 0xa2, 0x77, 0x91, 0x34, 0xd9, 0x23,
	0xed, 0x84, 0x51, 0xf0, 0x5f, 0x32, 0x29, 0xda, 0x19, 0x6b, 0x2b, 0xa6, 0xd8, 0xa1, 0xa1, 0xdf,
	0x60, 0xa2, 0x09, 0xd1, 0xce, 0x5e, 0xaa, 0xda, 0x2b, 0xe3, 0x0e, 0x95, 0x96, 0x6d, 0x0b, 0xcc,
	0x86, 0xba, 0x05, 0x96, 0xde, 0xf3, 0x0e, 0x88, 0xed, 0x21, 0xe4, 0x82, 0x6a, 0xcb, 0xa4, 0xe6,
	0x2d, 0x2e, 0xe8, 0x7d, 0x1f, 0x54, 0xdb, 0x5d, 0x7c, 0xf7, 0x3b, 0xc0, 0x9a, 0x2d, 0xc5, 0x5a,
	0xa1, 0x05, 0x43, 0xbf, 0xc5, 0x3a, 0x01, 0xf3, 0xb0, 0xa5, 0x9e, 0x38, 0xc4, 0x4d, 0x05, 0x0c,
	0x4b, 0x57, 0x13, 0xde, 0xcc, 0x39, 0xdf, 0xf9, 0xa9, 0x80, 0xd9, 0x76, 0xa0, 0x67, 0xdd, 0x22,
	0x34, 0x0b, 0x23, 0x96, 0xa4, 0x6d, 0x73, 0xc0, 0x45, 0x8b, 0x65, 0xae, 0xcd, 0x44, 0x52, 0x24,
	0xbc, 0x45, 0x1f, 0x20, 0x7f, 0x3e, 0x0b, 0xa3, 0xc7, 0x79, 0x79, 0x47, 0xc6, 0xf0, 0x08, 0x8b,
	0x4e, 0x98, 0xca, 0x30, 0x66, 0xcd, 0x30, 0x0d, 0x45, 0x04, 0x43, 0xc2, 0x87, 0x5e, 0xe8, 0xea,
	0x0f, 0x7d, 0x79, 0x40, 0x78, 0x9f, 0x5c, 0x1a, 0x15, 0x72, 0xc3, 0x62, 0x48, 0xc2, 0x76, 0x6a,
	0xe9, 0x23, 0xbc, 0xfe, 0xb4, 0x20, 0xae, 0x9a, 0x0d, 0x5f, 0x0f, 0xd6, 0xc8, 0xf9, 0x51, 0xbd,
	0xeb, 0xb7, 0x74, 0xc3, 0x1f, 0x95, 0x82, 0x72, 0x57, 0x81, 0x0e, 0x3e, 0x23, 0xe7, 0x12, 0xa9,
	0x23, 0x60, 0x86, 0x8b, 0x56, 0x0a, 0xec, 0x40, 0x66, 0x40, 0x37, 0x71, 0xa4, 0x19, 0x2c, 0x34,
	0x10, 0xdf, 0x92, 0x19, 0x04, 0xd7, 0xc8, 0x8c, 0x91, 0x6d, 0x47, 0x76, 0x2b, 0x83, 0xce, 0x8f,
	0x7d, 0x53, 0xf5, 0xf0, 0x4e, 0x18, 0xf5, 0x3c, 0x07, 0x79, 0x7e, 0x9f, 0x9e, 0x20, 0x73, 0xe6,
	0x88, 0xe9, 0xb7, 0xeb, 0x0a, 0x29, 0x29, 0x00, 0xed, 0x56, 0xd8, 0xf2, 0x0c, 0x34, 0xdd, 0xc2,
	0xf3, 0x30, 0x95, 0x83, 0x7b, 0x0e, 0x0b, 0xbe, 0x20, 0xc1, 0x10, 0x89, 0xa5, 0x90, 0x58, 0x5a,
	0x45, 0xe6, 0xd9, 0x41, 0xe6, 0x36, 0x24, 0xd6, 0xb5, 0x7e, 0x0d, 0x91, 0xec, 0x80, 0xee, 0xe6,
	0x9e, 0x4f, 0x91, 0x59, 0xea, 0xa1, 0xde, 0xb4, 0x42, 0x66, 0x87, 0x69, 0xde, 0xf5, 0x99, 0xbf,
	0x88, 0x43, 0x5c, 0xb4, 0xbd, 0x42, 0x4a, 0x51, 0xa8, 0x3b, 0x47, 0x49, 0xb7, 0x7d, 0xd2, 0x1c,
	0xec, 0x27, 0x1d, 0x22, 0x79, 0xcf, 0x1d, 0x9f, 0x74, 0x90, 0x89, 0x96, 0x5f, 0x92, 0xb9, 0x5e,
	0xef, 0xed, 0xa9, 0xdc, 0xa6, 0xd1, 0xe7, 0xb8, 0x56, 0x41, 0x5e, 0x7b, 0xe4, 0x4b, 0x6e, 0xcb,
	0x82, 0xeb, 0x64, 0x46, 0x69, 0x9e, 0x85, 0xba, 0xdb, 0xeb, 0xda, 0x74, 0x17, 0x5b, 0xdc, 0x74,
	0x0e, 0x37, 0x3c, 0xea, 0x5a, 0x40, 0x81, 0x68, 0x18, 0x17, 0xaa, 0x6d, 0x69, 0x0d, 0xcd, 0xe7,
	0x86, 0xf9, 0xa6, 0xea, 0x6a, 0xc1, 0xe7, 0xe4, 0x9c, 0x81, 0x48, 0x8a, 0x78, 0x70, 0x80, 0x17,
	0x38, 0xc0, 0xd9, 0x7e, 0xa1, 0x37, 0xc4, 0x6d, 0x42, 0x47, 0xc8, 0xbd, 0x41, 0xea, 0x38, 0xc8,
	0xf9, 0xa2, 0x26, 0x1f, 0xe6, 0xf7, 0x81, 0xbe, 0xdf, 0x58, 0x3a, 0xb1, 0x32, 0xb9, 0xda, 0xac,
	0xfc, 0x6b, 0x9f, 0x0a, 0x95, 0xf7, 0xfc, 0xd6, 0x1e, 0xfd, 0xb6, 0x5c, 0x27, 0xee, 0x68, 0x1f,
	0x86, 0x3a, 0x06, 0x8d, 0xad, 0xca, 0xd0, 0x3d, 0xdc, 0xa1, 0xe9, 0x3e, 0xec, 0x9a, 0x95, 0x09,
	0xee, 0x90, 0x0b, 0x0a, 0x74, 0x16, 0x0a, 0xd7, 0x40, 0x8b, 0x92, 0x97, 0x28, 0x59, 0xe8, 0x13,
	0x1e, 0x0f, 0x6b, 0x6f, 0x92, 0x05, 0x48, 0x21, 0xb2, 0x10, 0x8f, 0x28, 0xf7, 0x51, 0x39, 0x9f,
	0x97, 0x0b, 0xba, 0x7b, 0x64, 0x51, 0x48, 0xcb, 0xde, 0xa7, 0x7d, 0x85, 0x5a, 0x2a, 0xa4, 0xdd,
	0x3c, 0x56, 0x7e, 0x97, 0x5c, 0x74, 0x72, 0xdf, 0x6f, 0x46, 0xd4, 0xdf, 0xfb, 0xcc, 0x42, 0x5a,
	0xdf, 0x73, 0x0a, 0xe2, 0x65, 0x32, 0x95, 0x29, 0xe6, 0xbe, 0xdd, 0xd0, 0x9a, 0xbe, 0xc6, 0x3e,
	0x30, 0x99, 0xa9, 0x5a, 0x0f, 0x0a, 0x56, 0xc8, 0x59, 0xd1, 0x01, 0x16, 0x8a, 0x6e, 0x14, 0x1a,
	0xcb, 0x3a, 0x16, 0x14, 0xfd, 0x01, 0x69, 0xd3, 0xa2, 0x03, 0x0f, 0x3c, 0xbc, 0x6f, 0x41, 0xb9,
	0x05, 0x70, 0x4c, 0x2e, 0x5a, 0x1a, 0x8c, 0x61, 0x1a, 0x54, 0xca, 0xa3, 0xd0, 0x72, 0x29, 0xe8,
	0x8f, 0x28, 0x98, 0x17, 0x1d, 0xa8, 0xfa, 0x6a, 0xfd, 0xa8, 0x18, 0xbc, 0x1b, 0x23, 0x57, 0x8b,
	0xb7, 0xa2, 0xa3, 0x0e, 0x0d, 0x3b, 0xda, 0x0a, 0x0d, 0xc6, 0x35, 0xc4, 0x9f, 0xf0, 0xe8, 0x24,
	0x9f, 0xe8, 0xe8, 0x14, 0x3e, 0x8c, 0xea, 0xcb, 0xc3, 0xd7, 0x71, 0x5f, 0x1d, 0x9a, 0x5a, 0x2f,
	0x51, 0x1d, 0x03, 0x05, 0xdb, 0xe4, 0x4a, 0x31, 0x39, 0x67, 0x86, 0xc7, 0xbd, 0xcd, 0xcc, 0x73,
	0xff, 0x8c, 0x17, 0xea, 0xf2, 0xb0, 0x5f, 0xb5, 0xd1, 0xe7, 0xe5, 0x6e, 0x0d, 0x72, 0x7d, 0xc4,
	0xcd, 0xf0, 0x98, 0x0d, 0x1e, 0x8f, 0xdc, 0xf1, 0x17, 0x74, 0x2c, 0x17, 0x1c, 0x0d, 0x8f, 0x9f,
	0xf7, 0xcf, 0x49, 0x6e, 0xfa, 0x94, 0x94, 0x8b, 0xa6, 0xd0, 0xe1, 0x45, 0x3f, 0x86, 0x7e, 0xff,
	0x1b, 0xf6, 0xdb, 0xec, 0xf0, 0x61, 0xaf, 0x17, 0xe4, 0xda, 0x71, 0x5e, 0xc7, 0xe4, 0xfb, 0x15,
	0xfd, 0x96, 0x47, 0xfc, 0xfe, 0x49, 0xbc, 0x8e, 0x18, 0x89, 0x17, 0x1e, 0x17, 0x6f, 0x5f, 0x7c,
	0x3c, 0x9e, 0xf3, 0x3a, 0x26, 0x5e, 0xf3, 0xb8, 0x78, 0xfb, 0x62, 0x34, 0xde, 0x2e, 0xb9, 0x9a,
	0xca, 0x28, 0x4c, 0x99, 0x51, 0x29, 0x77, 0xfd, 0x4b, 0xf3, 0xdf, 0xa4, 0x60, 0x2d, 0x2d, 0xdb,
	0x8a, 0xe1, 0x27, 0x2a, 0xeb, 0x84, 0x29, 0x8f, 0x69, 0x84, 0x27, 0x7c, 0x09, 0xc9, 0x0d, 0xc7,
	0xdd, 0xf2, 0xd4, 0x27, 0x8e, 0xb9, 0xed, 0x88, 0xfb, 0x8e, 0x17, 0x6c, 0x90, 0xcb, 0x1f, 0x31,
	0xa4, 0x31, 0xde, 0xd9, 0xc5, 0x0f, 0x58, 0x05, 0x7f, 0x8e, 0x91, 0xa5, 0xde, 0xb7, 0xf5, 0x7b,
	0x7d, 0x00, 0x6f, 0x0b, 0x7c, 0x9a, 0xdb, 0x52, 0xf8, 0xb2, 0xaf, 0x5f, 0xf2, 0xc0, 0xf1, 0x79,
	0x9b, 0xe3, 0xf8, 0xef, 0xe1, 0xda, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x70, 0x6b, 0x4d, 0x03,
	0x35, 0x0e, 0x00, 0x00,
}
