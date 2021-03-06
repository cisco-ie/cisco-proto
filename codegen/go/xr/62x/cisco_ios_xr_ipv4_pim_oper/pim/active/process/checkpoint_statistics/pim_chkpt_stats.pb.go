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
// source: pim_chkpt_stats.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_process_checkpoint_statistics

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

type PimChkptStats_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimChkptStats_KEYS) Reset()         { *m = PimChkptStats_KEYS{} }
func (m *PimChkptStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimChkptStats_KEYS) ProtoMessage()    {}
func (*PimChkptStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb0aa4ddd17d903, []int{0}
}

func (m *PimChkptStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimChkptStats_KEYS.Unmarshal(m, b)
}
func (m *PimChkptStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimChkptStats_KEYS.Marshal(b, m, deterministic)
}
func (m *PimChkptStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimChkptStats_KEYS.Merge(m, src)
}
func (m *PimChkptStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimChkptStats_KEYS.Size(m)
}
func (m *PimChkptStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimChkptStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimChkptStats_KEYS proto.InternalMessageInfo

type PimChkptTableStats struct {
	TableDescription     string   `protobuf:"bytes,1,opt,name=table_description,json=tableDescription,proto3" json:"table_description,omitempty"`
	TableName            uint32   `protobuf:"varint,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	IsMirrored           bool     `protobuf:"varint,3,opt,name=is_mirrored,json=isMirrored,proto3" json:"is_mirrored,omitempty"`
	Statistic            []uint32 `protobuf:"varint,4,rep,packed,name=statistic,proto3" json:"statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimChkptTableStats) Reset()         { *m = PimChkptTableStats{} }
func (m *PimChkptTableStats) String() string { return proto.CompactTextString(m) }
func (*PimChkptTableStats) ProtoMessage()    {}
func (*PimChkptTableStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb0aa4ddd17d903, []int{1}
}

func (m *PimChkptTableStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimChkptTableStats.Unmarshal(m, b)
}
func (m *PimChkptTableStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimChkptTableStats.Marshal(b, m, deterministic)
}
func (m *PimChkptTableStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimChkptTableStats.Merge(m, src)
}
func (m *PimChkptTableStats) XXX_Size() int {
	return xxx_messageInfo_PimChkptTableStats.Size(m)
}
func (m *PimChkptTableStats) XXX_DiscardUnknown() {
	xxx_messageInfo_PimChkptTableStats.DiscardUnknown(m)
}

var xxx_messageInfo_PimChkptTableStats proto.InternalMessageInfo

func (m *PimChkptTableStats) GetTableDescription() string {
	if m != nil {
		return m.TableDescription
	}
	return ""
}

func (m *PimChkptTableStats) GetTableName() uint32 {
	if m != nil {
		return m.TableName
	}
	return 0
}

func (m *PimChkptTableStats) GetIsMirrored() bool {
	if m != nil {
		return m.IsMirrored
	}
	return false
}

func (m *PimChkptTableStats) GetStatistic() []uint32 {
	if m != nil {
		return m.Statistic
	}
	return nil
}

type PimChkptStats struct {
	CheckpointTable      []*PimChkptTableStats `protobuf:"bytes,50,rep,name=checkpoint_table,json=checkpointTable,proto3" json:"checkpoint_table,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PimChkptStats) Reset()         { *m = PimChkptStats{} }
func (m *PimChkptStats) String() string { return proto.CompactTextString(m) }
func (*PimChkptStats) ProtoMessage()    {}
func (*PimChkptStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eb0aa4ddd17d903, []int{2}
}

func (m *PimChkptStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimChkptStats.Unmarshal(m, b)
}
func (m *PimChkptStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimChkptStats.Marshal(b, m, deterministic)
}
func (m *PimChkptStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimChkptStats.Merge(m, src)
}
func (m *PimChkptStats) XXX_Size() int {
	return xxx_messageInfo_PimChkptStats.Size(m)
}
func (m *PimChkptStats) XXX_DiscardUnknown() {
	xxx_messageInfo_PimChkptStats.DiscardUnknown(m)
}

var xxx_messageInfo_PimChkptStats proto.InternalMessageInfo

func (m *PimChkptStats) GetCheckpointTable() []*PimChkptTableStats {
	if m != nil {
		return m.CheckpointTable
	}
	return nil
}

func init() {
	proto.RegisterType((*PimChkptStats_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.checkpoint_statistics.pim_chkpt_stats_KEYS")
	proto.RegisterType((*PimChkptTableStats)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.checkpoint_statistics.pim_chkpt_table_stats")
	proto.RegisterType((*PimChkptStats)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.process.checkpoint_statistics.pim_chkpt_stats")
}

func init() { proto.RegisterFile("pim_chkpt_stats.proto", fileDescriptor_8eb0aa4ddd17d903) }

var fileDescriptor_8eb0aa4ddd17d903 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0x2b, 0xe2, 0x4e, 0x91, 0xd6, 0xa0, 0x92, 0x83, 0x62, 0xd8, 0xd3, 0x82, 0x90,
	0x43, 0xf5, 0x0d, 0xd4, 0x93, 0xe8, 0x61, 0xf5, 0xa2, 0x97, 0x90, 0xa6, 0x81, 0x0e, 0x35, 0x9b,
	0x90, 0x09, 0xc5, 0x17, 0xf0, 0x21, 0xbc, 0xfa, 0xa4, 0xb2, 0xbb, 0xd2, 0x96, 0xe2, 0xd1, 0x63,
	0xbe, 0xf9, 0x32, 0x33, 0xf9, 0x03, 0xa7, 0x11, 0xbd, 0xb6, 0x8b, 0x65, 0xcc, 0x9a, 0xb2, 0xc9,
	0xa4, 0x62, 0x0a, 0x39, 0xf0, 0x5b, 0x8b, 0x64, 0x83, 0xc6, 0x40, 0xfa, 0x23, 0x69, 0x8c, 0xab,
	0x1b, 0xdd, 0x89, 0x21, 0xba, 0xa4, 0x22, 0x7a, 0x65, 0x6c, 0xc6, 0x95, 0xeb, 0x64, 0xeb, 0x88,
	0x94, 0x5d, 0x38, 0xbb, 0x8c, 0x01, 0xdb, 0xa1, 0x0b, 0x52, 0x46, 0x4b, 0xd5, 0x19, 0x9c, 0xec,
	0x74, 0xd7, 0x0f, 0xf7, 0xaf, 0xcf, 0xd5, 0x37, 0xdb, 0x1e, 0x9b, 0xcd, 0xec, 0xdd, 0x0d, 0x65,
	0x7e, 0x05, 0xc7, 0xc3, 0x71, 0xee, 0xc8, 0x26, 0x8c, 0x19, 0x43, 0x2b, 0x98, 0x64, 0x75, 0xd9,
	0x4c, 0xfa, 0xc2, 0xdd, 0x86, 0xf3, 0x0b, 0x80, 0x41, 0x6e, 0x8d, 0x77, 0x62, 0x4f, 0xb2, 0xfa,
	0xa8, 0x29, 0x7b, 0xf2, 0x64, 0xbc, 0xe3, 0x97, 0x30, 0x42, 0xd2, 0x1e, 0x53, 0x0a, 0xc9, 0xcd,
	0x45, 0x21, 0x59, 0x7d, 0xd8, 0x00, 0xd2, 0xe3, 0x2f, 0xe1, 0xe7, 0x50, 0xae, 0x97, 0x15, 0xfb,
	0xb2, 0xe8, 0xae, 0xaf, 0x41, 0xf5, 0xc5, 0x60, 0xbc, 0xb3, 0x3d, 0xff, 0x64, 0x30, 0xd9, 0x7a,
	0x6a, 0x3f, 0x4b, 0x4c, 0x65, 0x51, 0x8f, 0xa6, 0x6f, 0xea, 0x1f, 0x12, 0x53, 0x7f, 0xa6, 0xd2,
	0x8c, 0x37, 0xf2, 0x4b, 0x87, 0x67, 0x07, 0xfd, 0x27, 0x5d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xb2, 0x6d, 0xdb, 0xc6, 0xbd, 0x01, 0x00, 0x00,
}
