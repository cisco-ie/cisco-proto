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
// source: radius_dead_criteria.proto

package cisco_ios_xr_aaa_protocol_radius_oper_radius_nodes_node_dead_criteria_hosts_host

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

type RadiusDeadCriteria_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	IpAddress            string   `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	AuthPortNumber       uint32   `protobuf:"varint,3,opt,name=auth_port_number,json=authPortNumber,proto3" json:"auth_port_number,omitempty"`
	AcctPortNumber       uint32   `protobuf:"varint,4,opt,name=acct_port_number,json=acctPortNumber,proto3" json:"acct_port_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusDeadCriteria_KEYS) Reset()         { *m = RadiusDeadCriteria_KEYS{} }
func (m *RadiusDeadCriteria_KEYS) String() string { return proto.CompactTextString(m) }
func (*RadiusDeadCriteria_KEYS) ProtoMessage()    {}
func (*RadiusDeadCriteria_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e28d1ef6b660396e, []int{0}
}

func (m *RadiusDeadCriteria_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDeadCriteria_KEYS.Unmarshal(m, b)
}
func (m *RadiusDeadCriteria_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDeadCriteria_KEYS.Marshal(b, m, deterministic)
}
func (m *RadiusDeadCriteria_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDeadCriteria_KEYS.Merge(m, src)
}
func (m *RadiusDeadCriteria_KEYS) XXX_Size() int {
	return xxx_messageInfo_RadiusDeadCriteria_KEYS.Size(m)
}
func (m *RadiusDeadCriteria_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDeadCriteria_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDeadCriteria_KEYS proto.InternalMessageInfo

func (m *RadiusDeadCriteria_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RadiusDeadCriteria_KEYS) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *RadiusDeadCriteria_KEYS) GetAuthPortNumber() uint32 {
	if m != nil {
		return m.AuthPortNumber
	}
	return 0
}

func (m *RadiusDeadCriteria_KEYS) GetAcctPortNumber() uint32 {
	if m != nil {
		return m.AcctPortNumber
	}
	return 0
}

type RadiusTimeTriesData struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	IsComputed           bool     `protobuf:"varint,2,opt,name=is_computed,json=isComputed,proto3" json:"is_computed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusTimeTriesData) Reset()         { *m = RadiusTimeTriesData{} }
func (m *RadiusTimeTriesData) String() string { return proto.CompactTextString(m) }
func (*RadiusTimeTriesData) ProtoMessage()    {}
func (*RadiusTimeTriesData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e28d1ef6b660396e, []int{1}
}

func (m *RadiusTimeTriesData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusTimeTriesData.Unmarshal(m, b)
}
func (m *RadiusTimeTriesData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusTimeTriesData.Marshal(b, m, deterministic)
}
func (m *RadiusTimeTriesData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusTimeTriesData.Merge(m, src)
}
func (m *RadiusTimeTriesData) XXX_Size() int {
	return xxx_messageInfo_RadiusTimeTriesData.Size(m)
}
func (m *RadiusTimeTriesData) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusTimeTriesData.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusTimeTriesData proto.InternalMessageInfo

func (m *RadiusTimeTriesData) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *RadiusTimeTriesData) GetIsComputed() bool {
	if m != nil {
		return m.IsComputed
	}
	return false
}

type RadiusDeadCriteria struct {
	Time                 *RadiusTimeTriesData `protobuf:"bytes,50,opt,name=time,proto3" json:"time,omitempty"`
	Tries                *RadiusTimeTriesData `protobuf:"bytes,51,opt,name=tries,proto3" json:"tries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RadiusDeadCriteria) Reset()         { *m = RadiusDeadCriteria{} }
func (m *RadiusDeadCriteria) String() string { return proto.CompactTextString(m) }
func (*RadiusDeadCriteria) ProtoMessage()    {}
func (*RadiusDeadCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_e28d1ef6b660396e, []int{2}
}

func (m *RadiusDeadCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDeadCriteria.Unmarshal(m, b)
}
func (m *RadiusDeadCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDeadCriteria.Marshal(b, m, deterministic)
}
func (m *RadiusDeadCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDeadCriteria.Merge(m, src)
}
func (m *RadiusDeadCriteria) XXX_Size() int {
	return xxx_messageInfo_RadiusDeadCriteria.Size(m)
}
func (m *RadiusDeadCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDeadCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDeadCriteria proto.InternalMessageInfo

func (m *RadiusDeadCriteria) GetTime() *RadiusTimeTriesData {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *RadiusDeadCriteria) GetTries() *RadiusTimeTriesData {
	if m != nil {
		return m.Tries
	}
	return nil
}

func init() {
	proto.RegisterType((*RadiusDeadCriteria_KEYS)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dead_criteria.hosts.host.radius_dead_criteria_KEYS")
	proto.RegisterType((*RadiusTimeTriesData)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dead_criteria.hosts.host.radius_time_tries_data")
	proto.RegisterType((*RadiusDeadCriteria)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dead_criteria.hosts.host.radius_dead_criteria")
}

func init() { proto.RegisterFile("radius_dead_criteria.proto", fileDescriptor_e28d1ef6b660396e) }

var fileDescriptor_e28d1ef6b660396e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x90, 0xc1, 0x4a, 0x3b, 0x31,
	0x10, 0xc6, 0xd9, 0xfe, 0xdb, 0x3f, 0xed, 0x94, 0x8a, 0x84, 0x22, 0xab, 0x22, 0x96, 0x9e, 0xf6,
	0xb4, 0x87, 0xf6, 0x09, 0x44, 0x3c, 0x09, 0xb5, 0xac, 0x27, 0x4f, 0xc3, 0x34, 0x09, 0x34, 0xd0,
	0x6d, 0x42, 0x32, 0x2b, 0x1e, 0xc4, 0xa7, 0xf0, 0x2d, 0x7c, 0x49, 0xd9, 0x64, 0x41, 0x0a, 0x7b,
	0x14, 0x2f, 0x43, 0xf2, 0xcb, 0x37, 0xf3, 0x7d, 0x13, 0xb8, 0xf2, 0xa4, 0x4c, 0x13, 0x50, 0x69,
	0x52, 0x28, 0xbd, 0x61, 0xed, 0x0d, 0x95, 0xce, 0x5b, 0xb6, 0x62, 0x2b, 0x4d, 0x90, 0x16, 0x8d,
	0x0d, 0xf8, 0xe6, 0x91, 0x88, 0x30, 0x72, 0x69, 0x0f, 0xd8, 0x75, 0x59, 0xa7, 0x7d, 0x99, 0xce,
	0xe5, 0xd1, 0x2a, 0x9d, 0x6a, 0x79, 0x3a, 0x6c, 0x6f, 0x03, 0x87, 0x58, 0x97, 0x5f, 0x19, 0x5c,
	0xf6, 0x19, 0xe2, 0xe3, 0xc3, 0xcb, 0xb3, 0xb8, 0x86, 0x49, 0xdb, 0x8e, 0x47, 0xaa, 0x75, 0x9e,
	0x2d, 0xb2, 0x62, 0x52, 0x8d, 0x5b, 0xb0, 0xa1, 0x5a, 0x8b, 0x1b, 0x00, 0xe3, 0x90, 0x94, 0xf2,
	0x3a, 0x84, 0x7c, 0x10, 0x5f, 0x27, 0xc6, 0xdd, 0x25, 0x20, 0x0a, 0x38, 0xa7, 0x86, 0xf7, 0xe8,
	0xac, 0x67, 0x3c, 0x36, 0xf5, 0x4e, 0xfb, 0xfc, 0xdf, 0x22, 0x2b, 0x66, 0xd5, 0x59, 0xcb, 0xb7,
	0xd6, 0xf3, 0x26, 0xd2, 0xa8, 0x94, 0x92, 0x4f, 0x94, 0xc3, 0x4e, 0x29, 0x25, 0xff, 0x28, 0x97,
	0x4f, 0x70, 0xd1, 0x85, 0x65, 0x53, 0x6b, 0x64, 0x6f, 0x74, 0x40, 0x45, 0x4c, 0x62, 0x0e, 0xa3,
	0x57, 0x3a, 0x34, 0x29, 0xe5, 0xac, 0x4a, 0x17, 0x71, 0x0b, 0x53, 0x13, 0x50, 0xda, 0xda, 0x35,
	0xac, 0x55, 0xcc, 0x38, 0xae, 0xc0, 0x84, 0xfb, 0x8e, 0x2c, 0x3f, 0x07, 0x30, 0xef, 0x5b, 0x5f,
	0xbc, 0xc3, 0xb0, 0xb5, 0xc8, 0x57, 0x8b, 0xac, 0x98, 0xae, 0xf6, 0xe5, 0x6f, 0x7f, 0x7c, 0xd9,
	0xbf, 0x47, 0x15, 0x5d, 0xc5, 0x07, 0x8c, 0x22, 0xcb, 0xd7, 0x7f, 0x6c, 0x9f, 0x6c, 0x77, 0xff,
	0xe3, 0xf8, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda, 0x72, 0x14, 0xf2, 0x8c, 0x02, 0x00,
	0x00,
}
