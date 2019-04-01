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

package cisco_ios_xr_ipv4_pim_oper_pim_standby_default_context_bsr_candidate_rps_candidate_rp

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
	RpAddress            string   `protobuf:"bytes,1,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
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
	proto.RegisterType((*PimBsrCandidateRpBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.candidate_rps.candidate_rp.pim_bsr_candidate_rp_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.candidate_rps.candidate_rp.pim_addrtype")
	proto.RegisterType((*PimBsrCrpAclBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.candidate_rps.candidate_rp.pim_bsr_crp_acl_bag")
	proto.RegisterType((*PimBsrCandidateRpBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.standby.default_context.bsr.candidate_rps.candidate_rp.pim_bsr_candidate_rp_bag")
}

func init() { proto.RegisterFile("pim_bsr_candidate_rp_bag.proto", fileDescriptor_56781592f929d3c3) }

var fileDescriptor_56781592f929d3c3 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x5d, 0x8b, 0xd4, 0x30,
	0x14, 0xa5, 0xae, 0xec, 0x3a, 0xb7, 0x23, 0x3a, 0xd9, 0x07, 0xab, 0x30, 0x52, 0xe7, 0xa9, 0x88,
	0xe4, 0x61, 0x76, 0x1d, 0x3f, 0x1e, 0x84, 0x3e, 0x08, 0x8a, 0x28, 0xd2, 0x65, 0x1f, 0x04, 0x21,
	0xa4, 0x49, 0x46, 0x23, 0x6d, 0x13, 0x92, 0xec, 0xb0, 0xf3, 0x07, 0x04, 0x7f, 0xa4, 0xff, 0x45,
	0x92, 0x7e, 0xd0, 0xca, 0xee, 0xe3, 0x3c, 0xf6, 0xdc, 0x73, 0xee, 0x39, 0xb9, 0xf7, 0x16, 0x9e,
	0x6a, 0x59, 0x93, 0xd2, 0x1a, 0xc2, 0x68, 0xc3, 0x25, 0xa7, 0x4e, 0x10, 0xa3, 0x49, 0x49, 0x7f,
	0x60, 0x6d, 0x94, 0x53, 0xe8, 0x92, 0x49, 0xcb, 0x14, 0x91, 0xca, 0x92, 0x6b, 0x43, 0xa4, 0xde,
	0x9d, 0x13, 0xaf, 0x50, 0x5a, 0x18, 0xac, 0x65, 0x8d, 0xad, 0xa3, 0x0d, 0x2f, 0xf7, 0x98, 0x8b,
	0x2d, 0xbd, 0xaa, 0x1c, 0x61, 0xaa, 0x71, 0xe2, 0xda, 0xe1, 0xd2, 0x1a, 0x3c, 0x6e, 0x69, 0x27,
	0x5f, 0xab, 0x77, 0xb0, 0xbc, 0xcd, 0x98, 0x7c, 0x7a, 0xff, 0xed, 0x02, 0x2d, 0x01, 0x8c, 0x26,
	0x94, 0x73, 0x23, 0xac, 0x4d, 0xa2, 0x34, 0xca, 0x66, 0xc5, 0xcc, 0xe8, 0xbc, 0x05, 0x56, 0x35,
	0xcc, 0xbd, 0xde, 0xd7, 0xdd, 0x5e, 0x0b, 0xf4, 0x08, 0x4e, 0xe8, 0x96, 0x34, 0xb4, 0x16, 0x1d,
	0xf7, 0x98, 0x6e, 0xbf, 0xd0, 0x5a, 0xa0, 0x67, 0x30, 0x0f, 0xa1, 0xfb, 0x4e, 0x77, 0x42, 0x35,
	0xf6, 0x58, 0xd7, 0xab, 0xa3, 0x6c, 0x06, 0xca, 0xd1, 0x40, 0xd9, 0xf4, 0x76, 0xdf, 0xe1, 0x74,
	0x88, 0xeb, 0x63, 0xb1, 0xca, 0x27, 0x45, 0xcf, 0x61, 0x31, 0x49, 0x5f, 0x2b, 0xde, 0xfb, 0x3f,
	0x18, 0x0a, 0x85, 0xfe, 0xac, 0xb8, 0x40, 0x8f, 0xe1, 0x9e, 0x97, 0x85, 0x88, 0x6d, 0x88, 0x13,
	0xca, 0x2a, 0x9f, 0x71, 0xf5, 0xf7, 0x2e, 0x24, 0xb7, 0x4d, 0x03, 0xfd, 0x8e, 0x60, 0x3e, 0x06,
	0x93, 0x75, 0x1a, 0x65, 0xf1, 0x9a, 0xe1, 0x83, 0x2c, 0x06, 0x8f, 0xa7, 0x5a, 0xc4, 0xa3, 0x47,
	0xdc, 0xfc, 0xd8, 0xb3, 0x9b, 0x1f, 0xfb, 0x02, 0xd0, 0x84, 0x6b, 0x99, 0xd2, 0x22, 0x39, 0x4f,
	0xa3, 0x6c, 0x51, 0x3c, 0x1c, 0x91, 0x2f, 0x3c, 0xee, 0x17, 0xe0, 0xa7, 0xaa, 0x8d, 0x54, 0x46,
	0xba, 0x7d, 0xf2, 0x32, 0x8d, 0xb2, 0xfb, 0x45, 0xcc, 0x8c, 0xfe, 0xda, 0x41, 0x3d, 0xe5, 0xa7,
	0xaa, 0xb8, 0x93, 0xb5, 0x48, 0x36, 0x03, 0xe5, 0x43, 0x07, 0xa1, 0x1c, 0x96, 0x13, 0x4f, 0xca,
	0x77, 0xb4, 0x61, 0x82, 0xc8, 0xc6, 0x09, 0xb3, 0xa3, 0x55, 0xf2, 0x2a, 0x68, 0x9e, 0x8c, 0xec,
	0xf3, 0x96, 0xf2, 0xb1, 0x63, 0x20, 0x0c, 0xa7, 0x93, 0x16, 0x57, 0x3a, 0x98, 0xbd, 0x0e, 0xc2,
	0xc5, 0x48, 0x78, 0x19, 0x0a, 0x93, 0x9d, 0xbe, 0x99, 0xec, 0x14, 0xfd, 0x89, 0x00, 0xda, 0x53,
	0x61, 0xfe, 0xa6, 0xde, 0xa6, 0x47, 0x59, 0xbc, 0xfe, 0x75, 0xc0, 0xa5, 0xfd, 0x77, 0x9b, 0xc5,
	0x8c, 0x19, 0x9d, 0x07, 0xf3, 0xf2, 0x38, 0xfc, 0xca, 0x67, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x79, 0xa8, 0xb4, 0xd6, 0xec, 0x03, 0x00, 0x00,
}
