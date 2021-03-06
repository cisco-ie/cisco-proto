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
// source: mpls_te_lsp_wrap_protection_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_wrap_protection_protected_lsps_protected_lsp

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

type MplsTeLspWrapProtectionBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,2,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,3,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,4,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	LspId                uint32   `protobuf:"varint,5,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,6,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	SubGroupOriginator   string   `protobuf:"bytes,7,opt,name=sub_group_originator,json=subGroupOriginator,proto3" json:"sub_group_originator,omitempty"`
	P2MpId               uint32   `protobuf:"varint,8,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	SessionType          string   `protobuf:"bytes,9,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeLspWrapProtectionBag_KEYS) Reset()         { *m = MplsTeLspWrapProtectionBag_KEYS{} }
func (m *MplsTeLspWrapProtectionBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeLspWrapProtectionBag_KEYS) ProtoMessage()    {}
func (*MplsTeLspWrapProtectionBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a64665727e71db0, []int{0}
}

func (m *MplsTeLspWrapProtectionBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeLspWrapProtectionBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeLspWrapProtectionBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS.Merge(m, src)
}
func (m *MplsTeLspWrapProtectionBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS.Size(m)
}
func (m *MplsTeLspWrapProtectionBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeLspWrapProtectionBag_KEYS proto.InternalMessageInfo

func (m *MplsTeLspWrapProtectionBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetLspId() uint32 {
	if m != nil {
		return m.LspId
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetSubGroupOriginator() string {
	if m != nil {
		return m.SubGroupOriginator
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag_KEYS) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

type MplsTeLspWrapProtectionBag struct {
	SourceAddressXr         string   `protobuf:"bytes,50,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	DestinationAddressXr    string   `protobuf:"bytes,51,opt,name=destination_address_xr,json=destinationAddressXr,proto3" json:"destination_address_xr,omitempty"`
	TunnelIdXr              uint32   `protobuf:"varint,52,opt,name=tunnel_id_xr,json=tunnelIdXr,proto3" json:"tunnel_id_xr,omitempty"`
	ExtendedTunnelIdXr      string   `protobuf:"bytes,53,opt,name=extended_tunnel_id_xr,json=extendedTunnelIdXr,proto3" json:"extended_tunnel_id_xr,omitempty"`
	TunnelName              string   `protobuf:"bytes,54,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TunnelInstance          uint32   `protobuf:"varint,55,opt,name=tunnel_instance,json=tunnelInstance,proto3" json:"tunnel_instance,omitempty"`
	LspWrapProtectionEnable bool     `protobuf:"varint,56,opt,name=lsp_wrap_protection_enable,json=lspWrapProtectionEnable,proto3" json:"lsp_wrap_protection_enable,omitempty"`
	LspWrapProtectionLabel  uint32   `protobuf:"varint,57,opt,name=lsp_wrap_protection_label,json=lspWrapProtectionLabel,proto3" json:"lsp_wrap_protection_label,omitempty"`
	ReverseEgressInterface  string   `protobuf:"bytes,58,opt,name=reverse_egress_interface,json=reverseEgressInterface,proto3" json:"reverse_egress_interface,omitempty"`
	ReverseLspLabel         uint32   `protobuf:"varint,59,opt,name=reverse_lsp_label,json=reverseLspLabel,proto3" json:"reverse_lsp_label,omitempty"`
	LspWrapProtectionState  string   `protobuf:"bytes,60,opt,name=lsp_wrap_protection_state,json=lspWrapProtectionState,proto3" json:"lsp_wrap_protection_state,omitempty"`
	SubGroupOriginalId      string   `protobuf:"bytes,61,opt,name=sub_group_original_id,json=subGroupOriginalId,proto3" json:"sub_group_original_id,omitempty"`
	SubGroupIdXr            uint32   `protobuf:"varint,62,opt,name=sub_group_id_xr,json=subGroupIdXr,proto3" json:"sub_group_id_xr,omitempty"`
	P2MpIdXr                uint32   `protobuf:"varint,63,opt,name=p2mp_id_xr,json=p2mpIdXr,proto3" json:"p2mp_id_xr,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *MplsTeLspWrapProtectionBag) Reset()         { *m = MplsTeLspWrapProtectionBag{} }
func (m *MplsTeLspWrapProtectionBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeLspWrapProtectionBag) ProtoMessage()    {}
func (*MplsTeLspWrapProtectionBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a64665727e71db0, []int{1}
}

func (m *MplsTeLspWrapProtectionBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag.Unmarshal(m, b)
}
func (m *MplsTeLspWrapProtectionBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag.Marshal(b, m, deterministic)
}
func (m *MplsTeLspWrapProtectionBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeLspWrapProtectionBag.Merge(m, src)
}
func (m *MplsTeLspWrapProtectionBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeLspWrapProtectionBag.Size(m)
}
func (m *MplsTeLspWrapProtectionBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeLspWrapProtectionBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeLspWrapProtectionBag proto.InternalMessageInfo

func (m *MplsTeLspWrapProtectionBag) GetSourceAddressXr() string {
	if m != nil {
		return m.SourceAddressXr
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetDestinationAddressXr() string {
	if m != nil {
		return m.DestinationAddressXr
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetTunnelIdXr() uint32 {
	if m != nil {
		return m.TunnelIdXr
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag) GetExtendedTunnelIdXr() string {
	if m != nil {
		return m.ExtendedTunnelIdXr
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetTunnelInstance() uint32 {
	if m != nil {
		return m.TunnelInstance
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag) GetLspWrapProtectionEnable() bool {
	if m != nil {
		return m.LspWrapProtectionEnable
	}
	return false
}

func (m *MplsTeLspWrapProtectionBag) GetLspWrapProtectionLabel() uint32 {
	if m != nil {
		return m.LspWrapProtectionLabel
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag) GetReverseEgressInterface() string {
	if m != nil {
		return m.ReverseEgressInterface
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetReverseLspLabel() uint32 {
	if m != nil {
		return m.ReverseLspLabel
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag) GetLspWrapProtectionState() string {
	if m != nil {
		return m.LspWrapProtectionState
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetSubGroupOriginalId() string {
	if m != nil {
		return m.SubGroupOriginalId
	}
	return ""
}

func (m *MplsTeLspWrapProtectionBag) GetSubGroupIdXr() uint32 {
	if m != nil {
		return m.SubGroupIdXr
	}
	return 0
}

func (m *MplsTeLspWrapProtectionBag) GetP2MpIdXr() uint32 {
	if m != nil {
		return m.P2MpIdXr
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsTeLspWrapProtectionBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.wrap_protection.protected_lsps.protected_lsp.mpls_te_lsp_wrap_protection_bag_KEYS")
	proto.RegisterType((*MplsTeLspWrapProtectionBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.wrap_protection.protected_lsps.protected_lsp.mpls_te_lsp_wrap_protection_bag")
}

func init() {
	proto.RegisterFile("mpls_te_lsp_wrap_protection_bag.proto", fileDescriptor_7a64665727e71db0)
}

var fileDescriptor_7a64665727e71db0 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x55, 0xfa, 0x7d, 0x49, 0x93, 0x69, 0xda, 0xc0, 0xd2, 0xa6, 0xcb, 0x8f, 0xd4, 0x50, 0x11,
	0x11, 0x21, 0x14, 0xa0, 0x2d, 0xd0, 0x52, 0x7e, 0xc4, 0x45, 0x84, 0x22, 0x2a, 0x40, 0x69, 0x05,
	0xe6, 0x6a, 0xb5, 0xf6, 0x0e, 0x91, 0x25, 0xc7, 0x5e, 0xed, 0x3a, 0x90, 0x3c, 0x17, 0x2f, 0xc6,
	0x23, 0xa0, 0x1d, 0xdb, 0x49, 0xd3, 0x04, 0xf5, 0xce, 0x3b, 0xe7, 0xcc, 0x9c, 0xd9, 0xb3, 0x33,
	0x86, 0xf6, 0x48, 0x47, 0x56, 0xa4, 0x28, 0x22, 0xab, 0xc5, 0x2f, 0x23, 0xb5, 0xd0, 0x26, 0x49,
	0x31, 0x48, 0xc3, 0x24, 0x16, 0xbe, 0x1c, 0x76, 0xdd, 0x31, 0x61, 0x5f, 0x83, 0xd0, 0x06, 0x89,
	0x08, 0x13, 0x2b, 0x26, 0x46, 0x14, 0x39, 0x89, 0x46, 0xd3, 0x2d, 0x0e, 0x36, 0x95, 0xb1, 0xf2,
	0xa7, 0xdd, 0x2b, 0x45, 0xba, 0xf9, 0x27, 0x2a, 0x27, 0x61, 0x17, 0x8f, 0xfb, 0x7f, 0xd6, 0xe0,
	0xc1, 0x35, 0x1d, 0x88, 0x8f, 0xbd, 0xef, 0xe7, 0xac, 0x0d, 0x5b, 0x36, 0x19, 0x9b, 0x00, 0x85,
	0x54, 0xca, 0xa0, 0xb5, 0xbc, 0xd4, 0x2a, 0x75, 0x6a, 0x83, 0xcd, 0x2c, 0xfa, 0x3e, 0x0b, 0xb2,
	0x27, 0x70, 0x4b, 0xa1, 0x4d, 0xc3, 0x58, 0x52, 0x7a, 0xc1, 0x5d, 0x23, 0x2e, 0xbb, 0x04, 0x15,
	0x09, 0x77, 0xa1, 0x96, 0x8e, 0xe3, 0x18, 0x23, 0x11, 0x2a, 0xfe, 0x5f, 0xab, 0xd4, 0xd9, 0x1c,
	0x54, 0xb3, 0x40, 0x5f, 0xb1, 0xc7, 0xc0, 0x70, 0x92, 0x62, 0xac, 0x50, 0x89, 0x39, 0xeb, 0x7f,
	0x2a, 0x76, 0xa3, 0x40, 0x2e, 0x0a, 0xf6, 0x0e, 0x54, 0xdc, 0x15, 0x42, 0xc5, 0xcb, 0x54, 0xa7,
	0x1c, 0x59, 0xdd, 0x57, 0xac, 0x05, 0x75, 0x3b, 0xf6, 0xc5, 0xd0, 0x24, 0x63, 0x02, 0x2b, 0x04,
	0x82, 0x1d, 0xfb, 0x1f, 0x5c, 0xa8, 0xaf, 0xd8, 0x53, 0xd8, 0x9e, 0x33, 0x12, 0x13, 0x0e, 0x5d,
	0x8f, 0x89, 0xe1, 0xeb, 0x59, 0xd7, 0x05, 0xf3, 0xf3, 0x0c, 0x61, 0xbb, 0xb0, 0xae, 0x0f, 0x46,
	0x54, 0xae, 0x4a, 0xe5, 0x2a, 0xee, 0xd8, 0x57, 0xec, 0x3e, 0xd4, 0x2d, 0x5a, 0xeb, 0xee, 0x9e,
	0x4e, 0x35, 0xf2, 0x1a, 0x95, 0xd8, 0xc8, 0x63, 0x17, 0x53, 0x8d, 0xfb, 0xbf, 0xcb, 0xb0, 0x77,
	0x8d, 0xe5, 0xec, 0x11, 0xdc, 0x5c, 0x74, 0x5b, 0x4c, 0x0c, 0x3f, 0xa0, 0x5a, 0x8d, 0x05, 0xc3,
	0x3d, 0xc3, 0x8e, 0xa0, 0xb9, 0xc2, 0x72, 0x97, 0x70, 0x48, 0x09, 0xdb, 0xcb, 0xae, 0x7b, 0xc6,
	0xb9, 0x32, 0x73, 0xd4, 0x71, 0x8f, 0x32, 0x57, 0x0a, 0xeb, 0x3d, 0xc3, 0x9e, 0xc1, 0xce, 0xb2,
	0xf9, 0x8e, 0xfa, 0x3c, 0xb3, 0xe5, 0xaa, 0xff, 0x9e, 0x61, 0x7b, 0xb0, 0x91, 0x33, 0x63, 0x39,
	0x42, 0xfe, 0x82, 0x88, 0x79, 0xcd, 0x4f, 0x72, 0x84, 0xec, 0x21, 0x34, 0x8a, 0x52, 0xb1, 0x9b,
	0xd7, 0x00, 0xf9, 0x4b, 0x12, 0xde, 0xca, 0x85, 0xf3, 0x28, 0x3b, 0x85, 0x3b, 0xab, 0xbc, 0xc1,
	0x58, 0xfa, 0x11, 0xf2, 0xe3, 0x56, 0xa9, 0x53, 0x1d, 0xec, 0x46, 0x56, 0x7f, 0x33, 0x52, 0x7f,
	0x99, 0xe1, 0x3d, 0x82, 0xd9, 0x09, 0xdc, 0x5e, 0x95, 0x1c, 0x49, 0x1f, 0x23, 0x7e, 0x42, 0x7a,
	0xcd, 0xa5, 0xdc, 0x33, 0x87, 0xb2, 0x63, 0xe0, 0x06, 0x7f, 0xa2, 0xb1, 0x28, 0x70, 0x48, 0x3e,
	0x86, 0x71, 0x8a, 0xe6, 0x87, 0x0c, 0x90, 0xbf, 0xa2, 0xeb, 0x34, 0x73, 0xbc, 0x47, 0x70, 0xbf,
	0x40, 0xdd, 0x93, 0x15, 0x99, 0x4e, 0x3c, 0x13, 0x3b, 0x25, 0xb1, 0x46, 0x0e, 0x9c, 0x59, 0x9d,
	0xa9, 0xfc, 0xa3, 0x41, 0x9b, 0xca, 0x14, 0xf9, 0xeb, 0x4c, 0x66, 0xa9, 0xc1, 0x73, 0x87, 0xba,
	0x57, 0x59, 0x9a, 0x55, 0xda, 0x8a, 0x37, 0x2b, 0x87, 0xd5, 0xed, 0x45, 0x1b, 0x1a, 0x97, 0x17,
	0xc0, 0x3d, 0xe1, 0x5b, 0xea, 0xab, 0x3e, 0xdf, 0x01, 0xcf, 0xb0, 0x7b, 0x00, 0xf9, 0x4c, 0x3b,
	0xc6, 0xbb, 0x6c, 0x15, 0xb3, 0xb1, 0xf6, 0x8c, 0x5f, 0xa1, 0xff, 0xd0, 0xe1, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x89, 0x6f, 0x3e, 0xb0, 0x04, 0x00, 0x00,
}
