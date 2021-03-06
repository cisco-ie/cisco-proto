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
// source: rsvp_mgmt_count_prefix_acl_info.proto

package cisco_ios_xr_ip_rsvp_oper_rsvp_standby_counters_prefix_filtering_accesses_access

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

type RsvpMgmtCountPrefixAclInfo_KEYS struct {
	AccessListName       string   `protobuf:"bytes,1,opt,name=access_list_name,json=accessListName,proto3" json:"access_list_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountPrefixAclInfo_KEYS) Reset()         { *m = RsvpMgmtCountPrefixAclInfo_KEYS{} }
func (m *RsvpMgmtCountPrefixAclInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixAclInfo_KEYS) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixAclInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73bfdfd9d32e191, []int{0}
}

func (m *RsvpMgmtCountPrefixAclInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixAclInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtCountPrefixAclInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS.Merge(m, src)
}
func (m *RsvpMgmtCountPrefixAclInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS.Size(m)
}
func (m *RsvpMgmtCountPrefixAclInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixAclInfo_KEYS proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixAclInfo_KEYS) GetAccessListName() string {
	if m != nil {
		return m.AccessListName
	}
	return ""
}

type RsvpMgmtCountPrefixMsg struct {
	Path                 uint32   `protobuf:"varint,1,opt,name=path,proto3" json:"path,omitempty"`
	PathTear             uint32   `protobuf:"varint,2,opt,name=path_tear,json=pathTear,proto3" json:"path_tear,omitempty"`
	ReservationConfirm   uint32   `protobuf:"varint,3,opt,name=reservation_confirm,json=reservationConfirm,proto3" json:"reservation_confirm,omitempty"`
	Total                uint32   `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsvpMgmtCountPrefixMsg) Reset()         { *m = RsvpMgmtCountPrefixMsg{} }
func (m *RsvpMgmtCountPrefixMsg) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixMsg) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73bfdfd9d32e191, []int{1}
}

func (m *RsvpMgmtCountPrefixMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixMsg.Merge(m, src)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixMsg.Size(m)
}
func (m *RsvpMgmtCountPrefixMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixMsg.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixMsg proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixMsg) GetPath() uint32 {
	if m != nil {
		return m.Path
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetPathTear() uint32 {
	if m != nil {
		return m.PathTear
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetReservationConfirm() uint32 {
	if m != nil {
		return m.ReservationConfirm
	}
	return 0
}

func (m *RsvpMgmtCountPrefixMsg) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type RsvpMgmtCountPrefixAclInfo struct {
	Forwarded            *RsvpMgmtCountPrefixMsg `protobuf:"bytes,50,opt,name=forwarded,proto3" json:"forwarded,omitempty"`
	LocallyDestined      *RsvpMgmtCountPrefixMsg `protobuf:"bytes,51,opt,name=locally_destined,json=locallyDestined,proto3" json:"locally_destined,omitempty"`
	Dropped              *RsvpMgmtCountPrefixMsg `protobuf:"bytes,52,opt,name=dropped,proto3" json:"dropped,omitempty"`
	Total                *RsvpMgmtCountPrefixMsg `protobuf:"bytes,53,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RsvpMgmtCountPrefixAclInfo) Reset()         { *m = RsvpMgmtCountPrefixAclInfo{} }
func (m *RsvpMgmtCountPrefixAclInfo) String() string { return proto.CompactTextString(m) }
func (*RsvpMgmtCountPrefixAclInfo) ProtoMessage()    {}
func (*RsvpMgmtCountPrefixAclInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c73bfdfd9d32e191, []int{2}
}

func (m *RsvpMgmtCountPrefixAclInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo.Unmarshal(m, b)
}
func (m *RsvpMgmtCountPrefixAclInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo.Marshal(b, m, deterministic)
}
func (m *RsvpMgmtCountPrefixAclInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsvpMgmtCountPrefixAclInfo.Merge(m, src)
}
func (m *RsvpMgmtCountPrefixAclInfo) XXX_Size() int {
	return xxx_messageInfo_RsvpMgmtCountPrefixAclInfo.Size(m)
}
func (m *RsvpMgmtCountPrefixAclInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RsvpMgmtCountPrefixAclInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RsvpMgmtCountPrefixAclInfo proto.InternalMessageInfo

func (m *RsvpMgmtCountPrefixAclInfo) GetForwarded() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Forwarded
	}
	return nil
}

func (m *RsvpMgmtCountPrefixAclInfo) GetLocallyDestined() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.LocallyDestined
	}
	return nil
}

func (m *RsvpMgmtCountPrefixAclInfo) GetDropped() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Dropped
	}
	return nil
}

func (m *RsvpMgmtCountPrefixAclInfo) GetTotal() *RsvpMgmtCountPrefixMsg {
	if m != nil {
		return m.Total
	}
	return nil
}

func init() {
	proto.RegisterType((*RsvpMgmtCountPrefixAclInfo_KEYS)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.prefix_filtering.accesses.access.rsvp_mgmt_count_prefix_acl_info_KEYS")
	proto.RegisterType((*RsvpMgmtCountPrefixMsg)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.prefix_filtering.accesses.access.rsvp_mgmt_count_prefix_msg")
	proto.RegisterType((*RsvpMgmtCountPrefixAclInfo)(nil), "cisco_ios_xr_ip_rsvp_oper.rsvp_standby.counters.prefix_filtering.accesses.access.rsvp_mgmt_count_prefix_acl_info")
}

func init() {
	proto.RegisterFile("rsvp_mgmt_count_prefix_acl_info.proto", fileDescriptor_c73bfdfd9d32e191)
}

var fileDescriptor_c73bfdfd9d32e191 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xbd, 0x4e, 0x23, 0x31,
	0x10, 0x96, 0xef, 0x72, 0x3f, 0xf1, 0xe9, 0xee, 0x22, 0xdf, 0x15, 0x2b, 0x28, 0x88, 0x22, 0x90,
	0x52, 0x2d, 0x52, 0x02, 0x4f, 0x00, 0x54, 0x20, 0x14, 0x05, 0x1a, 0xaa, 0x91, 0xe3, 0x9d, 0x0d,
	0x96, 0xbc, 0xb6, 0x35, 0x36, 0x21, 0x69, 0x29, 0x28, 0xa8, 0x91, 0x78, 0x33, 0x9e, 0x07, 0xc5,
	0xbb, 0x11, 0x34, 0x51, 0x1a, 0xa4, 0x54, 0x3b, 0x3b, 0xdf, 0x7c, 0x3f, 0x1e, 0xcb, 0xfc, 0x80,
	0xc2, 0xcc, 0x43, 0x35, 0xad, 0x22, 0x28, 0x77, 0x67, 0x23, 0x78, 0xc2, 0x52, 0xcf, 0x41, 0x2a,
	0x03, 0xda, 0x96, 0x2e, 0xf7, 0xe4, 0xa2, 0x13, 0x23, 0xa5, 0x83, 0x72, 0xa0, 0x5d, 0x80, 0x39,
	0x81, 0xf6, 0x90, 0x68, 0xce, 0x23, 0xe5, 0xa9, 0x0a, 0x51, 0xda, 0x62, 0xb2, 0xc8, 0x93, 0x06,
	0x52, 0xc8, 0x1b, 0x99, 0x52, 0x9b, 0x88, 0xa4, 0xed, 0x34, 0x97, 0x4a, 0x61, 0x08, 0x18, 0x9a,
	0xa2, 0x37, 0xe2, 0xfb, 0x1b, 0xac, 0xe1, 0xfc, 0xec, 0xe6, 0x4a, 0xf4, 0x79, 0xa7, 0x66, 0x80,
	0xd1, 0x21, 0x82, 0x95, 0x15, 0x66, 0xac, 0xcb, 0xfa, 0xed, 0xf1, 0x9f, 0xba, 0x7f, 0xa1, 0x43,
	0xbc, 0x94, 0x15, 0xf6, 0x9e, 0x19, 0xdf, 0x59, 0x23, 0x59, 0x85, 0xa9, 0x10, 0xbc, 0xe5, 0x65,
	0xbc, 0x4d, 0xe4, 0xdf, 0xe3, 0x54, 0x8b, 0x5d, 0xde, 0x5e, 0x7e, 0x21, 0xa2, 0xa4, 0xec, 0x4b,
	0x02, 0x7e, 0x2e, 0x1b, 0xd7, 0x28, 0x49, 0x1c, 0xf2, 0x7f, 0x84, 0x01, 0x69, 0x26, 0xa3, 0x76,
	0x16, 0x94, 0xb3, 0xa5, 0xa6, 0x2a, 0xfb, 0x9a, 0xc6, 0xc4, 0x07, 0xe8, 0xa4, 0x46, 0xc4, 0x7f,
	0xfe, 0x2d, 0xba, 0x28, 0x4d, 0xd6, 0x4a, 0x23, 0xf5, 0x4f, 0xef, 0xb5, 0xc5, 0xf7, 0x36, 0x9c,
	0x54, 0x3c, 0x31, 0xde, 0x2e, 0x1d, 0xdd, 0x4b, 0x2a, 0xb0, 0xc8, 0x06, 0x5d, 0xd6, 0xff, 0x35,
	0x30, 0xf9, 0x67, 0xef, 0x3c, 0x5f, 0xbf, 0x9d, 0xf1, 0xbb, 0xbd, 0x78, 0x61, 0xbc, 0x63, 0x9c,
	0x92, 0xc6, 0x2c, 0xa0, 0xc0, 0x10, 0xb5, 0xc5, 0x22, 0x1b, 0x6e, 0x21, 0xd3, 0xdf, 0x26, 0xc5,
	0x69, 0x13, 0x42, 0x3c, 0x32, 0xfe, 0xa3, 0x20, 0xe7, 0x3d, 0x16, 0xd9, 0xd1, 0x16, 0x02, 0xad,
	0xcc, 0xc5, 0x03, 0x5b, 0x5d, 0xf5, 0xf1, 0x16, 0x62, 0xd4, 0xd6, 0x93, 0xef, 0xe9, 0x69, 0x0e,
	0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x22, 0xa1, 0x9a, 0xc3, 0x03, 0x00, 0x00,
}
