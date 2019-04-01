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
// source: isis_sh_nbr.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_neighbors_neighbor

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

type IsisShNbr_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	SystemId             string   `protobuf:"bytes,2,opt,name=system_id,json=systemId,proto3" json:"system_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShNbr_KEYS) Reset()         { *m = IsisShNbr_KEYS{} }
func (m *IsisShNbr_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShNbr_KEYS) ProtoMessage()    {}
func (*IsisShNbr_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{0}
}

func (m *IsisShNbr_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShNbr_KEYS.Unmarshal(m, b)
}
func (m *IsisShNbr_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShNbr_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShNbr_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShNbr_KEYS.Merge(m, src)
}
func (m *IsisShNbr_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShNbr_KEYS.Size(m)
}
func (m *IsisShNbr_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShNbr_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShNbr_KEYS proto.InternalMessageInfo

func (m *IsisShNbr_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShNbr_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

func (m *IsisShNbr_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type NeighborActiveAreaAddressesItem struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeighborActiveAreaAddressesItem) Reset()         { *m = NeighborActiveAreaAddressesItem{} }
func (m *NeighborActiveAreaAddressesItem) String() string { return proto.CompactTextString(m) }
func (*NeighborActiveAreaAddressesItem) ProtoMessage()    {}
func (*NeighborActiveAreaAddressesItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{1}
}

func (m *NeighborActiveAreaAddressesItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighborActiveAreaAddressesItem.Unmarshal(m, b)
}
func (m *NeighborActiveAreaAddressesItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighborActiveAreaAddressesItem.Marshal(b, m, deterministic)
}
func (m *NeighborActiveAreaAddressesItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighborActiveAreaAddressesItem.Merge(m, src)
}
func (m *NeighborActiveAreaAddressesItem) XXX_Size() int {
	return xxx_messageInfo_NeighborActiveAreaAddressesItem.Size(m)
}
func (m *NeighborActiveAreaAddressesItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighborActiveAreaAddressesItem.DiscardUnknown(m)
}

var xxx_messageInfo_NeighborActiveAreaAddressesItem proto.InternalMessageInfo

func (m *NeighborActiveAreaAddressesItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type IsisTopoIdType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	SafName              string   `protobuf:"bytes,2,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	VrfName              string   `protobuf:"bytes,3,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,4,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisTopoIdType) Reset()         { *m = IsisTopoIdType{} }
func (m *IsisTopoIdType) String() string { return proto.CompactTextString(m) }
func (*IsisTopoIdType) ProtoMessage()    {}
func (*IsisTopoIdType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{2}
}

func (m *IsisTopoIdType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisTopoIdType.Unmarshal(m, b)
}
func (m *IsisTopoIdType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisTopoIdType.Marshal(b, m, deterministic)
}
func (m *IsisTopoIdType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisTopoIdType.Merge(m, src)
}
func (m *IsisTopoIdType) XXX_Size() int {
	return xxx_messageInfo_IsisTopoIdType.Size(m)
}
func (m *IsisTopoIdType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisTopoIdType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisTopoIdType proto.InternalMessageInfo

func (m *IsisTopoIdType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisTopoIdType) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisTopoIdType) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IsisTopoIdType) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

type IsisIpv4AdjSidBackupInfo struct {
	BackupLabelStackSize uint32   `protobuf:"varint,1,opt,name=backup_label_stack_size,json=backupLabelStackSize,proto3" json:"backup_label_stack_size,omitempty"`
	BackupLabelStack     []uint32 `protobuf:"varint,2,rep,packed,name=backup_label_stack,json=backupLabelStack,proto3" json:"backup_label_stack,omitempty"`
	BackupNodeAddress    string   `protobuf:"bytes,3,opt,name=backup_node_address,json=backupNodeAddress,proto3" json:"backup_node_address,omitempty"`
	BackupNexthop        string   `protobuf:"bytes,4,opt,name=backup_nexthop,json=backupNexthop,proto3" json:"backup_nexthop,omitempty"`
	BackupInterface      string   `protobuf:"bytes,5,opt,name=backup_interface,json=backupInterface,proto3" json:"backup_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisIpv4AdjSidBackupInfo) Reset()         { *m = IsisIpv4AdjSidBackupInfo{} }
func (m *IsisIpv4AdjSidBackupInfo) String() string { return proto.CompactTextString(m) }
func (*IsisIpv4AdjSidBackupInfo) ProtoMessage()    {}
func (*IsisIpv4AdjSidBackupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{3}
}

func (m *IsisIpv4AdjSidBackupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv4AdjSidBackupInfo.Unmarshal(m, b)
}
func (m *IsisIpv4AdjSidBackupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv4AdjSidBackupInfo.Marshal(b, m, deterministic)
}
func (m *IsisIpv4AdjSidBackupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv4AdjSidBackupInfo.Merge(m, src)
}
func (m *IsisIpv4AdjSidBackupInfo) XXX_Size() int {
	return xxx_messageInfo_IsisIpv4AdjSidBackupInfo.Size(m)
}
func (m *IsisIpv4AdjSidBackupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv4AdjSidBackupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv4AdjSidBackupInfo proto.InternalMessageInfo

func (m *IsisIpv4AdjSidBackupInfo) GetBackupLabelStackSize() uint32 {
	if m != nil {
		return m.BackupLabelStackSize
	}
	return 0
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupLabelStack() []uint32 {
	if m != nil {
		return m.BackupLabelStack
	}
	return nil
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupNodeAddress() string {
	if m != nil {
		return m.BackupNodeAddress
	}
	return ""
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupNexthop() string {
	if m != nil {
		return m.BackupNexthop
	}
	return ""
}

func (m *IsisIpv4AdjSidBackupInfo) GetBackupInterface() string {
	if m != nil {
		return m.BackupInterface
	}
	return ""
}

type IsisIpv4AdjSid struct {
	AdjacencySidValue    uint32                    `protobuf:"varint,1,opt,name=adjacency_sid_value,json=adjacencySidValue,proto3" json:"adjacency_sid_value,omitempty"`
	AdjacencySidBackup   *IsisIpv4AdjSidBackupInfo `protobuf:"bytes,2,opt,name=adjacency_sid_backup,json=adjacencySidBackup,proto3" json:"adjacency_sid_backup,omitempty"`
	AdjacencySidBackupTe *IsisIpv4AdjSidBackupInfo `protobuf:"bytes,3,opt,name=adjacency_sid_backup_te,json=adjacencySidBackupTe,proto3" json:"adjacency_sid_backup_te,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IsisIpv4AdjSid) Reset()         { *m = IsisIpv4AdjSid{} }
func (m *IsisIpv4AdjSid) String() string { return proto.CompactTextString(m) }
func (*IsisIpv4AdjSid) ProtoMessage()    {}
func (*IsisIpv4AdjSid) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{4}
}

func (m *IsisIpv4AdjSid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv4AdjSid.Unmarshal(m, b)
}
func (m *IsisIpv4AdjSid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv4AdjSid.Marshal(b, m, deterministic)
}
func (m *IsisIpv4AdjSid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv4AdjSid.Merge(m, src)
}
func (m *IsisIpv4AdjSid) XXX_Size() int {
	return xxx_messageInfo_IsisIpv4AdjSid.Size(m)
}
func (m *IsisIpv4AdjSid) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv4AdjSid.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv4AdjSid proto.InternalMessageInfo

func (m *IsisIpv4AdjSid) GetAdjacencySidValue() uint32 {
	if m != nil {
		return m.AdjacencySidValue
	}
	return 0
}

func (m *IsisIpv4AdjSid) GetAdjacencySidBackup() *IsisIpv4AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackup
	}
	return nil
}

func (m *IsisIpv4AdjSid) GetAdjacencySidBackupTe() *IsisIpv4AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackupTe
	}
	return nil
}

type IsisShIntfDet struct {
	InterfaceIndex       uint32   `protobuf:"varint,1,opt,name=interface_index,json=interfaceIndex,proto3" json:"interface_index,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShIntfDet) Reset()         { *m = IsisShIntfDet{} }
func (m *IsisShIntfDet) String() string { return proto.CompactTextString(m) }
func (*IsisShIntfDet) ProtoMessage()    {}
func (*IsisShIntfDet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{5}
}

func (m *IsisShIntfDet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShIntfDet.Unmarshal(m, b)
}
func (m *IsisShIntfDet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShIntfDet.Marshal(b, m, deterministic)
}
func (m *IsisShIntfDet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShIntfDet.Merge(m, src)
}
func (m *IsisShIntfDet) XXX_Size() int {
	return xxx_messageInfo_IsisShIntfDet.Size(m)
}
func (m *IsisShIntfDet) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShIntfDet.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShIntfDet proto.InternalMessageInfo

func (m *IsisShIntfDet) GetInterfaceIndex() uint32 {
	if m != nil {
		return m.InterfaceIndex
	}
	return 0
}

func (m *IsisShIntfDet) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type IsisShAdjIpv4 struct {
	NextHop                string           `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	InterfaceAddress       []string         `protobuf:"bytes,2,rep,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	AdjacencySid           *IsisIpv4AdjSid  `protobuf:"bytes,3,opt,name=adjacency_sid,json=adjacencySid,proto3" json:"adjacency_sid,omitempty"`
	NonFrrAdjacencySid     *IsisIpv4AdjSid  `protobuf:"bytes,4,opt,name=non_frr_adjacency_sid,json=nonFrrAdjacencySid,proto3" json:"non_frr_adjacency_sid,omitempty"`
	UnderlyingInterface    []*IsisShIntfDet `protobuf:"bytes,5,rep,name=underlying_interface,json=underlyingInterface,proto3" json:"underlying_interface,omitempty"`
	UnderlyingAdjacencySid []uint32         `protobuf:"varint,6,rep,packed,name=underlying_adjacency_sid,json=underlyingAdjacencySid,proto3" json:"underlying_adjacency_sid,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *IsisShAdjIpv4) Reset()         { *m = IsisShAdjIpv4{} }
func (m *IsisShAdjIpv4) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjIpv4) ProtoMessage()    {}
func (*IsisShAdjIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{6}
}

func (m *IsisShAdjIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjIpv4.Unmarshal(m, b)
}
func (m *IsisShAdjIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjIpv4.Marshal(b, m, deterministic)
}
func (m *IsisShAdjIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjIpv4.Merge(m, src)
}
func (m *IsisShAdjIpv4) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjIpv4.Size(m)
}
func (m *IsisShAdjIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjIpv4 proto.InternalMessageInfo

func (m *IsisShAdjIpv4) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *IsisShAdjIpv4) GetInterfaceAddress() []string {
	if m != nil {
		return m.InterfaceAddress
	}
	return nil
}

func (m *IsisShAdjIpv4) GetAdjacencySid() *IsisIpv4AdjSid {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv4) GetNonFrrAdjacencySid() *IsisIpv4AdjSid {
	if m != nil {
		return m.NonFrrAdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv4) GetUnderlyingInterface() []*IsisShIntfDet {
	if m != nil {
		return m.UnderlyingInterface
	}
	return nil
}

func (m *IsisShAdjIpv4) GetUnderlyingAdjacencySid() []uint32 {
	if m != nil {
		return m.UnderlyingAdjacencySid
	}
	return nil
}

type IsisIpv6AdjSidBackupInfo struct {
	BackupLabelStackSize uint32   `protobuf:"varint,1,opt,name=backup_label_stack_size,json=backupLabelStackSize,proto3" json:"backup_label_stack_size,omitempty"`
	BackupLabelStack     []uint32 `protobuf:"varint,2,rep,packed,name=backup_label_stack,json=backupLabelStack,proto3" json:"backup_label_stack,omitempty"`
	BackupNodeAddress    string   `protobuf:"bytes,3,opt,name=backup_node_address,json=backupNodeAddress,proto3" json:"backup_node_address,omitempty"`
	BackupNexthop        string   `protobuf:"bytes,4,opt,name=backup_nexthop,json=backupNexthop,proto3" json:"backup_nexthop,omitempty"`
	BackupInterface      string   `protobuf:"bytes,5,opt,name=backup_interface,json=backupInterface,proto3" json:"backup_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisIpv6AdjSidBackupInfo) Reset()         { *m = IsisIpv6AdjSidBackupInfo{} }
func (m *IsisIpv6AdjSidBackupInfo) String() string { return proto.CompactTextString(m) }
func (*IsisIpv6AdjSidBackupInfo) ProtoMessage()    {}
func (*IsisIpv6AdjSidBackupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{7}
}

func (m *IsisIpv6AdjSidBackupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv6AdjSidBackupInfo.Unmarshal(m, b)
}
func (m *IsisIpv6AdjSidBackupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv6AdjSidBackupInfo.Marshal(b, m, deterministic)
}
func (m *IsisIpv6AdjSidBackupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv6AdjSidBackupInfo.Merge(m, src)
}
func (m *IsisIpv6AdjSidBackupInfo) XXX_Size() int {
	return xxx_messageInfo_IsisIpv6AdjSidBackupInfo.Size(m)
}
func (m *IsisIpv6AdjSidBackupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv6AdjSidBackupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv6AdjSidBackupInfo proto.InternalMessageInfo

func (m *IsisIpv6AdjSidBackupInfo) GetBackupLabelStackSize() uint32 {
	if m != nil {
		return m.BackupLabelStackSize
	}
	return 0
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupLabelStack() []uint32 {
	if m != nil {
		return m.BackupLabelStack
	}
	return nil
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupNodeAddress() string {
	if m != nil {
		return m.BackupNodeAddress
	}
	return ""
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupNexthop() string {
	if m != nil {
		return m.BackupNexthop
	}
	return ""
}

func (m *IsisIpv6AdjSidBackupInfo) GetBackupInterface() string {
	if m != nil {
		return m.BackupInterface
	}
	return ""
}

type IsisIpv6AdjSid struct {
	AdjacencySidValue    uint32                    `protobuf:"varint,1,opt,name=adjacency_sid_value,json=adjacencySidValue,proto3" json:"adjacency_sid_value,omitempty"`
	AdjacencySidBackup   *IsisIpv6AdjSidBackupInfo `protobuf:"bytes,2,opt,name=adjacency_sid_backup,json=adjacencySidBackup,proto3" json:"adjacency_sid_backup,omitempty"`
	AdjacencySidBackupTe *IsisIpv6AdjSidBackupInfo `protobuf:"bytes,3,opt,name=adjacency_sid_backup_te,json=adjacencySidBackupTe,proto3" json:"adjacency_sid_backup_te,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IsisIpv6AdjSid) Reset()         { *m = IsisIpv6AdjSid{} }
func (m *IsisIpv6AdjSid) String() string { return proto.CompactTextString(m) }
func (*IsisIpv6AdjSid) ProtoMessage()    {}
func (*IsisIpv6AdjSid) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{8}
}

func (m *IsisIpv6AdjSid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisIpv6AdjSid.Unmarshal(m, b)
}
func (m *IsisIpv6AdjSid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisIpv6AdjSid.Marshal(b, m, deterministic)
}
func (m *IsisIpv6AdjSid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisIpv6AdjSid.Merge(m, src)
}
func (m *IsisIpv6AdjSid) XXX_Size() int {
	return xxx_messageInfo_IsisIpv6AdjSid.Size(m)
}
func (m *IsisIpv6AdjSid) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisIpv6AdjSid.DiscardUnknown(m)
}

var xxx_messageInfo_IsisIpv6AdjSid proto.InternalMessageInfo

func (m *IsisIpv6AdjSid) GetAdjacencySidValue() uint32 {
	if m != nil {
		return m.AdjacencySidValue
	}
	return 0
}

func (m *IsisIpv6AdjSid) GetAdjacencySidBackup() *IsisIpv6AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackup
	}
	return nil
}

func (m *IsisIpv6AdjSid) GetAdjacencySidBackupTe() *IsisIpv6AdjSidBackupInfo {
	if m != nil {
		return m.AdjacencySidBackupTe
	}
	return nil
}

type IsisShAdjIpv6 struct {
	NextHop                string           `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	InterfaceAddress       []string         `protobuf:"bytes,2,rep,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	AdjacencySid           *IsisIpv6AdjSid  `protobuf:"bytes,3,opt,name=adjacency_sid,json=adjacencySid,proto3" json:"adjacency_sid,omitempty"`
	NonFrrAdjacencySid     *IsisIpv6AdjSid  `protobuf:"bytes,4,opt,name=non_frr_adjacency_sid,json=nonFrrAdjacencySid,proto3" json:"non_frr_adjacency_sid,omitempty"`
	UnderlyingInterface    []*IsisShIntfDet `protobuf:"bytes,5,rep,name=underlying_interface,json=underlyingInterface,proto3" json:"underlying_interface,omitempty"`
	UnderlyingAdjacencySid []uint32         `protobuf:"varint,6,rep,packed,name=underlying_adjacency_sid,json=underlyingAdjacencySid,proto3" json:"underlying_adjacency_sid,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *IsisShAdjIpv6) Reset()         { *m = IsisShAdjIpv6{} }
func (m *IsisShAdjIpv6) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjIpv6) ProtoMessage()    {}
func (*IsisShAdjIpv6) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{9}
}

func (m *IsisShAdjIpv6) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjIpv6.Unmarshal(m, b)
}
func (m *IsisShAdjIpv6) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjIpv6.Marshal(b, m, deterministic)
}
func (m *IsisShAdjIpv6) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjIpv6.Merge(m, src)
}
func (m *IsisShAdjIpv6) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjIpv6.Size(m)
}
func (m *IsisShAdjIpv6) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjIpv6.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjIpv6 proto.InternalMessageInfo

func (m *IsisShAdjIpv6) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *IsisShAdjIpv6) GetInterfaceAddress() []string {
	if m != nil {
		return m.InterfaceAddress
	}
	return nil
}

func (m *IsisShAdjIpv6) GetAdjacencySid() *IsisIpv6AdjSid {
	if m != nil {
		return m.AdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv6) GetNonFrrAdjacencySid() *IsisIpv6AdjSid {
	if m != nil {
		return m.NonFrrAdjacencySid
	}
	return nil
}

func (m *IsisShAdjIpv6) GetUnderlyingInterface() []*IsisShIntfDet {
	if m != nil {
		return m.UnderlyingInterface
	}
	return nil
}

func (m *IsisShAdjIpv6) GetUnderlyingAdjacencySid() []uint32 {
	if m != nil {
		return m.UnderlyingAdjacencySid
	}
	return nil
}

type IsisShAdjAf struct {
	AfName               string         `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 *IsisShAdjIpv4 `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 *IsisShAdjIpv6 `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IsisShAdjAf) Reset()         { *m = IsisShAdjAf{} }
func (m *IsisShAdjAf) String() string { return proto.CompactTextString(m) }
func (*IsisShAdjAf) ProtoMessage()    {}
func (*IsisShAdjAf) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{10}
}

func (m *IsisShAdjAf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShAdjAf.Unmarshal(m, b)
}
func (m *IsisShAdjAf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShAdjAf.Marshal(b, m, deterministic)
}
func (m *IsisShAdjAf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShAdjAf.Merge(m, src)
}
func (m *IsisShAdjAf) XXX_Size() int {
	return xxx_messageInfo_IsisShAdjAf.Size(m)
}
func (m *IsisShAdjAf) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShAdjAf.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShAdjAf proto.InternalMessageInfo

func (m *IsisShAdjAf) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShAdjAf) GetIpv4() *IsisShAdjIpv4 {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *IsisShAdjAf) GetIpv6() *IsisShAdjIpv6 {
	if m != nil {
		return m.Ipv6
	}
	return nil
}

type IsisShNbr struct {
	NeighborSystemId             string                             `protobuf:"bytes,50,opt,name=neighbor_system_id,json=neighborSystemId,proto3" json:"neighbor_system_id,omitempty"`
	NeighborSnpa                 string                             `protobuf:"bytes,51,opt,name=neighbor_snpa,json=neighborSnpa,proto3" json:"neighbor_snpa,omitempty"`
	LocalInterface               string                             `protobuf:"bytes,52,opt,name=local_interface,json=localInterface,proto3" json:"local_interface,omitempty"`
	NeighborState                string                             `protobuf:"bytes,53,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	NeighborCircuitType          string                             `protobuf:"bytes,54,opt,name=neighbor_circuit_type,json=neighborCircuitType,proto3" json:"neighbor_circuit_type,omitempty"`
	NeighborIetfNsfCapableFlag   uint32                             `protobuf:"varint,55,opt,name=neighbor_ietf_nsf_capable_flag,json=neighborIetfNsfCapableFlag,proto3" json:"neighbor_ietf_nsf_capable_flag,omitempty"`
	NeighborMediaType            string                             `protobuf:"bytes,56,opt,name=neighbor_media_type,json=neighborMediaType,proto3" json:"neighbor_media_type,omitempty"`
	NeighborHoldtime             uint32                             `protobuf:"varint,57,opt,name=neighbor_holdtime,json=neighborHoldtime,proto3" json:"neighbor_holdtime,omitempty"`
	NeighborActiveAreaAddress    []*NeighborActiveAreaAddressesItem `protobuf:"bytes,58,rep,name=neighbor_active_area_address,json=neighborActiveAreaAddress,proto3" json:"neighbor_active_area_address,omitempty"`
	NeighborUptimeValidFlag      bool                               `protobuf:"varint,59,opt,name=neighbor_uptime_valid_flag,json=neighborUptimeValidFlag,proto3" json:"neighbor_uptime_valid_flag,omitempty"`
	NeighborUptime               uint32                             `protobuf:"varint,60,opt,name=neighbor_uptime,json=neighborUptime,proto3" json:"neighbor_uptime,omitempty"`
	TopologiesSupported          []*IsisTopoIdType                  `protobuf:"bytes,61,rep,name=topologies_supported,json=topologiesSupported,proto3" json:"topologies_supported,omitempty"`
	NeighborPerAddressFamilyData []*IsisShAdjAf                     `protobuf:"bytes,62,rep,name=neighbor_per_address_family_data,json=neighborPerAddressFamilyData,proto3" json:"neighbor_per_address_family_data,omitempty"`
	NsrStandby                   bool                               `protobuf:"varint,63,opt,name=nsr_standby,json=nsrStandby,proto3" json:"nsr_standby,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                           `json:"-"`
	XXX_unrecognized             []byte                             `json:"-"`
	XXX_sizecache                int32                              `json:"-"`
}

func (m *IsisShNbr) Reset()         { *m = IsisShNbr{} }
func (m *IsisShNbr) String() string { return proto.CompactTextString(m) }
func (*IsisShNbr) ProtoMessage()    {}
func (*IsisShNbr) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea641ebc36dc2941, []int{11}
}

func (m *IsisShNbr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShNbr.Unmarshal(m, b)
}
func (m *IsisShNbr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShNbr.Marshal(b, m, deterministic)
}
func (m *IsisShNbr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShNbr.Merge(m, src)
}
func (m *IsisShNbr) XXX_Size() int {
	return xxx_messageInfo_IsisShNbr.Size(m)
}
func (m *IsisShNbr) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShNbr.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShNbr proto.InternalMessageInfo

func (m *IsisShNbr) GetNeighborSystemId() string {
	if m != nil {
		return m.NeighborSystemId
	}
	return ""
}

func (m *IsisShNbr) GetNeighborSnpa() string {
	if m != nil {
		return m.NeighborSnpa
	}
	return ""
}

func (m *IsisShNbr) GetLocalInterface() string {
	if m != nil {
		return m.LocalInterface
	}
	return ""
}

func (m *IsisShNbr) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *IsisShNbr) GetNeighborCircuitType() string {
	if m != nil {
		return m.NeighborCircuitType
	}
	return ""
}

func (m *IsisShNbr) GetNeighborIetfNsfCapableFlag() uint32 {
	if m != nil {
		return m.NeighborIetfNsfCapableFlag
	}
	return 0
}

func (m *IsisShNbr) GetNeighborMediaType() string {
	if m != nil {
		return m.NeighborMediaType
	}
	return ""
}

func (m *IsisShNbr) GetNeighborHoldtime() uint32 {
	if m != nil {
		return m.NeighborHoldtime
	}
	return 0
}

func (m *IsisShNbr) GetNeighborActiveAreaAddress() []*NeighborActiveAreaAddressesItem {
	if m != nil {
		return m.NeighborActiveAreaAddress
	}
	return nil
}

func (m *IsisShNbr) GetNeighborUptimeValidFlag() bool {
	if m != nil {
		return m.NeighborUptimeValidFlag
	}
	return false
}

func (m *IsisShNbr) GetNeighborUptime() uint32 {
	if m != nil {
		return m.NeighborUptime
	}
	return 0
}

func (m *IsisShNbr) GetTopologiesSupported() []*IsisTopoIdType {
	if m != nil {
		return m.TopologiesSupported
	}
	return nil
}

func (m *IsisShNbr) GetNeighborPerAddressFamilyData() []*IsisShAdjAf {
	if m != nil {
		return m.NeighborPerAddressFamilyData
	}
	return nil
}

func (m *IsisShNbr) GetNsrStandby() bool {
	if m != nil {
		return m.NsrStandby
	}
	return false
}

func init() {
	proto.RegisterType((*IsisShNbr_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_nbr_KEYS")
	proto.RegisterType((*NeighborActiveAreaAddressesItem)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.neighbor_active_area_addresses_item")
	proto.RegisterType((*IsisTopoIdType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_topo_id_type")
	proto.RegisterType((*IsisIpv4AdjSidBackupInfo)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv4_adj_sid_backup_info")
	proto.RegisterType((*IsisIpv4AdjSid)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv4_adj_sid")
	proto.RegisterType((*IsisShIntfDet)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_intf_det")
	proto.RegisterType((*IsisShAdjIpv4)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_ipv4")
	proto.RegisterType((*IsisIpv6AdjSidBackupInfo)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv6_adj_sid_backup_info")
	proto.RegisterType((*IsisIpv6AdjSid)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_ipv6_adj_sid")
	proto.RegisterType((*IsisShAdjIpv6)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_ipv6")
	proto.RegisterType((*IsisShAdjAf)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_adj_af")
	proto.RegisterType((*IsisShNbr)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.neighbors.neighbor.isis_sh_nbr")
}

func init() { proto.RegisterFile("isis_sh_nbr.proto", fileDescriptor_ea641ebc36dc2941) }

var fileDescriptor_ea641ebc36dc2941 = []byte{
	// 1067 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x58, 0x4f, 0x6f, 0x1c, 0x35,
	0x14, 0xd7, 0x6c, 0xd3, 0xfc, 0x71, 0xba, 0xc9, 0xc6, 0xd9, 0x92, 0x29, 0x7f, 0xa3, 0x8d, 0xaa,
	0x06, 0x15, 0xed, 0x21, 0x6d, 0x87, 0x42, 0xf8, 0xa3, 0xb4, 0x10, 0x35, 0x02, 0x22, 0x34, 0x5b,
	0x22, 0x7a, 0xb2, 0xbc, 0x63, 0xcf, 0xc6, 0xcd, 0xac, 0x3d, 0xb2, 0xbd, 0xab, 0x6c, 0x91, 0xb8,
	0x01, 0x07, 0xb8, 0x73, 0x41, 0x7c, 0x02, 0x2e, 0x7c, 0x2e, 0xbe, 0x02, 0x07, 0x64, 0x7b, 0x3c,
	0x33, 0x9b, 0x86, 0x0a, 0xa4, 0x0d, 0x0a, 0x82, 0xdb, 0xce, 0xfb, 0x3d, 0xfb, 0xfd, 0xfc, 0xf3,
	0x7b, 0x6f, 0xe6, 0x2d, 0x58, 0x63, 0x8a, 0x29, 0xa4, 0x8e, 0x11, 0xef, 0xcb, 0x6e, 0x2e, 0x85,
	0x16, 0x70, 0x3f, 0x61, 0x2a, 0x11, 0x88, 0x09, 0x85, 0x4e, 0x25, 0x4a, 0x32, 0xae, 0x90, 0x75,
	0x12, 0x39, 0x95, 0x5d, 0xf3, 0xab, 0xcb, 0xb8, 0xd2, 0x98, 0x27, 0xb4, 0xfa, 0xd5, 0xe5, 0x94,
	0x0d, 0x8e, 0xfb, 0x42, 0xaa, 0xf2, 0x57, 0xe7, 0x2b, 0xd0, 0xaa, 0x6d, 0x8e, 0x3e, 0xf9, 0xf8,
	0x49, 0x0f, 0x6e, 0x81, 0xa6, 0x5f, 0x82, 0x38, 0x1e, 0xd2, 0x30, 0xd8, 0x0c, 0xb6, 0x97, 0xe2,
	0x6b, 0xde, 0x78, 0x88, 0x87, 0x14, 0xbe, 0x02, 0x96, 0xd4, 0x44, 0x69, 0x3a, 0x44, 0x8c, 0x84,
	0x0d, 0xeb, 0xb0, 0xe8, 0x0c, 0x07, 0x04, 0xde, 0x04, 0x2b, 0x8c, 0x6b, 0x2a, 0x53, 0xec, 0xb7,
	0xb8, 0x62, 0x3d, 0x9a, 0xa5, 0xd5, 0xec, 0xd1, 0xd9, 0x05, 0x5b, 0x9e, 0x08, 0xc2, 0x89, 0x66,
	0x63, 0x8a, 0xb0, 0xa4, 0x18, 0x61, 0x42, 0x24, 0x55, 0x8a, 0x2a, 0xc4, 0x34, 0x1d, 0xc2, 0x36,
	0xb8, 0x3a, 0xc6, 0xd9, 0xc8, 0xf3, 0x70, 0x0f, 0x9d, 0xef, 0x82, 0x42, 0x17, 0x2d, 0x72, 0x81,
	0x18, 0x41, 0x7a, 0x92, 0x53, 0xb8, 0x01, 0x16, 0x70, 0x5a, 0x67, 0x3d, 0x8f, 0x53, 0xcb, 0xf7,
	0x06, 0x58, 0x54, 0x1e, 0x71, 0x74, 0x17, 0x54, 0x05, 0x8d, 0x65, 0x5a, 0xe7, 0xb9, 0x30, 0x96,
	0x0e, 0xda, 0x02, 0x4d, 0xb3, 0x7d, 0x26, 0x06, 0x13, 0x87, 0xcf, 0x39, 0x29, 0xbc, 0xd1, 0x1e,
	0xe3, 0xdb, 0x06, 0x78, 0xcd, 0x32, 0x61, 0xf9, 0xf8, 0x2e, 0xc2, 0xe4, 0x29, 0x52, 0x8c, 0xa0,
	0x3e, 0x4e, 0x4e, 0x46, 0x39, 0x62, 0x3c, 0x15, 0xf0, 0x1e, 0xd8, 0x28, 0x1e, 0x33, 0xdc, 0xa7,
	0x19, 0x52, 0x1a, 0x27, 0x27, 0x48, 0xb1, 0x67, 0x8e, 0x65, 0x33, 0x6e, 0x3b, 0xf8, 0x53, 0x83,
	0xf6, 0x0c, 0xd8, 0x63, 0xcf, 0x28, 0x7c, 0x0b, 0xc0, 0xe7, 0x97, 0x85, 0x8d, 0xcd, 0x2b, 0xdb,
	0xcd, 0xb8, 0x75, 0x76, 0x05, 0xec, 0x82, 0xf5, 0xc2, 0x9b, 0x0b, 0x42, 0xbd, 0x88, 0xc5, 0x89,
	0xd6, 0x1c, 0x74, 0x28, 0x08, 0xdd, 0x73, 0x80, 0xb9, 0x24, 0xef, 0x4f, 0x4f, 0xf5, 0xb1, 0xc8,
	0x8b, 0xc3, 0x35, 0x0b, 0x57, 0x67, 0x84, 0x6f, 0x82, 0x56, 0x79, 0x94, 0xe2, 0xf2, 0xc2, 0xab,
	0xd6, 0x71, 0xd5, 0xd9, 0x0f, 0xbc, 0xb9, 0xf3, 0x7b, 0xa3, 0xb8, 0x92, 0xba, 0x10, 0x86, 0x17,
	0x26, 0x4f, 0x71, 0x42, 0x79, 0x32, 0xb1, 0xca, 0x54, 0x97, 0xd9, 0x8c, 0xd7, 0x4a, 0xa8, 0xc7,
	0xc8, 0x91, 0x01, 0xe0, 0x8f, 0x01, 0x68, 0x4f, 0x2f, 0x70, 0x71, 0xec, 0xb5, 0x2d, 0xef, 0xd0,
	0xee, 0x6c, 0x52, 0xbf, 0xfb, 0xc2, 0x2b, 0x8b, 0x61, 0x9d, 0xd8, 0x03, 0x0b, 0xc0, 0x9f, 0x02,
	0xb0, 0x71, 0x1e, 0x33, 0xa4, 0x5d, 0xe2, 0xfc, 0x63, 0xe4, 0xda, 0xcf, 0x93, 0x7b, 0x4c, 0x3b,
	0xfd, 0xaa, 0x96, 0x19, 0xd7, 0x29, 0x22, 0x54, 0xc3, 0x5b, 0x60, 0xb5, 0xaa, 0x44, 0xc6, 0x09,
	0x3d, 0x2d, 0x84, 0xaf, 0x0a, 0xf4, 0xc0, 0x58, 0xcf, 0x29, 0xd9, 0xc6, 0x79, 0x25, 0xfb, 0xeb,
	0x5c, 0x15, 0xc4, 0x30, 0x33, 0x14, 0x4d, 0x01, 0x99, 0x14, 0x42, 0x26, 0x87, 0x5c, 0xd5, 0x2d,
	0x98, 0xe7, 0x47, 0x22, 0x87, 0xb7, 0xc1, 0x5a, 0xb5, 0xad, 0x4f, 0x49, 0x93, 0xc1, 0x4b, 0x71,
	0xab, 0x04, 0x7c, 0x46, 0x7e, 0x0d, 0x9a, 0x53, 0xf2, 0x16, 0xa2, 0x3e, 0xb9, 0x30, 0x51, 0xe3,
	0x6b, 0x75, 0x21, 0xe1, 0x0f, 0x01, 0xb8, 0xce, 0x05, 0x47, 0xa9, 0x94, 0x68, 0x9a, 0xc8, 0xdc,
	0x45, 0x13, 0x81, 0x5c, 0xf0, 0x7d, 0x29, 0xf7, 0xea, 0x74, 0xbe, 0x0f, 0x40, 0x7b, 0xc4, 0x09,
	0x95, 0xd9, 0x84, 0xf1, 0xc1, 0x54, 0xf9, 0x5d, 0xd9, 0x5e, 0xde, 0xf9, 0x72, 0xa6, 0x6c, 0x6a,
	0x49, 0x13, 0xaf, 0x57, 0x51, 0xcb, 0xe2, 0x86, 0xf7, 0x41, 0x58, 0x23, 0x33, 0x2d, 0xcf, 0xbc,
	0x6d, 0x49, 0x2f, 0x55, 0x78, 0xfd, 0x1c, 0x53, 0xfd, 0x31, 0xfa, 0xbf, 0x3f, 0x9e, 0x11, 0xe2,
	0xd2, 0xf7, 0xc7, 0xe8, 0x32, 0xf7, 0xc7, 0xe8, 0x6f, 0xf4, 0xc7, 0x73, 0x7a, 0x57, 0xf4, 0xaf,
	0xeb, 0x5d, 0xd1, 0x65, 0xe9, 0x5d, 0xd1, 0x7f, 0xa2, 0x77, 0x7d, 0xd3, 0x00, 0x2b, 0xf5, 0x9c,
	0xc1, 0xe9, 0x9f, 0x7f, 0x62, 0x66, 0x60, 0xce, 0xf4, 0xf4, 0xa2, 0x0e, 0x67, 0x7e, 0x44, 0xff,
	0xba, 0x8d, 0x6d, 0x94, 0x22, 0x5a, 0x54, 0xe4, 0xd9, 0x45, 0x45, 0x8b, 0x6c, 0xb4, 0xa8, 0xf3,
	0xdb, 0x02, 0x58, 0xae, 0x0d, 0x0a, 0xa6, 0xf5, 0x96, 0x9f, 0xee, 0xd5, 0x1c, 0xb0, 0x63, 0xf5,
	0x68, 0x79, 0xa4, 0xe7, 0xe7, 0x81, 0x2d, 0xd0, 0xac, 0xbc, 0x79, 0x8e, 0xc3, 0x3b, 0xee, 0x33,
	0xba, 0x74, 0xe4, 0x39, 0x36, 0x9f, 0x2a, 0x99, 0x48, 0x70, 0x56, 0x4b, 0x96, 0xbb, 0xd6, 0x6d,
	0xc5, 0x9a, 0xab, 0xdb, 0xbc, 0x09, 0x56, 0xaa, 0xdd, 0x34, 0xd6, 0x34, 0xbc, 0xe7, 0x1a, 0x73,
	0xb9, 0x9d, 0x31, 0xc2, 0x1d, 0x70, 0xbd, 0x74, 0x4b, 0x98, 0x4c, 0x46, 0x4c, 0xdb, 0x19, 0x21,
	0x8c, 0xac, 0xf7, 0xba, 0x07, 0x1f, 0x3a, 0xec, 0xb1, 0x19, 0x1f, 0x1e, 0x80, 0xd7, 0xcb, 0x35,
	0x8c, 0xea, 0x14, 0x71, 0x95, 0xa2, 0x04, 0xe7, 0xb8, 0x9f, 0x51, 0x94, 0x66, 0x78, 0x10, 0xbe,
	0x6d, 0xdb, 0xf2, 0xcb, 0xde, 0xeb, 0x80, 0xea, 0xf4, 0x50, 0xa5, 0x0f, 0x9d, 0xcb, 0x7e, 0x86,
	0x07, 0xa6, 0x9f, 0x97, 0x7b, 0x0c, 0x29, 0x61, 0xd8, 0x45, 0xbd, 0xef, 0xde, 0x33, 0x1e, 0xfa,
	0xcc, 0x20, 0x36, 0xe6, 0x6d, 0x50, 0x1a, 0xd1, 0xb1, 0xc8, 0x88, 0x66, 0x43, 0x1a, 0xbe, 0x63,
	0xc3, 0x94, 0x4a, 0x3e, 0x2a, 0xec, 0xf0, 0x97, 0x00, 0xbc, 0xfa, 0xa2, 0x99, 0x29, 0x7c, 0xd7,
	0xd6, 0xd7, 0xc9, 0xac, 0xd2, 0xe1, 0x2f, 0xcc, 0x67, 0xf1, 0x0d, 0xef, 0xb4, 0x67, 0x7d, 0xf6,
	0x24, 0xc5, 0xbe, 0x2b, 0xee, 0x82, 0x52, 0x29, 0x34, 0xca, 0xcd, 0x09, 0xcc, 0xdb, 0x8d, 0x11,
	0xa7, 0xe5, 0xee, 0x66, 0xb0, 0xbd, 0x18, 0x6f, 0x78, 0x8f, 0x2f, 0xac, 0xc3, 0x91, 0xc1, 0xad,
	0x90, 0xb7, 0xc0, 0xea, 0x99, 0xc5, 0xe1, 0x7b, 0xee, 0xdb, 0x75, 0x7a, 0x85, 0xe9, 0x7d, 0xed,
	0x62, 0x22, 0x63, 0x54, 0x21, 0x35, 0xca, 0x73, 0x21, 0x35, 0x25, 0xe1, 0xfb, 0x56, 0x8c, 0xd9,
	0xb6, 0xbe, 0xfa, 0xb8, 0x19, 0xaf, 0x57, 0x61, 0x7b, 0x3e, 0x2a, 0xfc, 0x39, 0x00, 0x9b, 0x25,
	0xf1, 0x9c, 0x4a, 0xaf, 0x17, 0x4a, 0xf1, 0x90, 0x65, 0x13, 0x44, 0xb0, 0xc6, 0xe1, 0x07, 0x96,
	0xda, 0xd1, 0x45, 0x94, 0x2d, 0x4e, 0xe3, 0x32, 0x47, 0x3e, 0xa7, 0xb2, 0xb8, 0x8b, 0x7d, 0x1b,
	0xfc, 0x23, 0xac, 0x31, 0x7c, 0x03, 0x2c, 0x73, 0x65, 0x6b, 0x87, 0x93, 0xfe, 0x24, 0xfc, 0xd0,
	0x5e, 0x03, 0xe0, 0xca, 0x14, 0x8e, 0xb1, 0xf4, 0xe7, 0xed, 0x9f, 0x0c, 0x77, 0xfe, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x3c, 0x25, 0x43, 0x79, 0x10, 0x00, 0x00,
}
