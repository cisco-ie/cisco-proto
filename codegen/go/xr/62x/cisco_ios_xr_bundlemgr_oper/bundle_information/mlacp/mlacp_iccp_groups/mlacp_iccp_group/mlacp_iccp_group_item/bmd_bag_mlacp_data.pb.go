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
// source: bmd_bag_mlacp_data.proto

package cisco_ios_xr_bundlemgr_oper_bundle_information_mlacp_mlacp_iccp_groups_mlacp_iccp_group_mlacp_iccp_group_item

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

type BmdBagMlacpData_KEYS struct {
	IccpGroup            uint32   `protobuf:"varint,1,opt,name=iccp_group,json=iccpGroup,proto3" json:"iccp_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BmdBagMlacpData_KEYS) Reset()         { *m = BmdBagMlacpData_KEYS{} }
func (m *BmdBagMlacpData_KEYS) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpData_KEYS) ProtoMessage()    {}
func (*BmdBagMlacpData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f07b39087c2bc24, []int{0}
}

func (m *BmdBagMlacpData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BmdBagMlacpData_KEYS.Unmarshal(m, b)
}
func (m *BmdBagMlacpData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BmdBagMlacpData_KEYS.Marshal(b, m, deterministic)
}
func (m *BmdBagMlacpData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BmdBagMlacpData_KEYS.Merge(m, src)
}
func (m *BmdBagMlacpData_KEYS) XXX_Size() int {
	return xxx_messageInfo_BmdBagMlacpData_KEYS.Size(m)
}
func (m *BmdBagMlacpData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BmdBagMlacpData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BmdBagMlacpData_KEYS proto.InternalMessageInfo

func (m *BmdBagMlacpData_KEYS) GetIccpGroup() uint32 {
	if m != nil {
		return m.IccpGroup
	}
	return 0
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
	return fileDescriptor_4f07b39087c2bc24, []int{1}
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
	return fileDescriptor_4f07b39087c2bc24, []int{2}
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
	return fileDescriptor_4f07b39087c2bc24, []int{3}
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
	return fileDescriptor_4f07b39087c2bc24, []int{4}
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
	return fileDescriptor_4f07b39087c2bc24, []int{5}
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
	return fileDescriptor_4f07b39087c2bc24, []int{6}
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
	return fileDescriptor_4f07b39087c2bc24, []int{7}
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
	return fileDescriptor_4f07b39087c2bc24, []int{8}
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
	IccpGroupData        *BmdBagMlacpRg    `protobuf:"bytes,50,opt,name=iccp_group_data,json=iccpGroupData,proto3" json:"iccp_group_data,omitempty"`
	BundleData           []*BmdBagMlacpBdl `protobuf:"bytes,51,rep,name=bundle_data,json=bundleData,proto3" json:"bundle_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BmdBagMlacpData) Reset()         { *m = BmdBagMlacpData{} }
func (m *BmdBagMlacpData) String() string { return proto.CompactTextString(m) }
func (*BmdBagMlacpData) ProtoMessage()    {}
func (*BmdBagMlacpData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f07b39087c2bc24, []int{9}
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

func init() {
	proto.RegisterType((*BmdBagMlacpData_KEYS)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bmd_bag_mlacp_data_KEYS")
	proto.RegisterType((*EtherMacaddr_)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.ether_macaddr_")
	proto.RegisterType((*BmSystemIdSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bm_system_id_st")
	proto.RegisterType((*MlacpRgNodeInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.mlacp_rg_node_info_st")
	proto.RegisterType((*BmdBagMlacpRg)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bmd_bag_mlacp_rg")
	proto.RegisterType((*BmMacAddrSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bm_mac_addr_st")
	proto.RegisterType((*MlacpBdlInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.mlacp_bdl_info_st")
	proto.RegisterType((*MlacpMbrInfoSt)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.mlacp_mbr_info_st")
	proto.RegisterType((*BmdBagMlacpBdl)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bmd_bag_mlacp_bdl")
	proto.RegisterType((*BmdBagMlacpData)(nil), "cisco_ios_xr_bundlemgr_oper.bundle_information.mlacp.mlacp_iccp_groups.mlacp_iccp_group.mlacp_iccp_group_item.bmd_bag_mlacp_data")
}

func init() { proto.RegisterFile("bmd_bag_mlacp_data.proto", fileDescriptor_4f07b39087c2bc24) }

var fileDescriptor_4f07b39087c2bc24 = []byte{
	// 870 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4d, 0x8f, 0x23, 0x35,
	0x10, 0x55, 0x27, 0xf3, 0x95, 0xca, 0xf4, 0x64, 0xd6, 0xbb, 0x23, 0x22, 0x01, 0x9a, 0xd9, 0xac,
	0x90, 0x02, 0x87, 0x00, 0xb3, 0x17, 0xae, 0x20, 0x10, 0x44, 0xab, 0x1d, 0x56, 0xbd, 0x7b, 0xe1,
	0x64, 0xb9, 0xdb, 0x9e, 0x5e, 0x43, 0x6c, 0xb7, 0xdc, 0x0e, 0x22, 0x07, 0x3e, 0x0e, 0xfc, 0x0b,
	0x24, 0x0e, 0x2b, 0xa1, 0x3d, 0xed, 0x1f, 0xe2, 0x84, 0x84, 0xc4, 0x4f, 0xe0, 0x8c, 0x5c, 0x76,
	0x77, 0x3a, 0xd3, 0x9c, 0x67, 0x2e, 0x51, 0xfc, 0xca, 0x76, 0xbd, 0x7a, 0xae, 0x7a, 0x09, 0x4c,
	0x73, 0xc5, 0x69, 0xce, 0x4a, 0xaa, 0x56, 0xac, 0xa8, 0x28, 0x67, 0x8e, 0x2d, 0x2a, 0x6b, 0x9c,
	0x21, 0xaa, 0x90, 0x75, 0x61, 0xa8, 0x34, 0x35, 0xfd, 0xc1, 0xd2, 0x7c, 0xad, 0xf9, 0x4a, 0xa8,
	0xd2, 0x52, 0x53, 0x09, 0xbb, 0x08, 0x4b, 0x2a, 0xf5, 0xb5, 0xb1, 0x8a, 0x39, 0x69, 0xf4, 0x02,
	0x2f, 0x08, 0x9f, 0x54, 0x16, 0x45, 0x45, 0x4b, 0x6b, 0xd6, 0x55, 0xdd, 0x43, 0x7a, 0x00, 0x95,
	0x4e, 0xa8, 0xd9, 0x27, 0xf0, 0x56, 0x9f, 0x0a, 0x7d, 0xf2, 0xc5, 0x37, 0xcf, 0xc9, 0xbb, 0x00,
	0xdb, 0xdd, 0xd3, 0xe4, 0x22, 0x99, 0xa7, 0xd9, 0xc8, 0x23, 0x5f, 0x7a, 0x60, 0xf6, 0x01, 0x9c,
	0x08, 0xf7, 0x52, 0x58, 0xaa, 0x58, 0xc1, 0x38, 0xb7, 0x94, 0x4c, 0xe1, 0x30, 0x7e, 0x9f, 0x26,
	0x17, 0xc3, 0x79, 0x9a, 0x35, 0xcb, 0xd9, 0x9f, 0x09, 0x4c, 0x72, 0x45, 0xeb, 0x4d, 0xed, 0x84,
	0xa2, 0x92, 0xd3, 0xda, 0x91, 0x73, 0x18, 0xc7, 0x75, 0x65, 0xa5, 0x89, 0xf7, 0x43, 0x80, 0x9e,
	0x59, 0x69, 0xc8, 0x1f, 0x09, 0x4c, 0xe2, 0x0e, 0xc5, 0x0a, 0x8a, 0xf7, 0x0e, 0x2e, 0x92, 0xf9,
	0xf8, 0xf2, 0xc7, 0xc5, 0xad, 0x8a, 0xb4, 0xd8, 0xad, 0x33, 0x4b, 0x03, 0xab, 0xa7, 0xac, 0xf8,
	0xd4, 0x17, 0xf7, 0xd7, 0x00, 0xce, 0xc2, 0x39, 0x5b, 0x52, 0x6d, 0x78, 0x48, 0xed, 0x4b, 0x9c,
	0x41, 0x1a, 0x02, 0x01, 0xe5, 0xb1, 0xc8, 0x31, 0x82, 0x57, 0x86, 0x8b, 0x25, 0x27, 0x67, 0x70,
	0xb0, 0xe2, 0x95, 0x0f, 0xfa, 0xda, 0x46, 0xd9, 0xfe, 0x8a, 0x57, 0x4b, 0x4e, 0xde, 0x83, 0x93,
	0xef, 0x85, 0xad, 0xa5, 0xd1, 0x54, 0xaf, 0x55, 0x2e, 0xec, 0x74, 0x88, 0x67, 0xd3, 0x88, 0x5e,
	0x21, 0x48, 0x7e, 0x4b, 0x60, 0xd4, 0xaa, 0x3a, 0xdd, 0x43, 0x75, 0x7e, 0xba, 0x65, 0x75, 0x6e,
	0x3c, 0x6c, 0x76, 0x14, 0x56, 0x4b, 0xee, 0x3b, 0x08, 0x2b, 0xaf, 0x1d, 0x73, 0x62, 0xba, 0x8f,
	0xf5, 0x8d, 0x3c, 0xf2, 0xdc, 0x03, 0x64, 0x0e, 0xa7, 0x9d, 0xbb, 0xc2, 0xa6, 0x03, 0xdc, 0x74,
	0xd2, 0xb6, 0x19, 0xee, 0x9c, 0xbd, 0x19, 0xc0, 0xe9, 0x6e, 0x9b, 0xda, 0xd2, 0xab, 0xdb, 0xa5,
	0xd2, 0xaa, 0xdb, 0x9e, 0x5d, 0x72, 0xf2, 0x0e, 0x8c, 0x6a, 0xa9, 0xcb, 0x95, 0x70, 0x46, 0xa3,
	0xc0, 0x69, 0xb6, 0x05, 0xc8, 0x25, 0x9c, 0x15, 0x46, 0x6b, 0x51, 0x38, 0xea, 0xa4, 0x12, 0x96,
	0xda, 0xb5, 0xd6, 0x52, 0x97, 0xa8, 0xf5, 0x5e, 0x76, 0x3f, 0x06, 0x5f, 0xf8, 0x58, 0x16, 0x42,
	0xe4, 0x55, 0x02, 0x58, 0x02, 0x0e, 0xca, 0x74, 0xef, 0x62, 0x38, 0x1f, 0x5f, 0xfe, 0x9a, 0xdc,
	0xb2, 0xe4, 0xff, 0xdb, 0x6e, 0xd9, 0x91, 0x5f, 0x7d, 0xce, 0x1c, 0xf3, 0xb3, 0x99, 0x6f, 0xa7,
	0xc6, 0xb7, 0xe2, 0x14, 0x0e, 0xfd, 0x57, 0x51, 0xd7, 0x28, 0xd3, 0x28, 0x6b, 0x96, 0xb3, 0x7f,
	0x06, 0x70, 0x2f, 0xdc, 0x97, 0xf3, 0x55, 0xdb, 0xba, 0xe7, 0x30, 0x8e, 0xb4, 0x35, 0x53, 0x22,
	0x9e, 0x81, 0x00, 0x5d, 0x31, 0x25, 0xfa, 0xbd, 0x3d, 0xe8, 0xf7, 0xf6, 0xef, 0x09, 0x8c, 0x1b,
	0x12, 0x3e, 0xf3, 0xf0, 0x4e, 0xa6, 0x77, 0x57, 0x89, 0x0c, 0x54, 0x98, 0x5b, 0x51, 0xd7, 0xe4,
	0x11, 0xa4, 0xac, 0x2c, 0xad, 0x28, 0x99, 0x33, 0xb6, 0x99, 0xa0, 0x34, 0x3b, 0xde, 0x82, 0x4b,
	0x4e, 0x1e, 0xc2, 0x71, 0x24, 0xd5, 0xed, 0xe3, 0x28, 0x4f, 0xe8, 0xe4, 0x47, 0x90, 0x56, 0xc6,
	0x3a, 0x74, 0x32, 0x2b, 0xdd, 0x06, 0xdb, 0x38, 0xcd, 0x8e, 0x3d, 0xf8, 0x2c, 0x62, 0xb3, 0xd7,
	0xad, 0xd0, 0x2a, 0xb7, 0xad, 0xd0, 0x6f, 0xc3, 0x08, 0x8f, 0x76, 0x64, 0x3e, 0xf2, 0x00, 0x8a,
	0xfc, 0x3e, 0x9c, 0x4a, 0xed, 0x84, 0xbd, 0x66, 0x85, 0xa0, 0x2f, 0x99, 0x4f, 0x18, 0x6d, 0x62,
	0xd2, 0xe2, 0x5f, 0x21, 0xdc, 0x7f, 0x8f, 0x61, 0xff, 0x3d, 0xce, 0x61, 0x1c, 0x72, 0x05, 0x47,
	0x09, 0xc5, 0x02, 0x66, 0x0b, 0x76, 0xf2, 0x31, 0x3c, 0xf0, 0x8f, 0x80, 0xb2, 0xb3, 0xd5, 0xb6,
	0x9c, 0x7d, 0xdc, 0x79, 0xbf, 0x13, 0x6b, 0xaa, 0x22, 0x1f, 0x82, 0x1f, 0x93, 0x6b, 0x59, 0xae,
	0xad, 0xe0, 0x37, 0x05, 0x20, 0xdb, 0x50, 0x7b, 0xe0, 0x21, 0x1c, 0x2b, 0xe1, 0xb3, 0x45, 0x39,
	0x0f, 0x83, 0x9c, 0x01, 0x0b, 0xe3, 0xfe, 0xef, 0x10, 0xee, 0xed, 0x8e, 0x7b, 0xce, 0x57, 0xe4,
	0x23, 0x78, 0xd0, 0x36, 0x47, 0xa3, 0xc9, 0x77, 0x62, 0x13, 0xc7, 0x9e, 0x84, 0xd8, 0xb2, 0x09,
	0x3d, 0x11, 0x1b, 0xef, 0x3f, 0x4a, 0x70, 0xc9, 0xa8, 0xdb, 0x54, 0x8d, 0x70, 0x23, 0x44, 0x5e,
	0x6c, 0x2a, 0xe1, 0x2f, 0xb4, 0x82, 0xaf, 0x35, 0x67, 0xba, 0xd8, 0x50, 0x93, 0x7f, 0xeb, 0x8d,
	0x20, 0x2a, 0xb7, 0x97, 0x91, 0x6d, 0xec, 0x6b, 0x0c, 0x2d, 0x39, 0x79, 0x93, 0xb4, 0xb3, 0x12,
	0x98, 0x74, 0x4c, 0xe0, 0x97, 0xbb, 0x31, 0x81, 0xce, 0xd0, 0x66, 0x13, 0x84, 0x3e, 0xc3, 0x4c,
	0xde, 0x07, 0x3a, 0x7c, 0xa3, 0xe4, 0xc8, 0x77, 0xff, 0x2e, 0xf9, 0x76, 0x7a, 0x3f, 0xf2, 0x7d,
	0x8a, 0xd4, 0xd0, 0xb7, 0xfe, 0x1e, 0x00, 0xe9, 0xff, 0x1d, 0x21, 0xaf, 0x13, 0x98, 0x74, 0x6e,
	0xc3, 0x22, 0x2e, 0xd1, 0x4b, 0x7e, 0xbe, 0x75, 0x2f, 0xd9, 0xfd, 0x11, 0xca, 0xd2, 0xf6, 0xd7,
	0x06, 0x05, 0x7f, 0x95, 0xb4, 0xbe, 0x89, 0x2c, 0x1f, 0xdf, 0x8d, 0xd4, 0xbd, 0xe1, 0x69, 0xac,
	0xdb, 0x93, 0xcc, 0x0f, 0xf0, 0x9f, 0xe6, 0xe3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xf4,
	0x09, 0xbf, 0x85, 0x0a, 0x00, 0x00,
}
