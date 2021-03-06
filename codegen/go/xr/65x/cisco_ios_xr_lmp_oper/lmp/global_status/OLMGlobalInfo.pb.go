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
// source: OLMGlobalInfo.proto

package cisco_ios_xr_lmp_oper_lmp_global_status

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

type OLMGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OLMGlobalInfo_KEYS) Reset()         { *m = OLMGlobalInfo_KEYS{} }
func (m *OLMGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OLMGlobalInfo_KEYS) ProtoMessage()    {}
func (*OLMGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bda31b4ffaf29a, []int{0}
}

func (m *OLMGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *OLMGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *OLMGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMGlobalInfo_KEYS.Merge(m, src)
}
func (m *OLMGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OLMGlobalInfo_KEYS.Size(m)
}
func (m *OLMGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OLMGlobalInfo_KEYS proto.InternalMessageInfo

type OLMAddrU struct {
	AddressType          string   `protobuf:"bytes,1,opt,name=address_type,json=addressType,proto3" json:"address_type,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	UnnumberedAddress    uint32   `protobuf:"varint,4,opt,name=unnumbered_address,json=unnumberedAddress,proto3" json:"unnumbered_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OLMAddrU) Reset()         { *m = OLMAddrU{} }
func (m *OLMAddrU) String() string { return proto.CompactTextString(m) }
func (*OLMAddrU) ProtoMessage()    {}
func (*OLMAddrU) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bda31b4ffaf29a, []int{1}
}

func (m *OLMAddrU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMAddrU.Unmarshal(m, b)
}
func (m *OLMAddrU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMAddrU.Marshal(b, m, deterministic)
}
func (m *OLMAddrU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMAddrU.Merge(m, src)
}
func (m *OLMAddrU) XXX_Size() int {
	return xxx_messageInfo_OLMAddrU.Size(m)
}
func (m *OLMAddrU) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMAddrU.DiscardUnknown(m)
}

var xxx_messageInfo_OLMAddrU proto.InternalMessageInfo

func (m *OLMAddrU) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *OLMAddrU) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *OLMAddrU) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *OLMAddrU) GetUnnumberedAddress() uint32 {
	if m != nil {
		return m.UnnumberedAddress
	}
	return 0
}

type OLMAddr_ struct {
	Address              *OLMAddrU `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OLMAddr_) Reset()         { *m = OLMAddr_{} }
func (m *OLMAddr_) String() string { return proto.CompactTextString(m) }
func (*OLMAddr_) ProtoMessage()    {}
func (*OLMAddr_) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bda31b4ffaf29a, []int{2}
}

func (m *OLMAddr_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMAddr_.Unmarshal(m, b)
}
func (m *OLMAddr_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMAddr_.Marshal(b, m, deterministic)
}
func (m *OLMAddr_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMAddr_.Merge(m, src)
}
func (m *OLMAddr_) XXX_Size() int {
	return xxx_messageInfo_OLMAddr_.Size(m)
}
func (m *OLMAddr_) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMAddr_.DiscardUnknown(m)
}

var xxx_messageInfo_OLMAddr_ proto.InternalMessageInfo

func (m *OLMAddr_) GetAddress() *OLMAddrU {
	if m != nil {
		return m.Address
	}
	return nil
}

type OLMGlobalInfo struct {
	LocalOuniLmpNodeId          *OLMAddr_ `protobuf:"bytes,50,opt,name=local_ouni_lmp_node_id,json=localOuniLmpNodeId,proto3" json:"local_ouni_lmp_node_id,omitempty"`
	LocalMplsTeLmpNodeId        *OLMAddr_ `protobuf:"bytes,51,opt,name=local_mpls_te_lmp_node_id,json=localMplsTeLmpNodeId,proto3" json:"local_mpls_te_lmp_node_id,omitempty"`
	LocalGmplsUniLmpNodeId      *OLMAddr_ `protobuf:"bytes,52,opt,name=local_gmpls_uni_lmp_node_id,json=localGmplsUniLmpNodeId,proto3" json:"local_gmpls_uni_lmp_node_id,omitempty"`
	LocalOuniLmpNodeIdInterface string    `protobuf:"bytes,53,opt,name=local_ouni_lmp_node_id_interface,json=localOuniLmpNodeIdInterface,proto3" json:"local_ouni_lmp_node_id_interface,omitempty"`
	LocalOuniLmpNodeIdType      string    `protobuf:"bytes,54,opt,name=local_ouni_lmp_node_id_type,json=localOuniLmpNodeIdType,proto3" json:"local_ouni_lmp_node_id_type,omitempty"`
	IsOuniConfigExist           bool      `protobuf:"varint,55,opt,name=is_ouni_config_exist,json=isOuniConfigExist,proto3" json:"is_ouni_config_exist,omitempty"`
	IsGmplsNniConfigExist       bool      `protobuf:"varint,56,opt,name=is_gmpls_nni_config_exist,json=isGmplsNniConfigExist,proto3" json:"is_gmpls_nni_config_exist,omitempty"`
	IsGmplsUniConfigExist       bool      `protobuf:"varint,57,opt,name=is_gmpls_uni_config_exist,json=isGmplsUniConfigExist,proto3" json:"is_gmpls_uni_config_exist,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}  `json:"-"`
	XXX_unrecognized            []byte    `json:"-"`
	XXX_sizecache               int32     `json:"-"`
}

func (m *OLMGlobalInfo) Reset()         { *m = OLMGlobalInfo{} }
func (m *OLMGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*OLMGlobalInfo) ProtoMessage()    {}
func (*OLMGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_79bda31b4ffaf29a, []int{3}
}

func (m *OLMGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OLMGlobalInfo.Unmarshal(m, b)
}
func (m *OLMGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OLMGlobalInfo.Marshal(b, m, deterministic)
}
func (m *OLMGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OLMGlobalInfo.Merge(m, src)
}
func (m *OLMGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_OLMGlobalInfo.Size(m)
}
func (m *OLMGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OLMGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OLMGlobalInfo proto.InternalMessageInfo

func (m *OLMGlobalInfo) GetLocalOuniLmpNodeId() *OLMAddr_ {
	if m != nil {
		return m.LocalOuniLmpNodeId
	}
	return nil
}

func (m *OLMGlobalInfo) GetLocalMplsTeLmpNodeId() *OLMAddr_ {
	if m != nil {
		return m.LocalMplsTeLmpNodeId
	}
	return nil
}

func (m *OLMGlobalInfo) GetLocalGmplsUniLmpNodeId() *OLMAddr_ {
	if m != nil {
		return m.LocalGmplsUniLmpNodeId
	}
	return nil
}

func (m *OLMGlobalInfo) GetLocalOuniLmpNodeIdInterface() string {
	if m != nil {
		return m.LocalOuniLmpNodeIdInterface
	}
	return ""
}

func (m *OLMGlobalInfo) GetLocalOuniLmpNodeIdType() string {
	if m != nil {
		return m.LocalOuniLmpNodeIdType
	}
	return ""
}

func (m *OLMGlobalInfo) GetIsOuniConfigExist() bool {
	if m != nil {
		return m.IsOuniConfigExist
	}
	return false
}

func (m *OLMGlobalInfo) GetIsGmplsNniConfigExist() bool {
	if m != nil {
		return m.IsGmplsNniConfigExist
	}
	return false
}

func (m *OLMGlobalInfo) GetIsGmplsUniConfigExist() bool {
	if m != nil {
		return m.IsGmplsUniConfigExist
	}
	return false
}

func init() {
	proto.RegisterType((*OLMGlobalInfo_KEYS)(nil), "cisco_ios_xr_lmp_oper.lmp.global_status.OLMGlobalInfo_KEYS")
	proto.RegisterType((*OLMAddrU)(nil), "cisco_ios_xr_lmp_oper.lmp.global_status.OLMAddrU")
	proto.RegisterType((*OLMAddr_)(nil), "cisco_ios_xr_lmp_oper.lmp.global_status.OLM_addr_")
	proto.RegisterType((*OLMGlobalInfo)(nil), "cisco_ios_xr_lmp_oper.lmp.global_status.OLMGlobalInfo")
}

func init() { proto.RegisterFile("OLMGlobalInfo.proto", fileDescriptor_79bda31b4ffaf29a) }

var fileDescriptor_79bda31b4ffaf29a = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x51, 0xaf, 0xd2, 0x30,
	0x1c, 0xc5, 0x33, 0xbd, 0xd1, 0x7b, 0x7b, 0xbd, 0x0f, 0xb7, 0xe2, 0xcd, 0x08, 0x2f, 0x73, 0x2f,
	0xf2, 0xe2, 0x8c, 0x80, 0x88, 0xf1, 0x89, 0x18, 0x42, 0x08, 0x1b, 0x4b, 0x26, 0x24, 0xfa, 0xd4,
	0x8c, 0xb5, 0x23, 0x8d, 0x5d, 0xdb, 0xac, 0x9b, 0x81, 0xaf, 0xe3, 0x07, 0x35, 0x66, 0xdd, 0x06,
	0x8e, 0x69, 0xa2, 0xbc, 0xfe, 0xfb, 0x3b, 0xe7, 0xf4, 0x7f, 0xd2, 0x82, 0xe7, 0xbe, 0xeb, 0xcd,
	0x99, 0xd8, 0x86, 0x6c, 0xc1, 0x63, 0xe1, 0xc8, 0x54, 0x64, 0x02, 0xbe, 0x8a, 0xa8, 0x8a, 0x04,
	0xa2, 0x42, 0xa1, 0x7d, 0x8a, 0x58, 0x22, 0x91, 0x90, 0x24, 0x75, 0x58, 0x22, 0x9d, 0x9d, 0x66,
	0x91, 0xca, 0xc2, 0x2c, 0x57, 0x76, 0x07, 0xc0, 0x86, 0x1e, 0x2d, 0x67, 0x5f, 0x3f, 0xdb, 0x3f,
	0x0c, 0x70, 0xed, 0xbb, 0xde, 0x14, 0xe3, 0x74, 0x03, 0x5f, 0x82, 0x67, 0x21, 0xc6, 0x29, 0x51,
	0x0a, 0x65, 0x07, 0x49, 0x4c, 0xc3, 0x32, 0xfa, 0x37, 0xc1, 0x6d, 0x35, 0x5b, 0x1f, 0x24, 0x29,
	0x10, 0x2a, 0xbf, 0x8f, 0x50, 0x35, 0x33, 0x1f, 0x95, 0x48, 0x31, 0x9b, 0x96, 0xa3, 0x0a, 0x19,
	0x1f, 0x91, 0xc7, 0x47, 0x64, 0x5c, 0x23, 0xaf, 0x01, 0xcc, 0x39, 0xcf, 0x93, 0x2d, 0x49, 0x09,
	0x3e, 0x82, 0x57, 0x96, 0xd1, 0xbf, 0x0b, 0xee, 0x4f, 0x27, 0x15, 0x6e, 0x7f, 0x01, 0x37, 0xbe,
	0xeb, 0x69, 0x0e, 0xc1, 0x25, 0x78, 0x5a, 0x0b, 0x8a, 0xfb, 0xdd, 0x0e, 0xde, 0x3a, 0xff, 0x58,
	0x81, 0x53, 0x2f, 0x1a, 0xd4, 0x0e, 0xf6, 0xcf, 0x2b, 0x70, 0xd7, 0x68, 0x05, 0xc6, 0xe0, 0x81,
	0x89, 0x28, 0x64, 0x48, 0xe4, 0x9c, 0x6a, 0x33, 0x2e, 0x30, 0x41, 0x14, 0x9b, 0x03, 0x9d, 0x36,
	0xf8, 0x9f, 0xb4, 0xf2, 0xca, 0x01, 0xd4, 0x8e, 0x7e, 0xce, 0xa9, 0x9b, 0xc8, 0x95, 0xc0, 0x64,
	0x81, 0xe1, 0x37, 0xd0, 0x2d, 0x73, 0x12, 0xc9, 0x14, 0xca, 0x48, 0x23, 0x6a, 0x78, 0x71, 0x54,
	0x47, 0x9b, 0x7a, 0x92, 0xa9, 0x35, 0x39, 0x85, 0x09, 0xd0, 0x2b, 0xc3, 0x76, 0x3a, 0xed, 0x7c,
	0xb3, 0xd1, 0xc5, 0x71, 0x65, 0x57, 0xf3, 0xc2, 0x75, 0xf3, 0xfb, 0x76, 0x33, 0x60, 0xfd, 0xb9,
	0x45, 0x44, 0x79, 0x46, 0xd2, 0x38, 0x8c, 0x88, 0xf9, 0x4e, 0xbf, 0x8b, 0x5e, 0xbb, 0x9b, 0x45,
	0x8d, 0xc0, 0x8f, 0xf5, 0xbd, 0x5b, 0x36, 0xfa, 0x7d, 0x8e, 0xb5, 0xc3, 0x43, 0xdb, 0x41, 0x3f,
	0xd5, 0x37, 0xa0, 0x43, 0x55, 0xa9, 0x8c, 0x04, 0x8f, 0xe9, 0x0e, 0x91, 0x3d, 0x55, 0x99, 0xf9,
	0xde, 0x32, 0xfa, 0xd7, 0xc1, 0x3d, 0x55, 0x85, 0xe4, 0x93, 0x3e, 0x99, 0x15, 0x07, 0x70, 0x02,
	0xba, 0x54, 0x55, 0x15, 0xf1, 0x73, 0xd5, 0x44, 0xab, 0x5e, 0x50, 0xa5, 0x97, 0x5d, 0xfd, 0x5d,
	0xd9, 0xca, 0xfb, 0xd0, 0x50, 0x6e, 0x1a, 0xca, 0xed, 0x13, 0xfd, 0x8b, 0x87, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xa6, 0xee, 0x0d, 0xdb, 0xdc, 0x03, 0x00, 0x00,
}
