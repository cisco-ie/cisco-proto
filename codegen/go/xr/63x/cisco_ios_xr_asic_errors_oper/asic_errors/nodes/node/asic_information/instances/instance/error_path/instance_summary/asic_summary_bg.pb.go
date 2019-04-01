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

package cisco_ios_xr_asic_errors_oper_asic_errors_nodes_node_asic_information_instances_instance_error_path_instance_summary

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
	AsicInstance         uint32   `protobuf:"varint,3,opt,name=asic_instance,json=asicInstance,proto3" json:"asic_instance,omitempty"`
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

func (m *AsicSummaryBg_KEYS) GetAsicInstance() uint32 {
	if m != nil {
		return m.AsicInstance
	}
	return 0
}

type SummaryDataBg struct {
	NumNodes             uint32   `protobuf:"varint,1,opt,name=num_nodes,json=numNodes,proto3" json:"num_nodes,omitempty"`
	CrcErrCount          uint32   `protobuf:"varint,2,opt,name=crc_err_count,json=crcErrCount,proto3" json:"crc_err_count,omitempty"`
	SbeErrCount          uint32   `protobuf:"varint,3,opt,name=sbe_err_count,json=sbeErrCount,proto3" json:"sbe_err_count,omitempty"`
	MbeErrCount          uint32   `protobuf:"varint,4,opt,name=mbe_err_count,json=mbeErrCount,proto3" json:"mbe_err_count,omitempty"`
	ParErrCount          uint32   `protobuf:"varint,5,opt,name=par_err_count,json=parErrCount,proto3" json:"par_err_count,omitempty"`
	GenErrCount          uint32   `protobuf:"varint,6,opt,name=gen_err_count,json=genErrCount,proto3" json:"gen_err_count,omitempty"`
	ResetErrCount        uint32   `protobuf:"varint,7,opt,name=reset_err_count,json=resetErrCount,proto3" json:"reset_err_count,omitempty"`
	ErrCount             []uint32 `protobuf:"varint,8,rep,packed,name=err_count,json=errCount,proto3" json:"err_count,omitempty"`
	PcieErrCount         []uint32 `protobuf:"varint,9,rep,packed,name=pcie_err_count,json=pcieErrCount,proto3" json:"pcie_err_count,omitempty"`
	NodeKey              []uint32 `protobuf:"varint,10,rep,packed,name=node_key,json=nodeKey,proto3" json:"node_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryDataBg) Reset()         { *m = SummaryDataBg{} }
func (m *SummaryDataBg) String() string { return proto.CompactTextString(m) }
func (*SummaryDataBg) ProtoMessage()    {}
func (*SummaryDataBg) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08b0ec06e7c1495, []int{1}
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

func (m *SummaryDataBg) GetErrCount() []uint32 {
	if m != nil {
		return m.ErrCount
	}
	return nil
}

func (m *SummaryDataBg) GetPcieErrCount() []uint32 {
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
	return fileDescriptor_d08b0ec06e7c1495, []int{2}
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
	proto.RegisterType((*AsicSummaryBg_KEYS)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.instance_summary.asic_summary_bg_KEYS")
	proto.RegisterType((*SummaryDataBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.instance_summary.summary_data_bg")
	proto.RegisterType((*AsicSummaryBg)(nil), "cisco_ios_xr_asic_errors_oper.asic_errors.nodes.node.asic_information.instances.instance.error_path.instance_summary.asic_summary_bg")
}

func init() { proto.RegisterFile("asic_summary_bg.proto", fileDescriptor_d08b0ec06e7c1495) }

var fileDescriptor_d08b0ec06e7c1495 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x41, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x95, 0x4e, 0xe9, 0x24, 0x6f, 0x26, 0x8c, 0x14, 0x81, 0x14, 0x54, 0x21, 0x45, 0x03,
	0x42, 0xb3, 0xca, 0xa2, 0xe5, 0x06, 0xa5, 0x0b, 0x54, 0xa9, 0x8b, 0xb0, 0x62, 0x65, 0x39, 0xee,
	0x23, 0x63, 0x31, 0xb6, 0x23, 0xdb, 0x91, 0x98, 0x03, 0x70, 0x10, 0x6e, 0xc7, 0x15, 0xd8, 0x21,
	0xbf, 0x4c, 0x82, 0x3b, 0x17, 0x60, 0x13, 0xd9, 0x9f, 0x3f, 0xfb, 0xf7, 0xd3, 0x73, 0xe0, 0x35,
	0x77, 0x52, 0x30, 0x37, 0x28, 0xc5, 0xed, 0x91, 0xb5, 0x5d, 0xdd, 0x5b, 0xe3, 0x4d, 0xe1, 0x85,
	0x74, 0xc2, 0x30, 0x69, 0x1c, 0xfb, 0x61, 0x19, 0x39, 0x68, 0xad, 0xb1, 0x8e, 0x99, 0x1e, 0x6d,
	0x1d, 0x81, 0x5a, 0x9b, 0x27, 0x1c, 0xbf, 0x23, 0x96, 0xfa, 0x9b, 0xb1, 0x8a, 0x7b, 0x69, 0x74,
	0x2d, 0xb5, 0xf3, 0x5c, 0x0b, 0x74, 0xf3, 0xa8, 0xa6, 0x5d, 0xac, 0xe7, 0x7e, 0x3f, 0xb3, 0x29,
	0x7e, 0x7b, 0x80, 0x57, 0x67, 0xd7, 0x61, 0x0f, 0xf7, 0x5f, 0xbf, 0x14, 0xd7, 0x90, 0x85, 0xf3,
	0x99, 0xe6, 0x0a, 0xcb, 0xa4, 0x4a, 0x76, 0x59, 0x93, 0x06, 0xf0, 0xc8, 0x15, 0x16, 0x05, 0x5c,
	0x86, 0x4d, 0xe5, 0x05, 0x71, 0x1a, 0x17, 0xef, 0x20, 0x3f, 0xdd, 0x65, 0x4c, 0x28, 0x17, 0x55,
	0xb2, 0xcb, 0x9b, 0x75, 0x80, 0x9f, 0x4f, 0x6c, 0xfb, 0xfb, 0x02, 0x36, 0x53, 0xd2, 0x13, 0xf7,
	0x9c, 0xb5, 0x1d, 0x25, 0x0d, 0x8a, 0x51, 0x4d, 0x94, 0x94, 0x37, 0xa9, 0x1e, 0xd4, 0x63, 0x98,
	0x17, 0x5b, 0xc8, 0x85, 0xa5, 0xba, 0x99, 0x30, 0x83, 0xf6, 0x14, 0x99, 0x37, 0x2b, 0x61, 0xc5,
	0xbd, 0xb5, 0x77, 0x01, 0x05, 0xc7, 0xb5, 0x18, 0x39, 0x63, 0xf2, 0xca, 0xb5, 0x18, 0x3b, 0xea,
	0x99, 0x73, 0x39, 0x3a, 0xea, 0xb9, 0xd3, 0x73, 0x1b, 0x39, 0x2f, 0x46, 0xa7, 0xe7, 0x36, 0x76,
	0x3a, 0xd4, 0x91, 0x73, 0x35, 0x3a, 0x1d, 0xea, 0xd9, 0xf9, 0x00, 0x1b, 0x8b, 0x0e, 0x7d, 0x64,
	0x2d, 0xc9, 0xca, 0x09, 0xcf, 0xde, 0x35, 0x64, 0xff, 0x8c, 0xb4, 0x5a, 0x84, 0xc2, 0x71, 0x5a,
	0x7c, 0x0f, 0x2f, 0x7b, 0x21, 0xe3, 0x1b, 0x67, 0x64, 0xac, 0x03, 0x9d, 0x8f, 0x78, 0x03, 0xd4,
	0x14, 0xf6, 0x1d, 0x8f, 0x25, 0xd0, 0xfa, 0x32, 0xcc, 0x1f, 0xf0, 0xb8, 0xfd, 0x93, 0xc0, 0xe6,
	0xac, 0xb3, 0xa1, 0x47, 0x07, 0xec, 0xb8, 0x38, 0x32, 0x71, 0x90, 0xa8, 0x7d, 0x79, 0x53, 0x25,
	0xbb, 0xb4, 0x59, 0x8f, 0xf0, 0x8e, 0x58, 0xf1, 0x16, 0x40, 0xc8, 0xfd, 0x64, 0xdc, 0x92, 0x91,
	0x09, 0xb9, 0x3f, 0x2d, 0xff, 0x4a, 0x20, 0x75, 0x83, 0xa2, 0xf6, 0x95, 0x1f, 0xab, 0xc5, 0x6e,
	0x75, 0xf3, 0x33, 0xa9, 0xff, 0xc7, 0xdb, 0xad, 0xcf, 0x5e, 0x52, 0xb3, 0x74, 0x83, 0xfa, 0xc4,
	0x3d, 0x6f, 0xaf, 0xe8, 0x8f, 0xba, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x18, 0xae, 0x66, 0xcc,
	0x6a, 0x03, 0x00, 0x00,
}
