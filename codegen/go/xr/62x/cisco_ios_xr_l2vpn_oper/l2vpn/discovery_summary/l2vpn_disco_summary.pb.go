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
// source: l2vpn_disco_summary.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_discovery_summary

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

type L2VpnDiscoSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnDiscoSummary_KEYS) Reset()         { *m = L2VpnDiscoSummary_KEYS{} }
func (m *L2VpnDiscoSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoSummary_KEYS) ProtoMessage()    {}
func (*L2VpnDiscoSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfcf485cee683412, []int{0}
}

func (m *L2VpnDiscoSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoSummary_KEYS.Unmarshal(m, b)
}
func (m *L2VpnDiscoSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoSummary_KEYS.Merge(m, src)
}
func (m *L2VpnDiscoSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoSummary_KEYS.Size(m)
}
func (m *L2VpnDiscoSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoSummary_KEYS proto.InternalMessageInfo

type L2VpnDiscoSummary struct {
	NumberBridgeDomainvpns        uint32   `protobuf:"varint,50,opt,name=number_bridge_domainvpns,json=numberBridgeDomainvpns,proto3" json:"number_bridge_domainvpns,omitempty"`
	NumberMp2MPxconnectVpns       uint32   `protobuf:"varint,51,opt,name=number_mp2m_pxconnect_vpns,json=numberMp2mPxconnectVpns,proto3" json:"number_mp2m_pxconnect_vpns,omitempty"`
	NumberLocalEdgesBridgeDomain  uint32   `protobuf:"varint,52,opt,name=number_local_edges_bridge_domain,json=numberLocalEdgesBridgeDomain,proto3" json:"number_local_edges_bridge_domain,omitempty"`
	NumberRemoteEdgesBridgeDomain uint32   `protobuf:"varint,53,opt,name=number_remote_edges_bridge_domain,json=numberRemoteEdgesBridgeDomain,proto3" json:"number_remote_edges_bridge_domain,omitempty"`
	NumberNlriBridgeDomain        uint32   `protobuf:"varint,54,opt,name=number_nlri_bridge_domain,json=numberNlriBridgeDomain,proto3" json:"number_nlri_bridge_domain,omitempty"`
	NumberLocalEdgesXconnect      uint32   `protobuf:"varint,55,opt,name=number_local_edges_xconnect,json=numberLocalEdgesXconnect,proto3" json:"number_local_edges_xconnect,omitempty"`
	NumberRemoteEdgesXconnect     uint32   `protobuf:"varint,56,opt,name=number_remote_edges_xconnect,json=numberRemoteEdgesXconnect,proto3" json:"number_remote_edges_xconnect,omitempty"`
	NumberNlriXconnect            uint32   `protobuf:"varint,57,opt,name=number_nlri_xconnect,json=numberNlriXconnect,proto3" json:"number_nlri_xconnect,omitempty"`
	BgpStateonActiveRp            bool     `protobuf:"varint,58,opt,name=bgp_stateon_active_rp,json=bgpStateonActiveRp,proto3" json:"bgp_stateon_active_rp,omitempty"`
	BgpStateonStandbyRp           bool     `protobuf:"varint,59,opt,name=bgp_stateon_standby_rp,json=bgpStateonStandbyRp,proto3" json:"bgp_stateon_standby_rp,omitempty"`
	VplsRegistered                bool     `protobuf:"varint,60,opt,name=vpls_registered,json=vplsRegistered,proto3" json:"vpls_registered,omitempty"`
	VpwsRegistered                bool     `protobuf:"varint,61,opt,name=vpws_registered,json=vpwsRegistered,proto3" json:"vpws_registered,omitempty"`
	BgpIpcTransportMode           string   `protobuf:"bytes,62,opt,name=bgp_ipc_transport_mode,json=bgpIpcTransportMode,proto3" json:"bgp_ipc_transport_mode,omitempty"`
	BgpCurrentNodeId              string   `protobuf:"bytes,63,opt,name=bgp_current_node_id,json=bgpCurrentNodeId,proto3" json:"bgp_current_node_id,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *L2VpnDiscoSummary) Reset()         { *m = L2VpnDiscoSummary{} }
func (m *L2VpnDiscoSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoSummary) ProtoMessage()    {}
func (*L2VpnDiscoSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfcf485cee683412, []int{1}
}

func (m *L2VpnDiscoSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoSummary.Unmarshal(m, b)
}
func (m *L2VpnDiscoSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoSummary.Merge(m, src)
}
func (m *L2VpnDiscoSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoSummary.Size(m)
}
func (m *L2VpnDiscoSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoSummary proto.InternalMessageInfo

func (m *L2VpnDiscoSummary) GetNumberBridgeDomainvpns() uint32 {
	if m != nil {
		return m.NumberBridgeDomainvpns
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberMp2MPxconnectVpns() uint32 {
	if m != nil {
		return m.NumberMp2MPxconnectVpns
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberLocalEdgesBridgeDomain() uint32 {
	if m != nil {
		return m.NumberLocalEdgesBridgeDomain
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberRemoteEdgesBridgeDomain() uint32 {
	if m != nil {
		return m.NumberRemoteEdgesBridgeDomain
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberNlriBridgeDomain() uint32 {
	if m != nil {
		return m.NumberNlriBridgeDomain
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberLocalEdgesXconnect() uint32 {
	if m != nil {
		return m.NumberLocalEdgesXconnect
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberRemoteEdgesXconnect() uint32 {
	if m != nil {
		return m.NumberRemoteEdgesXconnect
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetNumberNlriXconnect() uint32 {
	if m != nil {
		return m.NumberNlriXconnect
	}
	return 0
}

func (m *L2VpnDiscoSummary) GetBgpStateonActiveRp() bool {
	if m != nil {
		return m.BgpStateonActiveRp
	}
	return false
}

func (m *L2VpnDiscoSummary) GetBgpStateonStandbyRp() bool {
	if m != nil {
		return m.BgpStateonStandbyRp
	}
	return false
}

func (m *L2VpnDiscoSummary) GetVplsRegistered() bool {
	if m != nil {
		return m.VplsRegistered
	}
	return false
}

func (m *L2VpnDiscoSummary) GetVpwsRegistered() bool {
	if m != nil {
		return m.VpwsRegistered
	}
	return false
}

func (m *L2VpnDiscoSummary) GetBgpIpcTransportMode() string {
	if m != nil {
		return m.BgpIpcTransportMode
	}
	return ""
}

func (m *L2VpnDiscoSummary) GetBgpCurrentNodeId() string {
	if m != nil {
		return m.BgpCurrentNodeId
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnDiscoSummary_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.discovery_summary.l2vpn_disco_summary_KEYS")
	proto.RegisterType((*L2VpnDiscoSummary)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.discovery_summary.l2vpn_disco_summary")
}

func init() { proto.RegisterFile("l2vpn_disco_summary.proto", fileDescriptor_cfcf485cee683412) }

var fileDescriptor_cfcf485cee683412 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xdb, 0x6b, 0x13, 0x41,
	0x14, 0xc6, 0xc9, 0x4b, 0xd1, 0x01, 0x2f, 0x4c, 0xb5, 0x4e, 0x6a, 0x85, 0xd8, 0x17, 0xf3, 0x62,
	0xd4, 0xc4, 0x4b, 0x6b, 0xad, 0xc5, 0x4b, 0xc5, 0xa2, 0x2d, 0xb2, 0x11, 0xd1, 0xa7, 0xc3, 0xee,
	0xce, 0x61, 0x19, 0xd8, 0x9d, 0x19, 0x66, 0x26, 0x69, 0xf3, 0x27, 0xf9, 0x5f, 0x96, 0x3d, 0x7b,
	0xe9, 0xa6, 0xdd, 0xb7, 0x30, 0xdf, 0xef, 0xf7, 0xf1, 0x9d, 0xc0, 0xb2, 0x61, 0x3e, 0x5d, 0x5a,
	0x0d, 0x52, 0xf9, 0xd4, 0x80, 0x5f, 0x14, 0x45, 0xec, 0x56, 0x13, 0xeb, 0x4c, 0x30, 0xfc, 0x45,
	0x4a, 0x8f, 0xca, 0x78, 0xb8, 0x70, 0x50, 0x71, 0xc6, 0xa2, 0x9b, 0xd0, 0xcf, 0x09, 0x29, 0x4b,
	0x74, 0xab, 0x46, 0xdb, 0xdd, 0x66, 0xa2, 0xa7, 0x0d, 0x7e, 0x1c, 0xff, 0x9b, 0xef, 0xfe, 0xdf,
	0x60, 0x9b, 0x3d, 0x21, 0xdf, 0x63, 0x42, 0x2f, 0x8a, 0x04, 0x1d, 0x24, 0x4e, 0xc9, 0x0c, 0x41,
	0x9a, 0x22, 0x56, 0x7a, 0x69, 0xb5, 0x17, 0xd3, 0xd1, 0x60, 0x7c, 0x27, 0xda, 0xaa, 0xf2, 0xcf,
	0x14, 0x7f, 0x6d, 0x53, 0x7e, 0xc0, 0xb6, 0x6b, 0xb3, 0xb0, 0xd3, 0x02, 0xec, 0x45, 0x6a, 0xb4,
	0xc6, 0x34, 0x00, 0xb9, 0x33, 0x72, 0x1f, 0x55, 0xc4, 0xa9, 0x9d, 0x16, 0xbf, 0x9a, 0xfc, 0x4f,
	0x29, 0x7f, 0x63, 0xa3, 0x5a, 0xce, 0x4d, 0x1a, 0xe7, 0x80, 0x32, 0x43, 0xbf, 0x3e, 0x41, 0xbc,
	0xa6, 0x8a, 0x9d, 0x8a, 0xfb, 0x59, 0x62, 0xc7, 0x25, 0xd5, 0x1d, 0xc2, 0xbf, 0xb3, 0xa7, 0x75,
	0x8f, 0xc3, 0xc2, 0x04, 0xec, 0x2d, 0x7a, 0x43, 0x45, 0x4f, 0x2a, 0x30, 0x22, 0xee, 0x66, 0xd3,
	0x3e, 0x1b, 0xd6, 0x4d, 0x3a, 0x77, 0xea, 0x5a, 0xc3, 0xdb, 0xee, 0x3f, 0x71, 0x96, 0x3b, 0xb5,
	0xa6, 0x1e, 0xb2, 0xc7, 0x3d, 0xc7, 0x34, 0xf7, 0x8a, 0x77, 0x24, 0x8b, 0xeb, 0x77, 0xfc, 0xad,
	0x73, 0x7e, 0xc4, 0x76, 0xfa, 0x6e, 0x68, 0xfd, 0x3d, 0xf2, 0x87, 0x37, 0xe6, 0xb7, 0x05, 0x2f,
	0xd9, 0x83, 0xee, 0xf4, 0x56, 0xdc, 0x27, 0x91, 0x5f, 0xad, 0x6e, 0x8d, 0x57, 0xec, 0x61, 0x92,
	0x59, 0xf0, 0x21, 0x0e, 0x68, 0x34, 0xc4, 0x69, 0x50, 0x4b, 0x04, 0x67, 0xc5, 0xfb, 0xd1, 0x60,
	0x7c, 0x2b, 0xe2, 0x49, 0x66, 0xe7, 0x55, 0xf6, 0x89, 0xa2, 0xc8, 0xf2, 0x19, 0xdb, 0xea, 0x2a,
	0x3e, 0xc4, 0x5a, 0x26, 0xab, 0xd2, 0x39, 0x20, 0x67, 0xf3, 0xca, 0x99, 0x57, 0x59, 0x64, 0xf9,
	0x33, 0x76, 0x6f, 0x69, 0x73, 0x0f, 0x0e, 0x33, 0xe5, 0x03, 0x3a, 0x94, 0xe2, 0x03, 0xd1, 0x77,
	0xcb, 0xe7, 0xa8, 0x7d, 0xad, 0xc0, 0xf3, 0x35, 0xf0, 0xb0, 0x01, 0xcf, 0xbb, 0x60, 0x3d, 0x43,
	0xd9, 0x14, 0x82, 0x8b, 0xb5, 0xb7, 0xc6, 0x05, 0x28, 0x8c, 0x44, 0xf1, 0x71, 0x34, 0x18, 0xdf,
	0xa6, 0x19, 0x27, 0x36, 0xfd, 0xdd, 0x64, 0xa7, 0x46, 0x22, 0x7f, 0xce, 0xca, 0x67, 0x48, 0x17,
	0xce, 0xa1, 0x0e, 0xa0, 0x8d, 0x44, 0x50, 0x52, 0x1c, 0x91, 0x71, 0x3f, 0xc9, 0xec, 0x97, 0x2a,
	0x39, 0x33, 0x12, 0x4f, 0x64, 0xb2, 0x41, 0xdf, 0xdf, 0xec, 0x32, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0xb5, 0xdf, 0x0e, 0x9c, 0x03, 0x00, 0x00,
}
