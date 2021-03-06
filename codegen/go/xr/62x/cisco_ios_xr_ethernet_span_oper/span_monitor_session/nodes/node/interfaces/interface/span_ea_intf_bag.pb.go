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
// source: span_ea_intf_bag.proto

package cisco_ios_xr_ethernet_span_oper_span_monitor_session_nodes_node_interfaces_interface

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

type SpanEaIntfBag_KEYS struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanEaIntfBag_KEYS) Reset()         { *m = SpanEaIntfBag_KEYS{} }
func (m *SpanEaIntfBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*SpanEaIntfBag_KEYS) ProtoMessage()    {}
func (*SpanEaIntfBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f67db4ce557cb5f, []int{0}
}

func (m *SpanEaIntfBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanEaIntfBag_KEYS.Unmarshal(m, b)
}
func (m *SpanEaIntfBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanEaIntfBag_KEYS.Marshal(b, m, deterministic)
}
func (m *SpanEaIntfBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanEaIntfBag_KEYS.Merge(m, src)
}
func (m *SpanEaIntfBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_SpanEaIntfBag_KEYS.Size(m)
}
func (m *SpanEaIntfBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanEaIntfBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SpanEaIntfBag_KEYS proto.InternalMessageInfo

func (m *SpanEaIntfBag_KEYS) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *SpanEaIntfBag_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
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
	return fileDescriptor_2f67db4ce557cb5f, []int{1}
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
	return fileDescriptor_2f67db4ce557cb5f, []int{2}
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
	return fileDescriptor_2f67db4ce557cb5f, []int{3}
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

type SpanBagSrcInfo_ struct {
	TrafficDirection     string   `protobuf:"bytes,1,opt,name=traffic_direction,json=trafficDirection,proto3" json:"traffic_direction,omitempty"`
	PortLevel            bool     `protobuf:"varint,2,opt,name=port_level,json=portLevel,proto3" json:"port_level,omitempty"`
	IsAclEnabled         bool     `protobuf:"varint,3,opt,name=is_acl_enabled,json=isAclEnabled,proto3" json:"is_acl_enabled,omitempty"`
	MirrorBytes          uint32   `protobuf:"varint,4,opt,name=mirror_bytes,json=mirrorBytes,proto3" json:"mirror_bytes,omitempty"`
	MirrorInterval       string   `protobuf:"bytes,5,opt,name=mirror_interval,json=mirrorInterval,proto3" json:"mirror_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpanBagSrcInfo_) Reset()         { *m = SpanBagSrcInfo_{} }
func (m *SpanBagSrcInfo_) String() string { return proto.CompactTextString(m) }
func (*SpanBagSrcInfo_) ProtoMessage()    {}
func (*SpanBagSrcInfo_) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f67db4ce557cb5f, []int{4}
}

func (m *SpanBagSrcInfo_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanBagSrcInfo_.Unmarshal(m, b)
}
func (m *SpanBagSrcInfo_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanBagSrcInfo_.Marshal(b, m, deterministic)
}
func (m *SpanBagSrcInfo_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanBagSrcInfo_.Merge(m, src)
}
func (m *SpanBagSrcInfo_) XXX_Size() int {
	return xxx_messageInfo_SpanBagSrcInfo_.Size(m)
}
func (m *SpanBagSrcInfo_) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanBagSrcInfo_.DiscardUnknown(m)
}

var xxx_messageInfo_SpanBagSrcInfo_ proto.InternalMessageInfo

func (m *SpanBagSrcInfo_) GetTrafficDirection() string {
	if m != nil {
		return m.TrafficDirection
	}
	return ""
}

func (m *SpanBagSrcInfo_) GetPortLevel() bool {
	if m != nil {
		return m.PortLevel
	}
	return false
}

func (m *SpanBagSrcInfo_) GetIsAclEnabled() bool {
	if m != nil {
		return m.IsAclEnabled
	}
	return false
}

func (m *SpanBagSrcInfo_) GetMirrorBytes() uint32 {
	if m != nil {
		return m.MirrorBytes
	}
	return 0
}

func (m *SpanBagSrcInfo_) GetMirrorInterval() string {
	if m != nil {
		return m.MirrorInterval
	}
	return ""
}

type SpanEaAttachmentBag struct {
	Class                      string           `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	DestinationId              *SpanBagDstId    `protobuf:"bytes,2,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	TrafficMirroringParameters *SpanBagSrcInfo_ `protobuf:"bytes,3,opt,name=traffic_mirroring_parameters,json=trafficMirroringParameters,proto3" json:"traffic_mirroring_parameters,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}         `json:"-"`
	XXX_unrecognized           []byte           `json:"-"`
	XXX_sizecache              int32            `json:"-"`
}

func (m *SpanEaAttachmentBag) Reset()         { *m = SpanEaAttachmentBag{} }
func (m *SpanEaAttachmentBag) String() string { return proto.CompactTextString(m) }
func (*SpanEaAttachmentBag) ProtoMessage()    {}
func (*SpanEaAttachmentBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f67db4ce557cb5f, []int{5}
}

func (m *SpanEaAttachmentBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanEaAttachmentBag.Unmarshal(m, b)
}
func (m *SpanEaAttachmentBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanEaAttachmentBag.Marshal(b, m, deterministic)
}
func (m *SpanEaAttachmentBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanEaAttachmentBag.Merge(m, src)
}
func (m *SpanEaAttachmentBag) XXX_Size() int {
	return xxx_messageInfo_SpanEaAttachmentBag.Size(m)
}
func (m *SpanEaAttachmentBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanEaAttachmentBag.DiscardUnknown(m)
}

var xxx_messageInfo_SpanEaAttachmentBag proto.InternalMessageInfo

func (m *SpanEaAttachmentBag) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *SpanEaAttachmentBag) GetDestinationId() *SpanBagDstId {
	if m != nil {
		return m.DestinationId
	}
	return nil
}

func (m *SpanEaAttachmentBag) GetTrafficMirroringParameters() *SpanBagSrcInfo_ {
	if m != nil {
		return m.TrafficMirroringParameters
	}
	return nil
}

type SpanEaIntfBag struct {
	SourceInterface            string                 `protobuf:"bytes,50,opt,name=source_interface,json=sourceInterface,proto3" json:"source_interface,omitempty"`
	Attachment                 []*SpanEaAttachmentBag `protobuf:"bytes,51,rep,name=attachment,proto3" json:"attachment,omitempty"`
	PlatformError              uint32                 `protobuf:"varint,52,opt,name=platform_error,json=platformError,proto3" json:"platform_error,omitempty"`
	DestinationInterface       string                 `protobuf:"bytes,53,opt,name=destination_interface,json=destinationInterface,proto3" json:"destination_interface,omitempty"`
	TrafficDirection           string                 `protobuf:"bytes,54,opt,name=traffic_direction,json=trafficDirection,proto3" json:"traffic_direction,omitempty"`
	DestinationId              *SpanBagDstId          `protobuf:"bytes,55,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	TrafficMirroringParameters *SpanBagSrcInfo_       `protobuf:"bytes,56,opt,name=traffic_mirroring_parameters,json=trafficMirroringParameters,proto3" json:"traffic_mirroring_parameters,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}               `json:"-"`
	XXX_unrecognized           []byte                 `json:"-"`
	XXX_sizecache              int32                  `json:"-"`
}

func (m *SpanEaIntfBag) Reset()         { *m = SpanEaIntfBag{} }
func (m *SpanEaIntfBag) String() string { return proto.CompactTextString(m) }
func (*SpanEaIntfBag) ProtoMessage()    {}
func (*SpanEaIntfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f67db4ce557cb5f, []int{6}
}

func (m *SpanEaIntfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpanEaIntfBag.Unmarshal(m, b)
}
func (m *SpanEaIntfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpanEaIntfBag.Marshal(b, m, deterministic)
}
func (m *SpanEaIntfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpanEaIntfBag.Merge(m, src)
}
func (m *SpanEaIntfBag) XXX_Size() int {
	return xxx_messageInfo_SpanEaIntfBag.Size(m)
}
func (m *SpanEaIntfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_SpanEaIntfBag.DiscardUnknown(m)
}

var xxx_messageInfo_SpanEaIntfBag proto.InternalMessageInfo

func (m *SpanEaIntfBag) GetSourceInterface() string {
	if m != nil {
		return m.SourceInterface
	}
	return ""
}

func (m *SpanEaIntfBag) GetAttachment() []*SpanEaAttachmentBag {
	if m != nil {
		return m.Attachment
	}
	return nil
}

func (m *SpanEaIntfBag) GetPlatformError() uint32 {
	if m != nil {
		return m.PlatformError
	}
	return 0
}

func (m *SpanEaIntfBag) GetDestinationInterface() string {
	if m != nil {
		return m.DestinationInterface
	}
	return ""
}

func (m *SpanEaIntfBag) GetTrafficDirection() string {
	if m != nil {
		return m.TrafficDirection
	}
	return ""
}

func (m *SpanEaIntfBag) GetDestinationId() *SpanBagDstId {
	if m != nil {
		return m.DestinationId
	}
	return nil
}

func (m *SpanEaIntfBag) GetTrafficMirroringParameters() *SpanBagSrcInfo_ {
	if m != nil {
		return m.TrafficMirroringParameters
	}
	return nil
}

func init() {
	proto.RegisterType((*SpanEaIntfBag_KEYS)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_ea_intf_bag_KEYS")
	proto.RegisterType((*SpanBagDstIdNexthopV4_)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_bag_dst_id_nexthop_v4_")
	proto.RegisterType((*SpanBagDstIdNexthopV6_)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_bag_dst_id_nexthop_v6_")
	proto.RegisterType((*SpanBagDstId)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_bag_dst_id")
	proto.RegisterType((*SpanBagSrcInfo_)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_bag_src_info_")
	proto.RegisterType((*SpanEaAttachmentBag)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_ea_attachment_bag")
	proto.RegisterType((*SpanEaIntfBag)(nil), "cisco_ios_xr_ethernet_span_oper.span_monitor_session.nodes.node.interfaces.interface.span_ea_intf_bag")
}

func init() { proto.RegisterFile("span_ea_intf_bag.proto", fileDescriptor_2f67db4ce557cb5f) }

var fileDescriptor_2f67db4ce557cb5f = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0xcb, 0x6e, 0x13, 0x3f,
	0x14, 0xc6, 0x35, 0xbd, 0xfd, 0x1b, 0xa7, 0x49, 0x5b, 0xab, 0xfd, 0x2b, 0x40, 0x91, 0x42, 0x0a,
	0x22, 0x08, 0x29, 0x8b, 0xb4, 0x04, 0xb6, 0x05, 0xba, 0x88, 0xb8, 0x08, 0x0d, 0xa8, 0x12, 0x62,
	0x71, 0xe4, 0x8c, 0xcf, 0x34, 0x96, 0x26, 0xf6, 0x60, 0xbb, 0x43, 0x79, 0x87, 0x6e, 0x78, 0x05,
	0x96, 0x3c, 0x0e, 0x8f, 0xc1, 0x8a, 0x47, 0x40, 0xf6, 0x64, 0x26, 0xd3, 0x0b, 0x15, 0x9b, 0x42,
	0x37, 0x95, 0xfb, 0xf3, 0xcc, 0xe7, 0x73, 0xce, 0xf7, 0x8d, 0x15, 0xf2, 0xbf, 0x49, 0x99, 0x04,
	0x64, 0x20, 0xa4, 0x8d, 0x61, 0xc4, 0x0e, 0x7b, 0xa9, 0x56, 0x56, 0xd1, 0x77, 0x91, 0x30, 0x91,
	0x02, 0xa1, 0x0c, 0x1c, 0x6b, 0x40, 0x3b, 0x46, 0x2d, 0xd1, 0x82, 0x7f, 0x5a, 0xa5, 0xa8, 0x7b,
	0x7e, 0x35, 0x51, 0x52, 0x58, 0xa5, 0xc1, 0xa0, 0x31, 0x42, 0xc9, 0x9e, 0x54, 0x1c, 0x8d, 0xff,
	0xdb, 0x13, 0xd2, 0xa2, 0x8e, 0x59, 0x84, 0x66, 0xb6, 0xec, 0x0c, 0xc9, 0xe6, 0xd9, 0xf3, 0xe0,
	0xc5, 0xfe, 0xfb, 0xb7, 0x94, 0x92, 0x05, 0xf7, 0x56, 0x2b, 0x68, 0x07, 0xdd, 0x5a, 0xe8, 0xd7,
	0x74, 0x8b, 0xd4, 0xca, 0x37, 0x5b, 0x73, 0x7e, 0x63, 0x06, 0x3a, 0x1f, 0xc8, 0x2d, 0x2f, 0xe5,
	0x24, 0xb8, 0xb1, 0x20, 0x38, 0x48, 0x3c, 0xb6, 0x63, 0x95, 0x42, 0xb6, 0x0b, 0xf4, 0x0e, 0x59,
	0x11, 0x69, 0xb6, 0x0b, 0x8c, 0x73, 0x8d, 0xc6, 0x4c, 0x85, 0xeb, 0x8e, 0xed, 0xe5, 0x88, 0xde,
	0x20, 0xcb, 0x99, 0x8e, 0x41, 0xb2, 0x49, 0x21, 0xff, 0x5f, 0xa6, 0xe3, 0xd7, 0x6c, 0x72, 0xb9,
	0xf8, 0xa0, 0x10, 0x1f, 0x5c, 0x20, 0x3e, 0xf8, 0x03, 0xf1, 0x9f, 0xf3, 0x64, 0xf5, 0x8c, 0x3a,
	0x7d, 0x48, 0xd6, 0x39, 0x1a, 0x2b, 0x24, 0xb3, 0x42, 0x49, 0x88, 0x12, 0x56, 0xca, 0xae, 0x55,
	0x36, 0x9e, 0x39, 0x7e, 0xf9, 0x60, 0xe8, 0x36, 0x69, 0xa4, 0x06, 0x8f, 0xb8, 0xfa, 0x24, 0x34,
	0x82, 0xe0, 0xad, 0xf9, 0x76, 0xd0, 0x6d, 0x84, 0x2b, 0x33, 0x38, 0xe4, 0xf4, 0x6b, 0x40, 0x36,
	0xaa, 0xf3, 0x01, 0x26, 0x39, 0x64, 0x3a, 0x6e, 0x2d, 0xb4, 0x83, 0x6e, 0xbd, 0xff, 0xb1, 0x77,
	0x15, 0xf6, 0xf7, 0x2e, 0x31, 0x2c, 0x5c, 0xaf, 0x58, 0xb3, 0x27, 0xf9, 0x81, 0x8e, 0x8b, 0x22,
	0x07, 0xe7, 0x8a, 0x5c, 0xfc, 0x27, 0x45, 0x0e, 0xf2, 0x22, 0x07, 0xa7, 0x8b, 0xdc, 0x26, 0x0d,
	0x21, 0x33, 0x96, 0x08, 0x0e, 0x19, 0x4b, 0x8e, 0xb0, 0xb5, 0x94, 0x8f, 0x7b, 0x0a, 0x0f, 0x1c,
	0xeb, 0x7c, 0x0f, 0x08, 0x2d, 0x75, 0x8d, 0x8e, 0x40, 0xc8, 0x58, 0x81, 0x73, 0xdd, 0x6a, 0x16,
	0xc7, 0x22, 0x02, 0x2e, 0x34, 0x46, 0xce, 0xe2, 0xc2, 0xf5, 0xe9, 0xc6, 0xf3, 0x82, 0xd3, 0xdb,
	0x84, 0xa4, 0x4a, 0x5b, 0x48, 0x30, 0xc3, 0xc4, 0xdb, 0xbe, 0x1c, 0xd6, 0x1c, 0x79, 0xe9, 0x00,
	0xbd, 0x4b, 0x9a, 0xc2, 0x00, 0x8b, 0x12, 0x40, 0xc9, 0x46, 0x09, 0xe6, 0xbe, 0x2f, 0x87, 0x2b,
	0xc2, 0xec, 0x45, 0xc9, 0x7e, 0xce, 0x5c, 0x72, 0x27, 0x42, 0x6b, 0xa5, 0x61, 0xf4, 0xd9, 0xa2,
	0xf1, 0x76, 0x37, 0xc2, 0x7a, 0xce, 0x9e, 0x3a, 0x44, 0xef, 0x93, 0xd5, 0xe9, 0x23, 0x7e, 0x26,
	0x19, 0x4b, 0xfc, 0xbc, 0x6b, 0x61, 0x33, 0xc7, 0xc3, 0x29, 0xed, 0xfc, 0x98, 0x9b, 0xdd, 0x1e,
	0xcc, 0x5a, 0x16, 0x8d, 0x27, 0x28, 0xad, 0x6b, 0x91, 0x6e, 0x90, 0xc5, 0x6a, 0x84, 0xf3, 0x7f,
	0xe8, 0x49, 0x40, 0x9a, 0xd5, 0x94, 0x0b, 0xee, 0xdb, 0xa8, 0xf7, 0xf1, 0xaf, 0x38, 0x19, 0x36,
	0x2a, 0x87, 0x0f, 0x39, 0xfd, 0x16, 0x90, 0xad, 0x62, 0xfc, 0x79, 0x6b, 0x42, 0x1e, 0x42, 0xca,
	0x34, 0x9b, 0xa0, 0x45, 0x6d, 0xfc, 0x00, 0xeb, 0xfd, 0xf1, 0x15, 0x17, 0x57, 0xc6, 0x21, 0xbc,
	0x39, 0xad, 0xe6, 0x55, 0x51, 0xcc, 0x9b, 0xb2, 0x96, 0xce, 0x97, 0x45, 0xb2, 0x76, 0xf6, 0xea,
	0xa4, 0x0f, 0xc8, 0x9a, 0x51, 0x47, 0x3a, 0x42, 0x98, 0xdd, 0x07, 0x7d, 0x3f, 0xf1, 0xd5, 0x9c,
	0x0f, 0xcb, 0x5b, 0xe1, 0x24, 0x20, 0x64, 0x66, 0x52, 0x6b, 0xa7, 0x3d, 0xdf, 0xad, 0xf7, 0x93,
	0x2b, 0x6c, 0xed, 0x5c, 0x28, 0xc2, 0xca, 0xf9, 0xf4, 0x1e, 0x69, 0xa6, 0x09, 0xb3, 0xb1, 0xd2,
	0x13, 0x40, 0xd7, 0x6d, 0x6b, 0xd7, 0x27, 0xb1, 0x51, 0xd0, 0x7d, 0x07, 0xe9, 0x0e, 0xd9, 0x3c,
	0x15, 0x98, 0xb2, 0xcb, 0x47, 0xbe, 0xcb, 0x8d, 0xaa, 0xa1, 0x65, 0xab, 0x17, 0x7e, 0x55, 0x83,
	0xdf, 0x7c, 0x55, 0x17, 0x64, 0xf2, 0xf1, 0x75, 0xce, 0xe4, 0x93, 0xeb, 0x93, 0xc9, 0xd1, 0x92,
	0xff, 0xa9, 0xb0, 0xf3, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x9d, 0x66, 0x41, 0x44, 0x08, 0x00,
	0x00,
}
