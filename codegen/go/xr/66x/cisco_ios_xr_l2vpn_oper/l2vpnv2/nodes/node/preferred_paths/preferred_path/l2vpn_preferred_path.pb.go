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
// source: l2vpn_preferred_path.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_preferred_paths_preferred_path

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

type L2VpnPreferredPath_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PreferredType        string   `protobuf:"bytes,2,opt,name=preferred_type,json=preferredType,proto3" json:"preferred_type,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPreferredPath_KEYS) Reset()         { *m = L2VpnPreferredPath_KEYS{} }
func (m *L2VpnPreferredPath_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnPreferredPath_KEYS) ProtoMessage()    {}
func (*L2VpnPreferredPath_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_91713ea6b68a0ff2, []int{0}
}

func (m *L2VpnPreferredPath_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPreferredPath_KEYS.Unmarshal(m, b)
}
func (m *L2VpnPreferredPath_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPreferredPath_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnPreferredPath_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPreferredPath_KEYS.Merge(m, src)
}
func (m *L2VpnPreferredPath_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnPreferredPath_KEYS.Size(m)
}
func (m *L2VpnPreferredPath_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPreferredPath_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPreferredPath_KEYS proto.InternalMessageInfo

func (m *L2VpnPreferredPath_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnPreferredPath_KEYS) GetPreferredType() string {
	if m != nil {
		return m.PreferredType
	}
	return ""
}

func (m *L2VpnPreferredPath_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type VcInfo struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	PwidType             string   `protobuf:"bytes,3,opt,name=pwid_type,json=pwidType,proto3" json:"pwid_type,omitempty"`
	Pwid                 uint64   `protobuf:"varint,4,opt,name=pwid,proto3" json:"pwid,omitempty"`
	FeCtype              string   `protobuf:"bytes,5,opt,name=fe_ctype,json=feCtype,proto3" json:"fe_ctype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VcInfo) Reset()         { *m = VcInfo{} }
func (m *VcInfo) String() string { return proto.CompactTextString(m) }
func (*VcInfo) ProtoMessage()    {}
func (*VcInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_91713ea6b68a0ff2, []int{1}
}

func (m *VcInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VcInfo.Unmarshal(m, b)
}
func (m *VcInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VcInfo.Marshal(b, m, deterministic)
}
func (m *VcInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VcInfo.Merge(m, src)
}
func (m *VcInfo) XXX_Size() int {
	return xxx_messageInfo_VcInfo.Size(m)
}
func (m *VcInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VcInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VcInfo proto.InternalMessageInfo

func (m *VcInfo) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *VcInfo) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *VcInfo) GetPwidType() string {
	if m != nil {
		return m.PwidType
	}
	return ""
}

func (m *VcInfo) GetPwid() uint64 {
	if m != nil {
		return m.Pwid
	}
	return 0
}

func (m *VcInfo) GetFeCtype() string {
	if m != nil {
		return m.FeCtype
	}
	return ""
}

type L2VpnPreferredPath struct {
	Type                 string    `protobuf:"bytes,50,opt,name=type,proto3" json:"type,omitempty"`
	InterfaceNameXr      string    `protobuf:"bytes,51,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	TotalBandwidth       uint32    `protobuf:"varint,52,opt,name=total_bandwidth,json=totalBandwidth,proto3" json:"total_bandwidth,omitempty"`
	AvailableBandwidth   uint32    `protobuf:"varint,53,opt,name=available_bandwidth,json=availableBandwidth,proto3" json:"available_bandwidth,omitempty"`
	ReservedBandwidth    uint32    `protobuf:"varint,54,opt,name=reserved_bandwidth,json=reservedBandwidth,proto3" json:"reserved_bandwidth,omitempty"`
	VirtualCircuit       []*VcInfo `protobuf:"bytes,55,rep,name=virtual_circuit,json=virtualCircuit,proto3" json:"virtual_circuit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *L2VpnPreferredPath) Reset()         { *m = L2VpnPreferredPath{} }
func (m *L2VpnPreferredPath) String() string { return proto.CompactTextString(m) }
func (*L2VpnPreferredPath) ProtoMessage()    {}
func (*L2VpnPreferredPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_91713ea6b68a0ff2, []int{2}
}

func (m *L2VpnPreferredPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPreferredPath.Unmarshal(m, b)
}
func (m *L2VpnPreferredPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPreferredPath.Marshal(b, m, deterministic)
}
func (m *L2VpnPreferredPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPreferredPath.Merge(m, src)
}
func (m *L2VpnPreferredPath) XXX_Size() int {
	return xxx_messageInfo_L2VpnPreferredPath.Size(m)
}
func (m *L2VpnPreferredPath) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPreferredPath.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPreferredPath proto.InternalMessageInfo

func (m *L2VpnPreferredPath) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *L2VpnPreferredPath) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *L2VpnPreferredPath) GetTotalBandwidth() uint32 {
	if m != nil {
		return m.TotalBandwidth
	}
	return 0
}

func (m *L2VpnPreferredPath) GetAvailableBandwidth() uint32 {
	if m != nil {
		return m.AvailableBandwidth
	}
	return 0
}

func (m *L2VpnPreferredPath) GetReservedBandwidth() uint32 {
	if m != nil {
		return m.ReservedBandwidth
	}
	return 0
}

func (m *L2VpnPreferredPath) GetVirtualCircuit() []*VcInfo {
	if m != nil {
		return m.VirtualCircuit
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnPreferredPath_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.preferred_paths.preferred_path.l2vpn_preferred_path_KEYS")
	proto.RegisterType((*VcInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.preferred_paths.preferred_path.vc_info")
	proto.RegisterType((*L2VpnPreferredPath)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.preferred_paths.preferred_path.l2vpn_preferred_path")
}

func init() { proto.RegisterFile("l2vpn_preferred_path.proto", fileDescriptor_91713ea6b68a0ff2) }

var fileDescriptor_91713ea6b68a0ff2 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0xbd, 0xf5, 0xf6, 0x76, 0xa4, 0x2d, 0x1d, 0x05, 0x53, 0xdd, 0x84, 0x82, 0x18,
	0x04, 0x23, 0xa4, 0xfe, 0x59, 0x6b, 0x71, 0x51, 0x04, 0x17, 0xd1, 0x85, 0xae, 0x86, 0xe9, 0xcc,
	0x09, 0x1d, 0x48, 0x33, 0xc3, 0x99, 0x69, 0xda, 0xe2, 0xce, 0x77, 0xf0, 0xed, 0x7c, 0x18, 0x99,
	0x49, 0xff, 0x4a, 0x97, 0x6e, 0xc2, 0xf9, 0xbe, 0xf3, 0xcb, 0xf9, 0x66, 0x92, 0x43, 0x9e, 0x56,
	0x79, 0x63, 0x6a, 0x66, 0x10, 0x4a, 0x40, 0x04, 0xc9, 0x0c, 0x77, 0xcb, 0xcc, 0xa0, 0x76, 0x9a,
	0xce, 0x85, 0xb2, 0x42, 0x33, 0xa5, 0x2d, 0xdb, 0x22, 0x6b, 0x41, 0x6d, 0x00, 0xb3, 0x50, 0x36,
	0x79, 0x56, 0x6b, 0x09, 0x36, 0x3c, 0xb3, 0xcb, 0x01, 0xf6, 0x1f, 0x3d, 0xf9, 0x15, 0x91, 0xf1,
	0xb5, 0x24, 0xf6, 0xf9, 0xd3, 0x8f, 0xaf, 0xf4, 0x09, 0xe9, 0xfa, 0x21, 0x4c, 0xc9, 0x38, 0x4a,
	0xa2, 0xb4, 0x57, 0xdc, 0x79, 0x39, 0x97, 0xf4, 0x39, 0x19, 0x9c, 0x78, 0xb7, 0x33, 0x10, 0xdf,
	0x84, 0x7e, 0xff, 0xe8, 0x7e, 0xdb, 0x19, 0xf0, 0x98, 0xaa, 0x1d, 0x60, 0xc9, 0x05, 0xb0, 0x9a,
	0xaf, 0x20, 0xbe, 0x6d, 0xb1, 0xa3, 0xfb, 0x85, 0xaf, 0x60, 0xf2, 0x3b, 0x22, 0xdd, 0x46, 0x30,
	0x55, 0x97, 0xda, 0x47, 0x1a, 0x00, 0x3c, 0x8b, 0xf4, 0xb2, 0x8d, 0xb4, 0x7a, 0x8d, 0x02, 0x18,
	0x97, 0x12, 0xc1, 0xda, 0x43, 0x64, 0xeb, 0x7e, 0x68, 0x4d, 0xfa, 0x8c, 0xf4, 0xcc, 0x46, 0xed,
	0x0f, 0xd5, 0xa6, 0xdd, 0x7b, 0x23, 0x9c, 0x87, 0x92, 0x8e, 0xaf, 0xe3, 0x4e, 0x12, 0xa5, 0x9d,
	0x22, 0xd4, 0x74, 0x4c, 0xee, 0x4b, 0x60, 0x22, 0xf0, 0x0f, 0x02, 0xdf, 0x2d, 0x61, 0xe6, 0xe5,
	0xe4, 0xcf, 0x0d, 0x79, 0x7c, 0xed, 0xe3, 0xf8, 0x39, 0x81, 0xcf, 0x03, 0x1f, 0x6a, 0xfa, 0x92,
	0x8c, 0x2e, 0xef, 0xca, 0xb6, 0x18, 0x4f, 0x03, 0x30, 0xbc, 0xb8, 0xee, 0x77, 0xa4, 0x2f, 0xc8,
	0xd0, 0x69, 0xc7, 0x2b, 0xb6, 0xe0, 0xb5, 0xdc, 0x28, 0xe9, 0x96, 0xf1, 0x9b, 0x24, 0x4a, 0xfb,
	0xc5, 0x20, 0xd8, 0x1f, 0x0f, 0x2e, 0x7d, 0x4d, 0x1e, 0xf1, 0x86, 0xab, 0x8a, 0x2f, 0x2a, 0x38,
	0x83, 0xdf, 0x06, 0x98, 0x1e, 0x5b, 0xa7, 0x17, 0x5e, 0x11, 0x8a, 0x60, 0x01, 0x1b, 0x90, 0x67,
	0xfc, 0xbb, 0xc0, 0x8f, 0x0e, 0x9d, 0x13, 0xfe, 0x93, 0x0c, 0x1b, 0x85, 0x6e, 0xcd, 0x2b, 0x26,
	0x14, 0x8a, 0xb5, 0x72, 0xf1, 0xfb, 0xe4, 0x36, 0x7d, 0x98, 0x17, 0xd9, 0x7f, 0xdb, 0xb1, 0x6c,
	0xff, 0x6b, 0x8b, 0xc1, 0x3e, 0x6a, 0xd6, 0x26, 0x2d, 0xee, 0xc2, 0x36, 0x4f, 0xff, 0x06, 0x00,
	0x00, 0xff, 0xff, 0x85, 0xb8, 0xd5, 0x4d, 0xeb, 0x02, 0x00, 0x00,
}