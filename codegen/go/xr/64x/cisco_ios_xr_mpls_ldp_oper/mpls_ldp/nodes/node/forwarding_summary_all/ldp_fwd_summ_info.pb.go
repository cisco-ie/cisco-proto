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
// source: ldp_fwd_summ_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_forwarding_summary_all

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

type LdpFwdSummInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpFwdSummInfo_KEYS) Reset()         { *m = LdpFwdSummInfo_KEYS{} }
func (m *LdpFwdSummInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpFwdSummInfo_KEYS) ProtoMessage()    {}
func (*LdpFwdSummInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{0}
}

func (m *LdpFwdSummInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdSummInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpFwdSummInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdSummInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpFwdSummInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdSummInfo_KEYS.Merge(m, src)
}
func (m *LdpFwdSummInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpFwdSummInfo_KEYS.Size(m)
}
func (m *LdpFwdSummInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdSummInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdSummInfo_KEYS proto.InternalMessageInfo

func (m *LdpFwdSummInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type LdpVrfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpVrfInfo) Reset()         { *m = LdpVrfInfo{} }
func (m *LdpVrfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpVrfInfo) ProtoMessage()    {}
func (*LdpVrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{1}
}

func (m *LdpVrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpVrfInfo.Unmarshal(m, b)
}
func (m *LdpVrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpVrfInfo.Marshal(b, m, deterministic)
}
func (m *LdpVrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpVrfInfo.Merge(m, src)
}
func (m *LdpVrfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpVrfInfo.Size(m)
}
func (m *LdpVrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpVrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpVrfInfo proto.InternalMessageInfo

func (m *LdpVrfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LdpVrfInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LdpFwdRwPfxLblSumm struct {
	LabeledPfxs          uint32   `protobuf:"varint,1,opt,name=labeled_pfxs,json=labeledPfxs,proto3" json:"labeled_pfxs,omitempty"`
	LabeledPfxsPartial   uint32   `protobuf:"varint,2,opt,name=labeled_pfxs_partial,json=labeledPfxsPartial,proto3" json:"labeled_pfxs_partial,omitempty"`
	UnlabeledPfxs        uint32   `protobuf:"varint,3,opt,name=unlabeled_pfxs,json=unlabeledPfxs,proto3" json:"unlabeled_pfxs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpFwdRwPfxLblSumm) Reset()         { *m = LdpFwdRwPfxLblSumm{} }
func (m *LdpFwdRwPfxLblSumm) String() string { return proto.CompactTextString(m) }
func (*LdpFwdRwPfxLblSumm) ProtoMessage()    {}
func (*LdpFwdRwPfxLblSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{2}
}

func (m *LdpFwdRwPfxLblSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdRwPfxLblSumm.Unmarshal(m, b)
}
func (m *LdpFwdRwPfxLblSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdRwPfxLblSumm.Marshal(b, m, deterministic)
}
func (m *LdpFwdRwPfxLblSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdRwPfxLblSumm.Merge(m, src)
}
func (m *LdpFwdRwPfxLblSumm) XXX_Size() int {
	return xxx_messageInfo_LdpFwdRwPfxLblSumm.Size(m)
}
func (m *LdpFwdRwPfxLblSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdRwPfxLblSumm.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdRwPfxLblSumm proto.InternalMessageInfo

func (m *LdpFwdRwPfxLblSumm) GetLabeledPfxs() uint32 {
	if m != nil {
		return m.LabeledPfxs
	}
	return 0
}

func (m *LdpFwdRwPfxLblSumm) GetLabeledPfxsPartial() uint32 {
	if m != nil {
		return m.LabeledPfxsPartial
	}
	return 0
}

func (m *LdpFwdRwPfxLblSumm) GetUnlabeledPfxs() uint32 {
	if m != nil {
		return m.UnlabeledPfxs
	}
	return 0
}

type LdpFwdRwPfxSumm struct {
	TotalPfxs            uint32              `protobuf:"varint,1,opt,name=total_pfxs,json=totalPfxs,proto3" json:"total_pfxs,omitempty"`
	EcmpPfxs             uint32              `protobuf:"varint,2,opt,name=ecmp_pfxs,json=ecmpPfxs,proto3" json:"ecmp_pfxs,omitempty"`
	ProtectedPfxs        uint32              `protobuf:"varint,3,opt,name=protected_pfxs,json=protectedPfxs,proto3" json:"protected_pfxs,omitempty"`
	LabeledPfxsAggr      *LdpFwdRwPfxLblSumm `protobuf:"bytes,4,opt,name=labeled_pfxs_aggr,json=labeledPfxsAggr,proto3" json:"labeled_pfxs_aggr,omitempty"`
	LabeledPfxsPrimary   *LdpFwdRwPfxLblSumm `protobuf:"bytes,5,opt,name=labeled_pfxs_primary,json=labeledPfxsPrimary,proto3" json:"labeled_pfxs_primary,omitempty"`
	LabeledPfxsBackup    *LdpFwdRwPfxLblSumm `protobuf:"bytes,6,opt,name=labeled_pfxs_backup,json=labeledPfxsBackup,proto3" json:"labeled_pfxs_backup,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LdpFwdRwPfxSumm) Reset()         { *m = LdpFwdRwPfxSumm{} }
func (m *LdpFwdRwPfxSumm) String() string { return proto.CompactTextString(m) }
func (*LdpFwdRwPfxSumm) ProtoMessage()    {}
func (*LdpFwdRwPfxSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{3}
}

func (m *LdpFwdRwPfxSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdRwPfxSumm.Unmarshal(m, b)
}
func (m *LdpFwdRwPfxSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdRwPfxSumm.Marshal(b, m, deterministic)
}
func (m *LdpFwdRwPfxSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdRwPfxSumm.Merge(m, src)
}
func (m *LdpFwdRwPfxSumm) XXX_Size() int {
	return xxx_messageInfo_LdpFwdRwPfxSumm.Size(m)
}
func (m *LdpFwdRwPfxSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdRwPfxSumm.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdRwPfxSumm proto.InternalMessageInfo

func (m *LdpFwdRwPfxSumm) GetTotalPfxs() uint32 {
	if m != nil {
		return m.TotalPfxs
	}
	return 0
}

func (m *LdpFwdRwPfxSumm) GetEcmpPfxs() uint32 {
	if m != nil {
		return m.EcmpPfxs
	}
	return 0
}

func (m *LdpFwdRwPfxSumm) GetProtectedPfxs() uint32 {
	if m != nil {
		return m.ProtectedPfxs
	}
	return 0
}

func (m *LdpFwdRwPfxSumm) GetLabeledPfxsAggr() *LdpFwdRwPfxLblSumm {
	if m != nil {
		return m.LabeledPfxsAggr
	}
	return nil
}

func (m *LdpFwdRwPfxSumm) GetLabeledPfxsPrimary() *LdpFwdRwPfxLblSumm {
	if m != nil {
		return m.LabeledPfxsPrimary
	}
	return nil
}

func (m *LdpFwdRwPfxSumm) GetLabeledPfxsBackup() *LdpFwdRwPfxLblSumm {
	if m != nil {
		return m.LabeledPfxsBackup
	}
	return nil
}

type LdpFwdRwPathSumm struct {
	TotalPaths           uint32   `protobuf:"varint,1,opt,name=total_paths,json=totalPaths,proto3" json:"total_paths,omitempty"`
	ProtectedPaths       uint32   `protobuf:"varint,2,opt,name=protected_paths,json=protectedPaths,proto3" json:"protected_paths,omitempty"`
	BackupPaths          uint32   `protobuf:"varint,3,opt,name=backup_paths,json=backupPaths,proto3" json:"backup_paths,omitempty"`
	RemoteBackupPaths    uint32   `protobuf:"varint,4,opt,name=remote_backup_paths,json=remoteBackupPaths,proto3" json:"remote_backup_paths,omitempty"`
	LabeledPaths         uint32   `protobuf:"varint,5,opt,name=labeled_paths,json=labeledPaths,proto3" json:"labeled_paths,omitempty"`
	LabeledBackupPaths   uint32   `protobuf:"varint,6,opt,name=labeled_backup_paths,json=labeledBackupPaths,proto3" json:"labeled_backup_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpFwdRwPathSumm) Reset()         { *m = LdpFwdRwPathSumm{} }
func (m *LdpFwdRwPathSumm) String() string { return proto.CompactTextString(m) }
func (*LdpFwdRwPathSumm) ProtoMessage()    {}
func (*LdpFwdRwPathSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{4}
}

func (m *LdpFwdRwPathSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdRwPathSumm.Unmarshal(m, b)
}
func (m *LdpFwdRwPathSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdRwPathSumm.Marshal(b, m, deterministic)
}
func (m *LdpFwdRwPathSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdRwPathSumm.Merge(m, src)
}
func (m *LdpFwdRwPathSumm) XXX_Size() int {
	return xxx_messageInfo_LdpFwdRwPathSumm.Size(m)
}
func (m *LdpFwdRwPathSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdRwPathSumm.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdRwPathSumm proto.InternalMessageInfo

func (m *LdpFwdRwPathSumm) GetTotalPaths() uint32 {
	if m != nil {
		return m.TotalPaths
	}
	return 0
}

func (m *LdpFwdRwPathSumm) GetProtectedPaths() uint32 {
	if m != nil {
		return m.ProtectedPaths
	}
	return 0
}

func (m *LdpFwdRwPathSumm) GetBackupPaths() uint32 {
	if m != nil {
		return m.BackupPaths
	}
	return 0
}

func (m *LdpFwdRwPathSumm) GetRemoteBackupPaths() uint32 {
	if m != nil {
		return m.RemoteBackupPaths
	}
	return 0
}

func (m *LdpFwdRwPathSumm) GetLabeledPaths() uint32 {
	if m != nil {
		return m.LabeledPaths
	}
	return 0
}

func (m *LdpFwdRwPathSumm) GetLabeledBackupPaths() uint32 {
	if m != nil {
		return m.LabeledBackupPaths
	}
	return 0
}

type LdpFwdRwSumm struct {
	Pfxs                 *LdpFwdRwPfxSumm  `protobuf:"bytes,1,opt,name=pfxs,proto3" json:"pfxs,omitempty"`
	Nhs                  *LdpFwdRwPathSumm `protobuf:"bytes,2,opt,name=nhs,proto3" json:"nhs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *LdpFwdRwSumm) Reset()         { *m = LdpFwdRwSumm{} }
func (m *LdpFwdRwSumm) String() string { return proto.CompactTextString(m) }
func (*LdpFwdRwSumm) ProtoMessage()    {}
func (*LdpFwdRwSumm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{5}
}

func (m *LdpFwdRwSumm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdRwSumm.Unmarshal(m, b)
}
func (m *LdpFwdRwSumm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdRwSumm.Marshal(b, m, deterministic)
}
func (m *LdpFwdRwSumm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdRwSumm.Merge(m, src)
}
func (m *LdpFwdRwSumm) XXX_Size() int {
	return xxx_messageInfo_LdpFwdRwSumm.Size(m)
}
func (m *LdpFwdRwSumm) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdRwSumm.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdRwSumm proto.InternalMessageInfo

func (m *LdpFwdRwSumm) GetPfxs() *LdpFwdRwPfxSumm {
	if m != nil {
		return m.Pfxs
	}
	return nil
}

func (m *LdpFwdRwSumm) GetNhs() *LdpFwdRwPathSumm {
	if m != nil {
		return m.Nhs
	}
	return nil
}

type LdpFwdSummInfo struct {
	Vrf                  *LdpVrfInfo   `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	IsLsdBound           bool          `protobuf:"varint,51,opt,name=is_lsd_bound,json=isLsdBound,proto3" json:"is_lsd_bound,omitempty"`
	Fsht                 uint32        `protobuf:"varint,52,opt,name=fsht,proto3" json:"fsht,omitempty"`
	Intfs                uint32        `protobuf:"varint,53,opt,name=intfs,proto3" json:"intfs,omitempty"`
	Lbls                 uint32        `protobuf:"varint,54,opt,name=lbls,proto3" json:"lbls,omitempty"`
	Rws                  *LdpFwdRwSumm `protobuf:"bytes,55,opt,name=rws,proto3" json:"rws,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *LdpFwdSummInfo) Reset()         { *m = LdpFwdSummInfo{} }
func (m *LdpFwdSummInfo) String() string { return proto.CompactTextString(m) }
func (*LdpFwdSummInfo) ProtoMessage()    {}
func (*LdpFwdSummInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d26714c6b631a3e5, []int{6}
}

func (m *LdpFwdSummInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpFwdSummInfo.Unmarshal(m, b)
}
func (m *LdpFwdSummInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpFwdSummInfo.Marshal(b, m, deterministic)
}
func (m *LdpFwdSummInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpFwdSummInfo.Merge(m, src)
}
func (m *LdpFwdSummInfo) XXX_Size() int {
	return xxx_messageInfo_LdpFwdSummInfo.Size(m)
}
func (m *LdpFwdSummInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpFwdSummInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpFwdSummInfo proto.InternalMessageInfo

func (m *LdpFwdSummInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpFwdSummInfo) GetIsLsdBound() bool {
	if m != nil {
		return m.IsLsdBound
	}
	return false
}

func (m *LdpFwdSummInfo) GetFsht() uint32 {
	if m != nil {
		return m.Fsht
	}
	return 0
}

func (m *LdpFwdSummInfo) GetIntfs() uint32 {
	if m != nil {
		return m.Intfs
	}
	return 0
}

func (m *LdpFwdSummInfo) GetLbls() uint32 {
	if m != nil {
		return m.Lbls
	}
	return 0
}

func (m *LdpFwdSummInfo) GetRws() *LdpFwdRwSumm {
	if m != nil {
		return m.Rws
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpFwdSummInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_summ_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_vrf_info")
	proto.RegisterType((*LdpFwdRwPfxLblSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_rw_pfx_lbl_summ")
	proto.RegisterType((*LdpFwdRwPfxSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_rw_pfx_summ")
	proto.RegisterType((*LdpFwdRwPathSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_rw_path_summ")
	proto.RegisterType((*LdpFwdRwSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_rw_summ")
	proto.RegisterType((*LdpFwdSummInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.forwarding_summary_all.ldp_fwd_summ_info")
}

func init() { proto.RegisterFile("ldp_fwd_summ_info.proto", fileDescriptor_d26714c6b631a3e5) }

var fileDescriptor_d26714c6b631a3e5 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcf, 0x6a, 0xdb, 0x4e,
	0x10, 0xc7, 0x91, 0xed, 0x84, 0x78, 0xe4, 0x3f, 0x78, 0x6d, 0x7e, 0x3f, 0x41, 0x28, 0x75, 0x54,
	0x42, 0x7d, 0x12, 0xc5, 0x69, 0xda, 0x73, 0x0d, 0x39, 0xb5, 0x14, 0xa3, 0x40, 0xa1, 0x2d, 0x74,
	0x59, 0x59, 0x92, 0x2d, 0xba, 0xfa, 0xc3, 0xae, 0x1c, 0xbb, 0xaf, 0x90, 0x43, 0xe9, 0xa5, 0x0f,
	0xd1, 0xf7, 0xe8, 0x03, 0xf5, 0x11, 0xca, 0xce, 0xda, 0xce, 0xca, 0xa1, 0x37, 0x27, 0x17, 0x23,
	0x7d, 0xf7, 0xa3, 0x99, 0xef, 0xce, 0xcc, 0xae, 0xe1, 0x7f, 0x1e, 0x16, 0x34, 0x5e, 0x85, 0x54,
	0x2e, 0xd3, 0x94, 0x26, 0x59, 0x9c, 0x7b, 0x85, 0xc8, 0xcb, 0x9c, 0x5c, 0xcd, 0x12, 0x39, 0xcb,
	0x69, 0x92, 0x4b, 0xba, 0x16, 0x34, 0x2d, 0xb8, 0xa4, 0x0a, 0xcd, 0x8b, 0x48, 0x78, 0xdb, 0x37,
	0x2f, 0xcb, 0xc3, 0x48, 0xe2, 0xaf, 0x17, 0xe7, 0x62, 0xc5, 0x44, 0x98, 0x64, 0x73, 0x0c, 0xc5,
	0xc4, 0x37, 0xca, 0x38, 0x77, 0x2f, 0xe1, 0xbf, 0x7b, 0x19, 0xe8, 0xdb, 0xab, 0x8f, 0xd7, 0xe4,
	0x14, 0x9a, 0xea, 0x43, 0x9a, 0xb1, 0x34, 0x72, 0xac, 0xa1, 0x35, 0x6a, 0xfa, 0x27, 0x4a, 0x78,
	0xcf, 0xd2, 0xc8, 0x1d, 0x43, 0x4b, 0x7d, 0x76, 0x23, 0x62, 0xfc, 0x82, 0x10, 0x68, 0x18, 0x1c,
	0x3e, 0x93, 0x0e, 0xd4, 0x92, 0xd0, 0xa9, 0x0d, 0xad, 0x51, 0xdb, 0xaf, 0x25, 0xa1, 0xfb, 0xd3,
	0xba, 0xdb, 0x8d, 0x58, 0xd1, 0x22, 0x5e, 0x53, 0x1e, 0x70, 0x4c, 0x4b, 0xce, 0xa0, 0xc5, 0x59,
	0x10, 0xf1, 0x28, 0x54, 0xba, 0xc4, 0x38, 0x6d, 0xdf, 0xde, 0x68, 0xd3, 0x78, 0x2d, 0xc9, 0x0b,
	0x18, 0x98, 0x08, 0x2d, 0x98, 0x28, 0x13, 0xc6, 0x37, 0x09, 0x88, 0x81, 0x4e, 0xf5, 0x0a, 0x39,
	0x87, 0xce, 0x32, 0xab, 0x84, 0xad, 0x23, 0xdb, 0xde, 0xa9, 0x8a, 0x76, 0x7f, 0x35, 0xa0, 0xbf,
	0xe7, 0x0b, 0x3d, 0x3d, 0x01, 0x28, 0xf3, 0x92, 0x71, 0xd3, 0x51, 0x13, 0x15, 0xf4, 0x73, 0x0a,
	0xcd, 0x68, 0x96, 0x16, 0x7a, 0x55, 0x9b, 0x38, 0x51, 0x02, 0x2e, 0x9e, 0x43, 0x47, 0xb5, 0x29,
	0x9a, 0x95, 0x7b, 0xa9, 0x77, 0x2a, 0x62, 0xb7, 0x16, 0xf4, 0x2a, 0x9b, 0x62, 0xf3, 0xb9, 0x70,
	0x1a, 0x43, 0x6b, 0x64, 0x8f, 0xbf, 0x78, 0x07, 0xe9, 0xb0, 0xf7, 0x8f, 0x92, 0xfb, 0x5d, 0xa3,
	0x06, 0x6f, 0xe6, 0x73, 0x41, 0x7e, 0x58, 0xfb, 0x15, 0x16, 0x89, 0x8a, 0xe2, 0x1c, 0x3d, 0x8a,
	0x9f, 0x4a, 0x07, 0x75, 0x66, 0xf2, 0xdd, 0x82, 0x7e, 0xc5, 0x52, 0xc0, 0x66, 0x5f, 0x97, 0x85,
	0x73, 0xfc, 0x28, 0x8e, 0x7a, 0x86, 0xa3, 0x09, 0x26, 0x76, 0x6f, 0x6b, 0x30, 0x30, 0x71, 0x56,
	0x2e, 0xf4, 0xb0, 0x3c, 0x05, 0x7b, 0x33, 0x2c, 0xac, 0x5c, 0x6c, 0xa7, 0x45, 0xcf, 0xcf, 0x54,
	0x29, 0xe4, 0x39, 0x74, 0x8d, 0x89, 0x40, 0x48, 0x0f, 0xcd, 0xdd, 0xa0, 0x68, 0xf0, 0x0c, 0x5a,
	0x7a, 0x97, 0x1b, 0x4a, 0x0f, 0x8e, 0xad, 0x35, 0x8d, 0x78, 0xd0, 0x17, 0x51, 0x9a, 0x97, 0x11,
	0xad, 0x90, 0x0d, 0x24, 0x7b, 0x7a, 0x69, 0x62, 0xf0, 0xcf, 0xa0, 0xbd, 0xab, 0x22, 0x92, 0x47,
	0x48, 0x6e, 0x8f, 0x9c, 0x86, 0x8c, 0xf3, 0x55, 0x89, 0x7a, 0x5c, 0x39, 0x5f, 0x46, 0x58, 0xf7,
	0x8f, 0x05, 0x5d, 0xa3, 0x18, 0x58, 0x87, 0x0c, 0x1a, 0xbb, 0xe3, 0x62, 0x8f, 0x3f, 0x3d, 0x4c,
	0x87, 0xb0, 0x3b, 0x98, 0x87, 0xa4, 0x50, 0xcf, 0x36, 0xa5, 0xb4, 0xc7, 0x9f, 0x1f, 0x20, 0xdd,
	0xb6, 0xc3, 0xbe, 0xca, 0xe3, 0xfe, 0xae, 0x41, 0xef, 0xde, 0x7d, 0x49, 0x22, 0xa8, 0xdf, 0x88,
	0xd8, 0x19, 0xa3, 0x89, 0xeb, 0x03, 0x9a, 0xd8, 0xde, 0xaf, 0xbe, 0x8a, 0x4f, 0x86, 0xd0, 0x4a,
	0x24, 0xe5, 0x32, 0xa4, 0x41, 0xbe, 0xcc, 0x42, 0xe7, 0x62, 0x68, 0x8d, 0x4e, 0x7c, 0x48, 0xe4,
	0x3b, 0x19, 0x4e, 0x94, 0xa2, 0xae, 0xe1, 0x58, 0x2e, 0x4a, 0xe7, 0x25, 0xf6, 0x0c, 0x9f, 0xc9,
	0x00, 0x8e, 0x92, 0xac, 0x8c, 0xa5, 0x73, 0x89, 0xa2, 0x7e, 0x51, 0x24, 0x0f, 0xb8, 0x74, 0x5e,
	0x69, 0x52, 0x3d, 0x93, 0x05, 0xd4, 0xc5, 0x4a, 0x3a, 0xaf, 0x71, 0x1b, 0x1f, 0x0e, 0x5f, 0x4b,
	0x5d, 0x46, 0xb1, 0x92, 0xc1, 0x31, 0xfe, 0x87, 0x5d, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfe,
	0x6b, 0x10, 0x27, 0xde, 0x06, 0x00, 0x00,
}
