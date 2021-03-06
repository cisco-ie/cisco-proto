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
// source: ldp_ldpid_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_ldp_id

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

type LdpLdpidInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpLdpidInfo_KEYS) Reset()         { *m = LdpLdpidInfo_KEYS{} }
func (m *LdpLdpidInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpLdpidInfo_KEYS) ProtoMessage()    {}
func (*LdpLdpidInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_78168b2f4678499a, []int{0}
}

func (m *LdpLdpidInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLdpidInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpLdpidInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLdpidInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpLdpidInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLdpidInfo_KEYS.Merge(m, src)
}
func (m *LdpLdpidInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpLdpidInfo_KEYS.Size(m)
}
func (m *LdpLdpidInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLdpidInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLdpidInfo_KEYS proto.InternalMessageInfo

func (m *LdpLdpidInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpLdpidInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type LdpLdpidInfo struct {
	LsrId                string   `protobuf:"bytes,50,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,51,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	LdpId                string   `protobuf:"bytes,52,opt,name=ldp_id,json=ldpId,proto3" json:"ldp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpLdpidInfo) Reset()         { *m = LdpLdpidInfo{} }
func (m *LdpLdpidInfo) String() string { return proto.CompactTextString(m) }
func (*LdpLdpidInfo) ProtoMessage()    {}
func (*LdpLdpidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_78168b2f4678499a, []int{1}
}

func (m *LdpLdpidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLdpidInfo.Unmarshal(m, b)
}
func (m *LdpLdpidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLdpidInfo.Marshal(b, m, deterministic)
}
func (m *LdpLdpidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLdpidInfo.Merge(m, src)
}
func (m *LdpLdpidInfo) XXX_Size() int {
	return xxx_messageInfo_LdpLdpidInfo.Size(m)
}
func (m *LdpLdpidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLdpidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLdpidInfo proto.InternalMessageInfo

func (m *LdpLdpidInfo) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpLdpidInfo) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

func (m *LdpLdpidInfo) GetLdpId() string {
	if m != nil {
		return m.LdpId
	}
	return ""
}

func init() {
	proto.RegisterType((*LdpLdpidInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.ldp_id.ldp_ldpid_info_KEYS")
	proto.RegisterType((*LdpLdpidInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.ldp_id.ldp_ldpid_info")
}

func init() { proto.RegisterFile("ldp_ldpid_info.proto", fileDescriptor_78168b2f4678499a) }

var fileDescriptor_78168b2f4678499a = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x59, 0xc1, 0xf3, 0x2e, 0xe8, 0x15, 0x51, 0x61, 0xc5, 0xe6, 0x38, 0x2c, 0xae, 0x4a,
	0xe1, 0x5a, 0xdb, 0x59, 0x2c, 0xa2, 0xc5, 0x6e, 0x65, 0x35, 0x64, 0x77, 0x12, 0x08, 0x64, 0x37,
	0x61, 0x46, 0x16, 0x7f, 0xbe, 0x4c, 0xf4, 0x8a, 0x6d, 0x06, 0xf2, 0xbe, 0x6f, 0x1e, 0x19, 0x75,
	0x17, 0x31, 0x43, 0xc4, 0x1c, 0x10, 0xc2, 0xec, 0x93, 0xc9, 0x94, 0xbe, 0x93, 0x7e, 0x1d, 0x03,
	0x8f, 0x09, 0x42, 0x62, 0xf8, 0x21, 0x98, 0x72, 0x64, 0x71, 0x20, 0x65, 0x47, 0xe6, 0xfc, 0x32,
	0x73, 0x42, 0xc7, 0x65, 0x9a, 0x85, 0x3c, 0xcb, 0x30, 0x62, 0x05, 0x3c, 0x7e, 0xa8, 0xdb, 0x75,
	0x2f, 0xbc, 0xbf, 0x7d, 0xf5, 0xfa, 0x51, 0xed, 0x44, 0x87, 0xd9, 0x4e, 0xae, 0xae, 0x0e, 0xd5,
	0x69, 0xd7, 0x6d, 0x25, 0xf8, 0xb4, 0x93, 0xd3, 0x0f, 0x6a, 0xbb, 0x90, 0xff, 0x63, 0x17, 0x85,
	0x5d, 0x2d, 0xe4, 0x05, 0x1d, 0x51, 0xed, 0xd7, 0x75, 0xfa, 0x5e, 0x6d, 0x22, 0x13, 0x04, 0xac,
	0x9f, 0x8b, 0x7a, 0x19, 0x99, 0x5a, 0xd4, 0x4f, 0x6a, 0x1f, 0xed, 0xe0, 0x22, 0x70, 0xb6, 0xa3,
	0x13, 0xdc, 0x1c, 0xaa, 0xd3, 0x4d, 0x77, 0x5d, 0xd2, 0x5e, 0xc2, 0x16, 0xcb, 0x72, 0xf9, 0x67,
	0xfd, 0xf2, 0xbf, 0x8c, 0xb9, 0xc5, 0x61, 0x53, 0x6e, 0x6f, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x85, 0x67, 0xc2, 0xb2, 0x13, 0x01, 0x00, 0x00,
}
