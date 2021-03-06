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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_interface_summary

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
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
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
	proto.RegisterType((*LdpIntfSum_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.interface_summary.ldp_intf_sum_KEYS")
	proto.RegisterType((*LdpIntfSum)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.interface_summary.ldp_intf_sum")
}

func init() { proto.RegisterFile("ldp_intf_sum.proto", fileDescriptor_b93a20219fd1fd8e) }

var fileDescriptor_b93a20219fd1fd8e = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x09, 0x83, 0x8c, 0x69, 0xdb, 0x21, 0x62, 0x63, 0x62, 0x83, 0x2c, 0x0b, 0x3b, 0xe4,
	0xb0, 0x89, 0xb1, 0x6c, 0xcb, 0x7a, 0xe8, 0xa1, 0x4d, 0x53, 0x08, 0x2d, 0x3d, 0xa4, 0xa5, 0xa5,
	0x27, 0xf1, 0x6c, 0x3d, 0xa5, 0xa6, 0xb6, 0x64, 0x24, 0xb9, 0x69, 0x3e, 0x75, 0xbf, 0x42, 0xb1,
	0xe2, 0xd8, 0x6e, 0xda, 0xa3, 0xfd, 0xff, 0xff, 0x7e, 0xef, 0x3d, 0x10, 0xa1, 0xa9, 0xcc, 0x45,
	0xa2, 0xbd, 0x12, 0xae, 0xc8, 0x78, 0x6e, 0x8d, 0x37, 0xf4, 0x2a, 0x4e, 0x5c, 0x6c, 0x44, 0x62,
	0x9c, 0xb8, 0xb7, 0x22, 0xcb, 0x53, 0x27, 0xca, 0x96, 0xc9, 0xd1, 0xf2, 0xed, 0x17, 0x5f, 0xa6,
	0x26, 0x82, 0x94, 0x3b, 0x0f, 0x5a, 0x46, 0x6b, 0x2e, 0x51, 0x41, 0x91, 0x7a, 0x71, 0x67, 0x15,
	0x07, 0xe5, 0x38, 0x28, 0x9e, 0x68, 0x8f, 0x56, 0x41, 0x8c, 0xa5, 0x3a, 0x03, 0xbb, 0x1e, 0xfe,
	0x20, 0xbd, 0xf6, 0x38, 0x71, 0x32, 0xbb, 0x3e, 0xa7, 0x9f, 0xc8, 0x6b, 0x50, 0x42, 0x43, 0x86,
	0xac, 0x33, 0xe8, 0x8c, 0xde, 0x2c, 0xba, 0xa0, 0xce, 0x20, 0xc3, 0xe1, 0xc3, 0x2b, 0xf2, 0xae,
	0x5d, 0xa7, 0x13, 0xc2, 0x6e, 0xb5, 0x59, 0x69, 0x91, 0x84, 0x9f, 0x95, 0x3c, 0x36, 0x85, 0xf6,
	0xec, 0xf7, 0xa0, 0x33, 0x7a, 0xbf, 0xf8, 0x18, 0xf2, 0x79, 0x3e, 0xdf, 0xa6, 0xd3, 0x32, 0xa4,
	0x87, 0xa4, 0xff, 0x02, 0x58, 0xba, 0x51, 0x43, 0x94, 0xa2, 0x64, 0xe3, 0x80, 0x7f, 0xde, 0xc5,
	0x4f, 0x65, 0x3e, 0xdb, 0x34, 0xe8, 0x9c, 0x7c, 0x2b, 0x81, 0xd8, 0x68, 0x95, 0x2c, 0x0b, 0x8b,
	0x52, 0x80, 0xf7, 0x10, 0xdf, 0xa0, 0x6c, 0x94, 0xec, 0x4f, 0xd0, 0xf4, 0x53, 0x99, 0x4f, 0xeb,
	0xde, 0x41, 0x55, 0xab, 0xa5, 0x74, 0x9f, 0x7c, 0xd9, 0x51, 0x79, 0x6c, 0x49, 0xfe, 0x06, 0x09,
	0x7b, 0x22, 0xb9, 0xc0, 0x06, 0xff, 0x49, 0xa8, 0x32, 0x76, 0x05, 0x56, 0x0a, 0x8b, 0x0a, 0x2d,
	0xea, 0x18, 0x1d, 0xfb, 0x17, 0xa8, 0x5e, 0x95, 0x2c, 0xea, 0x80, 0xfe, 0x22, 0x1f, 0xa0, 0xf0,
	0xa6, 0x1a, 0x27, 0x64, 0xe2, 0x36, 0x27, 0x4f, 0x02, 0x40, 0xcb, 0x6c, 0x33, 0xe7, 0xa8, 0x4a,
	0xe8, 0x57, 0xf2, 0xb6, 0x45, 0xb0, 0xff, 0xa1, 0x48, 0x9a, 0x22, 0xbd, 0x24, 0xa3, 0xb6, 0xf2,
	0xd9, 0x36, 0xcd, 0x31, 0x8e, 0xed, 0x05, 0xfa, 0x7b, 0x43, 0x1f, 0xef, 0x6c, 0x58, 0x1f, 0xe6,
	0xa2, 0x6e, 0x78, 0x7f, 0xe3, 0xc7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x22, 0x84, 0x75, 0x95,
	0x02, 0x00, 0x00,
}
