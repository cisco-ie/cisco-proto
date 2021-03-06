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
// source: ldp_nsr_sum.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_nsr_ha_summary

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

type LdpNsrSum_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrSum_KEYS) Reset()         { *m = LdpNsrSum_KEYS{} }
func (m *LdpNsrSum_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpNsrSum_KEYS) ProtoMessage()    {}
func (*LdpNsrSum_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed47cdc4b337e857, []int{0}
}

func (m *LdpNsrSum_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrSum_KEYS.Unmarshal(m, b)
}
func (m *LdpNsrSum_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrSum_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpNsrSum_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrSum_KEYS.Merge(m, src)
}
func (m *LdpNsrSum_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpNsrSum_KEYS.Size(m)
}
func (m *LdpNsrSum_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrSum_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrSum_KEYS proto.InternalMessageInfo

func (m *LdpNsrSum_KEYS) GetNodeName() string {
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
	return fileDescriptor_ed47cdc4b337e857, []int{1}
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

type LdpNsrSumSess struct {
	Total                uint32   `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	NsrEligible          uint32   `protobuf:"varint,2,opt,name=nsr_eligible,json=nsrEligible,proto3" json:"nsr_eligible,omitempty"`
	NsrStateNone         uint32   `protobuf:"varint,3,opt,name=nsr_state_none,json=nsrStateNone,proto3" json:"nsr_state_none,omitempty"`
	NsrStateWait         uint32   `protobuf:"varint,4,opt,name=nsr_state_wait,json=nsrStateWait,proto3" json:"nsr_state_wait,omitempty"`
	NsrStateReady        uint32   `protobuf:"varint,5,opt,name=nsr_state_ready,json=nsrStateReady,proto3" json:"nsr_state_ready,omitempty"`
	NsrStatePrepare      uint32   `protobuf:"varint,6,opt,name=nsr_state_prepare,json=nsrStatePrepare,proto3" json:"nsr_state_prepare,omitempty"`
	NsrStateAppWait      uint32   `protobuf:"varint,7,opt,name=nsr_state_app_wait,json=nsrStateAppWait,proto3" json:"nsr_state_app_wait,omitempty"`
	NsrStateOperational  uint32   `protobuf:"varint,8,opt,name=nsr_state_operational,json=nsrStateOperational,proto3" json:"nsr_state_operational,omitempty"`
	NsrStateTcpPhase1    uint32   `protobuf:"varint,9,opt,name=nsr_state_tcp_phase1,json=nsrStateTcpPhase1,proto3" json:"nsr_state_tcp_phase1,omitempty"`
	NsrStateTcpPhase2    uint32   `protobuf:"varint,10,opt,name=nsr_state_tcp_phase2,json=nsrStateTcpPhase2,proto3" json:"nsr_state_tcp_phase2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpNsrSumSess) Reset()         { *m = LdpNsrSumSess{} }
func (m *LdpNsrSumSess) String() string { return proto.CompactTextString(m) }
func (*LdpNsrSumSess) ProtoMessage()    {}
func (*LdpNsrSumSess) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed47cdc4b337e857, []int{2}
}

func (m *LdpNsrSumSess) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrSumSess.Unmarshal(m, b)
}
func (m *LdpNsrSumSess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrSumSess.Marshal(b, m, deterministic)
}
func (m *LdpNsrSumSess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrSumSess.Merge(m, src)
}
func (m *LdpNsrSumSess) XXX_Size() int {
	return xxx_messageInfo_LdpNsrSumSess.Size(m)
}
func (m *LdpNsrSumSess) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrSumSess.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrSumSess proto.InternalMessageInfo

func (m *LdpNsrSumSess) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrEligible() uint32 {
	if m != nil {
		return m.NsrEligible
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateNone() uint32 {
	if m != nil {
		return m.NsrStateNone
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateWait() uint32 {
	if m != nil {
		return m.NsrStateWait
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateReady() uint32 {
	if m != nil {
		return m.NsrStateReady
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStatePrepare() uint32 {
	if m != nil {
		return m.NsrStatePrepare
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateAppWait() uint32 {
	if m != nil {
		return m.NsrStateAppWait
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateOperational() uint32 {
	if m != nil {
		return m.NsrStateOperational
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateTcpPhase1() uint32 {
	if m != nil {
		return m.NsrStateTcpPhase1
	}
	return 0
}

func (m *LdpNsrSumSess) GetNsrStateTcpPhase2() uint32 {
	if m != nil {
		return m.NsrStateTcpPhase2
	}
	return 0
}

type LdpNsrSum struct {
	Vrf                  *LdpVrfInfo    `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	Sessions             *LdpNsrSumSess `protobuf:"bytes,51,opt,name=sessions,proto3" json:"sessions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LdpNsrSum) Reset()         { *m = LdpNsrSum{} }
func (m *LdpNsrSum) String() string { return proto.CompactTextString(m) }
func (*LdpNsrSum) ProtoMessage()    {}
func (*LdpNsrSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed47cdc4b337e857, []int{3}
}

func (m *LdpNsrSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpNsrSum.Unmarshal(m, b)
}
func (m *LdpNsrSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpNsrSum.Marshal(b, m, deterministic)
}
func (m *LdpNsrSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpNsrSum.Merge(m, src)
}
func (m *LdpNsrSum) XXX_Size() int {
	return xxx_messageInfo_LdpNsrSum.Size(m)
}
func (m *LdpNsrSum) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpNsrSum.DiscardUnknown(m)
}

var xxx_messageInfo_LdpNsrSum proto.InternalMessageInfo

func (m *LdpNsrSum) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpNsrSum) GetSessions() *LdpNsrSumSess {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpNsrSum_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.nsr.ha_summary.ldp_nsr_sum_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.nsr.ha_summary.ldp_vrf_info")
	proto.RegisterType((*LdpNsrSumSess)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.nsr.ha_summary.ldp_nsr_sum_sess")
	proto.RegisterType((*LdpNsrSum)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.nsr.ha_summary.ldp_nsr_sum")
}

func init() { proto.RegisterFile("ldp_nsr_sum.proto", fileDescriptor_ed47cdc4b337e857) }

var fileDescriptor_ed47cdc4b337e857 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x86, 0x71, 0xb2, 0xbb, 0x4d, 0x26, 0x9b, 0xdd, 0x46, 0xdd, 0x82, 0xa0, 0x97, 0xd4, 0x94,
	0x12, 0x5a, 0xd0, 0x52, 0xe7, 0x09, 0x7a, 0xc8, 0xa1, 0x14, 0xd2, 0xe0, 0x14, 0x42, 0xe9, 0x41,
	0x28, 0xb6, 0xdc, 0x08, 0x6c, 0x49, 0x48, 0x4a, 0xd2, 0xbc, 0x54, 0xdf, 0xad, 0x6f, 0x50, 0xa4,
	0xc4, 0xb5, 0x53, 0xba, 0xb7, 0x5c, 0x8c, 0x35, 0xfa, 0xbe, 0xf9, 0x0d, 0x33, 0x86, 0x51, 0x99,
	0x6b, 0x2a, 0xad, 0xa1, 0x76, 0x5b, 0x11, 0x6d, 0x94, 0x53, 0xe8, 0x53, 0x26, 0x6c, 0xa6, 0xa8,
	0x50, 0x96, 0xfe, 0x34, 0xb4, 0xd2, 0xa5, 0xa5, 0x1e, 0x52, 0x9a, 0x1b, 0x52, 0x9f, 0x88, 0x54,
	0x39, 0xb7, 0xe1, 0x49, 0x72, 0x5e, 0xb0, 0x6d, 0xe9, 0xe8, 0xce, 0x14, 0x44, 0x5a, 0x43, 0x36,
	0xcc, 0x37, 0xab, 0x98, 0x39, 0xc4, 0x8f, 0xf0, 0xbc, 0xd5, 0x9f, 0x7e, 0x9e, 0x7d, 0x5b, 0xa2,
	0x57, 0xd0, 0xf7, 0x1a, 0x95, 0xac, 0xe2, 0x38, 0x1a, 0x47, 0x93, 0x7e, 0xda, 0xf3, 0x85, 0x39,
	0xab, 0x78, 0x9c, 0xc0, 0xad, 0x17, 0x76, 0xa6, 0xa0, 0x42, 0x16, 0x0a, 0x21, 0xb8, 0x6a, 0x71,
	0xe1, 0x1d, 0xdd, 0x41, 0x47, 0xe4, 0xb8, 0x33, 0x8e, 0x26, 0xc3, 0xb4, 0x23, 0xf2, 0xf8, 0x57,
	0xf7, 0x3c, 0xc5, 0x72, 0x6b, 0xd1, 0x03, 0x5c, 0x3b, 0xe5, 0x58, 0x19, 0xcc, 0x61, 0x7a, 0x3c,
	0xa0, 0xd7, 0x70, 0xeb, 0x29, 0x5e, 0x8a, 0x1f, 0x62, 0x5d, 0xf2, 0x53, 0x93, 0x81, 0xb4, 0x66,
	0x76, 0x2a, 0xa1, 0x37, 0x70, 0x17, 0x1a, 0x39, 0xe6, 0x38, 0x95, 0x4a, 0x72, 0xdc, 0x0d, 0x90,
	0x17, 0x97, 0xbe, 0x38, 0x57, 0xf2, 0x1f, 0x6a, 0xcf, 0x84, 0xc3, 0x57, 0xe7, 0xd4, 0x8a, 0x09,
	0x87, 0xde, 0xc2, 0x7d, 0x43, 0x19, 0xce, 0xf2, 0x03, 0xbe, 0x0e, 0xd8, 0xb0, 0xc6, 0x52, 0x5f,
	0x44, 0xef, 0x60, 0xd4, 0x70, 0xda, 0x70, 0xcd, 0x0c, 0xc7, 0x37, 0x81, 0xbc, 0xaf, 0xc9, 0xc5,
	0xb1, 0x8c, 0xde, 0x03, 0x6a, 0x58, 0xa6, 0xf5, 0x31, 0xfd, 0xd9, 0x39, 0xfc, 0x51, 0xeb, 0xf0,
	0x01, 0x09, 0xbc, 0x6c, 0x60, 0x3f, 0x40, 0xe6, 0x84, 0x92, 0xac, 0xc4, 0xbd, 0xc0, 0xbf, 0xa8,
	0xf9, 0x2f, 0xcd, 0x15, 0x7a, 0x84, 0x87, 0xc6, 0x71, 0x99, 0xa6, 0x7a, 0xc3, 0x2c, 0xff, 0x80,
	0xfb, 0x41, 0x19, 0xd5, 0xca, 0xd7, 0x4c, 0x2f, 0xc2, 0xc5, 0x13, 0x42, 0x82, 0xe1, 0xff, 0x42,
	0x12, 0xff, 0x8e, 0x60, 0xd0, 0x1a, 0x18, 0x12, 0xd0, 0xdd, 0x99, 0x02, 0x27, 0xe3, 0x68, 0x32,
	0x48, 0x56, 0xe4, 0x62, 0xeb, 0x47, 0xda, 0xab, 0x94, 0xfa, 0x0c, 0xb4, 0x87, 0x9e, 0x5f, 0x0f,
	0xa1, 0xa4, 0xc5, 0xd3, 0x90, 0xf7, 0xfd, 0xc2, 0x79, 0xed, 0x2d, 0x4c, 0xff, 0x86, 0xad, 0x6f,
	0xc2, 0xbf, 0x35, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x00, 0x59, 0x2d, 0x70, 0x03, 0x00,
	0x00,
}
