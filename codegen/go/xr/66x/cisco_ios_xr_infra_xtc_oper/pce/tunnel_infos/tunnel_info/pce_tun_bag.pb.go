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
// source: pce_tun_bag.proto

package cisco_ios_xr_infra_xtc_oper_pce_tunnel_infos_tunnel_info

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

type PceTunBag_KEYS struct {
	PeerAddress          string   `protobuf:"bytes,1,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	PlspId               uint32   `protobuf:"varint,2,opt,name=plsp_id,json=plspId,proto3" json:"plsp_id,omitempty"`
	TunnelName           string   `protobuf:"bytes,3,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceTunBag_KEYS) Reset()         { *m = PceTunBag_KEYS{} }
func (m *PceTunBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceTunBag_KEYS) ProtoMessage()    {}
func (*PceTunBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_946468a27bd8a754, []int{0}
}

func (m *PceTunBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTunBag_KEYS.Unmarshal(m, b)
}
func (m *PceTunBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTunBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceTunBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTunBag_KEYS.Merge(m, src)
}
func (m *PceTunBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceTunBag_KEYS.Size(m)
}
func (m *PceTunBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTunBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceTunBag_KEYS proto.InternalMessageInfo

func (m *PceTunBag_KEYS) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *PceTunBag_KEYS) GetPlspId() uint32 {
	if m != nil {
		return m.PlspId
	}
	return 0
}

func (m *PceTunBag_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

type PceIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceIpAddrType) Reset()         { *m = PceIpAddrType{} }
func (m *PceIpAddrType) String() string { return proto.CompactTextString(m) }
func (*PceIpAddrType) ProtoMessage()    {}
func (*PceIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_946468a27bd8a754, []int{1}
}

func (m *PceIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceIpAddrType.Unmarshal(m, b)
}
func (m *PceIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceIpAddrType.Marshal(b, m, deterministic)
}
func (m *PceIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceIpAddrType.Merge(m, src)
}
func (m *PceIpAddrType) XXX_Size() int {
	return xxx_messageInfo_PceIpAddrType.Size(m)
}
func (m *PceIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PceIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PceIpAddrType proto.InternalMessageInfo

func (m *PceIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PceIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *PceIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type PceLspInfo struct {
	SourceAddress        *PceIpAddrType `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   *PceIpAddrType `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	TunnelId             uint32         `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	Lspid                uint32         `protobuf:"varint,4,opt,name=lspid,proto3" json:"lspid,omitempty"`
	BindingSid           uint32         `protobuf:"varint,5,opt,name=binding_sid,json=bindingSid,proto3" json:"binding_sid,omitempty"`
	LspSetupType         string         `protobuf:"bytes,6,opt,name=lsp_setup_type,json=lspSetupType,proto3" json:"lsp_setup_type,omitempty"`
	OperationalState     string         `protobuf:"bytes,7,opt,name=operational_state,json=operationalState,proto3" json:"operational_state,omitempty"`
	AdministrativeState  string         `protobuf:"bytes,8,opt,name=administrative_state,json=administrativeState,proto3" json:"administrative_state,omitempty"`
	Msd                  uint32         `protobuf:"varint,9,opt,name=msd,proto3" json:"msd,omitempty"`
	AbsoluteMargin       uint32         `protobuf:"varint,10,opt,name=absolute_margin,json=absoluteMargin,proto3" json:"absolute_margin,omitempty"`
	RelativeMargin       uint32         `protobuf:"varint,11,opt,name=relative_margin,json=relativeMargin,proto3" json:"relative_margin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PceLspInfo) Reset()         { *m = PceLspInfo{} }
func (m *PceLspInfo) String() string { return proto.CompactTextString(m) }
func (*PceLspInfo) ProtoMessage()    {}
func (*PceLspInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_946468a27bd8a754, []int{2}
}

func (m *PceLspInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceLspInfo.Unmarshal(m, b)
}
func (m *PceLspInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceLspInfo.Marshal(b, m, deterministic)
}
func (m *PceLspInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceLspInfo.Merge(m, src)
}
func (m *PceLspInfo) XXX_Size() int {
	return xxx_messageInfo_PceLspInfo.Size(m)
}
func (m *PceLspInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PceLspInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PceLspInfo proto.InternalMessageInfo

func (m *PceLspInfo) GetSourceAddress() *PceIpAddrType {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *PceLspInfo) GetDestinationAddress() *PceIpAddrType {
	if m != nil {
		return m.DestinationAddress
	}
	return nil
}

func (m *PceLspInfo) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *PceLspInfo) GetLspid() uint32 {
	if m != nil {
		return m.Lspid
	}
	return 0
}

func (m *PceLspInfo) GetBindingSid() uint32 {
	if m != nil {
		return m.BindingSid
	}
	return 0
}

func (m *PceLspInfo) GetLspSetupType() string {
	if m != nil {
		return m.LspSetupType
	}
	return ""
}

func (m *PceLspInfo) GetOperationalState() string {
	if m != nil {
		return m.OperationalState
	}
	return ""
}

func (m *PceLspInfo) GetAdministrativeState() string {
	if m != nil {
		return m.AdministrativeState
	}
	return ""
}

func (m *PceLspInfo) GetMsd() uint32 {
	if m != nil {
		return m.Msd
	}
	return 0
}

func (m *PceLspInfo) GetAbsoluteMargin() uint32 {
	if m != nil {
		return m.AbsoluteMargin
	}
	return 0
}

func (m *PceLspInfo) GetRelativeMargin() uint32 {
	if m != nil {
		return m.RelativeMargin
	}
	return 0
}

type PceTunBag struct {
	PccAddress           *PceIpAddrType `protobuf:"bytes,50,opt,name=pcc_address,json=pccAddress,proto3" json:"pcc_address,omitempty"`
	TunnelNameXr         string         `protobuf:"bytes,51,opt,name=tunnel_name_xr,json=tunnelNameXr,proto3" json:"tunnel_name_xr,omitempty"`
	BriefLspInformation  []*PceLspInfo  `protobuf:"bytes,52,rep,name=brief_lsp_information,json=briefLspInformation,proto3" json:"brief_lsp_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PceTunBag) Reset()         { *m = PceTunBag{} }
func (m *PceTunBag) String() string { return proto.CompactTextString(m) }
func (*PceTunBag) ProtoMessage()    {}
func (*PceTunBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_946468a27bd8a754, []int{3}
}

func (m *PceTunBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTunBag.Unmarshal(m, b)
}
func (m *PceTunBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTunBag.Marshal(b, m, deterministic)
}
func (m *PceTunBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTunBag.Merge(m, src)
}
func (m *PceTunBag) XXX_Size() int {
	return xxx_messageInfo_PceTunBag.Size(m)
}
func (m *PceTunBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTunBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceTunBag proto.InternalMessageInfo

func (m *PceTunBag) GetPccAddress() *PceIpAddrType {
	if m != nil {
		return m.PccAddress
	}
	return nil
}

func (m *PceTunBag) GetTunnelNameXr() string {
	if m != nil {
		return m.TunnelNameXr
	}
	return ""
}

func (m *PceTunBag) GetBriefLspInformation() []*PceLspInfo {
	if m != nil {
		return m.BriefLspInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*PceTunBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_oper.pce.tunnel_infos.tunnel_info.pce_tun_bag_KEYS")
	proto.RegisterType((*PceIpAddrType)(nil), "cisco_ios_xr_infra_xtc_oper.pce.tunnel_infos.tunnel_info.pce_ip_addr_type")
	proto.RegisterType((*PceLspInfo)(nil), "cisco_ios_xr_infra_xtc_oper.pce.tunnel_infos.tunnel_info.pce_lsp_info")
	proto.RegisterType((*PceTunBag)(nil), "cisco_ios_xr_infra_xtc_oper.pce.tunnel_infos.tunnel_info.pce_tun_bag")
}

func init() { proto.RegisterFile("pce_tun_bag.proto", fileDescriptor_946468a27bd8a754) }

var fileDescriptor_946468a27bd8a754 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0x55, 0xfe, 0x34, 0x6d, 0x66, 0x93, 0xfc, 0x52, 0xa7, 0x3f, 0x61, 0x89, 0x03, 0x21, 0xaa,
	0x44, 0x24, 0xa4, 0x48, 0xa4, 0x15, 0xe2, 0xca, 0x01, 0xa4, 0xf0, 0xef, 0xb0, 0xe1, 0x00, 0x27,
	0xcb, 0x59, 0x7b, 0x23, 0x8b, 0x5d, 0xdb, 0xd8, 0xde, 0x2a, 0xc0, 0x57, 0xe1, 0xa3, 0xf0, 0xe1,
	0x90, 0xed, 0xdd, 0x74, 0xcb, 0x11, 0x7a, 0xf3, 0xbc, 0x79, 0x33, 0x6f, 0xfc, 0x66, 0xbd, 0x70,
	0xae, 0x33, 0x4e, 0x5c, 0x25, 0xc9, 0x8e, 0xee, 0x57, 0xda, 0x28, 0xa7, 0xd0, 0x8b, 0x4c, 0xd8,
	0x4c, 0x11, 0xa1, 0x2c, 0x39, 0x18, 0x22, 0x64, 0x6e, 0x28, 0x39, 0xb8, 0x8c, 0x28, 0xcd, 0xcd,
	0x4a, 0x67, 0x7c, 0xe5, 0x2a, 0x29, 0x79, 0xe1, 0x33, 0xca, 0xb6, 0x83, 0x85, 0x82, 0x69, 0xab,
	0x1d, 0x79, 0xfb, 0xea, 0xf3, 0x16, 0x3d, 0x86, 0x91, 0xe6, 0xdc, 0x10, 0xca, 0x98, 0xe1, 0xd6,
	0xe2, 0xce, 0xbc, 0xb3, 0x1c, 0xa6, 0x89, 0xc7, 0x5e, 0x46, 0x08, 0x3d, 0x80, 0x53, 0x5d, 0x58,
	0x4d, 0x04, 0xc3, 0xdd, 0x79, 0x67, 0x39, 0x4e, 0x07, 0x3e, 0xdc, 0x30, 0xf4, 0x08, 0x92, 0xba,
	0xbd, 0xa4, 0x25, 0xc7, 0xbd, 0x50, 0x0a, 0x11, 0xfa, 0x40, 0x4b, 0xbe, 0xd8, 0x46, 0x41, 0xa1,
	0x43, 0x7b, 0xe2, 0xbe, 0x69, 0xee, 0xbb, 0xd1, 0x3c, 0x16, 0x44, 0xad, 0x01, 0xcd, 0x3d, 0x19,
	0x21, 0xe8, 0x0b, 0x7d, 0x73, 0x1d, 0x34, 0x86, 0x69, 0x38, 0xd7, 0xd8, 0xf3, 0xba, 0x75, 0x38,
	0x2f, 0x7e, 0xf5, 0x61, 0xe4, 0xbb, 0x86, 0x91, 0x64, 0xae, 0xd0, 0x57, 0x98, 0x58, 0x55, 0x99,
	0x8c, 0xdf, 0xb9, 0x44, 0xb2, 0x7e, 0xb3, 0xfa, 0x5b, 0xa7, 0x56, 0x7f, 0x4e, 0x9d, 0x8e, 0xa3,
	0x42, 0x63, 0xc9, 0x0f, 0x98, 0x31, 0x6e, 0x9d, 0x90, 0xd4, 0x09, 0x25, 0x8f, 0xba, 0xdd, 0x7b,
	0xd7, 0x45, 0x2d, 0x99, 0x46, 0xfc, 0x21, 0x0c, 0x9b, 0x1a, 0x16, 0x9c, 0x19, 0xa7, 0x67, 0x11,
	0xd8, 0x30, 0x74, 0x01, 0x27, 0x85, 0xd5, 0x82, 0xe1, 0x7e, 0x48, 0xc4, 0xc0, 0x6f, 0x6a, 0x27,
	0x24, 0x13, 0x72, 0x4f, 0xac, 0x60, 0xf8, 0x24, 0xe4, 0xa0, 0x86, 0xb6, 0x82, 0xa1, 0x4b, 0x98,
	0x78, 0x3f, 0x2d, 0x77, 0x95, 0x0e, 0xca, 0x78, 0x10, 0x2c, 0x1f, 0x15, 0x56, 0x6f, 0x3d, 0xf8,
	0xd1, 0xef, 0xee, 0x29, 0x9c, 0xfb, 0x3b, 0x84, 0x69, 0x68, 0x41, 0xac, 0xa3, 0x8e, 0xe3, 0xd3,
	0x40, 0x9c, 0xb6, 0x12, 0x5b, 0x8f, 0xa3, 0x67, 0x70, 0x41, 0x59, 0x29, 0xa4, 0xb0, 0xce, 0x27,
	0x6e, 0x78, 0xcd, 0x3f, 0x0b, 0xfc, 0xd9, 0xdd, 0x5c, 0x2c, 0x99, 0x42, 0xaf, 0xb4, 0x0c, 0x0f,
	0xc3, 0x78, 0xfe, 0x88, 0x9e, 0xc0, 0x7f, 0x74, 0x67, 0x55, 0x51, 0x39, 0x4e, 0x4a, 0x6a, 0xf6,
	0x42, 0x62, 0x08, 0xd9, 0x49, 0x03, 0xbf, 0x0f, 0xa8, 0x27, 0x1a, 0x5e, 0x44, 0x9d, 0x9a, 0x98,
	0x44, 0x62, 0x03, 0x47, 0xe2, 0xe2, 0x67, 0x17, 0x92, 0xd6, 0x2b, 0x40, 0x5f, 0x7c, 0x98, 0x1d,
	0x57, 0xb8, 0xbe, 0xf7, 0x15, 0x82, 0xce, 0xb2, 0x66, 0x75, 0x97, 0x30, 0x69, 0xbd, 0x18, 0x72,
	0x30, 0xf8, 0x2a, 0xda, 0x7c, 0xfb, 0x68, 0x3e, 0x19, 0xf4, 0x1d, 0xfe, 0xdf, 0x19, 0xc1, 0xf3,
	0xe3, 0x27, 0x6e, 0xca, 0xe0, 0x2c, 0xbe, 0x9e, 0xf7, 0x96, 0xc9, 0xfa, 0xf5, 0xbf, 0x0d, 0xd7,
	0x34, 0x4d, 0x67, 0x41, 0xe4, 0x9d, 0xd5, 0x9b, 0x5b, 0x89, 0xdd, 0x20, 0xfc, 0x64, 0xae, 0x7e,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xd0, 0xa2, 0xa2, 0x79, 0x04, 0x00, 0x00,
}
