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
// source: tcp_nsr_sscb_detailed_bag.proto

package cisco_ios_xr_ip_tcp_oper_tcp_nsr_nodes_node_session_set_detail_sets_detail_set

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

type TcpNsrSscbDetailedBag_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Id_1                 string   `protobuf:"bytes,2,opt,name=id_1,json=id1,proto3" json:"id_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpNsrSscbDetailedBag_KEYS) Reset()         { *m = TcpNsrSscbDetailedBag_KEYS{} }
func (m *TcpNsrSscbDetailedBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSscbDetailedBag_KEYS) ProtoMessage()    {}
func (*TcpNsrSscbDetailedBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f131f96c81678fc, []int{0}
}

func (m *TcpNsrSscbDetailedBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS.Unmarshal(m, b)
}
func (m *TcpNsrSscbDetailedBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS.Marshal(b, m, deterministic)
}
func (m *TcpNsrSscbDetailedBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS.Merge(m, src)
}
func (m *TcpNsrSscbDetailedBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS.Size(m)
}
func (m *TcpNsrSscbDetailedBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSscbDetailedBag_KEYS proto.InternalMessageInfo

func (m *TcpNsrSscbDetailedBag_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TcpNsrSscbDetailedBag_KEYS) GetId_1() string {
	if m != nil {
		return m.Id_1
	}
	return ""
}

type TcpNsrSscbDetailedBag struct {
	Sscb                             uint64   `protobuf:"varint,50,opt,name=sscb,proto3" json:"sscb,omitempty"`
	Pid                              uint32   `protobuf:"varint,51,opt,name=pid,proto3" json:"pid,omitempty"`
	SetId                            uint32   `protobuf:"varint,52,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty"`
	SsoRole                          uint32   `protobuf:"varint,53,opt,name=sso_role,json=ssoRole,proto3" json:"sso_role,omitempty"`
	Mode                             uint32   `protobuf:"varint,54,opt,name=mode,proto3" json:"mode,omitempty"`
	AddressFamily                    string   `protobuf:"bytes,55,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	WellKnownPort                    uint32   `protobuf:"varint,56,opt,name=well_known_port,json=wellKnownPort,proto3" json:"well_known_port,omitempty"`
	LocalNode                        string   `protobuf:"bytes,57,opt,name=local_node,json=localNode,proto3" json:"local_node,omitempty"`
	LocalInstance                    uint32   `protobuf:"varint,58,opt,name=local_instance,json=localInstance,proto3" json:"local_instance,omitempty"`
	ProtectNode                      string   `protobuf:"bytes,59,opt,name=protect_node,json=protectNode,proto3" json:"protect_node,omitempty"`
	ProtectInstance                  uint32   `protobuf:"varint,60,opt,name=protect_instance,json=protectInstance,proto3" json:"protect_instance,omitempty"`
	NumberOfSessions                 uint32   `protobuf:"varint,61,opt,name=number_of_sessions,json=numberOfSessions,proto3" json:"number_of_sessions,omitempty"`
	NumberOfSyncedSessionsUpStream   uint32   `protobuf:"varint,62,opt,name=number_of_synced_sessions_up_stream,json=numberOfSyncedSessionsUpStream,proto3" json:"number_of_synced_sessions_up_stream,omitempty"`
	NumberOfSyncedSessionsDownStream uint32   `protobuf:"varint,63,opt,name=number_of_synced_sessions_down_stream,json=numberOfSyncedSessionsDownStream,proto3" json:"number_of_synced_sessions_down_stream,omitempty"`
	IsInitSyncInProgress             bool     `protobuf:"varint,64,opt,name=is_init_sync_in_progress,json=isInitSyncInProgress,proto3" json:"is_init_sync_in_progress,omitempty"`
	IsInitSyncSecondPhase            bool     `protobuf:"varint,65,opt,name=is_init_sync_second_phase,json=isInitSyncSecondPhase,proto3" json:"is_init_sync_second_phase,omitempty"`
	SequenceNumberOfInitSync         uint32   `protobuf:"varint,66,opt,name=sequence_number_of_init_sync,json=sequenceNumberOfInitSync,proto3" json:"sequence_number_of_init_sync,omitempty"`
	InitSyncTimer                    uint32   `protobuf:"varint,67,opt,name=init_sync_timer,json=initSyncTimer,proto3" json:"init_sync_timer,omitempty"`
	TotalNumberOfInitSyncSessions    uint32   `protobuf:"varint,68,opt,name=total_number_of_init_sync_sessions,json=totalNumberOfInitSyncSessions,proto3" json:"total_number_of_init_sync_sessions,omitempty"`
	NumberOfInitSyncedSessions       uint32   `protobuf:"varint,69,opt,name=number_of_init_synced_sessions,json=numberOfInitSyncedSessions,proto3" json:"number_of_init_synced_sessions,omitempty"`
	NumberOfSessionsInitSyncFailed   uint32   `protobuf:"varint,70,opt,name=number_of_sessions_init_sync_failed,json=numberOfSessionsInitSyncFailed,proto3" json:"number_of_sessions_init_sync_failed,omitempty"`
	InitSyncError                    []string `protobuf:"bytes,71,rep,name=init_sync_error,json=initSyncError,proto3" json:"init_sync_error,omitempty"`
	IsInitSyncErrorLocal             bool     `protobuf:"varint,72,opt,name=is_init_sync_error_local,json=isInitSyncErrorLocal,proto3" json:"is_init_sync_error_local,omitempty"`
	InitSyncStartTime                uint32   `protobuf:"varint,73,opt,name=init_sync_start_time,json=initSyncStartTime,proto3" json:"init_sync_start_time,omitempty"`
	InitSyncEndTime                  uint32   `protobuf:"varint,74,opt,name=init_sync_end_time,json=initSyncEndTime,proto3" json:"init_sync_end_time,omitempty"`
	IsSscbInitSyncReady              bool     `protobuf:"varint,75,opt,name=is_sscb_init_sync_ready,json=isSscbInitSyncReady,proto3" json:"is_sscb_init_sync_ready,omitempty"`
	InitSyncReadyStartTime           uint32   `protobuf:"varint,76,opt,name=init_sync_ready_start_time,json=initSyncReadyStartTime,proto3" json:"init_sync_ready_start_time,omitempty"`
	InitSyncReadyEndTime             uint32   `protobuf:"varint,77,opt,name=init_sync_ready_end_time,json=initSyncReadyEndTime,proto3" json:"init_sync_ready_end_time,omitempty"`
	NsrResetTime                     uint32   `protobuf:"varint,78,opt,name=nsr_reset_time,json=nsrResetTime,proto3" json:"nsr_reset_time,omitempty"`
	IsAuditInProgress                bool     `protobuf:"varint,79,opt,name=is_audit_in_progress,json=isAuditInProgress,proto3" json:"is_audit_in_progress,omitempty"`
	AuditSeqNumber                   uint32   `protobuf:"varint,80,opt,name=audit_seq_number,json=auditSeqNumber,proto3" json:"audit_seq_number,omitempty"`
	AuditStartTime                   uint32   `protobuf:"varint,81,opt,name=audit_start_time,json=auditStartTime,proto3" json:"audit_start_time,omitempty"`
	AuditEndTime                     uint32   `protobuf:"varint,82,opt,name=audit_end_time,json=auditEndTime,proto3" json:"audit_end_time,omitempty"`
	XXX_NoUnkeyedLiteral             struct{} `json:"-"`
	XXX_unrecognized                 []byte   `json:"-"`
	XXX_sizecache                    int32    `json:"-"`
}

func (m *TcpNsrSscbDetailedBag) Reset()         { *m = TcpNsrSscbDetailedBag{} }
func (m *TcpNsrSscbDetailedBag) String() string { return proto.CompactTextString(m) }
func (*TcpNsrSscbDetailedBag) ProtoMessage()    {}
func (*TcpNsrSscbDetailedBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f131f96c81678fc, []int{1}
}

func (m *TcpNsrSscbDetailedBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpNsrSscbDetailedBag.Unmarshal(m, b)
}
func (m *TcpNsrSscbDetailedBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpNsrSscbDetailedBag.Marshal(b, m, deterministic)
}
func (m *TcpNsrSscbDetailedBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpNsrSscbDetailedBag.Merge(m, src)
}
func (m *TcpNsrSscbDetailedBag) XXX_Size() int {
	return xxx_messageInfo_TcpNsrSscbDetailedBag.Size(m)
}
func (m *TcpNsrSscbDetailedBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpNsrSscbDetailedBag.DiscardUnknown(m)
}

var xxx_messageInfo_TcpNsrSscbDetailedBag proto.InternalMessageInfo

func (m *TcpNsrSscbDetailedBag) GetSscb() uint64 {
	if m != nil {
		return m.Sscb
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetSetId() uint32 {
	if m != nil {
		return m.SetId
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetSsoRole() uint32 {
	if m != nil {
		return m.SsoRole
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *TcpNsrSscbDetailedBag) GetWellKnownPort() uint32 {
	if m != nil {
		return m.WellKnownPort
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetLocalNode() string {
	if m != nil {
		return m.LocalNode
	}
	return ""
}

func (m *TcpNsrSscbDetailedBag) GetLocalInstance() uint32 {
	if m != nil {
		return m.LocalInstance
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetProtectNode() string {
	if m != nil {
		return m.ProtectNode
	}
	return ""
}

func (m *TcpNsrSscbDetailedBag) GetProtectInstance() uint32 {
	if m != nil {
		return m.ProtectInstance
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNumberOfSessions() uint32 {
	if m != nil {
		return m.NumberOfSessions
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNumberOfSyncedSessionsUpStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsUpStream
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNumberOfSyncedSessionsDownStream() uint32 {
	if m != nil {
		return m.NumberOfSyncedSessionsDownStream
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetIsInitSyncInProgress() bool {
	if m != nil {
		return m.IsInitSyncInProgress
	}
	return false
}

func (m *TcpNsrSscbDetailedBag) GetIsInitSyncSecondPhase() bool {
	if m != nil {
		return m.IsInitSyncSecondPhase
	}
	return false
}

func (m *TcpNsrSscbDetailedBag) GetSequenceNumberOfInitSync() uint32 {
	if m != nil {
		return m.SequenceNumberOfInitSync
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncTimer() uint32 {
	if m != nil {
		return m.InitSyncTimer
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetTotalNumberOfInitSyncSessions() uint32 {
	if m != nil {
		return m.TotalNumberOfInitSyncSessions
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNumberOfInitSyncedSessions() uint32 {
	if m != nil {
		return m.NumberOfInitSyncedSessions
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNumberOfSessionsInitSyncFailed() uint32 {
	if m != nil {
		return m.NumberOfSessionsInitSyncFailed
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncError() []string {
	if m != nil {
		return m.InitSyncError
	}
	return nil
}

func (m *TcpNsrSscbDetailedBag) GetIsInitSyncErrorLocal() bool {
	if m != nil {
		return m.IsInitSyncErrorLocal
	}
	return false
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncStartTime() uint32 {
	if m != nil {
		return m.InitSyncStartTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncEndTime() uint32 {
	if m != nil {
		return m.InitSyncEndTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetIsSscbInitSyncReady() bool {
	if m != nil {
		return m.IsSscbInitSyncReady
	}
	return false
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncReadyStartTime() uint32 {
	if m != nil {
		return m.InitSyncReadyStartTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetInitSyncReadyEndTime() uint32 {
	if m != nil {
		return m.InitSyncReadyEndTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetNsrResetTime() uint32 {
	if m != nil {
		return m.NsrResetTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetIsAuditInProgress() bool {
	if m != nil {
		return m.IsAuditInProgress
	}
	return false
}

func (m *TcpNsrSscbDetailedBag) GetAuditSeqNumber() uint32 {
	if m != nil {
		return m.AuditSeqNumber
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetAuditStartTime() uint32 {
	if m != nil {
		return m.AuditStartTime
	}
	return 0
}

func (m *TcpNsrSscbDetailedBag) GetAuditEndTime() uint32 {
	if m != nil {
		return m.AuditEndTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpNsrSscbDetailedBag_KEYS)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session_set.detail_sets.detail_set.tcp_nsr_sscb_detailed_bag_KEYS")
	proto.RegisterType((*TcpNsrSscbDetailedBag)(nil), "cisco_ios_xr_ip_tcp_oper.tcp_nsr.nodes.node.session_set.detail_sets.detail_set.tcp_nsr_sscb_detailed_bag")
}

func init() { proto.RegisterFile("tcp_nsr_sscb_detailed_bag.proto", fileDescriptor_9f131f96c81678fc) }

var fileDescriptor_9f131f96c81678fc = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x6d, 0x53, 0x63, 0x35,
	0x14, 0x9e, 0x02, 0xae, 0xbb, 0x11, 0xba, 0x25, 0xb2, 0x1a, 0x76, 0x5c, 0xac, 0xb8, 0xeb, 0xd4,
	0xd1, 0xa9, 0x83, 0x20, 0x22, 0x2a, 0xca, 0xab, 0xd6, 0x62, 0xa9, 0xb7, 0xfa, 0xc1, 0x4f, 0x99,
	0xdb, 0x9b, 0x14, 0x33, 0xde, 0x26, 0x97, 0x9c, 0x74, 0x90, 0x1f, 0xe4, 0xff, 0xdc, 0xc9, 0xb9,
	0xb9, 0x2f, 0x2d, 0xf0, 0xa5, 0x93, 0x79, 0xce, 0xf3, 0x3c, 0xe7, 0x2d, 0xb9, 0x25, 0x1f, 0xbb,
	0x24, 0xe3, 0x1a, 0x2c, 0x07, 0x48, 0xc6, 0x5c, 0x48, 0x17, 0xab, 0x54, 0x0a, 0x3e, 0x8e, 0xaf,
	0xbb, 0x99, 0x35, 0xce, 0xd0, 0x41, 0xa2, 0x20, 0x31, 0x5c, 0x19, 0xe0, 0xff, 0x59, 0xae, 0x32,
	0xee, 0x05, 0x26, 0x93, 0xb6, 0x1b, 0x94, 0x5d, 0x6d, 0x84, 0x04, 0xfc, 0xed, 0x82, 0x04, 0x50,
	0x46, 0x73, 0x90, 0xae, 0x9b, 0x7b, 0xf9, 0x23, 0xd4, 0xce, 0xdb, 0xa7, 0x64, 0xeb, 0xd1, 0x94,
	0xbc, 0x7f, 0xfe, 0xf7, 0x88, 0x36, 0xc9, 0x92, 0x12, 0xac, 0xd1, 0x6e, 0x74, 0x9e, 0x45, 0x4b,
	0x4a, 0xd0, 0x75, 0xb2, 0xa2, 0x04, 0xdf, 0x61, 0x4b, 0x88, 0x2c, 0x2b, 0xb1, 0xb3, 0xfd, 0xff,
	0x2a, 0xd9, 0x7c, 0xd4, 0x85, 0x52, 0xb2, 0xe2, 0x41, 0xf6, 0x75, 0xbb, 0xd1, 0x59, 0x89, 0xf0,
	0x4c, 0x5b, 0x64, 0x39, 0x53, 0x82, 0xed, 0xb6, 0x1b, 0x9d, 0xb5, 0xc8, 0x1f, 0xe9, 0x0b, 0xf2,
	0x04, 0xa4, 0xe3, 0x4a, 0xb0, 0x3d, 0x04, 0xdf, 0x01, 0xe9, 0x7a, 0x82, 0x6e, 0x92, 0xa7, 0x00,
	0x86, 0x5b, 0x93, 0x4a, 0xf6, 0x0d, 0x06, 0xde, 0x05, 0x30, 0x91, 0x49, 0xa5, 0xf7, 0x9d, 0x1a,
	0x21, 0xd9, 0x3e, 0xc2, 0x78, 0xa6, 0x6f, 0x48, 0x33, 0x16, 0xc2, 0x4a, 0x00, 0x3e, 0x89, 0xa7,
	0x2a, 0xbd, 0x63, 0xdf, 0x62, 0x99, 0x6b, 0x01, 0xbd, 0x40, 0x90, 0x7e, 0x46, 0x9e, 0xdf, 0xca,
	0x34, 0xe5, 0xff, 0x6a, 0x73, 0xab, 0x79, 0x66, 0xac, 0x63, 0x07, 0xe8, 0xb2, 0xe6, 0xe1, 0xbe,
	0x47, 0x87, 0xc6, 0x3a, 0xfa, 0x8a, 0x90, 0xd4, 0x24, 0x71, 0xca, 0xfd, 0x38, 0xd9, 0x77, 0x68,
	0xf5, 0x0c, 0x91, 0x41, 0xc8, 0x96, 0x87, 0x95, 0x06, 0x17, 0xeb, 0x44, 0xb2, 0xc3, 0xdc, 0x05,
	0xd1, 0x5e, 0x00, 0xe9, 0x27, 0x64, 0xd5, 0x2f, 0x4f, 0x26, 0x2e, 0xf7, 0xf9, 0x1e, 0x7d, 0xde,
	0x0b, 0x18, 0x3a, 0x7d, 0x4e, 0x5a, 0x05, 0xa5, 0xf4, 0xfa, 0x01, 0xbd, 0x9e, 0x07, 0xbc, 0x74,
	0xfb, 0x92, 0x50, 0x3d, 0x9b, 0x8e, 0xa5, 0xe5, 0x66, 0xc2, 0xc3, 0x86, 0x81, 0xfd, 0x88, 0xe4,
	0x56, 0x1e, 0xb9, 0x9a, 0x8c, 0x02, 0x4e, 0xfb, 0xe4, 0xd3, 0x1a, 0xfb, 0x4e, 0x27, 0x52, 0x94,
	0x22, 0x3e, 0xcb, 0x38, 0x38, 0x2b, 0xe3, 0x29, 0x3b, 0x42, 0xf9, 0x56, 0x29, 0x47, 0x62, 0x61,
	0xf2, 0x57, 0x36, 0x42, 0x16, 0xbd, 0x22, 0x6f, 0x1e, 0x37, 0x13, 0x7e, 0x94, 0xc1, 0xee, 0x27,
	0xb4, 0x6b, 0x3f, 0x6c, 0x77, 0x66, 0x6e, 0x75, 0x30, 0xdc, 0x27, 0x4c, 0x01, 0x57, 0x5a, 0x39,
	0xb4, 0xe3, 0x4a, 0xf3, 0xcc, 0x9a, 0x6b, 0xbf, 0x29, 0xf6, 0x73, 0xbb, 0xd1, 0x79, 0x1a, 0x6d,
	0x28, 0xe8, 0x69, 0xe5, 0xbc, 0x43, 0x4f, 0x0f, 0x43, 0x8c, 0x1e, 0x90, 0xcd, 0x39, 0x1d, 0xc8,
	0xc4, 0x68, 0xc1, 0xb3, 0x7f, 0x62, 0x90, 0xec, 0x18, 0x85, 0x2f, 0x2a, 0xe1, 0x08, 0xa3, 0x43,
	0x1f, 0xa4, 0x47, 0xe4, 0x23, 0x90, 0x37, 0x33, 0xa9, 0x13, 0xc9, 0xab, 0x5e, 0x4a, 0x27, 0x76,
	0x82, 0x95, 0xb3, 0x82, 0x33, 0x08, 0x1d, 0x14, 0x56, 0xfe, 0xe6, 0x54, 0x69, 0x9d, 0x9a, 0x4a,
	0xcb, 0x4e, 0xf3, 0x9d, 0xab, 0x40, 0xf9, 0xd3, 0x83, 0xb4, 0x47, 0xb6, 0x9d, 0x71, 0xfe, 0xe6,
	0xdc, 0x4f, 0x52, 0x6d, 0xed, 0x0c, 0xa5, 0xaf, 0x90, 0xb9, 0x98, 0xaa, 0x5c, 0xe1, 0x09, 0xd9,
	0x7a, 0xc0, 0xa4, 0x36, 0x7a, 0x76, 0x8e, 0x36, 0x2f, 0xf5, 0x82, 0x43, 0x35, 0xf2, 0x85, 0x6b,
	0x50, 0xac, 0xac, 0xaa, 0x68, 0x82, 0xcf, 0x95, 0x5d, 0x2c, 0x5c, 0x83, 0x40, 0x2c, 0x0c, 0x2f,
	0x90, 0x35, 0x3f, 0x03, 0x69, 0xad, 0xb1, 0xec, 0x97, 0xf6, 0xb2, 0x7f, 0x65, 0xc5, 0x0c, 0xce,
	0x3d, 0x78, 0x6f, 0xbb, 0x48, 0xe5, 0xf8, 0x36, 0xd8, 0xaf, 0x8b, 0xdb, 0x45, 0xc9, 0xa5, 0x8f,
	0xd1, 0xaf, 0xc8, 0x46, 0x6d, 0x56, 0x2e, 0xb6, 0x0e, 0x27, 0xcd, 0x7a, 0x58, 0xdd, 0x7a, 0x91,
	0x64, 0xe4, 0x23, 0x7e, 0xda, 0xf4, 0x0b, 0x42, 0x6b, 0x59, 0xb4, 0xc8, 0xe9, 0xbf, 0xe5, 0xef,
	0xa7, 0xac, 0x49, 0x0b, 0x24, 0xef, 0x91, 0x0f, 0x15, 0xe4, 0x9f, 0xa9, 0x4a, 0x64, 0x65, 0x2c,
	0xee, 0x58, 0x1f, 0x8b, 0x7a, 0x5f, 0xc1, 0x08, 0x92, 0x71, 0x51, 0x58, 0xe4, 0x43, 0xf4, 0x90,
	0xbc, 0x5c, 0x60, 0xd7, 0x2b, 0xbb, 0xc4, 0x54, 0x1f, 0xa8, 0xba, 0xa4, 0x2a, 0xcf, 0xcf, 0x61,
	0x41, 0x5b, 0x16, 0xf9, 0x3b, 0x2a, 0x37, 0xe6, 0x94, 0x45, 0xa5, 0xaf, 0x49, 0xd3, 0x7f, 0x51,
	0xad, 0xf4, 0x1f, 0x46, 0x64, 0x0f, 0x90, 0xbd, 0xaa, 0xc1, 0x46, 0x1e, 0x44, 0x96, 0x9f, 0x16,
	0xf0, 0x78, 0x26, 0x94, 0x9b, 0x7b, 0x3f, 0x57, 0xd8, 0xcc, 0xba, 0x82, 0x63, 0x1f, 0xaa, 0x3d,
	0x9e, 0x0e, 0x69, 0xe5, 0x6c, 0x90, 0x37, 0xe1, 0x7a, 0xb2, 0x21, 0x1a, 0x37, 0x11, 0x1f, 0xc9,
	0x9b, 0xfc, 0x2e, 0xd6, 0x98, 0x55, 0xab, 0x7f, 0xd4, 0x99, 0x65, 0x8b, 0xaf, 0x49, 0x8e, 0x54,
	0x8d, 0x45, 0x79, 0xa9, 0x88, 0x86, 0x86, 0xc6, 0x4f, 0xf0, 0x3f, 0x6c, 0xf7, 0x6d, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x07, 0x6e, 0x2f, 0x30, 0xe6, 0x06, 0x00, 0x00,
}