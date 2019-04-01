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
	LspId                           string   `protobuf:"bytes,1,opt,name=lsp_id,json=lspId,proto3" json:"lsp_id,omitempty"`
	LocalLspFlag                    bool     `protobuf:"varint,2,opt,name=local_lsp_flag,json=localLspFlag,proto3" json:"local_lsp_flag,omitempty"`
	LspActiveFlag                   bool     `protobuf:"varint,3,opt,name=lsp_active_flag,json=lspActiveFlag,proto3" json:"lsp_active_flag,omitempty"`
	LspHoldtime                     uint32   `protobuf:"varint,4,opt,name=lsp_holdtime,json=lspHoldtime,proto3" json:"lsp_holdtime,omitempty"`
	LspSequenceNumber               uint32   `protobuf:"varint,5,opt,name=lsp_sequence_number,json=lspSequenceNumber,proto3" json:"lsp_sequence_number,omitempty"`
	LspChecksum                     uint32   `protobuf:"varint,6,opt,name=lsp_checksum,json=lspChecksum,proto3" json:"lsp_checksum,omitempty"`
	LspPartitionRepairSupportedFlag bool     `protobuf:"varint,7,opt,name=lsp_partition_repair_supported_flag,json=lspPartitionRepairSupportedFlag,proto3" json:"lsp_partition_repair_supported_flag,omitempty"`
	LspAttachedFlag                 bool     `protobuf:"varint,8,opt,name=lsp_attached_flag,json=lspAttachedFlag,proto3" json:"lsp_attached_flag,omitempty"`
	LspOverloadedFlag               bool     `protobuf:"varint,9,opt,name=lsp_overloaded_flag,json=lspOverloadedFlag,proto3" json:"lsp_overloaded_flag,omitempty"`
	LspNonv1AFlag                   uint32   `protobuf:"varint,10,opt,name=lsp_nonv1a_flag,json=lspNonv1aFlag,proto3" json:"lsp_nonv1a_flag,omitempty"`
	LspLevel                        string   `protobuf:"bytes,11,opt,name=lsp_level,json=lspLevel,proto3" json:"lsp_level,omitempty"`
	LspLength                       uint32   `protobuf:"varint,12,opt,name=lsp_length,json=lspLength,proto3" json:"lsp_length,omitempty"`
	MaxAreaAddresses                uint32   `protobuf:"varint,13,opt,name=max_area_addresses,json=maxAreaAddresses,proto3" json:"max_area_addresses,omitempty"`
	IdLength                        uint32   `protobuf:"varint,14,opt,name=id_length,json=idLength,proto3" json:"id_length,omitempty"`
	Version                         uint32   `protobuf:"varint,15,opt,name=version,proto3" json:"version,omitempty"`
	Version2                        uint32   `protobuf:"varint,16,opt,name=version2,proto3" json:"version2,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
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

func (m *IsisShLspHeader) GetLspPartitionRepairSupportedFlag() bool {
	if m != nil {
		return m.LspPartitionRepairSupportedFlag
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

func (m *IsisShLspHeader) GetMaxAreaAddresses() uint32 {
	if m != nil {
		return m.MaxAreaAddresses
	}
	return 0
}

func (m *IsisShLspHeader) GetIdLength() uint32 {
	if m != nil {
		return m.IdLength
	}
	return 0
}

func (m *IsisShLspHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *IsisShLspHeader) GetVersion2() uint32 {
	if m != nil {
		return m.Version2
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
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xdb, 0x6e, 0x13, 0x3d,
	0x10, 0x56, 0x7a, 0x4c, 0x9c, 0xa4, 0x07, 0xff, 0xa7, 0xd5, 0x8f, 0x0a, 0x6d, 0x8a, 0x00, 0x21,
	0xb4, 0x12, 0xe1, 0x09, 0x22, 0x28, 0x02, 0x11, 0x5a, 0xb4, 0x11, 0x17, 0x70, 0x63, 0x39, 0xeb,
	0x69, 0xb2, 0xc2, 0x6b, 0x1b, 0xdb, 0x49, 0xdb, 0x5b, 0x10, 0x17, 0x3c, 0x05, 0x4f, 0xc6, 0xab,
	0x20, 0xe4, 0xd9, 0x75, 0x13, 0x10, 0xbd, 0xa3, 0x5c, 0x65, 0xe7, 0x9b, 0xcf, 0xdf, 0x1c, 0x3c,
	0xe3, 0x90, 0x9b, 0x85, 0x2b, 0x1c, 0x73, 0x53, 0x26, 0x9d, 0x61, 0x62, 0xcc, 0xa4, 0x9e, 0x30,
	0xe7, 0xed, 0x2c, 0x4f, 0x8d, 0xd5, 0x5e, 0xd3, 0x97, 0x79, 0xe1, 0x72, 0xcd, 0x0a, 0xed, 0xd8,
	0xb9, 0x65, 0xb9, 0x54, 0x8e, 0xe1, 0x09, 0x6d, 0xc0, 0xa6, 0xe1, 0x2b, 0x2d, 0x94, 0xf3, 0x5c,
	0xe5, 0xb0, 0xf8, 0x4a, 0x25, 0xcc, 0x41, 0xba, 0xea, 0x27, 0x15, 0xdc, 0xf3, 0x31, 0x77, 0x10,
	0xa4, 0x7b, 0x6f, 0xc9, 0xde, 0x55, 0x01, 0xd9, 0x8b, 0xa3, 0x37, 0x23, 0x7a, 0x48, 0xba, 0x51,
	0x86, 0x29, 0x5e, 0x42, 0xd2, 0xd8, 0x6f, 0xdc, 0x6b, 0x65, 0x9d, 0x08, 0x1e, 0xf3, 0x12, 0xe8,
	0xdf, 0x64, 0x1d, 0xb5, 0x93, 0x15, 0x74, 0x56, 0x46, 0xef, 0x35, 0xf9, 0x37, 0x6a, 0xfb, 0xa2,
	0x04, 0xe7, 0x79, 0x69, 0x98, 0xbf, 0x30, 0x40, 0x13, 0xb2, 0xe9, 0x20, 0xd7, 0x4a, 0x38, 0x94,
	0xeb, 0x66, 0xd1, 0xa4, 0x07, 0xa4, 0xa3, 0xb8, 0xd2, 0x2c, 0xba, 0x57, 0xd0, 0xdd, 0x0e, 0xd8,
	0xa8, 0x82, 0x7a, 0x5f, 0x1a, 0xe4, 0xbf, 0xa8, 0x3b, 0x01, 0x05, 0xb6, 0xc8, 0x31, 0x69, 0x50,
	0x9e, 0x7e, 0x6c, 0x90, 0xd6, 0x65, 0x2c, 0xd4, 0x6e, 0xf7, 0x21, 0xfd, 0xad, 0x2d, 0x4b, 0x7f,
	0x5d, 0x53, 0xb6, 0x88, 0xdb, 0xfb, 0xb6, 0x46, 0xe8, 0x72, 0x57, 0xa7, 0xc0, 0x05, 0x58, 0xfa,
	0x0f, 0xd9, 0x08, 0x56, 0x21, 0xea, 0x1e, 0xae, 0x4b, 0x67, 0x9e, 0x0b, 0x7a, 0x9b, 0x6c, 0x49,
	0x9d, 0x73, 0x89, 0xd4, 0x53, 0xc9, 0x27, 0x58, 0x74, 0x33, 0xeb, 0x20, 0x3a, 0x74, 0xe6, 0xa9,
	0xe4, 0x13, 0x7a, 0x87, 0x6c, 0x07, 0x3f, 0xcf, 0x7d, 0x31, 0x87, 0x8a, 0xb6, 0x8a, 0xb4, 0xae,
	0x74, 0x66, 0x80, 0x28, 0xf2, 0x0e, 0x48, 0x07, 0x43, 0x6a, 0x29, 0x42, 0x42, 0xc9, 0x5a, 0xd5,
	0x40, 0xe9, 0xcc, 0xb3, 0x1a, 0xa2, 0x29, 0xf9, 0x2b, 0x50, 0x1c, 0xbc, 0x9f, 0x01, 0x5e, 0xeb,
	0xac, 0x1c, 0x83, 0x4d, 0xd6, 0x91, 0xb9, 0x2b, 0x9d, 0x19, 0xd5, 0x9e, 0x63, 0x74, 0x44, 0xc9,
	0x7c, 0x0a, 0xf9, 0x3b, 0x37, 0x2b, 0x93, 0x8d, 0x4b, 0xc9, 0xc7, 0x35, 0x44, 0x87, 0xe4, 0x30,
	0x50, 0x0c, 0xb7, 0xbe, 0xf0, 0x85, 0x56, 0xcc, 0x82, 0xe1, 0x85, 0x65, 0x6e, 0x66, 0x8c, 0xb6,
	0x1e, 0x44, 0x95, 0xf1, 0x26, 0x66, 0x7c, 0x4b, 0x3a, 0xf3, 0x2a, 0x32, 0x33, 0x24, 0x8e, 0x22,
	0x0f, 0x6b, 0xb8, 0x4f, 0x76, 0xb1, 0x56, 0xef, 0x79, 0x3e, 0x8d, 0x67, 0x9b, 0x78, 0x36, 0x34,
	0x61, 0x50, 0xe3, 0xc8, 0xad, 0x8b, 0xd1, 0x73, 0xb0, 0x52, 0x73, 0x11, 0xd9, 0x2d, 0x64, 0x07,
	0x99, 0x93, 0x4b, 0xcf, 0x72, 0x1f, 0x95, 0x56, 0xf3, 0x87, 0xbc, 0xe2, 0x12, 0xac, 0x27, 0xf4,
	0xf1, 0x18, 0x51, 0xe4, 0xdd, 0x20, 0xad, 0xc0, 0xab, 0xc6, 0xba, 0x8d, 0xf7, 0xd5, 0x94, 0xce,
	0x0c, 0x83, 0x4d, 0xf7, 0x08, 0xa9, 0x9c, 0x6a, 0xe2, 0xa7, 0x49, 0x07, 0xcf, 0xb7, 0xd0, 0x1b,
	0x00, 0xfa, 0x80, 0xd0, 0x92, 0x9f, 0x33, 0x6e, 0x81, 0x33, 0x2e, 0x84, 0x05, 0xe7, 0xc0, 0x25,
	0x5d, 0xa4, 0xed, 0x94, 0xfc, 0x7c, 0x60, 0x81, 0x0f, 0x22, 0x1e, 0x22, 0x15, 0x22, 0x6a, 0x6d,
	0x21, 0xa9, 0x59, 0x88, 0x5a, 0x2a, 0x21, 0x9b, 0x73, 0xb0, 0xae, 0xd0, 0x2a, 0xd9, 0xae, 0x36,
	0xa5, 0x36, 0xe9, 0xff, 0xa4, 0x59, 0x7f, 0xf6, 0x93, 0x9d, 0xea, 0x54, 0xb4, 0x7b, 0x5f, 0x57,
	0x17, 0xab, 0xb7, 0xb4, 0xd6, 0x61, 0x43, 0x3e, 0x37, 0x48, 0x27, 0x6e, 0x4d, 0x18, 0xeb, 0x7a,
	0x49, 0x4e, 0xaf, 0x69, 0x49, 0x7e, 0x5a, 0xd0, 0xac, 0x5d, 0x03, 0x4f, 0xb8, 0xe7, 0xf4, 0x2e,
	0xde, 0x85, 0x18, 0x63, 0x10, 0x1e, 0xa6, 0xa1, 0x7e, 0x40, 0xb6, 0x10, 0x3e, 0x89, 0x28, 0xfd,
	0xd4, 0x20, 0x5d, 0x05, 0x67, 0x58, 0x0b, 0x28, 0x6f, 0x2f, 0x70, 0xf6, 0xdb, 0x7d, 0x7e, 0x4d,
	0x59, 0x2f, 0x96, 0x36, 0x6b, 0x2b, 0x38, 0x1b, 0x3a, 0x73, 0x14, 0xa2, 0x62, 0x1e, 0x5a, 0x8a,
	0xa5, 0x3c, 0xd6, 0xfe, 0x58, 0x1e, 0x5a, 0x8a, 0x98, 0x47, 0x78, 0x02, 0x93, 0xab, 0x9e, 0x6d,
	0xfa, 0xa1, 0x41, 0x5a, 0x75, 0xbb, 0xed, 0x45, 0xd2, 0xdf, 0x5f, 0xbd, 0xc6, 0x37, 0xf0, 0xc7,
	0xe1, 0xca, 0x9a, 0x52, 0x4f, 0x30, 0xc3, 0xf1, 0x06, 0xfe, 0x5b, 0x3d, 0xfa, 0x1e, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x3a, 0xa9, 0x23, 0xcf, 0x06, 0x00, 0x00,
}