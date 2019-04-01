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
// source: fib_oc_aft_prefix.proto

package cisco_ios_xr_fib_common_oper_oc_aft_l3_vrfs_vrf_abstract_forwarding_tables_ipv4_unicast_prefix_entries_prefix_entry

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

type FibOcAftPrefix_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	NetworkAddress       string   `protobuf:"bytes,2,opt,name=network_address,json=networkAddress,proto3" json:"network_address,omitempty"`
	MaskLength           uint32   `protobuf:"varint,3,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcAftPrefix_KEYS) Reset()         { *m = FibOcAftPrefix_KEYS{} }
func (m *FibOcAftPrefix_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPrefix_KEYS) ProtoMessage()    {}
func (*FibOcAftPrefix_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{0}
}

func (m *FibOcAftPrefix_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPrefix_KEYS.Unmarshal(m, b)
}
func (m *FibOcAftPrefix_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPrefix_KEYS.Marshal(b, m, deterministic)
}
func (m *FibOcAftPrefix_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPrefix_KEYS.Merge(m, src)
}
func (m *FibOcAftPrefix_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPrefix_KEYS.Size(m)
}
func (m *FibOcAftPrefix_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPrefix_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPrefix_KEYS proto.InternalMessageInfo

func (m *FibOcAftPrefix_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibOcAftPrefix_KEYS) GetNetworkAddress() string {
	if m != nil {
		return m.NetworkAddress
	}
	return ""
}

func (m *FibOcAftPrefix_KEYS) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

type FibOcPrefixState struct {
	PrefixIndex          string   `protobuf:"bytes,1,opt,name=prefix_index,json=prefixIndex,proto3" json:"prefix_index,omitempty"`
	LabelIndex           string   `protobuf:"bytes,2,opt,name=label_index,json=labelIndex,proto3" json:"label_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcPrefixState) Reset()         { *m = FibOcPrefixState{} }
func (m *FibOcPrefixState) String() string { return proto.CompactTextString(m) }
func (*FibOcPrefixState) ProtoMessage()    {}
func (*FibOcPrefixState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{1}
}

func (m *FibOcPrefixState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcPrefixState.Unmarshal(m, b)
}
func (m *FibOcPrefixState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcPrefixState.Marshal(b, m, deterministic)
}
func (m *FibOcPrefixState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcPrefixState.Merge(m, src)
}
func (m *FibOcPrefixState) XXX_Size() int {
	return xxx_messageInfo_FibOcPrefixState.Size(m)
}
func (m *FibOcPrefixState) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcPrefixState.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcPrefixState proto.InternalMessageInfo

func (m *FibOcPrefixState) GetPrefixIndex() string {
	if m != nil {
		return m.PrefixIndex
	}
	return ""
}

func (m *FibOcPrefixState) GetLabelIndex() string {
	if m != nil {
		return m.LabelIndex
	}
	return ""
}

type FibOcAftPathState struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	IpAddress            string   `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PoppedMplsLabelStack []string `protobuf:"bytes,4,rep,name=popped_mpls_label_stack,json=poppedMplsLabelStack,proto3" json:"popped_mpls_label_stack,omitempty"`
	PushedMplsLabelStack []string `protobuf:"bytes,5,rep,name=pushed_mpls_label_stack,json=pushedMplsLabelStack,proto3" json:"pushed_mpls_label_stack,omitempty"`
	NetworkInstance      string   `protobuf:"bytes,6,opt,name=network_instance,json=networkInstance,proto3" json:"network_instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcAftPathState) Reset()         { *m = FibOcAftPathState{} }
func (m *FibOcAftPathState) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPathState) ProtoMessage()    {}
func (*FibOcAftPathState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{2}
}

func (m *FibOcAftPathState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPathState.Unmarshal(m, b)
}
func (m *FibOcAftPathState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPathState.Marshal(b, m, deterministic)
}
func (m *FibOcAftPathState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPathState.Merge(m, src)
}
func (m *FibOcAftPathState) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPathState.Size(m)
}
func (m *FibOcAftPathState) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPathState.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPathState proto.InternalMessageInfo

func (m *FibOcAftPathState) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *FibOcAftPathState) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *FibOcAftPathState) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *FibOcAftPathState) GetPoppedMplsLabelStack() []string {
	if m != nil {
		return m.PoppedMplsLabelStack
	}
	return nil
}

func (m *FibOcAftPathState) GetPushedMplsLabelStack() []string {
	if m != nil {
		return m.PushedMplsLabelStack
	}
	return nil
}

func (m *FibOcAftPathState) GetNetworkInstance() string {
	if m != nil {
		return m.NetworkInstance
	}
	return ""
}

type FibOcInterfRefState struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	Subinterface         uint32   `protobuf:"varint,2,opt,name=subinterface,proto3" json:"subinterface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcInterfRefState) Reset()         { *m = FibOcInterfRefState{} }
func (m *FibOcInterfRefState) String() string { return proto.CompactTextString(m) }
func (*FibOcInterfRefState) ProtoMessage()    {}
func (*FibOcInterfRefState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{3}
}

func (m *FibOcInterfRefState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcInterfRefState.Unmarshal(m, b)
}
func (m *FibOcInterfRefState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcInterfRefState.Marshal(b, m, deterministic)
}
func (m *FibOcInterfRefState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcInterfRefState.Merge(m, src)
}
func (m *FibOcInterfRefState) XXX_Size() int {
	return xxx_messageInfo_FibOcInterfRefState.Size(m)
}
func (m *FibOcInterfRefState) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcInterfRefState.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcInterfRefState proto.InternalMessageInfo

func (m *FibOcInterfRefState) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *FibOcInterfRefState) GetSubinterface() uint32 {
	if m != nil {
		return m.Subinterface
	}
	return 0
}

type FibOcInterfRef struct {
	State                *FibOcInterfRefState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FibOcInterfRef) Reset()         { *m = FibOcInterfRef{} }
func (m *FibOcInterfRef) String() string { return proto.CompactTextString(m) }
func (*FibOcInterfRef) ProtoMessage()    {}
func (*FibOcInterfRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{4}
}

func (m *FibOcInterfRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcInterfRef.Unmarshal(m, b)
}
func (m *FibOcInterfRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcInterfRef.Marshal(b, m, deterministic)
}
func (m *FibOcInterfRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcInterfRef.Merge(m, src)
}
func (m *FibOcInterfRef) XXX_Size() int {
	return xxx_messageInfo_FibOcInterfRef.Size(m)
}
func (m *FibOcInterfRef) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcInterfRef.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcInterfRef proto.InternalMessageInfo

func (m *FibOcInterfRef) GetState() *FibOcInterfRefState {
	if m != nil {
		return m.State
	}
	return nil
}

type FibOcAftPath struct {
	State                *FibOcAftPathState `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	InterfaceRef         *FibOcInterfRef    `protobuf:"bytes,2,opt,name=interface_ref,json=interfaceRef,proto3" json:"interface_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FibOcAftPath) Reset()         { *m = FibOcAftPath{} }
func (m *FibOcAftPath) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPath) ProtoMessage()    {}
func (*FibOcAftPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{5}
}

func (m *FibOcAftPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPath.Unmarshal(m, b)
}
func (m *FibOcAftPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPath.Marshal(b, m, deterministic)
}
func (m *FibOcAftPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPath.Merge(m, src)
}
func (m *FibOcAftPath) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPath.Size(m)
}
func (m *FibOcAftPath) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPath.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPath proto.InternalMessageInfo

func (m *FibOcAftPath) GetState() *FibOcAftPathState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *FibOcAftPath) GetInterfaceRef() *FibOcInterfRef {
	if m != nil {
		return m.InterfaceRef
	}
	return nil
}

type FibOcAftPrefix struct {
	State                *FibOcPrefixState `protobuf:"bytes,50,opt,name=state,proto3" json:"state,omitempty"`
	NextHop              []*FibOcAftPath   `protobuf:"bytes,51,rep,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FibOcAftPrefix) Reset()         { *m = FibOcAftPrefix{} }
func (m *FibOcAftPrefix) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPrefix) ProtoMessage()    {}
func (*FibOcAftPrefix) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0240fb97cee26e2, []int{6}
}

func (m *FibOcAftPrefix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPrefix.Unmarshal(m, b)
}
func (m *FibOcAftPrefix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPrefix.Marshal(b, m, deterministic)
}
func (m *FibOcAftPrefix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPrefix.Merge(m, src)
}
func (m *FibOcAftPrefix) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPrefix.Size(m)
}
func (m *FibOcAftPrefix) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPrefix.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPrefix proto.InternalMessageInfo

func (m *FibOcAftPrefix) GetState() *FibOcPrefixState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *FibOcAftPrefix) GetNextHop() []*FibOcAftPath {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func init() {
	proto.RegisterType((*FibOcAftPrefix_KEYS)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_aft_prefix_KEYS")
	proto.RegisterType((*FibOcPrefixState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_prefix_state")
	proto.RegisterType((*FibOcAftPathState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_aft_path_state")
	proto.RegisterType((*FibOcInterfRefState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_interf_ref_state")
	proto.RegisterType((*FibOcInterfRef)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_interf_ref")
	proto.RegisterType((*FibOcAftPath)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_aft_path")
	proto.RegisterType((*FibOcAftPrefix)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefix_entries.prefix_entry.fib_oc_aft_prefix")
}

func init() { proto.RegisterFile("fib_oc_aft_prefix.proto", fileDescriptor_e0240fb97cee26e2) }

var fileDescriptor_e0240fb97cee26e2 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x96, 0x93, 0x26, 0x6d, 0x5e, 0x12, 0x0a, 0x43, 0x69, 0x8d, 0x04, 0x22, 0x78, 0x43, 0xd8,
	0x78, 0x91, 0xc0, 0x01, 0x58, 0x20, 0x51, 0x51, 0x58, 0xb8, 0xab, 0x8a, 0xc5, 0x68, 0xec, 0x8c,
	0x93, 0x51, 0xec, 0x99, 0xd1, 0xcc, 0xe4, 0x87, 0x05, 0x3b, 0x84, 0x10, 0xe2, 0x0e, 0x80, 0xd8,
	0x72, 0x01, 0x6e, 0xc2, 0x01, 0xb8, 0x06, 0x12, 0xf2, 0x78, 0xe2, 0xa4, 0x29, 0xac, 0xd3, 0x8d,
	0xa5, 0xf7, 0xbd, 0x9f, 0xef, 0x7b, 0x3f, 0x1a, 0xc3, 0x49, 0xca, 0x62, 0x2c, 0x12, 0x4c, 0x52,
	0x83, 0xa5, 0xa2, 0x29, 0x5b, 0x86, 0x52, 0x09, 0x23, 0x90, 0x4e, 0x98, 0x4e, 0x04, 0x66, 0x42,
	0xe3, 0xa5, 0xc2, 0x45, 0x54, 0x22, 0xf2, 0x5c, 0x70, 0x2c, 0x24, 0x55, 0xa1, 0xcb, 0xc8, 0x86,
	0xe1, 0x5c, 0xa5, 0xba, 0xf8, 0x84, 0x24, 0xd6, 0x46, 0x91, 0xc4, 0xe0, 0x54, 0xa8, 0x05, 0x51,
	0x23, 0xc6, 0xc7, 0xd8, 0x90, 0x38, 0xa3, 0x3a, 0x64, 0x72, 0xfe, 0x04, 0xcf, 0x38, 0x4b, 0x88,
	0x36, 0x61, 0xc9, 0x82, 0x29, 0x37, 0x8a, 0x51, 0xbd, 0x69, 0xbe, 0x0d, 0xde, 0xc1, 0xf1, 0x15,
	0x3d, 0xf8, 0xe5, 0xf3, 0x8b, 0x73, 0x74, 0x17, 0x0e, 0xe6, 0x2a, 0xc5, 0x9c, 0xe4, 0xd4, 0xf7,
	0x7a, 0x5e, 0xbf, 0x15, 0xed, 0xcf, 0x55, 0xfa, 0x9a, 0xe4, 0x14, 0x3d, 0x82, 0x43, 0x4e, 0xcd,
	0x42, 0xa8, 0x29, 0x26, 0xa3, 0x91, 0xa2, 0x5a, 0xfb, 0x35, 0x1b, 0x71, 0xc3, 0xc1, 0xcf, 0x4a,
	0x14, 0x3d, 0x80, 0x76, 0x4e, 0xf4, 0x14, 0x67, 0x94, 0x8f, 0xcd, 0xc4, 0xaf, 0xf7, 0xbc, 0x7e,
	0x37, 0x82, 0x02, 0x3a, 0xb3, 0x48, 0x70, 0x01, 0xb7, 0x1d, 0xbd, 0xa3, 0xd6, 0x86, 0x18, 0x8a,
	0x1e, 0x42, 0xc7, 0xd9, 0x8c, 0x8f, 0xe8, 0xd2, 0xf1, 0xb7, 0x4b, 0xec, 0xb4, 0x80, 0x8a, 0xd2,
	0x19, 0x89, 0x69, 0xe6, 0x22, 0x4a, 0x7e, 0xb0, 0x90, 0x0d, 0x08, 0xfe, 0x78, 0x70, 0x67, 0xb3,
	0x35, 0x62, 0x26, 0xae, 0xfa, 0x11, 0x34, 0xd6, 0x65, 0xf7, 0xa2, 0xd2, 0x40, 0xc7, 0xd0, 0x5c,
	0x50, 0x36, 0x9e, 0x18, 0x5b, 0xab, 0x1b, 0x39, 0x0b, 0xdd, 0x07, 0x60, 0xb2, 0xea, 0xb3, 0x6e,
	0x79, 0x5a, 0x4c, 0xae, 0x5a, 0x7c, 0x0a, 0x27, 0x52, 0x48, 0x49, 0x47, 0x38, 0x97, 0x99, 0xc6,
	0xa5, 0x26, 0x6d, 0x48, 0x32, 0xf5, 0xf7, 0x7a, 0xf5, 0x7e, 0x2b, 0x3a, 0x2a, 0xdd, 0xaf, 0x64,
	0xa6, 0xcf, 0x0a, 0xe7, 0x79, 0xe1, 0xb3, 0x69, 0x33, 0x3d, 0xf9, 0x57, 0x5a, 0xc3, 0xa5, 0x59,
	0xf7, 0x56, 0xda, 0x63, 0xb8, 0xb9, 0x9a, 0x3c, 0xe3, 0xda, 0x10, 0x9e, 0x50, 0xbf, 0x69, 0x25,
	0xad, 0x36, 0x72, 0xea, 0xe0, 0xe0, 0x4d, 0x75, 0x69, 0x8c, 0x1b, 0xaa, 0x52, 0xac, 0x68, 0xea,
	0x06, 0x70, 0x0f, 0x5a, 0x25, 0x46, 0x92, 0xd5, 0x6e, 0xd7, 0x00, 0x0a, 0xa0, 0xa3, 0x67, 0xf1,
	0x3a, 0xa0, 0x1c, 0xc7, 0x25, 0x2c, 0xf8, 0xe9, 0xc1, 0xad, 0x2b, 0xd5, 0xd1, 0x77, 0x0f, 0x1a,
	0x96, 0xc1, 0x16, 0x6d, 0x0f, 0x3e, 0x7b, 0xe1, 0x0e, 0x6e, 0x3a, 0xfc, 0x4f, 0xdb, 0x51, 0xa9,
	0x2d, 0xf8, 0x5d, 0x83, 0xc3, 0xad, 0xc3, 0x40, 0xdf, 0xb6, 0x94, 0x7f, 0xda, 0xa9, 0xf2, 0xcb,
	0xf7, 0xea, 0x74, 0xa3, 0x1f, 0x1e, 0x74, 0xab, 0x0d, 0x14, 0x6d, 0xd9, 0xcd, 0xb4, 0x07, 0x1f,
	0xae, 0xc9, 0x94, 0xa3, 0x4e, 0xa5, 0x2e, 0xa2, 0x69, 0xf0, 0xab, 0x56, 0x9d, 0xc8, 0xfa, 0x69,
	0x41, 0x5f, 0xaa, 0x41, 0x0f, 0xac, 0xf8, 0x8f, 0x3b, 0x15, 0xbf, 0xf9, 0xe8, 0xac, 0xc6, 0xfc,
	0xd5, 0x83, 0x03, 0x4e, 0x97, 0x06, 0x4f, 0x84, 0xf4, 0x87, 0xbd, 0x7a, 0xbf, 0x3d, 0x78, 0x7f,
	0x2d, 0xae, 0x21, 0xda, 0x2f, 0x64, 0xbd, 0x10, 0x32, 0x6e, 0xda, 0x1f, 0xc6, 0xf0, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0xae, 0x52, 0xde, 0x4b, 0x06, 0x00, 0x00,
}