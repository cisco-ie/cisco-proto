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
// source: fsync_sp_input_info.proto

package cisco_ios_xr_ncs4k_freqsync_oper_frequency_synchronization_nodes_node_selection_point_inputs_selection_point_input

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

type FsyncSpInputInfo_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SelectionPoint       uint32   `protobuf:"varint,2,opt,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	StreamType           string   `protobuf:"bytes,3,opt,name=stream_type,json=streamType,proto3" json:"stream_type,omitempty"`
	SourceType           string   `protobuf:"bytes,4,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	Interface            string   `protobuf:"bytes,5,opt,name=interface,proto3" json:"interface,omitempty"`
	ClockPort            uint32   `protobuf:"varint,6,opt,name=clock_port,json=clockPort,proto3" json:"clock_port,omitempty"`
	LastNode             string   `protobuf:"bytes,7,opt,name=last_node,json=lastNode,proto3" json:"last_node,omitempty"`
	LastSelectionPoint   uint32   `protobuf:"varint,8,opt,name=last_selection_point,json=lastSelectionPoint,proto3" json:"last_selection_point,omitempty"`
	OutputId             uint32   `protobuf:"varint,9,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncSpInputInfo_KEYS) Reset()         { *m = FsyncSpInputInfo_KEYS{} }
func (m *FsyncSpInputInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncSpInputInfo_KEYS) ProtoMessage()    {}
func (*FsyncSpInputInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0954f74c333e7db1, []int{0}
}

func (m *FsyncSpInputInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSpInputInfo_KEYS.Unmarshal(m, b)
}
func (m *FsyncSpInputInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSpInputInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncSpInputInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSpInputInfo_KEYS.Merge(m, src)
}
func (m *FsyncSpInputInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncSpInputInfo_KEYS.Size(m)
}
func (m *FsyncSpInputInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSpInputInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSpInputInfo_KEYS proto.InternalMessageInfo

func (m *FsyncSpInputInfo_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *FsyncSpInputInfo_KEYS) GetSelectionPoint() uint32 {
	if m != nil {
		return m.SelectionPoint
	}
	return 0
}

func (m *FsyncSpInputInfo_KEYS) GetStreamType() string {
	if m != nil {
		return m.StreamType
	}
	return ""
}

func (m *FsyncSpInputInfo_KEYS) GetSourceType() string {
	if m != nil {
		return m.SourceType
	}
	return ""
}

func (m *FsyncSpInputInfo_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *FsyncSpInputInfo_KEYS) GetClockPort() uint32 {
	if m != nil {
		return m.ClockPort
	}
	return 0
}

func (m *FsyncSpInputInfo_KEYS) GetLastNode() string {
	if m != nil {
		return m.LastNode
	}
	return ""
}

func (m *FsyncSpInputInfo_KEYS) GetLastSelectionPoint() uint32 {
	if m != nil {
		return m.LastSelectionPoint
	}
	return 0
}

func (m *FsyncSpInputInfo_KEYS) GetOutputId() uint32 {
	if m != nil {
		return m.OutputId
	}
	return 0
}

type FsyncBagSpId struct {
	SelectionPointType        uint32   `protobuf:"varint,1,opt,name=selection_point_type,json=selectionPointType,proto3" json:"selection_point_type,omitempty"`
	SelectionPointDescription string   `protobuf:"bytes,2,opt,name=selection_point_description,json=selectionPointDescription,proto3" json:"selection_point_description,omitempty"`
	Node                      string   `protobuf:"bytes,3,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *FsyncBagSpId) Reset()         { *m = FsyncBagSpId{} }
func (m *FsyncBagSpId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSpId) ProtoMessage()    {}
func (*FsyncBagSpId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0954f74c333e7db1, []int{1}
}

func (m *FsyncBagSpId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSpId.Unmarshal(m, b)
}
func (m *FsyncBagSpId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSpId.Marshal(b, m, deterministic)
}
func (m *FsyncBagSpId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSpId.Merge(m, src)
}
func (m *FsyncBagSpId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSpId.Size(m)
}
func (m *FsyncBagSpId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSpId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSpId proto.InternalMessageInfo

func (m *FsyncBagSpId) GetSelectionPointType() uint32 {
	if m != nil {
		return m.SelectionPointType
	}
	return 0
}

func (m *FsyncBagSpId) GetSelectionPointDescription() string {
	if m != nil {
		return m.SelectionPointDescription
	}
	return ""
}

func (m *FsyncBagSpId) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
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
	return fileDescriptor_0954f74c333e7db1, []int{2}
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
	return fileDescriptor_0954f74c333e7db1, []int{3}
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

type FsyncBagLastSpId struct {
	SelectionPoint       *FsyncBagSpId `protobuf:"bytes,1,opt,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	OutputId             uint32        `protobuf:"varint,2,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FsyncBagLastSpId) Reset()         { *m = FsyncBagLastSpId{} }
func (m *FsyncBagLastSpId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagLastSpId) ProtoMessage()    {}
func (*FsyncBagLastSpId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0954f74c333e7db1, []int{4}
}

func (m *FsyncBagLastSpId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagLastSpId.Unmarshal(m, b)
}
func (m *FsyncBagLastSpId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagLastSpId.Marshal(b, m, deterministic)
}
func (m *FsyncBagLastSpId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagLastSpId.Merge(m, src)
}
func (m *FsyncBagLastSpId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagLastSpId.Size(m)
}
func (m *FsyncBagLastSpId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagLastSpId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagLastSpId proto.InternalMessageInfo

func (m *FsyncBagLastSpId) GetSelectionPoint() *FsyncBagSpId {
	if m != nil {
		return m.SelectionPoint
	}
	return nil
}

func (m *FsyncBagLastSpId) GetOutputId() uint32 {
	if m != nil {
		return m.OutputId
	}
	return 0
}

type FsyncBagStreamId struct {
	StreamInput          string            `protobuf:"bytes,1,opt,name=stream_input,json=streamInput,proto3" json:"stream_input,omitempty"`
	SourceId             *FsyncBagSourceId `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SelectionPointId     *FsyncBagLastSpId `protobuf:"bytes,3,opt,name=selection_point_id,json=selectionPointId,proto3" json:"selection_point_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncBagStreamId) Reset()         { *m = FsyncBagStreamId{} }
func (m *FsyncBagStreamId) String() string { return proto.CompactTextString(m) }
func (*FsyncBagStreamId) ProtoMessage()    {}
func (*FsyncBagStreamId) Descriptor() ([]byte, []int) {
	return fileDescriptor_0954f74c333e7db1, []int{5}
}

func (m *FsyncBagStreamId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagStreamId.Unmarshal(m, b)
}
func (m *FsyncBagStreamId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagStreamId.Marshal(b, m, deterministic)
}
func (m *FsyncBagStreamId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagStreamId.Merge(m, src)
}
func (m *FsyncBagStreamId) XXX_Size() int {
	return xxx_messageInfo_FsyncBagStreamId.Size(m)
}
func (m *FsyncBagStreamId) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagStreamId.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagStreamId proto.InternalMessageInfo

func (m *FsyncBagStreamId) GetStreamInput() string {
	if m != nil {
		return m.StreamInput
	}
	return ""
}

func (m *FsyncBagStreamId) GetSourceId() *FsyncBagSourceId {
	if m != nil {
		return m.SourceId
	}
	return nil
}

func (m *FsyncBagStreamId) GetSelectionPointId() *FsyncBagLastSpId {
	if m != nil {
		return m.SelectionPointId
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
	return fileDescriptor_0954f74c333e7db1, []int{6}
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

type FsyncSpInputInfo struct {
	InputSelectionPoint  *FsyncBagSpId     `protobuf:"bytes,50,opt,name=input_selection_point,json=inputSelectionPoint,proto3" json:"input_selection_point,omitempty"`
	Stream               *FsyncBagStreamId `protobuf:"bytes,51,opt,name=stream,proto3" json:"stream,omitempty"`
	OriginalSource       *FsyncBagSourceId `protobuf:"bytes,52,opt,name=original_source,json=originalSource,proto3" json:"original_source,omitempty"`
	SupportsFrequency    bool              `protobuf:"varint,53,opt,name=supports_frequency,json=supportsFrequency,proto3" json:"supports_frequency,omitempty"`
	SupportsTimeOfDay    bool              `protobuf:"varint,54,opt,name=supports_time_of_day,json=supportsTimeOfDay,proto3" json:"supports_time_of_day,omitempty"`
	QualityLevel         *FsyncBagQl       `protobuf:"bytes,55,opt,name=quality_level,json=qualityLevel,proto3" json:"quality_level,omitempty"`
	Priority             uint32            `protobuf:"varint,56,opt,name=priority,proto3" json:"priority,omitempty"`
	TimeOfDayPriority    uint32            `protobuf:"varint,57,opt,name=time_of_day_priority,json=timeOfDayPriority,proto3" json:"time_of_day_priority,omitempty"`
	Selected             bool              `protobuf:"varint,58,opt,name=selected,proto3" json:"selected,omitempty"`
	OutputId             uint32            `protobuf:"varint,59,opt,name=output_id,json=outputId,proto3" json:"output_id,omitempty"`
	PlatformStatus       string            `protobuf:"bytes,60,opt,name=platform_status,json=platformStatus,proto3" json:"platform_status,omitempty"`
	PlatformFailedReason string            `protobuf:"bytes,61,opt,name=platform_failed_reason,json=platformFailedReason,proto3" json:"platform_failed_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncSpInputInfo) Reset()         { *m = FsyncSpInputInfo{} }
func (m *FsyncSpInputInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncSpInputInfo) ProtoMessage()    {}
func (*FsyncSpInputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0954f74c333e7db1, []int{7}
}

func (m *FsyncSpInputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSpInputInfo.Unmarshal(m, b)
}
func (m *FsyncSpInputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSpInputInfo.Marshal(b, m, deterministic)
}
func (m *FsyncSpInputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSpInputInfo.Merge(m, src)
}
func (m *FsyncSpInputInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncSpInputInfo.Size(m)
}
func (m *FsyncSpInputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSpInputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSpInputInfo proto.InternalMessageInfo

func (m *FsyncSpInputInfo) GetInputSelectionPoint() *FsyncBagSpId {
	if m != nil {
		return m.InputSelectionPoint
	}
	return nil
}

func (m *FsyncSpInputInfo) GetStream() *FsyncBagStreamId {
	if m != nil {
		return m.Stream
	}
	return nil
}

func (m *FsyncSpInputInfo) GetOriginalSource() *FsyncBagSourceId {
	if m != nil {
		return m.OriginalSource
	}
	return nil
}

func (m *FsyncSpInputInfo) GetSupportsFrequency() bool {
	if m != nil {
		return m.SupportsFrequency
	}
	return false
}

func (m *FsyncSpInputInfo) GetSupportsTimeOfDay() bool {
	if m != nil {
		return m.SupportsTimeOfDay
	}
	return false
}

func (m *FsyncSpInputInfo) GetQualityLevel() *FsyncBagQl {
	if m != nil {
		return m.QualityLevel
	}
	return nil
}

func (m *FsyncSpInputInfo) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *FsyncSpInputInfo) GetTimeOfDayPriority() uint32 {
	if m != nil {
		return m.TimeOfDayPriority
	}
	return 0
}

func (m *FsyncSpInputInfo) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

func (m *FsyncSpInputInfo) GetOutputId() uint32 {
	if m != nil {
		return m.OutputId
	}
	return 0
}

func (m *FsyncSpInputInfo) GetPlatformStatus() string {
	if m != nil {
		return m.PlatformStatus
	}
	return ""
}

func (m *FsyncSpInputInfo) GetPlatformFailedReason() string {
	if m != nil {
		return m.PlatformFailedReason
	}
	return ""
}

func init() {
	proto.RegisterType((*FsyncSpInputInfo_KEYS)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_sp_input_info_KEYS")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_sp_id")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagLastSpId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_last_sp_id")
	proto.RegisterType((*FsyncBagStreamId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_stream_id")
	proto.RegisterType((*FsyncBagQl)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_ql")
	proto.RegisterType((*FsyncSpInputInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_sp_input_info")
}

func init() { proto.RegisterFile("fsync_sp_input_info.proto", fileDescriptor_0954f74c333e7db1) }

var fileDescriptor_0954f74c333e7db1 = []byte{
	// 994 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4f, 0x6f, 0x23, 0x35,
	0x14, 0xd7, 0xa4, 0x25, 0x9d, 0xbc, 0xb6, 0xe9, 0xd6, 0x5b, 0x60, 0xd2, 0x82, 0x28, 0xe1, 0x40,
	0x2f, 0x04, 0xc8, 0x96, 0x7f, 0xcb, 0x82, 0x84, 0x76, 0x59, 0x14, 0x81, 0x76, 0xab, 0xe9, 0x0a,
	0xc4, 0xc9, 0xf2, 0xce, 0x38, 0x5d, 0xab, 0x93, 0xf1, 0xd4, 0x76, 0x2a, 0xc2, 0x09, 0x09, 0x09,
	0xb8, 0x23, 0x84, 0xc4, 0x09, 0x21, 0xc1, 0x8d, 0x03, 0xe2, 0x00, 0x67, 0xbe, 0x07, 0x27, 0xbe,
	0x08, 0xf2, 0xf3, 0x78, 0x26, 0x93, 0x16, 0x8e, 0xdb, 0x5c, 0xa2, 0xf1, 0xef, 0xfd, 0xde, 0xf3,
	0xfb, 0xe3, 0xf7, 0xec, 0x40, 0x6f, 0xac, 0x67, 0x79, 0x42, 0x75, 0x41, 0x45, 0x5e, 0x4c, 0x0d,
	0x15, 0xf9, 0x58, 0x0e, 0x0a, 0x25, 0x8d, 0x24, 0x2a, 0x11, 0x3a, 0x91, 0x54, 0x48, 0x4d, 0x3f,
	0x53, 0x34, 0x4f, 0xf4, 0xe1, 0x29, 0x1d, 0x2b, 0x7e, 0x86, 0x0a, 0xb2, 0xe0, 0x6a, 0x60, 0x57,
	0x53, 0x9e, 0x27, 0x33, 0x6a, 0xb1, 0x47, 0x4a, 0xe6, 0xe2, 0x73, 0x66, 0x84, 0xcc, 0x07, 0xb9,
	0x4c, 0xb9, 0xc6, 0xdf, 0x81, 0xe6, 0x19, 0x4f, 0x2c, 0x48, 0x0b, 0x29, 0x72, 0xe3, 0xf6, 0xd1,
	0x97, 0xc3, 0xfd, 0xbf, 0x5a, 0x10, 0x5d, 0xe2, 0x11, 0xfd, 0xf0, 0xfd, 0x4f, 0x8f, 0x09, 0x81,
	0x55, 0x6b, 0x31, 0x0a, 0xf6, 0x83, 0x83, 0x4e, 0x8c, 0xdf, 0xe4, 0x45, 0xd8, 0x5a, 0xb0, 0x14,
	0xb5, 0xf6, 0x83, 0x83, 0xcd, 0xb8, 0x5b, 0xc1, 0x47, 0x16, 0x25, 0xcf, 0xc1, 0xba, 0x36, 0x8a,
	0xb3, 0x09, 0x35, 0xb3, 0x82, 0x47, 0x2b, 0x68, 0x03, 0x1c, 0xf4, 0x60, 0x56, 0x70, 0x24, 0xc8,
	0xa9, 0x4a, 0xb8, 0x23, 0xac, 0x96, 0x04, 0x84, 0x90, 0xf0, 0x0c, 0x74, 0x44, 0x6e, 0xb8, 0x1a,
	0xb3, 0x84, 0x47, 0x4f, 0xa0, 0xb8, 0x06, 0xc8, 0xb3, 0x00, 0x49, 0x26, 0x93, 0x53, 0x5a, 0x48,
	0x65, 0xa2, 0x36, 0xfa, 0xd0, 0x41, 0xe4, 0x48, 0x2a, 0x43, 0xf6, 0xa0, 0x93, 0x31, 0x6d, 0x28,
	0x06, 0xb0, 0x86, 0xca, 0xa1, 0x05, 0xee, 0xd9, 0x20, 0x5e, 0x81, 0x1d, 0x14, 0x2e, 0x46, 0x12,
	0xa2, 0x15, 0x62, 0x65, 0xc7, 0xcd, 0x68, 0xf6, 0xa0, 0x23, 0xa7, 0x06, 0xd3, 0x93, 0x46, 0x1d,
	0xa4, 0x85, 0x0e, 0x18, 0xa5, 0xfd, 0xef, 0x03, 0xd8, 0x72, 0x49, 0x7c, 0xc8, 0x4e, 0x30, 0x91,
	0xa9, 0xdd, 0x62, 0x31, 0xe3, 0x18, 0x66, 0xe0, 0xb6, 0x68, 0x26, 0x0b, 0xc3, 0x7d, 0x17, 0xf6,
	0x16, 0x35, 0x52, 0xae, 0x13, 0x25, 0x0a, 0x8b, 0x60, 0x96, 0x3b, 0x71, 0xaf, 0xa9, 0x78, 0xa7,
	0x26, 0x54, 0xd5, 0x5a, 0xa9, 0xab, 0xd5, 0xff, 0x04, 0x48, 0xed, 0x98, 0x4b, 0x97, 0x48, 0x2f,
	0xad, 0x6b, 0x17, 0x5a, 0x22, 0x2d, 0x4b, 0xd9, 0x12, 0x69, 0x9d, 0xde, 0x9c, 0x4d, 0xbc, 0x4d,
	0x97, 0xde, 0x7b, 0x6c, 0xc2, 0xfb, 0xdf, 0xb6, 0xe1, 0xfa, 0x5c, 0xc8, 0xae, 0x8e, 0x22, 0x25,
	0xcf, 0xc3, 0x46, 0xb9, 0x48, 0x32, 0xa6, 0x75, 0xb9, 0x45, 0x59, 0xe8, 0xdb, 0x16, 0x22, 0x2f,
	0x01, 0xe1, 0xe6, 0x11, 0x57, 0x39, 0xb7, 0x67, 0xcd, 0xd7, 0xd7, 0x85, 0xb7, 0xed, 0x25, 0xa3,
	0xaa, 0xce, 0xf6, 0xc0, 0xc9, 0x26, 0xd7, 0x79, 0xd3, 0x45, 0xb8, 0x26, 0xfe, 0x14, 0x40, 0xe8,
	0x43, 0xc4, 0xd3, 0xb4, 0x3e, 0xfc, 0x2a, 0x18, 0x3c, 0xfe, 0x9e, 0x1a, 0x5c, 0xcc, 0x78, 0xbc,
	0x86, 0x5f, 0xa3, 0x94, 0xfc, 0x1e, 0xc0, 0x36, 0x06, 0x92, 0xb3, 0xac, 0x12, 0xe3, 0xe1, 0x5e,
	0x22, 0x6f, 0xb7, 0xbc, 0x87, 0xb7, 0x4b, 0xaf, 0x7b, 0x10, 0x16, 0xa6, 0x70, 0xbd, 0xd4, 0xc6,
	0xe4, 0xaf, 0x15, 0xa6, 0xc0, 0x56, 0xba, 0x05, 0xbb, 0x9a, 0x19, 0x9e, 0x65, 0xc2, 0x70, 0xca,
	0x92, 0x84, 0x6b, 0x3d, 0x57, 0x29, 0xd7, 0x78, 0x51, 0xc5, 0x78, 0x0f, 0x09, 0x75, 0xcd, 0x7a,
	0x10, 0xe6, 0xde, 0x70, 0xe8, 0x0c, 0xe7, 0xa5, 0xe1, 0xdf, 0x02, 0xb8, 0x76, 0x92, 0x6b, 0x4d,
	0x15, 0x4f, 0xb8, 0x38, 0xe7, 0xca, 0x77, 0xde, 0x12, 0x25, 0xaa, 0x6b, 0x1d, 0x8c, 0x4b, 0xff,
	0x46, 0x69, 0xff, 0x9f, 0x00, 0x76, 0x6a, 0x9a, 0x1b, 0x31, 0x38, 0x0d, 0x7e, 0x0d, 0x2e, 0x8e,
	0xcd, 0x00, 0x63, 0xf9, 0xf2, 0x8a, 0x63, 0x41, 0xff, 0x2e, 0x0c, 0xef, 0xc6, 0xb8, 0x6b, 0x2d,
	0x8c, 0xbb, 0x1f, 0x56, 0x1a, 0xbd, 0xef, 0x86, 0x7c, 0xd9, 0xfb, 0xe5, 0xc2, 0x6e, 0x52, 0xf5,
	0x3e, 0x62, 0x23, 0x0b, 0x91, 0x9f, 0x03, 0xe8, 0x54, 0xc3, 0x02, 0x0d, 0xaf, 0x0f, 0xbf, 0xbe,
	0xea, 0x0c, 0x78, 0x7f, 0xe2, 0xd0, 0x7d, 0x8e, 0x52, 0xf2, 0x47, 0x00, 0xe4, 0x82, 0x5e, 0x8a,
	0x83, 0x67, 0x7d, 0xf8, 0xcd, 0x15, 0x3b, 0x5c, 0x9f, 0xab, 0xf8, 0x5a, 0xb3, 0x6e, 0xa3, 0xb4,
	0xff, 0x77, 0x00, 0x1b, 0x35, 0xf5, 0x2c, 0xb3, 0x17, 0xd1, 0xd9, 0x94, 0x65, 0xc2, 0xcc, 0x68,
	0xc6, 0xcf, 0x79, 0x46, 0xa5, 0xbb, 0x4f, 0x5c, 0x75, 0x48, 0x29, 0xfb, 0xc8, 0x8a, 0xee, 0xbb,
	0x8b, 0xe4, 0x05, 0xd8, 0x74, 0x9c, 0x57, 0xe9, 0x39, 0xcb, 0xa6, 0x7e, 0x36, 0x6f, 0x94, 0xe0,
	0xc7, 0x16, 0x23, 0x37, 0xa1, 0xe7, 0xd6, 0x43, 0x7a, 0xc2, 0x73, 0xae, 0xd8, 0xbc, 0x82, 0x1b,
	0xd0, 0x4f, 0x97, 0x84, 0x0f, 0x6a, 0xf9, 0xff, 0xe8, 0x0e, 0x4b, 0xdd, 0xd5, 0xff, 0xd0, 0x1d,
	0xa2, 0x6e, 0xff, 0xbb, 0xd0, 0x1f, 0xbe, 0xc6, 0x83, 0x85, 0xfc, 0x19, 0xc0, 0x93, 0x6e, 0xb9,
	0xd8, 0x67, 0xc3, 0x25, 0xea, 0xb3, 0xeb, 0x08, 0x2f, 0xbc, 0x2d, 0x7e, 0x0c, 0xa0, 0xed, 0x9a,
	0x24, 0xba, 0xb1, 0x1c, 0x1d, 0xe1, 0x5b, 0x3a, 0x2e, 0xfd, 0xb2, 0xc3, 0x78, 0x4b, 0x2a, 0x71,
	0x22, 0xec, 0xb5, 0xe5, 0xba, 0x24, 0x3a, 0x5c, 0xb2, 0xee, 0xed, 0x7a, 0x07, 0x8f, 0x11, 0xb2,
	0xef, 0x0c, 0x3d, 0x2d, 0xec, 0xeb, 0x50, 0xd3, 0xca, 0x83, 0xe8, 0xb5, 0xfd, 0xe0, 0x20, 0x8c,
	0xb7, 0xbd, 0xe4, 0xae, 0x17, 0x90, 0x97, 0x61, 0xa7, 0xa2, 0x1b, 0x31, 0xe1, 0x54, 0x8e, 0x69,
	0xca, 0x66, 0xd1, 0xeb, 0x4d, 0x85, 0x07, 0x62, 0xc2, 0xef, 0x8f, 0xef, 0xb0, 0x19, 0xf9, 0x25,
	0x80, 0xcd, 0x46, 0x67, 0x45, 0x6f, 0x60, 0x46, 0xbe, 0xb8, 0xe2, 0x8c, 0x9c, 0x65, 0xf1, 0xc6,
	0x7c, 0x57, 0x93, 0x5d, 0x08, 0x0b, 0x25, 0xa4, 0x12, 0x66, 0x16, 0xbd, 0xe9, 0x66, 0xb9, 0x5f,
	0xdb, 0xa8, 0xe7, 0x82, 0xa5, 0x15, 0xef, 0x2d, 0xe4, 0x6d, 0x1b, 0x1f, 0xed, 0x91, 0x57, 0xd8,
	0x85, 0xd0, 0x39, 0xc0, 0xd3, 0xe8, 0x26, 0xa6, 0xa6, 0x5a, 0x37, 0x6f, 0x8d, 0xb7, 0x9b, 0xb7,
	0x86, 0x7d, 0xc7, 0x15, 0x19, 0x33, 0x63, 0xa9, 0x26, 0x54, 0x1b, 0x66, 0xa6, 0x3a, 0xba, 0xe5,
	0xde, 0x71, 0x1e, 0x3e, 0x46, 0x94, 0x1c, 0xc2, 0x53, 0x15, 0x71, 0xcc, 0x44, 0xc6, 0x53, 0xaa,
	0x38, 0xd3, 0x32, 0x8f, 0xde, 0x41, 0xfe, 0x8e, 0x97, 0xde, 0x45, 0x61, 0x8c, 0xb2, 0x87, 0x6d,
	0xfc, 0x0f, 0x75, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x72, 0x66, 0x05, 0x17, 0x60, 0x0d,
	0x00, 0x00,
}