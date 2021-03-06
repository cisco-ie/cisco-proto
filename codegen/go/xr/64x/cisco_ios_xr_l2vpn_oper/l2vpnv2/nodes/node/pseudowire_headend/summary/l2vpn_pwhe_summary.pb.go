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
// source: l2vpn_pwhe_summary.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_pseudowire_headend_summary

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

type L2VpnPwheSummary_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPwheSummary_KEYS) Reset()         { *m = L2VpnPwheSummary_KEYS{} }
func (m *L2VpnPwheSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheSummary_KEYS) ProtoMessage()    {}
func (*L2VpnPwheSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_24e0a7689ff7e834, []int{0}
}

func (m *L2VpnPwheSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheSummary_KEYS.Unmarshal(m, b)
}
func (m *L2VpnPwheSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheSummary_KEYS.Merge(m, src)
}
func (m *L2VpnPwheSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheSummary_KEYS.Size(m)
}
func (m *L2VpnPwheSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheSummary_KEYS proto.InternalMessageInfo

func (m *L2VpnPwheSummary_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type L2VpnPwheSummary struct {
	Interfaces                         uint32   `protobuf:"varint,50,opt,name=interfaces,proto3" json:"interfaces,omitempty"`
	UpInterfaces                       uint32   `protobuf:"varint,51,opt,name=up_interfaces,json=upInterfaces,proto3" json:"up_interfaces,omitempty"`
	DownInterfaces                     uint32   `protobuf:"varint,52,opt,name=down_interfaces,json=downInterfaces,proto3" json:"down_interfaces,omitempty"`
	AdminDownInterfaces                uint32   `protobuf:"varint,53,opt,name=admin_down_interfaces,json=adminDownInterfaces,proto3" json:"admin_down_interfaces,omitempty"`
	PsuedowireEtherInterfaces          uint32   `protobuf:"varint,54,opt,name=psuedowire_ether_interfaces,json=psuedowireEtherInterfaces,proto3" json:"psuedowire_ether_interfaces,omitempty"`
	UpPsuedowireEtherInterfaces        uint32   `protobuf:"varint,55,opt,name=up_psuedowire_ether_interfaces,json=upPsuedowireEtherInterfaces,proto3" json:"up_psuedowire_ether_interfaces,omitempty"`
	DownPseudowireEtherInterfaces      uint32   `protobuf:"varint,56,opt,name=down_pseudowire_ether_interfaces,json=downPseudowireEtherInterfaces,proto3" json:"down_pseudowire_ether_interfaces,omitempty"`
	AdminDownPseudowireEtherInterfaces uint32   `protobuf:"varint,57,opt,name=admin_down_pseudowire_ether_interfaces,json=adminDownPseudowireEtherInterfaces,proto3" json:"admin_down_pseudowire_ether_interfaces,omitempty"`
	PseudowireIwInterfaces             uint32   `protobuf:"varint,58,opt,name=pseudowire_iw_interfaces,json=pseudowireIwInterfaces,proto3" json:"pseudowire_iw_interfaces,omitempty"`
	UpPseudowireIwInterfaces           uint32   `protobuf:"varint,59,opt,name=up_pseudowire_iw_interfaces,json=upPseudowireIwInterfaces,proto3" json:"up_pseudowire_iw_interfaces,omitempty"`
	DownPseudowireIwInterfaces         uint32   `protobuf:"varint,60,opt,name=down_pseudowire_iw_interfaces,json=downPseudowireIwInterfaces,proto3" json:"down_pseudowire_iw_interfaces,omitempty"`
	AdminDownPseudowireIwInterfaces    uint32   `protobuf:"varint,61,opt,name=admin_down_pseudowire_iw_interfaces,json=adminDownPseudowireIwInterfaces,proto3" json:"admin_down_pseudowire_iw_interfaces,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *L2VpnPwheSummary) Reset()         { *m = L2VpnPwheSummary{} }
func (m *L2VpnPwheSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnPwheSummary) ProtoMessage()    {}
func (*L2VpnPwheSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_24e0a7689ff7e834, []int{1}
}

func (m *L2VpnPwheSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPwheSummary.Unmarshal(m, b)
}
func (m *L2VpnPwheSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPwheSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnPwheSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPwheSummary.Merge(m, src)
}
func (m *L2VpnPwheSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnPwheSummary.Size(m)
}
func (m *L2VpnPwheSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPwheSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPwheSummary proto.InternalMessageInfo

func (m *L2VpnPwheSummary) GetInterfaces() uint32 {
	if m != nil {
		return m.Interfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetUpInterfaces() uint32 {
	if m != nil {
		return m.UpInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetDownInterfaces() uint32 {
	if m != nil {
		return m.DownInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetAdminDownInterfaces() uint32 {
	if m != nil {
		return m.AdminDownInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetPsuedowireEtherInterfaces() uint32 {
	if m != nil {
		return m.PsuedowireEtherInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetUpPsuedowireEtherInterfaces() uint32 {
	if m != nil {
		return m.UpPsuedowireEtherInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetDownPseudowireEtherInterfaces() uint32 {
	if m != nil {
		return m.DownPseudowireEtherInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetAdminDownPseudowireEtherInterfaces() uint32 {
	if m != nil {
		return m.AdminDownPseudowireEtherInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetPseudowireIwInterfaces() uint32 {
	if m != nil {
		return m.PseudowireIwInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetUpPseudowireIwInterfaces() uint32 {
	if m != nil {
		return m.UpPseudowireIwInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetDownPseudowireIwInterfaces() uint32 {
	if m != nil {
		return m.DownPseudowireIwInterfaces
	}
	return 0
}

func (m *L2VpnPwheSummary) GetAdminDownPseudowireIwInterfaces() uint32 {
	if m != nil {
		return m.AdminDownPseudowireIwInterfaces
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnPwheSummary_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.pseudowire_headend.summary.l2vpn_pwhe_summary_KEYS")
	proto.RegisterType((*L2VpnPwheSummary)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.pseudowire_headend.summary.l2vpn_pwhe_summary")
}

func init() { proto.RegisterFile("l2vpn_pwhe_summary.proto", fileDescriptor_24e0a7689ff7e834) }

var fileDescriptor_24e0a7689ff7e834 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4b, 0x4f, 0xf2, 0x40,
	0x14, 0x86, 0xc3, 0xe2, 0xe3, 0x8b, 0x27, 0xa2, 0x49, 0x8d, 0x52, 0x25, 0x20, 0x81, 0x44, 0x59,
	0x75, 0x51, 0xbc, 0xe0, 0x05, 0x13, 0xa3, 0xc4, 0x10, 0x5d, 0x10, 0x5c, 0xb9, 0x9a, 0x54, 0x66,
	0x0c, 0x93, 0x48, 0x67, 0x32, 0xc3, 0x50, 0xfd, 0xbb, 0xfe, 0x12, 0xd3, 0x23, 0x81, 0x53, 0x6e,
	0x9b, 0xa6, 0x3d, 0xe7, 0x79, 0x9f, 0xb6, 0x6f, 0x53, 0xf0, 0x3f, 0xc3, 0x89, 0x8e, 0x99, 0x4e,
	0x86, 0x82, 0x59, 0x37, 0x1a, 0x45, 0xe6, 0x3b, 0xd0, 0x46, 0x8d, 0x95, 0xd7, 0x19, 0x48, 0x3b,
	0x50, 0x4c, 0x2a, 0xcb, 0xbe, 0x0c, 0xfb, 0xc3, 0x94, 0x16, 0x26, 0xc0, 0xd3, 0x49, 0x18, 0xc4,
	0x8a, 0x0b, 0x8b, 0xc7, 0x40, 0x5b, 0xe1, 0xb8, 0x4a, 0xa4, 0x11, 0x6c, 0x28, 0x22, 0x2e, 0x62,
	0x1e, 0x4c, 0x65, 0xb5, 0x10, 0x8a, 0xcb, 0xb7, 0x60, 0xcf, 0x9d, 0xb7, 0x57, 0xaf, 0x08, 0xff,
	0xd3, 0x34, 0x93, 0xdc, 0xcf, 0x55, 0x73, 0x8d, 0xad, 0x7e, 0x3e, 0xbd, 0xec, 0xf2, 0xda, 0xcf,
	0x3f, 0xf0, 0x96, 0x43, 0x5e, 0x05, 0x40, 0xc6, 0x63, 0x61, 0x3e, 0xa2, 0x81, 0xb0, 0x7e, 0x58,
	0xcd, 0x35, 0x0a, 0x7d, 0x32, 0xf1, 0xea, 0x50, 0x70, 0x9a, 0x11, 0xa4, 0x89, 0xc8, 0xb6, 0xd3,
	0xdd, 0x39, 0x74, 0x0a, 0xbb, 0x5c, 0x25, 0x31, 0xc5, 0xce, 0x10, 0xdb, 0x49, 0xc7, 0x04, 0x0c,
	0x61, 0x3f, 0xe2, 0x23, 0x19, 0xb3, 0x45, 0xfc, 0x1c, 0xf1, 0x3d, 0x5c, 0x3e, 0x66, 0x33, 0x77,
	0x50, 0xd2, 0xd6, 0x89, 0x69, 0x15, 0x62, 0x3c, 0x14, 0x86, 0x26, 0x2f, 0x30, 0x79, 0x38, 0x47,
	0x3a, 0x29, 0x41, 0xf2, 0x0f, 0x50, 0x71, 0x9a, 0x6d, 0x52, 0x5c, 0xa2, 0xa2, 0xe4, 0x74, 0x6f,
	0xad, 0xe4, 0x09, 0xaa, 0xf8, 0xc8, 0xe4, 0xa3, 0x2c, 0x69, 0x5a, 0xa8, 0x29, 0xa7, 0x5c, 0x6f,
	0x86, 0x2d, 0x8a, 0xfa, 0x70, 0x42, 0x1a, 0xd8, 0xa4, 0xbb, 0x42, 0x5d, 0x6d, 0x56, 0xc9, 0x7a,
	0x67, 0x0b, 0x7c, 0x22, 0x92, 0x09, 0xb5, 0x5c, 0xa3, 0xe5, 0x60, 0xbe, 0xef, 0x26, 0x24, 0xd9,
	0x86, 0x12, 0x76, 0xb3, 0x26, 0x7c, 0x83, 0x61, 0x3f, 0x2d, 0x66, 0x65, 0xfc, 0x1e, 0xca, 0x8b,
	0xaf, 0x91, 0x15, 0xdc, 0xa2, 0xe0, 0x28, 0x5b, 0x49, 0x46, 0xf1, 0x02, 0xf5, 0xd5, 0x7d, 0x64,
	0x45, 0x6d, 0x14, 0x1d, 0xaf, 0x28, 0x83, 0xda, 0xde, 0xf3, 0xf8, 0x9b, 0x35, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x15, 0xdc, 0x4a, 0xb1, 0x82, 0x03, 0x00, 0x00,
}
