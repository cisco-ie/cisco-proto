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
// source: pce_asso_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_association_infos_association_info

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

type PceAssoBag_KEYS struct {
	GroupId              uint32   `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	SubId                uint32   `protobuf:"varint,3,opt,name=sub_id,json=subId,proto3" json:"sub_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceAssoBag_KEYS) Reset()         { *m = PceAssoBag_KEYS{} }
func (m *PceAssoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceAssoBag_KEYS) ProtoMessage()    {}
func (*PceAssoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6934e182f49bd6f3, []int{0}
}

func (m *PceAssoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceAssoBag_KEYS.Unmarshal(m, b)
}
func (m *PceAssoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceAssoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceAssoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceAssoBag_KEYS.Merge(m, src)
}
func (m *PceAssoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceAssoBag_KEYS.Size(m)
}
func (m *PceAssoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceAssoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceAssoBag_KEYS proto.InternalMessageInfo

func (m *PceAssoBag_KEYS) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *PceAssoBag_KEYS) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PceAssoBag_KEYS) GetSubId() uint32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type PceAssoLspInfo struct {
	PccAddress           string   `protobuf:"bytes,1,opt,name=pcc_address,json=pccAddress,proto3" json:"pcc_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Lspid                uint32   `protobuf:"varint,3,opt,name=lspid,proto3" json:"lspid,omitempty"`
	TunnelName           string   `protobuf:"bytes,4,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	PceBased             bool     `protobuf:"varint,5,opt,name=pce_based,json=pceBased,proto3" json:"pce_based,omitempty"`
	PlspId               uint32   `protobuf:"varint,6,opt,name=plsp_id,json=plspId,proto3" json:"plsp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceAssoLspInfo) Reset()         { *m = PceAssoLspInfo{} }
func (m *PceAssoLspInfo) String() string { return proto.CompactTextString(m) }
func (*PceAssoLspInfo) ProtoMessage()    {}
func (*PceAssoLspInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6934e182f49bd6f3, []int{1}
}

func (m *PceAssoLspInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceAssoLspInfo.Unmarshal(m, b)
}
func (m *PceAssoLspInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceAssoLspInfo.Marshal(b, m, deterministic)
}
func (m *PceAssoLspInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceAssoLspInfo.Merge(m, src)
}
func (m *PceAssoLspInfo) XXX_Size() int {
	return xxx_messageInfo_PceAssoLspInfo.Size(m)
}
func (m *PceAssoLspInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PceAssoLspInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PceAssoLspInfo proto.InternalMessageInfo

func (m *PceAssoLspInfo) GetPccAddress() string {
	if m != nil {
		return m.PccAddress
	}
	return ""
}

func (m *PceAssoLspInfo) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *PceAssoLspInfo) GetLspid() uint32 {
	if m != nil {
		return m.Lspid
	}
	return 0
}

func (m *PceAssoLspInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *PceAssoLspInfo) GetPceBased() bool {
	if m != nil {
		return m.PceBased
	}
	return false
}

func (m *PceAssoLspInfo) GetPlspId() uint32 {
	if m != nil {
		return m.PlspId
	}
	return 0
}

type PceAssoBag struct {
	AssociationType      uint32            `protobuf:"varint,50,opt,name=association_type,json=associationType,proto3" json:"association_type,omitempty"`
	AssociationId        uint32            `protobuf:"varint,51,opt,name=association_id,json=associationId,proto3" json:"association_id,omitempty"`
	AssociationSource    string            `protobuf:"bytes,52,opt,name=association_source,json=associationSource,proto3" json:"association_source,omitempty"`
	AssociationLsp       []*PceAssoLspInfo `protobuf:"bytes,53,rep,name=association_lsp,json=associationLsp,proto3" json:"association_lsp,omitempty"`
	Strict               bool              `protobuf:"varint,54,opt,name=strict,proto3" json:"strict,omitempty"`
	Status               uint32            `protobuf:"varint,55,opt,name=status,proto3" json:"status,omitempty"`
	HeadendsSwapped      uint32            `protobuf:"varint,56,opt,name=headends_swapped,json=headendsSwapped,proto3" json:"headends_swapped,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PceAssoBag) Reset()         { *m = PceAssoBag{} }
func (m *PceAssoBag) String() string { return proto.CompactTextString(m) }
func (*PceAssoBag) ProtoMessage()    {}
func (*PceAssoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6934e182f49bd6f3, []int{2}
}

func (m *PceAssoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceAssoBag.Unmarshal(m, b)
}
func (m *PceAssoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceAssoBag.Marshal(b, m, deterministic)
}
func (m *PceAssoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceAssoBag.Merge(m, src)
}
func (m *PceAssoBag) XXX_Size() int {
	return xxx_messageInfo_PceAssoBag.Size(m)
}
func (m *PceAssoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceAssoBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceAssoBag proto.InternalMessageInfo

func (m *PceAssoBag) GetAssociationType() uint32 {
	if m != nil {
		return m.AssociationType
	}
	return 0
}

func (m *PceAssoBag) GetAssociationId() uint32 {
	if m != nil {
		return m.AssociationId
	}
	return 0
}

func (m *PceAssoBag) GetAssociationSource() string {
	if m != nil {
		return m.AssociationSource
	}
	return ""
}

func (m *PceAssoBag) GetAssociationLsp() []*PceAssoLspInfo {
	if m != nil {
		return m.AssociationLsp
	}
	return nil
}

func (m *PceAssoBag) GetStrict() bool {
	if m != nil {
		return m.Strict
	}
	return false
}

func (m *PceAssoBag) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *PceAssoBag) GetHeadendsSwapped() uint32 {
	if m != nil {
		return m.HeadendsSwapped
	}
	return 0
}

func init() {
	proto.RegisterType((*PceAssoBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.association_infos.association_info.pce_asso_bag_KEYS")
	proto.RegisterType((*PceAssoLspInfo)(nil), "cisco_ios_xr_infra_xtc_oper.pce.association_infos.association_info.pce_asso_lsp_info")
	proto.RegisterType((*PceAssoBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.association_infos.association_info.pce_asso_bag")
}

func init() { proto.RegisterFile("pce_asso_bag.proto", fileDescriptor_6934e182f49bd6f3) }

var fileDescriptor_6934e182f49bd6f3 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0x95, 0x69, 0x27, 0xcd, 0xdc, 0x52, 0xa0, 0x16, 0x3f, 0x46, 0x2c, 0x88, 0x46, 0x42,
	0x0a, 0x0b, 0xb2, 0x68, 0xf9, 0xdb, 0x52, 0x89, 0x45, 0x04, 0x62, 0x91, 0x81, 0x45, 0x57, 0x96,
	0x63, 0x9b, 0x12, 0x69, 0x1a, 0x5b, 0xbe, 0x8e, 0x28, 0x1b, 0x1e, 0x8d, 0x07, 0xe2, 0x29, 0x90,
	0xaf, 0x07, 0x61, 0x34, 0x4b, 0x76, 0x39, 0xc7, 0xc7, 0xe7, 0xda, 0x5f, 0x0c, 0xcc, 0x29, 0x23,
	0x24, 0xa2, 0x15, 0x83, 0xbc, 0x6a, 0x9d, 0xb7, 0xc1, 0xb2, 0x0b, 0x35, 0xa2, 0xb2, 0x62, 0xb4,
	0x28, 0x6e, 0xbc, 0x18, 0xa7, 0x2f, 0x5e, 0x8a, 0x9b, 0xa0, 0x84, 0x75, 0xc6, 0xb7, 0x4e, 0x99,
	0x36, 0xe6, 0xd5, 0x28, 0xc3, 0x68, 0xa7, 0xb8, 0x6c, 0x71, 0xcf, 0x59, 0x5f, 0xc2, 0x69, 0xde,
	0x2c, 0xde, 0xbf, 0xbb, 0xdc, 0xb0, 0x47, 0x50, 0x5d, 0x79, 0x3b, 0x3b, 0x31, 0x6a, 0x5e, 0xd4,
	0x45, 0x73, 0xd2, 0x1f, 0x91, 0xee, 0x34, 0x63, 0x70, 0x18, 0xbe, 0x3b, 0xc3, 0x17, 0x75, 0xd1,
	0xac, 0x7a, 0xfa, 0x66, 0xf7, 0xa1, 0xc4, 0x79, 0x88, 0xe1, 0x03, 0x0a, 0x2f, 0x71, 0x1e, 0x3a,
	0xbd, 0xfe, 0x59, 0x64, 0xdd, 0x5b, 0x74, 0x34, 0x90, 0x3d, 0x81, 0x63, 0xa7, 0x94, 0x90, 0x5a,
	0x7b, 0x83, 0x48, 0xf5, 0xab, 0x1e, 0x9c, 0x52, 0x6f, 0x93, 0xc3, 0x1e, 0xc3, 0x2a, 0xcc, 0xd3,
	0x64, 0xb6, 0xb1, 0x70, 0x41, 0x85, 0x55, 0x32, 0x3a, 0xcd, 0xee, 0xc1, 0x72, 0x8b, 0xee, 0xef,
	0x24, 0x12, 0xb1, 0x73, 0xb7, 0x65, 0x92, 0xd7, 0x86, 0x1f, 0xa6, 0xce, 0x64, 0x7d, 0x94, 0xd7,
	0x26, 0x76, 0xc6, 0x93, 0x0c, 0x12, 0x8d, 0xe6, 0xcb, 0xba, 0x68, 0xaa, 0xbe, 0x72, 0xca, 0x5c,
	0x44, 0xcd, 0x1e, 0xc2, 0x91, 0xa3, 0xe3, 0x69, 0x5e, 0x52, 0x6b, 0x19, 0x65, 0xa7, 0xd7, 0xbf,
	0x16, 0x70, 0x2b, 0x87, 0xc3, 0x9e, 0xc1, 0xdd, 0x1c, 0x20, 0x81, 0x38, 0xa3, 0x2d, 0x77, 0x32,
	0xff, 0x53, 0x64, 0xf2, 0x14, 0x6e, 0xff, 0xc3, 0x5a, 0xf3, 0x73, 0x0a, 0x9e, 0x64, 0x6e, 0xa7,
	0xd9, 0x73, 0x60, 0x79, 0x0c, 0xed, 0xec, 0x95, 0xe1, 0x2f, 0xe8, 0x02, 0xa7, 0xd9, 0xca, 0x86,
	0x16, 0xd8, 0x0f, 0xc8, 0x07, 0x45, 0xa8, 0xfc, 0x65, 0x7d, 0xd0, 0x1c, 0x9f, 0x7d, 0x6e, 0xff,
	0xff, 0x2d, 0xb4, 0x7b, 0x3f, 0xab, 0xcf, 0xef, 0xf0, 0x01, 0x1d, 0x7b, 0x00, 0x25, 0x06, 0x3f,
	0xaa, 0xc0, 0x5f, 0x11, 0xc4, 0x9d, 0x4a, 0xbe, 0x0c, 0x33, 0xf2, 0xd7, 0x89, 0x60, 0x52, 0x11,
	0xd8, 0x57, 0x23, 0xb5, 0x99, 0x34, 0x0a, 0xfc, 0x26, 0x9d, 0x33, 0x9a, 0xbf, 0x49, 0xc0, 0xfe,
	0xf8, 0x9b, 0x64, 0x0f, 0x25, 0xbd, 0xe9, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x75,
	0xc5, 0x9e, 0xe9, 0x02, 0x00, 0x00,
}
