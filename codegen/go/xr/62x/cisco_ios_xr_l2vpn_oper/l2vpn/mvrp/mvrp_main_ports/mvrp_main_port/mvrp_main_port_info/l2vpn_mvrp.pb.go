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
// source: l2vpn_mvrp.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_mvrp_mvrp_main_ports_mvrp_main_port_mvrp_main_port_info

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

type L2VpnMvrp_KEYS struct {
	MainPortInterfaceName string   `protobuf:"bytes,1,opt,name=main_port_interface_name,json=mainPortInterfaceName,proto3" json:"main_port_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *L2VpnMvrp_KEYS) Reset()         { *m = L2VpnMvrp_KEYS{} }
func (m *L2VpnMvrp_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMvrp_KEYS) ProtoMessage()    {}
func (*L2VpnMvrp_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7db156249ddcd0, []int{0}
}

func (m *L2VpnMvrp_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMvrp_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMvrp_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMvrp_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMvrp_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMvrp_KEYS.Merge(m, src)
}
func (m *L2VpnMvrp_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMvrp_KEYS.Size(m)
}
func (m *L2VpnMvrp_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMvrp_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMvrp_KEYS proto.InternalMessageInfo

func (m *L2VpnMvrp_KEYS) GetMainPortInterfaceName() string {
	if m != nil {
		return m.MainPortInterfaceName
	}
	return ""
}

type L2VpnEfpRange struct {
	Lower                uint32   `protobuf:"varint,1,opt,name=lower,proto3" json:"lower,omitempty"`
	Upper                uint32   `protobuf:"varint,2,opt,name=upper,proto3" json:"upper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEfpRange) Reset()         { *m = L2VpnEfpRange{} }
func (m *L2VpnEfpRange) String() string { return proto.CompactTextString(m) }
func (*L2VpnEfpRange) ProtoMessage()    {}
func (*L2VpnEfpRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7db156249ddcd0, []int{1}
}

func (m *L2VpnEfpRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEfpRange.Unmarshal(m, b)
}
func (m *L2VpnEfpRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEfpRange.Marshal(b, m, deterministic)
}
func (m *L2VpnEfpRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEfpRange.Merge(m, src)
}
func (m *L2VpnEfpRange) XXX_Size() int {
	return xxx_messageInfo_L2VpnEfpRange.Size(m)
}
func (m *L2VpnEfpRange) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEfpRange.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEfpRange proto.InternalMessageInfo

func (m *L2VpnEfpRange) GetLower() uint32 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *L2VpnEfpRange) GetUpper() uint32 {
	if m != nil {
		return m.Upper
	}
	return 0
}

type L2VpnMvrpBp struct {
	BridgePortInterfaceName string           `protobuf:"bytes,1,opt,name=bridge_port_interface_name,json=bridgePortInterfaceName,proto3" json:"bridge_port_interface_name,omitempty"`
	BridgePortXconnectId    uint32           `protobuf:"varint,2,opt,name=bridge_port_xconnect_id,json=bridgePortXconnectId,proto3" json:"bridge_port_xconnect_id,omitempty"`
	MvrpSequenceNumber      uint32           `protobuf:"varint,3,opt,name=mvrp_sequence_number,json=mvrpSequenceNumber,proto3" json:"mvrp_sequence_number,omitempty"`
	VlanRange               []*L2VpnEfpRange `protobuf:"bytes,4,rep,name=vlan_range,json=vlanRange,proto3" json:"vlan_range,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *L2VpnMvrpBp) Reset()         { *m = L2VpnMvrpBp{} }
func (m *L2VpnMvrpBp) String() string { return proto.CompactTextString(m) }
func (*L2VpnMvrpBp) ProtoMessage()    {}
func (*L2VpnMvrpBp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7db156249ddcd0, []int{2}
}

func (m *L2VpnMvrpBp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMvrpBp.Unmarshal(m, b)
}
func (m *L2VpnMvrpBp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMvrpBp.Marshal(b, m, deterministic)
}
func (m *L2VpnMvrpBp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMvrpBp.Merge(m, src)
}
func (m *L2VpnMvrpBp) XXX_Size() int {
	return xxx_messageInfo_L2VpnMvrpBp.Size(m)
}
func (m *L2VpnMvrpBp) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMvrpBp.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMvrpBp proto.InternalMessageInfo

func (m *L2VpnMvrpBp) GetBridgePortInterfaceName() string {
	if m != nil {
		return m.BridgePortInterfaceName
	}
	return ""
}

func (m *L2VpnMvrpBp) GetBridgePortXconnectId() uint32 {
	if m != nil {
		return m.BridgePortXconnectId
	}
	return 0
}

func (m *L2VpnMvrpBp) GetMvrpSequenceNumber() uint32 {
	if m != nil {
		return m.MvrpSequenceNumber
	}
	return 0
}

func (m *L2VpnMvrpBp) GetVlanRange() []*L2VpnEfpRange {
	if m != nil {
		return m.VlanRange
	}
	return nil
}

type L2VpnMvrp struct {
	MainPortInterfaceName string       `protobuf:"bytes,50,opt,name=main_port_interface_name,json=mainPortInterfaceName,proto3" json:"main_port_interface_name,omitempty"`
	IsTrunk               bool         `protobuf:"varint,51,opt,name=is_trunk,json=isTrunk,proto3" json:"is_trunk,omitempty"`
	IsDefaultEncap        bool         `protobuf:"varint,52,opt,name=is_default_encap,json=isDefaultEncap,proto3" json:"is_default_encap,omitempty"`
	DefaultBridgePort     *L2VpnMvrpBp `protobuf:"bytes,53,opt,name=default_bridge_port,json=defaultBridgePort,proto3" json:"default_bridge_port,omitempty"`
	NumberOfBridgePorts   uint32       `protobuf:"varint,54,opt,name=number_of_bridge_ports,json=numberOfBridgePorts,proto3" json:"number_of_bridge_ports,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}     `json:"-"`
	XXX_unrecognized      []byte       `json:"-"`
	XXX_sizecache         int32        `json:"-"`
}

func (m *L2VpnMvrp) Reset()         { *m = L2VpnMvrp{} }
func (m *L2VpnMvrp) String() string { return proto.CompactTextString(m) }
func (*L2VpnMvrp) ProtoMessage()    {}
func (*L2VpnMvrp) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e7db156249ddcd0, []int{3}
}

func (m *L2VpnMvrp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMvrp.Unmarshal(m, b)
}
func (m *L2VpnMvrp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMvrp.Marshal(b, m, deterministic)
}
func (m *L2VpnMvrp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMvrp.Merge(m, src)
}
func (m *L2VpnMvrp) XXX_Size() int {
	return xxx_messageInfo_L2VpnMvrp.Size(m)
}
func (m *L2VpnMvrp) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMvrp.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMvrp proto.InternalMessageInfo

func (m *L2VpnMvrp) GetMainPortInterfaceName() string {
	if m != nil {
		return m.MainPortInterfaceName
	}
	return ""
}

func (m *L2VpnMvrp) GetIsTrunk() bool {
	if m != nil {
		return m.IsTrunk
	}
	return false
}

func (m *L2VpnMvrp) GetIsDefaultEncap() bool {
	if m != nil {
		return m.IsDefaultEncap
	}
	return false
}

func (m *L2VpnMvrp) GetDefaultBridgePort() *L2VpnMvrpBp {
	if m != nil {
		return m.DefaultBridgePort
	}
	return nil
}

func (m *L2VpnMvrp) GetNumberOfBridgePorts() uint32 {
	if m != nil {
		return m.NumberOfBridgePorts
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnMvrp_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mvrp.mvrp_main_ports.mvrp_main_port.mvrp_main_port_info.l2vpn_mvrp_KEYS")
	proto.RegisterType((*L2VpnEfpRange)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mvrp.mvrp_main_ports.mvrp_main_port.mvrp_main_port_info.l2vpn_efp_range")
	proto.RegisterType((*L2VpnMvrpBp)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mvrp.mvrp_main_ports.mvrp_main_port.mvrp_main_port_info.l2vpn_mvrp_bp")
	proto.RegisterType((*L2VpnMvrp)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mvrp.mvrp_main_ports.mvrp_main_port.mvrp_main_port_info.l2vpn_mvrp")
}

func init() { proto.RegisterFile("l2vpn_mvrp.proto", fileDescriptor_5e7db156249ddcd0) }

var fileDescriptor_5e7db156249ddcd0 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0x56, 0xb6, 0xfc, 0xb4, 0x53, 0x15, 0x8a, 0xbb, 0x50, 0xc3, 0x69, 0x95, 0x53, 0x4e, 0x11,
	0xca, 0x52, 0x38, 0x20, 0x2e, 0x88, 0x1e, 0x0a, 0x52, 0x41, 0x29, 0x48, 0x70, 0xb2, 0x9c, 0x64,
	0x52, 0x59, 0x24, 0xb6, 0xb1, 0x93, 0xa5, 0x2f, 0xc0, 0x1b, 0xf0, 0x3a, 0xbc, 0x06, 0xcf, 0x83,
	0x6c, 0x67, 0x49, 0xb4, 0x95, 0xf6, 0xb6, 0x97, 0xc8, 0x33, 0xdf, 0x7c, 0xf9, 0xbe, 0xcc, 0x17,
	0xc3, 0x71, 0x93, 0xad, 0xb4, 0x64, 0xed, 0xca, 0xe8, 0x54, 0x1b, 0xd5, 0x29, 0xf2, 0xa5, 0x14,
	0xb6, 0x54, 0x4c, 0x28, 0xcb, 0x6e, 0x0c, 0x0b, 0xb0, 0xd2, 0x68, 0x52, 0x7f, 0x4c, 0xfd, 0xa4,
	0x7b, 0xb0, 0x96, 0x0b, 0xc9, 0xb4, 0x32, 0x9d, 0xdd, 0xa8, 0x37, 0x4a, 0x26, 0x64, 0xad, 0xe2,
	0xf7, 0xf0, 0x70, 0x94, 0x62, 0x1f, 0xce, 0xbf, 0x5d, 0x91, 0x57, 0x40, 0xa7, 0x43, 0x1d, 0x9a,
	0x9a, 0x97, 0xc8, 0x24, 0x6f, 0x91, 0x46, 0x8b, 0x28, 0x39, 0xc8, 0x1f, 0x3b, 0xfc, 0x93, 0x32,
	0xdd, 0xc5, 0x1a, 0xbd, 0xe4, 0x2d, 0xc6, 0x6f, 0xd6, 0xef, 0xc2, 0x5a, 0x33, 0xc3, 0xe5, 0x35,
	0x92, 0x39, 0xdc, 0x6d, 0xd4, 0x4f, 0x34, 0x9e, 0x78, 0x94, 0x87, 0xc2, 0x75, 0x7b, 0xad, 0xd1,
	0xd0, 0x59, 0xe8, 0xfa, 0x22, 0xfe, 0x33, 0x83, 0xa3, 0x89, 0x97, 0x42, 0x93, 0xd7, 0xf0, 0xac,
	0x30, 0xa2, 0xba, 0xc6, 0x2d, 0x5e, 0x4e, 0xc3, 0xc4, 0x2d, 0x37, 0xe4, 0x0c, 0x4e, 0xa7, 0xe4,
	0x9b, 0x52, 0x49, 0x89, 0x65, 0xc7, 0x44, 0x35, 0xc8, 0xce, 0x47, 0xe6, 0xd7, 0x01, 0xbc, 0xa8,
	0xc8, 0x73, 0x98, 0x7b, 0x79, 0x8b, 0x3f, 0x7a, 0x94, 0x4e, 0xab, 0x6f, 0x0b, 0x34, 0x74, 0xcf,
	0x73, 0x88, 0xc3, 0xae, 0x06, 0xe8, 0xd2, 0x23, 0xe4, 0x57, 0x04, 0xb0, 0x6a, 0xb8, 0x0c, 0x9f,
	0x4c, 0xef, 0x2c, 0xf6, 0x92, 0xc3, 0xac, 0x4e, 0x77, 0x92, 0x57, 0xba, 0xb1, 0xe0, 0xfc, 0xc0,
	0x29, 0xe7, 0xee, 0x18, 0xff, 0x9d, 0x01, 0x8c, 0xfb, 0xdb, 0x1a, 0x63, 0xb6, 0x25, 0x46, 0xf2,
	0x14, 0xf6, 0x85, 0x65, 0x9d, 0xe9, 0xe5, 0x77, 0xba, 0x5c, 0x44, 0xc9, 0x7e, 0x7e, 0x5f, 0xd8,
	0xcf, 0xae, 0x24, 0x09, 0x1c, 0x0b, 0xcb, 0x2a, 0xac, 0x79, 0xdf, 0x74, 0x0c, 0x65, 0xc9, 0x35,
	0x7d, 0xe1, 0x47, 0x1e, 0x08, 0xfb, 0x2e, 0xb4, 0xcf, 0x5d, 0x97, 0xfc, 0x8e, 0xe0, 0x64, 0x3d,
	0x37, 0x89, 0x81, 0x9e, 0x2d, 0xa2, 0xe4, 0x30, 0xab, 0x76, 0xba, 0x9d, 0xe1, 0xf7, 0xc9, 0x1f,
	0x0d, 0x06, 0xde, 0xfe, 0xcf, 0x99, 0x2c, 0xe1, 0x49, 0xc8, 0x93, 0xa9, 0x7a, 0xea, 0xcb, 0xd2,
	0x97, 0x3e, 0xdf, 0x93, 0x80, 0x7e, 0xac, 0x47, 0x8e, 0x2d, 0xee, 0xf9, 0x1b, 0xb8, 0xfc, 0x17,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0x32, 0xa0, 0x20, 0x95, 0x03, 0x00, 0x00,
}
