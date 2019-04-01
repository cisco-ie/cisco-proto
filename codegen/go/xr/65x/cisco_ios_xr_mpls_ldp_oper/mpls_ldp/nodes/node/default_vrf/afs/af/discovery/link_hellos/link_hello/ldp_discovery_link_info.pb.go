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
// source: ldp_discovery_link_info.proto

package cisco_ios_xr_mpls_ldp_oper_mpls_ldp_nodes_node_default_vrf_afs_af_discovery_link_hellos_link_hello

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

type LdpDiscoveryLinkInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	AfName               string   `protobuf:"bytes,2,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LdpDiscoveryLinkInfo_KEYS) Reset()         { *m = LdpDiscoveryLinkInfo_KEYS{} }
func (m *LdpDiscoveryLinkInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryLinkInfo_KEYS) ProtoMessage()    {}
func (*LdpDiscoveryLinkInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_332f2a09e882d2e8, []int{0}
}

func (m *LdpDiscoveryLinkInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS.Unmarshal(m, b)
}
func (m *LdpDiscoveryLinkInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryLinkInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS.Merge(m, src)
}
func (m *LdpDiscoveryLinkInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS.Size(m)
}
func (m *LdpDiscoveryLinkInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryLinkInfo_KEYS proto.InternalMessageInfo

func (m *LdpDiscoveryLinkInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *LdpDiscoveryLinkInfo_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *LdpDiscoveryLinkInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_332f2a09e882d2e8, []int{1}
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
	return fileDescriptor_332f2a09e882d2e8, []int{2}
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

type LdpHelloInfo struct {
	NeighborSrcAddress          *LdpIpAddrTUnion   `protobuf:"bytes,1,opt,name=neighbor_src_address,json=neighborSrcAddress,proto3" json:"neighbor_src_address,omitempty"`
	NeighborTransportAddress    *LdpIpAddrTUnion   `protobuf:"bytes,2,opt,name=neighbor_transport_address,json=neighborTransportAddress,proto3" json:"neighbor_transport_address,omitempty"`
	NeighborLdpIdentifier       string             `protobuf:"bytes,3,opt,name=neighbor_ldp_identifier,json=neighborLdpIdentifier,proto3" json:"neighbor_ldp_identifier,omitempty"`
	IsNoRoute                   bool               `protobuf:"varint,4,opt,name=is_no_route,json=isNoRoute,proto3" json:"is_no_route,omitempty"`
	HoldTime                    uint32             `protobuf:"varint,5,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	LocalHoldTime               uint32             `protobuf:"varint,6,opt,name=local_hold_time,json=localHoldTime,proto3" json:"local_hold_time,omitempty"`
	NeighborHoldTime            uint32             `protobuf:"varint,7,opt,name=neighbor_hold_time,json=neighborHoldTime,proto3" json:"neighbor_hold_time,omitempty"`
	DiscExpiry                  uint32             `protobuf:"varint,8,opt,name=disc_expiry,json=discExpiry,proto3" json:"disc_expiry,omitempty"`
	IsTargeted                  bool               `protobuf:"varint,9,opt,name=is_targeted,json=isTargeted,proto3" json:"is_targeted,omitempty"`
	Target                      *LdpIpAddrTUnion   `protobuf:"bytes,10,opt,name=target,proto3" json:"target,omitempty"`
	SessionUp                   bool               `protobuf:"varint,11,opt,name=session_up,json=sessionUp,proto3" json:"session_up,omitempty"`
	EstablishedTime             uint64             `protobuf:"varint,12,opt,name=established_time,json=establishedTime,proto3" json:"established_time,omitempty"`
	EstablishedAge              uint64             `protobuf:"varint,13,opt,name=established_age,json=establishedAge,proto3" json:"established_age,omitempty"`
	SessionBringupFailureReason string             `protobuf:"bytes,14,opt,name=session_bringup_failure_reason,json=sessionBringupFailureReason,proto3" json:"session_bringup_failure_reason,omitempty"`
	LastSessionDownInfo         []*LdpLastSessInfo `protobuf:"bytes,15,rep,name=last_session_down_info,json=lastSessionDownInfo,proto3" json:"last_session_down_info,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}           `json:"-"`
	XXX_unrecognized            []byte             `json:"-"`
	XXX_sizecache               int32              `json:"-"`
}

func (m *LdpHelloInfo) Reset()         { *m = LdpHelloInfo{} }
func (m *LdpHelloInfo) String() string { return proto.CompactTextString(m) }
func (*LdpHelloInfo) ProtoMessage()    {}
func (*LdpHelloInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_332f2a09e882d2e8, []int{3}
}

func (m *LdpHelloInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpHelloInfo.Unmarshal(m, b)
}
func (m *LdpHelloInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpHelloInfo.Marshal(b, m, deterministic)
}
func (m *LdpHelloInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpHelloInfo.Merge(m, src)
}
func (m *LdpHelloInfo) XXX_Size() int {
	return xxx_messageInfo_LdpHelloInfo.Size(m)
}
func (m *LdpHelloInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpHelloInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpHelloInfo proto.InternalMessageInfo

func (m *LdpHelloInfo) GetNeighborSrcAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.NeighborSrcAddress
	}
	return nil
}

func (m *LdpHelloInfo) GetNeighborTransportAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.NeighborTransportAddress
	}
	return nil
}

func (m *LdpHelloInfo) GetNeighborLdpIdentifier() string {
	if m != nil {
		return m.NeighborLdpIdentifier
	}
	return ""
}

func (m *LdpHelloInfo) GetIsNoRoute() bool {
	if m != nil {
		return m.IsNoRoute
	}
	return false
}

func (m *LdpHelloInfo) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *LdpHelloInfo) GetLocalHoldTime() uint32 {
	if m != nil {
		return m.LocalHoldTime
	}
	return 0
}

func (m *LdpHelloInfo) GetNeighborHoldTime() uint32 {
	if m != nil {
		return m.NeighborHoldTime
	}
	return 0
}

func (m *LdpHelloInfo) GetDiscExpiry() uint32 {
	if m != nil {
		return m.DiscExpiry
	}
	return 0
}

func (m *LdpHelloInfo) GetIsTargeted() bool {
	if m != nil {
		return m.IsTargeted
	}
	return false
}

func (m *LdpHelloInfo) GetTarget() *LdpIpAddrTUnion {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *LdpHelloInfo) GetSessionUp() bool {
	if m != nil {
		return m.SessionUp
	}
	return false
}

func (m *LdpHelloInfo) GetEstablishedTime() uint64 {
	if m != nil {
		return m.EstablishedTime
	}
	return 0
}

func (m *LdpHelloInfo) GetEstablishedAge() uint64 {
	if m != nil {
		return m.EstablishedAge
	}
	return 0
}

func (m *LdpHelloInfo) GetSessionBringupFailureReason() string {
	if m != nil {
		return m.SessionBringupFailureReason
	}
	return ""
}

func (m *LdpHelloInfo) GetLastSessionDownInfo() []*LdpLastSessInfo {
	if m != nil {
		return m.LastSessionDownInfo
	}
	return nil
}

type LdpDiscoveryLinkAfInfo struct {
	LocalSrcAddress       *LdpIpAddrTUnion `protobuf:"bytes,1,opt,name=local_src_address,json=localSrcAddress,proto3" json:"local_src_address,omitempty"`
	LocalTransportAddress *LdpIpAddrTUnion `protobuf:"bytes,2,opt,name=local_transport_address,json=localTransportAddress,proto3" json:"local_transport_address,omitempty"`
	Interval              uint32           `protobuf:"varint,3,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}         `json:"-"`
	XXX_unrecognized      []byte           `json:"-"`
	XXX_sizecache         int32            `json:"-"`
}

func (m *LdpDiscoveryLinkAfInfo) Reset()         { *m = LdpDiscoveryLinkAfInfo{} }
func (m *LdpDiscoveryLinkAfInfo) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryLinkAfInfo) ProtoMessage()    {}
func (*LdpDiscoveryLinkAfInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_332f2a09e882d2e8, []int{4}
}

func (m *LdpDiscoveryLinkAfInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryLinkAfInfo.Unmarshal(m, b)
}
func (m *LdpDiscoveryLinkAfInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryLinkAfInfo.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryLinkAfInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryLinkAfInfo.Merge(m, src)
}
func (m *LdpDiscoveryLinkAfInfo) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryLinkAfInfo.Size(m)
}
func (m *LdpDiscoveryLinkAfInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryLinkAfInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryLinkAfInfo proto.InternalMessageInfo

func (m *LdpDiscoveryLinkAfInfo) GetLocalSrcAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.LocalSrcAddress
	}
	return nil
}

func (m *LdpDiscoveryLinkAfInfo) GetLocalTransportAddress() *LdpIpAddrTUnion {
	if m != nil {
		return m.LocalTransportAddress
	}
	return nil
}

func (m *LdpDiscoveryLinkAfInfo) GetInterval() uint32 {
	if m != nil {
		return m.Interval
	}
	return 0
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
	return fileDescriptor_332f2a09e882d2e8, []int{5}
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

type LdpDiscoveryLinkInfo struct {
	HelloInformation     []*LdpHelloInfo           `protobuf:"bytes,50,rep,name=hello_information,json=helloInformation,proto3" json:"hello_information,omitempty"`
	DiscoveryLinkAf      []*LdpDiscoveryLinkAfInfo `protobuf:"bytes,51,rep,name=discovery_link_af,json=discoveryLinkAf,proto3" json:"discovery_link_af,omitempty"`
	NextHello            uint32                    `protobuf:"varint,52,opt,name=next_hello,json=nextHello,proto3" json:"next_hello,omitempty"`
	Interface            string                    `protobuf:"bytes,53,opt,name=interface,proto3" json:"interface,omitempty"`
	InterfaceNameXr      string                    `protobuf:"bytes,54,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	Vrf                  *LdpVrfInfo               `protobuf:"bytes,55,opt,name=vrf,proto3" json:"vrf,omitempty"`
	QuickStartDisabled   bool                      `protobuf:"varint,56,opt,name=quick_start_disabled,json=quickStartDisabled,proto3" json:"quick_start_disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LdpDiscoveryLinkInfo) Reset()         { *m = LdpDiscoveryLinkInfo{} }
func (m *LdpDiscoveryLinkInfo) String() string { return proto.CompactTextString(m) }
func (*LdpDiscoveryLinkInfo) ProtoMessage()    {}
func (*LdpDiscoveryLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_332f2a09e882d2e8, []int{6}
}

func (m *LdpDiscoveryLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LdpDiscoveryLinkInfo.Unmarshal(m, b)
}
func (m *LdpDiscoveryLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LdpDiscoveryLinkInfo.Marshal(b, m, deterministic)
}
func (m *LdpDiscoveryLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LdpDiscoveryLinkInfo.Merge(m, src)
}
func (m *LdpDiscoveryLinkInfo) XXX_Size() int {
	return xxx_messageInfo_LdpDiscoveryLinkInfo.Size(m)
}
func (m *LdpDiscoveryLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LdpDiscoveryLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LdpDiscoveryLinkInfo proto.InternalMessageInfo

func (m *LdpDiscoveryLinkInfo) GetHelloInformation() []*LdpHelloInfo {
	if m != nil {
		return m.HelloInformation
	}
	return nil
}

func (m *LdpDiscoveryLinkInfo) GetDiscoveryLinkAf() []*LdpDiscoveryLinkAfInfo {
	if m != nil {
		return m.DiscoveryLinkAf
	}
	return nil
}

func (m *LdpDiscoveryLinkInfo) GetNextHello() uint32 {
	if m != nil {
		return m.NextHello
	}
	return 0
}

func (m *LdpDiscoveryLinkInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *LdpDiscoveryLinkInfo) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *LdpDiscoveryLinkInfo) GetVrf() *LdpVrfInfo {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *LdpDiscoveryLinkInfo) GetQuickStartDisabled() bool {
	if m != nil {
		return m.QuickStartDisabled
	}
	return false
}

func init() {
	proto.RegisterType((*LdpDiscoveryLinkInfo_KEYS)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_discovery_link_info_KEYS")
	proto.RegisterType((*LdpIpAddrTUnion)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_ip_addr_t_union")
	proto.RegisterType((*LdpLastSessInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_last_sess_info")
	proto.RegisterType((*LdpHelloInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_hello_info")
	proto.RegisterType((*LdpDiscoveryLinkAfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_discovery_link_af_info")
	proto.RegisterType((*LdpVrfInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_vrf_info")
	proto.RegisterType((*LdpDiscoveryLinkInfo)(nil), "cisco_ios_xr_mpls_ldp_oper.mpls_ldp.nodes.node.default_vrf.afs.af.discovery.link_hellos.link_hello.ldp_discovery_link_info")
}

func init() { proto.RegisterFile("ldp_discovery_link_info.proto", fileDescriptor_332f2a09e882d2e8) }

var fileDescriptor_332f2a09e882d2e8 = []byte{
	// 901 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcf, 0x6f, 0x23, 0x35,
	0x14, 0xd6, 0x34, 0x6d, 0x37, 0x79, 0xd9, 0x26, 0x8d, 0xb7, 0xa5, 0x56, 0xf7, 0x07, 0xab, 0x48,
	0xc0, 0x82, 0xd0, 0x80, 0xda, 0xa5, 0xcb, 0xb5, 0xb0, 0x8b, 0xb6, 0x62, 0xb5, 0x87, 0x69, 0x91,
	0xe0, 0x64, 0x39, 0x19, 0x4f, 0x6a, 0x75, 0x62, 0x0f, 0xb6, 0x93, 0xb6, 0x42, 0xe2, 0xc0, 0x89,
	0x1b, 0xfc, 0x07, 0x1c, 0x38, 0x01, 0x07, 0x0e, 0x88, 0x3f, 0x8e, 0x1b, 0xf2, 0x1b, 0xcf, 0x64,
	0x4a, 0xbb, 0x47, 0xb2, 0x7b, 0x89, 0x3c, 0xdf, 0xf7, 0x9e, 0xfd, 0xe5, 0xbd, 0xe7, 0xf7, 0x0c,
	0xf7, 0xf3, 0xb4, 0x60, 0xa9, 0xb4, 0x63, 0x3d, 0x17, 0xe6, 0x92, 0xe5, 0x52, 0x9d, 0x31, 0xa9,
	0x32, 0x1d, 0x17, 0x46, 0x3b, 0x4d, 0x46, 0x63, 0x4f, 0x31, 0xa9, 0x2d, 0xbb, 0x30, 0x6c, 0x5a,
	0xe4, 0x96, 0x79, 0x07, 0x5d, 0x08, 0x13, 0x57, 0x5f, 0xb1, 0xd2, 0xa9, 0xb0, 0xf8, 0x1b, 0xa7,
	0x22, 0xe3, 0xb3, 0xdc, 0xb1, 0xb9, 0xc9, 0x62, 0x9e, 0xd9, 0x98, 0x67, 0x71, 0xbd, 0x79, 0x8c,
	0x9b, 0x9f, 0x8a, 0x3c, 0xd7, 0xb6, 0xb1, 0x1e, 0x7e, 0x07, 0xf7, 0x5e, 0x21, 0x82, 0x7d, 0xf9,
	0xec, 0x9b, 0x63, 0x72, 0x17, 0x3a, 0x7e, 0x6f, 0xa6, 0xf8, 0x54, 0xd0, 0xe8, 0x61, 0xf4, 0xa8,
	0x93, 0xb4, 0x3d, 0xf0, 0x92, 0x4f, 0x05, 0xd9, 0x81, 0x5b, 0x3c, 0x2b, 0xa9, 0x15, 0xa4, 0xd6,
	0x79, 0x86, 0xc4, 0x3b, 0xd0, 0x93, 0xca, 0x09, 0x93, 0xf1, 0x71, 0x70, 0x6d, 0x21, 0xbf, 0x51,
	0xa3, 0xde, 0x6c, 0x28, 0xe0, 0x8e, 0x3f, 0x5c, 0x16, 0x8c, 0xa7, 0xa9, 0x61, 0x8e, 0xcd, 0x94,
	0xd4, 0x8a, 0x6c, 0x42, 0x8b, 0x67, 0x32, 0x9c, 0xe6, 0x97, 0x64, 0x0b, 0xd6, 0xd2, 0xd9, 0x74,
	0x7a, 0x89, 0xc7, 0x6c, 0x24, 0xe5, 0x07, 0x21, 0xb0, 0x2a, 0x8b, 0xf9, 0xe3, 0xb0, 0x37, 0xae,
	0x03, 0x76, 0x40, 0x57, 0x6b, 0xec, 0x60, 0xf8, 0x57, 0x04, 0xc4, 0x9f, 0x93, 0x73, 0xeb, 0x98,
	0x15, 0xd6, 0xe2, 0xff, 0x23, 0x4f, 0x80, 0xd6, 0x88, 0xd4, 0x8a, 0xa5, 0xfa, 0x5c, 0x31, 0x23,
	0xb8, 0xd5, 0x2a, 0x9c, 0xbd, 0xed, 0xf9, 0xe3, 0x92, 0x7e, 0xaa, 0xcf, 0x55, 0x82, 0x24, 0xd9,
	0x87, 0xb7, 0xae, 0x3b, 0x3a, 0x19, 0xa2, 0xb0, 0x9a, 0xdc, 0xf9, 0x8f, 0xdb, 0x89, 0x9c, 0x0a,
	0xf2, 0x11, 0x6c, 0x5d, 0x71, 0x9a, 0x15, 0xa5, 0x4b, 0x0b, 0xff, 0xd1, 0xa0, 0xe1, 0xf2, 0x55,
	0xe1, 0x1d, 0x86, 0xff, 0xb4, 0xa1, 0xe7, 0x55, 0x63, 0x9e, 0x4a, 0xc5, 0xbf, 0x45, 0xb0, 0xa5,
	0x84, 0x9c, 0x9c, 0x8e, 0xb4, 0x61, 0xd6, 0x8c, 0x31, 0x6c, 0xc2, 0x5a, 0x94, 0xdb, 0xdd, 0x3b,
	0x8f, 0xff, 0xff, 0x82, 0x89, 0x6f, 0x48, 0x58, 0x42, 0x2a, 0x51, 0xc7, 0x66, 0x7c, 0x58, 0x4a,
	0x22, 0x7f, 0x47, 0xb0, 0x5b, 0x6b, 0x75, 0x86, 0x2b, 0x5b, 0x68, 0xe3, 0x6a, 0xc5, 0x2b, 0xaf,
	0x57, 0x31, 0xad, 0xa4, 0x9d, 0x54, 0xca, 0x2a, 0xdd, 0x07, 0xb0, 0x53, 0xcb, 0x46, 0xcf, 0x54,
	0x28, 0x27, 0x33, 0x29, 0x4c, 0xa8, 0xb3, 0xed, 0x8a, 0x7e, 0x91, 0x16, 0x47, 0x35, 0x49, 0x1e,
	0x40, 0x57, 0x5a, 0xa6, 0x34, 0x33, 0x7a, 0xe6, 0x04, 0xd6, 0x5f, 0x3b, 0xe9, 0x48, 0xfb, 0x52,
	0x27, 0x1e, 0xf0, 0x17, 0xe9, 0x54, 0xe7, 0x69, 0x99, 0xf4, 0x35, 0x4c, 0x7a, 0xdb, 0x03, 0x58,
	0x1c, 0xef, 0x42, 0x3f, 0xd7, 0x63, 0x9e, 0xb3, 0x85, 0xc9, 0x3a, 0x9a, 0x6c, 0x20, 0xfc, 0xbc,
	0xb2, 0xfb, 0x10, 0xea, 0x50, 0x37, 0x4c, 0x6f, 0xa1, 0xe9, 0x66, 0xc5, 0xd4, 0xd6, 0x6f, 0x43,
	0xd7, 0x07, 0x87, 0x89, 0x8b, 0x42, 0x9a, 0x4b, 0xda, 0x46, 0x33, 0xf0, 0xd0, 0x33, 0x44, 0xbc,
	0x81, 0xb4, 0xcc, 0x71, 0x33, 0x11, 0x4e, 0xa4, 0xb4, 0x83, 0x9a, 0x41, 0xda, 0x93, 0x80, 0x90,
	0x9f, 0x22, 0x58, 0x2f, 0x69, 0x0a, 0xaf, 0x37, 0x61, 0x41, 0x06, 0xb9, 0x0f, 0xb0, 0xb8, 0x41,
	0xb4, 0x5b, 0x46, 0xd9, 0x56, 0x17, 0x87, 0xbc, 0x0f, 0x9b, 0xc2, 0x3a, 0x3e, 0xca, 0xa5, 0x3d,
	0x15, 0x21, 0x3c, 0xb7, 0xf1, 0x52, 0xf6, 0x1b, 0x38, 0x46, 0xe7, 0x3d, 0x68, 0x42, 0x8c, 0x4f,
	0x04, 0xdd, 0x40, 0xcb, 0x5e, 0x03, 0x3e, 0x9c, 0x08, 0xf2, 0x39, 0x3c, 0xa8, 0x8e, 0x1c, 0x19,
	0xa9, 0x26, 0xb3, 0x82, 0x65, 0x5c, 0xe6, 0x33, 0x23, 0xaa, 0x6e, 0xd1, 0xc3, 0xc2, 0xb8, 0x1b,
	0xac, 0x3e, 0x2b, 0x8d, 0xbe, 0x28, 0x6d, 0x42, 0xcf, 0xf8, 0x23, 0xba, 0xa9, 0x69, 0xf8, 0x5b,
	0x4d, 0xfb, 0x0f, 0x5b, 0x8f, 0xba, 0x7b, 0xf3, 0x65, 0x45, 0xf6, 0x6a, 0x17, 0xbc, 0xd6, 0xac,
	0x8e, 0x54, 0xa6, 0x87, 0x3f, 0xb7, 0x60, 0xf7, 0x86, 0xb1, 0xc0, 0xb3, 0xb2, 0x0f, 0xfd, 0x1a,
	0xc1, 0xa0, 0xac, 0xd7, 0x37, 0xa8, 0x09, 0x95, 0x37, 0xa8, 0xd1, 0x81, 0xfe, 0x8c, 0x60, 0xa7,
	0x54, 0xf9, 0xc6, 0xb5, 0x9f, 0x6d, 0xd4, 0x75, 0xad, 0xf7, 0xec, 0x42, 0x1b, 0x07, 0xe4, 0x9c,
	0xe7, 0x61, 0x2e, 0xd4, 0xdf, 0xc3, 0x3d, 0xb8, 0xed, 0x77, 0x9a, 0x9b, 0x90, 0x03, 0x02, 0xab,
	0x8d, 0x99, 0x8c, 0x6b, 0xd2, 0x83, 0x15, 0x99, 0x86, 0x19, 0xb9, 0x22, 0xd3, 0xe1, 0x8f, 0x6b,
	0xb0, 0xf3, 0x8a, 0xe9, 0x4e, 0x7e, 0x89, 0x60, 0xb0, 0x18, 0x2d, 0x66, 0xca, 0x9d, 0xd4, 0x8a,
	0xee, 0x61, 0x2d, 0x9a, 0x65, 0xc5, 0x65, 0x21, 0x20, 0xd9, 0xc4, 0xf5, 0xd1, 0x42, 0x0b, 0xf9,
	0x3d, 0x82, 0xc1, 0xb5, 0x02, 0xa4, 0xfb, 0xa8, 0xf0, 0xfb, 0x65, 0x29, 0xbc, 0xf9, 0x06, 0x24,
	0xfd, 0x1a, 0x7f, 0x21, 0xd5, 0xd9, 0x61, 0xe6, 0xfb, 0x92, 0x12, 0x17, 0xae, 0xf4, 0xa4, 0x8f,
	0x31, 0x05, 0x1d, 0x8f, 0x3c, 0xf7, 0x00, 0xb9, 0x07, 0x9d, 0xfa, 0xe9, 0x43, 0x3f, 0xc1, 0x94,
	0x2d, 0x00, 0xf2, 0x01, 0x0c, 0xae, 0x3e, 0x97, 0xd8, 0x85, 0xa1, 0x07, 0x68, 0xd5, 0xbf, 0xf2,
	0x62, 0xfa, 0xda, 0x90, 0x1f, 0x22, 0x68, 0xcd, 0x4d, 0x46, 0x9f, 0x60, 0x05, 0x17, 0xcb, 0x8a,
	0x43, 0x55, 0x77, 0x89, 0x3f, 0x9c, 0x7c, 0x0c, 0x5b, 0xdf, 0xce, 0xe4, 0xf8, 0x8c, 0x59, 0xc7,
	0x8d, 0xf3, 0x41, 0xe2, 0xa3, 0x5c, 0xa4, 0xf4, 0x53, 0xec, 0xc7, 0x04, 0xb9, 0x63, 0x4f, 0x3d,
	0x0d, 0xcc, 0x68, 0x1d, 0x9f, 0xb4, 0xfb, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x54, 0x83, 0x37,
	0x76, 0xf3, 0x0a, 0x00, 0x00,
}
