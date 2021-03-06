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
// source: fsync_summary_info.proto

package cisco_ios_xr_freqsync_oper_frequency_synchronization_summary

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

type FsyncSummaryInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsyncSummaryInfo_KEYS) Reset()         { *m = FsyncSummaryInfo_KEYS{} }
func (m *FsyncSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsyncSummaryInfo_KEYS) ProtoMessage()    {}
func (*FsyncSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7995d1811c32674, []int{0}
}

func (m *FsyncSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *FsyncSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FsyncSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSummaryInfo_KEYS.Merge(m, src)
}
func (m *FsyncSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsyncSummaryInfo_KEYS.Size(m)
}
func (m *FsyncSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSummaryInfo_KEYS proto.InternalMessageInfo

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
	return fileDescriptor_b7995d1811c32674, []int{1}
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
	return fileDescriptor_b7995d1811c32674, []int{2}
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

type FsyncBagSummaryFreqInfo struct {
	Source               *FsyncBagSourceId `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	ClockCount           uint32            `protobuf:"varint,2,opt,name=clock_count,json=clockCount,proto3" json:"clock_count,omitempty"`
	EthernetCount        uint32            `protobuf:"varint,3,opt,name=ethernet_count,json=ethernetCount,proto3" json:"ethernet_count,omitempty"`
	SonetCount           uint32            `protobuf:"varint,4,opt,name=sonet_count,json=sonetCount,proto3" json:"sonet_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncBagSummaryFreqInfo) Reset()         { *m = FsyncBagSummaryFreqInfo{} }
func (m *FsyncBagSummaryFreqInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSummaryFreqInfo) ProtoMessage()    {}
func (*FsyncBagSummaryFreqInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7995d1811c32674, []int{3}
}

func (m *FsyncBagSummaryFreqInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSummaryFreqInfo.Unmarshal(m, b)
}
func (m *FsyncBagSummaryFreqInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSummaryFreqInfo.Marshal(b, m, deterministic)
}
func (m *FsyncBagSummaryFreqInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSummaryFreqInfo.Merge(m, src)
}
func (m *FsyncBagSummaryFreqInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSummaryFreqInfo.Size(m)
}
func (m *FsyncBagSummaryFreqInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSummaryFreqInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSummaryFreqInfo proto.InternalMessageInfo

func (m *FsyncBagSummaryFreqInfo) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FsyncBagSummaryFreqInfo) GetClockCount() uint32 {
	if m != nil {
		return m.ClockCount
	}
	return 0
}

func (m *FsyncBagSummaryFreqInfo) GetEthernetCount() uint32 {
	if m != nil {
		return m.EthernetCount
	}
	return 0
}

func (m *FsyncBagSummaryFreqInfo) GetSonetCount() uint32 {
	if m != nil {
		return m.SonetCount
	}
	return 0
}

type FsyncBagSummaryTodInfo struct {
	Source               *FsyncBagSourceId `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	NodeCount            uint32            `protobuf:"varint,2,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FsyncBagSummaryTodInfo) Reset()         { *m = FsyncBagSummaryTodInfo{} }
func (m *FsyncBagSummaryTodInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncBagSummaryTodInfo) ProtoMessage()    {}
func (*FsyncBagSummaryTodInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7995d1811c32674, []int{4}
}

func (m *FsyncBagSummaryTodInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncBagSummaryTodInfo.Unmarshal(m, b)
}
func (m *FsyncBagSummaryTodInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncBagSummaryTodInfo.Marshal(b, m, deterministic)
}
func (m *FsyncBagSummaryTodInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncBagSummaryTodInfo.Merge(m, src)
}
func (m *FsyncBagSummaryTodInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncBagSummaryTodInfo.Size(m)
}
func (m *FsyncBagSummaryTodInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncBagSummaryTodInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncBagSummaryTodInfo proto.InternalMessageInfo

func (m *FsyncBagSummaryTodInfo) GetSource() *FsyncBagSourceId {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FsyncBagSummaryTodInfo) GetNodeCount() uint32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

type FsyncSummaryInfo struct {
	FrequencySummary     []*FsyncBagSummaryFreqInfo `protobuf:"bytes,50,rep,name=frequency_summary,json=frequencySummary,proto3" json:"frequency_summary,omitempty"`
	TimeOfDaySummary     []*FsyncBagSummaryTodInfo  `protobuf:"bytes,51,rep,name=time_of_day_summary,json=timeOfDaySummary,proto3" json:"time_of_day_summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FsyncSummaryInfo) Reset()         { *m = FsyncSummaryInfo{} }
func (m *FsyncSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*FsyncSummaryInfo) ProtoMessage()    {}
func (*FsyncSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7995d1811c32674, []int{5}
}

func (m *FsyncSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsyncSummaryInfo.Unmarshal(m, b)
}
func (m *FsyncSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsyncSummaryInfo.Marshal(b, m, deterministic)
}
func (m *FsyncSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsyncSummaryInfo.Merge(m, src)
}
func (m *FsyncSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_FsyncSummaryInfo.Size(m)
}
func (m *FsyncSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FsyncSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FsyncSummaryInfo proto.InternalMessageInfo

func (m *FsyncSummaryInfo) GetFrequencySummary() []*FsyncBagSummaryFreqInfo {
	if m != nil {
		return m.FrequencySummary
	}
	return nil
}

func (m *FsyncSummaryInfo) GetTimeOfDaySummary() []*FsyncBagSummaryTodInfo {
	if m != nil {
		return m.TimeOfDaySummary
	}
	return nil
}

func init() {
	proto.RegisterType((*FsyncSummaryInfo_KEYS)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_summary_info_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagSummaryFreqInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_bag_summary_freq_info")
	proto.RegisterType((*FsyncBagSummaryTodInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_bag_summary_tod_info")
	proto.RegisterType((*FsyncSummaryInfo)(nil), "cisco_ios_xr_freqsync_oper.frequency_synchronization.summary.fsync_summary_info")
}

func init() { proto.RegisterFile("fsync_summary_info.proto", fileDescriptor_b7995d1811c32674) }

var fileDescriptor_b7995d1811c32674 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0x21, 0x1f, 0xb3, 0xb4, 0xb4, 0xee, 0x81, 0x4d, 0x11, 0x6a, 0x58, 0x09, 0x91,
	0x0b, 0x7b, 0x48, 0xaf, 0xb9, 0xa0, 0xc2, 0xa1, 0x42, 0xe2, 0x63, 0x7b, 0xa1, 0x27, 0x6b, 0xeb,
	0x78, 0xa9, 0xd5, 0xc4, 0x5e, 0x6c, 0x47, 0x22, 0xdc, 0x81, 0x7f, 0x02, 0xbf, 0x8f, 0x7f, 0x80,
	0x3c, 0x5e, 0x67, 0xb7, 0xe4, 0x48, 0xd4, 0x9b, 0xf7, 0xcd, 0xf3, 0xf3, 0xcc, 0x7b, 0xf6, 0x42,
	0x52, 0x9a, 0xb5, 0x64, 0xd4, 0xac, 0x96, 0xcb, 0x42, 0xaf, 0xa9, 0x90, 0xa5, 0xca, 0x2a, 0xad,
	0xac, 0x22, 0x33, 0x26, 0x0c, 0x53, 0x54, 0x28, 0x43, 0xbf, 0x6a, 0x5a, 0x6a, 0xfe, 0x05, 0x99,
	0xaa, 0xe2, 0x3a, 0x73, 0x5f, 0x2b, 0x2e, 0xd9, 0x9a, 0x3a, 0xec, 0x46, 0x2b, 0x29, 0xbe, 0x15,
	0x56, 0x28, 0x99, 0xd5, 0x42, 0xe9, 0x08, 0x1e, 0x6f, 0x2b, 0xd3, 0xb7, 0x6f, 0xae, 0x2e, 0xd3,
	0x19, 0x10, 0x5f, 0xba, 0x2e, 0x3e, 0x53, 0xb6, 0x50, 0xec, 0x96, 0x8a, 0x39, 0x21, 0xd0, 0x95,
	0x6a, 0xce, 0x93, 0x68, 0x1c, 0x4d, 0x86, 0x39, 0xae, 0x1d, 0x56, 0x29, 0x6d, 0x93, 0xbd, 0x71,
	0x34, 0xd9, 0xcf, 0x71, 0x9d, 0x7e, 0xef, 0xc0, 0x71, 0xb3, 0xdd, 0xa8, 0x95, 0x66, 0xdc, 0xed,
	0x7f, 0x06, 0x0f, 0xeb, 0x0f, 0xb6, 0x28, 0x8c, 0xa9, 0x75, 0x62, 0x8f, 0x9d, 0x3b, 0x88, 0xbc,
	0x04, 0xc2, 0xed, 0x0d, 0xd7, 0x92, 0x5b, 0x2a, 0xa4, 0xe5, 0xba, 0x2c, 0x18, 0x47, 0xf1, 0x61,
	0x7e, 0x14, 0x2a, 0x17, 0xa1, 0x40, 0x5e, 0xc0, 0x23, 0xa3, 0xee, 0x72, 0x3b, 0xc8, 0x3d, 0x40,
	0xb8, 0x21, 0xde, 0xc2, 0x20, 0x8c, 0x91, 0x74, 0xc7, 0xd1, 0x24, 0x9e, 0x7e, 0xc8, 0xfe, 0xc7,
	0xbc, 0x6c, 0xdb, 0x9e, 0xbc, 0x8f, 0xab, 0x8b, 0xc6, 0xa7, 0x07, 0x2d, 0x9f, 0x46, 0x30, 0xa8,
	0x6c, 0x45, 0x11, 0xef, 0x21, 0xde, 0xaf, 0x6c, 0xf5, 0xce, 0x95, 0x66, 0x70, 0x62, 0x0a, 0xcb,
	0x17, 0x0b, 0x61, 0x39, 0x2d, 0x18, 0xe3, 0xc6, 0xb4, 0xe6, 0xe9, 0x23, 0x39, 0xd9, 0x30, 0x5e,
	0x21, 0xa1, 0x99, 0x6c, 0x04, 0x03, 0x19, 0x84, 0x07, 0x5e, 0x58, 0x7a, 0xe1, 0xf4, 0x4f, 0x04,
	0x4f, 0x5a, 0x39, 0xd4, 0x29, 0xbb, 0xd9, 0x30, 0x6a, 0x22, 0xa0, 0xe7, 0xbd, 0xc7, 0x24, 0xe2,
	0xe9, 0xc7, 0x5d, 0x59, 0xb2, 0x89, 0x3c, 0xaf, 0x0f, 0x20, 0xa7, 0x10, 0x7b, 0x9f, 0x98, 0x5a,
	0xc9, 0x70, 0x5b, 0x00, 0xa1, 0x73, 0x87, 0x90, 0xe7, 0x70, 0xb0, 0x09, 0xde, 0x73, 0x3a, 0xc8,
	0xd9, 0x0f, 0xa8, 0xa7, 0x9d, 0x42, 0xec, 0x03, 0xf7, 0x9c, 0xae, 0xd7, 0x41, 0x08, 0x09, 0xe9,
	0xaf, 0x08, 0x4e, 0xb6, 0x67, 0xb6, 0x6a, 0x7e, 0xef, 0x23, 0x3f, 0x05, 0x70, 0xa1, 0xdc, 0x99,
	0x78, 0xe8, 0x10, 0xdf, 0xe8, 0xef, 0xbd, 0xf0, 0xc6, 0xda, 0xcf, 0x8f, 0xfc, 0x88, 0xe0, 0xa8,
	0x75, 0xae, 0x2f, 0x25, 0xd3, 0x71, 0x67, 0x12, 0x4f, 0xaf, 0x76, 0xd6, 0xec, 0xbf, 0x57, 0x21,
	0x3f, 0xdc, 0x6c, 0xbf, 0xf4, 0x35, 0xf2, 0x33, 0x82, 0x63, 0x2b, 0x96, 0x9c, 0xaa, 0x92, 0xce,
	0x8b, 0xa6, 0x95, 0x33, 0x6c, 0xe5, 0xd3, 0xae, 0x5b, 0x09, 0x09, 0xe5, 0x87, 0xee, 0xd0, 0xf7,
	0xe5, 0xeb, 0x22, 0x74, 0x72, 0xdd, 0xc3, 0x9f, 0xdd, 0xd9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x04, 0xdb, 0xc9, 0x32, 0x08, 0x05, 0x00, 0x00,
}
