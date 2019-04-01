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

type L2RibDdParams struct {
	DdParamsDisable      bool     `protobuf:"varint,1,opt,name=dd_params_disable,json=ddParamsDisable,proto3" json:"dd_params_disable,omitempty"`
	DdParamsFreezeTime   uint32   `protobuf:"varint,2,opt,name=dd_params_freeze_time,json=ddParamsFreezeTime,proto3" json:"dd_params_freeze_time,omitempty"`
	DdParamsRetryCount   uint32   `protobuf:"varint,3,opt,name=dd_params_retry_count,json=ddParamsRetryCount,proto3" json:"dd_params_retry_count,omitempty"`
	DdParamsMoveCount    uint32   `protobuf:"varint,4,opt,name=dd_params_move_count,json=ddParamsMoveCount,proto3" json:"dd_params_move_count,omitempty"`
	DdParamsMoveInterval uint32   `protobuf:"varint,5,opt,name=dd_params_move_interval,json=ddParamsMoveInterval,proto3" json:"dd_params_move_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2RibDdParams) Reset()         { *m = L2RibDdParams{} }
func (m *L2RibDdParams) String() string { return proto.CompactTextString(m) }
func (*L2RibDdParams) ProtoMessage()    {}
func (*L2RibDdParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{5}
}

func (m *L2RibDdParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibDdParams.Unmarshal(m, b)
}
func (m *L2RibDdParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibDdParams.Marshal(b, m, deterministic)
}
func (m *L2RibDdParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibDdParams.Merge(m, src)
}
func (m *L2RibDdParams) XXX_Size() int {
	return xxx_messageInfo_L2RibDdParams.Size(m)
}
func (m *L2RibDdParams) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibDdParams.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibDdParams proto.InternalMessageInfo

func (m *L2RibDdParams) GetDdParamsDisable() bool {
	if m != nil {
		return m.DdParamsDisable
	}
	return false
}

func (m *L2RibDdParams) GetDdParamsFreezeTime() uint32 {
	if m != nil {
		return m.DdParamsFreezeTime
	}
	return 0
}

func (m *L2RibDdParams) GetDdParamsRetryCount() uint32 {
	if m != nil {
		return m.DdParamsRetryCount
	}
	return 0
}

func (m *L2RibDdParams) GetDdParamsMoveCount() uint32 {
	if m != nil {
		return m.DdParamsMoveCount
	}
	return 0
}

func (m *L2RibDdParams) GetDdParamsMoveInterval() uint32 {
	if m != nil {
		return m.DdParamsMoveInterval
	}
	return 0
}

type L2RibSummary struct {
	TableSummary         []*L2RibTblSummary `protobuf:"bytes,50,rep,name=table_summary,json=tableSummary,proto3" json:"table_summary,omitempty"`
	ConvergedTablesCount uint32             `protobuf:"varint,51,opt,name=converged_tables_count,json=convergedTablesCount,proto3" json:"converged_tables_count,omitempty"`
	TotalMemory          uint32             `protobuf:"varint,52,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	MacDdParams          *L2RibDdParams     `protobuf:"bytes,53,opt,name=mac_dd_params,json=macDdParams,proto3" json:"mac_dd_params,omitempty"`
	Ipv4DdParams         *L2RibDdParams     `protobuf:"bytes,54,opt,name=ipv4_dd_params,json=ipv4DdParams,proto3" json:"ipv4_dd_params,omitempty"`
	Ipv6DdParams         *L2RibDdParams     `protobuf:"bytes,55,opt,name=ipv6_dd_params,json=ipv6DdParams,proto3" json:"ipv6_dd_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2RibSummary) Reset()         { *m = L2RibSummary{} }
func (m *L2RibSummary) String() string { return proto.CompactTextString(m) }
func (*L2RibSummary) ProtoMessage()    {}
func (*L2RibSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3df371c99cc020d1, []int{6}
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

func (m *L2RibSummary) GetMacDdParams() *L2RibDdParams {
	if m != nil {
		return m.MacDdParams
	}
	return nil
}

func (m *L2RibSummary) GetIpv4DdParams() *L2RibDdParams {
	if m != nil {
		return m.Ipv4DdParams
	}
	return nil
}

func (m *L2RibSummary) GetIpv6DdParams() *L2RibDdParams {
	if m != nil {
		return m.Ipv6DdParams
	}
	return nil
}

func init() {
	proto.RegisterType((*L2RibSummary_KEYS)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_summary_KEYS")
	proto.RegisterType((*L2RibBagEc)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_bag_ec")
	proto.RegisterType((*L2RibProdUpdateStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_prod_update_stats")
	proto.RegisterType((*L2RibProdStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_prod_stats")
	proto.RegisterType((*L2RibTblSummary)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_tbl_summary")
	proto.RegisterType((*L2RibDdParams)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_dd_params")
	proto.RegisterType((*L2RibSummary)(nil), "cisco_ios_xr_l2rib_oper.l2rib.summary.l2rib_summary")
}

func init() { proto.RegisterFile("l2rib_summary.proto", fileDescriptor_3df371c99cc020d1) }

var fileDescriptor_3df371c99cc020d1 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4b, 0x6e, 0x13, 0x4d,
	0x10, 0x96, 0x1f, 0x7f, 0xf4, 0xa7, 0x6c, 0xe3, 0x64, 0x08, 0x89, 0x59, 0xe1, 0x18, 0x21, 0x59,
	0x20, 0x8c, 0xb0, 0x13, 0x07, 0x36, 0x6c, 0xf2, 0x90, 0x22, 0x08, 0x42, 0x63, 0x6f, 0x40, 0x21,
	0xad, 0x9e, 0x99, 0x4e, 0xd4, 0xc8, 0xe3, 0x1e, 0x75, 0xb7, 0xad, 0x38, 0x57, 0xe0, 0x12, 0x1c,
	0x83, 0xa3, 0xb0, 0x63, 0xcd, 0x2d, 0x50, 0x57, 0xf7, 0x3c, 0x62, 0x16, 0x60, 0x65, 0x17, 0x7f,
	0x8f, 0x9a, 0xaa, 0xaf, 0x6a, 0x26, 0x70, 0x7f, 0xd2, 0x97, 0x3c, 0x20, 0x6a, 0x16, 0xc7, 0x54,
	0x2e, 0x7a, 0x89, 0x14, 0x5a, 0x78, 0x4f, 0x42, 0xae, 0x42, 0x41, 0xb8, 0x50, 0xe4, 0x5a, 0x12,
	0xab, 0x10, 0x09, 0x93, 0x3d, 0xfc, 0xb3, 0xe7, 0xc4, 0x9d, 0x2d, 0xf0, 0x6e, 0xb9, 0xc9, 0xdb,
	0xe3, 0x8f, 0xa3, 0xce, 0xb7, 0x32, 0xd4, 0x2d, 0x1c, 0xd0, 0x2b, 0xc2, 0x42, 0x6f, 0x17, 0xea,
	0xa1, 0x98, 0x4d, 0x35, 0x93, 0x44, 0x2f, 0x12, 0xd6, 0x2a, 0xb5, 0x4b, 0xdd, 0x86, 0x5f, 0x73,
	0xd8, 0x78, 0x91, 0xb0, 0xa2, 0x64, 0x4a, 0x63, 0xd6, 0x2a, 0xb7, 0x4b, 0xdd, 0xf5, 0x4c, 0xf2,
	0x9e, 0xc6, 0xcc, 0x7b, 0x8e, 0xad, 0x06, 0xe4, 0x92, 0x4b, 0xa5, 0x09, 0x9b, 0xb3, 0xa9, 0x26,
	0x5a, 0xb5, 0x2a, 0xed, 0x52, 0xb7, 0xea, 0x6f, 0x18, 0xea, 0xc4, 0x30, 0xc7, 0x86, 0x18, 0x2b,
	0xef, 0x19, 0xf6, 0x16, 0x90, 0x09, 0x2d, 0xaa, 0xab, 0xa8, 0x6e, 0x1a, 0xe6, 0x1d, 0xcd, 0xc5,
	0xaf, 0xe1, 0x21, 0x8a, 0xb9, 0x79, 0xda, 0x9c, 0x4e, 0x9c, 0x01, 0x1f, 0xdf, 0xfa, 0x0f, 0xdb,
	0xdd, 0x36, 0x82, 0x53, 0xc7, 0xa3, 0xef, 0xd0, 0xb0, 0xde, 0x00, 0x90, 0x21, 0x5a, 0xe8, 0x25,
	0xdf, 0x1a, 0xfa, 0xb0, 0xe9, 0xb1, 0x21, 0x73, 0x53, 0xe7, 0x57, 0x09, 0x76, 0x6c, 0x44, 0x89,
	0x14, 0x11, 0x99, 0x25, 0x11, 0xd5, 0x8c, 0x28, 0x4d, 0xb5, 0xf2, 0x1e, 0x41, 0x2d, 0x66, 0xb1,
	0x90, 0x0b, 0xa2, 0xf8, 0x4d, 0x1a, 0x16, 0x58, 0x68, 0xc4, 0x6f, 0x30, 0x2b, 0x11, 0x7c, 0x61,
	0x61, 0xfa, 0x9c, 0xb2, 0x8d, 0xd3, 0x62, 0xb6, 0xa9, 0xa7, 0xb0, 0xc9, 0xa6, 0x91, 0xb8, 0xcc,
	0x07, 0xca, 0x92, 0x6a, 0x22, 0x91, 0x0e, 0x32, 0x56, 0xde, 0x05, 0x6c, 0xb0, 0x6b, 0xcd, 0xa6,
	0x11, 0x8b, 0x88, 0xcb, 0xbb, 0x55, 0x6d, 0x57, 0xba, 0xb5, 0xfe, 0xa0, 0xf7, 0x4f, 0x67, 0xd0,
	0x2b, 0x2e, 0xdb, 0x6f, 0xa6, 0xc5, 0x0e, 0x6d, 0xad, 0xce, 0xf7, 0x12, 0x6c, 0x14, 0x66, 0xcd,
	0x86, 0x34, 0xbf, 0x66, 0x21, 0x93, 0x84, 0x47, 0x38, 0xe4, 0xba, 0x0f, 0x29, 0x74, 0x1a, 0x79,
	0x8f, 0xa1, 0x91, 0x09, 0x0a, 0x17, 0x51, 0x4f, 0x41, 0x3c, 0x89, 0x0b, 0x00, 0x53, 0x8e, 0x2b,
	0xcd, 0x43, 0x3b, 0x5f, 0xad, 0xff, 0x66, 0xa5, 0xa6, 0xff, 0x88, 0xdf, 0x2f, 0x54, 0xec, 0xfc,
	0x28, 0xc1, 0xa6, 0xd5, 0xe9, 0x60, 0x92, 0x1e, 0xb9, 0xe9, 0xdd, 0xe5, 0x9f, 0x5d, 0xf3, 0xba,
	0x0f, 0x16, 0x4a, 0x8f, 0xf9, 0x6f, 0x0b, 0xda, 0x85, 0xba, 0xa6, 0xc1, 0x84, 0x11, 0xbb, 0x57,
	0xec, 0xbd, 0xe1, 0xd7, 0x10, 0x3b, 0x43, 0xc8, 0x3b, 0x2f, 0x24, 0x60, 0x7a, 0x72, 0x4b, 0x39,
	0x58, 0x7d, 0x3e, 0x3b, 0x58, 0x16, 0xdd, 0x48, 0x53, 0xdd, 0xf9, 0x5a, 0x86, 0xa6, 0x95, 0x44,
	0x11, 0x49, 0xa8, 0xa4, 0xb1, 0x32, 0x57, 0x93, 0xfd, 0x20, 0x11, 0x57, 0xa6, 0x19, 0x1c, 0xef,
	0x7f, 0xbf, 0x19, 0x45, 0x1f, 0x10, 0x3f, 0xb2, 0xb0, 0xf7, 0x12, 0x1e, 0xe4, 0xda, 0x4b, 0xc9,
	0xd8, 0x0d, 0x23, 0x9a, 0xbb, 0x3d, 0x35, 0x7c, 0x2f, 0xd5, 0x9f, 0x20, 0x35, 0xe6, 0xf1, 0x92,
	0x45, 0x32, 0x2d, 0x17, 0x2e, 0x9f, 0xca, 0x6d, 0x8b, 0x6f, 0x28, 0x1b, 0xd3, 0x0b, 0xd8, 0xca,
	0x2d, 0xb1, 0x98, 0x33, 0xe7, 0xa8, 0xa2, 0x63, 0x33, 0x75, 0x9c, 0x89, 0x39, 0xb3, 0x86, 0x7d,
	0xd8, 0x59, 0x32, 0xa4, 0x6f, 0x80, 0x7b, 0x8d, 0xb7, 0x8a, 0x9e, 0xf4, 0x2d, 0xe8, 0xfc, 0xac,
	0x40, 0xe3, 0xd6, 0x97, 0xcc, 0xfb, 0x0c, 0x0d, 0xbb, 0x20, 0x07, 0xb4, 0xfa, 0x98, 0xfe, 0xab,
	0x95, 0xd2, 0x2f, 0x5c, 0x8d, 0x6f, 0xf7, 0x3d, 0x72, 0xe5, 0xf7, 0x60, 0x3b, 0x14, 0xd3, 0x39,
	0x93, 0x57, 0x2c, 0x22, 0xc8, 0x28, 0x37, 0xda, 0xc0, 0xb6, 0x99, 0xb1, 0x63, 0x24, 0xf3, 0xab,
	0xc1, 0xcf, 0x8c, 0xbb, 0x9a, 0x3d, 0x77, 0x35, 0x06, 0x73, 0x57, 0xf3, 0x09, 0x1a, 0x31, 0x0d,
	0xf3, 0xa5, 0xb6, 0xf6, 0xf1, 0xad, 0x18, 0xae, 0xd4, 0x77, 0xe6, 0xf6, 0x6b, 0x31, 0x0d, 0x8f,
	0x5c, 0x62, 0xde, 0x39, 0xdc, 0xe3, 0xc9, 0x7c, 0xaf, 0x50, 0x7c, 0x78, 0xa7, 0xe2, 0x75, 0x53,
	0x6d, 0xa9, 0xfa, 0xb0, 0x50, 0xfd, 0xe0, 0xce, 0xd5, 0x87, 0x69, 0xf5, 0x60, 0x0d, 0xff, 0xb1,
	0x0d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xc1, 0xbe, 0x2e, 0xef, 0x06, 0x00, 0x00,
}
