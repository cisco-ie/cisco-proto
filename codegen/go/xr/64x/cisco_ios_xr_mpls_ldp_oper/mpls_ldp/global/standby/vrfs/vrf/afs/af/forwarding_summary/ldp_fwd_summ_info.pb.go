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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_afs_af_forwarding_summary

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
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

func (m *LdpFwdSummInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpFwdSummInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
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
	proto.RegisterType((*LdpFwdSummInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_summ_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_vrf_info")
	proto.RegisterType((*LdpFwdRwPfxLblSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_rw_pfx_lbl_summ")
	proto.RegisterType((*LdpFwdRwPfxSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_rw_pfx_summ")
	proto.RegisterType((*LdpFwdRwPathSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_rw_path_summ")
	proto.RegisterType((*LdpFwdRwSumm)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_rw_summ")
	proto.RegisterType((*LdpFwdSummInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.forwarding_summary.ldp_fwd_summ_info")
}

func init() { proto.RegisterFile("ldp_fwd_summ_info.proto", fileDescriptor_d26714c6b631a3e5) }

var fileDescriptor_d26714c6b631a3e5 = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x34, 0x6d, 0x27, 0x49, 0xab, 0x6c, 0x2b, 0x6a, 0x84, 0x10, 0xa9, 0x51, 0x45,
	0x4f, 0x16, 0x4a, 0xf9, 0x38, 0x53, 0x89, 0x13, 0x15, 0xaa, 0x8c, 0x38, 0x70, 0x5a, 0xad, 0x63,
	0xaf, 0x6b, 0xba, 0xfe, 0xd0, 0xae, 0xf3, 0xd1, 0x03, 0xfc, 0x00, 0x4e, 0x5c, 0xb8, 0x70, 0x40,
	0xfc, 0x2d, 0xee, 0xfc, 0x10, 0xb4, 0xb3, 0x4e, 0xba, 0x6e, 0xc5, 0x31, 0xed, 0x25, 0x5a, 0xbf,
	0x79, 0x3b, 0xf3, 0x76, 0xe6, 0xed, 0x06, 0x0e, 0x44, 0x54, 0x52, 0x3e, 0x8f, 0xa8, 0x9a, 0x66,
	0x19, 0x4d, 0x73, 0x5e, 0xf8, 0xa5, 0x2c, 0xaa, 0x82, 0x7c, 0x9c, 0xa4, 0x6a, 0x52, 0xd0, 0xb4,
	0x50, 0x74, 0x21, 0x69, 0x56, 0x0a, 0x45, 0x35, 0xb5, 0x28, 0x63, 0xe9, 0x2f, 0xbf, 0xfc, 0x44,
	0x14, 0x21, 0x13, 0xbe, 0xaa, 0x58, 0x1e, 0x85, 0x57, 0xfe, 0x4c, 0x72, 0xa5, 0x7f, 0x7c, 0xc6,
	0x95, 0xcf, 0xb8, 0xcf, 0x0b, 0x39, 0x67, 0x32, 0x4a, 0xf3, 0x04, 0xd3, 0x33, 0x79, 0xe5, 0x9d,
	0xc1, 0x83, 0x5b, 0x15, 0xe9, 0xbb, 0xb7, 0x9f, 0x3e, 0x90, 0x87, 0xb0, 0x35, 0x93, 0x9c, 0xe6,
	0x2c, 0x8b, 0x5d, 0x67, 0xe4, 0x1c, 0x6f, 0x07, 0x9b, 0x33, 0xc9, 0xdf, 0xb3, 0x2c, 0x26, 0x07,
	0xb0, 0xc9, 0xea, 0x48, 0x0b, 0x23, 0x5d, 0x86, 0x01, 0x6f, 0x0c, 0x7d, 0x9d, 0x4d, 0xef, 0xd3,
	0x89, 0x08, 0x81, 0x8e, 0xb5, 0x1f, 0xd7, 0x64, 0x07, 0x5a, 0x69, 0x84, 0xfb, 0x06, 0x41, 0x2b,
	0x8d, 0xbc, 0x1f, 0xce, 0xf5, 0xa1, 0xe5, 0x9c, 0x96, 0x7c, 0x41, 0x45, 0x28, 0x50, 0x0d, 0x39,
	0x84, 0xbe, 0x60, 0x61, 0x2c, 0xe2, 0x48, 0xe3, 0x0a, 0xf3, 0x0c, 0x82, 0x5e, 0x8d, 0x9d, 0xf3,
	0x85, 0x22, 0xcf, 0x61, 0xdf, 0xa6, 0xd0, 0x92, 0xc9, 0x2a, 0x65, 0xa2, 0x2e, 0x40, 0x2c, 0xea,
	0xb9, 0x89, 0x90, 0x23, 0xd8, 0x99, 0xe6, 0x8d, 0xb4, 0x6d, 0xe4, 0x0e, 0x56, 0xa8, 0x66, 0x7b,
	0x7f, 0x3a, 0xb0, 0x77, 0x43, 0x17, 0x6a, 0x7a, 0x0c, 0x50, 0x15, 0x15, 0x13, 0xb6, 0xa2, 0x6d,
	0x44, 0x50, 0xcf, 0x23, 0xd8, 0x8e, 0x27, 0x59, 0x69, 0xa2, 0x46, 0xc4, 0x96, 0x06, 0x30, 0x78,
	0x04, 0x3b, 0x7a, 0x9a, 0xf1, 0xa4, 0xba, 0x51, 0x7a, 0x85, 0x22, 0xed, 0xa7, 0x03, 0xc3, 0xc6,
	0xa1, 0x58, 0x92, 0x48, 0xb7, 0x33, 0x72, 0x8e, 0x7b, 0xe3, 0xdc, 0x5f, 0x8b, 0x11, 0xfc, 0xff,
	0x8c, 0x20, 0xd8, 0xb5, 0x7a, 0xf2, 0x26, 0x49, 0x24, 0xf9, 0xed, 0xdc, 0xec, 0xb8, 0x4c, 0x75,
	0x06, 0x77, 0xe3, 0x5e, 0xf4, 0x35, 0x26, 0x6c, 0x94, 0x90, 0x5f, 0x0e, 0xec, 0x35, 0x24, 0x86,
	0x6c, 0x72, 0x39, 0x2d, 0xdd, 0xee, 0xbd, 0x28, 0x1c, 0x5a, 0x0a, 0x4f, 0x51, 0x88, 0xf7, 0xad,
	0x05, 0xfb, 0x36, 0x9d, 0x55, 0x17, 0xc6, 0x5c, 0x4f, 0xa0, 0x57, 0x9b, 0x8b, 0x55, 0x17, 0x4b,
	0x77, 0x19, 0xbf, 0x9d, 0x6b, 0x84, 0x3c, 0x83, 0x5d, 0xcb, 0x41, 0x48, 0x32, 0x26, 0xbb, 0x36,
	0x96, 0x21, 0x1e, 0x42, 0xdf, 0x9c, 0xba, 0x66, 0x19, 0xa3, 0xf5, 0x0c, 0x66, 0x28, 0x3e, 0xec,
	0xc9, 0x38, 0x2b, 0xaa, 0x98, 0x36, 0x98, 0x1d, 0x64, 0x0e, 0x4d, 0xe8, 0xd4, 0xe2, 0x3f, 0x85,
	0xc1, 0xaa, 0xab, 0xc8, 0xdc, 0x40, 0xe6, 0xf2, 0x8a, 0x1a, 0x92, 0x75, 0x1f, 0x1b, 0x59, 0xbb,
	0x8d, 0xfb, 0x68, 0xa5, 0xf5, 0xbe, 0xb7, 0x60, 0xd7, 0x6a, 0x06, 0xf6, 0xe1, 0x2b, 0x74, 0x56,
	0xd7, 0xab, 0x37, 0xfe, 0x7c, 0x37, 0x13, 0xc3, 0x69, 0x61, 0x5d, 0xf2, 0x05, 0xda, 0x79, 0xdd,
	0xda, 0xde, 0xf8, 0xf2, 0x0e, 0xca, 0x2f, 0x1d, 0x10, 0xe8, 0xba, 0xde, 0xdf, 0x16, 0x0c, 0x6f,
	0x3d, 0xcb, 0x64, 0x0a, 0xed, 0x99, 0xe4, 0xee, 0x18, 0x45, 0x4d, 0xd6, 0x28, 0x6a, 0xf9, 0x7e,
	0x07, 0xba, 0x1e, 0x19, 0x41, 0x3f, 0x55, 0x54, 0xa8, 0x88, 0x86, 0xc5, 0x34, 0x8f, 0xdc, 0x93,
	0x91, 0x73, 0xbc, 0x15, 0x40, 0xaa, 0xce, 0x54, 0x74, 0xaa, 0x11, 0xfd, 0xcc, 0x73, 0x75, 0x51,
	0xb9, 0x2f, 0x70, 0xc6, 0xb8, 0x26, 0xfb, 0xb0, 0x91, 0xe6, 0x15, 0x57, 0xee, 0x4b, 0x04, 0xcd,
	0x87, 0x66, 0x8a, 0x50, 0x28, 0xf7, 0x95, 0x61, 0xea, 0x35, 0x59, 0x40, 0x5b, 0xce, 0x95, 0xfb,
	0x1a, 0x8f, 0xc5, 0xd7, 0xdf, 0x6b, 0xd3, 0x66, 0x39, 0x57, 0x61, 0x17, 0xff, 0x5a, 0x4f, 0xfe,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x08, 0xb1, 0xf1, 0x40, 0x75, 0x07, 0x00, 0x00,
}
