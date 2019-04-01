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
// source: bmd_bag_mbr_info_mlacp_counters_rg.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_mlacp_member_counters_members_member_member_item

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

type BmdBagMbrInfoMlacpCountersRg_KEYS struct {
	MemberInterface      string   `protobuf:"bytes,1,opt,name=member_interface,json=memberInterface,proto3" json:"member_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) Reset()         { *m = BmdBagMbrInfoMlacpCountersRg_KEYS{} }
func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersRg_KEYS) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersRg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc94c7200227a96, []int{0}
}

func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg_KEYS proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersRg_KEYS) GetMemberInterface() string {
	if m != nil {
		return m.MemberInterface
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
	return fileDescriptor_ccc94c7200227a96, []int{1}
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
	return fileDescriptor_ccc94c7200227a96, []int{2}
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
	return fileDescriptor_ccc94c7200227a96, []int{3}
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
	return fileDescriptor_ccc94c7200227a96, []int{4}
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
	return fileDescriptor_ccc94c7200227a96, []int{5}
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
	return fileDescriptor_ccc94c7200227a96, []int{6}
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
	return fileDescriptor_ccc94c7200227a96, []int{7}
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
	BundleData           *BmdBagMbrInfoMlacpCountersBdlData `protobuf:"bytes,1,opt,name=bundle_data,json=bundleData,proto3" json:"bundle_data,omitempty"`
	NodeData             []*BmdBagMbrInfoMlacpCountersNode  `protobuf:"bytes,2,rep,name=node_data,json=nodeData,proto3" json:"node_data,omitempty"`
	MemberData           []*BmdBagMbrInfoMlacpCountersMbr   `protobuf:"bytes,3,rep,name=member_data,json=memberData,proto3" json:"member_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersBdl) Reset()         { *m = BmdBagMbrInfoMlacpCountersBdl{} }
func (m *BmdBagMbrInfoMlacpCountersBdl) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersBdl) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersBdl) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc94c7200227a96, []int{8}
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

type BmdBagMbrInfoMlacpCountersRg struct {
	Items                []*BmdBagMbrInfoMlacpCountersBdl `protobuf:"bytes,50,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *BmdBagMbrInfoMlacpCountersRg) Reset()         { *m = BmdBagMbrInfoMlacpCountersRg{} }
func (m *BmdBagMbrInfoMlacpCountersRg) String() string { return proto.CompactTextString(m) }
func (*BmdBagMbrInfoMlacpCountersRg) ProtoMessage()    {}
func (*BmdBagMbrInfoMlacpCountersRg) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc94c7200227a96, []int{9}
}

func (m *BmdBagMbrInfoMlacpCountersRg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg.Unmarshal(m, b)
}
func (m *BmdBagMbrInfoMlacpCountersRg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg.Marshal(b, m, deterministic)
}
func (m *BmdBagMbrInfoMlacpCountersRg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg.Merge(m, src)
}
func (m *BmdBagMbrInfoMlacpCountersRg) XXX_Size() int {
	return xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg.Size(m)
}
func (m *BmdBagMbrInfoMlacpCountersRg) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMbrInfoMlacpCountersRg proto.InternalMessageInfo

func (m *BmdBagMbrInfoMlacpCountersRg) GetItems() []*BmdBagMbrInfoMlacpCountersBdl {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersRg_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_rg_KEYS")
	proto.RegisterType((*BmdBagMlacpSyncRqstTlvsRcvd)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mlacp_sync_rqst_tlvs_rcvd")
	proto.RegisterType((*BmdBagMlacpSyncRqstCounters)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mlacp_sync_rqst_counters")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersBdlData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_bdl_data")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersNodeData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_node_data")
	proto.RegisterType((*BmdBagMlacpTlvCounters)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mlacp_tlv_counters")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersMbr)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_mbr")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersNode)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_node")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_bdl")
	proto.RegisterType((*BmdBagMbrInfoMlacpCountersRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_member_counters.members.member.member_item.bmd_bag_mbr_info_mlacp_counters_rg")
}

func init() {
	proto.RegisterFile("bmd_bag_mbr_info_mlacp_counters_rg.proto", fileDescriptor_ccc94c7200227a96)
}

var fileDescriptor_ccc94c7200227a96 = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x57, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0xd6, 0xd8, 0xce, 0xc6, 0x2e, 0xb3, 0xac, 0xd3, 0x01, 0xb2, 0x82, 0x03, 0xce, 0x80, 0x88,
	0x97, 0x1f, 0x1f, 0x8c, 0x78, 0x00, 0xb4, 0x04, 0xb4, 0x02, 0x85, 0x68, 0x6c, 0x90, 0x38, 0xb5,
	0x7a, 0x66, 0xda, 0xa6, 0x95, 0x9e, 0xee, 0x49, 0x4f, 0xdb, 0x4a, 0xc4, 0x09, 0x14, 0x0e, 0xf0,
	0x0a, 0xdc, 0x40, 0x42, 0x20, 0x5e, 0x00, 0xae, 0x1c, 0xb8, 0x71, 0x03, 0x89, 0xa7, 0xe0, 0x19,
	0x50, 0x57, 0xf7, 0x4c, 0xbc, 0x10, 0x32, 0xcb, 0x81, 0xfd, 0x39, 0xcd, 0x76, 0x7d, 0x5f, 0xd5,
	0x56, 0x57, 0xd5, 0x57, 0xbd, 0x0b, 0x93, 0xb4, 0xc8, 0x69, 0xca, 0x56, 0xb4, 0x48, 0x0d, 0x15,
	0x6a, 0xa9, 0x69, 0x21, 0x59, 0x56, 0xd2, 0x4c, 0xaf, 0x95, 0xe5, 0xa6, 0xa2, 0x66, 0x35, 0x2d,
	0x8d, 0xb6, 0x9a, 0xd0, 0x4c, 0x54, 0x99, 0xa6, 0x42, 0x57, 0xf4, 0x9e, 0xa1, 0xe9, 0x5a, 0xe5,
	0x92, 0x17, 0x2b, 0x43, 0x75, 0xc9, 0xcd, 0xd4, 0x1f, 0x31, 0x80, 0x29, 0x98, 0x15, 0x5a, 0x4d,
	0x7d, 0x9c, 0x82, 0x17, 0x29, 0x37, 0x4d, 0xb8, 0xa9, 0x3f, 0xd7, 0xdf, 0xf0, 0xa1, 0xc2, 0xf2,
	0x22, 0x5e, 0xc0, 0x8d, 0xf6, 0x64, 0xe8, 0xbb, 0x37, 0x3f, 0x9a, 0x93, 0x03, 0x18, 0xd5, 0x9e,
	0x0e, 0x58, 0xb2, 0x8c, 0xef, 0x47, 0xe3, 0x68, 0x32, 0x48, 0xf6, 0xbc, 0xfd, 0xa8, 0x36, 0xc7,
	0x0f, 0x22, 0xb8, 0xde, 0x84, 0xc5, 0x68, 0xd5, 0x7d, 0x95, 0x51, 0x73, 0xb7, 0xb2, 0xd4, 0xca,
	0x4d, 0x45, 0x4d, 0xb6, 0xc9, 0xc9, 0x73, 0x30, 0x60, 0x52, 0x22, 0x54, 0x61, 0xa4, 0xdd, 0xa4,
	0xcf, 0xa4, 0x9c, 0xbb, 0x33, 0xb9, 0x0e, 0x4f, 0x64, 0x5a, 0x2d, 0xc5, 0x2a, 0xe0, 0x1d, 0xc4,
	0x87, 0xde, 0xe6, 0x29, 0xcf, 0xc3, 0xb0, 0xb2, 0xcc, 0xf2, 0xc0, 0xe8, 0x22, 0x03, 0xd0, 0x84,
	0x84, 0xf8, 0xc7, 0x0e, 0x8c, 0xff, 0x2d, 0x8d, 0xfa, 0x7a, 0xe4, 0xa7, 0x08, 0x9e, 0x31, 0x3c,
	0xe3, 0x62, 0xc3, 0xf3, 0x80, 0xf3, 0xbb, 0x6b, 0x5e, 0x59, 0x9f, 0xd3, 0x70, 0xf6, 0x59, 0x34,
	0xfd, 0x9f, 0xbb, 0x30, 0x6d, 0xad, 0x55, 0xf2, 0x54, 0x9d, 0xa2, 0xbb, 0x59, 0x12, 0x12, 0x24,
	0x2f, 0xc3, 0x15, 0xc9, 0x1c, 0x4f, 0x14, 0x9c, 0x66, 0x92, 0x33, 0xc3, 0x73, 0xac, 0x54, 0x2f,
	0xd9, 0x73, 0xc0, 0x42, 0x14, 0xfc, 0xd0, 0x9b, 0xc9, 0xab, 0x40, 0x90, 0x56, 0x09, 0x95, 0x3d,
	0x24, 0x77, 0x91, 0x3c, 0x72, 0xc8, 0xdc, 0x01, 0x81, 0x1d, 0xff, 0xd2, 0x69, 0x9f, 0xd2, 0x34,
	0x97, 0x34, 0x67, 0x96, 0xb9, 0x46, 0x84, 0x2a, 0x28, 0x56, 0xd4, 0x43, 0x01, 0xde, 0x74, 0x8b,
	0x15, 0x9c, 0xc4, 0xb0, 0x2b, 0xb2, 0xac, 0xa4, 0x2b, 0xa3, 0xd7, 0x25, 0x15, 0x79, 0xdd, 0x4d,
	0x67, 0x7c, 0xc7, 0xd9, 0x8e, 0x72, 0xf2, 0x5b, 0x04, 0x2f, 0x6d, 0xdf, 0x3f, 0xdc, 0x91, 0x6a,
	0x45, 0xdd, 0x94, 0x48, 0x9d, 0x31, 0x49, 0x4b, 0x6d, 0xac, 0xef, 0xf4, 0x70, 0xf6, 0xe9, 0xd9,
	0xf5, 0xa5, 0xf6, 0x4d, 0xc6, 0x88, 0x6c, 0xf7, 0xe4, 0x7d, 0xf5, 0xa6, 0x94, 0xef, 0xb9, 0x64,
	0x6f, 0xbb, 0x5c, 0xe3, 0x8f, 0xe1, 0xa0, 0xad, 0x8e, 0x4a, 0xe7, 0xdc, 0x17, 0xf2, 0x1a, 0x5c,
	0xc6, 0x83, 0xc8, 0x83, 0x1e, 0x76, 0xdc, 0xf1, 0x28, 0x3f, 0x49, 0x01, 0xe3, 0x5f, 0x7b, 0xf0,
	0xec, 0xf1, 0x84, 0xad, 0xdc, 0x3c, 0x9c, 0xf3, 0x03, 0xb8, 0x52, 0x71, 0xe5, 0x72, 0x47, 0x55,
	0x59, 0x49, 0x37, 0xb5, 0xea, 0x9e, 0x74, 0xc0, 0x21, 0xda, 0x17, 0xf2, 0xc3, 0x8a, 0xdc, 0x80,
	0x11, 0x52, 0xbd, 0xba, 0x3c, 0xd3, 0xff, 0xc2, 0x5d, 0x67, 0x9f, 0x3b, 0x33, 0x12, 0x5f, 0x83,
	0xab, 0x48, 0x2c, 0x8d, 0xd0, 0x46, 0xd8, 0xfb, 0x81, 0xeb, 0x95, 0x88, 0x31, 0x6e, 0x07, 0x04,
	0xe9, 0x6f, 0xc0, 0xb5, 0x46, 0x69, 0x7f, 0x73, 0xe9, 0xa1, 0x4b, 0x33, 0xe5, 0xc7, 0xdc, 0x5e,
	0x01, 0xd2, 0xb8, 0x29, 0x76, 0x27, 0x78, 0x5c, 0x42, 0x8f, 0xbd, 0x1a, 0xb9, 0xc5, 0xee, 0x20,
	0xf9, 0x31, 0x72, 0xde, 0xb9, 0x90, 0x72, 0xbe, 0xfc, 0x5f, 0xe4, 0xdc, 0x7f, 0xb4, 0x9c, 0xc9,
	0x0c, 0x9e, 0xc6, 0xc8, 0x6b, 0xc5, 0xef, 0x95, 0x3c, 0xb3, 0x3c, 0xa7, 0x7c, 0xc3, 0x95, 0xdd,
	0x1f, 0xa0, 0xc3, 0x55, 0x07, 0x7e, 0xd0, 0x60, 0x37, 0x1d, 0x14, 0xff, 0x19, 0xc1, 0x0b, 0x6d,
	0xa3, 0x5b, 0xa4, 0xc6, 0xad, 0x71, 0x27, 0xcb, 0x6d, 0xed, 0xf7, 0x9d, 0x01, 0x95, 0xff, 0x7d,
	0x04, 0xe4, 0x9f, 0xc3, 0x88, 0xd3, 0x34, 0x9c, 0x7d, 0x72, 0xca, 0x9d, 0xd8, 0x4e, 0x21, 0x19,
	0xa1, 0x6d, 0x21, 0x37, 0x87, 0xc1, 0x12, 0xff, 0xde, 0x83, 0x17, 0x4f, 0xa2, 0x55, 0xf2, 0x5d,
	0x04, 0x83, 0x46, 0xb4, 0xe1, 0x95, 0xf8, 0xf2, 0x14, 0xc7, 0xaa, 0x6d, 0x8f, 0x24, 0x7d, 0xf7,
	0xe3, 0x5b, 0x6e, 0xa3, 0xfc, 0x11, 0xc1, 0xe4, 0x31, 0x5b, 0x75, 0xa9, 0x0d, 0x17, 0x2b, 0x15,
	0xf6, 0x6a, 0xe7, 0xdc, 0xec, 0xd5, 0xf8, 0xd1, 0x7b, 0xf5, 0x6d, 0x9f, 0x2e, 0x6e, 0x56, 0xf2,
	0x6d, 0x04, 0xc3, 0x10, 0x14, 0xfb, 0xd0, 0x1d, 0x77, 0x27, 0xc3, 0xd9, 0x83, 0xb3, 0xef, 0x43,
	0x91, 0x9a, 0x04, 0xbc, 0x83, 0xeb, 0x41, 0xfc, 0x75, 0xaf, 0x5d, 0x48, 0x69, 0x2e, 0xc9, 0x0f,
	0x51, 0xf3, 0x8e, 0x6e, 0x0d, 0xd6, 0x17, 0x67, 0x7f, 0xa1, 0xfa, 0xa1, 0xaf, 0xdf, 0x74, 0x9c,
	0xac, 0x6f, 0x8e, 0x89, 0xa0, 0x83, 0xc5, 0xff, 0xfc, 0x7c, 0x88, 0x60, 0x6b, 0xfe, 0x2f, 0xcc,
	0x90, 0xfc, 0x1c, 0x41, 0xdc, 0xfe, 0x97, 0x38, 0xf9, 0x2a, 0x82, 0x4b, 0x2e, 0x7e, 0xb5, 0x3f,
	0x3b, 0x2f, 0x37, 0x49, 0x73, 0x99, 0xf8, 0x9c, 0xd2, 0x1d, 0xfc, 0xaf, 0xe5, 0xf5, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x49, 0x3c, 0x5f, 0xd9, 0xe1, 0x0c, 0x00, 0x00,
}