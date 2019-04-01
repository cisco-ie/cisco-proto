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
// source: ldp_nsr_stats_nbr_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_issu_ha_statistics_ha_neighbors_ha_neighbor

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

type LdpNsrStatsNbrInfo_KEYS struct {
	LsrId                string   `protobuf:"bytes,1,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,2,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrStatsNbrInfo_KEYS) Reset()         { *m = LdpNsrStatsNbrInfo_KEYS{} }
func (m *LdpNsrStatsNbrInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpNsrStatsNbrInfo_KEYS) ProtoMessage()    {}
func (*LdpNsrStatsNbrInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b668fae369a575f, []int{0}
}

func (m *LdpNsrStatsNbrInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpNsrStatsNbrInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpNsrStatsNbrInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS.Merge(m, src)
}
func (m *LdpNsrStatsNbrInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS.Size(m)
}
func (m *LdpNsrStatsNbrInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrStatsNbrInfo_KEYS proto.InternalMessageInfo

func (m *LdpNsrStatsNbrInfo_KEYS) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpNsrStatsNbrInfo_KEYS) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

type LdpNsrNbrSynciInfo struct {
	InitSyncStart        uint32   `protobuf:"varint,1,opt,name=init_sync_start,json=initSyncStart,proto3" json:"init_sync_start,omitempty"`
	InitSyncEnd          uint32   `protobuf:"varint,2,opt,name=init_sync_end,json=initSyncEnd,proto3" json:"init_sync_end,omitempty"`
	NumAddr              uint32   `protobuf:"varint,3,opt,name=num_addr,json=numAddr,proto3" json:"num_addr,omitempty"`
	NumDuplicateAddr     uint32   `protobuf:"varint,4,opt,name=num_duplicate_addr,json=numDuplicateAddr,proto3" json:"num_duplicate_addr,omitempty"`
	NumRxBytes           uint32   `protobuf:"varint,5,opt,name=num_rx_bytes,json=numRxBytes,proto3" json:"num_rx_bytes,omitempty"`
	NumCapSent           uint32   `protobuf:"varint,6,opt,name=num_cap_sent,json=numCapSent,proto3" json:"num_cap_sent,omitempty"`
	NumCapRcvd           uint32   `protobuf:"varint,7,opt,name=num_cap_rcvd,json=numCapRcvd,proto3" json:"num_cap_rcvd,omitempty"`
	NumLbl               uint32   `protobuf:"varint,8,opt,name=num_lbl,json=numLbl,proto3" json:"num_lbl,omitempty"`
	NumAppBytes          uint32   `protobuf:"varint,9,opt,name=num_app_bytes,json=numAppBytes,proto3" json:"num_app_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrNbrSynciInfo) Reset()         { *m = LdpNsrNbrSynciInfo{} }
func (m *LdpNsrNbrSynciInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNsrNbrSynciInfo) ProtoMessage()    {}
func (*LdpNsrNbrSynciInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b668fae369a575f, []int{1}
}

func (m *LdpNsrNbrSynciInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrNbrSynciInfo.Unmarshal(m, b)
}
func (m *LdpNsrNbrSynciInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrNbrSynciInfo.Marshal(b, m, deterministic)
}
func (m *LdpNsrNbrSynciInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrNbrSynciInfo.Merge(m, src)
}
func (m *LdpNsrNbrSynciInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNsrNbrSynciInfo.Size(m)
}
func (m *LdpNsrNbrSynciInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrNbrSynciInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrNbrSynciInfo proto.InternalMessageInfo

func (m *LdpNsrNbrSynciInfo) GetInitSyncStart() uint32 {
	if m != nil {
		return m.InitSyncStart
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetInitSyncEnd() uint32 {
	if m != nil {
		return m.InitSyncEnd
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumAddr() uint32 {
	if m != nil {
		return m.NumAddr
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumDuplicateAddr() uint32 {
	if m != nil {
		return m.NumDuplicateAddr
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumRxBytes() uint32 {
	if m != nil {
		return m.NumRxBytes
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumCapSent() uint32 {
	if m != nil {
		return m.NumCapSent
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumCapRcvd() uint32 {
	if m != nil {
		return m.NumCapRcvd
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumLbl() uint32 {
	if m != nil {
		return m.NumLbl
	}
	return 0
}

func (m *LdpNsrNbrSynciInfo) GetNumAppBytes() uint32 {
	if m != nil {
		return m.NumAppBytes
	}
	return 0
}

type LdpNsrNbrSyncsInfo struct {
	NumCapSent           uint32   `protobuf:"varint,1,opt,name=num_cap_sent,json=numCapSent,proto3" json:"num_cap_sent,omitempty"`
	NumCapRcvd           uint32   `protobuf:"varint,2,opt,name=num_cap_rcvd,json=numCapRcvd,proto3" json:"num_cap_rcvd,omitempty"`
	RemLblWd             uint32   `protobuf:"varint,3,opt,name=rem_lbl_wd,json=remLblWd,proto3" json:"rem_lbl_wd,omitempty"`
	RemLblRq             uint32   `protobuf:"varint,4,opt,name=rem_lbl_rq,json=remLblRq,proto3" json:"rem_lbl_rq,omitempty"`
	NumStdbyAdjJoin      uint32   `protobuf:"varint,5,opt,name=num_stdby_adj_join,json=numStdbyAdjJoin,proto3" json:"num_stdby_adj_join,omitempty"`
	NumStdbyAdjLeave     uint32   `protobuf:"varint,6,opt,name=num_stdby_adj_leave,json=numStdbyAdjLeave,proto3" json:"num_stdby_adj_leave,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrNbrSyncsInfo) Reset()         { *m = LdpNsrNbrSyncsInfo{} }
func (m *LdpNsrNbrSyncsInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNsrNbrSyncsInfo) ProtoMessage()    {}
func (*LdpNsrNbrSyncsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b668fae369a575f, []int{2}
}

func (m *LdpNsrNbrSyncsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrNbrSyncsInfo.Unmarshal(m, b)
}
func (m *LdpNsrNbrSyncsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrNbrSyncsInfo.Marshal(b, m, deterministic)
}
func (m *LdpNsrNbrSyncsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrNbrSyncsInfo.Merge(m, src)
}
func (m *LdpNsrNbrSyncsInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNsrNbrSyncsInfo.Size(m)
}
func (m *LdpNsrNbrSyncsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrNbrSyncsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrNbrSyncsInfo proto.InternalMessageInfo

func (m *LdpNsrNbrSyncsInfo) GetNumCapSent() uint32 {
	if m != nil {
		return m.NumCapSent
	}
	return 0
}

func (m *LdpNsrNbrSyncsInfo) GetNumCapRcvd() uint32 {
	if m != nil {
		return m.NumCapRcvd
	}
	return 0
}

func (m *LdpNsrNbrSyncsInfo) GetRemLblWd() uint32 {
	if m != nil {
		return m.RemLblWd
	}
	return 0
}

func (m *LdpNsrNbrSyncsInfo) GetRemLblRq() uint32 {
	if m != nil {
		return m.RemLblRq
	}
	return 0
}

func (m *LdpNsrNbrSyncsInfo) GetNumStdbyAdjJoin() uint32 {
	if m != nil {
		return m.NumStdbyAdjJoin
	}
	return 0
}

func (m *LdpNsrNbrSyncsInfo) GetNumStdbyAdjLeave() uint32 {
	if m != nil {
		return m.NumStdbyAdjLeave
	}
	return 0
}

type LdpNsrStatsNbrInfo struct {
	LsrIdXr              uint32              `protobuf:"varint,50,opt,name=lsr_id_xr,json=lsrIdXr,proto3" json:"lsr_id_xr,omitempty"`
	LblSpcId             uint32              `protobuf:"varint,51,opt,name=lbl_spc_id,json=lblSpcId,proto3" json:"lbl_spc_id,omitempty"`
	NsrSyncState         int32               `protobuf:"zigzag32,52,opt,name=nsr_sync_state,json=nsrSyncState,proto3" json:"nsr_sync_state,omitempty"`
	NumMsg               uint32              `protobuf:"varint,53,opt,name=num_msg,json=numMsg,proto3" json:"num_msg,omitempty"`
	InitSyncInfo         *LdpNsrNbrSynciInfo `protobuf:"bytes,54,opt,name=init_sync_info,json=initSyncInfo,proto3" json:"init_sync_info,omitempty"`
	SteadyStateSyncInfo  *LdpNsrNbrSyncsInfo `protobuf:"bytes,55,opt,name=steady_state_sync_info,json=steadyStateSyncInfo,proto3" json:"steady_state_sync_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LdpNsrStatsNbrInfo) Reset()         { *m = LdpNsrStatsNbrInfo{} }
func (m *LdpNsrStatsNbrInfo) String() string { return proto.CompactTextString(m) }
func (*LdpNsrStatsNbrInfo) ProtoMessage()    {}
func (*LdpNsrStatsNbrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b668fae369a575f, []int{3}
}

func (m *LdpNsrStatsNbrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrStatsNbrInfo.Unmarshal(m, b)
}
func (m *LdpNsrStatsNbrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrStatsNbrInfo.Marshal(b, m, deterministic)
}
func (m *LdpNsrStatsNbrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrStatsNbrInfo.Merge(m, src)
}
func (m *LdpNsrStatsNbrInfo) XXX_Size() int {
	return xxx_messageInfo_LdpNsrStatsNbrInfo.Size(m)
}
func (m *LdpNsrStatsNbrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrStatsNbrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrStatsNbrInfo proto.InternalMessageInfo

func (m *LdpNsrStatsNbrInfo) GetLsrIdXr() uint32 {
	if m != nil {
		return m.LsrIdXr
	}
	return 0
}

func (m *LdpNsrStatsNbrInfo) GetLblSpcId() uint32 {
	if m != nil {
		return m.LblSpcId
	}
	return 0
}

func (m *LdpNsrStatsNbrInfo) GetNsrSyncState() int32 {
	if m != nil {
		return m.NsrSyncState
	}
	return 0
}

func (m *LdpNsrStatsNbrInfo) GetNumMsg() uint32 {
	if m != nil {
		return m.NumMsg
	}
	return 0
}

func (m *LdpNsrStatsNbrInfo) GetInitSyncInfo() *LdpNsrNbrSynciInfo {
	if m != nil {
		return m.InitSyncInfo
	}
	return nil
}

func (m *LdpNsrStatsNbrInfo) GetSteadyStateSyncInfo() *LdpNsrNbrSyncsInfo {
	if m != nil {
		return m.SteadyStateSyncInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpNsrStatsNbrInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_neighbors.ha_neighbor.ldp_nsr_stats_nbr_info_KEYS")
	proto.RegisterType((*LdpNsrNbrSynciInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_neighbors.ha_neighbor.ldp_nsr_nbr_synci_info")
	proto.RegisterType((*LdpNsrNbrSyncsInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_neighbors.ha_neighbor.ldp_nsr_nbr_syncs_info")
	proto.RegisterType((*LdpNsrStatsNbrInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.issu.ha_statistics.ha_neighbors.ha_neighbor.ldp_nsr_stats_nbr_info")
}

func init() { proto.RegisterFile("ldp_nsr_stats_nbr_info.proto", fileDescriptor_7b668fae369a575f) }

var fileDescriptor_7b668fae369a575f = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x14, 0x85, 0x35, 0x2d, 0x4d, 0x5b, 0x37, 0x69, 0x8b, 0x2b, 0xca, 0x00, 0x5d, 0x44, 0x11, 0x42,
	0x91, 0x80, 0x59, 0xb4, 0xfc, 0xac, 0x0b, 0x74, 0x51, 0x28, 0x9b, 0x99, 0x05, 0x3f, 0x1b, 0xcb,
	0x63, 0xbb, 0xa9, 0x23, 0x8f, 0xc7, 0xf5, 0xf5, 0x84, 0x64, 0xc9, 0xdb, 0xf0, 0x06, 0xbc, 0x06,
	0x0f, 0x83, 0x58, 0x23, 0x7b, 0x66, 0x9a, 0x84, 0x56, 0x62, 0x07, 0xcb, 0xb9, 0xf7, 0xe8, 0xfa,
	0xbb, 0xc7, 0xc7, 0x83, 0x0e, 0x14, 0x37, 0x44, 0x83, 0x25, 0xe0, 0xa8, 0x03, 0xa2, 0x73, 0x4b,
	0xa4, 0x3e, 0x2f, 0x13, 0x63, 0x4b, 0x57, 0x62, 0xc9, 0x24, 0xb0, 0x92, 0xc8, 0x12, 0xc8, 0xd4,
	0x92, 0xc2, 0x28, 0x20, 0x5e, 0x5f, 0x1a, 0x61, 0x93, 0xf6, 0x2b, 0x19, 0xa9, 0x32, 0xa7, 0x2a,
	0xa1, 0xcc, 0xc9, 0x89, 0x48, 0xb8, 0x38, 0xa7, 0x95, 0x72, 0x64, 0x62, 0xcf, 0x13, 0x09, 0x50,
	0x25, 0x17, 0x34, 0xcc, 0x96, 0xe0, 0x24, 0x03, 0xff, 0xa5, 0x85, 0x1c, 0x5d, 0xe4, 0xa5, 0x5d,
	0xfa, 0x18, 0x7c, 0x46, 0x0f, 0x6e, 0x46, 0x21, 0xef, 0x4e, 0x3e, 0x65, 0xf8, 0x0e, 0xea, 0x28,
	0xb0, 0x44, 0xf2, 0x38, 0xea, 0x47, 0xc3, 0xcd, 0x74, 0x4d, 0x81, 0x3d, 0xe5, 0xf8, 0x21, 0xda,
	0x56, 0x34, 0x17, 0x8a, 0x80, 0xa1, 0x4c, 0xf8, 0xf6, 0x4a, 0x3f, 0x1a, 0xf6, 0xd2, 0x6e, 0xa8,
	0x66, 0xbe, 0x78, 0xca, 0x07, 0x3f, 0x56, 0xd0, 0x7e, 0x3b, 0xdc, 0x8f, 0x85, 0x99, 0x66, 0x32,
	0x0c, 0xc7, 0x8f, 0xd0, 0x8e, 0xd4, 0xd2, 0x85, 0x92, 0x3f, 0xd8, 0xba, 0x70, 0x40, 0x2f, 0xed,
	0xf9, 0x72, 0x36, 0xd3, 0x2c, 0xf3, 0x45, 0x3c, 0x40, 0xbd, 0xb9, 0x4e, 0xe8, 0xf6, 0x9c, 0xad,
	0x56, 0x75, 0xa2, 0x39, 0xbe, 0x87, 0x36, 0x74, 0x55, 0x10, 0xca, 0xb9, 0x8d, 0x57, 0x43, 0x7b,
	0x5d, 0x57, 0xc5, 0x31, 0xe7, 0x16, 0x3f, 0x41, 0xd8, 0xb7, 0x78, 0x65, 0x94, 0x64, 0xd4, 0x89,
	0x5a, 0x74, 0x2b, 0x88, 0x76, 0x75, 0x55, 0xbc, 0x69, 0x1b, 0x41, 0xdd, 0x47, 0x5d, 0xaf, 0xb6,
	0x53, 0x92, 0xcf, 0x9c, 0x80, 0x78, 0x2d, 0xe8, 0x90, 0xae, 0x8a, 0x74, 0xfa, 0xca, 0x57, 0x5a,
	0x05, 0xa3, 0x86, 0x80, 0xd0, 0x2e, 0xee, 0x5c, 0x29, 0x5e, 0x53, 0x93, 0x09, 0xed, 0x16, 0x15,
	0x96, 0x4d, 0x78, 0xbc, 0xbe, 0xa8, 0x48, 0xd9, 0x84, 0xe3, 0xbb, 0xc8, 0xe3, 0x11, 0x95, 0xab,
	0x78, 0x23, 0x34, 0x3b, 0xba, 0x2a, 0xce, 0x72, 0xe5, 0x77, 0x0d, 0x7b, 0x18, 0xd3, 0x9c, 0xbf,
	0x59, 0xef, 0xea, 0x97, 0x31, 0x26, 0x00, 0x0c, 0x7e, 0x46, 0xd7, 0x2d, 0x85, 0xda, 0xd2, 0x3f,
	0xd9, 0xa2, 0xbf, 0xb2, 0xad, 0x5c, 0x63, 0x3b, 0x40, 0xc8, 0x8a, 0xc0, 0x46, 0xbe, 0xf0, 0xc6,
	0xcc, 0x0d, 0x2b, 0x3c, 0xde, 0x87, 0xa5, 0xae, 0xbd, 0x6c, 0x5c, 0x6c, 0xba, 0xe9, 0x25, 0x7e,
	0x5c, 0x7b, 0x0d, 0x8e, 0xe7, 0x33, 0x42, 0xf9, 0x98, 0x8c, 0x4b, 0xa9, 0x1b, 0x0f, 0x77, 0x74,
	0x55, 0x64, 0xbe, 0x71, 0xcc, 0xc7, 0x6f, 0x4b, 0xa9, 0xf1, 0x53, 0xb4, 0xb7, 0x2c, 0x56, 0x82,
	0x4e, 0x44, 0xe3, 0xe7, 0xee, 0x82, 0xfa, 0xcc, 0xd7, 0x07, 0xbf, 0x56, 0xe7, 0x6b, 0x2f, 0xc7,
	0x14, 0xdf, 0x47, 0x9b, 0x75, 0x42, 0xc9, 0xd4, 0xc6, 0x87, 0xf5, 0xf5, 0x87, 0x90, 0x7e, 0xb4,
	0x1e, 0xd8, 0xc3, 0x82, 0x61, 0x3e, 0xa2, 0x47, 0x35, 0xb0, 0xca, 0x55, 0x66, 0x58, 0x1d, 0xe2,
	0x30, 0xaf, 0x89, 0xa0, 0x13, 0xf1, 0xb3, 0x7e, 0x34, 0xbc, 0x9d, 0x76, 0x35, 0xd8, 0x26, 0x81,
	0x4e, 0xb4, 0xd7, 0x55, 0xc0, 0x28, 0x7e, 0x7e, 0x75, 0x5d, 0xef, 0x61, 0x84, 0xbf, 0x45, 0x68,
	0x7b, 0x9e, 0x4d, 0xcf, 0x12, 0xbf, 0xe8, 0x47, 0xc3, 0xad, 0xc3, 0xaf, 0x51, 0xf2, 0xcf, 0xde,
	0x6f, 0x72, 0xf3, 0xfb, 0x4a, 0xbb, 0xed, 0x03, 0x39, 0xf5, 0x1e, 0x7d, 0x8f, 0xd0, 0x3e, 0x38,
	0x41, 0xf9, 0xac, 0x5e, 0x74, 0x01, 0xf9, 0xe5, 0xff, 0x47, 0xae, 0xf3, 0x9b, 0xee, 0xd5, 0x84,
	0xc1, 0xf5, 0x96, 0x3c, 0xef, 0x84, 0x1f, 0xe2, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7,
	0x2a, 0x6f, 0xa8, 0x30, 0x05, 0x00, 0x00,
}