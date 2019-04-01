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
// source: l2fib_main_port_mstp_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_mstp_summaries_l2fib_mstp_summary

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

type L2FibMainPortMstpSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ParentInterface      string   `protobuf:"bytes,2,opt,name=parent_interface,json=parentInterface,proto3" json:"parent_interface,omitempty"`
	Msti                 uint32   `protobuf:"varint,3,opt,name=msti,proto3" json:"msti,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainPortMstpSummaryInfo_KEYS) Reset()         { *m = L2FibMainPortMstpSummaryInfo_KEYS{} }
func (m *L2FibMainPortMstpSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibMainPortMstpSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2FibMainPortMstpSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c5b7dcf1cb2190, []int{0}
}

func (m *L2FibMainPortMstpSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMainPortMstpSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMainPortMstpSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibMainPortMstpSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS.Size(m)
}
func (m *L2FibMainPortMstpSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainPortMstpSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMainPortMstpSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMainPortMstpSummaryInfo_KEYS) GetParentInterface() string {
	if m != nil {
		return m.ParentInterface
	}
	return ""
}

func (m *L2FibMainPortMstpSummaryInfo_KEYS) GetMsti() uint32 {
	if m != nil {
		return m.Msti
	}
	return 0
}

type L2FibBaseInfo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibBaseInfo) Reset()         { *m = L2FibBaseInfo{} }
func (m *L2FibBaseInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibBaseInfo) ProtoMessage()    {}
func (*L2FibBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c5b7dcf1cb2190, []int{1}
}

func (m *L2FibBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibBaseInfo.Unmarshal(m, b)
}
func (m *L2FibBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibBaseInfo.Marshal(b, m, deterministic)
}
func (m *L2FibBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibBaseInfo.Merge(m, src)
}
func (m *L2FibBaseInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibBaseInfo.Size(m)
}
func (m *L2FibBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibBaseInfo proto.InternalMessageInfo

type L2FibMainPortMstpSummaryInfo struct {
	Base                 *L2FibBaseInfo `protobuf:"bytes,50,opt,name=base,proto3" json:"base,omitempty"`
	ParentInterfaceXr    string         `protobuf:"bytes,51,opt,name=parent_interface_xr,json=parentInterfaceXr,proto3" json:"parent_interface_xr,omitempty"`
	MstiXr               uint32         `protobuf:"varint,52,opt,name=msti_xr,json=mstiXr,proto3" json:"msti_xr,omitempty"`
	State                uint32         `protobuf:"varint,53,opt,name=state,proto3" json:"state,omitempty"`
	BridgePortCount      uint32         `protobuf:"varint,54,opt,name=bridge_port_count,json=bridgePortCount,proto3" json:"bridge_port_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *L2FibMainPortMstpSummaryInfo) Reset()         { *m = L2FibMainPortMstpSummaryInfo{} }
func (m *L2FibMainPortMstpSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMainPortMstpSummaryInfo) ProtoMessage()    {}
func (*L2FibMainPortMstpSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_09c5b7dcf1cb2190, []int{2}
}

func (m *L2FibMainPortMstpSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibMainPortMstpSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMainPortMstpSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainPortMstpSummaryInfo.Merge(m, src)
}
func (m *L2FibMainPortMstpSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMainPortMstpSummaryInfo.Size(m)
}
func (m *L2FibMainPortMstpSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainPortMstpSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainPortMstpSummaryInfo proto.InternalMessageInfo

func (m *L2FibMainPortMstpSummaryInfo) GetBase() *L2FibBaseInfo {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *L2FibMainPortMstpSummaryInfo) GetParentInterfaceXr() string {
	if m != nil {
		return m.ParentInterfaceXr
	}
	return ""
}

func (m *L2FibMainPortMstpSummaryInfo) GetMstiXr() uint32 {
	if m != nil {
		return m.MstiXr
	}
	return 0
}

func (m *L2FibMainPortMstpSummaryInfo) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *L2FibMainPortMstpSummaryInfo) GetBridgePortCount() uint32 {
	if m != nil {
		return m.BridgePortCount
	}
	return 0
}

func init() {
	proto.RegisterType((*L2FibMainPortMstpSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mstp_summaries.l2fib_mstp_summary.l2fib_main_port_mstp_summary_info_KEYS")
	proto.RegisterType((*L2FibBaseInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mstp_summaries.l2fib_mstp_summary.l2fib_base_info")
	proto.RegisterType((*L2FibMainPortMstpSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mstp_summaries.l2fib_mstp_summary.l2fib_main_port_mstp_summary_info")
}

func init() {
	proto.RegisterFile("l2fib_main_port_mstp_summary_info.proto", fileDescriptor_09c5b7dcf1cb2190)
}

var fileDescriptor_09c5b7dcf1cb2190 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x65, 0x6b, 0xad, 0x18, 0x91, 0xda, 0x28, 0x98, 0x63, 0xed, 0x41, 0xab, 0x87, 0x1c, 0xb6,
	0xea, 0x0f, 0x88, 0x87, 0xe2, 0x45, 0xea, 0xa5, 0xe2, 0x61, 0xc8, 0xee, 0x66, 0x4b, 0xa0, 0x9b,
	0x84, 0x49, 0xaa, 0xed, 0xad, 0xbf, 0xe4, 0x1f, 0x4a, 0x12, 0x45, 0x58, 0x0f, 0x3d, 0x79, 0x09,
	0x33, 0x2f, 0x93, 0x37, 0xef, 0x3d, 0x42, 0xae, 0x96, 0x79, 0xad, 0x0a, 0x68, 0x84, 0xd2, 0x60,
	0x0d, 0x7a, 0x68, 0x9c, 0xb7, 0xe0, 0x56, 0x4d, 0x23, 0x70, 0x03, 0x4a, 0xd7, 0x86, 0x5b, 0x34,
	0xde, 0xd0, 0xb7, 0x52, 0xb9, 0xd2, 0x80, 0x32, 0x0e, 0xd6, 0x08, 0xcb, 0xfc, 0xdd, 0x6a, 0x30,
	0x56, 0x22, 0x4f, 0x65, 0x6d, 0xf0, 0x43, 0x60, 0xa5, 0xf4, 0x82, 0x6b, 0x53, 0x49, 0x17, 0x4f,
	0xfe, 0x4d, 0xfe, 0x4b, 0xa9, 0xa4, 0xfb, 0x0b, 0x6e, 0x46, 0xdb, 0x8c, 0x5c, 0xee, 0x14, 0x02,
	0x4f, 0x8f, 0xaf, 0x2f, 0xf4, 0x9c, 0x1c, 0x04, 0x6a, 0x50, 0x15, 0xcb, 0x86, 0xd9, 0xf8, 0x70,
	0xd6, 0x0b, 0xed, 0xb4, 0xa2, 0xd7, 0xe4, 0xc4, 0x0a, 0x94, 0xda, 0x83, 0xd2, 0x5e, 0x62, 0x2d,
	0x4a, 0xc9, 0x3a, 0x71, 0xa2, 0x9f, 0xf0, 0xe9, 0x0f, 0x4c, 0x29, 0xe9, 0x36, 0xce, 0x2b, 0xb6,
	0x37, 0xcc, 0xc6, 0xc7, 0xb3, 0x58, 0x8f, 0x06, 0xa4, 0x9f, 0x14, 0x14, 0xc2, 0xc9, 0xb8, 0x6f,
	0xf4, 0xd9, 0x21, 0x17, 0x3b, 0x55, 0xd1, 0x6d, 0x46, 0xba, 0xe1, 0x0d, 0xcb, 0x87, 0xd9, 0xf8,
	0x28, 0x5f, 0xf2, 0x7f, 0x0c, 0x8a, 0xb7, 0x24, 0xce, 0xe2, 0x66, 0xca, 0xc9, 0x69, 0xdb, 0x3a,
	0xac, 0x91, 0x4d, 0xa2, 0xfb, 0x41, 0xcb, 0xfd, 0x1c, 0x43, 0x86, 0xc1, 0x73, 0x98, 0xb9, 0x8d,
	0x11, 0xf4, 0x42, 0x3b, 0x47, 0x7a, 0x46, 0xf6, 0x9d, 0x17, 0x5e, 0xb2, 0xbb, 0x08, 0xa7, 0x86,
	0xde, 0x90, 0x41, 0x81, 0xaa, 0x5a, 0xc8, 0x14, 0x41, 0x69, 0x56, 0xda, 0xb3, 0xfb, 0x38, 0xd1,
	0x4f, 0x17, 0xcf, 0x06, 0xfd, 0x43, 0x80, 0x8b, 0x5e, 0xfc, 0x2d, 0x93, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x53, 0x96, 0x2f, 0xcc, 0x58, 0x02, 0x00, 0x00,
}
