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
// source: fsync_selection_forwardtrace_info.proto

package cisco_ios_xr_freqsync_oper_frequency_synchronization_global_interfaces_global_interface_interface_selection_forward_trace

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

type FsyncSelectionForwardtraceInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncSelectionForwardtraceInfo_KEYS) Reset()         { *m = FsyncSelectionForwardtraceInfo_KEYS{} }
func (m *FsyncSelectionForwardtraceInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncSelectionForwardtraceInfo_KEYS) ProtoMessage()    {}
func (*FsyncSelectionForwardtraceInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb48806671644731, []int{0}
}

func (m *FsyncSelectionForwardtraceInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS.Unmarshal(m, b)
}
func (m *FsyncSelectionForwardtraceInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncSelectionForwardtraceInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS.Merge(m, src)
}
func (m *FsyncSelectionForwardtraceInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS.Size(m)
}
func (m *FsyncSelectionForwardtraceInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSelectionForwardtraceInfo_KEYS proto.InternalMessageInfo

func (m *FsyncSelectionForwardtraceInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_cb48806671644731, []int{1}
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
	return fileDescriptor_cb48806671644731, []int{2}
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
	return fileDescriptor_cb48806671644731, []int{3}
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

type FsyncBagForwardtraceNode struct {
	NodeType             string            `protobuf:"bytes,1,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	SelectionPoint       *FsyncBagSpId     `protobuf:"bytes,2,opt,name=selection_point,json=selectionPoint,proto3" json:"selection_point,omitempty"`
	Source               *FsyncBagSourceId `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncBagForwardtraceNode) Reset()         { *m = FsyncBagForwardtraceNode{} }
func (m *FsyncBagForwardtraceNode) String() string { return proto.CompactTextString(m) }
func (*FsyncBagForwardtraceNode) ProtoMessage()    {}
func (*FsyncBagForwardtraceNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb48806671644731, []int{4}
}

func (m *FsyncBagForwardtraceNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagForwardtraceNode.Unmarshal(m, b)
}
func (m *FsyncBagForwardtraceNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagForwardtraceNode.Marshal(b, m, deterministic)
}
func (m *FsyncBagForwardtraceNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagForwardtraceNode.Merge(m, src)
}
func (m *FsyncBagForwardtraceNode) XXX_Size() int {
	return xxx_messageInfo_FsyncBagForwardtraceNode.Size(m)
}
func (m *FsyncBagForwardtraceNode) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagForwardtraceNode.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagForwardtraceNode proto.InternalMessageInfo

func (m *FsyncBagForwardtraceNode) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func (m *FsyncBagForwardtraceNode) GetSelectionPoint() *FsyncBagSpId {
	if m != nil {
		return m.SelectionPoint
	}
	return nil
}

func (m *FsyncBagForwardtraceNode) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

type FsyncBagForwardtrace struct {
	ForwardTraceNode     *FsyncBagForwardtraceNode `protobuf:"bytes,1,opt,name=forward_trace_node,json=forwardTraceNode,proto3" json:"forward_trace_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FsyncBagForwardtrace) Reset()         { *m = FsyncBagForwardtrace{} }
func (m *FsyncBagForwardtrace) String() string { return proto.CompactTextString(m) }
func (*FsyncBagForwardtrace) ProtoMessage()    {}
func (*FsyncBagForwardtrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb48806671644731, []int{5}
}

func (m *FsyncBagForwardtrace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagForwardtrace.Unmarshal(m, b)
}
func (m *FsyncBagForwardtrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagForwardtrace.Marshal(b, m, deterministic)
}
func (m *FsyncBagForwardtrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagForwardtrace.Merge(m, src)
}
func (m *FsyncBagForwardtrace) XXX_Size() int {
	return xxx_messageInfo_FsyncBagForwardtrace.Size(m)
}
func (m *FsyncBagForwardtrace) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagForwardtrace.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagForwardtrace proto.InternalMessageInfo

func (m *FsyncBagForwardtrace) GetForwardTraceNode() *FsyncBagForwardtraceNode {
	if m != nil {
		return m.ForwardTraceNode
	}
	return nil
}

type FsyncSelectionForwardtraceInfo struct {
	ForwardTrace         []*FsyncBagForwardtrace `protobuf:"bytes,50,rep,name=forward_trace,json=forwardTrace,proto3" json:"forward_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FsyncSelectionForwardtraceInfo) Reset()         { *m = FsyncSelectionForwardtraceInfo{} }
func (m *FsyncSelectionForwardtraceInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncSelectionForwardtraceInfo) ProtoMessage()    {}
func (*FsyncSelectionForwardtraceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb48806671644731, []int{6}
}

func (m *FsyncSelectionForwardtraceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo.Unmarshal(m, b)
}
func (m *FsyncSelectionForwardtraceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo.Marshal(b, m, deterministic)
}
func (m *FsyncSelectionForwardtraceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSelectionForwardtraceInfo.Merge(m, src)
}
func (m *FsyncSelectionForwardtraceInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncSelectionForwardtraceInfo.Size(m)
}
func (m *FsyncSelectionForwardtraceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSelectionForwardtraceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSelectionForwardtraceInfo proto.InternalMessageInfo

func (m *FsyncSelectionForwardtraceInfo) GetForwardTrace() []*FsyncBagForwardtrace {
	if m != nil {
		return m.ForwardTrace
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncSelectionForwardtraceInfo_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_selection_forwardtrace_info_KEYS")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_sp_id")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagForwardtraceNode)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_forwardtrace_node")
	proto.RegisterType((*FsyncBagForwardtrace)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_forwardtrace")
	proto.RegisterType((*FsyncSelectionForwardtraceInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_selection_forwardtrace_info")
}

func init() {
	proto.RegisterFile("fsync_selection_forwardtrace_info.proto", fileDescriptor_cb48806671644731)
}

var fileDescriptor_cb48806671644731 = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6b, 0x14, 0x31,
	0x14, 0x27, 0x6d, 0xdd, 0xee, 0xbe, 0xed, 0xee, 0xb6, 0x51, 0x64, 0x6a, 0x11, 0xda, 0x05, 0x6d,
	0x2f, 0x0e, 0x32, 0x5e, 0x45, 0x90, 0xea, 0xa1, 0x08, 0x55, 0xc6, 0x82, 0x78, 0x0a, 0x69, 0x26,
	0xdb, 0x06, 0xa7, 0xc9, 0x98, 0xa4, 0xea, 0xfa, 0x11, 0x04, 0x05, 0x2f, 0x7a, 0xf0, 0x03, 0xa8,
	0x77, 0x41, 0xf1, 0xee, 0xc7, 0x10, 0xfc, 0x28, 0x92, 0xcc, 0xbf, 0x9d, 0xae, 0xd0, 0xa3, 0x73,
	0x6a, 0xe6, 0xf7, 0x7e, 0xf9, 0xe5, 0xf7, 0xf2, 0xf2, 0x5e, 0x17, 0xb6, 0x27, 0x66, 0x2a, 0x19,
	0x31, 0x3c, 0xe5, 0xcc, 0x0a, 0x25, 0xc9, 0x44, 0xe9, 0x97, 0x54, 0x27, 0x56, 0x53, 0xc6, 0x89,
	0x90, 0x13, 0x15, 0x66, 0x5a, 0x59, 0x85, 0xa7, 0x4c, 0x18, 0xa6, 0x88, 0x50, 0x86, 0xbc, 0xd2,
	0x64, 0xa2, 0xf9, 0x73, 0xbf, 0x51, 0x65, 0x5c, 0x87, 0xee, 0xeb, 0x94, 0x4b, 0x36, 0x25, 0x0e,
	0x3b, 0xd6, 0x4a, 0x8a, 0xd7, 0xd4, 0xa9, 0x85, 0x47, 0xa9, 0x3a, 0xa4, 0x29, 0x11, 0xd2, 0x72,
	0x3d, 0xa1, 0x8c, 0x9b, 0x39, 0x24, 0xac, 0x56, 0xf3, 0x26, 0x88, 0x77, 0x31, 0x7e, 0x08, 0xd7,
	0xcf, 0x75, 0x49, 0x1e, 0xdc, 0x7f, 0xfa, 0x18, 0x5f, 0x83, 0x61, 0x2d, 0x27, 0xe9, 0x09, 0x0f,
	0xd0, 0x26, 0xda, 0xe9, 0xc5, 0x83, 0x0a, 0xdd, 0xa7, 0x27, 0x7c, 0xfc, 0x11, 0xc1, 0x28, 0x57,
	0x3c, 0xa4, 0x47, 0xc4, 0x64, 0x44, 0x24, 0xf8, 0x26, 0x5c, 0xaa, 0xe5, 0x33, 0x25, 0xa4, 0x25,
	0x76, 0x9a, 0xe5, 0x02, 0x83, 0x18, 0x57, 0xb1, 0x47, 0x2e, 0x74, 0x30, 0xcd, 0x38, 0xbe, 0x03,
	0x1b, 0x67, 0x77, 0x24, 0xdc, 0x30, 0x2d, 0x32, 0x87, 0x04, 0x0b, 0xfe, 0xe4, 0xf5, 0xe6, 0xc6,
	0x7b, 0x35, 0x01, 0x63, 0x58, 0x92, 0x2a, 0xe1, 0xc1, 0xa2, 0x27, 0xfa, 0xf5, 0xf8, 0x09, 0xe0,
	0xda, 0x18, 0x4b, 0x15, 0x7b, 0xe6, 0xbc, 0x95, 0x4c, 0x54, 0x33, 0xf1, 0x10, 0x16, 0x44, 0xe2,
	0x0f, 0x19, 0xc4, 0x0b, 0x22, 0xc1, 0x57, 0x01, 0x72, 0xbe, 0x4f, 0x3b, 0xd7, 0xec, 0x79, 0xc4,
	0xa7, 0xfc, 0xa5, 0x03, 0x17, 0x67, 0x52, 0x56, 0xa7, 0xda, 0x5d, 0x5c, 0x82, 0xb7, 0x60, 0xa5,
	0xf8, 0x60, 0x29, 0x35, 0xa6, 0x38, 0xa2, 0x9f, 0x63, 0xbb, 0x0e, 0xc2, 0x37, 0x00, 0x73, 0x7b,
	0xcc, 0xb5, 0xe4, 0xb6, 0x2e, 0x5b, 0x91, 0xde, 0x5a, 0x19, 0xd9, 0x2b, 0x03, 0x78, 0x1b, 0x46,
	0x46, 0x35, 0xb9, 0xb9, 0x9b, 0xa1, 0x87, 0x6b, 0xe2, 0x57, 0x04, 0xdd, 0x32, 0xc5, 0x60, 0x69,
	0x13, 0xed, 0xf4, 0xa3, 0xb7, 0x28, 0xfc, 0x6f, 0xcf, 0x2c, 0x9c, 0xbf, 0xf8, 0x78, 0xd9, 0xaf,
	0xf6, 0x12, 0xfc, 0x13, 0xc1, 0x9a, 0x57, 0x90, 0x34, 0xad, 0xc2, 0xc1, 0x85, 0x56, 0x9a, 0x1e,
	0x95, 0x46, 0x77, 0x0b, 0xf3, 0xeb, 0xd0, 0xcd, 0x6c, 0x46, 0xfc, 0x13, 0xea, 0xf8, 0x52, 0x2c,
	0x67, 0x36, 0xdb, 0x77, 0xaf, 0xe8, 0x36, 0x5c, 0x31, 0xd4, 0xf2, 0x34, 0x15, 0x96, 0x13, 0xca,
	0x18, 0x37, 0x66, 0xa6, 0x6e, 0xcb, 0x9e, 0x1c, 0x54, 0x8c, 0xbb, 0x9e, 0x50, 0x57, 0x70, 0x1d,
	0xba, 0xb2, 0x14, 0xee, 0xe6, 0xc2, 0xb2, 0x10, 0xfe, 0x81, 0x60, 0xf5, 0x48, 0x1a, 0x43, 0x34,
	0x67, 0x5c, 0xbc, 0xe0, 0xda, 0xdd, 0x57, 0xaf, 0x95, 0xf7, 0x35, 0x74, 0x3e, 0xe3, 0xc2, 0xe6,
	0x5e, 0x32, 0xfe, 0xb4, 0x08, 0x1b, 0x35, 0xad, 0x31, 0x69, 0x7c, 0xe7, 0x6d, 0x40, 0xcf, 0xfd,
	0xad, 0xc7, 0x43, 0x2f, 0xee, 0x3a, 0xc0, 0x0f, 0x85, 0x6f, 0x08, 0x46, 0x67, 0xa6, 0x82, 0x6f,
	0x95, 0x7e, 0xf4, 0xa6, 0x1d, 0x69, 0xfb, 0x69, 0x17, 0x0f, 0x9b, 0x63, 0x09, 0x7f, 0x46, 0xd0,
	0xc9, 0x7b, 0xde, 0x37, 0x6b, 0x3f, 0x7a, 0xd7, 0x12, 0xb7, 0xe5, 0xa0, 0x8a, 0x0b, 0x7b, 0xe3,
	0x3f, 0x08, 0x2e, 0xff, 0xbb, 0x3a, 0xf8, 0x17, 0x02, 0xdc, 0x10, 0x22, 0xd5, 0xd4, 0xec, 0x47,
	0x1f, 0xda, 0x91, 0xd0, 0xdc, 0x73, 0x8a, 0x57, 0x0b, 0xe8, 0xc0, 0x41, 0xae, 0x77, 0xc6, 0xbf,
	0x11, 0x6c, 0x9d, 0xfb, 0x0f, 0x0f, 0x7f, 0x47, 0x30, 0x68, 0x9c, 0x12, 0x44, 0x9b, 0x8b, 0x3b,
	0xfd, 0xe8, 0x7d, 0xfb, 0x12, 0x8d, 0x57, 0x66, 0x73, 0x3c, 0xec, 0xf8, 0x5f, 0x14, 0xb7, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x2a, 0xdb, 0xde, 0x7c, 0x08, 0x00, 0x00,
}