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
// source: l2vpn_evpn_evi_bgp_rt.proto

package cisco_ios_xr_evpn_oper_evpn_active_evi_detail_evi_children_route_targets_route_target

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

type L2VpnEvpnEviBgpRt_KEYS struct {
	Evi                  uint32   `protobuf:"varint,1,opt,name=evi,proto3" json:"evi,omitempty"`
	Encapsulation        uint32   `protobuf:"varint,2,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Format               string   `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	As                   uint32   `protobuf:"varint,5,opt,name=as,proto3" json:"as,omitempty"`
	AsIndex              uint32   `protobuf:"varint,6,opt,name=as_index,json=asIndex,proto3" json:"as_index,omitempty"`
	AddrIndex            uint32   `protobuf:"varint,7,opt,name=addr_index,json=addrIndex,proto3" json:"addr_index,omitempty"`
	Address              string   `protobuf:"bytes,8,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEviBgpRt_KEYS) Reset()         { *m = L2VpnEvpnEviBgpRt_KEYS{} }
func (m *L2VpnEvpnEviBgpRt_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviBgpRt_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnEviBgpRt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{0}
}

func (m *L2VpnEvpnEviBgpRt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviBgpRt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviBgpRt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnEviBgpRt_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS.Size(m)
}
func (m *L2VpnEvpnEviBgpRt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviBgpRt_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnEviBgpRt_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetEncapsulation() uint32 {
	if m != nil {
		return m.Encapsulation
	}
	return 0
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetAsIndex() uint32 {
	if m != nil {
		return m.AsIndex
	}
	return 0
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetAddrIndex() uint32 {
	if m != nil {
		return m.AddrIndex
	}
	return 0
}

func (m *L2VpnEvpnEviBgpRt_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type L2VpnEvpnEviSummary struct {
	EthernetVpnId        uint32   `protobuf:"varint,1,opt,name=ethernet_vpn_id,json=ethernetVpnId,proto3" json:"ethernet_vpn_id,omitempty"`
	EncapsulationXr      uint32   `protobuf:"varint,2,opt,name=encapsulation_xr,json=encapsulationXr,proto3" json:"encapsulation_xr,omitempty"`
	BdName               string   `protobuf:"bytes,3,opt,name=bd_name,json=bdName,proto3" json:"bd_name,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnEviSummary) Reset()         { *m = L2VpnEvpnEviSummary{} }
func (m *L2VpnEvpnEviSummary) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviSummary) ProtoMessage()    {}
func (*L2VpnEvpnEviSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{1}
}

func (m *L2VpnEvpnEviSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviSummary.Merge(m, src)
}
func (m *L2VpnEvpnEviSummary) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviSummary.Size(m)
}
func (m *L2VpnEvpnEviSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviSummary.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviSummary proto.InternalMessageInfo

func (m *L2VpnEvpnEviSummary) GetEthernetVpnId() uint32 {
	if m != nil {
		return m.EthernetVpnId
	}
	return 0
}

func (m *L2VpnEvpnEviSummary) GetEncapsulationXr() uint32 {
	if m != nil {
		return m.EncapsulationXr
	}
	return 0
}

func (m *L2VpnEvpnEviSummary) GetBdName() string {
	if m != nil {
		return m.BdName
	}
	return ""
}

func (m *L2VpnEvpnEviSummary) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type L2VpnRt_2ByteAs struct {
	TwoByteAs            uint32   `protobuf:"varint,1,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteIndex        uint32   `protobuf:"varint,2,opt,name=four_byte_index,json=fourByteIndex,proto3" json:"four_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRt_2ByteAs) Reset()         { *m = L2VpnRt_2ByteAs{} }
func (m *L2VpnRt_2ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt_2ByteAs) ProtoMessage()    {}
func (*L2VpnRt_2ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{2}
}

func (m *L2VpnRt_2ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRt_2ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRt_2ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt_2ByteAs.Merge(m, src)
}
func (m *L2VpnRt_2ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt_2ByteAs.Size(m)
}
func (m *L2VpnRt_2ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt_2ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt_2ByteAs proto.InternalMessageInfo

func (m *L2VpnRt_2ByteAs) GetTwoByteAs() uint32 {
	if m != nil {
		return m.TwoByteAs
	}
	return 0
}

func (m *L2VpnRt_2ByteAs) GetFourByteIndex() uint32 {
	if m != nil {
		return m.FourByteIndex
	}
	return 0
}

type L2VpnRt_4ByteAs struct {
	FourByteAs           uint32   `protobuf:"varint,1,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRt_4ByteAs) Reset()         { *m = L2VpnRt_4ByteAs{} }
func (m *L2VpnRt_4ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt_4ByteAs) ProtoMessage()    {}
func (*L2VpnRt_4ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{3}
}

func (m *L2VpnRt_4ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Unmarshal(m, b)
}
func (m *L2VpnRt_4ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnRt_4ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt_4ByteAs.Merge(m, src)
}
func (m *L2VpnRt_4ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt_4ByteAs.Size(m)
}
func (m *L2VpnRt_4ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt_4ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt_4ByteAs proto.InternalMessageInfo

func (m *L2VpnRt_4ByteAs) GetFourByteAs() uint32 {
	if m != nil {
		return m.FourByteAs
	}
	return 0
}

func (m *L2VpnRt_4ByteAs) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRtV4Addr struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRtV4Addr) Reset()         { *m = L2VpnRtV4Addr{} }
func (m *L2VpnRtV4Addr) String() string { return proto.CompactTextString(m) }
func (*L2VpnRtV4Addr) ProtoMessage()    {}
func (*L2VpnRtV4Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{4}
}

func (m *L2VpnRtV4Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRtV4Addr.Unmarshal(m, b)
}
func (m *L2VpnRtV4Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRtV4Addr.Marshal(b, m, deterministic)
}
func (m *L2VpnRtV4Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRtV4Addr.Merge(m, src)
}
func (m *L2VpnRtV4Addr) XXX_Size() int {
	return xxx_messageInfo_L2VpnRtV4Addr.Size(m)
}
func (m *L2VpnRtV4Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRtV4Addr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRtV4Addr proto.InternalMessageInfo

func (m *L2VpnRtV4Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *L2VpnRtV4Addr) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnRtEsImport struct {
	HighBytes            uint32   `protobuf:"varint,1,opt,name=high_bytes,json=highBytes,proto3" json:"high_bytes,omitempty"`
	LowBytes             uint32   `protobuf:"varint,2,opt,name=low_bytes,json=lowBytes,proto3" json:"low_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnRtEsImport) Reset()         { *m = L2VpnRtEsImport{} }
func (m *L2VpnRtEsImport) String() string { return proto.CompactTextString(m) }
func (*L2VpnRtEsImport) ProtoMessage()    {}
func (*L2VpnRtEsImport) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{5}
}

func (m *L2VpnRtEsImport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRtEsImport.Unmarshal(m, b)
}
func (m *L2VpnRtEsImport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRtEsImport.Marshal(b, m, deterministic)
}
func (m *L2VpnRtEsImport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRtEsImport.Merge(m, src)
}
func (m *L2VpnRtEsImport) XXX_Size() int {
	return xxx_messageInfo_L2VpnRtEsImport.Size(m)
}
func (m *L2VpnRtEsImport) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRtEsImport.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRtEsImport proto.InternalMessageInfo

func (m *L2VpnRtEsImport) GetHighBytes() uint32 {
	if m != nil {
		return m.HighBytes
	}
	return 0
}

func (m *L2VpnRtEsImport) GetLowBytes() uint32 {
	if m != nil {
		return m.LowBytes
	}
	return 0
}

type L2VpnRt struct {
	Rt                   string           `protobuf:"bytes,1,opt,name=rt,proto3" json:"rt,omitempty"`
	TwoByteAs            *L2VpnRt_2ByteAs `protobuf:"bytes,2,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteAs           *L2VpnRt_4ByteAs `protobuf:"bytes,3,opt,name=four_byte_as,json=fourByteAs,proto3" json:"four_byte_as,omitempty"`
	V4Addr               *L2VpnRtV4Addr   `protobuf:"bytes,4,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	EsImport             *L2VpnRtEsImport `protobuf:"bytes,5,opt,name=es_import,json=esImport,proto3" json:"es_import,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *L2VpnRt) Reset()         { *m = L2VpnRt{} }
func (m *L2VpnRt) String() string { return proto.CompactTextString(m) }
func (*L2VpnRt) ProtoMessage()    {}
func (*L2VpnRt) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{6}
}

func (m *L2VpnRt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnRt.Unmarshal(m, b)
}
func (m *L2VpnRt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnRt.Marshal(b, m, deterministic)
}
func (m *L2VpnRt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnRt.Merge(m, src)
}
func (m *L2VpnRt) XXX_Size() int {
	return xxx_messageInfo_L2VpnRt.Size(m)
}
func (m *L2VpnRt) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnRt.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnRt proto.InternalMessageInfo

func (m *L2VpnRt) GetRt() string {
	if m != nil {
		return m.Rt
	}
	return ""
}

func (m *L2VpnRt) GetTwoByteAs() *L2VpnRt_2ByteAs {
	if m != nil {
		return m.TwoByteAs
	}
	return nil
}

func (m *L2VpnRt) GetFourByteAs() *L2VpnRt_4ByteAs {
	if m != nil {
		return m.FourByteAs
	}
	return nil
}

func (m *L2VpnRt) GetV4Addr() *L2VpnRtV4Addr {
	if m != nil {
		return m.V4Addr
	}
	return nil
}

func (m *L2VpnRt) GetEsImport() *L2VpnRtEsImport {
	if m != nil {
		return m.EsImport
	}
	return nil
}

type L2VpnEvpnEviBgpRt struct {
	EvpnInstance         *L2VpnEvpnEviSummary `protobuf:"bytes,50,opt,name=evpn_instance,json=evpnInstance,proto3" json:"evpn_instance,omitempty"`
	RouteTarget          *L2VpnRt             `protobuf:"bytes,51,opt,name=route_target,json=routeTarget,proto3" json:"route_target,omitempty"`
	RouteTargetRole      string               `protobuf:"bytes,52,opt,name=route_target_role,json=routeTargetRole,proto3" json:"route_target_role,omitempty"`
	RouteTargetStitching bool                 `protobuf:"varint,53,opt,name=route_target_stitching,json=routeTargetStitching,proto3" json:"route_target_stitching,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *L2VpnEvpnEviBgpRt) Reset()         { *m = L2VpnEvpnEviBgpRt{} }
func (m *L2VpnEvpnEviBgpRt) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnEviBgpRt) ProtoMessage()    {}
func (*L2VpnEvpnEviBgpRt) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c7229cd1c3513f, []int{7}
}

func (m *L2VpnEvpnEviBgpRt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt.Unmarshal(m, b)
}
func (m *L2VpnEvpnEviBgpRt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnEviBgpRt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnEviBgpRt.Merge(m, src)
}
func (m *L2VpnEvpnEviBgpRt) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnEviBgpRt.Size(m)
}
func (m *L2VpnEvpnEviBgpRt) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnEviBgpRt.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnEviBgpRt proto.InternalMessageInfo

func (m *L2VpnEvpnEviBgpRt) GetEvpnInstance() *L2VpnEvpnEviSummary {
	if m != nil {
		return m.EvpnInstance
	}
	return nil
}

func (m *L2VpnEvpnEviBgpRt) GetRouteTarget() *L2VpnRt {
	if m != nil {
		return m.RouteTarget
	}
	return nil
}

func (m *L2VpnEvpnEviBgpRt) GetRouteTargetRole() string {
	if m != nil {
		return m.RouteTargetRole
	}
	return ""
}

func (m *L2VpnEvpnEviBgpRt) GetRouteTargetStitching() bool {
	if m != nil {
		return m.RouteTargetStitching
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnEvpnEviBgpRt_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_evpn_evi_bgp_rt_KEYS")
	proto.RegisterType((*L2VpnEvpnEviSummary)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_evpn_evi_summary")
	proto.RegisterType((*L2VpnRt_2ByteAs)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_rt_2byte_as")
	proto.RegisterType((*L2VpnRt_4ByteAs)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_rt_4byte_as")
	proto.RegisterType((*L2VpnRtV4Addr)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_rt_v4addr")
	proto.RegisterType((*L2VpnRtEsImport)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_rt_es_import")
	proto.RegisterType((*L2VpnRt)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_rt")
	proto.RegisterType((*L2VpnEvpnEviBgpRt)(nil), "cisco_ios_xr_evpn_oper.evpn.active.evi_detail.evi_children.route_targets.route_target.l2vpn_evpn_evi_bgp_rt")
}

func init() { proto.RegisterFile("l2vpn_evpn_evi_bgp_rt.proto", fileDescriptor_54c7229cd1c3513f) }

var fileDescriptor_54c7229cd1c3513f = []byte{
	// 663 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x96, 0x93, 0x92, 0x9f, 0x49, 0xd2, 0xb4, 0x16, 0x94, 0x85, 0x0a, 0x14, 0xa2, 0x0a, 0x15,
	0x0e, 0x7b, 0x48, 0xc3, 0x03, 0x14, 0x89, 0x43, 0x84, 0x84, 0xd0, 0x16, 0x10, 0x3f, 0x07, 0xcb,
	0xd9, 0x75, 0x13, 0x4b, 0xbb, 0xeb, 0x95, 0xed, 0x24, 0xcd, 0x09, 0x89, 0x03, 0x42, 0x1c, 0xe1,
	0xc2, 0xc3, 0xf0, 0x38, 0x3c, 0x08, 0xb2, 0xd7, 0xbb, 0x64, 0x43, 0x0f, 0x5c, 0xc2, 0x25, 0x9a,
	0xf9, 0x66, 0xe6, 0xf3, 0xe7, 0x99, 0xc9, 0x1a, 0x8e, 0xe3, 0xd1, 0x32, 0x4b, 0x09, 0xcb, 0x7f,
	0x38, 0x99, 0xce, 0x32, 0x22, 0xb5, 0x9f, 0x49, 0xa1, 0x05, 0x7e, 0x1d, 0x72, 0x15, 0x0a, 0xc2,
	0x85, 0x22, 0x57, 0x32, 0xcf, 0x11, 0x19, 0x93, 0xbe, 0xb1, 0x7c, 0x1a, 0x6a, 0xbe, 0x64, 0xbe,
	0x29, 0x8a, 0x98, 0xa6, 0x3c, 0xb6, 0x66, 0x38, 0xe7, 0x71, 0x24, 0x59, 0xea, 0x4b, 0xb1, 0xd0,
	0x8c, 0x68, 0x2a, 0x67, 0x4c, 0xab, 0x8a, 0x37, 0xfc, 0x85, 0xe0, 0xee, 0xb5, 0xc7, 0x92, 0xe7,
	0xcf, 0xde, 0x5d, 0xe0, 0x03, 0xa8, 0xb3, 0x25, 0xf7, 0xd0, 0x00, 0x9d, 0xf6, 0x02, 0x63, 0xe2,
	0x13, 0xe8, 0xb1, 0x34, 0xa4, 0x99, 0x5a, 0xc4, 0x54, 0x73, 0x91, 0x7a, 0x35, 0x1b, 0xab, 0x82,
	0x18, 0xc3, 0x9e, 0x14, 0x31, 0xf3, 0xea, 0x03, 0x74, 0xda, 0x0e, 0xac, 0x8d, 0x8f, 0xa0, 0x71,
	0x29, 0x64, 0x42, 0xb5, 0xb7, 0x67, 0x51, 0xe7, 0xe1, 0x7d, 0xa8, 0x51, 0xe5, 0xdd, 0xb0, 0x34,
	0x35, 0xaa, 0xf0, 0x1d, 0x68, 0x51, 0x45, 0x78, 0x1a, 0xb1, 0x2b, 0xaf, 0x61, 0xd1, 0x26, 0x55,
	0x13, 0xe3, 0xe2, 0x7b, 0x00, 0x34, 0x8a, 0xa4, 0x0b, 0x36, 0x6d, 0xb0, 0x6d, 0x90, 0x3c, 0xec,
	0x41, 0xd3, 0x38, 0x4c, 0x29, 0xaf, 0x65, 0x8f, 0x28, 0xdc, 0xe1, 0x0f, 0x04, 0x47, 0x5b, 0xd7,
	0x54, 0x8b, 0x24, 0xa1, 0x72, 0x8d, 0x1f, 0x42, 0x9f, 0xe9, 0x39, 0x93, 0x29, 0xd3, 0xc4, 0xc4,
	0x78, 0xe4, 0xae, 0xdb, 0x2b, 0xe0, 0x37, 0x59, 0x3a, 0x89, 0xf0, 0x23, 0x38, 0xa8, 0xdc, 0x91,
	0x5c, 0x49, 0x77, 0xf7, 0x7e, 0x05, 0x7f, 0x2b, 0xf1, 0x6d, 0x68, 0x4e, 0x23, 0x92, 0xd2, 0xa4,
	0x68, 0x40, 0x63, 0x1a, 0xbd, 0xa0, 0x09, 0x33, 0x6d, 0xd1, 0xeb, 0x8c, 0xb9, 0x06, 0x58, 0x7b,
	0xf8, 0x01, 0x0e, 0x73, 0x65, 0x52, 0x93, 0xd1, 0x74, 0xad, 0x19, 0xa1, 0x0a, 0xdf, 0x87, 0x8e,
	0x5e, 0x09, 0xe2, 0x5c, 0x27, 0xa8, 0xad, 0x57, 0xe2, 0xe9, 0x5a, 0xb3, 0x73, 0x65, 0x44, 0x5f,
	0x8a, 0x85, 0xcc, 0x13, 0xf2, 0x6e, 0xb8, 0x39, 0x18, 0xd8, 0x24, 0xd9, 0x8e, 0x54, 0xc8, 0xc7,
	0x05, 0xf9, 0x00, 0xba, 0x7f, 0x8a, 0x4b, 0x76, 0x28, 0x2a, 0xcf, 0x15, 0x3e, 0x81, 0xfd, 0xf2,
	0xf8, 0x4d, 0xf6, 0xae, 0x53, 0x90, 0x93, 0xbf, 0x87, 0x7e, 0x49, 0xbe, 0x1c, 0x9b, 0x56, 0xe3,
	0x07, 0xd0, 0xe5, 0xd9, 0x72, 0x4c, 0x8a, 0x31, 0x20, 0x7b, 0xd1, 0x8e, 0xc1, 0xce, 0x73, 0xe8,
	0x1f, 0xb9, 0x5f, 0x02, 0x2e, 0xb9, 0x99, 0x22, 0x3c, 0xc9, 0x84, 0xd4, 0x66, 0xfe, 0x73, 0x3e,
	0x9b, 0xdb, 0xe2, 0xb2, 0x2b, 0x06, 0x31, 0x85, 0x0a, 0x1f, 0x43, 0x3b, 0x16, 0x2b, 0x17, 0xcd,
	0x59, 0x5b, 0xb1, 0x58, 0xd9, 0xe0, 0xf0, 0xe7, 0x1e, 0xb4, 0x0a, 0x4a, 0xb3, 0x73, 0x52, 0x3b,
	0x75, 0x35, 0xa9, 0xf1, 0x17, 0x54, 0x6d, 0xb8, 0x29, 0xee, 0x8c, 0xe6, 0xfe, 0x4e, 0xfe, 0x74,
	0xfe, 0x5f, 0xf3, 0xde, 0x1c, 0xed, 0x57, 0xb4, 0x35, 0x9e, 0xfa, 0xff, 0xd1, 0x52, 0xac, 0x47,
	0x65, 0x11, 0x3e, 0x42, 0xd3, 0x4d, 0xd3, 0xee, 0x6c, 0x67, 0x74, 0xb9, 0x6b, 0x19, 0xf9, 0x22,
	0x05, 0x8d, 0x7c, 0x61, 0xf0, 0x67, 0x04, 0xed, 0x72, 0xfe, 0xf6, 0x23, 0xd1, 0x19, 0xf1, 0x5d,
	0x6b, 0x28, 0x0f, 0x0c, 0x5a, 0x4c, 0x4d, 0xac, 0x35, 0xfc, 0x5e, 0x87, 0x5b, 0xd7, 0x7e, 0x28,
	0xf1, 0x37, 0x04, 0x3d, 0x8b, 0xf1, 0x54, 0x69, 0x9a, 0x86, 0xcc, 0x1b, 0x59, 0x99, 0xc9, 0x4e,
	0x65, 0x6e, 0x7f, 0xc7, 0x82, 0xae, 0x41, 0x26, 0x4e, 0x02, 0xfe, 0x84, 0xa0, 0xbb, 0x59, 0xe5,
	0x9d, 0x59, 0x4d, 0x64, 0xc7, 0xad, 0x0b, 0x3a, 0x16, 0x7e, 0x65, 0x51, 0xfc, 0x18, 0x0e, 0x37,
	0xb3, 0x88, 0x7d, 0x12, 0xc6, 0xf6, 0x4f, 0xd7, 0xdf, 0xc8, 0x0b, 0xcc, 0xeb, 0x30, 0x86, 0xa3,
	0x4a, 0xae, 0xd2, 0x5c, 0x87, 0x73, 0x9e, 0xce, 0xbc, 0x27, 0x03, 0x74, 0xda, 0x0a, 0x6e, 0x6e,
	0x14, 0x5c, 0x14, 0xb1, 0x69, 0xc3, 0x3e, 0x8e, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3f,
	0xf0, 0x4d, 0xa7, 0x3b, 0x07, 0x00, 0x00,
}
