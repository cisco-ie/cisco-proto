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
// source: ldp_tib_entry_detail.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_afs_af_bindings_binding

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

type LdpTibEntryDetail_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpTibEntryDetail_KEYS) Reset()         { *m = LdpTibEntryDetail_KEYS{} }
func (m *LdpTibEntryDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpTibEntryDetail_KEYS) ProtoMessage()    {}
func (*LdpTibEntryDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbd261a6df5ce59, []int{0}
}

func (m *LdpTibEntryDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTibEntryDetail_KEYS.Unmarshal(m, b)
}
func (m *LdpTibEntryDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTibEntryDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpTibEntryDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTibEntryDetail_KEYS.Merge(m, src)
}
func (m *LdpTibEntryDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpTibEntryDetail_KEYS.Size(m)
}
func (m *LdpTibEntryDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTibEntryDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTibEntryDetail_KEYS proto.InternalMessageInfo

func (m *LdpTibEntryDetail_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpTibEntryDetail_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
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
	return fileDescriptor_fdbd261a6df5ce59, []int{1}
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
	return fileDescriptor_fdbd261a6df5ce59, []int{2}
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

type LdpLdpidInfo struct {
	LsrId                string   `protobuf:"bytes,1,opt,name=lsr_id,json=lsrId,proto3" json:"lsr_id,omitempty"`
	LabelSpaceId         uint32   `protobuf:"varint,2,opt,name=label_space_id,json=labelSpaceId,proto3" json:"label_space_id,omitempty"`
	LdpId                string   `protobuf:"bytes,3,opt,name=ldp_id,json=ldpId,proto3" json:"ldp_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpLdpidInfo) Reset()         { *m = LdpLdpidInfo{} }
func (m *LdpLdpidInfo) String() string { return proto.CompactTextString(m) }
func (*LdpLdpidInfo) ProtoMessage()    {}
func (*LdpLdpidInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbd261a6df5ce59, []int{3}
}

func (m *LdpLdpidInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLdpidInfo.Unmarshal(m, b)
}
func (m *LdpLdpidInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLdpidInfo.Marshal(b, m, deterministic)
}
func (m *LdpLdpidInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLdpidInfo.Merge(m, src)
}
func (m *LdpLdpidInfo) XXX_Size() int {
	return xxx_messageInfo_LdpLdpidInfo.Size(m)
}
func (m *LdpLdpidInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLdpidInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLdpidInfo proto.InternalMessageInfo

func (m *LdpLdpidInfo) GetLsrId() string {
	if m != nil {
		return m.LsrId
	}
	return ""
}

func (m *LdpLdpidInfo) GetLabelSpaceId() uint32 {
	if m != nil {
		return m.LabelSpaceId
	}
	return 0
}

func (m *LdpLdpidInfo) GetLdpId() string {
	if m != nil {
		return m.LdpId
	}
	return ""
}

type LdpTibEntryRemote struct {
	RemoteLabel           uint32        `protobuf:"varint,1,opt,name=remote_label,json=remoteLabel,proto3" json:"remote_label,omitempty"`
	AssigningPeerLdpIdent *LdpLdpidInfo `protobuf:"bytes,2,opt,name=assigning_peer_ldp_ident,json=assigningPeerLdpIdent,proto3" json:"assigning_peer_ldp_ident,omitempty"`
	IsStale               bool          `protobuf:"varint,3,opt,name=is_stale,json=isStale,proto3" json:"is_stale,omitempty"`
	IsElc                 bool          `protobuf:"varint,4,opt,name=is_elc,json=isElc,proto3" json:"is_elc,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}      `json:"-"`
	XXX_unrecognized      []byte        `json:"-"`
	XXX_sizecache         int32         `json:"-"`
}

func (m *LdpTibEntryRemote) Reset()         { *m = LdpTibEntryRemote{} }
func (m *LdpTibEntryRemote) String() string { return proto.CompactTextString(m) }
func (*LdpTibEntryRemote) ProtoMessage()    {}
func (*LdpTibEntryRemote) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbd261a6df5ce59, []int{4}
}

func (m *LdpTibEntryRemote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTibEntryRemote.Unmarshal(m, b)
}
func (m *LdpTibEntryRemote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTibEntryRemote.Marshal(b, m, deterministic)
}
func (m *LdpTibEntryRemote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTibEntryRemote.Merge(m, src)
}
func (m *LdpTibEntryRemote) XXX_Size() int {
	return xxx_messageInfo_LdpTibEntryRemote.Size(m)
}
func (m *LdpTibEntryRemote) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTibEntryRemote.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTibEntryRemote proto.InternalMessageInfo

func (m *LdpTibEntryRemote) GetRemoteLabel() uint32 {
	if m != nil {
		return m.RemoteLabel
	}
	return 0
}

func (m *LdpTibEntryRemote) GetAssigningPeerLdpIdent() *LdpLdpidInfo {
	if m != nil {
		return m.AssigningPeerLdpIdent
	}
	return nil
}

func (m *LdpTibEntryRemote) GetIsStale() bool {
	if m != nil {
		return m.IsStale
	}
	return false
}

func (m *LdpTibEntryRemote) GetIsElc() bool {
	if m != nil {
		return m.IsElc
	}
	return false
}

type LdpTibEntryDetail struct {
	Vrf                           *LdpVrfInfo          `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	PrefixXr                      *LdpIpAddrTUnion     `protobuf:"bytes,51,opt,name=prefix_xr,json=prefixXr,proto3" json:"prefix_xr,omitempty"`
	PrefixLength                  uint32               `protobuf:"varint,52,opt,name=prefix_length,json=prefixLength,proto3" json:"prefix_length,omitempty"`
	LocalLabel                    uint32               `protobuf:"varint,53,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	LeLocalBindingRevision        uint32               `protobuf:"varint,54,opt,name=le_local_binding_revision,json=leLocalBindingRevision,proto3" json:"le_local_binding_revision,omitempty"`
	LeLocalLabelState             string               `protobuf:"bytes,55,opt,name=le_local_label_state,json=leLocalLabelState,proto3" json:"le_local_label_state,omitempty"`
	IsNoRoute                     bool                 `protobuf:"varint,56,opt,name=is_no_route,json=isNoRoute,proto3" json:"is_no_route,omitempty"`
	LabelOor                      bool                 `protobuf:"varint,57,opt,name=label_oor,json=labelOor,proto3" json:"label_oor,omitempty"`
	AdvertisePrefixAcl            string               `protobuf:"bytes,58,opt,name=advertise_prefix_acl,json=advertisePrefixAcl,proto3" json:"advertise_prefix_acl,omitempty"`
	AdvertiseTsrAcl               string               `protobuf:"bytes,59,opt,name=advertise_tsr_acl,json=advertiseTsrAcl,proto3" json:"advertise_tsr_acl,omitempty"`
	RemoteBinding                 []*LdpTibEntryRemote `protobuf:"bytes,60,rep,name=remote_binding,json=remoteBinding,proto3" json:"remote_binding,omitempty"`
	PeersAdvertisedTo             []*LdpLdpidInfo      `protobuf:"bytes,61,rep,name=peers_advertised_to,json=peersAdvertisedTo,proto3" json:"peers_advertised_to,omitempty"`
	PeersAcked                    []*LdpLdpidInfo      `protobuf:"bytes,62,rep,name=peers_acked,json=peersAcked,proto3" json:"peers_acked,omitempty"`
	ConfigEnforcedLocalLabelValue bool                 `protobuf:"varint,63,opt,name=config_enforced_local_label_value,json=configEnforcedLocalLabelValue,proto3" json:"config_enforced_local_label_value,omitempty"`
	IsElc                         bool                 `protobuf:"varint,64,opt,name=is_elc,json=isElc,proto3" json:"is_elc,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *LdpTibEntryDetail) Reset()         { *m = LdpTibEntryDetail{} }
func (m *LdpTibEntryDetail) String() string { return proto.CompactTextString(m) }
func (*LdpTibEntryDetail) ProtoMessage()    {}
func (*LdpTibEntryDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdbd261a6df5ce59, []int{5}
}

func (m *LdpTibEntryDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTibEntryDetail.Unmarshal(m, b)
}
func (m *LdpTibEntryDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTibEntryDetail.Marshal(b, m, deterministic)
}
func (m *LdpTibEntryDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTibEntryDetail.Merge(m, src)
}
func (m *LdpTibEntryDetail) XXX_Size() int {
	return xxx_messageInfo_LdpTibEntryDetail.Size(m)
}
func (m *LdpTibEntryDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTibEntryDetail.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTibEntryDetail proto.InternalMessageInfo

func (m *LdpTibEntryDetail) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpTibEntryDetail) GetPrefixXr() *LdpIpAddrTUnion {
	if m != nil {
		return m.PrefixXr
	}
	return nil
}

func (m *LdpTibEntryDetail) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

func (m *LdpTibEntryDetail) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *LdpTibEntryDetail) GetLeLocalBindingRevision() uint32 {
	if m != nil {
		return m.LeLocalBindingRevision
	}
	return 0
}

func (m *LdpTibEntryDetail) GetLeLocalLabelState() string {
	if m != nil {
		return m.LeLocalLabelState
	}
	return ""
}

func (m *LdpTibEntryDetail) GetIsNoRoute() bool {
	if m != nil {
		return m.IsNoRoute
	}
	return false
}

func (m *LdpTibEntryDetail) GetLabelOor() bool {
	if m != nil {
		return m.LabelOor
	}
	return false
}

func (m *LdpTibEntryDetail) GetAdvertisePrefixAcl() string {
	if m != nil {
		return m.AdvertisePrefixAcl
	}
	return ""
}

func (m *LdpTibEntryDetail) GetAdvertiseTsrAcl() string {
	if m != nil {
		return m.AdvertiseTsrAcl
	}
	return ""
}

func (m *LdpTibEntryDetail) GetRemoteBinding() []*LdpTibEntryRemote {
	if m != nil {
		return m.RemoteBinding
	}
	return nil
}

func (m *LdpTibEntryDetail) GetPeersAdvertisedTo() []*LdpLdpidInfo {
	if m != nil {
		return m.PeersAdvertisedTo
	}
	return nil
}

func (m *LdpTibEntryDetail) GetPeersAcked() []*LdpLdpidInfo {
	if m != nil {
		return m.PeersAcked
	}
	return nil
}

func (m *LdpTibEntryDetail) GetConfigEnforcedLocalLabelValue() bool {
	if m != nil {
		return m.ConfigEnforcedLocalLabelValue
	}
	return false
}

func (m *LdpTibEntryDetail) GetIsElc() bool {
	if m != nil {
		return m.IsElc
	}
	return false
}

func init() {
	proto.RegisterType((*LdpTibEntryDetail_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_tib_entry_detail_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_vrf_info")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpLdpidInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_ldpid_info")
	proto.RegisterType((*LdpTibEntryRemote)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_tib_entry_remote")
	proto.RegisterType((*LdpTibEntryDetail)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.afs.af.bindings.binding.ldp_tib_entry_detail")
}

func init() { proto.RegisterFile("ldp_tib_entry_detail.proto", fileDescriptor_fdbd261a6df5ce59) }

var fileDescriptor_fdbd261a6df5ce59 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcf, 0x8f, 0x1b, 0x35,
	0x14, 0x56, 0x76, 0xbb, 0x69, 0xf2, 0xb2, 0x09, 0xac, 0x9b, 0x16, 0x2f, 0x08, 0x68, 0x03, 0x87,
	0x8a, 0xc3, 0x80, 0xd2, 0xb2, 0x50, 0x7e, 0x17, 0x69, 0x25, 0x56, 0x44, 0xa5, 0x9a, 0x14, 0x04,
	0x27, 0xcb, 0x19, 0xbf, 0x09, 0xa6, 0x8e, 0x3d, 0xb2, 0x9d, 0xd1, 0xf6, 0x1f, 0x80, 0x0b, 0x27,
	0x24, 0xae, 0xfc, 0xa3, 0x5c, 0x90, 0x7f, 0xec, 0x74, 0x77, 0xc5, 0x35, 0x3d, 0xc5, 0xfe, 0xde,
	0xf7, 0xe6, 0xfb, 0xfc, 0xe6, 0xf3, 0x04, 0xde, 0x54, 0xa2, 0x61, 0x5e, 0xae, 0x18, 0x6a, 0x6f,
	0x5f, 0x30, 0x81, 0x9e, 0x4b, 0x55, 0x34, 0xd6, 0x78, 0x43, 0x7e, 0xac, 0xa4, 0xab, 0x0c, 0x93,
	0xc6, 0xb1, 0x73, 0xcb, 0x36, 0x8d, 0x72, 0x2c, 0xb0, 0x4d, 0x83, 0xb6, 0xb8, 0xd8, 0x15, 0x6b,
	0x65, 0x56, 0x5c, 0x15, 0xbc, 0xf2, 0xb2, 0xc5, 0x42, 0x60, 0xcd, 0xb7, 0xca, 0xb3, 0xd6, 0xd6,
	0x05, 0xaf, 0x5d, 0xc1, 0xeb, 0x62, 0x25, 0xb5, 0x90, 0x7a, 0xed, 0x2e, 0x16, 0xb3, 0x05, 0x1c,
	0xff, 0x9f, 0x28, 0xfb, 0xfe, 0xf4, 0x97, 0x25, 0x79, 0x03, 0x6e, 0xf2, 0x9a, 0x69, 0xbe, 0x41,
	0xda, 0xbb, 0xdb, 0xbb, 0x3f, 0x2c, 0xfb, 0xbc, 0x7e, 0xc2, 0x37, 0x48, 0xee, 0x40, 0xbf, 0xb1,
	0x58, 0xcb, 0x73, 0xba, 0x97, 0xf0, 0xb4, 0x9b, 0xcd, 0xe1, 0x30, 0x3c, 0xad, 0xb5, 0x35, 0x93,
	0xba, 0x36, 0x84, 0xc0, 0x8d, 0x4b, 0xdd, 0x71, 0x4d, 0x26, 0xb0, 0x27, 0x45, 0xec, 0x1b, 0x97,
	0x7b, 0x52, 0xcc, 0x10, 0x6e, 0x85, 0x1e, 0xd9, 0x30, 0x2e, 0x84, 0x65, 0x9e, 0x6d, 0xb5, 0x34,
	0x9a, 0xbc, 0x0e, 0xfb, 0xbc, 0x96, 0xb9, 0x33, 0x2c, 0xc9, 0x14, 0x0e, 0xc4, 0x76, 0xb3, 0x79,
	0x91, 0x7b, 0xd3, 0x26, 0x48, 0xc8, 0xa6, 0x7d, 0x48, 0xf7, 0x93, 0x44, 0x58, 0x67, 0xec, 0x84,
	0xde, 0xe8, 0xb0, 0x93, 0x99, 0x80, 0x49, 0x90, 0x51, 0xa2, 0x91, 0x22, 0x99, 0xbb, 0x0d, 0x7d,
	0xe5, 0x2c, 0x93, 0x22, 0x8b, 0x1c, 0x28, 0x67, 0xcf, 0x04, 0x79, 0x1f, 0x26, 0x8a, 0xaf, 0x50,
	0x31, 0xd7, 0xf0, 0x0a, 0x59, 0xe7, 0xf5, 0x30, 0xa2, 0xcb, 0x00, 0x9e, 0x89, 0xd8, 0x1c, 0x5c,
	0x8b, 0x2c, 0x7c, 0xa0, 0x44, 0x73, 0x26, 0x66, 0x7f, 0xee, 0xc1, 0xf4, 0xea, 0x3c, 0x2d, 0x6e,
	0x8c, 0x47, 0x72, 0x0f, 0x0e, 0xd3, 0x8a, 0xc5, 0xc7, 0x44, 0xc9, 0x71, 0x39, 0x4a, 0xd8, 0x22,
	0x40, 0xe4, 0x9f, 0x1e, 0x50, 0xee, 0x9c, 0x5c, 0x6b, 0xa9, 0xd7, 0xac, 0x41, 0xb4, 0x2c, 0x49,
	0xa0, 0xf6, 0xd1, 0xc3, 0x68, 0x8e, 0xc5, 0x4e, 0x52, 0x50, 0x5c, 0x9d, 0x4c, 0x79, 0xbb, 0xb3,
	0xf1, 0x14, 0xd1, 0x2e, 0xc2, 0xc9, 0x50, 0x7b, 0x72, 0x0c, 0x03, 0xe9, 0x98, 0xf3, 0x5c, 0x61,
	0x3c, 0xf5, 0xa0, 0xbc, 0x29, 0xdd, 0x32, 0x6c, 0xc3, 0x38, 0xa4, 0x63, 0xa8, 0xaa, 0x38, 0xf3,
	0x41, 0x79, 0x20, 0xdd, 0xa9, 0xaa, 0x66, 0xff, 0x0e, 0xae, 0x8f, 0x23, 0xc5, 0x8b, 0x6c, 0x61,
	0xbf, 0xb5, 0x35, 0x9d, 0xc7, 0x53, 0x55, 0x3b, 0x3c, 0xd5, 0x45, 0x14, 0xcb, 0xa0, 0x47, 0xfe,
	0xe8, 0xc1, 0x30, 0x45, 0x95, 0x9d, 0x5b, 0xfa, 0x20, 0xaa, 0xff, 0xb6, 0x43, 0xf5, 0x6b, 0xa1,
	0x2e, 0x07, 0x49, 0xfc, 0x67, 0x4b, 0xde, 0x83, 0x71, 0x36, 0xa2, 0x50, 0xaf, 0xfd, 0xaf, 0xf4,
	0x61, 0x0a, 0x59, 0x02, 0x17, 0x11, 0x23, 0xef, 0xc2, 0x48, 0x99, 0x8a, 0xab, 0x9c, 0x99, 0x8f,
	0x23, 0x05, 0x22, 0x94, 0x22, 0xf3, 0x08, 0x8e, 0x15, 0xb2, 0xc4, 0xc9, 0xba, 0xcc, 0x62, 0x2b,
	0x9d, 0x34, 0x9a, 0x9e, 0x44, 0xfa, 0x1d, 0x85, 0x8b, 0x50, 0xff, 0x36, 0x95, 0xcb, 0x5c, 0x25,
	0x1f, 0xc2, 0xb4, 0x6b, 0xcd, 0x79, 0xf7, 0xdc, 0x23, 0xfd, 0x24, 0xc6, 0xf9, 0x28, 0x77, 0x45,
	0x99, 0x65, 0x28, 0x90, 0x77, 0x60, 0x24, 0x1d, 0xd3, 0x86, 0x59, 0xb3, 0xf5, 0x48, 0x3f, 0x8d,
	0xef, 0x79, 0x28, 0xdd, 0x13, 0x53, 0x06, 0x80, 0xbc, 0x05, 0xc3, 0xf4, 0x1c, 0x63, 0x2c, 0x7d,
	0x14, 0xab, 0x83, 0x08, 0xfc, 0x60, 0x2c, 0xf9, 0x08, 0xa6, 0x5c, 0xb4, 0x68, 0xbd, 0x74, 0xc8,
	0xf2, 0xc1, 0x79, 0xa5, 0xe8, 0x67, 0x51, 0x8d, 0x74, 0xb5, 0xa7, 0xb1, 0xf4, 0xb8, 0x52, 0xe4,
	0x03, 0x38, 0x7a, 0xd9, 0xe1, 0x9d, 0x8d, 0xf4, 0xcf, 0x23, 0xfd, 0xb5, 0xae, 0xf0, 0xcc, 0xd9,
	0xc0, 0xfd, 0xab, 0x07, 0x93, 0x7c, 0xbb, 0xf2, 0x14, 0xe8, 0x17, 0x77, 0xf7, 0xef, 0x8f, 0xe6,
	0xcf, 0x77, 0xf8, 0x6e, 0xaf, 0x5f, 0xf1, 0x72, 0x9c, 0x7e, 0xf3, 0xa0, 0xc9, 0xdf, 0x3d, 0xb8,
	0x15, 0x2e, 0xb1, 0x63, 0x9d, 0x5d, 0xc1, 0xbc, 0xa1, 0x5f, 0x46, 0x67, 0xaf, 0xe8, 0x26, 0x1f,
	0x45, 0x07, 0x8f, 0x3b, 0x03, 0xcf, 0x0c, 0xf9, 0xbd, 0x07, 0xa3, 0xec, 0xab, 0x7a, 0x8e, 0x82,
	0x7e, 0xf5, 0x2a, 0xfd, 0x40, 0xf2, 0x13, 0x84, 0xc9, 0x77, 0x70, 0xaf, 0x32, 0xba, 0x96, 0x6b,
	0x86, 0xba, 0x36, 0xb6, 0x42, 0x71, 0x25, 0x8e, 0x2d, 0x57, 0x5b, 0xa4, 0x5f, 0xc7, 0x20, 0xbd,
	0x9d, 0x88, 0xa7, 0x99, 0xf7, 0x32, 0x9a, 0x3f, 0x05, 0xd2, 0xa5, 0xaf, 0xcf, 0x37, 0x97, 0xbe,
	0x3e, 0xab, 0x7e, 0xfc, 0xe7, 0x7c, 0xf0, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xbd, 0x61,
	0xe0, 0x57, 0x07, 0x00, 0x00,
}
