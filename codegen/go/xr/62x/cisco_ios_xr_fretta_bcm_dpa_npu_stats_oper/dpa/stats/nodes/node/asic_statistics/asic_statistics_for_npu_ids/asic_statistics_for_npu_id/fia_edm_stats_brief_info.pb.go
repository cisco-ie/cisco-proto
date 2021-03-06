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
// source: fia_edm_stats_brief_info.proto

package cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper_dpa_stats_nodes_node_asic_statistics_asic_statistics_for_npu_ids_asic_statistics_for_npu_id

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

type FiaEdmStatsBriefInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NpuId                uint32   `protobuf:"varint,2,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FiaEdmStatsBriefInfo_KEYS) Reset()         { *m = FiaEdmStatsBriefInfo_KEYS{} }
func (m *FiaEdmStatsBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*FiaEdmStatsBriefInfo_KEYS) ProtoMessage()    {}
func (*FiaEdmStatsBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b969035e2969373, []int{0}
}

func (m *FiaEdmStatsBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *FiaEdmStatsBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *FiaEdmStatsBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS.Merge(m, src)
}
func (m *FiaEdmStatsBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS.Size(m)
}
func (m *FiaEdmStatsBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmStatsBriefInfo_KEYS proto.InternalMessageInfo

func (m *FiaEdmStatsBriefInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FiaEdmStatsBriefInfo_KEYS) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

type FiaEdmDeviceStatsAsicBriefInfo struct {
	NbiRxTotalByteCnt           uint64   `protobuf:"varint,1,opt,name=nbi_rx_total_byte_cnt,json=nbiRxTotalByteCnt,proto3" json:"nbi_rx_total_byte_cnt,omitempty"`
	NbiRxTotalPktCnt            uint64   `protobuf:"varint,2,opt,name=nbi_rx_total_pkt_cnt,json=nbiRxTotalPktCnt,proto3" json:"nbi_rx_total_pkt_cnt,omitempty"`
	IreCpuPktCnt                uint64   `protobuf:"varint,3,opt,name=ire_cpu_pkt_cnt,json=ireCpuPktCnt,proto3" json:"ire_cpu_pkt_cnt,omitempty"`
	IreNifPktCnt                uint64   `protobuf:"varint,4,opt,name=ire_nif_pkt_cnt,json=ireNifPktCnt,proto3" json:"ire_nif_pkt_cnt,omitempty"`
	IreOampPktCnt               uint64   `protobuf:"varint,5,opt,name=ire_oamp_pkt_cnt,json=ireOampPktCnt,proto3" json:"ire_oamp_pkt_cnt,omitempty"`
	IreOlpPktCnt                uint64   `protobuf:"varint,6,opt,name=ire_olp_pkt_cnt,json=ireOlpPktCnt,proto3" json:"ire_olp_pkt_cnt,omitempty"`
	IreRcyPktCnt                uint64   `protobuf:"varint,7,opt,name=ire_rcy_pkt_cnt,json=ireRcyPktCnt,proto3" json:"ire_rcy_pkt_cnt,omitempty"`
	IreFdtIfCnt                 uint64   `protobuf:"varint,8,opt,name=ire_fdt_if_cnt,json=ireFdtIfCnt,proto3" json:"ire_fdt_if_cnt,omitempty"`
	IdrMmuIfCnt                 uint64   `protobuf:"varint,9,opt,name=idr_mmu_if_cnt,json=idrMmuIfCnt,proto3" json:"idr_mmu_if_cnt,omitempty"`
	IdrOcbIfCnt                 uint64   `protobuf:"varint,10,opt,name=idr_ocb_if_cnt,json=idrOcbIfCnt,proto3" json:"idr_ocb_if_cnt,omitempty"`
	IqmEnqueuePktCnt            uint64   `protobuf:"varint,11,opt,name=iqm_enqueue_pkt_cnt,json=iqmEnqueuePktCnt,proto3" json:"iqm_enqueue_pkt_cnt,omitempty"`
	IqmDequeuePktCnt            uint64   `protobuf:"varint,12,opt,name=iqm_dequeue_pkt_cnt,json=iqmDequeuePktCnt,proto3" json:"iqm_dequeue_pkt_cnt,omitempty"`
	IqmDeletedPktCnt            uint64   `protobuf:"varint,13,opt,name=iqm_deleted_pkt_cnt,json=iqmDeletedPktCnt,proto3" json:"iqm_deleted_pkt_cnt,omitempty"`
	IqmEnqDiscardedPktCnt       uint64   `protobuf:"varint,14,opt,name=iqm_enq_discarded_pkt_cnt,json=iqmEnqDiscardedPktCnt,proto3" json:"iqm_enq_discarded_pkt_cnt,omitempty"`
	IptEgqPktCnt                uint64   `protobuf:"varint,15,opt,name=ipt_egq_pkt_cnt,json=iptEgqPktCnt,proto3" json:"ipt_egq_pkt_cnt,omitempty"`
	IptEnqPktCnt                uint64   `protobuf:"varint,16,opt,name=ipt_enq_pkt_cnt,json=iptEnqPktCnt,proto3" json:"ipt_enq_pkt_cnt,omitempty"`
	IptFdtPktCnt                uint64   `protobuf:"varint,17,opt,name=ipt_fdt_pkt_cnt,json=iptFdtPktCnt,proto3" json:"ipt_fdt_pkt_cnt,omitempty"`
	IptCfgEventCnt              uint64   `protobuf:"varint,18,opt,name=ipt_cfg_event_cnt,json=iptCfgEventCnt,proto3" json:"ipt_cfg_event_cnt,omitempty"`
	IptCfgByteCnt               uint64   `protobuf:"varint,19,opt,name=ipt_cfg_byte_cnt,json=iptCfgByteCnt,proto3" json:"ipt_cfg_byte_cnt,omitempty"`
	FdtIptDescCellCnt           uint64   `protobuf:"varint,20,opt,name=fdt_ipt_desc_cell_cnt,json=fdtIptDescCellCnt,proto3" json:"fdt_ipt_desc_cell_cnt,omitempty"`
	FdtIreDescCellCnt           uint64   `protobuf:"varint,21,opt,name=fdt_ire_desc_cell_cnt,json=fdtIreDescCellCnt,proto3" json:"fdt_ire_desc_cell_cnt,omitempty"`
	FdtTransmittedDataCellsCnt  uint64   `protobuf:"varint,22,opt,name=fdt_transmitted_data_cells_cnt,json=fdtTransmittedDataCellsCnt,proto3" json:"fdt_transmitted_data_cells_cnt,omitempty"`
	FdrP1CellInCnt              uint64   `protobuf:"varint,23,opt,name=fdr_p1_cell_in_cnt,json=fdrP1CellInCnt,proto3" json:"fdr_p1_cell_in_cnt,omitempty"`
	FdrP2CellInCnt              uint64   `protobuf:"varint,24,opt,name=fdr_p2_cell_in_cnt,json=fdrP2CellInCnt,proto3" json:"fdr_p2_cell_in_cnt,omitempty"`
	FdrP3CellInCnt              uint64   `protobuf:"varint,25,opt,name=fdr_p3_cell_in_cnt,json=fdrP3CellInCnt,proto3" json:"fdr_p3_cell_in_cnt,omitempty"`
	FdrCellInCntTotal           uint64   `protobuf:"varint,26,opt,name=fdr_cell_in_cnt_total,json=fdrCellInCntTotal,proto3" json:"fdr_cell_in_cnt_total,omitempty"`
	FdaCellsInCntP1             uint64   `protobuf:"varint,27,opt,name=fda_cells_in_cnt_p1,json=fdaCellsInCntP1,proto3" json:"fda_cells_in_cnt_p1,omitempty"`
	FdaCellsInCntP2             uint64   `protobuf:"varint,28,opt,name=fda_cells_in_cnt_p2,json=fdaCellsInCntP2,proto3" json:"fda_cells_in_cnt_p2,omitempty"`
	FdaCellsInCntP3             uint64   `protobuf:"varint,29,opt,name=fda_cells_in_cnt_p3,json=fdaCellsInCntP3,proto3" json:"fda_cells_in_cnt_p3,omitempty"`
	FdaCellsInTdmCnt            uint64   `protobuf:"varint,30,opt,name=fda_cells_in_tdm_cnt,json=fdaCellsInTdmCnt,proto3" json:"fda_cells_in_tdm_cnt,omitempty"`
	FdaCellsInMeshmcCnt         uint64   `protobuf:"varint,31,opt,name=fda_cells_in_meshmc_cnt,json=fdaCellsInMeshmcCnt,proto3" json:"fda_cells_in_meshmc_cnt,omitempty"`
	FdaCellsInIptCnt            uint64   `protobuf:"varint,32,opt,name=fda_cells_in_ipt_cnt,json=fdaCellsInIptCnt,proto3" json:"fda_cells_in_ipt_cnt,omitempty"`
	FdaCellsOutCntP1            uint64   `protobuf:"varint,33,opt,name=fda_cells_out_cnt_p1,json=fdaCellsOutCntP1,proto3" json:"fda_cells_out_cnt_p1,omitempty"`
	FdaCellsOutCntP2            uint64   `protobuf:"varint,34,opt,name=fda_cells_out_cnt_p2,json=fdaCellsOutCntP2,proto3" json:"fda_cells_out_cnt_p2,omitempty"`
	FdaCellsOutCntP3            uint64   `protobuf:"varint,35,opt,name=fda_cells_out_cnt_p3,json=fdaCellsOutCntP3,proto3" json:"fda_cells_out_cnt_p3,omitempty"`
	FdaCellsOutTdmCnt           uint64   `protobuf:"varint,36,opt,name=fda_cells_out_tdm_cnt,json=fdaCellsOutTdmCnt,proto3" json:"fda_cells_out_tdm_cnt,omitempty"`
	FdaCellsOutMeshmcCnt        uint64   `protobuf:"varint,37,opt,name=fda_cells_out_meshmc_cnt,json=fdaCellsOutMeshmcCnt,proto3" json:"fda_cells_out_meshmc_cnt,omitempty"`
	FdaCellsOutIptCnt           uint64   `protobuf:"varint,38,opt,name=fda_cells_out_ipt_cnt,json=fdaCellsOutIptCnt,proto3" json:"fda_cells_out_ipt_cnt,omitempty"`
	FdaEgqDropCnt               uint64   `protobuf:"varint,39,opt,name=fda_egq_drop_cnt,json=fdaEgqDropCnt,proto3" json:"fda_egq_drop_cnt,omitempty"`
	FdaEgqMeshmcDropCnt         uint64   `protobuf:"varint,40,opt,name=fda_egq_meshmc_drop_cnt,json=fdaEgqMeshmcDropCnt,proto3" json:"fda_egq_meshmc_drop_cnt,omitempty"`
	EgqFqpPktCnt                uint64   `protobuf:"varint,41,opt,name=egq_fqp_pkt_cnt,json=egqFqpPktCnt,proto3" json:"egq_fqp_pkt_cnt,omitempty"`
	EgqPqpUcPktCnt              uint64   `protobuf:"varint,42,opt,name=egq_pqp_uc_pkt_cnt,json=egqPqpUcPktCnt,proto3" json:"egq_pqp_uc_pkt_cnt,omitempty"`
	EgqPqpDiscardUcPktCnt       uint64   `protobuf:"varint,43,opt,name=egq_pqp_discard_uc_pkt_cnt,json=egqPqpDiscardUcPktCnt,proto3" json:"egq_pqp_discard_uc_pkt_cnt,omitempty"`
	EgqPqpUcBytesCnt            uint64   `protobuf:"varint,44,opt,name=egq_pqp_uc_bytes_cnt,json=egqPqpUcBytesCnt,proto3" json:"egq_pqp_uc_bytes_cnt,omitempty"`
	EgqPqpMcPktCnt              uint64   `protobuf:"varint,45,opt,name=egq_pqp_mc_pkt_cnt,json=egqPqpMcPktCnt,proto3" json:"egq_pqp_mc_pkt_cnt,omitempty"`
	EgqPqpDiscardMcPktCnt       uint64   `protobuf:"varint,46,opt,name=egq_pqp_discard_mc_pkt_cnt,json=egqPqpDiscardMcPktCnt,proto3" json:"egq_pqp_discard_mc_pkt_cnt,omitempty"`
	EgqPqpMcBytesCnt            uint64   `protobuf:"varint,47,opt,name=egq_pqp_mc_bytes_cnt,json=egqPqpMcBytesCnt,proto3" json:"egq_pqp_mc_bytes_cnt,omitempty"`
	EgqEhpUcPktCnt              uint64   `protobuf:"varint,48,opt,name=egq_ehp_uc_pkt_cnt,json=egqEhpUcPktCnt,proto3" json:"egq_ehp_uc_pkt_cnt,omitempty"`
	EgqEhpMcHighPktCnt          uint64   `protobuf:"varint,49,opt,name=egq_ehp_mc_high_pkt_cnt,json=egqEhpMcHighPktCnt,proto3" json:"egq_ehp_mc_high_pkt_cnt,omitempty"`
	EgqEhpMcLowPktCnt           uint64   `protobuf:"varint,50,opt,name=egq_ehp_mc_low_pkt_cnt,json=egqEhpMcLowPktCnt,proto3" json:"egq_ehp_mc_low_pkt_cnt,omitempty"`
	EgqDeletedPktCnt            uint64   `protobuf:"varint,51,opt,name=egq_deleted_pkt_cnt,json=egqDeletedPktCnt,proto3" json:"egq_deleted_pkt_cnt,omitempty"`
	EgqEhpMcHighDiscardCnt      uint64   `protobuf:"varint,52,opt,name=egq_ehp_mc_high_discard_cnt,json=egqEhpMcHighDiscardCnt,proto3" json:"egq_ehp_mc_high_discard_cnt,omitempty"`
	EgqEhpMcLowDiscardCnt       uint64   `protobuf:"varint,53,opt,name=egq_ehp_mc_low_discard_cnt,json=egqEhpMcLowDiscardCnt,proto3" json:"egq_ehp_mc_low_discard_cnt,omitempty"`
	EgqErppLagPruningDiscardCnt uint64   `protobuf:"varint,54,opt,name=egq_erpp_lag_pruning_discard_cnt,json=egqErppLagPruningDiscardCnt,proto3" json:"egq_erpp_lag_pruning_discard_cnt,omitempty"`
	EgqErppPmfDiscardCnt        uint64   `protobuf:"varint,55,opt,name=egq_erpp_pmf_discard_cnt,json=egqErppPmfDiscardCnt,proto3" json:"egq_erpp_pmf_discard_cnt,omitempty"`
	EgqErppVlanMbrDiscardCnt    uint64   `protobuf:"varint,56,opt,name=egq_erpp_vlan_mbr_discard_cnt,json=egqErppVlanMbrDiscardCnt,proto3" json:"egq_erpp_vlan_mbr_discard_cnt,omitempty"`
	EpniEpeByteCnt              uint64   `protobuf:"varint,57,opt,name=epni_epe_byte_cnt,json=epniEpeByteCnt,proto3" json:"epni_epe_byte_cnt,omitempty"`
	EpniEpePktCnt               uint64   `protobuf:"varint,58,opt,name=epni_epe_pkt_cnt,json=epniEpePktCnt,proto3" json:"epni_epe_pkt_cnt,omitempty"`
	EpniEpeDiscardCnt           uint64   `protobuf:"varint,59,opt,name=epni_epe_discard_cnt,json=epniEpeDiscardCnt,proto3" json:"epni_epe_discard_cnt,omitempty"`
	NbiTxTotalByteCnt           uint64   `protobuf:"varint,60,opt,name=nbi_tx_total_byte_cnt,json=nbiTxTotalByteCnt,proto3" json:"nbi_tx_total_byte_cnt,omitempty"`
	NbiTxTotalPktCnt            uint64   `protobuf:"varint,61,opt,name=nbi_tx_total_pkt_cnt,json=nbiTxTotalPktCnt,proto3" json:"nbi_tx_total_pkt_cnt,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) Reset()         { *m = FiaEdmDeviceStatsAsicBriefInfo{} }
func (m *FiaEdmDeviceStatsAsicBriefInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmDeviceStatsAsicBriefInfo) ProtoMessage()    {}
func (*FiaEdmDeviceStatsAsicBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b969035e2969373, []int{1}
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo.Unmarshal(m, b)
}
func (m *FiaEdmDeviceStatsAsicBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmDeviceStatsAsicBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo.Merge(m, src)
}
func (m *FiaEdmDeviceStatsAsicBriefInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo.Size(m)
}
func (m *FiaEdmDeviceStatsAsicBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmDeviceStatsAsicBriefInfo proto.InternalMessageInfo

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetNbiRxTotalByteCnt() uint64 {
	if m != nil {
		return m.NbiRxTotalByteCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetNbiRxTotalPktCnt() uint64 {
	if m != nil {
		return m.NbiRxTotalPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreCpuPktCnt() uint64 {
	if m != nil {
		return m.IreCpuPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreNifPktCnt() uint64 {
	if m != nil {
		return m.IreNifPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreOampPktCnt() uint64 {
	if m != nil {
		return m.IreOampPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreOlpPktCnt() uint64 {
	if m != nil {
		return m.IreOlpPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreRcyPktCnt() uint64 {
	if m != nil {
		return m.IreRcyPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIreFdtIfCnt() uint64 {
	if m != nil {
		return m.IreFdtIfCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIdrMmuIfCnt() uint64 {
	if m != nil {
		return m.IdrMmuIfCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIdrOcbIfCnt() uint64 {
	if m != nil {
		return m.IdrOcbIfCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIqmEnqueuePktCnt() uint64 {
	if m != nil {
		return m.IqmEnqueuePktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIqmDequeuePktCnt() uint64 {
	if m != nil {
		return m.IqmDequeuePktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIqmDeletedPktCnt() uint64 {
	if m != nil {
		return m.IqmDeletedPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIqmEnqDiscardedPktCnt() uint64 {
	if m != nil {
		return m.IqmEnqDiscardedPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIptEgqPktCnt() uint64 {
	if m != nil {
		return m.IptEgqPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIptEnqPktCnt() uint64 {
	if m != nil {
		return m.IptEnqPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIptFdtPktCnt() uint64 {
	if m != nil {
		return m.IptFdtPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIptCfgEventCnt() uint64 {
	if m != nil {
		return m.IptCfgEventCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetIptCfgByteCnt() uint64 {
	if m != nil {
		return m.IptCfgByteCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdtIptDescCellCnt() uint64 {
	if m != nil {
		return m.FdtIptDescCellCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdtIreDescCellCnt() uint64 {
	if m != nil {
		return m.FdtIreDescCellCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdtTransmittedDataCellsCnt() uint64 {
	if m != nil {
		return m.FdtTransmittedDataCellsCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdrP1CellInCnt() uint64 {
	if m != nil {
		return m.FdrP1CellInCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdrP2CellInCnt() uint64 {
	if m != nil {
		return m.FdrP2CellInCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdrP3CellInCnt() uint64 {
	if m != nil {
		return m.FdrP3CellInCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdrCellInCntTotal() uint64 {
	if m != nil {
		return m.FdrCellInCntTotal
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInCntP1() uint64 {
	if m != nil {
		return m.FdaCellsInCntP1
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInCntP2() uint64 {
	if m != nil {
		return m.FdaCellsInCntP2
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInCntP3() uint64 {
	if m != nil {
		return m.FdaCellsInCntP3
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInTdmCnt() uint64 {
	if m != nil {
		return m.FdaCellsInTdmCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInMeshmcCnt() uint64 {
	if m != nil {
		return m.FdaCellsInMeshmcCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsInIptCnt() uint64 {
	if m != nil {
		return m.FdaCellsInIptCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutCntP1() uint64 {
	if m != nil {
		return m.FdaCellsOutCntP1
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutCntP2() uint64 {
	if m != nil {
		return m.FdaCellsOutCntP2
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutCntP3() uint64 {
	if m != nil {
		return m.FdaCellsOutCntP3
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutTdmCnt() uint64 {
	if m != nil {
		return m.FdaCellsOutTdmCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutMeshmcCnt() uint64 {
	if m != nil {
		return m.FdaCellsOutMeshmcCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaCellsOutIptCnt() uint64 {
	if m != nil {
		return m.FdaCellsOutIptCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaEgqDropCnt() uint64 {
	if m != nil {
		return m.FdaEgqDropCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetFdaEgqMeshmcDropCnt() uint64 {
	if m != nil {
		return m.FdaEgqMeshmcDropCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqFqpPktCnt() uint64 {
	if m != nil {
		return m.EgqFqpPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpUcPktCnt() uint64 {
	if m != nil {
		return m.EgqPqpUcPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpDiscardUcPktCnt() uint64 {
	if m != nil {
		return m.EgqPqpDiscardUcPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpUcBytesCnt() uint64 {
	if m != nil {
		return m.EgqPqpUcBytesCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpMcPktCnt() uint64 {
	if m != nil {
		return m.EgqPqpMcPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpDiscardMcPktCnt() uint64 {
	if m != nil {
		return m.EgqPqpDiscardMcPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqPqpMcBytesCnt() uint64 {
	if m != nil {
		return m.EgqPqpMcBytesCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqEhpUcPktCnt() uint64 {
	if m != nil {
		return m.EgqEhpUcPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqEhpMcHighPktCnt() uint64 {
	if m != nil {
		return m.EgqEhpMcHighPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqEhpMcLowPktCnt() uint64 {
	if m != nil {
		return m.EgqEhpMcLowPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqDeletedPktCnt() uint64 {
	if m != nil {
		return m.EgqDeletedPktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqEhpMcHighDiscardCnt() uint64 {
	if m != nil {
		return m.EgqEhpMcHighDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqEhpMcLowDiscardCnt() uint64 {
	if m != nil {
		return m.EgqEhpMcLowDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqErppLagPruningDiscardCnt() uint64 {
	if m != nil {
		return m.EgqErppLagPruningDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqErppPmfDiscardCnt() uint64 {
	if m != nil {
		return m.EgqErppPmfDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEgqErppVlanMbrDiscardCnt() uint64 {
	if m != nil {
		return m.EgqErppVlanMbrDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEpniEpeByteCnt() uint64 {
	if m != nil {
		return m.EpniEpeByteCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEpniEpePktCnt() uint64 {
	if m != nil {
		return m.EpniEpePktCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetEpniEpeDiscardCnt() uint64 {
	if m != nil {
		return m.EpniEpeDiscardCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetNbiTxTotalByteCnt() uint64 {
	if m != nil {
		return m.NbiTxTotalByteCnt
	}
	return 0
}

func (m *FiaEdmDeviceStatsAsicBriefInfo) GetNbiTxTotalPktCnt() uint64 {
	if m != nil {
		return m.NbiTxTotalPktCnt
	}
	return 0
}

type FiaEdmStatsBriefInfo struct {
	Valid                bool                            `protobuf:"varint,50,opt,name=valid,proto3" json:"valid,omitempty"`
	RackNumber           uint32                          `protobuf:"varint,51,opt,name=rack_number,json=rackNumber,proto3" json:"rack_number,omitempty"`
	SlotNumber           uint32                          `protobuf:"varint,52,opt,name=slot_number,json=slotNumber,proto3" json:"slot_number,omitempty"`
	AsicInstance         uint32                          `protobuf:"varint,53,opt,name=asic_instance,json=asicInstance,proto3" json:"asic_instance,omitempty"`
	ChipVersion          uint32                          `protobuf:"varint,54,opt,name=chip_version,json=chipVersion,proto3" json:"chip_version,omitempty"`
	Statistics           *FiaEdmDeviceStatsAsicBriefInfo `protobuf:"bytes,55,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FiaEdmStatsBriefInfo) Reset()         { *m = FiaEdmStatsBriefInfo{} }
func (m *FiaEdmStatsBriefInfo) String() string { return proto.CompactTextString(m) }
func (*FiaEdmStatsBriefInfo) ProtoMessage()    {}
func (*FiaEdmStatsBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b969035e2969373, []int{2}
}

func (m *FiaEdmStatsBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaEdmStatsBriefInfo.Unmarshal(m, b)
}
func (m *FiaEdmStatsBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaEdmStatsBriefInfo.Marshal(b, m, deterministic)
}
func (m *FiaEdmStatsBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaEdmStatsBriefInfo.Merge(m, src)
}
func (m *FiaEdmStatsBriefInfo) XXX_Size() int {
	return xxx_messageInfo_FiaEdmStatsBriefInfo.Size(m)
}
func (m *FiaEdmStatsBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaEdmStatsBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FiaEdmStatsBriefInfo proto.InternalMessageInfo

func (m *FiaEdmStatsBriefInfo) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *FiaEdmStatsBriefInfo) GetRackNumber() uint32 {
	if m != nil {
		return m.RackNumber
	}
	return 0
}

func (m *FiaEdmStatsBriefInfo) GetSlotNumber() uint32 {
	if m != nil {
		return m.SlotNumber
	}
	return 0
}

func (m *FiaEdmStatsBriefInfo) GetAsicInstance() uint32 {
	if m != nil {
		return m.AsicInstance
	}
	return 0
}

func (m *FiaEdmStatsBriefInfo) GetChipVersion() uint32 {
	if m != nil {
		return m.ChipVersion
	}
	return 0
}

func (m *FiaEdmStatsBriefInfo) GetStatistics() *FiaEdmDeviceStatsAsicBriefInfo {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func init() {
	proto.RegisterType((*FiaEdmStatsBriefInfo_KEYS)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_for_npu_ids.asic_statistics_for_npu_id.fia_edm_stats_brief_info_KEYS")
	proto.RegisterType((*FiaEdmDeviceStatsAsicBriefInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_for_npu_ids.asic_statistics_for_npu_id.fia_edm_device_stats_asic_brief_info")
	proto.RegisterType((*FiaEdmStatsBriefInfo)(nil), "cisco_ios_xr_fretta_bcm_dpa_npu_stats_oper.dpa.stats.nodes.node.asic_statistics.asic_statistics_for_npu_ids.asic_statistics_for_npu_id.fia_edm_stats_brief_info")
}

func init() { proto.RegisterFile("fia_edm_stats_brief_info.proto", fileDescriptor_5b969035e2969373) }

var fileDescriptor_5b969035e2969373 = []byte{
	// 1309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x97, 0x6f, 0x6f, 0xd4, 0x36,
	0x18, 0xc0, 0x55, 0xfe, 0x0d, 0xdc, 0x1e, 0x6d, 0xd3, 0x96, 0x06, 0x18, 0x50, 0x0a, 0xac, 0xc0,
	0xe0, 0xa0, 0x77, 0x8c, 0xc1, 0xd8, 0x34, 0x89, 0xf6, 0xd0, 0xaa, 0x71, 0xb4, 0x3a, 0x3a, 0xa4,
	0xbd, 0xb2, 0x72, 0xb1, 0x93, 0xb3, 0x48, 0x1c, 0xc7, 0x71, 0x0a, 0xfd, 0x02, 0x7b, 0x31, 0xed,
	0xa3, 0xec, 0xab, 0xec, 0x3b, 0x4d, 0xfe, 0x1b, 0xe7, 0x7a, 0x45, 0x7b, 0xb9, 0x37, 0x27, 0x9d,
	0xf3, 0xfb, 0xd9, 0x8f, 0x1f, 0x3b, 0x8f, 0x1d, 0x70, 0x33, 0x21, 0x11, 0xc4, 0x28, 0x87, 0x95,
	0x88, 0x44, 0x05, 0xc7, 0x9c, 0xe0, 0x04, 0x12, 0x9a, 0x14, 0x5d, 0xc6, 0x0b, 0x51, 0x04, 0x7f,
	0xcc, 0xc5, 0xa4, 0x8a, 0x0b, 0x48, 0x8a, 0x0a, 0x7e, 0xe6, 0x30, 0xe1, 0x58, 0x88, 0x08, 0x8e,
	0xe3, 0x1c, 0x22, 0x16, 0x41, 0xca, 0x6a, 0x23, 0x16, 0x0c, 0xf3, 0x2e, 0x62, 0x51, 0x57, 0xfd,
	0xed, 0xd2, 0x02, 0x61, 0xfd, 0xdb, 0x8d, 0x2a, 0x12, 0x2b, 0x88, 0x54, 0x82, 0xc4, 0xd5, 0xf4,
	0x7f, 0x98, 0x14, 0x5c, 0xf5, 0x43, 0xd0, 0x97, 0x9e, 0x6d, 0xbe, 0x07, 0x37, 0x4e, 0x0b, 0x15,
	0xfe, 0x3a, 0xf8, 0xfd, 0x7d, 0x70, 0x1d, 0x5c, 0x92, 0xe3, 0x41, 0x1a, 0xe5, 0x38, 0x9c, 0xdb,
	0x98, 0xbb, 0x7f, 0x69, 0x74, 0x51, 0x36, 0xbc, 0x8b, 0x72, 0x1c, 0xac, 0x81, 0x0b, 0xba, 0x9f,
	0xf0, 0xcc, 0xc6, 0xdc, 0xfd, 0xce, 0xe8, 0x3c, 0x65, 0xf5, 0x1e, 0xda, 0xfc, 0x6b, 0x1d, 0xdc,
	0xb5, 0xbd, 0x22, 0x7c, 0x44, 0x62, 0x6c, 0x3a, 0x57, 0x81, 0x34, 0x23, 0x04, 0x4f, 0xc1, 0x1a,
	0x1d, 0x13, 0xc8, 0x3f, 0x43, 0x51, 0x88, 0x28, 0x83, 0xe3, 0x63, 0x81, 0x61, 0x4c, 0x85, 0x1a,
	0xe8, 0xdc, 0x68, 0x99, 0x8e, 0xc9, 0xe8, 0xf3, 0xa1, 0x7c, 0xf4, 0xfa, 0x58, 0xe0, 0x1d, 0x2a,
	0x82, 0x2e, 0x58, 0x6d, 0x19, 0xec, 0xa3, 0x50, 0xc2, 0x19, 0x25, 0x2c, 0x35, 0xc2, 0xc1, 0x47,
	0x21, 0xf9, 0x7b, 0x60, 0x91, 0x70, 0x0c, 0x63, 0x56, 0x3b, 0xf4, 0xac, 0x42, 0x17, 0x08, 0xc7,
	0x3b, 0xac, 0x6e, 0x63, 0x94, 0x24, 0x0e, 0x3b, 0xe7, 0xb0, 0x77, 0x24, 0x31, 0xd8, 0x16, 0x58,
	0x92, 0x58, 0x11, 0xe5, 0xcc, 0x71, 0xe7, 0x15, 0xd7, 0x21, 0x1c, 0xef, 0x47, 0x39, 0x6b, 0xf7,
	0x57, 0x64, 0x0d, 0x77, 0xc1, 0xf5, 0xb7, 0x9f, 0x4d, 0x61, 0x3c, 0x3e, 0x76, 0xd8, 0x57, 0x0e,
	0x1b, 0xc5, 0xc7, 0x06, 0xbb, 0x03, 0x2e, 0x4b, 0x2c, 0x41, 0x02, 0x92, 0x44, 0x51, 0x17, 0x15,
	0x35, 0x4f, 0x38, 0x7e, 0x83, 0xc4, 0x5e, 0x62, 0x21, 0xc4, 0x61, 0x9e, 0xd7, 0x16, 0xba, 0x64,
	0x20, 0xc4, 0x87, 0x79, 0xdd, 0x82, 0x8a, 0x78, 0x6c, 0x21, 0xe0, 0xa0, 0xfd, 0x78, 0xac, 0xa1,
	0xc7, 0x60, 0x85, 0x94, 0x39, 0xc4, 0xb4, 0xac, 0x71, 0x8d, 0x5d, 0x64, 0xf3, 0x3a, 0xc5, 0xa4,
	0xcc, 0x07, 0xfa, 0x89, 0x89, 0xce, 0xe0, 0x08, 0xb7, 0xf1, 0x05, 0x87, 0xef, 0xe2, 0x99, 0x78,
	0x86, 0x05, 0x46, 0x0e, 0xef, 0x78, 0xb8, 0x7a, 0x62, 0xf0, 0x17, 0xe0, 0xaa, 0x09, 0x06, 0x22,
	0x52, 0xc5, 0x11, 0x47, 0x9e, 0x74, 0x59, 0x49, 0x6b, 0x3a, 0xa4, 0x5d, 0xfb, 0xd8, 0x4b, 0x2e,
	0x13, 0x10, 0xa7, 0xa5, 0xe3, 0x17, 0x4d, 0x72, 0x99, 0x18, 0xa4, 0xe5, 0x14, 0x46, 0x1b, 0x6c,
	0xa9, 0xc1, 0xe8, 0x14, 0x26, 0xd7, 0xc0, 0x62, 0xcb, 0x0e, 0x7b, 0x83, 0x84, 0xc1, 0x1e, 0x80,
	0x65, 0x89, 0xc5, 0x49, 0x0a, 0xf1, 0x11, 0xa6, 0x1a, 0x0c, 0x14, 0x78, 0x99, 0x30, 0xb1, 0x93,
	0xa4, 0x03, 0xd9, 0x6c, 0x37, 0x93, 0x41, 0xdd, 0xbe, 0x5f, 0x31, 0x9b, 0x49, 0x91, 0x76, 0xcf,
	0x3f, 0x05, 0x6b, 0x6a, 0xe9, 0x99, 0x80, 0x08, 0x57, 0x31, 0x8c, 0x71, 0x96, 0x29, 0x7a, 0x55,
	0xbf, 0x25, 0x09, 0x12, 0x7b, 0x4c, 0xec, 0xe2, 0x2a, 0xde, 0xc1, 0x59, 0xe6, 0x1b, 0x1c, 0x4f,
	0x19, 0x6b, 0x8d, 0xc1, 0xb1, 0x6f, 0xbc, 0x06, 0x37, 0xa5, 0x21, 0x78, 0x44, 0xab, 0x9c, 0x08,
	0xb9, 0x32, 0x28, 0x12, 0x91, 0x32, 0x2b, 0xa5, 0x5e, 0x51, 0xea, 0xb5, 0x04, 0x89, 0xc3, 0x06,
	0xda, 0x8d, 0x44, 0x24, 0xbb, 0xa8, 0x64, 0x1f, 0x0f, 0x41, 0x90, 0x20, 0x0e, 0xd9, 0xb6, 0x1e,
	0x8f, 0x50, 0xe5, 0xad, 0xeb, 0xc9, 0x27, 0x88, 0x1f, 0x6c, 0x4b, 0x74, 0x8f, 0xfa, 0x6c, 0xaf,
	0xc5, 0x86, 0x0d, 0xdb, 0x3b, 0xc9, 0xf6, 0x5b, 0xec, 0xd5, 0x86, 0xed, 0x37, 0xac, 0x9a, 0x39,
	0xf7, 0x41, 0x5d, 0x28, 0xc2, 0x6b, 0x76, 0xe6, 0xdc, 0xc1, 0xaa, 0x4e, 0x04, 0x8f, 0xc0, 0x4a,
	0x82, 0xec, 0x44, 0x8d, 0xc2, 0xb6, 0xc3, 0xeb, 0x8a, 0x5f, 0x4c, 0x90, 0x9e, 0x9f, 0x12, 0x0e,
	0xb6, 0x67, 0xd3, 0xbd, 0xf0, 0xeb, 0x59, 0x74, 0x6f, 0x36, 0xdd, 0x0f, 0x6f, 0xcc, 0xa2, 0xfb,
	0xb2, 0xb6, 0xb5, 0x68, 0x81, 0x72, 0x35, 0xd3, 0x9b, 0xfa, 0xd5, 0x68, 0xf0, 0x43, 0x94, 0xcb,
	0xb9, 0x3e, 0x03, 0xeb, 0x2d, 0x3e, 0xc7, 0xd5, 0x24, 0x8f, 0x95, 0x72, 0x4b, 0x29, 0x2b, 0x8d,
	0x32, 0x54, 0xcf, 0x4c, 0x05, 0x6d, 0x59, 0x6a, 0x0f, 0x52, 0x11, 0x6e, 0x4c, 0x8f, 0xb2, 0xc7,
	0xc4, 0x09, 0xbe, 0xa8, 0x85, 0x4d, 0xd0, 0xed, 0x36, 0xbf, 0x5f, 0x0b, 0x9d, 0xa1, 0xd9, 0x7c,
	0x2f, 0xdc, 0x9c, 0xc9, 0xf7, 0x4e, 0xe1, 0xfb, 0xe1, 0x9d, 0x99, 0x7c, 0x5f, 0xaf, 0xb0, 0xcf,
	0xdb, 0x34, 0xdd, 0xb5, 0x2b, 0xec, 0x04, 0x93, 0xa7, 0xe7, 0x20, 0x6c, 0x1b, 0x5e, 0xa2, 0xee,
	0x29, 0x69, 0xd5, 0x93, 0x9a, 0x4c, 0x9d, 0x18, 0xc9, 0xa6, 0xea, 0x9b, 0x13, 0x23, 0x99, 0x5c,
	0x6d, 0x01, 0x19, 0xaf, 0x2a, 0x39, 0x88, 0x17, 0x4c, 0xc1, 0x5b, 0xfa, 0x95, 0x4e, 0x50, 0x34,
	0x48, 0xcb, 0x5d, 0x5e, 0x30, 0x6f, 0xe9, 0x24, 0x68, 0x82, 0x71, 0xfc, 0x7d, 0xb7, 0x74, 0x83,
	0xb4, 0xd4, 0xc1, 0x58, 0xeb, 0x1e, 0x58, 0x94, 0x46, 0x52, 0x36, 0xa7, 0xca, 0x03, 0x5d, 0x83,
	0x70, 0x5a, 0xbe, 0x29, 0xed, 0xa9, 0xf2, 0x10, 0x04, 0xaa, 0xe8, 0x95, 0x0c, 0xd6, 0xb1, 0x23,
	0x1f, 0xea, 0xf7, 0x05, 0xa7, 0xe5, 0x41, 0xc9, 0x7e, 0x8b, 0x0d, 0xfb, 0x12, 0x5c, 0xb3, 0xac,
	0x29, 0xaf, 0xbe, 0xf3, 0xad, 0xae, 0xaf, 0xda, 0x31, 0xf5, 0xd5, 0xa9, 0x5d, 0xb0, 0xea, 0x0d,
	0x23, 0x4b, 0x98, 0x2e, 0x14, 0x8f, 0xf4, 0xc2, 0xd9, 0x81, 0x64, 0x15, 0xab, 0xa6, 0xc2, 0xca,
	0x9b, 0x21, 0x1e, 0xfb, 0x61, 0x0d, 0xbf, 0x10, 0x96, 0xe7, 0x74, 0x67, 0x84, 0x35, 0x9c, 0x11,
	0x56, 0xee, 0x87, 0xf5, 0xc4, 0x0f, 0x6b, 0x78, 0x22, 0x2c, 0x3c, 0x69, 0x65, 0xeb, 0xa9, 0x0b,
	0x6b, 0x30, 0x69, 0xb2, 0xd5, 0x07, 0xeb, 0x96, 0xcd, 0x63, 0x38, 0x21, 0xe9, 0xc4, 0x09, 0xdb,
	0x4a, 0x08, 0xb4, 0x30, 0x8c, 0x7f, 0x21, 0xe9, 0xc4, 0x48, 0xdb, 0xe0, 0x8a, 0x27, 0x65, 0xc5,
	0x27, 0xe7, 0xf4, 0xf4, 0x3e, 0xb2, 0xce, 0xdb, 0xe2, 0x53, 0x73, 0x46, 0xaa, 0x3d, 0x34, 0x75,
	0x46, 0xf6, 0xdd, 0x14, 0xda, 0x67, 0xe4, 0x2b, 0x70, 0x7d, 0x3a, 0x2c, 0x9b, 0x35, 0xa9, 0x3d,
	0x53, 0xda, 0x15, 0x3f, 0x34, 0x93, 0x34, 0x2f, 0xd5, 0x5e, 0x78, 0xbe, 0xfb, 0x9d, 0x4b, 0xb5,
	0x0d, 0xd1, 0x53, 0x07, 0x60, 0x43, 0xa9, 0x9c, 0x31, 0x98, 0x45, 0x29, 0x64, 0xbc, 0xa6, 0x84,
	0xa6, 0xad, 0x0e, 0x9e, 0xab, 0x0e, 0x64, 0x7c, 0x03, 0xce, 0xd8, 0xdb, 0x28, 0x3d, 0xd0, 0x90,
	0xd7, 0xcd, 0x73, 0x10, 0xba, 0x6e, 0x58, 0x9e, 0xb4, 0xf4, 0xef, 0xf5, 0xfb, 0x69, 0xf4, 0x83,
	0x3c, 0xf1, 0xbc, 0x9f, 0xc1, 0x0d, 0xe7, 0x1d, 0x65, 0x11, 0x85, 0xf9, 0x98, 0xb7, 0xe4, 0x17,
	0x4a, 0x0e, 0x8d, 0xfc, 0x21, 0x8b, 0xe8, 0x70, 0xcc, 0xbd, 0x0e, 0x1e, 0x80, 0x65, 0xcc, 0x28,
	0x81, 0x98, 0xe1, 0xe6, 0x08, 0x7e, 0x69, 0x56, 0x9e, 0x51, 0x32, 0x60, 0xd8, 0x9e, 0xc1, 0x5b,
	0x60, 0xc9, 0xa1, 0x76, 0x39, 0x7e, 0xd0, 0x6f, 0xb6, 0x21, 0xcd, 0x5a, 0x3c, 0x01, 0xab, 0x0e,
	0xf4, 0x63, 0x79, 0x65, 0xd6, 0x5a, 0xc3, 0x5e, 0x10, 0xe6, 0x0e, 0x2c, 0x4e, 0xdc, 0x81, 0x7f,
	0x74, 0x77, 0xe0, 0xc3, 0x99, 0x77, 0x60, 0x31, 0x7d, 0x07, 0xfe, 0xc9, 0xdd, 0x81, 0x0f, 0xfd,
	0x3b, 0xf0, 0xe6, 0x9f, 0x67, 0x41, 0x78, 0xda, 0x25, 0x3f, 0x58, 0x05, 0xe7, 0x8f, 0xa2, 0x8c,
	0x20, 0xb5, 0x19, 0x2f, 0x8e, 0xf4, 0x9f, 0xe0, 0x16, 0x98, 0xe7, 0x51, 0xfc, 0x11, 0xd2, 0x3a,
	0x1f, 0x63, 0xae, 0x36, 0x5e, 0x67, 0x04, 0x64, 0xd3, 0x3b, 0xd5, 0x22, 0x81, 0x2a, 0x2b, 0x84,
	0x05, 0x9e, 0x69, 0x40, 0x36, 0x19, 0xe0, 0x0e, 0xe8, 0xa8, 0xdb, 0x3e, 0xa1, 0x95, 0x88, 0x68,
	0x8c, 0xd5, 0x4e, 0xea, 0x8c, 0x16, 0x64, 0xe3, 0x9e, 0x69, 0x0b, 0x6e, 0x83, 0x85, 0x78, 0x42,
	0x18, 0x3c, 0xc2, 0xbc, 0x22, 0x05, 0x55, 0x9b, 0xa5, 0x33, 0x9a, 0x97, 0x6d, 0x1f, 0x74, 0x53,
	0xf0, 0xcf, 0x1c, 0x00, 0xcd, 0xa7, 0x8b, 0xda, 0x0f, 0xf3, 0xbd, 0xbf, 0xe7, 0xba, 0xff, 0x8f,
	0xef, 0xa7, 0xee, 0x7f, 0xf9, 0xcc, 0x19, 0x79, 0x13, 0x18, 0x5f, 0x50, 0x1f, 0x80, 0xfd, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x09, 0xaa, 0xd2, 0x76, 0x22, 0x0e, 0x00, 0x00,
}
