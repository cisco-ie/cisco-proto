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
// source: crypto_sam_certificate_common.proto

package cisco_ios_xr_crypto_sam_oper_sam_packages_package

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

type CryptoSamCertificateCommon_KEYS struct {
	PackageName          string   `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoSamCertificateCommon_KEYS) Reset()         { *m = CryptoSamCertificateCommon_KEYS{} }
func (m *CryptoSamCertificateCommon_KEYS) String() string { return proto.CompactTextString(m) }
func (*CryptoSamCertificateCommon_KEYS) ProtoMessage()    {}
func (*CryptoSamCertificateCommon_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_38408d80371d2fcb, []int{0}
}

func (m *CryptoSamCertificateCommon_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoSamCertificateCommon_KEYS.Unmarshal(m, b)
}
func (m *CryptoSamCertificateCommon_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoSamCertificateCommon_KEYS.Marshal(b, m, deterministic)
}
func (m *CryptoSamCertificateCommon_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoSamCertificateCommon_KEYS.Merge(m, src)
}
func (m *CryptoSamCertificateCommon_KEYS) XXX_Size() int {
	return xxx_messageInfo_CryptoSamCertificateCommon_KEYS.Size(m)
}
func (m *CryptoSamCertificateCommon_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoSamCertificateCommon_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoSamCertificateCommon_KEYS proto.InternalMessageInfo

func (m *CryptoSamCertificateCommon_KEYS) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

type Flags struct {
	IsTrusted            bool     `protobuf:"varint,1,opt,name=is_trusted,json=isTrusted,proto3" json:"is_trusted,omitempty"`
	IsRevoked            bool     `protobuf:"varint,2,opt,name=is_revoked,json=isRevoked,proto3" json:"is_revoked,omitempty"`
	IsExpired            bool     `protobuf:"varint,3,opt,name=is_expired,json=isExpired,proto3" json:"is_expired,omitempty"`
	IsValidated          bool     `protobuf:"varint,4,opt,name=is_validated,json=isValidated,proto3" json:"is_validated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Flags) Reset()         { *m = Flags{} }
func (m *Flags) String() string { return proto.CompactTextString(m) }
func (*Flags) ProtoMessage()    {}
func (*Flags) Descriptor() ([]byte, []int) {
	return fileDescriptor_38408d80371d2fcb, []int{1}
}

func (m *Flags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flags.Unmarshal(m, b)
}
func (m *Flags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flags.Marshal(b, m, deterministic)
}
func (m *Flags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flags.Merge(m, src)
}
func (m *Flags) XXX_Size() int {
	return xxx_messageInfo_Flags.Size(m)
}
func (m *Flags) XXX_DiscardUnknown() {
	xxx_messageInfo_Flags.DiscardUnknown(m)
}

var xxx_messageInfo_Flags proto.InternalMessageInfo

func (m *Flags) GetIsTrusted() bool {
	if m != nil {
		return m.IsTrusted
	}
	return false
}

func (m *Flags) GetIsRevoked() bool {
	if m != nil {
		return m.IsRevoked
	}
	return false
}

func (m *Flags) GetIsExpired() bool {
	if m != nil {
		return m.IsExpired
	}
	return false
}

func (m *Flags) GetIsValidated() bool {
	if m != nil {
		return m.IsValidated
	}
	return false
}

type CryptoSamCertificateCommon struct {
	Location             string   `protobuf:"bytes,50,opt,name=location,proto3" json:"location,omitempty"`
	CertificateIndex     uint32   `protobuf:"varint,51,opt,name=certificate_index,json=certificateIndex,proto3" json:"certificate_index,omitempty"`
	CertificateFlags     *Flags   `protobuf:"bytes,52,opt,name=certificate_flags,json=certificateFlags,proto3" json:"certificate_flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptoSamCertificateCommon) Reset()         { *m = CryptoSamCertificateCommon{} }
func (m *CryptoSamCertificateCommon) String() string { return proto.CompactTextString(m) }
func (*CryptoSamCertificateCommon) ProtoMessage()    {}
func (*CryptoSamCertificateCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_38408d80371d2fcb, []int{2}
}

func (m *CryptoSamCertificateCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoSamCertificateCommon.Unmarshal(m, b)
}
func (m *CryptoSamCertificateCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoSamCertificateCommon.Marshal(b, m, deterministic)
}
func (m *CryptoSamCertificateCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoSamCertificateCommon.Merge(m, src)
}
func (m *CryptoSamCertificateCommon) XXX_Size() int {
	return xxx_messageInfo_CryptoSamCertificateCommon.Size(m)
}
func (m *CryptoSamCertificateCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoSamCertificateCommon.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoSamCertificateCommon proto.InternalMessageInfo

func (m *CryptoSamCertificateCommon) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *CryptoSamCertificateCommon) GetCertificateIndex() uint32 {
	if m != nil {
		return m.CertificateIndex
	}
	return 0
}

func (m *CryptoSamCertificateCommon) GetCertificateFlags() *Flags {
	if m != nil {
		return m.CertificateFlags
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoSamCertificateCommon_KEYS)(nil), "cisco_ios_xr_crypto_sam_oper.sam.packages.package.crypto_sam_certificate_common_KEYS")
	proto.RegisterType((*Flags)(nil), "cisco_ios_xr_crypto_sam_oper.sam.packages.package.flags")
	proto.RegisterType((*CryptoSamCertificateCommon)(nil), "cisco_ios_xr_crypto_sam_oper.sam.packages.package.crypto_sam_certificate_common")
}

func init() {
	proto.RegisterFile("crypto_sam_certificate_common.proto", fileDescriptor_38408d80371d2fcb)
}

var fileDescriptor_38408d80371d2fcb = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xc9, 0xff, 0xab, 0xcc, 0xa4, 0x0a, 0x9a, 0x55, 0x10, 0x06, 0x6a, 0xdd, 0x14, 0x84,
	0x82, 0x33, 0x2e, 0x7c, 0x81, 0x51, 0x44, 0x70, 0x51, 0x45, 0x70, 0x15, 0x62, 0x7a, 0x67, 0xb8,
	0x4c, 0xd3, 0x94, 0x24, 0x0e, 0xf5, 0x09, 0x7c, 0x3c, 0x5f, 0x49, 0x9a, 0xb6, 0xd2, 0xd9, 0x0c,
	0xb8, 0x4b, 0xbe, 0x73, 0xce, 0x4d, 0x0e, 0x97, 0x5e, 0x2a, 0xfb, 0x59, 0x7b, 0x23, 0x9c, 0xd4,
	0x42, 0x81, 0xf5, 0xb8, 0x42, 0x25, 0x3d, 0x08, 0x65, 0xb4, 0x36, 0x55, 0x56, 0x5b, 0xe3, 0x0d,
	0xbb, 0x56, 0xe8, 0x94, 0x11, 0x68, 0x9c, 0x68, 0xac, 0x18, 0x25, 0x4c, 0x0d, 0x36, 0x73, 0x52,
	0x67, 0xb5, 0x54, 0x1b, 0xb9, 0x06, 0x37, 0x1c, 0x92, 0x7b, 0x9a, 0xec, 0x9d, 0x2c, 0x1e, 0x97,
	0x6f, 0xcf, 0xec, 0x82, 0x1e, 0xf7, 0x01, 0x51, 0x49, 0x0d, 0x9c, 0xc4, 0x24, 0x9d, 0xe6, 0x51,
	0xcf, 0x9e, 0xa4, 0x86, 0xe4, 0x8b, 0xd0, 0xc3, 0x55, 0x29, 0xd7, 0x8e, 0xcd, 0x28, 0x45, 0x27,
	0xbc, 0xfd, 0x70, 0x1e, 0x8a, 0x60, 0x9d, 0xe4, 0x53, 0x74, 0x2f, 0x1d, 0xe8, 0x65, 0x0b, 0x5b,
	0xb3, 0x81, 0x82, 0xff, 0x1b, 0xe4, 0xbc, 0x03, 0xbd, 0x0c, 0x4d, 0x8d, 0x16, 0x0a, 0xfe, 0x7f,
	0x90, 0x97, 0x1d, 0x68, 0x7f, 0x82, 0x4e, 0x6c, 0x65, 0x89, 0x85, 0x6c, 0xc7, 0x1f, 0x04, 0x43,
	0x84, 0xee, 0x75, 0x40, 0xc9, 0x37, 0xa1, 0xb3, 0xbd, 0x9d, 0xd8, 0x39, 0x9d, 0x94, 0x46, 0x49,
	0x8f, 0xa6, 0xe2, 0xf3, 0x50, 0xe5, 0xf7, 0xce, 0xae, 0xe8, 0xd9, 0x38, 0x81, 0x55, 0x01, 0x0d,
	0x5f, 0xc4, 0x24, 0x3d, 0xc9, 0x4f, 0x47, 0xc2, 0x43, 0xcb, 0x19, 0xec, 0x9a, 0x43, 0x7f, 0x7e,
	0x13, 0x93, 0x34, 0x9a, 0xdf, 0x66, 0x7f, 0x5e, 0x46, 0x16, 0xf2, 0x3b, 0xcf, 0xdc, 0xb5, 0xe4,
	0xfd, 0x28, 0xac, 0x77, 0xf1, 0x13, 0x00, 0x00, 0xff, 0xff, 0x25, 0x05, 0x49, 0x26, 0x05, 0x02,
	0x00, 0x00,
}
