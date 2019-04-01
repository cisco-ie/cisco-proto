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
// source: l2vpn_disco.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_nodes_node_discoveries_discovery

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

type L2VpnDisco_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ServiceName          string   `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	GroupName            string   `protobuf:"bytes,3,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	VpnName              string   `protobuf:"bytes,4,opt,name=vpn_name,json=vpnName,proto3" json:"vpn_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnDisco_KEYS) Reset()         { *m = L2VpnDisco_KEYS{} }
func (m *L2VpnDisco_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnDisco_KEYS) ProtoMessage()    {}
func (*L2VpnDisco_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{0}
}

func (m *L2VpnDisco_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDisco_KEYS.Unmarshal(m, b)
}
func (m *L2VpnDisco_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDisco_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnDisco_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDisco_KEYS.Merge(m, src)
}
func (m *L2VpnDisco_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnDisco_KEYS.Size(m)
}
func (m *L2VpnDisco_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDisco_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDisco_KEYS proto.InternalMessageInfo

func (m *L2VpnDisco_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnDisco_KEYS) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *L2VpnDisco_KEYS) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *L2VpnDisco_KEYS) GetVpnName() string {
	if m != nil {
		return m.VpnName
	}
	return ""
}

type L2VpnLabelBlock struct {
	LabelTimeCreated     uint32   `protobuf:"varint,1,opt,name=label_time_created,json=labelTimeCreated,proto3" json:"label_time_created,omitempty"`
	LabelBase            uint32   `protobuf:"varint,2,opt,name=label_base,json=labelBase,proto3" json:"label_base,omitempty"`
	BlockOffset          uint32   `protobuf:"varint,3,opt,name=block_offset,json=blockOffset,proto3" json:"block_offset,omitempty"`
	BlockSize            uint32   `protobuf:"varint,4,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	LocalEdgeId          uint32   `protobuf:"varint,5,opt,name=local_edge_id,json=localEdgeId,proto3" json:"local_edge_id,omitempty"`
	NextHop              string   `protobuf:"bytes,6,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	StatusVector         []uint32 `protobuf:"varint,7,rep,packed,name=status_vector,json=statusVector,proto3" json:"status_vector,omitempty"`
	LabelError           string   `protobuf:"bytes,8,opt,name=label_error,json=labelError,proto3" json:"label_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLabelBlock) Reset()         { *m = L2VpnLabelBlock{} }
func (m *L2VpnLabelBlock) String() string { return proto.CompactTextString(m) }
func (*L2VpnLabelBlock) ProtoMessage()    {}
func (*L2VpnLabelBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{1}
}

func (m *L2VpnLabelBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLabelBlock.Unmarshal(m, b)
}
func (m *L2VpnLabelBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLabelBlock.Marshal(b, m, deterministic)
}
func (m *L2VpnLabelBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLabelBlock.Merge(m, src)
}
func (m *L2VpnLabelBlock) XXX_Size() int {
	return xxx_messageInfo_L2VpnLabelBlock.Size(m)
}
func (m *L2VpnLabelBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLabelBlock.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLabelBlock proto.InternalMessageInfo

func (m *L2VpnLabelBlock) GetLabelTimeCreated() uint32 {
	if m != nil {
		return m.LabelTimeCreated
	}
	return 0
}

func (m *L2VpnLabelBlock) GetLabelBase() uint32 {
	if m != nil {
		return m.LabelBase
	}
	return 0
}

func (m *L2VpnLabelBlock) GetBlockOffset() uint32 {
	if m != nil {
		return m.BlockOffset
	}
	return 0
}

func (m *L2VpnLabelBlock) GetBlockSize() uint32 {
	if m != nil {
		return m.BlockSize
	}
	return 0
}

func (m *L2VpnLabelBlock) GetLocalEdgeId() uint32 {
	if m != nil {
		return m.LocalEdgeId
	}
	return 0
}

func (m *L2VpnLabelBlock) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnLabelBlock) GetStatusVector() []uint32 {
	if m != nil {
		return m.StatusVector
	}
	return nil
}

func (m *L2VpnLabelBlock) GetLabelError() string {
	if m != nil {
		return m.LabelError
	}
	return ""
}

type L2VpnDiscoEdge struct {
	EdgeId               uint32             `protobuf:"varint,1,opt,name=edge_id,json=edgeId,proto3" json:"edge_id,omitempty"`
	LabelCount           uint32             `protobuf:"varint,2,opt,name=label_count,json=labelCount,proto3" json:"label_count,omitempty"`
	LabelBlock           []*L2VpnLabelBlock `protobuf:"bytes,3,rep,name=label_block,json=labelBlock,proto3" json:"label_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnDiscoEdge) Reset()         { *m = L2VpnDiscoEdge{} }
func (m *L2VpnDiscoEdge) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoEdge) ProtoMessage()    {}
func (*L2VpnDiscoEdge) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{2}
}

func (m *L2VpnDiscoEdge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoEdge.Unmarshal(m, b)
}
func (m *L2VpnDiscoEdge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoEdge.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoEdge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoEdge.Merge(m, src)
}
func (m *L2VpnDiscoEdge) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoEdge.Size(m)
}
func (m *L2VpnDiscoEdge) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoEdge.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoEdge proto.InternalMessageInfo

func (m *L2VpnDiscoEdge) GetEdgeId() uint32 {
	if m != nil {
		return m.EdgeId
	}
	return 0
}

func (m *L2VpnDiscoEdge) GetLabelCount() uint32 {
	if m != nil {
		return m.LabelCount
	}
	return 0
}

func (m *L2VpnDiscoEdge) GetLabelBlock() []*L2VpnLabelBlock {
	if m != nil {
		return m.LabelBlock
	}
	return nil
}

type L2VpnDiscoBgpSig struct {
	NumberEdges          uint32            `protobuf:"varint,1,opt,name=number_edges,json=numberEdges,proto3" json:"number_edges,omitempty"`
	NumberRemoteEdges    uint32            `protobuf:"varint,2,opt,name=number_remote_edges,json=numberRemoteEdges,proto3" json:"number_remote_edges,omitempty"`
	Edge                 []*L2VpnDiscoEdge `protobuf:"bytes,3,rep,name=edge,proto3" json:"edge,omitempty"`
	Redge                []*L2VpnDiscoEdge `protobuf:"bytes,4,rep,name=redge,proto3" json:"redge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2VpnDiscoBgpSig) Reset()         { *m = L2VpnDiscoBgpSig{} }
func (m *L2VpnDiscoBgpSig) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoBgpSig) ProtoMessage()    {}
func (*L2VpnDiscoBgpSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{3}
}

func (m *L2VpnDiscoBgpSig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoBgpSig.Unmarshal(m, b)
}
func (m *L2VpnDiscoBgpSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoBgpSig.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoBgpSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoBgpSig.Merge(m, src)
}
func (m *L2VpnDiscoBgpSig) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoBgpSig.Size(m)
}
func (m *L2VpnDiscoBgpSig) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoBgpSig.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoBgpSig proto.InternalMessageInfo

func (m *L2VpnDiscoBgpSig) GetNumberEdges() uint32 {
	if m != nil {
		return m.NumberEdges
	}
	return 0
}

func (m *L2VpnDiscoBgpSig) GetNumberRemoteEdges() uint32 {
	if m != nil {
		return m.NumberRemoteEdges
	}
	return 0
}

func (m *L2VpnDiscoBgpSig) GetEdge() []*L2VpnDiscoEdge {
	if m != nil {
		return m.Edge
	}
	return nil
}

func (m *L2VpnDiscoBgpSig) GetRedge() []*L2VpnDiscoEdge {
	if m != nil {
		return m.Redge
	}
	return nil
}

type L2VpnVplsIdAuto struct {
	Asn                  uint32   `protobuf:"varint,1,opt,name=asn,proto3" json:"asn,omitempty"`
	VpnId                uint32   `protobuf:"varint,2,opt,name=vpn_id,json=vpnId,proto3" json:"vpn_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnVplsIdAuto) Reset()         { *m = L2VpnVplsIdAuto{} }
func (m *L2VpnVplsIdAuto) String() string { return proto.CompactTextString(m) }
func (*L2VpnVplsIdAuto) ProtoMessage()    {}
func (*L2VpnVplsIdAuto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{4}
}

func (m *L2VpnVplsIdAuto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnVplsIdAuto.Unmarshal(m, b)
}
func (m *L2VpnVplsIdAuto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnVplsIdAuto.Marshal(b, m, deterministic)
}
func (m *L2VpnVplsIdAuto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnVplsIdAuto.Merge(m, src)
}
func (m *L2VpnVplsIdAuto) XXX_Size() int {
	return xxx_messageInfo_L2VpnVplsIdAuto.Size(m)
}
func (m *L2VpnVplsIdAuto) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnVplsIdAuto.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnVplsIdAuto proto.InternalMessageInfo

func (m *L2VpnVplsIdAuto) GetAsn() uint32 {
	if m != nil {
		return m.Asn
	}
	return 0
}

func (m *L2VpnVplsIdAuto) GetVpnId() uint32 {
	if m != nil {
		return m.VpnId
	}
	return 0
}

type L2VpnVplsId_2ByteAs struct {
	TwoByteAs            uint32   `protobuf:"varint,1,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	FourByteIndex        uint32   `protobuf:"varint,2,opt,name=four_byte_index,json=fourByteIndex,proto3" json:"four_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnVplsId_2ByteAs) Reset()         { *m = L2VpnVplsId_2ByteAs{} }
func (m *L2VpnVplsId_2ByteAs) String() string { return proto.CompactTextString(m) }
func (*L2VpnVplsId_2ByteAs) ProtoMessage()    {}
func (*L2VpnVplsId_2ByteAs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{5}
}

func (m *L2VpnVplsId_2ByteAs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnVplsId_2ByteAs.Unmarshal(m, b)
}
func (m *L2VpnVplsId_2ByteAs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnVplsId_2ByteAs.Marshal(b, m, deterministic)
}
func (m *L2VpnVplsId_2ByteAs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnVplsId_2ByteAs.Merge(m, src)
}
func (m *L2VpnVplsId_2ByteAs) XXX_Size() int {
	return xxx_messageInfo_L2VpnVplsId_2ByteAs.Size(m)
}
func (m *L2VpnVplsId_2ByteAs) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnVplsId_2ByteAs.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnVplsId_2ByteAs proto.InternalMessageInfo

func (m *L2VpnVplsId_2ByteAs) GetTwoByteAs() uint32 {
	if m != nil {
		return m.TwoByteAs
	}
	return 0
}

func (m *L2VpnVplsId_2ByteAs) GetFourByteIndex() uint32 {
	if m != nil {
		return m.FourByteIndex
	}
	return 0
}

type L2VpnVplsIdV4Addr struct {
	Ipv4Address          string   `protobuf:"bytes,1,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	TwoByteIndex         uint32   `protobuf:"varint,2,opt,name=two_byte_index,json=twoByteIndex,proto3" json:"two_byte_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnVplsIdV4Addr) Reset()         { *m = L2VpnVplsIdV4Addr{} }
func (m *L2VpnVplsIdV4Addr) String() string { return proto.CompactTextString(m) }
func (*L2VpnVplsIdV4Addr) ProtoMessage()    {}
func (*L2VpnVplsIdV4Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{6}
}

func (m *L2VpnVplsIdV4Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnVplsIdV4Addr.Unmarshal(m, b)
}
func (m *L2VpnVplsIdV4Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnVplsIdV4Addr.Marshal(b, m, deterministic)
}
func (m *L2VpnVplsIdV4Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnVplsIdV4Addr.Merge(m, src)
}
func (m *L2VpnVplsIdV4Addr) XXX_Size() int {
	return xxx_messageInfo_L2VpnVplsIdV4Addr.Size(m)
}
func (m *L2VpnVplsIdV4Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnVplsIdV4Addr.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnVplsIdV4Addr proto.InternalMessageInfo

func (m *L2VpnVplsIdV4Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *L2VpnVplsIdV4Addr) GetTwoByteIndex() uint32 {
	if m != nil {
		return m.TwoByteIndex
	}
	return 0
}

type L2VpnVplsId struct {
	VplsIdType           string               `protobuf:"bytes,1,opt,name=vpls_id_type,json=vplsIdType,proto3" json:"vpls_id_type,omitempty"`
	Auto                 *L2VpnVplsIdAuto     `protobuf:"bytes,2,opt,name=auto,proto3" json:"auto,omitempty"`
	TwoByteAs            *L2VpnVplsId_2ByteAs `protobuf:"bytes,3,opt,name=two_byte_as,json=twoByteAs,proto3" json:"two_byte_as,omitempty"`
	V4Addr               *L2VpnVplsIdV4Addr   `protobuf:"bytes,4,opt,name=v4_addr,json=v4Addr,proto3" json:"v4_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *L2VpnVplsId) Reset()         { *m = L2VpnVplsId{} }
func (m *L2VpnVplsId) String() string { return proto.CompactTextString(m) }
func (*L2VpnVplsId) ProtoMessage()    {}
func (*L2VpnVplsId) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{7}
}

func (m *L2VpnVplsId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnVplsId.Unmarshal(m, b)
}
func (m *L2VpnVplsId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnVplsId.Marshal(b, m, deterministic)
}
func (m *L2VpnVplsId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnVplsId.Merge(m, src)
}
func (m *L2VpnVplsId) XXX_Size() int {
	return xxx_messageInfo_L2VpnVplsId.Size(m)
}
func (m *L2VpnVplsId) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnVplsId.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnVplsId proto.InternalMessageInfo

func (m *L2VpnVplsId) GetVplsIdType() string {
	if m != nil {
		return m.VplsIdType
	}
	return ""
}

func (m *L2VpnVplsId) GetAuto() *L2VpnVplsIdAuto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *L2VpnVplsId) GetTwoByteAs() *L2VpnVplsId_2ByteAs {
	if m != nil {
		return m.TwoByteAs
	}
	return nil
}

func (m *L2VpnVplsId) GetV4Addr() *L2VpnVplsIdV4Addr {
	if m != nil {
		return m.V4Addr
	}
	return nil
}

type L2VpnNlriLdpSig struct {
	NlriTimeCreated      uint32   `protobuf:"varint,1,opt,name=nlri_time_created,json=nlriTimeCreated,proto3" json:"nlri_time_created,omitempty"`
	LocalAddress         string   `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        string   `protobuf:"bytes,3,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	RemoteL2RouterId     string   `protobuf:"bytes,4,opt,name=remote_l2_router_id,json=remoteL2RouterId,proto3" json:"remote_l2_router_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnNlriLdpSig) Reset()         { *m = L2VpnNlriLdpSig{} }
func (m *L2VpnNlriLdpSig) String() string { return proto.CompactTextString(m) }
func (*L2VpnNlriLdpSig) ProtoMessage()    {}
func (*L2VpnNlriLdpSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{8}
}

func (m *L2VpnNlriLdpSig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnNlriLdpSig.Unmarshal(m, b)
}
func (m *L2VpnNlriLdpSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnNlriLdpSig.Marshal(b, m, deterministic)
}
func (m *L2VpnNlriLdpSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnNlriLdpSig.Merge(m, src)
}
func (m *L2VpnNlriLdpSig) XXX_Size() int {
	return xxx_messageInfo_L2VpnNlriLdpSig.Size(m)
}
func (m *L2VpnNlriLdpSig) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnNlriLdpSig.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnNlriLdpSig proto.InternalMessageInfo

func (m *L2VpnNlriLdpSig) GetNlriTimeCreated() uint32 {
	if m != nil {
		return m.NlriTimeCreated
	}
	return 0
}

func (m *L2VpnNlriLdpSig) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *L2VpnNlriLdpSig) GetRemoteAddress() string {
	if m != nil {
		return m.RemoteAddress
	}
	return ""
}

func (m *L2VpnNlriLdpSig) GetRemoteL2RouterId() string {
	if m != nil {
		return m.RemoteL2RouterId
	}
	return ""
}

type L2VpnDiscoLdpSig struct {
	LocalVplsId          *L2VpnVplsId       `protobuf:"bytes,1,opt,name=local_vpls_id,json=localVplsId,proto3" json:"local_vpls_id,omitempty"`
	LocalL2RouterId      string             `protobuf:"bytes,2,opt,name=local_l2_router_id,json=localL2RouterId,proto3" json:"local_l2_router_id,omitempty"`
	NumberRemoteEdges    uint32             `protobuf:"varint,3,opt,name=number_remote_edges,json=numberRemoteEdges,proto3" json:"number_remote_edges,omitempty"`
	RemoteNlri           []*L2VpnNlriLdpSig `protobuf:"bytes,4,rep,name=remote_nlri,json=remoteNlri,proto3" json:"remote_nlri,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnDiscoLdpSig) Reset()         { *m = L2VpnDiscoLdpSig{} }
func (m *L2VpnDiscoLdpSig) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoLdpSig) ProtoMessage()    {}
func (*L2VpnDiscoLdpSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{9}
}

func (m *L2VpnDiscoLdpSig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoLdpSig.Unmarshal(m, b)
}
func (m *L2VpnDiscoLdpSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoLdpSig.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoLdpSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoLdpSig.Merge(m, src)
}
func (m *L2VpnDiscoLdpSig) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoLdpSig.Size(m)
}
func (m *L2VpnDiscoLdpSig) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoLdpSig.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoLdpSig proto.InternalMessageInfo

func (m *L2VpnDiscoLdpSig) GetLocalVplsId() *L2VpnVplsId {
	if m != nil {
		return m.LocalVplsId
	}
	return nil
}

func (m *L2VpnDiscoLdpSig) GetLocalL2RouterId() string {
	if m != nil {
		return m.LocalL2RouterId
	}
	return ""
}

func (m *L2VpnDiscoLdpSig) GetNumberRemoteEdges() uint32 {
	if m != nil {
		return m.NumberRemoteEdges
	}
	return 0
}

func (m *L2VpnDiscoLdpSig) GetRemoteNlri() []*L2VpnNlriLdpSig {
	if m != nil {
		return m.RemoteNlri
	}
	return nil
}

type L2VpnDiscoSigInfo struct {
	AdSignallingMethod   string            `protobuf:"bytes,1,opt,name=ad_signalling_method,json=adSignallingMethod,proto3" json:"ad_signalling_method,omitempty"`
	BgpSigInfo           *L2VpnDiscoBgpSig `protobuf:"bytes,2,opt,name=bgp_sig_info,json=bgpSigInfo,proto3" json:"bgp_sig_info,omitempty"`
	LdpSigInfo           *L2VpnDiscoLdpSig `protobuf:"bytes,3,opt,name=ldp_sig_info,json=ldpSigInfo,proto3" json:"ldp_sig_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2VpnDiscoSigInfo) Reset()         { *m = L2VpnDiscoSigInfo{} }
func (m *L2VpnDiscoSigInfo) String() string { return proto.CompactTextString(m) }
func (*L2VpnDiscoSigInfo) ProtoMessage()    {}
func (*L2VpnDiscoSigInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{10}
}

func (m *L2VpnDiscoSigInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDiscoSigInfo.Unmarshal(m, b)
}
func (m *L2VpnDiscoSigInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDiscoSigInfo.Marshal(b, m, deterministic)
}
func (m *L2VpnDiscoSigInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDiscoSigInfo.Merge(m, src)
}
func (m *L2VpnDiscoSigInfo) XXX_Size() int {
	return xxx_messageInfo_L2VpnDiscoSigInfo.Size(m)
}
func (m *L2VpnDiscoSigInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDiscoSigInfo.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDiscoSigInfo proto.InternalMessageInfo

func (m *L2VpnDiscoSigInfo) GetAdSignallingMethod() string {
	if m != nil {
		return m.AdSignallingMethod
	}
	return ""
}

func (m *L2VpnDiscoSigInfo) GetBgpSigInfo() *L2VpnDiscoBgpSig {
	if m != nil {
		return m.BgpSigInfo
	}
	return nil
}

func (m *L2VpnDiscoSigInfo) GetLdpSigInfo() *L2VpnDiscoLdpSig {
	if m != nil {
		return m.LdpSigInfo
	}
	return nil
}

type L2VpnDisco struct {
	MtuMismatchIgnore    bool               `protobuf:"varint,50,opt,name=mtu_mismatch_ignore,json=mtuMismatchIgnore,proto3" json:"mtu_mismatch_ignore,omitempty"`
	NumberVpn            uint32             `protobuf:"varint,51,opt,name=number_vpn,json=numberVpn,proto3" json:"number_vpn,omitempty"`
	VpnId                uint32             `protobuf:"varint,52,opt,name=vpn_id,json=vpnId,proto3" json:"vpn_id,omitempty"`
	ServiceNameXr        string             `protobuf:"bytes,53,opt,name=service_name_xr,json=serviceNameXr,proto3" json:"service_name_xr,omitempty"`
	GroupNameXr          string             `protobuf:"bytes,54,opt,name=group_name_xr,json=groupNameXr,proto3" json:"group_name_xr,omitempty"`
	VpnNameXr            string             `protobuf:"bytes,55,opt,name=vpn_name_xr,json=vpnNameXr,proto3" json:"vpn_name_xr,omitempty"`
	IsServiceConnected   bool               `protobuf:"varint,56,opt,name=is_service_connected,json=isServiceConnected,proto3" json:"is_service_connected,omitempty"`
	SignallingInfo       *L2VpnDiscoSigInfo `protobuf:"bytes,57,opt,name=signalling_info,json=signallingInfo,proto3" json:"signalling_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *L2VpnDisco) Reset()         { *m = L2VpnDisco{} }
func (m *L2VpnDisco) String() string { return proto.CompactTextString(m) }
func (*L2VpnDisco) ProtoMessage()    {}
func (*L2VpnDisco) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9a95e650ec6a2d7, []int{11}
}

func (m *L2VpnDisco) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnDisco.Unmarshal(m, b)
}
func (m *L2VpnDisco) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnDisco.Marshal(b, m, deterministic)
}
func (m *L2VpnDisco) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnDisco.Merge(m, src)
}
func (m *L2VpnDisco) XXX_Size() int {
	return xxx_messageInfo_L2VpnDisco.Size(m)
}
func (m *L2VpnDisco) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnDisco.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnDisco proto.InternalMessageInfo

func (m *L2VpnDisco) GetMtuMismatchIgnore() bool {
	if m != nil {
		return m.MtuMismatchIgnore
	}
	return false
}

func (m *L2VpnDisco) GetNumberVpn() uint32 {
	if m != nil {
		return m.NumberVpn
	}
	return 0
}

func (m *L2VpnDisco) GetVpnId() uint32 {
	if m != nil {
		return m.VpnId
	}
	return 0
}

func (m *L2VpnDisco) GetServiceNameXr() string {
	if m != nil {
		return m.ServiceNameXr
	}
	return ""
}

func (m *L2VpnDisco) GetGroupNameXr() string {
	if m != nil {
		return m.GroupNameXr
	}
	return ""
}

func (m *L2VpnDisco) GetVpnNameXr() string {
	if m != nil {
		return m.VpnNameXr
	}
	return ""
}

func (m *L2VpnDisco) GetIsServiceConnected() bool {
	if m != nil {
		return m.IsServiceConnected
	}
	return false
}

func (m *L2VpnDisco) GetSignallingInfo() *L2VpnDiscoSigInfo {
	if m != nil {
		return m.SignallingInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnDisco_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco_KEYS")
	proto.RegisterType((*L2VpnLabelBlock)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_label_block")
	proto.RegisterType((*L2VpnDiscoEdge)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco_edge")
	proto.RegisterType((*L2VpnDiscoBgpSig)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco_bgp_sig")
	proto.RegisterType((*L2VpnVplsIdAuto)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_vpls_id_auto")
	proto.RegisterType((*L2VpnVplsId_2ByteAs)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_vpls_id_2byte_as")
	proto.RegisterType((*L2VpnVplsIdV4Addr)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_vpls_id_v4addr")
	proto.RegisterType((*L2VpnVplsId)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_vpls_id")
	proto.RegisterType((*L2VpnNlriLdpSig)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_nlri_ldp_sig")
	proto.RegisterType((*L2VpnDiscoLdpSig)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco_ldp_sig")
	proto.RegisterType((*L2VpnDiscoSigInfo)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco_sig_info")
	proto.RegisterType((*L2VpnDisco)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.nodes.node.discoveries.discovery.l2vpn_disco")
}

func init() { proto.RegisterFile("l2vpn_disco.proto", fileDescriptor_a9a95e650ec6a2d7) }

var fileDescriptor_a9a95e650ec6a2d7 = []byte{
	// 1051 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x56, 0xe2, 0xc4, 0x49, 0x7a, 0xec, 0xfc, 0x74, 0x02, 0x78, 0x0f, 0x40, 0x98, 0x05, 0x14,
	0xf1, 0x63, 0x21, 0x6f, 0xf8, 0x3b, 0x20, 0xb1, 0x89, 0x22, 0x61, 0xc1, 0xee, 0x4a, 0xe3, 0x10,
	0x85, 0x53, 0x33, 0x9e, 0xe9, 0x38, 0x2d, 0x66, 0xa6, 0x47, 0xdd, 0x3d, 0x93, 0x64, 0x2f, 0xdc,
	0x10, 0x67, 0x5e, 0x82, 0x77, 0x80, 0x1b, 0x0f, 0xc2, 0xb3, 0xa0, 0xaa, 0xea, 0x71, 0xc6, 0x0b,
	0x7b, 0xc2, 0x7b, 0xb1, 0xdc, 0x5f, 0xd5, 0xd4, 0x57, 0x7f, 0x5d, 0x5d, 0x6c, 0x2f, 0x1b, 0xd5,
	0x65, 0x21, 0x52, 0x65, 0x13, 0x3d, 0x2c, 0x8d, 0x76, 0x9a, 0x7f, 0x9d, 0xc0, 0x41, 0x28, 0x6d,
	0xc5, 0xad, 0x11, 0x24, 0xd7, 0xa5, 0x34, 0x43, 0xfc, 0x5b, 0x8f, 0x86, 0x85, 0x4e, 0xa5, 0xc5,
	0xdf, 0x21, 0x7e, 0x57, 0x4b, 0xa3, 0xa4, 0x9d, 0xff, 0xbf, 0x0b, 0x7f, 0x5d, 0x61, 0xbb, 0x2d,
	0xbb, 0xe2, 0xdb, 0xb3, 0x1f, 0x26, 0xfc, 0x0d, 0xb6, 0x01, 0x9f, 0x08, 0x95, 0x0e, 0x56, 0x0e,
	0x57, 0x8e, 0xb6, 0xa2, 0x2e, 0x1c, 0xc7, 0x29, 0x7f, 0x87, 0xf5, 0xac, 0x34, 0xb5, 0x4a, 0xa4,
	0x28, 0xe2, 0x5c, 0x0e, 0x56, 0x51, 0x1a, 0x78, 0xec, 0x69, 0x9c, 0x4b, 0xfe, 0x26, 0x63, 0x33,
	0xa3, 0xab, 0x92, 0x14, 0x3a, 0xa8, 0xb0, 0x85, 0x08, 0x8a, 0x1f, 0xb0, 0x4d, 0x20, 0x43, 0xe1,
	0x1a, 0x0a, 0x37, 0xea, 0xb2, 0x00, 0x51, 0xf8, 0xfb, 0x6a, 0x13, 0x62, 0x16, 0x4f, 0x65, 0x26,
	0xa6, 0x99, 0x4e, 0x7e, 0xe2, 0x1f, 0x31, 0x4e, 0x47, 0xa7, 0x72, 0x29, 0x12, 0x23, 0x63, 0x27,
	0xc9, 0xad, 0x7e, 0xb4, 0x8b, 0x92, 0x73, 0x95, 0xcb, 0x53, 0xc2, 0x81, 0xdd, 0x7f, 0x1c, 0x5b,
	0x72, 0xaf, 0x1f, 0x6d, 0x21, 0x72, 0x12, 0x5b, 0x09, 0xfe, 0xa3, 0x55, 0xa1, 0xaf, 0xae, 0xac,
	0x74, 0xe8, 0x5e, 0x3f, 0x0a, 0x10, 0x7b, 0x86, 0x10, 0x58, 0x20, 0x15, 0xab, 0x9e, 0x93, 0x8b,
	0xfd, 0x68, 0x0b, 0x91, 0x89, 0x7a, 0x2e, 0x79, 0xc8, 0xfa, 0x99, 0x4e, 0xe2, 0x4c, 0xc8, 0x74,
	0x86, 0x09, 0x5a, 0x27, 0x13, 0x08, 0x9e, 0xa5, 0x33, 0xc8, 0xd2, 0x03, 0xb6, 0x59, 0xc8, 0x5b,
	0x27, 0xae, 0x75, 0x39, 0xe8, 0x52, 0x8c, 0x70, 0xfe, 0x46, 0x97, 0xfc, 0x21, 0xeb, 0x5b, 0x17,
	0xbb, 0xca, 0x8a, 0x5a, 0x26, 0x4e, 0x9b, 0xc1, 0xc6, 0x61, 0xe7, 0xa8, 0x1f, 0xf5, 0x08, 0xbc,
	0x40, 0x8c, 0xbf, 0xcd, 0x02, 0x0a, 0x42, 0x1a, 0xa3, 0xcd, 0x60, 0x13, 0x4d, 0x50, 0x5c, 0x67,
	0x80, 0x84, 0x7f, 0xbd, 0x50, 0x34, 0xf0, 0x05, 0x8a, 0xd6, 0xf8, 0x44, 0xd9, 0xe9, 0x4a, 0x72,
	0x67, 0x6e, 0x2e, 0xd1, 0x55, 0xe1, 0x7c, 0x52, 0xc8, 0xdc, 0x29, 0x20, 0xdc, 0x35, 0x0a, 0x18,
	0xe6, 0xa0, 0x73, 0xd8, 0x39, 0x0a, 0x46, 0x93, 0xe1, 0xff, 0xed, 0xad, 0xe1, 0xbf, 0x8a, 0xe9,
	0x59, 0x4f, 0xe0, 0x7f, 0xf8, 0xc7, 0x2a, 0xdb, 0x6f, 0x07, 0x31, 0x9d, 0x95, 0xc2, 0xaa, 0x19,
	0xd4, 0xa8, 0xa8, 0xf2, 0xa9, 0x34, 0x18, 0x96, 0xf5, 0xc1, 0x04, 0x84, 0x41, 0x86, 0x2d, 0x1f,
	0xb2, 0x7d, 0xaf, 0x62, 0x64, 0xae, 0x9d, 0xf4, 0x9a, 0x14, 0xd9, 0x1e, 0x89, 0x22, 0x94, 0x90,
	0xfe, 0x15, 0x5b, 0x03, 0x0d, 0x1f, 0x59, 0xb4, 0xac, 0xc8, 0xee, 0x93, 0x1f, 0xa1, 0x7d, 0x7e,
	0xcd, 0xd6, 0x0d, 0x12, 0xad, 0xbd, 0x32, 0x22, 0x22, 0x08, 0xbf, 0x62, 0x9c, 0x44, 0x75, 0x99,
	0x59, 0xa1, 0x52, 0x11, 0x57, 0x4e, 0xf3, 0x5d, 0xd6, 0x89, 0x6d, 0xe1, 0x33, 0x06, 0x7f, 0xf9,
	0x6b, 0xac, 0x0b, 0x5a, 0x2a, 0xf5, 0xc9, 0x59, 0xaf, 0xcb, 0x62, 0x9c, 0x86, 0x3f, 0xb2, 0xd7,
	0x17, 0x3f, 0x1f, 0x4d, 0xef, 0x9c, 0x14, 0xb1, 0xe5, 0x6f, 0xb1, 0xc0, 0xdd, 0x68, 0xe1, 0x8f,
	0xde, 0xd4, 0x96, 0xbb, 0xd1, 0x27, 0x77, 0x4e, 0x3e, 0xb6, 0xfc, 0x7d, 0xb6, 0x73, 0xa5, 0x2b,
	0x43, 0x0a, 0xaa, 0x48, 0xe5, 0xad, 0xb7, 0xdc, 0x07, 0x18, 0x94, 0xc6, 0x00, 0x86, 0x82, 0x1d,
	0x2c, 0x32, 0xd4, 0xc7, 0x71, 0x9a, 0x1a, 0xa8, 0xae, 0x2a, 0xeb, 0x63, 0x01, 0x07, 0x69, 0xad,
	0x9f, 0x2f, 0x01, 0x60, 0x8f, 0x09, 0xe2, 0xef, 0xb2, 0xed, 0xb9, 0x0b, 0x6d, 0x86, 0x9e, 0xf7,
	0x82, 0x08, 0x7e, 0xe9, 0xb0, 0xfe, 0x02, 0x03, 0x3f, 0x64, 0xbd, 0x86, 0xcc, 0xdd, 0x95, 0xd2,
	0x9b, 0x66, 0x80, 0x8d, 0xd3, 0xf3, 0xbb, 0x12, 0xea, 0xb3, 0x06, 0x79, 0x42, 0x7b, 0xc1, 0xe8,
	0x7c, 0x59, 0xe5, 0x69, 0xd7, 0x20, 0x42, 0x06, 0x7e, 0xbb, 0x98, 0xc6, 0x0e, 0x12, 0x5e, 0x2e,
	0x9b, 0xb0, 0xa9, 0x5a, 0xbb, 0x40, 0x9a, 0x6d, 0xf8, 0xf4, 0xe2, 0xf0, 0x0a, 0x46, 0x17, 0xcb,
	0x66, 0xa5, 0x4a, 0x46, 0x5d, 0xaa, 0x58, 0xf8, 0xe7, 0x4a, 0xd3, 0x8b, 0x45, 0x66, 0x94, 0xc8,
	0x52, 0xba, 0xc6, 0x1f, 0xb0, 0x3d, 0x3c, 0xff, 0xc7, 0xd8, 0xde, 0x01, 0x41, 0x7b, 0x6a, 0x3f,
	0x6c, 0x86, 0x6a, 0xd3, 0x15, 0xf4, 0xae, 0xf4, 0x10, 0x6c, 0xda, 0xe2, 0x3d, 0xb6, 0xed, 0x6f,
	0x7b, 0xa3, 0x45, 0x8f, 0x4b, 0x9f, 0xd0, 0x46, 0xed, 0x63, 0xb6, 0xef, 0xd5, 0xb2, 0x91, 0x30,
	0xba, 0x72, 0xd2, 0x40, 0xfb, 0xd3, 0x5b, 0xb3, 0x4b, 0xa2, 0xef, 0x46, 0x11, 0x0a, 0xc6, 0x69,
	0xf8, 0xf7, 0x0b, 0x53, 0xa8, 0x71, 0xdf, 0x36, 0x2e, 0xf9, 0xa8, 0xd1, 0xf5, 0x60, 0xf4, 0x6c,
	0xc9, 0xc9, 0xf4, 0x0f, 0xc7, 0x05, 0xf6, 0x28, 0xff, 0x90, 0x71, 0x22, 0x5d, 0x70, 0x9d, 0x92,
	0xb1, 0x83, 0x92, 0x7b, 0xcf, 0x5f, 0x36, 0x04, 0x3b, 0x2f, 0x1b, 0x82, 0x15, 0x0b, 0xbc, 0x22,
	0xa4, 0xdf, 0x8f, 0xa8, 0xa5, 0xdd, 0x81, 0x76, 0xed, 0x23, 0x46, 0x44, 0x4f, 0x33, 0xa3, 0x60,
	0xcc, 0x1f, 0xb4, 0x13, 0x6c, 0xd5, 0x4c, 0xa8, 0xe2, 0x4a, 0xf3, 0x4f, 0xd8, 0x41, 0x9c, 0xc2,
	0xb1, 0x88, 0xb3, 0x4c, 0x15, 0x33, 0x91, 0x4b, 0x77, 0xad, 0x9b, 0x8d, 0x83, 0xc7, 0xe9, 0x64,
	0x2e, 0x7a, 0x82, 0x12, 0x7e, 0xc3, 0x7a, 0xfe, 0x91, 0x40, 0x0b, 0xfe, 0x1a, 0x7f, 0xbf, 0xdc,
	0x29, 0xeb, 0x19, 0x22, 0x36, 0x9d, 0x95, 0x13, 0x35, 0x1b, 0x83, 0xab, 0x37, 0xac, 0xe7, 0x43,
	0x23, 0xe2, 0xce, 0xab, 0x20, 0x9e, 0x27, 0x2f, 0x4b, 0x1b, 0xe2, 0xf0, 0xb7, 0x0e, 0x0b, 0x5a,
	0x3a, 0x50, 0xf3, 0xdc, 0x55, 0x22, 0x57, 0x36, 0x8f, 0x5d, 0x72, 0x2d, 0xd4, 0xac, 0xd0, 0x46,
	0x0e, 0x46, 0x87, 0x2b, 0x47, 0x9b, 0xd1, 0x5e, 0xee, 0xaa, 0x27, 0x5e, 0x32, 0x46, 0x01, 0x2c,
	0x33, 0xbe, 0x47, 0xea, 0xb2, 0x18, 0x3c, 0xa2, 0x61, 0x4e, 0xc8, 0x45, 0xd9, 0x7e, 0x1d, 0x8e,
	0x5b, 0xaf, 0x03, 0xcc, 0xf8, 0xf6, 0x96, 0x27, 0x6e, 0xcd, 0xe0, 0x53, 0xba, 0x6a, 0xad, 0x45,
	0xef, 0xd2, 0xc0, 0x2e, 0x74, 0xbf, 0xea, 0x81, 0xd6, 0x67, 0x34, 0xcc, 0xe7, 0xdb, 0xde, 0xa5,
	0x81, 0xf7, 0xa4, 0xd9, 0xf7, 0x40, 0xe3, 0x73, 0xda, 0x07, 0xfd, 0xca, 0x77, 0x69, 0xa0, 0x0b,
	0x94, 0x15, 0x0d, 0x5d, 0xa2, 0x8b, 0x42, 0x26, 0x30, 0x29, 0xbe, 0xc0, 0x90, 0xb8, 0xb2, 0x13,
	0x12, 0x9d, 0x36, 0x12, 0xfe, 0x33, 0xdb, 0x69, 0x35, 0x0d, 0xd6, 0xe3, 0xcb, 0xe5, 0x0e, 0xba,
	0xc5, 0x46, 0x8d, 0xb6, 0xef, 0xe9, 0xa0, 0x28, 0xd3, 0x2e, 0xee, 0xde, 0x8f, 0xfe, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x7b, 0xf7, 0xfd, 0x90, 0x0b, 0x00, 0x00,
}
