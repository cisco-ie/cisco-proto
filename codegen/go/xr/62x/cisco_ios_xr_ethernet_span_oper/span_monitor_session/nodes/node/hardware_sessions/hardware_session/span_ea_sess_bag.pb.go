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
// source: span_ea_sess_bag.proto

package cisco_ios_xr_ethernet_span_oper_span_monitor_session_nodes_node_hardware_sessions_hardware_session

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

type SpanEaSessBag_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SessionClass         string   `protobuf:"bytes,2,opt,name=session_class,json=sessionClass,proto3" json:"session_class,omitempty"`
	SessionId            int32    `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanEaSessBag_KEYS) Reset()         { *m = SpanEaSessBag_KEYS{} }
func (m *SpanEaSessBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*SpanEaSessBag_KEYS) ProtoMessage()    {}
func (*SpanEaSessBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011381c324faa19, []int{0}
}

func (m *SpanEaSessBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanEaSessBag_KEYS.Unmarshal(m, b)
}
func (m *SpanEaSessBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanEaSessBag_KEYS.Marshal(b, m, deterministic)
}
func (m *SpanEaSessBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanEaSessBag_KEYS.Merge(m, src)
}
func (m *SpanEaSessBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_SpanEaSessBag_KEYS.Size(m)
}
func (m *SpanEaSessBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanEaSessBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SpanEaSessBag_KEYS proto.InternalMessageInfo

func (m *SpanEaSessBag_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *SpanEaSessBag_KEYS) GetSessionClass() string {
	if m != nil {
		return m.SessionClass
	}
	return ""
}

func (m *SpanEaSessBag_KEYS) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

type SpanBagDstIdNexthopV4_ struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanBagDstIdNexthopV4_) Reset()         { *m = SpanBagDstIdNexthopV4_{} }
func (m *SpanBagDstIdNexthopV4_) String() string { return proto.CompactTextString(m) }
func (*SpanBagDstIdNexthopV4_) ProtoMessage()    {}
func (*SpanBagDstIdNexthopV4_) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011381c324faa19, []int{1}
}

func (m *SpanBagDstIdNexthopV4_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBagDstIdNexthopV4_.Unmarshal(m, b)
}
func (m *SpanBagDstIdNexthopV4_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBagDstIdNexthopV4_.Marshal(b, m, deterministic)
}
func (m *SpanBagDstIdNexthopV4_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBagDstIdNexthopV4_.Merge(m, src)
}
func (m *SpanBagDstIdNexthopV4_) XXX_Size() int {
	return xxx_messageInfo_SpanBagDstIdNexthopV4_.Size(m)
}
func (m *SpanBagDstIdNexthopV4_) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBagDstIdNexthopV4_.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBagDstIdNexthopV4_ proto.InternalMessageInfo

func (m *SpanBagDstIdNexthopV4_) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *SpanBagDstIdNexthopV4_) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type SpanBagDstIdNexthopV6_ struct {
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanBagDstIdNexthopV6_) Reset()         { *m = SpanBagDstIdNexthopV6_{} }
func (m *SpanBagDstIdNexthopV6_) String() string { return proto.CompactTextString(m) }
func (*SpanBagDstIdNexthopV6_) ProtoMessage()    {}
func (*SpanBagDstIdNexthopV6_) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011381c324faa19, []int{2}
}

func (m *SpanBagDstIdNexthopV6_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBagDstIdNexthopV6_.Unmarshal(m, b)
}
func (m *SpanBagDstIdNexthopV6_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBagDstIdNexthopV6_.Marshal(b, m, deterministic)
}
func (m *SpanBagDstIdNexthopV6_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBagDstIdNexthopV6_.Merge(m, src)
}
func (m *SpanBagDstIdNexthopV6_) XXX_Size() int {
	return xxx_messageInfo_SpanBagDstIdNexthopV6_.Size(m)
}
func (m *SpanBagDstIdNexthopV6_) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBagDstIdNexthopV6_.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBagDstIdNexthopV6_ proto.InternalMessageInfo

func (m *SpanBagDstIdNexthopV6_) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func (m *SpanBagDstIdNexthopV6_) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type SpanBagDstId struct {
	DestinationClass     string                  `protobuf:"bytes,1,opt,name=destination_class,json=destinationClass,proto3" json:"destination_class,omitempty"`
	Interface            string                  `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	PseudowireId         uint32                  `protobuf:"varint,3,opt,name=pseudowire_id,json=pseudowireId,proto3" json:"pseudowire_id,omitempty"`
	Ipv4AddressAndVrf    *SpanBagDstIdNexthopV4_ `protobuf:"bytes,4,opt,name=ipv4_address_and_vrf,json=ipv4AddressAndVrf,proto3" json:"ipv4_address_and_vrf,omitempty"`
	Ipv6AddressAndVrf    *SpanBagDstIdNexthopV6_ `protobuf:"bytes,5,opt,name=ipv6_address_and_vrf,json=ipv6AddressAndVrf,proto3" json:"ipv6_address_and_vrf,omitempty"`
	InvalidValue         uint32                  `protobuf:"varint,6,opt,name=invalid_value,json=invalidValue,proto3" json:"invalid_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SpanBagDstId) Reset()         { *m = SpanBagDstId{} }
func (m *SpanBagDstId) String() string { return proto.CompactTextString(m) }
func (*SpanBagDstId) ProtoMessage()    {}
func (*SpanBagDstId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011381c324faa19, []int{3}
}

func (m *SpanBagDstId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBagDstId.Unmarshal(m, b)
}
func (m *SpanBagDstId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBagDstId.Marshal(b, m, deterministic)
}
func (m *SpanBagDstId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBagDstId.Merge(m, src)
}
func (m *SpanBagDstId) XXX_Size() int {
	return xxx_messageInfo_SpanBagDstId.Size(m)
}
func (m *SpanBagDstId) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBagDstId.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBagDstId proto.InternalMessageInfo

func (m *SpanBagDstId) GetDestinationClass() string {
	if m != nil {
		return m.DestinationClass
	}
	return ""
}

func (m *SpanBagDstId) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *SpanBagDstId) GetPseudowireId() uint32 {
	if m != nil {
		return m.PseudowireId
	}
	return 0
}

func (m *SpanBagDstId) GetIpv4AddressAndVrf() *SpanBagDstIdNexthopV4_ {
	if m != nil {
		return m.Ipv4AddressAndVrf
	}
	return nil
}

func (m *SpanBagDstId) GetIpv6AddressAndVrf() *SpanBagDstIdNexthopV6_ {
	if m != nil {
		return m.Ipv6AddressAndVrf
	}
	return nil
}

func (m *SpanBagDstId) GetInvalidValue() uint32 {
	if m != nil {
		return m.InvalidValue
	}
	return 0
}

type SpanEaSessBag struct {
	Id                   uint32        `protobuf:"varint,50,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string        `protobuf:"bytes,51,opt,name=name,proto3" json:"name,omitempty"`
	SessionClassXr       string        `protobuf:"bytes,52,opt,name=session_class_xr,json=sessionClassXr,proto3" json:"session_class_xr,omitempty"`
	DestinationInterface string        `protobuf:"bytes,53,opt,name=destination_interface,json=destinationInterface,proto3" json:"destination_interface,omitempty"`
	DestinationId        *SpanBagDstId `protobuf:"bytes,54,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	PlatformError        uint32        `protobuf:"varint,55,opt,name=platform_error,json=platformError,proto3" json:"platform_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SpanEaSessBag) Reset()         { *m = SpanEaSessBag{} }
func (m *SpanEaSessBag) String() string { return proto.CompactTextString(m) }
func (*SpanEaSessBag) ProtoMessage()    {}
func (*SpanEaSessBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2011381c324faa19, []int{4}
}

func (m *SpanEaSessBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanEaSessBag.Unmarshal(m, b)
}
func (m *SpanEaSessBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanEaSessBag.Marshal(b, m, deterministic)
}
func (m *SpanEaSessBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanEaSessBag.Merge(m, src)
}
func (m *SpanEaSessBag) XXX_Size() int {
	return xxx_messageInfo_SpanEaSessBag.Size(m)
}
func (m *SpanEaSessBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanEaSessBag.DiscardUnknown(m)
}

var xxx_messageInfo_SpanEaSessBag proto.InternalMessageInfo

func (m *SpanEaSessBag) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpanEaSessBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SpanEaSessBag) GetSessionClassXr() string {
	if m != nil {
		return m.SessionClassXr
	}
	return ""
}

func (m *SpanEaSessBag) GetDestinationInterface() string {
	if m != nil {
		return m.DestinationInterface
	}
	return ""
}

func (m *SpanEaSessBag) GetDestinationId() *SpanBagDstId {
	if m != nil {
		return m.DestinationId
	}
	return nil
}

func (m *SpanEaSessBag) GetPlatformError() uint32 {
	if m != nil {
		return m.PlatformError
	}
	return 0
}

func init() {
	proto.RegisterType((*SpanEaSessBag_KEYS)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.hardware_sessions.hardware_session.span_ea_sess_bag_KEYS")
	proto.RegisterType((*SpanBagDstIdNexthopV4_)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.hardware_sessions.hardware_session.span_bag_dst_id_nexthop_v4_")
	proto.RegisterType((*SpanBagDstIdNexthopV6_)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.hardware_sessions.hardware_session.span_bag_dst_id_nexthop_v6_")
	proto.RegisterType((*SpanBagDstId)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.hardware_sessions.hardware_session.span_bag_dst_id")
	proto.RegisterType((*SpanEaSessBag)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.hardware_sessions.hardware_session.span_ea_sess_bag")
}

func init() { proto.RegisterFile("span_ea_sess_bag.proto", fileDescriptor_2011381c324faa19) }

var fileDescriptor_2011381c324faa19 = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xc1, 0x8a, 0x13, 0x41,
	0x10, 0x65, 0x62, 0x76, 0x35, 0xb5, 0x49, 0xcc, 0x36, 0xbb, 0x32, 0xa2, 0x42, 0xcc, 0x22, 0x04,
	0x84, 0x39, 0x6c, 0xe2, 0x78, 0x5e, 0x64, 0x0f, 0x41, 0xf0, 0x10, 0x61, 0x51, 0x3c, 0x14, 0x9d,
	0x74, 0x8f, 0x69, 0xc8, 0x74, 0x0f, 0xd5, 0xbd, 0xb3, 0xb9, 0xf9, 0x03, 0x7e, 0x81, 0x5f, 0xe1,
	0xef, 0xf8, 0x37, 0x32, 0x3d, 0x33, 0x66, 0x92, 0x95, 0xc5, 0x8b, 0xe8, 0x65, 0x68, 0x5e, 0x55,
	0xbd, 0x7a, 0xbc, 0xaa, 0x1a, 0x78, 0x64, 0x33, 0xae, 0x51, 0x72, 0xb4, 0xd2, 0x5a, 0x5c, 0xf0,
	0xcf, 0x51, 0x46, 0xc6, 0x19, 0xb6, 0x58, 0x2a, 0xbb, 0x34, 0xa8, 0x8c, 0xc5, 0x0d, 0xa1, 0x74,
	0x2b, 0x49, 0x5a, 0x3a, 0xf4, 0xd9, 0x26, 0x93, 0x14, 0xf9, 0x57, 0x6a, 0xb4, 0x72, 0x86, 0x7c,
	0xb1, 0x32, 0x3a, 0xd2, 0x46, 0x48, 0xeb, 0xbf, 0xd1, 0x8a, 0x93, 0xb8, 0xe1, 0x24, 0xeb, 0x98,
	0xbd, 0x85, 0x8c, 0x0c, 0x9c, 0xee, 0x77, 0xc7, 0xb7, 0x97, 0x1f, 0xdf, 0x33, 0x06, 0xed, 0x82,
	0x23, 0x0c, 0x86, 0xc1, 0xb8, 0x33, 0xf7, 0x6f, 0x76, 0x06, 0xbd, 0xaa, 0x0e, 0x97, 0x6b, 0x6e,
	0x6d, 0xd8, 0xf2, 0xc1, 0x6e, 0x05, 0xbe, 0x29, 0x30, 0xf6, 0x0c, 0xa0, 0x4e, 0x52, 0x22, 0xbc,
	0x37, 0x0c, 0xc6, 0x07, 0xf3, 0x4e, 0x85, 0xcc, 0xc4, 0xe8, 0x13, 0x3c, 0xf1, 0x0d, 0x8b, 0x46,
	0xc2, 0x3a, 0x54, 0x02, 0xb5, 0xdc, 0xb8, 0x95, 0xc9, 0x30, 0x9f, 0x22, 0x7b, 0x0e, 0x5d, 0x95,
	0xe5, 0x53, 0xe4, 0x42, 0x90, 0xb4, 0xb6, 0x6a, 0x7f, 0x54, 0x60, 0x17, 0x25, 0xc4, 0x1e, 0xc3,
	0x83, 0x9c, 0x12, 0xd4, 0x3c, 0x95, 0x95, 0x80, 0xfb, 0x39, 0x25, 0xef, 0x78, 0x2a, 0xef, 0x24,
	0x8f, 0x6b, 0xf2, 0xf8, 0x37, 0xe4, 0xf1, 0x1f, 0x90, 0x7f, 0x6d, 0xc3, 0xc3, 0x3d, 0x76, 0xf6,
	0x12, 0x8e, 0x85, 0xb4, 0x4e, 0x69, 0xee, 0xb6, 0xae, 0x94, 0xb4, 0x83, 0x46, 0xa0, 0x74, 0xe6,
	0x29, 0x74, 0x94, 0x76, 0x92, 0x12, 0xbe, 0xac, 0xc9, 0xb7, 0x40, 0x61, 0x6e, 0x66, 0xe5, 0xb5,
	0x30, 0x37, 0x8a, 0x64, 0x6d, 0x5d, 0x6f, 0xde, 0xdd, 0x82, 0x33, 0xc1, 0xbe, 0x07, 0x70, 0xd2,
	0xf4, 0x07, 0xb9, 0x16, 0x98, 0x53, 0x12, 0xb6, 0x87, 0xc1, 0xf8, 0xe8, 0xfc, 0x4b, 0xf4, 0xf7,
	0x57, 0x26, 0xba, 0x63, 0x7c, 0xf3, 0xe3, 0xc6, 0xa0, 0x2e, 0xb4, 0xb8, 0xa2, 0xa4, 0x96, 0x1c,
	0xdf, 0x92, 0x7c, 0xf0, 0x1f, 0x48, 0x8e, 0x4b, 0xc9, 0xf1, 0xae, 0xe4, 0x33, 0xe8, 0x29, 0x9d,
	0xf3, 0xb5, 0x12, 0x98, 0xf3, 0xf5, 0xb5, 0x0c, 0x0f, 0xcb, 0x51, 0x54, 0xe0, 0x55, 0x81, 0x8d,
	0x7e, 0xb4, 0x60, 0xb0, 0x7f, 0x3a, 0xac, 0x0f, 0x2d, 0x25, 0xc2, 0x73, 0x9f, 0xde, 0x52, 0xc2,
	0x5f, 0x51, 0xb1, 0x4a, 0x93, 0xea, 0x8a, 0x78, 0x2a, 0xd9, 0x18, 0x06, 0x3b, 0x57, 0x84, 0x1b,
	0x0a, 0xa7, 0x3e, 0xde, 0x6f, 0x1e, 0xd2, 0x07, 0x62, 0x13, 0x38, 0x6d, 0x6e, 0xd7, 0x76, 0x79,
	0x5e, 0xf9, 0xf4, 0x93, 0x46, 0x70, 0xf6, 0x6b, 0x8f, 0xbe, 0x05, 0xd0, 0xdf, 0xa9, 0x12, 0x61,
	0xec, 0x9d, 0xb6, 0xff, 0xc0, 0xe9, 0x79, 0xaf, 0xa9, 0x51, 0xb0, 0x17, 0xd0, 0xcf, 0xd6, 0xdc,
	0x25, 0x86, 0x52, 0x94, 0x44, 0x86, 0xc2, 0xd7, 0xde, 0xab, 0x5e, 0x8d, 0x5e, 0x16, 0xe0, 0xe2,
	0xd0, 0xff, 0x00, 0x27, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x45, 0xcb, 0x62, 0x49, 0x1a, 0x05,
	0x00, 0x00,
}
