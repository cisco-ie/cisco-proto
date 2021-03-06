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
// source: l2vpn_mstp_vlan.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_mstp_vlans_mstp_vlan

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

type L2VpnMstpVlan_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	VlanId               uint32   `protobuf:"varint,2,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMstpVlan_KEYS) Reset()         { *m = L2VpnMstpVlan_KEYS{} }
func (m *L2VpnMstpVlan_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstpVlan_KEYS) ProtoMessage()    {}
func (*L2VpnMstpVlan_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cdbda0015a78c49, []int{0}
}

func (m *L2VpnMstpVlan_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstpVlan_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMstpVlan_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstpVlan_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMstpVlan_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstpVlan_KEYS.Merge(m, src)
}
func (m *L2VpnMstpVlan_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstpVlan_KEYS.Size(m)
}
func (m *L2VpnMstpVlan_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstpVlan_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstpVlan_KEYS proto.InternalMessageInfo

func (m *L2VpnMstpVlan_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnMstpVlan_KEYS) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

type L2VpnMstpSubint struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMstpSubint) Reset()         { *m = L2VpnMstpSubint{} }
func (m *L2VpnMstpSubint) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstpSubint) ProtoMessage()    {}
func (*L2VpnMstpSubint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cdbda0015a78c49, []int{1}
}

func (m *L2VpnMstpSubint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstpSubint.Unmarshal(m, b)
}
func (m *L2VpnMstpSubint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstpSubint.Marshal(b, m, deterministic)
}
func (m *L2VpnMstpSubint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstpSubint.Merge(m, src)
}
func (m *L2VpnMstpSubint) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstpSubint.Size(m)
}
func (m *L2VpnMstpSubint) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstpSubint.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstpSubint proto.InternalMessageInfo

func (m *L2VpnMstpSubint) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type L2VpnMstpVlan struct {
	VlanIdXr             uint32             `protobuf:"varint,50,opt,name=vlan_id_xr,json=vlanIdXr,proto3" json:"vlan_id_xr,omitempty"`
	MstiId               uint32             `protobuf:"varint,51,opt,name=msti_id,json=mstiId,proto3" json:"msti_id,omitempty"`
	PortCount            uint32             `protobuf:"varint,52,opt,name=port_count,json=portCount,proto3" json:"port_count,omitempty"`
	SubInterface         []*L2VpnMstpSubint `protobuf:"bytes,53,rep,name=sub_interface,json=subInterface,proto3" json:"sub_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnMstpVlan) Reset()         { *m = L2VpnMstpVlan{} }
func (m *L2VpnMstpVlan) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstpVlan) ProtoMessage()    {}
func (*L2VpnMstpVlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_6cdbda0015a78c49, []int{2}
}

func (m *L2VpnMstpVlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstpVlan.Unmarshal(m, b)
}
func (m *L2VpnMstpVlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstpVlan.Marshal(b, m, deterministic)
}
func (m *L2VpnMstpVlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstpVlan.Merge(m, src)
}
func (m *L2VpnMstpVlan) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstpVlan.Size(m)
}
func (m *L2VpnMstpVlan) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstpVlan.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstpVlan proto.InternalMessageInfo

func (m *L2VpnMstpVlan) GetVlanIdXr() uint32 {
	if m != nil {
		return m.VlanIdXr
	}
	return 0
}

func (m *L2VpnMstpVlan) GetMstiId() uint32 {
	if m != nil {
		return m.MstiId
	}
	return 0
}

func (m *L2VpnMstpVlan) GetPortCount() uint32 {
	if m != nil {
		return m.PortCount
	}
	return 0
}

func (m *L2VpnMstpVlan) GetSubInterface() []*L2VpnMstpSubint {
	if m != nil {
		return m.SubInterface
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnMstpVlan_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.mstp_vlans.mstp_vlan.l2vpn_mstp_vlan_KEYS")
	proto.RegisterType((*L2VpnMstpSubint)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.mstp_vlans.mstp_vlan.l2vpn_mstp_subint")
	proto.RegisterType((*L2VpnMstpVlan)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.mstp_vlans.mstp_vlan.l2vpn_mstp_vlan")
}

func init() { proto.RegisterFile("l2vpn_mstp_vlan.proto", fileDescriptor_6cdbda0015a78c49) }

var fileDescriptor_6cdbda0015a78c49 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x51, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0xa5, 0x0a, 0xd5, 0x1d, 0xad, 0x62, 0x50, 0xec, 0x41, 0xa1, 0x14, 0x84, 0x9e, 0x72, 0xe8,
	0xea, 0xc5, 0x8b, 0x07, 0x11, 0x2c, 0x82, 0x87, 0x7a, 0xd1, 0xd3, 0xd0, 0x36, 0x11, 0x02, 0xdb,
	0x24, 0x24, 0x69, 0xdd, 0xcf, 0xf5, 0x53, 0x24, 0xe9, 0x5a, 0x65, 0x3d, 0x7a, 0x09, 0xf3, 0xde,
	0xcc, 0xbc, 0xf7, 0x92, 0xc0, 0xd9, 0xaa, 0x1c, 0xb5, 0xc4, 0xde, 0x3a, 0x8d, 0xe3, 0xaa, 0x91,
	0x54, 0x1b, 0xe5, 0x14, 0xb9, 0xeb, 0x84, 0xed, 0x14, 0x0a, 0x65, 0x71, 0x6d, 0x70, 0x9a, 0x51,
	0x9a, 0x1b, 0x1a, 0xca, 0xb1, 0xa4, 0x52, 0x31, 0x6e, 0xc3, 0x49, 0xe7, 0x5d, 0xfb, 0x53, 0xe6,
	0x8f, 0x70, 0xba, 0xa5, 0x8c, 0x4f, 0x0f, 0x6f, 0x2f, 0xe4, 0x1c, 0xf6, 0xfc, 0x12, 0x0a, 0x96,
	0x46, 0x59, 0x54, 0x2c, 0xea, 0xd8, 0xc3, 0x8a, 0xf9, 0x46, 0x98, 0x12, 0x2c, 0xdd, 0xc9, 0xa2,
	0x22, 0xa9, 0x63, 0x0f, 0x2b, 0x96, 0xdf, 0xc2, 0xc9, 0x2f, 0x25, 0x3b, 0xb4, 0x42, 0x3a, 0x72,
	0x05, 0x47, 0x42, 0x3a, 0x6e, 0xde, 0x9b, 0x8e, 0xa3, 0x6c, 0x7a, 0xbe, 0x51, 0x4b, 0x66, 0xf6,
	0xb9, 0xe9, 0x79, 0xfe, 0x19, 0xc1, 0xf1, 0x56, 0x0c, 0x72, 0x01, 0xb0, 0x31, 0xc2, 0xb5, 0x49,
	0xcb, 0xe0, 0xb5, 0x3f, 0x79, 0xbd, 0x1a, 0x1f, 0xa3, 0xb7, 0x4e, 0xf8, 0x18, 0xcb, 0x29, 0x86,
	0x87, 0x15, 0x23, 0x97, 0x00, 0x5a, 0x19, 0x87, 0x9d, 0x1a, 0xa4, 0x4b, 0xaf, 0x43, 0x6f, 0xe1,
	0x99, 0x7b, 0x4f, 0x90, 0x0f, 0x48, 0xec, 0xd0, 0xe2, 0x6c, 0x9f, 0xde, 0x64, 0xbb, 0xc5, 0x41,
	0x59, 0xd3, 0x7f, 0x3e, 0x24, 0xfd, 0x73, 0xf7, 0xfa, 0xd0, 0x0e, 0x6d, 0xf5, 0xed, 0xd3, 0xc6,
	0xe1, 0xc3, 0x96, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xa4, 0x97, 0x18, 0xc9, 0x01, 0x00,
	0x00,
}
