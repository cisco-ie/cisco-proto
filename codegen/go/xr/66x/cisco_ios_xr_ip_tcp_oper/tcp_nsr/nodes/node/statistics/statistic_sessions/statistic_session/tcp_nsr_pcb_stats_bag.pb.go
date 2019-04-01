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
// source: tcp_nsr_pcb_stats_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_nsr_nodes_node_statistics_statistic_sessions_statistic_session

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

type TcpNsrPcbStatsBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 string   `protobuf:"bytes,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrPcbStatsBag_KEYS) Reset()         { *m = TcpNsrPcbStatsBag_KEYS{} }
func (m *TcpNsrPcbStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbStatsBag_KEYS) ProtoMessage()    {}
func (*TcpNsrPcbStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e700e9063a6e0d7, []int{0}
}

func (m *TcpNsrPcbStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbStatsBag_KEYS.Unmarshal(m, b)
}
func (m *TcpNsrPcbStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbStatsBag_KEYS.Merge(m, src)
}
func (m *TcpNsrPcbStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbStatsBag_KEYS.Size(m)
}
func (m *TcpNsrPcbStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbStatsBag_KEYS proto.InternalMessageInfo

func (m *TcpNsrPcbStatsBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpNsrPcbStatsBag_KEYS) GetId_1() string {
	if m != nil {
		return m.Id_1
	}
	return ""
}

type TcpNsrSndCountersNode struct {
	DataXferSend                   uint32   `protobuf:"varint,1,opt,name=data_xfer_send,json=dataXferSend,proto3" json:"data_xfer_send,omitempty"`
	DataXferSendTotal              uint64   `protobuf:"varint,2,opt,name=data_xfer_send_total,json=dataXferSendTotal,proto3" json:"data_xfer_send_total,omitempty"`
	DataXferSendDrop               uint32   `protobuf:"varint,3,opt,name=data_xfer_send_drop,json=dataXferSendDrop,proto3" json:"data_xfer_send_drop,omitempty"`
	DataXferSendIovAlloc           uint32   `protobuf:"varint,4,opt,name=data_xfer_send_iov_alloc,json=dataXferSendIovAlloc,proto3" json:"data_xfer_send_iov_alloc,omitempty"`
	DataXferRcv                    uint32   `protobuf:"varint,5,opt,name=data_xfer_rcv,json=dataXferRcv,proto3" json:"data_xfer_rcv,omitempty"`
	DataXferRcvSuccess             uint32   `protobuf:"varint,6,opt,name=data_xfer_rcv_success,json=dataXferRcvSuccess,proto3" json:"data_xfer_rcv_success,omitempty"`
	DataXferRcvFailBufferTrim      uint32   `protobuf:"varint,7,opt,name=data_xfer_rcv_fail_buffer_trim,json=dataXferRcvFailBufferTrim,proto3" json:"data_xfer_rcv_fail_buffer_trim,omitempty"`
	DataXferRcvFailSndUnaOutOfSync uint32   `protobuf:"varint,8,opt,name=data_xfer_rcv_fail_snd_una_out_of_sync,json=dataXferRcvFailSndUnaOutOfSync,proto3" json:"data_xfer_rcv_fail_snd_una_out_of_sync,omitempty"`
	SegInstrSend                   uint32   `protobuf:"varint,9,opt,name=seg_instr_send,json=segInstrSend,proto3" json:"seg_instr_send,omitempty"`
	SegInstrSendUnits              uint32   `protobuf:"varint,10,opt,name=seg_instr_send_units,json=segInstrSendUnits,proto3" json:"seg_instr_send_units,omitempty"`
	SegInstrSendDrop               uint32   `protobuf:"varint,11,opt,name=seg_instr_send_drop,json=segInstrSendDrop,proto3" json:"seg_instr_send_drop,omitempty"`
	SegInstrRcv                    uint32   `protobuf:"varint,12,opt,name=seg_instr_rcv,json=segInstrRcv,proto3" json:"seg_instr_rcv,omitempty"`
	SegInstrRcvSuccess             uint32   `protobuf:"varint,13,opt,name=seg_instr_rcv_success,json=segInstrRcvSuccess,proto3" json:"seg_instr_rcv_success,omitempty"`
	SegInstrRcvFailBufferTrim      uint32   `protobuf:"varint,14,opt,name=seg_instr_rcv_fail_buffer_trim,json=segInstrRcvFailBufferTrim,proto3" json:"seg_instr_rcv_fail_buffer_trim,omitempty"`
	SegInstrRcvFailTcpProcess      uint32   `protobuf:"varint,15,opt,name=seg_instr_rcv_fail_tcp_process,json=segInstrRcvFailTcpProcess,proto3" json:"seg_instr_rcv_fail_tcp_process,omitempty"`
	NackSend                       uint32   `protobuf:"varint,16,opt,name=nack_send,json=nackSend,proto3" json:"nack_send,omitempty"`
	NackSendDrop                   uint32   `protobuf:"varint,17,opt,name=nack_send_drop,json=nackSendDrop,proto3" json:"nack_send_drop,omitempty"`
	NackRcv                        uint32   `protobuf:"varint,18,opt,name=nack_rcv,json=nackRcv,proto3" json:"nack_rcv,omitempty"`
	NackRcvSuccess                 uint32   `protobuf:"varint,19,opt,name=nack_rcv_success,json=nackRcvSuccess,proto3" json:"nack_rcv_success,omitempty"`
	NackRcvFailDataSend            uint32   `protobuf:"varint,20,opt,name=nack_rcv_fail_data_send,json=nackRcvFailDataSend,proto3" json:"nack_rcv_fail_data_send,omitempty"`
	CleanupSend                    uint32   `protobuf:"varint,21,opt,name=cleanup_send,json=cleanupSend,proto3" json:"cleanup_send,omitempty"`
	CleanupSendDrop                uint32   `protobuf:"varint,22,opt,name=cleanup_send_drop,json=cleanupSendDrop,proto3" json:"cleanup_send_drop,omitempty"`
	CleanupRcv                     uint32   `protobuf:"varint,23,opt,name=cleanup_rcv,json=cleanupRcv,proto3" json:"cleanup_rcv,omitempty"`
	CleanupRcvSuccess              uint32   `protobuf:"varint,24,opt,name=cleanup_rcv_success,json=cleanupRcvSuccess,proto3" json:"cleanup_rcv_success,omitempty"`
	CleanupRcvFailBufferTrim       uint32   `protobuf:"varint,25,opt,name=cleanup_rcv_fail_buffer_trim,json=cleanupRcvFailBufferTrim,proto3" json:"cleanup_rcv_fail_buffer_trim,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *TcpNsrSndCountersNode) Reset()         { *m = TcpNsrSndCountersNode{} }
func (m *TcpNsrSndCountersNode) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSndCountersNode) ProtoMessage()    {}
func (*TcpNsrSndCountersNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e700e9063a6e0d7, []int{1}
}

func (m *TcpNsrSndCountersNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSndCountersNode.Unmarshal(m, b)
}
func (m *TcpNsrSndCountersNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSndCountersNode.Marshal(b, m, deterministic)
}
func (m *TcpNsrSndCountersNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSndCountersNode.Merge(m, src)
}
func (m *TcpNsrSndCountersNode) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSndCountersNode.Size(m)
}
func (m *TcpNsrSndCountersNode) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSndCountersNode.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSndCountersNode proto.InternalMessageInfo

func (m *TcpNsrSndCountersNode) GetDataXferSend() uint32 {
	if m != nil {
		return m.DataXferSend
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferSendTotal() uint64 {
	if m != nil {
		return m.DataXferSendTotal
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferSendDrop() uint32 {
	if m != nil {
		return m.DataXferSendDrop
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferSendIovAlloc() uint32 {
	if m != nil {
		return m.DataXferSendIovAlloc
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferRcv() uint32 {
	if m != nil {
		return m.DataXferRcv
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferRcvSuccess() uint32 {
	if m != nil {
		return m.DataXferRcvSuccess
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferRcvFailBufferTrim() uint32 {
	if m != nil {
		return m.DataXferRcvFailBufferTrim
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetDataXferRcvFailSndUnaOutOfSync() uint32 {
	if m != nil {
		return m.DataXferRcvFailSndUnaOutOfSync
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrSend() uint32 {
	if m != nil {
		return m.SegInstrSend
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrSendUnits() uint32 {
	if m != nil {
		return m.SegInstrSendUnits
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrSendDrop() uint32 {
	if m != nil {
		return m.SegInstrSendDrop
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrRcv() uint32 {
	if m != nil {
		return m.SegInstrRcv
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrRcvSuccess() uint32 {
	if m != nil {
		return m.SegInstrRcvSuccess
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrRcvFailBufferTrim() uint32 {
	if m != nil {
		return m.SegInstrRcvFailBufferTrim
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetSegInstrRcvFailTcpProcess() uint32 {
	if m != nil {
		return m.SegInstrRcvFailTcpProcess
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetNackSend() uint32 {
	if m != nil {
		return m.NackSend
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetNackSendDrop() uint32 {
	if m != nil {
		return m.NackSendDrop
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetNackRcv() uint32 {
	if m != nil {
		return m.NackRcv
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetNackRcvSuccess() uint32 {
	if m != nil {
		return m.NackRcvSuccess
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetNackRcvFailDataSend() uint32 {
	if m != nil {
		return m.NackRcvFailDataSend
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetCleanupSend() uint32 {
	if m != nil {
		return m.CleanupSend
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetCleanupSendDrop() uint32 {
	if m != nil {
		return m.CleanupSendDrop
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetCleanupRcv() uint32 {
	if m != nil {
		return m.CleanupRcv
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetCleanupRcvSuccess() uint32 {
	if m != nil {
		return m.CleanupRcvSuccess
	}
	return 0
}

func (m *TcpNsrSndCountersNode) GetCleanupRcvFailBufferTrim() uint32 {
	if m != nil {
		return m.CleanupRcvFailBufferTrim
	}
	return 0
}

type TcpNsrPcbStatsBag struct {
	Pcb                                uint64                 `protobuf:"varint,50,opt,name=pcb,proto3" json:"pcb,omitempty"`
	NumberOfTimesNsrUp                 uint32                 `protobuf:"varint,51,opt,name=number_of_times_nsr_up,json=numberOfTimesNsrUp,proto3" json:"number_of_times_nsr_up,omitempty"`
	NumberOfTimersNsrDown              uint32                 `protobuf:"varint,52,opt,name=number_of_timers_nsr_down,json=numberOfTimersNsrDown,proto3" json:"number_of_timers_nsr_down,omitempty"`
	NumberOfTimesNsrDisabled           uint32                 `protobuf:"varint,53,opt,name=number_of_times_nsr_disabled,json=numberOfTimesNsrDisabled,proto3" json:"number_of_times_nsr_disabled,omitempty"`
	NumberOfTimesNsrFailOver           uint32                 `protobuf:"varint,54,opt,name=number_of_times_nsr_fail_over,json=numberOfTimesNsrFailOver,proto3" json:"number_of_times_nsr_fail_over,omitempty"`
	InternalAckDropsNotReplicated      uint64                 `protobuf:"varint,55,opt,name=internal_ack_drops_not_replicated,json=internalAckDropsNotReplicated,proto3" json:"internal_ack_drops_not_replicated,omitempty"`
	InternalAckDropsInitsyncFirstPhase uint64                 `protobuf:"varint,56,opt,name=internal_ack_drops_initsync_first_phase,json=internalAckDropsInitsyncFirstPhase,proto3" json:"internal_ack_drops_initsync_first_phase,omitempty"`
	InternalAckDropsStale              uint64                 `protobuf:"varint,57,opt,name=internal_ack_drops_stale,json=internalAckDropsStale,proto3" json:"internal_ack_drops_stale,omitempty"`
	InternalAckDropsImmediateMatch     uint64                 `protobuf:"varint,58,opt,name=internal_ack_drops_immediate_match,json=internalAckDropsImmediateMatch,proto3" json:"internal_ack_drops_immediate_match,omitempty"`
	SndCounters                        *TcpNsrSndCountersNode `protobuf:"bytes,59,opt,name=snd_counters,json=sndCounters,proto3" json:"snd_counters,omitempty"`
	LastClearedTime                    uint32                 `protobuf:"varint,60,opt,name=last_cleared_time,json=lastClearedTime,proto3" json:"last_cleared_time,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}               `json:"-"`
	XXX_unrecognized                   []byte                 `json:"-"`
	XXX_sizecache                      int32                  `json:"-"`
}

func (m *TcpNsrPcbStatsBag) Reset()         { *m = TcpNsrPcbStatsBag{} }
func (m *TcpNsrPcbStatsBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrPcbStatsBag) ProtoMessage()    {}
func (*TcpNsrPcbStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e700e9063a6e0d7, []int{2}
}

func (m *TcpNsrPcbStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrPcbStatsBag.Unmarshal(m, b)
}
func (m *TcpNsrPcbStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrPcbStatsBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrPcbStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrPcbStatsBag.Merge(m, src)
}
func (m *TcpNsrPcbStatsBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrPcbStatsBag.Size(m)
}
func (m *TcpNsrPcbStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrPcbStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrPcbStatsBag proto.InternalMessageInfo

func (m *TcpNsrPcbStatsBag) GetPcb() uint64 {
	if m != nil {
		return m.Pcb
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetNumberOfTimesNsrUp() uint32 {
	if m != nil {
		return m.NumberOfTimesNsrUp
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetNumberOfTimersNsrDown() uint32 {
	if m != nil {
		return m.NumberOfTimersNsrDown
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetNumberOfTimesNsrDisabled() uint32 {
	if m != nil {
		return m.NumberOfTimesNsrDisabled
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetNumberOfTimesNsrFailOver() uint32 {
	if m != nil {
		return m.NumberOfTimesNsrFailOver
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetInternalAckDropsNotReplicated() uint64 {
	if m != nil {
		return m.InternalAckDropsNotReplicated
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetInternalAckDropsInitsyncFirstPhase() uint64 {
	if m != nil {
		return m.InternalAckDropsInitsyncFirstPhase
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetInternalAckDropsStale() uint64 {
	if m != nil {
		return m.InternalAckDropsStale
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetInternalAckDropsImmediateMatch() uint64 {
	if m != nil {
		return m.InternalAckDropsImmediateMatch
	}
	return 0
}

func (m *TcpNsrPcbStatsBag) GetSndCounters() *TcpNsrSndCountersNode {
	if m != nil {
		return m.SndCounters
	}
	return nil
}

func (m *TcpNsrPcbStatsBag) GetLastClearedTime() uint32 {
	if m != nil {
		return m.LastClearedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpNsrPcbStatsBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.statistics.statistic_sessions.statistic_session.tcp_nsr_pcb_stats_bag_KEYS")
	proto.RegisterType((*TcpNsrSndCountersNode)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.statistics.statistic_sessions.statistic_session.tcp_nsr_snd_counters_node")
	proto.RegisterType((*TcpNsrPcbStatsBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.statistics.statistic_sessions.statistic_session.tcp_nsr_pcb_stats_bag")
}

func init() { proto.RegisterFile("tcp_nsr_pcb_stats_bag.proto", fileDescriptor_0e700e9063a6e0d7) }

var fileDescriptor_0e700e9063a6e0d7 = []byte{
	// 908 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xdb, 0x6e, 0x1b, 0x37,
	0x10, 0x86, 0xe1, 0x43, 0x9d, 0x98, 0x3e, 0xc4, 0xa2, 0xad, 0x84, 0x6a, 0x1a, 0x37, 0x11, 0x82,
	0xd6, 0x28, 0x50, 0x15, 0x76, 0xd2, 0x24, 0x3d, 0xa0, 0x81, 0x1b, 0x37, 0xa8, 0x5b, 0xd4, 0x0e,
	0x24, 0x1b, 0x68, 0xd1, 0x0b, 0x82, 0x22, 0x29, 0x87, 0xc8, 0x8a, 0x5c, 0x90, 0x5c, 0x25, 0xbe,
	0xef, 0x4b, 0xf4, 0x7d, 0xfa, 0x60, 0xc5, 0x0c, 0x77, 0xa5, 0x95, 0xbc, 0xbd, 0xec, 0x8d, 0xa1,
	0x9d, 0xf9, 0xfe, 0x21, 0xe7, 0x1f, 0x2e, 0xd7, 0xe4, 0x7e, 0x94, 0x39, 0xb7, 0xc1, 0xf3, 0x5c,
	0x0e, 0x79, 0x88, 0x22, 0x06, 0x3e, 0x14, 0x57, 0xbd, 0xdc, 0xbb, 0xe8, 0xe8, 0x9f, 0xd2, 0x04,
	0xe9, 0xb8, 0x71, 0x81, 0x7f, 0xf0, 0xdc, 0xe4, 0x1c, 0x60, 0x97, 0x6b, 0xdf, 0x2b, 0x55, 0x3d,
	0xeb, 0x94, 0x0e, 0xf8, 0xb7, 0x07, 0x62, 0x13, 0xa2, 0x91, 0x61, 0xf6, 0x93, 0x07, 0x1d, 0x82,
	0x71, 0xb6, 0x21, 0xd4, 0x7d, 0x49, 0x3e, 0x6e, 0x5c, 0x9b, 0xff, 0xfa, 0xd3, 0x1f, 0x03, 0xba,
	0x4d, 0x96, 0x8d, 0x62, 0x4b, 0x0f, 0x97, 0x0e, 0xd6, 0xfb, 0xcb, 0x46, 0xd1, 0x16, 0x59, 0x35,
	0x8a, 0x1f, 0xb2, 0x65, 0x8c, 0xac, 0x18, 0x75, 0xd8, 0xfd, 0x67, 0x9d, 0x74, 0xaa, 0x0a, 0xc1,
	0x2a, 0x2e, 0x5d, 0x61, 0xa3, 0xf6, 0x81, 0xc3, 0x76, 0xe8, 0x63, 0xb2, 0xad, 0x44, 0x14, 0xfc,
	0xc3, 0x48, 0x7b, 0x1e, 0xb4, 0x4d, 0xc5, 0xb6, 0xfa, 0x9b, 0x10, 0xfd, 0x7d, 0xa4, 0xfd, 0x40,
	0x5b, 0x45, 0xbf, 0x22, 0x7b, 0xf3, 0x14, 0x8f, 0x2e, 0x8a, 0x0c, 0x97, 0x59, 0xed, 0xb7, 0xea,
	0xec, 0x05, 0x24, 0xe8, 0x97, 0x64, 0x77, 0x41, 0xa0, 0xbc, 0xcb, 0xd9, 0x0a, 0xd6, 0xde, 0xa9,
	0xf3, 0x27, 0xde, 0xe5, 0xf4, 0x19, 0x61, 0x0b, 0xb8, 0x71, 0x13, 0x2e, 0xb2, 0xcc, 0x49, 0xb6,
	0x8a, 0x9a, 0xbd, 0xba, 0xe6, 0xd4, 0x4d, 0x8e, 0x21, 0x47, 0xbb, 0x64, 0x6b, 0xa6, 0xf3, 0x72,
	0xc2, 0x3e, 0x42, 0x78, 0xa3, 0x82, 0xfb, 0x72, 0x42, 0x0f, 0x49, 0x7b, 0x8e, 0xe1, 0xa1, 0x90,
	0x52, 0x87, 0xc0, 0xd6, 0x90, 0xa5, 0x35, 0x76, 0x90, 0x32, 0xf4, 0x98, 0xec, 0xcf, 0x4b, 0x46,
	0xc2, 0x64, 0x7c, 0x58, 0x8c, 0xe0, 0x39, 0x7a, 0x33, 0x66, 0xb7, 0x50, 0xdb, 0xa9, 0x69, 0x5f,
	0x0b, 0x93, 0xfd, 0x88, 0xc4, 0x85, 0x37, 0x63, 0x7a, 0x46, 0x3e, 0x6b, 0x28, 0x01, 0xfe, 0x17,
	0x56, 0x70, 0x57, 0x44, 0xee, 0x46, 0x3c, 0x5c, 0x5b, 0xc9, 0x6e, 0x63, 0xa9, 0xfd, 0x85, 0x52,
	0x03, 0xab, 0x2e, 0xad, 0x38, 0x2f, 0xe2, 0xf9, 0x68, 0x70, 0x6d, 0x25, 0xcc, 0x29, 0xe8, 0x2b,
	0x6e, 0x6c, 0x88, 0xe5, 0x9c, 0xd6, 0xd3, 0x9c, 0x82, 0xbe, 0x3a, 0x85, 0x60, 0x35, 0xa7, 0x79,
	0x8a, 0x17, 0xd6, 0xc4, 0xc0, 0x08, 0xb2, 0xad, 0x3a, 0x7b, 0x09, 0x09, 0x98, 0xd3, 0x82, 0x00,
	0xe7, 0xb4, 0x91, 0xe6, 0x54, 0xe7, 0x71, 0x4e, 0x5d, 0xb2, 0x35, 0xc3, 0xc1, 0xef, 0xcd, 0xe4,
	0x77, 0x05, 0x96, 0x7e, 0xcf, 0x31, 0x53, 0xbf, 0xb7, 0x92, 0xdf, 0x35, 0xb6, 0xe6, 0xf7, 0xbc,
	0xe4, 0x86, 0xdf, 0xdb, 0xc9, 0xef, 0x9a, 0x76, 0xc1, 0xef, 0xe6, 0x12, 0x70, 0xee, 0x73, 0xef,
	0x70, 0xf9, 0x3b, 0x8d, 0x25, 0x2e, 0x64, 0xfe, 0x26, 0x01, 0xf4, 0x3e, 0x59, 0xb7, 0x42, 0xbe,
	0x4b, 0xee, 0xee, 0x20, 0x7d, 0x1b, 0x02, 0xe8, 0xec, 0x63, 0xb2, 0x3d, 0x4d, 0x26, 0x8f, 0x5a,
	0xc9, 0xff, 0x8a, 0x40, 0x7f, 0x3a, 0x04, 0x15, 0x68, 0x0d, 0xc5, 0xfc, 0x2d, 0x78, 0x06, 0x5b,
	0x0e, 0xc8, 0x4e, 0x95, 0x9a, 0x3a, 0xb2, 0x8b, 0xc8, 0x76, 0x89, 0x54, 0x6e, 0x3c, 0x25, 0xf7,
	0xa6, 0x24, 0x76, 0x81, 0x07, 0x09, 0x77, 0xb5, 0x87, 0x82, 0xdd, 0x52, 0x00, 0xfb, 0x3f, 0x11,
	0x51, 0xe0, 0x06, 0x1f, 0x91, 0x4d, 0x99, 0x69, 0x61, 0x8b, 0x3c, 0xa1, 0xed, 0x34, 0x99, 0x32,
	0x86, 0xc8, 0x17, 0xa4, 0x55, 0x47, 0x52, 0x1b, 0x77, 0x91, 0xbb, 0x53, 0xe3, 0xb0, 0x93, 0x4f,
	0x49, 0x25, 0xc5, 0x66, 0xee, 0x21, 0x45, 0xca, 0x10, 0xf4, 0xd3, 0x23, 0xbb, 0x35, 0x60, 0xda,
	0x12, 0x4b, 0x27, 0x6d, 0x06, 0x56, 0x5d, 0xfd, 0x40, 0x3e, 0xa9, 0xf3, 0x37, 0x26, 0xdc, 0x41,
	0x21, 0x9b, 0x09, 0xe7, 0x07, 0xdc, 0xfd, 0x6b, 0x8d, 0xb4, 0x1b, 0x2f, 0x42, 0xba, 0x43, 0x56,
	0x72, 0x39, 0x64, 0x47, 0x78, 0x17, 0xc1, 0x4f, 0x7a, 0x44, 0xee, 0xda, 0x62, 0x3c, 0xd4, 0x1e,
	0x5e, 0xb2, 0x68, 0xc6, 0x3a, 0xa0, 0xac, 0xc8, 0xd9, 0x93, 0x74, 0x06, 0x53, 0xf6, 0x7c, 0x74,
	0x01, 0xb9, 0xb3, 0xe0, 0x2f, 0x73, 0xfa, 0x82, 0x74, 0xe6, 0x35, 0x3e, 0x89, 0x94, 0x7b, 0x6f,
	0xd9, 0x53, 0x94, 0xb5, 0xeb, 0x32, 0x0f, 0xba, 0x13, 0xf7, 0xde, 0x42, 0x67, 0x4d, 0xab, 0x29,
	0x13, 0xc4, 0x30, 0xd3, 0x8a, 0x7d, 0x9d, 0x3a, 0x5b, 0x5c, 0xf3, 0xa4, 0xcc, 0xd3, 0x97, 0xe4,
	0x41, 0x93, 0x1e, 0x1d, 0x72, 0x13, 0xed, 0xd9, 0xb3, 0xe6, 0x02, 0x60, 0xd0, 0xf9, 0x44, 0x7b,
	0xfa, 0x33, 0x79, 0x64, 0xe0, 0x4a, 0xb7, 0x22, 0xe3, 0x70, 0x70, 0x60, 0xae, 0x70, 0xbd, 0x47,
	0xee, 0x75, 0x9e, 0x19, 0x29, 0xa2, 0x56, 0xec, 0x39, 0xda, 0xf3, 0xa0, 0x02, 0x8f, 0xe5, 0x3b,
	0x98, 0x73, 0x38, 0x73, 0xb1, 0x3f, 0x85, 0xe8, 0x80, 0x7c, 0xde, 0x50, 0xc9, 0xc0, 0x55, 0x71,
	0x6d, 0x25, 0x1f, 0x19, 0x1f, 0x22, 0xcf, 0xdf, 0x8a, 0xa0, 0xd9, 0x0b, 0xac, 0xd7, 0x5d, 0xac,
	0x77, 0x5a, 0xb2, 0xaf, 0x01, 0x7d, 0x03, 0x24, 0x7d, 0x4e, 0x58, 0x43, 0xd1, 0x10, 0x45, 0xa6,
	0xd9, 0x37, 0x58, 0xa5, 0xbd, 0x58, 0x65, 0x00, 0x49, 0xfa, 0x0b, 0xe9, 0x36, 0xed, 0x66, 0x3c,
	0xd6, 0xca, 0x88, 0xa8, 0xf9, 0x58, 0x44, 0xf9, 0x96, 0x7d, 0x8b, 0x25, 0xf6, 0x6f, 0x6c, 0xa4,
	0xc2, 0x7e, 0x03, 0x8a, 0xfe, 0xbd, 0x44, 0x36, 0xeb, 0x5f, 0x3f, 0xf6, 0xdd, 0xc3, 0xa5, 0x83,
	0x8d, 0xa3, 0x49, 0xef, 0x7f, 0xfc, 0x76, 0xf7, 0xfe, 0xf3, 0xb3, 0xdb, 0xdf, 0x08, 0x56, 0xbd,
	0x2a, 0x23, 0xf0, 0x5e, 0x66, 0x22, 0x44, 0x0e, 0x67, 0xdf, 0x6b, 0x85, 0x67, 0x80, 0x7d, 0x9f,
	0xde, 0x4b, 0x48, 0xbc, 0x4a, 0x71, 0x98, 0xfb, 0x70, 0x0d, 0xff, 0xe5, 0x78, 0xf2, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x0e, 0x9f, 0x1a, 0x91, 0x08, 0x00, 0x00,
}
