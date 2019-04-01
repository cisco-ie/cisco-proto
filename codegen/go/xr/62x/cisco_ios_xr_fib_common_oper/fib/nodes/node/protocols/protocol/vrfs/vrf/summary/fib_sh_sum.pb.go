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

type FibShSum struct {
	Prefix                        [][]byte       `protobuf:"bytes,50,rep,name=prefix,proto3" json:"prefix,omitempty"`
	SsTblId                       uint32         `protobuf:"varint,51,opt,name=ss_tbl_id,json=ssTblId,proto3" json:"ss_tbl_id,omitempty"`
	SsTblIdPtr                    uint32         `protobuf:"varint,52,opt,name=ss_tbl_id_ptr,json=ssTblIdPtr,proto3" json:"ss_tbl_id_ptr,omitempty"`
	SsVrfId                       uint32         `protobuf:"varint,53,opt,name=ss_vrf_id,json=ssVrfId,proto3" json:"ss_vrf_id,omitempty"`
	SsVrId                        uint32         `protobuf:"varint,54,opt,name=ss_vr_id,json=ssVrId,proto3" json:"ss_vr_id,omitempty"`
	LoadBalancing                 string         `protobuf:"bytes,55,opt,name=load_balancing,json=loadBalancing,proto3" json:"load_balancing,omitempty"`
	ForwardingElements            uint32         `protobuf:"varint,56,opt,name=forwarding_elements,json=forwardingElements,proto3" json:"forwarding_elements,omitempty"`
	Routes                        uint32         `protobuf:"varint,57,opt,name=routes,proto3" json:"routes,omitempty"`
	PrefixInPlaceModifications    uint32         `protobuf:"varint,58,opt,name=prefix_in_place_modifications,json=prefixInPlaceModifications,proto3" json:"prefix_in_place_modifications,omitempty"`
	StalePrefixDeletes            uint32         `protobuf:"varint,59,opt,name=stale_prefix_deletes,json=stalePrefixDeletes,proto3" json:"stale_prefix_deletes,omitempty"`
	LoadSharingElements           uint32         `protobuf:"varint,60,opt,name=load_sharing_elements,json=loadSharingElements,proto3" json:"load_sharing_elements,omitempty"`
	LoadSharingReferences         uint64         `protobuf:"varint,61,opt,name=load_sharing_references,json=loadSharingReferences,proto3" json:"load_sharing_references,omitempty"`
	TotalLoadShareElementBytes    uint32         `protobuf:"varint,62,opt,name=total_load_share_element_bytes,json=totalLoadShareElementBytes,proto3" json:"total_load_share_element_bytes,omitempty"`
	ExclusiveLoadSharingElement   *FibPlLdiCount `protobuf:"bytes,63,opt,name=exclusive_load_sharing_element,json=exclusiveLoadSharingElement,proto3" json:"exclusive_load_sharing_element,omitempty"`
	SharedLoadSharingElement      *FibPlLdiCount `protobuf:"bytes,64,opt,name=shared_load_sharing_element,json=sharedLoadSharingElement,proto3" json:"shared_load_sharing_element,omitempty"`
	CrossSharedLoadSharingElement *FibPlLdiCount `protobuf:"bytes,65,opt,name=cross_shared_load_sharing_element,json=crossSharedLoadSharingElement,proto3" json:"cross_shared_load_sharing_element,omitempty"`
	LabelSharedLoadSharingElement *FibPlLdiCount `protobuf:"bytes,66,opt,name=label_shared_load_sharing_element,json=labelSharedLoadSharingElement,proto3" json:"label_shared_load_sharing_element,omitempty"`
	LeavesUsedBytes               uint32         `protobuf:"varint,67,opt,name=leaves_used_bytes,json=leavesUsedBytes,proto3" json:"leaves_used_bytes,omitempty"`
	ReresolveEntries              uint32         `protobuf:"varint,68,opt,name=reresolve_entries,json=reresolveEntries,proto3" json:"reresolve_entries,omitempty"`
	OldUnresolveEntries           uint32         `protobuf:"varint,69,opt,name=old_unresolve_entries,json=oldUnresolveEntries,proto3" json:"old_unresolve_entries,omitempty"`
	NewUnresolveEntries           uint32         `protobuf:"varint,70,opt,name=new_unresolve_entries,json=newUnresolveEntries,proto3" json:"new_unresolve_entries,omitempty"`
	UnresolveEntries              uint32         `protobuf:"varint,71,opt,name=unresolve_entries,json=unresolveEntries,proto3" json:"unresolve_entries,omitempty"`
	CefRouteDrops                 uint32         `protobuf:"varint,72,opt,name=cef_route_drops,json=cefRouteDrops,proto3" json:"cef_route_drops,omitempty"`
	CefVersionMismatchRouteDrops  uint64         `protobuf:"varint,73,opt,name=cef_version_mismatch_route_drops,json=cefVersionMismatchRouteDrops,proto3" json:"cef_version_mismatch_route_drops,omitempty"`
	DeleteCacheNumEntries         uint32         `protobuf:"varint,74,opt,name=delete_cache_num_entries,json=deleteCacheNumEntries,proto3" json:"delete_cache_num_entries,omitempty"`
	ExistingLeavesRevisions       uint32         `protobuf:"varint,75,opt,name=existing_leaves_revisions,json=existingLeavesRevisions,proto3" json:"existing_leaves_revisions,omitempty"`
	FibDefaultPrefix              uint32         `protobuf:"varint,76,opt,name=fib_default_prefix,json=fibDefaultPrefix,proto3" json:"fib_default_prefix,omitempty"`
	FibDefaultPrefixMaskLength    uint32         `protobuf:"varint,77,opt,name=fib_default_prefix_mask_length,json=fibDefaultPrefixMaskLength,proto3" json:"fib_default_prefix_mask_length,omitempty"`
	NextHops                      uint32         `protobuf:"varint,78,opt,name=next_hops,json=nextHops,proto3" json:"next_hops,omitempty"`
	IncompleteNextHops            uint32         `protobuf:"varint,79,opt,name=incomplete_next_hops,json=incompleteNextHops,proto3" json:"incomplete_next_hops,omitempty"`
	ResolutionTimer               uint32         `protobuf:"varint,80,opt,name=resolution_timer,json=resolutionTimer,proto3" json:"resolution_timer,omitempty"`
	SlowProcessTimer              uint32         `protobuf:"varint,81,opt,name=slow_process_timer,json=slowProcessTimer,proto3" json:"slow_process_timer,omitempty"`
	MaxResolutionTimer            uint32         `protobuf:"varint,82,opt,name=max_resolution_timer,json=maxResolutionTimer,proto3" json:"max_resolution_timer,omitempty"`
	ImpositionPrefixes            uint32         `protobuf:"varint,83,opt,name=imposition_prefixes,json=impositionPrefixes,proto3" json:"imposition_prefixes,omitempty"`
	ExtendedPrefixes              uint32         `protobuf:"varint,84,opt,name=extended_prefixes,json=extendedPrefixes,proto3" json:"extended_prefixes,omitempty"`
	CeflBlRecycledRoutes          uint32         `protobuf:"varint,85,opt,name=cefl_bl_recycled_routes,json=ceflBlRecycledRoutes,proto3" json:"cefl_bl_recycled_routes,omitempty"`
	LdiBackwalks                  uint32         `protobuf:"varint,86,opt,name=ldi_backwalks,json=ldiBackwalks,proto3" json:"ldi_backwalks,omitempty"`
	SsProtRouteCount              uint32         `protobuf:"varint,87,opt,name=ss_prot_route_count,json=ssProtRouteCount,proto3" json:"ss_prot_route_count,omitempty"`
	LispEidPrefixes               uint32         `protobuf:"varint,88,opt,name=lisp_eid_prefixes,json=lispEidPrefixes,proto3" json:"lisp_eid_prefixes,omitempty"`
	LispEidValidPrefixes          uint32         `protobuf:"varint,89,opt,name=lisp_eid_valid_prefixes,json=lispEidValidPrefixes,proto3" json:"lisp_eid_valid_prefixes,omitempty"`
	LispRlocObjects               uint32         `protobuf:"varint,90,opt,name=lisp_rloc_objects,json=lispRlocObjects,proto3" json:"lisp_rloc_objects,omitempty"`
	SsVxlanLtepIfh                string         `protobuf:"bytes,91,opt,name=ss_vxlan_ltep_ifh,json=ssVxlanLtepIfh,proto3" json:"ss_vxlan_ltep_ifh,omitempty"`
	SsDropPlCount                 uint32         `protobuf:"varint,92,opt,name=ss_drop_pl_count,json=ssDropPlCount,proto3" json:"ss_drop_pl_count,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}       `json:"-"`
	XXX_unrecognized              []byte         `json:"-"`
	XXX_sizecache                 int32          `json:"-"`
}

func (m *FibShSum) Reset()         { *m = FibShSum{} }
func (m *FibShSum) String() string { return proto.CompactTextString(m) }
func (*FibShSum) ProtoMessage()    {}
func (*FibShSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b124c48422c0f95, []int{2}
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

func (m *FibShSum) GetPrefix() [][]byte {
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

func init() {
	proto.RegisterType((*FibShSum_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_sum_KEYS")
	proto.RegisterType((*FibPlLdiCount)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_pl_ldi_count")
	proto.RegisterType((*FibShSum)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.vrfs.vrf.summary.fib_sh_sum")
}

func init() { proto.RegisterFile("fib_sh_sum.proto", fileDescriptor_6b124c48422c0f95) }

var fileDescriptor_6b124c48422c0f95 = []byte{
	// 1292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x6d, 0x73, 0x13, 0x37,
	0x10, 0x9e, 0x14, 0x08, 0x89, 0x4a, 0x20, 0xb9, 0x04, 0x38, 0x5e, 0x27, 0x84, 0xa1, 0x85, 0xd2,
	0x1a, 0x06, 0x0a, 0x69, 0xe9, 0x0b, 0x25, 0x24, 0x14, 0x43, 0x12, 0xdc, 0x0b, 0xa4, 0xa5, 0xed,
	0x8c, 0x46, 0xd6, 0xed, 0x61, 0x35, 0x3a, 0xe9, 0x46, 0x2b, 0x3b, 0xce, 0xbf, 0xe9, 0x27, 0xbe,
	0xf4, 0x57, 0xf4, 0x9f, 0x75, 0x24, 0xdd, 0x9d, 0x1d, 0xdb, 0xe4, 0x53, 0x87, 0x2f, 0x99, 0x78,
	0x9f, 0xe7, 0xd9, 0x7b, 0x76, 0x4f, 0xbb, 0x27, 0x32, 0x9f, 0x89, 0x36, 0xc5, 0x0e, 0xc5, 0x6e,
	0xde, 0x28, 0x8c, 0xb6, 0x3a, 0x7a, 0xc5, 0x05, 0x72, 0x4d, 0x85, 0x46, 0xda, 0x37, 0xd4, 0xc1,
	0x5c, 0xe7, 0xb9, 0x56, 0x54, 0x17, 0x60, 0x1a, 0x99, 0x68, 0x37, 0x94, 0x4e, 0x01, 0xfd, 0xdf,
	0x20, 0xe1, 0x5a, 0x62, 0xfd, 0x5f, 0xa3, 0x67, 0x32, 0x74, 0x7f, 0x1a, 0xd8, 0xcd, 0x73, 0x66,
	0x0e, 0x56, 0x14, 0x39, 0x33, 0x78, 0x08, 0x7d, 0xb9, 0xf1, 0x76, 0x27, 0xba, 0x44, 0x66, 0x9d,
	0x9c, 0x2a, 0x96, 0x43, 0x3c, 0xb5, 0x3c, 0x75, 0x73, 0x36, 0x99, 0x71, 0x81, 0x6d, 0x96, 0x43,
	0x74, 0x9d, 0xcc, 0x55, 0xc9, 0x02, 0xe1, 0x13, 0x4f, 0x38, 0x55, 0x05, 0x3d, 0xe9, 0x02, 0x99,
	0xe9, 0x99, 0x2c, 0xe0, 0xc7, 0x3c, 0x7e, 0xb2, 0x67, 0x32, 0x07, 0xad, 0xbc, 0x3f, 0x11, 0xaa,
	0x2a, 0x24, 0x95, 0xa9, 0xa0, 0x5c, 0x77, 0x95, 0x8d, 0x9e, 0x91, 0x65, 0xab, 0x2d, 0x93, 0x54,
	0x6a, 0x96, 0x52, 0xec, 0x30, 0x23, 0xd4, 0x3b, 0x0a, 0x12, 0x72, 0x50, 0x96, 0xb6, 0x0f, 0x2c,
	0xa0, 0x37, 0x32, 0x97, 0x5c, 0xf6, 0xbc, 0x4d, 0xcd, 0xd2, 0x9d, 0xc0, 0xda, 0x08, 0xa4, 0x35,
	0xc7, 0x89, 0x5a, 0xe4, 0xc6, 0x11, 0x79, 0x0c, 0x64, 0x60, 0x40, 0x71, 0x40, 0x6f, 0xfa, 0x78,
	0x72, 0xed, 0x03, 0xc9, 0x92, 0x9a, 0x18, 0xad, 0x92, 0x38, 0x64, 0x2c, 0x98, 0xed, 0x50, 0x29,
	0xd0, 0x56, 0xe9, 0xd0, 0x57, 0x36, 0x97, 0x9c, 0xf5, 0x78, 0x8b, 0xd9, 0xce, 0xa6, 0x40, 0x5b,
	0x66, 0xc0, 0xe8, 0x31, 0xb9, 0x6c, 0x80, 0x77, 0x0d, 0x8a, 0x1e, 0x4c, 0x12, 0x1f, 0xf7, 0xe2,
	0x0b, 0x35, 0x67, 0x2c, 0xc1, 0x0b, 0xb2, 0x52, 0x48, 0x66, 0x33, 0x6d, 0x72, 0x5f, 0x09, 0xa4,
	0x93, 0xd2, 0x9c, 0xf0, 0x69, 0xae, 0x56, 0xcc, 0x1d, 0x4f, 0x1c, 0xcb, 0xb5, 0x4a, 0x62, 0x03,
	0xd6, 0x1c, 0x4c, 0xca, 0x30, 0x1d, 0xaa, 0xf0, 0xf8, 0x24, 0xe1, 0x50, 0x43, 0x85, 0xca, 0xf4,
	0x40, 0x78, 0x72, 0xa8, 0x7c, 0xd7, 0xc3, 0xa6, 0xca, 0xf4, 0xe4, 0xf2, 0x27, 0x88, 0x67, 0x46,
	0xca, 0x1f, 0x4b, 0x30, 0xa1, 0xfc, 0x09, 0x69, 0x66, 0x27, 0x95, 0x3f, 0x96, 0xeb, 0x3e, 0x39,
	0xd7, 0x2f, 0x26, 0xd6, 0x40, 0xbc, 0x7e, 0xb1, 0x5f, 0x8c, 0x55, 0xb0, 0xf2, 0xef, 0x22, 0x21,
	0x83, 0xc9, 0x88, 0xce, 0x91, 0xe9, 0xc2, 0x40, 0x26, 0xfa, 0xf1, 0xbd, 0xe5, 0x63, 0x37, 0x4f,
	0x25, 0xe5, 0xaf, 0xe8, 0x22, 0x99, 0x45, 0xa4, 0xb6, 0x2d, 0xa9, 0x48, 0xe3, 0xfb, 0x3e, 0xdd,
	0x49, 0xc4, 0xd7, 0x6d, 0xd9, 0x4c, 0xa3, 0x6b, 0x64, 0xae, 0xc6, 0x68, 0x61, 0x4d, 0xfc, 0xb5,
	0xc7, 0x49, 0x89, 0xb7, 0xac, 0x29, 0xe5, 0x6e, 0x58, 0x44, 0x1a, 0x3f, 0xa8, 0xe4, 0xbb, 0x26,
	0x6b, 0xa6, 0x51, 0x4c, 0x66, 0x3c, 0xe6, 0xa0, 0x87, 0x1e, 0x9a, 0x76, 0x50, 0x33, 0x8d, 0x6e,
	0x90, 0xd3, 0xbe, 0x98, 0x36, 0x93, 0x4c, 0x71, 0xa1, 0xde, 0xc5, 0xab, 0x7e, 0xca, 0xe6, 0x5c,
	0x74, 0xad, 0x0a, 0x46, 0x77, 0xc8, 0x62, 0xa6, 0xcd, 0x3e, 0x33, 0xe9, 0xd0, 0x18, 0x60, 0xfc,
	0x8d, 0xcf, 0x15, 0x0d, 0xa0, 0xba, 0x51, 0xe7, 0xc8, 0xb4, 0xd1, 0x5d, 0x37, 0x6d, 0xdf, 0x86,
	0xe7, 0x85, 0x5f, 0xd1, 0x13, 0x72, 0x25, 0x94, 0x4b, 0x85, 0xa2, 0x85, 0x64, 0x1c, 0x68, 0xae,
	0x53, 0x91, 0x09, 0xce, 0xac, 0xd0, 0x0a, 0xe3, 0x47, 0x9e, 0x7e, 0x31, 0x90, 0x9a, 0xaa, 0xe5,
	0x28, 0x5b, 0xc3, 0x8c, 0xe8, 0x2e, 0x59, 0x42, 0xcb, 0x24, 0xd0, 0x32, 0x51, 0x0a, 0x12, 0xdc,
	0x83, 0xbe, 0x0b, 0x66, 0x3c, 0xd6, 0xf2, 0xd0, 0x7a, 0x40, 0xa2, 0x7b, 0xe4, 0xec, 0xa4, 0x31,
	0xc6, 0xf8, 0xfb, 0xf0, 0xd2, 0xe4, 0xd8, 0xdc, 0x62, 0xf4, 0x90, 0x9c, 0x3f, 0xa4, 0x19, 0x1a,
	0xf9, 0x1f, 0xfc, 0xc8, 0x9f, 0x1d, 0x52, 0x0d, 0x8d, 0xf9, 0x1a, 0xb9, 0x3a, 0xb2, 0x38, 0x60,
	0x64, 0xfd, 0xfc, 0x18, 0x2a, 0x3c, 0xb4, 0x31, 0xe0, 0xd0, 0xf2, 0x79, 0x3f, 0x45, 0xae, 0x42,
	0x9f, 0xcb, 0xee, 0xe0, 0xcc, 0x8f, 0x58, 0x8f, 0x1f, 0x2f, 0x4f, 0xdd, 0xfc, 0xf4, 0x1e, 0x6b,
	0xfc, 0xcf, 0x4b, 0xbc, 0x31, 0xba, 0x50, 0x93, 0x4b, 0xb5, 0x91, 0xf1, 0xed, 0x16, 0xfd, 0x3d,
	0x45, 0x2e, 0x0d, 0x8f, 0xd4, 0xa8, 0xcb, 0x9f, 0x3e, 0x96, 0xcb, 0x18, 0xeb, 0x79, 0x1d, 0xb1,
	0xf8, 0xcf, 0x14, 0xb9, 0xc6, 0x8d, 0x46, 0xa4, 0x47, 0x19, 0x7d, 0xf2, 0xb1, 0x8c, 0x5e, 0xf1,
	0x5e, 0x76, 0x8e, 0x72, 0x2b, 0x59, 0x1b, 0xe4, 0x91, 0x6e, 0xd7, 0x3e, 0x9a, 0x5b, 0xef, 0xe5,
	0x83, 0x6e, 0xbf, 0x20, 0x0b, 0x12, 0x58, 0x0f, 0x90, 0x76, 0x11, 0xd2, 0xf2, 0x78, 0x3f, 0xf5,
	0xc7, 0xfb, 0x4c, 0x00, 0xde, 0x20, 0xa4, 0xe1, 0x4c, 0xdf, 0x26, 0x0b, 0x06, 0x0c, 0xa0, 0x96,
	0x3d, 0xa0, 0xa0, 0xac, 0x11, 0x80, 0xf1, 0xba, 0xe7, 0xce, 0xd7, 0xc0, 0x46, 0x88, 0xbb, 0x81,
	0xd5, 0x32, 0xa5, 0x5d, 0x35, 0x2a, 0xd8, 0x08, 0x03, 0xab, 0x65, 0xfa, 0x46, 0x8d, 0x6b, 0x14,
	0xec, 0x4f, 0xd0, 0x3c, 0x0b, 0x1a, 0x05, 0xfb, 0x63, 0x9a, 0xdb, 0x64, 0x61, 0x9c, 0xff, 0x73,
	0x30, 0xd5, 0x1d, 0x25, 0x7f, 0x46, 0xce, 0x70, 0xc8, 0xa8, 0x5f, 0x64, 0x34, 0x35, 0xba, 0xc0,
	0xf8, 0xb9, 0xa7, 0xce, 0x71, 0xc8, 0x12, 0x17, 0x5d, 0x77, 0x41, 0x77, 0x05, 0x71, 0xbc, 0x1e,
	0x18, 0x14, 0x5a, 0xd1, 0x5c, 0x60, 0xce, 0x2c, 0xef, 0x1c, 0x12, 0x36, 0xfd, 0x0a, 0xb9, 0xcc,
	0x21, 0xdb, 0x0d, 0xb4, 0xad, 0x92, 0x35, 0x94, 0x67, 0x95, 0xc4, 0x61, 0xb5, 0x51, 0xce, 0x78,
	0x07, 0xa8, 0xea, 0xe6, 0xb5, 0xc7, 0x17, 0xe1, 0x8b, 0x19, 0xf0, 0xa7, 0x0e, 0xde, 0xee, 0xe6,
	0x95, 0xd1, 0x47, 0xe4, 0x02, 0xf4, 0x05, 0x5a, 0x77, 0x64, 0xca, 0xf7, 0x63, 0xa0, 0x27, 0xd0,
	0xef, 0xd7, 0x97, 0x5e, 0x79, 0xbe, 0x22, 0x6c, 0x7a, 0x3c, 0xa9, 0xe0, 0xe8, 0x4b, 0x12, 0xb9,
	0x53, 0x90, 0x42, 0xc6, 0xba, 0xd2, 0x96, 0x2b, 0x36, 0xde, 0x0c, 0x2d, 0xc9, 0x44, 0x7b, 0x3d,
	0x00, 0x61, 0xbf, 0xba, 0x65, 0x37, 0xce, 0xa6, 0x39, 0xc3, 0x3d, 0x2a, 0x41, 0xbd, 0xb3, 0x9d,
	0x78, 0x2b, 0x2c, 0xbb, 0x51, 0xe5, 0x16, 0xc3, 0xbd, 0x4d, 0xcf, 0xf0, 0x77, 0x44, 0xe8, 0x5b,
	0xda, 0x71, 0x7d, 0xd9, 0xf6, 0xf4, 0x19, 0x17, 0x78, 0xee, 0x7a, 0x70, 0x97, 0x2c, 0x09, 0xc5,
	0x75, 0x5e, 0xf8, 0x3e, 0x0c, 0x78, 0xaf, 0xc2, 0xae, 0x1f, 0x60, 0xdb, 0x95, 0xe2, 0x16, 0x99,
	0xf7, 0xef, 0xad, 0xeb, 0x3e, 0x16, 0xd4, 0x8a, 0x1c, 0x4c, 0xdc, 0x0a, 0x47, 0x72, 0x10, 0x7f,
	0xed, 0xc2, 0xae, 0x56, 0x94, 0x7a, 0x9f, 0x16, 0x46, 0x73, 0x70, 0x9f, 0x57, 0x4f, 0xfe, 0x25,
	0xd4, 0xea, 0x90, 0x56, 0x00, 0x02, 0xfb, 0x2e, 0x59, 0xca, 0x59, 0x9f, 0x8e, 0x25, 0x4f, 0x82,
	0x95, 0x9c, 0xf5, 0x93, 0x91, 0xfc, 0x77, 0xc8, 0xa2, 0xc8, 0x0b, 0x8d, 0xc2, 0xb3, 0x43, 0x73,
	0x00, 0xe3, 0x9d, 0xd2, 0x7b, 0x0d, 0xb5, 0x4a, 0xc4, 0x1d, 0x47, 0xe8, 0x5b, 0x50, 0xa9, 0xbb,
	0xa1, 0x55, 0xf4, 0xd7, 0xc1, 0x4f, 0x05, 0xd4, 0xe4, 0x07, 0xe4, 0x3c, 0x87, 0x4c, 0xd2, 0xb6,
	0xa4, 0x06, 0xf8, 0x01, 0x97, 0x90, 0xd2, 0xf2, 0x93, 0xfb, 0xc6, 0x4b, 0x96, 0x1c, 0xbc, 0x26,
	0x93, 0x12, 0x4c, 0xc2, 0x07, 0xf8, 0x3a, 0x99, 0x73, 0xf3, 0xdd, 0x66, 0x7c, 0x6f, 0x9f, 0xc9,
	0x3d, 0x8c, 0x77, 0x3d, 0xf9, 0x94, 0x4c, 0xc5, 0x5a, 0x15, 0x8b, 0xbe, 0x22, 0x8b, 0x88, 0xae,
	0x2f, 0xb6, 0x3c, 0xb5, 0x7e, 0x1d, 0xc4, 0xbf, 0x96, 0xad, 0xc1, 0x96, 0xd1, 0xd6, 0xe7, 0x7b,
	0xea, 0x2f, 0xdd, 0x6e, 0x0f, 0x08, 0x2c, 0x28, 0x88, 0x21, 0xdf, 0xbf, 0x95, 0x7b, 0x40, 0x60,
	0xb1, 0x21, 0x0e, 0xd9, 0xae, 0xb9, 0x3d, 0x26, 0x87, 0x15, 0x6f, 0x83, 0xed, 0x52, 0xb1, 0xeb,
	0xc0, 0x5a, 0x56, 0x3d, 0xc2, 0x48, 0xcd, 0xa9, 0x6e, 0xff, 0x05, 0xdc, 0x62, 0xfc, 0xfb, 0xe0,
	0x11, 0x89, 0xd4, 0xfc, 0x55, 0x08, 0x47, 0xb7, 0xc8, 0x82, 0xbb, 0xed, 0xf4, 0x25, 0x53, 0x54,
	0x5a, 0x28, 0xa8, 0xc8, 0x3a, 0xf1, 0x1f, 0xfe, 0x5a, 0x73, 0x1a, 0x71, 0xd7, 0xc5, 0x37, 0x2d,
	0x14, 0xcd, 0xac, 0x13, 0x7d, 0x4e, 0xe6, 0x11, 0xfd, 0x4c, 0xba, 0xc5, 0x17, 0xaa, 0xfc, 0x33,
	0x0c, 0x35, 0xa2, 0x1b, 0xc3, 0x96, 0xf4, 0x25, 0xb6, 0xa7, 0xfd, 0xd6, 0xbc, 0xff, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x78, 0xf1, 0x35, 0xd2, 0x48, 0x0d, 0x00, 0x00,
}
