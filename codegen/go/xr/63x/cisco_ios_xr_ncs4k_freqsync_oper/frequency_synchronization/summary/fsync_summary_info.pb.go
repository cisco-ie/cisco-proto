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

package cisco_ios_xr_ncs4k_freqsync_oper_frequency_synchronization_summary

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
	proto.RegisterType((*FsyncSummaryInfo_KEYS)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_summary_info_KEYS")
	proto.RegisterType((*FsyncBagClockId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_bag_clock_id")
	proto.RegisterType((*FsyncBagSourceId)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_bag_source_id")
	proto.RegisterType((*FsyncBagSummaryFreqInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_bag_summary_freq_info")
	proto.RegisterType((*FsyncBagSummaryTodInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_bag_summary_tod_info")
	proto.RegisterType((*FsyncSummaryInfo)(nil), "cisco_ios_xr_ncs4k_freqsync_oper.frequency_synchronization.summary.fsync_summary_info")
}

func init() { proto.RegisterFile("fsync_summary_info.proto", fileDescriptor_b7995d1811c32674) }

var fileDescriptor_b7995d1811c32674 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xd6, 0xd2, 0x1f, 0x2f, 0x6c, 0x6c, 0xde, 0x81, 0x74, 0x08, 0xad, 0x44, 0x42, 0xf4,
	0x42, 0x0e, 0x1d, 0xc7, 0x5d, 0x60, 0x70, 0x98, 0x90, 0x40, 0xca, 0x24, 0x10, 0x17, 0xac, 0xcc,
	0x75, 0x98, 0xb5, 0xd6, 0xce, 0x6c, 0x57, 0xa2, 0xfc, 0x0b, 0xc0, 0x7f, 0xc2, 0x8d, 0x3f, 0x8e,
	0x2b, 0xf2, 0x73, 0xdc, 0x64, 0xf4, 0x48, 0x11, 0x37, 0xe7, 0x7b, 0x9f, 0x3f, 0xbf, 0xef, 0x7d,
	0x76, 0x20, 0x29, 0xcd, 0x4a, 0x32, 0x6a, 0x96, 0x8b, 0x45, 0xa1, 0x57, 0x54, 0xc8, 0x52, 0x65,
	0x95, 0x56, 0x56, 0x91, 0x17, 0x4c, 0x18, 0xa6, 0xa8, 0x50, 0x86, 0x7e, 0xd6, 0x54, 0x32, 0xf3,
	0xec, 0x9a, 0x96, 0x9a, 0xdf, 0x20, 0x5f, 0x55, 0x5c, 0x67, 0xee, 0x6b, 0xc9, 0x25, 0x5b, 0x51,
	0x87, 0x5d, 0x69, 0x25, 0xc5, 0x97, 0xc2, 0x0a, 0x25, 0xb3, 0x5a, 0x2e, 0x1d, 0xc1, 0xfd, 0x4d,
	0x7d, 0xfa, 0xfa, 0xd5, 0x87, 0x8b, 0xf4, 0x14, 0x88, 0x2f, 0x5d, 0x16, 0x9f, 0x28, 0x9b, 0x2b,
	0x76, 0x4d, 0xc5, 0x8c, 0x10, 0xe8, 0x4a, 0x35, 0xe3, 0x49, 0x34, 0x8e, 0x26, 0xc3, 0x1c, 0xd7,
	0x0e, 0xab, 0x94, 0xb6, 0xc9, 0xce, 0x38, 0x9a, 0xec, 0xe6, 0xb8, 0x4e, 0xbf, 0x76, 0xe0, 0xb0,
	0xd9, 0x6e, 0xd4, 0x52, 0x33, 0xee, 0xf6, 0x3f, 0x82, 0xbb, 0xf5, 0x07, 0x9b, 0x17, 0xc6, 0xd4,
	0x3a, 0xb1, 0xc7, 0xce, 0x1c, 0x44, 0x9e, 0x02, 0xe1, 0xf6, 0x8a, 0x6b, 0xc9, 0x2d, 0x15, 0xd2,
	0x72, 0x5d, 0x16, 0x8c, 0xa3, 0xf8, 0x30, 0x3f, 0x08, 0x95, 0xf3, 0x50, 0x20, 0x4f, 0xe0, 0x9e,
	0x51, 0xb7, 0xb9, 0x1d, 0xe4, 0xee, 0x21, 0xdc, 0x10, 0x6f, 0x60, 0x10, 0x6c, 0x24, 0xdd, 0x71,
	0x34, 0x89, 0xa7, 0xef, 0xb2, 0xbf, 0x1f, 0x61, 0xb6, 0x39, 0xa4, 0xbc, 0x8f, 0xab, 0xf3, 0x66,
	0x5a, 0x77, 0x5a, 0xd3, 0x1a, 0xc1, 0xa0, 0xb2, 0x15, 0x45, 0xbc, 0x87, 0x78, 0xbf, 0xb2, 0xd5,
	0x1b, 0x57, 0x3a, 0x85, 0x23, 0x53, 0x58, 0x3e, 0x9f, 0x0b, 0xcb, 0x69, 0xc1, 0x18, 0x37, 0xa6,
	0xe5, 0xaa, 0x8f, 0xe4, 0x64, 0xcd, 0x78, 0x8e, 0x84, 0xc6, 0xdf, 0x08, 0x06, 0x32, 0x08, 0x0f,
	0xbc, 0xb0, 0xf4, 0xc2, 0xe9, 0xaf, 0x08, 0x1e, 0xb4, 0xd2, 0xa8, 0xb3, 0x76, 0xde, 0x30, 0x70,
	0xa2, 0xa0, 0xe7, 0x13, 0xc0, 0x3c, 0xe2, 0xe9, 0xfb, 0xed, 0x0e, 0x66, 0x1d, 0x7f, 0x5e, 0x1f,
	0x43, 0x8e, 0x21, 0xf6, 0xd3, 0x62, 0x6a, 0x29, 0xc3, 0xcd, 0x01, 0x84, 0xce, 0x1c, 0x42, 0x1e,
	0xc3, 0xde, 0xfa, 0x12, 0x78, 0x4e, 0x07, 0x39, 0xbb, 0x01, 0xf5, 0xb4, 0x63, 0x88, 0x7d, 0xf8,
	0x9e, 0xd3, 0xf5, 0x3a, 0x08, 0x21, 0x21, 0xfd, 0x11, 0xc1, 0xd1, 0xa6, 0x73, 0xab, 0x66, 0xff,
	0xc9, 0xf8, 0x43, 0x00, 0x17, 0xd0, 0x2d, 0xdf, 0x43, 0x87, 0xf8, 0x76, 0x7f, 0xee, 0x84, 0x57,
	0xd7, 0x7e, 0x90, 0xe4, 0x5b, 0x04, 0x07, 0xad, 0x73, 0x7d, 0x29, 0x99, 0x8e, 0x3b, 0x93, 0x78,
	0x4a, 0xb7, 0xdc, 0xf2, 0x9f, 0x97, 0x23, 0xdf, 0x5f, 0x6f, 0xbf, 0xf0, 0x35, 0xf2, 0x3d, 0x82,
	0x43, 0x2b, 0x16, 0x9c, 0xaa, 0x92, 0xce, 0x8a, 0xa6, 0xa1, 0x13, 0x6c, 0xe8, 0xe3, 0xbf, 0x69,
	0x28, 0x64, 0x96, 0xef, 0xbb, 0xa3, 0xdf, 0x96, 0x2f, 0x8b, 0xd0, 0xcf, 0x65, 0x0f, 0x7f, 0x88,
	0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd0, 0x50, 0x06, 0x2c, 0x05, 0x00, 0x00,
}
