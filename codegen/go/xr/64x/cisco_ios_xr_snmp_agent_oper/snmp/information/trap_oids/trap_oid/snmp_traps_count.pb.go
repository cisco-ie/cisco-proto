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
// source: snmp_traps_count.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_trap_oids_trap_oid

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

type SnmpTrapsCount_KEYS struct {
	TrapOid              string   `protobuf:"bytes,1,opt,name=trap_oid,json=trapOid,proto3" json:"trap_oid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpTrapsCount_KEYS) Reset()         { *m = SnmpTrapsCount_KEYS{} }
func (m *SnmpTrapsCount_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapsCount_KEYS) ProtoMessage()    {}
func (*SnmpTrapsCount_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cf287dcad5313d, []int{0}
}

func (m *SnmpTrapsCount_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapsCount_KEYS.Unmarshal(m, b)
}
func (m *SnmpTrapsCount_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapsCount_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpTrapsCount_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapsCount_KEYS.Merge(m, src)
}
func (m *SnmpTrapsCount_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapsCount_KEYS.Size(m)
}
func (m *SnmpTrapsCount_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapsCount_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapsCount_KEYS proto.InternalMessageInfo

func (m *SnmpTrapsCount_KEYS) GetTrapOid() string {
	if m != nil {
		return m.TrapOid
	}
	return ""
}

type SnmpTrapsCount struct {
	TrapOidCount         uint32   `protobuf:"varint,50,opt,name=trap_oid_count,json=trapOidCount,proto3" json:"trap_oid_count,omitempty"`
	TrapOidXr            string   `protobuf:"bytes,51,opt,name=trap_oid_xr,json=trapOidXr,proto3" json:"trap_oid_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpTrapsCount) Reset()         { *m = SnmpTrapsCount{} }
func (m *SnmpTrapsCount) String() string { return proto.CompactTextString(m) }
func (*SnmpTrapsCount) ProtoMessage()    {}
func (*SnmpTrapsCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1cf287dcad5313d, []int{1}
}

func (m *SnmpTrapsCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpTrapsCount.Unmarshal(m, b)
}
func (m *SnmpTrapsCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpTrapsCount.Marshal(b, m, deterministic)
}
func (m *SnmpTrapsCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpTrapsCount.Merge(m, src)
}
func (m *SnmpTrapsCount) XXX_Size() int {
	return xxx_messageInfo_SnmpTrapsCount.Size(m)
}
func (m *SnmpTrapsCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpTrapsCount.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpTrapsCount proto.InternalMessageInfo

func (m *SnmpTrapsCount) GetTrapOidCount() uint32 {
	if m != nil {
		return m.TrapOidCount
	}
	return 0
}

func (m *SnmpTrapsCount) GetTrapOidXr() string {
	if m != nil {
		return m.TrapOidXr
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpTrapsCount_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_oids.trap_oid.snmp_traps_count_KEYS")
	proto.RegisterType((*SnmpTrapsCount)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.trap_oids.trap_oid.snmp_traps_count")
}

func init() { proto.RegisterFile("snmp_traps_count.proto", fileDescriptor_e1cf287dcad5313d) }

var fileDescriptor_e1cf287dcad5313d = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0xce, 0xcb, 0x2d,
	0x88, 0x2f, 0x29, 0x4a, 0x2c, 0x28, 0x8e, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x72, 0x48, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28,
	0x8a, 0x07, 0x2b, 0x4a, 0x4c, 0x4f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1,
	0xf5, 0x32, 0xf3, 0xd2, 0xf2, 0x8b, 0x72, 0x13, 0x4b, 0x32, 0xf3, 0xf3, 0xf4, 0x40, 0x06, 0xc4,
	0xe7, 0x67, 0xa6, 0x14, 0xc3, 0x59, 0x4a, 0x46, 0x5c, 0xa2, 0xe8, 0x66, 0xc7, 0x7b, 0xbb, 0x46,
	0x06, 0x0b, 0x49, 0x72, 0x71, 0xc0, 0x14, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x83,
	0xf8, 0xfe, 0x99, 0x29, 0x4a, 0x11, 0x5c, 0x02, 0xe8, 0x7a, 0x84, 0x54, 0xb8, 0xf8, 0x60, 0xca,
	0x21, 0x22, 0x12, 0x46, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x3c, 0x50, 0x4d, 0xce, 0x60, 0x55, 0x72,
	0x5c, 0xdc, 0x70, 0x55, 0x15, 0x45, 0x12, 0xc6, 0x60, 0x73, 0x39, 0xa1, 0x4a, 0x22, 0x8a, 0x92,
	0xd8, 0xc0, 0xde, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xc9, 0xd2, 0xfb, 0xf0, 0x00,
	0x00, 0x00,
}
