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
// source: rsvp_mgmt_frr_states_detail.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_frrs_frr

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

type RsvpMgmtFrrStatesDetail_KEYS struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	DestinationPort      uint32   `protobuf:"varint,2,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	Protocol             uint32   `protobuf:"varint,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,4,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	SessionType          string   `protobuf:"bytes,5,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	P2MpId               uint32   `protobuf:"varint,6,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	SourceAddress        string   `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	SourcePort           uint32   `protobuf:"varint,8,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	SubGroupOrigin       string   `protobuf:"bytes,9,opt,name=sub_group_origin,json=subGroupOrigin,proto3" json:"sub_group_origin,omitempty"`
	SubGroupId           uint32   `protobuf:"varint,10,opt,name=sub_group_id,json=subGroupId,proto3" json:"sub_group_id,omitempty"`
	VrfName              string   `protobuf:"bytes,11,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) Reset()         { *m = RsvpMgmtFrrStatesDetail_KEYS{} }
func (m *RsvpMgmtFrrStatesDetail_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesDetail_KEYS) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesDetail_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{0}
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Merge(m, src)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.Size(m)
}
func (m *RsvpMgmtFrrStatesDetail_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesDetail_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetDestinationPort() uint32 {
	if m != nil {
		return m.DestinationPort
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSourcePort() uint32 {
	if m != nil {
		return m.SourcePort
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSubGroupOrigin() string {
	if m != nil {
		return m.SubGroupOrigin
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetSubGroupId() uint32 {
	if m != nil {
		return m.SubGroupId
	}
	return 0
}

func (m *RsvpMgmtFrrStatesDetail_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type RsvpMgmtSessionUdpIpv4 struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	Protocol             uint32   `protobuf:"varint,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	DestinationPort      uint32   `protobuf:"varint,3,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionUdpIpv4) Reset()         { *m = RsvpMgmtSessionUdpIpv4{} }
func (m *RsvpMgmtSessionUdpIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionUdpIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionUdpIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{1}
}

func (m *RsvpMgmtSessionUdpIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Merge(m, src)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionUdpIpv4.Size(m)
}
func (m *RsvpMgmtSessionUdpIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionUdpIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionUdpIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionUdpIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionUdpIpv4) GetProtocol() uint32 {
	if m != nil {
		return m.Protocol
	}
	return 0
}

func (m *RsvpMgmtSessionUdpIpv4) GetDestinationPort() uint32 {
	if m != nil {
		return m.DestinationPort
	}
	return 0
}

type RsvpMgmtSessionLspTunnelIpv4 struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionLspTunnelIpv4) Reset()         { *m = RsvpMgmtSessionLspTunnelIpv4{} }
func (m *RsvpMgmtSessionLspTunnelIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionLspTunnelIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionLspTunnelIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{2}
}

func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Merge(m, src)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.Size(m)
}
func (m *RsvpMgmtSessionLspTunnelIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionLspTunnelIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionLspTunnelIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionLspTunnelIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionLspTunnelIpv4) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

type RsvpMgmtSessionUniIpv4 struct {
	DestinationAddress   string   `protobuf:"bytes,1,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedAddress      string   `protobuf:"bytes,3,opt,name=extended_address,json=extendedAddress,proto3" json:"extended_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionUniIpv4) Reset()         { *m = RsvpMgmtSessionUniIpv4{} }
func (m *RsvpMgmtSessionUniIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionUniIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionUniIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{3}
}

func (m *RsvpMgmtSessionUniIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionUniIpv4.Merge(m, src)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionUniIpv4.Size(m)
}
func (m *RsvpMgmtSessionUniIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionUniIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionUniIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionUniIpv4) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

func (m *RsvpMgmtSessionUniIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionUniIpv4) GetExtendedAddress() string {
	if m != nil {
		return m.ExtendedAddress
	}
	return ""
}

type RsvpMgmtSessionP2MpLspTunnelIpv4 struct {
	P2MpId               uint32   `protobuf:"varint,1,opt,name=p2mp_id,json=p2mpId,proto3" json:"p2mp_id,omitempty"`
	TunnelId             uint32   `protobuf:"varint,2,opt,name=tunnel_id,json=tunnelId,proto3" json:"tunnel_id,omitempty"`
	ExtendedTunnelId     string   `protobuf:"bytes,3,opt,name=extended_tunnel_id,json=extendedTunnelId,proto3" json:"extended_tunnel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) Reset()         { *m = RsvpMgmtSessionP2MpLspTunnelIpv4{} }
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionP2MpLspTunnelIpv4) ProtoMessage()    {}
func (*RsvpMgmtSessionP2MpLspTunnelIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{4}
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Merge(m, src)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.Size(m)
}
func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionP2MpLspTunnelIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetP2MpId() uint32 {
	if m != nil {
		return m.P2MpId
	}
	return 0
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetTunnelId() uint32 {
	if m != nil {
		return m.TunnelId
	}
	return 0
}

func (m *RsvpMgmtSessionP2MpLspTunnelIpv4) GetExtendedTunnelId() string {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return ""
}

type RsvpSessionUnion struct {
	SessionType          string                            `protobuf:"bytes,1,opt,name=session_type,json=sessionType,proto3" json:"session_type,omitempty"`
	Ipv4                 *RsvpMgmtSessionUdpIpv4           `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv4LspSession       *RsvpMgmtSessionLspTunnelIpv4     `protobuf:"bytes,3,opt,name=ipv4_lsp_session,json=ipv4LspSession,proto3" json:"ipv4_lsp_session,omitempty"`
	Ipv4UniSession       *RsvpMgmtSessionUniIpv4           `protobuf:"bytes,4,opt,name=ipv4_uni_session,json=ipv4UniSession,proto3" json:"ipv4_uni_session,omitempty"`
	Ipv4P2MpLspSession   *RsvpMgmtSessionP2MpLspTunnelIpv4 `protobuf:"bytes,5,opt,name=ipv4_p2mp_lsp_session,json=ipv4P2mpLspSession,proto3" json:"ipv4_p2mp_lsp_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *RsvpSessionUnion) Reset()         { *m = RsvpSessionUnion{} }
func (m *RsvpSessionUnion) String() string { return proto.CompactTextString(m) }
func (*RsvpSessionUnion) ProtoMessage()    {}
func (*RsvpSessionUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{5}
}

func (m *RsvpSessionUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpSessionUnion.Unmarshal(m, b)
}
func (m *RsvpSessionUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpSessionUnion.Marshal(b, m, deterministic)
}
func (m *RsvpSessionUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpSessionUnion.Merge(m, src)
}
func (m *RsvpSessionUnion) XXX_Size() int {
	return xxx_messageInfo_RsvpSessionUnion.Size(m)
}
func (m *RsvpSessionUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpSessionUnion.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpSessionUnion proto.InternalMessageInfo

func (m *RsvpSessionUnion) GetSessionType() string {
	if m != nil {
		return m.SessionType
	}
	return ""
}

func (m *RsvpSessionUnion) GetIpv4() *RsvpMgmtSessionUdpIpv4 {
	if m != nil {
		return m.Ipv4
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4LspSession() *RsvpMgmtSessionLspTunnelIpv4 {
	if m != nil {
		return m.Ipv4LspSession
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4UniSession() *RsvpMgmtSessionUniIpv4 {
	if m != nil {
		return m.Ipv4UniSession
	}
	return nil
}

func (m *RsvpSessionUnion) GetIpv4P2MpLspSession() *RsvpMgmtSessionP2MpLspTunnelIpv4 {
	if m != nil {
		return m.Ipv4P2MpLspSession
	}
	return nil
}

type RsvpMgmtSessionInfo struct {
	RsvpSession          *RsvpSessionUnion `protobuf:"bytes,1,opt,name=rsvp_session,json=rsvpSession,proto3" json:"rsvp_session,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RsvpMgmtSessionInfo) Reset()         { *m = RsvpMgmtSessionInfo{} }
func (m *RsvpMgmtSessionInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtSessionInfo) ProtoMessage()    {}
func (*RsvpMgmtSessionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{6}
}

func (m *RsvpMgmtSessionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtSessionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtSessionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtSessionInfo.Merge(m, src)
}
func (m *RsvpMgmtSessionInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtSessionInfo.Size(m)
}
func (m *RsvpMgmtSessionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtSessionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtSessionInfo proto.InternalMessageInfo

func (m *RsvpMgmtSessionInfo) GetRsvpSession() *RsvpSessionUnion {
	if m != nil {
		return m.RsvpSession
	}
	return nil
}

type RsvpMgmtS2LSubLspIpv4 struct {
	S2LDestinationAddress string   `protobuf:"bytes,1,opt,name=s2l_destination_address,json=s2lDestinationAddress,proto3" json:"s2l_destination_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *RsvpMgmtS2LSubLspIpv4) Reset()         { *m = RsvpMgmtS2LSubLspIpv4{} }
func (m *RsvpMgmtS2LSubLspIpv4) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtS2LSubLspIpv4) ProtoMessage()    {}
func (*RsvpMgmtS2LSubLspIpv4) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{7}
}

func (m *RsvpMgmtS2LSubLspIpv4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Unmarshal(m, b)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Merge(m, src)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.Size(m)
}
func (m *RsvpMgmtS2LSubLspIpv4) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtS2LSubLspIpv4.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtS2LSubLspIpv4 proto.InternalMessageInfo

func (m *RsvpMgmtS2LSubLspIpv4) GetS2LDestinationAddress() string {
	if m != nil {
		return m.S2LDestinationAddress
	}
	return ""
}

type RsvpMgmtFrrStatesDetail struct {
	Session              *RsvpMgmtSessionInfo   `protobuf:"bytes,50,opt,name=session,proto3" json:"session,omitempty"`
	S2LSubLsp            *RsvpMgmtS2LSubLspIpv4 `protobuf:"bytes,51,opt,name=s2l_sub_lsp,json=s2lSubLsp,proto3" json:"s2l_sub_lsp,omitempty"`
	PathStatus           string                 `protobuf:"bytes,52,opt,name=path_status,json=pathStatus,proto3" json:"path_status,omitempty"`
	ReservationStatus    string                 `protobuf:"bytes,53,opt,name=reservation_status,json=reservationStatus,proto3" json:"reservation_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RsvpMgmtFrrStatesDetail) Reset()         { *m = RsvpMgmtFrrStatesDetail{} }
func (m *RsvpMgmtFrrStatesDetail) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtFrrStatesDetail) ProtoMessage()    {}
func (*RsvpMgmtFrrStatesDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e6155f478f19dd0, []int{8}
}

func (m *RsvpMgmtFrrStatesDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Unmarshal(m, b)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail.Merge(m, src)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtFrrStatesDetail.Size(m)
}
func (m *RsvpMgmtFrrStatesDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtFrrStatesDetail.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtFrrStatesDetail proto.InternalMessageInfo

func (m *RsvpMgmtFrrStatesDetail) GetSession() *RsvpMgmtSessionInfo {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *RsvpMgmtFrrStatesDetail) GetS2LSubLsp() *RsvpMgmtS2LSubLspIpv4 {
	if m != nil {
		return m.S2LSubLsp
	}
	return nil
}

func (m *RsvpMgmtFrrStatesDetail) GetPathStatus() string {
	if m != nil {
		return m.PathStatus
	}
	return ""
}

func (m *RsvpMgmtFrrStatesDetail) GetReservationStatus() string {
	if m != nil {
		return m.ReservationStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*RsvpMgmtFrrStatesDetail_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_frr_states_detail_KEYS")
	proto.RegisterType((*RsvpMgmtSessionUdpIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_session_udp_ipv4")
	proto.RegisterType((*RsvpMgmtSessionLspTunnelIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_session_lsp_tunnel_ipv4")
	proto.RegisterType((*RsvpMgmtSessionUniIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_session_uni_ipv4")
	proto.RegisterType((*RsvpMgmtSessionP2MpLspTunnelIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_session_p2mp_lsp_tunnel_ipv4")
	proto.RegisterType((*RsvpSessionUnion)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_session_union")
	proto.RegisterType((*RsvpMgmtSessionInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_session_info")
	proto.RegisterType((*RsvpMgmtS2LSubLspIpv4)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_s2l_sub_lsp_ipv4")
	proto.RegisterType((*RsvpMgmtFrrStatesDetail)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.frrs.frr.rsvp_mgmt_frr_states_detail")
}

func init() { proto.RegisterFile("rsvp_mgmt_frr_states_detail.proto", fileDescriptor_6e6155f478f19dd0) }

var fileDescriptor_6e6155f478f19dd0 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0x96, 0x9b, 0xb4, 0x69, 0xc6, 0xfd, 0x91, 0xb7, 0x4f, 0x7d, 0xf5, 0x6b, 0x0f, 0x2f, 0x8d,
	0xf4, 0x50, 0x2b, 0x41, 0x2a, 0xb9, 0x85, 0x3b, 0x02, 0x54, 0x55, 0xad, 0xa0, 0x72, 0x8a, 0x10,
	0xa7, 0x95, 0x13, 0x6f, 0xca, 0x0a, 0xc7, 0xbb, 0xda, 0x5d, 0x47, 0xad, 0xc4, 0x81, 0x33, 0xe2,
	0x08, 0x17, 0x2e, 0xfc, 0x11, 0x9c, 0xf9, 0xdf, 0xd0, 0x8e, 0xed, 0xc4, 0x4d, 0xd2, 0x4a, 0x2d,
	0xe1, 0xe2, 0xc4, 0xb3, 0xb3, 0xf3, 0x7d, 0xf3, 0xcd, 0xb7, 0x5e, 0xd8, 0x51, 0x7a, 0x28, 0xe9,
	0xe0, 0x62, 0x60, 0x68, 0x5f, 0x29, 0xaa, 0x4d, 0x68, 0x98, 0xa6, 0x11, 0x33, 0x21, 0x8f, 0xdb,
	0x52, 0x09, 0x23, 0xc8, 0x7e, 0x8f, 0xeb, 0x9e, 0xa0, 0x5c, 0x68, 0x7a, 0xa9, 0x28, 0x97, 0x14,
	0xb7, 0x08, 0xc9, 0x54, 0x1b, 0xff, 0x69, 0x13, 0x26, 0x51, 0xf7, 0xaa, 0xdd, 0x57, 0x4a, 0xdb,
	0x47, 0xeb, 0x47, 0x05, 0x9a, 0xb7, 0x94, 0xa5, 0x27, 0x2f, 0xde, 0x76, 0xc8, 0x3e, 0xfc, 0x1d,
	0x31, 0x6d, 0x78, 0x12, 0x1a, 0x2e, 0x12, 0x1a, 0x46, 0x91, 0x62, 0x5a, 0x7b, 0x4e, 0xd3, 0xd9,
	0xad, 0x07, 0xa4, 0xb4, 0xf4, 0x34, 0x5b, 0x21, 0x7b, 0xd0, 0x28, 0x6f, 0x90, 0x42, 0x19, 0x6f,
	0xa1, 0xe9, 0xec, 0xae, 0x06, 0xeb, 0xa5, 0xf8, 0x99, 0x50, 0x86, 0x6c, 0xc1, 0x32, 0x52, 0xef,
	0x89, 0xd8, 0xab, 0x60, 0xca, 0xe8, 0x9d, 0x3c, 0x04, 0xc2, 0x2e, 0x0d, 0x4b, 0x22, 0x16, 0x51,
	0x93, 0x26, 0x09, 0x8b, 0x29, 0x8f, 0xbc, 0x2a, 0xc2, 0x36, 0x8a, 0x95, 0x73, 0x5c, 0x38, 0x8e,
	0xc8, 0x0e, 0xac, 0x68, 0xa6, 0xb5, 0x05, 0x34, 0x57, 0x92, 0x79, 0x8b, 0x98, 0xe7, 0xe6, 0xb1,
	0xf3, 0x2b, 0xc9, 0xc8, 0x26, 0xd4, 0xa4, 0x3f, 0x90, 0xb6, 0xca, 0x12, 0x62, 0x2d, 0xd9, 0xd7,
	0xe3, 0x88, 0xfc, 0x0f, 0x6b, 0x5a, 0xa4, 0xaa, 0xc7, 0x46, 0xcd, 0xd5, 0x70, 0xf7, 0x6a, 0x16,
	0x2d, 0xfa, 0xfa, 0x0f, 0xdc, 0x3c, 0x0d, 0x5b, 0x5a, 0xc6, 0x1a, 0x90, 0x85, 0xb0, 0x9b, 0x5d,
	0x68, 0xe8, 0xb4, 0x4b, 0x2f, 0x94, 0x48, 0x25, 0x15, 0x8a, 0x5f, 0xf0, 0xc4, 0xab, 0x63, 0xa5,
	0x35, 0x9d, 0x76, 0x8f, 0x6c, 0xf8, 0x15, 0x46, 0x49, 0x13, 0x56, 0xc6, 0x99, 0x3c, 0xf2, 0x20,
	0xaf, 0x95, 0x67, 0x1d, 0x47, 0xe4, 0x5f, 0x58, 0x1e, 0xaa, 0x3e, 0x4d, 0xc2, 0x01, 0xf3, 0x5c,
	0xac, 0x51, 0x1b, 0xaa, 0xfe, 0xcb, 0x70, 0xc0, 0x5a, 0x5f, 0x1c, 0xd8, 0x1a, 0x4f, 0xad, 0xe8,
	0x3a, 0x8d, 0x24, 0xe5, 0x72, 0x78, 0x78, 0xf7, 0x79, 0x95, 0x87, 0xb0, 0x30, 0x31, 0x84, 0x59,
	0xb3, 0xac, 0xcc, 0x9c, 0x65, 0xeb, 0xbb, 0x53, 0xf6, 0x68, 0x41, 0x2b, 0xd6, 0x72, 0x34, 0xbd,
	0x7b, 0xb1, 0xdb, 0x86, 0xfa, 0x78, 0xfa, 0x39, 0x3d, 0x53, 0x4c, 0x7d, 0xb6, 0x47, 0x2a, 0xb3,
	0x3d, 0xd2, 0xfa, 0x3a, 0x5b, 0xb8, 0x84, 0xff, 0x09, 0x6a, 0x7b, 0x30, 0x22, 0x30, 0x2a, 0x95,
	0x11, 0x5b, 0x2f, 0xe2, 0x79, 0x9d, 0xd6, 0x67, 0x07, 0x1e, 0x4c, 0xf3, 0x42, 0xaf, 0x4e, 0xca,
	0x57, 0xf2, 0xb0, 0x73, 0xcd, 0xc3, 0x73, 0x94, 0xe9, 0x5b, 0x15, 0x48, 0xf6, 0xbd, 0x18, 0x2b,
	0x24, 0x92, 0xa9, 0x13, 0xe6, 0x4c, 0x9f, 0x30, 0x0a, 0x55, 0xcb, 0x12, 0xf1, 0x5d, 0xff, 0xa4,
	0x7d, 0xc7, 0xef, 0x51, 0xfb, 0x66, 0x57, 0x07, 0x58, 0x98, 0x7c, 0x80, 0x86, 0xfd, 0x45, 0x59,
	0xf2, 0x14, 0x6c, 0xc3, 0xf5, 0x83, 0x39, 0x80, 0x4d, 0x88, 0x1d, 0xac, 0xd9, 0xe7, 0xa9, 0x96,
	0x9d, 0x6c, 0x9d, 0xa4, 0x39, 0xba, 0x75, 0x4c, 0x81, 0x5e, 0x9d, 0x5f, 0xab, 0xb9, 0x0f, 0x33,
	0xd8, 0xd7, 0x09, 0x2f, 0x60, 0x3f, 0x39, 0xb0, 0x81, 0xb8, 0x23, 0x47, 0x14, 0xe0, 0x8b, 0x08,
	0xfe, 0x66, 0x0e, 0xe0, 0xb3, 0xcc, 0x16, 0x10, 0xfb, 0x3c, 0xf3, 0x07, 0x72, 0xac, 0x41, 0xeb,
	0xa3, 0x03, 0xff, 0x4c, 0x6f, 0xe7, 0x49, 0x5f, 0x90, 0x3e, 0xac, 0x94, 0x6d, 0x83, 0x06, 0x71,
	0xfd, 0x67, 0xf7, 0x63, 0x77, 0xcd, 0x7b, 0x81, 0x6b, 0x63, 0x05, 0x85, 0xf3, 0x6b, 0xa7, 0xd8,
	0x8f, 0xa9, 0xfd, 0x94, 0x5a, 0xee, 0x68, 0x91, 0x27, 0xb0, 0x69, 0x63, 0x37, 0x9f, 0xe4, 0x0d,
	0xed, 0xc7, 0xcf, 0xa7, 0x0e, 0x73, 0xeb, 0xe7, 0x02, 0x6c, 0xdf, 0x72, 0x17, 0x92, 0x10, 0x6a,
	0x45, 0x63, 0x3e, 0x36, 0x76, 0x34, 0x07, 0xd9, 0xad, 0x6e, 0x41, 0x51, 0x97, 0xbc, 0x07, 0xb7,
	0xd4, 0x8e, 0x77, 0xf0, 0xfb, 0xd6, 0x9a, 0x10, 0x27, 0xa8, 0x6b, 0x3f, 0xee, 0xa4, 0xdd, 0x53,
	0x2d, 0xed, 0x6d, 0x26, 0x43, 0xf3, 0x0e, 0xbb, 0x4c, 0xb5, 0x77, 0x88, 0xda, 0x80, 0x0d, 0x75,
	0x30, 0x42, 0x1e, 0x01, 0x51, 0x4c, 0x33, 0x35, 0xcc, 0x44, 0xcc, 0xf3, 0x1e, 0x63, 0xde, 0x5f,
	0xa5, 0x95, 0x2c, 0xbd, 0xbb, 0x84, 0x77, 0xc6, 0xc1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x8f, 0x27, 0xd0, 0xa8, 0x08, 0x00, 0x00,
}
