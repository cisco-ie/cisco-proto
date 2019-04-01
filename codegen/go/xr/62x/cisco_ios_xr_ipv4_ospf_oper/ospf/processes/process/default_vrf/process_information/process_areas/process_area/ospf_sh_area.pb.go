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
	proto.RegisterType((*OspfShArea_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area_KEYS")
	proto.RegisterType((*OspfShAreaRange)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area_range")
	proto.RegisterType((*OspfShArea)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.process_areas.process_area.ospf_sh_area")
}

func init() { proto.RegisterFile("ospf_sh_area.proto", fileDescriptor_c064299b5cec207a) }

var fileDescriptor_c064299b5cec207a = []byte{
	// 1136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x72, 0x1b, 0x45,
	0x13, 0x2d, 0x7d, 0x49, 0x25, 0xd1, 0x38, 0x76, 0xa2, 0xb1, 0xf3, 0x31, 0x60, 0x4c, 0x39, 0x26,
	0x80, 0x49, 0x15, 0x4a, 0x48, 0x0c, 0x09, 0x3f, 0x21, 0x28, 0xfe, 0x49, 0x84, 0x6d, 0xd9, 0x48,
	0x26, 0x10, 0xb8, 0x98, 0x1a, 0xed, 0xf6, 0xca, 0x13, 0xef, 0xce, 0x2c, 0x33, 0xb3, 0x2a, 0xfb,
	0x92, 0x47, 0xe0, 0x11, 0x78, 0x46, 0x5e, 0x80, 0x9a, 0x9e, 0x5d, 0xad, 0x84, 0xed, 0x5b, 0xee,
	0x56, 0xa7, 0x4f, 0x9f, 0x33, 0xdd, 0xdb, 0xea, 0x59, 0x42, 0xb5, 0xcd, 0x13, 0x6e, 0x8f, 0xb9,
	0x30, 0x20, 0xda, 0xb9, 0xd1, 0x4e, 0xd3, 0x2c, 0x92, 0x36, 0xd2, 0x5c, 0x6a, 0xcb, 0x4f, 0x0d,
	0x97, 0xf9, 0x78, 0x83, 0x23, 0x4b, 0xe7, 0x60, 0xda, 0xfe, 0xc9, 0xf3, 0x22, 0xb0, 0x16, 0x6c,
	0xf5, 0xd4, 0x8e, 0x21, 0x11, 0x45, 0xea, 0xf8, 0xd8, 0x4c, 0xa2, 0x5c, 0xaa, 0x44, 0x9b, 0x4c,
	0x38, 0xa9, 0xd5, 0x04, 0xf3, 0x36, 0x76, 0xe6, 0xd7, 0x9a, 0x24, 0xad, 0xe9, 0x43, 0xf0, 0xdd,
	0xed, 0x37, 0x03, 0x7a, 0x97, 0xdc, 0xac, 0x48, 0x4a, 0x64, 0xc0, 0x1a, 0xab, 0x8d, 0xf5, 0x66,
	0x7f, 0xae, 0xc4, 0x7a, 0x22, 0x03, 0xfa, 0x0e, 0xb9, 0x8e, 0x7c, 0x19, 0xb3, 0xff, 0xad, 0x36,
	0xd6, 0xe7, 0xfb, 0xd7, 0xfc, 0xcf, 0x6e, 0x4c, 0x19, 0xb9, 0x2e, 0xe2, 0xd8, 0x80, 0xb5, 0xec,
	0x0a, 0xa6, 0x55, 0x3f, 0xd7, 0xfe, 0x6c, 0xcc, 0x16, 0xcc, 0x8d, 0x50, 0x23, 0xf0, 0x66, 0xf8,
	0xc0, 0x73, 0x03, 0x89, 0x3c, 0xad, 0xcc, 0x10, 0x3b, 0x44, 0x88, 0xae, 0x10, 0x12, 0x28, 0x99,
	0xb0, 0x27, 0xe8, 0xd7, 0xec, 0x37, 0x11, 0xd9, 0x17, 0xf6, 0x84, 0x52, 0x72, 0x35, 0xd2, 0xd6,
	0xa1, 0xdf, 0x7c, 0x1f, 0x9f, 0xe9, 0x47, 0x64, 0x41, 0xc4, 0x63, 0x30, 0x4e, 0x5a, 0xe0, 0x49,
	0x2a, 0x46, 0xec, 0xea, 0x6a, 0x63, 0xfd, 0x46, 0x7f, 0x7e, 0x82, 0xee, 0xa4, 0x62, 0xb4, 0xf6,
	0xf7, 0x6d, 0x72, 0x73, 0xfa, 0x4c, 0xf4, 0x1e, 0x59, 0x28, 0xeb, 0xe2, 0xd6, 0x19, 0xa9, 0x46,
	0xec, 0x11, 0xda, 0xdd, 0x0c, 0xe5, 0x0d, 0x10, 0xa3, 0x0f, 0xc9, 0xd2, 0x50, 0x44, 0x27, 0x43,
	0xad, 0x20, 0x94, 0x22, 0x22, 0x27, 0xc7, 0xc0, 0x1e, 0xa3, 0x07, 0xad, 0x62, 0x1d, 0x03, 0xa2,
	0x83, 0x11, 0x9f, 0x11, 0x74, 0x95, 0x03, 0x93, 0x88, 0x08, 0x78, 0xa4, 0x0b, 0xe5, 0xd8, 0x06,
	0x9e, 0x99, 0xa2, 0x7a, 0x15, 0xda, 0xf4, 0x11, 0xba, 0x4c, 0x9a, 0x98, 0x61, 0x5d, 0x31, 0x64,
	0x5f, 0xa0, 0xf0, 0x0d, 0x0f, 0x0c, 0x5c, 0x31, 0xa4, 0x1f, 0x93, 0x5b, 0x18, 0x74, 0xda, 0x89,
	0x34, 0x50, 0xbe, 0x2c, 0xeb, 0x33, 0x20, 0x8e, 0x3c, 0x8a, 0xbc, 0xfb, 0xa4, 0xe5, 0x83, 0xbc,
	0x9a, 0x10, 0xec, 0xd3, 0x13, 0xf4, 0xbc, 0xe5, 0x03, 0x5b, 0x01, 0xdf, 0xf4, 0x2d, 0xab, 0x0c,
	0x95, 0xb5, 0x82, 0x3d, 0xad, 0x0d, 0x7b, 0xd6, 0x0a, 0xba, 0x41, 0xfe, 0xef, 0x71, 0xae, 0x34,
	0x37, 0x10, 0x4b, 0xdf, 0x9b, 0x61, 0xe1, 0x87, 0x8b, 0x7d, 0x85, 0xcc, 0x25, 0x1f, 0xed, 0xe9,
	0xfe, 0x4c, 0xcc, 0xbf, 0x05, 0xcc, 0x72, 0x46, 0x28, 0x9b, 0x0a, 0x07, 0xec, 0xeb, 0x70, 0x4a,
	0x8f, 0x1e, 0x55, 0xa0, 0x1f, 0x01, 0xa4, 0x95, 0xa7, 0x64, 0xdf, 0x20, 0x69, 0xce, 0x63, 0xe5,
	0x01, 0xfd, 0x08, 0x38, 0xe0, 0xa0, 0xc4, 0x30, 0x85, 0x98, 0x7d, 0x8b, 0x84, 0xa6, 0x83, 0xed,
	0x00, 0xd0, 0x36, 0x59, 0x74, 0xc0, 0x9d, 0xce, 0x75, 0xaa, 0x47, 0x67, 0x7c, 0x0c, 0xc6, 0xfa,
	0xb3, 0x3d, 0xc3, 0x4a, 0x5b, 0x0e, 0x8e, 0xca, 0xc8, 0xeb, 0x10, 0xf0, 0x8e, 0x70, 0xea, 0xc0,
	0x28, 0x91, 0x72, 0x5d, 0x38, 0xf6, 0x5d, 0x70, 0xac, 0xb0, 0x83, 0x02, 0x1d, 0x6d, 0x91, 0x65,
	0xc2, 0x9c, 0x71, 0xa9, 0xd8, 0xf3, 0xe0, 0x58, 0x22, 0x5d, 0x45, 0x3f, 0x21, 0xb7, 0x2c, 0x8c,
	0x32, 0x50, 0x8e, 0x1b, 0x5d, 0x38, 0x3f, 0x29, 0xdf, 0xa3, 0xdb, 0x42, 0x09, 0xf7, 0x03, 0x4a,
	0x3f, 0x25, 0x2d, 0x6b, 0x70, 0x98, 0x22, 0xc7, 0xfd, 0xa8, 0x45, 0x22, 0x67, 0x1d, 0x94, 0x5b,
	0xb0, 0x66, 0x80, 0xf8, 0x20, 0x4f, 0x36, 0x45, 0x4e, 0x9f, 0x91, 0xe5, 0x19, 0xaa, 0x95, 0xb1,
	0xe5, 0x62, 0x2c, 0x64, 0xea, 0xab, 0x64, 0x2f, 0x30, 0x89, 0x4d, 0x25, 0x79, 0x42, 0xa7, 0x8a,
	0xd3, 0x0e, 0x59, 0xb1, 0x86, 0x67, 0x32, 0x32, 0x3a, 0xd5, 0x3a, 0xe7, 0x62, 0xac, 0x65, 0x2c,
	0x54, 0x04, 0xd5, 0x78, 0x6e, 0xa2, 0xc0, 0x7b, 0xd6, 0xec, 0x57, 0x9c, 0x4e, 0x45, 0x29, 0xc7,
	0xf4, 0x15, 0xb9, 0x7b, 0x89, 0x04, 0x8c, 0x7d, 0xa9, 0xee, 0x2c, 0x07, 0xb6, 0x85, 0xff, 0x88,
	0x95, 0x8b, 0x64, 0xb6, 0x3d, 0xeb, 0xe8, 0x2c, 0x07, 0xda, 0x25, 0x6b, 0x97, 0x28, 0x29, 0x10,
	0x86, 0x83, 0x8a, 0xfd, 0xee, 0xd8, 0xbe, 0x5c, 0xaa, 0x07, 0xc2, 0x6c, 0xab, 0xb8, 0x1b, 0xd3,
	0x97, 0x97, 0x1e, 0x2a, 0xa9, 0x95, 0x76, 0x50, 0xe9, 0xfd, 0x8b, 0x94, 0x76, 0x2a, 0xa1, 0x1e,
	0xb9, 0x77, 0x89, 0x50, 0x6e, 0xa1, 0x88, 0xb5, 0xd2, 0x31, 0x78, 0xad, 0x97, 0xa8, 0xb5, 0x7a,
	0x91, 0xd6, 0xe1, 0x84, 0xd8, 0x8d, 0xe9, 0x03, 0xb2, 0x28, 0x0a, 0x77, 0x0c, 0xca, 0xc9, 0x08,
	0x37, 0x6d, 0xe8, 0xcf, 0x2b, 0x4c, 0xa7, 0xb3, 0x21, 0x6c, 0xca, 0x32, 0x69, 0xe2, 0x04, 0xe0,
	0x5f, 0xbf, 0x8b, 0xe3, 0x72, 0xc3, 0xe6, 0x49, 0xf8, 0xc3, 0x57, 0xab, 0x27, 0xd7, 0xa9, 0x8c,
	0x70, 0xe8, 0x7e, 0xc0, 0xf7, 0x85, 0xab, 0xe7, 0x10, 0xc1, 0xae, 0xa2, 0x9f, 0x91, 0xc5, 0x59,
	0x56, 0x58, 0xd1, 0xbb, 0xe8, 0x79, 0x7b, 0x9a, 0x8a, 0x7b, 0xba, 0x5a, 0x14, 0x25, 0xdd, 0xcf,
	0xfa, 0x5e, 0xbd, 0x28, 0x02, 0xd5, 0x4f, 0xfb, 0x83, 0x72, 0x3f, 0xd5, 0xbc, 0xa0, 0xbb, 0x8f,
	0xba, 0xad, 0x19, 0x32, 0x0a, 0xff, 0xd5, 0x20, 0xa4, 0xde, 0xe2, 0xac, 0xb7, 0x7a, 0x65, 0x7d,
	0xee, 0xd1, 0x1f, 0x8d, 0xf6, 0x7f, 0x7a, 0x7d, 0xb5, 0xcf, 0xdf, 0x27, 0x7d, 0x5c, 0x62, 0x7d,
	0xbc, 0x5a, 0xaa, 0x8e, 0xa6, 0x56, 0x94, 0x3d, 0x3f, 0xc0, 0x9e, 0x63, 0x47, 0xf7, 0xac, 0x08,
	0x7d, 0xbf, 0x4f, 0x5a, 0x35, 0xeb, 0x18, 0xa2, 0x13, 0x5b, 0x64, 0xec, 0x30, 0xec, 0xc8, 0x8a,
	0x58, 0xc2, 0xf4, 0x73, 0x72, 0x07, 0xb9, 0x3a, 0x17, 0xbf, 0x17, 0x30, 0x25, 0xfc, 0x63, 0xbd,
	0xc7, 0x0f, 0x30, 0x36, 0x91, 0x7f, 0x42, 0xd8, 0xb9, 0x94, 0xca, 0xa5, 0x8f, 0x59, 0x77, 0x66,
	0xb3, 0x2a, 0xaf, 0xa7, 0xe4, 0x5d, 0x4c, 0x8c, 0x23, 0x3e, 0x94, 0x2e, 0xf5, 0x35, 0xd7, 0x7e,
	0x83, 0x3a, 0x73, 0x2b, 0x7a, 0x11, 0xc2, 0x13, 0xcb, 0x87, 0x64, 0x49, 0xaa, 0xb8, 0x9a, 0xc9,
	0x3a, 0xe9, 0x28, 0x1c, 0xb2, 0x8e, 0x4d, 0x32, 0xd6, 0xc8, 0x7c, 0xac, 0xa6, 0x1b, 0xf5, 0x13,
	0x52, 0xe7, 0x62, 0x35, 0xd3, 0xa7, 0x24, 0xd5, 0x3a, 0xe6, 0xa9, 0xb4, 0x8e, 0xa7, 0xa0, 0x46,
	0xee, 0x98, 0xbd, 0x0e, 0x7d, 0xc2, 0xc0, 0x9e, 0xb4, 0x6e, 0x0f, 0xe1, 0x49, 0xd1, 0x69, 0x72,
	0xfe, 0xca, 0xfb, 0xb9, 0x3e, 0xfa, 0x5e, 0xf2, 0xef, 0x5b, 0x6f, 0x97, 0x7c, 0x18, 0xe6, 0x10,
	0x4c, 0xf9, 0x41, 0x70, 0xa1, 0xc6, 0x2f, 0xa8, 0xf1, 0x01, 0x8e, 0x25, 0x98, 0xf0, 0x99, 0x70,
	0x5e, 0x6c, 0xf2, 0x66, 0x13, 0xc1, 0x0d, 0x8c, 0x25, 0xde, 0x09, 0x6f, 0xa6, 0xde, 0x6c, 0x22,
	0xfa, 0x25, 0x4c, 0x9f, 0x93, 0x95, 0x70, 0x93, 0xc7, 0x6f, 0xb9, 0x75, 0x62, 0xc4, 0x55, 0x91,
	0x71, 0x35, 0x34, 0xdc, 0xcf, 0xa1, 0xdf, 0xee, 0xbf, 0x62, 0x1e, 0x96, 0xd5, 0x89, 0xdf, 0x0e,
	0x9c, 0x18, 0xf5, 0x8a, 0xac, 0x37, 0x34, 0x3b, 0x21, 0xee, 0xf7, 0x7c, 0xb8, 0x3e, 0xab, 0xbc,
	0x22, 0x4d, 0xd9, 0x6f, 0xe1, 0x4a, 0xc0, 0x6b, 0x34, 0xb0, 0x8b, 0x34, 0x1d, 0x5e, 0xc3, 0x4f,
	0xbd, 0xc7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x8e, 0x98, 0x0d, 0x00, 0x0a, 0x00, 0x00,
}
