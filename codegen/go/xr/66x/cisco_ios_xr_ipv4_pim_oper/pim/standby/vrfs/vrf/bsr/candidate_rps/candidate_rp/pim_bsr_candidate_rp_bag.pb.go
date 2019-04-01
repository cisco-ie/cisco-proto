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
// source: pim_bsr_candidate_rp_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_standby_vrfs_vrf_bsr_candidate_rps_candidate_rp

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

type PimBsrCandidateRpBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RpAddress            string   `protobuf:"bytes,2,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBsrCandidateRpBag_KEYS) Reset()         { *m = PimBsrCandidateRpBag_KEYS{} }
func (m *PimBsrCandidateRpBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBsrCandidateRpBag_KEYS) ProtoMessage()    {}
func (*PimBsrCandidateRpBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_56781592f929d3c3, []int{0}
}

func (m *PimBsrCandidateRpBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrCandidateRpBag_KEYS.Unmarshal(m, b)
}
func (m *PimBsrCandidateRpBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrCandidateRpBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBsrCandidateRpBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrCandidateRpBag_KEYS.Merge(m, src)
}
func (m *PimBsrCandidateRpBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBsrCandidateRpBag_KEYS.Size(m)
}
func (m *PimBsrCandidateRpBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrCandidateRpBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrCandidateRpBag_KEYS proto.InternalMessageInfo

func (m *PimBsrCandidateRpBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimBsrCandidateRpBag_KEYS) GetRpAddress() string {
	if m != nil {
		return m.RpAddress
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_56781592f929d3c3, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimBsrCrpAclBag struct {
	CandidateRpMode      string   `protobuf:"bytes,1,opt,name=candidate_rp_mode,json=candidateRpMode,proto3" json:"candidate_rp_mode,omitempty"`
	AclName              string   `protobuf:"bytes,2,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBsrCrpAclBag) Reset()         { *m = PimBsrCrpAclBag{} }
func (m *PimBsrCrpAclBag) String() string { return proto.CompactTextString(m) }
func (*PimBsrCrpAclBag) ProtoMessage()    {}
func (*PimBsrCrpAclBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_56781592f929d3c3, []int{2}
}

func (m *PimBsrCrpAclBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrCrpAclBag.Unmarshal(m, b)
}
func (m *PimBsrCrpAclBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrCrpAclBag.Marshal(b, m, deterministic)
}
func (m *PimBsrCrpAclBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrCrpAclBag.Merge(m, src)
}
func (m *PimBsrCrpAclBag) XXX_Size() int {
	return xxx_messageInfo_PimBsrCrpAclBag.Size(m)
}
func (m *PimBsrCrpAclBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrCrpAclBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrCrpAclBag proto.InternalMessageInfo

func (m *PimBsrCrpAclBag) GetCandidateRpMode() string {
	if m != nil {
		return m.CandidateRpMode
	}
	return ""
}

func (m *PimBsrCrpAclBag) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

type PimBsrCandidateRpBag struct {
	CandidateRp                *PimAddrtype       `protobuf:"bytes,50,opt,name=candidate_rp,json=candidateRp,proto3" json:"candidate_rp,omitempty"`
	CandidateRpMode            string             `protobuf:"bytes,51,opt,name=candidate_rp_mode,json=candidateRpMode,proto3" json:"candidate_rp_mode,omitempty"`
	CandidateRpScope           int32              `protobuf:"zigzag32,52,opt,name=candidate_rp_scope,json=candidateRpScope,proto3" json:"candidate_rp_scope,omitempty"`
	CrpPriority                uint32             `protobuf:"varint,53,opt,name=crp_priority,json=crpPriority,proto3" json:"crp_priority,omitempty"`
	CrpHoldtime                uint32             `protobuf:"varint,54,opt,name=crp_holdtime,json=crpHoldtime,proto3" json:"crp_holdtime,omitempty"`
	CandidateRpAdvanceInterval uint32             `protobuf:"varint,55,opt,name=candidate_rp_advance_interval,json=candidateRpAdvanceInterval,proto3" json:"candidate_rp_advance_interval,omitempty"`
	CandidateRpUptime          uint32             `protobuf:"varint,56,opt,name=candidate_rp_uptime,json=candidateRpUptime,proto3" json:"candidate_rp_uptime,omitempty"`
	AclName                    string             `protobuf:"bytes,57,opt,name=acl_name,json=aclName,proto3" json:"acl_name,omitempty"`
	CrpAccess                  []*PimBsrCrpAclBag `protobuf:"bytes,58,rep,name=crp_access,json=crpAccess,proto3" json:"crp_access,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}           `json:"-"`
	XXX_unrecognized           []byte             `json:"-"`
	XXX_sizecache              int32              `json:"-"`
}

func (m *PimBsrCandidateRpBag) Reset()         { *m = PimBsrCandidateRpBag{} }
func (m *PimBsrCandidateRpBag) String() string { return proto.CompactTextString(m) }
func (*PimBsrCandidateRpBag) ProtoMessage()    {}
func (*PimBsrCandidateRpBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_56781592f929d3c3, []int{3}
}

func (m *PimBsrCandidateRpBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrCandidateRpBag.Unmarshal(m, b)
}
func (m *PimBsrCandidateRpBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrCandidateRpBag.Marshal(b, m, deterministic)
}
func (m *PimBsrCandidateRpBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrCandidateRpBag.Merge(m, src)
}
func (m *PimBsrCandidateRpBag) XXX_Size() int {
	return xxx_messageInfo_PimBsrCandidateRpBag.Size(m)
}
func (m *PimBsrCandidateRpBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrCandidateRpBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrCandidateRpBag proto.InternalMessageInfo

func (m *PimBsrCandidateRpBag) GetCandidateRp() *PimAddrtype {
	if m != nil {
		return m.CandidateRp
	}
	return nil
}

func (m *PimBsrCandidateRpBag) GetCandidateRpMode() string {
	if m != nil {
		return m.CandidateRpMode
	}
	return ""
}

func (m *PimBsrCandidateRpBag) GetCandidateRpScope() int32 {
	if m != nil {
		return m.CandidateRpScope
	}
	return 0
}

func (m *PimBsrCandidateRpBag) GetCrpPriority() uint32 {
	if m != nil {
		return m.CrpPriority
	}
	return 0
}

func (m *PimBsrCandidateRpBag) GetCrpHoldtime() uint32 {
	if m != nil {
		return m.CrpHoldtime
	}
	return 0
}

func (m *PimBsrCandidateRpBag) GetCandidateRpAdvanceInterval() uint32 {
	if m != nil {
		return m.CandidateRpAdvanceInterval
	}
	return 0
}

func (m *PimBsrCandidateRpBag) GetCandidateRpUptime() uint32 {
	if m != nil {
		return m.CandidateRpUptime
	}
	return 0
}

func (m *PimBsrCandidateRpBag) GetAclName() string {
	if m != nil {
		return m.AclName
	}
	return ""
}

func (m *PimBsrCandidateRpBag) GetCrpAccess() []*PimBsrCrpAclBag {
	if m != nil {
		return m.CrpAccess
	}
	return nil
}

func init() {
	proto.RegisterType((*PimBsrCandidateRpBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bsr.candidate_rps.candidate_rp.pim_bsr_candidate_rp_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bsr.candidate_rps.candidate_rp.pim_addrtype")
	proto.RegisterType((*PimBsrCrpAclBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bsr.candidate_rps.candidate_rp.pim_bsr_crp_acl_bag")
	proto.RegisterType((*PimBsrCandidateRpBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.vrfs.vrf.bsr.candidate_rps.candidate_rp.pim_bsr_candidate_rp_bag")
}

func init() { proto.RegisterFile("pim_bsr_candidate_rp_bag.proto", fileDescriptor_56781592f929d3c3) }

var fileDescriptor_56781592f929d3c3 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0x26, 0xae, 0x6c, 0xed, 0x4b, 0x45, 0x3b, 0x7b, 0x30, 0x0a, 0x95, 0xd8, 0x53, 0x10, 0xc9,
	0xa1, 0xbb, 0xd6, 0x1f, 0xb7, 0x1e, 0x04, 0x45, 0x5c, 0x24, 0x8b, 0x87, 0x85, 0x85, 0x61, 0x32,
	0x33, 0xd5, 0x81, 0x26, 0x33, 0xbc, 0x89, 0xc1, 0x9e, 0x04, 0xff, 0x40, 0xff, 0x26, 0x99, 0x49,
	0x1a, 0x32, 0xb2, 0xbd, 0xe9, 0xa5, 0xd0, 0xef, 0x7d, 0xef, 0xfb, 0xde, 0xf7, 0xde, 0x04, 0x9e,
	0x1a, 0x55, 0xd1, 0xd2, 0x22, 0xe5, 0xac, 0x16, 0x4a, 0xb0, 0x46, 0x52, 0x34, 0xb4, 0x64, 0x5f,
	0x73, 0x83, 0xba, 0xd1, 0xe4, 0x92, 0x2b, 0xcb, 0x35, 0x55, 0xda, 0xd2, 0x1f, 0x48, 0x95, 0x69,
	0x2f, 0xa8, 0xeb, 0xd0, 0x46, 0x62, 0x6e, 0x54, 0x95, 0xdb, 0x86, 0xd5, 0xa2, 0xdc, 0xe7, 0x2d,
	0x6e, 0xad, 0xfb, 0xc9, 0x4b, 0x8b, 0xf9, 0x58, 0xcb, 0x06, 0xff, 0x96, 0xd7, 0xb0, 0x38, 0xe6,
	0x48, 0x3f, 0xbe, 0xbb, 0xbe, 0x22, 0x8f, 0xe1, 0x5e, 0x8b, 0x5b, 0x5a, 0xb3, 0x4a, 0x26, 0x51,
	0x1a, 0x65, 0xd3, 0x62, 0xd2, 0xe2, 0xf6, 0x92, 0x55, 0x92, 0x2c, 0x00, 0xd0, 0x50, 0x26, 0x04,
	0x4a, 0x6b, 0x93, 0x3b, 0xbe, 0x38, 0x45, 0xb3, 0xe9, 0x80, 0x65, 0x05, 0x33, 0x27, 0xed, 0xea,
	0xcd, 0xde, 0x48, 0xf2, 0x08, 0x26, 0x2c, 0x10, 0x3a, 0x65, 0x9d, 0xce, 0x33, 0x98, 0xf9, 0x20,
	0xa1, 0x52, 0xec, 0xb0, 0x5e, 0xab, 0xa7, 0xac, 0x07, 0xca, 0xc9, 0x40, 0x59, 0x1f, 0xec, 0x6e,
	0xe0, 0x6c, 0x48, 0xe2, 0xc6, 0xe2, 0x3b, 0x17, 0x82, 0x3c, 0x87, 0x79, 0x10, 0xac, 0xd2, 0xe2,
	0xe0, 0xff, 0x60, 0x28, 0x14, 0xe6, 0x93, 0x16, 0xd2, 0x65, 0x75, 0x6d, 0x7e, 0xc4, 0x6e, 0x88,
	0x09, 0xe3, 0x3b, 0x37, 0xe3, 0xf2, 0xf7, 0x5d, 0x48, 0x8e, 0x2d, 0x8a, 0xfc, 0x84, 0xd9, 0x18,
	0x4b, 0x56, 0x69, 0x94, 0xc5, 0xab, 0x9b, 0xfc, 0xdf, 0xde, 0x2a, 0x1f, 0x6f, 0xb3, 0x88, 0x47,
	0xc3, 0xdf, 0x1e, 0xf2, 0xfc, 0xf6, 0x90, 0x2f, 0x80, 0x04, 0x5c, 0xcb, 0xb5, 0x91, 0xc9, 0x45,
	0x1a, 0x65, 0xf3, 0xe2, 0xe1, 0x88, 0x7c, 0xe5, 0x70, 0xb7, 0x78, 0xb7, 0x4d, 0x83, 0x4a, 0xa3,
	0x6a, 0xf6, 0xc9, 0xcb, 0x34, 0xca, 0xee, 0x17, 0x31, 0x47, 0xf3, 0xb9, 0x87, 0x0e, 0x94, 0x6f,
	0x7a, 0x27, 0x1a, 0x55, 0xc9, 0x64, 0x3d, 0x50, 0xde, 0xf7, 0x10, 0xd9, 0xc0, 0x22, 0xf0, 0x64,
	0xa2, 0x65, 0x35, 0x97, 0x54, 0xd5, 0x8d, 0xc4, 0x96, 0xed, 0x92, 0x57, 0xbe, 0xe7, 0xc9, 0xc8,
	0x7e, 0xd3, 0x51, 0x3e, 0xf4, 0x0c, 0x92, 0xc3, 0x59, 0x20, 0xf1, 0xdd, 0x78, 0xb3, 0xd7, 0xbe,
	0x71, 0x3e, 0x6a, 0xfc, 0xe2, 0x0b, 0xc1, 0x2d, 0xdf, 0x04, 0xb7, 0x24, 0xbf, 0x22, 0x80, 0xee,
	0x89, 0x70, 0xf7, 0x96, 0xde, 0xa6, 0x27, 0x59, 0xbc, 0xe2, 0xff, 0xe3, 0x5a, 0x7f, 0x3d, 0xc6,
	0x62, 0xca, 0xd1, 0x6c, 0xbc, 0x6b, 0x79, 0xea, 0xbf, 0xe7, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5a, 0xb9, 0xdb, 0xde, 0xf1, 0x03, 0x00, 0x00,
}