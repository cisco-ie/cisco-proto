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
// source: ospfv3_edm_nsr_stats.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_statistics_nsr_stats

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

type Ospfv3EdmNsrStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmNsrStats_KEYS) Reset()         { *m = Ospfv3EdmNsrStats_KEYS{} }
func (m *Ospfv3EdmNsrStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNsrStats_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmNsrStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fee77c6427e3651, []int{0}
}

func (m *Ospfv3EdmNsrStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNsrStats_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmNsrStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNsrStats_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNsrStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNsrStats_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmNsrStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNsrStats_KEYS.Size(m)
}
func (m *Ospfv3EdmNsrStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNsrStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNsrStats_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmNsrStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShNsrStatsPri struct {
	NsrSchedPri          string   `protobuf:"bytes,1,opt,name=nsr_sched_pri,json=nsrSchedPri,proto3" json:"nsr_sched_pri,omitempty"`
	NsrSchedQuant        uint32   `protobuf:"varint,2,opt,name=nsr_sched_quant,json=nsrSchedQuant,proto3" json:"nsr_sched_quant,omitempty"`
	NsrSchedRemainQuant  uint32   `protobuf:"varint,3,opt,name=nsr_sched_remain_quant,json=nsrSchedRemainQuant,proto3" json:"nsr_sched_remain_quant,omitempty"`
	NsrSchedEvsInQ       uint32   `protobuf:"varint,4,opt,name=nsr_sched_evs_in_q,json=nsrSchedEvsInQ,proto3" json:"nsr_sched_evs_in_q,omitempty"`
	NsrSchedMaxEvs       uint32   `protobuf:"varint,5,opt,name=nsr_sched_max_evs,json=nsrSchedMaxEvs,proto3" json:"nsr_sched_max_evs,omitempty"`
	NsrSchedPeakQLen     uint32   `protobuf:"varint,6,opt,name=nsr_sched_peak_q_len,json=nsrSchedPeakQLen,proto3" json:"nsr_sched_peak_q_len,omitempty"`
	NsrSchedEvsQd        uint64   `protobuf:"varint,7,opt,name=nsr_sched_evs_qd,json=nsrSchedEvsQd,proto3" json:"nsr_sched_evs_qd,omitempty"`
	NsrSchedEnqFails     uint64   `protobuf:"varint,8,opt,name=nsr_sched_enq_fails,json=nsrSchedEnqFails,proto3" json:"nsr_sched_enq_fails,omitempty"`
	NsrSchedEvsDeqd      uint64   `protobuf:"varint,9,opt,name=nsr_sched_evs_deqd,json=nsrSchedEvsDeqd,proto3" json:"nsr_sched_evs_deqd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNsrStatsPri) Reset()         { *m = OspfShNsrStatsPri{} }
func (m *OspfShNsrStatsPri) String() string { return proto.CompactTextString(m) }
func (*OspfShNsrStatsPri) ProtoMessage()    {}
func (*OspfShNsrStatsPri) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fee77c6427e3651, []int{1}
}

func (m *OspfShNsrStatsPri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNsrStatsPri.Unmarshal(m, b)
}
func (m *OspfShNsrStatsPri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNsrStatsPri.Marshal(b, m, deterministic)
}
func (m *OspfShNsrStatsPri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNsrStatsPri.Merge(m, src)
}
func (m *OspfShNsrStatsPri) XXX_Size() int {
	return xxx_messageInfo_OspfShNsrStatsPri.Size(m)
}
func (m *OspfShNsrStatsPri) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNsrStatsPri.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNsrStatsPri proto.InternalMessageInfo

func (m *OspfShNsrStatsPri) GetNsrSchedPri() string {
	if m != nil {
		return m.NsrSchedPri
	}
	return ""
}

func (m *OspfShNsrStatsPri) GetNsrSchedQuant() uint32 {
	if m != nil {
		return m.NsrSchedQuant
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedRemainQuant() uint32 {
	if m != nil {
		return m.NsrSchedRemainQuant
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedEvsInQ() uint32 {
	if m != nil {
		return m.NsrSchedEvsInQ
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedMaxEvs() uint32 {
	if m != nil {
		return m.NsrSchedMaxEvs
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedPeakQLen() uint32 {
	if m != nil {
		return m.NsrSchedPeakQLen
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedEvsQd() uint64 {
	if m != nil {
		return m.NsrSchedEvsQd
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedEnqFails() uint64 {
	if m != nil {
		return m.NsrSchedEnqFails
	}
	return 0
}

func (m *OspfShNsrStatsPri) GetNsrSchedEvsDeqd() uint64 {
	if m != nil {
		return m.NsrSchedEvsDeqd
	}
	return 0
}

type OspfShNsrSchedStats struct {
	NsrPulseQuant        int32                `protobuf:"zigzag32,1,opt,name=nsr_pulse_quant,json=nsrPulseQuant,proto3" json:"nsr_pulse_quant,omitempty"`
	NsrEventsInQ         uint32               `protobuf:"varint,2,opt,name=nsr_events_in_q,json=nsrEventsInQ,proto3" json:"nsr_events_in_q,omitempty"`
	NsrEventsTx          uint64               `protobuf:"varint,3,opt,name=nsr_events_tx,json=nsrEventsTx,proto3" json:"nsr_events_tx,omitempty"`
	NsrEventsRx          uint64               `protobuf:"varint,4,opt,name=nsr_events_rx,json=nsrEventsRx,proto3" json:"nsr_events_rx,omitempty"`
	NsrBadPulsesRx       uint64               `protobuf:"varint,5,opt,name=nsr_bad_pulses_rx,json=nsrBadPulsesRx,proto3" json:"nsr_bad_pulses_rx,omitempty"`
	NsrGoodPulsesRx      uint64               `protobuf:"varint,6,opt,name=nsr_good_pulses_rx,json=nsrGoodPulsesRx,proto3" json:"nsr_good_pulses_rx,omitempty"`
	NsrPulsesTx          uint64               `protobuf:"varint,7,opt,name=nsr_pulses_tx,json=nsrPulsesTx,proto3" json:"nsr_pulses_tx,omitempty"`
	NsrPulseTxFails      uint64               `protobuf:"varint,8,opt,name=nsr_pulse_tx_fails,json=nsrPulseTxFails,proto3" json:"nsr_pulse_tx_fails,omitempty"`
	NsrPri               []*OspfShNsrStatsPri `protobuf:"bytes,9,rep,name=nsr_pri,json=nsrPri,proto3" json:"nsr_pri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OspfShNsrSchedStats) Reset()         { *m = OspfShNsrSchedStats{} }
func (m *OspfShNsrSchedStats) String() string { return proto.CompactTextString(m) }
func (*OspfShNsrSchedStats) ProtoMessage()    {}
func (*OspfShNsrSchedStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fee77c6427e3651, []int{2}
}

func (m *OspfShNsrSchedStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNsrSchedStats.Unmarshal(m, b)
}
func (m *OspfShNsrSchedStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNsrSchedStats.Marshal(b, m, deterministic)
}
func (m *OspfShNsrSchedStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNsrSchedStats.Merge(m, src)
}
func (m *OspfShNsrSchedStats) XXX_Size() int {
	return xxx_messageInfo_OspfShNsrSchedStats.Size(m)
}
func (m *OspfShNsrSchedStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNsrSchedStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNsrSchedStats proto.InternalMessageInfo

func (m *OspfShNsrSchedStats) GetNsrPulseQuant() int32 {
	if m != nil {
		return m.NsrPulseQuant
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrEventsInQ() uint32 {
	if m != nil {
		return m.NsrEventsInQ
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrEventsTx() uint64 {
	if m != nil {
		return m.NsrEventsTx
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrEventsRx() uint64 {
	if m != nil {
		return m.NsrEventsRx
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrBadPulsesRx() uint64 {
	if m != nil {
		return m.NsrBadPulsesRx
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrGoodPulsesRx() uint64 {
	if m != nil {
		return m.NsrGoodPulsesRx
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrPulsesTx() uint64 {
	if m != nil {
		return m.NsrPulsesTx
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrPulseTxFails() uint64 {
	if m != nil {
		return m.NsrPulseTxFails
	}
	return 0
}

func (m *OspfShNsrSchedStats) GetNsrPri() []*OspfShNsrStatsPri {
	if m != nil {
		return m.NsrPri
	}
	return nil
}

type Ospfv3EdmNsrStats struct {
	NsrRevision             int32                `protobuf:"zigzag32,50,opt,name=nsr_revision,json=nsrRevision,proto3" json:"nsr_revision,omitempty"`
	NsrFsmState             int32                `protobuf:"zigzag32,51,opt,name=nsr_fsm_state,json=nsrFsmState,proto3" json:"nsr_fsm_state,omitempty"`
	NsrVersion              uint32               `protobuf:"varint,52,opt,name=nsr_version,json=nsrVersion,proto3" json:"nsr_version,omitempty"`
	NsrNodeId               uint32               `protobuf:"varint,53,opt,name=nsr_node_id,json=nsrNodeId,proto3" json:"nsr_node_id,omitempty"`
	NsrPeerVersion          uint32               `protobuf:"varint,54,opt,name=nsr_peer_version,json=nsrPeerVersion,proto3" json:"nsr_peer_version,omitempty"`
	NsrPeerNodeId           uint32               `protobuf:"varint,55,opt,name=nsr_peer_node_id,json=nsrPeerNodeId,proto3" json:"nsr_peer_node_id,omitempty"`
	NsrMtu                  uint32               `protobuf:"varint,56,opt,name=nsr_mtu,json=nsrMtu,proto3" json:"nsr_mtu,omitempty"`
	NsrNbrQadQid            uint32               `protobuf:"varint,57,opt,name=nsr_nbr_qad_qid,json=nsrNbrQadQid,proto3" json:"nsr_nbr_qad_qid,omitempty"`
	NsrLsaQadQid            uint32               `protobuf:"varint,58,opt,name=nsr_lsa_qad_qid,json=nsrLsaQadQid,proto3" json:"nsr_lsa_qad_qid,omitempty"`
	NsrNbrQadMdataCount     uint32               `protobuf:"varint,59,opt,name=nsr_nbr_qad_mdata_count,json=nsrNbrQadMdataCount,proto3" json:"nsr_nbr_qad_mdata_count,omitempty"`
	NsrLsaQadMdataCount     uint32               `protobuf:"varint,60,opt,name=nsr_lsa_qad_mdata_count,json=nsrLsaQadMdataCount,proto3" json:"nsr_lsa_qad_mdata_count,omitempty"`
	NsrNbrInitSyncPendCount int32                `protobuf:"zigzag32,61,opt,name=nsr_nbr_init_sync_pend_count,json=nsrNbrInitSyncPendCount,proto3" json:"nsr_nbr_init_sync_pend_count,omitempty"`
	NsrLsaInitSyncPendCount int32                `protobuf:"zigzag32,62,opt,name=nsr_lsa_init_sync_pend_count,json=nsrLsaInitSyncPendCount,proto3" json:"nsr_lsa_init_sync_pend_count,omitempty"`
	NsrNbrSeqNo             uint32               `protobuf:"varint,63,opt,name=nsr_nbr_seq_no,json=nsrNbrSeqNo,proto3" json:"nsr_nbr_seq_no,omitempty"`
	NsrIntfSeqNo            uint32               `protobuf:"varint,64,opt,name=nsr_intf_seq_no,json=nsrIntfSeqNo,proto3" json:"nsr_intf_seq_no,omitempty"`
	NsrTmrQuant             int32                `protobuf:"zigzag32,65,opt,name=nsr_tmr_quant,json=nsrTmrQuant,proto3" json:"nsr_tmr_quant,omitempty"`
	NsrConnToActiveAttempts uint64               `protobuf:"varint,66,opt,name=nsr_conn_to_active_attempts,json=nsrConnToActiveAttempts,proto3" json:"nsr_conn_to_active_attempts,omitempty"`
	NsrConnToActiveFails    uint64               `protobuf:"varint,67,opt,name=nsr_conn_to_active_fails,json=nsrConnToActiveFails,proto3" json:"nsr_conn_to_active_fails,omitempty"`
	NsrConnToActiveOpens    uint64               `protobuf:"varint,68,opt,name=nsr_conn_to_active_opens,json=nsrConnToActiveOpens,proto3" json:"nsr_conn_to_active_opens,omitempty"`
	NsrConnToActiveCloses   uint64               `protobuf:"varint,69,opt,name=nsr_conn_to_active_closes,json=nsrConnToActiveCloses,proto3" json:"nsr_conn_to_active_closes,omitempty"`
	NsrConnToActiveErrors   uint64               `protobuf:"varint,70,opt,name=nsr_conn_to_active_errors,json=nsrConnToActiveErrors,proto3" json:"nsr_conn_to_active_errors,omitempty"`
	NsrThdStats             *OspfShNsrSchedStats `protobuf:"bytes,71,opt,name=nsr_thd_stats,json=nsrThdStats,proto3" json:"nsr_thd_stats,omitempty"`
	NsrRtrThdSched          *OspfShNsrSchedStats `protobuf:"bytes,72,opt,name=nsr_rtr_thd_sched,json=nsrRtrThdSched,proto3" json:"nsr_rtr_thd_sched,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *Ospfv3EdmNsrStats) Reset()         { *m = Ospfv3EdmNsrStats{} }
func (m *Ospfv3EdmNsrStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmNsrStats) ProtoMessage()    {}
func (*Ospfv3EdmNsrStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fee77c6427e3651, []int{3}
}

func (m *Ospfv3EdmNsrStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmNsrStats.Unmarshal(m, b)
}
func (m *Ospfv3EdmNsrStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmNsrStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmNsrStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmNsrStats.Merge(m, src)
}
func (m *Ospfv3EdmNsrStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmNsrStats.Size(m)
}
func (m *Ospfv3EdmNsrStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmNsrStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmNsrStats proto.InternalMessageInfo

func (m *Ospfv3EdmNsrStats) GetNsrRevision() int32 {
	if m != nil {
		return m.NsrRevision
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrFsmState() int32 {
	if m != nil {
		return m.NsrFsmState
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrVersion() uint32 {
	if m != nil {
		return m.NsrVersion
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrNodeId() uint32 {
	if m != nil {
		return m.NsrNodeId
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrPeerVersion() uint32 {
	if m != nil {
		return m.NsrPeerVersion
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrPeerNodeId() uint32 {
	if m != nil {
		return m.NsrPeerNodeId
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrMtu() uint32 {
	if m != nil {
		return m.NsrMtu
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrNbrQadQid() uint32 {
	if m != nil {
		return m.NsrNbrQadQid
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrLsaQadQid() uint32 {
	if m != nil {
		return m.NsrLsaQadQid
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrNbrQadMdataCount() uint32 {
	if m != nil {
		return m.NsrNbrQadMdataCount
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrLsaQadMdataCount() uint32 {
	if m != nil {
		return m.NsrLsaQadMdataCount
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrNbrInitSyncPendCount() int32 {
	if m != nil {
		return m.NsrNbrInitSyncPendCount
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrLsaInitSyncPendCount() int32 {
	if m != nil {
		return m.NsrLsaInitSyncPendCount
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrNbrSeqNo() uint32 {
	if m != nil {
		return m.NsrNbrSeqNo
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrIntfSeqNo() uint32 {
	if m != nil {
		return m.NsrIntfSeqNo
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrTmrQuant() int32 {
	if m != nil {
		return m.NsrTmrQuant
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrConnToActiveAttempts() uint64 {
	if m != nil {
		return m.NsrConnToActiveAttempts
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrConnToActiveFails() uint64 {
	if m != nil {
		return m.NsrConnToActiveFails
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrConnToActiveOpens() uint64 {
	if m != nil {
		return m.NsrConnToActiveOpens
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrConnToActiveCloses() uint64 {
	if m != nil {
		return m.NsrConnToActiveCloses
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrConnToActiveErrors() uint64 {
	if m != nil {
		return m.NsrConnToActiveErrors
	}
	return 0
}

func (m *Ospfv3EdmNsrStats) GetNsrThdStats() *OspfShNsrSchedStats {
	if m != nil {
		return m.NsrThdStats
	}
	return nil
}

func (m *Ospfv3EdmNsrStats) GetNsrRtrThdSched() *OspfShNsrSchedStats {
	if m != nil {
		return m.NsrRtrThdSched
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmNsrStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.nsr_stats.ospfv3_edm_nsr_stats_KEYS")
	proto.RegisterType((*OspfShNsrStatsPri)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.nsr_stats.ospf_sh_nsr_stats_pri")
	proto.RegisterType((*OspfShNsrSchedStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.nsr_stats.ospf_sh_nsr_sched_stats")
	proto.RegisterType((*Ospfv3EdmNsrStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.nsr_stats.ospfv3_edm_nsr_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_nsr_stats.proto", fileDescriptor_8fee77c6427e3651) }

var fileDescriptor_8fee77c6427e3651 = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xdf, 0x72, 0x1b, 0x35,
	0x14, 0xc6, 0xc7, 0xa4, 0x4d, 0x89, 0xdc, 0x34, 0xed, 0x36, 0x25, 0x2a, 0x30, 0x60, 0xcc, 0x40,
	0x0d, 0x0c, 0xbe, 0x48, 0x4a, 0x28, 0xd0, 0x16, 0xd2, 0xd4, 0x29, 0x99, 0x26, 0x21, 0xb6, 0x33,
	0xcc, 0x70, 0xa5, 0x91, 0x57, 0xc7, 0x44, 0x53, 0xaf, 0xb4, 0x2b, 0xc9, 0x3b, 0xdb, 0xe1, 0x05,
	0xb8, 0xe3, 0x25, 0x78, 0x26, 0xde, 0x83, 0x37, 0x60, 0xf4, 0x6f, 0xbd, 0x0e, 0xe6, 0x0e, 0xb8,
	0xf3, 0x9e, 0xf3, 0xfb, 0x8e, 0x8e, 0xce, 0x27, 0x29, 0x41, 0x6f, 0x4b, 0x9d, 0x4f, 0xcb, 0x3d,
	0x02, 0x2c, 0x23, 0x42, 0x2b, 0xa2, 0x0d, 0x35, 0xba, 0x9f, 0x2b, 0x69, 0x64, 0xf2, 0x32, 0xe5,
	0x3a, 0x95, 0x84, 0x4b, 0x4d, 0x2a, 0x45, 0x78, 0x5e, 0xee, 0x93, 0x40, 0xcb, 0x1c, 0x54, 0xdf,
	0xff, 0xb6, 0x6c, 0x0a, 0x5a, 0x83, 0x8e, 0xbf, 0xfa, 0xb6, 0x06, 0xd7, 0x86, 0xa7, 0xba, 0x5f,
	0x97, 0xec, 0x3e, 0x45, 0xf7, 0x57, 0x2d, 0x45, 0x5e, 0x0e, 0x7e, 0x1a, 0x27, 0x1f, 0xa0, 0x9b,
	0x41, 0x4c, 0x04, 0xcd, 0x00, 0xb7, 0x3a, 0xad, 0xde, 0xc6, 0xa8, 0x1d, 0x62, 0x67, 0x34, 0x83,
	0xee, 0xef, 0x6b, 0xe8, 0x9e, 0x2d, 0x40, 0xf4, 0x65, 0x43, 0x9d, 0x2b, 0x9e, 0x74, 0xd1, 0xa6,
	0x0b, 0xa4, 0x97, 0xc0, 0x6c, 0x20, 0xaa, 0x85, 0x56, 0x63, 0x1b, 0x3b, 0x57, 0x3c, 0xf9, 0x18,
	0x6d, 0x2d, 0x98, 0x62, 0x4e, 0x85, 0xc1, 0x6f, 0x74, 0x5a, 0xbd, 0xcd, 0xd1, 0x66, 0xa4, 0x86,
	0x36, 0x98, 0xec, 0xa1, 0xb7, 0x16, 0x9c, 0x82, 0x8c, 0x72, 0x11, 0xf0, 0x35, 0x87, 0xdf, 0x8d,
	0xf8, 0xc8, 0xe5, 0xbc, 0xe8, 0x53, 0x94, 0x2c, 0x44, 0x50, 0x6a, 0x62, 0x45, 0xf8, 0x9a, 0x13,
	0xdc, 0x8a, 0x82, 0x41, 0xa9, 0x8f, 0xc5, 0x30, 0xf9, 0x04, 0xdd, 0x59, 0xb0, 0x19, 0xad, 0x2c,
	0x8f, 0xaf, 0x2f, 0xa3, 0xa7, 0xb4, 0x1a, 0x94, 0x3a, 0xe9, 0xa3, 0xed, 0xc6, 0xbe, 0x80, 0xbe,
	0x22, 0x05, 0x99, 0x81, 0xc0, 0xeb, 0x8e, 0xbe, 0x5d, 0x6f, 0x0f, 0xe8, 0xab, 0xe1, 0x09, 0x88,
	0xe4, 0x01, 0xba, 0xbd, 0xdc, 0x46, 0xc1, 0xf0, 0x8d, 0x4e, 0xab, 0x77, 0x6d, 0xb1, 0xc9, 0x41,
	0xa9, 0x87, 0x2c, 0xf9, 0x1c, 0xdd, 0x6d, 0x80, 0xa2, 0x20, 0x53, 0xca, 0x67, 0x1a, 0xbf, 0xe9,
	0xd8, 0xba, 0xee, 0x40, 0x14, 0x47, 0x36, 0x9e, 0x7c, 0x76, 0x75, 0x7b, 0x0c, 0x0a, 0x86, 0x37,
	0x1c, 0xbd, 0xd5, 0xa8, 0xfc, 0x1c, 0x0a, 0xd6, 0xfd, 0x63, 0x0d, 0xed, 0x2c, 0xd9, 0xe4, 0x54,
	0xce, 0xac, 0x68, 0x42, 0x3e, 0x9f, 0x69, 0x08, 0x53, 0xb5, 0x56, 0xdd, 0x71, 0xfd, 0x9d, 0xdb,
	0xa8, 0x9f, 0xe7, 0x47, 0x9e, 0x83, 0x12, 0x84, 0x09, 0xc3, 0xf4, 0x66, 0xdd, 0x14, 0x5a, 0x0d,
	0x5c, 0xd4, 0x8e, 0x32, 0xf8, 0x1e, 0x30, 0x53, 0x39, 0x8b, 0xae, 0x39, 0xdf, 0x3d, 0x74, 0x51,
	0x5d, 0x61, 0x54, 0xe5, 0x5c, 0x69, 0x32, 0xa3, 0x2a, 0x5a, 0x32, 0xa1, 0xcc, 0xb7, 0xe6, 0xb8,
	0xeb, 0x8e, 0xb3, 0x96, 0x3c, 0xa3, 0xcc, 0xf5, 0x66, 0xd1, 0x30, 0x8a, 0x9f, 0xa5, 0x6c, 0xb2,
	0xeb, 0xf5, 0x28, 0x5e, 0x48, 0xb9, 0x80, 0xc3, 0xda, 0x81, 0x33, 0x55, 0x30, 0xa3, 0x1d, 0x37,
	0x6b, 0xfb, 0x0b, 0x05, 0xfd, 0x48, 0x4c, 0xb5, 0xe4, 0xc4, 0x56, 0x04, 0x2f, 0x2a, 0x6f, 0xc4,
	0x2f, 0xe8, 0x86, 0x83, 0x15, 0xc7, 0x1b, 0x9d, 0xb5, 0x5e, 0x7b, 0x77, 0xd2, 0xff, 0x17, 0x6f,
	0x68, 0x7f, 0xe5, 0xed, 0x1a, 0xad, 0xdb, 0x2e, 0x14, 0xef, 0xfe, 0xb9, 0x81, 0xb6, 0x57, 0x5d,
	0x60, 0x7b, 0x77, 0xed, 0x87, 0x82, 0x92, 0x6b, 0x2e, 0x05, 0xde, 0x75, 0x96, 0xda, 0x5d, 0x8e,
	0x42, 0x28, 0x4e, 0x62, 0xaa, 0x33, 0xa7, 0x01, 0xbc, 0x57, 0x33, 0x47, 0x3a, 0x1b, 0xdb, 0x50,
	0xf2, 0x3e, 0xb2, 0x9f, 0xa4, 0x04, 0xe5, 0xaa, 0x3c, 0x74, 0x86, 0x23, 0xa1, 0xd5, 0x8f, 0x3e,
	0x92, 0xbc, 0xe7, 0x01, 0x21, 0x19, 0x10, 0xce, 0xf0, 0x17, 0x0e, 0xd8, 0x10, 0x5a, 0x9d, 0x49,
	0x06, 0xc7, 0x2c, 0xe9, 0xf9, 0xe3, 0x9f, 0x03, 0x2c, 0xaa, 0xec, 0xd7, 0x17, 0xeb, 0x1c, 0xa0,
	0xae, 0xf4, 0xa0, 0x41, 0xc6, 0x72, 0x5f, 0xd6, 0xaf, 0x81, 0x25, 0x43, 0xc9, 0x1d, 0x3f, 0xf0,
	0xcc, 0xcc, 0xf1, 0x23, 0x97, 0xb7, 0xc3, 0x38, 0x35, 0xf3, 0x78, 0x42, 0xc5, 0x44, 0x91, 0x82,
	0x32, 0x52, 0x70, 0x86, 0xbf, 0xaa, 0x4f, 0xe8, 0xd9, 0x44, 0x0d, 0x29, 0x1b, 0x72, 0x16, 0xb1,
	0x99, 0xa6, 0x35, 0xf6, 0x75, 0x8d, 0x9d, 0x68, 0x1a, 0xb0, 0x87, 0x68, 0xa7, 0x59, 0x2d, 0x63,
	0xd4, 0x50, 0x92, 0xca, 0xb9, 0x30, 0xf8, 0x9b, 0xfa, 0xd5, 0xf1, 0x55, 0x4f, 0x6d, 0xee, 0xd0,
	0xa6, 0xa2, 0x2a, 0x16, 0x6f, 0xaa, 0x1e, 0xd7, 0x2a, 0xbf, 0x48, 0x43, 0xf5, 0x04, 0xbd, 0x1b,
	0xd7, 0xe2, 0x82, 0x1b, 0xa2, 0x5f, 0x8b, 0x94, 0xe4, 0x20, 0x58, 0x90, 0x3e, 0x71, 0xce, 0xec,
	0xf8, 0x05, 0x8f, 0x05, 0x37, 0xe3, 0xd7, 0x22, 0x3d, 0x07, 0xc1, 0x96, 0xe4, 0x76, 0xd1, 0x95,
	0xf2, 0xa7, 0xb5, 0xfc, 0x44, 0xd3, 0xbf, 0xcb, 0x3f, 0x44, 0xb7, 0xe2, 0xea, 0x1a, 0x0a, 0x22,
	0x24, 0xfe, 0xd6, 0xb5, 0xda, 0xf6, 0xeb, 0x8d, 0xa1, 0x38, 0x93, 0x71, 0x6a, 0x5c, 0x98, 0x69,
	0xa4, 0xbe, 0xab, 0xa7, 0x76, 0x2c, 0xcc, 0xd4, 0x63, 0xe1, 0x50, 0x99, 0x4c, 0x85, 0xb7, 0xe4,
	0xa0, 0x3e, 0x54, 0x17, 0x99, 0xf2, 0x2f, 0xc9, 0x63, 0xf4, 0x8e, 0x65, 0x52, 0x29, 0x04, 0x31,
	0x92, 0xd0, 0xd4, 0xf0, 0x12, 0x08, 0x35, 0x06, 0xb2, 0xdc, 0x68, 0xfc, 0xcc, 0xdd, 0x33, 0xdb,
	0xed, 0xa1, 0x14, 0xe2, 0x42, 0x1e, 0xb8, 0xfc, 0x41, 0x48, 0x27, 0xfb, 0x08, 0xaf, 0x50, 0xfb,
	0x2b, 0x7a, 0xe8, 0xa4, 0xdb, 0x57, 0xa4, 0xfe, 0x9e, 0xae, 0xd6, 0xc9, 0x1c, 0x84, 0xc6, 0xcf,
	0x57, 0xea, 0x7e, 0xb0, 0xb9, 0xe4, 0x11, 0xba, 0xbf, 0x42, 0x97, 0xce, 0xa4, 0x06, 0x8d, 0x07,
	0x4e, 0x78, 0xef, 0x8a, 0xf0, 0xd0, 0x25, 0xff, 0x41, 0x09, 0x4a, 0x49, 0xa5, 0xf1, 0xd1, 0x4a,
	0xe5, 0xc0, 0x25, 0x93, 0x5f, 0x5b, 0x61, 0x8c, 0x97, 0xe1, 0x95, 0xc6, 0x2f, 0x3a, 0xad, 0x5e,
	0x7b, 0x97, 0xfd, 0x77, 0x4f, 0xcb, 0xe2, 0x2f, 0x82, 0x37, 0xeb, 0x92, 0x8d, 0xdd, 0x43, 0xf2,
	0x5b, 0xcb, 0x3f, 0xc4, 0xca, 0x84, 0x76, 0x2c, 0x88, 0xbf, 0xff, 0x1f, 0xdb, 0xb1, 0x87, 0x73,
	0x64, 0x5c, 0x47, 0x36, 0x3a, 0x59, 0x77, 0xff, 0x07, 0xed, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff,
	0x29, 0x5f, 0x11, 0x28, 0x25, 0x09, 0x00, 0x00,
}
