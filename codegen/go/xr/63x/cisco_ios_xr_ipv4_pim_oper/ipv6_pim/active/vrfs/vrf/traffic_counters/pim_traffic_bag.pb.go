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
// source: pim_traffic_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_traffic_counters

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

type PimTrafficBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimTrafficBag_KEYS) Reset()         { *m = PimTrafficBag_KEYS{} }
func (m *PimTrafficBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimTrafficBag_KEYS) ProtoMessage()    {}
func (*PimTrafficBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad93b554274fdcdc, []int{0}
}

func (m *PimTrafficBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTrafficBag_KEYS.Unmarshal(m, b)
}
func (m *PimTrafficBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTrafficBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimTrafficBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTrafficBag_KEYS.Merge(m, src)
}
func (m *PimTrafficBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimTrafficBag_KEYS.Size(m)
}
func (m *PimTrafficBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTrafficBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimTrafficBag_KEYS proto.InternalMessageInfo

func (m *PimTrafficBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type PimPktqStateT struct {
	MaxQueueSize         uint32   `protobuf:"varint,1,opt,name=max_queue_size,json=maxQueueSize,proto3" json:"max_queue_size,omitempty"`
	QueueSizeBytes       uint32   `protobuf:"varint,2,opt,name=queue_size_bytes,json=queueSizeBytes,proto3" json:"queue_size_bytes,omitempty"`
	QueueSizePackets     uint32   `protobuf:"varint,3,opt,name=queue_size_packets,json=queueSizePackets,proto3" json:"queue_size_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimPktqStateT) Reset()         { *m = PimPktqStateT{} }
func (m *PimPktqStateT) String() string { return proto.CompactTextString(m) }
func (*PimPktqStateT) ProtoMessage()    {}
func (*PimPktqStateT) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad93b554274fdcdc, []int{1}
}

func (m *PimPktqStateT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimPktqStateT.Unmarshal(m, b)
}
func (m *PimPktqStateT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimPktqStateT.Marshal(b, m, deterministic)
}
func (m *PimPktqStateT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimPktqStateT.Merge(m, src)
}
func (m *PimPktqStateT) XXX_Size() int {
	return xxx_messageInfo_PimPktqStateT.Size(m)
}
func (m *PimPktqStateT) XXX_DiscardUnknown() {
	xxx_messageInfo_PimPktqStateT.DiscardUnknown(m)
}

var xxx_messageInfo_PimPktqStateT proto.InternalMessageInfo

func (m *PimPktqStateT) GetMaxQueueSize() uint32 {
	if m != nil {
		return m.MaxQueueSize
	}
	return 0
}

func (m *PimPktqStateT) GetQueueSizeBytes() uint32 {
	if m != nil {
		return m.QueueSizeBytes
	}
	return 0
}

func (m *PimPktqStateT) GetQueueSizePackets() uint32 {
	if m != nil {
		return m.QueueSizePackets
	}
	return 0
}

type PimPktqStatsT struct {
	EnqueuedPackets      uint32   `protobuf:"varint,1,opt,name=enqueued_packets,json=enqueuedPackets,proto3" json:"enqueued_packets,omitempty"`
	DequeuedPackets      uint32   `protobuf:"varint,2,opt,name=dequeued_packets,json=dequeuedPackets,proto3" json:"dequeued_packets,omitempty"`
	HighWaterMarkPackets uint32   `protobuf:"varint,3,opt,name=high_water_mark_packets,json=highWaterMarkPackets,proto3" json:"high_water_mark_packets,omitempty"`
	HighWaterMarkBytes   uint32   `protobuf:"varint,4,opt,name=high_water_mark_bytes,json=highWaterMarkBytes,proto3" json:"high_water_mark_bytes,omitempty"`
	TailDrops            uint32   `protobuf:"varint,5,opt,name=tail_drops,json=tailDrops,proto3" json:"tail_drops,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimPktqStatsT) Reset()         { *m = PimPktqStatsT{} }
func (m *PimPktqStatsT) String() string { return proto.CompactTextString(m) }
func (*PimPktqStatsT) ProtoMessage()    {}
func (*PimPktqStatsT) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad93b554274fdcdc, []int{2}
}

func (m *PimPktqStatsT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimPktqStatsT.Unmarshal(m, b)
}
func (m *PimPktqStatsT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimPktqStatsT.Marshal(b, m, deterministic)
}
func (m *PimPktqStatsT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimPktqStatsT.Merge(m, src)
}
func (m *PimPktqStatsT) XXX_Size() int {
	return xxx_messageInfo_PimPktqStatsT.Size(m)
}
func (m *PimPktqStatsT) XXX_DiscardUnknown() {
	xxx_messageInfo_PimPktqStatsT.DiscardUnknown(m)
}

var xxx_messageInfo_PimPktqStatsT proto.InternalMessageInfo

func (m *PimPktqStatsT) GetEnqueuedPackets() uint32 {
	if m != nil {
		return m.EnqueuedPackets
	}
	return 0
}

func (m *PimPktqStatsT) GetDequeuedPackets() uint32 {
	if m != nil {
		return m.DequeuedPackets
	}
	return 0
}

func (m *PimPktqStatsT) GetHighWaterMarkPackets() uint32 {
	if m != nil {
		return m.HighWaterMarkPackets
	}
	return 0
}

func (m *PimPktqStatsT) GetHighWaterMarkBytes() uint32 {
	if m != nil {
		return m.HighWaterMarkBytes
	}
	return 0
}

func (m *PimPktqStatsT) GetTailDrops() uint32 {
	if m != nil {
		return m.TailDrops
	}
	return 0
}

type PimPktqT struct {
	PacketQueuePriority  uint32         `protobuf:"varint,1,opt,name=packet_queue_priority,json=packetQueuePriority,proto3" json:"packet_queue_priority,omitempty"`
	PacketQueueState     *PimPktqStateT `protobuf:"bytes,2,opt,name=packet_queue_state,json=packetQueueState,proto3" json:"packet_queue_state,omitempty"`
	PacketQueueStats     *PimPktqStatsT `protobuf:"bytes,3,opt,name=packet_queue_stats,json=packetQueueStats,proto3" json:"packet_queue_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PimPktqT) Reset()         { *m = PimPktqT{} }
func (m *PimPktqT) String() string { return proto.CompactTextString(m) }
func (*PimPktqT) ProtoMessage()    {}
func (*PimPktqT) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad93b554274fdcdc, []int{3}
}

func (m *PimPktqT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimPktqT.Unmarshal(m, b)
}
func (m *PimPktqT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimPktqT.Marshal(b, m, deterministic)
}
func (m *PimPktqT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimPktqT.Merge(m, src)
}
func (m *PimPktqT) XXX_Size() int {
	return xxx_messageInfo_PimPktqT.Size(m)
}
func (m *PimPktqT) XXX_DiscardUnknown() {
	xxx_messageInfo_PimPktqT.DiscardUnknown(m)
}

var xxx_messageInfo_PimPktqT proto.InternalMessageInfo

func (m *PimPktqT) GetPacketQueuePriority() uint32 {
	if m != nil {
		return m.PacketQueuePriority
	}
	return 0
}

func (m *PimPktqT) GetPacketQueueState() *PimPktqStateT {
	if m != nil {
		return m.PacketQueueState
	}
	return nil
}

func (m *PimPktqT) GetPacketQueueStats() *PimPktqStatsT {
	if m != nil {
		return m.PacketQueueStats
	}
	return nil
}

type PimTrafficBag struct {
	ElapsedTime                       uint32      `protobuf:"varint,50,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
	Inputs                            uint32      `protobuf:"varint,51,opt,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs                           uint32      `protobuf:"varint,52,opt,name=outputs,proto3" json:"outputs,omitempty"`
	FormatError                       uint32      `protobuf:"varint,53,opt,name=format_error,json=formatError,proto3" json:"format_error,omitempty"`
	PakmanError                       uint32      `protobuf:"varint,54,opt,name=pakman_error,json=pakmanError,proto3" json:"pakman_error,omitempty"`
	StandbyPacketsError               uint32      `protobuf:"varint,55,opt,name=standby_packets_error,json=standbyPacketsError,proto3" json:"standby_packets_error,omitempty"`
	ChecksumError                     uint32      `protobuf:"varint,56,opt,name=checksum_error,json=checksumError,proto3" json:"checksum_error,omitempty"`
	SocketError                       uint32      `protobuf:"varint,57,opt,name=socket_error,json=socketError,proto3" json:"socket_error,omitempty"`
	SendQueueFull                     uint32      `protobuf:"varint,58,opt,name=send_queue_full,json=sendQueueFull,proto3" json:"send_queue_full,omitempty"`
	BoundaryAclRxDrop                 uint32      `protobuf:"varint,59,opt,name=boundary_acl_rx_drop,json=boundaryAclRxDrop,proto3" json:"boundary_acl_rx_drop,omitempty"`
	BoundaryAclTxDrop                 uint32      `protobuf:"varint,60,opt,name=boundary_acl_tx_drop,json=boundaryAclTxDrop,proto3" json:"boundary_acl_tx_drop,omitempty"`
	NoSocketConnection                uint32      `protobuf:"varint,61,opt,name=no_socket_connection,json=noSocketConnection,proto3" json:"no_socket_connection,omitempty"`
	NoSourceAddress                   uint32      `protobuf:"varint,62,opt,name=no_source_address,json=noSourceAddress,proto3" json:"no_source_address,omitempty"`
	InputHello                        uint32      `protobuf:"varint,63,opt,name=input_hello,json=inputHello,proto3" json:"input_hello,omitempty"`
	OutputHello                       uint32      `protobuf:"varint,64,opt,name=output_hello,json=outputHello,proto3" json:"output_hello,omitempty"`
	InputJp                           uint32      `protobuf:"varint,65,opt,name=input_jp,json=inputJp,proto3" json:"input_jp,omitempty"`
	OutputJp                          uint32      `protobuf:"varint,66,opt,name=output_jp,json=outputJp,proto3" json:"output_jp,omitempty"`
	InputDataRegister                 uint32      `protobuf:"varint,67,opt,name=input_data_register,json=inputDataRegister,proto3" json:"input_data_register,omitempty"`
	InputNullRegister                 uint32      `protobuf:"varint,68,opt,name=input_null_register,json=inputNullRegister,proto3" json:"input_null_register,omitempty"`
	OutputNullRegister                uint32      `protobuf:"varint,69,opt,name=output_null_register,json=outputNullRegister,proto3" json:"output_null_register,omitempty"`
	InputRegisterStop                 uint32      `protobuf:"varint,70,opt,name=input_register_stop,json=inputRegisterStop,proto3" json:"input_register_stop,omitempty"`
	OutputRegisterStop                uint32      `protobuf:"varint,71,opt,name=output_register_stop,json=outputRegisterStop,proto3" json:"output_register_stop,omitempty"`
	InputAssert                       uint32      `protobuf:"varint,72,opt,name=input_assert,json=inputAssert,proto3" json:"input_assert,omitempty"`
	InputAssertBatched                uint32      `protobuf:"varint,73,opt,name=input_assert_batched,json=inputAssertBatched,proto3" json:"input_assert_batched,omitempty"`
	OutputAssert                      uint32      `protobuf:"varint,74,opt,name=output_assert,json=outputAssert,proto3" json:"output_assert,omitempty"`
	OutputAssertBatched               uint32      `protobuf:"varint,75,opt,name=output_assert_batched,json=outputAssertBatched,proto3" json:"output_assert_batched,omitempty"`
	InputDfElection                   uint32      `protobuf:"varint,76,opt,name=input_df_election,json=inputDfElection,proto3" json:"input_df_election,omitempty"`
	OutputDfElection                  uint32      `protobuf:"varint,77,opt,name=output_df_election,json=outputDfElection,proto3" json:"output_df_election,omitempty"`
	InputBsrMessage                   uint32      `protobuf:"varint,78,opt,name=input_bsr_message,json=inputBsrMessage,proto3" json:"input_bsr_message,omitempty"`
	OutputBsrMessage                  uint32      `protobuf:"varint,79,opt,name=output_bsr_message,json=outputBsrMessage,proto3" json:"output_bsr_message,omitempty"`
	InputCandidateRpAdvertisement     uint32      `protobuf:"varint,80,opt,name=input_candidate_rp_advertisement,json=inputCandidateRpAdvertisement,proto3" json:"input_candidate_rp_advertisement,omitempty"`
	OutputCandidateRpAdvertisement    uint32      `protobuf:"varint,81,opt,name=output_candidate_rp_advertisement,json=outputCandidateRpAdvertisement,proto3" json:"output_candidate_rp_advertisement,omitempty"`
	InputEcmpRedirect                 uint32      `protobuf:"varint,82,opt,name=input_ecmp_redirect,json=inputEcmpRedirect,proto3" json:"input_ecmp_redirect,omitempty"`
	OutputEcmpRedirect                uint32      `protobuf:"varint,83,opt,name=output_ecmp_redirect,json=outputEcmpRedirect,proto3" json:"output_ecmp_redirect,omitempty"`
	OutputLoopError                   uint32      `protobuf:"varint,84,opt,name=output_loop_error,json=outputLoopError,proto3" json:"output_loop_error,omitempty"`
	MldpMdtInvalidLsmIdentifier       uint32      `protobuf:"varint,85,opt,name=mldp_mdt_invalid_lsm_identifier,json=mldpMdtInvalidLsmIdentifier,proto3" json:"mldp_mdt_invalid_lsm_identifier,omitempty"`
	InputNoIdbError                   uint32      `protobuf:"varint,86,opt,name=input_no_idb_error,json=inputNoIdbError,proto3" json:"input_no_idb_error,omitempty"`
	InputNoVrfError                   uint32      `protobuf:"varint,87,opt,name=input_no_vrf_error,json=inputNoVrfError,proto3" json:"input_no_vrf_error,omitempty"`
	InputNoPimError                   uint32      `protobuf:"varint,88,opt,name=input_no_pim_error,json=inputNoPimError,proto3" json:"input_no_pim_error,omitempty"`
	InputPimVersionError              uint32      `protobuf:"varint,89,opt,name=input_pim_version_error,json=inputPimVersionError,proto3" json:"input_pim_version_error,omitempty"`
	OutputJoinGroup                   uint32      `protobuf:"varint,90,opt,name=output_join_group,json=outputJoinGroup,proto3" json:"output_join_group,omitempty"`
	OutputPruneGroup                  uint32      `protobuf:"varint,91,opt,name=output_prune_group,json=outputPruneGroup,proto3" json:"output_prune_group,omitempty"`
	OutputJoinPruneBytes              uint32      `protobuf:"varint,92,opt,name=output_join_prune_bytes,json=outputJoinPruneBytes,proto3" json:"output_join_prune_bytes,omitempty"`
	OutputHelloBytes                  uint32      `protobuf:"varint,93,opt,name=output_hello_bytes,json=outputHelloBytes,proto3" json:"output_hello_bytes,omitempty"`
	NonSupportedPackets               uint32      `protobuf:"varint,94,opt,name=non_supported_packets,json=nonSupportedPackets,proto3" json:"non_supported_packets,omitempty"`
	InvalidRegisters                  uint32      `protobuf:"varint,95,opt,name=invalid_registers,json=invalidRegisters,proto3" json:"invalid_registers,omitempty"`
	InvalidJoinPrunes                 uint32      `protobuf:"varint,96,opt,name=invalid_join_prunes,json=invalidJoinPrunes,proto3" json:"invalid_join_prunes,omitempty"`
	PacketPackmanError                uint32      `protobuf:"varint,97,opt,name=packet_packman_error,json=packetPackmanError,proto3" json:"packet_packman_error,omitempty"`
	PacketReadSocketError             uint32      `protobuf:"varint,98,opt,name=packet_read_socket_error,json=packetReadSocketError,proto3" json:"packet_read_socket_error,omitempty"`
	PacketQueue                       []*PimPktqT `protobuf:"bytes,99,rep,name=packet_queue,json=packetQueue,proto3" json:"packet_queue,omitempty"`
	PacketQueueLastClear              uint32      `protobuf:"varint,100,opt,name=packet_queue_last_clear,json=packetQueueLastClear,proto3" json:"packet_queue_last_clear,omitempty"`
	PacketsStandby                    uint32      `protobuf:"varint,101,opt,name=packets_standby,json=packetsStandby,proto3" json:"packets_standby,omitempty"`
	NoMdtSocketConnection             uint32      `protobuf:"varint,102,opt,name=no_mdt_socket_connection,json=noMdtSocketConnection,proto3" json:"no_mdt_socket_connection,omitempty"`
	MdtSendQueueFull                  uint32      `protobuf:"varint,103,opt,name=mdt_send_queue_full,json=mdtSendQueueFull,proto3" json:"mdt_send_queue_full,omitempty"`
	MdtSocketError                    uint32      `protobuf:"varint,104,opt,name=mdt_socket_error,json=mdtSocketError,proto3" json:"mdt_socket_error,omitempty"`
	MdtJoinTlvSent                    uint32      `protobuf:"varint,105,opt,name=mdt_join_tlv_sent,json=mdtJoinTlvSent,proto3" json:"mdt_join_tlv_sent,omitempty"`
	MdtJoinTlvReceived                uint32      `protobuf:"varint,106,opt,name=mdt_join_tlv_received,json=mdtJoinTlvReceived,proto3" json:"mdt_join_tlv_received,omitempty"`
	MdtJoinBadType                    uint32      `protobuf:"varint,107,opt,name=mdt_join_bad_type,json=mdtJoinBadType,proto3" json:"mdt_join_bad_type,omitempty"`
	MdtDropLocalSourceAddress         uint32      `protobuf:"varint,108,opt,name=mdt_drop_local_source_address,json=mdtDropLocalSourceAddress,proto3" json:"mdt_drop_local_source_address,omitempty"`
	MdtDropNullLocalAddress           uint32      `protobuf:"varint,109,opt,name=mdt_drop_null_local_address,json=mdtDropNullLocalAddress,proto3" json:"mdt_drop_null_local_address,omitempty"`
	MdtDropNoIdb                      uint32      `protobuf:"varint,110,opt,name=mdt_drop_no_idb,json=mdtDropNoIdb,proto3" json:"mdt_drop_no_idb,omitempty"`
	MdtDropNoVrf                      uint32      `protobuf:"varint,111,opt,name=mdt_drop_no_vrf,json=mdtDropNoVrf,proto3" json:"mdt_drop_no_vrf,omitempty"`
	InvalidDestinationPackets         uint32      `protobuf:"varint,112,opt,name=invalid_destination_packets,json=invalidDestinationPackets,proto3" json:"invalid_destination_packets,omitempty"`
	MdtJoinsDropMultipleEncapsulation uint32      `protobuf:"varint,113,opt,name=mdt_joins_drop_multiple_encapsulation,json=mdtJoinsDropMultipleEncapsulation,proto3" json:"mdt_joins_drop_multiple_encapsulation,omitempty"`
	TruncatedPimPackets               uint32      `protobuf:"varint,114,opt,name=truncated_pim_packets,json=truncatedPimPackets,proto3" json:"truncated_pim_packets,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}    `json:"-"`
	XXX_unrecognized                  []byte      `json:"-"`
	XXX_sizecache                     int32       `json:"-"`
}

func (m *PimTrafficBag) Reset()         { *m = PimTrafficBag{} }
func (m *PimTrafficBag) String() string { return proto.CompactTextString(m) }
func (*PimTrafficBag) ProtoMessage()    {}
func (*PimTrafficBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad93b554274fdcdc, []int{4}
}

func (m *PimTrafficBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimTrafficBag.Unmarshal(m, b)
}
func (m *PimTrafficBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimTrafficBag.Marshal(b, m, deterministic)
}
func (m *PimTrafficBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimTrafficBag.Merge(m, src)
}
func (m *PimTrafficBag) XXX_Size() int {
	return xxx_messageInfo_PimTrafficBag.Size(m)
}
func (m *PimTrafficBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimTrafficBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimTrafficBag proto.InternalMessageInfo

func (m *PimTrafficBag) GetElapsedTime() uint32 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func (m *PimTrafficBag) GetInputs() uint32 {
	if m != nil {
		return m.Inputs
	}
	return 0
}

func (m *PimTrafficBag) GetOutputs() uint32 {
	if m != nil {
		return m.Outputs
	}
	return 0
}

func (m *PimTrafficBag) GetFormatError() uint32 {
	if m != nil {
		return m.FormatError
	}
	return 0
}

func (m *PimTrafficBag) GetPakmanError() uint32 {
	if m != nil {
		return m.PakmanError
	}
	return 0
}

func (m *PimTrafficBag) GetStandbyPacketsError() uint32 {
	if m != nil {
		return m.StandbyPacketsError
	}
	return 0
}

func (m *PimTrafficBag) GetChecksumError() uint32 {
	if m != nil {
		return m.ChecksumError
	}
	return 0
}

func (m *PimTrafficBag) GetSocketError() uint32 {
	if m != nil {
		return m.SocketError
	}
	return 0
}

func (m *PimTrafficBag) GetSendQueueFull() uint32 {
	if m != nil {
		return m.SendQueueFull
	}
	return 0
}

func (m *PimTrafficBag) GetBoundaryAclRxDrop() uint32 {
	if m != nil {
		return m.BoundaryAclRxDrop
	}
	return 0
}

func (m *PimTrafficBag) GetBoundaryAclTxDrop() uint32 {
	if m != nil {
		return m.BoundaryAclTxDrop
	}
	return 0
}

func (m *PimTrafficBag) GetNoSocketConnection() uint32 {
	if m != nil {
		return m.NoSocketConnection
	}
	return 0
}

func (m *PimTrafficBag) GetNoSourceAddress() uint32 {
	if m != nil {
		return m.NoSourceAddress
	}
	return 0
}

func (m *PimTrafficBag) GetInputHello() uint32 {
	if m != nil {
		return m.InputHello
	}
	return 0
}

func (m *PimTrafficBag) GetOutputHello() uint32 {
	if m != nil {
		return m.OutputHello
	}
	return 0
}

func (m *PimTrafficBag) GetInputJp() uint32 {
	if m != nil {
		return m.InputJp
	}
	return 0
}

func (m *PimTrafficBag) GetOutputJp() uint32 {
	if m != nil {
		return m.OutputJp
	}
	return 0
}

func (m *PimTrafficBag) GetInputDataRegister() uint32 {
	if m != nil {
		return m.InputDataRegister
	}
	return 0
}

func (m *PimTrafficBag) GetInputNullRegister() uint32 {
	if m != nil {
		return m.InputNullRegister
	}
	return 0
}

func (m *PimTrafficBag) GetOutputNullRegister() uint32 {
	if m != nil {
		return m.OutputNullRegister
	}
	return 0
}

func (m *PimTrafficBag) GetInputRegisterStop() uint32 {
	if m != nil {
		return m.InputRegisterStop
	}
	return 0
}

func (m *PimTrafficBag) GetOutputRegisterStop() uint32 {
	if m != nil {
		return m.OutputRegisterStop
	}
	return 0
}

func (m *PimTrafficBag) GetInputAssert() uint32 {
	if m != nil {
		return m.InputAssert
	}
	return 0
}

func (m *PimTrafficBag) GetInputAssertBatched() uint32 {
	if m != nil {
		return m.InputAssertBatched
	}
	return 0
}

func (m *PimTrafficBag) GetOutputAssert() uint32 {
	if m != nil {
		return m.OutputAssert
	}
	return 0
}

func (m *PimTrafficBag) GetOutputAssertBatched() uint32 {
	if m != nil {
		return m.OutputAssertBatched
	}
	return 0
}

func (m *PimTrafficBag) GetInputDfElection() uint32 {
	if m != nil {
		return m.InputDfElection
	}
	return 0
}

func (m *PimTrafficBag) GetOutputDfElection() uint32 {
	if m != nil {
		return m.OutputDfElection
	}
	return 0
}

func (m *PimTrafficBag) GetInputBsrMessage() uint32 {
	if m != nil {
		return m.InputBsrMessage
	}
	return 0
}

func (m *PimTrafficBag) GetOutputBsrMessage() uint32 {
	if m != nil {
		return m.OutputBsrMessage
	}
	return 0
}

func (m *PimTrafficBag) GetInputCandidateRpAdvertisement() uint32 {
	if m != nil {
		return m.InputCandidateRpAdvertisement
	}
	return 0
}

func (m *PimTrafficBag) GetOutputCandidateRpAdvertisement() uint32 {
	if m != nil {
		return m.OutputCandidateRpAdvertisement
	}
	return 0
}

func (m *PimTrafficBag) GetInputEcmpRedirect() uint32 {
	if m != nil {
		return m.InputEcmpRedirect
	}
	return 0
}

func (m *PimTrafficBag) GetOutputEcmpRedirect() uint32 {
	if m != nil {
		return m.OutputEcmpRedirect
	}
	return 0
}

func (m *PimTrafficBag) GetOutputLoopError() uint32 {
	if m != nil {
		return m.OutputLoopError
	}
	return 0
}

func (m *PimTrafficBag) GetMldpMdtInvalidLsmIdentifier() uint32 {
	if m != nil {
		return m.MldpMdtInvalidLsmIdentifier
	}
	return 0
}

func (m *PimTrafficBag) GetInputNoIdbError() uint32 {
	if m != nil {
		return m.InputNoIdbError
	}
	return 0
}

func (m *PimTrafficBag) GetInputNoVrfError() uint32 {
	if m != nil {
		return m.InputNoVrfError
	}
	return 0
}

func (m *PimTrafficBag) GetInputNoPimError() uint32 {
	if m != nil {
		return m.InputNoPimError
	}
	return 0
}

func (m *PimTrafficBag) GetInputPimVersionError() uint32 {
	if m != nil {
		return m.InputPimVersionError
	}
	return 0
}

func (m *PimTrafficBag) GetOutputJoinGroup() uint32 {
	if m != nil {
		return m.OutputJoinGroup
	}
	return 0
}

func (m *PimTrafficBag) GetOutputPruneGroup() uint32 {
	if m != nil {
		return m.OutputPruneGroup
	}
	return 0
}

func (m *PimTrafficBag) GetOutputJoinPruneBytes() uint32 {
	if m != nil {
		return m.OutputJoinPruneBytes
	}
	return 0
}

func (m *PimTrafficBag) GetOutputHelloBytes() uint32 {
	if m != nil {
		return m.OutputHelloBytes
	}
	return 0
}

func (m *PimTrafficBag) GetNonSupportedPackets() uint32 {
	if m != nil {
		return m.NonSupportedPackets
	}
	return 0
}

func (m *PimTrafficBag) GetInvalidRegisters() uint32 {
	if m != nil {
		return m.InvalidRegisters
	}
	return 0
}

func (m *PimTrafficBag) GetInvalidJoinPrunes() uint32 {
	if m != nil {
		return m.InvalidJoinPrunes
	}
	return 0
}

func (m *PimTrafficBag) GetPacketPackmanError() uint32 {
	if m != nil {
		return m.PacketPackmanError
	}
	return 0
}

func (m *PimTrafficBag) GetPacketReadSocketError() uint32 {
	if m != nil {
		return m.PacketReadSocketError
	}
	return 0
}

func (m *PimTrafficBag) GetPacketQueue() []*PimPktqT {
	if m != nil {
		return m.PacketQueue
	}
	return nil
}

func (m *PimTrafficBag) GetPacketQueueLastClear() uint32 {
	if m != nil {
		return m.PacketQueueLastClear
	}
	return 0
}

func (m *PimTrafficBag) GetPacketsStandby() uint32 {
	if m != nil {
		return m.PacketsStandby
	}
	return 0
}

func (m *PimTrafficBag) GetNoMdtSocketConnection() uint32 {
	if m != nil {
		return m.NoMdtSocketConnection
	}
	return 0
}

func (m *PimTrafficBag) GetMdtSendQueueFull() uint32 {
	if m != nil {
		return m.MdtSendQueueFull
	}
	return 0
}

func (m *PimTrafficBag) GetMdtSocketError() uint32 {
	if m != nil {
		return m.MdtSocketError
	}
	return 0
}

func (m *PimTrafficBag) GetMdtJoinTlvSent() uint32 {
	if m != nil {
		return m.MdtJoinTlvSent
	}
	return 0
}

func (m *PimTrafficBag) GetMdtJoinTlvReceived() uint32 {
	if m != nil {
		return m.MdtJoinTlvReceived
	}
	return 0
}

func (m *PimTrafficBag) GetMdtJoinBadType() uint32 {
	if m != nil {
		return m.MdtJoinBadType
	}
	return 0
}

func (m *PimTrafficBag) GetMdtDropLocalSourceAddress() uint32 {
	if m != nil {
		return m.MdtDropLocalSourceAddress
	}
	return 0
}

func (m *PimTrafficBag) GetMdtDropNullLocalAddress() uint32 {
	if m != nil {
		return m.MdtDropNullLocalAddress
	}
	return 0
}

func (m *PimTrafficBag) GetMdtDropNoIdb() uint32 {
	if m != nil {
		return m.MdtDropNoIdb
	}
	return 0
}

func (m *PimTrafficBag) GetMdtDropNoVrf() uint32 {
	if m != nil {
		return m.MdtDropNoVrf
	}
	return 0
}

func (m *PimTrafficBag) GetInvalidDestinationPackets() uint32 {
	if m != nil {
		return m.InvalidDestinationPackets
	}
	return 0
}

func (m *PimTrafficBag) GetMdtJoinsDropMultipleEncapsulation() uint32 {
	if m != nil {
		return m.MdtJoinsDropMultipleEncapsulation
	}
	return 0
}

func (m *PimTrafficBag) GetTruncatedPimPackets() uint32 {
	if m != nil {
		return m.TruncatedPimPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*PimTrafficBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.traffic_counters.pim_traffic_bag_KEYS")
	proto.RegisterType((*PimPktqStateT)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.traffic_counters.pim_pktq_state_t")
	proto.RegisterType((*PimPktqStatsT)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.traffic_counters.pim_pktq_stats_t")
	proto.RegisterType((*PimPktqT)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.traffic_counters.pim_pktq_t")
	proto.RegisterType((*PimTrafficBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.traffic_counters.pim_traffic_bag")
}

func init() { proto.RegisterFile("pim_traffic_bag.proto", fileDescriptor_ad93b554274fdcdc) }

var fileDescriptor_ad93b554274fdcdc = []byte{
	// 1585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x97, 0xdd, 0x76, 0x14, 0xb9,
	0x11, 0xc7, 0x8f, 0x21, 0xe1, 0x43, 0x06, 0x6c, 0x37, 0x76, 0x2c, 0x8e, 0x0f, 0x01, 0x3b, 0x21,
	0x31, 0x90, 0x4c, 0xc0, 0x60, 0xc8, 0x07, 0x21, 0xf8, 0x0b, 0xb0, 0xb1, 0x9d, 0x61, 0xc6, 0x31,
	0x21, 0x5f, 0x8a, 0xa6, 0xa5, 0xb1, 0x65, 0x77, 0xb7, 0x64, 0x49, 0x3d, 0xb1, 0xb9, 0xce, 0x03,
	0xec, 0x6b, 0xed, 0xcd, 0xbe, 0xcb, 0xbe, 0xc1, 0x1e, 0xa9, 0xd4, 0x3d, 0xea, 0x31, 0xbb, 0x57,
	0xec, 0x0d, 0x9c, 0xa9, 0xfa, 0xfd, 0x4b, 0xea, 0xaa, 0x92, 0x54, 0x46, 0x33, 0x4a, 0xe4, 0xc4,
	0x6a, 0xda, 0xef, 0x8b, 0x94, 0xf4, 0xe8, 0x41, 0x4b, 0x69, 0x69, 0x65, 0xb2, 0x9e, 0x0a, 0x93,
	0x4a, 0x22, 0xa4, 0x21, 0xa7, 0x9a, 0x08, 0x35, 0x78, 0x4a, 0x1c, 0x28, 0x15, 0xd7, 0x2d, 0xa1,
	0x06, 0xcf, 0xdc, 0xaf, 0x16, 0x4d, 0xad, 0x18, 0xf0, 0xd6, 0x40, 0xf7, 0x8d, 0xfb, 0xa7, 0x55,
	0x85, 0x49, 0x65, 0x59, 0x58, 0xae, 0xcd, 0xc2, 0x63, 0x34, 0x3d, 0x12, 0x9e, 0xbc, 0xdb, 0xf8,
	0xd8, 0x4d, 0x6e, 0xa1, 0x2b, 0x03, 0xdd, 0x27, 0x05, 0xcd, 0x39, 0x1e, 0xbb, 0x3b, 0xb6, 0x78,
	0xb5, 0x73, 0x79, 0xa0, 0xfb, 0xbb, 0x34, 0xe7, 0x0b, 0x5f, 0x8d, 0xa1, 0x49, 0xa7, 0x51, 0xc7,
	0xf6, 0x84, 0x18, 0x4b, 0x2d, 0x27, 0x36, 0xf9, 0x25, 0xba, 0x91, 0xd3, 0x53, 0x72, 0x52, 0xf2,
	0x92, 0x13, 0x23, 0x3e, 0x81, 0xea, 0x7a, 0xe7, 0x5a, 0x4e, 0x4f, 0xdf, 0x3b, 0x63, 0x57, 0x7c,
	0xe2, 0xc9, 0x22, 0x9a, 0x1c, 0x12, 0xa4, 0x77, 0x66, 0xb9, 0xc1, 0x17, 0x3c, 0x77, 0xe3, 0xa4,
	0x82, 0x56, 0x9d, 0x35, 0xf9, 0x0d, 0x4a, 0x22, 0x52, 0xd1, 0xf4, 0x98, 0x5b, 0x83, 0x2f, 0x7a,
	0x76, 0xb2, 0x66, 0xdb, 0x60, 0x5f, 0xf8, 0x76, 0x74, 0x4b, 0x86, 0xd8, 0xe4, 0x3e, 0x9a, 0xe4,
	0x85, 0x47, 0x59, 0x1d, 0x00, 0x36, 0x35, 0x51, 0xd9, 0x83, 0xde, 0xa1, 0x8c, 0x8f, 0xa0, 0xb0,
	0xaf, 0x89, 0xca, 0x5e, 0xa1, 0xcb, 0x68, 0xf6, 0x50, 0x1c, 0x1c, 0x92, 0xff, 0x51, 0xcb, 0x35,
	0xc9, 0xa9, 0x3e, 0x1e, 0xd9, 0xdd, 0xb4, 0x73, 0x7f, 0x70, 0xde, 0x1d, 0xaa, 0x8f, 0x2b, 0xd9,
	0x63, 0x34, 0x33, 0x2a, 0x83, 0xcf, 0xff, 0x89, 0x17, 0x25, 0x0d, 0x11, 0xa4, 0xe0, 0x36, 0x42,
	0x96, 0x8a, 0x8c, 0x30, 0x2d, 0x95, 0xc1, 0x3f, 0xf5, 0xdc, 0x55, 0x67, 0x59, 0x77, 0x86, 0x85,
	0x6f, 0x2e, 0x20, 0x54, 0x7f, 0xb3, 0x4d, 0x96, 0xd0, 0x0c, 0xec, 0x23, 0xd4, 0x40, 0x69, 0x21,
	0xb5, 0xb0, 0x67, 0xe1, 0x93, 0x6f, 0x82, 0xd3, 0x97, 0xa2, 0x1d, 0x5c, 0xc9, 0xff, 0xc7, 0x50,
	0xd2, 0x10, 0xf9, 0x6a, 0xfa, 0x2f, 0x1f, 0x5f, 0xda, 0x6f, 0x7d, 0x89, 0x06, 0x6b, 0x8d, 0x76,
	0x4a, 0x67, 0x32, 0xda, 0x49, 0xd7, 0xd9, 0x3e, 0xbf, 0x0d, 0x48, 0xe7, 0x8f, 0xb3, 0x0d, 0xf3,
	0x99, 0x6d, 0x98, 0x85, 0xaf, 0x31, 0x9a, 0x18, 0x39, 0x0b, 0xc9, 0x3c, 0xba, 0xc6, 0x33, 0xaa,
	0x0c, 0x67, 0xc4, 0x8a, 0x9c, 0xe3, 0x25, 0x9f, 0xcc, 0xf1, 0x60, 0xdb, 0x13, 0x39, 0x4f, 0x7e,
	0x86, 0x2e, 0x89, 0x42, 0x95, 0xd6, 0xe0, 0x27, 0xde, 0x19, 0x7e, 0x25, 0x18, 0x5d, 0x96, 0xa5,
	0xf5, 0x8e, 0xa7, 0xde, 0x51, 0xfd, 0x74, 0x41, 0xfb, 0x52, 0xe7, 0xd4, 0x12, 0xae, 0xb5, 0xd4,
	0x78, 0x19, 0x82, 0x82, 0x6d, 0xc3, 0x99, 0x1c, 0xa2, 0xe8, 0x71, 0x4e, 0x8b, 0x80, 0x3c, 0x03,
	0x04, 0x6c, 0x80, 0x2c, 0xa1, 0x19, 0x63, 0x69, 0xc1, 0x7a, 0x67, 0x55, 0x03, 0x06, 0xf6, 0x39,
	0x14, 0x3c, 0x38, 0x43, 0x03, 0x82, 0xe6, 0x1e, 0xba, 0x91, 0x1e, 0xf2, 0xf4, 0xd8, 0x94, 0x79,
	0x80, 0x7f, 0xef, 0xe1, 0xeb, 0x95, 0xb5, 0x5e, 0xdd, 0x48, 0x5f, 0x0f, 0x80, 0xfe, 0x00, 0xab,
	0x83, 0x0d, 0x90, 0x5f, 0xa1, 0x09, 0xc3, 0x0b, 0x16, 0x0a, 0xd6, 0x2f, 0xb3, 0x0c, 0xff, 0x11,
	0x42, 0x39, 0xb3, 0xcf, 0xea, 0xeb, 0x32, 0xcb, 0x92, 0xdf, 0xa1, 0xe9, 0x9e, 0x2c, 0x0b, 0x46,
	0xf5, 0x19, 0xa1, 0x69, 0x46, 0xf4, 0xa9, 0xef, 0x67, 0xfc, 0x27, 0x0f, 0x4f, 0x55, 0xbe, 0x95,
	0x34, 0xeb, 0x9c, 0xba, 0xbe, 0x3e, 0x27, 0xb0, 0x41, 0xf0, 0xe2, 0x9c, 0x60, 0x0f, 0x04, 0x8f,
	0xd0, 0x74, 0x21, 0x49, 0xd8, 0x6f, 0x2a, 0x8b, 0x82, 0xa7, 0x56, 0xc8, 0x02, 0xff, 0x19, 0x0e,
	0x56, 0x21, 0xbb, 0xde, 0xb5, 0x56, 0x7b, 0x92, 0x07, 0x68, 0xca, 0x2b, 0x4a, 0x9d, 0x72, 0x42,
	0x19, 0xd3, 0xdc, 0x18, 0xfc, 0x12, 0x8e, 0xbb, 0xc3, 0x9d, 0x7d, 0x05, 0xcc, 0xc9, 0x1d, 0x34,
	0xee, 0xeb, 0x49, 0x0e, 0x79, 0x96, 0x49, 0xfc, 0x17, 0x4f, 0x21, 0x6f, 0x7a, 0xeb, 0x2c, 0x2e,
	0x57, 0x50, 0xd7, 0x40, 0xbc, 0x82, 0x5c, 0x81, 0x0d, 0x90, 0x5b, 0xe8, 0x0a, 0xc4, 0x38, 0x52,
	0x78, 0x05, 0x5a, 0xc1, 0xff, 0xde, 0x52, 0xc9, 0x1c, 0xba, 0x1a, 0xd4, 0x47, 0x0a, 0xaf, 0x7a,
	0xdf, 0x15, 0x30, 0x6c, 0xa9, 0xa4, 0x85, 0x6e, 0x82, 0x8e, 0x51, 0x4b, 0x89, 0xe6, 0x07, 0xc2,
	0x58, 0xae, 0xf1, 0x1a, 0x64, 0xc2, 0xbb, 0xd6, 0xa9, 0xa5, 0x9d, 0xe0, 0x18, 0xf2, 0x45, 0x99,
	0x65, 0x43, 0x7e, 0x3d, 0xe2, 0x77, 0xcb, 0x2c, 0xab, 0xf9, 0x47, 0x68, 0x3a, 0x2c, 0xde, 0x14,
	0x6c, 0x40, 0xe6, 0xc0, 0xd7, 0x50, 0xd4, 0x2b, 0x54, 0x2c, 0x31, 0x56, 0x2a, 0xfc, 0x3a, 0x5a,
	0xa1, 0x62, 0xbb, 0x16, 0x6a, 0x13, 0x56, 0x68, 0x0a, 0xde, 0xc4, 0x2b, 0x34, 0x14, 0xf3, 0xe8,
	0x1a, 0xac, 0x40, 0x8d, 0xe1, 0xda, 0xe2, 0xb7, 0x90, 0x4e, 0x6f, 0x5b, 0xf1, 0x26, 0x17, 0x34,
	0x46, 0x48, 0x8f, 0xda, 0xf4, 0x90, 0x33, 0xbc, 0x09, 0x41, 0x23, 0x74, 0x15, 0x3c, 0xc9, 0x2f,
	0xd0, 0xf5, 0xb0, 0x8d, 0x10, 0x75, 0x0b, 0xde, 0x26, 0x30, 0x86, 0xb0, 0x4b, 0x68, 0xa6, 0x01,
	0xd5, 0x71, 0xdf, 0xc1, 0x79, 0x8a, 0xe1, 0x2a, 0xf0, 0x03, 0x34, 0x15, 0x2a, 0xd4, 0x27, 0x3c,
	0x0b, 0x8d, 0xb7, 0x0d, 0x9d, 0x04, 0xf5, 0xe9, 0x6f, 0x04, 0xb3, 0x7b, 0xd1, 0x42, 0xfc, 0x18,
	0xde, 0x81, 0x17, 0x0d, 0x3c, 0x11, 0x5d, 0x47, 0xee, 0x19, 0x4d, 0x72, 0x6e, 0x0c, 0x3d, 0xe0,
	0x78, 0x37, 0x8a, 0xbc, 0x6a, 0xf4, 0x0e, 0x98, 0xa3, 0xc8, 0x31, 0xfc, 0xd7, 0x38, 0x72, 0x44,
	0xbf, 0x41, 0x77, 0x21, 0x72, 0x4a, 0x0b, 0x26, 0x98, 0xbb, 0x94, 0xb5, 0x22, 0x94, 0x0d, 0xb8,
	0xb6, 0xc2, 0xf0, 0x9c, 0x17, 0x16, 0xb7, 0xbd, 0xf6, 0xb6, 0xe7, 0xd6, 0x2a, 0xac, 0xa3, 0x56,
	0x62, 0x28, 0xd9, 0x44, 0xf3, 0x61, 0xd9, 0x1f, 0x88, 0xf4, 0xde, 0x47, 0xfa, 0x39, 0x80, 0xdf,
	0x1b, 0xaa, 0xee, 0x2b, 0x9e, 0xe6, 0x8a, 0x68, 0xce, 0x84, 0xe6, 0xa9, 0xc5, 0x9d, 0xa8, 0xaf,
	0x36, 0xd2, 0x5c, 0x75, 0x82, 0x23, 0xea, 0xab, 0xa6, 0xa0, 0x1b, 0xf7, 0x55, 0x43, 0xf1, 0x00,
	0x4d, 0x05, 0x45, 0x26, 0xa5, 0x0a, 0xf7, 0xda, 0x1e, 0xe4, 0x13, 0x1c, 0xdb, 0x52, 0x2a, 0xb8,
	0xdb, 0xd6, 0xd1, 0x9d, 0x3c, 0x63, 0x8a, 0xe4, 0xcc, 0x12, 0x51, 0x0c, 0x68, 0x26, 0x18, 0xc9,
	0x4c, 0x4e, 0x04, 0xe3, 0x85, 0x15, 0x7d, 0xc1, 0x35, 0xfe, 0x9b, 0x57, 0xce, 0x39, 0x6c, 0x87,
	0xd9, 0x4d, 0x80, 0xb6, 0x4d, 0xbe, 0x59, 0x23, 0xc9, 0x43, 0x94, 0x84, 0xd3, 0x28, 0x89, 0x60,
	0xbd, 0xb0, 0xe4, 0x7e, 0x54, 0xc2, 0x5d, 0xb9, 0xc9, 0x7a, 0xb0, 0x64, 0x0c, 0xbb, 0xb9, 0x0b,
	0xe0, 0x0f, 0x0d, 0x78, 0x5f, 0xf7, 0xcf, 0xc3, 0xee, 0xc1, 0x02, 0xf8, 0xef, 0x0d, 0xb8, 0x2d,
	0xc2, 0x5d, 0xbe, 0x8c, 0x66, 0x01, 0x76, 0xe4, 0x80, 0x6b, 0x23, 0x64, 0xf5, 0xa8, 0x7c, 0x84,
	0x79, 0xc5, 0xbb, 0xdb, 0x22, 0xdf, 0x07, 0x27, 0xc8, 0x86, 0xf9, 0x3a, 0x92, 0xa2, 0x20, 0x07,
	0x5a, 0x96, 0x0a, 0xff, 0x23, 0xce, 0xd7, 0x96, 0x14, 0xc5, 0x1b, 0x67, 0x8e, 0xfa, 0x4f, 0xe9,
	0xb2, 0xe0, 0x01, 0xfe, 0x67, 0xdc, 0x7f, 0x6d, 0xe7, 0x00, 0x7a, 0x19, 0xcd, 0xc6, 0x91, 0x41,
	0x02, 0xb3, 0xd0, 0xbf, 0x60, 0x43, 0xc3, 0xf8, 0x5e, 0x56, 0x0f, 0x84, 0xf1, 0x3d, 0x1b, 0x14,
	0xff, 0x8e, 0x17, 0xf1, 0xb7, 0x2d, 0xd0, 0x4b, 0x68, 0xa6, 0x90, 0x05, 0x31, 0xa5, 0x52, 0x52,
	0xdb, 0x68, 0xaa, 0xfb, 0x0f, 0x1c, 0xe6, 0x42, 0x16, 0xdd, 0xca, 0x57, 0x8d, 0x68, 0x0f, 0xdd,
	0x91, 0x83, 0x6a, 0x57, 0xb7, 0x95, 0xc1, 0x04, 0x16, 0x08, 0x8e, 0xea, 0xaa, 0x32, 0xd0, 0xb1,
	0x00, 0x0f, 0x3f, 0xc3, 0xe0, 0xff, 0x56, 0x1d, 0xeb, 0x5d, 0xf5, 0x27, 0x18, 0xd7, 0xb1, 0x61,
	0xc4, 0x71, 0xff, 0x0d, 0x1f, 0x76, 0x0a, 0x1d, 0x0b, 0xbe, 0x36, 0xb8, 0xa0, 0x02, 0xcf, 0x11,
	0x0e, 0x0a, 0xcd, 0x29, 0x23, 0x8d, 0x07, 0xb9, 0xe7, 0x55, 0x61, 0xe0, 0xeb, 0x70, 0xca, 0xba,
	0xd1, 0xd3, 0x6c, 0xdc, 0xec, 0x30, 0x9c, 0xa6, 0x70, 0x7a, 0xf7, 0xe2, 0xe2, 0xf8, 0x52, 0xfb,
	0x0b, 0xcf, 0x51, 0xd6, 0x4d, 0x23, 0xf5, 0x04, 0xe5, 0xaa, 0xda, 0x18, 0xe1, 0x32, 0x6a, 0x2c,
	0x49, 0x33, 0x4e, 0x35, 0x66, 0x50, 0xd5, 0x88, 0xde, 0xa6, 0xc6, 0xae, 0x39, 0x5f, 0xf2, 0x6b,
	0x34, 0x51, 0x0d, 0x2f, 0x61, 0x5e, 0xc1, 0x1c, 0xfe, 0x1e, 0x08, 0xe6, 0x2e, 0x58, 0x5d, 0x36,
	0x0a, 0xe9, 0x4f, 0xe4, 0xf9, 0x97, 0xbe, 0x0f, 0xd9, 0x28, 0xe4, 0x0e, 0xb3, 0xe7, 0x1e, 0xfb,
	0xdf, 0xa2, 0x9b, 0x5e, 0x35, 0x32, 0xac, 0x1c, 0x40, 0x5d, 0x73, 0x66, 0xbb, 0x8d, 0x79, 0x65,
	0x11, 0x4d, 0x46, 0x8b, 0x40, 0xb6, 0x0f, 0x61, 0x47, 0x79, 0x15, 0x1d, 0xd2, 0x7c, 0x1f, 0x4d,
	0x39, 0xd2, 0x57, 0xdf, 0x66, 0x03, 0xb7, 0x82, 0xc5, 0xa2, 0x46, 0x5d, 0xed, 0xf7, 0xb2, 0x41,
	0xd7, 0x5d, 0x6f, 0x8f, 0xd1, 0x4c, 0x03, 0xd5, 0x3c, 0xe5, 0x62, 0xc0, 0x19, 0x3e, 0x82, 0xea,
	0x0f, 0xf1, 0x4e, 0xf0, 0x34, 0xa2, 0xf7, 0x28, 0x23, 0xf6, 0x4c, 0x71, 0x7c, 0xdc, 0x88, 0xbe,
	0x4a, 0xd9, 0xde, 0x99, 0xe2, 0xc9, 0x2b, 0x74, 0xdb, 0xa1, 0x6e, 0x4a, 0x22, 0x99, 0x4c, 0x69,
	0x36, 0x3a, 0xda, 0x64, 0x5e, 0x76, 0x2b, 0x67, 0xd6, 0x0d, 0x4c, 0xdb, 0x0e, 0x69, 0x0e, 0x39,
	0x2f, 0xd0, 0x5c, 0x1d, 0xc1, 0x8f, 0x02, 0x10, 0xa6, 0xd2, 0xe7, 0x5e, 0x3f, 0x1b, 0xf4, 0x6e,
	0x20, 0xf0, 0x31, 0x2a, 0xf5, 0x3d, 0x34, 0x31, 0x54, 0xfb, 0xbb, 0x0e, 0x17, 0xe1, 0x6f, 0xbf,
	0xa0, 0x70, 0xf7, 0xdc, 0x28, 0x36, 0xd0, 0x7d, 0x2c, 0x47, 0xb0, 0x7d, 0xdd, 0x4f, 0x5e, 0xa2,
	0xb9, 0xea, 0x60, 0x31, 0x6e, 0xac, 0x28, 0xa8, 0x2b, 0x63, 0x7d, 0x7e, 0x15, 0x7c, 0x4b, 0x40,
	0xd6, 0x87, 0x44, 0x75, 0x8a, 0xdb, 0xe8, 0x5e, 0x95, 0x38, 0x03, 0x8b, 0xe5, 0x65, 0x66, 0x85,
	0xca, 0x38, 0xe1, 0x45, 0x4a, 0x95, 0x29, 0x33, 0xcf, 0xe3, 0x13, 0x1f, 0x69, 0x3e, 0x24, 0xd3,
	0xb8, 0x1d, 0xec, 0x04, 0x72, 0x23, 0x06, 0xdd, 0x5d, 0x62, 0x75, 0x59, 0xa4, 0xd4, 0xdf, 0x23,
	0xae, 0xff, 0xc3, 0x5e, 0x34, 0xdc, 0x25, 0xb5, 0xb3, 0x2d, 0xf2, 0xb0, 0x8b, 0xde, 0x25, 0xff,
	0x37, 0xfa, 0x93, 0xef, 0x02, 0x00, 0x00, 0xff, 0xff, 0x52, 0x4d, 0xc5, 0x23, 0xbc, 0x0f, 0x00,
	0x00,
}
