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

package cisco_ios_xr_freqsync_oper_frequency_synchronization_global_nodes_global_node_clock_interface_selection_forward_traces_clock_interface_selection_forward_trace

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
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClockType            string   `protobuf:"bytes,2,opt,name=clock_type,json=clockType,proto3" json:"clock_type,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
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

func (m *FsyncSelectionForwardtraceInfo_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *FsyncSelectionForwardtraceInfo_KEYS) GetClockType() string {
	if m != nil {
		return m.ClockType
	}
	return ""
}

func (m *FsyncSelectionForwardtraceInfo_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
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
	proto.RegisterType((*FsyncSelectionForwardtraceInfo_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_selection_forwardtrace_info_KEYS")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_bag_sp_id")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagForwardtraceNode)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_bag_forwardtrace_node")
	proto.RegisterType((*FsyncBagForwardtrace)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_bag_forwardtrace")
	proto.RegisterType((*FsyncSelectionForwardtraceInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.global_nodes.global_node.clock_interface_selection_forward_traces.clock_interface_selection_forward_trace.fsync_selection_forwardtrace_info")
}

func init() {
	proto.RegisterFile("fsync_selection_forwardtrace_info.proto", fileDescriptor_cb48806671644731)
}

var fileDescriptor_cb48806671644731 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x41, 0x6b, 0x14, 0x3f,
	0x14, 0x27, 0xbb, 0xfd, 0x77, 0x77, 0xdf, 0xb6, 0xbb, 0x6d, 0xfe, 0x22, 0x53, 0x8b, 0xb0, 0x9d,
	0x83, 0xdd, 0x8b, 0x8b, 0x8c, 0x57, 0x11, 0xa4, 0x7a, 0x28, 0x42, 0x91, 0xb1, 0x20, 0x9e, 0xc2,
	0x34, 0x93, 0x6d, 0x83, 0xd3, 0x24, 0x26, 0xa9, 0xba, 0x1e, 0x05, 0x3f, 0x43, 0x6f, 0x7e, 0x05,
	0x41, 0xc1, 0x8b, 0x5f, 0x41, 0xbc, 0x0b, 0x8a, 0x27, 0x2f, 0x7e, 0x0a, 0x49, 0x66, 0x66, 0x67,
	0xa7, 0x2d, 0xd4, 0xa3, 0x73, 0xda, 0xe4, 0xbd, 0x97, 0xf7, 0xde, 0xef, 0x97, 0x97, 0xdf, 0x0e,
	0x6c, 0x4f, 0xcd, 0x4c, 0x50, 0x62, 0x58, 0xc6, 0xa8, 0xe5, 0x52, 0x90, 0xa9, 0xd4, 0x2f, 0x13,
	0x9d, 0x5a, 0x9d, 0x50, 0x46, 0xb8, 0x98, 0xca, 0x89, 0xd2, 0xd2, 0x4a, 0xfc, 0x0e, 0x51, 0x6e,
	0xa8, 0x24, 0x5c, 0x1a, 0xf2, 0x4a, 0x93, 0xa9, 0x66, 0xcf, 0xfd, 0x49, 0xa9, 0x98, 0x9e, 0xb8,
	0xdd, 0x09, 0x13, 0x74, 0x46, 0x9c, 0xed, 0x48, 0x4b, 0xc1, 0x5f, 0x27, 0x2e, 0xdd, 0xe4, 0x30,
	0x93, 0x07, 0x49, 0x46, 0x84, 0x4c, 0x99, 0x59, 0xdc, 0x4c, 0x68, 0x26, 0xe9, 0x33, 0xc2, 0x85,
	0x65, 0x7a, 0xea, 0x4a, 0x9d, 0xeb, 0x80, 0xf8, 0x16, 0xcc, 0xdf, 0x06, 0x86, 0x12, 0x6e, 0x5c,
	0x8a, 0x85, 0x3c, 0x7c, 0xf0, 0xf4, 0x31, 0xc6, 0xb0, 0xe4, 0x5a, 0x08, 0xd0, 0x08, 0x8d, 0x7b,
	0xb1, 0x5f, 0xe3, 0xeb, 0x00, 0x79, 0x21, 0x3b, 0x53, 0x2c, 0x68, 0x79, 0x4f, 0xcf, 0x5b, 0xf6,
	0x67, 0x8a, 0xb9, 0x23, 0x4a, 0x6a, 0x1b, 0xb4, 0x47, 0x68, 0xbc, 0x1a, 0xfb, 0x75, 0x78, 0x8a,
	0x60, 0x98, 0x57, 0x3c, 0x48, 0x0e, 0x89, 0x51, 0x84, 0xa7, 0xf8, 0x16, 0x5c, 0xa9, 0xca, 0x2b,
	0xc9, 0x85, 0xcd, 0x13, 0x22, 0x7f, 0x0e, 0xcf, 0x7d, 0x8f, 0x9c, 0xcb, 0x67, 0xbe, 0x0b, 0x9b,
	0x67, 0x4f, 0xa4, 0xcc, 0x50, 0xcd, 0x95, 0xb3, 0x14, 0x9d, 0x6c, 0xd4, 0x0f, 0xde, 0xaf, 0x02,
	0xe6, 0x60, 0xda, 0x15, 0x98, 0xf0, 0x09, 0xe0, 0xaa, 0xb1, 0x82, 0xbf, 0xf4, 0x42, 0xd8, 0x03,
	0x68, 0xf1, 0xd4, 0x17, 0x59, 0x8d, 0x5b, 0x3c, 0xad, 0x68, 0x10, 0xc9, 0x71, 0x99, 0x33, 0xa7,
	0x61, 0x2f, 0x39, 0x66, 0xe1, 0x69, 0x07, 0xfe, 0x5f, 0x80, 0x2c, 0x4f, 0xb4, 0x23, 0x36, 0xc5,
	0x5b, 0xb0, 0x52, 0x6c, 0x68, 0x96, 0x18, 0x53, 0x94, 0xe8, 0xe7, 0xb6, 0x1d, 0x67, 0xc2, 0x37,
	0x01, 0x33, 0x7b, 0xc4, 0xb4, 0x60, 0xb6, 0xba, 0xcc, 0x02, 0xde, 0x7a, 0xe9, 0xd9, 0x2d, 0x1d,
	0x78, 0x1b, 0x86, 0x46, 0xd6, 0x63, 0xf3, 0x6e, 0x06, 0xde, 0x5c, 0x05, 0x7e, 0x45, 0xd0, 0x2d,
	0x21, 0x06, 0x4b, 0x23, 0x34, 0xee, 0x47, 0x1f, 0xd0, 0xe4, 0xdf, 0x9e, 0xd5, 0xc9, 0xf9, 0xdb,
	0x89, 0x3b, 0x7e, 0xb5, 0x9b, 0xe2, 0x9f, 0x08, 0xd6, 0x7d, 0x06, 0x91, 0x64, 0x73, 0x77, 0xf0,
	0x5f, 0x73, 0x91, 0x0d, 0x4b, 0x34, 0x3b, 0x05, 0xc2, 0x0d, 0xe8, 0x2a, 0xab, 0x7c, 0x03, 0xc1,
	0xb2, 0xbf, 0xd4, 0x8e, 0xb2, 0x6a, 0xcf, 0xcd, 0xe3, 0x1d, 0xb8, 0x66, 0x12, 0xcb, 0xb2, 0x8c,
	0x5b, 0x46, 0x12, 0x4a, 0x99, 0x31, 0x0b, 0x13, 0xd0, 0xf1, 0xc1, 0xc1, 0x3c, 0xe2, 0x9e, 0x0f,
	0xa8, 0x66, 0x61, 0x03, 0xba, 0xa2, 0x4c, 0xdc, 0xcd, 0x13, 0x8b, 0x22, 0xf1, 0x0f, 0x04, 0x6b,
	0x87, 0xc2, 0x18, 0xa2, 0x19, 0x65, 0xfc, 0x05, 0xd3, 0x8e, 0xd4, 0x5e, 0x73, 0x49, 0x1d, 0x38,
	0x30, 0x71, 0x81, 0x65, 0x37, 0x0d, 0x7f, 0xb5, 0x61, 0xb3, 0x0a, 0xab, 0x09, 0x9f, 0x7f, 0xe8,
	0x9b, 0xd0, 0x73, 0xbf, 0x95, 0x1a, 0xf5, 0xe2, 0xae, 0x33, 0x78, 0x0d, 0xfa, 0x86, 0x60, 0x78,
	0x46, 0x84, 0xfc, 0xcb, 0xec, 0x47, 0xef, 0x1b, 0xc4, 0x8d, 0x57, 0xe0, 0x78, 0x50, 0x97, 0x4a,
	0xfc, 0x05, 0xc1, 0x72, 0xae, 0x43, 0x5e, 0x40, 0xfa, 0xd1, 0xc7, 0x26, 0x41, 0x2a, 0x15, 0x36,
	0x2e, 0x30, 0x84, 0x6f, 0x5b, 0x70, 0xf5, 0xe2, 0x7b, 0xc6, 0xbf, 0x11, 0xe0, 0x5a, 0x22, 0x32,
	0x97, 0xfb, 0x7e, 0xf4, 0xb9, 0x41, 0xa8, 0xcf, 0x4d, 0x6f, 0xbc, 0x56, 0x98, 0xf6, 0x9d, 0xc9,
	0xbd, 0xe7, 0xf0, 0x4d, 0x0b, 0xb6, 0x2e, 0xfd, 0xbb, 0xc7, 0xdf, 0x11, 0xac, 0xd6, 0xaa, 0x04,
	0xd1, 0xa8, 0x3d, 0xee, 0x47, 0x9f, 0x1a, 0xca, 0x46, 0xbc, 0xb2, 0x48, 0xc4, 0xc1, 0xb2, 0xff,
	0x34, 0xbb, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x44, 0xd8, 0xa7, 0xc5, 0x09, 0x00, 0x00,
}
