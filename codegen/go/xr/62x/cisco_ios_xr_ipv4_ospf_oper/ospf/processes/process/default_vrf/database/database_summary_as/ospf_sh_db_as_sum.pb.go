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
// source: ospf_sh_db_as_sum.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_database_database_summary_as

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

type OspfShDbAsSum_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbAsSum_KEYS) Reset()         { *m = OspfShDbAsSum_KEYS{} }
func (m *OspfShDbAsSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAsSum_KEYS) ProtoMessage()    {}
func (*OspfShDbAsSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dd6c9f4478cf1b, []int{0}
}

func (m *OspfShDbAsSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Unmarshal(m, b)
}
func (m *OspfShDbAsSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShDbAsSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAsSum_KEYS.Merge(m, src)
}
func (m *OspfShDbAsSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAsSum_KEYS.Size(m)
}
func (m *OspfShDbAsSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAsSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAsSum_KEYS proto.InternalMessageInfo

func (m *OspfShDbAsSum_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShDbSumCntr struct {
	LsaType              string   `protobuf:"bytes,1,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	LsaCount             int32    `protobuf:"zigzag32,2,opt,name=lsa_count,json=lsaCount,proto3" json:"lsa_count,omitempty"`
	LsaDeleteCount       int32    `protobuf:"zigzag32,3,opt,name=lsa_delete_count,json=lsaDeleteCount,proto3" json:"lsa_delete_count,omitempty"`
	LsaMaxageCount       int32    `protobuf:"zigzag32,4,opt,name=lsa_maxage_count,json=lsaMaxageCount,proto3" json:"lsa_maxage_count,omitempty"`
	LsaSelfCount         int32    `protobuf:"zigzag32,5,opt,name=lsa_self_count,json=lsaSelfCount,proto3" json:"lsa_self_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShDbSumCntr) Reset()         { *m = OspfShDbSumCntr{} }
func (m *OspfShDbSumCntr) String() string { return proto.CompactTextString(m) }
func (*OspfShDbSumCntr) ProtoMessage()    {}
func (*OspfShDbSumCntr) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dd6c9f4478cf1b, []int{1}
}

func (m *OspfShDbSumCntr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbSumCntr.Unmarshal(m, b)
}
func (m *OspfShDbSumCntr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbSumCntr.Marshal(b, m, deterministic)
}
func (m *OspfShDbSumCntr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbSumCntr.Merge(m, src)
}
func (m *OspfShDbSumCntr) XXX_Size() int {
	return xxx_messageInfo_OspfShDbSumCntr.Size(m)
}
func (m *OspfShDbSumCntr) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbSumCntr.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbSumCntr proto.InternalMessageInfo

func (m *OspfShDbSumCntr) GetLsaType() string {
	if m != nil {
		return m.LsaType
	}
	return ""
}

func (m *OspfShDbSumCntr) GetLsaCount() int32 {
	if m != nil {
		return m.LsaCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaDeleteCount() int32 {
	if m != nil {
		return m.LsaDeleteCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaMaxageCount() int32 {
	if m != nil {
		return m.LsaMaxageCount
	}
	return 0
}

func (m *OspfShDbSumCntr) GetLsaSelfCount() int32 {
	if m != nil {
		return m.LsaSelfCount
	}
	return 0
}

type OspfShDbAsSum struct {
	AsLsaCounter         []*OspfShDbSumCntr `protobuf:"bytes,50,rep,name=as_lsa_counter,json=asLsaCounter,proto3" json:"as_lsa_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OspfShDbAsSum) Reset()         { *m = OspfShDbAsSum{} }
func (m *OspfShDbAsSum) String() string { return proto.CompactTextString(m) }
func (*OspfShDbAsSum) ProtoMessage()    {}
func (*OspfShDbAsSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4dd6c9f4478cf1b, []int{2}
}

func (m *OspfShDbAsSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShDbAsSum.Unmarshal(m, b)
}
func (m *OspfShDbAsSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShDbAsSum.Marshal(b, m, deterministic)
}
func (m *OspfShDbAsSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShDbAsSum.Merge(m, src)
}
func (m *OspfShDbAsSum) XXX_Size() int {
	return xxx_messageInfo_OspfShDbAsSum.Size(m)
}
func (m *OspfShDbAsSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShDbAsSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShDbAsSum proto.InternalMessageInfo

func (m *OspfShDbAsSum) GetAsLsaCounter() []*OspfShDbSumCntr {
	if m != nil {
		return m.AsLsaCounter
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShDbAsSum_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_as_sum_KEYS")
	proto.RegisterType((*OspfShDbSumCntr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_sum_cntr")
	proto.RegisterType((*OspfShDbAsSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.database.database_summary_as.ospf_sh_db_as_sum")
}

func init() { proto.RegisterFile("ospf_sh_db_as_sum.proto", fileDescriptor_b4dd6c9f4478cf1b) }

var fileDescriptor_b4dd6c9f4478cf1b = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x15, 0xca, 0x5f, 0xdd, 0xaa, 0xa2, 0x46, 0x82, 0x22, 0x96, 0x52, 0x31, 0x64, 0xf2,
	0x50, 0xd8, 0x18, 0x81, 0x89, 0x9f, 0xa1, 0x65, 0x41, 0x0c, 0x57, 0xb7, 0xc9, 0x0d, 0x54, 0x72,
	0x6a, 0xcb, 0xd7, 0xad, 0xda, 0x17, 0xe1, 0x29, 0x78, 0x0f, 0x5e, 0x0b, 0xd9, 0xa4, 0x11, 0xa8,
	0xac, 0x6c, 0x27, 0xe7, 0x7e, 0xe7, 0xc6, 0x3a, 0xb6, 0x38, 0x36, 0x6c, 0x0b, 0xe0, 0x37, 0xc8,
	0x27, 0x80, 0x0c, 0x3c, 0x2f, 0x95, 0x75, 0xc6, 0x1b, 0xf9, 0x92, 0x4d, 0x39, 0x33, 0x30, 0x35,
	0x0c, 0x4b, 0x07, 0x53, 0xbb, 0xb8, 0x84, 0x88, 0x1a, 0x4b, 0x4e, 0x05, 0x15, 0xb8, 0x8c, 0x98,
	0x89, 0xd7, 0x4a, 0xe5, 0x54, 0xe0, 0x5c, 0x7b, 0x58, 0xb8, 0x42, 0xe5, 0xe8, 0x71, 0x82, 0x4c,
	0xb5, 0x08, 0xbb, 0x4b, 0x74, 0x2b, 0x40, 0x1e, 0x5c, 0x89, 0xa3, 0x8d, 0xff, 0xc2, 0xdd, 0xed,
	0xf3, 0x58, 0x9e, 0x89, 0x76, 0xb5, 0x0d, 0x66, 0x58, 0x52, 0x2f, 0xe9, 0x27, 0x69, 0x73, 0xd4,
	0xaa, 0xbc, 0x47, 0x2c, 0x69, 0xf0, 0x99, 0x88, 0xc3, 0x1f, 0xe9, 0x10, 0xcd, 0x66, 0xde, 0xc9,
	0x13, 0xb1, 0xaf, 0x19, 0xc1, 0xaf, 0xec, 0x3a, 0xb6, 0xa7, 0x19, 0x9f, 0x56, 0x96, 0xe4, 0xa9,
	0x68, 0x86, 0x51, 0x66, 0xe6, 0x33, 0xdf, 0xdb, 0xea, 0x27, 0x69, 0x77, 0x14, 0xd8, 0xeb, 0xf0,
	0x2d, 0x53, 0x71, 0x10, 0x86, 0x39, 0x69, 0xf2, 0x54, 0x31, 0x8d, 0xc8, 0x74, 0x34, 0xe3, 0x4d,
	0xb4, 0x7f, 0x91, 0x25, 0x2e, 0xf1, 0x75, 0x4d, 0x6e, 0xd7, 0xe4, 0x43, 0xb4, 0xbf, 0xc9, 0x73,
	0x11, 0x1c, 0x60, 0xd2, 0x45, 0xc5, 0xed, 0x44, 0xae, 0xad, 0x19, 0xc7, 0xa4, 0x8b, 0x48, 0x0d,
	0x3e, 0x12, 0xd1, 0xdd, 0xe8, 0x41, 0xbe, 0x27, 0xa2, 0x83, 0x0c, 0xf5, 0x81, 0xc9, 0xf5, 0x86,
	0xfd, 0x46, 0xda, 0x1a, 0x5a, 0xf5, 0x8f, 0x77, 0xa2, 0xfe, 0xa8, 0x74, 0xd4, 0x46, 0xbe, 0xaf,
	0x6a, 0x22, 0x37, 0xd9, 0x8d, 0x2f, 0xe3, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x3b, 0xc8,
	0xdd, 0x34, 0x02, 0x00, 0x00,
}
