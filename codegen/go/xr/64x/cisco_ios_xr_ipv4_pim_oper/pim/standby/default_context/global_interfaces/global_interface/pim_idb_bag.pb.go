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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_global_interfaces_global_interface

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
	proto.RegisterType((*PimIdbBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.global_interfaces.global_interface.pim_idb_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.global_interfaces.global_interface.pim_addrtype")
	proto.RegisterType((*PimIdbBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.global_interfaces.global_interface.pim_idb_bag")
}

func init() { proto.RegisterFile("pim_idb_bag.proto", fileDescriptor_9cef3e95a85205ea) }

var fileDescriptor_9cef3e95a85205ea = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xed, 0x8e, 0xdb, 0x44,
	0x14, 0x86, 0x15, 0x8a, 0x5a, 0xf6, 0x24, 0xd9, 0x26, 0xde, 0x7e, 0x58, 0x15, 0x88, 0x74, 0xd1,
	0xa2, 0x88, 0x4a, 0x16, 0x6a, 0x97, 0x6d, 0xcb, 0x67, 0xf7, 0xab, 0xb0, 0x94, 0xdd, 0x0d, 0x01,
	0x09, 0xfa, 0x6b, 0x34, 0xf6, 0x8c, 0x93, 0x43, 0x1d, 0x8f, 0x75, 0x66, 0x12, 0x92, 0x1b, 0xe0,
	0x0a, 0xb8, 0x5f, 0xd0, 0xcc, 0xd8, 0x8e, 0x37, 0xf9, 0x0d, 0x7f, 0xdf, 0xf7, 0x39, 0x33, 0xef,
	0x39, 0x33, 0x63, 0x43, 0xbf, 0xc0, 0x19, 0x43, 0x11, 0xb3, 0x98, 0x4f, 0xa2, 0x82, 0x94, 0x51,
	0xc1, 0xdb, 0x04, 0x75, 0xa2, 0x18, 0x2a, 0xcd, 0x96, 0xc4, 0xb0, 0x58, 0x1c, 0x32, 0x0b, 0xa9,
	0x42, 0x52, 0x54, 0xe0, 0x2c, 0xd2, 0x86, 0xe7, 0x22, 0x5e, 0x45, 0x42, 0xa6, 0x7c, 0x9e, 0x19,
	0x96, 0xa8, 0xdc, 0xc8, 0xa5, 0x89, 0x26, 0x99, 0x8a, 0x79, 0xc6, 0x30, 0x37, 0x92, 0x52, 0x9e,
	0x48, 0xbd, 0xa5, 0xec, 0xbf, 0x84, 0x5e, 0x63, 0x3f, 0xf6, 0xe6, 0xfc, 0xed, 0x2f, 0xc1, 0x01,
	0xec, 0xd6, 0x00, 0xcb, 0xf9, 0x4c, 0x86, 0xad, 0x41, 0x6b, 0xb8, 0x33, 0xee, 0xd6, 0xea, 0x15,
	0x9f, 0xc9, 0xfd, 0x19, 0x74, 0x6c, 0x29, 0x17, 0x82, 0xcc, 0xaa, 0x90, 0xc1, 0x43, 0xb8, 0xc3,
	0xd3, 0x26, 0x7f, 0x9b, 0xa7, 0x16, 0x0c, 0x1e, 0x43, 0xc7, 0x65, 0xb6, 0xa4, 0xd4, 0x3a, 0x7c,
	0xcf, 0xb9, 0x6d, 0xab, 0x1d, 0x7b, 0xa9, 0x44, 0x8e, 0x6a, 0xe4, 0x56, 0x8d, 0x1c, 0x95, 0xc8,
	0xfe, 0x3f, 0x1d, 0x68, 0x37, 0xa2, 0x06, 0x9f, 0x41, 0xff, 0x66, 0x4a, 0xb6, 0xa4, 0xf0, 0xa9,
	0xab, 0xbb, 0x7b, 0x23, 0xe8, 0xef, 0x14, 0xfc, 0xdd, 0x6a, 0xc2, 0xd5, 0x26, 0xcf, 0x06, 0xb7,
	0x86, 0xed, 0xa7, 0x93, 0xe8, 0x3f, 0x9b, 0x6e, 0xd4, 0x9c, 0xcf, 0xb8, 0x57, 0xeb, 0x55, 0xd7,
	0x87, 0xf0, 0x60, 0x2b, 0x15, 0x9b, 0x71, 0xfd, 0x2e, 0x3c, 0x1c, 0xb4, 0x86, 0xdd, 0xf1, 0xbd,
	0xcd, 0x8a, 0x4b, 0xae, 0xdf, 0x05, 0x1f, 0x01, 0xa0, 0x66, 0x32, 0xe7, 0x71, 0x26, 0x45, 0xf8,
	0xc5, 0xa0, 0x35, 0xfc, 0x60, 0xbc, 0x83, 0xfa, 0xdc, 0x0b, 0xf6, 0xf4, 0x72, 0x89, 0x93, 0x69,
	0xac, 0x88, 0x25, 0x6a, 0x9e, 0x9b, 0xf0, 0xc8, 0x2d, 0xd6, 0xad, 0xd4, 0x53, 0x2b, 0x06, 0x47,
	0xf0, 0x50, 0x2e, 0x8d, 0xa4, 0x9c, 0x67, 0x6c, 0x83, 0x7f, 0xee, 0xf8, 0xfb, 0x95, 0x7d, 0x75,
	0xa3, 0xee, 0x00, 0x76, 0xa7, 0x32, 0xcb, 0x94, 0xef, 0x72, 0xc1, 0xb3, 0xf0, 0x85, 0x5f, 0xde,
	0xa9, 0x17, 0xa5, 0x18, 0xfc, 0xd5, 0x02, 0x10, 0x54, 0x8f, 0xfa, 0xe5, 0xa0, 0xf5, 0x7f, 0x8e,
	0x7a, 0x47, 0x50, 0x35, 0xe3, 0x8f, 0xa1, 0x2d, 0x88, 0x15, 0x84, 0x8a, 0xd0, 0xac, 0xc2, 0x2f,
	0x5d, 0x58, 0x10, 0x34, 0x2a, 0x15, 0x7b, 0xf5, 0x7c, 0x43, 0x72, 0x59, 0x20, 0xad, 0xc2, 0xaf,
	0x06, 0xad, 0xe1, 0xfb, 0xe3, 0xb6, 0xd3, 0xce, 0x9d, 0x14, 0x3c, 0x82, 0x1d, 0x4e, 0x92, 0xfd,
	0x29, 0x99, 0xa0, 0xf0, 0x6b, 0x37, 0xf0, 0x3b, 0x9c, 0xe4, 0x6f, 0xf2, 0x8c, 0xec, 0xfa, 0x71,
	0x2a, 0xea, 0xe3, 0xf8, 0xc6, 0xb9, 0x10, 0xa7, 0xa2, 0x3a, 0x8f, 0xc7, 0xd0, 0xb1, 0x40, 0x3d,
	0xae, 0x6f, 0x5d, 0x02, 0x5b, 0x54, 0x0f, 0xeb, 0x00, 0x76, 0x2d, 0x32, 0x9b, 0x67, 0x06, 0x8b,
	0x0c, 0x25, 0x85, 0xdf, 0xf9, 0x99, 0xc6, 0xa9, 0xb8, 0xac, 0xc5, 0xe0, 0x09, 0xf4, 0x17, 0x48,
	0x66, 0xde, 0xec, 0x3b, 0x7c, 0xe5, 0x36, 0xec, 0x95, 0xc6, 0x45, 0xa5, 0x5b, 0xb8, 0xe0, 0x5a,
	0xe3, 0x42, 0x36, 0xe0, 0x63, 0x0f, 0x97, 0xc6, 0x1a, 0xfe, 0x1c, 0xee, 0xd5, 0x77, 0x20, 0xc5,
	0xcc, 0x48, 0xf2, 0xef, 0xf8, 0xc4, 0x3d, 0xa7, 0xa0, 0xf2, 0x5e, 0x3b, 0xcb, 0xbd, 0xe9, 0x08,
	0xf6, 0xfe, 0x50, 0x98, 0xb3, 0x82, 0xe6, 0xb9, 0x5c, 0x37, 0x77, 0xea, 0x72, 0xf7, 0xad, 0x35,
	0xb2, 0x4e, 0xdd, 0x62, 0x04, 0x7b, 0x1e, 0x15, 0x32, 0xe3, 0xab, 0x7a, 0x5c, 0x67, 0x2e, 0x50,
	0xdf, 0x59, 0x67, 0xd6, 0xa9, 0xa6, 0xf6, 0x0a, 0x3e, 0x4c, 0x54, 0x9e, 0xe2, 0x64, 0x4e, 0x52,
	0xb0, 0x82, 0x54, 0xc1, 0x27, 0xdc, 0xa0, 0xca, 0xfd, 0x02, 0xe1, 0xb9, 0xdb, 0xe8, 0xd1, 0x9a,
	0x19, 0xad, 0x11, 0xb7, 0x90, 0x1b, 0xc0, 0x56, 0xd9, 0x6b, 0x57, 0xd6, 0x2b, 0x36, 0xe1, 0x9b,
	0xdb, 0xa9, 0x85, 0x24, 0x42, 0xd1, 0xe8, 0xeb, 0xfb, 0xcd, 0xed, 0xae, 0x4b, 0xa4, 0x6e, 0xf0,
	0x09, 0xf4, 0xb7, 0xcb, 0x7e, 0xf0, 0xdb, 0xa9, 0x4d, 0xf8, 0x13, 0xe8, 0x4e, 0x64, 0x2e, 0xc9,
	0x47, 0x43, 0x11, 0x5e, 0x38, 0xb0, 0xb3, 0x16, 0x2f, 0x44, 0xf0, 0x02, 0x42, 0xd4, 0x2c, 0x46,
	0x81, 0x24, 0x13, 0xab, 0xf1, 0x8c, 0x25, 0xbc, 0xb0, 0xf3, 0x09, 0x7f, 0x74, 0x73, 0x7b, 0x80,
	0xfa, 0xa4, 0x69, 0x9f, 0x7a, 0x37, 0x18, 0x42, 0x0f, 0xb5, 0x1d, 0xda, 0x72, 0x55, 0x57, 0xbc,
	0x71, 0x15, 0xbb, 0xa8, 0x47, 0x56, 0xae, 0xc8, 0xe7, 0x7e, 0x0f, 0x6e, 0x92, 0x29, 0xe3, 0x5a,
	0x4b, 0x32, 0xba, 0xae, 0xf8, 0xc9, 0x55, 0xdc, 0x47, 0x7d, 0x62, 0xed, 0x63, 0xef, 0x56, 0x85,
	0x9f, 0xc2, 0x5d, 0xfb, 0x21, 0x56, 0x8a, 0xea, 0xb3, 0xbc, 0x74, 0x7c, 0x17, 0x45, 0x7c, 0xad,
	0xa8, 0x3a, 0x47, 0x1b, 0x45, 0xc4, 0x8c, 0x27, 0x99, 0xcd, 0xb3, 0x40, 0x21, 0x45, 0x78, 0x55,
	0x46, 0x11, 0xf1, 0x71, 0x92, 0x8d, 0x4a, 0x35, 0xd8, 0x07, 0x5b, 0xca, 0x66, 0x7c, 0x59, 0x7e,
	0x86, 0xae, 0xfd, 0x43, 0x41, 0x11, 0x5f, 0xf2, 0xa5, 0xff, 0xf8, 0x44, 0xb0, 0x67, 0x19, 0x33,
	0x25, 0xa9, 0xa7, 0x2a, 0x13, 0x25, 0x39, 0xf2, 0xb7, 0x0e, 0x45, 0xfc, 0x6b, 0xe5, 0x78, 0xde,
	0xfe, 0x23, 0x44, 0xcc, 0x92, 0x39, 0x91, 0xcc, 0x4d, 0x49, 0xff, 0xec, 0x68, 0x1b, 0xff, 0xd4,
	0xeb, 0x9e, 0x1d, 0x40, 0xa7, 0x4a, 0xea, 0xee, 0xfe, 0xd8, 0xdd, 0x7d, 0xf0, 0x29, 0xed, 0x9d,
	0x8f, 0x6f, 0xbb, 0xbf, 0xf1, 0xb3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xbe, 0x18, 0x14,
	0xa2, 0x07, 0x00, 0x00,
}
