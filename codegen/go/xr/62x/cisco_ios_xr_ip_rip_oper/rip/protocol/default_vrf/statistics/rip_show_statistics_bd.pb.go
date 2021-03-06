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

package cisco_ios_xr_ip_rip_oper_rip_protocol_default_vrf_statistics

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
	proto.RegisterType((*RipShowStatisticsBd_KEYS)(nil), "cisco_ios_xr_ip_rip_oper.rip.protocol.default_vrf.statistics.rip_show_statistics_bd_KEYS")
	proto.RegisterType((*RipShowStatisticsBd)(nil), "cisco_ios_xr_ip_rip_oper.rip.protocol.default_vrf.statistics.rip_show_statistics_bd")
}

func init() { proto.RegisterFile("rip_show_statistics_bd.proto", fileDescriptor_66227cbf5e51e264) }

var fileDescriptor_66227cbf5e51e264 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x6b, 0x15, 0x31,
	0x14, 0xc5, 0x11, 0x44, 0xf0, 0x6a, 0x6d, 0x8d, 0xb6, 0x04, 0xb4, 0x28, 0x75, 0xa1, 0x45, 0x18,
	0xe4, 0xd5, 0x3f, 0x55, 0xab, 0x1b, 0xd1, 0x8d, 0x14, 0x64, 0xc4, 0x85, 0xab, 0x4b, 0x26, 0xc9,
	0xb3, 0xc1, 0xe9, 0x24, 0xe6, 0x66, 0xaa, 0xfd, 0x88, 0x7e, 0xab, 0x92, 0x9b, 0xcc, 0xbc, 0x47,
	0xe9, 0xf6, 0x77, 0x7e, 0xf7, 0xbc, 0x9c, 0xc7, 0xc0, 0xc3, 0xe8, 0x02, 0xd2, 0x89, 0xff, 0x8b,
	0x94, 0x54, 0x72, 0x94, 0x9c, 0x26, 0xec, 0x4c, 0x13, 0xa2, 0x4f, 0x5e, 0x1c, 0x69, 0x47, 0xda,
	0xa3, 0xf3, 0x84, 0xff, 0x22, 0xba, 0x80, 0xd9, 0xf6, 0xc1, 0xc6, 0x26, 0xba, 0x50, 0x1c, 0xed,
	0xfb, 0xc6, 0xd8, 0xa5, 0x1a, 0xfb, 0x84, 0x67, 0x71, 0xd9, 0xac, 0x6a, 0xf6, 0x76, 0xe1, 0xc1,
	0xd5, 0xed, 0xf8, 0xf5, 0xf3, 0xcf, 0xef, 0x7b, 0xff, 0xaf, 0xc3, 0xce, 0xd5, 0xb9, 0xd8, 0x87,
	0xad, 0x68, 0xb5, 0x75, 0x67, 0xd6, 0x60, 0x50, 0xfa, 0xb7, 0x4d, 0x24, 0x17, 0x8f, 0xaf, 0x3d,
	0xdb, 0x68, 0x37, 0x27, 0xfe, 0xad, 0x60, 0xf1, 0x1c, 0xee, 0x1a, 0x47, 0x5a, 0x45, 0xb3, 0xe6,
	0x1e, 0xb0, 0xbb, 0x35, 0x07, 0x93, 0xbc, 0x0f, 0x2b, 0x86, 0xd1, 0x8f, 0xc9, 0x92, 0x7c, 0x59,
	0x7a, 0x67, 0xde, 0x32, 0x16, 0x87, 0x20, 0x29, 0xa9, 0xc1, 0x74, 0xe7, 0x53, 0x2b, 0x4e, 0x3f,
	0x2d, 0x5f, 0xf1, 0xc9, 0x4e, 0xcd, 0x6b, 0x79, 0x5b, 0x53, 0xf1, 0x04, 0x36, 0xc8, 0x0e, 0x09,
	0x4f, 0x2d, 0x91, 0xfa, 0x65, 0x49, 0xbe, 0x66, 0xfd, 0x76, 0x86, 0xc7, 0x95, 0x89, 0x05, 0x6c,
	0xaf, 0x4b, 0xb8, 0x54, 0xae, 0x1f, 0xa3, 0x25, 0xf9, 0x86, 0xe5, 0x7b, 0x6b, 0xf2, 0x97, 0x1a,
	0x89, 0xa7, 0xb0, 0xf9, 0x67, 0xb4, 0xf1, 0x1c, 0xa3, 0xa5, 0xe0, 0x07, 0xb2, 0x24, 0x0f, 0xd9,
	0xbe, 0xc3, 0xb8, 0x9d, 0x68, 0x9e, 0x19, 0x6c, 0x74, 0xde, 0x38, 0x8d, 0x63, 0x30, 0x2a, 0xcf,
	0x7c, 0x5b, 0x66, 0x4e, 0xfc, 0x47, 0xc1, 0xe2, 0x11, 0xdc, 0xe2, 0xff, 0x01, 0xb5, 0x1f, 0x87,
	0x24, 0xdf, 0xb1, 0x05, 0x8c, 0x3e, 0x65, 0x22, 0x76, 0x01, 0x82, 0x4a, 0x27, 0x35, 0x7f, 0xcf,
	0xf9, 0xcd, 0x4c, 0x4a, 0xbc, 0x80, 0xed, 0x72, 0x7f, 0xaa, 0xfa, 0xde, 0xeb, 0xd5, 0x8e, 0xa3,
	0xb2, 0x83, 0xc3, 0x63, 0xce, 0xe6, 0x1d, 0x2f, 0xe0, 0x3e, 0x57, 0x5e, 0x3e, 0xf9, 0xc0, 0x27,
	0x22, 0x67, 0x97, 0x2e, 0xf2, 0x2b, 0x5d, 0x37, 0x6f, 0xf9, 0x58, 0x5f, 0xe9, 0xba, 0x3a, 0xa3,
	0xbb, 0xc1, 0xdf, 0xe2, 0xc1, 0x45, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xee, 0x86, 0x9b, 0xcf,
	0x02, 0x00, 0x00,
}
