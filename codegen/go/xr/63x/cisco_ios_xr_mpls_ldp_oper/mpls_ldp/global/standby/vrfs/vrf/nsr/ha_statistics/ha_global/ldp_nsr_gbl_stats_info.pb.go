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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_nsr_ha_statistics_ha_global

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
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

func (m *LdpNsrGblStatsInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*LdpNsrGblStatsInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.nsr.ha_statistics.ha_global.ldp_nsr_gbl_stats_info_KEYS")
	proto.RegisterType((*LdpNsrGblSynciInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.nsr.ha_statistics.ha_global.ldp_nsr_gbl_synci_info")
	proto.RegisterType((*LdpNsrGblStatsInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.nsr.ha_statistics.ha_global.ldp_nsr_gbl_stats_info")
}

func init() { proto.RegisterFile("ldp_nsr_gbl_stats_info.proto", fileDescriptor_485776815a97187c) }

var fileDescriptor_485776815a97187c = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5f, 0x6b, 0xdb, 0x3c,
	0x14, 0xc6, 0xc9, 0xfb, 0x27, 0x4d, 0x94, 0xa6, 0x69, 0xf5, 0xb6, 0x7d, 0x35, 0xba, 0x8d, 0x90,
	0xd1, 0xae, 0x57, 0x66, 0x74, 0x30, 0x76, 0xdb, 0x66, 0x1d, 0x8c, 0x35, 0x63, 0x38, 0x85, 0xb2,
	0x2b, 0x21, 0x4b, 0x72, 0x2a, 0x90, 0x65, 0x23, 0xc9, 0x9e, 0xfb, 0x1d, 0xf6, 0x3d, 0xf6, 0xb1,
	0xf6, 0x55, 0x86, 0x8e, 0x9d, 0xcc, 0xed, 0x7a, 0xbb, 0x1b, 0x83, 0xce, 0xf3, 0x7b, 0x9e, 0x73,
	0x2c, 0x7c, 0x8c, 0x9e, 0x6a, 0x51, 0x50, 0xe3, 0x2c, 0x5d, 0x25, 0x9a, 0x3a, 0xcf, 0xbc, 0xa3,
	0xca, 0xa4, 0x79, 0x54, 0xd8, 0xdc, 0xe7, 0xf8, 0x86, 0x2b, 0xc7, 0x73, 0xaa, 0x72, 0x47, 0x6b,
	0x4b, 0xb3, 0x42, 0x3b, 0x1a, 0xf8, 0xbc, 0x90, 0x36, 0x5a, 0x9f, 0xa2, 0x95, 0xce, 0x13, 0xa6,
	0x23, 0xe7, 0x99, 0x11, 0xc9, 0x5d, 0x54, 0xd9, 0xd4, 0x85, 0x47, 0x64, 0x9c, 0x8d, 0x6e, 0x19,
	0x84, 0x2a, 0xe7, 0x15, 0x77, 0xe1, 0xd4, 0xc0, 0xb3, 0xb7, 0xe8, 0xe8, 0xf1, 0xc6, 0xf4, 0xe3,
	0xe5, 0x97, 0x25, 0x7e, 0x82, 0x06, 0x95, 0x4d, 0xa9, 0x61, 0x99, 0x24, 0xbd, 0x69, 0xef, 0x74,
	0x18, 0x6f, 0x55, 0x36, 0xfd, 0xc4, 0x32, 0x39, 0xfb, 0xd1, 0x47, 0x87, 0xf7, 0xac, 0x77, 0x86,
	0x2b, 0xb0, 0xe2, 0x23, 0x34, 0x0c, 0x55, 0x9e, 0xae, 0xa4, 0x00, 0xdb, 0x20, 0x1e, 0x18, 0x67,
	0xe7, 0xe1, 0x8c, 0x9f, 0x21, 0x14, 0xc4, 0x80, 0x4b, 0x41, 0xfe, 0x02, 0x35, 0xe0, 0x4b, 0x28,
	0xe0, 0x13, 0x34, 0x51, 0x46, 0x79, 0xd0, 0xc3, 0x38, 0xd6, 0x93, 0xbf, 0xa7, 0xbd, 0xd3, 0x71,
	0x3c, 0x0e, 0xe5, 0x00, 0x2d, 0x43, 0x11, 0xcf, 0xd0, 0xf8, 0x17, 0x27, 0x8d, 0x20, 0xff, 0x00,
	0x35, 0x5a, 0x53, 0x97, 0x46, 0xc0, 0x1c, 0x65, 0x46, 0x0b, 0x29, 0xad, 0x23, 0xff, 0x82, 0x3e,
	0x30, 0x65, 0xf6, 0x39, 0x9c, 0xf1, 0x14, 0x6d, 0x07, 0x91, 0xb3, 0x82, 0x3a, 0x69, 0x3c, 0xe9,
	0x83, 0x8e, 0x4c, 0x99, 0xcd, 0x59, 0xb1, 0x94, 0xc6, 0x77, 0x09, 0xcb, 0x2b, 0x41, 0xb6, 0xba,
	0x44, 0xcc, 0x2b, 0x81, 0xff, 0x47, 0x5b, 0xd0, 0x20, 0xad, 0xc9, 0x00, 0xc4, 0x7e, 0x88, 0x4f,
	0xeb, 0xb5, 0xa0, 0x13, 0x4d, 0x86, 0x1b, 0xe1, 0x2a, 0xd1, 0xf8, 0x18, 0x4d, 0x40, 0xe0, 0x9a,
	0x32, 0x21, 0x2c, 0xfd, 0x2a, 0x08, 0x02, 0x20, 0xb4, 0xba, 0xe2, 0xfa, 0x5c, 0x08, 0x7b, 0x23,
	0xf0, 0x73, 0x34, 0x6a, 0xfd, 0x94, 0x89, 0x8a, 0x8c, 0x00, 0x19, 0x36, 0x19, 0xe7, 0xa2, 0xc2,
	0x2f, 0xd0, 0x8e, 0x2a, 0x38, 0xcd, 0xdc, 0x8a, 0xfa, 0x9a, 0x72, 0xe3, 0xc9, 0x76, 0xfb, 0xfa,
	0x05, 0x5f, 0xb8, 0xd5, 0x75, 0x3d, 0x37, 0x1e, 0xbf, 0x44, 0xbb, 0x1d, 0x28, 0xb9, 0xf3, 0xd2,
	0x91, 0x71, 0x7b, 0x97, 0x2d, 0x76, 0x11, 0x8a, 0xdd, 0x34, 0xdb, 0xa4, 0xed, 0x74, 0xd3, 0xe2,
	0x87, 0x69, 0x76, 0x9d, 0x36, 0xe9, 0xa6, 0xc5, 0x6d, 0xda, 0x19, 0x3a, 0x04, 0x90, 0xd5, 0xd0,
	0x96, 0x79, 0x7e, 0xdb, 0xe2, 0xbb, 0x80, 0xe3, 0x80, 0xb3, 0xfa, 0xba, 0xbe, 0x08, 0xd2, 0x6f,
	0x1e, 0x7b, 0xdf, 0xb3, 0xd7, 0xf5, 0xc4, 0x5d, 0xcf, 0x31, 0x9a, 0x04, 0x8f, 0xaf, 0x69, 0xca,
	0x94, 0x86, 0xb1, 0x71, 0x73, 0x95, 0xaa, 0xe0, 0xd7, 0xf5, 0x7b, 0xa6, 0x74, 0x98, 0xfb, 0x15,
	0x3a, 0xf0, 0xb9, 0x67, 0x9a, 0x3e, 0x84, 0xff, 0x03, 0x78, 0x0f, 0xc4, 0x0f, 0x5d, 0xc7, 0x49,
	0x13, 0x6c, 0x25, 0x7c, 0x7e, 0xc0, 0xee, 0x6f, 0x5e, 0x34, 0x6e, 0xaa, 0x1d, 0x4e, 0xc8, 0x94,
	0x95, 0xda, 0xd3, 0xcc, 0x97, 0xe4, 0x60, 0xc3, 0xbd, 0x6b, 0xaa, 0x0b, 0x5f, 0xe2, 0x37, 0x88,
	0x04, 0x4e, 0xd6, 0x5c, 0x4a, 0x21, 0x45, 0x00, 0xe1, 0x1a, 0x43, 0xf0, 0x21, 0x18, 0xf6, 0x55,
	0xc1, 0x2f, 0x5b, 0x79, 0xe1, 0xcb, 0x85, 0x5b, 0xcd, 0x8d, 0x9f, 0x7d, 0xef, 0x3d, 0xd8, 0xb0,
	0xcd, 0x72, 0xe2, 0x6f, 0x3d, 0x34, 0xdc, 0x7c, 0xfe, 0xe4, 0x6c, 0xda, 0x3b, 0x1d, 0x9d, 0xe5,
	0xd1, 0x1f, 0xfa, 0x49, 0x44, 0x8f, 0xaf, 0x79, 0x3c, 0x58, 0xef, 0x5a, 0xd2, 0x87, 0xbf, 0xd4,
	0xeb, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x28, 0x52, 0xa3, 0xc5, 0x04, 0x00, 0x00,
}