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
// source: rip_show_statistics_bd.proto

package cisco_ios_xr_ip_rip_oper_rip_default_vrf_statistics

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

type RipShowStatisticsBd_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RipShowStatisticsBd_KEYS) Reset()         { *m = RipShowStatisticsBd_KEYS{} }
func (m *RipShowStatisticsBd_KEYS) String() string { return proto.CompactTextString(m) }
func (*RipShowStatisticsBd_KEYS) ProtoMessage()    {}
func (*RipShowStatisticsBd_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_66227cbf5e51e264, []int{0}
}

func (m *RipShowStatisticsBd_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowStatisticsBd_KEYS.Unmarshal(m, b)
}
func (m *RipShowStatisticsBd_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowStatisticsBd_KEYS.Marshal(b, m, deterministic)
}
func (m *RipShowStatisticsBd_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowStatisticsBd_KEYS.Merge(m, src)
}
func (m *RipShowStatisticsBd_KEYS) XXX_Size() int {
	return xxx_messageInfo_RipShowStatisticsBd_KEYS.Size(m)
}
func (m *RipShowStatisticsBd_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowStatisticsBd_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowStatisticsBd_KEYS proto.InternalMessageInfo

type RipShowStatisticsBd struct {
	ReceivedPackets        uint32   `protobuf:"varint,50,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	DiscardedPackets       uint32   `protobuf:"varint,51,opt,name=discarded_packets,json=discardedPackets,proto3" json:"discarded_packets,omitempty"`
	DiscardedRoutes        uint32   `protobuf:"varint,52,opt,name=discarded_routes,json=discardedRoutes,proto3" json:"discarded_routes,omitempty"`
	StandbyPacketsReceived uint32   `protobuf:"varint,53,opt,name=standby_packets_received,json=standbyPacketsReceived,proto3" json:"standby_packets_received,omitempty"`
	SentMessages           uint32   `protobuf:"varint,54,opt,name=sent_messages,json=sentMessages,proto3" json:"sent_messages,omitempty"`
	SentMessageFailures    uint32   `protobuf:"varint,55,opt,name=sent_message_failures,json=sentMessageFailures,proto3" json:"sent_message_failures,omitempty"`
	QueryResponses         uint32   `protobuf:"varint,56,opt,name=query_responses,json=queryResponses,proto3" json:"query_responses,omitempty"`
	PeriodicUpdates        uint32   `protobuf:"varint,57,opt,name=periodic_updates,json=periodicUpdates,proto3" json:"periodic_updates,omitempty"`
	RouteCount             uint32   `protobuf:"varint,58,opt,name=route_count,json=routeCount,proto3" json:"route_count,omitempty"`
	PathCount              uint32   `protobuf:"varint,59,opt,name=path_count,json=pathCount,proto3" json:"path_count,omitempty"`
	RouteMallocFailures    uint32   `protobuf:"varint,60,opt,name=route_malloc_failures,json=routeMallocFailures,proto3" json:"route_malloc_failures,omitempty"`
	PathMallocFailures     uint32   `protobuf:"varint,61,opt,name=path_malloc_failures,json=pathMallocFailures,proto3" json:"path_malloc_failures,omitempty"`
	RibUpdates             uint32   `protobuf:"varint,62,opt,name=rib_updates,json=ribUpdates,proto3" json:"rib_updates,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *RipShowStatisticsBd) Reset()         { *m = RipShowStatisticsBd{} }
func (m *RipShowStatisticsBd) String() string { return proto.CompactTextString(m) }
func (*RipShowStatisticsBd) ProtoMessage()    {}
func (*RipShowStatisticsBd) Descriptor() ([]byte, []int) {
	return fileDescriptor_66227cbf5e51e264, []int{1}
}

func (m *RipShowStatisticsBd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RipShowStatisticsBd.Unmarshal(m, b)
}
func (m *RipShowStatisticsBd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RipShowStatisticsBd.Marshal(b, m, deterministic)
}
func (m *RipShowStatisticsBd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RipShowStatisticsBd.Merge(m, src)
}
func (m *RipShowStatisticsBd) XXX_Size() int {
	return xxx_messageInfo_RipShowStatisticsBd.Size(m)
}
func (m *RipShowStatisticsBd) XXX_DiscardUnknown() {
	xxx_messageInfo_RipShowStatisticsBd.DiscardUnknown(m)
}

var xxx_messageInfo_RipShowStatisticsBd proto.InternalMessageInfo

func (m *RipShowStatisticsBd) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func (m *RipShowStatisticsBd) GetDiscardedPackets() uint32 {
	if m != nil {
		return m.DiscardedPackets
	}
	return 0
}

func (m *RipShowStatisticsBd) GetDiscardedRoutes() uint32 {
	if m != nil {
		return m.DiscardedRoutes
	}
	return 0
}

func (m *RipShowStatisticsBd) GetStandbyPacketsReceived() uint32 {
	if m != nil {
		return m.StandbyPacketsReceived
	}
	return 0
}

func (m *RipShowStatisticsBd) GetSentMessages() uint32 {
	if m != nil {
		return m.SentMessages
	}
	return 0
}

func (m *RipShowStatisticsBd) GetSentMessageFailures() uint32 {
	if m != nil {
		return m.SentMessageFailures
	}
	return 0
}

func (m *RipShowStatisticsBd) GetQueryResponses() uint32 {
	if m != nil {
		return m.QueryResponses
	}
	return 0
}

func (m *RipShowStatisticsBd) GetPeriodicUpdates() uint32 {
	if m != nil {
		return m.PeriodicUpdates
	}
	return 0
}

func (m *RipShowStatisticsBd) GetRouteCount() uint32 {
	if m != nil {
		return m.RouteCount
	}
	return 0
}

func (m *RipShowStatisticsBd) GetPathCount() uint32 {
	if m != nil {
		return m.PathCount
	}
	return 0
}

func (m *RipShowStatisticsBd) GetRouteMallocFailures() uint32 {
	if m != nil {
		return m.RouteMallocFailures
	}
	return 0
}

func (m *RipShowStatisticsBd) GetPathMallocFailures() uint32 {
	if m != nil {
		return m.PathMallocFailures
	}
	return 0
}

func (m *RipShowStatisticsBd) GetRibUpdates() uint32 {
	if m != nil {
		return m.RibUpdates
	}
	return 0
}

func init() {
	proto.RegisterType((*RipShowStatisticsBd_KEYS)(nil), "cisco_ios_xr_ip_rip_oper.rip.default_vrf.statistics.rip_show_statistics_bd_KEYS")
	proto.RegisterType((*RipShowStatisticsBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.default_vrf.statistics.rip_show_statistics_bd")
}

func init() { proto.RegisterFile("rip_show_statistics_bd.proto", fileDescriptor_66227cbf5e51e264) }

var fileDescriptor_66227cbf5e51e264 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x5d, 0x6b, 0x14, 0x31,
	0x14, 0x86, 0x11, 0x44, 0xf0, 0x68, 0x6d, 0x8d, 0xb6, 0x04, 0xb4, 0x28, 0xf5, 0x42, 0x8b, 0x30,
	0xc8, 0xae, 0x1f, 0xf5, 0xf3, 0x46, 0xf4, 0x46, 0x0a, 0x32, 0xe2, 0x85, 0x57, 0x87, 0x4c, 0x92,
	0xb5, 0xc1, 0xe9, 0x24, 0xe6, 0x64, 0x56, 0xf7, 0x27, 0xfa, 0xaf, 0x24, 0x67, 0x32, 0xb3, 0xcb,
	0xb2, 0xb7, 0xcf, 0xf3, 0xbe, 0x67, 0xf3, 0x2e, 0x03, 0xf7, 0xa3, 0x0b, 0x48, 0x17, 0xfe, 0x0f,
	0x52, 0x52, 0xc9, 0x51, 0x72, 0x9a, 0xb0, 0x31, 0x55, 0x88, 0x3e, 0x79, 0x31, 0xd7, 0x8e, 0xb4,
	0x47, 0xe7, 0x09, 0xff, 0x46, 0x74, 0x01, 0x73, 0xda, 0x07, 0x1b, 0xab, 0xe8, 0x42, 0x65, 0xec,
	0x42, 0xf5, 0x6d, 0xc2, 0x65, 0x5c, 0x54, 0xeb, 0xf6, 0xc9, 0x31, 0xdc, 0xdb, 0x7d, 0x14, 0xbf,
	0x7c, 0xfa, 0xf1, 0xed, 0xe4, 0xdf, 0x55, 0x38, 0xda, 0xed, 0xc5, 0x29, 0x1c, 0x44, 0xab, 0xad,
	0x5b, 0x5a, 0x83, 0x41, 0xe9, 0x5f, 0x36, 0x91, 0x9c, 0x3d, 0xbc, 0xf2, 0x64, 0xaf, 0xde, 0x1f,
	0xf9, 0xd7, 0x01, 0x8b, 0xa7, 0x70, 0xdb, 0x38, 0xd2, 0x2a, 0x9a, 0x8d, 0xec, 0x9c, 0xb3, 0x07,
	0x93, 0x18, 0xc3, 0xa7, 0xb0, 0x66, 0x18, 0x7d, 0x9f, 0x2c, 0xc9, 0xe7, 0xc3, 0xdd, 0x89, 0xd7,
	0x8c, 0xc5, 0x19, 0x48, 0x4a, 0xaa, 0x33, 0xcd, 0x6a, 0xbc, 0x8a, 0xe3, 0x4f, 0xcb, 0x17, 0x5c,
	0x39, 0x2a, 0xbe, 0x1c, 0xaf, 0x8b, 0x15, 0x8f, 0x60, 0x8f, 0x6c, 0x97, 0xf0, 0xd2, 0x12, 0xa9,
	0x9f, 0x96, 0xe4, 0x4b, 0x8e, 0xdf, 0xcc, 0xf0, 0xbc, 0x30, 0x31, 0x83, 0xc3, 0xcd, 0x10, 0x2e,
	0x94, 0x6b, 0xfb, 0x68, 0x49, 0xbe, 0xe2, 0xf0, 0x9d, 0x8d, 0xf0, 0xe7, 0xa2, 0xc4, 0x63, 0xd8,
	0xff, 0xdd, 0xdb, 0xb8, 0xc2, 0x68, 0x29, 0xf8, 0x8e, 0x2c, 0xc9, 0x33, 0x4e, 0xdf, 0x62, 0x5c,
	0x8f, 0x34, 0xcf, 0x0c, 0x36, 0x3a, 0x6f, 0x9c, 0xc6, 0x3e, 0x18, 0x95, 0x67, 0xbe, 0x1e, 0x66,
	0x8e, 0xfc, 0xfb, 0x80, 0xc5, 0x03, 0xb8, 0xc1, 0xff, 0x03, 0x6a, 0xdf, 0x77, 0x49, 0xbe, 0xe1,
	0x14, 0x30, 0xfa, 0x98, 0x89, 0x38, 0x06, 0x08, 0x2a, 0x5d, 0x14, 0xff, 0x96, 0xfd, 0xf5, 0x4c,
	0x06, 0x3d, 0x83, 0xc3, 0xa1, 0x7f, 0xa9, 0xda, 0xd6, 0xeb, 0xf5, 0x8e, 0x77, 0xc3, 0x0e, 0x96,
	0xe7, 0xec, 0xa6, 0x1d, 0xcf, 0xe0, 0x2e, 0x9f, 0xdc, 0xae, 0xbc, 0xe7, 0x8a, 0xc8, 0x6e, 0xab,
	0x91, 0x5f, 0xe9, 0x9a, 0x69, 0xcb, 0x87, 0xf2, 0x4a, 0xd7, 0x94, 0x19, 0xcd, 0x35, 0xfe, 0x4c,
	0xe7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x37, 0x3d, 0x87, 0xc6, 0x02, 0x00, 0x00,
}