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
// source: l2fib_main_interface_summary_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_main_interfaces_l2fib_main_interface_l2fib_main_interface_hardware_ingress_info

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

type L2FibMainInterfaceSummaryInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	MainInterfaceId      uint32   `protobuf:"varint,2,opt,name=main_interface_id,json=mainInterfaceId,proto3" json:"main_interface_id,omitempty"`
	MainInterfaceType    string   `protobuf:"bytes,3,opt,name=main_interface_type,json=mainInterfaceType,proto3" json:"main_interface_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceSummaryInfo_KEYS) Reset()         { *m = L2FibMainInterfaceSummaryInfo_KEYS{} }
func (m *L2FibMainInterfaceSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceSummaryInfo_KEYS) ProtoMessage()    {}
func (*L2FibMainInterfaceSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_883cfd7e8b2656b4, []int{0}
}

func (m *L2FibMainInterfaceSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS.Merge(m, src)
}
func (m *L2FibMainInterfaceSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS.Size(m)
}
func (m *L2FibMainInterfaceSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceSummaryInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMainInterfaceSummaryInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMainInterfaceSummaryInfo_KEYS) GetMainInterfaceId() uint32 {
	if m != nil {
		return m.MainInterfaceId
	}
	return 0
}

func (m *L2FibMainInterfaceSummaryInfo_KEYS) GetMainInterfaceType() string {
	if m != nil {
		return m.MainInterfaceType
	}
	return ""
}

type L2FibMainInterfaceSummaryInfo struct {
	MainInterface        string   `protobuf:"bytes,50,opt,name=main_interface,json=mainInterface,proto3" json:"main_interface,omitempty"`
	MainIfType           uint32   `protobuf:"varint,51,opt,name=main_if_type,json=mainIfType,proto3" json:"main_if_type,omitempty"`
	VirtualIfName        string   `protobuf:"bytes,52,opt,name=virtual_if_name,json=virtualIfName,proto3" json:"virtual_if_name,omitempty"`
	InstanceId           []uint32 `protobuf:"varint,53,rep,packed,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstancesState       []string `protobuf:"bytes,54,rep,name=instances_state,json=instancesState,proto3" json:"instances_state,omitempty"`
	InstancesProvisioned []bool   `protobuf:"varint,55,rep,packed,name=instances_provisioned,json=instancesProvisioned,proto3" json:"instances_provisioned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMainInterfaceSummaryInfo) Reset()         { *m = L2FibMainInterfaceSummaryInfo{} }
func (m *L2FibMainInterfaceSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMainInterfaceSummaryInfo) ProtoMessage()    {}
func (*L2FibMainInterfaceSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_883cfd7e8b2656b4, []int{1}
}

func (m *L2FibMainInterfaceSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo.Unmarshal(m, b)
}
func (m *L2FibMainInterfaceSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMainInterfaceSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMainInterfaceSummaryInfo.Merge(m, src)
}
func (m *L2FibMainInterfaceSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMainInterfaceSummaryInfo.Size(m)
}
func (m *L2FibMainInterfaceSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMainInterfaceSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMainInterfaceSummaryInfo proto.InternalMessageInfo

func (m *L2FibMainInterfaceSummaryInfo) GetMainInterface() string {
	if m != nil {
		return m.MainInterface
	}
	return ""
}

func (m *L2FibMainInterfaceSummaryInfo) GetMainIfType() uint32 {
	if m != nil {
		return m.MainIfType
	}
	return 0
}

func (m *L2FibMainInterfaceSummaryInfo) GetVirtualIfName() string {
	if m != nil {
		return m.VirtualIfName
	}
	return ""
}

func (m *L2FibMainInterfaceSummaryInfo) GetInstanceId() []uint32 {
	if m != nil {
		return m.InstanceId
	}
	return nil
}

func (m *L2FibMainInterfaceSummaryInfo) GetInstancesState() []string {
	if m != nil {
		return m.InstancesState
	}
	return nil
}

func (m *L2FibMainInterfaceSummaryInfo) GetInstancesProvisioned() []bool {
	if m != nil {
		return m.InstancesProvisioned
	}
	return nil
}

func init() {
	proto.RegisterType((*L2FibMainInterfaceSummaryInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_hardware_ingress_info.l2fib_main_interface_summary_info_KEYS")
	proto.RegisterType((*L2FibMainInterfaceSummaryInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_main_interfaces.l2fib_main_interface.l2fib_main_interface_hardware_ingress_info.l2fib_main_interface_summary_info")
}

func init() {
	proto.RegisterFile("l2fib_main_interface_summary_info.proto", fileDescriptor_883cfd7e8b2656b4)
}

var fileDescriptor_883cfd7e8b2656b4 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6a, 0x32, 0x31,
	0x14, 0xc5, 0x19, 0x07, 0xfc, 0x3e, 0xd3, 0xaa, 0x98, 0xb6, 0x74, 0x76, 0x9d, 0x0a, 0xd5, 0xa1,
	0x8b, 0x59, 0x68, 0xff, 0x3c, 0x41, 0x17, 0x43, 0xa1, 0x94, 0xb1, 0x9b, 0xae, 0x42, 0x34, 0x89,
	0xbd, 0xe0, 0x24, 0x43, 0x12, 0xb5, 0x3e, 0x42, 0xd7, 0x5d, 0xf7, 0x5d, 0x4b, 0x32, 0x55, 0x51,
	0x04, 0x37, 0xc3, 0xdc, 0xdf, 0x39, 0x39, 0xf7, 0x04, 0x82, 0xfa, 0xb3, 0x81, 0x80, 0x31, 0x29,
	0x28, 0x48, 0x02, 0xd2, 0x72, 0x2d, 0xe8, 0x84, 0x13, 0x33, 0x2f, 0x0a, 0xaa, 0x57, 0x04, 0xa4,
	0x50, 0x69, 0xa9, 0x95, 0x55, 0xf8, 0x2b, 0x98, 0x80, 0x99, 0x28, 0x02, 0xca, 0x90, 0x4f, 0x4d,
	0x66, 0x83, 0x45, 0x29, 0x89, 0x2a, 0xb9, 0x4e, 0xab, 0x5f, 0xa1, 0xf4, 0x92, 0x6a, 0x06, 0x72,
	0x9a, 0x4a, 0xc5, 0xb8, 0xf1, 0xdf, 0xf4, 0x50, 0xba, 0x39, 0x48, 0x0f, 0x42, 0xf2, 0x41, 0x35,
	0x5b, 0x52, 0xcd, 0x09, 0xc8, 0xa9, 0xe6, 0xc6, 0xf8, 0x46, 0xdd, 0x9f, 0x00, 0xf5, 0x8e, 0xf6,
	0x26, 0xcf, 0x4f, 0xef, 0x23, 0x7c, 0x89, 0xfe, 0xb9, 0x22, 0x04, 0x58, 0x14, 0xc4, 0x41, 0xd2,
	0xc8, 0xeb, 0x6e, 0xcc, 0x18, 0xbe, 0x45, 0x9d, 0xbd, 0xc3, 0xc0, 0xa2, 0x5a, 0x1c, 0x24, 0xcd,
	0xbc, 0xed, 0x84, 0x6c, 0xcd, 0x33, 0x86, 0x53, 0x74, 0xb6, 0xe7, 0xb5, 0xab, 0x92, 0x47, 0xa1,
	0x0f, 0xec, 0xec, 0xb8, 0xdf, 0x56, 0x25, 0xef, 0x7e, 0xd7, 0xd0, 0xf5, 0xd1, 0x7e, 0xf8, 0x06,
	0xb5, 0x76, 0xe5, 0x68, 0xe0, 0x03, 0x9b, 0x3b, 0x81, 0x38, 0x46, 0xa7, 0x95, 0x4d, 0x54, 0x5b,
	0x87, 0xbe, 0x23, 0xf2, 0x26, 0xe1, 0xd6, 0xe1, 0x1e, 0x6a, 0x2f, 0x40, 0xdb, 0x39, 0x9d, 0x39,
	0x93, 0xa4, 0x05, 0x8f, 0xee, 0xaa, 0xa4, 0x3f, 0x9c, 0x89, 0x17, 0x5a, 0x70, 0x7c, 0x85, 0x4e,
	0x40, 0x1a, 0x4b, 0x65, 0x75, 0xd9, 0xfb, 0x38, 0x74, 0x41, 0x6b, 0x94, 0x31, 0xdc, 0x47, 0xed,
	0xf5, 0x64, 0x88, 0xb1, 0xd4, 0xf2, 0xe8, 0x21, 0x0e, 0x93, 0x46, 0xde, 0xda, 0xe0, 0x91, 0xa3,
	0x78, 0x88, 0x2e, 0xb6, 0xc6, 0x52, 0xab, 0x05, 0x18, 0x50, 0x92, 0xb3, 0xe8, 0x31, 0x0e, 0x93,
	0xff, 0xf9, 0xf9, 0x46, 0x7c, 0xdd, 0x6a, 0xe3, 0xba, 0x7f, 0x48, 0xc3, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x44, 0x68, 0x28, 0x2b, 0x73, 0x02, 0x00, 0x00,
}
