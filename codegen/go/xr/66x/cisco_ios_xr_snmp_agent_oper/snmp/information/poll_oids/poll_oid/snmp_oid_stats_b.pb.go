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
// source: snmp_oid_stats_b.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_poll_oids_poll_oid

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

type SnmpOidStatsB_KEYS struct {
	ObjectId             string   `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpOidStatsB_KEYS) Reset()         { *m = SnmpOidStatsB_KEYS{} }
func (m *SnmpOidStatsB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpOidStatsB_KEYS) ProtoMessage()    {}
func (*SnmpOidStatsB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b611c12304300f5c, []int{0}
}

func (m *SnmpOidStatsB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpOidStatsB_KEYS.Unmarshal(m, b)
}
func (m *SnmpOidStatsB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpOidStatsB_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpOidStatsB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpOidStatsB_KEYS.Merge(m, src)
}
func (m *SnmpOidStatsB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpOidStatsB_KEYS.Size(m)
}
func (m *SnmpOidStatsB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpOidStatsB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpOidStatsB_KEYS proto.InternalMessageInfo

func (m *SnmpOidStatsB_KEYS) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type SnmpOidStatsB struct {
	Nms                  []string `protobuf:"bytes,50,rep,name=nms,proto3" json:"nms,omitempty"`
	RequestCount         []uint32 `protobuf:"varint,51,rep,packed,name=request_count,json=requestCount,proto3" json:"request_count,omitempty"`
	NmsCount             uint32   `protobuf:"varint,52,opt,name=nms_count,json=nmsCount,proto3" json:"nms_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpOidStatsB) Reset()         { *m = SnmpOidStatsB{} }
func (m *SnmpOidStatsB) String() string { return proto.CompactTextString(m) }
func (*SnmpOidStatsB) ProtoMessage()    {}
func (*SnmpOidStatsB) Descriptor() ([]byte, []int) {
	return fileDescriptor_b611c12304300f5c, []int{1}
}

func (m *SnmpOidStatsB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpOidStatsB.Unmarshal(m, b)
}
func (m *SnmpOidStatsB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpOidStatsB.Marshal(b, m, deterministic)
}
func (m *SnmpOidStatsB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpOidStatsB.Merge(m, src)
}
func (m *SnmpOidStatsB) XXX_Size() int {
	return xxx_messageInfo_SnmpOidStatsB.Size(m)
}
func (m *SnmpOidStatsB) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpOidStatsB.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpOidStatsB proto.InternalMessageInfo

func (m *SnmpOidStatsB) GetNms() []string {
	if m != nil {
		return m.Nms
	}
	return nil
}

func (m *SnmpOidStatsB) GetRequestCount() []uint32 {
	if m != nil {
		return m.RequestCount
	}
	return nil
}

func (m *SnmpOidStatsB) GetNmsCount() uint32 {
	if m != nil {
		return m.NmsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpOidStatsB_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.poll_oids.poll_oid.snmp_oid_stats_b_KEYS")
	proto.RegisterType((*SnmpOidStatsB)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.poll_oids.poll_oid.snmp_oid_stats_b")
}

func init() { proto.RegisterFile("snmp_oid_stats_b.proto", fileDescriptor_b611c12304300f5c) }

var fileDescriptor_b611c12304300f5c = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x29, 0x05, 0xd9, 0x06, 0x0b, 0x4b, 0x41, 0x29, 0x78, 0x29, 0xeb, 0x25, 0xa7, 0x1c,
	0xdc, 0xfd, 0x01, 0x82, 0x78, 0x10, 0x6f, 0xf1, 0xe4, 0x69, 0x68, 0xd3, 0xac, 0x44, 0x36, 0x33,
	0x31, 0x33, 0x0b, 0xfe, 0x7c, 0x49, 0x5d, 0xf6, 0xd0, 0xdb, 0x9b, 0x6f, 0xde, 0xbc, 0xe1, 0xa9,
	0x7b, 0xc6, 0x98, 0x80, 0xc2, 0x0c, 0x2c, 0xa3, 0x30, 0x4c, 0x26, 0x65, 0x12, 0xea, 0x9e, 0x5d,
	0x60, 0x47, 0x10, 0x88, 0xe1, 0x37, 0xc3, 0x62, 0x1a, 0xbf, 0x3c, 0x0a, 0x50, 0xf2, 0xd9, 0x94,
	0xd9, 0x04, 0x3c, 0x52, 0x8e, 0xa3, 0x04, 0x42, 0x93, 0xe8, 0x74, 0x2a, 0x29, 0x7c, 0x55, 0xbb,
	0x83, 0xba, 0x5b, 0x67, 0xc3, 0xfb, 0xeb, 0xe7, 0x47, 0xf7, 0xa0, 0x1a, 0x9a, 0xbe, 0xbd, 0x13,
	0x08, 0x73, 0x5f, 0x0d, 0x95, 0x6e, 0xec, 0xe6, 0x1f, 0xbc, 0xcd, 0xbb, 0xa3, 0xda, 0xae, 0xaf,
	0xba, 0xad, 0xaa, 0x31, 0x72, 0xff, 0x34, 0xd4, 0xba, 0xb1, 0x45, 0x76, 0x8f, 0xaa, 0xcd, 0xfe,
	0xe7, 0xec, 0x59, 0xc0, 0xd1, 0x19, 0xa5, 0xdf, 0x0f, 0xb5, 0x6e, 0xed, 0xed, 0x05, 0xbe, 0x14,
	0x56, 0xfe, 0x60, 0xe4, 0x8b, 0xe1, 0x30, 0x54, 0xba, 0xb5, 0x1b, 0x8c, 0xbc, 0x2c, 0xa7, 0x9b,
	0xa5, 0xe6, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xea, 0xfe, 0x41, 0x81, 0x00, 0x01, 0x00, 0x00,
}