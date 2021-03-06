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
// source: mpls_lsd_rw_sum.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_rewrite_summary_vrfs_rewrite_summary_vrf

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

type MplsLsdRwSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdRwSum_KEYS) Reset()         { *m = MplsLsdRwSum_KEYS{} }
func (m *MplsLsdRwSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdRwSum_KEYS) ProtoMessage()    {}
func (*MplsLsdRwSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3f4657d5bf8c9f, []int{0}
}

func (m *MplsLsdRwSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdRwSum_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdRwSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdRwSum_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdRwSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdRwSum_KEYS.Merge(m, src)
}
func (m *MplsLsdRwSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdRwSum_KEYS.Size(m)
}
func (m *MplsLsdRwSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdRwSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdRwSum_KEYS proto.InternalMessageInfo

func (m *MplsLsdRwSum_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *MplsLsdRwSum_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type MplsLsdRwIpPathSum struct {
	TotalPaths             uint32   `protobuf:"varint,1,opt,name=total_paths,json=totalPaths,proto3" json:"total_paths,omitempty"`
	TotalBackupPaths       uint32   `protobuf:"varint,2,opt,name=total_backup_paths,json=totalBackupPaths,proto3" json:"total_backup_paths,omitempty"`
	TotalRemoteBackupPaths uint32   `protobuf:"varint,3,opt,name=total_remote_backup_paths,json=totalRemoteBackupPaths,proto3" json:"total_remote_backup_paths,omitempty"`
	TotalProtectedPaths    uint32   `protobuf:"varint,4,opt,name=total_protected_paths,json=totalProtectedPaths,proto3" json:"total_protected_paths,omitempty"`
	TotalPopLkupPaths      uint32   `protobuf:"varint,5,opt,name=total_pop_lkup_paths,json=totalPopLkupPaths,proto3" json:"total_pop_lkup_paths,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *MplsLsdRwIpPathSum) Reset()         { *m = MplsLsdRwIpPathSum{} }
func (m *MplsLsdRwIpPathSum) String() string { return proto.CompactTextString(m) }
func (*MplsLsdRwIpPathSum) ProtoMessage()    {}
func (*MplsLsdRwIpPathSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3f4657d5bf8c9f, []int{1}
}

func (m *MplsLsdRwIpPathSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdRwIpPathSum.Unmarshal(m, b)
}
func (m *MplsLsdRwIpPathSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdRwIpPathSum.Marshal(b, m, deterministic)
}
func (m *MplsLsdRwIpPathSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdRwIpPathSum.Merge(m, src)
}
func (m *MplsLsdRwIpPathSum) XXX_Size() int {
	return xxx_messageInfo_MplsLsdRwIpPathSum.Size(m)
}
func (m *MplsLsdRwIpPathSum) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdRwIpPathSum.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdRwIpPathSum proto.InternalMessageInfo

func (m *MplsLsdRwIpPathSum) GetTotalPaths() uint32 {
	if m != nil {
		return m.TotalPaths
	}
	return 0
}

func (m *MplsLsdRwIpPathSum) GetTotalBackupPaths() uint32 {
	if m != nil {
		return m.TotalBackupPaths
	}
	return 0
}

func (m *MplsLsdRwIpPathSum) GetTotalRemoteBackupPaths() uint32 {
	if m != nil {
		return m.TotalRemoteBackupPaths
	}
	return 0
}

func (m *MplsLsdRwIpPathSum) GetTotalProtectedPaths() uint32 {
	if m != nil {
		return m.TotalProtectedPaths
	}
	return 0
}

func (m *MplsLsdRwIpPathSum) GetTotalPopLkupPaths() uint32 {
	if m != nil {
		return m.TotalPopLkupPaths
	}
	return 0
}

type MplsLsdRwSum struct {
	NumLabelFpi              uint32              `protobuf:"varint,50,opt,name=num_label_fpi,json=numLabelFpi,proto3" json:"num_label_fpi,omitempty"`
	NumTefpi                 uint32              `protobuf:"varint,51,opt,name=num_tefpi,json=numTefpi,proto3" json:"num_tefpi,omitempty"`
	NumIpv4Fpi               uint32              `protobuf:"varint,52,opt,name=num_ipv4fpi,json=numIpv4fpi,proto3" json:"num_ipv4fpi,omitempty"`
	NumIpv6Fpi               uint32              `protobuf:"varint,53,opt,name=num_ipv6fpi,json=numIpv6fpi,proto3" json:"num_ipv6fpi,omitempty"`
	NumPwListTefpi           uint32              `protobuf:"varint,54,opt,name=num_pw_list_tefpi,json=numPwListTefpi,proto3" json:"num_pw_list_tefpi,omitempty"`
	NumDmtctefpi             uint32              `protobuf:"varint,55,opt,name=num_dmtctefpi,json=numDmtctefpi,proto3" json:"num_dmtctefpi,omitempty"`
	NumRewrite               uint32              `protobuf:"varint,56,opt,name=num_rewrite,json=numRewrite,proto3" json:"num_rewrite,omitempty"`
	TotalForwardUpdate       uint32              `protobuf:"varint,57,opt,name=total_forward_update,json=totalForwardUpdate,proto3" json:"total_forward_update,omitempty"`
	TotalForwadUpdateMessage uint32              `protobuf:"varint,58,opt,name=total_forwad_update_message,json=totalForwadUpdateMessage,proto3" json:"total_forwad_update_message,omitempty"`
	TotalPaths               uint32              `protobuf:"varint,59,opt,name=total_paths,json=totalPaths,proto3" json:"total_paths,omitempty"`
	Ipv4Paths                *MplsLsdRwIpPathSum `protobuf:"bytes,60,opt,name=ipv4_paths,json=ipv4Paths,proto3" json:"ipv4_paths,omitempty"`
	Ipv6Paths                *MplsLsdRwIpPathSum `protobuf:"bytes,61,opt,name=ipv6_paths,json=ipv6Paths,proto3" json:"ipv6_paths,omitempty"`
	TotalTEv4Paths           uint32              `protobuf:"varint,62,opt,name=total_t_ev4_paths,json=totalTEv4Paths,proto3" json:"total_t_ev4_paths,omitempty"`
	TotalTeHeadPaths         uint32              `protobuf:"varint,63,opt,name=total_te_head_paths,json=totalTeHeadPaths,proto3" json:"total_te_head_paths,omitempty"`
	TotalPwPaths             uint32              `protobuf:"varint,64,opt,name=total_pw_paths,json=totalPwPaths,proto3" json:"total_pw_paths,omitempty"`
	TotalIpSubPaths          uint32              `protobuf:"varint,65,opt,name=total_ip_sub_paths,json=totalIpSubPaths,proto3" json:"total_ip_sub_paths,omitempty"`
	TotalIpv4RpfNeighbors    uint32              `protobuf:"varint,66,opt,name=total_ipv4rpf_neighbors,json=totalIpv4rpfNeighbors,proto3" json:"total_ipv4rpf_neighbors,omitempty"`
	TotalIpv6RpfNeighbors    uint32              `protobuf:"varint,67,opt,name=total_ipv6rpf_neighbors,json=totalIpv6rpfNeighbors,proto3" json:"total_ipv6rpf_neighbors,omitempty"`
	NumRewriteRpfNeighbors   uint32              `protobuf:"varint,68,opt,name=num_rewrite_rpf_neighbors,json=numRewriteRpfNeighbors,proto3" json:"num_rewrite_rpf_neighbors,omitempty"`
	TotalDmtcIntf            uint32              `protobuf:"varint,69,opt,name=total_dmtc_intf,json=totalDmtcIntf,proto3" json:"total_dmtc_intf,omitempty"`
	VrfNameXr                string              `protobuf:"bytes,70,opt,name=vrf_name_xr,json=vrfNameXr,proto3" json:"vrf_name_xr,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}            `json:"-"`
	XXX_unrecognized         []byte              `json:"-"`
	XXX_sizecache            int32               `json:"-"`
}

func (m *MplsLsdRwSum) Reset()         { *m = MplsLsdRwSum{} }
func (m *MplsLsdRwSum) String() string { return proto.CompactTextString(m) }
func (*MplsLsdRwSum) ProtoMessage()    {}
func (*MplsLsdRwSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be3f4657d5bf8c9f, []int{2}
}

func (m *MplsLsdRwSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdRwSum.Unmarshal(m, b)
}
func (m *MplsLsdRwSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdRwSum.Marshal(b, m, deterministic)
}
func (m *MplsLsdRwSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdRwSum.Merge(m, src)
}
func (m *MplsLsdRwSum) XXX_Size() int {
	return xxx_messageInfo_MplsLsdRwSum.Size(m)
}
func (m *MplsLsdRwSum) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdRwSum.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdRwSum proto.InternalMessageInfo

func (m *MplsLsdRwSum) GetNumLabelFpi() uint32 {
	if m != nil {
		return m.NumLabelFpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumTefpi() uint32 {
	if m != nil {
		return m.NumTefpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumIpv4Fpi() uint32 {
	if m != nil {
		return m.NumIpv4Fpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumIpv6Fpi() uint32 {
	if m != nil {
		return m.NumIpv6Fpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumPwListTefpi() uint32 {
	if m != nil {
		return m.NumPwListTefpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumDmtctefpi() uint32 {
	if m != nil {
		return m.NumDmtctefpi
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumRewrite() uint32 {
	if m != nil {
		return m.NumRewrite
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalForwardUpdate() uint32 {
	if m != nil {
		return m.TotalForwardUpdate
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalForwadUpdateMessage() uint32 {
	if m != nil {
		return m.TotalForwadUpdateMessage
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalPaths() uint32 {
	if m != nil {
		return m.TotalPaths
	}
	return 0
}

func (m *MplsLsdRwSum) GetIpv4Paths() *MplsLsdRwIpPathSum {
	if m != nil {
		return m.Ipv4Paths
	}
	return nil
}

func (m *MplsLsdRwSum) GetIpv6Paths() *MplsLsdRwIpPathSum {
	if m != nil {
		return m.Ipv6Paths
	}
	return nil
}

func (m *MplsLsdRwSum) GetTotalTEv4Paths() uint32 {
	if m != nil {
		return m.TotalTEv4Paths
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalTeHeadPaths() uint32 {
	if m != nil {
		return m.TotalTeHeadPaths
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalPwPaths() uint32 {
	if m != nil {
		return m.TotalPwPaths
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalIpSubPaths() uint32 {
	if m != nil {
		return m.TotalIpSubPaths
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalIpv4RpfNeighbors() uint32 {
	if m != nil {
		return m.TotalIpv4RpfNeighbors
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalIpv6RpfNeighbors() uint32 {
	if m != nil {
		return m.TotalIpv6RpfNeighbors
	}
	return 0
}

func (m *MplsLsdRwSum) GetNumRewriteRpfNeighbors() uint32 {
	if m != nil {
		return m.NumRewriteRpfNeighbors
	}
	return 0
}

func (m *MplsLsdRwSum) GetTotalDmtcIntf() uint32 {
	if m != nil {
		return m.TotalDmtcIntf
	}
	return 0
}

func (m *MplsLsdRwSum) GetVrfNameXr() string {
	if m != nil {
		return m.VrfNameXr
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsLsdRwSum_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary_vrfs.rewrite_summary_vrf.mpls_lsd_rw_sum_KEYS")
	proto.RegisterType((*MplsLsdRwIpPathSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary_vrfs.rewrite_summary_vrf.mpls_lsd_rw_ip_path_sum")
	proto.RegisterType((*MplsLsdRwSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary_vrfs.rewrite_summary_vrf.mpls_lsd_rw_sum")
}

func init() { proto.RegisterFile("mpls_lsd_rw_sum.proto", fileDescriptor_be3f4657d5bf8c9f) }

var fileDescriptor_be3f4657d5bf8c9f = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0x6d, 0x4f, 0xd4, 0x40,
	0x10, 0xc7, 0x73, 0xf8, 0x74, 0xb7, 0x27, 0x20, 0x05, 0xa4, 0x84, 0x44, 0xc9, 0x69, 0x0c, 0x44,
	0x3d, 0x0d, 0x60, 0x15, 0x15, 0x1f, 0x10, 0x88, 0x44, 0x24, 0xe4, 0xc0, 0x44, 0x5f, 0xad, 0xbd,
	0xeb, 0x16, 0x1a, 0xbb, 0xdd, 0xcd, 0xee, 0xf6, 0x0a, 0xdf, 0xc0, 0x8f, 0xe0, 0x97, 0xf4, 0x3b,
	0x98, 0x9d, 0xd9, 0x1e, 0x2d, 0xea, 0x4b, 0x13, 0x5f, 0x76, 0xfe, 0xff, 0xdf, 0xce, 0xdc, 0xec,
	0xcc, 0x1e, 0x99, 0xe5, 0x32, 0xd5, 0x34, 0xd5, 0x11, 0x55, 0x05, 0xd5, 0x39, 0xef, 0x4a, 0x25,
	0x8c, 0xf0, 0xbe, 0x0e, 0x12, 0x3d, 0x10, 0x34, 0x11, 0x9a, 0x9e, 0x2a, 0x3a, 0xf2, 0x08, 0xc9,
	0x54, 0x77, 0xf4, 0x95, 0x89, 0x88, 0xe9, 0xfa, 0x67, 0x57, 0xb1, 0x42, 0x25, 0x86, 0xd9, 0xa3,
	0x78, 0xa8, 0xce, 0xe8, 0x50, 0xc5, 0xfa, 0x4f, 0xc1, 0xce, 0x3e, 0x99, 0xb9, 0x90, 0x9a, 0x7e,
	0xd8, 0xfe, 0x72, 0xe8, 0x2d, 0x90, 0x96, 0x3d, 0x88, 0x66, 0x21, 0x67, 0x7e, 0x63, 0xb1, 0xb1,
	0xd4, 0xea, 0x35, 0x6d, 0x60, 0x3f, 0xe4, 0xcc, 0x9b, 0x27, 0xcd, 0xa1, 0x8a, 0x51, 0x1b, 0x03,
	0xed, 0xda, 0x50, 0xc5, 0x56, 0xea, 0x7c, 0x1f, 0x23, 0x73, 0xd5, 0x03, 0x13, 0x49, 0x65, 0x68,
	0x4e, 0xec, 0xc1, 0xde, 0x6d, 0xd2, 0x36, 0xc2, 0x84, 0x29, 0x44, 0x34, 0x9c, 0x3a, 0xde, 0x23,
	0x10, 0x3a, 0xb0, 0x11, 0xef, 0x01, 0xf1, 0xd0, 0xd0, 0x0f, 0x07, 0xdf, 0x72, 0xe9, 0x7c, 0x63,
	0xe0, 0xbb, 0x01, 0xca, 0x26, 0x08, 0xe8, 0x5e, 0x27, 0xf3, 0xe8, 0x56, 0x8c, 0x0b, 0xc3, 0xea,
	0xd0, 0x25, 0x80, 0x6e, 0x82, 0xa1, 0x07, 0x7a, 0x15, 0x5d, 0x21, 0xb3, 0xae, 0x12, 0x25, 0x0c,
	0x1b, 0x18, 0x16, 0x39, 0xec, 0x32, 0x60, 0xd3, 0x58, 0x53, 0xa9, 0x21, 0xf3, 0x88, 0xcc, 0x38,
	0x46, 0x48, 0x9a, 0x9e, 0x67, 0xba, 0x02, 0xc8, 0x14, 0x22, 0x42, 0xee, 0x95, 0x49, 0x3a, 0x3f,
	0x9b, 0x64, 0xf2, 0x42, 0x6f, 0xbd, 0x0e, 0x19, 0xcf, 0x72, 0x4e, 0xd3, 0xb0, 0xcf, 0x52, 0x1a,
	0xcb, 0xc4, 0x5f, 0x01, 0xba, 0x9d, 0xe5, 0x7c, 0xcf, 0xc6, 0x76, 0x64, 0x02, 0xad, 0xcf, 0x39,
	0x35, 0xcc, 0xea, 0xab, 0xa0, 0x37, 0xb3, 0x9c, 0x1f, 0xd9, 0x6f, 0xdb, 0x43, 0x2b, 0x26, 0x72,
	0xb8, 0x66, 0xe5, 0x35, 0xec, 0x61, 0x96, 0xf3, 0x5d, 0x8c, 0x54, 0x0c, 0x81, 0x35, 0x3c, 0xa9,
	0x1a, 0x6c, 0xc4, 0x5b, 0x26, 0x53, 0xd6, 0x20, 0x0b, 0x9a, 0x26, 0xda, 0xb8, 0x34, 0x01, 0xd8,
	0x26, 0xb2, 0x9c, 0x1f, 0x14, 0x7b, 0x89, 0x36, 0x98, 0xec, 0x0e, 0x56, 0x1b, 0x71, 0x33, 0x40,
	0xdb, 0x53, 0xb0, 0x5d, 0xcf, 0x72, 0xbe, 0x55, 0xc6, 0xca, 0x84, 0x6e, 0xb8, 0xfc, 0x67, 0xa3,
	0x84, 0x3d, 0x8c, 0x78, 0x8f, 0xcb, 0xc6, 0xc5, 0x42, 0x15, 0xa1, 0x8a, 0x68, 0x2e, 0xa3, 0xd0,
	0x30, 0x7f, 0x1d, 0x9c, 0x78, 0xe3, 0x3b, 0x28, 0x7d, 0x02, 0xc5, 0xdb, 0x20, 0x0b, 0x15, 0xa2,
	0x04, 0x28, 0x67, 0x5a, 0x87, 0xc7, 0xcc, 0x7f, 0x0e, 0xa0, 0x7f, 0x0e, 0x3a, 0xee, 0x23, 0xea,
	0x17, 0xe7, 0xec, 0xc5, 0x6f, 0x73, 0xf6, 0xa3, 0x41, 0x88, 0xed, 0xa0, 0x33, 0xbc, 0x5c, 0x6c,
	0x2c, 0xb5, 0x57, 0xce, 0xba, 0xff, 0x7a, 0xd9, 0xba, 0x7f, 0x59, 0x8c, 0x5e, 0xcb, 0x16, 0x53,
	0x2b, 0x2d, 0x70, 0xa5, 0x6d, 0xfc, 0x0f, 0xa5, 0x05, 0x58, 0xda, 0x32, 0xc1, 0x21, 0xa7, 0x86,
	0xb2, 0x51, 0xef, 0x5e, 0xe1, 0xe0, 0x80, 0x70, 0xb4, 0x5d, 0xfe, 0x8a, 0x87, 0x64, 0xda, 0x59,
	0x19, 0x3d, 0x61, 0x61, 0xb9, 0x5d, 0xaf, 0x2b, 0x9b, 0x7c, 0xc4, 0xde, 0xb3, 0xd0, 0xad, 0xd6,
	0x5d, 0x32, 0xe1, 0x2e, 0xac, 0x70, 0xce, 0x37, 0x38, 0x68, 0x78, 0x67, 0x05, 0xba, 0xee, 0x97,
	0xaf, 0x43, 0x22, 0xa9, 0xce, 0xfb, 0xce, 0xf9, 0x16, 0x9c, 0x93, 0xa0, 0xec, 0xca, 0xc3, 0xbc,
	0x8f, 0xe6, 0x80, 0xcc, 0x95, 0xe6, 0xe1, 0x9a, 0x92, 0x31, 0xcd, 0x58, 0x72, 0x7c, 0xd2, 0x17,
	0x4a, 0xfb, 0x9b, 0x40, 0xcc, 0x3a, 0x02, 0xd4, 0xfd, 0x52, 0xac, 0x71, 0x41, 0x9d, 0x7b, 0x57,
	0xe7, 0x82, 0x1a, 0xb7, 0x4e, 0xe6, 0x2b, 0x5b, 0x40, 0xeb, 0xe4, 0x16, 0x3e, 0x46, 0xe7, 0x3b,
	0xd1, 0xab, 0xa2, 0xf7, 0x08, 0x56, 0x0f, 0x7b, 0x46, 0x93, 0xcc, 0xc4, 0xfe, 0x36, 0x00, 0xe3,
	0x10, 0xb6, 0x9b, 0xb6, 0x9b, 0x99, 0xd8, 0xbb, 0x45, 0xda, 0xe5, 0xab, 0x4b, 0x4f, 0x95, 0xbf,
	0x03, 0x0f, 0x6f, 0xcb, 0x3d, 0xbc, 0x9f, 0x55, 0xff, 0x2a, 0xfc, 0x67, 0xac, 0xfe, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x25, 0xa6, 0xd7, 0xc1, 0x4c, 0x06, 0x00, 0x00,
}
