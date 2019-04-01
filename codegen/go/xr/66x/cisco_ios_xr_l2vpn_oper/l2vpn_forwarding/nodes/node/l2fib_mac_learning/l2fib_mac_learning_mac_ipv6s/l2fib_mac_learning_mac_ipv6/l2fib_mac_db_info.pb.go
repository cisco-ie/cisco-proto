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
// source: l2fib_mac_db_info.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_forwarding_nodes_node_l2fib_mac_learning_l2fib_mac_learning_mac_ipv6s_l2fib_mac_learning_mac_ipv6

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

type L2FibMacDbInfo_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	InterfaceHandle      string   `protobuf:"bytes,2,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Bdid                 uint32   `protobuf:"varint,4,opt,name=bdid,proto3" json:"bdid,omitempty"`
	MacAddress           string   `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibMacDbInfo_KEYS) Reset()         { *m = L2FibMacDbInfo_KEYS{} }
func (m *L2FibMacDbInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2FibMacDbInfo_KEYS) ProtoMessage()    {}
func (*L2FibMacDbInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_46a420ab8ce9420f, []int{0}
}

func (m *L2FibMacDbInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMacDbInfo_KEYS.Unmarshal(m, b)
}
func (m *L2FibMacDbInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMacDbInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *L2FibMacDbInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMacDbInfo_KEYS.Merge(m, src)
}
func (m *L2FibMacDbInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2FibMacDbInfo_KEYS.Size(m)
}
func (m *L2FibMacDbInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMacDbInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMacDbInfo_KEYS proto.InternalMessageInfo

func (m *L2FibMacDbInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMacDbInfo_KEYS) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
}

func (m *L2FibMacDbInfo_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *L2FibMacDbInfo_KEYS) GetBdid() uint32 {
	if m != nil {
		return m.Bdid
	}
	return 0
}

func (m *L2FibMacDbInfo_KEYS) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

type L2FibIpAddrT struct {
	AddrType             string   `protobuf:"bytes,1,opt,name=addr_type,json=addrType,proto3" json:"addr_type,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2FibIpAddrT) Reset()         { *m = L2FibIpAddrT{} }
func (m *L2FibIpAddrT) String() string { return proto.CompactTextString(m) }
func (*L2FibIpAddrT) ProtoMessage()    {}
func (*L2FibIpAddrT) Descriptor() ([]byte, []int) {
	return fileDescriptor_46a420ab8ce9420f, []int{1}
}

func (m *L2FibIpAddrT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibIpAddrT.Unmarshal(m, b)
}
func (m *L2FibIpAddrT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibIpAddrT.Marshal(b, m, deterministic)
}
func (m *L2FibIpAddrT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibIpAddrT.Merge(m, src)
}
func (m *L2FibIpAddrT) XXX_Size() int {
	return xxx_messageInfo_L2FibIpAddrT.Size(m)
}
func (m *L2FibIpAddrT) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibIpAddrT.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibIpAddrT proto.InternalMessageInfo

func (m *L2FibIpAddrT) GetAddrType() string {
	if m != nil {
		return m.AddrType
	}
	return ""
}

func (m *L2FibIpAddrT) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type L2FibMacDbInfo struct {
	ElementType          uint32        `protobuf:"varint,50,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
	NodeId               string        `protobuf:"bytes,51,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	TopologyId           uint32        `protobuf:"varint,52,opt,name=topology_id,json=topologyId,proto3" json:"topology_id,omitempty"`
	MacAddressXr         string        `protobuf:"bytes,53,opt,name=mac_address_xr,json=macAddressXr,proto3" json:"mac_address_xr,omitempty"`
	InterfaceHandleXr    uint32        `protobuf:"varint,54,opt,name=interface_handle_xr,json=interfaceHandleXr,proto3" json:"interface_handle_xr,omitempty"`
	Xcid                 uint32        `protobuf:"varint,55,opt,name=xcid,proto3" json:"xcid,omitempty"`
	Generation           uint32        `protobuf:"varint,56,opt,name=generation,proto3" json:"generation,omitempty"`
	Addr                 *L2FibIpAddrT `protobuf:"bytes,57,opt,name=addr,proto3" json:"addr,omitempty"`
	Nhaddr               *L2FibIpAddrT `protobuf:"bytes,58,opt,name=nhaddr,proto3" json:"nhaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *L2FibMacDbInfo) Reset()         { *m = L2FibMacDbInfo{} }
func (m *L2FibMacDbInfo) String() string { return proto.CompactTextString(m) }
func (*L2FibMacDbInfo) ProtoMessage()    {}
func (*L2FibMacDbInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_46a420ab8ce9420f, []int{2}
}

func (m *L2FibMacDbInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2FibMacDbInfo.Unmarshal(m, b)
}
func (m *L2FibMacDbInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2FibMacDbInfo.Marshal(b, m, deterministic)
}
func (m *L2FibMacDbInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2FibMacDbInfo.Merge(m, src)
}
func (m *L2FibMacDbInfo) XXX_Size() int {
	return xxx_messageInfo_L2FibMacDbInfo.Size(m)
}
func (m *L2FibMacDbInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2FibMacDbInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2FibMacDbInfo proto.InternalMessageInfo

func (m *L2FibMacDbInfo) GetElementType() uint32 {
	if m != nil {
		return m.ElementType
	}
	return 0
}

func (m *L2FibMacDbInfo) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2FibMacDbInfo) GetTopologyId() uint32 {
	if m != nil {
		return m.TopologyId
	}
	return 0
}

func (m *L2FibMacDbInfo) GetMacAddressXr() string {
	if m != nil {
		return m.MacAddressXr
	}
	return ""
}

func (m *L2FibMacDbInfo) GetInterfaceHandleXr() uint32 {
	if m != nil {
		return m.InterfaceHandleXr
	}
	return 0
}

func (m *L2FibMacDbInfo) GetXcid() uint32 {
	if m != nil {
		return m.Xcid
	}
	return 0
}

func (m *L2FibMacDbInfo) GetGeneration() uint32 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *L2FibMacDbInfo) GetAddr() *L2FibIpAddrT {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *L2FibMacDbInfo) GetNhaddr() *L2FibIpAddrT {
	if m != nil {
		return m.Nhaddr
	}
	return nil
}

func init() {
	proto.RegisterType((*L2FibMacDbInfo_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mac_learning.l2fib_mac_learning_mac_ipv6s.l2fib_mac_learning_mac_ipv6.l2fib_mac_db_info_KEYS")
	proto.RegisterType((*L2FibIpAddrT)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mac_learning.l2fib_mac_learning_mac_ipv6s.l2fib_mac_learning_mac_ipv6.l2fib_ip_addr_t")
	proto.RegisterType((*L2FibMacDbInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn_forwarding.nodes.node.l2fib_mac_learning.l2fib_mac_learning_mac_ipv6s.l2fib_mac_learning_mac_ipv6.l2fib_mac_db_info")
}

func init() { proto.RegisterFile("l2fib_mac_db_info.proto", fileDescriptor_46a420ab8ce9420f) }

var fileDescriptor_46a420ab8ce9420f = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0xe5, 0x12, 0xba, 0xec, 0x64, 0xd9, 0x65, 0x8d, 0x44, 0x2d, 0x21, 0xd1, 0x52, 0x71,
	0x28, 0x97, 0x1c, 0x52, 0x28, 0x1f, 0x07, 0x24, 0x0e, 0x48, 0x54, 0xdc, 0x0a, 0x87, 0x72, 0xb2,
	0x9c, 0xd8, 0x69, 0x2d, 0xa5, 0xb6, 0xe5, 0x44, 0xa5, 0x3d, 0xf1, 0x0a, 0xbc, 0x04, 0x07, 0xc4,
	0xbb, 0xf0, 0x4c, 0xc8, 0x93, 0x94, 0x96, 0x56, 0xe2, 0xca, 0x5e, 0xaa, 0x99, 0xff, 0x7c, 0xf4,
	0x37, 0xff, 0x24, 0xd0, 0x2b, 0xd3, 0x42, 0x67, 0x7c, 0x25, 0x72, 0x2e, 0x33, 0xae, 0x4d, 0x61,
	0x13, 0xe7, 0x6d, 0x6d, 0xe9, 0xd7, 0x5c, 0x57, 0xb9, 0xe5, 0xda, 0x56, 0x7c, 0xe3, 0x79, 0x99,
	0xae, 0x9d, 0xe1, 0xd6, 0x29, 0x9f, 0x34, 0x61, 0x61, 0xfd, 0x17, 0xe1, 0xa5, 0x36, 0x8b, 0xc4,
	0x58, 0xa9, 0x2a, 0xfc, 0x4d, 0xf6, 0xcb, 0x4a, 0x25, 0xbc, 0x09, 0xd5, 0x53, 0x09, 0x13, 0xed,
	0xd6, 0x93, 0xea, 0x5f, 0xc5, 0xe1, 0x4f, 0x02, 0x0f, 0x4e, 0xe0, 0xf8, 0x87, 0x77, 0x9f, 0x3f,
	0xd2, 0x1e, 0x9c, 0x85, 0xbf, 0xe3, 0x5a, 0x32, 0x32, 0x20, 0xa3, 0xf3, 0x59, 0x37, 0xa4, 0x53,
	0x49, 0x9f, 0xc2, 0x3d, 0x6d, 0x6a, 0xe5, 0x0b, 0x91, 0x2b, 0xbe, 0x14, 0x46, 0x96, 0x8a, 0x75,
	0xb0, 0xe3, 0xea, 0x8f, 0xfe, 0x1e, 0x65, 0xca, 0xe0, 0x4c, 0x48, 0xe9, 0x55, 0x55, 0xb1, 0x5b,
	0xd8, 0xb1, 0x4b, 0x29, 0x85, 0x28, 0x93, 0x5a, 0xb2, 0x68, 0x40, 0x46, 0x77, 0x67, 0x18, 0xd3,
	0x3e, 0xc4, 0x81, 0x62, 0x37, 0x71, 0x1b, 0x27, 0x60, 0x25, 0xf2, 0xb7, 0x8d, 0x32, 0x7c, 0x03,
	0x57, 0x0d, 0xac, 0x76, 0xd8, 0xc5, 0x6b, 0xfa, 0x10, 0xce, 0x9b, 0x68, 0xeb, 0x54, 0xcb, 0x79,
	0x27, 0x08, 0x9f, 0xb6, 0x4e, 0xd1, 0x4b, 0xe8, 0x68, 0xd7, 0xb2, 0x75, 0xb4, 0x1b, 0xfe, 0x8a,
	0xe0, 0xfa, 0xe4, 0x5a, 0xfa, 0x18, 0x2e, 0x54, 0xa9, 0x56, 0xca, 0xd4, 0xcd, 0x96, 0x14, 0x91,
	0xe2, 0x56, 0xc3, 0x45, 0x07, 0x5e, 0x8c, 0xff, 0xf2, 0xa2, 0x0f, 0x71, 0x6d, 0x9d, 0x2d, 0xed,
	0x62, 0x1b, 0x8a, 0xcf, 0x70, 0x14, 0x76, 0xd2, 0x54, 0xd2, 0x27, 0x70, 0x79, 0x70, 0x13, 0xdf,
	0x78, 0xf6, 0x1c, 0x17, 0x5c, 0xec, 0xcf, 0x9a, 0x7b, 0x9a, 0xc0, 0xfd, 0x63, 0x4b, 0x43, 0xeb,
	0x04, 0xd7, 0x5d, 0x1f, 0xb9, 0x3a, 0xf7, 0xc1, 0xbd, 0x4d, 0xae, 0x25, 0x7b, 0xd1, 0xb8, 0x17,
	0x62, 0xfa, 0x08, 0x60, 0xa1, 0x8c, 0xf2, 0xa2, 0xd6, 0xd6, 0xb0, 0x97, 0x0d, 0xc9, 0x5e, 0xa1,
	0xdf, 0x09, 0x44, 0x01, 0x83, 0xbd, 0x1a, 0x90, 0x51, 0x9c, 0x7e, 0x23, 0xc9, 0x7f, 0x7e, 0xf9,
	0x92, 0xa3, 0x67, 0x39, 0x43, 0x3c, 0xfa, 0x83, 0x40, 0xd7, 0x2c, 0x91, 0xf4, 0xf5, 0x4d, 0x25,
	0x6d, 0x01, 0xb3, 0x2e, 0x7e, 0xc6, 0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0x2d, 0xb2,
	0xb7, 0xe1, 0x03, 0x00, 0x00,
}