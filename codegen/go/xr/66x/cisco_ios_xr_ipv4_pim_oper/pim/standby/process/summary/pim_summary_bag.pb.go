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
// source: pim_summary_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_process_summary

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

type PimSummaryBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimSummaryBag_KEYS) Reset()         { *m = PimSummaryBag_KEYS{} }
func (m *PimSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimSummaryBag_KEYS) ProtoMessage()    {}
func (*PimSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43a4755d2ad21ff, []int{0}
}

func (m *PimSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *PimSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimSummaryBag_KEYS.Merge(m, src)
}
func (m *PimSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimSummaryBag_KEYS.Size(m)
}
func (m *PimSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimSummaryBag_KEYS proto.InternalMessageInfo

type PimSummaryBag struct {
	RouteLimit                       uint32   `protobuf:"varint,50,opt,name=route_limit,json=routeLimit,proto3" json:"route_limit,omitempty"`
	RouteCount                       uint32   `protobuf:"varint,51,opt,name=route_count,json=routeCount,proto3" json:"route_count,omitempty"`
	RouteLowWaterMark                uint32   `protobuf:"varint,52,opt,name=route_low_water_mark,json=routeLowWaterMark,proto3" json:"route_low_water_mark,omitempty"`
	IsRouteLimitReached              bool     `protobuf:"varint,53,opt,name=is_route_limit_reached,json=isRouteLimitReached,proto3" json:"is_route_limit_reached,omitempty"`
	TopologyInterfaceStateLimit      uint32   `protobuf:"varint,54,opt,name=topology_interface_state_limit,json=topologyInterfaceStateLimit,proto3" json:"topology_interface_state_limit,omitempty"`
	TopologyInterfaceStateCount      uint32   `protobuf:"varint,55,opt,name=topology_interface_state_count,json=topologyInterfaceStateCount,proto3" json:"topology_interface_state_count,omitempty"`
	RxiLowWaterMark                  uint32   `protobuf:"varint,56,opt,name=rxi_low_water_mark,json=rxiLowWaterMark,proto3" json:"rxi_low_water_mark,omitempty"`
	RxiLimitReached                  bool     `protobuf:"varint,57,opt,name=rxi_limit_reached,json=rxiLimitReached,proto3" json:"rxi_limit_reached,omitempty"`
	RegisterLimit                    uint32   `protobuf:"varint,58,opt,name=register_limit,json=registerLimit,proto3" json:"register_limit,omitempty"`
	RegisterCount                    uint32   `protobuf:"varint,59,opt,name=register_count,json=registerCount,proto3" json:"register_count,omitempty"`
	RegisterLimitReached             bool     `protobuf:"varint,60,opt,name=register_limit_reached,json=registerLimitReached,proto3" json:"register_limit_reached,omitempty"`
	RangesLimit                      uint32   `protobuf:"varint,61,opt,name=ranges_limit,json=rangesLimit,proto3" json:"ranges_limit,omitempty"`
	RangesCount                      uint32   `protobuf:"varint,62,opt,name=ranges_count,json=rangesCount,proto3" json:"ranges_count,omitempty"`
	RangesThreshold                  uint32   `protobuf:"varint,63,opt,name=ranges_threshold,json=rangesThreshold,proto3" json:"ranges_threshold,omitempty"`
	IsRangesLimitReached             bool     `protobuf:"varint,64,opt,name=is_ranges_limit_reached,json=isRangesLimitReached,proto3" json:"is_ranges_limit_reached,omitempty"`
	BsrRangesLimit                   uint32   `protobuf:"varint,65,opt,name=bsr_ranges_limit,json=bsrRangesLimit,proto3" json:"bsr_ranges_limit,omitempty"`
	BsrRangesCount                   uint32   `protobuf:"varint,66,opt,name=bsr_ranges_count,json=bsrRangesCount,proto3" json:"bsr_ranges_count,omitempty"`
	BsrRangeThreshold                uint32   `protobuf:"varint,67,opt,name=bsr_range_threshold,json=bsrRangeThreshold,proto3" json:"bsr_range_threshold,omitempty"`
	IsBsrRangesThresholdReached      bool     `protobuf:"varint,68,opt,name=is_bsr_ranges_threshold_reached,json=isBsrRangesThresholdReached,proto3" json:"is_bsr_ranges_threshold_reached,omitempty"`
	BsrCandidateRpSetLimit           uint32   `protobuf:"varint,69,opt,name=bsr_candidate_rp_set_limit,json=bsrCandidateRpSetLimit,proto3" json:"bsr_candidate_rp_set_limit,omitempty"`
	BsrCandidateRpSetCount           uint32   `protobuf:"varint,70,opt,name=bsr_candidate_rp_set_count,json=bsrCandidateRpSetCount,proto3" json:"bsr_candidate_rp_set_count,omitempty"`
	BsrCandidateRpSetThreshold       uint32   `protobuf:"varint,71,opt,name=bsr_candidate_rp_set_threshold,json=bsrCandidateRpSetThreshold,proto3" json:"bsr_candidate_rp_set_threshold,omitempty"`
	IsMaximumEnforcementDisabled     bool     `protobuf:"varint,72,opt,name=is_maximum_enforcement_disabled,json=isMaximumEnforcementDisabled,proto3" json:"is_maximum_enforcement_disabled,omitempty"`
	IsNodeLowMemory                  bool     `protobuf:"varint,73,opt,name=is_node_low_memory,json=isNodeLowMemory,proto3" json:"is_node_low_memory,omitempty"`
	RouteThreshold                   uint32   `protobuf:"varint,74,opt,name=route_threshold,json=routeThreshold,proto3" json:"route_threshold,omitempty"`
	GlobalAutoRpRangesLimit          uint32   `protobuf:"varint,75,opt,name=global_auto_rp_ranges_limit,json=globalAutoRpRangesLimit,proto3" json:"global_auto_rp_ranges_limit,omitempty"`
	IsGlobalAutoRpRangesLimitReached bool     `protobuf:"varint,76,opt,name=is_global_auto_rp_ranges_limit_reached,json=isGlobalAutoRpRangesLimitReached,proto3" json:"is_global_auto_rp_ranges_limit_reached,omitempty"`
	GlobalBsrRangesLimit             uint32   `protobuf:"varint,77,opt,name=global_bsr_ranges_limit,json=globalBsrRangesLimit,proto3" json:"global_bsr_ranges_limit,omitempty"`
	GlobalBsrRangesCount             uint32   `protobuf:"varint,78,opt,name=global_bsr_ranges_count,json=globalBsrRangesCount,proto3" json:"global_bsr_ranges_count,omitempty"`
	GlobalBsrRangesThreshold         uint32   `protobuf:"varint,79,opt,name=global_bsr_ranges_threshold,json=globalBsrRangesThreshold,proto3" json:"global_bsr_ranges_threshold,omitempty"`
	IsGlobalBsrRangesLimitReached    bool     `protobuf:"varint,80,opt,name=is_global_bsr_ranges_limit_reached,json=isGlobalBsrRangesLimitReached,proto3" json:"is_global_bsr_ranges_limit_reached,omitempty"`
	GlobalBsrCandidateRpSetLimit     uint32   `protobuf:"varint,81,opt,name=global_bsr_candidate_rp_set_limit,json=globalBsrCandidateRpSetLimit,proto3" json:"global_bsr_candidate_rp_set_limit,omitempty"`
	GlobalBsrCandidateRpSetCount     uint32   `protobuf:"varint,82,opt,name=global_bsr_candidate_rp_set_count,json=globalBsrCandidateRpSetCount,proto3" json:"global_bsr_candidate_rp_set_count,omitempty"`
	GlobalBsrCandidateRpSetThreshold uint32   `protobuf:"varint,83,opt,name=global_bsr_candidate_rp_set_threshold,json=globalBsrCandidateRpSetThreshold,proto3" json:"global_bsr_candidate_rp_set_threshold,omitempty"`
	IsGlobalRouteLimitReached        bool     `protobuf:"varint,84,opt,name=is_global_route_limit_reached,json=isGlobalRouteLimitReached,proto3" json:"is_global_route_limit_reached,omitempty"`
	TopologyInterfaceStateThreshold  uint32   `protobuf:"varint,85,opt,name=topology_interface_state_threshold,json=topologyInterfaceStateThreshold,proto3" json:"topology_interface_state_threshold,omitempty"`
	IsGlobalRxiLimitReached          bool     `protobuf:"varint,86,opt,name=is_global_rxi_limit_reached,json=isGlobalRxiLimitReached,proto3" json:"is_global_rxi_limit_reached,omitempty"`
	RegisterThreshold                uint32   `protobuf:"varint,87,opt,name=register_threshold,json=registerThreshold,proto3" json:"register_threshold,omitempty"`
	GlobalRegisterLimit              uint32   `protobuf:"varint,88,opt,name=global_register_limit,json=globalRegisterLimit,proto3" json:"global_register_limit,omitempty"`
	IsGlobalRegisterLimitReached     bool     `protobuf:"varint,89,opt,name=is_global_register_limit_reached,json=isGlobalRegisterLimitReached,proto3" json:"is_global_register_limit_reached,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *PimSummaryBag) Reset()         { *m = PimSummaryBag{} }
func (m *PimSummaryBag) String() string { return proto.CompactTextString(m) }
func (*PimSummaryBag) ProtoMessage()    {}
func (*PimSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c43a4755d2ad21ff, []int{1}
}

func (m *PimSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimSummaryBag.Unmarshal(m, b)
}
func (m *PimSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimSummaryBag.Marshal(b, m, deterministic)
}
func (m *PimSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimSummaryBag.Merge(m, src)
}
func (m *PimSummaryBag) XXX_Size() int {
	return xxx_messageInfo_PimSummaryBag.Size(m)
}
func (m *PimSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimSummaryBag proto.InternalMessageInfo

func (m *PimSummaryBag) GetRouteLimit() uint32 {
	if m != nil {
		return m.RouteLimit
	}
	return 0
}

func (m *PimSummaryBag) GetRouteCount() uint32 {
	if m != nil {
		return m.RouteCount
	}
	return 0
}

func (m *PimSummaryBag) GetRouteLowWaterMark() uint32 {
	if m != nil {
		return m.RouteLowWaterMark
	}
	return 0
}

func (m *PimSummaryBag) GetIsRouteLimitReached() bool {
	if m != nil {
		return m.IsRouteLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetTopologyInterfaceStateLimit() uint32 {
	if m != nil {
		return m.TopologyInterfaceStateLimit
	}
	return 0
}

func (m *PimSummaryBag) GetTopologyInterfaceStateCount() uint32 {
	if m != nil {
		return m.TopologyInterfaceStateCount
	}
	return 0
}

func (m *PimSummaryBag) GetRxiLowWaterMark() uint32 {
	if m != nil {
		return m.RxiLowWaterMark
	}
	return 0
}

func (m *PimSummaryBag) GetRxiLimitReached() bool {
	if m != nil {
		return m.RxiLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetRegisterLimit() uint32 {
	if m != nil {
		return m.RegisterLimit
	}
	return 0
}

func (m *PimSummaryBag) GetRegisterCount() uint32 {
	if m != nil {
		return m.RegisterCount
	}
	return 0
}

func (m *PimSummaryBag) GetRegisterLimitReached() bool {
	if m != nil {
		return m.RegisterLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetRangesLimit() uint32 {
	if m != nil {
		return m.RangesLimit
	}
	return 0
}

func (m *PimSummaryBag) GetRangesCount() uint32 {
	if m != nil {
		return m.RangesCount
	}
	return 0
}

func (m *PimSummaryBag) GetRangesThreshold() uint32 {
	if m != nil {
		return m.RangesThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsRangesLimitReached() bool {
	if m != nil {
		return m.IsRangesLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetBsrRangesLimit() uint32 {
	if m != nil {
		return m.BsrRangesLimit
	}
	return 0
}

func (m *PimSummaryBag) GetBsrRangesCount() uint32 {
	if m != nil {
		return m.BsrRangesCount
	}
	return 0
}

func (m *PimSummaryBag) GetBsrRangeThreshold() uint32 {
	if m != nil {
		return m.BsrRangeThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsBsrRangesThresholdReached() bool {
	if m != nil {
		return m.IsBsrRangesThresholdReached
	}
	return false
}

func (m *PimSummaryBag) GetBsrCandidateRpSetLimit() uint32 {
	if m != nil {
		return m.BsrCandidateRpSetLimit
	}
	return 0
}

func (m *PimSummaryBag) GetBsrCandidateRpSetCount() uint32 {
	if m != nil {
		return m.BsrCandidateRpSetCount
	}
	return 0
}

func (m *PimSummaryBag) GetBsrCandidateRpSetThreshold() uint32 {
	if m != nil {
		return m.BsrCandidateRpSetThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsMaximumEnforcementDisabled() bool {
	if m != nil {
		return m.IsMaximumEnforcementDisabled
	}
	return false
}

func (m *PimSummaryBag) GetIsNodeLowMemory() bool {
	if m != nil {
		return m.IsNodeLowMemory
	}
	return false
}

func (m *PimSummaryBag) GetRouteThreshold() uint32 {
	if m != nil {
		return m.RouteThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalAutoRpRangesLimit() uint32 {
	if m != nil {
		return m.GlobalAutoRpRangesLimit
	}
	return 0
}

func (m *PimSummaryBag) GetIsGlobalAutoRpRangesLimitReached() bool {
	if m != nil {
		return m.IsGlobalAutoRpRangesLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetGlobalBsrRangesLimit() uint32 {
	if m != nil {
		return m.GlobalBsrRangesLimit
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalBsrRangesCount() uint32 {
	if m != nil {
		return m.GlobalBsrRangesCount
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalBsrRangesThreshold() uint32 {
	if m != nil {
		return m.GlobalBsrRangesThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsGlobalBsrRangesLimitReached() bool {
	if m != nil {
		return m.IsGlobalBsrRangesLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetGlobalBsrCandidateRpSetLimit() uint32 {
	if m != nil {
		return m.GlobalBsrCandidateRpSetLimit
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalBsrCandidateRpSetCount() uint32 {
	if m != nil {
		return m.GlobalBsrCandidateRpSetCount
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalBsrCandidateRpSetThreshold() uint32 {
	if m != nil {
		return m.GlobalBsrCandidateRpSetThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsGlobalRouteLimitReached() bool {
	if m != nil {
		return m.IsGlobalRouteLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetTopologyInterfaceStateThreshold() uint32 {
	if m != nil {
		return m.TopologyInterfaceStateThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetIsGlobalRxiLimitReached() bool {
	if m != nil {
		return m.IsGlobalRxiLimitReached
	}
	return false
}

func (m *PimSummaryBag) GetRegisterThreshold() uint32 {
	if m != nil {
		return m.RegisterThreshold
	}
	return 0
}

func (m *PimSummaryBag) GetGlobalRegisterLimit() uint32 {
	if m != nil {
		return m.GlobalRegisterLimit
	}
	return 0
}

func (m *PimSummaryBag) GetIsGlobalRegisterLimitReached() bool {
	if m != nil {
		return m.IsGlobalRegisterLimitReached
	}
	return false
}

func init() {
	proto.RegisterType((*PimSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.process.summary.pim_summary_bag_KEYS")
	proto.RegisterType((*PimSummaryBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.process.summary.pim_summary_bag")
}

func init() { proto.RegisterFile("pim_summary_bag.proto", fileDescriptor_c43a4755d2ad21ff) }

var fileDescriptor_c43a4755d2ad21ff = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x6d, 0x73, 0xdb, 0x44,
	0x10, 0x9e, 0x7e, 0x61, 0x98, 0x2d, 0x6d, 0x1a, 0xc5, 0x75, 0x44, 0xd3, 0x36, 0xae, 0x67, 0x0a,
	0x01, 0x06, 0x33, 0xd3, 0x24, 0x05, 0x4a, 0x0b, 0x8d, 0x1d, 0xc7, 0x84, 0xc4, 0x49, 0x90, 0x03,
	0x21, 0x9f, 0x6e, 0xf4, 0x72, 0x71, 0x6e, 0x62, 0xe9, 0x34, 0x77, 0x32, 0xb1, 0xff, 0x24, 0xbf,
	0x89, 0xd1, 0xad, 0x4e, 0x27, 0xc9, 0xb2, 0xfb, 0xf5, 0xee, 0xd9, 0xe7, 0x65, 0x6f, 0xad, 0x35,
	0x3c, 0x8d, 0x59, 0x48, 0xe4, 0x34, 0x0c, 0x5d, 0x31, 0x27, 0x9e, 0x3b, 0xee, 0xc4, 0x82, 0x27,
	0xdc, 0x7a, 0xeb, 0x33, 0xe9, 0x73, 0xc2, 0xb8, 0x24, 0x33, 0x41, 0x58, 0xfc, 0xef, 0x1e, 0x49,
	0x81, 0x3c, 0xa6, 0xa2, 0x13, 0xb3, 0xb0, 0x23, 0x13, 0x37, 0x0a, 0xbc, 0x79, 0x8a, 0xf6, 0xa9,
	0x94, 0x9d, 0x8c, 0xa1, 0xdd, 0x84, 0x46, 0x85, 0x90, 0x9c, 0xf4, 0xaf, 0x47, 0xed, 0xff, 0xd6,
	0x61, 0xad, 0x72, 0x61, 0x6d, 0xc3, 0x43, 0xc1, 0xa7, 0x09, 0x25, 0x13, 0x16, 0xb2, 0xc4, 0x7e,
	0xd3, 0x7a, 0xb0, 0xf3, 0xc8, 0x01, 0x75, 0x74, 0x9a, 0x9e, 0x18, 0x80, 0xcf, 0xa7, 0x51, 0x62,
	0xef, 0x16, 0x00, 0xbd, 0xf4, 0xc4, 0xfa, 0x01, 0x1a, 0x19, 0x03, 0xbf, 0x27, 0xf7, 0x6e, 0x42,
	0x05, 0x09, 0x5d, 0x71, 0x67, 0xef, 0x29, 0xe4, 0x3a, 0x52, 0xf1, 0xfb, 0xab, 0xf4, 0x66, 0xe8,
	0x8a, 0x3b, 0x6b, 0x17, 0x9a, 0x4c, 0x92, 0x82, 0x2a, 0x11, 0xd4, 0xf5, 0x6f, 0x69, 0x60, 0xef,
	0xb7, 0x1e, 0xec, 0x7c, 0xee, 0x6c, 0x30, 0xe9, 0xe4, 0xfa, 0x0e, 0x5e, 0x59, 0x3d, 0x78, 0x99,
	0xf0, 0x98, 0x4f, 0xf8, 0x78, 0x4e, 0x58, 0x94, 0x50, 0x71, 0xe3, 0xfa, 0x94, 0xc8, 0xc4, 0xcd,
	0xad, 0xbf, 0x55, 0x7a, 0x5b, 0x1a, 0x75, 0xac, 0x41, 0xa3, 0x14, 0x83, 0x59, 0x56, 0x91, 0x60,
	0xbc, 0x1f, 0x57, 0x91, 0x60, 0xde, 0xef, 0xc0, 0x12, 0x33, 0x56, 0x4d, 0xfb, 0x93, 0x2a, 0x5c,
	0x13, 0x33, 0x56, 0xca, 0xfa, 0x2d, 0xac, 0x2b, 0x70, 0x29, 0xe6, 0xcf, 0x2a, 0xa6, 0xc2, 0x16,
	0x23, 0xbe, 0x86, 0xc7, 0x82, 0x8e, 0x99, 0x4c, 0x39, 0x31, 0xd2, 0x3b, 0x45, 0xfa, 0x48, 0x9f,
	0x62, 0x88, 0x22, 0x0c, 0x4d, 0xff, 0x52, 0x86, 0xa1, 0xcd, 0x3d, 0x68, 0x96, 0xd9, 0x72, 0xf9,
	0xf7, 0x4a, 0xbe, 0x51, 0x62, 0xd5, 0x1e, 0x5e, 0xc1, 0x17, 0xc2, 0x8d, 0xc6, 0x54, 0x66, 0x0e,
	0x3e, 0x28, 0xea, 0x87, 0x78, 0x86, 0xfa, 0x06, 0x82, 0xea, 0xbf, 0x16, 0x21, 0xa8, 0xfd, 0x0d,
	0x3c, 0xc9, 0x20, 0xc9, 0xad, 0xa0, 0xf2, 0x96, 0x4f, 0x02, 0xfb, 0xb7, 0xac, 0x41, 0xea, 0xfc,
	0x52, 0x1f, 0x5b, 0xfb, 0xb0, 0x99, 0x0e, 0x43, 0x41, 0x33, 0xf7, 0xf9, 0x11, 0x7d, 0x32, 0xe9,
	0x18, 0x75, 0xed, 0x73, 0x07, 0x9e, 0x78, 0x52, 0x94, 0xea, 0xec, 0x03, 0xa5, 0xf0, 0xd8, 0x93,
	0xa2, 0x50, 0x50, 0x41, 0xa2, 0xe5, 0x6e, 0x05, 0x89, 0xae, 0x3b, 0xb0, 0x91, 0x23, 0x0b, 0xc6,
	0x7b, 0x38, 0xc7, 0x1a, 0x6c, 0xac, 0x1f, 0xc2, 0x36, 0x93, 0xa4, 0x40, 0x9e, 0xd7, 0xe4, 0x11,
	0x0e, 0x55, 0x84, 0x2d, 0x26, 0xbb, 0x5a, 0x2a, 0x2f, 0xd7, 0x49, 0xde, 0xc1, 0xb3, 0x94, 0xc2,
	0x77, 0xa3, 0x80, 0x05, 0xe9, 0x20, 0x8a, 0x98, 0x48, 0x9a, 0x64, 0x99, 0xfa, 0x4a, 0xbc, 0xe9,
	0x49, 0xd1, 0xd3, 0x00, 0x27, 0x1e, 0xd1, 0x04, 0xb3, 0x2d, 0xab, 0xc5, 0x94, 0x47, 0x4b, 0x6a,
	0x31, 0x6d, 0x17, 0x5e, 0xd6, 0xd6, 0x9a, 0xe0, 0x03, 0x55, 0xff, 0x6c, 0xa1, 0xde, 0x74, 0xa0,
	0xaf, 0x3a, 0x10, 0xba, 0x33, 0x16, 0x4e, 0x43, 0x42, 0xa3, 0x1b, 0x2e, 0x7c, 0x1a, 0xd2, 0x28,
	0x21, 0x01, 0x93, 0xae, 0x37, 0xa1, 0x81, 0xfd, 0xbb, 0xea, 0xc0, 0x73, 0x26, 0x87, 0x88, 0xea,
	0x1b, 0xd0, 0x61, 0x86, 0x49, 0x7f, 0x51, 0x4c, 0x92, 0x88, 0x07, 0xf8, 0x0d, 0x09, 0x69, 0xc8,
	0xc5, 0xdc, 0x3e, 0xc6, 0x5f, 0x09, 0x93, 0x67, 0x3c, 0x48, 0x3f, 0x20, 0x43, 0x75, 0x6c, 0x7d,
	0x0d, 0x6b, 0xf8, 0xe9, 0x30, 0x46, 0xff, 0xc0, 0xe7, 0x54, 0xc7, 0xc6, 0xdc, 0x7b, 0xd8, 0x1a,
	0x4f, 0xb8, 0xe7, 0x4e, 0x88, 0x3b, 0x4d, 0x78, 0x1a, 0xaf, 0x34, 0x2d, 0x27, 0xaa, 0x68, 0x13,
	0x21, 0x07, 0xd3, 0x84, 0x3b, 0x71, 0x71, 0x6c, 0x2e, 0xe0, 0x2b, 0x26, 0xc9, 0x0a, 0x82, 0xfc,
	0x8d, 0x4f, 0x95, 0xcf, 0x16, 0x93, 0x83, 0x7a, 0x2a, 0xfd, 0xd0, 0xfb, 0x90, 0x89, 0x91, 0x85,
	0xc9, 0x1d, 0x2a, 0x2f, 0x0d, 0xbc, 0xee, 0x96, 0xe7, 0xb7, 0xb6, 0x0c, 0x1f, 0xf8, 0xac, 0xb6,
	0x0c, 0x9f, 0xf7, 0x43, 0x9e, 0xbe, 0x6e, 0x40, 0xed, 0x73, 0x55, 0x6a, 0x57, 0x4a, 0x4d, 0xf3,
	0x8e, 0xa1, 0x6d, 0xe2, 0x57, 0xfd, 0xe6, 0xd1, 0x2f, 0x54, 0xf4, 0x17, 0x3a, 0x7a, 0xd9, 0xb9,
	0xce, 0x3d, 0x80, 0x57, 0x05, 0x9e, 0x25, 0x73, 0xfe, 0xa7, 0xf2, 0xf3, 0x3c, 0xf7, 0x53, 0x37,
	0xed, 0x9f, 0x20, 0xc2, 0x9e, 0x38, 0x2b, 0x89, 0xb0, 0x37, 0xe7, 0xf0, 0x7a, 0x15, 0x91, 0xe9,
	0xd2, 0x48, 0x91, 0xb5, 0x96, 0x90, 0x99, 0x6e, 0x7d, 0x84, 0x17, 0xa6, 0x5b, 0x75, 0x8b, 0xed,
	0x52, 0x35, 0xea, 0x4b, 0xdd, 0xa8, 0xc5, 0xf5, 0x76, 0x02, 0xed, 0xa5, 0x9b, 0xc9, 0xf8, 0xf9,
	0x4b, 0xf9, 0xd9, 0xae, 0xdf, 0x4e, 0xa5, 0xc9, 0x2f, 0xd8, 0x59, 0x58, 0x3f, 0x7f, 0x2b, 0x33,
	0x9b, 0xb9, 0x99, 0xca, 0x1a, 0xfa, 0x1e, 0xac, 0x7c, 0x71, 0x18, 0xe9, 0xab, 0x6c, 0x9b, 0x67,
	0x37, 0x46, 0xec, 0x0d, 0x3c, 0xd5, 0x4a, 0xe5, 0xe5, 0xf5, 0x8f, 0xaa, 0xd8, 0xc0, 0x4b, 0xa7,
	0xb4, 0xc2, 0x8e, 0xa0, 0x55, 0x30, 0x58, 0xbf, 0xa5, 0xae, 0xf5, 0x87, 0x63, 0xb0, 0x48, 0x90,
	0x59, 0xf5, 0x3e, 0x53, 0xff, 0x93, 0x76, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xb7, 0x13,
	0x7d, 0x40, 0x09, 0x00, 0x00,
}
