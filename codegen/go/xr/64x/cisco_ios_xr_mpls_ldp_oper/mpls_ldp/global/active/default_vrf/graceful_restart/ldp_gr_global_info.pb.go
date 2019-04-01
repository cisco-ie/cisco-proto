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
// source: ldp_gr_global_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_global_active_default_vrf_graceful_restart

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

type LdpGrGlobalInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpGrGlobalInfo_KEYS) Reset()         { *m = LdpGrGlobalInfo_KEYS{} }
func (m *LdpGrGlobalInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpGrGlobalInfo_KEYS) ProtoMessage()    {}
func (*LdpGrGlobalInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f02f7481d141c0, []int{0}
}

func (m *LdpGrGlobalInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpGrGlobalInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpGrGlobalInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpGrGlobalInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpGrGlobalInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpGrGlobalInfo_KEYS.Merge(m, src)
}
func (m *LdpGrGlobalInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpGrGlobalInfo_KEYS.Size(m)
}
func (m *LdpGrGlobalInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpGrGlobalInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpGrGlobalInfo_KEYS proto.InternalMessageInfo

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
	return fileDescriptor_16f02f7481d141c0, []int{1}
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
	return fileDescriptor_16f02f7481d141c0, []int{2}
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

type LdpGrDnbrIntfInfo struct {
	AddressFamily        string   `protobuf:"bytes,1,opt,name=address_family,json=addressFamily,proto3" json:"address_family,omitempty"`
	InterfaceHandle      string   `protobuf:"bytes,2,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpGrDnbrIntfInfo) Reset()         { *m = LdpGrDnbrIntfInfo{} }
func (m *LdpGrDnbrIntfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpGrDnbrIntfInfo) ProtoMessage()    {}
func (*LdpGrDnbrIntfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f02f7481d141c0, []int{3}
}

func (m *LdpGrDnbrIntfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpGrDnbrIntfInfo.Unmarshal(m, b)
}
func (m *LdpGrDnbrIntfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpGrDnbrIntfInfo.Marshal(b, m, deterministic)
}
func (m *LdpGrDnbrIntfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpGrDnbrIntfInfo.Merge(m, src)
}
func (m *LdpGrDnbrIntfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpGrDnbrIntfInfo.Size(m)
}
func (m *LdpGrDnbrIntfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpGrDnbrIntfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpGrDnbrIntfInfo proto.InternalMessageInfo

func (m *LdpGrDnbrIntfInfo) GetAddressFamily() string {
	if m != nil {
		return m.AddressFamily
	}
	return ""
}

func (m *LdpGrDnbrIntfInfo) GetInterfaceHandle() string {
	if m != nil {
		return m.InterfaceHandle
	}
	return ""
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
	return fileDescriptor_16f02f7481d141c0, []int{4}
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

type LdpGrNbrInfo struct {
	GrPeer                        *LdpLdpidInfo        `protobuf:"bytes,1,opt,name=gr_peer,json=grPeer,proto3" json:"gr_peer,omitempty"`
	ConnectCount                  uint32               `protobuf:"varint,2,opt,name=connect_count,json=connectCount,proto3" json:"connect_count,omitempty"`
	IsNeighborUp                  bool                 `protobuf:"varint,3,opt,name=is_neighbor_up,json=isNeighborUp,proto3" json:"is_neighbor_up,omitempty"`
	IsLivenessTimerRunning        bool                 `protobuf:"varint,4,opt,name=is_liveness_timer_running,json=isLivenessTimerRunning,proto3" json:"is_liveness_timer_running,omitempty"`
	LivenessTimerRemainingSeconds uint32               `protobuf:"varint,5,opt,name=liveness_timer_remaining_seconds,json=livenessTimerRemainingSeconds,proto3" json:"liveness_timer_remaining_seconds,omitempty"`
	IsRecoveryTimerRunning        bool                 `protobuf:"varint,6,opt,name=is_recovery_timer_running,json=isRecoveryTimerRunning,proto3" json:"is_recovery_timer_running,omitempty"`
	RecoveryTimerRemainingSeconds uint32               `protobuf:"varint,7,opt,name=recovery_timer_remaining_seconds,json=recoveryTimerRemainingSeconds,proto3" json:"recovery_timer_remaining_seconds,omitempty"`
	DownNbrInterface              []*LdpGrDnbrIntfInfo `protobuf:"bytes,8,rep,name=down_nbr_interface,json=downNbrInterface,proto3" json:"down_nbr_interface,omitempty"`
	DownNbrAddress                []*LdpIpAddrTUnion   `protobuf:"bytes,9,rep,name=down_nbr_address,json=downNbrAddress,proto3" json:"down_nbr_address,omitempty"`
	DownNbrFlapCount              uint32               `protobuf:"varint,10,opt,name=down_nbr_flap_count,json=downNbrFlapCount,proto3" json:"down_nbr_flap_count,omitempty"`
	DownNbrFlags                  uint32               `protobuf:"varint,11,opt,name=down_nbr_flags,json=downNbrFlags,proto3" json:"down_nbr_flags,omitempty"`
	DownNbrDownReason             uint32               `protobuf:"varint,12,opt,name=down_nbr_down_reason,json=downNbrDownReason,proto3" json:"down_nbr_down_reason,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *LdpGrNbrInfo) Reset()         { *m = LdpGrNbrInfo{} }
func (m *LdpGrNbrInfo) String() string { return proto.CompactTextString(m) }
func (*LdpGrNbrInfo) ProtoMessage()    {}
func (*LdpGrNbrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f02f7481d141c0, []int{5}
}

func (m *LdpGrNbrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpGrNbrInfo.Unmarshal(m, b)
}
func (m *LdpGrNbrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpGrNbrInfo.Marshal(b, m, deterministic)
}
func (m *LdpGrNbrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpGrNbrInfo.Merge(m, src)
}
func (m *LdpGrNbrInfo) XXX_Size() int {
	return xxx_messageInfo_LdpGrNbrInfo.Size(m)
}
func (m *LdpGrNbrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpGrNbrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpGrNbrInfo proto.InternalMessageInfo

func (m *LdpGrNbrInfo) GetGrPeer() *LdpLdpidInfo {
	if m != nil {
		return m.GrPeer
	}
	return nil
}

func (m *LdpGrNbrInfo) GetConnectCount() uint32 {
	if m != nil {
		return m.ConnectCount
	}
	return 0
}

func (m *LdpGrNbrInfo) GetIsNeighborUp() bool {
	if m != nil {
		return m.IsNeighborUp
	}
	return false
}

func (m *LdpGrNbrInfo) GetIsLivenessTimerRunning() bool {
	if m != nil {
		return m.IsLivenessTimerRunning
	}
	return false
}

func (m *LdpGrNbrInfo) GetLivenessTimerRemainingSeconds() uint32 {
	if m != nil {
		return m.LivenessTimerRemainingSeconds
	}
	return 0
}

func (m *LdpGrNbrInfo) GetIsRecoveryTimerRunning() bool {
	if m != nil {
		return m.IsRecoveryTimerRunning
	}
	return false
}

func (m *LdpGrNbrInfo) GetRecoveryTimerRemainingSeconds() uint32 {
	if m != nil {
		return m.RecoveryTimerRemainingSeconds
	}
	return 0
}

func (m *LdpGrNbrInfo) GetDownNbrInterface() []*LdpGrDnbrIntfInfo {
	if m != nil {
		return m.DownNbrInterface
	}
	return nil
}

func (m *LdpGrNbrInfo) GetDownNbrAddress() []*LdpIpAddrTUnion {
	if m != nil {
		return m.DownNbrAddress
	}
	return nil
}

func (m *LdpGrNbrInfo) GetDownNbrFlapCount() uint32 {
	if m != nil {
		return m.DownNbrFlapCount
	}
	return 0
}

func (m *LdpGrNbrInfo) GetDownNbrFlags() uint32 {
	if m != nil {
		return m.DownNbrFlags
	}
	return 0
}

func (m *LdpGrNbrInfo) GetDownNbrDownReason() uint32 {
	if m != nil {
		return m.DownNbrDownReason
	}
	return 0
}

type LdpGrGlobalInfo struct {
	Vrf                                      *LdpVrfInfo     `protobuf:"bytes,50,opt,name=vrf,proto3" json:"vrf,omitempty"`
	IsForwardingStateHoldTimerRunning        bool            `protobuf:"varint,51,opt,name=is_forwarding_state_hold_timer_running,json=isForwardingStateHoldTimerRunning,proto3" json:"is_forwarding_state_hold_timer_running,omitempty"`
	ForwardingStateHoldTimerRemainingSeconds uint32          `protobuf:"varint,52,opt,name=forwarding_state_hold_timer_remaining_seconds,json=forwardingStateHoldTimerRemainingSeconds,proto3" json:"forwarding_state_hold_timer_remaining_seconds,omitempty"`
	GracefulRestartableNeighbor              []*LdpGrNbrInfo `protobuf:"bytes,53,rep,name=graceful_restartable_neighbor,json=gracefulRestartableNeighbor,proto3" json:"graceful_restartable_neighbor,omitempty"`
	XXX_NoUnkeyedLiteral                     struct{}        `json:"-"`
	XXX_unrecognized                         []byte          `json:"-"`
	XXX_sizecache                            int32           `json:"-"`
}

func (m *LdpGrGlobalInfo) Reset()         { *m = LdpGrGlobalInfo{} }
func (m *LdpGrGlobalInfo) String() string { return proto.CompactTextString(m) }
func (*LdpGrGlobalInfo) ProtoMessage()    {}
func (*LdpGrGlobalInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_16f02f7481d141c0, []int{6}
}

func (m *LdpGrGlobalInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpGrGlobalInfo.Unmarshal(m, b)
}
func (m *LdpGrGlobalInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpGrGlobalInfo.Marshal(b, m, deterministic)
}
func (m *LdpGrGlobalInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpGrGlobalInfo.Merge(m, src)
}
func (m *LdpGrGlobalInfo) XXX_Size() int {
	return xxx_messageInfo_LdpGrGlobalInfo.Size(m)
}
func (m *LdpGrGlobalInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpGrGlobalInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpGrGlobalInfo proto.InternalMessageInfo

func (m *LdpGrGlobalInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpGrGlobalInfo) GetIsForwardingStateHoldTimerRunning() bool {
	if m != nil {
		return m.IsForwardingStateHoldTimerRunning
	}
	return false
}

func (m *LdpGrGlobalInfo) GetForwardingStateHoldTimerRemainingSeconds() uint32 {
	if m != nil {
		return m.ForwardingStateHoldTimerRemainingSeconds
	}
	return 0
}

func (m *LdpGrGlobalInfo) GetGracefulRestartableNeighbor() []*LdpGrNbrInfo {
	if m != nil {
		return m.GracefulRestartableNeighbor
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpGrGlobalInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_gr_global_info_KEYS")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_vrf_info")
	proto.RegisterType((*LdpLdpidInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_ldpid_info")
	proto.RegisterType((*LdpGrDnbrIntfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_gr_dnbr_intf_info")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpGrNbrInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_gr_nbr_info")
	proto.RegisterType((*LdpGrGlobalInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.global.active.default_vrf.graceful_restart.ldp_gr_global_info")
}

func init() { proto.RegisterFile("ldp_gr_global_info.proto", fileDescriptor_16f02f7481d141c0) }

var fileDescriptor_16f02f7481d141c0 = []byte{
	// 742 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0x56, 0x9a, 0x36, 0x6d, 0xdd, 0x34, 0xcd, 0x75, 0xdb, 0x7b, 0xa7, 0xba, 0xaa, 0x14, 0xc2,
	0x8f, 0xc2, 0xa2, 0x83, 0x94, 0x96, 0x4a, 0x2c, 0x11, 0x50, 0x1a, 0x81, 0x2a, 0x98, 0xc0, 0x02,
	0x09, 0x61, 0x39, 0x63, 0xcf, 0xd4, 0x92, 0x63, 0x5b, 0xf6, 0x64, 0x42, 0x1f, 0x81, 0x05, 0x2b,
	0x1e, 0x81, 0xf7, 0xe2, 0x59, 0x90, 0x3d, 0x9e, 0x29, 0x99, 0x08, 0x56, 0x64, 0xe7, 0xf9, 0xce,
	0x77, 0x7c, 0x3e, 0xfb, 0x7c, 0x3e, 0x03, 0x02, 0x4e, 0x14, 0x4a, 0x35, 0x4a, 0xb9, 0x9c, 0x60,
	0x8e, 0x98, 0x48, 0x64, 0xa8, 0xb4, 0xcc, 0x24, 0xbc, 0x8a, 0x99, 0x89, 0x25, 0x62, 0xd2, 0xa0,
	0xcf, 0x1a, 0x4d, 0x15, 0x37, 0xc8, 0x72, 0xa5, 0xa2, 0x3a, 0x2c, 0xbf, 0xc2, 0x22, 0x2d, 0xc4,
	0x71, 0xc6, 0x72, 0x1a, 0x12, 0x9a, 0xe0, 0x19, 0xcf, 0x50, 0xae, 0x93, 0x30, 0xd5, 0x38, 0xa6,
	0xc9, 0x8c, 0x23, 0x4d, 0x4d, 0x86, 0x75, 0xd6, 0x3f, 0x02, 0xff, 0x2d, 0xd7, 0x42, 0xaf, 0x5e,
	0x7c, 0x18, 0xf7, 0x87, 0xa0, 0x6d, 0x43, 0xb9, 0x4e, 0x1c, 0x08, 0x21, 0x58, 0x17, 0x78, 0x4a,
	0x83, 0x46, 0xaf, 0x31, 0xd8, 0x8e, 0xdc, 0x1a, 0x76, 0xc0, 0x1a, 0x23, 0xc1, 0x5a, 0xaf, 0x31,
	0xd8, 0x8d, 0xd6, 0x18, 0xe9, 0x13, 0xd0, 0xb1, 0x39, 0x9c, 0x28, 0x46, 0x8a, 0xac, 0x43, 0xd0,
	0xe2, 0x46, 0x23, 0x46, 0x7c, 0xde, 0x06, 0x37, 0x7a, 0x44, 0xe0, 0x3d, 0xd0, 0xe1, 0x78, 0x42,
	0x39, 0x32, 0x0a, 0xc7, 0x14, 0x55, 0x9b, 0xb4, 0x1d, 0x3a, 0xb6, 0xe0, 0x88, 0xb8, 0x64, 0xa2,
	0x6c, 0xb4, 0xe9, 0x93, 0x89, 0x1a, 0x91, 0x3e, 0x03, 0x87, 0x5e, 0x34, 0x11, 0x13, 0x8d, 0x98,
	0xc8, 0xbc, 0xc4, 0xfb, 0xa0, 0x83, 0x09, 0xd1, 0xd4, 0x18, 0x94, 0xe0, 0x29, 0xe3, 0x37, 0xbe,
	0xe8, 0xae, 0x47, 0x2f, 0x1c, 0x08, 0x1f, 0x82, 0x2e, 0x13, 0x19, 0xd5, 0x89, 0x2d, 0x7d, 0x8d,
	0x05, 0xe1, 0xd4, 0x95, 0xdf, 0x8e, 0xf6, 0x2a, 0xfc, 0xd2, 0xc1, 0x7d, 0x0a, 0xf6, 0x9d, 0x02,
	0x85, 0xec, 0x16, 0x28, 0x43, 0x33, 0xc1, 0xa4, 0x80, 0x5d, 0xd0, 0xc4, 0x09, 0xf3, 0xbb, 0xdb,
	0x25, 0x3c, 0x00, 0x1b, 0x64, 0x36, 0x9d, 0xde, 0xf8, 0x73, 0x14, 0x1f, 0xf6, 0xce, 0x98, 0xca,
	0xcf, 0xbc, 0x7c, 0xb7, 0xf6, 0xd8, 0x79, 0xb0, 0x5e, 0x61, 0xe7, 0xfd, 0x2f, 0x9b, 0x60, 0xcf,
	0x1f, 0xa9, 0x38, 0x51, 0x22, 0xe1, 0x1c, 0x6c, 0xa6, 0x1a, 0x29, 0x4a, 0xb5, 0xab, 0xb3, 0x33,
	0xfc, 0x14, 0xfe, 0xdd, 0xe6, 0x87, 0x8b, 0xad, 0x8a, 0x5a, 0xa9, 0x7e, 0x43, 0xa9, 0x86, 0x77,
	0xc1, 0x6e, 0x2c, 0x85, 0xa0, 0x71, 0x86, 0x62, 0x39, 0x13, 0x59, 0xd9, 0x1a, 0x0f, 0x3e, 0xb3,
	0x98, 0x6d, 0x20, 0x33, 0x48, 0x50, 0x96, 0x5e, 0x4f, 0xa4, 0x46, 0x33, 0xe5, 0xce, 0xb8, 0x15,
	0xb5, 0x99, 0xb9, 0xf2, 0xe0, 0x7b, 0x05, 0x9f, 0x80, 0x23, 0x66, 0x10, 0x67, 0x39, 0x15, 0xb6,
	0x29, 0x19, 0x9b, 0x52, 0x8d, 0xf4, 0x4c, 0x08, 0x26, 0x52, 0x77, 0x01, 0x5b, 0xd1, 0xbf, 0xcc,
	0xbc, 0xf6, 0xf1, 0x77, 0x36, 0x1c, 0x15, 0x51, 0xf8, 0x12, 0xf4, 0xea, 0x79, 0x74, 0x8a, 0x99,
	0x8d, 0x21, 0x43, 0x63, 0x29, 0x88, 0x09, 0x36, 0x9c, 0xb0, 0x63, 0xbe, 0x90, 0x5f, 0xb2, 0xc6,
	0x05, 0xc9, 0x6b, 0xd0, 0x34, 0x96, 0x39, 0xd5, 0x37, 0x35, 0x0d, 0xad, 0x52, 0x43, 0xe4, 0xe3,
	0x75, 0x0d, 0xf5, 0xbc, 0x25, 0x0d, 0x9b, 0x85, 0x06, 0xbd, 0x90, 0x5f, 0xd7, 0xf0, 0xad, 0x01,
	0x20, 0x91, 0x73, 0xe1, 0xbb, 0xeb, 0x3d, 0x16, 0x6c, 0xf5, 0x9a, 0x83, 0x9d, 0x21, 0x5d, 0x45,
	0x5f, 0x97, 0x1e, 0x47, 0xd4, 0xb5, 0x02, 0xae, 0x26, 0x7a, 0x54, 0x96, 0x87, 0x5f, 0x1b, 0xa0,
	0x5b, 0xa9, 0xf2, 0x4f, 0x24, 0xd8, 0x76, 0x9a, 0xe2, 0x55, 0x68, 0xaa, 0xbd, 0xa2, 0xa8, 0xe3,
	0x15, 0x3d, 0x2d, 0x4a, 0xc3, 0x13, 0xb0, 0x5f, 0xc9, 0x49, 0x38, 0x56, 0xde, 0x7e, 0xc0, 0xdd,
	0x70, 0x29, 0xff, 0x82, 0x63, 0x55, 0x59, 0xf0, 0x57, 0x7a, 0x6a, 0x82, 0x9d, 0xc2, 0xa8, 0xb7,
	0xcc, 0xd4, 0xc0, 0x47, 0xe0, 0xa0, 0x62, 0xb9, 0x85, 0xa6, 0xd8, 0x48, 0x11, 0xb4, 0x1d, 0xf7,
	0x1f, 0xcf, 0x7d, 0x2e, 0xe7, 0x22, 0x72, 0x81, 0xfe, 0x8f, 0x26, 0x80, 0xcb, 0x33, 0x11, 0x0a,
	0xd0, 0xcc, 0x75, 0x12, 0x0c, 0xdd, 0x53, 0xfc, 0xb8, 0x8a, 0xeb, 0x29, 0x27, 0x6d, 0x64, 0x0b,
	0xc1, 0xb7, 0xe0, 0x01, 0x33, 0x28, 0x91, 0x7a, 0x8e, 0x35, 0x71, 0x86, 0xcb, 0x70, 0x46, 0xd1,
	0xb5, 0xe4, 0xa4, 0xe6, 0xe1, 0x53, 0xe7, 0xe1, 0x3b, 0xcc, 0x5c, 0x54, 0xe4, 0xb1, 0xe5, 0x5e,
	0x4a, 0x4e, 0x16, 0xec, 0x8c, 0xc0, 0xc9, 0x1f, 0xf7, 0x5b, 0xf2, 0xf6, 0x99, 0xbb, 0xa3, 0x41,
	0xf2, 0xbb, 0x7d, 0xeb, 0x36, 0xff, 0xde, 0x00, 0xc7, 0xf5, 0xa3, 0xe1, 0x09, 0xa7, 0xd5, 0x9c,
	0x08, 0x1e, 0x3b, 0x77, 0xa1, 0x15, 0x39, 0xbe, 0x9c, 0x9d, 0xd1, 0xff, 0x25, 0x23, 0xba, 0x15,
	0x51, 0x8e, 0xa5, 0x49, 0xcb, 0xfd, 0x4a, 0x4f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xae, 0xbb,
	0xd5, 0x4e, 0x66, 0x07, 0x00, 0x00,
}
