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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_summary

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
	proto.RegisterType((*PimSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.summary.pim_summary_bag_KEYS")
	proto.RegisterType((*PimSummaryBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.summary.pim_summary_bag")
}

func init() { proto.RegisterFile("pim_summary_bag.proto", fileDescriptor_c43a4755d2ad21ff) }

var fileDescriptor_c43a4755d2ad21ff = []byte{
	// 844 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x6d, 0x73, 0xd3, 0x46,
	0x10, 0x1e, 0xbe, 0x74, 0x3a, 0x4b, 0x21, 0x44, 0x09, 0x8e, 0x4a, 0x80, 0x18, 0xcf, 0xd0, 0xa6,
	0xed, 0xd4, 0x9d, 0x21, 0x49, 0x5f, 0x28, 0x69, 0xc1, 0x8e, 0x71, 0xd3, 0x60, 0xa0, 0x32, 0x2d,
	0xe5, 0xd3, 0xcd, 0x59, 0xba, 0x38, 0x37, 0x58, 0x3a, 0xcd, 0xdd, 0x39, 0x71, 0xfe, 0x64, 0x7f,
	0x53, 0x47, 0xb7, 0x3a, 0x9d, 0x24, 0xcb, 0xe6, 0xeb, 0xdd, 0xb3, 0xcf, 0xcb, 0xde, 0x5a, 0x6b,
	0xb8, 0x9b, 0xf2, 0x98, 0xa8, 0x79, 0x1c, 0x53, 0x79, 0x4d, 0x26, 0x74, 0xda, 0x4d, 0xa5, 0xd0,
	0xc2, 0x3b, 0x0e, 0xb9, 0x0a, 0x05, 0xe1, 0x42, 0x91, 0x85, 0x24, 0x3c, 0xbd, 0x3c, 0x24, 0x19,
	0x50, 0xa4, 0x4c, 0x76, 0x53, 0x1e, 0x77, 0x69, 0xa8, 0xf9, 0x25, 0xeb, 0x46, 0xec, 0x9c, 0xce,
	0x67, 0x9a, 0x84, 0x22, 0xd1, 0x6c, 0xa1, 0xbb, 0x39, 0x51, 0xa7, 0x05, 0xdb, 0x35, 0x5e, 0x72,
	0x36, 0xf8, 0x30, 0xee, 0xfc, 0xb7, 0x09, 0x1b, 0xb5, 0x0b, 0x6f, 0x0f, 0x6e, 0x4a, 0x31, 0xd7,
	0x8c, 0xcc, 0x78, 0xcc, 0xb5, 0xff, 0xa4, 0x7d, 0x63, 0xff, 0x56, 0x00, 0xe6, 0xe8, 0x55, 0x76,
	0xe2, 0x00, 0xa1, 0x98, 0x27, 0xda, 0x3f, 0x28, 0x01, 0xfa, 0xd9, 0x89, 0xf7, 0x03, 0x6c, 0xe7,
	0x0c, 0xe2, 0x8a, 0x5c, 0x51, 0xcd, 0x24, 0x89, 0xa9, 0xfc, 0xe8, 0x1f, 0x1a, 0xe4, 0x26, 0x52,
	0x89, 0xab, 0xf7, 0xd9, 0xcd, 0x88, 0xca, 0x8f, 0xde, 0x01, 0xb4, 0xb8, 0x22, 0x25, 0x55, 0x22,
	0x19, 0x0d, 0x2f, 0x58, 0xe4, 0x1f, 0xb5, 0x6f, 0xec, 0x7f, 0x1e, 0x6c, 0x71, 0x15, 0x14, 0xfa,
	0x01, 0x5e, 0x79, 0x7d, 0x78, 0xa8, 0x45, 0x2a, 0x66, 0x62, 0x7a, 0x4d, 0x78, 0xa2, 0x99, 0x3c,
	0xa7, 0x21, 0x23, 0x4a, 0xd3, 0xc2, 0xfa, 0x8f, 0x46, 0x6f, 0xd7, 0xa2, 0x4e, 0x2d, 0x68, 0x9c,
	0x61, 0x30, 0xcb, 0x3a, 0x12, 0x8c, 0xf7, 0xd3, 0x3a, 0x12, 0xcc, 0xfb, 0x1d, 0x78, 0x72, 0xc1,
	0xeb, 0x69, 0x7f, 0x36, 0x85, 0x1b, 0x72, 0xc1, 0x2b, 0x59, 0xbf, 0x85, 0x4d, 0x03, 0xae, 0xc4,
	0xfc, 0xc5, 0xc4, 0x34, 0xd8, 0x72, 0xc4, 0xc7, 0x70, 0x5b, 0xb2, 0x29, 0x57, 0x19, 0x27, 0x46,
	0x7a, 0x6a, 0x48, 0x6f, 0xd9, 0x53, 0x0c, 0x51, 0x86, 0xa1, 0xe9, 0x5f, 0xab, 0x30, 0xb4, 0x79,
	0x08, 0xad, 0x2a, 0x5b, 0x21, 0xff, 0xcc, 0xc8, 0x6f, 0x57, 0x58, 0xad, 0x87, 0x47, 0xf0, 0x85,
	0xa4, 0xc9, 0x94, 0xa9, 0xdc, 0xc1, 0xb1, 0xa1, 0xbe, 0x89, 0x67, 0xa8, 0xef, 0x20, 0xa8, 0xfe,
	0x5b, 0x19, 0x82, 0xda, 0xdf, 0xc0, 0x9d, 0x1c, 0xa2, 0x2f, 0x24, 0x53, 0x17, 0x62, 0x16, 0xf9,
	0xbf, 0xe7, 0x0d, 0x32, 0xe7, 0xef, 0xec, 0xb1, 0x77, 0x04, 0x3b, 0xd9, 0x30, 0x94, 0x34, 0x0b,
	0x9f, 0xcf, 0xd1, 0x27, 0x57, 0x81, 0x53, 0xb7, 0x3e, 0xf7, 0xe1, 0xce, 0x44, 0xc9, 0x4a, 0x9d,
	0xff, 0xc2, 0x28, 0xdc, 0x9e, 0x28, 0x59, 0x2a, 0xa8, 0x21, 0xd1, 0x72, 0xaf, 0x86, 0x44, 0xd7,
	0x5d, 0xd8, 0x2a, 0x90, 0x25, 0xe3, 0x7d, 0x9c, 0x63, 0x0b, 0x76, 0xd6, 0x4f, 0x60, 0x8f, 0x2b,
	0x52, 0x22, 0x2f, 0x6a, 0x8a, 0x08, 0x27, 0x26, 0xc2, 0x2e, 0x57, 0x3d, 0x2b, 0x55, 0x94, 0xdb,
	0x24, 0x4f, 0xe1, 0x5e, 0x46, 0x11, 0xd2, 0x24, 0xe2, 0x51, 0x36, 0x88, 0x32, 0x25, 0x8a, 0xe9,
	0x3c, 0xd3, 0xc0, 0x88, 0xb7, 0x26, 0x4a, 0xf6, 0x2d, 0x20, 0x48, 0xc7, 0x4c, 0x63, 0xb6, 0x55,
	0xb5, 0x98, 0xf2, 0xe5, 0x8a, 0x5a, 0x4c, 0xdb, 0x83, 0x87, 0x8d, 0xb5, 0x2e, 0xf8, 0xd0, 0xd4,
	0xdf, 0x5b, 0xaa, 0x77, 0x1d, 0x18, 0x98, 0x0e, 0xc4, 0x74, 0xc1, 0xe3, 0x79, 0x4c, 0x58, 0x72,
	0x2e, 0x64, 0xc8, 0x62, 0x96, 0x68, 0x12, 0x71, 0x45, 0x27, 0x33, 0x16, 0xf9, 0x7f, 0x98, 0x0e,
	0xdc, 0xe7, 0x6a, 0x84, 0xa8, 0x81, 0x03, 0x9d, 0xe4, 0x98, 0xec, 0x17, 0xc5, 0x15, 0x49, 0x44,
	0x84, 0xdf, 0x90, 0x98, 0xc5, 0x42, 0x5e, 0xfb, 0xa7, 0xf8, 0x2b, 0xe1, 0xea, 0xb5, 0x88, 0xb2,
	0x0f, 0xc8, 0xc8, 0x1c, 0x7b, 0x5f, 0xc3, 0x06, 0x7e, 0x3a, 0x9c, 0xd1, 0x3f, 0xf1, 0x39, 0xcd,
	0xb1, 0x33, 0xf7, 0x0c, 0x76, 0xa7, 0x33, 0x31, 0xa1, 0x33, 0x42, 0xe7, 0x5a, 0x64, 0xf1, 0x2a,
	0xd3, 0x72, 0x66, 0x8a, 0x76, 0x10, 0xf2, 0x62, 0xae, 0x45, 0x90, 0x96, 0xc7, 0xe6, 0x2d, 0x7c,
	0xc5, 0x15, 0x59, 0x43, 0x50, 0xbc, 0xf1, 0x2b, 0xe3, 0xb3, 0xcd, 0xd5, 0xb0, 0x99, 0xca, 0x3e,
	0xf4, 0x11, 0xe4, 0x62, 0x64, 0x69, 0x72, 0x47, 0xc6, 0xcb, 0x36, 0x5e, 0xf7, 0xaa, 0xf3, 0xdb,
	0x58, 0x86, 0x0f, 0xfc, 0xba, 0xb1, 0x0c, 0x9f, 0xf7, 0xb8, 0x48, 0xdf, 0x34, 0xa0, 0xfe, 0x1b,
	0x53, 0xea, 0xd7, 0x4a, 0x5d, 0xf3, 0x4e, 0xa1, 0xe3, 0xe2, 0xd7, 0xfd, 0x16, 0xd1, 0xdf, 0x9a,
	0xe8, 0x0f, 0x6c, 0xf4, 0xaa, 0x73, 0x9b, 0x7b, 0x08, 0x8f, 0x4a, 0x3c, 0x2b, 0xe6, 0xfc, 0x2f,
	0xe3, 0xe7, 0x7e, 0xe1, 0xa7, 0x69, 0xda, 0x3f, 0x41, 0x84, 0x3d, 0x09, 0xd6, 0x12, 0x61, 0x6f,
	0xde, 0xc0, 0xe3, 0x75, 0x44, 0xae, 0x4b, 0x63, 0x43, 0xd6, 0x5e, 0x41, 0xe6, 0xba, 0xf5, 0x1c,
	0x1e, 0xb8, 0x6e, 0x35, 0x2d, 0xb6, 0x77, 0xa6, 0x51, 0x5f, 0xda, 0x46, 0x2d, 0xaf, 0xb7, 0x33,
	0xe8, 0xac, 0xdc, 0x4c, 0xce, 0xcf, 0xdf, 0xc6, 0xcf, 0x5e, 0xf3, 0x76, 0xaa, 0x4c, 0x7e, 0xc9,
	0xce, 0xd2, 0xfa, 0xf9, 0xc7, 0x98, 0xd9, 0x29, 0xcc, 0xd4, 0xd6, 0xd0, 0xf7, 0xe0, 0x15, 0x8b,
	0xc3, 0x49, 0xbf, 0xcf, 0xb7, 0x79, 0x7e, 0xe3, 0xc4, 0x9e, 0xc0, 0x5d, 0xab, 0x54, 0x5d, 0x5e,
	0xff, 0x9a, 0x8a, 0x2d, 0xbc, 0x0c, 0x2a, 0x2b, 0xec, 0x25, 0xb4, 0x4b, 0x06, 0x9b, 0xb7, 0xd4,
	0x07, 0xfb, 0xe1, 0x18, 0x2e, 0x13, 0xe4, 0x56, 0x27, 0x9f, 0x99, 0xbf, 0x4b, 0x07, 0xff, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0xb3, 0x1b, 0x6a, 0x47, 0x09, 0x00, 0x00,
}
