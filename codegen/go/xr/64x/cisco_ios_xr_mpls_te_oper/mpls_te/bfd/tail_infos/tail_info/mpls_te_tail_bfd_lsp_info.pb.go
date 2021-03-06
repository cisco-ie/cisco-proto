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
// source: mpls_te_tail_bfd_lsp_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_bfd_tail_infos_tail_info

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

type MplsTeTailBfdLspInfo_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,3,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,5,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	Ctype                string   `protobuf:"bytes,6,opt,name=ctype,proto3" json:"ctype,omitempty"`
	P2MpId               uint32   `protobuf:"varint,7,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeTailBfdLspInfo_KEYS) Reset()         { *m = MplsTeTailBfdLspInfo_KEYS{} }
func (m *MplsTeTailBfdLspInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeTailBfdLspInfo_KEYS) ProtoMessage()    {}
func (*MplsTeTailBfdLspInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_325e955e031246e6, []int{0}
}

func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Merge(m, src)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.Size(m)
}
func (m *MplsTeTailBfdLspInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTailBfdLspInfo_KEYS proto.InternalMessageInfo

func (m *MplsTeTailBfdLspInfo_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetCtype() string {
	if m != nil {
		return m.Ctype
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

type TeLspFecCtypeDataT struct {
	FecCtype             string   `protobuf:"bytes,1,opt,name=fec_ctype,json=fecCtype,proto3" json:"fec_ctype,omitempty"`
	P2PLspDestination    string   `protobuf:"bytes,2,opt,name=p2p_lsp_destination,json=p2pLspDestination,proto3" json:"p2p_lsp_destination,omitempty"`
	FecDestinationP2MpId uint32   `protobuf:"varint,3,opt,name=fec_destination_p2mp_id,json=fecDestinationP2mpId,proto3" json:"fec_destination_p2mp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeLspFecCtypeDataT) Reset()         { *m = TeLspFecCtypeDataT{} }
func (m *TeLspFecCtypeDataT) String() string { return proto.CompactTextString(m) }
func (*TeLspFecCtypeDataT) ProtoMessage()    {}
func (*TeLspFecCtypeDataT) Descriptor() ([]byte, []int) {
	return fileDescriptor_325e955e031246e6, []int{1}
}

func (m *TeLspFecCtypeDataT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeLspFecCtypeDataT.Unmarshal(m, b)
}
func (m *TeLspFecCtypeDataT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeLspFecCtypeDataT.Marshal(b, m, deterministic)
}
func (m *TeLspFecCtypeDataT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeLspFecCtypeDataT.Merge(m, src)
}
func (m *TeLspFecCtypeDataT) XXX_Size() int {
	return xxx_messageInfo_TeLspFecCtypeDataT.Size(m)
}
func (m *TeLspFecCtypeDataT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeLspFecCtypeDataT.DiscardUnknown(m)
}

var xxx_messageInfo_TeLspFecCtypeDataT proto.InternalMessageInfo

func (m *TeLspFecCtypeDataT) GetFecCtype() string {
	if m != nil {
		return m.FecCtype
	}
	return ""
}

func (m *TeLspFecCtypeDataT) GetP2PLspDestination() string {
	if m != nil {
		return m.P2PLspDestination
	}
	return ""
}

func (m *TeLspFecCtypeDataT) GetFecDestinationP2MpId() uint32 {
	if m != nil {
		return m.FecDestinationP2MpId
	}
	return 0
}

type TeLspFecT struct {
	FecLspId             uint32              `protobuf:"varint,1,opt,name=fec_lsp_id,json=fecLspId,proto3" json:"fec_lsp_id,omitempty"`
	FecTunnelId          uint32              `protobuf:"varint,2,opt,name=fec_tunnel_id,json=fecTunnelId,proto3" json:"fec_tunnel_id,omitempty"`
	FecExtendedTunnelId  string              `protobuf:"bytes,3,opt,name=fec_extended_tunnel_id,json=fecExtendedTunnelId,proto3" json:"fec_extended_tunnel_id,omitempty"`
	FecSource            string              `protobuf:"bytes,4,opt,name=fec_source,json=fecSource,proto3" json:"fec_source,omitempty"`
	FecDestinationInfo   *TeLspFecCtypeDataT `protobuf:"bytes,5,opt,name=fec_destination_info,json=fecDestinationInfo,proto3" json:"fec_destination_info,omitempty"`
	FecVrf               string              `protobuf:"bytes,6,opt,name=fec_vrf,json=fecVrf,proto3" json:"fec_vrf,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TeLspFecT) Reset()         { *m = TeLspFecT{} }
func (m *TeLspFecT) String() string { return proto.CompactTextString(m) }
func (*TeLspFecT) ProtoMessage()    {}
func (*TeLspFecT) Descriptor() ([]byte, []int) {
	return fileDescriptor_325e955e031246e6, []int{2}
}

func (m *TeLspFecT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeLspFecT.Unmarshal(m, b)
}
func (m *TeLspFecT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeLspFecT.Marshal(b, m, deterministic)
}
func (m *TeLspFecT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeLspFecT.Merge(m, src)
}
func (m *TeLspFecT) XXX_Size() int {
	return xxx_messageInfo_TeLspFecT.Size(m)
}
func (m *TeLspFecT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeLspFecT.DiscardUnknown(m)
}

var xxx_messageInfo_TeLspFecT proto.InternalMessageInfo

func (m *TeLspFecT) GetFecLspId() uint32 {
	if m != nil {
		return m.FecLspId
	}
	return 0
}

func (m *TeLspFecT) GetFecTunnelId() uint32 {
	if m != nil {
		return m.FecTunnelId
	}
	return 0
}

func (m *TeLspFecT) GetFecExtendedTunnelId() string {
	if m != nil {
		return m.FecExtendedTunnelId
	}
	return ""
}

func (m *TeLspFecT) GetFecSource() string {
	if m != nil {
		return m.FecSource
	}
	return ""
}

func (m *TeLspFecT) GetFecDestinationInfo() *TeLspFecCtypeDataT {
	if m != nil {
		return m.FecDestinationInfo
	}
	return nil
}

func (m *TeLspFecT) GetFecVrf() string {
	if m != nil {
		return m.FecVrf
	}
	return ""
}

type MplsTeTailBfdLspInfo struct {
	SignaledName         string     `protobuf:"bytes,50,opt,name=signaled_name,json=signaledName,proto3" json:"signaled_name,omitempty"`
	LspFec               *TeLspFecT `protobuf:"bytes,51,opt,name=lsp_fec,json=lspFec,proto3" json:"lsp_fec,omitempty"`
	BfdSessionState      string     `protobuf:"bytes,52,opt,name=bfd_session_state,json=bfdSessionState,proto3" json:"bfd_session_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MplsTeTailBfdLspInfo) Reset()         { *m = MplsTeTailBfdLspInfo{} }
func (m *MplsTeTailBfdLspInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeTailBfdLspInfo) ProtoMessage()    {}
func (*MplsTeTailBfdLspInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_325e955e031246e6, []int{3}
}

func (m *MplsTeTailBfdLspInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Unmarshal(m, b)
}
func (m *MplsTeTailBfdLspInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeTailBfdLspInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeTailBfdLspInfo.Merge(m, src)
}
func (m *MplsTeTailBfdLspInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeTailBfdLspInfo.Size(m)
}
func (m *MplsTeTailBfdLspInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeTailBfdLspInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeTailBfdLspInfo proto.InternalMessageInfo

func (m *MplsTeTailBfdLspInfo) GetSignaledName() string {
	if m != nil {
		return m.SignaledName
	}
	return ""
}

func (m *MplsTeTailBfdLspInfo) GetLspFec() *TeLspFecT {
	if m != nil {
		return m.LspFec
	}
	return nil
}

func (m *MplsTeTailBfdLspInfo) GetBfdSessionState() string {
	if m != nil {
		return m.BfdSessionState
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsTeTailBfdLspInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.tail_infos.tail_info.mpls_te_tail_bfd_lsp_info_KEYS")
	proto.RegisterType((*TeLspFecCtypeDataT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.tail_infos.tail_info.te_lsp_fec_ctype_data_t")
	proto.RegisterType((*TeLspFecT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.tail_infos.tail_info.te_lsp_fec_t")
	proto.RegisterType((*MplsTeTailBfdLspInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.bfd.tail_infos.tail_info.mpls_te_tail_bfd_lsp_info")
}

func init() { proto.RegisterFile("mpls_te_tail_bfd_lsp_info.proto", fileDescriptor_325e955e031246e6) }

var fileDescriptor_325e955e031246e6 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x26, 0xbb, 0xb6, 0xdd, 0x9e, 0xdd, 0xaa, 0x3b, 0xad, 0x36, 0xe2, 0xdf, 0x52, 0x11, 0x16,
	0x91, 0x08, 0xad, 0xde, 0x78, 0x27, 0xba, 0x62, 0x71, 0x11, 0x69, 0x44, 0xf0, 0x6a, 0x98, 0xce,
	0x9c, 0x91, 0x40, 0x9a, 0x0c, 0x99, 0x59, 0x59, 0x1f, 0xc0, 0x0b, 0x1f, 0xc2, 0xc7, 0xf1, 0xda,
	0x57, 0x92, 0x39, 0x93, 0xb4, 0xb1, 0xd8, 0x1b, 0xf1, 0x6e, 0xe6, 0x3b, 0x7f, 0xdf, 0xf9, 0xbe,
	0x49, 0xe0, 0xfe, 0xca, 0xe4, 0x96, 0x3b, 0xe4, 0x4e, 0x64, 0x39, 0x5f, 0x6a, 0xc5, 0x73, 0x6b,
	0x78, 0x56, 0xe8, 0x32, 0x31, 0x55, 0xe9, 0x4a, 0xf6, 0x5c, 0x66, 0x56, 0x96, 0x3c, 0x2b, 0x2d,
	0xbf, 0xac, 0x78, 0x93, 0x5d, 0x1a, 0xac, 0x92, 0xfa, 0x92, 0x2c, 0xb5, 0x4a, 0xa8, 0xdc, 0x97,
	0xd9, 0xcd, 0x71, 0xf2, 0x7d, 0x0f, 0xee, 0xed, 0xec, 0xcf, 0xdf, 0x9e, 0x7d, 0x4a, 0xd9, 0x43,
	0xb8, 0x6a, 0xcb, 0x8b, 0x4a, 0x22, 0x17, 0x4a, 0x55, 0x68, 0x6d, 0x1c, 0x9d, 0x44, 0xa7, 0xfd,
	0xc5, 0x20, 0xa0, 0x2f, 0x02, 0xc8, 0x6e, 0x43, 0xdf, 0x5d, 0x14, 0x05, 0xe6, 0x3c, 0x53, 0xf1,
	0xde, 0x49, 0x74, 0x3a, 0x58, 0x1c, 0x04, 0x60, 0xae, 0xd8, 0x0d, 0xe8, 0x52, 0x53, 0x15, 0xef,
	0x53, 0xa4, 0x93, 0x5b, 0x33, 0x57, 0xec, 0x09, 0x0c, 0x15, 0x5a, 0x97, 0x15, 0xc2, 0x65, 0x65,
	0xb1, 0xee, 0x7f, 0x85, 0xfa, 0xb3, 0x56, 0xa8, 0x19, 0xf2, 0x18, 0x18, 0x5e, 0x3a, 0x2c, 0x14,
	0x2a, 0xbe, 0x99, 0xd6, 0xa1, 0xfc, 0xeb, 0x4d, 0xe4, 0x43, 0x33, 0x75, 0x04, 0x1d, 0xe9, 0xbe,
	0x1a, 0x8c, 0xbb, 0x94, 0x10, 0x2e, 0x6c, 0x0c, 0x3d, 0x33, 0x5d, 0x11, 0x99, 0x1e, 0x91, 0xe9,
	0xfa, 0xeb, 0x5c, 0x4d, 0x7e, 0x44, 0x30, 0x76, 0x48, 0xdb, 0x6b, 0x94, 0x9c, 0xb2, 0xb9, 0x12,
	0x4e, 0x70, 0xe7, 0xb7, 0x5b, 0x63, 0xf5, 0xfe, 0x07, 0x1a, 0xe5, 0x4b, 0xea, 0x98, 0xc0, 0xd0,
	0x4c, 0x0d, 0x15, 0xb6, 0x38, 0x93, 0x08, 0xfd, 0xc5, 0xb1, 0x99, 0x9a, 0x73, 0x6b, 0x5e, 0x6d,
	0x02, 0xec, 0x19, 0x8c, 0x7d, 0xb3, 0xf6, 0xea, 0x0d, 0xa3, 0x20, 0xcf, 0x48, 0xa3, 0x6c, 0x15,
	0xbc, 0x0f, 0xfc, 0x7e, 0xee, 0xc1, 0x51, 0x8b, 0x9f, 0x63, 0x77, 0x00, 0xfc, 0xa1, 0x56, 0x36,
	0x0a, 0x9a, 0x6b, 0x94, 0xe7, 0x24, 0xee, 0x04, 0x06, 0x94, 0xb6, 0x65, 0xca, 0xa1, 0x46, 0xb9,
	0x56, 0x68, 0x06, 0x37, 0x7d, 0xce, 0x5f, 0x34, 0xdd, 0x27, 0xf2, 0x43, 0x8d, 0xf2, 0x6c, 0x5b,
	0xd6, 0xbb, 0x61, 0x6c, 0xb0, 0xbf, 0x36, 0xcb, 0xab, 0x93, 0x12, 0xc0, 0xbe, 0x45, 0x30, 0xda,
	0x5e, 0xcf, 0xbf, 0x26, 0xb2, 0xe9, 0x70, 0x9a, 0x26, 0xff, 0xfe, 0x5c, 0x93, 0x1d, 0xf6, 0x2c,
	0xd8, 0x9f, 0x82, 0xcd, 0x0b, 0x5d, 0x7a, 0x9f, 0x7d, 0xde, 0x97, 0x4a, 0xd7, 0xfe, 0x77, 0x35,
	0xca, 0x8f, 0x95, 0x9e, 0xfc, 0x8a, 0xe0, 0xd6, 0xce, 0x37, 0xcf, 0x1e, 0xc0, 0xc0, 0x66, 0x9f,
	0x0b, 0x91, 0xa3, 0xe2, 0x85, 0x58, 0x61, 0x3c, 0xa5, 0xe2, 0xa3, 0x06, 0x7c, 0x27, 0x56, 0xc8,
	0x04, 0xf4, 0x6a, 0x1e, 0xf1, 0x8c, 0xb6, 0x7a, 0xf3, 0x9f, 0xb6, 0x72, 0x0b, 0xff, 0xa1, 0xbc,
	0x46, 0xc9, 0x1e, 0xc1, 0xb1, 0xe7, 0x65, 0xd1, 0x5a, 0xaf, 0xa0, 0x75, 0xc2, 0x61, 0xfc, 0x94,
	0xb8, 0x5c, 0x5b, 0x6a, 0x95, 0x06, 0x3c, 0xf5, 0xf0, 0xb2, 0x4b, 0x3f, 0x82, 0xd9, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x77, 0x44, 0x34, 0x2b, 0x04, 0x00, 0x00,
}
