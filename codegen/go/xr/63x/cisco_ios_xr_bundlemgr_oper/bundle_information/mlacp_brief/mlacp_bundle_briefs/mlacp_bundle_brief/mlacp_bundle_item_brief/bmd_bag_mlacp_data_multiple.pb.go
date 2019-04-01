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
// source: bmd_bag_mlacp_data_multiple.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_mlacp_brief_mlacp_bundle_briefs_mlacp_bundle_brief_mlacp_bundle_item_brief

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

type BmdBagMlacpDataMultiple_KEYS struct {
	BundleInterface      string   `protobuf:"bytes,1,opt,name=bundle_interface,json=bundleInterface,proto3" json:"bundle_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMlacpDataMultiple_KEYS) Reset()         { *m = BmdBagMlacpDataMultiple_KEYS{} }
func (m *BmdBagMlacpDataMultiple_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpDataMultiple_KEYS) ProtoMessage()    {}
func (*BmdBagMlacpDataMultiple_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{0}
}

func (m *BmdBagMlacpDataMultiple_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS.Unmarshal(m, b)
}
func (m *BmdBagMlacpDataMultiple_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpDataMultiple_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS.Merge(m, src)
}
func (m *BmdBagMlacpDataMultiple_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS.Size(m)
}
func (m *BmdBagMlacpDataMultiple_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpDataMultiple_KEYS proto.InternalMessageInfo

func (m *BmdBagMlacpDataMultiple_KEYS) GetBundleInterface() string {
	if m != nil {
		return m.BundleInterface
	}
	return ""
}

type EtherMacaddr_ struct {
	Macaddr              []uint32 `protobuf:"varint,1,rep,packed,name=macaddr,proto3" json:"macaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EtherMacaddr_) Reset()         { *m = EtherMacaddr_{} }
func (m *EtherMacaddr_) String() string { return proto.CompactTextString(m) }
func (*EtherMacaddr_) ProtoMessage()    {}
func (*EtherMacaddr_) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{1}
}

func (m *EtherMacaddr_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EtherMacaddr_.Unmarshal(m, b)
}
func (m *EtherMacaddr_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EtherMacaddr_.Marshal(b, m, deterministic)
}
func (m *EtherMacaddr_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EtherMacaddr_.Merge(m, src)
}
func (m *EtherMacaddr_) XXX_Size() int {
	return xxx_messageInfo_EtherMacaddr_.Size(m)
}
func (m *EtherMacaddr_) XXX_DiscardUnknown() {
	xxx_messageInfo_EtherMacaddr_.DiscardUnknown(m)
}

var xxx_messageInfo_EtherMacaddr_ proto.InternalMessageInfo

func (m *EtherMacaddr_) GetMacaddr() []uint32 {
	if m != nil {
		return m.Macaddr
	}
	return nil
}

type BmSystemIdSt struct {
	SystemPrio           uint32         `protobuf:"varint,1,opt,name=system_prio,json=systemPrio,proto3" json:"system_prio,omitempty"`
	SystemMacAddr        *EtherMacaddr_ `protobuf:"bytes,2,opt,name=system_mac_addr,json=systemMacAddr,proto3" json:"system_mac_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BmSystemIdSt) Reset()         { *m = BmSystemIdSt{} }
func (m *BmSystemIdSt) String() string { return proto.CompactTextString(m) }
func (*BmSystemIdSt) ProtoMessage()    {}
func (*BmSystemIdSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{2}
}

func (m *BmSystemIdSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmSystemIdSt.Unmarshal(m, b)
}
func (m *BmSystemIdSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmSystemIdSt.Marshal(b, m, deterministic)
}
func (m *BmSystemIdSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmSystemIdSt.Merge(m, src)
}
func (m *BmSystemIdSt) XXX_Size() int {
	return xxx_messageInfo_BmSystemIdSt.Size(m)
}
func (m *BmSystemIdSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmSystemIdSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmSystemIdSt proto.InternalMessageInfo

func (m *BmSystemIdSt) GetSystemPrio() uint32 {
	if m != nil {
		return m.SystemPrio
	}
	return 0
}

func (m *BmSystemIdSt) GetSystemMacAddr() *EtherMacaddr_ {
	if m != nil {
		return m.SystemMacAddr
	}
	return nil
}

type MlacpRgNodeInfoSt struct {
	MlacpNodeId          uint32        `protobuf:"varint,1,opt,name=mlacp_node_id,json=mlacpNodeId,proto3" json:"mlacp_node_id,omitempty"`
	LdpId                string        `protobuf:"bytes,2,opt,name=ldp_id,json=ldpId,proto3" json:"ldp_id,omitempty"`
	VersionNumber        uint32        `protobuf:"varint,3,opt,name=version_number,json=versionNumber,proto3" json:"version_number,omitempty"`
	SystemId             *BmSystemIdSt `protobuf:"bytes,4,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	NodeState            string        `protobuf:"bytes,5,opt,name=node_state,json=nodeState,proto3" json:"node_state,omitempty"`
	IccpGroupState       string        `protobuf:"bytes,6,opt,name=iccp_group_state,json=iccpGroupState,proto3" json:"iccp_group_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MlacpRgNodeInfoSt) Reset()         { *m = MlacpRgNodeInfoSt{} }
func (m *MlacpRgNodeInfoSt) String() string { return proto.CompactTextString(m) }
func (*MlacpRgNodeInfoSt) ProtoMessage()    {}
func (*MlacpRgNodeInfoSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{3}
}

func (m *MlacpRgNodeInfoSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MlacpRgNodeInfoSt.Unmarshal(m, b)
}
func (m *MlacpRgNodeInfoSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MlacpRgNodeInfoSt.Marshal(b, m, deterministic)
}
func (m *MlacpRgNodeInfoSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MlacpRgNodeInfoSt.Merge(m, src)
}
func (m *MlacpRgNodeInfoSt) XXX_Size() int {
	return xxx_messageInfo_MlacpRgNodeInfoSt.Size(m)
}
func (m *MlacpRgNodeInfoSt) XXX_DiscardUnknown() {
	xxx_messageInfo_MlacpRgNodeInfoSt.DiscardUnknown(m)
}

var xxx_messageInfo_MlacpRgNodeInfoSt proto.InternalMessageInfo

func (m *MlacpRgNodeInfoSt) GetMlacpNodeId() uint32 {
	if m != nil {
		return m.MlacpNodeId
	}
	return 0
}

func (m *MlacpRgNodeInfoSt) GetLdpId() string {
	if m != nil {
		return m.LdpId
	}
	return ""
}

func (m *MlacpRgNodeInfoSt) GetVersionNumber() uint32 {
	if m != nil {
		return m.VersionNumber
	}
	return 0
}

func (m *MlacpRgNodeInfoSt) GetSystemId() *BmSystemIdSt {
	if m != nil {
		return m.SystemId
	}
	return nil
}

func (m *MlacpRgNodeInfoSt) GetNodeState() string {
	if m != nil {
		return m.NodeState
	}
	return ""
}

func (m *MlacpRgNodeInfoSt) GetIccpGroupState() string {
	if m != nil {
		return m.IccpGroupState
	}
	return ""
}

type BmdBagMlacpRg struct {
	IccpGroupId          uint32               `protobuf:"varint,1,opt,name=iccp_group_id,json=iccpGroupId,proto3" json:"iccp_group_id,omitempty"`
	Singleton            uint32               `protobuf:"varint,2,opt,name=singleton,proto3" json:"singleton,omitempty"`
	ConnectTimerRunning  uint64               `protobuf:"varint,3,opt,name=connect_timer_running,json=connectTimerRunning,proto3" json:"connect_timer_running,omitempty"`
	NodeData             []*MlacpRgNodeInfoSt `protobuf:"bytes,4,rep,name=node_data,json=nodeData,proto3" json:"node_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BmdBagMlacpRg) Reset()         { *m = BmdBagMlacpRg{} }
func (m *BmdBagMlacpRg) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpRg) ProtoMessage()    {}
func (*BmdBagMlacpRg) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{4}
}

func (m *BmdBagMlacpRg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpRg.Unmarshal(m, b)
}
func (m *BmdBagMlacpRg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpRg.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpRg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpRg.Merge(m, src)
}
func (m *BmdBagMlacpRg) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpRg.Size(m)
}
func (m *BmdBagMlacpRg) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpRg.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpRg proto.InternalMessageInfo

func (m *BmdBagMlacpRg) GetIccpGroupId() uint32 {
	if m != nil {
		return m.IccpGroupId
	}
	return 0
}

func (m *BmdBagMlacpRg) GetSingleton() uint32 {
	if m != nil {
		return m.Singleton
	}
	return 0
}

func (m *BmdBagMlacpRg) GetConnectTimerRunning() uint64 {
	if m != nil {
		return m.ConnectTimerRunning
	}
	return 0
}

func (m *BmdBagMlacpRg) GetNodeData() []*MlacpRgNodeInfoSt {
	if m != nil {
		return m.NodeData
	}
	return nil
}

type BmMacAddrSt struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmMacAddrSt) Reset()         { *m = BmMacAddrSt{} }
func (m *BmMacAddrSt) String() string { return proto.CompactTextString(m) }
func (*BmMacAddrSt) ProtoMessage()    {}
func (*BmMacAddrSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{5}
}

func (m *BmMacAddrSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmMacAddrSt.Unmarshal(m, b)
}
func (m *BmMacAddrSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmMacAddrSt.Marshal(b, m, deterministic)
}
func (m *BmMacAddrSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmMacAddrSt.Merge(m, src)
}
func (m *BmMacAddrSt) XXX_Size() int {
	return xxx_messageInfo_BmMacAddrSt.Size(m)
}
func (m *BmMacAddrSt) XXX_DiscardUnknown() {
	xxx_messageInfo_BmMacAddrSt.DiscardUnknown(m)
}

var xxx_messageInfo_BmMacAddrSt proto.InternalMessageInfo

func (m *BmMacAddrSt) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MlacpBdlInfoSt struct {
	BundleName           string       `protobuf:"bytes,1,opt,name=bundle_name,json=bundleName,proto3" json:"bundle_name,omitempty"`
	MlacpNodeId          uint32       `protobuf:"varint,2,opt,name=mlacp_node_id,json=mlacpNodeId,proto3" json:"mlacp_node_id,omitempty"`
	MacAddress           *BmMacAddrSt `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	AggregatorId         uint32       `protobuf:"varint,4,opt,name=aggregator_id,json=aggregatorId,proto3" json:"aggregator_id,omitempty"`
	BundleState          string       `protobuf:"bytes,5,opt,name=bundle_state,json=bundleState,proto3" json:"bundle_state,omitempty"`
	PortPriority         uint32       `protobuf:"varint,6,opt,name=port_priority,json=portPriority,proto3" json:"port_priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MlacpBdlInfoSt) Reset()         { *m = MlacpBdlInfoSt{} }
func (m *MlacpBdlInfoSt) String() string { return proto.CompactTextString(m) }
func (*MlacpBdlInfoSt) ProtoMessage()    {}
func (*MlacpBdlInfoSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{6}
}

func (m *MlacpBdlInfoSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MlacpBdlInfoSt.Unmarshal(m, b)
}
func (m *MlacpBdlInfoSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MlacpBdlInfoSt.Marshal(b, m, deterministic)
}
func (m *MlacpBdlInfoSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MlacpBdlInfoSt.Merge(m, src)
}
func (m *MlacpBdlInfoSt) XXX_Size() int {
	return xxx_messageInfo_MlacpBdlInfoSt.Size(m)
}
func (m *MlacpBdlInfoSt) XXX_DiscardUnknown() {
	xxx_messageInfo_MlacpBdlInfoSt.DiscardUnknown(m)
}

var xxx_messageInfo_MlacpBdlInfoSt proto.InternalMessageInfo

func (m *MlacpBdlInfoSt) GetBundleName() string {
	if m != nil {
		return m.BundleName
	}
	return ""
}

func (m *MlacpBdlInfoSt) GetMlacpNodeId() uint32 {
	if m != nil {
		return m.MlacpNodeId
	}
	return 0
}

func (m *MlacpBdlInfoSt) GetMacAddress() *BmMacAddrSt {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *MlacpBdlInfoSt) GetAggregatorId() uint32 {
	if m != nil {
		return m.AggregatorId
	}
	return 0
}

func (m *MlacpBdlInfoSt) GetBundleState() string {
	if m != nil {
		return m.BundleState
	}
	return ""
}

func (m *MlacpBdlInfoSt) GetPortPriority() uint32 {
	if m != nil {
		return m.PortPriority
	}
	return 0
}

type MlacpMbrInfoSt struct {
	PortName             string   `protobuf:"bytes,1,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	InterfaceHandle      string   `protobuf:"bytes,2,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	MlacpNodeId          uint32   `protobuf:"varint,3,opt,name=mlacp_node_id,json=mlacpNodeId,proto3" json:"mlacp_node_id,omitempty"`
	PortNumber           uint32   `protobuf:"varint,4,opt,name=port_number,json=portNumber,proto3" json:"port_number,omitempty"`
	OperationalPriority  uint32   `protobuf:"varint,5,opt,name=operational_priority,json=operationalPriority,proto3" json:"operational_priority,omitempty"`
	ConfiguredPriority   uint32   `protobuf:"varint,6,opt,name=configured_priority,json=configuredPriority,proto3" json:"configured_priority,omitempty"`
	MemberState          string   `protobuf:"bytes,7,opt,name=member_state,json=memberState,proto3" json:"member_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MlacpMbrInfoSt) Reset()         { *m = MlacpMbrInfoSt{} }
func (m *MlacpMbrInfoSt) String() string { return proto.CompactTextString(m) }
func (*MlacpMbrInfoSt) ProtoMessage()    {}
func (*MlacpMbrInfoSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{7}
}

func (m *MlacpMbrInfoSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MlacpMbrInfoSt.Unmarshal(m, b)
}
func (m *MlacpMbrInfoSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MlacpMbrInfoSt.Marshal(b, m, deterministic)
}
func (m *MlacpMbrInfoSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MlacpMbrInfoSt.Merge(m, src)
}
func (m *MlacpMbrInfoSt) XXX_Size() int {
	return xxx_messageInfo_MlacpMbrInfoSt.Size(m)
}
func (m *MlacpMbrInfoSt) XXX_DiscardUnknown() {
	xxx_messageInfo_MlacpMbrInfoSt.DiscardUnknown(m)
}

var xxx_messageInfo_MlacpMbrInfoSt proto.InternalMessageInfo

func (m *MlacpMbrInfoSt) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *MlacpMbrInfoSt) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

func (m *MlacpMbrInfoSt) GetMlacpNodeId() uint32 {
	if m != nil {
		return m.MlacpNodeId
	}
	return 0
}

func (m *MlacpMbrInfoSt) GetPortNumber() uint32 {
	if m != nil {
		return m.PortNumber
	}
	return 0
}

func (m *MlacpMbrInfoSt) GetOperationalPriority() uint32 {
	if m != nil {
		return m.OperationalPriority
	}
	return 0
}

func (m *MlacpMbrInfoSt) GetConfiguredPriority() uint32 {
	if m != nil {
		return m.ConfiguredPriority
	}
	return 0
}

func (m *MlacpMbrInfoSt) GetMemberState() string {
	if m != nil {
		return m.MemberState
	}
	return ""
}

type BmdBagMlacpBdl struct {
	BundleInterfaceKey   uint32            `protobuf:"varint,1,opt,name=bundle_interface_key,json=bundleInterfaceKey,proto3" json:"bundle_interface_key,omitempty"`
	MediaType            string            `protobuf:"bytes,2,opt,name=media_type,json=mediaType,proto3" json:"media_type,omitempty"`
	RedundancyObjectId   uint64            `protobuf:"varint,3,opt,name=redundancy_object_id,json=redundancyObjectId,proto3" json:"redundancy_object_id,omitempty"`
	MlacpBundleData      []*MlacpBdlInfoSt `protobuf:"bytes,4,rep,name=mlacp_bundle_data,json=mlacpBundleData,proto3" json:"mlacp_bundle_data,omitempty"`
	MlacpMemberData      []*MlacpMbrInfoSt `protobuf:"bytes,5,rep,name=mlacp_member_data,json=mlacpMemberData,proto3" json:"mlacp_member_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BmdBagMlacpBdl) Reset()         { *m = BmdBagMlacpBdl{} }
func (m *BmdBagMlacpBdl) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpBdl) ProtoMessage()    {}
func (*BmdBagMlacpBdl) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{8}
}

func (m *BmdBagMlacpBdl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpBdl.Unmarshal(m, b)
}
func (m *BmdBagMlacpBdl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpBdl.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpBdl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpBdl.Merge(m, src)
}
func (m *BmdBagMlacpBdl) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpBdl.Size(m)
}
func (m *BmdBagMlacpBdl) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpBdl.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpBdl proto.InternalMessageInfo

func (m *BmdBagMlacpBdl) GetBundleInterfaceKey() uint32 {
	if m != nil {
		return m.BundleInterfaceKey
	}
	return 0
}

func (m *BmdBagMlacpBdl) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *BmdBagMlacpBdl) GetRedundancyObjectId() uint64 {
	if m != nil {
		return m.RedundancyObjectId
	}
	return 0
}

func (m *BmdBagMlacpBdl) GetMlacpBundleData() []*MlacpBdlInfoSt {
	if m != nil {
		return m.MlacpBundleData
	}
	return nil
}

func (m *BmdBagMlacpBdl) GetMlacpMemberData() []*MlacpMbrInfoSt {
	if m != nil {
		return m.MlacpMemberData
	}
	return nil
}

type BmdBagMlacpData struct {
	IccpGroupData        *BmdBagMlacpRg    `protobuf:"bytes,1,opt,name=iccp_group_data,json=iccpGroupData,proto3" json:"iccp_group_data,omitempty"`
	BundleData           []*BmdBagMlacpBdl `protobuf:"bytes,2,rep,name=bundle_data,json=bundleData,proto3" json:"bundle_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BmdBagMlacpData) Reset()         { *m = BmdBagMlacpData{} }
func (m *BmdBagMlacpData) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpData) ProtoMessage()    {}
func (*BmdBagMlacpData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{9}
}

func (m *BmdBagMlacpData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpData.Unmarshal(m, b)
}
func (m *BmdBagMlacpData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpData.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpData.Merge(m, src)
}
func (m *BmdBagMlacpData) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpData.Size(m)
}
func (m *BmdBagMlacpData) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpData.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpData proto.InternalMessageInfo

func (m *BmdBagMlacpData) GetIccpGroupData() *BmdBagMlacpRg {
	if m != nil {
		return m.IccpGroupData
	}
	return nil
}

func (m *BmdBagMlacpData) GetBundleData() []*BmdBagMlacpBdl {
	if m != nil {
		return m.BundleData
	}
	return nil
}

type BmdBagMlacpDataMultiple struct {
	MlacpData            []*BmdBagMlacpData `protobuf:"bytes,50,rep,name=mlacp_data,json=mlacpData,proto3" json:"mlacp_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *BmdBagMlacpDataMultiple) Reset()         { *m = BmdBagMlacpDataMultiple{} }
func (m *BmdBagMlacpDataMultiple) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpDataMultiple) ProtoMessage()    {}
func (*BmdBagMlacpDataMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_f226b3f9a25ce63b, []int{10}
}

func (m *BmdBagMlacpDataMultiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpDataMultiple.Unmarshal(m, b)
}
func (m *BmdBagMlacpDataMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpDataMultiple.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpDataMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpDataMultiple.Merge(m, src)
}
func (m *BmdBagMlacpDataMultiple) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpDataMultiple.Size(m)
}
func (m *BmdBagMlacpDataMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpDataMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpDataMultiple proto.InternalMessageInfo

func (m *BmdBagMlacpDataMultiple) GetMlacpData() []*BmdBagMlacpData {
	if m != nil {
		return m.MlacpData
	}
	return nil
}

func init() {
	proto.RegisterType((*BmdBagMlacpDataMultiple_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bmd_bag_mlacp_data_multiple_KEYS")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bm_system_id_st")
	proto.RegisterType((*MlacpRgNodeInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.mlacp_rg_node_info_st")
	proto.RegisterType((*BmdBagMlacpRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bmd_bag_mlacp_rg")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bm_mac_addr_st")
	proto.RegisterType((*MlacpBdlInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.mlacp_bdl_info_st")
	proto.RegisterType((*MlacpMbrInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.mlacp_mbr_info_st")
	proto.RegisterType((*BmdBagMlacpBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bmd_bag_mlacp_bdl")
	proto.RegisterType((*BmdBagMlacpData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bmd_bag_mlacp_data")
	proto.RegisterType((*BmdBagMlacpDataMultiple)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp_brief.mlacp_bundle_briefs.mlacp_bundle_brief.mlacp_bundle_item_brief.bmd_bag_mlacp_data_multiple")
}

func init() { proto.RegisterFile("bmd_bag_mlacp_data_multiple.proto", fileDescriptor_f226b3f9a25ce63b) }

var fileDescriptor_f226b3f9a25ce63b = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x8e, 0xdb, 0x36,
	0x10, 0x86, 0x6c, 0xef, 0x8f, 0xc7, 0xab, 0xf5, 0x86, 0xc9, 0x02, 0x02, 0xd2, 0x62, 0xbd, 0x0e,
	0x0a, 0x38, 0x3d, 0xb8, 0xed, 0xf6, 0x09, 0x5a, 0xb4, 0x68, 0x8d, 0x60, 0xb7, 0x81, 0x92, 0x4b,
	0x4f, 0x04, 0x25, 0x72, 0x15, 0xb6, 0x92, 0x28, 0x50, 0x74, 0x51, 0x3f, 0x42, 0x7f, 0xd1, 0xb7,
	0x68, 0x81, 0xa2, 0x39, 0x14, 0xbd, 0xe4, 0x21, 0xfa, 0x1c, 0x3d, 0xf5, 0xd4, 0x6b, 0x0f, 0x01,
	0x87, 0x94, 0x2c, 0xdb, 0x41, 0xae, 0xce, 0x4d, 0xfc, 0x86, 0xe4, 0xcc, 0x7c, 0x33, 0xf3, 0x51,
	0x70, 0x99, 0x14, 0x9c, 0x26, 0x2c, 0xa3, 0x45, 0xce, 0xd2, 0x8a, 0x72, 0x66, 0x18, 0x2d, 0x96,
	0xb9, 0x91, 0x55, 0x2e, 0xe6, 0x95, 0x56, 0x46, 0x91, 0x55, 0x2a, 0xeb, 0x54, 0x51, 0xa9, 0x6a,
	0xfa, 0xad, 0xa6, 0xc9, 0xb2, 0xe4, 0xb9, 0x28, 0x32, 0x4d, 0x55, 0x25, 0xf4, 0xdc, 0x2d, 0xa9,
	0x2c, 0x6f, 0x95, 0x2e, 0x98, 0x91, 0xaa, 0x9c, 0xbb, 0x9b, 0x12, 0x2d, 0xc5, 0x6d, 0xf3, 0xed,
	0x36, 0x21, 0x54, 0xbf, 0x02, 0xdb, 0x84, 0xa4, 0x11, 0x85, 0xc3, 0xa7, 0xd7, 0x30, 0x79, 0x4d,
	0x7c, 0xf4, 0xd1, 0xa7, 0x5f, 0x3e, 0x21, 0x0f, 0xe1, 0xac, 0x0d, 0xc2, 0x08, 0x7d, 0xcb, 0x52,
	0x11, 0x05, 0x93, 0x60, 0x36, 0x8c, 0xc7, 0x0e, 0x5f, 0x34, 0xf0, 0xf4, 0x5d, 0x38, 0x15, 0xe6,
	0x99, 0xd0, 0xb4, 0x60, 0x29, 0xe3, 0x5c, 0x53, 0x12, 0xc1, 0x91, 0xff, 0x8e, 0x82, 0x49, 0x7f,
	0x16, 0xc6, 0xcd, 0x72, 0xfa, 0x4f, 0x00, 0xe3, 0xa4, 0xa0, 0xf5, 0xaa, 0xb6, 0xe1, 0x48, 0x4e,
	0x6b, 0x43, 0x2e, 0x60, 0xe4, 0xd7, 0x95, 0x96, 0x0a, 0xbd, 0x84, 0x31, 0x38, 0xe8, 0xb1, 0x96,
	0x8a, 0xfc, 0x19, 0xc0, 0xd8, 0xef, 0x28, 0x58, 0x4a, 0xf1, 0xde, 0xde, 0x24, 0x98, 0x8d, 0xae,
	0xbe, 0x0b, 0xe6, 0x7b, 0xa3, 0x71, 0xbe, 0x99, 0x74, 0x1c, 0xba, 0x10, 0xaf, 0x59, 0xfa, 0x91,
	0xcd, 0xf4, 0xdf, 0x1e, 0x9c, 0xbb, 0x93, 0x3a, 0xa3, 0xa5, 0xe2, 0x2e, 0x0c, 0x9b, 0xef, 0x14,
	0x42, 0x67, 0x70, 0x28, 0xf7, 0x19, 0x8f, 0x10, 0xbc, 0x51, 0x5c, 0x2c, 0x38, 0x39, 0x87, 0xc3,
	0x9c, 0x57, 0xd6, 0xd8, 0x43, 0xd2, 0x0f, 0x72, 0x5e, 0x2d, 0x38, 0x79, 0x07, 0x4e, 0xbf, 0x11,
	0xba, 0x96, 0xaa, 0xa4, 0xe5, 0xb2, 0x48, 0x84, 0x8e, 0xfa, 0x78, 0x36, 0xf4, 0xe8, 0x0d, 0x82,
	0xe4, 0xb7, 0x00, 0x86, 0x2d, 0xc5, 0xd1, 0x00, 0xa9, 0xfa, 0x7e, 0x9f, 0x54, 0x6d, 0x95, 0x3c,
	0x3e, 0x76, 0xab, 0x05, 0x27, 0x6f, 0x03, 0x20, 0x0d, 0xb5, 0x61, 0x46, 0x44, 0x07, 0x98, 0xec,
	0xd0, 0x22, 0x4f, 0x2c, 0x40, 0x66, 0x70, 0x26, 0xd3, 0xb4, 0xa2, 0x99, 0x56, 0xcb, 0xca, 0x6f,
	0x3a, 0xc4, 0x4d, 0xa7, 0x16, 0xff, 0xcc, 0xc2, 0xb8, 0x73, 0xfa, 0xa2, 0x07, 0x67, 0x9b, 0x5d,
	0xad, 0x33, 0x4b, 0x75, 0xe7, 0xf8, 0x9a, 0xea, 0xf6, 0xec, 0x82, 0x93, 0xb7, 0x60, 0x58, 0xcb,
	0x32, 0xcb, 0x85, 0x51, 0x25, 0xb2, 0x1d, 0xc6, 0x6b, 0x80, 0x5c, 0xc1, 0x79, 0xaa, 0xca, 0x52,
	0xa4, 0x86, 0x1a, 0x59, 0x08, 0x4d, 0xf5, 0xb2, 0x2c, 0x65, 0x99, 0x21, 0xf1, 0x83, 0xf8, 0xae,
	0x37, 0x3e, 0xb5, 0xb6, 0xd8, 0x99, 0xc8, 0xf3, 0x00, 0x30, 0x05, 0x9c, 0xab, 0x68, 0x30, 0xe9,
	0xcf, 0x46, 0x57, 0xbf, 0xec, 0x93, 0xfe, 0x57, 0xf6, 0x61, 0x7c, 0x6c, 0x57, 0x9f, 0x30, 0xc3,
	0xec, 0x04, 0x27, 0xeb, 0xd9, 0xb2, 0x3d, 0x1a, 0xc1, 0x91, 0xfd, 0x14, 0x75, 0xed, 0xa7, 0xbe,
	0x59, 0x4e, 0xff, 0xeb, 0xc1, 0x1d, 0xef, 0x87, 0xe7, 0x6d, 0x4f, 0x5f, 0xc0, 0xc8, 0xbb, 0x2d,
	0x59, 0xd1, 0x28, 0x05, 0x38, 0xe8, 0x86, 0x15, 0x62, 0xb7, 0xe9, 0x7b, 0xbb, 0x4d, 0xff, 0x7b,
	0x00, 0xa3, 0x26, 0x08, 0xeb, 0xb9, 0xbf, 0xff, 0x19, 0xdf, 0xa4, 0x25, 0x86, 0xc2, 0x4d, 0xb7,
	0xa8, 0x6b, 0xf2, 0x00, 0x42, 0x96, 0x65, 0x5a, 0x64, 0xcc, 0x28, 0xdd, 0xcc, 0x59, 0x18, 0x9f,
	0xac, 0xc1, 0x05, 0x27, 0x97, 0x70, 0xe2, 0x2f, 0xed, 0x36, 0xb8, 0xe7, 0xca, 0xb5, 0xf8, 0x03,
	0x08, 0x2b, 0xa5, 0x0d, 0x8a, 0x9f, 0x96, 0x66, 0x85, 0xfd, 0x1d, 0xc6, 0x27, 0x16, 0x7c, 0xec,
	0xb1, 0xe9, 0xaf, 0x2d, 0xeb, 0x45, 0xa2, 0x5b, 0xd6, 0xef, 0xc3, 0x10, 0x8f, 0x76, 0x38, 0x3f,
	0xb6, 0x00, 0x32, 0xfe, 0x10, 0xce, 0x5a, 0xe9, 0xa6, 0xcf, 0x98, 0x75, 0xe8, 0xc5, 0x64, 0xdc,
	0xe2, 0x9f, 0x23, 0xbc, 0x5b, 0x9c, 0xfe, 0x6e, 0x71, 0x2e, 0x60, 0xe4, 0x7c, 0x39, 0xdd, 0x71,
	0xc9, 0x02, 0x7a, 0x73, 0xa2, 0xf3, 0x01, 0xdc, 0xb3, 0x05, 0xc1, 0x12, 0xb0, 0x7c, 0x9d, 0xce,
	0x01, 0xee, 0xbc, 0xdb, 0xb1, 0x35, 0x59, 0x91, 0xf7, 0xc0, 0xce, 0xcf, 0xad, 0xcc, 0x96, 0x5a,
	0xf0, 0x6d, 0x02, 0xc8, 0xda, 0xd4, 0x1e, 0xb8, 0x84, 0x93, 0x42, 0x58, 0x6f, 0x9e, 0xce, 0x23,
	0x47, 0xa7, 0xc3, 0x9c, 0x0e, 0xfc, 0x3c, 0x80, 0x3b, 0x9b, 0x3a, 0x90, 0xf0, 0x9c, 0xbc, 0x0f,
	0xf7, 0xb6, 0x9f, 0x33, 0xfa, 0xb5, 0x58, 0x79, 0x3d, 0x20, 0x5b, 0x4f, 0xda, 0x23, 0xb1, 0xb2,
	0xc2, 0x54, 0x08, 0x2e, 0x19, 0x35, 0xab, 0xaa, 0x21, 0x6e, 0x88, 0xc8, 0xd3, 0x55, 0x25, 0xec,
	0x85, 0x5a, 0xf0, 0x65, 0xc9, 0x59, 0x99, 0xae, 0xa8, 0x4a, 0xbe, 0xb2, 0x0a, 0xe1, 0x99, 0x1b,
	0xc4, 0x64, 0x6d, 0xfb, 0x02, 0x4d, 0x0b, 0x4e, 0x5e, 0x04, 0xed, 0xe0, 0xb8, 0x48, 0x3a, 0xea,
	0xf0, 0xe3, 0xfe, 0xd5, 0xa1, 0x33, 0xcd, 0xf1, 0x18, 0xa1, 0x8f, 0x71, 0xa7, 0x15, 0x88, 0x4e,
	0xec, 0x9e, 0x7e, 0x8c, 0xfd, 0xe0, 0x4d, 0x89, 0xbd, 0x33, 0x13, 0x3e, 0xf6, 0x6b, 0x0c, 0x13,
	0xc5, 0xed, 0xff, 0x1e, 0x90, 0xdd, 0xdf, 0x1d, 0xf2, 0x57, 0x00, 0xe3, 0xce, 0xdb, 0x80, 0x09,
	0x05, 0x28, 0x38, 0x3f, 0xec, 0x57, 0x70, 0x36, 0x9f, 0xb0, 0x38, 0x6c, 0xdf, 0x2a, 0xac, 0xc4,
	0xf3, 0xa0, 0x55, 0x5a, 0x0c, 0xb9, 0xb7, 0xff, 0x1a, 0xec, 0x4c, 0x5b, 0x23, 0xfc, 0x48, 0xff,
	0xdf, 0x01, 0xdc, 0x7f, 0xcd, 0xdf, 0x26, 0xf9, 0x23, 0x00, 0x58, 0xe3, 0xd1, 0x15, 0xe6, 0xf3,
	0xd3, 0x9b, 0x93, 0x8f, 0x8d, 0x2a, 0x1e, 0xe2, 0xb7, 0xcd, 0x27, 0x39, 0xc4, 0xdf, 0xf7, 0x0f,
	0x5f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x71, 0xcd, 0xa0, 0xf4, 0xe3, 0x0b, 0x00, 0x00,
}