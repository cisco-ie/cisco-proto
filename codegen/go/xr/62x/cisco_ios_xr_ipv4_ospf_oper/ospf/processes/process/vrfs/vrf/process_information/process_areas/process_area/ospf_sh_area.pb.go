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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_process_information_process_areas_process_area

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
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

func (m *OspfShArea_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	NssaDefault                      bool               `protobuf:"varint,59,opt,name=nssa_default,json=nssaDefault,proto3" json:"nssa_default,omitempty"`
	TeEnabled                        bool               `protobuf:"varint,60,opt,name=te_enabled,json=teEnabled,proto3" json:"te_enabled,omitempty"`
	TeTopologyVersion                uint32             `protobuf:"varint,61,opt,name=te_topology_version,json=teTopologyVersion,proto3" json:"te_topology_version,omitempty"`
	ExternalOut                      bool               `protobuf:"varint,62,opt,name=external_out,json=externalOut,proto3" json:"external_out,omitempty"`
	SummaryIn                        bool               `protobuf:"varint,63,opt,name=summary_in,json=summaryIn,proto3" json:"summary_in,omitempty"`
	SegmentRouting                   uint32             `protobuf:"varint,64,opt,name=segment_routing,json=segmentRouting,proto3" json:"segment_routing,omitempty"`
	SrStrictSpfCap                   bool               `protobuf:"varint,65,opt,name=sr_strict_spf_cap,json=srStrictSpfCap,proto3" json:"sr_strict_spf_cap,omitempty"`
	SrStrictSpfsidsAvailable         bool               `protobuf:"varint,66,opt,name=sr_strict_spfsids_available,json=srStrictSpfsidsAvailable,proto3" json:"sr_strict_spfsids_available,omitempty"`
	SrMicroloopAvoidanceActive       bool               `protobuf:"varint,67,opt,name=sr_microloop_avoidance_active,json=srMicroloopAvoidanceActive,proto3" json:"sr_microloop_avoidance_active,omitempty"`
	SrMicroloopAvoidanceEventType    string             `protobuf:"bytes,68,opt,name=sr_microloop_avoidance_event_type,json=srMicroloopAvoidanceEventType,proto3" json:"sr_microloop_avoidance_event_type,omitempty"`
	SrMicroloopAvoidanceNearEndId    string             `protobuf:"bytes,69,opt,name=sr_microloop_avoidance_near_end_id,json=srMicroloopAvoidanceNearEndId,proto3" json:"sr_microloop_avoidance_near_end_id,omitempty"`
	SrMicroloopAvoidanceFarEndId     string             `protobuf:"bytes,70,opt,name=sr_microloop_avoidance_far_end_id,json=srMicroloopAvoidanceFarEndId,proto3" json:"sr_microloop_avoidance_far_end_id,omitempty"`
	SrMicroloopAvoidancePseudonodeId string             `protobuf:"bytes,71,opt,name=sr_microloop_avoidance_pseudonode_id,json=srMicroloopAvoidancePseudonodeId,proto3" json:"sr_microloop_avoidance_pseudonode_id,omitempty"`
	AuthenticationType               string             `protobuf:"bytes,72,opt,name=authentication_type,json=authenticationType,proto3" json:"authentication_type,omitempty"`
	SpfCount                         uint32             `protobuf:"varint,73,opt,name=spf_count,json=spfCount,proto3" json:"spf_count,omitempty"`
	AreaPolicyIn                     bool               `protobuf:"varint,74,opt,name=area_policy_in,json=areaPolicyIn,proto3" json:"area_policy_in,omitempty"`
	AreaPolicyInName                 string             `protobuf:"bytes,75,opt,name=area_policy_in_name,json=areaPolicyInName,proto3" json:"area_policy_in_name,omitempty"`
	AreaPolicyOut                    bool               `protobuf:"varint,76,opt,name=area_policy_out,json=areaPolicyOut,proto3" json:"area_policy_out,omitempty"`
	AreaPolicyOutName                string             `protobuf:"bytes,77,opt,name=area_policy_out_name,json=areaPolicyOutName,proto3" json:"area_policy_out_name,omitempty"`
	AreaRange                        []*OspfShAreaRange `protobuf:"bytes,78,rep,name=area_range,json=areaRange,proto3" json:"area_range,omitempty"`
	AreaLsaCount                     uint32             `protobuf:"varint,79,opt,name=area_lsa_count,json=areaLsaCount,proto3" json:"area_lsa_count,omitempty"`
	AreaLsaChecksum                  uint32             `protobuf:"varint,80,opt,name=area_lsa_checksum,json=areaLsaChecksum,proto3" json:"area_lsa_checksum,omitempty"`
	AreaOpaqueLsaCount               uint32             `protobuf:"varint,81,opt,name=area_opaque_lsa_count,json=areaOpaqueLsaCount,proto3" json:"area_opaque_lsa_count,omitempty"`
	AreaOpaqueLsaChecksum            uint32             `protobuf:"varint,82,opt,name=area_opaque_lsa_checksum,json=areaOpaqueLsaChecksum,proto3" json:"area_opaque_lsa_checksum,omitempty"`
	AreaDcBitlessLsaCount            uint32             `protobuf:"varint,83,opt,name=area_dc_bitless_lsa_count,json=areaDcBitlessLsaCount,proto3" json:"area_dc_bitless_lsa_count,omitempty"`
	IndicationLsaCount               uint32             `protobuf:"varint,84,opt,name=indication_lsa_count,json=indicationLsaCount,proto3" json:"indication_lsa_count,omitempty"`
	DnaLsaCount                      uint32             `protobuf:"varint,85,opt,name=dna_lsa_count,json=dnaLsaCount,proto3" json:"dna_lsa_count,omitempty"`
	FloodListLength                  uint32             `protobuf:"varint,86,opt,name=flood_list_length,json=floodListLength,proto3" json:"flood_list_length,omitempty"`
	AreaLfaInterfaceCount            uint32             `protobuf:"varint,87,opt,name=area_lfa_interface_count,json=areaLfaInterfaceCount,proto3" json:"area_lfa_interface_count,omitempty"`
	AreaPerPrefixLfaInterfaceCount   uint32             `protobuf:"varint,88,opt,name=area_per_prefix_lfa_interface_count,json=areaPerPrefixLfaInterfaceCount,proto3" json:"area_per_prefix_lfa_interface_count,omitempty"`
	AreaLfaRevision                  uint32             `protobuf:"varint,89,opt,name=area_lfa_revision,json=areaLfaRevision,proto3" json:"area_lfa_revision,omitempty"`
	AreaAdjStagNumNbrForming         uint32             `protobuf:"varint,90,opt,name=area_adj_stag_num_nbr_forming,json=areaAdjStagNumNbrForming,proto3" json:"area_adj_stag_num_nbr_forming,omitempty"`
	AreaNumNbrFull                   uint32             `protobuf:"varint,91,opt,name=area_num_nbr_full,json=areaNumNbrFull,proto3" json:"area_num_nbr_full,omitempty"`
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

func (m *OspfShArea) GetSegmentRouting() uint32 {
	if m != nil {
		return m.SegmentRouting
	}
	return 0
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
	proto.RegisterType((*OspfShArea_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.process_areas.process_area.ospf_sh_area_KEYS")
	proto.RegisterType((*OspfShAreaRange)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.process_areas.process_area.ospf_sh_area_range")
	proto.RegisterType((*OspfShArea)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.process_information.process_areas.process_area.ospf_sh_area")
}

func init() { proto.RegisterFile("ospf_sh_area.proto", fileDescriptor_c064299b5cec207a) }

var fileDescriptor_c064299b5cec207a = []byte{
	// 1147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x72, 0x1b, 0x45,
	0x13, 0x2d, 0x7d, 0x49, 0xc5, 0xf6, 0xf8, 0x2f, 0x1a, 0x3b, 0x1f, 0x13, 0x8c, 0x29, 0xc7, 0x04,
	0x30, 0xa9, 0x42, 0x09, 0x89, 0x21, 0xe1, 0x27, 0x04, 0xc5, 0x3f, 0x89, 0xb0, 0x2d, 0x1b, 0xc9,
	0x04, 0x02, 0x17, 0x53, 0xa3, 0xdd, 0x5e, 0x79, 0xec, 0xd5, 0xcc, 0x32, 0x33, 0xab, 0xb2, 0x6f,
	0xb8, 0xe2, 0x05, 0x78, 0x00, 0x9e, 0x91, 0x57, 0xa0, 0xa6, 0x67, 0x57, 0x2b, 0x61, 0xfb, 0x96,
	0x1b, 0xd5, 0xea, 0x9c, 0xd3, 0xa7, 0x7b, 0x7a, 0x5b, 0x3d, 0x22, 0x54, 0xdb, 0x2c, 0xe1, 0xf6,
	0x84, 0x0b, 0x03, 0xa2, 0x91, 0x19, 0xed, 0x34, 0x3d, 0x8d, 0xa4, 0x8d, 0x34, 0x97, 0xda, 0xf2,
	0x73, 0xc3, 0x65, 0x36, 0xdc, 0xe4, 0xa8, 0xd2, 0x19, 0x98, 0x86, 0x7f, 0xf2, 0xba, 0x08, 0xac,
	0x05, 0x5b, 0x3e, 0x35, 0x86, 0x26, 0xc1, 0x8f, 0x12, 0xe0, 0x52, 0x25, 0xda, 0x0c, 0x84, 0x93,
	0x5a, 0x8d, 0x30, 0x9f, 0xc3, 0x4e, 0x7c, 0x5b, 0xff, 0xa3, 0x46, 0xea, 0xe3, 0x25, 0xf0, 0xbd,
	0x9d, 0xb7, 0x5d, 0x7a, 0x8f, 0xcc, 0x95, 0x2a, 0x25, 0x06, 0xc0, 0x6a, 0x6b, 0xb5, 0x8d, 0x99,
	0xce, 0x6c, 0x81, 0xb5, 0xc5, 0x00, 0xe8, 0x5d, 0x32, 0x3d, 0x34, 0x49, 0xa0, 0xff, 0x87, 0xf4,
	0xd4, 0xd0, 0x24, 0x48, 0xbd, 0x43, 0xa6, 0xd0, 0x4a, 0xc6, 0xec, 0xc6, 0x5a, 0x6d, 0x63, 0xbe,
	0x73, 0xcb, 0x7f, 0x6d, 0xc5, 0x94, 0x91, 0x29, 0x11, 0xc7, 0x06, 0xac, 0x65, 0x37, 0x43, 0x48,
	0xf1, 0x75, 0xfd, 0xcf, 0xda, 0x64, 0x27, 0xb8, 0x11, 0xaa, 0x0f, 0xbe, 0x0e, 0x7c, 0xe0, 0x99,
	0x81, 0x44, 0x9e, 0x97, 0x75, 0x20, 0x76, 0x84, 0x10, 0x5d, 0x25, 0x24, 0x48, 0x06, 0xc2, 0x9e,
	0x15, 0x95, 0xcc, 0x20, 0x72, 0x20, 0xec, 0x19, 0xa5, 0xe4, 0x66, 0xa4, 0xad, 0x2b, 0x0a, 0xc1,
	0x67, 0xfa, 0x21, 0x59, 0x10, 0xf1, 0x10, 0x8c, 0x93, 0x16, 0x78, 0x92, 0x8a, 0x3e, 0x56, 0x33,
	0xdd, 0x99, 0x1f, 0xa1, 0xbb, 0xa9, 0xe8, 0xaf, 0xff, 0x7d, 0x9b, 0xcc, 0x8d, 0xd7, 0x44, 0xef,
	0x93, 0x85, 0xe2, 0x5c, 0xdc, 0x3a, 0x23, 0x55, 0x9f, 0x3d, 0xc6, 0x74, 0x73, 0xe1, 0x78, 0x5d,
	0xc4, 0xe8, 0x23, 0xb2, 0xdc, 0x13, 0xd1, 0x59, 0x4f, 0x2b, 0x08, 0x47, 0x11, 0x91, 0x93, 0x43,
	0x60, 0x4f, 0x30, 0x07, 0x2d, 0xb9, 0xa6, 0x01, 0xd1, 0x44, 0xc6, 0x47, 0x04, 0x5f, 0xe5, 0xc0,
	0x24, 0x22, 0x02, 0x1e, 0xe9, 0x5c, 0x39, 0xb6, 0x89, 0x35, 0x53, 0x74, 0x2f, 0xa9, 0x2d, 0xcf,
	0xd0, 0x15, 0x32, 0x83, 0x11, 0xd6, 0xe5, 0x3d, 0xf6, 0x39, 0x1a, 0x4f, 0x7b, 0xa0, 0xeb, 0xf2,
	0x1e, 0xfd, 0x88, 0x2c, 0x22, 0xe9, 0xb4, 0x13, 0x69, 0x90, 0x7c, 0x51, 0x9c, 0xcf, 0x80, 0x38,
	0xf6, 0x28, 0xea, 0x1e, 0x90, 0xba, 0x27, 0x79, 0x0c, 0x89, 0xc8, 0x53, 0xc7, 0xb1, 0x4f, 0x4f,
	0x31, 0xe7, 0xa2, 0x27, 0xb6, 0x03, 0xbe, 0xe5, 0x5b, 0x56, 0x26, 0x54, 0xd6, 0x0a, 0xf6, 0xac,
	0x4a, 0xd8, 0xb6, 0x56, 0xd0, 0x4d, 0xf2, 0x7f, 0x8f, 0x73, 0xa5, 0xb9, 0x81, 0x58, 0xfa, 0xde,
	0xf4, 0x72, 0x3f, 0x78, 0xec, 0x4b, 0x54, 0x2e, 0x7b, 0xb6, 0xad, 0x3b, 0x13, 0x9c, 0x7f, 0x0b,
	0x18, 0xe5, 0x8c, 0x50, 0x36, 0x15, 0x0e, 0xd8, 0x57, 0xa1, 0x4a, 0x8f, 0x1e, 0x97, 0xa0, 0x1f,
	0x01, 0x94, 0x15, 0x55, 0xb2, 0xaf, 0x51, 0x34, 0xeb, 0xb1, 0xa2, 0x40, 0x3f, 0x02, 0x0e, 0x38,
	0x28, 0xd1, 0x4b, 0x21, 0x66, 0xdf, 0xa0, 0x60, 0xc6, 0xc1, 0x4e, 0x00, 0x68, 0x83, 0x2c, 0x39,
	0xe0, 0x4e, 0x67, 0x3a, 0xd5, 0xfd, 0x0b, 0x3e, 0x04, 0x63, 0x7d, 0x6d, 0xcf, 0xf1, 0xa4, 0x75,
	0x07, 0xc7, 0x05, 0xf3, 0x26, 0x10, 0x3e, 0x23, 0x9c, 0x3b, 0x30, 0x4a, 0xa4, 0x5c, 0xe7, 0x8e,
	0x7d, 0x1b, 0x32, 0x96, 0xd8, 0x61, 0x8e, 0x19, 0x6d, 0x3e, 0x18, 0x08, 0x73, 0xc1, 0xa5, 0x62,
	0x2f, 0x42, 0xc6, 0x02, 0x69, 0x29, 0xfa, 0x31, 0x59, 0xb4, 0xd0, 0x1f, 0x80, 0x72, 0xdc, 0xe8,
	0xdc, 0xf9, 0x49, 0xf9, 0x0e, 0xb3, 0x2d, 0x14, 0x70, 0x27, 0xa0, 0xf4, 0x13, 0x52, 0xb7, 0x06,
	0x87, 0x29, 0x72, 0xdc, 0x8f, 0x5a, 0x24, 0x32, 0xd6, 0x44, 0xbb, 0x05, 0x6b, 0xba, 0x88, 0x77,
	0xb3, 0x64, 0x4b, 0x64, 0xf4, 0x39, 0x59, 0x99, 0x90, 0x5a, 0x19, 0x5b, 0x2e, 0x86, 0x42, 0xa6,
	0xfe, 0x94, 0xec, 0x25, 0x06, 0xb1, 0xb1, 0x20, 0x2f, 0x68, 0x96, 0x3c, 0x6d, 0x92, 0x55, 0x6b,
	0xf8, 0x40, 0x46, 0x46, 0xa7, 0x5a, 0x67, 0x5c, 0x0c, 0xb5, 0x8c, 0x85, 0x8a, 0xa0, 0x1c, 0xcf,
	0x2d, 0x34, 0x78, 0xd7, 0x9a, 0x83, 0x52, 0xd3, 0x2c, 0x25, 0xc5, 0x98, 0xbe, 0x26, 0xf7, 0xae,
	0xb1, 0x80, 0xa1, 0x3f, 0xaa, 0xbb, 0xc8, 0x80, 0x6d, 0xe3, 0x2f, 0x62, 0xf5, 0x2a, 0x9b, 0x1d,
	0xaf, 0x3a, 0xbe, 0xc8, 0x80, 0xb6, 0xc8, 0xfa, 0x35, 0x4e, 0x0a, 0x84, 0xe1, 0xa0, 0x62, 0xbf,
	0x3b, 0x76, 0xae, 0xb7, 0x6a, 0x83, 0x30, 0x3b, 0x2a, 0x6e, 0xc5, 0xf4, 0xd5, 0xb5, 0x45, 0x25,
	0x95, 0xd3, 0x2e, 0x3a, 0xbd, 0x77, 0x95, 0xd3, 0x6e, 0x69, 0xd4, 0x26, 0xf7, 0xaf, 0x31, 0xca,
	0x2c, 0xe4, 0xb1, 0x56, 0x3a, 0x06, 0xef, 0xf5, 0x0a, 0xbd, 0xd6, 0xae, 0xf2, 0x3a, 0x1a, 0x09,
	0x5b, 0x31, 0x7d, 0x48, 0x96, 0x44, 0xee, 0x4e, 0x40, 0x39, 0x19, 0xe1, 0x16, 0x0e, 0xfd, 0x79,
	0x8d, 0xe1, 0x74, 0x92, 0xc2, 0xa6, 0xac, 0x90, 0x19, 0x9c, 0x00, 0xfc, 0xe9, 0xb7, 0x70, 0x5c,
	0xa6, 0x6d, 0x96, 0x84, 0x1f, 0x7c, 0xb9, 0x7a, 0x32, 0x9d, 0xca, 0x08, 0x87, 0xee, 0x7b, 0x7c,
	0x5f, 0xb8, 0x7a, 0x8e, 0x10, 0x6c, 0x29, 0xfa, 0x29, 0x59, 0x9a, 0x54, 0x85, 0xf5, 0xbc, 0x87,
	0x39, 0x6f, 0x8f, 0x4b, 0x71, 0x4f, 0x97, 0x8b, 0xa2, 0x90, 0xfb, 0x59, 0xdf, 0xaf, 0x16, 0x45,
	0x90, 0xfa, 0x69, 0x7f, 0x58, 0xec, 0xa7, 0x4a, 0x17, 0x7c, 0x0f, 0xd0, 0xb7, 0x3e, 0x21, 0x46,
	0xe3, 0xbf, 0x6a, 0x84, 0x54, 0x5b, 0x9c, 0xb5, 0xd7, 0x6e, 0x6c, 0xcc, 0x3e, 0xfe, 0xbd, 0xf1,
	0xdf, 0x5d, 0x6b, 0x8d, 0xcb, 0x77, 0x49, 0x07, 0x17, 0x58, 0x07, 0xaf, 0x95, 0xb2, 0x9b, 0xa9,
	0x15, 0x45, 0xbf, 0x0f, 0xb1, 0xdf, 0xd8, 0xcd, 0x7d, 0x2b, 0x42, 0xcf, 0x1f, 0x90, 0x7a, 0xa5,
	0x3a, 0x81, 0xe8, 0xcc, 0xe6, 0x03, 0x76, 0x14, 0xf6, 0x63, 0x29, 0x2c, 0x60, 0xfa, 0x19, 0xb9,
	0x83, 0x5a, 0x9d, 0x89, 0xdf, 0x72, 0x18, 0x33, 0xfe, 0xa1, 0xda, 0xe1, 0x87, 0xc8, 0x8d, 0xec,
	0x9f, 0x12, 0x76, 0x29, 0xa4, 0xcc, 0xd2, 0xc1, 0xa8, 0x3b, 0x93, 0x51, 0x65, 0xae, 0x67, 0xe4,
	0x2e, 0x06, 0xc6, 0x11, 0xef, 0x49, 0x97, 0xfa, 0x33, 0x57, 0xf9, 0xba, 0x55, 0xe4, 0x76, 0xf4,
	0x32, 0xd0, 0xa3, 0x94, 0x8f, 0xc8, 0xb2, 0x54, 0x71, 0x39, 0x8f, 0x55, 0xd0, 0x71, 0x28, 0xb2,
	0xe2, 0x46, 0x11, 0xeb, 0x64, 0x3e, 0x56, 0xe3, 0x8d, 0xfa, 0x11, 0xa5, 0xb3, 0xb1, 0x9a, 0xe8,
	0x53, 0x92, 0x6a, 0x1d, 0xf3, 0x54, 0x5a, 0xc7, 0x53, 0x50, 0x7d, 0x77, 0xc2, 0xde, 0x84, 0x3e,
	0x21, 0xb1, 0x2f, 0xad, 0xdb, 0x47, 0x78, 0x74, 0xe8, 0x34, 0xb9, 0x7c, 0xdd, 0xfd, 0x54, 0x95,
	0xbe, 0x9f, 0xfc, 0xfb, 0xc6, 0xdb, 0x23, 0x1f, 0x84, 0x19, 0x04, 0x53, 0xfc, 0x19, 0xb8, 0xd2,
	0xe3, 0x67, 0xf4, 0x78, 0x1f, 0x47, 0x12, 0x4c, 0xf8, 0x8b, 0x70, 0xd9, 0x6c, 0xf4, 0x66, 0x13,
	0xc1, 0x0d, 0x0c, 0x25, 0xde, 0x07, 0x6f, 0xc7, 0xde, 0x6c, 0x22, 0x3a, 0x05, 0x4c, 0x5f, 0x90,
	0xd5, 0x70, 0x8b, 0xc7, 0xa7, 0xdc, 0x3a, 0xd1, 0xe7, 0x2a, 0x1f, 0x70, 0xd5, 0x33, 0xdc, 0xcf,
	0xa1, 0xdf, 0xec, 0xbf, 0x60, 0x1c, 0x1e, 0xab, 0x19, 0x9f, 0x76, 0x9d, 0xe8, 0xb7, 0xf3, 0x41,
	0xbb, 0x67, 0x76, 0x03, 0xef, 0x77, 0x7c, 0xb8, 0x3a, 0xcb, 0xb8, 0x3c, 0x4d, 0xd9, 0xaf, 0xe1,
	0x3a, 0xc0, 0x2b, 0x34, 0xa8, 0xf3, 0x34, 0xed, 0xdd, 0xc2, 0xff, 0x7f, 0x4f, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x29, 0x03, 0xf5, 0x81, 0x15, 0x0a, 0x00, 0x00,
}
