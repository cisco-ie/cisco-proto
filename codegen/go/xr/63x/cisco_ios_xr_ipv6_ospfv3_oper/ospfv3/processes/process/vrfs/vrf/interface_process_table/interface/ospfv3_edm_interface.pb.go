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
// source: ospfv3_edm_interface.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_interface_process_table_interface

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

type Ospfv3EdmInterface_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmInterface_KEYS) Reset()         { *m = Ospfv3EdmInterface_KEYS{} }
func (m *Ospfv3EdmInterface_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterface_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmInterface_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd988687fd2d89, []int{0}
}

func (m *Ospfv3EdmInterface_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmInterface_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmInterface_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmInterface_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmInterface_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmInterface_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmInterface_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmInterface_KEYS.Size(m)
}
func (m *Ospfv3EdmInterface_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmInterface_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmInterface_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmInterface_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmInterface_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3EdmInterface_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmInterfaceNbr struct {
	InterfaceNeighborId   string   `protobuf:"bytes,1,opt,name=interface_neighbor_id,json=interfaceNeighborId,proto3" json:"interface_neighbor_id,omitempty"`
	InterfaceNeighborCost uint32   `protobuf:"varint,2,opt,name=interface_neighbor_cost,json=interfaceNeighborCost,proto3" json:"interface_neighbor_cost,omitempty"`
	IsNeighborDr          bool     `protobuf:"varint,3,opt,name=is_neighbor_dr,json=isNeighborDr,proto3" json:"is_neighbor_dr,omitempty"`
	IsNeighborBdr         bool     `protobuf:"varint,4,opt,name=is_neighbor_bdr,json=isNeighborBdr,proto3" json:"is_neighbor_bdr,omitempty"`
	IsHelloSuppressed     bool     `protobuf:"varint,5,opt,name=is_hello_suppressed,json=isHelloSuppressed,proto3" json:"is_hello_suppressed,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ospfv3EdmInterfaceNbr) Reset()         { *m = Ospfv3EdmInterfaceNbr{} }
func (m *Ospfv3EdmInterfaceNbr) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceNbr) ProtoMessage()    {}
func (*Ospfv3EdmInterfaceNbr) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd988687fd2d89, []int{1}
}

func (m *Ospfv3EdmInterfaceNbr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmInterfaceNbr.Unmarshal(m, b)
}
func (m *Ospfv3EdmInterfaceNbr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmInterfaceNbr.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmInterfaceNbr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmInterfaceNbr.Merge(m, src)
}
func (m *Ospfv3EdmInterfaceNbr) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmInterfaceNbr.Size(m)
}
func (m *Ospfv3EdmInterfaceNbr) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmInterfaceNbr.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmInterfaceNbr proto.InternalMessageInfo

func (m *Ospfv3EdmInterfaceNbr) GetInterfaceNeighborId() string {
	if m != nil {
		return m.InterfaceNeighborId
	}
	return ""
}

func (m *Ospfv3EdmInterfaceNbr) GetInterfaceNeighborCost() uint32 {
	if m != nil {
		return m.InterfaceNeighborCost
	}
	return 0
}

func (m *Ospfv3EdmInterfaceNbr) GetIsNeighborDr() bool {
	if m != nil {
		return m.IsNeighborDr
	}
	return false
}

func (m *Ospfv3EdmInterfaceNbr) GetIsNeighborBdr() bool {
	if m != nil {
		return m.IsNeighborBdr
	}
	return false
}

func (m *Ospfv3EdmInterfaceNbr) GetIsHelloSuppressed() bool {
	if m != nil {
		return m.IsHelloSuppressed
	}
	return false
}

type Ospfv3EdmInterfaceUp struct {
	WaitTime                    uint32   `protobuf:"varint,1,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	InterfaceAreaFloodIndex     uint32   `protobuf:"varint,2,opt,name=interface_area_flood_index,json=interfaceAreaFloodIndex,proto3" json:"interface_area_flood_index,omitempty"`
	InterfaceAsFloodIndex       uint32   `protobuf:"varint,3,opt,name=interface_as_flood_index,json=interfaceAsFloodIndex,proto3" json:"interface_as_flood_index,omitempty"`
	InterfaceLinkFloodIndex     uint32   `protobuf:"varint,4,opt,name=interface_link_flood_index,json=interfaceLinkFloodIndex,proto3" json:"interface_link_flood_index,omitempty"`
	FloodQueueLength            uint32   `protobuf:"varint,5,opt,name=flood_queue_length,json=floodQueueLength,proto3" json:"flood_queue_length,omitempty"`
	InterfaceAreaNextFlood      uint32   `protobuf:"varint,6,opt,name=interface_area_next_flood,json=interfaceAreaNextFlood,proto3" json:"interface_area_next_flood,omitempty"`
	InterfaceAreaNextFloodIndex uint32   `protobuf:"varint,7,opt,name=interface_area_next_flood_index,json=interfaceAreaNextFloodIndex,proto3" json:"interface_area_next_flood_index,omitempty"`
	InterfaceAsNextFlood        uint32   `protobuf:"varint,8,opt,name=interface_as_next_flood,json=interfaceAsNextFlood,proto3" json:"interface_as_next_flood,omitempty"`
	InterfaceAsNextFloodIndex   uint32   `protobuf:"varint,9,opt,name=interface_as_next_flood_index,json=interfaceAsNextFloodIndex,proto3" json:"interface_as_next_flood_index,omitempty"`
	InterfaceLinkNextFlood      uint32   `protobuf:"varint,10,opt,name=interface_link_next_flood,json=interfaceLinkNextFlood,proto3" json:"interface_link_next_flood,omitempty"`
	InterfaceLinkNextIndex      uint32   `protobuf:"varint,11,opt,name=interface_link_next_index,json=interfaceLinkNextIndex,proto3" json:"interface_link_next_index,omitempty"`
	FloodScanLength             uint32   `protobuf:"varint,12,opt,name=flood_scan_length,json=floodScanLength,proto3" json:"flood_scan_length,omitempty"`
	MaximumFloodLength          uint32   `protobuf:"varint,13,opt,name=maximum_flood_length,json=maximumFloodLength,proto3" json:"maximum_flood_length,omitempty"`
	LastFloodTime               uint32   `protobuf:"varint,14,opt,name=last_flood_time,json=lastFloodTime,proto3" json:"last_flood_time,omitempty"`
	MaximumFloodTime            uint32   `protobuf:"varint,15,opt,name=maximum_flood_time,json=maximumFloodTime,proto3" json:"maximum_flood_time,omitempty"`
	InterfaceFloodPacingTimer   uint32   `protobuf:"varint,16,opt,name=interface_flood_pacing_timer,json=interfaceFloodPacingTimer,proto3" json:"interface_flood_pacing_timer,omitempty"`
	InterfaceNeighbors          uint32   `protobuf:"varint,17,opt,name=interface_neighbors,json=interfaceNeighbors,proto3" json:"interface_neighbors,omitempty"`
	SuppressedHellos            uint32   `protobuf:"varint,18,opt,name=suppressed_hellos,json=suppressedHellos,proto3" json:"suppressed_hellos,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *Ospfv3EdmInterfaceUp) Reset()         { *m = Ospfv3EdmInterfaceUp{} }
func (m *Ospfv3EdmInterfaceUp) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceUp) ProtoMessage()    {}
func (*Ospfv3EdmInterfaceUp) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd988687fd2d89, []int{2}
}

func (m *Ospfv3EdmInterfaceUp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmInterfaceUp.Unmarshal(m, b)
}
func (m *Ospfv3EdmInterfaceUp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmInterfaceUp.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmInterfaceUp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmInterfaceUp.Merge(m, src)
}
func (m *Ospfv3EdmInterfaceUp) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmInterfaceUp.Size(m)
}
func (m *Ospfv3EdmInterfaceUp) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmInterfaceUp.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmInterfaceUp proto.InternalMessageInfo

func (m *Ospfv3EdmInterfaceUp) GetWaitTime() uint32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAreaFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAsFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceLinkFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetFloodQueueLength() uint32 {
	if m != nil {
		return m.FloodQueueLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaNextFlood() uint32 {
	if m != nil {
		return m.InterfaceAreaNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAreaNextFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAreaNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsNextFlood() uint32 {
	if m != nil {
		return m.InterfaceAsNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceAsNextFloodIndex() uint32 {
	if m != nil {
		return m.InterfaceAsNextFloodIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkNextFlood() uint32 {
	if m != nil {
		return m.InterfaceLinkNextFlood
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceLinkNextIndex() uint32 {
	if m != nil {
		return m.InterfaceLinkNextIndex
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetFloodScanLength() uint32 {
	if m != nil {
		return m.FloodScanLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetMaximumFloodLength() uint32 {
	if m != nil {
		return m.MaximumFloodLength
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetLastFloodTime() uint32 {
	if m != nil {
		return m.LastFloodTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetMaximumFloodTime() uint32 {
	if m != nil {
		return m.MaximumFloodTime
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceFloodPacingTimer() uint32 {
	if m != nil {
		return m.InterfaceFloodPacingTimer
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetInterfaceNeighbors() uint32 {
	if m != nil {
		return m.InterfaceNeighbors
	}
	return 0
}

func (m *Ospfv3EdmInterfaceUp) GetSuppressedHellos() uint32 {
	if m != nil {
		return m.SuppressedHellos
	}
	return 0
}

type Ospfv3EdmInterfaceBfd struct {
	BfdIntfEnableMode      uint32   `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	BfdInterval            uint32   `protobuf:"varint,2,opt,name=bfd_interval,json=bfdInterval,proto3" json:"bfd_interval,omitempty"`
	BfdDetectionMultiplier uint32   `protobuf:"varint,3,opt,name=bfd_detection_multiplier,json=bfdDetectionMultiplier,proto3" json:"bfd_detection_multiplier,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Ospfv3EdmInterfaceBfd) Reset()         { *m = Ospfv3EdmInterfaceBfd{} }
func (m *Ospfv3EdmInterfaceBfd) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterfaceBfd) ProtoMessage()    {}
func (*Ospfv3EdmInterfaceBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd988687fd2d89, []int{3}
}

func (m *Ospfv3EdmInterfaceBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmInterfaceBfd.Unmarshal(m, b)
}
func (m *Ospfv3EdmInterfaceBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmInterfaceBfd.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmInterfaceBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmInterfaceBfd.Merge(m, src)
}
func (m *Ospfv3EdmInterfaceBfd) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmInterfaceBfd.Size(m)
}
func (m *Ospfv3EdmInterfaceBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmInterfaceBfd.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmInterfaceBfd proto.InternalMessageInfo

func (m *Ospfv3EdmInterfaceBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *Ospfv3EdmInterfaceBfd) GetBfdInterval() uint32 {
	if m != nil {
		return m.BfdInterval
	}
	return 0
}

func (m *Ospfv3EdmInterfaceBfd) GetBfdDetectionMultiplier() uint32 {
	if m != nil {
		return m.BfdDetectionMultiplier
	}
	return 0
}

type Ospfv3EdmInterface struct {
	InterfaceState                              string                   `protobuf:"bytes,50,opt,name=interface_state,json=interfaceState,proto3" json:"interface_state,omitempty"`
	IsInterfaceLineUp                           bool                     `protobuf:"varint,51,opt,name=is_interface_line_up,json=isInterfaceLineUp,proto3" json:"is_interface_line_up,omitempty"`
	IsInterfaceIpSecurityRequired               bool                     `protobuf:"varint,52,opt,name=is_interface_ip_security_required,json=isInterfaceIpSecurityRequired,proto3" json:"is_interface_ip_security_required,omitempty"`
	IsInterfaceIpSecurityActive                 bool                     `protobuf:"varint,53,opt,name=is_interface_ip_security_active,json=isInterfaceIpSecurityActive,proto3" json:"is_interface_ip_security_active,omitempty"`
	InterfaceAddress                            string                   `protobuf:"bytes,54,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	InterfaceNumber                             uint32                   `protobuf:"varint,55,opt,name=interface_number,json=interfaceNumber,proto3" json:"interface_number,omitempty"`
	InterfaceRouterId                           string                   `protobuf:"bytes,56,opt,name=interface_router_id,json=interfaceRouterId,proto3" json:"interface_router_id,omitempty"`
	NetworkType                                 string                   `protobuf:"bytes,57,opt,name=network_type,json=networkType,proto3" json:"network_type,omitempty"`
	InterfaceLinkCost                           uint32                   `protobuf:"varint,58,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	IsInterfaceFloodReduction                   bool                     `protobuf:"varint,59,opt,name=is_interface_flood_reduction,json=isInterfaceFloodReduction,proto3" json:"is_interface_flood_reduction,omitempty"`
	IsDemandCircuitConfigured                   bool                     `protobuf:"varint,60,opt,name=is_demand_circuit_configured,json=isDemandCircuitConfigured,proto3" json:"is_demand_circuit_configured,omitempty"`
	IsInterfaceDemandCircuit                    bool                     `protobuf:"varint,61,opt,name=is_interface_demand_circuit,json=isInterfaceDemandCircuit,proto3" json:"is_interface_demand_circuit,omitempty"`
	InterfaceDcBitlessLsAs                      uint32                   `protobuf:"varint,62,opt,name=interface_dc_bitless_ls_as,json=interfaceDcBitlessLsAs,proto3" json:"interface_dc_bitless_ls_as,omitempty"`
	TransmissionDelay                           uint32                   `protobuf:"varint,63,opt,name=transmission_delay,json=transmissionDelay,proto3" json:"transmission_delay,omitempty"`
	OspfInterfaceState                          string                   `protobuf:"bytes,64,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	InterfacePriority                           uint32                   `protobuf:"varint,65,opt,name=interface_priority,json=interfacePriority,proto3" json:"interface_priority,omitempty"`
	IsDesignatedRouter                          bool                     `protobuf:"varint,66,opt,name=is_designated_router,json=isDesignatedRouter,proto3" json:"is_designated_router,omitempty"`
	DesignatedRouterId                          string                   `protobuf:"bytes,67,opt,name=designated_router_id,json=designatedRouterId,proto3" json:"designated_router_id,omitempty"`
	DesignatedRouterAddress                     string                   `protobuf:"bytes,68,opt,name=designated_router_address,json=designatedRouterAddress,proto3" json:"designated_router_address,omitempty"`
	BackupDesignatedRouterId                    string                   `protobuf:"bytes,69,opt,name=backup_designated_router_id,json=backupDesignatedRouterId,proto3" json:"backup_designated_router_id,omitempty"`
	BackupDesignatedRouterAddress               string                   `protobuf:"bytes,70,opt,name=backup_designated_router_address,json=backupDesignatedRouterAddress,proto3" json:"backup_designated_router_address,omitempty"`
	NetworkLsaFlushTimer                        uint32                   `protobuf:"varint,71,opt,name=network_lsa_flush_timer,json=networkLsaFlushTimer,proto3" json:"network_lsa_flush_timer,omitempty"`
	IsInterfaceLsaFiltered                      bool                     `protobuf:"varint,72,opt,name=is_interface_lsa_filtered,json=isInterfaceLsaFiltered,proto3" json:"is_interface_lsa_filtered,omitempty"`
	HelloInterval                               uint32                   `protobuf:"varint,73,opt,name=hello_interval,json=helloInterval,proto3" json:"hello_interval,omitempty"`
	DeadInterval                                uint32                   `protobuf:"varint,74,opt,name=dead_interval,json=deadInterval,proto3" json:"dead_interval,omitempty"`
	WaitInterval                                uint32                   `protobuf:"varint,75,opt,name=wait_interval,json=waitInterval,proto3" json:"wait_interval,omitempty"`
	InterfaceRetransmissionInterval             uint32                   `protobuf:"varint,76,opt,name=interface_retransmission_interval,json=interfaceRetransmissionInterval,proto3" json:"interface_retransmission_interval,omitempty"`
	NextHelloTime                               uint32                   `protobuf:"varint,77,opt,name=next_hello_time,json=nextHelloTime,proto3" json:"next_hello_time,omitempty"`
	InterfaceAuthenticationSpi                  uint32                   `protobuf:"varint,78,opt,name=interface_authentication_spi,json=interfaceAuthenticationSpi,proto3" json:"interface_authentication_spi,omitempty"`
	InterfaceAuthenticationTransmit             uint32                   `protobuf:"varint,79,opt,name=interface_authentication_transmit,json=interfaceAuthenticationTransmit,proto3" json:"interface_authentication_transmit,omitempty"`
	IsInterfaceEncryptionEnabled                bool                     `protobuf:"varint,80,opt,name=is_interface_encryption_enabled,json=isInterfaceEncryptionEnabled,proto3" json:"is_interface_encryption_enabled,omitempty"`
	IsPrefixSuppress                            bool                     `protobuf:"varint,81,opt,name=is_prefix_suppress,json=isPrefixSuppress,proto3" json:"is_prefix_suppress,omitempty"`
	InterfaceEncryptionSpi                      uint32                   `protobuf:"varint,82,opt,name=interface_encryption_spi,json=interfaceEncryptionSpi,proto3" json:"interface_encryption_spi,omitempty"`
	InterfaceEncryptionTransmitted              uint32                   `protobuf:"varint,83,opt,name=interface_encryption_transmitted,json=interfaceEncryptionTransmitted,proto3" json:"interface_encryption_transmitted,omitempty"`
	InterfaceEncryptedAuthenticationTransmitted uint32                   `protobuf:"varint,84,opt,name=interface_encrypted_authentication_transmitted,json=interfaceEncryptedAuthenticationTransmitted,proto3" json:"interface_encrypted_authentication_transmitted,omitempty"`
	InterfaceNeighbor                           []*Ospfv3EdmInterfaceNbr `protobuf:"bytes,85,rep,name=interface_neighbor,json=interfaceNeighbor,proto3" json:"interface_neighbor,omitempty"`
	AdjacentNeighbor                            uint32                   `protobuf:"varint,86,opt,name=adjacent_neighbor,json=adjacentNeighbor,proto3" json:"adjacent_neighbor,omitempty"`
	ActiveInterface                             *Ospfv3EdmInterfaceUp    `protobuf:"bytes,87,opt,name=active_interface,json=activeInterface,proto3" json:"active_interface,omitempty"`
	InterfaceBfd                                *Ospfv3EdmInterfaceBfd   `protobuf:"bytes,88,opt,name=interface_bfd,json=interfaceBfd,proto3" json:"interface_bfd,omitempty"`
	InterfaceReferences                         uint32                   `protobuf:"varint,89,opt,name=interface_references,json=interfaceReferences,proto3" json:"interface_references,omitempty"`
	ConfiguredLdpSync                           bool                     `protobuf:"varint,90,opt,name=configured_ldp_sync,json=configuredLdpSync,proto3" json:"configured_ldp_sync,omitempty"`
	InterfaceLdpSync                            bool                     `protobuf:"varint,91,opt,name=interface_ldp_sync,json=interfaceLdpSync,proto3" json:"interface_ldp_sync,omitempty"`
	XXX_NoUnkeyedLiteral                        struct{}                 `json:"-"`
	XXX_unrecognized                            []byte                   `json:"-"`
	XXX_sizecache                               int32                    `json:"-"`
}

func (m *Ospfv3EdmInterface) Reset()         { *m = Ospfv3EdmInterface{} }
func (m *Ospfv3EdmInterface) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmInterface) ProtoMessage()    {}
func (*Ospfv3EdmInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_25fd988687fd2d89, []int{4}
}

func (m *Ospfv3EdmInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmInterface.Unmarshal(m, b)
}
func (m *Ospfv3EdmInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmInterface.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmInterface.Merge(m, src)
}
func (m *Ospfv3EdmInterface) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmInterface.Size(m)
}
func (m *Ospfv3EdmInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmInterface.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmInterface proto.InternalMessageInfo

func (m *Ospfv3EdmInterface) GetInterfaceState() string {
	if m != nil {
		return m.InterfaceState
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetIsInterfaceLineUp() bool {
	if m != nil {
		return m.IsInterfaceLineUp
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceIpSecurityRequired() bool {
	if m != nil {
		return m.IsInterfaceIpSecurityRequired
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceIpSecurityActive() bool {
	if m != nil {
		return m.IsInterfaceIpSecurityActive
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfaceNumber() uint32 {
	if m != nil {
		return m.InterfaceNumber
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceRouterId() string {
	if m != nil {
		return m.InterfaceRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetNetworkType() string {
	if m != nil {
		return m.NetworkType
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceFloodReduction() bool {
	if m != nil {
		return m.IsInterfaceFloodReduction
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsDemandCircuitConfigured() bool {
	if m != nil {
		return m.IsDemandCircuitConfigured
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsInterfaceDemandCircuit() bool {
	if m != nil {
		return m.IsInterfaceDemandCircuit
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceDcBitlessLsAs() uint32 {
	if m != nil {
		return m.InterfaceDcBitlessLsAs
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetTransmissionDelay() uint32 {
	if m != nil {
		return m.TransmissionDelay
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetInterfacePriority() uint32 {
	if m != nil {
		return m.InterfacePriority
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsDesignatedRouter() bool {
	if m != nil {
		return m.IsDesignatedRouter
	}
	return false
}

func (m *Ospfv3EdmInterface) GetDesignatedRouterId() string {
	if m != nil {
		return m.DesignatedRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetDesignatedRouterAddress() string {
	if m != nil {
		return m.DesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetBackupDesignatedRouterId() string {
	if m != nil {
		return m.BackupDesignatedRouterId
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetBackupDesignatedRouterAddress() string {
	if m != nil {
		return m.BackupDesignatedRouterAddress
	}
	return ""
}

func (m *Ospfv3EdmInterface) GetNetworkLsaFlushTimer() uint32 {
	if m != nil {
		return m.NetworkLsaFlushTimer
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceLsaFiltered() bool {
	if m != nil {
		return m.IsInterfaceLsaFiltered
	}
	return false
}

func (m *Ospfv3EdmInterface) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetDeadInterval() uint32 {
	if m != nil {
		return m.DeadInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetWaitInterval() uint32 {
	if m != nil {
		return m.WaitInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceRetransmissionInterval() uint32 {
	if m != nil {
		return m.InterfaceRetransmissionInterval
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetNextHelloTime() uint32 {
	if m != nil {
		return m.NextHelloTime
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceAuthenticationSpi() uint32 {
	if m != nil {
		return m.InterfaceAuthenticationSpi
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceAuthenticationTransmit() uint32 {
	if m != nil {
		return m.InterfaceAuthenticationTransmit
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetIsInterfaceEncryptionEnabled() bool {
	if m != nil {
		return m.IsInterfaceEncryptionEnabled
	}
	return false
}

func (m *Ospfv3EdmInterface) GetIsPrefixSuppress() bool {
	if m != nil {
		return m.IsPrefixSuppress
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptionSpi() uint32 {
	if m != nil {
		return m.InterfaceEncryptionSpi
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptionTransmitted() uint32 {
	if m != nil {
		return m.InterfaceEncryptionTransmitted
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceEncryptedAuthenticationTransmitted() uint32 {
	if m != nil {
		return m.InterfaceEncryptedAuthenticationTransmitted
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetInterfaceNeighbor() []*Ospfv3EdmInterfaceNbr {
	if m != nil {
		return m.InterfaceNeighbor
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetAdjacentNeighbor() uint32 {
	if m != nil {
		return m.AdjacentNeighbor
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetActiveInterface() *Ospfv3EdmInterfaceUp {
	if m != nil {
		return m.ActiveInterface
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetInterfaceBfd() *Ospfv3EdmInterfaceBfd {
	if m != nil {
		return m.InterfaceBfd
	}
	return nil
}

func (m *Ospfv3EdmInterface) GetInterfaceReferences() uint32 {
	if m != nil {
		return m.InterfaceReferences
	}
	return 0
}

func (m *Ospfv3EdmInterface) GetConfiguredLdpSync() bool {
	if m != nil {
		return m.ConfiguredLdpSync
	}
	return false
}

func (m *Ospfv3EdmInterface) GetInterfaceLdpSync() bool {
	if m != nil {
		return m.InterfaceLdpSync
	}
	return false
}

func init() {
	proto.RegisterType((*Ospfv3EdmInterface_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.interface_process_table.interface.ospfv3_edm_interface_KEYS")
	proto.RegisterType((*Ospfv3EdmInterfaceNbr)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.interface_process_table.interface.ospfv3_edm_interface_nbr")
	proto.RegisterType((*Ospfv3EdmInterfaceUp)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.interface_process_table.interface.ospfv3_edm_interface_up")
	proto.RegisterType((*Ospfv3EdmInterfaceBfd)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.interface_process_table.interface.ospfv3_edm_interface_bfd")
	proto.RegisterType((*Ospfv3EdmInterface)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.interface_process_table.interface.ospfv3_edm_interface")
}

func init() { proto.RegisterFile("ospfv3_edm_interface.proto", fileDescriptor_25fd988687fd2d89) }

var fileDescriptor_25fd988687fd2d89 = []byte{
	// 1562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x58, 0x5f, 0x57, 0x54, 0xb7,
	0x16, 0x5f, 0xa3, 0x5e, 0xc5, 0xc0, 0x08, 0x04, 0x94, 0x20, 0x7a, 0x2f, 0x72, 0xef, 0xb5, 0xb6,
	0xea, 0xd8, 0x62, 0xfd, 0x5f, 0xab, 0xc0, 0x80, 0x8c, 0xa2, 0xc5, 0x19, 0x6c, 0x6b, 0xfb, 0x90,
	0x95, 0x39, 0xc9, 0x81, 0x94, 0x33, 0xe7, 0x1c, 0x93, 0x1c, 0x84, 0xf6, 0xa5, 0x0f, 0x7d, 0xe9,
	0xb7, 0xe8, 0x4b, 0xd7, 0xea, 0xea, 0x47, 0xea, 0xa7, 0xe9, 0xca, 0xce, 0xf9, 0xcf, 0xcc, 0x63,
	0x97, 0x2f, 0xac, 0x21, 0xfb, 0xf7, 0xdb, 0x7b, 0x67, 0x67, 0xff, 0x9b, 0x41, 0x17, 0x23, 0x1d,
	0xfb, 0x07, 0xb7, 0xa9, 0xe0, 0x03, 0x2a, 0x43, 0x23, 0x94, 0xcf, 0x3c, 0xd1, 0x8a, 0x55, 0x64,
	0x22, 0xcc, 0x3c, 0xa9, 0xbd, 0x88, 0xca, 0x48, 0xd3, 0x43, 0x45, 0x65, 0x7c, 0x70, 0x97, 0xa6,
	0xe8, 0x28, 0x16, 0xaa, 0xe5, 0x3e, 0x5b, 0xac, 0x27, 0xb4, 0x16, 0x3a, 0xfb, 0xd4, 0x3a, 0x50,
	0x3e, 0xfc, 0x69, 0xe5, 0x0a, 0x69, 0x2a, 0xa2, 0x86, 0xf5, 0x03, 0x51, 0x9c, 0x2f, 0xfd, 0xdc,
	0x40, 0xf3, 0xc3, 0x3c, 0xa0, 0x2f, 0xd6, 0xdf, 0xf6, 0xf0, 0x15, 0x34, 0x91, 0x11, 0x43, 0x36,
	0x10, 0xa4, 0xb1, 0xd8, 0xb8, 0x76, 0xb6, 0x3b, 0x9e, 0x9e, 0xbd, 0x62, 0x03, 0x81, 0xe7, 0xd1,
	0xd8, 0x81, 0xf2, 0x9d, 0xf8, 0x04, 0x88, 0xcf, 0x1c, 0x28, 0x1f, 0x44, 0xff, 0x47, 0xe7, 0x0a,
	0x7d, 0x00, 0x38, 0x09, 0x80, 0x66, 0x7e, 0x6a, 0x61, 0x4b, 0xbf, 0x9c, 0x40, 0x64, 0xa8, 0x0b,
	0x61, 0x5f, 0xe1, 0x65, 0x74, 0xbe, 0x74, 0x20, 0xe4, 0xee, 0x5e, 0x3f, 0x52, 0x54, 0xf2, 0xd4,
	0x95, 0x99, 0x42, 0x55, 0x2a, 0xeb, 0x70, 0x7c, 0x17, 0xcd, 0x0d, 0xe1, 0x78, 0x91, 0x36, 0xe0,
	0x61, 0xb3, 0x7b, 0xfe, 0x18, 0x6b, 0x2d, 0xd2, 0x06, 0xff, 0x0f, 0x9d, 0x93, 0xba, 0x20, 0x70,
	0x05, 0xfe, 0x8e, 0x75, 0x27, 0xa4, 0xce, 0x70, 0x6d, 0x85, 0xaf, 0xa2, 0xc9, 0x32, 0xaa, 0xcf,
	0x15, 0x39, 0x05, 0xb0, 0x66, 0x01, 0x5b, 0xe5, 0x0a, 0xb7, 0xd0, 0x8c, 0xd4, 0x74, 0x4f, 0x04,
	0x41, 0x44, 0x75, 0x12, 0xc7, 0xca, 0xbe, 0x11, 0x27, 0xff, 0x02, 0xec, 0xb4, 0xd4, 0x9b, 0x56,
	0xd2, 0xcb, 0x05, 0x4b, 0x7f, 0x9d, 0x41, 0x73, 0x43, 0xc3, 0x90, 0xc4, 0x78, 0x01, 0x9d, 0x7d,
	0xcf, 0xa4, 0xa1, 0x46, 0xa6, 0x8f, 0xd0, 0xec, 0x8e, 0xd9, 0x83, 0x1d, 0x39, 0x10, 0xf8, 0x11,
	0xba, 0x58, 0x80, 0x99, 0x12, 0x8c, 0xfa, 0x41, 0x14, 0x71, 0x2a, 0x43, 0x2e, 0x0e, 0xd3, 0x1b,
	0x17, 0x01, 0x59, 0x51, 0x82, 0x6d, 0x58, 0x79, 0xc7, 0x8a, 0xf1, 0x3d, 0x44, 0x4a, 0x64, 0x5d,
	0xa1, 0x9e, 0xac, 0x05, 0x6b, 0x45, 0x97, 0x88, 0x15, 0xab, 0x81, 0x0c, 0xf7, 0x2b, 0xd4, 0x53,
	0x35, 0xab, 0x5b, 0x32, 0xdc, 0x2f, 0x91, 0x6f, 0x20, 0xec, 0xd0, 0xef, 0x12, 0x91, 0x08, 0x1a,
	0x88, 0x70, 0xd7, 0xec, 0x41, 0x68, 0x9a, 0xdd, 0x29, 0x90, 0xbc, 0xb6, 0x82, 0x2d, 0x38, 0xc7,
	0x0f, 0xd0, 0x7c, 0xed, 0x82, 0xa1, 0x38, 0x34, 0xce, 0x1e, 0x39, 0x0d, 0xa4, 0x0b, 0x95, 0xfb,
	0xbd, 0x12, 0x87, 0x06, 0xac, 0xe1, 0x36, 0xfa, 0xcf, 0x48, 0x6a, 0xea, 0xea, 0x19, 0x50, 0xb0,
	0x30, 0x5c, 0x81, 0x73, 0xf7, 0x4e, 0x39, 0xa1, 0x98, 0x2e, 0x9b, 0x1f, 0x03, 0xf6, 0x6c, 0x29,
	0x46, 0x85, 0xf1, 0xa7, 0xe8, 0xf2, 0x08, 0x5a, 0x6a, 0xfa, 0x2c, 0x90, 0xe7, 0x87, 0x91, 0x9d,
	0xe1, 0xca, 0xcd, 0x21, 0xc8, 0x25, 0xd3, 0xa8, 0x76, 0x73, 0x1b, 0xe3, 0xc2, 0xf8, 0x08, 0xaa,
	0x33, 0x3c, 0x3e, 0x82, 0xea, 0xac, 0x7e, 0x82, 0xa6, 0x9d, 0x97, 0xda, 0x63, 0x61, 0xf6, 0x38,
	0x13, 0x40, 0x99, 0x04, 0x41, 0xcf, 0x63, 0x61, 0xfa, 0x36, 0x9f, 0xa2, 0xd9, 0x01, 0x3b, 0x94,
	0x83, 0x64, 0x90, 0xde, 0x2c, 0x85, 0x37, 0x01, 0x8e, 0x53, 0x19, 0xb8, 0x94, 0x32, 0xae, 0xa2,
	0xc9, 0x80, 0xe9, 0x2c, 0x10, 0x90, 0xd1, 0xe7, 0x00, 0xdc, 0xb4, 0xc7, 0x80, 0x84, 0xb4, 0xbe,
	0x81, 0x70, 0x55, 0x33, 0x40, 0x27, 0x5d, 0x8e, 0x94, 0xf5, 0x02, 0xfa, 0x09, 0xba, 0x54, 0x5c,
	0xd7, 0xe1, 0x63, 0xe6, 0xc9, 0x70, 0x17, 0x68, 0x8a, 0x4c, 0xd5, 0x42, 0x0d, 0xcc, 0x6d, 0x40,
	0x58, 0xbe, 0xc2, 0xb7, 0xd0, 0xcc, 0xf1, 0xa6, 0xa1, 0xc9, 0xb4, 0xbb, 0xc7, 0xb1, 0x86, 0xa1,
	0xf1, 0x75, 0x34, 0x5d, 0x94, 0xb5, 0xab, 0x73, 0x4d, 0xb0, 0x73, 0xaf, 0x10, 0x40, 0x95, 0xeb,
	0xa5, 0x3f, 0x1a, 0x23, 0x7a, 0x5c, 0xdf, 0xe7, 0xf8, 0x16, 0x9a, 0xed, 0xfb, 0x36, 0x27, 0x8c,
	0x4f, 0x45, 0x68, 0x1b, 0x34, 0x1d, 0x44, 0x3c, 0x2b, 0xf4, 0xe9, 0xbe, 0xcf, 0x3b, 0xa1, 0xf1,
	0xd7, 0x41, 0xf2, 0x32, 0xe2, 0xc2, 0xb6, 0xe5, 0x94, 0x20, 0xd4, 0x01, 0x0b, 0xd2, 0x1a, 0x1f,
	0x77, 0x40, 0x38, 0xc2, 0xf7, 0x11, 0xb1, 0x10, 0x2e, 0x8c, 0xf0, 0x8c, 0x8c, 0x42, 0x3a, 0x48,
	0x02, 0x23, 0xe3, 0x40, 0x0a, 0x95, 0xd6, 0xf5, 0x85, 0xbe, 0xcf, 0xdb, 0x99, 0xf8, 0x65, 0x2e,
	0x5d, 0xfa, 0xf5, 0x3c, 0x9a, 0x1d, 0xe6, 0x2a, 0xfe, 0x08, 0x4d, 0x16, 0x7e, 0x6b, 0xc3, 0x8c,
	0x20, 0xcb, 0xd0, 0x84, 0x8b, 0x2e, 0xdf, 0xb3, 0xa7, 0xf6, 0x3e, 0x52, 0xd3, 0x4a, 0xf6, 0xd9,
	0x2e, 0x46, 0x6e, 0x67, 0xad, 0xaf, 0x53, 0xca, 0x3b, 0xf1, 0x26, 0xc6, 0x9b, 0xe8, 0x4a, 0x85,
	0x20, 0x63, 0xaa, 0x85, 0x97, 0x28, 0x69, 0x8e, 0xa8, 0x12, 0xef, 0x12, 0xa9, 0x04, 0x27, 0x9f,
	0x03, 0xfb, 0x72, 0x89, 0xdd, 0x89, 0x7b, 0x29, 0xaa, 0x9b, 0x82, 0xa0, 0xde, 0x47, 0x69, 0x62,
	0x9e, 0x91, 0x07, 0x82, 0xdc, 0x01, 0x3d, 0x0b, 0x43, 0xf5, 0xac, 0x00, 0xc4, 0x3e, 0x6d, 0xa9,
	0x70, 0x39, 0xb7, 0x4f, 0x49, 0xee, 0xc2, 0x5d, 0xa7, 0x8a, 0x62, 0x75, 0xe7, 0xf8, 0x63, 0x34,
	0x55, 0x4a, 0x9c, 0x64, 0xd0, 0x17, 0x8a, 0xdc, 0x73, 0xc5, 0x52, 0x64, 0x0d, 0x1c, 0xc3, 0x48,
	0xc8, 0xa1, 0x2a, 0x4a, 0x8c, 0x80, 0x51, 0x76, 0x1f, 0x34, 0x17, 0x26, 0xbb, 0x20, 0xe9, 0x70,
	0xfb, 0xce, 0xa1, 0x30, 0xef, 0x23, 0xb5, 0x4f, 0xcd, 0x51, 0x2c, 0xc8, 0x03, 0x37, 0x7e, 0xd3,
	0xb3, 0x9d, 0xa3, 0x58, 0x54, 0x55, 0x42, 0x99, 0xc3, 0x9c, 0x7b, 0xe8, 0x52, 0xa7, 0x52, 0xe0,
	0x30, 0xe3, 0x6c, 0x9d, 0x94, 0x03, 0xe4, 0x4a, 0x45, 0x09, 0x9e, 0x40, 0x1e, 0x90, 0x47, 0x10,
	0x9d, 0xf9, 0x52, 0x74, 0xa0, 0x52, 0xba, 0x19, 0x20, 0x55, 0xc0, 0xc5, 0x80, 0x85, 0x9c, 0x7a,
	0x52, 0x79, 0x89, 0x34, 0xd4, 0x8b, 0x42, 0x5f, 0xee, 0x26, 0xf6, 0x99, 0xbe, 0xc8, 0x14, 0xb4,
	0x01, 0xb2, 0xe6, 0x10, 0x6b, 0x39, 0x00, 0x3f, 0x46, 0x0b, 0x15, 0x0f, 0xaa, 0xaa, 0xc8, 0x63,
	0xe0, 0x93, 0x92, 0x03, 0x15, 0x45, 0xf8, 0x61, 0x79, 0xee, 0x70, 0x8f, 0xf6, 0xa5, 0x09, 0xec,
	0x7e, 0x12, 0x68, 0xca, 0x34, 0xf9, 0xb2, 0xd6, 0xd8, 0xda, 0xde, 0xaa, 0x93, 0x6f, 0xe9, 0x15,
	0x8d, 0x6f, 0x22, 0x6c, 0x14, 0x0b, 0xf5, 0x40, 0x6a, 0x6d, 0x6b, 0x82, 0x8b, 0x80, 0x1d, 0x91,
	0x27, 0x2e, 0x56, 0x65, 0x49, 0xdb, 0x0a, 0x6c, 0x6f, 0xb3, 0x85, 0x40, 0xeb, 0x59, 0xff, 0x14,
	0x9e, 0x01, 0x5b, 0x59, 0xa7, 0x9a, 0xf9, 0x37, 0x11, 0x2e, 0xaf, 0x5c, 0x32, 0xb2, 0x49, 0x45,
	0x56, 0x6a, 0x8f, 0xb1, 0x9d, 0x0a, 0xac, 0x01, 0x88, 0xa5, 0x96, 0xbb, 0x21, 0x33, 0x82, 0xa7,
	0x39, 0x41, 0x56, 0x21, 0x06, 0xd8, 0xc6, 0x30, 0x13, 0xb9, 0x9c, 0xb0, 0x8c, 0x63, 0x70, 0x9b,
	0x42, 0x6b, 0xce, 0x25, 0x5e, 0xc3, 0x77, 0x38, 0x7e, 0x88, 0xe6, 0x8f, 0x33, 0xb2, 0x9c, 0x6e,
	0x03, 0x6d, 0xae, 0x4e, 0xcb, 0x52, 0xfb, 0x31, 0x5a, 0xe8, 0x33, 0x6f, 0x3f, 0x89, 0xe9, 0x50,
	0xa3, 0xeb, 0xc0, 0x26, 0x0e, 0xd2, 0x3e, 0x6e, 0xfa, 0x19, 0x5a, 0x1c, 0x49, 0xcf, 0x3c, 0xd8,
	0x00, 0x1d, 0x97, 0x87, 0xeb, 0xc8, 0xfc, 0xb8, 0x83, 0xe6, 0xb2, 0x3a, 0x08, 0xb4, 0x5d, 0x6f,
	0x12, 0xbd, 0x97, 0xf6, 0xf5, 0x67, 0x6e, 0xfe, 0xa6, 0xe2, 0x2d, 0xcd, 0x36, 0xac, 0xd0, 0xb5,
	0x74, 0x3b, 0x02, 0x2b, 0x7d, 0xc8, 0x72, 0x65, 0x60, 0x84, 0xcd, 0xd3, 0x4d, 0x88, 0xf1, 0x85,
	0x72, 0x33, 0xd2, 0x6c, 0x23, 0x95, 0xda, 0xd5, 0xd5, 0x6d, 0x6e, 0x79, 0x8f, 0xed, 0xb8, 0x19,
	0x05, 0xa7, 0x79, 0x97, 0xfd, 0x2f, 0x6a, 0x72, 0xc1, 0x4a, 0x9d, 0xf8, 0x39, 0xa0, 0x26, 0xec,
	0x61, 0x19, 0x04, 0xcb, 0x5b, 0x0e, 0x7a, 0xe1, 0x40, 0xf6, 0x30, 0x07, 0x3d, 0x47, 0x57, 0x4a,
	0xad, 0x41, 0x54, 0xd2, 0x34, 0x27, 0x6e, 0x01, 0xb1, 0xd8, 0x68, 0xba, 0x15, 0x5c, 0xae, 0xeb,
	0x2a, 0x9a, 0x84, 0x59, 0xef, 0x6e, 0x00, 0x63, 0xf3, 0xa5, 0xf3, 0xde, 0x1e, 0xc3, 0x44, 0x82,
	0x99, 0xf9, 0xb4, 0x3c, 0x33, 0x59, 0x62, 0xf6, 0x44, 0x68, 0xa4, 0xc7, 0x60, 0x5c, 0xe8, 0x58,
	0x92, 0x57, 0x40, 0x2a, 0xca, 0x6d, 0xa5, 0x02, 0xe9, 0xc5, 0xb2, 0xea, 0x75, 0x4d, 0x43, 0xea,
	0x9a, 0x21, 0x5f, 0xd5, 0xbc, 0xae, 0xaa, 0xd9, 0x49, 0x61, 0x78, 0xbd, 0xd6, 0xba, 0x45, 0xe8,
	0xa9, 0xa3, 0x18, 0x54, 0xb9, 0xa1, 0xc8, 0xc9, 0x36, 0xbc, 0xd9, 0xa5, 0xd2, 0x9b, 0xad, 0xe7,
	0x20, 0x37, 0x1e, 0xb9, 0x5d, 0x1b, 0xa4, 0xa6, 0xb1, 0x12, 0xbe, 0x3c, 0xcc, 0xf7, 0x6e, 0xf2,
	0x1a, 0x98, 0x53, 0x52, 0x6f, 0x83, 0x20, 0x5b, 0xbb, 0xed, 0x98, 0x1c, 0x6a, 0xd1, 0x5e, 0xbf,
	0x5b, 0xeb, 0x25, 0x85, 0x2d, 0x7b, 0xf5, 0x4d, 0xb4, 0x38, 0x94, 0x99, 0x5d, 0xdb, 0x08, 0x4e,
	0x7a, 0xa0, 0xe1, 0xdf, 0x43, 0x34, 0xec, 0x14, 0x28, 0xec, 0xa1, 0xd6, 0x31, 0x4d, 0x82, 0x8f,
	0x0a, 0xa7, 0xd5, 0xbb, 0x03, 0x7a, 0xaf, 0xd7, 0xf5, 0x0a, 0x3e, 0x3c, 0xb4, 0xd6, 0xc8, 0x9f,
	0x8d, 0x72, 0x6b, 0xca, 0xf6, 0x1b, 0xf2, 0x66, 0xf1, 0xe4, 0xb5, 0xf1, 0xe5, 0x9f, 0x5a, 0xff,
	0xf8, 0x17, 0xcd, 0xd6, 0xa8, 0x6f, 0x78, 0xa5, 0xbe, 0x98, 0xed, 0x56, 0x76, 0xfe, 0x32, 0xfe,
	0x03, 0xf3, 0x44, 0x68, 0x0a, 0x57, 0xbf, 0x76, 0xab, 0x55, 0x26, 0xc8, 0xc1, 0xbf, 0x37, 0xd0,
	0x94, 0x1b, 0xed, 0x85, 0x62, 0xf2, 0xcd, 0x62, 0xe3, 0xda, 0xf8, 0xf2, 0x8f, 0x1f, 0xea, 0x5e,
	0x49, 0xdc, 0x9d, 0x74, 0x3e, 0xe5, 0x99, 0x8a, 0x7f, 0x6b, 0xa0, 0x66, 0x65, 0xef, 0x23, 0xdf,
	0x82, 0x93, 0x1f, 0x2c, 0xf8, 0x7d, 0x9f, 0x77, 0x27, 0xf2, 0x7f, 0x57, 0x7d, 0x8e, 0x3f, 0x43,
	0xb3, 0xe5, 0x26, 0xe4, 0x0b, 0x25, 0x42, 0x4f, 0x68, 0xf2, 0x16, 0x42, 0x3f, 0x53, 0xea, 0x3b,
	0x99, 0xc8, 0xee, 0x1f, 0xc5, 0xf0, 0xa7, 0x01, 0x8f, 0xa9, 0x3e, 0x0a, 0x3d, 0xf2, 0x9d, 0x5b,
	0xf5, 0x0a, 0xd1, 0x16, 0x8f, 0x7b, 0x47, 0xa1, 0x07, 0xe5, 0x59, 0x34, 0xe4, 0x0c, 0xfe, 0x7d,
	0x5a, 0x9e, 0x79, 0x2b, 0x76, 0xe8, 0xfe, 0x69, 0xf8, 0x1d, 0xe4, 0xf6, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x7c, 0x2e, 0x6c, 0xf0, 0x25, 0x11, 0x00, 0x00,
}