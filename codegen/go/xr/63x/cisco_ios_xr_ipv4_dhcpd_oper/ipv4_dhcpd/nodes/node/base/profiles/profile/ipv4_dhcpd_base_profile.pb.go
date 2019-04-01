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
// source: ipv4_dhcpd_base_profile.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_base_profiles_profile

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

type Ipv4DhcpdBaseProfile_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	ProfileName          string   `protobuf:"bytes,2,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdBaseProfile_KEYS) Reset()         { *m = Ipv4DhcpdBaseProfile_KEYS{} }
func (m *Ipv4DhcpdBaseProfile_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseProfile_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdBaseProfile_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{0}
}

func (m *Ipv4DhcpdBaseProfile_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseProfile_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseProfile_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdBaseProfile_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS.Size(m)
}
func (m *Ipv4DhcpdBaseProfile_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseProfile_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseProfile_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

func (m *Ipv4DhcpdBaseProfile_KEYS) GetProfileName() string {
	if m != nil {
		return m.ProfileName
	}
	return ""
}

type Ipv4DhcpdBaseInterfaceReferenceItem struct {
	BaseReferenceInterfaceName string   `protobuf:"bytes,1,opt,name=base_reference_interface_name,json=baseReferenceInterfaceName,proto3" json:"base_reference_interface_name,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *Ipv4DhcpdBaseInterfaceReferenceItem) Reset()         { *m = Ipv4DhcpdBaseInterfaceReferenceItem{} }
func (m *Ipv4DhcpdBaseInterfaceReferenceItem) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseInterfaceReferenceItem) ProtoMessage()    {}
func (*Ipv4DhcpdBaseInterfaceReferenceItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{1}
}

func (m *Ipv4DhcpdBaseInterfaceReferenceItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem.Merge(m, src)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceItem) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem.Size(m)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceItem proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseInterfaceReferenceItem) GetBaseReferenceInterfaceName() string {
	if m != nil {
		return m.BaseReferenceInterfaceName
	}
	return ""
}

type Ipv4DhcpdBaseInterfaceReferenceEntry struct {
	Ipv4DhcpdBaseInterfaceReference []*Ipv4DhcpdBaseInterfaceReferenceItem `protobuf:"bytes,2,rep,name=ipv4_dhcpd_base_interface_reference,json=ipv4DhcpdBaseInterfaceReference,proto3" json:"ipv4_dhcpd_base_interface_reference,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                               `json:"-"`
	XXX_unrecognized                []byte                                 `json:"-"`
	XXX_sizecache                   int32                                  `json:"-"`
}

func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) Reset()         { *m = Ipv4DhcpdBaseInterfaceReferenceEntry{} }
func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseInterfaceReferenceEntry) ProtoMessage()    {}
func (*Ipv4DhcpdBaseInterfaceReferenceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{2}
}

func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry.Merge(m, src)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry.Size(m)
}
func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseInterfaceReferenceEntry proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseInterfaceReferenceEntry) GetIpv4DhcpdBaseInterfaceReference() []*Ipv4DhcpdBaseInterfaceReferenceItem {
	if m != nil {
		return m.Ipv4DhcpdBaseInterfaceReference
	}
	return nil
}

type Ipv4DhcpdBaseChildProfileInfoItem struct {
	BaseChildProfileName string   `protobuf:"bytes,1,opt,name=base_child_profile_name,json=baseChildProfileName,proto3" json:"base_child_profile_name,omitempty"`
	Mode                 uint32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	MatchedOptionCode    uint32   `protobuf:"varint,3,opt,name=matched_option_code,json=matchedOptionCode,proto3" json:"matched_option_code,omitempty"`
	MatchedOptionLen     uint32   `protobuf:"varint,4,opt,name=matched_option_len,json=matchedOptionLen,proto3" json:"matched_option_len,omitempty"`
	OptionData           []uint32 `protobuf:"varint,5,rep,packed,name=option_data,json=optionData,proto3" json:"option_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) Reset()         { *m = Ipv4DhcpdBaseChildProfileInfoItem{} }
func (m *Ipv4DhcpdBaseChildProfileInfoItem) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseChildProfileInfoItem) ProtoMessage()    {}
func (*Ipv4DhcpdBaseChildProfileInfoItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{3}
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseChildProfileInfoItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseChildProfileInfoItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem.Merge(m, src)
}
func (m *Ipv4DhcpdBaseChildProfileInfoItem) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem.Size(m)
}
func (m *Ipv4DhcpdBaseChildProfileInfoItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoItem proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseChildProfileInfoItem) GetBaseChildProfileName() string {
	if m != nil {
		return m.BaseChildProfileName
	}
	return ""
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) GetMatchedOptionCode() uint32 {
	if m != nil {
		return m.MatchedOptionCode
	}
	return 0
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) GetMatchedOptionLen() uint32 {
	if m != nil {
		return m.MatchedOptionLen
	}
	return 0
}

func (m *Ipv4DhcpdBaseChildProfileInfoItem) GetOptionData() []uint32 {
	if m != nil {
		return m.OptionData
	}
	return nil
}

type Ipv4DhcpdBaseChildProfileInfoEntry struct {
	Ipv4DhcpdBaseChildProfileInfo []*Ipv4DhcpdBaseChildProfileInfoItem `protobuf:"bytes,6,rep,name=ipv4_dhcpd_base_child_profile_info,json=ipv4DhcpdBaseChildProfileInfo,proto3" json:"ipv4_dhcpd_base_child_profile_info,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                             `json:"-"`
	XXX_unrecognized              []byte                               `json:"-"`
	XXX_sizecache                 int32                                `json:"-"`
}

func (m *Ipv4DhcpdBaseChildProfileInfoEntry) Reset()         { *m = Ipv4DhcpdBaseChildProfileInfoEntry{} }
func (m *Ipv4DhcpdBaseChildProfileInfoEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseChildProfileInfoEntry) ProtoMessage()    {}
func (*Ipv4DhcpdBaseChildProfileInfoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{4}
}

func (m *Ipv4DhcpdBaseChildProfileInfoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseChildProfileInfoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseChildProfileInfoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry.Merge(m, src)
}
func (m *Ipv4DhcpdBaseChildProfileInfoEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry.Size(m)
}
func (m *Ipv4DhcpdBaseChildProfileInfoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseChildProfileInfoEntry proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseChildProfileInfoEntry) GetIpv4DhcpdBaseChildProfileInfo() []*Ipv4DhcpdBaseChildProfileInfoItem {
	if m != nil {
		return m.Ipv4DhcpdBaseChildProfileInfo
	}
	return nil
}

type Ipv4DhcpdBaseProfile struct {
	BaseDefaultProfileName string                                `protobuf:"bytes,50,opt,name=base_default_profile_name,json=baseDefaultProfileName,proto3" json:"base_default_profile_name,omitempty"`
	DefaultProfileMode     uint32                                `protobuf:"varint,51,opt,name=default_profile_mode,json=defaultProfileMode,proto3" json:"default_profile_mode,omitempty"`
	RelayAuthenticate      string                                `protobuf:"bytes,52,opt,name=relay_authenticate,json=relayAuthenticate,proto3" json:"relay_authenticate,omitempty"`
	RemoteId               string                                `protobuf:"bytes,53,opt,name=remote_id,json=remoteId,proto3" json:"remote_id,omitempty"`
	ChildProfileCount      uint32                                `protobuf:"varint,54,opt,name=child_profile_count,json=childProfileCount,proto3" json:"child_profile_count,omitempty"`
	IntfRefCount           uint32                                `protobuf:"varint,55,opt,name=intf_ref_count,json=intfRefCount,proto3" json:"intf_ref_count,omitempty"`
	InterfaceReferences    *Ipv4DhcpdBaseInterfaceReferenceEntry `protobuf:"bytes,56,opt,name=interface_references,json=interfaceReferences,proto3" json:"interface_references,omitempty"`
	ChildProfileInfo       *Ipv4DhcpdBaseChildProfileInfoEntry   `protobuf:"bytes,57,opt,name=child_profile_info,json=childProfileInfo,proto3" json:"child_profile_info,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                              `json:"-"`
	XXX_unrecognized       []byte                                `json:"-"`
	XXX_sizecache          int32                                 `json:"-"`
}

func (m *Ipv4DhcpdBaseProfile) Reset()         { *m = Ipv4DhcpdBaseProfile{} }
func (m *Ipv4DhcpdBaseProfile) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdBaseProfile) ProtoMessage()    {}
func (*Ipv4DhcpdBaseProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9dd6d4f59f5ecf3, []int{5}
}

func (m *Ipv4DhcpdBaseProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile.Unmarshal(m, b)
}
func (m *Ipv4DhcpdBaseProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdBaseProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdBaseProfile.Merge(m, src)
}
func (m *Ipv4DhcpdBaseProfile) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdBaseProfile.Size(m)
}
func (m *Ipv4DhcpdBaseProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdBaseProfile.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdBaseProfile proto.InternalMessageInfo

func (m *Ipv4DhcpdBaseProfile) GetBaseDefaultProfileName() string {
	if m != nil {
		return m.BaseDefaultProfileName
	}
	return ""
}

func (m *Ipv4DhcpdBaseProfile) GetDefaultProfileMode() uint32 {
	if m != nil {
		return m.DefaultProfileMode
	}
	return 0
}

func (m *Ipv4DhcpdBaseProfile) GetRelayAuthenticate() string {
	if m != nil {
		return m.RelayAuthenticate
	}
	return ""
}

func (m *Ipv4DhcpdBaseProfile) GetRemoteId() string {
	if m != nil {
		return m.RemoteId
	}
	return ""
}

func (m *Ipv4DhcpdBaseProfile) GetChildProfileCount() uint32 {
	if m != nil {
		return m.ChildProfileCount
	}
	return 0
}

func (m *Ipv4DhcpdBaseProfile) GetIntfRefCount() uint32 {
	if m != nil {
		return m.IntfRefCount
	}
	return 0
}

func (m *Ipv4DhcpdBaseProfile) GetInterfaceReferences() *Ipv4DhcpdBaseInterfaceReferenceEntry {
	if m != nil {
		return m.InterfaceReferences
	}
	return nil
}

func (m *Ipv4DhcpdBaseProfile) GetChildProfileInfo() *Ipv4DhcpdBaseChildProfileInfoEntry {
	if m != nil {
		return m.ChildProfileInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*Ipv4DhcpdBaseProfile_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_profile_KEYS")
	proto.RegisterType((*Ipv4DhcpdBaseInterfaceReferenceItem)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_interface_reference_item")
	proto.RegisterType((*Ipv4DhcpdBaseInterfaceReferenceEntry)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_interface_reference_entry")
	proto.RegisterType((*Ipv4DhcpdBaseChildProfileInfoItem)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_child_profile_info_item")
	proto.RegisterType((*Ipv4DhcpdBaseChildProfileInfoEntry)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_child_profile_info_entry")
	proto.RegisterType((*Ipv4DhcpdBaseProfile)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.profiles.profile.ipv4_dhcpd_base_profile")
}

func init() { proto.RegisterFile("ipv4_dhcpd_base_profile.proto", fileDescriptor_e9dd6d4f59f5ecf3) }

var fileDescriptor_e9dd6d4f59f5ecf3 = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xe5, 0x36, 0xbf, 0xe8, 0xd7, 0x49, 0x8b, 0xda, 0x6d, 0xd4, 0x9a, 0x3f, 0x51, 0x83,
	0x41, 0x22, 0x48, 0x60, 0xa1, 0xb4, 0x05, 0x7a, 0x2c, 0x09, 0x12, 0x11, 0x7f, 0x65, 0x4e, 0x3d,
	0xad, 0xb6, 0xeb, 0xb1, 0xb2, 0x52, 0xbc, 0x1b, 0xd6, 0x1b, 0x44, 0xdf, 0x86, 0x13, 0x4f, 0xc1,
	0x81, 0xd7, 0xe1, 0xc2, 0x05, 0x1e, 0x00, 0xed, 0xda, 0x49, 0x9c, 0x34, 0x11, 0x39, 0x40, 0x2f,
	0x91, 0x3d, 0xdf, 0xcf, 0xce, 0x78, 0xe7, 0x3b, 0xa3, 0x40, 0x43, 0x0c, 0x3f, 0x1e, 0xd1, 0xb8,
	0xcf, 0x87, 0x31, 0x3d, 0x67, 0x19, 0xd2, 0xa1, 0x56, 0x89, 0x18, 0x60, 0x38, 0xd4, 0xca, 0x28,
	0xf2, 0x82, 0x8b, 0x8c, 0x2b, 0x2a, 0x54, 0x46, 0x3f, 0x69, 0x5a, 0x62, 0xd5, 0x10, 0x75, 0x38,
	0x7d, 0x0f, 0xa5, 0x8a, 0x31, 0x73, 0xbf, 0xa1, 0x4d, 0x13, 0x16, 0x69, 0xb2, 0xf1, 0x43, 0x70,
	0x06, 0xb7, 0x96, 0x94, 0xa2, 0x2f, 0x9f, 0x9f, 0xbd, 0x27, 0x7b, 0x50, 0xb5, 0xa7, 0x45, 0xec,
	0x7b, 0x4d, 0xaf, 0xb5, 0x11, 0x15, 0x6f, 0xe4, 0x36, 0x6c, 0x8e, 0x39, 0xc9, 0x52, 0xf4, 0xd7,
	0x9c, 0x5a, 0x2b, 0x62, 0x6f, 0x58, 0x8a, 0x41, 0x0a, 0xad, 0xf9, 0xd4, 0x42, 0x1a, 0xd4, 0x09,
	0xe3, 0x48, 0x35, 0x26, 0xa8, 0x51, 0x72, 0xa4, 0xc2, 0x60, 0x4a, 0x4e, 0xa1, 0xe1, 0x80, 0x52,
	0x78, 0x82, 0xba, 0xfc, 0x79, 0xf5, 0x1b, 0x16, 0x8a, 0xc6, 0x4c, 0x6f, 0x8c, 0xb8, 0x72, 0x3f,
	0x3c, 0xb8, 0xbf, 0x4a, 0x3d, 0x94, 0x46, 0x5f, 0x90, 0x6f, 0x1e, 0xdc, 0x59, 0x81, 0xf6, 0xd7,
	0x9a, 0xeb, 0xad, 0x5a, 0x5b, 0x87, 0x7f, 0xab, 0xe1, 0xe1, 0xaa, 0x2d, 0x89, 0x0e, 0x2c, 0xd9,
	0xb5, 0xe0, 0x33, 0x96, 0x4d, 0xef, 0x3a, 0xb9, 0x7d, 0xf0, 0xd3, 0x83, 0x7b, 0xf3, 0xd9, 0x78,
	0x5f, 0x0c, 0xe2, 0x89, 0x83, 0x42, 0x26, 0x2a, 0xef, 0xef, 0x31, 0xec, 0x2f, 0xd0, 0x4b, 0x9d,
	0xad, 0x5b, 0xb9, 0x63, 0xd5, 0x77, 0x53, 0x0b, 0x09, 0x81, 0x4a, 0xaa, 0xe2, 0xdc, 0xdd, 0xad,
	0xc8, 0x3d, 0x93, 0x10, 0x76, 0x53, 0x66, 0x78, 0x1f, 0xed, 0xe5, 0x8d, 0x50, 0x92, 0x72, 0x8b,
	0xac, 0x3b, 0x64, 0xa7, 0x90, 0xde, 0x3a, 0xa5, 0x63, 0xf9, 0x07, 0x40, 0xe6, 0xf8, 0x01, 0x4a,
	0xbf, 0xe2, 0xf0, 0xed, 0x19, 0xfc, 0x15, 0x4a, 0x72, 0x00, 0xb5, 0x82, 0x8a, 0x99, 0x61, 0xfe,
	0x7f, 0xcd, 0xf5, 0xd6, 0x56, 0x04, 0x79, 0xa8, 0xcb, 0x0c, 0x0b, 0xbe, 0x7b, 0x97, 0xc7, 0x6a,
	0xc1, 0xad, 0x73, 0x97, 0xbf, 0x7a, 0x10, 0xfc, 0x19, 0xf6, 0xab, 0xce, 0xe4, 0x0f, 0xff, 0xce,
	0xe4, 0x25, 0xb6, 0x44, 0x8d, 0x19, 0x8f, 0xcb, 0xed, 0xef, 0xc9, 0x44, 0x05, 0xbf, 0x2a, 0xb0,
	0xbf, 0x64, 0x3b, 0xc9, 0x09, 0x5c, 0x77, 0xef, 0x31, 0x26, 0x6c, 0x34, 0x30, 0xb3, 0x9e, 0xb6,
	0x9d, 0xa7, 0x7b, 0x16, 0xe8, 0xe6, 0x7a, 0xd9, 0xd5, 0x47, 0x50, 0x9f, 0x3f, 0xe5, 0x5c, 0x3e,
	0x74, 0x9e, 0x90, 0x78, 0xe6, 0xc4, 0x6b, 0xeb, 0xe1, 0x43, 0x20, 0x1a, 0x07, 0xec, 0x82, 0xb2,
	0x91, 0xe9, 0xa3, 0x34, 0x82, 0x33, 0x83, 0xfe, 0x91, 0xab, 0xb2, 0xe3, 0x94, 0xd3, 0x92, 0x40,
	0x6e, 0xc2, 0x86, 0xc6, 0x54, 0x19, 0xa4, 0x22, 0xf6, 0x8f, 0x1d, 0xf5, 0x7f, 0x1e, 0xe8, 0xc5,
	0x76, 0x7e, 0x66, 0xdb, 0xc1, 0xd5, 0x48, 0x1a, 0xff, 0x71, 0x3e, 0x3f, 0xbc, 0xd4, 0x83, 0x8e,
	0x15, 0xc8, 0x5d, 0xb8, 0x26, 0xa4, 0x49, 0xec, 0x7a, 0x14, 0xe8, 0x13, 0x87, 0x6e, 0xda, 0x68,
	0x84, 0x49, 0x4e, 0x7d, 0xf1, 0xa0, 0xbe, 0x60, 0x95, 0x32, 0xff, 0x69, 0xd3, 0x6b, 0xd5, 0xda,
	0xd9, 0xd5, 0x2e, 0xb0, 0x9b, 0xbe, 0x68, 0x57, 0x5c, 0x5a, 0xda, 0x8c, 0x7c, 0xf6, 0x80, 0x2c,
	0x18, 0xc1, 0x13, 0xf7, 0x99, 0xfa, 0x4a, 0x47, 0x30, 0xff, 0xca, 0x6d, 0x3e, 0x37, 0x76, 0xe7,
	0x55, 0xf7, 0x27, 0x73, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x77, 0xf5, 0x75, 0xbc, 0x85, 0x06,
	0x00, 0x00,
}
