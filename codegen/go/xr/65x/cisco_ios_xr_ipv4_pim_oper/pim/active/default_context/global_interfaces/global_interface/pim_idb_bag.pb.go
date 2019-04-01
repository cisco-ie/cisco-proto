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
// source: pim_idb_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_global_interfaces_global_interface

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

type PimIdbBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimIdbBag_KEYS) Reset()         { *m = PimIdbBag_KEYS{} }
func (m *PimIdbBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimIdbBag_KEYS) ProtoMessage()    {}
func (*PimIdbBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cef3e95a85205ea, []int{0}
}

func (m *PimIdbBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIdbBag_KEYS.Unmarshal(m, b)
}
func (m *PimIdbBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIdbBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimIdbBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIdbBag_KEYS.Merge(m, src)
}
func (m *PimIdbBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimIdbBag_KEYS.Size(m)
}
func (m *PimIdbBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIdbBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimIdbBag_KEYS proto.InternalMessageInfo

func (m *PimIdbBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cef3e95a85205ea, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimIdbBag struct {
	InterfaceNameXr            string         `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	InterfaceAddress           []*PimAddrtype `protobuf:"bytes,51,rep,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	InterfaceAddressMask       uint32         `protobuf:"varint,52,opt,name=interface_address_mask,json=interfaceAddressMask,proto3" json:"interface_address_mask,omitempty"`
	IsEnabled                  bool           `protobuf:"varint,53,opt,name=is_enabled,json=isEnabled,proto3" json:"is_enabled,omitempty"`
	NeighborCount              uint32         `protobuf:"varint,54,opt,name=neighbor_count,json=neighborCount,proto3" json:"neighbor_count,omitempty"`
	ExternalNeighborCount      uint32         `protobuf:"varint,55,opt,name=external_neighbor_count,json=externalNeighborCount,proto3" json:"external_neighbor_count,omitempty"`
	HelloInterval              uint32         `protobuf:"varint,56,opt,name=hello_interval,json=helloInterval,proto3" json:"hello_interval,omitempty"`
	DrAddress                  *PimAddrtype   `protobuf:"bytes,57,opt,name=dr_address,json=drAddress,proto3" json:"dr_address,omitempty"`
	DrPriority                 uint32         `protobuf:"varint,58,opt,name=dr_priority,json=drPriority,proto3" json:"dr_priority,omitempty"`
	HelloExpiry                uint64         `protobuf:"varint,59,opt,name=hello_expiry,json=helloExpiry,proto3" json:"hello_expiry,omitempty"`
	AreWeDr                    bool           `protobuf:"varint,60,opt,name=are_we_dr,json=areWeDr,proto3" json:"are_we_dr,omitempty"`
	BfdEnabled                 bool           `protobuf:"varint,61,opt,name=bfd_enabled,json=bfdEnabled,proto3" json:"bfd_enabled,omitempty"`
	BfdInterval                uint32         `protobuf:"varint,62,opt,name=bfd_interval,json=bfdInterval,proto3" json:"bfd_interval,omitempty"`
	BfdMultiplier              uint32         `protobuf:"varint,63,opt,name=bfd_multiplier,json=bfdMultiplier,proto3" json:"bfd_multiplier,omitempty"`
	VirtualInterface           bool           `protobuf:"varint,64,opt,name=virtual_interface,json=virtualInterface,proto3" json:"virtual_interface,omitempty"`
	PassiveInterface           bool           `protobuf:"varint,65,opt,name=passive_interface,json=passiveInterface,proto3" json:"passive_interface,omitempty"`
	NeighborFilterName         string         `protobuf:"bytes,66,opt,name=neighbor_filter_name,json=neighborFilterName,proto3" json:"neighbor_filter_name,omitempty"`
	JoinPruneInterval          uint32         `protobuf:"varint,67,opt,name=join_prune_interval,json=joinPruneInterval,proto3" json:"join_prune_interval,omitempty"`
	PruneDelayEnabled          bool           `protobuf:"varint,68,opt,name=prune_delay_enabled,json=pruneDelayEnabled,proto3" json:"prune_delay_enabled,omitempty"`
	ConfiguredPropagationDelay uint32         `protobuf:"varint,69,opt,name=configured_propagation_delay,json=configuredPropagationDelay,proto3" json:"configured_propagation_delay,omitempty"`
	PropagationDelay           uint32         `protobuf:"varint,70,opt,name=propagation_delay,json=propagationDelay,proto3" json:"propagation_delay,omitempty"`
	ConfiguredOverrideInterval uint32         `protobuf:"varint,71,opt,name=configured_override_interval,json=configuredOverrideInterval,proto3" json:"configured_override_interval,omitempty"`
	OverrideInterval           uint32         `protobuf:"varint,72,opt,name=override_interval,json=overrideInterval,proto3" json:"override_interval,omitempty"`
	GenerationId               uint32         `protobuf:"varint,73,opt,name=generation_id,json=generationId,proto3" json:"generation_id,omitempty"`
	IsBidirectionalCapable     bool           `protobuf:"varint,74,opt,name=is_bidirectional_capable,json=isBidirectionalCapable,proto3" json:"is_bidirectional_capable,omitempty"`
	IsProxyCapable             bool           `protobuf:"varint,75,opt,name=is_proxy_capable,json=isProxyCapable,proto3" json:"is_proxy_capable,omitempty"`
	IsBatchAssertsCapable      bool           `protobuf:"varint,76,opt,name=is_batch_asserts_capable,json=isBatchAssertsCapable,proto3" json:"is_batch_asserts_capable,omitempty"`
	IdbOorEnabled              bool           `protobuf:"varint,77,opt,name=idb_oor_enabled,json=idbOorEnabled,proto3" json:"idb_oor_enabled,omitempty"`
	IdbAclProvided             bool           `protobuf:"varint,78,opt,name=idb_acl_provided,json=idbAclProvided,proto3" json:"idb_acl_provided,omitempty"`
	IdbMaxCount                uint32         `protobuf:"varint,79,opt,name=idb_max_count,json=idbMaxCount,proto3" json:"idb_max_count,omitempty"`
	IdbThresholdCount          uint32         `protobuf:"varint,80,opt,name=idb_threshold_count,json=idbThresholdCount,proto3" json:"idb_threshold_count,omitempty"`
	IdbCurrentCount            uint32         `protobuf:"varint,81,opt,name=idb_current_count,json=idbCurrentCount,proto3" json:"idb_current_count,omitempty"`
	IdbAclName                 string         `protobuf:"bytes,82,opt,name=idb_acl_name,json=idbAclName,proto3" json:"idb_acl_name,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}       `json:"-"`
	XXX_unrecognized           []byte         `json:"-"`
	XXX_sizecache              int32          `json:"-"`
}

func (m *PimIdbBag) Reset()         { *m = PimIdbBag{} }
func (m *PimIdbBag) String() string { return proto.CompactTextString(m) }
func (*PimIdbBag) ProtoMessage()    {}
func (*PimIdbBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cef3e95a85205ea, []int{2}
}

func (m *PimIdbBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimIdbBag.Unmarshal(m, b)
}
func (m *PimIdbBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimIdbBag.Marshal(b, m, deterministic)
}
func (m *PimIdbBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimIdbBag.Merge(m, src)
}
func (m *PimIdbBag) XXX_Size() int {
	return xxx_messageInfo_PimIdbBag.Size(m)
}
func (m *PimIdbBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimIdbBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimIdbBag proto.InternalMessageInfo

func (m *PimIdbBag) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *PimIdbBag) GetInterfaceAddress() []*PimAddrtype {
	if m != nil {
		return m.InterfaceAddress
	}
	return nil
}

func (m *PimIdbBag) GetInterfaceAddressMask() uint32 {
	if m != nil {
		return m.InterfaceAddressMask
	}
	return 0
}

func (m *PimIdbBag) GetIsEnabled() bool {
	if m != nil {
		return m.IsEnabled
	}
	return false
}

func (m *PimIdbBag) GetNeighborCount() uint32 {
	if m != nil {
		return m.NeighborCount
	}
	return 0
}

func (m *PimIdbBag) GetExternalNeighborCount() uint32 {
	if m != nil {
		return m.ExternalNeighborCount
	}
	return 0
}

func (m *PimIdbBag) GetHelloInterval() uint32 {
	if m != nil {
		return m.HelloInterval
	}
	return 0
}

func (m *PimIdbBag) GetDrAddress() *PimAddrtype {
	if m != nil {
		return m.DrAddress
	}
	return nil
}

func (m *PimIdbBag) GetDrPriority() uint32 {
	if m != nil {
		return m.DrPriority
	}
	return 0
}

func (m *PimIdbBag) GetHelloExpiry() uint64 {
	if m != nil {
		return m.HelloExpiry
	}
	return 0
}

func (m *PimIdbBag) GetAreWeDr() bool {
	if m != nil {
		return m.AreWeDr
	}
	return false
}

func (m *PimIdbBag) GetBfdEnabled() bool {
	if m != nil {
		return m.BfdEnabled
	}
	return false
}

func (m *PimIdbBag) GetBfdInterval() uint32 {
	if m != nil {
		return m.BfdInterval
	}
	return 0
}

func (m *PimIdbBag) GetBfdMultiplier() uint32 {
	if m != nil {
		return m.BfdMultiplier
	}
	return 0
}

func (m *PimIdbBag) GetVirtualInterface() bool {
	if m != nil {
		return m.VirtualInterface
	}
	return false
}

func (m *PimIdbBag) GetPassiveInterface() bool {
	if m != nil {
		return m.PassiveInterface
	}
	return false
}

func (m *PimIdbBag) GetNeighborFilterName() string {
	if m != nil {
		return m.NeighborFilterName
	}
	return ""
}

func (m *PimIdbBag) GetJoinPruneInterval() uint32 {
	if m != nil {
		return m.JoinPruneInterval
	}
	return 0
}

func (m *PimIdbBag) GetPruneDelayEnabled() bool {
	if m != nil {
		return m.PruneDelayEnabled
	}
	return false
}

func (m *PimIdbBag) GetConfiguredPropagationDelay() uint32 {
	if m != nil {
		return m.ConfiguredPropagationDelay
	}
	return 0
}

func (m *PimIdbBag) GetPropagationDelay() uint32 {
	if m != nil {
		return m.PropagationDelay
	}
	return 0
}

func (m *PimIdbBag) GetConfiguredOverrideInterval() uint32 {
	if m != nil {
		return m.ConfiguredOverrideInterval
	}
	return 0
}

func (m *PimIdbBag) GetOverrideInterval() uint32 {
	if m != nil {
		return m.OverrideInterval
	}
	return 0
}

func (m *PimIdbBag) GetGenerationId() uint32 {
	if m != nil {
		return m.GenerationId
	}
	return 0
}

func (m *PimIdbBag) GetIsBidirectionalCapable() bool {
	if m != nil {
		return m.IsBidirectionalCapable
	}
	return false
}

func (m *PimIdbBag) GetIsProxyCapable() bool {
	if m != nil {
		return m.IsProxyCapable
	}
	return false
}

func (m *PimIdbBag) GetIsBatchAssertsCapable() bool {
	if m != nil {
		return m.IsBatchAssertsCapable
	}
	return false
}

func (m *PimIdbBag) GetIdbOorEnabled() bool {
	if m != nil {
		return m.IdbOorEnabled
	}
	return false
}

func (m *PimIdbBag) GetIdbAclProvided() bool {
	if m != nil {
		return m.IdbAclProvided
	}
	return false
}

func (m *PimIdbBag) GetIdbMaxCount() uint32 {
	if m != nil {
		return m.IdbMaxCount
	}
	return 0
}

func (m *PimIdbBag) GetIdbThresholdCount() uint32 {
	if m != nil {
		return m.IdbThresholdCount
	}
	return 0
}

func (m *PimIdbBag) GetIdbCurrentCount() uint32 {
	if m != nil {
		return m.IdbCurrentCount
	}
	return 0
}

func (m *PimIdbBag) GetIdbAclName() string {
	if m != nil {
		return m.IdbAclName
	}
	return ""
}

func init() {
	proto.RegisterType((*PimIdbBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.global_interfaces.global_interface.pim_idb_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.global_interfaces.global_interface.pim_addrtype")
	proto.RegisterType((*PimIdbBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.global_interfaces.global_interface.pim_idb_bag")
}

func init() { proto.RegisterFile("pim_idb_bag.proto", fileDescriptor_9cef3e95a85205ea) }

var fileDescriptor_9cef3e95a85205ea = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xed, 0x8e, 0xdb, 0x44,
	0x14, 0x86, 0x15, 0x8a, 0x5a, 0x76, 0x92, 0x6c, 0x13, 0x6f, 0x3f, 0xac, 0x0a, 0x44, 0xba, 0x68,
	0x51, 0x44, 0x25, 0x0b, 0xb5, 0xcb, 0xb6, 0xe5, 0xb3, 0xfb, 0x55, 0x58, 0xca, 0xee, 0x86, 0x80,
	0x44, 0xf9, 0x35, 0x1a, 0x7b, 0x8e, 0x93, 0x43, 0x1d, 0x8f, 0x75, 0x66, 0x12, 0x92, 0x0b, 0xe0,
	0x0a, 0xb8, 0x5e, 0x24, 0x34, 0x33, 0xb6, 0xe3, 0x4d, 0x7e, 0xc3, 0xdf, 0xf7, 0x7d, 0xce, 0xcc,
	0x7b, 0xce, 0xcc, 0xd8, 0xac, 0x5f, 0xe0, 0x8c, 0xa3, 0x8c, 0x79, 0x2c, 0x26, 0x51, 0x41, 0xca,
	0xa8, 0xe0, 0x6d, 0x82, 0x3a, 0x51, 0x1c, 0x95, 0xe6, 0x4b, 0xe2, 0x58, 0x2c, 0x0e, 0xb9, 0x85,
	0x54, 0x01, 0x14, 0x15, 0x38, 0x8b, 0x44, 0x62, 0x70, 0x01, 0x91, 0x84, 0x54, 0xcc, 0x33, 0xc3,
	0x13, 0x95, 0x1b, 0x58, 0x9a, 0x68, 0x92, 0xa9, 0x58, 0x64, 0x1c, 0x73, 0x03, 0x94, 0x8a, 0x04,
	0xf4, 0x96, 0xb2, 0xff, 0x92, 0xf5, 0x1a, 0xdb, 0xf1, 0x37, 0xe7, 0xbf, 0xff, 0x12, 0x1c, 0xb0,
	0xdd, 0x1a, 0xe0, 0xb9, 0x98, 0x41, 0xd8, 0x1a, 0xb4, 0x86, 0x3b, 0xe3, 0x6e, 0xad, 0x5e, 0x89,
	0x19, 0xec, 0xcf, 0x58, 0xc7, 0x96, 0x0a, 0x29, 0xc9, 0xac, 0x0a, 0x08, 0x1e, 0xb2, 0x3b, 0x22,
	0x6d, 0xf2, 0xb7, 0x45, 0x6a, 0xc1, 0xe0, 0x31, 0xeb, 0xb8, 0xc8, 0x96, 0x04, 0xad, 0xc3, 0xf7,
	0x9c, 0xdb, 0xb6, 0xda, 0xb1, 0x97, 0x4a, 0xe4, 0xa8, 0x46, 0x6e, 0xd5, 0xc8, 0x51, 0x89, 0xec,
	0xff, 0xd3, 0x61, 0xed, 0x46, 0xd4, 0xe0, 0x33, 0xd6, 0xbf, 0x99, 0x92, 0x2f, 0x29, 0x7c, 0xea,
	0xea, 0xee, 0xde, 0x08, 0xfa, 0x96, 0x82, 0xbf, 0x5b, 0x4d, 0xb8, 0xda, 0xe4, 0xd9, 0xe0, 0xd6,
	0xb0, 0xfd, 0x34, 0x8d, 0xfe, 0xab, 0xe1, 0x46, 0xcd, 0xf1, 0x8c, 0x7b, 0xb5, 0x5e, 0x35, 0x7d,
	0xc8, 0x1e, 0x6c, 0x85, 0xe2, 0x33, 0xa1, 0xdf, 0x85, 0x87, 0x83, 0xd6, 0xb0, 0x3b, 0xbe, 0xb7,
	0x59, 0x71, 0x29, 0xf4, 0xbb, 0xe0, 0x23, 0xc6, 0x50, 0x73, 0xc8, 0x45, 0x9c, 0x81, 0x0c, 0xbf,
	0x18, 0xb4, 0x86, 0x1f, 0x8c, 0x77, 0x50, 0x9f, 0x7b, 0xc1, 0x1e, 0x5e, 0x0e, 0x38, 0x99, 0xc6,
	0x8a, 0x78, 0xa2, 0xe6, 0xb9, 0x09, 0x8f, 0xdc, 0x62, 0xdd, 0x4a, 0x3d, 0xb5, 0x62, 0x70, 0xc4,
	0x1e, 0xc2, 0xd2, 0x00, 0xe5, 0x22, 0xe3, 0x1b, 0xfc, 0x73, 0xc7, 0xdf, 0xaf, 0xec, 0xab, 0x1b,
	0x75, 0x07, 0x6c, 0x77, 0x0a, 0x59, 0xa6, 0x7c, 0x97, 0x0b, 0x91, 0x85, 0x2f, 0xfc, 0xf2, 0x4e,
	0xbd, 0x28, 0xc5, 0xe0, 0xaf, 0x16, 0x63, 0x92, 0xea, 0x49, 0xbf, 0x1c, 0xb4, 0xfe, 0xc7, 0x49,
	0xef, 0x48, 0xaa, 0x46, 0xfc, 0x31, 0x6b, 0x4b, 0xe2, 0x05, 0xa1, 0x22, 0x34, 0xab, 0xf0, 0x4b,
	0x97, 0x95, 0x49, 0x1a, 0x95, 0x8a, 0xbd, 0x78, 0xbe, 0x1f, 0x58, 0x16, 0x48, 0xab, 0xf0, 0xab,
	0x41, 0x6b, 0xf8, 0xfe, 0xb8, 0xed, 0xb4, 0x73, 0x27, 0x05, 0x8f, 0xd8, 0x8e, 0x20, 0xe0, 0x7f,
	0x02, 0x97, 0x14, 0x7e, 0xed, 0xe6, 0x7d, 0x47, 0x10, 0xfc, 0x06, 0x67, 0x64, 0xd7, 0x8f, 0x53,
	0x59, 0x9f, 0xc6, 0x37, 0xce, 0x65, 0x71, 0x2a, 0xab, 0xe3, 0x78, 0xcc, 0x3a, 0x16, 0xa8, 0xa7,
	0xf5, 0xad, 0x4b, 0x60, 0x8b, 0xea, 0x59, 0x1d, 0xb0, 0x5d, 0x8b, 0xcc, 0xe6, 0x99, 0xc1, 0x22,
	0x43, 0xa0, 0xf0, 0x3b, 0x3f, 0xd2, 0x38, 0x95, 0x97, 0xb5, 0x18, 0x3c, 0x61, 0xfd, 0x05, 0x92,
	0x99, 0x37, 0xfb, 0x0e, 0x5f, 0xb9, 0x0d, 0x7b, 0xa5, 0x71, 0x51, 0xe9, 0x16, 0x2e, 0x84, 0xd6,
	0xb8, 0x80, 0x06, 0x7c, 0xec, 0xe1, 0xd2, 0x58, 0xc3, 0x9f, 0xb3, 0x7b, 0xf5, 0x15, 0x48, 0x31,
	0x33, 0x40, 0xfe, 0x15, 0x9f, 0xb8, 0xc7, 0x14, 0x54, 0xde, 0x6b, 0x67, 0xb9, 0x17, 0x1d, 0xb1,
	0xbd, 0x3f, 0x14, 0xe6, 0xbc, 0xa0, 0x79, 0x0e, 0xeb, 0xe6, 0x4e, 0x5d, 0xee, 0xbe, 0xb5, 0x46,
	0xd6, 0xa9, 0x5b, 0x8c, 0xd8, 0x9e, 0x47, 0x25, 0x64, 0x62, 0x55, 0x8f, 0xeb, 0xcc, 0x05, 0xea,
	0x3b, 0xeb, 0xcc, 0x3a, 0xd5, 0xd4, 0x5e, 0xb1, 0x0f, 0x13, 0x95, 0xa7, 0x38, 0x99, 0x13, 0x48,
	0x5e, 0x90, 0x2a, 0xc4, 0x44, 0x18, 0x54, 0xb9, 0x5f, 0x20, 0x3c, 0x77, 0x1b, 0x3d, 0x5a, 0x33,
	0xa3, 0x35, 0xe2, 0x16, 0x72, 0x03, 0xd8, 0x2a, 0x7b, 0xed, 0xca, 0x7a, 0xc5, 0x26, 0x7c, 0x73,
	0x3b, 0xb5, 0x00, 0x22, 0x94, 0x8d, 0xbe, 0xbe, 0xdf, 0xdc, 0xee, 0xba, 0x44, 0xea, 0x06, 0x9f,
	0xb0, 0xfe, 0x76, 0xd9, 0x0f, 0x7e, 0x3b, 0xb5, 0x09, 0x7f, 0xc2, 0xba, 0x13, 0xc8, 0x81, 0x7c,
	0x34, 0x94, 0xe1, 0x85, 0x03, 0x3b, 0x6b, 0xf1, 0x42, 0x06, 0x2f, 0x58, 0x88, 0x9a, 0xc7, 0x28,
	0x91, 0x20, 0xb1, 0x9a, 0xc8, 0x78, 0x22, 0x0a, 0x3b, 0x9f, 0xf0, 0x47, 0x37, 0xb7, 0x07, 0xa8,
	0x4f, 0x9a, 0xf6, 0xa9, 0x77, 0x83, 0x21, 0xeb, 0xa1, 0xb6, 0x43, 0x5b, 0xae, 0xea, 0x8a, 0x37,
	0xae, 0x62, 0x17, 0xf5, 0xc8, 0xca, 0x15, 0xf9, 0xdc, 0xef, 0x21, 0x4c, 0x32, 0xe5, 0x42, 0x6b,
	0x20, 0xa3, 0xeb, 0x8a, 0x9f, 0x5c, 0xc5, 0x7d, 0xd4, 0x27, 0xd6, 0x3e, 0xf6, 0x6e, 0x55, 0xf8,
	0x29, 0xbb, 0x6b, 0x3f, 0xc3, 0x4a, 0x51, 0x7d, 0x96, 0x97, 0x8e, 0xef, 0xa2, 0x8c, 0xaf, 0x15,
	0x55, 0xe7, 0x68, 0xa3, 0xc8, 0x98, 0x8b, 0x24, 0xb3, 0x79, 0x16, 0x28, 0x41, 0x86, 0x57, 0x65,
	0x14, 0x19, 0x1f, 0x27, 0xd9, 0xa8, 0x54, 0x83, 0x7d, 0x66, 0x4b, 0xf9, 0x4c, 0x2c, 0xcb, 0xaf,
	0xd0, 0xb5, 0x7f, 0x28, 0x28, 0xe3, 0x4b, 0xb1, 0xf4, 0xdf, 0x9e, 0x88, 0xed, 0x59, 0xc6, 0x4c,
	0x09, 0xf4, 0x54, 0x65, 0xb2, 0x24, 0x47, 0xfe, 0xd6, 0xa1, 0x8c, 0x7f, 0xad, 0x1c, 0xcf, 0xdb,
	0x3f, 0x84, 0x8c, 0x79, 0x32, 0x27, 0x82, 0xdc, 0x94, 0xf4, 0xcf, 0x8e, 0xb6, 0xf1, 0x4f, 0xbd,
	0xee, 0xd9, 0x01, 0xeb, 0x54, 0x49, 0xdd, 0xdd, 0x1f, 0xbb, 0xbb, 0xcf, 0x7c, 0x4a, 0x7b, 0xe7,
	0xe3, 0xdb, 0xee, 0x57, 0xfc, 0xec, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x50, 0x0b, 0x89,
	0x9f, 0x07, 0x00, 0x00,
}
