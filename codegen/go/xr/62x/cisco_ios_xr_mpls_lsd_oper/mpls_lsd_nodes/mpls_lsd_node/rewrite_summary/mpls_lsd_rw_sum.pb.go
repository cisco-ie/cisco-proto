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

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_rewrite_summary

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
	proto.RegisterType((*MplsLsdRwSum_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary.mpls_lsd_rw_sum_KEYS")
	proto.RegisterType((*MplsLsdRwIpPathSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary.mpls_lsd_rw_ip_path_sum")
	proto.RegisterType((*MplsLsdRwSum)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.rewrite_summary.mpls_lsd_rw_sum")
}

func init() { proto.RegisterFile("mpls_lsd_rw_sum.proto", fileDescriptor_be3f4657d5bf8c9f) }

var fileDescriptor_be3f4657d5bf8c9f = []byte{
	// 645 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xdb, 0x4f, 0x13, 0x41,
	0x14, 0xc6, 0x03, 0x5e, 0x42, 0xa7, 0x02, 0xb2, 0x80, 0x2c, 0x21, 0x51, 0x52, 0x8d, 0x81, 0xa8,
	0xd5, 0x00, 0xae, 0xa2, 0xe2, 0x05, 0x01, 0x25, 0x22, 0x21, 0x05, 0x13, 0x7d, 0x1a, 0xb7, 0xed,
	0x2c, 0x6c, 0xdc, 0xd9, 0x99, 0xcc, 0xa5, 0x8b, 0x4f, 0xbe, 0xfa, 0x2f, 0xf9, 0xdf, 0x99, 0x39,
	0x67, 0xb6, 0xec, 0xd6, 0xf8, 0x66, 0x7c, 0xdc, 0xef, 0xfb, 0x7e, 0x3d, 0x27, 0xa7, 0xe7, 0x0c,
	0x99, 0xe7, 0x32, 0xd3, 0x34, 0xd3, 0x7d, 0xaa, 0x0a, 0xaa, 0x2d, 0x6f, 0x4b, 0x25, 0x8c, 0x08,
	0xde, 0xf5, 0x52, 0xdd, 0x13, 0x34, 0x15, 0x9a, 0x9e, 0x2b, 0x3a, 0xcc, 0x08, 0xc9, 0x54, 0x7b,
	0xf8, 0x95, 0x8b, 0x3e, 0xd3, 0xf5, 0xcf, 0xb6, 0x62, 0x85, 0x4a, 0x0d, 0x73, 0x3f, 0xc5, 0x63,
	0xf5, 0xbd, 0xb5, 0x4e, 0xe6, 0x46, 0x2a, 0xd0, 0x0f, 0xbb, 0x5f, 0x8e, 0x83, 0x25, 0xd2, 0x70,
	0x79, 0x9a, 0xc7, 0x9c, 0x85, 0x63, 0xcb, 0x63, 0x2b, 0x8d, 0xce, 0x84, 0x13, 0x0e, 0x63, 0xce,
	0x5a, 0x3f, 0xc7, 0xc9, 0x42, 0x95, 0x4a, 0x25, 0x95, 0xb1, 0x39, 0x73, 0x74, 0x70, 0x8b, 0x34,
	0x8d, 0x30, 0x71, 0x06, 0x8a, 0x06, 0x74, 0xb2, 0x43, 0x40, 0x3a, 0x72, 0x4a, 0x70, 0x9f, 0x04,
	0x18, 0xe8, 0xc6, 0xbd, 0x6f, 0x56, 0xfa, 0xdc, 0x38, 0xe4, 0xae, 0x83, 0xb3, 0x0d, 0x06, 0xa6,
	0x37, 0xc9, 0x22, 0xa6, 0x15, 0xe3, 0xc2, 0xb0, 0x3a, 0x74, 0x09, 0xa0, 0x1b, 0x10, 0xe8, 0x80,
	0x5f, 0x45, 0xd7, 0xc8, 0xbc, 0xef, 0x44, 0x09, 0xc3, 0x7a, 0x86, 0xf5, 0x3d, 0x76, 0x19, 0xb0,
	0x59, 0xec, 0xa9, 0xf4, 0x90, 0x79, 0x48, 0xe6, 0x3c, 0x23, 0x24, 0xcd, 0x2e, 0x2a, 0x5d, 0x01,
	0x64, 0x06, 0x11, 0x21, 0x0f, 0xca, 0x22, 0xad, 0x5f, 0x13, 0x64, 0x7a, 0x64, 0x80, 0x41, 0x8b,
	0x4c, 0xe6, 0x96, 0xd3, 0x2c, 0xee, 0xb2, 0x8c, 0x26, 0x32, 0x0d, 0xd7, 0x80, 0x6e, 0xe6, 0x96,
	0x1f, 0x38, 0x6d, 0x4f, 0xa6, 0x30, 0x5f, 0xcb, 0xa9, 0x61, 0xce, 0x5f, 0x07, 0x7f, 0x22, 0xb7,
	0xfc, 0xc4, 0x7d, 0xbb, 0x19, 0x3a, 0x33, 0x95, 0x83, 0x0d, 0x67, 0x6f, 0xe0, 0x0c, 0x73, 0xcb,
	0xf7, 0x51, 0xa9, 0x04, 0x22, 0x17, 0x78, 0x5c, 0x0d, 0x38, 0x25, 0x58, 0x25, 0x33, 0x2e, 0x20,
	0x0b, 0x9a, 0xa5, 0xda, 0xf8, 0x32, 0x11, 0xc4, 0xa6, 0x72, 0xcb, 0x8f, 0x8a, 0x83, 0x54, 0x1b,
	0x2c, 0x76, 0x1b, 0xbb, 0xed, 0x73, 0xd3, 0xc3, 0xd8, 0x13, 0x88, 0x5d, 0xcb, 0x2d, 0xdf, 0x29,
	0xb5, 0xb2, 0xa0, 0xdf, 0x9e, 0xf0, 0xe9, 0xb0, 0x60, 0x07, 0x95, 0xe0, 0x51, 0x39, 0xb8, 0x44,
	0xa8, 0x22, 0x56, 0x7d, 0x6a, 0x65, 0x3f, 0x36, 0x2c, 0xdc, 0x84, 0x24, 0xfe, 0xe3, 0x7b, 0x68,
	0x7d, 0x02, 0x27, 0xd8, 0x22, 0x4b, 0x15, 0xa2, 0x04, 0x28, 0x67, 0x5a, 0xc7, 0xa7, 0x2c, 0x7c,
	0x06, 0x60, 0x78, 0x01, 0x7a, 0xee, 0x23, 0xfa, 0xa3, 0x7b, 0xf6, 0xfc, 0x8f, 0x3d, 0xfb, 0x41,
	0x88, 0x1b, 0xa0, 0xf7, 0x5f, 0x2c, 0x8f, 0xad, 0x34, 0xd7, 0xbe, 0xb6, 0xff, 0xd1, 0xdd, 0xb4,
	0xff, 0xb2, 0xfe, 0x9d, 0x86, 0xab, 0x59, 0x6d, 0x20, 0xf2, 0x0d, 0x6c, 0xfd, 0xc7, 0x06, 0x22,
	0x6c, 0x60, 0x95, 0xe0, 0xc2, 0x52, 0x43, 0xd9, 0x70, 0x10, 0x2f, 0x71, 0x09, 0xc0, 0x38, 0xd9,
	0x2d, 0x7b, 0x7d, 0x40, 0x66, 0x7d, 0x94, 0xd1, 0x33, 0x16, 0x97, 0x97, 0xf2, 0xaa, 0x72, 0x95,
	0x27, 0xec, 0x3d, 0x8b, 0xfd, 0x99, 0xdc, 0x21, 0x53, 0x7e, 0xf8, 0x85, 0x4f, 0xbe, 0xc6, 0xa5,
	0xc1, 0xf9, 0x17, 0x98, 0xba, 0x57, 0x5e, 0x7a, 0x2a, 0xa9, 0xb6, 0x5d, 0x9f, 0x7c, 0x03, 0xc9,
	0x69, 0x70, 0xf6, 0xe5, 0xb1, 0xed, 0x62, 0x38, 0x22, 0x0b, 0x65, 0x78, 0xb0, 0xa1, 0x64, 0x42,
	0x73, 0x96, 0x9e, 0x9e, 0x75, 0x85, 0xd2, 0xe1, 0x36, 0x10, 0xf3, 0x9e, 0x00, 0xf7, 0xb0, 0x34,
	0x6b, 0x5c, 0x54, 0xe7, 0xde, 0xd6, 0xb9, 0xa8, 0xc6, 0x6d, 0x92, 0xc5, 0xca, 0x46, 0xd3, 0x3a,
	0xb9, 0x83, 0x0f, 0xcb, 0xc5, 0x7e, 0x77, 0xaa, 0xe8, 0x5d, 0x82, 0xdd, 0xc3, 0xcd, 0xd0, 0x34,
	0x37, 0x49, 0xb8, 0x0b, 0xc0, 0x24, 0xc8, 0xee, 0x6a, 0xf6, 0x73, 0x93, 0x04, 0x37, 0x49, 0x73,
	0xa0, 0x12, 0x78, 0x42, 0xe9, 0xb9, 0x0a, 0xf7, 0xe0, 0x15, 0x6d, 0x0c, 0x54, 0xe2, 0x1e, 0xd1,
	0xcf, 0xaa, 0x7b, 0x15, 0xde, 0xf2, 0xf5, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x9a, 0x18,
	0x53, 0xe4, 0x05, 0x00, 0x00,
}
