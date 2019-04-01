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
// source: ipv6_dhcpv6d_server_profile.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_server_profiles_profile_info

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

type Ipv6Dhcpv6DServerProfile_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProfileName          string   `protobuf:"bytes,2,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerProfile_KEYS) Reset()         { *m = Ipv6Dhcpv6DServerProfile_KEYS{} }
func (m *Ipv6Dhcpv6DServerProfile_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerProfile_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerProfile_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac60509bf736c3ca, []int{0}
}

func (m *Ipv6Dhcpv6DServerProfile_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerProfile_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerProfile_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerProfile_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DServerProfile_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerProfile_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerProfile_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile_KEYS) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

type TimeBag struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeBag) Reset()         { *m = TimeBag{} }
func (m *TimeBag) String() string { return proto.CompactTextString(m) }
func (*TimeBag) ProtoMessage()    {}
func (*TimeBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac60509bf736c3ca, []int{1}
}

func (m *TimeBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeBag.Unmarshal(m, b)
}
func (m *TimeBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeBag.Marshal(b, m, deterministic)
}
func (m *TimeBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeBag.Merge(m, src)
}
func (m *TimeBag) XXX_Size() int {
	return xxx_messageInfo_TimeBag.Size(m)
}
func (m *TimeBag) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeBag.DiscardUnknown(m)
}

var xxx_messageInfo_TimeBag proto.InternalMessageInfo

func (m *TimeBag) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TimeBag) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

type Ipv6Dhcpv6DServerInterfaceReferenceItem struct {
	ServerReferenceInterfaceName string   `protobuf:"bytes,1,opt,name=server_reference_interface_name,json=serverReferenceInterfaceName,proto3" json:"server_reference_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) Reset() {
	*m = Ipv6Dhcpv6DServerInterfaceReferenceItem{}
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerInterfaceReferenceItem) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerInterfaceReferenceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac60509bf736c3ca, []int{2}
}

func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem.Size(m)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceItem proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerInterfaceReferenceItem) GetServerReferenceInterfaceName() string {
	if m != nil {
		return m.ServerReferenceInterfaceName
	}
	return ""
}

type Ipv6Dhcpv6DServerInterfaceReferenceEntry struct {
	Ipv6Dhcpv6DServerInterfaceReference []*Ipv6Dhcpv6DServerInterfaceReferenceItem `protobuf:"bytes,1,rep,name=ipv6_dhcpv6d_server_interface_reference,json=ipv6Dhcpv6dServerInterfaceReference,proto3" json:"ipv6_dhcpv6d_server_interface_reference,omitempty"`
	XXX_NoUnkeyedLiteral                struct{}                                   `json:"-"`
	XXX_unrecognized                    []byte                                     `json:"-"`
	XXX_sizecache                       int32                                      `json:"-"`
}

func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) Reset() {
	*m = Ipv6Dhcpv6DServerInterfaceReferenceEntry{}
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerInterfaceReferenceEntry) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerInterfaceReferenceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac60509bf736c3ca, []int{3}
}

func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry.Size(m)
}
func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerInterfaceReferenceEntry proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerInterfaceReferenceEntry) GetIpv6Dhcpv6DServerInterfaceReference() []*Ipv6Dhcpv6DServerInterfaceReferenceItem {
	if m != nil {
		return m.Ipv6Dhcpv6DServerInterfaceReference
	}
	return nil
}

type Ipv6Dhcpv6DServerProfile struct {
	ProfileName             string                                    `protobuf:"bytes,50,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	ProfileAllowedDuidType  uint32                                    `protobuf:"varint,51,opt,name=profile_allowed_duid_type,json=profileAllowedDuidType,proto3" json:"profile_allowed_duid_type,omitempty"`
	DomainName              string                                    `protobuf:"bytes,52,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	ProfileDns              uint32                                    `protobuf:"varint,53,opt,name=profile_dns,json=profileDns,proto3" json:"profile_dns,omitempty"`
	ProfileDnsAddress       []string                                  `protobuf:"bytes,54,rep,name=profile_dns_address,json=profileDnsAddress,proto3" json:"profile_dns_address,omitempty"`
	Lease                   *TimeBag                                  `protobuf:"bytes,55,opt,name=lease,proto3" json:"lease,omitempty"`
	AftrName                string                                    `protobuf:"bytes,56,opt,name=aftr_name,json=aftrName,proto3" json:"aftr_name,omitempty"`
	FramedAddrPoolName      string                                    `protobuf:"bytes,57,opt,name=framed_addr_pool_name,json=framedAddrPoolName,proto3" json:"framed_addr_pool_name,omitempty"`
	DelegatedPrefixPoolName string                                    `protobuf:"bytes,58,opt,name=delegated_prefix_pool_name,json=delegatedPrefixPoolName,proto3" json:"delegated_prefix_pool_name,omitempty"`
	InterfaceReferences     *Ipv6Dhcpv6DServerInterfaceReferenceEntry `protobuf:"bytes,59,opt,name=interface_references,json=interfaceReferences,proto3" json:"interface_references,omitempty"`
	RapidCommit             bool                                      `protobuf:"varint,60,opt,name=rapid_commit,json=rapidCommit,proto3" json:"rapid_commit,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                  `json:"-"`
	XXX_unrecognized        []byte                                    `json:"-"`
	XXX_sizecache           int32                                     `json:"-"`
}

func (m *Ipv6Dhcpv6DServerProfile) Reset()         { *m = Ipv6Dhcpv6DServerProfile{} }
func (m *Ipv6Dhcpv6DServerProfile) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DServerProfile) ProtoMessage()    {}
func (*Ipv6Dhcpv6DServerProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac60509bf736c3ca, []int{4}
}

func (m *Ipv6Dhcpv6DServerProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DServerProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DServerProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DServerProfile.Merge(m, src)
}
func (m *Ipv6Dhcpv6DServerProfile) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DServerProfile.Size(m)
}
func (m *Ipv6Dhcpv6DServerProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DServerProfile.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DServerProfile proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DServerProfile) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile) GetProfileAllowedDuidType() uint32 {
	if m != nil {
		return m.ProfileAllowedDuidType
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerProfile) GetDomainName() string {
	if m != nil {
		return m.DomainName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile) GetProfileDns() uint32 {
	if m != nil {
		return m.ProfileDns
	}
	return 0
}

func (m *Ipv6Dhcpv6DServerProfile) GetProfileDnsAddress() []string {
	if m != nil {
		return m.ProfileDnsAddress
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerProfile) GetLease() *TimeBag {
	if m != nil {
		return m.Lease
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerProfile) GetAftrName() string {
	if m != nil {
		return m.AftrName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile) GetFramedAddrPoolName() string {
	if m != nil {
		return m.FramedAddrPoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile) GetDelegatedPrefixPoolName() string {
	if m != nil {
		return m.DelegatedPrefixPoolName
	}
	return ""
}

func (m *Ipv6Dhcpv6DServerProfile) GetInterfaceReferences() *Ipv6Dhcpv6DServerInterfaceReferenceEntry {
	if m != nil {
		return m.InterfaceReferences
	}
	return nil
}

func (m *Ipv6Dhcpv6DServerProfile) GetRapidCommit() bool {
	if m != nil {
		return m.RapidCommit
	}
	return false
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DServerProfile_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.info.ipv6_dhcpv6d_server_profile_KEYS")
	proto.RegisterType((*TimeBag)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.info.time_bag")
	proto.RegisterType((*Ipv6Dhcpv6DServerInterfaceReferenceItem)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.info.ipv6_dhcpv6d_server_interface_reference_item")
	proto.RegisterType((*Ipv6Dhcpv6DServerInterfaceReferenceEntry)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.info.ipv6_dhcpv6d_server_interface_reference_entry")
	proto.RegisterType((*Ipv6Dhcpv6DServerProfile)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.server.profiles.profile.info.ipv6_dhcpv6d_server_profile")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_server_profile.proto", fileDescriptor_ac60509bf736c3ca) }

var fileDescriptor_ac60509bf736c3ca = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0xf6, 0xd7, 0x1f, 0xc9, 0x04, 0x0e, 0x6c, 0xf9, 0x63, 0x28, 0x52, 0xdc, 0x70,
	0x20, 0x07, 0xb0, 0x44, 0x0a, 0xa1, 0xa5, 0x5c, 0x2a, 0xd2, 0x03, 0x42, 0x42, 0xc5, 0xe5, 0x82,
	0x38, 0xac, 0x36, 0xde, 0x71, 0x59, 0xc9, 0xf6, 0x9a, 0x5d, 0x27, 0x6d, 0x4e, 0x70, 0xe3, 0x85,
	0x78, 0x07, 0x9e, 0x80, 0xf7, 0x41, 0x9e, 0xb5, 0x4d, 0xa4, 0x54, 0x55, 0x91, 0x40, 0x5c, 0x22,
	0x7b, 0xe6, 0x33, 0xdf, 0x19, 0xcf, 0x9f, 0xc0, 0xb6, 0x2a, 0xe6, 0x63, 0x2e, 0x3f, 0xc6, 0xc5,
	0x7c, 0x2c, 0xb9, 0x45, 0x33, 0x47, 0xc3, 0x0b, 0xa3, 0x13, 0x95, 0x62, 0x58, 0x18, 0x5d, 0x6a,
	0xf6, 0x36, 0x56, 0x36, 0xd6, 0x5c, 0x69, 0xcb, 0xcf, 0x0c, 0x27, 0x3e, 0xc7, 0xd3, 0x36, 0x46,
	0x17, 0x68, 0x42, 0xf7, 0x12, 0xe6, 0x5a, 0xa2, 0xa5, 0xdf, 0xd0, 0x49, 0x85, 0xb5, 0x94, 0x6d,
	0x1e, 0x42, 0x95, 0x27, 0x7a, 0x30, 0x85, 0xe0, 0x82, 0xbc, 0xfc, 0xf5, 0xe1, 0xfb, 0x63, 0xb6,
	0x05, 0xdd, 0x4a, 0x88, 0xe7, 0x22, 0x43, 0xdf, 0x0b, 0xbc, 0x61, 0x37, 0xea, 0x54, 0x86, 0x37,
	0x22, 0x43, 0xb6, 0x0d, 0x57, 0x1b, 0x98, 0xfc, 0x6b, 0xe4, 0xef, 0xd5, 0xb6, 0x0a, 0x19, 0xec,
	0x42, 0xa7, 0x54, 0x19, 0xf2, 0xa9, 0x38, 0x61, 0x3e, 0x5c, 0xb1, 0x18, 0xeb, 0x5c, 0x5a, 0x52,
	0xba, 0x16, 0x35, 0xaf, 0x8c, 0xc1, 0x7f, 0x15, 0x55, 0x0b, 0xd0, 0xf3, 0x60, 0x06, 0x0f, 0xcf,
	0xab, 0x4e, 0xe5, 0x25, 0x9a, 0x44, 0xc4, 0xc8, 0x0d, 0x26, 0x68, 0x30, 0x8f, 0x91, 0xab, 0x12,
	0x33, 0x76, 0x08, 0xfd, 0x1a, 0x59, 0x72, 0xb4, 0xf0, 0x52, 0xfd, 0xf7, 0x1c, 0x16, 0x35, 0xd4,
	0xab, 0x06, 0xa2, 0x82, 0xbf, 0xae, 0xc1, 0xa3, 0xcb, 0xe6, 0xc5, 0xbc, 0x34, 0x0b, 0xf6, 0xc3,
	0x83, 0x07, 0x97, 0x8c, 0xf0, 0xbd, 0x60, 0x7d, 0xd8, 0x1b, 0x7d, 0x0e, 0xff, 0xf8, 0x30, 0xc3,
	0xdf, 0xe9, 0x55, 0x74, 0xbf, 0xa2, 0x27, 0x0e, 0x3e, 0x26, 0xb6, 0x6d, 0x42, 0xdb, 0x96, 0xc1,
	0xf7, 0x0d, 0xd8, 0xba, 0x60, 0x3f, 0x56, 0xa6, 0x3f, 0x5a, 0x99, 0x3e, 0xdb, 0x83, 0x3b, 0x0d,
	0x22, 0xd2, 0x54, 0x9f, 0xa2, 0xe4, 0x72, 0xa6, 0x24, 0x2f, 0x17, 0x05, 0xfa, 0x3b, 0xb4, 0x03,
	0xb7, 0x6a, 0xe0, 0xc0, 0xf9, 0x27, 0x33, 0x25, 0xdf, 0x2d, 0x0a, 0x64, 0x7d, 0xe8, 0x49, 0x9d,
	0x09, 0x95, 0x3b, 0xf1, 0x27, 0x24, 0x0e, 0xce, 0x44, 0xda, 0x7d, 0x68, 0x52, 0x71, 0x99, 0x5b,
	0xff, 0x29, 0xa9, 0x41, 0x6d, 0x9a, 0xe4, 0x96, 0x85, 0xb0, 0xb9, 0x04, 0x70, 0x21, 0xa5, 0x41,
	0x6b, 0xfd, 0x71, 0xb0, 0x3e, 0xec, 0x46, 0xd7, 0x7f, 0x81, 0x07, 0xce, 0xc1, 0x3e, 0xc1, 0x46,
	0x8a, 0xc2, 0xa2, 0xff, 0x2c, 0xf0, 0x86, 0xbd, 0xd1, 0x87, 0xbf, 0x30, 0xa4, 0xe6, 0x14, 0x22,
	0x97, 0xa9, 0xba, 0x2e, 0x91, 0x94, 0xc6, 0x7d, 0xe2, 0xae, 0xbb, 0xae, 0xca, 0x40, 0x1f, 0xf8,
	0x18, 0x6e, 0x26, 0x46, 0x64, 0x28, 0xa9, 0x74, 0x5e, 0x68, 0x9d, 0x3a, 0x70, 0x8f, 0x40, 0xe6,
	0x9c, 0x55, 0xf5, 0x47, 0x5a, 0xa7, 0x14, 0xb2, 0x0f, 0x77, 0x25, 0xa6, 0x78, 0x22, 0x4a, 0x94,
	0xbc, 0x30, 0x98, 0xa8, 0xb3, 0xa5, 0xb8, 0xe7, 0x14, 0x77, 0xbb, 0x25, 0x8e, 0x08, 0x68, 0x83,
	0xbf, 0x79, 0x70, 0xe3, 0x9c, 0x8d, 0xb1, 0xfe, 0x3e, 0xf5, 0xe3, 0x8b, 0xf7, 0x0f, 0xb7, 0x96,
	0x2e, 0x2d, 0xda, 0x54, 0x2b, 0x5b, 0x6a, 0xab, 0x35, 0x34, 0xa2, 0x50, 0x92, 0xc7, 0x3a, 0xcb,
	0x54, 0xe9, 0xbf, 0x08, 0xbc, 0x61, 0x27, 0xea, 0x91, 0xed, 0x25, 0x99, 0xa6, 0xff, 0xd3, 0x5f,
	0xe8, 0xce, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x67, 0xff, 0xe3, 0x67, 0x05, 0x00, 0x00,
}