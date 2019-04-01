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
// source: fs_mgr_tbl_summary.proto

package cisco_ios_xr_flowspec_oper_flow_spec_vrfs_vrf_afs_af_table_summary

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

type FsMgrTblSummary_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrTblSummary_KEYS) Reset()         { *m = FsMgrTblSummary_KEYS{} }
func (m *FsMgrTblSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsMgrTblSummary_KEYS) ProtoMessage()    {}
func (*FsMgrTblSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d305bb1f39e30744, []int{0}
}

func (m *FsMgrTblSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrTblSummary_KEYS.Unmarshal(m, b)
}
func (m *FsMgrTblSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrTblSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *FsMgrTblSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrTblSummary_KEYS.Merge(m, src)
}
func (m *FsMgrTblSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsMgrTblSummary_KEYS.Size(m)
}
func (m *FsMgrTblSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrTblSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrTblSummary_KEYS proto.InternalMessageInfo

func (m *FsMgrTblSummary_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *FsMgrTblSummary_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

type FsMgrTblSummary struct {
	TotalFlows           uint32   `protobuf:"varint,50,opt,name=total_flows,json=totalFlows,proto3" json:"total_flows,omitempty"`
	ServicePolicies      uint32   `protobuf:"varint,51,opt,name=service_policies,json=servicePolicies,proto3" json:"service_policies,omitempty"`
	LocalInstallEnabled  bool     `protobuf:"varint,52,opt,name=local_install_enabled,json=localInstallEnabled,proto3" json:"local_install_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrTblSummary) Reset()         { *m = FsMgrTblSummary{} }
func (m *FsMgrTblSummary) String() string { return proto.CompactTextString(m) }
func (*FsMgrTblSummary) ProtoMessage()    {}
func (*FsMgrTblSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_d305bb1f39e30744, []int{1}
}

func (m *FsMgrTblSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrTblSummary.Unmarshal(m, b)
}
func (m *FsMgrTblSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrTblSummary.Marshal(b, m, deterministic)
}
func (m *FsMgrTblSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrTblSummary.Merge(m, src)
}
func (m *FsMgrTblSummary) XXX_Size() int {
	return xxx_messageInfo_FsMgrTblSummary.Size(m)
}
func (m *FsMgrTblSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrTblSummary.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrTblSummary proto.InternalMessageInfo

func (m *FsMgrTblSummary) GetTotalFlows() uint32 {
	if m != nil {
		return m.TotalFlows
	}
	return 0
}

func (m *FsMgrTblSummary) GetServicePolicies() uint32 {
	if m != nil {
		return m.ServicePolicies
	}
	return 0
}

func (m *FsMgrTblSummary) GetLocalInstallEnabled() bool {
	if m != nil {
		return m.LocalInstallEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*FsMgrTblSummary_KEYS)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.table_summary.fs_mgr_tbl_summary_KEYS")
	proto.RegisterType((*FsMgrTblSummary)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.vrfs.vrf.afs.af.table_summary.fs_mgr_tbl_summary")
}

func init() { proto.RegisterFile("fs_mgr_tbl_summary.proto", fileDescriptor_d305bb1f39e30744) }

var fileDescriptor_d305bb1f39e30744 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x87, 0xdd, 0x35, 0x22, 0x4a, 0x44, 0x36, 0x9e, 0x2c, 0x7b, 0xaa, 0x97, 0x1c,
	0x76, 0x7d, 0x02, 0x61, 0x05, 0x11, 0x45, 0xea, 0xc9, 0xd3, 0x30, 0x8d, 0x89, 0x04, 0x92, 0xa6,
	0x64, 0x62, 0xd5, 0xf7, 0xf0, 0x81, 0x25, 0x69, 0x3d, 0xed, 0x25, 0xf0, 0x7f, 0x5f, 0xf8, 0x99,
	0x19, 0x26, 0x0c, 0x81, 0xff, 0x88, 0x90, 0x3a, 0x07, 0xf4, 0xe9, 0x3d, 0xc6, 0x1f, 0x39, 0xc4,
	0x90, 0x02, 0xbf, 0x53, 0x96, 0x54, 0x00, 0x1b, 0x08, 0xbe, 0x23, 0x18, 0x17, 0xbe, 0x68, 0xd0,
	0x0a, 0xc2, 0xa0, 0xa3, 0xcc, 0x09, 0x72, 0x94, 0x63, 0x34, 0x94, 0x1f, 0x89, 0x86, 0x24, 0x1a,
	0x99, 0xb0, 0x73, 0xfa, 0xbf, 0x69, 0xf3, 0xc4, 0xd6, 0x87, 0xfd, 0xf0, 0xb8, 0x7f, 0x7b, 0xe5,
	0x57, 0x6c, 0x35, 0x46, 0x03, 0x3d, 0x7a, 0x2d, 0xaa, 0xba, 0x6a, 0x8e, 0xdb, 0xe5, 0x18, 0xcd,
	0x33, 0x7a, 0xcd, 0xd7, 0x6c, 0x89, 0xb3, 0x39, 0x2a, 0x66, 0x81, 0x45, 0x6c, 0x7e, 0x2b, 0xc6,
	0x0f, 0xfb, 0xf8, 0x35, 0x3b, 0x49, 0x21, 0xa1, 0x9b, 0x86, 0x14, 0xdb, 0xba, 0x6a, 0x4e, 0x5b,
	0x56, 0xd0, 0x7d, 0x26, 0xfc, 0x86, 0x9d, 0x93, 0x8e, 0xa3, 0x55, 0x1a, 0x86, 0xe0, 0xac, 0xb2,
	0x9a, 0xc4, 0xae, 0xfc, 0x3a, 0x9b, 0xf9, 0xcb, 0x8c, 0xf9, 0x96, 0x5d, 0xba, 0xa0, 0xd0, 0x81,
	0xed, 0x29, 0xa1, 0x73, 0xa0, 0xfb, 0xbc, 0xd1, 0xbb, 0xb8, 0xad, 0xab, 0x66, 0xd5, 0x5e, 0x14,
	0xf9, 0x30, 0xb9, 0xfd, 0xa4, 0xba, 0x45, 0x39, 0xd8, 0xee, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x36,
	0xcc, 0x67, 0x82, 0x4c, 0x01, 0x00, 0x00,
}