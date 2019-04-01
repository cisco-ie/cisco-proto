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
// source: rcmd_show_memory.proto

package cisco_ios_xr_infra_rcmd_oper_rcmd_memory

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

type RcmdShowMemory_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowMemory_KEYS) Reset()         { *m = RcmdShowMemory_KEYS{} }
func (m *RcmdShowMemory_KEYS) String() string { return proto.CompactTextString(m) }
func (*RcmdShowMemory_KEYS) ProtoMessage()    {}
func (*RcmdShowMemory_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1162942eb0a682db, []int{0}
}

func (m *RcmdShowMemory_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowMemory_KEYS.Unmarshal(m, b)
}
func (m *RcmdShowMemory_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowMemory_KEYS.Marshal(b, m, deterministic)
}
func (m *RcmdShowMemory_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowMemory_KEYS.Merge(m, src)
}
func (m *RcmdShowMemory_KEYS) XXX_Size() int {
	return xxx_messageInfo_RcmdShowMemory_KEYS.Size(m)
}
func (m *RcmdShowMemory_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowMemory_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowMemory_KEYS proto.InternalMessageInfo

type RcmdShowMemoryInfo struct {
	StructureName        string   `protobuf:"bytes,1,opt,name=structure_name,json=structureName,proto3" json:"structure_name,omitempty"`
	Size                 uint32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	CurrentCount         uint32   `protobuf:"varint,3,opt,name=current_count,json=currentCount,proto3" json:"current_count,omitempty"`
	AllocFails           uint32   `protobuf:"varint,4,opt,name=alloc_fails,json=allocFails,proto3" json:"alloc_fails,omitempty"`
	AllocCount           uint32   `protobuf:"varint,5,opt,name=alloc_count,json=allocCount,proto3" json:"alloc_count,omitempty"`
	FreedCount           uint32   `protobuf:"varint,6,opt,name=freed_count,json=freedCount,proto3" json:"freed_count,omitempty"`
	MemoryType           string   `protobuf:"bytes,7,opt,name=memory_type,json=memoryType,proto3" json:"memory_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowMemoryInfo) Reset()         { *m = RcmdShowMemoryInfo{} }
func (m *RcmdShowMemoryInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowMemoryInfo) ProtoMessage()    {}
func (*RcmdShowMemoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1162942eb0a682db, []int{1}
}

func (m *RcmdShowMemoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowMemoryInfo.Unmarshal(m, b)
}
func (m *RcmdShowMemoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowMemoryInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowMemoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowMemoryInfo.Merge(m, src)
}
func (m *RcmdShowMemoryInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowMemoryInfo.Size(m)
}
func (m *RcmdShowMemoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowMemoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowMemoryInfo proto.InternalMessageInfo

func (m *RcmdShowMemoryInfo) GetStructureName() string {
	if m != nil {
		return m.StructureName
	}
	return ""
}

func (m *RcmdShowMemoryInfo) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RcmdShowMemoryInfo) GetCurrentCount() uint32 {
	if m != nil {
		return m.CurrentCount
	}
	return 0
}

func (m *RcmdShowMemoryInfo) GetAllocFails() uint32 {
	if m != nil {
		return m.AllocFails
	}
	return 0
}

func (m *RcmdShowMemoryInfo) GetAllocCount() uint32 {
	if m != nil {
		return m.AllocCount
	}
	return 0
}

func (m *RcmdShowMemoryInfo) GetFreedCount() uint32 {
	if m != nil {
		return m.FreedCount
	}
	return 0
}

func (m *RcmdShowMemoryInfo) GetMemoryType() string {
	if m != nil {
		return m.MemoryType
	}
	return ""
}

type RcmdShowEdmMemoryInfo struct {
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Total                uint32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Success              uint32   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Failure              uint32   `protobuf:"varint,4,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowEdmMemoryInfo) Reset()         { *m = RcmdShowEdmMemoryInfo{} }
func (m *RcmdShowEdmMemoryInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowEdmMemoryInfo) ProtoMessage()    {}
func (*RcmdShowEdmMemoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1162942eb0a682db, []int{2}
}

func (m *RcmdShowEdmMemoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowEdmMemoryInfo.Unmarshal(m, b)
}
func (m *RcmdShowEdmMemoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowEdmMemoryInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowEdmMemoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowEdmMemoryInfo.Merge(m, src)
}
func (m *RcmdShowEdmMemoryInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowEdmMemoryInfo.Size(m)
}
func (m *RcmdShowEdmMemoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowEdmMemoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowEdmMemoryInfo proto.InternalMessageInfo

func (m *RcmdShowEdmMemoryInfo) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RcmdShowEdmMemoryInfo) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RcmdShowEdmMemoryInfo) GetSuccess() uint32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *RcmdShowEdmMemoryInfo) GetFailure() uint32 {
	if m != nil {
		return m.Failure
	}
	return 0
}

type RcmdShowStringMemoryInfo struct {
	Size                 uint32   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	Total                uint32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Success              uint32   `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Failure              uint32   `protobuf:"varint,4,opt,name=failure,proto3" json:"failure,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowStringMemoryInfo) Reset()         { *m = RcmdShowStringMemoryInfo{} }
func (m *RcmdShowStringMemoryInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowStringMemoryInfo) ProtoMessage()    {}
func (*RcmdShowStringMemoryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1162942eb0a682db, []int{3}
}

func (m *RcmdShowStringMemoryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowStringMemoryInfo.Unmarshal(m, b)
}
func (m *RcmdShowStringMemoryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowStringMemoryInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowStringMemoryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowStringMemoryInfo.Merge(m, src)
}
func (m *RcmdShowStringMemoryInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowStringMemoryInfo.Size(m)
}
func (m *RcmdShowStringMemoryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowStringMemoryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowStringMemoryInfo proto.InternalMessageInfo

func (m *RcmdShowStringMemoryInfo) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RcmdShowStringMemoryInfo) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RcmdShowStringMemoryInfo) GetSuccess() uint32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *RcmdShowStringMemoryInfo) GetFailure() uint32 {
	if m != nil {
		return m.Failure
	}
	return 0
}

type RcmdShowMemory struct {
	MemoryInfo           []*RcmdShowMemoryInfo       `protobuf:"bytes,50,rep,name=memory_info,json=memoryInfo,proto3" json:"memory_info,omitempty"`
	EdmMemoryInfo        []*RcmdShowEdmMemoryInfo    `protobuf:"bytes,51,rep,name=edm_memory_info,json=edmMemoryInfo,proto3" json:"edm_memory_info,omitempty"`
	StringMemoryInfo     []*RcmdShowStringMemoryInfo `protobuf:"bytes,52,rep,name=string_memory_info,json=stringMemoryInfo,proto3" json:"string_memory_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RcmdShowMemory) Reset()         { *m = RcmdShowMemory{} }
func (m *RcmdShowMemory) String() string { return proto.CompactTextString(m) }
func (*RcmdShowMemory) ProtoMessage()    {}
func (*RcmdShowMemory) Descriptor() ([]byte, []int) {
	return fileDescriptor_1162942eb0a682db, []int{4}
}

func (m *RcmdShowMemory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowMemory.Unmarshal(m, b)
}
func (m *RcmdShowMemory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowMemory.Marshal(b, m, deterministic)
}
func (m *RcmdShowMemory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowMemory.Merge(m, src)
}
func (m *RcmdShowMemory) XXX_Size() int {
	return xxx_messageInfo_RcmdShowMemory.Size(m)
}
func (m *RcmdShowMemory) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowMemory.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowMemory proto.InternalMessageInfo

func (m *RcmdShowMemory) GetMemoryInfo() []*RcmdShowMemoryInfo {
	if m != nil {
		return m.MemoryInfo
	}
	return nil
}

func (m *RcmdShowMemory) GetEdmMemoryInfo() []*RcmdShowEdmMemoryInfo {
	if m != nil {
		return m.EdmMemoryInfo
	}
	return nil
}

func (m *RcmdShowMemory) GetStringMemoryInfo() []*RcmdShowStringMemoryInfo {
	if m != nil {
		return m.StringMemoryInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*RcmdShowMemory_KEYS)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.memory.rcmd_show_memory_KEYS")
	proto.RegisterType((*RcmdShowMemoryInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.memory.rcmd_show_memory_info")
	proto.RegisterType((*RcmdShowEdmMemoryInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.memory.rcmd_show_edm_memory_info")
	proto.RegisterType((*RcmdShowStringMemoryInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.memory.rcmd_show_string_memory_info")
	proto.RegisterType((*RcmdShowMemory)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.memory.rcmd_show_memory")
}

func init() { proto.RegisterFile("rcmd_show_memory.proto", fileDescriptor_1162942eb0a682db) }

var fileDescriptor_1162942eb0a682db = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x4b, 0x4b, 0xec, 0x30,
	0x18, 0xa5, 0xf3, 0xe4, 0x66, 0x6e, 0xef, 0x1d, 0xc2, 0x7d, 0x44, 0x10, 0x1c, 0x2a, 0x42, 0x57,
	0x5d, 0xcc, 0xb8, 0x77, 0x31, 0x38, 0x20, 0xa2, 0x8b, 0xea, 0xc6, 0x55, 0xac, 0xe9, 0x57, 0x2d,
	0xb6, 0x4d, 0x49, 0x52, 0xb4, 0xe2, 0xff, 0xf4, 0xc7, 0xb8, 0x91, 0xa6, 0xaf, 0xb1, 0x33, 0x0b,
	0x99, 0x85, 0xbb, 0x7c, 0xe7, 0x3b, 0x9c, 0x93, 0x73, 0x9a, 0xa2, 0x7f, 0x82, 0xc5, 0x3e, 0x95,
	0x0f, 0xfc, 0x89, 0xc6, 0x10, 0x73, 0x91, 0x3b, 0xa9, 0xe0, 0x8a, 0x63, 0x9b, 0x85, 0x92, 0x71,
	0x1a, 0x72, 0x49, 0x9f, 0x05, 0x0d, 0x93, 0x40, 0x78, 0x54, 0x53, 0x79, 0x0a, 0xc2, 0x29, 0x4e,
	0x4e, 0xc9, 0xb7, 0xfe, 0xa3, 0xbf, 0x5d, 0x0d, 0x7a, 0x7e, 0x7a, 0x73, 0x65, 0xbd, 0x1b, 0x5b,
	0x36, 0x61, 0x12, 0x70, 0x7c, 0x84, 0x7e, 0x49, 0x25, 0x32, 0xa6, 0x32, 0x01, 0x34, 0xf1, 0x62,
	0x20, 0xc6, 0xcc, 0xb0, 0x7f, 0xb8, 0x66, 0x83, 0x5e, 0x7a, 0x31, 0x60, 0x8c, 0x06, 0x32, 0x7c,
	0x01, 0xd2, 0x9b, 0x19, 0xb6, 0xe9, 0xea, 0x33, 0x3e, 0x44, 0x26, 0xcb, 0x84, 0x80, 0x44, 0x51,
	0xc6, 0xb3, 0x44, 0x91, 0xbe, 0x5e, 0xfe, 0xac, 0xc0, 0x65, 0x81, 0xe1, 0x03, 0x34, 0xf1, 0xa2,
	0x88, 0x33, 0x1a, 0x78, 0x61, 0x24, 0xc9, 0x40, 0x53, 0x90, 0x86, 0x56, 0x05, 0xd2, 0x12, 0x4a,
	0x8d, 0xe1, 0x1a, 0xa1, 0x51, 0x08, 0x04, 0x80, 0x5f, 0x11, 0x46, 0x25, 0x41, 0x43, 0x0d, 0xa1,
	0x4a, 0xa4, 0xf2, 0x14, 0xc8, 0x58, 0xdf, 0x1f, 0x95, 0xd0, 0x75, 0x9e, 0x82, 0x95, 0xa3, 0xbd,
	0x36, 0x3c, 0xf8, 0xf1, 0xa7, 0x02, 0xea, 0x64, 0xc6, 0x5a, 0xb2, 0x3f, 0x68, 0xa8, 0xb8, 0xf2,
	0xa2, 0x2a, 0x6e, 0x39, 0x60, 0x82, 0xc6, 0x32, 0x63, 0x0c, 0xa4, 0xac, 0x92, 0xd6, 0x63, 0xb1,
	0x29, 0xe2, 0x65, 0x02, 0xaa, 0x80, 0xf5, 0x68, 0xbd, 0xa2, 0xfd, 0xd6, 0x5a, 0x2a, 0x11, 0x26,
	0xf7, 0xdf, 0xe8, 0xfe, 0xd6, 0x43, 0xd3, 0xee, 0x67, 0xc7, 0xb7, 0x4d, 0x5d, 0xc5, 0x0d, 0xc8,
	0x7c, 0xd6, 0xb7, 0x27, 0xf3, 0x13, 0xe7, 0xab, 0x8f, 0xcc, 0xd9, 0xfa, 0x8e, 0xea, 0xbe, 0xcf,
	0x8a, 0x50, 0x8f, 0xe8, 0x77, 0xa7, 0x65, 0xb2, 0xd0, 0x2e, 0xcb, 0x5d, 0x5c, 0x3a, 0x52, 0xae,
	0x09, 0x7e, 0x7c, 0xd1, 0x9a, 0x29, 0x84, 0x37, 0x7b, 0x25, 0xc7, 0xda, 0x6f, 0xb5, 0x8b, 0xdf,
	0xa6, 0x9a, 0x3b, 0x2d, 0xb1, 0xd6, 0xf5, 0x6e, 0xa4, 0x7f, 0xcd, 0xc5, 0x47, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x65, 0x26, 0xba, 0x1d, 0xb4, 0x03, 0x00, 0x00,
}
