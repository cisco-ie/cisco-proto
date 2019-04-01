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

package cisco_ios_xr_freqsync_oper_frequency_synchronization_nodes_node_selection_point_inputs_selection_point_input

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
	proto.RegisterType((*FsyncSpInputInfo_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_sp_input_info_KEYS")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_sp_id")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagLastSpId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_last_sp_id")
	proto.RegisterType((*FsyncBagStreamId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_stream_id")
	proto.RegisterType((*FsyncBagQl)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_bag_ql")
	proto.RegisterType((*FsyncSpInputInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.nodes.node.selection_point_inputs.selection_point_input.fsync_sp_input_info")
}

func init() { proto.RegisterFile("fsync_sp_input_info.proto", fileDescriptor_0954f74c333e7db1) }

var fileDescriptor_0954f74c333e7db1 = []byte{
	// 986 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4f, 0x6f, 0x23, 0x35,
	0x14, 0xd7, 0xa4, 0x25, 0x9d, 0xbc, 0xb6, 0xe9, 0xd6, 0x5b, 0x60, 0xd2, 0x82, 0x28, 0xe1, 0x40,
	0x2f, 0x04, 0xc8, 0x2e, 0xff, 0x96, 0x05, 0x09, 0xed, 0xb2, 0x28, 0x02, 0xed, 0x56, 0xd3, 0x15,
	0x88, 0x93, 0xe5, 0x9d, 0x71, 0xba, 0x16, 0x93, 0xf1, 0xd4, 0x76, 0x2a, 0x66, 0x25, 0x24, 0x96,
	0x0f, 0xc0, 0x15, 0x09, 0x21, 0x71, 0xe3, 0x82, 0x38, 0xf0, 0x15, 0xf8, 0x1c, 0xf0, 0x49, 0xb8,
	0x20, 0x3f, 0x8f, 0x67, 0x32, 0x69, 0xe1, 0x48, 0x7a, 0x89, 0xc6, 0xbf, 0xf7, 0x7b, 0xcf, 0xef,
	0x8f, 0xdf, 0xb3, 0x03, 0x83, 0xa9, 0x2e, 0xf3, 0x84, 0xea, 0x82, 0x8a, 0xbc, 0x98, 0x1b, 0x2a,
	0xf2, 0xa9, 0x1c, 0x15, 0x4a, 0x1a, 0x49, 0xb2, 0x44, 0xe8, 0x44, 0x52, 0x21, 0x35, 0xfd, 0x5a,
	0xd1, 0xa9, 0xe2, 0x67, 0x48, 0x95, 0x05, 0x57, 0x23, 0xbb, 0x9a, 0xf3, 0x3c, 0x29, 0xa9, 0xc5,
	0x1e, 0x2b, 0x99, 0x8b, 0x27, 0xcc, 0x08, 0x99, 0x8f, 0x72, 0x99, 0x72, 0x8d, 0xbf, 0x23, 0xcd,
	0x33, 0x9e, 0x58, 0x90, 0x16, 0x52, 0xe4, 0xc6, 0xed, 0xa0, 0x2f, 0x87, 0x87, 0x7f, 0x74, 0x20,
	0xba, 0xc4, 0x17, 0xfa, 0xe9, 0xc7, 0x5f, 0x9e, 0x10, 0x02, 0xeb, 0xd6, 0x62, 0x14, 0x1c, 0x06,
	0x47, 0xbd, 0x18, 0xbf, 0xc9, 0xab, 0xb0, 0xb3, 0x64, 0x29, 0xea, 0x1c, 0x06, 0x47, 0xdb, 0x71,
	0xbf, 0x86, 0x8f, 0x2d, 0x4a, 0x5e, 0x82, 0x4d, 0x6d, 0x14, 0x67, 0x33, 0x6a, 0xca, 0x82, 0x47,
	0x6b, 0x68, 0x03, 0x1c, 0xf4, 0xb0, 0x2c, 0x38, 0x12, 0xe4, 0x5c, 0x25, 0xdc, 0x11, 0xd6, 0x2b,
	0x02, 0x42, 0x48, 0x78, 0x01, 0x7a, 0x22, 0x37, 0x5c, 0x4d, 0x59, 0xc2, 0xa3, 0x67, 0x50, 0xdc,
	0x00, 0xe4, 0x45, 0x80, 0x24, 0x93, 0xc9, 0x57, 0xb4, 0x90, 0xca, 0x44, 0x5d, 0xf4, 0xa1, 0x87,
	0xc8, 0xb1, 0x54, 0x86, 0x1c, 0x40, 0x2f, 0x63, 0xda, 0x50, 0x0c, 0x60, 0x03, 0x95, 0x43, 0x0b,
	0xdc, 0xb7, 0x41, 0xbc, 0x01, 0x7b, 0x28, 0x5c, 0x8e, 0x24, 0x44, 0x2b, 0xc4, 0xca, 0x4e, 0xda,
	0xd1, 0x1c, 0x40, 0x4f, 0xce, 0x0d, 0xa6, 0x27, 0x8d, 0x7a, 0x48, 0x0b, 0x1d, 0x30, 0x49, 0x87,
	0x3f, 0x04, 0xb0, 0xe3, 0x92, 0xf8, 0x88, 0x9d, 0x62, 0x22, 0x53, 0xbb, 0xc5, 0x72, 0xc6, 0x31,
	0xcc, 0xc0, 0x6d, 0xd1, 0x4e, 0x16, 0x86, 0xfb, 0x21, 0x1c, 0x2c, 0x6b, 0xa4, 0x5c, 0x27, 0x4a,
	0x14, 0x16, 0xc1, 0x2c, 0xf7, 0xe2, 0x41, 0x5b, 0xf1, 0x6e, 0x43, 0xa8, 0xab, 0xb5, 0xd6, 0x54,
	0x6b, 0xf8, 0x05, 0x90, 0xc6, 0x31, 0x97, 0x2e, 0x91, 0x5e, 0x5a, 0xd7, 0x3e, 0x74, 0x44, 0x5a,
	0x95, 0xb2, 0x23, 0xd2, 0x26, 0xbd, 0x39, 0x9b, 0x79, 0x9b, 0x2e, 0xbd, 0xf7, 0xd9, 0x8c, 0x0f,
	0x9f, 0x76, 0xe1, 0xfa, 0x42, 0xc8, 0xae, 0x8e, 0x22, 0x25, 0x2f, 0xc3, 0x56, 0xb5, 0x48, 0x32,
	0xa6, 0x75, 0xb5, 0x45, 0x55, 0xe8, 0x3b, 0x16, 0x22, 0xaf, 0x01, 0xe1, 0xe6, 0x31, 0x57, 0x39,
	0xb7, 0x67, 0xcd, 0xd7, 0xd7, 0x85, 0xb7, 0xeb, 0x25, 0x93, 0xba, 0xce, 0xf6, 0xc0, 0xc9, 0x36,
	0xd7, 0x79, 0xd3, 0x47, 0xb8, 0x21, 0xfe, 0x14, 0x40, 0xe8, 0x43, 0xc4, 0xd3, 0xb4, 0x39, 0xfe,
	0x36, 0x18, 0xfd, 0x9f, 0xdd, 0x34, 0xba, 0x98, 0xeb, 0x78, 0x03, 0xbf, 0x26, 0x29, 0xf9, 0x2d,
	0x80, 0x5d, 0x0c, 0x21, 0x67, 0x59, 0x2d, 0xc6, 0x63, 0x7d, 0x25, 0xfc, 0xdc, 0xf1, 0xbe, 0xdd,
	0xa9, 0xfc, 0x1d, 0x40, 0x58, 0x98, 0xc2, 0xf5, 0x4f, 0x17, 0x13, 0xbe, 0x51, 0x98, 0x02, 0xdb,
	0xe7, 0x36, 0xec, 0x6b, 0x66, 0x78, 0x96, 0x09, 0xc3, 0x29, 0x4b, 0x12, 0xae, 0xf5, 0x42, 0x75,
	0x5c, 0xb3, 0x45, 0x35, 0xe3, 0x23, 0x24, 0x34, 0x75, 0x1a, 0x40, 0x98, 0x7b, 0xc3, 0xa1, 0x33,
	0x9c, 0x57, 0x86, 0x7f, 0x0d, 0xe0, 0xda, 0x69, 0xae, 0x35, 0x55, 0x3c, 0xe1, 0xe2, 0x9c, 0x2b,
	0xdf, 0x6d, 0x57, 0x22, 0x45, 0x7d, 0xeb, 0x5a, 0x5c, 0x79, 0x36, 0x49, 0x87, 0x7f, 0x06, 0xb0,
	0xd7, 0xd0, 0xdc, 0x40, 0xc1, 0xde, 0xff, 0x25, 0xb8, 0x38, 0x24, 0x03, 0x8c, 0xe2, 0x9b, 0x55,
	0x05, 0x81, 0x8e, 0x5d, 0x98, 0xd1, 0xad, 0xa9, 0xd6, 0x59, 0x9a, 0x6a, 0xdf, 0xaf, 0xb5, 0x5a,
	0xdc, 0xcd, 0xf2, 0xaa, 0xc5, 0xab, 0x85, 0xdd, 0xa4, 0x6e, 0x71, 0xc4, 0x26, 0x16, 0x22, 0x3f,
	0x07, 0xd0, 0xab, 0x67, 0x02, 0x1a, 0xde, 0x1c, 0x3f, 0x5d, 0x59, 0x01, 0x6b, 0x4f, 0xe2, 0xd0,
	0x7d, 0x4e, 0x52, 0xf2, 0x7b, 0x00, 0xe4, 0x82, 0x5e, 0x8a, 0x93, 0x65, 0x73, 0xfc, 0xdd, 0xca,
	0x5c, 0x6d, 0x0e, 0x51, 0x7c, 0xad, 0x5d, 0xab, 0x49, 0x3a, 0xfc, 0x2b, 0x80, 0xad, 0x86, 0x7a,
	0x96, 0xd9, 0x3b, 0xe6, 0x6c, 0xce, 0x32, 0x61, 0x4a, 0x9a, 0xf1, 0x73, 0x9e, 0x51, 0xe9, 0xae,
	0x0a, 0x57, 0x11, 0x52, 0xc9, 0x3e, 0xb3, 0xa2, 0x07, 0xee, 0x8e, 0x78, 0x05, 0xb6, 0x1d, 0xe7,
	0x4d, 0x7a, 0xce, 0xb2, 0xb9, 0x1f, 0xbb, 0x5b, 0x15, 0xf8, 0xb9, 0xc5, 0xc8, 0x2d, 0x18, 0xb8,
	0xf5, 0x98, 0x9e, 0xf2, 0x9c, 0x2b, 0xb6, 0xa8, 0xe0, 0x66, 0xef, 0xf3, 0x15, 0xe1, 0x93, 0x46,
	0xfe, 0x1f, 0xba, 0xe3, 0x4a, 0x77, 0xfd, 0x5f, 0x74, 0xc7, 0xa8, 0x3b, 0xfc, 0x7b, 0xc3, 0x1f,
	0xb8, 0xd6, 0x5b, 0xc4, 0xd6, 0xea, 0x59, 0xb7, 0x5c, 0x6e, 0xaa, 0xf1, 0x55, 0x68, 0xaa, 0xeb,
	0x08, 0x2f, 0xbd, 0x17, 0x7e, 0x0c, 0xa0, 0xeb, 0x3a, 0x22, 0xba, 0xb1, 0xea, 0xe3, 0xef, 0x3b,
	0x37, 0xae, 0x3c, 0xb2, 0x63, 0x76, 0x47, 0x2a, 0x71, 0x2a, 0xec, 0x55, 0xe4, 0x5a, 0x22, 0xba,
	0x79, 0x65, 0x9a, 0xb4, 0xef, 0x5d, 0x3b, 0x41, 0xc8, 0xbe, 0x17, 0xf4, 0xbc, 0xb0, 0xaf, 0x3c,
	0x4d, 0x6b, 0x0f, 0xa2, 0xb7, 0x0e, 0x83, 0xa3, 0x30, 0xde, 0xf5, 0x92, 0x7b, 0x5e, 0x40, 0x5e,
	0x87, 0xbd, 0x9a, 0x6e, 0xc4, 0x8c, 0x53, 0x39, 0xa5, 0x29, 0x2b, 0xa3, 0xb7, 0xdb, 0x0a, 0x0f,
	0xc5, 0x8c, 0x3f, 0x98, 0xde, 0x65, 0xa5, 0x1d, 0x56, 0xdb, 0xad, 0x36, 0x8a, 0xde, 0xc1, 0x5c,
	0x3c, 0x59, 0x55, 0x2a, 0xce, 0xb2, 0x78, 0x6b, 0xb1, 0x77, 0xc9, 0x3e, 0x84, 0x85, 0x12, 0x52,
	0x09, 0x53, 0x46, 0xef, 0xba, 0x29, 0xed, 0xd7, 0x36, 0xdc, 0x85, 0x28, 0x69, 0xcd, 0x7b, 0x0f,
	0x79, 0xbb, 0xc6, 0x87, 0x79, 0xec, 0x15, 0xf6, 0x21, 0x74, 0x0e, 0xf0, 0x34, 0xba, 0x85, 0x39,
	0xa9, 0xd7, 0xed, 0xfb, 0xe0, 0xfd, 0xf6, 0x7d, 0x60, 0x1f, 0x62, 0x45, 0xc6, 0xcc, 0x54, 0xaa,
	0x19, 0xd5, 0x86, 0x99, 0xb9, 0x8e, 0x6e, 0xbb, 0x87, 0x98, 0x87, 0x4f, 0x10, 0x25, 0x37, 0xe1,
	0xb9, 0x9a, 0x38, 0x65, 0x22, 0xe3, 0x29, 0x55, 0x9c, 0x69, 0x99, 0x47, 0x1f, 0x20, 0x7f, 0xcf,
	0x4b, 0xef, 0xa1, 0x30, 0x46, 0xd9, 0xa3, 0x2e, 0xfe, 0xfd, 0xb9, 0xf1, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x59, 0x53, 0x34, 0x1b, 0x0d, 0x00, 0x00,
}