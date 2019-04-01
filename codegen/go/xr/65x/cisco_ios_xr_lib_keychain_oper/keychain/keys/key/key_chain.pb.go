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
// source: key_chain.proto

package cisco_ios_xr_lib_keychain_oper_keychain_keys_key

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

type KeyChain_KEYS struct {
	KeyName              string   `protobuf:"bytes,1,opt,name=key_name,json=keyName,proto3" json:"key_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyChain_KEYS) Reset()         { *m = KeyChain_KEYS{} }
func (m *KeyChain_KEYS) String() string { return proto.CompactTextString(m) }
func (*KeyChain_KEYS) ProtoMessage()    {}
func (*KeyChain_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{0}
}

func (m *KeyChain_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyChain_KEYS.Unmarshal(m, b)
}
func (m *KeyChain_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyChain_KEYS.Marshal(b, m, deterministic)
}
func (m *KeyChain_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyChain_KEYS.Merge(m, src)
}
func (m *KeyChain_KEYS) XXX_Size() int {
	return xxx_messageInfo_KeyChain_KEYS.Size(m)
}
func (m *KeyChain_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyChain_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_KeyChain_KEYS proto.InternalMessageInfo

func (m *KeyChain_KEYS) GetKeyName() string {
	if m != nil {
		return m.KeyName
	}
	return ""
}

type Macsec struct {
	IsMacsecKey          bool     `protobuf:"varint,1,opt,name=is_macsec_key,json=isMacsecKey,proto3" json:"is_macsec_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Macsec) Reset()         { *m = Macsec{} }
func (m *Macsec) String() string { return proto.CompactTextString(m) }
func (*Macsec) ProtoMessage()    {}
func (*Macsec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{1}
}

func (m *Macsec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Macsec.Unmarshal(m, b)
}
func (m *Macsec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Macsec.Marshal(b, m, deterministic)
}
func (m *Macsec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Macsec.Merge(m, src)
}
func (m *Macsec) XXX_Size() int {
	return xxx_messageInfo_Macsec.Size(m)
}
func (m *Macsec) XXX_DiscardUnknown() {
	xxx_messageInfo_Macsec.DiscardUnknown(m)
}

var xxx_messageInfo_Macsec proto.InternalMessageInfo

func (m *Macsec) GetIsMacsecKey() bool {
	if m != nil {
		return m.IsMacsecKey
	}
	return false
}

type Lifetime struct {
	Start                string   `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  string   `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	Duration             string   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	IsAlwaysValid        bool     `protobuf:"varint,4,opt,name=is_always_valid,json=isAlwaysValid,proto3" json:"is_always_valid,omitempty"`
	IsValidNow           bool     `protobuf:"varint,5,opt,name=is_valid_now,json=isValidNow,proto3" json:"is_valid_now,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lifetime) Reset()         { *m = Lifetime{} }
func (m *Lifetime) String() string { return proto.CompactTextString(m) }
func (*Lifetime) ProtoMessage()    {}
func (*Lifetime) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{2}
}

func (m *Lifetime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lifetime.Unmarshal(m, b)
}
func (m *Lifetime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lifetime.Marshal(b, m, deterministic)
}
func (m *Lifetime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lifetime.Merge(m, src)
}
func (m *Lifetime) XXX_Size() int {
	return xxx_messageInfo_Lifetime.Size(m)
}
func (m *Lifetime) XXX_DiscardUnknown() {
	xxx_messageInfo_Lifetime.DiscardUnknown(m)
}

var xxx_messageInfo_Lifetime proto.InternalMessageInfo

func (m *Lifetime) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Lifetime) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Lifetime) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *Lifetime) GetIsAlwaysValid() bool {
	if m != nil {
		return m.IsAlwaysValid
	}
	return false
}

func (m *Lifetime) GetIsValidNow() bool {
	if m != nil {
		return m.IsValidNow
	}
	return false
}

type KeyIdItem struct {
	Macsec                 *Macsec   `protobuf:"bytes,1,opt,name=macsec,proto3" json:"macsec,omitempty"`
	KeyString              string    `protobuf:"bytes,2,opt,name=key_string,json=keyString,proto3" json:"key_string,omitempty"`
	Type                   string    `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	KeyId                  string    `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	CryptographicAlgorithm string    `protobuf:"bytes,5,opt,name=cryptographic_algorithm,json=cryptographicAlgorithm,proto3" json:"cryptographic_algorithm,omitempty"`
	SendLifetime           *Lifetime `protobuf:"bytes,6,opt,name=send_lifetime,json=sendLifetime,proto3" json:"send_lifetime,omitempty"`
	AcceptLifetime         *Lifetime `protobuf:"bytes,7,opt,name=accept_lifetime,json=acceptLifetime,proto3" json:"accept_lifetime,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}  `json:"-"`
	XXX_unrecognized       []byte    `json:"-"`
	XXX_sizecache          int32     `json:"-"`
}

func (m *KeyIdItem) Reset()         { *m = KeyIdItem{} }
func (m *KeyIdItem) String() string { return proto.CompactTextString(m) }
func (*KeyIdItem) ProtoMessage()    {}
func (*KeyIdItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{3}
}

func (m *KeyIdItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyIdItem.Unmarshal(m, b)
}
func (m *KeyIdItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyIdItem.Marshal(b, m, deterministic)
}
func (m *KeyIdItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyIdItem.Merge(m, src)
}
func (m *KeyIdItem) XXX_Size() int {
	return xxx_messageInfo_KeyIdItem.Size(m)
}
func (m *KeyIdItem) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyIdItem.DiscardUnknown(m)
}

var xxx_messageInfo_KeyIdItem proto.InternalMessageInfo

func (m *KeyIdItem) GetMacsec() *Macsec {
	if m != nil {
		return m.Macsec
	}
	return nil
}

func (m *KeyIdItem) GetKeyString() string {
	if m != nil {
		return m.KeyString
	}
	return ""
}

func (m *KeyIdItem) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *KeyIdItem) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *KeyIdItem) GetCryptographicAlgorithm() string {
	if m != nil {
		return m.CryptographicAlgorithm
	}
	return ""
}

func (m *KeyIdItem) GetSendLifetime() *Lifetime {
	if m != nil {
		return m.SendLifetime
	}
	return nil
}

func (m *KeyIdItem) GetAcceptLifetime() *Lifetime {
	if m != nil {
		return m.AcceptLifetime
	}
	return nil
}

type KeyIdEntry struct {
	KeyId                []*KeyIdItem `protobuf:"bytes,1,rep,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *KeyIdEntry) Reset()         { *m = KeyIdEntry{} }
func (m *KeyIdEntry) String() string { return proto.CompactTextString(m) }
func (*KeyIdEntry) ProtoMessage()    {}
func (*KeyIdEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{4}
}

func (m *KeyIdEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyIdEntry.Unmarshal(m, b)
}
func (m *KeyIdEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyIdEntry.Marshal(b, m, deterministic)
}
func (m *KeyIdEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyIdEntry.Merge(m, src)
}
func (m *KeyIdEntry) XXX_Size() int {
	return xxx_messageInfo_KeyIdEntry.Size(m)
}
func (m *KeyIdEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyIdEntry.DiscardUnknown(m)
}

var xxx_messageInfo_KeyIdEntry proto.InternalMessageInfo

func (m *KeyIdEntry) GetKeyId() []*KeyIdItem {
	if m != nil {
		return m.KeyId
	}
	return nil
}

type KeyChain struct {
	AcceptTolerance      string      `protobuf:"bytes,50,opt,name=accept_tolerance,json=acceptTolerance,proto3" json:"accept_tolerance,omitempty"`
	Key                  *KeyIdEntry `protobuf:"bytes,51,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KeyChain) Reset()         { *m = KeyChain{} }
func (m *KeyChain) String() string { return proto.CompactTextString(m) }
func (*KeyChain) ProtoMessage()    {}
func (*KeyChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf8417ec4d6f0e9, []int{5}
}

func (m *KeyChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyChain.Unmarshal(m, b)
}
func (m *KeyChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyChain.Marshal(b, m, deterministic)
}
func (m *KeyChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyChain.Merge(m, src)
}
func (m *KeyChain) XXX_Size() int {
	return xxx_messageInfo_KeyChain.Size(m)
}
func (m *KeyChain) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyChain.DiscardUnknown(m)
}

var xxx_messageInfo_KeyChain proto.InternalMessageInfo

func (m *KeyChain) GetAcceptTolerance() string {
	if m != nil {
		return m.AcceptTolerance
	}
	return ""
}

func (m *KeyChain) GetKey() *KeyIdEntry {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyChain_KEYS)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.key_chain_KEYS")
	proto.RegisterType((*Macsec)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.macsec")
	proto.RegisterType((*Lifetime)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.lifetime")
	proto.RegisterType((*KeyIdItem)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.key_id_item")
	proto.RegisterType((*KeyIdEntry)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.key_id_entry")
	proto.RegisterType((*KeyChain)(nil), "cisco_ios_xr_lib_keychain_oper.keychain.keys.key.key_chain")
}

func init() { proto.RegisterFile("key_chain.proto", fileDescriptor_5bf8417ec4d6f0e9) }

var fileDescriptor_5bf8417ec4d6f0e9 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0xc7, 0x89, 0xb3, 0xc9, 0x66, 0x2a, 0xc9, 0x66, 0x69, 0xfc, 0x18, 0x05, 0x21, 0xcc, 0x41,
	0x56, 0x94, 0x20, 0xd9, 0x83, 0x22, 0x28, 0xec, 0xc1, 0x83, 0xac, 0x2e, 0xcb, 0xec, 0x22, 0x78,
	0x6a, 0x7a, 0x7b, 0xca, 0xa4, 0x98, 0x8f, 0x1e, 0xba, 0x5b, 0x63, 0xbf, 0x81, 0x8f, 0xe0, 0x23,
	0xf8, 0x98, 0xd2, 0x3d, 0x33, 0x51, 0x8f, 0x71, 0x2f, 0xa1, 0xeb, 0xdf, 0x95, 0xfa, 0xd7, 0xaf,
	0x6a, 0x1a, 0xe6, 0x05, 0x3a, 0x2e, 0x37, 0x82, 0xea, 0x65, 0xa3, 0x95, 0x55, 0xec, 0x85, 0x24,
	0x23, 0x15, 0x27, 0x65, 0xf8, 0x77, 0xcd, 0x4b, 0xba, 0xe1, 0x05, 0xba, 0x90, 0xc0, 0x55, 0x83,
	0x7a, 0xd9, 0x47, 0xfe, 0x60, 0xfc, 0x4f, 0xfa, 0x0c, 0x8e, 0x76, 0x45, 0xf8, 0xf9, 0xbb, 0xcf,
	0x57, 0xec, 0x21, 0x8c, 0xbd, 0x52, 0x8b, 0x0a, 0x93, 0xc1, 0x62, 0x70, 0x12, 0x67, 0x87, 0x05,
	0xba, 0x0b, 0x51, 0x61, 0xfa, 0x1c, 0x46, 0x95, 0x90, 0x06, 0x25, 0x4b, 0x61, 0x46, 0x86, 0xb7,
	0x81, 0xf7, 0x08, 0x99, 0xe3, 0x6c, 0x42, 0xe6, 0x63, 0xd0, 0xce, 0xd1, 0xa5, 0x3f, 0x07, 0x30,
	0x2e, 0xe9, 0x0b, 0x5a, 0xaa, 0x90, 0xdd, 0x85, 0xa1, 0xb1, 0x42, 0xdb, 0xae, 0x64, 0x1b, 0xb0,
	0x63, 0x88, 0xb0, 0xce, 0x93, 0x3b, 0x41, 0xf3, 0x47, 0xf6, 0x08, 0xc6, 0xf9, 0x57, 0x2d, 0x2c,
	0xa9, 0x3a, 0x89, 0x82, 0xbc, 0x8b, 0xd9, 0x13, 0x98, 0x93, 0xe1, 0xa2, 0xdc, 0x0a, 0x67, 0xf8,
	0x37, 0x51, 0x52, 0x9e, 0x1c, 0x04, 0xdb, 0x19, 0x99, 0xb3, 0xa0, 0x7e, 0xf2, 0x22, 0x5b, 0xc0,
	0x94, 0xba, 0x04, 0x5e, 0xab, 0x6d, 0x32, 0x0c, 0x49, 0x40, 0xed, 0xf5, 0x85, 0xda, 0xa6, 0xbf,
	0x22, 0x98, 0x78, 0x48, 0xca, 0x39, 0x59, 0xac, 0xd8, 0x65, 0x0f, 0x16, 0xda, 0x9b, 0xac, 0x5e,
	0x2d, 0xf7, 0x1d, 0xe4, 0xb2, 0xfd, 0x7f, 0xd6, 0x0f, 0xe8, 0x31, 0x80, 0x37, 0x30, 0x56, 0x53,
	0xbd, 0xee, 0x00, 0xe3, 0x02, 0xdd, 0x55, 0x10, 0x18, 0x83, 0x03, 0xeb, 0x1a, 0xec, 0x10, 0xc3,
	0x99, 0xdd, 0x83, 0x51, 0xdb, 0x53, 0xa0, 0x8a, 0xb3, 0x61, 0x81, 0xee, 0x7d, 0xce, 0x5e, 0xc2,
	0x03, 0xa9, 0x5d, 0x63, 0xd5, 0x5a, 0x8b, 0x66, 0x43, 0x92, 0x8b, 0x72, 0xad, 0x34, 0xd9, 0x4d,
	0x15, 0xc0, 0xe2, 0xec, 0xfe, 0x3f, 0xd7, 0x67, 0xfd, 0x2d, 0xe3, 0x30, 0x33, 0x58, 0xe7, 0xbc,
	0xdf, 0x41, 0x32, 0x0a, 0x6c, 0xaf, 0xf7, 0x67, 0xeb, 0x2b, 0x64, 0x53, 0x5f, 0xf0, 0x43, 0xbf,
	0x53, 0x09, 0x73, 0x21, 0x25, 0x36, 0xf6, 0x8f, 0xc5, 0xe1, 0xad, 0x2d, 0x8e, 0xda, 0x92, 0xbd,
	0x49, 0x9a, 0xc3, 0xb4, 0xdb, 0x14, 0xd6, 0x56, 0x3b, 0x76, 0xbd, 0x9b, 0xd2, 0x60, 0x11, 0x9d,
	0x4c, 0x56, 0x6f, 0xf6, 0xf7, 0xfa, 0x6b, 0xf3, 0xdd, 0x90, 0xd3, 0x1f, 0x03, 0x88, 0x77, 0xef,
	0x80, 0x3d, 0x85, 0xe3, 0x0e, 0xcc, 0xaa, 0x12, 0xb5, 0xa8, 0x25, 0x26, 0xab, 0x30, 0xeb, 0x0e,
	0xf8, 0xba, 0x97, 0xd9, 0x25, 0x44, 0xfe, 0xf3, 0x3f, 0x0d, 0xdc, 0x6f, 0xff, 0xbb, 0x97, 0xc0,
	0x96, 0xf9, 0x52, 0x37, 0xa3, 0xf0, 0x94, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x07,
	0x02, 0x7a, 0xdd, 0x03, 0x00, 0x00,
}