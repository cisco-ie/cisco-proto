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
// source: ospf_sh_nsr_stats.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_statistics_nsr_stats

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

type OspfShNsrStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNsrStats_KEYS) Reset()         { *m = OspfShNsrStats_KEYS{} }
func (m *OspfShNsrStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNsrStats_KEYS) ProtoMessage()    {}
func (*OspfShNsrStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38cc0a5f94f1aa3, []int{0}
}

func (m *OspfShNsrStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNsrStats_KEYS.Unmarshal(m, b)
}
func (m *OspfShNsrStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNsrStats_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShNsrStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNsrStats_KEYS.Merge(m, src)
}
func (m *OspfShNsrStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNsrStats_KEYS.Size(m)
}
func (m *OspfShNsrStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNsrStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNsrStats_KEYS proto.InternalMessageInfo

func (m *OspfShNsrStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShNsrStats_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type OspfShNsrStatsPri struct {
	NsrSchedPri          uint32   `protobuf:"varint,1,opt,name=nsr_sched_pri,json=nsrSchedPri,proto3" json:"nsr_sched_pri,omitempty"`
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
	return fileDescriptor_e38cc0a5f94f1aa3, []int{1}
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

func (m *OspfShNsrStatsPri) GetNsrSchedPri() uint32 {
	if m != nil {
		return m.NsrSchedPri
	}
	return 0
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
	return fileDescriptor_e38cc0a5f94f1aa3, []int{2}
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

type OspfShNsrStats struct {
	NsrRev                  int32                `protobuf:"zigzag32,50,opt,name=nsr_rev,json=nsrRev,proto3" json:"nsr_rev,omitempty"`
	NsrFsmState             int32                `protobuf:"zigzag32,51,opt,name=nsr_fsm_state,json=nsrFsmState,proto3" json:"nsr_fsm_state,omitempty"`
	NsrVersion              uint32               `protobuf:"varint,52,opt,name=nsr_version,json=nsrVersion,proto3" json:"nsr_version,omitempty"`
	NsrNodeid               uint32               `protobuf:"varint,53,opt,name=nsr_nodeid,json=nsrNodeid,proto3" json:"nsr_nodeid,omitempty"`
	NsrPeerVersion          uint32               `protobuf:"varint,54,opt,name=nsr_peer_version,json=nsrPeerVersion,proto3" json:"nsr_peer_version,omitempty"`
	NsrPeerNodeid           uint32               `protobuf:"varint,55,opt,name=nsr_peer_nodeid,json=nsrPeerNodeid,proto3" json:"nsr_peer_nodeid,omitempty"`
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
	NsrThdSched             *OspfShNsrSchedStats `protobuf:"bytes,71,opt,name=nsr_thd_sched,json=nsrThdSched,proto3" json:"nsr_thd_sched,omitempty"`
	NsrRtrThdSched          *OspfShNsrSchedStats `protobuf:"bytes,72,opt,name=nsr_rtr_thd_sched,json=nsrRtrThdSched,proto3" json:"nsr_rtr_thd_sched,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *OspfShNsrStats) Reset()         { *m = OspfShNsrStats{} }
func (m *OspfShNsrStats) String() string { return proto.CompactTextString(m) }
func (*OspfShNsrStats) ProtoMessage()    {}
func (*OspfShNsrStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e38cc0a5f94f1aa3, []int{3}
}

func (m *OspfShNsrStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNsrStats.Unmarshal(m, b)
}
func (m *OspfShNsrStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNsrStats.Marshal(b, m, deterministic)
}
func (m *OspfShNsrStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNsrStats.Merge(m, src)
}
func (m *OspfShNsrStats) XXX_Size() int {
	return xxx_messageInfo_OspfShNsrStats.Size(m)
}
func (m *OspfShNsrStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNsrStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNsrStats proto.InternalMessageInfo

func (m *OspfShNsrStats) GetNsrRev() int32 {
	if m != nil {
		return m.NsrRev
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrFsmState() int32 {
	if m != nil {
		return m.NsrFsmState
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrVersion() uint32 {
	if m != nil {
		return m.NsrVersion
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrNodeid() uint32 {
	if m != nil {
		return m.NsrNodeid
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrPeerVersion() uint32 {
	if m != nil {
		return m.NsrPeerVersion
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrPeerNodeid() uint32 {
	if m != nil {
		return m.NsrPeerNodeid
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrMtu() uint32 {
	if m != nil {
		return m.NsrMtu
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrNbrQadQid() uint32 {
	if m != nil {
		return m.NsrNbrQadQid
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrLsaQadQid() uint32 {
	if m != nil {
		return m.NsrLsaQadQid
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrNbrQadMdataCount() uint32 {
	if m != nil {
		return m.NsrNbrQadMdataCount
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrLsaQadMdataCount() uint32 {
	if m != nil {
		return m.NsrLsaQadMdataCount
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrNbrInitSyncPendCount() int32 {
	if m != nil {
		return m.NsrNbrInitSyncPendCount
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrLsaInitSyncPendCount() int32 {
	if m != nil {
		return m.NsrLsaInitSyncPendCount
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrNbrSeqNo() uint32 {
	if m != nil {
		return m.NsrNbrSeqNo
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrIntfSeqNo() uint32 {
	if m != nil {
		return m.NsrIntfSeqNo
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrTmrQuant() int32 {
	if m != nil {
		return m.NsrTmrQuant
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrConnToActiveAttempts() uint64 {
	if m != nil {
		return m.NsrConnToActiveAttempts
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrConnToActiveFails() uint64 {
	if m != nil {
		return m.NsrConnToActiveFails
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrConnToActiveOpens() uint64 {
	if m != nil {
		return m.NsrConnToActiveOpens
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrConnToActiveCloses() uint64 {
	if m != nil {
		return m.NsrConnToActiveCloses
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrConnToActiveErrors() uint64 {
	if m != nil {
		return m.NsrConnToActiveErrors
	}
	return 0
}

func (m *OspfShNsrStats) GetNsrThdSched() *OspfShNsrSchedStats {
	if m != nil {
		return m.NsrThdSched
	}
	return nil
}

func (m *OspfShNsrStats) GetNsrRtrThdSched() *OspfShNsrSchedStats {
	if m != nil {
		return m.NsrRtrThdSched
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShNsrStats_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.nsr_stats.ospf_sh_nsr_stats_KEYS")
	proto.RegisterType((*OspfShNsrStatsPri)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.nsr_stats.ospf_sh_nsr_stats_pri")
	proto.RegisterType((*OspfShNsrSchedStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.nsr_stats.ospf_sh_nsr_sched_stats")
	proto.RegisterType((*OspfShNsrStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.nsr_stats.ospf_sh_nsr_stats")
}

func init() { proto.RegisterFile("ospf_sh_nsr_stats.proto", fileDescriptor_e38cc0a5f94f1aa3) }

var fileDescriptor_e38cc0a5f94f1aa3 = []byte{
	// 947 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0xc7, 0xc7, 0x24, 0x4d, 0x1a, 0xa5, 0x69, 0x1a, 0xf5, 0x23, 0xea, 0x00, 0x43, 0x30, 0x03,
	0x18, 0x18, 0x7c, 0x91, 0x84, 0x52, 0xa0, 0x05, 0xd2, 0xd4, 0x29, 0x19, 0x12, 0xe3, 0x8f, 0x4c,
	0x67, 0xb8, 0xd2, 0xc8, 0x2b, 0xb9, 0xd1, 0xd4, 0x2b, 0xed, 0x4a, 0xf2, 0xce, 0xf6, 0x8e, 0x17,
	0xe0, 0x8e, 0x47, 0xe0, 0xb1, 0x78, 0x04, 0x1e, 0x82, 0xd1, 0x91, 0x76, 0xed, 0x38, 0xe6, 0x0e,
	0x7a, 0x93, 0xb1, 0xcf, 0xf9, 0xfd, 0x8f, 0x8e, 0xce, 0x5f, 0x92, 0x83, 0x76, 0xb5, 0xcd, 0xc6,
	0xd4, 0x5e, 0x52, 0x65, 0x0d, 0xb5, 0x8e, 0x39, 0xdb, 0xce, 0x8c, 0x76, 0x1a, 0xf7, 0x12, 0x69,
	0x13, 0x4d, 0xa5, 0xb6, 0xb4, 0x34, 0x54, 0x66, 0xc5, 0x21, 0x05, 0x54, 0x67, 0xc2, 0xb4, 0xfd,
	0x27, 0xcf, 0x25, 0xc2, 0x5a, 0x61, 0xab, 0x4f, 0xed, 0xc2, 0x8c, 0xe1, 0x4f, 0xdb, 0x17, 0x92,
	0xd6, 0xc9, 0xc4, 0xb6, 0xeb, 0xba, 0xcd, 0x97, 0xe8, 0xc1, 0xb5, 0xc5, 0xe8, 0xcf, 0x9d, 0x5f,
	0x87, 0xf8, 0x43, 0x74, 0x2b, 0x96, 0xa0, 0x8a, 0xa5, 0x82, 0x34, 0xf6, 0x1a, 0xad, 0x8d, 0xc1,
	0x66, 0x8c, 0x75, 0x59, 0x2a, 0xf0, 0x43, 0x74, 0xb3, 0x30, 0xe3, 0x90, 0x7e, 0x07, 0xd2, 0xeb,
	0x85, 0x19, 0xfb, 0x54, 0xf3, 0xcf, 0x15, 0x74, 0xff, 0x7a, 0xe1, 0xcc, 0x48, 0xdc, 0x44, 0x5b,
	0x10, 0x48, 0x2e, 0x05, 0xf7, 0x01, 0x28, 0xbc, 0x35, 0xd8, 0x54, 0xd6, 0x0c, 0x7d, 0xac, 0x67,
	0x24, 0xfe, 0x04, 0x6d, 0xcf, 0x98, 0x7c, 0xca, 0x94, 0x83, 0xfa, 0x5b, 0x83, 0xad, 0x8a, 0xea,
	0xfb, 0x20, 0x3e, 0x40, 0x0f, 0x66, 0x9c, 0x11, 0x29, 0x93, 0x2a, 0xe2, 0x2b, 0x80, 0xdf, 0xad,
	0xf0, 0x01, 0xe4, 0x82, 0xe8, 0x73, 0x84, 0x67, 0x22, 0x51, 0x58, 0xea, 0x45, 0x64, 0x15, 0x04,
	0xb7, 0x2b, 0x41, 0xa7, 0xb0, 0xa7, 0xaa, 0x8f, 0x3f, 0x43, 0x3b, 0x33, 0x36, 0x65, 0xa5, 0xe7,
	0xc9, 0x8d, 0xab, 0xe8, 0x39, 0x2b, 0x3b, 0x85, 0xc5, 0x6d, 0x74, 0x6f, 0x6e, 0x5f, 0x82, 0xbd,
	0xa6, 0x39, 0x9d, 0x08, 0x45, 0xd6, 0x80, 0xbe, 0x53, 0x6f, 0x4f, 0xb0, 0xd7, 0xfd, 0x33, 0xa1,
	0xf0, 0xa7, 0xe8, 0xce, 0xd5, 0x36, 0x72, 0x4e, 0xd6, 0xf7, 0x1a, 0xad, 0xd5, 0xd9, 0x26, 0x3b,
	0x85, 0xed, 0x73, 0xfc, 0x25, 0xba, 0x3b, 0x07, 0xaa, 0x9c, 0x8e, 0x99, 0x9c, 0x58, 0x72, 0x13,
	0xd8, 0xba, 0x6e, 0x47, 0xe5, 0x27, 0x3e, 0x8e, 0xbf, 0x58, 0xdc, 0x1e, 0x17, 0x39, 0x27, 0x1b,
	0x40, 0x6f, 0xcf, 0x55, 0x7e, 0x2e, 0x72, 0xde, 0xfc, 0x6b, 0x65, 0xe1, 0xb0, 0x81, 0x0a, 0xcc,
	0xaa, 0x4c, 0xc8, 0xa6, 0x13, 0x2b, 0xe2, 0x54, 0xbd, 0x55, 0x3b, 0xd0, 0x5f, 0xcf, 0x47, 0xc3,
	0x3c, 0x3f, 0x0e, 0x9c, 0x28, 0x84, 0x72, 0x71, 0x98, 0xc1, 0xac, 0x5b, 0xca, 0x9a, 0x0e, 0x44,
	0xfd, 0x28, 0xa3, 0xef, 0x11, 0x73, 0x25, 0x58, 0xb4, 0x0a, 0xbe, 0x07, 0xe8, 0xa2, 0x5c, 0x60,
	0x4c, 0x09, 0xae, 0xcc, 0x33, 0x83, 0xb2, 0xb2, 0x64, 0xc4, 0x78, 0x68, 0x0d, 0xb8, 0x1b, 0xc0,
	0x79, 0x4b, 0x9e, 0x31, 0x0e, 0xbd, 0x79, 0x34, 0x8e, 0xe2, 0x95, 0xd6, 0xf3, 0xec, 0x5a, 0x3d,
	0x8a, 0x17, 0x5a, 0xcf, 0xe0, 0xb8, 0x76, 0xe4, 0x5c, 0x19, 0xcd, 0xd8, 0xac, 0x36, 0xeb, 0xfb,
	0x8b, 0x05, 0xc3, 0x48, 0x5c, 0x79, 0xc5, 0x89, 0xed, 0x0a, 0xbc, 0x28, 0x83, 0x11, 0xbf, 0x35,
	0xd0, 0x3a, 0xd0, 0x46, 0x92, 0x8d, 0xbd, 0x95, 0xd6, 0xe6, 0xfe, 0xab, 0xf6, 0x7f, 0x7d, 0x7f,
	0xdb, 0x4b, 0xef, 0xd8, 0x60, 0xcd, 0xf7, 0x62, 0x64, 0xf3, 0xef, 0x0d, 0xb4, 0x73, 0x8d, 0xc0,
	0xbb, 0xa1, 0x2f, 0x23, 0x0a, 0xb2, 0x0f, 0x86, 0x7a, 0x7c, 0x20, 0x8a, 0x6a, 0x04, 0x63, 0x9b,
	0x02, 0x29, 0xc8, 0x01, 0xa4, 0xfd, 0x08, 0x4e, 0x6c, 0x3a, 0xf4, 0x21, 0xfc, 0x01, 0xf2, 0x5f,
	0x69, 0x21, 0x8c, 0x95, 0x5a, 0x91, 0x43, 0x70, 0x1a, 0x29, 0x6b, 0x5e, 0x86, 0x08, 0x7e, 0x1f,
	0xf9, 0x6f, 0x54, 0x69, 0x2e, 0x24, 0x27, 0x5f, 0x41, 0x7e, 0x43, 0x59, 0xd3, 0x85, 0x00, 0x6e,
	0x85, 0x63, 0x9f, 0x09, 0x31, 0x2b, 0xf2, 0xa8, 0xbe, 0x50, 0x3d, 0x21, 0xea, 0x42, 0xd5, 0xf9,
	0xf3, 0x64, 0xac, 0xf6, 0x75, 0xfd, 0x08, 0x78, 0x30, 0x56, 0x8c, 0xdb, 0x49, 0xdd, 0x94, 0x3c,
	0x86, 0xbc, 0xdf, 0xce, 0xb9, 0x9b, 0x56, 0x07, 0x53, 0x8d, 0x0c, 0xcd, 0x19, 0xa7, 0xb9, 0xe4,
	0xe4, 0x9b, 0xfa, 0x60, 0x76, 0x47, 0xa6, 0xcf, 0x78, 0x5f, 0xf2, 0x0a, 0x9b, 0x58, 0x56, 0x63,
	0xdf, 0xd6, 0xd8, 0x99, 0x65, 0x11, 0x3b, 0x44, 0xbb, 0xf3, 0xd5, 0x52, 0xce, 0x1c, 0xa3, 0x89,
	0x9e, 0x2a, 0x47, 0xbe, 0xab, 0x1f, 0x9b, 0x50, 0xf5, 0xdc, 0xe7, 0x8e, 0x7d, 0xaa, 0x52, 0x55,
	0xc5, 0xe7, 0x55, 0x4f, 0x6a, 0x55, 0x58, 0x64, 0x4e, 0xf5, 0x14, 0xbd, 0x57, 0xad, 0x25, 0x95,
	0x74, 0xd4, 0xbe, 0x51, 0x09, 0xcd, 0x84, 0xe2, 0x51, 0xfa, 0x14, 0x7c, 0xd9, 0x0d, 0x0b, 0x9e,
	0x2a, 0xe9, 0x86, 0x6f, 0x54, 0xd2, 0x13, 0x8a, 0x5f, 0x91, 0xfb, 0x45, 0x97, 0xca, 0xbf, 0xaf,
	0xe5, 0x67, 0x96, 0x5d, 0x97, 0x7f, 0x84, 0x6e, 0x57, 0xab, 0x5b, 0x91, 0x53, 0xa5, 0xc9, 0x0f,
	0xf5, 0x13, 0xdd, 0x1d, 0x99, 0xa1, 0xc8, 0xbb, 0xba, 0x9a, 0x9a, 0x54, 0x6e, 0x5c, 0x51, 0x3f,
	0xd6, 0x53, 0x3b, 0x55, 0x6e, 0x1c, 0xb0, 0x78, 0xa4, 0x5c, 0x6a, 0xe2, 0x13, 0x72, 0x54, 0x1f,
	0xa9, 0x8b, 0xd4, 0x84, 0x07, 0xe4, 0x09, 0x7a, 0xd7, 0x33, 0x89, 0x56, 0x8a, 0x3a, 0x4d, 0x59,
	0xe2, 0x64, 0x21, 0x28, 0x73, 0x4e, 0xa4, 0x99, 0xb3, 0xe4, 0x19, 0x5c, 0x2f, 0xdf, 0xed, 0xb1,
	0x56, 0xea, 0x42, 0x1f, 0x41, 0xfe, 0x28, 0xa6, 0xf1, 0x23, 0x44, 0x96, 0xa8, 0xc3, 0xcd, 0x3c,
	0x06, 0xe9, 0xbd, 0x05, 0x69, 0xb8, 0x9e, 0xcb, 0x75, 0x3a, 0x13, 0xca, 0x92, 0xe7, 0x4b, 0x75,
	0xbf, 0xf8, 0x1c, 0x7e, 0x8c, 0x1e, 0x2e, 0xd1, 0x25, 0x13, 0x6d, 0x85, 0x25, 0x1d, 0x10, 0xde,
	0x5f, 0x10, 0x1e, 0x43, 0xf2, 0x5f, 0x94, 0xc2, 0x18, 0x6d, 0x2c, 0x39, 0x59, 0xaa, 0xec, 0x40,
	0x12, 0xff, 0xde, 0x88, 0x63, 0xbc, 0xe4, 0xe1, 0x89, 0x26, 0x2f, 0xf6, 0x1a, 0xad, 0xcd, 0x7d,
	0xf9, 0x3f, 0x3f, 0x28, 0xb3, 0x5f, 0x83, 0xe0, 0xd8, 0x25, 0x87, 0x5f, 0x0f, 0xfc, 0x47, 0x23,
	0x3c, 0xc2, 0xc6, 0xcd, 0xf7, 0xf4, 0xd3, 0xdb, 0xee, 0xc9, 0x1f, 0xd3, 0x81, 0xab, 0xdb, 0x1a,
	0xad, 0xc1, 0x7f, 0x49, 0x07, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x6e, 0xb5, 0xc8, 0x40,
	0x09, 0x00, 0x00,
}
