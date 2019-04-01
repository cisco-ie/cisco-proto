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
// source: csi_ma_admin_info.proto

package cisco_ios_xr_cofo_csi_ma_oper_csima_nodes_node_csi_indexes_csi_index

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

type CsiMaAdminInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	CsiIndex             uint32   `protobuf:"varint,2,opt,name=csi_index,json=csiIndex,proto3" json:"csi_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsiMaAdminInfo_KEYS) Reset()         { *m = CsiMaAdminInfo_KEYS{} }
func (m *CsiMaAdminInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*CsiMaAdminInfo_KEYS) ProtoMessage()    {}
func (*CsiMaAdminInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_254bb90b13b2ea34, []int{0}
}

func (m *CsiMaAdminInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsiMaAdminInfo_KEYS.Unmarshal(m, b)
}
func (m *CsiMaAdminInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsiMaAdminInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *CsiMaAdminInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsiMaAdminInfo_KEYS.Merge(m, src)
}
func (m *CsiMaAdminInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_CsiMaAdminInfo_KEYS.Size(m)
}
func (m *CsiMaAdminInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CsiMaAdminInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CsiMaAdminInfo_KEYS proto.InternalMessageInfo

func (m *CsiMaAdminInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *CsiMaAdminInfo_KEYS) GetCsiIndex() uint32 {
	if m != nil {
		return m.CsiIndex
	}
	return 0
}

type CsiMaAdminInfo struct {
	CsiIndexXr           uint32   `protobuf:"varint,50,opt,name=csi_index_xr,json=csiIndexXr,proto3" json:"csi_index_xr,omitempty"`
	SdrId                uint32   `protobuf:"varint,51,opt,name=sdr_id,json=sdrId,proto3" json:"sdr_id,omitempty"`
	PeerSdrId            uint32   `protobuf:"varint,52,opt,name=peer_sdr_id,json=peerSdrId,proto3" json:"peer_sdr_id,omitempty"`
	PeerSdrName          string   `protobuf:"bytes,53,opt,name=peer_sdr_name,json=peerSdrName,proto3" json:"peer_sdr_name,omitempty"`
	State                string   `protobuf:"bytes,54,opt,name=state,proto3" json:"state,omitempty"`
	CsiHelperReg         bool     `protobuf:"varint,55,opt,name=csi_helper_reg,json=csiHelperReg,proto3" json:"csi_helper_reg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsiMaAdminInfo) Reset()         { *m = CsiMaAdminInfo{} }
func (m *CsiMaAdminInfo) String() string { return proto.CompactTextString(m) }
func (*CsiMaAdminInfo) ProtoMessage()    {}
func (*CsiMaAdminInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_254bb90b13b2ea34, []int{1}
}

func (m *CsiMaAdminInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsiMaAdminInfo.Unmarshal(m, b)
}
func (m *CsiMaAdminInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsiMaAdminInfo.Marshal(b, m, deterministic)
}
func (m *CsiMaAdminInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsiMaAdminInfo.Merge(m, src)
}
func (m *CsiMaAdminInfo) XXX_Size() int {
	return xxx_messageInfo_CsiMaAdminInfo.Size(m)
}
func (m *CsiMaAdminInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CsiMaAdminInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CsiMaAdminInfo proto.InternalMessageInfo

func (m *CsiMaAdminInfo) GetCsiIndexXr() uint32 {
	if m != nil {
		return m.CsiIndexXr
	}
	return 0
}

func (m *CsiMaAdminInfo) GetSdrId() uint32 {
	if m != nil {
		return m.SdrId
	}
	return 0
}

func (m *CsiMaAdminInfo) GetPeerSdrId() uint32 {
	if m != nil {
		return m.PeerSdrId
	}
	return 0
}

func (m *CsiMaAdminInfo) GetPeerSdrName() string {
	if m != nil {
		return m.PeerSdrName
	}
	return ""
}

func (m *CsiMaAdminInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CsiMaAdminInfo) GetCsiHelperReg() bool {
	if m != nil {
		return m.CsiHelperReg
	}
	return false
}

func init() {
	proto.RegisterType((*CsiMaAdminInfo_KEYS)(nil), "cisco_ios_xr_cofo_csi_ma_oper.csima.nodes.node.csi_indexes.csi_index.csi_ma_admin_info_KEYS")
	proto.RegisterType((*CsiMaAdminInfo)(nil), "cisco_ios_xr_cofo_csi_ma_oper.csima.nodes.node.csi_indexes.csi_index.csi_ma_admin_info")
}

func init() { proto.RegisterFile("csi_ma_admin_info.proto", fileDescriptor_254bb90b13b2ea34) }

var fileDescriptor_254bb90b13b2ea34 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xfc, 0x30,
	0x14, 0xc5, 0xe9, 0x1f, 0x66, 0x68, 0xef, 0xfc, 0x47, 0x30, 0xf8, 0x11, 0x10, 0xa4, 0x0c, 0x2e,
	0xba, 0xea, 0xc2, 0xf1, 0xe3, 0x05, 0x14, 0x1c, 0x04, 0x17, 0x9d, 0x8d, 0xae, 0x2e, 0x31, 0xb9,
	0x33, 0x06, 0x6c, 0x52, 0x92, 0x2e, 0xfa, 0x90, 0x3e, 0x94, 0xe4, 0x5a, 0xeb, 0x62, 0x36, 0x21,
	0x39, 0xbf, 0x7b, 0x4e, 0x0e, 0x17, 0xce, 0x75, 0xb4, 0xd8, 0x2a, 0x54, 0xa6, 0xb5, 0x0e, 0xad,
	0xdb, 0xf9, 0xba, 0x0b, 0xbe, 0xf7, 0xe2, 0x41, 0xdb, 0xa8, 0x3d, 0x5a, 0x1f, 0x71, 0x08, 0xa8,
	0xfd, 0xce, 0xe3, 0x38, 0xea, 0x3b, 0x0a, 0xb5, 0x8e, 0xb6, 0x55, 0xb5, 0xf3, 0x86, 0x22, 0x9f,
	0x49, 0x40, 0xeb, 0x0c, 0x0d, 0x14, 0xff, 0xee, 0xab, 0x06, 0xce, 0x0e, 0x3e, 0xc0, 0xe7, 0xc7,
	0xb7, 0xad, 0xb8, 0x80, 0x22, 0xf9, 0xd0, 0xa9, 0x96, 0x64, 0x56, 0x66, 0x55, 0xd1, 0xe4, 0x49,
	0x78, 0x51, 0x2d, 0x25, 0x38, 0x65, 0xc8, 0x7f, 0x65, 0x56, 0x2d, 0x9b, 0x5c, 0x47, 0xbb, 0xe1,
	0xcc, 0xaf, 0x0c, 0x8e, 0x0f, 0x42, 0x45, 0x09, 0xff, 0x27, 0x0b, 0x0e, 0x41, 0x5e, 0xb3, 0x0b,
	0x7e, 0x5d, 0xaf, 0x41, 0x9c, 0xc2, 0x3c, 0x9a, 0x80, 0xd6, 0xc8, 0x35, 0xb3, 0x59, 0x34, 0x61,
	0x63, 0xc4, 0x25, 0x2c, 0x3a, 0xa2, 0x80, 0x23, 0xbb, 0x61, 0x56, 0x24, 0x69, 0xcb, 0x7c, 0x05,
	0xcb, 0x89, 0x73, 0xd9, 0x5b, 0x2e, 0xbb, 0x18, 0x27, 0xb8, 0xef, 0x09, 0xcc, 0x62, 0xaf, 0x7a,
	0x92, 0x77, 0xcc, 0x7e, 0x1e, 0xe2, 0x0a, 0x8e, 0x52, 0xa5, 0x0f, 0xfa, 0xec, 0x28, 0x60, 0xa0,
	0xbd, 0xbc, 0x2f, 0xb3, 0x2a, 0x6f, 0x52, 0xd1, 0x27, 0x16, 0x1b, 0xda, 0xbf, 0xcf, 0x79, 0xdf,
	0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x8c, 0xe6, 0x54, 0x8a, 0x01, 0x00, 0x00,
}