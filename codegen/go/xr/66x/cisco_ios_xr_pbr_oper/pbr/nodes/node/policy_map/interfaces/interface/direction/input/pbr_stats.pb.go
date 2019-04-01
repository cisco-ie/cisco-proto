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

type HttpEnrichStatsSt struct {
	RqstRcvdPackets      uint64   `protobuf:"varint,1,opt,name=rqst_rcvd_packets,json=rqstRcvdPackets,proto3" json:"rqst_rcvd_packets,omitempty"`
	RqstRcvdBytes        uint64   `protobuf:"varint,2,opt,name=rqst_rcvd_bytes,json=rqstRcvdBytes,proto3" json:"rqst_rcvd_bytes,omitempty"`
	DropPackets          uint64   `protobuf:"varint,3,opt,name=drop_packets,json=dropPackets,proto3" json:"drop_packets,omitempty"`
	DropBytes            uint64   `protobuf:"varint,4,opt,name=drop_bytes,json=dropBytes,proto3" json:"drop_bytes,omitempty"`
	RespSentPackets      uint64   `protobuf:"varint,5,opt,name=resp_sent_packets,json=respSentPackets,proto3" json:"resp_sent_packets,omitempty"`
	RespSentBytes        uint64   `protobuf:"varint,6,opt,name=resp_sent_bytes,json=respSentBytes,proto3" json:"resp_sent_bytes,omitempty"`
	ReqSentPackets       uint64   `protobuf:"varint,7,opt,name=req_sent_packets,json=reqSentPackets,proto3" json:"req_sent_packets,omitempty"`
	TcpSentPackets       uint64   `protobuf:"varint,8,opt,name=tcp_sent_packets,json=tcpSentPackets,proto3" json:"tcp_sent_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpEnrichStatsSt) Reset()         { *m = HttpEnrichStatsSt{} }
func (m *HttpEnrichStatsSt) String() string { return proto.CompactTextString(m) }
func (*HttpEnrichStatsSt) ProtoMessage()    {}
func (*HttpEnrichStatsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{3}
}

func (m *HttpEnrichStatsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpEnrichStatsSt.Unmarshal(m, b)
}
func (m *HttpEnrichStatsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpEnrichStatsSt.Marshal(b, m, deterministic)
}
func (m *HttpEnrichStatsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpEnrichStatsSt.Merge(m, src)
}
func (m *HttpEnrichStatsSt) XXX_Size() int {
	return xxx_messageInfo_HttpEnrichStatsSt.Size(m)
}
func (m *HttpEnrichStatsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpEnrichStatsSt.DiscardUnknown(m)
}

var xxx_messageInfo_HttpEnrichStatsSt proto.InternalMessageInfo

func (m *HttpEnrichStatsSt) GetRqstRcvdPackets() uint64 {
	if m != nil {
		return m.RqstRcvdPackets
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetRqstRcvdBytes() uint64 {
	if m != nil {
		return m.RqstRcvdBytes
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetDropPackets() uint64 {
	if m != nil {
		return m.DropPackets
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetDropBytes() uint64 {
	if m != nil {
		return m.DropBytes
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetRespSentPackets() uint64 {
	if m != nil {
		return m.RespSentPackets
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetRespSentBytes() uint64 {
	if m != nil {
		return m.RespSentBytes
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetReqSentPackets() uint64 {
	if m != nil {
		return m.ReqSentPackets
	}
	return 0
}

func (m *HttpEnrichStatsSt) GetTcpSentPackets() uint64 {
	if m != nil {
		return m.TcpSentPackets
	}
	return 0
}

type ClassStats struct {
	CounterValidityBitmask uint64             `protobuf:"varint,1,opt,name=counter_validity_bitmask,json=counterValidityBitmask,proto3" json:"counter_validity_bitmask,omitempty"`
	ClassName              string             `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ClassId                uint32             `protobuf:"varint,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	GeneralStats           *GenStatsSt        `protobuf:"bytes,4,opt,name=general_stats,json=generalStats,proto3" json:"general_stats,omitempty"`
	HttprStats             *HttprStatsSt      `protobuf:"bytes,5,opt,name=httpr_stats,json=httprStats,proto3" json:"httpr_stats,omitempty"`
	HttpEnrichStats        *HttpEnrichStatsSt `protobuf:"bytes,6,opt,name=http_enrich_stats,json=httpEnrichStats,proto3" json:"http_enrich_stats,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *ClassStats) Reset()         { *m = ClassStats{} }
func (m *ClassStats) String() string { return proto.CompactTextString(m) }
func (*ClassStats) ProtoMessage()    {}
func (*ClassStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_6233b79aff5b417f, []int{4}
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

func (m *ClassStats) GetHttpEnrichStats() *HttpEnrichStatsSt {
	if m != nil {
		return m.HttpEnrichStats
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
	return fileDescriptor_6233b79aff5b417f, []int{5}
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
	proto.RegisterType((*HttpEnrichStatsSt)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.http_enrich_stats_st")
	proto.RegisterType((*ClassStats)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.class_stats")
	proto.RegisterType((*PbrStats)(nil), "cisco_ios_xr_pbr_oper.pbr.nodes.node.policy_map.interfaces.interface.direction.input.pbr_stats")
}

func init() { proto.RegisterFile("pbr_stats.proto", fileDescriptor_6233b79aff5b417f) }

var fileDescriptor_6233b79aff5b417f = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0xcb, 0x6e, 0x13, 0x3d,
	0x14, 0xd6, 0xb4, 0x69, 0x9b, 0x9c, 0x69, 0x6e, 0xf3, 0xf7, 0x87, 0x01, 0x84, 0x28, 0x91, 0xa8,
	0xc2, 0x45, 0xb3, 0x48, 0xb9, 0x49, 0xec, 0xaa, 0x76, 0x81, 0x10, 0xa8, 0x9a, 0x56, 0x48, 0xac,
	0x2c, 0xc7, 0x63, 0x1a, 0xd3, 0x64, 0xc6, 0xb5, 0xdd, 0x8a, 0xee, 0xd8, 0x20, 0x96, 0xec, 0x78,
	0x0a, 0xc4, 0xd3, 0xf0, 0x34, 0xac, 0x90, 0x8f, 0x27, 0xce, 0x24, 0xed, 0x36, 0x2b, 0x36, 0xa3,
	0xf1, 0xf9, 0x3e, 0x7f, 0xe7, 0xe6, 0x63, 0x43, 0x5b, 0x0e, 0x15, 0xd1, 0x86, 0x1a, 0x9d, 0x48,
	0x55, 0x98, 0x22, 0x3a, 0x66, 0x42, 0xb3, 0x82, 0x88, 0x42, 0x93, 0xcf, 0x8a, 0x58, 0xb4, 0x90,
	0x5c, 0x25, 0x72, 0xa8, 0x92, 0xbc, 0xc8, 0xb8, 0xc6, 0x6f, 0x22, 0x8b, 0xb1, 0x60, 0x97, 0x64,
	0x42, 0x65, 0x22, 0x72, 0xc3, 0xd5, 0x47, 0xca, 0xb8, 0x9e, 0xfd, 0x26, 0x99, 0x50, 0x9c, 0x19,
	0x51, 0xe4, 0x89, 0xc8, 0xe5, 0xb9, 0xe9, 0x1d, 0x43, 0xcb, 0x3b, 0x22, 0x6f, 0x0e, 0x3e, 0x1c,
	0x45, 0x77, 0xa0, 0x61, 0x95, 0x48, 0x4e, 0x27, 0x3c, 0x0e, 0xb6, 0x83, 0x7e, 0x23, 0xad, 0x5b,
	0xc3, 0x3b, 0x3a, 0xe1, 0xd1, 0x03, 0x68, 0x79, 0x2d, 0xc7, 0x58, 0x41, 0x46, 0xd3, 0x5b, 0x2d,
	0xad, 0xf7, 0x6b, 0x15, 0x36, 0x4f, 0x78, 0x5e, 0xca, 0x6a, 0x13, 0x3d, 0x84, 0x8e, 0x51, 0x34,
	0xd7, 0x13, 0x61, 0x88, 0xa4, 0xec, 0x94, 0x1b, 0x8d, 0xda, 0xb5, 0xb4, 0x3d, 0xb5, 0x1f, 0x3a,
	0xb3, 0x75, 0xe1, 0xa9, 0xc3, 0x4b, 0xc3, 0x35, 0xba, 0xa8, 0xa5, 0xcd, 0xa9, 0x75, 0xcf, 0x1a,
	0xa3, 0x27, 0x10, 0x99, 0xc2, 0xd0, 0x31, 0xc9, 0x54, 0x21, 0xbd, 0xe6, 0x2a, 0x52, 0x3b, 0x88,
	0xec, 0xab, 0x42, 0x4e, 0x45, 0xfb, 0xd0, 0xa9, 0xb0, 0x9d, 0x6c, 0x0d, 0xb9, 0x2d, 0xcf, 0x75,
	0xba, 0x3b, 0xd0, 0xae, 0x30, 0x15, 0x35, 0x3c, 0x5e, 0xdb, 0x0e, 0xfa, 0xcd, 0xb4, 0xe9, 0x89,
	0x29, 0x35, 0xdc, 0xf2, 0x26, 0xd4, 0xb0, 0x11, 0xc9, 0xa8, 0xa1, 0x8e, 0xb7, 0xee, 0x78, 0x68,
	0xde, 0xa7, 0x86, 0x22, 0x2f, 0x81, 0xff, 0x9c, 0x9e, 0x4f, 0x0a, 0xb9, 0x1b, 0xc8, 0xed, 0x22,
	0x74, 0x5c, 0x22, 0xc8, 0x7f, 0x05, 0xb7, 0xa5, 0xe2, 0xc4, 0xf7, 0xd1, 0xb0, 0x11, 0xcf, 0x7c,
	0x7e, 0x75, 0x8c, 0xf9, 0xa6, 0x54, 0xfc, 0x10, 0x09, 0x6f, 0x1d, 0x3e, 0x4d, 0xf3, 0x05, 0xc4,
	0xd7, 0x6c, 0x76, 0xe9, 0x36, 0x70, 0xeb, 0xff, 0x8b, 0x5b, 0x31, 0xeb, 0xde, 0x9f, 0x00, 0x5a,
	0x23, 0x63, 0xa4, 0x9a, 0xb5, 0xec, 0x11, 0x74, 0xd5, 0x99, 0x36, 0x44, 0xb1, 0x8b, 0x6c, 0xb1,
	0x67, 0x16, 0x48, 0xd9, 0x85, 0xf7, 0xbb, 0x03, 0xed, 0x19, 0x77, 0xae, 0x69, 0x53, 0xa6, 0x2b,
	0xee, 0x7d, 0xd8, 0xbc, 0xa6, 0x5d, 0x61, 0x56, 0xe9, 0xd4, 0x5d, 0x80, 0x2b, 0x3d, 0x6a, 0x64,
	0xbe, 0x3d, 0x36, 0x2a, 0xae, 0x25, 0xd1, 0x3c, 0x9f, 0x9d, 0xa4, 0xb5, 0x32, 0x2a, 0xae, 0xe5,
	0x11, 0xcf, 0x4d, 0x35, 0x2a, 0xcf, 0x75, 0x7a, 0xeb, 0x65, 0x54, 0x25, 0xd3, 0x25, 0xff, 0x7b,
	0x05, 0xb6, 0x6c, 0xf2, 0x84, 0xe7, 0x4a, 0xb0, 0xd1, 0x3f, 0x58, 0x02, 0x3b, 0x1f, 0x8a, 0x9f,
	0xcd, 0x4b, 0x6e, 0xb8, 0xf9, 0x50, 0xfc, 0xac, 0xaa, 0x68, 0x27, 0x89, 0x2d, 0x38, 0xaf, 0x97,
	0x93, 0xc4, 0xaa, 0xbe, 0x7b, 0x3f, 0x6b, 0x10, 0xb2, 0x31, 0xd5, 0xda, 0x15, 0x34, 0x7a, 0x09,
	0x31, 0x2b, 0xce, 0xed, 0x3d, 0x41, 0x2e, 0xe8, 0x58, 0x64, 0xc2, 0x5c, 0x92, 0xa1, 0x30, 0x13,
	0xaa, 0x4f, 0xcb, 0xa2, 0xde, 0x28, 0xf1, 0xf7, 0x25, 0xbc, 0xe7, 0x50, 0x5b, 0x10, 0x27, 0x54,
	0xb9, 0x71, 0x1a, 0x68, 0xc1, 0x4b, 0xe9, 0x16, 0xd4, 0x1d, 0x2c, 0x32, 0x2c, 0x67, 0x33, 0xdd,
	0xc0, 0xf5, 0xeb, 0x2c, 0xfa, 0x16, 0x40, 0xf3, 0x84, 0xe7, 0x5c, 0xd1, 0xb1, 0x8b, 0x02, 0xcb,
	0x19, 0x0e, 0x86, 0xc9, 0x32, 0x6e, 0xd3, 0xa4, 0x7a, 0xe7, 0xa5, 0x9b, 0xa5, 0xe3, 0x23, 0xcc,
	0xfe, 0x6b, 0x00, 0x61, 0x65, 0xc2, 0xb0, 0x61, 0xe1, 0x20, 0x5b, 0x4e, 0x1c, 0xf3, 0xa3, 0x9c,
	0x02, 0xae, 0x5d, 0x1c, 0x3f, 0x02, 0xe8, 0x5e, 0x39, 0xec, 0x78, 0x28, 0xc2, 0xc1, 0xa7, 0xe5,
	0x45, 0xb3, 0x38, 0x5b, 0x69, 0xdb, 0x5a, 0x0f, 0xd0, 0x88, 0x81, 0xf5, 0xbe, 0xaf, 0x40, 0xc3,
	0x3f, 0x45, 0xf3, 0xaf, 0xd0, 0x60, 0xe1, 0x15, 0xba, 0x07, 0x61, 0x19, 0x03, 0xc2, 0xbb, 0x08,
	0x83, 0x33, 0x21, 0x61, 0x0b, 0xd6, 0xac, 0x0c, 0x8f, 0x9f, 0x22, 0xe4, 0x16, 0xd1, 0x63, 0xe8,
	0xe2, 0x0f, 0xc9, 0xb8, 0x66, 0x4a, 0x48, 0x1b, 0x62, 0xfc, 0x0c, 0x19, 0x1d, 0x04, 0xf6, 0x67,
	0xf6, 0xe8, 0x4b, 0x30, 0x3d, 0x74, 0x16, 0x8a, 0x9f, 0x6f, 0xaf, 0xf6, 0xc3, 0x01, 0x5d, 0x4e,
	0x81, 0x2a, 0x53, 0x52, 0x9e, 0x6b, 0x5b, 0x92, 0xe1, 0x3a, 0x3e, 0xfc, 0xbb, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x55, 0x12, 0x03, 0x0b, 0x0b, 0x08, 0x00, 0x00,
}