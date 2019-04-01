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
// source: lacp_mbr_data.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_lacp_lacp_bundles_lacp_bundle_lacp_bundle_children_members_lacp_bundle_children_member

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

type LacpMbrData_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	MemberInterface      string   `protobuf:"bytes,2,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpMbrData_KEYS) Reset()         { *m = LacpMbrData_KEYS{} }
func (m *LacpMbrData_KEYS) String() string { return proto.CompactTextString(m) }
func (*LacpMbrData_KEYS) ProtoMessage()    {}
func (*LacpMbrData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{0}
}

func (m *LacpMbrData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpMbrData_KEYS.Unmarshal(m, b)
}
func (m *LacpMbrData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpMbrData_KEYS.Marshal(b, m, deterministic)
}
func (m *LacpMbrData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpMbrData_KEYS.Merge(m, src)
}
func (m *LacpMbrData_KEYS) XXX_Size() int {
	return xxx_messageInfo_LacpMbrData_KEYS.Size(m)
}
func (m *LacpMbrData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpMbrData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LacpMbrData_KEYS proto.InternalMessageInfo

func (m *LacpMbrData_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

func (m *LacpMbrData_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
	}
	return ""
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
	return fileDescriptor_8af27d3c02c3c096, []int{1}
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
	return fileDescriptor_8af27d3c02c3c096, []int{2}
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

type BmLinkIdSt struct {
	LinkPriority         uint32   `protobuf:"varint,1,opt,name=link_priority,json=linkPriority,proto3" json:"link_priority,omitempty"`
	LinkNumber           uint32   `protobuf:"varint,2,opt,name=link_number,json=linkNumber,proto3" json:"link_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmLinkIdSt) Reset()         { *m = BmLinkIdSt{} }
func (m *BmLinkIdSt) String() string { return proto.CompactTextString(m) }
func (*BmLinkIdSt) ProtoMessage()    {}
func (*BmLinkIdSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{3}
}

func (m *BmLinkIdSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmLinkIdSt.Unmarshal(m, b)
}
func (m *BmLinkIdSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmLinkIdSt.Marshal(b, m, deterministic)
}
func (m *BmLinkIdSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmLinkIdSt.Merge(m, src)
}
func (m *BmLinkIdSt) XXX_Size() int {
	return xxx_messageInfo_BmLinkIdSt.Size(m)
}
func (m *BmLinkIdSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmLinkIdSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmLinkIdSt proto.InternalMessageInfo

func (m *BmLinkIdSt) GetLinkPriority() uint32 {
	if m != nil {
		return m.LinkPriority
	}
	return 0
}

func (m *BmLinkIdSt) GetLinkNumber() uint32 {
	if m != nil {
		return m.LinkNumber
	}
	return 0
}

type BmLacpPortInfo_ struct {
	System               *BmSystemIdSt `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Key                  uint32        `protobuf:"varint,2,opt,name=key,proto3" json:"key,omitempty"`
	Port                 *BmLinkIdSt   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	State                uint32        `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BmLacpPortInfo_) Reset()         { *m = BmLacpPortInfo_{} }
func (m *BmLacpPortInfo_) String() string { return proto.CompactTextString(m) }
func (*BmLacpPortInfo_) ProtoMessage()    {}
func (*BmLacpPortInfo_) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{4}
}

func (m *BmLacpPortInfo_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmLacpPortInfo_.Unmarshal(m, b)
}
func (m *BmLacpPortInfo_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmLacpPortInfo_.Marshal(b, m, deterministic)
}
func (m *BmLacpPortInfo_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmLacpPortInfo_.Merge(m, src)
}
func (m *BmLacpPortInfo_) XXX_Size() int {
	return xxx_messageInfo_BmLacpPortInfo_.Size(m)
}
func (m *BmLacpPortInfo_) XXX_DiscardUnknown() {
	xxx_messageInfo_BmLacpPortInfo_.DiscardUnknown(m)
}

var xxx_messageInfo_BmLacpPortInfo_ proto.InternalMessageInfo

func (m *BmLacpPortInfo_) GetSystem() *BmSystemIdSt {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *BmLacpPortInfo_) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *BmLacpPortInfo_) GetPort() *BmLinkIdSt {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *BmLacpPortInfo_) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type LacpLinkDeviceInfoSt struct {
	PortInfo             *BmLacpPortInfo_ `protobuf:"bytes,1,opt,name=port_info,json=portInfo,proto3" json:"port_info,omitempty"`
	TxPeriod             uint32           `protobuf:"varint,2,opt,name=tx_period,json=txPeriod,proto3" json:"tx_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LacpLinkDeviceInfoSt) Reset()         { *m = LacpLinkDeviceInfoSt{} }
func (m *LacpLinkDeviceInfoSt) String() string { return proto.CompactTextString(m) }
func (*LacpLinkDeviceInfoSt) ProtoMessage()    {}
func (*LacpLinkDeviceInfoSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{5}
}

func (m *LacpLinkDeviceInfoSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLinkDeviceInfoSt.Unmarshal(m, b)
}
func (m *LacpLinkDeviceInfoSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLinkDeviceInfoSt.Marshal(b, m, deterministic)
}
func (m *LacpLinkDeviceInfoSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLinkDeviceInfoSt.Merge(m, src)
}
func (m *LacpLinkDeviceInfoSt) XXX_Size() int {
	return xxx_messageInfo_LacpLinkDeviceInfoSt.Size(m)
}
func (m *LacpLinkDeviceInfoSt) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLinkDeviceInfoSt.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLinkDeviceInfoSt proto.InternalMessageInfo

func (m *LacpLinkDeviceInfoSt) GetPortInfo() *BmLacpPortInfo_ {
	if m != nil {
		return m.PortInfo
	}
	return nil
}

func (m *LacpLinkDeviceInfoSt) GetTxPeriod() uint32 {
	if m != nil {
		return m.TxPeriod
	}
	return 0
}

type LacpLinkAdditionalInfoLocal struct {
	InterfaceHandle      string   `protobuf:"bytes,1,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpLinkAdditionalInfoLocal) Reset()         { *m = LacpLinkAdditionalInfoLocal{} }
func (m *LacpLinkAdditionalInfoLocal) String() string { return proto.CompactTextString(m) }
func (*LacpLinkAdditionalInfoLocal) ProtoMessage()    {}
func (*LacpLinkAdditionalInfoLocal) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{6}
}

func (m *LacpLinkAdditionalInfoLocal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLinkAdditionalInfoLocal.Unmarshal(m, b)
}
func (m *LacpLinkAdditionalInfoLocal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLinkAdditionalInfoLocal.Marshal(b, m, deterministic)
}
func (m *LacpLinkAdditionalInfoLocal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLinkAdditionalInfoLocal.Merge(m, src)
}
func (m *LacpLinkAdditionalInfoLocal) XXX_Size() int {
	return xxx_messageInfo_LacpLinkAdditionalInfoLocal.Size(m)
}
func (m *LacpLinkAdditionalInfoLocal) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLinkAdditionalInfoLocal.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLinkAdditionalInfoLocal proto.InternalMessageInfo

func (m *LacpLinkAdditionalInfoLocal) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

type LacpLinkAdditionalInfoForeign struct {
	PeerAddress          string   `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	MemberName           string   `protobuf:"bytes,2,opt,name=member_name,json=memberName,proto3" json:"member_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LacpLinkAdditionalInfoForeign) Reset()         { *m = LacpLinkAdditionalInfoForeign{} }
func (m *LacpLinkAdditionalInfoForeign) String() string { return proto.CompactTextString(m) }
func (*LacpLinkAdditionalInfoForeign) ProtoMessage()    {}
func (*LacpLinkAdditionalInfoForeign) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{7}
}

func (m *LacpLinkAdditionalInfoForeign) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLinkAdditionalInfoForeign.Unmarshal(m, b)
}
func (m *LacpLinkAdditionalInfoForeign) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLinkAdditionalInfoForeign.Marshal(b, m, deterministic)
}
func (m *LacpLinkAdditionalInfoForeign) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLinkAdditionalInfoForeign.Merge(m, src)
}
func (m *LacpLinkAdditionalInfoForeign) XXX_Size() int {
	return xxx_messageInfo_LacpLinkAdditionalInfoForeign.Size(m)
}
func (m *LacpLinkAdditionalInfoForeign) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLinkAdditionalInfoForeign.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLinkAdditionalInfoForeign proto.InternalMessageInfo

func (m *LacpLinkAdditionalInfoForeign) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *LacpLinkAdditionalInfoForeign) GetMemberName() string {
	if m != nil {
		return m.MemberName
	}
	return ""
}

type LacpLinkAdditionalInfo struct {
	MbrType              string                         `protobuf:"bytes,1,opt,name=mbr_type,json=mbrType,proto3" json:"mbr_type,omitempty"`
	Local                *LacpLinkAdditionalInfoLocal   `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
	Foreign              *LacpLinkAdditionalInfoForeign `protobuf:"bytes,3,opt,name=foreign,proto3" json:"foreign,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *LacpLinkAdditionalInfo) Reset()         { *m = LacpLinkAdditionalInfo{} }
func (m *LacpLinkAdditionalInfo) String() string { return proto.CompactTextString(m) }
func (*LacpLinkAdditionalInfo) ProtoMessage()    {}
func (*LacpLinkAdditionalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{8}
}

func (m *LacpLinkAdditionalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpLinkAdditionalInfo.Unmarshal(m, b)
}
func (m *LacpLinkAdditionalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpLinkAdditionalInfo.Marshal(b, m, deterministic)
}
func (m *LacpLinkAdditionalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpLinkAdditionalInfo.Merge(m, src)
}
func (m *LacpLinkAdditionalInfo) XXX_Size() int {
	return xxx_messageInfo_LacpLinkAdditionalInfo.Size(m)
}
func (m *LacpLinkAdditionalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpLinkAdditionalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LacpLinkAdditionalInfo proto.InternalMessageInfo

func (m *LacpLinkAdditionalInfo) GetMbrType() string {
	if m != nil {
		return m.MbrType
	}
	return ""
}

func (m *LacpLinkAdditionalInfo) GetLocal() *LacpLinkAdditionalInfoLocal {
	if m != nil {
		return m.Local
	}
	return nil
}

func (m *LacpLinkAdditionalInfo) GetForeign() *LacpLinkAdditionalInfoForeign {
	if m != nil {
		return m.Foreign
	}
	return nil
}

type LacpMbrData struct {
	SelectedAggregatorId uint32                  `protobuf:"varint,50,opt,name=selected_aggregator_id,json=selectedAggregatorId,proto3" json:"selected_aggregator_id,omitempty"`
	AttachedAggregatorId uint32                  `protobuf:"varint,51,opt,name=attached_aggregator_id,json=attachedAggregatorId,proto3" json:"attached_aggregator_id,omitempty"`
	ActorInfo            *LacpLinkDeviceInfoSt   `protobuf:"bytes,52,opt,name=actor_info,json=actorInfo,proto3" json:"actor_info,omitempty"`
	PartnerInfo          *LacpLinkDeviceInfoSt   `protobuf:"bytes,53,opt,name=partner_info,json=partnerInfo,proto3" json:"partner_info,omitempty"`
	SelectionState       string                  `protobuf:"bytes,54,opt,name=selection_state,json=selectionState,proto3" json:"selection_state,omitempty"`
	PeriodState          string                  `protobuf:"bytes,55,opt,name=period_state,json=periodState,proto3" json:"period_state,omitempty"`
	ReceiveMachineState  string                  `protobuf:"bytes,56,opt,name=receive_machine_state,json=receiveMachineState,proto3" json:"receive_machine_state,omitempty"`
	MuxState             string                  `protobuf:"bytes,57,opt,name=mux_state,json=muxState,proto3" json:"mux_state,omitempty"`
	ActorChurnState      string                  `protobuf:"bytes,58,opt,name=actor_churn_state,json=actorChurnState,proto3" json:"actor_churn_state,omitempty"`
	PartnerChurnState    string                  `protobuf:"bytes,59,opt,name=partner_churn_state,json=partnerChurnState,proto3" json:"partner_churn_state,omitempty"`
	IccpGroupId          uint32                  `protobuf:"varint,60,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	AdditionalInfo       *LacpLinkAdditionalInfo `protobuf:"bytes,61,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LacpMbrData) Reset()         { *m = LacpMbrData{} }
func (m *LacpMbrData) String() string { return proto.CompactTextString(m) }
func (*LacpMbrData) ProtoMessage()    {}
func (*LacpMbrData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8af27d3c02c3c096, []int{9}
}

func (m *LacpMbrData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LacpMbrData.Unmarshal(m, b)
}
func (m *LacpMbrData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LacpMbrData.Marshal(b, m, deterministic)
}
func (m *LacpMbrData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LacpMbrData.Merge(m, src)
}
func (m *LacpMbrData) XXX_Size() int {
	return xxx_messageInfo_LacpMbrData.Size(m)
}
func (m *LacpMbrData) XXX_DiscardUnknown() {
	xxx_messageInfo_LacpMbrData.DiscardUnknown(m)
}

var xxx_messageInfo_LacpMbrData proto.InternalMessageInfo

func (m *LacpMbrData) GetSelectedAggregatorId() uint32 {
	if m != nil {
		return m.SelectedAggregatorId
	}
	return 0
}

func (m *LacpMbrData) GetAttachedAggregatorId() uint32 {
	if m != nil {
		return m.AttachedAggregatorId
	}
	return 0
}

func (m *LacpMbrData) GetActorInfo() *LacpLinkDeviceInfoSt {
	if m != nil {
		return m.ActorInfo
	}
	return nil
}

func (m *LacpMbrData) GetPartnerInfo() *LacpLinkDeviceInfoSt {
	if m != nil {
		return m.PartnerInfo
	}
	return nil
}

func (m *LacpMbrData) GetSelectionState() string {
	if m != nil {
		return m.SelectionState
	}
	return ""
}

func (m *LacpMbrData) GetPeriodState() string {
	if m != nil {
		return m.PeriodState
	}
	return ""
}

func (m *LacpMbrData) GetReceiveMachineState() string {
	if m != nil {
		return m.ReceiveMachineState
	}
	return ""
}

func (m *LacpMbrData) GetMuxState() string {
	if m != nil {
		return m.MuxState
	}
	return ""
}

func (m *LacpMbrData) GetActorChurnState() string {
	if m != nil {
		return m.ActorChurnState
	}
	return ""
}

func (m *LacpMbrData) GetPartnerChurnState() string {
	if m != nil {
		return m.PartnerChurnState
	}
	return ""
}

func (m *LacpMbrData) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

func (m *LacpMbrData) GetAdditionalInfo() *LacpLinkAdditionalInfo {
	if m != nil {
		return m.AdditionalInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LacpMbrData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_mbr_data_KEYS")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.bm_system_id_st")
	proto.RegisterType((*BmLinkIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.bm_link_id_st")
	proto.RegisterType((*BmLacpPortInfo_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.bm_lacp_port_info_")
	proto.RegisterType((*LacpLinkDeviceInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_link_device_info_st")
	proto.RegisterType((*LacpLinkAdditionalInfoLocal)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_link_additional_info_local")
	proto.RegisterType((*LacpLinkAdditionalInfoForeign)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_link_additional_info_foreign")
	proto.RegisterType((*LacpLinkAdditionalInfo)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_link_additional_info")
	proto.RegisterType((*LacpMbrData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.lacp.lacp_bundles.lacp_bundle.lacp_bundle_children_members.lacp_bundle_children_member.lacp_mbr_data")
}

func init() { proto.RegisterFile("lacp_mbr_data.proto", fileDescriptor_8af27d3c02c3c096) }

var fileDescriptor_8af27d3c02c3c096 = []byte{
	// 839 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xbd, 0x8e, 0x1c, 0x45,
	0x10, 0x56, 0xb3, 0xb6, 0xef, 0xb6, 0xf6, 0xe6, 0xd6, 0xee, 0x33, 0x68, 0x2c, 0x07, 0xb6, 0x87,
	0x00, 0xe3, 0x60, 0x83, 0xb3, 0xf9, 0x87, 0xe0, 0x84, 0x10, 0x9c, 0xc0, 0xd6, 0x69, 0x0c, 0x01,
	0x51, 0xab, 0xb7, 0xa7, 0x6f, 0xb7, 0xf1, 0xcc, 0xf4, 0xa8, 0xa7, 0xd7, 0xda, 0x7d, 0x00, 0xde,
	0x00, 0x90, 0x20, 0x81, 0x08, 0x41, 0x44, 0x48, 0xc2, 0x4f, 0xe0, 0x8c, 0x77, 0x21, 0x40, 0xbc,
	0x00, 0xaa, 0xea, 0x9e, 0xbd, 0x5d, 0xa3, 0x83, 0x74, 0x9c, 0x9c, 0xba, 0xbf, 0xfa, 0xba, 0xe6,
	0xfb, 0xba, 0xba, 0x6a, 0x0f, 0x0e, 0x4a, 0xa9, 0x1a, 0x51, 0x4d, 0x9d, 0x28, 0xa4, 0x97, 0x93,
	0xc6, 0x59, 0x6f, 0xf9, 0xe7, 0x4c, 0x99, 0x56, 0x59, 0x61, 0x6c, 0x2b, 0x96, 0x4e, 0x4c, 0x17,
	0x75, 0x51, 0xea, 0x6a, 0xe6, 0x84, 0x6d, 0xb4, 0x9b, 0x84, 0xad, 0x30, 0xf5, 0xa9, 0x75, 0x95,
	0xf4, 0xc6, 0xd6, 0x13, 0x4c, 0x42, 0x7f, 0x22, 0xb7, 0xdd, 0xdc, 0x6c, 0xae, 0x85, 0x9a, 0x9b,
	0xb2, 0x70, 0xba, 0x16, 0x95, 0xae, 0xa6, 0xda, 0xb5, 0xff, 0x15, 0xcc, 0x3e, 0x03, 0xbe, 0x25,
	0x4f, 0x7c, 0xf8, 0xde, 0xa7, 0x0f, 0xf9, 0xcb, 0x70, 0x79, 0x2d, 0xc1, 0x6b, 0x77, 0x2a, 0x95,
	0x4e, 0xd9, 0x4d, 0x76, 0x7b, 0x98, 0x8f, 0x03, 0x7e, 0xdc, 0xc1, 0x48, 0x0d, 0xa9, 0x36, 0xa8,
	0xcf, 0x05, 0x6a, 0xc0, 0xd7, 0xd4, 0xec, 0x0e, 0xec, 0x6b, 0x3f, 0xd7, 0x4e, 0x54, 0x52, 0xc9,
	0xa2, 0x70, 0x82, 0xa7, 0xb0, 0x13, 0xd7, 0x29, 0xbb, 0x39, 0xb8, 0x9d, 0xe4, 0xdd, 0x36, 0xfb,
	0x9b, 0xc1, 0x78, 0x5a, 0x89, 0x76, 0xd5, 0x7a, 0x5d, 0x09, 0x53, 0x88, 0xd6, 0xf3, 0x1b, 0x30,
	0x8a, 0xfb, 0xc6, 0x19, 0x4b, 0x82, 0x92, 0x1c, 0x02, 0x74, 0xe2, 0x8c, 0xe5, 0xbf, 0x31, 0x18,
	0x47, 0x46, 0x25, 0x95, 0xa0, 0xbc, 0xa8, 0x65, 0x74, 0xf8, 0x25, 0x9b, 0xf4, 0xe2, 0xbe, 0x27,
	0xdb, 0x17, 0x90, 0x27, 0x41, 0xee, 0x7d, 0xa9, 0x8e, 0xd0, 0xf5, 0x27, 0x90, 0x4c, 0x2b, 0x51,
	0x9a, 0xfa, 0x51, 0xb4, 0xfc, 0x22, 0x24, 0xb4, 0x43, 0xc3, 0xce, 0xf8, 0x55, 0x34, 0xbd, 0x87,
	0xe0, 0x49, 0xc4, 0xf0, 0x5e, 0x88, 0x54, 0x2f, 0xf0, 0x13, 0xe4, 0x38, 0xc9, 0x01, 0xa1, 0x07,
	0x84, 0x64, 0xdf, 0x0c, 0x80, 0x63, 0x5e, 0xd4, 0xd5, 0x58, 0xe7, 0xc9, 0xa8, 0xe0, 0x3f, 0x31,
	0xb8, 0x14, 0xbe, 0x4f, 0x69, 0x47, 0x87, 0x5f, 0xf5, 0xe5, 0x96, 0x9e, 0xaa, 0x7c, 0x1e, 0x65,
	0xf2, 0xcb, 0x30, 0x78, 0xa4, 0x57, 0xd1, 0x21, 0x2e, 0xf9, 0x8f, 0x0c, 0x2e, 0xa0, 0xa5, 0x74,
	0x40, 0x0e, 0xbe, 0xe8, 0x91, 0x83, 0xb3, 0x32, 0xe6, 0x24, 0x91, 0x5f, 0x85, 0x8b, 0xad, 0x97,
	0x5e, 0xa7, 0x17, 0x48, 0x7f, 0xd8, 0x64, 0x7f, 0x31, 0x48, 0x29, 0x13, 0xf1, 0x0b, 0xfd, 0xd8,
	0xa8, 0xa0, 0x10, 0xeb, 0xff, 0x33, 0x83, 0xe1, 0xba, 0x62, 0xb1, 0x4a, 0x5f, 0xf7, 0xc9, 0xe3,
	0xf6, 0x93, 0xca, 0x77, 0x71, 0x7d, 0x5c, 0x9f, 0x5a, 0x7e, 0x1d, 0x86, 0x7e, 0x29, 0x1a, 0xed,
	0x8c, 0x2d, 0x62, 0xc1, 0x76, 0xfd, 0xf2, 0x84, 0xf6, 0xd9, 0x47, 0x70, 0xe3, 0xcc, 0xb2, 0x2c,
	0x0a, 0x83, 0x32, 0x65, 0x19, 0x92, 0x94, 0x56, 0xc9, 0x12, 0xe7, 0xca, 0x7a, 0xa0, 0x88, 0xb9,
	0x44, 0x01, 0xdd, 0x08, 0x5a, 0xe3, 0x1f, 0x10, 0x9c, 0xcd, 0xe0, 0xd6, 0xf9, 0xd9, 0x4e, 0xad,
	0xd3, 0x66, 0x56, 0xf3, 0x5b, 0xb0, 0xd7, 0x68, 0xed, 0x68, 0x28, 0xe8, 0xb6, 0x8d, 0xb9, 0x46,
	0x88, 0x1d, 0x05, 0x08, 0xfb, 0x28, 0x8e, 0xb2, 0x5a, 0x56, 0xdd, 0x14, 0x83, 0x00, 0x3d, 0x90,
	0x95, 0xce, 0x7e, 0x19, 0xc0, 0xb5, 0x73, 0xbf, 0xc4, 0xaf, 0xc1, 0x2e, 0x4e, 0x51, 0xbf, 0x6a,
	0x3a, 0xa5, 0x3b, 0xd5, 0xd4, 0x7d, 0xbc, 0x6a, 0x34, 0xff, 0x95, 0xc1, 0x45, 0xb2, 0x15, 0xc7,
	0xd1, 0xf7, 0x7d, 0x29, 0xe1, 0xff, 0x54, 0x21, 0x0f, 0xaa, 0xf9, 0x13, 0x06, 0x3b, 0xf1, 0x22,
	0x63, 0xa3, 0xfd, 0xd0, 0x7f, 0x07, 0x51, 0x70, 0xde, 0x29, 0xcf, 0xfe, 0xdc, 0x81, 0x64, 0xeb,
	0xc7, 0x8e, 0xdf, 0x83, 0x17, 0x5a, 0x5d, 0x6a, 0xe5, 0x75, 0x21, 0xe4, 0x6c, 0xe6, 0xf4, 0x4c,
	0x7a, 0xeb, 0x84, 0x29, 0xd2, 0x43, 0x7a, 0xb1, 0x57, 0xbb, 0xe8, 0xd1, 0x3a, 0x78, 0x5c, 0xe0,
	0x29, 0xe9, 0xbd, 0x54, 0xf3, 0x7f, 0x9d, 0xba, 0x1b, 0x4e, 0x75, 0xd1, 0xad, 0x53, 0xbf, 0x33,
	0x00, 0xa9, 0x88, 0x88, 0xbd, 0x7c, 0x8f, 0xae, 0xf1, 0xdb, 0xfe, 0x5d, 0xe3, 0xf6, 0x04, 0xca,
	0x87, 0xa4, 0x99, 0x5a, 0xfa, 0x09, 0x83, 0xbd, 0x46, 0x3a, 0x5f, 0xeb, 0xe8, 0xe1, 0x95, 0x67,
	0xc4, 0xc3, 0x28, 0xaa, 0x26, 0x17, 0x2f, 0xc1, 0x38, 0x54, 0xd5, 0xd8, 0x5a, 0x84, 0x79, 0xfc,
	0x2a, 0x75, 0xeb, 0xfe, 0x1a, 0x7e, 0x88, 0x68, 0x98, 0x18, 0x38, 0xae, 0x22, 0xeb, 0xb5, 0x6e,
	0x62, 0x20, 0x16, 0x28, 0x87, 0xf0, 0xbc, 0xd3, 0x4a, 0x9b, 0xc7, 0x1a, 0x7f, 0xd2, 0xe7, 0xa6,
	0xd6, 0x91, 0xfb, 0x3a, 0x71, 0x0f, 0x62, 0xf0, 0x7e, 0x88, 0x85, 0x33, 0xd7, 0x61, 0x58, 0x2d,
	0x96, 0x91, 0xf7, 0x06, 0xf1, 0x76, 0xab, 0xc5, 0x32, 0x04, 0xef, 0xc0, 0x95, 0xf0, 0x46, 0xd4,
	0x7c, 0xe1, 0x3a, 0x79, 0x6f, 0x86, 0xb1, 0x47, 0x81, 0x77, 0x11, 0x0f, 0xdc, 0x09, 0x1c, 0x74,
	0xd5, 0xd8, 0x64, 0xbf, 0x45, 0xec, 0x2b, 0x31, 0xb4, 0xc1, 0xcf, 0x20, 0x31, 0x4a, 0x35, 0x62,
	0xe6, 0xec, 0xa2, 0xc1, 0xd7, 0xfa, 0x36, 0xbd, 0xd6, 0x11, 0x82, 0xef, 0x23, 0x76, 0x5c, 0xf0,
	0x3f, 0x18, 0x8c, 0x9f, 0xea, 0xa3, 0xf4, 0x1d, 0xaa, 0xf2, 0x77, 0xbd, 0x6f, 0xf8, 0x7c, 0xff,
	0x0c, 0xc0, 0x4a, 0x4f, 0x2f, 0xd1, 0xbf, 0xda, 0x77, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x06,
	0xde, 0x03, 0x91, 0x81, 0x0b, 0x00, 0x00,
}
