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
// source: tcp_sh_keychain_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_connection_nodes_node_keychains_keychain

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

type TcpShKeychainBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	KeychainName         string   `protobuf:"bytes,2,opt,name=keychain_name,json=keychainName,proto3" json:"keychain_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpShKeychainBag_KEYS) Reset()         { *m = TcpShKeychainBag_KEYS{} }
func (m *TcpShKeychainBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpShKeychainBag_KEYS) ProtoMessage()    {}
func (*TcpShKeychainBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76af125e7882aca, []int{0}
}

func (m *TcpShKeychainBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShKeychainBag_KEYS.Unmarshal(m, b)
}
func (m *TcpShKeychainBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShKeychainBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpShKeychainBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShKeychainBag_KEYS.Merge(m, src)
}
func (m *TcpShKeychainBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpShKeychainBag_KEYS.Size(m)
}
func (m *TcpShKeychainBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShKeychainBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShKeychainBag_KEYS proto.InternalMessageInfo

func (m *TcpShKeychainBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpShKeychainBag_KEYS) GetKeychainName() string {
	if m != nil {
		return m.KeychainName
	}
	return ""
}

type TcpShKeyIdBag struct {
	KeyId                uint64   `protobuf:"varint,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpShKeyIdBag) Reset()         { *m = TcpShKeyIdBag{} }
func (m *TcpShKeyIdBag) String() string { return proto.CompactTextString(m) }
func (*TcpShKeyIdBag) ProtoMessage()    {}
func (*TcpShKeyIdBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76af125e7882aca, []int{1}
}

func (m *TcpShKeyIdBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShKeyIdBag.Unmarshal(m, b)
}
func (m *TcpShKeyIdBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShKeyIdBag.Marshal(b, m, deterministic)
}
func (m *TcpShKeyIdBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShKeyIdBag.Merge(m, src)
}
func (m *TcpShKeyIdBag) XXX_Size() int {
	return xxx_messageInfo_TcpShKeyIdBag.Size(m)
}
func (m *TcpShKeyIdBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShKeyIdBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShKeyIdBag proto.InternalMessageInfo

func (m *TcpShKeyIdBag) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

type TcpShKeyBag struct {
	KeyId                   uint64           `protobuf:"varint,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	IsActive                bool             `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	IsExpired               bool             `protobuf:"varint,3,opt,name=is_expired,json=isExpired,proto3" json:"is_expired,omitempty"`
	IsValid                 bool             `protobuf:"varint,4,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	Reason                  string           `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	SendId                  uint32           `protobuf:"varint,6,opt,name=send_id,json=sendId,proto3" json:"send_id,omitempty"`
	RecvId                  uint32           `protobuf:"varint,7,opt,name=recv_id,json=recvId,proto3" json:"recv_id,omitempty"`
	CryptAlgo               string           `protobuf:"bytes,8,opt,name=crypt_algo,json=cryptAlgo,proto3" json:"crypt_algo,omitempty"`
	IsConfigured            bool             `protobuf:"varint,9,opt,name=is_configured,json=isConfigured,proto3" json:"is_configured,omitempty"`
	InvalidatedKey          []*TcpShKeyIdBag `protobuf:"bytes,10,rep,name=invalidated_key,json=invalidatedKey,proto3" json:"invalidated_key,omitempty"`
	OverlappingKeyAvailable bool             `protobuf:"varint,11,opt,name=overlapping_key_available,json=overlappingKeyAvailable,proto3" json:"overlapping_key_available,omitempty"`
	OverlappingKey          uint64           `protobuf:"varint,12,opt,name=overlapping_key,json=overlappingKey,proto3" json:"overlapping_key,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *TcpShKeyBag) Reset()         { *m = TcpShKeyBag{} }
func (m *TcpShKeyBag) String() string { return proto.CompactTextString(m) }
func (*TcpShKeyBag) ProtoMessage()    {}
func (*TcpShKeyBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76af125e7882aca, []int{2}
}

func (m *TcpShKeyBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShKeyBag.Unmarshal(m, b)
}
func (m *TcpShKeyBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShKeyBag.Marshal(b, m, deterministic)
}
func (m *TcpShKeyBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShKeyBag.Merge(m, src)
}
func (m *TcpShKeyBag) XXX_Size() int {
	return xxx_messageInfo_TcpShKeyBag.Size(m)
}
func (m *TcpShKeyBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShKeyBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShKeyBag proto.InternalMessageInfo

func (m *TcpShKeyBag) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *TcpShKeyBag) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *TcpShKeyBag) GetIsExpired() bool {
	if m != nil {
		return m.IsExpired
	}
	return false
}

func (m *TcpShKeyBag) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *TcpShKeyBag) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *TcpShKeyBag) GetSendId() uint32 {
	if m != nil {
		return m.SendId
	}
	return 0
}

func (m *TcpShKeyBag) GetRecvId() uint32 {
	if m != nil {
		return m.RecvId
	}
	return 0
}

func (m *TcpShKeyBag) GetCryptAlgo() string {
	if m != nil {
		return m.CryptAlgo
	}
	return ""
}

func (m *TcpShKeyBag) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *TcpShKeyBag) GetInvalidatedKey() []*TcpShKeyIdBag {
	if m != nil {
		return m.InvalidatedKey
	}
	return nil
}

func (m *TcpShKeyBag) GetOverlappingKeyAvailable() bool {
	if m != nil {
		return m.OverlappingKeyAvailable
	}
	return false
}

func (m *TcpShKeyBag) GetOverlappingKey() uint64 {
	if m != nil {
		return m.OverlappingKey
	}
	return 0
}

type TcpShIdBag struct {
	Id                   uint32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Keys                 []*TcpShKeyIdBag `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TcpShIdBag) Reset()         { *m = TcpShIdBag{} }
func (m *TcpShIdBag) String() string { return proto.CompactTextString(m) }
func (*TcpShIdBag) ProtoMessage()    {}
func (*TcpShIdBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76af125e7882aca, []int{3}
}

func (m *TcpShIdBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShIdBag.Unmarshal(m, b)
}
func (m *TcpShIdBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShIdBag.Marshal(b, m, deterministic)
}
func (m *TcpShIdBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShIdBag.Merge(m, src)
}
func (m *TcpShIdBag) XXX_Size() int {
	return xxx_messageInfo_TcpShIdBag.Size(m)
}
func (m *TcpShIdBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShIdBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShIdBag proto.InternalMessageInfo

func (m *TcpShIdBag) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TcpShIdBag) GetKeys() []*TcpShKeyIdBag {
	if m != nil {
		return m.Keys
	}
	return nil
}

type TcpShKeychainBag struct {
	ChainName            string           `protobuf:"bytes,50,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	IsConfigured         bool             `protobuf:"varint,51,opt,name=is_configured,json=isConfigured,proto3" json:"is_configured,omitempty"`
	DesiredKeyAvailable  bool             `protobuf:"varint,52,opt,name=desired_key_available,json=desiredKeyAvailable,proto3" json:"desired_key_available,omitempty"`
	DesiredKeyId         uint64           `protobuf:"varint,53,opt,name=desired_key_id,json=desiredKeyId,proto3" json:"desired_key_id,omitempty"`
	Keys                 []*TcpShKeyBag   `protobuf:"bytes,54,rep,name=keys,proto3" json:"keys,omitempty"`
	ActiveKey            []*TcpShKeyIdBag `protobuf:"bytes,55,rep,name=active_key,json=activeKey,proto3" json:"active_key,omitempty"`
	SendId               []*TcpShIdBag    `protobuf:"bytes,56,rep,name=send_id,json=sendId,proto3" json:"send_id,omitempty"`
	ReceiveId            []*TcpShIdBag    `protobuf:"bytes,57,rep,name=receive_id,json=receiveId,proto3" json:"receive_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TcpShKeychainBag) Reset()         { *m = TcpShKeychainBag{} }
func (m *TcpShKeychainBag) String() string { return proto.CompactTextString(m) }
func (*TcpShKeychainBag) ProtoMessage()    {}
func (*TcpShKeychainBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76af125e7882aca, []int{4}
}

func (m *TcpShKeychainBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShKeychainBag.Unmarshal(m, b)
}
func (m *TcpShKeychainBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShKeychainBag.Marshal(b, m, deterministic)
}
func (m *TcpShKeychainBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShKeychainBag.Merge(m, src)
}
func (m *TcpShKeychainBag) XXX_Size() int {
	return xxx_messageInfo_TcpShKeychainBag.Size(m)
}
func (m *TcpShKeychainBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShKeychainBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShKeychainBag proto.InternalMessageInfo

func (m *TcpShKeychainBag) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *TcpShKeychainBag) GetIsConfigured() bool {
	if m != nil {
		return m.IsConfigured
	}
	return false
}

func (m *TcpShKeychainBag) GetDesiredKeyAvailable() bool {
	if m != nil {
		return m.DesiredKeyAvailable
	}
	return false
}

func (m *TcpShKeychainBag) GetDesiredKeyId() uint64 {
	if m != nil {
		return m.DesiredKeyId
	}
	return 0
}

func (m *TcpShKeychainBag) GetKeys() []*TcpShKeyBag {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *TcpShKeychainBag) GetActiveKey() []*TcpShKeyIdBag {
	if m != nil {
		return m.ActiveKey
	}
	return nil
}

func (m *TcpShKeychainBag) GetSendId() []*TcpShIdBag {
	if m != nil {
		return m.SendId
	}
	return nil
}

func (m *TcpShKeychainBag) GetReceiveId() []*TcpShIdBag {
	if m != nil {
		return m.ReceiveId
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpShKeychainBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.keychains.keychain.tcp_sh_keychain_bag_KEYS")
	proto.RegisterType((*TcpShKeyIdBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.keychains.keychain.tcp_sh_key_id_bag")
	proto.RegisterType((*TcpShKeyBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.keychains.keychain.tcp_sh_key_bag")
	proto.RegisterType((*TcpShIdBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.keychains.keychain.tcp_sh_id_bag")
	proto.RegisterType((*TcpShKeychainBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.keychains.keychain.tcp_sh_keychain_bag")
}

func init() { proto.RegisterFile("tcp_sh_keychain_bag.proto", fileDescriptor_b76af125e7882aca) }

var fileDescriptor_b76af125e7882aca = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcb, 0x8b, 0x1a, 0x4f,
	0x10, 0xc6, 0xc7, 0xaa, 0x53, 0xab, 0x2e, 0xbf, 0x5e, 0xf6, 0xb7, 0x2d, 0x21, 0x20, 0x26, 0x10,
	0xc9, 0x61, 0x0e, 0x6e, 0xde, 0x37, 0x09, 0x1e, 0x44, 0x48, 0xc0, 0x3c, 0x48, 0x4e, 0x4d, 0x3b,
	0xdd, 0x71, 0x0b, 0xc7, 0xe9, 0x61, 0x7a, 0x62, 0x76, 0x8e, 0xf9, 0x03, 0xf2, 0xe7, 0xe6, 0x94,
	0x4b, 0xe8, 0x9a, 0xf1, 0xb1, 0xbb, 0xe6, 0xa6, 0x17, 0x99, 0xea, 0xfa, 0xba, 0xbe, 0xea, 0xef,
	0xab, 0x12, 0x3a, 0x69, 0x10, 0x0b, 0x7b, 0x2d, 0x16, 0x3a, 0x0b, 0xae, 0x25, 0x46, 0x62, 0x26,
	0xe7, 0x7e, 0x9c, 0x98, 0xd4, 0xb0, 0x51, 0x80, 0x36, 0x30, 0x02, 0x8d, 0x15, 0x37, 0x89, 0xc0,
	0x58, 0x38, 0xa8, 0x89, 0x75, 0xe2, 0xbb, 0x8f, 0xc0, 0x44, 0x91, 0x0e, 0x52, 0x34, 0x91, 0x1f,
	0x19, 0xa5, 0x2d, 0xfd, 0xfa, 0xeb, 0x32, 0x76, 0xf3, 0xd5, 0x7b, 0x0f, 0x7c, 0x0f, 0x87, 0x98,
	0x8c, 0xbe, 0x7e, 0x60, 0x6d, 0x28, 0xa3, 0xe2, 0xa5, 0x6e, 0xa9, 0xef, 0x4d, 0xcb, 0xa8, 0xd8,
	0x23, 0x68, 0x6d, 0x40, 0x91, 0x5c, 0x6a, 0x5e, 0xa6, 0x54, 0x73, 0x7d, 0xf8, 0x4e, 0x2e, 0x75,
	0xef, 0x29, 0xfc, 0xb7, 0x2d, 0x28, 0x50, 0xb9, 0x72, 0xec, 0x02, 0x6a, 0x79, 0x44, 0xd5, 0xaa,
	0xd3, 0x93, 0x85, 0xce, 0xc6, 0xaa, 0xf7, 0xbb, 0x02, 0xed, 0x1d, 0xf0, 0xbf, 0x91, 0xec, 0x01,
	0x78, 0x68, 0x85, 0x0c, 0x52, 0x5c, 0xe5, 0xb4, 0x8d, 0x69, 0x03, 0xed, 0x90, 0x62, 0xf6, 0x10,
	0x00, 0xad, 0xd0, 0x37, 0x31, 0x26, 0x5a, 0xf1, 0x0a, 0x65, 0x3d, 0xb4, 0xa3, 0xfc, 0x80, 0x75,
	0xa0, 0x81, 0x56, 0xac, 0x64, 0x88, 0x8a, 0x57, 0x29, 0x59, 0x47, 0xfb, 0xd9, 0x85, 0xec, 0x7f,
	0xa8, 0x25, 0x5a, 0x5a, 0x13, 0xf1, 0x13, 0x7a, 0x4a, 0x11, 0xb1, 0x4b, 0xa8, 0x5b, 0x1d, 0x29,
	0xd7, 0x46, 0xad, 0x5b, 0xea, 0xb7, 0xa6, 0x35, 0x17, 0x8e, 0x95, 0x4b, 0x24, 0x3a, 0x58, 0xb9,
	0x44, 0x3d, 0x4f, 0xb8, 0x70, 0xac, 0x5c, 0x0f, 0x41, 0x92, 0xc5, 0xa9, 0x90, 0xe1, 0xdc, 0xf0,
	0x06, 0x55, 0xf3, 0xe8, 0x64, 0x18, 0xce, 0x8d, 0x93, 0x0e, 0xad, 0x73, 0xe5, 0x1b, 0xce, 0xbf,
	0xbb, 0x2e, 0x3d, 0x6a, 0xa4, 0x89, 0xf6, 0xed, 0xe6, 0x8c, 0xfd, 0x2c, 0xc1, 0x19, 0x46, 0xd4,
	0xa8, 0x4c, 0xb5, 0x72, 0x9a, 0x70, 0xe8, 0x56, 0xfa, 0xa7, 0x83, 0x2f, 0xfe, 0x41, 0xdc, 0xf6,
	0xef, 0x39, 0x33, 0x6d, 0xef, 0x10, 0x4e, 0x74, 0xc6, 0xde, 0x40, 0xc7, 0xac, 0x74, 0x12, 0xca,
	0x38, 0xc6, 0x68, 0x4e, 0x48, 0xb9, 0x92, 0x18, 0xca, 0x59, 0xa8, 0xf9, 0x29, 0x35, 0x7d, 0xb9,
	0x03, 0x98, 0xe8, 0x6c, 0xb8, 0x4e, 0xb3, 0x27, 0x70, 0x76, 0xe7, 0x2e, 0x6f, 0x92, 0x89, 0xed,
	0xdb, 0x37, 0x7a, 0xbf, 0x4a, 0xd0, 0x2a, 0x5a, 0x29, 0x06, 0x64, 0x3b, 0x6a, 0x2d, 0x1a, 0xb5,
	0x10, 0xaa, 0x0b, 0x9d, 0x59, 0x5e, 0x3e, 0xf2, 0xf3, 0x89, 0xa5, 0xf7, 0xa7, 0x0a, 0xe7, 0x7b,
	0xb6, 0x80, 0x4c, 0xdd, 0x4e, 0xfb, 0xa0, 0x30, 0x75, 0x3d, 0xea, 0xf7, 0x4d, 0xbd, 0xda, 0x63,
	0xea, 0x00, 0x2e, 0x94, 0xb6, 0x6e, 0x10, 0xef, 0x88, 0xf9, 0x8c, 0xc0, 0xe7, 0x45, 0xf2, 0x96,
	0x90, 0x8f, 0xa1, 0xbd, 0x7b, 0x07, 0x15, 0x7f, 0x4e, 0x3a, 0x36, 0xb7, 0xe0, 0xb1, 0x62, 0x58,
	0x68, 0xf4, 0x82, 0x34, 0xfa, 0x74, 0x78, 0x8d, 0x36, 0x02, 0xb1, 0x1f, 0x00, 0xf9, 0xee, 0x91,
	0xa9, 0x2f, 0x8f, 0x6c, 0x8a, 0x97, 0x73, 0xb9, 0x71, 0x5c, 0x6e, 0x17, 0xf1, 0x15, 0xb1, 0x7e,
	0x3c, 0x2c, 0x6b, 0xc1, 0xb8, 0x5e, 0x6f, 0x0b, 0x90, 0xe8, 0x40, 0xbb, 0x87, 0xa2, 0xe2, 0xaf,
	0x8f, 0xc8, 0xe8, 0x15, 0x3c, 0x63, 0x35, 0xab, 0xd1, 0x1f, 0xfa, 0xd5, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x63, 0xeb, 0x5f, 0x8d, 0xed, 0x05, 0x00, 0x00,
}
