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
// source: iep_explicit_path.proto

package cisco_ios_xr_ip_iep_oper_explicit_paths_identifiers_identifier

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

type IepExplicitPath_KEYS struct {
	IdentifierId         uint32   `protobuf:"varint,1,opt,name=identifier_id,json=identifierId,proto3" json:"identifier_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IepExplicitPath_KEYS) Reset()         { *m = IepExplicitPath_KEYS{} }
func (m *IepExplicitPath_KEYS) String() string { return proto.CompactTextString(m) }
func (*IepExplicitPath_KEYS) ProtoMessage()    {}
func (*IepExplicitPath_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d794e0861d98c35, []int{0}
}

func (m *IepExplicitPath_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IepExplicitPath_KEYS.Unmarshal(m, b)
}
func (m *IepExplicitPath_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IepExplicitPath_KEYS.Marshal(b, m, deterministic)
}
func (m *IepExplicitPath_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IepExplicitPath_KEYS.Merge(m, src)
}
func (m *IepExplicitPath_KEYS) XXX_Size() int {
	return xxx_messageInfo_IepExplicitPath_KEYS.Size(m)
}
func (m *IepExplicitPath_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IepExplicitPath_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IepExplicitPath_KEYS proto.InternalMessageInfo

func (m *IepExplicitPath_KEYS) GetIdentifierId() uint32 {
	if m != nil {
		return m.IdentifierId
	}
	return 0
}

type IepIpAddressEntry struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	IfIndex              uint32   `protobuf:"varint,2,opt,name=if_index,json=ifIndex,proto3" json:"if_index,omitempty"`
	AddressType          string   `protobuf:"bytes,3,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	HopType              string   `protobuf:"bytes,4,opt,name=hop_type,json=hopType,proto3" json:"hop_type,omitempty"`
	Address              string   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	MplsLabel            uint32   `protobuf:"varint,6,opt,name=mpls_label,json=mplsLabel,proto3" json:"mpls_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IepIpAddressEntry) Reset()         { *m = IepIpAddressEntry{} }
func (m *IepIpAddressEntry) String() string { return proto.CompactTextString(m) }
func (*IepIpAddressEntry) ProtoMessage()    {}
func (*IepIpAddressEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d794e0861d98c35, []int{1}
}

func (m *IepIpAddressEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IepIpAddressEntry.Unmarshal(m, b)
}
func (m *IepIpAddressEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IepIpAddressEntry.Marshal(b, m, deterministic)
}
func (m *IepIpAddressEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IepIpAddressEntry.Merge(m, src)
}
func (m *IepIpAddressEntry) XXX_Size() int {
	return xxx_messageInfo_IepIpAddressEntry.Size(m)
}
func (m *IepIpAddressEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_IepIpAddressEntry.DiscardUnknown(m)
}

var xxx_messageInfo_IepIpAddressEntry proto.InternalMessageInfo

func (m *IepIpAddressEntry) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *IepIpAddressEntry) GetIfIndex() uint32 {
	if m != nil {
		return m.IfIndex
	}
	return 0
}

func (m *IepIpAddressEntry) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *IepIpAddressEntry) GetHopType() string {
	if m != nil {
		return m.HopType
	}
	return ""
}

func (m *IepIpAddressEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IepIpAddressEntry) GetMplsLabel() uint32 {
	if m != nil {
		return m.MplsLabel
	}
	return 0
}

type IepExplicitPath struct {
	Status               string               `protobuf:"bytes,50,opt,name=status,proto3" json:"status,omitempty"`
	Address              []*IepIpAddressEntry `protobuf:"bytes,51,rep,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IepExplicitPath) Reset()         { *m = IepExplicitPath{} }
func (m *IepExplicitPath) String() string { return proto.CompactTextString(m) }
func (*IepExplicitPath) ProtoMessage()    {}
func (*IepExplicitPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d794e0861d98c35, []int{2}
}

func (m *IepExplicitPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IepExplicitPath.Unmarshal(m, b)
}
func (m *IepExplicitPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IepExplicitPath.Marshal(b, m, deterministic)
}
func (m *IepExplicitPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IepExplicitPath.Merge(m, src)
}
func (m *IepExplicitPath) XXX_Size() int {
	return xxx_messageInfo_IepExplicitPath.Size(m)
}
func (m *IepExplicitPath) XXX_DiscardUnknown() {
	xxx_messageInfo_IepExplicitPath.DiscardUnknown(m)
}

var xxx_messageInfo_IepExplicitPath proto.InternalMessageInfo

func (m *IepExplicitPath) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *IepExplicitPath) GetAddress() []*IepIpAddressEntry {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*IepExplicitPath_KEYS)(nil), "cisco_ios_xr_ip_iep_oper.explicit_paths.identifiers.identifier.iep_explicit_path_KEYS")
	proto.RegisterType((*IepIpAddressEntry)(nil), "cisco_ios_xr_ip_iep_oper.explicit_paths.identifiers.identifier.iep_ip_address_entry")
	proto.RegisterType((*IepExplicitPath)(nil), "cisco_ios_xr_ip_iep_oper.explicit_paths.identifiers.identifier.iep_explicit_path")
}

func init() { proto.RegisterFile("iep_explicit_path.proto", fileDescriptor_2d794e0861d98c35) }

var fileDescriptor_2d794e0861d98c35 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x95, 0x7f, 0xff, 0x6d, 0xe9, 0xb5, 0x1d, 0xb0, 0xaa, 0x62, 0x06, 0xa4, 0x52, 0x96,
	0x4e, 0x19, 0xda, 0x19, 0x36, 0x86, 0x0a, 0xa6, 0xd2, 0x85, 0xe9, 0x94, 0xd6, 0x8e, 0x72, 0x52,
	0x88, 0x4f, 0xb6, 0x91, 0x92, 0xcf, 0xc2, 0x77, 0xe1, 0xb3, 0xa1, 0x38, 0x81, 0x80, 0xca, 0xc6,
	0x96, 0x7b, 0xbf, 0x77, 0xef, 0xa2, 0x67, 0xb8, 0x20, 0xcd, 0xa8, 0x4b, 0xce, 0xe9, 0x48, 0x1e,
	0x39, 0xf1, 0x59, 0xcc, 0xd6, 0x78, 0x23, 0xee, 0x8e, 0xe4, 0x8e, 0x06, 0xc9, 0x38, 0x2c, 0x2d,
	0x12, 0x63, 0x6d, 0x34, 0xac, 0x6d, 0xfc, 0xc3, 0xed, 0x62, 0x52, 0xba, 0xf0, 0x94, 0x92, 0xb6,
	0xdf, 0xbf, 0x97, 0xb7, 0x30, 0x3f, 0x89, 0xc6, 0x87, 0xfb, 0xe7, 0x27, 0x71, 0x03, 0xd3, 0xce,
	0x87, 0xa4, 0x64, 0xb4, 0x88, 0x56, 0xd3, 0xdd, 0xa4, 0x13, 0xb7, 0x6a, 0xf9, 0x1e, 0xc1, 0xac,
	0xde, 0x27, 0xc6, 0x44, 0x29, 0xab, 0x9d, 0x43, 0x5d, 0x78, 0x5b, 0x89, 0x19, 0xf4, 0xa9, 0x50,
	0xba, 0x6c, 0xb7, 0x9a, 0x41, 0x5c, 0xc2, 0x19, 0xa5, 0xd8, 0x80, 0x7f, 0x01, 0x0c, 0x29, 0xdd,
	0x06, 0x74, 0x0d, 0x93, 0xcf, 0x04, 0x5f, 0xb1, 0x96, 0xbd, 0x45, 0xb4, 0x1a, 0xed, 0xc6, 0xad,
	0xb6, 0xaf, 0x58, 0xd7, 0xdb, 0x99, 0xe1, 0x06, 0xff, 0x0f, 0x78, 0x98, 0x19, 0x0e, 0x48, 0xc2,
	0xb0, 0x75, 0xca, 0x7e, 0x43, 0xda, 0x51, 0x5c, 0x01, 0xbc, 0x70, 0xee, 0x30, 0x4f, 0x0e, 0x3a,
	0x97, 0x83, 0x70, 0x74, 0x54, 0x2b, 0x8f, 0xb5, 0xb0, 0x7c, 0x8b, 0xe0, 0xfc, 0xa4, 0x00, 0x31,
	0x87, 0x81, 0xf3, 0x89, 0x7f, 0x75, 0x72, 0x1d, 0xd2, 0xda, 0x49, 0x14, 0xdd, 0x99, 0xcd, 0xa2,
	0xb7, 0x1a, 0xaf, 0xf7, 0xf1, 0xdf, 0xfa, 0x8f, 0x7f, 0x2b, 0xef, 0xeb, 0xe7, 0x0f, 0x83, 0xf0,
	0xc8, 0x9b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0xe6, 0x8c, 0xfd, 0xff, 0x01, 0x00, 0x00,
}
