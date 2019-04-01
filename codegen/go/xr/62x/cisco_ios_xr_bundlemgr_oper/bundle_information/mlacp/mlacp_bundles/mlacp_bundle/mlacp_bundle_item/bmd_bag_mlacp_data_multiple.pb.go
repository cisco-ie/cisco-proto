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

package cisco_ios_xr_bundlemgr_oper_bundle_information_mlacp_mlacp_bundles_mlacp_bundle_mlacp_bundle_item

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
	proto.RegisterType((*BmdBagMlacpDataMultiple_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bmd_bag_mlacp_data_multiple_KEYS")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bm_system_id_st")
	proto.RegisterType((*MlacpRgNodeInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.mlacp_rg_node_info_st")
	proto.RegisterType((*BmdBagMlacpRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bmd_bag_mlacp_rg")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bm_mac_addr_st")
	proto.RegisterType((*MlacpBdlInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.mlacp_bdl_info_st")
	proto.RegisterType((*MlacpMbrInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.mlacp_mbr_info_st")
	proto.RegisterType((*BmdBagMlacpBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bmd_bag_mlacp_bdl")
	proto.RegisterType((*BmdBagMlacpData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bmd_bag_mlacp_data")
	proto.RegisterType((*BmdBagMlacpDataMultiple)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_bundles.mlacp_bundle.mlacp_bundle_item.bmd_bag_mlacp_data_multiple")
}

func init() { proto.RegisterFile("bmd_bag_mlacp_data_multiple.proto", fileDescriptor_f226b3f9a25ce63b) }

var fileDescriptor_f226b3f9a25ce63b = []byte{
	// 898 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcf, 0x8e, 0xe4, 0xb4,
	0x13, 0x56, 0xba, 0xe7, 0x5f, 0x57, 0x4f, 0xa6, 0x67, 0xbd, 0x3b, 0x52, 0xa4, 0xfd, 0xfd, 0x34,
	0x3d, 0xbd, 0x42, 0xea, 0xe5, 0xd0, 0xc0, 0xf0, 0x04, 0x20, 0x10, 0xb4, 0x56, 0x33, 0xac, 0xb2,
	0x7b, 0xe1, 0x64, 0x39, 0xb1, 0x27, 0x6b, 0x88, 0xe3, 0xe0, 0xb8, 0xd1, 0xf6, 0x53, 0x20, 0x21,
	0xf6, 0x80, 0xb8, 0xec, 0x8d, 0x17, 0xe0, 0x29, 0x38, 0x70, 0xe2, 0x84, 0x78, 0x18, 0xe4, 0xb2,
	0x93, 0x4e, 0x4f, 0xa3, 0xbd, 0xd1, 0x97, 0x28, 0xf5, 0x95, 0xed, 0xaa, 0xfa, 0x52, 0x5f, 0x39,
	0x70, 0x95, 0x29, 0x4e, 0x33, 0x56, 0x50, 0x55, 0xb2, 0xbc, 0xa6, 0x9c, 0x59, 0x46, 0xd5, 0xaa,
	0xb4, 0xb2, 0x2e, 0xc5, 0xa2, 0x36, 0xda, 0x6a, 0xc2, 0x72, 0xd9, 0xe4, 0x9a, 0x4a, 0xdd, 0xd0,
	0xd7, 0x86, 0x66, 0xab, 0x8a, 0x97, 0x42, 0x15, 0x86, 0xea, 0x5a, 0x98, 0x85, 0x37, 0xa9, 0xac,
	0xee, 0xb4, 0x51, 0xcc, 0x4a, 0x5d, 0x2d, 0xf0, 0x24, 0xff, 0x0c, 0xab, 0x9b, 0x2d, 0x6b, 0xcb,
	0xa0, 0xd2, 0x0a, 0x35, 0xbb, 0x81, 0xe9, 0x3b, 0xf2, 0xa0, 0xcf, 0x3e, 0xff, 0xfa, 0x05, 0x79,
	0x0a, 0xe7, 0x5d, 0x30, 0x2b, 0xcc, 0x1d, 0xcb, 0x45, 0x12, 0x4d, 0xa3, 0xf9, 0x28, 0x9d, 0x78,
	0x7c, 0xd9, 0xc2, 0xb3, 0xf7, 0xe1, 0x4c, 0xd8, 0x57, 0xc2, 0x50, 0xc5, 0x72, 0xc6, 0xb9, 0xa1,
	0x24, 0x81, 0xe3, 0xf0, 0x9e, 0x44, 0xd3, 0xe1, 0x3c, 0x4e, 0x5b, 0x73, 0xf6, 0x47, 0x04, 0x93,
	0x4c, 0xd1, 0x66, 0xdd, 0x58, 0xa1, 0xa8, 0xe4, 0xb4, 0xb1, 0xe4, 0x12, 0xc6, 0xc1, 0xae, 0x8d,
	0xd4, 0x18, 0x25, 0x4e, 0xc1, 0x43, 0xcf, 0x8d, 0xd4, 0xe4, 0xe7, 0x08, 0x26, 0x61, 0x85, 0x62,
	0x39, 0xc5, 0x73, 0x07, 0xd3, 0x68, 0x3e, 0xbe, 0xfe, 0x6e, 0xf1, 0x9f, 0xb3, 0xb5, 0xd8, 0xae,
	0x2d, 0x8d, 0x7d, 0x26, 0x37, 0x2c, 0xff, 0xc4, 0x15, 0xf4, 0xe7, 0x00, 0x2e, 0xfc, 0x1e, 0x53,
	0xd0, 0x4a, 0x73, 0x1f, 0xce, 0x95, 0x35, 0x83, 0xd8, 0x3b, 0x3c, 0xca, 0x43, 0x61, 0x63, 0x04,
	0x6f, 0x35, 0x17, 0x4b, 0x4e, 0x2e, 0xe0, 0xa8, 0xe4, 0xb5, 0x73, 0x0e, 0x90, 0xdb, 0xc3, 0x92,
	0xd7, 0x4b, 0x4e, 0xde, 0x83, 0xb3, 0xef, 0x85, 0x69, 0xa4, 0xae, 0x68, 0xb5, 0x52, 0x99, 0x30,
	0xc9, 0x10, 0xf7, 0xc6, 0x01, 0xbd, 0x45, 0x90, 0xfc, 0x10, 0xc1, 0xa8, 0x63, 0x32, 0x39, 0x40,
	0x46, 0xcc, 0x1e, 0x18, 0xb9, 0xf7, 0x01, 0xd3, 0x13, 0x6f, 0x2d, 0x39, 0xf9, 0x3f, 0x00, 0x56,
	0xdb, 0x58, 0x66, 0x45, 0x72, 0x88, 0x35, 0x8d, 0x1c, 0xf2, 0xc2, 0x01, 0x64, 0x0e, 0xe7, 0x32,
	0xcf, 0x6b, 0x5a, 0x18, 0xbd, 0xaa, 0xc3, 0xa2, 0x23, 0x5c, 0x74, 0xe6, 0xf0, 0x2f, 0x1c, 0x8c,
	0x2b, 0x67, 0x6f, 0x07, 0x70, 0xbe, 0xdd, 0xa3, 0xa6, 0x70, 0x8c, 0xf6, 0xb6, 0x6f, 0x18, 0xed,
	0xf6, 0x2e, 0x39, 0xf9, 0x1f, 0x8c, 0x1a, 0x59, 0x15, 0xa5, 0xb0, 0xba, 0x42, 0x52, 0xe3, 0x74,
	0x03, 0x90, 0x6b, 0xb8, 0xc8, 0x75, 0x55, 0x89, 0xdc, 0x52, 0x2b, 0x95, 0x30, 0xd4, 0xac, 0xaa,
	0x4a, 0x56, 0x05, 0xf2, 0x7b, 0x90, 0x3e, 0x0c, 0xce, 0x97, 0xce, 0x97, 0x7a, 0x17, 0x79, 0x13,
	0x01, 0x96, 0x80, 0x2a, 0x49, 0x0e, 0xa6, 0xc3, 0xf9, 0xf8, 0xfa, 0xf5, 0x1e, 0x58, 0xfe, 0xd7,
	0xae, 0x4a, 0x4f, 0x9c, 0xf5, 0x19, 0xb3, 0xcc, 0xc9, 0x2e, 0xdb, 0x08, 0xc2, 0x75, 0x5c, 0x02,
	0xc7, 0xee, 0x55, 0x34, 0x4d, 0x90, 0x6a, 0x6b, 0xce, 0xfe, 0x1a, 0xc0, 0x83, 0x10, 0x81, 0x97,
	0x5d, 0x87, 0x5e, 0xc2, 0x38, 0x04, 0xac, 0x98, 0x6a, 0xe5, 0x0d, 0x1e, 0xba, 0x65, 0x4a, 0xec,
	0xb6, 0xf0, 0x60, 0xb7, 0x85, 0x7f, 0x8c, 0x60, 0xdc, 0x26, 0xe1, 0x22, 0x0f, 0xf7, 0x26, 0xcc,
	0xed, 0xea, 0x53, 0x50, 0x5e, 0x92, 0xa2, 0x69, 0xc8, 0x13, 0x88, 0x59, 0x51, 0x18, 0x51, 0x30,
	0xab, 0x4d, 0x2b, 0x8e, 0x38, 0x3d, 0xdd, 0x80, 0x4b, 0x4e, 0xae, 0xe0, 0x34, 0x1c, 0xd7, 0x6f,
	0xd7, 0x40, 0x89, 0x6f, 0xd8, 0x27, 0x10, 0xd7, 0xda, 0x58, 0x1c, 0x4c, 0x46, 0xda, 0x35, 0x76,
	0x6b, 0x9c, 0x9e, 0x3a, 0xf0, 0x79, 0xc0, 0x66, 0xbf, 0x76, 0xe4, 0xaa, 0xcc, 0x74, 0xe4, 0x3e,
	0x86, 0x11, 0x6e, 0xed, 0x51, 0x7b, 0xe2, 0x00, 0x24, 0xf6, 0x29, 0x9c, 0x77, 0x63, 0x95, 0xbe,
	0x62, 0x2e, 0x60, 0x98, 0x00, 0x93, 0x0e, 0xff, 0x12, 0xe1, 0xdd, 0x6f, 0x30, 0xdc, 0xfd, 0x06,
	0x97, 0x30, 0xf6, 0xb1, 0xfc, 0xb0, 0xf0, 0xc5, 0x02, 0x46, 0xf3, 0x93, 0xe2, 0x23, 0x78, 0xe4,
	0x88, 0x47, 0xaa, 0x59, 0xb9, 0x29, 0xe7, 0x10, 0x57, 0x3e, 0xec, 0xf9, 0xda, 0xaa, 0xc8, 0x07,
	0xe0, 0xd4, 0x70, 0x27, 0x8b, 0x95, 0x11, 0xfc, 0x3e, 0x01, 0x64, 0xe3, 0xea, 0x36, 0x5c, 0xc1,
	0xa9, 0x12, 0x2e, 0x5a, 0xa0, 0xf3, 0xd8, 0xd3, 0xe9, 0x31, 0xaf, 0xea, 0xbf, 0x87, 0xf0, 0x60,
	0x5b, 0xd5, 0x19, 0x2f, 0xc9, 0x87, 0xf0, 0xe8, 0xfe, 0x55, 0x43, 0xbf, 0x15, 0xeb, 0xa0, 0x6e,
	0x72, 0xef, 0xba, 0x79, 0x26, 0xd6, 0x6e, 0xcc, 0x28, 0xc1, 0x25, 0xa3, 0x76, 0x5d, 0xb7, 0xc4,
	0x8d, 0x10, 0x79, 0xb9, 0xae, 0x85, 0x3b, 0xd0, 0x08, 0xbe, 0xaa, 0x38, 0xab, 0xf2, 0x35, 0xd5,
	0xd9, 0x37, 0x4e, 0xef, 0x81, 0xb9, 0x83, 0x94, 0x6c, 0x7c, 0x5f, 0xa1, 0x6b, 0xc9, 0xc9, 0xdb,
	0xa8, 0xd3, 0x87, 0xcf, 0xa4, 0xa7, 0x75, 0xbb, 0x37, 0xad, 0xf7, 0xb4, 0x99, 0x4e, 0x10, 0xfa,
	0x14, 0xd7, 0x38, 0xb9, 0xf7, 0x52, 0x0c, 0x2c, 0x63, 0x8a, 0x87, 0x7b, 0x4e, 0xb1, 0xd7, 0xe1,
	0x21, 0xc5, 0x1b, 0xcc, 0x06, 0x27, 0xd2, 0xef, 0x03, 0x20, 0xbb, 0x3f, 0x16, 0xe4, 0x97, 0x08,
	0x26, 0xbd, 0xb9, 0x8d, 0x79, 0x47, 0x38, 0x25, 0x9a, 0xbd, 0x4c, 0x89, 0xed, 0x5b, 0x24, 0x8d,
	0xbb, 0xeb, 0x02, 0x79, 0x7d, 0x13, 0x75, 0x53, 0x10, 0x33, 0x1b, 0xec, 0x8d, 0xd1, 0x1d, 0x25,
	0xb4, 0xb3, 0x17, 0xc9, 0xfc, 0x2d, 0x82, 0xc7, 0xef, 0xf8, 0x4b, 0x23, 0x3f, 0x45, 0x00, 0x1b,
	0x3c, 0xb9, 0xc6, 0xb4, 0x57, 0x7b, 0x4f, 0xdb, 0x05, 0x4f, 0x47, 0xf8, 0xee, 0xd2, 0xce, 0x8e,
	0xf0, 0x2f, 0xf6, 0xe3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x20, 0x69, 0xef, 0xea, 0x0a,
	0x00, 0x00,
}