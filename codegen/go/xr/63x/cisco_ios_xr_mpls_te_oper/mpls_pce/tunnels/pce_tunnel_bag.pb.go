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
// source: pce_tunnel_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_pce_tunnels

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

type PceTunnelBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceTunnelBag_KEYS) Reset()         { *m = PceTunnelBag_KEYS{} }
func (m *PceTunnelBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PceTunnelBag_KEYS) ProtoMessage()    {}
func (*PceTunnelBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0850897a5e49f5fc, []int{0}
}

func (m *PceTunnelBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTunnelBag_KEYS.Unmarshal(m, b)
}
func (m *PceTunnelBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTunnelBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PceTunnelBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTunnelBag_KEYS.Merge(m, src)
}
func (m *PceTunnelBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PceTunnelBag_KEYS.Size(m)
}
func (m *PceTunnelBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTunnelBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PceTunnelBag_KEYS proto.InternalMessageInfo

type PceTunnelBag struct {
	SourceAddress        string   `protobuf:"bytes,50,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress   string   `protobuf:"bytes,51,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	IsCurrPathOptionPce  bool     `protobuf:"varint,52,opt,name=is_curr_path_option_pce,json=isCurrPathOptionPce,proto3" json:"is_curr_path_option_pce,omitempty"`
	CurrPathOptionIndex  uint32   `protobuf:"varint,53,opt,name=curr_path_option_index,json=currPathOptionIndex,proto3" json:"curr_path_option_index,omitempty"`
	ConfiguredPceAddress string   `protobuf:"bytes,54,opt,name=configured_pce_address,json=configuredPceAddress,proto3" json:"configured_pce_address,omitempty"`
	SenderPceAddress     string   `protobuf:"bytes,55,opt,name=sender_pce_address,json=senderPceAddress,proto3" json:"sender_pce_address,omitempty"`
	PathState            string   `protobuf:"bytes,56,opt,name=path_state,json=pathState,proto3" json:"path_state,omitempty"`
	TunnelState          string   `protobuf:"bytes,57,opt,name=tunnel_state,json=tunnelState,proto3" json:"tunnel_state,omitempty"`
	AdminWeight          uint32   `protobuf:"varint,58,opt,name=admin_weight,json=adminWeight,proto3" json:"admin_weight,omitempty"`
	HopCount             uint32   `protobuf:"varint,59,opt,name=hop_count,json=hopCount,proto3" json:"hop_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PceTunnelBag) Reset()         { *m = PceTunnelBag{} }
func (m *PceTunnelBag) String() string { return proto.CompactTextString(m) }
func (*PceTunnelBag) ProtoMessage()    {}
func (*PceTunnelBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0850897a5e49f5fc, []int{1}
}

func (m *PceTunnelBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PceTunnelBag.Unmarshal(m, b)
}
func (m *PceTunnelBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PceTunnelBag.Marshal(b, m, deterministic)
}
func (m *PceTunnelBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PceTunnelBag.Merge(m, src)
}
func (m *PceTunnelBag) XXX_Size() int {
	return xxx_messageInfo_PceTunnelBag.Size(m)
}
func (m *PceTunnelBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PceTunnelBag.DiscardUnknown(m)
}

var xxx_messageInfo_PceTunnelBag proto.InternalMessageInfo

func (m *PceTunnelBag) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PceTunnelBag) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *PceTunnelBag) GetIsCurrPathOptionPce() bool {
	if m != nil {
		return m.IsCurrPathOptionPce
	}
	return false
}

func (m *PceTunnelBag) GetCurrPathOptionIndex() uint32 {
	if m != nil {
		return m.CurrPathOptionIndex
	}
	return 0
}

func (m *PceTunnelBag) GetConfiguredPceAddress() string {
	if m != nil {
		return m.ConfiguredPceAddress
	}
	return ""
}

func (m *PceTunnelBag) GetSenderPceAddress() string {
	if m != nil {
		return m.SenderPceAddress
	}
	return ""
}

func (m *PceTunnelBag) GetPathState() string {
	if m != nil {
		return m.PathState
	}
	return ""
}

func (m *PceTunnelBag) GetTunnelState() string {
	if m != nil {
		return m.TunnelState
	}
	return ""
}

func (m *PceTunnelBag) GetAdminWeight() uint32 {
	if m != nil {
		return m.AdminWeight
	}
	return 0
}

func (m *PceTunnelBag) GetHopCount() uint32 {
	if m != nil {
		return m.HopCount
	}
	return 0
}

func init() {
	proto.RegisterType((*PceTunnelBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.tunnels.pce_tunnel_bag_KEYS")
	proto.RegisterType((*PceTunnelBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce.tunnels.pce_tunnel_bag")
}

func init() { proto.RegisterFile("pce_tunnel_bag.proto", fileDescriptor_0850897a5e49f5fc) }

var fileDescriptor_0850897a5e49f5fc = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x6b, 0x2a, 0x31,
	0x14, 0x86, 0x91, 0x0b, 0x17, 0x8d, 0x57, 0xb9, 0x64, 0xbc, 0xde, 0x81, 0x52, 0xb0, 0x42, 0x41,
	0x4a, 0xb1, 0x50, 0xed, 0xf7, 0xaa, 0x48, 0x17, 0xa5, 0x8b, 0x8a, 0x2e, 0x4a, 0x57, 0x87, 0x98,
	0x9c, 0x3a, 0x01, 0x4d, 0x42, 0x92, 0xa1, 0xfe, 0xb5, 0xfe, 0xbb, 0x32, 0x67, 0x5a, 0x75, 0xe8,
	0xf2, 0x3c, 0xef, 0xfb, 0xcc, 0x9c, 0x33, 0x0c, 0xeb, 0x38, 0x89, 0x10, 0x73, 0x63, 0x70, 0x05,
	0x0b, 0xb1, 0x1c, 0x3a, 0x6f, 0xa3, 0xe5, 0x27, 0x52, 0x07, 0x69, 0x41, 0xdb, 0x00, 0x1b, 0x0f,
	0x6b, 0xb7, 0x0a, 0x10, 0x11, 0xac, 0x43, 0x3f, 0xa4, 0xc1, 0x49, 0x1c, 0x96, 0x52, 0xe8, 0xff,
	0x63, 0x49, 0xf5, 0x19, 0xf0, 0xf4, 0xf0, 0x3a, 0xef, 0x7f, 0xfc, 0x62, 0xed, 0x2a, 0xe7, 0xc7,
	0xac, 0x1d, 0x6c, 0xee, 0x25, 0x82, 0x50, 0xca, 0x63, 0x08, 0xe9, 0x79, 0xaf, 0x36, 0x68, 0xcc,
	0x5a, 0x25, 0xbd, 0x2f, 0x21, 0x3f, 0x63, 0x89, 0xc2, 0x10, 0xb5, 0x11, 0x51, 0x5b, 0xb3, 0xed,
	0x8e, 0xa8, 0xcb, 0xf7, 0xa2, 0x6f, 0x61, 0xcc, 0xfe, 0xeb, 0x00, 0x32, 0xf7, 0x1e, 0x9c, 0x88,
	0x19, 0x58, 0x47, 0xa2, 0x93, 0x98, 0x8e, 0x7b, 0xb5, 0x41, 0x7d, 0x96, 0xe8, 0x30, 0xc9, 0xbd,
	0x9f, 0x8a, 0x98, 0x3d, 0x53, 0x36, 0x95, 0xc8, 0x47, 0xac, 0xfb, 0x43, 0xd1, 0x46, 0xe1, 0x26,
	0xbd, 0xe8, 0xd5, 0x06, 0xad, 0x59, 0x22, 0x2b, 0xca, 0x63, 0x11, 0xf1, 0x31, 0xeb, 0x4a, 0x6b,
	0xde, 0xf4, 0x32, 0xf7, 0xa8, 0x8a, 0x37, 0x6c, 0xd7, 0xbb, 0xa4, 0xf5, 0x3a, 0xbb, 0x74, 0xba,
	0xbb, 0xe8, 0x94, 0xf1, 0x80, 0x46, 0xa1, 0xaf, 0x18, 0x57, 0x64, 0xfc, 0x2d, 0x93, 0xbd, 0xf6,
	0x21, 0x63, 0xb4, 0x53, 0x88, 0x22, 0x62, 0x7a, 0x4d, 0xad, 0x46, 0x41, 0xe6, 0x05, 0xe0, 0x47,
	0xec, 0xcf, 0xd7, 0x37, 0x2d, 0x0b, 0x37, 0x54, 0x68, 0x96, 0x6c, 0x5b, 0x11, 0x6a, 0xad, 0x0d,
	0xbc, 0xa3, 0x5e, 0x66, 0x31, 0xbd, 0xa5, 0x83, 0x9a, 0xc4, 0x5e, 0x08, 0xf1, 0x03, 0xd6, 0xc8,
	0xac, 0x03, 0x69, 0x73, 0x13, 0xd3, 0x3b, 0xca, 0xeb, 0x99, 0x75, 0x93, 0x62, 0x5e, 0xfc, 0xa6,
	0xbf, 0x60, 0xf4, 0x19, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x31, 0x75, 0xc2, 0x1d, 0x02, 0x00, 0x00,
}
