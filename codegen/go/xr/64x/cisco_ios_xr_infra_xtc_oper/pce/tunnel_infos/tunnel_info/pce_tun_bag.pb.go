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
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x55, 0x92, 0x36, 0x6d, 0x66, 0x93, 0x28, 0x75, 0x8a, 0xb0, 0xc4, 0x81, 0x10, 0xf5, 0x10,
	0x09, 0x69, 0x25, 0xd2, 0x0a, 0x71, 0xe5, 0x00, 0x52, 0x00, 0x71, 0xd8, 0xe5, 0x00, 0x27, 0xcb,
	0x59, 0x7b, 0x2b, 0x8b, 0x5d, 0xdb, 0xd8, 0x4e, 0x15, 0xe0, 0x57, 0xf8, 0x42, 0xbe, 0x02, 0xd9,
	0xde, 0x4d, 0xb7, 0x1c, 0x21, 0xb7, 0x99, 0x37, 0x33, 0xef, 0xcd, 0x3e, 0x8f, 0x16, 0x2e, 0x74,
	0xc1, 0x89, 0xdb, 0x49, 0xb2, 0xa5, 0xb7, 0xa9, 0x36, 0xca, 0x29, 0xf4, 0xaa, 0x10, 0xb6, 0x50,
	0x44, 0x28, 0x4b, 0xf6, 0x86, 0x08, 0x59, 0x1a, 0x4a, 0xf6, 0xae, 0x20, 0x4a, 0x73, 0x93, 0xea,
	0x82, 0xa7, 0x6e, 0x27, 0x25, 0xaf, 0x7c, 0x45, 0xd9, 0x6e, 0xb2, 0x54, 0x30, 0xeb, 0xd0, 0x91,
	0xf7, 0x6f, 0xbe, 0xe4, 0xe8, 0x19, 0x8c, 0x35, 0xe7, 0x86, 0x50, 0xc6, 0x0c, 0xb7, 0x16, 0xf7,
	0x16, 0xbd, 0xd5, 0x28, 0x4b, 0x3c, 0xf6, 0x3a, 0x42, 0xe8, 0x31, 0x9c, 0xe9, 0xca, 0x6a, 0x22,
	0x18, 0xee, 0x2f, 0x7a, 0xab, 0x49, 0x36, 0xf4, 0xe9, 0x86, 0xa1, 0xa7, 0x90, 0x34, 0xf4, 0x92,
	0xd6, 0x1c, 0x0f, 0xc2, 0x28, 0x44, 0xe8, 0x23, 0xad, 0xf9, 0x32, 0x8f, 0x82, 0x42, 0x07, 0x7a,
	0xe2, 0xbe, 0x6b, 0xee, 0xd9, 0x68, 0x19, 0x07, 0xa2, 0xd6, 0x90, 0x96, 0xbe, 0x19, 0x21, 0x38,
	0x11, 0xfa, 0xee, 0x26, 0x68, 0x8c, 0xb2, 0x10, 0x37, 0xd8, 0xcb, 0x86, 0x3a, 0xc4, 0xcb, 0xdf,
	0x03, 0x18, 0x7b, 0xd6, 0xb0, 0x92, 0x2c, 0x15, 0xfa, 0x06, 0x53, 0xab, 0x76, 0xa6, 0xe0, 0x0f,
	0x3e, 0x22, 0x59, 0xbf, 0x4b, 0xff, 0xd5, 0xa9, 0xf4, 0xef, 0xad, 0xb3, 0x49, 0x54, 0x68, 0x2d,
	0xf9, 0x09, 0x73, 0xc6, 0xad, 0x13, 0x92, 0x3a, 0xa1, 0xe4, 0x41, 0xb7, 0x7f, 0x74, 0x5d, 0xd4,
	0x91, 0x69, 0xc5, 0x9f, 0xc0, 0xa8, 0x9d, 0x61, 0xc1, 0x99, 0x49, 0x76, 0x1e, 0x81, 0x0d, 0x43,
	0x97, 0x70, 0x5a, 0x59, 0x2d, 0x18, 0x3e, 0x09, 0x85, 0x98, 0xf8, 0x97, 0xda, 0x0a, 0xc9, 0x84,
	0xbc, 0x25, 0x56, 0x30, 0x7c, 0x1a, 0x6a, 0xd0, 0x40, 0xb9, 0x60, 0xe8, 0x0a, 0xa6, 0xde, 0x4f,
	0xcb, 0xdd, 0x4e, 0x07, 0x65, 0x3c, 0x0c, 0x96, 0x8f, 0x2b, 0xab, 0x73, 0x0f, 0x7e, 0xf2, 0x6f,
	0xf7, 0x1c, 0x2e, 0xfc, 0x37, 0x84, 0x6d, 0x68, 0x45, 0xac, 0xa3, 0x8e, 0xe3, 0xb3, 0xd0, 0x38,
	0xeb, 0x14, 0x72, 0x8f, 0xa3, 0x17, 0x70, 0x49, 0x59, 0x2d, 0xa4, 0xb0, 0xce, 0x17, 0xee, 0x78,
	0xd3, 0x7f, 0x1e, 0xfa, 0xe7, 0x0f, 0x6b, 0x71, 0x64, 0x06, 0x83, 0xda, 0x32, 0x3c, 0x0a, 0xeb,
	0xf9, 0x70, 0xf9, 0xab, 0x0f, 0x49, 0xe7, 0x66, 0xd1, 0x57, 0x9f, 0x16, 0x07, 0xc3, 0xd7, 0x47,
	0x37, 0x1c, 0x74, 0x51, 0xb4, 0x46, 0x5f, 0xc1, 0xb4, 0x73, 0xdf, 0x64, 0x6f, 0xf0, 0x75, 0x34,
	0xe5, 0xfe, 0xc4, 0x3f, 0x1b, 0xf4, 0x03, 0x1e, 0x6d, 0x8d, 0xe0, 0xe5, 0xe1, 0x20, 0x4d, 0x1d,
	0x7c, 0xc0, 0x37, 0x8b, 0xc1, 0x2a, 0x59, 0xbf, 0xfd, 0xbf, 0xe5, 0x5a, 0xd2, 0x6c, 0x1e, 0x44,
	0x3e, 0x58, 0xbd, 0xb9, 0x97, 0xd8, 0x0e, 0xc3, 0x2f, 0xe1, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0xa4, 0xfa, 0x59, 0x27, 0x04, 0x00, 0x00,
}