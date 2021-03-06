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
// source: te_stats_bag_lsp_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_signalling_counters_remote_signalling_counters_remote_signalling_counter

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

type TeStatsBagLspT_KEYS struct {
	Ctype                string   `protobuf:"bytes,1,opt,name=ctype,proto3" json:"ctype,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	P2MpId               uint32   `protobuf:"varint,4,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	LspId                uint32   `protobuf:"varint,5,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,6,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,7,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	SubGroupOriginator   string   `protobuf:"bytes,8,opt,name=sub_group_originator,json=subGroupOriginator,proto3" json:"sub_group_originator,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,9,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsBagLspT_KEYS) Reset()         { *m = TeStatsBagLspT_KEYS{} }
func (m *TeStatsBagLspT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT_KEYS) ProtoMessage()    {}
func (*TeStatsBagLspT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2a17ab9d5569e7, []int{0}
}

func (m *TeStatsBagLspT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsBagLspT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Marshal(b, m, deterministic)
}
func (m *TeStatsBagLspT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagLspT_KEYS.Merge(m, src)
}
func (m *TeStatsBagLspT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagLspT_KEYS.Size(m)
}
func (m *TeStatsBagLspT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagLspT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagLspT_KEYS proto.InternalMessageInfo

func (m *TeStatsBagLspT_KEYS) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *TeStatsBagLspT_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *TeStatsBagLspT_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *TeStatsBagLspT_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsBagLspT_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

type TeStatsSigT struct {
	TxEventUnknown         uint32   `protobuf:"varint,1,opt,name=tx_event_unknown,json=txEventUnknown,proto3" json:"tx_event_unknown,omitempty"`
	TxPathCreateEvent      uint32   `protobuf:"varint,2,opt,name=tx_path_create_event,json=txPathCreateEvent,proto3" json:"tx_path_create_event,omitempty"`
	TxPathChangeEvent      uint32   `protobuf:"varint,3,opt,name=tx_path_change_event,json=txPathChangeEvent,proto3" json:"tx_path_change_event,omitempty"`
	TxPathDeleteEvent      uint32   `protobuf:"varint,4,opt,name=tx_path_delete_event,json=txPathDeleteEvent,proto3" json:"tx_path_delete_event,omitempty"`
	TxPathErrorEvent       uint32   `protobuf:"varint,5,opt,name=tx_path_error_event,json=txPathErrorEvent,proto3" json:"tx_path_error_event,omitempty"`
	TxResvCreateEvent      uint32   `protobuf:"varint,6,opt,name=tx_resv_create_event,json=txResvCreateEvent,proto3" json:"tx_resv_create_event,omitempty"`
	TxResvChangeEvent      uint32   `protobuf:"varint,7,opt,name=tx_resv_change_event,json=txResvChangeEvent,proto3" json:"tx_resv_change_event,omitempty"`
	TxResvDeleteEvent      uint32   `protobuf:"varint,8,opt,name=tx_resv_delete_event,json=txResvDeleteEvent,proto3" json:"tx_resv_delete_event,omitempty"`
	TxResvErrorEvent       uint32   `protobuf:"varint,9,opt,name=tx_resv_error_event,json=txResvErrorEvent,proto3" json:"tx_resv_error_event,omitempty"`
	TxPathReevalQueryEvent uint32   `protobuf:"varint,10,opt,name=tx_path_reeval_query_event,json=txPathReevalQueryEvent,proto3" json:"tx_path_reeval_query_event,omitempty"`
	RxEventUnknown         uint32   `protobuf:"varint,11,opt,name=rx_event_unknown,json=rxEventUnknown,proto3" json:"rx_event_unknown,omitempty"`
	RxPathCreateEvent      uint32   `protobuf:"varint,12,opt,name=rx_path_create_event,json=rxPathCreateEvent,proto3" json:"rx_path_create_event,omitempty"`
	RxPathChangeEvent      uint32   `protobuf:"varint,13,opt,name=rx_path_change_event,json=rxPathChangeEvent,proto3" json:"rx_path_change_event,omitempty"`
	RxPathDeleteEvent      uint32   `protobuf:"varint,14,opt,name=rx_path_delete_event,json=rxPathDeleteEvent,proto3" json:"rx_path_delete_event,omitempty"`
	RxPathErrorEvent       uint32   `protobuf:"varint,15,opt,name=rx_path_error_event,json=rxPathErrorEvent,proto3" json:"rx_path_error_event,omitempty"`
	RxResvCreateEvent      uint32   `protobuf:"varint,16,opt,name=rx_resv_create_event,json=rxResvCreateEvent,proto3" json:"rx_resv_create_event,omitempty"`
	RxResvChangeEvent      uint32   `protobuf:"varint,17,opt,name=rx_resv_change_event,json=rxResvChangeEvent,proto3" json:"rx_resv_change_event,omitempty"`
	RxResvDeleteEvent      uint32   `protobuf:"varint,18,opt,name=rx_resv_delete_event,json=rxResvDeleteEvent,proto3" json:"rx_resv_delete_event,omitempty"`
	RxResvErrorEvent       uint32   `protobuf:"varint,19,opt,name=rx_resv_error_event,json=rxResvErrorEvent,proto3" json:"rx_resv_error_event,omitempty"`
	RxPathReevalQueryEvent uint32   `protobuf:"varint,20,opt,name=rx_path_reeval_query_event,json=rxPathReevalQueryEvent,proto3" json:"rx_path_reeval_query_event,omitempty"`
	TxBackupAssignEvent    uint32   `protobuf:"varint,21,opt,name=tx_backup_assign_event,json=txBackupAssignEvent,proto3" json:"tx_backup_assign_event,omitempty"`
	RxBackupAssignErrEvent uint32   `protobuf:"varint,22,opt,name=rx_backup_assign_err_event,json=rxBackupAssignErrEvent,proto3" json:"rx_backup_assign_err_event,omitempty"`
	EventsTotalCount       uint32   `protobuf:"varint,23,opt,name=events_total_count,json=eventsTotalCount,proto3" json:"events_total_count,omitempty"`
	EventsCount            uint32   `protobuf:"varint,24,opt,name=events_count,json=eventsCount,proto3" json:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TeStatsSigT) Reset()         { *m = TeStatsSigT{} }
func (m *TeStatsSigT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT) ProtoMessage()    {}
func (*TeStatsSigT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2a17ab9d5569e7, []int{1}
}

func (m *TeStatsSigT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigT.Unmarshal(m, b)
}
func (m *TeStatsSigT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigT.Marshal(b, m, deterministic)
}
func (m *TeStatsSigT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigT.Merge(m, src)
}
func (m *TeStatsSigT) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigT.Size(m)
}
func (m *TeStatsSigT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigT proto.InternalMessageInfo

func (m *TeStatsSigT) GetTxEventUnknown() uint32 {
	if m != nil {
		return m.TxEventUnknown
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathCreateEvent() uint32 {
	if m != nil {
		return m.TxPathCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathChangeEvent() uint32 {
	if m != nil {
		return m.TxPathChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathDeleteEvent() uint32 {
	if m != nil {
		return m.TxPathDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathErrorEvent() uint32 {
	if m != nil {
		return m.TxPathErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvCreateEvent() uint32 {
	if m != nil {
		return m.TxResvCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvChangeEvent() uint32 {
	if m != nil {
		return m.TxResvChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvDeleteEvent() uint32 {
	if m != nil {
		return m.TxResvDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxResvErrorEvent() uint32 {
	if m != nil {
		return m.TxResvErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxPathReevalQueryEvent() uint32 {
	if m != nil {
		return m.TxPathReevalQueryEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxEventUnknown() uint32 {
	if m != nil {
		return m.RxEventUnknown
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathCreateEvent() uint32 {
	if m != nil {
		return m.RxPathCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathChangeEvent() uint32 {
	if m != nil {
		return m.RxPathChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathDeleteEvent() uint32 {
	if m != nil {
		return m.RxPathDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathErrorEvent() uint32 {
	if m != nil {
		return m.RxPathErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvCreateEvent() uint32 {
	if m != nil {
		return m.RxResvCreateEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvChangeEvent() uint32 {
	if m != nil {
		return m.RxResvChangeEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvDeleteEvent() uint32 {
	if m != nil {
		return m.RxResvDeleteEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxResvErrorEvent() uint32 {
	if m != nil {
		return m.RxResvErrorEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxPathReevalQueryEvent() uint32 {
	if m != nil {
		return m.RxPathReevalQueryEvent
	}
	return 0
}

func (m *TeStatsSigT) GetTxBackupAssignEvent() uint32 {
	if m != nil {
		return m.TxBackupAssignEvent
	}
	return 0
}

func (m *TeStatsSigT) GetRxBackupAssignErrEvent() uint32 {
	if m != nil {
		return m.RxBackupAssignErrEvent
	}
	return 0
}

func (m *TeStatsSigT) GetEventsTotalCount() uint32 {
	if m != nil {
		return m.EventsTotalCount
	}
	return 0
}

func (m *TeStatsSigT) GetEventsCount() uint32 {
	if m != nil {
		return m.EventsCount
	}
	return 0
}

type TeStatsBagS2LT struct {
	SubGroupOriginator   string       `protobuf:"bytes,1,opt,name=sub_group_originator,json=subGroupOriginator,proto3" json:"sub_group_originator,omitempty"`
	SubGroupId           uint32       `protobuf:"varint,2,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	DestinationAddress   string       `protobuf:"bytes,3,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Statistics           *TeStatsSigT `protobuf:"bytes,4,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TeStatsBagS2LT) Reset()         { *m = TeStatsBagS2LT{} }
func (m *TeStatsBagS2LT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagS2LT) ProtoMessage()    {}
func (*TeStatsBagS2LT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2a17ab9d5569e7, []int{2}
}

func (m *TeStatsBagS2LT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagS2LT.Unmarshal(m, b)
}
func (m *TeStatsBagS2LT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagS2LT.Marshal(b, m, deterministic)
}
func (m *TeStatsBagS2LT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagS2LT.Merge(m, src)
}
func (m *TeStatsBagS2LT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagS2LT.Size(m)
}
func (m *TeStatsBagS2LT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagS2LT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagS2LT proto.InternalMessageInfo

func (m *TeStatsBagS2LT) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsBagS2LT) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *TeStatsBagS2LT) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagS2LT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

type TeStatsBagLspT struct {
	TunnelName           string            `protobuf:"bytes,50,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	Statistics           *TeStatsSigT      `protobuf:"bytes,51,opt,name=statistics,proto3" json:"statistics,omitempty"`
	S2LStatistic         []*TeStatsBagS2LT `protobuf:"bytes,52,rep,name=s2l_statistic,json=s2lStatistic,proto3" json:"s2l_statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagLspT) Reset()         { *m = TeStatsBagLspT{} }
func (m *TeStatsBagLspT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT) ProtoMessage()    {}
func (*TeStatsBagLspT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2a17ab9d5569e7, []int{3}
}

func (m *TeStatsBagLspT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagLspT.Unmarshal(m, b)
}
func (m *TeStatsBagLspT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagLspT.Marshal(b, m, deterministic)
}
func (m *TeStatsBagLspT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagLspT.Merge(m, src)
}
func (m *TeStatsBagLspT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagLspT.Size(m)
}
func (m *TeStatsBagLspT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagLspT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagLspT proto.InternalMessageInfo

func (m *TeStatsBagLspT) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *TeStatsBagLspT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagLspT) GetS2LStatistic() []*TeStatsBagS2LT {
	if m != nil {
		return m.S2LStatistic
	}
	return nil
}

func init() {
	proto.RegisterType((*TeStatsBagLspT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_lsp_t_KEYS")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_s2l_t")
	proto.RegisterType((*TeStatsBagLspT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.remote_signalling_counters.remote_signalling_counter.te_stats_bag_lsp_t")
}

func init() { proto.RegisterFile("te_stats_bag_lsp_t.proto", fileDescriptor_0a2a17ab9d5569e7) }

var fileDescriptor_0a2a17ab9d5569e7 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcf, 0x6e, 0xd3, 0x4a,
	0x14, 0xc6, 0x95, 0xa4, 0x4d, 0x9b, 0xc9, 0x9f, 0x9b, 0x4e, 0xd2, 0xd6, 0xba, 0x77, 0x71, 0x43,
	0x24, 0xa4, 0x2c, 0x20, 0x45, 0x29, 0x2b, 0x76, 0xa5, 0x54, 0xa8, 0x42, 0xe2, 0x4f, 0x5a, 0x16,
	0xac, 0x46, 0x13, 0x7b, 0x94, 0x9a, 0x3a, 0x1e, 0x73, 0x66, 0x1c, 0x5c, 0x16, 0x3c, 0x08, 0x5b,
	0x24, 0x1e, 0x81, 0x47, 0xe1, 0x59, 0x58, 0xa2, 0x99, 0xf1, 0xa4, 0x76, 0xec, 0x48, 0xb0, 0x01,
	0x76, 0xf5, 0xf9, 0xce, 0x6f, 0x7c, 0xbe, 0xf9, 0xa2, 0xe3, 0x22, 0x47, 0x32, 0x22, 0x24, 0x95,
	0x82, 0xcc, 0xe8, 0x9c, 0x04, 0x22, 0x22, 0x72, 0x1c, 0x01, 0x97, 0x1c, 0xbf, 0x75, 0x7d, 0xe1,
	0x72, 0xe2, 0x73, 0x41, 0x12, 0x20, 0x8b, 0x28, 0x10, 0x44, 0x32, 0xc2, 0x23, 0x06, 0xe3, 0xf4,
	0x61, 0x2c, 0xfc, 0x79, 0x48, 0x83, 0xc0, 0x0f, 0xe7, 0xc4, 0xe5, 0x71, 0x28, 0x19, 0x88, 0x31,
	0xb0, 0x05, 0x57, 0x67, 0xfe, 0x8a, 0x34, 0xfc, 0x56, 0x45, 0x87, 0xc5, 0x41, 0xc8, 0xb3, 0xb3,
	0x37, 0x17, 0xb8, 0x8f, 0xb6, 0x5d, 0x79, 0x13, 0x31, 0xa7, 0x32, 0xa8, 0x8c, 0x1a, 0x53, 0xf3,
	0x80, 0xff, 0x43, 0x0d, 0x19, 0x87, 0x21, 0x0b, 0x88, 0xef, 0x39, 0xd5, 0x41, 0x65, 0xd4, 0x9e,
	0xee, 0x9a, 0xc2, 0xb9, 0x87, 0xef, 0x21, 0xcc, 0x12, 0xc9, 0x42, 0x8f, 0x79, 0xe4, 0xb6, 0xab,
	0xa6, 0xf9, 0xae, 0x55, 0x2e, 0x6d, 0xf7, 0x21, 0xda, 0x89, 0x26, 0x8b, 0x48, 0xb5, 0x6c, 0xe9,
	0x83, 0xea, 0xea, 0xf1, 0xdc, 0xc3, 0xfb, 0xa8, 0xae, 0xe6, 0xf0, 0x3d, 0x67, 0x5b, 0xd7, 0xb7,
	0x03, 0xa1, 0xca, 0x77, 0x51, 0x47, 0xf0, 0x18, 0x5c, 0x46, 0xa8, 0xe7, 0x01, 0x13, 0xc2, 0xa9,
	0xeb, 0x93, 0xdb, 0xa6, 0x7a, 0x62, 0x8a, 0xf8, 0x08, 0xf5, 0x3c, 0x26, 0xa4, 0x1f, 0x52, 0xe9,
	0xf3, 0x70, 0xd5, 0xbb, 0xa3, 0x7b, 0x71, 0x46, 0xb2, 0xc0, 0x03, 0xd4, 0x17, 0xf1, 0x8c, 0xcc,
	0x81, 0xc7, 0x11, 0xe1, 0xe0, 0xcf, 0x95, 0xce, 0xc1, 0xd9, 0x35, 0x84, 0x88, 0x67, 0x4f, 0x95,
	0xf4, 0x62, 0xa5, 0xe0, 0x01, 0x6a, 0xdd, 0x12, 0xbe, 0xe7, 0x34, 0xf4, 0x98, 0xc8, 0x76, 0x9e,
	0x7b, 0xc3, 0x2f, 0x0d, 0xd4, 0x59, 0x5d, 0xac, 0xf0, 0xe7, 0x44, 0xe2, 0x11, 0xea, 0xca, 0x84,
	0xb0, 0x25, 0x0b, 0x25, 0x89, 0xc3, 0xeb, 0x90, 0xbf, 0x0f, 0xf5, 0xd5, 0xb6, 0xa7, 0x1d, 0x99,
	0x9c, 0xa9, 0xf2, 0x6b, 0x53, 0xc5, 0x47, 0xa8, 0x2f, 0x13, 0x12, 0x51, 0x79, 0x45, 0x5c, 0x60,
	0x54, 0x32, 0x43, 0xa5, 0xd7, 0xbd, 0x27, 0x93, 0x97, 0x54, 0x5e, 0x9d, 0x6a, 0x45, 0x73, 0x39,
	0xe0, 0x8a, 0x86, 0x73, 0x0b, 0xd4, 0x72, 0x80, 0x56, 0x0a, 0x80, 0xc7, 0x02, 0xb6, 0x7a, 0xc3,
	0x56, 0x16, 0x78, 0xa2, 0x15, 0x03, 0xdc, 0x47, 0x3d, 0x0b, 0x30, 0x00, 0x0e, 0x69, 0xbf, 0xc9,
	0xa7, 0x6b, 0xfa, 0xcf, 0x94, 0x90, 0x3d, 0x1f, 0x98, 0x58, 0xe6, 0x1d, 0xd4, 0xed, 0xf9, 0x53,
	0x26, 0x96, 0x45, 0x07, 0x06, 0xc8, 0x3a, 0xd8, 0xc9, 0x01, 0x05, 0x07, 0x1a, 0xc8, 0x39, 0xd8,
	0xcd, 0x02, 0x45, 0x07, 0x1a, 0xc8, 0x3a, 0x68, 0x58, 0x07, 0xaa, 0x3f, 0xe3, 0xe0, 0x11, 0xfa,
	0xd7, 0x1a, 0x06, 0xc6, 0x96, 0x34, 0x20, 0xef, 0x62, 0x06, 0x37, 0x29, 0x85, 0x34, 0x75, 0x60,
	0x7c, 0x4f, 0xb5, 0xfe, 0x4a, 0xc9, 0x86, 0x1d, 0xa1, 0x2e, 0xac, 0x27, 0xdd, 0x34, 0x49, 0x43,
	0x21, 0x69, 0x28, 0x4b, 0xba, 0x65, 0x5c, 0x40, 0x59, 0xd2, 0x50, 0x96, 0x74, 0x3b, 0x07, 0xe4,
	0xef, 0x09, 0xca, 0x92, 0xee, 0x64, 0x81, 0xb5, 0x7b, 0x82, 0x92, 0xa4, 0xff, 0x31, 0xf7, 0x04,
	0x25, 0x49, 0x43, 0x59, 0xd2, 0x5d, 0x7b, 0x7e, 0x49, 0xd2, 0x50, 0x96, 0xf4, 0x5e, 0x0e, 0x28,
	0x38, 0x28, 0x26, 0x8d, 0xb3, 0x40, 0xd1, 0x41, 0x21, 0xe9, 0x9e, 0x75, 0x50, 0x4c, 0x1a, 0x36,
	0x27, 0xdd, 0x37, 0x49, 0x43, 0x79, 0xd2, 0xc7, 0xe8, 0x40, 0x26, 0x64, 0x46, 0xdd, 0xeb, 0x38,
	0x22, 0x54, 0xa8, 0x0d, 0x9b, 0x72, 0xfb, 0x9a, 0xeb, 0xc9, 0xe4, 0xb1, 0x16, 0x4f, 0xb4, 0x96,
	0x7d, 0xe1, 0x1a, 0x04, 0x76, 0xcc, 0x03, 0xfb, 0xc2, 0x1c, 0x08, 0xe9, 0xb0, 0x6a, 0xc3, 0xaa,
	0x3f, 0x04, 0x91, 0x5c, 0xd2, 0xc0, 0x2c, 0x72, 0xe7, 0xd0, 0x58, 0x33, 0xca, 0xa5, 0x12, 0x4e,
	0x55, 0x1d, 0xdf, 0x41, 0xad, 0xb4, 0xdb, 0xf4, 0x39, 0xba, 0xaf, 0x69, 0x6a, 0xba, 0x65, 0xf8,
	0xb5, 0x8a, 0x70, 0xee, 0x0b, 0x20, 0x26, 0x01, 0x91, 0x1b, 0x77, 0x62, 0xe5, 0xa7, 0x77, 0x62,
	0x75, 0x7d, 0x27, 0x6e, 0x5a, 0xcc, 0xb5, 0x8d, 0x8b, 0xf9, 0x53, 0x05, 0x21, 0x35, 0x98, 0x2f,
	0xa4, 0xef, 0x0a, 0xbd, 0x9c, 0x9a, 0x93, 0x0f, 0xe3, 0xdf, 0xf7, 0x7d, 0x1c, 0xe7, 0x57, 0xf8,
	0x34, 0x33, 0xcd, 0xf0, 0xfb, 0xfa, 0xc5, 0xe9, 0x4f, 0x27, 0xfe, 0x1f, 0x35, 0xd3, 0x2f, 0x5f,
	0x48, 0x17, 0xcc, 0x99, 0x68, 0x73, 0xc8, 0x94, 0x9e, 0xd3, 0x05, 0x5b, 0x37, 0x75, 0xfc, 0x37,
	0x99, 0xc2, 0x9f, 0x2b, 0xa8, 0xad, 0x7e, 0x00, 0xab, 0x92, 0xf3, 0x70, 0x50, 0x1b, 0x35, 0x27,
	0x1f, 0xff, 0xc8, 0x7c, 0xab, 0x9f, 0xe3, 0xb4, 0x25, 0x26, 0xc1, 0x85, 0x9d, 0x69, 0x56, 0xd7,
	0xff, 0x28, 0x1d, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xeb, 0x9e, 0xf6, 0x44, 0x09, 0x00,
	0x00,
}
