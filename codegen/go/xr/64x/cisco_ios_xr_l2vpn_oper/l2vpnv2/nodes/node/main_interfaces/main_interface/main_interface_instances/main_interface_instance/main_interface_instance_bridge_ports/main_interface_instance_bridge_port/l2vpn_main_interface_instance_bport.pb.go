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
// source: l2vpn_main_interface_instance_bport.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_main_interfaces_main_interface_main_interface_instances_main_interface_instance_main_interface_instance_bridge_ports_main_interface_instance_bridge_port

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

type L2VpnMainInterfaceInstanceBport_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Instance             uint32   `protobuf:"varint,3,opt,name=instance,proto3" json:"instance,omitempty"`
	BridgePort           string   `protobuf:"bytes,4,opt,name=bridge_port,json=bridgePort,proto3" json:"bridge_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMainInterfaceInstanceBport_KEYS) Reset()         { *m = L2VpnMainInterfaceInstanceBport_KEYS{} }
func (m *L2VpnMainInterfaceInstanceBport_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMainInterfaceInstanceBport_KEYS) ProtoMessage()    {}
func (*L2VpnMainInterfaceInstanceBport_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f83dea64adeed0, []int{0}
}

func (m *L2VpnMainInterfaceInstanceBport_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMainInterfaceInstanceBport_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMainInterfaceInstanceBport_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS.Merge(m, src)
}
func (m *L2VpnMainInterfaceInstanceBport_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS.Size(m)
}
func (m *L2VpnMainInterfaceInstanceBport_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMainInterfaceInstanceBport_KEYS proto.InternalMessageInfo

func (m *L2VpnMainInterfaceInstanceBport_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnMainInterfaceInstanceBport_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2VpnMainInterfaceInstanceBport_KEYS) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

func (m *L2VpnMainInterfaceInstanceBport_KEYS) GetBridgePort() string {
	if m != nil {
		return m.BridgePort
	}
	return ""
}

type L2VpnMainInterfaceInstanceBport struct {
	BridgePortXr         string   `protobuf:"bytes,50,opt,name=bridge_port_xr,json=bridgePortXr,proto3" json:"bridge_port_xr,omitempty"`
	InstanceId           uint32   `protobuf:"varint,51,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMainInterfaceInstanceBport) Reset()         { *m = L2VpnMainInterfaceInstanceBport{} }
func (m *L2VpnMainInterfaceInstanceBport) String() string { return proto.CompactTextString(m) }
func (*L2VpnMainInterfaceInstanceBport) ProtoMessage()    {}
func (*L2VpnMainInterfaceInstanceBport) Descriptor() ([]byte, []int) {
	return fileDescriptor_91f83dea64adeed0, []int{1}
}

func (m *L2VpnMainInterfaceInstanceBport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport.Unmarshal(m, b)
}
func (m *L2VpnMainInterfaceInstanceBport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport.Marshal(b, m, deterministic)
}
func (m *L2VpnMainInterfaceInstanceBport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMainInterfaceInstanceBport.Merge(m, src)
}
func (m *L2VpnMainInterfaceInstanceBport) XXX_Size() int {
	return xxx_messageInfo_L2VpnMainInterfaceInstanceBport.Size(m)
}
func (m *L2VpnMainInterfaceInstanceBport) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMainInterfaceInstanceBport.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMainInterfaceInstanceBport proto.InternalMessageInfo

func (m *L2VpnMainInterfaceInstanceBport) GetBridgePortXr() string {
	if m != nil {
		return m.BridgePortXr
	}
	return ""
}

func (m *L2VpnMainInterfaceInstanceBport) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnMainInterfaceInstanceBport_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_bridge_ports.main_interface_instance_bridge_port.l2vpn_main_interface_instance_bport_KEYS")
	proto.RegisterType((*L2VpnMainInterfaceInstanceBport)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_bridge_ports.main_interface_instance_bridge_port.l2vpn_main_interface_instance_bport")
}

func init() {
	proto.RegisterFile("l2vpn_main_interface_instance_bport.proto", fileDescriptor_91f83dea64adeed0)
}

var fileDescriptor_91f83dea64adeed0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x89, 0x4a, 0xd5, 0xd1, 0xf6, 0x90, 0x8b, 0xc1, 0x8b, 0xc5, 0x3f, 0xb0, 0x5e, 0x72,
	0xd8, 0x3e, 0x83, 0x87, 0x22, 0x88, 0xd4, 0x8b, 0x9e, 0x86, 0x74, 0x13, 0x65, 0xa0, 0x9b, 0x2c,
	0x93, 0xa5, 0xf4, 0x8d, 0x7c, 0x10, 0x5f, 0x4c, 0x92, 0xd2, 0x6d, 0x15, 0x16, 0xf6, 0x12, 0x32,
	0xdf, 0xe4, 0xfb, 0xe6, 0x37, 0x04, 0x1e, 0x57, 0xe5, 0xba, 0xf1, 0x58, 0x1b, 0xf2, 0x48, 0xbe,
	0x75, 0xfc, 0x69, 0x2a, 0x87, 0xe4, 0x63, 0x6b, 0x7c, 0xe5, 0x70, 0xd9, 0x04, 0x6e, 0x75, 0xc3,
	0xa1, 0x0d, 0xf2, 0x47, 0x54, 0x14, 0xab, 0x80, 0x14, 0x22, 0x6e, 0x18, 0xb7, 0xc6, 0xd0, 0x38,
	0xd6, 0xf9, 0xba, 0x2e, 0xb5, 0x0f, 0xd6, 0xc5, 0x7c, 0xea, 0xbf, 0x81, 0xf1, 0x5f, 0xad, 0x7b,
	0xe6, 0xc5, 0xbe, 0x86, 0xee, 0x05, 0x64, 0xb2, 0x5f, 0x0e, 0x13, 0x66, 0x1c, 0xf2, 0xe8, 0xf6,
	0x5b, 0x40, 0x31, 0x60, 0x67, 0x7c, 0x7e, 0xfa, 0x78, 0x93, 0x57, 0x70, 0x9a, 0xb6, 0x41, 0xb2,
	0x4a, 0x4c, 0x45, 0x71, 0xbe, 0x18, 0xa5, 0x72, 0x6e, 0xe5, 0x03, 0x4c, 0xf6, 0x4e, 0x6f, 0x6a,
	0xa7, 0x8e, 0x72, 0x7f, 0xdc, 0xa9, 0x2f, 0xa6, 0x76, 0xf2, 0x1a, 0xce, 0x76, 0xb1, 0xea, 0x78,
	0x2a, 0x8a, 0xf1, 0xa2, 0xab, 0xe5, 0x0d, 0x5c, 0x1c, 0x70, 0xa9, 0x93, 0xec, 0x87, 0xad, 0xf4,
	0x9a, 0x48, 0x57, 0x70, 0x37, 0x00, 0x54, 0xde, 0xc3, 0xe4, 0x20, 0x07, 0x37, 0xac, 0xca, 0x1c,
	0x75, 0xb9, 0x8f, 0x7a, 0xe7, 0x34, 0xad, 0xf3, 0x91, 0x55, 0xb3, 0x0c, 0x03, 0x3b, 0x69, 0x6e,
	0x97, 0xa3, 0xfc, 0xc9, 0xb3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x12, 0xdb, 0x26, 0x11,
	0x02, 0x00, 0x00,
}
