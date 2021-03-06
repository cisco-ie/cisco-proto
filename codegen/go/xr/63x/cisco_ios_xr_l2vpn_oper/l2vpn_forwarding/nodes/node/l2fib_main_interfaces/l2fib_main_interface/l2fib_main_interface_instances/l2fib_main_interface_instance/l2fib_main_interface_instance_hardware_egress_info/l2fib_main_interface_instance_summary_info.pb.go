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
// source: l2fib_main_interface_instance_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_main_interfaces_l2fib_main_interface_l2fib_main_interface_instances_l2fib_main_interface_instance_l2fib_main_interface_instance_hardware_egress_info

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

type L2FibMainInterfaceInstanceSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	MainInterfaceId      uint32   `protobuf:"varint,2,opt,name=main_interface_id,json=mainInterfaceId,proto3" json:"main_interface_id,omitempty"`
	MainInterfaceType    string   `protobuf:"bytes,3,opt,name=main_interface_type,json=mainInterfaceType,proto3" json:"main_interface_type,omitempty"`
	Instance             uint32   `protobuf:"varint,4,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) Reset() {
	*m = L2FibMainInterfaceInstanceSummaryInfo_KEYS{}
}
func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) String() string {
	return proto.CompactTextString(m)
}
func (*L2FibMainInterfaceInstanceSummaryInfo_KEYS) ProtoMessage() {}
func (*L2FibMainInterfaceInstanceSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1495ba78f8609e64, []int{0}
}

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS.Size(m)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) GetMainInterfaceId() uint32 {
	if m != nil {
		return m.MainInterfaceId
	}
	return 0
}

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) GetMainInterfaceType() string {
	if m != nil {
		return m.MainInterfaceType
	}
	return ""
}

func (m *L2FibMainInterfaceInstanceSummaryInfo_KEYS) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

type L2FibMainInterfaceInstanceSummaryInfo struct {
	MainInterface        string   `protobuf:"bytes,50,opt,name=main_interface,json=mainInterface,proto3" json:"main_interface,omitempty"`
	MainIfType           uint32   `protobuf:"varint,51,opt,name=main_if_type,json=mainIfType,proto3" json:"main_if_type,omitempty"`
	VirtualIfName        string   `protobuf:"bytes,52,opt,name=virtual_if_name,json=virtualIfName,proto3" json:"virtual_if_name,omitempty"`
	Instance             uint32   `protobuf:"varint,53,opt,name=instance,proto3" json:"instance,omitempty"`
	State                string   `protobuf:"bytes,54,opt,name=state,proto3" json:"state,omitempty"`
	BridgePortCount      uint32   `protobuf:"varint,55,opt,name=bridge_port_count,json=bridgePortCount,proto3" json:"bridge_port_count,omitempty"`
	InstanceProvisioned  bool     `protobuf:"varint,56,opt,name=instance_provisioned,json=instanceProvisioned,proto3" json:"instance_provisioned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) Reset()         { *m = L2FibMainInterfaceInstanceSummaryInfo{} }
func (m *L2FibMainInterfaceInstanceSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceInstanceSummaryInfo) ProtoMessage()    {}
func (*L2FibMainInterfaceInstanceSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1495ba78f8609e64, []int{1}
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo.Merge(m, src)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo.Size(m)
}
func (m *L2FibMainInterfaceInstanceSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceInstanceSummaryInfo proto.InternalMessageInfo

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetMainInterface() string {
	if m != nil {
		return m.MainInterface
	}
	return ""
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetMainIfType() uint32 {
	if m != nil {
		return m.MainIfType
	}
	return 0
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetVirtualIfName() string {
	if m != nil {
		return m.VirtualIfName
	}
	return ""
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetBridgePortCount() uint32 {
	if m != nil {
		return m.BridgePortCount
	}
	return 0
}

func (m *L2FibMainInterfaceInstanceSummaryInfo) GetInstanceProvisioned() bool {
	if m != nil {
		return m.InstanceProvisioned
	}
	return false
}

func init() {
	proto.RegisterType((*L2FibMainInterfaceInstanceSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_instances.l2fib_main_interface_instance.l2fib_main_interface_instance_hardware_egress_info.l2fib_main_interface_instance_summary_info_KEYS")
	proto.RegisterType((*L2FibMainInterfaceInstanceSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_instances.l2fib_main_interface_instance.l2fib_main_interface_instance_hardware_egress_info.l2fib_main_interface_instance_summary_info")
}

func init() {
	proto.RegisterFile("l2fib_main_interface_instance_summary_info.proto", fileDescriptor_1495ba78f8609e64)
}

var fileDescriptor_1495ba78f8609e64 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6a, 0xe2, 0x40,
	0x14, 0x26, 0xee, 0xae, 0xeb, 0x0e, 0xeb, 0x8a, 0xa3, 0xb0, 0xc3, 0x5e, 0x05, 0x61, 0x17, 0xf1,
	0x22, 0xdb, 0x6a, 0xff, 0xee, 0x4b, 0x2f, 0x42, 0xa1, 0x88, 0xed, 0x4d, 0xaf, 0x0e, 0x63, 0x66,
	0x62, 0x07, 0xcc, 0x4c, 0x38, 0x33, 0x6a, 0x7d, 0xa9, 0x3e, 0x44, 0x1f, 0xa0, 0xcf, 0x54, 0x32,
	0xa9, 0x96, 0x88, 0x88, 0x37, 0x21, 0xe7, 0x7c, 0x5f, 0xbe, 0x9f, 0x43, 0xc8, 0xc9, 0x7c, 0x98,
	0xaa, 0x29, 0x64, 0x5c, 0x69, 0x50, 0xda, 0x49, 0x4c, 0x79, 0x22, 0x41, 0x69, 0xeb, 0xb8, 0x4e,
	0x24, 0xd8, 0x45, 0x96, 0x71, 0x5c, 0x83, 0xd2, 0xa9, 0x89, 0x72, 0x34, 0xce, 0xd0, 0xb7, 0x20,
	0x51, 0x36, 0x31, 0xa0, 0x8c, 0x85, 0x67, 0x84, 0xf9, 0x70, 0x99, 0x6b, 0x30, 0xb9, 0xc4, 0xa8,
	0x7c, 0x4d, 0x0d, 0xae, 0x38, 0x0a, 0xa5, 0x67, 0x91, 0x36, 0x42, 0x5a, 0xff, 0x8c, 0xf6, 0xd9,
	0xd8, 0xbd, 0xdb, 0xe8, 0x60, 0x22, 0x7b, 0x18, 0x3e, 0x8c, 0xc2, 0x13, 0x47, 0xb1, 0xe2, 0x28,
	0x41, 0xce, 0x50, 0x5a, 0xeb, 0x6b, 0xf5, 0x5e, 0x03, 0xf2, 0xff, 0xf8, 0x2b, 0xc0, 0xed, 0xcd,
	0xe3, 0x3d, 0xfd, 0x4d, 0xbe, 0x17, 0xb5, 0x40, 0x09, 0x16, 0x84, 0x41, 0xff, 0xc7, 0xa4, 0x5e,
	0x8c, 0xb1, 0xa0, 0x03, 0xd2, 0xde, 0x55, 0x11, 0xac, 0x16, 0x06, 0xfd, 0xe6, 0xa4, 0x55, 0x00,
	0xf1, 0x66, 0x1f, 0x0b, 0x1a, 0x91, 0xce, 0x0e, 0xd7, 0xad, 0x73, 0xc9, 0xbe, 0x78, 0xc1, 0x76,
	0x85, 0xfd, 0xb0, 0xce, 0x25, 0xfd, 0x43, 0x1a, 0x9b, 0x48, 0xec, 0xab, 0x97, 0xdc, 0xce, 0xbd,
	0x97, 0x1a, 0x19, 0x1c, 0x5f, 0x82, 0xfe, 0x25, 0xbf, 0xaa, 0x3c, 0x36, 0xf4, 0xae, 0xcd, 0x8a,
	0x2b, 0x0d, 0xc9, 0xcf, 0x92, 0x96, 0x96, 0xd1, 0x46, 0xde, 0x95, 0x78, 0x52, 0xea, 0x33, 0xfd,
	0x23, 0xad, 0xa5, 0x42, 0xb7, 0xe0, 0xf3, 0x82, 0xa4, 0x79, 0x26, 0xd9, 0x59, 0xa9, 0xf4, 0xb1,
	0x8e, 0xd3, 0x3b, 0x9e, 0x55, 0xb3, 0x9f, 0x57, 0xb3, 0xd3, 0x2e, 0xf9, 0x66, 0x1d, 0x77, 0x92,
	0x5d, 0xf8, 0x2f, 0xcb, 0xa1, 0xb8, 0xe4, 0x14, 0x95, 0x98, 0x49, 0xc8, 0x0d, 0x3a, 0x48, 0xcc,
	0x42, 0x3b, 0x76, 0x59, 0x5e, 0xb2, 0x04, 0xc6, 0x06, 0xdd, 0x75, 0xb1, 0xa6, 0xa7, 0xa4, 0xbb,
	0xed, 0x99, 0xa3, 0x59, 0x2a, 0xab, 0x8c, 0x96, 0x82, 0x5d, 0x85, 0x41, 0xbf, 0x31, 0xe9, 0x6c,
	0xb0, 0xf1, 0x27, 0x34, 0xad, 0xfb, 0xbf, 0x79, 0xf4, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xe3,
	0xfc, 0xf0, 0x01, 0x03, 0x00, 0x00,
}
