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
// source: snmp_mib_info.proto

package cisco_ios_xr_snmp_agent_oper_snmp_information_mibs_mib_mib_information

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

type SnmpMibInfo_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpMibInfo_KEYS) Reset()         { *m = SnmpMibInfo_KEYS{} }
func (m *SnmpMibInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*SnmpMibInfo_KEYS) ProtoMessage()    {}
func (*SnmpMibInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3908566ddefab5, []int{0}
}

func (m *SnmpMibInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpMibInfo_KEYS.Unmarshal(m, b)
}
func (m *SnmpMibInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpMibInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *SnmpMibInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpMibInfo_KEYS.Merge(m, src)
}
func (m *SnmpMibInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_SnmpMibInfo_KEYS.Size(m)
}
func (m *SnmpMibInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpMibInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpMibInfo_KEYS proto.InternalMessageInfo

func (m *SnmpMibInfo_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SnmpMibInfo struct {
	MibName              string   `protobuf:"bytes,50,opt,name=mib_name,json=mibName,proto3" json:"mib_name,omitempty"`
	DllName              string   `protobuf:"bytes,51,opt,name=dll_name,json=dllName,proto3" json:"dll_name,omitempty"`
	MibConfigFilename    string   `protobuf:"bytes,52,opt,name=mib_config_filename,json=mibConfigFilename,proto3" json:"mib_config_filename,omitempty"`
	IsMibLoaded          bool     `protobuf:"varint,53,opt,name=is_mib_loaded,json=isMibLoaded,proto3" json:"is_mib_loaded,omitempty"`
	DllCapabilities      uint32   `protobuf:"varint,54,opt,name=dll_capabilities,json=dllCapabilities,proto3" json:"dll_capabilities,omitempty"`
	TrapStrings          string   `protobuf:"bytes,55,opt,name=trap_strings,json=trapStrings,proto3" json:"trap_strings,omitempty"`
	Timeout              bool     `protobuf:"varint,56,opt,name=timeout,proto3" json:"timeout,omitempty"`
	LoadTime             uint32   `protobuf:"varint,57,opt,name=load_time,json=loadTime,proto3" json:"load_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SnmpMibInfo) Reset()         { *m = SnmpMibInfo{} }
func (m *SnmpMibInfo) String() string { return proto.CompactTextString(m) }
func (*SnmpMibInfo) ProtoMessage()    {}
func (*SnmpMibInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3908566ddefab5, []int{1}
}

func (m *SnmpMibInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnmpMibInfo.Unmarshal(m, b)
}
func (m *SnmpMibInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnmpMibInfo.Marshal(b, m, deterministic)
}
func (m *SnmpMibInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnmpMibInfo.Merge(m, src)
}
func (m *SnmpMibInfo) XXX_Size() int {
	return xxx_messageInfo_SnmpMibInfo.Size(m)
}
func (m *SnmpMibInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SnmpMibInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SnmpMibInfo proto.InternalMessageInfo

func (m *SnmpMibInfo) GetMibName() string {
	if m != nil {
		return m.MibName
	}
	return ""
}

func (m *SnmpMibInfo) GetDllName() string {
	if m != nil {
		return m.DllName
	}
	return ""
}

func (m *SnmpMibInfo) GetMibConfigFilename() string {
	if m != nil {
		return m.MibConfigFilename
	}
	return ""
}

func (m *SnmpMibInfo) GetIsMibLoaded() bool {
	if m != nil {
		return m.IsMibLoaded
	}
	return false
}

func (m *SnmpMibInfo) GetDllCapabilities() uint32 {
	if m != nil {
		return m.DllCapabilities
	}
	return 0
}

func (m *SnmpMibInfo) GetTrapStrings() string {
	if m != nil {
		return m.TrapStrings
	}
	return ""
}

func (m *SnmpMibInfo) GetTimeout() bool {
	if m != nil {
		return m.Timeout
	}
	return false
}

func (m *SnmpMibInfo) GetLoadTime() uint32 {
	if m != nil {
		return m.LoadTime
	}
	return 0
}

func init() {
	proto.RegisterType((*SnmpMibInfo_KEYS)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.mibs.mib.mib_information.snmp_mib_info_KEYS")
	proto.RegisterType((*SnmpMibInfo)(nil), "cisco_ios_xr_snmp_agent_oper.snmp.information.mibs.mib.mib_information.snmp_mib_info")
}

func init() { proto.RegisterFile("snmp_mib_info.proto", fileDescriptor_3d3908566ddefab5) }

var fileDescriptor_3d3908566ddefab5 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x84, 0x68, 0xeb, 0x52, 0x01, 0xee, 0x62, 0xc4, 0x52, 0x3a, 0x85, 0xc5, 0x03,
	0xe5, 0x73, 0xae, 0xe8, 0xc2, 0xc7, 0x90, 0xb2, 0x30, 0x59, 0x76, 0xe2, 0x44, 0x27, 0xf9, 0x23,
	0xb2, 0x8d, 0xc4, 0xaf, 0xe1, 0xb7, 0x22, 0x5f, 0x5a, 0x89, 0x2e, 0xd1, 0xdd, 0xfb, 0x3c, 0xb9,
	0x37, 0x52, 0xc8, 0x3c, 0x3a, 0xdb, 0x0b, 0x0b, 0x4a, 0x80, 0x6b, 0x3d, 0xef, 0x83, 0x4f, 0x9e,
	0x6e, 0x6a, 0x88, 0xb5, 0x17, 0xe0, 0xa3, 0xf8, 0x09, 0x02, 0x0d, 0xd9, 0x69, 0x97, 0x84, 0xef,
	0x75, 0xe0, 0x79, 0xe7, 0xd9, 0x0e, 0x56, 0x26, 0xf0, 0x8e, 0x5b, 0x50, 0x31, 0x3f, 0xf8, 0xfe,
	0xcc, 0x0e, 0x2c, 0x4b, 0x42, 0x0f, 0xce, 0x8b, 0xd7, 0x97, 0xaf, 0x2d, 0xa5, 0xe4, 0xd8, 0x49,
	0xab, 0x59, 0xb1, 0x28, 0xca, 0x49, 0x85, 0xf3, 0xf2, 0xf7, 0x88, 0xcc, 0x0e, 0x54, 0x7a, 0x49,
	0xc6, 0x79, 0x46, 0xf3, 0x16, 0xcd, 0x91, 0x05, 0xf5, 0x21, 0xad, 0xce, 0xa8, 0x31, 0x66, 0x40,
	0xab, 0x01, 0x35, 0xc6, 0x20, 0xe2, 0x64, 0x9e, 0xdf, 0xaa, 0xbd, 0x6b, 0xa1, 0x13, 0x2d, 0x18,
	0x8d, 0xd6, 0x1d, 0x5a, 0x17, 0x16, 0xd4, 0x1a, 0xc9, 0x66, 0x07, 0xe8, 0x92, 0xcc, 0x20, 0x62,
	0xa9, 0xf1, 0xb2, 0xd1, 0x0d, 0xbb, 0x5f, 0x14, 0xe5, 0xb8, 0x9a, 0x42, 0x7c, 0x07, 0xf5, 0x86,
	0x11, 0xbd, 0x21, 0xe7, 0xb9, 0xae, 0x96, 0xbd, 0x54, 0x60, 0x20, 0x81, 0x8e, 0xec, 0x61, 0x51,
	0x94, 0xb3, 0xea, 0xac, 0x31, 0x66, 0xfd, 0x2f, 0xa6, 0xd7, 0xe4, 0x34, 0x05, 0xd9, 0x8b, 0x98,
	0x02, 0xb8, 0x2e, 0xb2, 0x47, 0xec, 0x9d, 0xe6, 0x6c, 0x3b, 0x44, 0x94, 0x91, 0x51, 0x02, 0xab,
	0xfd, 0x77, 0x62, 0x4f, 0xd8, 0xb5, 0x5f, 0xe9, 0x15, 0x99, 0xe4, 0x8f, 0x10, 0x79, 0x67, 0xcf,
	0x58, 0x30, 0xce, 0xc1, 0x27, 0x58, 0xad, 0x4e, 0xf0, 0xcf, 0xac, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xce, 0x8f, 0x8a, 0x5b, 0xb0, 0x01, 0x00, 0x00,
}
