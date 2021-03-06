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
// source: mpls_te_autotun_backup_summary_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_auto_tunnel_backup_summary

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

type MplsTeAutotunBackupSummaryInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAutotunBackupSummaryInfo_KEYS) Reset()         { *m = MplsTeAutotunBackupSummaryInfo_KEYS{} }
func (m *MplsTeAutotunBackupSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunBackupSummaryInfo_KEYS) ProtoMessage()    {}
func (*MplsTeAutotunBackupSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d77850068d25b3, []int{0}
}

func (m *MplsTeAutotunBackupSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeAutotunBackupSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeAutotunBackupSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS.Merge(m, src)
}
func (m *MplsTeAutotunBackupSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS.Size(m)
}
func (m *MplsTeAutotunBackupSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunBackupSummaryInfo_KEYS proto.InternalMessageInfo

type MplsTeAutotunBackupSummaryInfo struct {
	Autobackups                                   uint32   `protobuf:"varint,50,opt,name=autobackups,proto3" json:"autobackups,omitempty"`
	UpAutobackups                                 uint32   `protobuf:"varint,51,opt,name=up_autobackups,json=upAutobackups,proto3" json:"up_autobackups,omitempty"`
	DownAutobackups                               uint32   `protobuf:"varint,52,opt,name=down_autobackups,json=downAutobackups,proto3" json:"down_autobackups,omitempty"`
	UnusedAutobackups                             uint32   `protobuf:"varint,53,opt,name=unused_autobackups,json=unusedAutobackups,proto3" json:"unused_autobackups,omitempty"`
	NextHopAutobackups                            uint32   `protobuf:"varint,54,opt,name=next_hop_autobackups,json=nextHopAutobackups,proto3" json:"next_hop_autobackups,omitempty"`
	NextNextHopAutobackups                        uint32   `protobuf:"varint,55,opt,name=next_next_hop_autobackups,json=nextNextHopAutobackups,proto3" json:"next_next_hop_autobackups,omitempty"`
	SrlgStrictAutobackups                         uint32   `protobuf:"varint,56,opt,name=srlg_strict_autobackups,json=srlgStrictAutobackups,proto3" json:"srlg_strict_autobackups,omitempty"`
	SrlgPreferredAutobackups                      uint32   `protobuf:"varint,57,opt,name=srlg_preferred_autobackups,json=srlgPreferredAutobackups,proto3" json:"srlg_preferred_autobackups,omitempty"`
	SrlgWeightedAutobackups                       uint32   `protobuf:"varint,58,opt,name=srlg_weighted_autobackups,json=srlgWeightedAutobackups,proto3" json:"srlg_weighted_autobackups,omitempty"`
	NextHopAutobackupProtectedLsPs                uint32   `protobuf:"varint,59,opt,name=next_hop_autobackup_protected_ls_ps,json=nextHopAutobackupProtectedLsPs,proto3" json:"next_hop_autobackup_protected_ls_ps,omitempty"`
	NextNextHopAutobackupProtectedLsPs            uint32   `protobuf:"varint,60,opt,name=next_next_hop_autobackup_protected_ls_ps,json=nextNextHopAutobackupProtectedLsPs,proto3" json:"next_next_hop_autobackup_protected_ls_ps,omitempty"`
	NextHopSrlgAutobackupProtectedLsPs            uint32   `protobuf:"varint,61,opt,name=next_hop_srlg_autobackup_protected_ls_ps,json=nextHopSrlgAutobackupProtectedLsPs,proto3" json:"next_hop_srlg_autobackup_protected_ls_ps,omitempty"`
	NextNextHopSrlgAutobackupProtectedLsPs        uint32   `protobuf:"varint,62,opt,name=next_next_hop_srlg_autobackup_protected_ls_ps,json=nextNextHopSrlgAutobackupProtectedLsPs,proto3" json:"next_next_hop_srlg_autobackup_protected_ls_ps,omitempty"`
	NextHopAutobackupProtectedS2LFamilies         uint32   `protobuf:"varint,63,opt,name=next_hop_autobackup_protected_s2l_families,json=nextHopAutobackupProtectedS2lFamilies,proto3" json:"next_hop_autobackup_protected_s2l_families,omitempty"`
	NextNextHopAutobackupProtectedS2LFamilies     uint32   `protobuf:"varint,64,opt,name=next_next_hop_autobackup_protected_s2l_families,json=nextNextHopAutobackupProtectedS2lFamilies,proto3" json:"next_next_hop_autobackup_protected_s2l_families,omitempty"`
	NextHopSrlgAutobackupProtectedS2LFamilies     uint32   `protobuf:"varint,65,opt,name=next_hop_srlg_autobackup_protected_s2l_families,json=nextHopSrlgAutobackupProtectedS2lFamilies,proto3" json:"next_hop_srlg_autobackup_protected_s2l_families,omitempty"`
	NextNextHopSrlgAutobackupProtectedS2LFamilies uint32   `protobuf:"varint,66,opt,name=next_next_hop_srlg_autobackup_protected_s2l_families,json=nextNextHopSrlgAutobackupProtectedS2lFamilies,proto3" json:"next_next_hop_srlg_autobackup_protected_s2l_families,omitempty"`
	NextHopAutobackupProtectedS2Ls                uint32   `protobuf:"varint,67,opt,name=next_hop_autobackup_protected_s2_ls,json=nextHopAutobackupProtectedS2Ls,proto3" json:"next_hop_autobackup_protected_s2_ls,omitempty"`
	NextNextHopAutobackupProtectedS2Ls            uint32   `protobuf:"varint,68,opt,name=next_next_hop_autobackup_protected_s2_ls,json=nextNextHopAutobackupProtectedS2Ls,proto3" json:"next_next_hop_autobackup_protected_s2_ls,omitempty"`
	NextHopSrlgAutobackupProtectedS2Ls            uint32   `protobuf:"varint,69,opt,name=next_hop_srlg_autobackup_protected_s2_ls,json=nextHopSrlgAutobackupProtectedS2Ls,proto3" json:"next_hop_srlg_autobackup_protected_s2_ls,omitempty"`
	NextNextHopSrlgAutobackupProtectedS2Ls        uint32   `protobuf:"varint,70,opt,name=next_next_hop_srlg_autobackup_protected_s2_ls,json=nextNextHopSrlgAutobackupProtectedS2Ls,proto3" json:"next_next_hop_srlg_autobackup_protected_s2_ls,omitempty"`
	XXX_NoUnkeyedLiteral                          struct{} `json:"-"`
	XXX_unrecognized                              []byte   `json:"-"`
	XXX_sizecache                                 int32    `json:"-"`
}

func (m *MplsTeAutotunBackupSummaryInfo) Reset()         { *m = MplsTeAutotunBackupSummaryInfo{} }
func (m *MplsTeAutotunBackupSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAutotunBackupSummaryInfo) ProtoMessage()    {}
func (*MplsTeAutotunBackupSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d77850068d25b3, []int{1}
}

func (m *MplsTeAutotunBackupSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo.Unmarshal(m, b)
}
func (m *MplsTeAutotunBackupSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeAutotunBackupSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAutotunBackupSummaryInfo.Merge(m, src)
}
func (m *MplsTeAutotunBackupSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAutotunBackupSummaryInfo.Size(m)
}
func (m *MplsTeAutotunBackupSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAutotunBackupSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAutotunBackupSummaryInfo proto.InternalMessageInfo

func (m *MplsTeAutotunBackupSummaryInfo) GetAutobackups() uint32 {
	if m != nil {
		return m.Autobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetUpAutobackups() uint32 {
	if m != nil {
		return m.UpAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetDownAutobackups() uint32 {
	if m != nil {
		return m.DownAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetUnusedAutobackups() uint32 {
	if m != nil {
		return m.UnusedAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopAutobackups() uint32 {
	if m != nil {
		return m.NextHopAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopAutobackups() uint32 {
	if m != nil {
		return m.NextNextHopAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetSrlgStrictAutobackups() uint32 {
	if m != nil {
		return m.SrlgStrictAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetSrlgPreferredAutobackups() uint32 {
	if m != nil {
		return m.SrlgPreferredAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetSrlgWeightedAutobackups() uint32 {
	if m != nil {
		return m.SrlgWeightedAutobackups
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopAutobackupProtectedLsPs() uint32 {
	if m != nil {
		return m.NextHopAutobackupProtectedLsPs
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopAutobackupProtectedLsPs() uint32 {
	if m != nil {
		return m.NextNextHopAutobackupProtectedLsPs
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopSrlgAutobackupProtectedLsPs() uint32 {
	if m != nil {
		return m.NextHopSrlgAutobackupProtectedLsPs
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopSrlgAutobackupProtectedLsPs() uint32 {
	if m != nil {
		return m.NextNextHopSrlgAutobackupProtectedLsPs
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopAutobackupProtectedS2LFamilies() uint32 {
	if m != nil {
		return m.NextHopAutobackupProtectedS2LFamilies
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopAutobackupProtectedS2LFamilies() uint32 {
	if m != nil {
		return m.NextNextHopAutobackupProtectedS2LFamilies
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopSrlgAutobackupProtectedS2LFamilies() uint32 {
	if m != nil {
		return m.NextHopSrlgAutobackupProtectedS2LFamilies
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopSrlgAutobackupProtectedS2LFamilies() uint32 {
	if m != nil {
		return m.NextNextHopSrlgAutobackupProtectedS2LFamilies
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopAutobackupProtectedS2Ls() uint32 {
	if m != nil {
		return m.NextHopAutobackupProtectedS2Ls
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopAutobackupProtectedS2Ls() uint32 {
	if m != nil {
		return m.NextNextHopAutobackupProtectedS2Ls
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextHopSrlgAutobackupProtectedS2Ls() uint32 {
	if m != nil {
		return m.NextHopSrlgAutobackupProtectedS2Ls
	}
	return 0
}

func (m *MplsTeAutotunBackupSummaryInfo) GetNextNextHopSrlgAutobackupProtectedS2Ls() uint32 {
	if m != nil {
		return m.NextNextHopSrlgAutobackupProtectedS2Ls
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeAutotunBackupSummaryInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.backup.summary.mpls_te_autotun_backup_summary_info_KEYS")
	proto.RegisterType((*MplsTeAutotunBackupSummaryInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.auto_tunnel.backup.summary.mpls_te_autotun_backup_summary_info")
}

func init() {
	proto.RegisterFile("mpls_te_autotun_backup_summary_info.proto", fileDescriptor_c0d77850068d25b3)
}

var fileDescriptor_c0d77850068d25b3 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdb, 0x6b, 0x13, 0x41,
	0x14, 0xc6, 0xf1, 0xc5, 0x87, 0xa3, 0xf5, 0xb2, 0x78, 0xd9, 0xfa, 0x20, 0x25, 0xa5, 0x92, 0x14,
	0xb2, 0xca, 0xb6, 0x56, 0x5b, 0xe3, 0xa5, 0x6a, 0x4b, 0xa1, 0x45, 0x82, 0x2b, 0x48, 0x1f, 0x64,
	0x48, 0xb6, 0x93, 0x74, 0xe9, 0x64, 0x67, 0x98, 0x0b, 0xad, 0x7f, 0x91, 0xff, 0xa6, 0xcc, 0xd9,
	0x84, 0xce, 0x34, 0x7b, 0xeb, 0xe3, 0xce, 0xf9, 0xbe, 0xdf, 0x39, 0xf3, 0xcd, 0x81, 0x85, 0xde,
	0x4c, 0x30, 0x45, 0x34, 0x25, 0x23, 0xa3, 0xb9, 0x36, 0x39, 0x19, 0x8f, 0xd2, 0x0b, 0x23, 0x88,
	0x32, 0xb3, 0xd9, 0x48, 0xfe, 0x25, 0x59, 0x3e, 0xe1, 0x91, 0x90, 0x5c, 0xf3, 0x60, 0x90, 0x66,
	0x2a, 0xe5, 0x24, 0xe3, 0x8a, 0x5c, 0x49, 0xb2, 0xf0, 0x71, 0x41, 0x65, 0x34, 0xff, 0x88, 0x2c,
	0x84, 0x68, 0x93, 0xe7, 0x94, 0x45, 0x05, 0x28, 0x9a, 0x83, 0x3a, 0x9b, 0xd0, 0x6d, 0xd1, 0x8a,
	0x1c, 0x1f, 0x9c, 0x26, 0x9d, 0x7f, 0xf7, 0x61, 0xbd, 0x85, 0x38, 0x58, 0x83, 0x7b, 0xb6, 0x5c,
	0x94, 0x54, 0x18, 0xaf, 0xdd, 0xe9, 0xae, 0xfc, 0x74, 0x8f, 0x82, 0x0d, 0x78, 0x60, 0x04, 0x71,
	0x45, 0x5b, 0x28, 0x5a, 0x31, 0x62, 0xdf, 0x91, 0xf5, 0xe0, 0xd1, 0x19, 0xbf, 0xcc, 0x3d, 0xe1,
	0x36, 0x0a, 0x1f, 0xda, 0x73, 0x57, 0xda, 0x87, 0xc0, 0xe4, 0x46, 0xd1, 0x33, 0x4f, 0xfc, 0x16,
	0xc5, 0x8f, 0x8b, 0x8a, 0x2b, 0x7f, 0x03, 0x4f, 0x72, 0x7a, 0xa5, 0xc9, 0x39, 0xf7, 0xc7, 0xd8,
	0x41, 0x43, 0x60, 0x6b, 0x47, 0xdc, 0x9b, 0x65, 0x17, 0x56, 0xd1, 0x51, 0x6a, 0x7b, 0x87, 0xb6,
	0x67, 0xb6, 0xf6, 0x63, 0xd9, 0xba, 0x03, 0xcf, 0x95, 0x64, 0x53, 0xa2, 0xb4, 0xcc, 0x52, 0xed,
	0x19, 0xdf, 0xa3, 0xf1, 0xa9, 0x2d, 0x27, 0x58, 0x75, 0x7d, 0x03, 0x78, 0x81, 0x3e, 0x21, 0xe9,
	0x84, 0x4a, 0x79, 0xe3, 0x6e, 0xbb, 0x68, 0x0d, 0xad, 0x62, 0xb8, 0x10, 0xb8, 0xee, 0x3d, 0x58,
	0x45, 0xf7, 0x25, 0xcd, 0xa6, 0xe7, 0xfa, 0x86, 0x79, 0x0f, 0xcd, 0x38, 0xd6, 0xef, 0x79, 0xdd,
	0xf5, 0x1e, 0xc3, 0x7a, 0xc9, 0x3d, 0x89, 0x5d, 0x38, 0x9a, 0x5a, 0x16, 0x53, 0x44, 0xa8, 0xf0,
	0x03, 0x52, 0x5e, 0x2e, 0xa5, 0x35, 0x5c, 0xe8, 0x4e, 0xd4, 0x50, 0x05, 0xbf, 0xa0, 0x5b, 0x95,
	0xdc, 0x12, 0x71, 0x80, 0xc4, 0x4e, 0x69, 0x90, 0xe5, 0x54, 0x0b, 0xc4, 0x7b, 0xd6, 0x50, 0x3f,
	0x5e, 0x53, 0x8f, 0xb8, 0x48, 0x24, 0x9b, 0x56, 0x51, 0xff, 0x40, 0xdf, 0x9f, 0xb5, 0x09, 0xfd,
	0x09, 0xd1, 0xaf, 0x9c, 0x81, 0xeb, 0xf0, 0xa7, 0xb0, 0x59, 0x9f, 0x82, 0x8a, 0x19, 0x99, 0x8c,
	0x66, 0x19, 0xcb, 0xa8, 0x0a, 0x3f, 0x23, 0x7b, 0xa3, 0x3a, 0xde, 0x24, 0x66, 0x87, 0x73, 0x71,
	0x30, 0x86, 0xd7, 0x2d, 0x52, 0xf6, 0xf8, 0x5f, 0x90, 0xdf, 0xab, 0x0f, 0xbb, 0xac, 0x47, 0x7d,
	0x30, 0x5e, 0x8f, 0xfd, 0xeb, 0x1e, 0xd5, 0xd9, 0xb8, 0x3d, 0x2e, 0x60, 0xbb, 0xed, 0x0b, 0x78,
	0x8d, 0xbe, 0x62, 0xa3, 0x7e, 0xf3, 0x43, 0xb8, 0xcd, 0x1a, 0xf7, 0x5c, 0xc5, 0x84, 0xa9, 0xf0,
	0x5b, 0xd3, 0x9e, 0x27, 0xf1, 0x49, 0xdb, 0x3d, 0x2f, 0x88, 0xdf, 0xdb, 0xec, 0xb9, 0x47, 0x6d,
	0x8a, 0xc2, 0x52, 0x0f, 0xda, 0xec, 0x39, 0x52, 0x6f, 0xb1, 0xe7, 0x05, 0xfa, 0xb0, 0xed, 0x9e,
	0x5b, 0xfc, 0xf8, 0x2e, 0xfe, 0x9a, 0xb6, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xa0, 0xc3,
	0xaf, 0xc7, 0x06, 0x00, 0x00,
}
