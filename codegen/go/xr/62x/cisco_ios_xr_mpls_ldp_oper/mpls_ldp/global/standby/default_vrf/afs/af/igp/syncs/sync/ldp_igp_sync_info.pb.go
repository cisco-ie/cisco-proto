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
// source: ldp_igp_sync_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_standby_default_vrf_afs_af_igp_syncs_sync

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

type LdpIgpSyncInfo_KEYS struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIgpSyncInfo_KEYS) Reset()         { *m = LdpIgpSyncInfo_KEYS{} }
func (m *LdpIgpSyncInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncInfo_KEYS) ProtoMessage()    {}
func (*LdpIgpSyncInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebc3d2375525d14, []int{0}
}

func (m *LdpIgpSyncInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpIgpSyncInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncInfo_KEYS.Merge(m, src)
}
func (m *LdpIgpSyncInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncInfo_KEYS.Size(m)
}
func (m *LdpIgpSyncInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncInfo_KEYS proto.InternalMessageInfo

func (m *LdpIgpSyncInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpIgpSyncInfo_KEYS) GetInterfaceName() string {
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
	return fileDescriptor_9ebc3d2375525d14, []int{1}
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

type LdpIgpSyncIntfPeer struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	IsGrEnabled          bool     `protobuf:"varint,2,opt,name=is_gr_enabled,json=isGrEnabled,proto3" json:"is_gr_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIgpSyncIntfPeer) Reset()         { *m = LdpIgpSyncIntfPeer{} }
func (m *LdpIgpSyncIntfPeer) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncIntfPeer) ProtoMessage()    {}
func (*LdpIgpSyncIntfPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebc3d2375525d14, []int{2}
}

func (m *LdpIgpSyncIntfPeer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncIntfPeer.Unmarshal(m, b)
}
func (m *LdpIgpSyncIntfPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncIntfPeer.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncIntfPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncIntfPeer.Merge(m, src)
}
func (m *LdpIgpSyncIntfPeer) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncIntfPeer.Size(m)
}
func (m *LdpIgpSyncIntfPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncIntfPeer.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncIntfPeer proto.InternalMessageInfo

func (m *LdpIgpSyncIntfPeer) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *LdpIgpSyncIntfPeer) GetIsGrEnabled() bool {
	if m != nil {
		return m.IsGrEnabled
	}
	return false
}

type LdpIgpSyncIntfGrOnly struct {
	PeerId               string   `protobuf:"bytes,1,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	IsChkptCreated       bool     `protobuf:"varint,2,opt,name=is_chkpt_created,json=isChkptCreated,proto3" json:"is_chkpt_created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpIgpSyncIntfGrOnly) Reset()         { *m = LdpIgpSyncIntfGrOnly{} }
func (m *LdpIgpSyncIntfGrOnly) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncIntfGrOnly) ProtoMessage()    {}
func (*LdpIgpSyncIntfGrOnly) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebc3d2375525d14, []int{3}
}

func (m *LdpIgpSyncIntfGrOnly) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncIntfGrOnly.Unmarshal(m, b)
}
func (m *LdpIgpSyncIntfGrOnly) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncIntfGrOnly.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncIntfGrOnly) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncIntfGrOnly.Merge(m, src)
}
func (m *LdpIgpSyncIntfGrOnly) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncIntfGrOnly.Size(m)
}
func (m *LdpIgpSyncIntfGrOnly) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncIntfGrOnly.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncIntfGrOnly proto.InternalMessageInfo

func (m *LdpIgpSyncIntfGrOnly) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *LdpIgpSyncIntfGrOnly) GetIsChkptCreated() bool {
	if m != nil {
		return m.IsChkptCreated
	}
	return false
}

type LdpIgpSyncInfo struct {
	Vrf                  *LdpVrfInfo             `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	InterfaceNameXr      string                  `protobuf:"bytes,51,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	IgpSyncState         string                  `protobuf:"bytes,52,opt,name=igp_sync_state,json=igpSyncState,proto3" json:"igp_sync_state,omitempty"`
	IgpSyncDelay         uint32                  `protobuf:"varint,53,opt,name=igp_sync_delay,json=igpSyncDelay,proto3" json:"igp_sync_delay,omitempty"`
	IsDelayTimerRunning  bool                    `protobuf:"varint,54,opt,name=is_delay_timer_running,json=isDelayTimerRunning,proto3" json:"is_delay_timer_running,omitempty"`
	DelayTimerRemaining  uint32                  `protobuf:"varint,55,opt,name=delay_timer_remaining,json=delayTimerRemaining,proto3" json:"delay_timer_remaining,omitempty"`
	Peers                []*LdpIgpSyncIntfPeer   `protobuf:"bytes,56,rep,name=peers,proto3" json:"peers,omitempty"`
	GrOnlyPeer           []*LdpIgpSyncIntfGrOnly `protobuf:"bytes,57,rep,name=gr_only_peer,json=grOnlyPeer,proto3" json:"gr_only_peer,omitempty"`
	IgpSyncDownReason    string                  `protobuf:"bytes,58,opt,name=igp_sync_down_reason,json=igpSyncDownReason,proto3" json:"igp_sync_down_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LdpIgpSyncInfo) Reset()         { *m = LdpIgpSyncInfo{} }
func (m *LdpIgpSyncInfo) String() string { return proto.CompactTextString(m) }
func (*LdpIgpSyncInfo) ProtoMessage()    {}
func (*LdpIgpSyncInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ebc3d2375525d14, []int{4}
}

func (m *LdpIgpSyncInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpIgpSyncInfo.Unmarshal(m, b)
}
func (m *LdpIgpSyncInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpIgpSyncInfo.Marshal(b, m, deterministic)
}
func (m *LdpIgpSyncInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpIgpSyncInfo.Merge(m, src)
}
func (m *LdpIgpSyncInfo) XXX_Size() int {
	return xxx_messageInfo_LdpIgpSyncInfo.Size(m)
}
func (m *LdpIgpSyncInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpIgpSyncInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpIgpSyncInfo proto.InternalMessageInfo

func (m *LdpIgpSyncInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpIgpSyncInfo) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *LdpIgpSyncInfo) GetIgpSyncState() string {
	if m != nil {
		return m.IgpSyncState
	}
	return ""
}

func (m *LdpIgpSyncInfo) GetIgpSyncDelay() uint32 {
	if m != nil {
		return m.IgpSyncDelay
	}
	return 0
}

func (m *LdpIgpSyncInfo) GetIsDelayTimerRunning() bool {
	if m != nil {
		return m.IsDelayTimerRunning
	}
	return false
}

func (m *LdpIgpSyncInfo) GetDelayTimerRemaining() uint32 {
	if m != nil {
		return m.DelayTimerRemaining
	}
	return 0
}

func (m *LdpIgpSyncInfo) GetPeers() []*LdpIgpSyncIntfPeer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *LdpIgpSyncInfo) GetGrOnlyPeer() []*LdpIgpSyncIntfGrOnly {
	if m != nil {
		return m.GrOnlyPeer
	}
	return nil
}

func (m *LdpIgpSyncInfo) GetIgpSyncDownReason() string {
	if m != nil {
		return m.IgpSyncDownReason
	}
	return ""
}

func init() {
	proto.RegisterType((*LdpIgpSyncInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.syncs.sync.ldp_igp_sync_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.syncs.sync.ldp_vrf_info")
	proto.RegisterType((*LdpIgpSyncIntfPeer)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.syncs.sync.ldp_igp_sync_intf_peer")
	proto.RegisterType((*LdpIgpSyncIntfGrOnly)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.syncs.sync.ldp_igp_sync_intf_gr_only")
	proto.RegisterType((*LdpIgpSyncInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.standby.default_vrf.afs.af.igp.syncs.sync.ldp_igp_sync_info")
}

func init() { proto.RegisterFile("ldp_igp_sync_info.proto", fileDescriptor_9ebc3d2375525d14) }

var fileDescriptor_9ebc3d2375525d14 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x55, 0x52, 0x5a, 0x60, 0xf2, 0x01, 0x71, 0xa1, 0x5d, 0x6e, 0xd1, 0x0a, 0xa4, 0x88, 0x83,
	0x91, 0x12, 0xbe, 0xaf, 0xa5, 0x42, 0x08, 0x09, 0xd0, 0xa6, 0x48, 0xe5, 0xc2, 0xc8, 0xd9, 0xb5,
	0x17, 0x0b, 0xc7, 0x5e, 0xd9, 0x6e, 0xda, 0xbd, 0xf2, 0x0f, 0xf8, 0x2f, 0xfc, 0x40, 0x64, 0x6f,
	0x93, 0x36, 0x5a, 0x71, 0x2c, 0x97, 0x24, 0x7e, 0x6f, 0xe6, 0xcd, 0xcb, 0x3c, 0xef, 0xc2, 0xa1,
	0x2a, 0x2a, 0x94, 0x65, 0x85, 0xae, 0xd6, 0x39, 0x4a, 0x2d, 0x0c, 0xad, 0xac, 0xf1, 0x86, 0x9c,
	0xe4, 0xd2, 0xe5, 0x06, 0xa5, 0x71, 0x78, 0x61, 0x71, 0x59, 0x29, 0x87, 0xa1, 0xd4, 0x54, 0xdc,
	0xd2, 0xf5, 0x89, 0x96, 0xca, 0x2c, 0x98, 0xa2, 0xce, 0x33, 0x5d, 0x2c, 0x6a, 0x5a, 0x70, 0xc1,
	0xce, 0x94, 0xc7, 0x95, 0x15, 0x94, 0x09, 0x47, 0x99, 0xa0, 0xb2, 0xac, 0x68, 0x90, 0x76, 0xf1,
	0x33, 0x3d, 0x85, 0x83, 0xd6, 0x40, 0xfc, 0x78, 0xfc, 0x6d, 0x4e, 0x0e, 0xe1, 0x36, 0x13, 0xa8,
	0xd9, 0x92, 0x27, 0x9d, 0x71, 0x67, 0x72, 0x37, 0xdb, 0x63, 0xe2, 0x13, 0x5b, 0x72, 0xf2, 0x04,
	0x86, 0x52, 0x7b, 0x6e, 0x05, 0xcb, 0x79, 0xc3, 0x77, 0x23, 0x3f, 0xd8, 0xa0, 0xa1, 0x2c, 0x9d,
	0x42, 0x3f, 0x28, 0xaf, 0xac, 0x88, 0xa2, 0x84, 0xc0, 0xad, 0x6b, 0x62, 0xf1, 0x37, 0x19, 0x42,
	0x57, 0x16, 0xb1, 0x7d, 0x90, 0x75, 0x65, 0x91, 0x7e, 0x6d, 0xb9, 0xf1, 0x02, 0x2b, 0xce, 0x6d,
	0x70, 0x13, 0xbe, 0x51, 0x16, 0x6b, 0x37, 0xe1, 0xf8, 0xa1, 0x20, 0x29, 0x0c, 0xa4, 0xc3, 0xd2,
	0x22, 0xd7, 0x6c, 0xa1, 0x78, 0xa3, 0x76, 0x27, 0xeb, 0x49, 0xf7, 0xde, 0x1e, 0x37, 0x50, 0xfa,
	0x1d, 0x1e, 0xb5, 0x65, 0x4b, 0x8b, 0x46, 0xab, 0xfa, 0xdf, 0xca, 0x13, 0xb8, 0x2f, 0x1d, 0xe6,
	0x3f, 0x7e, 0x56, 0x1e, 0x73, 0xcb, 0x99, 0xdf, 0x88, 0x0f, 0xa5, 0x3b, 0x0a, 0xf0, 0x51, 0x83,
	0xa6, 0x7f, 0x76, 0x61, 0xd4, 0xda, 0x22, 0xf1, 0xb0, 0xb3, 0xb2, 0x22, 0x99, 0x8e, 0x3b, 0x93,
	0xde, 0x74, 0x41, 0x6f, 0x22, 0x3e, 0x7a, 0x7d, 0xc3, 0x59, 0x18, 0x47, 0x9e, 0xc2, 0x68, 0x3b,
	0x1d, 0xbc, 0xb0, 0xc9, 0x2c, 0xfe, 0xb1, 0x7b, 0x5b, 0x01, 0x9d, 0x5a, 0xf2, 0x18, 0x86, 0x1b,
	0xcb, 0xce, 0x33, 0xcf, 0x93, 0xe7, 0xb1, 0xb0, 0x2f, 0xcb, 0x6a, 0x5e, 0xeb, 0x7c, 0x1e, 0xb0,
	0xad, 0xaa, 0x82, 0x2b, 0x56, 0x27, 0x2f, 0x62, 0x60, 0xeb, 0xaa, 0x77, 0x01, 0x23, 0x33, 0x38,
	0x90, 0xae, 0xe1, 0xd1, 0xcb, 0x25, 0xb7, 0x68, 0xcf, 0xb4, 0x96, 0xba, 0x4c, 0x5e, 0xc6, 0x9d,
	0xed, 0x4b, 0x17, 0x0b, 0x4f, 0x02, 0x97, 0x35, 0x14, 0x99, 0xc2, 0xc3, 0xad, 0x0e, 0xbe, 0x64,
	0x32, 0xf6, 0xbc, 0x8a, 0x13, 0xf6, 0x8b, 0xab, 0x8e, 0x35, 0x45, 0x7e, 0x75, 0x60, 0x37, 0x24,
	0xe4, 0x92, 0xd7, 0xe3, 0x9d, 0x49, 0x6f, 0xaa, 0x6e, 0x6e, 0xb3, 0xed, 0x7b, 0x98, 0x35, 0xa3,
	0xc9, 0xef, 0x0e, 0xf4, 0x2f, 0x2f, 0x50, 0xc4, 0x93, 0x37, 0xd1, 0x8b, 0xf9, 0x5f, 0x5e, 0x2e,
	0x67, 0x67, 0x50, 0xda, 0xcf, 0x5a, 0xd5, 0x5f, 0xc2, 0x23, 0xf2, 0x0c, 0x1e, 0x5c, 0xe5, 0x64,
	0xce, 0x35, 0x5a, 0xce, 0x9c, 0xd1, 0xc9, 0xdb, 0x98, 0xe9, 0x68, 0x9d, 0x96, 0x39, 0xd7, 0x59,
	0x24, 0x16, 0x7b, 0xf1, 0xc5, 0x32, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x96, 0xeb, 0x8f, 0xf7,
	0x73, 0x04, 0x00, 0x00,
}