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
// source: ldp_targeted_hello_brief_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_afs_af_discovery_brief_targeted_hello_briefs_targeted_hello_brief

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

type LdpTargetedHelloBriefInfo_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LocalAddress         string   `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	TargetAddress        string   `protobuf:"bytes,3,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpTargetedHelloBriefInfo_KEYS) Reset()         { *m = LdpTargetedHelloBriefInfo_KEYS{} }
func (m *LdpTargetedHelloBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpTargetedHelloBriefInfo_KEYS) ProtoMessage()    {}
func (*LdpTargetedHelloBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e66bbc349281ebba, []int{0}
}

func (m *LdpTargetedHelloBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpTargetedHelloBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpTargetedHelloBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS.Merge(m, src)
}
func (m *LdpTargetedHelloBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS.Size(m)
}
func (m *LdpTargetedHelloBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTargetedHelloBriefInfo_KEYS proto.InternalMessageInfo

func (m *LdpTargetedHelloBriefInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpTargetedHelloBriefInfo_KEYS) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *LdpTargetedHelloBriefInfo_KEYS) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

type LdpVrfInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpVrfInfo) Reset()         { *m = LdpVrfInfo{} }
func (m *LdpVrfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpVrfInfo) ProtoMessage()    {}
func (*LdpVrfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e66bbc349281ebba, []int{1}
}

func (m *LdpVrfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpVrfInfo.Unmarshal(m, b)
}
func (m *LdpVrfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpVrfInfo.Marshal(b, m, deterministic)
}
func (m *LdpVrfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpVrfInfo.Merge(m, src)
}
func (m *LdpVrfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpVrfInfo.Size(m)
}
func (m *LdpVrfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpVrfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpVrfInfo proto.InternalMessageInfo

func (m *LdpVrfInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LdpVrfInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type LdpIpAddrTUnion struct {
	Afi                  string   `protobuf:"bytes,1,opt,name=afi,proto3" json:"afi,omitempty"`
	Dummy                uint32   `protobuf:"varint,2,opt,name=dummy,proto3" json:"dummy,omitempty"`
	Ipv4                 string   `protobuf:"bytes,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,4,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIpAddrTUnion) Reset()         { *m = LdpIpAddrTUnion{} }
func (m *LdpIpAddrTUnion) String() string { return proto.CompactTextString(m) }
func (*LdpIpAddrTUnion) ProtoMessage()    {}
func (*LdpIpAddrTUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_e66bbc349281ebba, []int{2}
}

func (m *LdpIpAddrTUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIpAddrTUnion.Unmarshal(m, b)
}
func (m *LdpIpAddrTUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIpAddrTUnion.Marshal(b, m, deterministic)
}
func (m *LdpIpAddrTUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIpAddrTUnion.Merge(m, src)
}
func (m *LdpIpAddrTUnion) XXX_Size() int {
	return xxx_messageInfo_LdpIpAddrTUnion.Size(m)
}
func (m *LdpIpAddrTUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIpAddrTUnion.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIpAddrTUnion proto.InternalMessageInfo

func (m *LdpIpAddrTUnion) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

func (m *LdpIpAddrTUnion) GetDummy() uint32 {
	if m != nil {
		return m.Dummy
	}
	return 0
}

func (m *LdpIpAddrTUnion) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *LdpIpAddrTUnion) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type LdpHelloBriefInfo struct {
	NeighborLdpIdentifier string   `protobuf:"bytes,1,opt,name=neighbor_ldp_identifier,json=neighborLdpIdentifier,proto3" json:"neighbor_ldp_identifier,omitempty"`
	HoldTime              uint32   `protobuf:"varint,2,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	SessionUp             bool     `protobuf:"varint,3,opt,name=session_up,json=sessionUp,proto3" json:"session_up,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LdpHelloBriefInfo) Reset()         { *m = LdpHelloBriefInfo{} }
func (m *LdpHelloBriefInfo) String() string { return proto.CompactTextString(m) }
func (*LdpHelloBriefInfo) ProtoMessage()    {}
func (*LdpHelloBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e66bbc349281ebba, []int{3}
}

func (m *LdpHelloBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpHelloBriefInfo.Unmarshal(m, b)
}
func (m *LdpHelloBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpHelloBriefInfo.Marshal(b, m, deterministic)
}
func (m *LdpHelloBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpHelloBriefInfo.Merge(m, src)
}
func (m *LdpHelloBriefInfo) XXX_Size() int {
	return xxx_messageInfo_LdpHelloBriefInfo.Size(m)
}
func (m *LdpHelloBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpHelloBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpHelloBriefInfo proto.InternalMessageInfo

func (m *LdpHelloBriefInfo) GetNeighborLdpIdentifier() string {
	if m != nil {
		return m.NeighborLdpIdentifier
	}
	return ""
}

func (m *LdpHelloBriefInfo) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LdpHelloBriefInfo) GetSessionUp() bool {
	if m != nil {
		return m.SessionUp
	}
	return false
}

type LdpTargetedHelloBriefInfo struct {
	Vrf                  *LdpVrfInfo          `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AddressFamily        string               `protobuf:"bytes,51,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	DhcbTargetAddress    *LdpIpAddrTUnion     `protobuf:"bytes,52,opt,name=dhcb_target_address,json=dhcbTargetAddress,proto3" json:"dhcb_target_address,omitempty"`
	HelloInformation     []*LdpHelloBriefInfo `protobuf:"bytes,53,rep,name=hello_information,json=helloInformation,proto3" json:"hello_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LdpTargetedHelloBriefInfo) Reset()         { *m = LdpTargetedHelloBriefInfo{} }
func (m *LdpTargetedHelloBriefInfo) String() string { return proto.CompactTextString(m) }
func (*LdpTargetedHelloBriefInfo) ProtoMessage()    {}
func (*LdpTargetedHelloBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e66bbc349281ebba, []int{4}
}

func (m *LdpTargetedHelloBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo.Unmarshal(m, b)
}
func (m *LdpTargetedHelloBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo.Marshal(b, m, deterministic)
}
func (m *LdpTargetedHelloBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTargetedHelloBriefInfo.Merge(m, src)
}
func (m *LdpTargetedHelloBriefInfo) XXX_Size() int {
	return xxx_messageInfo_LdpTargetedHelloBriefInfo.Size(m)
}
func (m *LdpTargetedHelloBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTargetedHelloBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTargetedHelloBriefInfo proto.InternalMessageInfo

func (m *LdpTargetedHelloBriefInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpTargetedHelloBriefInfo) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpTargetedHelloBriefInfo) GetDhcbTargetAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.DhcbTargetAddress
	}
	return nil
}

func (m *LdpTargetedHelloBriefInfo) GetHelloInformation() []*LdpHelloBriefInfo {
	if m != nil {
		return m.HelloInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpTargetedHelloBriefInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.discovery.brief.targeted_hello_briefs.targeted_hello_brief.ldp_targeted_hello_brief_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.discovery.brief.targeted_hello_briefs.targeted_hello_brief.ldp_vrf_info")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.discovery.brief.targeted_hello_briefs.targeted_hello_brief.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpHelloBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.discovery.brief.targeted_hello_briefs.targeted_hello_brief.ldp_hello_brief_info")
	proto.RegisterType((*LdpTargetedHelloBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.discovery.brief.targeted_hello_briefs.targeted_hello_brief.ldp_targeted_hello_brief_info")
}

func init() {
	proto.RegisterFile("ldp_targeted_hello_brief_info.proto", fileDescriptor_e66bbc349281ebba)
}

var fileDescriptor_e66bbc349281ebba = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x4f, 0x6b, 0x1b, 0x3f,
	0x10, 0x65, 0x6d, 0xff, 0xf2, 0x8b, 0x27, 0x71, 0x48, 0x94, 0x94, 0x2c, 0x94, 0x80, 0xd9, 0x50,
	0xf0, 0x49, 0x07, 0x27, 0xf5, 0xbd, 0x87, 0x16, 0x42, 0x4b, 0x0f, 0xdb, 0xf4, 0xd0, 0x93, 0x90,
	0x57, 0x92, 0x3d, 0xa0, 0x5d, 0x2d, 0xd2, 0x7a, 0xa9, 0x4f, 0x3d, 0xb7, 0xd0, 0x5b, 0x7b, 0xef,
	0x77, 0x29, 0xfd, 0x4c, 0xbd, 0x16, 0x69, 0x15, 0xd7, 0x98, 0x90, 0x6b, 0x73, 0x9b, 0x7d, 0xf3,
	0xe7, 0xbd, 0xd1, 0x3c, 0x16, 0x2e, 0xb5, 0xa8, 0x59, 0xc3, 0xed, 0x42, 0x36, 0x52, 0xb0, 0xa5,
	0xd4, 0xda, 0xb0, 0xb9, 0x45, 0xa9, 0x18, 0x56, 0xca, 0xd0, 0xda, 0x9a, 0xc6, 0x90, 0x4f, 0x05,
	0xba, 0xc2, 0x30, 0x34, 0x8e, 0x7d, 0xb4, 0xac, 0xac, 0xb5, 0x63, 0xbe, 0xcd, 0xd4, 0xd2, 0xd2,
	0xbb, 0x2f, 0xba, 0xd0, 0x66, 0xce, 0x35, 0xe5, 0x45, 0x83, 0xad, 0xa4, 0x42, 0x2a, 0xbe, 0xd2,
	0x0d, 0x6b, 0xad, 0xa2, 0x5c, 0x39, 0xca, 0x15, 0x15, 0x7e, 0x48, 0x2b, 0xed, 0x9a, 0x06, 0x02,
	0x7a, 0x1f, 0xab, 0xbb, 0x17, 0xcd, 0x3e, 0x27, 0x90, 0x3d, 0x28, 0x94, 0xbd, 0x7e, 0xf9, 0xe1,
	0x1d, 0x39, 0x87, 0xff, 0xb9, 0x62, 0x15, 0x2f, 0x65, 0x9a, 0x8c, 0x93, 0xc9, 0x30, 0xdf, 0xe3,
	0xea, 0x2d, 0x2f, 0x25, 0xb9, 0x84, 0x91, 0x36, 0x05, 0xd7, 0x8c, 0x0b, 0x61, 0xa5, 0x73, 0x69,
	0x2f, 0xa4, 0x0f, 0x03, 0xf8, 0xa2, 0xc3, 0xc8, 0x33, 0x38, 0xea, 0xe6, 0x6f, 0xaa, 0xfa, 0xa1,
	0x6a, 0xd4, 0xa1, 0xb1, 0x2c, 0x9b, 0xc2, 0xa1, 0x97, 0xd2, 0xda, 0x8e, 0x99, 0x10, 0x18, 0x6c,
	0x31, 0x86, 0x98, 0x1c, 0x41, 0x0f, 0x45, 0x20, 0x19, 0xe5, 0x3d, 0x14, 0x99, 0x84, 0x53, 0xdf,
	0x83, 0x75, 0x18, 0xcd, 0x1a, 0xb6, 0xaa, 0xd0, 0x54, 0xe4, 0x18, 0xfa, 0x5c, 0x61, 0xec, 0xf4,
	0x21, 0x39, 0x83, 0xff, 0xc4, 0xaa, 0x2c, 0xd7, 0xb1, 0xb7, 0xfb, 0xf0, 0x14, 0x58, 0xb7, 0xd7,
	0x51, 0x4f, 0x88, 0x23, 0x36, 0x4b, 0x07, 0x1b, 0x6c, 0x96, 0x7d, 0x49, 0xe0, 0xcc, 0xf3, 0xec,
	0xbe, 0x0e, 0x99, 0xc1, 0x79, 0x25, 0x71, 0xb1, 0x9c, 0x1b, 0x1b, 0x2e, 0x87, 0x42, 0x56, 0x0d,
	0x2a, 0x94, 0x36, 0x92, 0x3f, 0xb9, 0x4b, 0xbf, 0x11, 0xf5, 0xcd, 0x26, 0x49, 0x9e, 0xc2, 0x70,
	0x69, 0xb4, 0x60, 0x0d, 0x96, 0x32, 0x4a, 0xda, 0xf7, 0xc0, 0x2d, 0x96, 0x92, 0x5c, 0x00, 0x38,
	0xe9, 0x1c, 0x9a, 0x8a, 0xad, 0xea, 0xa0, 0x6d, 0x3f, 0x1f, 0x46, 0xe4, 0x7d, 0x9d, 0xfd, 0x1e,
	0xc0, 0xc5, 0x83, 0x37, 0x23, 0x3f, 0x12, 0xe8, 0xb7, 0x56, 0xa5, 0xd3, 0x71, 0x32, 0x39, 0x98,
	0x7e, 0x4d, 0xe8, 0x3f, 0xb6, 0x19, 0xdd, 0xbe, 0x6b, 0xee, 0xa5, 0x79, 0x4f, 0x44, 0x33, 0x30,
	0xc5, 0x4b, 0xd4, 0xeb, 0xf4, 0xaa, 0xf3, 0x44, 0x44, 0x5f, 0x05, 0x90, 0xfc, 0x4a, 0xe0, 0x54,
	0x2c, 0x8b, 0x39, 0xdb, 0x31, 0xd0, 0x75, 0xd8, 0xec, 0xdb, 0xe3, 0xd8, 0x6c, 0xc7, 0x7d, 0xf9,
	0x89, 0x57, 0x7c, 0xbb, 0xed, 0x6d, 0xf2, 0x33, 0x81, 0x93, 0xae, 0xcd, 0x3f, 0x81, 0x2d, 0x79,
	0x83, 0xa6, 0x4a, 0x9f, 0x8f, 0xfb, 0x93, 0x83, 0xe9, 0xf7, 0xc7, 0xb1, 0xc5, 0xae, 0x8b, 0xf2,
	0xe3, 0x80, 0xdc, 0xfc, 0xd5, 0x3b, 0xdf, 0x0b, 0x7f, 0xad, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0xd3, 0x25, 0xe1, 0xdc, 0x04, 0x00, 0x00,
}
