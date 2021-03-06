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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_bsr_bsr_elections_bsr_election

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
	proto.RegisterType((*PimBsrElectionBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bsr.bsr_elections.bsr_election.pim_bsr_election_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bsr.bsr_elections.bsr_election.pim_addrtype")
	proto.RegisterType((*PimBsrElectionBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.bsr.bsr_elections.bsr_election.pim_bsr_election_bag")
}

func init() { proto.RegisterFile("pim_bsr_election_bag.proto", fileDescriptor_3fd166ab4631d04a) }

var fileDescriptor_3fd166ab4631d04a = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x15, 0x10, 0x63, 0x7d, 0xba, 0x89, 0xcd, 0x1b, 0x23, 0x63, 0x12, 0x2a, 0x3d, 0xa0,
	0x4a, 0x20, 0x1f, 0xb6, 0x52, 0x36, 0x5e, 0x0e, 0x20, 0x8d, 0x0b, 0x2f, 0x42, 0x6d, 0x39, 0x70,
	0xb2, 0xec, 0xc4, 0x2d, 0xd6, 0xe2, 0xd8, 0xf2, 0xe3, 0x4e, 0xdb, 0x99, 0xef, 0xc0, 0x57, 0xe4,
	0x6b, 0x20, 0x3b, 0x69, 0x48, 0xa4, 0x5d, 0xd9, 0x31, 0xcf, 0xff, 0xf7, 0xbc, 0xfa, 0x1f, 0x78,
	0x6c, 0x95, 0x66, 0x02, 0x1d, 0x93, 0x85, 0xcc, 0xbc, 0x32, 0x25, 0x13, 0x7c, 0x49, 0xad, 0x33,
	0xde, 0x90, 0x79, 0xa6, 0x30, 0x33, 0x4c, 0x19, 0x64, 0x57, 0x8e, 0x29, 0x7b, 0x39, 0x66, 0x81,
	0x36, 0x56, 0x3a, 0x6a, 0x95, 0xa6, 0x3c, 0xf3, 0xea, 0x52, 0xd2, 0x5c, 0x2e, 0xf8, 0xaa, 0xf0,
	0x2c, 0x33, 0xa5, 0x97, 0x57, 0x9e, 0x0a, 0x74, 0xb4, 0x5d, 0x11, 0x3b, 0x5f, 0xc3, 0x53, 0x38,
	0xbc, 0xa9, 0x27, 0xfb, 0x74, 0xfe, 0x63, 0x46, 0x8e, 0xa0, 0x17, 0x44, 0xcc, 0x8c, 0x95, 0x69,
	0x32, 0x48, 0x46, 0xdb, 0xd3, 0x4d, 0xab, 0xf4, 0x2c, 0x7c, 0x0f, 0x35, 0x6c, 0x05, 0x91, 0xe7,
	0xb9, 0xf3, 0xd7, 0x56, 0x92, 0x47, 0x70, 0x9f, 0x2f, 0x58, 0xc9, 0x75, 0x85, 0xf6, 0xa6, 0x1b,
	0x7c, 0xf1, 0x95, 0x6b, 0x49, 0x9e, 0xc2, 0x56, 0x9c, 0x36, 0x90, 0x12, 0x31, 0xbd, 0x13, 0xd5,
	0x7e, 0x88, 0xbd, 0xaf, 0x42, 0x35, 0x32, 0x69, 0x90, 0xbb, 0x0d, 0x32, 0xa9, 0x91, 0xe1, 0x9f,
	0x7b, 0xb0, 0x7f, 0xd3, 0xa4, 0xe4, 0x57, 0x02, 0xfd, 0x10, 0x5c, 0xe7, 0x1e, 0x0f, 0x92, 0x51,
	0xff, 0x58, 0xd0, 0xff, 0x71, 0x2e, 0xda, 0xde, 0x78, 0x0a, 0x02, 0x5d, 0x6b, 0x83, 0x00, 0x5a,
	0xa7, 0x8c, 0x53, 0xfe, 0x3a, 0x3d, 0x89, 0xd7, 0x0a, 0x83, 0x7d, 0xab, 0x43, 0xe4, 0x19, 0x3c,
	0x08, 0x88, 0xe6, 0x78, 0xc1, 0x0a, 0x59, 0x2e, 0xfd, 0xcf, 0x74, 0x1c, 0xa9, 0x6d, 0x81, 0xee,
	0x0b, 0xc7, 0x8b, 0xcf, 0x31, 0x48, 0x9e, 0x54, 0xfb, 0xac, 0x2c, 0xf3, 0x4a, 0xcb, 0xf4, 0x65,
	0x64, 0x7a, 0x02, 0xdd, 0x77, 0x3b, 0x57, 0x5a, 0x92, 0xe7, 0xb0, 0x2b, 0x8c, 0xf1, 0xe8, 0x1d,
	0xaf, 0x10, 0xb3, 0xf2, 0xe9, 0x24, 0x52, 0x3b, 0x8d, 0x30, 0xaf, 0xe2, 0x84, 0xc2, 0x5e, 0xc6,
	0xcb, 0x5c, 0xe5, 0xdc, 0xcb, 0x78, 0x3b, 0xf4, 0xdc, 0xcb, 0xf4, 0x55, 0xc4, 0x77, 0x1b, 0xe9,
	0x03, 0xba, 0x59, 0x10, 0xc8, 0x0b, 0x20, 0x9d, 0x0b, 0x57, 0xf8, 0x69, 0x5d, 0x1d, 0xdd, 0x79,
	0x2d, 0x54, 0xf4, 0x11, 0xf4, 0x62, 0xcd, 0x68, 0x90, 0xb3, 0xca, 0x20, 0x02, 0x5d, 0x34, 0x48,
	0x28, 0xd5, 0x6d, 0xbd, 0x28, 0xf8, 0x32, 0x7d, 0x3d, 0x48, 0x46, 0x9b, 0xd3, 0x9d, 0x76, 0xe7,
	0x8f, 0x05, 0x5f, 0x92, 0x31, 0x1c, 0x74, 0xe9, 0xe6, 0x94, 0x6f, 0x62, 0xdd, 0xfd, 0x76, 0x46,
	0x73, 0xd3, 0xdf, 0x09, 0x3c, 0xec, 0xa6, 0xad, 0x6d, 0xf0, 0xf6, 0xd6, 0x6c, 0xb0, 0xd7, 0x9e,
	0x6c, 0xed, 0x87, 0x33, 0x38, 0xec, 0xce, 0xd5, 0x7e, 0xf6, 0x77, 0x71, 0xa3, 0x83, 0x76, 0xde,
	0xbf, 0xf7, 0x17, 0x1b, 0xf1, 0x7f, 0x3f, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x46, 0x2d, 0x70,
	0x65, 0x0d, 0x04, 0x00, 0x00,
}
