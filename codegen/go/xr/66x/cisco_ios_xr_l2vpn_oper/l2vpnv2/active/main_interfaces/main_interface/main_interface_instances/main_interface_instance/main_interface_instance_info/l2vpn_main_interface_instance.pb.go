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
// source: l2vpn_main_interface_instance.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_main_interfaces_main_interface_main_interface_instances_main_interface_instance_main_interface_instance_info

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

type L2VpnMainInterfaceInstance_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Instance             uint32   `protobuf:"varint,2,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMainInterfaceInstance_KEYS) Reset()         { *m = L2VpnMainInterfaceInstance_KEYS{} }
func (m *L2VpnMainInterfaceInstance_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMainInterfaceInstance_KEYS) ProtoMessage()    {}
func (*L2VpnMainInterfaceInstance_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7d7294270a3af3, []int{0}
}

func (m *L2VpnMainInterfaceInstance_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMainInterfaceInstance_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMainInterfaceInstance_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS.Merge(m, src)
}
func (m *L2VpnMainInterfaceInstance_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS.Size(m)
}
func (m *L2VpnMainInterfaceInstance_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMainInterfaceInstance_KEYS proto.InternalMessageInfo

func (m *L2VpnMainInterfaceInstance_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2VpnMainInterfaceInstance_KEYS) GetInstance() uint32 {
	if m != nil {
		return m.Instance
	}
	return 0
}

type L2VpnMainInterfaceInstance struct {
	ConfiguredInstance   uint32   `protobuf:"varint,50,opt,name=configured_instance,json=configuredInstance,proto3" json:"configured_instance,omitempty"`
	FlushCount           uint32   `protobuf:"varint,51,opt,name=flush_count,json=flushCount,proto3" json:"flush_count,omitempty"`
	InterfaceCount       uint32   `protobuf:"varint,52,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	InstanceFlags        uint32   `protobuf:"varint,53,opt,name=instance_flags,json=instanceFlags,proto3" json:"instance_flags,omitempty"`
	InstanceId           uint32   `protobuf:"varint,54,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	InstanceState        string   `protobuf:"bytes,55,opt,name=instance_state,json=instanceState,proto3" json:"instance_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMainInterfaceInstance) Reset()         { *m = L2VpnMainInterfaceInstance{} }
func (m *L2VpnMainInterfaceInstance) String() string { return proto.CompactTextString(m) }
func (*L2VpnMainInterfaceInstance) ProtoMessage()    {}
func (*L2VpnMainInterfaceInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd7d7294270a3af3, []int{1}
}

func (m *L2VpnMainInterfaceInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMainInterfaceInstance.Unmarshal(m, b)
}
func (m *L2VpnMainInterfaceInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMainInterfaceInstance.Marshal(b, m, deterministic)
}
func (m *L2VpnMainInterfaceInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMainInterfaceInstance.Merge(m, src)
}
func (m *L2VpnMainInterfaceInstance) XXX_Size() int {
	return xxx_messageInfo_L2VpnMainInterfaceInstance.Size(m)
}
func (m *L2VpnMainInterfaceInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMainInterfaceInstance.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMainInterfaceInstance proto.InternalMessageInfo

func (m *L2VpnMainInterfaceInstance) GetConfiguredInstance() uint32 {
	if m != nil {
		return m.ConfiguredInstance
	}
	return 0
}

func (m *L2VpnMainInterfaceInstance) GetFlushCount() uint32 {
	if m != nil {
		return m.FlushCount
	}
	return 0
}

func (m *L2VpnMainInterfaceInstance) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *L2VpnMainInterfaceInstance) GetInstanceFlags() uint32 {
	if m != nil {
		return m.InstanceFlags
	}
	return 0
}

func (m *L2VpnMainInterfaceInstance) GetInstanceId() uint32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *L2VpnMainInterfaceInstance) GetInstanceState() string {
	if m != nil {
		return m.InstanceState
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnMainInterfaceInstance_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_info.l2vpn_main_interface_instance_KEYS")
	proto.RegisterType((*L2VpnMainInterfaceInstance)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.main_interfaces.main_interface.main_interface_instances.main_interface_instance.main_interface_instance_info.l2vpn_main_interface_instance")
}

func init() {
	proto.RegisterFile("l2vpn_main_interface_instance.proto", fileDescriptor_dd7d7294270a3af3)
}

var fileDescriptor_dd7d7294270a3af3 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xe9, 0x1e, 0x44, 0x23, 0xab, 0x10, 0x2f, 0x41, 0x10, 0x97, 0x8a, 0xb8, 0xa7, 0x08,
	0x5d, 0xff, 0x3c, 0x80, 0x28, 0x2c, 0x82, 0x87, 0xdd, 0x93, 0xa7, 0x10, 0xd3, 0xa4, 0x06, 0xda,
	0xa4, 0x24, 0x69, 0xf1, 0x15, 0x3c, 0xfb, 0xc2, 0xd2, 0xc4, 0x34, 0xac, 0xd0, 0xbd, 0x75, 0x7e,
	0xf3, 0xcd, 0x7c, 0xfd, 0x86, 0x80, 0xab, 0xba, 0xe8, 0x5b, 0x45, 0x1a, 0x2a, 0x15, 0x91, 0xca,
	0x71, 0x23, 0x28, 0xe3, 0x44, 0x2a, 0xeb, 0xa8, 0x62, 0x1c, 0xb7, 0x46, 0x3b, 0x0d, 0x7f, 0x32,
	0x26, 0x2d, 0xd3, 0x44, 0x6a, 0x4b, 0xbe, 0x0c, 0x09, 0x23, 0xba, 0xe5, 0x06, 0xfb, 0xcf, 0xbe,
	0xc0, 0x94, 0x39, 0xd9, 0x73, 0xbc, 0xbb, 0xc6, 0xfe, 0xab, 0xf1, 0x84, 0x8b, 0x9d, 0x6a, 0x4c,
	0x71, 0x22, 0x95, 0xd0, 0x79, 0x05, 0xf2, 0xbd, 0x3f, 0x4f, 0x5e, 0x9f, 0xdf, 0xb7, 0xf0, 0x1a,
	0x9c, 0xa4, 0x96, 0xa2, 0x0d, 0x47, 0xd9, 0x22, 0x5b, 0x1e, 0x6d, 0xe6, 0x23, 0x7d, 0xa3, 0x0d,
	0x87, 0xe7, 0xe0, 0x30, 0xce, 0xa1, 0xd9, 0x22, 0x5b, 0xce, 0x37, 0x63, 0x9d, 0x7f, 0xcf, 0xc0,
	0xc5, 0x5e, 0x27, 0x78, 0x0b, 0xce, 0x98, 0x56, 0x42, 0x56, 0x9d, 0xe1, 0xe5, 0x88, 0x51, 0xe1,
	0x17, 0xc1, 0xd4, 0x5a, 0xc7, 0x81, 0x4b, 0x70, 0x2c, 0xea, 0xce, 0x7e, 0x12, 0xa6, 0x3b, 0xe5,
	0xd0, 0xca, 0x0b, 0x81, 0x47, 0x4f, 0x03, 0x81, 0x37, 0xe0, 0x34, 0xf9, 0x04, 0xd1, 0x9d, 0x17,
	0xa5, 0x34, 0x41, 0xe8, 0xf3, 0xfd, 0x05, 0x16, 0x35, 0xad, 0x2c, 0xba, 0xf7, 0xba, 0x79, 0xa4,
	0x2f, 0x03, 0x1c, 0x0c, 0xd3, 0xf5, 0x4a, 0xf4, 0x10, 0x0c, 0x23, 0x5a, 0x97, 0x3b, 0x7b, 0xac,
	0xa3, 0x8e, 0xa3, 0xc7, 0x78, 0xa7, 0x40, 0xb7, 0x03, 0xfc, 0x38, 0xf0, 0x2f, 0x62, 0xf5, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0x06, 0xd5, 0x19, 0x38, 0x02, 0x00, 0x00,
}
