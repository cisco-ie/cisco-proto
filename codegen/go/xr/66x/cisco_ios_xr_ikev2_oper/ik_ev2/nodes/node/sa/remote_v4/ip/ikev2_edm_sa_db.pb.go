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

package cisco_ios_xr_ikev2_oper_ik_ev2_nodes_node_sa_remote_v4_ip

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
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
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

func (m *Ikev2EdmSaDb_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
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
	proto.RegisterType((*Ikev2EdmSaDb_KEYS)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.remote_v4.ip.ikev2_edm_sa_db_KEYS")
	proto.RegisterType((*Ikev2EdmSa)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.remote_v4.ip.ikev2_edm_sa")
	proto.RegisterType((*Ikev2EdmSaDb)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.sa.remote_v4.ip.ikev2_edm_sa_db")
}

func init() { proto.RegisterFile("ikev2_edm_sa_db.proto", fileDescriptor_afb18716fb1bab5f) }

var fileDescriptor_afb18716fb1bab5f = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x95, 0x3f, 0x6d, 0x93, 0x49, 0xd2, 0x04, 0x53, 0x24, 0x57, 0x08, 0x11, 0x72, 0x40,
	0x39, 0xed, 0x21, 0x70, 0xe1, 0x58, 0x09, 0x54, 0x45, 0x08, 0x84, 0x12, 0x09, 0xc4, 0xc9, 0x72,
	0xd7, 0x93, 0xac, 0x49, 0xb2, 0x5e, 0x79, 0x9c, 0x15, 0xe5, 0x93, 0xf1, 0xf1, 0xd0, 0xd8, 0x1b,
	0xa8, 0x7a, 0xe4, 0x12, 0x79, 0x7e, 0xef, 0xcd, 0xdb, 0xf1, 0x64, 0x17, 0x9e, 0xd9, 0x1d, 0xd6,
	0x0b, 0x85, 0xe6, 0xa0, 0x48, 0x2b, 0x73, 0x97, 0x55, 0xde, 0x05, 0x27, 0xde, 0xe5, 0x96, 0x72,
	0xa7, 0xac, 0x23, 0xf5, 0xd3, 0xab, 0xe4, 0x71, 0x15, 0xfa, 0xcc, 0xee, 0x14, 0xd6, 0x8b, 0xac,
	0x74, 0x06, 0x29, 0xfe, 0x66, 0xa4, 0x33, 0x8f, 0x07, 0x17, 0x50, 0xd5, 0x6f, 0x33, 0x5b, 0xcd,
	0x3e, 0xc1, 0xd5, 0xa3, 0x4c, 0xf5, 0xf1, 0xc3, 0xf7, 0xb5, 0x78, 0x0e, 0x7d, 0xf6, 0xab, 0x52,
	0x1f, 0x50, 0xb6, 0xa6, 0xad, 0x79, 0x7f, 0xd5, 0x63, 0xf0, 0x59, 0x1f, 0x50, 0x48, 0xb8, 0xd0,
	0xc6, 0x78, 0x24, 0x92, 0xed, 0x28, 0x9d, 0xca, 0xd9, 0xef, 0x2e, 0x0c, 0x1f, 0xe6, 0x71, 0x4e,
	0x38, 0x96, 0x25, 0xee, 0x95, 0x35, 0x31, 0x67, 0xb4, 0xea, 0x25, 0xb0, 0x34, 0xe2, 0x35, 0x8c,
	0xf7, 0x2e, 0xd7, 0x7b, 0xc5, 0xed, 0xaa, 0x72, 0x3e, 0x34, 0x79, 0xa3, 0x88, 0x6f, 0x8c, 0xf1,
	0x5f, 0x9c, 0x0f, 0x62, 0x0e, 0x93, 0x66, 0xe8, 0x7f, 0xc6, 0x4e, 0x34, 0x5e, 0x26, 0xfe, 0xd7,
	0x79, 0x05, 0x67, 0x14, 0x74, 0x40, 0xd9, 0x8d, 0x72, 0x2a, 0x84, 0x80, 0x2e, 0x96, 0xb9, 0x97,
	0x67, 0x11, 0xc6, 0x33, 0xdf, 0x61, 0x87, 0xf7, 0x64, 0x7f, 0xa1, 0x3c, 0x8f, 0x63, 0x9d, 0x4a,
	0x31, 0x81, 0x4e, 0xe5, 0x37, 0xf2, 0x22, 0x9a, 0xf9, 0xc8, 0xfd, 0x85, 0xa6, 0x42, 0xf6, 0x52,
	0x3f, 0x9f, 0xc5, 0x35, 0xf4, 0x4c, 0xa1, 0xb6, 0xde, 0x1d, 0x2b, 0xd9, 0x4f, 0x01, 0xa6, 0xb8,
	0xe5, 0x92, 0xef, 0xac, 0x8f, 0xa1, 0x50, 0x64, 0xb7, 0xa5, 0x84, 0xb4, 0x3b, 0x06, 0x6b, 0xbb,
	0x2d, 0xc5, 0x4b, 0x18, 0x44, 0xb1, 0x46, 0x6f, 0x37, 0xf7, 0x72, 0x10, 0x65, 0x60, 0xf4, 0x35,
	0x12, 0x36, 0xec, 0xed, 0x06, 0x95, 0xce, 0x83, 0xad, 0x51, 0x0e, 0x93, 0x81, 0xd1, 0x4d, 0x24,
	0xe2, 0x05, 0x00, 0x21, 0x91, 0x75, 0x25, 0xef, 0x74, 0x14, 0x9f, 0xdd, 0x6f, 0xc8, 0xd2, 0x70,
	0x3f, 0xdf, 0xfa, 0x48, 0xca, 0x20, 0xe5, 0xf2, 0x32, 0xf5, 0x27, 0xf4, 0x1e, 0x29, 0xe7, 0xf1,
	0xd2, 0xd6, 0xa9, 0xb2, 0x72, 0x9c, 0xc6, 0x8b, 0x60, 0x5d, 0x59, 0x0e, 0x6f, 0x56, 0xcd, 0xea,
	0x24, 0xaa, 0xfd, 0x44, 0x58, 0xbe, 0x86, 0x64, 0xe5, 0x27, 0x3f, 0x49, 0x7f, 0x7d, 0xac, 0x97,
	0x86, 0x63, 0x9b, 0x4e, 0x6b, 0xa4, 0x48, 0xb1, 0x09, 0x2c, 0x8d, 0x78, 0x05, 0x43, 0xd2, 0xca,
	0x96, 0x36, 0x58, 0x1d, 0x9c, 0x97, 0x4f, 0xa7, 0xad, 0x79, 0x6f, 0x35, 0x20, 0xbd, 0x3c, 0xa1,
	0xd9, 0x0f, 0x18, 0x3f, 0x7a, 0x13, 0xc5, 0x37, 0x68, 0x93, 0x96, 0x8b, 0x69, 0x67, 0x3e, 0x58,
	0xdc, 0x66, 0xff, 0xfd, 0x92, 0x67, 0x0f, 0x73, 0x57, 0x6d, 0xd2, 0x77, 0xe7, 0xf1, 0xbb, 0x79,
	0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2f, 0xcd, 0xe8, 0x4a, 0x50, 0x03, 0x00, 0x00,
}
