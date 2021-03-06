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
// source: pl_macsec_flow_t.proto

package cisco_ios_xr_crypto_macsec_pl_oper_macsec_platform_nodes_node_interfaces_interface_hw_flow_s_hw_flow

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

type PlMacsecFlowT_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FlowId               uint32   `protobuf:"varint,3,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlMacsecFlowT_KEYS) Reset()         { *m = PlMacsecFlowT_KEYS{} }
func (m *PlMacsecFlowT_KEYS) String() string { return proto.CompactTextString(m) }
func (*PlMacsecFlowT_KEYS) ProtoMessage()    {}
func (*PlMacsecFlowT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{0}
}

func (m *PlMacsecFlowT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlMacsecFlowT_KEYS.Unmarshal(m, b)
}
func (m *PlMacsecFlowT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlMacsecFlowT_KEYS.Marshal(b, m, deterministic)
}
func (m *PlMacsecFlowT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlMacsecFlowT_KEYS.Merge(m, src)
}
func (m *PlMacsecFlowT_KEYS) XXX_Size() int {
	return xxx_messageInfo_PlMacsecFlowT_KEYS.Size(m)
}
func (m *PlMacsecFlowT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PlMacsecFlowT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PlMacsecFlowT_KEYS proto.InternalMessageInfo

func (m *PlMacsecFlowT_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *PlMacsecFlowT_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PlMacsecFlowT_KEYS) GetFlowId() uint32 {
	if m != nil {
		return m.FlowId
	}
	return 0
}

type PlMsfpgaFlowInfoT struct {
	FlowId               uint32   `protobuf:"varint,1,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	IsEgress             bool     `protobuf:"varint,3,opt,name=is_egress,json=isEgress,proto3" json:"is_egress,omitempty"`
	InUse                bool     `protobuf:"varint,4,opt,name=in_use,json=inUse,proto3" json:"in_use,omitempty"`
	Action               uint32   `protobuf:"varint,5,opt,name=action,proto3" json:"action,omitempty"`
	Macsa                []uint32 `protobuf:"varint,6,rep,packed,name=macsa,proto3" json:"macsa,omitempty"`
	SmacInuse            bool     `protobuf:"varint,7,opt,name=smac_inuse,json=smacInuse,proto3" json:"smac_inuse,omitempty"`
	Macda                []uint32 `protobuf:"varint,8,rep,packed,name=macda,proto3" json:"macda,omitempty"`
	DmacInuse            bool     `protobuf:"varint,9,opt,name=dmac_inuse,json=dmacInuse,proto3" json:"dmac_inuse,omitempty"`
	Ethertype            uint32   `protobuf:"varint,10,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	OuterVlan            uint32   `protobuf:"varint,11,opt,name=outer_vlan,json=outerVlan,proto3" json:"outer_vlan,omitempty"`
	OuterVlanUp          uint32   `protobuf:"varint,12,opt,name=outer_vlan_up,json=outerVlanUp,proto3" json:"outer_vlan_up,omitempty"`
	OuterVlanTpid        uint32   `protobuf:"varint,13,opt,name=outer_vlan_tpid,json=outerVlanTpid,proto3" json:"outer_vlan_tpid,omitempty"`
	InnerVlan            uint32   `protobuf:"varint,14,opt,name=inner_vlan,json=innerVlan,proto3" json:"inner_vlan,omitempty"`
	InnerVlanUp          uint32   `protobuf:"varint,15,opt,name=inner_vlan_up,json=innerVlanUp,proto3" json:"inner_vlan_up,omitempty"`
	InnerVlanTpid        uint32   `protobuf:"varint,16,opt,name=inner_vlan_tpid,json=innerVlanTpid,proto3" json:"inner_vlan_tpid,omitempty"`
	SourcePort           uint32   `protobuf:"varint,17,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	SourcePortChk        bool     `protobuf:"varint,18,opt,name=source_port_chk,json=sourcePortChk,proto3" json:"source_port_chk,omitempty"`
	SciInuse             bool     `protobuf:"varint,19,opt,name=sci_inuse,json=sciInuse,proto3" json:"sci_inuse,omitempty"`
	Sci                  uint64   `protobuf:"varint,20,opt,name=sci,proto3" json:"sci,omitempty"`
	MatchPri             uint32   `protobuf:"varint,21,opt,name=match_pri,json=matchPri,proto3" json:"match_pri,omitempty"`
	IsCtrlPkt            bool     `protobuf:"varint,22,opt,name=is_ctrl_pkt,json=isCtrlPkt,proto3" json:"is_ctrl_pkt,omitempty"`
	CtrlCheck            bool     `protobuf:"varint,23,opt,name=ctrl_check,json=ctrlCheck,proto3" json:"ctrl_check,omitempty"`
	MatchUntagged        bool     `protobuf:"varint,24,opt,name=match_untagged,json=matchUntagged,proto3" json:"match_untagged,omitempty"`
	MatchTagged          bool     `protobuf:"varint,25,opt,name=match_tagged,json=matchTagged,proto3" json:"match_tagged,omitempty"`
	MatchBadTag          bool     `protobuf:"varint,26,opt,name=match_bad_tag,json=matchBadTag,proto3" json:"match_bad_tag,omitempty"`
	MatchKayTag          bool     `protobuf:"varint,27,opt,name=match_kay_tag,json=matchKayTag,proto3" json:"match_kay_tag,omitempty"`
	TciV                 uint32   `protobuf:"varint,28,opt,name=tci_v,json=tciV,proto3" json:"tci_v,omitempty"`
	TciEXr               uint32   `protobuf:"varint,29,opt,name=tci_e_xr,json=tciEXr,proto3" json:"tci_e_xr,omitempty"`
	TciSc                uint32   `protobuf:"varint,30,opt,name=tci_sc,json=tciSc,proto3" json:"tci_sc,omitempty"`
	TciScb               uint32   `protobuf:"varint,31,opt,name=tci_scb,json=tciScb,proto3" json:"tci_scb,omitempty"`
	Tci                  uint32   `protobuf:"varint,32,opt,name=tci,proto3" json:"tci,omitempty"`
	TciC                 uint32   `protobuf:"varint,33,opt,name=tci_c,json=tciC,proto3" json:"tci_c,omitempty"`
	TciAn                uint32   `protobuf:"varint,34,opt,name=tci_an,json=tciAn,proto3" json:"tci_an,omitempty"`
	TciAnChk             bool     `protobuf:"varint,35,opt,name=tci_an_chk,json=tciAnChk,proto3" json:"tci_an_chk,omitempty"`
	TciChk               bool     `protobuf:"varint,36,opt,name=tci_chk,json=tciChk,proto3" json:"tci_chk,omitempty"`
	Sai                  uint32   `protobuf:"varint,37,opt,name=sai,proto3" json:"sai,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlMsfpgaFlowInfoT) Reset()         { *m = PlMsfpgaFlowInfoT{} }
func (m *PlMsfpgaFlowInfoT) String() string { return proto.CompactTextString(m) }
func (*PlMsfpgaFlowInfoT) ProtoMessage()    {}
func (*PlMsfpgaFlowInfoT) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{1}
}

func (m *PlMsfpgaFlowInfoT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlMsfpgaFlowInfoT.Unmarshal(m, b)
}
func (m *PlMsfpgaFlowInfoT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlMsfpgaFlowInfoT.Marshal(b, m, deterministic)
}
func (m *PlMsfpgaFlowInfoT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlMsfpgaFlowInfoT.Merge(m, src)
}
func (m *PlMsfpgaFlowInfoT) XXX_Size() int {
	return xxx_messageInfo_PlMsfpgaFlowInfoT.Size(m)
}
func (m *PlMsfpgaFlowInfoT) XXX_DiscardUnknown() {
	xxx_messageInfo_PlMsfpgaFlowInfoT.DiscardUnknown(m)
}

var xxx_messageInfo_PlMsfpgaFlowInfoT proto.InternalMessageInfo

func (m *PlMsfpgaFlowInfoT) GetFlowId() uint32 {
	if m != nil {
		return m.FlowId
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetIsEgress() bool {
	if m != nil {
		return m.IsEgress
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetMacsa() []uint32 {
	if m != nil {
		return m.Macsa
	}
	return nil
}

func (m *PlMsfpgaFlowInfoT) GetSmacInuse() bool {
	if m != nil {
		return m.SmacInuse
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetMacda() []uint32 {
	if m != nil {
		return m.Macda
	}
	return nil
}

func (m *PlMsfpgaFlowInfoT) GetDmacInuse() bool {
	if m != nil {
		return m.DmacInuse
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetEthertype() uint32 {
	if m != nil {
		return m.Ethertype
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetOuterVlan() uint32 {
	if m != nil {
		return m.OuterVlan
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetOuterVlanUp() uint32 {
	if m != nil {
		return m.OuterVlanUp
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetOuterVlanTpid() uint32 {
	if m != nil {
		return m.OuterVlanTpid
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetInnerVlan() uint32 {
	if m != nil {
		return m.InnerVlan
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetInnerVlanUp() uint32 {
	if m != nil {
		return m.InnerVlanUp
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetInnerVlanTpid() uint32 {
	if m != nil {
		return m.InnerVlanTpid
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetSourcePort() uint32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetSourcePortChk() bool {
	if m != nil {
		return m.SourcePortChk
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetSciInuse() bool {
	if m != nil {
		return m.SciInuse
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetSci() uint64 {
	if m != nil {
		return m.Sci
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetMatchPri() uint32 {
	if m != nil {
		return m.MatchPri
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetIsCtrlPkt() bool {
	if m != nil {
		return m.IsCtrlPkt
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetCtrlCheck() bool {
	if m != nil {
		return m.CtrlCheck
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetMatchUntagged() bool {
	if m != nil {
		return m.MatchUntagged
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetMatchTagged() bool {
	if m != nil {
		return m.MatchTagged
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetMatchBadTag() bool {
	if m != nil {
		return m.MatchBadTag
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetMatchKayTag() bool {
	if m != nil {
		return m.MatchKayTag
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetTciV() uint32 {
	if m != nil {
		return m.TciV
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciEXr() uint32 {
	if m != nil {
		return m.TciEXr
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciSc() uint32 {
	if m != nil {
		return m.TciSc
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciScb() uint32 {
	if m != nil {
		return m.TciScb
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTci() uint32 {
	if m != nil {
		return m.Tci
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciC() uint32 {
	if m != nil {
		return m.TciC
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciAn() uint32 {
	if m != nil {
		return m.TciAn
	}
	return 0
}

func (m *PlMsfpgaFlowInfoT) GetTciAnChk() bool {
	if m != nil {
		return m.TciAnChk
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetTciChk() bool {
	if m != nil {
		return m.TciChk
	}
	return false
}

func (m *PlMsfpgaFlowInfoT) GetSai() uint32 {
	if m != nil {
		return m.Sai
	}
	return 0
}

type PlMsfpgaFlowT struct {
	TxFlow               *PlMsfpgaFlowInfoT `protobuf:"bytes,1,opt,name=tx_flow,json=txFlow,proto3" json:"tx_flow,omitempty"`
	RxFlow               *PlMsfpgaFlowInfoT `protobuf:"bytes,2,opt,name=rx_flow,json=rxFlow,proto3" json:"rx_flow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PlMsfpgaFlowT) Reset()         { *m = PlMsfpgaFlowT{} }
func (m *PlMsfpgaFlowT) String() string { return proto.CompactTextString(m) }
func (*PlMsfpgaFlowT) ProtoMessage()    {}
func (*PlMsfpgaFlowT) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{2}
}

func (m *PlMsfpgaFlowT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlMsfpgaFlowT.Unmarshal(m, b)
}
func (m *PlMsfpgaFlowT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlMsfpgaFlowT.Marshal(b, m, deterministic)
}
func (m *PlMsfpgaFlowT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlMsfpgaFlowT.Merge(m, src)
}
func (m *PlMsfpgaFlowT) XXX_Size() int {
	return xxx_messageInfo_PlMsfpgaFlowT.Size(m)
}
func (m *PlMsfpgaFlowT) XXX_DiscardUnknown() {
	xxx_messageInfo_PlMsfpgaFlowT.DiscardUnknown(m)
}

var xxx_messageInfo_PlMsfpgaFlowT proto.InternalMessageInfo

func (m *PlMsfpgaFlowT) GetTxFlow() *PlMsfpgaFlowInfoT {
	if m != nil {
		return m.TxFlow
	}
	return nil
}

func (m *PlMsfpgaFlowT) GetRxFlow() *PlMsfpgaFlowInfoT {
	if m != nil {
		return m.RxFlow
	}
	return nil
}

type PlEs200FlowInfoT struct {
	FlowNo               uint32   `protobuf:"varint,1,opt,name=flow_no,json=flowNo,proto3" json:"flow_no,omitempty"`
	IsFlowEnabled        bool     `protobuf:"varint,2,opt,name=is_flow_enabled,json=isFlowEnabled,proto3" json:"is_flow_enabled,omitempty"`
	Macda                []uint32 `protobuf:"varint,3,rep,packed,name=macda,proto3" json:"macda,omitempty"`
	Ethertype            uint32   `protobuf:"varint,4,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	OuterVlanId          uint32   `protobuf:"varint,5,opt,name=outer_vlan_id,json=outerVlanId,proto3" json:"outer_vlan_id,omitempty"`
	OuterVlanUserPri     uint32   `protobuf:"varint,6,opt,name=outer_vlan_user_pri,json=outerVlanUserPri,proto3" json:"outer_vlan_user_pri,omitempty"`
	InnerVlanId          uint32   `protobuf:"varint,7,opt,name=inner_vlan_id,json=innerVlanId,proto3" json:"inner_vlan_id,omitempty"`
	InnerVlanUserPri     uint32   `protobuf:"varint,8,opt,name=inner_vlan_user_pri,json=innerVlanUserPri,proto3" json:"inner_vlan_user_pri,omitempty"`
	Psci                 uint64   `protobuf:"varint,9,opt,name=psci,proto3" json:"psci,omitempty"`
	MatchPriority        uint32   `protobuf:"varint,10,opt,name=match_priority,json=matchPriority,proto3" json:"match_priority,omitempty"`
	TciV                 uint32   `protobuf:"varint,11,opt,name=tci_v,json=tciV,proto3" json:"tci_v,omitempty"`
	TciEXr               uint32   `protobuf:"varint,12,opt,name=tci_e_xr,json=tciEXr,proto3" json:"tci_e_xr,omitempty"`
	TciSc                uint32   `protobuf:"varint,13,opt,name=tci_sc,json=tciSc,proto3" json:"tci_sc,omitempty"`
	TciScb               uint32   `protobuf:"varint,14,opt,name=tci_scb,json=tciScb,proto3" json:"tci_scb,omitempty"`
	Tci                  uint32   `protobuf:"varint,15,opt,name=tci,proto3" json:"tci,omitempty"`
	TciC                 uint32   `protobuf:"varint,16,opt,name=tci_c,json=tciC,proto3" json:"tci_c,omitempty"`
	TciChk               bool     `protobuf:"varint,17,opt,name=tci_chk,json=tciChk,proto3" json:"tci_chk,omitempty"`
	PktType              string   `protobuf:"bytes,18,opt,name=pkt_type,json=pktType,proto3" json:"pkt_type,omitempty"`
	TagNum               string   `protobuf:"bytes,19,opt,name=tag_num,json=tagNum,proto3" json:"tag_num,omitempty"`
	InnerVlanDei         bool     `protobuf:"varint,20,opt,name=inner_vlan_dei,json=innerVlanDei,proto3" json:"inner_vlan_dei,omitempty"`
	OuterVlanDei         bool     `protobuf:"varint,21,opt,name=outer_vlan_dei,json=outerVlanDei,proto3" json:"outer_vlan_dei,omitempty"`
	PbbSid               uint32   `protobuf:"varint,22,opt,name=pbb_sid,json=pbbSid,proto3" json:"pbb_sid,omitempty"`
	PbbBvid              uint32   `protobuf:"varint,23,opt,name=pbb_bvid,json=pbbBvid,proto3" json:"pbb_bvid,omitempty"`
	PbbPcp               uint32   `protobuf:"varint,24,opt,name=pbb_pcp,json=pbbPcp,proto3" json:"pbb_pcp,omitempty"`
	PbbDei               uint32   `protobuf:"varint,25,opt,name=pbb_dei,json=pbbDei,proto3" json:"pbb_dei,omitempty"`
	Mpls1Label           uint32   `protobuf:"varint,26,opt,name=mpls1_label,json=mpls1Label,proto3" json:"mpls1_label,omitempty"`
	Mpls1Exp             uint32   `protobuf:"varint,27,opt,name=mpls1_exp,json=mpls1Exp,proto3" json:"mpls1_exp,omitempty"`
	Mpls1Bos             uint32   `protobuf:"varint,28,opt,name=mpls1_bos,json=mpls1Bos,proto3" json:"mpls1_bos,omitempty"`
	Mpls2Label           uint32   `protobuf:"varint,29,opt,name=mpls2_label,json=mpls2Label,proto3" json:"mpls2_label,omitempty"`
	Mpls2Exp             uint32   `protobuf:"varint,30,opt,name=mpls2_exp,json=mpls2Exp,proto3" json:"mpls2_exp,omitempty"`
	Mpls2Bos             uint32   `protobuf:"varint,31,opt,name=mpls2_bos,json=mpls2Bos,proto3" json:"mpls2_bos,omitempty"`
	PlainBits            uint64   `protobuf:"varint,32,opt,name=plain_bits,json=plainBits,proto3" json:"plain_bits,omitempty"`
	PlainBitsSize        uint32   `protobuf:"varint,33,opt,name=plain_bits_size,json=plainBitsSize,proto3" json:"plain_bits_size,omitempty"`
	ForceCtrl            bool     `protobuf:"varint,34,opt,name=force_ctrl,json=forceCtrl,proto3" json:"force_ctrl,omitempty"`
	Drop                 bool     `protobuf:"varint,35,opt,name=drop,proto3" json:"drop,omitempty"`
	MaskDa               uint64   `protobuf:"varint,36,opt,name=mask_da,json=maskDa,proto3" json:"mask_da,omitempty"`
	MaskEthertype        uint32   `protobuf:"varint,37,opt,name=mask_ethertype,json=maskEthertype,proto3" json:"mask_ethertype,omitempty"`
	MaskPlainBits        uint64   `protobuf:"varint,38,opt,name=mask_plain_bits,json=maskPlainBits,proto3" json:"mask_plain_bits,omitempty"`
	FlowHits             uint64   `protobuf:"varint,39,opt,name=flow_hits,json=flowHits,proto3" json:"flow_hits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlEs200FlowInfoT) Reset()         { *m = PlEs200FlowInfoT{} }
func (m *PlEs200FlowInfoT) String() string { return proto.CompactTextString(m) }
func (*PlEs200FlowInfoT) ProtoMessage()    {}
func (*PlEs200FlowInfoT) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{3}
}

func (m *PlEs200FlowInfoT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlEs200FlowInfoT.Unmarshal(m, b)
}
func (m *PlEs200FlowInfoT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlEs200FlowInfoT.Marshal(b, m, deterministic)
}
func (m *PlEs200FlowInfoT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlEs200FlowInfoT.Merge(m, src)
}
func (m *PlEs200FlowInfoT) XXX_Size() int {
	return xxx_messageInfo_PlEs200FlowInfoT.Size(m)
}
func (m *PlEs200FlowInfoT) XXX_DiscardUnknown() {
	xxx_messageInfo_PlEs200FlowInfoT.DiscardUnknown(m)
}

var xxx_messageInfo_PlEs200FlowInfoT proto.InternalMessageInfo

func (m *PlEs200FlowInfoT) GetFlowNo() uint32 {
	if m != nil {
		return m.FlowNo
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetIsFlowEnabled() bool {
	if m != nil {
		return m.IsFlowEnabled
	}
	return false
}

func (m *PlEs200FlowInfoT) GetMacda() []uint32 {
	if m != nil {
		return m.Macda
	}
	return nil
}

func (m *PlEs200FlowInfoT) GetEthertype() uint32 {
	if m != nil {
		return m.Ethertype
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetOuterVlanId() uint32 {
	if m != nil {
		return m.OuterVlanId
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetOuterVlanUserPri() uint32 {
	if m != nil {
		return m.OuterVlanUserPri
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetInnerVlanId() uint32 {
	if m != nil {
		return m.InnerVlanId
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetInnerVlanUserPri() uint32 {
	if m != nil {
		return m.InnerVlanUserPri
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPsci() uint64 {
	if m != nil {
		return m.Psci
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMatchPriority() uint32 {
	if m != nil {
		return m.MatchPriority
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciV() uint32 {
	if m != nil {
		return m.TciV
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciEXr() uint32 {
	if m != nil {
		return m.TciEXr
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciSc() uint32 {
	if m != nil {
		return m.TciSc
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciScb() uint32 {
	if m != nil {
		return m.TciScb
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTci() uint32 {
	if m != nil {
		return m.Tci
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciC() uint32 {
	if m != nil {
		return m.TciC
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetTciChk() bool {
	if m != nil {
		return m.TciChk
	}
	return false
}

func (m *PlEs200FlowInfoT) GetPktType() string {
	if m != nil {
		return m.PktType
	}
	return ""
}

func (m *PlEs200FlowInfoT) GetTagNum() string {
	if m != nil {
		return m.TagNum
	}
	return ""
}

func (m *PlEs200FlowInfoT) GetInnerVlanDei() bool {
	if m != nil {
		return m.InnerVlanDei
	}
	return false
}

func (m *PlEs200FlowInfoT) GetOuterVlanDei() bool {
	if m != nil {
		return m.OuterVlanDei
	}
	return false
}

func (m *PlEs200FlowInfoT) GetPbbSid() uint32 {
	if m != nil {
		return m.PbbSid
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPbbBvid() uint32 {
	if m != nil {
		return m.PbbBvid
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPbbPcp() uint32 {
	if m != nil {
		return m.PbbPcp
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPbbDei() uint32 {
	if m != nil {
		return m.PbbDei
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls1Label() uint32 {
	if m != nil {
		return m.Mpls1Label
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls1Exp() uint32 {
	if m != nil {
		return m.Mpls1Exp
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls1Bos() uint32 {
	if m != nil {
		return m.Mpls1Bos
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls2Label() uint32 {
	if m != nil {
		return m.Mpls2Label
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls2Exp() uint32 {
	if m != nil {
		return m.Mpls2Exp
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMpls2Bos() uint32 {
	if m != nil {
		return m.Mpls2Bos
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPlainBits() uint64 {
	if m != nil {
		return m.PlainBits
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetPlainBitsSize() uint32 {
	if m != nil {
		return m.PlainBitsSize
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetForceCtrl() bool {
	if m != nil {
		return m.ForceCtrl
	}
	return false
}

func (m *PlEs200FlowInfoT) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *PlEs200FlowInfoT) GetMaskDa() uint64 {
	if m != nil {
		return m.MaskDa
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMaskEthertype() uint32 {
	if m != nil {
		return m.MaskEthertype
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetMaskPlainBits() uint64 {
	if m != nil {
		return m.MaskPlainBits
	}
	return 0
}

func (m *PlEs200FlowInfoT) GetFlowHits() uint64 {
	if m != nil {
		return m.FlowHits
	}
	return 0
}

type PlEs200FlowT struct {
	TxFlow               *PlEs200FlowInfoT `protobuf:"bytes,1,opt,name=tx_flow,json=txFlow,proto3" json:"tx_flow,omitempty"`
	RxFlow               *PlEs200FlowInfoT `protobuf:"bytes,2,opt,name=rx_flow,json=rxFlow,proto3" json:"rx_flow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PlEs200FlowT) Reset()         { *m = PlEs200FlowT{} }
func (m *PlEs200FlowT) String() string { return proto.CompactTextString(m) }
func (*PlEs200FlowT) ProtoMessage()    {}
func (*PlEs200FlowT) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{4}
}

func (m *PlEs200FlowT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlEs200FlowT.Unmarshal(m, b)
}
func (m *PlEs200FlowT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlEs200FlowT.Marshal(b, m, deterministic)
}
func (m *PlEs200FlowT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlEs200FlowT.Merge(m, src)
}
func (m *PlEs200FlowT) XXX_Size() int {
	return xxx_messageInfo_PlEs200FlowT.Size(m)
}
func (m *PlEs200FlowT) XXX_DiscardUnknown() {
	xxx_messageInfo_PlEs200FlowT.DiscardUnknown(m)
}

var xxx_messageInfo_PlEs200FlowT proto.InternalMessageInfo

func (m *PlEs200FlowT) GetTxFlow() *PlEs200FlowInfoT {
	if m != nil {
		return m.TxFlow
	}
	return nil
}

func (m *PlEs200FlowT) GetRxFlow() *PlEs200FlowInfoT {
	if m != nil {
		return m.RxFlow
	}
	return nil
}

type PlMacsecFlowExt struct {
	Type                 string         `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	MsfpgaFlow           *PlMsfpgaFlowT `protobuf:"bytes,2,opt,name=msfpga_flow,json=msfpgaFlow,proto3" json:"msfpga_flow,omitempty"`
	Es200Flow            *PlEs200FlowT  `protobuf:"bytes,3,opt,name=es200_flow,json=es200Flow,proto3" json:"es200_flow,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PlMacsecFlowExt) Reset()         { *m = PlMacsecFlowExt{} }
func (m *PlMacsecFlowExt) String() string { return proto.CompactTextString(m) }
func (*PlMacsecFlowExt) ProtoMessage()    {}
func (*PlMacsecFlowExt) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{5}
}

func (m *PlMacsecFlowExt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlMacsecFlowExt.Unmarshal(m, b)
}
func (m *PlMacsecFlowExt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlMacsecFlowExt.Marshal(b, m, deterministic)
}
func (m *PlMacsecFlowExt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlMacsecFlowExt.Merge(m, src)
}
func (m *PlMacsecFlowExt) XXX_Size() int {
	return xxx_messageInfo_PlMacsecFlowExt.Size(m)
}
func (m *PlMacsecFlowExt) XXX_DiscardUnknown() {
	xxx_messageInfo_PlMacsecFlowExt.DiscardUnknown(m)
}

var xxx_messageInfo_PlMacsecFlowExt proto.InternalMessageInfo

func (m *PlMacsecFlowExt) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PlMacsecFlowExt) GetMsfpgaFlow() *PlMsfpgaFlowT {
	if m != nil {
		return m.MsfpgaFlow
	}
	return nil
}

func (m *PlMacsecFlowExt) GetEs200Flow() *PlEs200FlowT {
	if m != nil {
		return m.Es200Flow
	}
	return nil
}

type PlMacsecFlowT struct {
	Ext                  *PlMacsecFlowExt `protobuf:"bytes,50,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PlMacsecFlowT) Reset()         { *m = PlMacsecFlowT{} }
func (m *PlMacsecFlowT) String() string { return proto.CompactTextString(m) }
func (*PlMacsecFlowT) ProtoMessage()    {}
func (*PlMacsecFlowT) Descriptor() ([]byte, []int) {
	return fileDescriptor_da9007011b313edb, []int{6}
}

func (m *PlMacsecFlowT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlMacsecFlowT.Unmarshal(m, b)
}
func (m *PlMacsecFlowT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlMacsecFlowT.Marshal(b, m, deterministic)
}
func (m *PlMacsecFlowT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlMacsecFlowT.Merge(m, src)
}
func (m *PlMacsecFlowT) XXX_Size() int {
	return xxx_messageInfo_PlMacsecFlowT.Size(m)
}
func (m *PlMacsecFlowT) XXX_DiscardUnknown() {
	xxx_messageInfo_PlMacsecFlowT.DiscardUnknown(m)
}

var xxx_messageInfo_PlMacsecFlowT proto.InternalMessageInfo

func (m *PlMacsecFlowT) GetExt() *PlMacsecFlowExt {
	if m != nil {
		return m.Ext
	}
	return nil
}

func init() {
	proto.RegisterType((*PlMacsecFlowT_KEYS)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_macsec_flow_t_KEYS")
	proto.RegisterType((*PlMsfpgaFlowInfoT)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_msfpga_flow_info_t")
	proto.RegisterType((*PlMsfpgaFlowT)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_msfpga_flow_t")
	proto.RegisterType((*PlEs200FlowInfoT)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_es200_flow_info_t")
	proto.RegisterType((*PlEs200FlowT)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_es200_flow_t")
	proto.RegisterType((*PlMacsecFlowExt)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_macsec_flow_ext")
	proto.RegisterType((*PlMacsecFlowT)(nil), "cisco_ios_xr_crypto_macsec_pl_oper.macsec_platform.nodes.node.interfaces.interface.hw_flow_s.hw_flow.pl_macsec_flow_t")
}

func init() { proto.RegisterFile("pl_macsec_flow_t.proto", fileDescriptor_da9007011b313edb) }

var fileDescriptor_da9007011b313edb = []byte{
	// 1312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x86, 0x13, 0xd7, 0x91, 0xe9, 0x28, 0x49, 0xd9, 0xa6, 0x65, 0xd7, 0xa6, 0x49, 0xbd, 0xb6,
	0xcb, 0xcd, 0x8c, 0xce, 0x7b, 0x82, 0x25, 0xcd, 0xb0, 0xa2, 0x43, 0x11, 0x28, 0x69, 0xb1, 0x5d,
	0x11, 0x14, 0xc5, 0xd8, 0x84, 0x65, 0x89, 0x10, 0xe9, 0xd4, 0xe9, 0xde, 0x60, 0xb7, 0x1d, 0xb0,
	0xcb, 0xdd, 0xed, 0x29, 0x36, 0x60, 0xaf, 0xb1, 0xb7, 0x19, 0xce, 0xa1, 0xfe, 0x92, 0xa6, 0xb7,
	0xed, 0x6e, 0x0c, 0xf2, 0x3b, 0xdf, 0xf9, 0x11, 0x8f, 0xce, 0x47, 0x99, 0xdc, 0x31, 0x29, 0x9f,
	0x0b, 0x69, 0x95, 0xe4, 0x67, 0x69, 0xfe, 0x96, 0xbb, 0x91, 0x29, 0x72, 0x97, 0xd3, 0x44, 0x6a,
	0x2b, 0x73, 0xae, 0x73, 0xcb, 0x97, 0x05, 0x97, 0xc5, 0x85, 0x71, 0x79, 0x45, 0x34, 0x29, 0xcf,
	0x8d, 0x2a, 0x46, 0xf5, 0x56, 0xb8, 0xb3, 0xbc, 0x98, 0x8f, 0xb2, 0x3c, 0x51, 0x16, 0x7f, 0x47,
	0x3a, 0x73, 0xaa, 0x38, 0x13, 0x52, 0xd9, 0x66, 0x39, 0x9a, 0xbe, 0xf5, 0x39, 0x6c, 0xb5, 0x1a,
	0x0a, 0xb2, 0x7d, 0x35, 0x3f, 0x7f, 0x79, 0xf4, 0xf3, 0x09, 0xbd, 0x4f, 0xfa, 0x10, 0x87, 0x67,
	0x62, 0xae, 0x58, 0x67, 0xaf, 0xb3, 0xdf, 0x8f, 0x02, 0x00, 0x5e, 0x89, 0xb9, 0xa2, 0x94, 0x74,
	0x11, 0x5f, 0x41, 0x1c, 0xd7, 0xf4, 0x2e, 0x59, 0x43, 0x7f, 0x9d, 0xb0, 0xd5, 0xbd, 0xce, 0x7e,
	0x18, 0xf5, 0x60, 0xfb, 0x22, 0x19, 0xfe, 0x15, 0xf8, 0x1c, 0xf6, 0xcc, 0x4c, 0x84, 0xcf, 0xa1,
	0xb3, 0xb3, 0x9c, 0xbb, 0xb6, 0x4b, 0xa7, 0xed, 0x42, 0x6f, 0x93, 0x1b, 0xe7, 0x22, 0xd5, 0x09,
	0x26, 0x08, 0x22, 0xbf, 0x81, 0x92, 0xb4, 0xe5, 0x6a, 0x52, 0x28, 0x6b, 0x31, 0x47, 0x10, 0x05,
	0xda, 0x1e, 0xe1, 0x9e, 0x6e, 0x93, 0x9e, 0xce, 0xf8, 0xc2, 0x2a, 0xd6, 0xf5, 0x3e, 0x3a, 0x7b,
	0x6d, 0x15, 0xbd, 0x43, 0x7a, 0x42, 0x3a, 0x9d, 0x67, 0xec, 0x86, 0xcf, 0xe0, 0x77, 0x90, 0x01,
	0x1e, 0x5a, 0xb0, 0xde, 0xde, 0xea, 0x7e, 0x18, 0xf9, 0x0d, 0xdd, 0x21, 0xc4, 0xce, 0x85, 0xe4,
	0x3a, 0x83, 0x40, 0x6b, 0x18, 0xa8, 0x0f, 0xc8, 0x0b, 0x00, 0x4a, 0xa7, 0x44, 0xb0, 0xa0, 0x76,
	0x4a, 0xd0, 0x29, 0x69, 0x9c, 0xfa, 0xde, 0x29, 0xa9, 0x9d, 0x1e, 0x90, 0xbe, 0x72, 0x53, 0x55,
	0xb8, 0x0b, 0xa3, 0x18, 0xc1, 0x22, 0x1a, 0x00, 0x9c, 0xf3, 0x85, 0x53, 0x05, 0x3f, 0x4f, 0x45,
	0xc6, 0x06, 0xde, 0x8c, 0xc8, 0x9b, 0x54, 0x64, 0x74, 0x48, 0xc2, 0xc6, 0xcc, 0x17, 0x86, 0xad,
	0x23, 0x63, 0x50, 0x33, 0x5e, 0x1b, 0xfa, 0x94, 0x6c, 0xb6, 0x38, 0xce, 0xe8, 0x84, 0x85, 0xc8,
	0x0a, 0x6b, 0xd6, 0xa9, 0xd1, 0x09, 0xa4, 0xd2, 0x59, 0x56, 0xa5, 0xda, 0xf0, 0xa9, 0x10, 0xa9,
	0x52, 0x35, 0x66, 0x48, 0xb5, 0xe9, 0x53, 0xd5, 0x0c, 0x9f, 0xaa, 0xc5, 0xc1, 0x54, 0x5b, 0x3e,
	0x55, 0xcd, 0xc2, 0x54, 0xbb, 0x64, 0x60, 0xf3, 0x45, 0x21, 0x15, 0x37, 0x79, 0xe1, 0xd8, 0x4d,
	0xe4, 0x10, 0x0f, 0x1d, 0xe7, 0x85, 0x83, 0x40, 0x2d, 0x02, 0x97, 0xd3, 0x19, 0xa3, 0x78, 0x70,
	0x61, 0x43, 0x3a, 0x9c, 0xce, 0xa0, 0xe5, 0x56, 0xea, 0xf2, 0x68, 0x6f, 0xf9, 0x96, 0x5b, 0xa9,
	0xfd, 0xc9, 0x6e, 0x91, 0x55, 0x2b, 0x35, 0xbb, 0xbd, 0xd7, 0xd9, 0xef, 0x46, 0xb0, 0x04, 0xfa,
	0x5c, 0x38, 0x39, 0xe5, 0xa6, 0xd0, 0x6c, 0x1b, 0xb3, 0x06, 0x08, 0x1c, 0x17, 0x9a, 0x3e, 0x24,
	0x03, 0x6d, 0xb9, 0x74, 0x45, 0xca, 0xcd, 0xcc, 0xb1, 0x3b, 0xbe, 0x51, 0xda, 0x1e, 0xba, 0x22,
	0x3d, 0x9e, 0x39, 0x38, 0x1f, 0x34, 0xca, 0xa9, 0x92, 0x33, 0x76, 0xd7, 0x9b, 0x01, 0x39, 0x04,
	0x80, 0x3e, 0x21, 0x1b, 0x3e, 0xf6, 0x22, 0x73, 0x62, 0x32, 0x51, 0x09, 0x63, 0xbe, 0x62, 0x44,
	0x5f, 0x97, 0x20, 0x7d, 0x44, 0xd6, 0x3d, 0xad, 0x24, 0xdd, 0x43, 0xd2, 0x00, 0xb1, 0x53, 0x4f,
	0x19, 0x12, 0xef, 0xc3, 0x63, 0x91, 0x00, 0x8d, 0x7d, 0xd1, 0xe2, 0x1c, 0x88, 0xe4, 0x54, 0x4c,
	0x1a, 0xce, 0x4c, 0x5c, 0x20, 0xe7, 0x7e, 0x8b, 0xf3, 0x52, 0x5c, 0x00, 0xe7, 0x16, 0xb9, 0xe1,
	0xa4, 0xe6, 0xe7, 0xec, 0x01, 0x3e, 0x69, 0xd7, 0x49, 0xfd, 0x86, 0x32, 0x12, 0x00, 0xa8, 0xf8,
	0xb2, 0x60, 0x3b, 0xfe, 0x95, 0x77, 0x52, 0x1f, 0xfd, 0x54, 0xc0, 0x84, 0x80, 0xc5, 0x4a, 0xf6,
	0x10, 0x71, 0x70, 0x3e, 0x91, 0x30, 0x84, 0x1e, 0x8e, 0xd9, 0x6e, 0xcd, 0x3f, 0x91, 0x31, 0x1c,
	0xaf, 0x93, 0x9a, 0xed, 0x21, 0x08, 0xcb, 0x2a, 0xa1, 0x64, 0x8f, 0xea, 0x84, 0x87, 0x55, 0x58,
	0x91, 0xb1, 0x61, 0x1d, 0xf6, 0xbb, 0x8c, 0x3e, 0x20, 0xc4, 0xc3, 0xd8, 0xdc, 0x2f, 0x7d, 0xeb,
	0xd0, 0x04, 0x7d, 0x2d, 0x93, 0x82, 0xe9, 0x31, 0x9a, 0x20, 0x06, 0x18, 0xa0, 0xa7, 0x42, 0xb3,
	0x27, 0x3e, 0xa9, 0x15, 0x7a, 0xf8, 0xcf, 0x0a, 0xd9, 0xba, 0x22, 0x1f, 0x8e, 0xfe, 0xd6, 0x21,
	0x6b, 0x6e, 0x89, 0x3b, 0x94, 0x8e, 0xc1, 0xf8, 0x97, 0xd1, 0xa7, 0xd0, 0xcb, 0xd1, 0xb5, 0x42,
	0x16, 0xf5, 0xdc, 0xf2, 0xfb, 0x34, 0x7f, 0x8b, 0x65, 0x15, 0x65, 0x59, 0x2b, 0xff, 0x83, 0xb2,
	0x0a, 0x2c, 0x6b, 0xf8, 0x67, 0x9f, 0xdc, 0x36, 0x29, 0x57, 0x76, 0xfc, 0xec, 0xd9, 0xb5, 0x02,
	0x9c, 0xe5, 0x6d, 0x01, 0x7e, 0x95, 0xe3, 0xa0, 0x5b, 0x4f, 0x55, 0x99, 0x88, 0x53, 0x55, 0x49,
	0x71, 0xa8, 0x2d, 0x84, 0x3c, 0xf2, 0x60, 0xa3, 0x88, 0xab, 0x6d, 0x45, 0xbc, 0x24, 0x79, 0xdd,
	0xab, 0x92, 0x77, 0x59, 0xd3, 0x74, 0x52, 0x2a, 0x73, 0xa3, 0x69, 0x2f, 0x12, 0xfa, 0x35, 0xb9,
	0xd5, 0xd6, 0x3d, 0xab, 0x0a, 0x1c, 0xe9, 0x1e, 0x32, 0xb7, 0x1a, 0xf5, 0xb3, 0xaa, 0x80, 0xd1,
	0xbe, 0xac, 0x5d, 0x3a, 0x41, 0xe9, 0x6e, 0x6b, 0x97, 0x0f, 0xd9, 0xd6, 0xb7, 0x2a, 0x64, 0xe0,
	0x43, 0x36, 0x2a, 0x57, 0x86, 0xa4, 0xa4, 0x6b, 0x40, 0x5d, 0xfa, 0xa8, 0x2e, 0xb8, 0x6e, 0x24,
	0xc0, 0x14, 0x3a, 0x2f, 0xb4, 0xbb, 0x28, 0xf5, 0x3c, 0xac, 0x34, 0x06, 0xc1, 0x66, 0x2e, 0x07,
	0x1f, 0x99, 0xcb, 0xf5, 0x8f, 0xcc, 0x65, 0xf8, 0x91, 0xb9, 0xdc, 0xb8, 0x6e, 0x2e, 0x37, 0xaf,
	0x99, 0xcb, 0xad, 0xd6, 0x5c, 0xb6, 0x46, 0xec, 0xe6, 0xa5, 0x11, 0xbb, 0x47, 0x02, 0x33, 0x73,
	0x1c, 0x9b, 0x43, 0xf1, 0x02, 0x5f, 0x33, 0x33, 0x77, 0x0a, 0xad, 0x01, 0x1f, 0x31, 0xe1, 0xd9,
	0x62, 0x8e, 0x62, 0xdb, 0x8f, 0x7a, 0x4e, 0x4c, 0x5e, 0x2d, 0xe6, 0xf4, 0x31, 0xd9, 0x68, 0x1d,
	0x5e, 0xa2, 0xbc, 0xea, 0x06, 0xd1, 0x7a, 0x7d, 0x6e, 0xcf, 0x95, 0x06, 0x56, 0xab, 0x6b, 0xc0,
	0xda, 0xf6, 0xac, 0xba, 0x61, 0xc0, 0xba, 0x4b, 0xd6, 0x4c, 0x1c, 0x73, 0xab, 0x13, 0xd4, 0xe0,
	0x30, 0xea, 0x99, 0x38, 0x3e, 0xd1, 0x09, 0x16, 0x16, 0xc7, 0x3c, 0x3e, 0xd7, 0x09, 0xca, 0x6f,
	0x18, 0x01, 0xf1, 0xe0, 0x5c, 0x27, 0x95, 0x8f, 0x91, 0x06, 0x55, 0xd7, 0xfb, 0x1c, 0x4b, 0x53,
	0x19, 0x20, 0xd7, 0xbd, 0xda, 0x00, 0x59, 0x76, 0xc9, 0x60, 0x6e, 0x52, 0xfb, 0x0d, 0x4f, 0x45,
	0xac, 0x52, 0x94, 0xd8, 0x30, 0x22, 0x08, 0xfd, 0x08, 0x08, 0xde, 0x15, 0x48, 0x50, 0x4b, 0x83,
	0xea, 0x0a, 0x77, 0x05, 0x00, 0x47, 0x4b, 0xd3, 0x18, 0xe3, 0xdc, 0x96, 0xf2, 0xea, 0x8d, 0x07,
	0xb9, 0xad, 0x42, 0x8f, 0xcb, 0xd0, 0x3b, 0x4d, 0xe8, 0xf1, 0xa5, 0xd0, 0x63, 0x0c, 0xfd, 0xb0,
	0xf1, 0x1e, 0xb7, 0x42, 0x8f, 0x31, 0xf4, 0x6e, 0xcb, 0x08, 0xa1, 0x77, 0x08, 0x31, 0xa9, 0xd0,
	0x19, 0x8f, 0xb5, 0xb3, 0x28, 0xbd, 0xdd, 0xa8, 0x8f, 0xc8, 0x81, 0x76, 0x16, 0xc6, 0xb2, 0x31,
	0x73, 0xab, 0xdf, 0xa9, 0x52, 0x8a, 0xc3, 0x9a, 0x73, 0xa2, 0xdf, 0xe1, 0x57, 0xc5, 0x59, 0x0e,
	0xb7, 0x2b, 0x5c, 0x5f, 0xa8, 0xcb, 0x41, 0xd4, 0x47, 0x04, 0x2e, 0x3b, 0x78, 0xb7, 0x93, 0x22,
	0x37, 0xa5, 0x2a, 0xe3, 0x1a, 0x0e, 0x72, 0x2e, 0xec, 0x8c, 0x27, 0x02, 0x15, 0xb9, 0x1b, 0xf5,
	0x60, 0xfb, 0x5c, 0xf8, 0x97, 0xde, 0xce, 0x78, 0x33, 0xd1, 0x4f, 0xaa, 0x97, 0xde, 0xce, 0x8e,
	0xea, 0xa9, 0x7e, 0x4a, 0x36, 0x91, 0xd6, 0x2a, 0xff, 0x29, 0xc6, 0x41, 0xde, 0x71, 0xfd, 0x08,
	0xf7, 0x49, 0x1f, 0x65, 0x65, 0x0a, 0x8c, 0xaf, 0x90, 0x11, 0x00, 0xf0, 0x83, 0x76, 0x76, 0xf8,
	0xf7, 0x0a, 0x3c, 0x60, 0x5b, 0xa8, 0x1c, 0x7d, 0xff, 0x81, 0xd4, 0xbf, 0xfb, 0x64, 0x9a, 0xfa,
	0x81, 0x62, 0xd6, 0x4a, 0xff, 0xfe, 0x03, 0xa5, 0xff, 0xac, 0x55, 0x95, 0x42, 0xff, 0xef, 0x0a,
	0xa1, 0x57, 0x3e, 0xe7, 0xd5, 0xd2, 0x41, 0xbf, 0xb1, 0x71, 0xfe, 0x33, 0x1e, 0xd7, 0xf4, 0xf7,
	0x0e, 0x19, 0xb4, 0xae, 0x8c, 0xf2, 0x21, 0xce, 0x3f, 0xcb, 0x75, 0xe5, 0x22, 0xe2, 0xb7, 0xd5,
	0x25, 0x4a, 0x9a, 0x47, 0xc4, 0x0f, 0xfd, 0xc1, 0x78, 0xf1, 0x39, 0x4e, 0xd7, 0x45, 0x7d, 0xdc,
	0xe1, 0xd9, 0xfe, 0xd1, 0xf1, 0xdf, 0x21, 0xed, 0xbf, 0x4a, 0xf4, 0xd7, 0x0e, 0x59, 0x55, 0x4b,
	0xc7, 0xc6, 0x58, 0xe4, 0xf2, 0xd3, 0x9d, 0xde, 0xe5, 0x0e, 0x47, 0x50, 0x44, 0xdc, 0xc3, 0x3f,
	0x8e, 0xdf, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x22, 0xf7, 0x74, 0x42, 0x52, 0x0e, 0x00, 0x00,
}
