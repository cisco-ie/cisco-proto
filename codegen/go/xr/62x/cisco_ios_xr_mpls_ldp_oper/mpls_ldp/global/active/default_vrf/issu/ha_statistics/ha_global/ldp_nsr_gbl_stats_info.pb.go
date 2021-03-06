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
// source: ldp_nsr_gbl_stats_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_issu_ha_statistics_ha_global

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

type LdpNsrGblStatsInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrGblStatsInfo_KEYS) Reset()         { *m = LdpNsrGblStatsInfo_KEYS{} }
func (m *LdpNsrGblStatsInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpNsrGblStatsInfo_KEYS) ProtoMessage()    {}
func (*LdpNsrGblStatsInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_485776815a97187c, []int{0}
}

func (m *LdpNsrGblStatsInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrGblStatsInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpNsrGblStatsInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrGblStatsInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpNsrGblStatsInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrGblStatsInfo_KEYS.Merge(m, src)
}
func (m *LdpNsrGblStatsInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpNsrGblStatsInfo_KEYS.Size(m)
}
func (m *LdpNsrGblStatsInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrGblStatsInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrGblStatsInfo_KEYS proto.InternalMessageInfo

type LdpNsrGblSynciInfo struct {
	NsrCfged             bool     `protobuf:"varint,1,opt,name=nsr_cfged,json=nsrCfged,proto3" json:"nsr_cfged,omitempty"`
	NsrSynced            bool     `protobuf:"varint,2,opt,name=nsr_synced,json=nsrSynced,proto3" json:"nsr_synced,omitempty"`
	InitSyncStart        uint32   `protobuf:"varint,3,opt,name=init_sync_start,json=initSyncStart,proto3" json:"init_sync_start,omitempty"`
	InitSyncEnd          uint32   `protobuf:"varint,4,opt,name=init_sync_end,json=initSyncEnd,proto3" json:"init_sync_end,omitempty"`
	NumPeers             uint32   `protobuf:"varint,5,opt,name=num_peers,json=numPeers,proto3" json:"num_peers,omitempty"`
	NumCapSent           uint32   `protobuf:"varint,6,opt,name=num_cap_sent,json=numCapSent,proto3" json:"num_cap_sent,omitempty"`
	NumCapRcvd           uint32   `protobuf:"varint,7,opt,name=num_cap_rcvd,json=numCapRcvd,proto3" json:"num_cap_rcvd,omitempty"`
	NumPfx               uint32   `protobuf:"varint,8,opt,name=num_pfx,json=numPfx,proto3" json:"num_pfx,omitempty"`
	NumLbl               uint32   `protobuf:"varint,9,opt,name=num_lbl,json=numLbl,proto3" json:"num_lbl,omitempty"`
	NumLclAddrWd         uint32   `protobuf:"varint,10,opt,name=num_lcl_addr_wd,json=numLclAddrWd,proto3" json:"num_lcl_addr_wd,omitempty"`
	NumLblAdv            uint32   `protobuf:"varint,11,opt,name=num_lbl_adv,json=numLblAdv,proto3" json:"num_lbl_adv,omitempty"`
	IpcMsgTxCnt          uint32   `protobuf:"varint,12,opt,name=ipc_msg_tx_cnt,json=ipcMsgTxCnt,proto3" json:"ipc_msg_tx_cnt,omitempty"`
	IpcMsgTxBytes        uint32   `protobuf:"varint,13,opt,name=ipc_msg_tx_bytes,json=ipcMsgTxBytes,proto3" json:"ipc_msg_tx_bytes,omitempty"`
	IpcMsgRxCnt          uint32   `protobuf:"varint,14,opt,name=ipc_msg_rx_cnt,json=ipcMsgRxCnt,proto3" json:"ipc_msg_rx_cnt,omitempty"`
	IpcMsgRxBytes        uint32   `protobuf:"varint,15,opt,name=ipc_msg_rx_bytes,json=ipcMsgRxBytes,proto3" json:"ipc_msg_rx_bytes,omitempty"`
	IpcMaxTxBatchBytes   uint32   `protobuf:"varint,16,opt,name=ipc_max_tx_batch_bytes,json=ipcMaxTxBatchBytes,proto3" json:"ipc_max_tx_batch_bytes,omitempty"`
	IpcMaxRxBatchBytes   uint32   `protobuf:"varint,17,opt,name=ipc_max_rx_batch_bytes,json=ipcMaxRxBatchBytes,proto3" json:"ipc_max_rx_batch_bytes,omitempty"`
	IpcTxFailCnt         uint32   `protobuf:"varint,18,opt,name=ipc_tx_fail_cnt,json=ipcTxFailCnt,proto3" json:"ipc_tx_fail_cnt,omitempty"`
	TotalIpcTxFailCnt    uint32   `protobuf:"varint,19,opt,name=total_ipc_tx_fail_cnt,json=totalIpcTxFailCnt,proto3" json:"total_ipc_tx_fail_cnt,omitempty"`
	IpcRestartCnt        uint32   `protobuf:"varint,20,opt,name=ipc_restart_cnt,json=ipcRestartCnt,proto3" json:"ipc_restart_cnt,omitempty"`
	IpcDefaultMtu        uint32   `protobuf:"varint,21,opt,name=ipc_default_mtu,json=ipcDefaultMtu,proto3" json:"ipc_default_mtu,omitempty"`
	IpcExceededMtuMsgCnt uint32   `protobuf:"varint,22,opt,name=ipc_exceeded_mtu_msg_cnt,json=ipcExceededMtuMsgCnt,proto3" json:"ipc_exceeded_mtu_msg_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrGblSynciInfo) Reset()         { *m = LdpNsrGblSynciInfo{} }
func (m *LdpNsrGblSynciInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNsrGblSynciInfo) ProtoMessage()    {}
func (*LdpNsrGblSynciInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_485776815a97187c, []int{1}
}

func (m *LdpNsrGblSynciInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrGblSynciInfo.Unmarshal(m, b)
}
func (m *LdpNsrGblSynciInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrGblSynciInfo.Marshal(b, m, deterministic)
}
func (m *LdpNsrGblSynciInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrGblSynciInfo.Merge(m, src)
}
func (m *LdpNsrGblSynciInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNsrGblSynciInfo.Size(m)
}
func (m *LdpNsrGblSynciInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrGblSynciInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrGblSynciInfo proto.InternalMessageInfo

func (m *LdpNsrGblSynciInfo) GetNsrCfged() bool {
	if m != nil {
		return m.NsrCfged
	}
	return false
}

func (m *LdpNsrGblSynciInfo) GetNsrSynced() bool {
	if m != nil {
		return m.NsrSynced
	}
	return false
}

func (m *LdpNsrGblSynciInfo) GetInitSyncStart() uint32 {
	if m != nil {
		return m.InitSyncStart
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetInitSyncEnd() uint32 {
	if m != nil {
		return m.InitSyncEnd
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumPeers() uint32 {
	if m != nil {
		return m.NumPeers
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumCapSent() uint32 {
	if m != nil {
		return m.NumCapSent
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumCapRcvd() uint32 {
	if m != nil {
		return m.NumCapRcvd
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumPfx() uint32 {
	if m != nil {
		return m.NumPfx
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumLbl() uint32 {
	if m != nil {
		return m.NumLbl
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumLclAddrWd() uint32 {
	if m != nil {
		return m.NumLclAddrWd
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetNumLblAdv() uint32 {
	if m != nil {
		return m.NumLblAdv
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMsgTxCnt() uint32 {
	if m != nil {
		return m.IpcMsgTxCnt
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMsgTxBytes() uint32 {
	if m != nil {
		return m.IpcMsgTxBytes
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMsgRxCnt() uint32 {
	if m != nil {
		return m.IpcMsgRxCnt
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMsgRxBytes() uint32 {
	if m != nil {
		return m.IpcMsgRxBytes
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMaxTxBatchBytes() uint32 {
	if m != nil {
		return m.IpcMaxTxBatchBytes
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcMaxRxBatchBytes() uint32 {
	if m != nil {
		return m.IpcMaxRxBatchBytes
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcTxFailCnt() uint32 {
	if m != nil {
		return m.IpcTxFailCnt
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetTotalIpcTxFailCnt() uint32 {
	if m != nil {
		return m.TotalIpcTxFailCnt
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcRestartCnt() uint32 {
	if m != nil {
		return m.IpcRestartCnt
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcDefaultMtu() uint32 {
	if m != nil {
		return m.IpcDefaultMtu
	}
	return 0
}

func (m *LdpNsrGblSynciInfo) GetIpcExceededMtuMsgCnt() uint32 {
	if m != nil {
		return m.IpcExceededMtuMsgCnt
	}
	return 0
}

type LdpNsrGblStatsInfo struct {
	InitSync             *LdpNsrGblSynciInfo `protobuf:"bytes,50,opt,name=init_sync,json=initSync,proto3" json:"init_sync,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LdpNsrGblStatsInfo) Reset()         { *m = LdpNsrGblStatsInfo{} }
func (m *LdpNsrGblStatsInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNsrGblStatsInfo) ProtoMessage()    {}
func (*LdpNsrGblStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_485776815a97187c, []int{2}
}

func (m *LdpNsrGblStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrGblStatsInfo.Unmarshal(m, b)
}
func (m *LdpNsrGblStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrGblStatsInfo.Marshal(b, m, deterministic)
}
func (m *LdpNsrGblStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrGblStatsInfo.Merge(m, src)
}
func (m *LdpNsrGblStatsInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNsrGblStatsInfo.Size(m)
}
func (m *LdpNsrGblStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrGblStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrGblStatsInfo proto.InternalMessageInfo

func (m *LdpNsrGblStatsInfo) GetInitSync() *LdpNsrGblSynciInfo {
	if m != nil {
		return m.InitSync
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpNsrGblStatsInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_global.ldp_nsr_gbl_stats_info_KEYS")
	proto.RegisterType((*LdpNsrGblSynciInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_global.ldp_nsr_gbl_synci_info")
	proto.RegisterType((*LdpNsrGblStatsInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_global.ldp_nsr_gbl_stats_info")
}

func init() { proto.RegisterFile("ldp_nsr_gbl_stats_info.proto", fileDescriptor_485776815a97187c) }

var fileDescriptor_485776815a97187c = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xd5, 0xff, 0x1f, 0xba, 0xd6, 0x5d, 0xd7, 0xcd, 0x6c, 0xc3, 0xd2, 0x18, 0xaa, 0x8a,
	0x36, 0x76, 0x15, 0xa1, 0x21, 0x71, 0xbf, 0x95, 0x21, 0x21, 0x56, 0x09, 0xa5, 0x93, 0x10, 0xdc,
	0x58, 0xae, 0xed, 0x74, 0x96, 0x1c, 0x27, 0xb2, 0x9d, 0x90, 0x3d, 0x05, 0xef, 0xc1, 0x4b, 0xf1,
	0x2a, 0xc8, 0x27, 0x49, 0xc9, 0xc6, 0x6e, 0xb9, 0x3c, 0xe7, 0xfb, 0x7d, 0xdf, 0x71, 0xac, 0x1c,
	0xa3, 0x17, 0x5a, 0xe4, 0xd4, 0x38, 0x4b, 0xd7, 0x2b, 0x4d, 0x9d, 0x67, 0xde, 0x51, 0x65, 0x92,
	0x2c, 0xca, 0x6d, 0xe6, 0x33, 0xfc, 0x8d, 0x2b, 0xc7, 0x33, 0xaa, 0x32, 0x47, 0x2b, 0x4b, 0xd3,
	0x5c, 0x3b, 0x1a, 0xf8, 0x2c, 0x97, 0x36, 0x6a, 0xab, 0x68, 0xad, 0xb3, 0x15, 0xd3, 0x11, 0xe3,
	0x5e, 0x95, 0x32, 0x12, 0x32, 0x61, 0x85, 0xf6, 0xb4, 0xb4, 0x49, 0xa4, 0x9c, 0x2b, 0xa2, 0x5b,
	0x06, 0xb9, 0xca, 0x79, 0xc5, 0x5d, 0xa8, 0x6a, 0x7e, 0x76, 0x8c, 0x8e, 0x1e, 0x9f, 0x4d, 0x3f,
	0x5d, 0x7d, 0x5d, 0xce, 0x7e, 0xf5, 0xd1, 0xe1, 0x3d, 0xfd, 0xce, 0x70, 0x05, 0x3a, 0x3e, 0x42,
	0xc3, 0xd0, 0xe5, 0xc9, 0x5a, 0x0a, 0xd2, 0x9b, 0xf6, 0xce, 0x06, 0xf1, 0xc0, 0x38, 0x3b, 0x0f,
	0x35, 0x3e, 0x46, 0x28, 0x88, 0x01, 0x97, 0x82, 0xfc, 0x07, 0x6a, 0xc0, 0x97, 0xd0, 0xc0, 0xa7,
	0x68, 0xa2, 0x8c, 0xf2, 0xa0, 0x87, 0x99, 0xd6, 0x93, 0xff, 0xa7, 0xbd, 0xb3, 0x71, 0x3c, 0x0e,
	0xed, 0x00, 0x2d, 0x43, 0x13, 0xcf, 0xd0, 0xf8, 0x0f, 0x27, 0x8d, 0x20, 0x4f, 0x80, 0x1a, 0xb5,
	0xd4, 0x95, 0x11, 0x70, 0x8e, 0x22, 0xa5, 0xb9, 0x94, 0xd6, 0x91, 0xa7, 0xa0, 0x0f, 0x4c, 0x91,
	0x7e, 0x0e, 0x35, 0x9e, 0xa2, 0xed, 0x20, 0x72, 0x96, 0x53, 0x27, 0x8d, 0x27, 0x7d, 0xd0, 0x91,
	0x29, 0xd2, 0x39, 0xcb, 0x97, 0xd2, 0xf8, 0x2e, 0x61, 0x79, 0x29, 0xc8, 0x56, 0x97, 0x88, 0x79,
	0x29, 0xf0, 0x73, 0xb4, 0x05, 0x03, 0x92, 0x8a, 0x0c, 0x40, 0xec, 0x87, 0xf8, 0xa4, 0x6a, 0x05,
	0xbd, 0xd2, 0x64, 0xb8, 0x11, 0xae, 0x57, 0x1a, 0x9f, 0xa0, 0x09, 0x08, 0x5c, 0x53, 0x26, 0x84,
	0xa5, 0xdf, 0x05, 0x41, 0x00, 0x84, 0x51, 0xd7, 0x5c, 0x5f, 0x08, 0x61, 0xbf, 0x08, 0xfc, 0x12,
	0x8d, 0x1a, 0x3f, 0x65, 0xa2, 0x24, 0x23, 0x40, 0x86, 0x75, 0xc6, 0x85, 0x28, 0xf1, 0x2b, 0xb4,
	0xa3, 0x72, 0x4e, 0x53, 0xb7, 0xa6, 0xbe, 0xa2, 0xdc, 0x78, 0xb2, 0xdd, 0x7c, 0x7e, 0xce, 0x17,
	0x6e, 0x7d, 0x53, 0xcd, 0x8d, 0xc7, 0xaf, 0xd1, 0x6e, 0x07, 0x5a, 0xdd, 0x79, 0xe9, 0xc8, 0xb8,
	0xb9, 0xcb, 0x06, 0xbb, 0x0c, 0xcd, 0x6e, 0x9a, 0xad, 0xd3, 0x76, 0xba, 0x69, 0xf1, 0xc3, 0x34,
	0xdb, 0xa6, 0x4d, 0xba, 0x69, 0x71, 0x93, 0x76, 0x8e, 0x0e, 0x01, 0x64, 0x15, 0x8c, 0x65, 0x9e,
	0xdf, 0x36, 0xf8, 0x2e, 0xe0, 0x38, 0xe0, 0xac, 0xba, 0xa9, 0x2e, 0x83, 0xf4, 0x97, 0xc7, 0xde,
	0xf7, 0xec, 0x75, 0x3d, 0x71, 0xd7, 0x73, 0x82, 0x26, 0xc1, 0xe3, 0x2b, 0x9a, 0x30, 0xa5, 0xe1,
	0xd8, 0xb8, 0xbe, 0x4a, 0x95, 0xf3, 0x9b, 0xea, 0x03, 0x53, 0x3a, 0x9c, 0xfb, 0x0d, 0x3a, 0xf0,
	0x99, 0x67, 0x9a, 0x3e, 0x84, 0x9f, 0x01, 0xbc, 0x07, 0xe2, 0xc7, 0xae, 0xe3, 0xb4, 0x0e, 0xb6,
	0x12, 0x7e, 0x3f, 0x60, 0xf7, 0x37, 0x1f, 0x1a, 0xd7, 0xdd, 0x0e, 0xd7, 0x2e, 0x54, 0xea, 0x0b,
	0x72, 0xb0, 0xe1, 0xde, 0xd7, 0xdd, 0x85, 0x2f, 0xf0, 0x3b, 0x44, 0x02, 0x27, 0x2b, 0x2e, 0xa5,
	0x90, 0x22, 0x80, 0x70, 0x8d, 0x21, 0xf8, 0x10, 0x0c, 0xfb, 0x2a, 0xe7, 0x57, 0x8d, 0xbc, 0xf0,
	0xc5, 0xc2, 0xad, 0xe7, 0xc6, 0xcf, 0x7e, 0xf6, 0x1e, 0x6c, 0xd8, 0x66, 0x03, 0xf1, 0x8f, 0x1e,
	0x1a, 0x6e, 0x7e, 0x7f, 0x72, 0x3e, 0xed, 0x9d, 0x8d, 0xce, 0x6d, 0xf4, 0xef, 0x1e, 0x83, 0xe8,
	0xf1, 0x4d, 0x8f, 0x07, 0xed, 0xba, 0xad, 0xfa, 0xf0, 0x20, 0xbd, 0xfd, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x70, 0xb0, 0x76, 0xd7, 0xb0, 0x04, 0x00, 0x00,
}
