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
// source: l2rib_client_detail.proto

package cisco_ios_xr_l2rib_oper_l2rib_clients_details_clients_detail

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

type L2RibClientDetail_KEYS struct {
	ClientId             uint32   `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2RibClientDetail_KEYS) Reset()         { *m = L2RibClientDetail_KEYS{} }
func (m *L2RibClientDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2RibClientDetail_KEYS) ProtoMessage()    {}
func (*L2RibClientDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5dd6bd0cc39409, []int{0}
}

func (m *L2RibClientDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibClientDetail_KEYS.Unmarshal(m, b)
}
func (m *L2RibClientDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibClientDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *L2RibClientDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibClientDetail_KEYS.Merge(m, src)
}
func (m *L2RibClientDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2RibClientDetail_KEYS.Size(m)
}
func (m *L2RibClientDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibClientDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibClientDetail_KEYS proto.InternalMessageInfo

func (m *L2RibClientDetail_KEYS) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type L2RibClient struct {
	ClientIdXr           uint32   `protobuf:"varint,1,opt,name=client_id_xr,json=clientIdXr,proto3" json:"client_id_xr,omitempty"`
	ProcessId            uint32   `protobuf:"varint,2,opt,name=process_id,json=processId,proto3" json:"process_id,omitempty"`
	NodeId               string   `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ProcName             string   `protobuf:"bytes,4,opt,name=proc_name,json=procName,proto3" json:"proc_name,omitempty"`
	ProcSuffix           string   `protobuf:"bytes,5,opt,name=proc_suffix,json=procSuffix,proto3" json:"proc_suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2RibClient) Reset()         { *m = L2RibClient{} }
func (m *L2RibClient) String() string { return proto.CompactTextString(m) }
func (*L2RibClient) ProtoMessage()    {}
func (*L2RibClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5dd6bd0cc39409, []int{1}
}

func (m *L2RibClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibClient.Unmarshal(m, b)
}
func (m *L2RibClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibClient.Marshal(b, m, deterministic)
}
func (m *L2RibClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibClient.Merge(m, src)
}
func (m *L2RibClient) XXX_Size() int {
	return xxx_messageInfo_L2RibClient.Size(m)
}
func (m *L2RibClient) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibClient.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibClient proto.InternalMessageInfo

func (m *L2RibClient) GetClientIdXr() uint32 {
	if m != nil {
		return m.ClientIdXr
	}
	return 0
}

func (m *L2RibClient) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *L2RibClient) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2RibClient) GetProcName() string {
	if m != nil {
		return m.ProcName
	}
	return ""
}

func (m *L2RibClient) GetProcSuffix() string {
	if m != nil {
		return m.ProcSuffix
	}
	return ""
}

type L2RibProdInfo struct {
	ObjectType           string   `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	ProducerId           string   `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	ProducerName         string   `protobuf:"bytes,3,opt,name=producer_name,json=producerName,proto3" json:"producer_name,omitempty"`
	AdminDistance        uint32   `protobuf:"varint,4,opt,name=admin_distance,json=adminDistance,proto3" json:"admin_distance,omitempty"`
	PurgeTime            uint32   `protobuf:"varint,5,opt,name=purge_time,json=purgeTime,proto3" json:"purge_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2RibProdInfo) Reset()         { *m = L2RibProdInfo{} }
func (m *L2RibProdInfo) String() string { return proto.CompactTextString(m) }
func (*L2RibProdInfo) ProtoMessage()    {}
func (*L2RibProdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5dd6bd0cc39409, []int{2}
}

func (m *L2RibProdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibProdInfo.Unmarshal(m, b)
}
func (m *L2RibProdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibProdInfo.Marshal(b, m, deterministic)
}
func (m *L2RibProdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibProdInfo.Merge(m, src)
}
func (m *L2RibProdInfo) XXX_Size() int {
	return xxx_messageInfo_L2RibProdInfo.Size(m)
}
func (m *L2RibProdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibProdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibProdInfo proto.InternalMessageInfo

func (m *L2RibProdInfo) GetObjectType() string {
	if m != nil {
		return m.ObjectType
	}
	return ""
}

func (m *L2RibProdInfo) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *L2RibProdInfo) GetProducerName() string {
	if m != nil {
		return m.ProducerName
	}
	return ""
}

func (m *L2RibProdInfo) GetAdminDistance() uint32 {
	if m != nil {
		return m.AdminDistance
	}
	return 0
}

func (m *L2RibProdInfo) GetPurgeTime() uint32 {
	if m != nil {
		return m.PurgeTime
	}
	return 0
}

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
	return fileDescriptor_6a5dd6bd0cc39409, []int{3}
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
	return fileDescriptor_6a5dd6bd0cc39409, []int{4}
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
	return fileDescriptor_6a5dd6bd0cc39409, []int{5}
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

type L2RibClientDetail struct {
	Client                      *L2RibClient     `protobuf:"bytes,50,opt,name=client,proto3" json:"client,omitempty"`
	ProducerCount               uint32           `protobuf:"varint,51,opt,name=producer_count,json=producerCount,proto3" json:"producer_count,omitempty"`
	ProducerArray               []*L2RibProdInfo `protobuf:"bytes,52,rep,name=producer_array,json=producerArray,proto3" json:"producer_array,omitempty"`
	RegistrationTableStatistics *L2RibProdStats  `protobuf:"bytes,53,opt,name=registration_table_statistics,json=registrationTableStatistics,proto3" json:"registration_table_statistics,omitempty"`
	LastUpdateTimestamp         uint64           `protobuf:"varint,54,opt,name=last_update_timestamp,json=lastUpdateTimestamp,proto3" json:"last_update_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}         `json:"-"`
	XXX_unrecognized            []byte           `json:"-"`
	XXX_sizecache               int32            `json:"-"`
}

func (m *L2RibClientDetail) Reset()         { *m = L2RibClientDetail{} }
func (m *L2RibClientDetail) String() string { return proto.CompactTextString(m) }
func (*L2RibClientDetail) ProtoMessage()    {}
func (*L2RibClientDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a5dd6bd0cc39409, []int{6}
}

func (m *L2RibClientDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2RibClientDetail.Unmarshal(m, b)
}
func (m *L2RibClientDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2RibClientDetail.Marshal(b, m, deterministic)
}
func (m *L2RibClientDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2RibClientDetail.Merge(m, src)
}
func (m *L2RibClientDetail) XXX_Size() int {
	return xxx_messageInfo_L2RibClientDetail.Size(m)
}
func (m *L2RibClientDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_L2RibClientDetail.DiscardUnknown(m)
}

var xxx_messageInfo_L2RibClientDetail proto.InternalMessageInfo

func (m *L2RibClientDetail) GetClient() *L2RibClient {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *L2RibClientDetail) GetProducerCount() uint32 {
	if m != nil {
		return m.ProducerCount
	}
	return 0
}

func (m *L2RibClientDetail) GetProducerArray() []*L2RibProdInfo {
	if m != nil {
		return m.ProducerArray
	}
	return nil
}

func (m *L2RibClientDetail) GetRegistrationTableStatistics() *L2RibProdStats {
	if m != nil {
		return m.RegistrationTableStatistics
	}
	return nil
}

func (m *L2RibClientDetail) GetLastUpdateTimestamp() uint64 {
	if m != nil {
		return m.LastUpdateTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*L2RibClientDetail_KEYS)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_client_detail_KEYS")
	proto.RegisterType((*L2RibClient)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_client")
	proto.RegisterType((*L2RibProdInfo)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_prod_info")
	proto.RegisterType((*L2RibBagEc)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_bag_ec")
	proto.RegisterType((*L2RibProdUpdateStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_prod_update_stats")
	proto.RegisterType((*L2RibProdStats)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_prod_stats")
	proto.RegisterType((*L2RibClientDetail)(nil), "cisco_ios_xr_l2rib_oper.l2rib.clients_details.clients_detail.l2rib_client_detail")
}

func init() { proto.RegisterFile("l2rib_client_detail.proto", fileDescriptor_6a5dd6bd0cc39409) }

var fileDescriptor_6a5dd6bd0cc39409 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0xd6, 0xa4, 0xb9, 0xb9, 0x8d, 0x93, 0xdc, 0xf6, 0xba, 0xba, 0xb7, 0x83, 0xaa, 0x8a, 0x10,
	0x84, 0x54, 0x81, 0xc8, 0x22, 0xe5, 0x47, 0x48, 0x6c, 0x50, 0x29, 0x52, 0xf8, 0xe9, 0x62, 0x32,
	0x95, 0x60, 0x65, 0x79, 0xc6, 0x4e, 0x65, 0x34, 0x33, 0x1e, 0xd9, 0x9e, 0x2a, 0xe9, 0x8a, 0x67,
	0xe0, 0x09, 0x58, 0xf1, 0x18, 0xbc, 0x02, 0x8f, 0xc0, 0xab, 0x20, 0x1f, 0x7b, 0x92, 0x51, 0xa9,
	0xc4, 0x82, 0xec, 0x32, 0xdf, 0x77, 0x8e, 0xcf, 0x77, 0xbe, 0x73, 0xec, 0xa0, 0x5b, 0xd9, 0x44,
	0x89, 0x84, 0xa4, 0x99, 0xe0, 0x85, 0x21, 0x8c, 0x1b, 0x2a, 0xb2, 0x71, 0xa9, 0xa4, 0x91, 0xf8,
	0x79, 0x2a, 0x74, 0x2a, 0x89, 0x90, 0x9a, 0x2c, 0x14, 0x71, 0x71, 0xb2, 0xe4, 0x6a, 0x0c, 0x3f,
	0xc7, 0x2e, 0x45, 0xfb, 0x1c, 0x7d, 0xed, 0x7b, 0xf4, 0x14, 0x85, 0x37, 0x1c, 0x4d, 0xde, 0x9c,
	0x7e, 0x98, 0xe1, 0x03, 0xd4, 0xf5, 0xa8, 0x60, 0x61, 0x30, 0x0c, 0x8e, 0x06, 0xd1, 0xb6, 0x03,
	0xa6, 0x6c, 0xf4, 0x35, 0x40, 0xfd, 0x66, 0x26, 0x1e, 0xa2, 0xfe, 0x2a, 0x9a, 0x2c, 0x94, 0x4f,
	0x40, 0x75, 0xc2, 0x7b, 0x85, 0x0f, 0x11, 0x2a, 0x95, 0x4c, 0xb9, 0xd6, 0xf6, 0xc0, 0x16, 0xf0,
	0x5d, 0x8f, 0x4c, 0x19, 0xde, 0x47, 0x7f, 0x17, 0x92, 0x71, 0xcb, 0x6d, 0x0d, 0x83, 0xa3, 0x6e,
	0xd4, 0xb1, 0x9f, 0x53, 0x66, 0x75, 0xd8, 0x28, 0x52, 0xd0, 0x9c, 0x87, 0x6d, 0xa0, 0xb6, 0x2d,
	0x70, 0x46, 0x73, 0x8e, 0x6f, 0xa3, 0x1e, 0x90, 0xba, 0x9a, 0xcf, 0xc5, 0x22, 0xfc, 0x0b, 0x68,
	0xa8, 0x33, 0x03, 0x64, 0xf4, 0x2d, 0x40, 0x3b, 0x4e, 0x68, 0xa9, 0x24, 0x23, 0xa2, 0x98, 0x4b,
	0x9b, 0x24, 0x93, 0x8f, 0x3c, 0x35, 0xc4, 0x2c, 0x4b, 0x0e, 0x52, 0xbb, 0x11, 0x72, 0x50, 0xbc,
	0x2c, 0xeb, 0x53, 0x59, 0x95, 0x72, 0x55, 0x6b, 0x75, 0xa7, 0x02, 0x34, 0x65, 0xf8, 0x2e, 0x1a,
	0xac, 0x02, 0x40, 0x97, 0x93, 0xdc, 0xaf, 0x41, 0xd0, 0x76, 0x0f, 0xfd, 0x43, 0x59, 0x2e, 0x0a,
	0xc2, 0x84, 0x36, 0xb4, 0x48, 0x9d, 0xfa, 0x41, 0x34, 0x00, 0xf4, 0xa5, 0x07, 0xc1, 0x97, 0x4a,
	0x5d, 0x70, 0x62, 0x44, 0xce, 0xa1, 0x03, 0xeb, 0x8b, 0x45, 0x62, 0x91, 0xf3, 0xd1, 0x97, 0x56,
	0xed, 0x74, 0x42, 0x2f, 0x08, 0x4f, 0xf1, 0x1d, 0xd4, 0x4f, 0x65, 0x55, 0x18, 0xae, 0xd6, 0xf2,
	0x07, 0x51, 0xcf, 0x63, 0xa0, 0xbf, 0x11, 0x02, 0xea, 0x5c, 0x03, 0x75, 0x08, 0x88, 0x7b, 0x88,
	0xf6, 0xb2, 0x89, 0x4a, 0xc8, 0x5c, 0x28, 0x6d, 0x08, 0xbf, 0xb4, 0x93, 0x33, 0x1a, 0xfa, 0x68,
	0x47, 0xbb, 0x96, 0x7a, 0x65, 0x99, 0x53, 0x4b, 0xc4, 0x1a, 0x3f, 0x40, 0x18, 0xc2, 0x33, 0xda,
	0x8c, 0x6e, 0x43, 0xb4, 0xf5, 0x37, 0x79, 0x4b, 0xd7, 0xc1, 0xcf, 0x60, 0x61, 0x13, 0x22, 0x6c,
	0xb5, 0x4b, 0x9a, 0xf9, 0x04, 0x28, 0xef, 0x1b, 0xfc, 0xdf, 0x06, 0x4c, 0x3d, 0x0f, 0x79, 0x27,
	0x96, 0xc5, 0xc7, 0x08, 0x18, 0x62, 0xa4, 0xb9, 0x96, 0xd7, 0x81, 0x3c, 0x10, 0x1d, 0x5b, 0x72,
	0x9d, 0x34, 0xfa, 0xd4, 0x42, 0xfb, 0x8d, 0x19, 0x57, 0x25, 0xa3, 0x86, 0x13, 0x6d, 0xa8, 0xd1,
	0x76, 0x94, 0x39, 0xcf, 0xa5, 0x5a, 0x12, 0x2d, 0xae, 0x6a, 0xb3, 0x90, 0x83, 0x66, 0xe2, 0x0a,
	0xbc, 0xf2, 0xcb, 0xe0, 0xea, 0xb8, 0xc5, 0xf4, 0x0b, 0xe2, 0x44, 0xdd, 0x47, 0xff, 0xf2, 0x82,
	0xc9, 0xf9, 0xba, 0xa1, 0x95, 0x53, 0x3b, 0x40, 0xd4, 0x8d, 0xc4, 0x1a, 0x57, 0x68, 0x97, 0x2f,
	0x0c, 0x2f, 0x18, 0x67, 0xc4, 0xfb, 0x1d, 0xb6, 0x87, 0x5b, 0x47, 0xbd, 0xc9, 0xeb, 0xf1, 0x9f,
	0x5c, 0xd5, 0x71, 0x73, 0x07, 0xa2, 0x9d, 0xba, 0xc6, 0x89, 0x2b, 0x31, 0xfa, 0x1e, 0xa0, 0xdd,
	0x86, 0x05, 0xab, 0xde, 0x9b, 0x6b, 0x1c, 0xfc, 0x7e, 0x8d, 0x5b, 0x37, 0xac, 0x71, 0x85, 0x90,
	0x3d, 0x4e, 0x68, 0x23, 0x52, 0xd7, 0x76, 0x6f, 0x72, 0xbe, 0x89, 0x5e, 0x7e, 0x19, 0x56, 0xd4,
	0x28, 0x34, 0xfa, 0xb1, 0x05, 0x1b, 0x7a, 0xfd, 0x6d, 0xc2, 0x09, 0xea, 0x38, 0x20, 0x9c, 0x80,
	0x94, 0x8d, 0xd8, 0xea, 0xc0, 0xc8, 0x9f, 0x6c, 0x6f, 0xee, 0xca, 0x17, 0xb7, 0x15, 0xc7, 0xee,
	0xe6, 0xd6, 0xa8, 0xdb, 0x0b, 0xd3, 0x08, 0xa3, 0x4a, 0xd1, 0x65, 0xf8, 0x08, 0x26, 0xfd, 0x6e,
	0x63, 0xee, 0xd8, 0xe7, 0x6a, 0x5d, 0xf5, 0x85, 0xad, 0x81, 0x3f, 0x07, 0xe8, 0x50, 0xf1, 0x0b,
	0xa1, 0x8d, 0xa2, 0x46, 0xc8, 0x82, 0x18, 0x9a, 0x64, 0xce, 0x41, 0x3f, 0xa3, 0xc7, 0x60, 0xcc,
	0xd9, 0xc6, 0x54, 0xb8, 0xe1, 0x1c, 0x34, 0x8b, 0xc6, 0xb6, 0xe6, 0x6c, 0x55, 0x12, 0x4f, 0xd0,
	0x7f, 0xf0, 0x34, 0xf8, 0x71, 0xda, 0xa7, 0x4c, 0x1b, 0x9a, 0x97, 0xe1, 0x13, 0xb8, 0x26, 0x7b,
	0x96, 0x3c, 0x07, 0x2e, 0xae, 0xa9, 0xa4, 0x03, 0xff, 0x60, 0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0xb2, 0x7a, 0x56, 0xde, 0x06, 0x00, 0x00,
}