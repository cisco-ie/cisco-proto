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
// source: fsync_config_errors_info.proto

package cisco_ios_xr_ncs4k_freqsync_oper_frequency_synchronization_nodes_node_configuration_errors

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

type FsyncConfigErrorsInfo_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncConfigErrorsInfo_KEYS) Reset()         { *m = FsyncConfigErrorsInfo_KEYS{} }
func (m *FsyncConfigErrorsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncConfigErrorsInfo_KEYS) ProtoMessage()    {}
func (*FsyncConfigErrorsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5789f570eefdec, []int{0}
}

func (m *FsyncConfigErrorsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncConfigErrorsInfo_KEYS.Unmarshal(m, b)
}
func (m *FsyncConfigErrorsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncConfigErrorsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncConfigErrorsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncConfigErrorsInfo_KEYS.Merge(m, src)
}
func (m *FsyncConfigErrorsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncConfigErrorsInfo_KEYS.Size(m)
}
func (m *FsyncConfigErrorsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncConfigErrorsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncConfigErrorsInfo_KEYS proto.InternalMessageInfo

func (m *FsyncConfigErrorsInfo_KEYS) GetNode() string {
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
	return fileDescriptor_7d5789f570eefdec, []int{1}
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
	Node                     string           `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
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
	return fileDescriptor_7d5789f570eefdec, []int{2}
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

func (m *FsyncBagSourceId) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
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
	return fileDescriptor_7d5789f570eefdec, []int{3}
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

type FsyncBagError struct {
	Source               *FsyncBagSourceId `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	EnableError          bool              `protobuf:"varint,2,opt,name=enable_error,json=enableError,proto3" json:"enable_error,omitempty"`
	InputMinError        bool              `protobuf:"varint,3,opt,name=input_min_error,json=inputMinError,proto3" json:"input_min_error,omitempty"`
	InputExactError      bool              `protobuf:"varint,4,opt,name=input_exact_error,json=inputExactError,proto3" json:"input_exact_error,omitempty"`
	InputMaxError        bool              `protobuf:"varint,5,opt,name=input_max_error,json=inputMaxError,proto3" json:"input_max_error,omitempty"`
	OuputMinError        bool              `protobuf:"varint,6,opt,name=ouput_min_error,json=ouputMinError,proto3" json:"ouput_min_error,omitempty"`
	OutputExactError     bool              `protobuf:"varint,7,opt,name=output_exact_error,json=outputExactError,proto3" json:"output_exact_error,omitempty"`
	OutputMaxError       bool              `protobuf:"varint,8,opt,name=output_max_error,json=outputMaxError,proto3" json:"output_max_error,omitempty"`
	InputOutputMismatch  bool              `protobuf:"varint,9,opt,name=input_output_mismatch,json=inputOutputMismatch,proto3" json:"input_output_mismatch,omitempty"`
	InputMinQl           *FsyncBagQl       `protobuf:"bytes,10,opt,name=input_min_ql,json=inputMinQl,proto3" json:"input_min_ql,omitempty"`
	InputExactQl         *FsyncBagQl       `protobuf:"bytes,11,opt,name=input_exact_ql,json=inputExactQl,proto3" json:"input_exact_ql,omitempty"`
	InputMaxQl           *FsyncBagQl       `protobuf:"bytes,12,opt,name=input_max_ql,json=inputMaxQl,proto3" json:"input_max_ql,omitempty"`
	OutputMinQl          *FsyncBagQl       `protobuf:"bytes,13,opt,name=output_min_ql,json=outputMinQl,proto3" json:"output_min_ql,omitempty"`
	OutputExactQl        *FsyncBagQl       `protobuf:"bytes,14,opt,name=output_exact_ql,json=outputExactQl,proto3" json:"output_exact_ql,omitempty"`
	OutputMaxQl          *FsyncBagQl       `protobuf:"bytes,15,opt,name=output_max_ql,json=outputMaxQl,proto3" json:"output_max_ql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncBagError) Reset()         { *m = FsyncBagError{} }
func (m *FsyncBagError) String() string { return proto.CompactTextString(m) }
func (*FsyncBagError) ProtoMessage()    {}
func (*FsyncBagError) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5789f570eefdec, []int{4}
}

func (m *FsyncBagError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagError.Unmarshal(m, b)
}
func (m *FsyncBagError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagError.Marshal(b, m, deterministic)
}
func (m *FsyncBagError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagError.Merge(m, src)
}
func (m *FsyncBagError) XXX_Size() int {
	return xxx_messageInfo_FsyncBagError.Size(m)
}
func (m *FsyncBagError) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagError.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagError proto.InternalMessageInfo

func (m *FsyncBagError) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FsyncBagError) GetEnableError() bool {
	if m != nil {
		return m.EnableError
	}
	return false
}

func (m *FsyncBagError) GetInputMinError() bool {
	if m != nil {
		return m.InputMinError
	}
	return false
}

func (m *FsyncBagError) GetInputExactError() bool {
	if m != nil {
		return m.InputExactError
	}
	return false
}

func (m *FsyncBagError) GetInputMaxError() bool {
	if m != nil {
		return m.InputMaxError
	}
	return false
}

func (m *FsyncBagError) GetOuputMinError() bool {
	if m != nil {
		return m.OuputMinError
	}
	return false
}

func (m *FsyncBagError) GetOutputExactError() bool {
	if m != nil {
		return m.OutputExactError
	}
	return false
}

func (m *FsyncBagError) GetOutputMaxError() bool {
	if m != nil {
		return m.OutputMaxError
	}
	return false
}

func (m *FsyncBagError) GetInputOutputMismatch() bool {
	if m != nil {
		return m.InputOutputMismatch
	}
	return false
}

func (m *FsyncBagError) GetInputMinQl() *FsyncBagQl {
	if m != nil {
		return m.InputMinQl
	}
	return nil
}

func (m *FsyncBagError) GetInputExactQl() *FsyncBagQl {
	if m != nil {
		return m.InputExactQl
	}
	return nil
}

func (m *FsyncBagError) GetInputMaxQl() *FsyncBagQl {
	if m != nil {
		return m.InputMaxQl
	}
	return nil
}

func (m *FsyncBagError) GetOutputMinQl() *FsyncBagQl {
	if m != nil {
		return m.OutputMinQl
	}
	return nil
}

func (m *FsyncBagError) GetOutputExactQl() *FsyncBagQl {
	if m != nil {
		return m.OutputExactQl
	}
	return nil
}

func (m *FsyncBagError) GetOutputMaxQl() *FsyncBagQl {
	if m != nil {
		return m.OutputMaxQl
	}
	return nil
}

type FsyncConfigErrorsInfo struct {
	ErrorSource          []*FsyncBagError `protobuf:"bytes,50,rep,name=error_source,json=errorSource,proto3" json:"error_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FsyncConfigErrorsInfo) Reset()         { *m = FsyncConfigErrorsInfo{} }
func (m *FsyncConfigErrorsInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncConfigErrorsInfo) ProtoMessage()    {}
func (*FsyncConfigErrorsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d5789f570eefdec, []int{5}
}

func (m *FsyncConfigErrorsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncConfigErrorsInfo.Unmarshal(m, b)
}
func (m *FsyncConfigErrorsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncConfigErrorsInfo.Marshal(b, m, deterministic)
}
func (m *FsyncConfigErrorsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncConfigErrorsInfo.Merge(m, src)
}
func (m *FsyncConfigErrorsInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncConfigErrorsInfo.Size(m)
}
func (m *FsyncConfigErrorsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncConfigErrorsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncConfigErrorsInfo proto.InternalMessageInfo

func (m *FsyncConfigErrorsInfo) GetErrorSource() []*FsyncBagError {
	if m != nil {
		return m.ErrorSource
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncConfigErrorsInfo_KEYS)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_config_errors_info_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagQl)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_bag_ql")
	proto.RegisterType((*FsyncBagError)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_bag_error")
	proto.RegisterType((*FsyncConfigErrorsInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.nodes.node.configuration_errors.fsync_config_errors_info")
}

func init() { proto.RegisterFile("fsync_config_errors_info.proto", fileDescriptor_7d5789f570eefdec) }

var fileDescriptor_7d5789f570eefdec = []byte{
	// 765 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0xd6, 0x74, 0xb3, 0xed, 0xf4, 0x4c, 0x7e, 0x76, 0xbd, 0x20, 0xa6, 0x48, 0x8b, 0x42, 0x90,
	0x20, 0x42, 0x10, 0xc1, 0x2c, 0x57, 0x88, 0x1b, 0x84, 0x2a, 0x54, 0x01, 0xbb, 0xea, 0xac, 0x04,
	0x82, 0x1b, 0xcb, 0x75, 0x9c, 0xc4, 0xaa, 0x63, 0x4f, 0xc6, 0x33, 0x55, 0xca, 0x0b, 0x20, 0x10,
	0x42, 0x5c, 0xf2, 0x0e, 0x3c, 0x04, 0xf7, 0xbc, 0x03, 0xef, 0x82, 0x7c, 0xec, 0xc9, 0x4c, 0x0b,
	0xe5, 0x8e, 0xf4, 0x26, 0x19, 0x7f, 0xe7, 0x3b, 0x3f, 0xfe, 0x3e, 0x7b, 0x34, 0xf0, 0xc6, 0xc2,
	0x5e, 0x6b, 0x4e, 0xb9, 0xd1, 0x0b, 0xb9, 0xa4, 0xa2, 0x2c, 0x4d, 0x69, 0xa9, 0xd4, 0x0b, 0x33,
	0x2b, 0x4a, 0x53, 0x19, 0xf2, 0x1d, 0x97, 0x96, 0x1b, 0x2a, 0x8d, 0xa5, 0xdb, 0x92, 0x6a, 0x6e,
	0x3f, 0xba, 0xa4, 0x8b, 0x52, 0x6c, 0x30, 0xcb, 0x14, 0xa2, 0x9c, 0xb9, 0x55, 0x2d, 0x34, 0xbf,
	0xa6, 0x0e, 0x5b, 0x95, 0x46, 0xcb, 0xef, 0x59, 0x25, 0x8d, 0x9e, 0x69, 0x33, 0x17, 0x16, 0x7f,
	0x67, 0xbe, 0x7e, 0x5d, 0x62, 0x20, 0xb4, 0x99, 0x3c, 0x83, 0xa7, 0x77, 0x75, 0xa7, 0x5f, 0x9c,
	0x7e, 0xfb, 0x92, 0x10, 0xe8, 0xb9, 0xec, 0x34, 0x1a, 0x47, 0xd3, 0xe3, 0x1c, 0x9f, 0x27, 0xdf,
	0x00, 0xf1, 0x49, 0x17, 0x6c, 0x49, 0xb9, 0x32, 0xfc, 0x92, 0xca, 0xf9, 0xbf, 0x31, 0xc9, 0x10,
	0x0e, 0xe4, 0x3c, 0x3d, 0x18, 0x47, 0xd3, 0x41, 0x7e, 0x20, 0xe7, 0xe4, 0x29, 0x80, 0xe7, 0x6b,
	0xb6, 0x16, 0xe9, 0x03, 0x64, 0x1e, 0x23, 0xf2, 0x9c, 0xad, 0xc5, 0xe4, 0x8f, 0x1e, 0x3c, 0x69,
	0x2b, 0x5b, 0x53, 0x97, 0x5c, 0xb8, 0xd2, 0x6f, 0x42, 0x3f, 0x2c, 0xb8, 0x62, 0xd6, 0x86, 0x16,
	0x89, 0xc7, 0x3e, 0x73, 0x10, 0x79, 0x1f, 0x88, 0xa8, 0x56, 0xa2, 0xd4, 0xa2, 0xa2, 0x52, 0x57,
	0xa2, 0x5c, 0x30, 0x2e, 0xb0, 0xf3, 0x71, 0xfe, 0xb8, 0x89, 0x9c, 0x35, 0x01, 0xf2, 0x0e, 0x8c,
	0xac, 0xb9, 0xc9, 0xf5, 0xd3, 0x0c, 0x11, 0x6e, 0x89, 0x3f, 0x46, 0x10, 0x37, 0x5b, 0x4c, 0x7b,
	0xe3, 0x68, 0x9a, 0x64, 0x7a, 0xf6, 0xff, 0x19, 0x32, 0xfb, 0xa7, 0xb0, 0xf9, 0x11, 0x3e, 0x9d,
	0xb5, 0x0a, 0x3f, 0xec, 0x28, 0x7c, 0x02, 0x71, 0x51, 0x15, 0x14, 0xf1, 0x43, 0xc4, 0x8f, 0x8a,
	0xaa, 0x78, 0xee, 0x42, 0x9f, 0xc0, 0xeb, 0x96, 0x55, 0x42, 0x29, 0x59, 0x09, 0xca, 0x38, 0x17,
	0xd6, 0x76, 0xb6, 0x7b, 0x84, 0xe4, 0x74, 0xc7, 0xf8, 0x14, 0x09, 0xed, 0xc6, 0x4f, 0x20, 0xd6,
	0x4d, 0xe1, 0xd8, 0x17, 0xd6, 0xa1, 0xf0, 0x6f, 0x11, 0x3c, 0x5a, 0x6a, 0x6b, 0x69, 0x29, 0xb8,
	0x90, 0x57, 0xa2, 0x74, 0xda, 0x1c, 0xdf, 0x8b, 0x36, 0x43, 0x37, 0x47, 0x1e, 0xc6, 0x38, 0x9b,
	0x4f, 0xfe, 0x8a, 0xa0, 0xdf, 0xd2, 0x36, 0x8a, 0x7c, 0x00, 0xaf, 0x6c, 0x6a, 0xa6, 0x64, 0x75,
	0x4d, 0x95, 0xb8, 0x12, 0x8a, 0x9a, 0xc2, 0x95, 0x0c, 0x47, 0x88, 0x84, 0xd8, 0x97, 0x2e, 0xf4,
	0x02, 0x23, 0xe4, 0x2d, 0x18, 0x78, 0xce, 0x87, 0xf4, 0x8a, 0xa9, 0xba, 0x39, 0x44, 0xfd, 0x00,
	0x7e, 0xed, 0x30, 0xf2, 0x31, 0x9c, 0xf8, 0x75, 0x46, 0x97, 0x42, 0x0b, 0x3f, 0x67, 0x93, 0xe0,
	0x4f, 0xd2, 0x6b, 0x81, 0xf0, 0x79, 0x1b, 0xff, 0x8f, 0xdc, 0x2c, 0xe4, 0xf6, 0xee, 0xc8, 0xcd,
	0x30, 0x77, 0xf2, 0x27, 0xc0, 0xa8, 0xdd, 0x1f, 0xea, 0x42, 0x7e, 0x88, 0xe0, 0xd0, 0x5f, 0x05,
	0xdc, 0x55, 0x92, 0x99, 0xfd, 0x98, 0xb0, 0xbb, 0x9f, 0x79, 0x68, 0xef, 0xee, 0xa9, 0xd0, 0xec,
	0x42, 0x09, 0x4f, 0x46, 0xe5, 0xe2, 0x3c, 0xf1, 0xd8, 0x29, 0x0e, 0xfb, 0x36, 0x8c, 0xa4, 0x2e,
	0xea, 0x8a, 0xae, 0x65, 0x28, 0x89, 0x72, 0xc5, 0xf9, 0x00, 0xe1, 0xaf, 0xa4, 0xf6, 0xbc, 0x77,
	0xe1, 0xb1, 0xe7, 0x89, 0x2d, 0xe3, 0x55, 0x60, 0xf6, 0x90, 0xe9, 0x0b, 0x9c, 0x3a, 0xfc, 0x76,
	0x4d, 0xb6, 0x0d, 0xcc, 0x87, 0xdd, 0x9a, 0x6c, 0xbb, 0xe3, 0x99, 0xfa, 0x66, 0xef, 0x43, 0xcf,
	0x43, 0x78, 0xd7, 0xfb, 0x3d, 0x20, 0xa6, 0xae, 0x6e, 0x37, 0x3f, 0x42, 0xea, 0x23, 0x1f, 0xe9,
	0x74, 0x9f, 0x42, 0xc0, 0x3a, 0xed, 0x63, 0xe4, 0x0e, 0x3d, 0xbe, 0xeb, 0x9f, 0xc1, 0xab, 0x7e,
	0xce, 0x86, 0x2f, 0xed, 0x9a, 0x55, 0x7c, 0x85, 0x77, 0x27, 0xce, 0x9f, 0x60, 0xf0, 0x85, 0xcf,
	0x09, 0x21, 0xf2, 0x53, 0x04, 0xfd, 0x56, 0xb0, 0x8d, 0x4a, 0x01, 0x2d, 0x5e, 0xed, 0xc7, 0xe2,
	0x8d, 0xca, 0xa1, 0xf1, 0xe5, 0x5c, 0x91, 0x5f, 0x22, 0x18, 0x76, 0x5d, 0xd9, 0xa8, 0x34, 0xd9,
	0xf3, 0x38, 0xfd, 0xd6, 0xfc, 0x73, 0xd5, 0x55, 0x87, 0x6d, 0xdd, 0x38, 0xfd, 0xfb, 0x51, 0x87,
	0x6d, 0xcf, 0x15, 0xf9, 0x39, 0x82, 0xc1, 0xce, 0x59, 0xf4, 0x6a, 0xb0, 0xe7, 0x69, 0x92, 0x70,
	0xe0, 0xd0, 0xac, 0x5f, 0x23, 0x77, 0xdc, 0xab, 0x1b, 0x6e, 0x0d, 0xf7, 0x3c, 0xd0, 0xa0, 0x73,
	0x5b, 0x6e, 0x29, 0xe4, 0xfd, 0x1a, 0xdd, 0x93, 0x42, 0xce, 0xb0, 0xc9, 0xef, 0x11, 0xa4, 0x77,
	0x7d, 0xfd, 0xb8, 0xb3, 0xde, 0xc7, 0x75, 0x78, 0xcf, 0xa5, 0xd9, 0xf8, 0xc1, 0x34, 0xc9, 0x2e,
	0xf7, 0x33, 0x2a, 0x02, 0x79, 0x82, 0x7f, 0x2f, 0xb1, 0xff, 0xc5, 0x21, 0x7e, 0x0d, 0x3e, 0xfb,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb9, 0x99, 0x64, 0x2f, 0x0a, 0x00, 0x00,
}
