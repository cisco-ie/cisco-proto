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
// source: ldp_summary.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_summary

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

type LdpSummary_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpSummary_KEYS) Reset()         { *m = LdpSummary_KEYS{} }
func (m *LdpSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpSummary_KEYS) ProtoMessage()    {}
func (*LdpSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7f2abcb6a0f7f76, []int{0}
}

func (m *LdpSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummary_KEYS.Unmarshal(m, b)
}
func (m *LdpSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummary_KEYS.Merge(m, src)
}
func (m *LdpSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpSummary_KEYS.Size(m)
}
func (m *LdpSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummary_KEYS proto.InternalMessageInfo

func (m *LdpSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type LdpSummaryCommon struct {
	AddressFamilies                     string   `protobuf:"bytes,1,opt,name=address_families,json=addressFamilies,proto3" json:"address_families,omitempty"`
	NumberOfIpv4Af                      uint32   `protobuf:"varint,2,opt,name=number_of_ipv4af,json=numberOfIpv4af,proto3" json:"number_of_ipv4af,omitempty"`
	NumberOfIpv6Af                      uint32   `protobuf:"varint,3,opt,name=number_of_ipv6af,json=numberOfIpv6af,proto3" json:"number_of_ipv6af,omitempty"`
	NumberOfNeighbors                   uint32   `protobuf:"varint,4,opt,name=number_of_neighbors,json=numberOfNeighbors,proto3" json:"number_of_neighbors,omitempty"`
	NumberOfNsrSyncedNeighbors          uint32   `protobuf:"varint,5,opt,name=number_of_nsr_synced_neighbors,json=numberOfNsrSyncedNeighbors,proto3" json:"number_of_nsr_synced_neighbors,omitempty"`
	NumberOfGracefulRestartNeighbors    uint32   `protobuf:"varint,6,opt,name=number_of_graceful_restart_neighbors,json=numberOfGracefulRestartNeighbors,proto3" json:"number_of_graceful_restart_neighbors,omitempty"`
	NumberOfDownstreamOnDemandNeighbors uint32   `protobuf:"varint,7,opt,name=number_of_downstream_on_demand_neighbors,json=numberOfDownstreamOnDemandNeighbors,proto3" json:"number_of_downstream_on_demand_neighbors,omitempty"`
	NumberOfIpv4HelloAdj                uint32   `protobuf:"varint,8,opt,name=number_of_ipv4_hello_adj,json=numberOfIpv4HelloAdj,proto3" json:"number_of_ipv4_hello_adj,omitempty"`
	NumberOfIpv6HelloAdj                uint32   `protobuf:"varint,9,opt,name=number_of_ipv6_hello_adj,json=numberOfIpv6HelloAdj,proto3" json:"number_of_ipv6_hello_adj,omitempty"`
	NumberOfIpv4Routes                  uint32   `protobuf:"varint,10,opt,name=number_of_ipv4_routes,json=numberOfIpv4Routes,proto3" json:"number_of_ipv4_routes,omitempty"`
	NumberOfIpv6Routes                  uint32   `protobuf:"varint,11,opt,name=number_of_ipv6_routes,json=numberOfIpv6Routes,proto3" json:"number_of_ipv6_routes,omitempty"`
	NumberOfIpv4LocalAddresses          uint32   `protobuf:"varint,12,opt,name=number_of_ipv4_local_addresses,json=numberOfIpv4LocalAddresses,proto3" json:"number_of_ipv4_local_addresses,omitempty"`
	NumberOfIpv6LocalAddresses          uint32   `protobuf:"varint,13,opt,name=number_of_ipv6_local_addresses,json=numberOfIpv6LocalAddresses,proto3" json:"number_of_ipv6_local_addresses,omitempty"`
	NumberOfLdpInterfaces               uint32   `protobuf:"varint,14,opt,name=number_of_ldp_interfaces,json=numberOfLdpInterfaces,proto3" json:"number_of_ldp_interfaces,omitempty"`
	NumberOfIpv4LdpInterfaces           uint32   `protobuf:"varint,15,opt,name=number_of_ipv4ldp_interfaces,json=numberOfIpv4ldpInterfaces,proto3" json:"number_of_ipv4ldp_interfaces,omitempty"`
	NumberOfIpv6LdpInterfaces           uint32   `protobuf:"varint,16,opt,name=number_of_ipv6ldp_interfaces,json=numberOfIpv6ldpInterfaces,proto3" json:"number_of_ipv6ldp_interfaces,omitempty"`
	NumberOfBindingsIpv4                uint32   `protobuf:"varint,17,opt,name=number_of_bindings_ipv4,json=numberOfBindingsIpv4,proto3" json:"number_of_bindings_ipv4,omitempty"`
	NumberOfBindingsIpv6                uint32   `protobuf:"varint,18,opt,name=number_of_bindings_ipv6,json=numberOfBindingsIpv6,proto3" json:"number_of_bindings_ipv6,omitempty"`
	NumberOfLocalBindingsIpv4           uint32   `protobuf:"varint,19,opt,name=number_of_local_bindings_ipv4,json=numberOfLocalBindingsIpv4,proto3" json:"number_of_local_bindings_ipv4,omitempty"`
	NumberOfLocalBindingsIpv6           uint32   `protobuf:"varint,20,opt,name=number_of_local_bindings_ipv6,json=numberOfLocalBindingsIpv6,proto3" json:"number_of_local_bindings_ipv6,omitempty"`
	NumberOfRemoteBindingsIpv4          uint32   `protobuf:"varint,21,opt,name=number_of_remote_bindings_ipv4,json=numberOfRemoteBindingsIpv4,proto3" json:"number_of_remote_bindings_ipv4,omitempty"`
	NumberOfRemoteBindingsIpv6          uint32   `protobuf:"varint,22,opt,name=number_of_remote_bindings_ipv6,json=numberOfRemoteBindingsIpv6,proto3" json:"number_of_remote_bindings_ipv6,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *LdpSummaryCommon) Reset()         { *m = LdpSummaryCommon{} }
func (m *LdpSummaryCommon) String() string { return proto.CompactTextString(m) }
func (*LdpSummaryCommon) ProtoMessage()    {}
func (*LdpSummaryCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7f2abcb6a0f7f76, []int{1}
}

func (m *LdpSummaryCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummaryCommon.Unmarshal(m, b)
}
func (m *LdpSummaryCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummaryCommon.Marshal(b, m, deterministic)
}
func (m *LdpSummaryCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummaryCommon.Merge(m, src)
}
func (m *LdpSummaryCommon) XXX_Size() int {
	return xxx_messageInfo_LdpSummaryCommon.Size(m)
}
func (m *LdpSummaryCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummaryCommon.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummaryCommon proto.InternalMessageInfo

func (m *LdpSummaryCommon) GetAddressFamilies() string {
	if m != nil {
		return m.AddressFamilies
	}
	return ""
}

func (m *LdpSummaryCommon) GetNumberOfIpv4Af() uint32 {
	if m != nil {
		return m.NumberOfIpv4Af
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6Af() uint32 {
	if m != nil {
		return m.NumberOfIpv6Af
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfNeighbors() uint32 {
	if m != nil {
		return m.NumberOfNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfNsrSyncedNeighbors() uint32 {
	if m != nil {
		return m.NumberOfNsrSyncedNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfGracefulRestartNeighbors() uint32 {
	if m != nil {
		return m.NumberOfGracefulRestartNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfDownstreamOnDemandNeighbors() uint32 {
	if m != nil {
		return m.NumberOfDownstreamOnDemandNeighbors
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4HelloAdj() uint32 {
	if m != nil {
		return m.NumberOfIpv4HelloAdj
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6HelloAdj() uint32 {
	if m != nil {
		return m.NumberOfIpv6HelloAdj
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4Routes() uint32 {
	if m != nil {
		return m.NumberOfIpv4Routes
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6Routes() uint32 {
	if m != nil {
		return m.NumberOfIpv6Routes
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4LocalAddresses() uint32 {
	if m != nil {
		return m.NumberOfIpv4LocalAddresses
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6LocalAddresses() uint32 {
	if m != nil {
		return m.NumberOfIpv6LocalAddresses
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfLdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv4LdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfIpv4LdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfIpv6LdpInterfaces() uint32 {
	if m != nil {
		return m.NumberOfIpv6LdpInterfaces
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfBindingsIpv6
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLocalBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfLocalBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfLocalBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfLocalBindingsIpv6
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfRemoteBindingsIpv4() uint32 {
	if m != nil {
		return m.NumberOfRemoteBindingsIpv4
	}
	return 0
}

func (m *LdpSummaryCommon) GetNumberOfRemoteBindingsIpv6() uint32 {
	if m != nil {
		return m.NumberOfRemoteBindingsIpv6
	}
	return 0
}

type LdpSummary struct {
	NumberOfVrf                      uint32            `protobuf:"varint,50,opt,name=number_of_vrf,json=numberOfVrf,proto3" json:"number_of_vrf,omitempty"`
	NumberOfVrfOper                  uint32            `protobuf:"varint,51,opt,name=number_of_vrf_oper,json=numberOfVrfOper,proto3" json:"number_of_vrf_oper,omitempty"`
	Common                           *LdpSummaryCommon `protobuf:"bytes,52,opt,name=common,proto3" json:"common,omitempty"`
	NumberOfInterfaces               uint32            `protobuf:"varint,53,opt,name=number_of_interfaces,json=numberOfInterfaces,proto3" json:"number_of_interfaces,omitempty"`
	NumberOfFwdRefInterfaces         uint32            `protobuf:"varint,54,opt,name=number_of_fwd_ref_interfaces,json=numberOfFwdRefInterfaces,proto3" json:"number_of_fwd_ref_interfaces,omitempty"`
	NumberOfAutocfgInterfaces        uint32            `protobuf:"varint,55,opt,name=number_of_autocfg_interfaces,json=numberOfAutocfgInterfaces,proto3" json:"number_of_autocfg_interfaces,omitempty"`
	IsBoundWithSysdb                 bool              `protobuf:"varint,56,opt,name=is_bound_with_sysdb,json=isBoundWithSysdb,proto3" json:"is_bound_with_sysdb,omitempty"`
	IsRegisteredWithSysdb            bool              `protobuf:"varint,57,opt,name=is_registered_with_sysdb,json=isRegisteredWithSysdb,proto3" json:"is_registered_with_sysdb,omitempty"`
	IsBoundWithRsi                   bool              `protobuf:"varint,58,opt,name=is_bound_with_rsi,json=isBoundWithRsi,proto3" json:"is_bound_with_rsi,omitempty"`
	IsBoundWithInterfaceManager      bool              `protobuf:"varint,59,opt,name=is_bound_with_interface_manager,json=isBoundWithInterfaceManager,proto3" json:"is_bound_with_interface_manager,omitempty"`
	IsRegisteredWithInterfaceManager bool              `protobuf:"varint,60,opt,name=is_registered_with_interface_manager,json=isRegisteredWithInterfaceManager,proto3" json:"is_registered_with_interface_manager,omitempty"`
	IsBoundWithIpArm                 bool              `protobuf:"varint,61,opt,name=is_bound_with_ip_arm,json=isBoundWithIpArm,proto3" json:"is_bound_with_ip_arm,omitempty"`
	IsBoundWithLsd                   bool              `protobuf:"varint,62,opt,name=is_bound_with_lsd,json=isBoundWithLsd,proto3" json:"is_bound_with_lsd,omitempty"`
	IsRegisteredWithLsd              bool              `protobuf:"varint,63,opt,name=is_registered_with_lsd,json=isRegisteredWithLsd,proto3" json:"is_registered_with_lsd,omitempty"`
	IsBoundWithIpv4Rib               bool              `protobuf:"varint,64,opt,name=is_bound_with_ipv4_rib,json=isBoundWithIpv4Rib,proto3" json:"is_bound_with_ipv4_rib,omitempty"`
	IsRegisteredWithIpv4Rib          bool              `protobuf:"varint,65,opt,name=is_registered_with_ipv4_rib,json=isRegisteredWithIpv4Rib,proto3" json:"is_registered_with_ipv4_rib,omitempty"`
	NumberOfIpv4RibTables            uint32            `protobuf:"varint,66,opt,name=number_of_ipv4rib_tables,json=numberOfIpv4ribTables,proto3" json:"number_of_ipv4rib_tables,omitempty"`
	NumberOfRegisteredIpv4RibTables  uint32            `protobuf:"varint,67,opt,name=number_of_registered_ipv4rib_tables,json=numberOfRegisteredIpv4ribTables,proto3" json:"number_of_registered_ipv4rib_tables,omitempty"`
	IsBoundWithIpv6Rib               bool              `protobuf:"varint,68,opt,name=is_bound_with_ipv6_rib,json=isBoundWithIpv6Rib,proto3" json:"is_bound_with_ipv6_rib,omitempty"`
	IsRegisteredWithIpv6Rib          bool              `protobuf:"varint,69,opt,name=is_registered_with_ipv6_rib,json=isRegisteredWithIpv6Rib,proto3" json:"is_registered_with_ipv6_rib,omitempty"`
	NumberOfIpv6RibTables            uint32            `protobuf:"varint,70,opt,name=number_of_ipv6rib_tables,json=numberOfIpv6ribTables,proto3" json:"number_of_ipv6rib_tables,omitempty"`
	NumberOfRegisteredIpv6RibTables  uint32            `protobuf:"varint,71,opt,name=number_of_registered_ipv6rib_tables,json=numberOfRegisteredIpv6ribTables,proto3" json:"number_of_registered_ipv6rib_tables,omitempty"`
	IsBoundWithAtom                  bool              `protobuf:"varint,72,opt,name=is_bound_with_atom,json=isBoundWithAtom,proto3" json:"is_bound_with_atom,omitempty"`
	IsBoundWithNsrMate               bool              `protobuf:"varint,73,opt,name=is_bound_with_nsr_mate,json=isBoundWithNsrMate,proto3" json:"is_bound_with_nsr_mate,omitempty"`
	IsNsrConfigured                  bool              `protobuf:"varint,74,opt,name=is_nsr_configured,json=isNsrConfigured,proto3" json:"is_nsr_configured,omitempty"`
	IsMldpRegistered                 bool              `protobuf:"varint,75,opt,name=is_mldp_registered,json=isMldpRegistered,proto3" json:"is_mldp_registered,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}          `json:"-"`
	XXX_unrecognized                 []byte            `json:"-"`
	XXX_sizecache                    int32             `json:"-"`
}

func (m *LdpSummary) Reset()         { *m = LdpSummary{} }
func (m *LdpSummary) String() string { return proto.CompactTextString(m) }
func (*LdpSummary) ProtoMessage()    {}
func (*LdpSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7f2abcb6a0f7f76, []int{2}
}

func (m *LdpSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpSummary.Unmarshal(m, b)
}
func (m *LdpSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpSummary.Marshal(b, m, deterministic)
}
func (m *LdpSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpSummary.Merge(m, src)
}
func (m *LdpSummary) XXX_Size() int {
	return xxx_messageInfo_LdpSummary.Size(m)
}
func (m *LdpSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpSummary.DiscardUnknown(m)
}

var xxx_messageInfo_LdpSummary proto.InternalMessageInfo

func (m *LdpSummary) GetNumberOfVrf() uint32 {
	if m != nil {
		return m.NumberOfVrf
	}
	return 0
}

func (m *LdpSummary) GetNumberOfVrfOper() uint32 {
	if m != nil {
		return m.NumberOfVrfOper
	}
	return 0
}

func (m *LdpSummary) GetCommon() *LdpSummaryCommon {
	if m != nil {
		return m.Common
	}
	return nil
}

func (m *LdpSummary) GetNumberOfInterfaces() uint32 {
	if m != nil {
		return m.NumberOfInterfaces
	}
	return 0
}

func (m *LdpSummary) GetNumberOfFwdRefInterfaces() uint32 {
	if m != nil {
		return m.NumberOfFwdRefInterfaces
	}
	return 0
}

func (m *LdpSummary) GetNumberOfAutocfgInterfaces() uint32 {
	if m != nil {
		return m.NumberOfAutocfgInterfaces
	}
	return 0
}

func (m *LdpSummary) GetIsBoundWithSysdb() bool {
	if m != nil {
		return m.IsBoundWithSysdb
	}
	return false
}

func (m *LdpSummary) GetIsRegisteredWithSysdb() bool {
	if m != nil {
		return m.IsRegisteredWithSysdb
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithRsi() bool {
	if m != nil {
		return m.IsBoundWithRsi
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithInterfaceManager() bool {
	if m != nil {
		return m.IsBoundWithInterfaceManager
	}
	return false
}

func (m *LdpSummary) GetIsRegisteredWithInterfaceManager() bool {
	if m != nil {
		return m.IsRegisteredWithInterfaceManager
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithIpArm() bool {
	if m != nil {
		return m.IsBoundWithIpArm
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithLsd() bool {
	if m != nil {
		return m.IsBoundWithLsd
	}
	return false
}

func (m *LdpSummary) GetIsRegisteredWithLsd() bool {
	if m != nil {
		return m.IsRegisteredWithLsd
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithIpv4Rib() bool {
	if m != nil {
		return m.IsBoundWithIpv4Rib
	}
	return false
}

func (m *LdpSummary) GetIsRegisteredWithIpv4Rib() bool {
	if m != nil {
		return m.IsRegisteredWithIpv4Rib
	}
	return false
}

func (m *LdpSummary) GetNumberOfIpv4RibTables() uint32 {
	if m != nil {
		return m.NumberOfIpv4RibTables
	}
	return 0
}

func (m *LdpSummary) GetNumberOfRegisteredIpv4RibTables() uint32 {
	if m != nil {
		return m.NumberOfRegisteredIpv4RibTables
	}
	return 0
}

func (m *LdpSummary) GetIsBoundWithIpv6Rib() bool {
	if m != nil {
		return m.IsBoundWithIpv6Rib
	}
	return false
}

func (m *LdpSummary) GetIsRegisteredWithIpv6Rib() bool {
	if m != nil {
		return m.IsRegisteredWithIpv6Rib
	}
	return false
}

func (m *LdpSummary) GetNumberOfIpv6RibTables() uint32 {
	if m != nil {
		return m.NumberOfIpv6RibTables
	}
	return 0
}

func (m *LdpSummary) GetNumberOfRegisteredIpv6RibTables() uint32 {
	if m != nil {
		return m.NumberOfRegisteredIpv6RibTables
	}
	return 0
}

func (m *LdpSummary) GetIsBoundWithAtom() bool {
	if m != nil {
		return m.IsBoundWithAtom
	}
	return false
}

func (m *LdpSummary) GetIsBoundWithNsrMate() bool {
	if m != nil {
		return m.IsBoundWithNsrMate
	}
	return false
}

func (m *LdpSummary) GetIsNsrConfigured() bool {
	if m != nil {
		return m.IsNsrConfigured
	}
	return false
}

func (m *LdpSummary) GetIsMldpRegistered() bool {
	if m != nil {
		return m.IsMldpRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*LdpSummary_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.summary.ldp_summary_KEYS")
	proto.RegisterType((*LdpSummaryCommon)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.summary.ldp_summary_common")
	proto.RegisterType((*LdpSummary)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.summary.ldp_summary")
}

func init() { proto.RegisterFile("ldp_summary.proto", fileDescriptor_c7f2abcb6a0f7f76) }

var fileDescriptor_c7f2abcb6a0f7f76 = []byte{
	// 1003 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdb, 0x52, 0x1b, 0x37,
	0x1c, 0xc6, 0x87, 0x1e, 0x52, 0x90, 0x0b, 0x18, 0x01, 0x89, 0x5a, 0xda, 0xc6, 0x43, 0x7a, 0xe1,
	0x34, 0xad, 0xdb, 0x02, 0x11, 0x3d, 0xa4, 0x49, 0x4c, 0x08, 0x89, 0x13, 0x30, 0x33, 0x4b, 0x0f,
	0xd3, 0x2b, 0x8d, 0xd6, 0xab, 0x35, 0xca, 0xac, 0x56, 0x3b, 0xd2, 0x1a, 0xca, 0xe3, 0xf5, 0x39,
	0xfa, 0x32, 0x9d, 0x95, 0x77, 0xbd, 0xd2, 0xca, 0x40, 0xa7, 0x37, 0xcc, 0x20, 0x7d, 0xdf, 0x4f,
	0xdf, 0xff, 0x2f, 0xad, 0x64, 0xb0, 0x96, 0x44, 0x19, 0xd1, 0x13, 0x21, 0xa8, 0xba, 0xea, 0x65,
	0x4a, 0xe6, 0x12, 0xe2, 0x11, 0xd7, 0x23, 0x49, 0xb8, 0xd4, 0xe4, 0x2f, 0x45, 0x44, 0x96, 0x68,
	0x52, 0x88, 0x64, 0xc6, 0x54, 0xaf, 0xfa, 0xaf, 0x97, 0xca, 0x88, 0x69, 0xf3, 0xb7, 0x57, 0xba,
	0xb7, 0xbf, 0x05, 0x6d, 0x0b, 0x46, 0xde, 0xbe, 0xfc, 0xf3, 0x0c, 0x6e, 0x81, 0xa5, 0x42, 0x43,
	0x52, 0x2a, 0x18, 0x5a, 0xe8, 0x2c, 0x74, 0x97, 0x82, 0xc5, 0x62, 0x60, 0x48, 0x05, 0xdb, 0xfe,
	0x1b, 0x00, 0x68, 0x3b, 0x46, 0x52, 0x08, 0x99, 0xc2, 0x87, 0xa0, 0x4d, 0xa3, 0x48, 0x31, 0xad,
	0x49, 0x4c, 0x05, 0x4f, 0x38, 0xd3, 0xa5, 0x75, 0xb5, 0x1c, 0x3f, 0x2a, 0x87, 0x61, 0x17, 0xb4,
	0xd3, 0x89, 0x08, 0x99, 0x22, 0x32, 0x26, 0x3c, 0xbb, 0xd8, 0xa3, 0x31, 0x7a, 0xaf, 0xb3, 0xd0,
	0x5d, 0x0e, 0x56, 0xa6, 0xe3, 0xa7, 0xf1, 0xc0, 0x8c, 0x7a, 0x4a, 0x4c, 0x63, 0xf4, 0xbe, 0xa7,
	0xc4, 0x34, 0x86, 0x3d, 0xb0, 0x5e, 0x2b, 0x53, 0xc6, 0xc7, 0xe7, 0xa1, 0x54, 0x1a, 0x7d, 0x60,
	0xc4, 0x6b, 0x95, 0x78, 0x58, 0x4d, 0xc0, 0x03, 0xf0, 0x85, 0xa5, 0xd7, 0x8a, 0xe8, 0xab, 0x74,
	0xc4, 0x22, 0xcb, 0xfa, 0xa1, 0xb1, 0x7e, 0x3a, 0xb3, 0x6a, 0x75, 0x66, 0x24, 0x35, 0x63, 0x08,
	0xbe, 0xac, 0x19, 0x63, 0x45, 0x47, 0x2c, 0x9e, 0x24, 0x44, 0x31, 0x9d, 0x53, 0x95, 0x5b, 0xa4,
	0x3b, 0x86, 0xd4, 0xa9, 0x48, 0xaf, 0x4a, 0x65, 0x30, 0x15, 0xd6, 0xbc, 0xdf, 0x40, 0xb7, 0xe6,
	0x45, 0xf2, 0x32, 0xd5, 0xb9, 0x62, 0x54, 0x10, 0x99, 0x92, 0x88, 0x09, 0x9a, 0xda, 0xe9, 0x3e,
	0x32, 0xcc, 0x07, 0x15, 0xf3, 0x70, 0xa6, 0x3e, 0x4d, 0x0f, 0x8d, 0xb6, 0xc6, 0x62, 0x80, 0xdc,
	0x76, 0x93, 0x73, 0x96, 0x24, 0x92, 0xd0, 0xe8, 0x1d, 0x5a, 0x34, 0x98, 0x0d, 0xbb, 0xed, 0xaf,
	0x8b, 0xc9, 0x7e, 0xf4, 0xce, 0xf3, 0x61, 0xcb, 0xb7, 0xe4, 0xf9, 0xf0, 0xcc, 0xf7, 0x3d, 0xd8,
	0x6c, 0xac, 0xa7, 0xe4, 0x24, 0x67, 0x1a, 0x01, 0x63, 0x82, 0xf6, 0x62, 0x81, 0x99, 0xf1, 0x2c,
	0xb8, 0xb2, 0xb4, 0x3c, 0x0b, 0x2e, 0x2d, 0xce, 0x06, 0x9a, 0x55, 0x12, 0x39, 0xa2, 0x09, 0x29,
	0x0f, 0x1b, 0xd3, 0xe8, 0x63, 0x77, 0x03, 0x8b, 0xe5, 0x8e, 0x0b, 0x49, 0xbf, 0x52, 0x78, 0x0c,
	0xec, 0x31, 0x96, 0x3d, 0x06, 0x6e, 0x30, 0xf6, 0xed, 0x2e, 0x15, 0xdf, 0x05, 0x4f, 0x73, 0xa6,
	0x62, 0x3a, 0x62, 0x1a, 0xad, 0x18, 0xf7, 0x66, 0xe5, 0x3e, 0x8e, 0xb2, 0xc1, 0x6c, 0x12, 0x3e,
	0x03, 0x9f, 0xb9, 0x05, 0x34, 0xcc, 0xab, 0xc6, 0xfc, 0x89, 0x1d, 0x3f, 0xb9, 0x11, 0x80, 0x1b,
	0x80, 0xb6, 0x07, 0xc0, 0x2e, 0xe0, 0x31, 0xb8, 0x57, 0x03, 0x42, 0x9e, 0x46, 0x3c, 0x1d, 0x6b,
	0x13, 0x05, 0xad, 0xb9, 0xfb, 0x7b, 0x50, 0x4e, 0x16, 0x21, 0xae, 0xb7, 0x61, 0x04, 0xaf, 0xb5,
	0x61, 0xf8, 0x1c, 0x7c, 0x6e, 0x35, 0xca, 0xf4, 0xd9, 0x5d, 0x73, 0xdd, 0xcd, 0x6b, 0xfa, 0xec,
	0x2c, 0x7c, 0x0b, 0x01, 0xa3, 0x8d, 0x9b, 0x09, 0xd8, 0xdd, 0x70, 0xc5, 0x84, 0xcc, 0x59, 0x23,
	0xc4, 0xa6, 0xbb, 0xe1, 0x81, 0xd1, 0x38, 0x29, 0x6e, 0x63, 0x60, 0x74, 0xf7, 0x16, 0x06, 0xde,
	0xfe, 0xa7, 0x05, 0x5a, 0xd6, 0x1d, 0x0a, 0xb7, 0xc1, 0x72, 0xcd, 0xbc, 0x50, 0x31, 0xda, 0x31,
	0x88, 0x56, 0x85, 0xf8, 0x5d, 0xc5, 0xf0, 0x11, 0x80, 0x8e, 0xc6, 0x5c, 0xed, 0x68, 0xd7, 0x08,
	0x57, 0x2d, 0xe1, 0x69, 0xc6, 0x14, 0x0c, 0xc1, 0x9d, 0xe9, 0xbd, 0x8c, 0xf6, 0x3a, 0x0b, 0xdd,
	0xd6, 0xce, 0x9b, 0xde, 0xff, 0x7b, 0x1e, 0x7a, 0xfe, 0x4d, 0x1f, 0x94, 0x64, 0xf8, 0x1d, 0xd8,
	0xb0, 0xce, 0x5f, 0x7d, 0xee, 0x1e, 0x37, 0xbe, 0xd9, 0xfa, 0xc0, 0x3d, 0xb5, 0x4f, 0x6c, 0x7c,
	0x19, 0x11, 0xc5, 0x1c, 0x27, 0x36, 0x4e, 0x54, 0x39, 0x8f, 0x2e, 0xa3, 0x80, 0xc5, 0xd7, 0x9d,
	0x78, 0x3a, 0xc9, 0xe5, 0x28, 0x1e, 0xdb, 0xfe, 0x7d, 0x77, 0xff, 0xfb, 0x53, 0x85, 0x05, 0xf8,
	0x06, 0xac, 0x73, 0x4d, 0x42, 0x39, 0x49, 0x23, 0x72, 0xc9, 0xf3, 0x73, 0xa2, 0xaf, 0x74, 0x14,
	0xa2, 0x1f, 0x3a, 0x0b, 0xdd, 0xc5, 0xa0, 0xcd, 0xf5, 0x41, 0x31, 0xf3, 0x07, 0xcf, 0xcf, 0xcf,
	0x8a, 0xf1, 0xe2, 0xdb, 0xe6, 0x9a, 0x28, 0x36, 0xe6, 0x3a, 0x67, 0x8a, 0x39, 0x9e, 0x1f, 0x8d,
	0x67, 0x93, 0xeb, 0x60, 0x36, 0x5d, 0x1b, 0x1f, 0x82, 0x35, 0x77, 0x1d, 0xa5, 0x39, 0xfa, 0xc9,
	0x38, 0x56, 0xac, 0x55, 0x02, 0xcd, 0xe1, 0x21, 0xb8, 0xef, 0x4a, 0x67, 0xf5, 0x10, 0x41, 0x53,
	0x3a, 0x66, 0x0a, 0xfd, 0x6c, 0x8c, 0x5b, 0x96, 0x71, 0x56, 0xd2, 0xc9, 0x54, 0x52, 0x3c, 0x45,
	0x73, 0x92, 0xfa, 0xa8, 0x27, 0x06, 0xd5, 0x69, 0xa6, 0xf6, 0x78, 0x3d, 0xb0, 0xd1, 0x48, 0x95,
	0x11, 0xaa, 0x04, 0xfa, 0xc5, 0xeb, 0xd4, 0x20, 0xeb, 0x2b, 0xe1, 0x17, 0x9c, 0xe8, 0x08, 0x3d,
	0xf5, 0x0a, 0x3e, 0xd6, 0x11, 0xdc, 0x05, 0x77, 0xe7, 0x44, 0x2d, 0xf4, 0xcf, 0x8c, 0x7e, 0xbd,
	0x19, 0xae, 0x30, 0xed, 0x18, 0x93, 0x93, 0xa7, 0x78, 0x57, 0x78, 0x88, 0x9e, 0x1b, 0x13, 0x74,
	0x12, 0x5d, 0xec, 0x05, 0x3c, 0x84, 0x4f, 0xc0, 0xd6, 0xbc, 0x9e, 0x54, 0xc6, 0xbe, 0x31, 0xde,
	0xf3, 0x5a, 0x51, 0xba, 0xf7, 0x9b, 0xaf, 0xa6, 0xe2, 0x21, 0xc9, 0x69, 0x98, 0x30, 0x8d, 0x0e,
	0xdc, 0x7b, 0x7d, 0x30, 0x9d, 0xfd, 0xd5, 0x4c, 0xc2, 0x63, 0xf0, 0xc0, 0xbe, 0x1f, 0x66, 0xab,
	0x37, 0x18, 0x2f, 0x0c, 0xe3, 0x7e, 0x7d, 0x49, 0x54, 0x42, 0x97, 0x36, 0xaf, 0x70, 0x6c, 0xf2,
	0x1f, 0xce, 0x2b, 0x1c, 0xdf, 0x58, 0xf8, 0xd4, 0xf8, 0xf2, 0xda, 0xc2, 0xf1, 0xbc, 0xc2, 0xb1,
	0x15, 0xfa, 0xc8, 0x2b, 0x1c, 0xff, 0xa7, 0xc2, 0x6d, 0xc6, 0xab, 0x1b, 0x0a, 0xb7, 0x68, 0x8f,
	0x00, 0x74, 0x0b, 0xa7, 0xb9, 0x14, 0xe8, 0xb5, 0xc9, 0xbe, 0x6a, 0x15, 0xdd, 0xcf, 0xa5, 0xf0,
	0xbb, 0x54, 0xfc, 0xa2, 0x13, 0x34, 0x67, 0x68, 0xe0, 0x75, 0x69, 0xa8, 0xd5, 0x09, 0xcd, 0x19,
	0xfc, 0xca, 0x1c, 0xd9, 0x42, 0x38, 0x92, 0x69, 0xcc, 0xc7, 0x13, 0xc5, 0x22, 0xf4, 0xa6, 0xe2,
	0x0f, 0xb5, 0x7a, 0x31, 0x1b, 0x86, 0x5f, 0x9b, 0x30, 0xa2, 0xb8, 0x0c, 0xeb, 0xc2, 0xd0, 0xdb,
	0xea, 0x63, 0x38, 0x49, 0xa2, 0xac, 0xae, 0x23, 0xbc, 0x63, 0x7e, 0x91, 0xef, 0xfe, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x9f, 0x75, 0x61, 0x5b, 0xa6, 0x0b, 0x00, 0x00,
}