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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_areas_area_neighbor_details_neighbor_detail

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
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
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

func (m *OspfShNeighborDetail_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
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
	NeighborIpAddress      string             `protobuf:"bytes,2,opt,name=neighbor_ip_address,json=neighborIpAddress,proto3" json:"neighbor_ip_address,omitempty"`
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

func (m *OspfShNeighbor) GetNeighborIpAddress() string {
	if m != nil {
		return m.NeighborIpAddress
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
	proto.RegisterType((*OspfShNeighborDetail_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail_KEYS")
	proto.RegisterType((*OspfShNeighborBfd)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.neighbor_details.neighbor_detail.ospf_sh_neighbor_bfd")
	proto.RegisterType((*OspfShNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.neighbor_details.neighbor_detail.ospf_sh_neighbor")
	proto.RegisterType((*OspfShNeighborRetrans)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.neighbor_details.neighbor_detail.ospf_sh_neighbor_retrans")
	proto.RegisterType((*OspfShNeighborDetail)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.neighbor_details.neighbor_detail.ospf_sh_neighbor_detail")
}

func init() { proto.RegisterFile("ospf_sh_neighbor_detail.proto", fileDescriptor_c66ebdf8fd8c36e2) }

var fileDescriptor_c66ebdf8fd8c36e2 = []byte{
	// 1404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xcd, 0x72, 0xdb, 0x36,
	0x10, 0x1e, 0xa5, 0xf9, 0x33, 0x6c, 0x39, 0x36, 0x6d, 0xc7, 0xcc, 0xaf, 0x15, 0xa5, 0x69, 0x9d,
	0x99, 0x8e, 0xda, 0x71, 0x7e, 0x9a, 0xff, 0xd4, 0x76, 0xec, 0x5a, 0x89, 0xf3, 0x33, 0x72, 0x32,
	0x6d, 0x4f, 0x18, 0x50, 0x00, 0x25, 0xd8, 0x24, 0xc1, 0x01, 0x20, 0xd7, 0xee, 0xf4, 0xd0, 0x43,
	0xdf, 0xa5, 0x3d, 0xf4, 0x0d, 0x7a, 0xea, 0x43, 0xf4, 0x15, 0xfa, 0x1c, 0x1d, 0x2c, 0x40, 0x8a,
	0x94, 0x28, 0x1f, 0x9b, 0x5e, 0x34, 0xe2, 0xee, 0xb7, 0xbb, 0xdf, 0x2e, 0x97, 0xd8, 0x05, 0xba,
	0x26, 0x54, 0x1a, 0x62, 0xd5, 0xc7, 0x09, 0xe3, 0xbd, 0x7e, 0x20, 0x24, 0xa6, 0x4c, 0x13, 0x1e,
	0xb5, 0x52, 0x29, 0xb4, 0xf0, 0xf6, 0xbb, 0x5c, 0x75, 0x05, 0xe6, 0x42, 0xe1, 0x23, 0x89, 0x79,
	0x7a, 0x78, 0x17, 0x83, 0x81, 0x48, 0x99, 0x6c, 0x99, 0x7f, 0x06, 0xd7, 0x65, 0x4a, 0x31, 0x95,
	0xfd, 0x6b, 0x51, 0x16, 0x92, 0x41, 0xa4, 0xf1, 0xa1, 0x0c, 0x5b, 0x44, 0x32, 0xa2, 0xe0, 0xb7,
	0x35, 0x12, 0x40, 0x8d, 0x0a, 0x9a, 0x7f, 0xd4, 0xd0, 0xd5, 0x09, 0x6c, 0xf0, 0xab, 0xad, 0x1f,
	0xf6, 0xbc, 0x1b, 0x68, 0xc6, 0xc5, 0xc0, 0x09, 0x89, 0x99, 0x5f, 0x6b, 0xd4, 0x56, 0xa7, 0x3a,
	0xd3, 0x4e, 0xf6, 0x86, 0xc4, 0xcc, 0x5b, 0x46, 0xe7, 0x4c, 0x30, 0xcc, 0xa9, 0x7f, 0xaa, 0x51,
	0x5b, 0xad, 0x77, 0xce, 0x9a, 0xc7, 0x36, 0xf5, 0x6e, 0xa1, 0x59, 0x9e, 0x68, 0x26, 0x43, 0xd2,
	0x65, 0xd6, 0xfa, 0x13, 0xb0, 0xae, 0xe7, 0x52, 0xb0, 0xbf, 0x8d, 0xe6, 0xf2, 0xd0, 0x84, 0x52,
	0xc9, 0x94, 0xf2, 0x4f, 0x03, 0xf0, 0x42, 0x26, 0x5f, 0xb7, 0xe2, 0xa6, 0x40, 0x8b, 0x63, 0x6c,
	0x83, 0x90, 0x7a, 0x5f, 0xa2, 0xc5, 0x20, 0xa4, 0x98, 0x27, 0x3a, 0xc4, 0x2c, 0x21, 0x41, 0xc4,
	0x70, 0x2c, 0xa8, 0x65, 0x5b, 0xef, 0xcc, 0x07, 0x21, 0x6d, 0x27, 0x3a, 0xdc, 0x02, 0xcd, 0x6b,
	0x41, 0x99, 0xf7, 0x19, 0xba, 0x60, 0x0c, 0x94, 0x26, 0x7a, 0xa0, 0x70, 0x18, 0x91, 0x9e, 0xe3,
	0x5e, 0x0f, 0x42, 0xba, 0x07, 0xd2, 0xed, 0x88, 0xf4, 0x9a, 0xff, 0x9c, 0x46, 0x73, 0xa3, 0x11,
	0xbd, 0x15, 0x34, 0x9d, 0x47, 0xe7, 0xd4, 0x95, 0x04, 0x65, 0xa2, 0x36, 0xf5, 0x5a, 0x68, 0x61,
	0x08, 0x48, 0xf3, 0xa4, 0x4e, 0x01, 0x70, 0x3e, 0x07, 0xa6, 0x2e, 0x2d, 0xef, 0x3e, 0x5a, 0x1e,
	0xe2, 0xab, 0x2a, 0xb6, 0x94, 0xdb, 0x94, 0x2a, 0xf7, 0x15, 0x5a, 0x1c, 0xbe, 0x34, 0x89, 0x53,
	0xc9, 0x85, 0xe4, 0xfa, 0x18, 0xaa, 0x57, 0xef, 0x78, 0x99, 0xee, 0x85, 0x7c, 0xe7, 0x34, 0xe6,
	0x95, 0xe4, 0x16, 0x26, 0x79, 0xe6, 0x9f, 0xb1, 0xaf, 0x24, 0x93, 0x9a, 0xdc, 0x99, 0xd7, 0x40,
	0x33, 0x54, 0xe2, 0x80, 0x66, 0xa0, 0xb3, 0x36, 0x45, 0x2a, 0x37, 0xa8, 0x43, 0x14, 0x53, 0xa4,
	0x8c, 0x50, 0xac, 0x79, 0xcc, 0xa4, 0x7f, 0xce, 0x16, 0x3c, 0x8f, 0xcc, 0x08, 0x7d, 0x6f, 0x14,
	0xde, 0x6a, 0xe1, 0x25, 0x0f, 0x52, 0x40, 0xfb, 0xe7, 0x01, 0x9c, 0x13, 0xfa, 0x90, 0x1a, 0x68,
	0xa9, 0x18, 0x31, 0xa1, 0xfb, 0xc3, 0x8a, 0xf8, 0x53, 0x8d, 0xda, 0xea, 0xf9, 0x61, 0x31, 0x5e,
	0x13, 0xba, 0x9f, 0x17, 0xc4, 0xfb, 0xb3, 0x86, 0xfc, 0x62, 0x53, 0x60, 0x9e, 0x84, 0x42, 0xc6,
	0x44, 0x73, 0x91, 0xf8, 0xa8, 0x51, 0x5b, 0x9d, 0x5e, 0xfb, 0xa5, 0xd6, 0xfa, 0xef, 0xbe, 0xad,
	0x56, 0x55, 0xa7, 0x76, 0x2e, 0x66, 0x4f, 0x1b, 0xa6, 0x27, 0x73, 0x82, 0xcd, 0xbf, 0xa6, 0x90,
	0x3f, 0x66, 0x20, 0x99, 0x96, 0x24, 0x51, 0xde, 0x03, 0xe4, 0xd3, 0x80, 0x66, 0x8f, 0x31, 0x57,
	0x8a, 0x8b, 0x04, 0x77, 0xc5, 0x20, 0xd1, 0xae, 0xc5, 0x2f, 0xd2, 0x80, 0x76, 0x4a, 0xea, 0x4d,
	0xa3, 0xf5, 0x36, 0xd1, 0xf5, 0x0a, 0x4b, 0x2d, 0x34, 0x89, 0x9c, 0xbd, 0x6d, 0xfb, 0x2b, 0x63,
	0xf6, 0xef, 0x0d, 0xc6, 0x3a, 0x69, 0xa1, 0x05, 0xf8, 0xc0, 0xc3, 0x48, 0x08, 0xca, 0x93, 0x1e,
	0xe6, 0x09, 0x65, 0x47, 0xd0, 0x9a, 0xf5, 0xce, 0xbc, 0x51, 0x6d, 0x3b, 0x4d, 0xdb, 0x28, 0xbc,
	0x4f, 0xd1, 0x2c, 0x51, 0x16, 0xed, 0xa0, 0xb6, 0x21, 0x67, 0x88, 0x02, 0xa0, 0x45, 0x6d, 0xa0,
	0x6b, 0xa3, 0x89, 0x96, 0x33, 0x3b, 0x63, 0x99, 0x65, 0xa0, 0xaa, 0xf4, 0x1e, 0xa1, 0x4b, 0xc9,
	0x20, 0x0e, 0x98, 0xc4, 0x22, 0x1c, 0x71, 0xa2, 0xa0, 0x69, 0xeb, 0x9d, 0x65, 0x0b, 0x78, 0x1b,
	0x96, 0xed, 0x95, 0xf7, 0x1c, 0x5d, 0xb5, 0x59, 0x71, 0xa9, 0x74, 0xce, 0x76, 0xd8, 0x32, 0xb6,
	0x95, 0x2f, 0x41, 0x7a, 0x06, 0xe2, 0xa8, 0xe7, 0x00, 0xef, 0x25, 0x6a, 0x9e, 0xe4, 0xc0, 0xa5,
	0x6e, 0x9b, 0xfc, 0xfa, 0x44, 0x37, 0xb6, 0x18, 0x8f, 0xd1, 0x65, 0x53, 0xb2, 0x09, 0x54, 0xa6,
	0x6c, 0x26, 0x44, 0x55, 0x13, 0xd9, 0x46, 0x8d, 0xc9, 0xc6, 0x8e, 0x06, 0x02, 0x17, 0x57, 0x27,
	0xb8, 0xb0, 0x24, 0x9e, 0xa2, 0x2b, 0x90, 0x50, 0xc2, 0x8e, 0xaa, 0x58, 0x4c, 0x83, 0x0b, 0xdf,
	0x40, 0xde, 0xb0, 0xa3, 0x71, 0x1a, 0x3b, 0xe8, 0xc6, 0x09, 0xe6, 0x8e, 0xc7, 0x0c, 0x38, 0xb9,
	0x36, 0xc9, 0x89, 0x25, 0xf2, 0x10, 0x5d, 0x22, 0x6a, 0x12, 0x8d, 0xba, 0x6d, 0x78, 0xa2, 0x2a,
	0x49, 0xbc, 0x40, 0x2b, 0x13, 0x4d, 0x1d, 0x85, 0x59, 0xdb, 0x57, 0xd5, 0x0e, 0x2c, 0x81, 0x27,
	0xe8, 0x72, 0x44, 0x94, 0x1e, 0xed, 0xcb, 0x88, 0x25, 0x3d, 0xdd, 0xf7, 0x2f, 0xd8, 0x42, 0x18,
	0x44, 0xb9, 0xa9, 0x76, 0x41, 0x6f, 0x3a, 0x3b, 0x26, 0x47, 0x3c, 0x1e, 0xc4, 0x13, 0x1c, 0xcc,
	0x59, 0x06, 0x0e, 0x54, 0xe9, 0xe3, 0x01, 0xf2, 0xab, 0x18, 0xc0, 0xb9, 0x39, 0x6f, 0x2b, 0x30,
	0x1e, 0x1f, 0xce, 0xcf, 0x67, 0xe8, 0xca, 0x84, 0xe8, 0x60, 0xec, 0xd9, 0xb6, 0xae, 0x8c, 0x0d,
	0xf6, 0x26, 0xb2, 0x22, 0x55, 0xb6, 0xd2, 0x5f, 0x70, 0x91, 0x15, 0x19, 0x37, 0x94, 0xcd, 0xdf,
	0x67, 0xd1, 0xf2, 0x84, 0x65, 0xc2, 0xfb, 0xad, 0x56, 0x18, 0x00, 0x6a, 0x10, 0xc7, 0x44, 0x1e,
	0xfb, 0x6b, 0x70, 0x2a, 0xff, 0xfc, 0x31, 0x0f, 0xe5, 0xe1, 0x8e, 0xb1, 0x67, 0x49, 0x95, 0x26,
	0x55, 0xb6, 0xd7, 0xdc, 0x81, 0xf9, 0x97, 0x4f, 0xaa, 0x75, 0xbb, 0xdf, 0x7c, 0x81, 0x3c, 0x18,
	0x8f, 0xb8, 0xdb, 0x27, 0x49, 0x8f, 0xb9, 0x63, 0xeb, 0x2e, 0xd4, 0x68, 0x0e, 0x34, 0x9b, 0xa0,
	0xb0, 0x67, 0xd5, 0x4d, 0x94, 0x0f, 0x59, 0xdc, 0x15, 0x4a, 0xfb, 0xf7, 0xec, 0xa1, 0x98, 0x09,
	0x37, 0x85, 0xd2, 0xde, 0xe7, 0x28, 0xe7, 0x83, 0x43, 0x1e, 0x69, 0x26, 0xfd, 0xfb, 0x30, 0xf4,
	0xf2, 0xd8, 0xdb, 0x20, 0xf5, 0x5e, 0xa1, 0x66, 0x21, 0x35, 0xc5, 0x7b, 0x09, 0xd1, 0x8c, 0x62,
	0x29, 0x06, 0x9a, 0x0d, 0xd7, 0xa8, 0xaf, 0x81, 0xf7, 0xca, 0x70, 0x1c, 0x67, 0xc0, 0x0e, 0xe0,
	0xb2, 0xfd, 0xe3, 0x7b, 0x74, 0x7b, 0x38, 0xa4, 0x48, 0xf7, 0x60, 0x90, 0x9e, 0xe0, 0xf3, 0x01,
	0xf8, 0xbc, 0x95, 0xcf, 0x31, 0xc0, 0x4f, 0xf2, 0x5c, 0x5a, 0x01, 0xf5, 0x71, 0xca, 0xfc, 0x87,
	0x23, 0x2b, 0xe0, 0xfb, 0xe3, 0x94, 0x99, 0xda, 0xa4, 0x22, 0x8a, 0xec, 0xa8, 0x3f, 0x24, 0x91,
	0xff, 0xc8, 0xd6, 0xc6, 0x08, 0xdb, 0x4e, 0x66, 0xca, 0x0d, 0xdf, 0x75, 0x19, 0xf9, 0xd8, 0x96,
	0xdb, 0x68, 0xde, 0x15, 0xd1, 0xc5, 0x4a, 0x8a, 0x14, 0x4e, 0x8e, 0x27, 0xe5, 0x7d, 0xe3, 0x2d,
	0x48, 0x0d, 0xc5, 0x94, 0x25, 0x30, 0xd7, 0xd8, 0x21, 0x4b, 0xb4, 0xf2, 0x9f, 0xda, 0x4d, 0xd0,
	0x49, 0xb7, 0x40, 0x58, 0x5a, 0x78, 0xa2, 0x48, 0x65, 0x3e, 0x9f, 0x95, 0x17, 0x9e, 0xdd, 0x48,
	0x39, 0xb7, 0x77, 0xd0, 0x92, 0x10, 0x01, 0x96, 0x4c, 0x1d, 0x27, 0xdd, 0xbe, 0x14, 0x09, 0xff,
	0xc9, 0x9e, 0x5f, 0xcf, 0xe1, 0x7d, 0x2e, 0x0a, 0x11, 0x74, 0x46, 0x75, 0xd0, 0x7b, 0x2a, 0xcc,
	0x2a, 0x6e, 0x77, 0xaf, 0x6f, 0x5c, 0xef, 0xa9, 0xd0, 0x96, 0xd6, 0xee, 0x5f, 0x4d, 0x54, 0x87,
	0xf3, 0xc1, 0xc4, 0x80, 0xef, 0x7a, 0x1d, 0x88, 0x4c, 0x1b, 0xe1, 0x5b, 0x11, 0xc0, 0x97, 0x7c,
	0xe2, 0x46, 0xb4, 0xf1, 0x3f, 0xdf, 0x88, 0xbc, 0xbf, 0x6b, 0xe8, 0xe6, 0xa4, 0x05, 0xa1, 0x98,
	0xc8, 0x26, 0x24, 0xf2, 0xeb, 0xc7, 0x4d, 0xc4, 0xf1, 0xeb, 0xdc, 0xa8, 0xde, 0x56, 0x8a, 0x79,
	0xdd, 0x44, 0xf5, 0x28, 0x24, 0x85, 0xad, 0xf6, 0x05, 0xbc, 0xe0, 0x99, 0x28, 0x24, 0xc3, 0x65,
	0xb6, 0x81, 0xcc, 0xb3, 0x9d, 0x63, 0x7d, 0x91, 0xfa, 0x5b, 0x76, 0x01, 0x8f, 0x42, 0x98, 0x9a,
	0x3b, 0x22, 0x35, 0x37, 0x18, 0x8b, 0x18, 0x5e, 0x44, 0xb6, 0xed, 0xa7, 0x05, 0xa0, 0xfc, 0x2e,
	0xb2, 0x86, 0x96, 0x4a, 0x38, 0xc9, 0x0e, 0xb9, 0xa1, 0xe4, 0x7f, 0x0b, 0x0d, 0xb3, 0x50, 0x40,
	0x77, 0x9c, 0xca, 0xbb, 0x57, 0x58, 0xc1, 0x49, 0xf7, 0x00, 0x47, 0x5c, 0x69, 0x77, 0xba, 0xed,
	0x80, 0x55, 0x7e, 0xed, 0x58, 0xef, 0x1e, 0xec, 0x72, 0xa5, 0xed, 0x09, 0xb7, 0x8d, 0x1a, 0xe3,
	0x66, 0x7d, 0xde, 0xeb, 0xe3, 0x1f, 0x89, 0x66, 0x32, 0x26, 0xf2, 0xc0, 0x6f, 0xdb, 0x3d, 0x64,
	0xc4, 0x7e, 0x87, 0xf7, 0xfa, 0xdf, 0x65, 0x18, 0xd8, 0x37, 0xe9, 0x3e, 0xe9, 0xb2, 0xa4, 0x7b,
	0x8c, 0x15, 0xa7, 0x38, 0x22, 0x01, 0x8b, 0xfc, 0x97, 0x6e, 0xdf, 0xcc, 0x54, 0x7b, 0x9c, 0xee,
	0x1a, 0x85, 0xb9, 0x31, 0x94, 0xf1, 0xe6, 0x1e, 0xcd, 0xba, 0x9a, 0x51, 0xff, 0x95, 0xbd, 0x31,
	0x14, 0x6d, 0xde, 0x65, 0x4a, 0x6f, 0x0b, 0xad, 0x94, 0xed, 0x06, 0x49, 0x6e, 0xe9, 0x62, 0xee,
	0xba, 0xb5, 0xa9, 0x60, 0xff, 0x61, 0x08, 0xb2, 0xe1, 0xd7, 0xd0, 0x52, 0xc5, 0xed, 0x8d, 0x53,
	0xff, 0xb5, 0xad, 0xf0, 0xd8, 0xdd, 0xad, 0x4d, 0x83, 0xb3, 0x70, 0xd5, 0xbf, 0xf3, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x55, 0xc6, 0x6d, 0x2b, 0x0b, 0x10, 0x00, 0x00,
}