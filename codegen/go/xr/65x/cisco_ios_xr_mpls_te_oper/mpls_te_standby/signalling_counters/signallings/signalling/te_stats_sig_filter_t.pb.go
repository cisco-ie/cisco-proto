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
// source: te_stats_sig_filter_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_signalling_counters_signallings_signalling

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

type TeStatsSigFilterT_KEYS struct {
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

func (m *TeStatsSigFilterT_KEYS) Reset()         { *m = TeStatsSigFilterT_KEYS{} }
func (m *TeStatsSigFilterT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterT_KEYS) ProtoMessage()    {}
func (*TeStatsSigFilterT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{0}
}

func (m *TeStatsSigFilterT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsSigFilterT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Marshal(b, m, deterministic)
}
func (m *TeStatsSigFilterT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterT_KEYS.Merge(m, src)
}
func (m *TeStatsSigFilterT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterT_KEYS.Size(m)
}
func (m *TeStatsSigFilterT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterT_KEYS proto.InternalMessageInfo

func (m *TeStatsSigFilterT_KEYS) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *TeStatsSigFilterT_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *TeStatsSigFilterT_KEYS) GetSubGroupId() uint32 {
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
	return fileDescriptor_adee6c75523ff5ba, []int{1}
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
	return fileDescriptor_adee6c75523ff5ba, []int{2}
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

type TeStatsBagDestT struct {
	DestinationAddress   string            `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Statistics           *TeStatsSigT      `protobuf:"bytes,2,opt,name=statistics,proto3" json:"statistics,omitempty"`
	S2LStatistic         []*TeStatsBagS2LT `protobuf:"bytes,3,rep,name=s2l_statistic,json=s2lStatistic,proto3" json:"s2l_statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagDestT) Reset()         { *m = TeStatsBagDestT{} }
func (m *TeStatsBagDestT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagDestT) ProtoMessage()    {}
func (*TeStatsBagDestT) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{3}
}

func (m *TeStatsBagDestT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagDestT.Unmarshal(m, b)
}
func (m *TeStatsBagDestT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagDestT.Marshal(b, m, deterministic)
}
func (m *TeStatsBagDestT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagDestT.Merge(m, src)
}
func (m *TeStatsBagDestT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagDestT.Size(m)
}
func (m *TeStatsBagDestT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagDestT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagDestT proto.InternalMessageInfo

func (m *TeStatsBagDestT) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *TeStatsBagDestT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagDestT) GetS2LStatistic() []*TeStatsBagS2LT {
	if m != nil {
		return m.S2LStatistic
	}
	return nil
}

type TeStatsBagVifT struct {
	TunnelName           string             `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TunnelSigName        string             `protobuf:"bytes,2,opt,name=tunnel_sig_name,json=tunnelSigName,proto3" json:"tunnel_sig_name,omitempty"`
	LspId                uint32             `protobuf:"varint,3,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	Statistics           *TeStatsSigT       `protobuf:"bytes,4,opt,name=statistics,proto3" json:"statistics,omitempty"`
	DestinationStatistic []*TeStatsBagDestT `protobuf:"bytes,5,rep,name=destination_statistic,json=destinationStatistic,proto3" json:"destination_statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TeStatsBagVifT) Reset()         { *m = TeStatsBagVifT{} }
func (m *TeStatsBagVifT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagVifT) ProtoMessage()    {}
func (*TeStatsBagVifT) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{4}
}

func (m *TeStatsBagVifT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsBagVifT.Unmarshal(m, b)
}
func (m *TeStatsBagVifT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsBagVifT.Marshal(b, m, deterministic)
}
func (m *TeStatsBagVifT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsBagVifT.Merge(m, src)
}
func (m *TeStatsBagVifT) XXX_Size() int {
	return xxx_messageInfo_TeStatsBagVifT.Size(m)
}
func (m *TeStatsBagVifT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsBagVifT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsBagVifT proto.InternalMessageInfo

func (m *TeStatsBagVifT) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *TeStatsBagVifT) GetTunnelSigName() string {
	if m != nil {
		return m.TunnelSigName
	}
	return ""
}

func (m *TeStatsBagVifT) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *TeStatsBagVifT) GetStatistics() *TeStatsSigT {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *TeStatsBagVifT) GetDestinationStatistic() []*TeStatsBagDestT {
	if m != nil {
		return m.DestinationStatistic
	}
	return nil
}

type TeStatsBagLspT struct {
	TunnelName           string            `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	Statistics           *TeStatsSigT      `protobuf:"bytes,2,opt,name=statistics,proto3" json:"statistics,omitempty"`
	S2LStatistic         []*TeStatsBagS2LT `protobuf:"bytes,3,rep,name=s2l_statistic,json=s2lStatistic,proto3" json:"s2l_statistic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TeStatsBagLspT) Reset()         { *m = TeStatsBagLspT{} }
func (m *TeStatsBagLspT) String() string { return proto.CompactTextString(m) }
func (*TeStatsBagLspT) ProtoMessage()    {}
func (*TeStatsBagLspT) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{5}
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

type TeStatsSigFilterDataU struct {
	StatisticsFilter      string          `protobuf:"bytes,1,opt,name=statistics_filter,json=statisticsFilter,proto3" json:"statistics_filter,omitempty"`
	TeSignallingFilterVif *TeStatsBagVifT `protobuf:"bytes,2,opt,name=te_signalling_filter_vif,json=teSignallingFilterVif,proto3" json:"te_signalling_filter_vif,omitempty"`
	TeSignallingFilterLsp *TeStatsBagLspT `protobuf:"bytes,3,opt,name=te_signalling_filter_lsp,json=teSignallingFilterLsp,proto3" json:"te_signalling_filter_lsp,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}        `json:"-"`
	XXX_unrecognized      []byte          `json:"-"`
	XXX_sizecache         int32           `json:"-"`
}

func (m *TeStatsSigFilterDataU) Reset()         { *m = TeStatsSigFilterDataU{} }
func (m *TeStatsSigFilterDataU) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterDataU) ProtoMessage()    {}
func (*TeStatsSigFilterDataU) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{6}
}

func (m *TeStatsSigFilterDataU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterDataU.Unmarshal(m, b)
}
func (m *TeStatsSigFilterDataU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterDataU.Marshal(b, m, deterministic)
}
func (m *TeStatsSigFilterDataU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterDataU.Merge(m, src)
}
func (m *TeStatsSigFilterDataU) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterDataU.Size(m)
}
func (m *TeStatsSigFilterDataU) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterDataU.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterDataU proto.InternalMessageInfo

func (m *TeStatsSigFilterDataU) GetStatisticsFilter() string {
	if m != nil {
		return m.StatisticsFilter
	}
	return ""
}

func (m *TeStatsSigFilterDataU) GetTeSignallingFilterVif() *TeStatsBagVifT {
	if m != nil {
		return m.TeSignallingFilterVif
	}
	return nil
}

func (m *TeStatsSigFilterDataU) GetTeSignallingFilterLsp() *TeStatsBagLspT {
	if m != nil {
		return m.TeSignallingFilterLsp
	}
	return nil
}

type TeStatsSigFilterT struct {
	TeSignallingFilterData *TeStatsSigFilterDataU `protobuf:"bytes,50,opt,name=te_signalling_filter_data,json=teSignallingFilterData,proto3" json:"te_signalling_filter_data,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}               `json:"-"`
	XXX_unrecognized       []byte                 `json:"-"`
	XXX_sizecache          int32                  `json:"-"`
}

func (m *TeStatsSigFilterT) Reset()         { *m = TeStatsSigFilterT{} }
func (m *TeStatsSigFilterT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigFilterT) ProtoMessage()    {}
func (*TeStatsSigFilterT) Descriptor() ([]byte, []int) {
	return fileDescriptor_adee6c75523ff5ba, []int{7}
}

func (m *TeStatsSigFilterT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigFilterT.Unmarshal(m, b)
}
func (m *TeStatsSigFilterT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigFilterT.Marshal(b, m, deterministic)
}
func (m *TeStatsSigFilterT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigFilterT.Merge(m, src)
}
func (m *TeStatsSigFilterT) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigFilterT.Size(m)
}
func (m *TeStatsSigFilterT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigFilterT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigFilterT proto.InternalMessageInfo

func (m *TeStatsSigFilterT) GetTeSignallingFilterData() *TeStatsSigFilterDataU {
	if m != nil {
		return m.TeSignallingFilterData
	}
	return nil
}

func init() {
	proto.RegisterType((*TeStatsSigFilterT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_t_KEYS")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_t")
	proto.RegisterType((*TeStatsBagS2LT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_s2l_t")
	proto.RegisterType((*TeStatsBagDestT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_dest_t")
	proto.RegisterType((*TeStatsBagVifT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_vif_t")
	proto.RegisterType((*TeStatsBagLspT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_bag_lsp_t")
	proto.RegisterType((*TeStatsSigFilterDataU)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_data_u")
	proto.RegisterType((*TeStatsSigFilterT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.signalling_counters.signallings.signalling.te_stats_sig_filter_t")
}

func init() { proto.RegisterFile("te_stats_sig_filter_t.proto", fileDescriptor_adee6c75523ff5ba) }

var fileDescriptor_adee6c75523ff5ba = []byte{
	// 962 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x96, 0x93, 0x6d, 0xda, 0xbc, 0x34, 0xdd, 0x74, 0x92, 0x74, 0xcd, 0xee, 0x81, 0x12, 0x09,
	0x14, 0x09, 0xc8, 0xa2, 0xec, 0x8d, 0xdb, 0xb2, 0x5b, 0x50, 0x05, 0xe2, 0x87, 0x5b, 0x90, 0x38,
	0x8d, 0x26, 0xf1, 0x34, 0x1d, 0xad, 0x6b, 0x9b, 0x99, 0xe7, 0x90, 0xde, 0xb9, 0x72, 0xe0, 0x8e,
	0x10, 0x5c, 0xb8, 0xf0, 0x77, 0x70, 0xe3, 0x1f, 0xe0, 0xc0, 0xff, 0x82, 0x66, 0xc6, 0x8e, 0xed,
	0xd8, 0x91, 0x7a, 0x69, 0xe0, 0xb0, 0xb7, 0xe6, 0xbd, 0xef, 0x9b, 0x79, 0xdf, 0xfb, 0x5e, 0x5e,
	0x27, 0xf0, 0x04, 0x39, 0x55, 0xc8, 0x50, 0x51, 0x25, 0x16, 0xf4, 0x4a, 0x04, 0xc8, 0x25, 0xc5,
	0x49, 0x2c, 0x23, 0x8c, 0xc8, 0xe5, 0x5c, 0xa8, 0x79, 0x44, 0x45, 0xa4, 0xe8, 0x4a, 0xd2, 0x9b,
	0x38, 0x50, 0x14, 0x39, 0x8d, 0x62, 0x2e, 0x27, 0xd9, 0x07, 0x85, 0x2c, 0xf4, 0x67, 0xb7, 0x13,
	0x25, 0x16, 0x21, 0x0b, 0x02, 0x11, 0x2e, 0xe8, 0x3c, 0x4a, 0x42, 0xe4, 0x52, 0x15, 0x62, 0xc5,
	0xbf, 0x47, 0x7f, 0x37, 0xe0, 0x71, 0xed, 0xad, 0xf4, 0xd3, 0xb3, 0x6f, 0x2f, 0xc8, 0x00, 0xf6,
	0xe6, 0x78, 0x1b, 0x73, 0xd7, 0x39, 0x75, 0xc6, 0x6d, 0xcf, 0x7e, 0x20, 0x4f, 0xa0, 0x8d, 0x49,
	0x18, 0xf2, 0x80, 0x0a, 0xdf, 0x6d, 0x9c, 0x3a, 0xe3, 0xae, 0x77, 0x60, 0x03, 0xe7, 0x3e, 0x79,
	0x0f, 0x08, 0x5f, 0x21, 0x0f, 0x7d, 0xee, 0xd3, 0x1c, 0xd5, 0x34, 0xfc, 0x5e, 0x96, 0xb9, 0xcc,
	0xd0, 0x8f, 0x60, 0x3f, 0x9e, 0xde, 0xc4, 0x1a, 0xf2, 0xc0, 0x1c, 0xd4, 0xd2, 0x1f, 0xcf, 0x7d,
	0x32, 0x84, 0x56, 0xa0, 0x4c, 0x7c, 0xcf, 0xc4, 0xf7, 0x02, 0xa5, 0xc3, 0x6f, 0xc3, 0x91, 0x8a,
	0x12, 0x39, 0xe7, 0x94, 0xf9, 0xbe, 0xe4, 0x4a, 0xb9, 0x2d, 0x73, 0x72, 0xd7, 0x46, 0x9f, 0xdb,
	0x20, 0x79, 0x0a, 0x7d, 0x9f, 0x2b, 0x14, 0x21, 0x43, 0x11, 0x85, 0x6b, 0xec, 0xbe, 0xc1, 0x92,
	0x42, 0x2a, 0x23, 0x7c, 0x00, 0x03, 0x95, 0xcc, 0xe8, 0x42, 0x46, 0x49, 0x4c, 0x23, 0x29, 0x16,
	0x3a, 0x1f, 0x49, 0xf7, 0xc0, 0x32, 0x54, 0x32, 0xfb, 0x44, 0xa7, 0xbe, 0x58, 0x67, 0xc8, 0x29,
	0x1c, 0xe6, 0x0c, 0xe1, 0xbb, 0x6d, 0x53, 0x26, 0x64, 0xc8, 0x73, 0x7f, 0xf4, 0x7b, 0x1b, 0x8e,
	0x4a, 0xbd, 0x45, 0x32, 0x86, 0x1e, 0xae, 0x28, 0x5f, 0xf2, 0x10, 0x69, 0x12, 0xbe, 0x0a, 0xa3,
	0xef, 0x43, 0xd3, 0xda, 0xae, 0x77, 0x84, 0xab, 0x33, 0x1d, 0xfe, 0xda, 0x46, 0xc9, 0x53, 0x18,
	0xe0, 0x8a, 0xc6, 0x0c, 0xaf, 0xe9, 0x5c, 0x72, 0x86, 0xdc, 0xb2, 0xd2, 0x76, 0x1f, 0xe3, 0xea,
	0x4b, 0x86, 0xd7, 0x2f, 0x4c, 0xc6, 0xf0, 0x4a, 0x84, 0x6b, 0x16, 0x2e, 0x32, 0x42, 0xb3, 0x44,
	0x30, 0x99, 0x0a, 0xc1, 0xe7, 0x01, 0x5f, 0xdf, 0xf0, 0xa0, 0x48, 0x78, 0x69, 0x32, 0x96, 0xf0,
	0x3e, 0xf4, 0x33, 0x02, 0x97, 0x32, 0x92, 0x29, 0xde, 0xfa, 0xd3, 0xb3, 0xf8, 0x33, 0x9d, 0x28,
	0x9e, 0x2f, 0xb9, 0x5a, 0x96, 0x15, 0xb4, 0xb2, 0xf3, 0x3d, 0xae, 0x96, 0x55, 0x05, 0x96, 0x50,
	0x54, 0xb0, 0x5f, 0x22, 0x54, 0x14, 0x18, 0x42, 0x49, 0xc1, 0x41, 0x91, 0x50, 0x55, 0x60, 0x08,
	0x45, 0x05, 0xed, 0x4c, 0x81, 0xc6, 0x17, 0x14, 0x7c, 0x08, 0x8f, 0x33, 0xc1, 0x92, 0xf3, 0x25,
	0x0b, 0xe8, 0x77, 0x09, 0x97, 0xb7, 0x29, 0x0b, 0x0c, 0xeb, 0xc4, 0xea, 0xf6, 0x4c, 0xfe, 0x2b,
	0x9d, 0xb6, 0xdc, 0x31, 0xf4, 0xe4, 0xa6, 0xd3, 0x1d, 0xeb, 0xb4, 0xac, 0x38, 0x2d, 0xeb, 0x9c,
	0x3e, 0xb4, 0x2a, 0x64, 0x9d, 0xd3, 0xb2, 0xce, 0xe9, 0x6e, 0x89, 0x50, 0xee, 0x93, 0xac, 0x73,
	0xfa, 0xa8, 0x48, 0xd8, 0xe8, 0x93, 0xac, 0x71, 0xfa, 0xa1, 0xed, 0x93, 0xac, 0x71, 0x5a, 0xd6,
	0x39, 0xdd, 0xcb, 0xce, 0xaf, 0x71, 0x5a, 0xd6, 0x39, 0x7d, 0x5c, 0x22, 0x54, 0x14, 0x54, 0x9d,
	0x26, 0x45, 0x42, 0x55, 0x41, 0xc5, 0xe9, 0x7e, 0xa6, 0xa0, 0xea, 0xb4, 0xdc, 0xee, 0xf4, 0xc0,
	0x3a, 0x2d, 0xeb, 0x9d, 0x7e, 0x06, 0x27, 0xb8, 0xa2, 0x33, 0x36, 0x7f, 0x95, 0xc4, 0x94, 0x29,
	0xbd, 0x5c, 0x53, 0xde, 0xd0, 0xf0, 0xfa, 0xb8, 0xfa, 0xc8, 0x24, 0x9f, 0x9b, 0x5c, 0xf1, 0xc2,
	0x0d, 0x92, 0xcc, 0xca, 0x3c, 0xc9, 0x2e, 0x2c, 0x11, 0x65, 0x5a, 0xac, 0xde, 0xb0, 0xfa, 0x0f,
	0x45, 0x31, 0x42, 0x16, 0xd8, 0x1d, 0xef, 0x3e, 0xb2, 0xd2, 0x6c, 0xe6, 0x52, 0x27, 0x5e, 0xe8,
	0x38, 0x79, 0x0b, 0x0e, 0x53, 0xb4, 0xc5, 0xb9, 0x06, 0xd7, 0xb1, 0x31, 0x03, 0x19, 0xfd, 0xda,
	0x00, 0xb2, 0x5e, 0x54, 0x33, 0xb6, 0xa0, 0x6a, 0x1a, 0x50, 0xdc, 0xba, 0x13, 0x9d, 0x3b, 0xef,
	0xc4, 0xc6, 0xe6, 0x4e, 0xdc, 0xb6, 0x98, 0x9b, 0x5b, 0x17, 0xf3, 0x0f, 0x0e, 0x80, 0x2e, 0x4c,
	0x28, 0x14, 0x73, 0x65, 0x96, 0x53, 0x67, 0xea, 0x4f, 0xee, 0xe3, 0x9f, 0xe1, 0xa4, 0xbc, 0xac,
	0xbd, 0xc2, 0xbd, 0xa3, 0x7f, 0x1a, 0xd0, 0x2f, 0xb5, 0x48, 0x97, 0x4a, 0x71, 0x9b, 0x1e, 0xe7,
	0xae, 0x7a, 0x1a, 0xff, 0x8d, 0x1e, 0xf2, 0xa3, 0x03, 0x5d, 0xed, 0xf2, 0x3a, 0xe4, 0x36, 0x4f,
	0x9b, 0xe3, 0xce, 0xf4, 0xfa, 0x9e, 0x2b, 0x59, 0x4f, 0x97, 0x77, 0xa8, 0xa6, 0xc1, 0x45, 0x76,
	0xfb, 0xe8, 0xe7, 0xe6, 0xc6, 0x08, 0x2e, 0xc5, 0x15, 0x45, 0xf2, 0x26, 0x74, 0xd2, 0x37, 0x44,
	0xc8, 0x6e, 0xb2, 0x57, 0x08, 0xd8, 0xd0, 0xe7, 0xec, 0x86, 0x93, 0x77, 0xe0, 0x61, 0x0a, 0xd0,
	0x1a, 0x0d, 0xa8, 0x61, 0x1f, 0x04, 0x36, 0x7c, 0x21, 0x16, 0x06, 0x97, 0x3f, 0x27, 0x9a, 0xc5,
	0xe7, 0xc4, 0xff, 0x63, 0xba, 0xc8, 0x2f, 0x0e, 0x0c, 0x8b, 0x63, 0x94, 0xbb, 0xb2, 0x67, 0x5c,
	0x11, 0x3b, 0x70, 0xc5, 0x0e, 0xb4, 0x37, 0x28, 0xd4, 0x91, 0xdb, 0xf3, 0xd7, 0xe6, 0x86, 0xd0,
	0xcd, 0xbc, 0x83, 0x3d, 0xaf, 0xa7, 0xbd, 0x76, 0xda, 0x7f, 0x6a, 0xd6, 0xbf, 0xba, 0x7d, 0x86,
	0x8c, 0x26, 0xe4, 0x5d, 0x38, 0xce, 0x8b, 0x4f, 0x73, 0x69, 0x73, 0x7b, 0x79, 0xe2, 0x63, 0x13,
	0x27, 0xbf, 0x39, 0xe0, 0xea, 0xb3, 0xf2, 0xf2, 0xd2, 0xc3, 0x96, 0xe2, 0x2a, 0x6d, 0xf8, 0x2e,
	0x64, 0x9a, 0xef, 0xab, 0x37, 0x44, 0x7e, 0xb1, 0x06, 0xd8, 0x02, 0xbf, 0x11, 0x57, 0xdb, 0x6b,
	0x0c, 0x54, 0x6c, 0xbe, 0x90, 0xbb, 0xa9, 0xd1, 0x0c, 0x6d, 0x5d, 0x8d, 0x9f, 0xa9, 0x78, 0xf4,
	0xa7, 0x03, 0xc3, 0xda, 0x5f, 0x42, 0xe4, 0x0f, 0x07, 0xde, 0xa8, 0xad, 0x5e, 0xdb, 0xe5, 0x4e,
	0x4d, 0xf9, 0xf1, 0x0e, 0x66, 0xba, 0x34, 0x24, 0xde, 0x49, 0x55, 0xc6, 0x4b, 0x86, 0x6c, 0xd6,
	0x32, 0x3f, 0x17, 0x9f, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xdb, 0x8e, 0x42, 0x4d, 0x0e,
	0x00, 0x00,
}
