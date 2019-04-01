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
// source: ldp_targeted_hello_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_afs_af_discovery_targeted_hellos_targeted_hello

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

type LdpTargetedHelloInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	LocalAddress         string   `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	TargetAddress        string   `protobuf:"bytes,4,opt,name=target_address,json=targetAddress,proto3" json:"target_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpTargetedHelloInfo_KEYS) Reset()         { *m = LdpTargetedHelloInfo_KEYS{} }
func (m *LdpTargetedHelloInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpTargetedHelloInfo_KEYS) ProtoMessage()    {}
func (*LdpTargetedHelloInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cab00c979ee379, []int{0}
}

func (m *LdpTargetedHelloInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTargetedHelloInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpTargetedHelloInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTargetedHelloInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpTargetedHelloInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTargetedHelloInfo_KEYS.Merge(m, src)
}
func (m *LdpTargetedHelloInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpTargetedHelloInfo_KEYS.Size(m)
}
func (m *LdpTargetedHelloInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTargetedHelloInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTargetedHelloInfo_KEYS proto.InternalMessageInfo

func (m *LdpTargetedHelloInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpTargetedHelloInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpTargetedHelloInfo_KEYS) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *LdpTargetedHelloInfo_KEYS) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
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
	return fileDescriptor_d8cab00c979ee379, []int{1}
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

type LdpLastSessInfo struct {
	LastSessionDownReason string   `protobuf:"bytes,1,opt,name=last_session_down_reason,json=lastSessionDownReason,proto3" json:"last_session_down_reason,omitempty"`
	LastSessionDownTime   uint64   `protobuf:"varint,2,opt,name=last_session_down_time,json=lastSessionDownTime,proto3" json:"last_session_down_time,omitempty"`
	LastSessionUpTime     uint32   `protobuf:"varint,3,opt,name=last_session_up_time,json=lastSessionUpTime,proto3" json:"last_session_up_time,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LdpLastSessInfo) Reset()         { *m = LdpLastSessInfo{} }
func (m *LdpLastSessInfo) String() string { return proto.CompactTextString(m) }
func (*LdpLastSessInfo) ProtoMessage()    {}
func (*LdpLastSessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cab00c979ee379, []int{2}
}

func (m *LdpLastSessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpLastSessInfo.Unmarshal(m, b)
}
func (m *LdpLastSessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpLastSessInfo.Marshal(b, m, deterministic)
}
func (m *LdpLastSessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpLastSessInfo.Merge(m, src)
}
func (m *LdpLastSessInfo) XXX_Size() int {
	return xxx_messageInfo_LdpLastSessInfo.Size(m)
}
func (m *LdpLastSessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpLastSessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpLastSessInfo proto.InternalMessageInfo

func (m *LdpLastSessInfo) GetLastSessionDownReason() string {
	if m != nil {
		return m.LastSessionDownReason
	}
	return ""
}

func (m *LdpLastSessInfo) GetLastSessionDownTime() uint64 {
	if m != nil {
		return m.LastSessionDownTime
	}
	return 0
}

func (m *LdpLastSessInfo) GetLastSessionUpTime() uint32 {
	if m != nil {
		return m.LastSessionUpTime
	}
	return 0
}

type LdpTargetedHelloInfo struct {
	DhcbLocalAddress            *LdpIpAddrTUnion   `protobuf:"bytes,50,opt,name=dhcb_local_address,json=dhcbLocalAddress,proto3" json:"dhcb_local_address,omitempty"`
	DhcbTargetAddress           *LdpIpAddrTUnion   `protobuf:"bytes,51,opt,name=dhcb_target_address,json=dhcbTargetAddress,proto3" json:"dhcb_target_address,omitempty"`
	State                       string             `protobuf:"bytes,52,opt,name=state,proto3" json:"state,omitempty"`
	AdjacencyLdpIdentifier      string             `protobuf:"bytes,53,opt,name=adjacency_ldp_identifier,json=adjacencyLdpIdentifier,proto3" json:"adjacency_ldp_identifier,omitempty"`
	Interval                    uint32             `protobuf:"varint,54,opt,name=interval,proto3" json:"interval,omitempty"`
	NextHello                   uint32             `protobuf:"varint,55,opt,name=next_hello,json=nextHello,proto3" json:"next_hello,omitempty"`
	HoldTime                    uint32             `protobuf:"varint,56,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	LocalHoldTime               uint32             `protobuf:"varint,57,opt,name=local_hold_time,json=localHoldTime,proto3" json:"local_hold_time,omitempty"`
	NeighborHoldTime            uint32             `protobuf:"varint,58,opt,name=neighbor_hold_time,json=neighborHoldTime,proto3" json:"neighbor_hold_time,omitempty"`
	DiscExpiry                  uint32             `protobuf:"varint,59,opt,name=disc_expiry,json=discExpiry,proto3" json:"disc_expiry,omitempty"`
	QuickStartDisabled          bool               `protobuf:"varint,60,opt,name=quick_start_disabled,json=quickStartDisabled,proto3" json:"quick_start_disabled,omitempty"`
	EstablishedTime             uint64             `protobuf:"varint,61,opt,name=established_time,json=establishedTime,proto3" json:"established_time,omitempty"`
	EstablishedAge              uint64             `protobuf:"varint,62,opt,name=established_age,json=establishedAge,proto3" json:"established_age,omitempty"`
	SessionUp                   bool               `protobuf:"varint,63,opt,name=session_up,json=sessionUp,proto3" json:"session_up,omitempty"`
	SessionBringupFailureReason string             `protobuf:"bytes,64,opt,name=session_bringup_failure_reason,json=sessionBringupFailureReason,proto3" json:"session_bringup_failure_reason,omitempty"`
	LastSessionDownInfo         []*LdpLastSessInfo `protobuf:"bytes,65,rep,name=last_session_down_info,json=lastSessionDownInfo,proto3" json:"last_session_down_info,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}           `json:"-"`
	XXX_unrecognized            []byte             `json:"-"`
	XXX_sizecache               int32              `json:"-"`
}

func (m *LdpTargetedHelloInfo) Reset()         { *m = LdpTargetedHelloInfo{} }
func (m *LdpTargetedHelloInfo) String() string { return proto.CompactTextString(m) }
func (*LdpTargetedHelloInfo) ProtoMessage()    {}
func (*LdpTargetedHelloInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8cab00c979ee379, []int{3}
}

func (m *LdpTargetedHelloInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpTargetedHelloInfo.Unmarshal(m, b)
}
func (m *LdpTargetedHelloInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpTargetedHelloInfo.Marshal(b, m, deterministic)
}
func (m *LdpTargetedHelloInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpTargetedHelloInfo.Merge(m, src)
}
func (m *LdpTargetedHelloInfo) XXX_Size() int {
	return xxx_messageInfo_LdpTargetedHelloInfo.Size(m)
}
func (m *LdpTargetedHelloInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpTargetedHelloInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpTargetedHelloInfo proto.InternalMessageInfo

func (m *LdpTargetedHelloInfo) GetDhcbLocalAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.DhcbLocalAddress
	}
	return nil
}

func (m *LdpTargetedHelloInfo) GetDhcbTargetAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.DhcbTargetAddress
	}
	return nil
}

func (m *LdpTargetedHelloInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *LdpTargetedHelloInfo) GetAdjacencyLdpIdentifier() string {
	if m != nil {
		return m.AdjacencyLdpIdentifier
	}
	return ""
}

func (m *LdpTargetedHelloInfo) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetNextHello() uint32 {
	if m != nil {
		return m.NextHello
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetLocalHoldTime() uint32 {
	if m != nil {
		return m.LocalHoldTime
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetNeighborHoldTime() uint32 {
	if m != nil {
		return m.NeighborHoldTime
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetDiscExpiry() uint32 {
	if m != nil {
		return m.DiscExpiry
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetQuickStartDisabled() bool {
	if m != nil {
		return m.QuickStartDisabled
	}
	return false
}

func (m *LdpTargetedHelloInfo) GetEstablishedTime() uint64 {
	if m != nil {
		return m.EstablishedTime
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetEstablishedAge() uint64 {
	if m != nil {
		return m.EstablishedAge
	}
	return 0
}

func (m *LdpTargetedHelloInfo) GetSessionUp() bool {
	if m != nil {
		return m.SessionUp
	}
	return false
}

func (m *LdpTargetedHelloInfo) GetSessionBringupFailureReason() string {
	if m != nil {
		return m.SessionBringupFailureReason
	}
	return ""
}

func (m *LdpTargetedHelloInfo) GetLastSessionDownInfo() []*LdpLastSessInfo {
	if m != nil {
		return m.LastSessionDownInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LdpTargetedHelloInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.targeted_hellos.targeted_hello.ldp_targeted_hello_info_KEYS")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.targeted_hellos.targeted_hello.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpLastSessInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.targeted_hellos.targeted_hello.ldp_last_sess_info")
	proto.RegisterType((*LdpTargetedHelloInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.targeted_hellos.targeted_hello.ldp_targeted_hello_info")
}

func init() { proto.RegisterFile("ldp_targeted_hello_info.proto", fileDescriptor_d8cab00c979ee379) }

var fileDescriptor_d8cab00c979ee379 = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0x5f, 0x53, 0x13, 0x3f,
	0x14, 0x9d, 0xfd, 0x95, 0x1f, 0xb6, 0x41, 0xa0, 0x04, 0x84, 0x1d, 0x11, 0xed, 0xd4, 0x51, 0xeb,
	0x8c, 0xb3, 0x3a, 0x80, 0x80, 0xff, 0x45, 0xc1, 0x81, 0x91, 0xf1, 0xa1, 0xe0, 0x83, 0x4f, 0x99,
	0xb4, 0xc9, 0xb6, 0xc1, 0x6d, 0xb2, 0x26, 0x69, 0xa1, 0x2f, 0xfa, 0x55, 0xfc, 0x02, 0x8e, 0x0f,
	0xbc, 0xf9, 0xe9, 0x9c, 0xdc, 0x6c, 0xcb, 0x96, 0x3f, 0xaf, 0x8c, 0x2f, 0x9d, 0xcd, 0x39, 0xe7,
	0xf6, 0x9e, 0xe6, 0x9e, 0xbb, 0x45, 0x4b, 0x09, 0x4b, 0x89, 0xa5, 0xba, 0xc5, 0x2d, 0x67, 0xa4,
	0xcd, 0x93, 0x44, 0x11, 0x21, 0x63, 0x15, 0xa5, 0x5a, 0x59, 0x85, 0x0f, 0x9b, 0xc2, 0x34, 0x15,
	0x11, 0xca, 0x90, 0x63, 0x4d, 0x3a, 0x69, 0x62, 0x88, 0x2b, 0x50, 0x29, 0xd7, 0xd1, 0xe0, 0x14,
	0x49, 0xc5, 0xb8, 0x81, 0xcf, 0x88, 0xf1, 0x98, 0x76, 0x13, 0x4b, 0x7a, 0x3a, 0x8e, 0x68, 0x6c,
	0x22, 0x1a, 0x47, 0xcc, 0x7d, 0x43, 0x8f, 0xeb, 0x7e, 0x34, 0xda, 0xc6, 0x9c, 0x39, 0x57, 0x7f,
	0x06, 0xe8, 0xd6, 0x25, 0x6e, 0xc8, 0xc7, 0xed, 0x2f, 0xfb, 0x78, 0x11, 0x95, 0x5c, 0x13, 0x22,
	0x69, 0x87, 0x87, 0x41, 0x25, 0xa8, 0x95, 0xea, 0x45, 0x07, 0x7c, 0xa2, 0x1d, 0x8e, 0x17, 0xd0,
	0x35, 0x1a, 0x7b, 0xea, 0x3f, 0xa0, 0xc6, 0x69, 0x0c, 0xc4, 0x5d, 0x34, 0x99, 0xa8, 0x26, 0x4d,
	0x08, 0x65, 0x4c, 0x73, 0x63, 0xc2, 0x02, 0xd0, 0xd7, 0x01, 0xdc, 0xf4, 0x18, 0xbe, 0x87, 0xa6,
	0x7c, 0xdb, 0xa1, 0x6a, 0x0c, 0x54, 0x93, 0x1e, 0xcd, 0x64, 0x55, 0x8e, 0x66, 0x9d, 0x43, 0x91,
	0x82, 0x8c, 0x58, 0xd2, 0x95, 0x42, 0x49, 0x5c, 0x46, 0x05, 0x1a, 0x8b, 0xcc, 0x92, 0x7b, 0xc4,
	0x73, 0xe8, 0x7f, 0xd6, 0xed, 0x74, 0xfa, 0xe0, 0x65, 0xb2, 0xee, 0x0f, 0x18, 0xa3, 0x31, 0x91,
	0xf6, 0x56, 0x33, 0x07, 0xf0, 0x9c, 0x61, 0x6b, 0x59, 0x3f, 0x78, 0xae, 0x9e, 0x04, 0x08, 0xbb,
	0x3e, 0x09, 0x35, 0x96, 0x18, 0x6e, 0x0c, 0x5c, 0x02, 0x5e, 0x47, 0xe1, 0x10, 0x11, 0x4a, 0x12,
	0xa6, 0x8e, 0x24, 0xd1, 0x9c, 0x1a, 0x25, 0xb3, 0xde, 0x37, 0x1c, 0xbf, 0xef, 0xe9, 0x2d, 0x75,
	0x24, 0xeb, 0x40, 0xe2, 0x15, 0x34, 0x7f, 0xbe, 0xd0, 0x8a, 0xec, 0xaa, 0xc6, 0xea, 0xb3, 0x67,
	0xca, 0x0e, 0x44, 0x87, 0xe3, 0xc7, 0x68, 0x6e, 0xa4, 0xa8, 0x9b, 0xfa, 0x92, 0x02, 0xfc, 0xa2,
	0x99, 0x5c, 0xc9, 0xe7, 0xd4, 0x15, 0x54, 0xff, 0x14, 0xd1, 0xc2, 0x25, 0xf3, 0xc3, 0xbf, 0x02,
	0x84, 0x59, 0xbb, 0xd9, 0x20, 0xa3, 0xa3, 0x58, 0xae, 0x04, 0xb5, 0x89, 0xe5, 0x1f, 0xd1, 0xd5,
	0xa5, 0x2c, 0xba, 0x60, 0x7e, 0xf5, 0xb2, 0xb3, 0xb6, 0x97, 0xcf, 0xc3, 0xef, 0x00, 0xcd, 0x82,
	0xdf, 0x33, 0xa9, 0x58, 0xf9, 0x37, 0x0c, 0xcf, 0x38, 0x6f, 0x07, 0xf9, 0x68, 0xba, 0xc4, 0x19,
	0x4b, 0x2d, 0x0f, 0x57, 0x21, 0x09, 0xfe, 0x80, 0x37, 0x50, 0x48, 0xd9, 0x21, 0x6d, 0x72, 0xd9,
	0xec, 0x83, 0x45, 0xc1, 0xb8, 0xb4, 0x22, 0x16, 0x5c, 0x87, 0x4f, 0x41, 0x38, 0x3f, 0xe4, 0xf7,
	0x58, 0xba, 0x3b, 0x64, 0xf1, 0x4d, 0x54, 0x14, 0xd2, 0x72, 0xdd, 0xa3, 0x49, 0xb8, 0x06, 0x23,
	0x1f, 0x9e, 0xf1, 0x12, 0x42, 0x92, 0x1f, 0x5b, 0x6f, 0x30, 0x5c, 0x07, 0xb6, 0xe4, 0x90, 0x1d,
	0x07, 0xb8, 0x3d, 0x6d, 0xab, 0x84, 0xf9, 0xb8, 0x6c, 0xf8, 0x5a, 0x07, 0x40, 0xac, 0xee, 0xa3,
	0x69, 0x9f, 0x81, 0x53, 0xc9, 0x33, 0x90, 0xf8, 0x2d, 0xdd, 0x19, 0xe8, 0x1e, 0x21, 0x2c, 0xb9,
	0x68, 0xb5, 0x1b, 0x4a, 0xe7, 0xa4, 0xcf, 0x41, 0x5a, 0x1e, 0x30, 0x43, 0xf5, 0x1d, 0x34, 0xe1,
	0x2e, 0x94, 0xf0, 0xe3, 0x54, 0xe8, 0x7e, 0xf8, 0x02, 0x64, 0xc8, 0x41, 0xdb, 0x80, 0xe0, 0x27,
	0x68, 0xee, 0x5b, 0x57, 0x34, 0xbf, 0x12, 0x63, 0xa9, 0xb6, 0x84, 0x09, 0x43, 0x1b, 0x09, 0x67,
	0xe1, 0xcb, 0x4a, 0x50, 0x2b, 0xd6, 0x31, 0x70, 0xfb, 0x8e, 0xda, 0xca, 0x18, 0xfc, 0x10, 0x95,
	0xb9, 0xb1, 0xb4, 0x91, 0x08, 0xd3, 0xe6, 0x59, 0xfb, 0x57, 0xb0, 0x2e, 0xd3, 0x39, 0x1c, 0xba,
	0x3f, 0x40, 0x79, 0x88, 0xd0, 0x16, 0x0f, 0x5f, 0x83, 0x72, 0x2a, 0x07, 0x6f, 0xb6, 0xb8, 0xbb,
	0xb8, 0xd3, 0x75, 0x0a, 0xdf, 0x40, 0xef, 0x92, 0x19, 0x6c, 0x11, 0x7e, 0x8f, 0x6e, 0x0f, 0xe8,
	0x86, 0x16, 0xb2, 0xd5, 0x4d, 0x49, 0x4c, 0x45, 0xd2, 0xd5, 0x7c, 0xb0, 0xe6, 0x6f, 0x61, 0x66,
	0x8b, 0x99, 0xea, 0x9d, 0x17, 0x7d, 0xf0, 0x9a, 0x6c, 0xd9, 0x4f, 0x82, 0x8b, 0xb6, 0xdd, 0x6d,
	0x61, 0xb8, 0x59, 0x29, 0xd4, 0x26, 0x96, 0xbf, 0x5f, 0x75, 0x7a, 0x47, 0x5f, 0x63, 0xe7, 0xde,
	0x36, 0xbb, 0x32, 0x56, 0x8d, 0x71, 0xf8, 0xbf, 0x59, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0x06, 0xca, 0x9e, 0x90, 0x06, 0x00, 0x00,
}