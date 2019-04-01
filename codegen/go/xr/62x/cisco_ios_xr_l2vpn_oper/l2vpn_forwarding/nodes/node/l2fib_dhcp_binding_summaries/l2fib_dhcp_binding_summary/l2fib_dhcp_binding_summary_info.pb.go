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
// source: l2fib_dhcp_binding_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_dhcp_binding_summaries_l2fib_dhcp_binding_summary

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

type L2FibDhcpBindingSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Xcid                 string   `protobuf:"bytes,2,opt,name=xcid,proto3" json:"xcid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibDhcpBindingSummaryInfo_KEYS) Reset()         { *m = L2FibDhcpBindingSummaryInfo_KEYS{} }
func (m *L2FibDhcpBindingSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibDhcpBindingSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2FibDhcpBindingSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e022ec1fd2da9fed, []int{0}
}

func (m *L2FibDhcpBindingSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibDhcpBindingSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibDhcpBindingSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibDhcpBindingSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS.Size(m)
}
func (m *L2FibDhcpBindingSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibDhcpBindingSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibDhcpBindingSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibDhcpBindingSummaryInfo_KEYS) GetXcid() string {
	if m != nil {
		return m.Xcid
	}
	return ""
}

type L2FibAcKeyInfo struct {
	InterfaceHandle      string   `protobuf:"bytes,1,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibAcKeyInfo) Reset()         { *m = L2FibAcKeyInfo{} }
func (m *L2FibAcKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibAcKeyInfo) ProtoMessage()    {}
func (*L2FibAcKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e022ec1fd2da9fed, []int{1}
}

func (m *L2FibAcKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibAcKeyInfo.Unmarshal(m, b)
}
func (m *L2FibAcKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibAcKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibAcKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibAcKeyInfo.Merge(m, src)
}
func (m *L2FibAcKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibAcKeyInfo.Size(m)
}
func (m *L2FibAcKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibAcKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibAcKeyInfo proto.InternalMessageInfo

func (m *L2FibAcKeyInfo) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

type L2FibPwKeyInfo struct {
	PwId                 uint64   `protobuf:"varint,1,opt,name=pw_id,json=pwId,proto3" json:"pw_id,omitempty"`
	NextHopAddress       string   `protobuf:"bytes,2,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	PseudoWireIdType     string   `protobuf:"bytes,3,opt,name=pseudo_wire_id_type,json=pseudoWireIdType,proto3" json:"pseudo_wire_id_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibPwKeyInfo) Reset()         { *m = L2FibPwKeyInfo{} }
func (m *L2FibPwKeyInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibPwKeyInfo) ProtoMessage()    {}
func (*L2FibPwKeyInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e022ec1fd2da9fed, []int{2}
}

func (m *L2FibPwKeyInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibPwKeyInfo.Unmarshal(m, b)
}
func (m *L2FibPwKeyInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibPwKeyInfo.Marshal(b, m, deterministic)
}
func (m *L2FibPwKeyInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibPwKeyInfo.Merge(m, src)
}
func (m *L2FibPwKeyInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibPwKeyInfo.Size(m)
}
func (m *L2FibPwKeyInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibPwKeyInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibPwKeyInfo proto.InternalMessageInfo

func (m *L2FibPwKeyInfo) GetPwId() uint64 {
	if m != nil {
		return m.PwId
	}
	return 0
}

func (m *L2FibPwKeyInfo) GetNextHopAddress() string {
	if m != nil {
		return m.NextHopAddress
	}
	return ""
}

func (m *L2FibPwKeyInfo) GetPseudoWireIdType() string {
	if m != nil {
		return m.PseudoWireIdType
	}
	return ""
}

type L2FibDhcpBindUn struct {
	DataType             string          `protobuf:"bytes,1,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty"`
	Ac                   *L2FibAcKeyInfo `protobuf:"bytes,2,opt,name=ac,proto3" json:"ac,omitempty"`
	Pw                   *L2FibPwKeyInfo `protobuf:"bytes,3,opt,name=pw,proto3" json:"pw,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *L2FibDhcpBindUn) Reset()         { *m = L2FibDhcpBindUn{} }
func (m *L2FibDhcpBindUn) String() string { return proto.CompactTextString(m) }
func (*L2FibDhcpBindUn) ProtoMessage()    {}
func (*L2FibDhcpBindUn) Descriptor() ([]byte, []int) {
	return fileDescriptor_e022ec1fd2da9fed, []int{3}
}

func (m *L2FibDhcpBindUn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibDhcpBindUn.Unmarshal(m, b)
}
func (m *L2FibDhcpBindUn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibDhcpBindUn.Marshal(b, m, deterministic)
}
func (m *L2FibDhcpBindUn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibDhcpBindUn.Merge(m, src)
}
func (m *L2FibDhcpBindUn) XXX_Size() int {
	return xxx_messageInfo_L2FibDhcpBindUn.Size(m)
}
func (m *L2FibDhcpBindUn) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibDhcpBindUn.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibDhcpBindUn proto.InternalMessageInfo

func (m *L2FibDhcpBindUn) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *L2FibDhcpBindUn) GetAc() *L2FibAcKeyInfo {
	if m != nil {
		return m.Ac
	}
	return nil
}

func (m *L2FibDhcpBindUn) GetPw() *L2FibPwKeyInfo {
	if m != nil {
		return m.Pw
	}
	return nil
}

type L2FibDhcpBindingSummaryInfo struct {
	Port                 *L2FibDhcpBindUn `protobuf:"bytes,50,opt,name=port,proto3" json:"port,omitempty"`
	Bindings             uint32           `protobuf:"varint,51,opt,name=bindings,proto3" json:"bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2FibDhcpBindingSummaryInfo) Reset()         { *m = L2FibDhcpBindingSummaryInfo{} }
func (m *L2FibDhcpBindingSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibDhcpBindingSummaryInfo) ProtoMessage()    {}
func (*L2FibDhcpBindingSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e022ec1fd2da9fed, []int{4}
}

func (m *L2FibDhcpBindingSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibDhcpBindingSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibDhcpBindingSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibDhcpBindingSummaryInfo.Merge(m, src)
}
func (m *L2FibDhcpBindingSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibDhcpBindingSummaryInfo.Size(m)
}
func (m *L2FibDhcpBindingSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibDhcpBindingSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibDhcpBindingSummaryInfo proto.InternalMessageInfo

func (m *L2FibDhcpBindingSummaryInfo) GetPort() *L2FibDhcpBindUn {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *L2FibDhcpBindingSummaryInfo) GetBindings() uint32 {
	if m != nil {
		return m.Bindings
	}
	return 0
}

func init() {
	proto.RegisterType((*L2FibDhcpBindingSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_dhcp_binding_summaries.l2fib_dhcp_binding_summary.l2fib_dhcp_binding_summary_info_KEYS")
	proto.RegisterType((*L2FibAcKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_dhcp_binding_summaries.l2fib_dhcp_binding_summary.l2fib_ac_key_info")
	proto.RegisterType((*L2FibPwKeyInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_dhcp_binding_summaries.l2fib_dhcp_binding_summary.l2fib_pw_key_info")
	proto.RegisterType((*L2FibDhcpBindUn)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_dhcp_binding_summaries.l2fib_dhcp_binding_summary.l2fib_dhcp_bind_un")
	proto.RegisterType((*L2FibDhcpBindingSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_dhcp_binding_summaries.l2fib_dhcp_binding_summary.l2fib_dhcp_binding_summary_info")
}

func init() {
	proto.RegisterFile("l2fib_dhcp_binding_summary_info.proto", fileDescriptor_e022ec1fd2da9fed)
}

var fileDescriptor_e022ec1fd2da9fed = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0x41, 0x6f, 0xda, 0x30,
	0x14, 0xc7, 0x95, 0x2c, 0x63, 0xe0, 0x69, 0x1b, 0x33, 0x87, 0x45, 0xdb, 0x61, 0x28, 0xda, 0x24,
	0x7a, 0x68, 0x0e, 0xe1, 0x5e, 0xa9, 0x87, 0x4a, 0xa0, 0xde, 0x42, 0xa5, 0xaa, 0xa7, 0x27, 0x63,
	0x3b, 0xc5, 0x02, 0x6c, 0xcb, 0x0e, 0x0d, 0xb9, 0xb4, 0xea, 0x17, 0xe8, 0xa5, 0x1f, 0x8d, 0x2f,
	0x54, 0xc5, 0x01, 0x8a, 0x5a, 0x21, 0x6e, 0xe5, 0x12, 0xd9, 0x7f, 0xfb, 0xfd, 0xdf, 0xef, 0xbd,
	0x17, 0xa3, 0xff, 0xb3, 0x24, 0x13, 0x63, 0x60, 0x13, 0xaa, 0x61, 0x2c, 0x24, 0x13, 0xf2, 0x16,
	0xec, 0x62, 0x3e, 0x27, 0xa6, 0x04, 0x21, 0x33, 0x15, 0x6b, 0xa3, 0x72, 0x85, 0xa7, 0x54, 0x58,
	0xaa, 0x40, 0x28, 0x0b, 0x4b, 0x03, 0xb3, 0xe4, 0x4e, 0x4b, 0x50, 0x9a, 0x9b, 0xb8, 0x5e, 0x66,
	0xca, 0x14, 0xc4, 0x54, 0xc1, 0xb1, 0x54, 0x8c, 0x5b, 0xf7, 0x8d, 0xf7, 0x5a, 0x0b, 0x6e, 0xf7,
	0x1f, 0x96, 0xd1, 0x08, 0xfd, 0x3b, 0x40, 0x05, 0x97, 0x17, 0x37, 0x23, 0xfc, 0x0b, 0x7d, 0xa9,
	0xf2, 0x80, 0x60, 0xa1, 0xd7, 0xf5, 0x7a, 0xad, 0xb4, 0x51, 0x6d, 0x87, 0x0c, 0x63, 0x14, 0x2c,
	0xa9, 0x60, 0xa1, 0xef, 0x54, 0xb7, 0x8e, 0xce, 0xd0, 0xcf, 0xda, 0x94, 0x50, 0x98, 0xf2, 0xda,
	0x06, 0x9f, 0xa0, 0xb6, 0x90, 0x39, 0x37, 0x19, 0xa1, 0x1c, 0x26, 0x44, 0xb2, 0x19, 0x5f, 0x5b,
	0xfd, 0xd8, 0xea, 0x03, 0x27, 0x47, 0x8f, 0xde, 0xc6, 0x40, 0x17, 0xaf, 0x06, 0x1d, 0xf4, 0x59,
	0x17, 0x1b, 0x80, 0x20, 0x0d, 0x74, 0x31, 0x64, 0xb8, 0x87, 0xda, 0x92, 0x2f, 0x73, 0x98, 0x28,
	0x0d, 0x84, 0x31, 0xc3, 0xad, 0x5d, 0xa3, 0x7c, 0xaf, 0xf4, 0x81, 0xd2, 0xe7, 0xb5, 0x8a, 0x4f,
	0x51, 0x47, 0x5b, 0xbe, 0x60, 0x0a, 0x0a, 0x61, 0xaa, 0x42, 0x20, 0x2f, 0x35, 0x0f, 0x3f, 0xb9,
	0xcb, 0xed, 0xfa, 0xe8, 0x5a, 0x18, 0x3e, 0x64, 0x57, 0xa5, 0xe6, 0xd1, 0xca, 0x47, 0xf8, 0x4d,
	0x67, 0x60, 0x21, 0xf1, 0x1f, 0xd4, 0x62, 0x24, 0x27, 0x75, 0x6c, 0x8d, 0xdf, 0xac, 0x84, 0x2a,
	0x06, 0x3f, 0x79, 0xc8, 0x27, 0xd4, 0xe5, 0xff, 0x9a, 0xdc, 0xc7, 0x1f, 0x38, 0xc7, 0xf8, 0x5d,
	0xbf, 0x53, 0x9f, 0x50, 0x07, 0xa4, 0x0b, 0x57, 0xe3, 0x71, 0x80, 0x76, 0xe6, 0x97, 0xfa, 0xba,
	0x88, 0x56, 0x1e, 0xfa, 0x7b, 0xe0, 0x7f, 0xc3, 0xcf, 0x1e, 0x0a, 0xb4, 0x32, 0x79, 0x98, 0x38,
	0xec, 0x87, 0x23, 0x60, 0xef, 0x8e, 0x3c, 0x75, 0x30, 0xf8, 0x37, 0x6a, 0xae, 0xef, 0xda, 0xb0,
	0xdf, 0xf5, 0x7a, 0xdf, 0xd2, 0xed, 0x7e, 0xdc, 0x70, 0x0f, 0xb7, 0xff, 0x12, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0x1b, 0x4c, 0x3f, 0xe1, 0x03, 0x00, 0x00,
}
