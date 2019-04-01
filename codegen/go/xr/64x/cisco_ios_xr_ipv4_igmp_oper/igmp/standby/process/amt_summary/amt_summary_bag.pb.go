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
// source: amt_summary_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_process_amt_summary

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

type AmtSummaryBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmtSummaryBag_KEYS) Reset()         { *m = AmtSummaryBag_KEYS{} }
func (m *AmtSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AmtSummaryBag_KEYS) ProtoMessage()    {}
func (*AmtSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_692236ad8513a9eb, []int{0}
}

func (m *AmtSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmtSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *AmtSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmtSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AmtSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmtSummaryBag_KEYS.Merge(m, src)
}
func (m *AmtSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AmtSummaryBag_KEYS.Size(m)
}
func (m *AmtSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AmtSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AmtSummaryBag_KEYS proto.InternalMessageInfo

type AmtSummaryBag struct {
	AnycastPrefix           string   `protobuf:"bytes,50,opt,name=anycast_prefix,json=anycastPrefix,proto3" json:"anycast_prefix,omitempty"`
	PrefixLength            uint32   `protobuf:"varint,51,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	RelayAddress            string   `protobuf:"bytes,52,opt,name=relay_address,json=relayAddress,proto3" json:"relay_address,omitempty"`
	Mtu                     uint32   `protobuf:"varint,53,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Tos                     uint32   `protobuf:"varint,54,opt,name=tos,proto3" json:"tos,omitempty"`
	Ttl                     uint32   `protobuf:"varint,55,opt,name=ttl,proto3" json:"ttl,omitempty"`
	QueryInterval           uint32   `protobuf:"varint,56,opt,name=query_interval,json=queryInterval,proto3" json:"query_interval,omitempty"`
	GatewayCount            uint32   `protobuf:"varint,57,opt,name=gateway_count,json=gatewayCount,proto3" json:"gateway_count,omitempty"`
	MaxGateway              uint32   `protobuf:"varint,58,opt,name=max_gateway,json=maxGateway,proto3" json:"max_gateway,omitempty"`
	TunnelCount             uint32   `protobuf:"varint,59,opt,name=tunnel_count,json=tunnelCount,proto3" json:"tunnel_count,omitempty"`
	TunnelConfiguredMaximum uint32   `protobuf:"varint,60,opt,name=tunnel_configured_maximum,json=tunnelConfiguredMaximum,proto3" json:"tunnel_configured_maximum,omitempty"`
	IsAclConfigured         bool     `protobuf:"varint,61,opt,name=is_acl_configured,json=isAclConfigured,proto3" json:"is_acl_configured,omitempty"`
	IsGatewaySimulation     bool     `protobuf:"varint,62,opt,name=is_gateway_simulation,json=isGatewaySimulation,proto3" json:"is_gateway_simulation,omitempty"`
	IsOuOfResource          bool     `protobuf:"varint,63,opt,name=is_ou_of_resource,json=isOuOfResource,proto3" json:"is_ou_of_resource,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *AmtSummaryBag) Reset()         { *m = AmtSummaryBag{} }
func (m *AmtSummaryBag) String() string { return proto.CompactTextString(m) }
func (*AmtSummaryBag) ProtoMessage()    {}
func (*AmtSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_692236ad8513a9eb, []int{1}
}

func (m *AmtSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmtSummaryBag.Unmarshal(m, b)
}
func (m *AmtSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmtSummaryBag.Marshal(b, m, deterministic)
}
func (m *AmtSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmtSummaryBag.Merge(m, src)
}
func (m *AmtSummaryBag) XXX_Size() int {
	return xxx_messageInfo_AmtSummaryBag.Size(m)
}
func (m *AmtSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AmtSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_AmtSummaryBag proto.InternalMessageInfo

func (m *AmtSummaryBag) GetAnycastPrefix() string {
	if m != nil {
		return m.AnycastPrefix
	}
	return ""
}

func (m *AmtSummaryBag) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *AmtSummaryBag) GetRelayAddress() string {
	if m != nil {
		return m.RelayAddress
	}
	return ""
}

func (m *AmtSummaryBag) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *AmtSummaryBag) GetTos() uint32 {
	if m != nil {
		return m.Tos
	}
	return 0
}

func (m *AmtSummaryBag) GetTtl() uint32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *AmtSummaryBag) GetQueryInterval() uint32 {
	if m != nil {
		return m.QueryInterval
	}
	return 0
}

func (m *AmtSummaryBag) GetGatewayCount() uint32 {
	if m != nil {
		return m.GatewayCount
	}
	return 0
}

func (m *AmtSummaryBag) GetMaxGateway() uint32 {
	if m != nil {
		return m.MaxGateway
	}
	return 0
}

func (m *AmtSummaryBag) GetTunnelCount() uint32 {
	if m != nil {
		return m.TunnelCount
	}
	return 0
}

func (m *AmtSummaryBag) GetTunnelConfiguredMaximum() uint32 {
	if m != nil {
		return m.TunnelConfiguredMaximum
	}
	return 0
}

func (m *AmtSummaryBag) GetIsAclConfigured() bool {
	if m != nil {
		return m.IsAclConfigured
	}
	return false
}

func (m *AmtSummaryBag) GetIsGatewaySimulation() bool {
	if m != nil {
		return m.IsGatewaySimulation
	}
	return false
}

func (m *AmtSummaryBag) GetIsOuOfResource() bool {
	if m != nil {
		return m.IsOuOfResource
	}
	return false
}

func init() {
	proto.RegisterType((*AmtSummaryBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.process.amt_summary.amt_summary_bag_KEYS")
	proto.RegisterType((*AmtSummaryBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.process.amt_summary.amt_summary_bag")
}

func init() { proto.RegisterFile("amt_summary_bag.proto", fileDescriptor_692236ad8513a9eb) }

var fileDescriptor_692236ad8513a9eb = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x6f, 0x13, 0x31,
	0x10, 0xc5, 0x15, 0x81, 0x10, 0xb8, 0x49, 0x0b, 0x86, 0x82, 0x39, 0x11, 0x8a, 0x90, 0x02, 0x87,
	0x1c, 0xda, 0xf2, 0xaf, 0x14, 0x50, 0x85, 0x10, 0x42, 0x80, 0x8a, 0xd2, 0x13, 0xa7, 0xd1, 0x64,
	0xe3, 0x5d, 0x2c, 0xad, 0xed, 0xc5, 0x33, 0x2e, 0xbb, 0x9f, 0x82, 0xaf, 0x5c, 0xad, 0xbd, 0xa9,
	0xa2, 0xdc, 0xac, 0xdf, 0xfb, 0xcd, 0xd3, 0x3b, 0x58, 0xec, 0xa3, 0x65, 0xa0, 0x68, 0x2d, 0x86,
	0x0e, 0x96, 0x58, 0xcd, 0x9b, 0xe0, 0xd9, 0xcb, 0xd3, 0xc2, 0x50, 0xe1, 0xc1, 0x78, 0x82, 0x36,
	0x80, 0x69, 0x2e, 0x8f, 0xc1, 0x54, 0xb6, 0x01, 0xdf, 0xe8, 0x30, 0xef, 0x5f, 0x73, 0x62, 0x74,
	0xab, 0x65, 0xd7, 0xfb, 0x85, 0x26, 0x9a, 0x6f, 0xf4, 0x1c, 0x3c, 0x14, 0x0f, 0xb6, 0x6a, 0xe1,
	0xfb, 0x97, 0xdf, 0x17, 0x07, 0xff, 0x6f, 0x8a, 0xbd, 0xad, 0x40, 0x3e, 0x17, 0xbb, 0xe8, 0xba,
	0x02, 0x89, 0xa1, 0x09, 0xba, 0x34, 0xad, 0x3a, 0x9c, 0x8e, 0x66, 0x77, 0x16, 0x93, 0x81, 0xfe,
	0x4a, 0x50, 0x3e, 0x13, 0x93, 0x1c, 0x43, 0xad, 0x5d, 0xc5, 0x7f, 0xd4, 0xd1, 0x74, 0x34, 0x9b,
	0x2c, 0xc6, 0x19, 0xfe, 0x48, 0xac, 0x97, 0x82, 0xae, 0xb1, 0x03, 0x5c, 0xad, 0x82, 0x26, 0x52,
	0xc7, 0xa9, 0x6a, 0x9c, 0xe0, 0x59, 0x66, 0xf2, 0xae, 0xb8, 0x61, 0x39, 0xaa, 0x57, 0xe9, 0xbe,
	0x7f, 0xf6, 0x84, 0x3d, 0xa9, 0xd7, 0x99, 0xb0, 0x4f, 0x0e, 0x73, 0xad, 0xde, 0x0c, 0x84, 0xeb,
	0x7e, 0xe6, 0xdf, 0xa8, 0x43, 0x07, 0xc6, 0xb1, 0x0e, 0x97, 0x58, 0xab, 0xb7, 0x29, 0x9c, 0x24,
	0xfa, 0x6d, 0x80, 0xfd, 0x82, 0x0a, 0x59, 0xff, 0xc3, 0x0e, 0x0a, 0x1f, 0x1d, 0xab, 0x77, 0x79,
	0xe6, 0x00, 0x3f, 0xf7, 0x4c, 0x3e, 0x11, 0x3b, 0x16, 0x5b, 0x18, 0x98, 0x3a, 0x49, 0x8a, 0xb0,
	0xd8, 0x7e, 0xcd, 0x44, 0x3e, 0x15, 0x63, 0x8e, 0xce, 0xe9, 0x7a, 0x28, 0x79, 0x9f, 0x8c, 0x9d,
	0xcc, 0x72, 0xc7, 0x89, 0x78, 0x7c, 0xad, 0xb8, 0xd2, 0x54, 0x31, 0xe8, 0x15, 0x58, 0x6c, 0x8d,
	0x8d, 0x56, 0x9d, 0x26, 0xff, 0xd1, 0xda, 0x5f, 0xe7, 0x3f, 0x73, 0x2c, 0x5f, 0x8a, 0x7b, 0x86,
	0x00, 0x8b, 0xcd, 0x5b, 0xf5, 0x61, 0x3a, 0x9a, 0xdd, 0x5e, 0xec, 0x19, 0x3a, 0x2b, 0x36, 0x4e,
	0xe4, 0xa1, 0xd8, 0x37, 0xb4, 0x9e, 0x0a, 0x64, 0x6c, 0xac, 0x91, 0x8d, 0x77, 0xea, 0x63, 0xf2,
	0xef, 0x1b, 0x1a, 0x46, 0x5f, 0x5c, 0x47, 0xf2, 0x45, 0xea, 0xf7, 0x11, 0x7c, 0x09, 0x41, 0x93,
	0x8f, 0xa1, 0xd0, 0xea, 0x53, 0xf2, 0x77, 0x0d, 0x9d, 0xc7, 0xf3, 0x72, 0x31, 0xd0, 0xe5, 0xad,
	0xf4, 0xdd, 0x8e, 0xae, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xc0, 0x4b, 0x19, 0x87, 0x02, 0x00,
	0x00,
}
