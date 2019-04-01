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
// source: fsync_clock_info_v2.proto

package cisco_ios_xr_freqsync_oper_frequency_synchronization_nodes_node_clock_datas_clock_data

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

type FsyncClockInfoV2_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClockType            string   `protobuf:"bytes,2,opt,name=clock_type,json=clockType,proto3" json:"clock_type,omitempty"`
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncClockInfoV2_KEYS) Reset()         { *m = FsyncClockInfoV2_KEYS{} }
func (m *FsyncClockInfoV2_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncClockInfoV2_KEYS) ProtoMessage()    {}
func (*FsyncClockInfoV2_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{0}
}

func (m *FsyncClockInfoV2_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncClockInfoV2_KEYS.Unmarshal(m, b)
}
func (m *FsyncClockInfoV2_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncClockInfoV2_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncClockInfoV2_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncClockInfoV2_KEYS.Merge(m, src)
}
func (m *FsyncClockInfoV2_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncClockInfoV2_KEYS.Size(m)
}
func (m *FsyncClockInfoV2_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncClockInfoV2_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncClockInfoV2_KEYS proto.InternalMessageInfo

func (m *FsyncClockInfoV2_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *FsyncClockInfoV2_KEYS) GetClockType() string {
	if m != nil {
		return m.ClockType
	}
	return ""
}

func (m *FsyncClockInfoV2_KEYS) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FsyncBagClockId struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	ClockName            string   `protobuf:"bytes,3,opt,name=clock_name,json=clockName,proto3" json:"clock_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncBagClockId) Reset()         { *m = FsyncBagClockId{} }
func (m *FsyncBagClockId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagClockId) ProtoMessage()    {}
func (*FsyncBagClockId) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{1}
}

func (m *FsyncBagClockId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagClockId.Unmarshal(m, b)
}
func (m *FsyncBagClockId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagClockId.Marshal(b, m, deterministic)
}
func (m *FsyncBagClockId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagClockId.Merge(m, src)
}
func (m *FsyncBagClockId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagClockId.Size(m)
}
func (m *FsyncBagClockId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagClockId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagClockId proto.InternalMessageInfo

func (m *FsyncBagClockId) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *FsyncBagClockId) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FsyncBagClockId) GetClockName() string {
	if m != nil {
		return m.ClockName
	}
	return ""
}

type FsyncBagSourceId struct {
	SourceClass              string           `protobuf:"bytes,1,opt,name=source_class,json=sourceClass,proto3" json:"source_class,omitempty"`
	EthernetInterface        string           `protobuf:"bytes,2,opt,name=ethernet_interface,json=ethernetInterface,proto3" json:"ethernet_interface,omitempty"`
	SonetInterface           string           `protobuf:"bytes,3,opt,name=sonet_interface,json=sonetInterface,proto3" json:"sonet_interface,omitempty"`
	ClockId                  *FsyncBagClockId `protobuf:"bytes,4,opt,name=clock_id,json=clockId,proto3" json:"clock_id,omitempty"`
	InternalClockId          *FsyncBagClockId `protobuf:"bytes,5,opt,name=internal_clock_id,json=internalClockId,proto3" json:"internal_clock_id,omitempty"`
	PtpNode                  string           `protobuf:"bytes,6,opt,name=ptp_node,json=ptpNode,proto3" json:"ptp_node,omitempty"`
	SatelliteAccessInterface string           `protobuf:"bytes,7,opt,name=satellite_access_interface,json=satelliteAccessInterface,proto3" json:"satellite_access_interface,omitempty"`
	NtpNode                  string           `protobuf:"bytes,8,opt,name=ntp_node,json=ntpNode,proto3" json:"ntp_node,omitempty"`
	GnssReceiverId           *FsyncBagClockId `protobuf:"bytes,9,opt,name=gnss_receiver_id,json=gnssReceiverId,proto3" json:"gnss_receiver_id,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}         `json:"-"`
	XXX_unrecognized         []byte           `json:"-"`
	XXX_sizecache            int32            `json:"-"`
}

func (m *FsyncBagSourceId) Reset()         { *m = FsyncBagSourceId{} }
func (m *FsyncBagSourceId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSourceId) ProtoMessage()    {}
func (*FsyncBagSourceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{2}
}

func (m *FsyncBagSourceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSourceId.Unmarshal(m, b)
}
func (m *FsyncBagSourceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSourceId.Marshal(b, m, deterministic)
}
func (m *FsyncBagSourceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSourceId.Merge(m, src)
}
func (m *FsyncBagSourceId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSourceId.Size(m)
}
func (m *FsyncBagSourceId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSourceId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSourceId proto.InternalMessageInfo

func (m *FsyncBagSourceId) GetSourceClass() string {
	if m != nil {
		return m.SourceClass
	}
	return ""
}

func (m *FsyncBagSourceId) GetEthernetInterface() string {
	if m != nil {
		return m.EthernetInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetSonetInterface() string {
	if m != nil {
		return m.SonetInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetClockId() *FsyncBagClockId {
	if m != nil {
		return m.ClockId
	}
	return nil
}

func (m *FsyncBagSourceId) GetInternalClockId() *FsyncBagClockId {
	if m != nil {
		return m.InternalClockId
	}
	return nil
}

func (m *FsyncBagSourceId) GetPtpNode() string {
	if m != nil {
		return m.PtpNode
	}
	return ""
}

func (m *FsyncBagSourceId) GetSatelliteAccessInterface() string {
	if m != nil {
		return m.SatelliteAccessInterface
	}
	return ""
}

func (m *FsyncBagSourceId) GetNtpNode() string {
	if m != nil {
		return m.NtpNode
	}
	return ""
}

func (m *FsyncBagSourceId) GetGnssReceiverId() *FsyncBagClockId {
	if m != nil {
		return m.GnssReceiverId
	}
	return nil
}

type FsyncBagQl struct {
	QualityLevelOption      string   `protobuf:"bytes,1,opt,name=quality_level_option,json=qualityLevelOption,proto3" json:"quality_level_option,omitempty"`
	Option1Value            string   `protobuf:"bytes,2,opt,name=option1_value,json=option1Value,proto3" json:"option1_value,omitempty"`
	Option2Generation1Value string   `protobuf:"bytes,3,opt,name=option2_generation1_value,json=option2Generation1Value,proto3" json:"option2_generation1_value,omitempty"`
	Option2Generation2Value string   `protobuf:"bytes,4,opt,name=option2_generation2_value,json=option2Generation2Value,proto3" json:"option2_generation2_value,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *FsyncBagQl) Reset()         { *m = FsyncBagQl{} }
func (m *FsyncBagQl) String() string { return proto.CompactTextString(m) }
func (*FsyncBagQl) ProtoMessage()    {}
func (*FsyncBagQl) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{3}
}

func (m *FsyncBagQl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagQl.Unmarshal(m, b)
}
func (m *FsyncBagQl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagQl.Marshal(b, m, deterministic)
}
func (m *FsyncBagQl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagQl.Merge(m, src)
}
func (m *FsyncBagQl) XXX_Size() int {
	return xxx_messageInfo_FsyncBagQl.Size(m)
}
func (m *FsyncBagQl) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagQl.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagQl proto.InternalMessageInfo

func (m *FsyncBagQl) GetQualityLevelOption() string {
	if m != nil {
		return m.QualityLevelOption
	}
	return ""
}

func (m *FsyncBagQl) GetOption1Value() string {
	if m != nil {
		return m.Option1Value
	}
	return ""
}

func (m *FsyncBagQl) GetOption2Generation1Value() string {
	if m != nil {
		return m.Option2Generation1Value
	}
	return ""
}

func (m *FsyncBagQl) GetOption2Generation2Value() string {
	if m != nil {
		return m.Option2Generation2Value
	}
	return ""
}

type FsyncBagSpInfo struct {
	SelectionPoint            uint32   `protobuf:"varint,1,opt,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	SelectionPointDescription string   `protobuf:"bytes,2,opt,name=selection_point_description,json=selectionPointDescription,proto3" json:"selection_point_description,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *FsyncBagSpInfo) Reset()         { *m = FsyncBagSpInfo{} }
func (m *FsyncBagSpInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSpInfo) ProtoMessage()    {}
func (*FsyncBagSpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{4}
}

func (m *FsyncBagSpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSpInfo.Unmarshal(m, b)
}
func (m *FsyncBagSpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSpInfo.Marshal(b, m, deterministic)
}
func (m *FsyncBagSpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSpInfo.Merge(m, src)
}
func (m *FsyncBagSpInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSpInfo.Size(m)
}
func (m *FsyncBagSpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSpInfo proto.InternalMessageInfo

func (m *FsyncBagSpInfo) GetSelectionPoint() uint32 {
	if m != nil {
		return m.SelectionPoint
	}
	return 0
}

func (m *FsyncBagSpInfo) GetSelectionPointDescription() string {
	if m != nil {
		return m.SelectionPointDescription
	}
	return ""
}

type FsyncClockInfoV2 struct {
	Source                      *FsyncBagSourceId `protobuf:"bytes,50,opt,name=source,proto3" json:"source,omitempty"`
	State                       string            `protobuf:"bytes,51,opt,name=state,proto3" json:"state,omitempty"`
	DownReason                  string            `protobuf:"bytes,52,opt,name=down_reason,json=downReason,proto3" json:"down_reason,omitempty"`
	Description                 string            `protobuf:"bytes,53,opt,name=description,proto3" json:"description,omitempty"`
	Priority                    uint32            `protobuf:"varint,54,opt,name=priority,proto3" json:"priority,omitempty"`
	TimeOfDayPriority           uint32            `protobuf:"varint,55,opt,name=time_of_day_priority,json=timeOfDayPriority,proto3" json:"time_of_day_priority,omitempty"`
	SsmSupport                  bool              `protobuf:"varint,56,opt,name=ssm_support,json=ssmSupport,proto3" json:"ssm_support,omitempty"`
	SsmEnabled                  bool              `protobuf:"varint,57,opt,name=ssm_enabled,json=ssmEnabled,proto3" json:"ssm_enabled,omitempty"`
	Loopback                    bool              `protobuf:"varint,58,opt,name=loopback,proto3" json:"loopback,omitempty"`
	SelectionInput              bool              `protobuf:"varint,59,opt,name=selection_input,json=selectionInput,proto3" json:"selection_input,omitempty"`
	Squelched                   bool              `protobuf:"varint,60,opt,name=squelched,proto3" json:"squelched,omitempty"`
	SelectedSource              *FsyncBagSourceId `protobuf:"bytes,61,opt,name=selected_source,json=selectedSource,proto3" json:"selected_source,omitempty"`
	DampingState                string            `protobuf:"bytes,62,opt,name=damping_state,json=dampingState,proto3" json:"damping_state,omitempty"`
	DampingTime                 uint32            `protobuf:"varint,63,opt,name=damping_time,json=dampingTime,proto3" json:"damping_time,omitempty"`
	InputDisabled               bool              `protobuf:"varint,64,opt,name=input_disabled,json=inputDisabled,proto3" json:"input_disabled,omitempty"`
	OutputDisabled              bool              `protobuf:"varint,65,opt,name=output_disabled,json=outputDisabled,proto3" json:"output_disabled,omitempty"`
	WaitToRestoreTime           uint32            `protobuf:"varint,66,opt,name=wait_to_restore_time,json=waitToRestoreTime,proto3" json:"wait_to_restore_time,omitempty"`
	ClockTypeXr                 string            `protobuf:"bytes,67,opt,name=clock_type_xr,json=clockTypeXr,proto3" json:"clock_type_xr,omitempty"`
	SupportsFrequency           bool              `protobuf:"varint,68,opt,name=supports_frequency,json=supportsFrequency,proto3" json:"supports_frequency,omitempty"`
	SupportsTimeOfDay           bool              `protobuf:"varint,69,opt,name=supports_time_of_day,json=supportsTimeOfDay,proto3" json:"supports_time_of_day,omitempty"`
	QualityLevelReceived        *FsyncBagQl       `protobuf:"bytes,70,opt,name=quality_level_received,json=qualityLevelReceived,proto3" json:"quality_level_received,omitempty"`
	QualityLevelDamped          *FsyncBagQl       `protobuf:"bytes,71,opt,name=quality_level_damped,json=qualityLevelDamped,proto3" json:"quality_level_damped,omitempty"`
	QualityLevelEffectiveInput  *FsyncBagQl       `protobuf:"bytes,72,opt,name=quality_level_effective_input,json=qualityLevelEffectiveInput,proto3" json:"quality_level_effective_input,omitempty"`
	QualityLevelEffectiveOutput *FsyncBagQl       `protobuf:"bytes,73,opt,name=quality_level_effective_output,json=qualityLevelEffectiveOutput,proto3" json:"quality_level_effective_output,omitempty"`
	QualityLevelSelectedSource  *FsyncBagQl       `protobuf:"bytes,74,opt,name=quality_level_selected_source,json=qualityLevelSelectedSource,proto3" json:"quality_level_selected_source,omitempty"`
	SpaSelectionPoint           []*FsyncBagSpInfo `protobuf:"bytes,75,rep,name=spa_selection_point,json=spaSelectionPoint,proto3" json:"spa_selection_point,omitempty"`
	NodeSelectionPoint          []*FsyncBagSpInfo `protobuf:"bytes,76,rep,name=node_selection_point,json=nodeSelectionPoint,proto3" json:"node_selection_point,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}          `json:"-"`
	XXX_unrecognized            []byte            `json:"-"`
	XXX_sizecache               int32             `json:"-"`
}

func (m *FsyncClockInfoV2) Reset()         { *m = FsyncClockInfoV2{} }
func (m *FsyncClockInfoV2) String() string { return proto.CompactTextString(m) }
func (*FsyncClockInfoV2) ProtoMessage()    {}
func (*FsyncClockInfoV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_46d1abc5f9d3cefc, []int{5}
}

func (m *FsyncClockInfoV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncClockInfoV2.Unmarshal(m, b)
}
func (m *FsyncClockInfoV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncClockInfoV2.Marshal(b, m, deterministic)
}
func (m *FsyncClockInfoV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncClockInfoV2.Merge(m, src)
}
func (m *FsyncClockInfoV2) XXX_Size() int {
	return xxx_messageInfo_FsyncClockInfoV2.Size(m)
}
func (m *FsyncClockInfoV2) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncClockInfoV2.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncClockInfoV2 proto.InternalMessageInfo

func (m *FsyncClockInfoV2) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FsyncClockInfoV2) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *FsyncClockInfoV2) GetDownReason() string {
	if m != nil {
		return m.DownReason
	}
	return ""
}

func (m *FsyncClockInfoV2) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *FsyncClockInfoV2) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *FsyncClockInfoV2) GetTimeOfDayPriority() uint32 {
	if m != nil {
		return m.TimeOfDayPriority
	}
	return 0
}

func (m *FsyncClockInfoV2) GetSsmSupport() bool {
	if m != nil {
		return m.SsmSupport
	}
	return false
}

func (m *FsyncClockInfoV2) GetSsmEnabled() bool {
	if m != nil {
		return m.SsmEnabled
	}
	return false
}

func (m *FsyncClockInfoV2) GetLoopback() bool {
	if m != nil {
		return m.Loopback
	}
	return false
}

func (m *FsyncClockInfoV2) GetSelectionInput() bool {
	if m != nil {
		return m.SelectionInput
	}
	return false
}

func (m *FsyncClockInfoV2) GetSquelched() bool {
	if m != nil {
		return m.Squelched
	}
	return false
}

func (m *FsyncClockInfoV2) GetSelectedSource() *FsyncBagSourceId {
	if m != nil {
		return m.SelectedSource
	}
	return nil
}

func (m *FsyncClockInfoV2) GetDampingState() string {
	if m != nil {
		return m.DampingState
	}
	return ""
}

func (m *FsyncClockInfoV2) GetDampingTime() uint32 {
	if m != nil {
		return m.DampingTime
	}
	return 0
}

func (m *FsyncClockInfoV2) GetInputDisabled() bool {
	if m != nil {
		return m.InputDisabled
	}
	return false
}

func (m *FsyncClockInfoV2) GetOutputDisabled() bool {
	if m != nil {
		return m.OutputDisabled
	}
	return false
}

func (m *FsyncClockInfoV2) GetWaitToRestoreTime() uint32 {
	if m != nil {
		return m.WaitToRestoreTime
	}
	return 0
}

func (m *FsyncClockInfoV2) GetClockTypeXr() string {
	if m != nil {
		return m.ClockTypeXr
	}
	return ""
}

func (m *FsyncClockInfoV2) GetSupportsFrequency() bool {
	if m != nil {
		return m.SupportsFrequency
	}
	return false
}

func (m *FsyncClockInfoV2) GetSupportsTimeOfDay() bool {
	if m != nil {
		return m.SupportsTimeOfDay
	}
	return false
}

func (m *FsyncClockInfoV2) GetQualityLevelReceived() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelReceived
	}
	return nil
}

func (m *FsyncClockInfoV2) GetQualityLevelDamped() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelDamped
	}
	return nil
}

func (m *FsyncClockInfoV2) GetQualityLevelEffectiveInput() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelEffectiveInput
	}
	return nil
}

func (m *FsyncClockInfoV2) GetQualityLevelEffectiveOutput() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelEffectiveOutput
	}
	return nil
}

func (m *FsyncClockInfoV2) GetQualityLevelSelectedSource() *FsyncBagQl {
	if m != nil {
		return m.QualityLevelSelectedSource
	}
	return nil
}

func (m *FsyncClockInfoV2) GetSpaSelectionPoint() []*FsyncBagSpInfo {
	if m != nil {
		return m.SpaSelectionPoint
	}
	return nil
}

func (m *FsyncClockInfoV2) GetNodeSelectionPoint() []*FsyncBagSpInfo {
	if m != nil {
		return m.NodeSelectionPoint
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncClockInfoV2_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_clock_info_v2_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagQl)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_bag_ql")
	proto.RegisterType((*FsyncBagSpInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_bag_sp_info")
	proto.RegisterType((*FsyncClockInfoV2)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.clock_datas.clock_data.fsync_clock_info_v2")
}

func init() { proto.RegisterFile("fsync_clock_info_v2.proto", fileDescriptor_46d1abc5f9d3cefc) }

var fileDescriptor_46d1abc5f9d3cefc = []byte{
	// 1053 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xd6, 0xa6, 0xf9, 0x71, 0x8e, 0x63, 0xb7, 0x9e, 0x5a, 0x30, 0x49, 0x29, 0x75, 0x5d, 0x21,
	0x72, 0x83, 0x0b, 0x2e, 0xbf, 0xa5, 0x14, 0x4a, 0x92, 0x96, 0xd0, 0xaa, 0xa9, 0xd6, 0x51, 0x81,
	0x0b, 0x34, 0x9a, 0xec, 0x1e, 0x27, 0x43, 0xd6, 0x3b, 0x9b, 0x9d, 0xb1, 0x5b, 0x23, 0xee, 0x10,
	0x6f, 0x80, 0x90, 0x80, 0x07, 0x40, 0xbc, 0x02, 0xe2, 0x55, 0x78, 0x17, 0x34, 0x33, 0xfb, 0xe3,
	0x4d, 0x53, 0xee, 0x62, 0x6e, 0xac, 0x9d, 0xef, 0x7c, 0xe7, 0xcc, 0x37, 0xe7, 0x9c, 0x39, 0x23,
	0xc3, 0xfa, 0x50, 0x4d, 0xe3, 0x80, 0x05, 0x91, 0x0c, 0x8e, 0x99, 0x88, 0x87, 0x92, 0x4d, 0xfa,
	0xbd, 0x24, 0x95, 0x5a, 0x92, 0xa7, 0x81, 0x50, 0x81, 0x64, 0x42, 0x2a, 0xf6, 0x3c, 0x65, 0xc3,
	0x14, 0x4f, 0x2c, 0x55, 0x26, 0x98, 0xf6, 0xcc, 0x6a, 0x8c, 0x71, 0x30, 0x65, 0x06, 0x3b, 0x4a,
	0x65, 0x2c, 0xbe, 0xe7, 0x5a, 0xc8, 0xb8, 0x17, 0xcb, 0x10, 0x95, 0xfd, 0xed, 0xb9, 0xa0, 0x21,
	0xd7, 0x5c, 0xcd, 0x7c, 0x77, 0xbf, 0x05, 0x7a, 0xc6, 0xa6, 0xec, 0xe1, 0xce, 0x37, 0x03, 0x42,
	0x60, 0xd1, 0xb8, 0x52, 0xaf, 0xe3, 0x6d, 0xae, 0xfa, 0xf6, 0x9b, 0x5c, 0x05, 0x70, 0x4c, 0x3d,
	0x4d, 0x90, 0x2e, 0x58, 0xcb, 0xaa, 0x45, 0xf6, 0xa7, 0x09, 0x92, 0x26, 0x2c, 0x88, 0x90, 0x5e,
	0xe8, 0x78, 0x9b, 0x0d, 0x7f, 0x41, 0x84, 0xdd, 0xaf, 0x80, 0xb8, 0xf0, 0x07, 0xfc, 0x30, 0xdf,
	0x22, 0x3c, 0x33, 0xb0, 0xf3, 0x5c, 0xc8, 0x3d, 0xcb, 0x8d, 0x62, 0x3e, 0x42, 0x1b, 0x31, 0xdf,
	0xe8, 0x31, 0x1f, 0x61, 0xf7, 0xef, 0x25, 0xb8, 0x5c, 0x46, 0x56, 0x72, 0x9c, 0x06, 0x68, 0x42,
	0x5f, 0x87, 0xb5, 0x6c, 0x11, 0x44, 0x5c, 0xa9, 0x6c, 0x8b, 0xba, 0xc3, 0xb6, 0x0c, 0x44, 0xde,
	0x02, 0x82, 0xfa, 0x08, 0xd3, 0x18, 0x35, 0x13, 0xb1, 0xc6, 0x74, 0xc8, 0x83, 0xfc, 0x28, 0xad,
	0xdc, 0xb2, 0x9b, 0x1b, 0xc8, 0x9b, 0x70, 0x51, 0xc9, 0x2a, 0xd7, 0xa9, 0x69, 0x5a, 0xb8, 0x24,
	0xfe, 0xe4, 0x41, 0x2d, 0x3f, 0x22, 0x5d, 0xec, 0x78, 0x9b, 0xf5, 0xfe, 0x77, 0xbd, 0xf3, 0x29,
	0x5b, 0xef, 0xc5, 0xa4, 0xfa, 0x2b, 0xf6, 0x6b, 0x37, 0x24, 0xbf, 0x78, 0xd0, 0xb2, 0x5a, 0x63,
	0x1e, 0x15, 0x66, 0xba, 0x34, 0x77, 0x41, 0x17, 0x73, 0x11, 0x5b, 0x99, 0xb0, 0x75, 0xa8, 0x25,
	0x3a, 0x61, 0xb6, 0xf4, 0xcb, 0x36, 0x85, 0x2b, 0x89, 0x4e, 0x1e, 0x9b, 0xea, 0xdf, 0x81, 0x0d,
	0xc5, 0x35, 0x46, 0x91, 0xd0, 0xc8, 0x78, 0x10, 0xa0, 0x52, 0x33, 0xf9, 0x5e, 0xb1, 0x64, 0x5a,
	0x30, 0xee, 0x59, 0x42, 0x99, 0xf9, 0x75, 0xa8, 0xc5, 0x79, 0xe0, 0x9a, 0x0b, 0x1c, 0x67, 0x81,
	0x7f, 0xf6, 0xe0, 0xd2, 0x61, 0xac, 0x14, 0x4b, 0x31, 0x40, 0x31, 0xc1, 0xd4, 0xe4, 0x62, 0x75,
	0xee, 0xb9, 0x68, 0x1a, 0x0d, 0x7e, 0x26, 0x61, 0x37, 0xec, 0xfe, 0xe3, 0xc1, 0x5a, 0x49, 0x3b,
	0x89, 0xc8, 0xdb, 0xd0, 0x3e, 0x19, 0xf3, 0x48, 0xe8, 0x29, 0x8b, 0x70, 0x82, 0x11, 0x93, 0x89,
	0xd9, 0x2d, 0xeb, 0x5f, 0x92, 0xd9, 0x1e, 0x19, 0xd3, 0x9e, 0xb5, 0x90, 0x1b, 0xd0, 0x70, 0x9c,
	0x77, 0xd8, 0x84, 0x47, 0xe3, 0xbc, 0x83, 0xd7, 0x32, 0xf0, 0xa9, 0xc1, 0xc8, 0x6d, 0x58, 0x77,
	0xeb, 0x3e, 0x3b, 0xc4, 0x18, 0x53, 0x3e, 0xeb, 0xe0, 0xda, 0xf8, 0xd5, 0x8c, 0xf0, 0xa0, 0xb4,
	0xff, 0x87, 0x6f, 0x3f, 0xf3, 0x5d, 0x7c, 0x89, 0x6f, 0xdf, 0xfa, 0x76, 0x7f, 0x80, 0xd6, 0xcc,
	0xed, 0x4c, 0xec, 0x5c, 0xb1, 0x37, 0x09, 0x23, 0x0c, 0x0c, 0x8f, 0x25, 0x52, 0xc4, 0xda, 0x1e,
	0xaf, 0xe1, 0x37, 0x0b, 0xf8, 0x89, 0x41, 0xc9, 0x5d, 0xb8, 0x72, 0x8a, 0xc8, 0x42, 0x54, 0x41,
	0x2a, 0x5c, 0x4e, 0xdc, 0x41, 0xd7, 0xab, 0x4e, 0xdb, 0x25, 0xa1, 0xfb, 0xd7, 0xa5, 0x7c, 0x38,
	0x54, 0xa6, 0x1a, 0xf9, 0xd1, 0x83, 0x65, 0x37, 0x09, 0x68, 0xdf, 0xb6, 0xc0, 0xf1, 0xf9, 0xb7,
	0x40, 0x31, 0x9a, 0xfc, 0x6c, 0x6b, 0xd2, 0x86, 0x25, 0xa5, 0xb9, 0x46, 0x7a, 0xcb, 0x9e, 0xc3,
	0x2d, 0xc8, 0x35, 0xa8, 0x87, 0xf2, 0x59, 0xcc, 0x52, 0xe4, 0x4a, 0xc6, 0xf4, 0x5d, 0x6b, 0x03,
	0x03, 0xf9, 0x16, 0x21, 0x1d, 0xa8, 0xcf, 0x26, 0xe1, 0x3d, 0x37, 0xd8, 0x66, 0x20, 0xb2, 0x01,
	0xb5, 0x24, 0x15, 0x32, 0x15, 0x7a, 0x4a, 0xdf, 0xb7, 0x89, 0x2d, 0xd6, 0xe4, 0x26, 0xb4, 0xb5,
	0x18, 0x21, 0x93, 0x43, 0x16, 0xf2, 0x29, 0x2b, 0x78, 0x1f, 0x58, 0x5e, 0xcb, 0xd8, 0xf6, 0x86,
	0xdb, 0x7c, 0xfa, 0x24, 0x77, 0xb8, 0x06, 0x75, 0xa5, 0x46, 0x4c, 0x8d, 0x93, 0x44, 0xa6, 0x9a,
	0x7e, 0xd8, 0xf1, 0x36, 0x6b, 0x3e, 0x28, 0x35, 0x1a, 0x38, 0x24, 0x27, 0x60, 0xcc, 0x0f, 0x22,
	0x0c, 0xe9, 0x47, 0x05, 0x61, 0xc7, 0x21, 0x46, 0x4e, 0x24, 0x65, 0x72, 0xc0, 0x83, 0x63, 0x7a,
	0xdb, 0x5a, 0x8b, 0x75, 0xb5, 0x15, 0x44, 0x9c, 0x8c, 0x35, 0xfd, 0xd8, 0x52, 0xca, 0x56, 0xd8,
	0x35, 0x28, 0x79, 0x0d, 0x56, 0xd5, 0xc9, 0x18, 0xa3, 0xe0, 0x08, 0x43, 0x7a, 0xc7, 0x52, 0x4a,
	0xc0, 0xdc, 0xee, 0x2c, 0x0e, 0x86, 0x59, 0xa6, 0xe9, 0x27, 0xf3, 0xaf, 0x6c, 0x33, 0xd7, 0x30,
	0x70, 0x15, 0xbe, 0x01, 0x8d, 0x90, 0x8f, 0x12, 0x11, 0x1f, 0x32, 0x57, 0xe9, 0xbb, 0xee, 0x6a,
	0x66, 0xe0, 0xc0, 0x16, 0xfc, 0x3a, 0xe4, 0x6b, 0x66, 0xb2, 0x4f, 0x3f, 0xb5, 0x95, 0xa8, 0x67,
	0xd8, 0xbe, 0x18, 0x21, 0x79, 0x03, 0x9a, 0x36, 0x37, 0x2c, 0x14, 0xca, 0x65, 0xf9, 0x33, 0x9b,
	0x81, 0x86, 0x45, 0xb7, 0x33, 0xd0, 0x24, 0x53, 0x8e, 0x75, 0x85, 0x77, 0xcf, 0x25, 0xd3, 0xc1,
	0x05, 0xf1, 0x26, 0xb4, 0x9f, 0x71, 0xa1, 0x99, 0x96, 0x2c, 0x45, 0xa5, 0x65, 0x8a, 0x6e, 0xeb,
	0xcf, 0x5d, 0x13, 0x18, 0xdb, 0xbe, 0xf4, 0x9d, 0xc5, 0x0a, 0xe8, 0x42, 0xa3, 0x7c, 0xed, 0xd9,
	0xf3, 0x94, 0x6e, 0xb9, 0xae, 0x2b, 0x1e, 0xfc, 0xaf, 0x53, 0xf3, 0x9c, 0x66, 0x4d, 0xa2, 0x58,
	0x91, 0x58, 0xba, 0x6d, 0x05, 0xb4, 0x72, 0xcb, 0xfd, 0xdc, 0x60, 0x34, 0x14, 0xf4, 0x99, 0x8e,
	0xa4, 0x3b, 0x55, 0x87, 0xfd, 0xbc, 0x21, 0xc9, 0x6f, 0x1e, 0xbc, 0x52, 0x1d, 0x8d, 0xd9, 0x28,
	0x0f, 0xe9, 0x7d, 0x5b, 0xea, 0xf0, 0xfc, 0x4b, 0x7d, 0x12, 0xf9, 0xed, 0xd9, 0x11, 0x9c, 0x4d,
	0x72, 0xfb, 0xd6, 0x9e, 0x9a, 0xdb, 0xa6, 0x7e, 0x18, 0xd2, 0x07, 0x73, 0x94, 0x56, 0x79, 0x1d,
	0xb6, 0xed, 0xfe, 0xe4, 0x0f, 0x0f, 0xae, 0x56, 0x85, 0xe1, 0x70, 0x68, 0x2e, 0xd6, 0x04, 0xb3,
	0xfb, 0xf6, 0xc5, 0x1c, 0x15, 0x6e, 0xcc, 0x2a, 0xdc, 0xc9, 0x85, 0xb8, 0x1b, 0xfe, 0xa7, 0x07,
	0xaf, 0xbf, 0x4c, 0xa9, 0xeb, 0x5f, 0xba, 0x3b, 0x47, 0xa9, 0x57, 0xce, 0x94, 0xba, 0x67, 0x85,
	0x9c, 0x91, 0xd5, 0xd3, 0xd3, 0xe7, 0xcb, 0xff, 0x2b, 0xab, 0x83, 0xea, 0x08, 0xfa, 0xd5, 0x83,
	0xcb, 0x2a, 0xe1, 0xec, 0xf4, 0x83, 0xfb, 0xb0, 0x73, 0x61, 0xb3, 0xde, 0x17, 0x73, 0x98, 0x8e,
	0xee, 0xd1, 0xf7, 0x5b, 0x2a, 0xe1, 0x83, 0xea, 0xf3, 0xfe, 0xbb, 0x07, 0x6d, 0x13, 0xe4, 0x05,
	0x71, 0x8f, 0xe6, 0x2d, 0x8e, 0x18, 0x7a, 0x55, 0xdd, 0xc1, 0xb2, 0xfd, 0xc3, 0x75, 0xeb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x28, 0xe0, 0x33, 0x8d, 0x0d, 0x00, 0x00,
}
