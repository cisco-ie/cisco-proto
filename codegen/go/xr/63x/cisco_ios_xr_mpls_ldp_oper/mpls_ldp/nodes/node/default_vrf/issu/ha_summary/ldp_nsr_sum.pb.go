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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_issu_ha_summary

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
	proto.RegisterType((*LdpNsrSum_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.issu.ha_summary.ldp_nsr_sum_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.issu.ha_summary.ldp_vrf_info")
	proto.RegisterType((*LdpNsrSumSess)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.issu.ha_summary.ldp_nsr_sum_sess")
	proto.RegisterType((*LdpNsrSum)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.issu.ha_summary.ldp_nsr_sum")
}

func init() { proto.RegisterFile("ldp_nsr_sum.proto", fileDescriptor_ed47cdc4b337e857) }

var fileDescriptor_ed47cdc4b337e857 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0xaa, 0xd3, 0x30,
	0x18, 0xc7, 0xe9, 0x76, 0xce, 0x71, 0xfb, 0x76, 0x76, 0x8e, 0x8b, 0x47, 0x08, 0x78, 0x33, 0x8b,
	0xc8, 0x50, 0xc8, 0xb0, 0x7b, 0x02, 0x2f, 0x76, 0xa3, 0x30, 0x47, 0x27, 0xa8, 0x20, 0x84, 0xac,
	0x4d, 0x5d, 0xa4, 0x4d, 0x42, 0x92, 0xcd, 0xed, 0xa5, 0x7c, 0x39, 0x5f, 0x40, 0x92, 0xad, 0xb6,
	0x13, 0xbd, 0xdb, 0x4d, 0x69, 0xfe, 0xf9, 0xfd, 0xf2, 0x2f, 0xe4, 0x2b, 0x8c, 0xca, 0x5c, 0x53,
	0x69, 0x0d, 0xb5, 0xdb, 0x8a, 0x68, 0xa3, 0x9c, 0x42, 0xef, 0x32, 0x61, 0x33, 0x45, 0x85, 0xb2,
	0x74, 0x6f, 0x68, 0xa5, 0x4b, 0x4b, 0x3d, 0xa4, 0x34, 0x37, 0xa4, 0x5e, 0x11, 0xa9, 0x72, 0x6e,
	0xc3, 0x93, 0xe4, 0xbc, 0x60, 0xdb, 0xd2, 0xd1, 0x9d, 0x29, 0x88, 0xb0, 0x76, 0x4b, 0x36, 0xcc,
	0x9f, 0x56, 0x31, 0x73, 0x88, 0xa7, 0xf0, 0xb8, 0x55, 0x40, 0xdf, 0xcf, 0xbf, 0xac, 0xd0, 0x33,
	0xe8, 0x7b, 0x8f, 0x4a, 0x56, 0x71, 0x1c, 0x8d, 0xa3, 0x49, 0x3f, 0xed, 0xf9, 0x60, 0xc1, 0x2a,
	0x1e, 0x27, 0x70, 0xeb, 0x85, 0x9d, 0x29, 0xa8, 0x90, 0x85, 0x42, 0x08, 0xae, 0x5a, 0x5c, 0x78,
	0x47, 0x77, 0xd0, 0x11, 0x39, 0xee, 0x8c, 0xa3, 0xc9, 0x30, 0xed, 0x88, 0x3c, 0xfe, 0xd9, 0x3d,
	0x6f, 0xb1, 0xdc, 0x5a, 0xf4, 0x00, 0xd7, 0x4e, 0x39, 0x56, 0x06, 0x73, 0x98, 0x1e, 0x17, 0xe8,
	0x39, 0xdc, 0x7a, 0x8a, 0x97, 0xe2, 0x9b, 0x58, 0x97, 0xfc, 0x74, 0xc8, 0x40, 0x5a, 0x33, 0x3f,
	0x45, 0xe8, 0x05, 0xdc, 0x85, 0x83, 0x1c, 0x73, 0x9c, 0x4a, 0x25, 0x39, 0xee, 0x06, 0xc8, 0x8b,
	0x2b, 0x1f, 0x2e, 0x94, 0xfc, 0x8b, 0xfa, 0xc1, 0x84, 0xc3, 0x57, 0xe7, 0xd4, 0x27, 0x26, 0x1c,
	0x7a, 0x09, 0xf7, 0x0d, 0x65, 0x38, 0xcb, 0x0f, 0xf8, 0x3a, 0x60, 0xc3, 0x1a, 0x4b, 0x7d, 0x88,
	0x5e, 0xc1, 0xa8, 0xe1, 0xb4, 0xe1, 0x9a, 0x19, 0x8e, 0x6f, 0x02, 0x79, 0x5f, 0x93, 0xcb, 0x63,
	0x8c, 0x5e, 0x03, 0x6a, 0x58, 0xa6, 0xf5, 0xb1, 0xfd, 0xd1, 0x39, 0xfc, 0x56, 0xeb, 0xf0, 0x01,
	0x09, 0x3c, 0x6d, 0x60, 0x7f, 0x83, 0xcc, 0x09, 0x25, 0x59, 0x89, 0x7b, 0x81, 0x7f, 0x52, 0xf3,
	0x1f, 0x9a, 0x2d, 0x34, 0x85, 0x87, 0xc6, 0x71, 0x99, 0xa6, 0x7a, 0xc3, 0x2c, 0x7f, 0x83, 0xfb,
	0x41, 0x19, 0xd5, 0xca, 0xc7, 0x4c, 0x2f, 0xc3, 0xc6, 0x7f, 0x84, 0x04, 0xc3, 0xbf, 0x85, 0x24,
	0xfe, 0x15, 0xc1, 0xa0, 0x75, 0x61, 0xe8, 0x3b, 0x74, 0x77, 0xa6, 0xc0, 0xc9, 0x38, 0x9a, 0x0c,
	0x92, 0xcf, 0xe4, 0x72, 0xf3, 0x47, 0xda, 0xb3, 0x94, 0xfa, 0x12, 0xb4, 0x87, 0x9e, 0x9f, 0x0f,
	0xa1, 0xa4, 0xc5, 0xb3, 0x50, 0xf8, 0xf5, 0xd2, 0x85, 0xed, 0x39, 0x4c, 0xff, 0xb4, 0xad, 0x6f,
	0xc2, 0xef, 0x35, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x23, 0x71, 0xa8, 0x16, 0x73, 0x03, 0x00,
	0x00,
}