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
// source: ospf_sh_neighbor_detail.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_adjacency_information_neighbor_details_neighbor_detail

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

type OspfShNeighborDetail_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborDetail_KEYS) Reset()         { *m = OspfShNeighborDetail_KEYS{} }
func (m *OspfShNeighborDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborDetail_KEYS) ProtoMessage()    {}
func (*OspfShNeighborDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66ebdf8fd8c36e2, []int{0}
}

func (m *OspfShNeighborDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Unmarshal(m, b)
}
func (m *OspfShNeighborDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShNeighborDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborDetail_KEYS.Merge(m, src)
}
func (m *OspfShNeighborDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborDetail_KEYS.Size(m)
}
func (m *OspfShNeighborDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborDetail_KEYS proto.InternalMessageInfo

func (m *OspfShNeighborDetail_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShNeighborDetail_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShNeighborDetail_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShNeighborDetail_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShNeighborBfd struct {
	BfdIntfEnableMode    uint32   `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	BfdStatusFlag        uint32   `protobuf:"varint,2,opt,name=bfd_status_flag,json=bfdStatusFlag,proto3" json:"bfd_status_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborBfd) Reset()         { *m = OspfShNeighborBfd{} }
func (m *OspfShNeighborBfd) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborBfd) ProtoMessage()    {}
func (*OspfShNeighborBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66ebdf8fd8c36e2, []int{1}
}

func (m *OspfShNeighborBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborBfd.Unmarshal(m, b)
}
func (m *OspfShNeighborBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborBfd.Marshal(b, m, deterministic)
}
func (m *OspfShNeighborBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborBfd.Merge(m, src)
}
func (m *OspfShNeighborBfd) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborBfd.Size(m)
}
func (m *OspfShNeighborBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborBfd.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborBfd proto.InternalMessageInfo

func (m *OspfShNeighborBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *OspfShNeighborBfd) GetBfdStatusFlag() uint32 {
	if m != nil {
		return m.BfdStatusFlag
	}
	return 0
}

type OspfShNeighbor struct {
	NeighborId             string             `protobuf:"bytes,1,opt,name=neighbor_id,json=neighborId,proto3" json:"neighbor_id,omitempty"`
	NeighborAddressXr      string             `protobuf:"bytes,2,opt,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	NeighborInterfaceName  string             `protobuf:"bytes,3,opt,name=neighbor_interface_name,json=neighborInterfaceName,proto3" json:"neighbor_interface_name,omitempty"`
	NeighborDrPriority     uint32             `protobuf:"varint,4,opt,name=neighbor_dr_priority,json=neighborDrPriority,proto3" json:"neighbor_dr_priority,omitempty"`
	NeighborState          string             `protobuf:"bytes,5,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	DrBdrState             string             `protobuf:"bytes,6,opt,name=dr_bdr_state,json=drBdrState,proto3" json:"dr_bdr_state,omitempty"`
	NeighborDeadTimer      uint32             `protobuf:"varint,7,opt,name=neighbor_dead_timer,json=neighborDeadTimer,proto3" json:"neighbor_dead_timer,omitempty"`
	NeighborUpTime         uint32             `protobuf:"varint,8,opt,name=neighbor_up_time,json=neighborUpTime,proto3" json:"neighbor_up_time,omitempty"`
	NeighborMadjInterface  bool               `protobuf:"varint,9,opt,name=neighbor_madj_interface,json=neighborMadjInterface,proto3" json:"neighbor_madj_interface,omitempty"`
	NeighborBfdInformation *OspfShNeighborBfd `protobuf:"bytes,10,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShNeighbor) Reset()         { *m = OspfShNeighbor{} }
func (m *OspfShNeighbor) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor) ProtoMessage()    {}
func (*OspfShNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66ebdf8fd8c36e2, []int{2}
}

func (m *OspfShNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor.Unmarshal(m, b)
}
func (m *OspfShNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor.Marshal(b, m, deterministic)
}
func (m *OspfShNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor.Merge(m, src)
}
func (m *OspfShNeighbor) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor.Size(m)
}
func (m *OspfShNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor proto.InternalMessageInfo

func (m *OspfShNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborAddressXr() string {
	if m != nil {
		return m.NeighborAddressXr
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDrPriority() uint32 {
	if m != nil {
		return m.NeighborDrPriority
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *OspfShNeighbor) GetDrBdrState() string {
	if m != nil {
		return m.DrBdrState
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDeadTimer() uint32 {
	if m != nil {
		return m.NeighborDeadTimer
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborUpTime() uint32 {
	if m != nil {
		return m.NeighborUpTime
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborMadjInterface() bool {
	if m != nil {
		return m.NeighborMadjInterface
	}
	return false
}

func (m *OspfShNeighbor) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

type OspfShNeighborRetrans struct {
	DbdRetransmissionCount         uint32   `protobuf:"varint,1,opt,name=dbd_retransmission_count,json=dbdRetransmissionCount,proto3" json:"dbd_retransmission_count,omitempty"`
	DbdRetransmissionTotalCount    uint32   `protobuf:"varint,2,opt,name=dbd_retransmission_total_count,json=dbdRetransmissionTotalCount,proto3" json:"dbd_retransmission_total_count,omitempty"`
	AreaFloodingIndex              uint32   `protobuf:"varint,3,opt,name=area_flooding_index,json=areaFloodingIndex,proto3" json:"area_flooding_index,omitempty"`
	AsFloodIndex                   uint32   `protobuf:"varint,4,opt,name=as_flood_index,json=asFloodIndex,proto3" json:"as_flood_index,omitempty"`
	NeighborRetransmissionCount    uint32   `protobuf:"varint,5,opt,name=neighbor_retransmission_count,json=neighborRetransmissionCount,proto3" json:"neighbor_retransmission_count,omitempty"`
	NumberOfRetransmissions        uint32   `protobuf:"varint,6,opt,name=number_of_retransmissions,json=numberOfRetransmissions,proto3" json:"number_of_retransmissions,omitempty"`
	AreaFirstFloodInformation      uint32   `protobuf:"varint,7,opt,name=area_first_flood_information,json=areaFirstFloodInformation,proto3" json:"area_first_flood_information,omitempty"`
	AreaFirstFloodInformationIndex uint32   `protobuf:"varint,8,opt,name=area_first_flood_information_index,json=areaFirstFloodInformationIndex,proto3" json:"area_first_flood_information_index,omitempty"`
	AsFirstFloodInformation        uint32   `protobuf:"varint,9,opt,name=as_first_flood_information,json=asFirstFloodInformation,proto3" json:"as_first_flood_information,omitempty"`
	AsFirstFloodInformationIndex   uint32   `protobuf:"varint,10,opt,name=as_first_flood_information_index,json=asFirstFloodInformationIndex,proto3" json:"as_first_flood_information_index,omitempty"`
	AreaNextFloodInformation       uint32   `protobuf:"varint,11,opt,name=area_next_flood_information,json=areaNextFloodInformation,proto3" json:"area_next_flood_information,omitempty"`
	AreaNextFloodInformationIndex  uint32   `protobuf:"varint,12,opt,name=area_next_flood_information_index,json=areaNextFloodInformationIndex,proto3" json:"area_next_flood_information_index,omitempty"`
	AsNextFloodInformation         uint32   `protobuf:"varint,13,opt,name=as_next_flood_information,json=asNextFloodInformation,proto3" json:"as_next_flood_information,omitempty"`
	AsNextFloodInformationIndex    uint32   `protobuf:"varint,14,opt,name=as_next_flood_information_index,json=asNextFloodInformationIndex,proto3" json:"as_next_flood_information_index,omitempty"`
	LastRetransmissionLength       uint32   `protobuf:"varint,15,opt,name=last_retransmission_length,json=lastRetransmissionLength,proto3" json:"last_retransmission_length,omitempty"`
	MaximumRetransmissionLength    uint32   `protobuf:"varint,16,opt,name=maximum_retransmission_length,json=maximumRetransmissionLength,proto3" json:"maximum_retransmission_length,omitempty"`
	LastRetransmissionTime         uint32   `protobuf:"varint,17,opt,name=last_retransmission_time,json=lastRetransmissionTime,proto3" json:"last_retransmission_time,omitempty"`
	MaximumRetransmissionTime      uint32   `protobuf:"varint,18,opt,name=maximum_retransmission_time,json=maximumRetransmissionTime,proto3" json:"maximum_retransmission_time,omitempty"`
	LsaRetransmissionTimer         uint32   `protobuf:"varint,19,opt,name=lsa_retransmission_timer,json=lsaRetransmissionTimer,proto3" json:"lsa_retransmission_timer,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *OspfShNeighborRetrans) Reset()         { *m = OspfShNeighborRetrans{} }
func (m *OspfShNeighborRetrans) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborRetrans) ProtoMessage()    {}
func (*OspfShNeighborRetrans) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66ebdf8fd8c36e2, []int{3}
}

func (m *OspfShNeighborRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborRetrans.Unmarshal(m, b)
}
func (m *OspfShNeighborRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborRetrans.Marshal(b, m, deterministic)
}
func (m *OspfShNeighborRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborRetrans.Merge(m, src)
}
func (m *OspfShNeighborRetrans) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborRetrans.Size(m)
}
func (m *OspfShNeighborRetrans) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborRetrans.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborRetrans proto.InternalMessageInfo

func (m *OspfShNeighborRetrans) GetDbdRetransmissionCount() uint32 {
	if m != nil {
		return m.DbdRetransmissionCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetDbdRetransmissionTotalCount() uint32 {
	if m != nil {
		return m.DbdRetransmissionTotalCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFloodingIndex() uint32 {
	if m != nil {
		return m.AreaFloodingIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFloodIndex() uint32 {
	if m != nil {
		return m.AsFloodIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetNeighborRetransmissionCount() uint32 {
	if m != nil {
		return m.NeighborRetransmissionCount
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetNumberOfRetransmissions() uint32 {
	if m != nil {
		return m.NumberOfRetransmissions
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFirstFloodInformation() uint32 {
	if m != nil {
		return m.AreaFirstFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaFirstFloodInformationIndex() uint32 {
	if m != nil {
		return m.AreaFirstFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFirstFloodInformation() uint32 {
	if m != nil {
		return m.AsFirstFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsFirstFloodInformationIndex() uint32 {
	if m != nil {
		return m.AsFirstFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaNextFloodInformation() uint32 {
	if m != nil {
		return m.AreaNextFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAreaNextFloodInformationIndex() uint32 {
	if m != nil {
		return m.AreaNextFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsNextFloodInformation() uint32 {
	if m != nil {
		return m.AsNextFloodInformation
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetAsNextFloodInformationIndex() uint32 {
	if m != nil {
		return m.AsNextFloodInformationIndex
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLastRetransmissionLength() uint32 {
	if m != nil {
		return m.LastRetransmissionLength
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetMaximumRetransmissionLength() uint32 {
	if m != nil {
		return m.MaximumRetransmissionLength
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLastRetransmissionTime() uint32 {
	if m != nil {
		return m.LastRetransmissionTime
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetMaximumRetransmissionTime() uint32 {
	if m != nil {
		return m.MaximumRetransmissionTime
	}
	return 0
}

func (m *OspfShNeighborRetrans) GetLsaRetransmissionTimer() uint32 {
	if m != nil {
		return m.LsaRetransmissionTimer
	}
	return 0
}

type OspfShNeighborDetail struct {
	NeighborSummary                       *OspfShNeighbor        `protobuf:"bytes,50,opt,name=neighbor_summary,json=neighborSummary,proto3" json:"neighbor_summary,omitempty"`
	NeighborAreaId                        string                 `protobuf:"bytes,51,opt,name=neighbor_area_id,json=neighborAreaId,proto3" json:"neighbor_area_id,omitempty"`
	StateChangeCount                      uint32                 `protobuf:"varint,52,opt,name=state_change_count,json=stateChangeCount,proto3" json:"state_change_count,omitempty"`
	NeighborCost                          uint32                 `protobuf:"varint,53,opt,name=neighbor_cost,json=neighborCost,proto3" json:"neighbor_cost,omitempty"`
	NeighborFilter                        bool                   `protobuf:"varint,54,opt,name=neighbor_filter,json=neighborFilter,proto3" json:"neighbor_filter,omitempty"`
	NeighborDesignatedRouterAddress       string                 `protobuf:"bytes,55,opt,name=neighbor_designated_router_address,json=neighborDesignatedRouterAddress,proto3" json:"neighbor_designated_router_address,omitempty"`
	NeighborBackupDesignatedRouterAddress string                 `protobuf:"bytes,56,opt,name=neighbor_backup_designated_router_address,json=neighborBackupDesignatedRouterAddress,proto3" json:"neighbor_backup_designated_router_address,omitempty"`
	InterfaceType                         string                 `protobuf:"bytes,57,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	PollInterval                          uint32                 `protobuf:"varint,58,opt,name=poll_interval,json=pollInterval,proto3" json:"poll_interval,omitempty"`
	NextPollInterval                      uint32                 `protobuf:"varint,59,opt,name=next_poll_interval,json=nextPollInterval,proto3" json:"next_poll_interval,omitempty"`
	NeighborOption                        uint32                 `protobuf:"varint,60,opt,name=neighbor_option,json=neighborOption,proto3" json:"neighbor_option,omitempty"`
	PendingEvents                         uint32                 `protobuf:"varint,61,opt,name=pending_events,json=pendingEvents,proto3" json:"pending_events,omitempty"`
	NeighborLlsOption                     uint32                 `protobuf:"varint,62,opt,name=neighbor_lls_option,json=neighborLlsOption,proto3" json:"neighbor_lls_option,omitempty"`
	OobResynchronization                  bool                   `protobuf:"varint,63,opt,name=oob_resynchronization,json=oobResynchronization,proto3" json:"oob_resynchronization,omitempty"`
	NsfRouterState                        string                 `protobuf:"bytes,64,opt,name=nsf_router_state,json=nsfRouterState,proto3" json:"nsf_router_state,omitempty"`
	LastOobTime                           uint32                 `protobuf:"varint,65,opt,name=last_oob_time,json=lastOobTime,proto3" json:"last_oob_time,omitempty"`
	NeighborBfdInformation                *OspfShNeighborBfd     `protobuf:"bytes,66,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	NeighborRetransmissionInformation     *OspfShNeighborRetrans `protobuf:"bytes,67,opt,name=neighbor_retransmission_information,json=neighborRetransmissionInformation,proto3" json:"neighbor_retransmission_information,omitempty"`
	LfaInterface                          string                 `protobuf:"bytes,68,opt,name=lfa_interface,json=lfaInterface,proto3" json:"lfa_interface,omitempty"`
	LfaNextHop                            string                 `protobuf:"bytes,69,opt,name=lfa_next_hop,json=lfaNextHop,proto3" json:"lfa_next_hop,omitempty"`
	LfaNeighborId                         string                 `protobuf:"bytes,70,opt,name=lfa_neighbor_id,json=lfaNeighborId,proto3" json:"lfa_neighbor_id,omitempty"`
	LfaNeighborRevision                   uint32                 `protobuf:"varint,71,opt,name=lfa_neighbor_revision,json=lfaNeighborRevision,proto3" json:"lfa_neighbor_revision,omitempty"`
	NeighborAckListCount                  uint32                 `protobuf:"varint,72,opt,name=neighbor_ack_list_count,json=neighborAckListCount,proto3" json:"neighbor_ack_list_count,omitempty"`
	NeighborAckListHighWatermark          uint32                 `protobuf:"varint,73,opt,name=neighbor_ack_list_high_watermark,json=neighborAckListHighWatermark,proto3" json:"neighbor_ack_list_high_watermark,omitempty"`
	AdjacencySidLabel                     uint32                 `protobuf:"varint,74,opt,name=adjacency_sid_label,json=adjacencySidLabel,proto3" json:"adjacency_sid_label,omitempty"`
	AdjacencySidProtected                 bool                   `protobuf:"varint,75,opt,name=adjacency_sid_protected,json=adjacencySidProtected,proto3" json:"adjacency_sid_protected,omitempty"`
	AdjacencySidUnprotectedLabel          uint32                 `protobuf:"varint,76,opt,name=adjacency_sid_unprotected_label,json=adjacencySidUnprotectedLabel,proto3" json:"adjacency_sid_unprotected_label,omitempty"`
	NeighborInterfaceId                   uint32                 `protobuf:"varint,77,opt,name=neighbor_interface_id,json=neighborInterfaceId,proto3" json:"neighbor_interface_id,omitempty"`
	XXX_NoUnkeyedLiteral                  struct{}               `json:"-"`
	XXX_unrecognized                      []byte                 `json:"-"`
	XXX_sizecache                         int32                  `json:"-"`
}

func (m *OspfShNeighborDetail) Reset()         { *m = OspfShNeighborDetail{} }
func (m *OspfShNeighborDetail) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborDetail) ProtoMessage()    {}
func (*OspfShNeighborDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c66ebdf8fd8c36e2, []int{4}
}

func (m *OspfShNeighborDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborDetail.Unmarshal(m, b)
}
func (m *OspfShNeighborDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborDetail.Marshal(b, m, deterministic)
}
func (m *OspfShNeighborDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborDetail.Merge(m, src)
}
func (m *OspfShNeighborDetail) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborDetail.Size(m)
}
func (m *OspfShNeighborDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborDetail.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborDetail proto.InternalMessageInfo

func (m *OspfShNeighborDetail) GetNeighborSummary() *OspfShNeighbor {
	if m != nil {
		return m.NeighborSummary
	}
	return nil
}

func (m *OspfShNeighborDetail) GetNeighborAreaId() string {
	if m != nil {
		return m.NeighborAreaId
	}
	return ""
}

func (m *OspfShNeighborDetail) GetStateChangeCount() uint32 {
	if m != nil {
		return m.StateChangeCount
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborCost() uint32 {
	if m != nil {
		return m.NeighborCost
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborFilter() bool {
	if m != nil {
		return m.NeighborFilter
	}
	return false
}

func (m *OspfShNeighborDetail) GetNeighborDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborDesignatedRouterAddress
	}
	return ""
}

func (m *OspfShNeighborDetail) GetNeighborBackupDesignatedRouterAddress() string {
	if m != nil {
		return m.NeighborBackupDesignatedRouterAddress
	}
	return ""
}

func (m *OspfShNeighborDetail) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *OspfShNeighborDetail) GetPollInterval() uint32 {
	if m != nil {
		return m.PollInterval
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNextPollInterval() uint32 {
	if m != nil {
		return m.NextPollInterval
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborOption() uint32 {
	if m != nil {
		return m.NeighborOption
	}
	return 0
}

func (m *OspfShNeighborDetail) GetPendingEvents() uint32 {
	if m != nil {
		return m.PendingEvents
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborLlsOption() uint32 {
	if m != nil {
		return m.NeighborLlsOption
	}
	return 0
}

func (m *OspfShNeighborDetail) GetOobResynchronization() bool {
	if m != nil {
		return m.OobResynchronization
	}
	return false
}

func (m *OspfShNeighborDetail) GetNsfRouterState() string {
	if m != nil {
		return m.NsfRouterState
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLastOobTime() uint32 {
	if m != nil {
		return m.LastOobTime
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

func (m *OspfShNeighborDetail) GetNeighborRetransmissionInformation() *OspfShNeighborRetrans {
	if m != nil {
		return m.NeighborRetransmissionInformation
	}
	return nil
}

func (m *OspfShNeighborDetail) GetLfaInterface() string {
	if m != nil {
		return m.LfaInterface
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNextHop() string {
	if m != nil {
		return m.LfaNextHop
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNeighborId() string {
	if m != nil {
		return m.LfaNeighborId
	}
	return ""
}

func (m *OspfShNeighborDetail) GetLfaNeighborRevision() uint32 {
	if m != nil {
		return m.LfaNeighborRevision
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborAckListCount() uint32 {
	if m != nil {
		return m.NeighborAckListCount
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborAckListHighWatermark() uint32 {
	if m != nil {
		return m.NeighborAckListHighWatermark
	}
	return 0
}

func (m *OspfShNeighborDetail) GetAdjacencySidLabel() uint32 {
	if m != nil {
		return m.AdjacencySidLabel
	}
	return 0
}

func (m *OspfShNeighborDetail) GetAdjacencySidProtected() bool {
	if m != nil {
		return m.AdjacencySidProtected
	}
	return false
}

func (m *OspfShNeighborDetail) GetAdjacencySidUnprotectedLabel() uint32 {
	if m != nil {
		return m.AdjacencySidUnprotectedLabel
	}
	return 0
}

func (m *OspfShNeighborDetail) GetNeighborInterfaceId() uint32 {
	if m != nil {
		return m.NeighborInterfaceId
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShNeighborDetail_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail_KEYS")
	proto.RegisterType((*OspfShNeighborBfd)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_bfd")
	proto.RegisterType((*OspfShNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor")
	proto.RegisterType((*OspfShNeighborRetrans)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_retrans")
	proto.RegisterType((*OspfShNeighborDetail)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail")
}

func init() { proto.RegisterFile("ospf_sh_neighbor_detail.proto", fileDescriptor_c66ebdf8fd8c36e2) }

var fileDescriptor_c66ebdf8fd8c36e2 = []byte{
	// 1404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x58, 0xdb, 0x72, 0xdb, 0x36,
	0x13, 0x1e, 0xfd, 0x7f, 0x4e, 0x86, 0x2d, 0xc7, 0xa6, 0xed, 0x84, 0xce, 0xc9, 0x8e, 0xf2, 0xe7,
	0xaf, 0x33, 0xd3, 0x51, 0x3b, 0xce, 0xa1, 0x39, 0xa7, 0xb6, 0x63, 0xd7, 0x4e, 0x9c, 0xc3, 0xc8,
	0xc9, 0x34, 0xbd, 0xc2, 0x80, 0x04, 0x28, 0x21, 0x26, 0x09, 0x0e, 0x00, 0xa9, 0x76, 0xaf, 0x7b,
	0xd1, 0x8b, 0x3e, 0x4b, 0xa7, 0xaf, 0x90, 0xe9, 0x73, 0xf4, 0x15, 0xfa, 0x0c, 0x1d, 0x2c, 0x40,
	0x8a, 0x94, 0x28, 0x5f, 0x76, 0xd2, 0x9b, 0x8c, 0xb5, 0xfb, 0xed, 0xee, 0xb7, 0x8b, 0x25, 0x76,
	0x11, 0x74, 0x55, 0xa8, 0x2c, 0xc2, 0xaa, 0x87, 0x53, 0xc6, 0xbb, 0xbd, 0x40, 0x48, 0x4c, 0x99,
	0x26, 0x3c, 0x6e, 0x67, 0x52, 0x68, 0xe1, 0xc9, 0x90, 0xab, 0x50, 0x60, 0x2e, 0x14, 0x3e, 0x92,
	0x98, 0x67, 0x83, 0x3b, 0x18, 0x0c, 0x44, 0xc6, 0x64, 0xdb, 0xfc, 0x65, 0x70, 0x21, 0x53, 0x8a,
	0xa9, 0xfc, 0xaf, 0xf6, 0x40, 0x46, 0xf0, 0x4f, 0x9b, 0xd0, 0x8f, 0x24, 0x64, 0x69, 0x78, 0x8c,
	0x79, 0x1a, 0x09, 0x99, 0x10, 0xcd, 0x45, 0xda, 0x1e, 0x09, 0xa4, 0x46, 0x05, 0xad, 0xdf, 0x1a,
	0xe8, 0xca, 0x04, 0x56, 0xf8, 0xe5, 0xf6, 0x0f, 0x07, 0xde, 0x75, 0x34, 0xe3, 0x62, 0xe1, 0x94,
	0x24, 0xcc, 0x6f, 0xac, 0x36, 0xd6, 0xa6, 0x3a, 0xd3, 0x4e, 0xf6, 0x9a, 0x24, 0xcc, 0x5b, 0x46,
	0xe7, 0x06, 0x32, 0xb2, 0xea, 0xff, 0x80, 0xfa, 0xec, 0x40, 0x46, 0xa0, 0xba, 0x89, 0x66, 0x79,
	0xaa, 0x99, 0x8c, 0x48, 0xc8, 0x2c, 0xe0, 0xbf, 0x00, 0x68, 0x16, 0x52, 0x80, 0xdd, 0x42, 0x73,
	0x45, 0x70, 0x42, 0xa9, 0x64, 0x4a, 0xf9, 0xa7, 0x00, 0x78, 0x3e, 0x97, 0x6f, 0x58, 0x71, 0x4b,
	0xa0, 0xc5, 0x31, 0xbe, 0x41, 0x44, 0xbd, 0xaf, 0xd0, 0x62, 0x10, 0x51, 0xcc, 0x53, 0x1d, 0x61,
	0x96, 0x92, 0x20, 0x66, 0x38, 0x11, 0xd4, 0xf2, 0x6d, 0x76, 0xe6, 0x83, 0x88, 0xee, 0xa5, 0x3a,
	0xda, 0x06, 0xcd, 0x2b, 0x41, 0x99, 0xf7, 0x7f, 0x74, 0xde, 0x18, 0x28, 0x4d, 0x74, 0x5f, 0xe1,
	0x28, 0x26, 0x5d, 0x20, 0xdf, 0xec, 0x34, 0x83, 0x88, 0x1e, 0x80, 0x74, 0x27, 0x26, 0xdd, 0xd6,
	0x5f, 0xa7, 0xd0, 0xdc, 0x68, 0x44, 0x6f, 0x05, 0x4d, 0x17, 0xd1, 0x39, 0x75, 0x45, 0x41, 0xb9,
	0x68, 0x8f, 0x7a, 0x6d, 0xb4, 0x30, 0x9a, 0x11, 0x3e, 0x92, 0xae, 0x3c, 0xf3, 0x23, 0x49, 0x7d,
	0x90, 0xde, 0x3d, 0x74, 0x71, 0xe8, 0xb0, 0xae, 0x62, 0x4b, 0x85, 0xf3, 0x4a, 0xe5, 0xbe, 0x46,
	0x8b, 0xc3, 0x63, 0x93, 0x38, 0x93, 0x5c, 0x48, 0xae, 0x8f, 0xa1, 0x7a, 0xcd, 0x8e, 0x97, 0xeb,
	0x9e, 0xcb, 0xb7, 0x4e, 0x63, 0x8e, 0xa4, 0xb0, 0x30, 0xc9, 0x33, 0xff, 0xb4, 0x3d, 0x92, 0x5c,
	0x6a, 0x72, 0x67, 0xde, 0x2a, 0x9a, 0xa1, 0x12, 0x07, 0x34, 0x07, 0x9d, 0xb1, 0x29, 0x52, 0xb9,
	0x49, 0x1d, 0xa2, 0x9c, 0x22, 0x65, 0x84, 0x62, 0xcd, 0x13, 0x26, 0xfd, 0xb3, 0xb6, 0xe0, 0x45,
	0x64, 0x46, 0xe8, 0x3b, 0xa3, 0xf0, 0xd6, 0x4a, 0x87, 0xdc, 0xcf, 0x00, 0xed, 0x9f, 0x03, 0x70,
	0x41, 0xe8, 0x7d, 0x66, 0xa0, 0x95, 0x62, 0x24, 0x84, 0x7e, 0x1c, 0x56, 0xc4, 0x9f, 0x5a, 0x6d,
	0xac, 0x9d, 0x1b, 0x16, 0xe3, 0x15, 0xa1, 0x1f, 0x8b, 0x82, 0x78, 0x7f, 0x34, 0x90, 0x5f, 0x6e,
	0x8a, 0xf2, 0xa7, 0xe0, 0xa3, 0xd5, 0xc6, 0xda, 0xf4, 0xfa, 0x2f, 0x8d, 0xf6, 0x3f, 0xff, 0x95,
	0xb5, 0xeb, 0x3a, 0xb6, 0x73, 0x21, 0xff, 0xb5, 0x69, 0x7a, 0xb3, 0xf0, 0xd6, 0xfa, 0x34, 0x85,
	0xfc, 0x31, 0x03, 0xc9, 0xb4, 0x24, 0xa9, 0xf2, 0xee, 0x23, 0x9f, 0x06, 0x34, 0xff, 0x99, 0x70,
	0xa5, 0xb8, 0x48, 0x71, 0x28, 0xfa, 0xa9, 0x76, 0xad, 0x7e, 0x81, 0x06, 0xb4, 0x53, 0x51, 0x6f,
	0x19, 0xad, 0xb7, 0x85, 0xae, 0xd5, 0x58, 0x6a, 0xa1, 0x49, 0xec, 0xec, 0x6d, 0xfb, 0x5f, 0x1e,
	0xb3, 0x7f, 0x67, 0x30, 0xd6, 0x49, 0x1b, 0x2d, 0x10, 0xc9, 0x08, 0x8e, 0x62, 0x21, 0x28, 0x4f,
	0xbb, 0x98, 0xa7, 0x94, 0x1d, 0x41, 0x8b, 0x36, 0x3b, 0xf3, 0x46, 0xb5, 0xe3, 0x34, 0x7b, 0x46,
	0xe1, 0xfd, 0x0f, 0xcd, 0x12, 0x65, 0xd1, 0x0e, 0x6a, 0x1b, 0x73, 0x86, 0x28, 0x00, 0x5a, 0xd4,
	0x26, 0xba, 0x3a, 0x9a, 0x68, 0x35, 0xb3, 0xd3, 0x96, 0x59, 0x0e, 0xaa, 0x4b, 0xef, 0x21, 0x5a,
	0x4e, 0xfb, 0x49, 0xc0, 0x24, 0x16, 0xd1, 0x88, 0x13, 0x05, 0xcd, 0xdb, 0xec, 0x5c, 0xb4, 0x80,
	0x37, 0x51, 0xd5, 0x5e, 0x79, 0xcf, 0xd0, 0x15, 0x9b, 0x15, 0x97, 0x4a, 0x17, 0x6c, 0x87, 0xad,
	0x63, 0x5b, 0x7a, 0x19, 0xd2, 0x33, 0x10, 0x47, 0xbd, 0x00, 0x78, 0x2f, 0x50, 0xeb, 0x24, 0x07,
	0x2e, 0x75, 0xdb, 0xec, 0xd7, 0x26, 0xba, 0xb1, 0xc5, 0x78, 0x84, 0x2e, 0x99, 0x92, 0x4d, 0xa0,
	0x32, 0x65, 0x33, 0x21, 0xaa, 0x9e, 0xc8, 0x0e, 0x5a, 0x9d, 0x6c, 0xec, 0x68, 0x20, 0x70, 0x71,
	0x65, 0x82, 0x0b, 0x4b, 0xe2, 0x09, 0xba, 0x0c, 0x09, 0xa5, 0xec, 0xa8, 0x8e, 0xc5, 0x34, 0xb8,
	0xf0, 0x0d, 0xe4, 0x35, 0x3b, 0x1a, 0xa7, 0xb1, 0x8b, 0xae, 0x9f, 0x60, 0xee, 0x78, 0xcc, 0x80,
	0x93, 0xab, 0x93, 0x9c, 0x58, 0x22, 0x0f, 0xd0, 0x32, 0x51, 0x93, 0x68, 0x34, 0x6d, 0xc3, 0x13,
	0x55, 0x4b, 0xe2, 0x39, 0x5a, 0x99, 0x68, 0xea, 0x28, 0xcc, 0xda, 0xbe, 0xaa, 0x77, 0x60, 0x09,
	0x3c, 0x46, 0x97, 0x62, 0xa2, 0xf4, 0x68, 0x5f, 0xc6, 0x2c, 0xed, 0xea, 0x9e, 0x7f, 0xde, 0x16,
	0xc2, 0x20, 0xaa, 0x4d, 0xb5, 0x0f, 0x7a, 0xd3, 0xd9, 0x09, 0x39, 0xe2, 0x49, 0x3f, 0x99, 0xe0,
	0x60, 0xce, 0x32, 0x70, 0xa0, 0x5a, 0x1f, 0xf7, 0x91, 0x5f, 0xc7, 0x00, 0xee, 0xcf, 0x79, 0x5b,
	0x81, 0xf1, 0xf8, 0x70, 0x8f, 0x3e, 0x45, 0x97, 0x27, 0x44, 0x07, 0x63, 0xcf, 0xb6, 0x75, 0x6d,
	0x6c, 0xb0, 0x37, 0x91, 0x15, 0xa9, 0xb3, 0x95, 0xfe, 0x82, 0x8b, 0xac, 0xc8, 0xb8, 0xa1, 0x6c,
	0x7d, 0x9a, 0x45, 0x17, 0x27, 0xac, 0x15, 0xde, 0xef, 0x8d, 0xd2, 0x20, 0x50, 0xfd, 0x24, 0x21,
	0xf2, 0xd8, 0x5f, 0x87, 0xdb, 0xf9, 0xe7, 0xcf, 0xe2, 0x76, 0x1e, 0x2e, 0x1d, 0x07, 0x96, 0x5d,
	0x65, 0x74, 0x41, 0x63, 0x73, 0xea, 0xdf, 0x86, 0x81, 0x58, 0x8c, 0xae, 0x0d, 0xc9, 0xc8, 0x1e,
	0xf5, 0xbe, 0x44, 0x1e, 0xcc, 0x4b, 0x1c, 0xf6, 0x48, 0xda, 0x65, 0xee, 0xfe, 0xba, 0x03, 0xc5,
	0x9a, 0x03, 0xcd, 0x16, 0x28, 0xec, 0xa5, 0x75, 0x03, 0x15, 0x53, 0x17, 0x87, 0x42, 0x69, 0xff,
	0xae, 0xbd, 0x1d, 0x73, 0xe1, 0x96, 0x50, 0xda, 0xfb, 0x02, 0x15, 0x7c, 0x70, 0xc4, 0x63, 0xcd,
	0xa4, 0x7f, 0x0f, 0xa6, 0x60, 0x11, 0x7b, 0x07, 0xa4, 0xde, 0x4b, 0xd4, 0x2a, 0xa5, 0xa6, 0x78,
	0x37, 0x25, 0x9a, 0x51, 0x2c, 0x45, 0x5f, 0xb3, 0xe1, 0x5e, 0xf5, 0x0d, 0xf0, 0x5e, 0x19, 0xce,
	0xe7, 0x1c, 0xd8, 0x01, 0x9c, 0x5b, 0x49, 0xbc, 0x0f, 0xe8, 0xd6, 0x70, 0x5a, 0x91, 0xf0, 0xb0,
	0x9f, 0x9d, 0xe0, 0xf3, 0x3e, 0xf8, 0xbc, 0x59, 0x0c, 0x34, 0xc0, 0x4f, 0xf2, 0x5c, 0xd9, 0x09,
	0xf5, 0x71, 0xc6, 0xfc, 0x07, 0x23, 0x3b, 0xe1, 0xbb, 0xe3, 0x8c, 0x99, 0xda, 0x64, 0x22, 0x8e,
	0xed, 0xec, 0x1f, 0x90, 0xd8, 0x7f, 0x68, 0x6b, 0x63, 0x84, 0x7b, 0x4e, 0x66, 0xca, 0x0d, 0x1f,
	0x78, 0x15, 0xf9, 0xc8, 0x96, 0xdb, 0x68, 0xde, 0x96, 0xd1, 0xe5, 0x4a, 0x8a, 0x0c, 0xae, 0x90,
	0xc7, 0xd5, 0x05, 0xe4, 0x0d, 0x48, 0x0d, 0xc5, 0x8c, 0xa5, 0x30, 0xe0, 0xd8, 0x80, 0xa5, 0x5a,
	0xf9, 0x4f, 0xec, 0x6a, 0xe8, 0xa4, 0xdb, 0x20, 0xac, 0x6c, 0x40, 0x71, 0xac, 0x72, 0x9f, 0x4f,
	0xab, 0x1b, 0xd0, 0x7e, 0xac, 0x9c, 0xdb, 0xdb, 0x68, 0x49, 0x88, 0x00, 0x4b, 0xa6, 0x8e, 0xd3,
	0xb0, 0x27, 0x45, 0xca, 0x7f, 0xb2, 0x17, 0xd9, 0x33, 0x38, 0xcf, 0x45, 0x21, 0x82, 0xce, 0xa8,
	0x0e, 0x7a, 0x4f, 0x45, 0x79, 0xc5, 0xed, 0x32, 0xf6, 0xad, 0xeb, 0x3d, 0x15, 0xd9, 0xd2, 0xda,
	0x85, 0xac, 0x85, 0x9a, 0x70, 0x51, 0x98, 0x18, 0xf0, 0x81, 0x6f, 0x00, 0x91, 0x69, 0x23, 0x7c,
	0x23, 0x02, 0xf8, 0xa4, 0x4f, 0x5c, 0x91, 0x36, 0xff, 0x25, 0x2b, 0x92, 0xf7, 0x67, 0x03, 0xdd,
	0x98, 0xb4, 0x31, 0x94, 0x13, 0xda, 0x82, 0x84, 0x7e, 0xfd, 0x3c, 0x12, 0x72, 0x3c, 0x3b, 0xd7,
	0xeb, 0xd7, 0x98, 0x72, 0x7e, 0x37, 0x50, 0x33, 0x8e, 0x48, 0x69, 0xed, 0x7d, 0x0e, 0x07, 0x3e,
	0x13, 0x47, 0x64, 0xb8, 0xed, 0xae, 0x22, 0xf3, 0xdb, 0x0e, 0xb8, 0x9e, 0xc8, 0xfc, 0x6d, 0xbb,
	0xa1, 0xc7, 0x11, 0x8c, 0xd3, 0x5d, 0x91, 0x99, 0x27, 0x8e, 0x45, 0x0c, 0x5f, 0x2a, 0x3b, 0xf6,
	0x53, 0x03, 0x50, 0xf1, 0x58, 0x59, 0x47, 0x4b, 0x15, 0x9c, 0x64, 0x03, 0x6e, 0x28, 0xf9, 0xdf,
	0x41, 0x03, 0x2d, 0x94, 0xd0, 0x1d, 0xa7, 0xf2, 0xee, 0x96, 0x76, 0x74, 0x12, 0x1e, 0xe2, 0x98,
	0x2b, 0xed, 0x6e, 0xbb, 0x5d, 0xb0, 0x2a, 0xde, 0x25, 0x1b, 0xe1, 0xe1, 0x3e, 0x57, 0xda, 0xde,
	0x78, 0x3b, 0x68, 0x75, 0xdc, 0xac, 0xc7, 0xbb, 0x3d, 0xfc, 0x23, 0xd1, 0x4c, 0x26, 0x44, 0x1e,
	0xfa, 0x7b, 0x76, 0x41, 0x19, 0xb1, 0xdf, 0xe5, 0xdd, 0xde, 0xf7, 0x39, 0x06, 0x16, 0xd1, 0xe2,
	0x5c, 0x14, 0xa7, 0x38, 0x26, 0x01, 0x8b, 0xfd, 0x17, 0x6e, 0x11, 0xcd, 0x55, 0x07, 0x9c, 0xee,
	0x1b, 0x85, 0x79, 0x52, 0x54, 0xf1, 0xe6, 0xc9, 0xcd, 0x42, 0xcd, 0xa8, 0xff, 0xd2, 0x3e, 0x29,
	0xca, 0x36, 0x6f, 0x73, 0xa5, 0xb7, 0x8d, 0x56, 0xaa, 0x76, 0xfd, 0xb4, 0xb0, 0x74, 0x31, 0xf7,
	0xdd, 0x3e, 0x55, 0xb2, 0x7f, 0x3f, 0x04, 0xd9, 0xf0, 0xeb, 0x68, 0xa9, 0xe6, 0x79, 0xc7, 0xa9,
	0xff, 0xca, 0x56, 0x78, 0xec, 0x71, 0xb7, 0x47, 0x83, 0x33, 0xf0, 0xbf, 0x02, 0xb7, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x80, 0x10, 0xd2, 0x9b, 0x36, 0x10, 0x00, 0x00,
}
