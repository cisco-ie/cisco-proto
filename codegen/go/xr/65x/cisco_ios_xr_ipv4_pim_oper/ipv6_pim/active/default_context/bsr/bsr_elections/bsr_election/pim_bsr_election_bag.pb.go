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
// source: pim_bsr_election_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_bsr_bsr_elections_bsr_election

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

type PimBsrElectionBag_KEYS struct {
	PimScope             uint32   `protobuf:"varint,1,opt,name=pim_scope,json=pimScope,proto3" json:"pim_scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimBsrElectionBag_KEYS) Reset()         { *m = PimBsrElectionBag_KEYS{} }
func (m *PimBsrElectionBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimBsrElectionBag_KEYS) ProtoMessage()    {}
func (*PimBsrElectionBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fd166ab4631d04a, []int{0}
}

func (m *PimBsrElectionBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrElectionBag_KEYS.Unmarshal(m, b)
}
func (m *PimBsrElectionBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrElectionBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimBsrElectionBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrElectionBag_KEYS.Merge(m, src)
}
func (m *PimBsrElectionBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimBsrElectionBag_KEYS.Size(m)
}
func (m *PimBsrElectionBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrElectionBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrElectionBag_KEYS proto.InternalMessageInfo

func (m *PimBsrElectionBag_KEYS) GetPimScope() uint32 {
	if m != nil {
		return m.PimScope
	}
	return 0
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
	return fileDescriptor_3fd166ab4631d04a, []int{1}
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

type PimBsrElectionBag struct {
	BsrAddress             *PimAddrtype `protobuf:"bytes,50,opt,name=bsr_address,json=bsrAddress,proto3" json:"bsr_address,omitempty"`
	BsrPriority            uint32       `protobuf:"varint,51,opt,name=bsr_priority,json=bsrPriority,proto3" json:"bsr_priority,omitempty"`
	BsrMaskLength          uint32       `protobuf:"varint,52,opt,name=bsr_mask_length,json=bsrMaskLength,proto3" json:"bsr_mask_length,omitempty"`
	BsrUpTime              uint32       `protobuf:"varint,53,opt,name=bsr_up_time,json=bsrUpTime,proto3" json:"bsr_up_time,omitempty"`
	BootstrapTimeout       uint32       `protobuf:"varint,54,opt,name=bootstrap_timeout,json=bootstrapTimeout,proto3" json:"bootstrap_timeout,omitempty"`
	CandidateBsrState      uint32       `protobuf:"varint,55,opt,name=candidate_bsr_state,json=candidateBsrState,proto3" json:"candidate_bsr_state,omitempty"`
	BsrElectionState       uint32       `protobuf:"varint,56,opt,name=bsr_election_state,json=bsrElectionState,proto3" json:"bsr_election_state,omitempty"`
	BsrScope               uint32       `protobuf:"varint,57,opt,name=bsr_scope,json=bsrScope,proto3" json:"bsr_scope,omitempty"`
	CandidateBsrFlag       bool         `protobuf:"varint,58,opt,name=candidate_bsr_flag,json=candidateBsrFlag,proto3" json:"candidate_bsr_flag,omitempty"`
	CandidateBsrPriority   uint32       `protobuf:"varint,59,opt,name=candidate_bsr_priority,json=candidateBsrPriority,proto3" json:"candidate_bsr_priority,omitempty"`
	CandidateBsrAddress    *PimAddrtype `protobuf:"bytes,60,opt,name=candidate_bsr_address,json=candidateBsrAddress,proto3" json:"candidate_bsr_address,omitempty"`
	CandidateBsrMaskLength uint32       `protobuf:"varint,61,opt,name=candidate_bsr_mask_length,json=candidateBsrMaskLength,proto3" json:"candidate_bsr_mask_length,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}     `json:"-"`
	XXX_unrecognized       []byte       `json:"-"`
	XXX_sizecache          int32        `json:"-"`
}

func (m *PimBsrElectionBag) Reset()         { *m = PimBsrElectionBag{} }
func (m *PimBsrElectionBag) String() string { return proto.CompactTextString(m) }
func (*PimBsrElectionBag) ProtoMessage()    {}
func (*PimBsrElectionBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fd166ab4631d04a, []int{2}
}

func (m *PimBsrElectionBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimBsrElectionBag.Unmarshal(m, b)
}
func (m *PimBsrElectionBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimBsrElectionBag.Marshal(b, m, deterministic)
}
func (m *PimBsrElectionBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimBsrElectionBag.Merge(m, src)
}
func (m *PimBsrElectionBag) XXX_Size() int {
	return xxx_messageInfo_PimBsrElectionBag.Size(m)
}
func (m *PimBsrElectionBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimBsrElectionBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimBsrElectionBag proto.InternalMessageInfo

func (m *PimBsrElectionBag) GetBsrAddress() *PimAddrtype {
	if m != nil {
		return m.BsrAddress
	}
	return nil
}

func (m *PimBsrElectionBag) GetBsrPriority() uint32 {
	if m != nil {
		return m.BsrPriority
	}
	return 0
}

func (m *PimBsrElectionBag) GetBsrMaskLength() uint32 {
	if m != nil {
		return m.BsrMaskLength
	}
	return 0
}

func (m *PimBsrElectionBag) GetBsrUpTime() uint32 {
	if m != nil {
		return m.BsrUpTime
	}
	return 0
}

func (m *PimBsrElectionBag) GetBootstrapTimeout() uint32 {
	if m != nil {
		return m.BootstrapTimeout
	}
	return 0
}

func (m *PimBsrElectionBag) GetCandidateBsrState() uint32 {
	if m != nil {
		return m.CandidateBsrState
	}
	return 0
}

func (m *PimBsrElectionBag) GetBsrElectionState() uint32 {
	if m != nil {
		return m.BsrElectionState
	}
	return 0
}

func (m *PimBsrElectionBag) GetBsrScope() uint32 {
	if m != nil {
		return m.BsrScope
	}
	return 0
}

func (m *PimBsrElectionBag) GetCandidateBsrFlag() bool {
	if m != nil {
		return m.CandidateBsrFlag
	}
	return false
}

func (m *PimBsrElectionBag) GetCandidateBsrPriority() uint32 {
	if m != nil {
		return m.CandidateBsrPriority
	}
	return 0
}

func (m *PimBsrElectionBag) GetCandidateBsrAddress() *PimAddrtype {
	if m != nil {
		return m.CandidateBsrAddress
	}
	return nil
}

func (m *PimBsrElectionBag) GetCandidateBsrMaskLength() uint32 {
	if m != nil {
		return m.CandidateBsrMaskLength
	}
	return 0
}

func init() {
	proto.RegisterType((*PimBsrElectionBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.bsr.bsr_elections.bsr_election.pim_bsr_election_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.bsr.bsr_elections.bsr_election.pim_addrtype")
	proto.RegisterType((*PimBsrElectionBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.bsr.bsr_elections.bsr_election.pim_bsr_election_bag")
}

func init() { proto.RegisterFile("pim_bsr_election_bag.proto", fileDescriptor_3fd166ab4631d04a) }

var fileDescriptor_3fd166ab4631d04a = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0x15, 0x10, 0xa5, 0x73, 0xa6, 0x15, 0xad, 0x5b, 0x4a, 0x4a, 0x25, 0x34, 0xcc, 0x02,
	0x8d, 0x04, 0xf2, 0xa2, 0x1d, 0x86, 0x96, 0xcb, 0x02, 0xa4, 0xb2, 0xe1, 0x22, 0x34, 0x53, 0x16,
	0x5d, 0x59, 0x76, 0xe2, 0x04, 0xab, 0x49, 0x6c, 0xf9, 0x78, 0x46, 0xed, 0x1b, 0xf0, 0x0e, 0x3c,
	0x25, 0x6f, 0x80, 0xec, 0x64, 0x42, 0x22, 0x75, 0x0b, 0x4b, 0x9f, 0xff, 0x3b, 0xd7, 0xfc, 0x81,
	0xc7, 0x46, 0x95, 0x4c, 0xa0, 0x65, 0xb2, 0x90, 0x89, 0x53, 0xba, 0x62, 0x82, 0xe7, 0xd4, 0x58,
	0xed, 0x34, 0xb9, 0x4c, 0x14, 0x26, 0x9a, 0x29, 0x8d, 0xec, 0xda, 0x32, 0x65, 0x56, 0x53, 0xe6,
	0x69, 0x6d, 0xa4, 0xa5, 0xca, 0xac, 0x66, 0xfe, 0x45, 0x79, 0xe2, 0xd4, 0x4a, 0xd2, 0x54, 0x66,
	0x7c, 0x59, 0x38, 0x96, 0xe8, 0xca, 0xc9, 0x6b, 0x47, 0x05, 0x5a, 0xda, 0x2d, 0x8b, 0xbd, 0xd7,
	0xf8, 0x14, 0x0e, 0x6f, 0x6b, 0xcc, 0x3e, 0x9d, 0x5f, 0x2e, 0xc8, 0x11, 0x0c, 0xbc, 0x88, 0x89,
	0x36, 0x32, 0x8e, 0x46, 0xd1, 0x64, 0x7b, 0xbe, 0x69, 0x54, 0xb9, 0xf0, 0xef, 0x71, 0x09, 0x5b,
	0x5e, 0xe4, 0x69, 0x6a, 0xdd, 0x8d, 0x91, 0xe4, 0x11, 0xdc, 0xe7, 0x19, 0xab, 0x78, 0x59, 0xa3,
	0x83, 0xf9, 0x06, 0xcf, 0xbe, 0xf2, 0x52, 0x92, 0xa7, 0xb0, 0x15, 0x46, 0xf6, 0xa4, 0x44, 0x8c,
	0xef, 0x04, 0x75, 0xe8, 0x63, 0xef, 0xeb, 0x50, 0x83, 0xcc, 0x5a, 0xe4, 0x6e, 0x8b, 0xcc, 0x1a,
	0x64, 0xfc, 0xfb, 0x1e, 0xec, 0xdf, 0x36, 0x29, 0xf9, 0x19, 0xc1, 0xd0, 0x07, 0xd7, 0xb9, 0xc7,
	0xa3, 0x68, 0x32, 0x3c, 0xce, 0xe9, 0x3f, 0xbb, 0x19, 0xed, 0xae, 0x3d, 0x07, 0x81, 0xb6, 0xb3,
	0x86, 0x07, 0x8d, 0x55, 0xda, 0x2a, 0x77, 0x13, 0x9f, 0x84, 0x93, 0xf9, 0xe9, 0xbe, 0x35, 0x21,
	0xf2, 0x0c, 0x1e, 0x78, 0xa4, 0xe4, 0x78, 0xc5, 0x0a, 0x59, 0xe5, 0xee, 0x47, 0x3c, 0x0d, 0xd4,
	0xb6, 0x40, 0xfb, 0x85, 0xe3, 0xd5, 0xe7, 0x10, 0x24, 0x4f, 0xea, 0xa5, 0x96, 0x86, 0x39, 0x55,
	0xca, 0xf8, 0x65, 0x60, 0x06, 0x02, 0xed, 0x77, 0x73, 0xa1, 0x4a, 0x49, 0x9e, 0xc3, 0xae, 0xd0,
	0xda, 0xa1, 0xb3, 0xbc, 0x46, 0xf4, 0xd2, 0xc5, 0xb3, 0x40, 0xed, 0xb4, 0xc2, 0x45, 0x1d, 0x27,
	0x14, 0xf6, 0x12, 0x5e, 0xa5, 0x2a, 0xe5, 0x4e, 0x86, 0x03, 0xa2, 0xe3, 0x4e, 0xc6, 0xaf, 0x02,
	0xbe, 0xdb, 0x4a, 0x1f, 0xd0, 0x2e, 0xbc, 0x40, 0x5e, 0x00, 0xe9, 0x9d, 0xb9, 0xc6, 0x4f, 0x9b,
	0xea, 0x68, 0xcf, 0x1b, 0xa1, 0xa6, 0x8f, 0x60, 0x10, 0x6a, 0x06, 0x97, 0x9c, 0xd5, 0x2e, 0x11,
	0x68, 0x83, 0x4b, 0x7c, 0xa9, 0x7e, 0xeb, 0xac, 0xe0, 0x79, 0xfc, 0x7a, 0x14, 0x4d, 0x36, 0xe7,
	0x3b, 0xdd, 0xce, 0x1f, 0x0b, 0x9e, 0x93, 0x29, 0x1c, 0xf4, 0xe9, 0xf6, 0x94, 0x6f, 0x42, 0xdd,
	0xfd, 0x6e, 0x46, 0x7b, 0xd3, 0x5f, 0x11, 0x3c, 0xec, 0xa7, 0xad, 0xbd, 0xf0, 0xf6, 0xff, 0x7a,
	0x61, 0xaf, 0x3b, 0xde, 0xda, 0x14, 0x67, 0x70, 0xd8, 0x1f, 0xae, 0xfb, 0xed, 0xdf, 0x85, 0xb5,
	0x0e, 0xba, 0x79, 0x7f, 0x4d, 0x20, 0x36, 0xc2, 0xef, 0x7f, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0xf0, 0x39, 0x2b, 0x1c, 0x04, 0x00, 0x00,
}
