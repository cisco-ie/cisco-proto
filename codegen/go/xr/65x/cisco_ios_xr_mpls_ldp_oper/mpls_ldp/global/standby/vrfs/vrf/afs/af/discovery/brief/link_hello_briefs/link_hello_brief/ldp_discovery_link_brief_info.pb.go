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
// source: ldp_discovery_link_brief_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_vrfs_vrf_afs_af_discovery_brief_link_hello_briefs_link_hello_brief

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

type LdpDiscoveryLinkBriefInfo_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpDiscoveryLinkBriefInfo_KEYS) Reset()         { *m = LdpDiscoveryLinkBriefInfo_KEYS{} }
func (m *LdpDiscoveryLinkBriefInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryLinkBriefInfo_KEYS) ProtoMessage()    {}
func (*LdpDiscoveryLinkBriefInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_33dde35b4297c6d1, []int{0}
}

func (m *LdpDiscoveryLinkBriefInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpDiscoveryLinkBriefInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryLinkBriefInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS.Merge(m, src)
}
func (m *LdpDiscoveryLinkBriefInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS.Size(m)
}
func (m *LdpDiscoveryLinkBriefInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryLinkBriefInfo_KEYS proto.InternalMessageInfo

func (m *LdpDiscoveryLinkBriefInfo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_33dde35b4297c6d1, []int{1}
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
	return fileDescriptor_33dde35b4297c6d1, []int{2}
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

type LdpDiscoveryLinkBriefInfo struct {
	Vrf                  *LdpVrfInfo          `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	AddressFamily        string               `protobuf:"bytes,51,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	AddressFamilySet     string               `protobuf:"bytes,52,opt,name=address_family_set,json=addressFamilySet,proto3" json:"address_family_set,omitempty"`
	Interface            string               `protobuf:"bytes,53,opt,name=interface,proto3" json:"interface,omitempty"`
	InterfaceNameXr      string               `protobuf:"bytes,54,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	HelloInformation     []*LdpHelloBriefInfo `protobuf:"bytes,55,rep,name=hello_information,json=helloInformation,proto3" json:"hello_information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LdpDiscoveryLinkBriefInfo) Reset()         { *m = LdpDiscoveryLinkBriefInfo{} }
func (m *LdpDiscoveryLinkBriefInfo) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryLinkBriefInfo) ProtoMessage()    {}
func (*LdpDiscoveryLinkBriefInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_33dde35b4297c6d1, []int{3}
}

func (m *LdpDiscoveryLinkBriefInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo.Unmarshal(m, b)
}
func (m *LdpDiscoveryLinkBriefInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryLinkBriefInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryLinkBriefInfo.Merge(m, src)
}
func (m *LdpDiscoveryLinkBriefInfo) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryLinkBriefInfo.Size(m)
}
func (m *LdpDiscoveryLinkBriefInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryLinkBriefInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryLinkBriefInfo proto.InternalMessageInfo

func (m *LdpDiscoveryLinkBriefInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpDiscoveryLinkBriefInfo) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo) GetAddressFamilySet() string {
	if m != nil {
		return m.AddressFamilySet
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *LdpDiscoveryLinkBriefInfo) GetHelloInformation() []*LdpHelloBriefInfo {
	if m != nil {
		return m.HelloInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpDiscoveryLinkBriefInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.discovery.brief.link_hello_briefs.link_hello_brief.ldp_discovery_link_brief_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.discovery.brief.link_hello_briefs.link_hello_brief.ldp_vrf_info")
	proto.RegisterType((*LdpHelloBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.discovery.brief.link_hello_briefs.link_hello_brief.ldp_hello_brief_info")
	proto.RegisterType((*LdpDiscoveryLinkBriefInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.vrfs.vrf.afs.af.discovery.brief.link_hello_briefs.link_hello_brief.ldp_discovery_link_brief_info")
}

func init() {
	proto.RegisterFile("ldp_discovery_link_brief_info.proto", fileDescriptor_33dde35b4297c6d1)
}

var fileDescriptor_33dde35b4297c6d1 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0x26, 0xbb, 0xa5, 0xdd, 0x7d, 0x75, 0x6b, 0x3b, 0x28, 0x8d, 0x68, 0x61, 0x89, 0x08, 0x8b,
	0xc8, 0x1c, 0xb6, 0x5a, 0xff, 0x02, 0x85, 0xa2, 0x78, 0x48, 0x15, 0xf4, 0x34, 0x4c, 0x76, 0x66,
	0xba, 0x0f, 0x27, 0x33, 0x61, 0x26, 0x0d, 0xdd, 0x93, 0x07, 0x6f, 0x7a, 0x14, 0xfc, 0x5f, 0xfc,
	0xef, 0x64, 0x26, 0xe9, 0x76, 0xa3, 0xd0, 0xeb, 0x5e, 0x42, 0xf2, 0x7d, 0xdf, 0x7b, 0xef, 0x7b,
	0x3f, 0x02, 0x4f, 0xb5, 0xa8, 0x98, 0x40, 0xbf, 0xb0, 0x8d, 0x74, 0x2b, 0xa6, 0xd1, 0x7c, 0x65,
	0x85, 0x43, 0xa9, 0x18, 0x1a, 0x65, 0x69, 0xe5, 0x6c, 0x6d, 0xc9, 0xd5, 0x22, 0x08, 0x18, 0x5a,
	0xcf, 0xae, 0x1d, 0x2b, 0x2b, 0xed, 0x59, 0x08, 0xb3, 0x95, 0x74, 0xf4, 0xe6, 0x8b, 0x5e, 0x6a,
	0x5b, 0x70, 0x4d, 0x7d, 0xcd, 0x8d, 0x28, 0x56, 0xb4, 0x71, 0xca, 0x87, 0x07, 0xe5, 0xca, 0x53,
	0xae, 0xe8, 0xba, 0x06, 0x8d, 0xe9, 0x69, 0xac, 0xb4, 0x94, 0x5a, 0xdb, 0xb6, 0x9e, 0xff, 0x0f,
	0xc9, 0xbe, 0x41, 0x76, 0xa7, 0x3b, 0xf6, 0xee, 0xcd, 0x97, 0x0b, 0xf2, 0x08, 0x46, 0x8d, 0x53,
	0xcc, 0xf0, 0x52, 0xa6, 0xc9, 0x34, 0x99, 0x8d, 0xf3, 0xbd, 0xc6, 0xa9, 0x0f, 0xbc, 0x94, 0xe4,
	0x18, 0xf6, 0x78, 0xc7, 0x0c, 0x22, 0xb3, 0xcb, 0x5b, 0xe2, 0x19, 0x1c, 0xa0, 0xa9, 0xa5, 0x53,
	0x7c, 0x21, 0x5b, 0x7e, 0x18, 0xf9, 0xc9, 0x1a, 0x0d, 0xb2, 0x6c, 0x0e, 0xf7, 0x82, 0x81, 0x90,
	0x3e, 0xd4, 0x23, 0x04, 0x76, 0x36, 0xca, 0xc4, 0x77, 0x72, 0x00, 0x03, 0x14, 0x31, 0xfd, 0x24,
	0x1f, 0xa0, 0xc8, 0x7e, 0x24, 0xf0, 0x20, 0x04, 0x6d, 0x34, 0xd2, 0x06, 0x9f, 0xc1, 0xb1, 0x91,
	0x78, 0xb9, 0x2c, 0xac, 0x8b, 0xd3, 0x43, 0x21, 0x4d, 0x8d, 0x0a, 0xa5, 0xeb, 0xf2, 0x3d, 0xbc,
	0xa1, 0xdf, 0x8b, 0xea, 0x7c, 0x4d, 0x92, 0xc7, 0x30, 0x5e, 0x5a, 0x2d, 0x58, 0x8d, 0x5d, 0x1b,
	0x93, 0x7c, 0x14, 0x80, 0x8f, 0x58, 0x4a, 0x72, 0x02, 0xe0, 0xa5, 0xf7, 0x68, 0x0d, 0xbb, 0xaa,
	0x62, 0x13, 0xa3, 0x7c, 0xdc, 0x21, 0x9f, 0xaa, 0xec, 0xd7, 0x0e, 0x9c, 0xdc, 0x39, 0x42, 0xf2,
	0x3b, 0x81, 0x61, 0xe3, 0x54, 0x3a, 0x9f, 0x26, 0xb3, 0xfd, 0xf9, 0xf7, 0x84, 0x6e, 0x65, 0xd5,
	0x74, 0x73, 0xcc, 0x79, 0x30, 0x14, 0x56, 0xc4, 0x85, 0x70, 0xd2, 0x7b, 0xa6, 0x78, 0x89, 0x7a,
	0x95, 0x9e, 0xb6, 0x2b, 0xea, 0xd0, 0xb7, 0x11, 0x24, 0x2f, 0x80, 0xf4, 0x65, 0xcc, 0xcb, 0x3a,
	0x7d, 0x19, 0xa5, 0x87, 0x3d, 0xe9, 0x85, 0xac, 0xc9, 0x13, 0x18, 0xaf, 0x37, 0x9c, 0xbe, 0x8a,
	0xa2, 0x5b, 0x80, 0x3c, 0x87, 0xa3, 0xfe, 0x55, 0xb0, 0x6b, 0x97, 0x9e, 0x45, 0xd5, 0xfd, 0xde,
	0x61, 0x7c, 0x76, 0xe4, 0x4f, 0x02, 0x47, 0x6d, 0x03, 0xc1, 0xb2, 0x2b, 0x79, 0x8d, 0xd6, 0xa4,
	0xaf, 0xa7, 0xc3, 0xd9, 0xfe, 0xfc, 0xe7, 0x36, 0xa7, 0xf8, 0xef, 0xdd, 0xe5, 0x87, 0x11, 0x39,
	0xbf, 0x75, 0x59, 0xec, 0xc6, 0xbf, 0xfa, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x32,
	0x0b, 0x5c, 0xfc, 0x03, 0x00, 0x00,
}
