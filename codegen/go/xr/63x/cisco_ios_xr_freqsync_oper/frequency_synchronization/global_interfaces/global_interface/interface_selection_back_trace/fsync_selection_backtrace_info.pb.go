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
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
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

func (m *FsyncBagClockId) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
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
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x6f, 0x13, 0x3d,
	0x10, 0xd5, 0x7e, 0xdb, 0xaf, 0x49, 0x27, 0x24, 0x11, 0x86, 0xc3, 0x06, 0x24, 0x14, 0x16, 0x10,
	0xb9, 0xb0, 0x42, 0xe9, 0xb5, 0x42, 0x42, 0xc0, 0xa1, 0x02, 0x55, 0x68, 0xcb, 0x85, 0x93, 0xe5,
	0x78, 0x27, 0xad, 0xd5, 0xad, 0x6d, 0x6c, 0x17, 0x11, 0xfe, 0x00, 0x12, 0x70, 0xe7, 0x1f, 0x70,
	0xe7, 0xc0, 0x0f, 0xe0, 0x9f, 0x21, 0x7b, 0xb3, 0xeb, 0xb4, 0x95, 0x7a, 0xa5, 0xb7, 0xf1, 0x9b,
	0x99, 0xa7, 0xe7, 0xf7, 0xbc, 0x0b, 0x0f, 0x97, 0x76, 0x25, 0x39, 0xb5, 0x58, 0x23, 0x77, 0x42,
	0x49, 0xba, 0x60, 0xfc, 0xc4, 0x19, 0xc6, 0x91, 0x0a, 0xb9, 0x54, 0x85, 0x36, 0xca, 0x29, 0xf2,
	0x91, 0x0b, 0xcb, 0x15, 0x15, 0xca, 0xd2, 0x4f, 0x86, 0x2e, 0x0d, 0x7e, 0x08, 0x5b, 0x4a, 0xa3,
	0x29, 0xfc, 0xe9, 0x0c, 0x25, 0x5f, 0x51, 0x8f, 0x1d, 0x1b, 0x25, 0xc5, 0x67, 0xe6, 0xa9, 0x8a,
	0xa3, 0x5a, 0x2d, 0x58, 0x4d, 0x85, 0x74, 0x68, 0x96, 0x8c, 0xa3, 0xbd, 0x84, 0x14, 0x5d, 0x75,
	0x41, 0x01, 0x0d, 0x12, 0xf2, 0x37, 0xf0, 0xe0, 0x6a, 0x7d, 0xf4, 0xf5, 0xab, 0xf7, 0x87, 0xe4,
	0x11, 0x8c, 0x22, 0x91, 0x64, 0xa7, 0x98, 0x25, 0xd3, 0x64, 0xb6, 0x53, 0x0e, 0x3b, 0xf4, 0x80,
	0x9d, 0x62, 0xbe, 0x07, 0xa4, 0x61, 0x5b, 0xb0, 0x23, 0xca, 0x6b, 0xc5, 0x4f, 0xa8, 0xa8, 0x08,
	0x81, 0x2d, 0xa9, 0xaa, 0x76, 0x25, 0xd4, 0x1e, 0xd3, 0xca, 0xb8, 0xec, 0xbf, 0x69, 0x32, 0x1b,
	0x96, 0xa1, 0xce, 0xff, 0xa4, 0x70, 0x2b, 0xae, 0x5b, 0x75, 0x66, 0xbc, 0x86, 0x8a, 0xdc, 0x87,
	0x1b, 0xeb, 0x03, 0xaf, 0x99, 0xb5, 0x6b, 0x9e, 0x41, 0x83, 0xbd, 0xf0, 0x10, 0x79, 0x02, 0x04,
	0xdd, 0x31, 0x1a, 0x89, 0x2e, 0xde, 0x3d, 0x90, 0xef, 0x94, 0x37, 0xdb, 0xce, 0x7e, 0xdb, 0x20,
	0x8f, 0x61, 0x6c, 0xd5, 0xf9, 0xd9, 0x34, 0xcc, 0x8e, 0x02, 0x1c, 0x07, 0x7f, 0x26, 0xd0, 0x6f,
	0xef, 0x91, 0x6d, 0x4d, 0x93, 0xd9, 0x60, 0xfe, 0x35, 0x29, 0xfe, 0x4d, 0x56, 0xc5, 0x65, 0x6b,
	0xcb, 0x5e, 0xa8, 0xf6, 0xa3, 0xc7, 0xff, 0x6f, 0x78, 0x3c, 0x81, 0xbe, 0x76, 0x9a, 0x06, 0x7c,
	0x3b, 0xe0, 0x3d, 0xed, 0xf4, 0x81, 0x6f, 0xed, 0xc1, 0x1d, 0xcb, 0x1c, 0xd6, 0xb5, 0x70, 0x48,
	0x19, 0xe7, 0x68, 0xed, 0x86, 0x17, 0xbd, 0x30, 0x9c, 0x75, 0x13, 0xcf, 0xc3, 0x40, 0x74, 0x65,
	0x02, 0x7d, 0xd9, 0x12, 0xf7, 0x1b, 0x62, 0xd9, 0x10, 0xe7, 0x3f, 0x12, 0x18, 0x6f, 0x64, 0xa8,
	0x7d, 0x7e, 0x4f, 0xe1, 0x76, 0xbc, 0x8f, 0x56, 0x42, 0x3a, 0xea, 0x56, 0xba, 0x79, 0x0f, 0xc3,
	0x92, 0x74, 0xbd, 0xb7, 0xbe, 0xf5, 0x6e, 0xa5, 0x91, 0x3c, 0x83, 0xbb, 0x17, 0x37, 0x2a, 0xb4,
	0xdc, 0x08, 0xed, 0x91, 0x75, 0xae, 0x93, 0xf3, 0x8b, 0x2f, 0xe3, 0x40, 0xe7, 0x46, 0x1a, 0xdd,
	0xc8, 0xbf, 0xa7, 0x70, 0xef, 0xea, 0xa7, 0x4e, 0x7e, 0x27, 0x30, 0x6e, 0x9a, 0x58, 0xad, 0xdf,
	0x5f, 0x36, 0x0f, 0xa1, 0x7f, 0xbb, 0x06, 0xa1, 0x77, 0x1f, 0x44, 0x39, 0x6a, 0x45, 0x1e, 0x06,
	0x88, 0xfc, 0xea, 0x74, 0x77, 0x7e, 0x65, 0xbb, 0xd3, 0x74, 0x36, 0x98, 0x7f, 0xb9, 0x0e, 0xba,
	0xf5, 0x86, 0xe6, 0x36, 0xad, 0xc5, 0x76, 0xf8, 0xef, 0xed, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x0d, 0xbe, 0x0f, 0x8b, 0x1f, 0x05, 0x00, 0x00,
}
