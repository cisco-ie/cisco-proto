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
// source: vpdn_mirroring_stats.proto

package cisco_ios_xr_tunnel_vpdn_oper_vpdn_nodes_node_vpdn_mirroring

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

type VpdnMirroringStats_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpdnMirroringStats_KEYS) Reset()         { *m = VpdnMirroringStats_KEYS{} }
func (m *VpdnMirroringStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*VpdnMirroringStats_KEYS) ProtoMessage()    {}
func (*VpdnMirroringStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b96243b8184f04b, []int{0}
}

func (m *VpdnMirroringStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpdnMirroringStats_KEYS.Unmarshal(m, b)
}
func (m *VpdnMirroringStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpdnMirroringStats_KEYS.Marshal(b, m, deterministic)
}
func (m *VpdnMirroringStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpdnMirroringStats_KEYS.Merge(m, src)
}
func (m *VpdnMirroringStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_VpdnMirroringStats_KEYS.Size(m)
}
func (m *VpdnMirroringStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VpdnMirroringStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VpdnMirroringStats_KEYS proto.InternalMessageInfo

func (m *VpdnMirroringStats_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type L2TpNsrQadStatsSend struct {
	MsgsSent             uint32   `protobuf:"varint,1,opt,name=msgs_sent,json=msgsSent,proto3" json:"msgs_sent,omitempty"`
	AcksSent             uint32   `protobuf:"varint,2,opt,name=acks_sent,json=acksSent,proto3" json:"acks_sent,omitempty"`
	NoPartner            uint32   `protobuf:"varint,3,opt,name=no_partner,json=noPartner,proto3" json:"no_partner,omitempty"`
	SendsFailed          uint32   `protobuf:"varint,4,opt,name=sends_failed,json=sendsFailed,proto3" json:"sends_failed,omitempty"`
	AcksFailed           uint32   `protobuf:"varint,5,opt,name=acks_failed,json=acksFailed,proto3" json:"acks_failed,omitempty"`
	PendingAcks          uint32   `protobuf:"varint,6,opt,name=pending_acks,json=pendingAcks,proto3" json:"pending_acks,omitempty"`
	Timeouts             uint32   `protobuf:"varint,7,opt,name=timeouts,proto3" json:"timeouts,omitempty"`
	Suspends             uint32   `protobuf:"varint,8,opt,name=suspends,proto3" json:"suspends,omitempty"`
	Resumes              uint32   `protobuf:"varint,9,opt,name=resumes,proto3" json:"resumes,omitempty"`
	SendsFragment        uint32   `protobuf:"varint,10,opt,name=sends_fragment,json=sendsFragment,proto3" json:"sends_fragment,omitempty"`
	QadLastSeqNumber     uint32   `protobuf:"varint,11,opt,name=qad_last_seq_number,json=qadLastSeqNumber,proto3" json:"qad_last_seq_number,omitempty"`
	QadFragCount         uint32   `protobuf:"varint,12,opt,name=qad_frag_count,json=qadFragCount,proto3" json:"qad_frag_count,omitempty"`
	QadAckCount          uint32   `protobuf:"varint,13,opt,name=qad_ack_count,json=qadAckCount,proto3" json:"qad_ack_count,omitempty"`
	QadUnknownAcks       uint32   `protobuf:"varint,14,opt,name=qad_unknown_acks,json=qadUnknownAcks,proto3" json:"qad_unknown_acks,omitempty"`
	QadTimeouts          uint32   `protobuf:"varint,15,opt,name=qad_timeouts,json=qadTimeouts,proto3" json:"qad_timeouts,omitempty"`
	QadRxCount           uint32   `protobuf:"varint,16,opt,name=qad_rx_count,json=qadRxCount,proto3" json:"qad_rx_count,omitempty"`
	QadRxListCount       uint32   `protobuf:"varint,17,opt,name=qad_rx_list_count,json=qadRxListCount,proto3" json:"qad_rx_list_count,omitempty"`
	QadRxListQSize       uint32   `protobuf:"varint,18,opt,name=qad_rx_list_q_size,json=qadRxListQSize,proto3" json:"qad_rx_list_q_size,omitempty"`
	QadRxFirstSeqNumber  uint32   `protobuf:"varint,19,opt,name=qad_rx_first_seq_number,json=qadRxFirstSeqNumber,proto3" json:"qad_rx_first_seq_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpNsrQadStatsSend) Reset()         { *m = L2TpNsrQadStatsSend{} }
func (m *L2TpNsrQadStatsSend) String() string { return proto.CompactTextString(m) }
func (*L2TpNsrQadStatsSend) ProtoMessage()    {}
func (*L2TpNsrQadStatsSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b96243b8184f04b, []int{1}
}

func (m *L2TpNsrQadStatsSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpNsrQadStatsSend.Unmarshal(m, b)
}
func (m *L2TpNsrQadStatsSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpNsrQadStatsSend.Marshal(b, m, deterministic)
}
func (m *L2TpNsrQadStatsSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpNsrQadStatsSend.Merge(m, src)
}
func (m *L2TpNsrQadStatsSend) XXX_Size() int {
	return xxx_messageInfo_L2TpNsrQadStatsSend.Size(m)
}
func (m *L2TpNsrQadStatsSend) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpNsrQadStatsSend.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpNsrQadStatsSend proto.InternalMessageInfo

func (m *L2TpNsrQadStatsSend) GetMsgsSent() uint32 {
	if m != nil {
		return m.MsgsSent
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetAcksSent() uint32 {
	if m != nil {
		return m.AcksSent
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetNoPartner() uint32 {
	if m != nil {
		return m.NoPartner
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetSendsFailed() uint32 {
	if m != nil {
		return m.SendsFailed
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetAcksFailed() uint32 {
	if m != nil {
		return m.AcksFailed
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetPendingAcks() uint32 {
	if m != nil {
		return m.PendingAcks
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetTimeouts() uint32 {
	if m != nil {
		return m.Timeouts
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetSuspends() uint32 {
	if m != nil {
		return m.Suspends
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetResumes() uint32 {
	if m != nil {
		return m.Resumes
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetSendsFragment() uint32 {
	if m != nil {
		return m.SendsFragment
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadLastSeqNumber() uint32 {
	if m != nil {
		return m.QadLastSeqNumber
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadFragCount() uint32 {
	if m != nil {
		return m.QadFragCount
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadAckCount() uint32 {
	if m != nil {
		return m.QadAckCount
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadUnknownAcks() uint32 {
	if m != nil {
		return m.QadUnknownAcks
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadTimeouts() uint32 {
	if m != nil {
		return m.QadTimeouts
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadRxCount() uint32 {
	if m != nil {
		return m.QadRxCount
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadRxListCount() uint32 {
	if m != nil {
		return m.QadRxListCount
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadRxListQSize() uint32 {
	if m != nil {
		return m.QadRxListQSize
	}
	return 0
}

func (m *L2TpNsrQadStatsSend) GetQadRxFirstSeqNumber() uint32 {
	if m != nil {
		return m.QadRxFirstSeqNumber
	}
	return 0
}

type L2TpNsrQadStatsRecv struct {
	MsgsRecvd            uint32   `protobuf:"varint,1,opt,name=msgs_recvd,json=msgsRecvd,proto3" json:"msgs_recvd,omitempty"`
	AcksRecvd            uint32   `protobuf:"varint,2,opt,name=acks_recvd,json=acksRecvd,proto3" json:"acks_recvd,omitempty"`
	RecvdAcksFailed      uint32   `protobuf:"varint,3,opt,name=recvd_acks_failed,json=recvdAcksFailed,proto3" json:"recvd_acks_failed,omitempty"`
	InitDrops            uint32   `protobuf:"varint,4,opt,name=init_drops,json=initDrops,proto3" json:"init_drops,omitempty"`
	MsgDrops             uint32   `protobuf:"varint,5,opt,name=msg_drops,json=msgDrops,proto3" json:"msg_drops,omitempty"`
	OooDrops             uint32   `protobuf:"varint,6,opt,name=ooo_drops,json=oooDrops,proto3" json:"ooo_drops,omitempty"`
	StaleMsgs            uint32   `protobuf:"varint,7,opt,name=stale_msgs,json=staleMsgs,proto3" json:"stale_msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2TpNsrQadStatsRecv) Reset()         { *m = L2TpNsrQadStatsRecv{} }
func (m *L2TpNsrQadStatsRecv) String() string { return proto.CompactTextString(m) }
func (*L2TpNsrQadStatsRecv) ProtoMessage()    {}
func (*L2TpNsrQadStatsRecv) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b96243b8184f04b, []int{2}
}

func (m *L2TpNsrQadStatsRecv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2TpNsrQadStatsRecv.Unmarshal(m, b)
}
func (m *L2TpNsrQadStatsRecv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2TpNsrQadStatsRecv.Marshal(b, m, deterministic)
}
func (m *L2TpNsrQadStatsRecv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2TpNsrQadStatsRecv.Merge(m, src)
}
func (m *L2TpNsrQadStatsRecv) XXX_Size() int {
	return xxx_messageInfo_L2TpNsrQadStatsRecv.Size(m)
}
func (m *L2TpNsrQadStatsRecv) XXX_DiscardUnknown() {
	xxx_messageInfo_L2TpNsrQadStatsRecv.DiscardUnknown(m)
}

var xxx_messageInfo_L2TpNsrQadStatsRecv proto.InternalMessageInfo

func (m *L2TpNsrQadStatsRecv) GetMsgsRecvd() uint32 {
	if m != nil {
		return m.MsgsRecvd
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetAcksRecvd() uint32 {
	if m != nil {
		return m.AcksRecvd
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetRecvdAcksFailed() uint32 {
	if m != nil {
		return m.RecvdAcksFailed
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetInitDrops() uint32 {
	if m != nil {
		return m.InitDrops
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetMsgDrops() uint32 {
	if m != nil {
		return m.MsgDrops
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetOooDrops() uint32 {
	if m != nil {
		return m.OooDrops
	}
	return 0
}

func (m *L2TpNsrQadStatsRecv) GetStaleMsgs() uint32 {
	if m != nil {
		return m.StaleMsgs
	}
	return 0
}

type VpdnMirroringStats struct {
	SyncNotConnCnt        uint32               `protobuf:"varint,50,opt,name=sync_not_conn_cnt,json=syncNotConnCnt,proto3" json:"sync_not_conn_cnt,omitempty"`
	SsoErrCnt             uint32               `protobuf:"varint,51,opt,name=sso_err_cnt,json=ssoErrCnt,proto3" json:"sso_err_cnt,omitempty"`
	SsoBatchErrCnt        uint32               `protobuf:"varint,52,opt,name=sso_batch_err_cnt,json=ssoBatchErrCnt,proto3" json:"sso_batch_err_cnt,omitempty"`
	AllocErrCnt           uint32               `protobuf:"varint,53,opt,name=alloc_err_cnt,json=allocErrCnt,proto3" json:"alloc_err_cnt,omitempty"`
	AllocCnt              uint32               `protobuf:"varint,54,opt,name=alloc_cnt,json=allocCnt,proto3" json:"alloc_cnt,omitempty"`
	QadSendStats          *L2TpNsrQadStatsSend `protobuf:"bytes,55,opt,name=qad_send_stats,json=qadSendStats,proto3" json:"qad_send_stats,omitempty"`
	QadRecvStats          *L2TpNsrQadStatsRecv `protobuf:"bytes,56,opt,name=qad_recv_stats,json=qadRecvStats,proto3" json:"qad_recv_stats,omitempty"`
	QadSendStatsLastClear *L2TpNsrQadStatsSend `protobuf:"bytes,57,opt,name=qad_send_stats_last_clear,json=qadSendStatsLastClear,proto3" json:"qad_send_stats_last_clear,omitempty"`
	QadRecvStatsLastClear *L2TpNsrQadStatsRecv `protobuf:"bytes,58,opt,name=qad_recv_stats_last_clear,json=qadRecvStatsLastClear,proto3" json:"qad_recv_stats_last_clear,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}             `json:"-"`
	XXX_unrecognized      []byte               `json:"-"`
	XXX_sizecache         int32                `json:"-"`
}

func (m *VpdnMirroringStats) Reset()         { *m = VpdnMirroringStats{} }
func (m *VpdnMirroringStats) String() string { return proto.CompactTextString(m) }
func (*VpdnMirroringStats) ProtoMessage()    {}
func (*VpdnMirroringStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b96243b8184f04b, []int{3}
}

func (m *VpdnMirroringStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpdnMirroringStats.Unmarshal(m, b)
}
func (m *VpdnMirroringStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpdnMirroringStats.Marshal(b, m, deterministic)
}
func (m *VpdnMirroringStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpdnMirroringStats.Merge(m, src)
}
func (m *VpdnMirroringStats) XXX_Size() int {
	return xxx_messageInfo_VpdnMirroringStats.Size(m)
}
func (m *VpdnMirroringStats) XXX_DiscardUnknown() {
	xxx_messageInfo_VpdnMirroringStats.DiscardUnknown(m)
}

var xxx_messageInfo_VpdnMirroringStats proto.InternalMessageInfo

func (m *VpdnMirroringStats) GetSyncNotConnCnt() uint32 {
	if m != nil {
		return m.SyncNotConnCnt
	}
	return 0
}

func (m *VpdnMirroringStats) GetSsoErrCnt() uint32 {
	if m != nil {
		return m.SsoErrCnt
	}
	return 0
}

func (m *VpdnMirroringStats) GetSsoBatchErrCnt() uint32 {
	if m != nil {
		return m.SsoBatchErrCnt
	}
	return 0
}

func (m *VpdnMirroringStats) GetAllocErrCnt() uint32 {
	if m != nil {
		return m.AllocErrCnt
	}
	return 0
}

func (m *VpdnMirroringStats) GetAllocCnt() uint32 {
	if m != nil {
		return m.AllocCnt
	}
	return 0
}

func (m *VpdnMirroringStats) GetQadSendStats() *L2TpNsrQadStatsSend {
	if m != nil {
		return m.QadSendStats
	}
	return nil
}

func (m *VpdnMirroringStats) GetQadRecvStats() *L2TpNsrQadStatsRecv {
	if m != nil {
		return m.QadRecvStats
	}
	return nil
}

func (m *VpdnMirroringStats) GetQadSendStatsLastClear() *L2TpNsrQadStatsSend {
	if m != nil {
		return m.QadSendStatsLastClear
	}
	return nil
}

func (m *VpdnMirroringStats) GetQadRecvStatsLastClear() *L2TpNsrQadStatsRecv {
	if m != nil {
		return m.QadRecvStatsLastClear
	}
	return nil
}

func init() {
	proto.RegisterType((*VpdnMirroringStats_KEYS)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_mirroring.vpdn_mirroring_stats_KEYS")
	proto.RegisterType((*L2TpNsrQadStatsSend)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_mirroring.l2tp_nsr_qad_stats_send")
	proto.RegisterType((*L2TpNsrQadStatsRecv)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_mirroring.l2tp_nsr_qad_stats_recv")
	proto.RegisterType((*VpdnMirroringStats)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_mirroring.vpdn_mirroring_stats")
}

func init() { proto.RegisterFile("vpdn_mirroring_stats.proto", fileDescriptor_1b96243b8184f04b) }

var fileDescriptor_1b96243b8184f04b = []byte{
	// 778 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x4f, 0x23, 0x37,
	0x14, 0x55, 0xda, 0x06, 0x92, 0x9b, 0x84, 0x8f, 0xa1, 0x88, 0x81, 0x8a, 0x16, 0xa2, 0x56, 0xa2,
	0x48, 0xcd, 0x03, 0xd0, 0x96, 0x56, 0x7d, 0xa1, 0x29, 0xbc, 0x94, 0xa2, 0x36, 0x29, 0x0f, 0x7d,
	0xb2, 0xcc, 0x8c, 0x49, 0x47, 0x99, 0xb1, 0x33, 0xbe, 0x0e, 0x9b, 0x65, 0x7f, 0xc4, 0xfe, 0xac,
	0xfd, 0x31, 0xfb, 0x13, 0xf6, 0x65, 0x75, 0xaf, 0x9d, 0x21, 0xac, 0xd8, 0xa7, 0x15, 0x2f, 0x51,
	0x7c, 0xce, 0xf1, 0xbd, 0xf6, 0xf1, 0xb1, 0x07, 0x76, 0xee, 0x26, 0xa9, 0x16, 0x45, 0x66, 0xad,
	0xb1, 0x99, 0x1e, 0x09, 0x74, 0xd2, 0x61, 0x6f, 0x62, 0x8d, 0x33, 0xd1, 0x6f, 0x49, 0x86, 0x89,
	0x11, 0x99, 0x41, 0x31, 0xb3, 0xc2, 0x4d, 0xb5, 0x56, 0xb9, 0x60, 0xbd, 0x99, 0x28, 0xdb, 0xa3,
	0x7f, 0x3d, 0x6d, 0x52, 0x85, 0xfc, 0xdb, 0x7b, 0x5c, 0xa9, 0x7b, 0x0a, 0xdb, 0x4f, 0xd5, 0x16,
	0x7f, 0x9e, 0xff, 0x37, 0x8c, 0xbe, 0x82, 0x26, 0xcd, 0x11, 0x5a, 0x16, 0x2a, 0xae, 0xed, 0xd5,
	0x0e, 0x9a, 0x83, 0x06, 0x01, 0x57, 0xb2, 0x50, 0xdd, 0x37, 0x75, 0xd8, 0xca, 0x8f, 0xdc, 0x44,
	0x68, 0xb4, 0xa2, 0x94, 0x69, 0x98, 0x88, 0x4a, 0xa7, 0x34, 0xb1, 0xc0, 0x11, 0x0f, 0x1c, 0x4f,
	0xec, 0x0c, 0x1a, 0x04, 0x0c, 0x95, 0x76, 0x44, 0xca, 0x64, 0x1c, 0xc8, 0xcf, 0x3c, 0x49, 0x00,
	0x93, 0xbb, 0x00, 0xda, 0x88, 0x89, 0xb4, 0x4e, 0x2b, 0x1b, 0x7f, 0xce, 0x6c, 0x53, 0x9b, 0xbf,
	0x3d, 0x10, 0xed, 0x43, 0x9b, 0x1a, 0xa0, 0xb8, 0x95, 0x59, 0xae, 0xd2, 0xf8, 0x0b, 0x16, 0xb4,
	0x18, 0xbb, 0x60, 0x28, 0xfa, 0x06, 0x5a, 0x5c, 0x3e, 0x28, 0xea, 0xac, 0x00, 0x82, 0x82, 0x60,
	0x1f, 0xda, 0x13, 0xa5, 0x53, 0xda, 0x2b, 0xa1, 0xf1, 0x92, 0xaf, 0x11, 0xb0, 0xb3, 0x64, 0x8c,
	0xd1, 0x0e, 0x34, 0x5c, 0x56, 0x28, 0x33, 0x75, 0x18, 0x2f, 0xfb, 0x15, 0xce, 0xc7, 0xc4, 0xe1,
	0x14, 0x49, 0x8d, 0x71, 0xc3, 0x73, 0xf3, 0x71, 0x14, 0xc3, 0xb2, 0x55, 0x38, 0x2d, 0x14, 0xc6,
	0x4d, 0xa6, 0xe6, 0xc3, 0xe8, 0x3b, 0x58, 0x09, 0x0b, 0xb7, 0x72, 0x54, 0xd0, 0xce, 0x81, 0x05,
	0x1d, 0xbf, 0xf4, 0x00, 0x46, 0x3f, 0xc0, 0x06, 0x59, 0x99, 0x4b, 0x74, 0x02, 0x55, 0x29, 0xf4,
	0xb4, 0xb8, 0x51, 0x36, 0x6e, 0xb1, 0x76, 0xad, 0x94, 0xe9, 0xa5, 0x44, 0x37, 0x54, 0xe5, 0x15,
	0xe3, 0xd1, 0xb7, 0xb0, 0x42, 0x72, 0xaa, 0x29, 0x12, 0x33, 0xd5, 0x2e, 0x6e, 0xb3, 0xb2, 0x5d,
	0xca, 0x94, 0x6a, 0xf6, 0x09, 0x8b, 0xba, 0xd0, 0x21, 0x95, 0x4c, 0xc6, 0x41, 0xd4, 0xf1, 0x3b,
	0x2e, 0x65, 0x7a, 0x96, 0x8c, 0xbd, 0xe6, 0x00, 0xa8, 0xba, 0x98, 0xea, 0xb1, 0x36, 0x2f, 0xb4,
	0x37, 0x66, 0x85, 0x65, 0xd4, 0xe1, 0xda, 0xc3, 0xec, 0xcd, 0x3e, 0x50, 0x75, 0x51, 0xf9, 0xb3,
	0x5a, 0x15, 0xfb, 0x77, 0x6e, 0xd1, 0x9e, 0x97, 0xd8, 0x59, 0xe8, 0xb7, 0xe6, 0xcf, 0xa0, 0x94,
	0xe9, 0x60, 0xe6, 0xdb, 0x7d, 0x0f, 0xeb, 0x41, 0x91, 0x67, 0xe8, 0x82, 0x6c, 0xbd, 0xea, 0x37,
	0x98, 0x5d, 0x66, 0xe8, 0xbc, 0xf4, 0x10, 0xa2, 0x45, 0x69, 0x29, 0x30, 0xbb, 0x57, 0x71, 0xf4,
	0x81, 0xf6, 0x9f, 0x61, 0x76, 0xaf, 0xa2, 0x13, 0xd8, 0x0a, 0xda, 0xdb, 0xcc, 0x3e, 0xb6, 0x70,
	0x83, 0x27, 0x6c, 0xf0, 0x84, 0x0b, 0x22, 0x2b, 0x17, 0xbb, 0xef, 0x6a, 0x4f, 0x26, 0xd9, 0xaa,
	0xe4, 0x8e, 0xf2, 0xc8, 0x49, 0xa6, 0x41, 0x1a, 0xa2, 0xcc, 0xd9, 0x1e, 0x10, 0x40, 0x34, 0x87,
	0xcd, 0xd3, 0x3e, 0xcc, 0x9c, 0x6e, 0x4f, 0x1f, 0xc2, 0x3a, 0x33, 0x62, 0x31, 0x91, 0x3e, 0xd4,
	0xab, 0x4c, 0x9c, 0x3d, 0xc4, 0x72, 0x17, 0x20, 0xd3, 0x99, 0x13, 0xa9, 0x35, 0x13, 0x0c, 0xc1,
	0x6e, 0x12, 0xf2, 0x07, 0x01, 0xe1, 0x4a, 0x05, 0xb6, 0x5e, 0x5d, 0xa9, 0x8a, 0x34, 0xc6, 0x04,
	0xd2, 0xe7, 0xb9, 0x61, 0x8c, 0xf1, 0xe4, 0x2e, 0x00, 0x3a, 0x99, 0x2b, 0x41, 0xcb, 0x0e, 0x71,
	0x6e, 0x32, 0xf2, 0x17, 0x8e, 0xb0, 0xfb, 0xb6, 0x0e, 0x5f, 0x3e, 0xf5, 0x04, 0xd0, 0x19, 0xe1,
	0x4b, 0x9d, 0x08, 0x6d, 0xe8, 0x80, 0xb4, 0x16, 0x89, 0x76, 0xf1, 0x91, 0xf7, 0x9d, 0x88, 0x2b,
	0xe3, 0xfa, 0x46, 0xeb, 0xbe, 0x76, 0xd1, 0xd7, 0xd0, 0x42, 0x34, 0x42, 0x59, 0xcb, 0xa2, 0xe3,
	0xd0, 0x03, 0xcd, 0xb9, 0xb5, 0x7d, 0x7f, 0xdc, 0xc4, 0xdf, 0x48, 0x97, 0xfc, 0x5f, 0xa9, 0x4e,
	0x42, 0x29, 0x34, 0xbf, 0x13, 0x1e, 0xa4, 0x5d, 0xe8, 0xc8, 0x3c, 0x37, 0x49, 0x25, 0xfb, 0xd1,
	0xe7, 0x8b, 0xc1, 0xa0, 0xa1, 0x17, 0x84, 0x35, 0xc4, 0xff, 0x14, 0x5e, 0x10, 0x02, 0x88, 0x7c,
	0xe5, 0xef, 0x04, 0xdd, 0x2b, 0xbf, 0x91, 0xf8, 0xe7, 0xbd, 0xda, 0x41, 0xeb, 0xe8, 0xba, 0xf7,
	0x29, 0x0f, 0x65, 0xef, 0x23, 0x4f, 0x1d, 0x5f, 0xb5, 0xa1, 0xd2, 0xe9, 0x90, 0x3d, 0x0b, 0xcd,
	0xe9, 0x6c, 0x43, 0xf3, 0xd3, 0x67, 0x6a, 0x4e, 0x2d, 0xb8, 0x39, 0x25, 0xcd, 0x37, 0x7f, 0x5d,
	0x83, 0xed, 0xc7, 0x5b, 0xf7, 0x0f, 0x49, 0x92, 0x2b, 0x69, 0xe3, 0x5f, 0x9e, 0xd3, 0x85, 0xcd,
	0x45, 0x17, 0xe8, 0x8d, 0xea, 0x53, 0xcf, 0x6a, 0x45, 0x0f, 0x7e, 0x2c, 0xae, 0xe8, 0xd7, 0xe7,
	0xb4, 0x66, 0x73, 0xd1, 0x9a, 0x6a, 0x45, 0x37, 0x4b, 0xfc, 0xd1, 0x3c, 0x7e, 0x1f, 0x00, 0x00,
	0xff, 0xff, 0xc6, 0x94, 0x33, 0x7c, 0x52, 0x07, 0x00, 0x00,
}
