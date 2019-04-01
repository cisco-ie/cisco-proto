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
// source: bm_member_data.proto

package cisco_ios_xr_bundlemgr_oper_lacp_bundles_bundles_bundle_members_member

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

type BmMemberData_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	MemberInterface      string   `protobuf:"bytes,2,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmMemberData_KEYS) Reset()         { *m = BmMemberData_KEYS{} }
func (m *BmMemberData_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmMemberData_KEYS) ProtoMessage()    {}
func (*BmMemberData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{0}
}

func (m *BmMemberData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMemberData_KEYS.Unmarshal(m, b)
}
func (m *BmMemberData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMemberData_KEYS.Marshal(b, m, deterministic)
}
func (m *BmMemberData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMemberData_KEYS.Merge(m, src)
}
func (m *BmMemberData_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmMemberData_KEYS.Size(m)
}
func (m *BmMemberData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMemberData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmMemberData_KEYS proto.InternalMessageInfo

func (m *BmMemberData_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

func (m *BmMemberData_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
}

type BmLacpCounters struct {
	LacpdUsReceived             uint32   `protobuf:"varint,1,opt,name=lacpd_us_received,json=lacpdUsReceived,proto3" json:"lacpd_us_received,omitempty"`
	LacpdUsTransmitted          uint32   `protobuf:"varint,2,opt,name=lacpd_us_transmitted,json=lacpdUsTransmitted,proto3" json:"lacpd_us_transmitted,omitempty"`
	MarkerPacketsReceived       uint32   `protobuf:"varint,3,opt,name=marker_packets_received,json=markerPacketsReceived,proto3" json:"marker_packets_received,omitempty"`
	MarkerResponsesTransmitted  uint32   `protobuf:"varint,4,opt,name=marker_responses_transmitted,json=markerResponsesTransmitted,proto3" json:"marker_responses_transmitted,omitempty"`
	IllegalPacketsReceived      uint32   `protobuf:"varint,5,opt,name=illegal_packets_received,json=illegalPacketsReceived,proto3" json:"illegal_packets_received,omitempty"`
	ExcessLacpdUsReceived       uint32   `protobuf:"varint,6,opt,name=excess_lacpd_us_received,json=excessLacpdUsReceived,proto3" json:"excess_lacpd_us_received,omitempty"`
	ExcessMarkerPacketsReceived uint32   `protobuf:"varint,7,opt,name=excess_marker_packets_received,json=excessMarkerPacketsReceived,proto3" json:"excess_marker_packets_received,omitempty"`
	Defaulted                   uint32   `protobuf:"varint,8,opt,name=defaulted,proto3" json:"defaulted,omitempty"`
	Expired                     uint32   `protobuf:"varint,9,opt,name=expired,proto3" json:"expired,omitempty"`
	LastClearedSec              uint32   `protobuf:"varint,10,opt,name=last_cleared_sec,json=lastClearedSec,proto3" json:"last_cleared_sec,omitempty"`
	LastClearedNsec             uint32   `protobuf:"varint,11,opt,name=last_cleared_nsec,json=lastClearedNsec,proto3" json:"last_cleared_nsec,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *BmLacpCounters) Reset()         { *m = BmLacpCounters{} }
func (m *BmLacpCounters) String() string { return proto.CompactTextString(m) }
func (*BmLacpCounters) ProtoMessage()    {}
func (*BmLacpCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{1}
}

func (m *BmLacpCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmLacpCounters.Unmarshal(m, b)
}
func (m *BmLacpCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmLacpCounters.Marshal(b, m, deterministic)
}
func (m *BmLacpCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmLacpCounters.Merge(m, src)
}
func (m *BmLacpCounters) XXX_Size() int {
	return xxx_messageInfo_BmLacpCounters.Size(m)
}
func (m *BmLacpCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_BmLacpCounters.DiscardUnknown(m)
}

var xxx_messageInfo_BmLacpCounters proto.InternalMessageInfo

func (m *BmLacpCounters) GetLacpdUsReceived() uint32 {
	if m != nil {
		return m.LacpdUsReceived
	}
	return 0
}

func (m *BmLacpCounters) GetLacpdUsTransmitted() uint32 {
	if m != nil {
		return m.LacpdUsTransmitted
	}
	return 0
}

func (m *BmLacpCounters) GetMarkerPacketsReceived() uint32 {
	if m != nil {
		return m.MarkerPacketsReceived
	}
	return 0
}

func (m *BmLacpCounters) GetMarkerResponsesTransmitted() uint32 {
	if m != nil {
		return m.MarkerResponsesTransmitted
	}
	return 0
}

func (m *BmLacpCounters) GetIllegalPacketsReceived() uint32 {
	if m != nil {
		return m.IllegalPacketsReceived
	}
	return 0
}

func (m *BmLacpCounters) GetExcessLacpdUsReceived() uint32 {
	if m != nil {
		return m.ExcessLacpdUsReceived
	}
	return 0
}

func (m *BmLacpCounters) GetExcessMarkerPacketsReceived() uint32 {
	if m != nil {
		return m.ExcessMarkerPacketsReceived
	}
	return 0
}

func (m *BmLacpCounters) GetDefaulted() uint32 {
	if m != nil {
		return m.Defaulted
	}
	return 0
}

func (m *BmLacpCounters) GetExpired() uint32 {
	if m != nil {
		return m.Expired
	}
	return 0
}

func (m *BmLacpCounters) GetLastClearedSec() uint32 {
	if m != nil {
		return m.LastClearedSec
	}
	return 0
}

func (m *BmLacpCounters) GetLastClearedNsec() uint32 {
	if m != nil {
		return m.LastClearedNsec
	}
	return 0
}

type BmLacpLinkData struct {
	InterfaceHandle         string   `protobuf:"bytes,1,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	ActorSystemPriority     uint32   `protobuf:"varint,2,opt,name=actor_system_priority,json=actorSystemPriority,proto3" json:"actor_system_priority,omitempty"`
	ActorSystemMacAddress   string   `protobuf:"bytes,3,opt,name=actor_system_mac_address,json=actorSystemMacAddress,proto3" json:"actor_system_mac_address,omitempty"`
	ActorOperationalKey     uint32   `protobuf:"varint,4,opt,name=actor_operational_key,json=actorOperationalKey,proto3" json:"actor_operational_key,omitempty"`
	PartnerSystemPriority   uint32   `protobuf:"varint,5,opt,name=partner_system_priority,json=partnerSystemPriority,proto3" json:"partner_system_priority,omitempty"`
	PartnerSystemMacAddress string   `protobuf:"bytes,6,opt,name=partner_system_mac_address,json=partnerSystemMacAddress,proto3" json:"partner_system_mac_address,omitempty"`
	PartnerOperationalKey   uint32   `protobuf:"varint,7,opt,name=partner_operational_key,json=partnerOperationalKey,proto3" json:"partner_operational_key,omitempty"`
	SelectedAggregatorId    uint32   `protobuf:"varint,8,opt,name=selected_aggregator_id,json=selectedAggregatorId,proto3" json:"selected_aggregator_id,omitempty"`
	AttachedAggregatorId    uint32   `protobuf:"varint,9,opt,name=attached_aggregator_id,json=attachedAggregatorId,proto3" json:"attached_aggregator_id,omitempty"`
	ActorPortId             uint32   `protobuf:"varint,10,opt,name=actor_port_id,json=actorPortId,proto3" json:"actor_port_id,omitempty"`
	ActorPortPriority       uint32   `protobuf:"varint,11,opt,name=actor_port_priority,json=actorPortPriority,proto3" json:"actor_port_priority,omitempty"`
	PartnerPortId           uint32   `protobuf:"varint,12,opt,name=partner_port_id,json=partnerPortId,proto3" json:"partner_port_id,omitempty"`
	PartnerPortPriority     uint32   `protobuf:"varint,13,opt,name=partner_port_priority,json=partnerPortPriority,proto3" json:"partner_port_priority,omitempty"`
	ActorPortState          uint32   `protobuf:"varint,14,opt,name=actor_port_state,json=actorPortState,proto3" json:"actor_port_state,omitempty"`
	PartnerPortState        uint32   `protobuf:"varint,15,opt,name=partner_port_state,json=partnerPortState,proto3" json:"partner_port_state,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *BmLacpLinkData) Reset()         { *m = BmLacpLinkData{} }
func (m *BmLacpLinkData) String() string { return proto.CompactTextString(m) }
func (*BmLacpLinkData) ProtoMessage()    {}
func (*BmLacpLinkData) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{2}
}

func (m *BmLacpLinkData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmLacpLinkData.Unmarshal(m, b)
}
func (m *BmLacpLinkData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmLacpLinkData.Marshal(b, m, deterministic)
}
func (m *BmLacpLinkData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmLacpLinkData.Merge(m, src)
}
func (m *BmLacpLinkData) XXX_Size() int {
	return xxx_messageInfo_BmLacpLinkData.Size(m)
}
func (m *BmLacpLinkData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmLacpLinkData.DiscardUnknown(m)
}

var xxx_messageInfo_BmLacpLinkData proto.InternalMessageInfo

func (m *BmLacpLinkData) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

func (m *BmLacpLinkData) GetActorSystemPriority() uint32 {
	if m != nil {
		return m.ActorSystemPriority
	}
	return 0
}

func (m *BmLacpLinkData) GetActorSystemMacAddress() string {
	if m != nil {
		return m.ActorSystemMacAddress
	}
	return ""
}

func (m *BmLacpLinkData) GetActorOperationalKey() uint32 {
	if m != nil {
		return m.ActorOperationalKey
	}
	return 0
}

func (m *BmLacpLinkData) GetPartnerSystemPriority() uint32 {
	if m != nil {
		return m.PartnerSystemPriority
	}
	return 0
}

func (m *BmLacpLinkData) GetPartnerSystemMacAddress() string {
	if m != nil {
		return m.PartnerSystemMacAddress
	}
	return ""
}

func (m *BmLacpLinkData) GetPartnerOperationalKey() uint32 {
	if m != nil {
		return m.PartnerOperationalKey
	}
	return 0
}

func (m *BmLacpLinkData) GetSelectedAggregatorId() uint32 {
	if m != nil {
		return m.SelectedAggregatorId
	}
	return 0
}

func (m *BmLacpLinkData) GetAttachedAggregatorId() uint32 {
	if m != nil {
		return m.AttachedAggregatorId
	}
	return 0
}

func (m *BmLacpLinkData) GetActorPortId() uint32 {
	if m != nil {
		return m.ActorPortId
	}
	return 0
}

func (m *BmLacpLinkData) GetActorPortPriority() uint32 {
	if m != nil {
		return m.ActorPortPriority
	}
	return 0
}

func (m *BmLacpLinkData) GetPartnerPortId() uint32 {
	if m != nil {
		return m.PartnerPortId
	}
	return 0
}

func (m *BmLacpLinkData) GetPartnerPortPriority() uint32 {
	if m != nil {
		return m.PartnerPortPriority
	}
	return 0
}

func (m *BmLacpLinkData) GetActorPortState() uint32 {
	if m != nil {
		return m.ActorPortState
	}
	return 0
}

func (m *BmLacpLinkData) GetPartnerPortState() uint32 {
	if m != nil {
		return m.PartnerPortState
	}
	return 0
}

type BmMbrStateReasonDataType struct {
	ReasonType           string   `protobuf:"bytes,1,opt,name=reason_type,json=reasonType,proto3" json:"reason_type,omitempty"`
	Severity             string   `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmMbrStateReasonDataType) Reset()         { *m = BmMbrStateReasonDataType{} }
func (m *BmMbrStateReasonDataType) String() string { return proto.CompactTextString(m) }
func (*BmMbrStateReasonDataType) ProtoMessage()    {}
func (*BmMbrStateReasonDataType) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{3}
}

func (m *BmMbrStateReasonDataType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMbrStateReasonDataType.Unmarshal(m, b)
}
func (m *BmMbrStateReasonDataType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMbrStateReasonDataType.Marshal(b, m, deterministic)
}
func (m *BmMbrStateReasonDataType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMbrStateReasonDataType.Merge(m, src)
}
func (m *BmMbrStateReasonDataType) XXX_Size() int {
	return xxx_messageInfo_BmMbrStateReasonDataType.Size(m)
}
func (m *BmMbrStateReasonDataType) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMbrStateReasonDataType.DiscardUnknown(m)
}

var xxx_messageInfo_BmMbrStateReasonDataType proto.InternalMessageInfo

func (m *BmMbrStateReasonDataType) GetReasonType() string {
	if m != nil {
		return m.ReasonType
	}
	return ""
}

func (m *BmMbrStateReasonDataType) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

type BmMuxData struct {
	MuxState                 string                    `protobuf:"bytes,1,opt,name=mux_state,json=muxState,proto3" json:"mux_state,omitempty"`
	Error                    uint32                    `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
	MemberMuxStateReason     string                    `protobuf:"bytes,3,opt,name=member_mux_state_reason,json=memberMuxStateReason,proto3" json:"member_mux_state_reason,omitempty"`
	MemberMuxStateReasonData *BmMbrStateReasonDataType `protobuf:"bytes,4,opt,name=member_mux_state_reason_data,json=memberMuxStateReasonData,proto3" json:"member_mux_state_reason_data,omitempty"`
	MemberState              string                    `protobuf:"bytes,5,opt,name=member_state,json=memberState,proto3" json:"member_state,omitempty"`
	MuxStateReason           string                    `protobuf:"bytes,6,opt,name=mux_state_reason,json=muxStateReason,proto3" json:"mux_state_reason,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *BmMuxData) Reset()         { *m = BmMuxData{} }
func (m *BmMuxData) String() string { return proto.CompactTextString(m) }
func (*BmMuxData) ProtoMessage()    {}
func (*BmMuxData) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{4}
}

func (m *BmMuxData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMuxData.Unmarshal(m, b)
}
func (m *BmMuxData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMuxData.Marshal(b, m, deterministic)
}
func (m *BmMuxData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMuxData.Merge(m, src)
}
func (m *BmMuxData) XXX_Size() int {
	return xxx_messageInfo_BmMuxData.Size(m)
}
func (m *BmMuxData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMuxData.DiscardUnknown(m)
}

var xxx_messageInfo_BmMuxData proto.InternalMessageInfo

func (m *BmMuxData) GetMuxState() string {
	if m != nil {
		return m.MuxState
	}
	return ""
}

func (m *BmMuxData) GetError() uint32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *BmMuxData) GetMemberMuxStateReason() string {
	if m != nil {
		return m.MemberMuxStateReason
	}
	return ""
}

func (m *BmMuxData) GetMemberMuxStateReasonData() *BmMbrStateReasonDataType {
	if m != nil {
		return m.MemberMuxStateReasonData
	}
	return nil
}

func (m *BmMuxData) GetMemberState() string {
	if m != nil {
		return m.MemberState
	}
	return ""
}

func (m *BmMuxData) GetMuxStateReason() string {
	if m != nil {
		return m.MuxStateReason
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
	return fileDescriptor_af03d980950c6672, []int{5}
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

type BmMemberData struct {
	InterfaceName        string          `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	PortPriority         uint32          `protobuf:"varint,51,opt,name=port_priority,json=portPriority,proto3" json:"port_priority,omitempty"`
	PortNumber           uint32          `protobuf:"varint,52,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	UnderlyingLinkId     uint32          `protobuf:"varint,53,opt,name=underlying_link_id,json=underlyingLinkId,proto3" json:"underlying_link_id,omitempty"`
	LinkOrderNumber      uint32          `protobuf:"varint,54,opt,name=link_order_number,json=linkOrderNumber,proto3" json:"link_order_number,omitempty"`
	IccpNode             uint32          `protobuf:"varint,55,opt,name=iccp_node,json=iccpNode,proto3" json:"iccp_node,omitempty"`
	Bandwidth            uint32          `protobuf:"varint,56,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	LacpEnabled          string          `protobuf:"bytes,57,opt,name=lacp_enabled,json=lacpEnabled,proto3" json:"lacp_enabled,omitempty"`
	Counters             *BmLacpCounters `protobuf:"bytes,58,opt,name=counters,proto3" json:"counters,omitempty"`
	LinkData             *BmLacpLinkData `protobuf:"bytes,59,opt,name=link_data,json=linkData,proto3" json:"link_data,omitempty"`
	MemberMuxData        *BmMuxData      `protobuf:"bytes,60,opt,name=member_mux_data,json=memberMuxData,proto3" json:"member_mux_data,omitempty"`
	MemberType           string          `protobuf:"bytes,61,opt,name=member_type,json=memberType,proto3" json:"member_type,omitempty"`
	MemberName           string          `protobuf:"bytes,62,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	MacAddress           *BmMacAddrSt    `protobuf:"bytes,63,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BmMemberData) Reset()         { *m = BmMemberData{} }
func (m *BmMemberData) String() string { return proto.CompactTextString(m) }
func (*BmMemberData) ProtoMessage()    {}
func (*BmMemberData) Descriptor() ([]byte, []int) {
	return fileDescriptor_af03d980950c6672, []int{6}
}

func (m *BmMemberData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMemberData.Unmarshal(m, b)
}
func (m *BmMemberData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMemberData.Marshal(b, m, deterministic)
}
func (m *BmMemberData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMemberData.Merge(m, src)
}
func (m *BmMemberData) XXX_Size() int {
	return xxx_messageInfo_BmMemberData.Size(m)
}
func (m *BmMemberData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMemberData.DiscardUnknown(m)
}

var xxx_messageInfo_BmMemberData proto.InternalMessageInfo

func (m *BmMemberData) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *BmMemberData) GetPortPriority() uint32 {
	if m != nil {
		return m.PortPriority
	}
	return 0
}

func (m *BmMemberData) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

func (m *BmMemberData) GetUnderlyingLinkId() uint32 {
	if m != nil {
		return m.UnderlyingLinkId
	}
	return 0
}

func (m *BmMemberData) GetLinkOrderNumber() uint32 {
	if m != nil {
		return m.LinkOrderNumber
	}
	return 0
}

func (m *BmMemberData) GetIccpNode() uint32 {
	if m != nil {
		return m.IccpNode
	}
	return 0
}

func (m *BmMemberData) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *BmMemberData) GetLacpEnabled() string {
	if m != nil {
		return m.LacpEnabled
	}
	return ""
}

func (m *BmMemberData) GetCounters() *BmLacpCounters {
	if m != nil {
		return m.Counters
	}
	return nil
}

func (m *BmMemberData) GetLinkData() *BmLacpLinkData {
	if m != nil {
		return m.LinkData
	}
	return nil
}

func (m *BmMemberData) GetMemberMuxData() *BmMuxData {
	if m != nil {
		return m.MemberMuxData
	}
	return nil
}

func (m *BmMemberData) GetMemberType() string {
	if m != nil {
		return m.MemberType
	}
	return ""
}

func (m *BmMemberData) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

func (m *BmMemberData) GetMacAddress() *BmMacAddrSt {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*BmMemberData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_member_data_KEYS")
	proto.RegisterType((*BmLacpCounters)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_lacp_counters")
	proto.RegisterType((*BmLacpLinkData)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_lacp_link_data")
	proto.RegisterType((*BmMbrStateReasonDataType)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_mbr_state_reason_data_type")
	proto.RegisterType((*BmMuxData)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_mux_data")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_mac_addr_st")
	proto.RegisterType((*BmMemberData)(nil), "cisco_ios_xr_bundlemgr_oper.lacp_bundles.bundles.bundle.members.member.bm_member_data")
}

func init() { proto.RegisterFile("bm_member_data.proto", fileDescriptor_af03d980950c6672) }

var fileDescriptor_af03d980950c6672 = []byte{
	// 1083 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0x46, 0xf2, 0x27, 0xb6, 0x34, 0xb2, 0x6c, 0x79, 0x2d, 0xc7, 0x0b, 0xdb, 0x7f, 0xdd, 0xaa,
	0x68, 0xe1, 0x06, 0x85, 0x50, 0x38, 0xc7, 0x36, 0x3d, 0x05, 0x69, 0x8a, 0x1a, 0x89, 0x1d, 0x63,
	0xed, 0x16, 0x0d, 0x50, 0x80, 0xa0, 0x96, 0x13, 0x79, 0xa1, 0x3d, 0x81, 0xa4, 0x1c, 0x0b, 0xbd,
	0xe9, 0xcb, 0xf4, 0xbe, 0xef, 0xd3, 0x57, 0xe8, 0x43, 0x14, 0xe4, 0x70, 0xa9, 0x5d, 0x39, 0xe9,
	0x95, 0xaf, 0x24, 0x7e, 0xf3, 0xcd, 0x81, 0xe4, 0x37, 0xc3, 0x85, 0xfe, 0x28, 0x63, 0x19, 0x66,
	0x23, 0x94, 0x4c, 0x70, 0xcd, 0x87, 0xa5, 0x2c, 0x74, 0x11, 0xfc, 0x18, 0x27, 0x2a, 0x2e, 0x58,
	0x52, 0x28, 0x76, 0x29, 0xd9, 0x68, 0x9a, 0x8b, 0x14, 0xb3, 0xb1, 0x64, 0x45, 0x89, 0x72, 0x98,
	0xf2, 0xb8, 0x74, 0x98, 0x1a, 0x36, 0x7f, 0x87, 0x14, 0x4a, 0xb9, 0xdf, 0xc1, 0x04, 0x36, 0x9a,
	0xf1, 0xd9, 0x8b, 0xe7, 0xaf, 0x4f, 0x83, 0xcf, 0xa0, 0x47, 0x7c, 0x96, 0xe4, 0x1a, 0xe5, 0x1b,
	0x1e, 0x63, 0x78, 0xe3, 0xc3, 0x1b, 0xfb, 0xed, 0x68, 0x8d, 0xf0, 0xc3, 0x0a, 0x36, 0x54, 0xe7,
	0x3e, 0xa7, 0xde, 0x24, 0x2a, 0xe1, 0x9e, 0x3a, 0xf8, 0xeb, 0x16, 0xf4, 0x46, 0x19, 0xb3, 0xe5,
	0xc5, 0xc5, 0xd4, 0xe0, 0x2a, 0xb8, 0x0b, 0xeb, 0x06, 0x10, 0x6c, 0xaa, 0x98, 0xc4, 0x18, 0x93,
	0x0b, 0x14, 0x36, 0x57, 0x37, 0x5a, 0xb3, 0x86, 0x9f, 0x55, 0xe4, 0xe0, 0xe0, 0x0b, 0xe8, 0x7b,
	0xae, 0x96, 0x3c, 0x57, 0x59, 0xa2, 0x35, 0x0a, 0x9b, 0xaf, 0x1b, 0x05, 0x8e, 0x7e, 0x36, 0xb7,
	0x04, 0x0f, 0x61, 0x2b, 0xe3, 0x72, 0x82, 0x92, 0x95, 0x3c, 0x9e, 0xa0, 0xae, 0xe5, 0xf8, 0x9f,
	0x75, 0xda, 0x24, 0xf3, 0x09, 0x59, 0x7d, 0xa6, 0xef, 0x61, 0xd7, 0xf9, 0x49, 0x54, 0x65, 0x91,
	0x2b, 0x6c, 0x66, 0xbc, 0x65, 0x9d, 0xb7, 0x89, 0x13, 0x55, 0x94, 0x7a, 0xe6, 0xc7, 0x10, 0x26,
	0x69, 0x8a, 0x63, 0x9e, 0x5e, 0x4d, 0x7d, 0xdb, 0x7a, 0xdf, 0x71, 0xf6, 0xc5, 0xdc, 0x8f, 0x20,
	0xc4, 0xcb, 0x18, 0x95, 0x62, 0x57, 0x0f, 0x66, 0x89, 0x8a, 0x26, 0xfb, 0xcb, 0x85, 0xe3, 0x79,
	0x06, 0x1f, 0x38, 0xc7, 0xf7, 0xed, 0x79, 0xd9, 0xba, 0xef, 0x10, 0xeb, 0xe8, 0x9d, 0x3b, 0xdf,
	0x85, 0xb6, 0xc0, 0x37, 0x7c, 0x9a, 0x9a, 0x6d, 0xb6, 0x2c, 0x7f, 0x0e, 0x04, 0x21, 0x2c, 0xe3,
	0x65, 0x99, 0x48, 0x14, 0x61, 0xdb, 0xda, 0xaa, 0x65, 0xb0, 0x0f, 0xbd, 0x94, 0x2b, 0xcd, 0xe2,
	0x14, 0xb9, 0x44, 0xc1, 0x14, 0xc6, 0x21, 0x58, 0xca, 0xaa, 0xc1, 0x9f, 0x11, 0x7c, 0x8a, 0x31,
	0xdd, 0x78, 0x8d, 0x99, 0x1b, 0x6a, 0xa7, 0xba, 0x71, 0x4f, 0x3d, 0x56, 0x18, 0x0f, 0xfe, 0x58,
	0x82, 0xf5, 0x4a, 0x32, 0x69, 0x92, 0x4f, 0xac, 0x46, 0x8d, 0xe6, 0xbc, 0xd8, 0xd8, 0x39, 0x37,
	0x82, 0xac, 0xe4, 0xe9, 0xf1, 0x9f, 0x2c, 0x1c, 0x1c, 0xc0, 0x26, 0x8f, 0x75, 0x21, 0x99, 0x9a,
	0x29, 0x8d, 0x19, 0x2b, 0x65, 0x52, 0xc8, 0x44, 0xcf, 0x9c, 0x66, 0x36, 0xac, 0xf1, 0xd4, 0xda,
	0x4e, 0x9c, 0xc9, 0x5c, 0x40, 0xc3, 0x27, 0xe3, 0x31, 0xe3, 0x42, 0x48, 0x54, 0xca, 0xaa, 0xa6,
	0x1d, 0x6d, 0xd6, 0xdc, 0x8e, 0x78, 0xfc, 0x94, 0x8c, 0xf3, 0x64, 0xa6, 0x0d, 0xb9, 0x4e, 0x8a,
	0x9c, 0xa7, 0x6c, 0x82, 0x33, 0x27, 0x17, 0x4a, 0xf6, 0x6a, 0x6e, 0x7b, 0x81, 0x33, 0xa3, 0xd0,
	0x92, 0x4b, 0x9d, 0xe3, 0xd5, 0x12, 0x49, 0x26, 0x9b, 0xce, 0xbc, 0x50, 0xe4, 0x13, 0xd8, 0x5e,
	0xf0, 0xab, 0x97, 0xb9, 0x64, 0xcb, 0xdc, 0x6a, 0xb8, 0xd6, 0x0a, 0xad, 0x25, 0x5d, 0x2c, 0x75,
	0xb9, 0x91, 0x74, 0xa1, 0xd8, 0xfb, 0x70, 0x47, 0x61, 0x8a, 0xb1, 0x46, 0xc1, 0xf8, 0x78, 0x2c,
	0x71, 0xcc, 0xcd, 0x76, 0x93, 0x4a, 0x29, 0xfd, 0xca, 0xfa, 0xd4, 0x1b, 0x0f, 0x85, 0xf1, 0xe2,
	0x5a, 0xf3, 0xf8, 0xfc, 0x8a, 0x17, 0x69, 0xa8, 0x5f, 0x59, 0x1b, 0x5e, 0x03, 0xe8, 0xd2, 0x61,
	0x96, 0x85, 0xd4, 0x86, 0x4c, 0x6a, 0xea, 0x58, 0xf0, 0xa4, 0x90, 0xfa, 0x50, 0x04, 0x43, 0xd8,
	0xa8, 0x71, 0xfc, 0xc1, 0x91, 0x98, 0xd6, 0x3d, 0xd3, 0x1f, 0xda, 0xa7, 0xb0, 0x56, 0xed, 0xbb,
	0x8a, 0xba, 0x62, 0xb9, 0x5d, 0x07, 0xbb, 0xb8, 0x07, 0xb0, 0xd9, 0xe0, 0xf9, 0xc8, 0x5d, 0xba,
	0xc8, 0x1a, 0xdb, 0xc7, 0xde, 0x87, 0x5e, 0xad, 0x16, 0xa5, 0xb9, 0xc6, 0x70, 0x95, 0x1a, 0xc0,
	0x17, 0x72, 0x6a, 0xd0, 0xe0, 0x73, 0x08, 0x1a, 0xd1, 0x89, 0xbb, 0x66, 0xb9, 0xbd, 0x5a, 0x68,
	0xcb, 0x1e, 0xfc, 0x06, 0xff, 0x37, 0x23, 0x7a, 0x24, 0x89, 0xc7, 0x24, 0x72, 0x55, 0xe4, 0x34,
	0xac, 0xf5, 0xac, 0xc4, 0x60, 0x0f, 0x3a, 0x0e, 0x33, 0x4b, 0xd7, 0x08, 0x40, 0xd0, 0x99, 0x21,
	0x6c, 0x43, 0x4b, 0xe1, 0x05, 0x7a, 0xd9, 0xb7, 0x23, 0xbf, 0x1e, 0xfc, 0x73, 0x13, 0x3a, 0x26,
	0xfc, 0xf4, 0x92, 0x5a, 0x6b, 0x07, 0xda, 0xe6, 0x3f, 0x95, 0x44, 0xa1, 0x5a, 0xd9, 0xf4, 0x92,
	0x0a, 0xef, 0xc3, 0x6d, 0x94, 0xb2, 0x90, 0xae, 0x79, 0x68, 0x11, 0x3c, 0x80, 0x2d, 0xf7, 0x02,
	0x78, 0x4f, 0x57, 0xa4, 0xeb, 0x96, 0x3e, 0x99, 0x8f, 0x5c, 0x98, 0xc8, 0xda, 0x82, 0x3f, 0x6f,
	0xc0, 0xee, 0x7b, 0xfc, 0x6c, 0x29, 0xb6, 0x69, 0x3a, 0x07, 0x38, 0xbc, 0x9e, 0xa7, 0x6e, 0xf8,
	0x9f, 0x87, 0x18, 0x85, 0xef, 0xaa, 0xf1, 0x07, 0x73, 0x22, 0x1f, 0xc1, 0x8a, 0x2b, 0x93, 0x0e,
	0xe5, 0xb6, 0xdd, 0x53, 0x87, 0x30, 0x3a, 0x97, 0x7d, 0xe8, 0x5d, 0xd9, 0x3a, 0x75, 0xe0, 0x6a,
	0xd6, 0x08, 0x38, 0xb8, 0x0b, 0xab, 0xa3, 0x79, 0xa7, 0x32, 0xa5, 0xcd, 0x44, 0xad, 0x9a, 0x96,
	0x8e, 0xbb, 0x5a, 0x0e, 0xfe, 0x5e, 0x22, 0xf2, 0xfc, 0x71, 0x0e, 0x3e, 0x81, 0xd5, 0xf9, 0xe0,
	0xcb, 0x79, 0x86, 0xe1, 0x81, 0xf5, 0xe9, 0x7a, 0xf4, 0x98, 0x67, 0x18, 0x7c, 0x0c, 0xdd, 0xa6,
	0x6c, 0xef, 0xd9, 0xfb, 0x5a, 0x29, 0xeb, 0x7a, 0xdd, 0x83, 0x8e, 0x25, 0xe5, 0x53, 0x13, 0x3f,
	0xbc, 0x6f, 0x29, 0x60, 0xa0, 0x63, 0x8b, 0x18, 0x99, 0x4e, 0x73, 0x81, 0x32, 0x9d, 0x25, 0xf9,
	0x98, 0xa6, 0x6f, 0x22, 0xc2, 0x07, 0x24, 0xd3, 0xb9, 0xe5, 0x65, 0x92, 0x4f, 0x0e, 0x85, 0x9d,
	0xea, 0x86, 0x52, 0x48, 0x81, 0xb2, 0x0a, 0xfa, 0xd0, 0x4d, 0xf5, 0x24, 0x9f, 0xbc, 0x32, 0xb8,
	0x8b, 0xbc, 0x03, 0xed, 0x24, 0x8e, 0x4b, 0x96, 0x17, 0x02, 0xc3, 0x47, 0x96, 0xd3, 0x32, 0xc0,
	0x71, 0x21, 0xd0, 0x3c, 0x40, 0x23, 0x9e, 0x8b, 0xb7, 0x89, 0xd0, 0xe7, 0xe1, 0x63, 0x7a, 0x80,
	0x3c, 0x60, 0x6e, 0xc3, 0xde, 0x39, 0xe6, 0x7c, 0x94, 0xa2, 0x08, 0xbf, 0xa4, 0xdb, 0x30, 0xd8,
	0x73, 0x82, 0x02, 0x0d, 0xad, 0xea, 0xeb, 0x22, 0xfc, 0xca, 0x6a, 0xe8, 0xd7, 0x6b, 0xd4, 0x50,
	0xe3, 0xeb, 0x25, 0xf2, 0x99, 0x82, 0x0b, 0x68, 0xfb, 0x07, 0x2a, 0x7c, 0x62, 0xd3, 0xbe, 0xbe,
	0xee, 0xb4, 0x3e, 0x41, 0xd4, 0x32, 0x7f, 0xad, 0x3c, 0x7f, 0x87, 0xb5, 0x5a, 0x17, 0xd9, 0xec,
	0x5f, 0xdb, 0xec, 0xa7, 0xd7, 0xd9, 0x38, 0x2e, 0x74, 0xd4, 0xf5, 0x6d, 0x62, 0x93, 0xef, 0x81,
	0xeb, 0x03, 0x1a, 0x3d, 0xdf, 0xd0, 0xe8, 0x21, 0xe8, 0xcc, 0xcd, 0x26, 0x47, 0xb0, 0x6a, 0xfd,
	0xb6, 0x4e, 0xb0, 0x52, 0x7d, 0x0b, 0x9d, 0xfa, 0xbb, 0xf5, 0x9d, 0x2d, 0xfd, 0x97, 0xeb, 0x2c,
	0x7d, 0xde, 0x6b, 0x11, 0x64, 0xfe, 0x09, 0x1c, 0x2d, 0xd9, 0x0f, 0xe9, 0x7b, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0xff, 0x51, 0x0d, 0xc5, 0x60, 0x0b, 0x00, 0x00,
}
