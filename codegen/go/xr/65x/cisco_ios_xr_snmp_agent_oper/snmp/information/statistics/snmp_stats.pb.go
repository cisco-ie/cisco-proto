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
// source: snmp_stats.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_statistics

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

type SnmpStats_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpStats_KEYS) Reset()         { *m = SnmpStats_KEYS{} }
func (m *SnmpStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpStats_KEYS) ProtoMessage()    {}
func (*SnmpStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d7742d1ce4d167, []int{0}
}

func (m *SnmpStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpStats_KEYS.Unmarshal(m, b)
}
func (m *SnmpStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpStats_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpStats_KEYS.Merge(m, src)
}
func (m *SnmpStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpStats_KEYS.Size(m)
}
func (m *SnmpStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpStats_KEYS proto.InternalMessageInfo

type SnmpStats struct {
	PacketsReceived           uint32   `protobuf:"varint,50,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	BadVersionsReceived       uint32   `protobuf:"varint,51,opt,name=bad_versions_received,json=badVersionsReceived,proto3" json:"bad_versions_received,omitempty"`
	BadCommunityNamesReceived uint32   `protobuf:"varint,52,opt,name=bad_community_names_received,json=badCommunityNamesReceived,proto3" json:"bad_community_names_received,omitempty"`
	BadCommunityUsesReceived  uint32   `protobuf:"varint,53,opt,name=bad_community_uses_received,json=badCommunityUsesReceived,proto3" json:"bad_community_uses_received,omitempty"`
	AsnParseErrorsReceived    uint32   `protobuf:"varint,54,opt,name=asn_parse_errors_received,json=asnParseErrorsReceived,proto3" json:"asn_parse_errors_received,omitempty"`
	SilentDropCount           uint32   `protobuf:"varint,55,opt,name=silent_drop_count,json=silentDropCount,proto3" json:"silent_drop_count,omitempty"`
	ProxyDropCount            uint32   `protobuf:"varint,56,opt,name=proxy_drop_count,json=proxyDropCount,proto3" json:"proxy_drop_count,omitempty"`
	TooBigPacketReceived      uint32   `protobuf:"varint,57,opt,name=too_big_packet_received,json=tooBigPacketReceived,proto3" json:"too_big_packet_received,omitempty"`
	MaxPacketSize             uint32   `protobuf:"varint,58,opt,name=max_packet_size,json=maxPacketSize,proto3" json:"max_packet_size,omitempty"`
	NoSuchNamesReceived       uint32   `protobuf:"varint,59,opt,name=no_such_names_received,json=noSuchNamesReceived,proto3" json:"no_such_names_received,omitempty"`
	BadValuesReceived         uint32   `protobuf:"varint,60,opt,name=bad_values_received,json=badValuesReceived,proto3" json:"bad_values_received,omitempty"`
	ReadOnlyReceived          uint32   `protobuf:"varint,61,opt,name=read_only_received,json=readOnlyReceived,proto3" json:"read_only_received,omitempty"`
	TotalGeneralErrors        uint32   `protobuf:"varint,62,opt,name=total_general_errors,json=totalGeneralErrors,proto3" json:"total_general_errors,omitempty"`
	TotalRequestedVariables   uint32   `protobuf:"varint,63,opt,name=total_requested_variables,json=totalRequestedVariables,proto3" json:"total_requested_variables,omitempty"`
	TotalSetVariablesReceived uint32   `protobuf:"varint,64,opt,name=total_set_variables_received,json=totalSetVariablesReceived,proto3" json:"total_set_variables_received,omitempty"`
	GetRequestsReceived       uint32   `protobuf:"varint,65,opt,name=get_requests_received,json=getRequestsReceived,proto3" json:"get_requests_received,omitempty"`
	GetNextRequestsReceived   uint32   `protobuf:"varint,66,opt,name=get_next_requests_received,json=getNextRequestsReceived,proto3" json:"get_next_requests_received,omitempty"`
	SetRequestsReceived       uint32   `protobuf:"varint,67,opt,name=set_requests_received,json=setRequestsReceived,proto3" json:"set_requests_received,omitempty"`
	GetResponsesReceived      uint32   `protobuf:"varint,68,opt,name=get_responses_received,json=getResponsesReceived,proto3" json:"get_responses_received,omitempty"`
	TrapsReceived             uint32   `protobuf:"varint,69,opt,name=traps_received,json=trapsReceived,proto3" json:"traps_received,omitempty"`
	TotalPacketsSent          uint32   `protobuf:"varint,70,opt,name=total_packets_sent,json=totalPacketsSent,proto3" json:"total_packets_sent,omitempty"`
	TooBigPacketsSent         uint32   `protobuf:"varint,71,opt,name=too_big_packets_sent,json=tooBigPacketsSent,proto3" json:"too_big_packets_sent,omitempty"`
	NoSuchNamesSent           uint32   `protobuf:"varint,72,opt,name=no_such_names_sent,json=noSuchNamesSent,proto3" json:"no_such_names_sent,omitempty"`
	BadValuesSent             uint32   `protobuf:"varint,73,opt,name=bad_values_sent,json=badValuesSent,proto3" json:"bad_values_sent,omitempty"`
	GeneralErrorsSent         uint32   `protobuf:"varint,74,opt,name=general_errors_sent,json=generalErrorsSent,proto3" json:"general_errors_sent,omitempty"`
	GetRequestsSent           uint32   `protobuf:"varint,75,opt,name=get_requests_sent,json=getRequestsSent,proto3" json:"get_requests_sent,omitempty"`
	GetNextRequestSent        uint32   `protobuf:"varint,76,opt,name=get_next_request_sent,json=getNextRequestSent,proto3" json:"get_next_request_sent,omitempty"`
	SetRequestsSent           uint32   `protobuf:"varint,77,opt,name=set_requests_sent,json=setRequestsSent,proto3" json:"set_requests_sent,omitempty"`
	GetResponsesSent          uint32   `protobuf:"varint,78,opt,name=get_responses_sent,json=getResponsesSent,proto3" json:"get_responses_sent,omitempty"`
	TrapsSent                 uint32   `protobuf:"varint,79,opt,name=traps_sent,json=trapsSent,proto3" json:"traps_sent,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *SnmpStats) Reset()         { *m = SnmpStats{} }
func (m *SnmpStats) String() string { return proto.CompactTextString(m) }
func (*SnmpStats) ProtoMessage()    {}
func (*SnmpStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3d7742d1ce4d167, []int{1}
}

func (m *SnmpStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpStats.Unmarshal(m, b)
}
func (m *SnmpStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpStats.Marshal(b, m, deterministic)
}
func (m *SnmpStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpStats.Merge(m, src)
}
func (m *SnmpStats) XXX_Size() int {
	return xxx_messageInfo_SnmpStats.Size(m)
}
func (m *SnmpStats) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpStats.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpStats proto.InternalMessageInfo

func (m *SnmpStats) GetPacketsReceived() uint32 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *SnmpStats) GetBadVersionsReceived() uint32 {
	if m != nil {
		return m.BadVersionsReceived
	}
	return 0
}

func (m *SnmpStats) GetBadCommunityNamesReceived() uint32 {
	if m != nil {
		return m.BadCommunityNamesReceived
	}
	return 0
}

func (m *SnmpStats) GetBadCommunityUsesReceived() uint32 {
	if m != nil {
		return m.BadCommunityUsesReceived
	}
	return 0
}

func (m *SnmpStats) GetAsnParseErrorsReceived() uint32 {
	if m != nil {
		return m.AsnParseErrorsReceived
	}
	return 0
}

func (m *SnmpStats) GetSilentDropCount() uint32 {
	if m != nil {
		return m.SilentDropCount
	}
	return 0
}

func (m *SnmpStats) GetProxyDropCount() uint32 {
	if m != nil {
		return m.ProxyDropCount
	}
	return 0
}

func (m *SnmpStats) GetTooBigPacketReceived() uint32 {
	if m != nil {
		return m.TooBigPacketReceived
	}
	return 0
}

func (m *SnmpStats) GetMaxPacketSize() uint32 {
	if m != nil {
		return m.MaxPacketSize
	}
	return 0
}

func (m *SnmpStats) GetNoSuchNamesReceived() uint32 {
	if m != nil {
		return m.NoSuchNamesReceived
	}
	return 0
}

func (m *SnmpStats) GetBadValuesReceived() uint32 {
	if m != nil {
		return m.BadValuesReceived
	}
	return 0
}

func (m *SnmpStats) GetReadOnlyReceived() uint32 {
	if m != nil {
		return m.ReadOnlyReceived
	}
	return 0
}

func (m *SnmpStats) GetTotalGeneralErrors() uint32 {
	if m != nil {
		return m.TotalGeneralErrors
	}
	return 0
}

func (m *SnmpStats) GetTotalRequestedVariables() uint32 {
	if m != nil {
		return m.TotalRequestedVariables
	}
	return 0
}

func (m *SnmpStats) GetTotalSetVariablesReceived() uint32 {
	if m != nil {
		return m.TotalSetVariablesReceived
	}
	return 0
}

func (m *SnmpStats) GetGetRequestsReceived() uint32 {
	if m != nil {
		return m.GetRequestsReceived
	}
	return 0
}

func (m *SnmpStats) GetGetNextRequestsReceived() uint32 {
	if m != nil {
		return m.GetNextRequestsReceived
	}
	return 0
}

func (m *SnmpStats) GetSetRequestsReceived() uint32 {
	if m != nil {
		return m.SetRequestsReceived
	}
	return 0
}

func (m *SnmpStats) GetGetResponsesReceived() uint32 {
	if m != nil {
		return m.GetResponsesReceived
	}
	return 0
}

func (m *SnmpStats) GetTrapsReceived() uint32 {
	if m != nil {
		return m.TrapsReceived
	}
	return 0
}

func (m *SnmpStats) GetTotalPacketsSent() uint32 {
	if m != nil {
		return m.TotalPacketsSent
	}
	return 0
}

func (m *SnmpStats) GetTooBigPacketsSent() uint32 {
	if m != nil {
		return m.TooBigPacketsSent
	}
	return 0
}

func (m *SnmpStats) GetNoSuchNamesSent() uint32 {
	if m != nil {
		return m.NoSuchNamesSent
	}
	return 0
}

func (m *SnmpStats) GetBadValuesSent() uint32 {
	if m != nil {
		return m.BadValuesSent
	}
	return 0
}

func (m *SnmpStats) GetGeneralErrorsSent() uint32 {
	if m != nil {
		return m.GeneralErrorsSent
	}
	return 0
}

func (m *SnmpStats) GetGetRequestsSent() uint32 {
	if m != nil {
		return m.GetRequestsSent
	}
	return 0
}

func (m *SnmpStats) GetGetNextRequestSent() uint32 {
	if m != nil {
		return m.GetNextRequestSent
	}
	return 0
}

func (m *SnmpStats) GetSetRequestsSent() uint32 {
	if m != nil {
		return m.SetRequestsSent
	}
	return 0
}

func (m *SnmpStats) GetGetResponsesSent() uint32 {
	if m != nil {
		return m.GetResponsesSent
	}
	return 0
}

func (m *SnmpStats) GetTrapsSent() uint32 {
	if m != nil {
		return m.TrapsSent
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpStats_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.statistics.snmp_stats_KEYS")
	proto.RegisterType((*SnmpStats)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.statistics.snmp_stats")
}

func init() { proto.RegisterFile("snmp_stats.proto", fileDescriptor_e3d7742d1ce4d167) }

var fileDescriptor_e3d7742d1ce4d167 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0xeb, 0x4f, 0x13, 0x4d,
	0x14, 0xc6, 0xf3, 0x7e, 0x79, 0x93, 0xf7, 0xbc, 0x81, 0xb6, 0xcb, 0xad, 0x78, 0x49, 0x0c, 0x89,
	0x04, 0x91, 0x54, 0xe5, 0xa2, 0x5c, 0x44, 0x94, 0x8b, 0xa8, 0x28, 0x10, 0x1a, 0x49, 0xfc, 0x34,
	0x99, 0x6e, 0xc7, 0x65, 0x62, 0x77, 0x66, 0x9d, 0x99, 0x25, 0x2d, 0xff, 0xaa, 0xff, 0x8c, 0x99,
	0x73, 0xb6, 0xbb, 0xb3, 0xc0, 0xc7, 0xce, 0xf3, 0x7b, 0x66, 0x4f, 0xce, 0xf3, 0xec, 0x16, 0x9a,
	0x56, 0xa5, 0x19, 0xb3, 0x8e, 0x3b, 0xdb, 0xc9, 0x8c, 0x76, 0x3a, 0xda, 0x8c, 0xa5, 0x8d, 0x35,
	0x93, 0xda, 0xb2, 0xa1, 0x61, 0x28, 0xf3, 0x44, 0x28, 0xc7, 0x74, 0x26, 0x4c, 0xc7, 0xff, 0xee,
	0x48, 0xf5, 0x53, 0x9b, 0x94, 0x3b, 0xa9, 0x55, 0xc7, 0x5b, 0xa5, 0x75, 0x32, 0xb6, 0x0b, 0x2d,
	0x68, 0x54, 0xb7, 0xb1, 0x93, 0xa3, 0x1f, 0xdd, 0x85, 0x3f, 0xff, 0x03, 0x54, 0x67, 0xd1, 0x33,
	0x68, 0x66, 0x3c, 0xfe, 0x25, 0x9c, 0x65, 0x46, 0xc4, 0x42, 0x5e, 0x8b, 0x7e, 0x7b, 0xf5, 0xc9,
	0x3f, 0x4b, 0x13, 0x17, 0x8d, 0xe2, 0xfc, 0xa2, 0x38, 0x8e, 0x56, 0x61, 0xa6, 0xc7, 0xfb, 0xec,
	0x5a, 0x18, 0x2b, 0xb5, 0x0a, 0xf8, 0x35, 0xe4, 0xa7, 0x7a, 0xbc, 0x7f, 0x59, 0x68, 0xa5, 0x67,
	0x0f, 0x1e, 0x79, 0x4f, 0xac, 0xd3, 0x34, 0x57, 0xd2, 0x8d, 0x98, 0xe2, 0xa9, 0x08, 0xac, 0xeb,
	0x68, 0x9d, 0xef, 0xf1, 0xfe, 0xc1, 0x18, 0x39, 0xf5, 0x44, 0x79, 0xc1, 0x2e, 0x3c, 0xac, 0x5f,
	0x90, 0xdb, 0xd0, 0xbf, 0x81, 0xfe, 0x76, 0xe8, 0xff, 0x6e, 0x03, 0xfb, 0x16, 0xcc, 0x73, 0xab,
	0x58, 0xc6, 0x8d, 0x15, 0x4c, 0x18, 0xa3, 0x4d, 0x60, 0x7e, 0x8d, 0xe6, 0x59, 0x6e, 0xd5, 0xb9,
	0xd7, 0x8f, 0x50, 0x2e, 0xad, 0xcb, 0xd0, 0xb2, 0x72, 0xe0, 0xd7, 0xdc, 0x37, 0x3a, 0x63, 0xb1,
	0xce, 0x95, 0x6b, 0xbf, 0xa1, 0xd5, 0x90, 0x70, 0x68, 0x74, 0x76, 0xe0, 0x8f, 0xa3, 0x25, 0x68,
	0x66, 0x46, 0x0f, 0x47, 0x21, 0xba, 0x89, 0xe8, 0x24, 0x9e, 0x57, 0xe4, 0x06, 0xcc, 0x39, 0xad,
	0x59, 0x4f, 0x26, 0x8c, 0xf6, 0x5b, 0x8d, 0xb3, 0x85, 0x86, 0x69, 0xa7, 0xf5, 0xbe, 0x4c, 0xce,
	0x51, 0x2c, 0x87, 0x59, 0x84, 0x46, 0xca, 0x87, 0x63, 0x8b, 0x95, 0x37, 0xa2, 0xbd, 0x8d, 0xf8,
	0x44, 0xca, 0x87, 0xc4, 0x76, 0xe5, 0x8d, 0x88, 0xd6, 0x60, 0x56, 0x69, 0x66, 0xf3, 0xf8, 0xea,
	0xf6, 0xa6, 0x77, 0x28, 0x24, 0xa5, 0xbb, 0x79, 0x7c, 0x55, 0xdf, 0x71, 0x07, 0xa6, 0x30, 0x58,
	0x3e, 0xc8, 0x43, 0xc7, 0x5b, 0x74, 0xb4, 0x7c, 0xac, 0xa8, 0x94, 0xfc, 0x0a, 0x44, 0x46, 0xf0,
	0x3e, 0xd3, 0x6a, 0x30, 0xaa, 0xf0, 0x5d, 0xc4, 0x9b, 0x5e, 0x39, 0x53, 0x83, 0x51, 0x49, 0xbf,
	0x84, 0x69, 0xa7, 0x1d, 0x1f, 0xb0, 0x44, 0x28, 0x61, 0xf8, 0xa0, 0x88, 0xa1, 0xfd, 0x0e, 0xf9,
	0x08, 0xb5, 0x63, 0x92, 0x28, 0x81, 0x68, 0x1b, 0xe6, 0xc9, 0x61, 0xc4, 0xef, 0x5c, 0x58, 0x27,
	0xfc, 0x6c, 0x46, 0xf2, 0xde, 0x40, 0xd8, 0xf6, 0x1e, 0xda, 0xe6, 0x10, 0xb8, 0x18, 0xeb, 0x97,
	0x63, 0xd9, 0x17, 0x8e, 0xbc, 0x56, 0xb8, 0xca, 0x55, 0x4d, 0xf9, 0x9e, 0x0a, 0x87, 0x4c, 0x57,
	0xb8, 0xd2, 0x18, 0xb6, 0x3c, 0xc1, 0x54, 0xf0, 0xea, 0xc0, 0xf9, 0x81, 0x16, 0x98, 0xf8, 0x54,
	0x48, 0x2b, 0x3d, 0x3b, 0xf0, 0xc0, 0x7b, 0x94, 0x18, 0xde, 0x67, 0xdc, 0xa7, 0x89, 0x13, 0xe1,
	0x4e, 0xc5, 0xf0, 0xae, 0x79, 0x15, 0x66, 0xec, 0xbd, 0x0f, 0x3c, 0xa0, 0x07, 0xda, 0x7b, 0x1e,
	0xb8, 0x0e, 0xb3, 0x34, 0xa4, 0xcd, 0xb4, 0xaa, 0xbd, 0x10, 0x87, 0x54, 0x22, 0x9c, 0xb2, 0x10,
	0x4b, 0xd7, 0x53, 0x98, 0x74, 0x86, 0x67, 0x01, 0x7d, 0x44, 0x1d, 0xc2, 0xd3, 0x30, 0x5e, 0x5a,
	0xe1, 0xf8, 0xc3, 0x60, 0x85, 0x72, 0xed, 0x8f, 0x14, 0x2f, 0x2a, 0x54, 0x38, 0xdb, 0x15, 0xca,
	0x45, 0x2f, 0x7c, 0xbc, 0x61, 0xa1, 0x0b, 0xfe, 0x98, 0xda, 0x13, 0xb6, 0x99, 0x0c, 0xcf, 0x21,
	0xaa, 0x57, 0x14, 0xf1, 0x4f, 0xf4, 0x62, 0x05, 0xf5, 0x44, 0x78, 0x11, 0x1a, 0x41, 0x35, 0x91,
	0xfc, 0x4c, 0x33, 0x97, 0xb5, 0x44, 0xae, 0x03, 0x53, 0xf5, 0x7a, 0x11, 0xfb, 0x85, 0x86, 0x48,
	0xc2, 0x7a, 0x21, 0xbf, 0x0c, 0xad, 0x5a, 0xca, 0x48, 0x9f, 0xd0, 0x0c, 0x41, 0xc2, 0xc8, 0xbe,
	0xa2, 0x46, 0x84, 0xe9, 0x12, 0xff, 0x95, 0x1a, 0x5c, 0x0f, 0x76, 0x7c, 0xbd, 0xbd, 0x73, 0xfd,
	0xb7, 0xe2, 0xdb, 0x71, 0xeb, 0xfa, 0x15, 0x88, 0xea, 0x59, 0x22, 0x7c, 0x4a, 0xeb, 0x0e, 0x73,
	0x44, 0xfa, 0x31, 0x00, 0x65, 0x88, 0xd4, 0x19, 0x52, 0xff, 0xe1, 0x89, 0x97, 0x7b, 0xff, 0xe2,
	0x3f, 0xc6, 0xda, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x07, 0x2d, 0xa9, 0x45, 0x06, 0x00,
	0x00,
}
