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
// source: iedge_stats_coord_auth_data.proto

package cisco_ios_xr_iedge4710_oper_subscriber_manager_nodes_node_statistics_aaa_aggregate_authorization

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

type IedgeStatsCoordAuthData_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgeStatsCoordAuthData_KEYS) Reset()         { *m = IedgeStatsCoordAuthData_KEYS{} }
func (m *IedgeStatsCoordAuthData_KEYS) String() string { return proto.CompactTextString(m) }
func (*IedgeStatsCoordAuthData_KEYS) ProtoMessage()    {}
func (*IedgeStatsCoordAuthData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2510214c953205, []int{0}
}

func (m *IedgeStatsCoordAuthData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeStatsCoordAuthData_KEYS.Unmarshal(m, b)
}
func (m *IedgeStatsCoordAuthData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeStatsCoordAuthData_KEYS.Marshal(b, m, deterministic)
}
func (m *IedgeStatsCoordAuthData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeStatsCoordAuthData_KEYS.Merge(m, src)
}
func (m *IedgeStatsCoordAuthData_KEYS) XXX_Size() int {
	return xxx_messageInfo_IedgeStatsCoordAuthData_KEYS.Size(m)
}
func (m *IedgeStatsCoordAuthData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeStatsCoordAuthData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeStatsCoordAuthData_KEYS proto.InternalMessageInfo

func (m *IedgeStatsCoordAuthData_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type IedgeStatsCoordAuthData struct {
	SentRequests         uint64   `protobuf:"varint,50,opt,name=sent_requests,json=sentRequests,proto3" json:"sent_requests,omitempty"`
	AcceptedRequests     uint64   `protobuf:"varint,51,opt,name=accepted_requests,json=acceptedRequests,proto3" json:"accepted_requests,omitempty"`
	SuccessfulRequests   uint64   `protobuf:"varint,52,opt,name=successful_requests,json=successfulRequests,proto3" json:"successful_requests,omitempty"`
	RejectedRequests     uint64   `protobuf:"varint,53,opt,name=rejected_requests,json=rejectedRequests,proto3" json:"rejected_requests,omitempty"`
	UnreachableRequests  uint64   `protobuf:"varint,54,opt,name=unreachable_requests,json=unreachableRequests,proto3" json:"unreachable_requests,omitempty"`
	ErroredRequests      uint64   `protobuf:"varint,55,opt,name=errored_requests,json=erroredRequests,proto3" json:"errored_requests,omitempty"`
	TerminatedRequests   uint64   `protobuf:"varint,56,opt,name=terminated_requests,json=terminatedRequests,proto3" json:"terminated_requests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IedgeStatsCoordAuthData) Reset()         { *m = IedgeStatsCoordAuthData{} }
func (m *IedgeStatsCoordAuthData) String() string { return proto.CompactTextString(m) }
func (*IedgeStatsCoordAuthData) ProtoMessage()    {}
func (*IedgeStatsCoordAuthData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc2510214c953205, []int{1}
}

func (m *IedgeStatsCoordAuthData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IedgeStatsCoordAuthData.Unmarshal(m, b)
}
func (m *IedgeStatsCoordAuthData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IedgeStatsCoordAuthData.Marshal(b, m, deterministic)
}
func (m *IedgeStatsCoordAuthData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IedgeStatsCoordAuthData.Merge(m, src)
}
func (m *IedgeStatsCoordAuthData) XXX_Size() int {
	return xxx_messageInfo_IedgeStatsCoordAuthData.Size(m)
}
func (m *IedgeStatsCoordAuthData) XXX_DiscardUnknown() {
	xxx_messageInfo_IedgeStatsCoordAuthData.DiscardUnknown(m)
}

var xxx_messageInfo_IedgeStatsCoordAuthData proto.InternalMessageInfo

func (m *IedgeStatsCoordAuthData) GetSentRequests() uint64 {
	if m != nil {
		return m.SentRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetAcceptedRequests() uint64 {
	if m != nil {
		return m.AcceptedRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetSuccessfulRequests() uint64 {
	if m != nil {
		return m.SuccessfulRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetRejectedRequests() uint64 {
	if m != nil {
		return m.RejectedRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetUnreachableRequests() uint64 {
	if m != nil {
		return m.UnreachableRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetErroredRequests() uint64 {
	if m != nil {
		return m.ErroredRequests
	}
	return 0
}

func (m *IedgeStatsCoordAuthData) GetTerminatedRequests() uint64 {
	if m != nil {
		return m.TerminatedRequests
	}
	return 0
}

func init() {
	proto.RegisterType((*IedgeStatsCoordAuthData_KEYS)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.manager.nodes.node.statistics.aaa.aggregate_authorization.iedge_stats_coord_auth_data_KEYS")
	proto.RegisterType((*IedgeStatsCoordAuthData)(nil), "cisco_ios_xr_iedge4710_oper.subscriber.manager.nodes.node.statistics.aaa.aggregate_authorization.iedge_stats_coord_auth_data")
}

func init() { proto.RegisterFile("iedge_stats_coord_auth_data.proto", fileDescriptor_fc2510214c953205) }

var fileDescriptor_fc2510214c953205 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x31, 0x4f, 0x32, 0x41,
	0x10, 0x86, 0xc3, 0x97, 0x2f, 0x46, 0x36, 0x1a, 0xf1, 0xb0, 0x20, 0xa1, 0x41, 0x6c, 0x30, 0x26,
	0xab, 0x08, 0x8a, 0x9d, 0x95, 0x95, 0x89, 0x05, 0x56, 0x56, 0xeb, 0xdc, 0xde, 0x78, 0xac, 0xe1,
	0x76, 0x71, 0x66, 0x37, 0x31, 0xfe, 0x56, 0x7f, 0x8c, 0xb9, 0x05, 0xee, 0x8e, 0x86, 0xe6, 0x8a,
	0xf7, 0x7d, 0xe6, 0xb9, 0xb7, 0x58, 0x71, 0x6e, 0x30, 0xcb, 0x51, 0xb1, 0x07, 0xcf, 0x4a, 0x3b,
	0x47, 0x99, 0x82, 0xe0, 0x17, 0x2a, 0x03, 0x0f, 0x72, 0x45, 0xce, 0xbb, 0xe4, 0x5d, 0x1b, 0xd6,
	0x4e, 0x19, 0xc7, 0xea, 0x9b, 0x54, 0xe4, 0xa7, 0xb3, 0xf1, 0x8d, 0x72, 0x2b, 0x24, 0xc9, 0x21,
	0x65, 0x4d, 0x26, 0x45, 0x92, 0x05, 0x58, 0xc8, 0x91, 0xa4, 0x75, 0x19, 0x72, 0xfc, 0xca, 0x52,
	0x6b, 0xd8, 0x1b, 0xcd, 0x12, 0x00, 0x24, 0xe4, 0x39, 0x61, 0x0e, 0x1e, 0xe3, 0x3f, 0x1c, 0x99,
	0x1f, 0xf0, 0xc6, 0xd9, 0xe1, 0xa3, 0x18, 0xec, 0x99, 0xa1, 0x9e, 0x9f, 0xde, 0x5e, 0x93, 0xbe,
	0x68, 0x97, 0x4a, 0x65, 0xa1, 0xc0, 0x5e, 0x6b, 0xd0, 0x1a, 0xb5, 0xe7, 0x87, 0x65, 0xf0, 0x02,
	0x05, 0x0e, 0x7f, 0xff, 0x89, 0xfe, 0x1e, 0x43, 0x72, 0x21, 0x8e, 0x19, 0xad, 0x57, 0x84, 0x5f,
	0x01, 0xd9, 0x73, 0xef, 0x76, 0xd0, 0x1a, 0xfd, 0x9f, 0x1f, 0x95, 0xe1, 0x7c, 0x93, 0x25, 0x57,
	0xe2, 0x14, 0xb4, 0xc6, 0x95, 0xc7, 0xac, 0x06, 0x27, 0x11, 0xec, 0x6c, 0x8b, 0x0a, 0xbe, 0x16,
	0x5d, 0x0e, 0x5a, 0x23, 0xf3, 0x47, 0x58, 0xd6, 0xf8, 0x34, 0xe2, 0x49, 0x5d, 0x35, 0xed, 0x84,
	0x9f, 0xa8, 0x77, 0xec, 0x77, 0x6b, 0xfb, 0xb6, 0xa8, 0xe0, 0xb1, 0x38, 0x0b, 0x96, 0x10, 0xf4,
	0x02, 0xd2, 0x25, 0xd6, 0xfc, 0x7d, 0xe4, 0xbb, 0x8d, 0xae, 0x3a, 0xb9, 0x14, 0x1d, 0x24, 0x72,
	0xd4, 0xd4, 0xcf, 0x22, 0x7e, 0xb2, 0xc9, 0x9b, 0xdb, 0x3d, 0x52, 0x61, 0x2c, 0xec, 0x8c, 0x79,
	0x58, 0x6f, 0xaf, 0xab, 0xed, 0x41, 0x7a, 0x10, 0x1f, 0xc2, 0xe4, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0x6d, 0x5c, 0x41, 0x2d, 0x02, 0x00, 0x00,
}
