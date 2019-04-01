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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_vrfs_vrf_nsr_nsr_pending_ha_neighbors_ha_neighbor

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
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	LsrId                string   `protobuf:"bytes,3,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,4,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
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

func (m *LdpNsrStatsNbrInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpNsrStatsNbrInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*LdpNsrStatsNbrInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.nsr.nsr_pending.ha_neighbors.ha_neighbor.ldp_nsr_stats_nbr_info_KEYS")
	proto.RegisterType((*LdpNsrNbrSynciInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.nsr.nsr_pending.ha_neighbors.ha_neighbor.ldp_nsr_nbr_synci_info")
	proto.RegisterType((*LdpNsrNbrSyncsInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.nsr.nsr_pending.ha_neighbors.ha_neighbor.ldp_nsr_nbr_syncs_info")
	proto.RegisterType((*LdpNsrStatsNbrInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.vrfs.vrf.nsr.nsr_pending.ha_neighbors.ha_neighbor.ldp_nsr_stats_nbr_info")
}

func init() { proto.RegisterFile("ldp_nsr_stats_nbr_info.proto", fileDescriptor_7b668fae369a575f) }

var fileDescriptor_7b668fae369a575f = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x4f, 0x6e, 0xd4, 0x30,
	0x18, 0xc5, 0x95, 0x96, 0x4e, 0x67, 0xdc, 0x99, 0xb6, 0xb8, 0xa2, 0x04, 0xda, 0xc5, 0x68, 0x84,
	0xd0, 0x48, 0x40, 0x16, 0x2d, 0x7f, 0xd6, 0x05, 0xba, 0x28, 0x14, 0x16, 0xc9, 0x02, 0x58, 0x19,
	0x27, 0xf6, 0x4c, 0x53, 0x39, 0x8e, 0x6b, 0x3b, 0x61, 0xe6, 0x1a, 0x1c, 0x80, 0x1b, 0x70, 0x17,
	0xce, 0xc0, 0x19, 0x38, 0x00, 0xfa, 0xec, 0xa4, 0x9d, 0xa1, 0x95, 0x58, 0x21, 0x16, 0x8d, 0x94,
	0xef, 0x3d, 0xa5, 0xbf, 0xef, 0xf9, 0x79, 0xd0, 0xbe, 0x60, 0x8a, 0x48, 0xa3, 0x89, 0xb1, 0xd4,
	0x1a, 0x22, 0x53, 0x4d, 0x72, 0x39, 0x29, 0x23, 0xa5, 0x4b, 0x5b, 0xe2, 0xcf, 0x59, 0x6e, 0xb2,
	0x92, 0xe4, 0xa5, 0x21, 0x33, 0x4d, 0x0a, 0x25, 0x0c, 0x01, 0x7f, 0xa9, 0xb8, 0x8e, 0xda, 0xb7,
	0x48, 0x96, 0x8c, 0x1b, 0xf7, 0x8c, 0x6a, 0x3d, 0x31, 0xf0, 0x88, 0xa4, 0xd1, 0xf0, 0x47, 0x14,
	0x97, 0x2c, 0x97, 0xd3, 0xe8, 0x8c, 0x12, 0xc9, 0xf3, 0xe9, 0x59, 0x5a, 0x6a, 0xb3, 0xf8, 0x32,
	0xfa, 0x1a, 0xa0, 0xbd, 0x9b, 0x11, 0xc8, 0xdb, 0xe3, 0x4f, 0x09, 0xde, 0x43, 0x3d, 0xf8, 0x32,
	0x91, 0xb4, 0xe0, 0x61, 0x30, 0x0c, 0xc6, 0xbd, 0xb8, 0x0b, 0x83, 0xf7, 0xb4, 0xe0, 0xf8, 0x1e,
	0xea, 0xd6, 0x7a, 0xe2, 0xb5, 0x15, 0xa7, 0xad, 0xd7, 0x7a, 0xe2, 0xa4, 0x3b, 0xa8, 0x23, 0x8c,
	0x26, 0x39, 0x0b, 0x57, 0x9d, 0xb0, 0x26, 0x8c, 0x3e, 0x61, 0xf8, 0x01, 0xda, 0x14, 0x34, 0xe5,
	0x82, 0x18, 0x45, 0x33, 0x0e, 0xf2, 0xad, 0x61, 0x30, 0x1e, 0xc4, 0x7d, 0x37, 0x4d, 0x60, 0x78,
	0xc2, 0x46, 0x3f, 0x56, 0xd0, 0x6e, 0x0b, 0x05, 0x38, 0x66, 0x2e, 0xb3, 0xdc, 0x41, 0xe1, 0x87,
	0x68, 0x2b, 0x97, 0xb9, 0x75, 0x23, 0x00, 0xd6, 0xd6, 0x51, 0x0d, 0xe2, 0x01, 0x8c, 0x93, 0xb9,
	0xcc, 0x12, 0x18, 0xe2, 0x11, 0x1a, 0x5c, 0xf9, 0xb8, 0x64, 0x8e, 0x6f, 0x10, 0x6f, 0xb4, 0xae,
	0x63, 0xc9, 0x00, 0x5f, 0x56, 0x05, 0xa1, 0x8c, 0x69, 0x47, 0x39, 0x88, 0xd7, 0x65, 0x55, 0x1c,
	0x31, 0xa6, 0xf1, 0x63, 0x84, 0x41, 0x62, 0x95, 0x12, 0x79, 0x46, 0x2d, 0xf7, 0x26, 0xcf, 0xba,
	0x2d, 0xab, 0xe2, 0x75, 0x2b, 0x38, 0xf7, 0x10, 0xf5, 0xc1, 0xad, 0x67, 0x24, 0x9d, 0x5b, 0x6e,
	0xc2, 0x35, 0xe7, 0x43, 0xb2, 0x2a, 0xe2, 0xd9, 0x4b, 0x98, 0xb4, 0x8e, 0x8c, 0x2a, 0x62, 0xb8,
	0xb4, 0x61, 0xe7, 0xd2, 0xf1, 0x8a, 0xaa, 0x84, 0x4b, 0xbb, 0xe8, 0xd0, 0x59, 0xcd, 0xc2, 0xf5,
	0x45, 0x47, 0x9c, 0xd5, 0x0c, 0xdf, 0x45, 0x80, 0x47, 0x44, 0x2a, 0xc2, 0xae, 0x13, 0x3b, 0xb2,
	0x2a, 0x4e, 0x53, 0x01, 0xbb, 0xba, 0x3d, 0x94, 0x6a, 0xfe, 0x7f, 0xcf, 0xef, 0x0a, 0xcb, 0x28,
	0xe5, 0x00, 0x46, 0xbf, 0x82, 0xeb, 0x91, 0x1a, 0x1f, 0xe9, 0x9f, 0x6c, 0xc1, 0x5f, 0xd9, 0x56,
	0xae, 0xb1, 0xed, 0x23, 0xa4, 0xb9, 0x63, 0x23, 0x5f, 0x58, 0x13, 0x66, 0x57, 0x73, 0xc0, 0xfb,
	0xb0, 0xa4, 0xea, 0x8b, 0x26, 0xc5, 0x46, 0x8d, 0x2f, 0xf0, 0x23, 0x9f, 0xb5, 0xb1, 0x2c, 0x9d,
	0x13, 0xca, 0xce, 0xc9, 0x79, 0x99, 0xcb, 0x26, 0xc3, 0x2d, 0x59, 0x15, 0x09, 0x08, 0x47, 0xec,
	0xfc, 0x4d, 0x99, 0x4b, 0xfc, 0x04, 0xed, 0x2c, 0x9b, 0x05, 0xa7, 0x35, 0x6f, 0xf2, 0xdc, 0x5e,
	0x70, 0x9f, 0xc2, 0x7c, 0xf4, 0x73, 0xf5, 0x6a, 0xed, 0xe5, 0x7a, 0xe3, 0xfb, 0xa8, 0xe7, 0x1b,
	0x4a, 0x66, 0x3a, 0x3c, 0xf0, 0xc7, 0xef, 0x4a, 0xfa, 0x51, 0x03, 0x30, 0xc0, 0x1a, 0x95, 0x41,
	0x45, 0x0f, 0x3d, 0xb0, 0x48, 0x45, 0xa2, 0x32, 0x5f, 0x62, 0xf7, 0xbd, 0xa6, 0x82, 0x96, 0x87,
	0x4f, 0x87, 0xc1, 0xf8, 0x76, 0xdc, 0x97, 0x46, 0x37, 0x0d, 0xb4, 0xbc, 0x3d, 0xae, 0xc2, 0x4c,
	0xc3, 0x67, 0x97, 0xc7, 0xf5, 0xce, 0x4c, 0xf1, 0xb7, 0x00, 0x6d, 0x5e, 0x75, 0x13, 0x58, 0xc2,
	0xe7, 0xc3, 0x60, 0xbc, 0x71, 0x30, 0x8b, 0xfe, 0xf5, 0x75, 0x8f, 0x6e, 0xbe, 0x55, 0x71, 0xbf,
	0xbd, 0x16, 0x27, 0x90, 0xcc, 0xf7, 0x00, 0xed, 0x1a, 0xcb, 0x29, 0x9b, 0xfb, 0xf5, 0x16, 0x40,
	0x5f, 0xfc, 0x2f, 0x50, 0xdf, 0xd5, 0x78, 0xc7, 0x73, 0xb9, 0x84, 0x5b, 0xde, 0xb4, 0xe3, 0x7e,
	0x2c, 0x0f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x86, 0x4d, 0x0c, 0x4c, 0x05, 0x00, 0x00,
}