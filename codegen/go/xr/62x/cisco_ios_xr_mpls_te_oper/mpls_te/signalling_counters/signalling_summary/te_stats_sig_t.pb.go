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
// source: te_stats_sig_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_signalling_counters_signalling_summary

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

type TeStatsSigT_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsSigT_KEYS) Reset()         { *m = TeStatsSigT_KEYS{} }
func (m *TeStatsSigT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT_KEYS) ProtoMessage()    {}
func (*TeStatsSigT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e068ece89e51ea00, []int{0}
}

func (m *TeStatsSigT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsSigT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsSigT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsSigT_KEYS.Marshal(b, m, deterministic)
}
func (m *TeStatsSigT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsSigT_KEYS.Merge(m, src)
}
func (m *TeStatsSigT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsSigT_KEYS.Size(m)
}
func (m *TeStatsSigT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsSigT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsSigT_KEYS proto.InternalMessageInfo

type TeStatsSigT struct {
	TxEventUnknown         uint32   `protobuf:"varint,50,opt,name=tx_event_unknown,json=txEventUnknown,proto3" json:"tx_event_unknown,omitempty"`
	TxPathCreateEvent      uint32   `protobuf:"varint,51,opt,name=tx_path_create_event,json=txPathCreateEvent,proto3" json:"tx_path_create_event,omitempty"`
	TxPathChangeEvent      uint32   `protobuf:"varint,52,opt,name=tx_path_change_event,json=txPathChangeEvent,proto3" json:"tx_path_change_event,omitempty"`
	TxPathDeleteEvent      uint32   `protobuf:"varint,53,opt,name=tx_path_delete_event,json=txPathDeleteEvent,proto3" json:"tx_path_delete_event,omitempty"`
	TxPathErrorEvent       uint32   `protobuf:"varint,54,opt,name=tx_path_error_event,json=txPathErrorEvent,proto3" json:"tx_path_error_event,omitempty"`
	TxResvCreateEvent      uint32   `protobuf:"varint,55,opt,name=tx_resv_create_event,json=txResvCreateEvent,proto3" json:"tx_resv_create_event,omitempty"`
	TxResvChangeEvent      uint32   `protobuf:"varint,56,opt,name=tx_resv_change_event,json=txResvChangeEvent,proto3" json:"tx_resv_change_event,omitempty"`
	TxResvDeleteEvent      uint32   `protobuf:"varint,57,opt,name=tx_resv_delete_event,json=txResvDeleteEvent,proto3" json:"tx_resv_delete_event,omitempty"`
	TxResvErrorEvent       uint32   `protobuf:"varint,58,opt,name=tx_resv_error_event,json=txResvErrorEvent,proto3" json:"tx_resv_error_event,omitempty"`
	TxPathReevalQueryEvent uint32   `protobuf:"varint,59,opt,name=tx_path_reeval_query_event,json=txPathReevalQueryEvent,proto3" json:"tx_path_reeval_query_event,omitempty"`
	RxEventUnknown         uint32   `protobuf:"varint,60,opt,name=rx_event_unknown,json=rxEventUnknown,proto3" json:"rx_event_unknown,omitempty"`
	RxPathCreateEvent      uint32   `protobuf:"varint,61,opt,name=rx_path_create_event,json=rxPathCreateEvent,proto3" json:"rx_path_create_event,omitempty"`
	RxPathChangeEvent      uint32   `protobuf:"varint,62,opt,name=rx_path_change_event,json=rxPathChangeEvent,proto3" json:"rx_path_change_event,omitempty"`
	RxPathDeleteEvent      uint32   `protobuf:"varint,63,opt,name=rx_path_delete_event,json=rxPathDeleteEvent,proto3" json:"rx_path_delete_event,omitempty"`
	RxPathErrorEvent       uint32   `protobuf:"varint,64,opt,name=rx_path_error_event,json=rxPathErrorEvent,proto3" json:"rx_path_error_event,omitempty"`
	RxResvCreateEvent      uint32   `protobuf:"varint,65,opt,name=rx_resv_create_event,json=rxResvCreateEvent,proto3" json:"rx_resv_create_event,omitempty"`
	RxResvChangeEvent      uint32   `protobuf:"varint,66,opt,name=rx_resv_change_event,json=rxResvChangeEvent,proto3" json:"rx_resv_change_event,omitempty"`
	RxResvDeleteEvent      uint32   `protobuf:"varint,67,opt,name=rx_resv_delete_event,json=rxResvDeleteEvent,proto3" json:"rx_resv_delete_event,omitempty"`
	RxResvErrorEvent       uint32   `protobuf:"varint,68,opt,name=rx_resv_error_event,json=rxResvErrorEvent,proto3" json:"rx_resv_error_event,omitempty"`
	RxPathReevalQueryEvent uint32   `protobuf:"varint,69,opt,name=rx_path_reeval_query_event,json=rxPathReevalQueryEvent,proto3" json:"rx_path_reeval_query_event,omitempty"`
	TxBackupAssignEvent    uint32   `protobuf:"varint,70,opt,name=tx_backup_assign_event,json=txBackupAssignEvent,proto3" json:"tx_backup_assign_event,omitempty"`
	RxBackupAssignErrEvent uint32   `protobuf:"varint,71,opt,name=rx_backup_assign_err_event,json=rxBackupAssignErrEvent,proto3" json:"rx_backup_assign_err_event,omitempty"`
	EventsTotalCount       uint32   `protobuf:"varint,72,opt,name=events_total_count,json=eventsTotalCount,proto3" json:"events_total_count,omitempty"`
	EventsCount            uint32   `protobuf:"varint,73,opt,name=events_count,json=eventsCount,proto3" json:"events_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TeStatsSigT) Reset()         { *m = TeStatsSigT{} }
func (m *TeStatsSigT) String() string { return proto.CompactTextString(m) }
func (*TeStatsSigT) ProtoMessage()    {}
func (*TeStatsSigT) Descriptor() ([]byte, []int) {
	return fileDescriptor_e068ece89e51ea00, []int{1}
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

func init() {
	proto.RegisterType((*TeStatsSigT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.signalling_summary.te_stats_sig_t_KEYS")
	proto.RegisterType((*TeStatsSigT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.signalling_counters.signalling_summary.te_stats_sig_t")
}

func init() { proto.RegisterFile("te_stats_sig_t.proto", fileDescriptor_e068ece89e51ea00) }

var fileDescriptor_e068ece89e51ea00 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0xc5, 0x05, 0x09, 0x03, 0xab, 0x92, 0x2e, 0xab, 0x8a, 0x13, 0xec, 0xa9, 0x07, 0x28,
	0x12, 0xe5, 0xef, 0xf2, 0x77, 0xb7, 0x5b, 0x58, 0xc4, 0x05, 0x0a, 0x1c, 0x38, 0x8d, 0xbc, 0xc1,
	0x6a, 0xa3, 0x4d, 0xe3, 0x30, 0x9e, 0x84, 0xec, 0x17, 0xe2, 0x73, 0xa2, 0x8c, 0x63, 0x64, 0xd7,
	0xe6, 0x98, 0x79, 0xef, 0xe7, 0xc9, 0xf3, 0x93, 0x2c, 0xf6, 0x49, 0x81, 0x21, 0x49, 0x06, 0x4c,
	0xb1, 0x06, 0x9a, 0xd5, 0xa8, 0x49, 0x67, 0x67, 0x79, 0x61, 0x72, 0x0d, 0x85, 0x36, 0xd0, 0x21,
	0x6c, 0xeb, 0xd2, 0x00, 0x29, 0xd0, 0xb5, 0xc2, 0xd9, 0xf0, 0x31, 0x33, 0xc5, 0xba, 0x92, 0x65,
	0x59, 0x54, 0x6b, 0xc8, 0x75, 0x53, 0x91, 0x42, 0xe3, 0xcf, 0x4c, 0xb3, 0xdd, 0x4a, 0xbc, 0x3c,
	0xbc, 0x2d, 0xc6, 0xe1, 0x06, 0xf8, 0xb4, 0xfc, 0xf1, 0xf5, 0xf0, 0xcf, 0x35, 0xb1, 0x17, 0xce,
	0xb3, 0xa9, 0x18, 0x51, 0x07, 0xaa, 0x55, 0x15, 0x41, 0x53, 0x5d, 0x54, 0xfa, 0x77, 0x35, 0x79,
	0x74, 0xf7, 0xca, 0xf4, 0xe6, 0x6a, 0x8f, 0xba, 0x65, 0x3f, 0xfe, 0x6e, 0xa7, 0xd9, 0x43, 0xb1,
	0x4f, 0x1d, 0xd4, 0x92, 0x36, 0x90, 0xa3, 0x92, 0xa4, 0x2c, 0x35, 0x99, 0xb3, 0xfb, 0x16, 0x75,
	0x9f, 0x25, 0x6d, 0x16, 0xac, 0x30, 0x17, 0x00, 0x1b, 0x59, 0xad, 0x1d, 0xf0, 0x38, 0x00, 0x58,
	0x89, 0x80, 0x9f, 0xaa, 0x54, 0xff, 0x36, 0x3c, 0xf1, 0x81, 0x53, 0x56, 0x2c, 0xf0, 0x40, 0x8c,
	0x1d, 0xa0, 0x10, 0x35, 0x0e, 0xfe, 0xa7, 0xec, 0x1f, 0x59, 0xff, 0xb2, 0x17, 0xfc, 0xf3, 0x51,
	0x99, 0x36, 0x4c, 0xf0, 0xcc, 0x9d, 0xbf, 0x52, 0xa6, 0x8d, 0x13, 0x58, 0xc0, 0x4f, 0xf0, 0x3c,
	0x00, 0xa2, 0x04, 0x0c, 0x04, 0x09, 0x5e, 0xf8, 0x40, 0x9c, 0x80, 0x01, 0x3f, 0xc1, 0x91, 0x4b,
	0xd0, 0xfb, 0xbd, 0x04, 0x47, 0xe2, 0x8e, 0x0b, 0x8c, 0x4a, 0xb5, 0xb2, 0x84, 0x5f, 0x8d, 0xc2,
	0xcb, 0x81, 0x7a, 0xc9, 0xd4, 0x81, 0xcd, 0xbd, 0x62, 0xfd, 0x4b, 0x2f, 0x5b, 0x76, 0x2a, 0x46,
	0xb8, 0xdb, 0xf4, 0x2b, 0xdb, 0x34, 0x46, 0x4d, 0x63, 0xaa, 0xe9, 0xd7, 0x36, 0x05, 0xa6, 0x9a,
	0xc6, 0x54, 0xd3, 0x6f, 0x02, 0x20, 0xbc, 0x27, 0x4c, 0x35, 0xfd, 0xd6, 0x07, 0x76, 0xee, 0x09,
	0x13, 0x4d, 0xbf, 0xb3, 0xf7, 0x84, 0x89, 0xa6, 0x31, 0xd5, 0xf4, 0xb1, 0x3b, 0x3f, 0xd1, 0x34,
	0xa6, 0x9a, 0x3e, 0x09, 0x80, 0x28, 0x41, 0xdc, 0xf4, 0xc2, 0x07, 0xe2, 0x04, 0x51, 0xd3, 0xa7,
	0x2e, 0x41, 0xdc, 0x34, 0xfe, 0xbf, 0xe9, 0xa5, 0x6d, 0x1a, 0xd3, 0x4d, 0xcf, 0xc5, 0x01, 0x75,
	0x70, 0x2e, 0xf3, 0x8b, 0xa6, 0x06, 0x69, 0xfa, 0xf7, 0x61, 0xe0, 0xde, 0x33, 0x37, 0xa6, 0xee,
	0x84, 0xc5, 0x63, 0xd6, 0xfc, 0x85, 0x3b, 0x10, 0xba, 0xdf, 0xfc, 0xe0, 0x16, 0x06, 0x20, 0x0e,
	0x3f, 0x7b, 0x5f, 0x64, 0x6c, 0x33, 0x40, 0x9a, 0x64, 0x69, 0x9f, 0xa6, 0xc9, 0x99, 0x8d, 0x66,
	0x95, 0x6f, 0xbd, 0xb0, 0xe8, 0xe7, 0xd9, 0x3d, 0x71, 0x63, 0x70, 0x5b, 0xdf, 0x47, 0xf6, 0x5d,
	0xb7, 0x33, 0xb6, 0x9c, 0x5f, 0xe5, 0x07, 0x71, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xbe,
	0xb5, 0xe8, 0x28, 0x05, 0x00, 0x00,
}