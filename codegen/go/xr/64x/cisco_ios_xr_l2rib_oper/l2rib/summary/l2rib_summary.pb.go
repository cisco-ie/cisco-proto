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
// source: l2rib_summary.proto

package cisco_ios_xr_l2rib_oper_l2rib_summary

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

type L2RibSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2RibSummary_KEYS) Reset()         { *m = L2RibSummary_KEYS{} }
func (m *L2RibSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2RibSummary_KEYS) ProtoMessage()    {}
func (*L2RibSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{0}
}

func (m *L2RibSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibSummary_KEYS.Unmarshal(m, b)
}
func (m *L2RibSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2RibSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibSummary_KEYS.Merge(m, src)
}
func (m *L2RibSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2RibSummary_KEYS.Size(m)
}
func (m *L2RibSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibSummary_KEYS proto.InternalMessageInfo

type L2RibBagEc struct {
	CounterType            uint32   `protobuf:"varint,1,opt,name=counter_type,json=counterType,proto3" json:"counter_type,omitempty"`
	CounterName            string   `protobuf:"bytes,2,opt,name=counter_name,json=counterName,proto3" json:"counter_name,omitempty"`
	L2RbFirstEventTs       uint64   `protobuf:"varint,3,opt,name=l2rb_first_event_ts,json=l2rbFirstEventTs,proto3" json:"l2rb_first_event_ts,omitempty"`
	L2RbLastEventTs        uint64   `protobuf:"varint,4,opt,name=l2rb_last_event_ts,json=l2rbLastEventTs,proto3" json:"l2rb_last_event_ts,omitempty"`
	L2RbIntervalEventCount uint32   `protobuf:"varint,5,opt,name=l2rb_interval_event_count,json=l2rbIntervalEventCount,proto3" json:"l2rb_interval_event_count,omitempty"`
	L2RbTotalEventCount    uint32   `protobuf:"varint,6,opt,name=l2rb_total_event_count,json=l2rbTotalEventCount,proto3" json:"l2rb_total_event_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *L2RibBagEc) Reset()         { *m = L2RibBagEc{} }
func (m *L2RibBagEc) String() string { return proto.CompactTextString(m) }
func (*L2RibBagEc) ProtoMessage()    {}
func (*L2RibBagEc) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{1}
}

func (m *L2RibBagEc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibBagEc.Unmarshal(m, b)
}
func (m *L2RibBagEc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibBagEc.Marshal(b, m, deterministic)
}
func (m *L2RibBagEc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibBagEc.Merge(m, src)
}
func (m *L2RibBagEc) XXX_Size() int {
	return xxx_messageInfo_L2RibBagEc.Size(m)
}
func (m *L2RibBagEc) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibBagEc.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibBagEc proto.InternalMessageInfo

func (m *L2RibBagEc) GetCounterType() uint32 {
	if m != nil {
		return m.CounterType
	}
	return 0
}

func (m *L2RibBagEc) GetCounterName() string {
	if m != nil {
		return m.CounterName
	}
	return ""
}

func (m *L2RibBagEc) GetL2RbFirstEventTs() uint64 {
	if m != nil {
		return m.L2RbFirstEventTs
	}
	return 0
}

func (m *L2RibBagEc) GetL2RbLastEventTs() uint64 {
	if m != nil {
		return m.L2RbLastEventTs
	}
	return 0
}

func (m *L2RibBagEc) GetL2RbIntervalEventCount() uint32 {
	if m != nil {
		return m.L2RbIntervalEventCount
	}
	return 0
}

func (m *L2RibBagEc) GetL2RbTotalEventCount() uint32 {
	if m != nil {
		return m.L2RbTotalEventCount
	}
	return 0
}

type L2RibProdUpdateStats struct {
	MemorySize           uint32        `protobuf:"varint,1,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	ObjectCount          uint32        `protobuf:"varint,2,opt,name=object_count,json=objectCount,proto3" json:"object_count,omitempty"`
	EndofIntervalTs      uint64        `protobuf:"varint,3,opt,name=endof_interval_ts,json=endofIntervalTs,proto3" json:"endof_interval_ts,omitempty"`
	ExtendedCounter      []*L2RibBagEc `protobuf:"bytes,4,rep,name=extended_counter,json=extendedCounter,proto3" json:"extended_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *L2RibProdUpdateStats) Reset()         { *m = L2RibProdUpdateStats{} }
func (m *L2RibProdUpdateStats) String() string { return proto.CompactTextString(m) }
func (*L2RibProdUpdateStats) ProtoMessage()    {}
func (*L2RibProdUpdateStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{2}
}

func (m *L2RibProdUpdateStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibProdUpdateStats.Unmarshal(m, b)
}
func (m *L2RibProdUpdateStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibProdUpdateStats.Marshal(b, m, deterministic)
}
func (m *L2RibProdUpdateStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibProdUpdateStats.Merge(m, src)
}
func (m *L2RibProdUpdateStats) XXX_Size() int {
	return xxx_messageInfo_L2RibProdUpdateStats.Size(m)
}
func (m *L2RibProdUpdateStats) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibProdUpdateStats.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibProdUpdateStats proto.InternalMessageInfo

func (m *L2RibProdUpdateStats) GetMemorySize() uint32 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *L2RibProdUpdateStats) GetObjectCount() uint32 {
	if m != nil {
		return m.ObjectCount
	}
	return 0
}

func (m *L2RibProdUpdateStats) GetEndofIntervalTs() uint64 {
	if m != nil {
		return m.EndofIntervalTs
	}
	return 0
}

func (m *L2RibProdUpdateStats) GetExtendedCounter() []*L2RibBagEc {
	if m != nil {
		return m.ExtendedCounter
	}
	return nil
}

type L2RibProdStats struct {
	ProducerId           string                `protobuf:"bytes,1,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	ProducerName         string                `protobuf:"bytes,2,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty"`
	Statistics           *L2RibProdUpdateStats `protobuf:"bytes,3,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *L2RibProdStats) Reset()         { *m = L2RibProdStats{} }
func (m *L2RibProdStats) String() string { return proto.CompactTextString(m) }
func (*L2RibProdStats) ProtoMessage()    {}
func (*L2RibProdStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{3}
}

func (m *L2RibProdStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibProdStats.Unmarshal(m, b)
}
func (m *L2RibProdStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibProdStats.Marshal(b, m, deterministic)
}
func (m *L2RibProdStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibProdStats.Merge(m, src)
}
func (m *L2RibProdStats) XXX_Size() int {
	return xxx_messageInfo_L2RibProdStats.Size(m)
}
func (m *L2RibProdStats) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibProdStats.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibProdStats proto.InternalMessageInfo

func (m *L2RibProdStats) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *L2RibProdStats) GetProducerName() string {
	if m != nil {
		return m.ProducerName
	}
	return ""
}

func (m *L2RibProdStats) GetStatistics() *L2RibProdUpdateStats {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type L2RibTblSummary struct {
	ObjectType           string            `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	ObjectCount          uint32            `protobuf:"varint,2,opt,name=object_count,json=objectCount,proto3" json:"object_count,omitempty"`
	TableMemory          uint32            `protobuf:"varint,3,opt,name=table_memory,json=tableMemory,proto3" json:"table_memory,omitempty"`
	ProducerStat         []*L2RibProdStats `protobuf:"bytes,4,rep,name=producer_stat,json=producerStat,proto3" json:"producer_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2RibTblSummary) Reset()         { *m = L2RibTblSummary{} }
func (m *L2RibTblSummary) String() string { return proto.CompactTextString(m) }
func (*L2RibTblSummary) ProtoMessage()    {}
func (*L2RibTblSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{4}
}

func (m *L2RibTblSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibTblSummary.Unmarshal(m, b)
}
func (m *L2RibTblSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibTblSummary.Marshal(b, m, deterministic)
}
func (m *L2RibTblSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibTblSummary.Merge(m, src)
}
func (m *L2RibTblSummary) XXX_Size() int {
	return xxx_messageInfo_L2RibTblSummary.Size(m)
}
func (m *L2RibTblSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibTblSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibTblSummary proto.InternalMessageInfo

func (m *L2RibTblSummary) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *L2RibTblSummary) GetObjectCount() uint32 {
	if m != nil {
		return m.ObjectCount
	}
	return 0
}

func (m *L2RibTblSummary) GetTableMemory() uint32 {
	if m != nil {
		return m.TableMemory
	}
	return 0
}

func (m *L2RibTblSummary) GetProducerStat() []*L2RibProdStats {
	if m != nil {
		return m.ProducerStat
	}
	return nil
}

type L2RibSummary struct {
	TableSummary         []*L2RibTblSummary `protobuf:"bytes,50,rep,name=table_summary,json=tableSummary,proto3" json:"table_summary,omitempty"`
	ConvergedTablesCount uint32             `protobuf:"varint,51,opt,name=converged_tables_count,json=convergedTablesCount,proto3" json:"converged_tables_count,omitempty"`
	TotalMemory          uint32             `protobuf:"varint,52,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2RibSummary) Reset()         { *m = L2RibSummary{} }
func (m *L2RibSummary) String() string { return proto.CompactTextString(m) }
func (*L2RibSummary) ProtoMessage()    {}
func (*L2RibSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{5}
}

func (m *L2RibSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibSummary.Unmarshal(m, b)
}
func (m *L2RibSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibSummary.Marshal(b, m, deterministic)
}
func (m *L2RibSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibSummary.Merge(m, src)
}
func (m *L2RibSummary) XXX_Size() int {
	return xxx_messageInfo_L2RibSummary.Size(m)
}
func (m *L2RibSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibSummary proto.InternalMessageInfo

func (m *L2RibSummary) GetTableSummary() []*L2RibTblSummary {
	if m != nil {
		return m.TableSummary
	}
	return nil
}

func (m *L2RibSummary) GetConvergedTablesCount() uint32 {
	if m != nil {
		return m.ConvergedTablesCount
	}
	return 0
}

func (m *L2RibSummary) GetTotalMemory() uint32 {
	if m != nil {
		return m.TotalMemory
	}
	return 0
}

func init() {
	proto.RegisterType((*L2RibSummary_KEYS)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_summary_KEYS")
	proto.RegisterType((*L2RibBagEc)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_bag_ec")
	proto.RegisterType((*L2RibProdUpdateStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_prod_update_stats")
	proto.RegisterType((*L2RibProdStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_prod_stats")
	proto.RegisterType((*L2RibTblSummary)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_tbl_summary")
	proto.RegisterType((*L2RibSummary)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_summary")
}

func init() { proto.RegisterFile("l2rib_summary.proto", fileDescriptor_3df371c99cc020d1) }

var fileDescriptor_3df371c99cc020d1 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdf, 0x6e, 0xd3, 0x3e,
	0x18, 0x55, 0xba, 0xfe, 0x26, 0xf5, 0x6b, 0xab, 0x76, 0xfe, 0x4d, 0x23, 0x5c, 0xd1, 0x05, 0x21,
	0x55, 0x20, 0x7a, 0xd1, 0x4e, 0x02, 0x6e, 0xb8, 0x99, 0x86, 0x34, 0xf1, 0xe7, 0x22, 0xed, 0x0d,
	0x12, 0xcc, 0x72, 0x12, 0x6f, 0x32, 0x4a, 0xe2, 0xc8, 0x76, 0xab, 0x75, 0x4f, 0xc3, 0x63, 0xf0,
	0x06, 0xbc, 0x02, 0xaf, 0xc0, 0x5b, 0x20, 0x7f, 0x76, 0xd2, 0x0c, 0x2e, 0xa0, 0x77, 0xdb, 0x39,
	0xe7, 0xb3, 0xcf, 0x39, 0xfe, 0x52, 0xf8, 0x3f, 0x9f, 0x2b, 0x91, 0x50, 0xbd, 0x2e, 0x0a, 0xa6,
	0xb6, 0xb3, 0x4a, 0x49, 0x23, 0xc9, 0x93, 0x54, 0xe8, 0x54, 0x52, 0x21, 0x35, 0xbd, 0x55, 0xd4,
	0x29, 0x64, 0xc5, 0xd5, 0x0c, 0xff, 0x9c, 0x79, 0x71, 0x74, 0x0c, 0xe4, 0xde, 0x34, 0x7d, 0x7b,
	0xf1, 0x71, 0x19, 0x7d, 0xed, 0xc0, 0xc0, 0xc1, 0x09, 0xbb, 0xa1, 0x3c, 0x25, 0xa7, 0x30, 0x48,
	0xe5, 0xba, 0x34, 0x5c, 0x51, 0xb3, 0xad, 0x78, 0x18, 0x4c, 0x82, 0xe9, 0x30, 0xee, 0x7b, 0x6c,
	0xb5, 0xad, 0x78, 0x5b, 0x52, 0xb2, 0x82, 0x87, 0x9d, 0x49, 0x30, 0xed, 0x35, 0x92, 0x0f, 0xac,
	0xe0, 0xe4, 0x39, 0x5a, 0x4d, 0xe8, 0xb5, 0x50, 0xda, 0x50, 0xbe, 0xe1, 0xa5, 0xa1, 0x46, 0x87,
	0x07, 0x93, 0x60, 0xda, 0x8d, 0xc7, 0x96, 0x7a, 0x63, 0x99, 0x0b, 0x4b, 0xac, 0x34, 0x79, 0x86,
	0xde, 0x12, 0x9a, 0xb3, 0xb6, 0xba, 0x8b, 0xea, 0x91, 0x65, 0xde, 0xb1, 0x9d, 0xf8, 0x15, 0x3c,
	0x44, 0xb1, 0xb0, 0xb7, 0x6d, 0x58, 0xee, 0x07, 0xf0, 0xfa, 0xf0, 0x3f, 0xb4, 0x7b, 0x62, 0x05,
	0x97, 0x9e, 0xc7, 0xb9, 0x73, 0xcb, 0x92, 0x05, 0x20, 0x43, 0x8d, 0x34, 0xbf, 0xcd, 0x1d, 0xe2,
	0x1c, 0x9a, 0x5e, 0x59, 0x72, 0x37, 0x14, 0xfd, 0x0c, 0xe0, 0x81, 0xab, 0xa8, 0x52, 0x32, 0xa3,
	0xeb, 0x2a, 0x63, 0x86, 0x53, 0x6d, 0x98, 0xd1, 0xe4, 0x11, 0xf4, 0x0b, 0x5e, 0x48, 0xb5, 0xa5,
	0x5a, 0xdc, 0xd5, 0x65, 0x81, 0x83, 0x96, 0xe2, 0x0e, 0xbb, 0x92, 0xc9, 0x17, 0x9e, 0xd6, 0xf7,
	0x74, 0x5c, 0x9d, 0x0e, 0x73, 0xa6, 0x9e, 0xc2, 0x11, 0x2f, 0x33, 0x79, 0xbd, 0x0b, 0xd4, 0x34,
	0x35, 0x42, 0xa2, 0x0e, 0xb2, 0xd2, 0xe4, 0x0a, 0xc6, 0xfc, 0xd6, 0xf0, 0x32, 0xe3, 0x19, 0xf5,
	0x7d, 0x87, 0xdd, 0xc9, 0xc1, 0xb4, 0x3f, 0x5f, 0xcc, 0xfe, 0x69, 0x0d, 0x66, 0xed, 0xc7, 0x8e,
	0x47, 0xf5, 0x61, 0xe7, 0xee, 0xac, 0xe8, 0x5b, 0x00, 0xe3, 0x56, 0xd6, 0x26, 0xa4, 0xfd, 0x6f,
	0x9d, 0x72, 0x45, 0x45, 0x86, 0x21, 0x7b, 0x31, 0xd4, 0xd0, 0x65, 0x46, 0x1e, 0xc3, 0xb0, 0x11,
	0xb4, 0x36, 0x62, 0x50, 0x83, 0xb8, 0x12, 0x57, 0x00, 0xf6, 0x38, 0xa1, 0x8d, 0x48, 0x5d, 0xbe,
	0xfe, 0xfc, 0xf5, 0x5e, 0xa6, 0xff, 0xa8, 0x3f, 0x6e, 0x9d, 0x18, 0xfd, 0x08, 0xe0, 0xc8, 0xe9,
	0x4c, 0x92, 0xd7, 0x4b, 0x6e, 0xbd, 0xfb, 0xfe, 0x9b, 0x6d, 0xee, 0xc5, 0xe0, 0xa0, 0x7a, 0x99,
	0xff, 0xf6, 0x40, 0xa7, 0x30, 0x30, 0x2c, 0xc9, 0x39, 0x75, 0xef, 0x8a, 0xde, 0x87, 0x71, 0x1f,
	0xb1, 0xf7, 0x08, 0x91, 0x4f, 0xad, 0x06, 0xac, 0x27, 0xff, 0x28, 0x2f, 0xf6, 0xcf, 0xe7, 0x82,
	0x35, 0xd5, 0x2d, 0x0d, 0x33, 0xd1, 0xf7, 0x00, 0x86, 0xf7, 0xbe, 0x5d, 0xf2, 0x19, 0x86, 0xce,
	0x92, 0x07, 0xc2, 0x39, 0xde, 0xf7, 0x72, 0xaf, 0xfb, 0x5a, 0x3d, 0xc5, 0x2e, 0xe1, 0xd2, 0x1f,
	0x7f, 0x06, 0x27, 0xa9, 0x2c, 0x37, 0x5c, 0xdd, 0xf0, 0x8c, 0x22, 0xa3, 0x7d, 0x3d, 0x0b, 0xcc,
	0x7e, 0xdc, 0xb0, 0x2b, 0x24, 0x77, 0x3d, 0xe1, 0x87, 0xe5, 0x7b, 0x3a, 0xf3, 0x3d, 0x59, 0xcc,
	0xf5, 0x94, 0x1c, 0xe2, 0x4f, 0xd6, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x43, 0xc2,
	0x6b, 0xc9, 0x04, 0x00, 0x00,
}
