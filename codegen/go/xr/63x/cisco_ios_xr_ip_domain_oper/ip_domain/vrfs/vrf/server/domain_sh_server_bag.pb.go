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
// source: domain_sh_server_bag.proto

package cisco_ios_xr_ip_domain_oper_ip_domain_vrfs_vrf_server

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

type DomainShServerBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainShServerBag_KEYS) Reset()         { *m = DomainShServerBag_KEYS{} }
func (m *DomainShServerBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*DomainShServerBag_KEYS) ProtoMessage()    {}
func (*DomainShServerBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a635b0a742977297, []int{0}
}

func (m *DomainShServerBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainShServerBag_KEYS.Unmarshal(m, b)
}
func (m *DomainShServerBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainShServerBag_KEYS.Marshal(b, m, deterministic)
}
func (m *DomainShServerBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainShServerBag_KEYS.Merge(m, src)
}
func (m *DomainShServerBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_DomainShServerBag_KEYS.Size(m)
}
func (m *DomainShServerBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainShServerBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DomainShServerBag_KEYS proto.InternalMessageInfo

func (m *DomainShServerBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type DomainIpaddress struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainIpaddress) Reset()         { *m = DomainIpaddress{} }
func (m *DomainIpaddress) String() string { return proto.CompactTextString(m) }
func (*DomainIpaddress) ProtoMessage()    {}
func (*DomainIpaddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a635b0a742977297, []int{1}
}

func (m *DomainIpaddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainIpaddress.Unmarshal(m, b)
}
func (m *DomainIpaddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainIpaddress.Marshal(b, m, deterministic)
}
func (m *DomainIpaddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainIpaddress.Merge(m, src)
}
func (m *DomainIpaddress) XXX_Size() int {
	return xxx_messageInfo_DomainIpaddress.Size(m)
}
func (m *DomainIpaddress) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainIpaddress.DiscardUnknown(m)
}

var xxx_messageInfo_DomainIpaddress proto.InternalMessageInfo

func (m *DomainIpaddress) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *DomainIpaddress) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *DomainIpaddress) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type DomainShServerBag struct {
	DomainLookup         string             `protobuf:"bytes,50,opt,name=domain_lookup,json=domainLookup,proto3" json:"domain_lookup,omitempty"`
	DomainName           string             `protobuf:"bytes,51,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	ServerAddress        []*DomainIpaddress `protobuf:"bytes,52,rep,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	Domain               []string           `protobuf:"bytes,53,rep,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DomainShServerBag) Reset()         { *m = DomainShServerBag{} }
func (m *DomainShServerBag) String() string { return proto.CompactTextString(m) }
func (*DomainShServerBag) ProtoMessage()    {}
func (*DomainShServerBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a635b0a742977297, []int{2}
}

func (m *DomainShServerBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainShServerBag.Unmarshal(m, b)
}
func (m *DomainShServerBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainShServerBag.Marshal(b, m, deterministic)
}
func (m *DomainShServerBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainShServerBag.Merge(m, src)
}
func (m *DomainShServerBag) XXX_Size() int {
	return xxx_messageInfo_DomainShServerBag.Size(m)
}
func (m *DomainShServerBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainShServerBag.DiscardUnknown(m)
}

var xxx_messageInfo_DomainShServerBag proto.InternalMessageInfo

func (m *DomainShServerBag) GetDomainLookup() string {
	if m != nil {
		return m.DomainLookup
	}
	return ""
}

func (m *DomainShServerBag) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *DomainShServerBag) GetServerAddress() []*DomainIpaddress {
	if m != nil {
		return m.ServerAddress
	}
	return nil
}

func (m *DomainShServerBag) GetDomain() []string {
	if m != nil {
		return m.Domain
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainShServerBag_KEYS)(nil), "cisco_ios_xr_ip_domain_oper.ip_domain.vrfs.vrf.server.domain_sh_server_bag_KEYS")
	proto.RegisterType((*DomainIpaddress)(nil), "cisco_ios_xr_ip_domain_oper.ip_domain.vrfs.vrf.server.domain_ipaddress")
	proto.RegisterType((*DomainShServerBag)(nil), "cisco_ios_xr_ip_domain_oper.ip_domain.vrfs.vrf.server.domain_sh_server_bag")
}

func init() { proto.RegisterFile("domain_sh_server_bag.proto", fileDescriptor_a635b0a742977297) }

var fileDescriptor_a635b0a742977297 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x15, 0x22, 0xa5, 0x70, 0x69, 0x11, 0xb2, 0x10, 0xa4, 0x2c, 0x84, 0xb0, 0x74, 0xf2,
	0xd0, 0x3f, 0xd9, 0x19, 0x10, 0x03, 0x88, 0xa1, 0x4c, 0x4c, 0x27, 0xb7, 0x71, 0xc0, 0x82, 0xc4,
	0xc6, 0x06, 0x8b, 0x0f, 0xcc, 0x07, 0x41, 0xb1, 0xdd, 0x48, 0xa0, 0x4c, 0x5d, 0x22, 0xdd, 0xbb,
	0xdf, 0x3d, 0xbd, 0x3c, 0xc3, 0x45, 0x25, 0x1b, 0x26, 0x5a, 0x34, 0xaf, 0x68, 0xb8, 0xb6, 0x5c,
	0xe3, 0x86, 0xbd, 0x50, 0xa5, 0xe5, 0xa7, 0x24, 0xab, 0xad, 0x30, 0x5b, 0x89, 0x42, 0x1a, 0xfc,
	0xd6, 0x28, 0x14, 0x06, 0x56, 0x2a, 0xae, 0x69, 0x3f, 0x52, 0xab, 0x6b, 0xd3, 0x7d, 0xa8, 0x37,
	0x28, 0x4a, 0x98, 0x0e, 0x99, 0xe2, 0xfd, 0xed, 0xf3, 0x13, 0x99, 0xc2, 0xa1, 0xd5, 0x35, 0xb6,
	0xac, 0xe1, 0x59, 0x94, 0x47, 0xb3, 0xa3, 0xf5, 0xc8, 0xea, 0xfa, 0x91, 0x35, 0xbc, 0xf8, 0x80,
	0x93, 0x70, 0x27, 0x14, 0xab, 0x2a, 0xcd, 0x8d, 0x21, 0xe7, 0x30, 0x62, 0x7f, 0xe8, 0x84, 0x39,
	0x98, 0x5c, 0xc1, 0x58, 0x28, 0xbb, 0xc4, 0x00, 0x66, 0x07, 0x6e, 0x9b, 0x76, 0xda, 0x4d, 0xb8,
	0xf5, 0x48, 0xd9, 0x23, 0x71, 0x8f, 0x94, 0x01, 0x29, 0x7e, 0x22, 0x38, 0x1d, 0xca, 0x4a, 0xae,
	0x61, 0x12, 0xf4, 0x77, 0x29, 0xdf, 0xbe, 0x54, 0x36, 0x77, 0xc7, 0x63, 0x2f, 0x3e, 0x38, 0x8d,
	0x5c, 0x42, 0x1a, 0x20, 0x17, 0x70, 0xe1, 0x10, 0xf0, 0x92, 0x0b, 0xd9, 0xc2, 0x71, 0xf0, 0xdc,
	0x65, 0x58, 0xe6, 0xf1, 0x2c, 0x9d, 0xdf, 0xd1, 0xbd, 0x9a, 0xa5, 0xff, 0xeb, 0x59, 0x4f, 0xfc,
	0x62, 0xf7, 0xc7, 0x67, 0x90, 0x78, 0x24, 0x5b, 0xe5, 0x71, 0x57, 0x96, 0x9f, 0x36, 0x89, 0x7b,
	0xcf, 0xc5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xc7, 0xe5, 0x10, 0xed, 0x01, 0x00, 0x00,
}
