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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_afs_af_interface_summary

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
	proto.RegisterType((*LdpIntfSum_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.interface_summary.ldp_intf_sum_KEYS")
	proto.RegisterType((*LdpIntfSum)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.interface_summary.ldp_intf_sum")
}

func init() { proto.RegisterFile("ldp_intf_sum.proto", fileDescriptor_b93a20219fd1fd8e) }

var fileDescriptor_b93a20219fd1fd8e = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0x29, 0x42, 0xc5, 0xa8, 0x87, 0x06, 0xc5, 0xa0, 0x50, 0x6b, 0xf1, 0xd0, 0x83, 0x06,
	0xb1, 0x6a, 0xf5, 0xe0, 0x41, 0x6b, 0x85, 0xa2, 0x78, 0xa8, 0x52, 0xf0, 0x14, 0x5e, 0x93, 0x97,
	0x6e, 0xd8, 0x99, 0x64, 0x48, 0x32, 0xed, 0xee, 0xa7, 0xde, 0xaf, 0xb0, 0x4c, 0x3a, 0x9d, 0x99,
	0xed, 0xee, 0x71, 0xe6, 0xff, 0xff, 0xfd, 0xde, 0x7b, 0x10, 0x42, 0x33, 0x55, 0x08, 0x63, 0xa3,
	0x16, 0xa1, 0xcc, 0x79, 0xe1, 0x5d, 0x74, 0x74, 0x2d, 0x4d, 0x90, 0x4e, 0x18, 0x17, 0xc4, 0x85,
	0x17, 0x79, 0x91, 0x05, 0x51, 0xb5, 0x5c, 0x81, 0x9e, 0x1f, 0xbf, 0xf8, 0x36, 0x73, 0x1b, 0xc8,
	0x38, 0xc8, 0x68, 0x76, 0xc8, 0x15, 0x6a, 0x28, 0xb3, 0x28, 0x76, 0x5e, 0x73, 0xd0, 0x81, 0x83,
	0xe6, 0xc6, 0x46, 0xf4, 0x1a, 0x24, 0x56, 0xe6, 0x1c, 0xfc, 0xe5, 0xf8, 0x0d, 0x19, 0x74, 0xa7,
	0x89, 0x5f, 0x8b, 0xff, 0x7f, 0xe9, 0x33, 0x72, 0x1f, 0xb4, 0xb0, 0x90, 0x23, 0xeb, 0x8d, 0x7a,
	0x93, 0x07, 0xab, 0x3e, 0xe8, 0x3f, 0x90, 0xe3, 0xf8, 0xea, 0x1e, 0x79, 0xd4, 0xad, 0xd3, 0x19,
	0x61, 0xe7, 0xd6, 0xed, 0xad, 0x30, 0xe9, 0x67, 0x2d, 0x97, 0xae, 0xb4, 0x91, 0xbd, 0x1f, 0xf5,
	0x26, 0x8f, 0x57, 0x4f, 0x53, 0xbe, 0x2c, 0x96, 0xc7, 0x74, 0x5e, 0x85, 0xf4, 0x3b, 0x19, 0xde,
	0x01, 0x56, 0x6e, 0xb4, 0xb0, 0xc9, 0x50, 0xb1, 0x69, 0xc2, 0x9f, 0x9f, 0xe2, 0xbf, 0x55, 0xb1,
	0x38, 0x34, 0xe8, 0x92, 0xbc, 0xaa, 0x00, 0xe9, 0xac, 0x36, 0xdb, 0xd2, 0xa3, 0x12, 0x10, 0x23,
	0xc8, 0x33, 0x54, 0xad, 0x92, 0x7d, 0x48, 0x9a, 0x61, 0xa6, 0x8a, 0x79, 0xd3, 0xfb, 0x56, 0xd7,
	0x1a, 0x29, 0xfd, 0x4a, 0x5e, 0x9c, 0xa8, 0x22, 0x76, 0x24, 0x1f, 0x93, 0x84, 0xdd, 0x90, 0xfc,
	0xc3, 0x16, 0x7f, 0x4b, 0xa8, 0x76, 0x7e, 0x0f, 0x5e, 0x09, 0x8f, 0x1a, 0x3d, 0x5a, 0x89, 0x81,
	0x7d, 0x4a, 0xd4, 0xa0, 0x4e, 0x56, 0x4d, 0x40, 0xdf, 0x91, 0x27, 0x50, 0x46, 0x57, 0x8f, 0x13,
	0xca, 0x84, 0xc3, 0xc9, 0xb3, 0x04, 0xd0, 0x2a, 0x3b, 0xcc, 0xf9, 0x51, 0x27, 0xf4, 0x25, 0x79,
	0xd8, 0x21, 0xd8, 0xe7, 0x54, 0x24, 0x6d, 0x91, 0xae, 0xc9, 0xa4, 0xab, 0xbc, 0xb5, 0x4d, 0x7b,
	0x4c, 0x60, 0x5f, 0x12, 0xfd, 0xba, 0xa5, 0x7f, 0x9e, 0x6c, 0xd8, 0x1c, 0x16, 0x36, 0xfd, 0xf4,
	0xfc, 0xa6, 0xd7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x8c, 0xc8, 0x43, 0x94, 0x02, 0x00, 0x00,
}
