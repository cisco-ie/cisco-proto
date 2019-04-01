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
// source: igmp_edm_bvi_stats_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_process_bvi_statistics

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

type IgmpEdmBviStatsBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmBviStatsBag_KEYS) Reset()         { *m = IgmpEdmBviStatsBag_KEYS{} }
func (m *IgmpEdmBviStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmBviStatsBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmBviStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d212e7de14617b11, []int{0}
}

func (m *IgmpEdmBviStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmBviStatsBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmBviStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmBviStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmBviStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmBviStatsBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmBviStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmBviStatsBag_KEYS.Size(m)
}
func (m *IgmpEdmBviStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmBviStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmBviStatsBag_KEYS proto.InternalMessageInfo

type IgmpEdmBviStatsBag struct {
	ReceiveBuffers       uint32   `protobuf:"varint,50,opt,name=receive_buffers,json=receiveBuffers,proto3" json:"receive_buffers,omitempty"`
	ReleaseBuffers       uint32   `protobuf:"varint,51,opt,name=release_buffers,json=releaseBuffers,proto3" json:"release_buffers,omitempty"`
	SendBlocks           uint32   `protobuf:"varint,52,opt,name=send_blocks,json=sendBlocks,proto3" json:"send_blocks,omitempty"`
	ReleaseFailBuffers   uint32   `protobuf:"varint,53,opt,name=release_fail_buffers,json=releaseFailBuffers,proto3" json:"release_fail_buffers,omitempty"`
	NullBufferHandles    uint32   `protobuf:"varint,54,opt,name=null_buffer_handles,json=nullBufferHandles,proto3" json:"null_buffer_handles,omitempty"`
	RxIpcOpenNotif       uint32   `protobuf:"varint,55,opt,name=rx_ipc_open_notif,json=rxIpcOpenNotif,proto3" json:"rx_ipc_open_notif,omitempty"`
	RxIpcCloseNotif      uint32   `protobuf:"varint,56,opt,name=rx_ipc_close_notif,json=rxIpcCloseNotif,proto3" json:"rx_ipc_close_notif,omitempty"`
	RxIpcErrorNotif      uint32   `protobuf:"varint,57,opt,name=rx_ipc_error_notif,json=rxIpcErrorNotif,proto3" json:"rx_ipc_error_notif,omitempty"`
	RxIpcLwmNotif        uint32   `protobuf:"varint,58,opt,name=rx_ipc_lwm_notif,json=rxIpcLwmNotif,proto3" json:"rx_ipc_lwm_notif,omitempty"`
	RxIpcInputWaitNotif  uint32   `protobuf:"varint,59,opt,name=rx_ipc_input_wait_notif,json=rxIpcInputWaitNotif,proto3" json:"rx_ipc_input_wait_notif,omitempty"`
	RxIpcSendStatusNotif uint32   `protobuf:"varint,60,opt,name=rx_ipc_send_status_notif,json=rxIpcSendStatusNotif,proto3" json:"rx_ipc_send_status_notif,omitempty"`
	RxIpcPublishNotif    uint32   `protobuf:"varint,61,opt,name=rx_ipc_publish_notif,json=rxIpcPublishNotif,proto3" json:"rx_ipc_publish_notif,omitempty"`
	RxIpcQFullNotif      uint32   `protobuf:"varint,62,opt,name=rx_ipc_q_full_notif,json=rxIpcQFullNotif,proto3" json:"rx_ipc_q_full_notif,omitempty"`
	RxIpcOutputNotif     uint32   `protobuf:"varint,63,opt,name=rx_ipc_output_notif,json=rxIpcOutputNotif,proto3" json:"rx_ipc_output_notif,omitempty"`
	RxIpcConnectNotif    uint32   `protobuf:"varint,64,opt,name=rx_ipc_connect_notif,json=rxIpcConnectNotif,proto3" json:"rx_ipc_connect_notif,omitempty"`
	RxIgmpPacketSuccess  uint32   `protobuf:"varint,65,opt,name=rx_igmp_packet_success,json=rxIgmpPacketSuccess,proto3" json:"rx_igmp_packet_success,omitempty"`
	RxAddMrouterMsg      uint32   `protobuf:"varint,66,opt,name=rx_add_mrouter_msg,json=rxAddMrouterMsg,proto3" json:"rx_add_mrouter_msg,omitempty"`
	RxDeleteMrouterMsg   uint32   `protobuf:"varint,67,opt,name=rx_delete_mrouter_msg,json=rxDeleteMrouterMsg,proto3" json:"rx_delete_mrouter_msg,omitempty"`
	RxSweepMrouterMsg    uint32   `protobuf:"varint,68,opt,name=rx_sweep_mrouter_msg,json=rxSweepMrouterMsg,proto3" json:"rx_sweep_mrouter_msg,omitempty"`
	TxAddMrouterMsg      uint32   `protobuf:"varint,69,opt,name=tx_add_mrouter_msg,json=txAddMrouterMsg,proto3" json:"tx_add_mrouter_msg,omitempty"`
	TxDeleteMrouterMsg   uint32   `protobuf:"varint,70,opt,name=tx_delete_mrouter_msg,json=txDeleteMrouterMsg,proto3" json:"tx_delete_mrouter_msg,omitempty"`
	TxSweepMrouterMsg    uint32   `protobuf:"varint,71,opt,name=tx_sweep_mrouter_msg,json=txSweepMrouterMsg,proto3" json:"tx_sweep_mrouter_msg,omitempty"`
	RxUnknownMrouterMsg  uint32   `protobuf:"varint,72,opt,name=rx_unknown_mrouter_msg,json=rxUnknownMrouterMsg,proto3" json:"rx_unknown_mrouter_msg,omitempty"`
	TxUnknownMrouterMsg  uint32   `protobuf:"varint,73,opt,name=tx_unknown_mrouter_msg,json=txUnknownMrouterMsg,proto3" json:"tx_unknown_mrouter_msg,omitempty"`
	TxBufferErrors       uint32   `protobuf:"varint,74,opt,name=tx_buffer_errors,json=txBufferErrors,proto3" json:"tx_buffer_errors,omitempty"`
	TxBuffers            uint32   `protobuf:"varint,75,opt,name=tx_buffers,json=txBuffers,proto3" json:"tx_buffers,omitempty"`
	TxProtocolBuffers    uint32   `protobuf:"varint,76,opt,name=tx_protocol_buffers,json=txProtocolBuffers,proto3" json:"tx_protocol_buffers,omitempty"`
	TxMrouterBuffers     uint32   `protobuf:"varint,77,opt,name=tx_mrouter_buffers,json=txMrouterBuffers,proto3" json:"tx_mrouter_buffers,omitempty"`
	TxUnknownBuffers     uint32   `protobuf:"varint,78,opt,name=tx_unknown_buffers,json=txUnknownBuffers,proto3" json:"tx_unknown_buffers,omitempty"`
	WtxMsgRecvd          uint32   `protobuf:"varint,79,opt,name=wtx_msg_recvd,json=wtxMsgRecvd,proto3" json:"wtx_msg_recvd,omitempty"`
	WtxMsgSent           uint32   `protobuf:"varint,80,opt,name=wtx_msg_sent,json=wtxMsgSent,proto3" json:"wtx_msg_sent,omitempty"`
	WtxMsgProtoSent      uint32   `protobuf:"varint,81,opt,name=wtx_msg_proto_sent,json=wtxMsgProtoSent,proto3" json:"wtx_msg_proto_sent,omitempty"`
	WtxMsgDropDc         uint32   `protobuf:"varint,82,opt,name=wtx_msg_drop_dc,json=wtxMsgDropDc,proto3" json:"wtx_msg_drop_dc,omitempty"`
	WtxMsgDropNomem      uint32   `protobuf:"varint,83,opt,name=wtx_msg_drop_nomem,json=wtxMsgDropNomem,proto3" json:"wtx_msg_drop_nomem,omitempty"`
	WtxMsgFreed          uint32   `protobuf:"varint,84,opt,name=wtx_msg_freed,json=wtxMsgFreed,proto3" json:"wtx_msg_freed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmBviStatsBag) Reset()         { *m = IgmpEdmBviStatsBag{} }
func (m *IgmpEdmBviStatsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmBviStatsBag) ProtoMessage()    {}
func (*IgmpEdmBviStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d212e7de14617b11, []int{1}
}

func (m *IgmpEdmBviStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmBviStatsBag.Unmarshal(m, b)
}
func (m *IgmpEdmBviStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmBviStatsBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmBviStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmBviStatsBag.Merge(m, src)
}
func (m *IgmpEdmBviStatsBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmBviStatsBag.Size(m)
}
func (m *IgmpEdmBviStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmBviStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmBviStatsBag proto.InternalMessageInfo

func (m *IgmpEdmBviStatsBag) GetReceiveBuffers() uint32 {
	if m != nil {
		return m.ReceiveBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetReleaseBuffers() uint32 {
	if m != nil {
		return m.ReleaseBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetSendBlocks() uint32 {
	if m != nil {
		return m.SendBlocks
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetReleaseFailBuffers() uint32 {
	if m != nil {
		return m.ReleaseFailBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetNullBufferHandles() uint32 {
	if m != nil {
		return m.NullBufferHandles
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcOpenNotif() uint32 {
	if m != nil {
		return m.RxIpcOpenNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcCloseNotif() uint32 {
	if m != nil {
		return m.RxIpcCloseNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcErrorNotif() uint32 {
	if m != nil {
		return m.RxIpcErrorNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcLwmNotif() uint32 {
	if m != nil {
		return m.RxIpcLwmNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcInputWaitNotif() uint32 {
	if m != nil {
		return m.RxIpcInputWaitNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcSendStatusNotif() uint32 {
	if m != nil {
		return m.RxIpcSendStatusNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcPublishNotif() uint32 {
	if m != nil {
		return m.RxIpcPublishNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcQFullNotif() uint32 {
	if m != nil {
		return m.RxIpcQFullNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcOutputNotif() uint32 {
	if m != nil {
		return m.RxIpcOutputNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIpcConnectNotif() uint32 {
	if m != nil {
		return m.RxIpcConnectNotif
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxIgmpPacketSuccess() uint32 {
	if m != nil {
		return m.RxIgmpPacketSuccess
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxAddMrouterMsg() uint32 {
	if m != nil {
		return m.RxAddMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxDeleteMrouterMsg() uint32 {
	if m != nil {
		return m.RxDeleteMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxSweepMrouterMsg() uint32 {
	if m != nil {
		return m.RxSweepMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxAddMrouterMsg() uint32 {
	if m != nil {
		return m.TxAddMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxDeleteMrouterMsg() uint32 {
	if m != nil {
		return m.TxDeleteMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxSweepMrouterMsg() uint32 {
	if m != nil {
		return m.TxSweepMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetRxUnknownMrouterMsg() uint32 {
	if m != nil {
		return m.RxUnknownMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxUnknownMrouterMsg() uint32 {
	if m != nil {
		return m.TxUnknownMrouterMsg
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxBufferErrors() uint32 {
	if m != nil {
		return m.TxBufferErrors
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxBuffers() uint32 {
	if m != nil {
		return m.TxBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxProtocolBuffers() uint32 {
	if m != nil {
		return m.TxProtocolBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxMrouterBuffers() uint32 {
	if m != nil {
		return m.TxMrouterBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetTxUnknownBuffers() uint32 {
	if m != nil {
		return m.TxUnknownBuffers
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgRecvd() uint32 {
	if m != nil {
		return m.WtxMsgRecvd
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgSent() uint32 {
	if m != nil {
		return m.WtxMsgSent
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgProtoSent() uint32 {
	if m != nil {
		return m.WtxMsgProtoSent
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgDropDc() uint32 {
	if m != nil {
		return m.WtxMsgDropDc
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgDropNomem() uint32 {
	if m != nil {
		return m.WtxMsgDropNomem
	}
	return 0
}

func (m *IgmpEdmBviStatsBag) GetWtxMsgFreed() uint32 {
	if m != nil {
		return m.WtxMsgFreed
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmBviStatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.process.bvi_statistics.igmp_edm_bvi_stats_bag_KEYS")
	proto.RegisterType((*IgmpEdmBviStatsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.process.bvi_statistics.igmp_edm_bvi_stats_bag")
}

func init() { proto.RegisterFile("igmp_edm_bvi_stats_bag.proto", fileDescriptor_d212e7de14617b11) }

var fileDescriptor_d212e7de14617b11 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x6f, 0x4f, 0x1b, 0x39,
	0x10, 0xc6, 0x75, 0x6f, 0x4e, 0x3a, 0x03, 0x07, 0x2c, 0x1c, 0xb7, 0xd2, 0x1d, 0x3a, 0x84, 0x74,
	0x82, 0x13, 0x5c, 0xfa, 0x07, 0x4a, 0xff, 0x97, 0x02, 0x21, 0x85, 0x42, 0x20, 0x90, 0x56, 0x55,
	0x5f, 0x59, 0x1b, 0xaf, 0x13, 0x2c, 0x76, 0xed, 0xad, 0xed, 0x4d, 0xb6, 0x1f, 0xb1, 0xdf, 0xaa,
	0xf2, 0x8c, 0x1d, 0x36, 0x90, 0x77, 0xd1, 0x3c, 0xbf, 0x67, 0xf2, 0xcc, 0xac, 0x6d, 0xf2, 0xb7,
	0x18, 0xe4, 0x05, 0xe5, 0x69, 0x4e, 0x7b, 0x43, 0x41, 0x8d, 0x4d, 0xac, 0xa1, 0xbd, 0x64, 0xd0,
	0x28, 0xb4, 0xb2, 0x2a, 0xda, 0x67, 0xc2, 0x30, 0x45, 0x85, 0x32, 0xb4, 0xd2, 0x54, 0x14, 0xc3,
	0x5d, 0x0a, 0xbc, 0x2a, 0xb8, 0x6e, 0xb8, 0x5f, 0x0d, 0x63, 0x13, 0x99, 0xf6, 0xbe, 0x3b, 0x9e,
	0x71, 0x63, 0x1a, 0xa1, 0x8b, 0x30, 0x56, 0x30, 0xb3, 0xbe, 0x4a, 0xfe, 0x9a, 0xfe, 0x07, 0xf4,
	0xec, 0xf8, 0x6b, 0x77, 0xfd, 0xc7, 0x2c, 0x59, 0x99, 0xae, 0x47, 0x1b, 0x64, 0x5e, 0x73, 0xc6,
	0xc5, 0x90, 0xd3, 0x5e, 0xd9, 0xef, 0x73, 0x6d, 0xe2, 0xa7, 0x6b, 0xbf, 0x6c, 0xce, 0x5d, 0xff,
	0xee, 0xcb, 0x87, 0x58, 0x45, 0x30, 0xe3, 0x89, 0xb9, 0x03, 0x77, 0x02, 0x08, 0xe5, 0x00, 0xfe,
	0x43, 0x66, 0x0c, 0x97, 0x29, 0xed, 0x65, 0x8a, 0xdd, 0x9a, 0x78, 0x17, 0x20, 0xe2, 0x4a, 0x87,
	0x50, 0x89, 0x1e, 0x93, 0xe5, 0xd0, 0xa9, 0x9f, 0x88, 0x6c, 0xdc, 0xee, 0x19, 0x90, 0x91, 0xd7,
	0x5a, 0x89, 0xc8, 0x42, 0xcb, 0x06, 0x59, 0x92, 0x65, 0x16, 0x48, 0x7a, 0x93, 0xc8, 0x34, 0xe3,
	0x26, 0xde, 0x03, 0xc3, 0xa2, 0x93, 0x90, 0x3c, 0x41, 0x21, 0xfa, 0x8f, 0x2c, 0xea, 0x8a, 0x8a,
	0x82, 0xb9, 0x0d, 0x4a, 0x2a, 0x95, 0x15, 0xfd, 0xf8, 0xb9, 0x4f, 0x5b, 0x9d, 0x16, 0xec, 0xb2,
	0xe0, 0xf2, 0xc2, 0x55, 0xa3, 0x2d, 0x12, 0x79, 0x94, 0x65, 0xca, 0x70, 0xcf, 0xbe, 0x00, 0x76,
	0x1e, 0xd8, 0x23, 0x57, 0xbf, 0x0f, 0x73, 0xad, 0x95, 0xf6, 0xf0, 0xcb, 0x1a, 0x7c, 0xec, 0xea,
	0x08, 0x6f, 0x90, 0x05, 0x0f, 0x67, 0xa3, 0xdc, 0xa3, 0xaf, 0x00, 0x9d, 0x03, 0xf4, 0x7c, 0x94,
	0x23, 0xb8, 0x4b, 0xfe, 0xf4, 0xa0, 0x90, 0x45, 0x69, 0xe9, 0x28, 0x11, 0xd6, 0xf3, 0xaf, 0x81,
	0x5f, 0x02, 0xfe, 0xd4, 0x89, 0x5f, 0x12, 0x61, 0xd1, 0xb5, 0x47, 0x62, 0xef, 0x82, 0x6d, 0xbb,
	0x2f, 0x5a, 0x1a, 0x6f, 0x7b, 0x03, 0xb6, 0x65, 0xb0, 0x75, 0xb9, 0x4c, 0xbb, 0x20, 0xa2, 0xef,
	0x11, 0x59, 0xf6, 0xbe, 0xa2, 0xec, 0x65, 0xc2, 0xdc, 0x78, 0xcf, 0x5b, 0x5c, 0x26, 0x78, 0x3a,
	0xa8, 0xa0, 0x61, 0x9b, 0x2c, 0x79, 0xc3, 0x37, 0xda, 0x77, 0x5f, 0x01, 0xf9, 0x77, 0xb5, 0xa9,
	0xaf, 0x5a, 0x65, 0x96, 0x21, 0xfd, 0xff, 0x98, 0x56, 0xa5, 0x75, 0xd3, 0x20, 0xbd, 0x0f, 0xf4,
	0x02, 0x2e, 0x1f, 0x84, 0xfb, 0x69, 0x98, 0x92, 0x92, 0xb3, 0xc0, 0xbf, 0xaf, 0xa5, 0x39, 0x42,
	0x05, 0x0d, 0x3b, 0x64, 0xc5, 0x19, 0xdc, 0x61, 0x2e, 0x12, 0x76, 0xcb, 0x2d, 0x35, 0x25, 0x73,
	0x57, 0x22, 0x3e, 0x18, 0xef, 0x6a, 0x90, 0x17, 0x1d, 0xd0, 0xba, 0x28, 0xf9, 0xef, 0x96, 0xa4,
	0x29, 0xcd, 0xb5, 0x2a, 0x2d, 0xd7, 0x34, 0x37, 0x83, 0xf8, 0x30, 0x4c, 0x70, 0x90, 0xa6, 0x6d,
	0xac, 0xb7, 0xcd, 0x20, 0x7a, 0x42, 0xfe, 0xd0, 0x15, 0x4d, 0x79, 0xc6, 0x2d, 0x9f, 0xe0, 0x8f,
	0xfc, 0xf9, 0xac, 0x9a, 0xa0, 0xd5, 0x2c, 0x38, 0x85, 0x19, 0x71, 0x5e, 0x4c, 0x38, 0x9a, 0x61,
	0x8a, 0xae, 0x93, 0x6a, 0x86, 0x2d, 0x12, 0xd9, 0x87, 0x81, 0x8e, 0x31, 0x90, 0x7d, 0x18, 0xc8,
	0x4e, 0x0d, 0xd4, 0xc2, 0x40, 0x76, 0x6a, 0x20, 0x3b, 0x2d, 0xd0, 0x07, 0x0c, 0x64, 0x1f, 0x04,
	0xc2, 0xb5, 0x96, 0xf2, 0x56, 0xaa, 0x91, 0x9c, 0xb0, 0x9c, 0x84, 0xb5, 0x7e, 0x46, 0x71, 0xd2,
	0x64, 0xa7, 0x9b, 0x4e, 0xd1, 0x64, 0xa7, 0x98, 0x36, 0xc9, 0x82, 0xad, 0xc2, 0x4d, 0x86, 0x6b,
	0x64, 0xe2, 0x8f, 0x78, 0x35, 0x6d, 0x85, 0xd7, 0x18, 0x2e, 0x91, 0x89, 0x56, 0x09, 0x19, 0x93,
	0x26, 0x3e, 0x03, 0xe6, 0xb7, 0xc0, 0xc0, 0xa3, 0x60, 0x2b, 0x0a, 0x0f, 0x28, 0x53, 0x77, 0xaf,
	0xc8, 0x79, 0x18, 0xb1, 0xe3, 0x95, 0xc0, 0x6f, 0xc3, 0xce, 0x43, 0xca, 0x80, 0xb7, 0xf1, 0x60,
	0xda, 0xca, 0x47, 0x9c, 0xa4, 0xc3, 0x6c, 0x81, 0xbe, 0x08, 0xb4, 0x9f, 0x2b, 0xd0, 0xeb, 0x64,
	0x6e, 0xe4, 0x9a, 0x9b, 0x01, 0xd5, 0x9c, 0x0d, 0xd3, 0xf8, 0x12, 0xc0, 0x99, 0x91, 0xad, 0xda,
	0x66, 0x70, 0xed, 0x4a, 0xd1, 0x1a, 0x99, 0x0d, 0x8c, 0xe1, 0xd2, 0xc6, 0x1d, 0x7c, 0x18, 0x11,
	0xe9, 0x72, 0x69, 0xdd, 0xa9, 0x08, 0x04, 0x8c, 0x85, 0xdc, 0x15, 0x9e, 0x0a, 0xe4, 0x60, 0x28,
	0x80, 0xff, 0x25, 0xf3, 0x01, 0x4e, 0xb5, 0x2a, 0x68, 0xca, 0xe2, 0x6b, 0x20, 0x67, 0x91, 0x6c,
	0x6a, 0x55, 0x34, 0x59, 0xbd, 0x27, 0x60, 0x52, 0xe5, 0x3c, 0x8f, 0xbb, 0xf5, 0x9e, 0x8e, 0xbc,
	0x70, 0xe5, 0xfa, 0x18, 0x7d, 0xcd, 0x79, 0x1a, 0x7f, 0xaa, 0x8f, 0xd1, 0x72, 0xa5, 0xde, 0xaf,
	0x10, 0x6d, 0xe7, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x68, 0x77, 0x1f, 0xd2, 0x06, 0x00,
	0x00,
}
