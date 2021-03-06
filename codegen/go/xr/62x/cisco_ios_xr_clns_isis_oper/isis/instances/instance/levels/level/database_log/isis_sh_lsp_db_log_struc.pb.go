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
// source: isis_sh_lsp_db_log_struc.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_levels_level_database_log

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

type IsisShLspDbLogStruc_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	Level                string   `protobuf:"bytes,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShLspDbLogStruc_KEYS) Reset()         { *m = IsisShLspDbLogStruc_KEYS{} }
func (m *IsisShLspDbLogStruc_KEYS) String() string { return proto.CompactTextString(m) }
func (*IsisShLspDbLogStruc_KEYS) ProtoMessage()    {}
func (*IsisShLspDbLogStruc_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{0}
}

func (m *IsisShLspDbLogStruc_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLspDbLogStruc_KEYS.Unmarshal(m, b)
}
func (m *IsisShLspDbLogStruc_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLspDbLogStruc_KEYS.Marshal(b, m, deterministic)
}
func (m *IsisShLspDbLogStruc_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLspDbLogStruc_KEYS.Merge(m, src)
}
func (m *IsisShLspDbLogStruc_KEYS) XXX_Size() int {
	return xxx_messageInfo_IsisShLspDbLogStruc_KEYS.Size(m)
}
func (m *IsisShLspDbLogStruc_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLspDbLogStruc_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLspDbLogStruc_KEYS proto.InternalMessageInfo

func (m *IsisShLspDbLogStruc_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShLspDbLogStruc_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type IsisShTimestampType struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	NanoSeconds          uint32   `protobuf:"varint,2,opt,name=nano_seconds,json=nanoSeconds,proto3" json:"nano_seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsisShTimestampType) Reset()         { *m = IsisShTimestampType{} }
func (m *IsisShTimestampType) String() string { return proto.CompactTextString(m) }
func (*IsisShTimestampType) ProtoMessage()    {}
func (*IsisShTimestampType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{1}
}

func (m *IsisShTimestampType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShTimestampType.Unmarshal(m, b)
}
func (m *IsisShTimestampType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShTimestampType.Marshal(b, m, deterministic)
}
func (m *IsisShTimestampType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShTimestampType.Merge(m, src)
}
func (m *IsisShTimestampType) XXX_Size() int {
	return xxx_messageInfo_IsisShTimestampType.Size(m)
}
func (m *IsisShTimestampType) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShTimestampType.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShTimestampType proto.InternalMessageInfo

func (m *IsisShTimestampType) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *IsisShTimestampType) GetNanoSeconds() uint32 {
	if m != nil {
		return m.NanoSeconds
	}
	return 0
}

type IsisShGenericLogEnt struct {
	Timestamp            *IsisShTimestampType `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IsisShGenericLogEnt) Reset()         { *m = IsisShGenericLogEnt{} }
func (m *IsisShGenericLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShGenericLogEnt) ProtoMessage()    {}
func (*IsisShGenericLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{2}
}

func (m *IsisShGenericLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShGenericLogEnt.Unmarshal(m, b)
}
func (m *IsisShGenericLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShGenericLogEnt.Marshal(b, m, deterministic)
}
func (m *IsisShGenericLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShGenericLogEnt.Merge(m, src)
}
func (m *IsisShGenericLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShGenericLogEnt.Size(m)
}
func (m *IsisShGenericLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShGenericLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShGenericLogEnt proto.InternalMessageInfo

func (m *IsisShGenericLogEnt) GetTimestamp() *IsisShTimestampType {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type IsisShLspHeader struct {
	LspId                          string   `protobuf:"bytes,1,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	LocalLspFlag                   bool     `protobuf:"varint,2,opt,name=local_lsp_flag,json=localLspFlag,proto3" json:"local_lsp_flag,omitempty"`
	LspActiveFlag                  bool     `protobuf:"varint,3,opt,name=lsp_active_flag,json=lspActiveFlag,proto3" json:"lsp_active_flag,omitempty"`
	LspHoldtime                    uint32   `protobuf:"varint,4,opt,name=lsp_holdtime,json=lspHoldtime,proto3" json:"lsp_holdtime,omitempty"`
	LspSequenceNumber              uint32   `protobuf:"varint,5,opt,name=lsp_sequence_number,json=lspSequenceNumber,proto3" json:"lsp_sequence_number,omitempty"`
	LspChecksum                    uint32   `protobuf:"varint,6,opt,name=lsp_checksum,json=lspChecksum,proto3" json:"lsp_checksum,omitempty"`
	LspParitionRepairSupportedFlag bool     `protobuf:"varint,7,opt,name=lsp_parition_repair_supported_flag,json=lspParitionRepairSupportedFlag,proto3" json:"lsp_parition_repair_supported_flag,omitempty"`
	LspAttachedFlag                bool     `protobuf:"varint,8,opt,name=lsp_attached_flag,json=lspAttachedFlag,proto3" json:"lsp_attached_flag,omitempty"`
	LspOverloadedFlag              bool     `protobuf:"varint,9,opt,name=lsp_overloaded_flag,json=lspOverloadedFlag,proto3" json:"lsp_overloaded_flag,omitempty"`
	LspNonv1AFlag                  uint32   `protobuf:"varint,10,opt,name=lsp_nonv1a_flag,json=lspNonv1aFlag,proto3" json:"lsp_nonv1a_flag,omitempty"`
	LspLevel                       string   `protobuf:"bytes,11,opt,name=lsp_level,json=lspLevel,proto3" json:"lsp_level,omitempty"`
	LspLength                      uint32   `protobuf:"varint,12,opt,name=lsp_length,json=lspLength,proto3" json:"lsp_length,omitempty"`
	XXX_NoUnkeyedLiteral           struct{} `json:"-"`
	XXX_unrecognized               []byte   `json:"-"`
	XXX_sizecache                  int32    `json:"-"`
}

func (m *IsisShLspHeader) Reset()         { *m = IsisShLspHeader{} }
func (m *IsisShLspHeader) String() string { return proto.CompactTextString(m) }
func (*IsisShLspHeader) ProtoMessage()    {}
func (*IsisShLspHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{3}
}

func (m *IsisShLspHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLspHeader.Unmarshal(m, b)
}
func (m *IsisShLspHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLspHeader.Marshal(b, m, deterministic)
}
func (m *IsisShLspHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLspHeader.Merge(m, src)
}
func (m *IsisShLspHeader) XXX_Size() int {
	return xxx_messageInfo_IsisShLspHeader.Size(m)
}
func (m *IsisShLspHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLspHeader.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLspHeader proto.InternalMessageInfo

func (m *IsisShLspHeader) GetLspId() string {
	if m != nil {
		return m.LspId
	}
	return ""
}

func (m *IsisShLspHeader) GetLocalLspFlag() bool {
	if m != nil {
		return m.LocalLspFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspActiveFlag() bool {
	if m != nil {
		return m.LspActiveFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspHoldtime() uint32 {
	if m != nil {
		return m.LspHoldtime
	}
	return 0
}

func (m *IsisShLspHeader) GetLspSequenceNumber() uint32 {
	if m != nil {
		return m.LspSequenceNumber
	}
	return 0
}

func (m *IsisShLspHeader) GetLspChecksum() uint32 {
	if m != nil {
		return m.LspChecksum
	}
	return 0
}

func (m *IsisShLspHeader) GetLspParitionRepairSupportedFlag() bool {
	if m != nil {
		return m.LspParitionRepairSupportedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspAttachedFlag() bool {
	if m != nil {
		return m.LspAttachedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspOverloadedFlag() bool {
	if m != nil {
		return m.LspOverloadedFlag
	}
	return false
}

func (m *IsisShLspHeader) GetLspNonv1AFlag() uint32 {
	if m != nil {
		return m.LspNonv1AFlag
	}
	return 0
}

func (m *IsisShLspHeader) GetLspLevel() string {
	if m != nil {
		return m.LspLevel
	}
	return ""
}

func (m *IsisShLspHeader) GetLspLength() uint32 {
	if m != nil {
		return m.LspLength
	}
	return 0
}

type IsisShLspDbLogEnt struct {
	GenericData          *IsisShGenericLogEnt `protobuf:"bytes,1,opt,name=generic_data,json=genericData,proto3" json:"generic_data,omitempty"`
	LspdbOperation       string               `protobuf:"bytes,2,opt,name=lspdb_operation,json=lspdbOperation,proto3" json:"lspdb_operation,omitempty"`
	NewLspEntry          *IsisShLspHeader     `protobuf:"bytes,3,opt,name=new_lsp_entry,json=newLspEntry,proto3" json:"new_lsp_entry,omitempty"`
	OldLspEntry          *IsisShLspHeader     `protobuf:"bytes,4,opt,name=old_lsp_entry,json=oldLspEntry,proto3" json:"old_lsp_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IsisShLspDbLogEnt) Reset()         { *m = IsisShLspDbLogEnt{} }
func (m *IsisShLspDbLogEnt) String() string { return proto.CompactTextString(m) }
func (*IsisShLspDbLogEnt) ProtoMessage()    {}
func (*IsisShLspDbLogEnt) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{4}
}

func (m *IsisShLspDbLogEnt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLspDbLogEnt.Unmarshal(m, b)
}
func (m *IsisShLspDbLogEnt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLspDbLogEnt.Marshal(b, m, deterministic)
}
func (m *IsisShLspDbLogEnt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLspDbLogEnt.Merge(m, src)
}
func (m *IsisShLspDbLogEnt) XXX_Size() int {
	return xxx_messageInfo_IsisShLspDbLogEnt.Size(m)
}
func (m *IsisShLspDbLogEnt) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLspDbLogEnt.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLspDbLogEnt proto.InternalMessageInfo

func (m *IsisShLspDbLogEnt) GetGenericData() *IsisShGenericLogEnt {
	if m != nil {
		return m.GenericData
	}
	return nil
}

func (m *IsisShLspDbLogEnt) GetLspdbOperation() string {
	if m != nil {
		return m.LspdbOperation
	}
	return ""
}

func (m *IsisShLspDbLogEnt) GetNewLspEntry() *IsisShLspHeader {
	if m != nil {
		return m.NewLspEntry
	}
	return nil
}

func (m *IsisShLspDbLogEnt) GetOldLspEntry() *IsisShLspHeader {
	if m != nil {
		return m.OldLspEntry
	}
	return nil
}

type IsisShLspDbLogStruc struct {
	LogEntry             []*IsisShLspDbLogEnt `protobuf:"bytes,50,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *IsisShLspDbLogStruc) Reset()         { *m = IsisShLspDbLogStruc{} }
func (m *IsisShLspDbLogStruc) String() string { return proto.CompactTextString(m) }
func (*IsisShLspDbLogStruc) ProtoMessage()    {}
func (*IsisShLspDbLogStruc) Descriptor() ([]byte, []int) {
	return fileDescriptor_f73922a0bf04e811, []int{5}
}

func (m *IsisShLspDbLogStruc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsisShLspDbLogStruc.Unmarshal(m, b)
}
func (m *IsisShLspDbLogStruc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsisShLspDbLogStruc.Marshal(b, m, deterministic)
}
func (m *IsisShLspDbLogStruc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsisShLspDbLogStruc.Merge(m, src)
}
func (m *IsisShLspDbLogStruc) XXX_Size() int {
	return xxx_messageInfo_IsisShLspDbLogStruc.Size(m)
}
func (m *IsisShLspDbLogStruc) XXX_DiscardUnknown() {
	xxx_messageInfo_IsisShLspDbLogStruc.DiscardUnknown(m)
}

var xxx_messageInfo_IsisShLspDbLogStruc proto.InternalMessageInfo

func (m *IsisShLspDbLogStruc) GetLogEntry() []*IsisShLspDbLogEnt {
	if m != nil {
		return m.LogEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*IsisShLspDbLogStruc_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_lsp_db_log_struc_KEYS")
	proto.RegisterType((*IsisShTimestampType)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_timestamp_type")
	proto.RegisterType((*IsisShGenericLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_generic_log_ent")
	proto.RegisterType((*IsisShLspHeader)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_lsp_header")
	proto.RegisterType((*IsisShLspDbLogEnt)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_lsp_db_log_ent")
	proto.RegisterType((*IsisShLspDbLogStruc)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.levels.level.database_log.isis_sh_lsp_db_log_struc")
}

func init() { proto.RegisterFile("isis_sh_lsp_db_log_struc.proto", fileDescriptor_f73922a0bf04e811) }

var fileDescriptor_f73922a0bf04e811 = []byte{
	// 638 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xd6, 0x36, 0xfd, 0x48, 0x9c, 0xa4, 0xaf, 0xea, 0x97, 0x8f, 0x95, 0x50, 0xab, 0x36, 0x20,
	0x40, 0x1c, 0x56, 0xa2, 0xfc, 0x82, 0x0a, 0x8a, 0xf8, 0x28, 0x2d, 0xda, 0x88, 0x03, 0x5c, 0x2c,
	0x67, 0x77, 0x9a, 0xac, 0x70, 0x6c, 0xe3, 0x71, 0x52, 0x7a, 0x05, 0x71, 0xe0, 0x57, 0xf0, 0xcb,
	0xf8, 0x25, 0x5c, 0x90, 0x67, 0xd7, 0x4d, 0x40, 0xf4, 0x46, 0x39, 0x65, 0xfd, 0xcc, 0xe3, 0x67,
	0x9e, 0x19, 0x8f, 0x1d, 0xb6, 0x53, 0x61, 0x85, 0x02, 0x27, 0x42, 0xa1, 0x15, 0xe5, 0x48, 0x28,
	0x33, 0x16, 0xe8, 0xdd, 0xac, 0xc8, 0xac, 0x33, 0xde, 0xf0, 0x57, 0x45, 0x85, 0x85, 0x11, 0x95,
	0x41, 0xf1, 0xd1, 0x89, 0x42, 0x69, 0x14, 0xb4, 0xc3, 0x58, 0x70, 0x59, 0xf8, 0xca, 0x2a, 0x8d,
	0x5e, 0xea, 0x02, 0x16, 0x5f, 0x99, 0x82, 0x39, 0x28, 0xac, 0x7f, 0xb2, 0x52, 0x7a, 0x39, 0x92,
	0x08, 0x41, 0x7a, 0xf0, 0x8e, 0x6d, 0x5f, 0x96, 0x50, 0xbc, 0x3c, 0x7c, 0x3b, 0xe4, 0xb7, 0x59,
	0x3f, 0xca, 0x08, 0x2d, 0xa7, 0x90, 0x26, 0xbb, 0xc9, 0xfd, 0x4e, 0xde, 0x8b, 0xe0, 0xb1, 0x9c,
	0x02, 0xbf, 0xc6, 0xd6, 0x48, 0x3b, 0x5d, 0xa1, 0x60, 0xbd, 0x18, 0xbc, 0x61, 0x37, 0xa2, 0xb6,
	0xaf, 0xa6, 0x80, 0x5e, 0x4e, 0xad, 0xf0, 0xe7, 0x16, 0x78, 0xca, 0x36, 0x10, 0x0a, 0xa3, 0x4b,
	0x24, 0xb9, 0x7e, 0x1e, 0x97, 0x7c, 0x8f, 0xf5, 0xb4, 0xd4, 0x46, 0xc4, 0xf0, 0x0a, 0x85, 0xbb,
	0x01, 0x1b, 0xd6, 0xd0, 0xe0, 0x5b, 0xc2, 0x6e, 0x46, 0xdd, 0x31, 0x68, 0x70, 0x55, 0x41, 0xa6,
	0x41, 0x7b, 0xfe, 0x39, 0x61, 0x9d, 0x8b, 0x5c, 0xa4, 0xdd, 0xdd, 0x87, 0xec, 0xaf, 0xb6, 0x2c,
	0xfb, 0x73, 0x4d, 0xf9, 0x22, 0xef, 0xe0, 0x47, 0x8b, 0xf1, 0xe5, 0xae, 0x4e, 0x40, 0x96, 0xe0,
	0xf8, 0x75, 0xb6, 0x1e, 0x56, 0x55, 0xd9, 0xf4, 0x70, 0x4d, 0xa1, 0x7d, 0x5e, 0xf2, 0x3b, 0x6c,
	0x53, 0x99, 0x42, 0x2a, 0xa2, 0x9e, 0x2a, 0x39, 0xa6, 0xa2, 0xdb, 0x79, 0x8f, 0xd0, 0x23, 0xb4,
	0x4f, 0x95, 0x1c, 0xf3, 0xbb, 0xec, 0xbf, 0x10, 0x97, 0x85, 0xaf, 0xe6, 0x50, 0xd3, 0x5a, 0x44,
	0xeb, 0x2b, 0xb4, 0x07, 0x84, 0x12, 0x6f, 0x8f, 0xf5, 0x28, 0xa5, 0x51, 0x65, 0x30, 0x94, 0xae,
	0xd6, 0x0d, 0x54, 0x68, 0x9f, 0x35, 0x10, 0xcf, 0xd8, 0xff, 0x81, 0x82, 0xf0, 0x61, 0x06, 0x74,
	0xac, 0xb3, 0xe9, 0x08, 0x5c, 0xba, 0x46, 0xcc, 0x2d, 0x85, 0x76, 0xd8, 0x44, 0x8e, 0x29, 0x10,
	0x25, 0x8b, 0x09, 0x14, 0xef, 0x71, 0x36, 0x4d, 0xd7, 0x2f, 0x24, 0x1f, 0x37, 0x10, 0x7f, 0xc1,
	0x06, 0x81, 0x62, 0xa5, 0xab, 0x7c, 0x65, 0xb4, 0x70, 0x60, 0x65, 0xe5, 0x04, 0xce, 0xac, 0x35,
	0xce, 0x43, 0x59, 0x1b, 0xde, 0x20, 0xc3, 0x3b, 0x0a, 0xed, 0xeb, 0x86, 0x98, 0x13, 0x6f, 0x18,
	0x69, 0x54, 0xc1, 0x03, 0xb6, 0x45, 0x95, 0x7a, 0x2f, 0x8b, 0x49, 0xdc, 0xda, 0xa6, 0xad, 0xa1,
	0x05, 0x07, 0x0d, 0x4e, 0xdc, 0xa6, 0x14, 0x33, 0x07, 0xa7, 0x8c, 0x2c, 0x23, 0xbb, 0x43, 0xec,
	0x20, 0x73, 0x72, 0x11, 0x59, 0xee, 0xa2, 0x36, 0x7a, 0xfe, 0x50, 0xd6, 0x5c, 0x46, 0xd5, 0x84,
	0x2e, 0x1e, 0x13, 0x4a, 0xbc, 0x5b, 0xac, 0x13, 0x78, 0xf5, 0x50, 0x77, 0xe9, 0xb4, 0xda, 0x0a,
	0xed, 0x51, 0x58, 0xf3, 0x6d, 0xc6, 0xea, 0xa0, 0x1e, 0xfb, 0x49, 0xda, 0xa3, 0xfd, 0x1d, 0x8a,
	0x06, 0x60, 0xf0, 0xbd, 0xb5, 0x98, 0xfb, 0xa5, 0x3b, 0x15, 0xc6, 0xf3, 0x6b, 0xc2, 0x7a, 0x71,
	0x64, 0xc3, 0x4c, 0x35, 0x13, 0x7a, 0x7a, 0x45, 0x13, 0xfa, 0xdb, 0xed, 0xc8, 0xbb, 0x0d, 0xf0,
	0x44, 0x7a, 0xc9, 0xef, 0x51, 0x2b, 0xca, 0x11, 0x25, 0x91, 0xe1, 0x30, 0x9a, 0xdb, 0xbb, 0x49,
	0xf0, 0x49, 0x44, 0xf9, 0x97, 0x84, 0xf5, 0x35, 0x9c, 0x51, 0x2d, 0xa0, 0xbd, 0x3b, 0xa7, 0xc1,
	0xeb, 0xee, 0xcb, 0x2b, 0x72, 0xbd, 0xb8, 0x31, 0x79, 0x57, 0xc3, 0xd9, 0x11, 0xda, 0xc3, 0x90,
	0x95, 0x7c, 0x18, 0x55, 0x2e, 0xf9, 0x58, 0xfd, 0x67, 0x3e, 0x8c, 0x2a, 0xa3, 0x8f, 0xf0, 0xfe,
	0xa4, 0x97, 0xbd, 0x99, 0xfc, 0x53, 0xc2, 0x3a, 0x4d, 0xbb, 0xdd, 0x79, 0xba, 0xbf, 0xdb, 0xba,
	0xc2, 0x07, 0xe8, 0xd7, 0xe1, 0xca, 0xdb, 0xca, 0x8c, 0xc9, 0xe1, 0x68, 0x9d, 0xfe, 0x2a, 0x1e,
	0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x59, 0xea, 0x60, 0xf2, 0x4c, 0x06, 0x00, 0x00,
}
