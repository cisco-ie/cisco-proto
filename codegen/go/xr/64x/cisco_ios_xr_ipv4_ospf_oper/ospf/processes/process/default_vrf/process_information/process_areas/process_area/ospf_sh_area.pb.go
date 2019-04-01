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
// source: ospf_sh_area.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_process_information_process_areas_process_area

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

type OspfShArea_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShArea_KEYS) Reset()         { *m = OspfShArea_KEYS{} }
func (m *OspfShArea_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShArea_KEYS) ProtoMessage()    {}
func (*OspfShArea_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c064299b5cec207a, []int{0}
}

func (m *OspfShArea_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShArea_KEYS.Unmarshal(m, b)
}
func (m *OspfShArea_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShArea_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShArea_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShArea_KEYS.Merge(m, src)
}
func (m *OspfShArea_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShArea_KEYS.Size(m)
}
func (m *OspfShArea_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShArea_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShArea_KEYS proto.InternalMessageInfo

func (m *OspfShArea_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShArea_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShArea_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type OspfShAreaRange struct {
	RangePrefix          string   `protobuf:"bytes,1,opt,name=range_prefix,json=rangePrefix,proto3" json:"range_prefix,omitempty"`
	RangeMask            string   `protobuf:"bytes,2,opt,name=range_mask,json=rangeMask,proto3" json:"range_mask,omitempty"`
	Cost                 uint32   `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
	AdvertiseFlag        bool     `protobuf:"varint,4,opt,name=advertise_flag,json=advertiseFlag,proto3" json:"advertise_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShAreaRange) Reset()         { *m = OspfShAreaRange{} }
func (m *OspfShAreaRange) String() string { return proto.CompactTextString(m) }
func (*OspfShAreaRange) ProtoMessage()    {}
func (*OspfShAreaRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_c064299b5cec207a, []int{1}
}

func (m *OspfShAreaRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShAreaRange.Unmarshal(m, b)
}
func (m *OspfShAreaRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShAreaRange.Marshal(b, m, deterministic)
}
func (m *OspfShAreaRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShAreaRange.Merge(m, src)
}
func (m *OspfShAreaRange) XXX_Size() int {
	return xxx_messageInfo_OspfShAreaRange.Size(m)
}
func (m *OspfShAreaRange) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShAreaRange.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShAreaRange proto.InternalMessageInfo

func (m *OspfShAreaRange) GetRangePrefix() string {
	if m != nil {
		return m.RangePrefix
	}
	return ""
}

func (m *OspfShAreaRange) GetRangeMask() string {
	if m != nil {
		return m.RangeMask
	}
	return ""
}

func (m *OspfShAreaRange) GetCost() uint32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *OspfShAreaRange) GetAdvertiseFlag() bool {
	if m != nil {
		return m.AdvertiseFlag
	}
	return false
}

type OspfShArea struct {
	AreaIdString                     string             `protobuf:"bytes,50,opt,name=area_id_string,json=areaIdString,proto3" json:"area_id_string,omitempty"`
	BackboneAreaActive               bool               `protobuf:"varint,51,opt,name=backbone_area_active,json=backboneAreaActive,proto3" json:"backbone_area_active,omitempty"`
	AreaInterfaceCount               uint32             `protobuf:"varint,52,opt,name=area_interface_count,json=areaInterfaceCount,proto3" json:"area_interface_count,omitempty"`
	AreaStub                         bool               `protobuf:"varint,53,opt,name=area_stub,json=areaStub,proto3" json:"area_stub,omitempty"`
	AreaTotalStub                    bool               `protobuf:"varint,54,opt,name=area_total_stub,json=areaTotalStub,proto3" json:"area_total_stub,omitempty"`
	StubDefaultCost                  uint32             `protobuf:"varint,55,opt,name=stub_default_cost,json=stubDefaultCost,proto3" json:"stub_default_cost,omitempty"`
	AreaNssa                         bool               `protobuf:"varint,56,opt,name=area_nssa,json=areaNssa,proto3" json:"area_nssa,omitempty"`
	NssaNoRedistribution             bool               `protobuf:"varint,57,opt,name=nssa_no_redistribution,json=nssaNoRedistribution,proto3" json:"nssa_no_redistribution,omitempty"`
	NssaTranslate                    bool               `protobuf:"varint,58,opt,name=nssa_translate,json=nssaTranslate,proto3" json:"nssa_translate,omitempty"`
	NssaTranslateAlways              bool               `protobuf:"varint,59,opt,name=nssa_translate_always,json=nssaTranslateAlways,proto3" json:"nssa_translate_always,omitempty"`
	NssaDefault                      bool               `protobuf:"varint,60,opt,name=nssa_default,json=nssaDefault,proto3" json:"nssa_default,omitempty"`
	TeEnabled                        bool               `protobuf:"varint,61,opt,name=te_enabled,json=teEnabled,proto3" json:"te_enabled,omitempty"`
	TeTopologyVersion                uint32             `protobuf:"varint,62,opt,name=te_topology_version,json=teTopologyVersion,proto3" json:"te_topology_version,omitempty"`
	ExternalOut                      bool               `protobuf:"varint,63,opt,name=external_out,json=externalOut,proto3" json:"external_out,omitempty"`
	SummaryIn                        bool               `protobuf:"varint,64,opt,name=summary_in,json=summaryIn,proto3" json:"summary_in,omitempty"`
	SegmentRouting                   string             `protobuf:"bytes,65,opt,name=segment_routing,json=segmentRouting,proto3" json:"segment_routing,omitempty"`
	SrStrictSpfCap                   bool               `protobuf:"varint,66,opt,name=sr_strict_spf_cap,json=srStrictSpfCap,proto3" json:"sr_strict_spf_cap,omitempty"`
	SrStrictSpfsidsAvailable         bool               `protobuf:"varint,67,opt,name=sr_strict_spfsids_available,json=srStrictSpfsidsAvailable,proto3" json:"sr_strict_spfsids_available,omitempty"`
	SrMicroloopAvoidanceActive       bool               `protobuf:"varint,68,opt,name=sr_microloop_avoidance_active,json=srMicroloopAvoidanceActive,proto3" json:"sr_microloop_avoidance_active,omitempty"`
	SrMicroloopAvoidanceEventType    string             `protobuf:"bytes,69,opt,name=sr_microloop_avoidance_event_type,json=srMicroloopAvoidanceEventType,proto3" json:"sr_microloop_avoidance_event_type,omitempty"`
	SrMicroloopAvoidanceNearEndId    string             `protobuf:"bytes,70,opt,name=sr_microloop_avoidance_near_end_id,json=srMicroloopAvoidanceNearEndId,proto3" json:"sr_microloop_avoidance_near_end_id,omitempty"`
	SrMicroloopAvoidanceFarEndId     string             `protobuf:"bytes,71,opt,name=sr_microloop_avoidance_far_end_id,json=srMicroloopAvoidanceFarEndId,proto3" json:"sr_microloop_avoidance_far_end_id,omitempty"`
	SrMicroloopAvoidancePseudonodeId string             `protobuf:"bytes,72,opt,name=sr_microloop_avoidance_pseudonode_id,json=srMicroloopAvoidancePseudonodeId,proto3" json:"sr_microloop_avoidance_pseudonode_id,omitempty"`
	AuthenticationType               string             `protobuf:"bytes,73,opt,name=authentication_type,json=authenticationType,proto3" json:"authentication_type,omitempty"`
	SpfCount                         uint32             `protobuf:"varint,74,opt,name=spf_count,json=spfCount,proto3" json:"spf_count,omitempty"`
	AreaPolicyIn                     bool               `protobuf:"varint,75,opt,name=area_policy_in,json=areaPolicyIn,proto3" json:"area_policy_in,omitempty"`
	AreaPolicyInName                 string             `protobuf:"bytes,76,opt,name=area_policy_in_name,json=areaPolicyInName,proto3" json:"area_policy_in_name,omitempty"`
	AreaPolicyOut                    bool               `protobuf:"varint,77,opt,name=area_policy_out,json=areaPolicyOut,proto3" json:"area_policy_out,omitempty"`
	AreaPolicyOutName                string             `protobuf:"bytes,78,opt,name=area_policy_out_name,json=areaPolicyOutName,proto3" json:"area_policy_out_name,omitempty"`
	AreaRange                        []*OspfShAreaRange `protobuf:"bytes,79,rep,name=area_range,json=areaRange,proto3" json:"area_range,omitempty"`
	AreaLsaCount                     uint32             `protobuf:"varint,80,opt,name=area_lsa_count,json=areaLsaCount,proto3" json:"area_lsa_count,omitempty"`
	AreaLsaChecksum                  uint32             `protobuf:"varint,81,opt,name=area_lsa_checksum,json=areaLsaChecksum,proto3" json:"area_lsa_checksum,omitempty"`
	AreaOpaqueLsaCount               uint32             `protobuf:"varint,82,opt,name=area_opaque_lsa_count,json=areaOpaqueLsaCount,proto3" json:"area_opaque_lsa_count,omitempty"`
	AreaOpaqueLsaChecksum            uint32             `protobuf:"varint,83,opt,name=area_opaque_lsa_checksum,json=areaOpaqueLsaChecksum,proto3" json:"area_opaque_lsa_checksum,omitempty"`
	AreaDcBitlessLsaCount            uint32             `protobuf:"varint,84,opt,name=area_dc_bitless_lsa_count,json=areaDcBitlessLsaCount,proto3" json:"area_dc_bitless_lsa_count,omitempty"`
	IndicationLsaCount               uint32             `protobuf:"varint,85,opt,name=indication_lsa_count,json=indicationLsaCount,proto3" json:"indication_lsa_count,omitempty"`
	DnaLsaCount                      uint32             `protobuf:"varint,86,opt,name=dna_lsa_count,json=dnaLsaCount,proto3" json:"dna_lsa_count,omitempty"`
	FloodListLength                  uint32             `protobuf:"varint,87,opt,name=flood_list_length,json=floodListLength,proto3" json:"flood_list_length,omitempty"`
	AreaLfaInterfaceCount            uint32             `protobuf:"varint,88,opt,name=area_lfa_interface_count,json=areaLfaInterfaceCount,proto3" json:"area_lfa_interface_count,omitempty"`
	AreaPerPrefixLfaInterfaceCount   uint32             `protobuf:"varint,89,opt,name=area_per_prefix_lfa_interface_count,json=areaPerPrefixLfaInterfaceCount,proto3" json:"area_per_prefix_lfa_interface_count,omitempty"`
	AreaLfaRevision                  uint32             `protobuf:"varint,90,opt,name=area_lfa_revision,json=areaLfaRevision,proto3" json:"area_lfa_revision,omitempty"`
	AreaAdjStagNumNbrForming         uint32             `protobuf:"varint,91,opt,name=area_adj_stag_num_nbr_forming,json=areaAdjStagNumNbrForming,proto3" json:"area_adj_stag_num_nbr_forming,omitempty"`
	AreaNumNbrFull                   uint32             `protobuf:"varint,92,opt,name=area_num_nbr_full,json=areaNumNbrFull,proto3" json:"area_num_nbr_full,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}           `json:"-"`
	XXX_unrecognized                 []byte             `json:"-"`
	XXX_sizecache                    int32              `json:"-"`
}

func (m *OspfShArea) Reset()         { *m = OspfShArea{} }
func (m *OspfShArea) String() string { return proto.CompactTextString(m) }
func (*OspfShArea) ProtoMessage()    {}
func (*OspfShArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_c064299b5cec207a, []int{2}
}

func (m *OspfShArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShArea.Unmarshal(m, b)
}
func (m *OspfShArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShArea.Marshal(b, m, deterministic)
}
func (m *OspfShArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShArea.Merge(m, src)
}
func (m *OspfShArea) XXX_Size() int {
	return xxx_messageInfo_OspfShArea.Size(m)
}
func (m *OspfShArea) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShArea.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShArea proto.InternalMessageInfo

func (m *OspfShArea) GetAreaIdString() string {
	if m != nil {
		return m.AreaIdString
	}
	return ""
}

func (m *OspfShArea) GetBackboneAreaActive() bool {
	if m != nil {
		return m.BackboneAreaActive
	}
	return false
}

func (m *OspfShArea) GetAreaInterfaceCount() uint32 {
	if m != nil {
		return m.AreaInterfaceCount
	}
	return 0
}

func (m *OspfShArea) GetAreaStub() bool {
	if m != nil {
		return m.AreaStub
	}
	return false
}

func (m *OspfShArea) GetAreaTotalStub() bool {
	if m != nil {
		return m.AreaTotalStub
	}
	return false
}

func (m *OspfShArea) GetStubDefaultCost() uint32 {
	if m != nil {
		return m.StubDefaultCost
	}
	return 0
}

func (m *OspfShArea) GetAreaNssa() bool {
	if m != nil {
		return m.AreaNssa
	}
	return false
}

func (m *OspfShArea) GetNssaNoRedistribution() bool {
	if m != nil {
		return m.NssaNoRedistribution
	}
	return false
}

func (m *OspfShArea) GetNssaTranslate() bool {
	if m != nil {
		return m.NssaTranslate
	}
	return false
}

func (m *OspfShArea) GetNssaTranslateAlways() bool {
	if m != nil {
		return m.NssaTranslateAlways
	}
	return false
}

func (m *OspfShArea) GetNssaDefault() bool {
	if m != nil {
		return m.NssaDefault
	}
	return false
}

func (m *OspfShArea) GetTeEnabled() bool {
	if m != nil {
		return m.TeEnabled
	}
	return false
}

func (m *OspfShArea) GetTeTopologyVersion() uint32 {
	if m != nil {
		return m.TeTopologyVersion
	}
	return 0
}

func (m *OspfShArea) GetExternalOut() bool {
	if m != nil {
		return m.ExternalOut
	}
	return false
}

func (m *OspfShArea) GetSummaryIn() bool {
	if m != nil {
		return m.SummaryIn
	}
	return false
}

func (m *OspfShArea) GetSegmentRouting() string {
	if m != nil {
		return m.SegmentRouting
	}
	return ""
}

func (m *OspfShArea) GetSrStrictSpfCap() bool {
	if m != nil {
		return m.SrStrictSpfCap
	}
	return false
}

func (m *OspfShArea) GetSrStrictSpfsidsAvailable() bool {
	if m != nil {
		return m.SrStrictSpfsidsAvailable
	}
	return false
}

func (m *OspfShArea) GetSrMicroloopAvoidanceActive() bool {
	if m != nil {
		return m.SrMicroloopAvoidanceActive
	}
	return false
}

func (m *OspfShArea) GetSrMicroloopAvoidanceEventType() string {
	if m != nil {
		return m.SrMicroloopAvoidanceEventType
	}
	return ""
}

func (m *OspfShArea) GetSrMicroloopAvoidanceNearEndId() string {
	if m != nil {
		return m.SrMicroloopAvoidanceNearEndId
	}
	return ""
}

func (m *OspfShArea) GetSrMicroloopAvoidanceFarEndId() string {
	if m != nil {
		return m.SrMicroloopAvoidanceFarEndId
	}
	return ""
}

func (m *OspfShArea) GetSrMicroloopAvoidancePseudonodeId() string {
	if m != nil {
		return m.SrMicroloopAvoidancePseudonodeId
	}
	return ""
}

func (m *OspfShArea) GetAuthenticationType() string {
	if m != nil {
		return m.AuthenticationType
	}
	return ""
}

func (m *OspfShArea) GetSpfCount() uint32 {
	if m != nil {
		return m.SpfCount
	}
	return 0
}

func (m *OspfShArea) GetAreaPolicyIn() bool {
	if m != nil {
		return m.AreaPolicyIn
	}
	return false
}

func (m *OspfShArea) GetAreaPolicyInName() string {
	if m != nil {
		return m.AreaPolicyInName
	}
	return ""
}

func (m *OspfShArea) GetAreaPolicyOut() bool {
	if m != nil {
		return m.AreaPolicyOut
	}
	return false
}

func (m *OspfShArea) GetAreaPolicyOutName() string {
	if m != nil {
		return m.AreaPolicyOutName
	}
	return ""
}

func (m *OspfShArea) GetAreaRange() []*OspfShAreaRange {
	if m != nil {
		return m.AreaRange
	}
	return nil
}

func (m *OspfShArea) GetAreaLsaCount() uint32 {
	if m != nil {
		return m.AreaLsaCount
	}
	return 0
}

func (m *OspfShArea) GetAreaLsaChecksum() uint32 {
	if m != nil {
		return m.AreaLsaChecksum
	}
	return 0
}

func (m *OspfShArea) GetAreaOpaqueLsaCount() uint32 {
	if m != nil {
		return m.AreaOpaqueLsaCount
	}
	return 0
}

func (m *OspfShArea) GetAreaOpaqueLsaChecksum() uint32 {
	if m != nil {
		return m.AreaOpaqueLsaChecksum
	}
	return 0
}

func (m *OspfShArea) GetAreaDcBitlessLsaCount() uint32 {
	if m != nil {
		return m.AreaDcBitlessLsaCount
	}
	return 0
}

func (m *OspfShArea) GetIndicationLsaCount() uint32 {
	if m != nil {
		return m.IndicationLsaCount
	}
	return 0
}

func (m *OspfShArea) GetDnaLsaCount() uint32 {
	if m != nil {
		return m.DnaLsaCount
	}
	return 0
}

func (m *OspfShArea) GetFloodListLength() uint32 {
	if m != nil {
		return m.FloodListLength
	}
	return 0
}

func (m *OspfShArea) GetAreaLfaInterfaceCount() uint32 {
	if m != nil {
		return m.AreaLfaInterfaceCount
	}
	return 0
}

func (m *OspfShArea) GetAreaPerPrefixLfaInterfaceCount() uint32 {
	if m != nil {
		return m.AreaPerPrefixLfaInterfaceCount
	}
	return 0
}

func (m *OspfShArea) GetAreaLfaRevision() uint32 {
	if m != nil {
		return m.AreaLfaRevision
	}
	return 0
}

func (m *OspfShArea) GetAreaAdjStagNumNbrForming() uint32 {
	if m != nil {
		return m.AreaAdjStagNumNbrForming
	}
	return 0
}

func (m *OspfShArea) GetAreaNumNbrFull() uint32 {
	if m != nil {
		return m.AreaNumNbrFull
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShArea_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area_KEYS")
	proto.RegisterType((*OspfShAreaRange)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area_range")
	proto.RegisterType((*OspfShArea)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area")
}

func init() { proto.RegisterFile("ospf_sh_area.proto", fileDescriptor_c064299b5cec207a) }

var fileDescriptor_c064299b5cec207a = []byte{
	// 1151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x6d, 0x73, 0x1b, 0x35,
	0x10, 0x1e, 0xd3, 0x4e, 0x5b, 0x2b, 0x4d, 0x8a, 0x95, 0x16, 0x04, 0x25, 0x4c, 0x1a, 0x0a, 0x84,
	0xce, 0xe0, 0x96, 0x34, 0xd0, 0xf2, 0x52, 0x8a, 0x9b, 0x97, 0xd6, 0x24, 0x71, 0x82, 0x1d, 0x0a,
	0x05, 0x66, 0x34, 0xf2, 0xdd, 0x9e, 0xa3, 0xe6, 0x4e, 0x3a, 0x24, 0x9d, 0x49, 0x3e, 0xf2, 0x13,
	0xf8, 0x09, 0x7c, 0xe2, 0x6f, 0x32, 0x5a, 0xdd, 0xf9, 0x6c, 0x92, 0x7c, 0xe5, 0xdb, 0xdd, 0xb3,
	0xcf, 0x3e, 0xab, 0xdd, 0x7b, 0x66, 0x75, 0x84, 0x6a, 0x9b, 0x27, 0xdc, 0x1e, 0x71, 0x61, 0x40,
	0xb4, 0x73, 0xa3, 0x9d, 0xa6, 0x59, 0x24, 0x6d, 0xa4, 0xb9, 0xd4, 0x96, 0x9f, 0x18, 0x2e, 0xf3,
	0xf1, 0x3a, 0x47, 0x96, 0xce, 0xc1, 0xb4, 0xfd, 0x93, 0xe7, 0x45, 0x60, 0x2d, 0xd8, 0xea, 0xa9,
	0x1d, 0x43, 0x22, 0x8a, 0xd4, 0xf1, 0xb1, 0x99, 0x44, 0xb9, 0x54, 0x89, 0x36, 0x99, 0x70, 0x52,
	0xab, 0x09, 0xe6, 0xcb, 0xd8, 0x99, 0xb7, 0x15, 0x49, 0x5a, 0xd3, 0x87, 0xe0, 0x3b, 0x5b, 0xaf,
	0x06, 0xf4, 0x0e, 0xb9, 0x5e, 0x91, 0x94, 0xc8, 0x80, 0x35, 0x96, 0x1b, 0xab, 0xcd, 0xfe, 0x5c,
	0x89, 0xf5, 0x44, 0x06, 0xf4, 0x6d, 0x72, 0x15, 0xf9, 0x32, 0x66, 0x6f, 0x2c, 0x37, 0x56, 0xe7,
	0xfb, 0x57, 0xfc, 0x6b, 0x37, 0xa6, 0x8c, 0x5c, 0x15, 0x71, 0x6c, 0xc0, 0x5a, 0x76, 0x09, 0xd3,
	0xaa, 0xd7, 0x95, 0xbf, 0x1a, 0xb3, 0x0d, 0x73, 0x23, 0xd4, 0x08, 0x7c, 0x31, 0x7c, 0xe0, 0xb9,
	0x81, 0x44, 0x9e, 0x54, 0xc5, 0x10, 0x3b, 0x40, 0x88, 0x2e, 0x11, 0x12, 0x28, 0x99, 0xb0, 0xc7,
	0x58, 0xaf, 0xd9, 0x6f, 0x22, 0xb2, 0x27, 0xec, 0x31, 0xa5, 0xe4, 0x72, 0xa4, 0xad, 0xc3, 0x7a,
	0xf3, 0x7d, 0x7c, 0xa6, 0x1f, 0x92, 0x05, 0x11, 0x8f, 0xc1, 0x38, 0x69, 0x81, 0x27, 0xa9, 0x18,
	0xb1, 0xcb, 0xcb, 0x8d, 0xd5, 0x6b, 0xfd, 0xf9, 0x09, 0xba, 0x9d, 0x8a, 0xd1, 0xca, 0x3f, 0x2d,
	0x72, 0x7d, 0xfa, 0x4c, 0xf4, 0x2e, 0x59, 0x28, 0xfb, 0xe2, 0xd6, 0x19, 0xa9, 0x46, 0x6c, 0x0d,
	0xcb, 0x5d, 0x0f, 0xed, 0x0d, 0x10, 0xa3, 0x0f, 0xc8, 0xcd, 0xa1, 0x88, 0x8e, 0x87, 0x5a, 0x41,
	0x68, 0x45, 0x44, 0x4e, 0x8e, 0x81, 0x3d, 0xc4, 0x1a, 0xb4, 0x8a, 0x75, 0x0c, 0x88, 0x0e, 0x46,
	0x7c, 0x46, 0xd0, 0x55, 0x0e, 0x4c, 0x22, 0x22, 0xe0, 0x91, 0x2e, 0x94, 0x63, 0xeb, 0x78, 0x66,
	0x8a, 0xea, 0x55, 0x68, 0xc3, 0x47, 0xe8, 0x6d, 0xd2, 0xc4, 0x0c, 0xeb, 0x8a, 0x21, 0xfb, 0x1c,
	0x85, 0xaf, 0x79, 0x60, 0xe0, 0x8a, 0x21, 0xfd, 0x88, 0xdc, 0xc0, 0xa0, 0xd3, 0x4e, 0xa4, 0x81,
	0xf2, 0x45, 0xd9, 0x9f, 0x01, 0x71, 0xe8, 0x51, 0xe4, 0xdd, 0x23, 0x2d, 0x1f, 0xe4, 0x95, 0x43,
	0x70, 0x4e, 0x8f, 0xb0, 0xe6, 0x0d, 0x1f, 0xd8, 0x0c, 0xf8, 0x86, 0x1f, 0x59, 0x55, 0x50, 0x59,
	0x2b, 0xd8, 0xe3, 0xba, 0x60, 0xcf, 0x5a, 0x41, 0xd7, 0xc9, 0x5b, 0x1e, 0xe7, 0x4a, 0x73, 0x03,
	0xb1, 0xf4, 0xb3, 0x19, 0x16, 0xde, 0x5c, 0xec, 0x4b, 0x64, 0xde, 0xf4, 0xd1, 0x9e, 0xee, 0xcf,
	0xc4, 0xfc, 0x57, 0xc0, 0x2c, 0x67, 0x84, 0xb2, 0xa9, 0x70, 0xc0, 0xbe, 0x0a, 0xa7, 0xf4, 0xe8,
	0x61, 0x05, 0xd2, 0x35, 0x72, 0x6b, 0x96, 0xc6, 0x45, 0xfa, 0x87, 0x38, 0xb5, 0xec, 0x6b, 0x64,
	0x2f, 0xce, 0xb0, 0x3b, 0x18, 0xf2, 0xb6, 0xc1, 0x9c, 0xb2, 0x33, 0xf6, 0x0d, 0x52, 0xe7, 0x3c,
	0x56, 0x36, 0xe5, 0x6d, 0xe3, 0x80, 0x83, 0x12, 0xc3, 0x14, 0x62, 0xf6, 0x04, 0x09, 0x4d, 0x07,
	0x5b, 0x01, 0xa0, 0x6d, 0xb2, 0xe8, 0x80, 0x3b, 0x9d, 0xeb, 0x54, 0x8f, 0x4e, 0xf9, 0x18, 0x8c,
	0xf5, 0xfd, 0x7c, 0x8b, 0xd3, 0x69, 0x39, 0x38, 0x2c, 0x23, 0x2f, 0x43, 0xc0, 0x57, 0x84, 0x13,
	0x07, 0x46, 0x89, 0x94, 0xeb, 0xc2, 0xb1, 0xa7, 0xa1, 0x62, 0x85, 0xed, 0x17, 0x58, 0xd1, 0x16,
	0x59, 0x26, 0xcc, 0x29, 0x97, 0x8a, 0x7d, 0x17, 0x2a, 0x96, 0x48, 0x57, 0xd1, 0x8f, 0xc9, 0x0d,
	0x0b, 0xa3, 0x0c, 0x94, 0xe3, 0x46, 0x17, 0xce, 0xbb, 0xab, 0x83, 0xee, 0x5a, 0x28, 0xe1, 0x7e,
	0x40, 0xe9, 0x27, 0xa4, 0x65, 0x0d, 0x1a, 0x30, 0x72, 0xdc, 0xdb, 0x33, 0x12, 0x39, 0x7b, 0x86,
	0x72, 0x0b, 0xd6, 0x0c, 0x10, 0x1f, 0xe4, 0xc9, 0x86, 0xc8, 0xe9, 0x13, 0x72, 0x7b, 0x86, 0x6a,
	0x65, 0x6c, 0xb9, 0x18, 0x0b, 0x99, 0xfa, 0x2e, 0xd9, 0x06, 0x26, 0xb1, 0xa9, 0x24, 0x4f, 0xe8,
	0x54, 0x71, 0xda, 0x21, 0x4b, 0xd6, 0xf0, 0x4c, 0x46, 0x46, 0xa7, 0x5a, 0xe7, 0x5c, 0x8c, 0xb5,
	0x8c, 0x85, 0x8a, 0xa0, 0xb2, 0xf4, 0x26, 0x0a, 0xbc, 0x6b, 0xcd, 0x5e, 0xc5, 0xe9, 0x54, 0x94,
	0xd2, 0xda, 0x2f, 0xc8, 0x9d, 0x0b, 0x24, 0x60, 0xec, 0x5b, 0x75, 0xa7, 0x39, 0xb0, 0x2d, 0xec,
	0x73, 0xe9, 0x3c, 0x99, 0x2d, 0xcf, 0x3a, 0x3c, 0xcd, 0x81, 0x76, 0xc9, 0xca, 0x05, 0x4a, 0x0a,
	0x84, 0xe1, 0xa0, 0x62, 0xbf, 0x6f, 0xb6, 0x2f, 0x96, 0xea, 0x81, 0x30, 0x5b, 0x2a, 0xee, 0xc6,
	0xf4, 0xf9, 0x85, 0x87, 0x4a, 0x6a, 0xa5, 0xe7, 0xa8, 0xf4, 0xde, 0x79, 0x4a, 0xdb, 0x95, 0x50,
	0x8f, 0xdc, 0xbd, 0x40, 0x28, 0xb7, 0x50, 0xc4, 0x5a, 0xe9, 0x18, 0xbc, 0xd6, 0x0b, 0xd4, 0x5a,
	0x3e, 0x4f, 0xeb, 0x60, 0x42, 0xec, 0xc6, 0xf4, 0x3e, 0x59, 0x14, 0x85, 0x3b, 0x02, 0xe5, 0x64,
	0x84, 0xdb, 0x39, 0xcc, 0xa7, 0x8b, 0xe9, 0x74, 0x36, 0x84, 0x43, 0xb9, 0x4d, 0x9a, 0xe8, 0x00,
	0x5c, 0x17, 0xdf, 0xa3, 0x39, 0xaf, 0xd9, 0x3c, 0x09, 0x4b, 0xa2, 0x5a, 0x57, 0xb9, 0x4e, 0x65,
	0x84, 0xa6, 0xdb, 0xc1, 0xef, 0x85, 0xeb, 0xea, 0x00, 0xc1, 0xae, 0xa2, 0x9f, 0x92, 0xc5, 0x59,
	0x56, 0x58, 0xeb, 0xbb, 0x58, 0xf3, 0xcd, 0x69, 0x2a, 0xee, 0xf6, 0x6a, 0xb9, 0x94, 0x74, 0xef,
	0xf5, 0xbd, 0x7a, 0xb9, 0x04, 0xaa, 0x77, 0xfb, 0xfd, 0x72, 0xa7, 0xd5, 0xbc, 0xa0, 0xdb, 0x43,
	0xdd, 0xd6, 0x0c, 0x19, 0x85, 0xff, 0x6e, 0x10, 0x52, 0x6f, 0x7e, 0xb6, 0xbf, 0x7c, 0x69, 0x75,
	0x6e, 0xed, 0xcf, 0x46, 0xfb, 0x7f, 0xbd, 0xf2, 0xda, 0x67, 0xef, 0xa0, 0x3e, 0x2e, 0xbe, 0x3e,
	0x5e, 0x47, 0xd5, 0x44, 0x53, 0x2b, 0xca, 0x99, 0x1f, 0xe0, 0xcc, 0x71, 0xa2, 0xbb, 0x56, 0x84,
	0xb9, 0xdf, 0x23, 0xad, 0x9a, 0x75, 0x04, 0xd1, 0xb1, 0x2d, 0x32, 0xf6, 0x43, 0xd8, 0xab, 0x15,
	0xb1, 0x84, 0xe9, 0x67, 0xe4, 0x16, 0x72, 0x75, 0x2e, 0x7e, 0x2f, 0x60, 0x4a, 0xb8, 0x5f, 0xef,
	0xfe, 0x7d, 0x8c, 0x4d, 0xe4, 0x1f, 0x11, 0x76, 0x26, 0xa5, 0xaa, 0x32, 0xc0, 0xac, 0x5b, 0xb3,
	0x59, 0x55, 0xad, 0xc7, 0xe4, 0x1d, 0x4c, 0x8c, 0x23, 0x3e, 0x94, 0x2e, 0xf5, 0x3d, 0xd7, 0xf5,
	0x0e, 0xeb, 0xcc, 0xcd, 0xe8, 0x59, 0x08, 0x4f, 0x4a, 0x3e, 0x20, 0x37, 0xa5, 0x8a, 0x2b, 0x4f,
	0xd6, 0x49, 0x3f, 0x86, 0x43, 0xd6, 0xb1, 0x49, 0xc6, 0x0a, 0x99, 0x8f, 0xd5, 0xf4, 0xa0, 0x5e,
	0x22, 0x75, 0x2e, 0x56, 0x33, 0x73, 0x4a, 0x52, 0xad, 0x63, 0x9e, 0x4a, 0xeb, 0x78, 0x0a, 0x6a,
	0xe4, 0x8e, 0xd8, 0x4f, 0x61, 0x4e, 0x18, 0xd8, 0x95, 0xd6, 0xed, 0x22, 0x3c, 0x69, 0x3a, 0x4d,
	0xce, 0x5e, 0x93, 0x3f, 0xd7, 0x47, 0xdf, 0x4d, 0xfe, 0x7b, 0x53, 0xee, 0x90, 0x0f, 0x82, 0x0f,
	0xc1, 0x94, 0x3f, 0x11, 0xe7, 0x6a, 0xbc, 0x42, 0x8d, 0xf7, 0xd1, 0x96, 0x60, 0xc2, 0xaf, 0xc5,
	0x59, 0xb1, 0xc9, 0x97, 0x4d, 0x04, 0x37, 0x30, 0x96, 0x78, 0x27, 0xfc, 0x32, 0xf5, 0x65, 0x13,
	0xd1, 0x2f, 0x61, 0xfa, 0x94, 0x2c, 0x85, 0xdb, 0x3f, 0x7e, 0xcd, 0xad, 0x13, 0x23, 0xae, 0x8a,
	0x8c, 0xab, 0xa1, 0xe1, 0xde, 0x87, 0x7e, 0xbb, 0xff, 0x8a, 0x79, 0xd8, 0x56, 0x27, 0x7e, 0x3d,
	0x70, 0x62, 0xd4, 0x2b, 0xb2, 0xde, 0xd0, 0x6c, 0x87, 0xb8, 0xdf, 0xf3, 0xe1, 0xca, 0xad, 0xf2,
	0x8a, 0x34, 0x65, 0xbf, 0x61, 0x12, 0xba, 0xb0, 0x64, 0x17, 0x69, 0x3a, 0xbc, 0x82, 0xbf, 0x87,
	0x0f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x68, 0x80, 0x1a, 0x3d, 0x34, 0x0a, 0x00, 0x00,
}