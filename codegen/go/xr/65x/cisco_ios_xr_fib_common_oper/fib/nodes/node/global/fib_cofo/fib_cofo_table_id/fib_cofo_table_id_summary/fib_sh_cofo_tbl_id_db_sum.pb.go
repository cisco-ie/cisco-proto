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
// source: fib_sh_cofo_tbl_id_db_sum.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_global_fib_cofo_fib_cofo_table_id_fib_cofo_table_id_summary

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

type FibShCofoTblIdDbSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShCofoTblIdDbSum_KEYS) Reset()         { *m = FibShCofoTblIdDbSum_KEYS{} }
func (m *FibShCofoTblIdDbSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShCofoTblIdDbSum_KEYS) ProtoMessage()    {}
func (*FibShCofoTblIdDbSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db6e5dd4c547a5c, []int{0}
}

func (m *FibShCofoTblIdDbSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShCofoTblIdDbSum_KEYS.Unmarshal(m, b)
}
func (m *FibShCofoTblIdDbSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShCofoTblIdDbSum_KEYS.Marshal(b, m, deterministic)
}
func (m *FibShCofoTblIdDbSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShCofoTblIdDbSum_KEYS.Merge(m, src)
}
func (m *FibShCofoTblIdDbSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShCofoTblIdDbSum_KEYS.Size(m)
}
func (m *FibShCofoTblIdDbSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShCofoTblIdDbSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShCofoTblIdDbSum_KEYS proto.InternalMessageInfo

func (m *FibShCofoTblIdDbSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type FibShCofoTblIdPerSdr struct {
	Sdrid                uint32   `protobuf:"varint,1,opt,name=sdrid,proto3" json:"sdrid,omitempty"`
	NumOfipv4Tables      uint32   `protobuf:"varint,2,opt,name=num_ofipv4_tables,json=numOfipv4Tables,proto3" json:"num_ofipv4_tables,omitempty"`
	NumOfipv6Tables      uint32   `protobuf:"varint,3,opt,name=num_ofipv6_tables,json=numOfipv6Tables,proto3" json:"num_ofipv6_tables,omitempty"`
	NumOfmplsTables      uint32   `protobuf:"varint,4,opt,name=num_ofmpls_tables,json=numOfmplsTables,proto3" json:"num_ofmpls_tables,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShCofoTblIdPerSdr) Reset()         { *m = FibShCofoTblIdPerSdr{} }
func (m *FibShCofoTblIdPerSdr) String() string { return proto.CompactTextString(m) }
func (*FibShCofoTblIdPerSdr) ProtoMessage()    {}
func (*FibShCofoTblIdPerSdr) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db6e5dd4c547a5c, []int{1}
}

func (m *FibShCofoTblIdPerSdr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShCofoTblIdPerSdr.Unmarshal(m, b)
}
func (m *FibShCofoTblIdPerSdr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShCofoTblIdPerSdr.Marshal(b, m, deterministic)
}
func (m *FibShCofoTblIdPerSdr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShCofoTblIdPerSdr.Merge(m, src)
}
func (m *FibShCofoTblIdPerSdr) XXX_Size() int {
	return xxx_messageInfo_FibShCofoTblIdPerSdr.Size(m)
}
func (m *FibShCofoTblIdPerSdr) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShCofoTblIdPerSdr.DiscardUnknown(m)
}

var xxx_messageInfo_FibShCofoTblIdPerSdr proto.InternalMessageInfo

func (m *FibShCofoTblIdPerSdr) GetSdrid() uint32 {
	if m != nil {
		return m.Sdrid
	}
	return 0
}

func (m *FibShCofoTblIdPerSdr) GetNumOfipv4Tables() uint32 {
	if m != nil {
		return m.NumOfipv4Tables
	}
	return 0
}

func (m *FibShCofoTblIdPerSdr) GetNumOfipv6Tables() uint32 {
	if m != nil {
		return m.NumOfipv6Tables
	}
	return 0
}

func (m *FibShCofoTblIdPerSdr) GetNumOfmplsTables() uint32 {
	if m != nil {
		return m.NumOfmplsTables
	}
	return 0
}

type FibShCofoTblIdDbSum struct {
	NumberOfTblIdAllocated uint32                  `protobuf:"varint,50,opt,name=number_of_tbl_id_allocated,json=numberOfTblIdAllocated,proto3" json:"number_of_tbl_id_allocated,omitempty"`
	TblIdMinimum           uint32                  `protobuf:"varint,51,opt,name=tbl_id_minimum,json=tblIdMinimum,proto3" json:"tbl_id_minimum,omitempty"`
	TblIdMaximum           uint32                  `protobuf:"varint,52,opt,name=tbl_id_maximum,json=tblIdMaximum,proto3" json:"tbl_id_maximum,omitempty"`
	TblIdLastAllocated     uint32                  `protobuf:"varint,53,opt,name=tbl_id_last_allocated,json=tblIdLastAllocated,proto3" json:"tbl_id_last_allocated,omitempty"`
	TblIdDefaultV4         uint32                  `protobuf:"varint,54,opt,name=tbl_id_default_v4,json=tblIdDefaultV4,proto3" json:"tbl_id_default_v4,omitempty"`
	TblIdDefaultV6         uint32                  `protobuf:"varint,55,opt,name=tbl_id_default_v6,json=tblIdDefaultV6,proto3" json:"tbl_id_default_v6,omitempty"`
	NumberOfTblIdInGc      uint32                  `protobuf:"varint,56,opt,name=number_of_tbl_id_in_gc,json=numberOfTblIdInGc,proto3" json:"number_of_tbl_id_in_gc,omitempty"`
	PerSdrStat             []*FibShCofoTblIdPerSdr `protobuf:"bytes,57,rep,name=per_sdr_stat,json=perSdrStat,proto3" json:"per_sdr_stat,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *FibShCofoTblIdDbSum) Reset()         { *m = FibShCofoTblIdDbSum{} }
func (m *FibShCofoTblIdDbSum) String() string { return proto.CompactTextString(m) }
func (*FibShCofoTblIdDbSum) ProtoMessage()    {}
func (*FibShCofoTblIdDbSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5db6e5dd4c547a5c, []int{2}
}

func (m *FibShCofoTblIdDbSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShCofoTblIdDbSum.Unmarshal(m, b)
}
func (m *FibShCofoTblIdDbSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShCofoTblIdDbSum.Marshal(b, m, deterministic)
}
func (m *FibShCofoTblIdDbSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShCofoTblIdDbSum.Merge(m, src)
}
func (m *FibShCofoTblIdDbSum) XXX_Size() int {
	return xxx_messageInfo_FibShCofoTblIdDbSum.Size(m)
}
func (m *FibShCofoTblIdDbSum) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShCofoTblIdDbSum.DiscardUnknown(m)
}

var xxx_messageInfo_FibShCofoTblIdDbSum proto.InternalMessageInfo

func (m *FibShCofoTblIdDbSum) GetNumberOfTblIdAllocated() uint32 {
	if m != nil {
		return m.NumberOfTblIdAllocated
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetTblIdMinimum() uint32 {
	if m != nil {
		return m.TblIdMinimum
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetTblIdMaximum() uint32 {
	if m != nil {
		return m.TblIdMaximum
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetTblIdLastAllocated() uint32 {
	if m != nil {
		return m.TblIdLastAllocated
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetTblIdDefaultV4() uint32 {
	if m != nil {
		return m.TblIdDefaultV4
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetTblIdDefaultV6() uint32 {
	if m != nil {
		return m.TblIdDefaultV6
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetNumberOfTblIdInGc() uint32 {
	if m != nil {
		return m.NumberOfTblIdInGc
	}
	return 0
}

func (m *FibShCofoTblIdDbSum) GetPerSdrStat() []*FibShCofoTblIdPerSdr {
	if m != nil {
		return m.PerSdrStat
	}
	return nil
}

func init() {
	proto.RegisterType((*FibShCofoTblIdDbSum_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.global.fib_cofo.fib_cofo_table_id.fib_cofo_table_id_summary.fib_sh_cofo_tbl_id_db_sum_KEYS")
	proto.RegisterType((*FibShCofoTblIdPerSdr)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.global.fib_cofo.fib_cofo_table_id.fib_cofo_table_id_summary.fib_sh_cofo_tbl_id_per_sdr")
	proto.RegisterType((*FibShCofoTblIdDbSum)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.global.fib_cofo.fib_cofo_table_id.fib_cofo_table_id_summary.fib_sh_cofo_tbl_id_db_sum")
}

func init() { proto.RegisterFile("fib_sh_cofo_tbl_id_db_sum.proto", fileDescriptor_5db6e5dd4c547a5c) }

var fileDescriptor_5db6e5dd4c547a5c = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0xd3, 0x4d, 0x6b, 0x14, 0x31,
	0x18, 0x07, 0x70, 0xc6, 0x56, 0xb1, 0xb1, 0x56, 0x36, 0x68, 0x89, 0x15, 0xb4, 0x2c, 0x1e, 0xaa,
	0x87, 0x81, 0xb6, 0xeb, 0xf8, 0x02, 0x1e, 0x04, 0x45, 0x8a, 0x2f, 0x85, 0xdd, 0x22, 0x78, 0x7a,
	0x48, 0x26, 0xc9, 0x1a, 0xc8, 0xcb, 0x90, 0x64, 0x4a, 0x3d, 0x7b, 0xf1, 0x93, 0xf8, 0x11, 0xfc,
	0x7c, 0x32, 0xc9, 0x8e, 0xce, 0xda, 0xdd, 0x6b, 0x2f, 0xcb, 0xce, 0x3f, 0xbf, 0xcc, 0xf3, 0x24,
	0x3c, 0x83, 0x1e, 0x49, 0xc5, 0x20, 0x7c, 0x83, 0xda, 0x49, 0x07, 0x91, 0x69, 0x50, 0x1c, 0x38,
	0x83, 0xd0, 0x9a, 0xb2, 0xf1, 0x2e, 0x3a, 0x3c, 0xaf, 0x55, 0xa8, 0x1d, 0x28, 0x17, 0xe0, 0xc2,
	0x43, 0xa7, 0x6b, 0x67, 0x8c, 0xb3, 0xe0, 0x1a, 0xe1, 0x4b, 0xa9, 0x58, 0x69, 0x1d, 0x17, 0x21,
	0xfd, 0x96, 0x73, 0xed, 0x18, 0xd5, 0x65, 0x56, 0xd2, 0xfd, 0xfd, 0x03, 0x91, 0x32, 0x2d, 0x40,
	0xf1, 0xcb, 0x49, 0x57, 0xca, 0x50, 0xff, 0x7d, 0xfc, 0x1a, 0x3d, 0x5c, 0xdb, 0x0b, 0x7c, 0x78,
	0xf7, 0x75, 0x86, 0x1f, 0xa0, 0xad, 0xae, 0x0a, 0x58, 0x6a, 0x04, 0x29, 0xf6, 0x8b, 0x83, 0xad,
	0xe9, 0xcd, 0x2e, 0xf8, 0x4c, 0x8d, 0x18, 0xff, 0x2e, 0xd0, 0xde, 0x8a, 0xfd, 0x8d, 0xf0, 0x10,
	0xb8, 0xc7, 0x77, 0xd1, 0xf5, 0xc0, 0xbd, 0xe2, 0x69, 0xdf, 0xed, 0x69, 0x7e, 0xc0, 0x4f, 0xd1,
	0xc8, 0xb6, 0x06, 0x9c, 0x54, 0xcd, 0xf9, 0x24, 0xb7, 0x14, 0xc8, 0xb5, 0x24, 0xee, 0xd8, 0xd6,
	0x9c, 0xa6, 0xfc, 0x2c, 0xc5, 0x4b, 0xb6, 0xea, 0xed, 0xc6, 0xb2, 0xad, 0xfe, 0xb7, 0xa6, 0xd1,
	0xa1, 0xb7, 0x9b, 0x03, 0xdb, 0xe5, 0xd9, 0x8e, 0x7f, 0x6e, 0xa2, 0xfb, 0x6b, 0x0f, 0x8e, 0x5f,
	0xa1, 0x3d, 0xdb, 0x1a, 0x26, 0x3c, 0x38, 0xd9, 0x2f, 0x51, 0xad, 0x5d, 0x4d, 0xa3, 0xe0, 0xe4,
	0x28, 0xbd, 0x72, 0x37, 0x8b, 0x53, 0x79, 0xc6, 0xf4, 0x09, 0x7f, 0xd3, 0xaf, 0xe2, 0xc7, 0x68,
	0x67, 0xb1, 0xc3, 0x28, 0xab, 0x4c, 0x6b, 0xc8, 0x71, 0xf2, 0xdb, 0xb1, 0x73, 0x9f, 0x72, 0x36,
	0x54, 0xf4, 0x22, 0xa9, 0xc9, 0x50, 0xe5, 0x0c, 0x1f, 0xa2, 0x7b, 0x0b, 0xa5, 0x69, 0x88, 0x83,
	0x16, 0x9e, 0x25, 0x8c, 0x13, 0xfe, 0x48, 0x43, 0xfc, 0x57, 0xfe, 0x09, 0x1a, 0xf5, 0x67, 0x11,
	0x92, 0xb6, 0x3a, 0xc2, 0xf9, 0x84, 0x54, 0x89, 0xef, 0x24, 0xfe, 0x36, 0xc7, 0x5f, 0x26, 0xab,
	0x68, 0x45, 0x9e, 0xaf, 0xa0, 0x15, 0x3e, 0x44, 0xbb, 0x97, 0x2e, 0x44, 0x59, 0x98, 0xd7, 0xe4,
	0x45, 0xf2, 0xa3, 0xa5, 0xcb, 0x38, 0xb1, 0xef, 0x6b, 0xfc, 0xab, 0x40, 0xdb, 0x8b, 0x39, 0x80,
	0x10, 0x69, 0x24, 0x2f, 0xf7, 0x37, 0x0e, 0x6e, 0x1d, 0xfd, 0x28, 0xca, 0x2b, 0x9a, 0xed, 0x72,
	0xfd, 0x60, 0x4e, 0x51, 0x23, 0xfc, 0x8c, 0xfb, 0x59, 0xa4, 0x91, 0xdd, 0x48, 0x9f, 0xdc, 0xf1,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x18, 0xcf, 0xcc, 0x95, 0x03, 0x00, 0x00,
}
