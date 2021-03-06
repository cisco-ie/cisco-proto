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
// source: fib_oc_aft_path.proto

package cisco_ios_xr_fib_common_oper_oc_aft_l3_vrfs_vrf_abstract_forwarding_tables_ipv4_unicast_prefixes_prefix_next_hops_next_hop

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

type FibOcAftPath_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Index                uint32   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibOcAftPath_KEYS) Reset()         { *m = FibOcAftPath_KEYS{} }
func (m *FibOcAftPath_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPath_KEYS) ProtoMessage()    {}
func (*FibOcAftPath_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f730bb0e92074c59, []int{0}
}

func (m *FibOcAftPath_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibOcAftPath_KEYS.Unmarshal(m, b)
}
func (m *FibOcAftPath_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibOcAftPath_KEYS.Marshal(b, m, deterministic)
}
func (m *FibOcAftPath_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibOcAftPath_KEYS.Merge(m, src)
}
func (m *FibOcAftPath_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibOcAftPath_KEYS.Size(m)
}
func (m *FibOcAftPath_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibOcAftPath_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibOcAftPath_KEYS proto.InternalMessageInfo

func (m *FibOcAftPath_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FibOcAftPath_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *FibOcAftPath_KEYS) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
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
	return fileDescriptor_f730bb0e92074c59, []int{1}
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
	return fileDescriptor_f730bb0e92074c59, []int{2}
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
	return fileDescriptor_f730bb0e92074c59, []int{3}
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
	State                *FibOcAftPathState `protobuf:"bytes,50,opt,name=state,proto3" json:"state,omitempty"`
	InterfaceRef         *FibOcInterfRef    `protobuf:"bytes,51,opt,name=interface_ref,json=interfaceRef,proto3" json:"interface_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *FibOcAftPath) Reset()         { *m = FibOcAftPath{} }
func (m *FibOcAftPath) String() string { return proto.CompactTextString(m) }
func (*FibOcAftPath) ProtoMessage()    {}
func (*FibOcAftPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_f730bb0e92074c59, []int{4}
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

func init() {
	proto.RegisterType((*FibOcAftPath_KEYS)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefixes.prefix.next_hops.next_hop.fib_oc_aft_path_KEYS")
	proto.RegisterType((*FibOcAftPathState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefixes.prefix.next_hops.next_hop.fib_oc_aft_path_state")
	proto.RegisterType((*FibOcInterfRefState)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefixes.prefix.next_hops.next_hop.fib_oc_interf_ref_state")
	proto.RegisterType((*FibOcInterfRef)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefixes.prefix.next_hops.next_hop.fib_oc_interf_ref")
	proto.RegisterType((*FibOcAftPath)(nil), "cisco_ios_xr_fib_common_oper.oc_aft_l3.vrfs.vrf.abstract_forwarding_tables.ipv4_unicast.prefixes.prefix.next_hops.next_hop.fib_oc_aft_path")
}

func init() { proto.RegisterFile("fib_oc_aft_path.proto", fileDescriptor_f730bb0e92074c59) }

var fileDescriptor_f730bb0e92074c59 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xb6, 0x4d, 0x20, 0xa6, 0x51, 0xc1, 0x0a, 0x74, 0x91, 0x40, 0x8a, 0xf6, 0x14, 0x2e,
	0x3e, 0x34, 0xf0, 0x00, 0x1c, 0x38, 0x20, 0x7e, 0x0e, 0xdb, 0x13, 0xe2, 0x30, 0xf2, 0x7a, 0xc7,
	0x8d, 0xd5, 0x5d, 0xdb, 0xb2, 0x9d, 0x1f, 0xf1, 0x0e, 0x1c, 0x10, 0x2f, 0x81, 0xb8, 0xf2, 0x02,
	0x3c, 0x13, 0x77, 0x24, 0xb4, 0x5e, 0x93, 0xa8, 0x29, 0x9c, 0xdb, 0x4b, 0xe4, 0xf9, 0x66, 0xe6,
	0x9b, 0xef, 0x9b, 0x8c, 0x96, 0x3c, 0x94, 0xaa, 0x02, 0x23, 0x80, 0xcb, 0x00, 0x96, 0x87, 0x05,
	0xb3, 0xce, 0x04, 0x43, 0x3f, 0x09, 0xe5, 0x85, 0x01, 0x65, 0x3c, 0x6c, 0x1c, 0x74, 0x35, 0xc2,
	0xb4, 0xad, 0xd1, 0x60, 0x2c, 0x3a, 0x96, 0xea, 0x9b, 0x39, 0x5b, 0x39, 0xe9, 0xbb, 0x1f, 0xc6,
	0x2b, 0x1f, 0x1c, 0x17, 0x01, 0xa4, 0x71, 0x6b, 0xee, 0x6a, 0xa5, 0x2f, 0x20, 0xf0, 0xaa, 0x41,
	0xcf, 0x94, 0x5d, 0x3d, 0x87, 0xa5, 0x56, 0x82, 0xfb, 0xc0, 0xac, 0x43, 0xa9, 0x36, 0xe8, 0xd3,
	0x83, 0x69, 0xdc, 0x04, 0x58, 0x18, 0xeb, 0xb7, 0xaf, 0x02, 0xc8, 0x64, 0x4f, 0x14, 0xbc, 0x79,
	0xf5, 0xe1, 0x9c, 0x3e, 0x26, 0x77, 0x57, 0x4e, 0x82, 0xe6, 0x2d, 0xe6, 0xd9, 0x34, 0x9b, 0x8d,
	0xca, 0x3b, 0x2b, 0x27, 0xdf, 0xf3, 0x16, 0xe9, 0x23, 0x32, 0xec, 0xf9, 0xf2, 0x83, 0x98, 0x48,
	0x11, 0x9d, 0x90, 0x81, 0xd2, 0x35, 0x6e, 0xf2, 0xc3, 0x69, 0x36, 0x1b, 0x97, 0x7d, 0x50, 0xfc,
	0xce, 0xae, 0xd9, 0x06, 0x1f, 0x78, 0xc0, 0x5d, 0x7d, 0xc7, 0x7f, 0x94, 0xea, 0x3b, 0xf6, 0x35,
	0xaa, 0x8b, 0x45, 0x88, 0xec, 0xe3, 0x32, 0x45, 0xf4, 0x29, 0x21, 0xca, 0x02, 0xaf, 0x6b, 0x87,
	0xde, 0xc7, 0x11, 0xa3, 0x72, 0xa4, 0xec, 0xcb, 0x1e, 0xa0, 0x2f, 0xc8, 0xa9, 0x35, 0xd6, 0x62,
	0x0d, 0xad, 0x6d, 0x3c, 0x34, 0xbc, 0xc2, 0xa6, 0x9b, 0x23, 0x2e, 0xf3, 0xa3, 0xe9, 0xe1, 0x6c,
	0x54, 0x4e, 0xfa, 0xf4, 0x3b, 0xdb, 0xf8, 0xb7, 0x5d, 0xf2, 0xbc, 0xcb, 0xc5, 0xb6, 0xa5, 0x5f,
	0xfc, 0xab, 0x6d, 0x90, 0xda, 0x62, 0x7a, 0xaf, 0xed, 0x19, 0xb9, 0xaf, 0x31, 0xac, 0x8d, 0xbb,
	0x04, 0xa5, 0x7d, 0xe0, 0x5a, 0x60, 0x3e, 0x8c, 0x92, 0x4e, 0x12, 0xfe, 0x3a, 0xc1, 0xc5, 0x47,
	0x72, 0x9a, 0xec, 0x2b, 0x1d, 0xd0, 0x49, 0x70, 0x28, 0xd3, 0x02, 0x9e, 0x90, 0x51, 0x8f, 0x71,
	0xf1, 0x77, 0xc9, 0x3b, 0x80, 0x16, 0xe4, 0xd8, 0x2f, 0xab, 0x5d, 0x41, 0xbf, 0x8e, 0x2b, 0x58,
	0xf1, 0x33, 0x23, 0x0f, 0xae, 0xb1, 0xd3, 0xef, 0x19, 0x19, 0xc4, 0x09, 0x91, 0xf4, 0xde, 0xd9,
	0xd7, 0x8c, 0xdd, 0xdc, 0x85, 0xb1, 0xff, 0xb8, 0x2f, 0x7b, 0x89, 0xc5, 0xaf, 0x03, 0x72, 0xb2,
	0x77, 0x1f, 0xf4, 0xdb, 0xd6, 0xc0, 0x59, 0x34, 0xf0, 0xe5, 0x36, 0x18, 0xb8, 0x7a, 0xbd, 0x49,
	0x3e, 0xfd, 0x91, 0x91, 0xf1, 0xf6, 0xff, 0xe8, 0xdc, 0xe5, 0xf3, 0x28, 0xf9, 0xf3, 0xed, 0xda,
	0x79, 0x79, 0xbc, 0x15, 0x59, 0xa2, 0xac, 0x86, 0xf1, 0xc3, 0x33, 0xff, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x41, 0x81, 0x94, 0x8b, 0x91, 0x04, 0x00, 0x00,
}
