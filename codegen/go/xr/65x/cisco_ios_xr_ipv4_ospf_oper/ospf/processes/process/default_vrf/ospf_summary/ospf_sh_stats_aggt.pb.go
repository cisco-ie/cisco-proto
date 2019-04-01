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
// source: ospf_sh_stats_aggt.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_ospf_summary

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

type OspfShStatsAggt_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShStatsAggt_KEYS) Reset()         { *m = OspfShStatsAggt_KEYS{} }
func (m *OspfShStatsAggt_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShStatsAggt_KEYS) ProtoMessage()    {}
func (*OspfShStatsAggt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_77201b28090bdeda, []int{0}
}

func (m *OspfShStatsAggt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatsAggt_KEYS.Unmarshal(m, b)
}
func (m *OspfShStatsAggt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatsAggt_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShStatsAggt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatsAggt_KEYS.Merge(m, src)
}
func (m *OspfShStatsAggt_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShStatsAggt_KEYS.Size(m)
}
func (m *OspfShStatsAggt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatsAggt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatsAggt_KEYS proto.InternalMessageInfo

func (m *OspfShStatsAggt_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShStatsAggt struct {
	SaNumNbrs            uint32   `protobuf:"varint,50,opt,name=sa_num_nbrs,json=saNumNbrs,proto3" json:"sa_num_nbrs,omitempty"`
	SaNumNbrsUp          uint32   `protobuf:"varint,51,opt,name=sa_num_nbrs_up,json=saNumNbrsUp,proto3" json:"sa_num_nbrs_up,omitempty"`
	SaNumIntf            uint32   `protobuf:"varint,52,opt,name=sa_num_intf,json=saNumIntf,proto3" json:"sa_num_intf,omitempty"`
	SaNumIntfUp          uint32   `protobuf:"varint,53,opt,name=sa_num_intf_up,json=saNumIntfUp,proto3" json:"sa_num_intf_up,omitempty"`
	SaNumVintfUp         uint32   `protobuf:"varint,54,opt,name=sa_num_vintf_up,json=saNumVintfUp,proto3" json:"sa_num_vintf_up,omitempty"`
	SaNumSlintfUp        uint32   `protobuf:"varint,55,opt,name=sa_num_slintf_up,json=saNumSlintfUp,proto3" json:"sa_num_slintf_up,omitempty"`
	SaNumAreas           uint32   `protobuf:"varint,56,opt,name=sa_num_areas,json=saNumAreas,proto3" json:"sa_num_areas,omitempty"`
	SaLsaCntTypeRtr      uint32   `protobuf:"varint,57,opt,name=sa_lsa_cnt_type_rtr,json=saLsaCntTypeRtr,proto3" json:"sa_lsa_cnt_type_rtr,omitempty"`
	SaLsaCntTypeNet      uint32   `protobuf:"varint,58,opt,name=sa_lsa_cnt_type_net,json=saLsaCntTypeNet,proto3" json:"sa_lsa_cnt_type_net,omitempty"`
	SaLsaCntTypeSumNet   uint32   `protobuf:"varint,59,opt,name=sa_lsa_cnt_type_sum_net,json=saLsaCntTypeSumNet,proto3" json:"sa_lsa_cnt_type_sum_net,omitempty"`
	SaLsaCntTypeSumAsb   uint32   `protobuf:"varint,60,opt,name=sa_lsa_cnt_type_sum_asb,json=saLsaCntTypeSumAsb,proto3" json:"sa_lsa_cnt_type_sum_asb,omitempty"`
	SaLsaCntTypeAse      uint32   `protobuf:"varint,61,opt,name=sa_lsa_cnt_type_ase,json=saLsaCntTypeAse,proto3" json:"sa_lsa_cnt_type_ase,omitempty"`
	SaLsaCntTypeMospf    uint32   `protobuf:"varint,62,opt,name=sa_lsa_cnt_type_mospf,json=saLsaCntTypeMospf,proto3" json:"sa_lsa_cnt_type_mospf,omitempty"`
	SaLsaCntType_7Ase    uint32   `protobuf:"varint,63,opt,name=sa_lsa_cnt_type_7ase,json=saLsaCntType7ase,proto3" json:"sa_lsa_cnt_type_7ase,omitempty"`
	SaLsaCntType_8Ignore uint32   `protobuf:"varint,64,opt,name=sa_lsa_cnt_type_8_ignore,json=saLsaCntType8Ignore,proto3" json:"sa_lsa_cnt_type_8_ignore,omitempty"`
	SaLsaCntTypeOpqLink  uint32   `protobuf:"varint,65,opt,name=sa_lsa_cnt_type_opq_link,json=saLsaCntTypeOpqLink,proto3" json:"sa_lsa_cnt_type_opq_link,omitempty"`
	SaLsaCntTypeOpqArea  uint32   `protobuf:"varint,66,opt,name=sa_lsa_cnt_type_opq_area,json=saLsaCntTypeOpqArea,proto3" json:"sa_lsa_cnt_type_opq_area,omitempty"`
	SaLsaCntTypeOpqAs    uint32   `protobuf:"varint,67,opt,name=sa_lsa_cnt_type_opq_as,json=saLsaCntTypeOpqAs,proto3" json:"sa_lsa_cnt_type_opq_as,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShStatsAggt) Reset()         { *m = OspfShStatsAggt{} }
func (m *OspfShStatsAggt) String() string { return proto.CompactTextString(m) }
func (*OspfShStatsAggt) ProtoMessage()    {}
func (*OspfShStatsAggt) Descriptor() ([]byte, []int) {
	return fileDescriptor_77201b28090bdeda, []int{1}
}

func (m *OspfShStatsAggt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShStatsAggt.Unmarshal(m, b)
}
func (m *OspfShStatsAggt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShStatsAggt.Marshal(b, m, deterministic)
}
func (m *OspfShStatsAggt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShStatsAggt.Merge(m, src)
}
func (m *OspfShStatsAggt) XXX_Size() int {
	return xxx_messageInfo_OspfShStatsAggt.Size(m)
}
func (m *OspfShStatsAggt) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShStatsAggt.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShStatsAggt proto.InternalMessageInfo

func (m *OspfShStatsAggt) GetSaNumNbrs() uint32 {
	if m != nil {
		return m.SaNumNbrs
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumNbrsUp() uint32 {
	if m != nil {
		return m.SaNumNbrsUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumIntf() uint32 {
	if m != nil {
		return m.SaNumIntf
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumIntfUp() uint32 {
	if m != nil {
		return m.SaNumIntfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumVintfUp() uint32 {
	if m != nil {
		return m.SaNumVintfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumSlintfUp() uint32 {
	if m != nil {
		return m.SaNumSlintfUp
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaNumAreas() uint32 {
	if m != nil {
		return m.SaNumAreas
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeRtr() uint32 {
	if m != nil {
		return m.SaLsaCntTypeRtr
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeNet() uint32 {
	if m != nil {
		return m.SaLsaCntTypeNet
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeSumNet() uint32 {
	if m != nil {
		return m.SaLsaCntTypeSumNet
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeSumAsb() uint32 {
	if m != nil {
		return m.SaLsaCntTypeSumAsb
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeAse() uint32 {
	if m != nil {
		return m.SaLsaCntTypeAse
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeMospf() uint32 {
	if m != nil {
		return m.SaLsaCntTypeMospf
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntType_7Ase() uint32 {
	if m != nil {
		return m.SaLsaCntType_7Ase
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntType_8Ignore() uint32 {
	if m != nil {
		return m.SaLsaCntType_8Ignore
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqLink() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqLink
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqArea() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqArea
	}
	return 0
}

func (m *OspfShStatsAggt) GetSaLsaCntTypeOpqAs() uint32 {
	if m != nil {
		return m.SaLsaCntTypeOpqAs
	}
	return 0
}

func init() {
	proto.RegisterType((*OspfShStatsAggt_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.ospf_summary.ospf_sh_stats_aggt_KEYS")
	proto.RegisterType((*OspfShStatsAggt)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.ospf_summary.ospf_sh_stats_aggt")
}

func init() { proto.RegisterFile("ospf_sh_stats_aggt.proto", fileDescriptor_77201b28090bdeda) }

var fileDescriptor_77201b28090bdeda = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd3, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0x71, 0xf5, 0x32, 0x69, 0x6e, 0xc7, 0x86, 0x07, 0xec, 0x9d, 0x50, 0x19, 0x42, 0xec,
	0x80, 0x22, 0xa0, 0x1b, 0x2b, 0x30, 0x7e, 0x94, 0x89, 0xc3, 0xb4, 0x51, 0xa4, 0x96, 0x21, 0x71,
	0x7a, 0x72, 0x8a, 0x53, 0xa2, 0x25, 0xb6, 0xe7, 0xe7, 0x54, 0xf4, 0x9f, 0xe3, 0x6f, 0x43, 0xf6,
	0x52, 0x64, 0x35, 0x61, 0xb7, 0xc8, 0xef, 0xf3, 0x7d, 0x89, 0x92, 0x96, 0x81, 0x26, 0x93, 0x21,
	0xfd, 0x42, 0x72, 0xc2, 0x11, 0x8a, 0xf9, 0xdc, 0x25, 0xc6, 0x6a, 0xa7, 0xf9, 0xf9, 0x2c, 0xa7,
	0x99, 0xc6, 0x5c, 0x13, 0xfe, 0xb6, 0x98, 0x9b, 0xc5, 0x21, 0x06, 0xab, 0x8d, 0xb4, 0x89, 0xbf,
	0xf2, 0x6e, 0x26, 0x89, 0x24, 0xad, 0xae, 0x92, 0x9f, 0x32, 0x13, 0x55, 0xe1, 0x70, 0x61, 0xb3,
	0xe4, 0x66, 0x71, 0x55, 0x96, 0xc2, 0x2e, 0xf7, 0x4f, 0xd8, 0x5e, 0xf3, 0x46, 0x78, 0xfe, 0xf9,
	0xc7, 0x94, 0x3f, 0x62, 0xbd, 0x3a, 0x47, 0x25, 0x4a, 0x09, 0x9d, 0x7e, 0xe7, 0x60, 0x73, 0xd2,
	0xad, 0xcf, 0xc6, 0xa2, 0x94, 0xfb, 0x7f, 0x36, 0x18, 0x6f, 0xe6, 0xfc, 0x21, 0xeb, 0x92, 0x40,
	0x55, 0x95, 0xa8, 0x52, 0x4b, 0xf0, 0xb2, 0xdf, 0x39, 0xd8, 0x9a, 0x6c, 0x92, 0x18, 0x57, 0xe5,
	0x38, 0xb5, 0xc4, 0x1f, 0xb3, 0x3b, 0xd1, 0x1c, 0x2b, 0x03, 0x83, 0x40, 0xba, 0xff, 0xc8, 0xa5,
	0x89, 0x96, 0xe4, 0xca, 0x65, 0x70, 0x18, 0x2d, 0x39, 0x53, 0x2e, 0x8b, 0x96, 0xf8, 0xb9, 0x5f,
	0x72, 0x14, 0x2d, 0xf1, 0xe4, 0xd2, 0xf0, 0x27, 0x6c, 0xbb, 0x46, 0x8b, 0x95, 0x7a, 0x15, 0x54,
	0x2f, 0xa8, 0xef, 0xf9, 0x0d, 0x7b, 0xca, 0x76, 0x6a, 0x46, 0xc5, 0xca, 0x1d, 0x07, 0xb7, 0x15,
	0xdc, 0xb4, 0xa8, 0x61, 0x9f, 0xf5, 0x6a, 0x28, 0xac, 0x14, 0x04, 0xc3, 0x80, 0x58, 0x40, 0x23,
	0x7f, 0xc2, 0x9f, 0xb1, 0x5d, 0x12, 0x58, 0x90, 0xc0, 0x99, 0x72, 0xe8, 0x96, 0x46, 0xa2, 0x75,
	0x16, 0x5e, 0x07, 0xb8, 0x4d, 0xe2, 0x82, 0xc4, 0xa9, 0x72, 0xdf, 0x96, 0x46, 0x4e, 0x9c, 0x6d,
	0xd3, 0x4a, 0x3a, 0x78, 0xd3, 0xd4, 0x63, 0xe9, 0xf8, 0x80, 0xed, 0xad, 0x6b, 0xf2, 0x2f, 0x51,
	0x3a, 0x78, 0x1b, 0x0a, 0x1e, 0x17, 0xd3, 0xaa, 0xbc, 0x25, 0x12, 0x94, 0xc2, 0x49, 0x6b, 0x34,
	0xa2, 0xb4, 0xed, 0xb9, 0x04, 0x49, 0x78, 0xd7, 0x7c, 0xae, 0x11, 0x49, 0xfe, 0x9c, 0xdd, 0x5f,
	0xd7, 0xa5, 0xff, 0x59, 0xc0, 0xfb, 0xe0, 0xef, 0xc6, 0xfe, 0x8b, 0x1f, 0xf0, 0x84, 0xdd, 0x5b,
	0x2f, 0x8e, 0xfd, 0x0d, 0x3e, 0x84, 0x60, 0x27, 0x0e, 0xfc, 0x39, 0x3f, 0x62, 0xb0, 0xee, 0x87,
	0x98, 0xcf, 0x95, 0xb6, 0x12, 0x3e, 0x86, 0x66, 0x37, 0x6e, 0x86, 0x67, 0x61, 0xd4, 0x96, 0x69,
	0x73, 0x8d, 0x45, 0xae, 0xae, 0x60, 0xd4, 0xcc, 0xbe, 0x9a, 0xeb, 0x8b, 0x5c, 0x5d, 0xfd, 0x2f,
	0xf3, 0x9f, 0x1c, 0x3e, 0xb5, 0x66, 0xfe, 0xdb, 0xf3, 0x17, 0xec, 0x41, 0x6b, 0x46, 0x70, 0xda,
	0x7c, 0x0f, 0x3e, 0xa2, 0x74, 0x23, 0xfc, 0xa5, 0x07, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x29,
	0x92, 0x2c, 0x37, 0xee, 0x03, 0x00, 0x00,
}
