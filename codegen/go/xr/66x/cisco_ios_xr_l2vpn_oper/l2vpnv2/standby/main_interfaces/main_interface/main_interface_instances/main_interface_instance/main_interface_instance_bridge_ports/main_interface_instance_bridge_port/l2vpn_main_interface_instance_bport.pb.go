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

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_main_interfaces_main_interface_main_interface_instances_main_interface_instance_main_interface_instance_bridge_ports_main_interface_instance_bridge_port

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
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Instance             uint32   `protobuf:"varint,2,opt,name=instance,proto3" json:"instance,omitempty"`
	BridgePort           string   `protobuf:"bytes,3,opt,name=bridge_port,json=bridgePort,proto3" json:"bridge_port,omitempty"`
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
	proto.RegisterType((*L2VpnMainInterfaceInstanceBport_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_bridge_ports.main_interface_instance_bridge_port.l2vpn_main_interface_instance_bport_KEYS")
	proto.RegisterType((*L2VpnMainInterfaceInstanceBport)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_bridge_ports.main_interface_instance_bridge_port.l2vpn_main_interface_instance_bport")
}

func init() {
	proto.RegisterFile("l2vpn_main_interface_instance_bport.proto", fileDescriptor_91f83dea64adeed0)
}

var fileDescriptor_91f83dea64adeed0 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0x31, 0x2a, 0x2b,
	0xc8, 0x8b, 0xcf, 0x4d, 0xcc, 0xcc, 0x8b, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e,
	0x8d, 0xcf, 0xcc, 0x2b, 0x2e, 0x49, 0xcc, 0x4b, 0x4e, 0x8d, 0x4f, 0x2a, 0xc8, 0x2f, 0x2a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x3a, 0xc0, 0x98, 0x9c, 0x59, 0x9c, 0x9c, 0x1f, 0x9f, 0x99,
	0x5f, 0x1c, 0x5f, 0x51, 0x14, 0x0f, 0xd1, 0x98, 0x5f, 0x90, 0x5a, 0xa4, 0x07, 0x66, 0x96, 0x19,
	0xe9, 0x81, 0x34, 0xa5, 0x24, 0x55, 0xea, 0xa1, 0x9a, 0x56, 0x8c, 0xc6, 0xd7, 0xc3, 0x61, 0x59,
	0x31, 0x2e, 0x09, 0x3d, 0x9c, 0xae, 0x2b, 0xca, 0x4c, 0x49, 0x4f, 0x8d, 0x07, 0xb9, 0xb1, 0x98,
	0x18, 0x45, 0x4a, 0x7d, 0x8c, 0x5c, 0x1a, 0x44, 0x78, 0x38, 0xde, 0xdb, 0x35, 0x32, 0x58, 0x48,
	0x95, 0x8b, 0x0f, 0xa1, 0x20, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88,
	0x17, 0x2e, 0xea, 0x97, 0x98, 0x9b, 0x2a, 0x24, 0xc5, 0xc5, 0x01, 0xd3, 0x2d, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x1b, 0x04, 0xe7, 0x0b, 0xc9, 0x73, 0x71, 0x23, 0x59, 0x2f, 0xc1, 0x0c, 0xd6, 0xcf,
	0x05, 0x11, 0x0a, 0x00, 0x39, 0x28, 0x87, 0x4b, 0x99, 0x08, 0xf7, 0x08, 0xa9, 0x70, 0xf1, 0x21,
	0x99, 0x13, 0x5f, 0x51, 0x24, 0x61, 0x04, 0x36, 0x8a, 0x07, 0x61, 0x54, 0x44, 0x11, 0xc8, 0x36,
	0xb8, 0xbe, 0xcc, 0x14, 0x09, 0x63, 0xb0, 0x63, 0xb8, 0x60, 0x42, 0x9e, 0x29, 0x49, 0x6c, 0xe0,
	0x88, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xe1, 0x5d, 0x33, 0xf5, 0x01, 0x00, 0x00,
}
