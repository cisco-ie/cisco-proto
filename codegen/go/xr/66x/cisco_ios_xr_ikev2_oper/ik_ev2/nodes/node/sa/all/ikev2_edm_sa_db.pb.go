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
// source: ikev2_edm_sa_db.proto

package cisco_ios_xr_ikev2_oper_ik_ev2_nodes_node_sa_all

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

type Ikev2EdmSaDb_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ikev2EdmSaDb_KEYS) Reset()         { *m = Ikev2EdmSaDb_KEYS{} }
func (m *Ikev2EdmSaDb_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ikev2EdmSaDb_KEYS) ProtoMessage()    {}
func (*Ikev2EdmSaDb_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_afb18716fb1bab5f, []int{0}
}

func (m *Ikev2EdmSaDb_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ikev2EdmSaDb_KEYS.Unmarshal(m, b)
}
func (m *Ikev2EdmSaDb_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ikev2EdmSaDb_KEYS.Marshal(b, m, deterministic)
}
func (m *Ikev2EdmSaDb_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ikev2EdmSaDb_KEYS.Merge(m, src)
}
func (m *Ikev2EdmSaDb_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ikev2EdmSaDb_KEYS.Size(m)
}
func (m *Ikev2EdmSaDb_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ikev2EdmSaDb_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ikev2EdmSaDb_KEYS proto.InternalMessageInfo

func (m *Ikev2EdmSaDb_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ikev2EdmSa struct {
	TunnelId             uint32   `protobuf:"varint,1,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	LocalAddrPort        string   `protobuf:"bytes,2,opt,name=local_addr_port,json=localAddrPort,proto3" json:"local_addr_port,omitempty"`
	RemoteAddrPort       string   `protobuf:"bytes,3,opt,name=remote_addr_port,json=remoteAddrPort,proto3" json:"remote_addr_port,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Encr                 string   `protobuf:"bytes,5,opt,name=encr,proto3" json:"encr,omitempty"`
	Keysize              uint32   `protobuf:"varint,6,opt,name=keysize,proto3" json:"keysize,omitempty"`
	Prf                  string   `protobuf:"bytes,7,opt,name=prf,proto3" json:"prf,omitempty"`
	Hash                 string   `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
	DhGroup              uint32   `protobuf:"varint,9,opt,name=dh_group,json=dhGroup,proto3" json:"dh_group,omitempty"`
	AuthSign             string   `protobuf:"bytes,10,opt,name=auth_sign,json=authSign,proto3" json:"auth_sign,omitempty"`
	AuthVerify           string   `protobuf:"bytes,11,opt,name=auth_verify,json=authVerify,proto3" json:"auth_verify,omitempty"`
	LifeActive           string   `protobuf:"bytes,12,opt,name=life_active,json=lifeActive,proto3" json:"life_active,omitempty"`
	SessionId            uint32   `protobuf:"varint,13,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	StatusDesc           string   `protobuf:"bytes,14,opt,name=status_desc,json=statusDesc,proto3" json:"status_desc,omitempty"`
	LocalSpi             string   `protobuf:"bytes,15,opt,name=local_spi,json=localSpi,proto3" json:"local_spi,omitempty"`
	RemoteSpi            string   `protobuf:"bytes,16,opt,name=remote_spi,json=remoteSpi,proto3" json:"remote_spi,omitempty"`
	LocalId              string   `protobuf:"bytes,17,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	RemoteId             string   `protobuf:"bytes,18,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	SaInitiator          bool     `protobuf:"varint,19,opt,name=sa_initiator,json=saInitiator,proto3" json:"sa_initiator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ikev2EdmSa) Reset()         { *m = Ikev2EdmSa{} }
func (m *Ikev2EdmSa) String() string { return proto.CompactTextString(m) }
func (*Ikev2EdmSa) ProtoMessage()    {}
func (*Ikev2EdmSa) Descriptor() ([]byte, []int) {
	return fileDescriptor_afb18716fb1bab5f, []int{1}
}

func (m *Ikev2EdmSa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ikev2EdmSa.Unmarshal(m, b)
}
func (m *Ikev2EdmSa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ikev2EdmSa.Marshal(b, m, deterministic)
}
func (m *Ikev2EdmSa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ikev2EdmSa.Merge(m, src)
}
func (m *Ikev2EdmSa) XXX_Size() int {
	return xxx_messageInfo_Ikev2EdmSa.Size(m)
}
func (m *Ikev2EdmSa) XXX_DiscardUnknown() {
	xxx_messageInfo_Ikev2EdmSa.DiscardUnknown(m)
}

var xxx_messageInfo_Ikev2EdmSa proto.InternalMessageInfo

func (m *Ikev2EdmSa) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *Ikev2EdmSa) GetLocalAddrPort() string {
	if m != nil {
		return m.LocalAddrPort
	}
	return ""
}

func (m *Ikev2EdmSa) GetRemoteAddrPort() string {
	if m != nil {
		return m.RemoteAddrPort
	}
	return ""
}

func (m *Ikev2EdmSa) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Ikev2EdmSa) GetEncr() string {
	if m != nil {
		return m.Encr
	}
	return ""
}

func (m *Ikev2EdmSa) GetKeysize() uint32 {
	if m != nil {
		return m.Keysize
	}
	return 0
}

func (m *Ikev2EdmSa) GetPrf() string {
	if m != nil {
		return m.Prf
	}
	return ""
}

func (m *Ikev2EdmSa) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Ikev2EdmSa) GetDhGroup() uint32 {
	if m != nil {
		return m.DhGroup
	}
	return 0
}

func (m *Ikev2EdmSa) GetAuthSign() string {
	if m != nil {
		return m.AuthSign
	}
	return ""
}

func (m *Ikev2EdmSa) GetAuthVerify() string {
	if m != nil {
		return m.AuthVerify
	}
	return ""
}

func (m *Ikev2EdmSa) GetLifeActive() string {
	if m != nil {
		return m.LifeActive
	}
	return ""
}

func (m *Ikev2EdmSa) GetSessionId() uint32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *Ikev2EdmSa) GetStatusDesc() string {
	if m != nil {
		return m.StatusDesc
	}
	return ""
}

func (m *Ikev2EdmSa) GetLocalSpi() string {
	if m != nil {
		return m.LocalSpi
	}
	return ""
}

func (m *Ikev2EdmSa) GetRemoteSpi() string {
	if m != nil {
		return m.RemoteSpi
	}
	return ""
}

func (m *Ikev2EdmSa) GetLocalId() string {
	if m != nil {
		return m.LocalId
	}
	return ""
}

func (m *Ikev2EdmSa) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *Ikev2EdmSa) GetSaInitiator() bool {
	if m != nil {
		return m.SaInitiator
	}
	return false
}

type Ikev2EdmSaDb struct {
	Sa                   []*Ikev2EdmSa `protobuf:"bytes,50,rep,name=sa,proto3" json:"sa,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Ikev2EdmSaDb) Reset()         { *m = Ikev2EdmSaDb{} }
func (m *Ikev2EdmSaDb) String() string { return proto.CompactTextString(m) }
func (*Ikev2EdmSaDb) ProtoMessage()    {}
func (*Ikev2EdmSaDb) Descriptor() ([]byte, []int) {
	return fileDescriptor_afb18716fb1bab5f, []int{2}
}

func (m *Ikev2EdmSaDb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ikev2EdmSaDb.Unmarshal(m, b)
}
func (m *Ikev2EdmSaDb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ikev2EdmSaDb.Marshal(b, m, deterministic)
}
func (m *Ikev2EdmSaDb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ikev2EdmSaDb.Merge(m, src)
}
func (m *Ikev2EdmSaDb) XXX_Size() int {
	return xxx_messageInfo_Ikev2EdmSaDb.Size(m)
}
func (m *Ikev2EdmSaDb) XXX_DiscardUnknown() {
	xxx_messageInfo_Ikev2EdmSaDb.DiscardUnknown(m)
}

var xxx_messageInfo_Ikev2EdmSaDb proto.InternalMessageInfo

func (m *Ikev2EdmSaDb) GetSa() []*Ikev2EdmSa {
	if m != nil {
		return m.Sa
	}
	return nil
}

func init() {
	proto.RegisterType((*Ikev2EdmSaDb_KEYS)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.all.ikev2_edm_sa_db_KEYS")
	proto.RegisterType((*Ikev2EdmSa)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.all.ikev2_edm_sa")
	proto.RegisterType((*Ikev2EdmSaDb)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.all.ikev2_edm_sa_db")
}

func init() { proto.RegisterFile("ikev2_edm_sa_db.proto", fileDescriptor_afb18716fb1bab5f) }

var fileDescriptor_afb18716fb1bab5f = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6b, 0x1b, 0x31,
	0x10, 0xc5, 0xf1, 0x9f, 0xc4, 0xb6, 0x6c, 0xc7, 0xae, 0x9a, 0x82, 0x42, 0x29, 0x75, 0x7d, 0x28,
	0x3e, 0x2d, 0xc5, 0xb9, 0x17, 0x02, 0x2d, 0x65, 0x29, 0x84, 0x62, 0x43, 0xa1, 0x27, 0xa1, 0xac,
	0xc6, 0x5e, 0xe1, 0xf5, 0x6a, 0xd1, 0xc8, 0xa6, 0xe9, 0x27, 0xeb, 0xc7, 0x2b, 0x33, 0x72, 0x8a,
	0xc9, 0x2d, 0x97, 0x45, 0xf3, 0x7b, 0x6f, 0x9e, 0xa4, 0x41, 0x2b, 0xde, 0xb8, 0x1d, 0x1c, 0x97,
	0x1a, 0xec, 0x5e, 0xa3, 0xd1, 0xf6, 0x21, 0x6b, 0x82, 0x8f, 0x5e, 0x7e, 0x2a, 0x1c, 0x16, 0x5e,
	0x3b, 0x8f, 0xfa, 0x77, 0xd0, 0xc9, 0xe3, 0x1b, 0x08, 0x99, 0xdb, 0x69, 0x38, 0x2e, 0xb3, 0xda,
	0x5b, 0x40, 0xfe, 0x66, 0x68, 0x32, 0x53, 0x55, 0xf3, 0x5b, 0x71, 0xfd, 0x2c, 0x4a, 0x7f, 0xff,
	0xfa, 0x6b, 0x2d, 0xdf, 0x8a, 0x01, 0xd9, 0x74, 0x6d, 0xf6, 0xa0, 0x5a, 0xb3, 0xd6, 0x62, 0xb0,
	0xea, 0x13, 0xb8, 0x37, 0x7b, 0x98, 0xff, 0xed, 0x8a, 0xd1, 0x79, 0x17, 0xb9, 0xe3, 0xa1, 0xae,
	0xa1, 0xd2, 0xce, 0xb2, 0x7b, 0xbc, 0xea, 0x27, 0x90, 0x5b, 0xf9, 0x51, 0x4c, 0x2a, 0x5f, 0x98,
	0x4a, 0x1b, 0x6b, 0x83, 0x6e, 0x7c, 0x88, 0xaa, 0xcd, 0x81, 0x63, 0xc6, 0x77, 0xd6, 0x86, 0x1f,
	0x3e, 0x44, 0xb9, 0x10, 0xd3, 0x00, 0x7b, 0x1f, 0xe1, 0xcc, 0xd8, 0x61, 0xe3, 0x55, 0xe2, 0xff,
	0x9d, 0xd7, 0xe2, 0x02, 0xa3, 0x89, 0xa0, 0xba, 0x2c, 0xa7, 0x42, 0x4a, 0xd1, 0x85, 0xba, 0x08,
	0xea, 0x82, 0x21, 0xaf, 0xa5, 0x12, 0xbd, 0x1d, 0x3c, 0xa2, 0xfb, 0x03, 0xea, 0x92, 0x8f, 0xf5,
	0x54, 0xca, 0xa9, 0xe8, 0x34, 0x61, 0xa3, 0x7a, 0x6c, 0xa6, 0x25, 0xf5, 0x97, 0x06, 0x4b, 0xd5,
	0x4f, 0xfd, 0xb4, 0x96, 0x37, 0xa2, 0x6f, 0x4b, 0xbd, 0x0d, 0xfe, 0xd0, 0xa8, 0x41, 0x0a, 0xb0,
	0xe5, 0x37, 0x2a, 0xe9, 0xce, 0xe6, 0x10, 0x4b, 0x8d, 0x6e, 0x5b, 0x2b, 0x91, 0x26, 0x44, 0x60,
	0xed, 0xb6, 0xb5, 0x7c, 0x2f, 0x86, 0x2c, 0x1e, 0x21, 0xb8, 0xcd, 0xa3, 0x1a, 0xb2, 0x2c, 0x08,
	0xfd, 0x64, 0x42, 0x86, 0xca, 0x6d, 0x40, 0x9b, 0x22, 0xba, 0x23, 0xa8, 0x51, 0x32, 0x10, 0xba,
	0x63, 0x22, 0xdf, 0x09, 0x81, 0x80, 0xe8, 0x7c, 0x4d, 0x33, 0x1d, 0xf3, 0xde, 0x83, 0x13, 0xc9,
	0x2d, 0xf5, 0xd3, 0xad, 0x0f, 0xa8, 0x2d, 0x60, 0xa1, 0xae, 0x52, 0x7f, 0x42, 0x5f, 0x00, 0x0b,
	0x3a, 0x5e, 0x9a, 0x3a, 0x36, 0x4e, 0x4d, 0xd2, 0xf1, 0x18, 0xac, 0x1b, 0x47, 0xe1, 0xa7, 0x51,
	0x93, 0x3a, 0x65, 0x75, 0x90, 0x08, 0xc9, 0x37, 0x22, 0x59, 0x69, 0xe7, 0x57, 0x2c, 0xf6, 0xb8,
	0xce, 0x2d, 0xc5, 0x9e, 0x3a, 0x9d, 0x55, 0x32, 0xc5, 0x26, 0x90, 0x5b, 0xf9, 0x41, 0x8c, 0xd0,
	0x68, 0x57, 0xbb, 0xe8, 0x4c, 0xf4, 0x41, 0xbd, 0x9e, 0xb5, 0x16, 0xfd, 0xd5, 0x10, 0x4d, 0xfe,
	0x84, 0xe6, 0x46, 0x4c, 0x9e, 0xbd, 0x37, 0x79, 0x2f, 0xda, 0x68, 0xd4, 0x72, 0xd6, 0x59, 0x0c,
	0x97, 0x9f, 0xb3, 0x97, 0xbe, 0xe0, 0xec, 0x3c, 0x6e, 0xd5, 0x46, 0xf3, 0x70, 0xc9, 0xff, 0xc2,
	0xed, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x1e, 0x2f, 0xd0, 0x24, 0x03, 0x00, 0x00,
}
