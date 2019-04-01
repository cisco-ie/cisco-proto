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
// source: mpls_lm_bfd_nbrs_link_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_bfd_links_link

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

type MplsLmBfdNbrsLinkInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBfdNbrsLinkInfo_KEYS) Reset()         { *m = MplsLmBfdNbrsLinkInfo_KEYS{} }
func (m *MplsLmBfdNbrsLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrsLinkInfo_KEYS) ProtoMessage()    {}
func (*MplsLmBfdNbrsLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ceecba296887812, []int{0}
}

func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Merge(m, src)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.Size(m)
}
func (m *MplsLmBfdNbrsLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrsLinkInfo_KEYS proto.InternalMessageInfo

func (m *MplsLmBfdNbrsLinkInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type MplsLmBfdNbrInfo struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	IsBfdUp              bool     `protobuf:"varint,2,opt,name=is_bfd_up,json=isBfdUp,proto3" json:"is_bfd_up,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmBfdNbrInfo) Reset()         { *m = MplsLmBfdNbrInfo{} }
func (m *MplsLmBfdNbrInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrInfo) ProtoMessage()    {}
func (*MplsLmBfdNbrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ceecba296887812, []int{1}
}

func (m *MplsLmBfdNbrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmBfdNbrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrInfo.Merge(m, src)
}
func (m *MplsLmBfdNbrInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrInfo.Size(m)
}
func (m *MplsLmBfdNbrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrInfo proto.InternalMessageInfo

func (m *MplsLmBfdNbrInfo) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

func (m *MplsLmBfdNbrInfo) GetIsBfdUp() bool {
	if m != nil {
		return m.IsBfdUp
	}
	return false
}

type MplsLmBfdNbrsLinkInfo struct {
	Neighbor             []*MplsLmBfdNbrInfo `protobuf:"bytes,50,rep,name=neighbor,proto3" json:"neighbor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MplsLmBfdNbrsLinkInfo) Reset()         { *m = MplsLmBfdNbrsLinkInfo{} }
func (m *MplsLmBfdNbrsLinkInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmBfdNbrsLinkInfo) ProtoMessage()    {}
func (*MplsLmBfdNbrsLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ceecba296887812, []int{2}
}

func (m *MplsLmBfdNbrsLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Unmarshal(m, b)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Merge(m, src)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmBfdNbrsLinkInfo.Size(m)
}
func (m *MplsLmBfdNbrsLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmBfdNbrsLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmBfdNbrsLinkInfo proto.InternalMessageInfo

func (m *MplsLmBfdNbrsLinkInfo) GetNeighbor() []*MplsLmBfdNbrInfo {
	if m != nil {
		return m.Neighbor
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsLmBfdNbrsLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.links.link.mpls_lm_bfd_nbrs_link_info_KEYS")
	proto.RegisterType((*MplsLmBfdNbrInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.links.link.mpls_lm_bfd_nbr_info")
	proto.RegisterType((*MplsLmBfdNbrsLinkInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.links.link.mpls_lm_bfd_nbrs_link_info")
}

func init() { proto.RegisterFile("mpls_lm_bfd_nbrs_link_info.proto", fileDescriptor_1ceecba296887812) }

var fileDescriptor_1ceecba296887812 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0x82, 0xee, 0x46, 0xfc, 0x20, 0x78, 0x28, 0x7b, 0xb1, 0x14, 0x84, 0x7a, 0x09,
	0xb2, 0xfe, 0x02, 0x05, 0x45, 0x10, 0x3c, 0x54, 0x3c, 0x78, 0x90, 0xa1, 0x69, 0x26, 0x3a, 0xd8,
	0x7c, 0x90, 0xa9, 0xe0, 0xd1, 0x9f, 0x2e, 0xc6, 0xdd, 0x3d, 0x88, 0x3d, 0x78, 0x09, 0xe4, 0x61,
	0xde, 0x87, 0x97, 0x57, 0x54, 0x2e, 0x0e, 0x0c, 0x83, 0x03, 0x6d, 0x0d, 0x78, 0x9d, 0x18, 0x06,
	0xf2, 0x6f, 0x40, 0xde, 0x06, 0x15, 0x53, 0x18, 0x83, 0x3c, 0xef, 0x89, 0xfb, 0x00, 0x14, 0x18,
	0x3e, 0x12, 0xe4, 0xf3, 0x11, 0x21, 0x44, 0x4c, 0x6a, 0xf5, 0x51, 0xda, 0x1a, 0xf5, 0x1d, 0xe3,
	0xfc, 0xd6, 0xb7, 0xe2, 0x64, 0xda, 0x0a, 0x77, 0xd7, 0x4f, 0x0f, 0xf2, 0x54, 0x1c, 0x90, 0x1f,
	0x31, 0xd9, 0xae, 0x47, 0xf0, 0x9d, 0xc3, 0xb2, 0xa8, 0x8a, 0x66, 0xde, 0xee, 0x6f, 0xe8, 0x7d,
	0xe7, 0xb0, 0x7e, 0x16, 0xc7, 0xbf, 0x4c, 0xd9, 0x21, 0xcf, 0xc4, 0x91, 0x47, 0x7a, 0x79, 0xd5,
	0x21, 0x41, 0x67, 0x4c, 0x42, 0xe6, 0x95, 0xe0, 0x70, 0xcd, 0x2f, 0x7f, 0xb0, 0x5c, 0x88, 0x39,
	0x71, 0x4e, 0xbf, 0xc7, 0x72, 0xab, 0x2a, 0x9a, 0x59, 0xbb, 0x4b, 0x7c, 0x65, 0xcd, 0x63, 0xac,
	0x3f, 0x0b, 0xb1, 0x98, 0x6e, 0x2a, 0xb5, 0x98, 0xad, 0x6d, 0xe5, 0xb2, 0xda, 0x6e, 0xf6, 0x96,
	0x37, 0xea, 0xbf, 0x63, 0xa8, 0xbf, 0xfa, 0xb7, 0x1b, 0xaf, 0xde, 0xc9, 0x23, 0x5f, 0x7c, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x99, 0xdb, 0x68, 0x7a, 0x88, 0x01, 0x00, 0x00,
}
