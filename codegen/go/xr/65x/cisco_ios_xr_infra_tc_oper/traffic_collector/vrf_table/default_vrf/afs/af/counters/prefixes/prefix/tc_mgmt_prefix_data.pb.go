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
// source: tc_mgmt_prefix_data.proto

package cisco_ios_xr_infra_tc_oper_traffic_collector_vrf_table_default_vrf_afs_af_counters_prefixes_prefix

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

type TcMgmtPrefixData_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipaddr               string   `protobuf:"bytes,2,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	Mask                 string   `protobuf:"bytes,3,opt,name=mask,proto3" json:"mask,omitempty"`
	Label                uint32   `protobuf:"varint,4,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcMgmtPrefixData_KEYS) Reset()         { *m = TcMgmtPrefixData_KEYS{} }
func (m *TcMgmtPrefixData_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcMgmtPrefixData_KEYS) ProtoMessage()    {}
func (*TcMgmtPrefixData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e39103d54d5e0d, []int{0}
}

func (m *TcMgmtPrefixData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcMgmtPrefixData_KEYS.Unmarshal(m, b)
}
func (m *TcMgmtPrefixData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcMgmtPrefixData_KEYS.Marshal(b, m, deterministic)
}
func (m *TcMgmtPrefixData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcMgmtPrefixData_KEYS.Merge(m, src)
}
func (m *TcMgmtPrefixData_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcMgmtPrefixData_KEYS.Size(m)
}
func (m *TcMgmtPrefixData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcMgmtPrefixData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcMgmtPrefixData_KEYS proto.InternalMessageInfo

func (m *TcMgmtPrefixData_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *TcMgmtPrefixData_KEYS) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

func (m *TcMgmtPrefixData_KEYS) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *TcMgmtPrefixData_KEYS) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

type TcMgmtCountHistory struct {
	EventStartTimestamp             uint64   `protobuf:"varint,1,opt,name=event_start_timestamp,json=eventStartTimestamp,proto3" json:"event_start_timestamp,omitempty"`
	EventEndTimestamp               uint64   `protobuf:"varint,2,opt,name=event_end_timestamp,json=eventEndTimestamp,proto3" json:"event_end_timestamp,omitempty"`
	TransmitNumberOfPacketsSwitched uint64   `protobuf:"varint,3,opt,name=transmit_number_of_packets_switched,json=transmitNumberOfPacketsSwitched,proto3" json:"transmit_number_of_packets_switched,omitempty"`
	TransmitNumberOfBytesSwitched   uint64   `protobuf:"varint,4,opt,name=transmit_number_of_bytes_switched,json=transmitNumberOfBytesSwitched,proto3" json:"transmit_number_of_bytes_switched,omitempty"`
	IsValid                         bool     `protobuf:"varint,5,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *TcMgmtCountHistory) Reset()         { *m = TcMgmtCountHistory{} }
func (m *TcMgmtCountHistory) String() string { return proto.CompactTextString(m) }
func (*TcMgmtCountHistory) ProtoMessage()    {}
func (*TcMgmtCountHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e39103d54d5e0d, []int{1}
}

func (m *TcMgmtCountHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcMgmtCountHistory.Unmarshal(m, b)
}
func (m *TcMgmtCountHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcMgmtCountHistory.Marshal(b, m, deterministic)
}
func (m *TcMgmtCountHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcMgmtCountHistory.Merge(m, src)
}
func (m *TcMgmtCountHistory) XXX_Size() int {
	return xxx_messageInfo_TcMgmtCountHistory.Size(m)
}
func (m *TcMgmtCountHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_TcMgmtCountHistory.DiscardUnknown(m)
}

var xxx_messageInfo_TcMgmtCountHistory proto.InternalMessageInfo

func (m *TcMgmtCountHistory) GetEventStartTimestamp() uint64 {
	if m != nil {
		return m.EventStartTimestamp
	}
	return 0
}

func (m *TcMgmtCountHistory) GetEventEndTimestamp() uint64 {
	if m != nil {
		return m.EventEndTimestamp
	}
	return 0
}

func (m *TcMgmtCountHistory) GetTransmitNumberOfPacketsSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfPacketsSwitched
	}
	return 0
}

func (m *TcMgmtCountHistory) GetTransmitNumberOfBytesSwitched() uint64 {
	if m != nil {
		return m.TransmitNumberOfBytesSwitched
	}
	return 0
}

func (m *TcMgmtCountHistory) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

type TcMgmtCountStats struct {
	TransmitPacketsPerSecondSwitched uint64                `protobuf:"varint,1,opt,name=transmit_packets_per_second_switched,json=transmitPacketsPerSecondSwitched,proto3" json:"transmit_packets_per_second_switched,omitempty"`
	TransmitBytesPerSecondSwitched   uint64                `protobuf:"varint,2,opt,name=transmit_bytes_per_second_switched,json=transmitBytesPerSecondSwitched,proto3" json:"transmit_bytes_per_second_switched,omitempty"`
	CountHistory                     []*TcMgmtCountHistory `protobuf:"bytes,3,rep,name=count_history,json=countHistory,proto3" json:"count_history,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}              `json:"-"`
	XXX_unrecognized                 []byte                `json:"-"`
	XXX_sizecache                    int32                 `json:"-"`
}

func (m *TcMgmtCountStats) Reset()         { *m = TcMgmtCountStats{} }
func (m *TcMgmtCountStats) String() string { return proto.CompactTextString(m) }
func (*TcMgmtCountStats) ProtoMessage()    {}
func (*TcMgmtCountStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e39103d54d5e0d, []int{2}
}

func (m *TcMgmtCountStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcMgmtCountStats.Unmarshal(m, b)
}
func (m *TcMgmtCountStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcMgmtCountStats.Marshal(b, m, deterministic)
}
func (m *TcMgmtCountStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcMgmtCountStats.Merge(m, src)
}
func (m *TcMgmtCountStats) XXX_Size() int {
	return xxx_messageInfo_TcMgmtCountStats.Size(m)
}
func (m *TcMgmtCountStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TcMgmtCountStats.DiscardUnknown(m)
}

var xxx_messageInfo_TcMgmtCountStats proto.InternalMessageInfo

func (m *TcMgmtCountStats) GetTransmitPacketsPerSecondSwitched() uint64 {
	if m != nil {
		return m.TransmitPacketsPerSecondSwitched
	}
	return 0
}

func (m *TcMgmtCountStats) GetTransmitBytesPerSecondSwitched() uint64 {
	if m != nil {
		return m.TransmitBytesPerSecondSwitched
	}
	return 0
}

func (m *TcMgmtCountStats) GetCountHistory() []*TcMgmtCountHistory {
	if m != nil {
		return m.CountHistory
	}
	return nil
}

type TcMgmtPrefixData struct {
	Prefix                         string            `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	LabelXr                        uint32            `protobuf:"varint,51,opt,name=label_xr,json=labelXr,proto3" json:"label_xr,omitempty"`
	LdpLabel                       uint32            `protobuf:"varint,52,opt,name=ldp_label,json=ldpLabel,proto3" json:"ldp_label,omitempty"`
	IsActive                       bool              `protobuf:"varint,53,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	BaseCounterStatistics          *TcMgmtCountStats `protobuf:"bytes,54,opt,name=base_counter_statistics,json=baseCounterStatistics,proto3" json:"base_counter_statistics,omitempty"`
	TrafficMatrixCounterStatistics *TcMgmtCountStats `protobuf:"bytes,55,opt,name=traffic_matrix_counter_statistics,json=trafficMatrixCounterStatistics,proto3" json:"traffic_matrix_counter_statistics,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}          `json:"-"`
	XXX_unrecognized               []byte            `json:"-"`
	XXX_sizecache                  int32             `json:"-"`
}

func (m *TcMgmtPrefixData) Reset()         { *m = TcMgmtPrefixData{} }
func (m *TcMgmtPrefixData) String() string { return proto.CompactTextString(m) }
func (*TcMgmtPrefixData) ProtoMessage()    {}
func (*TcMgmtPrefixData) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e39103d54d5e0d, []int{3}
}

func (m *TcMgmtPrefixData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcMgmtPrefixData.Unmarshal(m, b)
}
func (m *TcMgmtPrefixData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcMgmtPrefixData.Marshal(b, m, deterministic)
}
func (m *TcMgmtPrefixData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcMgmtPrefixData.Merge(m, src)
}
func (m *TcMgmtPrefixData) XXX_Size() int {
	return xxx_messageInfo_TcMgmtPrefixData.Size(m)
}
func (m *TcMgmtPrefixData) XXX_DiscardUnknown() {
	xxx_messageInfo_TcMgmtPrefixData.DiscardUnknown(m)
}

var xxx_messageInfo_TcMgmtPrefixData proto.InternalMessageInfo

func (m *TcMgmtPrefixData) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *TcMgmtPrefixData) GetLabelXr() uint32 {
	if m != nil {
		return m.LabelXr
	}
	return 0
}

func (m *TcMgmtPrefixData) GetLdpLabel() uint32 {
	if m != nil {
		return m.LdpLabel
	}
	return 0
}

func (m *TcMgmtPrefixData) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *TcMgmtPrefixData) GetBaseCounterStatistics() *TcMgmtCountStats {
	if m != nil {
		return m.BaseCounterStatistics
	}
	return nil
}

func (m *TcMgmtPrefixData) GetTrafficMatrixCounterStatistics() *TcMgmtCountStats {
	if m != nil {
		return m.TrafficMatrixCounterStatistics
	}
	return nil
}

func init() {
	proto.RegisterType((*TcMgmtPrefixData_KEYS)(nil), "cisco_ios_xr_infra_tc_oper.traffic_collector.vrf_table.default_vrf.afs.af.counters.prefixes.prefix.tc_mgmt_prefix_data_KEYS")
	proto.RegisterType((*TcMgmtCountHistory)(nil), "cisco_ios_xr_infra_tc_oper.traffic_collector.vrf_table.default_vrf.afs.af.counters.prefixes.prefix.tc_mgmt_count_history")
	proto.RegisterType((*TcMgmtCountStats)(nil), "cisco_ios_xr_infra_tc_oper.traffic_collector.vrf_table.default_vrf.afs.af.counters.prefixes.prefix.tc_mgmt_count_stats")
	proto.RegisterType((*TcMgmtPrefixData)(nil), "cisco_ios_xr_infra_tc_oper.traffic_collector.vrf_table.default_vrf.afs.af.counters.prefixes.prefix.tc_mgmt_prefix_data")
}

func init() { proto.RegisterFile("tc_mgmt_prefix_data.proto", fileDescriptor_43e39103d54d5e0d) }

var fileDescriptor_43e39103d54d5e0d = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x24, 0x69, 0x9b, 0xba, 0x5f, 0x17, 0x9f, 0x4b, 0xe9, 0x54, 0x88, 0x32, 0x04, 0x16,
	0x59, 0xcd, 0x22, 0xe5, 0x67, 0x0d, 0xa8, 0x52, 0x05, 0xa5, 0x54, 0x13, 0x84, 0x60, 0x75, 0xe5,
	0xf1, 0xd8, 0xd4, 0xea, 0xfc, 0xc9, 0xbe, 0x49, 0xd3, 0x97, 0xe1, 0x11, 0xe0, 0x1d, 0x78, 0x04,
	0xde, 0x86, 0x1d, 0x9a, 0xeb, 0x4c, 0x9a, 0xd2, 0x6c, 0x81, 0xdd, 0x1c, 0xdf, 0xe3, 0x73, 0xcf,
	0xbd, 0x3e, 0x1a, 0xb6, 0x8f, 0x12, 0x8a, 0xcf, 0x05, 0x42, 0x6d, 0x95, 0x36, 0x33, 0xc8, 0x04,
	0x8a, 0xb8, 0xb6, 0x15, 0x56, 0x3c, 0x95, 0xc6, 0xc9, 0x0a, 0x4c, 0xe5, 0x60, 0x66, 0xc1, 0x94,
	0xda, 0x0a, 0x40, 0x09, 0x55, 0xad, 0x6c, 0x8c, 0x56, 0x68, 0x6d, 0x24, 0xc8, 0x2a, 0xcf, 0x95,
	0xc4, 0xca, 0xc6, 0x53, 0xab, 0x01, 0x45, 0x9a, 0xab, 0x38, 0x53, 0x5a, 0x4c, 0x72, 0x84, 0xa9,
	0xd5, 0xb1, 0xd0, 0x2e, 0x16, 0x3a, 0x96, 0xd5, 0xa4, 0x44, 0x65, 0x5d, 0xec, 0xbb, 0xa8, 0xf6,
	0x63, 0x30, 0x61, 0xe1, 0x0a, 0x03, 0xf0, 0xe6, 0xe8, 0xd3, 0x98, 0xef, 0xb1, 0x0d, 0xa1, 0xa1,
	0x14, 0x85, 0x0a, 0x83, 0x28, 0x18, 0x6e, 0x26, 0xeb, 0x42, 0x9f, 0x8a, 0x42, 0xf1, 0xbb, 0x6c,
	0xdd, 0xd4, 0x22, 0xcb, 0x6c, 0xd8, 0xf1, 0xe7, 0x1e, 0x71, 0xce, 0x7a, 0x85, 0x70, 0x17, 0x61,
	0x97, 0x4e, 0xe9, 0x9b, 0xdf, 0x61, 0x6b, 0xb9, 0x48, 0x55, 0x1e, 0xf6, 0xa2, 0x60, 0xb8, 0x9d,
	0x78, 0x30, 0xf8, 0xda, 0x61, 0xbb, 0x6d, 0x5f, 0xf2, 0x06, 0xe7, 0xc6, 0x61, 0x65, 0xaf, 0xf8,
	0x88, 0xed, 0xaa, 0xa9, 0x2a, 0x11, 0x1c, 0x0a, 0x8b, 0x80, 0xa6, 0x50, 0x0e, 0x45, 0x51, 0x93,
	0x85, 0x5e, 0xb2, 0x43, 0xc5, 0x71, 0x53, 0x7b, 0xdf, 0x96, 0x78, 0xcc, 0xfc, 0x31, 0xa8, 0x32,
	0x5b, 0xba, 0xd1, 0xa1, 0x1b, 0xff, 0x53, 0xe9, 0xa8, 0xcc, 0xae, 0xf9, 0x27, 0xec, 0x11, 0x5a,
	0x51, 0xba, 0xc2, 0x20, 0x94, 0x93, 0x22, 0x55, 0x16, 0x2a, 0x0d, 0xb5, 0x90, 0x17, 0x0a, 0x1d,
	0xb8, 0x4b, 0x83, 0xf2, 0x5c, 0x65, 0x34, 0x46, 0x2f, 0x79, 0xd0, 0x52, 0x4f, 0x89, 0xf9, 0x4e,
	0x9f, 0x79, 0xde, 0x78, 0x4e, 0xe3, 0xc7, 0xec, 0xe1, 0x0a, 0xb5, 0xf4, 0x0a, 0xd5, 0x92, 0x56,
	0x8f, 0xb4, 0xee, 0xff, 0xae, 0xf5, 0xb2, 0x61, 0x2d, 0x94, 0xf6, 0x59, 0xdf, 0x38, 0x98, 0x8a,
	0xdc, 0x64, 0xe1, 0x5a, 0x14, 0x0c, 0xfb, 0xc9, 0x86, 0x71, 0x1f, 0x1a, 0x38, 0xf8, 0xd1, 0x61,
	0x3b, 0x37, 0x17, 0xe6, 0x50, 0xa0, 0xe3, 0xa7, 0xec, 0xf1, 0xa2, 0x79, 0x3b, 0x40, 0xad, 0x2c,
	0x38, 0x25, 0xab, 0x32, 0xbb, 0xee, 0xef, 0xb7, 0x17, 0xb5, 0xdc, 0xf9, 0x0c, 0x67, 0xca, 0x8e,
	0x89, 0xb8, 0xb0, 0xf0, 0x9a, 0x0d, 0x16, 0x7a, 0x7e, 0x84, 0x55, 0x6a, 0x7e, 0xb3, 0x07, 0x2d,
	0x93, 0xa6, 0xb8, 0xad, 0xf5, 0x25, 0x60, 0xdb, 0x37, 0x1e, 0x37, 0xec, 0x46, 0xdd, 0xe1, 0xd6,
	0xe8, 0x2a, 0xfe, 0xf3, 0xc1, 0x8e, 0x57, 0xa6, 0x2b, 0xf9, 0x8f, 0xe0, 0xb1, 0x47, 0x83, 0x9f,
	0xdd, 0xeb, 0xa5, 0x2e, 0xa5, 0xbf, 0xc9, 0xb7, 0x87, 0xe1, 0xc8, 0xe7, 0xdb, 0xa3, 0xe6, 0x7d,
	0x28, 0xbe, 0x30, 0xb3, 0xe1, 0x21, 0xc5, 0x79, 0x83, 0xf0, 0x47, 0xcb, 0xef, 0xb1, 0xcd, 0x3c,
	0xab, 0xc1, 0x47, 0xfd, 0x09, 0xd5, 0xfa, 0x79, 0x56, 0x9f, 0x34, 0xb8, 0x29, 0x1a, 0x07, 0x42,
	0xa2, 0x99, 0xaa, 0xf0, 0x29, 0x3d, 0x6c, 0xdf, 0xb8, 0x17, 0x84, 0xf9, 0xb7, 0x80, 0xed, 0xa5,
	0xc2, 0x29, 0x98, 0x8f, 0x42, 0x0f, 0x6b, 0x1c, 0x1a, 0xe9, 0xc2, 0x67, 0x51, 0x30, 0xdc, 0x1a,
	0x5d, 0xfe, 0xfd, 0x7d, 0x51, 0xb8, 0x92, 0xdd, 0xc6, 0xd7, 0x2b, 0x7f, 0x63, 0xbc, 0x70, 0xc5,
	0xbf, 0x07, 0x94, 0x78, 0x6a, 0x5b, 0x08, 0xb4, 0x66, 0xb6, 0xca, 0xfb, 0xf3, 0x7f, 0xeb, 0xfd,
	0x60, 0x2e, 0xfe, 0x96, 0x0c, 0xde, 0x1a, 0x22, 0x5d, 0xa7, 0x7f, 0xec, 0xe1, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0xcc, 0x1b, 0xc9, 0x80, 0x05, 0x00, 0x00,
}
