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
// source: tty_aaa_stats.proto

package cisco_ios_xr_tty_server_oper_tty_auxiliary_nodes_auxiliary_node_auxiliary_line_auxiliary_statistics_aaa

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

type TtyAaaStats_KEYS struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TtyAaaStats_KEYS) Reset()         { *m = TtyAaaStats_KEYS{} }
func (m *TtyAaaStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*TtyAaaStats_KEYS) ProtoMessage()    {}
func (*TtyAaaStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_104d74924245fef1, []int{0}
}

func (m *TtyAaaStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TtyAaaStats_KEYS.Unmarshal(m, b)
}
func (m *TtyAaaStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TtyAaaStats_KEYS.Marshal(b, m, deterministic)
}
func (m *TtyAaaStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TtyAaaStats_KEYS.Merge(m, src)
}
func (m *TtyAaaStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_TtyAaaStats_KEYS.Size(m)
}
func (m *TtyAaaStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TtyAaaStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TtyAaaStats_KEYS proto.InternalMessageInfo

func (m *TtyAaaStats_KEYS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type TtyAaaStats struct {
	UserName             string   `protobuf:"bytes,50,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TtyAaaStats) Reset()         { *m = TtyAaaStats{} }
func (m *TtyAaaStats) String() string { return proto.CompactTextString(m) }
func (*TtyAaaStats) ProtoMessage()    {}
func (*TtyAaaStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_104d74924245fef1, []int{1}
}

func (m *TtyAaaStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TtyAaaStats.Unmarshal(m, b)
}
func (m *TtyAaaStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TtyAaaStats.Marshal(b, m, deterministic)
}
func (m *TtyAaaStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TtyAaaStats.Merge(m, src)
}
func (m *TtyAaaStats) XXX_Size() int {
	return xxx_messageInfo_TtyAaaStats.Size(m)
}
func (m *TtyAaaStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TtyAaaStats.DiscardUnknown(m)
}

var xxx_messageInfo_TtyAaaStats proto.InternalMessageInfo

func (m *TtyAaaStats) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*TtyAaaStats_KEYS)(nil), "cisco_ios_xr_tty_server_oper.tty.auxiliary_nodes.auxiliary_node.auxiliary_line.auxiliary_statistics.aaa.tty_aaa_stats_KEYS")
	proto.RegisterType((*TtyAaaStats)(nil), "cisco_ios_xr_tty_server_oper.tty.auxiliary_nodes.auxiliary_node.auxiliary_line.auxiliary_statistics.aaa.tty_aaa_stats")
}

func init() { proto.RegisterFile("tty_aaa_stats.proto", fileDescriptor_104d74924245fef1) }

var fileDescriptor_104d74924245fef1 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x8b, 0x83, 0x40,
	0x10, 0x46, 0xd1, 0xe2, 0x38, 0x17, 0xee, 0x8a, 0xbd, 0x46, 0xb8, 0x26, 0x48, 0x8a, 0x14, 0x61,
	0x8b, 0xe4, 0x37, 0xa4, 0x0a, 0xa4, 0x48, 0xaa, 0x54, 0xc3, 0x44, 0x87, 0x30, 0xa0, 0xae, 0xcc,
	0x8c, 0x41, 0xff, 0x7d, 0xd0, 0x4a, 0x53, 0xbe, 0x8f, 0x07, 0x8f, 0xcf, 0xfd, 0x99, 0x8d, 0x80,
	0x88, 0xa0, 0x86, 0xa6, 0xa1, 0x93, 0x68, 0xd1, 0x3f, 0x4b, 0xd6, 0x32, 0x02, 0x47, 0x85, 0x41,
	0x60, 0x32, 0x94, 0xe4, 0x45, 0x02, 0xb1, 0x23, 0x09, 0x66, 0x63, 0xc0, 0x7e, 0xe0, 0x9a, 0x51,
	0x46, 0x68, 0x63, 0x45, 0xfa, 0xc1, 0x0b, 0xac, 0xb9, 0x5d, 0xe2, 0x54, 0x61, 0x35, 0x2e, 0x35,
	0x20, 0x62, 0xb1, 0x75, 0x7e, 0xd5, 0x87, 0xf3, 0xe9, 0x7e, 0xf3, 0xbf, 0x2e, 0xe5, 0x2a, 0x4f,
	0x36, 0xc9, 0x2e, 0xbb, 0xa6, 0x5c, 0x15, 0x7b, 0xf7, 0xb3, 0xb2, 0xfc, 0xbf, 0xcb, 0x7a, 0x25,
	0x81, 0x16, 0x1b, 0xca, 0x0f, 0xb3, 0xf7, 0x3d, 0x0d, 0x17, 0x6c, 0xe8, 0xf1, 0x35, 0x7f, 0x38,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x8d, 0x14, 0x94, 0xda, 0x00, 0x00, 0x00,
}