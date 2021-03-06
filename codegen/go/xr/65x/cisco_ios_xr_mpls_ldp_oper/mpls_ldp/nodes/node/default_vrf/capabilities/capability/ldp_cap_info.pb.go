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
// source: ldp_cap_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_capabilities_capability

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

type LdpCapInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	CapabilityType       uint32   `protobuf:"varint,2,opt,name=capability_type,json=capabilityType,proto3" json:"capability_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpCapInfo_KEYS) Reset()         { *m = LdpCapInfo_KEYS{} }
func (m *LdpCapInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpCapInfo_KEYS) ProtoMessage()    {}
func (*LdpCapInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30991e82f8c7a91, []int{0}
}

func (m *LdpCapInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpCapInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpCapInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpCapInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpCapInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpCapInfo_KEYS.Merge(m, src)
}
func (m *LdpCapInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpCapInfo_KEYS.Size(m)
}
func (m *LdpCapInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpCapInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpCapInfo_KEYS proto.InternalMessageInfo

func (m *LdpCapInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpCapInfo_KEYS) GetCapabilityType() uint32 {
	if m != nil {
		return m.CapabilityType
	}
	return 0
}

type LdpCapDesc struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CapabilityDataLength uint32   `protobuf:"varint,3,opt,name=capability_data_length,json=capabilityDataLength,proto3" json:"capability_data_length,omitempty"`
	CapabilityData       string   `protobuf:"bytes,4,opt,name=capability_data,json=capabilityData,proto3" json:"capability_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpCapDesc) Reset()         { *m = LdpCapDesc{} }
func (m *LdpCapDesc) String() string { return proto.CompactTextString(m) }
func (*LdpCapDesc) ProtoMessage()    {}
func (*LdpCapDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30991e82f8c7a91, []int{1}
}

func (m *LdpCapDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpCapDesc.Unmarshal(m, b)
}
func (m *LdpCapDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpCapDesc.Marshal(b, m, deterministic)
}
func (m *LdpCapDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpCapDesc.Merge(m, src)
}
func (m *LdpCapDesc) XXX_Size() int {
	return xxx_messageInfo_LdpCapDesc.Size(m)
}
func (m *LdpCapDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpCapDesc.DiscardUnknown(m)
}

var xxx_messageInfo_LdpCapDesc proto.InternalMessageInfo

func (m *LdpCapDesc) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *LdpCapDesc) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *LdpCapDesc) GetCapabilityDataLength() uint32 {
	if m != nil {
		return m.CapabilityDataLength
	}
	return 0
}

func (m *LdpCapDesc) GetCapabilityData() string {
	if m != nil {
		return m.CapabilityData
	}
	return ""
}

type LdpCapInfo struct {
	Capability           *LdpCapDesc `protobuf:"bytes,50,opt,name=capability,proto3" json:"capability,omitempty"`
	CapabilityOwner      string      `protobuf:"bytes,51,opt,name=capability_owner,json=capabilityOwner,proto3" json:"capability_owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LdpCapInfo) Reset()         { *m = LdpCapInfo{} }
func (m *LdpCapInfo) String() string { return proto.CompactTextString(m) }
func (*LdpCapInfo) ProtoMessage()    {}
func (*LdpCapInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30991e82f8c7a91, []int{2}
}

func (m *LdpCapInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpCapInfo.Unmarshal(m, b)
}
func (m *LdpCapInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpCapInfo.Marshal(b, m, deterministic)
}
func (m *LdpCapInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpCapInfo.Merge(m, src)
}
func (m *LdpCapInfo) XXX_Size() int {
	return xxx_messageInfo_LdpCapInfo.Size(m)
}
func (m *LdpCapInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpCapInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpCapInfo proto.InternalMessageInfo

func (m *LdpCapInfo) GetCapability() *LdpCapDesc {
	if m != nil {
		return m.Capability
	}
	return nil
}

func (m *LdpCapInfo) GetCapabilityOwner() string {
	if m != nil {
		return m.CapabilityOwner
	}
	return ""
}

func init() {
	proto.RegisterType((*LdpCapInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.capabilities.capability.ldp_cap_info_KEYS")
	proto.RegisterType((*LdpCapDesc)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.capabilities.capability.ldp_cap_desc")
	proto.RegisterType((*LdpCapInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.capabilities.capability.ldp_cap_info")
}

func init() { proto.RegisterFile("ldp_cap_info.proto", fileDescriptor_c30991e82f8c7a91) }

var fileDescriptor_c30991e82f8c7a91 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4a, 0x34, 0x31,
	0x10, 0x24, 0xdf, 0xb7, 0x88, 0x9b, 0xd5, 0x55, 0x83, 0x48, 0xc0, 0xcb, 0xb0, 0x17, 0xc7, 0x4b,
	0x0e, 0xbb, 0x3e, 0x82, 0x9e, 0x14, 0x85, 0xd1, 0xcb, 0x9e, 0xda, 0xec, 0x4c, 0x8f, 0x06, 0x66,
	0x92, 0x90, 0xc4, 0x9f, 0xb9, 0xf9, 0x2e, 0xbe, 0x86, 0x0f, 0x27, 0x09, 0x2c, 0x13, 0x3d, 0x7b,
	0x09, 0xe9, 0xea, 0xea, 0xea, 0x4a, 0x11, 0xca, 0xba, 0xc6, 0x42, 0x2d, 0x2d, 0x28, 0xdd, 0x1a,
	0x61, 0x9d, 0x09, 0x86, 0x55, 0xb5, 0xf2, 0xb5, 0x01, 0x65, 0x3c, 0xbc, 0x3b, 0xe8, 0x6d, 0xe7,
	0x21, 0xb2, 0x8c, 0x45, 0x27, 0xb6, 0x95, 0xd0, 0xa6, 0x41, 0x9f, 0x4e, 0xd1, 0x60, 0x2b, 0x5f,
	0xba, 0x00, 0xaf, 0xae, 0x15, 0xb5, 0xb4, 0x72, 0xa3, 0x3a, 0x15, 0x14, 0xfa, 0xb1, 0x18, 0x16,
	0x6b, 0x7a, 0x94, 0x6f, 0x82, 0xeb, 0xab, 0xf5, 0x3d, 0x3b, 0xa5, 0xd3, 0x28, 0x00, 0x5a, 0xf6,
	0xc8, 0x49, 0x41, 0xca, 0x69, 0xb5, 0x1b, 0x81, 0x5b, 0xd9, 0x23, 0x3b, 0xa3, 0x07, 0xe3, 0x3c,
	0x84, 0xc1, 0x22, 0xff, 0x57, 0x90, 0x72, 0xbf, 0x9a, 0x8f, 0xf0, 0xc3, 0x60, 0x71, 0xf1, 0x49,
	0xe8, 0xde, 0x56, 0xbb, 0x41, 0x5f, 0x33, 0x46, 0x27, 0x89, 0x4e, 0x12, 0x3d, 0xdd, 0x59, 0x41,
	0x67, 0xb1, 0xe7, 0x94, 0x0d, 0xca, 0xe8, 0xa4, 0x34, 0xad, 0x72, 0x88, 0x5d, 0xd0, 0x93, 0x6c,
	0x5f, 0x23, 0x83, 0x84, 0x0e, 0xf5, 0x53, 0x78, 0xe6, 0xff, 0x93, 0xce, 0xf1, 0xd8, 0xbd, 0x94,
	0x41, 0xde, 0xa4, 0xde, 0x2f, 0x97, 0x71, 0x8a, 0x4f, 0x92, 0xf6, 0xfc, 0x27, 0x7d, 0xf1, 0x95,
	0xb9, 0x8c, 0x09, 0xb0, 0x0f, 0x42, 0xe9, 0xc8, 0xe1, 0xcb, 0x82, 0x94, 0xb3, 0xe5, 0xa3, 0xf8,
	0xfb, 0xec, 0x45, 0x1e, 0x4e, 0x95, 0xed, 0x64, 0xe7, 0xf4, 0x30, 0x33, 0x6f, 0xde, 0x34, 0x3a,
	0xbe, 0x4a, 0xee, 0xb3, 0x47, 0xdd, 0x45, 0x78, 0xb3, 0x93, 0xbe, 0xc6, 0xea, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x60, 0xf7, 0xe3, 0xe0, 0x30, 0x02, 0x00, 0x00,
}
