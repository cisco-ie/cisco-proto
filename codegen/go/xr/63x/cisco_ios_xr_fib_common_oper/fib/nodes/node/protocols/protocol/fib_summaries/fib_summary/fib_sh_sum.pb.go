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
// source: fib_sh_sum.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_fib_summaries_fib_summary

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

type FibShSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	TableId              string   `protobuf:"bytes,4,opt,name=table_id,json=tableId,proto3" json:"table_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShSum_KEYS) Reset()         { *m = FibShSum_KEYS{} }
func (m *FibShSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShSum_KEYS) ProtoMessage()    {}
func (*FibShSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{0}
}

func (m *FibShSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShSum_KEYS.Unmarshal(m, b)
}
func (m *FibShSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShSum_KEYS.Marshal(b, m, deterministic)
}
func (m *FibShSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShSum_KEYS.Merge(m, src)
}
func (m *FibShSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShSum_KEYS.Size(m)
}
func (m *FibShSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShSum_KEYS proto.InternalMessageInfo

func (m *FibShSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibShSum_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibShSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibShSum_KEYS) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

type FibPlLdiCount struct {
	TotalLoadSharingElementBytes      uint32   `protobuf:"varint,1,opt,name=total_load_sharing_element_bytes,json=totalLoadSharingElementBytes,proto3" json:"total_load_sharing_element_bytes,omitempty"`
	TotalLoadSharingElementReferences uint64   `protobuf:"varint,2,opt,name=total_load_sharing_element_references,json=totalLoadSharingElementReferences,proto3" json:"total_load_sharing_element_references,omitempty"`
	TotalPathListElements             uint32   `protobuf:"varint,3,opt,name=total_path_list_elements,json=totalPathListElements,proto3" json:"total_path_list_elements,omitempty"`
	RecursivePathListElements         uint32   `protobuf:"varint,4,opt,name=recursive_path_list_elements,json=recursivePathListElements,proto3" json:"recursive_path_list_elements,omitempty"`
	PlatformSharedPathListElements    uint32   `protobuf:"varint,5,opt,name=platform_shared_path_list_elements,json=platformSharedPathListElements,proto3" json:"platform_shared_path_list_elements,omitempty"`
	RetryPathListElements             uint32   `protobuf:"varint,6,opt,name=retry_path_list_elements,json=retryPathListElements,proto3" json:"retry_path_list_elements,omitempty"`
	TotalLoadInfoElements             uint32   `protobuf:"varint,7,opt,name=total_load_info_elements,json=totalLoadInfoElements,proto3" json:"total_load_info_elements,omitempty"`
	RecursiveLoadInfoElements         uint32   `protobuf:"varint,8,opt,name=recursive_load_info_elements,json=recursiveLoadInfoElements,proto3" json:"recursive_load_info_elements,omitempty"`
	PlatformSharedLoadInfoElements    uint32   `protobuf:"varint,9,opt,name=platform_shared_load_info_elements,json=platformSharedLoadInfoElements,proto3" json:"platform_shared_load_info_elements,omitempty"`
	XplLoadInfoElements               uint32   `protobuf:"varint,10,opt,name=xpl_load_info_elements,json=xplLoadInfoElements,proto3" json:"xpl_load_info_elements,omitempty"`
	XXX_NoUnkeyedLiteral              struct{} `json:"-"`
	XXX_unrecognized                  []byte   `json:"-"`
	XXX_sizecache                     int32    `json:"-"`
}

func (m *FibPlLdiCount) Reset()         { *m = FibPlLdiCount{} }
func (m *FibPlLdiCount) String() string { return proto.CompactTextString(m) }
func (*FibPlLdiCount) ProtoMessage()    {}
func (*FibPlLdiCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{1}
}

func (m *FibPlLdiCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibPlLdiCount.Unmarshal(m, b)
}
func (m *FibPlLdiCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibPlLdiCount.Marshal(b, m, deterministic)
}
func (m *FibPlLdiCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibPlLdiCount.Merge(m, src)
}
func (m *FibPlLdiCount) XXX_Size() int {
	return xxx_messageInfo_FibPlLdiCount.Size(m)
}
func (m *FibPlLdiCount) XXX_DiscardUnknown() {
	xxx_messageInfo_FibPlLdiCount.DiscardUnknown(m)
}

var xxx_messageInfo_FibPlLdiCount proto.InternalMessageInfo

func (m *FibPlLdiCount) GetTotalLoadSharingElementBytes() uint32 {
	if m != nil {
		return m.TotalLoadSharingElementBytes
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalLoadSharingElementReferences() uint64 {
	if m != nil {
		return m.TotalLoadSharingElementReferences
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalPathListElements() uint32 {
	if m != nil {
		return m.TotalPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRecursivePathListElements() uint32 {
	if m != nil {
		return m.RecursivePathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetPlatformSharedPathListElements() uint32 {
	if m != nil {
		return m.PlatformSharedPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRetryPathListElements() uint32 {
	if m != nil {
		return m.RetryPathListElements
	}
	return 0
}

func (m *FibPlLdiCount) GetTotalLoadInfoElements() uint32 {
	if m != nil {
		return m.TotalLoadInfoElements
	}
	return 0
}

func (m *FibPlLdiCount) GetRecursiveLoadInfoElements() uint32 {
	if m != nil {
		return m.RecursiveLoadInfoElements
	}
	return 0
}

func (m *FibPlLdiCount) GetPlatformSharedLoadInfoElements() uint32 {
	if m != nil {
		return m.PlatformSharedLoadInfoElements
	}
	return 0
}

func (m *FibPlLdiCount) GetXplLoadInfoElements() uint32 {
	if m != nil {
		return m.XplLoadInfoElements
	}
	return 0
}

type FibShPfxMasklenCnt struct {
	MaskLength           uint32   `protobuf:"varint,1,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	NumberOfPrefixes     uint32   `protobuf:"varint,2,opt,name=number_of_prefixes,json=numberOfPrefixes,proto3" json:"number_of_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShPfxMasklenCnt) Reset()         { *m = FibShPfxMasklenCnt{} }
func (m *FibShPfxMasklenCnt) String() string { return proto.CompactTextString(m) }
func (*FibShPfxMasklenCnt) ProtoMessage()    {}
func (*FibShPfxMasklenCnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{2}
}

func (m *FibShPfxMasklenCnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShPfxMasklenCnt.Unmarshal(m, b)
}
func (m *FibShPfxMasklenCnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShPfxMasklenCnt.Marshal(b, m, deterministic)
}
func (m *FibShPfxMasklenCnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShPfxMasklenCnt.Merge(m, src)
}
func (m *FibShPfxMasklenCnt) XXX_Size() int {
	return xxx_messageInfo_FibShPfxMasklenCnt.Size(m)
}
func (m *FibShPfxMasklenCnt) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShPfxMasklenCnt.DiscardUnknown(m)
}

var xxx_messageInfo_FibShPfxMasklenCnt proto.InternalMessageInfo

func (m *FibShPfxMasklenCnt) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

func (m *FibShPfxMasklenCnt) GetNumberOfPrefixes() uint32 {
	if m != nil {
		return m.NumberOfPrefixes
	}
	return 0
}

type FibShPfxMasklenDistrib struct {
	UnicastPrefixe       []*FibShPfxMasklenCnt `protobuf:"bytes,1,rep,name=unicast_prefixe,json=unicastPrefixe,proto3" json:"unicast_prefixe,omitempty"`
	BroadcastPrefixe     []*FibShPfxMasklenCnt `protobuf:"bytes,2,rep,name=broadcast_prefixe,json=broadcastPrefixe,proto3" json:"broadcast_prefixe,omitempty"`
	MulticastPrefix      []*FibShPfxMasklenCnt `protobuf:"bytes,3,rep,name=multicast_prefix,json=multicastPrefix,proto3" json:"multicast_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FibShPfxMasklenDistrib) Reset()         { *m = FibShPfxMasklenDistrib{} }
func (m *FibShPfxMasklenDistrib) String() string { return proto.CompactTextString(m) }
func (*FibShPfxMasklenDistrib) ProtoMessage()    {}
func (*FibShPfxMasklenDistrib) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{3}
}

func (m *FibShPfxMasklenDistrib) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Unmarshal(m, b)
}
func (m *FibShPfxMasklenDistrib) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Marshal(b, m, deterministic)
}
func (m *FibShPfxMasklenDistrib) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShPfxMasklenDistrib.Merge(m, src)
}
func (m *FibShPfxMasklenDistrib) XXX_Size() int {
	return xxx_messageInfo_FibShPfxMasklenDistrib.Size(m)
}
func (m *FibShPfxMasklenDistrib) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShPfxMasklenDistrib.DiscardUnknown(m)
}

var xxx_messageInfo_FibShPfxMasklenDistrib proto.InternalMessageInfo

func (m *FibShPfxMasklenDistrib) GetUnicastPrefixe() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.UnicastPrefixe
	}
	return nil
}

func (m *FibShPfxMasklenDistrib) GetBroadcastPrefixe() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.BroadcastPrefixe
	}
	return nil
}

func (m *FibShPfxMasklenDistrib) GetMulticastPrefix() []*FibShPfxMasklenCnt {
	if m != nil {
		return m.MulticastPrefix
	}
	return nil
}

type FibShSum struct {
	Prefix                        []byte                  `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	SsTblId                       uint32                  `protobuf:"varint,51,opt,name=ss_tbl_id,json=ssTblId,proto3" json:"ss_tbl_id,omitempty"`
	SsTblIdPtr                    uint32                  `protobuf:"varint,52,opt,name=ss_tbl_id_ptr,json=ssTblIdPtr,proto3" json:"ss_tbl_id_ptr,omitempty"`
	SsVrfId                       uint32                  `protobuf:"varint,53,opt,name=ss_vrf_id,json=ssVrfId,proto3" json:"ss_vrf_id,omitempty"`
	SsVrId                        uint32                  `protobuf:"varint,54,opt,name=ss_vr_id,json=ssVrId,proto3" json:"ss_vr_id,omitempty"`
	LoadBalancing                 string                  `protobuf:"bytes,55,opt,name=load_balancing,json=loadBalancing,proto3" json:"load_balancing,omitempty"`
	ForwardingElements            uint32                  `protobuf:"varint,56,opt,name=forwarding_elements,json=forwardingElements,proto3" json:"forwarding_elements,omitempty"`
	Routes                        uint32                  `protobuf:"varint,57,opt,name=routes,proto3" json:"routes,omitempty"`
	PrefixInPlaceModifications    uint32                  `protobuf:"varint,58,opt,name=prefix_in_place_modifications,json=prefixInPlaceModifications,proto3" json:"prefix_in_place_modifications,omitempty"`
	StalePrefixDeletes            uint32                  `protobuf:"varint,59,opt,name=stale_prefix_deletes,json=stalePrefixDeletes,proto3" json:"stale_prefix_deletes,omitempty"`
	LoadSharingElements           uint32                  `protobuf:"varint,60,opt,name=load_sharing_elements,json=loadSharingElements,proto3" json:"load_sharing_elements,omitempty"`
	LoadSharingReferences         uint64                  `protobuf:"varint,61,opt,name=load_sharing_references,json=loadSharingReferences,proto3" json:"load_sharing_references,omitempty"`
	TotalLoadShareElementBytes    uint32                  `protobuf:"varint,62,opt,name=total_load_share_element_bytes,json=totalLoadShareElementBytes,proto3" json:"total_load_share_element_bytes,omitempty"`
	ExclusiveLoadSharingElement   *FibPlLdiCount          `protobuf:"bytes,63,opt,name=exclusive_load_sharing_element,json=exclusiveLoadSharingElement,proto3" json:"exclusive_load_sharing_element,omitempty"`
	SharedLoadSharingElement      *FibPlLdiCount          `protobuf:"bytes,64,opt,name=shared_load_sharing_element,json=sharedLoadSharingElement,proto3" json:"shared_load_sharing_element,omitempty"`
	CrossSharedLoadSharingElement *FibPlLdiCount          `protobuf:"bytes,65,opt,name=cross_shared_load_sharing_element,json=crossSharedLoadSharingElement,proto3" json:"cross_shared_load_sharing_element,omitempty"`
	LabelSharedLoadSharingElement *FibPlLdiCount          `protobuf:"bytes,66,opt,name=label_shared_load_sharing_element,json=labelSharedLoadSharingElement,proto3" json:"label_shared_load_sharing_element,omitempty"`
	LeavesUsedBytes               uint32                  `protobuf:"varint,67,opt,name=leaves_used_bytes,json=leavesUsedBytes,proto3" json:"leaves_used_bytes,omitempty"`
	ReresolveEntries              uint32                  `protobuf:"varint,68,opt,name=reresolve_entries,json=reresolveEntries,proto3" json:"reresolve_entries,omitempty"`
	OldUnresolveEntries           uint32                  `protobuf:"varint,69,opt,name=old_unresolve_entries,json=oldUnresolveEntries,proto3" json:"old_unresolve_entries,omitempty"`
	NewUnresolveEntries           uint32                  `protobuf:"varint,70,opt,name=new_unresolve_entries,json=newUnresolveEntries,proto3" json:"new_unresolve_entries,omitempty"`
	UnresolveEntries              uint32                  `protobuf:"varint,71,opt,name=unresolve_entries,json=unresolveEntries,proto3" json:"unresolve_entries,omitempty"`
	CefRouteDrops                 uint32                  `protobuf:"varint,72,opt,name=cef_route_drops,json=cefRouteDrops,proto3" json:"cef_route_drops,omitempty"`
	CefVersionMismatchRouteDrops  uint64                  `protobuf:"varint,73,opt,name=cef_version_mismatch_route_drops,json=cefVersionMismatchRouteDrops,proto3" json:"cef_version_mismatch_route_drops,omitempty"`
	DeleteCacheNumEntries         uint32                  `protobuf:"varint,74,opt,name=delete_cache_num_entries,json=deleteCacheNumEntries,proto3" json:"delete_cache_num_entries,omitempty"`
	ExistingLeavesRevisions       uint32                  `protobuf:"varint,75,opt,name=existing_leaves_revisions,json=existingLeavesRevisions,proto3" json:"existing_leaves_revisions,omitempty"`
	FibDefaultPrefix              uint32                  `protobuf:"varint,76,opt,name=fib_default_prefix,json=fibDefaultPrefix,proto3" json:"fib_default_prefix,omitempty"`
	FibDefaultPrefixMaskLength    uint32                  `protobuf:"varint,77,opt,name=fib_default_prefix_mask_length,json=fibDefaultPrefixMaskLength,proto3" json:"fib_default_prefix_mask_length,omitempty"`
	NextHops                      uint32                  `protobuf:"varint,78,opt,name=next_hops,json=nextHops,proto3" json:"next_hops,omitempty"`
	IncompleteNextHops            uint32                  `protobuf:"varint,79,opt,name=incomplete_next_hops,json=incompleteNextHops,proto3" json:"incomplete_next_hops,omitempty"`
	ResolutionTimer               uint32                  `protobuf:"varint,80,opt,name=resolution_timer,json=resolutionTimer,proto3" json:"resolution_timer,omitempty"`
	SlowProcessTimer              uint32                  `protobuf:"varint,81,opt,name=slow_process_timer,json=slowProcessTimer,proto3" json:"slow_process_timer,omitempty"`
	MaxResolutionTimer            uint32                  `protobuf:"varint,82,opt,name=max_resolution_timer,json=maxResolutionTimer,proto3" json:"max_resolution_timer,omitempty"`
	ImpositionPrefixes            uint32                  `protobuf:"varint,83,opt,name=imposition_prefixes,json=impositionPrefixes,proto3" json:"imposition_prefixes,omitempty"`
	ExtendedPrefixes              uint32                  `protobuf:"varint,84,opt,name=extended_prefixes,json=extendedPrefixes,proto3" json:"extended_prefixes,omitempty"`
	CeflBlRecycledRoutes          uint32                  `protobuf:"varint,85,opt,name=cefl_bl_recycled_routes,json=ceflBlRecycledRoutes,proto3" json:"cefl_bl_recycled_routes,omitempty"`
	LdiBackwalks                  uint32                  `protobuf:"varint,86,opt,name=ldi_backwalks,json=ldiBackwalks,proto3" json:"ldi_backwalks,omitempty"`
	SsProtRouteCount              uint32                  `protobuf:"varint,87,opt,name=ss_prot_route_count,json=ssProtRouteCount,proto3" json:"ss_prot_route_count,omitempty"`
	LispEidPrefixes               uint32                  `protobuf:"varint,88,opt,name=lisp_eid_prefixes,json=lispEidPrefixes,proto3" json:"lisp_eid_prefixes,omitempty"`
	LispEidValidPrefixes          uint32                  `protobuf:"varint,89,opt,name=lisp_eid_valid_prefixes,json=lispEidValidPrefixes,proto3" json:"lisp_eid_valid_prefixes,omitempty"`
	LispRlocObjects               uint32                  `protobuf:"varint,90,opt,name=lisp_rloc_objects,json=lispRlocObjects,proto3" json:"lisp_rloc_objects,omitempty"`
	SsVxlanLtepIfh                string                  `protobuf:"bytes,91,opt,name=ss_vxlan_ltep_ifh,json=ssVxlanLtepIfh,proto3" json:"ss_vxlan_ltep_ifh,omitempty"`
	SsDropPlCount                 uint32                  `protobuf:"varint,92,opt,name=ss_drop_pl_count,json=ssDropPlCount,proto3" json:"ss_drop_pl_count,omitempty"`
	PrefixMasklenDistribution     *FibShPfxMasklenDistrib `protobuf:"bytes,93,opt,name=prefix_masklen_distribution,json=prefixMasklenDistribution,proto3" json:"prefix_masklen_distribution,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                `json:"-"`
	XXX_unrecognized              []byte                  `json:"-"`
	XXX_sizecache                 int32                   `json:"-"`
}

func (m *FibShSum) Reset()         { *m = FibShSum{} }
func (m *FibShSum) String() string { return proto.CompactTextString(m) }
func (*FibShSum) ProtoMessage()    {}
func (*FibShSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{4}
}

func (m *FibShSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShSum.Unmarshal(m, b)
}
func (m *FibShSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShSum.Marshal(b, m, deterministic)
}
func (m *FibShSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShSum.Merge(m, src)
}
func (m *FibShSum) XXX_Size() int {
	return xxx_messageInfo_FibShSum.Size(m)
}
func (m *FibShSum) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShSum.DiscardUnknown(m)
}

var xxx_messageInfo_FibShSum proto.InternalMessageInfo

func (m *FibShSum) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *FibShSum) GetSsTblId() uint32 {
	if m != nil {
		return m.SsTblId
	}
	return 0
}

func (m *FibShSum) GetSsTblIdPtr() uint32 {
	if m != nil {
		return m.SsTblIdPtr
	}
	return 0
}

func (m *FibShSum) GetSsVrfId() uint32 {
	if m != nil {
		return m.SsVrfId
	}
	return 0
}

func (m *FibShSum) GetSsVrId() uint32 {
	if m != nil {
		return m.SsVrId
	}
	return 0
}

func (m *FibShSum) GetLoadBalancing() string {
	if m != nil {
		return m.LoadBalancing
	}
	return ""
}

func (m *FibShSum) GetForwardingElements() uint32 {
	if m != nil {
		return m.ForwardingElements
	}
	return 0
}

func (m *FibShSum) GetRoutes() uint32 {
	if m != nil {
		return m.Routes
	}
	return 0
}

func (m *FibShSum) GetPrefixInPlaceModifications() uint32 {
	if m != nil {
		return m.PrefixInPlaceModifications
	}
	return 0
}

func (m *FibShSum) GetStalePrefixDeletes() uint32 {
	if m != nil {
		return m.StalePrefixDeletes
	}
	return 0
}

func (m *FibShSum) GetLoadSharingElements() uint32 {
	if m != nil {
		return m.LoadSharingElements
	}
	return 0
}

func (m *FibShSum) GetLoadSharingReferences() uint64 {
	if m != nil {
		return m.LoadSharingReferences
	}
	return 0
}

func (m *FibShSum) GetTotalLoadShareElementBytes() uint32 {
	if m != nil {
		return m.TotalLoadShareElementBytes
	}
	return 0
}

func (m *FibShSum) GetExclusiveLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.ExclusiveLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.SharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetCrossSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.CrossSharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetLabelSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.LabelSharedLoadSharingElement
	}
	return nil
}

func (m *FibShSum) GetLeavesUsedBytes() uint32 {
	if m != nil {
		return m.LeavesUsedBytes
	}
	return 0
}

func (m *FibShSum) GetReresolveEntries() uint32 {
	if m != nil {
		return m.ReresolveEntries
	}
	return 0
}

func (m *FibShSum) GetOldUnresolveEntries() uint32 {
	if m != nil {
		return m.OldUnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetNewUnresolveEntries() uint32 {
	if m != nil {
		return m.NewUnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetUnresolveEntries() uint32 {
	if m != nil {
		return m.UnresolveEntries
	}
	return 0
}

func (m *FibShSum) GetCefRouteDrops() uint32 {
	if m != nil {
		return m.CefRouteDrops
	}
	return 0
}

func (m *FibShSum) GetCefVersionMismatchRouteDrops() uint64 {
	if m != nil {
		return m.CefVersionMismatchRouteDrops
	}
	return 0
}

func (m *FibShSum) GetDeleteCacheNumEntries() uint32 {
	if m != nil {
		return m.DeleteCacheNumEntries
	}
	return 0
}

func (m *FibShSum) GetExistingLeavesRevisions() uint32 {
	if m != nil {
		return m.ExistingLeavesRevisions
	}
	return 0
}

func (m *FibShSum) GetFibDefaultPrefix() uint32 {
	if m != nil {
		return m.FibDefaultPrefix
	}
	return 0
}

func (m *FibShSum) GetFibDefaultPrefixMaskLength() uint32 {
	if m != nil {
		return m.FibDefaultPrefixMaskLength
	}
	return 0
}

func (m *FibShSum) GetNextHops() uint32 {
	if m != nil {
		return m.NextHops
	}
	return 0
}

func (m *FibShSum) GetIncompleteNextHops() uint32 {
	if m != nil {
		return m.IncompleteNextHops
	}
	return 0
}

func (m *FibShSum) GetResolutionTimer() uint32 {
	if m != nil {
		return m.ResolutionTimer
	}
	return 0
}

func (m *FibShSum) GetSlowProcessTimer() uint32 {
	if m != nil {
		return m.SlowProcessTimer
	}
	return 0
}

func (m *FibShSum) GetMaxResolutionTimer() uint32 {
	if m != nil {
		return m.MaxResolutionTimer
	}
	return 0
}

func (m *FibShSum) GetImpositionPrefixes() uint32 {
	if m != nil {
		return m.ImpositionPrefixes
	}
	return 0
}

func (m *FibShSum) GetExtendedPrefixes() uint32 {
	if m != nil {
		return m.ExtendedPrefixes
	}
	return 0
}

func (m *FibShSum) GetCeflBlRecycledRoutes() uint32 {
	if m != nil {
		return m.CeflBlRecycledRoutes
	}
	return 0
}

func (m *FibShSum) GetLdiBackwalks() uint32 {
	if m != nil {
		return m.LdiBackwalks
	}
	return 0
}

func (m *FibShSum) GetSsProtRouteCount() uint32 {
	if m != nil {
		return m.SsProtRouteCount
	}
	return 0
}

func (m *FibShSum) GetLispEidPrefixes() uint32 {
	if m != nil {
		return m.LispEidPrefixes
	}
	return 0
}

func (m *FibShSum) GetLispEidValidPrefixes() uint32 {
	if m != nil {
		return m.LispEidValidPrefixes
	}
	return 0
}

func (m *FibShSum) GetLispRlocObjects() uint32 {
	if m != nil {
		return m.LispRlocObjects
	}
	return 0
}

func (m *FibShSum) GetSsVxlanLtepIfh() string {
	if m != nil {
		return m.SsVxlanLtepIfh
	}
	return ""
}

func (m *FibShSum) GetSsDropPlCount() uint32 {
	if m != nil {
		return m.SsDropPlCount
	}
	return 0
}

func (m *FibShSum) GetPrefixMasklenDistribution() *FibShPfxMasklenDistrib {
	if m != nil {
		return m.PrefixMasklenDistribution
	}
	return nil
}

func init() {
	proto.RegisterType((*FibShSum_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_sum_KEYS")
	proto.RegisterType((*FibPlLdiCount)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_pl_ldi_count")
	proto.RegisterType((*FibShPfxMasklenCnt)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_pfx_masklen_cnt")
	proto.RegisterType((*FibShPfxMasklenDistrib)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_pfx_masklen_distrib")
	proto.RegisterType((*FibShSum)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.fib_summaries.fib_summary.fib_sh_sum")
}

func init() { proto.RegisterFile("fib_sh_sum.proto", fileDescriptor_6b124c48422c0f95) }

var fileDescriptor_6b124c48422c0f95 = []byte{
	// 1474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x5d, 0x73, 0xd4, 0xbc,
	0x15, 0x9e, 0x25, 0x10, 0x12, 0xc1, 0x92, 0xc4, 0xf9, 0xc0, 0x21, 0x40, 0x43, 0x18, 0xda, 0x50,
	0xda, 0xc0, 0x84, 0x42, 0x5a, 0xfa, 0x41, 0x09, 0x09, 0x65, 0x21, 0x09, 0x5b, 0x07, 0x52, 0xe8,
	0xc7, 0x68, 0x64, 0xfb, 0x38, 0x2b, 0x90, 0x25, 0x8f, 0x24, 0x6f, 0x36, 0x3f, 0xa1, 0x3f, 0xa1,
	0xd3, 0xf6, 0xf6, 0x9d, 0x79, 0xaf, 0xde, 0x1b, 0x7e, 0xdb, 0x7b, 0xfd, 0x8e, 0x24, 0xdb, 0xeb,
	0xec, 0x2e, 0xb9, 0x23, 0xdc, 0x64, 0xe2, 0xf3, 0x3c, 0xcf, 0xf1, 0x73, 0x64, 0xe9, 0x1c, 0x2d,
	0x9a, 0x4e, 0x68, 0x88, 0x55, 0x07, 0xab, 0x3c, 0x5d, 0xcb, 0xa4, 0xd0, 0xc2, 0x7b, 0x1f, 0x51,
	0x15, 0x09, 0x4c, 0x85, 0xc2, 0x3d, 0x89, 0x0d, 0x1c, 0x89, 0x34, 0x15, 0x1c, 0x8b, 0x0c, 0xe4,
	0x5a, 0x42, 0xc3, 0x35, 0x2e, 0x62, 0x50, 0xf6, 0xaf, 0x93, 0x44, 0x82, 0xa9, 0xea, 0xbf, 0x35,
	0x9b, 0x30, 0x4f, 0x53, 0x22, 0x29, 0xa8, 0xda, 0xd3, 0xf1, 0xca, 0xbf, 0x1b, 0x68, 0xaa, 0xff,
	0x3a, 0xfc, 0x7a, 0xfb, 0xc3, 0xbe, 0xb7, 0x84, 0x26, 0x4d, 0x22, 0xcc, 0x49, 0x0a, 0x7e, 0x63,
	0xb9, 0xb1, 0x3a, 0x19, 0x4c, 0x98, 0xc0, 0x1e, 0x49, 0xc1, 0xbb, 0x8d, 0x9a, 0x65, 0x5a, 0x47,
	0x38, 0x67, 0x09, 0x97, 0xcb, 0xa0, 0x25, 0x2d, 0xa2, 0x89, 0xae, 0x4c, 0x1c, 0x3e, 0x66, 0xf1,
	0x8b, 0x5d, 0x99, 0x94, 0x90, 0x26, 0x21, 0x03, 0x4c, 0x63, 0xff, 0xbc, 0x83, 0xec, 0x73, 0x2b,
	0x5e, 0xf9, 0xee, 0x82, 0x2b, 0x3d, 0x63, 0x98, 0xc5, 0x14, 0x47, 0x22, 0xe7, 0xda, 0x7b, 0x81,
	0x96, 0xb5, 0xd0, 0x84, 0x61, 0x26, 0x48, 0x8c, 0x55, 0x87, 0x48, 0xca, 0x0f, 0x31, 0x30, 0x48,
	0x81, 0x6b, 0x1c, 0x1e, 0x6b, 0x50, 0xd6, 0x63, 0x33, 0xb8, 0x6e, 0x79, 0x3b, 0x82, 0xc4, 0xfb,
	0x8e, 0xb5, 0xed, 0x48, 0x9b, 0x86, 0xe3, 0xb5, 0xd1, 0x9d, 0x53, 0xf2, 0x48, 0x48, 0x40, 0x02,
	0x8f, 0x40, 0xd9, 0x7a, 0xce, 0x07, 0xb7, 0xbe, 0x90, 0x2c, 0xa8, 0x88, 0xde, 0x06, 0xf2, 0x5d,
	0xc6, 0x8c, 0xe8, 0x0e, 0x66, 0x54, 0xe9, 0x32, 0x9d, 0xb2, 0x45, 0x37, 0x83, 0x79, 0x8b, 0xb7,
	0x89, 0xee, 0xec, 0x50, 0xa5, 0x8b, 0x0c, 0xca, 0x7b, 0x8a, 0xae, 0x4b, 0x88, 0x72, 0xa9, 0x68,
	0x17, 0x46, 0x89, 0xcf, 0x5b, 0xf1, 0x62, 0xc5, 0x19, 0x4a, 0xf0, 0x0a, 0xad, 0x64, 0x8c, 0xe8,
	0x44, 0xc8, 0xd4, 0x56, 0x02, 0xf1, 0xa8, 0x34, 0x17, 0x6c, 0x9a, 0x9b, 0x25, 0x73, 0xdf, 0x12,
	0x87, 0x72, 0x6d, 0x20, 0x5f, 0x82, 0x96, 0xc7, 0xa3, 0x32, 0x8c, 0xbb, 0x2a, 0x2c, 0x3e, 0x4a,
	0x58, 0x5b, 0x50, 0xca, 0x13, 0xd1, 0x17, 0x5e, 0xac, 0x95, 0x6f, 0xd6, 0xb0, 0xc5, 0x13, 0x31,
	0xba, 0xfc, 0x11, 0xe2, 0x89, 0x81, 0xf2, 0x87, 0x12, 0x8c, 0x28, 0x7f, 0x44, 0x9a, 0xc9, 0x51,
	0xe5, 0x0f, 0xe5, 0x7a, 0x88, 0x16, 0x7a, 0xd9, 0xc8, 0x1a, 0x90, 0xd5, 0xcf, 0xf6, 0xb2, 0xa1,
	0x0a, 0x56, 0x0e, 0xd1, 0x42, 0x71, 0x66, 0xb2, 0xa4, 0x87, 0x53, 0xa2, 0x3e, 0x31, 0xe0, 0x38,
	0xe2, 0xda, 0xfb, 0x19, 0xba, 0x64, 0x1e, 0x31, 0x03, 0x7e, 0xa8, 0x3b, 0xc5, 0xc6, 0x44, 0x26,
	0xb4, 0x63, 0x23, 0xde, 0xaf, 0x90, 0xc7, 0xf3, 0x34, 0x04, 0x89, 0x45, 0x82, 0x33, 0x09, 0x09,
	0xed, 0x15, 0x7b, 0xae, 0x19, 0x4c, 0x3b, 0xe4, 0x4d, 0xd2, 0x2e, 0xe2, 0x2b, 0x3f, 0x8e, 0xa1,
	0x6b, 0x23, 0xde, 0x14, 0x53, 0xa5, 0x25, 0x0d, 0xbd, 0xff, 0x34, 0xd0, 0x54, 0xce, 0x69, 0x44,
	0x94, 0x2e, 0x73, 0xf9, 0x8d, 0xe5, 0xb1, 0xd5, 0x4b, 0xeb, 0xd9, 0xda, 0xd7, 0xea, 0x18, 0x6b,
	0xa3, 0x2b, 0x0f, 0xae, 0x14, 0x46, 0x0a, 0xef, 0xde, 0xff, 0x1b, 0x68, 0x26, 0x94, 0x82, 0xc4,
	0x27, 0xdc, 0x9d, 0xfb, 0x46, 0xee, 0xa6, 0x2b, 0x2b, 0xa5, 0xbf, 0xff, 0x36, 0xd0, 0x74, 0x9a,
	0x33, 0x5d, 0x5f, 0x3d, 0x7f, 0xec, 0x1b, 0xd9, 0x9b, 0xaa, 0x9c, 0x38, 0x7b, 0x2b, 0xff, 0x9b,
	0x47, 0xa8, 0xdf, 0x96, 0xbd, 0x05, 0x34, 0x5e, 0x38, 0x5c, 0x5f, 0x6e, 0xac, 0x5e, 0x0e, 0x8a,
	0x27, 0xef, 0x1a, 0x9a, 0x54, 0x0a, 0xeb, 0x90, 0x99, 0x6e, 0xfa, 0xd0, 0x6e, 0xa2, 0x8b, 0x4a,
	0xbd, 0x0d, 0x59, 0x2b, 0xf6, 0x6e, 0xa1, 0x66, 0x85, 0xe1, 0x4c, 0x4b, 0xff, 0x37, 0x6e, 0x33,
	0x16, 0x78, 0x5b, 0xcb, 0x42, 0x6e, 0x3a, 0x35, 0x8d, 0xfd, 0x47, 0xa5, 0xfc, 0x40, 0x26, 0xad,
	0xd8, 0xf3, 0xd1, 0x84, 0xc5, 0x0c, 0xf4, 0xd8, 0x42, 0xe3, 0x06, 0x6a, 0xc5, 0xde, 0x1d, 0x74,
	0xc5, 0x1e, 0x97, 0x90, 0x30, 0xc2, 0x23, 0xca, 0x0f, 0xfd, 0x0d, 0xdb, 0xc7, 0x9b, 0x26, 0xba,
	0x59, 0x06, 0xbd, 0xfb, 0x68, 0x36, 0x11, 0xf2, 0x88, 0xc8, 0xb8, 0xd6, 0x68, 0x95, 0xff, 0x5b,
	0x9b, 0xcb, 0xeb, 0x43, 0xd5, 0x51, 0x5c, 0x40, 0xe3, 0x52, 0xe4, 0xa6, 0x9f, 0xff, 0xce, 0xbd,
	0xcf, 0x3d, 0x79, 0xcf, 0xd0, 0x0d, 0x57, 0x2e, 0xa6, 0x1c, 0x67, 0x8c, 0x44, 0x80, 0x53, 0x11,
	0xd3, 0x84, 0x46, 0x44, 0x53, 0xc1, 0x95, 0xff, 0xc4, 0xd2, 0xaf, 0x39, 0x52, 0x8b, 0xb7, 0x0d,
	0x65, 0xb7, 0xce, 0xf0, 0x1e, 0xa0, 0x39, 0xa5, 0x09, 0x83, 0xe2, 0x3b, 0xe3, 0x18, 0x18, 0x98,
	0x17, 0xfd, 0xde, 0x99, 0xb1, 0x98, 0x5b, 0xf9, 0x2d, 0x87, 0x78, 0xeb, 0x68, 0x7e, 0xd4, 0xa0,
	0x50, 0xfe, 0x1f, 0x5c, 0x5b, 0x60, 0x43, 0x93, 0x41, 0x79, 0x8f, 0xd1, 0xd5, 0x13, 0x9a, 0xda,
	0x50, 0xf9, 0xa3, 0x1d, 0x2a, 0xf3, 0x35, 0x55, 0x6d, 0x90, 0x6c, 0xa2, 0x9b, 0x03, 0xa3, 0x09,
	0x06, 0x06, 0xdc, 0x9f, 0x5c, 0x85, 0x27, 0x66, 0x12, 0x9c, 0x18, 0x6f, 0x3f, 0x34, 0xd0, 0x4d,
	0xe8, 0x45, 0x2c, 0xef, 0x77, 0xd5, 0x01, 0xeb, 0xfe, 0xd3, 0xe5, 0xc6, 0xea, 0xa5, 0xf5, 0x8f,
	0x5f, 0x77, 0x73, 0xd7, 0x67, 0x77, 0xb0, 0x54, 0x39, 0x1a, 0x1e, 0xa4, 0xde, 0xf7, 0x0d, 0xb4,
	0x54, 0xef, 0xde, 0x83, 0x76, 0xff, 0x7c, 0xe6, 0x76, 0x7d, 0x55, 0xcd, 0x88, 0x01, 0xaf, 0x9f,
	0x1b, 0xe8, 0x56, 0x24, 0x85, 0x52, 0xf8, 0x34, 0xc7, 0xcf, 0xce, 0xdc, 0xf1, 0x0d, 0x6b, 0x6a,
	0xff, 0x34, 0xdb, 0x8c, 0x84, 0xc0, 0x4e, 0xb5, 0xbd, 0x79, 0xf6, 0xb6, 0xad, 0xa9, 0x2f, 0xda,
	0xfe, 0x25, 0x9a, 0x61, 0x40, 0xba, 0xa0, 0x70, 0xae, 0x20, 0x2e, 0x8e, 0xc0, 0x73, 0x7b, 0x04,
	0xa6, 0x1c, 0xf0, 0x4e, 0x41, 0xec, 0xf6, 0xfd, 0x3d, 0x34, 0x23, 0x41, 0x82, 0x12, 0xac, 0x0b,
	0x18, 0xb8, 0x36, 0x2f, 0xf7, 0xb7, 0xdc, 0x38, 0xad, 0x80, 0x6d, 0x17, 0x37, 0x87, 0x5a, 0xb0,
	0x18, 0xe7, 0x7c, 0x50, 0xb0, 0xed, 0x0e, 0xb5, 0x60, 0xf1, 0x3b, 0x3e, 0xac, 0xe1, 0x70, 0x34,
	0x42, 0xf3, 0xc2, 0x69, 0x38, 0x1c, 0x0d, 0x69, 0xee, 0xa1, 0x99, 0x61, 0xfe, 0x5f, 0x9c, 0xa9,
	0x7c, 0x90, 0xfc, 0x73, 0x34, 0x15, 0x41, 0x82, 0x6d, 0xb3, 0xc3, 0xb1, 0x14, 0x99, 0xf2, 0x5f,
	0x5a, 0x6a, 0x33, 0x82, 0x24, 0x30, 0xd1, 0x2d, 0x13, 0x34, 0x17, 0x61, 0xc3, 0xeb, 0x82, 0x54,
	0x54, 0x70, 0x9c, 0x52, 0x95, 0x12, 0x1d, 0x75, 0x4e, 0x08, 0x5b, 0xb6, 0xcd, 0x5c, 0x8f, 0x20,
	0x39, 0x70, 0xb4, 0xdd, 0x82, 0x55, 0xcb, 0xb3, 0x81, 0x7c, 0xd7, 0xfe, 0x70, 0x44, 0xa2, 0x0e,
	0x60, 0x9e, 0xa7, 0x95, 0xc7, 0x57, 0xee, 0xde, 0xe6, 0xf0, 0xe7, 0x06, 0xde, 0xcb, 0xd3, 0xd2,
	0xe8, 0x13, 0xb4, 0x08, 0x3d, 0xaa, 0xb4, 0xd9, 0x3b, 0xc5, 0xf7, 0x91, 0xd0, 0xa5, 0xca, 0xf6,
	0xe0, 0xd7, 0x56, 0x79, 0xb5, 0x24, 0xec, 0x58, 0x3c, 0x28, 0x61, 0x73, 0xed, 0x31, 0xbb, 0x20,
	0x86, 0x84, 0xe4, 0xac, 0x1a, 0xb7, 0x3b, 0x6e, 0x49, 0x12, 0x1a, 0x6e, 0x39, 0xc0, 0xf5, 0x60,
	0xd3, 0x10, 0x87, 0xd9, 0xb8, 0x7e, 0xb1, 0xda, 0x75, 0x0d, 0x71, 0x50, 0xb9, 0xdb, 0xbf, 0x68,
	0x99, 0x1f, 0x31, 0xd0, 0xd3, 0xb8, 0x63, 0xd6, 0x65, 0xcf, 0xd2, 0x27, 0x4c, 0xe0, 0xa5, 0x59,
	0x83, 0x07, 0x68, 0x8e, 0xf2, 0x48, 0xa4, 0x99, 0x5d, 0x87, 0x3e, 0xef, 0x8d, 0x9b, 0x07, 0x7d,
	0x6c, 0xaf, 0x54, 0xdc, 0x45, 0xd3, 0xf6, 0xbb, 0xe5, 0x66, 0xa0, 0x60, 0x4d, 0x53, 0x90, 0x7e,
	0xdb, 0x6d, 0xc9, 0x7e, 0xfc, 0xad, 0x09, 0x9b, 0x5a, 0x15, 0x13, 0x47, 0x38, 0x93, 0x22, 0x02,
	0x33, 0x82, 0x2d, 0xf9, 0xaf, 0xae, 0x56, 0x83, 0xb4, 0x1d, 0xe0, 0xd8, 0x0f, 0xd0, 0x5c, 0x4a,
	0x7a, 0x78, 0x28, 0x79, 0xe0, 0xac, 0xa4, 0xa4, 0x17, 0x0c, 0xe4, 0xbf, 0x8f, 0x66, 0x69, 0x9a,
	0x09, 0x45, 0x2d, 0xbb, 0xba, 0x43, 0xee, 0x17, 0xde, 0x2b, 0xa8, 0xbc, 0x45, 0x9a, 0xed, 0x08,
	0x3d, 0x0d, 0x3c, 0x36, 0xbf, 0x13, 0x4a, 0xfa, 0x5b, 0xe7, 0xa7, 0x04, 0x2a, 0xf2, 0x23, 0x74,
	0x35, 0x82, 0x84, 0xe1, 0x90, 0x61, 0x09, 0xd1, 0x71, 0xc4, 0x20, 0xc6, 0xc5, 0x58, 0x7e, 0x67,
	0x25, 0x73, 0x06, 0xde, 0x64, 0x41, 0x01, 0x06, 0x6e, 0x48, 0xdf, 0x46, 0x4d, 0x73, 0xbe, 0x43,
	0x12, 0x7d, 0x3a, 0x22, 0xec, 0x93, 0xf2, 0x0f, 0x2c, 0xf9, 0x32, 0x8b, 0xe9, 0x66, 0x19, 0xf3,
	0x7e, 0x8d, 0x66, 0x95, 0x32, 0xeb, 0xa2, 0x8b, 0x5d, 0x6b, 0xdb, 0x81, 0xff, 0xb7, 0x62, 0x69,
	0x54, 0x5b, 0x0a, 0x6d, 0xf3, 0x3d, 0xb7, 0x3f, 0xfd, 0x4c, 0x1f, 0xa0, 0x2a, 0xc3, 0x40, 0x6b,
	0xbe, 0xdf, 0x17, 0x7d, 0x80, 0xaa, 0x6c, 0x9b, 0x9e, 0xb0, 0x5d, 0x71, 0xbb, 0x84, 0xd5, 0x15,
	0x1f, 0x9c, 0xed, 0x42, 0x71, 0x60, 0xc0, 0x4a, 0x56, 0xbe, 0x42, 0x32, 0x11, 0x61, 0x11, 0x7e,
	0x84, 0x48, 0x2b, 0xff, 0xef, 0xfd, 0x57, 0x04, 0x4c, 0x44, 0x6f, 0x5c, 0xd8, 0xbb, 0x8b, 0x66,
	0xcc, 0x8d, 0xa8, 0xc7, 0x08, 0xc7, 0x4c, 0x43, 0x86, 0x69, 0xd2, 0xf1, 0xff, 0x61, 0xaf, 0x3e,
	0x57, 0x94, 0x3a, 0x30, 0xf1, 0x1d, 0x0d, 0x59, 0x2b, 0xe9, 0x78, 0xbf, 0x40, 0xd3, 0x4a, 0xd9,
	0x33, 0x69, 0x1a, 0x9f, 0xab, 0xf2, 0x9f, 0xee, 0x50, 0x2b, 0x65, 0x8e, 0x61, 0x9b, 0xb9, 0x12,
	0x3f, 0x37, 0xd0, 0x52, 0x6d, 0x7b, 0xd7, 0x2e, 0xf7, 0xf6, 0x83, 0xfb, 0xff, 0xb2, 0xbd, 0x59,
	0x9f, 0xe9, 0x85, 0xb4, 0x30, 0x10, 0x2c, 0x66, 0xd5, 0x99, 0x62, 0xc0, 0xb7, 0x6a, 0xb6, 0xc2,
	0x71, 0x9b, 0xf9, 0xe1, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x77, 0xba, 0x46, 0xaa, 0x10,
	0x00, 0x00,
}
