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
// source: ipv4_dhcpd_database.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_base_database

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

type Ipv4DhcpdDatabase_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdDatabase_KEYS) Reset()         { *m = Ipv4DhcpdDatabase_KEYS{} }
func (m *Ipv4DhcpdDatabase_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdDatabase_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdDatabase_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a9921a27a535460, []int{0}
}

func (m *Ipv4DhcpdDatabase_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdDatabase_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdDatabase_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdDatabase_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdDatabase_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdDatabase_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdDatabase_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdDatabase_KEYS.Size(m)
}
func (m *Ipv4DhcpdDatabase_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdDatabase_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdDatabase_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdDatabase_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

type Ipv4DhcpdDatabase struct {
	Configured                             bool     `protobuf:"varint,50,opt,name=configured,proto3" json:"configured,omitempty"`
	Version                                uint32   `protobuf:"varint,51,opt,name=version,proto3" json:"version,omitempty"`
	FullFileWriteInterval                  uint32   `protobuf:"varint,52,opt,name=full_file_write_interval,json=fullFileWriteInterval,proto3" json:"full_file_write_interval,omitempty"`
	LastFullWriteFileName                  string   `protobuf:"bytes,53,opt,name=last_full_write_file_name,json=lastFullWriteFileName,proto3" json:"last_full_write_file_name,omitempty"`
	LastFullWriteTime                      uint32   `protobuf:"varint,54,opt,name=last_full_write_time,json=lastFullWriteTime,proto3" json:"last_full_write_time,omitempty"`
	FullFileWriteCount                     uint32   `protobuf:"varint,55,opt,name=full_file_write_count,json=fullFileWriteCount,proto3" json:"full_file_write_count,omitempty"`
	FailedFullFileWriteCount               uint32   `protobuf:"varint,56,opt,name=failed_full_file_write_count,json=failedFullFileWriteCount,proto3" json:"failed_full_file_write_count,omitempty"`
	FullFileRecordCount                    uint32   `protobuf:"varint,57,opt,name=full_file_record_count,json=fullFileRecordCount,proto3" json:"full_file_record_count,omitempty"`
	LastFullFileWriteErrorTimestamp        uint32   `protobuf:"varint,58,opt,name=last_full_file_write_error_timestamp,json=lastFullFileWriteErrorTimestamp,proto3" json:"last_full_file_write_error_timestamp,omitempty"`
	IncrementalFileWriteInterval           uint32   `protobuf:"varint,59,opt,name=incremental_file_write_interval,json=incrementalFileWriteInterval,proto3" json:"incremental_file_write_interval,omitempty"`
	LastIncrementalWriteFileName           string   `protobuf:"bytes,60,opt,name=last_incremental_write_file_name,json=lastIncrementalWriteFileName,proto3" json:"last_incremental_write_file_name,omitempty"`
	LastIncrementalWriteTime               uint32   `protobuf:"varint,61,opt,name=last_incremental_write_time,json=lastIncrementalWriteTime,proto3" json:"last_incremental_write_time,omitempty"`
	IncrementalFileWriteCount              uint32   `protobuf:"varint,62,opt,name=incremental_file_write_count,json=incrementalFileWriteCount,proto3" json:"incremental_file_write_count,omitempty"`
	FailedIncrementalFileWriteCount        uint32   `protobuf:"varint,63,opt,name=failed_incremental_file_write_count,json=failedIncrementalFileWriteCount,proto3" json:"failed_incremental_file_write_count,omitempty"`
	IncrementalFileRecordCount             uint32   `protobuf:"varint,64,opt,name=incremental_file_record_count,json=incrementalFileRecordCount,proto3" json:"incremental_file_record_count,omitempty"`
	LastIncrementalFileWriteErrorTimestamp uint32   `protobuf:"varint,65,opt,name=last_incremental_file_write_error_timestamp,json=lastIncrementalFileWriteErrorTimestamp,proto3" json:"last_incremental_file_write_error_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral                   struct{} `json:"-"`
	XXX_unrecognized                       []byte   `json:"-"`
	XXX_sizecache                          int32    `json:"-"`
}

func (m *Ipv4DhcpdDatabase) Reset()         { *m = Ipv4DhcpdDatabase{} }
func (m *Ipv4DhcpdDatabase) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdDatabase) ProtoMessage()    {}
func (*Ipv4DhcpdDatabase) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a9921a27a535460, []int{1}
}

func (m *Ipv4DhcpdDatabase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdDatabase.Unmarshal(m, b)
}
func (m *Ipv4DhcpdDatabase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdDatabase.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdDatabase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdDatabase.Merge(m, src)
}
func (m *Ipv4DhcpdDatabase) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdDatabase.Size(m)
}
func (m *Ipv4DhcpdDatabase) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdDatabase.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdDatabase proto.InternalMessageInfo

func (m *Ipv4DhcpdDatabase) GetConfigured() bool {
	if m != nil {
		return m.Configured
	}
	return false
}

func (m *Ipv4DhcpdDatabase) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetFullFileWriteInterval() uint32 {
	if m != nil {
		return m.FullFileWriteInterval
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetLastFullWriteFileName() string {
	if m != nil {
		return m.LastFullWriteFileName
	}
	return ""
}

func (m *Ipv4DhcpdDatabase) GetLastFullWriteTime() uint32 {
	if m != nil {
		return m.LastFullWriteTime
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetFullFileWriteCount() uint32 {
	if m != nil {
		return m.FullFileWriteCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetFailedFullFileWriteCount() uint32 {
	if m != nil {
		return m.FailedFullFileWriteCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetFullFileRecordCount() uint32 {
	if m != nil {
		return m.FullFileRecordCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetLastFullFileWriteErrorTimestamp() uint32 {
	if m != nil {
		return m.LastFullFileWriteErrorTimestamp
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetIncrementalFileWriteInterval() uint32 {
	if m != nil {
		return m.IncrementalFileWriteInterval
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetLastIncrementalWriteFileName() string {
	if m != nil {
		return m.LastIncrementalWriteFileName
	}
	return ""
}

func (m *Ipv4DhcpdDatabase) GetLastIncrementalWriteTime() uint32 {
	if m != nil {
		return m.LastIncrementalWriteTime
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetIncrementalFileWriteCount() uint32 {
	if m != nil {
		return m.IncrementalFileWriteCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetFailedIncrementalFileWriteCount() uint32 {
	if m != nil {
		return m.FailedIncrementalFileWriteCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetIncrementalFileRecordCount() uint32 {
	if m != nil {
		return m.IncrementalFileRecordCount
	}
	return 0
}

func (m *Ipv4DhcpdDatabase) GetLastIncrementalFileWriteErrorTimestamp() uint32 {
	if m != nil {
		return m.LastIncrementalFileWriteErrorTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4DhcpdDatabase_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.database.ipv4_dhcpd_database_KEYS")
	proto.RegisterType((*Ipv4DhcpdDatabase)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.base.database.ipv4_dhcpd_database")
}

func init() { proto.RegisterFile("ipv4_dhcpd_database.proto", fileDescriptor_7a9921a27a535460) }

var fileDescriptor_7a9921a27a535460 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdb, 0x6b, 0x13, 0x41,
	0x14, 0xc6, 0xc9, 0x4b, 0xab, 0x07, 0x7c, 0x70, 0x6a, 0xca, 0x44, 0xa3, 0x0d, 0x55, 0x24, 0x20,
	0x44, 0x6c, 0xaa, 0xad, 0x97, 0xde, 0x90, 0x04, 0x82, 0x97, 0x87, 0x58, 0x10, 0xf1, 0x61, 0x98,
	0xee, 0x9e, 0xe8, 0xc0, 0xec, 0xce, 0x32, 0x3b, 0x89, 0xfe, 0xc7, 0xfe, 0x1b, 0x32, 0x67, 0xb2,
	0xcd, 0x66, 0xb3, 0x9b, 0x97, 0xc0, 0xe4, 0x7c, 0xbf, 0xef, 0x5c, 0x3e, 0x58, 0xe8, 0xa8, 0x6c,
	0x71, 0x2c, 0xe2, 0xdf, 0x51, 0x16, 0x8b, 0x58, 0x3a, 0x79, 0x23, 0x73, 0x1c, 0x64, 0xd6, 0x38,
	0xc3, 0x2e, 0x23, 0x95, 0x47, 0x46, 0x28, 0x93, 0x8b, 0xbf, 0x56, 0x94, 0x74, 0x26, 0x43, 0x3b,
	0x58, 0xbd, 0x07, 0xa9, 0x89, 0x31, 0xa7, 0xdf, 0x01, 0xe1, 0x85, 0xcf, 0xe1, 0x11, 0xf0, 0x1a,
	0x7b, 0xf1, 0x69, 0xf4, 0xe3, 0x1b, 0xdb, 0x87, 0x1d, 0x4f, 0xa8, 0x98, 0xb7, 0x7a, 0xad, 0xfe,
	0xdd, 0xe9, 0xf2, 0x75, 0xf8, 0x6f, 0x17, 0xf6, 0x6a, 0x20, 0xf6, 0x04, 0x20, 0x32, 0xe9, 0x4c,
	0xfd, 0x9a, 0x5b, 0x8c, 0xf9, 0x51, 0xaf, 0xd5, 0xbf, 0x33, 0x2d, 0xfd, 0xc3, 0x38, 0xec, 0x2e,
	0xd0, 0xe6, 0xca, 0xa4, 0x7c, 0xd8, 0x6b, 0xf5, 0xef, 0x4d, 0x8b, 0x27, 0x3b, 0x01, 0x3e, 0x9b,
	0x6b, 0x2d, 0x66, 0x4a, 0xa3, 0xf8, 0x63, 0x95, 0x43, 0xa1, 0x52, 0x87, 0x76, 0x21, 0x35, 0x3f,
	0x26, 0x69, 0xdb, 0xd7, 0xc7, 0x4a, 0xe3, 0x77, 0x5f, 0x9d, 0x2c, 0x8b, 0xec, 0x14, 0x3a, 0x5a,
	0xe6, 0x4e, 0x10, 0x1d, 0x40, 0xf2, 0x48, 0x65, 0x82, 0xfc, 0x35, 0x4d, 0xdd, 0xf6, 0x82, 0xf1,
	0x5c, 0x6b, 0x22, 0xbd, 0xc5, 0x57, 0x99, 0x20, 0x7b, 0x09, 0x0f, 0xaa, 0xa4, 0x53, 0x09, 0xf2,
	0x37, 0xd4, 0xee, 0xfe, 0x1a, 0x74, 0xad, 0x12, 0x64, 0xaf, 0xa0, 0x5d, 0x9d, 0x31, 0x32, 0xf3,
	0xd4, 0xf1, 0x13, 0x22, 0xd8, 0xda, 0x80, 0x1f, 0x7d, 0x85, 0x9d, 0x43, 0x77, 0x26, 0x95, 0xc6,
	0x58, 0xd4, 0x93, 0xa7, 0x44, 0xf2, 0xa0, 0x19, 0x6f, 0xf2, 0x43, 0xd8, 0x5f, 0x81, 0x16, 0x23,
	0x63, 0xe3, 0x25, 0xf9, 0x96, 0xc8, 0xbd, 0xa2, 0xe7, 0x94, 0x6a, 0x01, 0xfa, 0x02, 0xcf, 0x56,
	0x8b, 0x95, 0x5a, 0xa2, 0xb5, 0xc6, 0xd2, 0x8e, 0xb9, 0x93, 0x49, 0xc6, 0xdf, 0x91, 0xc5, 0x41,
	0xb1, 0xe8, 0x6d, 0xeb, 0x91, 0xd7, 0x5d, 0x17, 0x32, 0x36, 0x82, 0x03, 0x95, 0x46, 0x16, 0x13,
	0x4c, 0x9d, 0xac, 0x4f, 0xe8, 0x3d, 0x39, 0x75, 0x4b, 0xb2, 0xcd, 0xa0, 0xc6, 0xd0, 0xa3, 0xa9,
	0xca, 0x5e, 0xd5, 0xbc, 0x3e, 0x50, 0x5e, 0x5d, 0xaf, 0x9b, 0xac, 0x64, 0xeb, 0xb1, 0x9d, 0xc1,
	0xa3, 0x06, 0x1f, 0x4a, 0xef, 0x2c, 0x5c, 0xb4, 0xce, 0x82, 0x42, 0xbc, 0x80, 0x6e, 0xc3, 0x36,
	0xe1, 0xae, 0xe7, 0xc4, 0x77, 0xea, 0x56, 0x09, 0xd7, 0xfd, 0x0c, 0x4f, 0x97, 0x91, 0x6e, 0xf5,
	0xb9, 0x08, 0xc7, 0x0d, 0xd2, 0x49, 0xa3, 0xdb, 0x15, 0x3c, 0xde, 0xb0, 0x59, 0xcb, 0xf9, 0x92,
	0x7c, 0x1e, 0x56, 0xe6, 0x29, 0xc7, 0xfd, 0x13, 0x5e, 0x6c, 0x1c, 0x64, 0x4b, 0xea, 0x57, 0x64,
	0xf8, 0xbc, 0x72, 0xa0, 0x86, 0xf0, 0x6f, 0x76, 0xe8, 0x33, 0x33, 0xfc, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x65, 0xa4, 0x1d, 0xbc, 0x83, 0x04, 0x00, 0x00,
}