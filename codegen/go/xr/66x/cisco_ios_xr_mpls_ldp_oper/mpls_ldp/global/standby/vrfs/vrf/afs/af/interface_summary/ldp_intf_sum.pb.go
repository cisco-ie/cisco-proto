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
// source: ldp_intf_sum.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_afs_af_interface_summary

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

type LdpIntfSum_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIntfSum_KEYS) Reset()         { *m = LdpIntfSum_KEYS{} }
func (m *LdpIntfSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpIntfSum_KEYS) ProtoMessage()    {}
func (*LdpIntfSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b93a20219fd1fd8e, []int{0}
}

func (m *LdpIntfSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntfSum_KEYS.Unmarshal(m, b)
}
func (m *LdpIntfSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntfSum_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpIntfSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntfSum_KEYS.Merge(m, src)
}
func (m *LdpIntfSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpIntfSum_KEYS.Size(m)
}
func (m *LdpIntfSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntfSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntfSum_KEYS proto.InternalMessageInfo

func (m *LdpIntfSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpIntfSum_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type LdpIntfSum struct {
	KnownIpInterfaceCount                uint32   `protobuf:"varint,50,opt,name=known_ip_interface_count,json=knownIpInterfaceCount,proto3" json:"known_ip_interface_count,omitempty"`
	KnownIpInterfaceLdpEnabled           uint32   `protobuf:"varint,51,opt,name=known_ip_interface_ldp_enabled,json=knownIpInterfaceLdpEnabled,proto3" json:"known_ip_interface_ldp_enabled,omitempty"`
	LdpConfiguredAttachedInterface       uint32   `protobuf:"varint,52,opt,name=ldp_configured_attached_interface,json=ldpConfiguredAttachedInterface,proto3" json:"ldp_configured_attached_interface,omitempty"`
	LdpConfiguredTeInterface             uint32   `protobuf:"varint,53,opt,name=ldp_configured_te_interface,json=ldpConfiguredTeInterface,proto3" json:"ldp_configured_te_interface,omitempty"`
	ForwardReferences                    uint32   `protobuf:"varint,54,opt,name=forward_references,json=forwardReferences,proto3" json:"forward_references,omitempty"`
	AutoConfigDisabled                   uint32   `protobuf:"varint,55,opt,name=auto_config_disabled,json=autoConfigDisabled,proto3" json:"auto_config_disabled,omitempty"`
	AutoConfig                           uint32   `protobuf:"varint,56,opt,name=auto_config,json=autoConfig,proto3" json:"auto_config,omitempty"`
	AutoConfigForwardReferenceInterfaces uint32   `protobuf:"varint,57,opt,name=auto_config_forward_reference_interfaces,json=autoConfigForwardReferenceInterfaces,proto3" json:"auto_config_forward_reference_interfaces,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{} `json:"-"`
	XXX_unrecognized                     []byte   `json:"-"`
	XXX_sizecache                        int32    `json:"-"`
}

func (m *LdpIntfSum) Reset()         { *m = LdpIntfSum{} }
func (m *LdpIntfSum) String() string { return proto.CompactTextString(m) }
func (*LdpIntfSum) ProtoMessage()    {}
func (*LdpIntfSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b93a20219fd1fd8e, []int{1}
}

func (m *LdpIntfSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIntfSum.Unmarshal(m, b)
}
func (m *LdpIntfSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIntfSum.Marshal(b, m, deterministic)
}
func (m *LdpIntfSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIntfSum.Merge(m, src)
}
func (m *LdpIntfSum) XXX_Size() int {
	return xxx_messageInfo_LdpIntfSum.Size(m)
}
func (m *LdpIntfSum) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIntfSum.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIntfSum proto.InternalMessageInfo

func (m *LdpIntfSum) GetKnownIpInterfaceCount() uint32 {
	if m != nil {
		return m.KnownIpInterfaceCount
	}
	return 0
}

func (m *LdpIntfSum) GetKnownIpInterfaceLdpEnabled() uint32 {
	if m != nil {
		return m.KnownIpInterfaceLdpEnabled
	}
	return 0
}

func (m *LdpIntfSum) GetLdpConfiguredAttachedInterface() uint32 {
	if m != nil {
		return m.LdpConfiguredAttachedInterface
	}
	return 0
}

func (m *LdpIntfSum) GetLdpConfiguredTeInterface() uint32 {
	if m != nil {
		return m.LdpConfiguredTeInterface
	}
	return 0
}

func (m *LdpIntfSum) GetForwardReferences() uint32 {
	if m != nil {
		return m.ForwardReferences
	}
	return 0
}

func (m *LdpIntfSum) GetAutoConfigDisabled() uint32 {
	if m != nil {
		return m.AutoConfigDisabled
	}
	return 0
}

func (m *LdpIntfSum) GetAutoConfig() uint32 {
	if m != nil {
		return m.AutoConfig
	}
	return 0
}

func (m *LdpIntfSum) GetAutoConfigForwardReferenceInterfaces() uint32 {
	if m != nil {
		return m.AutoConfigForwardReferenceInterfaces
	}
	return 0
}

func init() {
	proto.RegisterType((*LdpIntfSum_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.interface_summary.ldp_intf_sum_KEYS")
	proto.RegisterType((*LdpIntfSum)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.interface_summary.ldp_intf_sum")
}

func init() { proto.RegisterFile("ldp_intf_sum.proto", fileDescriptor_b93a20219fd1fd8e) }

var fileDescriptor_b93a20219fd1fd8e = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x99, 0xc2, 0xa6, 0x51, 0x0f, 0x0b, 0x8a, 0x51, 0x61, 0xce, 0xe1, 0x61, 0x17, 0x83,
	0x38, 0x75, 0x7a, 0xf0, 0xa0, 0x73, 0xca, 0x50, 0x3c, 0xcc, 0x21, 0x78, 0x0a, 0x69, 0x93, 0xcc,
	0x62, 0x9b, 0x94, 0x24, 0xdd, 0xdc, 0xaf, 0xf6, 0x2f, 0x48, 0xb3, 0xae, 0xad, 0xd3, 0x4b, 0xa1,
	0x7d, 0xdf, 0xe7, 0xf9, 0xbe, 0x0f, 0x0a, 0x60, 0xc8, 0x62, 0x12, 0x48, 0x2b, 0x88, 0x49, 0x22,
	0x1c, 0x6b, 0x65, 0x15, 0x1c, 0xf9, 0x81, 0xf1, 0x15, 0x09, 0x94, 0x21, 0x5f, 0x9a, 0x44, 0x71,
	0x68, 0x48, 0xda, 0x52, 0x31, 0xd7, 0x78, 0xf1, 0x86, 0xc7, 0xa1, 0xf2, 0x68, 0x88, 0x8d, 0xa5,
	0x92, 0x79, 0x33, 0x3c, 0xd1, 0xc2, 0xa4, 0x0f, 0x4c, 0x85, 0xc1, 0x54, 0xe0, 0x40, 0x5a, 0xae,
	0x05, 0xf5, 0x79, 0xea, 0x8d, 0xa8, 0x9e, 0xb5, 0x1e, 0x41, 0xbd, 0x3c, 0x8b, 0x3c, 0xf5, 0xdf,
	0x5f, 0xe1, 0x1e, 0x58, 0x9b, 0x68, 0x41, 0x24, 0x8d, 0x38, 0xaa, 0x34, 0x2b, 0xed, 0xf5, 0x61,
	0x6d, 0xa2, 0xc5, 0x0b, 0x8d, 0x38, 0xdc, 0x05, 0x35, 0x9a, 0x25, 0x2b, 0x2e, 0xa9, 0x52, 0x17,
	0xb4, 0xbe, 0x57, 0xc1, 0x66, 0xd9, 0x04, 0xbb, 0x00, 0x7d, 0x4a, 0x35, 0x95, 0x24, 0x70, 0x1f,
	0xb3, 0xb9, 0xbe, 0x4a, 0xa4, 0x45, 0x67, 0xcd, 0x4a, 0x7b, 0x6b, 0xb8, 0xe3, 0xf2, 0x41, 0x3c,
	0x58, 0xa4, 0xbd, 0x34, 0x84, 0x77, 0xa0, 0xf1, 0x0f, 0x98, 0xba, 0xb9, 0xa4, 0x5e, 0xc8, 0x19,
	0xea, 0x38, 0x7c, 0x7f, 0x19, 0x7f, 0x66, 0x71, 0x7f, 0xde, 0x80, 0x03, 0x70, 0x94, 0x02, 0xbe,
	0x92, 0x22, 0x18, 0x27, 0x9a, 0x33, 0x42, 0xad, 0xa5, 0xfe, 0x07, 0x67, 0x85, 0x12, 0x9d, 0x3b,
	0x4d, 0x23, 0x64, 0x71, 0x2f, 0xef, 0xdd, 0x66, 0xb5, 0x5c, 0x0a, 0x6f, 0xc0, 0xc1, 0x92, 0xca,
	0xf2, 0x92, 0xe4, 0xc2, 0x49, 0xd0, 0x2f, 0xc9, 0x88, 0x17, 0xf8, 0x09, 0x80, 0x42, 0xe9, 0x29,
	0xd5, 0x8c, 0x68, 0x2e, 0xb8, 0xe6, 0xd2, 0xe7, 0x06, 0x5d, 0x3a, 0xaa, 0x9e, 0x25, 0xc3, 0x3c,
	0x80, 0xa7, 0x60, 0x9b, 0x26, 0x56, 0x65, 0xe3, 0x08, 0x0b, 0xcc, 0xfc, 0xe4, 0xae, 0x03, 0x60,
	0x9a, 0xcd, 0xe7, 0xdc, 0x67, 0x09, 0x3c, 0x04, 0x1b, 0x25, 0x02, 0x5d, 0xb9, 0x22, 0x28, 0x8a,
	0xf0, 0x0d, 0xb4, 0xcb, 0xca, 0x3f, 0xdb, 0x14, 0xc7, 0x18, 0x74, 0xed, 0xe8, 0xe3, 0x82, 0x7e,
	0x58, 0xda, 0x30, 0x3f, 0xcc, 0x78, 0x55, 0xf7, 0x5f, 0x76, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x48, 0x66, 0x2a, 0xa4, 0xad, 0x02, 0x00, 0x00,
}
