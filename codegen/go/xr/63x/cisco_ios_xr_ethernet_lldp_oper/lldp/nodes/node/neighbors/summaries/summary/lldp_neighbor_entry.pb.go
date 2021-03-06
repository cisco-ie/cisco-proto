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
// source: lldp_neighbor_entry.proto

package cisco_ios_xr_ethernet_lldp_oper_lldp_nodes_node_neighbors_summaries_summary

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

type LldpNeighborEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	DeviceId             string   `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpNeighborEntry_KEYS) Reset()         { *m = LldpNeighborEntry_KEYS{} }
func (m *LldpNeighborEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*LldpNeighborEntry_KEYS) ProtoMessage()    {}
func (*LldpNeighborEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{0}
}

func (m *LldpNeighborEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighborEntry_KEYS.Unmarshal(m, b)
}
func (m *LldpNeighborEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighborEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *LldpNeighborEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighborEntry_KEYS.Merge(m, src)
}
func (m *LldpNeighborEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_LldpNeighborEntry_KEYS.Size(m)
}
func (m *LldpNeighborEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighborEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighborEntry_KEYS proto.InternalMessageInfo

func (m *LldpNeighborEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LldpNeighborEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *LldpNeighborEntry_KEYS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type LldpL3Addr struct {
	AddressType          string   `protobuf:"bytes,1,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpL3Addr) Reset()         { *m = LldpL3Addr{} }
func (m *LldpL3Addr) String() string { return proto.CompactTextString(m) }
func (*LldpL3Addr) ProtoMessage()    {}
func (*LldpL3Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{1}
}

func (m *LldpL3Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpL3Addr.Unmarshal(m, b)
}
func (m *LldpL3Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpL3Addr.Marshal(b, m, deterministic)
}
func (m *LldpL3Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpL3Addr.Merge(m, src)
}
func (m *LldpL3Addr) XXX_Size() int {
	return xxx_messageInfo_LldpL3Addr.Size(m)
}
func (m *LldpL3Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpL3Addr.DiscardUnknown(m)
}

var xxx_messageInfo_LldpL3Addr proto.InternalMessageInfo

func (m *LldpL3Addr) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *LldpL3Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *LldpL3Addr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type LldpAddrEntryItem struct {
	Address              *LldpL3Addr `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MaSubtype            uint32      `protobuf:"varint,2,opt,name=ma_subtype,json=maSubtype,proto3" json:"ma_subtype,omitempty"`
	IfNum                uint32      `protobuf:"varint,3,opt,name=if_num,json=ifNum,proto3" json:"if_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LldpAddrEntryItem) Reset()         { *m = LldpAddrEntryItem{} }
func (m *LldpAddrEntryItem) String() string { return proto.CompactTextString(m) }
func (*LldpAddrEntryItem) ProtoMessage()    {}
func (*LldpAddrEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{2}
}

func (m *LldpAddrEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpAddrEntryItem.Unmarshal(m, b)
}
func (m *LldpAddrEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpAddrEntryItem.Marshal(b, m, deterministic)
}
func (m *LldpAddrEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpAddrEntryItem.Merge(m, src)
}
func (m *LldpAddrEntryItem) XXX_Size() int {
	return xxx_messageInfo_LldpAddrEntryItem.Size(m)
}
func (m *LldpAddrEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpAddrEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_LldpAddrEntryItem proto.InternalMessageInfo

func (m *LldpAddrEntryItem) GetAddress() *LldpL3Addr {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LldpAddrEntryItem) GetMaSubtype() uint32 {
	if m != nil {
		return m.MaSubtype
	}
	return 0
}

func (m *LldpAddrEntryItem) GetIfNum() uint32 {
	if m != nil {
		return m.IfNum
	}
	return 0
}

type LldpAddrEntryEntry struct {
	LldpAddrEntry        []*LldpAddrEntryItem `protobuf:"bytes,4,rep,name=lldp_addr_entry,json=lldpAddrEntry,proto3" json:"lldp_addr_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LldpAddrEntryEntry) Reset()         { *m = LldpAddrEntryEntry{} }
func (m *LldpAddrEntryEntry) String() string { return proto.CompactTextString(m) }
func (*LldpAddrEntryEntry) ProtoMessage()    {}
func (*LldpAddrEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{3}
}

func (m *LldpAddrEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpAddrEntryEntry.Unmarshal(m, b)
}
func (m *LldpAddrEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpAddrEntryEntry.Marshal(b, m, deterministic)
}
func (m *LldpAddrEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpAddrEntryEntry.Merge(m, src)
}
func (m *LldpAddrEntryEntry) XXX_Size() int {
	return xxx_messageInfo_LldpAddrEntryEntry.Size(m)
}
func (m *LldpAddrEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpAddrEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LldpAddrEntryEntry proto.InternalMessageInfo

func (m *LldpAddrEntryEntry) GetLldpAddrEntry() []*LldpAddrEntryItem {
	if m != nil {
		return m.LldpAddrEntry
	}
	return nil
}

type LldpNeighborDetail struct {
	PeerMacAddress            string              `protobuf:"bytes,1,opt,name=peer_mac_address,json=peerMacAddress,proto3" json:"peer_mac_address,omitempty"`
	PortDescription           string              `protobuf:"bytes,2,opt,name=port_description,json=portDescription,proto3" json:"port_description,omitempty"`
	SystemName                string              `protobuf:"bytes,3,opt,name=system_name,json=systemName,proto3" json:"system_name,omitempty"`
	SystemDescription         string              `protobuf:"bytes,4,opt,name=system_description,json=systemDescription,proto3" json:"system_description,omitempty"`
	TimeRemaining             uint32              `protobuf:"varint,5,opt,name=time_remaining,json=timeRemaining,proto3" json:"time_remaining,omitempty"`
	SystemCapabilities        string              `protobuf:"bytes,6,opt,name=system_capabilities,json=systemCapabilities,proto3" json:"system_capabilities,omitempty"`
	EnabledCapabilities       string              `protobuf:"bytes,7,opt,name=enabled_capabilities,json=enabledCapabilities,proto3" json:"enabled_capabilities,omitempty"`
	NetworkAddresses          *LldpAddrEntryEntry `protobuf:"bytes,8,opt,name=network_addresses,json=networkAddresses,proto3" json:"network_addresses,omitempty"`
	AutoNegotiation           string              `protobuf:"bytes,9,opt,name=auto_negotiation,json=autoNegotiation,proto3" json:"auto_negotiation,omitempty"`
	PhysicalMediaCapabilities string              `protobuf:"bytes,10,opt,name=physical_media_capabilities,json=physicalMediaCapabilities,proto3" json:"physical_media_capabilities,omitempty"`
	MediaAttachmentUnitType   uint32              `protobuf:"varint,11,opt,name=media_attachment_unit_type,json=mediaAttachmentUnitType,proto3" json:"media_attachment_unit_type,omitempty"`
	PortVlanId                uint32              `protobuf:"varint,12,opt,name=port_vlan_id,json=portVlanId,proto3" json:"port_vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}            `json:"-"`
	XXX_unrecognized          []byte              `json:"-"`
	XXX_sizecache             int32               `json:"-"`
}

func (m *LldpNeighborDetail) Reset()         { *m = LldpNeighborDetail{} }
func (m *LldpNeighborDetail) String() string { return proto.CompactTextString(m) }
func (*LldpNeighborDetail) ProtoMessage()    {}
func (*LldpNeighborDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{4}
}

func (m *LldpNeighborDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighborDetail.Unmarshal(m, b)
}
func (m *LldpNeighborDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighborDetail.Marshal(b, m, deterministic)
}
func (m *LldpNeighborDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighborDetail.Merge(m, src)
}
func (m *LldpNeighborDetail) XXX_Size() int {
	return xxx_messageInfo_LldpNeighborDetail.Size(m)
}
func (m *LldpNeighborDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighborDetail.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighborDetail proto.InternalMessageInfo

func (m *LldpNeighborDetail) GetPeerMacAddress() string {
	if m != nil {
		return m.PeerMacAddress
	}
	return ""
}

func (m *LldpNeighborDetail) GetPortDescription() string {
	if m != nil {
		return m.PortDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func (m *LldpNeighborDetail) GetSystemDescription() string {
	if m != nil {
		return m.SystemDescription
	}
	return ""
}

func (m *LldpNeighborDetail) GetTimeRemaining() uint32 {
	if m != nil {
		return m.TimeRemaining
	}
	return 0
}

func (m *LldpNeighborDetail) GetSystemCapabilities() string {
	if m != nil {
		return m.SystemCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetNetworkAddresses() *LldpAddrEntryEntry {
	if m != nil {
		return m.NetworkAddresses
	}
	return nil
}

func (m *LldpNeighborDetail) GetAutoNegotiation() string {
	if m != nil {
		return m.AutoNegotiation
	}
	return ""
}

func (m *LldpNeighborDetail) GetPhysicalMediaCapabilities() string {
	if m != nil {
		return m.PhysicalMediaCapabilities
	}
	return ""
}

func (m *LldpNeighborDetail) GetMediaAttachmentUnitType() uint32 {
	if m != nil {
		return m.MediaAttachmentUnitType
	}
	return 0
}

func (m *LldpNeighborDetail) GetPortVlanId() uint32 {
	if m != nil {
		return m.PortVlanId
	}
	return 0
}

type LldpUnknownTlvEntryItem struct {
	TlvType              uint32   `protobuf:"varint,1,opt,name=tlv_type,json=tlvType,proto3" json:"tlv_type,omitempty"`
	TlvValue             []byte   `protobuf:"bytes,2,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpUnknownTlvEntryItem) Reset()         { *m = LldpUnknownTlvEntryItem{} }
func (m *LldpUnknownTlvEntryItem) String() string { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntryItem) ProtoMessage()    {}
func (*LldpUnknownTlvEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{5}
}

func (m *LldpUnknownTlvEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpUnknownTlvEntryItem.Unmarshal(m, b)
}
func (m *LldpUnknownTlvEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpUnknownTlvEntryItem.Marshal(b, m, deterministic)
}
func (m *LldpUnknownTlvEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpUnknownTlvEntryItem.Merge(m, src)
}
func (m *LldpUnknownTlvEntryItem) XXX_Size() int {
	return xxx_messageInfo_LldpUnknownTlvEntryItem.Size(m)
}
func (m *LldpUnknownTlvEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpUnknownTlvEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_LldpUnknownTlvEntryItem proto.InternalMessageInfo

func (m *LldpUnknownTlvEntryItem) GetTlvType() uint32 {
	if m != nil {
		return m.TlvType
	}
	return 0
}

func (m *LldpUnknownTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpUnknownTlvEntryEntry struct {
	LldpUnknownTlvEntry  []*LldpUnknownTlvEntryItem `protobuf:"bytes,3,rep,name=lldp_unknown_tlv_entry,json=lldpUnknownTlvEntry,proto3" json:"lldp_unknown_tlv_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LldpUnknownTlvEntryEntry) Reset()         { *m = LldpUnknownTlvEntryEntry{} }
func (m *LldpUnknownTlvEntryEntry) String() string { return proto.CompactTextString(m) }
func (*LldpUnknownTlvEntryEntry) ProtoMessage()    {}
func (*LldpUnknownTlvEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{6}
}

func (m *LldpUnknownTlvEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpUnknownTlvEntryEntry.Unmarshal(m, b)
}
func (m *LldpUnknownTlvEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpUnknownTlvEntryEntry.Marshal(b, m, deterministic)
}
func (m *LldpUnknownTlvEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpUnknownTlvEntryEntry.Merge(m, src)
}
func (m *LldpUnknownTlvEntryEntry) XXX_Size() int {
	return xxx_messageInfo_LldpUnknownTlvEntryEntry.Size(m)
}
func (m *LldpUnknownTlvEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpUnknownTlvEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LldpUnknownTlvEntryEntry proto.InternalMessageInfo

func (m *LldpUnknownTlvEntryEntry) GetLldpUnknownTlvEntry() []*LldpUnknownTlvEntryItem {
	if m != nil {
		return m.LldpUnknownTlvEntry
	}
	return nil
}

type LldpOrgDefTlvEntryItem struct {
	Oui                  uint32   `protobuf:"varint,1,opt,name=oui,proto3" json:"oui,omitempty"`
	TlvSubtype           uint32   `protobuf:"varint,2,opt,name=tlv_subtype,json=tlvSubtype,proto3" json:"tlv_subtype,omitempty"`
	TlvInfoIndes         uint32   `protobuf:"varint,3,opt,name=tlv_info_indes,json=tlvInfoIndes,proto3" json:"tlv_info_indes,omitempty"`
	TlvValue             []byte   `protobuf:"bytes,4,opt,name=tlv_value,json=tlvValue,proto3" json:"tlv_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LldpOrgDefTlvEntryItem) Reset()         { *m = LldpOrgDefTlvEntryItem{} }
func (m *LldpOrgDefTlvEntryItem) String() string { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntryItem) ProtoMessage()    {}
func (*LldpOrgDefTlvEntryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{7}
}

func (m *LldpOrgDefTlvEntryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpOrgDefTlvEntryItem.Unmarshal(m, b)
}
func (m *LldpOrgDefTlvEntryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpOrgDefTlvEntryItem.Marshal(b, m, deterministic)
}
func (m *LldpOrgDefTlvEntryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpOrgDefTlvEntryItem.Merge(m, src)
}
func (m *LldpOrgDefTlvEntryItem) XXX_Size() int {
	return xxx_messageInfo_LldpOrgDefTlvEntryItem.Size(m)
}
func (m *LldpOrgDefTlvEntryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpOrgDefTlvEntryItem.DiscardUnknown(m)
}

var xxx_messageInfo_LldpOrgDefTlvEntryItem proto.InternalMessageInfo

func (m *LldpOrgDefTlvEntryItem) GetOui() uint32 {
	if m != nil {
		return m.Oui
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvSubtype() uint32 {
	if m != nil {
		return m.TlvSubtype
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvInfoIndes() uint32 {
	if m != nil {
		return m.TlvInfoIndes
	}
	return 0
}

func (m *LldpOrgDefTlvEntryItem) GetTlvValue() []byte {
	if m != nil {
		return m.TlvValue
	}
	return nil
}

type LldpOrgDefTlvEntryEntry struct {
	LldpOrgDefTlvEntry   []*LldpOrgDefTlvEntryItem `protobuf:"bytes,5,rep,name=lldp_org_def_tlv_entry,json=lldpOrgDefTlvEntry,proto3" json:"lldp_org_def_tlv_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LldpOrgDefTlvEntryEntry) Reset()         { *m = LldpOrgDefTlvEntryEntry{} }
func (m *LldpOrgDefTlvEntryEntry) String() string { return proto.CompactTextString(m) }
func (*LldpOrgDefTlvEntryEntry) ProtoMessage()    {}
func (*LldpOrgDefTlvEntryEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{8}
}

func (m *LldpOrgDefTlvEntryEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpOrgDefTlvEntryEntry.Unmarshal(m, b)
}
func (m *LldpOrgDefTlvEntryEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpOrgDefTlvEntryEntry.Marshal(b, m, deterministic)
}
func (m *LldpOrgDefTlvEntryEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpOrgDefTlvEntryEntry.Merge(m, src)
}
func (m *LldpOrgDefTlvEntryEntry) XXX_Size() int {
	return xxx_messageInfo_LldpOrgDefTlvEntryEntry.Size(m)
}
func (m *LldpOrgDefTlvEntryEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpOrgDefTlvEntryEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LldpOrgDefTlvEntryEntry proto.InternalMessageInfo

func (m *LldpOrgDefTlvEntryEntry) GetLldpOrgDefTlvEntry() []*LldpOrgDefTlvEntryItem {
	if m != nil {
		return m.LldpOrgDefTlvEntry
	}
	return nil
}

type LldpNeighborMib struct {
	RemTimeMark          uint32                    `protobuf:"varint,1,opt,name=rem_time_mark,json=remTimeMark,proto3" json:"rem_time_mark,omitempty"`
	RemLocalPortNum      uint32                    `protobuf:"varint,2,opt,name=rem_local_port_num,json=remLocalPortNum,proto3" json:"rem_local_port_num,omitempty"`
	RemIndex             uint32                    `protobuf:"varint,3,opt,name=rem_index,json=remIndex,proto3" json:"rem_index,omitempty"`
	ChassisIdSubType     uint32                    `protobuf:"varint,4,opt,name=chassis_id_sub_type,json=chassisIdSubType,proto3" json:"chassis_id_sub_type,omitempty"`
	ChassisIdLen         uint32                    `protobuf:"varint,5,opt,name=chassis_id_len,json=chassisIdLen,proto3" json:"chassis_id_len,omitempty"`
	PortIdSubType        uint32                    `protobuf:"varint,6,opt,name=port_id_sub_type,json=portIdSubType,proto3" json:"port_id_sub_type,omitempty"`
	PortIdLen            uint32                    `protobuf:"varint,7,opt,name=port_id_len,json=portIdLen,proto3" json:"port_id_len,omitempty"`
	CombinedCapabilities uint32                    `protobuf:"varint,8,opt,name=combined_capabilities,json=combinedCapabilities,proto3" json:"combined_capabilities,omitempty"`
	UnknownTlvList       *LldpUnknownTlvEntryEntry `protobuf:"bytes,9,opt,name=unknown_tlv_list,json=unknownTlvList,proto3" json:"unknown_tlv_list,omitempty"`
	OrgDefTlvList        *LldpOrgDefTlvEntryEntry  `protobuf:"bytes,10,opt,name=org_def_tlv_list,json=orgDefTlvList,proto3" json:"org_def_tlv_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LldpNeighborMib) Reset()         { *m = LldpNeighborMib{} }
func (m *LldpNeighborMib) String() string { return proto.CompactTextString(m) }
func (*LldpNeighborMib) ProtoMessage()    {}
func (*LldpNeighborMib) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{9}
}

func (m *LldpNeighborMib) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighborMib.Unmarshal(m, b)
}
func (m *LldpNeighborMib) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighborMib.Marshal(b, m, deterministic)
}
func (m *LldpNeighborMib) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighborMib.Merge(m, src)
}
func (m *LldpNeighborMib) XXX_Size() int {
	return xxx_messageInfo_LldpNeighborMib.Size(m)
}
func (m *LldpNeighborMib) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighborMib.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighborMib proto.InternalMessageInfo

func (m *LldpNeighborMib) GetRemTimeMark() uint32 {
	if m != nil {
		return m.RemTimeMark
	}
	return 0
}

func (m *LldpNeighborMib) GetRemLocalPortNum() uint32 {
	if m != nil {
		return m.RemLocalPortNum
	}
	return 0
}

func (m *LldpNeighborMib) GetRemIndex() uint32 {
	if m != nil {
		return m.RemIndex
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdSubType() uint32 {
	if m != nil {
		return m.ChassisIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetChassisIdLen() uint32 {
	if m != nil {
		return m.ChassisIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdSubType() uint32 {
	if m != nil {
		return m.PortIdSubType
	}
	return 0
}

func (m *LldpNeighborMib) GetPortIdLen() uint32 {
	if m != nil {
		return m.PortIdLen
	}
	return 0
}

func (m *LldpNeighborMib) GetCombinedCapabilities() uint32 {
	if m != nil {
		return m.CombinedCapabilities
	}
	return 0
}

func (m *LldpNeighborMib) GetUnknownTlvList() *LldpUnknownTlvEntryEntry {
	if m != nil {
		return m.UnknownTlvList
	}
	return nil
}

func (m *LldpNeighborMib) GetOrgDefTlvList() *LldpOrgDefTlvEntryEntry {
	if m != nil {
		return m.OrgDefTlvList
	}
	return nil
}

type LldpNeighborItem struct {
	ReceivingInterfaceName       string              `protobuf:"bytes,1,opt,name=receiving_interface_name,json=receivingInterfaceName,proto3" json:"receiving_interface_name,omitempty"`
	ReceivingParentInterfaceName string              `protobuf:"bytes,2,opt,name=receiving_parent_interface_name,json=receivingParentInterfaceName,proto3" json:"receiving_parent_interface_name,omitempty"`
	DeviceId                     string              `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ChassisId                    string              `protobuf:"bytes,4,opt,name=chassis_id,json=chassisId,proto3" json:"chassis_id,omitempty"`
	PortIdDetail                 string              `protobuf:"bytes,5,opt,name=port_id_detail,json=portIdDetail,proto3" json:"port_id_detail,omitempty"`
	HeaderVersion                uint32              `protobuf:"varint,6,opt,name=header_version,json=headerVersion,proto3" json:"header_version,omitempty"`
	HoldTime                     uint32              `protobuf:"varint,7,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	EnabledCapabilities          string              `protobuf:"bytes,8,opt,name=enabled_capabilities,json=enabledCapabilities,proto3" json:"enabled_capabilities,omitempty"`
	Platform                     string              `protobuf:"bytes,9,opt,name=platform,proto3" json:"platform,omitempty"`
	Detail                       *LldpNeighborDetail `protobuf:"bytes,10,opt,name=detail,proto3" json:"detail,omitempty"`
	Mib                          *LldpNeighborMib    `protobuf:"bytes,11,opt,name=mib,proto3" json:"mib,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}            `json:"-"`
	XXX_unrecognized             []byte              `json:"-"`
	XXX_sizecache                int32               `json:"-"`
}

func (m *LldpNeighborItem) Reset()         { *m = LldpNeighborItem{} }
func (m *LldpNeighborItem) String() string { return proto.CompactTextString(m) }
func (*LldpNeighborItem) ProtoMessage()    {}
func (*LldpNeighborItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{10}
}

func (m *LldpNeighborItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighborItem.Unmarshal(m, b)
}
func (m *LldpNeighborItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighborItem.Marshal(b, m, deterministic)
}
func (m *LldpNeighborItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighborItem.Merge(m, src)
}
func (m *LldpNeighborItem) XXX_Size() int {
	return xxx_messageInfo_LldpNeighborItem.Size(m)
}
func (m *LldpNeighborItem) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighborItem.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighborItem proto.InternalMessageInfo

func (m *LldpNeighborItem) GetReceivingInterfaceName() string {
	if m != nil {
		return m.ReceivingInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetReceivingParentInterfaceName() string {
	if m != nil {
		return m.ReceivingParentInterfaceName
	}
	return ""
}

func (m *LldpNeighborItem) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LldpNeighborItem) GetChassisId() string {
	if m != nil {
		return m.ChassisId
	}
	return ""
}

func (m *LldpNeighborItem) GetPortIdDetail() string {
	if m != nil {
		return m.PortIdDetail
	}
	return ""
}

func (m *LldpNeighborItem) GetHeaderVersion() uint32 {
	if m != nil {
		return m.HeaderVersion
	}
	return 0
}

func (m *LldpNeighborItem) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LldpNeighborItem) GetEnabledCapabilities() string {
	if m != nil {
		return m.EnabledCapabilities
	}
	return ""
}

func (m *LldpNeighborItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *LldpNeighborItem) GetDetail() *LldpNeighborDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *LldpNeighborItem) GetMib() *LldpNeighborMib {
	if m != nil {
		return m.Mib
	}
	return nil
}

type LldpNeighborEntry struct {
	LldpNeighbor         []*LldpNeighborItem `protobuf:"bytes,50,rep,name=lldp_neighbor,json=lldpNeighbor,proto3" json:"lldp_neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LldpNeighborEntry) Reset()         { *m = LldpNeighborEntry{} }
func (m *LldpNeighborEntry) String() string { return proto.CompactTextString(m) }
func (*LldpNeighborEntry) ProtoMessage()    {}
func (*LldpNeighborEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_d82a467d203f2c78, []int{11}
}

func (m *LldpNeighborEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LldpNeighborEntry.Unmarshal(m, b)
}
func (m *LldpNeighborEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LldpNeighborEntry.Marshal(b, m, deterministic)
}
func (m *LldpNeighborEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LldpNeighborEntry.Merge(m, src)
}
func (m *LldpNeighborEntry) XXX_Size() int {
	return xxx_messageInfo_LldpNeighborEntry.Size(m)
}
func (m *LldpNeighborEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LldpNeighborEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LldpNeighborEntry proto.InternalMessageInfo

func (m *LldpNeighborEntry) GetLldpNeighbor() []*LldpNeighborItem {
	if m != nil {
		return m.LldpNeighbor
	}
	return nil
}

func init() {
	proto.RegisterType((*LldpNeighborEntry_KEYS)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_neighbor_entry_KEYS")
	proto.RegisterType((*LldpL3Addr)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_l3_addr")
	proto.RegisterType((*LldpAddrEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_addr_entry_item")
	proto.RegisterType((*LldpAddrEntryEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_addr_entry_entry")
	proto.RegisterType((*LldpNeighborDetail)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_neighbor_detail")
	proto.RegisterType((*LldpUnknownTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_unknown_tlv_entry_item")
	proto.RegisterType((*LldpUnknownTlvEntryEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_unknown_tlv_entry_entry")
	proto.RegisterType((*LldpOrgDefTlvEntryItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_org_def_tlv_entry_item")
	proto.RegisterType((*LldpOrgDefTlvEntryEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_org_def_tlv_entry_entry")
	proto.RegisterType((*LldpNeighborMib)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_neighbor_mib")
	proto.RegisterType((*LldpNeighborItem)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_neighbor_item")
	proto.RegisterType((*LldpNeighborEntry)(nil), "cisco_ios_xr_ethernet_lldp_oper.lldp.nodes.node.neighbors.summaries.summary.lldp_neighbor_entry")
}

func init() { proto.RegisterFile("lldp_neighbor_entry.proto", fileDescriptor_d82a467d203f2c78) }

var fileDescriptor_d82a467d203f2c78 = []byte{
	// 1203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x57, 0xcf, 0x8f, 0x1b, 0x35,
	0x14, 0xd6, 0x90, 0xec, 0x6e, 0xf2, 0xf2, 0xa3, 0xa9, 0xb7, 0x2d, 0x69, 0xb7, 0xd0, 0x12, 0x51,
	0x51, 0x84, 0x1a, 0xc4, 0x16, 0x55, 0x48, 0x48, 0x48, 0x15, 0xed, 0x21, 0x6a, 0xbb, 0x54, 0xd3,
	0x6d, 0xa5, 0x5e, 0xb0, 0x9c, 0x19, 0x27, 0xb1, 0x76, 0xec, 0x19, 0x79, 0x3c, 0x69, 0x23, 0xae,
	0x5c, 0x38, 0x21, 0xd1, 0x6b, 0xb9, 0xc0, 0xdf, 0x01, 0x12, 0x7f, 0x15, 0x47, 0xd0, 0xb3, 0x3d,
	0xb3, 0x49, 0x9a, 0x95, 0x38, 0x6c, 0xb8, 0x4c, 0x26, 0xdf, 0xfb, 0xfc, 0x9e, 0xfd, 0xf9, 0xbd,
	0x67, 0x0f, 0x5c, 0x4d, 0x92, 0x38, 0xa3, 0x8a, 0x8b, 0xe9, 0x6c, 0x9c, 0x6a, 0xca, 0x95, 0xd1,
	0x8b, 0x61, 0xa6, 0x53, 0x93, 0x92, 0x47, 0x91, 0xc8, 0xa3, 0x94, 0x8a, 0x34, 0xa7, 0xaf, 0x35,
	0xe5, 0x66, 0xc6, 0xb5, 0xe2, 0x86, 0xda, 0x01, 0x69, 0xc6, 0xf5, 0x10, 0xdf, 0x86, 0x2a, 0x8d,
	0x79, 0x6e, 0x9f, 0xc3, 0xd2, 0x4b, 0x3e, 0xcc, 0x0b, 0x29, 0x99, 0x16, 0xbc, 0x7c, 0x5b, 0x0c,
	0x7e, 0x80, 0xfe, 0x86, 0x48, 0xf4, 0xd1, 0xc3, 0x97, 0xcf, 0xc8, 0x01, 0x34, 0x71, 0x3c, 0x55,
	0x4c, 0xf2, 0x7e, 0x70, 0x33, 0xb8, 0xdd, 0x0c, 0x1b, 0x08, 0x1c, 0x31, 0xc9, 0xc9, 0x2d, 0xe8,
	0x0a, 0x65, 0xb8, 0x9e, 0xb0, 0xc8, 0x33, 0xde, 0xb3, 0x8c, 0x4e, 0x85, 0x5a, 0xda, 0x01, 0x34,
	0x63, 0x3e, 0x17, 0x11, 0xa7, 0x22, 0xee, 0xd7, 0x9c, 0x0f, 0x07, 0x8c, 0xe2, 0xc1, 0x2b, 0x68,
	0xdb, 0xe0, 0xc9, 0x5d, 0xca, 0xe2, 0x58, 0x93, 0x8f, 0xa0, 0x8d, 0xbf, 0x3c, 0xcf, 0xa9, 0x59,
	0x64, 0x65, 0xcc, 0x96, 0xc7, 0x8e, 0x17, 0x19, 0x47, 0x8a, 0xc8, 0xe6, 0x5f, 0x52, 0x8f, 0xf9,
	0xa0, 0x2d, 0xc4, 0xee, 0x3b, 0xc8, 0x53, 0xee, 0x55, 0x94, 0x5a, 0x45, 0xb9, 0xe7, 0x29, 0x83,
	0xbf, 0x02, 0xb8, 0x64, 0x23, 0x23, 0xc7, 0x2f, 0x59, 0x18, 0x2e, 0x49, 0x0e, 0x7b, 0xe5, 0x30,
	0x0c, 0xde, 0x3a, 0x7c, 0x39, 0x3c, 0x47, 0xb5, 0x87, 0xcb, 0xab, 0x0d, 0xcb, 0x48, 0xe4, 0x03,
	0x00, 0xc9, 0x68, 0x5e, 0x8c, 0xed, 0xa2, 0x71, 0x45, 0x9d, 0xb0, 0x29, 0xd9, 0x33, 0x07, 0x90,
	0xcb, 0xb0, 0x2b, 0x26, 0x54, 0x15, 0xd2, 0xae, 0xa4, 0x13, 0xee, 0x88, 0xc9, 0x51, 0x21, 0x07,
	0xbf, 0x07, 0x70, 0x79, 0x7d, 0x0d, 0xf6, 0x49, 0x7e, 0x0a, 0xe0, 0xc2, 0x9a, 0xa5, 0x5f, 0xbf,
	0x59, 0xbb, 0xdd, 0x3a, 0x64, 0xe7, 0xbf, 0x9a, 0x35, 0x05, 0xc3, 0x0e, 0xa2, 0xa8, 0xf3, 0x43,
	0xc4, 0x06, 0xbf, 0xed, 0x78, 0xa5, 0xab, 0x04, 0x8b, 0xb9, 0x61, 0x22, 0x21, 0xb7, 0xa1, 0x97,
	0x71, 0xae, 0xa9, 0x64, 0x11, 0x5d, 0x96, 0xbc, 0x19, 0x76, 0x11, 0x7f, 0xc2, 0xa2, 0x72, 0x3f,
	0x3f, 0x85, 0x5e, 0x96, 0x6a, 0x43, 0x63, 0x9e, 0x47, 0x5a, 0x64, 0x46, 0xa4, 0xca, 0x6f, 0xfb,
	0x05, 0xc4, 0x1f, 0x9c, 0xc2, 0xe4, 0x06, 0xb4, 0xf2, 0x45, 0x6e, 0xb8, 0x74, 0x19, 0xe9, 0x76,
	0x1e, 0x1c, 0x64, 0xd3, 0xf1, 0x0e, 0x10, 0x4f, 0x58, 0xf6, 0x56, 0xb7, 0xbc, 0x8b, 0xce, 0xb2,
	0xec, 0xef, 0x16, 0x74, 0x8d, 0x90, 0x9c, 0x6a, 0x2e, 0x99, 0x50, 0x42, 0x4d, 0xfb, 0x3b, 0x76,
	0x0b, 0x3a, 0x88, 0x86, 0x25, 0x48, 0x3e, 0x87, 0x7d, 0xef, 0x35, 0x62, 0x19, 0x1b, 0x8b, 0x44,
	0x18, 0xc1, 0xf3, 0xfe, 0xae, 0x75, 0xeb, 0x03, 0x7e, 0xbb, 0x64, 0x21, 0x5f, 0xc0, 0x25, 0xae,
	0xd8, 0x38, 0xe1, 0xf1, 0xea, 0x88, 0x3d, 0x3b, 0x62, 0xdf, 0xdb, 0x56, 0x86, 0xfc, 0x1c, 0xc0,
	0x45, 0xc5, 0xcd, 0xab, 0x54, 0x9f, 0x94, 0x7a, 0xf1, 0xbc, 0xdf, 0xb0, 0x49, 0x3a, 0xde, 0xea,
	0xb6, 0xda, 0x67, 0xd8, 0xf3, 0xc1, 0xef, 0x97, 0xb1, 0x71, 0x5f, 0x58, 0x61, 0x52, 0xaa, 0xf8,
	0x34, 0x35, 0x82, 0x59, 0x25, 0x9b, 0x6e, 0x5f, 0x10, 0x3f, 0x3a, 0x85, 0xc9, 0x37, 0x70, 0x90,
	0xcd, 0x16, 0xb9, 0x88, 0x58, 0x42, 0x25, 0x8f, 0x05, 0x5b, 0x5d, 0x36, 0xd8, 0x51, 0x57, 0x4b,
	0xca, 0x13, 0x64, 0xac, 0x2c, 0xfe, 0x6b, 0xb8, 0xe6, 0x86, 0x31, 0x63, 0x58, 0x34, 0x93, 0x5c,
	0x19, 0x5a, 0x28, 0x61, 0x5c, 0x9b, 0x68, 0xd9, 0x3d, 0x79, 0xdf, 0x32, 0xee, 0x57, 0x84, 0xe7,
	0x4a, 0x18, 0xdb, 0x32, 0x6e, 0x42, 0xdb, 0xe6, 0xcf, 0x3c, 0x61, 0x0a, 0xbb, 0x50, 0xdb, 0xd2,
	0x01, 0xb1, 0x17, 0x09, 0x53, 0xa3, 0x78, 0xf0, 0x1c, 0x0e, 0xec, 0xa2, 0x0b, 0x75, 0xa2, 0xd2,
	0x57, 0x8a, 0x9a, 0x64, 0xbe, 0xdc, 0x14, 0xae, 0x42, 0x03, 0x91, 0xaa, 0x25, 0x75, 0xc2, 0x3d,
	0x93, 0xcc, 0xad, 0xef, 0x03, 0x68, 0xa2, 0x69, 0xce, 0x92, 0xc2, 0x55, 0x6e, 0x3b, 0x44, 0xee,
	0x0b, 0xfc, 0x3f, 0xf8, 0x33, 0x80, 0xeb, 0x67, 0xf8, 0x75, 0x85, 0xfa, 0x6b, 0x00, 0x57, 0x36,
	0x13, 0xfa, 0x35, 0x5b, 0xaf, 0xb3, 0xf3, 0xdf, 0xd8, 0xcd, 0x6b, 0x0c, 0xf7, 0xd1, 0xf8, 0xdc,
	0xd9, 0x8e, 0x93, 0xb9, 0x2b, 0xde, 0x37, 0x81, 0x17, 0x26, 0xd5, 0x53, 0x1a, 0xf3, 0xc9, 0xba,
	0x30, 0x3d, 0xa8, 0xa5, 0x85, 0xf0, 0x9a, 0xe0, 0x2b, 0x16, 0x20, 0x72, 0x56, 0x7b, 0x19, 0x98,
	0x64, 0x5e, 0x36, 0xb3, 0x8f, 0xa1, 0x8b, 0x04, 0xa1, 0x26, 0x29, 0x15, 0x2a, 0xe6, 0xb9, 0x6f,
	0x6a, 0x6d, 0x93, 0xcc, 0x47, 0x6a, 0x92, 0x8e, 0x10, 0x5b, 0x95, 0xb5, 0xbe, 0x26, 0xeb, 0x1f,
	0xa5, 0xac, 0xef, 0xce, 0xca, 0xc9, 0xfa, 0xb6, 0x94, 0xf5, 0x1d, 0x42, 0x7f, 0x67, 0x5b, 0xb2,
	0x6e, 0x56, 0x28, 0x24, 0x68, 0xfc, 0x4e, 0x4f, 0x1f, 0xf0, 0x49, 0xa5, 0xea, 0x3f, 0x75, 0xb8,
	0xb8, 0xda, 0x12, 0xa5, 0x18, 0x93, 0x01, 0x74, 0x34, 0x97, 0xd4, 0xb6, 0x1b, 0xc9, 0xf4, 0x89,
	0x57, 0xb5, 0xa5, 0xb9, 0x3c, 0x16, 0x92, 0x3f, 0x61, 0xfa, 0x84, 0x7c, 0x06, 0x04, 0x39, 0x49,
	0x8a, 0x75, 0x64, 0x73, 0x1a, 0x4f, 0x05, 0x27, 0xf2, 0x05, 0xcd, 0xe5, 0x63, 0x34, 0x3c, 0x4d,
	0xb5, 0x39, 0x2a, 0x24, 0x6a, 0x88, 0x64, 0x14, 0xf9, 0xb5, 0x17, 0xb9, 0xa1, 0xb9, 0x44, 0x81,
	0x5f, 0x93, 0x3b, 0xb0, 0x1f, 0xcd, 0x58, 0x9e, 0x8b, 0x9c, 0x8a, 0x18, 0xb7, 0xcb, 0x65, 0x77,
	0xdd, 0xd2, 0x7a, 0xde, 0x34, 0x8a, 0x9f, 0x15, 0xe3, 0x63, 0xbf, 0x6b, 0x4b, 0xf4, 0x84, 0x2b,
	0xdf, 0x07, 0xdb, 0x15, 0xf3, 0x31, 0x57, 0xe4, 0x13, 0xdf, 0xa8, 0x97, 0x3d, 0xee, 0xba, 0x7e,
	0x89, 0xf8, 0xa9, 0xbb, 0x0f, 0xa1, 0x55, 0x12, 0xd1, 0xd7, 0x9e, 0x3b, 0xf1, 0x1c, 0x07, 0x1d,
	0xdd, 0x85, 0xcb, 0x51, 0x2a, 0xc7, 0x42, 0xad, 0xf7, 0xc7, 0x86, 0x65, 0x5e, 0x2a, 0x8d, 0x2b,
	0x3d, 0xe2, 0x4d, 0x00, 0xbd, 0xe5, 0xe4, 0x4e, 0x44, 0x6e, 0x6c, 0x3f, 0x6a, 0x1d, 0x8a, 0xff,
	0xa3, 0x8c, 0x5c, 0x9b, 0xec, 0x16, 0x55, 0x0d, 0x3d, 0x16, 0xb9, 0x21, 0xbf, 0x04, 0xd0, 0x5b,
	0xce, 0x0d, 0x3b, 0x2b, 0xd8, 0xd6, 0xac, 0xce, 0xa8, 0x88, 0xb0, 0x93, 0x96, 0x29, 0x88, 0x93,
	0x1a, 0xfc, 0x5d, 0x07, 0xb2, 0x9a, 0x81, 0xb6, 0x9c, 0xbf, 0x82, 0xbe, 0xe6, 0x11, 0x17, 0x73,
	0xa1, 0xa6, 0x74, 0xed, 0x72, 0xe7, 0x8e, 0xe6, 0x2b, 0x95, 0x7d, 0xb4, 0x72, 0xcb, 0x7b, 0x08,
	0x37, 0x4e, 0x47, 0x66, 0x4c, 0x63, 0x7f, 0xde, 0x78, 0x3b, 0xbc, 0x5e, 0xd1, 0x9e, 0x5a, 0xd6,
	0xe8, 0x3f, 0x5f, 0x16, 0xf1, 0x96, 0x74, 0x9a, 0x83, 0xfe, 0xc8, 0x6e, 0x56, 0xf9, 0x87, 0x29,
	0x5a, 0xe6, 0x94, 0xbb, 0x61, 0xd8, 0x14, 0x6d, 0x86, 0x6d, 0x97, 0x56, 0x0f, 0xdc, 0xad, 0xe3,
	0x16, 0x74, 0x67, 0x9c, 0xc5, 0x5c, 0xd3, 0x39, 0xd7, 0x39, 0x9e, 0x58, 0x3e, 0x41, 0x1d, 0xfa,
	0xc2, 0x81, 0x38, 0x91, 0x59, 0x9a, 0xc4, 0xb6, 0x1a, 0x7d, 0x7a, 0x36, 0x10, 0xc0, 0x4a, 0x3c,
	0xf3, 0xf0, 0x6e, 0x9c, 0x7d, 0x78, 0x5f, 0x83, 0x46, 0x96, 0x30, 0x33, 0x49, 0xb5, 0xf4, 0x47,
	0x64, 0xf5, 0x9f, 0x2c, 0x60, 0xd7, 0x4f, 0xd8, 0xa5, 0xc5, 0x16, 0xee, 0x68, 0x6b, 0x77, 0xaf,
	0xd0, 0x07, 0x24, 0x19, 0xd4, 0xa4, 0x18, 0xdb, 0xf3, 0xb3, 0x75, 0xf8, 0xfd, 0x16, 0xe3, 0x4a,
	0x31, 0x0e, 0x31, 0xd4, 0xe0, 0x6d, 0x00, 0xfb, 0x1b, 0xbe, 0x37, 0xc8, 0x8f, 0x01, 0x74, 0x56,
	0xf0, 0xfe, 0xa1, 0xed, 0xd4, 0x74, 0x8b, 0x93, 0xb2, 0x0d, 0xda, 0x7e, 0x80, 0x1c, 0x79, 0x68,
	0xbc, 0x6b, 0xbf, 0xb0, 0xee, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x92, 0x39, 0x66, 0x39, 0x7e,
	0x0d, 0x00, 0x00,
}
