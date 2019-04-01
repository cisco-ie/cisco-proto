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
// source: snmp_snmpViewXmlEntry.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_views_view_view_information

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

type SnmpSnmpViewXmlEntry_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ObjectId             string   `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpSnmpViewXmlEntry_KEYS) Reset()         { *m = SnmpSnmpViewXmlEntry_KEYS{} }
func (m *SnmpSnmpViewXmlEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpSnmpViewXmlEntry_KEYS) ProtoMessage()    {}
func (*SnmpSnmpViewXmlEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f49eacafd425a504, []int{0}
}

func (m *SnmpSnmpViewXmlEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS.Unmarshal(m, b)
}
func (m *SnmpSnmpViewXmlEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpSnmpViewXmlEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS.Merge(m, src)
}
func (m *SnmpSnmpViewXmlEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS.Size(m)
}
func (m *SnmpSnmpViewXmlEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpSnmpViewXmlEntry_KEYS proto.InternalMessageInfo

func (m *SnmpSnmpViewXmlEntry_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SnmpSnmpViewXmlEntry_KEYS) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

type SnmpSnmpViewXmlEntry struct {
	SnmpViewFamilyType        string   `protobuf:"bytes,50,opt,name=snmp_view_family_type,json=snmpViewFamilyType,proto3" json:"snmp_view_family_type,omitempty"`
	SnmpViewFamilyStorageType string   `protobuf:"bytes,51,opt,name=snmp_view_family_storage_type,json=snmpViewFamilyStorageType,proto3" json:"snmp_view_family_storage_type,omitempty"`
	SnmpViewFamilyStatus      string   `protobuf:"bytes,52,opt,name=snmp_view_family_status,json=snmpViewFamilyStatus,proto3" json:"snmp_view_family_status,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *SnmpSnmpViewXmlEntry) Reset()         { *m = SnmpSnmpViewXmlEntry{} }
func (m *SnmpSnmpViewXmlEntry) String() string { return proto.CompactTextString(m) }
func (*SnmpSnmpViewXmlEntry) ProtoMessage()    {}
func (*SnmpSnmpViewXmlEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_f49eacafd425a504, []int{1}
}

func (m *SnmpSnmpViewXmlEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry.Unmarshal(m, b)
}
func (m *SnmpSnmpViewXmlEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry.Marshal(b, m, deterministic)
}
func (m *SnmpSnmpViewXmlEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpSnmpViewXmlEntry.Merge(m, src)
}
func (m *SnmpSnmpViewXmlEntry) XXX_Size() int {
	return xxx_messageInfo_SnmpSnmpViewXmlEntry.Size(m)
}
func (m *SnmpSnmpViewXmlEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpSnmpViewXmlEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpSnmpViewXmlEntry proto.InternalMessageInfo

func (m *SnmpSnmpViewXmlEntry) GetSnmpViewFamilyType() string {
	if m != nil {
		return m.SnmpViewFamilyType
	}
	return ""
}

func (m *SnmpSnmpViewXmlEntry) GetSnmpViewFamilyStorageType() string {
	if m != nil {
		return m.SnmpViewFamilyStorageType
	}
	return ""
}

func (m *SnmpSnmpViewXmlEntry) GetSnmpViewFamilyStatus() string {
	if m != nil {
		return m.SnmpViewFamilyStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*SnmpSnmpViewXmlEntry_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.views.view.view_information.snmp_snmpViewXmlEntry_KEYS")
	proto.RegisterType((*SnmpSnmpViewXmlEntry)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.views.view.view_information.snmp_snmpViewXmlEntry")
}

func init() { proto.RegisterFile("snmp_snmpViewXmlEntry.proto", fileDescriptor_f49eacafd425a504) }

var fileDescriptor_f49eacafd425a504 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x11, 0xb1, 0x39, 0x06, 0xc5, 0xd5, 0x22, 0x48, 0x4f, 0x9e, 0x16, 0xb4, 0x7a,
	0xf7, 0x52, 0xa1, 0x88, 0x17, 0x2b, 0xa2, 0xa7, 0x21, 0xdd, 0x4e, 0xcb, 0x48, 0x93, 0x09, 0xc9,
	0x68, 0xdd, 0xe7, 0xf3, 0xc5, 0x64, 0x67, 0x15, 0xd4, 0xee, 0xe5, 0x27, 0xc9, 0x37, 0xdf, 0x9f,
	0x10, 0x33, 0xcc, 0xc1, 0x47, 0x68, 0xe3, 0x89, 0x70, 0xf3, 0xec, 0xd7, 0x93, 0x20, 0xa9, 0xa9,
	0x62, 0x62, 0x61, 0x3b, 0xad, 0x29, 0xd7, 0x0c, 0xc4, 0x19, 0x3e, 0x92, 0x0e, 0x81, 0x5b, 0x61,
	0x10, 0xe0, 0x88, 0xa9, 0x6a, 0xf7, 0x15, 0x85, 0x25, 0x27, 0xef, 0x84, 0x38, 0x54, 0xef, 0x84,
	0x9b, 0xac, 0xa9, 0x01, 0xbf, 0xd8, 0xe8, 0xde, 0x9c, 0xf4, 0xde, 0x04, 0x77, 0x93, 0x97, 0x99,
	0xb5, 0x66, 0x37, 0x38, 0x8f, 0x65, 0x71, 0x56, 0x9c, 0x0f, 0x1e, 0x74, 0x6d, 0x87, 0x66, 0xc0,
	0xf3, 0x57, 0xac, 0x05, 0x68, 0x51, 0xee, 0x28, 0xd8, 0xef, 0x0e, 0xa6, 0x8b, 0xd1, 0x67, 0x61,
	0x0e, 0x7b, 0xfb, 0xec, 0xc5, 0x37, 0xd0, 0x17, 0x2c, 0x9d, 0xa7, 0x75, 0x03, 0xd2, 0x44, 0x2c,
	0x2f, 0xb5, 0xc2, 0xfe, 0x08, 0xb7, 0x8a, 0x1e, 0x9b, 0x88, 0xf6, 0xc6, 0x9c, 0x6e, 0x29, 0x59,
	0x38, 0xb9, 0x15, 0x76, 0xea, 0x58, 0xd5, 0xe3, 0xbf, 0xea, 0xac, 0x9b, 0xd0, 0x86, 0x6b, 0x73,
	0xd4, 0xd3, 0xe0, 0xe4, 0x2d, 0x97, 0x57, 0xea, 0x1e, 0xfc, 0x77, 0x5b, 0x36, 0xdf, 0xd3, 0x6f,
	0x1e, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x67, 0x67, 0x3d, 0xdd, 0x85, 0x01, 0x00, 0x00,
}