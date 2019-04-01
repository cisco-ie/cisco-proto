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
// source: fia_card_tcam.proto

package cisco_ios_xr_fia_internal_tcam_oper_controller_dpa_nodes_node_external_tcam_resources

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

type FiaCardTcam_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FiaCardTcam_KEYS) Reset()         { *m = FiaCardTcam_KEYS{} }
func (m *FiaCardTcam_KEYS) String() string { return proto.CompactTextString(m) }
func (*FiaCardTcam_KEYS) ProtoMessage()    {}
func (*FiaCardTcam_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9677eeb50a9072, []int{0}
}

func (m *FiaCardTcam_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaCardTcam_KEYS.Unmarshal(m, b)
}
func (m *FiaCardTcam_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaCardTcam_KEYS.Marshal(b, m, deterministic)
}
func (m *FiaCardTcam_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaCardTcam_KEYS.Merge(m, src)
}
func (m *FiaCardTcam_KEYS) XXX_Size() int {
	return xxx_messageInfo_FiaCardTcam_KEYS.Size(m)
}
func (m *FiaCardTcam_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaCardTcam_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FiaCardTcam_KEYS proto.InternalMessageInfo

func (m *FiaCardTcam_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type FiaTcamDb struct {
	DbId                 uint32   `protobuf:"varint,1,opt,name=db_id,json=dbId,proto3" json:"db_id,omitempty"`
	DbInuseEntries       uint32   `protobuf:"varint,2,opt,name=db_inuse_entries,json=dbInuseEntries,proto3" json:"db_inuse_entries,omitempty"`
	DbPrefix             string   `protobuf:"bytes,3,opt,name=db_prefix,json=dbPrefix,proto3" json:"db_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FiaTcamDb) Reset()         { *m = FiaTcamDb{} }
func (m *FiaTcamDb) String() string { return proto.CompactTextString(m) }
func (*FiaTcamDb) ProtoMessage()    {}
func (*FiaTcamDb) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9677eeb50a9072, []int{1}
}

func (m *FiaTcamDb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaTcamDb.Unmarshal(m, b)
}
func (m *FiaTcamDb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaTcamDb.Marshal(b, m, deterministic)
}
func (m *FiaTcamDb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaTcamDb.Merge(m, src)
}
func (m *FiaTcamDb) XXX_Size() int {
	return xxx_messageInfo_FiaTcamDb.Size(m)
}
func (m *FiaTcamDb) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaTcamDb.DiscardUnknown(m)
}

var xxx_messageInfo_FiaTcamDb proto.InternalMessageInfo

func (m *FiaTcamDb) GetDbId() uint32 {
	if m != nil {
		return m.DbId
	}
	return 0
}

func (m *FiaTcamDb) GetDbInuseEntries() uint32 {
	if m != nil {
		return m.DbInuseEntries
	}
	return 0
}

func (m *FiaTcamDb) GetDbPrefix() string {
	if m != nil {
		return m.DbPrefix
	}
	return ""
}

type FiaTcamBank struct {
	BankId               string       `protobuf:"bytes,1,opt,name=bank_id,json=bankId,proto3" json:"bank_id,omitempty"`
	BankKeySize          string       `protobuf:"bytes,2,opt,name=bank_key_size,json=bankKeySize,proto3" json:"bank_key_size,omitempty"`
	BankFreeEntries      uint32       `protobuf:"varint,3,opt,name=bank_free_entries,json=bankFreeEntries,proto3" json:"bank_free_entries,omitempty"`
	BankInuseEntries     uint32       `protobuf:"varint,4,opt,name=bank_inuse_entries,json=bankInuseEntries,proto3" json:"bank_inuse_entries,omitempty"`
	Owner                string       `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	NofDbs               uint32       `protobuf:"varint,6,opt,name=nof_dbs,json=nofDbs,proto3" json:"nof_dbs,omitempty"`
	BankDb               []*FiaTcamDb `protobuf:"bytes,7,rep,name=bank_db,json=bankDb,proto3" json:"bank_db,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FiaTcamBank) Reset()         { *m = FiaTcamBank{} }
func (m *FiaTcamBank) String() string { return proto.CompactTextString(m) }
func (*FiaTcamBank) ProtoMessage()    {}
func (*FiaTcamBank) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9677eeb50a9072, []int{2}
}

func (m *FiaTcamBank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaTcamBank.Unmarshal(m, b)
}
func (m *FiaTcamBank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaTcamBank.Marshal(b, m, deterministic)
}
func (m *FiaTcamBank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaTcamBank.Merge(m, src)
}
func (m *FiaTcamBank) XXX_Size() int {
	return xxx_messageInfo_FiaTcamBank.Size(m)
}
func (m *FiaTcamBank) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaTcamBank.DiscardUnknown(m)
}

var xxx_messageInfo_FiaTcamBank proto.InternalMessageInfo

func (m *FiaTcamBank) GetBankId() string {
	if m != nil {
		return m.BankId
	}
	return ""
}

func (m *FiaTcamBank) GetBankKeySize() string {
	if m != nil {
		return m.BankKeySize
	}
	return ""
}

func (m *FiaTcamBank) GetBankFreeEntries() uint32 {
	if m != nil {
		return m.BankFreeEntries
	}
	return 0
}

func (m *FiaTcamBank) GetBankInuseEntries() uint32 {
	if m != nil {
		return m.BankInuseEntries
	}
	return 0
}

func (m *FiaTcamBank) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FiaTcamBank) GetNofDbs() uint32 {
	if m != nil {
		return m.NofDbs
	}
	return 0
}

func (m *FiaTcamBank) GetBankDb() []*FiaTcamDb {
	if m != nil {
		return m.BankDb
	}
	return nil
}

type FiaNpuTcam struct {
	NpuId                uint32         `protobuf:"varint,1,opt,name=npu_id,json=npuId,proto3" json:"npu_id,omitempty"`
	TcamBank             []*FiaTcamBank `protobuf:"bytes,2,rep,name=tcam_bank,json=tcamBank,proto3" json:"tcam_bank,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FiaNpuTcam) Reset()         { *m = FiaNpuTcam{} }
func (m *FiaNpuTcam) String() string { return proto.CompactTextString(m) }
func (*FiaNpuTcam) ProtoMessage()    {}
func (*FiaNpuTcam) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9677eeb50a9072, []int{3}
}

func (m *FiaNpuTcam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaNpuTcam.Unmarshal(m, b)
}
func (m *FiaNpuTcam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaNpuTcam.Marshal(b, m, deterministic)
}
func (m *FiaNpuTcam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaNpuTcam.Merge(m, src)
}
func (m *FiaNpuTcam) XXX_Size() int {
	return xxx_messageInfo_FiaNpuTcam.Size(m)
}
func (m *FiaNpuTcam) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaNpuTcam.DiscardUnknown(m)
}

var xxx_messageInfo_FiaNpuTcam proto.InternalMessageInfo

func (m *FiaNpuTcam) GetNpuId() uint32 {
	if m != nil {
		return m.NpuId
	}
	return 0
}

func (m *FiaNpuTcam) GetTcamBank() []*FiaTcamBank {
	if m != nil {
		return m.TcamBank
	}
	return nil
}

type FiaCardTcam struct {
	NpuTcam              []*FiaNpuTcam `protobuf:"bytes,50,rep,name=npu_tcam,json=npuTcam,proto3" json:"npu_tcam,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FiaCardTcam) Reset()         { *m = FiaCardTcam{} }
func (m *FiaCardTcam) String() string { return proto.CompactTextString(m) }
func (*FiaCardTcam) ProtoMessage()    {}
func (*FiaCardTcam) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e9677eeb50a9072, []int{4}
}

func (m *FiaCardTcam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiaCardTcam.Unmarshal(m, b)
}
func (m *FiaCardTcam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiaCardTcam.Marshal(b, m, deterministic)
}
func (m *FiaCardTcam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiaCardTcam.Merge(m, src)
}
func (m *FiaCardTcam) XXX_Size() int {
	return xxx_messageInfo_FiaCardTcam.Size(m)
}
func (m *FiaCardTcam) XXX_DiscardUnknown() {
	xxx_messageInfo_FiaCardTcam.DiscardUnknown(m)
}

var xxx_messageInfo_FiaCardTcam proto.InternalMessageInfo

func (m *FiaCardTcam) GetNpuTcam() []*FiaNpuTcam {
	if m != nil {
		return m.NpuTcam
	}
	return nil
}

func init() {
	proto.RegisterType((*FiaCardTcam_KEYS)(nil), "cisco_ios_xr_fia_internal_tcam_oper.controller.dpa.nodes.node.external_tcam_resources.fia_card_tcam_KEYS")
	proto.RegisterType((*FiaTcamDb)(nil), "cisco_ios_xr_fia_internal_tcam_oper.controller.dpa.nodes.node.external_tcam_resources.fia_tcam_db")
	proto.RegisterType((*FiaTcamBank)(nil), "cisco_ios_xr_fia_internal_tcam_oper.controller.dpa.nodes.node.external_tcam_resources.fia_tcam_bank")
	proto.RegisterType((*FiaNpuTcam)(nil), "cisco_ios_xr_fia_internal_tcam_oper.controller.dpa.nodes.node.external_tcam_resources.fia_npu_tcam")
	proto.RegisterType((*FiaCardTcam)(nil), "cisco_ios_xr_fia_internal_tcam_oper.controller.dpa.nodes.node.external_tcam_resources.fia_card_tcam")
}

func init() { proto.RegisterFile("fia_card_tcam.proto", fileDescriptor_7e9677eeb50a9072) }

var fileDescriptor_7e9677eeb50a9072 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x49, 0xbb, 0x4d, 0x9b, 0xa9, 0xd5, 0x75, 0x56, 0x31, 0xe0, 0xcd, 0x92, 0xab, 0x22,
	0x12, 0x70, 0x7d, 0x03, 0xd9, 0x15, 0xca, 0x82, 0x48, 0x56, 0x2f, 0xbc, 0x1a, 0xe6, 0xcf, 0x09,
	0x0c, 0x6d, 0x67, 0xc2, 0x4c, 0x82, 0xdd, 0x15, 0x04, 0x9f, 0xc0, 0xe7, 0xf0, 0x75, 0x7c, 0x22,
	0x39, 0x27, 0x36, 0xdb, 0x3e, 0x40, 0x6f, 0xda, 0x99, 0xef, 0x7c, 0xc3, 0xef, 0x3b, 0x73, 0x26,
	0xec, 0xa2, 0xb6, 0x52, 0x68, 0x19, 0x8c, 0x68, 0xb5, 0xdc, 0x96, 0x4d, 0xf0, 0xad, 0xe7, 0x5f,
	0xb5, 0x8d, 0xda, 0x0b, 0xeb, 0xa3, 0xd8, 0x05, 0x81, 0x0e, 0xeb, 0x5a, 0x08, 0x4e, 0x6e, 0xc8,
	0x25, 0x7c, 0x03, 0xa1, 0xd4, 0xde, 0xb5, 0xc1, 0x6f, 0x36, 0x10, 0x4a, 0xd3, 0xc8, 0xd2, 0x79,
	0x03, 0x91, 0x7e, 0x4b, 0xd8, 0x1d, 0x9a, 0x03, 0x44, 0xdf, 0x05, 0x0d, 0xb1, 0x78, 0xc7, 0xf8,
	0x11, 0x4d, 0xdc, 0xde, 0x7c, 0xbb, 0xe3, 0xaf, 0x59, 0x86, 0xa7, 0x84, 0x93, 0x5b, 0xc8, 0x93,
	0xcb, 0x64, 0x99, 0x55, 0x33, 0x14, 0x3e, 0xc9, 0x2d, 0x14, 0x96, 0xcd, 0xf1, 0x08, 0xb9, 0x8d,
	0xe2, 0x17, 0x6c, 0x62, 0x94, 0xb0, 0x86, 0x7c, 0x8b, 0xea, 0xcc, 0xa8, 0x95, 0xe1, 0x4b, 0x76,
	0x8e, 0xa2, 0xeb, 0x22, 0x08, 0x70, 0x6d, 0xb0, 0x10, 0xf3, 0x11, 0xd5, 0x9f, 0x1a, 0xb5, 0x42,
	0xf9, 0xa6, 0x57, 0x11, 0x65, 0x94, 0x68, 0x02, 0xd4, 0x76, 0x97, 0x8f, 0x7b, 0x94, 0x51, 0x9f,
	0x69, 0x5f, 0xfc, 0x1d, 0xb1, 0xc5, 0xc0, 0x52, 0xd2, 0xad, 0xf9, 0x2b, 0x36, 0xc5, 0xff, 0x3d,
	0x2f, 0xab, 0x52, 0xdc, 0xae, 0x0c, 0x2f, 0xd8, 0x82, 0x0a, 0x6b, 0xb8, 0x17, 0xd1, 0x3e, 0x00,
	0xe1, 0xb2, 0x6a, 0x8e, 0xe2, 0x2d, 0xdc, 0xdf, 0xd9, 0x07, 0xe0, 0x6f, 0xd8, 0x73, 0xf2, 0xd4,
	0x01, 0x1e, 0x63, 0x8d, 0x29, 0xd6, 0x33, 0x2c, 0x7c, 0x0c, 0x30, 0xe4, 0x7a, 0xcb, 0x78, 0x0f,
	0x3a, 0xea, 0xe1, 0x8c, 0xcc, 0xe7, 0xc4, 0x3c, 0xec, 0xe2, 0x05, 0x9b, 0xf8, 0xef, 0x0e, 0x42,
	0x3e, 0x21, 0x6a, 0xbf, 0xc1, 0xb0, 0xce, 0xd7, 0xc2, 0xa8, 0x98, 0xa7, 0x74, 0x30, 0x75, 0xbe,
	0xbe, 0x56, 0x91, 0xff, 0xf8, 0xdf, 0x85, 0x51, 0xf9, 0xf4, 0x72, 0xbc, 0x9c, 0x5f, 0xa9, 0xf2,
	0x24, 0xe3, 0x2d, 0x0f, 0x06, 0xd5, 0xdf, 0xd4, 0xb5, 0x2a, 0xfe, 0x24, 0xec, 0x09, 0xea, 0xae,
	0xe9, 0xa8, 0xc6, 0x5f, 0xb2, 0x14, 0xd7, 0xc3, 0x08, 0x27, 0xae, 0xe9, 0x56, 0x86, 0xff, 0x4a,
	0x58, 0x36, 0x5c, 0x7c, 0x3e, 0xa2, 0x9c, 0xe6, 0xd4, 0x39, 0x91, 0x55, 0xcd, 0x70, 0xf9, 0x41,
	0xba, 0x75, 0xf1, 0x3b, 0xe9, 0x1f, 0xc0, 0xf0, 0x3e, 0xf9, 0x4f, 0x36, 0xdb, 0x07, 0xcf, 0xaf,
	0x28, 0x93, 0x3e, 0x61, 0xa6, 0x3d, 0xaa, 0x9a, 0xba, 0xa6, 0xfb, 0xa2, 0xe5, 0x56, 0xa5, 0xf4,
	0x39, 0xbe, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x31, 0xc0, 0xe4, 0xa5, 0x03, 0x00, 0x00,
}
