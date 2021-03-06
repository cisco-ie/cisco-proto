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

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_bindings_binding

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
	proto.RegisterType((*LdpTibEntryDetail_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_tib_entry_detail_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_vrf_info")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpLdpidInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_ldpid_info")
	proto.RegisterType((*LdpTibEntryRemote)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_tib_entry_remote")
	proto.RegisterType((*LdpTibEntryDetail)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.bindings.binding.ldp_tib_entry_detail")
}

func init() { proto.RegisterFile("ldp_tib_entry_detail.proto", fileDescriptor_fdbd261a6df5ce59) }

var fileDescriptor_fdbd261a6df5ce59 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0x96, 0x93, 0x26, 0xb5, 0x9f, 0x63, 0x43, 0xa6, 0x6e, 0x99, 0x80, 0x80, 0xd6, 0x70, 0xa8,
	0x38, 0x2c, 0xc8, 0x2d, 0x81, 0xf2, 0xbb, 0x48, 0x91, 0x88, 0xb0, 0x4a, 0xb5, 0xae, 0x2a, 0x38,
	0x8d, 0xc6, 0x3b, 0x6f, 0xcd, 0xa8, 0xe3, 0x99, 0xd5, 0xcc, 0x78, 0x95, 0xdc, 0xb8, 0xc1, 0x85,
	0x13, 0x12, 0x57, 0xfe, 0x55, 0x34, 0x3f, 0xb2, 0x4d, 0xa2, 0x5e, 0x9d, 0x93, 0x67, 0xbe, 0xf7,
	0xbd, 0xfd, 0xbe, 0x79, 0xfb, 0xcd, 0x1a, 0xde, 0x55, 0xa2, 0x61, 0x5e, 0x2e, 0x19, 0x6a, 0x6f,
	0xcf, 0x99, 0x40, 0xcf, 0xa5, 0x2a, 0x1a, 0x6b, 0xbc, 0x21, 0x2f, 0x2b, 0xe9, 0x2a, 0xc3, 0xa4,
	0x71, 0xec, 0xcc, 0xb2, 0x75, 0xa3, 0x1c, 0x0b, 0x6c, 0xd3, 0xa0, 0x2d, 0x2e, 0x76, 0xc5, 0x4a,
	0x99, 0x25, 0x57, 0x85, 0xf3, 0x5c, 0x8b, 0xe5, 0x79, 0x21, 0xb0, 0xe6, 0x1b, 0xe5, 0x59, 0x6b,
	0xeb, 0x82, 0xd7, 0xae, 0xe0, 0x75, 0xb1, 0x94, 0x5a, 0x48, 0xbd, 0x72, 0x17, 0x8b, 0xe9, 0x1c,
	0x8e, 0xde, 0xa4, 0xca, 0x7e, 0x3e, 0xf9, 0x6d, 0x41, 0xde, 0x81, 0xdb, 0xbc, 0x66, 0x9a, 0xaf,
	0x91, 0xf6, 0xee, 0xf7, 0x1e, 0x0e, 0xca, 0x7d, 0x5e, 0x3f, 0xe3, 0x6b, 0x24, 0xf7, 0x60, 0xbf,
	0xb1, 0x58, 0xcb, 0x33, 0xba, 0x93, 0xf0, 0xb4, 0x9b, 0xce, 0xe0, 0x20, 0x3c, 0xad, 0xb5, 0x35,
	0x93, 0xba, 0x36, 0x84, 0xc0, 0xad, 0x4b, 0xdd, 0x71, 0x4d, 0xc6, 0xb0, 0x23, 0x45, 0xec, 0x1b,
	0x95, 0x3b, 0x52, 0x4c, 0x11, 0xee, 0x84, 0x1e, 0xd9, 0x30, 0x2e, 0x84, 0x65, 0x9e, 0x6d, 0xb4,
	0x34, 0x9a, 0xbc, 0x0d, 0xbb, 0xbc, 0x96, 0xb9, 0x33, 0x2c, 0xc9, 0x04, 0xf6, 0xc4, 0x66, 0xbd,
	0x3e, 0xcf, 0xbd, 0x69, 0x13, 0x24, 0x64, 0xd3, 0x3e, 0xa6, 0xbb, 0x49, 0x22, 0xac, 0x33, 0x76,
	0x4c, 0x6f, 0x75, 0xd8, 0xf1, 0x54, 0xc0, 0x38, 0xc8, 0x28, 0xd1, 0x48, 0x91, 0xcc, 0xdd, 0x85,
	0x7d, 0xe5, 0x2c, 0x93, 0x22, 0x8b, 0xec, 0x29, 0x67, 0x4f, 0x05, 0xf9, 0x18, 0xc6, 0x8a, 0x2f,
	0x51, 0x31, 0xd7, 0xf0, 0x0a, 0x59, 0xe7, 0xf5, 0x20, 0xa2, 0x8b, 0x00, 0x9e, 0x8a, 0xd8, 0x1c,
	0x5c, 0x8b, 0x2c, 0xbc, 0xa7, 0x44, 0x73, 0x2a, 0xa6, 0x7f, 0xef, 0xc0, 0xe4, 0xea, 0x3c, 0x2d,
	0xae, 0x8d, 0x47, 0xf2, 0x00, 0x0e, 0xd2, 0x8a, 0xc5, 0xc7, 0x44, 0xc9, 0x51, 0x39, 0x4c, 0xd8,
	0x3c, 0x40, 0xe4, 0xbf, 0x1e, 0x50, 0xee, 0x9c, 0x5c, 0x69, 0xa9, 0x57, 0xac, 0x41, 0xb4, 0x2c,
	0x49, 0xa0, 0xf6, 0xd1, 0xc3, 0x70, 0x56, 0x17, 0xdb, 0x89, 0x41, 0x71, 0x75, 0x34, 0xe5, 0xdd,
	0xce, 0xc7, 0x73, 0x44, 0x3b, 0x0f, 0x47, 0x43, 0xed, 0xc9, 0x11, 0xf4, 0xa5, 0x63, 0xce, 0x73,
	0x85, 0xf1, 0xd8, 0xfd, 0xf2, 0xb6, 0x74, 0x8b, 0xb0, 0x0d, 0xf3, 0x90, 0x8e, 0xa1, 0xaa, 0xe2,
	0xd0, 0xfb, 0xe5, 0x9e, 0x74, 0x27, 0xaa, 0x9a, 0xfe, 0x31, 0xb8, 0x3e, 0x8f, 0x94, 0x2f, 0xd2,
	0xc2, 0x6e, 0x6b, 0x6b, 0x3a, 0x8b, 0xc7, 0x12, 0xdb, 0x3c, 0xd6, 0x45, 0x18, 0xcb, 0x20, 0x48,
	0xfe, 0xea, 0xc1, 0x20, 0x85, 0x95, 0x9d, 0x59, 0xfa, 0x28, 0xca, 0xbf, 0xda, 0xa6, 0xfc, 0xb5,
	0x5c, 0x97, 0xfd, 0xa4, 0xfe, 0xab, 0x25, 0x1f, 0xc1, 0x28, 0x3b, 0x51, 0xa8, 0x57, 0xfe, 0x77,
	0xfa, 0x38, 0xe5, 0x2c, 0x81, 0xf3, 0x88, 0x91, 0x0f, 0x61, 0xa8, 0x4c, 0xc5, 0x55, 0x8e, 0xcd,
	0xe7, 0x91, 0x02, 0x11, 0x4a, 0xa9, 0x79, 0x02, 0x47, 0x0a, 0x59, 0xe2, 0x64, 0x5d, 0x66, 0xb1,
	0x95, 0x4e, 0x1a, 0x4d, 0x8f, 0x23, 0xfd, 0x9e, 0xc2, 0x79, 0xa8, 0xff, 0x98, 0xca, 0x65, 0xae,
	0x92, 0x4f, 0x61, 0xd2, 0xb5, 0xe6, 0xc8, 0x7b, 0xee, 0x91, 0x7e, 0x11, 0x13, 0x7d, 0x98, 0xbb,
	0xa2, 0xcc, 0x22, 0x14, 0xc8, 0x07, 0x30, 0x94, 0x8e, 0x69, 0xc3, 0xac, 0xd9, 0x78, 0xa4, 0x5f,
	0xc6, 0x37, 0x3d, 0x90, 0xee, 0x99, 0x29, 0x03, 0x40, 0xde, 0x83, 0x41, 0x7a, 0x8e, 0x31, 0x96,
	0x3e, 0x89, 0xd5, 0x7e, 0x04, 0x7e, 0x31, 0x96, 0x7c, 0x06, 0x13, 0x2e, 0x5a, 0xb4, 0x5e, 0x3a,
	0x64, 0xf9, 0xe0, 0xbc, 0x52, 0xf4, 0xab, 0xa8, 0x46, 0xba, 0xda, 0xf3, 0x58, 0x7a, 0x5a, 0x29,
	0xf2, 0x09, 0x1c, 0xbe, 0xee, 0xf0, 0xce, 0x46, 0xfa, 0xd7, 0x91, 0xfe, 0x56, 0x57, 0x78, 0xe1,
	0x6c, 0xe0, 0xfe, 0xd3, 0x83, 0x71, 0xbe, 0x60, 0x79, 0x0a, 0xf4, 0x9b, 0xfb, 0xbb, 0x0f, 0x87,
	0x33, 0xb5, 0xcd, 0x97, 0x7b, 0xfd, 0x9a, 0x97, 0xa3, 0xf4, 0x9b, 0x27, 0x4d, 0xfe, 0xed, 0xc1,
	0x9d, 0x70, 0x91, 0x1d, 0xeb, 0xfc, 0x0a, 0xe6, 0x0d, 0xfd, 0x36, 0x5a, 0xbb, 0xa9, 0xdb, 0x7c,
	0x18, 0x2d, 0x3c, 0xed, 0x1c, 0xbc, 0x30, 0xe4, 0xcf, 0x1e, 0x0c, 0xb3, 0xb1, 0xea, 0x15, 0x0a,
	0xfa, 0xdd, 0x8d, 0x1a, 0x82, 0x64, 0x28, 0x28, 0x93, 0x9f, 0xe0, 0x41, 0x65, 0x74, 0x2d, 0x57,
	0x0c, 0x75, 0x6d, 0x6c, 0x85, 0xe2, 0x4a, 0x22, 0x5b, 0xae, 0x36, 0x48, 0xbf, 0x8f, 0x59, 0x7a,
	0x3f, 0x11, 0x4f, 0x32, 0xef, 0x75, 0x3a, 0x5f, 0x06, 0xd2, 0xa5, 0x4f, 0xd0, 0x0f, 0x97, 0x3e,
	0x41, 0xcb, 0xfd, 0xf8, 0x07, 0xfa, 0xe8, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa0, 0x5a,
	0x80, 0x5e, 0x07, 0x00, 0x00,
}
