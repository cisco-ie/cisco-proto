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
// source: asic_summary_bg.proto

package cisco_ios_xr_asic_errors_oper_asic_errors_nodes_node_asic_information_all_instances_all_error_path_summary

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

type AsicSummaryBg_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Asic                 string   `protobuf:"bytes,2,opt,name=asic,proto3" json:"asic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AsicSummaryBg_KEYS) Reset()         { *m = AsicSummaryBg_KEYS{} }
func (m *AsicSummaryBg_KEYS) String() string { return proto.CompactTextString(m) }
func (*AsicSummaryBg_KEYS) ProtoMessage()    {}
func (*AsicSummaryBg_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08b0ec06e7c1495, []int{0}
}

func (m *AsicSummaryBg_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AsicSummaryBg_KEYS.Unmarshal(m, b)
}
func (m *AsicSummaryBg_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AsicSummaryBg_KEYS.Marshal(b, m, deterministic)
}
func (m *AsicSummaryBg_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AsicSummaryBg_KEYS.Merge(m, src)
}
func (m *AsicSummaryBg_KEYS) XXX_Size() int {
	return xxx_messageInfo_AsicSummaryBg_KEYS.Size(m)
}
func (m *AsicSummaryBg_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AsicSummaryBg_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AsicSummaryBg_KEYS proto.InternalMessageInfo

func (m *AsicSummaryBg_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AsicSummaryBg_KEYS) GetAsic() string {
	if m != nil {
		return m.Asic
	}
	return ""
}

type CountInfoBg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountInfoBg) Reset()         { *m = CountInfoBg{} }
func (m *CountInfoBg) String() string { return proto.CompactTextString(m) }
func (*CountInfoBg) ProtoMessage()    {}
func (*CountInfoBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08b0ec06e7c1495, []int{1}
}

func (m *CountInfoBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountInfoBg.Unmarshal(m, b)
}
func (m *CountInfoBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountInfoBg.Marshal(b, m, deterministic)
}
func (m *CountInfoBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountInfoBg.Merge(m, src)
}
func (m *CountInfoBg) XXX_Size() int {
	return xxx_messageInfo_CountInfoBg.Size(m)
}
func (m *CountInfoBg) XXX_DiscardUnknown() {
	xxx_messageInfo_CountInfoBg.DiscardUnknown(m)
}

var xxx_messageInfo_CountInfoBg proto.InternalMessageInfo

func (m *CountInfoBg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CountInfoBg) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type SummaryDataBg struct {
	NumNodes             uint32         `protobuf:"varint,1,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	CrcErrCount          uint32         `protobuf:"varint,2,opt,name=crc_err_count,json=crcErrCount,proto3" json:"crc_err_count,omitempty"`
	SbeErrCount          uint32         `protobuf:"varint,3,opt,name=sbe_err_count,json=sbeErrCount,proto3" json:"sbe_err_count,omitempty"`
	MbeErrCount          uint32         `protobuf:"varint,4,opt,name=mbe_err_count,json=mbeErrCount,proto3" json:"mbe_err_count,omitempty"`
	ParErrCount          uint32         `protobuf:"varint,5,opt,name=par_err_count,json=parErrCount,proto3" json:"par_err_count,omitempty"`
	GenErrCount          uint32         `protobuf:"varint,6,opt,name=gen_err_count,json=genErrCount,proto3" json:"gen_err_count,omitempty"`
	ResetErrCount        uint32         `protobuf:"varint,7,opt,name=reset_err_count,json=resetErrCount,proto3" json:"reset_err_count,omitempty"`
	ErrCount             []*CountInfoBg `protobuf:"bytes,8,rep,name=err_count,json=errCount,proto3" json:"err_count,omitempty"`
	PcieErrCount         []*CountInfoBg `protobuf:"bytes,9,rep,name=pcie_err_count,json=pcieErrCount,proto3" json:"pcie_err_count,omitempty"`
	NodeKey              []uint32       `protobuf:"varint,10,rep,packed,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SummaryDataBg) Reset()         { *m = SummaryDataBg{} }
func (m *SummaryDataBg) String() string { return proto.CompactTextString(m) }
func (*SummaryDataBg) ProtoMessage()    {}
func (*SummaryDataBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08b0ec06e7c1495, []int{2}
}

func (m *SummaryDataBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryDataBg.Unmarshal(m, b)
}
func (m *SummaryDataBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryDataBg.Marshal(b, m, deterministic)
}
func (m *SummaryDataBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryDataBg.Merge(m, src)
}
func (m *SummaryDataBg) XXX_Size() int {
	return xxx_messageInfo_SummaryDataBg.Size(m)
}
func (m *SummaryDataBg) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryDataBg.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryDataBg proto.InternalMessageInfo

func (m *SummaryDataBg) GetNumNodes() uint32 {
	if m != nil {
		return m.NumNodes
	}
	return 0
}

func (m *SummaryDataBg) GetCrcErrCount() uint32 {
	if m != nil {
		return m.CrcErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetSbeErrCount() uint32 {
	if m != nil {
		return m.SbeErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetMbeErrCount() uint32 {
	if m != nil {
		return m.MbeErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetParErrCount() uint32 {
	if m != nil {
		return m.ParErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetGenErrCount() uint32 {
	if m != nil {
		return m.GenErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetResetErrCount() uint32 {
	if m != nil {
		return m.ResetErrCount
	}
	return 0
}

func (m *SummaryDataBg) GetErrCount() []*CountInfoBg {
	if m != nil {
		return m.ErrCount
	}
	return nil
}

func (m *SummaryDataBg) GetPcieErrCount() []*CountInfoBg {
	if m != nil {
		return m.PcieErrCount
	}
	return nil
}

func (m *SummaryDataBg) GetNodeKey() []uint32 {
	if m != nil {
		return m.NodeKey
	}
	return nil
}

type AsicSummaryBg struct {
	LegacyClient         bool             `protobuf:"varint,50,opt,name=legacy_client,json=legacyClient,proto3" json:"legacy_client,omitempty"`
	CihClient            bool             `protobuf:"varint,51,opt,name=cih_client,json=cihClient,proto3" json:"cih_client,omitempty"`
	SumData              []*SummaryDataBg `protobuf:"bytes,52,rep,name=sum_data,json=sumData,proto3" json:"sum_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AsicSummaryBg) Reset()         { *m = AsicSummaryBg{} }
func (m *AsicSummaryBg) String() string { return proto.CompactTextString(m) }
func (*AsicSummaryBg) ProtoMessage()    {}
func (*AsicSummaryBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08b0ec06e7c1495, []int{3}
}

func (m *AsicSummaryBg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AsicSummaryBg.Unmarshal(m, b)
}
func (m *AsicSummaryBg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AsicSummaryBg.Marshal(b, m, deterministic)
}
func (m *AsicSummaryBg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AsicSummaryBg.Merge(m, src)
}
func (m *AsicSummaryBg) XXX_Size() int {
	return xxx_messageInfo_AsicSummaryBg.Size(m)
}
func (m *AsicSummaryBg) XXX_DiscardUnknown() {
	xxx_messageInfo_AsicSummaryBg.DiscardUnknown(m)
}

var xxx_messageInfo_AsicSummaryBg proto.InternalMessageInfo

func (m *AsicSummaryBg) GetLegacyClient() bool {
	if m != nil {
		return m.LegacyClient
	}
	return false
}

func (m *AsicSummaryBg) GetCihClient() bool {
	if m != nil {
		return m.CihClient
	}
	return false
}

func (m *AsicSummaryBg) GetSumData() []*SummaryDataBg {
	if m != nil {
		return m.SumData
	}
	return nil
}

func init() {
	proto.RegisterType((*AsicSummaryBg_KEYS)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.all_instances.all_error_path.summary.asic_summary_bg_KEYS")
	proto.RegisterType((*CountInfoBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.all_instances.all_error_path.summary.count_info_bg")
	proto.RegisterType((*SummaryDataBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.all_instances.all_error_path.summary.summary_data_bg")
	proto.RegisterType((*AsicSummaryBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.all_instances.all_error_path.summary.asic_summary_bg")
}

func init() { proto.RegisterFile("asic_summary_bg.proto", fileDescriptor_d08b0ec06e7c1495) }

var fileDescriptor_d08b0ec06e7c1495 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x15, 0x76, 0xdb, 0x4d, 0xa6, 0x0d, 0x2b, 0x45, 0x45, 0x0a, 0x42, 0x48, 0xab, 0x20,
	0xa1, 0x3d, 0xe5, 0xd0, 0x72, 0xe1, 0x5c, 0x2a, 0x0e, 0x95, 0x7a, 0x08, 0x27, 0x4e, 0x96, 0xe3,
	0x0e, 0x59, 0xc3, 0xda, 0x8e, 0x6c, 0x47, 0x22, 0xe2, 0x3d, 0x7a, 0xe5, 0xd5, 0x78, 0x02, 0x9e,
	0x01, 0x79, 0xd2, 0x5d, 0xbc, 0xfb, 0x00, 0xed, 0x25, 0xf2, 0x7c, 0xf9, 0xcd, 0x9f, 0x7c, 0x99,
	0x04, 0x5e, 0x71, 0x27, 0x05, 0x73, 0x83, 0x52, 0xdc, 0x8e, 0xac, 0xed, 0xea, 0xde, 0x1a, 0x6f,
	0x8a, 0xef, 0x42, 0x3a, 0x61, 0x98, 0x34, 0x8e, 0xfd, 0xb4, 0x8c, 0x18, 0xb4, 0xd6, 0x58, 0xc7,
	0x4c, 0x8f, 0xb6, 0x8e, 0x84, 0x5a, 0x9b, 0x7b, 0x9c, 0xae, 0x93, 0x2c, 0xf5, 0x37, 0x63, 0x15,
	0xf7, 0xd2, 0xe8, 0x9a, 0x6f, 0xb7, 0x4c, 0x6a, 0xe7, 0xb9, 0x16, 0xe8, 0x28, 0xa2, 0x24, 0xd6,
	0x73, 0xbf, 0xa9, 0x1f, 0x9b, 0x56, 0x9f, 0xe1, 0xe2, 0x68, 0x08, 0x76, 0x7b, 0xf3, 0xf5, 0x4b,
	0xf1, 0x06, 0xb2, 0x50, 0x95, 0x69, 0xae, 0xb0, 0x4c, 0x56, 0xc9, 0x3a, 0x6b, 0xd2, 0x20, 0xdc,
	0x71, 0x85, 0x45, 0x01, 0xf3, 0x90, 0x54, 0xbe, 0x20, 0x9d, 0xce, 0xd5, 0x47, 0xc8, 0x85, 0x19,
	0xb4, 0xa7, 0x11, 0x58, 0xdb, 0x05, 0x28, 0x4a, 0xa6, 0x73, 0x71, 0x01, 0x27, 0x04, 0x51, 0x66,
	0xde, 0x4c, 0x41, 0xf5, 0x67, 0x0e, 0xcb, 0x5d, 0xff, 0x7b, 0xee, 0x79, 0xc8, 0x0e, 0xfd, 0x07,
	0xc5, 0xe8, 0xf9, 0xa8, 0x44, 0xde, 0xa4, 0x7a, 0x50, 0x77, 0x21, 0x2e, 0x2a, 0xc8, 0x85, 0x25,
	0x0f, 0x58, 0x5c, 0xee, 0x4c, 0x58, 0x71, 0x63, 0xed, 0x75, 0x90, 0x02, 0xe3, 0x5a, 0x8c, 0x98,
	0xd9, 0xc4, 0xb8, 0x16, 0x63, 0x46, 0x1d, 0x30, 0xf3, 0x89, 0x51, 0x87, 0x4c, 0xcf, 0x6d, 0xc4,
	0x9c, 0x4c, 0x4c, 0xcf, 0x6d, 0xcc, 0x74, 0xa8, 0x23, 0xe6, 0x74, 0x62, 0x3a, 0xd4, 0x7b, 0xe6,
	0x3d, 0x2c, 0x2d, 0x3a, 0xf4, 0x11, 0xb5, 0x20, 0x2a, 0x27, 0x79, 0xcf, 0x3d, 0x24, 0x90, 0xfd,
	0x47, 0xd2, 0xd5, 0x6c, 0x7d, 0x76, 0x39, 0xd6, 0x4f, 0xb7, 0x11, 0xf5, 0xc1, 0x5b, 0x6c, 0x52,
	0xdc, 0x0d, 0xf6, 0x3b, 0x81, 0x97, 0xbd, 0x90, 0xb1, 0x5d, 0xd9, 0x73, 0x4f, 0x77, 0x1e, 0x06,
	0xda, 0x5b, 0xf7, 0x1a, 0x68, 0x45, 0xd9, 0x0f, 0x1c, 0x4b, 0x58, 0xcd, 0xd6, 0x79, 0xb3, 0x08,
	0xf1, 0x2d, 0x8e, 0xd5, 0xdf, 0x04, 0x96, 0x47, 0x7b, 0x5e, 0xbc, 0x83, 0x7c, 0x8b, 0x1d, 0x17,
	0x23, 0x13, 0x5b, 0x89, 0xda, 0x97, 0x97, 0xab, 0x64, 0x9d, 0x36, 0xe7, 0x93, 0x78, 0x4d, 0x5a,
	0xf1, 0x16, 0x40, 0xc8, 0xcd, 0x8e, 0xb8, 0x22, 0x22, 0x13, 0x72, 0xf3, 0x78, 0xfb, 0x21, 0x81,
	0xd4, 0x0d, 0x8a, 0xd6, 0xb6, 0xfc, 0x40, 0x76, 0xfc, 0x7a, 0x4a, 0x3b, 0x8e, 0x3e, 0x9b, 0x66,
	0xe1, 0x06, 0xf5, 0x89, 0x7b, 0xde, 0x9e, 0xd2, 0xaf, 0xe4, 0xea, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x03, 0x33, 0x76, 0x63, 0x04, 0x00, 0x00,
}
