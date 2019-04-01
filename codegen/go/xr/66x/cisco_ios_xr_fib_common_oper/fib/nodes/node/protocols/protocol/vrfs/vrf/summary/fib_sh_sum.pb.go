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

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_vrfs_vrf_summary

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
	Prefix                        string                  `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
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
	EncapSharedLoadSharingElement *FibPlLdiCount          `protobuf:"bytes,66,opt,name=encap_shared_load_sharing_element,json=encapSharedLoadSharingElement,proto3" json:"encap_shared_load_sharing_element,omitempty"`
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
	NumberRoutesSrv6Transit       uint32                  `protobuf:"varint,91,opt,name=number_routes_srv6_transit,json=numberRoutesSrv6Transit,proto3" json:"number_routes_srv6_transit,omitempty"`
	NumberRoutesSrv6End           uint32                  `protobuf:"varint,92,opt,name=number_routes_srv6_end,json=numberRoutesSrv6End,proto3" json:"number_routes_srv6_end,omitempty"`
	NumberOfSrLabels              uint32                  `protobuf:"varint,93,opt,name=number_of_sr_labels,json=numberOfSrLabels,proto3" json:"number_of_sr_labels,omitempty"`
	SsDropPlCount                 uint32                  `protobuf:"varint,94,opt,name=ss_drop_pl_count,json=ssDropPlCount,proto3" json:"ss_drop_pl_count,omitempty"`
	PrefixMasklenDistribution     *FibShPfxMasklenDistrib `protobuf:"bytes,95,opt,name=prefix_masklen_distribution,json=prefixMasklenDistribution,proto3" json:"prefix_masklen_distribution,omitempty"`
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

func (m *FibShSum) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
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

func (m *FibShSum) GetEncapSharedLoadSharingElement() *FibPlLdiCount {
	if m != nil {
		return m.EncapSharedLoadSharingElement
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

func (m *FibShSum) GetNumberRoutesSrv6Transit() uint32 {
	if m != nil {
		return m.NumberRoutesSrv6Transit
	}
	return 0
}

func (m *FibShSum) GetNumberRoutesSrv6End() uint32 {
	if m != nil {
		return m.NumberRoutesSrv6End
	}
	return 0
}

func (m *FibShSum) GetNumberOfSrLabels() uint32 {
	if m != nil {
		return m.NumberOfSrLabels
	}
	return 0
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
	proto.RegisterType((*FibShSum_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_sum_KEYS")
	proto.RegisterType((*FibPlLdiCount)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_pl_ldi_count")
	proto.RegisterType((*FibShPfxMasklenCnt)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_pfx_masklen_cnt")
	proto.RegisterType((*FibShPfxMasklenDistrib)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_pfx_masklen_distrib")
	proto.RegisterType((*FibShSum)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_sum")
}

func init() { proto.RegisterFile("fib_sh_sum.proto", fileDescriptor_6b124c48422c0f95) }

var fileDescriptor_6b124c48422c0f95 = []byte{
	// 1495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0x6d, 0x57, 0x14, 0x47,
	0x16, 0x3e, 0x23, 0x8a, 0x50, 0x3a, 0x02, 0xcd, 0x8b, 0x0d, 0xa8, 0x8b, 0x78, 0xdc, 0x75, 0xd7,
	0x5d, 0xf4, 0xe0, 0x0a, 0xbb, 0x9a, 0xc4, 0x88, 0x60, 0x44, 0x79, 0x99, 0x34, 0x48, 0x62, 0xde,
	0xea, 0xd4, 0x74, 0xdf, 0x66, 0x2a, 0x54, 0x57, 0xf5, 0xa9, 0xaa, 0x1e, 0x86, 0x7f, 0x91, 0x93,
	0x93, 0x1f, 0x90, 0x4f, 0x7e, 0xf1, 0x27, 0xe5, 0x73, 0xfe, 0x47, 0x4e, 0x55, 0x75, 0xf7, 0x34,
	0x33, 0x23, 0x9f, 0x12, 0xbe, 0x70, 0x98, 0xfb, 0x3c, 0xf7, 0xd6, 0x73, 0xef, 0xdc, 0xba, 0xb7,
	0x06, 0x8d, 0xc7, 0xb4, 0x89, 0x55, 0x0b, 0xab, 0x2c, 0x59, 0x4a, 0xa5, 0xd0, 0xc2, 0xdb, 0x0d,
	0xa9, 0x0a, 0x05, 0xa6, 0x42, 0xe1, 0x8e, 0xc4, 0x06, 0x0e, 0x45, 0x92, 0x08, 0x8e, 0x45, 0x0a,
	0x72, 0x29, 0xa6, 0xcd, 0x25, 0x2e, 0x22, 0x50, 0xf6, 0xaf, 0x73, 0x09, 0x05, 0x53, 0xe5, 0x7f,
	0x4b, 0x6d, 0x19, 0x2b, 0xf3, 0x67, 0x49, 0x65, 0x49, 0x42, 0xe4, 0xc9, 0x22, 0x47, 0x63, 0xdd,
	0x43, 0xf0, 0x9b, 0x8d, 0x77, 0x7b, 0xde, 0x3c, 0x1a, 0x35, 0xee, 0x98, 0x93, 0x04, 0xfc, 0xda,
	0x42, 0xed, 0xde, 0x68, 0x30, 0x62, 0x0c, 0x3b, 0x24, 0x01, 0xef, 0x0e, 0xaa, 0x17, 0xc1, 0x1c,
	0xe1, 0x82, 0x25, 0x5c, 0x2d, 0x8c, 0x96, 0x34, 0x8b, 0x46, 0xda, 0x32, 0x76, 0xf8, 0x90, 0xc5,
	0x2f, 0xb7, 0x65, 0x6c, 0xa0, 0xc5, 0xf7, 0x97, 0x5c, 0x56, 0x29, 0xc3, 0x2c, 0xa2, 0x38, 0x14,
	0x19, 0xd7, 0xde, 0x4b, 0xb4, 0xa0, 0x85, 0x26, 0x0c, 0x33, 0x41, 0x22, 0xac, 0x5a, 0x44, 0x52,
	0x7e, 0x88, 0x81, 0x41, 0x02, 0x5c, 0xe3, 0xe6, 0x89, 0x06, 0x65, 0x85, 0xd4, 0x83, 0x1b, 0x96,
	0xb7, 0x25, 0x48, 0xb4, 0xe7, 0x58, 0x1b, 0x8e, 0xb4, 0x66, 0x38, 0x5e, 0x03, 0xdd, 0x3d, 0x23,
	0x8e, 0x84, 0x18, 0x24, 0xf0, 0x10, 0x94, 0x15, 0x7d, 0x31, 0xb8, 0xfd, 0x91, 0x60, 0x41, 0x49,
	0xf4, 0x56, 0x91, 0xef, 0x22, 0xa6, 0x44, 0xb7, 0x30, 0xa3, 0x4a, 0x17, 0xe1, 0x94, 0xcd, 0xac,
	0x1e, 0x4c, 0x5b, 0xbc, 0x41, 0x74, 0x6b, 0x8b, 0x2a, 0x9d, 0x47, 0x50, 0xde, 0x33, 0x74, 0x43,
	0x42, 0x98, 0x49, 0x45, 0xdb, 0x30, 0xc8, 0xf9, 0xa2, 0x75, 0x9e, 0x2d, 0x39, 0x7d, 0x01, 0x5e,
	0xa3, 0xc5, 0x94, 0x11, 0x1d, 0x0b, 0x99, 0xd8, 0x4c, 0x20, 0x1a, 0x14, 0xe6, 0x92, 0x0d, 0x73,
	0xab, 0x60, 0xee, 0x59, 0x62, 0x5f, 0xac, 0x55, 0xe4, 0x4b, 0xd0, 0xf2, 0x64, 0x50, 0x84, 0x61,
	0x97, 0x85, 0xc5, 0x07, 0x39, 0x56, 0x0a, 0x4a, 0x79, 0x2c, 0xba, 0x8e, 0x97, 0x2b, 0xe9, 0x9b,
	0x1a, 0x6e, 0xf2, 0x58, 0x0c, 0x4e, 0x7f, 0x80, 0xf3, 0x48, 0x4f, 0xfa, 0x7d, 0x01, 0x06, 0xa4,
	0x3f, 0x20, 0xcc, 0xe8, 0xa0, 0xf4, 0xfb, 0x62, 0x3d, 0x42, 0x33, 0x9d, 0x74, 0x60, 0x0e, 0xc8,
	0xfa, 0x4f, 0x76, 0xd2, 0xbe, 0x0c, 0x16, 0x0f, 0xd1, 0x4c, 0x7e, 0x31, 0xd2, 0xb8, 0x83, 0x13,
	0xa2, 0x8e, 0x18, 0x70, 0x1c, 0x72, 0xed, 0xfd, 0x0d, 0x5d, 0x31, 0x1f, 0x31, 0x03, 0x7e, 0xa8,
	0x5b, 0x79, 0x63, 0x22, 0x63, 0xda, 0xb2, 0x16, 0xef, 0xdf, 0xc8, 0xe3, 0x59, 0xd2, 0x04, 0x89,
	0x45, 0x8c, 0x53, 0x09, 0x31, 0xed, 0xe4, 0x3d, 0x57, 0x0f, 0xc6, 0x1d, 0xb2, 0x1b, 0x37, 0x72,
	0xfb, 0xe2, 0x6f, 0x43, 0x68, 0x6e, 0xc0, 0x49, 0x11, 0x55, 0x5a, 0xd2, 0xa6, 0xf7, 0x53, 0x0d,
	0x8d, 0x65, 0x9c, 0x86, 0x44, 0xe9, 0x22, 0x96, 0x5f, 0x5b, 0x18, 0xba, 0x77, 0x65, 0xf9, 0x70,
	0xe9, 0x4f, 0x1e, 0x06, 0x4b, 0x83, 0x13, 0x0e, 0xae, 0xe5, 0xe7, 0xe7, 0x92, 0xbd, 0x5f, 0x6a,
	0x68, 0xa2, 0x29, 0x05, 0x89, 0x4e, 0x89, 0xba, 0x70, 0xbe, 0xa2, 0xc6, 0x4b, 0x05, 0x85, 0xac,
	0x9f, 0x6b, 0x68, 0x3c, 0xc9, 0x98, 0xae, 0xd6, 0xca, 0x1f, 0x3a, 0x5f, 0x55, 0x63, 0xa5, 0x00,
	0xa7, 0x6a, 0xf1, 0xf7, 0x69, 0x84, 0xba, 0x03, 0xd6, 0x9b, 0x41, 0xc3, 0xb9, 0xb0, 0x65, 0x3b,
	0x17, 0xf3, 0x4f, 0xde, 0x1c, 0x1a, 0x55, 0x0a, 0xeb, 0x26, 0xc3, 0x34, 0xf2, 0x1f, 0xd9, 0x4e,
	0xb9, 0xac, 0xd4, 0x7e, 0x93, 0x6d, 0x46, 0xde, 0x6d, 0x54, 0x2f, 0x31, 0x9c, 0x6a, 0xe9, 0xff,
	0xd7, 0x75, 0x5c, 0x8e, 0x37, 0xb4, 0xcc, 0xdd, 0xcd, 0xcc, 0xa5, 0x91, 0xff, 0xb8, 0x70, 0x3f,
	0x90, 0xf1, 0x66, 0xe4, 0xf9, 0x68, 0xc4, 0x62, 0x06, 0x5a, 0xb1, 0xd0, 0xb0, 0x81, 0x36, 0x23,
	0xef, 0x2e, 0xba, 0x66, 0xef, 0x44, 0x93, 0x30, 0xc2, 0x43, 0xca, 0x0f, 0xfd, 0x55, 0x2b, 0xaa,
	0x6e, 0xac, 0x6b, 0x85, 0xd1, 0x7b, 0x80, 0x26, 0x63, 0x21, 0x8f, 0x89, 0x8c, 0x2a, 0xd3, 0x54,
	0xf9, 0xff, 0xb3, 0xb1, 0xbc, 0x2e, 0x54, 0xde, 0xb7, 0x19, 0x34, 0x2c, 0x45, 0x66, 0x86, 0xf6,
	0xff, 0xdd, 0x79, 0xee, 0x93, 0xf7, 0x1c, 0xdd, 0x74, 0xe9, 0x62, 0xca, 0x71, 0xca, 0x48, 0x08,
	0x38, 0x11, 0x11, 0x8d, 0x69, 0x48, 0x34, 0x15, 0x5c, 0xf9, 0x4f, 0x2c, 0x7d, 0xce, 0x91, 0x36,
	0x79, 0xc3, 0x50, 0xb6, 0xab, 0x0c, 0xef, 0x21, 0x9a, 0x52, 0x9a, 0x30, 0xc8, 0xbf, 0x5e, 0x1c,
	0x01, 0x03, 0x73, 0xd0, 0x53, 0x27, 0xc6, 0x62, 0xae, 0xf2, 0xeb, 0x0e, 0xf1, 0x96, 0xd1, 0xf4,
	0xa0, 0x6d, 0xa0, 0xfc, 0x4f, 0xdc, 0xdd, 0x67, 0x7d, 0xe3, 0x5f, 0x79, 0x2b, 0xe8, 0xfa, 0x29,
	0x9f, 0xca, 0xe6, 0xf8, 0xd4, 0x6e, 0x8e, 0xe9, 0x8a, 0x57, 0x65, 0x5b, 0xac, 0xa1, 0x5b, 0x3d,
	0xfb, 0x07, 0x7a, 0xb6, 0xd8, 0x67, 0x2e, 0xc3, 0x53, 0x8b, 0x07, 0x4e, 0xed, 0xb0, 0xf7, 0x35,
	0x74, 0x0b, 0x3a, 0x21, 0xcb, 0xba, 0xa3, 0xb3, 0x47, 0xba, 0xff, 0x6c, 0xa1, 0x76, 0xef, 0xca,
	0x32, 0xf9, 0x4b, 0x7a, 0xba, 0xba, 0x97, 0x83, 0xf9, 0x52, 0x48, 0xff, 0x92, 0xf4, 0x7e, 0xad,
	0xa1, 0xf9, 0xea, 0x64, 0xee, 0x55, 0xf9, 0xf9, 0x79, 0xa9, 0xf4, 0x55, 0x39, 0xf6, 0x7b, 0x24,
	0x7e, 0xa8, 0xa1, 0xdb, 0xa1, 0x14, 0x4a, 0xe1, 0xb3, 0x84, 0x3e, 0x3f, 0x2f, 0xa1, 0x37, 0xad,
	0x96, 0xbd, 0xb3, 0xd4, 0x02, 0x0f, 0x49, 0x7a, 0xa6, 0xda, 0xb5, 0x73, 0x53, 0x6b, 0xb5, 0x7c,
	0x54, 0xed, 0xbf, 0xd0, 0x04, 0x03, 0xd2, 0x06, 0x85, 0x33, 0x05, 0x51, 0xde, 0xde, 0x2f, 0x6c,
	0x7b, 0x8f, 0x39, 0xe0, 0xad, 0x82, 0xc8, 0xf5, 0xf4, 0x7d, 0x34, 0x21, 0x41, 0x82, 0x12, 0xac,
	0x0d, 0x18, 0xb8, 0x96, 0x14, 0x94, 0xbf, 0xee, 0xf6, 0x61, 0x09, 0x6c, 0x38, 0xbb, 0xb9, 0xb0,
	0x82, 0x45, 0x38, 0xe3, 0xbd, 0x0e, 0x1b, 0xee, 0xc2, 0x0a, 0x16, 0xbd, 0xe5, 0xfd, 0x3e, 0x1c,
	0x8e, 0x07, 0xf8, 0xbc, 0x74, 0x3e, 0x1c, 0x8e, 0xfb, 0x7c, 0xee, 0xa3, 0x89, 0x7e, 0xfe, 0x17,
	0x4e, 0x54, 0xd6, 0x4b, 0xfe, 0x3b, 0x1a, 0x0b, 0x21, 0xc6, 0x76, 0x90, 0xe1, 0x48, 0x8a, 0x54,
	0xf9, 0xaf, 0x2c, 0xb5, 0x1e, 0x42, 0x1c, 0x18, 0xeb, 0xba, 0x31, 0x9a, 0x97, 0xac, 0xe1, 0xb5,
	0x41, 0x2a, 0x2a, 0x38, 0x4e, 0xa8, 0x4a, 0x88, 0x0e, 0x5b, 0xa7, 0x1c, 0x37, 0xed, 0x08, 0xb9,
	0x11, 0x42, 0x7c, 0xe0, 0x68, 0xdb, 0x39, 0xab, 0x12, 0x67, 0x15, 0xf9, 0x6e, 0xb4, 0xe1, 0x90,
	0x84, 0x2d, 0xc0, 0x3c, 0x4b, 0x4a, 0x8d, 0xaf, 0xdd, 0xc3, 0xcb, 0xe1, 0x2f, 0x0c, 0xbc, 0x93,
	0x25, 0x85, 0xd0, 0x27, 0x68, 0x16, 0x3a, 0x54, 0x69, 0xd3, 0x32, 0xf9, 0xf7, 0x23, 0xa1, 0x4d,
	0x95, 0x9d, 0xaf, 0x6f, 0xac, 0xe7, 0xf5, 0x82, 0xb0, 0x65, 0xf1, 0xa0, 0x80, 0xcd, 0xbb, 0xc5,
	0x74, 0x41, 0x04, 0x31, 0xc9, 0x58, 0xb9, 0x41, 0xb7, 0x5c, 0x49, 0x62, 0xda, 0x5c, 0x77, 0x80,
	0x9b, 0xaf, 0x66, 0xd8, 0xf5, 0xb3, 0x71, 0xf5, 0x65, 0xb4, 0xed, 0x86, 0x5d, 0xaf, 0xe7, 0x76,
	0xf7, 0xa5, 0x64, 0x7e, 0x6a, 0x40, 0x47, 0xe3, 0x96, 0xa9, 0xcb, 0x8e, 0xa5, 0x8f, 0x18, 0xc3,
	0x2b, 0x53, 0x83, 0x87, 0x68, 0x8a, 0xf2, 0x50, 0x24, 0xa9, 0xad, 0x43, 0x97, 0xb7, 0xeb, 0x66,
	0x7d, 0x17, 0xdb, 0x29, 0x3c, 0xfe, 0x89, 0xc6, 0xed, 0xf7, 0x96, 0x99, 0x65, 0x81, 0x35, 0x4d,
	0x40, 0xfa, 0x0d, 0xd7, 0x92, 0x5d, 0xfb, 0xbe, 0x31, 0x9b, 0x5c, 0x15, 0x13, 0xc7, 0x38, 0x95,
	0x22, 0x04, 0xb3, 0x5e, 0x2d, 0xf9, 0x4b, 0x97, 0xab, 0x41, 0x1a, 0x0e, 0x70, 0xec, 0x87, 0x68,
	0x2a, 0x21, 0x1d, 0xdc, 0x17, 0x3c, 0x70, 0x52, 0x12, 0xd2, 0x09, 0x7a, 0xe2, 0x3f, 0x40, 0x93,
	0x34, 0x49, 0x85, 0xa2, 0x96, 0x5d, 0x3e, 0x02, 0xf7, 0x72, 0xed, 0x25, 0x54, 0x3c, 0x03, 0x4d,
	0x3b, 0x42, 0x47, 0x03, 0x8f, 0xcc, 0x43, 0xbf, 0xa0, 0xef, 0x3b, 0x3d, 0x05, 0x50, 0x92, 0x1f,
	0xa3, 0xeb, 0x21, 0xc4, 0x0c, 0x37, 0x19, 0x96, 0x10, 0x9e, 0x84, 0x0c, 0x22, 0x9c, 0xaf, 0xdc,
	0xb7, 0xd6, 0x65, 0xca, 0xc0, 0x6b, 0x2c, 0xc8, 0xc1, 0xc0, 0x2d, 0xe0, 0x3b, 0xa8, 0x6e, 0xee,
	0x77, 0x93, 0x84, 0x47, 0xc7, 0x84, 0x1d, 0x29, 0xff, 0xc0, 0x92, 0xaf, 0xb2, 0x88, 0xae, 0x15,
	0x36, 0xef, 0x3f, 0x68, 0x52, 0x29, 0x53, 0x17, 0x9d, 0x77, 0xad, 0x1d, 0x07, 0xfe, 0x57, 0x79,
	0x69, 0x54, 0x43, 0x0a, 0x6d, 0xe3, 0xbd, 0xb0, 0xbf, 0xdd, 0xcc, 0x1c, 0xa0, 0x2a, 0xc5, 0x40,
	0x2b, 0xba, 0xbf, 0xce, 0xe7, 0x00, 0x55, 0xe9, 0x06, 0x3d, 0x25, 0xbb, 0xe4, 0xb6, 0x09, 0xab,
	0x7a, 0xbc, 0x73, 0xb2, 0x73, 0x8f, 0x03, 0x03, 0x96, 0x6e, 0xc5, 0x11, 0x92, 0x89, 0x10, 0x8b,
	0xe6, 0x8f, 0x10, 0x6a, 0xe5, 0x7f, 0xd3, 0x3d, 0x22, 0x60, 0x22, 0xdc, 0x75, 0x66, 0xef, 0x29,
	0x9a, 0xcb, 0xdf, 0xde, 0xae, 0x1e, 0x58, 0xc9, 0xf6, 0x0a, 0xd6, 0x92, 0x70, 0x45, 0xb5, 0xff,
	0xad, 0xbb, 0x00, 0x8e, 0xe1, 0x8a, 0xb2, 0x27, 0xdb, 0x2b, 0xfb, 0x0e, 0x36, 0x3f, 0x14, 0x06,
	0x38, 0x03, 0x8f, 0xfc, 0xef, 0xf2, 0x39, 0xd2, 0xe3, 0xb8, 0xc1, 0x23, 0x53, 0xaf, 0xee, 0x6b,
	0x5f, 0x49, 0xcc, 0x48, 0x13, 0x98, 0xf2, 0xbf, 0x3f, 0xfd, 0xdc, 0xdf, 0x93, 0x5b, 0xd6, 0xee,
	0xfd, 0x03, 0x8d, 0x2b, 0x65, 0x27, 0x81, 0x19, 0xb7, 0xae, 0xb6, 0x3f, 0xb8, 0x51, 0xa2, 0x94,
	0xb9, 0xfc, 0x0d, 0xe6, 0x0a, 0xfb, 0xa1, 0x86, 0xe6, 0x2b, 0x97, 0xaa, 0xf2, 0x9b, 0xc0, 0xb6,
	0x99, 0x8f, 0xed, 0x22, 0x38, 0x3a, 0x8f, 0x97, 0x6d, 0x7e, 0x6e, 0x30, 0x9b, 0x96, 0x17, 0x98,
	0x01, 0x5f, 0xaf, 0xa8, 0x69, 0x0e, 0xdb, 0x80, 0x8f, 0xfe, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0x58, 0x07, 0xbf, 0xb3, 0x10, 0x00, 0x00,
}
