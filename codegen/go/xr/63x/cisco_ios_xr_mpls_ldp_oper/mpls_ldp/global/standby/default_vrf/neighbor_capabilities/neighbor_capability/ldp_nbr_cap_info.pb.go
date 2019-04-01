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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_neighbor_capabilities_neighbor_capability

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
	LsrId                string   `protobuf:"bytes,1,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,2,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
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
	proto.RegisterType((*LdpNbrCapInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.neighbor_capabilities.neighbor_capability.ldp_nbr_cap_info_KEYS")
	proto.RegisterType((*LdpCapDesc)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.neighbor_capabilities.neighbor_capability.ldp_cap_desc")
	proto.RegisterType((*LdpNbrCapInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.neighbor_capabilities.neighbor_capability.ldp_nbr_cap_info")
}

func init() { proto.RegisterFile("ldp_nbr_cap_info.proto", fileDescriptor_5c58ed60687e9592) }

var fileDescriptor_5c58ed60687e9592 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x4d, 0x4e, 0x02, 0x31,
	0x18, 0x86, 0x33, 0x80, 0x44, 0xca, 0x8f, 0xa6, 0x11, 0x32, 0xcb, 0x09, 0x31, 0x91, 0x55, 0x17,
	0xe0, 0x11, 0x74, 0x41, 0x74, 0x35, 0xb8, 0x71, 0xf5, 0xa5, 0x9d, 0x7e, 0x40, 0x93, 0xda, 0x36,
	0x6d, 0x25, 0x72, 0x05, 0x97, 0x5e, 0xc1, 0x2b, 0x78, 0x40, 0x33, 0x4d, 0x14, 0x9c, 0xb8, 0xd6,
	0xdd, 0xcc, 0xfb, 0x7e, 0x7d, 0xf2, 0xf4, 0x87, 0x4c, 0xb4, 0x74, 0x60, 0x84, 0x87, 0x8a, 0x3b,
	0x50, 0x66, 0x6d, 0x99, 0xf3, 0x36, 0x5a, 0xba, 0xad, 0x54, 0xa8, 0x2c, 0x28, 0x1b, 0xe0, 0xc5,
	0xc3, 0x93, 0xd3, 0x01, 0xea, 0x49, 0xeb, 0xd0, 0xb3, 0xaf, 0x3f, 0xb6, 0xd1, 0x56, 0x70, 0xcd,
	0x42, 0xe4, 0x46, 0x8a, 0x3d, 0x93, 0xb8, 0xe6, 0xcf, 0x3a, 0xc2, 0xce, 0xaf, 0x99, 0x41, 0xb5,
	0xd9, 0x0a, 0x9b, 0xb8, 0x5c, 0x28, 0xad, 0xa2, 0xc2, 0xf0, 0x4b, 0xba, 0x9f, 0x3e, 0x90, 0x71,
	0xd3, 0x01, 0xee, 0x6e, 0x1f, 0x57, 0x74, 0x4c, 0xba, 0x3a, 0x78, 0x50, 0x32, 0xcf, 0x8a, 0x6c,
	0xd6, 0x2b, 0x4f, 0x74, 0xf0, 0x4b, 0x49, 0x2f, 0xc9, 0x48, 0x73, 0x81, 0x1a, 0x82, 0xe3, 0x15,
	0xd6, 0x75, 0xab, 0xc8, 0x66, 0xc3, 0x72, 0x90, 0xd2, 0x55, 0x1d, 0x2e, 0xe5, 0xf4, 0x3d, 0x23,
	0x83, 0x1a, 0x5b, 0x23, 0x25, 0x86, 0x8a, 0x52, 0xd2, 0x89, 0x7b, 0x87, 0x89, 0x35, 0x2c, 0xd3,
	0x37, 0x2d, 0x48, 0xbf, 0xee, 0xbc, 0x72, 0x51, 0x59, 0x93, 0x38, 0xbd, 0xf2, 0x38, 0xa2, 0xd7,
	0x64, 0x72, 0x50, 0x05, 0xc9, 0x23, 0x07, 0x8d, 0x66, 0x13, 0xb7, 0x79, 0x3b, 0x71, 0x2e, 0x0e,
	0xed, 0x0d, 0x8f, 0xfc, 0x3e, 0x75, 0xf4, 0x8a, 0x9c, 0x35, 0x56, 0xe5, 0x9d, 0xc4, 0x1e, 0xfd,
	0x1c, 0x9f, 0x7e, 0xb4, 0xc8, 0x79, 0x73, 0xf3, 0xf4, 0x35, 0x23, 0x9d, 0x80, 0x26, 0xe6, 0xf3,
	0xa2, 0x3d, 0xeb, 0xcf, 0x77, 0xec, 0xaf, 0xae, 0x82, 0x1d, 0x1f, 0x58, 0x99, 0x1c, 0xe8, 0x5b,
	0x46, 0x4e, 0x3d, 0x56, 0xa8, 0x76, 0x28, 0xf3, 0xc5, 0xbf, 0x0a, 0x7d, 0x7b, 0x88, 0x6e, 0x7a,
	0xa3, 0x8b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x45, 0x46, 0xc4, 0x48, 0xbd, 0x02, 0x00, 0x00,
}