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
// source: tcp_sh_pcb_stats_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_connection_nodes_node_statistics_pcbs_pcb

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

type TcpShPcbStatsBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 uint32   `protobuf:"varint,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpShPcbStatsBag_KEYS) Reset()         { *m = TcpShPcbStatsBag_KEYS{} }
func (m *TcpShPcbStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpShPcbStatsBag_KEYS) ProtoMessage()    {}
func (*TcpShPcbStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f444975b9ae036c3, []int{0}
}

func (m *TcpShPcbStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShPcbStatsBag_KEYS.Unmarshal(m, b)
}
func (m *TcpShPcbStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShPcbStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpShPcbStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShPcbStatsBag_KEYS.Merge(m, src)
}
func (m *TcpShPcbStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpShPcbStatsBag_KEYS.Size(m)
}
func (m *TcpShPcbStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShPcbStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShPcbStatsBag_KEYS proto.InternalMessageInfo

func (m *TcpShPcbStatsBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpShPcbStatsBag_KEYS) GetId_1() uint32 {
	if m != nil {
		return m.Id_1
	}
	return 0
}

type TcpIoCountsBag struct {
	IoCount              uint32   `protobuf:"varint,1,opt,name=io_count,json=ioCount,proto3" json:"io_count,omitempty"`
	ArmCount             uint32   `protobuf:"varint,2,opt,name=arm_count,json=armCount,proto3" json:"arm_count,omitempty"`
	UnarmCount           uint32   `protobuf:"varint,3,opt,name=unarm_count,json=unarmCount,proto3" json:"unarm_count,omitempty"`
	AutoarmCount         uint32   `protobuf:"varint,4,opt,name=autoarm_count,json=autoarmCount,proto3" json:"autoarm_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpIoCountsBag) Reset()         { *m = TcpIoCountsBag{} }
func (m *TcpIoCountsBag) String() string { return proto.CompactTextString(m) }
func (*TcpIoCountsBag) ProtoMessage()    {}
func (*TcpIoCountsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f444975b9ae036c3, []int{1}
}

func (m *TcpIoCountsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpIoCountsBag.Unmarshal(m, b)
}
func (m *TcpIoCountsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpIoCountsBag.Marshal(b, m, deterministic)
}
func (m *TcpIoCountsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpIoCountsBag.Merge(m, src)
}
func (m *TcpIoCountsBag) XXX_Size() int {
	return xxx_messageInfo_TcpIoCountsBag.Size(m)
}
func (m *TcpIoCountsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpIoCountsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpIoCountsBag proto.InternalMessageInfo

func (m *TcpIoCountsBag) GetIoCount() uint32 {
	if m != nil {
		return m.IoCount
	}
	return 0
}

func (m *TcpIoCountsBag) GetArmCount() uint32 {
	if m != nil {
		return m.ArmCount
	}
	return 0
}

func (m *TcpIoCountsBag) GetUnarmCount() uint32 {
	if m != nil {
		return m.UnarmCount
	}
	return 0
}

func (m *TcpIoCountsBag) GetAutoarmCount() uint32 {
	if m != nil {
		return m.AutoarmCount
	}
	return 0
}

type TcpAsyncSessionStatsBag struct {
	AsyncSession           bool     `protobuf:"varint,1,opt,name=async_session,json=asyncSession,proto3" json:"async_session,omitempty"`
	DataWriteSuccessNum    []uint32 `protobuf:"varint,2,rep,packed,name=data_write_success_num,json=dataWriteSuccessNum,proto3" json:"data_write_success_num,omitempty"`
	DataReadSuccessNum     []uint32 `protobuf:"varint,3,rep,packed,name=data_read_success_num,json=dataReadSuccessNum,proto3" json:"data_read_success_num,omitempty"`
	DataWriteErrorNum      []uint32 `protobuf:"varint,4,rep,packed,name=data_write_error_num,json=dataWriteErrorNum,proto3" json:"data_write_error_num,omitempty"`
	DataReadErrorNum       []uint32 `protobuf:"varint,5,rep,packed,name=data_read_error_num,json=dataReadErrorNum,proto3" json:"data_read_error_num,omitempty"`
	ControlWriteSuccessNum []uint32 `protobuf:"varint,6,rep,packed,name=control_write_success_num,json=controlWriteSuccessNum,proto3" json:"control_write_success_num,omitempty"`
	ControlReadSuccessNum  []uint32 `protobuf:"varint,7,rep,packed,name=control_read_success_num,json=controlReadSuccessNum,proto3" json:"control_read_success_num,omitempty"`
	ControlWriteErrorNum   []uint32 `protobuf:"varint,8,rep,packed,name=control_write_error_num,json=controlWriteErrorNum,proto3" json:"control_write_error_num,omitempty"`
	ControlReadErrorNum    []uint32 `protobuf:"varint,9,rep,packed,name=control_read_error_num,json=controlReadErrorNum,proto3" json:"control_read_error_num,omitempty"`
	DataWriteByte          []uint64 `protobuf:"varint,10,rep,packed,name=data_write_byte,json=dataWriteByte,proto3" json:"data_write_byte,omitempty"`
	DataReadByte           []uint64 `protobuf:"varint,11,rep,packed,name=data_read_byte,json=dataReadByte,proto3" json:"data_read_byte,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TcpAsyncSessionStatsBag) Reset()         { *m = TcpAsyncSessionStatsBag{} }
func (m *TcpAsyncSessionStatsBag) String() string { return proto.CompactTextString(m) }
func (*TcpAsyncSessionStatsBag) ProtoMessage()    {}
func (*TcpAsyncSessionStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f444975b9ae036c3, []int{2}
}

func (m *TcpAsyncSessionStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpAsyncSessionStatsBag.Unmarshal(m, b)
}
func (m *TcpAsyncSessionStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpAsyncSessionStatsBag.Marshal(b, m, deterministic)
}
func (m *TcpAsyncSessionStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpAsyncSessionStatsBag.Merge(m, src)
}
func (m *TcpAsyncSessionStatsBag) XXX_Size() int {
	return xxx_messageInfo_TcpAsyncSessionStatsBag.Size(m)
}
func (m *TcpAsyncSessionStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpAsyncSessionStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpAsyncSessionStatsBag proto.InternalMessageInfo

func (m *TcpAsyncSessionStatsBag) GetAsyncSession() bool {
	if m != nil {
		return m.AsyncSession
	}
	return false
}

func (m *TcpAsyncSessionStatsBag) GetDataWriteSuccessNum() []uint32 {
	if m != nil {
		return m.DataWriteSuccessNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetDataReadSuccessNum() []uint32 {
	if m != nil {
		return m.DataReadSuccessNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetDataWriteErrorNum() []uint32 {
	if m != nil {
		return m.DataWriteErrorNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetDataReadErrorNum() []uint32 {
	if m != nil {
		return m.DataReadErrorNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetControlWriteSuccessNum() []uint32 {
	if m != nil {
		return m.ControlWriteSuccessNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetControlReadSuccessNum() []uint32 {
	if m != nil {
		return m.ControlReadSuccessNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetControlWriteErrorNum() []uint32 {
	if m != nil {
		return m.ControlWriteErrorNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetControlReadErrorNum() []uint32 {
	if m != nil {
		return m.ControlReadErrorNum
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetDataWriteByte() []uint64 {
	if m != nil {
		return m.DataWriteByte
	}
	return nil
}

func (m *TcpAsyncSessionStatsBag) GetDataReadByte() []uint64 {
	if m != nil {
		return m.DataReadByte
	}
	return nil
}

type TcpShPcbStatsBag struct {
	Pcb                        uint64                   `protobuf:"varint,50,opt,name=pcb,proto3" json:"pcb,omitempty"`
	VrfId                      uint32                   `protobuf:"varint,51,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	PacketsSent                uint64                   `protobuf:"varint,52,opt,name=packets_sent,json=packetsSent,proto3" json:"packets_sent,omitempty"`
	XipcPulseReceived          uint64                   `protobuf:"varint,53,opt,name=xipc_pulse_received,json=xipcPulseReceived,proto3" json:"xipc_pulse_received,omitempty"`
	SegmentInstructionReceived uint32                   `protobuf:"varint,54,opt,name=segment_instruction_received,json=segmentInstructionReceived,proto3" json:"segment_instruction_received,omitempty"`
	SendPacketsQueued          uint64                   `protobuf:"varint,55,opt,name=send_packets_queued,json=sendPacketsQueued,proto3" json:"send_packets_queued,omitempty"`
	SendPacketsQueuedNetIo     uint64                   `protobuf:"varint,56,opt,name=send_packets_queued_net_io,json=sendPacketsQueuedNetIo,proto3" json:"send_packets_queued_net_io,omitempty"`
	SendQueueFailed            uint32                   `protobuf:"varint,57,opt,name=send_queue_failed,json=sendQueueFailed,proto3" json:"send_queue_failed,omitempty"`
	SendQueueNetIoFailed       uint32                   `protobuf:"varint,58,opt,name=send_queue_net_io_failed,json=sendQueueNetIoFailed,proto3" json:"send_queue_net_io_failed,omitempty"`
	PacketsReceived            uint64                   `protobuf:"varint,59,opt,name=packets_received,json=packetsReceived,proto3" json:"packets_received,omitempty"`
	ReceiveQueueFailed         uint32                   `protobuf:"varint,60,opt,name=receive_queue_failed,json=receiveQueueFailed,proto3" json:"receive_queue_failed,omitempty"`
	ReceivedPacketsQueued      uint64                   `protobuf:"varint,61,opt,name=received_packets_queued,json=receivedPacketsQueued,proto3" json:"received_packets_queued,omitempty"`
	SendWindowShrinkIgnored    uint32                   `protobuf:"varint,62,opt,name=send_window_shrink_ignored,json=sendWindowShrinkIgnored,proto3" json:"send_window_shrink_ignored,omitempty"`
	IsPawSocket                bool                     `protobuf:"varint,63,opt,name=is_paw_socket,json=isPawSocket,proto3" json:"is_paw_socket,omitempty"`
	ReadIoCounts               *TcpIoCountsBag          `protobuf:"bytes,64,opt,name=read_io_counts,json=readIoCounts,proto3" json:"read_io_counts,omitempty"`
	WriteIoCounts              *TcpIoCountsBag          `protobuf:"bytes,65,opt,name=write_io_counts,json=writeIoCounts,proto3" json:"write_io_counts,omitempty"`
	ReadIoTime                 uint32                   `protobuf:"varint,66,opt,name=read_io_time,json=readIoTime,proto3" json:"read_io_time,omitempty"`
	WriteIoTime                uint32                   `protobuf:"varint,67,opt,name=write_io_time,json=writeIoTime,proto3" json:"write_io_time,omitempty"`
	AsyncSessionStats          *TcpAsyncSessionStatsBag `protobuf:"bytes,68,opt,name=async_session_stats,json=asyncSessionStats,proto3" json:"async_session_stats,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                 `json:"-"`
	XXX_unrecognized           []byte                   `json:"-"`
	XXX_sizecache              int32                    `json:"-"`
}

func (m *TcpShPcbStatsBag) Reset()         { *m = TcpShPcbStatsBag{} }
func (m *TcpShPcbStatsBag) String() string { return proto.CompactTextString(m) }
func (*TcpShPcbStatsBag) ProtoMessage()    {}
func (*TcpShPcbStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f444975b9ae036c3, []int{3}
}

func (m *TcpShPcbStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpShPcbStatsBag.Unmarshal(m, b)
}
func (m *TcpShPcbStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpShPcbStatsBag.Marshal(b, m, deterministic)
}
func (m *TcpShPcbStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpShPcbStatsBag.Merge(m, src)
}
func (m *TcpShPcbStatsBag) XXX_Size() int {
	return xxx_messageInfo_TcpShPcbStatsBag.Size(m)
}
func (m *TcpShPcbStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpShPcbStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpShPcbStatsBag proto.InternalMessageInfo

func (m *TcpShPcbStatsBag) GetPcb() uint64 {
	if m != nil {
		return m.Pcb
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetPacketsSent() uint64 {
	if m != nil {
		return m.PacketsSent
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetXipcPulseReceived() uint64 {
	if m != nil {
		return m.XipcPulseReceived
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSegmentInstructionReceived() uint32 {
	if m != nil {
		return m.SegmentInstructionReceived
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSendPacketsQueued() uint64 {
	if m != nil {
		return m.SendPacketsQueued
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSendPacketsQueuedNetIo() uint64 {
	if m != nil {
		return m.SendPacketsQueuedNetIo
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSendQueueFailed() uint32 {
	if m != nil {
		return m.SendQueueFailed
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSendQueueNetIoFailed() uint32 {
	if m != nil {
		return m.SendQueueNetIoFailed
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetPacketsReceived() uint64 {
	if m != nil {
		return m.PacketsReceived
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetReceiveQueueFailed() uint32 {
	if m != nil {
		return m.ReceiveQueueFailed
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetReceivedPacketsQueued() uint64 {
	if m != nil {
		return m.ReceivedPacketsQueued
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetSendWindowShrinkIgnored() uint32 {
	if m != nil {
		return m.SendWindowShrinkIgnored
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetIsPawSocket() bool {
	if m != nil {
		return m.IsPawSocket
	}
	return false
}

func (m *TcpShPcbStatsBag) GetReadIoCounts() *TcpIoCountsBag {
	if m != nil {
		return m.ReadIoCounts
	}
	return nil
}

func (m *TcpShPcbStatsBag) GetWriteIoCounts() *TcpIoCountsBag {
	if m != nil {
		return m.WriteIoCounts
	}
	return nil
}

func (m *TcpShPcbStatsBag) GetReadIoTime() uint32 {
	if m != nil {
		return m.ReadIoTime
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetWriteIoTime() uint32 {
	if m != nil {
		return m.WriteIoTime
	}
	return 0
}

func (m *TcpShPcbStatsBag) GetAsyncSessionStats() *TcpAsyncSessionStatsBag {
	if m != nil {
		return m.AsyncSessionStats
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpShPcbStatsBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.pcbs.pcb.tcp_sh_pcb_stats_bag_KEYS")
	proto.RegisterType((*TcpIoCountsBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.pcbs.pcb.tcp_io_counts_bag")
	proto.RegisterType((*TcpAsyncSessionStatsBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.pcbs.pcb.tcp_async_session_stats_bag")
	proto.RegisterType((*TcpShPcbStatsBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_connection.nodes.node.statistics.pcbs.pcb.tcp_sh_pcb_stats_bag")
}

func init() { proto.RegisterFile("tcp_sh_pcb_stats_bag.proto", fileDescriptor_f444975b9ae036c3) }

var fileDescriptor_f444975b9ae036c3 = []byte{
	// 854 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0x95, 0x63, 0x37, 0x75, 0xae, 0xed, 0x24, 0x1e, 0x3b, 0xc9, 0x24, 0x45, 0xc2, 0x18, 0x84,
	0x0c, 0x12, 0x0b, 0xa9, 0x69, 0x4b, 0x5b, 0x28, 0xa5, 0xa5, 0x95, 0x2c, 0xa4, 0x2a, 0xac, 0x91,
	0xaa, 0x3e, 0x8d, 0xd6, 0xb3, 0x93, 0x74, 0xd4, 0x78, 0x66, 0x99, 0x99, 0x8d, 0x9b, 0x27, 0xc4,
	0x4f, 0x40, 0xfc, 0x19, 0xfe, 0x1b, 0x2f, 0x68, 0xee, 0x7e, 0xd9, 0x89, 0x79, 0x4b, 0x5f, 0xac,
	0xf5, 0x3d, 0xe7, 0xdc, 0x73, 0xee, 0xd5, 0xee, 0x0c, 0x1c, 0x39, 0x9e, 0x30, 0xfb, 0x96, 0x25,
	0x7c, 0xc6, 0xac, 0x8b, 0x9c, 0x65, 0xb3, 0xe8, 0x2c, 0x48, 0x8c, 0x76, 0x9a, 0xbc, 0xe4, 0xd2,
	0x72, 0xcd, 0xa4, 0xb6, 0xec, 0xbd, 0x61, 0x32, 0x61, 0x9e, 0xab, 0x13, 0x61, 0x02, 0xff, 0xc0,
	0xb5, 0x52, 0x82, 0x3b, 0xa9, 0x55, 0xa0, 0x74, 0x2c, 0x2c, 0xfe, 0x06, 0xbe, 0x87, 0xb4, 0x4e,
	0x72, 0x1b, 0x24, 0x7c, 0x86, 0x3f, 0xc3, 0x27, 0x70, 0xb8, 0xce, 0x85, 0xfd, 0xf2, 0xe2, 0xcd,
	0x94, 0x6c, 0xc3, 0x86, 0x8c, 0x69, 0x6d, 0x50, 0x1b, 0x6d, 0x85, 0x1b, 0x32, 0x26, 0x5d, 0x68,
	0xc8, 0x98, 0x1d, 0xd3, 0x8d, 0x41, 0x6d, 0xd4, 0x09, 0xeb, 0x32, 0x3e, 0x1e, 0xfe, 0x55, 0x83,
	0xae, 0x6f, 0x20, 0x35, 0xe3, 0x3a, 0x55, 0x99, 0x9a, 0x1c, 0x42, 0xb3, 0x28, 0xa0, 0xbc, 0x13,
	0xde, 0x96, 0xfa, 0xb9, 0xff, 0x4b, 0xee, 0xc0, 0x56, 0x64, 0xe6, 0x39, 0x96, 0x35, 0x6a, 0x46,
	0x66, 0x9e, 0x81, 0x1f, 0x43, 0x2b, 0x55, 0x15, 0x5c, 0x47, 0x18, 0xb0, 0x94, 0x11, 0x3e, 0x85,
	0x4e, 0x94, 0x3a, 0x5d, 0x51, 0x1a, 0x48, 0x69, 0xe7, 0x45, 0x24, 0x0d, 0xff, 0x69, 0xc0, 0x1d,
	0x9f, 0x29, 0xb2, 0x97, 0x8a, 0x33, 0x2b, 0xac, 0x95, 0x5a, 0x55, 0xb3, 0x61, 0x93, 0x65, 0x08,
	0x23, 0x36, 0xc3, 0x36, 0x16, 0xa7, 0x59, 0x8d, 0x8c, 0x61, 0x3f, 0x8e, 0x5c, 0xc4, 0x16, 0x46,
	0x3a, 0xc1, 0x6c, 0xca, 0xb9, 0xb0, 0x96, 0xa9, 0x74, 0x4e, 0x37, 0x06, 0xf5, 0x51, 0x27, 0xec,
	0x79, 0xf4, 0xb5, 0x07, 0xa7, 0x19, 0xf6, 0x2a, 0x9d, 0x93, 0x63, 0xd8, 0x43, 0x91, 0x11, 0x51,
	0xbc, 0xa2, 0xa9, 0xa3, 0x86, 0x78, 0x30, 0x14, 0x51, 0xbc, 0x24, 0xf9, 0x1a, 0xfa, 0x4b, 0x3e,
	0xc2, 0x18, 0x6d, 0x50, 0xd1, 0x40, 0x45, 0xb7, 0x74, 0x79, 0xe1, 0x11, 0x2f, 0xf8, 0x0a, 0x7a,
	0x95, 0x47, 0xc5, 0xbf, 0x85, 0xfc, 0xdd, 0xc2, 0xa1, 0xa4, 0x3f, 0x84, 0x43, 0xae, 0x95, 0x33,
	0xfa, 0x7c, 0xcd, 0x28, 0x9b, 0x28, 0xda, 0xcf, 0x09, 0x57, 0xa7, 0x79, 0x00, 0xb4, 0x90, 0x5e,
	0x1b, 0xe8, 0x36, 0x2a, 0xf7, 0x72, 0xfc, 0xca, 0x4c, 0xf7, 0xe0, 0x60, 0xd5, 0xb3, 0x8a, 0xd9,
	0x44, 0x5d, 0x7f, 0xd9, 0xb1, 0x8c, 0x3a, 0x86, 0xfd, 0x15, 0xbf, 0x4a, 0xb5, 0x95, 0xad, 0x7c,
	0xc9, 0xad, 0x14, 0x7d, 0x0e, 0x3b, 0x4b, 0xfb, 0x9b, 0x5d, 0x3a, 0x41, 0x61, 0x50, 0x1f, 0x35,
	0xc2, 0x4e, 0xb9, 0xba, 0x67, 0x97, 0x4e, 0x90, 0xcf, 0x60, 0xbb, 0x5a, 0x1b, 0xd2, 0x5a, 0x48,
	0x6b, 0x17, 0x1b, 0xf3, 0xac, 0xe1, 0xbf, 0x4d, 0xe8, 0xaf, 0xfb, 0x1e, 0xc8, 0x2e, 0xd4, 0x13,
	0x3e, 0xa3, 0x77, 0x07, 0xb5, 0x51, 0x23, 0xf4, 0x8f, 0x64, 0x0f, 0x36, 0x2f, 0xcc, 0x29, 0x93,
	0x31, 0x1d, 0xe3, 0x3b, 0x78, 0xeb, 0xc2, 0x9c, 0x4e, 0x62, 0xf2, 0x09, 0xb4, 0x93, 0x88, 0xbf,
	0x13, 0xce, 0x32, 0x2b, 0x94, 0xa3, 0xdf, 0xa2, 0xa2, 0x95, 0xd7, 0xa6, 0x42, 0x39, 0x12, 0x40,
	0xef, 0xbd, 0x4c, 0x38, 0x4b, 0xd2, 0x73, 0x2b, 0x98, 0x11, 0x5c, 0xc8, 0x0b, 0x11, 0xd3, 0x7b,
	0xc8, 0xec, 0x7a, 0xe8, 0xc4, 0x23, 0x61, 0x0e, 0x90, 0xa7, 0xf0, 0x91, 0x15, 0x67, 0x73, 0xa1,
	0x1c, 0x93, 0xca, 0x3a, 0x93, 0xe2, 0x97, 0x5d, 0x09, 0xef, 0xa3, 0xff, 0x51, 0xce, 0x99, 0x54,
	0x94, 0xb2, 0x43, 0x00, 0x3d, 0x2b, 0x54, 0xcc, 0x8a, 0x64, 0xbf, 0xa7, 0x22, 0x15, 0x31, 0x7d,
	0x90, 0x39, 0x7a, 0xe8, 0x24, 0x43, 0x7e, 0x45, 0x80, 0x3c, 0x82, 0xa3, 0x35, 0x7c, 0xa6, 0x84,
	0x63, 0x52, 0xd3, 0xef, 0x50, 0xb6, 0x7f, 0x4d, 0xf6, 0x4a, 0xb8, 0x89, 0x26, 0x5f, 0x02, 0x36,
	0xcc, 0x34, 0xec, 0x34, 0x92, 0xe7, 0x22, 0xa6, 0x0f, 0x31, 0xe2, 0x8e, 0x07, 0x90, 0xfb, 0x12,
	0xcb, 0xe4, 0x3e, 0xd0, 0x25, 0x6e, 0xd6, 0xbe, 0x90, 0x3c, 0x42, 0x49, 0xbf, 0x94, 0x60, 0xf7,
	0x5c, 0xf7, 0x05, 0xec, 0x16, 0xd1, 0xca, 0x2d, 0x3c, 0xc6, 0x54, 0x3b, 0x79, 0xbd, 0x1c, 0xfd,
	0x1b, 0xe8, 0xe7, 0x94, 0xd5, 0x44, 0xdf, 0x63, 0x7b, 0x92, 0x63, 0xab, 0xa1, 0x0e, 0x8a, 0xa6,
	0x57, 0x17, 0xf6, 0x03, 0x7a, 0xec, 0x15, 0xf0, 0xea, 0xd2, 0x1e, 0xe7, 0x4b, 0x5b, 0x48, 0x15,
	0xeb, 0x05, 0xb3, 0x6f, 0x8d, 0x54, 0xef, 0x98, 0x3c, 0x53, 0xda, 0x88, 0x98, 0x3e, 0x41, 0xbf,
	0x03, 0xcf, 0x78, 0x8d, 0x84, 0x29, 0xe2, 0x93, 0x0c, 0x26, 0x43, 0xe8, 0x48, 0xcb, 0x92, 0x68,
	0xc1, 0xac, 0xf6, 0x4d, 0xe9, 0x8f, 0x78, 0x26, 0xb5, 0xa4, 0x3d, 0x89, 0x16, 0x53, 0x2c, 0x91,
	0x3f, 0x60, 0x1b, 0xdf, 0xde, 0xf2, 0xac, 0xa5, 0x4f, 0x07, 0xb5, 0x51, 0xeb, 0xee, 0x9b, 0xe0,
	0x66, 0x2e, 0x83, 0xe0, 0xda, 0x41, 0x1e, 0xb6, 0xbd, 0xe1, 0x24, 0x3b, 0xba, 0x2d, 0xf9, 0xb3,
	0x06, 0x3b, 0xd9, 0x77, 0x56, 0x45, 0xf8, 0xe9, 0x43, 0x47, 0xe8, 0xa0, 0x63, 0x99, 0x61, 0x00,
	0xed, 0x62, 0x09, 0x4e, 0xce, 0x05, 0x7d, 0x96, 0xdd, 0x11, 0x59, 0xce, 0xdf, 0xe4, 0x5c, 0xf8,
	0x55, 0x96, 0x21, 0x91, 0xf2, 0x1c, 0x29, 0xad, 0xbc, 0x0f, 0x72, 0xfe, 0xae, 0x41, 0x6f, 0xcd,
	0xf5, 0x40, 0x7f, 0xc6, 0x69, 0xf8, 0x4d, 0x4e, 0xf3, 0x3f, 0xb7, 0x50, 0xd8, 0x5d, 0xbe, 0x6e,
	0xa6, 0xbe, 0x3c, 0xdb, 0xc4, 0xbb, 0x7d, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x6d,
	0x67, 0x5b, 0xf9, 0x07, 0x00, 0x00,
}
