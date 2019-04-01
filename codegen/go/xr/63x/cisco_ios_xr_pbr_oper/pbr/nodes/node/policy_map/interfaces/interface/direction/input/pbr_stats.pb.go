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
// source: pbr_stats.proto

package cisco_ios_xr_pbr_oper_pbr_nodes_node_policy_map_interfaces_interface_direction_input

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

type PbrStats_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PbrStats_KEYS) Reset()         { *m = PbrStats_KEYS{} }
func (m *PbrStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*PbrStats_KEYS) ProtoMessage()    {}
func (*PbrStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{0}
}

func (m *PbrStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbrStats_KEYS.Unmarshal(m, b)
}
func (m *PbrStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbrStats_KEYS.Marshal(b, m, deterministic)
}
func (m *PbrStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbrStats_KEYS.Merge(m, src)
}
func (m *PbrStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_PbrStats_KEYS.Size(m)
}
func (m *PbrStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PbrStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PbrStats_KEYS proto.InternalMessageInfo

func (m *PbrStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PbrStats_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type GenStatsSt struct {
	TransmitPackets         uint64   `protobuf:"varint,1,opt,name=transmit_packets,json=transmitPackets,proto3" json:"transmit_packets,omitempty"`
	TransmitBytes           uint64   `protobuf:"varint,2,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	TotalDropPackets        uint64   `protobuf:"varint,3,opt,name=total_drop_packets,json=totalDropPackets,proto3" json:"total_drop_packets,omitempty"`
	TotalDropBytes          uint64   `protobuf:"varint,4,opt,name=total_drop_bytes,json=totalDropBytes,proto3" json:"total_drop_bytes,omitempty"`
	TotalDropRate           uint32   `protobuf:"varint,5,opt,name=total_drop_rate,json=totalDropRate,proto3" json:"total_drop_rate,omitempty"`
	MatchDataRate           uint32   `protobuf:"varint,6,opt,name=match_data_rate,json=matchDataRate,proto3" json:"match_data_rate,omitempty"`
	TotalTransmitRate       uint32   `protobuf:"varint,7,opt,name=total_transmit_rate,json=totalTransmitRate,proto3" json:"total_transmit_rate,omitempty"`
	PrePolicyMatchedPackets uint64   `protobuf:"varint,8,opt,name=pre_policy_matched_packets,json=prePolicyMatchedPackets,proto3" json:"pre_policy_matched_packets,omitempty"`
	PrePolicyMatchedBytes   uint64   `protobuf:"varint,9,opt,name=pre_policy_matched_bytes,json=prePolicyMatchedBytes,proto3" json:"pre_policy_matched_bytes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GenStatsSt) Reset()         { *m = GenStatsSt{} }
func (m *GenStatsSt) String() string { return proto.CompactTextString(m) }
func (*GenStatsSt) ProtoMessage()    {}
func (*GenStatsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{1}
}

func (m *GenStatsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenStatsSt.Unmarshal(m, b)
}
func (m *GenStatsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenStatsSt.Marshal(b, m, deterministic)
}
func (m *GenStatsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenStatsSt.Merge(m, src)
}
func (m *GenStatsSt) XXX_Size() int {
	return xxx_messageInfo_GenStatsSt.Size(m)
}
func (m *GenStatsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_GenStatsSt.DiscardUnknown(m)
}

var xxx_messageInfo_GenStatsSt proto.InternalMessageInfo

func (m *GenStatsSt) GetTransmitPackets() uint64 {
	if m != nil {
		return m.TransmitPackets
	}
	return 0
}

func (m *GenStatsSt) GetTransmitBytes() uint64 {
	if m != nil {
		return m.TransmitBytes
	}
	return 0
}

func (m *GenStatsSt) GetTotalDropPackets() uint64 {
	if m != nil {
		return m.TotalDropPackets
	}
	return 0
}

func (m *GenStatsSt) GetTotalDropBytes() uint64 {
	if m != nil {
		return m.TotalDropBytes
	}
	return 0
}

func (m *GenStatsSt) GetTotalDropRate() uint32 {
	if m != nil {
		return m.TotalDropRate
	}
	return 0
}

func (m *GenStatsSt) GetMatchDataRate() uint32 {
	if m != nil {
		return m.MatchDataRate
	}
	return 0
}

func (m *GenStatsSt) GetTotalTransmitRate() uint32 {
	if m != nil {
		return m.TotalTransmitRate
	}
	return 0
}

func (m *GenStatsSt) GetPrePolicyMatchedPackets() uint64 {
	if m != nil {
		return m.PrePolicyMatchedPackets
	}
	return 0
}

func (m *GenStatsSt) GetPrePolicyMatchedBytes() uint64 {
	if m != nil {
		return m.PrePolicyMatchedBytes
	}
	return 0
}

type HttprStatsSt struct {
	RqstRcvdPackets      uint64   `protobuf:"varint,1,opt,name=rqst_rcvd_packets,json=rqstRcvdPackets,proto3" json:"rqst_rcvd_packets,omitempty"`
	RqstRcvdBytes        uint64   `protobuf:"varint,2,opt,name=rqst_rcvd_bytes,json=rqstRcvdBytes,proto3" json:"rqst_rcvd_bytes,omitempty"`
	DropPackets          uint64   `protobuf:"varint,3,opt,name=drop_packets,json=dropPackets,proto3" json:"drop_packets,omitempty"`
	DropBytes            uint64   `protobuf:"varint,4,opt,name=drop_bytes,json=dropBytes,proto3" json:"drop_bytes,omitempty"`
	RespSentPackets      uint64   `protobuf:"varint,5,opt,name=resp_sent_packets,json=respSentPackets,proto3" json:"resp_sent_packets,omitempty"`
	RespSentBytes        uint64   `protobuf:"varint,6,opt,name=resp_sent_bytes,json=respSentBytes,proto3" json:"resp_sent_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttprStatsSt) Reset()         { *m = HttprStatsSt{} }
func (m *HttprStatsSt) String() string { return proto.CompactTextString(m) }
func (*HttprStatsSt) ProtoMessage()    {}
func (*HttprStatsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{2}
}

func (m *HttprStatsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttprStatsSt.Unmarshal(m, b)
}
func (m *HttprStatsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttprStatsSt.Marshal(b, m, deterministic)
}
func (m *HttprStatsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttprStatsSt.Merge(m, src)
}
func (m *HttprStatsSt) XXX_Size() int {
	return xxx_messageInfo_HttprStatsSt.Size(m)
}
func (m *HttprStatsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_HttprStatsSt.DiscardUnknown(m)
}

var xxx_messageInfo_HttprStatsSt proto.InternalMessageInfo

func (m *HttprStatsSt) GetRqstRcvdPackets() uint64 {
	if m != nil {
		return m.RqstRcvdPackets
	}
	return 0
}

func (m *HttprStatsSt) GetRqstRcvdBytes() uint64 {
	if m != nil {
		return m.RqstRcvdBytes
	}
	return 0
}

func (m *HttprStatsSt) GetDropPackets() uint64 {
	if m != nil {
		return m.DropPackets
	}
	return 0
}

func (m *HttprStatsSt) GetDropBytes() uint64 {
	if m != nil {
		return m.DropBytes
	}
	return 0
}

func (m *HttprStatsSt) GetRespSentPackets() uint64 {
	if m != nil {
		return m.RespSentPackets
	}
	return 0
}

func (m *HttprStatsSt) GetRespSentBytes() uint64 {
	if m != nil {
		return m.RespSentBytes
	}
	return 0
}

type ClassStats struct {
	CounterValidityBitmask uint64        `protobuf:"varint,1,opt,name=counter_validity_bitmask,json=counterValidityBitmask,proto3" json:"counter_validity_bitmask,omitempty"`
	ClassName              string        `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ClassId                uint32        `protobuf:"varint,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	GeneralStats           *GenStatsSt   `protobuf:"bytes,4,opt,name=general_stats,json=generalStats,proto3" json:"general_stats,omitempty"`
	HttprStats             *HttprStatsSt `protobuf:"bytes,5,opt,name=httpr_stats,json=httprStats,proto3" json:"httpr_stats,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}      `json:"-"`
	XXX_unrecognized       []byte        `json:"-"`
	XXX_sizecache          int32         `json:"-"`
}

func (m *ClassStats) Reset()         { *m = ClassStats{} }
func (m *ClassStats) String() string { return proto.CompactTextString(m) }
func (*ClassStats) ProtoMessage()    {}
func (*ClassStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{3}
}

func (m *ClassStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassStats.Unmarshal(m, b)
}
func (m *ClassStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassStats.Marshal(b, m, deterministic)
}
func (m *ClassStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassStats.Merge(m, src)
}
func (m *ClassStats) XXX_Size() int {
	return xxx_messageInfo_ClassStats.Size(m)
}
func (m *ClassStats) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassStats.DiscardUnknown(m)
}

var xxx_messageInfo_ClassStats proto.InternalMessageInfo

func (m *ClassStats) GetCounterValidityBitmask() uint64 {
	if m != nil {
		return m.CounterValidityBitmask
	}
	return 0
}

func (m *ClassStats) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *ClassStats) GetClassId() uint32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *ClassStats) GetGeneralStats() *GenStatsSt {
	if m != nil {
		return m.GeneralStats
	}
	return nil
}

func (m *ClassStats) GetHttprStats() *HttprStatsSt {
	if m != nil {
		return m.HttprStats
	}
	return nil
}

type PbrStats struct {
	NodeName             string        `protobuf:"bytes,50,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PolicyName           string        `protobuf:"bytes,51,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	State                string        `protobuf:"bytes,52,opt,name=state,proto3" json:"state,omitempty"`
	StateDescription     string        `protobuf:"bytes,53,opt,name=state_description,json=stateDescription,proto3" json:"state_description,omitempty"`
	ClassStat            []*ClassStats `protobuf:"bytes,54,rep,name=class_stat,json=classStat,proto3" json:"class_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PbrStats) Reset()         { *m = PbrStats{} }
func (m *PbrStats) String() string { return proto.CompactTextString(m) }
func (*PbrStats) ProtoMessage()    {}
func (*PbrStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{4}
}

func (m *PbrStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PbrStats.Unmarshal(m, b)
}
func (m *PbrStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PbrStats.Marshal(b, m, deterministic)
}
func (m *PbrStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PbrStats.Merge(m, src)
}
func (m *PbrStats) XXX_Size() int {
	return xxx_messageInfo_PbrStats.Size(m)
}
func (m *PbrStats) XXX_DiscardUnknown() {
	xxx_messageInfo_PbrStats.DiscardUnknown(m)
}

var xxx_messageInfo_PbrStats proto.InternalMessageInfo

func (m *PbrStats) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PbrStats) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *PbrStats) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *PbrStats) GetStateDescription() string {
	if m != nil {
		return m.StateDescription
	}
	return ""
}

func (m *PbrStats) GetClassStat() []*ClassStats {
	if m != nil {
		return m.ClassStat
	}
	return nil
}

func init() {
	proto.RegisterType((*PbrStats_KEYS)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.pbr_stats_KEYS")
	proto.RegisterType((*GenStatsSt)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.gen_stats_st")
	proto.RegisterType((*HttprStatsSt)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.httpr_stats_st")
	proto.RegisterType((*ClassStats)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.class_stats")
	proto.RegisterType((*PbrStats)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.pbr_stats")
}

func init() { proto.RegisterFile("pbr_stats.proto", fileDescriptor_6233b79aff5b417f) }

var fileDescriptor_6233b79aff5b417f = []byte{
	// 648 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcb, 0x6f, 0x13, 0x3f,
	0x10, 0xc7, 0x95, 0xa6, 0xaf, 0xcc, 0x36, 0x8f, 0xee, 0xef, 0x07, 0x2c, 0x20, 0x44, 0x89, 0x44,
	0x15, 0x1e, 0xda, 0x43, 0xca, 0x4b, 0xe2, 0x56, 0x95, 0x03, 0x42, 0xa0, 0x6a, 0x5b, 0x21, 0x71,
	0xb2, 0x1c, 0xdb, 0xb4, 0x56, 0x93, 0x5d, 0x63, 0x4f, 0x2b, 0x7a, 0xe3, 0x82, 0x38, 0x72, 0xe0,
	0xff, 0xe0, 0x0f, 0xe4, 0x84, 0x3c, 0xde, 0x38, 0xdb, 0xd2, 0x6b, 0x2f, 0xd1, 0xee, 0x7c, 0x3f,
	0xfb, 0xf5, 0x8c, 0x67, 0xec, 0x40, 0xdf, 0x4c, 0x2c, 0x73, 0xc8, 0xd1, 0xe5, 0xc6, 0x56, 0x58,
	0xa5, 0x87, 0x42, 0x3b, 0x51, 0x31, 0x5d, 0x39, 0xf6, 0xd5, 0x32, 0xaf, 0x56, 0x46, 0xd9, 0xdc,
	0x4c, 0x6c, 0x5e, 0x56, 0x52, 0x39, 0xfa, 0xcd, 0x4d, 0x35, 0xd5, 0xe2, 0x9c, 0xcd, 0xb8, 0xc9,
	0x75, 0x89, 0xca, 0x7e, 0xe6, 0x42, 0xb9, 0xc5, 0x63, 0x2e, 0xb5, 0x55, 0x02, 0x75, 0x55, 0xe6,
	0xba, 0x34, 0xa7, 0x38, 0x3c, 0x84, 0x5e, 0x5c, 0x88, 0xbd, 0x7b, 0xf3, 0xe9, 0x20, 0xbd, 0x0b,
	0x1d, 0xef, 0xc4, 0x4a, 0x3e, 0x53, 0x59, 0x6b, 0xab, 0x35, 0xea, 0x14, 0xeb, 0x3e, 0xf0, 0x81,
	0xcf, 0x54, 0xfa, 0x10, 0x7a, 0xd1, 0x2b, 0x10, 0x4b, 0x44, 0x74, 0x63, 0xd4, 0x63, 0xc3, 0xdf,
	0x6d, 0xd8, 0x38, 0x52, 0x65, 0x6d, 0xeb, 0x30, 0x7d, 0x04, 0x03, 0xb4, 0xbc, 0x74, 0x33, 0x8d,
	0xcc, 0x70, 0x71, 0xa2, 0xd0, 0x91, 0xf7, 0x72, 0xd1, 0x9f, 0xc7, 0xf7, 0x43, 0xd8, 0x2f, 0x11,
	0xd1, 0xc9, 0x39, 0x2a, 0x47, 0x4b, 0x2c, 0x17, 0xdd, 0x79, 0x74, 0xd7, 0x07, 0xd3, 0xa7, 0x90,
	0x62, 0x85, 0x7c, 0xca, 0xa4, 0xad, 0x4c, 0xf4, 0x6c, 0x13, 0x3a, 0x20, 0x65, 0xcf, 0x56, 0x66,
	0x6e, 0x3a, 0x82, 0x41, 0x83, 0x0e, 0xb6, 0xcb, 0xc4, 0xf6, 0x22, 0x1b, 0x7c, 0xb7, 0xa1, 0xdf,
	0x20, 0x2d, 0x47, 0x95, 0xad, 0x6c, 0xb5, 0x46, 0xdd, 0xa2, 0x1b, 0xc1, 0x82, 0xa3, 0xf2, 0xdc,
	0x8c, 0xa3, 0x38, 0x66, 0x92, 0x23, 0x0f, 0xdc, 0x6a, 0xe0, 0x28, 0xbc, 0xc7, 0x91, 0x13, 0x97,
	0xc3, 0x7f, 0xc1, 0x2f, 0x16, 0x45, 0xec, 0x1a, 0xb1, 0x9b, 0x24, 0x1d, 0xd6, 0x0a, 0xf1, 0xaf,
	0xe1, 0x8e, 0xb1, 0x8a, 0xc5, 0x3e, 0xa2, 0x38, 0x56, 0x32, 0xd6, 0xb7, 0x4e, 0x39, 0xdf, 0x32,
	0x56, 0xed, 0x13, 0xf0, 0x3e, 0xe8, 0xf3, 0x32, 0x5f, 0x42, 0x76, 0xc5, 0xc7, 0xa1, 0xdc, 0x0e,
	0x7d, 0x7a, 0xe3, 0xf2, 0xa7, 0x54, 0xf5, 0xf0, 0x4f, 0x0b, 0x7a, 0xc7, 0x88, 0xc6, 0x2e, 0x5a,
	0xf6, 0x18, 0x36, 0xed, 0x17, 0x87, 0xcc, 0x8a, 0x33, 0x79, 0xb9, 0x67, 0x5e, 0x28, 0xc4, 0x59,
	0x5c, 0x77, 0x1b, 0xfa, 0x0b, 0xf6, 0x42, 0xd3, 0xe6, 0x64, 0xd8, 0xdc, 0x07, 0xb0, 0x71, 0x45,
	0xbb, 0x12, 0xd9, 0xe8, 0xd4, 0x3d, 0x80, 0x7f, 0x7a, 0xd4, 0x91, 0xb1, 0x3d, 0x3e, 0x2b, 0xe5,
	0x0c, 0x73, 0xaa, 0x5c, 0x4c, 0xd2, 0x4a, 0x9d, 0x95, 0x72, 0xe6, 0x40, 0x95, 0xd8, 0xcc, 0x2a,
	0xb2, 0xc1, 0x6f, 0xb5, 0xce, 0xaa, 0x26, 0x43, 0xf1, 0xbf, 0xda, 0x90, 0x88, 0x29, 0x77, 0x2e,
	0x14, 0x9f, 0xbe, 0x82, 0x4c, 0x54, 0xa7, 0x7e, 0xa0, 0xd9, 0x19, 0x9f, 0x6a, 0xa9, 0xf1, 0x9c,
	0x4d, 0x34, 0xce, 0xb8, 0x3b, 0xa9, 0x37, 0xe0, 0x66, 0xad, 0x7f, 0xac, 0xe5, 0xdd, 0xa0, 0xfa,
	0xe4, 0x83, 0x51, 0xe3, 0x68, 0x74, 0x28, 0x42, 0xa7, 0xe7, 0x36, 0xac, 0x07, 0x59, 0x4b, 0x2a,
	0xbd, 0x5b, 0xac, 0xd1, 0xfb, 0x5b, 0x99, 0xfe, 0x68, 0x41, 0xf7, 0x48, 0x95, 0xca, 0xf2, 0x69,
	0xc8, 0x82, 0x4a, 0x4f, 0xc6, 0x93, 0xfc, 0x3a, 0x8e, 0x7d, 0xde, 0x3c, 0x9c, 0xc5, 0x46, 0xbd,
	0xf0, 0x01, 0x55, 0xff, 0xbd, 0x05, 0x49, 0x63, 0x14, 0x68, 0x73, 0x93, 0xb1, 0xbc, 0x9e, 0x3c,
	0x2e, 0xce, 0x5c, 0x01, 0xf4, 0x4e, 0x79, 0x0c, 0x7f, 0x2e, 0x41, 0x27, 0x5e, 0x4d, 0x17, 0x6f,
	0xa5, 0xf1, 0xa5, 0x5b, 0xe9, 0x3e, 0x24, 0xf5, 0xc2, 0x24, 0xef, 0x90, 0x0c, 0x21, 0x44, 0xc0,
	0xff, 0xb0, 0xe2, 0x6d, 0x54, 0xf6, 0x8c, 0xa4, 0xf0, 0x92, 0x3e, 0x81, 0x4d, 0x7a, 0x60, 0x52,
	0x39, 0x61, 0xb5, 0xf1, 0x79, 0x65, 0xcf, 0x89, 0x18, 0x90, 0xb0, 0xb7, 0x88, 0xa7, 0xdf, 0x5a,
	0xf3, 0xde, 0x7a, 0x29, 0x7b, 0xb1, 0xd5, 0x1e, 0x25, 0x63, 0x7e, 0x3d, 0xbb, 0xd2, 0x18, 0xc6,
	0x7a, 0x7c, 0xfc, 0x96, 0x4c, 0x56, 0xe9, 0x8f, 0x60, 0xe7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xdd, 0x43, 0xae, 0xe9, 0x1b, 0x06, 0x00, 0x00,
}