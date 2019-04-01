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

package cisco_ios_xr_mpls_te_oper_mpls_pce_stdby_tunnels

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
	proto.RegisterType((*PceTunnelBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce_stdby.tunnels.pce_tunnel_bag_KEYS")
	proto.RegisterType((*PceTunnelBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_pce_stdby.tunnels.pce_tunnel_bag")
}

func init() { proto.RegisterFile("pce_tunnel_bag.proto", fileDescriptor_0850897a5e49f5fc) }

var fileDescriptor_0850897a5e49f5fc = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0x29, 0x0f, 0x3c, 0xb4, 0xa9, 0x2d, 0x92, 0xad, 0x75, 0x41, 0x84, 0x5a, 0x10, 0x7a,
	0x90, 0x2a, 0xb6, 0xbe, 0x9f, 0xa4, 0x78, 0x10, 0x0f, 0x96, 0xf6, 0x20, 0x9e, 0x86, 0x6d, 0x32,
	0x76, 0x03, 0x6d, 0x12, 0x92, 0x2c, 0xd6, 0x8f, 0xe6, 0xb7, 0x93, 0x9d, 0xb5, 0x2f, 0x8b, 0xc7,
	0xfc, 0xfe, 0xff, 0xdf, 0x32, 0x33, 0x2c, 0x6b, 0x59, 0x81, 0x10, 0x32, 0xad, 0x71, 0x01, 0xb3,
	0x64, 0xde, 0xb7, 0xce, 0x04, 0xc3, 0x2f, 0x84, 0xf2, 0xc2, 0x80, 0x32, 0x1e, 0x56, 0x0e, 0x96,
	0x76, 0xe1, 0x21, 0x20, 0x18, 0x8b, 0xae, 0x4f, 0x8f, 0x5c, 0xf2, 0x41, 0xce, 0xbe, 0xfa, 0x85,
	0xea, 0xbb, 0x07, 0x2c, 0x2a, 0x7f, 0x09, 0x5e, 0x9e, 0xde, 0xa7, 0xdd, 0xef, 0x7f, 0xac, 0x59,
	0xe6, 0xfc, 0x94, 0x35, 0xbd, 0xc9, 0x9c, 0x40, 0x48, 0xa4, 0x74, 0xe8, 0x7d, 0x7c, 0xd9, 0xa9,
	0xf4, 0x6a, 0x93, 0x46, 0x41, 0x1f, 0x0b, 0xc8, 0xcf, 0x59, 0x24, 0xd1, 0x07, 0xa5, 0x93, 0xa0,
	0x8c, 0xde, 0x74, 0x07, 0xd4, 0xe5, 0x3b, 0xd1, 0x5a, 0x18, 0xb2, 0x43, 0xe5, 0x41, 0x64, 0xce,
	0x81, 0x4d, 0x42, 0x0a, 0xc6, 0x92, 0x68, 0x05, 0xc6, 0xc3, 0x4e, 0xa5, 0x57, 0x9d, 0x44, 0xca,
	0x8f, 0x32, 0xe7, 0xc6, 0x49, 0x48, 0x5f, 0x29, 0x1b, 0x0b, 0xe4, 0x03, 0xd6, 0xfe, 0xa3, 0x28,
	0x2d, 0x71, 0x15, 0x5f, 0x75, 0x2a, 0xbd, 0xc6, 0x24, 0x12, 0x25, 0xe5, 0x39, 0x8f, 0xf8, 0x90,
	0xb5, 0x85, 0xd1, 0x1f, 0x6a, 0x9e, 0x39, 0x94, 0x74, 0x8c, 0xf5, 0x78, 0xd7, 0x34, 0x5e, 0x6b,
	0x9b, 0x8e, 0xb7, 0x1b, 0x9d, 0x31, 0xee, 0x51, 0x4b, 0x74, 0x25, 0xe3, 0x86, 0x8c, 0xfd, 0x22,
	0xd9, 0x69, 0x1f, 0x33, 0x46, 0x33, 0xf9, 0x90, 0x04, 0x8c, 0x6f, 0xa9, 0x55, 0xcb, 0xc9, 0x34,
	0x07, 0xfc, 0x84, 0xed, 0xfd, 0xde, 0xb4, 0x28, 0xdc, 0x51, 0xa1, 0x5e, 0xb0, 0x4d, 0x25, 0x91,
	0x4b, 0xa5, 0xe1, 0x13, 0xd5, 0x3c, 0x0d, 0xf1, 0x3d, 0x2d, 0x54, 0x27, 0xf6, 0x46, 0x88, 0x1f,
	0xb1, 0x5a, 0x6a, 0x2c, 0x08, 0x93, 0xe9, 0x10, 0x3f, 0x50, 0x5e, 0x4d, 0x8d, 0x1d, 0xe5, 0xef,
	0xd9, 0x7f, 0xfa, 0x17, 0x06, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x53, 0xf6, 0x0e, 0x23,
	0x02, 0x00, 0x00,
}
