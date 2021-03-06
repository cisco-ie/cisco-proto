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
// source: l2tpv2_cc_acct_stats_data.proto

package cisco_ios_xr_tunnel_l2tun_oper_l2tpv2_tunnel_accounting_statistics

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

type L2Tpv2CcAcctStatsData_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2Tpv2CcAcctStatsData_KEYS) Reset()         { *m = L2Tpv2CcAcctStatsData_KEYS{} }
func (m *L2Tpv2CcAcctStatsData_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2Tpv2CcAcctStatsData_KEYS) ProtoMessage()    {}
func (*L2Tpv2CcAcctStatsData_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec43c69f1d20d03, []int{0}
}

func (m *L2Tpv2CcAcctStatsData_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS.Unmarshal(m, b)
}
func (m *L2Tpv2CcAcctStatsData_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS.Marshal(b, m, deterministic)
}
func (m *L2Tpv2CcAcctStatsData_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS.Merge(m, src)
}
func (m *L2Tpv2CcAcctStatsData_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS.Size(m)
}
func (m *L2Tpv2CcAcctStatsData_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2Tpv2CcAcctStatsData_KEYS proto.InternalMessageInfo

type L2Tpv2CcAcctStatsData struct {
	RecordsSentSuccessfully        uint64   `protobuf:"varint,50,opt,name=records_sent_successfully,json=recordsSentSuccessfully,proto3" json:"records_sent_successfully,omitempty"`
	Start                          uint64   `protobuf:"varint,51,opt,name=start,proto3" json:"start,omitempty"`
	Stop                           uint64   `protobuf:"varint,52,opt,name=stop,proto3" json:"stop,omitempty"`
	Reject                         uint64   `protobuf:"varint,53,opt,name=reject,proto3" json:"reject,omitempty"`
	TransportFailures              uint64   `protobuf:"varint,54,opt,name=transport_failures,json=transportFailures,proto3" json:"transport_failures,omitempty"`
	PositiveAcknowledgement        uint64   `protobuf:"varint,55,opt,name=positive_acknowledgement,json=positiveAcknowledgement,proto3" json:"positive_acknowledgement,omitempty"`
	NegativeAcknowledgement        uint64   `protobuf:"varint,56,opt,name=negative_acknowledgement,json=negativeAcknowledgement,proto3" json:"negative_acknowledgement,omitempty"`
	RecordsCheckpointed            uint64   `protobuf:"varint,57,opt,name=records_checkpointed,json=recordsCheckpointed,proto3" json:"records_checkpointed,omitempty"`
	RecordsFailedToCheckpoint      uint64   `protobuf:"varint,58,opt,name=records_failed_to_checkpoint,json=recordsFailedToCheckpoint,proto3" json:"records_failed_to_checkpoint,omitempty"`
	RecordsSentFromQueue           uint64   `protobuf:"varint,59,opt,name=records_sent_from_queue,json=recordsSentFromQueue,proto3" json:"records_sent_from_queue,omitempty"`
	MemoryFailures                 uint32   `protobuf:"varint,60,opt,name=memory_failures,json=memoryFailures,proto3" json:"memory_failures,omitempty"`
	CurrentSize                    uint32   `protobuf:"varint,61,opt,name=current_size,json=currentSize,proto3" json:"current_size,omitempty"`
	RecordsRecoveredFromCheckpoint uint32   `protobuf:"varint,62,opt,name=records_recovered_from_checkpoint,json=recordsRecoveredFromCheckpoint,proto3" json:"records_recovered_from_checkpoint,omitempty"`
	RecordsFailToRecover           uint32   `protobuf:"varint,63,opt,name=records_fail_to_recover,json=recordsFailToRecover,proto3" json:"records_fail_to_recover,omitempty"`
	QueueStatisticsSize            int32    `protobuf:"zigzag32,64,opt,name=queue_statistics_size,json=queueStatisticsSize,proto3" json:"queue_statistics_size,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *L2Tpv2CcAcctStatsData) Reset()         { *m = L2Tpv2CcAcctStatsData{} }
func (m *L2Tpv2CcAcctStatsData) String() string { return proto.CompactTextString(m) }
func (*L2Tpv2CcAcctStatsData) ProtoMessage()    {}
func (*L2Tpv2CcAcctStatsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ec43c69f1d20d03, []int{1}
}

func (m *L2Tpv2CcAcctStatsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData.Unmarshal(m, b)
}
func (m *L2Tpv2CcAcctStatsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData.Marshal(b, m, deterministic)
}
func (m *L2Tpv2CcAcctStatsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2Tpv2CcAcctStatsData.Merge(m, src)
}
func (m *L2Tpv2CcAcctStatsData) XXX_Size() int {
	return xxx_messageInfo_L2Tpv2CcAcctStatsData.Size(m)
}
func (m *L2Tpv2CcAcctStatsData) XXX_DiscardUnknown() {
	xxx_messageInfo_L2Tpv2CcAcctStatsData.DiscardUnknown(m)
}

var xxx_messageInfo_L2Tpv2CcAcctStatsData proto.InternalMessageInfo

func (m *L2Tpv2CcAcctStatsData) GetRecordsSentSuccessfully() uint64 {
	if m != nil {
		return m.RecordsSentSuccessfully
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetStart() uint64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetStop() uint64 {
	if m != nil {
		return m.Stop
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetReject() uint64 {
	if m != nil {
		return m.Reject
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetTransportFailures() uint64 {
	if m != nil {
		return m.TransportFailures
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetPositiveAcknowledgement() uint64 {
	if m != nil {
		return m.PositiveAcknowledgement
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetNegativeAcknowledgement() uint64 {
	if m != nil {
		return m.NegativeAcknowledgement
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetRecordsCheckpointed() uint64 {
	if m != nil {
		return m.RecordsCheckpointed
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetRecordsFailedToCheckpoint() uint64 {
	if m != nil {
		return m.RecordsFailedToCheckpoint
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetRecordsSentFromQueue() uint64 {
	if m != nil {
		return m.RecordsSentFromQueue
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetMemoryFailures() uint32 {
	if m != nil {
		return m.MemoryFailures
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetCurrentSize() uint32 {
	if m != nil {
		return m.CurrentSize
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetRecordsRecoveredFromCheckpoint() uint32 {
	if m != nil {
		return m.RecordsRecoveredFromCheckpoint
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetRecordsFailToRecover() uint32 {
	if m != nil {
		return m.RecordsFailToRecover
	}
	return 0
}

func (m *L2Tpv2CcAcctStatsData) GetQueueStatisticsSize() int32 {
	if m != nil {
		return m.QueueStatisticsSize
	}
	return 0
}

func init() {
	proto.RegisterType((*L2Tpv2CcAcctStatsData_KEYS)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.tunnel.accounting.statistics.l2tpv2_cc_acct_stats_data_KEYS")
	proto.RegisterType((*L2Tpv2CcAcctStatsData)(nil), "cisco_ios_xr_tunnel_l2tun_oper.l2tpv2.tunnel.accounting.statistics.l2tpv2_cc_acct_stats_data")
}

func init() { proto.RegisterFile("l2tpv2_cc_acct_stats_data.proto", fileDescriptor_9ec43c69f1d20d03) }

var fileDescriptor_9ec43c69f1d20d03 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdb, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x55, 0x69, 0xdb, 0x83, 0xb9, 0x69, 0x5e, 0x61, 0x9e, 0x84, 0x46, 0xb7, 0x17, 0xfa,
	0x42, 0x24, 0x32, 0x06, 0x6c, 0x5c, 0xc6, 0x45, 0x54, 0x42, 0x3c, 0xd1, 0xec, 0x85, 0x27, 0xcb,
	0x38, 0xa7, 0xc5, 0x2c, 0xf1, 0x09, 0xf6, 0x49, 0x61, 0xfb, 0x97, 0xf9, 0x27, 0x90, 0x9d, 0xa4,
	0xcd, 0xd0, 0xfa, 0xd4, 0xfa, 0xfc, 0xbe, 0xef, 0xe4, 0x3b, 0xc7, 0x09, 0x7b, 0x54, 0xa4, 0x54,
	0x2d, 0x52, 0xa9, 0xb5, 0x54, 0x5a, 0x93, 0xf4, 0xa4, 0xc8, 0xcb, 0x5c, 0x91, 0x4a, 0x2a, 0x87,
	0x84, 0xfc, 0x83, 0x36, 0x5e, 0xa3, 0x34, 0xe8, 0xe5, 0x1f, 0x27, 0xa9, 0xb6, 0x16, 0x0a, 0x59,
	0xa4, 0x54, 0x5b, 0x89, 0x15, 0xb8, 0xa4, 0xf1, 0x27, 0x0d, 0x48, 0x94, 0xd6, 0x58, 0x5b, 0x32,
	0x76, 0x9e, 0x84, 0x4e, 0xc6, 0x93, 0xd1, 0xfe, 0x70, 0xc4, 0xf6, 0xd7, 0x3e, 0x46, 0x7e, 0xf9,
	0xf4, 0x2d, 0x3b, 0xfc, 0xbb, 0xc9, 0xf6, 0xd6, 0x4a, 0xf8, 0x29, 0xdb, 0x73, 0xa0, 0xd1, 0xe5,
	0x5e, 0x7a, 0xb0, 0x24, 0x7d, 0xad, 0x35, 0x78, 0x3f, 0xab, 0x8b, 0xe2, 0x52, 0xa4, 0xa3, 0xc1,
	0x78, 0x63, 0xba, 0xdb, 0x0a, 0x32, 0xb0, 0x94, 0xf5, 0x30, 0x1f, 0xb2, 0x4d, 0x4f, 0xca, 0x91,
	0x38, 0x8a, 0xba, 0xe6, 0xc0, 0x39, 0xdb, 0xf0, 0x84, 0x95, 0x78, 0x16, 0x8b, 0xf1, 0x3f, 0x7f,
	0xc0, 0xb6, 0x1c, 0xfc, 0x04, 0x4d, 0xe2, 0x38, 0x56, 0xdb, 0x13, 0x7f, 0xc2, 0x38, 0x39, 0x65,
	0x7d, 0x85, 0x8e, 0xe4, 0x4c, 0x99, 0xa2, 0x76, 0xe0, 0xc5, 0xf3, 0xa8, 0xd9, 0x5e, 0x92, 0x49,
	0x0b, 0xf8, 0x09, 0x13, 0x15, 0x7a, 0x43, 0x66, 0x01, 0x52, 0xe9, 0x0b, 0x8b, 0xbf, 0x0b, 0xc8,
	0xe7, 0x50, 0x82, 0x25, 0xf1, 0xa2, 0xc9, 0xda, 0xf1, 0xf7, 0xd7, 0x71, 0xb0, 0x5a, 0x98, 0xab,
	0x1b, 0xad, 0x2f, 0x1b, 0x6b, 0xc7, 0xff, 0xb7, 0x3e, 0x65, 0xc3, 0x6e, 0x45, 0xfa, 0x07, 0xe8,
	0x8b, 0x0a, 0x8d, 0x25, 0xc8, 0xc5, 0x49, 0xb4, 0xed, 0xb4, 0xec, 0x63, 0x0f, 0xf1, 0x33, 0xf6,
	0xb0, 0xb3, 0x84, 0xa9, 0x20, 0x97, 0x84, 0x3d, 0xb3, 0x38, 0x8d, 0xd6, 0x6e, 0xf3, 0x93, 0x28,
	0x39, 0xc7, 0x55, 0x0b, 0x7e, 0xcc, 0x76, 0xaf, 0x5d, 0xcb, 0xcc, 0x61, 0x29, 0x7f, 0xd5, 0x50,
	0x83, 0x78, 0x15, 0xbd, 0xc3, 0xde, 0xa5, 0x4c, 0x1c, 0x96, 0x5f, 0x03, 0xe3, 0x8f, 0xd9, 0xbd,
	0x12, 0x4a, 0x74, 0x97, 0xab, 0x65, 0xbe, 0x1e, 0x0d, 0xc6, 0x77, 0xa6, 0x77, 0x9b, 0xf2, 0x72,
	0x93, 0x07, 0xec, 0xb6, 0xae, 0x9d, 0x8b, 0x37, 0x6e, 0xae, 0x40, 0xbc, 0x89, 0xaa, 0x5b, 0x6d,
	0x2d, 0x33, 0x57, 0xc0, 0x3f, 0xb3, 0x83, 0x2e, 0x42, 0xf8, 0x5d, 0x80, 0x83, 0xbc, 0xc9, 0xd1,
	0x1b, 0xe4, 0x6d, 0xf4, 0xed, 0xb7, 0xc2, 0x69, 0xa7, 0x0b, 0x89, 0x6e, 0x9e, 0x26, 0xe4, 0x0a,
	0xcb, 0x68, 0x5b, 0x8a, 0xb3, 0xd8, 0x60, 0xd8, 0xdb, 0xc4, 0x39, 0xb6, 0x6d, 0x78, 0xca, 0xee,
	0xc7, 0x91, 0xe5, 0xea, 0x7d, 0x6f, 0xd2, 0xbe, 0x1b, 0x0d, 0xc6, 0xdb, 0xd3, 0x9d, 0x08, 0xb3,
	0x25, 0x0b, 0xa9, 0xbf, 0x6f, 0xc5, 0x4f, 0xeb, 0xe8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3,
	0x53, 0xb3, 0xb9, 0x7d, 0x03, 0x00, 0x00,
}
