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
// source: fsync_selection_backtrace_info.proto

package cisco_ios_xr_freqsync_oper_frequency_synchronization_global_interfaces_global_interface_interface_selection_back_trace

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

type FsyncSelectionBacktraceInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncSelectionBacktraceInfo_KEYS) Reset()         { *m = FsyncSelectionBacktraceInfo_KEYS{} }
func (m *FsyncSelectionBacktraceInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncSelectionBacktraceInfo_KEYS) ProtoMessage()    {}
func (*FsyncSelectionBacktraceInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_365772aa9c279f29, []int{0}
}

func (m *FsyncSelectionBacktraceInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS.Unmarshal(m, b)
}
func (m *FsyncSelectionBacktraceInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncSelectionBacktraceInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS.Merge(m, src)
}
func (m *FsyncSelectionBacktraceInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS.Size(m)
}
func (m *FsyncSelectionBacktraceInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSelectionBacktraceInfo_KEYS proto.InternalMessageInfo

func (m *FsyncSelectionBacktraceInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_365772aa9c279f29, []int{1}
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
	return fileDescriptor_365772aa9c279f29, []int{2}
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
	return fileDescriptor_365772aa9c279f29, []int{3}
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

type FsyncSelectionBacktraceInfo struct {
	SelectedSource       *FsyncBagSourceId `protobuf:"bytes,50,opt,name=selected_source,json=selectedSource,proto3" json:"selected_source,omitempty"`
	SelectionPoint       []*FsyncBagSpId   `protobuf:"bytes,51,rep,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncSelectionBacktraceInfo) Reset()         { *m = FsyncSelectionBacktraceInfo{} }
func (m *FsyncSelectionBacktraceInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncSelectionBacktraceInfo) ProtoMessage()    {}
func (*FsyncSelectionBacktraceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_365772aa9c279f29, []int{4}
}

func (m *FsyncSelectionBacktraceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo.Unmarshal(m, b)
}
func (m *FsyncSelectionBacktraceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo.Marshal(b, m, deterministic)
}
func (m *FsyncSelectionBacktraceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSelectionBacktraceInfo.Merge(m, src)
}
func (m *FsyncSelectionBacktraceInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncSelectionBacktraceInfo.Size(m)
}
func (m *FsyncSelectionBacktraceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSelectionBacktraceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSelectionBacktraceInfo proto.InternalMessageInfo

func (m *FsyncSelectionBacktraceInfo) GetSelectedSource() *FsyncBagSourceId {
	if m != nil {
		return m.SelectedSource
	}
	return nil
}

func (m *FsyncSelectionBacktraceInfo) GetSelectionPoint() []*FsyncBagSpId {
	if m != nil {
		return m.SelectionPoint
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncSelectionBacktraceInfo_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_back_trace.fsync_selection_backtrace_info_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_back_trace.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_back_trace.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_back_trace.fsync_bag_sp_id")
	proto.RegisterType((*FsyncSelectionBacktraceInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_back_trace.fsync_selection_backtrace_info")
}

func init() {
	proto.RegisterFile("fsync_selection_backtrace_info.proto", fileDescriptor_365772aa9c279f29)
}

var fileDescriptor_365772aa9c279f29 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0xa5, 0x4d, 0xa6, 0x24, 0xa1, 0x86, 0xc3, 0x06, 0x04, 0x0a, 0x01, 0x44, 0x2e,
	0xac, 0x50, 0x7a, 0x45, 0x48, 0xa8, 0x70, 0x88, 0x40, 0x15, 0xda, 0x22, 0x21, 0x4e, 0x96, 0xe3,
	0x9d, 0xa4, 0x56, 0xb7, 0xb6, 0xb1, 0xdd, 0x8a, 0xf0, 0x07, 0x90, 0x80, 0x3b, 0x47, 0x6e, 0xdc,
	0x39, 0x20, 0xfe, 0x1e, 0xb2, 0xf7, 0x2b, 0x6d, 0xa5, 0x5e, 0xc9, 0xcd, 0xfb, 0xe6, 0xcd, 0xdb,
	0x97, 0x37, 0x3b, 0x0e, 0x3c, 0x9c, 0xdb, 0xa5, 0xe4, 0xd4, 0x62, 0x8e, 0xdc, 0x09, 0x25, 0xe9,
	0x8c, 0xf1, 0x63, 0x67, 0x18, 0x47, 0x2a, 0xe4, 0x5c, 0x25, 0xda, 0x28, 0xa7, 0xc8, 0x19, 0x17,
	0x96, 0x2b, 0x2a, 0x94, 0xa5, 0x9f, 0x0c, 0x9d, 0x1b, 0xfc, 0x18, 0xba, 0x94, 0x46, 0x93, 0xf8,
	0xa7, 0x53, 0x94, 0x7c, 0x49, 0x3d, 0x76, 0x64, 0x94, 0x14, 0x9f, 0x99, 0x97, 0x4a, 0x16, 0xb9,
	0x9a, 0xb1, 0x9c, 0x0a, 0xe9, 0xd0, 0xcc, 0x19, 0x47, 0x7b, 0x09, 0x49, 0xea, 0xd3, 0x05, 0x07,
	0x34, 0x58, 0x18, 0xbd, 0x81, 0x07, 0x57, 0xfb, 0xa3, 0xaf, 0x5f, 0x7d, 0x38, 0x24, 0x8f, 0xa0,
	0xd7, 0x08, 0x49, 0x76, 0x82, 0x71, 0x34, 0x8c, 0xc6, 0x9d, 0xb4, 0x5b, 0xa3, 0x07, 0xec, 0x04,
	0x47, 0xef, 0x81, 0x14, 0x6a, 0x33, 0xb6, 0xa0, 0x3c, 0x57, 0xfc, 0x98, 0x8a, 0x8c, 0x10, 0xd8,
	0x94, 0x2a, 0xab, 0x5a, 0xc2, 0x99, 0xf4, 0x60, 0x43, 0x64, 0xf1, 0xc6, 0x30, 0x1a, 0x77, 0xd3,
	0x0d, 0x91, 0x91, 0xbb, 0x00, 0x05, 0x3f, 0x88, 0xb7, 0x02, 0xb3, 0x13, 0x90, 0x20, 0xfc, 0x73,
	0x0b, 0x6e, 0x36, 0xca, 0x56, 0x9d, 0x1a, 0x6f, 0x2f, 0x23, 0xf7, 0xe1, 0x7a, 0xf9, 0xc0, 0x73,
	0x66, 0x6d, 0xf9, 0x8a, 0x9d, 0x02, 0xdb, 0xf7, 0x10, 0x79, 0x02, 0x04, 0xdd, 0x11, 0x1a, 0x89,
	0xae, 0x89, 0x25, 0xbc, 0xb9, 0x93, 0xee, 0x56, 0x95, 0x69, 0x55, 0x20, 0x8f, 0xa1, 0x6f, 0xd5,
	0x79, 0x6e, 0xe1, 0xa6, 0x17, 0xe0, 0x86, 0xf8, 0x2b, 0x82, 0x76, 0xf5, 0x13, 0xe3, 0xcd, 0x61,
	0x34, 0xde, 0x99, 0x7c, 0x8d, 0x92, 0xff, 0x33, 0xc6, 0xe4, 0x72, 0xea, 0xe9, 0x76, 0x38, 0x4d,
	0x33, 0xf2, 0x37, 0x82, 0xdd, 0xd0, 0x2e, 0x59, 0x5e, 0x97, 0xe3, 0x6b, 0xeb, 0xe7, 0xb8, 0x5f,
	0xb9, 0xdc, 0x2f, 0x9d, 0x0f, 0xa0, 0xad, 0x9d, 0xa6, 0xe1, 0xe3, 0xd9, 0x0a, 0x43, 0xd8, 0xd6,
	0x4e, 0x1f, 0xf8, 0xef, 0xe7, 0x19, 0xdc, 0xb6, 0xcc, 0x61, 0x9e, 0x0b, 0x87, 0x94, 0x71, 0x8e,
	0xd6, 0xae, 0x4c, 0x6c, 0x3b, 0x90, 0xe3, 0x9a, 0xf1, 0x22, 0x10, 0x9a, 0xd9, 0x0d, 0xa0, 0x2d,
	0x2b, 0xe1, 0x76, 0x21, 0x2c, 0x4b, 0xe1, 0x3f, 0x11, 0xdc, 0x58, 0x48, 0x6b, 0xa9, 0x41, 0x8e,
	0xe2, 0x0c, 0x8d, 0x0f, 0xab, 0xb3, 0x7e, 0x61, 0xf5, 0xbc, 0xc9, 0xb4, 0xf4, 0x38, 0xcd, 0x46,
	0x3f, 0x22, 0xe8, 0xaf, 0x6c, 0x88, 0xf6, 0xdb, 0xf1, 0x14, 0x6e, 0x35, 0x72, 0x5a, 0x09, 0xe9,
	0xa8, 0x5b, 0xea, 0x62, 0x11, 0xbb, 0x29, 0xa9, 0x6b, 0x6f, 0x7d, 0xe9, 0xdd, 0x52, 0x23, 0x79,
	0x0e, 0x77, 0x2e, 0x76, 0x64, 0x68, 0xb9, 0x11, 0xda, 0x23, 0xe5, 0xd6, 0x0c, 0xce, 0x37, 0xbe,
	0x6c, 0x08, 0xf5, 0xaa, 0xb7, 0x9a, 0x55, 0x1f, 0x7d, 0x6f, 0xc1, 0xbd, 0xab, 0xef, 0x18, 0x1f,
	0x7a, 0xbf, 0x28, 0x62, 0x56, 0x6e, 0x77, 0x3c, 0x09, 0x99, 0x7f, 0x5b, 0x83, 0xcc, 0xeb, 0xeb,
	0x26, 0xed, 0x55, 0x26, 0x0f, 0x03, 0x44, 0x7e, 0xd7, 0xbe, 0xeb, 0xbc, 0xe2, 0xbd, 0x61, 0x6b,
	0xbc, 0x33, 0xf9, 0xb2, 0x0e, 0xbe, 0xf5, 0x8a, 0xe7, 0x6a, 0x5a, 0xb3, 0xad, 0xf0, 0x87, 0xb3,
	0xf7, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd3, 0x64, 0xd3, 0x98, 0x06, 0x00, 0x00,
}
