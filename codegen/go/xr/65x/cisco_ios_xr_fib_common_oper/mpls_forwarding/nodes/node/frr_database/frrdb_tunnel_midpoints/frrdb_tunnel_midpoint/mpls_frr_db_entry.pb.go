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
// source: mpls_frr_db_entry.proto

package cisco_ios_xr_fib_common_oper_mpls_forwarding_nodes_node_frr_database_frrdb_tunnel_midpoints_frrdb_tunnel_midpoint

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

type MplsFrrDbEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	LocalLabel           uint32   `protobuf:"varint,2,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsFrrDbEntry_KEYS) Reset()         { *m = MplsFrrDbEntry_KEYS{} }
func (m *MplsFrrDbEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsFrrDbEntry_KEYS) ProtoMessage()    {}
func (*MplsFrrDbEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{0}
}

func (m *MplsFrrDbEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsFrrDbEntry_KEYS.Unmarshal(m, b)
}
func (m *MplsFrrDbEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsFrrDbEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsFrrDbEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsFrrDbEntry_KEYS.Merge(m, src)
}
func (m *MplsFrrDbEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsFrrDbEntry_KEYS.Size(m)
}
func (m *MplsFrrDbEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsFrrDbEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsFrrDbEntry_KEYS proto.InternalMessageInfo

func (m *MplsFrrDbEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsFrrDbEntry_KEYS) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

type FibMplsFrrDbEntryHeadId struct {
	DestinationPrefix       string   `protobuf:"bytes,1,opt,name=destination_prefix,json=destinationPrefix,proto3" json:"destination_prefix,omitempty"`
	DestinationPrefixLength uint32   `protobuf:"varint,2,opt,name=destination_prefix_length,json=destinationPrefixLength,proto3" json:"destination_prefix_length,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *FibMplsFrrDbEntryHeadId) Reset()         { *m = FibMplsFrrDbEntryHeadId{} }
func (m *FibMplsFrrDbEntryHeadId) String() string { return proto.CompactTextString(m) }
func (*FibMplsFrrDbEntryHeadId) ProtoMessage()    {}
func (*FibMplsFrrDbEntryHeadId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{1}
}

func (m *FibMplsFrrDbEntryHeadId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsFrrDbEntryHeadId.Unmarshal(m, b)
}
func (m *FibMplsFrrDbEntryHeadId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsFrrDbEntryHeadId.Marshal(b, m, deterministic)
}
func (m *FibMplsFrrDbEntryHeadId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsFrrDbEntryHeadId.Merge(m, src)
}
func (m *FibMplsFrrDbEntryHeadId) XXX_Size() int {
	return xxx_messageInfo_FibMplsFrrDbEntryHeadId.Size(m)
}
func (m *FibMplsFrrDbEntryHeadId) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsFrrDbEntryHeadId.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsFrrDbEntryHeadId proto.InternalMessageInfo

func (m *FibMplsFrrDbEntryHeadId) GetDestinationPrefix() string {
	if m != nil {
		return m.DestinationPrefix
	}
	return ""
}

func (m *FibMplsFrrDbEntryHeadId) GetDestinationPrefixLength() uint32 {
	if m != nil {
		return m.DestinationPrefixLength
	}
	return 0
}

type FibMplsFrrDbEntryMidId struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Lspid                uint32   `protobuf:"varint,2,opt,name=lspid,proto3" json:"lspid,omitempty"`
	TunnelId             uint32   `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibMplsFrrDbEntryMidId) Reset()         { *m = FibMplsFrrDbEntryMidId{} }
func (m *FibMplsFrrDbEntryMidId) String() string { return proto.CompactTextString(m) }
func (*FibMplsFrrDbEntryMidId) ProtoMessage()    {}
func (*FibMplsFrrDbEntryMidId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{2}
}

func (m *FibMplsFrrDbEntryMidId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsFrrDbEntryMidId.Unmarshal(m, b)
}
func (m *FibMplsFrrDbEntryMidId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsFrrDbEntryMidId.Marshal(b, m, deterministic)
}
func (m *FibMplsFrrDbEntryMidId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsFrrDbEntryMidId.Merge(m, src)
}
func (m *FibMplsFrrDbEntryMidId) XXX_Size() int {
	return xxx_messageInfo_FibMplsFrrDbEntryMidId.Size(m)
}
func (m *FibMplsFrrDbEntryMidId) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsFrrDbEntryMidId.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsFrrDbEntryMidId proto.InternalMessageInfo

func (m *FibMplsFrrDbEntryMidId) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *FibMplsFrrDbEntryMidId) GetLspid() uint32 {
	if m != nil {
		return m.Lspid
	}
	return 0
}

func (m *FibMplsFrrDbEntryMidId) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

type FibMplsFrrDbEntryGenId struct {
	Role                 string                   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Head                 *FibMplsFrrDbEntryHeadId `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	Midpoint             *FibMplsFrrDbEntryMidId  `protobuf:"bytes,3,opt,name=midpoint,proto3" json:"midpoint,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FibMplsFrrDbEntryGenId) Reset()         { *m = FibMplsFrrDbEntryGenId{} }
func (m *FibMplsFrrDbEntryGenId) String() string { return proto.CompactTextString(m) }
func (*FibMplsFrrDbEntryGenId) ProtoMessage()    {}
func (*FibMplsFrrDbEntryGenId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{3}
}

func (m *FibMplsFrrDbEntryGenId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsFrrDbEntryGenId.Unmarshal(m, b)
}
func (m *FibMplsFrrDbEntryGenId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsFrrDbEntryGenId.Marshal(b, m, deterministic)
}
func (m *FibMplsFrrDbEntryGenId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsFrrDbEntryGenId.Merge(m, src)
}
func (m *FibMplsFrrDbEntryGenId) XXX_Size() int {
	return xxx_messageInfo_FibMplsFrrDbEntryGenId.Size(m)
}
func (m *FibMplsFrrDbEntryGenId) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsFrrDbEntryGenId.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsFrrDbEntryGenId proto.InternalMessageInfo

func (m *FibMplsFrrDbEntryGenId) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *FibMplsFrrDbEntryGenId) GetHead() *FibMplsFrrDbEntryHeadId {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *FibMplsFrrDbEntryGenId) GetMidpoint() *FibMplsFrrDbEntryMidId {
	if m != nil {
		return m.Midpoint
	}
	return nil
}

type FibMplsFrrDbMcastLegs struct {
	FrrEntryId            *FibMplsFrrDbEntryGenId `protobuf:"bytes,1,opt,name=frr_entry_id,json=frrEntryId,proto3" json:"frr_entry_id,omitempty"`
	TunnelInterfaceName   string                  `protobuf:"bytes,2,opt,name=tunnel_interface_name,json=tunnelInterfaceName,proto3" json:"tunnel_interface_name,omitempty"`
	InputLabel            uint32                  `protobuf:"varint,3,opt,name=input_label,json=inputLabel,proto3" json:"input_label,omitempty"`
	OutgoingInterface     string                  `protobuf:"bytes,4,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	OutgoingLabel         uint32                  `protobuf:"varint,5,opt,name=outgoing_label,json=outgoingLabel,proto3" json:"outgoing_label,omitempty"`
	FrrInterfaceName      string                  `protobuf:"bytes,6,opt,name=frr_interface_name,json=frrInterfaceName,proto3" json:"frr_interface_name,omitempty"`
	FrrLabel              uint32                  `protobuf:"varint,7,opt,name=frr_label,json=frrLabel,proto3" json:"frr_label,omitempty"`
	EntryFrrState         string                  `protobuf:"bytes,8,opt,name=entry_frr_state,json=entryFrrState,proto3" json:"entry_frr_state,omitempty"`
	FrrNextHopIpv4Address string                  `protobuf:"bytes,9,opt,name=frr_next_hop_ipv4_address,json=frrNextHopIpv4Address,proto3" json:"frr_next_hop_ipv4_address,omitempty"`
	IsMldpLsp             bool                    `protobuf:"varint,10,opt,name=is_mldp_lsp,json=isMldpLsp,proto3" json:"is_mldp_lsp,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                `json:"-"`
	XXX_unrecognized      []byte                  `json:"-"`
	XXX_sizecache         int32                   `json:"-"`
}

func (m *FibMplsFrrDbMcastLegs) Reset()         { *m = FibMplsFrrDbMcastLegs{} }
func (m *FibMplsFrrDbMcastLegs) String() string { return proto.CompactTextString(m) }
func (*FibMplsFrrDbMcastLegs) ProtoMessage()    {}
func (*FibMplsFrrDbMcastLegs) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{4}
}

func (m *FibMplsFrrDbMcastLegs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsFrrDbMcastLegs.Unmarshal(m, b)
}
func (m *FibMplsFrrDbMcastLegs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsFrrDbMcastLegs.Marshal(b, m, deterministic)
}
func (m *FibMplsFrrDbMcastLegs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsFrrDbMcastLegs.Merge(m, src)
}
func (m *FibMplsFrrDbMcastLegs) XXX_Size() int {
	return xxx_messageInfo_FibMplsFrrDbMcastLegs.Size(m)
}
func (m *FibMplsFrrDbMcastLegs) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsFrrDbMcastLegs.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsFrrDbMcastLegs proto.InternalMessageInfo

func (m *FibMplsFrrDbMcastLegs) GetFrrEntryId() *FibMplsFrrDbEntryGenId {
	if m != nil {
		return m.FrrEntryId
	}
	return nil
}

func (m *FibMplsFrrDbMcastLegs) GetTunnelInterfaceName() string {
	if m != nil {
		return m.TunnelInterfaceName
	}
	return ""
}

func (m *FibMplsFrrDbMcastLegs) GetInputLabel() uint32 {
	if m != nil {
		return m.InputLabel
	}
	return 0
}

func (m *FibMplsFrrDbMcastLegs) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *FibMplsFrrDbMcastLegs) GetOutgoingLabel() uint32 {
	if m != nil {
		return m.OutgoingLabel
	}
	return 0
}

func (m *FibMplsFrrDbMcastLegs) GetFrrInterfaceName() string {
	if m != nil {
		return m.FrrInterfaceName
	}
	return ""
}

func (m *FibMplsFrrDbMcastLegs) GetFrrLabel() uint32 {
	if m != nil {
		return m.FrrLabel
	}
	return 0
}

func (m *FibMplsFrrDbMcastLegs) GetEntryFrrState() string {
	if m != nil {
		return m.EntryFrrState
	}
	return ""
}

func (m *FibMplsFrrDbMcastLegs) GetFrrNextHopIpv4Address() string {
	if m != nil {
		return m.FrrNextHopIpv4Address
	}
	return ""
}

func (m *FibMplsFrrDbMcastLegs) GetIsMldpLsp() bool {
	if m != nil {
		return m.IsMldpLsp
	}
	return false
}

type FibMplsFrrDbEntry struct {
	FrrEntryId            *FibMplsFrrDbEntryGenId  `protobuf:"bytes,1,opt,name=frr_entry_id,json=frrEntryId,proto3" json:"frr_entry_id,omitempty"`
	TunnelInterfaceName   string                   `protobuf:"bytes,2,opt,name=tunnel_interface_name,json=tunnelInterfaceName,proto3" json:"tunnel_interface_name,omitempty"`
	InputLabel            uint32                   `protobuf:"varint,3,opt,name=input_label,json=inputLabel,proto3" json:"input_label,omitempty"`
	OutgoingInterface     string                   `protobuf:"bytes,4,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	OutgoingLabel         uint32                   `protobuf:"varint,5,opt,name=outgoing_label,json=outgoingLabel,proto3" json:"outgoing_label,omitempty"`
	FrrInterfaceName      string                   `protobuf:"bytes,6,opt,name=frr_interface_name,json=frrInterfaceName,proto3" json:"frr_interface_name,omitempty"`
	FrrLabel              uint32                   `protobuf:"varint,7,opt,name=frr_label,json=frrLabel,proto3" json:"frr_label,omitempty"`
	EntryFrrState         string                   `protobuf:"bytes,8,opt,name=entry_frr_state,json=entryFrrState,proto3" json:"entry_frr_state,omitempty"`
	FrrNextHopIpv4Address string                   `protobuf:"bytes,9,opt,name=frr_next_hop_ipv4_address,json=frrNextHopIpv4Address,proto3" json:"frr_next_hop_ipv4_address,omitempty"`
	IsMldpLsp             bool                     `protobuf:"varint,10,opt,name=is_mldp_lsp,json=isMldpLsp,proto3" json:"is_mldp_lsp,omitempty"`
	IsMulticastTunnel     bool                     `protobuf:"varint,11,opt,name=is_multicast_tunnel,json=isMulticastTunnel,proto3" json:"is_multicast_tunnel,omitempty"`
	MulticastTunnelLegs   uint32                   `protobuf:"varint,12,opt,name=multicast_tunnel_legs,json=multicastTunnelLegs,proto3" json:"multicast_tunnel_legs,omitempty"`
	MulticastLeg          []*FibMplsFrrDbMcastLegs `protobuf:"bytes,13,rep,name=multicast_leg,json=multicastLeg,proto3" json:"multicast_leg,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *FibMplsFrrDbEntry) Reset()         { *m = FibMplsFrrDbEntry{} }
func (m *FibMplsFrrDbEntry) String() string { return proto.CompactTextString(m) }
func (*FibMplsFrrDbEntry) ProtoMessage()    {}
func (*FibMplsFrrDbEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{5}
}

func (m *FibMplsFrrDbEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibMplsFrrDbEntry.Unmarshal(m, b)
}
func (m *FibMplsFrrDbEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibMplsFrrDbEntry.Marshal(b, m, deterministic)
}
func (m *FibMplsFrrDbEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibMplsFrrDbEntry.Merge(m, src)
}
func (m *FibMplsFrrDbEntry) XXX_Size() int {
	return xxx_messageInfo_FibMplsFrrDbEntry.Size(m)
}
func (m *FibMplsFrrDbEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_FibMplsFrrDbEntry.DiscardUnknown(m)
}

var xxx_messageInfo_FibMplsFrrDbEntry proto.InternalMessageInfo

func (m *FibMplsFrrDbEntry) GetFrrEntryId() *FibMplsFrrDbEntryGenId {
	if m != nil {
		return m.FrrEntryId
	}
	return nil
}

func (m *FibMplsFrrDbEntry) GetTunnelInterfaceName() string {
	if m != nil {
		return m.TunnelInterfaceName
	}
	return ""
}

func (m *FibMplsFrrDbEntry) GetInputLabel() uint32 {
	if m != nil {
		return m.InputLabel
	}
	return 0
}

func (m *FibMplsFrrDbEntry) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *FibMplsFrrDbEntry) GetOutgoingLabel() uint32 {
	if m != nil {
		return m.OutgoingLabel
	}
	return 0
}

func (m *FibMplsFrrDbEntry) GetFrrInterfaceName() string {
	if m != nil {
		return m.FrrInterfaceName
	}
	return ""
}

func (m *FibMplsFrrDbEntry) GetFrrLabel() uint32 {
	if m != nil {
		return m.FrrLabel
	}
	return 0
}

func (m *FibMplsFrrDbEntry) GetEntryFrrState() string {
	if m != nil {
		return m.EntryFrrState
	}
	return ""
}

func (m *FibMplsFrrDbEntry) GetFrrNextHopIpv4Address() string {
	if m != nil {
		return m.FrrNextHopIpv4Address
	}
	return ""
}

func (m *FibMplsFrrDbEntry) GetIsMldpLsp() bool {
	if m != nil {
		return m.IsMldpLsp
	}
	return false
}

func (m *FibMplsFrrDbEntry) GetIsMulticastTunnel() bool {
	if m != nil {
		return m.IsMulticastTunnel
	}
	return false
}

func (m *FibMplsFrrDbEntry) GetMulticastTunnelLegs() uint32 {
	if m != nil {
		return m.MulticastTunnelLegs
	}
	return 0
}

func (m *FibMplsFrrDbEntry) GetMulticastLeg() []*FibMplsFrrDbMcastLegs {
	if m != nil {
		return m.MulticastLeg
	}
	return nil
}

type MplsFrrDbEntry struct {
	FrrDb                *FibMplsFrrDbEntry `protobuf:"bytes,50,opt,name=frr_db,json=frrDb,proto3" json:"frr_db,omitempty"`
	OutgoingLableString  string             `protobuf:"bytes,51,opt,name=outgoing_lable_string,json=outgoingLableString,proto3" json:"outgoing_lable_string,omitempty"`
	FrrLableString       string             `protobuf:"bytes,52,opt,name=frr_lable_string,json=frrLableString,proto3" json:"frr_lable_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MplsFrrDbEntry) Reset()         { *m = MplsFrrDbEntry{} }
func (m *MplsFrrDbEntry) String() string { return proto.CompactTextString(m) }
func (*MplsFrrDbEntry) ProtoMessage()    {}
func (*MplsFrrDbEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a68f3dcef7d5ac2, []int{6}
}

func (m *MplsFrrDbEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsFrrDbEntry.Unmarshal(m, b)
}
func (m *MplsFrrDbEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsFrrDbEntry.Marshal(b, m, deterministic)
}
func (m *MplsFrrDbEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsFrrDbEntry.Merge(m, src)
}
func (m *MplsFrrDbEntry) XXX_Size() int {
	return xxx_messageInfo_MplsFrrDbEntry.Size(m)
}
func (m *MplsFrrDbEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsFrrDbEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MplsFrrDbEntry proto.InternalMessageInfo

func (m *MplsFrrDbEntry) GetFrrDb() *FibMplsFrrDbEntry {
	if m != nil {
		return m.FrrDb
	}
	return nil
}

func (m *MplsFrrDbEntry) GetOutgoingLableString() string {
	if m != nil {
		return m.OutgoingLableString
	}
	return ""
}

func (m *MplsFrrDbEntry) GetFrrLableString() string {
	if m != nil {
		return m.FrrLableString
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsFrrDbEntry_KEYS)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.mpls_frr_db_entry_KEYS")
	proto.RegisterType((*FibMplsFrrDbEntryHeadId)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.fib_mpls_frr_db_entry_head_id")
	proto.RegisterType((*FibMplsFrrDbEntryMidId)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.fib_mpls_frr_db_entry_mid_id")
	proto.RegisterType((*FibMplsFrrDbEntryGenId)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.fib_mpls_frr_db_entry_gen_id")
	proto.RegisterType((*FibMplsFrrDbMcastLegs)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.fib_mpls_frr_db_mcast_legs")
	proto.RegisterType((*FibMplsFrrDbEntry)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.fib_mpls_frr_db_entry")
	proto.RegisterType((*MplsFrrDbEntry)(nil), "cisco_ios_xr_fib_common_oper.mpls_forwarding.nodes.node.frr_database.frrdb_tunnel_midpoints.frrdb_tunnel_midpoint.mpls_frr_db_entry")
}

func init() { proto.RegisterFile("mpls_frr_db_entry.proto", fileDescriptor_2a68f3dcef7d5ac2) }

var fileDescriptor_2a68f3dcef7d5ac2 = []byte{
	// 744 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0xd6, 0xb4, 0x69, 0x6e, 0x72, 0xd2, 0xf4, 0xb6, 0xee, 0xed, 0xed, 0x94, 0xf2, 0x13, 0x45,
	0x02, 0x65, 0x01, 0x59, 0xa4, 0x5d, 0x20, 0x76, 0x48, 0x14, 0x51, 0x91, 0x56, 0x68, 0x8a, 0x90,
	0x58, 0x59, 0x93, 0xd8, 0x33, 0xb5, 0xe4, 0xb1, 0x07, 0xdb, 0x29, 0xe1, 0x0d, 0x00, 0x09, 0xc1,
	0x82, 0x07, 0x60, 0xc3, 0x8e, 0x15, 0x3c, 0x05, 0x2f, 0xc1, 0xb3, 0x20, 0xdb, 0x33, 0x49, 0xdb,
	0x14, 0xc4, 0x0a, 0x65, 0xd1, 0x4d, 0xd5, 0xf9, 0xce, 0xe7, 0x73, 0x8e, 0xcf, 0xf9, 0x3e, 0x2b,
	0xb0, 0x99, 0xe5, 0x5c, 0xe3, 0x44, 0x29, 0x4c, 0x06, 0x98, 0x0a, 0xa3, 0x5e, 0x75, 0x73, 0x25,
	0x8d, 0x44, 0x2f, 0x86, 0x4c, 0x0f, 0x25, 0x66, 0x52, 0xe3, 0xb1, 0xc2, 0x09, 0x1b, 0xe0, 0xa1,
	0xcc, 0x32, 0x29, 0xb0, 0xcc, 0xa9, 0xea, 0xfa, 0x53, 0x52, 0xbd, 0x8c, 0x15, 0x61, 0x22, 0xed,
	0x0a, 0x49, 0xa8, 0x76, 0x7f, 0xbb, 0x2e, 0x57, 0x6c, 0xe2, 0x41, 0xac, 0xdd, 0x07, 0x19, 0x60,
	0x33, 0x12, 0x82, 0x72, 0x9c, 0x31, 0x92, 0x4b, 0x26, 0x8c, 0xbe, 0x18, 0x6e, 0x3f, 0x83, 0xff,
	0x67, 0xba, 0xc1, 0x8f, 0xf7, 0x9e, 0x1f, 0xa1, 0x6d, 0xa8, 0xdb, 0xdc, 0x58, 0xc4, 0x19, 0x0d,
	0x83, 0x56, 0xd0, 0xa9, 0x47, 0x35, 0x0b, 0x1c, 0xc6, 0x19, 0x45, 0x37, 0xa0, 0xc1, 0xe5, 0x30,
	0xe6, 0x98, 0xc7, 0x03, 0xca, 0xc3, 0x85, 0x56, 0xd0, 0x69, 0x46, 0xe0, 0xa0, 0xbe, 0x45, 0xda,
	0x6f, 0x03, 0xb8, 0x66, 0x2f, 0x30, 0x9b, 0xfc, 0x98, 0xc6, 0x04, 0x33, 0x82, 0xee, 0x00, 0x22,
	0x54, 0x1b, 0x26, 0x62, 0xc3, 0xa4, 0xc0, 0xb9, 0xa2, 0x09, 0x1b, 0x17, 0x85, 0xd6, 0x4e, 0x45,
	0x9e, 0xb8, 0x00, 0xba, 0x07, 0x5b, 0xb3, 0x74, 0xcc, 0xa9, 0x48, 0xcd, 0x71, 0x51, 0x7f, 0x73,
	0xe6, 0x54, 0xdf, 0x85, 0xdb, 0x63, 0xb8, 0x7a, 0x71, 0x2f, 0x19, 0x73, 0xad, 0xdc, 0x84, 0x15,
	0x2d, 0x47, 0x6a, 0x48, 0x71, 0x4c, 0x88, 0xa2, 0x5a, 0x17, 0x6d, 0x34, 0x3d, 0x7a, 0xdf, 0x83,
	0xe8, 0x3f, 0x58, 0xe2, 0x3a, 0x67, 0xa4, 0x28, 0xe7, 0x3f, 0xec, 0x9c, 0x8a, 0xa1, 0x32, 0x12,
	0x2e, 0xba, 0x48, 0xcd, 0x03, 0xfb, 0xa4, 0xfd, 0x66, 0xf1, 0x57, 0xa5, 0x53, 0x2a, 0x6c, 0x69,
	0x04, 0x15, 0x25, 0x79, 0x39, 0x60, 0xf7, 0x3f, 0xfa, 0x1c, 0x40, 0xc5, 0x4e, 0xc9, 0xd5, 0x69,
	0xf4, 0x3e, 0x04, 0xdd, 0xbf, 0xae, 0x8b, 0xee, 0x6f, 0x77, 0x17, 0xb9, 0xf6, 0xd0, 0x97, 0x00,
	0x6a, 0xe5, 0x01, 0x77, 0xf3, 0x46, 0xef, 0xfd, 0xfc, 0xf4, 0xea, 0x77, 0x1b, 0x4d, 0x3a, 0x6c,
	0x7f, 0xaf, 0xc0, 0x95, 0xf3, 0xd4, 0x6c, 0x18, 0x6b, 0x83, 0x39, 0x4d, 0x35, 0xfa, 0x1a, 0xc0,
	0xb2, 0x45, 0xfd, 0x69, 0x46, 0xdc, 0x4a, 0xe6, 0xea, 0x46, 0x5e, 0x32, 0x11, 0x24, 0x4a, 0xed,
	0x59, 0x60, 0x9f, 0xa0, 0x1e, 0x6c, 0x94, 0xe2, 0x13, 0x86, 0xaa, 0x24, 0x1e, 0x16, 0x86, 0x5d,
	0x70, 0x7a, 0x5a, 0x2f, 0x84, 0x58, 0xc6, 0x4a, 0xef, 0x32, 0x91, 0x8f, 0x4c, 0xe1, 0x5d, 0x2f,
	0x59, 0x70, 0x90, 0xf3, 0xae, 0x75, 0xa6, 0x1c, 0x99, 0x54, 0x32, 0x91, 0x4e, 0xd3, 0x86, 0x15,
	0xef, 0xcc, 0x32, 0x32, 0xc9, 0x69, 0xdd, 0x33, 0xa1, 0xfb, 0x94, 0x4b, 0x2e, 0x65, 0xb3, 0x44,
	0x7d, 0xd6, 0xdb, 0x80, 0xec, 0x6d, 0xce, 0xf5, 0x59, 0x75, 0x59, 0x57, 0x13, 0xa5, 0xce, 0x36,
	0xb9, 0x0d, 0x75, 0xcb, 0xf6, 0xf9, 0xfe, 0xf1, 0xae, 0x4a, 0x94, 0xf2, 0xa9, 0x6e, 0xc1, 0xbf,
	0x7e, 0x22, 0x96, 0xa2, 0x4d, 0x6c, 0x68, 0x58, 0xf3, 0x86, 0x75, 0xf0, 0x43, 0xa5, 0x8e, 0x2c,
	0x88, 0xee, 0xc2, 0x96, 0x65, 0x08, 0x3a, 0x36, 0xf8, 0x58, 0xe6, 0x98, 0xe5, 0x27, 0xbb, 0x13,
	0x8b, 0xd7, 0xdd, 0x89, 0x8d, 0x44, 0xa9, 0x43, 0x3a, 0x36, 0x8f, 0x64, 0xbe, 0x9f, 0x9f, 0xec,
	0x96, 0x56, 0xbf, 0x0e, 0x0d, 0xa6, 0x71, 0xc6, 0x49, 0x8e, 0xb9, 0xce, 0x43, 0x68, 0x05, 0x9d,
	0x5a, 0x54, 0x67, 0xfa, 0x80, 0x93, 0xbc, 0xaf, 0xf3, 0xf6, 0x8f, 0x2a, 0x6c, 0x5c, 0xb8, 0xa4,
	0x4b, 0x19, 0x5d, 0xca, 0xe8, 0x8f, 0x65, 0x84, 0xba, 0xb0, 0x6e, 0xe3, 0x23, 0x6e, 0x98, 0x7b,
	0x88, 0xfc, 0x9c, 0xc3, 0x86, 0xe3, 0xad, 0x31, 0x7d, 0x50, 0x46, 0x9e, 0xba, 0x80, 0xdd, 0xd3,
	0x79, 0xb2, 0x7b, 0xbc, 0xc2, 0x65, 0x77, 0xb5, 0xf5, 0xec, 0x2c, 0xbf, 0x6f, 0xdf, 0xb5, 0x6f,
	0x01, 0x34, 0xa7, 0x87, 0x38, 0x4d, 0xc3, 0x66, 0x6b, 0xb1, 0xd3, 0xe8, 0xbd, 0x9b, 0x07, 0x45,
	0x4e, 0xdf, 0xdf, 0x68, 0x79, 0xd2, 0x64, 0x9f, 0xa6, 0xed, 0x8f, 0x0b, 0xb0, 0x36, 0x6b, 0xae,
	0x4f, 0x01, 0x54, 0x3d, 0x10, 0xf6, 0x9c, 0xad, 0x5e, 0xcf, 0x8d, 0xad, 0xa2, 0xa5, 0x44, 0xa9,
	0x07, 0x03, 0xbb, 0xa2, 0xd3, 0x32, 0xe6, 0x14, 0x6b, 0xa3, 0x98, 0x48, 0xc3, 0x1d, 0x6f, 0xa5,
	0x53, 0x6a, 0xe6, 0xf4, 0xc8, 0x85, 0x50, 0x07, 0x56, 0x0b, 0x95, 0x4e, 0xe9, 0xbb, 0x8e, 0xbe,
	0xe2, 0xc5, 0x5a, 0x32, 0x07, 0x55, 0xf7, 0x43, 0x71, 0xe7, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x17, 0xac, 0x69, 0xc2, 0x43, 0x0a, 0x00, 0x00,
}
