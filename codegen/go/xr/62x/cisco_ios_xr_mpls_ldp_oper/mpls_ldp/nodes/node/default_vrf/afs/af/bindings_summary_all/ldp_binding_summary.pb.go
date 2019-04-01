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
// source: ldp_binding_summary.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_afs_af_bindings_summary_all

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

type LdpBindingSummary_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpBindingSummary_KEYS) Reset()         { *m = LdpBindingSummary_KEYS{} }
func (m *LdpBindingSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpBindingSummary_KEYS) ProtoMessage()    {}
func (*LdpBindingSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1f4b5951054fcc, []int{0}
}

func (m *LdpBindingSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBindingSummary_KEYS.Unmarshal(m, b)
}
func (m *LdpBindingSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBindingSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpBindingSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBindingSummary_KEYS.Merge(m, src)
}
func (m *LdpBindingSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpBindingSummary_KEYS.Size(m)
}
func (m *LdpBindingSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBindingSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBindingSummary_KEYS proto.InternalMessageInfo

func (m *LdpBindingSummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpBindingSummary_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type LdpVrfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpVrfInfo) Reset()         { *m = LdpVrfInfo{} }
func (m *LdpVrfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpVrfInfo) ProtoMessage()    {}
func (*LdpVrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1f4b5951054fcc, []int{1}
}

func (m *LdpVrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpVrfInfo.Unmarshal(m, b)
}
func (m *LdpVrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpVrfInfo.Marshal(b, m, deterministic)
}
func (m *LdpVrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpVrfInfo.Merge(m, src)
}
func (m *LdpVrfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpVrfInfo.Size(m)
}
func (m *LdpVrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpVrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpVrfInfo proto.InternalMessageInfo

func (m *LdpVrfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LdpVrfInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LdpBindingSummaryAf struct {
	AddressFamily             string   `protobuf:"bytes,1,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	LastLibUpdate             uint32   `protobuf:"varint,2,opt,name=last_lib_update,json=lastLibUpdate,proto3" json:"last_lib_update,omitempty"`
	LibMinimumRevisionSentAll uint32   `protobuf:"varint,3,opt,name=lib_minimum_revision_sent_all,json=libMinimumRevisionSentAll,proto3" json:"lib_minimum_revision_sent_all,omitempty"`
	BindingTotal              uint32   `protobuf:"varint,4,opt,name=binding_total,json=bindingTotal,proto3" json:"binding_total,omitempty"`
	BindingLocal              uint32   `protobuf:"varint,5,opt,name=binding_local,json=bindingLocal,proto3" json:"binding_local,omitempty"`
	BindingRemote             uint32   `protobuf:"varint,6,opt,name=binding_remote,json=bindingRemote,proto3" json:"binding_remote,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *LdpBindingSummaryAf) Reset()         { *m = LdpBindingSummaryAf{} }
func (m *LdpBindingSummaryAf) String() string { return proto.CompactTextString(m) }
func (*LdpBindingSummaryAf) ProtoMessage()    {}
func (*LdpBindingSummaryAf) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1f4b5951054fcc, []int{2}
}

func (m *LdpBindingSummaryAf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBindingSummaryAf.Unmarshal(m, b)
}
func (m *LdpBindingSummaryAf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBindingSummaryAf.Marshal(b, m, deterministic)
}
func (m *LdpBindingSummaryAf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBindingSummaryAf.Merge(m, src)
}
func (m *LdpBindingSummaryAf) XXX_Size() int {
	return xxx_messageInfo_LdpBindingSummaryAf.Size(m)
}
func (m *LdpBindingSummaryAf) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBindingSummaryAf.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBindingSummaryAf proto.InternalMessageInfo

func (m *LdpBindingSummaryAf) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpBindingSummaryAf) GetLastLibUpdate() uint32 {
	if m != nil {
		return m.LastLibUpdate
	}
	return 0
}

func (m *LdpBindingSummaryAf) GetLibMinimumRevisionSentAll() uint32 {
	if m != nil {
		return m.LibMinimumRevisionSentAll
	}
	return 0
}

func (m *LdpBindingSummaryAf) GetBindingTotal() uint32 {
	if m != nil {
		return m.BindingTotal
	}
	return 0
}

func (m *LdpBindingSummaryAf) GetBindingLocal() uint32 {
	if m != nil {
		return m.BindingLocal
	}
	return 0
}

func (m *LdpBindingSummaryAf) GetBindingRemote() uint32 {
	if m != nil {
		return m.BindingRemote
	}
	return 0
}

type LdpBindingSummary struct {
	Vrf                      *LdpVrfInfo            `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AddressFamily            string                 `protobuf:"bytes,51,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	BindAf                   []*LdpBindingSummaryAf `protobuf:"bytes,52,rep,name=bind_af,json=bindAf,proto3" json:"bind_af,omitempty"`
	BindingNoRoute           uint32                 `protobuf:"varint,53,opt,name=binding_no_route,json=bindingNoRoute,proto3" json:"binding_no_route,omitempty"`
	BindingLocalNoRoute      uint32                 `protobuf:"varint,54,opt,name=binding_local_no_route,json=bindingLocalNoRoute,proto3" json:"binding_local_no_route,omitempty"`
	BindingLocalNull         uint32                 `protobuf:"varint,55,opt,name=binding_local_null,json=bindingLocalNull,proto3" json:"binding_local_null,omitempty"`
	BindingLocalImplicitNull uint32                 `protobuf:"varint,56,opt,name=binding_local_implicit_null,json=bindingLocalImplicitNull,proto3" json:"binding_local_implicit_null,omitempty"`
	BindingLocalExplicitNull uint32                 `protobuf:"varint,57,opt,name=binding_local_explicit_null,json=bindingLocalExplicitNull,proto3" json:"binding_local_explicit_null,omitempty"`
	BindingLocalNonNull      uint32                 `protobuf:"varint,58,opt,name=binding_local_non_null,json=bindingLocalNonNull,proto3" json:"binding_local_non_null,omitempty"`
	BindingLocalOor          uint32                 `protobuf:"varint,59,opt,name=binding_local_oor,json=bindingLocalOor,proto3" json:"binding_local_oor,omitempty"`
	LowestAllocatedLabel     uint32                 `protobuf:"varint,60,opt,name=lowest_allocated_label,json=lowestAllocatedLabel,proto3" json:"lowest_allocated_label,omitempty"`
	HighestAllocatedLabel    uint32                 `protobuf:"varint,61,opt,name=highest_allocated_label,json=highestAllocatedLabel,proto3" json:"highest_allocated_label,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}               `json:"-"`
	XXX_unrecognized         []byte                 `json:"-"`
	XXX_sizecache            int32                  `json:"-"`
}

func (m *LdpBindingSummary) Reset()         { *m = LdpBindingSummary{} }
func (m *LdpBindingSummary) String() string { return proto.CompactTextString(m) }
func (*LdpBindingSummary) ProtoMessage()    {}
func (*LdpBindingSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_da1f4b5951054fcc, []int{3}
}

func (m *LdpBindingSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpBindingSummary.Unmarshal(m, b)
}
func (m *LdpBindingSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpBindingSummary.Marshal(b, m, deterministic)
}
func (m *LdpBindingSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpBindingSummary.Merge(m, src)
}
func (m *LdpBindingSummary) XXX_Size() int {
	return xxx_messageInfo_LdpBindingSummary.Size(m)
}
func (m *LdpBindingSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpBindingSummary.DiscardUnknown(m)
}

var xxx_messageInfo_LdpBindingSummary proto.InternalMessageInfo

func (m *LdpBindingSummary) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpBindingSummary) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpBindingSummary) GetBindAf() []*LdpBindingSummaryAf {
	if m != nil {
		return m.BindAf
	}
	return nil
}

func (m *LdpBindingSummary) GetBindingNoRoute() uint32 {
	if m != nil {
		return m.BindingNoRoute
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalNoRoute() uint32 {
	if m != nil {
		return m.BindingLocalNoRoute
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalNull() uint32 {
	if m != nil {
		return m.BindingLocalNull
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalImplicitNull() uint32 {
	if m != nil {
		return m.BindingLocalImplicitNull
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalExplicitNull() uint32 {
	if m != nil {
		return m.BindingLocalExplicitNull
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalNonNull() uint32 {
	if m != nil {
		return m.BindingLocalNonNull
	}
	return 0
}

func (m *LdpBindingSummary) GetBindingLocalOor() uint32 {
	if m != nil {
		return m.BindingLocalOor
	}
	return 0
}

func (m *LdpBindingSummary) GetLowestAllocatedLabel() uint32 {
	if m != nil {
		return m.LowestAllocatedLabel
	}
	return 0
}

func (m *LdpBindingSummary) GetHighestAllocatedLabel() uint32 {
	if m != nil {
		return m.HighestAllocatedLabel
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpBindingSummary_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.bindings_summary_all.ldp_binding_summary_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.bindings_summary_all.ldp_vrf_info")
	proto.RegisterType((*LdpBindingSummaryAf)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.bindings_summary_all.ldp_binding_summary_af")
	proto.RegisterType((*LdpBindingSummary)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.bindings_summary_all.ldp_binding_summary")
}

func init() { proto.RegisterFile("ldp_binding_summary.proto", fileDescriptor_da1f4b5951054fcc) }

var fileDescriptor_da1f4b5951054fcc = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x1c, 0xc5, 0xd5, 0x6e, 0x74, 0xcc, 0x5b, 0xb7, 0xe1, 0x41, 0xe7, 0x69, 0x42, 0xaa, 0x8a, 0x40,
	0x15, 0x42, 0xb9, 0x68, 0xc7, 0xf8, 0x9c, 0x44, 0x2f, 0x86, 0x84, 0x28, 0x03, 0x65, 0x80, 0xc4,
	0x95, 0xe5, 0x34, 0xf6, 0x66, 0xc9, 0x1f, 0x91, 0xed, 0x94, 0xee, 0x09, 0x78, 0x05, 0xc4, 0xd3,
	0x22, 0x3b, 0x89, 0x48, 0x58, 0x6e, 0x77, 0x53, 0xb5, 0xe7, 0x9c, 0xdf, 0x89, 0xed, 0xbf, 0x1b,
	0x70, 0x28, 0xd2, 0x0c, 0x27, 0x5c, 0xa5, 0x5c, 0x5d, 0x62, 0x9b, 0x4b, 0x49, 0xcc, 0x75, 0x94,
	0x19, 0xed, 0x34, 0xfc, 0xbe, 0xe0, 0x76, 0xa1, 0x31, 0xd7, 0x16, 0xaf, 0x0c, 0x96, 0x99, 0xb0,
	0xd8, 0x87, 0x75, 0x46, 0x4d, 0x54, 0xfd, 0x8a, 0x94, 0x4e, 0xa9, 0x0d, 0x9f, 0x51, 0x4a, 0x19,
	0xc9, 0x85, 0xc3, 0x4b, 0xc3, 0x22, 0xc2, 0x6c, 0x44, 0x58, 0x54, 0x16, 0xdb, 0xaa, 0x19, 0x13,
	0x21, 0x46, 0x5f, 0x00, 0x6a, 0x79, 0x28, 0xfe, 0x78, 0xf6, 0xe3, 0x02, 0x1e, 0x81, 0x4d, 0xdf,
	0x85, 0x15, 0x91, 0x14, 0x75, 0x86, 0x9d, 0xf1, 0x66, 0x7c, 0xd7, 0x0b, 0xe7, 0x44, 0x52, 0x78,
	0x00, 0x36, 0x08, 0x2b, 0xac, 0x6e, 0xb0, 0x7a, 0x84, 0x79, 0x63, 0x34, 0x01, 0xdb, 0xbe, 0x71,
	0x69, 0x18, 0xe6, 0x8a, 0x69, 0x08, 0xc1, 0x7a, 0xad, 0x20, 0x7c, 0x87, 0x3b, 0xa0, 0xcb, 0xd3,
	0xc0, 0xf5, 0xe3, 0x2e, 0x4f, 0x47, 0x7f, 0xba, 0x60, 0xd0, 0xb6, 0x0c, 0xc2, 0xe0, 0x63, 0xb0,
	0x43, 0xd2, 0xd4, 0x50, 0x6b, 0x31, 0x23, 0x92, 0x8b, 0xeb, 0xb2, 0xa8, 0x5f, 0xaa, 0xef, 0x83,
	0x08, 0x9f, 0x80, 0x5d, 0x41, 0xac, 0xc3, 0x82, 0x27, 0x38, 0xcf, 0x52, 0xe2, 0x68, 0x59, 0xdf,
	0xf7, 0xf2, 0x9c, 0x27, 0xdf, 0x82, 0x08, 0xdf, 0x81, 0x87, 0x3e, 0x22, 0xb9, 0xe2, 0x32, 0x97,
	0xd8, 0xd0, 0x25, 0xb7, 0x5c, 0x2b, 0x6c, 0xa9, 0x72, 0xfe, 0x40, 0xd0, 0x5a, 0xa0, 0x0e, 0x05,
	0x4f, 0x3e, 0x15, 0x99, 0xb8, 0x8c, 0x5c, 0x50, 0xe5, 0x66, 0x42, 0xc0, 0x47, 0xa0, 0x5f, 0x2d,
	0xd3, 0x69, 0x47, 0x04, 0x5a, 0x0f, 0xc4, 0x76, 0x29, 0x7e, 0xf5, 0x5a, 0x3d, 0x24, 0xf4, 0x82,
	0x08, 0x74, 0xa7, 0x11, 0x9a, 0x7b, 0xcd, 0x6f, 0xad, 0x0a, 0x19, 0x2a, 0xb5, 0xa3, 0xa8, 0x57,
	0x2c, 0xb9, 0x54, 0xe3, 0x20, 0x8e, 0x7e, 0xf7, 0xc0, 0x7e, 0xcb, 0xe1, 0xc0, 0x25, 0x58, 0x5b,
	0x1a, 0x86, 0x26, 0xc3, 0xce, 0x78, 0x6b, 0x92, 0x46, 0xb7, 0x73, 0x41, 0xa2, 0xfa, 0x2c, 0x63,
	0xff, 0xc0, 0x96, 0x89, 0x4c, 0xdb, 0x26, 0xf2, 0xab, 0x03, 0x36, 0x7c, 0x23, 0x26, 0x0c, 0x1d,
	0x0f, 0xd7, 0xc6, 0x5b, 0x13, 0x75, 0x9b, 0x6b, 0xbc, 0x79, 0x75, 0xe2, 0x9e, 0xd7, 0x66, 0x0c,
	0x8e, 0xc1, 0x5e, 0xe5, 0x2a, 0x8d, 0x8d, 0xce, 0x1d, 0x45, 0xcf, 0xc3, 0x49, 0x57, 0xe7, 0x7f,
	0xae, 0x63, 0xaf, 0xc2, 0x29, 0x18, 0x34, 0xc6, 0xf6, 0x2f, 0x7f, 0x12, 0xf2, 0xfb, 0xf5, 0xf9,
	0x55, 0xd0, 0x33, 0x00, 0xff, 0x83, 0x72, 0x21, 0xd0, 0x8b, 0x00, 0xec, 0x35, 0x80, 0x5c, 0x08,
	0x78, 0x0a, 0x8e, 0x9a, 0x69, 0x2e, 0x33, 0xc1, 0x17, 0xdc, 0x15, 0xd8, 0xcb, 0x80, 0xa1, 0x3a,
	0xf6, 0xa1, 0x0c, 0xb4, 0xe3, 0x74, 0x55, 0xc7, 0x5f, 0xdd, 0xc4, 0xcf, 0x56, 0x35, 0xbc, 0x65,
	0x83, 0xaa, 0x20, 0x5f, 0xb7, 0x6d, 0x50, 0x05, 0xe8, 0x29, 0xb8, 0xd7, 0x84, 0xb4, 0x36, 0xe8,
	0x4d, 0xc8, 0xef, 0xd6, 0xf3, 0x9f, 0xb5, 0x81, 0xc7, 0x60, 0x20, 0xf4, 0x4f, 0x6a, 0xc3, 0x9f,
	0x49, 0x2f, 0x88, 0xa3, 0x29, 0x16, 0x24, 0xa1, 0x02, 0xbd, 0x0d, 0xc0, 0xfd, 0xc2, 0x9d, 0x55,
	0xe6, 0xdc, 0x7b, 0xf0, 0x04, 0x1c, 0x5c, 0xf1, 0xcb, 0xab, 0x36, 0xec, 0x34, 0x60, 0x0f, 0x4a,
	0xbb, 0xc9, 0x25, 0xbd, 0xf0, 0x72, 0x9c, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x22, 0x0f, 0x72,
	0x0c, 0x39, 0x05, 0x00, 0x00,
}
