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
// source: bmd_bag_mbr_info_mlacp_counters_bdl.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_mlacp_member_counters_bundles_bundle_bundle_item

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

type BmdBagMbrInfoMlacpCountersBdl_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) Reset()         { *m = BmdBagMbrInfoMlacpCountersBdl_KEYS{} }
func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersBdl_KEYS) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersBdl_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{0}
}

func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl_KEYS proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersBdl_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

type BmdBagMlacpSyncRqstTlvsRcvd struct {
	AllSyncs             uint32   `protobuf:"varint,1,opt,name=all_syncs,json=allSyncs,proto3" json:"all_syncs,omitempty"`
	ConfigSyncs          uint32   `protobuf:"varint,2,opt,name=config_syncs,json=configSyncs,proto3" json:"config_syncs,omitempty"`
	StateSyncs           uint32   `protobuf:"varint,3,opt,name=state_syncs,json=stateSyncs,proto3" json:"state_syncs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMlacpSyncRqstTlvsRcvd) Reset()         { *m = BmdBagMlacpSyncRqstTlvsRcvd{} }
func (m *BmdBagMlacpSyncRqstTlvsRcvd) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpSyncRqstTlvsRcvd) ProtoMessage()    {}
func (*BmdBagMlacpSyncRqstTlvsRcvd) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{1}
}

func (m *BmdBagMlacpSyncRqstTlvsRcvd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd.Unmarshal(m, b)
}
func (m *BmdBagMlacpSyncRqstTlvsRcvd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpSyncRqstTlvsRcvd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd.Merge(m, src)
}
func (m *BmdBagMlacpSyncRqstTlvsRcvd) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd.Size(m)
}
func (m *BmdBagMlacpSyncRqstTlvsRcvd) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpSyncRqstTlvsRcvd proto.InternalMessageInfo

func (m *BmdBagMlacpSyncRqstTlvsRcvd) GetAllSyncs() uint32 {
	if m != nil {
		return m.AllSyncs
	}
	return 0
}

func (m *BmdBagMlacpSyncRqstTlvsRcvd) GetConfigSyncs() uint32 {
	if m != nil {
		return m.ConfigSyncs
	}
	return 0
}

func (m *BmdBagMlacpSyncRqstTlvsRcvd) GetStateSyncs() uint32 {
	if m != nil {
		return m.StateSyncs
	}
	return 0
}

type BmdBagMlacpSyncRqstCounters struct {
	ReceivedSyncRequests *BmdBagMlacpSyncRqstTlvsRcvd `protobuf:"bytes,1,opt,name=received_sync_requests,json=receivedSyncRequests,proto3" json:"received_sync_requests,omitempty"`
	LastTimeCleared      uint64                       `protobuf:"varint,2,opt,name=last_time_cleared,json=lastTimeCleared,proto3" json:"last_time_cleared,omitempty"`
	TimeSinceCleared     uint64                       `protobuf:"varint,3,opt,name=time_since_cleared,json=timeSinceCleared,proto3" json:"time_since_cleared,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BmdBagMlacpSyncRqstCounters) Reset()         { *m = BmdBagMlacpSyncRqstCounters{} }
func (m *BmdBagMlacpSyncRqstCounters) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpSyncRqstCounters) ProtoMessage()    {}
func (*BmdBagMlacpSyncRqstCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{2}
}

func (m *BmdBagMlacpSyncRqstCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpSyncRqstCounters.Unmarshal(m, b)
}
func (m *BmdBagMlacpSyncRqstCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpSyncRqstCounters.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpSyncRqstCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpSyncRqstCounters.Merge(m, src)
}
func (m *BmdBagMlacpSyncRqstCounters) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpSyncRqstCounters.Size(m)
}
func (m *BmdBagMlacpSyncRqstCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpSyncRqstCounters.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpSyncRqstCounters proto.InternalMessageInfo

func (m *BmdBagMlacpSyncRqstCounters) GetReceivedSyncRequests() *BmdBagMlacpSyncRqstTlvsRcvd {
	if m != nil {
		return m.ReceivedSyncRequests
	}
	return nil
}

func (m *BmdBagMlacpSyncRqstCounters) GetLastTimeCleared() uint64 {
	if m != nil {
		return m.LastTimeCleared
	}
	return 0
}

func (m *BmdBagMlacpSyncRqstCounters) GetTimeSinceCleared() uint64 {
	if m != nil {
		return m.TimeSinceCleared
	}
	return 0
}

type BmdBagMbrInfoMlacpCountersBdlData struct {
	BundleName                       string                       `protobuf:"bytes,1,opt,name=bundle_name,json=bundleName,proto3" json:"bundle_name,omitempty"`
	IccpGroupId                      uint32                       `protobuf:"varint,2,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	MlacpSyncRequestsOnAllLocalPorts *BmdBagMlacpSyncRqstCounters `protobuf:"bytes,3,opt,name=mlacp_sync_requests_on_all_local_ports,json=mlacpSyncRequestsOnAllLocalPorts,proto3" json:"mlacp_sync_requests_on_all_local_ports,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                     `json:"-"`
	XXX_unrecognized                 []byte                       `json:"-"`
	XXX_sizecache                    int32                        `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersBdlData) Reset()         { *m = BmdBagMbrInfoMlacpCountersBdlData{} }
func (m *BmdBagMbrInfoMlacpCountersBdlData) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersBdlData) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersBdlData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{3}
}

func (m *BmdBagMbrInfoMlacpCountersBdlData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersBdlData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersBdlData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersBdlData) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersBdlData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdlData proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersBdlData) GetBundleName() string {
	if m != nil {
		return m.BundleName
	}
	return ""
}

func (m *BmdBagMbrInfoMlacpCountersBdlData) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

func (m *BmdBagMbrInfoMlacpCountersBdlData) GetMlacpSyncRequestsOnAllLocalPorts() *BmdBagMlacpSyncRqstCounters {
	if m != nil {
		return m.MlacpSyncRequestsOnAllLocalPorts
	}
	return nil
}

type BmdBagMbrInfoMlacpCountersNodeData struct {
	NodeId               uint32   `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	IccpGroupId          uint32   `protobuf:"varint,2,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersNodeData) Reset()         { *m = BmdBagMbrInfoMlacpCountersNodeData{} }
func (m *BmdBagMbrInfoMlacpCountersNodeData) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersNodeData) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersNodeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{4}
}

func (m *BmdBagMbrInfoMlacpCountersNodeData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersNodeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersNodeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersNodeData) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersNodeData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersNodeData proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersNodeData) GetNodeId() uint32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *BmdBagMbrInfoMlacpCountersNodeData) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

type BmdBagMlacpTlvCounters struct {
	SentConfigTlVs       uint32                       `protobuf:"varint,1,opt,name=sent_config_tl_vs,json=sentConfigTlVs,proto3" json:"sent_config_tl_vs,omitempty"`
	SentStateTlVs        uint32                       `protobuf:"varint,2,opt,name=sent_state_tl_vs,json=sentStateTlVs,proto3" json:"sent_state_tl_vs,omitempty"`
	SentPriorityTlVs     uint32                       `protobuf:"varint,3,opt,name=sent_priority_tl_vs,json=sentPriorityTlVs,proto3" json:"sent_priority_tl_vs,omitempty"`
	ReceivedPriorityTlVs uint32                       `protobuf:"varint,4,opt,name=received_priority_tl_vs,json=receivedPriorityTlVs,proto3" json:"received_priority_tl_vs,omitempty"`
	ReceivedNakTlVs      uint32                       `protobuf:"varint,5,opt,name=received_nak_tl_vs,json=receivedNakTlVs,proto3" json:"received_nak_tl_vs,omitempty"`
	ReceivedSyncRequests *BmdBagMlacpSyncRqstTlvsRcvd `protobuf:"bytes,6,opt,name=received_sync_requests,json=receivedSyncRequests,proto3" json:"received_sync_requests,omitempty"`
	LastTimeCleared      uint64                       `protobuf:"varint,7,opt,name=last_time_cleared,json=lastTimeCleared,proto3" json:"last_time_cleared,omitempty"`
	TimeSinceCleared     uint64                       `protobuf:"varint,8,opt,name=time_since_cleared,json=timeSinceCleared,proto3" json:"time_since_cleared,omitempty"`
	LastUnexpectedEvent  uint64                       `protobuf:"varint,9,opt,name=last_unexpected_event,json=lastUnexpectedEvent,proto3" json:"last_unexpected_event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BmdBagMlacpTlvCounters) Reset()         { *m = BmdBagMlacpTlvCounters{} }
func (m *BmdBagMlacpTlvCounters) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpTlvCounters) ProtoMessage()    {}
func (*BmdBagMlacpTlvCounters) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{5}
}

func (m *BmdBagMlacpTlvCounters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpTlvCounters.Unmarshal(m, b)
}
func (m *BmdBagMlacpTlvCounters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpTlvCounters.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpTlvCounters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpTlvCounters.Merge(m, src)
}
func (m *BmdBagMlacpTlvCounters) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpTlvCounters.Size(m)
}
func (m *BmdBagMlacpTlvCounters) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpTlvCounters.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpTlvCounters proto.InternalMessageInfo

func (m *BmdBagMlacpTlvCounters) GetSentConfigTlVs() uint32 {
	if m != nil {
		return m.SentConfigTlVs
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetSentStateTlVs() uint32 {
	if m != nil {
		return m.SentStateTlVs
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetSentPriorityTlVs() uint32 {
	if m != nil {
		return m.SentPriorityTlVs
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetReceivedPriorityTlVs() uint32 {
	if m != nil {
		return m.ReceivedPriorityTlVs
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetReceivedNakTlVs() uint32 {
	if m != nil {
		return m.ReceivedNakTlVs
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetReceivedSyncRequests() *BmdBagMlacpSyncRqstTlvsRcvd {
	if m != nil {
		return m.ReceivedSyncRequests
	}
	return nil
}

func (m *BmdBagMlacpTlvCounters) GetLastTimeCleared() uint64 {
	if m != nil {
		return m.LastTimeCleared
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetTimeSinceCleared() uint64 {
	if m != nil {
		return m.TimeSinceCleared
	}
	return 0
}

func (m *BmdBagMlacpTlvCounters) GetLastUnexpectedEvent() uint64 {
	if m != nil {
		return m.LastUnexpectedEvent
	}
	return 0
}

type BmdBagMbrInfoMlacpCountersMbr struct {
	PortName             string                  `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	MlacpTlvCounters     *BmdBagMlacpTlvCounters `protobuf:"bytes,2,opt,name=mlacp_tlv_counters,json=mlacpTlvCounters,proto3" json:"mlacp_tlv_counters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersMbr) Reset()         { *m = BmdBagMbrInfoMlacpCountersMbr{} }
func (m *BmdBagMbrInfoMlacpCountersMbr) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersMbr) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersMbr) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{6}
}

func (m *BmdBagMbrInfoMlacpCountersMbr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersMbr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersMbr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersMbr) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersMbr) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersMbr proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersMbr) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *BmdBagMbrInfoMlacpCountersMbr) GetMlacpTlvCounters() *BmdBagMlacpTlvCounters {
	if m != nil {
		return m.MlacpTlvCounters
	}
	return nil
}

type BmdBagMbrInfoMlacpCountersNode struct {
	NodeData                           *BmdBagMbrInfoMlacpCountersNodeData `protobuf:"bytes,1,opt,name=node_data,json=nodeData,proto3" json:"node_data,omitempty"`
	MlacpSyncRequestsOnAllForeignPorts *BmdBagMlacpSyncRqstCounters        `protobuf:"bytes,2,opt,name=mlacp_sync_requests_on_all_foreign_ports,json=mlacpSyncRequestsOnAllForeignPorts,proto3" json:"mlacp_sync_requests_on_all_foreign_ports,omitempty"`
	MemberData                         []*BmdBagMbrInfoMlacpCountersMbr    `protobuf:"bytes,3,rep,name=member_data,json=memberData,proto3" json:"member_data,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}                            `json:"-"`
	XXX_unrecognized                   []byte                              `json:"-"`
	XXX_sizecache                      int32                               `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersNode) Reset()         { *m = BmdBagMbrInfoMlacpCountersNode{} }
func (m *BmdBagMbrInfoMlacpCountersNode) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersNode) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{7}
}

func (m *BmdBagMbrInfoMlacpCountersNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersNode) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersNode proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersNode) GetNodeData() *BmdBagMbrInfoMlacpCountersNodeData {
	if m != nil {
		return m.NodeData
	}
	return nil
}

func (m *BmdBagMbrInfoMlacpCountersNode) GetMlacpSyncRequestsOnAllForeignPorts() *BmdBagMlacpSyncRqstCounters {
	if m != nil {
		return m.MlacpSyncRequestsOnAllForeignPorts
	}
	return nil
}

func (m *BmdBagMbrInfoMlacpCountersNode) GetMemberData() []*BmdBagMbrInfoMlacpCountersMbr {
	if m != nil {
		return m.MemberData
	}
	return nil
}

type BmdBagMbrInfoMlacpCountersBdl struct {
	BundleData           *BmdBagMbrInfoMlacpCountersBdlData `protobuf:"bytes,50,opt,name=bundle_data,json=bundleData,proto3" json:"bundle_data,omitempty"`
	NodeData             []*BmdBagMbrInfoMlacpCountersNode  `protobuf:"bytes,51,rep,name=node_data,json=nodeData,proto3" json:"node_data,omitempty"`
	MemberData           []*BmdBagMbrInfoMlacpCountersMbr   `protobuf:"bytes,52,rep,name=member_data,json=memberData,proto3" json:"member_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersBdl) Reset()         { *m = BmdBagMbrInfoMlacpCountersBdl{} }
func (m *BmdBagMbrInfoMlacpCountersBdl) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersBdl) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersBdl) Descriptor() ([]byte, []int) {
	return fileDescriptor_63f9ff0118379c3a, []int{8}
}

func (m *BmdBagMbrInfoMlacpCountersBdl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersBdl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersBdl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersBdl) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersBdl) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersBdl proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersBdl) GetBundleData() *BmdBagMbrInfoMlacpCountersBdlData {
	if m != nil {
		return m.BundleData
	}
	return nil
}

func (m *BmdBagMbrInfoMlacpCountersBdl) GetNodeData() []*BmdBagMbrInfoMlacpCountersNode {
	if m != nil {
		return m.NodeData
	}
	return nil
}

func (m *BmdBagMbrInfoMlacpCountersBdl) GetMemberData() []*BmdBagMbrInfoMlacpCountersMbr {
	if m != nil {
		return m.MemberData
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersBdl_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_bdl_KEYS")
	proto.RegisterType((*BmdBagMlacpSyncRqstTlvsRcvd)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mlacp_sync_rqst_tlvs_rcvd")
	proto.RegisterType((*BmdBagMlacpSyncRqstCounters)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mlacp_sync_rqst_counters")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersBdlData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_bdl_data")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersNodeData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_node_data")
	proto.RegisterType((*BmdBagMlacpTlvCounters)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mlacp_tlv_counters")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersMbr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_mbr")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersNode)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_node")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.bundles.bundle.bundle_item.bmd_bag_mbr_info_mlacp_counters_bdl")
}

func init() {
	proto.RegisterFile("bmd_bag_mbr_info_mlacp_counters_bdl.proto", fileDescriptor_63f9ff0118379c3a)
}

var fileDescriptor_63f9ff0118379c3a = []byte{
	// 807 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x6e, 0x1b, 0x37,
	0x10, 0xc6, 0x4a, 0xaa, 0x2d, 0x51, 0x75, 0x2d, 0xd3, 0x6d, 0x6d, 0xb4, 0x87, 0xca, 0xdb, 0xa2,
	0x95, 0xfb, 0xa3, 0x83, 0xdc, 0x3e, 0x40, 0xe1, 0xba, 0x85, 0xd1, 0xc2, 0x35, 0x56, 0x76, 0x81,
	0x9e, 0x08, 0xee, 0x2e, 0xa5, 0x2e, 0xcc, 0x25, 0xd7, 0x5c, 0x4a, 0xb0, 0xd1, 0x53, 0x0b, 0xf7,
	0xd0, 0x3c, 0x46, 0x02, 0x04, 0x09, 0xf2, 0x02, 0xc9, 0x0b, 0xe4, 0x96, 0x5b, 0x02, 0xe4, 0x29,
	0xf2, 0x0c, 0x01, 0x87, 0xdc, 0xb5, 0xec, 0x38, 0x96, 0x72, 0x88, 0xed, 0x9c, 0x6c, 0xce, 0xf7,
	0xcd, 0xec, 0x0c, 0x67, 0xbe, 0xa1, 0xd0, 0x7a, 0x98, 0xc6, 0x24, 0xa4, 0x43, 0x92, 0x86, 0x8a,
	0x24, 0x62, 0x20, 0x49, 0xca, 0x69, 0x94, 0x91, 0x48, 0x8e, 0x84, 0x66, 0x2a, 0x27, 0x61, 0xcc,
	0xbb, 0x99, 0x92, 0x5a, 0x62, 0x12, 0x25, 0x79, 0x24, 0x49, 0x22, 0x73, 0x72, 0xa4, 0x48, 0x38,
	0x12, 0x31, 0x67, 0xe9, 0x50, 0x11, 0x99, 0x31, 0xd5, 0xb5, 0x47, 0x88, 0xa0, 0x52, 0xaa, 0x13,
	0x29, 0xba, 0x36, 0x50, 0xca, 0xd2, 0x90, 0xa9, 0x32, 0x9e, 0x23, 0x16, 0x7f, 0x4b, 0x3f, 0xcd,
	0x52, 0x7f, 0x1f, 0x75, 0x66, 0xc8, 0x86, 0xfc, 0xba, 0xf5, 0x67, 0x1f, 0xaf, 0xa3, 0x56, 0xf9,
	0x49, 0xcd, 0xd4, 0x80, 0x46, 0x6c, 0xd5, 0x6b, 0x7b, 0x9d, 0x46, 0xb0, 0x68, 0xed, 0xdb, 0x85,
	0xd9, 0x3f, 0xf1, 0xd0, 0x5a, 0x19, 0x17, 0xc2, 0xe5, 0xc7, 0x22, 0x22, 0xea, 0x30, 0xd7, 0x44,
	0xf3, 0x71, 0x4e, 0x54, 0x34, 0x8e, 0xf1, 0xa7, 0xa8, 0x41, 0x39, 0x07, 0x28, 0x87, 0x48, 0x0b,
	0x41, 0x9d, 0x72, 0xde, 0x37, 0x67, 0xbc, 0x86, 0xde, 0x8f, 0xa4, 0x18, 0x24, 0x43, 0x87, 0x57,
	0x00, 0x6f, 0x5a, 0x9b, 0xa5, 0x7c, 0x86, 0x9a, 0xb9, 0xa6, 0x9a, 0x39, 0x46, 0x15, 0x18, 0x08,
	0x4c, 0x40, 0xf0, 0x1f, 0x56, 0x50, 0xfb, 0x75, 0x69, 0x14, 0xf5, 0xe1, 0x47, 0x1e, 0xfa, 0x58,
	0xb1, 0x88, 0x25, 0x63, 0x16, 0x3b, 0x9c, 0x1d, 0x8e, 0x58, 0xae, 0x6d, 0x4e, 0xcd, 0xde, 0xbf,
	0x5e, 0xf7, 0x2d, 0xb7, 0xa1, 0x3b, 0xf5, 0xae, 0x82, 0x0f, 0x8b, 0x14, 0x4d, 0x65, 0x81, 0x4b,
	0x10, 0x7f, 0x8d, 0x96, 0x38, 0x35, 0xbc, 0x24, 0x65, 0x24, 0xe2, 0x8c, 0x2a, 0x16, 0xc3, 0x4d,
	0xd5, 0x82, 0x45, 0x03, 0xec, 0x25, 0x29, 0xdb, 0xb4, 0x66, 0xfc, 0x2d, 0xc2, 0x40, 0xcb, 0x13,
	0x11, 0x9d, 0x92, 0xab, 0x40, 0x6e, 0x19, 0xa4, 0x6f, 0x00, 0xc7, 0xf6, 0x1f, 0x57, 0x66, 0x9b,
	0x8c, 0x98, 0x6a, 0x6a, 0x1a, 0xe1, 0xaa, 0x11, 0x34, 0x2d, 0x86, 0x02, 0x59, 0xd3, 0x0e, 0x4d,
	0x19, 0xf6, 0xd1, 0x42, 0x12, 0x45, 0x19, 0x19, 0x2a, 0x39, 0xca, 0x48, 0x12, 0x17, 0xdd, 0x34,
	0xc6, 0x5f, 0x8c, 0x6d, 0x3b, 0xc6, 0x4f, 0x3d, 0xf4, 0xe5, 0x64, 0xfd, 0xae, 0x46, 0x22, 0x05,
	0x31, 0x53, 0xc2, 0x65, 0x44, 0x39, 0xc9, 0xa4, 0xd2, 0xb6, 0xd3, 0xcd, 0xde, 0x3f, 0xd7, 0xd7,
	0x97, 0xc2, 0x37, 0x68, 0x03, 0x32, 0xd9, 0x93, 0xdf, 0xc5, 0x8f, 0x9c, 0xff, 0x66, 0x92, 0xdd,
	0x35, 0xb9, 0xfa, 0x7f, 0x4d, 0xd7, 0xbb, 0x90, 0x31, 0xb3, 0x17, 0xb9, 0x82, 0xe6, 0xe1, 0x90,
	0xc4, 0x4e, 0x0f, 0x73, 0xe6, 0xb8, 0x1d, 0xcf, 0x72, 0x81, 0xfe, 0x93, 0x1a, 0xfa, 0xe4, 0x6c,
	0xc2, 0x9a, 0x8f, 0x4f, 0xe7, 0x7c, 0x1d, 0x2d, 0xe5, 0x4c, 0x98, 0xdc, 0x41, 0x55, 0x9a, 0x93,
	0x71, 0xa1, 0xba, 0x0f, 0x0c, 0xb0, 0x09, 0xf6, 0x3d, 0xfe, 0x47, 0x8e, 0xbf, 0x42, 0x2d, 0xa0,
	0x5a, 0x75, 0x59, 0xa6, 0xfd, 0xe0, 0x82, 0xb1, 0xf7, 0x8d, 0x19, 0x88, 0xdf, 0xa1, 0x65, 0x20,
	0x66, 0x2a, 0x91, 0x2a, 0xd1, 0xc7, 0x8e, 0x6b, 0x95, 0x08, 0x31, 0x76, 0x1d, 0x02, 0xf4, 0x1f,
	0xd0, 0x4a, 0xa9, 0xb4, 0x73, 0x2e, 0x35, 0x70, 0x29, 0xa7, 0xfc, 0x8c, 0xdb, 0x37, 0x08, 0x97,
	0x6e, 0x82, 0x1e, 0x38, 0x8f, 0xf7, 0xc0, 0x63, 0xb1, 0x40, 0x76, 0xe8, 0x01, 0x90, 0x2f, 0x91,
	0xf3, 0xdc, 0x3b, 0x29, 0xe7, 0xf9, 0x37, 0x91, 0x73, 0xfd, 0x62, 0x39, 0xe3, 0x1e, 0xfa, 0x08,
	0x22, 0x8f, 0x04, 0x3b, 0xca, 0x58, 0xa4, 0x59, 0x4c, 0xd8, 0x98, 0x09, 0xbd, 0xda, 0x00, 0x87,
	0x65, 0x03, 0xee, 0x97, 0xd8, 0x96, 0x81, 0xfc, 0x17, 0x1e, 0xfa, 0x7c, 0xda, 0xe8, 0xa6, 0xa1,
	0x32, 0x6b, 0xdc, 0xc8, 0x72, 0x52, 0xfb, 0x75, 0x63, 0x00, 0xe5, 0xdf, 0xf7, 0x10, 0x7e, 0x75,
	0x18, 0x61, 0x9a, 0x9a, 0xbd, 0xbf, 0xaf, 0xb8, 0x13, 0x93, 0x29, 0x04, 0x2d, 0xb0, 0xed, 0xf1,
	0xf1, 0xa6, 0xb3, 0xf8, 0xcf, 0x6a, 0xe8, 0x8b, 0x59, 0xb4, 0x8a, 0xef, 0x79, 0xa8, 0x51, 0x8a,
	0xd6, 0xbd, 0x12, 0xb7, 0xae, 0x70, 0xac, 0xa6, 0xed, 0x91, 0xa0, 0x6e, 0xfe, 0xfd, 0xc9, 0x6c,
	0x94, 0xe7, 0x1e, 0xea, 0x5c, 0xb2, 0x55, 0x07, 0x52, 0xb1, 0x64, 0x28, 0xdc, 0x5e, 0xad, 0xdc,
	0x98, 0xbd, 0xea, 0x5f, 0xbc, 0x57, 0x7f, 0xb6, 0xe9, 0xc2, 0x66, 0xc5, 0x77, 0x3d, 0xd4, 0x74,
	0xdf, 0x84, 0x3e, 0x54, 0xdb, 0xd5, 0x4e, 0xb3, 0x77, 0x72, 0xfd, 0x7d, 0x48, 0x43, 0x15, 0x20,
	0x1b, 0xcf, 0xf4, 0xc0, 0xbf, 0x5d, 0x9b, 0x2e, 0xa4, 0x30, 0xe6, 0xf8, 0x81, 0x57, 0xbe, 0xa3,
	0x50, 0x50, 0x0f, 0xda, 0xf1, 0xff, 0xf5, 0x17, 0x54, 0x3c, 0xf4, 0xc5, 0x9b, 0x0e, 0x93, 0x75,
	0xe7, 0x8c, 0x08, 0x36, 0xe0, 0xf2, 0xff, 0xbb, 0x19, 0x22, 0x98, 0x98, 0xff, 0xf3, 0x43, 0xf2,
	0xfd, 0x0d, 0x1d, 0x92, 0x70, 0x0e, 0x7e, 0xf1, 0x6f, 0xbc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x73, 0x99, 0x71, 0x1e, 0x0c, 0x00, 0x00,
}
