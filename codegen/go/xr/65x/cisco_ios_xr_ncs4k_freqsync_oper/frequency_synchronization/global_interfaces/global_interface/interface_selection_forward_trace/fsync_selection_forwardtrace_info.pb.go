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

package cisco_ios_xr_ncs4k_freqsync_oper_frequency_synchronization_global_interfaces_global_interface_interface_selection_forward_trace

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
	proto.RegisterType((*FsyncSelectionForwardtraceInfo_KEYS)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_selection_forwardtrace_info_KEYS")
	proto.RegisterType((*FsyncBagSpId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_sp_id")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagForwardtraceNode)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_forwardtrace_node")
	proto.RegisterType((*FsyncBagForwardtrace)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_bag_forwardtrace")
	proto.RegisterType((*FsyncSelectionForwardtraceInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.global_interfaces.global_interface.interface_selection_forward_trace.fsync_selection_forwardtrace_info")
}

func init() {
	proto.RegisterFile("fsync_selection_forwardtrace_info.proto", fileDescriptor_cb48806671644731)
}

var fileDescriptor_cb48806671644731 = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4f, 0x6b, 0x14, 0x3f,
	0x18, 0x26, 0x6d, 0x7f, 0xdb, 0xdd, 0x77, 0xbb, 0xbb, 0x6d, 0x7e, 0x22, 0x53, 0x8b, 0xd0, 0x2e,
	0x68, 0x7b, 0x71, 0x90, 0xd1, 0xa3, 0x08, 0x52, 0x3d, 0x14, 0xa1, 0xca, 0x58, 0x10, 0x4f, 0x21,
	0xcd, 0x64, 0xdb, 0xd0, 0x69, 0x32, 0x26, 0xa9, 0xba, 0x5e, 0xfc, 0x0a, 0x82, 0x42, 0x2f, 0xde,
	0x3d, 0x89, 0xf7, 0x82, 0xe2, 0xdd, 0x8f, 0x21, 0xf8, 0x39, 0x24, 0x99, 0x7f, 0x9d, 0xae, 0xd0,
	0xa3, 0x73, 0x6a, 0xe6, 0xc9, 0x93, 0x27, 0xcf, 0x93, 0x37, 0x79, 0xbb, 0xb0, 0x39, 0x31, 0x53,
	0xc9, 0x88, 0xe1, 0x29, 0x67, 0x56, 0x28, 0x49, 0x26, 0x4a, 0xbf, 0xa6, 0x3a, 0xb1, 0x9a, 0x32,
	0x4e, 0x84, 0x9c, 0xa8, 0x30, 0xd3, 0xca, 0x2a, 0xfc, 0x8e, 0x09, 0xc3, 0x14, 0x11, 0xca, 0x90,
	0x37, 0x9a, 0x48, 0x66, 0xee, 0x1e, 0x91, 0x89, 0xe6, 0x2f, 0xfd, 0x72, 0x95, 0x71, 0x1d, 0xba,
	0xaf, 0x13, 0x2e, 0xd9, 0x94, 0x38, 0xec, 0x50, 0x2b, 0x29, 0xde, 0x52, 0xa7, 0x19, 0x1e, 0xa4,
	0x6a, 0x9f, 0xa6, 0x44, 0x48, 0xcb, 0xf5, 0x84, 0x32, 0x6e, 0x66, 0x90, 0xb0, 0x1a, 0xcd, 0x5a,
	0x21, 0xde, 0xcb, 0xf8, 0x09, 0xdc, 0xbc, 0xd4, 0x2b, 0x79, 0xfc, 0xe8, 0xc5, 0x33, 0x7c, 0x03,
	0x86, 0xb5, 0x9c, 0xa4, 0xc7, 0x3c, 0x40, 0xeb, 0x68, 0xab, 0x17, 0x0f, 0x2a, 0x74, 0x97, 0x1e,
	0xf3, 0xf1, 0x29, 0x82, 0x51, 0xae, 0xb8, 0x4f, 0x0f, 0x88, 0xc9, 0x88, 0x48, 0xf0, 0x6d, 0xb8,
	0x52, 0xcb, 0x67, 0x4a, 0x48, 0x4b, 0xec, 0x34, 0xcb, 0x05, 0x06, 0x31, 0xae, 0xe6, 0x9e, 0xba,
	0xa9, 0xbd, 0x69, 0xc6, 0xf1, 0x7d, 0x58, 0xbb, 0xb8, 0x22, 0xe1, 0x86, 0x69, 0x91, 0x39, 0x24,
	0x98, 0xf3, 0x3b, 0xaf, 0x36, 0x17, 0x3e, 0xac, 0x09, 0x18, 0xc3, 0x82, 0x54, 0x09, 0x0f, 0xe6,
	0x3d, 0xd1, 0x8f, 0xc7, 0xcf, 0x01, 0xd7, 0xc6, 0x58, 0xaa, 0xd8, 0x91, 0xf3, 0x56, 0x32, 0x51,
	0xcd, 0xc4, 0x43, 0x98, 0x13, 0x89, 0xdf, 0x64, 0x10, 0xcf, 0x89, 0x04, 0x5f, 0x07, 0xc8, 0xf9,
	0x3e, 0x76, 0xae, 0xd9, 0xf3, 0x88, 0x8f, 0x7c, 0xd6, 0x81, 0xff, 0xcf, 0x45, 0x56, 0x27, 0xda,
	0x1d, 0x5c, 0x82, 0x37, 0x60, 0xa9, 0xf8, 0x60, 0x29, 0x35, 0xa6, 0xd8, 0xa2, 0x9f, 0x63, 0xdb,
	0x0e, 0xc2, 0xb7, 0x00, 0x73, 0x7b, 0xc8, 0xb5, 0xe4, 0xb6, 0x2e, 0x5b, 0x11, 0x6f, 0xa5, 0x9c,
	0xd9, 0x29, 0x27, 0xf0, 0x26, 0x8c, 0x8c, 0x6a, 0x72, 0x73, 0x37, 0x43, 0x0f, 0xd7, 0xc4, 0xaf,
	0x08, 0xba, 0x65, 0xc4, 0x60, 0x61, 0x1d, 0x6d, 0xf5, 0xa3, 0x0f, 0x28, 0xfc, 0xc7, 0x97, 0x2d,
	0x9c, 0x3d, 0xfe, 0x78, 0xd1, 0x8f, 0x76, 0x12, 0xfc, 0x03, 0xc1, 0x8a, 0x57, 0x90, 0x34, 0xad,
	0xa6, 0x83, 0xff, 0x5a, 0x6c, 0x7d, 0x54, 0xda, 0xdd, 0x2e, 0x22, 0xac, 0x42, 0x37, 0xb3, 0x19,
	0xf1, 0xd7, 0xa9, 0xe3, 0xcb, 0xb2, 0x98, 0xd9, 0x6c, 0xd7, 0xdd, 0xa8, 0x7b, 0x70, 0xcd, 0x50,
	0xcb, 0xd3, 0x54, 0x58, 0x4e, 0x28, 0x63, 0xdc, 0x98, 0x73, 0x35, 0x5c, 0xf4, 0xe4, 0xa0, 0x62,
	0x3c, 0xf0, 0x84, 0xba, 0x9a, 0xab, 0xd0, 0x95, 0xa5, 0x70, 0x37, 0x17, 0x96, 0x85, 0xf0, 0x77,
	0x04, 0xcb, 0x07, 0xd2, 0x18, 0xa2, 0x39, 0xe3, 0xe2, 0x15, 0xd7, 0xee, 0xd4, 0x7a, 0x2d, 0x3e,
	0xb5, 0xa1, 0x73, 0x1b, 0x17, 0x66, 0x77, 0x92, 0xf1, 0xe7, 0x79, 0x58, 0xab, 0x69, 0x8d, 0xde,
	0xe3, 0xdf, 0xe2, 0x1a, 0xf4, 0xdc, 0xdf, 0xba, 0x61, 0xf4, 0xe2, 0xae, 0x03, 0x7c, 0x9b, 0x38,
	0x43, 0x30, 0xba, 0xd0, 0x27, 0xfc, 0xe3, 0xe9, 0x47, 0xef, 0xdb, 0x14, 0xde, 0x77, 0xc1, 0x78,
	0xd8, 0x6c, 0x57, 0xf8, 0x0b, 0x82, 0x4e, 0xde, 0x0b, 0xfc, 0x23, 0xee, 0x47, 0x1f, 0x5b, 0xe5,
	0xb9, 0x6c, 0x63, 0x71, 0x61, 0x72, 0xfc, 0x1b, 0xc1, 0xd5, 0xbf, 0x57, 0x0a, 0xff, 0x44, 0x80,
	0x1b, 0x42, 0xa4, 0xea, 0xa9, 0xfd, 0xe8, 0x53, 0x9b, 0x62, 0xcd, 0x5c, 0xb0, 0x78, 0xb9, 0x80,
	0xf6, 0x1c, 0xe4, 0xde, 0xd4, 0xf8, 0x17, 0x82, 0x8d, 0x4b, 0xff, 0x29, 0xe2, 0x6f, 0x08, 0x06,
	0x8d, 0x5d, 0x82, 0x68, 0x7d, 0x7e, 0xab, 0x1f, 0x9d, 0xb6, 0x35, 0x6e, 0xbc, 0x74, 0x3e, 0xe9,
	0x7e, 0xc7, 0xff, 0x02, 0xb9, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x96, 0xd4, 0x59, 0xac,
	0x08, 0x00, 0x00,
}
