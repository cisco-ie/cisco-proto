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
// source: mpls_te_fwd_adj_bag.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_forwarding_adjacency_destinations_forwarding_adjacency_destination

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

type MplsTeFwdAdjBag_KEYS struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeFwdAdjBag_KEYS) Reset()         { *m = MplsTeFwdAdjBag_KEYS{} }
func (m *MplsTeFwdAdjBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsTeFwdAdjBag_KEYS) ProtoMessage()    {}
func (*MplsTeFwdAdjBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ede6f726d10987, []int{0}
}

func (m *MplsTeFwdAdjBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFwdAdjBag_KEYS.Unmarshal(m, b)
}
func (m *MplsTeFwdAdjBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFwdAdjBag_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsTeFwdAdjBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFwdAdjBag_KEYS.Merge(m, src)
}
func (m *MplsTeFwdAdjBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsTeFwdAdjBag_KEYS.Size(m)
}
func (m *MplsTeFwdAdjBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFwdAdjBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFwdAdjBag_KEYS proto.InternalMessageInfo

func (m *MplsTeFwdAdjBag_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type MplsTeAreaAfiInfo struct {
	IgpAreaId            string   `protobuf:"bytes,1,opt,name=igp_area_id,json=igpAreaId,proto3" json:"igp_area_id,omitempty"`
	Afi                  string   `protobuf:"bytes,2,opt,name=afi,proto3" json:"afi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsTeAreaAfiInfo) Reset()         { *m = MplsTeAreaAfiInfo{} }
func (m *MplsTeAreaAfiInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeAreaAfiInfo) ProtoMessage()    {}
func (*MplsTeAreaAfiInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ede6f726d10987, []int{1}
}

func (m *MplsTeAreaAfiInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Unmarshal(m, b)
}
func (m *MplsTeAreaAfiInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeAreaAfiInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeAreaAfiInfo.Merge(m, src)
}
func (m *MplsTeAreaAfiInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeAreaAfiInfo.Size(m)
}
func (m *MplsTeAreaAfiInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeAreaAfiInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeAreaAfiInfo proto.InternalMessageInfo

func (m *MplsTeAreaAfiInfo) GetIgpAreaId() string {
	if m != nil {
		return m.IgpAreaId
	}
	return ""
}

func (m *MplsTeAreaAfiInfo) GetAfi() string {
	if m != nil {
		return m.Afi
	}
	return ""
}

type MplsTeFwdAdjInfo struct {
	TunnelName           string               `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TrafficShare         uint32               `protobuf:"varint,2,opt,name=traffic_share,json=trafficShare,proto3" json:"traffic_share,omitempty"`
	HoldTime             uint32               `protobuf:"varint,3,opt,name=hold_time,json=holdTime,proto3" json:"hold_time,omitempty"`
	IsAdjacencyUp        bool                 `protobuf:"varint,4,opt,name=is_adjacency_up,json=isAdjacencyUp,proto3" json:"is_adjacency_up,omitempty"`
	IgPs                 []*MplsTeAreaAfiInfo `protobuf:"bytes,5,rep,name=ig_ps,json=igPs,proto3" json:"ig_ps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MplsTeFwdAdjInfo) Reset()         { *m = MplsTeFwdAdjInfo{} }
func (m *MplsTeFwdAdjInfo) String() string { return proto.CompactTextString(m) }
func (*MplsTeFwdAdjInfo) ProtoMessage()    {}
func (*MplsTeFwdAdjInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ede6f726d10987, []int{2}
}

func (m *MplsTeFwdAdjInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFwdAdjInfo.Unmarshal(m, b)
}
func (m *MplsTeFwdAdjInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFwdAdjInfo.Marshal(b, m, deterministic)
}
func (m *MplsTeFwdAdjInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFwdAdjInfo.Merge(m, src)
}
func (m *MplsTeFwdAdjInfo) XXX_Size() int {
	return xxx_messageInfo_MplsTeFwdAdjInfo.Size(m)
}
func (m *MplsTeFwdAdjInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFwdAdjInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFwdAdjInfo proto.InternalMessageInfo

func (m *MplsTeFwdAdjInfo) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *MplsTeFwdAdjInfo) GetTrafficShare() uint32 {
	if m != nil {
		return m.TrafficShare
	}
	return 0
}

func (m *MplsTeFwdAdjInfo) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *MplsTeFwdAdjInfo) GetIsAdjacencyUp() bool {
	if m != nil {
		return m.IsAdjacencyUp
	}
	return false
}

func (m *MplsTeFwdAdjInfo) GetIgPs() []*MplsTeAreaAfiInfo {
	if m != nil {
		return m.IgPs
	}
	return nil
}

type MplsTeFwdAdjBag struct {
	Adjacencies          uint32              `protobuf:"varint,50,opt,name=adjacencies,proto3" json:"adjacencies,omitempty"`
	Ipv6Adjacencies      uint32              `protobuf:"varint,51,opt,name=ipv6_adjacencies,json=ipv6Adjacencies,proto3" json:"ipv6_adjacencies,omitempty"`
	DestinationAddressXr string              `protobuf:"bytes,52,opt,name=destination_address_xr,json=destinationAddressXr,proto3" json:"destination_address_xr,omitempty"`
	ForwardAdjacency     []*MplsTeFwdAdjInfo `protobuf:"bytes,53,rep,name=forward_adjacency,json=forwardAdjacency,proto3" json:"forward_adjacency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MplsTeFwdAdjBag) Reset()         { *m = MplsTeFwdAdjBag{} }
func (m *MplsTeFwdAdjBag) String() string { return proto.CompactTextString(m) }
func (*MplsTeFwdAdjBag) ProtoMessage()    {}
func (*MplsTeFwdAdjBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_04ede6f726d10987, []int{3}
}

func (m *MplsTeFwdAdjBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsTeFwdAdjBag.Unmarshal(m, b)
}
func (m *MplsTeFwdAdjBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsTeFwdAdjBag.Marshal(b, m, deterministic)
}
func (m *MplsTeFwdAdjBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsTeFwdAdjBag.Merge(m, src)
}
func (m *MplsTeFwdAdjBag) XXX_Size() int {
	return xxx_messageInfo_MplsTeFwdAdjBag.Size(m)
}
func (m *MplsTeFwdAdjBag) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsTeFwdAdjBag.DiscardUnknown(m)
}

var xxx_messageInfo_MplsTeFwdAdjBag proto.InternalMessageInfo

func (m *MplsTeFwdAdjBag) GetAdjacencies() uint32 {
	if m != nil {
		return m.Adjacencies
	}
	return 0
}

func (m *MplsTeFwdAdjBag) GetIpv6Adjacencies() uint32 {
	if m != nil {
		return m.Ipv6Adjacencies
	}
	return 0
}

func (m *MplsTeFwdAdjBag) GetDestinationAddressXr() string {
	if m != nil {
		return m.DestinationAddressXr
	}
	return ""
}

func (m *MplsTeFwdAdjBag) GetForwardAdjacency() []*MplsTeFwdAdjInfo {
	if m != nil {
		return m.ForwardAdjacency
	}
	return nil
}

func init() {
	proto.RegisterType((*MplsTeFwdAdjBag_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.forwarding_adjacency_destinations.forwarding_adjacency_destination.mpls_te_fwd_adj_bag_KEYS")
	proto.RegisterType((*MplsTeAreaAfiInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.forwarding_adjacency_destinations.forwarding_adjacency_destination.mpls_te_area_afi_info")
	proto.RegisterType((*MplsTeFwdAdjInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.forwarding_adjacency_destinations.forwarding_adjacency_destination.mpls_te_fwd_adj_info")
	proto.RegisterType((*MplsTeFwdAdjBag)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.forwarding_adjacency_destinations.forwarding_adjacency_destination.mpls_te_fwd_adj_bag")
}

func init() { proto.RegisterFile("mpls_te_fwd_adj_bag.proto", fileDescriptor_04ede6f726d10987) }

var fileDescriptor_04ede6f726d10987 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0xc9, 0x6e, 0x2b, 0xdd, 0xb3, 0x2e, 0x5d, 0xa7, 0x55, 0x46, 0x04, 0x0d, 0x11, 0x64,
	0xbd, 0x89, 0xd0, 0x56, 0xef, 0x73, 0xe1, 0x45, 0x29, 0x88, 0xa4, 0x0a, 0x7a, 0x75, 0x98, 0x66,
	0x26, 0xe9, 0x91, 0xcd, 0xcc, 0x30, 0x93, 0xda, 0xaa, 0x2f, 0xe1, 0x5b, 0xf8, 0x02, 0x3e, 0x81,
	0x4f, 0x26, 0x99, 0x4d, 0x34, 0x68, 0xa0, 0x77, 0x7b, 0x97, 0x7c, 0xf3, 0xcf, 0xcf, 0x39, 0xe7,
	0x3f, 0x03, 0x0f, 0x6b, 0xbb, 0xf6, 0xd8, 0x28, 0x2c, 0xaf, 0x25, 0x0a, 0xf9, 0x09, 0x2f, 0x44,
	0x95, 0x5a, 0x67, 0x1a, 0xc3, 0x64, 0x41, 0xbe, 0x30, 0x48, 0xc6, 0xe3, 0x8d, 0xc3, 0x5e, 0x67,
	0xac, 0x72, 0x69, 0xf7, 0x93, 0x96, 0xc6, 0x5d, 0x0b, 0x27, 0x49, 0x57, 0xed, 0x5d, 0x51, 0x28,
	0x5d, 0x7c, 0x41, 0xa9, 0x7c, 0x43, 0x5a, 0x34, 0x64, 0xb4, 0xbf, 0x55, 0x91, 0x9c, 0x01, 0x1f,
	0x29, 0x01, 0xcf, 0x5e, 0x7f, 0x3c, 0x67, 0x2f, 0xe0, 0x60, 0x20, 0x45, 0x21, 0xa5, 0x53, 0xde,
	0xf3, 0x28, 0x8e, 0x56, 0xb3, 0x9c, 0x0d, 0x8e, 0xb2, 0xcd, 0x49, 0x72, 0x0a, 0xf7, 0x7b, 0x33,
	0xe1, 0x94, 0x40, 0x51, 0x12, 0x92, 0x2e, 0x0d, 0x7b, 0x0c, 0x73, 0xaa, 0xec, 0x06, 0x92, 0xec,
	0x1c, 0x66, 0x54, 0xd9, 0xcc, 0x29, 0x71, 0x2a, 0xd9, 0x12, 0xa6, 0xa2, 0x24, 0x3e, 0x09, 0xbc,
	0xfd, 0x4c, 0x7e, 0x4e, 0xe0, 0xf0, 0xdf, 0xc2, 0x82, 0xd5, 0x13, 0x98, 0x37, 0x57, 0x5a, 0xab,
	0x35, 0x6a, 0x51, 0xab, 0xce, 0x0a, 0x36, 0xe8, 0x8d, 0xa8, 0x15, 0x7b, 0x0a, 0x8b, 0xc6, 0x89,
	0xb2, 0xa4, 0x02, 0xfd, 0xa5, 0x70, 0x2a, 0xb8, 0x2e, 0xf2, 0xbb, 0x1d, 0x3c, 0x6f, 0x19, 0x7b,
	0x04, 0xb3, 0x4b, 0xb3, 0x96, 0xd8, 0x50, 0xad, 0xf8, 0x34, 0x08, 0xf6, 0x5a, 0xf0, 0x8e, 0x6a,
	0xc5, 0x9e, 0xc1, 0x3e, 0xf9, 0xc1, 0xbc, 0xae, 0x2c, 0xdf, 0x89, 0xa3, 0xd5, 0x5e, 0xbe, 0x20,
	0x9f, 0xf5, 0xf4, 0xbd, 0x65, 0xdf, 0x23, 0xd8, 0xa5, 0x0a, 0xad, 0xe7, 0xbb, 0xf1, 0x74, 0x35,
	0x3f, 0xfa, 0x96, 0x6e, 0x23, 0xb2, 0x74, 0x74, 0xc4, 0xf9, 0x0e, 0x55, 0x6f, 0x7d, 0xf2, 0x6b,
	0x02, 0x07, 0x23, 0x79, 0xb2, 0x18, 0xe6, 0xbd, 0x19, 0x29, 0xcf, 0x8f, 0x42, 0xc7, 0x43, 0xc4,
	0x9e, 0xc3, 0x92, 0xec, 0xe7, 0x57, 0x38, 0x94, 0x1d, 0x07, 0xd9, 0x7e, 0xcb, 0xb3, 0x81, 0xf4,
	0x04, 0x1e, 0x8c, 0xec, 0x05, 0xde, 0x38, 0x7e, 0x12, 0xd2, 0x38, 0xfc, 0x7f, 0x35, 0x3e, 0x38,
	0xf6, 0x23, 0x82, 0x7b, 0x5d, 0x6f, 0x7f, 0x1b, 0xe3, 0x2f, 0xc3, 0xe4, 0xbe, 0x6e, 0x77, 0x72,
	0xc3, 0x85, 0xca, 0x97, 0xdd, 0xb5, 0x3f, 0xc9, 0x5e, 0xdc, 0x09, 0x0f, 0xf0, 0xf8, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xdd, 0xa8, 0x9f, 0x16, 0x9d, 0x03, 0x00, 0x00,
}