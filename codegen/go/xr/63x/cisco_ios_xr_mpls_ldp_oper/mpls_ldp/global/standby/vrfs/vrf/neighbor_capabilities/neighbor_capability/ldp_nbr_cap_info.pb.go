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
// source: ldp_nbr_cap_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_neighbor_capabilities_neighbor_capability

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

type LdpNbrCapInfo_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LsrId                string   `protobuf:"bytes,2,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,3,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNbrCapInfo_KEYS) Reset()         { *m = LdpNbrCapInfo_KEYS{} }
func (m *LdpNbrCapInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpNbrCapInfo_KEYS) ProtoMessage()    {}
func (*LdpNbrCapInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c58ed60687e9592, []int{0}
}

func (m *LdpNbrCapInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNbrCapInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpNbrCapInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNbrCapInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpNbrCapInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNbrCapInfo_KEYS.Merge(m, src)
}
func (m *LdpNbrCapInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpNbrCapInfo_KEYS.Size(m)
}
func (m *LdpNbrCapInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNbrCapInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNbrCapInfo_KEYS proto.InternalMessageInfo

func (m *LdpNbrCapInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpNbrCapInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpNbrCapInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
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
	return fileDescriptor_5c58ed60687e9592, []int{1}
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

type LdpNbrCapInfo struct {
	Sent                 []*LdpCapDesc `protobuf:"bytes,50,rep,name=sent,proto3" json:"sent,omitempty"`
	Received             []*LdpCapDesc `protobuf:"bytes,51,rep,name=received,proto3" json:"received,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LdpNbrCapInfo) Reset()         { *m = LdpNbrCapInfo{} }
func (m *LdpNbrCapInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNbrCapInfo) ProtoMessage()    {}
func (*LdpNbrCapInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c58ed60687e9592, []int{2}
}

func (m *LdpNbrCapInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNbrCapInfo.Unmarshal(m, b)
}
func (m *LdpNbrCapInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNbrCapInfo.Marshal(b, m, deterministic)
}
func (m *LdpNbrCapInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNbrCapInfo.Merge(m, src)
}
func (m *LdpNbrCapInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNbrCapInfo.Size(m)
}
func (m *LdpNbrCapInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNbrCapInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNbrCapInfo proto.InternalMessageInfo

func (m *LdpNbrCapInfo) GetSent() []*LdpCapDesc {
	if m != nil {
		return m.Sent
	}
	return nil
}

func (m *LdpNbrCapInfo) GetReceived() []*LdpCapDesc {
	if m != nil {
		return m.Received
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpNbrCapInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.neighbor_capabilities.neighbor_capability.ldp_nbr_cap_info_KEYS")
	proto.RegisterType((*LdpCapDesc)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.neighbor_capabilities.neighbor_capability.ldp_cap_desc")
	proto.RegisterType((*LdpNbrCapInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.neighbor_capabilities.neighbor_capability.ldp_nbr_cap_info")
}

func init() { proto.RegisterFile("ldp_nbr_cap_info.proto", fileDescriptor_5c58ed60687e9592) }

var fileDescriptor_5c58ed60687e9592 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcf, 0x4e, 0x2a, 0x31,
	0x14, 0xc6, 0x33, 0xc0, 0xe5, 0x72, 0xcb, 0x9f, 0x7b, 0xd3, 0x5c, 0xc8, 0xb8, 0x9b, 0x10, 0x13,
	0x59, 0x75, 0x01, 0x3e, 0x82, 0x2e, 0x88, 0xc6, 0xc5, 0xb0, 0x72, 0xd5, 0xb4, 0xd3, 0x33, 0xd0,
	0xa4, 0xb4, 0x4d, 0xdb, 0x4c, 0xe4, 0x09, 0x5c, 0xfa, 0x10, 0xee, 0x7d, 0x46, 0xd3, 0x8a, 0x82,
	0xc4, 0xb5, 0x6e, 0x26, 0xd3, 0xef, 0x3b, 0x3d, 0xbf, 0xaf, 0xa7, 0x45, 0x13, 0x25, 0x2c, 0xd5,
	0xdc, 0xd1, 0x8a, 0x59, 0x2a, 0x75, 0x6d, 0x88, 0x75, 0x26, 0x18, 0x0c, 0x95, 0xf4, 0x95, 0xa1,
	0xd2, 0x78, 0xfa, 0xe0, 0xe8, 0xd6, 0x2a, 0x4f, 0x63, 0xa5, 0xb1, 0xe0, 0xc8, 0xfb, 0x8a, 0xac,
	0x95, 0xe1, 0x4c, 0x11, 0x1f, 0x98, 0x16, 0x7c, 0x47, 0x1a, 0x57, 0xfb, 0xf8, 0x21, 0x1a, 0xe4,
	0x7a, 0xc3, 0x4d, 0x6a, 0xca, 0xb8, 0x54, 0x32, 0x48, 0xf0, 0x5f, 0xa8, 0xbb, 0xa9, 0x41, 0xe3,
	0xd3, 0x00, 0xf4, 0xe6, 0xfa, 0x7e, 0x85, 0xcf, 0x50, 0xaf, 0x71, 0x35, 0xd5, 0x6c, 0x0b, 0x79,
	0x56, 0x64, 0xb3, 0x3f, 0xe5, 0xef, 0xc6, 0xd5, 0x77, 0x6c, 0x0b, 0x78, 0x8c, 0xba, 0xca, 0x3b,
	0x2a, 0x45, 0xde, 0x4a, 0xc6, 0x2f, 0xe5, 0xdd, 0x52, 0xe0, 0x73, 0x34, 0x52, 0x8c, 0x83, 0xa2,
	0xde, 0xb2, 0x0a, 0xa2, 0xdd, 0x2e, 0xb2, 0xd9, 0xb0, 0x1c, 0x24, 0x75, 0x15, 0xc5, 0xa5, 0x98,
	0x3e, 0x67, 0x68, 0x10, 0x89, 0x91, 0x26, 0xc0, 0x57, 0x18, 0xa3, 0x4e, 0xd8, 0xd9, 0x37, 0xc8,
	0xb0, 0x4c, 0xff, 0xb8, 0x40, 0xfd, 0xe8, 0x39, 0x69, 0x83, 0x34, 0x7a, 0x8f, 0x39, 0x96, 0xf0,
	0x25, 0x9a, 0x1c, 0x4e, 0x41, 0x05, 0x0b, 0x8c, 0x2a, 0xd0, 0xeb, 0xb0, 0xd9, 0x43, 0xff, 0x1f,
	0xdc, 0x2b, 0x16, 0xd8, 0x6d, 0xf2, 0xf0, 0x05, 0xfa, 0x7b, 0xb2, 0x2b, 0xef, 0xa4, 0xde, 0xa3,
	0xcf, 0xe5, 0xd3, 0x97, 0x16, 0xfa, 0x77, 0x3a, 0x17, 0xfc, 0x98, 0xa1, 0x8e, 0x07, 0x1d, 0xf2,
	0x79, 0xd1, 0x9e, 0xf5, 0xe7, 0x9e, 0x7c, 0xcb, 0x15, 0x91, 0xe3, 0x69, 0x95, 0x29, 0x00, 0x7e,
	0xca, 0x50, 0xcf, 0x41, 0x05, 0xb2, 0x01, 0x91, 0x2f, 0x7e, 0x2e, 0xcd, 0x47, 0x08, 0xde, 0x4d,
	0xaf, 0x76, 0xf1, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x77, 0x46, 0x23, 0x82, 0xcf, 0x02, 0x00, 0x00,
}
