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
// source: crypto_sam_crl_info.proto

package cisco_ios_xr_crypto_sam_oper_sam_certificate_revocations_certificate_revocation_certificate_revocation_list_detail

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

type CryptoSamCrlInfo_KEYS struct {
	CrlIndex             uint32   `protobuf:"varint,1,opt,name=crl_index,json=crlIndex,proto3" json:"crl_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoSamCrlInfo_KEYS) Reset()         { *m = CryptoSamCrlInfo_KEYS{} }
func (m *CryptoSamCrlInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*CryptoSamCrlInfo_KEYS) ProtoMessage()    {}
func (*CryptoSamCrlInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_928337a268aff277, []int{0}
}

func (m *CryptoSamCrlInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoSamCrlInfo_KEYS.Unmarshal(m, b)
}
func (m *CryptoSamCrlInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoSamCrlInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *CryptoSamCrlInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoSamCrlInfo_KEYS.Merge(m, src)
}
func (m *CryptoSamCrlInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_CryptoSamCrlInfo_KEYS.Size(m)
}
func (m *CryptoSamCrlInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoSamCrlInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoSamCrlInfo_KEYS proto.InternalMessageInfo

func (m *CryptoSamCrlInfo_KEYS) GetCrlIndex() uint32 {
	if m != nil {
		return m.CrlIndex
	}
	return 0
}

type IssuerInfo struct {
	CommonName           string   `protobuf:"bytes,1,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	Organization         string   `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	Country              string   `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssuerInfo) Reset()         { *m = IssuerInfo{} }
func (m *IssuerInfo) String() string { return proto.CompactTextString(m) }
func (*IssuerInfo) ProtoMessage()    {}
func (*IssuerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_928337a268aff277, []int{1}
}

func (m *IssuerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssuerInfo.Unmarshal(m, b)
}
func (m *IssuerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssuerInfo.Marshal(b, m, deterministic)
}
func (m *IssuerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssuerInfo.Merge(m, src)
}
func (m *IssuerInfo) XXX_Size() int {
	return xxx_messageInfo_IssuerInfo.Size(m)
}
func (m *IssuerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IssuerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IssuerInfo proto.InternalMessageInfo

func (m *IssuerInfo) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

func (m *IssuerInfo) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *IssuerInfo) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

type CryptoSamCrlInfo struct {
	CrlIndex             uint32      `protobuf:"varint,50,opt,name=crl_index,json=crlIndex,proto3" json:"crl_index,omitempty"`
	Issuer               *IssuerInfo `protobuf:"bytes,51,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Updates              string      `protobuf:"bytes,52,opt,name=updates,proto3" json:"updates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CryptoSamCrlInfo) Reset()         { *m = CryptoSamCrlInfo{} }
func (m *CryptoSamCrlInfo) String() string { return proto.CompactTextString(m) }
func (*CryptoSamCrlInfo) ProtoMessage()    {}
func (*CryptoSamCrlInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_928337a268aff277, []int{2}
}

func (m *CryptoSamCrlInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoSamCrlInfo.Unmarshal(m, b)
}
func (m *CryptoSamCrlInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoSamCrlInfo.Marshal(b, m, deterministic)
}
func (m *CryptoSamCrlInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoSamCrlInfo.Merge(m, src)
}
func (m *CryptoSamCrlInfo) XXX_Size() int {
	return xxx_messageInfo_CryptoSamCrlInfo.Size(m)
}
func (m *CryptoSamCrlInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoSamCrlInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoSamCrlInfo proto.InternalMessageInfo

func (m *CryptoSamCrlInfo) GetCrlIndex() uint32 {
	if m != nil {
		return m.CrlIndex
	}
	return 0
}

func (m *CryptoSamCrlInfo) GetIssuer() *IssuerInfo {
	if m != nil {
		return m.Issuer
	}
	return nil
}

func (m *CryptoSamCrlInfo) GetUpdates() string {
	if m != nil {
		return m.Updates
	}
	return ""
}

func init() {
	proto.RegisterType((*CryptoSamCrlInfo_KEYS)(nil), "cisco_ios_xr_crypto_sam_oper.sam.certificate_revocations.certificate_revocation.certificate_revocation_list_detail.crypto_sam_crl_info_KEYS")
	proto.RegisterType((*IssuerInfo)(nil), "cisco_ios_xr_crypto_sam_oper.sam.certificate_revocations.certificate_revocation.certificate_revocation_list_detail.issuer_info")
	proto.RegisterType((*CryptoSamCrlInfo)(nil), "cisco_ios_xr_crypto_sam_oper.sam.certificate_revocations.certificate_revocation.certificate_revocation_list_detail.crypto_sam_crl_info")
}

func init() { proto.RegisterFile("crypto_sam_crl_info.proto", fileDescriptor_928337a268aff277) }

var fileDescriptor_928337a268aff277 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x85, 0x6a, 0xb3, 0x7a, 0x89, 0x97, 0x88, 0x07, 0xcb, 0x9e, 0x7a, 0xda, 0x43,
	0x2b, 0xf8, 0x04, 0x1e, 0x44, 0xf0, 0xb0, 0x9e, 0x3c, 0x0d, 0x31, 0x9b, 0xca, 0x40, 0x92, 0x59,
	0x26, 0x59, 0x69, 0xbd, 0xf8, 0x18, 0xbe, 0xa1, 0xcf, 0x21, 0xcd, 0x56, 0xe8, 0x42, 0x7b, 0xf5,
	0xf8, 0x7f, 0xf9, 0x03, 0xdf, 0x0c, 0x23, 0xae, 0x0d, 0x6f, 0xba, 0x44, 0x10, 0xb5, 0x07, 0xc3,
	0x0e, 0x30, 0xac, 0xa8, 0xee, 0x98, 0x12, 0x49, 0x36, 0x18, 0x0d, 0x01, 0x52, 0x84, 0x35, 0xc3,
	0x5e, 0x8f, 0x3a, 0xcb, 0x75, 0xd4, 0xbe, 0x36, 0x96, 0x13, 0xae, 0xd0, 0xe8, 0x64, 0x81, 0xed,
	0x07, 0x19, 0x9d, 0x90, 0x42, 0x3c, 0xc2, 0x8f, 0x60, 0x70, 0x18, 0x13, 0xb4, 0x36, 0x69, 0x74,
	0xd5, 0xbd, 0x50, 0x07, 0x84, 0xe0, 0xe9, 0xe1, 0xf5, 0x45, 0xde, 0x88, 0xe9, 0x00, 0x5a, 0xbb,
	0x56, 0xc5, 0xac, 0x98, 0x5f, 0x36, 0xe7, 0x86, 0xdd, 0xe3, 0x36, 0x57, 0x4e, 0x94, 0x18, 0x63,
	0x6f, 0x39, 0x7f, 0x90, 0xb7, 0xa2, 0x34, 0xe4, 0x3d, 0x05, 0x08, 0xda, 0xdb, 0xdc, 0x9e, 0x36,
	0x62, 0x40, 0xcf, 0xda, 0x5b, 0x59, 0x89, 0x0b, 0xe2, 0x77, 0x1d, 0xf0, 0x33, 0x4b, 0xa8, 0x93,
	0xdc, 0x18, 0x31, 0xa9, 0xc4, 0x99, 0xa1, 0x3e, 0x24, 0xde, 0xa8, 0xd3, 0xfc, 0xfc, 0x17, 0xab,
	0x9f, 0x42, 0x5c, 0x1d, 0xf0, 0x1c, 0x2b, 0x2e, 0xc6, 0x8a, 0xf2, 0xbb, 0x10, 0x93, 0xc1, 0x51,
	0x2d, 0x67, 0xc5, 0xbc, 0x5c, 0x7c, 0xd5, 0xff, 0xbf, 0xe1, 0x7a, 0x6f, 0x4b, 0xcd, 0x4e, 0x67,
	0x3b, 0x68, 0xdf, 0xb5, 0x3a, 0xd9, 0xa8, 0xee, 0x86, 0x41, 0x77, 0xf1, 0x6d, 0x92, 0x4f, 0x61,
	0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x88, 0x82, 0x95, 0x27, 0x02, 0x00, 0x00,
}
