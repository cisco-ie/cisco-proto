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
// source: snmp_bulkstats_b.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_bulk_stats_transfers_bulk_stats_transfer

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

type SnmpBulkstatsB_KEYS struct {
	TransferName         string   `protobuf:"bytes,1,opt,name=transfer_name,json=transferName,proto3" json:"transfer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpBulkstatsB_KEYS) Reset()         { *m = SnmpBulkstatsB_KEYS{} }
func (m *SnmpBulkstatsB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpBulkstatsB_KEYS) ProtoMessage()    {}
func (*SnmpBulkstatsB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c63a450caece32cb, []int{0}
}

func (m *SnmpBulkstatsB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpBulkstatsB_KEYS.Unmarshal(m, b)
}
func (m *SnmpBulkstatsB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpBulkstatsB_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpBulkstatsB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpBulkstatsB_KEYS.Merge(m, src)
}
func (m *SnmpBulkstatsB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpBulkstatsB_KEYS.Size(m)
}
func (m *SnmpBulkstatsB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpBulkstatsB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpBulkstatsB_KEYS proto.InternalMessageInfo

func (m *SnmpBulkstatsB_KEYS) GetTransferName() string {
	if m != nil {
		return m.TransferName
	}
	return ""
}

type SnmpBulkstatsB struct {
	TransferNameXr       string   `protobuf:"bytes,50,opt,name=transfer_name_xr,json=transferNameXr,proto3" json:"transfer_name_xr,omitempty"`
	UrlPrimary           string   `protobuf:"bytes,51,opt,name=url_primary,json=urlPrimary,proto3" json:"url_primary,omitempty"`
	UrlSecondary         string   `protobuf:"bytes,52,opt,name=url_secondary,json=urlSecondary,proto3" json:"url_secondary,omitempty"`
	RetainedFile         string   `protobuf:"bytes,53,opt,name=retained_file,json=retainedFile,proto3" json:"retained_file,omitempty"`
	TimeLeft             uint32   `protobuf:"varint,54,opt,name=time_left,json=timeLeft,proto3" json:"time_left,omitempty"`
	RetryLeft            uint32   `protobuf:"varint,55,opt,name=retry_left,json=retryLeft,proto3" json:"retry_left,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpBulkstatsB) Reset()         { *m = SnmpBulkstatsB{} }
func (m *SnmpBulkstatsB) String() string { return proto.CompactTextString(m) }
func (*SnmpBulkstatsB) ProtoMessage()    {}
func (*SnmpBulkstatsB) Descriptor() ([]byte, []int) {
	return fileDescriptor_c63a450caece32cb, []int{1}
}

func (m *SnmpBulkstatsB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpBulkstatsB.Unmarshal(m, b)
}
func (m *SnmpBulkstatsB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpBulkstatsB.Marshal(b, m, deterministic)
}
func (m *SnmpBulkstatsB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpBulkstatsB.Merge(m, src)
}
func (m *SnmpBulkstatsB) XXX_Size() int {
	return xxx_messageInfo_SnmpBulkstatsB.Size(m)
}
func (m *SnmpBulkstatsB) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpBulkstatsB.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpBulkstatsB proto.InternalMessageInfo

func (m *SnmpBulkstatsB) GetTransferNameXr() string {
	if m != nil {
		return m.TransferNameXr
	}
	return ""
}

func (m *SnmpBulkstatsB) GetUrlPrimary() string {
	if m != nil {
		return m.UrlPrimary
	}
	return ""
}

func (m *SnmpBulkstatsB) GetUrlSecondary() string {
	if m != nil {
		return m.UrlSecondary
	}
	return ""
}

func (m *SnmpBulkstatsB) GetRetainedFile() string {
	if m != nil {
		return m.RetainedFile
	}
	return ""
}

func (m *SnmpBulkstatsB) GetTimeLeft() uint32 {
	if m != nil {
		return m.TimeLeft
	}
	return 0
}

func (m *SnmpBulkstatsB) GetRetryLeft() uint32 {
	if m != nil {
		return m.RetryLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpBulkstatsB_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.bulk_stats_transfers.bulk_stats_transfer.snmp_bulkstats_b_KEYS")
	proto.RegisterType((*SnmpBulkstatsB)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.bulk_stats_transfers.bulk_stats_transfer.snmp_bulkstats_b")
}

func init() { proto.RegisterFile("snmp_bulkstats_b.proto", fileDescriptor_c63a450caece32cb) }

var fileDescriptor_c63a450caece32cb = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xc9, 0x45, 0xcc, 0x6a, 0xa5, 0x2c, 0x28, 0x01, 0x11, 0x4b, 0xbd, 0xe4, 0x94, 0x83,
	0xf5, 0xe3, 0xe2, 0x55, 0x2f, 0x8a, 0x48, 0x0b, 0xa2, 0xa7, 0x61, 0x93, 0x4e, 0x64, 0x71, 0x3f,
	0xc2, 0xec, 0x04, 0xda, 0xdf, 0xeb, 0x1f, 0x91, 0xdd, 0x10, 0xd0, 0xd2, 0xe3, 0x3e, 0xef, 0xf3,
	0xce, 0x0e, 0x8c, 0x38, 0x0b, 0xce, 0x76, 0x50, 0xf7, 0xe6, 0x3b, 0xb0, 0xe2, 0x00, 0x75, 0xd5,
	0x91, 0x67, 0x2f, 0xdf, 0x1b, 0x1d, 0x1a, 0x0f, 0xda, 0x07, 0xd8, 0x10, 0x24, 0x49, 0x7d, 0xa1,
	0x63, 0xf0, 0x1d, 0x52, 0x15, 0xdf, 0x95, 0x76, 0xad, 0x27, 0xab, 0x58, 0x7b, 0x57, 0xc5, 0x01,
	0x30, 0x4c, 0x60, 0x52, 0x2e, 0xb4, 0x48, 0x61, 0x1f, 0x9c, 0x3f, 0x88, 0xd3, 0xdd, 0x1f, 0xe1,
	0xf9, 0xf1, 0x73, 0x25, 0xaf, 0xc4, 0x64, 0x94, 0xc0, 0x29, 0x8b, 0x45, 0x36, 0xcb, 0xca, 0x7c,
	0x79, 0x3c, 0xc2, 0x57, 0x65, 0x71, 0xfe, 0x93, 0x89, 0xe9, 0x6e, 0x5d, 0x96, 0x62, 0xfa, 0xaf,
	0x09, 0x1b, 0x2a, 0xae, 0x53, 0xf9, 0xe4, 0x6f, 0xf9, 0x83, 0xe4, 0xa5, 0x38, 0xea, 0xc9, 0x40,
	0x47, 0xda, 0x2a, 0xda, 0x16, 0x8b, 0x24, 0x89, 0x9e, 0xcc, 0xdb, 0x40, 0xe2, 0x12, 0x51, 0x08,
	0xd8, 0x78, 0xb7, 0x8e, 0xca, 0xcd, 0xb0, 0x44, 0x4f, 0x66, 0x35, 0xb2, 0x28, 0x11, 0xb2, 0xd2,
	0x0e, 0xd7, 0xd0, 0x6a, 0x83, 0xc5, 0xed, 0x20, 0x8d, 0xf0, 0x49, 0x1b, 0x94, 0xe7, 0x22, 0x67,
	0x6d, 0x11, 0x0c, 0xb6, 0x5c, 0xdc, 0xcd, 0xb2, 0x72, 0xb2, 0x3c, 0x8c, 0xe0, 0x05, 0x5b, 0x96,
	0x17, 0x42, 0x10, 0x32, 0x6d, 0x87, 0xf4, 0x3e, 0xa5, 0x79, 0x22, 0x31, 0xae, 0x0f, 0xd2, 0x09,
	0x16, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x14, 0xa0, 0xe7, 0x9c, 0x01, 0x00, 0x00,
}
