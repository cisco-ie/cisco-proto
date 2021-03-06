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
// source: elod_summary_info.proto

package cisco_ios_xr_ethernet_link_oam_oper_ether_link_oam_nodes_node_summary

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

type ElodSummaryInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElodSummaryInfo_KEYS) Reset()         { *m = ElodSummaryInfo_KEYS{} }
func (m *ElodSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ElodSummaryInfo_KEYS) ProtoMessage()    {}
func (*ElodSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_22bfa49103093fdc, []int{0}
}

func (m *ElodSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElodSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *ElodSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElodSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ElodSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElodSummaryInfo_KEYS.Merge(m, src)
}
func (m *ElodSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ElodSummaryInfo_KEYS.Size(m)
}
func (m *ElodSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ElodSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ElodSummaryInfo_KEYS proto.InternalMessageInfo

func (m *ElodSummaryInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type ElodSummaryInfo struct {
	Interfaces           uint32   `protobuf:"varint,50,opt,name=interfaces,proto3" json:"interfaces,omitempty"`
	PortDown             uint32   `protobuf:"varint,51,opt,name=port_down,json=portDown,proto3" json:"port_down,omitempty"`
	PassiveWait          uint32   `protobuf:"varint,52,opt,name=passive_wait,json=passiveWait,proto3" json:"passive_wait,omitempty"`
	ActiveSend           uint32   `protobuf:"varint,53,opt,name=active_send,json=activeSend,proto3" json:"active_send,omitempty"`
	Evaluating           uint32   `protobuf:"varint,54,opt,name=evaluating,proto3" json:"evaluating,omitempty"`
	LocalAccept          uint32   `protobuf:"varint,55,opt,name=local_accept,json=localAccept,proto3" json:"local_accept,omitempty"`
	LocalReject          uint32   `protobuf:"varint,56,opt,name=local_reject,json=localReject,proto3" json:"local_reject,omitempty"`
	RemoteReject         uint32   `protobuf:"varint,57,opt,name=remote_reject,json=remoteReject,proto3" json:"remote_reject,omitempty"`
	Operational          uint32   `protobuf:"varint,58,opt,name=operational,proto3" json:"operational,omitempty"`
	LoopbackMode         uint32   `protobuf:"varint,59,opt,name=loopback_mode,json=loopbackMode,proto3" json:"loopback_mode,omitempty"`
	MiswiredConnections  uint32   `protobuf:"varint,60,opt,name=miswired_connections,json=miswiredConnections,proto3" json:"miswired_connections,omitempty"`
	Events               uint64   `protobuf:"varint,61,opt,name=events,proto3" json:"events,omitempty"`
	LocalEvents          uint64   `protobuf:"varint,62,opt,name=local_events,json=localEvents,proto3" json:"local_events,omitempty"`
	LocalSymbolPeriod    uint64   `protobuf:"varint,63,opt,name=local_symbol_period,json=localSymbolPeriod,proto3" json:"local_symbol_period,omitempty"`
	LocalFrame           uint64   `protobuf:"varint,64,opt,name=local_frame,json=localFrame,proto3" json:"local_frame,omitempty"`
	LocalFramePeriod     uint64   `protobuf:"varint,65,opt,name=local_frame_period,json=localFramePeriod,proto3" json:"local_frame_period,omitempty"`
	LocalFrameSeconds    uint64   `protobuf:"varint,66,opt,name=local_frame_seconds,json=localFrameSeconds,proto3" json:"local_frame_seconds,omitempty"`
	RemoteEvents         uint64   `protobuf:"varint,67,opt,name=remote_events,json=remoteEvents,proto3" json:"remote_events,omitempty"`
	RemoteSymbolPeriod   uint64   `protobuf:"varint,68,opt,name=remote_symbol_period,json=remoteSymbolPeriod,proto3" json:"remote_symbol_period,omitempty"`
	RemoteFrame          uint64   `protobuf:"varint,69,opt,name=remote_frame,json=remoteFrame,proto3" json:"remote_frame,omitempty"`
	RemoteFramePeriod    uint64   `protobuf:"varint,70,opt,name=remote_frame_period,json=remoteFramePeriod,proto3" json:"remote_frame_period,omitempty"`
	RemoteFrameSeconds   uint64   `protobuf:"varint,71,opt,name=remote_frame_seconds,json=remoteFrameSeconds,proto3" json:"remote_frame_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElodSummaryInfo) Reset()         { *m = ElodSummaryInfo{} }
func (m *ElodSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*ElodSummaryInfo) ProtoMessage()    {}
func (*ElodSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_22bfa49103093fdc, []int{1}
}

func (m *ElodSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElodSummaryInfo.Unmarshal(m, b)
}
func (m *ElodSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElodSummaryInfo.Marshal(b, m, deterministic)
}
func (m *ElodSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElodSummaryInfo.Merge(m, src)
}
func (m *ElodSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_ElodSummaryInfo.Size(m)
}
func (m *ElodSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ElodSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ElodSummaryInfo proto.InternalMessageInfo

func (m *ElodSummaryInfo) GetInterfaces() uint32 {
	if m != nil {
		return m.Interfaces
	}
	return 0
}

func (m *ElodSummaryInfo) GetPortDown() uint32 {
	if m != nil {
		return m.PortDown
	}
	return 0
}

func (m *ElodSummaryInfo) GetPassiveWait() uint32 {
	if m != nil {
		return m.PassiveWait
	}
	return 0
}

func (m *ElodSummaryInfo) GetActiveSend() uint32 {
	if m != nil {
		return m.ActiveSend
	}
	return 0
}

func (m *ElodSummaryInfo) GetEvaluating() uint32 {
	if m != nil {
		return m.Evaluating
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalAccept() uint32 {
	if m != nil {
		return m.LocalAccept
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalReject() uint32 {
	if m != nil {
		return m.LocalReject
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteReject() uint32 {
	if m != nil {
		return m.RemoteReject
	}
	return 0
}

func (m *ElodSummaryInfo) GetOperational() uint32 {
	if m != nil {
		return m.Operational
	}
	return 0
}

func (m *ElodSummaryInfo) GetLoopbackMode() uint32 {
	if m != nil {
		return m.LoopbackMode
	}
	return 0
}

func (m *ElodSummaryInfo) GetMiswiredConnections() uint32 {
	if m != nil {
		return m.MiswiredConnections
	}
	return 0
}

func (m *ElodSummaryInfo) GetEvents() uint64 {
	if m != nil {
		return m.Events
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalEvents() uint64 {
	if m != nil {
		return m.LocalEvents
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalSymbolPeriod() uint64 {
	if m != nil {
		return m.LocalSymbolPeriod
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalFrame() uint64 {
	if m != nil {
		return m.LocalFrame
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalFramePeriod() uint64 {
	if m != nil {
		return m.LocalFramePeriod
	}
	return 0
}

func (m *ElodSummaryInfo) GetLocalFrameSeconds() uint64 {
	if m != nil {
		return m.LocalFrameSeconds
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteEvents() uint64 {
	if m != nil {
		return m.RemoteEvents
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteSymbolPeriod() uint64 {
	if m != nil {
		return m.RemoteSymbolPeriod
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteFrame() uint64 {
	if m != nil {
		return m.RemoteFrame
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteFramePeriod() uint64 {
	if m != nil {
		return m.RemoteFramePeriod
	}
	return 0
}

func (m *ElodSummaryInfo) GetRemoteFrameSeconds() uint64 {
	if m != nil {
		return m.RemoteFrameSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*ElodSummaryInfo_KEYS)(nil), "cisco_ios_xr_ethernet_link_oam_oper.ether_link_oam.nodes.node.summary.elod_summary_info_KEYS")
	proto.RegisterType((*ElodSummaryInfo)(nil), "cisco_ios_xr_ethernet_link_oam_oper.ether_link_oam.nodes.node.summary.elod_summary_info")
}

func init() { proto.RegisterFile("elod_summary_info.proto", fileDescriptor_22bfa49103093fdc) }

var fileDescriptor_22bfa49103093fdc = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0x4d, 0x6f, 0x13, 0x3d,
	0x14, 0x85, 0x15, 0xe9, 0x55, 0xd4, 0x3a, 0xad, 0xf4, 0xd6, 0xa9, 0x8a, 0x25, 0x24, 0x48, 0xcb,
	0x26, 0x0b, 0x34, 0x02, 0x4a, 0xf9, 0xfe, 0x2a, 0x6d, 0xca, 0x02, 0x81, 0x50, 0xb2, 0x40, 0xac,
	0x2c, 0xc7, 0x73, 0x03, 0xa6, 0x33, 0xbe, 0x23, 0xdb, 0x4d, 0xe8, 0xaf, 0xe5, 0xaf, 0x20, 0x5f,
	0x3b, 0x9d, 0x89, 0xba, 0x89, 0x94, 0x73, 0x9e, 0x39, 0xf7, 0x9e, 0xdc, 0x0c, 0xbb, 0x03, 0x15,
	0x96, 0xd2, 0x5f, 0xd5, 0xb5, 0x72, 0xd7, 0xd2, 0xd8, 0x05, 0x16, 0x8d, 0xc3, 0x80, 0x7c, 0xa2,
	0x8d, 0xd7, 0x28, 0x0d, 0x7a, 0xf9, 0xc7, 0x49, 0x08, 0xbf, 0xc0, 0x59, 0x08, 0xb2, 0x32, 0xf6,
	0x52, 0xa2, 0xaa, 0x25, 0x36, 0xe0, 0x0a, 0x92, 0x6f, 0xb4, 0xc2, 0x62, 0x09, 0x9e, 0x3e, 0x8b,
	0x9c, 0x78, 0x74, 0xc2, 0x0e, 0x6e, 0x4d, 0x90, 0x9f, 0x27, 0x3f, 0x66, 0xfc, 0x2e, 0xdb, 0x8e,
	0xa4, 0xb4, 0xaa, 0x06, 0xd1, 0x1b, 0xf5, 0xc6, 0xdb, 0xd3, 0xad, 0x28, 0x7c, 0x55, 0x35, 0x1c,
	0xfd, 0xed, 0xb3, 0xbd, 0x5b, 0xcf, 0xf1, 0x7b, 0x8c, 0x19, 0x1b, 0xc0, 0x2d, 0x94, 0x06, 0x2f,
	0x9e, 0x8c, 0x7a, 0xe3, 0xdd, 0x69, 0x47, 0x89, 0x91, 0x0d, 0xba, 0x20, 0x4b, 0x5c, 0x59, 0x71,
	0x4c, 0xf6, 0x56, 0x14, 0xce, 0x71, 0x65, 0xf9, 0x21, 0xdb, 0x69, 0x94, 0xf7, 0x66, 0x09, 0x72,
	0xa5, 0x4c, 0x10, 0x4f, 0xc9, 0x1f, 0x64, 0xed, 0xbb, 0x32, 0x81, 0xdf, 0x67, 0x03, 0xa5, 0x43,
	0x24, 0x3c, 0xd8, 0x52, 0x9c, 0xa4, 0x01, 0x49, 0x9a, 0x81, 0x2d, 0xe3, 0x02, 0xb0, 0x54, 0xd5,
	0x95, 0x0a, 0xc6, 0xfe, 0x14, 0xcf, 0x92, 0xdf, 0x2a, 0x71, 0x46, 0x85, 0x5a, 0x55, 0x52, 0x69,
	0x0d, 0x4d, 0x10, 0xcf, 0xd3, 0x0c, 0xd2, 0x4e, 0x49, 0x6a, 0x11, 0x07, 0xbf, 0x41, 0x07, 0xf1,
	0xa2, 0x83, 0x4c, 0x49, 0xe2, 0x0f, 0xd8, 0xae, 0x83, 0x1a, 0x03, 0xac, 0x99, 0x97, 0xc4, 0xec,
	0x24, 0x31, 0x43, 0x23, 0x36, 0x88, 0x27, 0x50, 0xc1, 0xa0, 0x55, 0x95, 0x78, 0x95, 0x62, 0x3a,
	0x52, 0x8c, 0xa9, 0x10, 0x9b, 0xb9, 0xd2, 0x97, 0xb2, 0xc6, 0x12, 0xc4, 0xeb, 0x14, 0xb3, 0x16,
	0xbf, 0x60, 0x09, 0xfc, 0x31, 0xdb, 0xaf, 0x8d, 0x5f, 0x19, 0x07, 0xa5, 0xd4, 0x68, 0x2d, 0xe8,
	0xf8, 0xb4, 0x17, 0x6f, 0x88, 0x1d, 0xae, 0xbd, 0xb3, 0xd6, 0xe2, 0x07, 0xac, 0x0f, 0x4b, 0xb0,
	0xc1, 0x8b, 0xb7, 0xa3, 0xde, 0xf8, 0xbf, 0x69, 0xfe, 0xd6, 0x36, 0xcb, 0xee, 0x3b, 0x72, 0x53,
	0xb3, 0x49, 0x42, 0x0a, 0x36, 0x4c, 0x88, 0xbf, 0xae, 0xe7, 0x58, 0xc9, 0x06, 0x9c, 0xc1, 0x52,
	0xbc, 0x27, 0x72, 0x8f, 0xac, 0x19, 0x39, 0xdf, 0xc8, 0x88, 0x07, 0x49, 0xfc, 0xc2, 0xc5, 0x7f,
	0xc9, 0x07, 0xe2, 0x18, 0x49, 0x17, 0x51, 0xe1, 0x0f, 0x19, 0xef, 0x00, 0xeb, 0xbc, 0x53, 0xe2,
	0xfe, 0x6f, 0xb9, 0x1c, 0x77, 0x33, 0x3e, 0xd1, 0x1e, 0x34, 0xda, 0xd2, 0x8b, 0x8f, 0x9d, 0xf1,
	0x84, 0xcf, 0x92, 0xd1, 0x39, 0x44, 0xae, 0x74, 0x46, 0x64, 0x3e, 0x44, 0xee, 0xf4, 0x88, 0xed,
	0x67, 0x68, 0xb3, 0xd4, 0x39, 0xb1, 0x3c, 0x79, 0x1b, 0xad, 0x0e, 0x59, 0x4e, 0xc8, 0xb5, 0x26,
	0xe9, 0x87, 0x4a, 0x5a, 0xea, 0x55, 0xb0, 0x61, 0x17, 0x59, 0x67, 0x5e, 0xa4, 0x4d, 0x3b, 0x64,
	0x8e, 0x6c, 0x97, 0xd8, 0xac, 0xf6, 0xa9, 0xbb, 0x44, 0xb7, 0xdb, 0xbc, 0x4f, 0xaf, 0xf9, 0xf1,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x24, 0x36, 0x36, 0x01, 0x04, 0x00, 0x00,
}
